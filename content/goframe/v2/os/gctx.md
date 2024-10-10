+++
title = "gctx"
date = 2024-03-21T17:55:11+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gctx](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gctx)

Package gctx wraps context.Context and provides extra context features.

​	软件包 gctx 包装上下文。上下文并提供额外的上下文功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func CtxId

```go
func CtxId(ctx context.Context) string
```

CtxId retrieves and returns the context id from context.

​	CtxId 从上下文中检索并返回上下文 ID。

#### func GetInitCtx <-2.1.3

```go
func GetInitCtx() context.Context
```

GetInitCtx returns the initialization context. Initialization context is used in `main` or `init` functions.

​	GetInitCtx 返回初始化上下文。初始化上下文用于 `main` or `init` 函数。

#### func NeverDone <-2.5.1

```go
func NeverDone(ctx context.Context) context.Context
```

NeverDone wraps and returns a new context object that will be never done, which forbids the context manually done, to make the context can be propagated to asynchronous goroutines.

​	NeverDone 包装并返回一个永远不会完成的新上下文对象，该对象禁止手动完成的上下文，以使上下文可以传播到异步 goroutine。

Note that, it does not affect the closing (canceling) of the parent context, as it is a wrapper for its parent, which only affects the next context handling.

​	请注意，它不会影响父上下文的关闭（取消），因为它是其父上下文的包装器，它只影响下一个上下文处理。

#### func New

```go
func New() context.Context
```

New creates and returns a context which contains context id.

​	New 创建并返回包含上下文 ID 的上下文。

#### func SetInitCtx <-2.1.3

```go
func SetInitCtx(ctx context.Context)
```

SetInitCtx sets custom initialization context. Note that this function cannot be called in multiple goroutines.

​	SetInitCtx 设置自定义初始化上下文。请注意，不能在多个 goroutine 中调用此函数。

#### func WithCtx

```go
func WithCtx(ctx context.Context) context.Context
```

WithCtx creates and returns a context containing context id upon given parent context `ctx`.

​	WithCtx 在给定的父上下文 `ctx` 上创建并返回包含上下文 ID 的上下文。

## 类型

### type Ctx

```go
type Ctx = context.Context // Ctx is short name alias for context.Context.
```

### type StrKey

``` go
type StrKey string // StrKey is a type for warps basic type string as context key.
```