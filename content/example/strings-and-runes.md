+++
title = "Strings and Runes"
date = 2023-08-07T13:34:38+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Strings and Runes

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

const s1 = "Hi"
const s2 = "你好"

func main() {
	const s3 = "Hello"
	const s4 = "您好"

	var s5 = "Good"
	s6 := "好的呀"

	fmt.Printf("%q,%T,len=%d\n", s1, s1, len(s1))
	fmt.Printf("%q,%T,len=%d\n", s2, s2, len(s2))
	fmt.Printf("%q,%T,len=%d\n", s3, s3, len(s3))
	fmt.Printf("%q,%T,len=%d\n", s4, s4, len(s4))
	fmt.Printf("%q,%T,len=%d\n", s5, s5, len(s5))
	fmt.Printf("%q,%T,len=%d\n", s6, s6, len(s6))

	// for -> 取出索引对应 字节的值
	fmt.Println("----------------for s5--------------------")
	for i := 0; i < len(s5); i++ {
		fmt.Printf("i=%d,%T,%v,%x,%X,%c,%q\n", i, s5[i], s5[i], s5[i], s5[i], s5[i], s5[i])
	}

	// for -> 取出索引对应 字节的值 <- 出现乱码
	fmt.Println("----------------for s6--------------------")
	for i := 0; i < len(s6); i++ {
		fmt.Printf("i=%d,%T,%v,%x,%X,%c,%q\n", i, s6[i], s6[i], s6[i], s6[i], s6[i], s6[i])
	}

	// for range
	fmt.Println("----------------for range s5--------------------")

	for i, runeV := range s5 {
		fmt.Printf("i=%d,%T,%v,%x,%X,%c,%q\n", i, runeV, runeV, runeV, runeV, runeV, runeV)
	}

	fmt.Println("----------------for range s6--------------------")

	for i, runeV := range s6 {
		fmt.Printf("i=%d,%T,%v,%x,%X,%c,%q\n", i, runeV, runeV, runeV, runeV, runeV, runeV)
	}

	// 使用 utf8.DecodeRuneInString 函数
	fmt.Println("----------------for s5--------------------")
	for len(s5) > 0 {
		r, size := utf8.DecodeRuneInString(s5)
		fmt.Printf("%T,%v,%x,%X,%c,%q,size=%d\n", r, r, r, r, r, r, size)
		s5 = s5[size:]
	}

	fmt.Println("----------------for s6--------------------")
	for len(s6) > 0 {
		r, size := utf8.DecodeRuneInString(s6)
		fmt.Printf("%T,%v,%x,%X,%c,%q,size=%d\n", r, r, r, r, r, r, size)
		s6 = s6[size:]
	}
}

// Output:
//"Hi",string,len=2
//"你好",string,len=6
//"Hello",string,len=5
//"您好",string,len=6
//"Good",string,len=4
//"好的呀",string,len=9
//----------------for s5--------------------
//i=0,uint8,71,47,47,G,'G'
//i=1,uint8,111,6f,6F,o,'o'
//i=2,uint8,111,6f,6F,o,'o'
//i=3,uint8,100,64,64,d,'d'
//----------------for s6--------------------
//i=0,uint8,229,e5,E5,å,'å'
//i=1,uint8,165,a5,A5,¥,'¥'
//i=2,uint8,189,bd,BD,½,'½'
//i=3,uint8,231,e7,E7,ç,'ç'
//i=4,uint8,154,9a,9A,,'\u009a'
//i=5,uint8,132,84,84,,'\u0084'
//i=6,uint8,229,e5,E5,å,'å'
//i=7,uint8,145,91,91,,'\u0091'
//i=8,uint8,128,80,80,,'\u0080'
//----------------for range s5--------------------
//i=0,int32,71,47,47,G,'G'
//i=1,int32,111,6f,6F,o,'o'
//i=2,int32,111,6f,6F,o,'o'
//i=3,int32,100,64,64,d,'d'
//----------------for range s6--------------------
//i=0,int32,22909,597d,597D,好,'好'
//i=3,int32,30340,7684,7684,的,'的'
//i=6,int32,21568,5440,5440,呀,'呀'
//----------------for s5--------------------
//int32,71,47,47,G,'G',size=1
//int32,111,6f,6F,o,'o',size=1
//int32,111,6f,6F,o,'o',size=1
//int32,100,64,64,d,'d',size=1
//----------------for s6--------------------
//int32,22909,597d,597D,好,'好',size=3
//int32,30340,7684,7684,的,'的',size=3
//int32,21568,5440,5440,呀,'呀',size=3

```

