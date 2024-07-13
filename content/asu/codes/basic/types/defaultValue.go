package main

import "fmt"

func main() {
	fmt.Println("default value:")
	var b bool
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	var ui uint
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var i int
	var s string
	var ai [3]int
	var as [3]string
	var ab [3]bool
	var sli []int
	var sls []string
	var slb []bool
	type St struct {
		a int
		b string
		c *int
		d chan int
	}
	var st St
	var pti *int
	var pts *string
	var ptb *bool
	var uptr uintptr
	var fti func(int, int) int
	var fts func(string, string) string
	var ftb func(bool, bool) bool
	type ITF interface {
		M1(int) int
		M2(string) string
		M3(bool) bool
	}
	var itf ITF
	var mii map[int]int
	var mss map[string]string
	var mbb map[bool]bool
	var chi chan int
	var chs chan string
	var chb chan bool

	fmt.Printf("bool -> %t\n", b)
	fmt.Printf("uint8 -> %d\n", ui8)
	fmt.Printf("uint16 -> %d\n", ui16)
	fmt.Printf("uint32 -> %d\n", ui32)
	fmt.Printf("uint64 -> %d\n", ui64)
	fmt.Printf("uint -> %d\n", ui)
	fmt.Printf("int8 -> %d\n", i8)
	fmt.Printf("int16 -> %d\n", i16)
	fmt.Printf("int32 -> %d\n", i32)
	fmt.Printf("int64 -> %d\n", i64)
	fmt.Printf("int -> %d\n", i)
	fmt.Printf("string -> %q\n", s)
	fmt.Printf("array [3]int -> %#v\n", ai)
	fmt.Printf("array [3]string -> %#v\n", as)
	fmt.Printf("array [3]bool -> %#v\n", ab)
	fmt.Printf("slice []int -> %#v\n", sli)
	fmt.Printf("slice []int -> %v\n", sli)
	if sli == nil {
		fmt.Println("sli is nil.")
	}
	fmt.Printf("slice []string -> %#v\n", sls)
	fmt.Printf("slice []string -> %v\n", sls)
	if sls == nil {
		fmt.Println("sls is nil.")
	}
	fmt.Printf("slice []bool -> %#v\n", slb)
	fmt.Printf("slice []bool -> %v\n", slb)
	if slb == nil {
		fmt.Println("slb is nil.")
	}
	fmt.Printf("struct st -> %#v\n", st)
	fmt.Printf("struct st -> %v\n", st)
	fmt.Printf("point *int -> %v\n", pti)
	if pti == nil {
		fmt.Println("pti is nil.")
	}
	fmt.Printf("point *string -> %v\n", pts)
	if pts == nil {
		fmt.Println("pts is nil.")
	}
	fmt.Printf("point *bool -> %v\n", ptb)
	if ptb == nil {
		fmt.Println("ptb is nil.")
	}
	fmt.Printf("uintptr -> %v\n", uptr)
	fmt.Printf("func(int,int) int -> %v\n", fti)
	if fti == nil {
		fmt.Println("fti is nil.")
	}
	fmt.Printf("func(string,string) string -> %v\n", fts)
	if fts == nil {
		fmt.Println("fts is nil.")
	}

	fmt.Printf("func(bool,bool) bool -> %v\n", ftb)
	if ftb == nil {
		fmt.Println("ftb is nil.")
	}

	fmt.Printf("interface -> %v\n", itf)
	if itf == nil {
		fmt.Println("itf is nil.")
	}

	fmt.Printf("map[int]int -> %#v\n", mii)
	fmt.Printf("map[int]int -> %v\n", mii)
	if mii == nil {
		fmt.Println("mii is nil.")
	}
	fmt.Printf("map[string]string -> %#v\n", mss)
	fmt.Printf("map[string]string -> %v\n", mss)
	if mss == nil {
		fmt.Println("mss is nil.")
	}
	fmt.Printf("map[bool]bool -> %#v\n", mbb)
	fmt.Printf("map[bool]bool -> %v\n", mbb)
	if mbb == nil {
		fmt.Println("mbb is nil.")
	}

	fmt.Printf("chan int -> %#v\n", chi)
	fmt.Printf("chan int -> %v\n", chi)
	if chi == nil {
		fmt.Println("chi is nil.")
	}
	fmt.Printf("chan string -> %#v\n", chs)
	fmt.Printf("chan string -> %v\n", chs)
	if chs == nil {
		fmt.Println("chs is nil.")
	}
	fmt.Printf("chan bool -> %#v\n", chb)
	fmt.Printf("chan bool -> %v\n", chb)
	if chb == nil {
		fmt.Println("chb is nil.")
	}
}

//default value:
//bool -> false
//uint8 -> 0
//uint16 -> 0
//uint32 -> 0
//uint64 -> 0
//uint -> 0
//int8 -> 0
//int16 -> 0
//int32 -> 0
//int64 -> 0
//int -> 0
//string -> ""
//array [3]int -> [3]int{0, 0, 0}
//array [3]string -> [3]string{"", "", ""}
//array [3]bool -> [3]bool{false, false, false}
//slice []int -> []int(nil)
//slice []int -> []
//sli is nil.
//slice []string -> []string(nil)
//slice []string -> []
//sls is nil.
//slice []bool -> []bool(nil)
//slice []bool -> []
//slb is nil.
//struct st -> main.St{a:0, b:"", c:(*int)(nil), d:(chan int)(nil)}
//struct st -> {0  <nil> <nil>}
//point *int -> <nil>
//pti is nil.
//point *string -> <nil>
//pts is nil.
//point *bool -> <nil>
//ptb is nil.
//uintptr -> 0
//func(int,int) int -> <nil>
//fti is nil.
//func(string,string) string -> <nil>
//fts is nil.
//func(bool,bool) bool -> <nil>
//ftb is nil.
//interface -> <nil>
//itf is nil.
//map[int]int -> map[int]int(nil)
//map[int]int -> map[]
//mii is nil.
//map[string]string -> map[string]string(nil)
//map[string]string -> map[]
//mss is nil.
//map[bool]bool -> map[bool]bool(nil)
//map[bool]bool -> map[]
//mbb is nil.
//chan int -> (chan int)(nil)
//chan int -> <nil>
//chi is nil.
//chan string -> (chan string)(nil)
//chan string -> <nil>
//chs is nil.
//chan bool -> (chan bool)(nil)
//chan bool -> <nil>
//chb is nil.
