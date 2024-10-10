package main

import "fmt"

func main() {
	a := int8(8)
	b := int8(11)
	c := int8(-8)
	d := int8(-11)
	fmt.Printf("a&^b\na=%08b a补=%08b\nb=%08b b补=%08b\nr=%08b r补=%08b\n----\n", a, uint8(a), b, uint8(b), a&^b, uint8(a&^b))
	fmt.Printf("c&^d\nc=%08b c补=%08b\nd=%08b d补=%08b\nr=%08b r补=%08b\n----\n", c, uint8(c), d, uint8(d), c&^d, uint8(c&^d))
	fmt.Printf("a&^d\na=%08b a补=%08b\nd=%08b c补=%08b\nr=%08b r补=%08b\n----\n", a, uint8(a), d, uint8(d), a&^d, uint8(a&^d))
	fmt.Printf("d&^a\nd=%08b d补=%08b\na=%08b a补=%08b\nr=%08b r补=%08b\n----\n", d, uint8(d), a, uint8(a), d&^a, uint8(d&^a))
	fmt.Printf("b&^c\nb=%08b b补=%08b\nc=%08b c补=%08b\nr=%08b r补=%08b\n----\n", b, uint8(b), c, uint8(c), b&^c, uint8(b&^c))
	fmt.Printf("c&^b\nc=%08b c补=%08b\nb=%08b b补=%08b\nr=%08b r补=%08b\n----\n", c, uint8(c), b, uint8(b), c&^b, uint8(c&^b))


	// e := uint(256255)
	// fmt.Printf("a&^e\na=%08b\ne=%08b\nr=%08b\n----\n", a, e, a&^e) //invalid operation: a &^ e (mismatched types int8 and uint)
	// f := int16(11)
	// fmt.Printf("a&^f\na=%08b\nf=%08b\nr=%08b\n----\n", a, f, a&^f) //invalid operation: a &^ f (mismatched types int8 and int16)
}
