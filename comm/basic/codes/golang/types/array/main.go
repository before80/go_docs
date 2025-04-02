package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

//var ga9 = [3][...]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}    // invalid use of [...] array (outside a composite literal)
//var ga10 = [...][...]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // invalid use of [...] array (outside a composite literal)

var verbs = []string{"T", "v", "#v"}

func init() {

}

func main() {
	fmt.Println("C创建")
	fmt.Println("一维数组")
	var a0 [3]int
	var a1 = [3]int{1, 2, 3}
	var a2 [3]int = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3}
	ad1 := [...]int{1, 2, 3}
	mfp.PrintFmtVal("a0", a0, verbs)
	mfp.PrintFmtVal("a1", a1, verbs)
	mfp.PrintFmtVal("a2", a2, verbs)
	mfp.PrintFmtVal("a3", a3, verbs)
	mfp.PrintFmtVal("ad1", ad1, verbs)

	//a4 := make([3]int{1, 2, 3}) // 报错：[3]int{…} is not a type
	//mfp.PrintFmtVal("a4", a4, verbs)

	//a5 := new([3]int{1, 2, 3}) // 报错：[3]int{…} is not a type
	//mfp.PrintFmtVal("a5", a5, verbs)

	fmt.Println("多维数组")
	fmt.Println("二维数组")
	var a6 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var a7 [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var a8 = [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//var a9 = [...][...]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // 报错：invalid use of [...] array (outside a composite literal)
	ad2 := [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mfp.PrintFmtVal("a6", a6, verbs)
	mfp.PrintFmtVal("a7", a7, verbs)
	mfp.PrintFmtVal("a8", a8, verbs)
	//mfp.PrintFmtVal("a9", a9, verbs)
	mfp.PrintFmtVal("ad2", ad2, verbs)

	fmt.Println("三维数组")
	var a10 = [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	var a11 [2][2][2]int = [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	var a12 = [...][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	//var a13 = [...][...][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}} // 报错：invalid use of [...] array (outside a composite literal) 和 missing type in composite literal
	ad3 := [...][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	mfp.PrintFmtVal("a10", a10, verbs)
	mfp.PrintFmtVal("a11", a11, verbs)
	mfp.PrintFmtVal("a12", a12, verbs)
	//mfp.PrintFmtVal("a13", a13, verbs)
	mfp.PrintFmtVal("ad3", ad3, verbs)

	fmt.Println("U修改")
	fmt.Println("一维数组")

	a14 := [...]int{1, 2, 3}
	mfp.PrintFmtVal("a14", a14, verbs)

	a14[0] = 11
	mfp.PrintFmtVal("a14", a14, verbs)

	a14[len(a14)-1] = 33
	mfp.PrintFmtVal("a14", a14, verbs)

	pa141 := &a14[0]
	*pa141 = 111
	mfp.PrintFmtVal("a14", a14, verbs)
	fmt.Println("二维数组")

	a15 := [...][2]int{{1, 2}, {3, 4}}
	a15[0][0] = 11
	mfp.PrintFmtVal("a15", a15, verbs)

	a15[len(a15)-1][0] = 33
	mfp.PrintFmtVal("a15", a15, verbs)

	pa151 := &a15[0][0]
	*pa151 = 111
	mfp.PrintFmtVal("a15", a15, verbs)
	fmt.Println("三维数组和二维数组类似")

	fmt.Println("整个数组赋值")
	a16 := [...]int{1, 2, 3}
	mfp.PrintFmtVal("a16", a16, verbs)
	a16 = [...]int{2, 3, 4}
	mfp.PrintFmtVal("赋值后 a16", a16, verbs)
	//a16 = [...]int{2, 3, 4, 5} // 报错：cannot use [...]int{…} (value of type [4]int) as [3]int value in assignment
	//a16 = [...]string{"a", "b", "c"} // 报错：cannot use [...]string{…} (value of type [3]string) as [3]int value in assignment

	a17 := [...]int{1, 2, 3}
	fmt.Println("直接访问指定索引下标的元素")
	fmt.Println(a17[0])
	fmt.Println(a17[1])
	fmt.Println(a17[len(a17)-1])

	fmt.Println("遍历数组")
	for k, v := range a17 {
		if k%2 == 0 {
			fmt.Println(k, "->", v)
		}
	}
	mfp.PrintHr()
	for k, v := range a17 {
		fmt.Println(k, "->", v)
	}

	fmt.Println("数组指针和指针数组")
	fmt.Println("数组指针")
	a18 := [...]int{1, 2, 3}
	a19 := [...]int{1, 2, 3, 4}
	_ = a19
	var ptrA181 *[3]int
	ptrA181 = &a18
	mfp.PrintFmtVal("ptrA181", ptrA181, []string{"T", "v", "#v"})
	mfp.PrintFmtVal("*ptrA181", *ptrA181, []string{"T", "v", "#v"})
	//ptrA181 = &a19 // 报错：cannot use &a19 (value of type *[4]int) as *[3]int value in assignment
	mfp.PrintHr()
	fmt.Println("指针数组")
	xa201, xa202, xa203 := 1, 2, 3
	a20 := [...]*int{&xa201, &xa202, &xa203}
	mfp.PrintFmtVal("a20", a20, []string{"T", "v", "#v"})
	for k, v := range a20 {
		fmt.Println(k, "->", *v)
	}

	fmt.Println("易错点")
	fmt.Println("访问数组的最后一个元素")
	a21 := [...]int{1, 2, 3}
	//fmt.Println(a21[len(a21)]) // 报错：invalid argument: index 3 out of bounds [0:3]
	fmt.Println(a21[len(a21)-1]) // 正确方式

	fmt.Println("获取数组的相关属性")
	a22 := [...]int{1, 2, 3}
	fmt.Println("a22数组的长度 len(a22)=", len(a22))
	fmt.Println("a22数组的容量 cap(a22)=", cap(a22))

	fmt.Println("错误赋值：a1 = [4]int{1, 2, 3, 4} // cannot use [4]int{…} (value of type [4]int) as [3]int value in assignment")
	//a1 = [4]int{1, 2, 3, 4} // cannot use [4]int{…} (value of type [4]int) as [3]int value in assignment

	//a2 := [3]int{1, 2, 3}
	//a3 := [...]int{1, 2, 3}
	//a4 := [...]int{1, 2, 33}
	//
	//if a2 == a3 {
	//	fmt.Println("a2和a3竟然相等！")
	//} else {
	//	fmt.Println("a2和a3不相等！")
	//}
	//
	//if a2 == a4 {
	//	fmt.Println("a2和a4竟然相等！")
	//} else {
	//	fmt.Println("a2和a4不相等！")
	//}
	//
	//// 产生切片
	//sl1 := a2[:]
	//sl2 := a2[0:]
	//sl3 := a2[:len(a1)]
	//sl4 := a2[0:len(a1)]
	//sl5 := a2[1:2]
	//
	//mfp.PrintFmtVal("sl1", sl1, []string{"T", "v"})
	//mfp.PrintFmtVal("sl2", sl2, []string{"T", "v"})
	//mfp.PrintFmtVal("sl3", sl3, []string{"T", "v"})
	//mfp.PrintFmtVal("sl4", sl4, []string{"T", "v"})
	//mfp.PrintFmtVal("sl5", sl5, []string{"T", "v"})
}
