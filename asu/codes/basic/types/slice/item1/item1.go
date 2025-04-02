package main

import "fmt"

func main() {
	var a []string
	b := []string{}
	c := make([]string, 0)
	if a == nil {
		fmt.Println("a is equal nil.")
	}

	if b == nil {
		fmt.Println("b is equal nil.")
	}

	if c == nil {
		fmt.Println("1 c is equal nil.")
	}
	c = nil
	if c == nil {
		fmt.Println("2 c is equal nil.")
	}
}
