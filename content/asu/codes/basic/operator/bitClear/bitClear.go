package main

import "fmt"

func main() {

	a := 0b1110
	b := 0b0101
	c := a &^ b
	fmt.Printf("%b\n", c) // 1010
}
