package main

import (
	"fmt"
	"math"
)

func main() {
	z := int8(127)

	fmt.Printf("数     原码     反码     补码\n")
	for z > -128 || z == -128 {
		var f int8
		f = z
		if z < 0 {
			f = int8((^uint8(math.Abs(float64(z)))) | 0b10000000)
		}
		fmt.Printf("%3d | %#08b | %#08b | %#08b\n", z, z, f, uint8(z))
		if z == -128 {
			break
		}
		z--
	}
}
