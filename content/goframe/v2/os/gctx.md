+++
title = "gctx"
date = 2024-03-21T17:55:11+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gctx

Package gctx wraps context.Context and provides extra context features.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func CtxId 

``` go
func CtxId(ctx context.Context) string
```

CtxId retrieves and returns the context id from context.

##### func GetInitCtx <-2.1.3

``` go
func GetInitCtx() context.Context
```

GetInitCtx returns the initialization context. Initialization context is used in `main` or `init` functions.

##### func NeverDone <-2.5.1

``` go
func NeverDone(ctx context.Context) context.Context
```

NeverDone wraps and returns a new context object that will be never done, which forbids the context manually done, to make the context can be propagated to asynchronous goroutines.

Note that, it does not affect the closing (canceling) of the parent context, as it is a wrapper for its parent, which only affects the next context handling.

##### func New 

``` go
func New() context.Context
```

New creates and returns a context which contains context id.

##### func SetInitCtx <-2.1.3

``` go
func SetInitCtx(ctx context.Context)
```

SetInitCtx sets custom initialization context. Note that this function cannot be called in multiple goroutines.

##### func WithCtx 

``` go
func WithCtx(ctx context.Context) context.Context
```

WithCtx creates and returns a context containing context id upon given parent context `ctx`.

### Types 

#### type Ctx 

``` go
type Ctx = context.Context // Ctx is short name alias for context.Context.
```

#### type StrKey 

``` go
type StrKey string // StrKey is a type for warps basic type string as context key.
```