+++
title = "dist"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# dist

> 原文：[https://pkg.go.dev/cmd/dist@go1.19.3](https://pkg.go.dev/cmd/dist@go1.19.3)

### Overview 

Dist helps bootstrap, build, and test the Go distribution.

​	dist 帮助引导、构建和测试 Go 发行版。

Usage:

​	用法：

```
go tool dist [command]
```

The commands are:

​	命令如下：

```
banner         print installation banner
				打印安装横幅
bootstrap      rebuild everything
				重新构建所有内容
clean          deletes all built files
				删除所有已构建的文件
env [-p]       print environment (-p: include $PATH)
				打印环境变量（-p：包括 $PATH）
install [dir]  install individual directory
				安装单个目录
list [-json]   list all supported platforms
				列出所有支持的平台
test [-h]      run Go test(s)
				运行 Go 测试
version        print Go version
				打印 Go 版本
```