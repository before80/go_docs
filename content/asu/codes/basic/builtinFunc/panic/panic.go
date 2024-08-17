package main

import "fmt"

func SayHello() {
	fmt.Println("Hello")
	defer fmt.Println("World")
	panic("Notice: This is a panic.")
}

func main() {
	SayHello()
}

//Hello
//World
//panic: Notice: This is a panic.
