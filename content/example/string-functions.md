+++
title = "string functions"
date = 2023-08-07T13:50:09+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# string functions

> 原文：https://gobyexample.com/string-functions

```go
package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
// Output:
//Contains:   true
//Count:      2
//HasPrefix:  true
//HasSuffix:  true
//Index:      1
//Join:       a-b
//Repeat:     aaaaa
//Replace:    f00
//Replace:    f0o
//Split:      [a b c d e]
//ToLower:    test
//ToUpper:    TEST
```

