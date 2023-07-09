+++
title = "file-download"
date = 2023-07-09T22:02:07+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# File Download

https://echo.labstack.com/docs/cookbook/file-download

## Download file

### Server

cookbook/file-download/server.go

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.GET("/file", func(c echo.Context) error {
		return c.File("echo.svg")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
```



### Client

xxxxxxxxxx21 1package main2​3import (4    "net/http"5​6    "github.com/GeertJohan/go.rice"7    "github.com/labstack/echo/v4"8)9​10func main() {11    e := echo.New()12    // the file server for rice. "app" is the folder where the files come from.13    assetHandler := http.FileServer(rice.MustFindBox("app").HTTPBox())14    // serves the index.html from rice15    e.GET("/", echo.WrapHandler(assetHandler))16​17    // servers other static files18    e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))19​20    e.Logger.Fatal(e.Start(":1323"))21}go

```html
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>File download</title>
</head>
<body>

    <p>
        <a href="/file">File download</a>
    </p>

</body>
</html>
```



## Download file as inline

### Server

cookbook/file-download/inline/server.go

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.GET("/inline", func(c echo.Context) error {
		return c.Inline("inline.txt", "inline.txt")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
```



### Client

cookbook/file-download/inline/index.html

```html
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>File download</title>
</head>
<body>

    <p>
        <a href="/inline">Inline file download</a>
    </p>

</body>
</html>
```



## Download file as attachment

### Server

cookbook/file-download/attachment/server.go

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.GET("/attachment", func(c echo.Context) error {
		return c.Attachment("attachment.txt", "attachment.txt")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
```



### Client

cookbook/file-download/attachment/index.html

```html
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>File download</title>
</head>
<body>

    <p>
        <a href="/attachment">Attachment file download</a>
    </p>

</body>
</html>
```