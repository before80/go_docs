+++
title = "go help go.mod "
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 550
description = ""
isCJKLanguage = true
draft = false

+++

​	

A module version is defined by a tree of source files, with a go.mod file in its root. When the go command is run, it looks in the current directory and then successive parent directories to find the go.mod marking the root of the main (current) module.

​	模块版本由一组源文件定义，其中包含根目录中的 go.mod 文件。运行 go 命令时，它会在当前目录以及随后的父目录中查找 go.mod 文件，以找到标记主（当前）模块的根。

The go.mod file format is described in detail at https://golang.org/ref/mod#go-mod-file.

​	go.mod 文件格式的详细说明，请参见https://golang.org/ref/mod#go-mod-file。

To create a new go.mod file, use 'go mod init'. For details see 'go help mod init' or https://golang.org/ref/mod#go-mod-init.

​	要创建新的 go.mod 文件，请使用 'go mod init'。有关详细信息，请参见 'go help mod init' 或 https://golang.org/ref/mod#go-mod-init。

To add missing module requirements or remove unneeded requirements, use 'go mod tidy'. For details, see 'go help mod tidy' or https://golang.org/ref/mod#go-mod-tidy.

​	要添加缺少的模块要求或删除不需要的要求，请使用 'go mod tidy'。有关详细信息，请参见 'go help mod tidy' 或 https://golang.org/ref/mod#go-mod-tidy。

To add, upgrade, downgrade, or remove a specific module requirement, use 'go get'. For details, see 'go help module-get' or https://golang.org/ref/mod#go-get.

​	要添加、升级、降级或删除特定的模块要求，请使用 'go get'。有关详细信息，请参见 'go help module-get' 或 https://golang.org/ref/mod#go-get。

To make other changes or to parse go.mod as JSON for use by other tools, use 'go mod edit'. See 'go help mod edit' or https://golang.org/ref/mod#go-mod-edit.

​	要进行其他更改或将 go.mod 解析为 JSON 以供其他工具使用，请使用 'go mod edit'。请参见 'go help mod edit' 或 https://golang.org/ref/mod#go-mod-edit。
