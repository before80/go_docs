+++
title = "gtest"
date = 2024-03-21T17:58:09+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/test/gtest](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/test/gtest)

Package gtest provides convenient test utilities for unit testing.

​	软件包 gtest 为单元测试提供了方便的测试实用程序。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Assert

```go
func Assert(value, expect interface{})
```

Assert checks `value` and `expect` EQUAL.

​	断言检查 `value` 和 `expect` EQUAL。

#### func AssertEQ

```go
func AssertEQ(value, expect interface{})
```

AssertEQ checks `value` and `expect` EQUAL, including their TYPES.

​	AssertEQ 检查 `value` 和 `expect` EQUAL，包括它们的 TYPES。

#### func AssertGE

```go
func AssertGE(value, expect interface{})
```

AssertGE checks `value` is GREATER OR EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertGTE, others are invalid.

​	AssertGE 检查 `value` 大于或等于 `expect` 。请注意，AssertGTE 只能比较字符串、整数和浮点类型，其他类型无效。

#### func AssertGT

```go
func AssertGT(value, expect interface{})
```

AssertGT checks `value` is GREATER THAN `expect`. Notice that, only string, integer and float types can be compared by AssertGT, others are invalid.

​	AssertGT 检查 `value` 大于 。 `expect` 请注意，AssertGT 只能比较字符串、整数和浮点类型，其他类型无效。

#### func AssertIN

```go
func AssertIN(value, expect interface{})
```

AssertIN checks `value` is IN `expect`. The `expect` should be a slice, but the `value` can be a slice or a basic type variable. TODO map support. TODO: gconv.Strings(0) is not [0]

​	AssertIN 检查 `value` 为 IN `expect` 。应该是 `expect` 切片，但可以 `value` 是切片或基本类型变量。TODO地图支持。待办事项： gconv.Strings（0） 不是 [0]

#### func AssertLE

```go
func AssertLE(value, expect interface{})
```

AssertLE checks `value` is LESS OR EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertLTE, others are invalid.

​	AssertLE 检查 `value` 小于或等于 `expect` 。请注意，AssertLTE 只能比较字符串、整数和浮点类型，其他类型无效。

#### func AssertLT

```go
func AssertLT(value, expect interface{})
```

AssertLT checks `value` is LESS EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertLT, others are invalid.

​	AssertLT 检查 `value` 小于 `expect` 。请注意，AssertLT 只能比较字符串、整数和浮点类型，其他类型无效。

#### func AssertNE

```go
func AssertNE(value, expect interface{})
```

AssertNE checks `value` and `expect` NOT EQUAL.

​	AssertNE 检查 `value` 和 `expect` NOT EQUAL。

#### func AssertNI

```go
func AssertNI(value, expect interface{})
```

AssertNI checks `value` is NOT IN `expect`. The `expect` should be a slice, but the `value` can be a slice or a basic type variable. TODO map support.

​	AssertNI 检查 `value` 不在 `expect` 。应该是 `expect` 切片，但可以 `value` 是切片或基本类型变量。TODO地图支持。

#### func AssertNQ

```go
func AssertNQ(value, expect interface{})
```

AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.

​	AssertNQ 检查 `value` 和 `expect` NOT EQUAL，包括它们的 TYPES。

#### func AssertNil

```go
func AssertNil(value interface{})
```

AssertNil asserts `value` is nil.

​	AssertNil 断言 `value` 为 nil。

#### func C

```go
func C(t *testing.T, f func(t *T))
```

C creates a unit testing case. The parameter `t` is the pointer to testing.T of stdlib (*testing.T). The parameter `f` is the closure function for unit testing case.

​	C 创建一个单元测试用例。该参数 `t` 是指向测试的指针。T 的 stdlib （*testing.该参数 `f` 是单元测试用例的闭包函数。

#### func DataContent <-2.0.5

```go
func DataContent(names ...string) string
```

DataContent retrieves and returns the file content for specified testdata path of current package

​	DataContent 检索并返回当前包的指定 testdata 路径的文件内容

#### func DataPath <-2.0.5

```go
func DataPath(names ...string) string
```

DataPath retrieves and returns the testdata path of current package, which is used for unit testing cases only. The optional parameter `names` specifies the sub-folders/sub-files, which will be joined with current system separator and returned with the path.

​	DataPath 检索并返回当前包的 testdata 路径，该路径仅用于单元测试用例。optional 参数 `names` 指定子文件夹/子文件，这些子文件夹/子文件将与当前系统分隔符联接并与路径一起返回。

#### func Error

```go
func Error(message ...interface{})
```

Error panics with given `message`.

​	给定 `message` 的 .

#### func Fatal

```go
func Fatal(message ...interface{})
```

Fatal prints `message` to stderr and exit the process.

​	致命打印 `message` 到 stderr 并退出进程。

## 类型

### type T

```go
type T struct {
	*testing.T
}
```

T is the testing unit case management object.

​	T为测试单元案例管理对象。

#### (*T) Assert

```go
func (t *T) Assert(value, expect interface{})
```

Assert checks `value` and `expect` EQUAL.

​	断言检查 `value` 和 `expect` EQUAL。

#### (*T) AssertEQ

```go
func (t *T) AssertEQ(value, expect interface{})
```

AssertEQ checks `value` and `expect` EQUAL, including their TYPES.

​	AssertEQ 检查 `value` 和 `expect` EQUAL，包括它们的 TYPES。

#### (*T) AssertGE

```go
func (t *T) AssertGE(value, expect interface{})
```

AssertGE checks `value` is GREATER OR EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertGTE, others are invalid.

​	AssertGE 检查 `value` 大于或等于 `expect` 。请注意，AssertGTE 只能比较字符串、整数和浮点类型，其他类型无效。

#### (*T) AssertGT

```go
func (t *T) AssertGT(value, expect interface{})
```

AssertGT checks `value` is GREATER THAN `expect`. Notice that, only string, integer and float types can be compared by AssertGT, others are invalid.

​	AssertGT 检查 `value` 大于 。 `expect` 请注意，AssertGT 只能比较字符串、整数和浮点类型，其他类型无效。

#### (*T) AssertIN

```go
func (t *T) AssertIN(value, expect interface{})
```

AssertIN checks `value` is IN `expect`. The `expect` should be a slice, but the `value` can be a slice or a basic type variable.

​	AssertIN 检查 `value` 为 IN `expect` 。应该是 `expect` 切片，但可以 `value` 是切片或基本类型变量。

#### (*T) AssertLE

```go
func (t *T) AssertLE(value, expect interface{})
```

AssertLE checks `value` is LESS OR EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertLTE, others are invalid.

​	AssertLE 检查 `value` 小于或等于 `expect` 。请注意，AssertLTE 只能比较字符串、整数和浮点类型，其他类型无效。

#### (*T) AssertLT

```go
func (t *T) AssertLT(value, expect interface{})
```

AssertLT checks `value` is LESS EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertLT, others are invalid.

​	AssertLT 检查 `value` 小于 `expect` 。请注意，AssertLT 只能比较字符串、整数和浮点类型，其他类型无效。

#### (*T) AssertNE

```go
func (t *T) AssertNE(value, expect interface{})
```

AssertNE checks `value` and `expect` NOT EQUAL.

​	AssertNE 检查 `value` 和 `expect` NOT EQUAL。

#### (*T) AssertNI

```go
func (t *T) AssertNI(value, expect interface{})
```

AssertNI checks `value` is NOT IN `expect`. The `expect` should be a slice, but the `value` can be a slice or a basic type variable.

​	AssertNI 检查 `value` 不在 `expect` 。应该是 `expect` 切片，但可以 `value` 是切片或基本类型变量。

#### (*T) AssertNQ

```go
func (t *T) AssertNQ(value, expect interface{})
```

AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.

​	AssertNQ 检查 `value` 和 `expect` NOT EQUAL，包括它们的 TYPES。

#### (*T) AssertNil

```go
func (t *T) AssertNil(value interface{})
```

AssertNil asserts `value` is nil.

​	AssertNil 断言 `value` 为 nil。

#### (*T) Error

```go
func (t *T) Error(message ...interface{})
```

Error panics with given `message`.

​	给定 `message` 的 .

#### (*T) Fatal

```go
func (t *T) Fatal(message ...interface{})
```

Fatal prints `message` to stderr and exit the process.

​	致命打印 `message` 到 stderr 并退出进程。