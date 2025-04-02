package main

import "fmt"

func main() {
	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	s2 = append(s2, 2)
	fmt.Printf("s1 len=%d,cap=%d,%v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2 len=%d,cap=%d,%v\n", len(s2), cap(s2), s2)
	s2 = append(s2, 3)
	s2 = append(s2, 4)
	s2 = append(s2, 5)
	fmt.Printf("s1 len=%d,cap=%d,%v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2 len=%d,cap=%d,%v\n", len(s2), cap(s2), s2)

	s1 = append(s1, 1)
	fmt.Printf("s1 len=%d,cap=%d,%v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 2)
	fmt.Printf("s1 len=%d,cap=%d,%v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 3)
	fmt.Printf("s1 len=%d,cap=%d,%v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 4)
	fmt.Printf("s1 len=%d,cap=%d,%v\n", len(s1), cap(s1), s1)
}
