+++
title = "timeout"
date = 2023-07-09T22:05:04+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Timeout

https://echo.labstack.com/docs/cookbook/timeout

## Server

cookbook/timeout/server.go

```go
package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second,
	}))

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		time.Sleep(10 * time.Second)
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
```