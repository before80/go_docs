+++
title = "atomic counter"
date = 2023-08-07T13:41:08+08:00
weight = 33
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# atomic counter

> 原文：https://gobyexample.com/atomic-counters

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("counter:", counter)
}

```

