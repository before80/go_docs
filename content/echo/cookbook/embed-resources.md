+++
title = "embed-resources"
date = 2023-07-09T22:01:58+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Embed Resources

> 原文：[https://echo.labstack.com/docs/cookbook/embed-resources](https://echo.labstack.com/docs/cookbook/embed-resources)

## With go 1.16 embed feature

cookbook/embed/server.go

```go
package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

//go:embed app
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("app"))
	}

	log.Print("using embed mode")
	fsys, err := fs.Sub(embededFiles, "app")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func main() {
	e := echo.New()
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	assetHandler := http.FileServer(getFileSystem(useOS))
	e.GET("/", echo.WrapHandler(assetHandler))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
	e.Logger.Fatal(e.Start(":1323"))
}
```



## With go.rice

cookbook/embed-resources/server.go

```go
package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// the file server for rice. "app" is the folder where the files come from.
	assetHandler := http.FileServer(rice.MustFindBox("app").HTTPBox())
	// serves the index.html from rice
	e.GET("/", echo.WrapHandler(assetHandler))

	// servers other static files
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	e.Logger.Fatal(e.Start(":1323"))
}
```