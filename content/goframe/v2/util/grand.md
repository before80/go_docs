+++
title = "grand"
date = 2024-03-21T17:59:49+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/grand](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/grand)

Package grand provides high performance random bytes/number/string generation functionality.

​	Package grand 提供高性能的随机字节/数字/字符串生成功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func B

```go
func B(n int) []byte
```

B retrieves and returns random bytes of given length `n`.

​	B 检索并返回给定长度 `n` 的随机字节。

#### func D

```go
func D(min, max time.Duration) time.Duration
```

D returns a random time.Duration between min and max: [min, max].

​	D 返回一个随机时间。最小值和最大值之间的持续时间：[最小值，最大值]。

#### func Digits

```go
func Digits(n int) string
```

Digits returns a random string which contains only digits, and its length is `n`.

​	Digits 返回一个仅包含数字的随机字符串，其长度为 `n` 。

#### func Intn

```go
func Intn(max int) int
```

Intn returns an int number which is between 0 and max: [0, max).

​	Intn 返回一个介于 0 和 max 之间的 int 数：[0， max]。

Note that: 1. The `max` can only be greater than 0, or else it returns `max` directly; 2. The result is greater than or equal to 0, but less than `max`; 3. The result number is 32bit and less than math.MaxUint32.

​	注意：1.只能 `max` 大于0，否则直接返回 `max` ;2.结果大于或等于0，但小于 `max` ;3. 结果数为 32 位，小于数学。MaxUint32 中。

#### func Letters

```go
func Letters(n int) string
```

Letters returns a random string which contains only letters, and its length is `n`.

​	Letters 返回一个仅包含字母的随机字符串，其长度为 `n` 。

#### func Meet

```go
func Meet(num, total int) bool
```

Meet randomly calculate whether the given probability `num`/`total` is met.

​	开会随机计算是否满足给定的概率 `num` / `total` 。

#### func MeetProb

```go
func MeetProb(prob float32) bool
```

MeetProb randomly calculate whether the given probability is met.

​	MeetProb 随机计算是否满足给定概率。

#### func N

```go
func N(min, max int) int
```

N returns a random int between min and max: [min, max]. The `min` and `max` also support negative numbers.

​	N 返回一个介于 min 和 max 之间的随机 int：[min， max]。 `min` 和 `max` 也支持负数。

#### func Perm

```go
func Perm(n int) []int
```

Perm returns, as a slice of n int numbers, a pseudo-random permutation of the integers [0,n). TODO performance improving for large slice producing.

​	Perm 作为 n 个整数的切片返回整数 [0，n] 的伪随机排列。TODO性能提升，实现大切片生产。

#### func S

```go
func S(n int, symbols ...bool) string
```

S returns a random string which contains digits and letters, and its length is `n`. The optional parameter `symbols` specifies whether the result could contain symbols, which is false in default.

​	S 返回一个包含数字和字母的随机字符串，其长度为 `n` 。可选参数 `symbols` 指定结果是否可以包含符号，默认为 false。

#### func Str

```go
func Str(s string, n int) string
```

Str randomly picks and returns `n` count of chars from given string `s`. It also supports unicode string like Chinese/Russian/Japanese, etc.

​	Str 从给定字符串 `s` 中随机选择并返回 `n` 字符计数。它还支持Unicode字符串，如中文/俄文/日文等。

#### func Symbols

```go
func Symbols(n int) string
```

Symbols returns a random string which contains only symbols, and its length is `n`.

​	Symbols 返回一个仅包含符号的随机字符串，其长度为 `n` 。

## 类型

This section is empty.