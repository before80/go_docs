+++
title = "ExpVar"
date = 2024-02-05T09:14:15+08:00
weight = 100
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/expvar]({{< ref "/fiber/API/Middleware/ExpVar" >}})

# ExpVar

Expvar middleware for [Fiber](https://github.com/gofiber/fiber) that serves via its HTTP server runtime exposed variants in the JSON format. The package is typically only imported for the side effect of registering its HTTP handlers. The handled path is `/debug/vars`.

​	适用于 Fiber 的 Expvar 中间件，通过其 HTTP 服务器运行时以 JSON 格式提供公开的变量。通常仅导入该软件包以注册其 HTTP 处理程序。处理的路径是 `/debug/vars` 。

## Signatures 签名

```go
func New() fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  expvarmw "github.com/gofiber/fiber/v2/middleware/expvar"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
var count = expvar.NewInt("count")

app.Use(expvarmw.New())
app.Get("/", func(c *fiber.Ctx) error {
    count.Add(1)

    return c.SendString(fmt.Sprintf("hello expvar count %d", count.Value()))
})
```



Visit path `/debug/vars` to see all vars and use query `r=key` to filter exposed variables.

​	访问路径 `/debug/vars` 以查看所有变量并使用查询 `r=key` 过滤公开的变量。

```bash
curl 127.0.0.1:3000
hello expvar count 1

curl 127.0.0.1:3000/debug/vars
{
    "cmdline": ["xxx"],
    "count": 1,
    "expvarHandlerCalls": 33,
    "expvarRegexpErrors": 0,
    "memstats": {...}
}

curl 127.0.0.1:3000/debug/vars?r=c
{
    "cmdline": ["xxx"],
    "count": 1
}
```



## Config 配置

| Property 属性 | Type 输入               | Description 说明                                             | Default 默认 |
| ------------- | ----------------------- | ------------------------------------------------------------ | ------------ |
| Next 下一步   | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`        |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next: nil,
}
```
