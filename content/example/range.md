+++
title = "range"
date = 2023-08-07T13:32:37+08:00
weight = 11
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# range

```go
package main

import "fmt"

func main() {
	// 字符串
	// 纯ASCII字符的字符串
	str1 := "Hello World"
	fmt.Println("------ for i := range str1 ------")
	for i := range str1 {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, _ := range str1 ------")
	for i, _ := range str1 {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, v := range str1 ------")
	for i, c := range str1 {
		fmt.Printf("i->%v,%T | c->%q,%T\n", i, i, c, c)
	}

	// 非纯ASCII字符的字符串
	str2 := "你好中国，你好世界！"
	fmt.Println("------ for i := range str2 ------")
	for i := range str2 {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, _ := range str2 ------")
	for i, _ := range str2 {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, v := range str2 ------")
	for i, c := range str2 {
		fmt.Printf("i->%v,%T | c->%q,%T\n", i, i, c, c)
	}

	// 数组
	a := [...]int{1, 2, 3}

	fmt.Println("------ for i := range a ------")
	for i := range a {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, _ := range a ------")
	for i, _ := range a {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, v := range a ------")
	for i, v := range a {
		fmt.Printf("i->%v,%T | v->%v,%T\n", i, i, v, v)
	}

	// 数组指针
	ap := &a
	fmt.Println("------ for i := range ap ------")
	for i := range ap {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, _ := range ap ------")
	for i, _ := range ap {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, v := range ap ------")
	for i, v := range ap {
		fmt.Printf("i->%v,%T | v->%q,%T\n", i, i, v, v)
	}

	// 切片
	s := []int{1, 2, 3}
	fmt.Println("------ for i := range s ------")
	for i := range s {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, _ := range s ------")
	for i, _ := range s {
		fmt.Printf("i->%v,%T\n", i, i)
	}

	fmt.Println("------ for i, v := range s ------")
	for i, v := range s {
		fmt.Printf("i->%v,%T | v->%v,%T\n", i, i, v, v)
	}

	// 切片指针 [range中的变量不能是切片指针]
	//sp := &s
	//fmt.Println("------ for i := range sp ------")
	//for i := range sp { // 编译报错：cannot range over sp (variable of type *[]int
	//	fmt.Printf("i->%v,%T\n", i, i)
	//}
	//
	//fmt.Println("------ for i, _ := range sp ------")
	//for i, _ := range sp { // 编译报错：cannot range over sp (variable of type *[]int
	//	fmt.Printf("i->%v,%T\n", i, i)
	//}
	//
	//fmt.Println("------ for i, v := range sp ------")
	//for i, v := range sp { // 编译报错：cannot range over sp (variable of type *[]int
	//	fmt.Printf("i->%v,%T | v->%v,%T\n", i, i, v, v)
	//}

	// map
	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3

	fmt.Println("------ for k := range m ------")
	for k := range m {
		fmt.Printf("k->%v,%T\n", k, k)
	}

	fmt.Println("------ for k, _ := range m ------")
	for k, _ := range m {
		fmt.Printf("k->%v,%T\n", k, k)
	}

	fmt.Println("------ for k, v := range m ------")
	for k, v := range m {
		fmt.Printf("k->%v,%T | v->%v,%T\n", k, k, v, v)
	}

	// channel
	chs := make(chan int, 3)
	chs <- 1
	chs <- 2
	chs <- 3
	// chs <- 4 // 将触发错误：fatal error: all goroutines are asleep - deadlock!

	close(chs)
	fmt.Println("------ for ch := range chs ------")
	for ch := range chs {
		fmt.Printf("%v,%T\n", ch, ch)
	}
}
```

