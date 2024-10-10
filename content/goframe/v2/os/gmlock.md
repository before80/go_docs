+++
title = "gmlock"
date = 2024-03-21T17:56:11+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gmlock)

Package gmlock implements a concurrent-safe memory-based locker.

​	软件包 gmlock 实现了一个基于并发安全内存的储物柜。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Lock

```go
func Lock(key string)
```

Lock locks the `key` with writing lock. If there’s a write/reading lock the `key`, it will blocks until the lock is released.

​	锁 `key` 用写入锁锁定。如果存在写/读锁， `key` 它将阻塞，直到锁被释放。

#### func LockFunc

```go
func LockFunc(key string, f func())
```

LockFunc locks the `key` with writing lock and callback function `f`. If there’s a write/reading lock the `key`, it will blocks until the lock is released.

​	LockFunc 用写入锁和回调函数 `f` 锁定。 `key` 如果存在写/读锁， `key` 它将阻塞，直到锁被释放。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### func RLock

```go
func RLock(key string)
```

RLock locks the `key` with reading lock. If there’s a writing lock on `key`, it will blocks until the writing lock is released.

​	RLock 使用读取锁锁定。 `key` 如果有一个写入锁， `key` 它将阻塞，直到写入锁被释放。

#### func RLockFunc

```go
func RLockFunc(key string, f func())
```

RLockFunc locks the `key` with reading lock and callback function `f`. If there’s a writing lock the `key`, it will blocks until the lock is released.

​	RLockFunc `key` 使用读取锁定和回调函数 `f` 锁定。如果有写入锁， `key` 它将阻塞，直到锁被释放。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### func RUnlock

```go
func RUnlock(key string)
```

RUnlock unlocks the reading lock of the `key`.

​	RUnlock 解锁 `key` .

#### func Remove

```go
func Remove(key string)
```

Remove removes mutex with given `key`.

​	Remove 删除具有给定 `key` .

#### func TryLock

```go
func TryLock(key string) bool
```

TryLock tries locking the `key` with writing lock, it returns true if success, or if there’s a write/reading lock the `key`, it returns false.

​	TryLock 尝试用写入锁锁定， `key` 如果成功，则返回 true，或者如果存在写/读锁定 `key` ，则返回 false。

#### func TryLockFunc

```go
func TryLockFunc(key string, f func()) bool
```

TryLockFunc locks the `key` with writing lock and callback function `f`. It returns true if success, or else if there’s a write/reading lock the `key`, it return false.

​	TryLockFunc `key` 用写入锁和回调函数 `f` 锁定 。如果成功，则返回 true，否则如果存在写入/读取锁定 `key` ，则返回 false。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### func TryRLock

```go
func TryRLock(key string) bool
```

TryRLock tries locking the `key` with reading lock. It returns true if success, or if there’s a writing lock on `key`, it returns false.

​	TryRLock 尝试 `key` 锁定带读取锁。如果成功，则返回 true，或者如果存在写入锁定， `key` 则返回 false。

#### func TryRLockFunc

```go
func TryRLockFunc(key string, f func()) bool
```

TryRLockFunc locks the `key` with reading lock and callback function `f`. It returns true if success, or else if there’s a writing lock the `key`, it returns false.

​	TryRLockFunc `key` 使用读取锁和回调函数 `f` 锁定 。如果成功，则返回 true，否则如果存在写入锁定 `key` ，则返回 false。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### func Unlock

```go
func Unlock(key string)
```

Unlock unlocks the writing lock of the `key`.

​	Unlock 解锁 `key` 的写入锁。

## 类型

### type Locker

```go
type Locker struct {
	// contains filtered or unexported fields
}
```

Locker is a memory based locker. Note that there’s no cache expire mechanism for mutex in locker. You need remove certain mutex manually when you do not want use it anymore.

​	储物柜是一个基于内存的储物柜。请注意，储物柜中没有互斥锁的缓存过期机制。当您不想再使用它时，您需要手动删除某些互斥锁。

#### func New

```go
func New() *Locker
```

New creates and returns a new memory locker. A memory locker can lock/unlock with dynamic string key.

​	new 创建并返回新的内存存储箱。内存储物柜可以使用动态字符串键锁定/解锁。

#### (*Locker) Clear

```go
func (l *Locker) Clear()
```

Clear removes all mutexes from locker.

​	清除会从储物柜中删除所有互斥锁。

#### (*Locker) Lock

```go
func (l *Locker) Lock(key string)
```

Lock locks the `key` with writing lock. If there’s a write/reading lock the `key`, it will block until the lock is released.

​	锁 `key` 用写入锁锁定。如果存在写/读锁， `key` 它将阻塞，直到锁被释放。

#### (*Locker) LockFunc

```go
func (l *Locker) LockFunc(key string, f func())
```

LockFunc locks the `key` with writing lock and callback function `f`. If there’s a write/reading lock the `key`, it will block until the lock is released.

​	LockFunc 用写入锁和回调函数 `f` 锁定。 `key` 如果存在写/读锁， `key` 它将阻塞，直到锁被释放。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### (*Locker) RLock

```go
func (l *Locker) RLock(key string)
```

RLock locks the `key` with reading lock. If there’s a writing lock on `key`, it will blocks until the writing lock is released.

​	RLock 使用读取锁锁定。 `key` 如果有一个写入锁打开 `key` ，它将阻塞，直到写入锁被释放。

#### (*Locker) RLockFunc

```go
func (l *Locker) RLockFunc(key string, f func())
```

RLockFunc locks the `key` with reading lock and callback function `f`. If there’s a writing lock the `key`, it will block until the lock is released.

​	RLockFunc `key` 使用读取锁定和回调函数 `f` 锁定 。如果有写入锁， `key` 它将阻塞，直到锁被释放。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### (*Locker) RUnlock

```go
func (l *Locker) RUnlock(key string)
```

RUnlock unlocks the reading lock of the `key`.

​	RUnlock 解锁 `key` .

#### (*Locker) Remove

```go
func (l *Locker) Remove(key string)
```

Remove removes mutex with given `key` from locker.

​	Remove 从储物柜中删除 given `key` 的互斥锁。

#### (*Locker) TryLock

```go
func (l *Locker) TryLock(key string) bool
```

TryLock tries locking the `key` with writing lock, it returns true if success, or it returns false if there’s a writing/reading lock the `key`.

​	TryLock 尝试 `key` 锁定 with 写入锁，如果成功，则返回 true，如果 有写入/读取锁， `key` 则返回 false。

#### (*Locker) TryLockFunc

```go
func (l *Locker) TryLockFunc(key string, f func()) bool
```

TryLockFunc locks the `key` with writing lock and callback function `f`. It returns true if success, or else if there’s a write/reading lock the `key`, it return false.

​	TryLockFunc `key` 用写入锁和回调函数 `f` 锁定 。如果成功，则返回 true，否则如果存在写入/读取锁定 `key` ，则返回 false。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### (*Locker) TryRLock

```go
func (l *Locker) TryRLock(key string) bool
```

TryRLock tries locking the `key` with reading lock. It returns true if success, or if there’s a writing lock on `key`, it returns false.

​	TryRLock 尝试 `key` 锁定带读取锁。如果成功，则返回 true，或者如果存在写入锁定， `key` 则返回 false。

#### (*Locker) TryRLockFunc

```go
func (l *Locker) TryRLockFunc(key string, f func()) bool
```

TryRLockFunc locks the `key` with reading lock and callback function `f`. It returns true if success, or else if there’s a writing lock the `key`, it returns false.

​	TryRLockFunc `key` 使用读取锁和回调函数 `f` 锁定 。如果成功，则返回 true，否则如果存在写入锁定 `key` ，则返回 false。

It releases the lock after `f` is executed.

​	它在执行后 `f` 释放锁。

#### (*Locker) Unlock

```go
func (l *Locker) Unlock(key string)
```

Unlock unlocks the writing lock of the `key`.

​	Unlock 解锁 `key` 的写入锁。