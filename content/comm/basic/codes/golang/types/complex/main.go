package main

import (
	"fmt"
	"github.com/before80/utils/mfp"
)

// 全局声明（这里的全局应该说是 包级别的全局，即相同包名（连路径都相同的包名）下，不可声明两个相同名称的全局变量）
var gc641 = complex(float32(1), float32(2)) // 注意这里需要使用byte()进行类型转换，这里的byte()并非函数，仅仅是一个类型+一对()而已
var gc642 complex64 = 1 + 2i
var gc1281 = complex(float64(1), float64(2))
var gc1282 = complex(1, 2)
var gc1283 complex128 = 1 + 2i

var verbs = []string{"T", "v", "+v", "#v"}

func init() {
	fmt.Println("---init 修改前---")
	mfp.PrintFmtVal("全局变量 gc641", gc641, verbs)
	mfp.PrintFmtVal("全局变量 gc642", gc642, verbs)
	mfp.PrintFmtVal("全局变量 gc1281", gc1281, verbs)
	mfp.PrintFmtVal("全局变量 gc1282", gc1282, verbs)
	mfp.PrintFmtVal("全局变量 gc1283", gc1283, verbs)
	// 对部分全局变量进行修改
	gc641 = complex(float32(1.1), float32(2.2))
	gc1281 = complex(1.1, 2.2)
}

func main() {
	fmt.Println("---init 执行完成后---")
	mfp.PrintFmtVal("全局变量 gc641", gc641, verbs)
	mfp.PrintFmtVal("全局变量 gc1281", gc1281, verbs)
	fmt.Println("---局部变量---")
	fmt.Println("---complex64---")
	// 声明方式1
	var c641 complex64 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 false
	mfp.PrintFmtVal("声明方式1 c641", c641, verbs)
	// 赋值
	c641 = complex(float32(1), float32(2))
	mfp.PrintFmtVal("赋值后", c641, verbs)
	c641 = complex(float32(1.1), float32(2.2))
	mfp.PrintFmtVal("赋值后", c641, verbs)

	// 声明方式2
	var c642 complex64 = 1 + 2i
	mfp.PrintFmtVal("声明方式2 c642", c642, verbs)

	//短变量声明，仅用于局部变量
	c643 := complex(float32(1), float32(2))
	mfp.PrintFmtVal("声明方式3（短变量声明） c643", c643, verbs)

	x6431 := imag(c643)
	mfp.PrintFmtVal("调用imag函数 x6431", x6431, verbs)
	x6432 := real(c643)
	mfp.PrintFmtVal("调用real函数 x6432", x6432, verbs)

	c644 := complex(float32(1), float32(2))
	_ = c644 //这一赋值语句，仅仅是用于防止‘定义了但未使用的变量’报错

	fmt.Println("---complex128---")
	// 声明方式1
	var c1281 complex128 // 看着是仅声明，实际上已经存在隐式给该变量赋予了该类型的零值，即 false
	mfp.PrintFmtVal("声明方式1 c1281", c1281, verbs)
	// 赋值
	c1281 = complex(1, 2)
	mfp.PrintFmtVal("赋值后", c1281, verbs)
	c1281 = complex(float64(1.1), float64(2.2))
	mfp.PrintFmtVal("赋值后", c1281, verbs)

	// 声明方式2
	var c1282 complex128 = 1 + 2i
	mfp.PrintFmtVal("声明方式2 c1282", c1282, verbs)

	//短变量声明，仅用于局部变量
	c1283 := complex(1, 2)
	mfp.PrintFmtVal("声明方式3（短变量声明） c1283", c1283, verbs)

	c1284 := complex(1, 2)
	_ = c1284
}
