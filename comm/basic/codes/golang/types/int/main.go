package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gi81 = int8(1) // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gi82 int8 = -2
var gi161 = int16(1)
var gi162 int16 = -2
var gi321 = int32(1)
var gi322 int32 = -2
var gi641 = int64(1)
var gi642 int64 = -2
var gi1 = 1
var gi2 int = -2
var verbs = []string{"T", "v", "+v", "#v", "b", "c", "d", "o", "O", "q", "x", "X", "U"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gi81", gi81, verbs)
	mfp.PrintFmtVal("全局变量 gi82", gi82, verbs)
	mfp.PrintFmtVal("全局变量 gi161", gi161, verbs)
	mfp.PrintFmtVal("全局变量 gi162", gi162, verbs)
	mfp.PrintFmtVal("全局变量 gi321", gi321, verbs)
	mfp.PrintFmtVal("全局变量 gi322", gi322, verbs)
	mfp.PrintFmtVal("全局变量 gi641", gi641, verbs)
	mfp.PrintFmtVal("全局变量 gi642", gi642, verbs)
	mfp.PrintFmtVal("全局变量 gi1", gi1, verbs)
	mfp.PrintFmtVal("全局变量 gi2", gi2, verbs)
	// 对部分全局变量进行修改
	gi81 = -12
	gi161 = -12
	gi321 = -12
	gi641 = -12
	gi1 = -12
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gi81", gi81, verbs)
	mfp.PrintFmtVal("全局变量 gi161", gi161, verbs)
	mfp.PrintFmtVal("全局变量 gi321", gi321, verbs)
	mfp.PrintFmtVal("全局变量 gi641", gi641, verbs)
	mfp.PrintFmtVal("全局变量 gi1", gi1, verbs)
	fmt.Println("---局部变量---")
	fmt.Println("---int8---")
	// 声明方式1
	var i81 int8 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i81", i81, verbs)
	// 赋值
	i81 = 1
	mfp.PrintFmtVal("赋值后", i81, verbs)
	i81 = 11
	mfp.PrintFmtVal("赋值后", i81, verbs)

	// 声明方式2
	var i82 int8 = 20
	mfp.PrintFmtVal("声明方式2 i82", i82, verbs)

	//短变量声明，仅用于局部变量
	i83 := int8(30)
	mfp.PrintFmtVal("声明方式3（短变量声明） i83", i83, verbs)

	i84 := int8(40)
	_ = i84 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错

	fmt.Println("---int16---")
	// 声明方式1
	var i161 int16 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i161", i161, verbs)
	// 赋值
	i161 = 1
	mfp.PrintFmtVal("赋值后", i161, verbs)
	i161 = 11
	mfp.PrintFmtVal("赋值后", i161, verbs)

	// 声明方式2
	var i162 int16 = 12
	mfp.PrintFmtVal("声明方式2 i162", i162, verbs)

	//短变量声明，仅用于局部变量
	i163 := int16(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） i163", i163, verbs)

	i164 := int16(1234)
	_ = i164

	fmt.Println("---int32---")
	// 声明方式1
	var i321 int32 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i321", i321, verbs)
	// 赋值
	i321 = 1
	mfp.PrintFmtVal("赋值后", i321, verbs)
	i321 = 11
	mfp.PrintFmtVal("赋值后", i321, verbs)

	// 声明方式2
	var i322 int32 = 12
	mfp.PrintFmtVal("声明方式2 i322", i322, verbs)

	//短变量声明，仅用于局部变量
	i323 := int32(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） i323", i323, verbs)

	i324 := int32(1234)
	_ = i324

	fmt.Println("---int64---")
	// 声明方式1
	var i641 int64 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i641", i641, verbs)
	// 赋值
	i641 = 1
	mfp.PrintFmtVal("赋值后", i641, verbs)
	i641 = 11
	mfp.PrintFmtVal("赋值后", i641, verbs)

	// 声明方式2
	var i642 int64 = 12
	mfp.PrintFmtVal("声明方式2 i642", i642, verbs)

	//短变量声明，仅用于局部变量
	i643 := int64(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） i643", i643, verbs)

	i644 := int64(1234)
	_ = i644

	fmt.Println("---int---")
	// 声明方式1
	var i1 int // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i1", i1, verbs)
	// 赋值
	i1 = 1
	mfp.PrintFmtVal("赋值后", i1, verbs)
	i1 = 11
	mfp.PrintFmtVal("赋值后", i1, verbs)

	// 声明方式2
	var i2 int = 12
	mfp.PrintFmtVal("声明方式2 i2", i2, verbs)

	//短变量声明，仅用于局部变量
	i3 := 123
	mfp.PrintFmtVal("声明方式3（短变量声明） i3", i3, verbs)

	i4 := 1234
	_ = i4
}
