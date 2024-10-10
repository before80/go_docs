+++
title = "values"
date = 2023-08-07T13:31:12+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

#  Values

```go
package main

import "fmt"

func main() {
	// 字符串
	fmt.Printf("%v,%T\n", "Hi", "Hi")
	fmt.Printf("%v,%T\n", "Hi "+"golang", "Hi "+"golang")

	// 数值
	fmt.Printf("%v,%T\n", 1, 1)
	fmt.Printf("%v,%T\n", 1.2, 1.2)
	fmt.Printf("%v,%T\n", 1.2+3.4i, 1.2+3.4i)

	// 布尔
	fmt.Printf("%v,%T\n", true, true)
	fmt.Printf("%v,%T\n", false, false)
	fmt.Printf("%v,%T\n", !true, !true)
	fmt.Printf("%v,%T\n", !false, !false)
	fmt.Printf("%v,%T\n", true && false, true && false)
	fmt.Printf("%v,%T\n", true || false, true || false)

	// 其他
	fmt.Printf("%v,%T\n", [10]int{}, [10]int{})
	fmt.Printf("%v,%T\n", []int{}, []int{})
	fmt.Printf("%v,%T\n", make([]int, 10), make([]int, 10))
	fmt.Printf("%v,%T\n", make(map[string]int, 10), make(map[string]int, 10))
	fmt.Printf("%v,%T\n", make(chan int), make(chan int))
	fmt.Printf("%v,%T\n", make(chan int, 10), make(chan int, 10))
	fmt.Printf("%v,%T\n", struct{ Name string }{"zlongx"}, struct{ Name string }{"zlongx"})
}
// Output:
//Hi,string
//Hi golang,string
//1,int
//1.2,float64
//(1.2+3.4i),complex128
//true,bool
//false,bool
//false,bool
//true,bool
//false,bool
//true,bool
//[0 0 0 0 0 0 0 0 0 0],[10]int
//[],[]int
//[0 0 0 0 0 0 0 0 0 0],[]int
//map[],map[string]int
//0xc000096120,chan int
//0xc0000e2000,chan int

```

