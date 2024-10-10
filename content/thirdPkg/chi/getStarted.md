+++
title = "开始入门"
date = 2024-01-31T19:04:19+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-chi.io/#/pages/getting_started](https://go-chi.io/#/pages/getting_started)

## Installation 安装

```
go get -u github.com/go-chi/chi/v5
```

## Running a Simple Server 运行一个简单的服务器

The simplest Hello World Api Can look like this.

​	最简单的 Hello World Api 可以像这样。

```go
package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
    http.ListenAndServe(":3000", r)
}Copy to clipboardErrorCopied
go run main.goCopy to clipboardErrorCopied
```

Browse to `http://localhost:3000`, and you should see `Hello World!` on the page.

​	浏览到 `http://localhost:3000` ，您应该在页面上看到 `Hello World!` 。

