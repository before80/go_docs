+++
title = "custom method"
date = 2024-01-31T19:11:07+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-chi/chi/blob/master/_examples/custom-method/main.go](https://github.com/go-chi/chi/blob/master/_examples/custom-method/main.go)

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	chi.RegisterMethod("LINK")
	chi.RegisterMethod("UNLINK")
	chi.RegisterMethod("WOOHOO")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.MethodFunc("LINK", "/link", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("custom link method"))
	})
	r.MethodFunc("WOOHOO", "/woo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("custom woohoo method"))
	})
	r.HandleFunc("/everything", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("capturing all standard http methods, as well as LINK, UNLINK and WOOHOO"))
	})
	http.ListenAndServe(":3333", r)
}
```

