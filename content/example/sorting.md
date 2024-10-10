+++
title = "sorting"
date = 2023-08-07T13:49:18+08:00
weight = 36
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# sorting

> 原文：https://gobyexample.com/sorting
>
> 存在修改

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	fmt.Println("Sorted:", sort.StringsAreSorted(strs))

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)
	fmt.Println("Sorted:", sort.IntsAreSorted(ints))

	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	// Sort by name, preserving original order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Printf("By name: %v\n", people)

	// Sort by age preserving name order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Printf("By age,name: %v\n", people)
}

```

