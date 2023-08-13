+++
title = "number-parsing"
date = 2023-08-07T13:52:11+08:00
weight = 50
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# number parsing

> 原文：https://gobyexample.com/number-parsing

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234

	i0, _ := strconv.ParseInt("111", 0, 64)
	fmt.Println(i0) // 111

	i2, _ := strconv.ParseInt("111", 2, 64)
	fmt.Println(i2) // 7

	i16, _ := strconv.ParseInt("111", 16, 64)
	fmt.Println(i16) // 273

	i21, _ := strconv.ParseInt("111", 21, 64)
	fmt.Println(i21) // 463

	i36, _ := strconv.ParseInt("111", 36, 64)
	fmt.Println(i36) // 1333

	d0, _ := strconv.ParseInt("0x111", 0, 64)
	fmt.Println(d0) // 273

	d2, e := strconv.ParseInt("0x111", 2, 64)
	fmt.Println(d2) // 0
	fmt.Println(e)  // strconv.ParseInt: parsing "0x111": invalid syntax

	d12, e := strconv.ParseInt("0x111", 12, 64)
	fmt.Println(d12) // 0
	fmt.Println(e)   // strconv.ParseInt: parsing "0x111": invalid syntax

	d36, _ := strconv.ParseInt("0x111", 36, 64)
	fmt.Println(d36) // 1540981

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u) // 789

	k, _ := strconv.Atoi("135")
	fmt.Printf("%v,%T\n", k, k) // 135,int

	m, e := strconv.Atoi("wat")
	fmt.Println(m) // 0
	fmt.Println(e) // strconv.Atoi: parsing "wat": invalid syntax
}

```

