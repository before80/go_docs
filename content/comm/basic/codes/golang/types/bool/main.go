package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gb1 = true
var gb2 bool = false

var verbs = []string{"T", "v", "+v", "#v", "t"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gb1", gb1, verbs)
	mfp.PrintFmtVal("全局变量 gb2", gb2, verbs)
	// 对部分全局变量进行修改
	gb1 = false
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gb1", gb1, verbs)
	fmt.Println("---局部变量---")
	// 声明方式1
	var b1 bool // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 false
	mfp.PrintFmtVal("声明方式1 b1", b1, verbs)
	// 赋值
	b1 = true
	mfp.PrintFmtVal("赋值后", b1, verbs)

	b1 = false
	mfp.PrintFmtVal("赋值后", b1, verbs)

	// 声明方式2
	var b2 = true
	mfp.PrintFmtVal("声明方式2 b2", b2, verbs)

	//短变量声明，仅用于局部变量
	b3 := true
	mfp.PrintFmtVal("声明方式3 b3", b3, verbs)

	b4 := false
	_ = b4 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错
}
