+++
title = "gmutex"
date = 2024-03-21T17:56:28+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmutex

Package gmutex inherits and extends sync.Mutex and sync.RWMutex with more futures.

Note that, it is refracted using stdlib mutex of package sync from GoFrame version v2.5.2.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type Mutex 

``` go
type Mutex struct {
	sync.Mutex
}
```

Mutex is a high level Mutex, which implements more rich features for mutex.

##### (*Mutex) LockFunc 

``` go
func (m *Mutex) LockFunc(f func())
```

LockFunc locks the mutex for writing with given callback function `f`. If there's a write/reading lock the mutex, it will block until the lock is released.

It releases the lock after `f` is executed.

##### (*Mutex) TryLockFunc 

``` go
func (m *Mutex) TryLockFunc(f func()) (result bool)
```

TryLockFunc tries locking the mutex for writing with given callback function `f`. it returns true immediately if success, or if there's a write/reading lock on the mutex, it returns false immediately.

It releases the lock after `f` is executed.

#### type RWMutex <-2.5.3

``` go
type RWMutex struct {
	sync.RWMutex
}
```

RWMutex is a high level RWMutex, which implements more rich features for mutex.

##### func New 

``` go
func New() *RWMutex
```

New creates and returns a new mutex. Deprecated: use Mutex or RWMutex instead.

##### (*RWMutex) LockFunc <-2.5.3

``` go
func (m *RWMutex) LockFunc(f func())
```

LockFunc locks the mutex for writing with given callback function `f`. If there's a write/reading lock the mutex, it will block until the lock is released.

It releases the lock after `f` is executed.

##### (*RWMutex) RLockFunc <-2.5.3

``` go
func (m *RWMutex) RLockFunc(f func())
```

RLockFunc locks the mutex for reading with given callback function `f`. If there's a writing lock the mutex, it will block until the lock is released.

It releases the lock after `f` is executed.

##### (*RWMutex) TryLockFunc <-2.5.3

``` go
func (m *RWMutex) TryLockFunc(f func()) (result bool)
```

TryLockFunc tries locking the mutex for writing with given callback function `f`. it returns true immediately if success, or if there's a write/reading lock on the mutex, it returns false immediately.

It releases the lock after `f` is executed.

##### (*RWMutex) TryRLockFunc <-2.5.3

``` go
func (m *RWMutex) TryRLockFunc(f func()) (result bool)
```

TryRLockFunc tries locking the mutex for reading with given callback function `f`. It returns true immediately if success, or if there's a writing lock on the mutex, it returns false immediately.

It releases the lock after `f` is executed.