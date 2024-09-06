+++
title = "big"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/math/big@go1.23.0](https://pkg.go.dev/math/big@go1.23.0)

Package big implements arbitrary-precision arithmetic (big numbers). The following numeric types are supported:

​	big 包 实现任意精度算术（大数字）。支持以下数字类型：

```
Int    signed integers
Rat    rational numbers
Float  floating-point numbers
```

The zero value for an Int, Rat, or Float correspond to 0. Thus, new values can be declared in the usual ways and denote 0 without further initialization:

​	Int、Rat 或 Float 的零值对应于 0。因此，可以用通常的方式声明新值，并表示 0，而无需进一步初始化：

```go
var x Int        // &x is an *Int of value 0
var r = &Rat{}   // r is a *Rat of value 0
y := new(Float)  // y is a *Float of value 0
```

Alternatively, new values can be allocated and initialized with factory functions of the form:

​	或者，可以使用以下形式的工厂函数分配和初始化新值：

```go
func NewT(v V) *T
```

For instance, NewInt(x) returns an *Int set to the value of the int64 argument x, NewRat(a, b) returns a *Rat set to the fraction a/b where a and b are int64 values, and NewFloat(f) returns a *Float initialized to the float64 argument f. More flexibility is provided with explicit setters, for instance:

​	例如，NewInt(x) 返回一个 *Int，设置为 int64 参数 x 的值，NewRat(a, b) 返回一个 *Rat，设置为分数 a/b，其中 a 和 b 是 int64 值，NewFloat(f) 返回一个 *Float，初始化为 float64 参数 f。显式设置器提供了更大的灵活性，例如：

```go
var z1 Int
z1.SetUint64(123)                 // z1 := 123
z2 := new(Rat).SetFloat64(1.25)   // z2 := 5/4
z3 := new(Float).SetInt(z1)       // z3 := 123.0
```

Setters, numeric operations and predicates are represented as methods of the form:

​	设置器、数字运算和谓词表示为以下形式的方法：

```go
func (z *T) SetV(v V) *T          // z = v
func (z *T) Unary(x *T) *T        // z = unary x
func (z *T) Binary(x, y *T) *T    // z = x binary y
func (x *T) Pred() P              // p = pred(x)
```

with T one of Int, Rat, or Float. For unary and binary operations, the result is the receiver (usually named z in that case; see below); if it is one of the operands x or y it may be safely overwritten (and its memory reused).

​	使用 Int、Rat 或 Float 中的一个 T。对于一元和二元运算，结果是接收者（通常在这种情况下命名为 z；见下文）；如果它是操作数 x 或 y 之一，则可以安全地覆盖它（并重新使用其内存）。

Arithmetic expressions are typically written as a sequence of individual method calls, with each call corresponding to an operation. The receiver denotes the result and the method arguments are the operation’s operands. For instance, given three *Int values a, b and c, the invocation

​	算术表达式通常写为一系列单独的方法调用，每个调用对应一个运算。接收者表示结果，方法参数是运算的操作数。例如，给定三个 *Int 值 a、b 和 c，调用

```
c.Add(a, b)
```

computes the sum a + b and stores the result in c, overwriting whatever value was held in c before. Unless specified otherwise, operations permit aliasing of parameters, so it is perfectly ok to write

​	计算总和 a + b 并将结果存储在 c 中，覆盖之前 c 中保存的任何值。除非另有说明，否则运算允许参数别名，因此完全可以写

```
sum.Add(sum, x)
```

to accumulate values x in a sum.

​	在总和中累加值 x。

(By always passing in a result value via the receiver, memory use can be much better controlled. Instead of having to allocate new memory for each result, an operation can reuse the space allocated for the result value, and overwrite that value with the new result in the process.)

​	（通过始终通过接收者传递结果值，可以更好地控制内存使用。不必为每个结果分配新内存，运算可以重用为结果值分配的空间，并在过程中用新结果覆盖该值。）

Notational convention: Incoming method parameters (including the receiver) are named consistently in the API to clarify their use. Incoming operands are usually named x, y, a, b, and so on, but never z. A parameter specifying the result is named z (typically the receiver).

​	符号约定：传入的方法参数（包括接收器）在 API 中命名一致，以明确其用途。传入的操作数通常命名为 x、y、a、b 等，但从不命名为 z。指定结果的参数命名为 z（通常是接收器）。

For instance, the arguments for (*Int).Add are named x and y, and because the receiver specifies the result destination, it is called z:

​	例如，(*Int).Add 的参数命名为 x 和 y，并且因为接收器指定了结果目标，所以它被称为 z：

```go
func (z *Int) Add(x, y *Int) *Int
```

Methods of this form typically return the incoming receiver as well, to enable simple call chaining.

​	这种形式的方法通常也会返回传入的接收器，以启用简单的调用链。

Methods which don’t require a result value to be passed in (for instance, Int.Sign), simply return the result. In this case, the receiver is typically the first operand, named x:

​	不需要传入结果值的方法（例如，Int.Sign）只需返回结果。在这种情况下，接收器通常是第一个操作数，命名为 x：

```go
func (x *Int) Sign() int
```

Various methods support conversions between strings and corresponding numeric values, and vice versa: *Int, *Rat, and *Float values implement the Stringer interface for a (default) string representation of the value, but also provide SetString methods to initialize a value from a string in a variety of supported formats (see the respective SetString documentation).

​	各种方法支持字符串和相应数值之间的转换，反之亦然：*Int、*Rat 和 *Float 值实现 Stringer 接口，用于值的（默认）字符串表示形式，但也提供 SetString 方法，以从各种受支持格式的字符串中初始化值（请参阅相应的 SetString 文档）。

Finally, *Int, *Rat, and *Float satisfy the fmt package’s Scanner interface for scanning and (except for *Rat) the Formatter interface for formatted printing.

​	最后，*Int、*Rat 和 *Float 满足 fmt 包的 Scanner 接口，用于扫描，并且（除了 *Rat）满足 Formatter 接口，用于格式化打印。

## Example (EConvergents)

This example demonstrates how to use big.Rat to compute the first 15 terms in the sequence of rational convergents for the constant e (base of natural logarithm).

​	此示例演示如何使用 big.Rat 计算常数 e（自然对数的底数）的理性收敛序列的前 15 个项。

```go
package main

import (
	"fmt"
	"math/big"
)

// Use the classic continued fraction for e
//
//	e = [1; 0, 1, 1, 2, 1, 1, ... 2n, 1, 1, ...]
//
// i.e., for the nth term, use
//
//	   1          if   n mod 3 != 1
//	(n-1)/3 * 2   if   n mod 3 == 1
func recur(n, lim int64) *big.Rat {
	term := new(big.Rat)
	if n%3 != 1 {
		term.SetInt64(1)
	} else {
		term.SetInt64((n - 1) / 3 * 2)
	}

	if n > lim {
		return term
	}

	// Directly initialize frac as the fractional
	// inverse of the result of recur.
	frac := new(big.Rat).Inv(recur(n+1, lim))

	return term.Add(term, frac)
}

// This example demonstrates how to use big.Rat to compute the
// first 15 terms in the sequence of rational convergents for
// the constant e (base of natural logarithm).
func main() {
	for i := 1; i <= 15; i++ {
		r := recur(0, int64(i))

		// Print r both as a fraction and as a floating-point number.
		// Since big.Rat implements fmt.Formatter, we can use %-13s to
		// get a left-aligned string representation of the fraction.
		fmt.Printf("%-13s = %s\n", r, r.FloatString(8))
	}

}
Output:

2/1           = 2.00000000
3/1           = 3.00000000
8/3           = 2.66666667
11/4          = 2.75000000
19/7          = 2.71428571
87/32         = 2.71875000
106/39        = 2.71794872
193/71        = 2.71830986
1264/465      = 2.71827957
1457/536      = 2.71828358
2721/1001     = 2.71828172
23225/8544    = 2.71828184
25946/9545    = 2.71828182
49171/18089   = 2.71828183
517656/190435 = 2.71828183
```

## Example (Fibonacci)

This example demonstrates how to use big.Int to compute the smallest Fibonacci number with 100 decimal digits and to test whether it is prime.

​	此示例演示如何使用 big.Int 计算具有 100 个十进制数字的最小斐波那契数，并测试它是否是质数。

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Initialize limit as 10^99, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	fmt.Println(a) // 100-digit Fibonacci number

	// Test a for primality.
	// (ProbablyPrimes' argument sets the number of Miller-Rabin
	// rounds to be performed. 20 is a good value.)
	fmt.Println(a.ProbablyPrime(20))

}
Output:

1344719667586153181419716641724567886890850696275767987106294472017884974410332069524504824747437757
false
```

## Example (Sqrt2) 

This example shows how to use big.Float to compute the square root of 2 with a precision of 200 bits, and how to print the result as a decimal number.

​	此示例演示如何使用 big.Float 计算精度为 200 位的 2 的平方根，以及如何将结果打印为十进制数。

```go
package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// We'll do computations with 200 bits of precision in the mantissa.
	const prec = 200

	// Compute the square root of 2 using Newton's Method. We start with
	// an initial estimate for sqrt(2), and then iterate:
	//     x_{n+1} = 1/2 * ( x_n + (2.0 / x_n) )

	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	steps := int(math.Log2(prec))

	// Initialize values we need for the computation.
	two := new(big.Float).SetPrec(prec).SetInt64(2)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	// Use 1 as the initial estimate.
	x := new(big.Float).SetPrec(prec).SetInt64(1)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	t := new(big.Float)

	// Iterate.
	for i := 0; i <= steps; i++ {
		t.Quo(two, x)  // t = 2.0 / x_n
		t.Add(x, t)    // t = x_n + (2.0 / x_n)
		x.Mul(half, t) // x_{n+1} = 0.5 * t
	}

	// We can use the usual fmt.Printf verbs since big.Float implements fmt.Formatter
	fmt.Printf("sqrt(2) = %.50f\n", x)

	// Print the error between 2 and x*x.
	t.Mul(x, x) // t = x*x
	fmt.Printf("error = %e\n", t.Sub(two, t))

}
Output:

sqrt(2) = 1.41421356237309504880168872420969807856967187537695
error = 0.000000e+00
```

## 常量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/math/big/float.go;l=96)

```go
const (
	MaxExp  = math.MaxInt32  // largest supported exponent
	MinExp  = math.MinInt32  // smallest supported exponent
	MaxPrec = math.MaxUint32 // largest (theoretically) supported precision; likely memory-limited
)
```

Exponent and precision limits.

​	指数和精度限制。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/math/big/natconv.go;l=24)

```go
const MaxBase = 10 + ('z' - 'a' + 1) + ('Z' - 'A' + 1)
```

MaxBase is the largest number base accepted for string conversions.

​	MaxBase 是字符串转换接受的最大数字基数。

## 变量

This section is empty.

## 函数

### func Jacobi <- go1.5

```go
func Jacobi(x, y *Int) int
```

Jacobi returns the Jacobi symbol (x/y), either +1, -1, or 0. The y argument must be an odd integer.

​	Jacobi 返回 Jacobi 符号 (x/y)，为 +1、-1 或 0。y 参数必须是奇数整数。

## 类型

### type Accuracy <- go1.5

```go
type Accuracy int8
```

Accuracy describes the rounding error produced by the most recent operation that generated a Float value, relative to the exact value.

​	Accuracy 描述了生成 Float 值的最近操作产生的舍入误差，相对于确切值。

```go
const (
	Below Accuracy = -1
	Exact Accuracy = 0
	Above Accuracy = +1
)
```

Constants describing the Accuracy of a Float.

​	描述 Float 精度的常量。

#### (Accuracy) String <- go1.5

```go
func (i Accuracy) String() string
```

### type ErrNaN <- go1.5

```go
type ErrNaN struct {
	// contains filtered or unexported fields
}
```

An ErrNaN panic is raised by a Float operation that would lead to a NaN under IEEE-754 rules. An ErrNaN implements the error interface.

​	ErrNaN 恐慌是由 Float 操作引发的，该操作会导致 IEEE-754 规则下的 NaN。ErrNaN 实现错误接口。

#### (ErrNaN) Error <- go1.5

```go
func (err ErrNaN) Error() string
```

### type Float <- go1.5

```go
type Float struct {
	// contains filtered or unexported fields
}
```

A nonzero finite Float represents a multi-precision floating point number

​	非零有限 Float 表示多精度浮点数

```
sign × mantissa × 2**exponent
```

with 0.5 <= mantissa < 1.0, and MinExp <= exponent <= MaxExp. A Float may also be zero (+0, -0) or infinite (+Inf, -Inf). All Floats are ordered, and the ordering of two Floats x and y is defined by x.Cmp(y).

​	，其中 0.5 <= 尾数 < 1.0，且 MinExp <= 指数 <= MaxExp。Float 还可以是零（+0、-0）或无穷大（+Inf、-Inf）。所有 Float 都已排序，两个 Float x 和 y 的排序由 x.Cmp(y) 定义。

Each Float value also has a precision, rounding mode, and accuracy. The precision is the maximum number of mantissa bits available to represent the value. The rounding mode specifies how a result should be rounded to fit into the mantissa bits, and accuracy describes the rounding error with respect to the exact result.

​	每个 Float 值还具有精度、舍入模式和准确度。精度是可用于表示该值的尾数位数的最大值。舍入模式指定结果应如何舍入以适合尾数位，准确度描述相对于精确结果的舍入误差。

Unless specified otherwise, all operations (including setters) that specify a *Float variable for the result (usually via the receiver with the exception of MantExp), round the numeric result according to the precision and rounding mode of the result variable.

​	除非另有说明，否则所有指定结果的 *Float 变量（通常通过接收器，MantExp 除外）的操作（包括设置器）都会根据结果变量的精度和舍入模式对数字结果进行舍入。

If the provided result precision is 0 (see below), it is set to the precision of the argument with the largest precision value before any rounding takes place, and the rounding mode remains unchanged. Thus, uninitialized Floats provided as result arguments will have their precision set to a reasonable value determined by the operands, and their mode is the zero value for RoundingMode (ToNearestEven).

​	如果提供的结果精度为 0（见下文），则将其设置为在进行任何舍入之前具有最大精度值的参数的精度，舍入模式保持不变。因此，作为结果参数提供的未初始化 Float 的精度将设置为由操作数确定的合理值，其模式是 RoundingMode（ToNearestEven）的零值。

By setting the desired precision to 24 or 53 and using matching rounding mode (typically ToNearestEven), Float operations produce the same results as the corresponding float32 or float64 IEEE-754 arithmetic for operands that correspond to normal (i.e., not denormal) float32 or float64 numbers. Exponent underflow and overflow lead to a 0 or an Infinity for different values than IEEE-754 because Float exponents have a much larger range.

​	通过将所需精度设置为 24 或 53 并使用匹配的舍入模式（通常为 ToNearestEven），Float 运算会对对应于正常（即非非规格化）float32 或 float64 数字的操作数产生与相应的 float32 或 float64 IEEE-754 算术相同的结果。由于 Float 指数的范围更大，因此指数下溢和上溢会导致 0 或无穷大，不同于 IEEE-754。

The zero (uninitialized) value for a Float is ready to use and represents the number +0.0 exactly, with precision 0 and rounding mode ToNearestEven.

​	Float 的零（未初始化）值可以使用，并精确表示数字 +0.0，精度为 0，舍入模式为 ToNearestEven。

Operations always take pointer arguments (*Float) rather than Float values, and each unique Float value requires its own unique *Float pointer. To “copy” a Float value, an existing (or newly allocated) Float must be set to a new value using the Float.Set method; shallow copies of Floats are not supported and may lead to errors.

​	运算总是采用指针参数 (*Float) 而不是 Float 值，并且每个唯一的 Float 值都需要其自己的唯一 *Float 指针。要“复制”Float 值，必须使用 Float.Set 方法将现有（或新分配的）Float 设置为新值；不支持 Float 的浅层副本，并且可能导致错误。

#### Example (Shift)

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Implement Float "shift" by modifying the (binary) exponents directly.
	for s := -5; s <= 5; s++ {
		x := big.NewFloat(0.5)
		x.SetMantExp(x, x.MantExp(nil)+s) // shift x by s
		fmt.Println(x)
	}
}

Output:

0.015625
0.03125
0.0625
0.125
0.25
0.5
1
2
4
8
16
```

#### func NewFloat <- go1.5

```go
func NewFloat(x float64) *Float
```

NewFloat allocates and returns a new Float set to x, with precision 53 and rounding mode ToNearestEven. NewFloat panics with ErrNaN if x is a NaN.

​	NewFloat 分配并返回一个新的 Float，设置为 x，精度为 53，舍入模式为 ToNearestEven。如果 x 是 NaN，NewFloat 会引发 ErrNaN。

#### func ParseFloat <- go1.5

```go
func ParseFloat(s string, base int, prec uint, mode RoundingMode) (f *Float, b int, err error)
```

ParseFloat is like f.Parse(s, base) with f set to the given precision and rounding mode.

​	ParseFloat 类似于 f.Parse(s, base)，其中 f 设置为给定的精度和舍入模式。

#### (*Float) Abs <- go1.5

```go
func (z *Float) Abs(x *Float) *Float
```

Abs sets z to the (possibly rounded) value |x| (the absolute value of x) and returns z.

​	Abs 将 z 设置为（可能已舍入）值 |x|（x 的绝对值），并返回 z。

#### (*Float) Acc <- go1.5

```go
func (x *Float) Acc() Accuracy
```

Acc returns the accuracy of x produced by the most recent operation, unless explicitly documented otherwise by that operation.

​	Acc 返回由最近的操作产生的 x 的精度，除非该操作明确记录了其他内容。

#### (*Float) Add <- go1.5

```go
func (z *Float) Add(x, y *Float) *Float
```

Add sets z to the rounded sum x+y and returns z. If z’s precision is 0, it is changed to the larger of x’s or y’s precision before the operation. Rounding is performed according to z’s precision and rounding mode; and z’s accuracy reports the result error relative to the exact (not rounded) result. Add panics with ErrNaN if x and y are infinities with opposite signs. The value of z is undefined in that case.

​	Add 将 z 设置为舍入的和 x+y，并返回 z。如果 z 的精度为 0，则在操作之前将其更改为 x 或 y 的较大精度。舍入根据 z 的精度和舍入模式执行；z 的精度报告相对于确切（未舍入）结果的结果误差。如果 x 和 y 是带相反符号的无穷大，则 Add 会引发 ErrNaN。在这种情况下，z 的值是未定义的。

##### Add Example

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Operate on numbers of different precision.
	var x, y, z big.Float
	x.SetInt64(1000)          // x is automatically set to 64bit precision
	y.SetFloat64(2.718281828) // y is automatically set to 53bit precision
	z.SetPrec(32)
	z.Add(&x, &y)
	fmt.Printf("x = %.10g (%s, prec = %d, acc = %s)\n", &x, x.Text('p', 0), x.Prec(), x.Acc())
	fmt.Printf("y = %.10g (%s, prec = %d, acc = %s)\n", &y, y.Text('p', 0), y.Prec(), y.Acc())
	fmt.Printf("z = %.10g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
}

Output:

x = 1000 (0x.fap+10, prec = 64, acc = Exact)
y = 2.718281828 (0x.adf85458248cd8p+2, prec = 53, acc = Exact)
z = 1002.718282 (0x.faadf854p+10, prec = 32, acc = Below)
```

#### (*Float) Append <- go1.5

```go
func (x *Float) Append(buf []byte, fmt byte, prec int) []byte
```

Append appends to buf the string form of the floating-point number x, as generated by x.Text, and returns the extended buffer.

​	Append 将浮点数 x 的字符串形式追加到 buf，由 x.Text 生成，并返回扩展的缓冲区。

#### (*Float) Cmp <- go1.5

```go
func (x *Float) Cmp(y *Float) int
```

Cmp compares x and y and returns:

​	Cmp 比较 x 和 y 并返回：

```
-1 if x <  y
 0 if x == y (incl. -0 == 0, -Inf == -Inf, and +Inf == +Inf)
+1 if x >  y
```

##### Cmp Example

```go
package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	inf := math.Inf(1)
	zero := 0.0

	operands := []float64{-inf, -1.2, -zero, 0, +1.2, +inf}

	fmt.Println("   x     y  cmp")
	fmt.Println("---------------")
	for _, x64 := range operands {
		x := big.NewFloat(x64)
		for _, y64 := range operands {
			y := big.NewFloat(y64)
			fmt.Printf("%4g  %4g  %3d\n", x, y, x.Cmp(y))
		}
		fmt.Println()
	}

}
Output:

   x     y  cmp
---------------
-Inf  -Inf    0
-Inf  -1.2   -1
-Inf    -0   -1
-Inf     0   -1
-Inf   1.2   -1
-Inf  +Inf   -1

-1.2  -Inf    1
-1.2  -1.2    0
-1.2    -0   -1
-1.2     0   -1
-1.2   1.2   -1
-1.2  +Inf   -1

  -0  -Inf    1
  -0  -1.2    1
  -0    -0    0
  -0     0    0
  -0   1.2   -1
  -0  +Inf   -1

   0  -Inf    1
   0  -1.2    1
   0    -0    0
   0     0    0
   0   1.2   -1
   0  +Inf   -1

 1.2  -Inf    1
 1.2  -1.2    1
 1.2    -0    1
 1.2     0    1
 1.2   1.2    0
 1.2  +Inf   -1

+Inf  -Inf    1
+Inf  -1.2    1
+Inf    -0    1
+Inf     0    1
+Inf   1.2    1
+Inf  +Inf    0
```

#### (*Float) Copy <- go1.5

```go
func (z *Float) Copy(x *Float) *Float
```

Copy sets z to x, with the same precision, rounding mode, and accuracy as x, and returns z. x is not changed even if z and x are the same.

​	Copy 将 z 设置为 x，其精度、舍入模式和准确性与 x 相同，并返回 z。即使 z 和 x 相同，x 也不会更改。

#### (*Float) Float32 <- go1.5

```go
func (x *Float) Float32() (float32, Accuracy)
```

Float32 returns the float32 value nearest to x. If x is too small to be represented by a float32 (|x| < math.SmallestNonzeroFloat32), the result is (0, Below) or (-0, Above), respectively, depending on the sign of x. If x is too large to be represented by a float32 (|x| > math.MaxFloat32), the result is (+Inf, Above) or (-Inf, Below), depending on the sign of x.

​	Float32 返回最接近 x 的 float32 值。如果 x 太小而无法用 float32 表示（|x| < math.SmallestNonzeroFloat32），则结果分别为 (0, Below) 或 (-0, Above)，具体取决于 x 的符号。如果 x 太大而无法用 float32 表示（|x| > math.MaxFloat32），则结果为 (+Inf, Above) 或 (-Inf, Below)，具体取决于 x 的符号。

#### (*Float) Float64 <- go1.5

```go
func (x *Float) Float64() (float64, Accuracy)
```

Float64 returns the float64 value nearest to x. If x is too small to be represented by a float64 (|x| < math.SmallestNonzeroFloat64), the result is (0, Below) or (-0, Above), respectively, depending on the sign of x. If x is too large to be represented by a float64 (|x| > math.MaxFloat64), the result is (+Inf, Above) or (-Inf, Below), depending on the sign of x.

​	Float64 返回最接近 x 的 float64 值。如果 x 太小而无法用 float64 表示（|x| < math.SmallestNonzeroFloat64），则结果分别为 (0, Below) 或 (-0, Above)，具体取决于 x 的符号。如果 x 太大而无法用 float64 表示（|x| > math.MaxFloat64），则结果为 (+Inf, Above) 或 (-Inf, Below)，具体取决于 x 的符号。

#### (*Float) Format <- go1.5

```go
func (x *Float) Format(s fmt.State, format rune)
```

Format implements fmt.Formatter. It accepts all the regular formats for floating-point numbers (‘b’, ’e’, ‘E’, ‘f’, ‘F’, ‘g’, ‘G’, ‘x’) as well as ‘p’ and ‘v’. See (*Float).Text for the interpretation of ‘p’. The ‘v’ format is handled like ‘g’. Format also supports specification of the minimum precision in digits, the output field width, as well as the format flags ‘+’ and ’ ’ for sign control, ‘0’ for space or zero padding, and ‘-’ for left or right justification. See the fmt package for details.

​	Format 实现 fmt.Formatter。它接受浮点数的所有常规格式（“b”、“e”、“E”、“f”、“F”、“g”、“G”、“x”）以及“p”和“v”。有关“p”的解释，请参阅 (*Float).Text。“v”格式的处理方式与“g”类似。Format 还支持指定最小精度（以数字表示）、输出字段宽度以及格式标志“+”和“ ”（用于控制符号）、“0”（用于空格或零填充）和“-”（用于左对齐或右对齐）。有关详细信息，请参阅 fmt 包。

#### (*Float) GobDecode <- go1.7

```go
func (z *Float) GobDecode(buf []byte) error
```

GobDecode implements the gob.GobDecoder interface. The result is rounded per the precision and rounding mode of z unless z’s precision is 0, in which case z is set exactly to the decoded value.

​	GobDecode 实现 gob.GobDecoder 接口。结果根据 z 的精度和舍入模式进行舍入，除非 z 的精度为 0，在这种情况下，z 将精确设置为解码值。

#### (*Float) GobEncode <- go1.7

```go
func (x *Float) GobEncode() ([]byte, error)
```

GobEncode implements the gob.GobEncoder interface. The Float value and all its attributes (precision, rounding mode, accuracy) are marshaled.

​	GobEncode 实现 gob.GobEncoder 接口。Float 值及其所有属性（精度、舍入模式、精度）都已编组。

#### (*Float) Int <- go1.5

```go
func (x *Float) Int(z *Int) (*Int, Accuracy)
```

Int returns the result of truncating x towards zero; or nil if x is an infinity. The result is Exact if x.IsInt(); otherwise it is Below for x > 0, and Above for x < 0. If a non-nil *Int argument z is provided, Int stores the result in z instead of allocating a new Int.

​	Int 返回将 x 向零截断的结果；如果 x 是无穷大，则返回 nil。如果 x.IsInt() 为 true，则结果为 Exact；否则，如果 x > 0，则结果为 Below，如果 x < 0，则结果为 Above。如果提供了非 nil 的 *Int 参数 z，则 Int 将结果存储在 z 中，而不是分配一个新的 Int。

#### (*Float) Int64 <- go1.5

```go
func (x *Float) Int64() (int64, Accuracy)
```

Int64 returns the integer resulting from truncating x towards zero. If math.MinInt64 <= x <= math.MaxInt64, the result is Exact if x is an integer, and Above (x < 0) or Below (x > 0) otherwise. The result is (math.MinInt64, Above) for x < math.MinInt64, and (math.MaxInt64, Below) for x > math.MaxInt64.

​	Int64 返回截断 x 朝向零的整数。如果 math.MinInt64 <= x <= math.MaxInt64，结果为 Exact（如果 x 是整数），否则为 Above（x < 0）或 Below（x > 0）。对于 x < math.MinInt64，结果为 (math.MinInt64, Above)；对于 x > math.MaxInt64，结果为 (math.MaxInt64, Below)。

#### (*Float) IsInf <- go1.5

```go
func (x *Float) IsInf() bool
```

IsInf reports whether x is +Inf or -Inf.

​	IsInf 报告 x 是否为 +Inf 或 -Inf。

#### (*Float) IsInt <- go1.5

```go
func (x *Float) IsInt() bool
```

IsInt reports whether x is an integer. ±Inf values are not integers.

​	IsInt 报告 x 是否为整数。±Inf 值不是整数。

#### (*Float) MantExp <- go1.5

```go
func (x *Float) MantExp(mant *Float) (exp int)
```

MantExp breaks x into its mantissa and exponent components and returns the exponent. If a non-nil mant argument is provided its value is set to the mantissa of x, with the same precision and rounding mode as x. The components satisfy x == mant × 2**exp, with 0.5 <= |mant| < 1.0. Calling MantExp with a nil argument is an efficient way to get the exponent of the receiver.

​	MantExp 将 x 分解为尾数和指数部分，并返回指数。如果提供了非 nil mant 参数，则将其值设置为 x 的尾数，精度和舍入模式与 x 相同。这些部分满足 x == mant × 2**exp，其中 0.5 <= |mant| < 1.0。使用 nil 参数调用 MantExp 是获取接收器指数的有效方法。

Special cases are:

​	特殊情况是：

```
(  ±0).MantExp(mant) = 0, with mant set to   ±0
(±Inf).MantExp(mant) = 0, with mant set to ±Inf
```

x and mant may be the same in which case x is set to its mantissa value.

​	x 和 mant 可能相同，在这种情况下，x 设置为其尾数值。

#### (*Float) MarshalText <- go1.6 

```go
func (x *Float) MarshalText() (text []byte, err error)
```

MarshalText implements the encoding.TextMarshaler interface. Only the Float value is marshaled (in full precision), other attributes such as precision or accuracy are ignored.

​	MarshalText 实现 encoding.TextMarshaler 接口。仅对 Float 值进行编组（以完全精度），忽略精度或准确性等其他属性。

#### (*Float) MinPrec <- go1.5 

```go
func (x *Float) MinPrec() uint
```

MinPrec returns the minimum precision required to represent x exactly (i.e., the smallest prec before x.SetPrec(prec) would start rounding x). The result is 0 for |x| == 0 and |x| == Inf.

​	MinPrec 返回表示 x 所需的最小精度（即，在 x.SetPrec(prec) 开始舍入 x 之前的最小 prec）。对于 |x| == 0 和 |x| == Inf，结果为 0。

#### (*Float) Mode <- go1.5

```go
func (x *Float) Mode() RoundingMode
```

Mode returns the rounding mode of x.

​	Mode 返回 x 的舍入模式。

#### (*Float) Mul <- go1.5

```go
func (z *Float) Mul(x, y *Float) *Float
```

Mul sets z to the rounded product x*y and returns z. Precision, rounding, and accuracy reporting are as for Add. Mul panics with ErrNaN if one operand is zero and the other operand an infinity. The value of z is undefined in that case.

​	Mul 将 z 设置为舍入的乘积 x*y，并返回 z。精度、舍入和准确性报告与 Add 相同。如果一个操作数为零，另一个操作数为无穷大，则 Mul 会引发 ErrNaN。在这种情况下，z 的值是未定义的。

#### (*Float) Neg <- go1.5 

```go
func (z *Float) Neg(x *Float) *Float
```

Neg sets z to the (possibly rounded) value of x with its sign negated, and returns z.

​	Neg 将 z 设置为 x 的（可能舍入的）值，其符号取反，并返回 z。

#### (*Float) Parse <- go1.5

```go
func (z *Float) Parse(s string, base int) (f *Float, b int, err error)
```

Parse parses s which must contain a text representation of a floating- point number with a mantissa in the given conversion base (the exponent is always a decimal number), or a string representing an infinite value.

​	Parse 解析 s，其中必须包含一个浮点数的文本表示形式，其尾数采用给定的转换基数（指数始终是十进制数），或一个表示无穷大值的字符串。

For base 0, an underscore character “_” may appear between a base prefix and an adjacent digit, and between successive digits; such underscores do not change the value of the number, or the returned digit count. Incorrect placement of underscores is reported as an error if there are no other errors. If base != 0, underscores are not recognized and thus terminate scanning like any other character that is not a valid radix point or digit.

​	对于基数 0，下划线字符“_”可能出现在基数前缀和相邻数字之间，以及连续数字之间；此类下划线不会更改数字的值或返回的数字计数。如果不存在其他错误，则将下划线放置不正确报告为错误。如果 base != 0，则不识别下划线，因此会像任何其他不是有效基点或数字的字符一样终止扫描。

It sets z to the (possibly rounded) value of the corresponding floating- point value, and returns z, the actual base b, and an error err, if any. The entire string (not just a prefix) must be consumed for success. If z’s precision is 0, it is changed to 64 before rounding takes effect. The number must be of the form:

​	它将 z 设置为相应浮点值（可能已舍入）的值，并返回 z、实际基数 b 和错误 err（如果有）。整个字符串（不仅仅是前缀）必须被使用才能成功。如果 z 的精度为 0，则在舍入生效之前将其更改为 64。数字必须采用以下形式：

```
number    = [ sign ] ( float | "inf" | "Inf" ) .
sign      = "+" | "-" .
float     = ( mantissa | prefix pmantissa ) [ exponent ] .
prefix    = "0" [ "b" | "B" | "o" | "O" | "x" | "X" ] .
mantissa  = digits "." [ digits ] | digits | "." digits .
pmantissa = [ "_" ] digits "." [ digits ] | [ "_" ] digits | "." digits .
exponent  = ( "e" | "E" | "p" | "P" ) [ sign ] digits .
digits    = digit { [ "_" ] digit } .
digit     = "0" ... "9" | "a" ... "z" | "A" ... "Z" .
```

The base argument must be 0, 2, 8, 10, or 16. Providing an invalid base argument will lead to a run-time panic.

​	基数参数必须为 0、2、8、10 或 16。提供无效的基数参数将导致运行时恐慌。

For base 0, the number prefix determines the actual base: A prefix of “0b” or “0B” selects base 2, “0o” or “0O” selects base 8, and “0x” or “0X” selects base 16. Otherwise, the actual base is 10 and no prefix is accepted. The octal prefix “0” is not supported (a leading “0” is simply considered a “0”).

​	对于基数 0，数字前缀决定实际基数：“0b”或“0B”前缀选择基数 2，“0o”或“0O”选择基数 8，“0x”或“0X”选择基数 16。否则，实际基数为 10 且不接受任何前缀。不支持八进制前缀“0”（前导“0”仅被视为“0”）。

A “p” or “P” exponent indicates a base 2 (rather then base 10) exponent; for instance, “0x1.fffffffffffffp1023” (using base 0) represents the maximum float64 value. For hexadecimal mantissae, the exponent character must be one of ‘p’ or ‘P’, if present (an “e” or “E” exponent indicator cannot be distinguished from a mantissa digit).

​	“p”或“P”指数表示基数 2（而不是基数 10）指数；例如，“0x1.fffffffffffffp1023”（使用基数 0）表示最大 float64 值。对于十六进制尾数，如果存在，指数字符必须是“p”或“P”之一（无法将“e”或“E”指数指示符与尾数数字区分开来）。

The returned *Float f is nil and the value of z is valid but not defined if an error is reported.

​	如果报告错误，则返回的 *Float f 为 nil，并且 z 的值有效但未定义。

#### (*Float) Prec <- go1.5

```go
func (x *Float) Prec() uint
```

Prec returns the mantissa precision of x in bits. The result may be 0 for |x| == 0 and |x| == Inf.

​	Prec 以位为单位返回 x 的尾数精度。对于 |x| == 0 和 |x| == Inf，结果可能为 0。

#### (*Float) Quo <- go1.5

```go
func (z *Float) Quo(x, y *Float) *Float
```

Quo sets z to the rounded quotient x/y and returns z. Precision, rounding, and accuracy reporting are as for Add. Quo panics with ErrNaN if both operands are zero or infinities. The value of z is undefined in that case.

​	Quo 将 z 设置为舍入的商 x/y 并返回 z。精度、舍入和准确性报告与 Add 相同。如果两个操作数都是零或无穷大，则 Quo 会引发 ErrNaN。在这种情况下，z 的值未定义。

#### (*Float) Rat <- go1.5

```go
func (x *Float) Rat(z *Rat) (*Rat, Accuracy)
```

Rat returns the rational number corresponding to x; or nil if x is an infinity. The result is Exact if x is not an Inf. If a non-nil *Rat argument z is provided, Rat stores the result in z instead of allocating a new Rat.

​	Rat 返回与 x 对应的有理数；如果 x 是无穷大，则返回 nil。如果 x 不是 Inf，则结果为 Exact。如果提供了非 nil 的 *Rat 参数 z，则 Rat 将结果存储在 z 中，而不是分配一个新的 Rat。

#### (*Float) Scan <- go1.8

```go
func (z *Float) Scan(s fmt.ScanState, ch rune) error
```

Scan is a support routine for fmt.Scanner; it sets z to the value of the scanned number. It accepts formats whose verbs are supported by fmt.Scan for floating point values, which are: ‘b’ (binary), ’e’, ‘E’, ‘f’, ‘F’, ‘g’ and ‘G’. Scan doesn’t handle ±Inf.

​	Scan 是 fmt.Scanner 的支持例程；它将 z 设置为扫描数字的值。它接受格式，其动词由 fmt.Scan 支持，用于浮点值，这些动词是：“b”（二进制）、“e”、“E”、“f”、“F”、“g”和“G”。Scan 不处理 ±Inf。

##### Scan Example

```go
package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implementation of fmt.Scanner.
	f := new(big.Float)
	_, err := fmt.Sscan("1.19282e99", f)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(f)
	}
}
Output:

1.19282e+99
```

#### (*Float) Set <- go1.5

```go
func (z *Float) Set(x *Float) *Float
```

Set sets z to the (possibly rounded) value of x and returns z. If z’s precision is 0, it is changed to the precision of x before setting z (and rounding will have no effect). Rounding is performed according to z’s precision and rounding mode; and z’s accuracy reports the result error relative to the exact (not rounded) result.

​	Set 将 z 设置为 x 的（可能已舍入）值并返回 z。如果 z 的精度为 0，则在设置 z 之前将其更改为 x 的精度（舍入不会产生任何影响）。舍入根据 z 的精度和舍入模式执行；z 的精度报告结果误差相对于确切（未舍入）结果。

#### (*Float) SetFloat64 <- go1.5

```go
func (z *Float) SetFloat64(x float64) *Float
```

SetFloat64 sets z to the (possibly rounded) value of x and returns z. If z’s precision is 0, it is changed to 53 (and rounding will have no effect). SetFloat64 panics with ErrNaN if x is a NaN.

​	SetFloat64 将 z 设置为 x 的（可能已舍入）值并返回 z。如果 z 的精度为 0，则将其更改为 53（舍入不会产生任何影响）。如果 x 是 NaN，则 SetFloat64 会引发 ErrNaN。

#### (*Float) SetInf <- go1.5

```go
func (z *Float) SetInf(signbit bool) *Float
```

SetInf sets z to the infinite Float -Inf if signbit is set, or +Inf if signbit is not set, and returns z. The precision of z is unchanged and the result is always Exact.

​	SetInf 将 z 设置为无限浮点数 -Inf（如果 signbit 设置）或 +Inf（如果 signbit 未设置），并返回 z。z 的精度保持不变，结果始终为 Exact。

#### (*Float) SetInt <- go1.5

```go
func (z *Float) SetInt(x *Int) *Float
```

SetInt sets z to the (possibly rounded) value of x and returns z. If z’s precision is 0, it is changed to the larger of x.BitLen() or 64 (and rounding will have no effect).

​	SetInt 将 z 设置为 x 的（可能已舍入）值，并返回 z。如果 z 的精度为 0，则将其更改为 x.BitLen() 或 64 中较大的一个（舍入不会产生任何影响）。

#### (*Float) SetInt64 <- go1.5

```go
func (z *Float) SetInt64(x int64) *Float
```

SetInt64 sets z to the (possibly rounded) value of x and returns z. If z’s precision is 0, it is changed to 64 (and rounding will have no effect).

​	SetInt64 将 z 设置为 x 的（可能已舍入）值，并返回 z。如果 z 的精度为 0，则将其更改为 64（舍入不会产生任何影响）。

#### (*Float) SetMantExp <- go1.5

```go
func (z *Float) SetMantExp(mant *Float, exp int) *Float
```

SetMantExp sets z to mant × 2**exp and returns z. The result z has the same precision and rounding mode as mant. SetMantExp is an inverse of MantExp but does not require 0.5 <= |mant| < 1.0. Specifically, for a given x of type *Float, SetMantExp relates to MantExp as follows:

​	SetMantExp 将 z 设置为 mant × 2**exp，并返回 z。结果 z 具有与 mant 相同的精度和舍入模式。SetMantExp 是 MantExp 的逆运算，但不要求 0.5 <= |mant| < 1.0。具体来说，对于给定类型的 *Float 的 x，SetMantExp 与 MantExp 的关系如下：

```
mant := new(Float)
new(Float).SetMantExp(mant, x.MantExp(mant)).Cmp(x) == 0
```

Special cases are:

​	特殊情况是：

```
z.SetMantExp(  ±0, exp) =   ±0
z.SetMantExp(±Inf, exp) = ±Inf
```

z and mant may be the same in which case z’s exponent is set to exp.

​	z 和 mant 可能相同，在这种情况下，z 的指数设置为 exp。

#### (*Float) SetMode <- go1.5

```go
func (z *Float) SetMode(mode RoundingMode) *Float
```

SetMode sets z’s rounding mode to mode and returns an exact z. z remains unchanged otherwise. z.SetMode(z.Mode()) is a cheap way to set z’s accuracy to Exact.

​	SetMode 将 z 的舍入模式设置为 mode 并返回一个精确的 z。否则 z 保持不变。z.SetMode(z.Mode()) 是将 z 的精度设置为 Exact 的一种廉价方式。

#### (*Float) SetPrec <- go1.5

```go
func (z *Float) SetPrec(prec uint) *Float
```

SetPrec sets z’s precision to prec and returns the (possibly) rounded value of z. Rounding occurs according to z’s rounding mode if the mantissa cannot be represented in prec bits without loss of precision. SetPrec(0) maps all finite values to ±0; infinite values remain unchanged. If prec > MaxPrec, it is set to MaxPrec.

​	SetPrec 将 z 的精度设置为 prec 并返回 z 的（可能）舍入值。如果无法在不损失精度的情况下用 prec 位表示尾数，则根据 z 的舍入模式进行舍入。SetPrec(0) 将所有有限值映射到 ±0；无限值保持不变。如果 prec > MaxPrec，则将其设置为 MaxPrec。

#### (*Float) SetRat <- go1.5

```go
func (z *Float) SetRat(x *Rat) *Float
```

SetRat sets z to the (possibly rounded) value of x and returns z. If z’s precision is 0, it is changed to the largest of a.BitLen(), b.BitLen(), or 64; with x = a/b.

​	SetRat 将 z 设置为 x 的（可能舍入的）值并返回 z。如果 z 的精度为 0，则将其更改为 a.BitLen()、b.BitLen() 或 64 中的最大值；其中 x = a/b。

#### (*Float) SetString <- go1.5

```go
func (z *Float) SetString(s string) (*Float, bool)
```

SetString sets z to the value of s and returns z and a boolean indicating success. s must be a floating-point number of the same format as accepted by Parse, with base argument 0. The entire string (not just a prefix) must be valid for success. If the operation failed, the value of z is undefined but the returned value is nil.

​	SetString 将 z 设置为 s 的值并返回 z 和一个布尔值来指示成功。s 必须是与 Parse 接受的相同格式的浮点数，基本参数为 0。整个字符串（不仅仅是前缀）必须有效才能成功。如果操作失败，z 的值未定义，但返回值为 nil。

##### SetString Example

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	f := new(big.Float)
	f.SetString("3.14159")
	fmt.Println(f)
}
Output:

3.14159
```

#### (*Float) SetUint64 <- go1.5

```go
func (z *Float) SetUint64(x uint64) *Float
```

SetUint64 sets z to the (possibly rounded) value of x and returns z. If z’s precision is 0, it is changed to 64 (and rounding will have no effect).

​	SetUint64 将 z 设置为 x 的（可能已舍入）值并返回 z。如果 z 的精度为 0，则将其更改为 64（舍入不会产生任何影响）。

#### (*Float) Sign <- go1.5

```go
func (x *Float) Sign() int
```

Sign returns:

​	Sign 返回：

```
-1 if x <   0
 0 if x is ±0
+1 if x >   0
```

#### (*Float) Signbit <- go1.5

```go
func (x *Float) Signbit() bool
```

Signbit reports whether x is negative or negative zero.

​	Signbit 报告 x 是否为负数或负零。

#### (*Float) Sqrt <- go1.10

```go
func (z *Float) Sqrt(x *Float) *Float
```

Sqrt sets z to the rounded square root of x, and returns it.

​	Sqrt 将 z 设置为 x 的舍入平方根，并返回它。

If z’s precision is 0, it is changed to x’s precision before the operation. Rounding is performed according to z’s precision and rounding mode, but z’s accuracy is not computed. Specifically, the result of z.Acc() is undefined.

​	如果 z 的精度为 0，则在执行操作之前将其更改为 x 的精度。舍入根据 z 的精度和舍入模式执行，但不计算 z 的精度。具体来说，z.Acc() 的结果是未定义的。

The function panics if z < 0. The value of z is undefined in that case.

​	如果 z < 0，则该函数会引发 panic。在这种情况下，z 的值是未定义的。

#### (*Float) String <- go1.5

```go
func (x *Float) String() string
```

String formats x like x.Text(‘g’, 10). (String must be called explicitly, Float.Format does not support %s verb.)

​	字符串格式 x 如 x.Text(‘g’, 10)。（必须显式调用 String，Float.Format 不支持 %s 动词。）

#### (*Float) Sub <- go1.5

```go
func (z *Float) Sub(x, y *Float) *Float
```

Sub sets z to the rounded difference x-y and returns z. Precision, rounding, and accuracy reporting are as for Add. Sub panics with ErrNaN if x and y are infinities with equal signs. The value of z is undefined in that case.

​	Sub 将 z 设置为舍入的差值 x-y 并返回 z。精度、舍入和准确性报告与 Add 相同。如果 x 和 y 是具有相同符号的无穷大，Sub 会引发 ErrNaN。在这种情况下，z 的值是未定义的。

#### (*Float) Text <- go1.5

```go
func (x *Float) Text(format byte, prec int) string
```

Text converts the floating-point number x to a string according to the given format and precision prec. The format is one of:

​	Text 根据给定的格式和精度 prec 将浮点数 x 转换为字符串。格式之一：

```
'e'	-d.dddde±dd, decimal exponent, at least two (possibly 0) exponent digits
'E'	-d.ddddE±dd, decimal exponent, at least two (possibly 0) exponent digits
'f'	-ddddd.dddd, no exponent
'g'	like 'e' for large exponents, like 'f' otherwise
'G'	like 'E' for large exponents, like 'f' otherwise
'x'	-0xd.dddddp±dd, hexadecimal mantissa, decimal power of two exponent
'p'	-0x.dddp±dd, hexadecimal mantissa, decimal power of two exponent (non-standard)
'b'	-ddddddp±dd, decimal mantissa, decimal power of two exponent (non-standard)
```

For the power-of-two exponent formats, the mantissa is printed in normalized form:

​	对于以 2 为底的指数格式，尾数以标准化形式打印：

```
'x'	hexadecimal mantissa in [1, 2), or 0
'p'	hexadecimal mantissa in [½, 1), or 0
'b'	decimal integer mantissa using x.Prec() bits, or 0
```

Note that the ‘x’ form is the one used by most other languages and libraries.

​	请注意，“x”格式是大多数其他语言和库使用的格式。

If format is a different character, Text returns a “%” followed by the unrecognized format character.

​	如果格式是其他字符，Text 会返回一个“%”，后跟无法识别的格式字符。

The precision prec controls the number of digits (excluding the exponent) printed by the ’e’, ‘E’, ‘f’, ‘g’, ‘G’, and ‘x’ formats. For ’e’, ‘E’, ‘f’, and ‘x’, it is the number of digits after the decimal point. For ‘g’ and ‘G’ it is the total number of digits. A negative precision selects the smallest number of decimal digits necessary to identify the value x uniquely using x.Prec() mantissa bits. The prec value is ignored for the ‘b’ and ‘p’ formats.

​	精度 prec 控制由“e”、“E”、“f”、“g”、“G”和“x”格式打印的数字位数（不包括指数）。对于“e”、“E”、“f”和“x”，它是小数点后的数字位数。对于“g”和“G”，它是数字的总位数。负精度选择唯一标识值 x 所需的最小小数位数，使用 x.Prec() 尾数位。对于“b”和“p”格式，忽略 prec 值。

#### (*Float) Uint64 <- go1.5

```go
func (x *Float) Uint64() (uint64, Accuracy)
```

Uint64 returns the unsigned integer resulting from truncating x towards zero. If 0 <= x <= math.MaxUint64, the result is Exact if x is an integer and Below otherwise. The result is (0, Above) for x < 0, and (math.MaxUint64, Below) for x > math.MaxUint64.

​	Uint64 返回截断 x 朝向零产生的无符号整数。如果 0 <= x <= math.MaxUint64，则结果为 Exact（如果 x 是整数）或 Below（否则）。对于 x < 0，结果为 (0, Above)；对于 x > math.MaxUint64，结果为 (math.MaxUint64, Below)。

#### (*Float) UnmarshalText <- go1.6

```go
func (z *Float) UnmarshalText(text []byte) error
```

UnmarshalText implements the encoding.TextUnmarshaler interface. The result is rounded per the precision and rounding mode of z. If z’s precision is 0, it is changed to 64 before rounding takes effect.

​	UnmarshalText 实现 encoding.TextUnmarshaler 接口。结果根据 z 的精度和舍入模式进行舍入。如果 z 的精度为 0，则在舍入生效之前将其更改为 64。

### type Int

```go
type Int struct {
	// contains filtered or unexported fields
}
```

An Int represents a signed multi-precision integer. The zero value for an Int represents the value 0.

​	Int 表示带符号多精度整数。Int 的零值表示值 0。

Operations always take pointer arguments (*Int) rather than Int values, and each unique Int value requires its own unique *Int pointer. To “copy” an Int value, an existing (or newly allocated) Int must be set to a new value using the Int.Set method; shallow copies of Ints are not supported and may lead to errors.

​	操作总是采用指针参数 (*Int) 而不是 Int 值，并且每个唯一的 Int 值都需要其自己的唯一 *Int 指针。要“复制”Int 值，必须使用 Int.Set 方法将现有（或新分配的）Int 设置为新值；不支持 Int 的浅层副本，并且可能导致错误。

#### func NewInt

```go
func NewInt(x int64) *Int
```

NewInt allocates and returns a new Int set to x.

​	NewInt 分配并返回一个设置为 x 的新 Int。

#### (*Int) Abs

```go
func (z *Int) Abs(x *Int) *Int
```

Abs sets z to |x| (the absolute value of x) and returns z.

​	Abs 将 z 设置为 |x|（x 的绝对值）并返回 z。

#### (*Int) Add

```go
func (z *Int) Add(x, y *Int) *Int
```

Add sets z to the sum x+y and returns z.

​	Add 将 z 设置为 x+y 的和并返回 z。

#### (*Int) And

```go
func (z *Int) And(x, y *Int) *Int
```

And sets z = x & y and returns z.

​	And 将 z 设置为 x & y 并返回 z。

#### (*Int) AndNot

```go
func (z *Int) AndNot(x, y *Int) *Int
```

AndNot sets z = x &^ y and returns z.

​	AndNot 设置 z = x &^ y 并返回 z。

#### (*Int) Append <- go1.6

```go
func (x *Int) Append(buf []byte, base int) []byte
```

Append appends the string representation of x, as generated by x.Text(base), to buf and returns the extended buffer.

​	Append 将 x 的字符串表示形式（由 x.Text(base) 生成）追加到 buf 并返回扩展的缓冲区。

#### (*Int) Binomial

```go
func (z *Int) Binomial(n, k int64) *Int
```

Binomial sets z to the binomial coefficient C(n, k) and returns z.

​	Binomial 将 z 设置为二项式系数 C(n, k) 并返回 z。

#### (*Int) Bit

```go
func (x *Int) Bit(i int) uint
```

Bit returns the value of the i’th bit of x. That is, it returns (x»i)&1. The bit index i must be >= 0.

​	Bit 返回 x 的第 i 位的数值。也就是说，它返回 (x»i)&1。位索引 i 必须 >= 0。

#### (*Int) BitLen

```go
func (x *Int) BitLen() int
```

BitLen returns the length of the absolute value of x in bits. The bit length of 0 is 0.

​	BitLen 返回 x 的绝对值在位数中的长度。0 的位长为 0。

#### (*Int) Bits

```go
func (x *Int) Bits() []Word
```

Bits provides raw (unchecked but fast) access to x by returning its absolute value as a little-endian Word slice. The result and x share the same underlying array. Bits is intended to support implementation of missing low-level Int functionality outside this package; it should be avoided otherwise.

​	Bits 通过返回其绝对值作为小端 Word 切片，提供对 x 的原始（未经检查但速度快）访问。结果和 x 共享相同的底层数组。Bits 旨在支持在此包之外实现缺失的低级 Int 功能；否则应避免使用它。

#### (*Int) Bytes

```go
func (x *Int) Bytes() []byte
```

Bytes returns the absolute value of x as a big-endian byte slice.

​	Bytes 将 x 的绝对值作为大端字节切片返回。

To use a fixed length slice, or a preallocated one, use FillBytes.

​	要使用固定长度的切片或预先分配的切片，请使用 FillBytes。

#### (*Int) Cmp

```go
func (x *Int) Cmp(y *Int) (r int)
```

Cmp compares x and y and returns:

​	Cmp 比较 x 和 y 并返回：

```
-1 if x <  y
 0 if x == y
+1 if x >  y
```

#### (*Int) CmpAbs <- go1.10

```go
func (x *Int) CmpAbs(y *Int) int
```

CmpAbs compares the absolute values of x and y and returns:

​	CmpAbs 比较 x 和 y 的绝对值并返回：

```
-1 if |x| <  |y|
 0 if |x| == |y|
+1 if |x| >  |y|
```

#### (*Int) Div

```go
func (z *Int) Div(x, y *Int) *Int
```

Div sets z to the quotient x/y for y != 0 and returns z. If y == 0, a division-by-zero run-time panic occurs. Div implements Euclidean division (unlike Go); see DivMod for more details.

​	Div 将 z 设置为 x/y 的商，其中 y != 0，并返回 z。如果 y == 0，则会发生除以零的运行时恐慌。Div 实现欧几里得除法（与 Go 不同）；有关更多详细信息，请参阅 DivMod。

#### (*Int) DivMod

```go
func (z *Int) DivMod(x, y, m *Int) (*Int, *Int)
```

DivMod sets z to the quotient x div y and m to the modulus x mod y and returns the pair (z, m) for y != 0. If y == 0, a division-by-zero run-time panic occurs.

​	DivMod 将 z 设置为商 x div y，将 m 设置为模 x mod y，并返回对 (z, m)（y != 0）。如果 y == 0，则会发生除以零的运行时恐慌。

DivMod implements Euclidean division and modulus (unlike Go):

​	DivMod 实现欧几里得除法和模（与 Go 不同）：

```
q = x div y  such that
m = x - y*q  with 0 <= m < |y|
```

(See Raymond T. Boute, “The Euclidean definition of the functions div and mod”. ACM Transactions on Programming Languages and Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992. ACM press.) See QuoRem for T-division and modulus (like Go).

​	（请参阅 Raymond T. Boute，“div 和 mod 函数的欧几里得定义”。ACM 编程语言和系统事务 (TOPLAS)，14(2):127-144，纽约，纽约，美国，4/1992。ACM 出版社。）有关 T 除法和模（如 Go），请参阅 QuoRem。

#### (*Int) Exp

```go
func (z *Int) Exp(x, y, m *Int) *Int
```

Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z. If m == nil or m == 0, z = x**y unless y <= 0 then z = 1. If m != 0, y < 0, and x and m are not relatively prime, z is unchanged and nil is returned.

​	Exp 将 z 设置为 xy mod |m|（即忽略 m 的符号），并返回 z。如果 m == nil 或 m == 0，则 z = xy，除非 y <= 0，则 z = 1。如果 m != 0、y < 0，并且 x 和 m 不是相对质数，则 z 保持不变并返回 nil。

Modular exponentiation of inputs of a particular size is not a cryptographically constant-time operation.

​	特定大小的输入的模幂不是密码学恒定时间操作。

#### (*Int) FillBytes <- go1.15

```go
func (x *Int) FillBytes(buf []byte) []byte
```

FillBytes sets buf to the absolute value of x, storing it as a zero-extended big-endian byte slice, and returns buf.

​	FillBytes 将 buf 设置为 x 的绝对值，将其存储为零扩展的大端字节切片，并返回 buf。

If the absolute value of x doesn’t fit in buf, FillBytes will panic.

​	如果 x 的绝对值不适合 buf，FillBytes 将引发恐慌。

#### (*Int) Float64 <-go1.21.0

```go
func (x *Int) Float64() (float64, Accuracy)
```

Float64 returns the float64 value nearest x, and an indication of any rounding that occurred.

​	Float64 返回最接近 x 的 float64 值，以及发生的任何舍入的指示。

#### (*Int) Format

```go
func (x *Int) Format(s fmt.State, ch rune)
```

Format implements fmt.Formatter. It accepts the formats ‘b’ (binary), ‘o’ (octal with 0 prefix), ‘O’ (octal with 0o prefix), ’d’ (decimal), ‘x’ (lowercase hexadecimal), and ‘X’ (uppercase hexadecimal). Also supported are the full suite of package fmt’s format flags for integral types, including ‘+’ and ’ ’ for sign control, ‘#’ for leading zero in octal and for hexadecimal, a leading “0x” or “0X” for “%#x” and “%#X” respectively, specification of minimum digits precision, output field width, space or zero padding, and ‘-’ for left or right justification.

​	Format 实现 fmt.Formatter。它接受格式“b”（二进制）、“o”（带 0 前缀的八进制）、“O”（带 0o 前缀的八进制）、“d”（十进制）、“x”（小写十六进制）和“X”（大写十六进制）。还支持包 fmt 的一整套格式标志，用于整数类型，包括“+”和“ ”用于控制符号，“#”用于八进制和十六进制中的前导零，“0x”或“0X”分别用于“%#x”和“%#X”，指定最小数字精度、输出字段宽度、空格或零填充，以及“ -”用于左或右对齐。

#### (*Int) GCD 

```go
func (z *Int) GCD(x, y, a, b *Int) *Int
```

GCD sets z to the greatest common divisor of a and b and returns z. If x or y are not nil, GCD sets their value such that z = a*x + b*y.

​	GCD 将 z 设置为 a 和 b 的最大公约数，并返回 z。如果 x 或 y 不为 nil，GCD 将设置它们的值，使得 z = ax + by。

a and b may be positive, zero or negative. (Before Go 1.14 both had to be > 0.) Regardless of the signs of a and b, z is always >= 0.

​	a 和 b 可以是正数、零或负数。（在 Go 1.14 之前，两者都必须 > 0。）无论 a 和 b 的符号如何，z 始终 >= 0。

If a == b == 0, GCD sets z = x = y = 0.

​	如果 a == b == 0，GCD 将 z = x = y = 0。

If a == 0 and b != 0, GCD sets z = |b|, x = 0, y = sign(b) * 1.

​	如果 a == 0 且 b != 0，GCD 将 z = |b|，x = 0，y = sign(b) * 1。

If a != 0 and b == 0, GCD sets z = |a|, x = sign(a) * 1, y = 0.

​	如果 a != 0 且 b == 0，GCD 将 z = |a|，x = sign(a) * 1，y = 0。

#### (*Int) GobDecode

```go
func (z *Int) GobDecode(buf []byte) error
```

GobDecode implements the gob.GobDecoder interface.

​	GobDecode 实现 gob.GobDecoder 接口。

#### (*Int) GobEncode

```go
func (x *Int) GobEncode() ([]byte, error)
```

GobEncode implements the gob.GobEncoder interface.

​	GobEncode 实现 gob.GobEncoder 接口。

#### (*Int) Int64

```go
func (x *Int) Int64() int64
```

Int64 returns the int64 representation of x. If x cannot be represented in an int64, the result is undefined.

​	Int64 返回 x 的 int64 表示形式。如果 x 无法用 int64 表示，则结果未定义。

#### (*Int) IsInt64 <- go1.9

```go
func (x *Int) IsInt64() bool
```

IsInt64 reports whether x can be represented as an int64.

​	IsInt64 报告 x 是否可以表示为 int64。

#### (*Int) IsUint64 <- go1.9

```go
func (x *Int) IsUint64() bool
```

IsUint64 reports whether x can be represented as a uint64.

​	IsUint64 报告 x 是否可以表示为 uint64。

#### (*Int) Lsh 

```go
func (z *Int) Lsh(x *Int, n uint) *Int
```

Lsh sets z = x « n and returns z.

​	Lsh 将 z 设置为 x « n 并返回 z。

#### (*Int) MarshalJSON <- go1.1 

```go
func (x *Int) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the json.Marshaler interface.

​	MarshalJSON 实现 json.Marshaler 接口。

#### (*Int) MarshalText <- go1.3

```go
func (x *Int) MarshalText() (text []byte, err error)
```

MarshalText implements the encoding.TextMarshaler interface.

​	MarshalText 实现 encoding.TextMarshaler 接口。

#### (*Int) Mod

```go
func (z *Int) Mod(x, y *Int) *Int
```

Mod sets z to the modulus x%y for y != 0 and returns z. If y == 0, a division-by-zero run-time panic occurs. Mod implements Euclidean modulus (unlike Go); see DivMod for more details.

​	Mod 将 z 设置为模数 x%y（y != 0），并返回 z。如果 y == 0，则会发生除以零的运行时恐慌。Mod 实现欧几里得模数（与 Go 不同）；有关更多详细信息，请参阅 DivMod。

#### (*Int) ModInverse

```go
func (z *Int) ModInverse(g, n *Int) *Int
```

ModInverse sets z to the multiplicative inverse of g in the ring ℤ/nℤ and returns z. If g and n are not relatively prime, g has no multiplicative inverse in the ring ℤ/nℤ. In this case, z is unchanged and the return value is nil. If n == 0, a division-by-zero run-time panic occurs.

​	ModInverse 将 z 设置为环 ℤ/nℤ 中 g 的乘法逆，并返回 z。如果 g 和 n 不是相对质数，则 g 在环 ℤ/nℤ 中没有乘法逆。在这种情况下，z 保持不变，返回值为 nil。如果 n == 0，则会发生除以零的运行时恐慌。

#### (*Int) ModSqrt <- go1.5 

```go
func (z *Int) ModSqrt(x, p *Int) *Int
```

ModSqrt sets z to a square root of x mod p if such a square root exists, and returns z. The modulus p must be an odd prime. If x is not a square mod p, ModSqrt leaves z unchanged and returns nil. This function panics if p is not an odd integer, its behavior is undefined if p is odd but not prime.

​	如果存在这样的平方根，ModSqrt 将 z 设置为 x mod p 的平方根，并返回 z。模数 p 必须是奇素数。如果 x 不是平方模数 p，ModSqrt 将 z 保持不变并返回 nil。如果 p 不是奇数，此函数会引发 panic，如果 p 是奇数但不是素数，则其行为未定义。

#### (*Int) Mul

```go
func (z *Int) Mul(x, y *Int) *Int
```

Mul sets z to the product x*y and returns z.

​	Mul 将 z 设置为 x*y 的乘积并返回 z。

#### (*Int) MulRange

```go
func (z *Int) MulRange(a, b int64) *Int
```

MulRange sets z to the product of all integers in the range [a, b] inclusively and returns z. If a > b (empty range), the result is 1.

​	MulRange 将 z 设置为 [a, b] 范围（包括 a 和 b）内所有整数的乘积并返回 z。如果 a > b（空范围），则结果为 1。

#### (*Int) Neg 

```go
func (z *Int) Neg(x *Int) *Int
```

Neg sets z to -x and returns z.

​	Neg 将 z 设置为 -x 并返回 z。

#### (*Int) Not 

```go
func (z *Int) Not(x *Int) *Int
```

Not sets z = ^x and returns z.

​	Not 将 z 设置为 ^x 并返回 z。

#### (*Int) Or

```go
func (z *Int) Or(x, y *Int) *Int
```

Or sets z = x | y and returns z.

​	或将 z 设置为 x | y 并返回 z。

#### (*Int) ProbablyPrime 

```go
func (x *Int) ProbablyPrime(n int) bool
```

ProbablyPrime reports whether x is probably prime, applying the Miller-Rabin test with n pseudorandomly chosen bases as well as a Baillie-PSW test.

​	ProbablyPrime 报告 x 是否可能是素数，应用 Miller-Rabin 测试，其中 n 个伪随机选择的基数以及 Baillie-PSW 测试。

If x is prime, ProbablyPrime returns true. If x is chosen randomly and not prime, ProbablyPrime probably returns false. The probability of returning true for a randomly chosen non-prime is at most ¼ⁿ.

​	如果 x 是素数，ProbablyPrime 返回 true。如果 x 是随机选择的且不是素数，ProbablyPrime 可能返回 false。对于随机选择的非素数，返回 true 的概率最多为 ¼ⁿ。

ProbablyPrime is 100% accurate for inputs less than 2⁶⁴. See Menezes et al., Handbook of Applied Cryptography, 1997, pp. 145-149, and FIPS 186-4 Appendix F for further discussion of the error probabilities.

​	ProbablyPrime 对于小于 2⁶⁴ 的输入是 100% 准确的。有关错误概率的进一步讨论，请参阅 Menezes 等人，应用密码学手册，1997 年，第 145-149 页，以及 FIPS 186-4 附录 F。

ProbablyPrime is not suitable for judging primes that an adversary may have crafted to fool the test.

​	ProbablyPrime 不适合判断对手可能精心设计以欺骗测试的素数。

As of Go 1.8, ProbablyPrime(0) is allowed and applies only a Baillie-PSW test. Before Go 1.8, ProbablyPrime applied only the Miller-Rabin tests, and ProbablyPrime(0) panicked.

​	从 Go 1.8 开始，允许 ProbablyPrime(0) 并且仅应用 Baillie-PSW 测试。在 Go 1.8 之前，ProbablyPrime 仅应用 Miller-Rabin 测试，并且 ProbablyPrime(0) 会引发 panic。

#### (*Int) Quo

```go
func (z *Int) Quo(x, y *Int) *Int
```

Quo sets z to the quotient x/y for y != 0 and returns z. If y == 0, a division-by-zero run-time panic occurs. Quo implements truncated division (like Go); see QuoRem for more details.

​	Quo 将 z 设置为 y != 0 的商 x/y 并返回 z。如果 y == 0，则会发生除以零的运行时 panic。Quo 实现截断除法（如 Go）；有关更多详细信息，请参阅 QuoRem。

#### (*Int) QuoRem

```go
func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int)
```

QuoRem sets z to the quotient x/y and r to the remainder x%y and returns the pair (z, r) for y != 0. If y == 0, a division-by-zero run-time panic occurs.

​	QuoRem 将 z 设置为商 x/y，将 r 设置为余数 x%y，并返回对 (z, r)（y != 0）。如果 y == 0，则会发生除以零的运行时恐慌。

QuoRem implements T-division and modulus (like Go):

​	QuoRem 实现 T 除法和模数（如 Go）：

```
q = x/y      with the result truncated to zero
r = x - y*q
```

(See Daan Leijen, “Division and Modulus for Computer Scientists”.) See DivMod for Euclidean division and modulus (unlike Go).

​	（请参阅 Daan Leijen，“计算机科学家的除法和模数”。）有关欧几里得除法和模数（与 Go 不同），请参阅 DivMod。

#### (*Int) Rand

```go
func (z *Int) Rand(rnd *rand.Rand, n *Int) *Int
```

Rand sets z to a pseudo-random number in [0, n) and returns z.

​	Rand 将 z 设置为 [0, n) 中的伪随机数并返回 z。

As this uses the math/rand package, it must not be used for security-sensitive work. Use crypto/rand.Int instead.

​	由于这使用了 math/rand 包，因此不得将其用于对安全敏感的工作。请改用 crypto/rand.Int。

#### (*Int) Rem

```go
func (z *Int) Rem(x, y *Int) *Int
```

Rem sets z to the remainder x%y for y != 0 and returns z. If y == 0, a division-by-zero run-time panic occurs. Rem implements truncated modulus (like Go); see QuoRem for more details.

​	Rem 将 z 设置为余数 x%y（y != 0）并返回 z。如果 y == 0，则会发生除以零的运行时恐慌。Rem 实现截断模数（如 Go）；有关更多详细信息，请参阅 QuoRem。

#### (*Int) Rsh

```go
func (z *Int) Rsh(x *Int, n uint) *Int
```

Rsh sets z = x » n and returns z.

​	Rsh 设置 z = x » n 并返回 z。

#### (*Int) Scan

```go
func (z *Int) Scan(s fmt.ScanState, ch rune) error
```

Scan is a support routine for fmt.Scanner; it sets z to the value of the scanned number. It accepts the formats ‘b’ (binary), ‘o’ (octal), ’d’ (decimal), ‘x’ (lowercase hexadecimal), and ‘X’ (uppercase hexadecimal).

​	Scan 是 fmt.Scanner 的支持例程；它将 z 设置为扫描数字的值。它接受格式“b”（二进制）、“o”（八进制）、“d”（十进制）、“x”（小写十六进制）和“X”（大写十六进制）。

##### Scan Example

```go
package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implementation of fmt.Scanner.
	i := new(big.Int)
	_, err := fmt.Sscan("18446744073709551617", i)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(i)
	}
}
```

#### (*Int) Set

```go
func (z *Int) Set(x *Int) *Int
```

Set sets z to x and returns z.

​	Set 将 z 设置为 x 并返回 z。

#### (*Int) SetBit

```go
func (z *Int) SetBit(x *Int, i int, b uint) *Int
```

SetBit sets z to x, with x’s i’th bit set to b (0 or 1). That is, if b is 1 SetBit sets z = x | (1 « i); if b is 0 SetBit sets z = x &^ (1 « i). If b is not 0 or 1, SetBit will panic.

​	SetBit 将 z 设置为 x，并将 x 的第 i 位比特设置为 b（0 或 1）。也就是说，如果 b 为 1，SetBit 将 z 设置为 x | (1 « i)；如果 b 为 0，SetBit 将 z 设置为 x &^ (1 « i)。如果 b 不是 0 或 1，SetBit 将会引发恐慌。

#### (*Int) SetBits 

```go
func (z *Int) SetBits(abs []Word) *Int
```

SetBits provides raw (unchecked but fast) access to z by setting its value to abs, interpreted as a little-endian Word slice, and returning z. The result and abs share the same underlying array. SetBits is intended to support implementation of missing low-level Int functionality outside this package; it should be avoided otherwise.

​	SetBits 通过将 z 的值设置为 abs（解释为 little-endian Word 片段）并返回 z 来提供对 z 的原始（未经检查但速度很快）访问。结果和 abs 共享相同的底层数组。SetBits 旨在支持在此包之外实现缺失的低级 Int 功能；否则应避免使用它。

#### (*Int) SetBytes 

```go
func (z *Int) SetBytes(buf []byte) *Int
```

SetBytes interprets buf as the bytes of a big-endian unsigned integer, sets z to that value, and returns z.

​	SetBytes 将 buf 解释为大端无符号整数的字节，将 z 设置为该值，并返回 z。

#### (*Int) SetInt64 

```go
func (z *Int) SetInt64(x int64) *Int
```

SetInt64 sets z to x and returns z.

​	SetInt64 将 z 设置为 x 并返回 z。

#### (*Int) SetString 

```go
func (z *Int) SetString(s string, base int) (*Int, bool)
```

SetString sets z to the value of s, interpreted in the given base, and returns z and a boolean indicating success. The entire string (not just a prefix) must be valid for success. If SetString fails, the value of z is undefined but the returned value is nil.

​	SetString 将 z 设置为 s 的值，以给定的基数解释，并返回 z 和一个布尔值来指示成功。整个字符串（不仅仅是前缀）必须有效才能成功。如果 SetString 失败，则 z 的值未定义，但返回值为 nil。

The base argument must be 0 or a value between 2 and MaxBase. For base 0, the number prefix determines the actual base: A prefix of “0b” or “0B” selects base 2, “0”, “0o” or “0O” selects base 8, and “0x” or “0X” selects base 16. Otherwise, the selected base is 10 and no prefix is accepted.

​	base 参数必须为 0 或 2 到 MaxBase 之间的值。对于基数 0，数字前缀确定实际基数：“0b”或“0B”前缀选择基数 2，“0”、“0o”或“0O”选择基数 8，“0x”或“0X”选择基数 16。否则，选定的基数为 10，不接受前缀。

For bases <= 36, lower and upper case letters are considered the same: The letters ‘a’ to ‘z’ and ‘A’ to ‘Z’ represent digit values 10 to 35. For bases > 36, the upper case letters ‘A’ to ‘Z’ represent the digit values 36 to 61.

​	对于基数 <= 36，大小写字母被视为相同：字母“a”到“z”和“A”到“Z”表示数字值 10 到 35。对于基数 > 36，大写字母“A”到“Z”表示数字值 36 到 61。

For base 0, an underscore character “_” may appear between a base prefix and an adjacent digit, and between successive digits; such underscores do not change the value of the number. Incorrect placement of underscores is reported as an error if there are no other errors. If base != 0, underscores are not recognized and act like any other character that is not a valid digit.

​	对于基数 0，下划线字符“_”可能出现在基数前缀和相邻数字之间，以及连续数字之间；此类下划线不会改变数字的值。如果不存在其他错误，则将下划线放置不正确报告为错误。如果基数 != 0，则不识别下划线，并且它们的行为与任何其他不是有效数字的字符一样。

##### SetString Example

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := new(big.Int)
	i.SetString("644", 8) // octal
	fmt.Println(i)
}
Output:

420
```

#### (*Int) SetUint64 <- go1.1

```go
func (z *Int) SetUint64(x uint64) *Int
```

SetUint64 sets z to x and returns z.

​	SetUint64 将 z 设置为 x 并返回 z。

#### (*Int) Sign

```go
func (x *Int) Sign() int
```

Sign returns:

​	Sign 返回：

```
-1 if x <  0
 0 if x == 0
+1 if x >  0
```

#### (*Int) Sqrt <- go1.8

```go
func (z *Int) Sqrt(x *Int) *Int
```

Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z. It panics if x is negative.

​	Sqrt 将 z 设置为 ⌊√x⌋，即满足 z² ≤ x 的最大整数，并返回 z。如果 x 为负数，则会引发 panic。

#### (*Int) String

```go
func (x *Int) String() string
```

String returns the decimal representation of x as generated by x.Text(10).

​	String 返回 x.Text(10) 生成的 x 的十进制表示形式。

#### (*Int) Sub

```go
func (z *Int) Sub(x, y *Int) *Int
```

Sub sets z to the difference x-y and returns z.

​	Sub 将 z 设置为 x-y 的差值并返回 z。

#### (*Int) Text <- go1.6

```go
func (x *Int) Text(base int) string
```

Text returns the string representation of x in the given base. Base must be between 2 and 62, inclusive. The result uses the lower-case letters ‘a’ to ‘z’ for digit values 10 to 35, and the upper-case letters ‘A’ to ‘Z’ for digit values 36 to 61. No prefix (such as “0x”) is added to the string. If x is a nil pointer it returns “”.

​	Text 返回 x 在给定基数中的字符串表示形式。基数必须在 2 到 62（包括 2 和 62）之间。结果使用小写字母“a”到“z”表示数字值 10 到 35，使用大写字母“A”到“Z”表示数字值 36 到 61。字符串中不添加任何前缀（例如“0x”）。如果 x 是 nil 指针，则返回“”。

#### (*Int) TrailingZeroBits <- go1.13

```go
func (x *Int) TrailingZeroBits() uint
```

TrailingZeroBits returns the number of consecutive least significant zero bits of |x|.

​	TrailingZeroBits 返回 |x| 的连续最低有效零位数。

#### (*Int) Uint64 <- go1.1

```go
func (x *Int) Uint64() uint64
```

Uint64 returns the uint64 representation of x. If x cannot be represented in a uint64, the result is undefined.

​	Uint64 返回 x 的 uint64 表示形式。如果无法在 uint64 中表示 x，则结果未定义。

#### (*Int) UnmarshalJSON <- go1.1

```go
func (z *Int) UnmarshalJSON(text []byte) error
```

UnmarshalJSON implements the json.Unmarshaler interface.

​	UnmarshalJSON 实现 json.Unmarshaler 接口。

#### (*Int) UnmarshalText <- go1.3

```go
func (z *Int) UnmarshalText(text []byte) error
```

UnmarshalText implements the encoding.TextUnmarshaler interface.

​	UnmarshalText 实现 encoding.TextUnmarshaler 接口。

#### (*Int) Xor

```go
func (z *Int) Xor(x, y *Int) *Int
```

Xor sets z = x ^ y and returns z.

​	Xor 设置 z = x ^ y 并返回 z。

### type Rat

```go
type Rat struct {
	// contains filtered or unexported fields
}
```

A Rat represents a quotient a/b of arbitrary precision. The zero value for a Rat represents the value 0.

​	A Rat 表示任意精度的商 a/b。Rat 的零值表示值 0。

Operations always take pointer arguments (*Rat) rather than Rat values, and each unique Rat value requires its own unique *Rat pointer. To “copy” a Rat value, an existing (or newly allocated) Rat must be set to a new value using the Rat.Set method; shallow copies of Rats are not supported and may lead to errors.

​	操作始终采用指针参数 (*Rat) 而不是 Rat 值，并且每个唯一的 Rat 值都需要其自己的唯一 *Rat 指针。要“复制”Rat 值，必须使用 Rat.Set 方法将现有（或新分配的）Rat 设置为新值；不支持 Rat 的浅层副本，并且可能导致错误。

#### func NewRat

```go
func NewRat(a, b int64) *Rat
```

NewRat creates a new Rat with numerator a and denominator b.

​	NewRat 使用分子 a 和分母 b 创建一个新的 Rat。

#### (*Rat) Abs

```go
func (z *Rat) Abs(x *Rat) *Rat
```

Abs sets z to |x| (the absolute value of x) and returns z.

​	Abs 将 z 设置为 |x|（x 的绝对值）并返回 z。

#### (*Rat) Add

```go
func (z *Rat) Add(x, y *Rat) *Rat
```

Add sets z to the sum x+y and returns z.

​	Add 将 z 设置为 x+y 的和并返回 z。

#### (*Rat) Cmp

```go
func (x *Rat) Cmp(y *Rat) int
```

Cmp compares x and y and returns:

​	Cmp 比较 x 和 y 并返回：

```
-1 if x <  y
 0 if x == y
+1 if x >  y
```

#### (*Rat) Denom

```go
func (x *Rat) Denom() *Int
```

Denom returns the denominator of x; it is always > 0. The result is a reference to x’s denominator, unless x is an uninitialized (zero value) Rat, in which case the result is a new Int of value 1. (To initialize x, any operation that sets x will do, including x.Set(x).) If the result is a reference to x’s denominator it may change if a new value is assigned to x, and vice versa.

​	Denom 返回 x 的分母；它始终 > 0。结果是对 x 的分母的引用，除非 x 是一个未初始化的（零值）Rat，在这种情况下，结果是一个值为 1 的新 Int。（要初始化 x，任何设置 x 的操作都可以，包括 x.Set(x)。）如果结果是对 x 的分母的引用，则如果为 x 分配了一个新值，它可能会发生变化，反之亦然。

#### (*Rat) Float32 <- go1.4

```go
func (x *Rat) Float32() (f float32, exact bool)
```

Float32 returns the nearest float32 value for x and a bool indicating whether f represents x exactly. If the magnitude of x is too large to be represented by a float32, f is an infinity and exact is false. The sign of f always matches the sign of x, even if f == 0.

​	Float32 返回 x 的最接近的 float32 值和一个布尔值，指示 f 是否精确地表示 x。如果 x 的大小太大而无法用 float32 表示，则 f 是无穷大，exact 为 false。f 的符号始终与 x 的符号匹配，即使 f == 0。

#### (*Rat) Float64 <- go1.1

```go
func (x *Rat) Float64() (f float64, exact bool)
```

Float64 returns the nearest float64 value for x and a bool indicating whether f represents x exactly. If the magnitude of x is too large to be represented by a float64, f is an infinity and exact is false. The sign of f always matches the sign of x, even if f == 0.

​	Float64 返回 x 的最接近的 float64 值和一个布尔值，指示 f 是否精确地表示 x。如果 x 的大小太大而无法用 float64 表示，则 f 是无穷大，exact 为 false。f 的符号始终与 x 的符号匹配，即使 f == 0。

####  (*Rat) FloatPrec <- go1.22.0

``` go
func (x *Rat) FloatPrec() (n int, exact bool)
```

FloatPrec returns the number n of non-repeating digits immediately following the decimal point of the decimal representation of x. The boolean result indicates whether a decimal representation of x with that many fractional digits is exact or rounded.

​	FloatPrec 返回 x 的十进制表示中，小数点后面立即跟随的非重复数字的数量 n。布尔值结果表示具有这么多小数位的 x 的十进制表示是否精确或已四舍五入。

Examples:

```
x      n    exact    decimal representation n fractional digits
0      0    true     0
1      0    true     1
1/2    1    true     0.5
1/3    0    false    0       (0.333... rounded)
1/4    2    true     0.25
1/6    1    false    0.2     (0.166... rounded)
```

#### (*Rat) FloatString

```go
func (x *Rat) FloatString(prec int) string
```

FloatString returns a string representation of x in decimal form with prec digits of precision after the radix point. The last digit is rounded to nearest, with halves rounded away from zero.

​	FloatString 以十进制形式返回 x 的字符串表示形式，小数点后有 prec 位精度。最后一位四舍五入到最接近的值，一半舍入到远离零的值。

#### (*Rat) GobDecode

```go
func (z *Rat) GobDecode(buf []byte) error
```

GobDecode implements the gob.GobDecoder interface.

​	GobDecode 实现 gob.GobDecoder 接口。

#### (*Rat) GobEncode

```go
func (x *Rat) GobEncode() ([]byte, error)
```

GobEncode implements the gob.GobEncoder interface.

​	GobEncode 实现 gob.GobEncoder 接口。

#### (*Rat) Inv

```go
func (z *Rat) Inv(x *Rat) *Rat
```

Inv sets z to 1/x and returns z. If x == 0, Inv panics.

​	Inv 将 z 设置为 1/x 并返回 z。如果 x == 0，则 Inv 会引发 panic。

#### (*Rat) IsInt

```go
func (x *Rat) IsInt() bool
```

IsInt reports whether the denominator of x is 1.

​	IsInt 报告 x 的分母是否为 1。

#### (*Rat) MarshalText <- go1.3

```go
func (x *Rat) MarshalText() (text []byte, err error)
```

MarshalText implements the encoding.TextMarshaler interface.

​	MarshalText 实现 encoding.TextMarshaler 接口。

#### (*Rat) Mul

```go
func (z *Rat) Mul(x, y *Rat) *Rat
```

Mul sets z to the product x*y and returns z.

​	Mul 将 z 设置为乘积 x*y 并返回 z。

#### (*Rat) Neg

```go
func (z *Rat) Neg(x *Rat) *Rat
```

Neg sets z to -x and returns z.

​	Neg 将 z 设置为 -x 并返回 z。

#### (*Rat) Num

```go
func (x *Rat) Num() *Int
```

Num returns the numerator of x; it may be <= 0. The result is a reference to x’s numerator; it may change if a new value is assigned to x, and vice versa. The sign of the numerator corresponds to the sign of x.

​	Num 返回 x 的分子；它可能 <= 0。结果是对 x 分子的引用；如果为 x 分配新值，它可能会发生变化，反之亦然。分子的符号对应于 x 的符号。

#### (*Rat) Quo

```go
func (z *Rat) Quo(x, y *Rat) *Rat
```

Quo sets z to the quotient x/y and returns z. If y == 0, Quo panics.

​	Quo 将 z 设置为商 x/y 并返回 z。如果 y == 0，则 Quo 会引发恐慌。

#### (*Rat) RatString

```go
func (x *Rat) RatString() string
```

RatString returns a string representation of x in the form “a/b” if b != 1, and in the form “a” if b == 1.

​	RatString 以“a/b”的形式返回 x 的字符串表示形式（如果 b != 1），如果 b == 1，则以“a”的形式返回。

#### (*Rat) Scan

```go
func (z *Rat) Scan(s fmt.ScanState, ch rune) error
```

Scan is a support routine for fmt.Scanner. It accepts the formats ’e’, ‘E’, ‘f’, ‘F’, ‘g’, ‘G’, and ‘v’. All formats are equivalent.

​	Scan 是 fmt.Scanner 的一个支持例程。它接受格式“e”、“E”、“f”、“F”、“g”、“G”和“v”。所有格式都是等效的。

##### Scan Example

```go
package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// The Scan function is rarely used directly;
	// the fmt package recognizes it as an implementation of fmt.Scanner.
	r := new(big.Rat)
	_, err := fmt.Sscan("1.5000", r)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(r)
	}
}
Output:

3/2
```

#### (*Rat) Set

```go
func (z *Rat) Set(x *Rat) *Rat
```

Set sets z to x (by making a copy of x) and returns z.

​	Set 将 z 设置为 x（通过复制 x）并返回 z。

#### (*Rat) SetFloat64 <- go1.1

```go
func (z *Rat) SetFloat64(f float64) *Rat
```

SetFloat64 sets z to exactly f and returns z. If f is not finite, SetFloat returns nil.

​	SetFloat64 将 z 精确设置为 f 并返回 z。如果 f 不是有限的，SetFloat 返回 nil。

#### (*Rat) SetFrac

```go
func (z *Rat) SetFrac(a, b *Int) *Rat
```

SetFrac sets z to a/b and returns z. If b == 0, SetFrac panics.

​	SetFrac 将 z 设置为 a/b 并返回 z。如果 b == 0，SetFrac 会引发 panic。

#### (*Rat) SetFrac64

```go
func (z *Rat) SetFrac64(a, b int64) *Rat
```

SetFrac64 sets z to a/b and returns z. If b == 0, SetFrac64 panics.

​	SetFrac64 将 z 设置为 a/b 并返回 z。如果 b == 0，SetFrac64 会引发 panic。

#### (*Rat) SetInt 

```go
func (z *Rat) SetInt(x *Int) *Rat
```

SetInt sets z to x (by making a copy of x) and returns z.

​	SetInt 将 z 设置为 x（通过复制 x）并返回 z。

#### (*Rat) SetInt64 

```go
func (z *Rat) SetInt64(x int64) *Rat
```

SetInt64 sets z to x and returns z.

​	SetInt64 将 z 设置为 x 并返回 z。

#### (*Rat) SetString

```go
func (z *Rat) SetString(s string) (*Rat, bool)
```

SetString sets z to the value of s and returns z and a boolean indicating success. s can be given as a (possibly signed) fraction “a/b”, or as a floating-point number optionally followed by an exponent. If a fraction is provided, both the dividend and the divisor may be a decimal integer or independently use a prefix of “0b”, “0” or “0o”, or “0x” (or their upper-case variants) to denote a binary, octal, or hexadecimal integer, respectively. The divisor may not be signed. If a floating-point number is provided, it may be in decimal form or use any of the same prefixes as above but for “0” to denote a non-decimal mantissa. A leading “0” is considered a decimal leading 0; it does not indicate octal representation in this case. An optional base-10 “e” or base-2 “p” (or their upper-case variants) exponent may be provided as well, except for hexadecimal floats which only accept an (optional) “p” exponent (because an “e” or “E” cannot be distinguished from a mantissa digit). If the exponent’s absolute value is too large, the operation may fail. The entire string, not just a prefix, must be valid for success. If the operation failed, the value of z is undefined but the returned value is nil.

​	 SetString 将 z 设置为 s 的值，并返回 z 和一个指示成功的布尔值。s 可以表示为（可能带符号的）分数“a/b”，或作为浮点数，后面可以跟一个指数。如果提供分数，则被除数和除数都可以是十进制整数，或独立使用前缀“0b”、“0”或“0o”或“0x”（或它们的大写变体）分别表示二进制、八进制或十六进制整数。除数不能带符号。如果提供浮点数，它可以是十进制形式，或使用上述任何相同的前缀，但对于“0”表示非十进制尾数。前导“0”被视为十进制前导 0；在这种情况下，它不表示八进制表示。也可以提供可选的以 10 为底的“e”或以 2 为底的“p”（或它们的大写变体）指数，但十六进制浮点数只接受（可选的）“p”指数（因为“e”或“E”无法与尾数数字区分开来）。如果指数的绝对值太大，则操作可能会失败。整个字符串（而不仅仅是前缀）必须有效才能成功。 如果操作失败，z 的值未定义，但返回的值为 nil。

##### SetString Example

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	r := new(big.Rat)
	r.SetString("355/113")
	fmt.Println(r.FloatString(3))
}

Output:

3.142
```

#### (*Rat) SetUint64 <- go1.13

```go
func (z *Rat) SetUint64(x uint64) *Rat
```

SetUint64 sets z to x and returns z.

​	SetUint64 将 z 设置为 x 并返回 z。

#### (*Rat) Sign

```go
func (x *Rat) Sign() int
```

Sign returns:

​	Sign 返回：

```
-1 if x <  0
 0 if x == 0
+1 if x >  0
```

#### (*Rat) String

```go
func (x *Rat) String() string
```

String returns a string representation of x in the form “a/b” (even if b == 1).

​	String 以“a/b”的形式返回 x 的字符串表示形式（即使 b == 1）。

#### (*Rat) Sub

```go
func (z *Rat) Sub(x, y *Rat) *Rat
```

Sub sets z to the difference x-y and returns z.

​	Sub 将 z 设置为差值 x-y 并返回 z。

#### (*Rat) UnmarshalText <- go1.3 

```go
func (z *Rat) UnmarshalText(text []byte) error
```

UnmarshalText implements the encoding.TextUnmarshaler interface.

​	UnmarshalText 实现 encoding.TextUnmarshaler 接口。

### type RoundingMode <- go1.5

```go
type RoundingMode byte
```

RoundingMode determines how a Float value is rounded to the desired precision. Rounding may change the Float value; the rounding error is described by the Float’s Accuracy.

​	RoundingMode 确定如何将 Float 值舍入到所需精度。舍入可能会更改 Float 值；舍入误差由 Float 的精度描述。

#### Example 示例

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	operands := []float64{2.6, 2.5, 2.1, -2.1, -2.5, -2.6}

	fmt.Print("   x")
	for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
		fmt.Printf("  %s", mode)
	}
	fmt.Println()

	for _, f64 := range operands {
		fmt.Printf("%4g", f64)
		for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
			// sample operands above require 2 bits to represent mantissa
			// set binary precision to 2 to round them to integer values
			f := new(big.Float).SetPrec(2).SetMode(mode).SetFloat64(f64)
			fmt.Printf("  %*g", len(mode.String()), f)
		}
		fmt.Println()
	}

}

Output:

   x  ToNearestEven  ToNearestAway  ToZero  AwayFromZero  ToNegativeInf  ToPositiveInf
 2.6              3              3       2             3              2              3
 2.5              2              3       2             3              2              3
 2.1              2              2       2             3              2              3
-2.1             -2             -2      -2            -3             -3             -2
-2.5             -2             -3      -2            -3             -3             -2
-2.6             -3             -3      -2            -3             -3             -2
const (
	ToNearestEven RoundingMode = iota // == IEEE 754-2008 roundTiesToEven
	ToNearestAway                     // == IEEE 754-2008 roundTiesToAway
	ToZero                            // == IEEE 754-2008 roundTowardZero
	AwayFromZero                      // no IEEE 754-2008 equivalent
	ToNegativeInf                     // == IEEE 754-2008 roundTowardNegative
	ToPositiveInf                     // == IEEE 754-2008 roundTowardPositive
)
```

These constants define supported rounding modes.

​	这些常量定义受支持的舍入模式。

#### (RoundingMode) String <- go1.5 （RoundingMode）String <- go1.5

```go
func (i RoundingMode) String() string
```

### type Word

```go
type Word uint
```

A Word represents a single digit of a multi-precision unsigned integer.

​	Word 表示多精度无符号整数的单个数字。