package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gf321 = float32(1.1) // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gf322 float32 = float32(2.2)
var gf641 = 1.1
var gf642 float64 = 2.2

var verbs = []string{"T", "v", "+v", "#v", "b", "e", "E", "f", "F", "g", "G", "x", "X"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gf321", gf321, verbs)
	mfp.PrintFmtVal("全局变量 gf322", gf322, verbs)
	mfp.PrintFmtVal("全局变量 gf641", gf641, verbs)
	mfp.PrintFmtVal("全局变量 gf642", gf642, verbs)
	// 对部分全局变量进行修改
	gf321 = 1234567890.123456789
	gf641 = 1234567890.123456789
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gf321", gf321, verbs)
	mfp.PrintFmtVal("全局变量 gf641", gf641, verbs)
	fmt.Println("---局部变量---")
	fmt.Println("---float32---")
	// 声明方式1
	var f321 float32 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 f321", f321, verbs)
	// 赋值
	f321 = 12
	mfp.PrintFmtVal("赋值后", f321, verbs)
	f321 = 1234567890.123456789
	mfp.PrintFmtVal("赋值后", f321, verbs)

	// 声明方式2
	var f322 float32 = 1.1
	mfp.PrintFmtVal("声明方式2 f322", f322, verbs)

	//短变量声明，仅用于局部变量
	f323 := float32(2.2)
	mfp.PrintFmtVal("声明方式3（短变量声明） f323", f323, verbs)

	f324 := float32(1.1)
	_ = f324 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错

	fmt.Println("---float64---")
	// 声明方式1
	var f641 float64 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 f641", f641, verbs)
	// 赋值
	f641 = 13
	mfp.PrintFmtVal("赋值后", f641, verbs)
	f641 = 1234567890.123456789
	mfp.PrintFmtVal("赋值后", f641, verbs)

	// 声明方式2
	var f642 float64 = 1.1
	mfp.PrintFmtVal("声明方式2 f642", f642, verbs)

	//短变量声明，仅用于局部变量
	f643 := 2.2
	mfp.PrintFmtVal("声明方式3（短变量声明） f643", f643, verbs)

	f644 := 1.1
	_ = f644
}
