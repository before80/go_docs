+++
title = "go help c"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 500
description = ""
isCJKLanguage = true
draft = false

+++

​	

There are two different ways to call between Go and C/C++ code.

​	有两种在 Go 与 C/C++ 代码之间进行调用的方法。

The first is the cgo tool, which is part of the Go distribution. For information on how to use it see the cgo documentation (go doc cmd/cgo).

​	第一种是 cgo 工具，它是 Go 发行版的一部分。有关如何使用它的信息，请参阅 cgo 文档（运行 'go doc cmd/cgo'）。

The second is the SWIG program, which is a general tool for interfacing between languages. For information on SWIG see http://swig.org/. When running go build, any file with a .swig extension will be passed to SWIG. Any file with a .swigcxx extension will be passed to SWIG with the -c++ option.

​	第二种是 SWIG 程序，它是一种用于在语言之间进行接口的通用工具。有关 SWIG 的信息，请参阅 http://swig.org/。在运行 go build 时，任何具有 .swig 扩展名的文件都将被传递给 SWIG。具有 .swigcxx 扩展名的任何文件都将以 -c++ 选项传递给 SWIG。

When either cgo or SWIG is used, go build will pass any .c, .m, .s, .S or .sx files to the C compiler, and any .cc, .cpp, .cxx files to the C++
compiler. The CC or CXX environment variables may be set to determine the C or C++ compiler, respectively, to use.

​	在使用 cgo 或 SWIG 时，go build 将任何 .c、.m、.s、.S 或 .sx 文件传递给 C 编译器，并将任何 .cc、.cpp、.cxx 文件传递给 C++ 编译器。CC 或 CXX 环境变量可以设置为确定要使用的 C 或 C++ 编译器。
