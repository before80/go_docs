package main

import "fmt"

func main() {
	a := int8(8)
	b := uint8(8)
	c := int8(-8)
	fmt.Printf("a=%08b\n", a)

	fmt.Printf("a << 1 = %08b -> %d\n", a<<1, a<<1)
	fmt.Printf("a << 2 = %08b -> %d\n", a<<2, a<<2)
	fmt.Printf("a << 4 = %08b -> %d\n", a<<4, a<<3)
	fmt.Printf("a << 5 = %08b -> %d\n", a<<5, a<<4)
	fmt.Printf("a << 6 = %08b -> %d\n", a<<6, a<<6)
	fmt.Printf("b=%08b\n", b)
	fmt.Printf("b << 1 = %08b -> %d\n", b<<1, b<<1)
	fmt.Printf("b << 2 = %08b -> %d\n", b<<2, b<<2)
	fmt.Printf("b << 4 = %08b -> %d\n", b<<4, b<<4)
	fmt.Printf("b << 5 = %08b -> %d\n", b<<5, b<<5)
	fmt.Printf("b << 6 = %08b -> %d\n", b<<6, b<<6)
	fmt.Printf("c=%09b\n", c)
	fmt.Printf("c << 1 = %08b -> %d\n", c<<1, c<<1)
	fmt.Printf("c << 2 = %08b -> %d\n", c<<2, c<<2)
	fmt.Printf("c << 4 = %08b -> %d\n", c<<4, c<<4)
	fmt.Printf("c << 5 = %08b -> %d\n", c<<5, c<<5)
	fmt.Printf("c << 6 = %08b -> %d\n", c<<6, c<<6)

}
