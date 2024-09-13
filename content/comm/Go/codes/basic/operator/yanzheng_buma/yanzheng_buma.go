package main

import (
	"fmt"
	"math"
)

func main() {
	a := int(-1)
	for a >= -127 {
		b := uint8(a)
		c := uint8(math.Abs(float64(a)))
		fmt.Printf("b=%d,c=%d,b+c=%d\n", b, c, uint16(b)+uint16(c))
		if uint16(b)+uint16(c) == 256 {
			fmt.Println(a)
		}
		a--
	}
}
