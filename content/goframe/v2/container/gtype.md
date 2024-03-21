+++
title = "gtype"
date = 2024-03-21T17:45:17+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/container/gtype

Package gtype provides high performance and concurrent-safe basic variable types.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type Bool 

``` go
type Bool struct {
	// contains filtered or unexported fields
}
```

Bool is a struct for concurrent-safe operation for type bool.

##### func NewBool 

``` go
func NewBool(value ...bool) *Bool
```

NewBool creates and returns a concurrent-safe object for bool type, with given initial value `value`.

##### (*Bool) Cas 

``` go
func (v *Bool) Cas(old, new bool) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Bool) Clone 

``` go
func (v *Bool) Clone() *Bool
```

Clone clones and returns a new concurrent-safe object for bool type.

##### (*Bool) DeepCopy <-2.1.0

``` go
func (v *Bool) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Bool) MarshalJSON 

``` go
func (v Bool) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Bool) Set 

``` go
func (v *Bool) Set(value bool) (old bool)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Bool) String 

``` go
func (v *Bool) String() string
```

String implements String interface for string printing.

##### (*Bool) UnmarshalJSON 

``` go
func (v *Bool) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Bool) UnmarshalValue 

``` go
func (v *Bool) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Bool) Val 

``` go
func (v *Bool) Val() bool
```

Val atomically loads and returns t.value.

#### type Byte 

``` go
type Byte struct {
	// contains filtered or unexported fields
}
```

Byte is a struct for concurrent-safe operation for type byte.

##### func NewByte 

``` go
func NewByte(value ...byte) *Byte
```

NewByte creates and returns a concurrent-safe object for byte type, with given initial value `value`.

##### (*Byte) Add 

``` go
func (v *Byte) Add(delta byte) (new byte)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Byte) Cas 

``` go
func (v *Byte) Cas(old, new byte) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Byte) Clone 

``` go
func (v *Byte) Clone() *Byte
```

Clone clones and returns a new concurrent-safe object for byte type.

##### (*Byte) DeepCopy <-2.1.0

``` go
func (v *Byte) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Byte) MarshalJSON 

``` go
func (v Byte) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Byte) Set 

``` go
func (v *Byte) Set(value byte) (old byte)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Byte) String 

``` go
func (v *Byte) String() string
```

String implements String interface for string printing.

##### (*Byte) UnmarshalJSON 

``` go
func (v *Byte) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Byte) UnmarshalValue 

``` go
func (v *Byte) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Byte) Val 

``` go
func (v *Byte) Val() byte
```

Val atomically loads and returns t.value.

#### type Bytes 

``` go
type Bytes struct {
	// contains filtered or unexported fields
}
```

Bytes is a struct for concurrent-safe operation for type []byte.

##### func NewBytes 

``` go
func NewBytes(value ...[]byte) *Bytes
```

NewBytes creates and returns a concurrent-safe object for []byte type, with given initial value `value`.

##### (*Bytes) Clone 

``` go
func (v *Bytes) Clone() *Bytes
```

Clone clones and returns a new shallow copy object for []byte type.

##### (*Bytes) DeepCopy <-2.1.0

``` go
func (v *Bytes) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Bytes) MarshalJSON 

``` go
func (v Bytes) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Bytes) Set 

``` go
func (v *Bytes) Set(value []byte) (old []byte)
```

Set atomically stores `value` into t.value and returns the previous value of t.value. Note: The parameter `value` cannot be nil.

##### (*Bytes) String 

``` go
func (v *Bytes) String() string
```

String implements String interface for string printing.

##### (*Bytes) UnmarshalJSON 

``` go
func (v *Bytes) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Bytes) UnmarshalValue 

``` go
func (v *Bytes) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Bytes) Val 

``` go
func (v *Bytes) Val() []byte
```

Val atomically loads and returns t.value.

#### type Float32 

``` go
type Float32 struct {
	// contains filtered or unexported fields
}
```

Float32 is a struct for concurrent-safe operation for type float32.

##### func NewFloat32 

``` go
func NewFloat32(value ...float32) *Float32
```

NewFloat32 creates and returns a concurrent-safe object for float32 type, with given initial value `value`.

##### (*Float32) Add 

``` go
func (v *Float32) Add(delta float32) (new float32)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Float32) Cas 

``` go
func (v *Float32) Cas(old, new float32) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Float32) Clone 

``` go
func (v *Float32) Clone() *Float32
```

Clone clones and returns a new concurrent-safe object for float32 type.

##### (*Float32) DeepCopy <-2.1.0

``` go
func (v *Float32) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Float32) MarshalJSON 

``` go
func (v Float32) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Float32) Set 

``` go
func (v *Float32) Set(value float32) (old float32)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Float32) String 

``` go
func (v *Float32) String() string
```

String implements String interface for string printing.

##### (*Float32) UnmarshalJSON 

``` go
func (v *Float32) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Float32) UnmarshalValue 

``` go
func (v *Float32) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Float32) Val 

``` go
func (v *Float32) Val() float32
```

Val atomically loads and returns t.value.

#### type Float64 

``` go
type Float64 struct {
	// contains filtered or unexported fields
}
```

Float64 is a struct for concurrent-safe operation for type float64.

##### func NewFloat64 

``` go
func NewFloat64(value ...float64) *Float64
```

NewFloat64 creates and returns a concurrent-safe object for float64 type, with given initial value `value`.

##### (*Float64) Add 

``` go
func (v *Float64) Add(delta float64) (new float64)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Float64) Cas 

``` go
func (v *Float64) Cas(old, new float64) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Float64) Clone 

``` go
func (v *Float64) Clone() *Float64
```

Clone clones and returns a new concurrent-safe object for float64 type.

##### (*Float64) DeepCopy <-2.1.0

``` go
func (v *Float64) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Float64) MarshalJSON 

``` go
func (v Float64) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Float64) Set 

``` go
func (v *Float64) Set(value float64) (old float64)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Float64) String 

``` go
func (v *Float64) String() string
```

String implements String interface for string printing.

##### (*Float64) UnmarshalJSON 

``` go
func (v *Float64) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Float64) UnmarshalValue 

``` go
func (v *Float64) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Float64) Val 

``` go
func (v *Float64) Val() float64
```

Val atomically loads and returns t.value.

#### type Int 

``` go
type Int struct {
	// contains filtered or unexported fields
}
```

Int is a struct for concurrent-safe operation for type int.

##### func NewInt 

``` go
func NewInt(value ...int) *Int
```

NewInt creates and returns a concurrent-safe object for int type, with given initial value `value`.

##### (*Int) Add 

``` go
func (v *Int) Add(delta int) (new int)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Int) Cas 

``` go
func (v *Int) Cas(old, new int) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Int) Clone 

``` go
func (v *Int) Clone() *Int
```

Clone clones and returns a new concurrent-safe object for int type.

##### (*Int) DeepCopy <-2.1.0

``` go
func (v *Int) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Int) MarshalJSON 

``` go
func (v Int) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Int) Set 

``` go
func (v *Int) Set(value int) (old int)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Int) String 

``` go
func (v *Int) String() string
```

String implements String interface for string printing.

##### (*Int) UnmarshalJSON 

``` go
func (v *Int) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Int) UnmarshalValue 

``` go
func (v *Int) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Int) Val 

``` go
func (v *Int) Val() int
```

Val atomically loads and returns t.value.

#### type Int32 

``` go
type Int32 struct {
	// contains filtered or unexported fields
}
```

Int32 is a struct for concurrent-safe operation for type int32.

##### func NewInt32 

``` go
func NewInt32(value ...int32) *Int32
```

NewInt32 creates and returns a concurrent-safe object for int32 type, with given initial value `value`.

##### (*Int32) Add 

``` go
func (v *Int32) Add(delta int32) (new int32)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Int32) Cas 

``` go
func (v *Int32) Cas(old, new int32) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Int32) Clone 

``` go
func (v *Int32) Clone() *Int32
```

Clone clones and returns a new concurrent-safe object for int32 type.

##### (*Int32) DeepCopy <-2.1.0

``` go
func (v *Int32) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Int32) MarshalJSON 

``` go
func (v Int32) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Int32) Set 

``` go
func (v *Int32) Set(value int32) (old int32)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Int32) String 

``` go
func (v *Int32) String() string
```

String implements String interface for string printing.

##### (*Int32) UnmarshalJSON 

``` go
func (v *Int32) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Int32) UnmarshalValue 

``` go
func (v *Int32) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Int32) Val 

``` go
func (v *Int32) Val() int32
```

Val atomically loads and returns t.value.

#### type Int64 

``` go
type Int64 struct {
	// contains filtered or unexported fields
}
```

Int64 is a struct for concurrent-safe operation for type int64.

##### func NewInt64 

``` go
func NewInt64(value ...int64) *Int64
```

NewInt64 creates and returns a concurrent-safe object for int64 type, with given initial value `value`.

##### (*Int64) Add 

``` go
func (v *Int64) Add(delta int64) (new int64)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Int64) Cas 

``` go
func (v *Int64) Cas(old, new int64) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Int64) Clone 

``` go
func (v *Int64) Clone() *Int64
```

Clone clones and returns a new concurrent-safe object for int64 type.

##### (*Int64) DeepCopy <-2.1.0

``` go
func (v *Int64) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Int64) MarshalJSON 

``` go
func (v Int64) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Int64) Set 

``` go
func (v *Int64) Set(value int64) (old int64)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Int64) String 

``` go
func (v *Int64) String() string
```

String implements String interface for string printing.

##### (*Int64) UnmarshalJSON 

``` go
func (v *Int64) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Int64) UnmarshalValue 

``` go
func (v *Int64) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Int64) Val 

``` go
func (v *Int64) Val() int64
```

Val atomically loads and returns t.value.

#### type Interface 

``` go
type Interface struct {
	// contains filtered or unexported fields
}
```

Interface is a struct for concurrent-safe operation for type interface{}.

##### func New 

``` go
func New(value ...interface{}) *Interface
```

New is alias of NewInterface. See NewInterface.

##### func NewInterface 

``` go
func NewInterface(value ...interface{}) *Interface
```

NewInterface creates and returns a concurrent-safe object for interface{} type, with given initial value `value`.

##### (*Interface) Clone 

``` go
func (v *Interface) Clone() *Interface
```

Clone clones and returns a new concurrent-safe object for interface{} type.

##### (*Interface) DeepCopy <-2.1.0

``` go
func (v *Interface) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Interface) MarshalJSON 

``` go
func (v Interface) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Interface) Set 

``` go
func (v *Interface) Set(value interface{}) (old interface{})
```

Set atomically stores `value` into t.value and returns the previous value of t.value. Note: The parameter `value` cannot be nil.

##### (*Interface) String 

``` go
func (v *Interface) String() string
```

String implements String interface for string printing.

##### (*Interface) UnmarshalJSON 

``` go
func (v *Interface) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Interface) UnmarshalValue 

``` go
func (v *Interface) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Interface) Val 

``` go
func (v *Interface) Val() interface{}
```

Val atomically loads and returns t.value.

#### type String 

``` go
type String struct {
	// contains filtered or unexported fields
}
```

String is a struct for concurrent-safe operation for type string.

##### func NewString 

``` go
func NewString(value ...string) *String
```

NewString creates and returns a concurrent-safe object for string type, with given initial value `value`.

##### (*String) Clone 

``` go
func (v *String) Clone() *String
```

Clone clones and returns a new concurrent-safe object for string type.

##### (*String) DeepCopy <-2.1.2

``` go
func (v *String) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (String) MarshalJSON 

``` go
func (v String) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*String) Set 

``` go
func (v *String) Set(value string) (old string)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*String) String 

``` go
func (v *String) String() string
```

String implements String interface for string printing.

##### (*String) UnmarshalJSON 

``` go
func (v *String) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*String) UnmarshalValue 

``` go
func (v *String) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*String) Val 

``` go
func (v *String) Val() string
```

Val atomically loads and returns t.value.

#### type Uint 

``` go
type Uint struct {
	// contains filtered or unexported fields
}
```

Uint is a struct for concurrent-safe operation for type uint.

##### func NewUint 

``` go
func NewUint(value ...uint) *Uint
```

NewUint creates and returns a concurrent-safe object for uint type, with given initial value `value`.

##### (*Uint) Add 

``` go
func (v *Uint) Add(delta uint) (new uint)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Uint) Cas 

``` go
func (v *Uint) Cas(old, new uint) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Uint) Clone 

``` go
func (v *Uint) Clone() *Uint
```

Clone clones and returns a new concurrent-safe object for uint type.

##### (*Uint) DeepCopy <-2.1.0

``` go
func (v *Uint) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Uint) MarshalJSON 

``` go
func (v Uint) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Uint) Set 

``` go
func (v *Uint) Set(value uint) (old uint)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Uint) String 

``` go
func (v *Uint) String() string
```

String implements String interface for string printing.

##### (*Uint) UnmarshalJSON 

``` go
func (v *Uint) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Uint) UnmarshalValue 

``` go
func (v *Uint) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Uint) Val 

``` go
func (v *Uint) Val() uint
```

Val atomically loads and returns t.value.

#### type Uint32 

``` go
type Uint32 struct {
	// contains filtered or unexported fields
}
```

Uint32 is a struct for concurrent-safe operation for type uint32.

##### func NewUint32 

``` go
func NewUint32(value ...uint32) *Uint32
```

NewUint32 creates and returns a concurrent-safe object for uint32 type, with given initial value `value`.

##### (*Uint32) Add 

``` go
func (v *Uint32) Add(delta uint32) (new uint32)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Uint32) Cas 

``` go
func (v *Uint32) Cas(old, new uint32) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Uint32) Clone 

``` go
func (v *Uint32) Clone() *Uint32
```

Clone clones and returns a new concurrent-safe object for uint32 type.

##### (*Uint32) DeepCopy <-2.1.0

``` go
func (v *Uint32) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Uint32) MarshalJSON 

``` go
func (v Uint32) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Uint32) Set 

``` go
func (v *Uint32) Set(value uint32) (old uint32)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Uint32) String 

``` go
func (v *Uint32) String() string
```

String implements String interface for string printing.

##### (*Uint32) UnmarshalJSON 

``` go
func (v *Uint32) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Uint32) UnmarshalValue 

``` go
func (v *Uint32) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Uint32) Val 

``` go
func (v *Uint32) Val() uint32
```

Val atomically loads and returns t.value.

#### type Uint64 

``` go
type Uint64 struct {
	// contains filtered or unexported fields
}
```

Uint64 is a struct for concurrent-safe operation for type uint64.

##### func NewUint64 

``` go
func NewUint64(value ...uint64) *Uint64
```

NewUint64 creates and returns a concurrent-safe object for uint64 type, with given initial value `value`.

##### (*Uint64) Add 

``` go
func (v *Uint64) Add(delta uint64) (new uint64)
```

Add atomically adds `delta` to t.value and returns the new value.

##### (*Uint64) Cas 

``` go
func (v *Uint64) Cas(old, new uint64) (swapped bool)
```

Cas executes the compare-and-swap operation for value.

##### (*Uint64) Clone 

``` go
func (v *Uint64) Clone() *Uint64
```

Clone clones and returns a new concurrent-safe object for uint64 type.

##### (*Uint64) DeepCopy <-2.1.0

``` go
func (v *Uint64) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (Uint64) MarshalJSON 

``` go
func (v Uint64) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*Uint64) Set 

``` go
func (v *Uint64) Set(value uint64) (old uint64)
```

Set atomically stores `value` into t.value and returns the previous value of t.value.

##### (*Uint64) String 

``` go
func (v *Uint64) String() string
```

String implements String interface for string printing.

##### (*Uint64) UnmarshalJSON 

``` go
func (v *Uint64) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### (*Uint64) UnmarshalValue 

``` go
func (v *Uint64) UnmarshalValue(value interface{}) error
```

UnmarshalValue is an interface implement which sets any type of value for `v`.

##### (*Uint64) Val 

``` go
func (v *Uint64) Val() uint64
```

Val atomically loads and returns t.value.