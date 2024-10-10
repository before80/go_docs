+++
title = "Regular Expressions"
date = 2023-08-07T13:50:46+08:00
weight = 43
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Regular Expressions

> 原文：https://gobyexample.com/regular-expressions

```go
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	fmt.Println(match) // true

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach")) // true

	fmt.Println(r.FindString("peach punch")) // peach

	fmt.Println("idx:", r.FindStringIndex("peach punch")) // idx: [0 5]

	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1)) // all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	fmt.Println(r.Match([]byte("peach"))) // true

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r) // regexp: p([a-z]+)ch

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH
}

```

