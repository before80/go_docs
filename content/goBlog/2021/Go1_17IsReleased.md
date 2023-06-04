+++
title = "go 1.17版发布了"
weight = 90
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.17 is released - go 1.17版发布了

https://go.dev/blog/go1.17

Matt Pearring and Alex Rakoczy
16 August 2021

Today the Go team is thrilled to release Go 1.17, which you can get by visiting the [download page](https://go.dev/dl/).

今天，Go团队很高兴地发布了Go 1.17，您可以通过访问下载页面获得该版本。

This release brings additional improvements to the compiler, namely a [new way of passing function arguments and results](https://go.dev/doc/go1.17#compiler). This change has shown about a 5% performance improvement in Go programs and reduction in binary sizes of around 2% for amd64 platforms. Support for more platforms will come in future releases.

这个版本给编译器带来了额外的改进，即采用了一种新的函数参数和结果传递方式。这一改变使Go程序的性能提高了约5%，并使amd64平台的二进制大小减少了约2%。对更多平台的支持将在未来的版本中出现。

Go 1.17 also adds support for the [64-bit ARM architecture on Windows](https://go.dev/doc/go1.17#ports), letting gophers run Go natively on more devices.

Go 1.17还增加了对Windows上64位ARM架构的支持，让地鼠在更多设备上原生运行Go。

We’ve also introduced [pruned module graphs](https://go.dev/doc/go1.17#go-command) in this release. Modules that specify `go 1.17` or higher in their `go.mod` file will have their module graphs include only the immediate dependencies of other Go 1.17 modules, not their full transitive dependencies. This should help avoid the need to download or read `go.mod` files for otherwise irrelevant dependencies—saving time in everyday development.

我们还在这个版本中引入了修剪模块图。在go.mod文件中指定go 1.17或更高版本的模块，其模块图只包括其他Go 1.17模块的直接依赖关系，而不是其全部的横向依赖关系。这将有助于避免下载或阅读go.mod文件，以获取不相关的依赖关系，从而在日常开发中节省时间。

Go 1.17 comes with three small [changes to the language](https://go.dev/doc/go1.17#language). The first two are new functions in the `unsafe` package to make it simpler for programs to conform to the `unsafe.Pointer` rules: `unsafe.Add` allows for [safer pointer arithmetic](https://go.dev/pkg/unsafe#Add), while `unsafe.Slice` allows for [safer conversions of pointers to slices](https://go.dev/pkg/unsafe#Slice). The third change is an extension to the language type conversion rules to allow conversions from [slices to array pointers](https://go.dev/ref/spec#Conversions_from_slice_to_array_pointer), provided the slice is at least as large as the array at runtime.

Go 1.17对语言有三个小改动。前两项是不安全包中的新函数，使程序更容易符合不安全指针规则：unsafe.Add允许更安全的指针运算，而unsafe.Slice允许更安全的指针转换为切片。第三个变化是对语言类型转换规则的扩展，允许从切片到数组指针的转换，只要在运行时切片至少与数组一样大。

Finally there are quite a few other improvements and bug fixes, including verification improvements to [crypto/x509](https://go.dev/doc/go1.17#crypto/x509), and alterations to [URL query parsing](https://go.dev/doc/go1.17#semicolons). For a complete list of changes and more information about the improvements above, see the [full release notes](https://go.dev/doc/go1.17).

最后还有一些其他的改进和错误修复，包括加密/x509的验证改进，以及URL查询解析的改变。关于完整的变化列表和有关上述改进的更多信息，请参见完整的发布说明。

Thanks to everyone who contributed to this release by writing code, filing bugs, sharing feedback, and testing the beta and release candidates. Your efforts helped to ensure that Go 1.17 is as stable as possible. As always, if you notice any problems, please [file an issue](https://go.dev/issue/new).

感谢所有通过编写代码、提交错误、分享反馈以及测试测试版和候选版而对该版本作出贡献的人。您的努力有助于确保Go 1.17尽可能的稳定。像往常一样，如果您发现任何问题，请提出问题。

We hope you enjoy the new release!

我们希望您喜欢这个新版本
