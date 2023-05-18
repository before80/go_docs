+++
title = "constant"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# constant

https://pkg.go.dev/go/constant@go1.20.1



Package constant implements Values representing untyped Go constants and their corresponding operations.

A special Unknown value may be used when a value is unknown due to an error. Operations on unknown values produce unknown values unless specified otherwise.

##### Example
``` go 
```












## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [BitLen](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=645) 

``` go 
func BitLen(x Value) int
```

BitLen returns the number of bits required to represent the absolute value x in binary representation; x must be an Int or an Unknown. If x is Unknown, the result is 0.

#### func [BoolVal](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=480) 

``` go 
func BoolVal(x Value) bool
```

BoolVal returns the Go boolean value of x, which must be a Bool or an Unknown. If x is Unknown, the result is false.

#### func [Bytes](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=702) 

``` go 
func Bytes(x Value) []byte
```

Bytes returns the bytes for the absolute value of x in little- endian binary representation; x must be an Int.

#### func [Compare](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=1337) 

``` go 
func Compare(x_ Value, op token.Token, y_ Value) bool
```

Compare returns the result of the comparison x op y. The comparison must be defined for the operands. If one of the operands is Unknown, the result is false.

##### Example
``` go 
```

#### func [Float32Val](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=537) 

``` go 
func Float32Val(x Value) (float32, bool)
```

Float32Val is like Float64Val but for float32 instead of float64.

#### func [Float64Val](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=562) 

``` go 
func Float64Val(x Value) (float64, bool)
```

Float64Val returns the nearest Go float64 value of x and whether the result is exact; x must be numeric or an Unknown, but not Complex. For values too small (too close to 0) to represent as float64, Float64Val silently underflows to 0. The result sign always matches the sign of x, even for 0. If x is Unknown, the result is (0, false).

#### func [Int64Val](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=507) 

``` go 
func Int64Val(x Value) (int64, bool)
```

Int64Val returns the Go int64 value of x and whether the result is exact; x must be an Int or an Unknown. If the result is not exact, its value is undefined. If x is Unknown, the result is (0, false).

#### func [Sign](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=665) 

``` go 
func Sign(x Value) int
```

Sign returns -1, 0, or 1 depending on whether x < 0, x == 0, or x > 0; x must be numeric or Unknown. For complex values x, the sign is 0 if x == 0, otherwise it is != 0. If x is Unknown, the result is 1.

##### Example
``` go 
```

#### func [StringVal](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=493) 

``` go 
func StringVal(x Value) string
```

StringVal returns the Go string value of x, which must be a String or an Unknown. If x is Unknown, the result is "".

#### func [Uint64Val](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=523) 

``` go 
func Uint64Val(x Value) (uint64, bool)
```

Uint64Val returns the Go uint64 value of x and whether the result is exact; x must be an Int or an Unknown. If the result is not exact, its value is undefined. If x is Unknown, the result is (0, false).

#### func [Val](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=593)  <- go1.13

``` go 
func Val(x Value) any
```

Val returns the underlying value for a given constant. Since it returns an interface, it is up to the caller to type assert the result to the expected type. The possible dynamic return types are:

```
x Kind             type of result
-----------------------------------------
Bool               bool
String             string
Int                int64 or *big.Int
Float              *big.Float or *big.Rat
everything else    nil
```

##### Example
``` go 
```

## 类型

### type [Kind](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=29) 

``` go 
type Kind int
```

Kind specifies the kind of value represented by a Value.

``` go 
const (
	// unknown values
	Unknown Kind = iota

	// non-numeric values
	Bool
	String

	// numeric values
	Int
	Float
	Complex
)
```

#### (Kind) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/kind_string.go;l=23)  <- go1.18

``` go 
func (i Kind) String() string
```

### type [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=46) 

``` go 
type Value interface {
	// Kind returns the value kind.
	Kind() Kind

	// String returns a short, quoted (human-readable) form of the value.
	// For numeric values, the result may be an approximation;
	// for String values the result may be a shortened string.
	// Use ExactString for a string representing a value exactly.
	String() string

	// ExactString returns an exact, quoted (human-readable) form of the value.
	// If the Value is of Kind String, use StringVal to obtain the unquoted string.
	ExactString() string
	// contains filtered or unexported methods
}
```

A Value represents the value of a Go constant.

#### func [BinaryOp](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=1103) 

``` go 
func BinaryOp(x_ Value, op token.Token, y_ Value) Value
```

BinaryOp returns the result of the binary expression x op y. The operation must be defined for the operands. If one of the operands is Unknown, the result is Unknown. BinaryOp doesn't handle comparisons or shifts; use Compare or Shift instead.

To force integer division of Int operands, use op == token.QUO_ASSIGN instead of token.QUO; the result is guaranteed to be Int in this case. Division by zero leads to a run-time panic.

##### Example
``` go 
```

#### func [Denom](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=788) 

``` go 
func Denom(x Value) Value
```

Denom returns the denominator of x; x must be Int, Float, or Unknown. If x is Unknown, or if it is too large or small to represent as a fraction, the result is Unknown. Otherwise the result is an Int >= 1.

#### func [Imag](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=836) 

``` go 
func Imag(x Value) Value
```

Imag returns the imaginary part of x, which must be a numeric or unknown value. If x is Unknown, the result is Unknown.

#### func [Make](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=623)  <- go1.13

``` go 
func Make(x any) Value
```

Make returns the Value for x.

``` go 
type of x        result Kind
----------------------------
bool             Bool
string           String
int64            Int
*big.Int         Int
*big.Float       Float
*big.Rat         Float
anything else    Unknown
```

#### func [MakeBool](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=386) 

``` go 
func MakeBool(b bool) Value
```

MakeBool returns the Bool value for b.

#### func [MakeFloat64](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=412) 

``` go 
func MakeFloat64(x float64) Value
```

MakeFloat64 returns the Float value for x. If x is -0.0, the result is 0.0. If x is not finite, the result is an Unknown.

#### func [MakeFromBytes](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=734) 

``` go 
func MakeFromBytes(bytes []byte) Value
```

MakeFromBytes returns the Int value given the bytes of its little-endian binary representation. An empty byte slice argument represents 0.

#### func [MakeFromLiteral](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=427) 

``` go 
func MakeFromLiteral(lit string, tok token.Token, zero uint) Value
```

MakeFromLiteral returns the corresponding integer, floating-point, imaginary, character, or string value for a Go literal string. The tok value must be one of token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING. The final argument must be zero. If the literal string syntax is invalid, the result is an Unknown.

#### func [MakeImag](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=810) 

``` go 
func MakeImag(x Value) Value
```

MakeImag returns the Complex value x*i; x must be Int, Float, or Unknown. If x is Unknown, the result is Unknown.

#### func [MakeInt64](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=399) 

``` go 
func MakeInt64(x int64) Value
```

MakeInt64 returns the Int value for x.

#### func [MakeString](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=389) 

``` go 
func MakeString(s string) Value
```

MakeString returns the String value for s.

#### func [MakeUint64](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=402) 

``` go 
func MakeUint64(x uint64) Value
```

MakeUint64 returns the Int value for x.

#### func [MakeUnknown](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=383) 

``` go 
func MakeUnknown() Value
```

MakeUnknown returns the Unknown value.

#### func [Num](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=766) 

``` go 
func Num(x Value) Value
```

Num returns the numerator of x; x must be Int, Float, or Unknown. If x is Unknown, or if it is too large or small to represent as a fraction, the result is Unknown. Otherwise the result is an Int with the same sign as x.

#### func [Real](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=823) 

``` go 
func Real(x Value) Value
```

Real returns the real part of x, which must be a numeric or unknown value. If x is Unknown, the result is Unknown.

#### func [Shift](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=1282) 

``` go 
func Shift(x Value, op token.Token, s uint) Value
```

Shift returns the result of the shift expression x op s with op == token.SHL or token.SHR (<< or >>). x must be an Int or an Unknown. If x is Unknown, the result is x.

#### func [ToComplex](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=929)  <- go1.6

``` go 
func ToComplex(x Value) Value
```

ToComplex converts x to a Complex value if x is representable as a Complex. Otherwise it returns an Unknown.

#### func [ToFloat](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=908)  <- go1.6

``` go 
func ToFloat(x Value) Value
```

ToFloat converts x to a Float value if x is representable as a Float. Otherwise it returns an Unknown.

#### func [ToInt](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=854)  <- go1.6

``` go 
func ToInt(x Value) Value
```

ToInt converts x to an Int value if x is representable as an Int. Otherwise it returns an Unknown.

#### func [UnaryOp](https://cs.opensource.google/go/go/+/go1.20.1:src/go/constant/value.go;l=958) 

``` go 
func UnaryOp(op token.Token, y Value, prec uint) Value
```

UnaryOp returns the result of the unary expression op y. The operation must be defined for the operand. If prec > 0 it specifies the ^ (xor) result size in bits. If y is Unknown, the result is Unknown.