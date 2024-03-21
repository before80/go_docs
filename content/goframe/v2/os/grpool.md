+++
title = "grpool"
date = 2024-03-21T17:56:58+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/grpool

Package grpool implements a goroutine reusable pool.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Add 

``` go
func Add(ctx context.Context, f Func) error
```

Add pushes a new job to the default goroutine pool. The job will be executed asynchronously.

##### func AddWithRecover 

``` go
func AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error
```

AddWithRecover pushes a new job to the default pool with specified recover function.

The optional `recoverFunc` is called when any panic during executing of `userFunc`. If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`. The job will be executed asynchronously.

##### func Jobs 

``` go
func Jobs() int
```

Jobs returns current job count of default goroutine pool.

##### func Size 

``` go
func Size() int
```

Size returns current goroutine count of default goroutine pool.

### Types 

#### type Func 

``` go
type Func func(ctx context.Context)
```

Func is the pool function which contains context parameter.

#### type Pool 

``` go
type Pool struct {
	// contains filtered or unexported fields
}
```

Pool manages the goroutines using pool.

##### func New 

``` go
func New(limit ...int) *Pool
```

New creates and returns a new goroutine pool object. The parameter `limit` is used to limit the max goroutine count, which is not limited in default.

##### (*Pool) Add 

``` go
func (p *Pool) Add(ctx context.Context, f Func) error
```

Add pushes a new job to the pool. The job will be executed asynchronously.

##### (*Pool) AddWithRecover 

``` go
func (p *Pool) AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error
```

AddWithRecover pushes a new job to the pool with specified recover function.

The optional `recoverFunc` is called when any panic during executing of `userFunc`. If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`. The job will be executed asynchronously.

##### (*Pool) Cap 

``` go
func (p *Pool) Cap() int
```

Cap returns the capacity of the pool. This capacity is defined when pool is created. It returns -1 if there's no limit.

##### (*Pool) Close 

``` go
func (p *Pool) Close()
```

Close closes the goroutine pool, which makes all goroutines exit.

##### (*Pool) IsClosed 

``` go
func (p *Pool) IsClosed() bool
```

IsClosed returns if pool is closed.

##### (*Pool) Jobs 

``` go
func (p *Pool) Jobs() int
```

Jobs returns current job count of the pool. Note that, it does not return worker/goroutine count but the job/task count.

##### (*Pool) Size 

``` go
func (p *Pool) Size() int
```

Size returns current goroutine count of the pool.

#### type RecoverFunc <-2.1.0

``` go
type RecoverFunc func(ctx context.Context, exception error)
```

RecoverFunc is the pool runtime panic recover function which contains context parameter.