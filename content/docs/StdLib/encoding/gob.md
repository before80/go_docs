+++
title = "gob"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# gob

https://pkg.go.dev/encoding/gob@go1.20.1








Package gob manages streams of gobs - binary values exchanged between an Encoder (transmitter) and a Decoder (receiver). A typical use is transporting arguments and results of remote procedure calls (RPCs) such as those provided by package "net/rpc".

The implementation compiles a custom codec for each data type in the stream and is most efficient when a single Encoder is used to transmit a stream of values, amortizing the cost of compilation.

#### Basics 

A stream of gobs is self-describing. Each data item in the stream is preceded by a specification of its type, expressed in terms of a small set of predefined types. Pointers are not transmitted, but the things they point to are transmitted; that is, the values are flattened. Nil pointers are not permitted, as they have no value. Recursive types work fine, but recursive values (data with cycles) are problematic. This may change.

To use gobs, create an Encoder and present it with a series of data items as values or addresses that can be dereferenced to values. The Encoder makes sure all type information is sent before it is needed. At the receive side, a Decoder retrieves values from the encoded stream and unpacks them into local variables.

### 类型

The source and destination values/types need not correspond exactly. For structs, fields (identified by name) that are in the source but absent from the receiving variable will be ignored. Fields that are in the receiving variable but missing from the transmitted type or value will be ignored in the destination. If a field with the same name is present in both, their types must be compatible. Both the receiver and transmitter will do all necessary indirection and dereferencing to convert between gobs and actual Go values. For instance, a gob type that is schematically,

```
struct { A, B int }
```

can be sent from or received into any of these Go types:

```
struct { A, B int }	// the same
*struct { A, B int }	// extra indirection of the struct
struct { *A, **B int }	// extra indirection of the fields
struct { A, B int64 }	// different concrete value type; see below
```

It may also be received into any of these:

```
struct { A, B int }	// the same
struct { B, A int }	// ordering doesn't matter; matching is by name
struct { A, B, C int }	// extra field (C) ignored
struct { B int }	// missing field (A) ignored; data will be dropped
struct { B, C int }	// missing field (A) ignored; extra field (C) ignored.
```

Attempting to receive into these types will draw a decode error:

```
struct { A int; B uint }	// change of signedness for B
struct { A int; B float }	// change of type for B
struct { }			// no field names in common
struct { C, D int }		// no field names in common
```

Integers are transmitted two ways: arbitrary precision signed integers or arbitrary precision unsigned integers. There is no int8, int16 etc. discrimination in the gob format; there are only signed and unsigned integers. As described below, the transmitter sends the value in a variable-length encoding; the receiver accepts the value and stores it in the destination variable. Floating-point numbers are always sent using IEEE-754 64-bit precision (see below).

Signed integers may be received into any signed integer variable: int, int16, etc.; unsigned integers may be received into any unsigned integer variable; and floating point values may be received into any floating point variable. However, the destination variable must be able to represent the value or the decode operation will fail.

Structs, arrays and slices are also supported. Structs encode and decode only exported fields. Strings and arrays of bytes are supported with a special, efficient representation (see below). When a slice is decoded, if the existing slice has capacity the slice will be extended in place; if not, a new array is allocated. Regardless, the length of the resulting slice reports the number of elements decoded.

In general, if allocation is required, the decoder will allocate memory. If not, it will update the destination variables with values read from the stream. It does not initialize them first, so if the destination is a compound value such as a map, struct, or slice, the decoded values will be merged elementwise into the existing variables.

Functions and channels will not be sent in a gob. Attempting to encode such a value at the top level will fail. A struct field of chan or func type is treated exactly like an unexported field and is ignored.

Gob can encode a value of any type implementing the GobEncoder or encoding.BinaryMarshaler interfaces by calling the corresponding method, in that order of preference.

Gob can decode a value of any type implementing the GobDecoder or encoding.BinaryUnmarshaler interfaces by calling the corresponding method, again in that order of preference.

#### Encoding Details 

This section documents the encoding, details that are not important for most users. Details are presented bottom-up.

An unsigned integer is sent one of two ways. If it is less than 128, it is sent as a byte with that value. Otherwise it is sent as a minimal-length big-endian (high byte first) byte stream holding the value, preceded by one byte holding the byte count, negated. Thus 0 is transmitted as (00), 7 is transmitted as (07) and 256 is transmitted as (FE 01 00).

A boolean is encoded within an unsigned integer: 0 for false, 1 for true.

A signed integer, i, is encoded within an unsigned integer, u. Within u, bits 1 upward contain the value; bit 0 says whether they should be complemented upon receipt. The encode algorithm looks like this:

``` go 
var u uint
if i < 0 {
	u = (^uint(i) << 1) | 1 // complement i, bit 0 is 1
} else {
	u = (uint(i) << 1) // do not complement i, bit 0 is 0
}
encodeUnsigned(u)
```

The low bit is therefore analogous to a sign bit, but making it the complement bit instead guarantees that the largest negative integer is not a special case. For example, -129=^128=(^256>>1) encodes as (FE 01 01).

Floating-point numbers are always sent as a representation of a float64 value. That value is converted to a uint64 using math.Float64bits. The uint64 is then byte-reversed and sent as a regular unsigned integer. The byte-reversal means the exponent and high-precision part of the mantissa go first. Since the low bits are often zero, this can save encoding bytes. For instance, 17.0 is encoded in only three bytes (FE 31 40).

Strings and slices of bytes are sent as an unsigned count followed by that many uninterpreted bytes of the value.

All other slices and arrays are sent as an unsigned count followed by that many elements using the standard gob encoding for their type, recursively.

Maps are sent as an unsigned count followed by that many key, element pairs. Empty but non-nil maps are sent, so if the receiver has not allocated one already, one will always be allocated on receipt unless the transmitted map is nil and not at the top level.

In slices and arrays, as well as maps, all elements, even zero-valued elements, are transmitted, even if all the elements are zero.

Structs are sent as a sequence of (field number, field value) pairs. The field value is sent using the standard gob encoding for its type, recursively. If a field has the zero value for its type (except for arrays; see above), it is omitted from the transmission. The field number is defined by the type of the encoded struct: the first field of the encoded type is field 0, the second is field 1, etc. When encoding a value, the field numbers are delta encoded for efficiency and the fields are always sent in order of increasing field number; the deltas are therefore unsigned. The initialization for the delta encoding sets the field number to -1, so an unsigned integer field 0 with value 7 is transmitted as unsigned delta = 1, unsigned value = 7 or (01 07). Finally, after all the fields have been sent a terminating mark denotes the end of the struct. That mark is a delta=0 value, which has representation (00).

Interface types are not checked for compatibility; all interface types are treated, for transmission, as members of a single "interface" type, analogous to int or []byte - in effect they're all treated as interface{}. Interface values are transmitted as a string identifying the concrete type being sent (a name that must be pre-defined by calling Register), followed by a byte count of the length of the following data (so the value can be skipped if it cannot be stored), followed by the usual encoding of concrete (dynamic) value stored in the interface value. (A nil interface value is identified by the empty string and transmits no value.) Upon receipt, the decoder verifies that the unpacked concrete item satisfies the interface of the receiving variable.

If a value is passed to Encode and the type is not a struct (or pointer to struct, etc.), for simplicity of processing it is represented as a struct of one field. The only visible effect of this is to encode a zero byte after the value, just as after the last field of an encoded struct, so that the decode algorithm knows when the top-level value is complete.

The representation of types is described below. When a type is defined on a given connection between an Encoder and Decoder, it is assigned a signed integer type id. When Encoder.Encode(v) is called, it makes sure there is an id assigned for the type of v and all its elements and then it sends the pair (typeid, encoded-v) where typeid is the type id of the encoded type of v and encoded-v is the gob encoding of the value v.

To define a type, the encoder chooses an unused, positive type id and sends the pair (-type id, encoded-type) where encoded-type is the gob encoding of a wireType description, constructed from these types:

``` go 
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

For simplicity in setup, the connection is defined to understand these types a priori, as well as the basic gob types int, uint, etc. Their ids are:

```
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

In summary, a gob stream looks like

```
(byteCount (-type id, encoding of a wireType)* (type id, encoding of a value))*
```

where * signifies zero or more repetitions and the type id of a value must be predefined or be defined before the value in the stream.

Compatibility: Any future changes to the package will endeavor to maintain compatibility with streams encoded using previous versions. That is, any released version of this package should be able to decode data written with any previously released version, subject to issues such as security fixes. See the Go compatibility document for background: https://golang.org/doc/go1compat

See "Gobs of data" for a design discussion of the gob wire format: https://blog.golang.org/gobs-of-data

#### Security 

This package is not designed to be hardened against adversarial inputs, and is outside the scope of https://go.dev/security/policy. In particular, the Decoder does only basic sanity checking on decoded input sizes, and its limits are not configurable. Care should be taken when decoding gob data from untrusted sources, which may consume significant resources.

##### Example (Basic)

This example shows the basic usage of the package: Create an encoder, transmit some values, receive them with a decoder.

``` go 
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

##### Example(EncodeDecode)

This example transmits a value that implements the custom encoding and decoding methods.

``` go 
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

##### Example (Interface)

This example shows how to encode an interface value. The key distinction from regular types is to register the concrete type that implements the interface.

``` go 
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

#### func Register 

``` go 
func Register(value any)
```

Register records a type, identified by a value for that type, under its internal type name. That name will identify the concrete type of a value sent or received as an interface variable. Only types that will be transferred as implementations of interface values need to be registered. Expecting to be used only during initialization, it panics if the mapping between types and names is not a bijection.

#### func RegisterName 

``` go 
func RegisterName(name string, value any)
```

RegisterName is like Register but uses the provided name rather than the type's default.

## 类型

### type CommonType 

``` go 
type CommonType struct {
	Name string
	Id   typeId
}
```

CommonType holds elements of all types. It is a historical artifact, kept for binary compatibility and exported only for the benefit of the package's encoding of type descriptors. It is not intended for direct use by clients.

### type Decoder 

``` go 
type Decoder struct {
	// contains filtered or unexported fields
}
```

A Decoder manages the receipt of type and data information read from the remote side of a connection. It is safe for concurrent use by multiple goroutines.

The Decoder does only basic sanity checking on decoded input sizes, and its limits are not configurable. Take caution when decoding gob data from untrusted sources.

#### func NewDecoder 

``` go 
func NewDecoder(r io.Reader) *Decoder
```

NewDecoder returns a new decoder that reads from the io.Reader. If r does not also implement io.ByteReader, it will be wrapped in a bufio.Reader.

#### (*Decoder) Decode 

``` go 
func (dec *Decoder) Decode(e any) error
```

Decode reads the next value from the input stream and stores it in the data represented by the empty interface value. If e is nil, the value will be discarded. Otherwise, the value underlying e must be a pointer to the correct type for the next data item received. If the input is at EOF, Decode returns io.EOF and does not modify e.

#### (*Decoder) DecodeValue 

``` go 
func (dec *Decoder) DecodeValue(v reflect.Value) error
```

DecodeValue reads the next value from the input stream. If v is the zero reflect.Value (v.Kind() == Invalid), DecodeValue discards the value. Otherwise, it stores the value into v. In that case, v must represent a non-nil pointer to data or be an assignable reflect.Value (v.CanSet()) If the input is at EOF, DecodeValue returns io.EOF and does not modify v.

### type Encoder 

``` go 
type Encoder struct {
	// contains filtered or unexported fields
}
```

An Encoder manages the transmission of type and data information to the other side of a connection. It is safe for concurrent use by multiple goroutines.

#### func NewEncoder 

``` go 
func NewEncoder(w io.Writer) *Encoder
```

NewEncoder returns a new encoder that will transmit on the io.Writer.

#### (*Encoder) Encode 

``` go 
func (enc *Encoder) Encode(e any) error
```

Encode transmits the data item represented by the empty interface value, guaranteeing that all necessary type information has been transmitted first. Passing a nil pointer to Encoder will panic, as they cannot be transmitted by gob.

#### (*Encoder) EncodeValue 

``` go 
func (enc *Encoder) EncodeValue(value reflect.Value) error
```

EncodeValue transmits the data item represented by the reflection value, guaranteeing that all necessary type information has been transmitted first. Passing a nil pointer to EncodeValue will panic, as they cannot be transmitted by gob.

### type GobDecoder 

``` go 
type GobDecoder interface {
	// GobDecode overwrites the receiver, which must be a pointer,
	// with the value represented by the byte slice, which was written
	// by GobEncode, usually for the same concrete type.
	GobDecode([]byte) error
}
```

GobDecoder is the interface describing data that provides its own routine for decoding transmitted values sent by a GobEncoder.

### type GobEncoder 

``` go 
type GobEncoder interface {
	// GobEncode returns a byte slice representing the encoding of the
	// receiver for transmission to a GobDecoder, usually of the same
	// concrete type.
	GobEncode() ([]byte, error)
}
```

GobEncoder is the interface describing data that provides its own representation for encoding values for transmission to a GobDecoder. A type that implements GobEncoder and GobDecoder has complete control over the representation of its data and may therefore contain things such as private fields, channels, and functions, which are not usually transmissible in gob streams.

Note: Since gobs can be stored permanently, it is good design to guarantee the encoding used by a GobEncoder is stable as the software evolves. For instance, it might make sense for GobEncode to include a version number in the encoding.