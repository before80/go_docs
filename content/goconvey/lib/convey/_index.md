+++
title = "convey"
date = 2024-12-15T21:19:13+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/convey](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/convey)
>
> 收录该文档时间： `2024-12-15T21:19:13+08:00`

## Overview 

Package convey contains all of the public-facing entry points to this project. This means that it should never be required of the user to import any other packages from this project as they serve internal purposes.

## 常量

[View Source](https://github.com/smartystreets/goconvey/blob/v1.8.1/convey/doc.go#L148)

``` go
const (

	// FailureContinues is a failure mode which prevents failing
	// So()-assertions from halting Convey-block execution, instead
	// allowing the test to continue past failing So()-assertions.
	FailureContinues FailureMode = "continue"

	// FailureHalts is the default setting for a top-level Convey()-block
	// and will cause all failing So()-assertions to halt further execution
	// in that test-arm and continue on to the next arm.
	FailureHalts FailureMode = "halt"

	// FailureInherits is the default setting for failure-mode, it will
	// default to the failure-mode of the parent block. You should never
	// need to specify this mode in your tests..
	FailureInherits FailureMode = "inherits"

	// StackError is a stack mode which tells Convey to print stack traces
	// only for errors and not for test failures
	StackError StackMode = "error"

	// StackFail is a stack mode which tells Convey to print stack traces
	// for both errors and test failures
	StackFail StackMode = "fail"

	// StackInherits is the default setting for stack-mode, it will
	// default to the stack-mode of the parent block. You should never
	// need to specify this mode in your tests..
	StackInherits StackMode = "inherits"
)
```

## 变量

[View Source](https://github.com/smartystreets/goconvey/blob/v1.8.1/convey/assertions.go#L11)

``` go
var (
	ShouldAlmostEqual            = assertions.ShouldAlmostEqual
	ShouldBeBetween              = assertions.ShouldBeBetween
	ShouldBeBetweenOrEqual       = assertions.ShouldBeBetweenOrEqual
	ShouldBeBlank                = assertions.ShouldBeBlank
	ShouldBeChronological        = assertions.ShouldBeChronological
	ShouldBeEmpty                = assertions.ShouldBeEmpty
	ShouldBeError                = assertions.ShouldBeError
	ShouldBeFalse                = assertions.ShouldBeFalse
	ShouldBeGreaterThan          = assertions.ShouldBeGreaterThan
	ShouldBeGreaterThanOrEqualTo = assertions.ShouldBeGreaterThanOrEqualTo
	ShouldBeIn                   = assertions.ShouldBeIn
	ShouldBeLessThan             = assertions.ShouldBeLessThan
	ShouldBeLessThanOrEqualTo    = assertions.ShouldBeLessThanOrEqualTo
	ShouldBeNil                  = assertions.ShouldBeNil
	ShouldBeTrue                 = assertions.ShouldBeTrue
	ShouldBeZeroValue            = assertions.ShouldBeZeroValue
	ShouldContain                = assertions.ShouldContain
	ShouldContainKey             = assertions.ShouldContainKey
	ShouldContainSubstring       = assertions.ShouldContainSubstring
	ShouldEndWith                = assertions.ShouldEndWith
	ShouldEqual                  = assertions.ShouldEqual
	ShouldEqualJSON              = assertions.ShouldEqualJSON
	ShouldEqualTrimSpace         = assertions.ShouldEqualTrimSpace
	ShouldEqualWithout           = assertions.ShouldEqualWithout
	ShouldHappenAfter            = assertions.ShouldHappenAfter
	ShouldHappenBefore           = assertions.ShouldHappenBefore
	ShouldHappenBetween          = assertions.ShouldHappenBetween
	ShouldHappenOnOrAfter        = assertions.ShouldHappenOnOrAfter
	ShouldHappenOnOrBefore       = assertions.ShouldHappenOnOrBefore
	ShouldHappenOnOrBetween      = assertions.ShouldHappenOnOrBetween
	ShouldHappenWithin           = assertions.ShouldHappenWithin
	ShouldHaveLength             = assertions.ShouldHaveLength
	ShouldHaveSameTypeAs         = assertions.ShouldHaveSameTypeAs
	ShouldImplement              = assertions.ShouldImplement
	ShouldNotAlmostEqual         = assertions.ShouldNotAlmostEqual
	ShouldNotBeBetween           = assertions.ShouldNotBeBetween
	ShouldNotBeBetweenOrEqual    = assertions.ShouldNotBeBetweenOrEqual
	ShouldNotBeBlank             = assertions.ShouldNotBeBlank
	ShouldNotBeChronological     = assertions.ShouldNotBeChronological
	ShouldNotBeEmpty             = assertions.ShouldNotBeEmpty
	ShouldNotBeIn                = assertions.ShouldNotBeIn
	ShouldNotBeNil               = assertions.ShouldNotBeNil
	ShouldNotBeZeroValue         = assertions.ShouldNotBeZeroValue
	ShouldNotContain             = assertions.ShouldNotContain
	ShouldNotContainKey          = assertions.ShouldNotContainKey
	ShouldNotContainSubstring    = assertions.ShouldNotContainSubstring
	ShouldNotEndWith             = assertions.ShouldNotEndWith
	ShouldNotEqual               = assertions.ShouldNotEqual
	ShouldNotHappenOnOrBetween   = assertions.ShouldNotHappenOnOrBetween
	ShouldNotHappenWithin        = assertions.ShouldNotHappenWithin
	ShouldNotHaveSameTypeAs      = assertions.ShouldNotHaveSameTypeAs
	ShouldNotImplement           = assertions.ShouldNotImplement
	ShouldNotPanic               = assertions.ShouldNotPanic
	ShouldNotPanicWith           = assertions.ShouldNotPanicWith
	ShouldNotPointTo             = assertions.ShouldNotPointTo
	ShouldNotResemble            = assertions.ShouldNotResemble
	ShouldNotStartWith           = assertions.ShouldNotStartWith
	ShouldPanic                  = assertions.ShouldPanic
	ShouldPanicWith              = assertions.ShouldPanicWith
	ShouldPointTo                = assertions.ShouldPointTo
	ShouldResemble               = assertions.ShouldResemble
	ShouldStartWith              = assertions.ShouldStartWith
	ShouldWrap                   = assertions.ShouldWrap
)
```

These assertions are forwarded from github.com/smarty/assertions in order to make convey self-contained.

## 函数

### func Convey 

``` go
func Convey(items ...any)
```

Convey is the method intended for use when declaring the scopes of a specification. Each scope has a description and a func() which may contain other calls to Convey(), Reset() or Should-style assertions. Convey calls can be nested as far as you see fit.

IMPORTANT NOTE: The top-level Convey() within a Test method must conform to the following signature:

```
Convey(description string, t *testing.T, action func())
```

All other calls should look like this (no need to pass in *testing.T):

```
Convey(description string, action func())
```

Don't worry, goconvey will panic if you get it wrong so you can fix it.

Additionally, you may explicitly obtain access to the Convey context by doing:

```
Convey(description string, action func(c C))
```

You may need to do this if you want to pass the context through to a goroutine, or to close over the context in a handler to a library which calls your handler in a goroutine (httptest comes to mind).

All Convey()-blocks also accept an optional parameter of FailureMode which sets how goconvey should treat failures for So()-assertions in the block and nested blocks. See the constants in this file for the available options.

By default it will inherit from its parent block and the top-level blocks default to the FailureHalts setting.

This parameter is inserted before the block itself:

```
Convey(description string, t *testing.T, mode FailureMode, action func())
Convey(description string, mode FailureMode, action func())
```

See the examples package for, well, examples.

### func FocusConvey <-v1.6.1

``` go
func FocusConvey(items ...any)
```

FocusConvey is has the inverse effect of SkipConvey. If the top-level Convey is changed to `FocusConvey`, only nested scopes that are defined with FocusConvey will be run. The rest will be ignored completely. This is handy when debugging a large suite that runs a misbehaving function repeatedly as you can disable all but one of that function without swaths of `SkipConvey` calls, just a targeted chain of calls to FocusConvey.

### func Print <-v1.6.1

``` go
func Print(items ...any) (written int, err error)
```

Print is analogous to fmt.Print (and it even calls fmt.Print). It ensures that output is aligned with the corresponding scopes in the web UI.

### func PrintConsoleStatistics <-v1.6.1

``` go
func PrintConsoleStatistics()
```

PrintConsoleStatistics may be called at any time to print assertion statistics. Generally, the best place to do this would be in a TestMain function, after all tests have been run. Something like this:

``` go
func TestMain(m *testing.M) {
    convey.SuppressConsoleStatistics()
    result := m.Run()
    convey.PrintConsoleStatistics()
    os.Exit(result)
}
```

### func Printf <-v1.6.1

``` go
func Printf(format string, items ...any) (written int, err error)
```

Print is analogous to fmt.Printf (and it even calls fmt.Printf). It ensures that output is aligned with the corresponding scopes in the web UI.

### func Println <-v1.6.1

``` go
func Println(items ...any) (written int, err error)
```

Print is analogous to fmt.Println (and it even calls fmt.Println). It ensures that output is aligned with the corresponding scopes in the web UI.

### func Reset 

``` go
func Reset(action func())
```

Reset registers a cleanup function to be run after each Convey() in the same scope. See the examples package for a simple use case.

### func SetDefaultFailureMode <-v1.6.1

``` go
func SetDefaultFailureMode(mode FailureMode)
```

SetDefaultFailureMode allows you to specify the default failure mode for all Convey blocks. It is meant to be used in an init function to allow the default mode to be changed across all tests for an entire packgae but it can be used anywhere.

### func SetDefaultStackMode <-v1.6.5

``` go
func SetDefaultStackMode(mode StackMode)
```

SetDefaultStackMode allows you to specify the default stack mode for all Convey blocks. It is meant to be used in an init function to allow the default mode to be changed across all tests for an entire packgae but it can be used anywhere.

### func SkipConvey 

``` go
func SkipConvey(items ...any)
```

SkipConvey is analogous to Convey except that the scope is not executed (which means that child scopes defined within this scope are not run either). The reporter will be notified that this step was skipped.

### func SkipSo 

``` go
func SkipSo(stuff ...any)
```

SkipSo is analogous to So except that the assertion that would have been passed to So is not executed and the reporter is notified that the assertion was skipped.

### func So 

``` go
func So(actual any, assert Assertion, expected ...any)
```

So is the means by which assertions are made against the system under test. The majority of exported names in the assertions package begin with the word 'Should' and describe how the first argument (actual) should compare with any of the final (expected) arguments. How many final arguments are accepted depends on the particular assertion that is passed in as the assert argument. See the examples package for use cases and the assertions package for documentation on specific assertion methods. A failing assertion will cause t.Fail() to be invoked--you should never call this method (or other failure-inducing methods) in your test code. Leave that to GoConvey.

### func SoMsg <-v1.6.5

``` go
func SoMsg(msg string, actual any, assert Assertion, expected ...any)
```

SoMsg is an extension of So that allows you to specify a message to report on error.

### func SuppressConsoleStatistics <-v1.6.1

``` go
func SuppressConsoleStatistics()
```

SuppressConsoleStatistics prevents automatic printing of console statistics. Calling PrintConsoleStatistics explicitly will force printing of statistics.

## 类型

### type Assertion <-v1.6.5

``` go
type Assertion func(actual any, expected ...any) string
```

Assertion is an alias for a function with a signature that the convey.So() method can handle. Any future or custom assertions should conform to this method signature. The return value should be an empty string if the assertion passes and a well-formed failure message if not.

### type C <-v1.6.1

``` go
type C interface {
	Convey(items ...any)
	SkipConvey(items ...any)
	FocusConvey(items ...any)

	So(actual any, assert Assertion, expected ...any)
	SoMsg(msg string, actual any, assert Assertion, expected ...any)
	SkipSo(stuff ...any)

	Reset(action func())

	Println(items ...any) (int, error)
	Print(items ...any) (int, error)
	Printf(format string, items ...any) (int, error)
}
```

C is the Convey context which you can optionally obtain in your action by calling Convey like:

```
Convey(..., func(c C) {
  ...
})
```

See the documentation on Convey for more details.

All methods in this context behave identically to the global functions of the same name in this package.

### type FailureMode <-v1.6.1

``` go
type FailureMode string
```

FailureMode is a type which determines how the So() blocks should fail if their assertion fails. See constants further down for acceptable values

### type StackMode <-v1.6.5

``` go
type StackMode string
```

StackMode is a type which determines whether the So() blocks should report stack traces their assertion fails. See constants further down for acceptable values
