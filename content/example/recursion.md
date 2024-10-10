+++
title = "Recursion"
date = 2023-08-07T13:34:13+08:00
weight = 15
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++



# Recursion - 递归

```go
package main

import "fmt"

// 递归计算阶乘
func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 递归计算斐波那契数列的第 n 项
func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(factorial(3)) // 6
	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120

	fmt.Println(fibonacci(1)) // 1
	fmt.Println(fibonacci(2)) // 1
	fmt.Println(fibonacci(3)) // 2
	fmt.Println(fibonacci(4)) // 3
	fmt.Println(fibonacci(5)) // 5
	fmt.Println(fibonacci(6)) // 8

}

```

