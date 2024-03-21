+++
title = "grand"
date = 2024-03-21T17:59:49+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/grand

Package grand provides high performance random bytes/number/string generation functionality.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func B 

``` go
func B(n int) []byte
```

B retrieves and returns random bytes of given length `n`.

##### func D 

``` go
func D(min, max time.Duration) time.Duration
```

D returns a random time.Duration between min and max: [min, max].

##### func Digits 

``` go
func Digits(n int) string
```

Digits returns a random string which contains only digits, and its length is `n`.

##### func Intn 

``` go
func Intn(max int) int
```

Intn returns an int number which is between 0 and max: [0, max).

Note that: 1. The `max` can only be greater than 0, or else it returns `max` directly; 2. The result is greater than or equal to 0, but less than `max`; 3. The result number is 32bit and less than math.MaxUint32.

##### func Letters 

``` go
func Letters(n int) string
```

Letters returns a random string which contains only letters, and its length is `n`.

##### func Meet 

``` go
func Meet(num, total int) bool
```

Meet randomly calculate whether the given probability `num`/`total` is met.

##### func MeetProb 

``` go
func MeetProb(prob float32) bool
```

MeetProb randomly calculate whether the given probability is met.

##### func N 

``` go
func N(min, max int) int
```

N returns a random int between min and max: [min, max]. The `min` and `max` also support negative numbers.

##### func Perm 

``` go
func Perm(n int) []int
```

Perm returns, as a slice of n int numbers, a pseudo-random permutation of the integers [0,n). TODO performance improving for large slice producing.

##### func S 

``` go
func S(n int, symbols ...bool) string
```

S returns a random string which contains digits and letters, and its length is `n`. The optional parameter `symbols` specifies whether the result could contain symbols, which is false in default.

##### func Str 

``` go
func Str(s string, n int) string
```

Str randomly picks and returns `n` count of chars from given string `s`. It also supports unicode string like Chinese/Russian/Japanese, etc.

##### func Symbols 

``` go
func Symbols(n int) string
```

Symbols returns a random string which contains only symbols, and its length is `n`.

### Types 

This section is empty.