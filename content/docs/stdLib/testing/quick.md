+++
title = "quick"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# quick

[https://pkg.go.dev/testing/quick@go1.20.1](https://pkg.go.dev/testing/quick@go1.20.1)

quick包实现了一些实用的函数，以帮助黑盒测试。

testing/quick包已被冻结，不接受新特性。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [Check](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=263) 

``` go 
func Check(f any, config *Config) error
```

​	Check函数会查找一个`f`的输入，即任何返回bool类型的函数，使得`f`返回false。它会重复调用`f`，对于每个参数都使用任意值。如果`f`在给定的输入上返回false，则Check会将该输入作为`*CheckError`返回。例如：

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

#### func [CheckEqual](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=302) 

``` go 
func CheckEqual(f, g any, config *Config) error
```

​	CheckEqual函数会查找一个输入，使得`f`和`g`返回不同的结果。它会重复调用`f`和`g`，对于每个参数都使用任意值。如果`f`和`g`返回不同的答案，则CheckEqual函数会返回一个`*CheckEqualError`，描述输入和输出。

#### func [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=59) 

``` go 
func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)
```

​	Value函数返回给定类型的任意值。如果类型实现了Generator接口，则会使用它。注意：要为结构体创建任意值，必须导出所有字段。

## 类型

### type [CheckEqualError](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=238) 

``` go 
type CheckEqualError struct {
	CheckError
	Out1 []any
	Out2 []any
}
```

​	CheckEqualError结构体是CheckEqual函数发现错误的结果。

#### (*CheckEqualError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=244) 

``` go 
func (s *CheckEqualError) Error() string
```

### type [CheckError](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=228) 

``` go 
type CheckError struct {
	Count int
	In    []any
}
```

​	CheckError结构体是Check函数发现错误的结果。

#### (*CheckError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=233) 

``` go 
func (s *CheckError) Error() string
```

### type [Config](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=177) 

``` go 
type Config struct {
	// MaxCount设置最大迭代次数。
	// 如果为零，则使用MaxCountScale。
	MaxCount int
	// MaxCountScale是默认最大值的非负比例因子。
	// 0表示默认值，通常为100，但可以通过-quickchecks标志设置。
	MaxCountScale float64
	// Rand指定随机数的源。
	// 如果为nil，则使用默认的伪随机源。
	Rand *rand.Rand
	// Values指定生成与正在测试的函数的参数相一致的
    // 任意reflect.Values切片的函数。
	// 如果为nil，则使用顶层Value函数来生成它们。
	Values func([]reflect.Value, *rand.Rand)
}
```

A Config structure contains options for running a test.

Config结构包含运行测试的选项。

### type [Generator](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=23) 

``` go 
type Generator interface {
	// Generate使用size作为大小提示，返回其所属类型的随机实例。
	Generate(rand *rand.Rand, size int) reflect.Value
}
```

​	Generator接口可以生成其自身类型的随机值。

### type [SetupError](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=223) 

``` go 
type SetupError string
```

​	SetupError类型是check的使用方式中出现的错误结果，与被测试的函数无关。

#### (SetupError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/quick/quick.go;l=225) 

``` go 
func (s SetupError) Error() string
```