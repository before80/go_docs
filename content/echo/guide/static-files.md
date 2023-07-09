+++
title = "static-files"
date = 2023-07-09T21:51:46+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Static Files

https://echo.labstack.com/docs/static-files

Images, JavaScript, CSS, PDF, Fonts and so on...

## Using Static Middleware

[See](https://echo.labstack.com/docs/middleware/static)

## Using Echo#Static()

`Echo#Static(prefix, root string)` registers a new route with path prefix to serve static files from the provided root directory.

*Usage 1*

```go
e := echo.New()
e.Static("/static", "assets")
```



Example above will serve any file from the assets directory for path `/static/*`. For example, a request to `/static/js/main.js` will fetch and serve `assets/js/main.js` file.

*Usage 2*

```go
e := echo.New()
e.Static("/", "assets")
```



Example above will serve any file from the assets directory for path `/*`. For example, a request to `/js/main.js` will fetch and serve `assets/js/main.js` file.

## Using Echo#File()

`Echo#File(path, file string)` registers a new route with path to serve a static file.

*Usage 1*

Serving an index page from `public/index.html`

```go
e.File("/", "public/index.html")
```



*Usage 2*

Serving a favicon from `images/favicon.ico`

```go
e.File("/favicon.ico", "images/favicon.ico")
```