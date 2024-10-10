+++
title = "Command-Line Subcommands"
date = 2023-08-07T13:56:50+08:00
weight = 64
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Command-Line Subcommands

> 原文：https://gobyexample.com/command-line-subcommands

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

```



```bash
PS D:\Dev\Go\byExample\command_line_subcommands> ./main.exe foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]
  
PS D:\Dev\Go\byExample\command_line_subcommands> ./main.exe bar -level 8 a1          
subcommand 'bar'
  level: 8
  tail: [a1]
  
PS D:\Dev\Go\byExample\command_line_subcommands> ./main.exe bar -enable a1 
flag provided but not defined: -enable
Usage of bar:
  -level int
        level

```

