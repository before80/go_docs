package main

import "fmt"

func main() {
	alphabet := []string{"a", "b", "c"}

	fmt.Printf("alphabet=%q,len=%d,cap=%d\n", alphabet, len(alphabet), cap(alphabet))
	subset1 := make([]string, 3)
	copy(subset1, alphabet)
	fmt.Printf("subset1=%q,len=%d,cap=%d\n", subset1, len(subset1), cap(subset1))

	subset2 := make([]string, 0, 3)
	copy(subset2, alphabet)
	fmt.Printf("subset2=%q,len=%d,cap=%d\n", subset2, len(subset2), cap(subset2))

	subset3 := make([]string, 3, 3)
	copy(subset3, alphabet)
	fmt.Printf("subset3=%q,len=%d,cap=%d\n", subset3, len(subset3), cap(subset3))

	subset4 := []string{}
	copy(subset4, alphabet)
	fmt.Printf("subset4=%q,len=%d,cap=%d\n", subset4, len(subset4), cap(subset4))

	subset1[0] = "A"
	//subset2[1] = "B" // panic: runtime error: index out of range [1] with length 0
	fmt.Printf("alphabet=%q,len=%d,cap=%d\n", alphabet, len(alphabet), cap(alphabet))
	fmt.Printf("subset1=%q,len=%d,cap=%d\n", subset1, len(subset1), cap(subset1))
	//fmt.Printf("subset2=%q,len=%d,cap=%d\n", subset2, len(subset2), cap(subset2))
}
