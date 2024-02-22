+++
title = "subtle"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/subtle@go1.21.3](https://pkg.go.dev/crypto/subtle@go1.21.3)

Package subtle implements functions that are often useful in cryptographic code but require careful thought to use correctly.

​	Package subtle 实现通常在加密代码中很有用但需要仔细考虑才能正确使用的函数。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func ConstantTimeByteEq 

``` go
func ConstantTimeByteEq(x, y uint8) int
```

ConstantTimeByteEq returns 1 if x == y and 0 otherwise.

​	ConstantTimeByteEq 在 x == y 时返回 1，否则返回 0。

### func ConstantTimeCompare 

``` go
func ConstantTimeCompare(x, y []byte) int
```

ConstantTimeCompare returns 1 if the two slices, x and y, have equal contents and 0 otherwise. The time taken is a function of the length of the slices and is independent of the contents. If the lengths of x and y do not match it returns 0 immediately.

​	ConstantTimeCompare 在两个切片 x 和 y 内容相等时返回 1，否则返回 0。所花费的时间是切片长度的函数，与内容无关。如果 x 和 y 的长度不匹配，它会立即返回 0。

### func ConstantTimeCopy 

``` go
func ConstantTimeCopy(v int, x, y []byte)
```

ConstantTimeCopy copies the contents of y into x (a slice of equal length) if v == 1. If v == 0, x is left unchanged. Its behavior is undefined if v takes any other value.

​	ConstantTimeCopy 在 v == 1 时将 y 的内容复制到 x（长度相等的切片）中。如果 v == 0，则 x 保持不变。如果 v 采用任何其他值，则其行为未定义。

### func ConstantTimeEq 

``` go
func ConstantTimeEq(x, y int32) int
```

ConstantTimeEq returns 1 if x == y and 0 otherwise.

​	ConstantTimeEq 在 x == y 时返回 1，否则返回 0。

### func ConstantTimeLessOrEq  <- go1.2

``` go
func ConstantTimeLessOrEq(x, y int) int
```

ConstantTimeLessOrEq returns 1 if x <= y and 0 otherwise. Its behavior is undefined if x or y are negative or > $2^{31} - 1$.

​	ConstantTimeLessOrEq 在 x <= y 时返回 1，否则返回 0。如果 x 或 y 为负数或 > $2^{31} - 1$，则其行为未定义。

### func ConstantTimeSelect 

``` go
func ConstantTimeSelect(v, x, y int) int
```

ConstantTimeSelect returns x if v == 1 and y if v == 0. Its behavior is undefined if v takes any other value.

​	如果 v == 1，则 ConstantTimeSelect 返回 x；如果 v == 0，则返回 y。如果 v 采用任何其他值，则其行为未定义。

### func XORBytes  <- go1.20

``` go
func XORBytes(dst, x, y []byte) int
```

XORBytes sets dst[i] = x[i] ^ y[i] for all i < n = min(len(x), len(y)), returning n, the number of bytes written to dst. If dst does not have length at least n, XORBytes panics without writing anything to dst.

​	XORBytes 将 dst[i] = x[i] ^ y[i] 设置为所有 i < n = min(len(x), len(y))，返回 n，即写入 dst 的字节数。如果 dst 的长度至少为 n，则 XORBytes 会在不向 dst 写入任何内容的情况下引发 panic。

## 类型

This section is empty.