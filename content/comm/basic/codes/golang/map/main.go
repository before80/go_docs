package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

var verbs = []string{"T", "v", "#v"}

func main() {
	fmt.Println("C创建")
	fmt.Println("直接创建")
	var m1 map[int]int
	var m2 map[string]int = map[string]int{"A": 1, "B": 2}
	var m3 = map[string]int{"A": 1, "B": 2}
	m4 := map[string]int{"A": 1, "B": 2}
	mfp.PrintFmtValWithL("m1", m1, verbs)
	mfp.PrintFmtValWithL("m2", m2, verbs)
	mfp.PrintFmtValWithL("m3", m3, verbs)
	mfp.PrintFmtValWithL("m4", m4, verbs)

	//fmt.Println(len(m4))
	//fmt.Println(cap(m4)) // 报错：invalid argument: m4 (variable of type map[string]int) for cap

	fmt.Println("用make创建")
	m5 := make(map[string]int)
	m6 := make(map[string]int, 3)
	//m7 := make(map[string]int, 3, 3) // 报错：invalid operation: make(map[string]int, 3, 3) expects 1 or 2 arguments; found 3
	mfp.PrintFmtValWithL("1 m5", m5, verbs)
	mfp.PrintFmtValWithL("2 m6", m6, verbs)
	//mfp.PrintFmtValWithL("m7", m7, verbs)

	m5["A"] = 1
	mfp.PrintFmtValWithL("3 m5", m5, verbs)
	m5["B"] = 2
	mfp.PrintFmtValWithL("4 m5", m5, verbs)
	m5["C"] = 3
	mfp.PrintFmtValWithL("5 m5", m5, verbs)
	m6["A"] = 1
	mfp.PrintFmtValWithL("6 m6", m6, verbs)
	m6["B"] = 2
	mfp.PrintFmtValWithL("7 m6", m6, verbs)
	m6["C"] = 3
	mfp.PrintFmtValWithL("8 m6", m6, verbs)
	m6["D"] = 4
	mfp.PrintFmtValWithL("9 m6", m6, verbs)

	fmt.Println("用new创建")
	m7 := *new(map[string]int)
	mfp.PrintFmtValWithL("m7", m7, verbs)

	//m7["A"] = 1 // 报错：panic: assignment to entry in nil map
	//mfp.PrintFmtValWithL("m7", m7, verbs)

	m7 = map[string]int{"A": 1}
	mfp.PrintFmtValWithL("m7", m7, verbs)

	fmt.Println("是否可以删除map中的某一元素？")
	m8 := map[string]int{"A": 1, "B": 2, "C": 3}
	mfp.PrintFmtValWithL("m8", m8, verbs)
	delete(m8, "A")
	mfp.PrintFmtValWithL("m8", m8, verbs)
	delete(m8, "A") // 重复删除，也不会报错
	mfp.PrintFmtValWithL("m8", m8, verbs)
	delete(m8, "B")
	mfp.PrintFmtValWithL("m8", m8, verbs)
	delete(m8, "C")
	mfp.PrintFmtValWithL("m8", m8, verbs)

	fmt.Println("修改元素")
	m9 := map[string]int{"A": 1, "B": 2, "C": 3}
	mfp.PrintFmtValWithL("1 m9", m9, verbs)
	m9["A"] = 11
	mfp.PrintFmtValWithL("2 m9", m9, verbs)
	m9["D"] = 4 // 修改不存在的Key
	mfp.PrintFmtValWithL("3 m9", m9, verbs)

	fmt.Println("用整个map赋值")
	m10 := map[string]int{"A": 1, "B": 2, "C": 3}
	mfp.PrintFmtValWithL("1 m10", m10, verbs)
	m10 = map[string]int{"A": 11, "B": 22, "C": 33, "D": 44}
	mfp.PrintFmtValWithL("2 m10", m10, verbs)
	m11 := map[string]int{"A": 111, "B": 222, "C": 333, "D": 444}
	m10 = m11
	mfp.PrintFmtValWithL("3 m10", m10, verbs)
	m11["A"] = 1
	mfp.PrintFmtValWithL("4 m10", m10, verbs)

	fmt.Println("直接访问指定Key的元素")
	m12 := map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println(m12["A"])
	fmt.Println(m12["B"])
	fmt.Println(m12["C"])
	fmt.Println(m12["D"]) // 访问不存在的Key

	fmt.Println("遍历map")
	for k, v := range m12 {
		fmt.Println(k, "->", v)
	}

	fmt.Println("获取相关map属性")
	fmt.Println("m12 map的长度 len(m12)=", len(m12))
}
