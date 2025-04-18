package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入姓名和年龄（格式：姓名 年龄）：")
	_, err := fmt.Scanf("%s%d\n\n", &name, &age)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)

	// 清除缓冲区中的剩余内容（包括换行符）
	s, _ := reader.ReadString('\n')
	fmt.Print("s=", s)

	fmt.Print("请输入姓名和年龄（格式：姓名 年龄）：")
	_, err = fmt.Scanf("%s%d", &name, &age)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
}
