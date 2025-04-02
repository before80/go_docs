package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	//a2 := [4]int{1, 2, 3, 4}
	a3 := [3]int{1, 2, 3}
	//a4 := [3]string{"1", "2", "3"}
	a5 := [3]int{2, 3, 4}
	// 报错：invalid operation: a1 == a2 (mismatched types [3]int and [4]int)
	//if a1 == a2 {
	//	fmt.Println("a1 == a2")
	//}

	if a1 == a3 {
		fmt.Println("a1 == a3")
	} else {
		fmt.Println("a1 != a3")
	}

	// 报错：invalid operation: a1 < a3 (operator < not defined on array)
	//if a1 < a3 {
	//	fmt.Println("a1 < a3")
	//}

	// 报错：invalid operation: a1 > a3 (operator > not defined on array)
	//if a1 > a3 {
	//	fmt.Println("a1 > a3")
	//}

	// 报错：invalid operation: a1 <= a3 (operator <= not defined on array)
	//if a1 <= a3 {
	//	fmt.Println("a1 <= a3")
	//}

	// 报错：invalid operation: a1 >= a3 (operator >= not defined on array)
	//if a1 >= a3 {
	//	fmt.Println("a1 >= a3")
	//}

	// 报错：invalid operation: a1 == a4 (mismatched types [3]int and [3]string)
	//if a1 == a4 {
	//	fmt.Println("a1 == a4")
	//}

	if a1 == a5 {
		fmt.Println("a1 == a5")
	} else {
		fmt.Println("a1 != a5")
	}
}
