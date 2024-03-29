+++
title = "gregex"
date = 2024-03-21T17:58:48+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/text/gregex](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/text/gregex)

Package gregex provides high performance API for regular expression functionality.

​	软件包 gregex 为正则表达式功能提供高性能 API。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func IsMatch

```go
func IsMatch(pattern string, src []byte) bool
```

IsMatch checks whether given bytes `src` matches `pattern`.

​	IsMatch 检查给定的字节是否 `src` 匹配 `pattern` 。

##### Example

``` go
```

#### func IsMatchString

```go
func IsMatchString(pattern string, src string) bool
```

IsMatchString checks whether given string `src` matches `pattern`.

​	IsMatchString 检查给定的字符串 `src` 是否匹配 `pattern` 。

##### Example

``` go
```

#### func Match

```go
func Match(pattern string, src []byte) ([][]byte, error)
```

Match return bytes slice that matched `pattern`.

​	匹配匹配 `pattern` 的返回字节切片。

##### Example

``` go
```

#### func MatchAll

```go
func MatchAll(pattern string, src []byte) ([][][]byte, error)
```

MatchAll return all bytes slices that matched `pattern`.

​	MatchAll 返回匹配 `pattern` 的所有字节切片。

##### Example

``` go
```

#### func MatchAllString

```go
func MatchAllString(pattern string, src string) ([][]string, error)
```

MatchAllString return all strings that matched `pattern`.

​	MatchAllString 返回匹配 `pattern` 的所有字符串。

##### Example

``` go
```

#### func MatchString

```go
func MatchString(pattern string, src string) ([]string, error)
```

MatchString return strings that matched `pattern`.

​	MatchString 返回匹配 `pattern` 的字符串。

##### Example

``` go
```

#### func Quote

```go
func Quote(s string) string
```

Quote quotes `s` by replacing special chars in `s` to match the rules of regular expression pattern. And returns the copy.

​	 `s` 通过替换特殊字符 `s` 来引用引号以匹配正则表达式模式的规则。并返回副本。

Eg: Quote(`[foo]`) returns `\[foo\]`.

​	例如：Quote（ `[foo]` ） 返回 `\[foo\]` 。

##### Example

``` go
```

#### func Replace

```go
func Replace(pattern string, replace, src []byte) ([]byte, error)
```

Replace replaces all matched `pattern` in bytes `src` with bytes `replace`.

​	Replace 将所有匹配 `pattern` 的字节替换 `src` 为 `replace` bytes 。

##### Example

``` go
```

#### func ReplaceFunc

```go
func ReplaceFunc(pattern string, src []byte, replaceFunc func(b []byte) []byte) ([]byte, error)
```

ReplaceFunc replace all matched `pattern` in bytes `src` with custom replacement function `replaceFunc`.

​	ReplaceFunc `src` 将所有匹配 `pattern` 的字节替换为自定义替换函数 `replaceFunc` 。

##### Example

``` go
```

#### func ReplaceFuncMatch

```go
func ReplaceFuncMatch(pattern string, src []byte, replaceFunc func(match [][]byte) []byte) ([]byte, error)
```

ReplaceFuncMatch replace all matched `pattern` in bytes `src` with custom replacement function `replaceFunc`. The parameter `match` type for `replaceFunc` is [][]byte, which is the result contains all sub-patterns of `pattern` using Match function.

​	ReplaceFuncMatch `src` 使用自定义替换函数 `replaceFunc` 替换所有匹配 `pattern` 的字节。的 `replaceFunc` 参数 `match` 类型为 [][]byte，即结果包含 `pattern` 使用 Match 函数的所有子模式。

##### Example

``` go
```

#### func ReplaceString

```go
func ReplaceString(pattern, replace, src string) (string, error)
```

ReplaceString replace all matched `pattern` in string `src` with string `replace`.

​	ReplaceString 将 string 中所有匹配 `pattern` 的字符串 `src` 替换为 string `replace` 。

##### Example

``` go
```

#### func ReplaceStringFunc

```go
func ReplaceStringFunc(pattern string, src string, replaceFunc func(s string) string) (string, error)
```

ReplaceStringFunc replace all matched `pattern` in string `src` with custom replacement function `replaceFunc`.

​	ReplaceStringFunc 使用自定义替换函数 `replaceFunc` 替换字符串 `src` 中的所有匹配 `pattern` 项。

##### Example

``` go
```

#### func ReplaceStringFuncMatch

```go
func ReplaceStringFuncMatch(pattern string, src string, replaceFunc func(match []string) string) (string, error)
```

ReplaceStringFuncMatch replace all matched `pattern` in string `src` with custom replacement function `replaceFunc`. The parameter `match` type for `replaceFunc` is []string, which is the result contains all sub-patterns of `pattern` using MatchString function.

​	ReplaceStringFuncMatch 将字符串 `src` 中的所有匹配 `pattern` 替换为自定义替换函数 `replaceFunc` 。的 `replaceFunc` 参数 `match` 类型为 []string，即结果包含使用 MatchString 函数的所有 `pattern` 子模式。

##### Example

``` go
```

#### func Split

```go
func Split(pattern string, src string) []string
```

Split slices `src` into substrings separated by the expression and returns a slice of the substrings between those expression matches.

​	将 `src` 切片拆分为由表达式分隔的子字符串，并在这些表达式匹配项之间返回子字符串的切片。

##### Example

``` go
```

#### func Validate

```go
func Validate(pattern string) error
```

Validate checks whether given regular expression pattern `pattern` valid.

​	验证检查给定的正则表达式模式 `pattern` 是否有效。

##### Example

``` go
```

## 类型

This section is empty.