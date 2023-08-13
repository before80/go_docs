+++
title = "generic"
date = 2023-08-07T13:36:02+08:00
weight = 22
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# generic

```go
package main

import "fmt"

// 泛型类型
type List[T any] struct {
	next  *List[T]
	value T
}

// 泛型函数
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	l := List[int]{}
	_ = l

	m1 := make(map[string]int64)
	m2 := make(map[string]float64)

	m1["A"] = 1
	m1["B"] = 2
	m1["C"] = 3
	fmt.Println(SumIntsOrFloats(m1)) // 6

	m2["A"] = 1.2
	m2["B"] = 2.3
	m2["C"] = 3.4
	fmt.Println(SumIntsOrFloats(m2)) // 6.9 或 6.8999999999999995 <- 奇怪的结果

}

```

