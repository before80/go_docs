+++
title = "gerror"
date = 2024-03-21T17:51:03+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/errors/gerror

Package gerror provides rich functionalities to manipulate errors.

For maintainers, please very note that, this package is quite a basic package, which SHOULD NOT import extra packages except standard packages and internal packages, to avoid cycle imports.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Cause 

``` go
func Cause(err error) error
```

Cause returns the root cause error of `err`.

##### func Code 

``` go
func Code(err error) gcode.Code
```

Code returns the error code of `current error`. It returns `CodeNil` if it has no error code neither it does not implement interface Code.

##### func Current 

``` go
func Current(err error) error
```

Current creates and returns the current level error. It returns nil if current level error is nil.

##### func Equal <-2.1.0

``` go
func Equal(err, target error) bool
```

Equal reports whether current error `err` equals to error `target`. Please note that, in default comparison logic for `Error`, the errors are considered the same if both the `code` and `text` of them are the same.

##### Example

``` go
```
##### func HasCode <-2.1.3

``` go
func HasCode(err error, code gcode.Code) bool
```

HasCode checks and reports whether `err` has `code` in its chaining errors.

##### func HasError <-2.1.3

``` go
func HasError(err, target error) bool
```

HasError is alias of Is, which more easily understanding semantics.

##### func HasStack 

``` go
func HasStack(err error) bool
```

HasStack checks and reports whether `err` implemented interface `gerror.IStack`.

##### func Is <-2.1.0

``` go
func Is(err, target error) bool
```

Is reports whether current error `err` has error `target` in its chaining errors. It is just for implements for stdlib errors.Is from Go version 1.17.

##### Example

``` go
```
##### func New 

``` go
func New(text string) error
```

New creates and returns an error which is formatted from given text.

##### func NewCode 

``` go
func NewCode(code gcode.Code, text ...string) error
```

NewCode creates and returns an error that has error code and given text.

##### Example

``` go
```
##### func NewCodeSkip 

``` go
func NewCodeSkip(code gcode.Code, skip int, text ...string) error
```

NewCodeSkip creates and returns an error which has error code and is formatted from given text. The parameter `skip` specifies the stack callers skipped amount.

##### func NewCodeSkipf 

``` go
func NewCodeSkipf(code gcode.Code, skip int, format string, args ...interface{}) error
```

NewCodeSkipf returns an error that has error code and formats as the given format and args. The parameter `skip` specifies the stack callers skipped amount.

##### func NewCodef 

``` go
func NewCodef(code gcode.Code, format string, args ...interface{}) error
```

NewCodef returns an error that has error code and formats as the given format and args.

##### Example

``` go
```
##### func NewOption 

``` go
func NewOption(option Option) error
```

NewOption creates and returns a custom error with Option. Deprecated: use NewWithOption instead.

##### func NewSkip 

``` go
func NewSkip(skip int, text string) error
```

NewSkip creates and returns an error which is formatted from given text. The parameter `skip` specifies the stack callers skipped amount.

##### func NewSkipf 

``` go
func NewSkipf(skip int, format string, args ...interface{}) error
```

NewSkipf returns an error that formats as the given format and args. The parameter `skip` specifies the stack callers skipped amount.

##### func NewWithOption <-2.6.0

``` go
func NewWithOption(option Option) error
```

NewWithOption creates and returns a custom error with Option. It is the senior usage for creating error, which is often used internally in framework.

##### func Newf 

``` go
func Newf(format string, args ...interface{}) error
```

Newf returns an error that formats as the given format and args.

##### func Stack 

``` go
func Stack(err error) string
```

Stack returns the stack callers as string. It returns the error string directly if the `err` does not support stacks.

##### func Unwrap <-2.1.0

``` go
func Unwrap(err error) error
```

Unwrap returns the next level error. It returns nil if current level error or the next level error is nil.

##### func Wrap 

``` go
func Wrap(err error, text string) error
```

Wrap wraps error with text. It returns nil if given err is nil. Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

##### func WrapCode 

``` go
func WrapCode(code gcode.Code, err error, text ...string) error
```

WrapCode wraps error with code and text. It returns nil if given err is nil.

##### Example

``` go
```
##### func WrapCodeSkip 

``` go
func WrapCodeSkip(code gcode.Code, skip int, err error, text ...string) error
```

WrapCodeSkip wraps error with code and text. It returns nil if given err is nil. The parameter `skip` specifies the stack callers skipped amount.

##### func WrapCodeSkipf 

``` go
func WrapCodeSkipf(code gcode.Code, skip int, err error, format string, args ...interface{}) error
```

WrapCodeSkipf wraps error with code and text that is formatted with given format and args. It returns nil if given err is nil. The parameter `skip` specifies the stack callers skipped amount.

##### func WrapCodef 

``` go
func WrapCodef(code gcode.Code, err error, format string, args ...interface{}) error
```

WrapCodef wraps error with code and format specifier. It returns nil if given `err` is nil.

##### Example

``` go
```
##### func WrapSkip 

``` go
func WrapSkip(skip int, err error, text string) error
```

WrapSkip wraps error with text. It returns nil if given err is nil. The parameter `skip` specifies the stack callers skipped amount. Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

##### func WrapSkipf 

``` go
func WrapSkipf(skip int, err error, format string, args ...interface{}) error
```

WrapSkipf wraps error with text that is formatted with given format and args. It returns nil if given err is nil. The parameter `skip` specifies the stack callers skipped amount. Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

##### func Wrapf 

``` go
func Wrapf(err error, format string, args ...interface{}) error
```

Wrapf returns an error annotating err with a stack trace at the point Wrapf is called, and the format specifier. It returns nil if given `err` is nil. Note that it does not lose the error code of wrapped error, as it inherits the error code from it.

### Types 

#### type Error 

``` go
type Error struct {
	// contains filtered or unexported fields
}
```

Error is custom error for additional features.

##### (*Error) Cause 

``` go
func (err *Error) Cause() error
```

Cause returns the root cause error.

##### (*Error) Code 

``` go
func (err *Error) Code() gcode.Code
```

Code returns the error code. It returns CodeNil if it has no error code.

##### (*Error) Current 

``` go
func (err *Error) Current() error
```

Current creates and returns the current level error. It returns nil if current level error is nil.

##### (*Error) Equal <-2.1.0

``` go
func (err *Error) Equal(target error) bool
```

Equal reports whether current error `err` equals to error `target`. Please note that, in default comparison for `Error`, the errors are considered the same if both the `code` and `text` of them are the same.

##### (*Error) Error 

``` go
func (err *Error) Error() string
```

Error implements the interface of Error, it returns all the error as string.

##### (*Error) Format 

``` go
func (err *Error) Format(s fmt.State, verb rune)
```

Format formats the frame according to the fmt.Formatter interface.

%v, %s : Print all the error string; %-v, %-s : Print current level error string; %+s : Print full stack error list; %+v : Print the error string and full stack error list

##### (*Error) Is <-2.1.0

``` go
func (err *Error) Is(target error) bool
```

Is reports whether current error `err` has error `target` in its chaining errors. It is just for implements for stdlib errors.Is from Go version 1.17.

##### (Error) MarshalJSON 

``` go
func (err Error) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that do not use pointer as its receiver here.

##### (*Error) SetCode 

``` go
func (err *Error) SetCode(code gcode.Code)
```

SetCode updates the internal code with given code.

##### (*Error) Stack 

``` go
func (err *Error) Stack() string
```

Stack returns the error stack information as string.

##### (*Error) Unwrap <-2.1.0

``` go
func (err *Error) Unwrap() error
```

Unwrap is alias of function `Next`. It is just for implements for stdlib errors.Unwrap from Go version 1.17.

#### type ICause <-2.1.3

``` go
type ICause interface {
	Error() string
	Cause() error
}
```

ICause is the interface for Cause feature.

#### type ICode <-2.1.3

``` go
type ICode interface {
	Error() string
	Code() gcode.Code
}
```

ICode is the interface for Code feature.

#### type ICurrent <-2.1.3

``` go
type ICurrent interface {
	Error() string
	Current() error
}
```

ICurrent is the interface for Current feature.

#### type IEqual <-2.1.3

``` go
type IEqual interface {
	Error() string
	Equal(target error) bool
}
```

IEqual is the interface for Equal feature.

#### type IIs <-2.1.3

``` go
type IIs interface {
	Error() string
	Is(target error) bool
}
```

IIs is the interface for Is feature.

#### type IStack <-2.1.3

``` go
type IStack interface {
	Error() string
	Stack() string
}
```

IStack is the interface for Stack feature.

#### type IUnwrap <-2.1.3

``` go
type IUnwrap interface {
	Error() string
	Unwrap() error
}
```

IUnwrap is the interface for Unwrap feature.

#### type Option 

``` go
type Option struct {
	Error error      // Wrapped error if any.
	Stack bool       // Whether recording stack information into error.
	Text  string     // Error text, which is created by New* functions.
	Code  gcode.Code // Error code if necessary.
}
```

Option is option for creating error.