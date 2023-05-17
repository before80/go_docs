+++
title = "reflect"
linkTitle = "reflect"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# reflect

[https://pkg.go.dev/reflect@go1.20.1](https://pkg.go.dev/reflect@go1.20.1)

​	reflect包实现了运行时反射，允许程序操作任意类型的对象。典型用法是将静态类型为interface{}的值传递给TypeOf函数提取其动态类型信息，TypeOf函数返回一个Type。

​	调用ValueOf函数返回一个Value类型的值，表示运行时数据。Zero函数接受一个Type参数，并返回表示该类型零值的Value。

​	请参阅《反射的法则》(The Laws of Reflection)了解Go语言中的反射介绍：https://golang.org/doc/articles/laws_of_reflection.html


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=273)

``` go linenums="1"
const Ptr = Pointer
```

​	Ptr是Pointer种类的旧名称。

## 变量

This section is empty.

## 函数

#### func [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2862) 

``` go linenums="1"
func Copy(dst, src Value) int
```

​	Copy函数将src的内容复制到dst，直到dst被填满或src用尽为止。它返回已复制的元素数。Dst和src都必须是Slice或Array类型，并且dst和src必须具有相同的元素类型。

​	作为特例，如果dst的元素类型为Uint8，则src可以具有String种类。

#### func [DeepEqual](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/deepequal.go;l=228) 

``` go linenums="1"
func DeepEqual(x, y any) bool
```

​	DeepEqual函数报告 x 和 y 是否"深度相等"，定义如下。如果以下情况之一适用，则具有相同类型的两个值是深度相等的。具有不同类型的值永远不会深度相等。

​	当它们对应的元素是深度相等的时，数组值是深度相等的。

​	当它们对应的字段(包括导出的和未导出的字段)是深度相等的时，结构体值是深度相等的。

​	当它们都为 nil 时，函数值是深度相等的；否则它们不是深度相等的。

​	当它们保存深度相等的具体值时，接口值是深度相等的。

​	当以下条件全部为 true 时，Map 值是深度相等的：它们都为 nil 或都不为 nil，它们具有相同的长度，并且它们是同一个 map 对象或它们的相应键(使用 Go 相等性匹配)映射到深度相等的值。

​	当它们使用 Go 的 == 运算符相等时，或者它们指向深度相等的值时，指针值是深度相等的。

​	当以下条件全部为 true 时，Slice 值是深度相等的：它们都为 nil 或都不为 nil，它们具有相同的长度，并且它们指向相同底层数组的相同初始元素(也就是，&x[0] == &y[0])，或它们的相应元素(最多是长度)是深度相等的。请注意，非 nil 的空 Slice 和 nil Slice(例如 []byte{} 和 []byte(nil))不是深度相等的。

​	其他值 - 数字、布尔值、字符串和通道 - 如果它们使用 Go 的 == 运算符相等，则它们是深度相等的。

​	一般而言，DeepEqual函数是 Go 的 == 运算符的递归放宽版。然而，这个想法在没有一些不一致性的情况下是不可能实现的。具体来说，一个值可能不等于它自身，因为它是 func 类型(通常是不可比较的)，或者因为它是浮点数 NaN 值(在浮点比较中不等于它本身)，或者因为它是包含这样的值的数组、结构体或接口。另一方面，指针值始终等于它们自己，即使它们指向或包含这样的问题值，因为它们使用 Go 的 == 运算符比较相等，而这是一个足够的条件来深度相等，而不考虑内容。DeepEqual函数已被定义为使得同样的快捷方式适用于切片和映射：如果 x 和 y 是同一个切片或同一个映射，那么它们无论内容如何都是深度相等的。

​	当 DeepEqual函数遍历数据值时，可能会找到一个循环。DeepEqual函数第二次及后续比较已经比较过的两个指针值时，它将这些值视为相等，而不是检查它们所指向的值。这确保了 DeepEqual函数终止。

#### func [Swapper](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/swapper.go;l=17)  <- go1.8

``` go linenums="1"
func Swapper(slice any) func(i, j int)
```

​	Swapper 返回一个函数，该函数交换所提供的切片中的元素。

​	如果提供的接口不是切片，Swapper 将会 panic。

## 类型

### type [ChanDir](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=354) 

``` go linenums="1"
type ChanDir int
```

​	ChanDir 表示通道类型的方向。

``` go linenums="1"
const (
	RecvDir ChanDir             = 1 << iota // <-chan
	SendDir                                 // chan<-
	BothDir = RecvDir | SendDir             // chan
)
```

#### (ChanDir) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1104) 

``` go linenums="1"
func (d ChanDir) String() string
```

​	String 返回 ChanDir 的字符串形式。

### type [Kind](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=240) 

``` go linenums="1"
type Kind uint
```

​	Kind 表示 Type 所代表的具体类型种类。零 Kind 不是有效的类型。

##### Kind Example

``` go linenums="1"
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



``` go linenums="1"
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

#### (Kind) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=634) 

``` go linenums="1"
func (k Kind) String() string
```

​	String方法返回k的名称。

### type [MapIter](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1820)  <- go1.12

``` go linenums="1"
type MapIter struct {
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

​	MapIter结构体是一个用于遍历映射的迭代器。参见Value.MapRange。

#### (*MapIter) [Key](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1826)  <- go1.12

``` go linenums="1"
func (iter *MapIter) Key() Value
```

​	Key方法返回iter当前映射条目的键。

#### (*MapIter) [Next](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1914)  <- go1.12

``` go linenums="1"
func (iter *MapIter) Next() bool
```

​	Next方法将映射迭代器前进，并报告是否有另一个条目。当iter耗尽时，它返回false；对Key，Value或Next的后续调用将引发panic。

#### (*MapIter) [Reset](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1933)  <- go1.18

``` go linenums="1"
func (iter *MapIter) Reset(v Value)
```

​	Reset方法修改iter以遍历v。如果v的Kind不是Map且v不是零值，则它会引发panic。Reset(Value{})会导致iter不引用任何映射，这可能允许之前遍历过的映射被垃圾回收。

#### (*MapIter) [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1869)  <- go1.12

``` go linenums="1"
func (iter *MapIter) Value() Value
```

​	Value方法返回iter当前映射条目的值。

### type [Method](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=606) 

``` go linenums="1"
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

​	Method表示单个方法。

#### (Method) [IsExported](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=623)  <- go1.17

``` go linenums="1"
func (m Method) IsExported() bool
```

​	IsExported 返回该方法是否为导出方法。

### type [SelectCase](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2957)  <- go1.1

``` go linenums="1"
type SelectCase struct {
	Dir  SelectDir // case 方向
	Chan Value     // 使用的通道 (用于发送或接收)
	Send Value     // 发送的值 (用于发送)
}
```

​	SelectCase 描述 select 操作中的一个 case。case 的类型取决于 Dir，通信方向。

​	如果 Dir 是 SelectDefault，则 case 表示默认情况。Chan 和 Send 必须是零值。

​	如果 Dir 是 SelectSend，则 case 表示发送操作。通常 Chan 的基础值必须是一个通道，Send 的基础值必须可以分配给通道的元素类型。作为特例，如果 Chan 是零值，则忽略该 case，Send 字段也将被忽略，并且可以是零值或非零值。

​	如果 Dir 是 SelectRecv，则 case 表示接收操作。通常 Chan 的基础值必须是一个通道，Send 必须是零值。如果 Chan 是零值，则忽略该 case，但 Send 仍必须是零值。当选择接收操作时，Select 将返回接收到的值。

### type [SelectDir](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2930)  <- go1.1

``` go linenums="1"
type SelectDir int
```

​	SelectDir 描述 select case 的通信方向。

``` go linenums="1"
const (
	SelectSend    SelectDir // case Chan <- Send
	SelectRecv              // case <-Chan:
	SelectDefault           // default
)
```

### type [SliceHeader](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2760) 

``` go linenums="1"
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

​	SliceHeader 是切片的运行时表示。它不能安全或便携地使用，并且它的表示可能在以后的版本中更改。此外，Data 字段不足以保证它引用的数据不会被垃圾回收，因此程序必须保持一个单独的、正确类型的指向底层数据的指针。

​	在新代码中，请改用 unsafe.Slice 或 unsafe.SliceData。

### type [StringHeader](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2747) 

``` go linenums="1"
type StringHeader struct {
	Data uintptr
	Len  int
}
```

​	StringHeader是一个字符串的底层运行时表示形式。它不能安全或便携地使用，并且它的表示形式可能会在以后的版本中更改。此外，Data字段不足以保证它引用的数据不会被垃圾回收，因此程序必须保持一个单独的、正确类型的指向底层数据的指针。

​	在新代码中，应使用unsafe.String或unsafe.StringData代替。

### type [StructField](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1154) 

``` go linenums="1"
type StructField struct {
	// Name 是字段名。
	Name string

	// PkgPath 是限定小写(未公开)字段名的包路径。
    // 对于大写(公开)字段名，它为空。
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath string

	Type      Type      // 字段类型
	Tag       StructTag // 字段标记字符串
	Offset    uintptr   // 在结构体内的偏移量(以字节为单位)
	Index     []int     // 用于Type.FieldByIndex的索引序列
	Anonymous bool      // 是否为嵌入字段
}
```

​	StructField结构体描述结构体中的一个字段。

#### func [VisibleFields](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/visiblefields.go;l=16)  <- go1.17

``` go linenums="1"
func VisibleFields(t Type) []StructField
```

​	VisibleFields返回t中所有可见的字段，t必须是结构体类型。如果可以直接通过FieldByName调用访问字段，则定义字段为可见字段。返回的字段包括匿名结构成员内部的字段和未公开的字段。它们遵循在结构中找到的相同顺序，其中匿名字段紧随其提升的字段之后。

​	对于返回的切片中的每个元素e，可以通过调用v.FieldByIndex(e.Index)从类型为t的值v中检索相应的字段。

#### (StructField) [IsExported](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1171)  <- go1.17

``` go linenums="1"
func (f StructField) IsExported() bool
```

​	IsExported报告字段是否已公开。

### type [StructTag](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1183) 

``` go linenums="1"
type StructTag string
```

​	StructTag 是结构体字段的标签字符串。

By convention, tag strings are a concatenation of optionally space-separated key:"value" pairs. Each key is a non-empty string consisting of non-control characters other than space (U+0020 ' '), quote (U+0022 '"'), and colon (U+003A ':'). Each value is quoted using U+0022 '"' characters and Go string literal syntax.

​	按照惯例，标签字符串是可选地用空格分隔的 key:"value" 键值对的串联。每个键是由非空字符组成的字符串，不能包含空格(U+0020 ' ')、引号(U+0022 '"')和冒号(U+003A ':')以外的控制字符。每个值都使用 U+0022 '"' 字符引用并采用 Go 字符串文字语法进行引用。

##### StructTag Example

``` go linenums="1"
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



#### (StructTag) [Get](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1190) 

``` go linenums="1"
func (tag StructTag) Get(key string) string
```

​	Get方法在标签字符串中返回与键关联的值。如果标签中没有这样的键，则 Get 返回空字符串。如果标签不具有常规格式，则 Get 返回的值是未指定的。要确定标记是否显式设置为空字符串，请使用 Lookup。

#### (StructTag) [Lookup](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1201)  <- go1.7

``` go linenums="1"
func (tag StructTag) Lookup(key string) (value string, ok bool)
```

​	Lookup方法在标签字符串中返回与键关联的值。如果标签中存在该键，则返回值(可能为空)。否则，返回的值将是空字符串。ok 返回值报告值是否在标签字符串中显式设置。如果标签不具有常规格式，则 Lookup 返回的值是未指定的。

##### Lookup Example

``` go linenums="1"
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



### type [Type](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=38) 

``` go linenums="1"
type Type interface {
	// Align方法返回该类型在内存中分配时所需的字节对齐方式。
	Align() int

	// FieldAlign方法返回该类型在结构体字段中使用时所需的字节对齐方式。
	FieldAlign() int

	// Method方法返回类型方法集中第i个方法。
    // 如果i不在[0，NumMethod())范围内，
    // 则会出现恐慌。
	// 对于非接口类型T或T，
    // 返回的Method的Type和Func字段描述的是第一个参数是接收器的函数，
    // 只有导出的方法可访问。
	// 对于接口类型，
    // 返回的Method的Type字段给出了方法签名，
    // 没有接收器，而Func字段为nil。
	// 方法按字典顺序排序。
	Method(int) Method

	// MethodByName方法返回类型方法集中具有该名称的方法以及指示方法是否被找到的布尔值。
	// 对于非接口类型T或T，
    // 返回的Method的Type和Func字段描述的是第一个参数是接收器的函数。
	// 对于接口类型，
    // 返回的Method的Type字段给出了方法签名，
    // 没有接收器，而Func字段为nil。
	MethodByName(string) (Method, bool)

	// NumMethod方法返回使用Method可访问的方法数。
	// 对于非接口类型，它返回导出的方法的数量。
	// 对于接口类型，它返回导出和未导出方法的数量。
	NumMethod() int

	// Name方法返回定义类型的类型在其包中的名称。
    // 对于其他(非定义)类型，它返回空字符串。	
	Name() string

	// PkgPath方法返回已定义类型的包路径，
    // 即唯一标识包的导入路径，例如"encoding/base64"。
    // 如果类型是预声明的(string，error)或未定义(T，struct{}，
    // []int或A是非定义类型的别名，其中A是别名)，
    // 则包路径将为空字符串。
	PkgPath() string

	// Size方法返回存储给定类型值所需的字节数；类似于unsafe.Sizeof。
	Size() uintptr

	// String方法返回类型的字符串表示形式。
    // 字符串表示可以使用缩短的包名(例如，base64代替"encoding/base64")，
    // 并且不能保证在类型中唯一。
	// 要测试类型标识，请直接比较Types。
	String() string

	// Kind方法返回此类型的具体种类。
	Kind() Kind

	// Implements方法报告类型是否实现了接口类型u。
	Implements(u Type) bool

	// AssignableTo方法报告该类型的值是否可分配给类型u。
	AssignableTo(u Type) bool

	// ConvertibleTo reports whether a value of the type is convertible to type u.
	// Even if ConvertibleTo returns true, the conversion may still panic.
	// For example, a slice of type []T is convertible to *[N]T,
	// but the conversion will panic if its length is less than N.
    // ConvertibleTo 报告该类型的值是否可以转换为u类型。
	// 即使ConvertibleTo返回true，转换仍然可能发生恐慌。
	// 例如，一个[]T类型的片断可以转换为*[N]T，但是如果它的长度小于N的话，转换就会发生恐慌。
	ConvertibleTo(u Type) bool

	// ConvertibleTo方法报告该类型的值是否可转换为类型u。
    // 即使ConvertibleTo返回true，
    // 转换仍可能出现恐慌。
    // 例如，类型[]T的切片可转换为[N]T，
    // 但是如果其长度小于N，则转换将导致恐慌。
	Comparable() bool

	// Bits方法返回类型以位为单位的大小。
    // 如果类型的种类不是有大小或无大小的Int，Uint，Float或Complex，
    // 则它会出现恐慌。
	Bits() int

	// ChanDir返回通道类型的方向。
	// 如果类型的种类不是Chan，则会引发panic。
	ChanDir() ChanDir

	// IsVariadic报告函数类型的最后一个输入参数是否为"…"参数。
    // 如果是，则t.In(t.NumIn()-1)返回参数的隐式实际类型[]T。
	//
	// 具体来说，如果t表示func(x int，y ...float64)，则
	//
	// 	t.NumIn()== 2
	// 	t.In(0)是"int"的reflect.Type
	// 	t.In(1)是"[]float64"的reflect.Type
	// 	t.IsVariadic()== true
	//
	// 如果类型的种类不是Func，则IsVariadic会引发panic。
	IsVariadic() bool

	// Elem返回类型的元素类型。
	// 如果类型的种类不是Array，Chan，Map，Pointer或Slice，则会引发panic。
	Elem() Type

	// Field返回结构体类型的第i个字段。
	// 如果类型的种类不是Struct，则会引发panic。
	// 如果i不在[0，NumField()]范围内，则会引发panic。
	Field(i int) StructField

	// FieldByIndex返回相应于索引序列的嵌套字段。
    // 等效于依次调用Field(i)。
	// 如果类型的种类不是Struct，则会引发panic。
	FieldByIndex(index []int) StructField

	// FieldByName返回具有给定名称的结构体字段和一个布尔值，
    // 指示是否找到了该字段。
	// 如果类型的种类不是Struct，则会引发panic。
	FieldByName(name string) (StructField, bool)

	// FieldByNameFunc返回具有满足匹配函数的名称的结构字段和一个布尔值，
    // 指示是否找到了该字段。
	//
	// FieldByNameFunc在结构体本身中考虑字段，
    // 然后在任何嵌入式结构体中考虑字段，
    // 以广度优先的顺序停止于包含一个或多个满足匹配函数的字段的最浅嵌套深度。
    // 如果该深度处的多个字段都满足匹配函数，
    // 则它们相互取消，
    // 并且FieldByNameFunc不返回任何匹配项。
	// 这种行为与Go处理包含嵌入字段的结构体的名称查找相同。
	FieldByNameFunc(match func(string) bool) (StructField, bool)

	// In返回函数类型的第i个输入参数的类型。
	// 如果类型的Kind不是Func，则会panic。
	// 如果i不在[0，NumIn())范围内，则会panic。
	In(i int) Type

	// Key返回映射类型的键类型。
	// 如果类型的Kind不是Map，则会panic。
	Key() Type

	// Len返回数组类型的长度。
	// 如果类型的Kind不是Array，则会panic。
	Len() int

	// NumField返回结构体类型的字段数。
	// 如果类型的Kind不是Struct，则会panic。
	NumField() int

	// NumIn返回函数类型的输入参数计数。
	// 如果类型的Kind不是Func，则会panic。
	NumIn() int

	// NumOut返回函数类型的输出参数计数。
	// 如果类型的Kind不是Func，则会panic。
	NumOut() int

	// Out返回函数类型的第i个输出参数的类型。
	// 如果类型的Kind不是Func，则会panic。
	// 如果i不在[0，NumOut())范围内，则会panic。
	Out(i int) Type
	// 包含已过滤或未导出的方法
}
```

​	Type 是 Go 语言类型的表示。

​	并不是所有的方法都适用于所有类型。如果有限制，则在每个方法的文档中进行说明。在调用特定于类型种类的方法之前，请使用 Kind 方法查找类型的种类。调用不适用于类型种类的方法会导致运行时恐慌。

​	Type 值是可比较的，例如使用 == 运算符，因此它们可以用作映射键。如果它们表示相同的类型，则两个 Type 值相等。

#### func [ArrayOf](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=2899)  <- go1.5

``` go linenums="1"
func ArrayOf(length int, elem Type) Type
```

​	ArrayOf函数返回具有给定长度和元素类型的数组类型。例如，如果 t 表示 int，则 ArrayOf(5, t) 表示 [5]int。

​	如果结果类型的大小大于可用地址空间，则 ArrayOf 会发生恐慌。

#### func [ChanOf](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1865)  <- go1.1

``` go linenums="1"
func ChanOf(dir ChanDir, t Type) Type
```

​	ChanOf函数返回具有给定方向和元素类型的通道类型。例如，如果 t 表示 int，则 ChanOf(RecvDir, t) 表示 <-chan int。

​	gc 运行时对通道元素类型施加了 64 kB 的限制。如果 t 的大小等于或超过此限制，则 ChanOf 会发生恐慌。

#### func [FuncOf](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=2030)  <- go1.5

``` go linenums="1"
func FuncOf(in, out []Type, variadic bool) Type
```

​	FuncOf函数返回具有给定参数和结果类型的函数类型。例如，如果 k 表示 int，e 表示 string，则 FuncOf([]Type{k}, []Type{e}, false) 表示 func(int) string。

​	可变参数控制函数是否为变参。如果 in[len(in)-1] 不表示一个切片并且 variadic 为 true，则 FuncOf 会发生恐慌。

#### func [MapOf](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1928)  <- go1.1

``` go linenums="1"
func MapOf(key, elem Type) Type
```

​	MapOf函数返回具有给定键和元素类型的映射类型。例如，如果 k 表示 int，e 表示 string，则 MapOf(k, e) 表示 map[int]string。

​	如果键类型不是有效的映射键类型(即，如果它不实现 Go 的 == 运算符)，则 MapOf 会发生恐慌。

#### func [PointerTo](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1461)  <- go1.18

``` go linenums="1"
func PointerTo(t Type) Type
```

​	PointerTo函数返回带有元素类型t的指针类型。例如，如果t表示类型Foo，则PointerTo(t)表示`*Foo`。

#### func [PtrTo](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1457) 

``` go linenums="1"
func PtrTo(t Type) Type
```

​	PtrTo返回带有元素类型t的指针类型。例如，如果t表示类型Foo，则PtrTo(t)表示`*Foo`。

​	PtrTo函数是PointerTo的旧拼写。这两个函数的行为相同。

#### func [SliceOf](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=2347)  <- go1.1

``` go linenums="1"
func SliceOf(t Type) Type
```

​	SliceOf函数返回具有元素类型t的切片类型。例如，如果t表示int，则SliceOf(t)表示[]int。

#### func [StructOf](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=2428)  <- go1.7

``` go linenums="1"
func StructOf(fields []StructField) Type
```

​	StructOf返回包含字段的结构类型。忽略Offset和Index字段，计算它们时，它们的值将与编译器生成的值相同。

​	StructOf函数当前不会为嵌入字段生成包装器方法，并且如果传递未公开的StructFields，则会引发错误。这些限制在将来的版本中可能会被解除。

##### StructOf Example

``` go linenums="1"
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



#### func [TypeOf](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/type.go;l=1438) 

``` go linenums="1"
func TypeOf(i any) Type
```

​	TypeOf函数返回表示i的动态类型的反射类型。如果i是nil接口值，则TypeOf返回nil。

#### typeOf Example

``` go linenums="1"
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// 由于接口类型只用于静态类型，
    // 因此通常可以使用*Foo值来查找接口类型Foo的反射Type的惯用语。
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))

}
Output:

true
```



### type [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=39) 

``` go linenums="1"
type Value struct {
	//包含已过滤或未导出的字段
}
```

​	Value结构体是表示Go值的反射接口。

​	不是所有的方法都适用于所有类型的值。每个方法的文档中都有限制，如果有的话。在调用特定于类型的方法之前，请使用Kind方法找出类型的种类。调用不适合类型种类的方法会导致运行时panic。

​	零值表示没有值。它的IsValid方法返回false，它的Kind方法返回Invalid，它的String方法返回"`<invalid Value>`"，所有其他方法都会panic。大多数函数和方法从不返回无效值。如果有，则它的文档明确说明条件。

​	Value结构体可以被多个goroutine并发使用，前提是底层的Go值可以被用于等效的直接操作。

​	要比较两个值，请比较Interface方法的结果。使用==比较两个值不会比较它们表示的底层值。

#### func [Append](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2832) 

``` go linenums="1"
func Append(s Value, x ...Value) Value
```

​	Append函数将值x附加到切片s并返回结果切片。与Go语言一样，x的每个值都必须可分配给切片的元素类型。

#### func [AppendSlice](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2844) 

``` go linenums="1"
func AppendSlice(s, t Value) Value
```

​	AppendSlice函数将切片t附加到切片s并返回结果切片。切片s和t必须具有相同的元素类型。

#### func [Indirect](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3130) 

``` go linenums="1"
func Indirect(v Value) Value
```

​	Indirect函数返回v所指向的值。如果v是一个nil指针，则Indirect返回零值。如果v不是一个指针，则Indirect返回v。

#### func [MakeChan](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3096) 

``` go linenums="1"
func MakeChan(typ Type, buffer int) Value
```

​	MakeChan函数创建一个具有指定类型和缓冲区大小的新通道。

#### func [MakeFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/makefunc.go;l=46)  <- go1.1

``` go linenums="1"
func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
```

​	MakeFunc函数返回一个给定Type的新函数，该函数包装了fn函数。当调用该新函数时，它会执行以下操作：

- 将其参数转换为Value切片。
- 运行results := fn(args)。
- 将结果作为值切片返回，每个切片元素表示一个形式上的结果。

​	fn实现可以假设Type参数描述了函数参数的数量和类型。如果Type描述了一个可变参数函数，则最后一个值本身是一个表示可变参数的切片，就像可变参数函数的主体中一样。fn返回的结果Value切片必须具有由Type给出的结果数量和类型。

​	Value.Call方法允许调用者根据Value值调用类型化函数；相比之下，MakeFunc允许调用者根据Value值实现类型化函数。

​	文档的示例部分包括了如何使用MakeFunc为不同类型构建交换函数的示例。

##### MakeFunc Example

``` go linenums="1"
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// swap 是传递给 MakeFunc 的实现函数。
	// 它必须使用 reflect.Values 来工作，以便能够编写不事先知道类型的代码。
	// will be.
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap 期望 fptr 是一个指向 nil 函数的指针。
	// 它将该指针设置为使用 MakeFunc 创建的新函数。
	// 当调用该函数时，reflect 将参数转换为 Values，
	// 调用 swap，然后将 swap 的结果切片转换为新函数返回的值。
	makeSwap := func(fptr any) {
		// fptr 是一个函数的指针。
		// 将该函数值本身(可能为 nil)作为 reflect.Value 获取，
		// 以便我们可以查询其类型，然后设置该值。
		fn := reflect.ValueOf(fptr).Elem()

		// 制作正确类型的函数。
		v := reflect.MakeFunc(fn.Type(), swap)

		// 将其分配给 fn 所代表的值。
		fn.Set(v)
	}

	// 制作并调用一个用于 int 的 swap 函数。
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// 制作并调用一个用于 float64 的 swap 函数。
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))

}
Output:

1 0
3.14 2.72
```



#### func [MakeMap](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3112) 

``` go linenums="1"
func MakeMap(typ Type) Value
```

​	MakeMap函数创建一个指定类型的新map。

#### func [MakeMapWithSize](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3118)  <- go1.9

``` go linenums="1"
func MakeMapWithSize(typ Type, n int) Value
```

​	MakeMapWithSize函数创建一个指定类型和大约n个元素的初始空间的新map。

#### func [MakeSlice](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3077) 

``` go linenums="1"
func MakeSlice(typ Type, len, cap int) Value
```

​	MakeSlice函数为指定的slice类型、长度和容量创建一个新的零值初始化的slice。

#### func [New](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3184) 

``` go linenums="1"
func New(typ Type) Value
```

​	New函数返回表示指向指定类型的新零值的指针的Value。也就是，返回的Value的Type为PointerTo(typ)。

#### func [NewAt](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3201) 

``` go linenums="1"
func NewAt(typ Type, p unsafe.Pointer) Value
```

​	NewAt函数返回一个表示指向指定类型值的指针的Value，使用p作为该指针。

#### func [Select](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2971)  <- go1.1

``` go linenums="1"
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
```

​	Select执行由cases列表描述的select操作。与Go select语句一样，它会阻塞，直到至少有一个case可以进行，然后做出统一的伪随机选择，然后执行该case。它返回选择的case的索引以及(如果该case是接收操作)接收到的值和一个布尔值，指示该值是否对应于对通道的发送(而不是接收到的零值，因为通道已关闭)。Select最多支持65536个case。

#### func [ValueOf](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3139) 

``` go linenums="1"
func ValueOf(i any) Value
```

​	ValueOf函数返回一个初始化为接口i存储的具体值的新Value。ValueOf(nil)返回零Value。

#### func [Zero](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3158) 

``` go linenums="1"
func Zero(typ Type) Value
```

​	Zero函数返回表示指定类型的零值的Value。结果与Value结构体的零值不同，后者表示没有值。例如，Zero(TypeOf(42))返回一个Kind为Int且值为0的Value。返回的值既不可寻址也不可设置。

#### (Value) [Addr](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=271) 

``` go linenums="1"
func (v Value) Addr() Value
```

​	Addr方法返回表示v地址的指针值。如果CanAddr()返回false，则会引发恐慌。Addr通常用于获取结构体字段或切片元素的指针，以便调用需要指针接收器的方法。

#### (Value) [Bool](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=283) 

``` go linenums="1"
func (v Value) Bool() bool
```

​	Bool方法返回v的底层值。如果v的kind不是Bool，则会引发恐慌。

#### (Value) [Bytes](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=300) 

``` go linenums="1"
func (v Value) Bytes() []byte
```

​	Bytes方法返回v的底层值。如果v的底层值不是字节的切片或可寻址的字节数组，则会引发恐慌。

#### (Value) [Call](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=367) 

``` go linenums="1"
func (v Value) Call(in []Value) []Value
```

​	Call方法使用输入参数in调用函数v。例如，如果len(in)==3，则v.Call(in)表示Go调用v(in[0]、in[1]、in[2])。如果v的Kind不是Func，则会引发恐慌。它返回输出结果作为Values。与Go一样，每个输入参数必须可分配给函数对应的输入参数的类型。如果v是可变参数函数，则Call创建可变参数切片参数本身，并复制相应的值。

#### (Value) [CallSlice](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=380) 

``` go linenums="1"
func (v Value) CallSlice(in []Value) []Value
```

​	CallSlice方法调用可变参数函数 v，并将 in 作为输入参数，将 in[len(in)-1] 作为 v 的最后一个可变参数赋值。例如，如果 len(in) == 3，则 v.CallSlice(in) 表示 Go 调用 v(in[0], in[1], in[2]...)。如果 v 的 Kind 不是 Func 或 v 不是可变参数，则 CallSlice 会引发 panic。它将输出结果作为 Values 返回。与 Go 语言相同，每个输入参数必须可以分配到函数的相应输入参数的类型。

#### (Value) [CanAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=346) 

``` go linenums="1"
func (v Value) CanAddr() bool
```

​	CanAddr方法报告值的地址是否可以使用 Addr 获取。这样的值称为可寻址值。如果它是 slice 的元素，是可寻址数组的元素，是可寻址结构的字段或是指针的解引用结果，它就是可寻址的。如果 CanAddr 返回 false，则调用 Addr 将引发 panic。

#### (Value) [CanComplex](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1188)  <- go1.18

``` go linenums="1"
func (v Value) CanComplex() bool
```

​	CanComplex方法报告是否可以使用 Complex 而不会引发 panic。

#### (Value) [CanConvert](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3264)  <- go1.17

``` go linenums="1"
func (v Value) CanConvert(t Type) bool
```

​	CanConvert方法报告值 v 是否可以转换为类型 t。如果 v.CanConvert(t) 返回 true，则 v.Convert(t) 不会引发 panic。

#### (Value) [CanFloat](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1364)  <- go1.18

``` go linenums="1"
func (v Value) CanFloat() bool
```

​	CanFloat方法报告是否可以使用 Float 而不会引发 panic。

#### (Value) [CanInt](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1435)  <- go1.18

``` go linenums="1"
func (v Value) CanInt() bool
```

​	CanInt方法报告是否可以使用 Int 而不会引发 panic。

#### (Value) [CanInterface](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1465) 

``` go linenums="1"
func (v Value) CanInterface() bool
```

​	CanInterface方法报告是否可以使用 Interface 而不会引发 panic。

#### (Value) [CanSet](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=355) 

``` go linenums="1"
func (v Value) CanSet() bool
```

​	CanSet方法报告是否可以更改 v 的值。只有可寻址的值且没有通过使用未公开的结构字段获得的值才能更改。如果 CanSet 返回 false，则调用 Set 或任何类型特定的 setter(例如 SetBool、SetInt)将引发 panic。

#### (Value) [CanUint](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2638)  <- go1.18

``` go linenums="1"
func (v Value) CanUint() bool
```

​	CanUint方法报告是否可以使用Uint而不会引起panic。

#### (Value) [Cap](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1155) 

``` go linenums="1"
func (v Value) Cap() int
```

​	Cap方法返回v的容量。如果v的Kind不是Array、Chan、Slice或指向Array的指针，则会引发panic。

#### (Value) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1181) 

``` go linenums="1"
func (v Value) Close()
```

​	Close方法关闭通道v。如果v的Kind不是Chan，则会引发panic。

#### (Value) [Comparable](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3289)  <- go1.20

``` go linenums="1"
func (v Value) Comparable() bool
```

​	Comparable方法报告值v是否可比较。如果v的类型是接口，则会检查动态类型。如果此报告为true，则v.Interface()==x对于任何x都不会引发panic，也不会对于任何Value u引发v.Equal(u)。

#### (Value) [Complex](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1199) 

``` go linenums="1"
func (v Value) Complex() complex128
```

​	Complex方法返回v的底层值，作为complex128。如果v的Kind不是Complex64或Complex128，则会引发panic。

#### (Value) [Convert](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3251)  <- go1.1

``` go linenums="1"
func (v Value) Convert(t Type) Value
```

​	Convert方法返回值v转换为类型t后的值。如果通常的Go转换规则不允许将值v转换为类型t，或者将v转换为类型t会引发panic，则Convert会引发panic。

#### (Value) [Elem](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1214) 

``` go linenums="1"
func (v Value) Elem() Value
```

​	Elem方法返回接口v包含的值或指针v指向的值。如果v的Kind不是Interface或Pointer，则会引发panic。如果v为nil，则返回零值。

#### (Value) [Equal](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=3331)  <- go1.20

``` go linenums="1"
func (v Value) Equal(u Value) bool
```

​	Equal方法报告v是否等于u。对于两个无效值，Equal将报告为true。对于接口值，Equal将比较接口中的值。否则，如果值具有不同的类型，则Equal将报告false。否则，对于数组和结构体，Equal将按顺序比较每个元素，并在找到不相等的元素时报告false。在所有比较期间，如果比较相同类型的值，并且类型不可比较，则Equal将引发panic。

#### (Value) [Field](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1266) 

``` go linenums="1"
func (v Value) Field(i int) Value
```

​	Field方法返回v的第i个结构体字段的值。如果v的Kind不是Struct或i越界，则会出现panic。

#### (Value) [FieldByIndex](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1299) 

``` go linenums="1"
func (v Value) FieldByIndex(index []int) Value
```

​	FieldByIndex方法返回对应于索引的嵌套字段。如果需要通过nil指针或不是结构体的字段进行步进，则会出现panic。

##### FieldByIndex Example

``` go linenums="1"
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



#### (Value) [FieldByIndexErr](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1322)  <- go1.18

``` go linenums="1"
func (v Value) FieldByIndexErr(index []int) (Value, error)
```

​	FieldByIndexErr方法返回对应于索引的嵌套字段。如果需要通过nil指针进行步进，则返回错误，但如果必须通过不是结构体的字段进行步进，则会出现panic。

#### (Value) [FieldByName](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1344) 

``` go linenums="1"
func (v Value) FieldByName(name string) Value
```

​	FieldByName方法返回具有给定名称的结构体字段。如果未找到字段，则返回零值。如果v的Kind不是struct，则会出现panic。

##### FieldByName Example

``` go linenums="1"
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





#### (Value) [FieldByNameFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1356) 

``` go linenums="1"
func (v Value) FieldByNameFunc(match func(string) bool) Value
```

​	FieldByNameFunc方法返回满足match函数的名称的结构体字段。如果v的Kind不是struct，则会出现panic。如果未找到字段，则返回零值。

#### (Value) [Float](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1375) 

``` go linenums="1"
func (v Value) Float() float64
```

​	Float方法返回v的基础值，作为float64。如果v的Kind不是Float32或Float64，则会出现panic。

#### (Value) [Grow](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2789)  <- go1.20

``` go linenums="1"
func (v Value) Grow(n int)
```

​	Grow方法增加切片的容量，如果需要的话，以保证为另外n个元素腾出空间。在Grow(n)之后，可以将至少n个元素附加到切片中而不进行另一个分配。

​	如果v的Kind不是Slice或n为负数或太大以至于无法分配内存，则会出现panic。

#### (Value) [Index](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1390) 

``` go linenums="1"
func (v Value) Index(i int) Value
```

​	Index方法返回v的第i个元素。如果v的Kind不是Array，Slice或String，或者i越界，则会panic。

#### (Value) [Int](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1446) 

``` go linenums="1"
func (v Value) Int() int64
```

​	Int方法返回v的基础值，作为int64。如果v的Kind不是Int，Int8，Int16，Int32或Int64，则会panic。

#### (Value) [Interface](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1479) 

``` go linenums="1"
func (v Value) Interface() (i any)
```

​	Interface方法将v的当前值作为interface{}返回。它等同于：

``` go linenums="1"
var i interface{} = (v's underlying value)
```

​	如果Value是通过访问不公开的struct字段获得的，则会panic。

#### (Value) [IsNil](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1539) 

``` go linenums="1"
func (v Value) IsNil() bool
```

​	IsNil方法报告其参数v是否为nil。参数必须是chan，func，interface，map，pointer或slice值；如果不是，IsNil会panic。请注意，IsNil并不总是等价于在Go中使用nil的常规比较。例如，如果v是通过使用未初始化的接口变量i调用ValueOf创建的，则i == nil将为true，但v.IsNil将会panic，因为v将是零值。

#### (Value) [IsValid](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1564) 

``` go linenums="1"
func (v Value) IsValid() bool
```

​	IsValid方法报告v是否表示一个值。如果v是零值，则返回false。如果IsValid返回false，则所有其他方法除了String之外都会panic。大多数函数和方法都不会返回无效的Value。如果返回，则其文档明确说明条件。

#### (Value) [IsZero](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1570)  <- go1.13

``` go linenums="1"
func (v Value) IsZero() bool
```

​	IsZero方法报告v是否为其类型的零值。如果参数无效，则会panic。

#### (Value) [Kind](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1682) 

``` go linenums="1"
func (v Value) Kind() Kind
```

​	Kind方法返回v的Kind。如果v是零值(IsValid返回false)，则Kind返回Invalid。

#### (Value) [Len](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1688) 

``` go linenums="1"
func (v Value) Len() int
```

​	Len方法返回v的长度。如果v的Kind不是Array，Chan，Map，Slice，String或指向Array的指针，则会panic。

#### (Value) [MapIndex](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1723) 

``` go linenums="1"
func (v Value) MapIndex(key Value) Value
```

​	MapIndex 方法返回 map v 中关联于 key 的值。若 v 的 Kind 不是 Map，则会 panic。若 key 不存在于 map 中或 v 代表一个 nil 的 map，则返回 zero Value。与 Go 语言类似，key 的值必须能赋值给 map 的 key 类型。

#### (Value) [MapKeys](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1762) 

``` go linenums="1"
func (v Value) MapKeys() []Value
```

​	MapKeys 方法返回 map 中所有键的 slice，顺序不确定。若 v 的 Kind 不是 Map，则会 panic。若 v 代表一个 nil 的 map，则返回一个空的 slice。

#### (Value) [MapRange](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1956)  <- go1.12

``` go linenums="1"
func (v Value) MapRange() *MapIter
```

​	MapRange 方法返回一个 map 的 range 迭代器。若 v 的 Kind 不是 Map，则会 panic。

​	调用 Next 方法来推进迭代器，调用 Key 和 Value 方法来访问每个 entry。当迭代器耗尽时，Next 返回 false。MapRange 遵循与 range 语句相同的迭代语义。

Example:

```
iter := reflect.ValueOf(m).MapRange()
for iter.Next() {
	k := iter.Key()
	v := iter.Value()
	...
}
```

#### (Value) [Method](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1988) 

``` go linenums="1"
func (v Value) Method(i int) Value
```

​	Method 方法返回 v 的第 i 个方法对应的函数值。调用返回函数时，不需要包含一个接收器；返回的函数总是使用 v 作为接收器。若 i 超出了范围或 v 是 nil 的接口值，则 Method 会 panic。

#### (Value) [MethodByName](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2024) 

``` go linenums="1"
func (v Value) MethodByName(name string) Value
```

​	MethodByName 方法返回名称为 name 的 v 方法对应的函数值。调用返回函数时，不需要包含一个接收器；返回的函数总是使用 v 作为接收器。如果没有找到方法，则返回 zero Value。

#### (Value) [NumField](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2040) 

``` go linenums="1"
func (v Value) NumField() int
```

​	NumField 方法返回结构体 v 中的字段数。若 v 的 Kind 不是 Struct，则会 panic。

#### (Value) [NumMethod](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2009) 

``` go linenums="1"
func (v Value) NumMethod() int
```

​	NumMethod方法返回值的方法集中的方法数量。

​	对于非接口类型，它返回导出方法的数量。

​	对于接口类型，它返回导出和非导出方法的数量。

#### (Value) [OverflowComplex](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2048) 

``` go linenums="1"
func (v Value) OverflowComplex(x complex128) bool
```

​	OverflowComplex方法报告 complex128 类型 x 是否无法被 v 的类型所表示。如果 v 的 Kind 不是 Complex64 或 Complex128，则会出现 panic。

#### (Value) [OverflowFloat](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2061) 

``` go linenums="1"
func (v Value) OverflowFloat(x float64) bool
```

​	OverflowFloat方法报告 float64 类型 x 是否无法被 v 的类型所表示。如果 v 的 Kind 不是 Float32 或 Float64，则会出现 panic。

#### (Value) [OverflowInt](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2081) 

``` go linenums="1"
func (v Value) OverflowInt(x int64) bool
```

​	OverflowInt方法报告 int64 类型 x 是否无法被 v 的类型所表示。如果 v 的 Kind 不是 Int、Int8、Int16、Int32 或 Int64，则会出现 panic。

#### (Value) [OverflowUint](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2094) 

``` go linenums="1"
func (v Value) OverflowUint(x uint64) bool
```

​	OverflowUint方法报告 uint64 类型 x 是否无法被 v 的类型所表示。如果 v 的 Kind 不是 Uint、Uintptr、Uint8、Uint16、Uint32 或 Uint64，则会出现 panic。

#### (Value) [Pointer](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2123) 

``` go linenums="1"
func (v Value) Pointer() uintptr
```

​	Pointer方法返回 v 的值作为 uintptr。如果 v 的 Kind 不是 Chan、Func、Map、Pointer、Slice 或 UnsafePointer，则会出现 panic。

​	如果 v 的 Kind 是 Func，则返回的指针是底层代码指针，但不一定足以唯一地标识单个函数。唯一的保证是如果 v 是 nil func Value，则结果为零。

​	如果 v 的 Kind 是 Slice，则返回的指针是切片的第一个元素。如果切片是 nil，则返回的值为 0。如果切片为空但非 nil，则返回值是非零的。

​	推荐使用 uintptr(Value.UnsafePointer()) 来获取等效的结果。

#### (Value) [Recv](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2168) 

``` go linenums="1"
func (v Value) Recv() (x Value, ok bool)
```

​	Recv方法从通道 v 中接收并返回一个值。如果 v 的 Kind 不是 Chan，则会 panic。该接收操作会阻塞直到值准备就绪。如果值 x 对应于通道的发送操作，则布尔值 ok 为 true，如果是一个零值，表示通道已关闭，则为 false。

#### (Value) [Send](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2201) 

``` go linenums="1"
func (v Value) Send(x Value)
```

​	Send方法在通道 v 上发送 x。如果 v 的 Kind 不是 Chan 或者 x 的类型与 v 的元素类型不同，则会 panic。与 Go 语言类似，x 的值必须可分配给通道的元素类型。

#### (Value) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2229) 

``` go linenums="1"
func (v Value) Set(x Value)
```

​	Set方法将 x 赋值给值 v。如果 CanSet 返回 false，则会 panic。与 Go 语言类似，x 的值必须可分配给 v 的类型，且不能是未公开字段派生的。

#### (Value) [SetBool](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2250) 

``` go linenums="1"
func (v Value) SetBool(x bool)
```

​	SetBool方法设置 v 的底层值为 x。如果 v 的 Kind 不是 Bool 或者 CanSet() 返回 false，则会 panic。

#### (Value) [SetBytes](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2258) 

``` go linenums="1"
func (v Value) SetBytes(x []byte)
```

​	SetBytes方法设置 v 的底层值为 x。如果 v 的底层值不是字节切片，则会 panic。

#### (Value) [SetCap](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2342)  <- go1.2

``` go linenums="1"
func (v Value) SetCap(n int)
```

​	SetCap方法将 v 的容量设置为 n。如果 v 的 Kind 不是 Slice 或者 n 小于切片的长度或大于切片的容量，则会 panic。

#### (Value) [SetComplex](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2280) 

``` go linenums="1"
func (v Value) SetComplex(x complex128)
```

​	SetComplex方法将 v 的底层值设置为 x。如果 v 的 Kind 不是 Complex64 或 Complex128，或者 CanSet() 返回 false，则会 panic。

#### (Value) [SetFloat](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2294) 

``` go linenums="1"
func (v Value) SetFloat(x float64)
```

​	SetFloat方法将 v 的底层值设置为 x。如果 v 的 Kind 不是 Float32 或 Float64，或者 CanSet() 返回 false，则会 panic。

#### (Value) [SetInt](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2308) 

``` go linenums="1"
func (v Value) SetInt(x int64)
```

​	SetInt方法将 v 的底层值设置为 x。如果 v 的 Kind 不是 Int、Int8、Int16、Int32 或 Int64，或者 CanSet() 返回 false，则会 panic。

#### (Value) [SetIterKey](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1844)  <- go1.18

``` go linenums="1"
func (v Value) SetIterKey(iter *MapIter)
```

​	SetIterKey方法将 iter 当前映射项的键赋值给 v。它等价于 v.Set(iter.Key())，但是它避免了分配新值。与 Go 语言类似，键必须可分配给 v 的类型，且不能是未公开字段派生的。

#### (Value) [SetIterValue](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1887)  <- go1.18

``` go linenums="1"
func (v Value) SetIterValue(iter *MapIter)
```

​	SetIterValue方法为 v 赋值为 iter 当前的键值。它等同于 v.Set(iter.Value())，但是它避免了分配新的 Value。与 Go 一样，值必须可分配给 v 的类型，且不能从未公开的字段派生。

#### (Value) [SetLen](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2329) 

``` go linenums="1"
func (v Value) SetLen(n int)
```

​	SetLen方法将 v 的长度设置为 n。如果 v 的 Kind 不是 Slice，或者 n 是负数或大于切片的容量，则会 panic。

#### (Value) [SetMapIndex](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2358) 

``` go linenums="1"
func (v Value) SetMapIndex(key, elem Value)
```

​	SetMapIndex方法将与键 key 相关联的元素设置为 elem。如果 v 的 Kind 不是 Map，则会 panic。如果 elem 是零值，则 SetMapIndex 会从 map 中删除该键。否则，如果 v 持有一个 nil map，则 SetMapIndex 会 panic。与 Go 一样，key 的值必须可分配给 map 的键类型，elem 的值必须可分配给 map 的值类型。

#### (Value) [SetPointer](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2428) 

``` go linenums="1"
func (v Value) SetPointer(x unsafe.Pointer)
```

​	SetPointer方法将 unsafe.Pointer 值 x 分配给 v。如果 v 的 Kind 不是 UnsafePointer，则会 panic。

#### (Value) [SetString](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2436) 

``` go linenums="1"
func (v Value) SetString(x string)
```

​	SetString方法将 v 的底层值设置为 x。如果 v 的 Kind 不是 String 或 CanSet() 为 false，则会 panic。

#### (Value) [SetUint](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2406) 

``` go linenums="1"
func (v Value) SetUint(x uint64)
```

​	SetUint方法将 v 的底层值设置为 x。如果 v 的 Kind 不是 Uint，Uintptr，Uint8，Uint16，Uint32 或 Uint64 或 CanSet() 为 false，则会 panic。

#### (Value) [SetZero](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=1628)  <- go1.20

``` go linenums="1"
func (v Value) SetZero()
```

​	SetZero方法将 v 设置为 v 类型的零值。如果 CanSet 返回 false，则会 panic。

#### (Value) [Slice](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2445) 

``` go linenums="1"
func (v Value) Slice(i, j int) Value
```

​	Slice方法返回 v[i:j]。如果 v 的 Kind 不是 Array，Slice 或 String，或者 v 是不可寻址的数组，或者索引超出范围，则会 panic。

#### (Value) [Slice3](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2507)  <- go1.2

``` go linenums="1"
func (v Value) Slice3(i, j, k int) Value
```

​	Slice3方法是 slice 操作的三个索引形式：它返回 v[i:j:k]。如果 v 的 Kind 不是 Array 或 Slice，或者 v 是不可寻址的数组，或者索引超出范围，则会 panic。

#### (Value) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2562) 

``` go linenums="1"
func (v Value) String() string
```

​	String方法返回值v的基础值作为字符串。String方法是一个特殊情况，因为Go的String方法约定。与其他的getter不同，如果v的Kind不是String，它不会抛出panic。相反，它返回一个字符串"<T value>"，其中T是v的类型。fmt包对Values进行了特殊处理。它不会隐式调用它们的String方法，而是打印它们所持有的具体值。

#### (Value) [TryRecv](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2584) 

``` go linenums="1"
func (v Value) TryRecv() (x Value, ok bool)
```

​	TryRecv方法尝试从通道v接收一个值，但不会阻塞。如果接收到值，则x是传输的值，ok为true。如果接收无法完成而不阻塞，则x是零值，并且ok为false。如果通道关闭，则x是通道元素类型的零值，ok为false。如果v的Kind不是Chan，它将panic。

#### (Value) [TrySend](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2594) 

``` go linenums="1"
func (v Value) TrySend(x Value) bool
```

​	TrySend方法尝试在通道v上发送x，但不会阻塞。如果v的Kind不是Chan，它将panic。它报告值是否已发送。与Go中一样，x的值必须可分配给通道的元素类型。

#### (Value) [Type](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2601) 

``` go linenums="1"
func (v Value) Type() Type
```

​	Type方法返回v的类型。

#### (Value) [Uint](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2649) 

``` go linenums="1"
func (v Value) Uint() uint64
```

​	Uint方法返回v的基础值作为uint64。如果v的Kind不是Uint、Uintptr、Uint8、Uint16、Uint32或Uint64，它将panic。

#### (Value) [UnsafeAddr](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2678) 

``` go linenums="1"
func (v Value) UnsafeAddr() uintptr
```

​	UnsafeAddr方法返回指向v的数据的指针，作为uintptr。如果v不可寻址，则它将panic。

​	最好使用uintptr(Value.Addr()。UnsafePointer())来获得等效的结果。

#### (Value) [UnsafePointer](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=2699)  <- go1.18

``` go linenums="1"
func (v Value) UnsafePointer() unsafe.Pointer
```

​	UnsafePointer方法返回v的值作为unsafe.Pointer。如果v的Kind不是Chan、Func、Map、Pointer、Slice或UnsafePointer，它将panic。

​	如果v的Kind是Func，则返回的指针是底层的代码指针，但不一定足以唯一地标识单个函数。唯一的保证是，当且仅当v是一个nil函数Value时，结果为零。

​	如果v的Kind是Slice，则返回的指针指向切片的第一个元素。如果切片是nil，则返回值为0。如果切片为空但非nil，则返回值为非零。

### type [ValueError](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=163) 

``` go linenums="1"
type ValueError struct {
	Method string
	Kind   Kind
}
```

​	ValueError结构体在对不支持它的Value调用Value方法时发生。这些情况在每个方法的描述中都有记录。

#### (*ValueError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/reflect/value.go;l=168) 

``` go linenums="1"
func (e *ValueError) Error() string
```

## Notes

## Bugs

- FieldByName和相关函数认为结构体字段名称相等，即使它们是来自不同包的未导出名称。这样的实际影响是，如果结构体类型t包含多个名为x的字段(来自不同的包)，则t.FieldByName("x")的结果不是定义良好的。FieldByName可能返回一个名为x的字段，也可能报告没有任何字段。有关更多详细信息，请参见https://golang.org/issue/4876。