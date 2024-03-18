package main

import (
	"cmp"
	"fmt"
	"github.com/before80/utils/mfp"
	"math"
	"slices"
	"strconv"
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
	slSrc43 := []int{1, 2, 3}
	mfp.PrintFmtValWithLC("slSrc43", slSrc43, verbs)
	slDst44 := make([]int, len(slSrc43))
	mfp.PrintFmtValWithLC("slDst44", slDst44, verbs)

	copy(slDst44, slSrc43) // func copy(dst []Type, src []Type) int
	fmt.Println("使用copy函数")
	slDst44[0] = 11
	fmt.Println("slDst44[0] = 11 之后")
	mfp.PrintFmtValWithLC("slDst43", slSrc43, verbs)
	mfp.PrintFmtValWithLC("slDst44", slDst44, verbs)
	slSrc43[1] = 22
	fmt.Println("slSrc43[1] = 22 之后")
	mfp.PrintFmtValWithLC("slDst43", slSrc43, verbs)
	mfp.PrintFmtValWithLC("slDst44", slDst44, verbs)

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
	fmt.Println("sl48 == sl49 -> ", slices.Equal(sl48, sl49))
	fmt.Println("sl48 == sl50 -> ", slices.Equal(sl48, sl50))
	fmt.Println("sl48 == sl51 -> ", slices.Equal(sl48, sl51))

	fmt.Println("使用slices.EqualFunc函数")
	fmt.Println("从go1.21版本开始才可以使用")
	sl52 := []int{1, 15, 8}
	sl53 := []int{1, 15, 8}
	sl54 := []int{11, 15, 8}
	sl55 := []string{"01", "0x0f", "0o10"}

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
	sl74 = slices.Replace(sl74, 0, 0, 111)
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
