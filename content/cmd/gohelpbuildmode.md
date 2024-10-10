+++
title = "go help buildmode"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 420
description = ""
isCJKLanguage = true
draft = false

+++

​	

The '`go build`' and '`go install`' commands take a `-buildmode` argument which indicates which kind of object file is to be built. Currently supported values are:

​	'go build' 和 'go install' 命令接受一个 `-buildmode` 参数，指示要构建哪种类型的目标文件。当前支持的值有：

## -buildmode=archive

​        Build the listed non-main packages into .a files. Packages named main are ignored.

​	 构建列出的非 main 包到 .a 文件中。忽略名为 main 的包。

## -buildmode=c-archive

​        Build the listed main package, plus all packages it imports, into a C archive file. The only callable symbols will be those functions exported using a cgo //export comment. Requires exactly one main package to be listed.

 	构建列出的主包，以及它导入的所有包，到一个 C 存档文件中。可调用的唯一符号将是使用 cgo //export 注释导出的那些函数。要求列出正好一个主包。

## -buildmode=c-shared

​        Build the listed main package, plus all packages it imports, into a C shared library. The only callable symbols will be those functions exported using a cgo //export comment. Requires exactly one main package to be listed.

​	 构建列出的主包，以及它导入的所有包，到一个 C 共享库中。可调用的唯一符号将是使用 cgo //export 注释导出的那些函数。要求列出正好一个主包。

## -buildmode=default

​        Listed main packages are built into executables and listed non-main packages are built into .a files (the default behavior).

​	 列出的主包将构建为可执行文件，而列出的非 main 包将构建为 .a 文件（默认行为）。

## -buildmode=shared

​        Combine all the listed non-main packages into a single shared library that will be used when building with the -linkshared option. Packages named main are ignored.

​	 将所有列出的非 main 包合并到一个单独的共享库中，该库将在使用 -linkshared 选项构建时使用。忽略名为 main 的包。

## -buildmode=exe

​        Build the listed main packages and everything they import into executables. Packages not named main are ignored.

 	构建列出的主包以及它们导入的所有内容到可执行文件中。忽略非主包。

## -buildmode=pie

​        Build the listed main packages and everything they import into position independent executables (PIE). Packages not named main are ignored.

​	 构建列出的主包以及它们导入的所有内容到位置无关可执行文件（PIE）中。忽略非主包。

## -buildmode=plugin

​        Build the listed main packages, plus all packages that they import, into a Go plugin. Packages not named main are ignored.

​	 构建列出的主包以及它们导入的所有包到一个 Go 插件中。忽略非主包。



On AIX, when linking a C program that uses a Go archive built with `-buildmode=c-archive`, you must pass `-Wl`,`-bnoobjreorder` to the C compiler.

​	在链接使用 `-buildmode=c-archive` 构建的 Go 存档的 C 程序时，必须向 C 编译器传递 `-Wl, -bnoobjreorder`。
