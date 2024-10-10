+++
title = "versions"
date = 2024-01-31T19:12:49+08:00
weight = 11
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：

```
|---data
|	|
|	---article.go
|	|
|	---errors.go
|---presenter
|    |
|    |---v1
|    |    |
|    |    ---article.go
|    |---v2
|    |    |
|    |    ---article.go
|    |---v3
|    |    |
|    |    ---article.go
|---main.go
|---chi.svg
```

## main.go

```go
// This example demonstrates the use of the render subpackage, with
// a quick concept for how to support multiple api versions.
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/_examples/versions/data"
	v1 "github.com/go-chi/chi/v5/_examples/versions/presenter/v1"
	v2 "github.com/go-chi/chi/v5/_examples/versions/presenter/v2"
	v3 "github.com/go-chi/chi/v5/_examples/versions/presenter/v3"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// API version 3.
	r.Route("/v3", func(r chi.Router) {
		r.Use(apiVersionCtx("v3"))
		r.Mount("/articles", articleRouter())
	})

	// API version 2.
	r.Route("/v2", func(r chi.Router) {
		r.Use(apiVersionCtx("v2"))
		r.Mount("/articles", articleRouter())
	})

	// API version 1.
	r.Route("/v1", func(r chi.Router) {
		r.Use(randomErrorMiddleware) // Simulate random error, ie. version 1 is buggy.
		r.Use(apiVersionCtx("v1"))
		r.Mount("/articles", articleRouter())
	})

	http.ListenAndServe(":3333", r)
}

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

func articleRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", listArticles)
	r.Route("/{articleID}", func(r chi.Router) {
		r.Get("/", getArticle)
		// r.Put("/", updateArticle)
		// r.Delete("/", deleteArticle)
	})
	return r
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	articles := make(chan render.Renderer, 5)

	// Load data asynchronously into the channel (simulate slow storage):
	go func() {
		for i := 1; i <= 10; i++ {
			article := &data.Article{
				ID:                     i,
				Title:                  fmt.Sprintf("Article #%v", i),
				Data:                   []string{"one", "two", "three", "four"},
				CustomDataForAuthUsers: "secret data for auth'd users only",
			}

			apiVersion := r.Context().Value("api.version").(string)
			switch apiVersion {
			case "v1":
				articles <- v1.NewArticleResponse(article)
			case "v2":
				articles <- v2.NewArticleResponse(article)
			default:
				articles <- v3.NewArticleResponse(article)
			}

			time.Sleep(100 * time.Millisecond)
		}
		close(articles)
	}()

	// Start streaming data from the channel.
	render.Respond(w, r, articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	// Load article.
	if chi.URLParam(r, "articleID") != "1" {
		render.Respond(w, r, data.ErrNotFound)
		return
	}
	article := &data.Article{
		ID:                     1,
		Title:                  "Article #1",
		Data:                   []string{"one", "two", "three", "four"},
		CustomDataForAuthUsers: "secret data for auth'd users only",
	}

	// Simulate some context values:
	// 1. ?auth=true simulates authenticated session/user.
	// 2. ?error=true simulates random error.
	if r.URL.Query().Get("auth") != "" {
		r = r.WithContext(context.WithValue(r.Context(), "auth", true))
	}
	if r.URL.Query().Get("error") != "" {
		render.Respond(w, r, errors.New("error"))
		return
	}

	var payload render.Renderer

	apiVersion := r.Context().Value("api.version").(string)
	switch apiVersion {
	case "v1":
		payload = v1.NewArticleResponse(article)
	case "v2":
		payload = v2.NewArticleResponse(article)
	default:
		payload = v3.NewArticleResponse(article)
	}

	render.Render(w, r, payload)
}

func randomErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().Unix())

		// One in three chance of random error.
		if rand.Int31n(3) == 0 {
			errors := []error{data.ErrUnauthorized, data.ErrForbidden, data.ErrNotFound}
			render.Respond(w, r, errors[rand.Intn(len(errors))])
			return
		}
		next.ServeHTTP(w, r)
	})
}
```

## data/article.go

```go
package data

// Article is runtime object, that's not meant to be sent via REST.
type Article struct {
	ID                     int      `db:"id" json:"id" xml:"id"`
	Title                  string   `db:"title" json:"title" xml:"title"`
	Data                   []string `db:"data,stringarray" json:"data" xml:"data"`
	CustomDataForAuthUsers string   `db:"custom_data" json:"-" xml:"-"`
}
```

## data/errors.go

```go
package data

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrForbidden    = errors.New("Forbidden")
	ErrNotFound     = errors.New("Resource not found")
)

func PresentError(r *http.Request, err error) (*http.Request, interface{}) {
	switch err {
	case ErrUnauthorized:
		render.Status(r, 401)
	case ErrForbidden:
		render.Status(r, 403)
	case ErrNotFound:
		render.Status(r, 404)
	default:
		render.Status(r, 500)
	}
	return r, map[string]string{"error": err.Error()}
}
```

## presenter/v1/article.go

```go
package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5/_examples/versions/data"
)

// Article presented in API version 1.
type Article struct {
	*data.Article

	Data map[string]bool `json:"data" xml:"data"`
}

func (a *Article) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewArticleResponse(article *data.Article) *Article {
	return &Article{Article: article}
}
```

## presenter/v2/article.go

```go
package v2

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/_examples/versions/data"
)

// Article presented in API version 2.
type Article struct {
	// *v3.Article `json:",inline" xml:",inline"`

	*data.Article

	// Additional fields.
	SelfURL string `json:"self_url" xml:"self_url"`

	// Omitted fields.
	URL interface{} `json:"url,omitempty" xml:"url,omitempty"`
}

func (a *Article) Render(w http.ResponseWriter, r *http.Request) error {
	a.SelfURL = fmt.Sprintf("http://localhost:3333/v2?id=%v", a.ID)
	return nil
}

func NewArticleResponse(article *data.Article) *Article {
	return &Article{Article: article}
}
```

## presenter/v3/article.go

```go
package v3

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5/_examples/versions/data"
)

// Article presented in API version 2.
type Article struct {
	*data.Article `json:",inline" xml:",inline"`

	// Additional fields.
	URL        string `json:"url" xml:"url"`
	ViewsCount int64  `json:"views_count" xml:"views_count"`
	APIVersion string `json:"api_version" xml:"api_version"`

	// Omitted fields.
	// Show custom_data explicitly for auth'd users only.
	CustomDataForAuthUsers interface{} `json:"custom_data,omitempty" xml:"custom_data,omitempty"`
}

func (a *Article) Render(w http.ResponseWriter, r *http.Request) error {
	a.ViewsCount = rand.Int63n(100000)
	a.URL = fmt.Sprintf("http://localhost:3333/v3/?id=%v", a.ID)

	// Only show to auth'd user.
	if _, ok := r.Context().Value("auth").(bool); ok {
		a.CustomDataForAuthUsers = a.Article.CustomDataForAuthUsers
	}

	return nil
}

func NewArticleResponse(article *data.Article) *Article {
	return &Article{Article: article}
}
```

## chi.svg

![img](./versions_img/chi.svg)