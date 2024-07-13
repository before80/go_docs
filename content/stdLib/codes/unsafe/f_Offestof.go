package main

import (
	"before80.github.com/go_docs/stdlib/codes/unsafe/subpkg"
	"fmt"
	"unsafe"
)

func main() {
	m := subpkg.MyStruct1{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%d\n", unsafe.Sizeof(m))

	name1 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&m)) + unsafe.Sizeof('a')))
	*name1 = "zlx"
	fmt.Printf("%#v\n", m)

	n := subpkg.MyStruct2{}
	fmt.Printf("%#v\n", n)
	fmt.Printf("%d\n", unsafe.Sizeof(n))

	name2 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&n)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof("")))
	*name2 = "zlx"
	fmt.Printf("%#v\n", n)

}
