+++
title = "gbuild"
date = 2024-03-21T17:54:25+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gbuild](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gbuild)

Package gbuild manages the build-in variables from “gf build”.

​	软件包 gbuild 管理来自 “gf build” 的内置变量。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gbuild/gbuild.go#L31)

```go
const (
	BuiltGit     = `builtGit`
	BuiltTime    = `builtTime`
	BuiltVersion = `builtVersion`
)
```

## 变量

This section is empty.

## 函数

#### func Data

```go
func Data() map[string]interface{}
```

Data returns the custom build-in variables as map.

​	数据以 map 的形式返回自定义内置变量。

#### func Get

```go
func Get(name string, def ...interface{}) *gvar.Var
```

Get retrieves and returns the build-in binary variable with given name.

​	Get 检索并返回具有给定名称的内置二进制变量。

## 类型

### type BuildInfo

```go
type BuildInfo struct {
	GoFrame string                 // Built used GoFrame version.
	Golang  string                 // Built used Golang version.
	Git     string                 // Built used git repo. commit id and datetime.
	Time    string                 // Built datetime.
	Version string                 // Built version.
	Data    map[string]interface{} // All custom built data key-value pairs.
}
```

BuildInfo maintains the built info of current binary.

​	BuildInfo 维护当前二进制文件的构建信息。

#### func Info

```go
func Info() BuildInfo
```

Info returns the basic built information of the binary as map. Note that it should be used with gf-cli tool “gf build”, which automatically injects necessary information into the binary.

​	Info 以 map 的形式返回二进制文件的基本构建信息。请注意，它应该与 gf-cli 工具“gf build”一起使用，它会自动将必要的信息注入到二进制文件中。