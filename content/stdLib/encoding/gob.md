+++
title = "gob"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/encoding/gob@go1.23.0](https://pkg.go.dev/encoding/gob@go1.23.0)

Package gob manages streams of gobs - binary values exchanged between an Encoder (transmitter) and a Decoder (receiver). A typical use is transporting arguments and results of remote procedure calls (RPCs) such as those provided by package “net/rpc”.

​	包 gob 管理 gob 流 - 在编码器（发送器）和解码器（接收器）之间交换的二进制值。典型用法是传输远程过程调用 (RPC) 的参数和结果，例如包“net/rpc”提供的那些。

The implementation compiles a custom codec for each data type in the stream and is most efficient when a single Encoder is used to transmit a stream of values, amortizing the cost of compilation.

​	该实现为流中的每个数据类型编译一个自定义编解码器，并且当使用单个编码器来传输值流时最有效，摊销了编译成本。

## Basics 基础

A stream of gobs is self-describing. Each data item in the stream is preceded by a specification of its type, expressed in terms of a small set of predefined types. Pointers are not transmitted, but the things they point to are transmitted; that is, the values are flattened. Nil pointers are not permitted, as they have no value. Recursive types work fine, but recursive values (data with cycles) are problematic. This may change.

​	一串 gobs 是自描述的。串中的每个数据项之前都有一个类型说明，该说明用一组小的预定义类型表示。不会传输指针，但会传输它们指向的内容；也就是说，这些值是扁平化的。不允许使用 nil 指针，因为它们没有值。递归类型工作正常，但递归值（带循环的数据）存在问题。这可能会发生变化。

To use gobs, create an Encoder and present it with a series of data items as values or addresses that can be dereferenced to values. The Encoder makes sure all type information is sent before it is needed. At the receive side, a Decoder retrieves values from the encoded stream and unpacks them into local variables.

​	要使用 gobs，请创建一个 Encoder，并向它提供一系列数据项，这些数据项可以是值或可以解引用为值的地址。Encoder 确保在需要之前发送所有类型信息。在接收端，Decoder 从编码的串中检索值，并将它们解包到本地变量中。

## Types and Values 类型和值

The source and destination values/types need not correspond exactly. For structs, fields (identified by name) that are in the source but absent from the receiving variable will be ignored. Fields that are in the receiving variable but missing from the transmitted type or value will be ignored in the destination. If a field with the same name is present in both, their types must be compatible. Both the receiver and transmitter will do all necessary indirection and dereferencing to convert between gobs and actual Go values. For instance, a gob type that is schematically,

​	源值和目标值/类型不必完全对应。对于结构体，源值中存在但接收变量中不存在的字段（由名称标识）将被忽略。接收变量中存在但传输类型或值中不存在的字段将在目标值中被忽略。如果两个字段具有相同的名称，则它们的类型必须兼容。接收器和发送器都将执行所有必要的间接引用和取消引用，以便在 gob 和实际 Go 值之间进行转换。例如，一个模式为 

```go
struct { A, B int }
```

can be sent from or received into any of these Go types:

的 gob 类型可以发送到或接收自以下任何 Go 类型： 

```go
struct { A, B int }	// the same
*struct { A, B int }	// extra indirection of the struct
struct { *A, **B int }	// extra indirection of the fields
struct { A, B int64 }	// different concrete value type; see below
```

It may also be received into any of these:

​	它还可以接收以下任何类型：

```go
struct { A, B int }	// the same
struct { B, A int }	// ordering doesn't matter; matching is by name
struct { A, B, C int }	// extra field (C) ignored
struct { B int }	// missing field (A) ignored; data will be dropped
struct { B, C int }	// missing field (A) ignored; extra field (C) ignored.
```

Attempting to receive into these types will draw a decode error:

​	尝试接收这些类型将导致解码错误：

```go
struct { A int; B uint }	// change of signedness for B
struct { A int; B float }	// change of type for B
struct { }			// no field names in common
struct { C, D int }		// no field names in common
```

Integers are transmitted two ways: arbitrary precision signed integers or arbitrary precision unsigned integers. There is no int8, int16 etc. discrimination in the gob format; there are only signed and unsigned integers. As described below, the transmitter sends the value in a variable-length encoding; the receiver accepts the value and stores it in the destination variable. Floating-point numbers are always sent using IEEE-754 64-bit precision (see below).

​	整数以两种方式传输：任意精度有符号整数或任意精度无符号整数。gob 格式中没有 int8、int16 等区别；只有有符号整数和无符号整数。如下所述，发送方以可变长度编码发送值；接收方接受该值并将其存储在目标变量中。浮点数始终使用 IEEE-754 64 位精度发送（见下文）。

Signed integers may be received into any signed integer variable: int, int16, etc.; unsigned integers may be received into any unsigned integer variable; and floating point values may be received into any floating point variable. However, the destination variable must be able to represent the value or the decode operation will fail.

​	有符号整数可以接收为任何有符号整数变量：int、int16 等；无符号整数可以接收为任何无符号整数变量；浮点值可以接收为任何浮点变量。但是，目标变量必须能够表示该值，否则解码操作将失败。

Structs, arrays and slices are also supported. Structs encode and decode only exported fields. Strings and arrays of bytes are supported with a special, efficient representation (see below). When a slice is decoded, if the existing slice has capacity the slice will be extended in place; if not, a new array is allocated. Regardless, the length of the resulting slice reports the number of elements decoded.

​	结构体、数组和切片也受支持。结构体仅编码和解码导出的字段。字符串和字节数组受支持，并具有特殊的高效表示形式（见下文）。当解码切片时，如果现有切片具有容量，则该切片将就地扩展；如果没有，则分配一个新数组。无论如何，结果切片的长度都会报告解码的元素数量。

In general, if allocation is required, the decoder will allocate memory. If not, it will update the destination variables with values read from the stream. It does not initialize them first, so if the destination is a compound value such as a map, struct, or slice, the decoded values will be merged elementwise into the existing variables.

​	通常，如果需要分配，解码器将分配内存。如果不是，它将使用从流中读取的值更新目标变量。它不会首先初始化它们，因此如果目标是复合值，例如映射、结构或切片，则解码的值将按元素合并到现有变量中。

Functions and channels will not be sent in a gob. Attempting to encode such a value at the top level will fail. A struct field of chan or func type is treated exactly like an unexported field and is ignored.

​	函数和通道不会在 gob 中发送。尝试在顶层对这样的值进行编码将会失败。chan 或 func 类型的结构字段被视为未导出的字段并被忽略。

Gob can encode a value of any type implementing the GobEncoder or encoding.BinaryMarshaler interfaces by calling the corresponding method, in that order of preference.

​	Gob 可以通过按优先顺序调用相应的方法来对实现 GobEncoder 或 encoding.BinaryMarshaler 接口的任何类型的变量进行编码。

Gob can decode a value of any type implementing the GobDecoder or encoding.BinaryUnmarshaler interfaces by calling the corresponding method, again in that order of preference.

​	Gob 可以通过调用相应的方法对实现 GobDecoder 或 encoding.BinaryUnmarshaler 接口的任何类型的变量进行解码，同样按照优先顺序进行。

## Encoding Details 编码详情

This section documents the encoding, details that are not important for most users. Details are presented bottom-up.

​	本节记录了编码，这些详细信息对大多数用户来说并不重要。详细信息自下而上呈现。

An unsigned integer is sent one of two ways. If it is less than 128, it is sent as a byte with that value. Otherwise it is sent as a minimal-length big-endian (high byte first) byte stream holding the value, preceded by one byte holding the byte count, negated. Thus 0 is transmitted as (00), 7 is transmitted as (07) and 256 is transmitted as (FE 01 00).

​	无符号整数以两种方式之一发送。如果它小于 128，则以具有该值的字节形式发送。否则，它将以最短长度的大端序（高字节在前）字节流形式发送，该字节流保存该值，前面有一个字节保存字节数，取反。因此 0 传输为 (00)，7 传输为 (07)，256 传输为 (FE 01 00)。

A boolean is encoded within an unsigned integer: 0 for false, 1 for true.

​	布尔值在无符号整数中编码：0 表示 false，1 表示 true。

A signed integer, i, is encoded within an unsigned integer, u. Within u, bits 1 upward contain the value; bit 0 says whether they should be complemented upon receipt. The encode algorithm looks like this:

​	有符号整数 i 在无符号整数 u 中编码。在 u 中，从第 1 位开始的位包含该值；第 0 位表示在接收时是否应对其进行补码。编码算法如下所示：

```go
var u uint
if i < 0 {
	u = (^uint(i) << 1) | 1 // complement i, bit 0 is 1
} else {
	u = (uint(i) << 1) // do not complement i, bit 0 is 0
}
encodeUnsigned(u)
```

The low bit is therefore analogous to a sign bit, but making it the complement bit instead guarantees that the largest negative integer is not a special case. For example, -129=^128=(^256»1) encodes as (FE 01 01).

​	因此，低位类似于符号位，但将其设为补码位可确保最大的负整数不是特殊情况。例如，-129=^128=(^256»1) 编码为 (FE 01 01)。

Floating-point numbers are always sent as a representation of a float64 value. That value is converted to a uint64 using math.Float64bits. The uint64 is then byte-reversed and sent as a regular unsigned integer. The byte-reversal means the exponent and high-precision part of the mantissa go first. Since the low bits are often zero, this can save encoding bytes. For instance, 17.0 is encoded in only three bytes (FE 31 40).

​	浮点数字总是作为 float64 值的表示形式发送。该值使用 math.Float64bits 转换为 uint64。然后对 uint64 进行字节反转，并作为常规无符号整数发送。字节反转意味着指数和尾数的高精度部分优先。由于低位通常为零，因此可以节省编码字节。例如，17.0 仅以三个字节（FE 31 40）编码。

Strings and slices of bytes are sent as an unsigned count followed by that many uninterpreted bytes of the value.

​	字符串和字节切片作为无符号计数发送，后跟该值中的许多未解释字节。

All other slices and arrays are sent as an unsigned count followed by that many elements using the standard gob encoding for their type, recursively.

​	所有其他切片和数组作为无符号计数发送，后跟使用其类型的标准 gob 编码的许多元素，递归。

Maps are sent as an unsigned count followed by that many key, element pairs. Empty but non-nil maps are sent, so if the receiver has not allocated one already, one will always be allocated on receipt unless the transmitted map is nil and not at the top level.

​	映射作为无符号计数发送，后跟许多键、元素对。发送空但非 nil 的映射，因此如果接收者尚未分配一个映射，则在收到时始终会分配一个映射，除非传输的映射为 nil 且不在顶层。

In slices and arrays, as well as maps, all elements, even zero-valued elements, are transmitted, even if all the elements are zero.

​	在切片和数组以及映射中，即使所有元素都为零，也会传输所有元素，即使是零值元素。

Structs are sent as a sequence of (field number, field value) pairs. The field value is sent using the standard gob encoding for its type, recursively. If a field has the zero value for its type (except for arrays; see above), it is omitted from the transmission. The field number is defined by the type of the encoded struct: the first field of the encoded type is field 0, the second is field 1, etc. When encoding a value, the field numbers are delta encoded for efficiency and the fields are always sent in order of increasing field number; the deltas are therefore unsigned. The initialization for the delta encoding sets the field number to -1, so an unsigned integer field 0 with value 7 is transmitted as unsigned delta = 1, unsigned value = 7 or (01 07). Finally, after all the fields have been sent a terminating mark denotes the end of the struct. That mark is a delta=0 value, which has representation (00).

​	结构以（字段编号，字段值）对的序列发送。字段值使用其类型的标准 gob 编码递归发送。如果字段具有其类型的零值（数组除外；见上文），则从传输中省略该字段。字段编号由编码结构的类型定义：编码类型的第一个字段是字段 0，第二个是字段 1，依此类推。在对值进行编码时，字段编号经过增量编码以提高效率，并且字段始终按字段编号递增的顺序发送；因此，增量是无符号的。增量编码的初始化将字段编号设置为 -1，因此值为 7 的无符号整数字段 0 传输为无符号增量 = 1，无符号值 = 7 或 (01 07)。最后，在发送所有字段后，终止标记表示结构的结束。该标记是增量=0 的值，表示为 (00)。

Interface types are not checked for compatibility; all interface types are treated, for transmission, as members of a single “interface” type, analogous to int or []byte - in effect they’re all treated as interface{}. Interface values are transmitted as a string identifying the concrete type being sent (a name that must be pre-defined by calling Register), followed by a byte count of the length of the following data (so the value can be skipped if it cannot be stored), followed by the usual encoding of concrete (dynamic) value stored in the interface value. (A nil interface value is identified by the empty string and transmits no value.) Upon receipt, the decoder verifies that the unpacked concrete item satisfies the interface of the receiving variable.

​	接口类型不会检查兼容性；所有接口类型在传输时都作为单个“接口”类型的成员进行处理，类似于 int 或 []byte - 实际上它们都被视为 interface{}。接口值以标识正在发送的具体类型的字符串形式传输（必须通过调用 Register 预定义的名称），后跟紧跟数据的长度的字节计数（因此如果无法存储该值，则可以跳过该值），然后是存储在接口值中的具体（动态）值的通常编码。（nil 接口值由空字符串标识，不传输任何值。）收到后，解码器会验证已解包的具体项是否满足接收变量的接口。

If a value is passed to Encode and the type is not a struct (or pointer to struct, etc.), for simplicity of processing it is represented as a struct of one field. The only visible effect of this is to encode a zero byte after the value, just as after the last field of an encoded struct, so that the decode algorithm knows when the top-level value is complete.

​	如果将值传递给 Encode 并且该类型不是结构（或指向结构的指针等），为了简化处理，它将表示为一个字段的结构。这样做的唯一可见效果是在值之后编码一个零字节，就像在编码结构的最后一个字段之后一样，以便解码算法知道顶级值何时完成。

The representation of types is described below. When a type is defined on a given connection between an Encoder and Decoder, it is assigned a signed integer type id. When Encoder.Encode(v) is called, it makes sure there is an id assigned for the type of v and all its elements and then it sends the pair (typeid, encoded-v) where typeid is the type id of the encoded type of v and encoded-v is the gob encoding of the value v.

​	下面描述了类型的表示。当在编码器和解码器之间的给定连接上定义类型时，会为其分配一个有符号整数类型 ID。当调用 Encoder.Encode(v) 时，它会确保为 v 的类型及其所有元素分配了一个 ID，然后发送对 (typeid, encoded-v)，其中 typeid 是 v 的编码类型的类型 ID，encoded-v 是 v 值的 gob 编码。

To define a type, the encoder chooses an unused, positive type id and sends the pair (-type id, encoded-type) where encoded-type is the gob encoding of a wireType description, constructed from these types:

​	要定义类型，编码器会选择一个未使用的正类型 ID，并发送对 (-type id, encoded-type)，其中 encoded-type 是 wireType 描述的 gob 编码，由以下类型构成：

```go
type wireType struct {
	ArrayT           *ArrayType
	SliceT           *SliceType
	StructT          *StructType
	MapT             *MapType
	GobEncoderT      *gobEncoderType
	BinaryMarshalerT *gobEncoderType
	TextMarshalerT   *gobEncoderType

}
type arrayType struct {
	CommonType
	Elem typeId
	Len  int
}
type CommonType struct {
	Name string // the name of the struct type
	Id  int    // the id of the type, repeated so it's inside the type
}
type sliceType struct {
	CommonType
	Elem typeId
}
type structType struct {
	CommonType
	Field []*fieldType // the fields of the struct.
}
type fieldType struct {
	Name string // the name of the field.
	Id   int    // the type id of the field, which must be already defined
}
type mapType struct {
	CommonType
	Key  typeId
	Elem typeId
}
type gobEncoderType struct {
	CommonType
}
```

If there are nested type ids, the types for all inner type ids must be defined before the top-level type id is used to describe an encoded-v.

​	如果存在嵌套类型 ID，则必须在使用顶级类型 ID 来描述 encoded-v 之前定义所有内部类型 ID 的类型。

For simplicity in setup, the connection is defined to understand these types a priori, as well as the basic gob types int, uint, etc. Their ids are:

​	为了简化设置，连接被定义为先验理解这些类型，以及基本 gob 类型 int、uint 等。它们的 ID 为：

```go
bool        1
int         2
uint        3
float       4
[]byte      5
string      6
complex     7
interface   8
// gap for reserved ids.
WireType    16
ArrayType   17
CommonType  18
SliceType   19
StructType  20
FieldType   21
// 22 is slice of fieldType.
MapType     23
```

Finally, each message created by a call to Encode is preceded by an encoded unsigned integer count of the number of bytes remaining in the message. After the initial type name, interface values are wrapped the same way; in effect, the interface value acts like a recursive invocation of Encode.

​	最后，由 Encode 调用创建的每条消息之前都有一个编码的无符号整数计数，表示消息中剩余的字节数。在初始类型名称之后，接口值以相同的方式进行包装；实际上，接口值就像 Encode 的递归调用。

In summary, a gob stream looks like

​	总之，一个 gob 流看起来像

```go
(byteCount (-type id, encoding of a wireType)* (type id, encoding of a value))*
```

where `*` signifies zero or more repetitions and the type id of a value must be predefined or be defined before the value in the stream.

​	其中 `*` 表示零次或多次重复，并且值的类型 ID 必须预先定义或在流中的值之前定义。

Compatibility: Any future changes to the package will endeavor to maintain compatibility with streams encoded using previous versions. That is, any released version of this package should be able to decode data written with any previously released version, subject to issues such as security fixes. See the Go compatibility document for background: https://golang.org/doc/go1compat

​	兼容性：软件包的任何未来更改都将努力保持与使用以前版本编码的流的兼容性。也就是说，此软件包的任何已发布版本都应该能够解码使用任何以前发布的版本编写的的数据，但需视安全修复等问题而定。有关背景信息，请参阅 Go 兼容性文档：https://golang.org/doc/go1compat

See “Gobs of data” for a design discussion of the gob wire format: https://blog.golang.org/gobs-of-data

​	有关 gob 线路格式的设计讨论，请参阅“大量数据”：https://blog.golang.org/gobs-of-data

## Security 安全

This package is not designed to be hardened against adversarial inputs, and is outside the scope of https://go.dev/security/policy. In particular, the Decoder does only basic sanity checking on decoded input sizes, and its limits are not configurable. Care should be taken when decoding gob data from untrusted sources, which may consume significant resources.

​	此软件包并非旨在抵御对抗性输入，并且超出了 https://go.dev/security/policy 的范围。特别是，Decoder 仅对解码的输入大小进行基本健全性检查，并且其限制不可配置。在解码来自不受信任来源的 gob 数据时应小心，这可能会消耗大量资源。

## Example (Basic) 

This example shows the basic usage of the package: Create an encoder, transmit some values, receive them with a decoder.

​	此示例演示了该软件包的基本用法：创建编码器，传输一些值，使用解码器接收它们。

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func main() {
	// Initialize the encoder and decoder. Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	// Encode (send) some values.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// Decode (receive) and print the values.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)

}

Output:

"Pythagoras": {3, 4}
"Treehouse": {1782, 1841}
```

## Example(EncodeDecode)

This example transmits a value that implements the custom encoding and decoding methods.

​	此示例传输一个实现了自定义编码和解码方法的值。

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// The Vector type has unexported fields, which the package cannot access.
// We therefore write a BinaryMarshal/BinaryUnmarshal method pair to allow us
// to send and receive the type with the gob package. These interfaces are
// defined in the "encoding" package.
// We could equivalently use the locally defined GobEncode/GobDecoder
// interfaces.
type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	// A simple encoding: plain text.
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

// UnmarshalBinary modifies the receiver so it must take a pointer receiver.
func (v *Vector) UnmarshalBinary(data []byte) error {
	// A simple encoding: plain text.
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

// This example transmits a value that implements the custom encoding and decoding methods.
func main() {
	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}

	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(v)

}

Output:

{3 4 5}
```

## Example (Interface)

This example shows how to encode an interface value. The key distinction from regular types is to register the concrete type that implements the interface.

​	此示例演示如何编码接口值。与常规类型的关键区别在于注册实现该接口的具体类型。

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

type Pythagoras interface {
	Hypotenuse() float64
}

// This example shows how to encode an interface value. The key
// distinction from regular types is to register the concrete type that
// implements the interface.
func main() {
	var network bytes.Buffer // Stand-in for the network.

	// We must register the concrete type for the encoder and decoder (which would
	// normally be on a separate machine from the encoder). On each end, this tells the
	// engine which concrete type is being sent that implements the interface.
	gob.Register(Point{})

	// Create an encoder and send some values.
	enc := gob.NewEncoder(&network)
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}

	// Create a decoder and receive some values.
	dec := gob.NewDecoder(&network)
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Println(result.Hypotenuse())
	}

}

// interfaceEncode encodes the interface value into the encoder.
func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	// The encode will fail unless the concrete type has been
	// registered. We registered it in the calling function.

	// Pass pointer to interface so Encode sees (and hence sends) a value of
	// interface type. If we passed p directly it would see the concrete type instead.
	// See the blog post, "The Laws of Reflection" for background.
	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

// interfaceDecode decodes the next interface value from the stream and returns it.
func interfaceDecode(dec *gob.Decoder) Pythagoras {
	// The decode will fail unless the concrete type on the wire has been
	// registered. We registered it in the calling function.
	var p Pythagoras
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return p
}

Output:

5
10
15
```

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Register

```go
func Register(value any)
```

Register records a type, identified by a value for that type, under its internal type name. That name will identify the concrete type of a value sent or received as an interface variable. Only types that will be transferred as implementations of interface values need to be registered. Expecting to be used only during initialization, it panics if the mapping between types and names is not a bijection.

​	Register 记录一个类型，由该类型的某个值标识，在其内部类型名称下。该名称将标识作为接口变量发送或接收的值的具体类型。仅需注册将作为接口值实现而传输的类型。预期仅在初始化期间使用，如果类型与名称之间的映射不是双射，则会引发恐慌。

### func RegisterName

```go
func RegisterName(name string, value any)
```

RegisterName is like Register but uses the provided name rather than the type’s default.

​	RegisterName 与 Register 类似，但使用提供的名称而不是类型的默认名称。

## 类型

### type CommonType

```go
type CommonType struct {
	Name string
	Id   typeId
}
```

CommonType holds elements of all types. It is a historical artifact, kept for binary compatibility and exported only for the benefit of the package’s encoding of type descriptors. It is not intended for direct use by clients.

​	CommonType 保存所有类型的元素。它是一个历史遗留物，保留用于二进制兼容性，并且仅为了类型描述符的包编码而导出。它不适用于客户端直接使用。

### type Decoder

```go
type Decoder struct {
	// contains filtered or unexported fields
}
```

A Decoder manages the receipt of type and data information read from the remote side of a connection. It is safe for concurrent use by multiple goroutines.

​	Decoder 管理从连接的远程端读取的类型和数据信息。它可供多个 goroutine 并发使用。

The Decoder does only basic sanity checking on decoded input sizes, and its limits are not configurable. Take caution when decoding gob data from untrusted sources.

​	Decoder 仅对解码的输入大小执行基本健全性检查，并且其限制不可配置。在从不受信任的来源解码 gob 数据时要小心。

#### func NewDecoder

```go
func NewDecoder(r io.Reader) *Decoder
```

NewDecoder returns a new decoder that reads from the io.Reader. If r does not also implement io.ByteReader, it will be wrapped in a bufio.Reader.

​	NewDecoder 返回一个从 io.Reader 读取的新解码器。如果 r 也没有实现 io.ByteReader，它将被包装在 bufio.Reader 中。

#### (*Decoder) Decode

```go
func (dec *Decoder) Decode(e any) error
```

Decode reads the next value from the input stream and stores it in the data represented by the empty interface value. If e is nil, the value will be discarded. Otherwise, the value underlying e must be a pointer to the correct type for the next data item received. If the input is at EOF, Decode returns io.EOF and does not modify e.

​	Decode 从输入流中读取下一个值，并将其存储在由空接口值表示的数据中。如果 e 为 nil，则将丢弃该值。否则，e 下面的值必须是指向接收到的下一个数据项的正确类型的指针。如果输入处于 EOF，则 Decode 返回 io.EOF 且不修改 e。

#### (*Decoder) DecodeValue

```go
func (dec *Decoder) DecodeValue(v reflect.Value) error
```

DecodeValue reads the next value from the input stream. If v is the zero reflect.Value (v.Kind() == Invalid), DecodeValue discards the value. Otherwise, it stores the value into v. In that case, v must represent a non-nil pointer to data or be an assignable reflect.Value (v.CanSet()) If the input is at EOF, DecodeValue returns io.EOF and does not modify v.

​	DecodeValue 从输入流中读取下一个值。如果 v 是零 reflect.Value (v.Kind() == Invalid)，则 DecodeValue 会丢弃该值。否则，它会将该值存储到 v 中。在这种情况下，v 必须表示指向数据的非 nil 指针，或者是一个可赋值的 reflect.Value (v.CanSet())。如果输入处于 EOF，则 DecodeValue 返回 io.EOF 且不修改 v。

### type Encoder

```go
type Encoder struct {
	// contains filtered or unexported fields
}
```

An Encoder manages the transmission of type and data information to the other side of a connection. It is safe for concurrent use by multiple goroutines.

​	Encoder 管理着类型和数据信息向连接另一端的传输。它可以安全地供多个 goroutine 并发使用。

#### func NewEncoder

```go
func NewEncoder(w io.Writer) *Encoder
```

NewEncoder returns a new encoder that will transmit on the io.Writer.

​	NewEncoder 返回一个将在 io.Writer 上传输的新编码器。

#### (*Encoder) Encode

```go
func (enc *Encoder) Encode(e any) error
```

Encode transmits the data item represented by the empty interface value, guaranteeing that all necessary type information has been transmitted first. Passing a nil pointer to Encoder will panic, as they cannot be transmitted by gob.

​	Encode 传输由空接口值表示的数据项，保证所有必要的类型信息已首先传输。将 nil 指针传递给 Encoder 会引发 panic，因为 gob 无法传输它们。

#### (*Encoder) EncodeValue

```go
func (enc *Encoder) EncodeValue(value reflect.Value) error
```

EncodeValue transmits the data item represented by the reflection value, guaranteeing that all necessary type information has been transmitted first. Passing a nil pointer to EncodeValue will panic, as they cannot be transmitted by gob.

​	EncodeValue 传输由反射值表示的数据项，保证所有必要的类型信息已首先传输。将 nil 指针传递给 EncodeValue 会引发 panic，因为 gob 无法传输它们。

### type GobDecoder

```go
type GobDecoder interface {
	// GobDecode overwrites the receiver, which must be a pointer,
	// with the value represented by the byte slice, which was written
	// by GobEncode, usually for the same concrete type.
	GobDecode([]byte) error
}
```

GobDecoder is the interface describing data that provides its own routine for decoding transmitted values sent by a GobEncoder.

​	GobDecoder 是描述数据的接口，该数据提供自己的例程来解码由 GobEncoder 发送的已传输值。

### type GobEncoder

```go
type GobEncoder interface {
	// GobEncode returns a byte slice representing the encoding of the
	// receiver for transmission to a GobDecoder, usually of the same
	// concrete type.
	GobEncode() ([]byte, error)
}
```

GobEncoder is the interface describing data that provides its own representation for encoding values for transmission to a GobDecoder. A type that implements GobEncoder and GobDecoder has complete control over the representation of its data and may therefore contain things such as private fields, channels, and functions, which are not usually transmissible in gob streams.

​	GobEncoder 是描述数据的接口，该数据提供自己的表示形式，以便对值进行编码以传输到 GobDecoder。实现 GobEncoder 和 GobDecoder 的类型可以完全控制其数据的表示形式，因此可能包含私有字段、通道和函数等通常无法在 gob 流中传输的内容。

Note: Since gobs can be stored permanently, it is good design to guarantee the encoding used by a GobEncoder is stable as the software evolves. For instance, it might make sense for GobEncode to include a version number in the encoding.

​	注意：由于 gob 可以永久存储，因此保证 GobEncoder 使用的编码在软件演进时保持稳定是一个好的设计。例如，GobEncode 将版本号包含在编码中可能是有意义的。