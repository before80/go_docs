+++
title = "switch"
date = 2023-08-07T13:31:58+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# switch

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//----------- 表达式选择 ------------
	fmt.Println("----------- 表达式选择 ------------")
	// 没有变量, 类似 if else
	score := rand.Intn(101)
	switch { // 相当于 switch true {
	case score < 60:
		fmt.Println("E")
	case 60 <= score && score <= 69:
		fmt.Println("D")
	case 70 <= score && score <= 79:
		fmt.Println("C")
	case 80 <= score && score <= 89:
		fmt.Println("B")
	case 90 <= score && score <= 100:
		fmt.Println("A")
	}

	// 有变量
	i := rand.Intn(3)
	switch i {
	case 0:
		fmt.Println("i is 0.")
	case 1:
		fmt.Println("i is 1.")
	case 2:
		fmt.Println("i is 2.")
	default:
		fmt.Println("i is unknown.")
	}

	// 有变量, 且是新定义局部变量
	switch j := rand.Intn(3); j {
	case 0:
		fmt.Println("j is 0.")
	case 1:
		fmt.Println("j is 1.")
	case 2:
		fmt.Println("j is 2.")
	default:
		fmt.Println("j is unknown.")
	}

	// 有变量, 且是新定义局部变量
	switch score := rand.Intn(101); { // 相当于 switch score := rand.Intn(101); true {
	case score < 60:
		fmt.Println("E")
	case 60 <= score && score <= 69:
		fmt.Println("D")
	case 70 <= score && score <= 79:
		fmt.Println("C")
	case 80 <= score && score <= 89:
		fmt.Println("B")
	case 90 <= score && score <= 100:
		fmt.Println("A")
	}

	//----------- 类型选择 ------------
	fmt.Println("----------- 类型选择 ------------")
	var x any
	x = 28

	switch x.(type) { // 这里若使用 switch x.(type); true { ， 则会编译报错
	case int, int8, int16, int32, int64:
		fmt.Println("x's type is ints.")
	case uint, uint8, uint16, uint32, uint64:
		fmt.Println("x's type is uints.")
	case float32, float64:
		fmt.Println("x's type is floats.")
	default:
		fmt.Println("x's type is unknown.")
	}
}

```

