package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gr1 = 'A' // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gr2 rune = 'j'

var verbs = []string{"T", "v", "+v", "#v", "q", "+q", "#q", "c"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gr1", gr1, verbs)
	mfp.PrintFmtVal("全局变量 gr2", gr2, verbs)

	// 对部分全局变量进行修改
	gr1 = 'n'
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gr1", gr1, verbs)
	fmt.Println("---局部变量---")
	// 声明方式1
	var r1 rune // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 false
	mfp.PrintFmtVal("声明方式1 r1", r1, verbs)
	// 赋值
	r1 = 'A'
	mfp.PrintFmtVal("赋值后", r1, verbs)
	r1 = '\a' // 执行时会响铃
	mfp.PrintFmtVal("赋值后", r1, verbs)

	// 声明方式2
	var r2 = '中'
	mfp.PrintFmtVal("声明方式2 r2", r2, verbs)

	//短变量声明，仅用于局部变量
	r3 := '国'
	mfp.PrintFmtVal("声明方式3 r3", r3, verbs)

	r4 := '\x00'
	_ = r4 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错
}
