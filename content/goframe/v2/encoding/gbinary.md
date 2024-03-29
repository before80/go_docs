+++
title = "gbinary"
date = 2024-03-21T17:49:00+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gbinary](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gbinary)

Package gbinary provides useful API for handling binary/bytes data.

​	软件包 gbinary 提供了用于处理二进制/字节数据的有用 API。

Note that package gbinary encodes the data using LittleEndian in default.

​	请注意，默认情况下，软件包 gbinary 使用 LittleEndian 对数据进行编码。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func BeDecode

```go
func BeDecode(b []byte, values ...interface{}) error
```

#### func BeDecodeToBool

```go
func BeDecodeToBool(b []byte) bool
```

#### func BeDecodeToFloat32

```go
func BeDecodeToFloat32(b []byte) float32
```

#### func BeDecodeToFloat64

```go
func BeDecodeToFloat64(b []byte) float64
```

#### func BeDecodeToInt

```go
func BeDecodeToInt(b []byte) int
```

#### func BeDecodeToInt16

```go
func BeDecodeToInt16(b []byte) int16
```

#### func BeDecodeToInt32

```go
func BeDecodeToInt32(b []byte) int32
```

#### func BeDecodeToInt64

```go
func BeDecodeToInt64(b []byte) int64
```

#### func BeDecodeToInt8

```go
func BeDecodeToInt8(b []byte) int8
```

#### func BeDecodeToString

```go
func BeDecodeToString(b []byte) string
```

#### func BeDecodeToUint

```go
func BeDecodeToUint(b []byte) uint
```

#### func BeDecodeToUint16

```go
func BeDecodeToUint16(b []byte) uint16
```

#### func BeDecodeToUint32

```go
func BeDecodeToUint32(b []byte) uint32
```

#### func BeDecodeToUint64

```go
func BeDecodeToUint64(b []byte) uint64
```

#### func BeDecodeToUint8

```go
func BeDecodeToUint8(b []byte) uint8
```

#### func BeEncode

```go
func BeEncode(values ...interface{}) []byte
```

BeEncode encodes one or multiple `values` into bytes using BigEndian. It uses type asserting checking the type of each value of `values` and internally calls corresponding converting function do the bytes converting.

​	BeEncode 使用 BigEndian 将一个或多个编码 `values` 为字节。它使用类型断言检查每个值的 `values` 类型，并在内部调用相应的转换函数进行字节转换。

It supports common variable type asserting, and finally it uses fmt.Sprintf converting value to string and then to bytes.

​	它支持公共变量类型断言，最后它使用 fmt。Sprintf 将值转换为字符串，然后转换为字节。

#### func BeEncodeBool

```go
func BeEncodeBool(b bool) []byte
```

#### func BeEncodeByLength

```go
func BeEncodeByLength(length int, values ...interface{}) []byte
```

#### func BeEncodeFloat32

```go
func BeEncodeFloat32(f float32) []byte
```

#### func BeEncodeFloat64

```go
func BeEncodeFloat64(f float64) []byte
```

#### func BeEncodeInt

```go
func BeEncodeInt(i int) []byte
```

#### func BeEncodeInt16

```go
func BeEncodeInt16(i int16) []byte
```

#### func BeEncodeInt32

```go
func BeEncodeInt32(i int32) []byte
```

#### func BeEncodeInt64

```go
func BeEncodeInt64(i int64) []byte
```

#### func BeEncodeInt8

```go
func BeEncodeInt8(i int8) []byte
```

#### func BeEncodeString

```go
func BeEncodeString(s string) []byte
```

#### func BeEncodeUint

```go
func BeEncodeUint(i uint) []byte
```

#### func BeEncodeUint16

```go
func BeEncodeUint16(i uint16) []byte
```

#### func BeEncodeUint32

```go
func BeEncodeUint32(i uint32) []byte
```

#### func BeEncodeUint64

```go
func BeEncodeUint64(i uint64) []byte
```

#### func BeEncodeUint8

```go
func BeEncodeUint8(i uint8) []byte
```

#### func BeFillUpSize

```go
func BeFillUpSize(b []byte, l int) []byte
```

BeFillUpSize fills up the bytes `b` to given length `l` using big BigEndian.

​	BeFillUpSize 使用 big BigEndian 将字节填充 `b` 到给定长度 `l` 。

Note that it creates a new bytes slice by copying the original one to avoid changing the original parameter bytes.

​	请注意，它通过复制原始字节切片来创建新的字节切片，以避免更改原始参数字节。

#### func Decode

```go
func Decode(b []byte, values ...interface{}) error
```

#### func DecodeBits

```go
func DecodeBits(bits []Bit) int
```

DecodeBits .does decode bits to int Resolve to int

​	DecodeBits .does 将位解码为 int 解析为 int

#### func DecodeBitsToUint

```go
func DecodeBitsToUint(bits []Bit) uint
```

DecodeBitsToUint .Resolve to uint

​	DecodeBitsToUint 中。解析为 uint

#### func DecodeToBool

```go
func DecodeToBool(b []byte) bool
```

#### func DecodeToFloat32

```go
func DecodeToFloat32(b []byte) float32
```

#### func DecodeToFloat64

```go
func DecodeToFloat64(b []byte) float64
```

#### func DecodeToInt

```go
func DecodeToInt(b []byte) int
```

#### func DecodeToInt16

```go
func DecodeToInt16(b []byte) int16
```

#### func DecodeToInt32

```go
func DecodeToInt32(b []byte) int32
```

#### func DecodeToInt64

```go
func DecodeToInt64(b []byte) int64
```

#### func DecodeToInt8

```go
func DecodeToInt8(b []byte) int8
```

#### func DecodeToString

```go
func DecodeToString(b []byte) string
```

#### func DecodeToUint

```go
func DecodeToUint(b []byte) uint
```

#### func DecodeToUint16

```go
func DecodeToUint16(b []byte) uint16
```

#### func DecodeToUint32

```go
func DecodeToUint32(b []byte) uint32
```

#### func DecodeToUint64

```go
func DecodeToUint64(b []byte) uint64
```

#### func DecodeToUint8

```go
func DecodeToUint8(b []byte) uint8
```

#### func Encode

```go
func Encode(values ...interface{}) []byte
```

#### func EncodeBitsToBytes

```go
func EncodeBitsToBytes(bits []Bit) []byte
```

EncodeBitsToBytes . does encode bits to bytes Convert bits to [] byte, encode from left to right, and add less than 1 byte from 0 to the end.

​	EncodeBitsToBytes 中。将位编码为字节 将位转换为 [] 字节，从左到右编码，从 0 到末尾添加少于 1 个字节。

#### func EncodeBool

```go
func EncodeBool(b bool) []byte
```

#### func EncodeByLength

```go
func EncodeByLength(length int, values ...interface{}) []byte
```

#### func EncodeFloat32

```go
func EncodeFloat32(f float32) []byte
```

#### func EncodeFloat64

```go
func EncodeFloat64(f float64) []byte
```

#### func EncodeInt

```go
func EncodeInt(i int) []byte
```

#### func EncodeInt16

```go
func EncodeInt16(i int16) []byte
```

#### func EncodeInt32

```go
func EncodeInt32(i int32) []byte
```

#### func EncodeInt64

```go
func EncodeInt64(i int64) []byte
```

#### func EncodeInt8

```go
func EncodeInt8(i int8) []byte
```

#### func EncodeString

```go
func EncodeString(s string) []byte
```

#### func EncodeUint

```go
func EncodeUint(i uint) []byte
```

#### func EncodeUint16

```go
func EncodeUint16(i uint16) []byte
```

#### func EncodeUint32

```go
func EncodeUint32(i uint32) []byte
```

#### func EncodeUint64

```go
func EncodeUint64(i uint64) []byte
```

#### func EncodeUint8

```go
func EncodeUint8(i uint8) []byte
```

#### func LeDecode

```go
func LeDecode(b []byte, values ...interface{}) error
```

#### func LeDecodeToBool

```go
func LeDecodeToBool(b []byte) bool
```

#### func LeDecodeToFloat32

```go
func LeDecodeToFloat32(b []byte) float32
```

#### func LeDecodeToFloat64

```go
func LeDecodeToFloat64(b []byte) float64
```

#### func LeDecodeToInt

```go
func LeDecodeToInt(b []byte) int
```

#### func LeDecodeToInt16

```go
func LeDecodeToInt16(b []byte) int16
```

#### func LeDecodeToInt32

```go
func LeDecodeToInt32(b []byte) int32
```

#### func LeDecodeToInt64

```go
func LeDecodeToInt64(b []byte) int64
```

#### func LeDecodeToInt8

```go
func LeDecodeToInt8(b []byte) int8
```

#### func LeDecodeToString

```go
func LeDecodeToString(b []byte) string
```

#### func LeDecodeToUint

```go
func LeDecodeToUint(b []byte) uint
```

#### func LeDecodeToUint16

```go
func LeDecodeToUint16(b []byte) uint16
```

#### func LeDecodeToUint32

```go
func LeDecodeToUint32(b []byte) uint32
```

#### func LeDecodeToUint64

```go
func LeDecodeToUint64(b []byte) uint64
```

#### func LeDecodeToUint8

```go
func LeDecodeToUint8(b []byte) uint8
```

#### func LeEncode

```go
func LeEncode(values ...interface{}) []byte
```

LeEncode encodes one or multiple `values` into bytes using LittleEndian. It uses type asserting checking the type of each value of `values` and internally calls corresponding converting function do the bytes converting.

​	LeEncode 使用 LittleEndian 将一个或多个编码 `values` 为字节。它使用类型断言检查每个值的 `values` 类型，并在内部调用相应的转换函数进行字节转换。

It supports common variable type asserting, and finally it uses fmt.Sprintf converting value to string and then to bytes.

​	它支持公共变量类型断言，最后它使用 fmt。Sprintf 将值转换为字符串，然后转换为字节。

#### func LeEncodeBool

```go
func LeEncodeBool(b bool) []byte
```

#### func LeEncodeByLength

```go
func LeEncodeByLength(length int, values ...interface{}) []byte
```

#### func LeEncodeFloat32

```go
func LeEncodeFloat32(f float32) []byte
```

#### func LeEncodeFloat64

```go
func LeEncodeFloat64(f float64) []byte
```

#### func LeEncodeInt

```go
func LeEncodeInt(i int) []byte
```

#### func LeEncodeInt16

```go
func LeEncodeInt16(i int16) []byte
```

#### func LeEncodeInt32

```go
func LeEncodeInt32(i int32) []byte
```

#### func LeEncodeInt64

```go
func LeEncodeInt64(i int64) []byte
```

#### func LeEncodeInt8

```go
func LeEncodeInt8(i int8) []byte
```

#### func LeEncodeString

```go
func LeEncodeString(s string) []byte
```

#### func LeEncodeUint

```go
func LeEncodeUint(i uint) []byte
```

#### func LeEncodeUint16

```go
func LeEncodeUint16(i uint16) []byte
```

#### func LeEncodeUint32

```go
func LeEncodeUint32(i uint32) []byte
```

#### func LeEncodeUint64

```go
func LeEncodeUint64(i uint64) []byte
```

#### func LeEncodeUint8

```go
func LeEncodeUint8(i uint8) []byte
```

#### func LeFillUpSize

```go
func LeFillUpSize(b []byte, l int) []byte
```

LeFillUpSize fills up the bytes `b` to given length `l` using LittleEndian.

​	LeFillUpSize 使用 LittleEndian 将字节填充 `b` 到给定长度 `l` 。

Note that it creates a new bytes slice by copying the original one to avoid changing the original parameter bytes.

​	请注意，它通过复制原始字节切片来创建新的字节切片，以避免更改原始参数字节。

## 类型

### type Bit

```go
type Bit int8
```

Bit Binary bit (0 | 1)

​	位 二进制位 （0 | 1）

#### func DecodeBytesToBits

```go
func DecodeBytesToBits(bs []byte) []Bit
```

DecodeBytesToBits .Parsing [] byte into character array [] uint8

​	DecodeBytesToBits 中。将 [] 字节解析为字符数组 [] uint8

#### func EncodeBits

```go
func EncodeBits(bits []Bit, i int, l int) []Bit
```

EncodeBits does encode bits return bits Default coding

​	EncodeBits 是否对位进行编码 返回位 默认编码

#### func EncodeBitsWithUint

```go
func EncodeBitsWithUint(bits []Bit, ui uint, l int) []Bit
```

EncodeBitsWithUint . Merge ui bitwise into the bits array and occupy the length bits (Note: binary 0 | 1 digits are stored in the uis array)

​	EncodeBitsWithUint 中。将 ui 按位合并到 bits 数组中并占用长度位（注意：二进制 0 | 1 位数字存储在 uis 数组中）