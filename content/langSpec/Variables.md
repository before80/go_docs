+++
title = "变量"
date = 2023-05-17T09:59:21+08:00
weight = 6
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## Variables 变量

> 原文：[https://go.dev/ref/spec#Variables ](https://go.dev/ref/spec#Variables )

A variable is a storage location for holding a *value*. The set of permissible values is determined by the variable's *[type](https://go.dev/ref/spec#Types)*.

​	变量是用于保存值的存储位置。允许值的集合是由变量的[类型](../Types)决定的。

A [variable declaration](https://go.dev/ref/spec#Variable_declarations) or, for function parameters and results, the signature of a [function declaration](https://go.dev/ref/spec#Function_declarations) or [function literal](https://go.dev/ref/spec#Function_literals) reserves storage for a named variable. Calling the built-in function [`new`](https://go.dev/ref/spec#Allocation) or taking the address of a [composite literal](https://go.dev/ref/spec#Composite_literals) allocates storage for a variable at run time. Such an anonymous variable is referred to via a (possibly implicit) [pointer indirection](https://go.dev/ref/spec#Address_operators).

​	[变量声明](../DeclarationsAndScope#variable-declarations-变量声明) 、函数参数和结果、[函数声明](../DeclarationsAndScope#function-declarations-函数声明)或[函数字面量](../Expressions#function-literals-函数字面量)的签名为指定的变量保留存储空间。调用内置函数`new`或获取[复合字面量](../Expressions#composite-literals-复合字面量)的地址会在运行时为变量分配存储空间。这样的匿名变量是通过(可能是隐式的)[指针间接](../Expressions#address-operators-地址运算符)引用的。

*Structured* variables of [array](https://go.dev/ref/spec#Array_types), [slice](https://go.dev/ref/spec#Slice_types), and [struct](https://go.dev/ref/spec#Struct_types) types have elements and fields that may be [addressed](https://go.dev/ref/spec#Address_operators) individually. Each such element acts like a variable.

​	[数组](../Types#array-types-数组型)、[切片](../Types#slice-types-切片型)和[结构体](../Types#struct-types-结构体型)等类型的结构化变量具有可以被单独[寻址](../Expressions#address-operators-地址运算符)的元素和字段。每个这样的元素都像一个变量。

> 个人注释
>
> ```go
> package main
> 
> import "fmt"
> 
> type St struct {
> 	Name string
> 	Age  int
> }
> 
> func main() {
> 	var arr = [...]int{1, 2, 3}	
> 	fmt.Println(&arr[0]) // 0xc000010120
> 	fmt.Println(&arr[1]) //0xc000010128
> 	fmt.Println(&arr[2]) //0xc000010130
>     
>     var sli = []int{1, 2, 3}
> 	fmt.Println(&sli[0]) //0xc000010138
> 	fmt.Println(&sli[1]) //0xc000010140
> 	fmt.Println(&sli[2]) //0xc000010148
> 
> 	var st1 = St{"zlongx", 32}
> 	fmt.Println(&st1.Name) //0xc000008078
> 	fmt.Println(&st1.Age)  //0xc000008088
> }
> 
> ```
>
> 

The *static type* (or just *type*) of a variable is the type given in its declaration, the type provided in the `new` call or composite literal, or the type of an element of a structured variable. Variables of interface type also have a distinct *dynamic type*, which is the (non-interface) type of the value assigned to the variable at run time (unless the value is the predeclared identifier `nil`, which has no type). The dynamic type may vary during execution but values stored in interface variables are always [assignable](https://go.dev/ref/spec#Assignability) to the static type of the variable.

​	变量的静态类型（或简称为类型）是在其声明中指定的类型、在`new`调用或复合字面量中提供的类型、或是结构化变量的元素的类型。接口类型的变量还具有独特的动态类型，它是在运行时分配给变量的值的（非接口）类型（除非该值是预先声明的标识符`nil`，它没有类型）。在执行过程中，动态类型可能会发生变化，但是存储在接口变量中的值始终可以赋给变量的静态类型。

```go 
var x interface{}  // x is nil and has static type interface{}  => x 为 nil，具有静态类型 interface{}
var v *T           // v has value nil, static type *T =>  v 具有值 nil，静态类型 * T
x = 42             // x has value 42 and dynamic type int =>  x 具有值42，动态类型 int
x = v              // x has value (*T)(nil) and dynamic type *T => x 具有值(* T)(nil)和动态类型 * T
```

A variable's value is retrieved by referring to the variable in an [expression](https://go.dev/ref/spec#Expressions); it is the most recent value [assigned](https://go.dev/ref/spec#Assignment_statements) to the variable. If a variable has not yet been assigned a value, its value is the [zero value](https://go.dev/ref/spec#The_zero_value) for its type.

​	变量的值是通过在[表达式](../Expressions)中引用该变量来检索的；它是最近[分配](../statements#assignment-statements-赋值语句)给该变量的值。如果一个变量还没有被赋值，它的值就是其类型的[零值](../ProgramInitializationAndExecution#the-zero-value-零值)。