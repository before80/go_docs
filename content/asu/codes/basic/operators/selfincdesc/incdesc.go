package main

import "fmt"

func main() {
	a := 5                 // 二进制表示为: 0000 0101
	b := ^a                // 按位取反后: 1111 1010 (对于8位无符号整数，这是 250)
	fmt.Printf("%0b\n", b) // 输出: -6
	fmt.Println(b)

}
