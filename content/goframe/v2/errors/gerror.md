+++
title = "gerror"
date = 2024-03-21T17:51:03+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/errors/gerror](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/errors/gerror)

Package gerror provides rich functionalities to manipulate errors.

​	软件包 gerror 提供了丰富的功能来操作错误。

For maintainers, please very note that, this package is quite a basic package, which SHOULD NOT import extra packages except standard packages and internal packages, to avoid cycle imports.

​	对于维护者来说，请注意，这个包是一个非常基本的包，除了标准包和内部包之外，不应该导入额外的包，以避免循环导入。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Cause

```go
func Cause(err error) error
```

Cause returns the root cause error of `err`.

​	Cause 返回 的 `err` 根本原因错误。

#### func Code

```go
func Code(err error) gcode.Code
```

Code returns the error code of `current error`. It returns `CodeNil` if it has no error code neither it does not implement interface Code.

​	代码返回 的错误 `current error` 代码。如果它没有错误代码，则返回 `CodeNil` 它，它也没有实现接口代码。

#### func Current

```go
func Current(err error) error
```

Current creates and returns the current level error. It returns nil if current level error is nil.

​	Current 创建并返回当前电平错误。如果当前电平误差为 nil，则返回 nil。

#### func Equal <-2.1.0

```go
func Equal(err, target error) bool
```

Equal reports whether current error `err` equals to error `target`. Please note that, in default comparison logic for `Error`, the errors are considered the same if both the `code` and `text` of them are the same.

​	等于 报告当前误差 `err` 是否等于误差 `target` 。请注意，在 的 `Error` 默认比较逻辑中，如果 `code` 和 `text` 的 相同，则认为错误相同。

##### Example

``` go
```

#### func HasCode <-2.1.3

```go
func HasCode(err error, code gcode.Code) bool
```

HasCode checks and reports whether `err` has `code` in its chaining errors.

​	HasCode 检查并报告其链接中是否有 `err` `code` 错误。

#### func HasError <-2.1.3

```go
func HasError(err, target error) bool
```

HasError is alias of Is, which more easily understanding semantics.

​	HasError 是 Is 的别名，更容易理解语义。

#### func HasStack

```go
func HasStack(err error) bool
```

HasStack checks and reports whether `err` implemented interface `gerror.IStack`.

​	HasStack 检查并报告是否 `err` 实现了接口 `gerror.IStack` 。

#### func Is <-2.1.0

```go
func Is(err, target error) bool
```

Is reports whether current error `err` has error `target` in its chaining errors. It is just for implements for stdlib errors.Is from Go version 1.17.

​	报告当前错误 `err` 是否在其链接错误中具有错误 `target` 。它仅适用于 stdlib 错误的实现。来自 Go 版本 1.17。

##### Example

``` go
```

#### func New

```go
func New(text string) error
```

New creates and returns an error which is formatted from given text.

​	New 创建并返回一个错误，该错误是从给定文本格式化的。

#### func NewCode

```go
func NewCode(code gcode.Code, text ...string) error
```

NewCode creates and returns an error that has error code and given text.

​	NewCode 创建并返回具有错误代码和给定文本的错误。

##### Example

``` go
```

#### func NewCodeSkip

```go
func NewCodeSkip(code gcode.Code, skip int, text ...string) error
```

NewCodeSkip creates and returns an error which has error code and is formatted from given text. The parameter `skip` specifies the stack callers skipped amount.

​	NewCodeSkip 创建并返回一个错误，该错误具有错误代码，并且是根据给定文本格式化的。该参数 `skip` 指定堆栈调用方跳过的数量。

#### func NewCodeSkipf

```go
func NewCodeSkipf(code gcode.Code, skip int, format string, args ...interface{}) error
```

NewCodeSkipf returns an error that has error code and formats as the given format and args. The parameter `skip` specifies the stack callers skipped amount.

​	NewCodeSkipf 返回一个错误，该错误将错误代码和格式作为给定的格式和参数。该参数 `skip` 指定堆栈调用方跳过的数量。

#### func NewCodef

```go
func NewCodef(code gcode.Code, format string, args ...interface{}) error
```

NewCodef returns an error that has error code and formats as the given format and args.

​	NewCodef 返回一个错误，该错误将错误代码和格式作为给定的格式和参数。

##### Example

``` go
```

#### func NewOption

```go
func NewOption(option Option) error
```

NewOption creates and returns a custom error with Option. Deprecated: use NewWithOption instead.

​	NewOption 使用 Option 创建并返回自定义错误。已弃用：请改用 NewWithOption。

#### func NewSkip

```go
func NewSkip(skip int, text string) error
```

NewSkip creates and returns an error which is formatted from given text. The parameter `skip` specifies the stack callers skipped amount.

​	NewSkip 创建并返回一个错误，该错误是根据给定文本格式化的。该参数 `skip` 指定堆栈调用方跳过的数量。

#### func NewSkipf

```go
func NewSkipf(skip int, format string, args ...interface{}) error
```

NewSkipf returns an error that formats as the given format and args. The parameter `skip` specifies the stack callers skipped amount.

​	NewSkipf 返回一个错误，该错误格式为给定的格式和 args。该参数 `skip` 指定堆栈调用方跳过的数量。

#### func NewWithOption <-2.6.0

```go
func NewWithOption(option Option) error
```

NewWithOption creates and returns a custom error with Option. It is the senior usage for creating error, which is often used internally in framework.

​	NewWithOption 使用 Option 创建并返回自定义错误。它是创建错误的高级用法，通常在框架内部使用。

#### func Newf

```go
func Newf(format string, args ...interface{}) error
```

Newf returns an error that formats as the given format and args.

​	Newf 返回一个错误，该错误格式为给定的格式和 args。

#### func Stack

```go
func Stack(err error) string
```

Stack returns the stack callers as string. It returns the error string directly if the `err` does not support stacks.

​	Stack 以字符串形式返回堆栈调用者。如果 不支持堆栈， `err` 则直接返回错误字符串。

#### func Unwrap <-2.1.0

```go
func Unwrap(err error) error
```

Unwrap returns the next level error. It returns nil if current level error or the next level error is nil.

​	Unwrap 返回下一级错误。如果当前级别错误或下一级错误为 nil，则返回 nil。

#### func Wrap

```go
func Wrap(err error, text string) error
```

Wrap wraps error with text. It returns nil if given err is nil. Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

​	用文本换行错误。如果给定的 err 为 nil，则返回 nil。请注意，它不会丢失包装错误的错误代码，因为它从中继承了错误代码。

#### func WrapCode

```go
func WrapCode(code gcode.Code, err error, text ...string) error
```

WrapCode wraps error with code and text. It returns nil if given err is nil.

​	WrapCode 用代码和文本包装错误。如果给定的 err 为 nil，则返回 nil。

##### Example

``` go
```

#### func WrapCodeSkip

```go
func WrapCodeSkip(code gcode.Code, skip int, err error, text ...string) error
```

WrapCodeSkip wraps error with code and text. It returns nil if given err is nil. The parameter `skip` specifies the stack callers skipped amount.

​	WrapCodeSkip 使用代码和文本包装错误。如果给定的 err 为 nil，则返回 nil。该参数 `skip` 指定堆栈调用方跳过的数量。

#### func WrapCodeSkipf

```go
func WrapCodeSkipf(code gcode.Code, skip int, err error, format string, args ...interface{}) error
```

WrapCodeSkipf wraps error with code and text that is formatted with given format and args. It returns nil if given err is nil. The parameter `skip` specifies the stack callers skipped amount.

​	WrapCodeSkipf 使用使用给定格式和参数格式化的代码和文本包装错误。如果给定的 err 为 nil，则返回 nil。该参数 `skip` 指定堆栈调用方跳过的数量。

#### func WrapCodef

```go
func WrapCodef(code gcode.Code, err error, format string, args ...interface{}) error
```

WrapCodef wraps error with code and format specifier. It returns nil if given `err` is nil.

​	WrapCodef 使用代码和格式说明符包装错误。如果给定 `err` 为 nil，则返回 nil。

##### Example

``` go
```

#### func WrapSkip

```go
func WrapSkip(skip int, err error, text string) error
```

WrapSkip wraps error with text. It returns nil if given err is nil. The parameter `skip` specifies the stack callers skipped amount. Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

​	WrapSkip 用文本换行错误。如果给定的 err 为 nil，则返回 nil。该参数 `skip` 指定堆栈调用方跳过的数量。请注意，它不会丢失包装错误的错误代码，因为它从中继承了错误代码。

#### func WrapSkipf

```go
func WrapSkipf(skip int, err error, format string, args ...interface{}) error
```

WrapSkipf wraps error with text that is formatted with given format and args. It returns nil if given err is nil. The parameter `skip` specifies the stack callers skipped amount. Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

​	WrapSkipf 使用使用给定格式和 args 格式化的文本包装错误。如果给定的 err 为 nil，则返回 nil。该参数 `skip` 指定堆栈调用方跳过的数量。请注意，它不会丢失包装错误的错误代码，因为它从中继承了错误代码。

#### func Wrapf

```go
func Wrapf(err error, format string, args ...interface{}) error
```

Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the format specifier. It returns nil if given `err` is nil. Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

​	Wrapf 返回一个错误，在调用 Wrapf 的点使用堆栈跟踪和格式说明符注释 err。如果给定 `err` 为 nil，则返回 nil。请注意，它不会丢失包装错误的错误代码，因为它从中继承了错误代码。

## 类型

### type Error

```go
type Error struct {
	// contains filtered or unexported fields
}
```

Error is custom error for additional features.

​	错误是附加功能的自定义错误。

#### (*Error) Cause

```go
func (err *Error) Cause() error
```

Cause returns the root cause error.

​	Cause 返回根本原因错误。

#### (*Error) Code

```go
func (err *Error) Code() gcode.Code
```

Code returns the error code. It returns CodeNil if it has no error code.

​	代码返回错误代码。如果没有错误代码，则返回 CodeNil。

#### (*Error) Current

```go
func (err *Error) Current() error
```

Current creates and returns the current level error. It returns nil if current level error is nil.

​	Current 创建并返回当前电平错误。如果当前电平误差为 nil，则返回 nil。

#### (*Error) Equal

```go
func (err *Error) Equal(target error) bool
```

Equal reports whether current error `err` equals to error `target`. Please note that, in default comparison for `Error`, the errors are considered the same if both the `code` and `text` of them are the same.

​	等于 报告当前误差 `err` 是否等于误差 `target` 。请注意，在 `Error` 的默认比较中，如果 `code` 和 `text` 的 和 相同，则认为错误相同。

#### (*Error) Error

```go
func (err *Error) Error() string
```

Error implements the interface of Error, it returns all the error as string.

​	Error 实现了 Error 的接口，它以字符串的形式返回所有错误。

#### (*Error) Format

```go
func (err *Error) Format(s fmt.State, verb rune)
```

Format formats the frame according to the fmt.Formatter interface.

​	格式根据 fmt 格式化帧。格式化程序界面。

%v, %s : Print all the error string; %-v, %-s : Print current level error string; %+s : Print full stack error list; %+v : Print the error string and full stack error list

​	%v， %s ： 打印所有错误字符串;%-v， %-s ： 打印当前电平错误字符串;%+s ： 打印全栈错误列表;%+v：打印错误字符串和全栈错误列表

#### (*Error) Is

```go
func (err *Error) Is(target error) bool
```

Is reports whether current error `err` has error `target` in its chaining errors. It is just for implements for stdlib errors.Is from Go version 1.17.

​	报告当前错误 `err` 是否在其链接错误中具有错误 `target` 。它仅适用于 stdlib 错误的实现。来自 Go 版本 1.17。

#### (Error) MarshalJSON

```go
func (err Error) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that do not use pointer as its receiver here.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。请注意，此处不要使用指针作为其接收器。

#### (*Error) SetCode

```go
func (err *Error) SetCode(code gcode.Code)
```

SetCode updates the internal code with given code.

​	SetCode 使用给定的代码更新内部代码。

#### (*Error) Stack

```go
func (err *Error) Stack() string
```

Stack returns the error stack information as string.

​	Stack 以字符串形式返回错误堆栈信息。

#### (*Error) Unwrap

```go
func (err *Error) Unwrap() error
```

Unwrap is alias of function `Next`. It is just for implements for stdlib errors.Unwrap from Go version 1.17.

​	Unwrap 是函数 `Next` 的别名。它仅适用于 stdlib 错误的实现。从 Go 版本 1.17 解包。

### type ICause <-2.1.3

```go
type ICause interface {
	Error() string
	Cause() error
}
```

ICause is the interface for Cause feature.

​	ICause 是 Cause 功能的接口。

### type ICode <-2.1.3

```go
type ICode interface {
	Error() string
	Code() gcode.Code
}
```

ICode is the interface for Code feature.

​	ICode 是代码功能的接口。

### type ICurrent <-2.1.3

```go
type ICurrent interface {
	Error() string
	Current() error
}
```

ICurrent is the interface for Current feature.

​	ICurrent 是 Current 功能的接口。

### type IEqual <-2.1.3

```go
type IEqual interface {
	Error() string
	Equal(target error) bool
}
```

IEqual is the interface for Equal feature.

​	IEqual 是 Equal 功能的接口。

### type IIs <-2.1.3

```go
type IIs interface {
	Error() string
	Is(target error) bool
}
```

IIs is the interface for Is feature.

​	IIs 是 Is 功能的接口。

### type IStack <-2.1.3

```go
type IStack interface {
	Error() string
	Stack() string
}
```

IStack is the interface for Stack feature.

​	IStack 是堆栈功能的接口。

### type IUnwrap <-2.1.3

```go
type IUnwrap interface {
	Error() string
	Unwrap() error
}
```

IUnwrap is the interface for Unwrap feature.

​	IUnwrap 是 Unwrap 功能的接口。

### type Option

```go
type Option struct {
	Error error      // Wrapped error if any.
	Stack bool       // Whether recording stack information into error.
	Text  string     // Error text, which is created by New* functions.
	Code  gcode.Code // Error code if necessary.
}
```

Option is option for creating error.

​	选项是创建错误的选项。