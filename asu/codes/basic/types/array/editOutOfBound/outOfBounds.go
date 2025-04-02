package main

import "fmt"

var a = [3]int{1, 2, 3}

func editA(i, v int) {
	a[i] = v
}

func main() {
	fmt.Println(a)
	editA(len(a), 4)
	fmt.Println(a)
}
