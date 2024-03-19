package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
	"maps"
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

	fmt.Println("复制map")
	fmt.Println("从go1.21版本开始才可以使用")
	fmt.Println("使用maps.Clone函数")
	m13 := map[string]int{"A": 1, "B": 2, "C": 3}
	mfp.PrintFmtValWithL("1 m13", m13, verbs)
	m14 := maps.Clone(m13)
	mfp.PrintFmtValWithL("2 m14", m14, verbs)
	m13["A"] = 11
	fmt.Println(`修改 m13["A"] = 11`)
	mfp.PrintFmtValWithL("3 m13", m13, verbs)
	mfp.PrintFmtValWithL("4 m14", m14, verbs)
	m14["B"] = 22
	fmt.Println(`修改 m14["B"] = 22`)
	mfp.PrintFmtValWithL("5 m13", m13, verbs)
	mfp.PrintFmtValWithL("6 m14", m14, verbs)
	mfp.PrintHr()
	fmt.Println("使用maps.Copy函数")
	m15 := map[string]int{"A": 1, "B": 2}
	m16 := map[string]int{"A": 11, "C": 33}
	fmt.Println(`使用Copy函数前`)
	mfp.PrintFmtValWithL("m15", m15, verbs)
	mfp.PrintFmtValWithL("m16", m16, verbs)
	maps.Copy(m16, m15) // func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)
	fmt.Println(`使用Copy函数后`)
	mfp.PrintFmtValWithL("m15", m15, verbs)
	mfp.PrintFmtValWithL("m16", m16, verbs)
	m15["A"] = 111
	fmt.Println(`修改 m15["A"] = 111`)
	mfp.PrintFmtValWithL("m15", m15, verbs)
	mfp.PrintFmtValWithL("m16", m16, verbs)
	m16["B"] = 222
	fmt.Println(`修改 m16["B"] = 222`)
	mfp.PrintFmtValWithL("m15", m15, verbs)
	mfp.PrintFmtValWithL("m16", m16, verbs)

	fmt.Println("使用maps.DeleteFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	m17 := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}
	fmt.Println("使用maps.DeleteFunc函数前")
	mfp.PrintFmtValWithL("m17", m17, verbs)
	maps.DeleteFunc(m17, func(k string, v int) bool {
		if v%2 == 1 {
			return true
		}
		return false
	})

	fmt.Println("使用maps.DeleteFunc函数后")
	mfp.PrintFmtValWithL("m17", m17, verbs)

	fmt.Println("判断相等")
	m18 := map[string]int{"A": 1, "B": 2, "C": 3}
	m19 := map[string]int{"A": 1, "B": 2, "C": 3}
	//fmt.Println("m18 == m19 -> ", m18 == m19) // 报错：invalid operation: m18 == m19 (map can only be compared to nil)
	//fmt.Println("m18 != m19 -> ", m18 != m19) // 报错：invalid operation: m18 != m19 (map can only be compared to nil)

	_ = m18
	_ = m19

	fmt.Println("使用maps.Equal函数")
	fmt.Println("从go1.21版本开始才可以使用")
	m20 := map[string]int{"A": 1, "B": 2}
	m21 := map[string]int{"A": 1, "B": 2}
	fmt.Println("m20 == m21 ->", maps.Equal(m20, m21))

	m22 := map[string]int{"A": 11, "B": 2}
	fmt.Println("m20 == m22 ->", maps.Equal(m20, m22))

	m23 := map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println("m20 == m23 ->", maps.Equal(m20, m23))

	fmt.Println("使用maps.EqualFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	m24 := map[string]int{"A": 1, "B": 2}
	m25 := map[string]int{"A": 1, "B": 2}
	fmt.Println("m24 == m25 -> ", maps.EqualFunc(m24, m25, func(v1 int, v2 int) bool {
		if v1 == v2 {
			return true
		}
		return false
	}))

	fmt.Println("直接对new函数创建的map进行key操作")
	m26 := *new(map[string]int)
	mfp.PrintFmtValWithL("1 m26", m26, verbs)
	//m26["A"] = 1 // 报错：panic: assignment to entry in nil map
	m26 = map[string]int{"A": 1} // 正确方式
	mfp.PrintFmtValWithL("2 m26", m26, verbs)
	m26["B"] = 2
	mfp.PrintFmtValWithL("3 m26", m26, verbs)

	fmt.Println("以为可以使用copy内置函数来复制一个map")
	m27 := map[string]int{"A": 1}
	m28 := make(map[string]int, 1)
	//copy(m28, m27) // 报错：invalid argument: copy expects slice arguments; found m28 (variable of type map[string]int) and m27 (variable of type map[string]int)
	_ = m27
	_ = m28
}
