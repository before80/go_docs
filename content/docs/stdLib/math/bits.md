+++
title = "bits"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# bits

https://pkg.go.dev/math/bits@go1.20.1



Package bits implements bit counting and manipulation functions for the predeclared unsigned integer types.

Functions in this package may be implemented directly by the compiler, for better performance. For those functions the code in this package will not be used. Which functions are implemented by the compiler depends on the architecture and the Go release.










































## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=20)

``` go linenums="1"
const UintSize = uintSize
```

UintSize is the size of a uint in bits.

## 变量

This section is empty.

## 函数

#### func [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=360)  <- go1.12

``` go linenums="1"
func Add(x, y, carry uint) (sum, carryOut uint)
```

Add returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.

#### func [Add32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=374)  <- go1.12

``` go linenums="1"
func Add32(x, y, carry uint32) (sum, carryOut uint32)
```

Add32 returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [Add64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=386)  <- go1.12

``` go linenums="1"
func Add64(x, y, carry uint64) (sum, carryOut uint64)
```

Add64 returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [Div](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=492)  <- go1.12

``` go linenums="1"
func Div(hi, lo, y uint) (quo, rem uint)
```

Div returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper half in parameter hi and the lower half in parameter lo. Div panics for y == 0 (division by zero) or y <= hi (quotient overflow).

#### func [Div32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=505)  <- go1.12

``` go linenums="1"
func Div32(hi, lo, y uint32) (quo, rem uint32)
```

Div32 returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper half in parameter hi and the lower half in parameter lo. Div32 panics for y == 0 (division by zero) or y <= hi (quotient overflow).

##### Example
``` go linenums="1"
```

#### func [Div64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=518)  <- go1.12

``` go linenums="1"
func Div64(hi, lo, y uint64) (quo, rem uint64)
```

Div64 returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper half in parameter hi and the lower half in parameter lo. Div64 panics for y == 0 (division by zero) or y <= hi (quotient overflow).

##### Example
``` go linenums="1"
```

#### func [LeadingZeros](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=25) 

``` go linenums="1"
func LeadingZeros(x uint) int
```

LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.

#### func [LeadingZeros16](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=31) 

``` go linenums="1"
func LeadingZeros16(x uint16) int
```

LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.

##### Example
``` go linenums="1"
```

#### func [LeadingZeros32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=34) 

``` go linenums="1"
func LeadingZeros32(x uint32) int
```

LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.

##### Example
``` go linenums="1"
```

#### func [LeadingZeros64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=37) 

``` go linenums="1"
func LeadingZeros64(x uint64) int
```

LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.

##### Example
``` go linenums="1"
```

#### func [LeadingZeros8](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=28) 

``` go linenums="1"
func LeadingZeros8(x uint8) int
```

LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.

##### Example
``` go linenums="1"
```

#### func [Len](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=302) 

``` go linenums="1"
func Len(x uint) int
```

Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.

#### func [Len16](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=315) 

``` go linenums="1"
func Len16(x uint16) (n int)
```

Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0.

##### Example
``` go linenums="1"
```

#### func [Len32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=324) 

``` go linenums="1"
func Len32(x uint32) (n int)
```

Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.

##### Example
``` go linenums="1"
```

#### func [Len64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=337) 

``` go linenums="1"
func Len64(x uint64) (n int)
```

Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.

##### Example
``` go linenums="1"
```

#### func [Len8](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=310) 

``` go linenums="1"
func Len8(x uint8) int
```

Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.

##### Example
``` go linenums="1"
```

#### func [Mul](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=445)  <- go1.12

``` go linenums="1"
func Mul(x, y uint) (hi, lo uint)
```

Mul returns the full-width product of x and y: (hi, lo) = x * y with the product bits' upper half returned in hi and the lower half returned in lo.

This function's execution time does not depend on the inputs.

#### func [Mul32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=459)  <- go1.12

``` go linenums="1"
func Mul32(x, y uint32) (hi, lo uint32)
```

Mul32 returns the 64-bit product of x and y: (hi, lo) = x * y with the product bits' upper half returned in hi and the lower half returned in lo.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [Mul64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=470)  <- go1.12

``` go linenums="1"
func Mul64(x, y uint64) (hi, lo uint64)
```

Mul64 returns the 128-bit product of x and y: (hi, lo) = x * y with the product bits' upper half returned in hi and the lower half returned in lo.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [OnesCount](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=117) 

``` go linenums="1"
func OnesCount(x uint) int
```

OnesCount returns the number of one bits ("population count") in x.

##### Example
``` go linenums="1"
```

#### func [OnesCount16](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=130) 

``` go linenums="1"
func OnesCount16(x uint16) int
```

OnesCount16 returns the number of one bits ("population count") in x.

##### Example
``` go linenums="1"
```

#### func [OnesCount32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=135) 

``` go linenums="1"
func OnesCount32(x uint32) int
```

OnesCount32 returns the number of one bits ("population count") in x.

##### Example
``` go linenums="1"
```

#### func [OnesCount64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=140) 

``` go linenums="1"
func OnesCount64(x uint64) int
```

OnesCount64 returns the number of one bits ("population count") in x.

##### Example
``` go linenums="1"
```

#### func [OnesCount8](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=125) 

``` go linenums="1"
func OnesCount8(x uint8) int
```

OnesCount8 returns the number of one bits ("population count") in x.

##### Example
``` go linenums="1"
```

#### func [Rem](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=573)  <- go1.14

``` go linenums="1"
func Rem(hi, lo, y uint) uint
```

Rem returns the remainder of (hi, lo) divided by y. Rem panics for y == 0 (division by zero) but, unlike Div, it doesn't panic on a quotient overflow.

#### func [Rem32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=583)  <- go1.14

``` go linenums="1"
func Rem32(hi, lo, y uint32) uint32
```

Rem32 returns the remainder of (hi, lo) divided by y. Rem32 panics for y == 0 (division by zero) but, unlike Div32, it doesn't panic on a quotient overflow.

#### func [Rem64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=590)  <- go1.14

``` go linenums="1"
func Rem64(hi, lo, y uint64) uint64
```

Rem64 returns the remainder of (hi, lo) divided by y. Rem64 panics for y == 0 (division by zero) but, unlike Div64, it doesn't panic on a quotient overflow.

#### func [Reverse](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=226) 

``` go linenums="1"
func Reverse(x uint) uint
```

Reverse returns the value of x with its bits in reversed order.

#### func [Reverse16](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=239) 

``` go linenums="1"
func Reverse16(x uint16) uint16
```

Reverse16 returns the value of x with its bits in reversed order.

##### Example
``` go linenums="1"
```

#### func [Reverse32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=244) 

``` go linenums="1"
func Reverse32(x uint32) uint32
```

Reverse32 returns the value of x with its bits in reversed order.

##### Example
``` go linenums="1"
```

#### func [Reverse64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=253) 

``` go linenums="1"
func Reverse64(x uint64) uint64
```

Reverse64 returns the value of x with its bits in reversed order.

##### Example
``` go linenums="1"
```

#### func [Reverse8](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=234) 

``` go linenums="1"
func Reverse8(x uint8) uint8
```

Reverse8 returns the value of x with its bits in reversed order.

##### Example
``` go linenums="1"
```

#### func [ReverseBytes](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=266) 

``` go linenums="1"
func ReverseBytes(x uint) uint
```

ReverseBytes returns the value of x with its bytes in reversed order.

This function's execution time does not depend on the inputs.

#### func [ReverseBytes16](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=276) 

``` go linenums="1"
func ReverseBytes16(x uint16) uint16
```

ReverseBytes16 returns the value of x with its bytes in reversed order.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [ReverseBytes32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=283) 

``` go linenums="1"
func ReverseBytes32(x uint32) uint32
```

ReverseBytes32 returns the value of x with its bytes in reversed order.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [ReverseBytes64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=292) 

``` go linenums="1"
func ReverseBytes64(x uint64) uint64
```

ReverseBytes64 returns the value of x with its bytes in reversed order.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [RotateLeft](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=176) 

``` go linenums="1"
func RotateLeft(x uint, k int) uint
```

RotateLeft returns the value of x rotated left by (k mod UintSize) bits. To rotate x right by k bits, call RotateLeft(x, -k).

This function's execution time does not depend on the inputs.

#### func [RotateLeft16](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=197) 

``` go linenums="1"
func RotateLeft16(x uint16, k int) uint16
```

RotateLeft16 returns the value of x rotated left by (k mod 16) bits. To rotate x right by k bits, call RotateLeft16(x, -k).

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [RotateLeft32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=207) 

``` go linenums="1"
func RotateLeft32(x uint32, k int) uint32
```

RotateLeft32 returns the value of x rotated left by (k mod 32) bits. To rotate x right by k bits, call RotateLeft32(x, -k).

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [RotateLeft64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=217) 

``` go linenums="1"
func RotateLeft64(x uint64, k int) uint64
```

RotateLeft64 returns the value of x rotated left by (k mod 64) bits. To rotate x right by k bits, call RotateLeft64(x, -k).

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [RotateLeft8](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=187) 

``` go linenums="1"
func RotateLeft8(x uint8, k int) uint8
```

RotateLeft8 returns the value of x rotated left by (k mod 8) bits. To rotate x right by k bits, call RotateLeft8(x, -k).

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [Sub](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=402)  <- go1.12

``` go linenums="1"
func Sub(x, y, borrow uint) (diff, borrowOut uint)
```

Sub returns the difference of x, y and borrow: diff = x - y - borrow. The borrow input must be 0 or 1; otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.

#### func [Sub32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=416)  <- go1.12

``` go linenums="1"
func Sub32(x, y, borrow uint32) (diff, borrowOut uint32)
```

Sub32 returns the difference of x, y and borrow, diff = x - y - borrow. The borrow input must be 0 or 1; otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [Sub64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=431)  <- go1.12

``` go linenums="1"
func Sub64(x, y, borrow uint64) (diff, borrowOut uint64)
```

Sub64 returns the difference of x, y and borrow: diff = x - y - borrow. The borrow input must be 0 or 1; otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.

This function's execution time does not depend on the inputs.

##### Example
``` go linenums="1"
```

#### func [TrailingZeros](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=59) 

``` go linenums="1"
func TrailingZeros(x uint) int
```

TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.

#### func [TrailingZeros16](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=72) 

``` go linenums="1"
func TrailingZeros16(x uint16) int
```

TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.

##### Example
``` go linenums="1"
```

#### func [TrailingZeros32](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=81) 

``` go linenums="1"
func TrailingZeros32(x uint32) int
```

TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.

##### Example
``` go linenums="1"
```

#### func [TrailingZeros64](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=90) 

``` go linenums="1"
func TrailingZeros64(x uint64) int
```

TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.

##### Example
``` go linenums="1"
```

#### func [TrailingZeros8](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=67) 

``` go linenums="1"
func TrailingZeros8(x uint8) int
```

TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.

##### Example
``` go linenums="1"
```

## 类型

This section is empty.