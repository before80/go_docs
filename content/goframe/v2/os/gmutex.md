+++
title = "gmutex"
date = 2024-03-21T17:56:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmutex](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmutex)

Package gmutex inherits and extends sync.Mutex and sync.RWMutex with more futures.

​	软件包 gmutex 继承并扩展了同步。互斥和同步。RWMutex拥有更多期货。

Note that, it is refracted using stdlib mutex of package sync from GoFrame version v2.5.2.

​	请注意，它是使用 GoFrame 版本 v2.5.2 的包同步的 stdlib 互斥锁折射的。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Mutex

```go
type Mutex struct {
	sync.Mutex
}
```

Mutex is a high level Mutex, which implements more rich features for mutex.

​	互斥锁是一个高级互斥锁，它为互斥锁实现了更丰富的功能。

#### (*Mutex) LockFunc

```go
func (m *Mutex) LockFunc(f func())
```

LockFunc locks the mutex for writing with given callback function `f`. If there’s a write/reading lock the mutex, it will block until the lock is released.

​	LockFunc 使用给定的回调函数 `f` 锁定互斥锁以进行写入。如果互斥锁存在写/读锁定，它将阻塞，直到释放锁。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### (*Mutex) TryLockFunc

```go
func (m *Mutex) TryLockFunc(f func()) (result bool)
```

TryLockFunc tries locking the mutex for writing with given callback function `f`. it returns true immediately if success, or if there’s a write/reading lock on the mutex, it returns false immediately.

​	TryLockFunc 尝试锁定互斥锁以使用给定的回调函数 `f` 进行写入。如果成功，它会立即返回 true，或者如果互斥锁上有写入/读取锁，它会立即返回 false。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

### type RWMutex <-2.5.3

```go
type RWMutex struct {
	sync.RWMutex
}
```

RWMutex is a high level RWMutex, which implements more rich features for mutex.

​	RWMutex 是一个高级 RWMutex，它为互斥锁实现了更丰富的功能。

#### func New

```go
func New() *RWMutex
```

New creates and returns a new mutex. Deprecated: use Mutex or RWMutex instead.

​	new 创建并返回新的互斥锁。已弃用：请改用 Mutex 或 RWMutex。

#### (*RWMutex) LockFunc

```go
func (m *RWMutex) LockFunc(f func())
```

LockFunc locks the mutex for writing with given callback function `f`. If there’s a write/reading lock the mutex, it will block until the lock is released.

​	LockFunc 锁定互斥锁，以便使用给定的回调函数 `f` 进行写入。如果互斥锁存在写/读锁定，它将阻塞，直到释放锁。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### (*RWMutex) RLockFunc

```go
func (m *RWMutex) RLockFunc(f func())
```

RLockFunc locks the mutex for reading with given callback function `f`. If there’s a writing lock the mutex, it will block until the lock is released.

​	RLockFunc 使用给定的回调函数 `f` 锁定互斥锁以进行读取。如果互斥锁有写入锁，它将阻塞，直到锁被释放。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### (*RWMutex) TryLockFunc

```go
func (m *RWMutex) TryLockFunc(f func()) (result bool)
```

TryLockFunc tries locking the mutex for writing with given callback function `f`. it returns true immediately if success, or if there’s a write/reading lock on the mutex, it returns false immediately.

​	TryLockFunc 尝试锁定互斥锁以使用给定的回调函数 `f` 进行写入。如果成功，它会立即返回 true，或者如果互斥锁上有写入/读取锁，它会立即返回 false。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### (*RWMutex) TryRLockFunc

```go
func (m *RWMutex) TryRLockFunc(f func()) (result bool)
```

TryRLockFunc tries locking the mutex for reading with given callback function `f`. It returns true immediately if success, or if there’s a writing lock on the mutex, it returns false immediately.

​	TryRLockFunc 尝试锁定互斥锁以使用给定的回调函数 `f` 进行读取。如果成功，它会立即返回 true，或者如果互斥锁上有写入锁，则会立即返回 false。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。