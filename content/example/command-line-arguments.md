+++
title = "command-line-arguments"
date = 2023-08-07T13:56:28+08:00
weight = 62
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Command-Line Arguments

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Printf("%#v,%T\n", argsWithProg, argsWithProg)
	fmt.Printf("%v,%T\n", argsWithoutProg, argsWithoutProg)
	fmt.Printf("%v,%T\n", arg, arg)
}

```



```bash
PS D:\Dev\Go\byExample\command_line_arguments> .\main.exe a b c d
[]string{"D:\\Dev\\Go\\byExample\\command_line_arguments\\main.exe", "a", "b", "c", "d"},[]string
[a b c d],[]string
c,string
```

