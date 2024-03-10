+++
title = "command-line-flags"
date = 2023-08-07T13:56:36+08:00
weight = 63
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# command-line flags



```go
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
PS D:\Dev\Go\byExample\command_line_flags> .\main.exe -word=opt -numb=7 -fork -svar=flag 
word: opt
numb: 7
fork: true
svar: flag
tail: []
PS D:\Dev\Go\byExample\command_line_flags> .\main.exe -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []
PS D:\Dev\Go\byExample\command_line_flags> .\main.exe -word=opt a1 a2 a3                
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3]
PS D:\Dev\Go\byExample\command_line_flags> .\main.exe -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]
PS D:\Dev\Go\byExample\command_line_flags> .\main.exe -h                        
Usage of D:\Dev\Go\byExample\command_line_flags\main.exe:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")
PS D:\Dev\Go\byExample\command_line_flags> .\main.exe -wat                      
flag provided but not defined: -wat
Usage of D:\Dev\Go\byExample\command_line_flags\main.exe:
  -fork
        a bool
  -numb int
        an int (default 42)
  -svar string
        a string var (default "bar")
  -word string
        a string (default "foo")

```

