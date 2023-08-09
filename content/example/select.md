+++
title = "select"
date = 2023-08-07T13:37:10+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# select

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	waitSeconds := 2

	go func() {
		t := time.After(time.Duration(waitSeconds) * time.Second)
		for {
			select {
			case ch1 <- rand.Intn(101):
				time.Sleep(time.Second)
			case <-t:
				fmt.Println("1 time out")
				close(ch1)
				return
			}
		}
	}()

	go func() {
		t := time.After(time.Duration(waitSeconds) * time.Second)
		for {
			select {
			case ch2 <- rand.Intn(101):
				time.Sleep(time.Second)
			case <-t:
				fmt.Println("2 time out")
				close(ch2)
				return
			}
		}
	}()

	ch1Closed := false
	ch2Closed := false
    
	for {
		select {
		case d1, ok := <-ch1:
			if !ok {
				ch1Closed = true
			} else {
				fmt.Println("received from ch1:", d1)
			}
		case d2, ok := <-ch2:
			if !ok {
				ch2Closed = true
			} else {
				fmt.Println("received from ch2:", d2)
			}
		}

		if ch1Closed && ch2Closed {
			break
		}
	}
}

```

