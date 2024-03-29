+++
title = "gcode"
date = 2024-03-21T17:50:56+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/errors/gcode](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/errors/gcode)

Package gcode provides universal error code definition and common error codes implements.

​	软件包 gcode 提供了通用的错误代码定义和常见的错误代码实现。

## 常量

This section is empty.

## 变量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/errors/gcode/gcode.go#L28)

```go
var (
	CodeNil                       = localCode{-1, "", nil}                             // No error code specified.
	CodeOK                        = localCode{0, "OK", nil}                            // It is OK.
	CodeInternalError             = localCode{50, "Internal Error", nil}               // An error occurred internally.
	CodeValidationFailed          = localCode{51, "Validation Failed", nil}            // Data validation failed.
	CodeDbOperationError          = localCode{52, "Database Operation Error", nil}     // Database operation error.
	CodeInvalidParameter          = localCode{53, "Invalid Parameter", nil}            // The given parameter for current operation is invalid.
	CodeMissingParameter          = localCode{54, "Missing Parameter", nil}            // Parameter for current operation is missing.
	CodeInvalidOperation          = localCode{55, "Invalid Operation", nil}            // The function cannot be used like this.
	CodeInvalidConfiguration      = localCode{56, "Invalid Configuration", nil}        // The configuration is invalid for current operation.
	CodeMissingConfiguration      = localCode{57, "Missing Configuration", nil}        // The configuration is missing for current operation.
	CodeNotImplemented            = localCode{58, "Not Implemented", nil}              // The operation is not implemented yet.
	CodeNotSupported              = localCode{59, "Not Supported", nil}                // The operation is not supported yet.
	CodeOperationFailed           = localCode{60, "Operation Failed", nil}             // I tried, but I cannot give you what you want.
	CodeNotAuthorized             = localCode{61, "Not Authorized", nil}               // Not Authorized.
	CodeSecurityReason            = localCode{62, "Security Reason", nil}              // Security Reason.
	CodeServerBusy                = localCode{63, "Server Is Busy", nil}               // Server is busy, please try again later.
	CodeUnknown                   = localCode{64, "Unknown Error", nil}                // Unknown error.
	CodeNotFound                  = localCode{65, "Not Found", nil}                    // Resource does not exist.
	CodeInvalidRequest            = localCode{66, "Invalid Request", nil}              // Invalid request.
	CodeNecessaryPackageNotImport = localCode{67, "Necessary Package Not Import", nil} // It needs necessary package import.
	CodeInternalPanic             = localCode{68, "Internal Panic", nil}               // An panic occurred internally.
	CodeBusinessValidationFailed  = localCode{300, "Business Validation Failed", nil}  // Business validation failed.
)
```

## 函数

This section is empty.

## 类型

### type Code

```go
type Code interface {
	// Code returns the integer number of current error code.
	Code() int

	// Message returns the brief message for current error code.
	Message() string

	// Detail returns the detailed information of current error code,
	// which is mainly designed as an extension field for error code.
	Detail() interface{}
}
```

Code is universal error code interface definition.

​	代码是通用的错误代码接口定义。

#### func New

```go
func New(code int, message string, detail interface{}) Code
```

New creates and returns an error code. Note that it returns an interface object of Code.

​	New 创建并返回错误代码。请注意，它返回 Code 的接口对象。

#### func WithCode

```go
func WithCode(code Code, detail interface{}) Code
```

WithCode creates and returns a new error code based on given Code. The code and message is from given `code`, but the detail if from given `detail`.

​	WithCode 根据给定的代码创建并返回新的错误代码。代码和消息来自给定 `code` 的，但细节如果来自给定 `detail` 的。