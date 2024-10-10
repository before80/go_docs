+++
title = "asn1"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/encoding/asn1@go1.23.0](https://pkg.go.dev/encoding/asn1@go1.23.0)

Package asn1 implements parsing of DER-encoded ASN.1 data structures, as defined in ITU-T Rec X.690.

​	Package asn1 实现对 DER 编码的 ASN.1 数据结构的解析，如 ITU-T Rec X.690 中定义的。

See also “A Layman’s Guide to a Subset of ASN.1, BER, and DER,” http://luca.ntop.org/Teaching/Appunti/asn1.html.

​	另请参阅“ASN.1、BER 和 DER 子集的非专业指南”，[http://luca.ntop.org/Teaching/Appunti/asn1.html](http://luca.ntop.org/Teaching/Appunti/asn1.html)。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/asn1/common.go;l=22)

``` go 
const (
	TagBoolean         = 1
	TagInteger         = 2
	TagBitString       = 3
	TagOctetString     = 4
	TagNull            = 5
	TagOID             = 6
	TagEnum            = 10
	TagUTF8String      = 12
	TagSequence        = 16
	TagSet             = 17
	TagNumericString   = 18
	TagPrintableString = 19
	TagT61String       = 20
	TagIA5String       = 22
	TagUTCTime         = 23
	TagGeneralizedTime = 24
	TagGeneralString   = 27
	TagBMPString       = 30
)
```

ASN.1 tags represent the type of the following object.

​	ASN.1 标记表示以下对象的类型。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/asn1/common.go;l=44)

``` go 
const (
	ClassUniversal       = 0
	ClassApplication     = 1
	ClassContextSpecific = 2
	ClassPrivate         = 3
)
```

ASN.1 class types represent the namespace of the tag.

​	ASN.1 类别类型表示标记的命名空间。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/asn1/asn1.go;l=217)

``` go 
var NullBytes = []byte{TagNull, 0}
```

NullBytes contains bytes representing the DER-encoded ASN.1 NULL type.

​	NullBytes 包含表示 DER 编码的 ASN.1 NULL 类型的字节。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/asn1/asn1.go;l=214)

``` go 
var NullRawValue = RawValue{Tag: TagNull}
```

NullRawValue is a RawValue with its Tag set to the ASN.1 NULL type tag (5).

​	NullRawValue 是一个 RawValue，其 Tag 设置为 ASN.1 NULL 类型标记 (5)。

## 函数

### func Marshal 

``` go 
func Marshal(val any) ([]byte, error)
```

Marshal returns the ASN.1 encoding of val.

​	Marshal 返回 val 的 ASN.1 编码。

In addition to the struct tags recognised by Unmarshal, the following can be used:

​	除了 Unmarshal 识别的结构标记外，还可以使用以下标记：

```
ia5:         causes strings to be marshaled as ASN.1, IA5String values
omitempty:   causes empty slices to be skipped
printable:   causes strings to be marshaled as ASN.1, PrintableString values
utf8:        causes strings to be marshaled as ASN.1, UTF8String values
utc:         causes time.Time to be marshaled as ASN.1, UTCTime values
generalized: causes time.Time to be marshaled as ASN.1, GeneralizedTime values
```

### func MarshalWithParams  <- go1.10

``` go 
func MarshalWithParams(val any, params string) ([]byte, error)
```

MarshalWithParams allows field parameters to be specified for the top-level element. The form of the params is the same as the field tags.

​	MarshalWithParams 允许为顶级元素指定字段参数。参数的形式与字段标记相同。

### func Unmarshal 

``` go 
func Unmarshal(b []byte, val any) (rest []byte, err error)
```

Unmarshal parses the DER-encoded ASN.1 data structure b and uses the reflect package to fill in an arbitrary value pointed at by val. Because Unmarshal uses the reflect package, the structs being written to must use upper case field names. If val is nil or not a pointer, Unmarshal returns an error.

​	Unmarshal 解析 DER 编码的 ASN.1 数据结构 b，并使用 reflect 包填充 val 指向的任意值。由于 Unmarshal 使用 reflect 包，因此要写入的结构必须使用大写字段名。如果 val 为 nil 或不是指针，Unmarshal 会返回错误。

After parsing b, any bytes that were leftover and not used to fill val will be returned in rest. When parsing a SEQUENCE into a struct, any trailing elements of the SEQUENCE that do not have matching fields in val will not be included in rest, as these are considered valid elements of the SEQUENCE and not trailing data.

​	在解析 b 之后，任何剩余的未使用以填充 val 的字节都将在 rest 中返回。将 SEQUENCE 解析为结构时，SEQUENCE 中的任何尾随元素（在 val 中没有匹配的字段）都不会包含在 rest 中，因为这些元素被视为 SEQUENCE 的有效元素，而不是尾随数据。

An ASN.1 INTEGER can be written to an int, int32, int64, or *big.Int (from the math/big package). If the encoded value does not fit in the Go type, Unmarshal returns a parse error.

​	ASN.1 INTEGER 可以写入 int、int32、int64 或 *big.Int（来自 math/big 包）。如果编码值不适合 Go 类型，Unmarshal 会返回解析错误。

An ASN.1 BIT STRING can be written to a BitString.

​	ASN.1 BIT STRING 可以写入 BitString。

An ASN.1 OCTET STRING can be written to a []byte.

​	ASN.1 OCTET STRING 可以写入 []byte。

An ASN.1 OBJECT IDENTIFIER can be written to an ObjectIdentifier.

​	ASN.1 对象标识符可以写入 ObjectIdentifier。

An ASN.1 ENUMERATED can be written to an Enumerated.

​	ASN.1 ENUMERATED 可以写入 Enumerated。

An ASN.1 UTCTIME or GENERALIZEDTIME can be written to a time.Time.

​	ASN.1 UTCTIME 或 GENERALIZEDTIME 可以写入 time.Time。

An ASN.1 PrintableString, IA5String, or NumericString can be written to a string.

​	ASN.1 PrintableString、IA5String 或 NumericString 可以写入字符串。

Any of the above ASN.1 values can be written to an interface{}. The value stored in the interface has the corresponding Go type. For integers, that type is int64.

​	上述任何 ASN.1 值都可以写入 interface{}。存储在接口中的值具有相应的 Go 类型。对于整数，该类型为 int64。

An ASN.1 SEQUENCE OF x or SET OF x can be written to a slice if an x can be written to the slice's element type.

​	如果 x 可以写入切片的元素类型，则 ASN.1 SEQUENCE OF x 或 SET OF x 可以写入切片。

An ASN.1 SEQUENCE or SET can be written to a struct if each of the elements in the sequence can be written to the corresponding element in the struct.

​	如果序列中的每个元素都可以写入结构中的相应元素，则 ASN.1 SEQUENCE 或 SET 可以写入结构。

The following tags on struct fields have special meaning to Unmarshal:

​	结构字段上的以下标记对 Unmarshal 具有特殊含义：

```
application specifies that an APPLICATION tag is used
private     specifies that a PRIVATE tag is used
default:x   sets the default value for optional integer fields (only used if optional is also present)
explicit    specifies that an additional, explicit tag wraps the implicit one
optional    marks the field as ASN.1 OPTIONAL
set         causes a SET, rather than a SEQUENCE type to be expected
tag:x       specifies the ASN.1 tag number; implies ASN.1 CONTEXT SPECIFIC
```

When decoding an ASN.1 value with an IMPLICIT tag into a string field, Unmarshal will default to a PrintableString, which doesn't support characters such as '@' and '&'. To force other encodings, use the following tags:

​	将带有 IMPLICIT 标记的 ASN.1 值解码为字符串字段时，Unmarshal 将默认为 PrintableString，它不支持“@”和“&”等字符。要强制使用其他编码，请使用以下标记：

```
ia5     causes strings to be unmarshaled as ASN.1 IA5String values
numeric causes strings to be unmarshaled as ASN.1 NumericString values
utf8    causes strings to be unmarshaled as ASN.1 UTF8String values
```

If the type of the first field of a structure is RawContent then the raw ASN1 contents of the struct will be stored in it.

​	如果结构的第一个字段的类型是 RawContent，则结构的原始 ASN1 内容将存储在其中。

If the name of a slice type ends with “SET” then it’s treated as if the “set” tag was set on it. This results in interpreting the type as a SET OF x rather than a SEQUENCE OF x. This can be used with nested slices where a struct tag cannot be given.

​	如果切片类型的名称以“SET”结尾，则将其视为已在其上设置“set”标记。这会导致将类型解释为 SET OF x 而不是 SEQUENCE OF x。这可用于无法给出结构标记的嵌套切片。

Other ASN.1 types are not supported; if it encounters them, Unmarshal returns a parse error.

​	不支持其他 ASN.1 类型；如果遇到它们，Unmarshal 会返回解析错误。

### func UnmarshalWithParams

```go
func UnmarshalWithParams(b []byte, val any, params string) (rest []byte, err error)
```

UnmarshalWithParams allows field parameters to be specified for the top-level element. The form of the params is the same as the field tags.

​	UnmarshalWithParams 允许为顶级元素指定字段参数。参数的形式与字段标记相同。

## 类型

### type BitString

```go
type BitString struct {
	Bytes     []byte // bits packed into bytes.
	BitLength int    // length in bits.
}
```

BitString is the structure to use when you want an ASN.1 BIT STRING type. A bit string is padded up to the nearest byte in memory and the number of valid bits is recorded. Padding bits will be zero.

​	当您想要 ASN.1 BIT STRING 类型时，BitString 是要使用的结构。位串在内存中填充到最接近的字节，并记录有效位的数量。填充位将为零。

#### (BitString) At

```go
func (b BitString) At(i int) int
```

At returns the bit at the given index. If the index is out of range it returns 0.

​	At 返回给定索引处的位。如果索引超出范围，则返回 0。

#### (BitString) RightAlign

```go
func (b BitString) RightAlign() []byte
```

RightAlign returns a slice where the padding bits are at the beginning. The slice may share memory with the BitString.

​	RightAlign 返回一个切片，其中填充位位于开头。该切片可能与 BitString 共享内存。

### type Enumerated

```go
type Enumerated int
```

An Enumerated is represented as a plain int.

​	枚举表示为普通 int。

### type Flag

```go
type Flag bool
```

A Flag accepts any data and is set to true if present.

​	标志接受任何数据，如果存在，则设置为 true。

### type ObjectIdentifier

```go
type ObjectIdentifier []int
```

An ObjectIdentifier represents an ASN.1 OBJECT IDENTIFIER.
An ObjectIdentifier 表示 ASN.1 OBJECT IDENTIFIER。

#### (ObjectIdentifier) Equal

```go
func (oi ObjectIdentifier) Equal(other ObjectIdentifier) bool
```

Equal reports whether oi and other represent the same identifier.

​	Equal 报告 oi 和 other 是否表示相同的标识符。

#### (ObjectIdentifier) String <- go1.3

```go
func (oi ObjectIdentifier) String() string
```

### type RawContent

```go
type RawContent []byte
```

RawContent is used to signal that the undecoded, DER data needs to be preserved for a struct. To use it, the first field of the struct must have this type. It’s an error for any of the other fields to have this type.

​	RawContent 用于指示未解码的 DER 数据需要为结构保留。要使用它，结构的第一个字段必须具有此类型。任何其他字段具有此类型都是错误的。

### type RawValue

```go
type RawValue struct {
	Class, Tag int
	IsCompound bool
	Bytes      []byte
	FullBytes  []byte // includes the tag and length
}
```

A RawValue represents an undecoded ASN.1 object.

​	RawValue 表示未解码的 ASN.1 对象。

### type StructuralError

```go
type StructuralError struct {
	Msg string
}
```

A StructuralError suggests that the ASN.1 data is valid, but the Go type which is receiving it doesn’t match.

​	StructuralError 表明 ASN.1 数据有效，但接收它的 Go 类型不匹配。

#### (StructuralError) Error

```go
func (e StructuralError) Error() string
```

### type SyntaxError

```go
type SyntaxError struct {
	Msg string
}
```

A SyntaxError suggests that the ASN.1 data is invalid.

​	SyntaxError 表示 ASN.1 数据无效。

#### (SyntaxError) Error 

``` go 
func (e SyntaxError) Error() string
```