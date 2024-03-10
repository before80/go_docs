+++
title = "function"
date = 2023-08-07T13:32:45+08:00
weight = 12
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# function

```go
package main

import "fmt"

// 求两 int 类型变量之和，未命名返回的变量
func AddInt1(a, b int) int {
	return a + b
}

// 求两 int 类型变量之和，命名返回的变量
func AddInt2(a, b int) (c int) {
	c = a + b
	return
}

// 交换两个 int 类型的数值
func SwapInt(a, b int) (int, int) {
	return b, a
}

// 交换两个 float64 类型的数值
func SwapFloat64(a, b float64) (float64, float64) {
	return b, a
}

// 泛型函数：用于交换两个类型都是 int 或 float64 的数值
func Swap[T int | float64](a, b T) (T, T) {
	return b, a
}

// 泛型函数：用于求和类型为 map[K]V 的各个元素的数值之和
func SumIntsOrFloats[K comparable, V int | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	fmt.Println(AddInt1(1, 2))         // 3
	fmt.Println(AddInt2(1, 2))         // 3
	fmt.Println(SwapInt(1, 2))         // 2 1
	fmt.Println(SwapFloat64(1.2, 3.4)) // 3.4 1.2
	fmt.Println(Swap(1, 2))            // 2 1
	fmt.Println(Swap(1.2, 3.4))        // 3.4 1.2

	m1 := make(map[string]int)
	m1["A"] = 1
	m1["B"] = 2
	m1["C"] = 3
	fmt.Println(SumIntsOrFloats(m1)) // 6

	m2 := make(map[int]float64)
	m2[1] = 1.2
	m2[2] = 2.3
	m2[3] = 3.4
	fmt.Println(SumIntsOrFloats(m2)) // 6.8999999999999995 或 6.9 =》 好奇怪 !
}

```

