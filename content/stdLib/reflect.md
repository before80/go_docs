+++
title = "reflect"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
[https://pkg.go.dev/reflect@go1.20.1](https://pkg.go.dev/reflect@go1.20.1)

Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types. The typical use is to take a value with static type interface{} and extract its dynamic type information by calling TypeOf, which returns a Type.

​	`reflect`包实现了运行时反射，允许程序操作任意类型的对象。典型用法是将静态类型为`interface{}`的值传递给`TypeOf`函数提取其动态类型信息，`TypeOf`函数返回一个`Type`。

A call to ValueOf returns a Value representing the run-time data. Zero takes a Type and returns a Value representing a zero value for that type.

​	调用`ValueOf`函数返回一个`Value`类型的值，表示运行时数据。`Zero`函数接受一个`Type`参数，并返回表示该类型零值的Value。

See "The Laws of Reflection" for an introduction to reflection in Go: https://golang.org/doc/articles/laws_of_reflection.html

​	请参阅《[反射法则]({{< ref "/goBlog/2011/TheLawsOfReflection">}})》(The Laws of Reflection)了解Go语言中的反射介绍。


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=273)

``` go 
const Ptr = Pointer
```

Ptr is the old name for the Pointer kind.

​	`Ptr`是`Pointer`种类的旧名称。

## 变量

This section is empty.

## 函数

### func Copy 

``` go 
func Copy(dst, src Value) int
```

Copy copies the contents of src into dst until either dst has been filled or src has been exhausted. It returns the number of elements copied. Dst and src each must have kind Slice or Array, and dst and src must have the same element type.

​	`Copy` 函数将 `src` 的内容复制到 `dst`，直到 `dst` 已满或 `src` 已耗尽。它返回已复制的元素数量。`dst` 和 `src` 必须都是 `Slice` 或 `Array` 类型，而且它们的元素类型必须相同。

As a special case, src can have kind String if the element type of dst is kind Uint8.

​	作为特殊情况，如果 `dst` 的元素类型是 `Uint8`，则 `src` （的类型，非元素的类型）可以是 `String` 类型。

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a0 := []int{1, 2, 3}
	va := reflect.ValueOf(a0)

	b := make([]int, 2)
	vb := reflect.ValueOf(b)
	fmt.Println(reflect.Copy(vb, va)) // 2

	c := make([]int, 3)
	vc := reflect.ValueOf(c)
	fmt.Println(reflect.Copy(vc, va)) // 3

	d := make([]int, 4)
	vd := reflect.ValueOf(d)
	fmt.Println(reflect.Copy(vd, va)) // 3

	//a1 := []string{"你好", "中国", "hello", "China"}
	//va1 := reflect.ValueOf(a1)
	//
	//b1 := make([]uint8, 2)
	//vb1 := reflect.ValueOf(b1)
	//fmt.Println(reflect.Copy(vb1, va1)) // panic: reflect.Copy: uint8 != string

	//a2 := []string{"你", "好", "中", "国"}
	//va2 := reflect.ValueOf(a2)
	//
	//b2 := make([]uint8, 2)
	//vb2 := reflect.ValueOf(b2)
	//fmt.Println(reflect.Copy(vb2, va2)) // panic: reflect.Copy: uint8 != string

	a3 := "你好中国"
	va3 := reflect.ValueOf(a3)

	b3 := make([]uint8, 2)
	vb3 := reflect.ValueOf(b3)
	fmt.Println(reflect.Copy(vb3, va3)) // 2

	b4 := make([]uint8, 3)
	vb4 := reflect.ValueOf(b4)
	fmt.Println(reflect.Copy(vb4, va3)) // 3

	b5 := make([]uint8, 12)
	vb5 := reflect.ValueOf(b5)
	fmt.Println(reflect.Copy(vb5, va3)) // 12

	b6 := make([]uint8, 13)
	vb6 := reflect.ValueOf(b6)
	fmt.Println(reflect.Copy(vb6, va3)) // 12
}

// Output:
//2
//3
//3
//2
//3
//12
//12

```



### func DeepEqual 

``` go 
func DeepEqual(x, y any) bool
```

DeepEqual reports whether x and y are “deeply equal,” defined as follows. Two values of identical type are deeply equal if one of the following cases applies. Values of distinct types are never deeply equal.

​	`DeepEqual` 函数报告 `x` 和 `y` 是否“深度相等（deeply equal）”，定义如下。两个具有相同类型的值在以下情况下被认为是深度相等。不同类型的值永远不会深度相等。

Array values are deeply equal when their corresponding elements are deeply equal.

​	`Array`值在它们对应的元素深度相等时是深度相等的。

Struct values are deeply equal if their corresponding fields, both exported and unexported, are deeply equal.

​	`Struct`值在它们对应的字段（包括导出的和未导出的）深度相等时是深度相等的。

Func values are deeply equal if both are nil; otherwise they are not deeply equal.

​	如果两个`Func`值都是`nil`，则它们是深度相等的；否则，它们不是深度相等的。

Interface values are deeply equal if they hold deeply equal concrete values.

​	如果 `Interface` 值持有深度相等的具体值，那么它们是深度相等的。

Map values are deeply equal when all of the following are true: they are both nil or both non-nil, they have the same length, and either they are the same map object or their corresponding keys (matched using Go equality) map to deeply equal values.

​	当以下所有条件都满足时，`Map`值是深度相等的：它们都是`nil`或者都不是`nil`，它们有相同的长度，要么它们是相同的映射对象，要么它们相应的键（使用Go的相等性匹配）映射到深度相等的值。

Pointer values are deeply equal if they are equal using Go's == operator or if they point to deeply equal values.

​	当它们使用 Go 的`==`运算符相等，或者它们指向深度相等的值时，`Pointer` 值是深度相等的。

Slice values are deeply equal when all of the following are true: they are both nil or both non-nil, they have the same length, and either they point to the same initial entry of the same underlying array (that is, &x[0] == &y[0]) or their corresponding elements (up to length) are deeply equal. Note that a non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil)) are not deeply equal.

​	当以下所有条件都满足时，`Slice`值是深度相等的：它们都是`nil`或者都不是`nil`，它们有相同的长度，要么它们指向相同底层数组的相同初始条目（即，`&x[0] == &y[0]`），要么它们相应的元素（达到长度）是深度相等的。请注意，非`nil`的空切片和`nil`切片（例如，`[]byte{}`和`[]byte(nil)`）不是深度相等的。

Other values - numbers, bools, strings, and channels - are deeply equal if they are equal using Go's == operator.

​	其他值，如数字、布尔值、字符串和通道，如果它们使用Go的`==`操作符相等，则是深度相等的。

In general DeepEqual is a recursive relaxation of Go's == operator. However, this idea is impossible to implement without some inconsistency. Specifically, it is possible for a value to be unequal to itself, either because it is of func type (uncomparable in general) or because it is a floating-point NaN value (not equal to itself in floating-point comparison), or because it is an array, struct, or interface containing such a value. On the other hand, pointer values are always equal to themselves, even if they point at or contain such problematic values, because they compare equal using Go's == operator, and that is a sufficient condition to be deeply equal, regardless of content. DeepEqual has been defined so that the same short-cut applies to slices and maps: if x and y are the same slice or the same map, they are deeply equal regardless of content.

​	一般来说，`DeepEqual`是Go的`==`操作符的递归放宽版本。然而，如果不存在一些不一致性，这个想法是无法实现的。具体来说，**一个值可能会与自己不相等，要么是因为它是函数类型（通常无法比较），要么是因为它是浮点`NaN`值（在浮点比较中与自己不相等），要么是因为它是包含这样的值的数组、结构体或接口**。另一方面，**指针值总是与自己相等，即使它们指向或包含这样有问题的值，因为它们使用Go的`==`操作符进行比较是相等的，这是一个足够的条件，使其无论内容如何都被视为深度相等**。`DeepEqual`的定义使得相同的捷径适用于切片和映射：如果`x`和`y`是相同的切片或相同的映射，则无论内容如何，它们都是深度相等的。

As DeepEqual traverses the data values it may find a cycle. The second and subsequent times that DeepEqual compares two pointer values that have been compared before, it treats the values as equal rather than examining the values to which they point. This ensures that DeepEqual terminates.

​	`DeepEqual`在遍历数据值时，可能会发现一个循环。`DeepEqual`比较两个之前已经比较过的指针值时，第二次及以后会将这些值视为相等，而不是检查它们所指向的值。这确保了`DeepEqual`能终止运行。

### func Swapper  <- go1.8

``` go 
func Swapper(slice any) func(i, j int)
```

Swapper returns a function that swaps the elements in the provided slice.

​	`Swapper` 函数返回一个用于交换提供的切片中元素的函数。

Swapper panics if the provided interface is not a slice.

​	如果提供的接口不是切片类型，`Swapper` 函数会引发 panic。

## 类型

### type ChanDir 

``` go 
type ChanDir int
```

ChanDir represents a channel type's direction.	

​	`ChanDir` 类型表示通道类型的方向。

``` go 
const (
	RecvDir ChanDir             = 1 << iota // <-chan
	SendDir                                 // chan<-
	BothDir = RecvDir | SendDir             // chan
)
```

#### (ChanDir) String 

``` go 
func (d ChanDir) String() string
```

​	`String` 方法返回 `ChanDir` 的字符串形式。

### type Kind 

``` go 
type Kind uint
```

A Kind represents the specific kind of type that a Type represents. The zero Kind is not a valid kind.

​	`Kind` 类型表示 `Type` 表示的特定类型。零 `Kind` 不是有效的类型。

### Kind Example

``` go 
package main

import (
	"fmt"
	"reflect"
)

func main() {
	for _, v := range []any{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}

}
Output:

hi
42
unhandled kind func
```



``` go 
const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)
```

#### (Kind) String 

``` go 
func (k Kind) String() string
```

String returns the name of k.

​	`String`方法返回`k`的名称。

### type MapIter  <- go1.12

``` go 
type MapIter struct {
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

A MapIter is an iterator for ranging over a map. See Value.MapRange.

​	`MapIter` 是用于遍历映射的迭代器。参见 Value.MapRange。

#### (*MapIter) Key  <- go1.12

``` go 
func (iter *MapIter) Key() Value
```

Key returns the key of iter's current map entry.

​	`Key`方法返回`iter`当前映射条目的键。

#### (*MapIter) Next  <- go1.12

``` go 
func (iter *MapIter) Next() bool
```

Next advances the map iterator and reports whether there is another entry. It returns false when iter is exhausted; subsequent calls to Key, Value, or Next will panic.

​	`Next`方法推进映射迭代器，并报告是否有另一个条目。当`iter`耗尽时，它返回`false`；对`Key`方法，`Value`方法或`Next`方法的后续调用将引发panic。

#### (*MapIter) Reset  <- go1.18

``` go 
func (iter *MapIter) Reset(v Value)
```

Reset modifies iter to iterate over v. It panics if v's Kind is not Map and v is not the zero Value. Reset(Value{}) causes iter to not to refer to any map, which may allow the previously iterated-over map to be garbage collected.

​	`Reset`方法修改`iter`以遍历`v`。如果`v`的`Kind`不是`Map`且`v`不是零`Value`，则它会引发panic。`Reset(Value{})`会导致`iter`不引用任何映射，这可能允许之前遍历过的映射被垃圾回收。

#### (*MapIter) Value  <- go1.12

``` go 
func (iter *MapIter) Value() Value
```

Value returns the value of iter's current map entry.

​	`Value`方法返回`iter`当前映射条目的值。

### type Method 

``` go 
type Method struct {
	// Name 是方法名。
	Name string

	// PkgPath 是包路径，用于标识一个小写(未导出)的方法名。
    // 对于大写(导出)的方法名，它为空。
	// PkgPath 和 Name 的组合在方法集中唯一标识一个方法。
	// 参见 https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath string


	Type  Type  // 方法类型
	Func  Value // 具有接收器为第一个参数的函数
	Index int   // Type.Method的索引值
}
```

Method represents a single method.

​	`Method`表示单个方法。

#### (Method) IsExported  <- go1.17

``` go 
func (m Method) IsExported() bool
```

IsExported reports whether the method is exported.

​	`IsExported` 方法返回该方法是否为导出方法。

### type SelectCase  <- go1.1

``` go 
type SelectCase struct {
	Dir  SelectDir // case 方向
	Chan Value     // 用于发送或接收的通道
	Send Value     // 发送的值 (用于发送)
}
```

A SelectCase describes a single case in a select operation. The kind of case depends on Dir, the communication direction.

​	`SelectCase` 描述 `select` 操作中的一个单独 `case`。情况的种类取决于 `Dir`，通信方向。

If Dir is SelectDefault, the case represents a default case. Chan and Send must be zero Values.

​	如果 `Dir` 是 `SelectDefault`，则该 `case` 表示默认 case。`Chan` 和 `Send` 字段必须是零 `Value`。

If Dir is SelectSend, the case represents a send operation. Normally Chan's underlying value must be a channel, and Send's underlying value must be assignable to the channel's element type. As a special case, if Chan is a zero Value, then the case is ignored, and the field Send will also be ignored and may be either zero or non-zero.

​	如果 `Dir` 是 `SelectSend`，则该 `case` 表示发送操作。通常情况下，`Chan`字段的底层值必须是通道，并且 `Send` 字段的底层值必须可以赋值给通道的元素类型。作为特殊情况，如果 `Chan`字段是零 `Value`，则 `case` 将被忽略，`Send` 字段也将被忽略，可以是零或非零。

If Dir is SelectRecv, the case represents a receive operation. Normally Chan's underlying value must be a channel and Send must be a zero Value. If Chan is a zero Value, then the case is ignored, but Send must still be a zero Value. When a receive operation is selected, the received Value is returned by Select.

​	如果 `Dir` 是 `SelectRecv`，则该 `case` 表示接收操作。通常情况下，`Chan` 字段的底层值必须是通道，`Send` 字段必须是零 `Value`。如果 `Chan` 是零 `Value`，则 `case` 将被忽略，但 `Send` 仍然必须是零 `Value`。当选择接收操作时，接收到的 `Value` 将由 `Select` 返回。	

### type SelectDir  <- go1.1

``` go 
type SelectDir int
```

A SelectDir describes the communication direction of a select case.

​	`SelectDir` 描述 `select` `case` 的通信方向。

``` go 
const (
	SelectSend    SelectDir // case Chan <- Send
	SelectRecv              // case <-Chan:
	SelectDefault           // default
)
```

### type SliceHeader 

``` go 
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

SliceHeader is the runtime representation of a slice. It cannot be used safely or portably and its representation may change in a later release. Moreover, the Data field is not sufficient to guarantee the data it references will not be garbage collected, so programs must keep a separate, correctly typed pointer to the underlying data.

​	`SliceHeader` 是切片的运行时表示。它不能安全或可移植地使用，其表示可能会在以后的版本中更改。此外，`Data` 字段不足以保证其引用的数据不会被垃圾收集，因此程序必须保留一个单独的、正确类型的指针来引用底层数据。

In new code, use unsafe.Slice or unsafe.SliceData instead.

​	在新代码中，请使用 `unsafe.Slice` 或 `unsafe.SliceData` 代替。	

### type StringHeader 

``` go 
type StringHeader struct {
	Data uintptr
	Len  int
}
```

StringHeader is the runtime representation of a string. It cannot be used safely or portably and its representation may change in a later release. Moreover, the Data field is not sufficient to guarantee the data it references will not be garbage collected, so programs must keep a separate, correctly typed pointer to the underlying data.

​	`StringHeader` 是字符串的运行时表示。它不能安全或可移植地使用，其表示可能会在以后的版本中更改。此外，`Data` 字段不足以保证其引用的数据不会被垃圾收集，因此程序必须保留一个单独的、正确类型的指针来引用底层数据。

In new code, use unsafe.String or unsafe.StringData instead.

​	在新代码中，请使用 `unsafe.String` 或 `unsafe.StringData` 代替。	

### type StructField 

``` go 
type StructField struct {
    // Name is the field name.
	// Name 是字段名。
	Name string

    // PkgPath is the package path that qualifies a lower case (unexported)
	// field name. It is empty for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	// PkgPath 是限定小写(未导出)字段名的包路径。
    // 对于大写(导出)字段名，它为空。
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath string

	Type      Type      // 字段类型 field type
	Tag       StructTag // 字段标签字符串 field tag string
	Offset    uintptr   // 在结构体内的偏移量(以字节为单位) offset within struct, in bytes
	Index     []int     // 用于Type.FieldByIndex的索引序列 index sequence for Type.FieldByIndex
	Anonymous bool      // 是否为嵌入字段 is an embedded field
}
```

A StructField describes a single field in a struct.

​	`StructField` 描述结构体中的一个字段。

#### func VisibleFields  <- go1.17

``` go 
func VisibleFields(t Type) []StructField
```

VisibleFields returns all the visible fields in t, which must be a struct type. A field is defined as visible if it's accessible directly with a FieldByName call. The returned fields include fields inside anonymous struct members and unexported fields. They follow the same order found in the struct, with anonymous fields followed immediately by their promoted fields.

​	`VisibleFields` 函数返回类型 `t` 中的所有可见字段，`t` 必须是结构体类型。字段定义为可直接通过 `FieldByName` 方法调用访问的字段。返回的字段包括嵌入结构体成员内的字段和未导出字段。它们遵循与结构体中找到的相同顺序，嵌入字段后面紧跟其被提升的字段。

For each element e of the returned slice, the corresponding field can be retrieved from a value v of type t by calling v.FieldByIndex(e.Index).

​	对于返回的切片的每个元素 `e`，可以通过调用 `v.FieldByIndex(e.Index)` 从类型 `t` 的值 `v` 中获取相应的字段。	

#### (StructField) IsExported  <- go1.17

``` go 
func (f StructField) IsExported() bool
```

IsExported reports whether the field is exported.

​	`IsExported`方法报告字段是否已导出。

### type StructTag 

``` go 
type StructTag string
```

A StructTag is the tag string in a struct field.

​	`StructTag` 是结构体字段中的标签字符串。

By convention, tag strings are a concatenation of optionally space-separated key:"value" pairs. Each key is a non-empty string consisting of non-control characters other than space (U+0020 ' '), quote (U+0022 '"'), and colon (U+003A ':'). Each value is quoted using U+0022 '"' characters and Go string literal syntax.

​	按照约定，标记字符串是一个可选的空格分隔的`key:"value"`对的串联。每个`key`都是一个由非空格（U+0020 ' '）、引号（U+0022 '`"`'）和冒号（U+003A '`:`'）之外的控制字符组成的非空字符串。每个`value`都使用U+0022 '`"`'字符和Go字符串文字语法引起引用。 

#### StructTag Example

``` go 
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

}
Output:

blue gopher
```



#### (StructTag) Get 

``` go 
func (tag StructTag) Get(key string) string
```

Get returns the value associated with key in the tag string. If there is no such key in the tag, Get returns the empty string. If the tag does not have the conventional format, the value returned by Get is unspecified. To determine whether a tag is explicitly set to the empty string, use Lookup.

​	`Get`方法在标签字符串中返回与`key`关联的值。如果标签中没有这样的键，则 `Get` 返回空字符串。如果标签不具有常规格式，则 `Get` 返回的值是未指定的。要确定标签是否显式设置为空字符串，请使用 `Lookup`方法。

#### (StructTag) Lookup  <- go1.7

``` go 
func (tag StructTag) Lookup(key string) (value string, ok bool)
```

Lookup returns the value associated with key in the tag string. If the key is present in the tag the value (which may be empty) is returned. Otherwise the returned value will be the empty string. The ok return value reports whether the value was explicitly set in the tag string. If the tag does not have the conventional format, the value returned by Lookup is unspecified.

​	`Lookup`方法在标签字符串中返回与`key`关联的值。如果标签中存在该键，则返回值(可能为空)。否则，返回的值将是空字符串。`ok` 返回值报告值是否在标签字符串中显式设置。如果标签不具有常规格式，则 `Lookup` 返回的值是未指定的。

##### Lookup Example

``` go 
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F0 string `alias:"field_0"`
		F1 string `alias:""`
		F2 string
	}

	s := S{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if alias, ok := field.Tag.Lookup("alias"); ok {
			if alias == "" {
				fmt.Println("(blank)")
			} else {
				fmt.Println(alias)
			}
		} else {
			fmt.Println("(not specified)")
		}
	}

}
Output:

field_0
(blank)
(not specified)
```



### type Type 

``` go 
type Type interface {
	// Align returns the alignment in bytes of a value of
	// this type when allocated in memory.
    // Align返回在内存中分配该类型的值的对齐字节数。
	Align() int

    // FieldAlign returns the alignment in bytes of a value of
	// this type when used as a field in a struct.
	// FieldAlign返回在结构体中使用该类型作为字段时的值的对齐字节数。
	FieldAlign() int
	
    // Method returns the i'th method in the type's method set.
    // Method返回类型方法集中的第i个方法。 
	// It panics if i is not in the range [0, NumMethod()).
    // 如果i不在[0, NumMethod())范围内，则会引发恐慌。
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver,
	// and only exported methods are accessible.
    // 对于非接口类型T或*T，返回方法的Type和Func字段描述了第一个参数为接收器的函数，
    // 并且只有导出的方法才可访问。
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
    // 对于接口类型，返回方法的Type字段给出没有接收器的方法签名，并且Func字段为nil。
	//
	// Methods are sorted in lexicographic order.
    // 方法按字母顺序排序。
	Method(int) Method

    // MethodByName returns the method with that name in the type's
	// method set and a boolean indicating if the method was found.
    // MethodByName返回类型方法集中具有该名称的方法以及一个布尔值，指示是否找到该方法。
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	// 对于接口类型，返回方法的Type字段给出没有接收器的方法签名，并且Func字段为nil。
	MethodByName(string) (Method, bool)

	// NumMethod returns the number of methods accessible using Method.
    // NumMethod返回使用Method可访问的方法数。
	//
	// For a non-interface type, it returns the number of exported methods.
    // 对于非接口类型，它返回导出的方法数。
	//
	// For an interface type, it returns the number of exported and unexported methods.
    // 对于接口类型，它返回导出的和未导出的方法数。
	NumMethod() int

    // Name returns the type's name within its package for a defined type.
	// For other (non-defined) types it returns the empty string.
    // Name返回已定义类型的类型的名称（在其包内）。 
    // 对于其他（未定义的）类型，它将返回空字符串。	
	Name() string

    // PkgPath returns a defined type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or not defined (*T, struct{},
	// []int, or A where A is an alias for a non-defined type), the package path
	// will be the empty string.
    // PkgPath() 方法返回一个已定义的类型的包路径，
    // 即唯一标识该包的导入路径，例如 "encoding/base64"。
    // 如果该类型是预先声明的（例如 string 或 error）或未定义的
    // （例如 *T，struct{}，[int]，或 A（其中A是非定义类型的别名）），则包路径将是空字符串。
	PkgPath() string

    // Size returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
	// Size 方法返回存储给定类型值的字节数，类似于 unsafe.Sizeof。
	Size() uintptr

    // String returns a string representation of the type.
	// The string representation may use shortened package names
	// (e.g., base64 instead of "encoding/base64") and is not
	// guaranteed to be unique among types. To test for type identity,
	// compare the Types directly.
	// String() 方法返回类型的字符串表示形式。
    // 该字符串表示形式可能使用简短的文件名（例如 base64 而不是 "encoding/base64"）
    // 并不保证在类型之间具有唯一性。要测试类型的身份，请直接比较 Types。
	String() string

    // Kind returns the specific kind of this type.
	// Kind方法返回此类型的特定种类（Kind）。
	Kind() Kind

    // Implements reports whether the type implements the interface type u.
	// Implements方法报告该类型是否实现了接口类型 u。
	Implements(u Type) bool

    // AssignableTo reports whether a value of the type is assignable to type u.
	// AssignableTo 方法报告一个该类型的值是否可分配给类型 u。
	AssignableTo(u Type) bool

	// ConvertibleTo reports whether a value of the type is convertible to type u.
	// Even if ConvertibleTo returns true, the conversion may still panic.
	// For example, a slice of type []T is convertible to *[N]T,
	// but the conversion will panic if its length is less than N.
    // ConvertibleTo 方法报告一个该类型的值是否可转换为类型 u。
    // 即使 ConvertibleTo() 返回 true，转换仍可能引发恐慌。
    // 例如，类型为 []T 的切片可转换为 *[N]T，但如果其长度小于 N，则转换会引发恐慌。
	ConvertibleTo(u Type) bool

    // Comparable reports whether values of this type are comparable.
	// Even if Comparable returns true, the comparison may still panic.
	// For example, values of interface type are comparable,
	// but the comparison will panic if their dynamic type is not comparable.
    // Comparable 方法报告该类型的值是否可以比较。
    // 即使 Comparable() 返回 true，比较仍可能引发恐慌。
    // 例如，接口类型的值是可以比较的，但如果它们的动态类型不可比较，则比较会引发恐慌。	
	Comparable() bool

    // Bits returns the size of the type in bits.
	// It panics if the type's Kind is not one of the
	// sized or unsized Int, Uint, Float, or Complex kinds.
	// Bits方法返回该类型的大小（以位为单位）。
    // 如果该类型的 Kind 不是有尺寸或无尺寸的 Int、Uint、Float 或 Complex 类型，
    // 则 Bits() 会引发恐慌。
	Bits() int

    // ChanDir returns a channel type's direction.
	// It panics if the type's Kind is not Chan.
	// ChanDir 返回一个通道类型的方向。  
	// 如果该类型的 Kind 不是 Chan，则会引发恐慌。  
	ChanDir() ChanDir

    // IsVariadic reports whether a function type's final input parameter
	// is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
	// implicit actual type []T.
    // IsVariadic 报告函数类型的最后一个输入参数  
	// 是否是 "..." 参数。如果是，t.In(t.NumIn() - 1) 返回参数的  
	// 隐式实际类型 []T。  
	//
	// For concreteness, if t represents func(x int, y ... float64), then
    // 具体来说，如果 t 代表 func(x int, y ... float64)，那么 
	//
	//	t.NumIn() == 2
	//	t.In(0) is the reflect.Type for "int"
	//	t.In(1) is the reflect.Type for "[]float64"
	//	t.IsVariadic() == true
	//
	// IsVariadic panics if the type's Kind is not Func.
	// 如果类型的 Kind 不是 Func，IsVariadic 会引发恐慌。      
	IsVariadic() bool

    // Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
	// Elem 返回一个类型的元素类型。  
	// 如果该类型的 Kind 不是 Array、Chan、Map、Pointer 或 Slice，则会引发恐慌。
	Elem() Type

    // Field returns a struct type's i'th field.
    // Field 返回结构类型的第 i 个字段。
	// It panics if the type's Kind is not Struct.
    // 如果类型的 Kind 不是 Struct，则会引发恐慌。
	// It panics if i is not in the range [0, NumField()).
	// 如果 i 不在范围 [0, NumField()) 内，也会引发恐慌。
	Field(i int) StructField

    // FieldByIndex returns the nested field corresponding
	// to the index sequence. It is equivalent to calling Field
	// successively for each index i.
    // FieldByIndex 返回与索引序列相对应的嵌套字段。  
	// 它相当于为每个索引 i 连续调用 Field。  
	// It panics if the type's Kind is not Struct.
	// 如果类型的 Kind 不是 Struct，则会引发恐慌。
	FieldByIndex(index []int) StructField

    // FieldByName returns the struct field with the given name
	// and a boolean indicating if the field was found.
	// FieldByName 返回给定名称的结构字段  
	// 以及一个布尔值，表示是否找到该字段。  
	FieldByName(name string) (StructField, bool)

    // FieldByNameFunc returns the struct field with a name
	// that satisfies the match function and a boolean indicating if
	// the field was found.
    // FieldByNameFunc 返回满足匹配函数的结构字段  
	// 以及一个布尔值，表示是否找到该字段。  
	//
	// FieldByNameFunc considers the fields in the struct itself
	// and then the fields in any embedded structs, in breadth first order,
	// stopping at the shallowest nesting depth containing one or more
	// fields satisfying the match function. If multiple fields at that depth
	// satisfy the match function, they cancel each other
	// and FieldByNameFunc returns no match.
	// This behavior mirrors Go's handling of name lookup in
	// structs containing embedded fields.
	// FieldByNameFunc 首先考虑结构本身的字段，  
	// 然后考虑任何嵌入结构的字段，按广度优先顺序，  
	// 在包含满足匹配函数的一个或多个字段的最浅嵌套深度停止。  
	// 如果在该深度的多个字段满足匹配函数，它们会相互抵消，  
	// 并且 FieldByNameFunc 返回无匹配项。  
	// 此行为反映了 Go 在包含嵌入字段的结构中处理名称查找的行为。  
	FieldByNameFunc(match func(string) bool) (StructField, bool)

    // In returns the type of a function type's i'th input parameter.
    // In 返回函数类型的第 i 个输入参数的类型。
	// It panics if the type's Kind is not Func.
    // 如果该类型的 Kind 不是 Func，则会引发恐慌。
	// It panics if i is not in the range [0, NumIn()).
	// 如果 i 不在范围 [0, NumIn()) 内，也会引发恐慌。  
	In(i int) Type

    // Key returns a map type's key type.
	// It panics if the type's Kind is not Map.
	// Key 返回映射类型的键类型。
	// 如果该类型的 Kind 不是 Map，则会引发恐慌。 
	Key() Type

    // Len returns an array type's length.
	// It panics if the type's Kind is not Array.
	// Len 返回数组类型的长度。  
	// 如果该类型的 Kind 不是 Array，则会引发恐慌。
	Len() int

    // NumField returns a struct type's field count.
    // NumField 返回结构类型的字段数量。 
	// It panics if the type's Kind is not Struct.
	// 如果该类型的 Kind 不是 Struct，则会引发恐慌。
	NumField() int

    // NumIn returns a function type's input parameter count.
    // NumIn 返回函数类型的输入参数数量。
	// It panics if the type's Kind is not Func.
	// 如果该类型的 Kind 不是 Func，则会引发恐慌。
	NumIn() int

    // NumOut returns a function type's output parameter count.
    // NumOut 返回函数类型的输出参数数量。
	// It panics if the type's Kind is not Func.
	// NumOut返回函数类型的输出参数计数。
	// 如果该类型的 Kind 不是 Func，则会引发恐慌。
	NumOut() int

    // Out returns the type of a function type's i'th output parameter.
    // Out 返回函数类型的第 i 个输出参数的类型。
	// It panics if the type's Kind is not Func.
    // 如果该类型的 Kind 不是 Func，则会引发恐慌。
	// It panics if i is not in the range [0, NumOut()).
	// 如果 i 不在范围 [0, NumOut()) 内，也会引发恐慌。
	Out(i int) Type
	// 包含已过滤或未导出的方法
}
```

Type is the representation of a Go type.

​	`Type` 是 Go 语言类型的表示。

Not all methods apply to all kinds of types. Restrictions, if any, are noted in the documentation for each method. Use the Kind method to find out the kind of type before calling kind-specific methods. Calling a method inappropriate to the kind of type causes a run-time panic.

​	并非所有的方法都适用于所有类型。如果有任何限制，会在每个方法的文档中注明。在调用特定类型的方法之前，请使用 `Kind` 方法来查明类型的种类。如果调用的方法是不适合类型的种类，会导致运行时恐慌（panic）。

Type values are comparable, such as with the == operator, so they can be used as map keys. Two Type values are equal if they represent identical types.

​	`Type` 值是可比较的，例如使用 `==` 运算符，因此它们可以用作映射键。如果两个 `Type` 值表示相同的类型，则它们是相等的。

#### func ArrayOf  <- go1.5

``` go 
func ArrayOf(length int, elem Type) Type
```

ArrayOf returns the array type with the given length and element type. For example, if t represents int, ArrayOf(5, t) represents [5]int.

​	`ArrayOf` 函数返回给定长度和元素类型的数组类型。例如，如果 `t` 代表 `int`，则 `ArrayOf(5, t)` 代表 `[5]int`。

If the resulting type would be larger than the available address space, ArrayOf panics.

​	如果结果类型大于可用的地址空间，`ArrayOf` 会引发恐慌。

#### func ChanOf  <- go1.1

``` go 
func ChanOf(dir ChanDir, t Type) Type
```

ChanOf returns the channel type with the given direction and element type. For example, if t represents int, ChanOf(RecvDir, t) represents <-chan int.

​	`ChanOf` 函数返回给定方向和元素类型的通道类型。例如，如果 `t` 代表 `int`，则 `ChanOf(RecvDir, t)` 代表 `<-chan int`。

The gc runtime imposes a limit of 64 kB on channel element types. If t's size is equal to or exceeds this limit, ChanOf panics.

​	gc 运行时对通道元素类型的大小限制为 64 kB。如果 `t` 的大小等于或超过此限制，`ChanOf` 会引发恐慌。

#### func FuncOf  <- go1.5

``` go 
func FuncOf(in, out []Type, variadic bool) Type
```

FuncOf returns the function type with the given argument and result types. For example if k represents int and e represents string, FuncOf([]Type{k}, []Type{e}, false) represents func(int) string.

​	`FuncOf` 函数返回给定实参和结果类型的函数类型。例如，如果 `k` 代表 `int`，`e` 代表 `string`，则 `FuncOf([]Type{k}, []Type{e}, false)` 代表 `func(int) string`。

The variadic argument controls whether the function is variadic. FuncOf panics if the in[len(in)-1] does not represent a slice and variadic is true.

​	变参实参控制函数是否为变参函数。如果 `in[len(in)-1]` 不表示一个切片且 `variadic` 为 `true`，`FuncOf` 会引发恐慌。

#### func MapOf  <- go1.1

``` go 
func MapOf(key, elem Type) Type
```

MapOf returns the map type with the given key and element types. For example, if k represents int and e represents string, MapOf(k, e) represents map[int]string.

​	`MapOf` 函数返回给定键和元素类型的映射类型。例如，如果 `k` 代表 `int`，`e` 代表 `string`，则 `MapOf(k, e)` 代表 `map[int]string`。

If the key type is not a valid map key type (that is, if it does not implement Go's == operator), MapOf panics.

​	如果键类型不是有效的映射键类型（即，如果它没有实现 Go 的 `==` 操作符），`MapOf` 会引发恐慌。

#### func PointerTo  <- go1.18

``` go 
func PointerTo(t Type) Type
```

PointerTo returns the pointer type with element t. For example, if t represents type Foo, PointerTo(t) represents *Foo.

​	`PointerTo` 函数返回元素为 `t` 的指针类型。例如，如果 `t` 代表 `Foo` 类型，`PointerTo(t)` 就代表 `*Foo`。 

#### func PtrTo 

``` go 
func PtrTo(t Type) Type
```

PtrTo returns the pointer type with element t. For example, if t represents type Foo, PtrTo(t) represents *Foo.

​	`PtrTo` 函数返回元素为 `t` 的指针类型。例如，如果 `t` 代表 `Foo` 类型，`PtrTo(t)` 就代表 `*Foo`。

PtrTo is the old spelling of PointerTo. The two functions behave identically.

​	`PtrTo` 是 `PointerTo` 的旧拼写方式。这两个函数的行为完全相同。  	

#### func SliceOf  <- go1.1

``` go 
func SliceOf(t Type) Type
```

SliceOf returns the slice type with element type t. For example, if t represents int, SliceOf(t) represents []int.

​	`SliceOf` 函数返回具有元素类型`t`的切片类型。例如，如果`t`表示`int`，则`SliceOf(t)`表示`[]int`。

#### func StructOf  <- go1.7

``` go 
func StructOf(fields []StructField) Type
```

StructOf returns the struct type containing fields. The Offset and Index fields are ignored and computed as they would be by the compiler.

​	`StructOf` 函数返回包含 `fields` 的结构体类型。`Offset` 和 `Index` 字段会被忽略，并会计算出编译器所期望的值。  

StructOf currently does not generate wrapper methods for embedded fields and panics if passed unexported StructFields. These limitations may be lifted in a future version.

​	`StructOf` 目前不会为嵌入字段生成包装方法，如果传入未导出的 `StructFields`，它会引发恐慌。这些限制在未来的版本中可能会被解除。  

##### StructOf Example

``` go 
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json:  %s", w.Bytes())

	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

}
Output:

value: &{Height:0.4 Age:2}
json:  {"height":0.4,"age":2}
value: &{Height:1.5 Age:10}
```



#### func TypeOf 

``` go 
func TypeOf(i any) Type
```

TypeOf returns the reflection Type that represents the dynamic type of i. If i is a nil interface value, TypeOf returns nil.

​	`TypeOf` 函数返回表示 `i` 的动态类型的反射 `Type`。如果 `i` 是 `nil` 接口值，`TypeOf` 返回 `nil`。  

#### typeOf Example

``` go 
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
    // As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value.
	// 由于接口类型仅用于静态类型检查，
    // 因此找到一个接口类型 Foo 的反射类型的常见做法是使用一个指向 Foo 类型的指针。
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))

}
// Output:
// true
```



### type Value 

``` go 
type Value struct {
	//包含已过滤或未导出的字段
}
```

Value is the reflection interface to a Go value.

​	`Value` 是Go值的反射接口。

Not all methods apply to all kinds of values. Restrictions, if any, are noted in the documentation for each method. Use the Kind method to find out the kind of value before calling kind-specific methods. Calling a method inappropriate to the kind of type causes a run time panic.

​	并非所有方法都适用于所有类型的值。如果有任何限制，会在每个方法的文档中注明。在调用特定类型的方法之前，使用Kind方法来查明值的类型。调用不适合该类型的方法会导致运行时恐慌。  

The zero Value represents no value. Its IsValid method returns false, its Kind method returns Invalid, its String method returns "<invalid Value>", and all other methods panic. Most functions and methods never return an invalid value. If one does, its documentation states the conditions explicitly.

​	零值Value代表没有值。其`IsValid`方法返回`false`，其`Kind`方法返回`Invalid`，其`String`方法返回"`<invalid Value>`"，所有其他方法都会引发恐慌。大多数函数和方法永远不会返回无效的值。如果有，其文档会明确地说明这些条件。  

A Value can be used concurrently by multiple goroutines provided that the underlying Go value can be used concurrently for the equivalent direct operations.

​	如果基础的Go值可以被并发地用于等价的直接操作，那么一个`Value`可以被多个goroutine并发地使用。  

To compare two Values, compare the results of the Interface method. Using == on two Values does not compare the underlying values they represent.

​	要比较两个`Value`，请比较它们的`Interface`方法的结果。在两个`Value`上使用`==`并不会比较它们所代表的基础值。  	

#### func Append 

``` go 
func Append(s Value, x ...Value) Value
```

Append appends the values x to a slice s and returns the resulting slice. As in Go, each x's value must be assignable to the slice's element type.

​	`Append`函数将值`x`追加到切片`s`上，并返回结果切片。像在Go中一样，每个`x`的值必须可以分配给切片的元素类型。  

#### func AppendSlice 

``` go 
func AppendSlice(s, t Value) Value
```

AppendSlice appends a slice t to a slice s and returns the resulting slice. The slices s and t must have the same element type.

​	`AppendSlice`函数将切片`t`追加到切片`s`上，并返回结果切片。切片`s`和`t`必须具有相同的元素类型。  

#### func Indirect 

``` go 
func Indirect(v Value) Value
```

Indirect returns the value that v points to. If v is a nil pointer, Indirect returns a zero Value. If v is not a pointer, Indirect returns v.

​	`Indirect`函数返回`v`指向的值。如果`v`是一个`nil`指针，`Indirect`返回一个零值。如果`v`不是一个指针，`Indirect`返回`v`。  

#### func MakeChan 

``` go 
func MakeChan(typ Type, buffer int) Value
```

MakeChan creates a new channel with the specified type and buffer size.

​	`MakeChan`函数使用指定的类型和缓冲区大小创建一个新的通道。

#### func MakeFunc  <- go1.1

``` go 
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
```

MakeFunc returns a new function of the given Type that wraps the function fn. When called, that new function does the following:

​	`MakeFunc`返回一个给定`Type`的新函数，该函数包装了`fn`。当被调用时，这个新函数执行以下操作：

- converts its arguments to a slice of Values.
- 将其实参转换为 `Value` 切片。
- runs results := fn(args).
- 运行 `results := fn(args)`。
- returns the results as a slice of Values, one per formal result.
- 将结果作为 `Value` 切片返回，每个形式结果一个。

The implementation fn can assume that the argument Value slice has the number and type of arguments given by typ. If typ describes a variadic function, the final Value is itself a slice representing the variadic arguments, as in the body of a variadic function. The result Value slice returned by fn must have the number and type of results given by typ.

​	 `fn` 实现可以假定实参`Value` 切片具有由 `typ` 给出的实参的数量和类型。如果 `typ` 描述了一个可变参数函数，那么最后一个 `Value` 本身就是一个代表可变实参的切片，就像在可变参数函数的主体中一样。由 `fn` 返回的结果 `Value` 切片必须具有由 `typ` 给出的结果的数量和类型。

The Value.Call method allows the caller to invoke a typed function in terms of Values; in contrast, MakeFunc allows the caller to implement a typed function in terms of Values.

​	`Value.Call` 方法允许调用者根据 `Value` 来调用类型化的函数；与此相反，`MakeFunc` 允许调用者根据 `Value` 来实现类型化的函数。

The Examples section of the documentation includes an illustration of how to use MakeFunc to build a swap function for different types.

​	文档中的Examples 部分包括一个关于如何使用 `MakeFunc` 为不同类型构建交换函数的说明。	

##### MakeFunc Example

``` go 
package main

import (
	"fmt"
	"reflect"
)

func main() {
    // swap is the implementation passed to MakeFunc.
	// It must work in terms of reflect.Values so that it is possible
	// to write code without knowing beforehand what the types
	// will be.
    // swap是传递给MakeFunc的实现。
	// 它必须以reflect.Values的方式工作，以便可以编写代码，而不需要预先知道类型。
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

    // makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
    // makeSwap期望fptr是指向nil函数的指针。
	// 它将该指针设置为使用MakeFunc创建的新函数。
	// 当调用该函数时，reflect将参数转换为Values，调用swap，然后将swap的结果切片转换为新函数返回的值。
	makeSwap := func(fptr any) {
        // fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
        // fptr是指向函数的指针。
		// 获取函数值本身（可能为nil）作为reflect.Value，以便我们可以查询其类型，然后设置该值。
		fn := reflect.ValueOf(fptr).Elem()
		
        // Make a function of the right type.
        // 制作正确类型的函数。
		v := reflect.MakeFunc(fn.Type(), swap)

		// Assign it to the value fn represents.
        // 将其分配给fn所代表的值。
		fn.Set(v)
	}

    // Make and call a swap function for ints.
    // 制作并调用一个用于整数的交换函数。
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

    // Make and call a swap function for float64s.
    // 制作并调用一个用于float64的交换函数。
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))

}
Output:

1 0
3.14 2.72
```



#### func MakeMap 

``` go 
func MakeMap(typ Type) Value
```

MakeMap creates a new map with the specified type.

​	`MakeMap` 函数使用指定的类型创建一个新的映射。

#### func MakeMapWithSize  <- go1.9

``` go 
func MakeMapWithSize(typ Type, n int) Value
```

MakeMapWithSize creates a new map with the specified type and initial space for approximately n elements.

​	`MakeMapWithSize`函数使用一个指定的类型和大约`n` 个元素的初始空间创建一个新的映射。

#### func MakeSlice 

``` go 
func MakeSlice(typ Type, len, cap int) Value
```

MakeSlice creates a new zero-initialized slice value for the specified slice type, length, and capacity.

​	`MakeSlice` 函数为指定的切片类型、长度和容量创建一个新的零值初始化切片值。

#### func New 

``` go 
func New(typ Type) Value
```

New returns a Value representing a pointer to a new zero value for the specified type. That is, the returned Value's Type is PointerTo(typ).

​	`New` 函数返回一个 `Value`，它代表指定类型的新零值的指针。也就是说，返回的 `Value` 的 `Type` 是 `PointerTo(typ)`。

#### func NewAt 

``` go 
func NewAt(typ Type, p unsafe.Pointer) Value
```

NewAt returns a Value representing a pointer to a value of the specified type, using p as that pointer.

​	`NewAt` 函数返回一个 `Value`，它表示指向指定类型值的指针，使用 `p` 作为该指针。

#### func Select  <- go1.1

``` go 
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
```

Select executes a select operation described by the list of cases. Like the Go select statement, it blocks until at least one of the cases can proceed, makes a uniform pseudo-random choice, and then executes that case. It returns the index of the chosen case and, if that case was a receive operation, the value received and a boolean indicating whether the value corresponds to a send on the channel (as opposed to a zero value received because the channel is closed). Select supports a maximum of 65536 cases.

​	`Select` 函数执行由 `cases` 列表描述的 `select` 操作。与 Go 的 `select` 语句类似，它会阻塞，直到至少有一个 `case` 可以进行，然后做出统一伪随机的（uniform pseudo-random）选择，并执行该 `case`。它返回所选 `case` 的索引，如果该 `case` 是接收操作，则返回接收的值和一个布尔值，该布尔值表示该值是否与通道上的发送相对应（而不是因为通道关闭而接收到的零值）。`Select` 支持最多 `65536` 个 cases。

#### func ValueOf 

``` go 
func ValueOf(i any) Value
```

ValueOf returns a new Value initialized to the concrete value stored in the interface i. ValueOf(nil) returns the zero Value.

​	`ValueOf` 函数返回一个新的 `Value`，初始化为接口 `i` 中存储的具体值。`ValueOf(nil)` 返回零值 `Value`。

#### func Zero 

``` go 
func Zero(typ Type) Value
```

Zero returns a Value representing the zero value for the specified type. The result is different from the zero value of the Value struct, which represents no value at all. For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0. The returned value is neither addressable nor settable.

​	`Zero` 函数返回一个表示指定类型的零值的 `Value`。该结果与`Value`结构体的零值不同，后者表示根本没有值。例如，`Zero(TypeOf(42))` 返回一个 `Kind` 为 `Int`，值为 `0` 的 `Value`。返回的值既不可寻址也不可设置。

#### (Value) Addr 

``` go 
func (v Value) Addr() Value
```

Addr returns a pointer value representing the address of v. It panics if CanAddr() returns false. Addr is typically used to obtain a pointer to a struct field or slice element in order to call a method that requires a pointer receiver.

​	`Addr`返回表示`v`的地址的指针值。如果`CanAddr()`返回`false`，它将引发`panic`。`Addr`通常用于获取结构体字段或切片元素的指针，以便调用需要指针接收者的方法。

#### (Value) Bool 

``` go 
func (v Value) Bool() bool
```

Bool returns v's underlying value. It panics if v's kind is not Bool.

​	`Bool`方法返回`v`的基础值。如果`v`的kind不是`Bool`，它会引发panic。

#### (Value) Bytes 

``` go 
func (v Value) Bytes() []byte
```

Bytes returns v's underlying value. It panics if v's underlying value is not a slice of bytes or an addressable array of bytes.

​	`Bytes`方法返回`v`的底层值。如果`v`的底层值既不是字节的切片，也不是可寻址的字节数组，它会引发panic。

#### (Value) Call 

``` go 
func (v Value) Call(in []Value) []Value
```

Call calls the function v with the input arguments in. For example, if len(in) == 3, v.Call(in) represents the Go call v(in[0], in[1], in[2]). Call panics if v's Kind is not Func. It returns the output results as Values. As in Go, each input argument must be assignable to the type of the function's corresponding input parameter. If v is a variadic function, Call creates the variadic slice parameter itself, copying in the corresponding values.

​	`Call`方法使用输入实参`in`调用函数`v`。例如，如果`len(in) == 3`，`v.Call(in)`表示Go调用`v(in[0], in[1], in[2])`。如果`v`的`Kind`不是`Func`，`Call`会引发panic。它以`Value`的形式返回输出结果。像在Go中一样，每个输入实参必须可以赋值给函数相应输入参数的类型。如果`v`是可变参数函数，`Call`会自己创建可变参数切片，并复制相应的值。

#### (Value) CallSlice 

``` go 
func (v Value) CallSlice(in []Value) []Value
```

CallSlice calls the variadic function v with the input arguments in, assigning the slice in[len(in)-1] to v's final variadic argument. For example, if len(in) == 3, v.CallSlice(in) represents the Go call v(in[0], in[1], in[2]...). CallSlice panics if v's Kind is not Func or if v is not variadic. It returns the output results as Values. As in Go, each input argument must be assignable to the type of the function's corresponding input parameter.

​	`CallSlice`使用输入实参`in`调用可变参数函数`v`，将切片`in[len(in)-1]`赋值给`v`的最后一个可变参数。例如，如果`len(in) == 3`，`v.CallSlice(in)`表示Go调用`v(in[0], in[1], in[2]...)`。如果`v`的`Kind`不是`Func`，或者`v`不是可变参数，`CallSlice`会引发panic。它以`Value`的形式返回输出结果。像在Go中一样，每个输入实参必须可以赋值给函数相应输入参数的类型。

#### (Value) CanAddr 

``` go 
func (v Value) CanAddr() bool
```

CanAddr reports whether the value's address can be obtained with Addr. Such values are called addressable. A value is addressable if it is an element of a slice, an element of an addressable array, a field of an addressable struct, or the result of dereferencing a pointer. If CanAddr returns false, calling Addr will panic.

​	`CanAddr` 方法报告是否可以使用 `Addr` 获取值的地址。这样的值被称为可寻址的。如果一个值是切片的一个元素、可寻址数组的一个元素、可寻址结构体的一个字段，或者是指针解引用的结果，那么它就是可寻址的。如果 `CanAddr` 返回 `false`，调用 `Addr` 方法将会引发 panic。  

#### (Value) CanComplex  <- go1.18

``` go 
func (v Value) CanComplex() bool
```

CanComplex reports whether Complex can be used without panicking.

​	`CanComplex` 方法报告是否可以使用 `Complex`方法而不引发 panic。  

#### (Value) CanConvert  <- go1.17

``` go 
func (v Value) CanConvert(t Type) bool
```

CanConvert reports whether the value v can be converted to type t. If v.CanConvert(t) returns true then v.Convert(t) will not panic.

​	`CanConvert` 方法报告值 `v` 是否可以被转换为类型 `t`。如果 `v.CanConvert(t)` 返回 `true`，那么 `v.Convert(t)` 将不会引发 panic。  

#### (Value) CanFloat  <- go1.18

``` go 
func (v Value) CanFloat() bool
```

CanFloat reports whether Float can be used without panicking.

​	`CanFloat` 方法报告是否可以使用 `Float`方法而不引发 panic。  

#### (Value) CanInt  <- go1.18

``` go 
func (v Value) CanInt() bool
```

CanInt reports whether Int can be used without panicking.

​	`CanInt` 方法报告是否可以使用 `Int`方法而不引发 panic。  

#### (Value) CanInterface 

``` go 
func (v Value) CanInterface() bool
```

CanInterface reports whether Interface can be used without panicking.

​	`CanInterface` 方法报告是否可以使用 `Interface` 方法而不引发 panic。  

#### (Value) CanSet 

``` go 
func (v Value) CanSet() bool
```

CanSet reports whether the value of v can be changed. A Value can be changed only if it is addressable and was not obtained by the use of unexported struct fields. If CanSet returns false, calling Set or any type-specific setter (e.g., SetBool, SetInt) will panic.

​	`CanSet` 方法报告是否可以更改 `v` 的值。只有当它是可寻址的，并且不是通过未导出的结构体字段获得时，才能更改 `Value`。如果 `CanSet` 返回 `false`，调用 `Set`方法或任何特定类型的设置器（例如，`SetBool`方法、`SetInt`方法）将会引发 panic。  

#### (Value) CanUint  <- go1.18

``` go 
func (v Value) CanUint() bool
```

CanUint reports whether Uint can be used without panicking.

​	`CanUint` 方法报告是否可以使用 `Uint` 方法而不引发 panic。

#### (Value) Cap 

``` go 
func (v Value) Cap() int
```

Cap returns v's capacity. It panics if v's Kind is not Array, Chan, Slice or pointer to Array.

​	 `Cap` 方法返回 `v` 的容量。如果 `v` 的 `Kind` 不是 `Array`、`Chan`、`Slice` 或 `Array` 的指针，它会引发 panic。  

#### (Value) Close 

``` go 
func (v Value) Close()
```

Close closes the channel v. It panics if v's Kind is not Chan.

​	`Close` 方法关闭通道 `v`。如果 `v` 的 `Kind` 不是 `Chan`，它会引发 panic。  

#### (Value) Comparable  <- go1.20

``` go 
func (v Value) Comparable() bool
```

Comparable reports whether the value v is comparable. If the type of v is an interface, this checks the dynamic type. If this reports true then v.Interface() == x will not panic for any x, nor will v.Equal(u) for any Value u.

​	`Comparable` 方法报告值 `v` 是否可比较。如果 `v` 的类型是接口，则会检查动态类型。如果此报告为 `true`，则 `v.Interface() == x` 不会对任何 `x` 引发 panic，对于任何 `Value u`，`v.Equal(u)` 也不会引发 panic。 

#### (Value) Complex 

``` go 
func (v Value) Complex() complex128
```

Complex returns v's underlying value, as a complex128. It panics if v's Kind is not Complex64 or Complex128.

​	`Complex` 方法返回 `v` 的底层值，作为 `complex128`。如果 `v` 的 `Kind` 不是 `Complex64` 或 `Complex128`，它会引发 panic。  

#### (Value) Convert  <- go1.1

``` go 
func (v Value) Convert(t Type) Value
```

Convert returns the value v converted to type t. If the usual Go conversion rules do not allow conversion of the value v to type t, or if converting v to type t panics, Convert panics.

​	`Convert` 方法返回将值 `v` 转换为类型 `t` 的结果。如果通常的 Go 转换规则不允许将值 `v` 转换为类型 `t`，或者如果将 `v` 转换为类型 `t` 引发 panic，`Convert` 就会引发 panic。  

#### (Value) Elem 

``` go 
func (v Value) Elem() Value
```

Elem returns the value that the interface v contains or that the pointer v points to. It panics if v's Kind is not Interface or Pointer. It returns the zero Value if v is nil.

​	`Elem` 方法返回接口 `v` 包含的值或指针 `v` 指向的值。如果 `v` 的 `Kind` 不是 `Interface` 或 `Pointer`，它会引发 panic。如果 `v` 是 `nil`，它会返回零值 `Value`。  

#### (Value) Equal  <- go1.20

``` go 
func (v Value) Equal(u Value) bool
```

Equal reports true if v is equal to u. For two invalid values, Equal will report true. For an interface value, Equal will compare the value within the interface. Otherwise, If the values have different types, Equal will report false. Otherwise, for arrays and structs Equal will compare each element in order, and report false if it finds non-equal elements. During all comparisons, if values of the same type are compared, and the type is not comparable, Equal will panic.

​	`Equal` 方法报告 `v` 是否等于 `u`。对于两个无效值，`Equal` 将报告 `true`。对于接口值，`Equal` 将比较接口中的值。否则，如果值的类型不同，`Equal` 将报告 `false`。否则，对于数组和结构体，`Equal` 将按顺序比较每个元素，并在找到不相等的元素时报告 `false`。在所有比较中，如果比较相同类型的值，并且类型不可比较，`Equal` 就会引发 panic。  

#### (Value) Field 

``` go 
func (v Value) Field(i int) Value
```

Field returns the i'th field of the struct v. It panics if v's Kind is not Struct or i is out of range.

​	`Field` 方法返回结构体 `v` 的第 `i` 个字段。如果 `v` 的 `Kind` 不是 `Struct` 或 `i` 越界，它会引发 panic。  

#### (Value) FieldByIndex 

``` go 
func (v Value) FieldByIndex(index []int) Value
```

FieldByIndex returns the nested field corresponding to index. It panics if evaluation requires stepping through a nil pointer or a field that is not a struct.

​	`FieldByIndex` 方法返回与`index`相对应的嵌套字段。如果求值需要通过nil指针或不是结构体的字段，则会引发panic。  

##### FieldByIndex Example

``` go 
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 本示例展示了一个情况，其中一个提升字段的名称被另一个字段隐藏：
	// FieldByName 将无法工作，因此必须改用 FieldByIndex。
	type user struct {
		firstName string
		lastName  string
	}

	type data struct {
		user
		firstName string
		lastName  string
	}

	u := data{
		user:      user{"Embedded John", "Embedded Doe"},
		firstName: "John",
		lastName:  "Doe",
	}

	s := reflect.ValueOf(u).FieldByIndex([]int{0, 1})
	fmt.Println("embedded last name:", s)

}
Output:

embedded last name: Embedded Doe
```



#### (Value) FieldByIndexErr  <- go1.18

``` go 
func (v Value) FieldByIndexErr(index []int) (Value, error)
```

FieldByIndexErr returns the nested field corresponding to index. It returns an error if evaluation requires stepping through a nil pointer, but panics if it must step through a field that is not a struct.

​	`FieldByIndexErr` 方法返回与`index`相对应的嵌套字段，并返回一个错误。如果求值需要通过`nil`指针，则返回错误；如果必须通过不是结构体的字段，则会引发panic。  

#### (Value) FieldByName 

``` go 
func (v Value) FieldByName(name string) Value
```

FieldByName returns the struct field with the given name. It returns the zero Value if no field was found. It panics if v's Kind is not struct.

​	`FieldByName` 方法返回给定名称的结构体字段。如果没有找到字段，则返回零值。如果`v`的`Kind`不是`Struct`，则会引发panic。  

##### FieldByName Example

``` go 
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type user struct {
		firstName string
		lastName  string
	}
	u := user{firstName: "John", lastName: "Doe"}
	s := reflect.ValueOf(u)

	fmt.Println("Name:", s.FieldByName("firstName"))
}
Output:

Name: John
```

#### (Value) FieldByNameFunc 

``` go 
func (v Value) FieldByNameFunc(match func(string) bool) Value
```

FieldByNameFunc returns the struct field with a name that satisfies the match function. It panics if v's Kind is not struct. It returns the zero Value if no field was found.

​	`FieldByNameFunc` 方法返回满足`match`函数的名称的结构体字段。如果`v`的`Kind`不是`Struct`，则会引发panic。如果没有找到字段，则返回零值。  

#### (Value) Float 

``` go 
func (v Value) Float() float64
```

Float returns v's underlying value, as a float64. It panics if v's Kind is not Float32 or Float64.

​	`Float` 方法将`v`的底层值作为`float64`返回。如果`v`的`Kind`不是`Float32`或`Float64`，则会引发panic。  

#### (Value) Grow  <- go1.20

``` go 
func (v Value) Grow(n int)
```

Grow increases the slice's capacity, if necessary, to guarantee space for another n elements. After Grow(n), at least n elements can be appended to the slice without another allocation.

​	`Grow` 方法增加切片的容量（如果需要），以保证有空间容纳另外`n`个元素。在`Grow(n)`之后，至少可以追加`n`个元素到切片，而不需要再次分配内存。  

It panics if v's Kind is not a Slice or if n is negative or too large to allocate the memory.

​	如果`v`的`Kind`不是`Slice`，或者`n`是负数，或者太大以至于无法分配内存，则会引发panic。

#### (Value) Index 

``` go 
func (v Value) Index(i int) Value
```

Index returns v's i'th element. It panics if v's Kind is not Array, Slice, or String or i is out of range.

​	`Index`方法返回`v`的第`i`个元素。如果`v`的`Kind`不是`Array`，`Slice`或`String`，或者`i`越界，则会引发panic。

#### (Value) Int 

``` go 
func (v Value) Int() int64
```

Int returns v's underlying value, as an int64. It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.

​	`Int`方法将`v`的底层值作为`int64`返回。如果`v`的`Kind`不是`Int`，`Int8`，`Int16`，`Int32`或`Int64`，则会引发panic。

#### (Value) Interface 

``` go 
func (v Value) Interface() (i any)
```

Interface returns v's current value as an interface{}. It is equivalent to:

​	`Interface`方法将`v`的当前值作为`interface{}`返回。它等同于：

``` go 
var i interface{} = (v's underlying value)
```

It panics if the Value was obtained by accessing unexported struct fields.

​	如果`Value`是通过访问未导出的结构体字段获得的，则会引发panic。

#### (Value) InterfaceData <- DEPRECATED

```go
func (v Value) InterfaceData() [2]uintptr
```

InterfaceData returns a pair of unspecified uintptr values. It panics if v's Kind is not Interface.

​	`InterfaceData`方法返回一个未指定的uintptr值对。如果`v`的`Kind`不是`Interface`，就会引发panic。

In earlier versions of Go, this function returned the interface's value as a uintptr pair. As of Go 1.4, the implementation of interface values precludes any defined use of InterfaceData.

​	在Go的早期版本中，这个函数将接口的值作为uintptr对返回。然而，自`Go 1.4`起，接口值的实现不再支持任何定义的`InterfaceData`用法。

Deprecated: The memory representation of interface values is not compatible with InterfaceData.

​	已弃用：接口值的内存表示与`InterfaceData`不兼容。

#### (Value) IsNil 

``` go 
func (v Value) IsNil() bool
```

IsNil reports whether its argument v is nil. The argument must be a chan, func, interface, map, pointer, or slice value; if it is not, IsNil panics. Note that IsNil is not always equivalent to a regular comparison with nil in Go. For example, if v was created by calling ValueOf with an uninitialized interface variable i, i==nil will be true but v.IsNil will panic as v will be the zero Value.

​	`IsNil`方法报告其实参`v`是否为`nil`。实参必须是`chan`、`func`、`interface`、`map`、`pointer`或`slice`值；如果不是，`IsNil`就会引发panic。请注意，`IsNil`并不总是等同于Go中使用`nil`的常规比较。例如，如果`v`是通过使用未初始化的接口变量`i`调用`ValueOf`函数创建的，那么`i==nil`将为`true`，但`v.IsNil`会引发panic，因为`v`将是零值`Value`。（参见以下个人给出的示例）

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var c chan int
	vc := reflect.ValueOf(c)
	fmt.Println(vc.IsNil()) // true

	var f func() error
	vf := reflect.ValueOf(f)
	fmt.Println(vf.IsNil()) // true

	type I interface {
		String()
	}

	//var itf I
	//vi1 := reflect.ValueOf(itf)
	//fmt.Println(vi1.IsNil()) // panic: reflect: call of reflect.Value.IsNil on zero Value

	//var interfaceVar interface{} // 未初始化的接口变量，为 nil
	//vi2 := reflect.ValueOf(interfaceVar)
	//fmt.Println(vi2.IsNil()) // panic: reflect: call of reflect.Value.IsNil on zero Value

	var m map[string]int
	vm := reflect.ValueOf(m)
	fmt.Println(vm.IsNil()) // true

	var ptr *int // 未初始化的指针，指向 nil
	vp := reflect.ValueOf(ptr)
	fmt.Println(vp.IsNil()) // true

	var s []int
	vs := reflect.ValueOf(s)
	fmt.Println(vs.IsNil()) // true

	//var i int
	//v3 := reflect.ValueOf(i)
	//fmt.Println(v3.IsNil()) // panic: reflect: call of reflect.Value.IsNil on int Value

	//var b bool
	//vb := reflect.ValueOf(b)
	//fmt.Println(vb.IsNil()) // panic: reflect: call of reflect.Value.IsNil on bool Value
}

// Output:
//true
//true
//true
//true
//true
```



#### (Value) IsValid 

``` go 
func (v Value) IsValid() bool
```

IsValid reports whether v represents a value. It returns false if v is the zero Value. If IsValid returns false, all other methods except String panic. Most functions and methods never return an invalid Value. If one does, its documentation states the conditions explicitly.

​	`IsValid`报告`v`是否代表一个值。如果`v`是零值，它会返回`false`。如果`IsValid`返回`false`，除`String`外的所有其他方法都会引发panic。大多数函数和方法永远不会返回一个无效的值。如果一个函数或方法确实返回了无效值，那么它的文档会明确地说明这些条件。

#### (Value) IsZero  <- go1.13

``` go 
func (v Value) IsZero() bool
```

IsZero reports whether v is the zero value for its type. It panics if the argument is invalid.

​	`IsZero`方法报告`v`是否为其类型的零值。如果实参无效，则会引发panic。  

#### (Value) Kind 

``` go 
func (v Value) Kind() Kind
```

Kind returns v's Kind. If v is the zero Value (IsValid returns false), Kind returns Invalid.

​	`Kind`方法返回`v`的`Kind`。如果`v`是零值（`IsValid`返回`false`），`Kind`方法返回`Invalid`。  

#### (Value) Len 

``` go 
func (v Value) Len() int
```

Len returns v's length. It panics if v's Kind is not Array, Chan, Map, Slice, String, or pointer to Array.

​	`Len`方法返回`v`的长度。如果`v`的`Kind`不是`Array`、`Chan`、`Map`、`Slice`、`String`或者指向`Array`的指针，它会引发panic。  

#### (Value) MapIndex 

``` go 
func (v Value) MapIndex(key Value) Value
```

MapIndex returns the value associated with key in the map v. It panics if v's Kind is not Map. It returns the zero Value if key is not found in the map or if v represents a nil map. As in Go, the key's value must be assignable to the map's key type.

​	`MapIndex`方法返回在映射`v`中与`key`关联的值。如果`v`的`Kind` 不是`Map`，它会引发panic。如果在映射中找不到`key`，或者`v`代表一个`nil`映射，它会返回零值。就像在Go中一样，`key`的值必须可分配给映射的键类型。  

#### (Value) MapKeys 

``` go 
func (v Value) MapKeys() []Value
```

MapKeys returns a slice containing all the keys present in the map, in unspecified order. It panics if v's Kind is not Map. It returns an empty slice if v represents a nil map.

​	`MapKeys`方法返回一个包含映射中所有存在的key的切片，顺序未指定。如果`v`的`Kind` 不是`Map`，它会引发panic。如果`v`代表一个`nil`映射，它会返回一个空切片。  

#### (Value) MapRange  <- go1.12

``` go 
func (v Value) MapRange() *MapIter
```

MapRange returns a range iterator for a map. It panics if v's Kind is not Map.

​	`MapRange`方法返回一个映射的范围迭代器。如果`v`的`Kind` 不是`Map`，它会引发panic。  

Call Next to advance the iterator, and Key/Value to access each entry. Next returns false when the iterator is exhausted. MapRange follows the same iteration semantics as a range statement.

​	调用`Next`方法来推进迭代器，并调用`Key`/`Value`方法来访问每个条目。当迭代器耗尽时，`Next`方法返回`false`。`MapRange`方法遵循与`range`语句相同的迭代语义。  	

Example:

示例：

```go
iter := reflect.ValueOf(m).MapRange()
for iter.Next() {
	k := iter.Key()
	v := iter.Value()
	...
}
```

#### (Value) Method 

``` go 
func (v Value) Method(i int) Value
```

Method returns a function value corresponding to v's i'th method. The arguments to a Call on the returned function should not include a receiver; the returned function will always use v as the receiver. Method panics if i is out of range or if v is a nil interface value.

​	`Method`方法返回与`v`的第`i`个方法对应的函数值。对返回的函数进行`Call`的实参不应包括接收器；返回的函数将始终使用`v`作为接收器。如果`i`超出范围，或者`v`是一个`nil`接口值，`Method`会引发panic。

#### (Value) MethodByName 

``` go 
func (v Value) MethodByName(name string) Value
```

MethodByName returns a function value corresponding to the method of v with the given name. The arguments to a Call on the returned function should not include a receiver; the returned function will always use v as the receiver. It returns the zero Value if no method was found.

​	`MethodByName`方法根据给定的名称返回与`v`的方法相对应的函数值。对返回的函数进行`Call`的实参不应包含接收器；返回的函数将始终使用`v`作为接收器。如果没有找到对应的方法，它会返回零值。

#### (Value) NumField 

``` go 
func (v Value) NumField() int
```

NumField returns the number of fields in the struct v. It panics if v's Kind is not Struct.

​	`NumField`方法返回结构体`v`中的字段数量。如果`v`的`Kind`不是`Struct`，则会引发panic。

#### (Value) NumMethod 

``` go 
func (v Value) NumMethod() int
```

NumMethod returns the number of methods in the value's method set.

​	`NumMethod`方法返回值的方法集中的方法数量。

For a non-interface type, it returns the number of exported methods.

​	对于非接口类型，它返回导出的方法数量。

For an interface type, it returns the number of exported and unexported methods.

​	对于接口类型，它返回导出和未导出的方法数量。	

#### (Value) OverflowComplex 

``` go 
func (v Value) OverflowComplex(x complex128) bool
```

OverflowComplex reports whether the complex128 x cannot be represented by v's type. It panics if v's Kind is not Complex64 or Complex128.

​	`OverflowComplex`方法报告`complex128` `x`是否无法由`v`的类型表示。如果`v`的Kind 不是`Complex64`或`Complex128`，则会引发panic。

#### (Value) OverflowFloat 

``` go 
func (v Value) OverflowFloat(x float64) bool
```

OverflowFloat reports whether the float64 x cannot be represented by v's type. It panics if v's Kind is not Float32 or Float64.

​	`OverflowFloat`方法报告`float64` `x`是否无法由`v`的类型表示。如果`v`的Kind 不是`Float32`或`Float64`，则会引发panic。

#### (Value) OverflowInt 

``` go 
func (v Value) OverflowInt(x int64) bool
```

OverflowInt reports whether the int64 x cannot be represented by v's type. It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.

​	`OverflowInt`方法报告`int64` `x`是否无法由`v`的类型表示。如果`v`的`Kind` 不是`Int`、`Int8`、`Int16`、`Int32`或`Int64`，则会引发panic。

#### (Value) OverflowUint 

``` go 
func (v Value) OverflowUint(x uint64) bool
```

OverflowUint reports whether the uint64 x cannot be represented by v's type. It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.

​	`OverflowUint`方法报告`uint64` `x`是否无法由`v`的类型表示。如果`v`的`Kind` 不是`Uint`、`Uintptr`、`Uint8`、`Uint16`、`Uint32`或`Uint64`，则会引发panic。

#### (Value) Pointer 

``` go 
func (v Value) Pointer() uintptr
```

Pointer returns v's value as a uintptr. It panics if v's Kind is not Chan, Func, Map, Pointer, Slice, or UnsafePointer.

​	`Pointer`方法将`v`的值作为`uintptr`返回。如果`v`的`Kind`不是`Chan`、`Func`、`Map`、`Pointer`、`Slice`或`UnsafePointer`，它会引发panic。  

If v's Kind is Func, the returned pointer is an underlying code pointer, but not necessarily enough to identify a single function uniquely. The only guarantee is that the result is zero if and only if v is a nil func Value.

​	如果`v`的`Kind`是`Func`，返回的指针是一个底层代码指针，但不一定足以唯一标识一个函数。唯一的保证是，当且仅当`v`是`nil func Value`时，结果为零。  

If v's Kind is Slice, the returned pointer is to the first element of the slice. If the slice is nil the returned value is 0. If the slice is empty but non-nil the return value is non-zero.

​	如果`v`的`Kind`是`Slice`，返回的指针指向切片的第一个元素。如果切片是`nil`，返回的值是`0`。如果切片是空的但非`nil`，返回值是非0。  

It's preferred to use uintptr(Value.UnsafePointer()) to get the equivalent result.

​	建议使用`uintptr(Value.UnsafePointer())`来获得等价的结果。  	

#### (Value) Recv 

``` go 
func (v Value) Recv() (x Value, ok bool)
```

Recv receives and returns a value from the channel v. It panics if v's Kind is not Chan. The receive blocks until a value is ready. The boolean value ok is true if the value x corresponds to a send on the channel, false if it is a zero value received because the channel is closed.

​	`Recv`方法从通道 `v` 中接收并返回一个值。如果 `v` 的 `Kind` 不是 `Chan`，则会 panic。该接收操作会阻塞直到值准备就绪。如果值 `x` 对应于通道的发送，则布尔值 `ok` 为 `true`，如果是一个零值，表示通道已关闭，则为 `false`。

#### (Value) Send 

``` go 
func (v Value) Send(x Value)
```

Send sends x on the channel v. It panics if v's kind is not Chan or if x's type is not the same type as v's element type. As in Go, x's value must be assignable to the channel's element type.

​	`Send`方法在通道 `v` 上发送 `x`。如果 `v` 的 `Kind` 不是 `Chan` 或者 `x` 的类型与 `v` 的元素类型不同，则会 panic。与 Go 语言类似，`x` 的值必须可分配给通道的元素类型。

#### (Value) Set 

``` go 
func (v Value) Set(x Value)
```

Set assigns x to the value v. It panics if CanSet returns false. As in Go, x's value must be assignable to v's type and must not be derived from an unexported field.

​	`Set`方法将 `x` 赋值给值 `v`。如果 `CanSet` 方法返回 `false`，则会 panic。与 Go 语言类似，`x` 的值必须可分配给 `v` 的类型，且不能是未导出字段派生的。

#### (Value) SetBool 

``` go 
func (v Value) SetBool(x bool)
```

SetBool sets v's underlying value. It panics if v's Kind is not Bool or if CanSet() is false.

​	`SetBool`方法设置 `v` 的底层值为 `x`。如果 `v` 的 `Kind` 不是 `Bool` 或者 `CanSet()` 返回 `false`，则会 panic。

#### (Value) SetBytes 

``` go 
func (v Value) SetBytes(x []byte)
```

SetBytes sets v's underlying value. It panics if v's underlying value is not a slice of bytes.

​	`SetBytes`方法设置 `v` 的底层值为 `x`。如果 `v` 的底层值不是字节切片，则会 panic。

#### (Value) SetCap  <- go1.2

``` go 
func (v Value) SetCap(n int)
```

SetCap sets v's capacity to n. It panics if v's Kind is not Slice or if n is smaller than the length or greater than the capacity of the slice.

​	`SetCap`方法将`v`的容量设置为`n`。如果`v`的`Kind`不是`Slice`，或者`n`小于切片的长度或大于切片的容量，它会引发panic。

#### (Value) SetComplex 

``` go 
func (v Value) SetComplex(x complex128)
```

SetComplex sets v's underlying value to x. It panics if v's Kind is not Complex64 or Complex128, or if CanSet() is false.

​	`SetComplex`方法将 `v` 的底层值设置为 `x`。如果 `v` 的 `Kind` 不是 `Complex64` 或 `Complex128`，或者 `CanSet()` 返回 `false`，则会 panic。

#### (Value) SetFloat 

``` go 
func (v Value) SetFloat(x float64)
```

SetFloat sets v's underlying value to x. It panics if v's Kind is not Float32 or Float64, or if CanSet() is false.

​	`SetFloat`方法将 `v` 的底层值设置为 `x`。如果 `v` 的 `Kind` 不是 `Float32` 或 `Float64`，或者 `CanSet()` 返回 `false`，则会 panic。

#### (Value) SetInt 

``` go 
func (v Value) SetInt(x int64)
```

SetInt sets v's underlying value to x. It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64, or if CanSet() is false.

​	`SetInt`方法将 `v` 的底层值设置为 `x`。如果 `v` 的 `Kind` 不是 `Int`、`Int8`、`Int16`、`Int32` 或 `Int64`，或者 `CanSet()` 返回 `false`，则会 panic。

#### (Value) SetIterKey  <- go1.18

``` go 
func (v Value) SetIterKey(iter *MapIter)
```

SetIterKey assigns to v the key of iter's current map entry. It is equivalent to v.Set(iter.Key()), but it avoids allocating a new Value. As in Go, the key must be assignable to v's type and must not be derived from an unexported field.

​	`SetIterKey`方法将 `iter` 当前映射项的键赋值给 `v`。它等价于 `v.Set(iter.Key())`，但是它避免了分配新值。与 Go 语言类似，键必须可分配给 `v` 的类型，且不能是未导出字段派生的。

#### (Value) SetIterValue  <- go1.18

``` go 
func (v Value) SetIterValue(iter *MapIter)
```

SetIterValue assigns to v the value of iter's current map entry. It is equivalent to v.Set(iter.Value()), but it avoids allocating a new Value. As in Go, the value must be assignable to v's type and must not be derived from an unexported field.

​	`SetIterValue`方法为 `v` 赋值为 `iter` 当前的键值。它等同于 `v.Set(iter.Value())`，但是它避免了分配新的 Value。与 Go 一样，值必须可分配给 v 的类型，且不能从未公开的字段派生。

#### (Value) SetLen 

``` go 
func (v Value) SetLen(n int)
```

SetLen sets v's length to n. It panics if v's Kind is not Slice or if n is negative or greater than the capacity of the slice.

​	`SetLen`方法将 `v` 的长度设置为 `n`。如果 `v` 的 `Kind` 不是 `Slice`，或者 `n` 是负数或大于切片的容量，则会 panic。

#### (Value) SetMapIndex 

``` go 
func (v Value) SetMapIndex(key, elem Value)
```

SetMapIndex sets the element associated with key in the map v to elem. It panics if v's Kind is not Map. If elem is the zero Value, SetMapIndex deletes the key from the map. Otherwise if v holds a nil map, SetMapIndex will panic. As in Go, key's elem must be assignable to the map's key type, and elem's value must be assignable to the map's elem type.

​	`SetMapIndex`方法将与键 `key` 相关联的元素设置为 `elem`。如果 `v` 的 `Kind` 不是 `Map`，则会 panic。如果 `elem` 是零值，则 `SetMapIndex` 会从 map 中删除该`key`。否则，如果 `v` 持有一个 `nil` map，则 `SetMapIndex` 会 panic。与 Go 一样，`key` 的值必须可分配给 map 的键类型，`elem` 的值必须可分配给 map 的值类型。

#### (Value) SetPointer 

``` go 
func (v Value) SetPointer(x unsafe.Pointer)
```

SetPointer sets the [unsafe.Pointer](https://pkg.go.dev/unsafe#Pointer) value v to x. It panics if v's Kind is not UnsafePointer.

​	`SetPointer`方法将 [unsafe.Pointer]({{< ref "/stdLib/unsafe#type-pointer">}}) 值 `x` 分配给 `v`。如果 `v` 的 `Kind` 不是 `UnsafePointer`，则会 panic。

#### (Value) SetString 

``` go 
func (v Value) SetString(x string)
```

SetString sets v's underlying value to x. It panics if v's Kind is not String or if CanSet() is false.

​	`SetString`方法将 `v` 的底层值设置为 `x`。如果 `v` 的 `Kind` 不是 `String` 或 `CanSet()` 为 `false`，则会 panic。

#### (Value) SetUint 

``` go 
func (v Value) SetUint(x uint64)
```

SetUint sets v's underlying value to x. It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64, or if CanSet() is false.

​	`SetUint`方法将 `v` 的底层值设置为 `x`。如果 `v` 的 `Kind` 不是 `Uint`，`Uintptr`，`Uint8`，`Uint16`，`Uint32` 或 `Uint64` 或 `CanSet()` 为 `false`，则会 panic。

#### (Value) SetZero  <- go1.20

``` go 
func (v Value) SetZero()
```

SetZero sets v to be the zero value of v's type. It panics if CanSet returns false.

​	`SetZero`方法将 `v` 设置为 `v` 类型的零值。如果 `CanSet` 返回 `false`，则会 panic。

#### (Value) Slice 

``` go 
func (v Value) Slice(i, j int) Value
```

Slice returns v[i:j]. It panics if v's Kind is not Array, Slice or String, or if v is an unaddressable array, or if the indexes are out of bounds.

​	`Slice`方法返回 `v[i:j]`。如果 `v` 的 `Kind` 不是 `Array`，`Slice` 或 `String`，或者 `v` 是不可寻址的数组，或者索引超出范围，则会 panic。

#### (Value) Slice3  <- go1.2

``` go 
func (v Value) Slice3(i, j, k int) Value
```

Slice3 is the 3-index form of the slice operation: it returns v[i:j:k]. It panics if v's Kind is not Array or Slice, or if v is an unaddressable array, or if the indexes are out of bounds.

​	`Slice3`方法是 `slice` 操作的三个索引形式：它返回 `v[i:j:k]`。如果 `v` 的 `Kind` 不是 `Array` 或 `Slice`，或者 `v` 是不可寻址的数组，或者索引超出范围，则会 panic。

#### (Value) String 

``` go 
func (v Value) String() string
```

String returns the string v's underlying value, as a string. String is a special case because of Go's String method convention. Unlike the other getters, it does not panic if v's Kind is not String. Instead, it returns a string of the form "<T value>" where T is v's type. The fmt package treats Values specially. It does not call their String method implicitly but instead prints the concrete values they hold.

​	`String`方法将`v`的底层值作为`string`返回。`String`方法是一个特殊情况，因为Go的`String`方法约定。与其他的getter不同，如果`v`的`Kind`不是`String`，它不会抛出panic。相反，它返回一个字符串"`<T value>`"，其中`T`是`v`的类型。`fmt`包对`Values`进行了特殊处理。它（指的是`fmt`包）不会隐式调用它们的`String`方法，而是打印它们所持有的具体值。

#### (Value) TryRecv 

``` go 
func (v Value) TryRecv() (x Value, ok bool)
```

TryRecv attempts to receive a value from the channel v but will not block. It panics if v's Kind is not Chan. If the receive delivers a value, x is the transferred value and ok is true. If the receive cannot finish without blocking, x is the zero Value and ok is false. If the channel is closed, x is the zero value for the channel's element type and ok is false.

​	`TryRecv`方法尝试从通道`v`接收一个值，但不会阻塞。如果接收到值，则`x`是传输的值，`ok`为`true`。如果接收无法完成而不阻塞，则`x`是零值，并且`ok`为false。如果通道关闭，则`x`是通道元素类型的零值，`ok`为`false`。如果`v`的`Kind`不是`Chan`，它将panic。

#### (Value) TrySend 

``` go 
func (v Value) TrySend(x Value) bool
```

TrySend attempts to send x on the channel v but will not block. It panics if v's Kind is not Chan. It reports whether the value was sent. As in Go, x's value must be assignable to the channel's element type.

​	`TrySend`方法尝试在通道`v`上发送`x`，但不会阻塞。如果`v`的`Kind`不是`Chan`，它将panic。它报告值是否已发送。与Go中一样，`x`的值必须可分配给通道的元素类型。

#### (Value) Type 

``` go 
func (v Value) Type() Type
```

Type returns v's type.

​	`Type`方法返回`v`的类型。

#### (Value) Uint 

``` go 
func (v Value) Uint() uint64
```

Uint returns v's underlying value, as a uint64. It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.

​	`Uint`方法将`v`的底层值作为`uint64`返回。如果`v`的Kind不是`Uint`、`Uintptr`、`Uint8`、`Uint16`、`Uint32`或`Uint64`，它将引发panic。

#### (Value) UnsafeAddr 

``` go 
func (v Value) UnsafeAddr() uintptr
```

UnsafeAddr returns a pointer to v's data, as a uintptr. It panics if v is not addressable.

​	`UnsafeAddr`方法返回指向`v`的数据的指针，作为`uintptr`。如果`v`不可寻址，则它将引发panic。

It's preferred to use uintptr(Value.Addr().UnsafePointer()) to get the equivalent result.

​	最好使用`uintptr(Value.Addr().UnsafePointer())`来获得等效的结果。

#### (Value) UnsafePointer  <- go1.18

``` go 
func (v Value) UnsafePointer() unsafe.Pointer
```

UnsafePointer returns v's value as a [unsafe.Pointer](https://pkg.go.dev/unsafe#Pointer). It panics if v's Kind is not Chan, Func, Map, Pointer, Slice, or UnsafePointer.

​	`UnsafePointer`方法将`v`的值作为[unsafe.Pointer]({{< ref "/stdLib/unsafe#type-pointer">}})返回。如果`v`的`Kind`不是`Chan`、`Func`、`Map`、`Pointer`、`Slice`或`UnsafePointer`，它将panic。

If v's Kind is Func, the returned pointer is an underlying code pointer, but not necessarily enough to identify a single function uniquely. The only guarantee is that the result is zero if and only if v is a nil func Value.

​	如果`v`的`Kind`是`Func`，则返回的指针是底层的代码指针，但不一定足以唯一地标识单个函数。唯一的保证是，当且仅当`v`是一个`nil`函数`Value`时，结果为零。

If v's Kind is Slice, the returned pointer is to the first element of the slice. If the slice is nil the returned value is nil. If the slice is empty but non-nil the return value is non-nil.

​	如果`v`的`Kind`是`Slice`，则返回的指针指向切片的第一个元素。如果切片是`nil`，则返回值为`nil`。如果切片为空但非`nil`，则返回值为非`nil`。

### type ValueError 

``` go 
type ValueError struct {
	Method string
	Kind   Kind
}
```

A ValueError occurs when a Value method is invoked on a Value that does not support it. Such cases are documented in the description of each method.

​	`ValueError`结构体在对不支持它的`Value`调用`Value`方法时发生。这些情况在每个方法的描述中都有记录。

#### (*ValueError) Error 

``` go 
func (e *ValueError) Error() string
```

## Notes

## Bugs

- FieldByName and related functions consider struct field names to be equal if the names are equal, even if they are unexported names originating in different packages. The practical effect of this is that the result of t.FieldByName("x") is not well defined if the struct type t contains multiple fields named x (embedded from different packages). FieldByName may return one of the fields named x or may report that there are none. See https://golang.org/issue/4876 for more details.
- `FieldByName`和相关函数认为结构体字段名称相等，即使它们是来自不同包的未导出名称。这样的实际影响是，如果结构体类型`t`包含多个名为`x`的字段(来自不同的包)，则`t.FieldByName("x")`的结果不是定义良好的。`FieldByName`可能返回一个名为`x`的字段，也可能报告没有任何字段。有关更多详细信息，请参见[https://golang.org/issue/4876](https://golang.org/issue/4876)。