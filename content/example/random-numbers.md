+++
title = "random number"
date = 2023-08-07T13:51:50+08:00
weight = 49
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++



# random number

> 原文：https://gobyexample.com/random-numbers

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100)) // 9
	fmt.Println(rand.Intn(100)) // 96

	fmt.Println(rand.Float64())           // 0.20107204222647884
	fmt.Println((rand.Float64() * 5) + 5) // 8.852847642797128
	fmt.Println((rand.Float64() * 5) + 5) // 9.47001132931635

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100)) // 22
	fmt.Println(r1.Intn(100)) // 22

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100)) // 5
	fmt.Println(r2.Intn(100)) // 87

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Println(r3.Intn(100)) // 5
	fmt.Println(r3.Intn(100)) // 87
}

```

