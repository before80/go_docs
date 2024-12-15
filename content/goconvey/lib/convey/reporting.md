+++
title = "reporting"
date = 2024-12-15T21:19:44+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/convey/reporting](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/convey/reporting)
>
> 收录该文档时间： `2024-12-15T21:19:44+08:00`

## Overview 

Package reporting contains internal functionality related to console reporting and output. Although this package has exported names is not intended for public consumption. See the examples package for how to use this project.

## 常量

[View Source](https://github.com/smartystreets/goconvey/blob/v1.8.1/convey/reporting/json.go#L80)

``` go
const CloseJson = "<-<-<-CLOSE-JSON<-<-<" // "⌫"
```

[View Source](https://github.com/smartystreets/goconvey/blob/v1.8.1/convey/reporting/json.go#L79)

``` go
const OpenJson = ">->->OPEN-JSON->->->" // "⌦"
```

## 变量

This section is empty.

## 函数

### func NewConsole 

``` go
func NewConsole() io.Writer
```

### func NewDotReporter 

``` go
func NewDotReporter(out *Printer) *dot
```

### func NewGoTestReporter 

``` go
func NewGoTestReporter() *gotestReporter
```

### func NewProblemReporter 

``` go
func NewProblemReporter(out *Printer) *problem
```

### func NewReporters 

``` go
func NewReporters(collection ...Reporter) *reporters
```

### func NewSilentProblemReporter 

``` go
func NewSilentProblemReporter(out *Printer) *problem
```

### func NewStatisticsReporter 

``` go
func NewStatisticsReporter(out *Printer) *statistics
```

### func NewStoryReporter 

``` go
func NewStoryReporter(out *Printer) *story
```

### func PrintConsoleStatistics 

``` go
func PrintConsoleStatistics()
```

### func QuietMode 

``` go
func QuietMode()
```

QuietMode disables all console output symbols. This is only meant to be used for tests that are internal to goconvey where the output is distracting or otherwise not needed in the test output.

### func SuppressConsoleStatistics 

``` go
func SuppressConsoleStatistics()
```

## 类型

### type AssertionResult 

``` go
type AssertionResult struct {
	File       string
	Line       int
	Expected   string
	Actual     string
	Failure    string
	Error      any
	StackTrace string
	Skipped    bool
}
```

#### func NewErrorReport 

``` go
func NewErrorReport(err any) *AssertionResult
```

#### func NewFailureReport 

``` go
func NewFailureReport(failure string, showStack bool) *AssertionResult
```

#### func NewSkipReport 

``` go
func NewSkipReport() *AssertionResult
```

#### func NewSuccessReport 

``` go
func NewSuccessReport() *AssertionResult
```

### type FailureView 

``` go
type FailureView struct {
	Message  string `json:"Message"`
	Expected string `json:"Expected"`
	Actual   string `json:"Actual"`
}
```

FailureView is also declared in github.com/smarty/assertions. The json struct tags should be equal in both declarations.

### type JsonReporter 

``` go
type JsonReporter struct {
	// contains filtered or unexported fields
}
```

#### func NewJsonReporter 

``` go
func NewJsonReporter(out *Printer) *JsonReporter
```

#### (*JsonReporter) BeginStory 

``` go
func (self *JsonReporter) BeginStory(story *StoryReport)
```

#### (*JsonReporter) EndStory 

``` go
func (self *JsonReporter) EndStory()
```

#### (*JsonReporter) Enter 

``` go
func (self *JsonReporter) Enter(scope *ScopeReport)
```

#### (*JsonReporter) Exit 

``` go
func (self *JsonReporter) Exit()
```

#### (*JsonReporter) Report 

``` go
func (self *JsonReporter) Report(report *AssertionResult)
```

#### (*JsonReporter) Write 

``` go
func (self *JsonReporter) Write(content []byte) (written int, err error)
```

### type Printer 

``` go
type Printer struct {
	// contains filtered or unexported fields
}
```

#### func NewPrinter 

``` go
func NewPrinter(out io.Writer) *Printer
```

#### (*Printer) Dedent 

``` go
func (self *Printer) Dedent()
```

#### (*Printer) Indent 

``` go
func (self *Printer) Indent()
```

#### (*Printer) Insert 

``` go
func (self *Printer) Insert(text string)
```

#### (*Printer) Print 

``` go
func (self *Printer) Print(message string, values ...any)
```

#### (*Printer) Println 

``` go
func (self *Printer) Println(message string, values ...any)
```

### type Reporter 

``` go
type Reporter interface {
	BeginStory(story *StoryReport)
	Enter(scope *ScopeReport)
	Report(r *AssertionResult)
	Exit()
	EndStory()
	io.Writer
}
```

#### func BuildDotReporter 

``` go
func BuildDotReporter() Reporter
```

#### func BuildJsonReporter 

``` go
func BuildJsonReporter() Reporter
```

#### func BuildSilentReporter 

``` go
func BuildSilentReporter() Reporter
```

#### func BuildStoryReporter 

``` go
func BuildStoryReporter() Reporter
```

### type ScopeReport 

``` go
type ScopeReport struct {
	Title string
	File  string
	Line  int
}
```

#### func NewScopeReport 

``` go
func NewScopeReport(title string) *ScopeReport
```

### type ScopeResult 

``` go
type ScopeResult struct {
	Title      string
	File       string
	Line       int
	Depth      int
	Assertions []*AssertionResult
	Output     string
}
```

### type StoryReport 

``` go
type StoryReport struct {
	Test T
	Name string
	File string
	Line int
}
```

#### func NewStoryReport 

``` go
func NewStoryReport(test T) *StoryReport
```

### type T 

``` go
type T interface {
	Fail()
}
```

This interface allows us to pass the *testing.T struct throughout the internals of this tool without ever having to import the "testing" package.
