+++
title = "custom handler"
date = 2024-01-31T19:10:50+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-chi/chi/blob/master/_examples/custom-handler/main.go](https://github.com/go-chi/chi/blob/master/_examples/custom-handler/main.go)

```go
package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func main() {
	r := chi.NewRouter()
	r.Method("GET", "/", Handler(customHandler))
	http.ListenAndServe(":3333", r)
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	w.Write([]byte("foo"))
	return nil
}
```

