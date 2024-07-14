package main

import "github.com/before80/utils/mfp"

type St struct {
	a int
	b string
}

type Itf interface {
	M1()
	M2()
}

var verbs = []string{"T", "v", "#v"}

func main() {
	var aby [3]byte
	var abl [3]bool
	var as [3]string
	var ar [3]rune
	var autr [3]uintptr
	var ai8 [3]int8
	var ai16 [3]int16
	var ai32 [3]int32
	var ai64 [3]int64
	var ai [3]int
	var aui8 [3]uint8
	var aui16 [3]uint16
	var aui64 [3]uint64
	var aui32 [3]uint32
	var aui [3]uint
	var af32 [3]float32
	var af64 [3]float64
	var acplx64 [3]complex64
	var acplx128 [3]complex128
	var asli [3][]int
	var ast [3]St
	var aitf [3]Itf
	var af [3]func(int) int
	var am [3]map[string]int
	var ach [3]chan int
	mfp.PrintFmtVal("aby", aby, verbs)
	mfp.PrintFmtVal("abl", abl, verbs)
	mfp.PrintFmtVal("as", as, verbs)
	mfp.PrintFmtVal("ar", ar, verbs)
	mfp.PrintFmtVal("autr", autr, verbs)
	mfp.PrintFmtVal("ai8", ai8, verbs)
	mfp.PrintFmtVal("ai16", ai16, verbs)
	mfp.PrintFmtVal("ai32", ai32, verbs)
	mfp.PrintFmtVal("ai64", ai64, verbs)
	mfp.PrintFmtVal("ai", ai, verbs)
	mfp.PrintFmtVal("aui8", aui8, verbs)
	mfp.PrintFmtVal("aui16", aui16, verbs)
	mfp.PrintFmtVal("aui32", aui32, verbs)
	mfp.PrintFmtVal("aui64", aui64, verbs)
	mfp.PrintFmtVal("aui", aui, verbs)
	mfp.PrintFmtVal("af32", af32, verbs)
	mfp.PrintFmtVal("af64", af64, verbs)
	mfp.PrintFmtVal("acplx64", acplx64, verbs)
	mfp.PrintFmtVal("acplx128", acplx128, verbs)
	mfp.PrintFmtVal("asli", asli, verbs)
	mfp.PrintFmtVal("ast", ast, verbs)
	mfp.PrintFmtVal("asli", asli, verbs)
	mfp.PrintFmtVal("aitf", aitf, verbs)
	mfp.PrintFmtVal("af", af, verbs)
	mfp.PrintFmtVal("am", am, verbs)
	mfp.PrintFmtVal("ach", ach, verbs)
}
