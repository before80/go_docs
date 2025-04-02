+++
title = "weak"
date = 2025-04-02T16:15:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/weak@go1.24.2](https://pkg.go.dev/weak@go1.24.2)
>

> 注意
>
> ​	从go1.24.0开始才可以使用该包。



## Overview 

Package weak provides ways to safely reference memory weakly, that is, without preventing its reclamation.

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Pointer 

```go
type Pointer[T any] struct {
	// contains filtered or unexported fields
}
```

Pointer is a weak pointer to a value of type T.

Just like regular pointers, Pointer may reference any part of an object, such as a field of a struct or an element of an array. Objects that are only pointed to by weak pointers are not considered reachable, and once the object becomes unreachable, [Pointer.Value](https://pkg.go.dev/weak@go1.24.2#Pointer.Value) may return nil.

The primary use-cases for weak pointers are for implementing caches, canonicalization maps (like the unique package), and for tying together the lifetimes of separate values (for example, through a map with weak keys).

Two Pointer values always compare equal if the pointers from which they were created compare equal. This property is retained even after the object referenced by the pointer used to create a weak reference is reclaimed. If multiple weak pointers are made to different offsets within the same object (for example, pointers to different fields of the same struct), those pointers will not compare equal. If a weak pointer is created from an object that becomes unreachable, but is then resurrected due to a finalizer, that weak pointer will not compare equal with weak pointers created after the resurrection.

Calling [Make](https://pkg.go.dev/weak@go1.24.2#Make) with a nil pointer returns a weak pointer whose [Pointer.Value](https://pkg.go.dev/weak@go1.24.2#Pointer.Value) always returns nil. The zero value of a Pointer behaves as if it were created by passing nil to [Make](https://pkg.go.dev/weak@go1.24.2#Make) and compares equal with such pointers.

[Pointer.Value](https://pkg.go.dev/weak@go1.24.2#Pointer.Value) is not guaranteed to eventually return nil. [Pointer.Value](https://pkg.go.dev/weak@go1.24.2#Pointer.Value) may return nil as soon as the object becomes unreachable. Values stored in global variables, or that can be found by tracing pointers from a global variable, are reachable. A function argument or receiver may become unreachable at the last point where the function mentions it. To ensure [Pointer.Value](https://pkg.go.dev/weak@go1.24.2#Pointer.Value) does not return nil, pass a pointer to the object to the [runtime.KeepAlive](https://pkg.go.dev/runtime#KeepAlive) function after the last point where the object must remain reachable.

Note that because [Pointer.Value](https://pkg.go.dev/weak@go1.24.2#Pointer.Value) is not guaranteed to eventually return nil, even after an object is no longer referenced, the runtime is allowed to perform a space-saving optimization that batches objects together in a single allocation slot. The weak pointer for an unreferenced object in such an allocation may never become nil if it always exists in the same batch as a referenced object. Typically, this batching only happens for tiny (on the order of 16 bytes or less) and pointer-free objects.

#### func Make

```go
func Make[T any](ptr *T) Pointer[T]
```

Make creates a weak pointer from a pointer to some value of type T.

#### (Pointer[T]) Value 

```go
func (p Pointer[T]) Value() *T
```

Value returns the original pointer used to create the weak pointer. It returns nil if the value pointed to by the original pointer was reclaimed by the garbage collector. If a weak pointer points to an object with a finalizer, then Value will return nil as soon as the object's finalizer is queued for execution.
