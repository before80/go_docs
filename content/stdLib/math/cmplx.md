+++
title = "cmplx"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/math/cmplx@go1.23.0](https://pkg.go.dev/math/cmplx@go1.23.0)

Package cmplx provides basic constants and mathematical functions for complex numbers. Special case handling conforms to the C99 standard Annex G IEC 60559-compatible complex arithmetic.

​	cmplx 包提供基本常量和复数的数学函数。特殊情况处理符合 C99 标准附录 G IEC 60559 兼容的复数运算。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Abs

```go
func Abs(x complex128) float64
```

Abs returns the absolute value (also called the modulus) of x.

​	Abs 返回 x 的绝对值（也称为模数）。

#### Abs Example

```go
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Printf("%.1f", cmplx.Abs(3+4i))
}
Output:

5.0
```

### func Acos

```go
func Acos(x complex128) complex128
```

Acos returns the inverse cosine of x.

​	Acos 返回 x 的反余弦。

### func Acosh

```go
func Acosh(x complex128) complex128
```

Acosh returns the inverse hyperbolic cosine of x.

​	Acosh 返回 x 的反双曲余弦。

### func Asin

```go
func Asin(x complex128) complex128
```

Asin returns the inverse sine of x.

​	Asin 返回 x 的反正弦。

### func Asinh

```go
func Asinh(x complex128) complex128
```

Asinh returns the inverse hyperbolic sine of x.

​	Asinh 返回 x 的反双曲正弦。

### func Atan

```go
func Atan(x complex128) complex128
```

Atan returns the inverse tangent of x.

​	Atan 返回 x 的反正切。

### func Atanh

```go
func Atanh(x complex128) complex128
```

Atanh returns the inverse hyperbolic tangent of x.

​	Atanh 返回 x 的反双曲正切。

### func Conj

```go
func Conj(x complex128) complex128
```

Conj returns the complex conjugate of x.

​	Conj 返回 x 的复共轭。

### func Cos

```go
func Cos(x complex128) complex128
```

Cos returns the cosine of x.

​	Cos 返回 x 的余弦。

### func Cosh

```go
func Cosh(x complex128) complex128
```

Cosh returns the hyperbolic cosine of x.

​	Cosh 返回 x 的双曲余弦。

### func Cot

```go
func Cot(x complex128) complex128
```

Cot returns the cotangent of x.

​	Cot 返回 x 的余切。

### func Exp

```go
func Exp(x complex128) complex128
```

Exp returns $e^x$, the base-e exponential of x.

​	Exp 返回 $e^x$，即 x 的 e 为底的指数。

#### Exp Example 

ExampleExp computes Euler’s identity.

​	ExampleExp 计算欧拉恒等式。

```go
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	fmt.Printf("%.1f", cmplx.Exp(1i*math.Pi)+1)
}
```

### func Inf

```go
func Inf() complex128
```

Inf returns a complex infinity, complex(+Inf, +Inf).

​	Inf 返回一个复数无穷大，complex(+Inf, +Inf)。

### func IsInf

```go
func IsInf(x complex128) bool
```

IsInf reports whether either real(x) or imag(x) is an infinity.

​	IsInf 报告 real(x) 或 imag(x) 是否为无穷大。

### func IsNaN

```go
func IsNaN(x complex128) bool
```

IsNaN reports whether either real(x) or imag(x) is NaN and neither is an infinity.

​	IsNaN 报告 real(x) 或 imag(x) 是否为 NaN，且两者都不是无穷大。

### func Log

```go
func Log(x complex128) complex128
```

Log returns the natural logarithm of x.

​	Log 返回 x 的自然对数。

### func Log10

```go
func Log10(x complex128) complex128
```

Log10 returns the decimal logarithm of x.

​	Log10 返回 x 的十进制对数。

### func NaN

```go
func NaN() complex128
```

NaN returns a complex “not-a-number” value.

​	NaN 返回一个复数“非数字”值。

### func Phase

```go
func Phase(x complex128) float64
```

Phase returns the phase (also called the argument) of x. The returned value is in the range [-Pi, Pi].

​	Phase 返回 x 的相位（也称为参数）。返回值在 [-Pi, Pi] 范围内。

### func Polar

```go
func Polar(x complex128) (r, θ float64)
```

Polar returns the absolute value r and phase θ of x, such that $x = r * e^{θi}$. The phase is in the range [-Pi, Pi].

​	Polar 返回 x 的绝对值 r 和相位 θ，使得 $x = r * e^{θi}$。相位在 [-Pi, Pi] 范围内。

#### Polar Example

```go
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	r, theta := cmplx.Polar(2i)
	fmt.Printf("r: %.1f, θ: %.1f*π", r, theta/math.Pi)
}
Output:

r: 2.0, θ: 0.5*π
```

### func Pow

```go
func Pow(x, y complex128) complex128
```

Pow returns $x^y$, the base-x exponential of y. For generalized compatibility with math.Pow:

​	Pow 返回 $x^y$，即 y 的以 x 为底的指数。为了与 math.Pow 实现通用的兼容性：

```
Pow(0, ±0) returns 1+0i
Pow(0, c) for real(c)<0 returns Inf+0i if imag(c) is zero, otherwise Inf+Inf i.
```

### func Rect

```go
func Rect(r, θ float64) complex128
```

Rect returns the complex number x with polar coordinates r, θ.

​	Rect 返回具有极坐标 r、θ 的复数 x。

### func Sin

```go
func Sin(x complex128) complex128
```

Sin returns the sine of x.

​	Sin 返回 x 的正弦值。

### func Sinh

```go
func Sinh(x complex128) complex128
```

Sinh returns the hyperbolic sine of x.

​	Sinh 返回 x 的双曲正弦。

### func Sqrt

```go
func Sqrt(x complex128) complex128
```

Sqrt returns the square root of x. The result r is chosen so that real(r) ≥ 0 and imag(r) has the same sign as imag(x).

​	Sqrt 返回 x 的平方根。结果 r 的选择方式使得 real(r) ≥ 0，imag(r) 与 imag(x) 同号。

### func Tan

```go
func Tan(x complex128) complex128
```

Tan returns the tangent of x.

​	Tan 返回 x 的正切。

### func Tanh

```go
func Tanh(x complex128) complex128
```

Tanh returns the hyperbolic tangent of x.

​	Tanh 返回 x 的双曲正切。

## 类型

This section is empty.