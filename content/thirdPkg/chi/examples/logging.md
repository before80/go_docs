+++
title = "logging"
date = 2024-01-31T19:12:06+08:00
weight =7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-chi/chi/blob/master/_examples/logging/main.go](https://github.com/go-chi/chi/blob/master/_examples/logging/main.go)

```go
package main

// Please see https://github.com/go-chi/httplog for a complete package
// and example for writing a structured logger on chi built on
// the Go 1.21+ "log/slog" package.

func main() {
	// See https://github.com/go-chi/httplog/blob/master/_example/main.go
}
```

> 原文：[https://github.com/go-chi/httplog/blob/master/_example/main.go](https://github.com/go-chi/httplog/blob/master/_example/main.go):

```go
package main

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
)

func main() {
	// Logger
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		LogLevel: slog.LevelDebug,
		// JSON:             true,
		Concise: true,
		// RequestHeaders:   true,
		// ResponseHeaders:  true,
		MessageFieldName: "message",
		LevelFieldName:   "severity",
		TimeFieldFormat:  time.RFC3339,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
		QuietDownRoutes: []string{
			"/",
			"/ping",
		},
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})

	// Service
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger, []string{"/ping"}))
	r.Use(middleware.Heartbeat("/ping"))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			httplog.LogEntrySetField(ctx, "user", slog.StringValue("user1"))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oh no")
	})

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		w.Header().Add("Content-Type", "text/plain")
		oplog.Info("info here")
		w.Write([]byte("info here"))
	})

	r.Get("/warn", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		oplog.Warn("warn here")
		w.WriteHeader(400)
		w.Write([]byte("warn here"))
	})

	r.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())

		// two varianets of syntax to specify "err" attr.
		err := errors.New("err here")
		// oplog.Error("msg here", "err", err)
		oplog.Error("msg here", httplog.ErrAttr(err))

		// logging with the global logger also works
		slog.Default().With(slog.Group("ImpGroup", slog.String("account", "id"))).Error("doesn't exist")
		slog.Default().Error("oops, err occured")
		w.WriteHeader(500)
		w.Write([]byte("oops, err"))
	})

	http.ListenAndServe("localhost:8000", r)
}
```

