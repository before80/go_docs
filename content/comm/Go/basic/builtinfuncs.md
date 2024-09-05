+++
title = "内置函数"
date = 2024-08-19T09:36:18+08:00
weight = 80
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 截止至go1.23，目前有18个内置函数！append(slice []Type, elems ...Type) []Type

## `append(slice []Type, elems ...Type) []Type`

> 注意
>
> ​	仅用于切片类型变量追加元素。

### `elems`为切片+`...`

```go

```

### `elems`为字符串+`...`

```go

```



## `cap(v Type) int`



## `clear[T ~[]Type | ~map[Type]Type1](t T)`

## `close(c chan<- Type)`

## `complex(r, i FloatType) ComplexType`

## `copy(dst, src []Type) int`

## `delete(m map[Type]Type1, key Type)`

## `imag(c ComplexType) FloatType`

## `len(v Type) int`

## `make(t Type, size ...IntegerType) Type`

## `max[T cmp.Ordered](x T, y ...T) T`

## `min[T cmp.Ordered](x T, y ...T) T`

## `new(Type) *Type`

## `panic(v any)`

## `print(args ...Type)`

## `println(args ...Type)`

## `real(c ComplexType) FloatType`

## `recover() any`

