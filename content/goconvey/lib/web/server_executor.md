+++
title = "server_executor"
date = 2024-12-15T21:22:01+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/executor](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/executor)
>
> 收录该文档时间： `2024-12-15T21:22:01+08:00`

## 常量

[View Source](https://github.com/smartystreets/goconvey/blob/v1.8.1/web/server/executor/executor.go#L10)

``` go
const (
	Idle      = "idle"
	Executing = "executing"
)
```

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type ConcurrentTester 

``` go
type ConcurrentTester struct {
	// contains filtered or unexported fields
}
```

#### func NewConcurrentTester 

``` go
func NewConcurrentTester(shell contract.Shell) *ConcurrentTester
```

#### (*ConcurrentTester) SetBatchSize 

``` go
func (self *ConcurrentTester) SetBatchSize(batchSize int)
```

#### (*ConcurrentTester) TestAll 

``` go
func (self *ConcurrentTester) TestAll(folders []*contract.Package)
```

### type Executor 

``` go
type Executor struct {
	// contains filtered or unexported fields
}
```

#### func NewExecutor 

``` go
func NewExecutor(tester Tester, parser Parser, ch chan chan string) *Executor
```

#### (*Executor) ClearStatusFlag 

``` go
func (self *Executor) ClearStatusFlag() bool
```

#### (*Executor) ExecuteTests 

``` go
func (self *Executor) ExecuteTests(folders []*contract.Package) *contract.CompleteOutput
```

#### (*Executor) Status 

``` go
func (self *Executor) Status() string
```

### type Parser 

``` go
type Parser interface {
	Parse([]*contract.Package)
}
```

### type Tester 

``` go
type Tester interface {
	SetBatchSize(batchSize int)
	TestAll(folders []*contract.Package)
}
```
