+++
title = "slices"
date = 2023-08-07T13:32:23+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# slices

```go
package main

import "fmt"

func main() {
	s1 := []bool{false, true, true}
	s2 := []uint{1, 2, 3}
	s3 := []uint8{1, 2, 3}
	s4 := []uint16{1, 2, 3}
	s5 := []uint32{1, 2, 3}
	s6 := []uint64{1, 2, 3}
	s7 := []int{1, 2, 3}
	s8 := []int8{1, 2, 3}
	s9 := []int16{1, 2, 3}
	s10 := []int32{1, 2, 3}
	s11 := []int64{1, 2, 3}
	s12 := []float32{1.2, 2.3, 3.4}
	s13 := []float64{1.2, 2.3, 3.4}
	s14 := []complex64{1.2, 2.3, 3.4}
	s15 := []complex128{1.2, 2.3, 3.4}

	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s2, s2, len(s2), cap(s2))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s3, s3, len(s3), cap(s3))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s4, s4, len(s4), cap(s4))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s5, s5, len(s5), cap(s5))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s6, s6, len(s6), cap(s6))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s7, s7, len(s7), cap(s7))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s8, s8, len(s8), cap(s8))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s9, s9, len(s9), cap(s9))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s10, s10, len(s10), cap(s10))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s11, s11, len(s11), cap(s11))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s12, s12, len(s12), cap(s12))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s13, s13, len(s13), cap(s13))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s14, s14, len(s14), cap(s14))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s15, s15, len(s15), cap(s15))

	// 从数组中获取切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}

	s16 := a1[:]
	s17 := a1[0:]
	s18 := a1[:len(a1)]
	s19 := a1[0:len(a1)]

	s20 := a1[1:]
	s21 := a1[1:len(a1)]

	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s16, s16, len(s16), cap(s16))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s17, s17, len(s17), cap(s17))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s18, s18, len(s18), cap(s18))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s19, s19, len(s19), cap(s19))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s20, s20, len(s20), cap(s20))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s21, s21, len(s21), cap(s21))

	// 从数组中获取切片
	sl := []int{1, 2, 3, 4, 5, 6}

	s22 := sl[:]
	s23 := sl[0:]
	s24 := sl[:len(sl)]
	s25 := sl[0:len(sl)]

	s26 := sl[1:]
	s27 := sl[1:len(sl)]

	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s22, s22, len(s22), cap(s22))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s23, s23, len(s23), cap(s23))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s24, s24, len(s24), cap(s24))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s25, s25, len(s25), cap(s25))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s26, s26, len(s26), cap(s26))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s27, s27, len(s27), cap(s27))

	// 使用make函数生成切片
	s28 := make([]int, 6)
	s29 := make([]int, 6, 6)
	s30 := make([]int, 6, 10)

	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s28, s28, len(s28), cap(s28))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s29, s29, len(s29), cap(s29))
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s30, s30, len(s30), cap(s30))

	// 使用 new 函数声明切片
	s31 := *new([]int) // 注意这里有一个 * 符号
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s31, s31, len(s31), cap(s31))

	s32 := []int{1, 2, 3}
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s32, s32, len(s32), cap(s32))
	// 修改
	s32[0] = 11
	s32[1] = 22
	s32[2] = 33
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s32, s32, len(s32), cap(s32))

	// append 操作
	// 仅追加一个元素
	s32 = append(s32, 4) // 发生扩容
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s32, s32, len(s32), cap(s32))

	// 追加多个元素
	s32 = append(s32, 5, 6, 7)
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s32, s32, len(s32), cap(s32))

	// 追加另外1个切片中的所有元素
	s320 := []int{8, 9, 10}
	s32 = append(s32, s320...)
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s32, s32, len(s32), cap(s32))

	// copy 操作
	s33 := []int{1, 2, 3}
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s33, s33, len(s33), cap(s33))
	copy(s33, s32)
	fmt.Printf("%#v,%T,len=%d,cap=%d\n", s33, s33, len(s33), cap(s33))

}

```

