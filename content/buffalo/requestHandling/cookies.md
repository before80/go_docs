+++
title = "Cookies"
date = 2024-02-04T21:08:55+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/request_handling/cookies/]({{< ref "/buffalo/requestHandling/cookies" >}})

# Cookies Cookie 

An HTTP cookie is a small piece of data that a server sends to the user’s web browser. The browser can store this data and send it back to the same server, even after the browser restart (unlike a [browser session]({{< ref "/buffalo/requestHandling/sessions" >}})).

​	HTTP cookie 是服务器发送给用户网络浏览器的少量数据。浏览器可以存储此数据，并在浏览器重新启动后将其发送回同一服务器（与浏览器会话不同）。

(HTTP) cookies are commonly used to save users state (like whether the user logged-in). See https://golang.org/pkg/net/http/#Cookie for more information on cookies in Go.

​	(HTTP) cookie 通常用于保存用户状态（例如用户是否已登录）。有关 Go 中 cookie 的更多信息，请参阅 https://golang.org/pkg/net/http/#Cookie。

## Setting a Cookie 设置 Cookie 

```go
func MyHandler(c buffalo.Context) error {
  // ...
  c.Cookies().Set("user_id", user.ID, 30 * 24 * time.Hour)
  // ...
}
```

## Setting a Cookie with Expiration 设置带过期时间的 Cookie 

```go
func MyHandler(c buffalo.Context) error {
  // ...
  exp := time.Now().Add(365 * 24 * time.Hour) // expire in 1 year
  c.Cookies().SetWithExpirationTime("user_id", user.ID, exp)
  // ...
}
```

## Setting a Cookie with Path 设置带路径的 Cookie 

```go
func MyHandler(c buffalo.Context) error {
  // ...
  c.Cookies().SetWithPath("user_id", user.ID, "/user")
  // ...
}
```

## Advanced setting a Cookie way 高级设置 Cookie 的方式 

```go
import "net/http"
func MyHandler(c buffalo.Context) error {
  // ...
  ck := http.Cookie{
    Name:    "token",
    Value:   token,
    Path:    "/",
    Expires: time.Now().Add(30 * 24 * time.Hour), // expire in 1 month
  }

  http.SetCookie(c.Response(), &ck)
  // ...
}
```

See [Cookie struct](https://golang.org/src/net/http/cookie.go) for other parameters.

​	有关其他参数，请参阅 Cookie 结构。

## Getting a Cookie 获取 Cookie 

```go
func MyHandler(c buffalo.Context) error {
  value, err := c.Cookies().Get("user_id")
  if err != nil {
    return err
  }
  return c.Render(200, r.String(value))
}
```

## Deleting a Cookie 删除 Cookie 

```go
func MyHandler(c buffalo.Context) error {
  c.Cookies().Delete("user_id")
  // ...
}
```
