+++
title = "map"
date = 2023-08-07T13:32:29+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# map

```go
package main

import "fmt"

type K interface{}

func main() {
	var m0 map[string]int
	fmt.Printf("%v,%T,len=%d\n", m0, m0, len(m0))  // map[],map[string]int,len=0
	fmt.Printf("%#v,%T,len=%d\n", m0, m0, len(m0)) // map[string]int(nil),map[string]int,len=0

	m01 := make(map[string]int)
	fmt.Printf("%v,%T,len=%d\n", m01, m01, len(m01))  // map[],map[string]int,len=0
	fmt.Printf("%#v,%T,len=%d\n", m01, m01, len(m01)) // map[string]int{},map[string]int,len=0

	//对 m01 添加或修改元素
	m01["A"] = 30
	m01["B"] = 20
	fmt.Printf("%#v,%T,len=%d\n", m01, m01, len(m01)) // map[string]int{"A":30, "B":20},map[string]int,len=2

	m1 := make(map[K]int)
	fmt.Printf("%v,%T,len=%d\n", m1, m1, len(m1))  // map[],map[main.K]int,len=0
	fmt.Printf("%#v,%T,len=%d\n", m1, m1, len(m1)) // map[main.K]int{},map[main.K]int,len=0

	//对 m1 添加或修改元素
	var k1 K
	k1 = 1

	m1[k1] = 1
	fmt.Printf("%#v,%T,len=%d\n", m1, m1, len(m1)) // map[main.K]int{1:1},map[main.K]int,len=1

	k1 = "a"
	m1[k1] = 2
	fmt.Printf("%#v,%T,len=%d\n", m1, m1, len(m1)) // map[main.K]int{1:1, "a":2},map[main.K]int,len=2

}

```

