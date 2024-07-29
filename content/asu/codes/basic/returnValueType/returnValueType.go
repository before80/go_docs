package main

import "fmt"

type St struct {
	x, y int
}

func returnSt() St {
	st := St{1, 2}
	fmt.Printf("1 st's address is %p\n", &st)
	return st
}

func returnArr() [3]int {
	arr := [3]int{}
	fmt.Printf("1 arr's address is %p\n", &arr)
	return arr
}

func main() {
	st := returnSt()
	fmt.Printf("2 st's address is %p\n", &st)

	arr := returnArr()
	fmt.Printf("2 arr's address is %p\n", &arr)
}
