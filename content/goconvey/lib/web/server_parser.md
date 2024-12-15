+++
title = "server_parser"
date = 2024-12-15T21:22:24+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/parser](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/parser)
>
> 收录该文档时间： `2024-12-15T21:22:24+08:00`

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func ParsePackageResults 

``` go
func ParsePackageResults(result *contract.PackageResult, rawOutput string)
```

## 类型

### type Parser <-v1.6.1

``` go
type Parser struct {
	// contains filtered or unexported fields
}
```

#### func NewParser <-v1.6.1

``` go
func NewParser(helper func(*contract.PackageResult, string)) *Parser
```

#### (*Parser) Parse <-v1.6.1

``` go
func (self *Parser) Parse(packages []*contract.Package)
```

### type TestResults <-v1.6.1

``` go
type TestResults []contract.TestResult
```

TestResults is a collection of TestResults that implements sort.Interface.

#### (TestResults) Len <-v1.6.1

``` go
func (r TestResults) Len() int
```

#### (TestResults) Less <-v1.6.1

``` go
func (r TestResults) Less(i, j int) bool
```

Less compares TestResults on TestName

#### (TestResults) Swap <-v1.6.1

``` go
func (r TestResults) Swap(i, j int)
```
