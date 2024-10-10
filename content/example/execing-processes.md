+++
title = "execing-processes"
date = 2023-08-07T13:59:15+08:00
weight = 70
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Exec'ing processes

> 原文：https://gobyexample.com/execing-processes

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

```

