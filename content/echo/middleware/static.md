+++
title = "static"
weight = 230
date = 2023-07-09T21:58:17+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Static

https://echo.labstack.com/docs/middleware/static

Static middleware can be used to serve static files from the provided root directory.

## Usage

```go
e := echo.New()
e.Use(middleware.Static("/static"))
```



This serves static files from `static` directory. For example, a request to `/js/main.js` will fetch and serve `static/js/main.js` file.

## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
  Root:   "static",
  Browse: true,
}))
```



This serves static files from `static` directory and enables directory browsing.

Default behavior when using with non root URL paths is to append the URL path to the filesystem path.

#### Example

```go
group := root.Group("somepath")
group.Use(middleware.Static(filepath.Join("filesystempath")))
// When an incoming request comes for `/somepath` the actual filesystem request goes to `filesystempath/somepath` instead of only `filesystempath`. 
```



TIP

To turn off this behavior set the `IgnoreBase` config param to `true`.

## Configuration

```go
StaticConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Root directory from where the static content is served.
  // Required.
  Root string `json:"root"`

  // Index file for serving a directory.
  // Optional. Default value "index.html".
  Index string `json:"index"`

  // Enable HTML5 mode by forwarding all not-found requests to root so that
  // SPA (single-page application) can handle the routing.
  // Optional. Default value false.
  HTML5 bool `json:"html5"`

  // Enable directory browsing.
  // Optional. Default value false.
  Browse bool `json:"browse"`
  
  // Enable ignoring of the base of the URL path.
  // Example: when assigning a static middleware to a non root path group,
  // the filesystem path is not doubled
  // Optional. Default value false.
  IgnoreBase bool `yaml:"ignoreBase"`
}
```



### Default Configuration

```go
DefaultStaticConfig = StaticConfig{
  Skipper: DefaultSkipper,
  Index:   "index.html",
}
```