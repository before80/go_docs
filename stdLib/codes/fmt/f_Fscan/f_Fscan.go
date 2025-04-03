package main

import (
	"fmt"
	"strings"
)

func main() {
	// 模拟从某个 io.Reader 读取的输入，这里使用 strings.Reader
	input := "张三 30 true 98.5"
	reader := strings.NewReader(input)

	var name string
	var age int
	var isStudent bool
	var score float64

	// 使用 fmt.Fscan 从 reader 中读取数据，并按顺序存储到变量中
	n, err := fmt.Fscan(reader, &name, &age, &isStudent, &score)
	if err != nil {
		fmt.Println("扫描错误:", err)
	}

	fmt.Printf("成功扫描了 %d 个值:\n", n)
	fmt.Printf("姓名: %s\n", name)
	fmt.Printf("年龄: %d\n", age)
	fmt.Printf("是学生: %t\n", isStudent)
	fmt.Printf("分数: %f\n", score)

	fmt.Println("\n--- 更多示例 ---")

	// 包含换行符的输入
	input2 := `Apple 10
Banana 20
Cherry 30`
	reader2 := strings.NewReader(input2)

	var fruit1 string
	var count1 int
	var fruit2 string
	var count2 int
	var fruit3 string
	var count3 int

	n2, err2 := fmt.Fscan(reader2, &fruit1, &count1, &fruit2, &count2, &fruit3, &count3)
	if err2 != nil {
		fmt.Println("扫描错误 (示例2):", err2)
	}

	fmt.Printf("成功扫描了 %d 个值 (示例2):\n", n2)
	fmt.Printf("%s: %d\n", fruit1, count1)
	fmt.Printf("%s: %d\n", fruit2, count2)
	fmt.Printf("%s: %d\n", fruit3, count3)

	fmt.Println("\n--- 类型不匹配示例 ---")

	input3 := "Hello world 123"
	reader3 := strings.NewReader(input3)

	var str1 string
	var num int
	var str2 string

	n3, err3 := fmt.Fscan(reader3, &str1, &num, &str2)
	if err3 != nil {
		fmt.Println("扫描错误 (示例3):", err3)
	}

	fmt.Printf("成功扫描了 %d 个值 (示例3):\n", n3)
	fmt.Printf("字符串 1: %s\n", str1)
	fmt.Printf("数字: %d\n", num)
	fmt.Printf("字符串 2: %s\n", str2)
}
