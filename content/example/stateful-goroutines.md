+++
title = "Stateful Goroutines"
date = 2023-08-07T13:49:05+08:00
weight = 35
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Stateful Goroutines

> 原文：https://gobyexample.com/stateful-goroutines

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

var shortDuration = 100 * time.Millisecond

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	go func(ctx context.Context) {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			case d := <-ctx.Done():
				_ = d
				//fmt.Printf("state done. d=%#v,%T\n", d, d)
				return
			}
		}
	}(ctx)

	for r := 0; r < 100; r++ {
		go func(ctx context.Context) {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				//fmt.Println("<-read.resp = ", <-read.resp)
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
				select {
				case d := <-ctx.Done():
					_ = d
					//fmt.Printf("read done. d=%#v,%T\n", d, d)
					return
				default:
				}
			}
		}(ctx)
	}

	for w := 0; w < 10; w++ {
		go func(ctx context.Context) {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				//fmt.Println("<-write.resp = ", <-write.resp)
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
				select {
				case d := <-ctx.Done():
					_ = d
					//fmt.Printf("write done. d=%#v,%T\n", d, d)
					return
				default:
				}
			}
		}(ctx)
	}

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	}

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// 在 Go 语言中，如果 main 协程（也就是主协程）退出，
	// 其他所有正在运行的协程（包括子协程）也会随之被终止。
	// 这是因为 Go 运行时会在 main 协程退出时自动结束整个程序。
	
	// 虽然如此，本示例还是在每个goroutine 中都加入了退出条件代码
}

```

