+++
title = "math"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/math@go1.21.3

Package math provides basic constants and mathematical functions.

​	`math`包提供基本常量和数学函数。

This package does not guarantee bit-identical results across architectures.

​	该包不能保证在不同架构下获得完全相同的结果。

## 常量 

### 数学常量

#### E

#### Pi

#### Phi

#### Sqrt2

#### SqrtE

#### SqrtPi

#### SqrtPhi

#### Ln2

#### Log2E

#### Ln10

#### Log10E

``` go 
const (
	E   = 2.71828182845904523536028747135266249775724709369995957496696763 // https://oeis.org/A001113
	Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
	Phi = 1.61803398874989484820458683436563811772030917980576286213544862 // https://oeis.org/A001622

	Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339

	Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 // https://oeis.org/A002162
	Log2E  = 1 / Ln2
	Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 // https://oeis.org/A002392
	Log10E = 1 / Ln10
)
```

Mathematical constants.

数学常量。

### 浮点数的极限值

#### MaxFloat32

#### SmallestNonzeroFloat32

#### MaxFloat64

#### SmallestNonzeroFloat64

``` go 
const (
	MaxFloat32             = 0x1p127 * (1 + (1 - 0x1p-23)) // 3.40282346638528859811704183484516925440e+38
	SmallestNonzeroFloat32 = 0x1p-126 * 0x1p-23            // 1.401298464324817070923729583289916131280e-45

	MaxFloat64             = 0x1p1023 * (1 + (1 - 0x1p-52)) // 1.79769313486231570814527423731704356798070e+308
	SmallestNonzeroFloat64 = 0x1p-1022 * 0x1p-52            // 4.9406564584124654417656879286822137236505980e-324
)
```

Floating-point limit values. Max is the largest finite value representable by the type. SmallestNonzero is the smallest positive, non-zero value representable by the type.

​	浮点数的极限值。Max 是该类型能表示的最大有限值。SmallestNonzero 是该类型能表示的最小正非零值。

### 整数的极限值

#### MaxInt

#### MinInt

#### MaxInt8

#### MinInt8

#### MaxInt16

#### MinInt16

#### MaxInt32

#### MinInt32

#### MaxInt64

#### MinInt64

#### MaxUint8

#### MaxUint16

#### MaxUint32

#### MaxUint64

``` go 
const (
	MaxInt    = 1<<(intSize-1) - 1  // MaxInt32 or MaxInt64 depending on intSize.
	MinInt    = -1 << (intSize - 1) // MinInt32 or MinInt64 depending on intSize.
	MaxInt8   = 1<<7 - 1            // 127
	MinInt8   = -1 << 7             // -128
	MaxInt16  = 1<<15 - 1           // 32767
	MinInt16  = -1 << 15            // -32768
	MaxInt32  = 1<<31 - 1           // 2147483647
	MinInt32  = -1 << 31            // -2147483648
	MaxInt64  = 1<<63 - 1           // 9223372036854775807
	MinInt64  = -1 << 63            // -9223372036854775808
	MaxUint   = 1<<intSize - 1      // MaxUint32 or MaxUint64 depending on intSize.
	MaxUint8  = 1<<8 - 1            // 255
	MaxUint16 = 1<<16 - 1           // 65535
	MaxUint32 = 1<<32 - 1           // 4294967295
	MaxUint64 = 1<<64 - 1           // 18446744073709551615
)
```

Integer limit values.

​	整数的极限值。

## 变量

This section is empty.

## 函数

### func Abs 

``` go 
func Abs(x float64) float64
```

Abs returns the absolute value of x.

​	Abs 返回 x 的绝对值。

Special cases are:

​	特殊情况如下：

```
Abs(±Inf) = +Inf
Abs(NaN) = NaN
```

#### Abs Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)
}

Output:

2.0
2.0
```

### func Acos 

``` go 
func Acos(x float64) float64
```

Acos returns the arccosine, in radians, of x.

​	Acos返回x的反余弦值（以弧度为单位）。

Special case is:

特殊情况如下：

```
Acos(x) = NaN if x < -1 or x > 1
```

#### Acos Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Acos(1))
}


Output:

0.00
```

### func Acosh 

``` go 
func Acosh(x float64) float64
```

Acosh returns the inverse hyperbolic cosine of x.

​	Acosh返回x的反双曲余弦值。

Special cases are:

特殊情况如下：

```
Acosh(+Inf) = +Inf
Acosh(x) = NaN if x < 1
Acosh(NaN) = NaN
```

#### Acosh  Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Acosh(1))
}

Output:

0.00
```

### func Asin 

``` go 
func Asin(x float64) float64
```

Asin returns the arcsine, in radians, of x.

​	Asin返回x的反正弦值（以弧度为单位）。

Special cases are:

​	特殊情况如下：

```
Asin(±0) = ±0
Asin(x) = NaN if x < -1 or x > 1
```

#### Asin Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Asin(0))
}

Output:

0.00
```

### func Asinh 

``` go 
func Asinh(x float64) float64
```

Asinh returns the inverse hyperbolic sine of x.

​	Asinh返回x的反双曲正弦值。

Special cases are:

​	特殊情况如下：

```
Asinh(±0) = ±0
Asinh(±Inf) = ±Inf
Asinh(NaN) = NaN
```

#### Asinh Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Asinh(0))
}


Output:

0.00
```

### func Atan 

``` go 
func Atan(x float64) float64
```

Atan returns the arctangent, in radians, of x.

​	Atan返回x的反正切值（以弧度为单位）。

Special cases are:

​	特殊情况如下：

```
Atan(±0) = ±0
Atan(±Inf) = ±Pi/2
```

#### Atan Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Atan(0))
}


Output:

0.00
```

### func Atan2 

``` go 
func Atan2(y, x float64) float64
```

Atan2 returns the arc tangent of y/x, using the signs of the two to determine the quadrant of the return value.

​	Atan2返回y/x的反正切值，根据y和x的符号确定返回值所在的象限。

Special cases are (in order):

​	特殊情况按顺序为：

```
Atan2(y, NaN) = NaN
Atan2(NaN, x) = NaN
Atan2(+0, x>=0) = +0
Atan2(-0, x>=0) = -0
Atan2(+0, x<=-0) = +Pi
Atan2(-0, x<=-0) = -Pi
Atan2(y>0, 0) = +Pi/2
Atan2(y<0, 0) = -Pi/2
Atan2(+Inf, +Inf) = +Pi/4
Atan2(-Inf, +Inf) = -Pi/4
Atan2(+Inf, -Inf) = 3Pi/4
Atan2(-Inf, -Inf) = -3Pi/4
Atan2(y, +Inf) = 0
Atan2(y>0, -Inf) = +Pi
Atan2(y<0, -Inf) = -Pi
Atan2(+Inf, x) = +Pi/2
Atan2(-Inf, x) = -Pi/2
```

#### Atan2 Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Atan2(0, 0))
}

Output:

0.00
```

### func Atanh 

``` go 
func Atanh(x float64) float64
```

Atanh returns the inverse hyperbolic tangent of x.

​	Atanh返回x的反双曲切线值。

Special cases are:

​	特殊情况如下：

```
Atanh(1) = +Inf
Atanh(±0) = ±0
Atanh(-1) = -Inf
Atanh(x) = NaN if x < -1 or x > 1
Atanh(NaN) = NaN
```

#### Atanh Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Atanh(0))
}

Output:

0.004
```

### func Cbrt 

``` go 
func Cbrt(x float64) float64
```

Cbrt returns the cube root of x.

​	Cbrt返回x的立方根。

Special cases are:

​	特殊情况如下：

```
Cbrt(±0) = ±0
Cbrt(±Inf) = ±Inf
Cbrt(NaN) = NaN
```

#### Cbrt Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f\n", math.Cbrt(8))
	fmt.Printf("%.2f\n", math.Cbrt(27))
}
Output:

2.00
3.00
```

### func Ceil 

``` go 
func Ceil(x float64) float64
```

Ceil returns the least integer value greater than or equal to x.

​	Ceil返回大于或等于x的最小整数值。

Special cases are:

​	特殊情况如下：

```
Ceil(±0) = ±0
Ceil(±Inf) = ±Inf
Ceil(NaN) = NaN
```

#### Ceil Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	c := math.Ceil(1.49)
	fmt.Printf("%.1f", c)
}
Output:

2.0
```

### func Copysign 

``` go 
func Copysign(f, sign float64) float64
```

Copysign returns a value with the magnitude of f and the sign of sign.

​	Copysign返回具有f的绝对值和sign的符号的值。

#### Copysign Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Copysign(3.2, -1))
}
Output:

-3.20
```

### func Cos 

``` go 
func Cos(x float64) float64
```

Cos returns the cosine of the radian argument x.

​	Cos返回弧度参数x的余弦值。

Special cases are:

​	特殊情况如下：

```
Cos(±Inf) = NaN
Cos(NaN) = NaN
```

#### Cos Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Cos(math.Pi/2))
}
Output:

0.00
```

### func Cosh 

``` go 
func Cosh(x float64) float64
```

Cosh returns the hyperbolic cosine of x.

​	Cosh返回x的双曲余弦值。

Special cases are:

​	特殊情况如下：

```
Cosh(±0) = 1
Cosh(±Inf) = +Inf
Cosh(NaN) = NaN
```

#### Cosh Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Cosh(0))
}
Output:

1.00
```

### func Dim 

``` go 
func Dim(x, y float64) float64
```

Dim returns the maximum of x-y or 0.

​	Dim返回x-y和0之间的最大值。

Special cases are:

​	特殊情况如下：

```
Dim(+Inf, +Inf) = NaN
Dim(-Inf, -Inf) = NaN
Dim(x, NaN) = Dim(NaN, x) = NaN
```

#### Dim Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f\n", math.Dim(4, -2))
	fmt.Printf("%.2f\n", math.Dim(-4, 2))
}
Output:

6.00
0.00
```

### func Erf 

``` go 
func Erf(x float64) float64
```

Erf returns the error function of x.

​	Erf返回x的误差函数值。

Special cases are:

​	特殊情况如下：

```
Erf(+Inf) = 1
Erf(-Inf) = -1
Erf(NaN) = NaN
```

### func Erfc 

``` go 
func Erfc(x float64) float64
```

Erfc returns the complementary error function of x.

​	Erfc返回x的互补误差函数值。

Special cases are:

​	特殊情况如下：

```
Erfc(+Inf) = 0
Erfc(-Inf) = 2
Erfc(NaN) = NaN
```

### func Erfcinv  <- go1.10

``` go 
func Erfcinv(x float64) float64
```

Erfcinv returns the inverse of Erfc(x).

​	Erfcinv返回Erfc(x)的反函数值。

Special cases are:

​	特殊情况如下：

```
Erfcinv(0) = +Inf
Erfcinv(2) = -Inf
Erfcinv(x) = NaN if x < 0 or x > 2
Erfcinv(NaN) = NaN
```

### func Erfinv  <- go1.10

``` go 
func Erfinv(x float64) float64
```

Erfinv returns the inverse error function of x.

​	Erfinv返回x的反误差函数值。

Special cases are:

​	特殊情况如下：

```
Erfinv(1) = +Inf
Erfinv(-1) = -Inf
Erfinv(x) = NaN if x < -1 or x > 1
Erfinv(NaN) = NaN
```

### func Exp 

``` go 
func Exp(x float64) float64
```

Exp returns `e**x`, the base-e exponential of x.

​	Exp返回`e**x`，即x的以e为底的指数函数值。

Special cases are:

​	特殊情况如下：

```
Exp(+Inf) = +Inf
Exp(NaN) = NaN
```

Very large values overflow to 0 or +Inf. Very small values underflow to 1.

​	非常大的值会溢出为0或+Inf。非常小的值会下溢为1。

#### Exp Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f\n", math.Exp(1))
	fmt.Printf("%.2f\n", math.Exp(2))
	fmt.Printf("%.2f\n", math.Exp(-1))
}
Output:

2.72
7.39
0.37
```

### func Exp2 

``` go 
func Exp2(x float64) float64
```

Exp2 returns `2**x`, the base-2 exponential of x.

​	Exp2返回`2**x`，即x的以2为底的指数函数值。

Special cases are the same as Exp.

​	特殊情况与Exp相同。

#### Exp2 Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f\n", math.Exp2(1))
	fmt.Printf("%.2f\n", math.Exp2(-3))
}
Output:

2.00
0.12
```

### func Expm1 

``` go 
func Expm1(x float64) float64
```

Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more accurate than Exp(x) - 1 when x is near zero.

​	Expm1返回`e**x` - 1，即x的以e为底的指数函数值减去1。当x接近零时，Expm1比Exp(x) - 1更准确。

Special cases are:

​	特殊情况如下：

```
Expm1(+Inf) = +Inf
Expm1(-Inf) = -1
Expm1(NaN) = NaN
```

Very large values overflow to -1 or +Inf.

​	非常大的值会溢出为-1或+Inf。

#### Expm1 Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.6f\n", math.Expm1(0.01))
	fmt.Printf("%.6f\n", math.Expm1(-1))
}
Output:

0.010050
-0.632121
```

### func FMA  <- go1.14

``` go 
func FMA(x, y, z float64) float64
```

FMA returns x * y + z, computed with only one rounding. (That is, FMA returns the fused multiply-add of x, y, and z.)

​	FMA返回x * y + z，使用仅进行一次舍入的计算。（也就是说，FMA返回x、y和z的融合乘加结果。）

### func Float32bits 

``` go 
func Float32bits(f float32) uint32
```

Float32bits returns the IEEE 754 binary representation of f, with the sign bit of f and the result in the same bit position. Float32bits(Float32frombits(x)) == x.

​	Float32bits返回f的IEEE 754二进制表示，其中包含f的符号位，并且结果与符号位处于相同的位位置。Float32bits(Float32frombits(x)) == x。

### func Float32frombits 

``` go 
func Float32frombits(b uint32) float32
```

Float32frombits returns the floating-point number corresponding to the IEEE 754 binary representation b, with the sign bit of b and the result in the same bit position. Float32frombits(Float32bits(x)) == x.

​	Float32frombits返回与IEEE 754二进制表示b对应的浮点数，其中包含b的符号位，并且结果与符号位处于相同的位位置。Float32frombits(Float32bits(x)) == x。

### func Float64bits 

``` go 
func Float64bits(f float64) uint64
```

Float64bits returns the IEEE 754 binary representation of f, with the sign bit of f and the result in the same bit position, and Float64bits(Float64frombits(x)) == x.

​	Float64bits返回f的IEEE 754二进制表示，其中包含f的符号位，并且结果与符号位处于相同的位位置。Float64bits(Float64frombits(x)) == x。

### func Float64frombits 

``` go 
func Float64frombits(b uint64) float64
```

Float64frombits returns the floating-point number corresponding to the IEEE 754 binary representation b, with the sign bit of b and the result in the same bit position. Float64frombits(Float64bits(x)) == x.

​	Float64frombits返回与IEEE 754二进制表示b对应的浮点数，其中包含b的符号位，并且结果与符号位处于相同的位位置。Float64frombits(Float64bits(x)) == x。

### func Floor 

``` go 
func Floor(x float64) float64
```

Floor returns the greatest integer value less than or equal to x.

​	Floor返回小于等于x的最大整数值。

Special cases are:

​	特殊情况如下：

```
Floor(±0) = ±0
Floor(±Inf) = ±Inf
Floor(NaN) = NaN
```

#### Floor Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	c := math.Floor(1.51)
	fmt.Printf("%.1f", c)
}
Output:

1.0
```

### func Frexp 

``` go 
func Frexp(f float64) (frac float64, exp int)
```

Frexp breaks f into a normalized fraction and an integral power of two. It returns frac and exp satisfying f == frac × 2**exp, with the absolute value of frac in the interval [½, 1).

​	Frexp将f拆分为标准化的分数和2的整数次幂。它返回满足f == frac × 2**exp的frac和exp，其中frac的绝对值位于[½, 1)区间内。

Special cases are:

​	特殊情况如下：

```
Frexp(±0) = ±0, 0
Frexp(±Inf) = ±Inf, 0
Frexp(NaN) = NaN, 0
```

### func Gamma 

``` go 
func Gamma(x float64) float64
```

Gamma returns the Gamma function of x.

​	Gamma返回x的Gamma函数值。

Special cases are:

​	特殊情况如下：

```
Gamma(+Inf) = +Inf
Gamma(+0) = +Inf
Gamma(-0) = -Inf
Gamma(x) = NaN for integer x < 0
Gamma(-Inf) = NaN
Gamma(NaN) = NaN
```

### func Hypot 

``` go 
func Hypot(p, q float64) float64
```

Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and underflow.

​	Hypot返回Sqrt(p*p + q*q)，注意避免不必要的溢出和下溢。

Special cases are:

​	特殊情况如下：

```
Hypot(±Inf, q) = +Inf
Hypot(p, ±Inf) = +Inf
Hypot(NaN, q) = NaN
Hypot(p, NaN) = NaN
```

### func Ilogb 

``` go 
func Ilogb(x float64) int
```

Ilogb returns the binary exponent of x as an integer.

​	Ilogb返回x的二进制指数作为整数。

Special cases are:

​	特殊情况如下：

```
Ilogb(±Inf) = MaxInt32
Ilogb(0) = MinInt32
Ilogb(NaN) = MaxInt32
```

### func Inf 

``` go 
func Inf(sign int) float64
```

Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.

​	如果sign >= 0，则Inf返回正无穷大；如果sign < 0，则返回负无穷大。

### func IsInf 

``` go 
func IsInf(f float64, sign int) bool
```

IsInf reports whether f is an infinity, according to sign. If sign > 0, IsInf reports whether f is positive infinity. If sign < 0, IsInf reports whether f is negative infinity. If sign == 0, IsInf reports whether f is either infinity.

​	IsInf报告f是否为无穷大，根据sign的值。如果sign > 0，则IsInf报告f是否为正无穷大。如果sign < 0，则IsInf报告f是否为负无穷大。如果sign == 0，则IsInf报告f是否为无穷大。

### func IsNaN 

``` go 
func IsNaN(f float64) (is bool)
```

IsNaN reports whether f is an IEEE 754 "not-a-number" value.

​	IsNaN报告f是否为IEEE 754的"非数"值。

### func J0 

``` go 
func J0(x float64) float64
```

J0 returns the order-zero Bessel function of the first kind.

​	J0返回第一类零阶贝塞尔函数。

Special cases are:

​	特殊情况如下：

```
J0(±Inf) = 0
J0(0) = 1
J0(NaN) = NaN
```

### func J1 

``` go 
func J1(x float64) float64
```

J1 returns the order-one Bessel function of the first kind.

​	J1返回第一类一阶贝塞尔函数。

Special cases are:

​	特殊情况如下：

```
J1(±Inf) = 0
J1(NaN) = NaN
```

### func Jn 

``` go 
func Jn(n int, x float64) float64
```

Jn returns the order-n Bessel function of the first kind.

​	Jn返回第一类n阶贝塞尔函数。

Special cases are:

​	特殊情况如下：

```
Jn(n, ±Inf) = 0
Jn(n, NaN) = NaN
```

### func Ldexp 

``` go 
func Ldexp(frac float64, exp int) float64
```

Ldexp is the inverse of Frexp. It returns frac × 2**exp.

​	Ldexp是Frexp的反函数。它返回frac × 2**exp。

Special cases are:

​	特殊情况如下：

```
Ldexp(±0, exp) = ±0
Ldexp(±Inf, exp) = ±Inf
Ldexp(NaN, exp) = NaN
```

### func Lgamma 

``` go 
func Lgamma(x float64) (lgamma float64, sign int)
```

Lgamma returns the natural logarithm and sign (-1 or +1) of Gamma(x).

​	Lgamma返回Gamma(x)的自然对数和符号（-1或+1）。

Special cases are:

​	特殊情况如下：

```
Lgamma(+Inf) = +Inf
Lgamma(0) = +Inf
Lgamma(-integer) = +Inf
Lgamma(-Inf) = -Inf
Lgamma(NaN) = NaN
```

### func Log 

``` go 
func Log(x float64) float64
```

Log returns the natural logarithm of x.

​	Log返回x的自然对数。

Special cases are:

​	特殊情况如下：

```
Log(+Inf) = +Inf
Log(0) = -Inf
Log(x < 0) = NaN
Log(NaN) = NaN
```

#### Log Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	x := math.Log(1)
	fmt.Printf("%.1f\n", x)

	y := math.Log(2.7183)
	fmt.Printf("%.1f\n", y)
}
Output:

0.0
1.0
```

### func Log10 

``` go 
func Log10(x float64) float64
```

Log10 returns the decimal logarithm of x. The special cases are the same as for Log.

​	Log10返回x的十进制对数。特殊情况与Log相同。

#### Log10 Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.1f", math.Log10(100))
}
Output:

2.0
```

### func Log1p 

``` go 
func Log1p(x float64) float64
```

Log1p returns the natural logarithm of 1 plus its argument x. It is more accurate than Log(1 + x) when x is near zero.

​	Log1p返回1加上其参数x的自然对数。当x接近零时，它比Log(1 + x)更准确。

Special cases are:

​	特殊情况如下：

```
Log1p(+Inf) = +Inf
Log1p(±0) = ±0
Log1p(-1) = -Inf
Log1p(x < -1) = NaN
Log1p(NaN) = NaN
```

### func Log2 

``` go 
func Log2(x float64) float64
```

Log2 returns the binary logarithm of x. The special cases are the same as for Log.

​	Log2返回x的二进制对数。特殊情况与Log相同。

#### Log2 Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.1f", math.Log2(256))
}
Output:

8.0
```

### func Logb 

``` go 
func Logb(x float64) float64
```

Logb returns the binary exponent of x.

​	Logb返回x的二进制指数。

Special cases are:

​	特殊情况如下：

```
Logb(±Inf) = +Inf
Logb(0) = -Inf
Logb(NaN) = NaN
```

### func Max 

``` go 
func Max(x, y float64) float64
```

Max returns the larger of x or y.

​	Max返回x和y中较大的值。

Special cases are:

​	特殊情况如下：

```
Max(x, +Inf) = Max(+Inf, x) = +Inf
Max(x, NaN) = Max(NaN, x) = NaN
Max(+0, ±0) = Max(±0, +0) = +0
Max(-0, -0) = -0
```

### func Min 

``` go 
func Min(x, y float64) float64
```

Min returns the smaller of x or y.

​	Min返回x和y中较小的值。

Special cases are:

​	特殊情况如下：

```
Min(x, -Inf) = Min(-Inf, x) = -Inf
Min(x, NaN) = Min(NaN, x) = NaN
Min(-0, ±0) = Min(±0, -0) = -0
```

### func Mod 

``` go 
func Mod(x, y float64) float64
```

Mod returns the floating-point remainder of x/y. The magnitude of the result is less than y and its sign agrees with that of x.

​	Mod返回x除以y的浮点余数。结果的绝对值小于y，符号与x相同。

Special cases are:

​	特殊情况如下：

```
Mod(±Inf, y) = NaN
Mod(NaN, y) = NaN
Mod(x, 0) = NaN
Mod(x, ±Inf) = x
Mod(x, NaN) = NaN
```

#### Mod Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	c := math.Mod(7, 4)
	fmt.Printf("%.1f", c)
}
Output:

3.0
```

### func Modf 

``` go 
func Modf(f float64) (int float64, frac float64)
```

Modf returns integer and fractional floating-point numbers that sum to f. Both values have the same sign as f.

​	Modf返回和为f的整数部分和小数部分的浮点数。两个值的符号与f相同。

Special cases are:

​	特殊情况如下：

```
Modf(±Inf) = ±Inf, NaN
Modf(NaN) = NaN, NaN
```

#### Modf Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	int, frac := math.Modf(3.14)
	fmt.Printf("%.2f, %.2f\n", int, frac)

	int, frac = math.Modf(-2.71)
	fmt.Printf("%.2f, %.2f\n", int, frac)
}
Output:

3.00, 0.14
-2.00, -0.71
```

### func NaN 

``` go 
func NaN() float64
```

NaN returns an IEEE 754 "not-a-number" value.

​	NaN返回一个IEEE 754的"非数"值。

### func Nextafter 

``` go 
func Nextafter(x, y float64) (r float64)
```

Nextafter returns the next representable float64 value after x towards y.

​	Nextafter返回x朝向y的下一个可表示的float64值。

Special cases are:

​	特殊情况如下：

```
Nextafter(x, x)   = x
Nextafter(NaN, y) = NaN
Nextafter(x, NaN) = NaN
```

### func Nextafter32  <- go1.4

``` go 
func Nextafter32(x, y float32) (r float32)
```

Nextafter32 returns the next representable float32 value after x towards y.

​	Nextafter32返回x朝向y的下一个可表示的float32值。

Special cases are:

​	特殊情况如下：

```
Nextafter32(x, x)   = x
Nextafter32(NaN, y) = NaN
Nextafter32(x, NaN) = NaN
```

### func Pow 

``` go 
func Pow(x, y float64) float64
```

Pow returns x**y, the base-x exponential of y.

​	Pow返回x的y次幂，即x**y。

Special cases are (in order):

```
Pow(x, ±0) = 1 for any x
Pow(1, y) = 1 for any y
Pow(x, 1) = x for any x
Pow(NaN, y) = NaN
Pow(x, NaN) = NaN
Pow(±0, y) = ±Inf for y an odd integer < 0
Pow(±0, -Inf) = +Inf
Pow(±0, +Inf) = +0
Pow(±0, y) = +Inf for finite y < 0 and not an odd integer
Pow(±0, y) = ±0 for y an odd integer > 0
Pow(±0, y) = +0 for finite y > 0 and not an odd integer
Pow(-1, ±Inf) = 1
Pow(x, +Inf) = +Inf for |x| > 1
Pow(x, -Inf) = +0 for |x| > 1
Pow(x, +Inf) = +0 for |x| < 1
Pow(x, -Inf) = +Inf for |x| < 1
Pow(+Inf, y) = +Inf for y > 0
Pow(+Inf, y) = +0 for y < 0
Pow(-Inf, y) = Pow(-0, -y)
Pow(x, y) = NaN for finite x < 0 and finite non-integer y
```

#### Pow Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	c := math.Pow(2, 3)
	fmt.Printf("%.1f", c)
}
Output:

8.0
```

### func Pow10 

``` go 
func Pow10(n int) float64
```

Pow10 returns 10**n, the base-10 exponential of n.

​	Pow10返回10的n次幂，即10**n。

Special cases are:

​	特殊情况如下：

```
Pow10(n) =    0 for n < -323
Pow10(n) = +Inf for n > 308
```

#### Pow10 Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	c := math.Pow10(2)
	fmt.Printf("%.1f", c)
}
Output:

100.0
```

### func Remainder 

``` go 
func Remainder(x, y float64) float64
```

Remainder returns the IEEE 754 floating-point remainder of x/y.

​	Pow10返回10的n次幂，即10**n。

Special cases are:

​	特殊情况如下：

```
Remainder(±Inf, y) = NaN
Remainder(NaN, y) = NaN
Remainder(x, 0) = NaN
Remainder(x, ±Inf) = x
Remainder(x, NaN) = NaN
```

#### Remainder Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.1f", math.Remainder(100, 30))
}
Output:

10.0
```

### func Round  <- go1.10

``` go 
func Round(x float64) float64
```

Round returns the nearest integer, rounding half away from zero.

​	Round返回最接近x的整数，向远离零的方向四舍五入。

Special cases are:

​	特殊情况如下：

```
Round(±0) = ±0
Round(±Inf) = ±Inf
Round(NaN) = NaN
```

#### Round  Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	p := math.Round(10.5)
	fmt.Printf("%.1f\n", p)

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)
}
Output:

11.0
-11.0
```

### func RoundToEven  <- go1.10

``` go 
func RoundToEven(x float64) float64
```

RoundToEven returns the nearest integer, rounding ties to even.

​	RoundToEven返回最接近x的整数，四舍五入时选择偶数。

Special cases are:

​	特殊情况如下：

```
RoundToEven(±0) = ±0
RoundToEven(±Inf) = ±Inf
RoundToEven(NaN) = NaN
```

#### RoundToEven  Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	u := math.RoundToEven(11.5)
	fmt.Printf("%.1f\n", u)

	d := math.RoundToEven(12.5)
	fmt.Printf("%.1f\n", d)
}
Output:

12.0
12.0
```

### func Signbit 

``` go 
func Signbit(x float64) bool
```

Signbit reports whether x is negative or negative zero.

​	Signbit报告x是否为负数或负零。

### func Sin 

``` go 
func Sin(x float64) float64
```

Sin returns the sine of the radian argument x.

​	Sin返回弧度参数x的正弦值。

Special cases are:

​	特殊情况如下：

```
Sin(±0) = ±0
Sin(±Inf) = NaN
Sin(NaN) = NaN
```

#### Sin Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Sin(math.Pi))
}
Output:

0.00
```

### func Sincos 

``` go 
func Sincos(x float64) (sin, cos float64)
```

Sincos returns Sin(x), Cos(x).

​	Sincos返回Sin(x)和Cos(x)。

Special cases are:

​	特殊情况如下：

```
Sincos(±0) = ±0, 1
Sincos(±Inf) = NaN, NaN
Sincos(NaN) = NaN, NaN
```

#### Sincos Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	sin, cos := math.Sincos(0)
	fmt.Printf("%.2f, %.2f", sin, cos)
}
Output:

0.00, 1.00
```

### func Sinh 

``` go 
func Sinh(x float64) float64
```

Sinh returns the hyperbolic sine of x.

​	Sinh返回x的双曲正弦值。

Special cases are:

​	特殊情况如下：

```
Sinh(±0) = ±0
Sinh(±Inf) = ±Inf
Sinh(NaN) = NaN
```

#### Sinh Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Sinh(0))
}
Output:

0.00
```

### func Sqrt 

``` go 
func Sqrt(x float64) float64
```

Sqrt returns the square root of x.

​	Sqrt返回x的平方根。

Special cases are:

​	特殊情况如下：

```
Sqrt(+Inf) = +Inf
Sqrt(±0) = ±0
Sqrt(x < 0) = NaN
Sqrt(NaN) = NaN
```

#### Sqrt Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		a = 3
		b = 4
	)
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("%.1f", c)
}
Output:

0.00
```

### func Tan 

``` go 
func Tan(x float64) float64
```

Tan returns the tangent of the radian argument x.

Tan返回弧度参数x的正切值。

Special cases are:

​	特殊情况如下：

```
Tan(±0) = ±0
Tan(±Inf) = NaN
Tan(NaN) = NaN
```

#### Tan Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Tan(0))
}
Output:

0.00
```

### func Tanh 

``` go 
func Tanh(x float64) float64
```

Tanh returns the hyperbolic tangent of x.

​	Tanh返回x的双曲正切值。

Special cases are:

​	特殊情况如下：

```
Tanh(±0) = ±0
Tanh(±Inf) = ±1
Tanh(NaN) = NaN
```

#### Tanh Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f", math.Tanh(0))
}
Output:

0.00
```

### func Trunc 

``` go 
func Trunc(x float64) float64
```

Trunc returns the integer value of x.

​	Trunc返回x的整数部分。

Special cases are:

​	特殊情况如下：

```
Trunc(±0) = ±0
Trunc(±Inf) = ±Inf
Trunc(NaN) = NaN
```

#### Trunc Example
``` go 
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.2f\n", math.Trunc(math.Pi))
	fmt.Printf("%.2f\n", math.Trunc(-1.2345))
}
Output:

3.00
-1.00
```

### func Y0 

``` go 
func Y0(x float64) float64
```

Y0 returns the order-zero Bessel function of the second kind.

​	Y0返回第二类零阶贝塞尔函数。

Special cases are:

​	特殊情况如下：

```
Y0(+Inf) = 0
Y0(0) = -Inf
Y0(x < 0) = NaN
Y0(NaN) = NaN
```

### func Y1 

``` go 
func Y1(x float64) float64
```

Y1 returns the order-one Bessel function of the second kind.

​	Y1返回第二类一阶贝塞尔函数。

Special cases are:

​	特殊情况如下：

```
Y1(+Inf) = 0
Y1(0) = -Inf
Y1(x < 0) = NaN
Y1(NaN) = NaN
```

### func Yn 

``` go 
func Yn(n int, x float64) float64
```

Yn returns the order-n Bessel function of the second kind.

​	Yn返回第二类n阶贝塞尔函数。

Special cases are:

​	特殊情况如下：

```
Yn(n, +Inf) = 0
Yn(n ≥ 0, 0) = -Inf
Yn(n < 0, 0) = +Inf if n is odd, -Inf if n is even
Yn(n, x < 0) = NaN
Yn(n, NaN) = NaN
```

## 类型

This section is empty.