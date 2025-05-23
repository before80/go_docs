+++
title = "maps"
date = 2023-11-05T14:27:43+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文：[https://pkg.go.dev/maps@go1.24.2](https://pkg.go.dev/maps@go1.24.2)

> 注意
>
> ​	从go1.21.0开始才有该包。

## 概述

Package maps defines various functions useful with maps of any type.

​	maps包定义了各种类型映射的有用函数。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func All <- go1.23.0

``` go
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V]
```

All returns an iterator over key-value pairs from m. The iteration order is not specified and is not guaranteed to be the same from one call to the next.

​	All 返回一个迭代器，用于遍历 m 中的键值对。迭代顺序未指定，并且从一次调用到下一次调用可能不相同。

### func Clone 

``` go
func Clone[M ~map[K]V, K comparable, V any](m M) M
```

Clone returns a copy of m. This is a shallow clone: the new keys and values are set using ordinary assignment.

​	Clone返回m的副本。这是一个浅拷贝：新键和值使用普通赋值设置。

#### func Collect <- go1.23.0

``` go
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V
```

Collect collects key-value pairs from seq into a new map and returns it.

​	Collect 将 seq 中的键值对收集到一个新的 map 中并返回。

### func Copy 

``` go
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)
```

Copy copies all key/value pairs in src adding them to dst. When a key in src is already present in dst, the value in dst will be overwritten by the value associated with the key in src.

​	`Copy`函数将`src`中的所有键/值对复制到`dst`中。当`src`中的某个键已经存在于`dst`中时，`dst`中的值将被`src`中与该键关联的值覆盖。

### func DeleteFunc 

``` go
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool)
```

DeleteFunc deletes any key/value pairs from m for which del returns true.

​	`DeleteFunc`函数从`m`中删除`del`返回`true`的任何键/值对。

#### DeleteFunc Example

``` go
package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	maps.DeleteFunc(m, func(k string, v int) bool {
		return v%2 != 0 // delete odd values
	})
	fmt.Println(m)
}
Output:

map[four:4 two:2]
```
### func Equal 

``` go
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool
```

Equal reports whether two maps contain the same key/value pairs. Values are compared using ==.

​	`Equal`函数报告两个映射是否包含相同的键/值对。值使用`==`进行比较。

### func EqualFunc 

``` go
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool
```

EqualFunc is like Equal, but compares values using eq. Keys are still compared with ==.

​	`EqualFunc`函数类似于`Equal`函数，但使用`eq`比较值。键仍然使用`==`进行比较。

#### EqualFunc  Example

```go
package main

import (
	"fmt"
	"maps"
	"strings"
)

func main() {
	m1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	m2 := map[int][]byte{
		1:    []byte("One"),
		10:   []byte("Ten"),
		1000: []byte("Thousand"),
	}
	eq := maps.EqualFunc(m1, m2, func(v1 string, v2 []byte) bool {
		return strings.ToLower(v1) == strings.ToLower(string(v2))
	})
	fmt.Println(eq)
}
Output:

true
```

#### func Insert <- go1.23.0

``` go
func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V])
```

Insert adds the key-value pairs from seq to m. If a key in seq already exists in m, its value will be overwritten.

​	Insert 将 seq 中的键值对添加到 m 中。如果 seq 中的键已经存在于 m 中，则其值会被覆盖。

#### func Keys <- go1.23.0

``` go
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K]
```

Keys returns an iterator over keys in m. The iteration order is not specified and is not guaranteed to be the same from one call to the next.

​		Keys 返回一个迭代器，用于遍历 m 中的键。迭代顺序未指定，并且从一次调用到下一次调用可能不相同。

#### func Values <- go1.23.0

``` go
func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V]
```

Values returns an iterator over values in m. The iteration order is not specified and is not guaranteed to be the same from one call to the next.

​	Values 返回一个迭代器，用于遍历 m 中的值。迭代顺序未指定，并且从一次调用到下一次调用可能不相同。

## 类型

This section is empty.

