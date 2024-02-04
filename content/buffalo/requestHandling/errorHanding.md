+++
title = "错误处理"
date = 2024-02-04T21:08:36+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/request_handling/errors/](https://gobuffalo.io/documentation/request_handling/errors/)

# Error Handling 错误处理 

An `error` is Go way to tell something went wrong. In this chapter, you’ll learn how to return errors from a route handler and how Buffalo will catch any non-handled error.

​	 `error` 是 Go 中表示出现错误的方式。在本章中，您将学习如何从路由处理程序返回错误，以及 Buffalo 将如何捕获任何未处理的错误。

## Returning Errors From a Handler 从处理程序返回错误 

The easiest way to produce an error response is to return a standard Go error:

​	生成错误响应的最简单方法是返回一个标准的 Go 错误：

```go
func MyHandler(c buffalo.Context) error {
  // Return any go error, this will result in a 500 status code.
  return errors.New("boom!")
}
```

A `nil` error will produce a raw HTTP 200 response:

​	 `nil` 错误将生成一个原始的 HTTP 200 响应：

```go
func MyHandler(c buffalo.Context) error {
  // HTTP 200
  return nil
}
```

If you need to customize the error message or the HTTP code, use the [Error](https://pkg.go.dev/github.com/gobuffalo/buffalo#DefaultContext.Error) method:

​	如果您需要自定义错误消息或 HTTP 代码，请使用 Error 方法：

```go
func MyHandler(c buffalo.Context) error {
  // Use the Error function on the context.
  // This will result in a status code of 401.
  return c.Error(401, errors.New("Unauthorized!"))
}
```

## Default Error Handling (Development) 默认错误处理（开发） 

In “development” mode (`GO_ENV=development`), Buffalo will generate some helpful errors pages for you.

​	在“开发”模式（ `GO_ENV=development` ）中，Buffalo 将为您生成一些有用的错误页面。

If you use a JSON or an XML content type, the error is returned in the proper type:

​	如果您使用 JSON 或 XML 内容类型，则错误将以适当的类型返回：

```json
{
  "error": "could not find test/",
  "trace": "could not find test/\ngithub.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).fileServer.func1\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/route_mappings.go:97\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\nnet/http.StripPrefix.func1\n\t/usr/local/go/src/net/http/server.go:1986\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\ngithub.com/gobuffalo/docs/vendor/github.com/gorilla/mux.(*Router).ServeHTTP\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gorilla/mux/mux.go:162\ngithub.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web.ErrorChecker.func1\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web/web.go:23\nnet/http.HandlerFunc.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:1947\ngithub.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).ServeHTTP\n\t/home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/server.go:127\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2694\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1830\nruntime.goexit\n\t/usr/local/go/src/runtime/asm_amd64.s:2361",
  "code": 404
}
<response code="404">
  <error>could not find test/</error>
  <trace>could not find test/ github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).fileServer.func1 /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/route_mappings.go:97 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 net/http.StripPrefix.func1 /usr/local/go/src/net/http/server.go:1986 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 github.com/gobuffalo/docs/vendor/github.com/gorilla/mux.(*Router).ServeHTTP /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gorilla/mux/mux.go:162 github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web.ErrorChecker.func1 /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/markbates/refresh/refresh/web/web.go:23 net/http.HandlerFunc.ServeHTTP /usr/local/go/src/net/http/server.go:1947 github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo.(*App).ServeHTTP /home/michalakst/go/src/github.com/gobuffalo/docs/vendor/github.com/gobuffalo/buffalo/server.go:127 net/http.serverHandler.ServeHTTP /usr/local/go/src/net/http/server.go:2694 net/http.(*conn).serve /usr/local/go/src/net/http/server.go:1830 runtime.goexit /usr/local/go/src/runtime/asm_amd64.s:2361</trace>
</response>
```

In “production” mode (`GO_ENV=production`), Buffalo will not generate pages that have developer style information, because this would give precious information to hackers. Instead the pages are simpler.

​	在“生产”模式（ `GO_ENV=production` ）中，Buffalo 不会生成包含开发人员样式信息的页面，因为这会向黑客提供宝贵的信息。相反，这些页面更简单。

## Custom Error Handling 自定义错误处理 

While Buffalo will handle errors for you out of the box, it can be useful to handle errors in a custom way. To accomplish this, Buffalo allows for the mapping of HTTP status codes to specific handlers. This means the error can be dealt with in a custom fashion.

​	虽然 Buffalo 会开箱即用地处理错误，但以自定义方式处理错误可能很有用。为了实现这一点，Buffalo 允许将 HTTP 状态码映射到特定处理程序。这意味着可以以自定义方式处理错误。

```go
app = buffalo.New(buffalo.Options{
  Env: ENV,
})

// We associate the HTTP 422 status to a specific handler.
// All the other status code will still use the default handler provided by Buffalo.
app.ErrorHandlers[422] = func(status int, err error, c buffalo.Context) error {
  res := c.Response()
  res.WriteHeader(422)
  res.Write([]byte(fmt.Sprintf("Oops!! There was an error: %s", err.Error())))
  return nil
}

app.GET("/oops", MyHandler)

func MyHandler(c buffalo.Context) error {
  return c.Error(422, errors.New("Oh no!"))
}
GET /oops -> [422] Oops!! There was an error: Oh no!
```

In the above example any error from your application that returns a status of `422` will be caught by the custom handler and will be dealt with accordingly.

​	在上面的示例中，应用程序中任何返回状态 `422` 的错误都将被自定义处理程序捕获并相应处理。