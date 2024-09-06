+++
title = "tzdata"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/time/tzdata@go1.23.0](https://pkg.go.dev/time/tzdata@go1.23.0)

Package tzdata provides an embedded copy of the timezone database. If this package is imported anywhere in the program, then if the time package cannot find tzdata files on the system, it will use this embedded information.

​	`tzdata`包 提供了时区数据库的嵌入式副本。如果程序的任何地方导入了该包，那么如果时间包在系统上找不到 tzdata 文件，它将使用这个嵌入式信息。

Importing this package will increase the size of a program by about 450 KB.

​	导入此包将使程序的大小增加约450 KB。

This package should normally be imported by a program's main package, not by a library. Libraries normally shouldn't decide whether to include the timezone database in a program.

​	通常应该由程序的main包导入此包，而不是由库导入。库通常不应决定是否在程序中包含时区数据库。

This package will be automatically imported if you build with -tags timetzdata.

​	如果您使用 -tags timetzdata 进行构建，将自动导入此包。