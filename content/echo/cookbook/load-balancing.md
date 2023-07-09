+++
title = "load-balancing"
date = 2023-07-09T22:03:52+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Load Balancing

https://echo.labstack.com/docs/cookbook/load-balancing

This recipe demonstrates how you can use Nginx as a reverse proxy server and load balance between multiple Echo servers.

## Echo

cookbook/load-balancing/upstream/server.go

```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		<p>
			Hello from upstream server %s
		</p>
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
	e.Logger.Fatal(e.Start(port))
}
```



### Start servers

- `cd upstream`
- `go run server.go server1 :8081`
- `go run server.go server2 :8082`

## Nginx

### 1) Install Nginx

https://www.nginx.com/resources/wiki/start/topics/tutorials/install

### 2) Configure Nginx

Create a file `/etc/nginx/sites-enabled/localhost` with the following content:

```reference
https://github.com/labstack/echox/blob/master/cookbook/load-balancing/nginx.conf
```



::: info

Change listen, server_name, access_log per your need.

:::

### 3) Restart Nginx

```sh
service nginx restart
```



Browse to https://localhost:8080, and you should see a webpage being served from either "server 1" or "server 2".

```sh
Hello from upstream server server1
```