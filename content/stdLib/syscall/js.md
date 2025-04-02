+++
title = "js"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++


> 原文：[https://pkg.go.dev/syscall/js@go1.24.2](https://pkg.go.dev/syscall/js@go1.24.2)

Package js gives access to the WebAssembly host environment when using the js/wasm architecture. Its API is based on JavaScript semantics.

​	js 包在使用 js/wasm 架构时提供对 WebAssembly 主机环境的访问。其 API 基于 JavaScript 语义。

This package is EXPERIMENTAL. Its current scope is only to allow tests to run, but not yet to provide a comprehensive API for users. It is exempt from the Go compatibility promise.

​	此包为实验性包。其当前范围仅允许运行测试，但尚未为用户提供全面的 API。它不受 Go 兼容性承诺的约束。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func CopyBytesToGo <- go1.13

```go
func CopyBytesToGo(dst []byte, src Value) int
```

CopyBytesToGo copies bytes from src to dst. It panics if src is not an Uint8Array or Uint8ClampedArray. It returns the number of bytes copied, which will be the minimum of the lengths of src and dst.

​	CopyBytesToGo 将字节从 src 复制到 dst。如果 src 不是 Uint8Array 或 Uint8ClampedArray，它会引发 panic。它返回已复制的字节数，该字节数将是 src 和 dst 长度的最小值。

### func CopyBytesToJS <- go1.13

```go
func CopyBytesToJS(dst Value, src []byte) int
```

CopyBytesToJS copies bytes from src to dst. It panics if dst is not an Uint8Array or Uint8ClampedArray. It returns the number of bytes copied, which will be the minimum of the lengths of src and dst.

​	CopyBytesToJS 将字节从 src 复制到 dst。如果 dst 不是 Uint8Array 或 Uint8ClampedArray，它会引发 panic。它返回已复制的字节数，该字节数将是 src 和 dst 长度的最小值。

## 类型

### type Error

```go
type Error struct {
	// Value is the underlying JavaScript error value.
	Value
}
```

Error wraps a JavaScript error.

​	Error 包装了一个 JavaScript 错误。

#### (Error) Error

```go
func (e Error) Error() string
```

Error implements the error interface.

​	Error 实现错误接口。

### type Func <- go1.12

```go
type Func struct {
	Value // the JavaScript function that invokes the Go function
	// contains filtered or unexported fields
}
```

Func is a wrapped Go function to be called by JavaScript.

​	Func 是一个包装的 Go 函数，可由 JavaScript 调用。

#### func FuncOf <- go1.12

```go
func FuncOf(fn func(this Value, args []Value) any) Func
```

FuncOf returns a function to be used by JavaScript.

​	FuncOf 返回一个供 JavaScript 使用的函数。

The Go function fn is called with the value of JavaScript’s “this” keyword and the arguments of the invocation. The return value of the invocation is the result of the Go function mapped back to JavaScript according to ValueOf.

​	Go 函数 fn 使用 JavaScript 的“this”关键字的值和调用的参数进行调用。调用的返回值是根据 ValueOf 映射回 JavaScript 的 Go 函数的结果。

Invoking the wrapped Go function from JavaScript will pause the event loop and spawn a new goroutine. Other wrapped functions which are triggered during a call from Go to JavaScript get executed on the same goroutine.

​	从 JavaScript 调用包装的 Go 函数将暂停事件循环并生成一个新的 goroutine。在从 Go 到 JavaScript 的调用期间触发的其他包装函数在同一个 goroutine 上执行。

As a consequence, if one wrapped function blocks, JavaScript’s event loop is blocked until that function returns. Hence, calling any async JavaScript API, which requires the event loop, like fetch (http.Client), will cause an immediate deadlock. Therefore a blocking function should explicitly start a new goroutine.

​	因此，如果一个包装函数阻塞，JavaScript 的事件循环将阻塞，直到该函数返回。因此，调用任何需要事件循环的异步 JavaScript API，如 fetch (http.Client)，将导致立即死锁。因此，阻塞函数应显式启动一个新的 goroutine。

Func.Release must be called to free up resources when the function will not be invoked any more.

​	当不再调用函数时，必须调用 Func.Release 以释放资源。

##### Example

```go
var cb js.Func
cb = js.FuncOf(func(this js.Value, args []js.Value) any {
	fmt.Println("button clicked")
	cb.Release() // release the function if the button will not be clicked again
	return nil
})
js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)

// Output:
```



#### (Func) Release <- go1.12

```go
func (c Func) Release()
```

Release frees up resources allocated for the function. The function must not be invoked after calling Release. It is allowed to call Release while the function is still running.

​	Release 释放为函数分配的资源。调用 Release 后，不得再调用该函数。允许在函数仍在运行时调用 Release。

### type Type

```go
type Type int
```

Type represents the JavaScript type of a Value.

​	Type 表示 Value 的 JavaScript 类型。

```go
const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)
```

#### (Type) String

```go
func (t Type) String() string
```

### type Value

```go
type Value struct {
	// contains filtered or unexported fields
}
```

Value represents a JavaScript value. The zero value is the JavaScript value “undefined”. Values can be checked for equality with the Equal method.

​	Value 表示 JavaScript 值。零值是 JavaScript 值“undefined”。可以使用 Equal 方法检查值是否相等。

#### func Global

```go
func Global() Value
```

Global returns the JavaScript global object, usually “window” or “global”.

​	Global 返回 JavaScript 全局对象，通常为“window”或“global”。

#### func Null

```go
func Null() Value
```

Null returns the JavaScript value “null”.

​	Null 返回 JavaScript 值“null”。

#### func Undefined

```go
func Undefined() Value
```

Undefined returns the JavaScript value “undefined”.

​	Undefined 返回 JavaScript 值“undefined”。

#### func ValueOf

```go
func ValueOf(x any) Value
```

ValueOf returns x as a JavaScript value:

​	ValueOf 将 x 作为 JavaScript 值返回：

```
| Go                     | JavaScript             |
| ---------------------- | ---------------------- |
| js.Value               | [its value]            |
| js.Func                | function               |
| nil                    | null                   |
| bool                   | boolean                |
| integers and floats    | number                 |
| string                 | string                 |
| []interface{}          | new array              |
| map[string]interface{} | new object             |
```

Panics if x is not one of the expected types.

​	如果 x 不是预期类型之一，则会引发恐慌。

#### (Value) Bool

```go
func (v Value) Bool() bool
```

Bool returns the value v as a bool. It panics if v is not a JavaScript boolean.

​	Bool 将值 v 作为布尔值返回。如果 v 不是 JavaScript 布尔值，则会引发恐慌。

#### (Value) Call

```go
func (v Value) Call(m string, args ...any) Value
```

Call does a JavaScript call to the method m of value v with the given arguments. It panics if v has no method m. The arguments get mapped to JavaScript values according to the ValueOf function.

​	Call 使用给定的参数对值 v 的方法 m 执行 JavaScript 调用。如果 v 没有方法 m，则会引发恐慌。根据 ValueOf 函数，参数会映射到 JavaScript 值。

#### (Value) Delete <- go1.14

```go
func (v Value) Delete(p string)
```

Delete deletes the JavaScript property p of value v. It panics if v is not a JavaScript object.

​	Delete 删除值 v 的 JavaScript 属性 p。如果 v 不是 JavaScript 对象，它会引发恐慌。

#### (Value) Equal <- go1.14

```go
func (v Value) Equal(w Value) bool
```

Equal reports whether v and w are equal according to JavaScript’s === operator.

​	Equal 报告 v 和 w 是否根据 JavaScript 的 === 运算符相等。

#### (Value) Float

```go
func (v Value) Float() float64
```

Float returns the value v as a float64. It panics if v is not a JavaScript number.

​	Float 将值 v 作为 float64 返回。如果 v 不是 JavaScript 数字，它会引发恐慌。

#### (Value) Get

```go
func (v Value) Get(p string) Value
```

Get returns the JavaScript property p of value v. It panics if v is not a JavaScript object.

​	Get 返回值 v 的 JavaScript 属性 p。如果 v 不是 JavaScript 对象，它会引发恐慌。

#### (Value) Index

```go
func (v Value) Index(i int) Value
```

Index returns JavaScript index i of value v. It panics if v is not a JavaScript object.

​	Index 返回值 v 的 JavaScript 索引 i。如果 v 不是 JavaScript 对象，它会引发恐慌。

#### (Value) InstanceOf

```go
func (v Value) InstanceOf(t Value) bool
```

InstanceOf reports whether v is an instance of type t according to JavaScript’s instanceof operator.

​	InstanceOf 报告 v 是否根据 JavaScript 的 instanceof 运算符是类型 t 的实例。

#### (Value) Int

```go
func (v Value) Int() int
```

Int returns the value v truncated to an int. It panics if v is not a JavaScript number.

​	Int 将值 v 截断为 int。如果 v 不是 JavaScript 数字，则会引发 panic。

#### (Value) Invoke

```go
func (v Value) Invoke(args ...any) Value
```

Invoke does a JavaScript call of the value v with the given arguments. It panics if v is not a JavaScript function. The arguments get mapped to JavaScript values according to the ValueOf function.

​	Invoke 使用给定的参数对值 v 进行 JavaScript 调用。如果 v 不是 JavaScript 函数，则会引发 panic。参数会根据 ValueOf 函数映射到 JavaScript 值。

#### (Value) IsNaN <- go1.14

```go
func (v Value) IsNaN() bool
```

IsNaN reports whether v is the JavaScript value “NaN”.

​	IsNaN 报告 v 是否是 JavaScript 值“NaN”。

#### (Value) IsNull <- go1.14

```go
func (v Value) IsNull() bool
```

IsNull reports whether v is the JavaScript value “null”.

​	IsNull 报告 v 是否是 JavaScript 值“null”。

#### (Value) IsUndefined <- go1.14

```go
func (v Value) IsUndefined() bool
```

IsUndefined reports whether v is the JavaScript value “undefined”.

​	IsUndefined 报告 v 是否为 JavaScript 值“undefined”。

#### (Value) Length

```go
func (v Value) Length() int
```

Length returns the JavaScript property “length” of v. It panics if v is not a JavaScript object.

​	Length 返回 v 的 JavaScript 属性“length”。如果 v 不是 JavaScript 对象，它会引发恐慌。

#### (Value) New

```go
func (v Value) New(args ...any) Value
```

New uses JavaScript’s “new” operator with value v as constructor and the given arguments. It panics if v is not a JavaScript function. The arguments get mapped to JavaScript values according to the ValueOf function.

​	New 使用 JavaScript 的“new”运算符，其中值 v 作为构造函数，并给出参数。如果 v 不是 JavaScript 函数，它会引发恐慌。参数根据 ValueOf 函数映射到 JavaScript 值。

#### (Value) Set

```go
func (v Value) Set(p string, x any)
```

Set sets the JavaScript property p of value v to ValueOf(x). It panics if v is not a JavaScript object.

​	Set 将值 v 的 JavaScript 属性 p 设置为 ValueOf(x)。如果 v 不是 JavaScript 对象，它会引发恐慌。

#### (Value) SetIndex

```go
func (v Value) SetIndex(i int, x any)
```

SetIndex sets the JavaScript index i of value v to ValueOf(x). It panics if v is not a JavaScript object.

​	SetIndex 将值 v 的 JavaScript 索引 i 设置为 ValueOf(x)。如果 v 不是 JavaScript 对象，它会引发恐慌。

#### (Value) String

```go
func (v Value) String() string
```

String returns the value v as a string. String is a special case because of Go’s String method convention. Unlike the other getters, it does not panic if v’s Type is not TypeString. Instead, it returns a string of the form “” or “<T: V>” where T is v’s type and V is a string representation of v’s value.

​	String 将值 v 作为字符串返回。String 是一个特殊情况，因为 Go 的 String 方法约定。与其他 getter 不同，如果 v 的 Type 不是 TypeString，它不会引发 panic。相反，它返回 “” 或 “” 形式的字符串，其中 T 是 v 的类型，V 是 v 的值的字符串表示形式。

#### (Value) Truthy <- go1.12

```go
func (v Value) Truthy() bool
```

Truthy returns the JavaScript “truthiness” of the value v. In JavaScript, false, 0, “”, null, undefined, and NaN are “falsy”, and everything else is “truthy”. See https://developer.mozilla.org/en-US/docs/Glossary/Truthy.

​	Truthy 返回值 v 的 JavaScript “真值性”。在 JavaScript 中，false、0、“”、null、undefined 和 NaN 是 “假值”，其他一切都是 “真值”。请参阅 https://developer.mozilla.org/en-US/docs/Glossary/Truthy。

#### (Value) Type

```go
func (v Value) Type() Type
```

Type returns the JavaScript type of the value v. It is similar to JavaScript’s typeof operator, except that it returns TypeNull instead of TypeObject for null.

​	Type 返回值 v 的 JavaScript 类型。它类似于 JavaScript 的 typeof 运算符，除了它为 null 返回 TypeNull 而不是 TypeObject。

### type ValueError

```go
type ValueError struct {
	Method string
	Type   Type
}
```

A ValueError occurs when a Value method is invoked on a Value that does not support it. Such cases are documented in the description of each method.

​	当对不支持它的 Value 调用 Value 方法时，就会发生 ValueError。此类情况在每个方法的说明中都有记录。

#### (*ValueError) Error 

``` go 
func (e *ValueError) Error() string
```