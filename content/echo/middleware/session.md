+++
title = "session"
weight = 220
date = 2023-07-09T21:58:02+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Session

> 原文：[https://echo.labstack.com/docs/middleware/session](https://echo.labstack.com/docs/middleware/session)

Session middleware facilitates HTTP session management backed by [gorilla sessions](https://github.com/gorilla/sessions). The default implementation provides cookie and filesystem based session store; however, you can take advantage of [community maintained implementation](https://github.com/gorilla/sessions#store-implementations) for various backends.

NOTE

Echo community contribution

## Dependencies

```go
import (
  "github.com/gorilla/sessions"
  "github.com/labstack/echo-contrib/session"
)
```



## Usage

```go
e := echo.New()
e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

e.GET("/", func(c echo.Context) error {
  sess, _ := session.Get("session", c)
  sess.Options = &sessions.Options{
    Path:     "/",
    MaxAge:   86400 * 7,
    HttpOnly: true,
  }
  sess.Values["foo"] = "bar"
  sess.Save(c.Request(), c.Response())
  return c.NoContent(http.StatusOK)
})
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(session.MiddlewareWithConfig(session.Config{}))
```



## Configuration

```go
Config struct {
  // Skipper defines a function to skip middleware.
  Skipper middleware.Skipper

  // Session store.
  // Required.
  Store sessions.Store
}
```



### Default Configuration

```go
DefaultConfig = Config{
  Skipper: DefaultSkipper,
}
```