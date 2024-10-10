+++
title = "错误处理"
weight = 60
date = 2023-07-09T21:50:29+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Error Handling - 错误处理

> 原文：[https://echo.labstack.com/docs/error-handling](https://echo.labstack.com/docs/error-handling)

​	Echo 提倡通过从中间件和处理程序返回错误来进行集中的 HTTP 错误处理。集中的错误处理程序允许我们从统一的位置将错误日志记录到外部服务，并向客户端发送自定义的 HTTP 响应。

​	您可以返回标准的 `error` 或 `echo.*HTTPError`。

​	例如，当基本身份验证（basic auth）中间件发现无效的凭证（credentials）时，它会返回 `401 - Unauthorized`的错误，并中止（aborting）当前的 HTTP 请求。

```go
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    // 从 HTTP 请求头提取凭证并进行安全性检查

    // 对于无效凭证
    return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

    // 对于有效凭证，调用 next
    // return next(c)
  }
})
```



​	您还可以使用没有消息的 `echo.NewHTTPError()`，在这种情况下，状态文本将被用作错误消息。例如，"Unauthorized"。

## 默认的 HTTP 错误处理程序

​	Echo 提供了一个默认的 HTTP 错误处理程序，它以 JSON 格式发送错误。

```js
{
  "message": "error connecting to redis"
}
```



​	对于标准的 `error`，响应将作为 `500 - Internal Server Error` 发送；但是，如果您在调试模式下运行，则发送原始错误消息。如果错误是 `*HTTPError`，则使用提供的状态码和消息发送响应。如果日志记录开启，则还会记录错误消息。

## 自定义的 HTTP 错误处理程序

​	可以通过 `e.HTTPErrorHandler` 设置自定义的 HTTP 错误处理程序。

​	对于大多数情况，默认的错误 HTTP 处理程序应该足够；然而，如果您想捕获不同类型的错误并采取相应的操作，例如发送通知电子邮件或将错误记录到集中（centralized）系统，自定义的 HTTP 错误处理程序可以派上用场。您还可以向客户端发送自定义的响应，例如错误页面或仅 JSON 响应。

### 错误页面

​	以下自定义的 HTTP 错误处理程序显示了如何为不同类型的错误显示错误页面并记录错误。错误页面的名称应该类似于 `<CODE>.html`，例如 `500.html`。您可以查看这个项目 [https://github.com/AndiDittrich/HttpErrorPages](https://github.com/AndiDittrich/HttpErrorPages) 来获取预构建（pre-built）的错误页面。

```go
func customHTTPErrorHandler(err error, c echo.Context) {
    code := http.StatusInternalServerError
    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
    }
    c.Logger().Error(err)
    errorPage := fmt.Sprintf("%d.html", code)
    if err := c.File(errorPage); err != nil {
        c.Logger().Error(err)
    }
}

e.HTTPErrorHandler = customHTTPErrorHandler
```



> 提示
>
> ​	您可以将日志写入外部服务，如 Elasticsearch 或 Splunk，而不是写入日志记录器（logger）。