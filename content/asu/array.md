+++
title = "数组"
date = 2023-07-31T20:55:01+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# 数组

## 可以对数组进行cap()操作？

```go
package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(len(arr)) // 6
	fmt.Println(cap(arr)) // 6
}

```

可见，可以对数组进行cap()操作，且数组的 容量和长度 是一致的！

## 数组的长度是固定的，那是否可以作为常量的值？

```go
package main

import "fmt"

var arr = [3]int{1, 2, 3}

const LEN = len(arr)

func main() {
	fmt.Println(LEN) // 3
}
```

可见，数组的长度是可以作为常量的值的！

## 对数组指针进行for range遍历

```go
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	p := &a

	// 遍历数组指针
	for i, v := range p {
		fmt.Println(i, v)
	}

	// 遍历数组
	for i, v := range *p {
		fmt.Println(i, v)
	}
}

Output:
0 1
1 2
2 3
0 1
1 2
2 3
```

## 使用new函数初始化数组

```go
package main

import "fmt"

func main() {
	a := new([10]int)
	fmt.Println(a) // &[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(*a) // [0 0 0 0 0 0 0 0 0 0]
    //b := make([10]int) // 编译报错：invalid argument: cannot make [10]int; type must be slice, map, or channel
}

```

可见，可以使用new函数来初始化数组，但make函数却不能用来初始化数组！
