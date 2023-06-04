+++
title = "泛型"
linkTitle = "泛型"
weight = 4
date = 2023-05-17T12:10:24+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Generics

##  Type parameters 类型参数

[https://go.dev/tour/generics/1](https://go.dev/tour/generics/1)

​	可以使用类型参数编写 Go 函数，以便对多种类型进行操作。一个函数的类型参数出现在括号中，在函数的参数之前。

```
func Index[T comparable](s []T, x T) int
```

​	这个声明意味着`s`是任何类型T的一个切片，满足内置的`comparable`约束。x也是同一类型的值。

​	`comparable`是一个有用的约束，它使我们可以在该类型的值上使用`==`和`!=`运算符。在这个例子中，我们用它来比较一个值和所有的切片元素，直到找到一个匹配。这个`Index`函数适用于任何支持比较的类型。


```go 
package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

```


##  Generic types 范型

[https://go.dev/tour/generics/2](https://go.dev/tour/generics/2)

​	除了泛型函数，Go 还支持泛型。一个类型可以用一个类型参数进行参数化，这对实现通用数据结构很有用。

​	这个例子演示了一个简单的类型声明，用于保存任何类型值的单链式列表。

作为练习，为这个列表的实现添加一些功能。

```go title="main.go" 
package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}





func main() {
}

```

##  Congratulations!

[https://go.dev/tour/generics/3](https://go.dev/tour/generics/3)

您完成了这一课!

您可以回到`模块`列表中寻找下一步要学习的内容，或者继续学习下一课。