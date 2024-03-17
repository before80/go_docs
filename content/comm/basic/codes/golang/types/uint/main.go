package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gui81 = uint8(1) // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gui82 uint8 = 2
var gui161 = uint16(1)
var gui162 uint16 = 2
var gui321 = uint32(1)
var gui322 uint32 = 2
var gui641 = uint64(1)
var gui642 uint64 = 2
var gui1 = uint(1)
var gui2 uint = 2
var verbs = []string{"T", "v", "+v", "#v", "b", "c", "d", "o", "O", "q", "x", "X", "U"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gui81", gui81, verbs)
	mfp.PrintFmtVal("全局变量 gui82", gui82, verbs)
	mfp.PrintFmtVal("全局变量 gui161", gui161, verbs)
	mfp.PrintFmtVal("全局变量 gui162", gui162, verbs)
	mfp.PrintFmtVal("全局变量 gui321", gui321, verbs)
	mfp.PrintFmtVal("全局变量 gui322", gui322, verbs)
	mfp.PrintFmtVal("全局变量 gui641", gui641, verbs)
	mfp.PrintFmtVal("全局变量 gui642", gui642, verbs)
	mfp.PrintFmtVal("全局变量 gui1", gui1, verbs)
	mfp.PrintFmtVal("全局变量 gui2", gui2, verbs)
	// 对部分全局变量进行修改
	gui81 = 12
	gui161 = 12
	gui321 = 12
	gui641 = 12
	gui1 = 12
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gui81", gui81, verbs)
	mfp.PrintFmtVal("全局变量 gui161", gui161, verbs)
	mfp.PrintFmtVal("全局变量 gui321", gui321, verbs)
	mfp.PrintFmtVal("全局变量 gui641", gui641, verbs)
	mfp.PrintFmtVal("全局变量 gui1", gui1, verbs)
	fmt.Println("---局部变量---")
	fmt.Println("---uint8---")
	// 声明方式1
	var ui81 uint8 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 ui81", ui81, verbs)
	// 赋值
	ui81 = 1
	mfp.PrintFmtVal("赋值后", ui81, verbs)
	ui81 = 11
	mfp.PrintFmtVal("赋值后", ui81, verbs)

	// 声明方式2
	var ui82 uint8 = 20
	mfp.PrintFmtVal("声明方式2 ui82", ui82, verbs)

	//短变量声明，仅用于局部变量
	ui83 := uint8(30)
	mfp.PrintFmtVal("声明方式3（短变量声明） ui83", ui83, verbs)

	ui84 := uint8(40)
	_ = ui84 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错

	fmt.Println("---uint16---")
	// 声明方式1
	var ui161 uint16 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 ui161", ui161, verbs)
	// 赋值
	ui161 = 1
	mfp.PrintFmtVal("赋值后", ui161, verbs)
	ui161 = 11
	mfp.PrintFmtVal("赋值后", ui161, verbs)

	// 声明方式2
	var ui162 uint16 = 12
	mfp.PrintFmtVal("声明方式2 ui162", ui162, verbs)

	//短变量声明，仅用于局部变量
	ui163 := uint16(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） ui163", ui163, verbs)

	ui164 := uint16(1234)
	_ = ui164

	fmt.Println("---uint32---")
	// 声明方式1
	var ui321 uint32 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 ui321", ui321, verbs)
	// 赋值
	ui321 = 1
	mfp.PrintFmtVal("赋值后", ui321, verbs)
	ui321 = 11
	mfp.PrintFmtVal("赋值后", ui321, verbs)

	// 声明方式2
	var ui322 uint32 = 12
	mfp.PrintFmtVal("声明方式2 ui322", ui322, verbs)

	//短变量声明，仅用于局部变量
	ui323 := uint32(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） ui323", ui323, verbs)

	ui324 := uint32(1234)
	_ = ui324

	fmt.Println("---uint64---")
	// 声明方式1
	var ui641 uint64 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 ui641", ui641, verbs)
	// 赋值
	ui641 = 1
	mfp.PrintFmtVal("赋值后", ui641, verbs)
	ui641 = 11
	mfp.PrintFmtVal("赋值后", ui641, verbs)

	// 声明方式2
	var ui642 uint64 = 12
	mfp.PrintFmtVal("声明方式2 ui642", ui642, verbs)

	//短变量声明，仅用于局部变量
	ui643 := uint64(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） ui643", ui643, verbs)

	ui644 := uint64(1234)
	_ = ui644

	fmt.Println("---uint---")
	// 声明方式1
	var i1 uint // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 0
	mfp.PrintFmtVal("声明方式1 i1", i1, verbs)
	// 赋值
	i1 = 1
	mfp.PrintFmtVal("赋值后", i1, verbs)
	i1 = 11
	mfp.PrintFmtVal("赋值后", i1, verbs)

	// 声明方式2
	var i2 uint = 12
	mfp.PrintFmtVal("声明方式2 i2", i2, verbs)

	//短变量声明，仅用于局部变量
	i3 := uint(123)
	mfp.PrintFmtVal("声明方式3（短变量声明） i3", i3, verbs)

	i4 := uint(1234)
	_ = i4
}
