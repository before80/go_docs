package main

import (
	"fmt"
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

	copy(slDst44, slSrc43)
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

	fmt.Println("排序")

}
