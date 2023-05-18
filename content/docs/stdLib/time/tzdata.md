+++
title = "tzdata"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# tzdata

https://pkg.go.dev/time/tzdata@go1.20.1



​	tzdata包 提供了时区数据库的嵌入式副本。如果程序的任何地方导入了该包，那么如果时间包在系统上找不到 tzdata 文件，它将使用这个嵌入式信息。

​	导入此包将使程序的大小增加约450 KB。

​	通常应该由程序的main包导入此包，而不是由库导入。库通常不应决定是否在程序中包含时区数据库。

​	如果您使用 -tags timetzdata 进行构建，将自动导入此包。