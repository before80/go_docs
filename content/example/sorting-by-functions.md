+++
title = "sorting by functions"
date = 2023-08-07T13:49:29+08:00
weight = 37
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# sorting by functions

> 原文：https://gobyexample.com/sorting-by-functions
>
> 存在修改

```go
package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi", "pear", "apple"}
	sort.Sort(byLength(fruits))
	fmt.Printf("%v\n", fruits)

	sort.Stable(byLength(fruits))
	fmt.Printf("%v\n", fruits)

}

```

