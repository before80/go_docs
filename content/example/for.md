+++
title = "for"
date = 2023-08-07T13:31:41+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# for

```go
package main

import "fmt"

func main() {
	i := 0
	// 没有语句
	for {
		i++
		break
	}

	// 1个语句
	j := 0
	for j <= 99 {
		j++
	}

	// 3个语句（都非空），1个变量
	for n := 0; n <= 99; n++ {
		// do something
	}

	// 3个语句（1个为空），1个变量
	m := 0
	// 相当于
	for ; m <= 99; m++ {
		// do something

	}

	// 3个语句（1个为空），1个变量
	for k := 0; k <= 99; {
		// do something
		k++
	}

	// 3个语句（2个为空） =》 自动变成1个语句，1个变量
	l := 0
	for l <= 99 {
		// do something
		l++
	}

	// 3个语句（都不为空），多个变量
	for p, q := 0, 0; p <= 20 && q <= 40; p, q = p+1, q+2 {
		// do something
	}

	// 3个语句（1个为空），多个变量
	for p, q := 0, 0; p <= 20 && q <= 40; {
		// do something
		p, q = p+1, q+2
	}

	// 3个语句（1个为空），多个变量
	p, q := 0, 0
	for ; p <= 20 && q <= 40; p, q = p+1, q+2 {
		// do something
	}

	// 3个语句（2个为空）=》 自动变成1个语句，多个变量
	p, q = 0, 0
	for p <= 20 && q <= 40 {
		// do something
		p, q = p+1, q+2
	}

	fmt.Println("Run Here")
}

```

