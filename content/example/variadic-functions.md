+++
title = "Variadic Function"
date = 2023-08-07T13:33:44+08:00
weight = 13
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Variadic Function - 可变参数函数



```go
package main

import (
	"fmt"
	"math"
)

func findMax(nums ...int) int {
	max := math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func ConcatStrs(sep string, strs ...string) (r string) {
	for i, str := range strs {
		if i > 0 {
			r += sep
		}
		r += str
	}
	return r
}

func main() {
	fmt.Println(findMax(1, 2, 3)) // 3

	fmt.Println(findMax([]int{1, 2, 3}...)) // 3

	fmt.Println(ConcatStrs("|", "a", "b", "c")) // a|b|c

	fmt.Println(ConcatStrs("|", []string{"a", "b", "c"}...)) // a|b|c
}

```

