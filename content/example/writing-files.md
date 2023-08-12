+++
title = "writing-files"
date = 2023-08-07T13:53:44+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# writing files

> 原文：https://gobyexample.com/writing-files

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("./tmp/dat1.txt", d1, 0644)
	check(err)

	f, err := os.Create("./tmp/dat2.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2) // wrote 5 bytes

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3) // wrote 7 bytes

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4) // wrote 9 bytes

	w.Flush()

}

```

