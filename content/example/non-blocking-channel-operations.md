+++
title = "non-blocking-channel-operations"
date = 2023-08-07T13:37:54+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Non-Blocking Channel Operations

```go
package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	chErr := make(chan error)
	signal := make(chan bool)

	select {
	case err := <-chErr:
		fmt.Println("found error:", err)
	default:
		fmt.Println("no found error")
	}

	go func() {
		time.Sleep(2 * time.Second)
		signal <- true
		close(signal)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chErr <- fmt.Errorf("%w", errors.New("can't continue"))
		close(chErr)
	}()

	select {
	case err := <-chErr:
		fmt.Println("found error:", err)
		break
	case sig := <-signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

```

