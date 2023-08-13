+++
title = "embed-directive"
date = 2023-08-07T13:55:49+08:00
weight = 60
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Embed Directive

> 原文：https://gobyexample.com/embed-directive

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"embed"
	"fmt"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	print(fileString)       // hello go
	print(string(fileByte)) // hello go

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1)) // 123

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2)) // 456

	content3, _ := folder.ReadFile("folder/single_file.txt")
	fmt.Println(content3)   // [104 101 108 108 111 32 103 111 13 10]
	print(string(content3)) // hello go
}

```

