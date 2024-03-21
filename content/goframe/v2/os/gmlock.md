+++
title = "gmlock"
date = 2024-03-21T17:56:11+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock

Package gmlock implements a concurrent-safe memory-based locker.

### Index 

- [func Lock(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Lock)
- [func LockFunc(key string, f func())](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#LockFunc)
- [func RLock(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#RLock)
- [func RLockFunc(key string, f func())](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#RLockFunc)
- [func RUnlock(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#RUnlock)
- [func Remove(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Remove)
- [func TryLock(key string) bool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#TryLock)
- [func TryLockFunc(key string, f func()) bool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#TryLockFunc)
- [func TryRLock(key string) bool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#TryRLock)
- [func TryRLockFunc(key string, f func()) bool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#TryRLockFunc)
- [func Unlock(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Unlock)
- [type Locker](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker)
- - [func New() *Locker](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#New)
- - [func (l *Locker) Clear()](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.Clear)
  - [func (l *Locker) Lock(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.Lock)
  - [func (l *Locker) LockFunc(key string, f func())](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.LockFunc)
  - [func (l *Locker) RLock(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.RLock)
  - [func (l *Locker) RLockFunc(key string, f func())](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.RLockFunc)
  - [func (l *Locker) RUnlock(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.RUnlock)
  - [func (l *Locker) Remove(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.Remove)
  - [func (l *Locker) TryLock(key string) bool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.TryLock)
  - [func (l *Locker) TryLockFunc(key string, f func()) bool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.TryLockFunc)
  - [func (l *Locker) TryRLock(key string) bool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.TryRLock)
  - [func (l *Locker) TryRLockFunc(key string, f func()) bool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.TryRLockFunc)
  - [func (l *Locker) Unlock(key string)](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock#Locker.Unlock)

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Lock 

``` go
func Lock(key string)
```

Lock locks the `key` with writing lock. If there's a write/reading lock the `key`, it will blocks until the lock is released.

##### func LockFunc 

``` go
func LockFunc(key string, f func())
```

LockFunc locks the `key` with writing lock and callback function `f`. If there's a write/reading lock the `key`, it will blocks until the lock is released.

It releases the lock after `f` is executed.

##### func RLock 

``` go
func RLock(key string)
```

RLock locks the `key` with reading lock. If there's a writing lock on `key`, it will blocks until the writing lock is released.

##### func RLockFunc 

``` go
func RLockFunc(key string, f func())
```

RLockFunc locks the `key` with reading lock and callback function `f`. If there's a writing lock the `key`, it will blocks until the lock is released.

It releases the lock after `f` is executed.

##### func RUnlock 

``` go
func RUnlock(key string)
```

RUnlock unlocks the reading lock of the `key`.

##### func Remove 

``` go
func Remove(key string)
```

Remove removes mutex with given `key`.

##### func TryLock 

``` go
func TryLock(key string) bool
```

TryLock tries locking the `key` with writing lock, it returns true if success, or if there's a write/reading lock the `key`, it returns false.

##### func TryLockFunc 

``` go
func TryLockFunc(key string, f func()) bool
```

TryLockFunc locks the `key` with writing lock and callback function `f`. It returns true if success, or else if there's a write/reading lock the `key`, it return false.

It releases the lock after `f` is executed.

##### func TryRLock 

``` go
func TryRLock(key string) bool
```

TryRLock tries locking the `key` with reading lock. It returns true if success, or if there's a writing lock on `key`, it returns false.

##### func TryRLockFunc 

``` go
func TryRLockFunc(key string, f func()) bool
```

TryRLockFunc locks the `key` with reading lock and callback function `f`. It returns true if success, or else if there's a writing lock the `key`, it returns false.

It releases the lock after `f` is executed.

##### func Unlock 

``` go
func Unlock(key string)
```

Unlock unlocks the writing lock of the `key`.

### Types 

#### type Locker 

``` go
type Locker struct {
	// contains filtered or unexported fields
}
```

Locker is a memory based locker. Note that there's no cache expire mechanism for mutex in locker. You need remove certain mutex manually when you do not want use it anymore.

##### func New 

``` go
func New() *Locker
```

New creates and returns a new memory locker. A memory locker can lock/unlock with dynamic string key.

##### (*Locker) Clear 

``` go
func (l *Locker) Clear()
```

Clear removes all mutexes from locker.

##### (*Locker) Lock 

``` go
func (l *Locker) Lock(key string)
```

Lock locks the `key` with writing lock. If there's a write/reading lock the `key`, it will block until the lock is released.

##### (*Locker) LockFunc 

``` go
func (l *Locker) LockFunc(key string, f func())
```

LockFunc locks the `key` with writing lock and callback function `f`. If there's a write/reading lock the `key`, it will block until the lock is released.

It releases the lock after `f` is executed.

##### (*Locker) RLock 

``` go
func (l *Locker) RLock(key string)
```

RLock locks the `key` with reading lock. If there's a writing lock on `key`, it will blocks until the writing lock is released.

##### (*Locker) RLockFunc 

``` go
func (l *Locker) RLockFunc(key string, f func())
```

RLockFunc locks the `key` with reading lock and callback function `f`. If there's a writing lock the `key`, it will block until the lock is released.

It releases the lock after `f` is executed.

##### (*Locker) RUnlock 

``` go
func (l *Locker) RUnlock(key string)
```

RUnlock unlocks the reading lock of the `key`.

##### (*Locker) Remove 

``` go
func (l *Locker) Remove(key string)
```

Remove removes mutex with given `key` from locker.

##### (*Locker) TryLock 

``` go
func (l *Locker) TryLock(key string) bool
```

TryLock tries locking the `key` with writing lock, it returns true if success, or it returns false if there's a writing/reading lock the `key`.

##### (*Locker) TryLockFunc 

``` go
func (l *Locker) TryLockFunc(key string, f func()) bool
```

TryLockFunc locks the `key` with writing lock and callback function `f`. It returns true if success, or else if there's a write/reading lock the `key`, it return false.

It releases the lock after `f` is executed.

##### (*Locker) TryRLock 

``` go
func (l *Locker) TryRLock(key string) bool
```

TryRLock tries locking the `key` with reading lock. It returns true if success, or if there's a writing lock on `key`, it returns false.

##### (*Locker) TryRLockFunc 

``` go
func (l *Locker) TryRLockFunc(key string, f func()) bool
```

TryRLockFunc locks the `key` with reading lock and callback function `f`. It returns true if success, or else if there's a writing lock the `key`, it returns false.

It releases the lock after `f` is executed.

##### (*Locker) Unlock 

``` go
func (l *Locker) Unlock(key string)
```

Unlock unlocks the writing lock of the `key`.