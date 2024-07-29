package main

import "fmt"

const a = 1
const b = 2
const c = 3
const d = 9.876543210123456789

func main() {
	fmt.Printf("a=%[1]d (%[1]T)\n", a)
	fmt.Printf("a=%[1]d (%[1]T); b=%[2]d (%[2]T)\n", a, b)
	fmt.Printf("a=%[1]d (%[1]T); b=%d (%T)\n", a, b, b)
	fmt.Printf("a=%d (%[1]T); b=%[3]d (%[3]T); c=%d (%T)\n", a, a, b, c, c)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 2, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 3, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 4, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 5, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 6, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 7, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 8, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 9, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 10, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 11, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 12, 9)
	fmt.Printf("d=%[3]*.[2]*[1]f\n", d, 13, 9)
}
