+++
title = "Directories"
date = 2023-08-07T13:54:47+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Directories

> 原文：https://gobyexample.com/directories

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	err := os.Mkdir("subdir", 0755)
	check(err)

	//defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent") // Listing subdir/parent
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// child true
	// file2 false
	// file3 false

	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child") // Listing subdir/parent/child
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	// file4 false

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir") // Visiting subdir
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	//   subdir\parent\file2 false
	//   subdir\parent\file3 false
	return nil
}
```

