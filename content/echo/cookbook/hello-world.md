+++
title = "hello-world"
date = 2023-07-09T22:03:02+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Hello World

> 原文：[https://echo.labstack.com/docs/cookbook/hello-world](https://echo.labstack.com/docs/cookbook/hello-world)

## Server

cookbook/hello-world/server.go

```go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
```