+++
title = "gregex"
date = 2024-03-21T17:58:48+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/text/gregex

Package gregex provides high performance API for regular expression functionality.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func IsMatch 

``` go
func IsMatch(pattern string, src []byte) bool
```

IsMatch checks whether given bytes `src` matches `pattern`.

##### Example

``` go
```
##### func IsMatchString 

``` go
func IsMatchString(pattern string, src string) bool
```

IsMatchString checks whether given string `src` matches `pattern`.

##### Example

``` go
```
##### func Match 

``` go
func Match(pattern string, src []byte) ([][]byte, error)
```

Match return bytes slice that matched `pattern`.

##### Example

``` go
```
##### func MatchAll 

``` go
func MatchAll(pattern string, src []byte) ([][][]byte, error)
```

MatchAll return all bytes slices that matched `pattern`.

##### Example

``` go
```
##### func MatchAllString 

``` go
func MatchAllString(pattern string, src string) ([][]string, error)
```

MatchAllString return all strings that matched `pattern`.

##### Example

``` go
```
##### func MatchString 

``` go
func MatchString(pattern string, src string) ([]string, error)
```

MatchString return strings that matched `pattern`.

##### Example

``` go
```
##### func Quote 

``` go
func Quote(s string) string
```

Quote quotes `s` by replacing special chars in `s` to match the rules of regular expression pattern. And returns the copy.

Eg: Quote(`[foo]`) returns `\[foo\]`.

##### Example

``` go
```
##### func Replace 

``` go
func Replace(pattern string, replace, src []byte) ([]byte, error)
```

Replace replaces all matched `pattern` in bytes `src` with bytes `replace`.

##### Example

``` go
```
##### func ReplaceFunc 

``` go
func ReplaceFunc(pattern string, src []byte, replaceFunc func(b []byte) []byte) ([]byte, error)
```

ReplaceFunc replace all matched `pattern` in bytes `src` with custom replacement function `replaceFunc`.

##### Example

``` go
```
##### func ReplaceFuncMatch 

``` go
func ReplaceFuncMatch(pattern string, src []byte, replaceFunc func(match [][]byte) []byte) ([]byte, error)
```

ReplaceFuncMatch replace all matched `pattern` in bytes `src` with custom replacement function `replaceFunc`. The parameter `match` type for `replaceFunc` is [][]byte, which is the result contains all sub-patterns of `pattern` using Match function.

##### Example

``` go
```
##### func ReplaceString 

``` go
func ReplaceString(pattern, replace, src string) (string, error)
```

ReplaceString replace all matched `pattern` in string `src` with string `replace`.

##### Example

``` go
```
##### func ReplaceStringFunc 

``` go
func ReplaceStringFunc(pattern string, src string, replaceFunc func(s string) string) (string, error)
```

ReplaceStringFunc replace all matched `pattern` in string `src` with custom replacement function `replaceFunc`.

##### Example

``` go
```
##### func ReplaceStringFuncMatch 

``` go
func ReplaceStringFuncMatch(pattern string, src string, replaceFunc func(match []string) string) (string, error)
```

ReplaceStringFuncMatch replace all matched `pattern` in string `src` with custom replacement function `replaceFunc`. The parameter `match` type for `replaceFunc` is []string, which is the result contains all sub-patterns of `pattern` using MatchString function.

##### Example

``` go
```
##### func Split 

``` go
func Split(pattern string, src string) []string
```

Split slices `src` into substrings separated by the expression and returns a slice of the substrings between those expression matches.

##### Example

``` go
```
##### func Validate 

``` go
func Validate(pattern string) error
```

Validate checks whether given regular expression pattern `pattern` valid.

##### Example

``` go
```
### Types 

This section is empty.