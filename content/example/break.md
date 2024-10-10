+++
title = "break"
weight = 7
date = 2023-08-26T11:06:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# break

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func breakFor() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			if j == 2 {
				break
			}
		}
	}
}

func breakForWithOuterLoopLabel() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			if j == 2 {
				break OuterLoop
			}
		}
	}
}

func breakForWithInnerLoopLabel() {
	for i := 0; i < 2; i++ {
	InnerLoop:
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			if j == 2 {
				break InnerLoop
			}
		}
	}
}

func breakSwitch() {
	i := rand.Intn(3)
	switch i {
	case 0:
		fmt.Println("i is 0.")
		break
	case 1, 2:
		fmt.Println("i is 1 or 2.")
		break
	default:
		fmt.Println("i is unknown.")
	}
}

func breakSwitchWithLabel() {
	i := rand.Intn(3)
Label:
	switch i {
	case 0:
		fmt.Println("i is 0.")
		break Label
	case 1, 2:
		fmt.Println("i is 1 or 2.")
		break Label
	default:
		fmt.Println("i is unknown.")
	}
}

func breakSelect() {
	sch1 := make(chan int, 3)
	sch2 := make(chan int, 3)

	go func() {
		sch1 <- 1
		sch1 <- 2
		sch1 <- 3
	}()

	go func() {
		sch2 <- 1
		sch2 <- 2
		sch2 <- 3
	}()

	num := 0
	for {
		select {
		case d := <-sch1:
			fmt.Println("from sch1 got ", d)
			break
		case d := <-sch2:
			fmt.Println("from sch2 got ", d)
			break
		default:
			fmt.Println("run in default")
			num++
			time.Sleep(time.Second)
			if num > 5 {
				return
			}
		}
	}
}

func breakSelectWithLabel() {
	sch1 := make(chan int, 3)
	sch2 := make(chan int, 3)

	go func() {
		sch1 <- 1
		sch1 <- 2
		sch1 <- 3
	}()

	go func() {
		sch2 <- 1
		sch2 <- 2
		sch2 <- 3
	}()

OuterLabel:
	for {
		select {
		case d := <-sch1:
			fmt.Println("from sch1 got ", d)
			break OuterLabel
		case d := <-sch2:
			fmt.Println("from sch2 got ", d)
			break OuterLabel
		default:
			fmt.Println("run in default")
		}
	}
}

func main() {
	fmt.Println("breakFor--------------")
	breakFor()

	fmt.Println("breakForWithOuterLoopLabel--------------")
	breakForWithOuterLoopLabel()

	fmt.Println("breakForWithInnerLoopLabel--------------")
	breakForWithInnerLoopLabel()

	fmt.Println("breakSwitch--------------")
	breakSwitch()

	fmt.Println("breakSwitchWithLabel--------------")
	breakSwitchWithLabel()

	fmt.Println("breakSelectWithLabel--------------")
	breakSelectWithLabel()

	fmt.Println("breakSelect--------------")
	breakSelect()
}

// Output:
//breakFor--------------
//0 0
//0 1
//0 2
//1 0
//1 1
//1 2
//breakForWithOuterLoopLabel--------------
//0 0
//0 1
//0 2
//breakForWithInnerLoopLabel--------------
//0 0
//0 1
//0 2
//1 0
//1 1
//1 2
//breakSwitch--------------
//i is 0.
//breakSwitchWithLabel--------------
//i is 0.
//breakSelectWithLabel--------------
//run in default
//from sch2 got  1
//breakSelect--------------
//run in default
//from sch2 got  1
//from sch1 got  1
//from sch1 got  2
//from sch1 got  3
//from sch2 got  2
//from sch2 got  3
//run in default
//run in default
//run in default
//run in default
//run in default
```

