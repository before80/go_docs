+++
title = "通道"
weight = 93
date = 2023-06-12T16:06:37+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

#通道

## 会触发panic的通道操作

### 关闭值为nil的通道

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic -> ", r) // panic ->  close of nil channel
		}
	}()

	var ch chan int
	close(ch)
}

```



### 关闭已被关闭的通道

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic -> ", r) // panic ->  close of closed channel
		}
	}()

	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	d := <-ch
	fmt.Println(d) // 1
	close(ch)
	// 关闭已关闭的ch
	close(ch)
}
```



### 向已关闭的通道写入数据

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic -> ", r) // panic ->  send on closed channel
		}
	}()

	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	d := <-ch
	fmt.Println(d) // 1
	close(ch)
	// 向已关闭的通道写入数据
	ch <- 2
}

```



## 注意点：for ... range 带缓冲区通道

只能有一个迭代变量

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic -> ", r)
		}
	}()

	ch := make(chan int, 10)

	go func(c1 chan<- int) {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}(ch)

	go func(c2 <-chan int) {
		// for d, ok := range c2 { // <- 这里的 ok 不能出现！
        // 否则，编译的时候就报错：range over c2 (variable of type <-chan int) permits only one iteration variable
		for d := range c2 {
			fmt.Println(d)
		}
	}(ch)
	
	time.Sleep(2 * time.Second)
}

```



## 注意点：for ... range 无缓冲区通道

只能有一个迭代变量

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic -> ", r)
		}
	}()

	ch := make(chan int)

	go func(c1 chan<- int) {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}(ch)

	go func(c2 <-chan int) {
		// for d, ok := range c2 { // <- 这里的 ok 不能出现！
        // 否则，编译的时候就报错：range over c2 (variable of type <-chan int) permits only one iteration variable
		for d := range c2 {
			fmt.Println(d)
		}
	}(ch)

	time.Sleep(2 * time.Second)
}

```

