+++
title = "上下文"
weight = 40
date = 2023-07-09T21:50:09+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Context - 上下文

> 原文：[https://echo.labstack.com/docs/context](https://echo.labstack.com/docs/context)

​	`echo.Context` 表示当前 HTTP 请求的上下文。它保存了请求和响应的引用、路径、路径参数、数据、已注册的处理程序以及读取请求和写入响应的 API。由于 Context 是一个接口，可以轻松地通过自定义 API 扩展它。

## 扩展

**定义一个自定义context**

```go
type CustomContext struct {
    echo.Context
}

func (c *CustomContext) Foo() {
    println("foo")
}

func (c *CustomContext) Bar() {
    println("bar")
}
```



**创建一个中间件来扩展默认context**

```go
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cc := &CustomContext{c}
        return next(cc)
    }
})
```



> 注意事项
>
> ​	这个中间件应该在其他中间件之前注册。

> 注意事项
>
> ​	在路由运行之前（Pre）不能在中间件中定义自定义上下文。

**在处理程序中使用**

```go
e.GET("/", func(c echo.Context) error {
    cc := c.(*CustomContext)
    cc.Foo()
    cc.Bar()
    return cc.String(200, "OK")
})
```



## 并发

> 注意事项
>
> ​	`Context` 不能在处理请求的 goroutine 之外被访问。原因有两个： 
>
> 1. `Context` 具有从多个 goroutine 执行时容易出现问题的函数。因此，只能有一个 goroutine 访问它。
> 2. Echo 使用池来创建 `Context`。当请求处理完成后，Echo 将 `Context` 返回到池中。
>
> ​	关于这个原因，参考 issue [1908](https://github.com/labstack/echo/issues/1908) 中发生的“警示故事（cautionary tale）”。并发是复杂的。在使用 goroutine 时要注意这个陷阱（pitfall）。

### 解决方案

​	使用通道

```go
func(c echo.Context) error {
    ca := make(chan string, 1) // 为了防止该通道阻塞，大小设置为1。
    r := c.Request()
    method := r.Method

    go func() {
        // 此函数不能操作 Context。

        fmt.Printf("Method: %s\n", method)

        // 执行一些耗时操作...

        ca <- "Hey!"
    }()

    select {
    case result := <-ca:
        return c.String(http.StatusOK, "Result: "+result)
    case <-c.Request().Context().Done(): // 检查上下文。
        // 如果到达这里，说明上下文已取消（达到了超时等）。
        return nil
    }
}
```



