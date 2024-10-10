+++
title = "closure"
date = 2023-08-07T13:33:53+08:00
weight = 14
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# closure - 闭包

```go
package main

import "fmt"

// 步骤
func step() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 累加器
func accumulator(initial int) func(int) int {
	sum := initial
	return func(increment int) int {
		sum += increment
		return sum
	}
}

// 打印消息（含给出第几次的消息）
func PrintMessage() func(string) {
	count := 0
	return func(message string) {
		count++
		fmt.Printf("Message #%d: %s\n", count, message)
	}
}

func main() {
	nextStep := step()

	fmt.Println(nextStep()) // 1
	fmt.Println(nextStep()) // 2
	fmt.Println(nextStep()) // 3

	a := accumulator(10)
	fmt.Println(a(5))  // 15
	fmt.Println(a(8))  // 23
	fmt.Println(a(-3)) // 20

	pm := PrintMessage()
	pm("Hello") // Message #1: Hello
	pm("World") // Message #2: World
}

```

