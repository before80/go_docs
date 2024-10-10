package main

import "fmt"

func main() {
	a := int8(8)
	c := int8(-8)
	fmt.Printf("^a\na=%08b a补=%08b <- %d\nr=%08b r补=%08b <- %d\n----\n", a, uint8(a), a, ^a, uint8(^a), ^a)
	fmt.Printf("^c\na=%08b c补=%08b <- %d\nr=%08b r补=%08b <- %d\n----\n", c, uint8(c), c, ^c, uint8(^c), ^c)

}
