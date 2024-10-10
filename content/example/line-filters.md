+++
title = "line-filters"
date = 2023-08-07T13:54:05+08:00
weight = 56
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# line filters

> 原文：https://gobyexample.com/line-filters

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

```

