+++
title = "gtype"
date = 2024-03-21T17:45:17+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gtype](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gtype)

Package gtype provides high performance and concurrent-safe basic variable types.

​	软件包 gtype 提供高性能和并发安全的基本变量类型。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Bool

```go
type Bool struct {
	// contains filtered or unexported fields
}
```

Bool is a struct for concurrent-safe operation for type bool.

​	Bool 是用于 bool 类型并发安全操作的结构。

#### func NewBool

```go
func NewBool(value ...bool) *Bool
```

NewBool creates and returns a concurrent-safe object for bool type, with given initial value `value`.

​	NewBool 创建并返回一个 bool 类型的并发安全对象，其初始值 `value` 为 。

#### (*Bool) Cas

```go
func (v *Bool) Cas(old, new bool) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Bool) Clone

```go
func (v *Bool) Clone() *Bool
```

Clone clones and returns a new concurrent-safe object for bool type.

​	克隆克隆并返回 bool 类型的新并发安全对象。

#### (*Bool) DeepCopy

```go
func (v *Bool) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Bool) MarshalJSON

```go
func (v Bool) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Bool) Set

```go
func (v *Bool) Set(value bool) (old bool)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Bool) String

```go
func (v *Bool) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Bool) UnmarshalJSON

```go
func (v *Bool) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Bool) UnmarshalValue

```go
func (v *Bool) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Bool) Val

```go
func (v *Bool) Val() bool
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Byte

```go
type Byte struct {
	// contains filtered or unexported fields
}
```

Byte is a struct for concurrent-safe operation for type byte.

​	Byte 是用于类型 byte 的并发安全操作的结构。

#### func NewByte

```go
func NewByte(value ...byte) *Byte
```

NewByte creates and returns a concurrent-safe object for byte type, with given initial value `value`.

​	NewByte 创建并返回字节类型的并发安全对象，其初始值 `value` 为 。

#### (*Byte) Add

```go
func (v *Byte) Add(delta byte) (new byte)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Byte) Cas

```go
func (v *Byte) Cas(old, new byte) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Byte) Clone

```go
func (v *Byte) Clone() *Byte
```

Clone clones and returns a new concurrent-safe object for byte type.

​	克隆克隆并返回字节类型的新并发安全对象。

#### (*Byte) DeepCopy

```go
func (v *Byte) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Byte) MarshalJSON

```go
func (v Byte) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Byte) Set

```go
func (v *Byte) Set(value byte) (old byte)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Byte) String

```go
func (v *Byte) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Byte) UnmarshalJSON

```go
func (v *Byte) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Byte) UnmarshalValue

```go
func (v *Byte) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Byte) Val

```go
func (v *Byte) Val() byte
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Bytes

```go
type Bytes struct {
	// contains filtered or unexported fields
}
```

Bytes is a struct for concurrent-safe operation for type []byte.

​	Bytes 是用于 []byte 类型的并发安全操作的结构。

#### func NewBytes

```go
func NewBytes(value ...[]byte) *Bytes
```

NewBytes creates and returns a concurrent-safe object for []byte type, with given initial value `value`.

​	NewBytes 为 []byte 类型创建并返回一个并发安全对象，其初始值 `value` 为 。

#### (*Bytes) Clone

```go
func (v *Bytes) Clone() *Bytes
```

Clone clones and returns a new shallow copy object for []byte type.

​	克隆克隆并返回 []byte 类型的新浅层复制对象。

#### (*Bytes) DeepCopy

```go
func (v *Bytes) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Bytes) MarshalJSON

```go
func (v Bytes) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Bytes) Set

```go
func (v *Bytes) Set(value []byte) (old []byte)
```

Set atomically stores `value` into t.value and returns the previous value of t.value. Note: The parameter `value` cannot be nil.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。注意：该参数 `value` 不能为零。

#### (*Bytes) String

```go
func (v *Bytes) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Bytes) UnmarshalJSON

```go
func (v *Bytes) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Bytes) UnmarshalValue

```go
func (v *Bytes) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Bytes) Val

```go
func (v *Bytes) Val() []byte
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Float32

```go
type Float32 struct {
	// contains filtered or unexported fields
}
```

Float32 is a struct for concurrent-safe operation for type float32.

​	float32 是用于 float32 类型的并发安全操作的结构。

#### func NewFloat32

```go
func NewFloat32(value ...float32) *Float32
```

NewFloat32 creates and returns a concurrent-safe object for float32 type, with given initial value `value`.

​	NewFloat32 为 float32 类型创建并返回一个并发安全对象，该对象具有给定的初始值 `value` 。

#### (*Float32) Add

```go
func (v *Float32) Add(delta float32) (new float32)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Float32) Cas

```go
func (v *Float32) Cas(old, new float32) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Float32) Clone

```go
func (v *Float32) Clone() *Float32
```

Clone clones and returns a new concurrent-safe object for float32 type.

​	克隆克隆并返回 float32 类型的新并发安全对象。

#### (*Float32) DeepCopy

```go
func (v *Float32) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Float32) MarshalJSON

```go
func (v Float32) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Float32) Set

```go
func (v *Float32) Set(value float32) (old float32)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Float32) String

```go
func (v *Float32) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Float32) UnmarshalJSON

```go
func (v *Float32) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Float32) UnmarshalValue

```go
func (v *Float32) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Float32) Val

```go
func (v *Float32) Val() float32
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Float64

```go
type Float64 struct {
	// contains filtered or unexported fields
}
```

Float64 is a struct for concurrent-safe operation for type float64.

​	float64 是用于 float64 类型的并发安全操作的结构。

#### func NewFloat64

```go
func NewFloat64(value ...float64) *Float64
```

NewFloat64 creates and returns a concurrent-safe object for float64 type, with given initial value `value`.

​	NewFloat64 为 float64 类型创建并返回一个并发安全对象，该对象具有给定的初始值 `value` 。

#### (*Float64) Add

```go
func (v *Float64) Add(delta float64) (new float64)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Float64) Cas

```go
func (v *Float64) Cas(old, new float64) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Float64) Clone

```go
func (v *Float64) Clone() *Float64
```

Clone clones and returns a new concurrent-safe object for float64 type.

​	克隆克隆并返回 float64 类型的新并发安全对象。

#### (*Float64) DeepCopy

```go
func (v *Float64) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Float64) MarshalJSON

```go
func (v Float64) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Float64) Set

```go
func (v *Float64) Set(value float64) (old float64)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Float64) String

```go
func (v *Float64) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Float64) UnmarshalJSON

```go
func (v *Float64) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Float64) UnmarshalValue

```go
func (v *Float64) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Float64) Val

```go
func (v *Float64) Val() float64
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Int

```go
type Int struct {
	// contains filtered or unexported fields
}
```

Int is a struct for concurrent-safe operation for type int.

​	Int 是用于 int 类型的并发安全操作的结构。

#### func NewInt

```go
func NewInt(value ...int) *Int
```

NewInt creates and returns a concurrent-safe object for int type, with given initial value `value`.

​	NewInt 创建并返回 int 类型的并发安全对象，其初始值 `value` 为 。

#### (*Int) Add

```go
func (v *Int) Add(delta int) (new int)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Int) Cas

```go
func (v *Int) Cas(old, new int) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Int) Clone

```go
func (v *Int) Clone() *Int
```

Clone clones and returns a new concurrent-safe object for int type.

​	克隆克隆并返回 int 类型的新并发安全对象。

#### (*Int) DeepCopy

```go
func (v *Int) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Int) MarshalJSON

```go
func (v Int) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Int) Set

```go
func (v *Int) Set(value int) (old int)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Int) String

```go
func (v *Int) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Int) UnmarshalJSON

```go
func (v *Int) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Int) UnmarshalValue

```go
func (v *Int) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Int) Val

```go
func (v *Int) Val() int
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Int32

```go
type Int32 struct {
	// contains filtered or unexported fields
}
```

Int32 is a struct for concurrent-safe operation for type int32.

​	Int32 是用于 int32 类型的并发安全操作的结构。

#### func NewInt32

```go
func NewInt32(value ...int32) *Int32
```

NewInt32 creates and returns a concurrent-safe object for int32 type, with given initial value `value`.

​	NewInt32 创建并返回 int32 类型的并发安全对象，其初始值 `value` 为 。

#### (*Int32) Add

```go
func (v *Int32) Add(delta int32) (new int32)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Int32) Cas

```go
func (v *Int32) Cas(old, new int32) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Int32) Clone

```go
func (v *Int32) Clone() *Int32
```

Clone clones and returns a new concurrent-safe object for int32 type.

​	克隆克隆并返回 int32 类型的新并发安全对象。

#### (*Int32) DeepCopy

```go
func (v *Int32) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Int32) MarshalJSON

```go
func (v Int32) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Int32) Set

```go
func (v *Int32) Set(value int32) (old int32)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Int32) String

```go
func (v *Int32) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Int32) UnmarshalJSON

```go
func (v *Int32) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Int32) UnmarshalValue

```go
func (v *Int32) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Int32) Val

```go
func (v *Int32) Val() int32
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Int64

```go
type Int64 struct {
	// contains filtered or unexported fields
}
```

Int64 is a struct for concurrent-safe operation for type int64.

​	Int64 是用于 int64 类型的并发安全操作的结构。

#### func NewInt64

```go
func NewInt64(value ...int64) *Int64
```

NewInt64 creates and returns a concurrent-safe object for int64 type, with given initial value `value`.

​	NewInt64 创建并返回 int64 类型的并发安全对象，其初始值 `value` 为 。

#### (*Int64) Add

```go
func (v *Int64) Add(delta int64) (new int64)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Int64) Cas

```go
func (v *Int64) Cas(old, new int64) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Int64) Clone

```go
func (v *Int64) Clone() *Int64
```

Clone clones and returns a new concurrent-safe object for int64 type.

​	克隆克隆并返回 int64 类型的新并发安全对象。

#### (*Int64) DeepCopy

```go
func (v *Int64) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Int64) MarshalJSON

```go
func (v Int64) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Int64) Set

```go
func (v *Int64) Set(value int64) (old int64)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Int64) String

```go
func (v *Int64) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Int64) UnmarshalJSON

```go
func (v *Int64) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Int64) UnmarshalValue

```go
func (v *Int64) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Int64) Val

```go
func (v *Int64) Val() int64
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Interface

```go
type Interface struct {
	// contains filtered or unexported fields
}
```

Interface is a struct for concurrent-safe operation for type interface{}.

​	interface 是 interface{} 类型的并发安全操作的结构。

#### func New

```go
func New(value ...interface{}) *Interface
```

New is alias of NewInterface. See NewInterface.

​	New 是 NewInterface 的别名。请参阅 NewInterface。

#### func NewInterface

```go
func NewInterface(value ...interface{}) *Interface
```

NewInterface creates and returns a concurrent-safe object for interface{} type, with given initial value `value`.

​	NewInterface 为 interface{} 类型创建并返回一个并发安全对象，该对象具有给定的初始值 `value` 。

#### (*Interface) Clone

```go
func (v *Interface) Clone() *Interface
```

Clone clones and returns a new concurrent-safe object for interface{} type.

​	克隆克隆并返回 interface{} 类型的新并发安全对象。

#### (*Interface) DeepCopy

```go
func (v *Interface) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Interface) MarshalJSON

```go
func (v Interface) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Interface) Set

```go
func (v *Interface) Set(value interface{}) (old interface{})
```

Set atomically stores `value` into t.value and returns the previous value of t.value. Note: The parameter `value` cannot be nil.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。注意：该参数 `value` 不能为零。

#### (*Interface) String

```go
func (v *Interface) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Interface) UnmarshalJSON

```go
func (v *Interface) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Interface) UnmarshalValue

```go
func (v *Interface) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Interface) Val

```go
func (v *Interface) Val() interface{}
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type String

```go
type String struct {
	// contains filtered or unexported fields
}
```

String is a struct for concurrent-safe operation for type string.

​	String 是用于字符串类型并发安全操作的结构。

#### func NewString

```go
func NewString(value ...string) *String
```

NewString creates and returns a concurrent-safe object for string type, with given initial value `value`.

​	NewString 创建并返回字符串类型的并发安全对象，该对象具有给定的初始值 `value` 。

#### (*String) Clone

```go
func (v *String) Clone() *String
```

Clone clones and returns a new concurrent-safe object for string type.

​	克隆克隆并返回字符串类型的新并发安全对象。

#### (*String) DeepCopy

```go
func (v *String) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (String) MarshalJSON

```go
func (v String) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*String) Set

```go
func (v *String) Set(value string) (old string)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*String) String

```go
func (v *String) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*String) UnmarshalJSON

```go
func (v *String) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*String) UnmarshalValue

```go
func (v *String) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*String) Val

```go
func (v *String) Val() string
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Uint

```go
type Uint struct {
	// contains filtered or unexported fields
}
```

Uint is a struct for concurrent-safe operation for type uint.

​	Uint 是用于 uint 类型并发安全操作的结构。

#### func NewUint

```go
func NewUint(value ...uint) *Uint
```

NewUint creates and returns a concurrent-safe object for uint type, with given initial value `value`.

​	NewUint 为 uint 类型创建并返回一个并发安全对象，该对象具有给定的初始值 `value` 。

#### (*Uint) Add

```go
func (v *Uint) Add(delta uint) (new uint)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Uint) Cas

```go
func (v *Uint) Cas(old, new uint) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Uint) Clone

```go
func (v *Uint) Clone() *Uint
```

Clone clones and returns a new concurrent-safe object for uint type.

​	克隆克隆并返回 uint 类型的新并发安全对象。

#### (*Uint) DeepCopy

```go
func (v *Uint) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Uint) MarshalJSON

```go
func (v Uint) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Uint) Set

```go
func (v *Uint) Set(value uint) (old uint)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Uint) String

```go
func (v *Uint) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Uint) UnmarshalJSON

```go
func (v *Uint) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Uint) UnmarshalValue

```go
func (v *Uint) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Uint) Val

```go
func (v *Uint) Val() uint
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Uint32

```go
type Uint32 struct {
	// contains filtered or unexported fields
}
```

Uint32 is a struct for concurrent-safe operation for type uint32.

​	Uint32 是用于 uint32 类型的并发安全操作的结构。

#### func NewUint32

```go
func NewUint32(value ...uint32) *Uint32
```

NewUint32 creates and returns a concurrent-safe object for uint32 type, with given initial value `value`.

​	NewUint32 为 uint32 类型创建并返回一个并发安全对象，该对象具有给定的初始值 `value` 。

#### (*Uint32) Add

```go
func (v *Uint32) Add(delta uint32) (new uint32)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Uint32) Cas

```go
func (v *Uint32) Cas(old, new uint32) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Uint32) Clone

```go
func (v *Uint32) Clone() *Uint32
```

Clone clones and returns a new concurrent-safe object for uint32 type.

​	克隆克隆并返回 uint32 类型的新并发安全对象。

#### (*Uint32) DeepCopy

```go
func (v *Uint32) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Uint32) MarshalJSON

```go
func (v Uint32) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Uint32) Set

```go
func (v *Uint32) Set(value uint32) (old uint32)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Uint32) String

```go
func (v *Uint32) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Uint32) UnmarshalJSON

```go
func (v *Uint32) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Uint32) UnmarshalValue

```go
func (v *Uint32) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Uint32) Val

```go
func (v *Uint32) Val() uint32
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。

### type Uint64

```go
type Uint64 struct {
	// contains filtered or unexported fields
}
```

Uint64 is a struct for concurrent-safe operation for type uint64.

​	Uint64 是用于 uint64 类型的并发安全操作的结构。

#### func NewUint64

```go
func NewUint64(value ...uint64) *Uint64
```

NewUint64 creates and returns a concurrent-safe object for uint64 type, with given initial value `value`.

​	NewUint64 为 uint64 类型创建并返回一个并发安全对象，该对象具有给定的初始值 `value` 。

#### (*Uint64) Add

```go
func (v *Uint64) Add(delta uint64) (new uint64)
```

Add atomically adds `delta` to t.value and returns the new value.

​	以原子方式添加 `delta` to t.value 并返回新值。

#### (*Uint64) Cas

```go
func (v *Uint64) Cas(old, new uint64) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

​	Cas 执行值的比较和交换操作。

#### (*Uint64) Clone

```go
func (v *Uint64) Clone() *Uint64
```

Clone clones and returns a new concurrent-safe object for uint64 type.

​	克隆克隆并返回 uint64 类型的新并发安全对象。

#### (*Uint64) DeepCopy

```go
func (v *Uint64) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (Uint64) MarshalJSON

```go
func (v Uint64) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*Uint64) Set

```go
func (v *Uint64) Set(value uint64) (old uint64)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

​	设置以原子方式存储 `value` 到 t.value 中，并返回 t.value 的上一个值。

#### (*Uint64) String

```go
func (v *Uint64) String() string
```

String implements String interface for string printing.

​	String 实现用于字符串打印的 String 接口。

#### (*Uint64) UnmarshalJSON

```go
func (v *Uint64) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

#### (*Uint64) UnmarshalValue

```go
func (v *Uint64) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

​	UnmarshalValue 是一个接口实现，它为 `v` .

#### (*Uint64) Val

```go
func (v *Uint64) Val() uint64
```

Val atomically loads and returns t.value.

​	Val 以原子方式加载并返回 t.value。