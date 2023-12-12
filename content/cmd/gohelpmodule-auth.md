+++
title = "go help module-auth"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

When the go command downloads a module zip file or go.mod file into the module cache, it computes a cryptographic hash and compares it with a known value to verify the file hasn't changed since it was first downloaded. Known hashes are stored in a file in the module root directory named go.sum. Hashes may also be downloaded from the checksum database depending on the values of GOSUMDB, GOPRIVATE, and GONOSUMDB.

​	当go命令将模块zip文件或go.mod文件下载到模块缓存时，它会计算加密哈希并与已知值进行比较，以验证文件自第一次下载以来未发生更改。已知哈希存储在模块根目录中名为go.sum的文件中。根据GOSUMDB、GOPRIVATE和GONOSUMDB的值，还可以从校验和数据库中下载哈希。

For details, see https://golang.org/ref/mod#authenticating.

​	有关详细信息，请参见 https://golang.org/ref/mod#authenticating。
