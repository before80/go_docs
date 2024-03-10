+++
title = "signal"
date = 2023-08-07T13:59:31+08:00
weight = 71
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++



# signals

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

```



```bash
PS D:\Dev\Go\byExample\signals> go run main.go
awaiting signal

interrupt <- CTRL + C
exiting
```

