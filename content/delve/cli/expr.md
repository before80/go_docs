+++
title = "表达式"
date = 2024-12-09T07:59:02+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/cli/expr.md](https://github.com/go-delve/delve/blob/master/Documentation/cli/expr.md)
>
> 收录该文档时间： `2024-12-09T07:59:02+08:00`

# Expressions

Delve can evaluate a subset of go expression language, specifically the following features are supported:

​	Delve 可以评估 Go 表达式语言的一个子集，具体支持以下功能：

- All (binary and unary) on basic types except <-, ++ and --
  - 所有基本类型上的二元和一元操作符，除了 `<-`、`++` 和 `--`

- Comparison operators on any type
  - 对任何类型的比较操作符

- Type casts between numeric types
  - 数值类型之间的类型转换

- Type casts of integer constants into any pointer type and vice versa
  - 整数常量到任何指针类型的类型转换，反之亦然

- Type casts between string, []byte and []rune
  - 字符串、`[]byte` 和 `[]rune` 之间的类型转换

- Struct member access (i.e. `somevar.memberfield`)
  - 结构体成员访问（即 `somevar.memberfield`）

- Slicing and indexing operators on arrays, slices and strings
  - 数组、切片和字符串上的切片和索引操作符

- Map access
  - 映射（map）访问

- Pointer dereference
  - 指针解引用

- Calls to builtin functions: `cap`, `len`, `complex`, `imag` and `real`
  - 调用内建函数：`cap`、`len`、`complex`、`imag` 和 `real`

- Type assertion on interface variables (i.e. `somevar.(concretetype)`)
  - 对接口变量的类型断言（即 `somevar.(concretetype)`）


# 嵌套限制 Nesting limit

When delve evaluates a memory address it will automatically return the value of nested struct members, array and slice items and dereference pointers. However to limit the size of the output evaluation will be limited to two levels deep. Beyond two levels only the address of the item will be returned, for example:

​	当 Delve 评估内存地址时，它将自动返回嵌套的结构体成员、数组和切片项，并进行指针解引用。然而，为了限制输出的大小，评估将限制为最多两级嵌套。超过两级时，仅返回项的地址，例如：

```
(dlv) print c1
main.cstruct {
	pb: *struct main.bstruct {
		a: (*main.astruct)(0xc82000a430),
	},
	sa: []*main.astruct len: 3, cap: 3, [
		*(*main.astruct)(0xc82000a440),
		*(*main.astruct)(0xc82000a450),
		*(*main.astruct)(0xc82000a460),
	],
}
```

To see the contents of the first item of the slice `c1.sa` there are two possibilities:

​	要查看切片 `c1.sa` 的第一个项的内容，有两种可能的方法：

1. Execute `print c1.sa[0]` 执行 `print c1.sa[0]`
2. Use the address directly, executing: `print *(*main.astruct)(0xc82000a440) `直接使用地址，执行：`print *(*main.astruct)(0xc82000a440)`

# 元素限制 Elements limit

For arrays, slices, strings and maps delve will only return a maximum of 64 elements at a time:

​	对于数组、切片、字符串和映射（map），Delve 每次只会返回最多 64 个元素：

```
(dlv) print ba
[]int len: 200, cap: 200, [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,...+136 more]
```

To see more values use the slice operator:

​	要查看更多的值，可以使用切片操作符：

```
(dlv) print ba[64:]
[]int len: 136, cap: 136, [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,...+72 more]
```

For this purpose delve allows use of the slice operator on maps, `m[64:]` will return the key/value pairs of map `m` that follow the first 64 key/value pairs (note that delve iterates over maps using a fixed ordering).

​	为此，Delve 允许在映射（map）上使用切片操作符，`m[64:]` 将返回映射 `m` 中在前 64 个键值对之后的键值对（请注意，Delve 对映射使用固定的迭代顺序）。

These limits can be configured with `max-string-len` and `max-array-values`. See [config](https://github.com/go-delve/delve/tree/master/Documentation/cli#config) for usage.

​	这些限制可以通过 `max-string-len` 和 `max-array-values` 进行配置。有关用法，请参见 [config](https://github.com/go-delve/delve/tree/master/Documentation/cli#config)。

# Interfaces

Interfaces will be printed using the following syntax:

​	接口将使用以下语法打印：

```
<interface name>(<concrete type>) <value>
```

For example:

​	例如：

```
(dlv) p iface1
(dlv) p iface1
interface {}(*struct main.astruct) *{A: 1, B: 2}
(dlv) p iface2
interface {}(*struct string) *"test"
(dlv) p err1
error(*struct main.astruct) *{A: 1, B: 2}
```

To use the contents of an interface variable use a type assertion:

​	要使用接口变量的内容，可以使用类型断言：

```
(dlv) p iface1.(*main.astruct).B
2
```

Or just use the special `.(data)` type assertion:

​	或者直接使用特殊的 `.(data)` 类型断言：

```
(dlv) p iface1.(data).B
2
```

If the contents of the interface variable are a struct or a pointer to struct the fields can also be accessed directly:

​	如果接口变量的内容是结构体或指向结构体的指针，则可以直接访问字段：

```
(dlv) p iface1.B
2
```

# 指定包路径 Specifying package paths

Packages with the same name can be disambiguated by using the full package path. For example, if the application imports two packages, `some/package` and `some/other/package`, both defining a variable `A`, the two variables can be accessed using this syntax:

​	具有相同名称的包可以通过使用完整的包路径来区分。例如，如果应用程序导入了两个包 `some/package` 和 `some/other/package`，并且这两个包都定义了变量 `A`，则可以使用以下语法访问这两个变量：

```
(dlv) p "some/package".A
(dlv) p "some/other/package".A
```

# Cgo 中的指针 Pointers in Cgo

Char pointers are always treated as NUL terminated strings, both indexing and the slice operator can be applied to them. Other C pointers can also be used similarly to Go slices, with indexing and the slice operator. In both of these cases it is up to the user to respect array bounds.

​	字符指针始终被视为 NUL 终止的字符串，既可以对其进行索引，也可以应用切片操作符。其他 C 指针也可以像 Go 切片一样使用，支持索引和切片操作符。在这两种情况下，用户需要自己保证数组边界。

# 特殊功能 Special Features

## 特殊变量 Special Variables

Delve defines two special variables:

​	Delve 定义了两个特殊变量：

* `runtime.curg` evaluates to the 'g' struct for the current goroutine, in particular `runtime.curg.goid` is the goroutine id of the current goroutine.
  * `runtime.curg` 评估为当前 goroutine 的 'g' 结构，特别是 `runtime.curg.goid` 是当前 goroutine 的 goroutine id。

* `runtime.frameoff` is the offset of the frame's base address from the bottom of the stack.
  * `runtime.frameoff` 是帧的基地址与栈底之间的偏移量。


## 访问前面帧的变量 Access to variables from previous frames

Variables from previous frames (i.e. stack frames other than the top of the stack) can be referred using the following notation `runtime.frame(n).name` which is the variable called 'name' on the n-th frame from the top of the stack.

​	可以使用以下表示法 `runtime.frame(n).name` 引用前面帧（即栈顶之外的栈帧）中的变量 `name`，其中 `n` 是从栈顶开始的第 `n` 个帧。

## CPU 寄存器 CPU Registers

The name of a CPU register, in all uppercase letters, will resolve to the value of that CPU register in the current frame. For example on AMD64 the expression `RAX` will evaluate to the value of the RAX register. 

​	CPU 寄存器的名称（全大写字母）将解析为当前帧中该寄存器的值。例如，在 AMD64 上，表达式 `RAX` 将评估为 RAX 寄存器的值。

Register names are shadowed by both local and global variables, so if a local variable called "RAX" exists, the `RAX` expression will evaluate to it instead of the CPU register.

​	寄存器名称会被局部变量和全局变量覆盖，因此，如果存在名为 `RAX` 的局部变量，`RAX` 表达式将解析为它而不是 CPU 寄存器。

Register names can optionally be prefixed by any number of underscore characters, so `RAX`, `_RAX`, `__RAX`, etc... can all be used to refer to the same RAX register and, in absence of shadowing from other variables, will all evaluate to the same value.

​	寄存器名称可以选择性地前缀任意数量的下划线字符，因此 `RAX`、`_RAX`、`__RAX` 等都可以用来引用同一个 RAX 寄存器，并且在没有其他变量的覆盖下，它们都会解析为相同的值。

Registers of 64bits or less are returned as uint64 variables. Larger registers are returned as strings of hexadecimal digits.

​	64 位或更小的寄存器返回为 `uint64` 类型的变量。更大的寄存器返回为十六进制数字的字符串。

Because many architectures have SIMD registers that can be used by the application in different ways the following syntax is also available:

​	因为许多架构有 SIMD 寄存器，可以以不同方式被应用程序使用，以下语法也可用：

* `REGNAME.intN` returns the register REGNAME as an array of intN elements.
  * `REGNAME.intN` 返回寄存器 REGNAME 作为 `intN` 元素的数组。

* `REGNAME.uintN` returns the register REGNAME as an array of uintN elements.
  * `REGNAME.uintN` 返回寄存器 REGNAME 作为 `uintN` 元素的数组。

* `REGNAME.floatN` returns the register REGNAME as an array of floatN elements.
  * `REGNAME.floatN` 返回寄存器 REGNAME 作为 `floatN` 元素的数组。


In all cases N must be a power of 2.

​	在所有这些情况下，`N` 必须是 2 的幂。

