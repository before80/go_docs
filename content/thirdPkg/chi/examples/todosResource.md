+++
title = "todos resource"
date = 2024-01-31T19:12:40+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-chi/chi/blob/master/_examples/todos-resource/main.go](https://github.com/go-chi/chi/blob/master/_examples/todos-resource/main.go)

## main.go

```go
// This example demonstrates a project structure that defines a subrouter and its
// handlers on a struct, and mounting them as subrouters to a parent router.
// See also _examples/rest for an in-depth example of a REST service, and apply
// those same patterns to this structure.
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Mount("/users", usersResource{}.Routes())
	r.Mount("/todos", todosResource{}.Routes())

	http.ListenAndServe(":3333", r)
}
```

## todos.go

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type todosResource struct{}

// Routes creates a REST router for the todos resource
func (rs todosResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List)    // GET /todos - read a list of todos
	r.Post("/", rs.Create) // POST /todos - create a new todo and persist it
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(rs.TodoCtx) // lets have a todos map, and lets actually load/manipulate
		r.Get("/", rs.Get)       // GET /todos/{id} - read a single todo by :id
		r.Put("/", rs.Update)    // PUT /todos/{id} - update a single todo by :id
		r.Delete("/", rs.Delete) // DELETE /todos/{id} - delete a single todo by :id
		r.Get("/sync", rs.Sync)
	})

	return r
}

func (rs todosResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos list of stuff.."))
}

func (rs todosResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todos create"))
}

func (rs todosResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo get"))
}

func (rs todosResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo update"))
}

func (rs todosResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo delete"))
}

func (rs todosResource) Sync(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo sync"))
}
```

## user.go

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type usersResource struct{}

// Routes creates a REST router for the todos resource
func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List)    // GET /users - read a list of users
	r.Post("/", rs.Create) // POST /users - create a new user and persist it
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(rs.TodoCtx) // lets have a users map, and lets actually load/manipulate
		r.Get("/", rs.Get)       // GET /users/{id} - read a single user by :id
		r.Put("/", rs.Update)    // PUT /users/{id} - update a single user by :id
		r.Delete("/", rs.Delete) // DELETE /users/{id} - delete a single user by :id
	})

	return r
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users list of stuff.."))
}

func (rs usersResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users create"))
}

func (rs usersResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user get"))
}

func (rs usersResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user update"))
}

func (rs usersResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user delete"))
}
```

