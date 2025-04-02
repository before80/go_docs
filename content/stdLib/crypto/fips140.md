+++
title = "fips140"
date = 2025-04-01T13:15:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
> 原文：[https://pkg.go.dev/crypto/fips140@go1.24.2](https://pkg.go.dev/crypto/fips140@go1.24.2)

> 注意
>
> ​	从go1.24.0开始才可以使用该包。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Enabled

```go
func Enabled() bool
```

Enabled reports whether the cryptography libraries are operating in FIPS 140-3 mode.

It can be controlled at runtime using the GODEBUG setting "fips140". If set to "on", FIPS 140-3 mode is enabled. If set to "only", non-approved cryptography functions will additionally return errors or panic.

This can't be changed after the program has started.

## 类型

This section is empty.