+++
title = "sha256-hashes"
date = 2023-08-07T13:52:41+08:00
weight = 52
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# sha256 hashes

> 原文：https://gobyexample.com/sha256-hashes

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)         // sha256 this string
	fmt.Printf("%x\n", bs) // 1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb
}

```

