+++
title = "http-client"
date = 2023-08-07T13:57:24+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# HTTP Client

> 原文：https://gobyexample.com/http-client

```go
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

```

