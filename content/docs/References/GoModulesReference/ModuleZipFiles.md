+++
title = "模块zip文件"
date = 2023-05-17T09:59:21+08:00
weight = 11
description = ""
isCJKLanguage = true
draft = false
+++
## Module zip files 模块zip文件

> 原文：[https://go.dev/ref/mod#zip-files](https://go.dev/ref/mod#zip-files)

​	模块版本是以`.zip`文件的形式发布的。很少有必要直接与这些文件互动，因为`go`命令会自动从[module proxies（模块代理）](../Glossary#module-proxy)和版本控制库中创建、下载和提取这些文件。但是，了解这些文件对了解跨平台的兼容性约束或在实现模块代理时仍然很有用。

​	[go mod download](../gomodFiles#go-mod-download)命令下载一个或多个模块的zip文件，然后将这些文件提取到模块缓存中。根据`GOPROXY`和其他[环境变量](../EnvironmentVariables)，`go`命令可以从代理下载zip文件，或者克隆源码管理存储库并从中创建zip文件。`-json`标志可以用来查找下载的 zip文件及其在模块缓存中提取的内容的位置。

​	[golang.org/x/mod/zip](https://pkg.go.dev/golang.org/x/mod/zip?tab=doc)包可以用来以编程方式创建、提取或检查压缩文件的内容。

### File path and size constraints 文件路径和大小约束

​	对模块zip文件的内容有一些限制。这些限制确保zip文件可以在各种平台上安全和一致地被提取。

- 模块 zip文件最多可以有500 MiB的大小。`go.mod`文件被限制在16 MB以内。`LICENSE`文件也被限制在16 MB以内。这些限制的存在是为了减轻对用户、代理和模块生态系统的其他部分的拒绝服务攻击。在模块目录树中包含超过500 MiB的文件的存储库应该在提交时标记模块版本，只包括构建模块包所需的文件；视频、模型和其他大型资产通常不需要构建。

- 模块zip文件中的每个文件必须以前缀`$module@$version/`开始，其中`$module`是模块路径，`$version`是版本，例如`golang.org/x/mod@v0.3.0/`。模块路径必须是有效的，版本必须是有效的和经典的，并且版本必须与模块路径的主版本后缀匹配。具体定义和限制，请参见[Module paths and versions（模块路径和版本）](../gomodFiles#module-paths-and-versions)。

- 文件模式、时间戳和其他元数据被忽略。

- 空目录（路径以斜线结尾的条目）可能包含在模块zip文件中，但不会被提取。`go`命令在它创建的压缩文件中不包括空目录。

- 符号链接和其他不规则的文件在创建zip文件时被忽略，因为它们在不同的操作系统和文件系统中是不可移植的，也没有可移植的方法在zip文件格式中表示它们。

- 在创建zip文件时，名为`vendor`的目录内的文件被忽略，因为主模块外的`vendor`目录从不被使用。

- 在创建zip文件时，包含`go.mod`文件的目录中的文件(模块根目录除外)将被忽略，因为它们不是模块的一部分。`go`命令在提取文件时忽略了包含`go.mod`文件的子目录。

- 在Unicode大小写折叠下，zip文件中的任何两个文件的路径都不可能相等（见[strings.EqualFold](https://pkg.go.dev/strings?tab=doc#EqualFold)）。这保证了在不区分大小写的文件系统中提取文件时不会出现冲突。

- `go.mod`文件可能出现在顶层目录（`$module@$version/go.mod`）中，也可能不出现。如果出现，它的名字必须是`go.mod`（全小写）。名为`go.mod`的文件不允许出现在任何其他目录中。

- 模块中的文件和目录名可以由Unicode字母、ASCII数字、ASCII空格字符（U+0020）和ASCII标点字符`!#$%&()+,-.=@[]^_{}~`组成。注意，包的路径可能不包含所有这些字符。参见[module.CheckFilePath](https://pkg.go.dev/golang.org/x/mod/module?tab=doc#CheckFilePath)和[module.CheckImportPath](https://pkg.go.dev/golang.org/x/mod/module?tab=doc#CheckImportPath)的区别。

- 在Windows上，第一个点之前的文件或目录名不能是保留文件名，无论大小写(`CON`、`com1`、`NuL`等)。

  