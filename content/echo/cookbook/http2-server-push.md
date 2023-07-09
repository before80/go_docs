+++
title = "http2-server-push"
date = 2023-07-09T22:03:17+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# HTTP/2 Server Push

https://echo.labstack.com/docs/cookbook/http2-server-push

NOTE

Requires go1.8+

## Send web assets using HTTP/2 server push

### [Generate a self-signed X.509 TLS certificate](https://echo.labstack.com/docs/cookbook/http2#step-1-generate-a-self-signed-x-509-tls-certificate)

### 1) Register a route to serve web assets

```go
e.Static("/", "static")
```



### 2) Create a handler to serve index.html and push it's dependencies

```go
e.GET("/", func(c echo.Context) (err error) {
  pusher, ok := c.Response().Writer.(http.Pusher)
  if ok {
    if err = pusher.Push("/app.css", nil); err != nil {
      return
    }
    if err = pusher.Push("/app.js", nil); err != nil {
      return
    }
    if err = pusher.Push("/echo.png", nil); err != nil {
      return
    }
  }
  return c.File("index.html")
})
```



INFO

If `http.Pusher` is supported, web assets are pushed; otherwise, client makes separate requests to get them.

### 3) Start TLS server using cert.pem and key.pem

```go
if err := e.StartTLS(":1323", "cert.pem", "key.pem"); err != http.ErrServerClosed {
  log.Fatal(err)
}
```



or use customized HTTP server with your own TLSConfig

```go
s := http.Server{
  Addr:    ":8443",
  Handler: e, // set Echo as handler
  TLSConfig: &tls.Config{
    //Certificates: nil, // <-- s.ListenAndServeTLS will populate this field
  },
  //ReadTimeout: 30 * time.Second, // use custom timeouts
}
if err := s.ListenAndServeTLS("cert.pem", "key.pem"); err != http.ErrServerClosed {
  log.Fatal(err)
}
```



### 4) Start the server and browse to https://localhost:1323

```sh
Protocol: HTTP/2.0
Host: localhost:1323
Remote Address: [::1]:60288
Method: GET
Path: /
```



## Source Code

cookbook/http2-server-push/index.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>HTTP/2 Server Push</title>
  <link rel="stylesheet" href="/app.css">
  <script src="/app.js"></script>
</head>
<body>
  <img class="echo" src="/echo.png">
  <h2>The following static files are served via HTTP/2 server push</h2>
  <ul>
    <li><code>/app.css</code></li>
    <li><code>/app.js</code></li>
    <li><code>/echo.png</code></li>
  </ul>
</body>
</html>
```



cookbook/http2-server-push/server.go

```go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "static")
	e.GET("/", func(c echo.Context) (err error) {
		pusher, ok := c.Response().Writer.(http.Pusher)
		if ok {
			if err = pusher.Push("/app.css", nil); err != nil {
				return
			}
			if err = pusher.Push("/app.js", nil); err != nil {
				return
			}
			if err = pusher.Push("/echo.png", nil); err != nil {
				return
			}
		}
		return c.File("index.html")
	})
	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}
```