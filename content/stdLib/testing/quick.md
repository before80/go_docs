+++
title = "quick"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/testing/quick@go1.23.0](https://pkg.go.dev/testing/quick@go1.23.0)

Package quick implements utility functions to help with black box testing.

​	`quick`包实现了一些实用的函数，以帮助黑盒测试。

The testing/quick package is frozen and is not accepting new features.

​	`testing/quick`包已被冻结，不接受新特性。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Check 

``` go 
func Check(f any, config *Config) error
```

Check looks for an input to f, any function that returns bool, such that f returns false. It calls f repeatedly, with arbitrary values for each argument. If f returns false on a given input, Check returns that input as a *CheckError. For example:

​	`Check`函数会查找一个`f`的输入，即任何返回bool类型的函数，使得`f`返回false。它会重复调用`f`，对于每个参数都使用任意值。如果`f`在给定的输入上返回false，则Check会将该输入作为`*CheckError`返回。例如：

``` go 
func TestOddMultipleOfThree(t *testing.T) {
	f := func(x int) bool {
		y := OddMultipleOfThree(x)
		return y%2 == 1 && y%3 == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
```

### func CheckEqual 

``` go 
func CheckEqual(f, g any, config *Config) error
```

CheckEqual looks for an input on which f and g return different results. It calls f and g repeatedly with arbitrary values for each argument. If f and g return different answers, CheckEqual returns a *CheckEqualError describing the input and the outputs.

​	`CheckEqual`函数会查找一个输入，使得`f`和`g`返回不同的结果。它会重复调用`f`和`g`，对于每个参数都使用任意值。如果`f`和`g`返回不同的答案，则CheckEqual函数会返回一个`*CheckEqualError`，描述输入和输出。

### func Value 

``` go 
func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)
```

Value returns an arbitrary value of the given type. If the type implements the Generator interface, that will be used. Note: To create arbitrary values for structs, all the fields must be exported.

​	`Value`函数返回给定类型的任意值。如果类型实现了Generator接口，则会使用它。注意：要为结构体创建任意值，必须导出所有字段。

## 类型

### type CheckEqualError 

``` go 
type CheckEqualError struct {
	CheckError
	Out1 []any
	Out2 []any
}
```

A CheckEqualError is the result CheckEqual finding an error.

​	`CheckEqualError`结构体是CheckEqual函数发现错误的结果。

#### (*CheckEqualError) Error 

``` go 
func (s *CheckEqualError) Error() string
```

### type CheckError 

``` go 
type CheckError struct {
	Count int
	In    []any
}
```

A CheckError is the result of Check finding an error.

​	`CheckError`结构体是Check函数发现错误的结果。

#### (*CheckError) Error 

``` go 
func (s *CheckError) Error() string
```

### type Config 

``` go 
type Config struct {
    // MaxCount sets the maximum number of iterations.
	// If zero, MaxCountScale is used.
	// MaxCount设置最大迭代次数。
	// 如果为零，则使用MaxCountScale。
	MaxCount int
    // MaxCountScale is a non-negative scale factor applied to the
	// default maximum.
	// A count of zero implies the default, which is usually 100
	// but can be set by the -quickchecks flag.
	// MaxCountScale是默认最大值的非负比例因子。
	// 0表示默认值，通常为100，但可以通过-quickchecks标志设置。
	MaxCountScale float64
    // Rand specifies a source of random numbers.
	// If nil, a default pseudo-random source will be used.
	// Rand指定随机数的源。
	// 如果为nil，则使用默认的伪随机源。
	Rand *rand.Rand
    // Values specifies a function to generate a slice of
	// arbitrary reflect.Values that are congruent with the
	// arguments to the function being tested.
	// If nil, the top-level Value function is used to generate them.
	// Values指定生成与正在测试的函数的参数相一致的
    // 任意reflect.Values切片的函数。
	// 如果为nil，则使用顶层Value函数来生成它们。
	Values func([]reflect.Value, *rand.Rand)
}
```

A Config structure contains options for running a test.

​	`Config`结构包含运行测试的选项。

### type Generator 

``` go 
type Generator interface {
    // Generate returns a random instance of the type on which it is a
	// method using the size as a size hint.
	// Generate使用size作为大小提示，返回其所属类型的随机实例。
	Generate(rand *rand.Rand, size int) reflect.Value
}
```

A Generator can generate random values of its own type.

​	Generator接口可以生成其自身类型的随机值。

### type SetupError 

``` go 
type SetupError string
```

A SetupError is the result of an error in the way that check is being used, independent of the functions being tested.

​	`SetupError`类型是check的使用方式中出现的错误结果，与被测试的函数无关。

#### (SetupError) Error 

``` go 
func (s SetupError) Error() string
```