package main

import "fmt"

func main() {
	a := int8(126)
	b := int8(127)
	c := int8(1)
	r1 := a + c
	r2 := b + c
	fmt.Printf("int8 %d + %d = %d\n", a, c, a+c)
	fmt.Printf("int8 %d + %d = %d\n", a, c, r1)
	fmt.Printf("int8 %d + %d = %d\n", b, c, b+c)
	fmt.Printf("int8 %d + %d = %d\n", b, c, r2)

	d := int8(-127)
	e := int8(-128)
	f := int8(-1)
	r3 := d + f
	r4 := e + f
	fmt.Printf("int8 %d + %d = %d\n", d, f, d+f)
	fmt.Printf("int8 %d + %d = %d\n", d, f, r3)
	fmt.Printf("int8 %d + %d = %d\n", e, f, e+f)
	fmt.Printf("int8 %d + %d = %d\n", e, f, r4)
}
