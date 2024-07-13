+++
title = "constants"
date = 2023-08-07T13:31:30+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# constants

```go
package main

import "fmt"

// 无类型常量
const B10 = false
const I10 = 1
const F10 = 1.2

// 有类型常量
const B20 bool = true
const I20 int8 = 1
const F20 float32 = 1.2

// 使用iota
const (
	AI = iota + 1
	BI
	CI
	_
	DI
)

func main() {
	fmt.Printf("%v,%T\n", B10, B10)
	fmt.Printf("%v,%T\n", I10, I10)
	fmt.Printf("%v,%T\n", I10, I10)

	fmt.Printf("%v,%T\n", B20, B20)
	fmt.Printf("%v,%T\n", I20, I20)
	fmt.Printf("%v,%T\n", F20, F20)

	fmt.Printf("%v,%T\n", AI, AI)
	fmt.Printf("%v,%T\n", BI, BI)
	fmt.Printf("%v,%T\n", CI, CI)
	fmt.Printf("%v,%T\n", DI, DI)

	// 无类型常量
	const b10 = false
	const i10 = 1
	const f10 = 1.2

	// 有类型常量
	const b20 bool = true
	const i20 int8 = 1
	const f20 float32 = 1.2

	// （1）使用iota
	const (
		ai = iota + 1
		bi
		ci
		_
		di
	)

	// （2）使用iota
	const (
		t1 = iota
		t2
		t3
		_
		t4 = "abcde"
		t5
		t6 = iota
		t7
	)

	fmt.Printf("%v,%T\n", b10, b10)
	fmt.Printf("%v,%T\n", i10, i10)
	fmt.Printf("%v,%T\n", f10, f10)

	fmt.Printf("%v,%T\n", b20, b20)
	fmt.Printf("%v,%T\n", i20, i20)
	fmt.Printf("%v,%T\n", f20, f20)

	fmt.Printf("%v,%T\n", ai, ai)
	fmt.Printf("%v,%T\n", bi, bi)
	fmt.Printf("%v,%T\n", ci, ci)
	fmt.Printf("%v,%T\n", di, di)

	fmt.Printf("%v,%T\n", t1, t1)
	fmt.Printf("%v,%T\n", t2, t2)
	fmt.Printf("%v,%T\n", t3, t3)
	fmt.Printf("%v,%T\n", t4, t4)
	fmt.Printf("%v,%T\n", t5, t5)
	fmt.Printf("%v,%T\n", t6, t6)
	fmt.Printf("%v,%T\n", t7, t7)
}

// Output:
//false,bool
//1,int
//1,int
//true,bool
//1,int8
//1.2,float32
//1,int
//2,int
//3,int
//5,int
//false,bool
//1,int
//1.2,float64
//true,bool
//1,int8
//1.2,float32
//1,int
//2,int
//3,int
//5,int
//0,int
//1,int
//2,int
//abcde,string
//abcde,string
//6,int
//7,int
```

