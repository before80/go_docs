+++
title = "reverse-proxy"
date = 2023-07-09T22:04:22+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Reverse Proxy

https://echo.labstack.com/docs/cookbook/reverse-proxy

This recipe demonstrates how you can use Echo as a reverse proxy server and load balancer in front of your favorite applications like WordPress, Node.js, Java, Python, Ruby or even Go. For simplicity, I will use Go upstream servers with WebSocket.

## 1) Identify upstream target URL(s)

```go
url1, err := url.Parse("http://localhost:8081")
if err != nil {
  e.Logger.Fatal(err)
}
url2, err := url.Parse("http://localhost:8082")
if err != nil {
  e.Logger.Fatal(err)
}
targets := []*middleware.ProxyTarget{
  {
    URL: url1,
  },
  {
    URL: url2,
  },
}
```



## 2) Setup proxy middleware with upstream targets

In the following code snippet we are using round-robin load balancing technique. You may also use `middleware.NewRandomBalancer()`.

```go
e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))
```



To setup proxy for a sub-route use `Echo#Group()`.

```go
g := e.Group("/blog")
g.Use(middleware.Proxy(...))
```



## 3) Start upstream servers

- `cd upstream`
- `go run server.go server1 :8081`
- `go run server.go server2 :8082`

## 4) Start the proxy server

```sh
go run server.go
```



Browse to http://localhost:1323, and you should see a webpage with an HTTP request being served from "server 1" and a WebSocket request being served from "server 2."

```sh
HTTP

Hello from upstream server server1

WebSocket

Hello from upstream server server2!
Hello from upstream server server2!
Hello from upstream server server2!
```



## Source Code

cookbook/reverse-proxy/upstream/server.go

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

var index = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Upstream Server</title>
		<style>
			h1, p {
				font-weight: 300;
			}
		</style>
	</head>
	<body>
		<h1>HTTP</h1>
		<p>
			Hello from upstream server %s
		</p>
		<h1>WebSocket</h1>
		<p id="output"></p>
		<script>
			var ws = new WebSocket('ws://localhost:1323/ws')

			ws.onmessage = function(evt) {
				var out = document.getElementById('output');
				out.innerHTML += evt.data + '<br>';
			}
		</script>
	</body>
	</html>
`

func main() {
	name := os.Args[1]
	port := os.Args[2]
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf(index, name))
	})

	// WebSocket handler
	e.GET("/ws", func(c echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			for {
				// Write
				err := websocket.Message.Send(ws, fmt.Sprintf("Hello from upstream server %s!", name))
				if err != nil {
					e.Logger.Error(err)
				}
				time.Sleep(1 * time.Second)
			}
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(port))
}
```



cookbook/reverse-proxy/server.go

```go
package main

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Setup proxy
	url1, err := url.Parse("http://localhost:8081")
	if err != nil {
		e.Logger.Fatal(err)
	}
	url2, err := url.Parse("http://localhost:8082")
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
		{
			URL: url2,
		},
	}
	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Fatal(e.Start(":1323"))
}
```