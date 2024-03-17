package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gs1 = "Hello World"
var gs2 string = "勇敢前行"

var verbs = []string{"T", "v", "+v", "#v", "s", "q", "+q", "#q", "x", "X"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gs1", gs1, verbs)
	mfp.PrintFmtVal("全局变量 gs2", gs2, verbs)
	// 对部分全局变量进行修改
	gs1 = "Hello 中国！"
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gs1", gs1, verbs)
	fmt.Println("---局部变量---")
	// 声明方式1
	var s1 string // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 ""
	mfp.PrintFmtVal("声明方式1 s1", s1, verbs)
	// 赋值
	s1 = "你好"
	mfp.PrintFmtVal("赋值后", s1, verbs)

	s1 = "Hello 你好"
	mfp.PrintFmtVal("赋值后", s1, verbs)

	// 声明方式2
	var b2 = "真诚勤勇"
	mfp.PrintFmtVal("声明方式2 b2", b2, verbs)
	mfp.PrintFmtVal("声明方式2 b2", b2, []string{"x"})
	mfp.PrintFmtVal("声明方式2 b2", b2, []string{"x", "X"})
	mfp.PrintFmtVal("声明方式2 b2", b2, []string{"#q", "x", "X"})

	//短变量声明，仅用于局部变量
	b3 := "Welcome to Go"
	mfp.PrintFmtVal("声明方式3 b3", b3, verbs)

	b4 := "Nice to meet you!很高兴见到你！"
	_ = b4 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错
}
