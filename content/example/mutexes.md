+++
title = "Mutex"
date = 2023-08-07T13:48:45+08:00
weight = 34
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Mutex

```go
package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3) // 必须在使用 go 关键字之前，使用wg.Add(x)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()

	fmt.Println(c.counters)
}

```

