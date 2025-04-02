+++
title = "cmp"
date = 2023-11-05T14:26:38+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文：[https://pkg.go.dev/cmp@go1.24.2](https://pkg.go.dev/cmp@go1.24.2)

> 注意
>
>  	从go1.21.0开始才有该包。

## 概述

Package cmp provides types and functions related to comparing ordered values.

​	`cmp`包提供了与比较有序值相关的类型和函数。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Compare 

``` go
func Compare[T Ordered](x, y T) int
```

Compare returns

​	`Compare`函数返回以下值：

```
-1 if x is less than y,
如果x小于y，返回-1；
 0 if x equals y,
 如果x等于y，返回0；
+1 if x is greater than y.
如果x大于y，返回+1。
```

For floating-point types, a NaN is considered less than any non-NaN, a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.

​	对于浮点类型，`NaN`被视为小于任何非`NaN`，`NaN`被视为等于`NaN`，`-0.0`等于`0.0`。

### func Less 

``` go
func Less[T Ordered](x, y T) bool
```

Less reports whether x is less than y. For floating-point types, a NaN is considered less than any non-NaN, and -0.0 is not less than (is equal to) 0.0.

​	`Less`函数报告`x`是否小于`y`。对于浮点类型，`NaN`被视为小于任何非`NaN`，并且`-0.0`不小于（等于）`0.0`。

### func Or <-go1.22.0

``` go
func Or[T comparable](vals ...T) T
```

Or returns the first of its arguments that is not equal to the zero value. If no argument is non-zero, it returns the zero value.

​	`Or` 函数返回其第一个不等于零值的实参。如果没有实参为非零，则返回零值。

#### Or Example 

```go
package main

import (
	"cmp"
	"fmt"
)

func main() {
	// Suppose we have some user input
	// that may or may not be an empty string
	userInput1 := ""
	userInput2 := "some text"

	fmt.Println(cmp.Or(userInput1, "default"))
	fmt.Println(cmp.Or(userInput2, "default"))
	fmt.Println(cmp.Or(userInput1, userInput2, "default"))
}

Output:

default
some text
some text
```

#### Or Example (Sort)

```go
import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Order struct {
		Product  string
		Customer string
		Price    float64
	}
	orders := []Order{
		{"foo", "alice", 1.00},
		{"bar", "bob", 3.00},
		{"baz", "carol", 4.00},
		{"foo", "alice", 2.00},
		{"bar", "carol", 1.00},
		{"foo", "bob", 4.00},
	}
	// Sort by customer first, product second, and last by higher price
	slices.SortFunc(orders, func(a, b Order) int {
		return cmp.Or(
			cmp.Compare(a.Customer, b.Customer),
			cmp.Compare(a.Product, b.Product),
			cmp.Compare(b.Price, a.Price),
		)
	})
	for _, order := range orders {
		fmt.Printf("%s %s %.2f\n", order.Product, order.Customer, order.Price)
	}

}

Output:

foo alice 2.00
foo alice 1.00
bar bob 3.00
foo bob 4.00
bar carol 1.00
baz carol 4.00
```



## 类型

### type Ordered 

``` go
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
```

Ordered is a constraint that permits any ordered type: any type that supports the operators < <= >= >. If future releases of Go add new ordered types, this constraint will be modified to include them.

​	`Ordered`是一个约束，它允许任何有序类型：任何支持`<` `<=` `>=` `>`操作符的类型。如果Go的未来版本添加了新的有序类型，此约束将被修改以包含它们。

Note that floating-point types may contain NaN ("not-a-number") values. An operator such as == or < will always report false when comparing a NaN value with any other value, NaN or not. See the [Compare](https://pkg.go.dev/cmp@go1.21.3#Compare) function for a consistent way to compare NaN values.

​	请注意，浮点类型可能包含`NaN`（“非数字”）值。当将`NaN`值与其他值（`NaN`或非`NaN`）进行比较时，`==`或`<`等操作符将始终报告`false`。请使用[Compare](#func-compare)函数作为比较`NaN`值的一致方式。