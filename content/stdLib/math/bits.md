+++
title = "bits"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/math/bits@go1.23.0](https://pkg.go.dev/math/bits@go1.23.0)

Package bits implements bit counting and manipulation functions for the predeclared unsigned integer types.

​	bits 包为预声明的无符号整数类型实现位计数和操作函数。

Functions in this package may be implemented directly by the compiler, for better performance. For those functions the code in this package will not be used. Which functions are implemented by the compiler depends on the architecture and the Go release.

​	此包中的函数可能由编译器直接实现，以获得更好的性能。对于这些函数，将不会使用此包中的代码。由编译器实现哪些函数取决于体系结构和 Go 版本。

## 常量

[View Source
查看源代码](https://cs.opensource.google/go/go/+/go1.20.1:src/math/bits/bits.go;l=20)

```go
const UintSize = uintSize
```

UintSize is the size of a uint in bits.

​	UintSize 是以位为单位的 uint 的大小。

## 变量

This section is empty.

## 函数

### func Add <- go1.12

```go
func Add(x, y, carry uint) (sum, carryOut uint)
```

Add returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1.

​	Add 返回 x、y 和进位之和：sum = x + y + 进位。进位输入必须为 0 或 1；否则行为是未定义的。carryOut 输出保证为 0 或 1。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

### func Add32 <- go1.12

```go
func Add32(x, y, carry uint32) (sum, carryOut uint32)
```

Add32 returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1.

​	Add32 返回 x、y 和进位之和：sum = x + y + 进位。进位输入必须为 0 或 1；否则行为是未定义的。carryOut 输出保证为 0 或 1。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### Add32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 33<<32 + 12
	n1 := []uint32{33, 12}
	// Second number is 21<<32 + 23
	n2 := []uint32{21, 23}
	// Add them together without producing carry.
	d1, carry := bits.Add32(n1[1], n2[1], 0)
	d0, _ := bits.Add32(n1[0], n2[0], carry)
	nsum := []uint32{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// First number is 1<<32 + 2147483648
	n1 = []uint32{1, 0x80000000}
	// Second number is 1<<32 + 2147483648
	n2 = []uint32{1, 0x80000000}
	// Add them together producing carry.
	d1, carry = bits.Add32(n1[1], n2[1], 0)
	d0, _ = bits.Add32(n1[0], n2[0], carry)
	nsum = []uint32{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
}
Output:

[33 12] + [21 23] = [54 35] (carry bit was 0)
[1 2147483648] + [1 2147483648] = [3 0] (carry bit was 1)
```

### func Add64 <- go1.12 

```go
func Add64(x, y, carry uint64) (sum, carryOut uint64)
```

Add64 returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1.

​	Add64 返回 x、y 和进位之和：sum = x + y + 进位。进位输入必须为 0 或 1；否则行为是未定义的。carryOut 输出保证为 0 或 1。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### Add64 Example 

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 33<<64 + 12
	n1 := []uint64{33, 12}
	// Second number is 21<<64 + 23
	n2 := []uint64{21, 23}
	// Add them together without producing carry.
	d1, carry := bits.Add64(n1[1], n2[1], 0)
	d0, _ := bits.Add64(n1[0], n2[0], carry)
	nsum := []uint64{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// First number is 1<<64 + 9223372036854775808
	n1 = []uint64{1, 0x8000000000000000}
	// Second number is 1<<64 + 9223372036854775808
	n2 = []uint64{1, 0x8000000000000000}
	// Add them together producing carry.
	d1, carry = bits.Add64(n1[1], n2[1], 0)
	d0, _ = bits.Add64(n1[0], n2[0], carry)
	nsum = []uint64{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
}
```

### func Div <- go1.12

```go
func Div(hi, lo, y uint) (quo, rem uint)
```

Div returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits’ upper half in parameter hi and the lower half in parameter lo. Div panics for y == 0 (division by zero) or y <= hi (quotient overflow).

​	Div 返回 (hi, lo)/y 的商和余数：quo = (hi, lo)/y，rem = (hi, lo)%y，其中被除数的二进制位高半部分在参数 hi 中，低半部分在参数 lo 中。如果 y == 0（除以零）或 y <= hi（商溢出），Div 会引发 panic。

### func Div32 <- go1.12

```go
func Div32(hi, lo, y uint32) (quo, rem uint32)
```

Div32 returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits’ upper half in parameter hi and the lower half in parameter lo. Div32 panics for y == 0 (division by zero) or y <= hi (quotient overflow).

​	Div32 返回 (hi, lo)/y 的商和余数：quo = (hi, lo)/y，rem = (hi, lo)%y，其中被除数的二进制位高半部分在参数 hi 中，低半部分在参数 lo 中。如果 y == 0（除以零）或 y <= hi（商溢出），Div32 会引发 panic。

#### Div32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 0<<32 + 6
	n1 := []uint32{0, 6}
	// Second number is 0<<32 + 3
	n2 := []uint32{0, 3}
	// Divide them together.
	quo, rem := bits.Div32(n1[0], n1[1], n2[1])
	nsum := []uint32{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)

	// First number is 2<<32 + 2147483648
	n1 = []uint32{2, 0x80000000}
	// Second number is 0<<32 + 2147483648
	n2 = []uint32{0, 0x80000000}
	// Divide them together.
	quo, rem = bits.Div32(n1[0], n1[1], n2[1])
	nsum = []uint32{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)
}
Output:

[0 6] / 3 = [2 0]
[2 2147483648] / 2147483648 = [5 0]
```

### func Div64 <- go1.12

```go
func Div64(hi, lo, y uint64) (quo, rem uint64)
```

Div64 returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits’ upper half in parameter hi and the lower half in parameter lo. Div64 panics for y == 0 (division by zero) or y <= hi (quotient overflow).

​	Div64 返回 (hi, lo)/y 的商和余数：quo = (hi, lo)/y，rem = (hi, lo)%y，其中被除数的二进制位高半部分在参数 hi 中，低半部分在参数 lo 中。如果 y == 0（除以零）或 y <= hi（商溢出），Div64 会引发 panic。

#### Div64 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 0<<64 + 6
	n1 := []uint64{0, 6}
	// Second number is 0<<64 + 3
	n2 := []uint64{0, 3}
	// Divide them together.
	quo, rem := bits.Div64(n1[0], n1[1], n2[1])
	nsum := []uint64{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)

	// First number is 2<<64 + 9223372036854775808
	n1 = []uint64{2, 0x8000000000000000}
	// Second number is 0<<64 + 9223372036854775808
	n2 = []uint64{0, 0x8000000000000000}
	// Divide them together.
	quo, rem = bits.Div64(n1[0], n1[1], n2[1])
	nsum = []uint64{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)
}
Output:

[0 6] / 3 = [2 0]
[2 9223372036854775808] / 9223372036854775808 = [5 0]
```

### func LeadingZeros

```go
func LeadingZeros(x uint) int
```

LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.

​	LeadingZeros 返回 x 中前导零位的数量；对于 x == 0，结果为 UintSize。

### func LeadingZeros16

```go
func LeadingZeros16(x uint16) int
```

LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.

​	LeadingZeros16 返回 x 中前导零位的数量；对于 x == 0，结果为 16。

#### LeadingZeros16 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("LeadingZeros16(%016b) = %d\n", 1, bits.LeadingZeros16(1))
}
Output:

LeadingZeros16(0000000000000001) = 15
```

### func LeadingZeros32

```go
func LeadingZeros32(x uint32) int
```

LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.

​	LeadingZeros32 返回 x 中前导零位的数量；对于 x == 0，结果为 32。

#### LeadingZeros32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("LeadingZeros32(%032b) = %d\n", 1, bits.LeadingZeros32(1))
}
Output:

LeadingZeros32(00000000000000000000000000000001) = 31
```

### func LeadingZeros64

```go
func LeadingZeros64(x uint64) int
```

LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.

​	LeadingZeros64 返回 x 中前导零位的数量；对于 x == 0，结果为 64。

#### LeadingZeros64 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1, bits.LeadingZeros64(1))
}
Output:

LeadingZeros64(0000000000000000000000000000000000000000000000000000000000000001) = 63
```

### func LeadingZeros8

```go
func LeadingZeros8(x uint8) int
```

LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.

​	LeadingZeros8 返回 x 中前导零位的数量；对于 x == 0，结果为 8。

#### LeadingZeros8 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("LeadingZeros8(%08b) = %d\n", 1, bits.LeadingZeros8(1))
}
Output:

LeadingZeros8(00000001) = 7
```

### func Len

```go
func Len(x uint) int
```

Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.

​	Len 返回表示 x 所需的最小位数；对于 x == 0，结果为 0。

### func Len16

```go
func Len16(x uint16) (n int)
```

Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0.

​	Len16 返回表示 x 所需的最小位数；对于 x == 0，结果为 0。

#### Len16 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("Len16(%016b) = %d\n", 8, bits.Len16(8))
}
Output:

Len16(0000000000001000) = 4
```

### func Len32

```go
func Len32(x uint32) (n int)
```

Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.

​	Len32 返回表示 x 所需的最小位数；对于 x == 0，结果为 0。

#### Len32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("Len32(%032b) = %d\n", 8, bits.Len32(8))
}
Output:

Len32(00000000000000000000000000001000) = 4
```

### func Len64

```go
func Len64(x uint64) (n int)
```

Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.

​	Len64 返回表示 x 所需的最小位数；对于 x == 0，结果为 0。

#### Len64 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("Len64(%064b) = %d\n", 8, bits.Len64(8))
}
Output:

Len64(0000000000000000000000000000000000000000000000000000000000001000) = 4
```

### func Len8

```go
func Len8(x uint8) int
```

Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.

​	Len8 返回表示 x 所需的最小位数；对于 x == 0，结果为 0。

#### Len8 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("Len8(%08b) = %d\n", 8, bits.Len8(8))
}
Output:

Len8(00001000) = 4
```

### func Mul <- go1.12

```go
func Mul(x, y uint) (hi, lo uint)
```

Mul returns the full-width product of x and y: (hi, lo) = x * y with the product bits’ upper half returned in hi and the lower half returned in lo.

​	Mul 返回 x 和 y 的全宽乘积：(hi, lo) = x * y，其中乘积位的高半部分在 hi 中返回，低半部分在 lo 中返回。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

### func Mul32 <- go1.12

```go
func Mul32(x, y uint32) (hi, lo uint32)
```

Mul32 returns the 64-bit product of x and y: (hi, lo) = x * y with the product bits’ upper half returned in hi and the lower half returned in lo.

​	Mul32 返回 x 和 y 的 64 位乘积：(hi, lo) = x * y，其中乘积位的高半部分在 hi 中返回，低半部分在 lo 中返回。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### Mul32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 0<<32 + 12
	n1 := []uint32{0, 12}
	// Second number is 0<<32 + 12
	n2 := []uint32{0, 12}
	// Multiply them together without producing overflow.
	hi, lo := bits.Mul32(n1[1], n2[1])
	nsum := []uint32{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)

	// First number is 0<<32 + 2147483648
	n1 = []uint32{0, 0x80000000}
	// Second number is 0<<32 + 2
	n2 = []uint32{0, 2}
	// Multiply them together producing overflow.
	hi, lo = bits.Mul32(n1[1], n2[1])
	nsum = []uint32{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)
}
Output:

12 * 12 = [0 144]
2147483648 * 2 = [1 0]
```

### func Mul64 <- go1.12

```go
func Mul64(x, y uint64) (hi, lo uint64)
```

Mul64 returns the 128-bit product of x and y: (hi, lo) = x * y with the product bits’ upper half returned in hi and the lower half returned in lo.

​	Mul64 返回 x 和 y 的 128 位乘积：(hi, lo) = x * y，其中乘积位的高半部分在 hi 中返回，低半部分在 lo 中返回。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### Mul64 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 0<<64 + 12
	n1 := []uint64{0, 12}
	// Second number is 0<<64 + 12
	n2 := []uint64{0, 12}
	// Multiply them together without producing overflow.
	hi, lo := bits.Mul64(n1[1], n2[1])
	nsum := []uint64{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)

	// First number is 0<<64 + 9223372036854775808
	n1 = []uint64{0, 0x8000000000000000}
	// Second number is 0<<64 + 2
	n2 = []uint64{0, 2}
	// Multiply them together producing overflow.
	hi, lo = bits.Mul64(n1[1], n2[1])
	nsum = []uint64{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)
}
Output:

12 * 12 = [0 144]
9223372036854775808 * 2 = [1 0]
```

### func OnesCount

```go
func OnesCount(x uint) int
```

OnesCount returns the number of one bits (“population count”) in x.

​	OnesCount 返回 x 中的 1 位数（“基数计数”）。

#### OnesCount Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount(%b) = %d\n", 14, bits.OnesCount(14))
}
Output:

OnesCount(1110) = 3
```

### func OnesCount16

```go
func OnesCount16(x uint16) int
```

OnesCount16 returns the number of one bits (“population count”) in x.

​	OnesCount16 返回 x 中的 1 位数（“基数计数”）。

#### OnesCount16 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount16(%016b) = %d\n", 14, bits.OnesCount16(14))
}
Output:

OnesCount16(0000000000001110) = 3
```

### func OnesCount32

```go
func OnesCount32(x uint32) int
```

OnesCount32 returns the number of one bits (“population count”) in x.

​	OnesCount32 返回 x 中的 1 位数（“基数计数”）。

#### OnesCount32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount32(%032b) = %d\n", 14, bits.OnesCount32(14))
}
Output:

OnesCount32(00000000000000000000000000001110) = 3
```

### func OnesCount64

```go
func OnesCount64(x uint64) int
```

OnesCount64 returns the number of one bits (“population count”) in x.

​	OnesCount64 返回 x 中的 1 位数（“基数计数”）。

#### OnesCount64 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount64(%064b) = %d\n", 14, bits.OnesCount64(14))
}
Output:

OnesCount64(0000000000000000000000000000000000000000000000000000000000001110) = 3
```

### func OnesCount8

```go
func OnesCount8(x uint8) int
```

OnesCount8 returns the number of one bits (“population count”) in x.

​	OnesCount8 返回 x 中的 1 位数（“基数计数”）。

#### OnesCount8 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount8(%08b) = %d\n", 14, bits.OnesCount8(14))
}
Output:

OnesCount8(00001110) = 3
```

### func Rem <- go1.14

```go
func Rem(hi, lo, y uint) uint
```

Rem returns the remainder of (hi, lo) divided by y. Rem panics for y == 0 (division by zero) but, unlike Div, it doesn’t panic on a quotient overflow.

​	Rem 返回 (hi, lo) 除以 y 的余数。Rem 在 y == 0（除以零）时会引发 panic，但与 Div 不同，它不会在商溢出时引发 panic。

### func Rem32 <- go1.14

```go
func Rem32(hi, lo, y uint32) uint32
```

Rem32 returns the remainder of (hi, lo) divided by y. Rem32 panics for y == 0 (division by zero) but, unlike Div32, it doesn’t panic on a quotient overflow.

​	Rem32 返回 (hi, lo) 除以 y 的余数。Rem32 在 y == 0（除以零）时会引发 panic，但与 Div32 不同，它不会在商溢出时引发 panic。

### func Rem64 <- go1.14

```go
func Rem64(hi, lo, y uint64) uint64
```

Rem64 returns the remainder of (hi, lo) divided by y. Rem64 panics for y == 0 (division by zero) but, unlike Div64, it doesn’t panic on a quotient overflow.

​	Rem64 返回 (hi, lo) 除以 y 的余数。Rem64 在 y == 0（除以零）时会引发 panic，但与 Div64 不同，它不会在商溢出时引发 panic。

### func Reverse

```go
func Reverse(x uint) uint
```

Reverse returns the value of x with its bits in reversed order.

​	Reverse 返回 x 的值，其位按相反的顺序排列。

### func Reverse16

```go
func Reverse16(x uint16) uint16
```

Reverse16 returns the value of x with its bits in reversed order.

​	Reverse16 返回 x 的值，其位按相反的顺序排列。

#### Reverse16 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%016b\n", 19)
	fmt.Printf("%016b\n", bits.Reverse16(19))
}
Output:

0000000000010011
1100100000000000
```

### func Reverse32

```go
func Reverse32(x uint32) uint32
```

Reverse32 returns the value of x with its bits in reversed order.

​	Reverse32 返回 x 的值，其位按相反的顺序排列。

#### Reverse32 Example 

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%032b\n", 19)
	fmt.Printf("%032b\n", bits.Reverse32(19))
}
Output:

00000000000000000000000000010011
11001000000000000000000000000000
```

### func Reverse64

```go
func Reverse64(x uint64) uint64
```

Reverse64 returns the value of x with its bits in reversed order.

​	Reverse64 返回 x 的值，其位按相反的顺序排列。

#### Reverse64 Example 

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%064b\n", 19)
	fmt.Printf("%064b\n", bits.Reverse64(19))
}
Output:

0000000000000000000000000000000000000000000000000000000000010011
1100100000000000000000000000000000000000000000000000000000000000
```

### func Reverse8

```go
func Reverse8(x uint8) uint8
```

Reverse8 returns the value of x with its bits in reversed order.

​	Reverse8 返回其位以相反顺序排列的 x 的值。

#### Reverse8 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%08b\n", 19)
	fmt.Printf("%08b\n", bits.Reverse8(19))
}
Output:

00010011
11001000
```

### func ReverseBytes

```go
func ReverseBytes(x uint) uint
```

ReverseBytes returns the value of x with its bytes in reversed order.

​	ReverseBytes 返回其字节以相反顺序排列的 x 的值。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

### func ReverseBytes16

```go
func ReverseBytes16(x uint16) uint16
```

ReverseBytes16 returns the value of x with its bytes in reversed order.

​	ReverseBytes16 返回其字节以相反顺序排列的 x 的值。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### ReverseBytes16 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%016b\n", 15)
	fmt.Printf("%016b\n", bits.ReverseBytes16(15))
}
Output:

0000000000001111
0000111100000000
```

### func ReverseBytes32

```go
func ReverseBytes32(x uint32) uint32
```

ReverseBytes32 returns the value of x with its bytes in reversed order.

​	ReverseBytes32 返回其字节顺序相反的 x 的值。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### ReverseBytes32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%032b\n", 15)
	fmt.Printf("%032b\n", bits.ReverseBytes32(15))
}
Output:

00000000000000000000000000001111
00001111000000000000000000000000
```

### func ReverseBytes64

```go
func ReverseBytes64(x uint64) uint64
```

ReverseBytes64 returns the value of x with its bytes in reversed order.

​	ReverseBytes64 返回其字节顺序相反的 x 的值。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### ReverseBytes64 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%064b\n", 15)
	fmt.Printf("%064b\n", bits.ReverseBytes64(15))
}
Output:

0000000000000000000000000000000000000000000000000000000000001111
0000111100000000000000000000000000000000000000000000000000000000
```

### func RotateLeft

```go
func RotateLeft(x uint, k int) uint
```

RotateLeft returns the value of x rotated left by (k mod UintSize) bits. To rotate x right by k bits, call RotateLeft(x, -k).

​	RotateLeft 返回 x 左移 (k mod UintSize) 位的值。要将 x 右移 k 位，请调用 RotateLeft(x, -k)。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

### func RotateLeft16

```go
func RotateLeft16(x uint16, k int) uint16
```

RotateLeft16 returns the value of x rotated left by (k mod 16) bits. To rotate x right by k bits, call RotateLeft16(x, -k).

​	RotateLeft16 返回 x 左移 (k mod 16) 位的值。要将 x 右移 k 位，请调用 RotateLeft16(x, -k)。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### RotateLeft16 Example 

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%016b\n", 15)
	fmt.Printf("%016b\n", bits.RotateLeft16(15, 2))
	fmt.Printf("%016b\n", bits.RotateLeft16(15, -2))
}
Output:

0000000000001111
0000000000111100
1100000000000011
```

### func RotateLeft32

```go
func RotateLeft32(x uint32, k int) uint32
```

RotateLeft32 returns the value of x rotated left by (k mod 32) bits. To rotate x right by k bits, call RotateLeft32(x, -k).

​	RotateLeft32 返回 x 左移 (k mod 32) 位的值。要将 x 右移 k 位，请调用 RotateLeft32(x, -k)。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### RotateLeft32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%032b\n", 15)
	fmt.Printf("%032b\n", bits.RotateLeft32(15, 2))
	fmt.Printf("%032b\n", bits.RotateLeft32(15, -2))
}
Output:

00000000000000000000000000001111
00000000000000000000000000111100
11000000000000000000000000000011
```

### func RotateLeft64

```go
func RotateLeft64(x uint64, k int) uint64
```

RotateLeft64 returns the value of x rotated left by (k mod 64) bits. To rotate x right by k bits, call RotateLeft64(x, -k).

​	RotateLeft64 返回 x 左移 (k mod 64) 位后的值。要将 x 右移 k 位，请调用 RotateLeft64(x, -k)。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### RotateLeft64 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%064b\n", 15)
	fmt.Printf("%064b\n", bits.RotateLeft64(15, 2))
	fmt.Printf("%064b\n", bits.RotateLeft64(15, -2))
}
Output:

0000000000000000000000000000000000000000000000000000000000001111
0000000000000000000000000000000000000000000000000000000000111100
1100000000000000000000000000000000000000000000000000000000000011
```

### func RotateLeft8

```go
func RotateLeft8(x uint8, k int) uint8
```

RotateLeft8 returns the value of x rotated left by (k mod 8) bits. To rotate x right by k bits, call RotateLeft8(x, -k).

​	RotateLeft8 返回 x 左移 (k mod 8) 位后的值。要将 x 右移 k 位，请调用 RotateLeft8(x, -k)。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### RotateLeft8 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", bits.RotateLeft8(15, 2))
	fmt.Printf("%08b\n", bits.RotateLeft8(15, -2))
}
Output:

00001111
00111100
11000011
```

### func Sub <- go1.12

```go
func Sub(x, y, borrow uint) (diff, borrowOut uint)
```

Sub returns the difference of x, y and borrow: diff = x - y - borrow. The borrow input must be 0 or 1; otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.

​	Sub 返回 x、y 和借位的差值：diff = x - y - borrow。借位输入必须为 0 或 1；否则行为未定义。borrowOut 输出保证为 0 或 1。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

### func Sub32 <- go1.12

```go
func Sub32(x, y, borrow uint32) (diff, borrowOut uint32)
```

Sub32 returns the difference of x, y and borrow, diff = x - y - borrow. The borrow input must be 0 or 1; otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.

​	Sub32 返回 x、y 和借位 diff = x - y - borrow 的差值。借位输入必须为 0 或 1；否则行为未定义。borrowOut 输出保证为 0 或 1。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### Sub32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 33<<32 + 23
	n1 := []uint32{33, 23}
	// Second number is 21<<32 + 12
	n2 := []uint32{21, 12}
	// Sub them together without producing carry.
	d1, carry := bits.Sub32(n1[1], n2[1], 0)
	d0, _ := bits.Sub32(n1[0], n2[0], carry)
	nsum := []uint32{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// First number is 3<<32 + 2147483647
	n1 = []uint32{3, 0x7fffffff}
	// Second number is 1<<32 + 2147483648
	n2 = []uint32{1, 0x80000000}
	// Sub them together producing carry.
	d1, carry = bits.Sub32(n1[1], n2[1], 0)
	d0, _ = bits.Sub32(n1[0], n2[0], carry)
	nsum = []uint32{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
}
Output:

[33 23] - [21 12] = [12 11] (carry bit was 0)
[3 2147483647] - [1 2147483648] = [1 4294967295] (carry bit was 1)
```

### func Sub64 <- go1.12

```go
func Sub64(x, y, borrow uint64) (diff, borrowOut uint64)
```

Sub64 returns the difference of x, y and borrow: diff = x - y - borrow. The borrow input must be 0 or 1; otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.

​	Sub64 返回 x、y 和借位 diff = x - y - borrow 的差值。借位输入必须为 0 或 1；否则行为未定义。borrowOut 输出保证为 0 或 1。

This function’s execution time does not depend on the inputs.

​	此函数的执行时间不依赖于输入。

#### Sub64 Example 

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// First number is 33<<64 + 23
	n1 := []uint64{33, 23}
	// Second number is 21<<64 + 12
	n2 := []uint64{21, 12}
	// Sub them together without producing carry.
	d1, carry := bits.Sub64(n1[1], n2[1], 0)
	d0, _ := bits.Sub64(n1[0], n2[0], carry)
	nsum := []uint64{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// First number is 3<<64 + 9223372036854775807
	n1 = []uint64{3, 0x7fffffffffffffff}
	// Second number is 1<<64 + 9223372036854775808
	n2 = []uint64{1, 0x8000000000000000}
	// Sub them together producing carry.
	d1, carry = bits.Sub64(n1[1], n2[1], 0)
	d0, _ = bits.Sub64(n1[0], n2[0], carry)
	nsum = []uint64{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
}
Output:

[33 23] - [21 12] = [12 11] (carry bit was 0)
[3 9223372036854775807] - [1 9223372036854775808] = [1 18446744073709551615] (carry bit was 1)
```

### func TrailingZeros

```go
func TrailingZeros(x uint) int
```

TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.

​	TrailingZeros 返回 x 中尾随零位的数量；结果为 x == 0 时的 UintSize。

### func TrailingZeros16

```go
func TrailingZeros16(x uint16) int
```

TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.

​	TrailingZeros16 返回 x 中尾随零位的数量；结果为 x == 0 时的 16。

#### TrailingZeros16 Example 

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros16(%016b) = %d\n", 14, bits.TrailingZeros16(14))
}
Output:

TrailingZeros16(0000000000001110) = 1
```

### func TrailingZeros32

```go
func TrailingZeros32(x uint32) int
```

TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.

​	TrailingZeros32 返回 x 中尾随零位的数量；结果为 x == 0 时的 32。

#### TrailingZeros32 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros32(%032b) = %d\n", 14, bits.TrailingZeros32(14))
}
Output:

TrailingZeros32(00000000000000000000000000001110) = 1
```

### func TrailingZeros64

```go
func TrailingZeros64(x uint64) int
```

TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.

​	TrailingZeros64 返回 x 中尾随零位的数量；结果为 x == 0 时的 64。

#### TrailingZeros64 Example

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros64(%064b) = %d\n", 14, bits.TrailingZeros64(14))
}
Output:

TrailingZeros64(0000000000000000000000000000000000000000000000000000000000001110) = 1
```

### func TrailingZeros8

```go
func TrailingZeros8(x uint8) int
```

TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.

​	TrailingZeros8 返回 x 中尾随零位的数量；对于 x == 0，结果为 8。

#### TrailingZeros8 Example
``` go 
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros8(%08b) = %d\n", 14, bits.TrailingZeros8(14))
}
Output:

TrailingZeros8(00001110) = 1
```

## 类型

This section is empty.