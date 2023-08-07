+++
title = "Variables"
date = 2023-08-07T13:31:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Variables

```go
package main

import "fmt"

func main() {
	// 布尔类型
	var b10 bool = false
	b11 := true
	_, _ = b10, b11

	// 数值
	// 整型类型
	var i10 uint8 = uint8(1)
	i11 := uint8(1)
	_, _ = i10, i11

	var i20 uint16 = uint16(2)
	i21 := uint16(2)
	_, _ = i20, i21

	var i30 uint32 = uint32(3)
	i31 := uint32(3)
	_, _ = i30, i31

	var i40 uint64 = uint64(4)
	i41 := uint64(4)
	_, _ = i40, i41

	var i50 uint = uint(5)
	i51 := uint(5)
	_, _ = i50, i51

	var i60 int8 = int8(6)
	i61 := uint8(6)
	_, _ = i60, i61

	var i70 int16 = int16(7)
	i71 := int16(2)
	_, _ = i70, i71

	var i80 int32 = int32(8)
	i81 := int32(8)
	_, _ = i80, i81

	var i90 int64 = int64(9)
	i91 := int64(9)
	_, _ = i90, i91

	var i100 int = int(10)
	i101 := int(10)
	_, _ = i100, i101

	// 浮点类型
	var f10 float32 = float32(1.2)
	f11 := float32(1.2)
	_, _ = f10, f11

	var f20 float64 = 1.2
	f21 := 1.2
	_, _ = f20, f21

	// 复数类型
	var cmpx10 complex64 = complex(float32(1.2), float32(3.4))
	cmpx11 := complex(float32(1.2), float32(3.4))
	_, _ = cmpx10, cmpx11

	var cmpx20 complex128 = complex(1.2, 3.4)
	cmpx21 := complex(1.2, 3.4)
	cmpx22 := 1.2 + 3.4i
	_, _, _ = cmpx20, cmpx21, cmpx22

	// byte类型 - uint8的别名
	var bt10 byte = byte('A')
	bt11 := byte('A')
	_, _ = bt10, bt11

	// rune类型 - int32的别名
	var r10 rune = 'A'
	r11 := 'A'
	_, _ = r10, r11

	// string类型
	var s10 string = "Hi"
	s11 := "Hi"
	_, _ = s10, s11

	// array类型
	var a10 [10]int = [10]int{}
	a11 := [10]int{}
	_, _ = a10, a11

	// slice类型
	var sl10 []int = []int{1, 2, 3}
	sl11 := []int{1, 2, 3}
	_, _ = sl10, sl11

	// map类型
	var m10 map[int]string = map[int]string{18: "zLongX-1.0", 30: "zLongX-2.0"}
	m11 := map[int]string{18: "zLongX-1.0", 30: "zLongX-2.0"}
	_, _ = m10, m11

	// channel 类型
	var ch10 chan int = make(chan int)
	ch11 := make(chan int)

	go func() {
		ch10 <- 1
		ch11 <- 1
	}()

	<-ch10
	<-ch11

	close(ch10)
	close(ch11)

	var ch20 chan int = make(chan int, 1)
	ch21 := make(chan int, 1)

	go func() {
		ch20 <- 1
		ch21 <- 1
	}()

	<-ch20
	<-ch21
	close(ch20)
	close(ch21)

	// struct类型
	var st10 struct{ Name string } = struct{ Name string }{"zLongX"}
	st11 := struct{ Name string }{"zLongX"}
	_, _ = st10, st11

	// interface类型
	var it10 interface{}
	_ = it10

	var it20 any
	_ = it20

	fmt.Println("Run Here")
}

```

