package main

import (
	"fmt"
)

func main() {
	//var a int8 = -8
	//// 获取 a 的绝对值部分
	//absValue := a & 0x7F // 取出除了符号位之外的部分
	//// 对绝对值部分取反
	//inverted := ^absValue
	//// 添加符号位（符号位保持为1）
	//result := inverted | int8(0x80)
	//
	//fmt.Printf("a 的二进制: %08b\n", a)
	//fmt.Printf("a 的反码: %08b\n", result)

	var a int8 = -8
	// 取出除了符号位之外的部分
	absValue := a & 0x7F // 保留符号位，取出绝对值部分
	// 对绝对值部分取反
	inverted := ^absValue
	// 符号位保持不变（负数的符号位为1）
	result := inverted | (a & 0x80) // 保持符号位

	fmt.Printf("a 的二进制: %08b\n", a)
	fmt.Printf("a 的反码: %08b\n", result)

	z := int8(127)

	fmt.Printf("数        原码         反码         补码\n")
	for z > -128 || z == -128 {
		var f int8
		f = z
		//if z < 0 {
		//	f = int8(^z | 0b10000000)
		//}
		fmt.Printf("%3d | %#08b | %#08b | %#08b\n", z, z, f, uint8(z))
		if z == -128 {
			break
		}
		z--
	}
}
