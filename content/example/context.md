+++
title = "context"
date = 2023-08-13T12:57:27+08:00
weight = 68
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# context

> 原文：https://gobyexample.com/context

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

```



```bash
PS D:\Dev\Go\byExample\context> go run main.go
server: hello handler started
server: hello handler ended
server: hello handler started
server: hello handler ended
server: hello handler started
server: context canceled <- 按了 CTRL +C
server: hello handler ended
server: hello handler started
server: context canceled <- 按了 CTRL +C
server: hello handler ended
server: hello handler started
server: hello handler ended
server: hello handler started
server: hello handler ended
server: hello handler started
server: context canceled <- 按了 CTRL +C
server: hello handler ended
```

