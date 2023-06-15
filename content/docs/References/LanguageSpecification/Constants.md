+++
title = "常量"
date = 2023-05-17T09:59:21+08:00
weight = 5
description = ""
isCJKLanguage = true
draft = false
+++
## Constants 常量

> 原文：[https://go.dev/ref/spec#Constants](https://go.dev/ref/spec#Constants)

​	有布尔常量、符文常量、整数常量、浮点常量、复数常量和字符串常量。符文、整数、浮点和复数常量统称为`数值常量`。

​	常量值由一个[符文字面量](../LexicalElements#rune-literals-符文字面量)、[整数字面量](../LexicalElements#imaginary-literals-虚数字面量)、[浮点数字面量](../LexicalElements#floating-point-literals-浮点数字面量)、[虚数字面量](../LexicalElements#imaginary-literals-虚数字面量)或[字符串字面量](../LexicalElements#string-literals-字符串字面量)，表示常量的标识符，[常量表达式](../Expressions#constant-expressions-常量表达式)，结果为常量的[转换](../Expressions#conversions-转换)，或一些内置函数的结果值表示，如`unsafe.Sizeof`应用于某些值，`cap`或`len`应用于一些表达式，`real`和`imag`应用于复数常量，`complex`应用于数值常量。布尔真值由预先声明的常数`true`和`false`表示。预先声明的标识符`iota`表示一个整数常量。

> 个人注释
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"unsafe"
> )
> 
> var arr = [3]int{1, 2, 3}
> var sli = []int{1, 2, 3}
> 
> const r = 'r'
> const i1 = 1
> const i2 = 11
> const f = 1.2
> const x = 1 + 2.3i
> const s = "This is a string."
> const s1 = s
> const (
> 	iotaNum1 = iota + 1
> 	iotaNum2 = iota + 2
> )
> 
> const fi = f + i1
> const i2f = float64(i2)
> 
> const unsafeSizeOfS = unsafe.Sizeof(s)
> 
> const builtIn1 = len(s)
> 
> // const builtIn2 = len(sli) // len(sli) (value of type int) is not constant
> // const buildIn3 = cap(sli) // cap(sli) (value of type int) is not constant
> const buildIn4 = cap(arr)
> const ir = real(1 + 2i)
> const ii = imag(1 + 2i)
> const ic = complex(1, 2)
> const isTrue = true
> const isFalse = false
> 
> func main() {
> 	fmt.Printf("r =%v,其类型是%T\n", r, r)                                     // r =114,其类型是int32
> 	fmt.Printf("i1 =%v,其类型是%T\n", i1, i1)                                  // i1 =1,其类型是int
> 	fmt.Printf("i2 =%v,其类型是%T\n", i2, i2)                                  // i2 =11,其类型是int
> 	fmt.Printf("f =%v,其类型是%T\n", f, f)                                     // f =1.2,其类型是float64
> 	fmt.Printf("x =%v,其类型是%T\n", x, x)                                     // x =(1+2.3i),其类型是complex128
> 	fmt.Printf("s =%v,其类型是%T\n", s, s)                                     // s =This is a string.,其类型是string
> 	fmt.Printf("s1 =%v,其类型是%T\n", s1, s1)                                  // s1 =This is a string.,其类型是string
> 	fmt.Printf("iotaNum1 =%v,其类型是%T\n", iotaNum1, iotaNum1)                // iotaNum1 =1,其类型是int
> 	fmt.Printf("iotaNum2 =%v,其类型是%T\n", iotaNum2, iotaNum2)                // iotaNum2 =3,其类型是int
> 	fmt.Printf("fi =%v,其类型是%T\n", fi, fi)                                  // fi =2.2,其类型是float64
> 	fmt.Printf("i2f =%v,其类型是%T\n", i2f, i2f)                               // i2f =11,其类型是float64
> 	fmt.Printf("unsafeSizeOfS =%v,其类型是%T\n", unsafeSizeOfS, unsafeSizeOfS) // unsafeSizeOfS =16,其类型是uintptr
> 	fmt.Printf("builtIn1 =%v,其类型是%T\n", builtIn1, builtIn1)                // builtIn1 =17,其类型是int
> 	//fmt.Printf("builtIn2 =%v,其类型是%T\n", builtIn2, builtIn2)
> 	//fmt.Printf("buildIn3 =%v,其类型是%T\n", buildIn3, buildIn3)
> 	fmt.Printf("buildIn4 =%v,其类型是%T\n", buildIn4, buildIn4) // buildIn4 =3,其类型是int
> 	fmt.Printf("ir =%v,其类型是%T\n", ir, ir)                   // ir =1,其类型是float64
> 	fmt.Printf("ii =%v,其类型是%T\n", ii, ii)                   // ii =2,其类型是float64
> 	fmt.Printf("ic =%v,其类型是%T\n", ic, ic)                   // ic =(1+2i),其类型是complex128
> 	fmt.Printf("isTrue =%v,其类型是%T\n", isTrue, isTrue)       // isTrue =true,其类型是bool
> 	fmt.Printf("isFalse =%v,其类型是%T\n", isFalse, isFalse)    // isFalse =false,其类型是bool
> }
> 
> ```
>
> 这么说，难道常量值的类型不能是数组、切片、map、channel、结构体？
>
> ```go
> package main
> 
> import "fmt"
> 
> const s1 = [3]int{1, 2, 3}                 // [3]int{…} (value of type [3]int) is not constant
> const s2 = []int{1, 2, 3}                  // []int{…} (value of type []int) is not constant
> const m1 = map[string]int{"a": 1, "b": 2}  // map[string]int{…} (value of type map[string]int) is not constant
> const c <-chan int = 2                     // invalid constant type <-chan int
> const s3 = struct{ Name string }{"zlongx"} // struct{Name string}{…} (value of type struct{Name string}) is not constant
> 
> func main() {
> 	fmt.Println(s1)
> 	fmt.Println(s2)
> 	fmt.Println(m1)
> 	fmt.Println(c)
> 	fmt.Println(s3)
> }
> 
> ```
>
> 确实不能为数组、切片、map、结构体、channel！
>
> 那常量的值可以是指针吗？
>
> ```go
> package main
> 
> import "unsafe"
> 
> const s1 = "This is a string"
> const s2 string = "This is a string"
> 
> var s3 = "This is a string"
> 
> const p1 = &s1                 // invalid operation: cannot take address of s1 (untyped string constant "This is a string")
> const p2 = uintptr(&s1)        // invalid operation: cannot take address of s1 (untyped string constant "This is a string")
> const p3 = unsafe.Pointer(&s1) // nvalid operation: cannot take address of s1 (untyped string constant "This is a string")
> 
> const p4 = &s2                 // invalid operation: cannot take address of s2 (constant "This is a string" of type string)
> const p5 = uintptr(&s2)        // invalid operation: cannot take address of s2 (constant "This is a string" of type string)
> const p6 = unsafe.Pointer(&s2) // invalid operation: cannot take address of s2 (constant "This is a string" of type string)
> 
> const p7 = &s3          // &s3 (value of type *string) is not constant
> const p8 = uintptr(&s3) // cannot convert &s3 (value of type *string) to type uintptr
> const p9 = unsafe.Pointer(&s3) // unsafe.Pointer(&s3) (value of type unsafe.Pointer) is not constant
> 
> func main() {
> }
> 
> ```
>
> 看来常量的值也是不能为指针！
>
> > ​	常量在**编译时**就确定了它们的值，并且在程序运行期间是不可修改的。**指针类型是动态的**，它们包含了变量的内存地址，而这个地址是在运行时确定的。因此，指针类型的值是不适合用作常量的。
>
> 常量值可以是变量吗？
>
> ```go
> package main
> 
> var s = "This is a string"
> 
> const ss = s // s (variable of type string) is not constant
> 
> func main() {
> }
> 
> ```
>
> 明显，常量的值不能为变量！

​	通常，复数常量是[常量表达式](../Expressions#constant-expressions-常量表达式)的一种形式，将在该节中讨论。

​	数值常量表示任意精度的精确值，不会溢出。因此，不存在表示IEEE-754负零、无穷大和非数字值的常量。

> 个人注释
>
> 关于“数值常量表示任意精度的精确值，不会溢出”，还是不能理解！—— 请查看下面的**实施限制**！！！
>
> ```go
> package main
> 
> import "fmt"
> 
> const f1 = 1.123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
> 
> var f2 = f1
> 
> // constant overflow
> //const i1 = 1123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
> 
> //var i2 = i1
> 
> func main() {
> 	fmt.Printf("f1=%v，其类型是%T\n", f1, f1) // f1=1.1234567890123457，其类型是float64
> 	fmt.Printf("f2=%v，其类型是%T\n", f2, f2) // f2=1.1234567890123457，其类型是float64
> 
> 	fmt.Printf("f1=%.20f，其类型是%T\n", f1, f1)
> 	fmt.Printf("f1=%.25f，其类型是%T\n", f1, f1)
> 	fmt.Printf("f1=%.30f，其类型是%T\n", f1, f1)
> 	fmt.Printf("f1=%.35f，其类型是%T\n", f1, f1)
> 	fmt.Printf("f1=%.40f，其类型是%T\n", f1, f1)
> 	fmt.Printf("f1=%.45f，其类型是%T\n", f1, f1)
> 	fmt.Printf("f1=%.50f，其类型是%T\n", f1, f1)
> 	fmt.Printf("f1=%.55f，其类型是%T\n", f1, f1)
> 	fmt.Printf("f1=%.60f，其类型是%T\n", f1, f1)
> 
> 	fmt.Printf("f2=%.20f，其类型是%T\n", f2, f2)
> 	fmt.Printf("f2=%.25f，其类型是%T\n", f2, f2)
> 	fmt.Printf("f2=%.30f，其类型是%T\n", f2, f2)
> 	fmt.Printf("f2=%.35f，其类型是%T\n", f2, f2)
> 	fmt.Printf("f2=%.40f，其类型是%T\n", f2, f2)
> 	fmt.Printf("f2=%.45f，其类型是%T\n", f2, f2)
> 	fmt.Printf("f2=%.50f，其类型是%T\n", f2, f2)
> 	fmt.Printf("f2=%.55f，其类型是%T\n", f2, f2)
> 	fmt.Printf("f2=%.60f，其类型是%T\n", f2, f2)
> 
> 	//fmt.Printf("i1=%v，其类型是%T\n", i1, i1)
> 	//fmt.Printf("i2=%v，其类型是%T\n", i2, i2)
> }
> f1=1.1234567890123457，其类型是float64
> f2=1.1234567890123457，其类型是float64
> f1=1.12345678901234569125，其类型是float64
> f1=1.1234567890123456912476740，其类型是float64
> f1=1.123456789012345691247674039914，其类型是float64
> f1=1.12345678901234569124767403991427273，其类型是float64
> f1=1.1234567890123456912476740399142727255821，其类型是float64
> f1=1.123456789012345691247674039914272725582122803，其类型是float64
> f1=1.12345678901234569124767403991427272558212280273438，其类型是float64
> f1=1.1234567890123456912476740399142727255821228027343750000，其类型是float64
> f1=1.123456789012345691247674039914272725582122802734375000000000，其类型是float64
> f2=1.12345678901234569125，其类型是float64
> f2=1.1234567890123456912476740，其类型是float64
> f2=1.123456789012345691247674039914，其类型是float64
> f2=1.12345678901234569124767403991427273，其类型是float64
> f2=1.1234567890123456912476740399142727255821，其类型是float64
> f2=1.123456789012345691247674039914272725582122803，其类型是float64
> f2=1.12345678901234569124767403991427272558212280273438，其类型是float64
> f2=1.1234567890123456912476740399142727255821228027343750000，其类型是float64
> f2=1.123456789012345691247674039914272725582122802734375000000000，其类型是float64
> 
> 
> ```
>
> 

​	常量可以是[有类型的](../Types)的或无类型的。`字面常量`、`true`、`false`、`iota`，以及某些只包含无类型的常量操作数的[常量表达式](../Expressions#constant-expressions-常量表达式)是无类型的。

​	常量可以通过[常量声明](../DeclarationsAndScope#constant-declarations-常量声明)或[转换](../Expressions#conversions-转换)显式地给出类型，也可以在[变量声明](../DeclarationsAndScope#variable-declarations-变量声明)、[赋值语句](../Statements#assignment-statements-赋值语句) 、作为[表达式](../Expressions)的操作数时，隐式赋予类型。如果常量值不能被[表示](../PropertiesOfTypesAndValues#representability-可表示性)为相应类型的值，那就是一个错误。如果类型是一个[类型参数](../DeclarationsAndScope#type-parameter-declarations-类型参数声明)，常量将被转换为类型参数的一个非常量值。

> 个人注释
>
> 这里提到的隐式赋予类型怎么解释？请看以下给出的示例
>
> ```go
> package main
> 
> import "fmt"
> 
> const A = 1
> const B = 2
> const C uint8 = 3
> const D = 4
> 
> func main() {
> 	var a int8 = 1
> 	fmt.Printf("A的类型是：%T\n", A)        //A的类型是：int
> 	fmt.Printf("a的类型是：%T\n", a)        //a的类型是：int8
> 	fmt.Println("a + A = ", a+A)       //a + A =  2
> 	fmt.Printf("a + A后的类型是：%T\n", a+A) //a + A后的类型是：int8
> 
> 	var b uint8 = 2
> 	fmt.Printf("B的类型是：%T\n", B)        //B的类型是：int
> 	fmt.Printf("b的类型是：%T\n", b)        //b的类型是：uint8
> 	fmt.Println("b + B = ", b+B)       //b + B =  4
> 	fmt.Printf("b + B后的类型是：%T\n", b+B) //b + B后的类型是：uint8
> 
> 	var c int = 3
> 	fmt.Printf("C的类型是：%T\n", C) //C的类型是：uint8
> 	fmt.Printf("c的类型是：%T\n", c) //c的类型是：int
> 	//fmt.Println("c + C = ", c+C)       // invalid operation: c + C (mismatched types int and uint8)
> 	//fmt.Printf("c + C后的类型是：%T\n", c+C) // invalid operation: c + C (mismatched types int and uint8)
> 
> 	var d float64 = 1.2
> 	fmt.Printf("D的类型是：%T\n", D)        //D的类型是：int
> 	fmt.Printf("d的类型是：%T\n", d)        //d的类型是：float64
> 	fmt.Println("d + D = ", d+D)       // d + D =  5.2
> 	fmt.Printf("d + D后的类型是：%T\n", d+D) // d + D后的类型是：float64
> 
> }
> 
> ```
>
> 我们知道算数运算符左右两个操作数的类型必须一致，才能进行运算（可参见 [数值型]({{< ref "/docs/References/LanguageSpecification/Types#numeric-types-数值型">}})的最后一段话：当不同的数值类型在表达式或赋值中混合使用时，需要进行显示转换。）。
>
> 以上示例a + A、b + B 、d + D中，+ 运算符两边的类型分明是不一致的（特别是d + D），但却可以进行运算，可见就是隐式赋予了无类型常量（**该常量声明时没有给出明确类型，采用了默认类型**）以类型，这应该是编译时就处理好的吧。TODO 待找出出处。

​	一个无类型常量有一个默认的类型，该类型是在需要类型化值的上下文中隐式转换为的类型，例如，在一个[短变量声明](../DeclarationsAndScope#short-variable-declarations-短变量声明)中，如`i := 0`，没有明确的类型。无类型常量的默认类型分别是`bool`, `rune`, `int`, `float64`, `complex128`或`string`，具体取决于它是一个布尔型常量、rune型常量、整数型常量、浮点型常量、复数型常量还是字符串型常量。

​	实现限制：尽管数值常量在语言中具有任意的精度，但编译器可以使用有限精度的内部表示法来实现它们。也就是说，每个实现都必须：

- 用至少256位来表示整数常量。
- 用至少256位的尾数和至少16位的有符号二进制指数来表示浮点常量，包括复数常量的对应部分。
- 如果不能精确表示一个整数常量，则给出一个错误。
- 如果由于溢出而无法表示一个浮点常量或复数常量，则给出一个错误。
- 如果由于精度的限制，无法表示一个浮点常量或复数常量，则四舍五入到最接近的可表示常量。

​	这些要求既适用于字面常量，也适用于[常量表达式](../Expressions#constant-expressions-常量表达式)的计算结果。