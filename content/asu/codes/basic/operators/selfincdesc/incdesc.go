package main

import "fmt"

func main() {
	type A int
	type B int
	var a A
	a = 2
	fmt.Println(a)
	var b any
	b = 2
	bv := b.(B)
	fmt.Println(bv)

}
