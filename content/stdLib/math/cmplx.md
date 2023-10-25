+++
title = "cmplx"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/math/cmplx@go1.20.1



Package cmplx provides basic constants and mathematical functions for complex numbers. Special case handling conforms to the C99 standard Annex G IEC 60559-compatible complex arithmetic.









## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Abs 

``` go 
func Abs(x complex128) float64
```

Abs returns the absolute value (also called the modulus) of x.

##### Example
``` go 
```

#### func Acos 

``` go 
func Acos(x complex128) complex128
```

Acos returns the inverse cosine of x.

#### func Acosh 

``` go 
func Acosh(x complex128) complex128
```

Acosh returns the inverse hyperbolic cosine of x.

#### func Asin 

``` go 
func Asin(x complex128) complex128
```

Asin returns the inverse sine of x.

#### func Asinh 

``` go 
func Asinh(x complex128) complex128
```

Asinh returns the inverse hyperbolic sine of x.

#### func Atan 

``` go 
func Atan(x complex128) complex128
```

Atan returns the inverse tangent of x.

#### func Atanh 

``` go 
func Atanh(x complex128) complex128
```

Atanh returns the inverse hyperbolic tangent of x.

#### func Conj 

``` go 
func Conj(x complex128) complex128
```

Conj returns the complex conjugate of x.

#### func Cos 

``` go 
func Cos(x complex128) complex128
```

Cos returns the cosine of x.

#### func Cosh 

``` go 
func Cosh(x complex128) complex128
```

Cosh returns the hyperbolic cosine of x.

#### func Cot 

``` go 
func Cot(x complex128) complex128
```

Cot returns the cotangent of x.

#### func Exp 

``` go 
func Exp(x complex128) complex128
```

Exp returns e**x, the base-e exponential of x.

##### Example
``` go 
```

#### func Inf 

``` go 
func Inf() complex128
```

Inf returns a complex infinity, complex(+Inf, +Inf).

#### func IsInf 

``` go 
func IsInf(x complex128) bool
```

IsInf reports whether either real(x) or imag(x) is an infinity.

#### func IsNaN 

``` go 
func IsNaN(x complex128) bool
```

IsNaN reports whether either real(x) or imag(x) is NaN and neither is an infinity.

#### func Log 

``` go 
func Log(x complex128) complex128
```

Log returns the natural logarithm of x.

#### func Log10 

``` go 
func Log10(x complex128) complex128
```

Log10 returns the decimal logarithm of x.

#### func NaN 

``` go 
func NaN() complex128
```

NaN returns a complex "not-a-number" value.

#### func Phase 

``` go 
func Phase(x complex128) float64
```

Phase returns the phase (also called the argument) of x. The returned value is in the range [-Pi, Pi].

#### func Polar 

``` go 
func Polar(x complex128) (r, θ float64)
```

Polar returns the absolute value r and phase θ of x, such that x = r * e**θi. The phase is in the range [-Pi, Pi].

##### Example
``` go 
```

#### func Pow 

``` go 
func Pow(x, y complex128) complex128
```

Pow returns x**y, the base-x exponential of y. For generalized compatibility with math.Pow:

```
Pow(0, ±0) returns 1+0i
Pow(0, c) for real(c)<0 returns Inf+0i if imag(c) is zero, otherwise Inf+Inf i.
```

#### func Rect 

``` go 
func Rect(r, θ float64) complex128
```

Rect returns the complex number x with polar coordinates r, θ.

#### func Sin 

``` go 
func Sin(x complex128) complex128
```

Sin returns the sine of x.

#### func Sinh 

``` go 
func Sinh(x complex128) complex128
```

Sinh returns the hyperbolic sine of x.

#### func Sqrt 

``` go 
func Sqrt(x complex128) complex128
```

Sqrt returns the square root of x. The result r is chosen so that real(r) ≥ 0 and imag(r) has the same sign as imag(x).

#### func Tan 

``` go 
func Tan(x complex128) complex128
```

Tan returns the tangent of x.

#### func Tanh 

``` go 
func Tanh(x complex128) complex128
```

Tanh returns the hyperbolic tangent of x.

## 类型

This section is empty.