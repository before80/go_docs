package main

import "fmt"

type St struct {
	x, y int
	z    *int
}

func main() {
	//str := "hello"
	//pStr := &str
	//for i, v := range pStr { // cannot range over pStr (variable of type *string)
	//	fmt.Printf("%d->%q\n", i, v)
	//}

	//sl := []int{1, 2, 3}
	//pSl := &sl
	//for i, v := range pSl { // cannot range over pSl (variable of type *[]int)
	//	fmt.Printf("%d->%d\n", i, v)
	//}

	//mp := map[int]string{1: "a", 2: "b"}
	//pMp := &mp
	//for i, v := range pMp { //cannot range over pMp (variable of type *map[int]string)
	//	fmt.Printf("%d->%q\n", i, v)
	//}

	arr := [3]int{1, 2, 3}
	pArr := &arr
	for i, v := range pArr {
		fmt.Printf("%d->%d\n", i, v)
	}

	a, b, c := 1, 2, 3
	d := 6
	sts := []*St{
		&St{1, 2, &a},
		&St{1, 2, &b},
		&St{1, 2, &c},
	}

	for _, st := range sts {
		st.z = &d
	}
	fmt.Printf("d's address is %p\n", &d)
	fmt.Printf("%#v\n", sts)
}
