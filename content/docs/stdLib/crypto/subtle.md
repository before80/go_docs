+++
title = "subtle"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# subtle

https://pkg.go.dev/crypto/subtle@go1.20.1



Package subtle implements functions that are often useful in cryptographic code but require careful thought to use correctly.











## 常量 ¶

This section is empty.

## 变量

This section is empty.

## 函数

#### func ConstantTimeByteEq [¶](https://pkg.go.dev/crypto/subtle@go1.20.1#ConstantTimeByteEq)

```
func ConstantTimeByteEq(x, y uint8) int
```

ConstantTimeByteEq returns 1 if x == y and 0 otherwise.

#### func ConstantTimeCompare [¶](https://pkg.go.dev/crypto/subtle@go1.20.1#ConstantTimeCompare)

```
func ConstantTimeCompare(x, y []byte) int
```

ConstantTimeCompare returns 1 if the two slices, x and y, have equal contents and 0 otherwise. The time taken is a function of the length of the slices and is independent of the contents. If the lengths of x and y do not match it returns 0 immediately.

#### func ConstantTimeCopy [¶](https://pkg.go.dev/crypto/subtle@go1.20.1#ConstantTimeCopy)

```
func ConstantTimeCopy(v int, x, y []byte)
```

ConstantTimeCopy copies the contents of y into x (a slice of equal length) if v == 1. If v == 0, x is left unchanged. Its behavior is undefined if v takes any other value.

#### func ConstantTimeEq [¶](https://pkg.go.dev/crypto/subtle@go1.20.1#ConstantTimeEq)

```
func ConstantTimeEq(x, y int32) int
```

ConstantTimeEq returns 1 if x == y and 0 otherwise.

#### func ConstantTimeLessOrEq [¶](https://pkg.go.dev/crypto/subtle@go1.20.1#ConstantTimeLessOrEq)added in go1.2

```
func ConstantTimeLessOrEq(x, y int) int
```

ConstantTimeLessOrEq returns 1 if x <= y and 0 otherwise. Its behavior is undefined if x or y are negative or > 2**31 - 1.

#### func ConstantTimeSelect [¶](https://pkg.go.dev/crypto/subtle@go1.20.1#ConstantTimeSelect)

```
func ConstantTimeSelect(v, x, y int) int
```

ConstantTimeSelect returns x if v == 1 and y if v == 0. Its behavior is undefined if v takes any other value.

#### func XORBytes [¶](https://pkg.go.dev/crypto/subtle@go1.20.1#XORBytes)added in go1.20

```
func XORBytes(dst, x, y []byte) int
```

XORBytes sets dst[i] = x[i] ^ y[i] for all i < n = min(len(x), len(y)), returning n, the number of bytes written to dst. If dst does not have length at least n, XORBytes panics without writing anything to dst.

## 类型

This section is empty.