+++
title = "panic 和 recover"
date = 2023-08-07T13:49:38+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# panic 和 recover

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v,%T\n", r, r)
		}
	}()
	
	panic("a problem")
}
// Output:
// a problem,string
```

