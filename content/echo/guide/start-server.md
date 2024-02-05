+++
title = "启动服务器"
weight = 70
date = 2023-07-09T21:50:57+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Start Server - 启动服务器

> 原文：[https://echo.labstack.com/docs/start-server](https://echo.labstack.com/docs/start-server)

​	Echo 提供以下方便的方法来启动服务器： 

- `Echo.Start(address string)`
- `Echo.StartTLS(address string, certFile, keyFile interface{})`
- `Echo.StartAutoTLS(address string)`
- `Echo.StartH2CServer(address string, h2s *http2.Server)`
- `Echo.StartServer(s *http.Server)`

## HTTP 服务器

​	`Echo.Start` 是一个方便的方法，它使用 Echo 来启动 HTTP 服务器以处理请求。

```go
func main() {
  e := echo.New()
  // 添加中间件和路由
  // ...
  if err := e.Start(":8080"); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}
```



​	下面的示例与前面的 `Echo.Start` 等效：

```go
func main() {
  e := echo.New()
  // 添加中间件和路由
  // ...
  s := http.Server{
    Addr:        ":8080",
    Handler:     e,
    //ReadTimeout: 30 * time.Second, // customize http.Server timeouts
  }
  if err := s.ListenAndServe(); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}
```



## HTTPS 服务器

​	`Echo.StartTLS` 是一个方便的方法，它使用 Echo 来启动 HTTPS 服务器，并使用 `server.crt` 和 `server.key` 作为 TLS 证书对。

```go
func main() {
  e := echo.New()
  // 添加中间件和路由
  // ...
  if err := e.StartTLS(":8443", "server.crt", "server.key"); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}
```



​	下面的示例与前面的 `Echo.StartTLS` 等效：

```go
func main() {
  e := echo.New()
  // 添加中间件和路由
  // ...
  s := http.Server{
    Addr:    ":8443",
    Handler: e, // set Echo as handler
    TLSConfig: &tls.Config{
      //MinVersion: 1, // customize TLS configuration
    },
    //ReadTimeout: 30 * time.Second, // use custom timeouts
  }
  if err := s.ListenAndServeTLS("server.crt", "server.key"); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}
```



## 使用 Let's Encrypt 的自动 TLS 服务器

​	请参阅 [Auto TLS Recipe](https://echo.labstack.com/docs/cookbook/auto-tls#server)。

## 明文 HTTP/2 服务器 (HTTP2 over HTTP)

​	`Echo.StartH2CServer` 是一个方便的方法，它在给定的地址上启动一个自定义的明文（cleartext ） HTTP/2 服务器。

```go
func main() {
  e := echo.New()
  // 添加中间件和路由
  // ...
  s := &http2.Server{
    MaxConcurrentStreams: 250,
    MaxReadFrameSize:     1048576,
    IdleTimeout:          10 * time.Second,
  }
  if err := e.StartH2CServer(":8080", s); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}
```



​	下面的示例与前面的 `Echo.StartH2CServer` 等效：

```go
func main() {
  e := echo.New()
  // 添加中间件和路由
  // ...
  h2s := &http2.Server{
    MaxConcurrentStreams: 250,
    MaxReadFrameSize:     1048576,
    IdleTimeout:          10 * time.Second,
  }
  s := http.Server{
    Addr:    ":8080",
    Handler: h2c.NewHandler(e, h2s),
  }
  if err := s.ListenAndServe(); err != http.ErrServerClosed {
    log.Fatal(err)
  }
}
```



