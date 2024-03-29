+++
title = "gmode"
date = 2024-03-21T17:59:37+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gmode](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gmode)

Package gmode provides release mode management for project.

​	软件包 gmode 为项目提供发布模式管理。

It uses string to mark the mode instead of integer, which is convenient for configuration.

​	它使用字符串而不是整数来标记模式，方便配置。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gmode/gmode.go#L18)

```go
const (
	NOT_SET = "not-set"
	DEVELOP = "develop"
	TESTING = "testing"
	STAGING = "staging"
	PRODUCT = "product"
)
```

## 变量

This section is empty.

## 函数

#### func IsDevelop

```go
func IsDevelop() bool
```

IsDevelop checks and returns whether current application is running in DEVELOP mode.

​	IsDevelop 检查并返回当前应用程序是否在 DEVELOP 模式下运行。

#### func IsProduct

```go
func IsProduct() bool
```

IsProduct checks and returns whether current application is running in PRODUCT mode.

​	IsProduct 检查并返回当前应用程序是否在 PRODUCT 模式下运行。

#### func IsStaging

```go
func IsStaging() bool
```

IsStaging checks and returns whether current application is running in STAGING mode.

​	IsStaging 检查并返回当前应用程序是否在暂存模式下运行。

#### func IsTesting

```go
func IsTesting() bool
```

IsTesting checks and returns whether current application is running in TESTING mode.

​	IsTesting 检查并返回当前应用程序是否在 TESTING 模式下运行。

#### func Mode

```go
func Mode() string
```

Mode returns current application mode set.

​	Mode 返回当前应用程序模式集。

#### func Set

```go
func Set(mode string)
```

Set sets the mode for current application.

​	Set 设置当前应用程序的模式。

#### func SetDevelop

```go
func SetDevelop()
```

SetDevelop sets current mode DEVELOP for current application.

​	SetDevelop 为当前应用程序设置当前模式 DEVELOP。

#### func SetProduct

```go
func SetProduct()
```

SetProduct sets current mode PRODUCT for current application.

​	SetProduct 为当前应用程序设置当前模式 PRODUCT。

#### func SetStaging

```go
func SetStaging()
```

SetStaging sets current mode STAGING for current application.

​	SetStaging 为当前应用程序设置当前模式 STAGING。

#### func SetTesting

```go
func SetTesting()
```

SetTesting sets current mode TESTING for current application.

​	SetTesting 为当前应用程序设置当前模式 TESTING。

## 类型

This section is empty.