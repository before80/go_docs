+++
title = "grpool"
date = 2024-03-21T17:56:58+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/grpool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/grpool)

Package grpool implements a goroutine reusable pool.

​	软件包 grpool 实现了一个 goroutine 可重用池。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Add

```go
func Add(ctx context.Context, f Func) error
```

Add pushes a new job to the default goroutine pool. The job will be executed asynchronously.

​	Add 将新作业推送到默认的 goroutine 池。作业将异步执行。

#### func AddWithRecover

```go
func AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error
```

AddWithRecover pushes a new job to the default pool with specified recover function.

​	AddWithRecover 使用指定的恢复函数将新作业推送到默认池。

The optional `recoverFunc` is called when any panic during executing of `userFunc`. If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`. The job will be executed asynchronously.

​	当 执行 期间出现任何 panic 时，将调用可选 `userFunc` 项 `recoverFunc` 。如果 `recoverFunc` 未传递或给出 nil，则忽略来自 `userFunc` 的恐慌。作业将异步执行。

#### func Jobs

```go
func Jobs() int
```

Jobs returns current job count of default goroutine pool.

​	Jobs 返回默认 goroutine 池的当前作业计数。

#### func Size

```go
func Size() int
```

Size returns current goroutine count of default goroutine pool.

​	Size 返回默认 goroutine 池的当前 goroutine 计数。

## 类型

### type Func

```go
type Func func(ctx context.Context)
```

Func is the pool function which contains context parameter.

​	Func 是包含 context 参数的池函数。

### type Pool

```go
type Pool struct {
	// contains filtered or unexported fields
}
```

Pool manages the goroutines using pool.

​	Pool 使用 pool 管理 goroutine。

#### func New

```go
func New(limit ...int) *Pool
```

New creates and returns a new goroutine pool object. The parameter `limit` is used to limit the max goroutine count, which is not limited in default.

​	New 创建并返回一个新的 goroutine 池对象。该参数 `limit` 用于限制最大 goroutine 计数，默认情况下不受限制。

#### (*Pool) Add

```go
func (p *Pool) Add(ctx context.Context, f Func) error
```

Add pushes a new job to the pool. The job will be executed asynchronously.

​	添加将新作业推送到池中。作业将异步执行。

#### (*Pool) AddWithRecover

```go
func (p *Pool) AddWithRecover(ctx context.Context, userFunc Func, recoverFunc RecoverFunc) error
```

AddWithRecover pushes a new job to the pool with specified recover function.

​	AddWithRecover 使用指定的 recover 函数将新作业推送到池中。

The optional `recoverFunc` is called when any panic during executing of `userFunc`. If `recoverFunc` is not passed or given nil, it ignores the panic from `userFunc`. The job will be executed asynchronously.

​	当 执行 期间出现任何 panic 时，将调用可选 `userFunc` 项 `recoverFunc` 。如果 `recoverFunc` 未传递或给出 nil，则忽略来自 `userFunc` 的恐慌。作业将异步执行。

#### (*Pool) Cap

```go
func (p *Pool) Cap() int
```

Cap returns the capacity of the pool. This capacity is defined when pool is created. It returns -1 if there’s no limit.

​	Cap 返回池的容量。此容量是在创建池时定义的。如果没有限制，则返回 -1。

#### (*Pool) Close

```go
func (p *Pool) Close()
```

Close closes the goroutine pool, which makes all goroutines exit.

​	Close 关闭 goroutine 池，这将使所有 goroutine 退出。

#### (*Pool) IsClosed

```go
func (p *Pool) IsClosed() bool
```

IsClosed returns if pool is closed.

​	如果池已关闭，则返回 IsClosed。

#### (*Pool) Jobs

```go
func (p *Pool) Jobs() int
```

Jobs returns current job count of the pool. Note that, it does not return worker/goroutine count but the job/task count.

​	作业返回池的当前作业计数。请注意，它不会返回 worker/goroutine 计数，而是返回作业/任务计数。

#### (*Pool) Size

```go
func (p *Pool) Size() int
```

Size returns current goroutine count of the pool.

​	size 返回池的当前 goroutine 计数。

### type RecoverFunc <-2.1.0

```go
type RecoverFunc func(ctx context.Context, exception error)
```

RecoverFunc is the pool runtime panic recover function which contains context parameter.

​	RecoverFunc 是包含上下文参数的池运行时紧急恢复函数。