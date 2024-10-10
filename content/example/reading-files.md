+++
title = "reading-files"
date = 2023-08-07T13:53:25+08:00
weight = 54
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# reading files

> 原文：https://gobyexample.com/reading-files

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("./tmp/dat.txt")
	check(err)
	fmt.Print(string(dat))
	//hello
	//go

	f, err := os.Open("./tmp/dat.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)

	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1])) // 5 bytes: hello

	o2, err := f.Seek(6, 0)
	check(err)

	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)

	fmt.Printf("%d bytes @ %d: ", n2, o2) // 2 bytes @ 6:
	fmt.Printf("%v\n", string(b2[:n2]))   // g

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3)) // 2 bytes @ 6:
	// g

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4)) // 5 bytes: hello

	f.Close()
}

```

./tmp/dat.txt

```txt
hello
go
```

