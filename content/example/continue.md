+++
title = "continue"
weight = 7
date = 2023-08-26T11:03:47+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# continue

```go
package main

import "fmt"

func continue1() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
}

func continue2() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 3:
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}
}

func continue3() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				fmt.Println(i, j)
				continue
			}
		}
	}
}

func continue4() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 3:
				fmt.Println(i, j)
				continue
			}
		}
	}
}

func main() {
	fmt.Println("continue 语句 仅用于for语句中")
	fmt.Println("continue + 标签的情况 --------------")
	fmt.Println("continue1() --------------")
	continue1()
	fmt.Println("continue2() --------------")
	continue2()

	fmt.Println("continue 没有标签的情况 --------------")
	fmt.Println("continue3() --------------")
	continue3()

	fmt.Println("continue4() --------------")
	continue4()
}

// Output:
//continue 语句 仅用于for语句中
//continue + 标签的情况 --------------
//continue1() --------------
//0 3
//1 3
//continue2() --------------
//0 3
//1 3
//continue 没有标签的情况 --------------
//continue3() --------------
//0 3
//1 3
//continue4() --------------
//0 3
//1 3 
```

