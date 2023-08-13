+++
title = "exit"
date = 2023-08-07T13:59:40+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++



# exit

原文：

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}

```

