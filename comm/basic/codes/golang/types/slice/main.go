package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/before80/utils/mfp"
)

var verbs = []string{"T", "v", "#v"}

func main() {
	fmt.Println("C创建")
	fmt.Println("直接创建")
	var sl1 []int
	var sl2 []int = []int{1, 2, 3}
	var sl3 = []int{1, 2, 3}
	sl4 := []int{1, 2, 3}
	mfp.PrintFmtValWithLC("sl1", sl1, verbs)
	mfp.PrintFmtValWithLC("sl2", sl2, verbs)
	mfp.PrintFmtValWithLC("sl3", sl3, verbs)
	mfp.PrintFmtValWithLC("sl4", sl4, verbs)

	fmt.Println("基于数组创建")

	a1 := [...]int{1, 2, 3, 4, 5, 6}
	mfp.PrintFmtValWithLC("a1", a1, verbs)
	sl5 := a1[:]
	sl6 := a1[0:]
	sl7 := a1[:len(a1)]
	sl8 := a1[0:len(a1)]
	//sl9 := a1[0:3:2] // 报错：invalid slice indices: 2 < 3
	sl10 := a1[0:3:3]
	sl11 := a1[0:3:4]
	sl12 := a1[0:3:5]
	sl13 := a1[0:3:6]
	//sl14 := a1[0:3:7] // 报错：invalid argument: index 7 out of bounds [0:7]
	mfp.PrintFmtValWithLC("sl5", sl5, verbs)
	mfp.PrintFmtValWithLC("sl6", sl6, verbs)
	mfp.PrintFmtValWithLC("sl7", sl7, verbs)
	mfp.PrintFmtValWithLC("sl8", sl8, verbs)
	//mfp.PrintFmtValWithLC("sl9", sl9, verbs)
	mfp.PrintFmtValWithLC("sl10", sl10, verbs)
	mfp.PrintFmtValWithLC("sl11", sl11, verbs)
	mfp.PrintFmtValWithLC("sl12", sl12, verbs)
	mfp.PrintFmtValWithLC("sl13", sl13, verbs)
	//mfp.PrintFmtValWithLC("sl14", sl14, verbs)

	fmt.Println("用make创建")
	sl15 := make([]int, 3)
	//sl16 := make([]int, 3, 2) // 报错：invalid argument: length and capacity swapped
	sl17 := make([]int, 3, 3)
	sl18 := make([]int, 3, 4)
	mfp.PrintFmtValWithLC("sl15", sl15, verbs)
	//mfp.PrintFmtValWithLC("sl16", sl16, verbs)
	mfp.PrintFmtValWithLC("sl17", sl17, verbs)
	mfp.PrintFmtValWithLC("sl18", sl18, verbs)

	fmt.Println("用new创建")
	sl19 := *new([]int) // 注意此时 sl19 为空切片，其长度和容量都为0
	mfp.PrintFmtValWithLC("sl19", sl19, verbs)
	sl19 = append(sl19, 1)
	mfp.PrintFmtValWithLC("sl19", sl19, verbs)
	sl19 = append(sl19, 2)
	mfp.PrintFmtValWithLC("sl19", sl19, verbs)
	sl19 = append(sl19, 3)
	mfp.PrintFmtValWithLC("sl19", sl19, verbs)

	fmt.Println("基于已有切片创建")
	a2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mfp.PrintFmtValWithLC("已有数组 a2", a2, verbs)

	sl20 := a2[0:6]
	mfp.PrintFmtValWithLC("已有切片 sl20", sl20, verbs)

	sl21 := sl20[:]
	sl22 := sl20[0:]
	sl23 := sl20[:len(sl20)]
	sl24 := sl20[:cap(sl20)]
	sl25 := sl20[0:len(sl20)]
	sl26 := sl20[0:cap(sl20)]
	//sl27 := sl20[0:cap(sl20)+1] // 报错：panic: runtime error: slice bounds out of range [:11] with capacity 10
	sl28 := sl20[1:3]
	sl29 := sl20[1:4]
	sl30 := sl20[2:4]
	//sl31 := sl20[2:4:2] // 报错：invalid slice indices: 2 < 4
	//sl32 := sl20[2:4:3] // 报错：invalid slice indices: 3 < 4
	sl33 := sl20[2:4:4]
	sl34 := sl20[2:4:5]
	sl35 := sl20[2:4:6]
	sl36 := sl20[2:4:7]

	mfp.PrintFmtValWithLC("sl21=sl20[:]", sl21, verbs)
	mfp.PrintFmtValWithLC("sl22=sl20[0:]", sl22, verbs)
	mfp.PrintFmtValWithLC("sl23=sl20[:len(sl20)]", sl23, verbs)
	mfp.PrintFmtValWithLC("sl24=sl20[:cap(sl20)]", sl24, verbs)
	mfp.PrintFmtValWithLC("sl25=[0:len(sl20)]", sl25, verbs)
	mfp.PrintFmtValWithLC("sl26=[0:cap(sl20)]", sl26, verbs)
	//mfp.PrintFmtValWithLC("sl27=sl20[0:cap(sl20)+1]", sl27, verbs)
	mfp.PrintFmtValWithLC("sl28=sl20[1:3]", sl28, verbs)
	mfp.PrintFmtValWithLC("sl29=sl20[1:4]", sl29, verbs)
	mfp.PrintFmtValWithLC("sl30=sl20[2:4]", sl30, verbs)
	//mfp.PrintFmtValWithLC("sl31=sl20[2:4:2]", sl31, verbs)
	//mfp.PrintFmtValWithLC("sl32=sl20[2:4:3]", sl32, verbs)
	mfp.PrintFmtValWithLC("sl33=sl20[2:4:4]", sl33, verbs)
	mfp.PrintFmtValWithLC("sl34=sl20[2:4:5]", sl34, verbs)
	mfp.PrintFmtValWithLC("sl35=sl20[2:4:6]", sl35, verbs)
	mfp.PrintFmtValWithLC("sl36=sl20[2:4:7]", sl36, verbs)

	fmt.Println("修改元素")
	sl37 := []int{1, 2, 3}
	mfp.PrintFmtValWithLC("sl37", sl37, verbs)
	sl37[0] = 11
	mfp.PrintFmtValWithLC("sl37", sl37, verbs)
	sl37[len(sl37)-1] = 33
	mfp.PrintFmtValWithLC("sl37", sl37, verbs)
	// 修改不存在的元素
	//sl37[3] = 4 // 报错：panic: runtime error: index out of range [3] with length 3

	fmt.Println("用整个切片赋值")

	sl38 := []int{1, 2, 3}
	mfp.PrintFmtValWithLC("1 sl38", sl38, verbs)
	sl38 = []int{1, 2, 3, 4}
	mfp.PrintFmtValWithLC("2 sl38", sl38, verbs)
	sl38 = make([]int, 5, 10)
	mfp.PrintFmtValWithLC("3 sl38", sl38, verbs)
	sl38 = *new([]int)
	mfp.PrintFmtValWithLC("4 sl38", sl38, verbs)
	sl39 := []int{1, 2, 3, 4, 5, 6}
	mfp.PrintFmtValWithLC("5 sl39", sl39, verbs)
	sl38 = sl39
	mfp.PrintFmtValWithLC("6 sl38", sl38, verbs)

	fmt.Println("直接访问指定索引下标的元素")
	sl40 := []int{1, 2, 3}
	fmt.Println("sl40[0]=", sl40[0])
	fmt.Println("sl40[1]=", sl40[1])
	fmt.Println("sl40[2]=", sl40[2])

	fmt.Println("遍历切片")
	for k, v := range sl40 {
		if k%2 == 0 {
			fmt.Println(k, "->", v)
		}
	}
	mfp.PrintHr()
	for k, v := range sl40 {
		fmt.Println(k, "->", v)
	}
	mfp.PrintHr()
	IamNaN5 := math.NaN()
	sl40x := []float64{0, 42.12, -10.123, 8, IamNaN5}
	for k, v := range sl40x {
		fmt.Println(k, "->", v)
	}

	fmt.Println("获取相关切片属性")
	sl41 := []int{1, 2, 3}
	fmt.Println("sl41切片的长度 len(sl41)=", len(sl41))
	fmt.Println("sl41切片的容量 cap(sl41)=", cap(sl41))

	fmt.Println("是否可以删除某一元素呢？")
	sl42 := []int{1, 2, 3, 4, 5, 6}
	i := 3 // 需要删除元素的索引下标
	mfp.PrintFmtValWithLC("1 sl42", sl42, verbs)
	sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 索引为3的元素
	mfp.PrintFmtValWithLC("2 sl42", sl42, verbs)
	sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 当前索引为3的元素
	mfp.PrintFmtValWithLC("3 sl42", sl42, verbs)
	i = 0
	sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 当前索引为0的元素
	mfp.PrintFmtValWithLC("4 sl42", sl42, verbs)
	sl42 = sl42[:len(sl42)-1] // 删除当前的最后一个元素
	mfp.PrintFmtValWithLC("5 sl42", sl42, verbs)
	sl42 = sl42[1:] // 删除当前的第一个元素
	mfp.PrintFmtValWithLC("6 sl42", sl42, verbs)

	fmt.Println("复制切片")
	fmt.Println("使用copy函数")
	slSrc43 := make([]int, 3, 6)
	slSrc43 = slices.Replace(slSrc43, 0, 3, []int{1, 2, 3}...)
	mfp.PrintFmtValWithLC("1 slSrc43", slSrc43, verbs)
	slDst44 := make([]int, len(slSrc43))
	mfp.PrintFmtValWithLC("2 slDst44", slDst44, verbs)

	copy(slDst44, slSrc43) // func copy(dst []Type, src []Type) int
	mfp.PrintFmtValWithLC("3 slDst44", slDst44, verbs)
	slDst44[0] = 11
	fmt.Println("slDst44[0] = 11 之后")
	mfp.PrintFmtValWithLC("4 slDst43", slSrc43, verbs)
	mfp.PrintFmtValWithLC("5 slDst44", slDst44, verbs)
	slSrc43[1] = 22
	fmt.Println("slSrc43[1] = 22 之后")
	mfp.PrintFmtValWithLC("6 slDst43", slSrc43, verbs)
	mfp.PrintFmtValWithLC("7 slDst44", slDst44, verbs)

	fmt.Println("访问切片的最后一个元素")
	sl45 := []int{1, 2, 3}
	//fmt.Println(sl45[len(sl45)])   // 报错：panic: runtime error: index out of range [3] with length 3
	fmt.Println(sl45[len(sl45)-1]) // 正确方式

	fmt.Println("判断相等")
	fmt.Println("是否可以使用`==`或`!=`?")
	sl46 := []int{1, 2, 3}
	sl47 := []int{1, 2, 3}
	//fmt.Println("sl46 == sl47 -> ", sl46 == sl47) // 报错：invalid operation: sl46 == sl47 (slice can only be compared to nil)
	//fmt.Println("sl46 != sl47 -> ", sl46 != sl47)// 报错：invalid operation: sl46 != sl47 (slice can only be compared to nil)

	_ = sl46
	_ = sl47

	fmt.Println("使用slices.Equal函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl48 := []int{1, 2, 3}
	sl49 := []int{1, 2, 3}
	sl50 := []int{11, 2, 3}
	sl51 := []int{1, 2, 3, 4}
	sl48x1 := make([]int, 3, 6)
	sl48x1 = slices.Replace(sl48x1, 0, 3, []int{1, 2, 3}...)
	mfp.PrintFmtValWithLC("sl48x1", sl48x1, verbs)
	fmt.Println("sl48 == sl49 -> ", slices.Equal(sl48, sl49))
	fmt.Println("sl48 == sl50 -> ", slices.Equal(sl48, sl50))
	fmt.Println("sl48 == sl51 -> ", slices.Equal(sl48, sl51))
	fmt.Println("sl48 == sl48x1 -> ", slices.Equal(sl48, sl48x1))

	fmt.Println("使用slices.EqualFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl52 := []int{1, 15, 8}
	sl53 := []int{1, 15, 8}
	sl54 := []int{11, 15, 8}
	sl55 := []string{"01", "0x0f", "0o10"}
	sl52x1 := make([]int, 3, 6)
	sl52x1 = slices.Replace(sl52x1, 0, 3, []int{1, 15, 8}...)
	mfp.PrintFmtValWithLC("sl52x1", sl52x1, verbs)
	feq1 := func(e1, e2 int) bool {
		return e1 == e2
	}
	feq2 := func(e1 int, e2 string) bool {
		sn, err := strconv.ParseInt(e2, 0, 64)
		if err != nil {
			return false
		}
		return e1 == int(sn)
	}
	fmt.Println("sl52 == sl53 -> ", slices.EqualFunc(sl52, sl53, feq1))
	fmt.Println("sl52 == sl54 -> ", slices.EqualFunc(sl52, sl54, feq1))
	fmt.Println("sl52 == sl55 -> ", slices.EqualFunc(sl52, sl55, feq2))
	fmt.Println("sl52 == sl52x1 -> ", slices.EqualFunc(sl52, sl52x1, feq1))

	fmt.Println("判断是否存在")
	fmt.Println("使用for循环")
	sl56 := []int{1, 2, 3}
	forFunc := func(src []int, target int) bool {
		for _, v := range src {
			if v == target {
				return true
			}
		}
		return false
	}

	fmt.Println("1 在 sl56中 -> ", forFunc(sl56, 1))
	fmt.Println("4 在 sl56中 -> ", forFunc(sl56, 4))

	fmt.Println("使用slices.Contains函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl57 := []int{1, 2, 3}
	fmt.Println("1 在 sl57中 -> ", slices.Contains(sl57, 1))
	fmt.Println("4 在 sl57中 -> ", slices.Contains(sl57, 4))

	fmt.Println("使用slices.ContainsFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl58 := []int{0, 42, -10, 8}

	fmt.Println("sl58中存在负数 -> ", slices.ContainsFunc(sl58, func(e int) bool {
		return e < 0
	}))
	fmt.Println("sl58中存在奇数 -> ", slices.ContainsFunc(sl58, func(e int) bool {
		return e%2 == 1
	}))
	fmt.Println("sl58中存在 8 -> ", slices.ContainsFunc(sl58, func(e int) bool {
		return e == 8
	}))

	fmt.Println("获取最大值")
	fmt.Println("使用for循环")
	sl59 := []int{0, 42, -10, 8}

	maxK := 0
	maxV := sl59[0]
	for k, v := range sl59 {
		if maxV < v {
			maxK = k
			maxV = v
		}
	}
	fmt.Printf("sl59中的最大值是sl59[%d]=%d\n", maxK, maxV)

	fmt.Println("使用slices.Max函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl60 := []int{0, 42, -10, 8}
	IamNaN := math.NaN()
	sl61 := []float64{0, 42.12, -10.123, 8, IamNaN}
	//sl62 := []int{0, 42, -10, 8, IamNaN} // 报错：cannot use IamNaN (variable of type float64) as int value in array or slice literal
	fmt.Printf("sl60中的最大值是%d\n", slices.Max(sl60))

	maxV2 := slices.Max(sl61)
	fmt.Printf("sl61中的最大值是%f（%T）\n", maxV2, maxV2)

	//sl63 := []int{}
	//fmt.Printf("sl60中的最大值是%d\n", slices.Max(sl63)) // 报错：panic: slices.Max: empty list

	fmt.Println("使用slices.MaxFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl64 := []int{0, 42, -10, 8}
	IamNaN2 := math.NaN()
	sl65 := []float64{0, 42.12, -10.123, 8, IamNaN2}
	fmt.Printf("sl64中最大值是%d\n", slices.MaxFunc(sl64, func(e1, e2 int) int {
		return cmp.Compare(e1, e2)
	}))

	fmt.Printf("sl65中最大值是%f\n", slices.MaxFunc(sl65, func(e1, e2 float64) int {
		return cmp.Compare(e1, e2)
	}))

	//sl66 := []int{}
	//fmt.Printf("sl66中最大值是%d\n", slices.MaxFunc(sl66, func(e1, e2 int) int {
	//	return cmp.Compare(e1, e2)
	//})) // 报错：panic: slices.Max: empty list

	fmt.Println("获取最小值")
	fmt.Println("使用for循环")
	sl67 := []int{0, 42, -10, 8}
	minK1, minV1 := findMin(0, sl67[0], sl67)
	fmt.Printf("sl67中的最小值是sl67[%d]=%d\n", minK1, minV1)

	sl68 := []float64{0, 42.12, -10.123, 8}
	minK2, minV2 := findMin(0, sl68[0], sl68)
	fmt.Printf("sl68中的最小值是sl68[%d]=%f\n", minK2, minV2)

	IamNaN3 := math.NaN()
	sl69 := []float64{0, 42.12, -10.123, 8, IamNaN3}
	minK3, minV3 := findMin(0, sl69[0], sl69)
	fmt.Printf("sl69中的最小值是sl69[%d]=%f\n", minK3, minV3)

	fmt.Println("使用slices.Min函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl70 := []int{0, 42, -10, 8}
	sl71 := []float64{0, 42.12, -10.123, 8}
	IamNaN4 := math.NaN()
	sl72 := []float64{0, 42.12, -10.123, 8, IamNaN4}
	fmt.Println("sl70中的最小值是", slices.Min(sl70))
	fmt.Println("sl71中的最小值是", slices.Min(sl71))
	fmt.Println("sl72中的最小值是", slices.Min(sl72))

	fmt.Println("使用slices.MinFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	fmt.Println("sl70中的最小值是", slices.MinFunc(sl70, func(a, b int) int {
		return cmp.Compare(a, b)
	}))
	fmt.Println("sl71中的最小值是", slices.MinFunc(sl71, func(a, b float64) int {
		return cmp.Compare(a, b)
	}))
	fmt.Println("sl72中的最小值是", slices.MinFunc(sl72, func(a, b float64) int {
		return cmp.Compare(a, b)
	}))

	fmt.Println("替换")
	fmt.Println("使用for循环")
	sl73 := make([]int, 6, 10)
	mfp.PrintFmtValWithLC("1 sl73", sl73, verbs)
	// 将 sl73[0]~sl73[6]依次替换为 1~6
	for k, _ := range sl73 {
		if k <= 6 {
			sl73[k] = k + 1
		}
	}
	mfp.PrintFmtValWithLC("2 sl73", sl73, verbs)

	fmt.Println("使用slices.Replace函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl74 := make([]int, 6, 10)
	mfp.PrintFmtValWithLC("1 sl74", sl74, verbs)
	sl74 = slices.Replace(sl74, 0, 6, []int{1, 2, 3, 4, 5, 6}...)
	mfp.PrintFmtValWithLC("2 sl74", sl74, verbs)
	sl74 = slices.Replace(sl74, 0, 1, 111)
	mfp.PrintFmtValWithLC("3 sl74", sl74, verbs)
	//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
	//mfp.PrintFmtValWithLC("4 sl74", sl74, verbs)
	//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6, 7}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
	//mfp.PrintFmtValWithLC("5 sl74", sl74, verbs)

	fmt.Println("长度和容量不一致的切片")
	sl75 := make([]int, 3, 6)
	mfp.PrintFmtValWithLC("1 sl75", sl75, verbs)
	//sl75[3] = 4 // 报错：panic: runtime error: index out of range [3] with length 3
	//mfp.PrintFmtValWithLC("2 sl75", sl75, verbs)
	//sl75[4] = 5 // 报错：panic: runtime error: index out of range [4] with length 3
	//mfp.PrintFmtValWithLC("3 sl75", sl75, verbs)

	fmt.Println("反转")
	fmt.Println("使用for循环")
	sl76 := []int{1, 2, 3, 4, 5, 6}
	mfp.PrintFmtValWithLC("1 sl76", sl76, verbs)
	reverseSlice(sl76)
	mfp.PrintFmtValWithLC("2 sl76", sl76, verbs)

	fmt.Println("使用slices.Reverse函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl77 := []int{1, 2, 3, 4, 5, 6}
	mfp.PrintFmtValWithLC("1 sl77", sl77, verbs)
	slices.Reverse(sl77)
	mfp.PrintFmtValWithLC("2 sl77", sl77, verbs)

	fmt.Println("移除未使用的容量")
	fmt.Println("使用slices.Clip函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl78 := make([]int, 3, 6)
	mfp.PrintFmtValWithLC("1 sl78", sl78, verbs)
	sl78 = slices.Clip(sl78)
	mfp.PrintFmtValWithLC("2 sl78", sl78, verbs)

	fmt.Println("易错点：使用slices.Replace函数")
	fmt.Println("错误的方式 1")
	sl79 := []int{1, 2, 3}
	mfp.PrintFmtValWithLC("1 sl79", sl79, verbs)
	// 要修改索引0处的元素值
	sl79 = slices.Replace(sl19, 0, 0, 111)
	mfp.PrintFmtValWithLC("2 sl79", sl79, verbs)

	fmt.Println("错误的方式 2")
	sl81 := []int{1, 2, 3}
	mfp.PrintFmtValWithLC("1 sl81", sl81, verbs)
	fmt.Println("若 i == j == len(sl) 呢？")
	sl81 = slices.Replace(sl81, 3, 3, 111)
	mfp.PrintFmtValWithLC("2 sl81", sl81, verbs)

	fmt.Println("正确的方式")
	sl80 := []int{1, 2, 3}
	mfp.PrintFmtValWithLC("1 sl80", sl80, verbs)
	sl80 = slices.Replace(sl80, 0, 1, 111)
	mfp.PrintFmtValWithLC("2 sl80", sl80, verbs)

	fmt.Println("排序")
	fmt.Println("使用slices.Sort函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl82 := []float64{0, 42.12, -10.123, 8, math.NaN()}
	mfp.PrintFmtValWithL("1 sl82", sl82, verbs)
	slices.Sort(sl82)
	mfp.PrintFmtValWithL("2 sl82", sl82, verbs)

	type Person struct {
		name string
		age  int8
	}

	sl83 := []Person{
		{"zlx2", 30},
		{"zlx1", 32},
		{"zlx3", 29},
	}

	mfp.PrintFmtValWithLC("1 sl83", sl83, verbs)
	//slices.Sort(sl83) // 报错：Person does not satisfy cmp.Ordered (Person missing in ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string)
	//mfp.PrintFmtValWithLC("2 sl83", sl83, verbs)

	sl84 := []int{2, 4, 6, 8, 1, 3, 5, 7}
	mfp.PrintFmtValWithLC("1 sl84", sl84, verbs)
	slices.Sort(sl84)
	mfp.PrintFmtValWithLC("2 sl84", sl84, verbs)

	sl85 := make([]int, 3, 6)
	sl85 = slices.Replace(sl85, 0, 3, []int{2, 1, 3}...)
	mfp.PrintFmtValWithLC("1 sl85", sl85, verbs)
	slices.Sort(sl85)
	mfp.PrintFmtValWithLC("2 sl85", sl85, verbs)

	fmt.Println("使用slices.SortFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl86 := []Person{
		{"zlx2", 30},
		{"zlx1", 32},
		{"zlx3", 29},
	}
	mfp.PrintFmtValWithLC("1 sl86", sl86, verbs)
	slices.SortFunc(sl86, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	mfp.PrintFmtValWithLC("2 sl86", sl86, verbs)

	sl88 := []Person{
		{"Gopher", 13},
		{"Alice", 55},
		{"Bob", 24},
		{"Alice", 20},
	}
	mfp.PrintFmtValWithLC("1 sl88", sl88, verbs)
	slices.SortFunc(sl88, func(a, b Person) int {
		if n := cmp.Compare(a.name, b.name); n != 0 {
			return n
		}
		// 如果 name 字段的值相等，则继续按 age 字段进行排序
		return cmp.Compare(a.age, b.age)
	})
	mfp.PrintFmtValWithLC("2 sl88", sl88, verbs)

	fmt.Println("使用slices.SortStableFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")

	sl89 := []Person{
		{"Gopher", 13},
		{"Alice", 55},
		{"Bob", 24},
		{"Alice", 30},
		{"Alice", 20},
	}
	mfp.PrintFmtValWithLC("1 sl89", sl89, verbs)
	slices.SortStableFunc(sl89, func(a, b Person) int {
		return cmp.Compare(a.name, b.name)
	})
	mfp.PrintFmtValWithLC("2 sl89", sl89, verbs)

	fmt.Println("使用slices.Delete函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl87 := []int{1, 2, 3, 4, 5, 6}
	mfp.PrintFmtValWithLC("1 sl87", sl87, verbs)
	sl87 = slices.Delete(sl87, 0, 0) // 注意这里并没有删除成功
	mfp.PrintFmtValWithLC("2 sl87", sl87, verbs)
	sl87 = slices.Delete(sl87, 0, 1) // 这里才会删除成功
	mfp.PrintFmtValWithLC("3 sl87", sl87, verbs)

	fmt.Println("使用slices.Compare函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl90 := []int{1, 2, 3}
	sl91 := []int{1, 2, 3}
	sl92 := []int{1, 2, 3, 4}
	sl93 := []int{11, 2, 3}
	fmt.Println("sl90 == sl91 ->", slices.Compare(sl90, sl91) == 0)
	fmt.Println("sl90 == sl92 ->", slices.Compare(sl90, sl92) == 0)
	fmt.Println("sl90 == sl93 ->", slices.Compare(sl90, sl93) == 0)

	fmt.Println("连接多个切片")
	fmt.Println("从go1.22版本开始才可以使用")
	sl94 := []int{1, 2, 3}
	sl95 := []int{4, 5, 6}
	sl96 := make([]int, 3, 6)
	sl96 = slices.Replace(sl96, 0, 3, []int{7, 8, 9}...)
	sl97 := make([]int, 3, 7)
	sl97 = slices.Replace(sl97, 0, 3, []int{7, 8, 9}...)
	sl98 := make([]int, 3, 8)
	sl98 = slices.Replace(sl98, 0, 3, []int{7, 8, 9}...)
	sl99 := slices.Concat(sl94, sl95, sl96)
	sl100 := slices.Concat(sl94, sl95, sl97)
	sl101 := slices.Concat(sl94, sl95, sl98)
	mfp.PrintFmtValWithLC("sl94", sl94, verbs)
	mfp.PrintFmtValWithLC("sl95", sl95, verbs)
	mfp.PrintFmtValWithLC("sl96", sl96, verbs)
	mfp.PrintFmtValWithLC("sl97", sl97, verbs)
	mfp.PrintFmtValWithLC("sl98", sl98, verbs)
	mfp.PrintFmtValWithLC("sl99", sl99, verbs)
	mfp.PrintFmtValWithLC("sl100", sl100, verbs)
	mfp.PrintFmtValWithLC("sl101", sl101, verbs)
	fmt.Println("去重")
	fmt.Println("使用slices.Compact函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl102 := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 8, 8, 1, 2, 3, 4, 5, 8}
	mfp.PrintFmtValWithLC("1 sl102", sl102, verbs)
	sl102 = slices.Compact(sl102)
	mfp.PrintFmtValWithLC("2 sl102", sl102, verbs)

	fmt.Println("使用slices.CompactFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl103 := []string{"bob", "Bob", "alice", "Vera", "VERA"}
	mfp.PrintFmtValWithLC("1 sl103", sl103, verbs)
	sl103 = slices.CompactFunc(sl103, func(a, b string) bool {
		return strings.ToLower(a) == strings.ToLower(b)
	})
	mfp.PrintFmtValWithLC("2 sl103", sl103, verbs)

	fmt.Println("判断是否已经排序")
	fmt.Println("使用slices.IsSorted函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl104 := []int{1, 2, 3}
	sl105 := []int{1, 3, 2}
	fmt.Println("sl104已排序？-> ", slices.IsSorted(sl104))
	fmt.Println("sl105已排序？-> ", slices.IsSorted(sl105))

	fmt.Println("使用slices.IsSortedFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl106 := []string{"alice", "Bob", "VERA"}
	fmt.Println("sl106已排序？-> ", slices.IsSortedFunc(sl106, func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	}))

	fmt.Println("使用slices.Clone函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl107 := make([]int, 3, 6)
	sl107 = slices.Replace(sl107, 0, 3, []int{1, 2, 3}...)
	mfp.PrintFmtValWithLC("1 sl107", sl107, verbs)
	sl108 := slices.Clone(sl107)
	mfp.PrintFmtValWithLC("2 sl108", sl108, verbs)
	sl108[0] = 11
	fmt.Println("sl108[0] = 11 之后")
	mfp.PrintFmtValWithLC("3 sl107", sl107, verbs)
	mfp.PrintFmtValWithLC("4 sl108", sl108, verbs)

	sl107[1] = 22
	fmt.Println("sl107[1] = 22 之后")
	mfp.PrintFmtValWithLC("5 sl107", sl107, verbs)
	mfp.PrintFmtValWithLC("6 sl108", sl108, verbs)

	fmt.Println("插入")
	fmt.Println("使用slices.Insert函数")
	fmt.Println("从go1.21版本开始才可以使用")

	sl109 := make([]int, 2, 3)
	sl109 = slices.Replace(sl109, 0, 2, []int{1, 2}...)
	mfp.PrintFmtValWithLC("1 sl109", sl109, verbs)
	sl109 = slices.Insert(sl109, 0, []int{11, 22, 33}...)
	mfp.PrintFmtValWithLC("2 sl109", sl109, verbs)

	sl110 := make([]int, 2)
	sl110 = slices.Replace(sl110, 0, 2, []int{1, 2}...)
	mfp.PrintFmtValWithLC("1 sl110", sl110, verbs)
	sl110 = slices.Insert(sl110, 0, []int{11, 22}...)
	mfp.PrintFmtValWithLC("2 sl110", sl110, verbs)

	sl111 := make([]int, 2)
	sl111 = slices.Replace(sl111, 0, 2, []int{1, 2}...)
	mfp.PrintFmtValWithLC("1 sl111", sl111, verbs)
	sl111 = slices.Insert(sl111, 0, []int{11, 22, 33}...)
	mfp.PrintFmtValWithLC("2 sl111", sl111, verbs)

	fmt.Println("使用slices.Replace函数")
	fmt.Println("从go1.21版本开始才可以使用")

	sl112 := make([]int, 2, 3)
	sl112 = slices.Replace(sl112, 0, 2, []int{1, 2}...)
	mfp.PrintFmtValWithLC("1 sl112", sl112, verbs)
	sl112 = slices.Replace(sl112, 0, 0, 11)
	mfp.PrintFmtValWithLC("2 sl112", sl112, verbs)
	sl112 = slices.Replace(sl112, 0, 0, 111)
	mfp.PrintFmtValWithLC("3 sl112", sl112, verbs)

	fmt.Println("获取索引")
	fmt.Println("使用slices.Index函数")
	fmt.Println("从go1.21版本开始才可以使用")

	sl113 := []string{"hello", "golang", "China", "World"}
	fmt.Println("golang在sl113中的索引是 ", slices.Index(sl113, "golang"))
	fmt.Println("China在sl113中的索引是 ", slices.Index(sl113, "China"))
	fmt.Println("xyz在sl113中的索引是 ", slices.Index(sl113, "xyz"))

	fmt.Println("使用slices.BinarySearch")
	fmt.Println("从go1.21版本开始才可以使用")
	sl114 := []string{"hello", "golang", "China", "World"}
	fmt.Println("未排序的sl114")
	mfp.PrintFmtValWithLC("1 sl114", sl114, verbs)
	i114, b114 := slices.BinarySearch(sl114, "golang")
	fmt.Printf("golang 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
	i114, b114 = slices.BinarySearch(sl114, "China")
	fmt.Printf("China 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
	i114, b114 = slices.BinarySearch(sl114, "xyz")
	fmt.Printf("xyz 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
	mfp.PrintHr()
	fmt.Println("已排序的sl114")
	slices.Sort(sl114)
	mfp.PrintFmtValWithLC("2 sl114", sl114, verbs)
	i114, b114 = slices.BinarySearch(sl114, "golang")
	fmt.Printf("golang 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
	i114, b114 = slices.BinarySearch(sl114, "China")
	fmt.Printf("China 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
	i114, b114 = slices.BinarySearch(sl114, "xyz")
	fmt.Printf("xyz 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)

	fmt.Println("格式化动词")
	newVerbs := []string{"T", "%v", "+v", "#v"}
	sl115 := []int{1, 2, 3}
	sl116 := []float32{1.1, 2.2, 3.3}
	sl117 := []string{"A", "B", "C"}
	sl118 := []Person{{"Alice", 12}, {"Bob", 28}}
	mfp.PrintFmtValWithLC("sl115", sl115, newVerbs)
	mfp.PrintFmtValWithLC("sl116", sl116, newVerbs)
	mfp.PrintFmtValWithLC("sl117", sl117, newVerbs)
	mfp.PrintFmtValWithLC("sl118", sl118, newVerbs)
}

func reverseSlice(slice []int) {
	length := len(slice)
	for i := 0; i < length/2; i++ {
		j := length - 1 - i
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func findMin[T1, T2 cmp.Ordered](minK T1, minV T2, src []T2) (T1, T2) {
	for k, v := range src {
		if minV > v {
			minK = T1(k)
			minV = v
		}
	}
	return minK, minV
}
