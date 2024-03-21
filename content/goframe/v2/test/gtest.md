+++
title = "gtest"
date = 2024-03-21T17:58:09+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/test/gtest

Package gtest provides convenient test utilities for unit testing.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Assert 

``` go
func Assert(value, expect interface{})
```

Assert checks `value` and `expect` EQUAL.

##### func AssertEQ 

``` go
func AssertEQ(value, expect interface{})
```

AssertEQ checks `value` and `expect` EQUAL, including their TYPES.

##### func AssertGE 

``` go
func AssertGE(value, expect interface{})
```

AssertGE checks `value` is GREATER OR EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertGTE, others are invalid.

##### func AssertGT 

``` go
func AssertGT(value, expect interface{})
```

AssertGT checks `value` is GREATER THAN `expect`. Notice that, only string, integer and float types can be compared by AssertGT, others are invalid.

##### func AssertIN 

``` go
func AssertIN(value, expect interface{})
```

AssertIN checks `value` is IN `expect`. The `expect` should be a slice, but the `value` can be a slice or a basic type variable. TODO map support. TODO: gconv.Strings(0) is not [0]

##### func AssertLE 

``` go
func AssertLE(value, expect interface{})
```

AssertLE checks `value` is LESS OR EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertLTE, others are invalid.

##### func AssertLT 

``` go
func AssertLT(value, expect interface{})
```

AssertLT checks `value` is LESS EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertLT, others are invalid.

##### func AssertNE 

``` go
func AssertNE(value, expect interface{})
```

AssertNE checks `value` and `expect` NOT EQUAL.

##### func AssertNI 

``` go
func AssertNI(value, expect interface{})
```

AssertNI checks `value` is NOT IN `expect`. The `expect` should be a slice, but the `value` can be a slice or a basic type variable. TODO map support.

##### func AssertNQ 

``` go
func AssertNQ(value, expect interface{})
```

AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.

##### func AssertNil 

``` go
func AssertNil(value interface{})
```

AssertNil asserts `value` is nil.

##### func C 

``` go
func C(t *testing.T, f func(t *T))
```

C creates a unit testing case. The parameter `t` is the pointer to testing.T of stdlib (*testing.T). The parameter `f` is the closure function for unit testing case.

##### func DataContent <-2.0.5

``` go
func DataContent(names ...string) string
```

DataContent retrieves and returns the file content for specified testdata path of current package

##### func DataPath <-2.0.5

``` go
func DataPath(names ...string) string
```

DataPath retrieves and returns the testdata path of current package, which is used for unit testing cases only. The optional parameter `names` specifies the sub-folders/sub-files, which will be joined with current system separator and returned with the path.

##### func Error 

``` go
func Error(message ...interface{})
```

Error panics with given `message`.

##### func Fatal 

``` go
func Fatal(message ...interface{})
```

Fatal prints `message` to stderr and exit the process.

### Types 

#### type T 

``` go
type T struct {
	*testing.T
}
```

T is the testing unit case management object.

##### (*T) Assert 

``` go
func (t *T) Assert(value, expect interface{})
```

Assert checks `value` and `expect` EQUAL.

##### (*T) AssertEQ 

``` go
func (t *T) AssertEQ(value, expect interface{})
```

AssertEQ checks `value` and `expect` EQUAL, including their TYPES.

##### (*T) AssertGE 

``` go
func (t *T) AssertGE(value, expect interface{})
```

AssertGE checks `value` is GREATER OR EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertGTE, others are invalid.

##### (*T) AssertGT 

``` go
func (t *T) AssertGT(value, expect interface{})
```

AssertGT checks `value` is GREATER THAN `expect`. Notice that, only string, integer and float types can be compared by AssertGT, others are invalid.

##### (*T) AssertIN 

``` go
func (t *T) AssertIN(value, expect interface{})
```

AssertIN checks `value` is IN `expect`. The `expect` should be a slice, but the `value` can be a slice or a basic type variable.

##### (*T) AssertLE 

``` go
func (t *T) AssertLE(value, expect interface{})
```

AssertLE checks `value` is LESS OR EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertLTE, others are invalid.

##### (*T) AssertLT 

``` go
func (t *T) AssertLT(value, expect interface{})
```

AssertLT checks `value` is LESS EQUAL THAN `expect`. Notice that, only string, integer and float types can be compared by AssertLT, others are invalid.

##### (*T) AssertNE 

``` go
func (t *T) AssertNE(value, expect interface{})
```

AssertNE checks `value` and `expect` NOT EQUAL.

##### (*T) AssertNI 

``` go
func (t *T) AssertNI(value, expect interface{})
```

AssertNI checks `value` is NOT IN `expect`. The `expect` should be a slice, but the `value` can be a slice or a basic type variable.

##### (*T) AssertNQ 

``` go
func (t *T) AssertNQ(value, expect interface{})
```

AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.

##### (*T) AssertNil 

``` go
func (t *T) AssertNil(value interface{})
```

AssertNil asserts `value` is nil.

##### (*T) Error 

``` go
func (t *T) Error(message ...interface{})
```

Error panics with given `message`.

##### (*T) Fatal 

``` go
func (t *T) Fatal(message ...interface{})
```

Fatal prints `message` to stderr and exit the process.