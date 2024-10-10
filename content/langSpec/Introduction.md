+++
title = "简介"
date = 2023-05-17T09:59:21+08:00
weight = 1
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
# The Go Programming Language Specification


## Version of June 29, 2022

## Introduction 简介
> 原文：[https://go.dev/ref/spec#Introduction](https://go.dev/ref/spec#Introduction)

This is the reference manual for the Go programming language. The pre-Go1.18 version, without generics, can be found [here](https://go.dev/doc/go1.17_spec.html). For more information and other documents, see [go.dev](https://go.dev/).

​	这是Go编程语言的参考手册。不带泛型的 `Go1.18` 之前的版本可以在[这里](https://go.dev/doc/go1.17_spec.html)找到。更多信息和其他文档，请参见 [go.dev](https://go.dev/)。

Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from *packages*, whose properties allow efficient management of dependencies.

​	Go是一种通用语言，设计时考虑到了系统编程。它是强类型和垃圾收集的，并且明确支持并发编程。程序是由包构成的，包的属性允许对依赖项进行有效的管理。

The syntax is compact and simple to parse, allowing for easy analysis by automatic tools such as integrated development environments.

​	语法紧凑且易于解析，允许使用集成开发环境等自动工具进行轻松的分析。