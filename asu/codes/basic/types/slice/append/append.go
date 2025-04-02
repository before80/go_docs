package main

func main() {
	var si []int
	var sb []byte
	si0 := []int{1, 2, 3}
	//ai0 := [3]int{1, 2, 3}
	si = append(si, si0...)
	sb = append(sb, "hello"...)

}
