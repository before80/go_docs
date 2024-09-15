package main

import (
	"fmt"
	"time"
)

// 打印原码、反码、补码
func printCode(n int8) {
	fmt.Printf("十进制: %4d  ", n)

	// 原码
	if n >= 0 {
		// 正数的原码、反码、补码相同
		fmt.Printf("原码: %08b  反码: %08b  补码: %08b\n", n, n, n)
	} else {
		// 负数
		original := n & 0x7F                 // 原码，保留符号位外的部分
		onesComplement := ^original          // 反码
		twosComplement := onesComplement + 1 // 补码

		// 打印原码、反码、补码
		fmt.Printf("原码: 1%07b  反码: 1%07b  补码: %08b\n", original, onesComplement, twosComplement)
	}
}

func main() {
	//num := int8(-8)
	//negativeNum := ^num
	//fmt.Printf("原数值: %b, 取反码: %b\n", num, negativeNum)
	//// 从 127 到 -128 打印原码、反码、补码
	//for i := int8(127); i > -128 || i == -128; i-- {
	//	if i == -128 {
	//		break
	//	}
	//	printCode(i)
	//}
	t := time.Now()
	fmt.Println(t)
	time.Sleep(10 * time.Millisecond)
	t1 := time.Now()
	fmt.Println(t1.Sub(t))
	fmt.Println(340049016 / 3600 / 24 / 365)

	t2, _ := time.Parse("2006-01-02 15:04:05", "2024-09-01 12:00:00")
	fmt.Println(t2)
}
