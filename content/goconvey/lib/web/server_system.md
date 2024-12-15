+++
title = "server_system"
date = 2024-12-15T21:22:32+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/system](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/system)
>
> 收录该文档时间： `2024-12-15T21:22:32+08:00`

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Command 

``` go
type Command struct {
	Output string
	Error  error
	// contains filtered or unexported fields
}
```

### func NewCommand 

``` go
func NewCommand(directory, executable string, arguments ...string) Command
```

#### (Command) Execute 

``` go
func (this Command) Execute() Command
```

### type Shell 

``` go
type Shell struct {
	// contains filtered or unexported fields
}
```

### func NewShell 

``` go
func NewShell(gobin, reportsPath string, coverage bool, defaultTimeout string) *Shell
```

#### (*Shell) GoTest 

``` go
func (self *Shell) GoTest(directory, packageName string, tags, arguments []string) (output string, err error)
```
