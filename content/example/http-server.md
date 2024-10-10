+++
title = "HTTP Server"
date = 2023-08-07T13:57:30+08:00
weight = 67
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# HTTP Server

> 原文：https://gobyexample.com/http-server

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

```



```bash
PS D:\Dev\Go\byExample> curl localhost:8090/hello
hello
PS D:\Dev\Go\byExample> curl localhost:8090/headers
User-Agent: curl/8.0.1
Accept: */*
PS D:\Dev\Go\byExample> 
```

