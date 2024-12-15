+++
title = "server_contract"
date = 2024-12-15T21:21:44+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/contract](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/contract)
>
> 收录该文档时间： `2024-12-15T21:21:44+08:00`

## 常量

This section is empty.

## 变量

[View Source](https://github.com/smartystreets/goconvey/blob/v1.8.1/web/server/contract/result.go#L49)

``` go
var (
	Ignored         = "ignored"
	Disabled        = "disabled"
	Passed          = "passed"
	Failed          = "failed"
	Panicked        = "panicked"
	BuildFailure    = "build failure"
	NoTestFiles     = "no test files"
	NoTestFunctions = "no test functions"
	NoGoFiles       = "no go code"

	TestRunAbortedUnexpectedly = "test run aborted unexpectedly"
)
```

## 函数

This section is empty.

## 类型

### type CompleteOutput 

``` go
type CompleteOutput struct {
	Packages []*PackageResult
	Revision string
	Paused   bool
}
```

### type Executor 

``` go
type Executor interface {
	ExecuteTests([]*Package) *CompleteOutput
	Status() string
	ClearStatusFlag() bool
}
```

### type Package 

``` go
type Package struct {
	Path          string
	Name          string
	Ignored       bool
	Disabled      bool
	BuildTags     []string
	TestArguments []string
	Error         error
	Output        string
	Result        *PackageResult

	HasImportCycle bool
}
```

#### func NewPackage 

``` go
func NewPackage(folder *messaging.Folder, name string, hasImportCycle bool) *Package
```

#### (*Package) Active 

``` go
func (self *Package) Active() bool
```

#### (*Package) HasUsableResult 

``` go
func (self *Package) HasUsableResult() bool
```

### type PackageResult 

``` go
type PackageResult struct {
	PackageName string
	Elapsed     float64
	Coverage    float64
	Outcome     string
	BuildOutput string
	TestResults []TestResult
}
```

#### func NewPackageResult 

``` go
func NewPackageResult(packageName string) *PackageResult
```

### type Server 

``` go
type Server interface {
	ReceiveUpdate(root string, update *CompleteOutput)
	Watch(writer http.ResponseWriter, request *http.Request)
	Ignore(writer http.ResponseWriter, request *http.Request)
	Reinstate(writer http.ResponseWriter, request *http.Request)
	Status(writer http.ResponseWriter, request *http.Request)
	LongPollStatus(writer http.ResponseWriter, request *http.Request)
	Results(writer http.ResponseWriter, request *http.Request)
	Execute(writer http.ResponseWriter, request *http.Request)
	TogglePause(writer http.ResponseWriter, request *http.Request)
}
```

### type Shell 

``` go
type Shell interface {
	GoTest(directory, packageName string, tags, arguments []string) (output string, err error)
}
```

### type TestResult 

``` go
type TestResult struct {
	TestName string
	Elapsed  float64
	Passed   bool
	Skipped  bool
	File     string
	Line     int
	Message  string
	Error    string
	Stories  []reporting.ScopeResult

	RawLines []string `json:",omitempty"`
}
```

#### func NewTestResult 

``` go
func NewTestResult(testName string) *TestResult
```
