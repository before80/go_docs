+++
title = "pflag"
type = "docs"
weight = 2
date = 2023-05-21T16:19:55+08:00
description = ""
isCJKLanguage = true
draft = false

+++

# pflag

> 原文：[https://pkg.go.dev/github.com/spf13/pflag](https://pkg.go.dev/github.com/spf13/pflag)
>
> 版本：v1.0.5
>
> 发布日期：2019.9.18
>
> github网址：[https://github.com/spf13/pflag](https://github.com/spf13/pflag)

## 描述

​	pflag是go语言的`flag`包的一个替代品，实现了POSIX/GNU风格的`--flags`。

​	pflag兼容[GNU对命令行选项的POSIX建议的扩展](http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html)。更详细的描述请参见下面的"命令行标志语法"部分。

​	pflag采用与Go语言相同的BSD许可证授权，许可证的内容可以在LICENSE文件中找到。

## 安装

​	使用标准的`go get`命令可以获取pflag。

​	运行以下命令进行安装：

```bash
go get github.com/spf13/pflag
```

​	运行以下命令进行测试：

```bash
go test github.com/spf13/pflag
```

## 用法

​	`pflag`是Go原生`flag`包的一个替代品。如果您将pflag导入为"flag"，则所有的代码都应该继续正常工作，无需进行任何更改。

```go
import flag "github.com/spf13/pflag"
```

​	这里有一个例外情况：如果你直接实例化 Flag 结构体，你需要设置一个额外的字段 "Shorthand"。大多数代码从不直接实例化这个结构体，而是使用 String()、BoolVar() 和 Var() 等函数，因此不受影响。

​	使用flag.String()、Bool()、Int()等函数来定义标志。

​	以下是一个声明整数标志的示例，标志名为`-flagname`，存储在指针`ip`中，类型为`*int`：

```go
var ip *int = flag.Int("flagname", 1234, "help message for flagname")
```

​	如果需要，您可以使用Var()函数将标志绑定到一个变量上：

```go
var flagvar int
func init() {
    flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

​	或者您可以创建满足Value接口的自定义标志（使用指针接收器），并通过以下方式将它们与标志解析相关联：

```go
flag.Var(&flagVal, "name", "help message for flagname")
```

​	对于这样的标志，其默认值就是变量的初始值。

​	在定义所有标志后，调用

```go
flag.Parse()
```

来将命令行解析进已定义的标志。

​	然后可以直接使用标志。**如果使用的是标志本身，则它们都是指针；如果绑定到变量，则它们是值**。

```go
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

​	如果您使用 `FlagSet`，并且发现在代码中跟踪所有指针变得困难，那么可以使用一些辅助函数来获取 `Flag`（结构体） 中存储的值。如果您有一个名为'flagname'、类型为`int`的`pflag.FlagSet`，您可以使用`GetInt()`来获取int值。但请注意，'flagname'必须存在且为int类型，否则`GetString("flagname")`将失败。

```go
i, err := flagset.GetInt("flagname")
```

​	在解析后，该标志之后的实参可以作为`flag.Args()`的切片或作为`flag.Arg(i)`单独使用。实参的索引范围是从0到`flag.NArg()-1`。

​	pflag包还定义了一些在 `flag` 包中不存在的新函数，这些函数为标志提供了一字母缩写。您可以通过在定义标志的任何函数名称后附加 '`P`' 来使用这些函数。

```go
var ip = flag.IntP("flagname", "f", 1234, "help message")
var flagvar bool
func init() {
	flag.BoolVarP(&flagvar, "boolname", "b", true, "help message")
}
flag.VarP(&flagVal, "varname", "v", "help message")
```

​	速记字母可以在命令行上使用单个短划线。布尔型缩写标志可以与其他缩写标志结合使用。

​	默认的命令行标志集由顶层函数控制。FlagSet类型允许定义独立的标志集，例如在命令行接口中实现子命令。FlagSet的方法类似于顶层函数用于命令行标志集的方法。

## 为标志设置无选项默认值

​	在创建标志后，可以为给定的标志设置pflag.NoOptDefVal。这样做会略微改变标志的含义。如果一个标志具有NoOptDefVal，并且在命令行上设置该标志而没有选项，那么该标志将被设置为NoOptDefVal。例如：

```go
var ip = flag.IntP("flagname", "f", 1234, "help message")
flag.Lookup("flagname").NoOptDefVal = "4321"
```

​	会产生类似以下的结果：

| Parsed Arguments  | 结果值  |
| ----------------- | ------- |
| `--flagname=1357` | ip=1357 |
| `--flagname`      | ip=4321 |
| [nothing]         | ip=1234 |

## 命令行标志语法

```bash
--flag    // boolean flags, or flags with no option default values
--flag x  // only on flags without a default value
--flag=x
```

Unlike the flag package, a single dash before an option means something different than a double dash. Single dashes signify a series of shorthand letters for flags. All but the last shorthand letter must be boolean flags or a flag with a default value

​	与`flag`包不同，选项之前的单个短划线和双破折号有不同的含义。单个短划线表示一系列标志的速记字母。除最后一个速记字母外，其它都必须是布尔标志或具有默认值的标志

```
// 布尔标志或带有 'no option default value' 的标志
-f
-f=true
-abc
but
-b true is INVALID

// 非布尔和没有 'no option default value' 的标志
-n 1234
-n=1234
-n1234

// 混合使用
-abcs "hello"
-absd="hello"
-abcs1234
```

​	在终止符 "`--`" 之后，标志解析会停止。与flag包不同，在这个终止符之前，标志可以与参数混合在命令行的任何位置。

​	整型标志接受1234、0664、0x1234等，并且可以为负数。布尔型标志（长格式）接受1、0、t、f、true、false、TRUE、FALSE、True、False。持续时间标志接受任何对于time.ParseDuration有效的输入。

## 修改或"规范化"标志名称

​	可以设置自定义的标志名称"规范化函数"。它允许标志名称在代码中创建时和在命令行上使用时以某种"规范化"的形式进行变换。比较时使用"规范化"的形式。下面是两个使用自定义规范化函数的示例。

**示例＃1**：您希望在标志中比较 `-`、`_` 和 `.` 时得到相同的结果。也就是说 `--my-flag == --my_flag == --my.flag`：

``` go
func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

myFlagSet.SetNormalizeFunc(wordSepNormalizeFunc)
```

**示例＃2**：您希望给两个标志设置别名。也就是说 `--old-flag-name == --new-flag-name`：

``` go
func aliasNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "old-flag-name":
		name = "new-flag-name"
		break
	}
	return pflag.NormalizedName(name)
}

myFlagSet.SetNormalizeFunc(aliasNormalizeFunc)
```

## 弃用标志或其缩写

​	可以弃用一个标志，或仅弃用它的缩写。弃用标志/缩写将在帮助文本中隐藏，并在使用被弃用的标志/缩写时显示使用消息。

**示例＃1**：您希望弃用一个名为"badflag"的标志，并告知用户应该使用哪个标志代替它：

```go
// 通过指定其名称和使用消息来废弃一个标志
flags.MarkDeprecated("badflag", "please use --good-flag instead")
```

​	这将在帮助文本中隐藏"badflag"，并在使用"badflag"时打印`Flag --badflag has been deprecated, please use --good-flag instead`。

**示例＃2**：您希望保留一个名为"noshorthandflag"的标志，但弃用其缩写"n"：

```go
// 通过指定其标志名称和使用消息来废弃一个标志缩写
flags.MarkShorthandDeprecated("noshorthandflag", "please use --noshorthandflag only")
```

​	这将在帮助文本中隐藏缩写"n"，并在使用缩写"n"时打印`Flag shorthand -n has been deprecated, please use --noshorthandflag only`。

​	请注意，这里的用法消息是必要的，不应为空。

## 隐藏标志

​	可以将一个标志标记为隐藏，这意味着它仍然会正常工作，但不会显示在用法/帮助文本中。

**示例**：您有一个名为"secretFlag"的标志，仅供内部使用，不希望它显示在帮助文本中或可用于用法文本：

```go
// 通过指定其名称来隐藏标志
flags.MarkHidden("secretFlag")
```

## 禁用标志的排序

​	`pflag`允许您禁用帮助和用法消息的标志排序。

示例：

```go
flags.BoolP("verbose", "v", false, "verbose output")
flags.String("coolflag", "yeaah", "it's really cool flag")
flags.Int("usefulflag", 777, "sometimes it's very useful")
flags.SortFlags = false
flags.PrintDefaults()
```

输出:

```
  -v, --verbose           verbose output
      --coolflag string   it's really cool flag (default "yeaah")
      --usefulflag int    sometimes it's very useful (default 777)
```

## 在使用pflag时支持Go（标准库中的flag定义的）标志

​	为了支持使用Go的`flag`包定义的标志，它们必须被添加到`pflag`的标志集中。这通常是为了支持由第三方依赖（例如`golang/glog`）定义的标志。

**示例**：您希望将Go标志添加到`CommandLine`的标志集中：

```go
import (
	goflag "flag"
	flag "github.com/spf13/pflag"
)

var ip *int = flag.Int("flagname", 1234, "help message for flagname")

func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()
}
```

## 更多信息

​	您可以在[godoc.org](http://godoc.org/github.com/spf13/pflag)上查看pflag包的完整参考文档，或者在安装后通过运行`godoc -http=:6060`并浏览http://localhost:6060/pkg/github.com/spf13/pflag来使用 go 的标准文档系统。



## 文档概述

​	pflag包是Go标准库flag包的一个替代品，实现了POSIX/GNU风格的`--flags`。

​	pflag 与 GNU 对命令行选项的 POSIX 建议的扩展兼容。请参阅 [http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html](http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html)。

### 常量

This section is empty.

### 变量

[View Source](https://github.com/spf13/pflag/blob/v1.0.5/flag.go#L1212)

```go
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

​	CommandLine 变量是默认的命令行标志集，从 os.Args 解析而来。

[View Source](https://github.com/spf13/pflag/blob/v1.0.5/flag.go#L113)

```go
var ErrHelp = errors.New("pflag: help requested")
```

ErrHelp is the error returned if the flag -help is invoked but no such flag is defined.

​	ErrHelp 变量是在调用标志 `-help` 但没有定义此类标志时返回的错误。

[View Source](https://github.com/spf13/pflag/blob/v1.0.5/flag.go#L773)

```go
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	PrintDefaults()
}
```

​	Usage 变量将一个用于文档化所有定义的命令行标志的用法消息打印到标准错误。该函数是一个变量，可以更改为指向自定义函数。默认情况下，它会打印一个简单的标题并调用 PrintDefaults；关于输出格式及如何控制它的详细信息，请参阅 PrintDefaults 的文档。

### 函数

#### func Arg 

``` go
func Arg(i int) string
```

​	Arg 函数返回第 i 个命令行实参。Arg(0) 是在处理标志后剩余的第一个实参。

#### func Args 

``` go
func Args() []string
```

​	Args 函数返回非标志的命令行实参。

#### func Bool 

``` go
func Bool(name string, value bool, usage string) *bool
```

​	Bool 函数定义一个具有指定名称、默认值和用法说明的 bool 标志。返回值是存储标志值的 bool 变量的地址。

#### func BoolP 

``` go
func BoolP(name, shorthand string, value bool, usage string) *bool
```

​	BoolP 函数类似于 Bool函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func BoolSlice 

``` go
func BoolSlice(name string, value []bool, usage string) *[]bool
```

​	BoolSlice 函数定义一个具有指定名称、默认值和用法说明的 `[]bool` 标志。返回值是存储标志值的 `[]bool` 变量的地址。

#### func BoolSliceP 

``` go
func BoolSliceP(name, shorthand string, value []bool, usage string) *[]bool
```

​	BoolSliceP 函数类似于 BoolSlice函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func BoolSliceVar 

``` go
func BoolSliceVar(p *[]bool, name string, value []bool, usage string)
```

​	BoolSliceVar 函数定义一个具有指定名称、默认值和用法说明的 `[]bool` 标志。参数 p 指向一个 `[]bool` 变量，用于存储标志值。

#### func BoolSliceVarP 

``` go
func BoolSliceVarP(p *[]bool, name, shorthand string, value []bool, usage string)
```

​	BoolSliceVarP 函数类似于 BoolSliceVar函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func BoolVar 

``` go
func BoolVar(p *bool, name string, value bool, usage string)
```

​	BoolVar 函数定义一个具有指定名称、默认值和用法说明的 bool 标志。参数 p 指向一个 bool 变量，用于存储标志值。

#### func BoolVarP 

``` go
func BoolVarP(p *bool, name, shorthand string, value bool, usage string)
```

​	BoolVarP 函数类似于 BoolVar函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func BytesBase64  <- v1.0.2

``` go
func BytesBase64(name string, value []byte, usage string) *[]byte
```

​	BytesBase64 函数定义一个具有指定名称、默认值和用法说明的 `[]byte` 标志。返回值是存储标志值的 `[]byte` 变量的地址。

#### func BytesBase64P  <- v1.0.2

``` go
func BytesBase64P(name, shorthand string, value []byte, usage string) *[]byte
```

​	BytesBase64P 函数类似于 BytesBase64函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func BytesBase64Var  <- v1.0.2

``` go
func BytesBase64Var(p *[]byte, name string, value []byte, usage string)
```

​	BytesBase64Var 函数定义一个具有指定名称、默认值和用法说明的 `[]byte` 标志。参数 p 指向一个 `[]byte` 变量，用于存储标志值。

#### func BytesBase64VarP  <- v1.0.2

``` go
func BytesBase64VarP(p *[]byte, name, shorthand string, value []byte, usage string)
```

​	BytesBase64VarP 函数类似于 BytesBase64Var函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func BytesHex  <- v1.0.1

``` go
func BytesHex(name string, value []byte, usage string) *[]byte
```

​	BytesHex 函数定义一个具有指定名称、默认值和用法说明的 `[]byte` 标志。返回值是存储标志值的 []byte 变量的地址。

#### func BytesHexP  <- v1.0.1

``` go
func BytesHexP(name, shorthand string, value []byte, usage string) *[]byte
```

​	BytesHexP 函数类似于 BytesHex函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func BytesHexVar  <- v1.0.1

``` go
func BytesHexVar(p *[]byte, name string, value []byte, usage string)
```

​	BytesHexVar 函数定义一个具有指定名称、默认值和用法说明的 `[]byte` 标志。参数 p 指向一个 []byte 变量，用于存储标志值。

#### func BytesHexVarP  <- v1.0.1

``` go
func BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string)
```

​	BytesHexVarP 函数类似于 BytesHexVar函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Count 

``` go
func Count(name string, usage string) *int
```

​	Count 函数定义一个具有指定名称、默认值和用法说明的计数标志。返回值是存储标志值的 int 变量的地址。计数标志每次在命令行上找到时都会将其值加 1。

#### func CountP 

``` go
func CountP(name, shorthand string, usage string) *int
```

​	CountP 函数类似于 Count函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func CountVar 

``` go
func CountVar(p *int, name string, usage string)
```

​	CountVar 函数类似于 Count函数，但将标志放在 CommandLine 上，而不是在给定的标志集中。

#### func CountVarP 

``` go
func CountVarP(p *int, name, shorthand string, usage string)
```

​	CountVarP 函数类似于 CountVar函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Duration 

``` go
func Duration(name string, value time.Duration, usage string) *time.Duration
```

​	Duration 函数定义一个具有指定名称、默认值和用法说明的 time.Duration 标志。返回值是存储标志值的 time.Duration 变量的地址。

#### func DurationP 

``` go
func DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration
```

​	DurationP 函数类似于 Duration函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func DurationSlice  <- v1.0.1

``` go
func DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration
```

​	DurationSlice 函数定义一个具有指定名称、默认值和用法说明的 []time.Duration 标志。返回值是存储标志值的 []time.Duration 变量的地址。

#### func DurationSliceP  <- v1.0.1

``` go
func DurationSliceP(name, shorthand string, value []time.Duration, usage string) *[]time.Duration
```

​	DurationSliceP 函数类似于 DurationSlice函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func DurationSliceVar  <- v1.0.1

``` go
func DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string)
```

​	DurationSliceVar 函数定义一个具有指定名称、默认值和用法说明的 `[]time.Duration` 标志。参数 p 指向一个 `[]time.Duration` 变量，用于存储标志值。

#### func DurationSliceVarP  <- v1.0.1

``` go
func DurationSliceVarP(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string)
```

​	DurationSliceVarP 函数类似于 DurationSliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func DurationVar 

``` go
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

​	DurationVar 函数定义一个具有指定名称、默认值和用法说明的 time.Duration 标志。参数 p 指向一个 time.Duration 变量，用于存储标志值。

#### func DurationVarP 

``` go
func DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string)
```

​	DurationVarP 函数类似于 DurationVar函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Float32 

``` go
func Float32(name string, value float32, usage string) *float32
```

​	Float32 函数定义一个具有指定名称、默认值和用法说明的 float32 标志。返回值是存储标志值的 float32 变量的地址。

#### func Float32P 

``` go
func Float32P(name, shorthand string, value float32, usage string) *float32
```

​	Float32P 函数类似于 Float32函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Float32Slice  <- v1.0.5

``` go
func Float32Slice(name string, value []float32, usage string) *[]float32
```

​	Float32Slice 函数定义一个具有指定名称、默认值和用法说明的 `[]float32` 标志。返回值是存储标志值的 `[]float32` 变量的地址。

#### func Float32SliceP  <- v1.0.5

``` go
func Float32SliceP(name, shorthand string, value []float32, usage string) *[]float32
```

​	Float32SliceP 函数类似于 Float32Slice函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Float32SliceVar  <- v1.0.5

``` go
func Float32SliceVar(p *[]float32, name string, value []float32, usage string)
```

​	Float32SliceVar 函数定义一个具有指定名称、默认值和用法说明的 `[]float32` 标志。参数 p 指向一个 `[]float32` 变量，用于存储标志值。

#### func Float32SliceVarP  <- v1.0.5

``` go
func Float32SliceVarP(p *[]float32, name, shorthand string, value []float32, usage string)
```

​	Float32SliceVarP 函数类似于 Float32SliceVar函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Float32Var 

``` go
func Float32Var(p *float32, name string, value float32, usage string)
```

​	Float32Var 函数定义一个具有指定名称、默认值和用法说明的 float32 标志。参数 p 指向一个 float32 变量，用于存储标志值。

#### func Float32VarP 

``` go
func Float32VarP(p *float32, name, shorthand string, value float32, usage string)
```

​	Float32VarP 函数类似于 Float32Var函数，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Float64 

``` go
func Float64(name string, value float64, usage string) *float64
```

​	Float64 函数定义一个具有指定名称、默认值和用法说明的 float64 标志。返回值是存储标志值的 float64 变量的地址。

#### func Float64P 

``` go
func Float64P(name, shorthand string, value float64, usage string) *float64
```

​	Float64P 类似于 Float64，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Float64Slice  <- v1.0.5

``` go
func Float64Slice(name string, value []float64, usage string) *[]float64
```

​	Float64Slice 定义一个具有指定名称、默认值和用法说明的 `[]float64` 标志。返回值是存储标志值的 `[]float64` 变量的地址。

#### func Float64SliceP  <- v1.0.5

``` go
func Float64SliceP(name, shorthand string, value []float64, usage string) *[]float64
```

​	Float64SliceP 类似于 Float64Slice，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Float64SliceVar  <- v1.0.5

``` go
func Float64SliceVar(p *[]float64, name string, value []float64, usage string)
```

​	Float64SliceVar 定义一个具有指定名称、默认值和用法说明的 []float64 标志。参数 p 指向一个 []float64 变量，用于存储标志值。

#### func Float64SliceVarP  <- v1.0.5

``` go
func Float64SliceVarP(p *[]float64, name, shorthand string, value []float64, usage string)
```

​	Float64SliceVarP 类似于 Float64SliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Float64Var 

``` go
func Float64Var(p *float64, name string, value float64, usage string)
```

​	Float64Var 定义一个具有指定名称、默认值和用法说明的 float64 标志。参数 p 指向一个 float64 变量，用于存储标志值。

#### func Float64VarP 

``` go
func Float64VarP(p *float64, name, shorthand string, value float64, usage string)
```

​	Float64VarP 类似于 Float64Var，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IP 

``` go
func IP(name string, value net.IP, usage string) *net.IP
```

​	IP 定义一个具有指定名称、默认值和用法说明的 net.IP 标志。返回值是存储标志值的 net.IP 变量的地址。

#### func IPMask 

``` go
func IPMask(name string, value net.IPMask, usage string) *net.IPMask
```

​	IPMask 定义一个具有指定名称、默认值和用法说明的 net.IPMask 标志。返回值是存储标志值的 net.IPMask 变量的地址。

#### func IPMaskP 

``` go
func IPMaskP(name, shorthand string, value net.IPMask, usage string) *net.IPMask
```

​	IPMaskP 类似于 IP，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IPMaskVar 

``` go
func IPMaskVar(p *net.IPMask, name string, value net.IPMask, usage string)
```

​	IPMaskVar 定义一个具有指定名称、默认值和用法说明的 net.IPMask 标志。参数 p 指向一个 net.IPMask 变量，用于存储标志值。

#### func IPMaskVarP 

``` go
func IPMaskVarP(p *net.IPMask, name, shorthand string, value net.IPMask, usage string)
```

​	IPMaskVarP 类似于 IPMaskVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IPNet 

``` go
func IPNet(name string, value net.IPNet, usage string) *net.IPNet
```

​	IPNet 定义一个具有指定名称、默认值和用法说明的 net.IPNet 标志。返回值是存储标志值的 net.IPNet 变量的地址。

#### func IPNetP 

``` go
func IPNetP(name, shorthand string, value net.IPNet, usage string) *net.IPNet
```

​	IPNetP 类似于 IPNet，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IPNetVar 

``` go
func IPNetVar(p *net.IPNet, name string, value net.IPNet, usage string)
```

​	IPNetVar 定义一个具有指定名称、默认值和用法说明的 net.IPNet 标志。参数 p 指向一个 net.IPNet 变量，用于存储标志值。

#### func IPNetVarP 

``` go
func IPNetVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string)
```

​	IPNetVarP 类似于 IPNetVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IPP 

``` go
func IPP(name, shorthand string, value net.IP, usage string) *net.IP
```

​	IPP 类似于 IP，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IPSlice 

``` go
func IPSlice(name string, value []net.IP, usage string) *[]net.IP
```

​	IPSlice 定义一个具有指定名称、默认值和用法说明的 `[]net.IP` 标志。返回值是存储标志值的 `[]net.IP` 变量的地址。

#### func IPSliceP 

``` go
func IPSliceP(name, shorthand string, value []net.IP, usage string) *[]net.IP
```

​	IPSliceP 类似于 IPSlice，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IPSliceVar 

``` go
func IPSliceVar(p *[]net.IP, name string, value []net.IP, usage string)
```

​	IPSliceVar 定义一个具有指定名称、默认值和用法说明的 `[]net.IP` 标志。参数 p 指向一个 `[]net.IP` 变量，用于存储标志值。

#### func IPSliceVarP 

``` go
func IPSliceVarP(p *[]net.IP, name, shorthand string, value []net.IP, usage string)
```

​	IPSliceVarP 类似于 IPSliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IPVar 

``` go
func IPVar(p *net.IP, name string, value net.IP, usage string)
```

​	IPVar 定义一个具有指定名称、默认值和用法说明的 net.IP 标志。参数 p 指向一个 net.IP 变量，用于存储标志值。

#### func IPVarP 

``` go
func IPVarP(p *net.IP, name, shorthand string, value net.IP, usage string)
```

​	IPVarP 类似于 IPVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int 

``` go
func Int(name string, value int, usage string) *int
```

​	Int 定义一个具有指定名称、默认值和用法说明的 int 标志。返回值是存储标志值的 int 变量的地址。

#### func Int16  <- v1.0.1

``` go
func Int16(name string, value int16, usage string) *int16
```

​	Int16 定义一个具有指定名称、默认值和用法说明的 int16 标志。返回值是存储标志值的 int16 变量的地址。

#### func Int16P  <- v1.0.1

``` go
func Int16P(name, shorthand string, value int16, usage string) *int16
```

​	Int16P 类似于 Int16，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int16Var  <- v1.0.1

``` go
func Int16Var(p *int16, name string, value int16, usage string)
```

​	Int16Var 定义一个具有指定名称、默认值和用法说明的 int16 标志。参数 p 指向一个 int16 变量，用于存储标志值。

#### func Int16VarP  <- v1.0.1

``` go
func Int16VarP(p *int16, name, shorthand string, value int16, usage string)
```

​	Int16VarP 类似于 Int16Var，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int32 

``` go
func Int32(name string, value int32, usage string) *int32
```

​	Int32 定义一个具有指定名称、默认值和用法说明的 int32 标志。返回值是存储标志值的 int32 变量的地址。

#### func Int32P 

``` go
func Int32P(name, shorthand string, value int32, usage string) *int32
```

​	Int32P 类似于 Int32，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int32Slice  <- v1.0.5

``` go
func Int32Slice(name string, value []int32, usage string) *[]int32
```

​	Int32Slice 定义一个具有指定名称、默认值和用法说明的 `[]int32` 标志。返回值是存储标志值的 `[]int32` 变量的地址。

#### func Int32SliceP  <- v1.0.5

``` go
func Int32SliceP(name, shorthand string, value []int32, usage string) *[]int32
```

​	Int32SliceP 类似于 Int32Slice，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int32SliceVar  <- v1.0.5

``` go
func Int32SliceVar(p *[]int32, name string, value []int32, usage string)
```

​	Int32SliceVar 定义一个具有指定名称、默认值和用法说明的 `[]int32` 标志。参数 p 指向一个 `[]int32` 变量，用于存储标志值。

#### func Int32SliceVarP  <- v1.0.5

``` go
func Int32SliceVarP(p *[]int32, name, shorthand string, value []int32, usage string)
```

​	Int32SliceVarP 类似于 Int32SliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int32Var 

``` go
func Int32Var(p *int32, name string, value int32, usage string)
```

​	Int32Var 定义一个具有指定名称、默认值和用法说明的 int32 标志。参数 p 指向一个 int32 变量，用于存储标志值。

#### func Int32VarP 

``` go
func Int32VarP(p *int32, name, shorthand string, value int32, usage string)
```

​	Int32VarP 类似于 Int32Var，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int64 

``` go
func Int64(name string, value int64, usage string) *int64
```

​	Int64 定义一个具有指定名称、默认值和用法说明的 int64 标志。返回值是存储标志值的 int64 变量的地址。

#### func Int64P 

``` go
func Int64P(name, shorthand string, value int64, usage string) *int64
```

​	Int64P 类似于 Int64，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int64Slice  <- v1.0.5

``` go
func Int64Slice(name string, value []int64, usage string) *[]int64
```

​	Int64Slice 定义一个具有指定名称、默认值和用法说明的 `[]int64` 标志。返回值是存储标志值的 `[]int64` 变量的地址。

#### func Int64SliceP  <- v1.0.5

``` go
func Int64SliceP(name, shorthand string, value []int64, usage string) *[]int64
```

​	Int64SliceP 类似于 Int64Slice，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int64SliceVar  <- v1.0.5

``` go
func Int64SliceVar(p *[]int64, name string, value []int64, usage string)
```

Int64SliceVar defines a int64[] flag with specified name, default value, and usage string. The argument p points to a int64[] variable in which to store the value of the flag.

​	Int64SliceVar 定义一个具有指定名称、默认值和用法说明的 `[]int64` 标志。参数 p 指向一个 `[]int64` 变量，用于存储标志值。

#### func Int64SliceVarP  <- v1.0.5

``` go
func Int64SliceVarP(p *[]int64, name, shorthand string, value []int64, usage string)
```

​	Int64SliceVarP 类似于 Int64SliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int64Var 

``` go
func Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var defines an int64 flag with specified name, default value, and usage string. The argument p points to an int64 variable in which to store the value of the flag.

​	Int64VarP 类似于 Int64Var，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int64VarP 

``` go
func Int64VarP(p *int64, name, shorthand string, value int64, usage string)
```

Int64VarP is like Int64Var, but accepts a shorthand letter that can be used after a single dash.

​	Int8 定义一个具有指定名称、默认值和用法说明的 int8 标志。返回值是存储标志值的 int8 变量的地址。

#### func Int8 

``` go
func Int8(name string, value int8, usage string) *int8
```

Int8 defines an int8 flag with specified name, default value, and usage string. The return value is the address of an int8 variable that stores the value of the flag.

​	Int8P 类似于 Int8，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int8P 

``` go
func Int8P(name, shorthand string, value int8, usage string) *int8
```

Int8P is like Int8, but accepts a shorthand letter that can be used after a single dash.

​	Int8Var 定义一个具有指定名称、默认值和用法说明的 int8 标志。参数 p 指向一个 int8 变量，用于存储标志值。

#### func Int8Var 

``` go
func Int8Var(p *int8, name string, value int8, usage string)
```

Int8Var defines an int8 flag with specified name, default value, and usage string. The argument p points to an int8 variable in which to store the value of the flag.

​	Int8VarP 类似于 Int8Var，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Int8VarP 

``` go
func Int8VarP(p *int8, name, shorthand string, value int8, usage string)
```

Int8VarP is like Int8Var, but accepts a shorthand letter that can be used after a single dash.

​	Int8VarP 类似于 Int8Var，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IntP 

``` go
func IntP(name, shorthand string, value int, usage string) *int
```

IntP is like Int, but accepts a shorthand letter that can be used after a single dash.

​	IntP 类似于 Int，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IntSlice 

``` go
func IntSlice(name string, value []int, usage string) *[]int
```

IntSlice defines a []int flag with specified name, default value, and usage string. The return value is the address of a []int variable that stores the value of the flag.

​	IntSlice 定义一个具有指定名称、默认值和用法说明的 `[]int` 标志。返回值是存储标志值的 `[]int` 变量的地址。

#### func IntSliceP 

``` go
func IntSliceP(name, shorthand string, value []int, usage string) *[]int
```

​	IntSliceP 类似于 IntSlice，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IntSliceVar 

``` go
func IntSliceVar(p *[]int, name string, value []int, usage string)
```

​	IntSliceVar 定义一个具有指定名称、默认值和用法说明的 `[]int` 标志。参数 p 指向一个 `[]int` 变量，用于存储标志值。

#### func IntSliceVarP 

``` go
func IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string)
```

​	IntSliceVarP 类似于 IntSliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func IntVar 

``` go
func IntVar(p *int, name string, value int, usage string)
```

​	IntVar 定义一个具有指定名称、默认值和用法说明的 int 标志。参数 p 指向一个 int 变量，用于存储标志值。

#### func IntVarP 

``` go
func IntVarP(p *int, name, shorthand string, value int, usage string)
```

​	IntVarP 类似于 IntVar，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func NArg 

``` go
func NArg() int
```

​	NArg 返回在处理标志之后剩余的参数数目。unc NFlag 

``` go
func NFlag() int
```

​	NFlag 返回已设置的命令行标志数目。

#### func Parse 

``` go
func Parse()
```

​	Parse 解析来自 os.Args[1:] 的命令行标志。必须在定义所有标志之后，程序访问标志之前调用。

#### func ParseAll 

``` go
func ParseAll(fn func(flag *Flag, value string) error)
```

​	ParseAll 解析来自 os.Args[1:] 的命令行标志，并为每个标志调用 fn。fn 的参数为 flag 和 value。必须在定义所有标志之后，程序访问标志之前调用。

#### func ParseIPv4Mask 

``` go
func ParseIPv4Mask(s string) net.IPMask
```

​	ParseIPv4Mask 将以 IP 形式（例如 255.255.255.0）编写的字符串解析为 net.IPMask。这个函数应该实际上属于 net 包。

#### func Parsed 

``` go
func Parsed() bool
```

​	Parsed 如果已解析命令行标志，则返回 true。

#### func PrintDefaults 

``` go
func PrintDefaults()
```

​	PrintDefaults 打印到标准错误流中所有已定义命令行标志的默认值。

#### func Set 

``` go
func Set(name, value string) error
```

​	Set 函数用于设置指定名称的命令行标志的值。

#### func SetInterspersed 

``` go
func SetInterspersed(interspersed bool)
```

​	SetInterspersed 函数用于设置是否支持交错的选项和非选项实参。

#### func String 

``` go
func String(name string, value string, usage string) *string
```

​	String 函数定义一个具有指定名称、默认值和用法说明的字符串标志。返回值是存储标志值的字符串变量的地址。

#### func StringArray 

``` go
func StringArray(name string, value []string, usage string) *[]string
```

​	StringArray 函数定义一个具有指定名称、默认值和用法说明的字符串数组标志。返回值是存储标志值的 []string 变量的地址。每个参数的值不会尝试用逗号分隔。对于这种情况，请使用 StringSlice函数。

#### func StringArrayP 

``` go
func StringArrayP(name, shorthand string, value []string, usage string) *[]string
```

​	StringArrayP 函数与 StringArray 函数类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringArrayVar 

``` go
func StringArrayVar(p *[]string, name string, value []string, usage string)
```

​	StringArrayVar 函数定义一个具有指定名称、默认值和用法说明的字符串数组标志。参数 p 指向一个 []string 变量，用于存储标志值。每个参数的值不会尝试用逗号分隔。对于这种情况，请使用 StringSlice。

#### func StringArrayVarP 

``` go
func StringArrayVarP(p *[]string, name, shorthand string, value []string, usage string)
```

​	StringArrayVarP 与 StringArrayVar 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringP 

``` go
func StringP(name, shorthand string, value string, usage string) *string
```

​	StringP 函数与 String 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringSlice 

``` go
func StringSlice(name string, value []string, usage string) *[]string
```

​	StringSlice 函数定义一个具有指定名称、默认值和用法说明的字符串切片标志。返回值是存储标志值的 []string 变量的地址。与 StringArray 标志不同，StringSlice 标志将以逗号分隔的值作为参数，并相应地进行拆分。例如：

```
--ss="v1,v2" --ss="v3"
```

将导致

```
[]string{"v1", "v2", "v3"}
```

#### func StringSliceP 

``` go
func StringSliceP(name, shorthand string, value []string, usage string) *[]string
```

​	StringSliceP 与 StringSlice 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringSliceVar 

``` go
func StringSliceVar(p *[]string, name string, value []string, usage string)
```

StringSliceVar 函数定义一个具有指定名称、默认值和用法说明的字符串切片标志。参数 p 指向一个 []string 变量，用于存储标志值。与 StringArray 标志不同，StringSlice 标志将以逗号分隔的值作为参数，并相应地进行拆分。例如：

```
--ss="v1,v2" --ss="v3"
```

将导致

```
[]string{"v1", "v2", "v3"}
```

#### func StringSliceVarP 

``` go
func StringSliceVarP(p *[]string, name, shorthand string, value []string, usage string)
```

​	StringSliceVarP 函数与 StringSliceVar 函数类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringToInt  <- v1.0.3

``` go
func StringToInt(name string, value map[string]int, usage string) *map[string]int
```

​	StringToInt 函数定义一个具有指定名称、默认值和用法说明的 map[string]int 标志。返回值是存储标志值的 map[string]int 变量的地址。每个参数的值不会尝试用逗号分隔。

#### func StringToInt64  <- v1.0.5

``` go
func StringToInt64(name string, value map[string]int64, usage string) *map[string]int64
```

​	StringToInt64 函数定义一个具有指定名称、默认值和用法说明的 map[string]int64 标志。返回值是存储标志值的 map[string]int64 变量的地址。每个参数的值不会尝试用逗号分隔。

#### func StringToInt64P  <- v1.0.5

``` go
func StringToInt64P(name, shorthand string, value map[string]int64, usage string) *map[string]int64
```

​	StringToInt64P 与 StringToInt64 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringToInt64Var  <- v1.0.5

``` go
func StringToInt64Var(p *map[string]int64, name string, value map[string]int64, usage string)
```

​	StringToInt64Var 函数定义一个具有指定名称、默认值和用法说明的 map[string]int64 标志。参数 p 指向一个 map[string]int64 变量，用于存储标志值。每个参数的值不会尝试用逗号分隔。

#### func StringToInt64VarP  <- v1.0.5

``` go
func StringToInt64VarP(p *map[string]int64, name, shorthand string, value map[string]int64, usage string)
```

​	StringToInt64VarP 与 StringToInt64Var 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringToIntP  <- v1.0.3

``` go
func StringToIntP(name, shorthand string, value map[string]int, usage string) *map[string]int
```

​	StringToIntP 与 StringToInt 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringToIntVar  <- v1.0.3

``` go
func StringToIntVar(p *map[string]int, name string, value map[string]int, usage string)
```

​	StringToIntVar 函数定义一个具有指定名称、默认值和用法说明的 map[string]int 标志。参数 p 指向一个 map[string]int 变量，用于存储标志值。每个参数的值不会尝试用逗号分隔。

#### func StringToIntVarP  <- v1.0.3

``` go
func StringToIntVarP(p *map[string]int, name, shorthand string, value map[string]int, usage string)
```

​	StringToIntVarP 与 StringToIntVar 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringToString  <- v1.0.3

``` go
func StringToString(name string, value map[string]string, usage string) *map[string]string
```

​	StringToString 函数定义一个具有指定名称、默认值和用法说明的 map[string]string 标志。返回值是存储标志值的 map[string]string 变量的地址。每个参数的值不会尝试用逗号分隔。

#### func StringToStringP  <- v1.0.3

``` go
func StringToStringP(name, shorthand string, value map[string]string, usage string) *map[string]string
```

​	StringToStringP 与 StringToString 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringToStringVar  <- v1.0.3

``` go
func StringToStringVar(p *map[string]string, name string, value map[string]string, usage string)
```

​	StringToStringVar 函数定义一个具有指定名称、默认值和用法说明的 map[string]string 标志。参数 p 指向一个 map[string]string 变量，用于存储标志值。每个参数的值不会尝试用逗号分隔。

#### func StringToStringVarP  <- v1.0.3

``` go
func StringToStringVarP(p *map[string]string, name, shorthand string, value map[string]string, usage string)
```

​	StringToStringVarP 与 StringToStringVar 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func StringVar 

``` go
func StringVar(p *string, name string, value string, usage string)
```

​	StringVar 函数定义一个具有指定名称、默认值和用法说明的字符串标志。参数 p 指向一个字符串变量，用于存储标志值。

#### func StringVarP 

``` go
func StringVarP(p *string, name, shorthand string, value string, usage string)
```

​	StringVarP 函数与 StringVar 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Uint 

``` go
func Uint(name string, value uint, usage string) *uint
```

​	Uint 函数定义一个具有指定名称、默认值和用法说明的无符号整数标志。返回值是存储标志值的无符号整数变量的地址。

#### func Uint16 

``` go
func Uint16(name string, value uint16, usage string) *uint16
```

​	Uint16 函数定义一个具有指定名称、默认值和用法说明的 uint16 标志。返回值是存储标志值的 uint16 变量的地址。

#### func Uint16P 

``` go
func Uint16P(name, shorthand string, value uint16, usage string) *uint16
```

​	Uint16P 函数与 Uint16 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Uint16Var 

``` go
func Uint16Var(p *uint16, name string, value uint16, usage string)
```

​	Uint16Var 函数定义一个具有指定名称、默认值和用法说明的 uint16 标志。参数 p 指向一个 uint16 变量，用于存储标志值。

#### func Uint16VarP 

``` go
func Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string)
```

​	Uint16VarP 函数与 Uint16Var 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Uint32 

``` go
func Uint32(name string, value uint32, usage string) *uint32
```

​	Uint32 函数定义一个具有指定名称、默认值和用法说明的 uint32 标志。返回值是存储标志值的 uint32 变量的地址。

#### func Uint32P 

``` go
func Uint32P(name, shorthand string, value uint32, usage string) *uint32
```

​	Uint32P 函数与 Uint32 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Uint32Var 

``` go
func Uint32Var(p *uint32, name string, value uint32, usage string)
```

​	Uint32Var 函数定义一个具有指定名称、默认值和用法说明的 uint32 标志。参数 p 指向一个 uint32 变量，用于存储标志值。

#### func Uint32VarP 

``` go
func Uint32VarP(p *uint32, name, shorthand string, value uint32, usage string)
```

​	Uint32VarP 函数与 Uint32Var 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Uint64 

``` go
func Uint64(name string, value uint64, usage string) *uint64
```

​	Uint64 函数定义一个具有指定名称、默认值和用法说明的 uint64 标志。返回值是存储标志值的 uint64 变量的地址。

#### func Uint64P 

``` go
func Uint64P(name, shorthand string, value uint64, usage string) *uint64
```

​	Uint64P 函数与 Uint64 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Uint64Var 

``` go
func Uint64Var(p *uint64, name string, value uint64, usage string)
```

​	Uint64Var 函数定义一个具有指定名称、默认值和用法说明的 uint64 标志。参数 p 指向一个 uint64 变量，用于存储标志值。

#### func Uint64VarP 

``` go
func Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string)
```

​	Uint64VarP 函数与 Uint64Var 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Uint8 

``` go
func Uint8(name string, value uint8, usage string) *uint8
```

​	Uint8 函数定义一个具有指定名称、默认值和用法说明的 uint8 标志。返回值是存储标志值的 uint8 变量的地址。

#### func Uint8P 

``` go
func Uint8P(name, shorthand string, value uint8, usage string) *uint8
```

​	Uint8P 函数与 Uint8 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Uint8Var 

``` go
func Uint8Var(p *uint8, name string, value uint8, usage string)
```

​	Uint8Var 函数定义一个具有指定名称、默认值和用法说明的 uint8 标志。参数 p 指向一个 uint8 变量，用于存储标志值。

#### func Uint8VarP 

``` go
func Uint8VarP(p *uint8, name, shorthand string, value uint8, usage string)
```

​	Uint8VarP 函数与 Uint8Var 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func UintP 

``` go
func UintP(name, shorthand string, value uint, usage string) *uint
```

​	UintP 函数与 Uint 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func UintSlice 

``` go
func UintSlice(name string, value []uint, usage string) *[]uint
```

​	UintSlice 函数定义一个具有指定名称、默认值和用法说明的 `[]uint` 标志。返回值是存储标志值的 `[]uint` 变量的地址。

#### func UintSliceP 

``` go
func UintSliceP(name, shorthand string, value []uint, usage string) *[]uint
```

​	UintSliceP 函数与 UintSlice 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func UintSliceVar 

``` go
func UintSliceVar(p *[]uint, name string, value []uint, usage string)
```

​	UintSliceVar 函数定义一个具有指定名称、默认值和用法说明的 `[]uint` 标志。参数 p 指向一个 `[]uint` 变量，用于存储标志值。

#### func UintSliceVarP 

``` go
func UintSliceVarP(p *[]uint, name, shorthand string, value []uint, usage string)
```

​	UintSliceVarP 函数与 UintSliceVar 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func UintVar 

``` go
func UintVar(p *uint, name string, value uint, usage string)
```

​	UintVar 函数定义一个具有指定名称、默认值和用法说明的无符号整数标志。参数 p 指向一个无符号整数变量，用于存储标志值。

#### func UintVarP 

``` go
func UintVarP(p *uint, name, shorthand string, value uint, usage string)
```

​	UintVarP 函数与 UintVar 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func UnquoteUsage 

``` go
func UnquoteUsage(flag *Flag) (name string, usage string)
```

​	UnquoteUsage 函数从标志的用法说明字符串中提取用反引号括起的名称，并返回名称和未引用的用法说明。对于给定的 "a `name` to show"，它返回 ("name", "a name to show")。如果没有反引号，则名称是标志值类型的有教育意义的猜测，如果标志是布尔值，则为空字符串。

#### func Var 

``` go
func Var(value Value, name string, usage string)
```

​	Var 函数通过第一个参数的 Value 类型来定义具有指定名称和用法说明的标志。Value 类型通常包含用户定义的 Value 实现，通过该实现可以将逗号分隔的字符串转换为切片。例如，调用者可以创建一个标志，通过给切片赋予 Value 方法，将逗号分隔的字符串转换为切片。

#### func VarP 

``` go
func VarP(value Value, name, shorthand, usage string)
```

VarP is like Var, but accepts a shorthand letter that can be used after a single dash.

​	VarP 函数与 Var 类似，但接受一个速记字母，该字母可以在单个短划线后使用。

#### func Visit 

``` go
func Visit(fn func(*Flag))
```

​	Visit 函数按照词法顺序或原始顺序（如果 f.SortFlags 为 false）访问命令行标志，对每个标志调用提供的函数 fn。它仅访问已设置的标志。

#### func VisitAll 

``` go
func VisitAll(fn func(*Flag))
```

​	VisitAll 函数按照词法顺序或原始顺序（如果 f.SortFlags 为 false）访问命令行标志，对每个标志调用提供的函数 fn。它访问所有标志，包括未设置的标志。

### 类型 

#### type ErrorHandling 

``` go
type ErrorHandling int
```

​	ErrorHandling 定义了如何处理标志解析错误。

``` go
const (
     // ContinueOnError 在发现错误时从 Parse() 返回一个 err
	ContinueOnError ErrorHandling = iota
    // ExitOnError 在发现错误时调用 os.Exit(2)
	ExitOnError
    // PanicOnError 在发现错误时调用 panic()
	PanicOnError
)
```

#### type Flag 

``` go
type Flag struct {
	Name                string              // name as it appears on command line 在命令行上显示的名称
	Shorthand           string              // one-letter abbreviated flag 一个字母的缩写标志
	Usage               string              // help message 帮助消息
	Value               Value               // value as set 设置的值
	DefValue            string              // default value (as text); for usage message 默认值（作为文本）；用于用法消息
	Changed             bool                // If the user set the value (or if left to default) 如果用户设置了值（或使用默认值），则为 true
	NoOptDefVal         string              // default value (as text); if the flag is on the command line without any options 默认值（作为文本）；如果标志在命令行上没有任何选项，但出现在命令行上
	Deprecated          string              // If this flag is deprecated, this string is the new or now thing to use 如果此标志已被弃用，则该字符串是新的或现在使用的内容
	Hidden              bool                // used by cobra.Command to allow flags to be hidden from help/usage text 由 cobra.Command 使用，允许将标志从帮助/用法文本中隐藏
	ShorthandDeprecated string              // If the shorthand of this flag is deprecated, this string is the new or now thing to use 如果此标志的缩写已被弃用，则该字符串是新的或现在使用的内容
	Annotations         map[string][]string // used by cobra.Command bash autocomple code 由 cobra.Command bash autocomple 代码使用
}
```

Flag 表示标志的状态。

##### func Lookup 

``` go
func Lookup(name string) *Flag
```

​	Lookup 返回指定命令行标志的 Flag 结构，如果不存在则返回 nil。

##### func PFlagFromGoFlag 

``` go
func PFlagFromGoFlag(goflag *goflag.Flag) *Flag
```

​	PFlagFromGoFlag 将给定的 *flag.Flag 转换为 *pflag.Flag。如果 *flag.Flag.Name 是一个单个字符（例如 `v`），则可以在标志中使用 `-v` 和 `--v`。如果 golang 标志不止一个字符（例如 `verbose`），则只能通过 `--verbose` 使用。

##### func ShorthandLookup 

``` go
func ShorthandLookup(name string) *Flag
```

​	ShorthandLookup 返回指定缩写标志的 Flag 结构，如果不存在则返回 nil。

###### ShorthandLookup Example

```go
package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	name := "verbose"
	short := name[:1]

	pflag.BoolP(name, short, false, "verbose output")

	// len(short) must be == 1
	flag := pflag.ShorthandLookup(short)

	fmt.Println(flag.Name)
}

```



##### type FlagSet 

``` go
type FlagSet struct {
    // 在解析标志时发生错误时调用的函数。
    // 该字段是一个函数（不是方法），可以更改以指向自定义的错误处理程序。
	Usage func()

    // 用于指示用户是否希望在帮助/用法消息中对标志进行排序。
	SortFlags bool

    // 用于配置错误白名单的 ParseErrorsWhitelist。
	ParseErrorsWhitelist ParseErrorsWhitelist
	// 包含过滤或未导出的字段
}
```

​	FlagSet 表示一组定义的标志。

##### func NewFlagSet 

``` go
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
```

​	NewFlagSet 函数返回一个新的、空的带有指定名称和错误处理属性的 FlagSet，SortFlags 设置为 true。

##### (*FlagSet) AddFlag 

``` go
func (f *FlagSet) AddFlag(flag *Flag)
```

​	AddFlag 方法将标志添加到 FlagSet。

##### (*FlagSet) AddFlagSet 

``` go
func (f *FlagSet) AddFlagSet(newSet *FlagSet)
```

​	AddFlagSet 将一个 FlagSet 添加到另一个 FlagSet。如果标志已经存在于 f 中，则忽略来自 newSet 的标志。

##### (*FlagSet) AddGoFlag 

``` go
func (f *FlagSet) AddGoFlag(goflag *goflag.Flag)
```

​	AddGoFlag 将给定的 *flag.Flag 添加到 pflag.FlagSet。

##### (*FlagSet) AddGoFlagSet 

``` go
func (f *FlagSet) AddGoFlagSet(newSet *goflag.FlagSet)
```

​	AddGoFlagSet 将给定的 *flag.FlagSet 添加到 pflag.FlagSet。

##### (*FlagSet) Arg 

``` go
func (f *FlagSet) Arg(i int) string
```

​	Arg 返回第 i 个参数。Arg(0) 是在处理标志后剩下的第一个参数。

##### (*FlagSet) Args 

``` go
func (f *FlagSet) Args() []string
```

​	Args 返回非标志参数。

##### (*FlagSet) ArgsLenAtDash 

``` go
func (f *FlagSet) ArgsLenAtDash() int
```

​	ArgsLenAtDash 返回在参数解析过程中找到 `--` 时 f.Args 的长度。这允许程序知道哪些参数在 `--` 之前，哪些参数在 `--` 之后。

##### (*FlagSet) Bool 

``` go
func (f *FlagSet) Bool(name string, value bool, usage string) *bool
```

​	Bool 定义具有指定名称、默认值和用法说明的布尔标志。返回值是存储标志值的 bool 变量的地址。

##### (*FlagSet) BoolP 

``` go
func (f *FlagSet) BoolP(name, shorthand string, value bool, usage string) *bool
```

​	BoolP 类似于 Bool，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) BoolSlice 

``` go
func (f *FlagSet) BoolSlice(name string, value []bool, usage string) *[]bool
```

​	BoolSlice 定义具有指定名称、默认值和用法说明的 []bool 标志。返回值是存储标志值的 []bool 变量的地址。

##### (*FlagSet) BoolSliceP 

``` go
func (f *FlagSet) BoolSliceP(name, shorthand string, value []bool, usage string) *[]bool
```

​	BoolSliceP 类似于 BoolSlice，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) BoolSliceVar 

``` go
func (f *FlagSet) BoolSliceVar(p *[]bool, name string, value []bool, usage string)
```

​	BoolSliceVar 定义具有指定名称、默认值和用法说明的 boolSlice 标志。参数 p 指向一个 []bool 变量，用于存储标志值。

##### (*FlagSet) BoolSliceVarP 

``` go
func (f *FlagSet) BoolSliceVarP(p *[]bool, name, shorthand string, value []bool, usage string)
```

​	BoolSliceVarP 类似于 BoolSliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) BoolVar 

``` go
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
```

​	BoolVar 定义具有指定名称、默认值和用法说明的 bool 标志。参数 p 指向一个 bool 变量，用于存储标志值。

##### (*FlagSet) BoolVarP 

``` go
func (f *FlagSet) BoolVarP(p *bool, name, shorthand string, value bool, usage string)
```

​	BoolVarP 类似于 BoolVar，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) BytesBase64  <- v1.0.2

``` go
func (f *FlagSet) BytesBase64(name string, value []byte, usage string) *[]byte
```

​	BytesBase64 定义具有指定名称、默认值和用法说明的 []byte 标志。返回值是存储标志值的 []byte 变量的地址。

##### (*FlagSet) BytesBase64P  <- v1.0.2

``` go
func (f *FlagSet) BytesBase64P(name, shorthand string, value []byte, usage string) *[]byte
```

​	BytesBase64P 类似于 BytesBase64，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) BytesBase64Var  <- v1.0.2

``` go
func (f *FlagSet) BytesBase64Var(p *[]byte, name string, value []byte, usage string)
```

​	BytesBase64Var 定义具有指定名称、默认值和用法说明的 []byte 标志。参数 p 指向一个 []byte 变量，用于存储标志值。

##### (*FlagSet) BytesBase64VarP  <- v1.0.2

``` go
func (f *FlagSet) BytesBase64VarP(p *[]byte, name, shorthand string, value []byte, usage string)
```

​	BytesBase64VarP 类似于 BytesBase64Var，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) BytesHex  <- v1.0.1

``` go
func (f *FlagSet) BytesHex(name string, value []byte, usage string) *[]byte
```

​	BytesHex 定义具有指定名称、默认值和用法说明的 []byte 标志。返回值是存储标志值的 []byte 变量的地址。

##### (*FlagSet) BytesHexP  <- v1.0.1

``` go
func (f *FlagSet) BytesHexP(name, shorthand string, value []byte, usage string) *[]byte
```

​	BytesHexP 类似于 BytesHex，但接受一个速记字母，该字母可以在单个短划线后使用。



##### (*FlagSet) BytesHexVar  <- v1.0.1

``` go
func (f *FlagSet) BytesHexVar(p *[]byte, name string, value []byte, usage string)
```

​	BytesHexVar 定义具有指定名称、默认值和用法说明的 []byte 标志。参数 p 指向一个 []byte 变量，用于存储标志值。

##### (*FlagSet) BytesHexVarP  <- v1.0.1

``` go
func (f *FlagSet) BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string)
```

​	BytesHexVarP 类似于 BytesHexVar，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) Changed 

``` go
func (f *FlagSet) Changed(name string) bool
```

​	Changed 返回 true，如果标志在 Parse() 期间被显式设置，否则返回 false。

##### (*FlagSet) Count 

``` go
func (f *FlagSet) Count(name string, usage string) *int
```

​	Count 定义具有指定名称、默认值和用法说明的计数标志。返回值是存储标志值的 int 变量的地址。计数标志会在每次在命令行上找到它时将其值加 1。

##### (*FlagSet) CountP 

``` go
func (f *FlagSet) CountP(name, shorthand string, usage string) *int
```

​	CountP 类似于 Count，只接受缩写标志名称。

##### (*FlagSet) CountVar 

``` go
func (f *FlagSet) CountVar(p *int, name string, usage string)
```

​	CountVar 定义具有指定名称、默认值和用法说明的计数标志。参数 p 指向一个 int 变量，用于存储标志值。计数标志会在每次在命令行上找到它时将其值加 1。

##### (*FlagSet) CountVarP 

``` go
func (f *FlagSet) CountVarP(p *int, name, shorthand string, usage string)
```

​	CountVarP 类似于 CountVar，只接受缩写标志名称。

##### (*FlagSet) Duration 

``` go
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
```

​	Duration 定义具有指定名称、默认值和用法说明的 time.Duration 标志。返回值是存储标志值的 time.Duration 变量的地址。

##### (*FlagSet) DurationP 

``` go
func (f *FlagSet) DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration
```

​	DurationP 类似于 Duration，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) DurationSlice  <- v1.0.1

``` go
func (f *FlagSet) DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration
```

​	DurationSlice 定义具有指定名称、默认值和用法说明的 []time.Duration 标志。返回值是存储标志值的 []time.Duration 变量的地址。

##### (*FlagSet) DurationSliceP  <- v1.0.1

``` go
func (f *FlagSet) DurationSliceP(name, shorthand string, value []time.Duration, usage string) *[]time.Duration
```

​	DurationSliceP 类似于 DurationSlice，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) DurationSliceVar  <- v1.0.1

``` go
func (f *FlagSet) DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string)
```

​	DurationSliceVar 定义具有指定名称、默认值和用法说明的 durationSlice 标志。参数 p 指向一个 []time.Duration 变量，用于存储标志值。

##### (*FlagSet) DurationSliceVarP  <- v1.0.1

``` go
func (f *FlagSet) DurationSliceVarP(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string)
```

​	DurationSliceVarP 类似于 DurationSliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) DurationVar 

``` go
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

​	DurationVar 定义具有指定名称、默认值和用法说明的 time.Duration 标志。参数 p 指向一个 time.Duration 变量，用于存储标志值。

##### (*FlagSet) DurationVarP 

``` go
func (f *FlagSet) DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string)
```

​	DurationVarP 类似于 DurationVar，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) FlagUsages 

``` go
func (f *FlagSet) FlagUsages() string
```

​	FlagUsages 返回一个包含 FlagSet 中所有标志的用法信息的字符串。

##### (*FlagSet) FlagUsagesWrapped 

``` go
func (f *FlagSet) FlagUsagesWrapped(cols int) string
```

​	FlagUsagesWrapped 返回一个包含 FlagSet 中所有标志的用法信息的字符串。用 cols 列进行换行（0 表示不换行）。

##### (*FlagSet) Float32 

``` go
func (f *FlagSet) Float32(name string, value float32, usage string) *float32
```

​	Float32 定义具有指定名称、默认值和用法说明的 float32 标志。返回值是存储标志值的 float32 变量的地址。

##### (*FlagSet) Float32P 

``` go
func (f *FlagSet) Float32P(name, shorthand string, value float32, usage string) *float32
```

​	Float32P 类似于 Float32，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) Float32Slice  <- v1.0.5

``` go
func (f *FlagSet) Float32Slice(name string, value []float32, usage string) *[]float32
```

​	Float32Slice 定义具有指定名称、默认值和用法说明的 []float32 标志。返回值是存储标志值的 []float32 变量的地址。

##### (*FlagSet) Float32SliceP  <- v1.0.5

``` go
func (f *FlagSet) Float32SliceP(name, shorthand string, value []float32, usage string) *[]float32
```

​	Float32SliceP 类似于 Float32Slice，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) Float32SliceVar  <- v1.0.5

``` go
func (f *FlagSet) Float32SliceVar(p *[]float32, name string, value []float32, usage string)
```

​	Float32SliceVar 定义具有指定名称、默认值和用法说明的 float32Slice 标志。参数 p 指向一个 []float32 变量，用于存储标志值。

##### (*FlagSet) Float32SliceVarP  <- v1.0.5

``` go
func (f *FlagSet) Float32SliceVarP(p *[]float32, name, shorthand string, value []float32, usage string)
```

​	Float32SliceVarP 类似于 Float32SliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) Float32Var 

``` go
func (f *FlagSet) Float32Var(p *float32, name string, value float32, usage string)
```

​	Float32Var 定义具有指定名称、默认值和用法说明的 float32 标志。参数 p 指向一个 float32 变量，用于存储标志值。

##### (*FlagSet) Float32VarP 

``` go
func (f *FlagSet) Float32VarP(p *float32, name, shorthand string, value float32, usage string)
```

​	Float32VarP 类似于 Float32Var，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) Float64 

``` go
func (f *FlagSet) Float64(name string, value float64, usage string) *float64
```

​	Float64 定义具有指定名称、默认值和用法说明的 float64 标志。返回值是存储标志值的 float64 变量的地址。

##### (*FlagSet) Float64P 

``` go
func (f *FlagSet) Float64P(name, shorthand string, value float64, usage string) *float64
```

​	Float64P 类似于 Float64，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) Float64Slice  <- v1.0.5

``` go
func (f *FlagSet) Float64Slice(name string, value []float64, usage string) *[]float64
```

​	Float64Slice 定义具有指定名称、默认值和用法说明的 []float64 标志。返回值是存储标志值的 []float64 变量的地址。

##### (*FlagSet) Float64SliceP  <- v1.0.5

``` go
func (f *FlagSet) Float64SliceP(name, shorthand string, value []float64, usage string) *[]float64
```

​	Float64SliceP 类似于 Float64Slice，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) Float64SliceVar  <- v1.0.5

``` go
func (f *FlagSet) Float64SliceVar(p *[]float64, name string, value []float64, usage string)
```

​	Float64SliceVar 定义具有指定名称、默认值和用法说明的 float64Slice 标志。参数 p 指向一个 []float64 变量，用于存储标志值。

##### (*FlagSet) Float64SliceVarP  <- v1.0.5

``` go
func (f *FlagSet) Float64SliceVarP(p *[]float64, name, shorthand string, value []float64, usage string)
```

​	Float64SliceVarP 类似于 Float64SliceVar，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) Float64Var 

``` go
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
```

​	Float64Var 定义具有指定名称、默认值和用法说明的 float64 标志。参数 p 指向一个 float64 变量，用于存储标志值。

##### (*FlagSet) Float64VarP 

``` go
func (f *FlagSet) Float64VarP(p *float64, name, shorthand string, value float64, usage string)
```

​	Float64VarP 类似于 Float64Var，但接受一个速记字母，该字母可以在单个短划线后使用。

##### (*FlagSet) GetBool 

``` go
func (f *FlagSet) GetBool(name string) (bool, error)
```

​	GetBool 返回具有指定名称的 bool 标志的值。

##### (*FlagSet) GetBoolSlice 

``` go
func (f *FlagSet) GetBoolSlice(name string) ([]bool, error)
```

​	GetBoolSlice 返回具有指定名称的 []bool 标志的值。

##### (*FlagSet) GetBytesBase64  <- v1.0.2

``` go
func (f *FlagSet) GetBytesBase64(name string) ([]byte, error)
```

​	GetBytesBase64 返回具有给定名称的标志的 []byte 值。

##### (*FlagSet) GetBytesHex  <- v1.0.1

``` go
func (f *FlagSet) GetBytesHex(name string) ([]byte, error)
```

​	GetBytesHex 返回具有给定名称的标志的 []byte 值。

##### (*FlagSet) GetCount 

``` go
func (f *FlagSet) GetCount(name string) (int, error)
```

​	GetCount 返回具有给定名称的标志的 int 值。

##### (*FlagSet) GetDuration 

``` go
func (f *FlagSet) GetDuration(name string) (time.Duration, error)
```

​	GetDuration 返回具有给定名称的标志的 time.Duration 值。

##### (*FlagSet) GetDurationSlice  <- v1.0.1

``` go
func (f *FlagSet) GetDurationSlice(name string) ([]time.Duration, error)
```

​	GetDurationSlice 返回具有给定名称的标志的 []time.Duration 值。

##### (*FlagSet) GetFloat32 

``` go
func (f *FlagSet) GetFloat32(name string) (float32, error)
```

​	GetFloat32 返回具有给定名称的标志的 float32 值。

##### (*FlagSet) GetFloat32Slice  <- v1.0.5

``` go
func (f *FlagSet) GetFloat32Slice(name string) ([]float32, error)
```

​	GetFloat32Slice 返回具有给定名称的标志的 []float32 值。

##### (*FlagSet) GetFloat64 

``` go
func (f *FlagSet) GetFloat64(name string) (float64, error)
```

​	GetFloat64 返回具有给定名称的标志的 float64 值。

##### (*FlagSet) GetFloat64Slice  <- v1.0.5

``` go
func (f *FlagSet) GetFloat64Slice(name string) ([]float64, error)
```

​	GetFloat64Slice 返回具有给定名称的标志的 []float64 值。

##### (*FlagSet) GetIP 

``` go
func (f *FlagSet) GetIP(name string) (net.IP, error)
```

​	GetIP 返回具有给定名称的标志的 net.IP 值。

##### (*FlagSet) GetIPNet 

``` go
func (f *FlagSet) GetIPNet(name string) (net.IPNet, error)
```

​	GetIPNet 返回具有给定名称的标志的 net.IPNet 值。

##### (*FlagSet) GetIPSlice 

``` go
func (f *FlagSet) GetIPSlice(name string) ([]net.IP, error)
```

​	GetIPSlice 返回具有给定名称的标志的 []net.IP 值。

##### (*FlagSet) GetIPv4Mask 

``` go
func (f *FlagSet) GetIPv4Mask(name string) (net.IPMask, error)
```

​	GetIPv4Mask 返回具有给定名称的标志的 net.IPv4Mask 值。

##### (*FlagSet) GetInt 

``` go
func (f *FlagSet) GetInt(name string) (int, error)
```

​	GetInt 返回具有给定名称的标志的 int 值。

##### (*FlagSet) GetInt16  <- v1.0.1

``` go
func (f *FlagSet) GetInt16(name string) (int16, error)
```

​	GetInt16 返回具有给定名称的标志的 int16 值。

##### (*FlagSet) GetInt32 

``` go
func (f *FlagSet) GetInt32(name string) (int32, error)
```

​	GetInt32 返回具有给定名称的标志的 int32 值。

##### (*FlagSet) GetInt32Slice  <- v1.0.5

``` go
func (f *FlagSet) GetInt32Slice(name string) ([]int32, error)
```

​	GetInt32Slice 返回具有给定名称的标志的 []int32 值。

##### (*FlagSet) GetInt64 

``` go
func (f *FlagSet) GetInt64(name string) (int64, error)
```

​	GetInt64 返回具有给定名称的标志的 int64 值。

##### (*FlagSet) GetInt64Slice  <- v1.0.5

``` go
func (f *FlagSet) GetInt64Slice(name string) ([]int64, error)
```

​	GetInt64Slice 返回具有给定名称的标志的 []int64 值。

##### (*FlagSet) GetInt8 

``` go
func (f *FlagSet) GetInt8(name string) (int8, error)
```

​	GetInt8 返回具有给定名称的标志的 int8 值。

##### (*FlagSet) GetIntSlice 

``` go
func (f *FlagSet) GetIntSlice(name string) ([]int, error)
```

​	GetIntSlice 返回具有给定名称的标志的 []int 值。

##### (*FlagSet) GetNormalizeFunc 

``` go
func (f *FlagSet) GetNormalizeFunc() func(f *FlagSet, name string) NormalizedName
```

​	GetNormalizeFunc 返回先前设置的 NormalizeFunc 函数，该函数不进行翻译，如果以前未设置。

##### (*FlagSet) GetString 

``` go
func (f *FlagSet) GetString(name string) (string, error)
```

​	GetString 返回具有给定名称的标志的 string 值。

##### (*FlagSet) GetStringArray 

``` go
func (f *FlagSet) GetStringArray(name string) ([]string, error)
```

​	GetStringArray 返回具有给定名称的标志的 []string 值。

##### (*FlagSet) GetStringSlice 

``` go
func (f *FlagSet) GetStringSlice(name string) ([]string, error)
```

​	GetStringSlice 返回具有给定名称的标志的 []string 值。

##### (*FlagSet) GetStringToInt  <- v1.0.3

``` go
func (f *FlagSet) GetStringToInt(name string) (map[string]int, error)
```

​	GetStringToInt 返回具有给定名称的标志的 map[string]int 值。

##### (*FlagSet) GetStringToInt64  <- v1.0.5

``` go
func (f *FlagSet) GetStringToInt64(name string) (map[string]int64, error)
```

​	GetStringToInt64 返回具有给定名称的标志的 map[string]int64 值。

##### (*FlagSet) GetStringToString  <- v1.0.3

``` go
func (f *FlagSet) GetStringToString(name string) (map[string]string, error)
```

​	GetStringToString 返回具有给定名称的标志的 map[string]string 值。

##### (*FlagSet) GetUint 

``` go
func (f *FlagSet) GetUint(name string) (uint, error)
```

​	GetUint 返回给定名称的标志的 uint 值。

##### (*FlagSet) GetUint16 

``` go
func (f *FlagSet) GetUint16(name string) (uint16, error)
```

​	GetUint16 返回给定名称的标志的 uint16 值。

##### (*FlagSet) GetUint32 

``` go
func (f *FlagSet) GetUint32(name string) (uint32, error)
```

​	GetUint32 返回给定名称的标志的 uint32 值。

##### (*FlagSet) GetUint64 

``` go
func (f *FlagSet) GetUint64(name string) (uint64, error)
```

​	GetUint64 返回给定名称的标志的 uint64 值。

##### (*FlagSet) GetUint8 

``` go
func (f *FlagSet) GetUint8(name string) (uint8, error)
```

​	GetUint8 返回给定名称的标志的 uint8 值。

##### (*FlagSet) GetUintSlice 

``` go
func (f *FlagSet) GetUintSlice(name string) ([]uint, error)
```

​	GetUintSlice 返回给定名称的标志的 []uint 值。

##### (*FlagSet) HasAvailableFlags 

``` go
func (f *FlagSet) HasAvailableFlags() bool
```

​	HasAvailableFlags 返回一个布尔值，指示 FlagSet 是否有任何未隐藏的标志。

##### (*FlagSet) HasFlags 

``` go
func (f *FlagSet) HasFlags() bool
```

​	HasFlags 返回一个布尔值，指示 FlagSet 是否有任何已定义的标志。

##### (*FlagSet) IP 

``` go
func (f *FlagSet) IP(name string, value net.IP, usage string) *net.IP
```

​	IP 使用指定的名称、默认值和用法字符串来定义 net.IP 标志。返回值是一个 net.IP 变量的地址，用于存储标志的值。

##### (*FlagSet) IPMask 

``` go
func (f *FlagSet) IPMask(name string, value net.IPMask, usage string) *net.IPMask
```

​	IPMask 使用指定的名称、默认值和用法字符串来定义 net.IPMask 标志。返回值是一个 net.IPMask 变量的地址，用于存储标志的值。

##### (*FlagSet) IPMaskP 

``` go
func (f *FlagSet) IPMaskP(name, shorthand string, value net.IPMask, usage string) *net.IPMask
```

​	IPMaskP 类似于 IPMask，但接受一个快捷字母，可以在单个短划线后使用。

##### (*FlagSet) IPMaskVar 

``` go
func (f *FlagSet) IPMaskVar(p *net.IPMask, name string, value net.IPMask, usage string)
```

​	IPMaskVar 使用指定的名称、默认值和用法字符串来定义 net.IPMask 标志。参数 p 指向一个 net.IPMask 变量，用于存储标志的值。

##### (*FlagSet) IPMaskVarP 

``` go
func (f *FlagSet) IPMaskVarP(p *net.IPMask, name, shorthand string, value net.IPMask, usage string)
```

​	IPMaskVarP 类似于 IPMaskVar，但接受一个快捷字母，可以在单个短划线后使用。

##### (*FlagSet) IPNet 

``` go
func (f *FlagSet) IPNet(name string, value net.IPNet, usage string) *net.IPNet
```

​	IPNet 使用指定的名称、默认值和用法字符串来定义 net.IPNet 标志。返回值是一个 net.IPNet 变量的地址，用于存储标志的值。

##### (*FlagSet) IPNetP 

``` go
func (f *FlagSet) IPNetP(name, shorthand string, value net.IPNet, usage string) *net.IPNet
```

​	IPNetP 类似于 IPNet，但接受一个快捷字母，可以在单个短划线后使用。

##### (*FlagSet) IPNetVar 

``` go
func (f *FlagSet) IPNetVar(p *net.IPNet, name string, value net.IPNet, usage string)
```

​	IPNetVar 使用指定的名称、默认值和用法字符串来定义 net.IPNet 标志。参数 p 指向一个 net.IPNet 变量，用于存储标志的值。

##### (*FlagSet) IPNetVarP 

``` go
func (f *FlagSet) IPNetVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string)
```

​	IPNetVarP 类似于 IPNetVar，但接受一个快捷字母，可以在单个短划线后使用。

##### (*FlagSet) IPP 

``` go
func (f *FlagSet) IPP(name, shorthand string, value net.IP, usage string) *net.IP
```

​	IPP 类似于 IP，但接受一个快捷字母，可以在单个短划线后使用。

##### (*FlagSet) IPSlice 

``` go
func (f *FlagSet) IPSlice(name string, value []net.IP, usage string) *[]net.IP
```

​	IPSlice 使用指定的名称、默认值和用法字符串来定义 []net.IP 标志。返回值是一个 []net.IP 变量的地址，用于存储标志的值。

##### (*FlagSet) IPSliceP 

``` go
func (f *FlagSet) IPSliceP(name, shorthand string, value []net.IP, usage string) *[]net.IP
```

​	IPSliceP 类似于 IPSlice，但接受一个快捷字母，可以在单个短划线后使用。

##### (*FlagSet) IPSliceVar 

``` go
func (f *FlagSet) IPSliceVar(p *[]net.IP, name string, value []net.IP, usage string)
```

​	IPSliceVar 定义了一个名为 ipSlice 的标志，指定了默认值和用法说明。参数 p 指向一个 []net.IP 变量，用于存储标志的值。

##### (*FlagSet) IPSliceVarP 

``` go
func (f *FlagSet) IPSliceVarP(p *[]net.IP, name, shorthand string, value []net.IP, usage string)
```

​	IPSliceVarP 类似于 IPSliceVar，但允许在单个短划线后使用速记字母。

##### (*FlagSet) IPVar 

``` go
func (f *FlagSet) IPVar(p *net.IP, name string, value net.IP, usage string)
```

​	IPVar 定义了一个带有指定名称、默认值和用法说明的 net.IP 标志。参数 p 指向一个 net.IP 变量，用于存储标志的值。

##### (*FlagSet) IPVarP 

``` go
func (f *FlagSet) IPVarP(p *net.IP, name, shorthand string, value net.IP, usage string)
```

​	IPVarP 类似于 IPVar，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Init 

``` go
func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
```

​	Init 为标志集设置名称和错误处理属性。默认情况下，零值的 FlagSet 使用空名称和 ContinueOnError 错误处理策略。

##### (*FlagSet) Int 

``` go
func (f *FlagSet) Int(name string, value int, usage string) *int
```

​	Int 定义了一个带有指定名称、默认值和用法说明的 int 标志。返回值是一个 int 变量的地址，用于存储标志的值。

##### (*FlagSet) Int16  <- v1.0.1

``` go
func (f *FlagSet) Int16(name string, value int16, usage string) *int16
```

​	Int16 定义了一个带有指定名称、默认值和用法说明的 int16 标志。返回值是一个 int16 变量的地址，用于存储标志的值。

##### (*FlagSet) Int16P  <- v1.0.1

``` go
func (f *FlagSet) Int16P(name, shorthand string, value int16, usage string) *int16
```

​	Int16P 类似于 Int16，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int16Var  <- v1.0.1

``` go
func (f *FlagSet) Int16Var(p *int16, name string, value int16, usage string)
```

​	Int16Var 定义了一个带有指定名称、默认值和用法说明的 int16 标志。参数 p 指向一个 int16 变量，用于存储标志的值。

##### (*FlagSet) Int16VarP  <- v1.0.1

``` go
func (f *FlagSet) Int16VarP(p *int16, name, shorthand string, value int16, usage string)
```

​	Int16VarP 类似于 Int16Var，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int32 

``` go
func (f *FlagSet) Int32(name string, value int32, usage string) *int32
```

​	Int32 定义了一个带有指定名称、默认值和用法说明的 int32 标志。返回值是一个 int32 变量的地址，用于存储标志的值。

##### (*FlagSet) Int32P 

``` go
func (f *FlagSet) Int32P(name, shorthand string, value int32, usage string) *int32
```

​	Int32P 类似于 Int32，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int32Slice  <- v1.0.5

``` go
func (f *FlagSet) Int32Slice(name string, value []int32, usage string) *[]int32
```

​	Int32Slice 定义了一个带有指定名称、默认值和用法说明的 []int32 标志。返回值是一个 []int32 变量的地址，用于存储标志的值。

##### (*FlagSet) Int32SliceP  <- v1.0.5

``` go
func (f *FlagSet) Int32SliceP(name, shorthand string, value []int32, usage string) *[]int32
```

​	Int32SliceP 类似于 Int32Slice，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int32SliceVar  <- v1.0.5

``` go
func (f *FlagSet) Int32SliceVar(p *[]int32, name string, value []int32, usage string)
```

​	Int32SliceVar 定义了一个带有指定名称、默认值和用法说明的 int32Slice 标志。参数 p 指向一个 []int32 变量，用于存储标志的值。

##### (*FlagSet) Int32SliceVarP  <- v1.0.5

``` go
func (f *FlagSet) Int32SliceVarP(p *[]int32, name, shorthand string, value []int32, usage string)
```

​	Int32SliceVarP 类似于 Int32SliceVar，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int32Var 

``` go
func (f *FlagSet) Int32Var(p *int32, name string, value int32, usage string)
```

​	Int32Var 定义了一个带有指定名称、默认值和用法说明的 int32 标志。参数 p 指向一个 int32 变量，用于存储标志的值。

##### (*FlagSet) Int32VarP 

``` go
func (f *FlagSet) Int32VarP(p *int32, name, shorthand string, value int32, usage string)
```

​	Int32VarP 类似于 Int32Var，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int64 

``` go
func (f *FlagSet) Int64(name string, value int64, usage string) *int64
```

​	Int64 定义了一个带有指定名称、默认值和用法说明的 int64 标志。返回值是一个 int64 变量的地址，用于存储标志的值。

##### (*FlagSet) Int64P 

``` go
func (f *FlagSet) Int64P(name, shorthand string, value int64, usage string) *int64
```

​	Int64P 类似于 Int64，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int64Slice  <- v1.0.5

``` go
func (f *FlagSet) Int64Slice(name string, value []int64, usage string) *[]int64
```

​	Int64Slice 定义了一个带有指定名称、默认值和用法说明的 []int64 标志。返回值是一个 []int64 变量的地址，用于存储标志的值。

##### (*FlagSet) Int64SliceP  <- v1.0.5

``` go
func (f *FlagSet) Int64SliceP(name, shorthand string, value []int64, usage string) *[]int64
```

​	Int64SliceP 类似于 Int64Slice，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int64SliceVar  <- v1.0.5

``` go
func (f *FlagSet) Int64SliceVar(p *[]int64, name string, value []int64, usage string)
```

​	Int64SliceVar 定义了一个带有指定名称、默认值和用法说明的 int64Slice 标志。参数 p 指向一个 []int64 变量，用于存储标志的值。

##### (*FlagSet) Int64SliceVarP  <- v1.0.5

``` go
func (f *FlagSet) Int64SliceVarP(p *[]int64, name, shorthand string, value []int64, usage string)
```

​	Int64SliceVarP 类似于 Int64SliceVar，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int64Var 

``` go
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
```

​	Int64Var 定义了一个带有指定名称、默认值和用法说明的 int64 标志。参数 p 指向一个 int64 变量，用于存储标志的值。

##### (*FlagSet) Int64VarP 

``` go
func (f *FlagSet) Int64VarP(p *int64, name, shorthand string, value int64, usage string)
```

​	Int64VarP 类似于 Int64Var，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int8 

``` go
func (f *FlagSet) Int8(name string, value int8, usage string) *int8
```

​	Int8 定义了一个带有指定名称、默认值和用法说明的 int8 标志。返回值是一个 int8 变量的地址，用于存储标志的值。

##### (*FlagSet) Int8P 

``` go
func (f *FlagSet) Int8P(name, shorthand string, value int8, usage string) *int8
```

​	Int8P 类似于 Int8，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Int8Var 

``` go
func (f *FlagSet) Int8Var(p *int8, name string, value int8, usage string)
```

​	Int8Var 定义了一个带有指定名称、默认值和用法说明的 int8 标志。参数 p 指向一个 int8 变量，用于存储标志的值。

##### (*FlagSet) Int8VarP 

``` go
func (f *FlagSet) Int8VarP(p *int8, name, shorthand string, value int8, usage string)
```

​	Int8VarP 类似于 Int8Var，但允许在单个短划线后使用速记字母。

##### (*FlagSet) IntP 

``` go
func (f *FlagSet) IntP(name, shorthand string, value int, usage string) *int
```

​	IntP 类似于 Int，但允许在单个短划线后使用速记字母。

##### (*FlagSet) IntSlice 

``` go
func (f *FlagSet) IntSlice(name string, value []int, usage string) *[]int
```

​	IntSlice 定义了一个带有指定名称、默认值和用法说明的 []int 标志。返回值是一个 []int 变量的地址，用于存储标志的值。

##### (*FlagSet) IntSliceP 

``` go
func (f *FlagSet) IntSliceP(name, shorthand string, value []int, usage string) *[]int
```

​	IntSliceP 类似于 IntSlice，但允许在单个短划线后使用速记字母。

##### (*FlagSet) IntSliceVar 

``` go
func (f *FlagSet) IntSliceVar(p *[]int, name string, value []int, usage string)
```

​	IntSliceVar 定义了一个带有指定名称、默认值和用法说明的 intSlice 标志。参数 p 指向一个 []int 变量，用于存储标志的值。

##### (*FlagSet) IntSliceVarP 

``` go
func (f *FlagSet) IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string)
```

​	IntSliceVarP 类似于 IntSliceVar，但允许在单个短划线后使用速记字母。

##### (*FlagSet) IntVar 

``` go
func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
```

​	IntVar 定义了一个带有指定名称、默认值和用法说明的 int 标志。参数 p 指向一个 int 变量，用于存储标志的值。

##### (*FlagSet) IntVarP 

``` go
func (f *FlagSet) IntVarP(p *int, name, shorthand string, value int, usage string)
```

​	IntVarP 类似于 IntVar，但允许在单个短划线后使用速记字母。

##### (*FlagSet) Lookup 

``` go
func (f *FlagSet) Lookup(name string) *Flag
```

​	Lookup 根据名称返回标志的 Flag 结构，如果不存在则返回 nil。

##### (*FlagSet) MarkDeprecated 

``` go
func (f *FlagSet) MarkDeprecated(name string, usageMessage string) error
```

​	MarkDeprecated 表示程序中的一个标志已被弃用。它将继续工作，但不会显示在帮助或用法消息中。使用此标志还将打印给定的 usageMessage。

##### (*FlagSet) MarkHidden 

``` go
func (f *FlagSet) MarkHidden(name string) error
```

​	MarkHidden 将标志标记为“隐藏”状态。它将继续工作，但不会显示在帮助或用法消息中。

##### (*FlagSet) MarkShorthandDeprecated 

``` go
func (f *FlagSet) MarkShorthandDeprecated(name string, usageMessage string) error
```

​	MarkShorthandDeprecated 将标志的缩写标记为在程序中已弃用。它将继续工作，但不会显示在帮助或用法消息中。使用此标志还将打印给定的 usageMessage。

##### (*FlagSet) NArg 

``` go
func (f *FlagSet) NArg() int
```

​	NArg 返回处理完标志后剩余的参数数量。

##### (*FlagSet) NFlag 

``` go
func (f *FlagSet) NFlag() int
```

​	NFlag 返回已设置的标志数量。

##### (*FlagSet) Parse 

``` go
func (f *FlagSet) Parse(arguments []string) error
```

​	Parse 从参数列表解析标志定义，参数列表不应包括命令名称。必须在 FlagSet 中的所有标志定义之后、程序访问标志之前调用。如果设置了 -help 但未定义，返回值将是 ErrHelp。

##### (*FlagSet) ParseAll 

``` go
func (f *FlagSet) ParseAll(arguments []string, fn func(flag *Flag, value string) error) error
```

​	ParseAll 从参数列表解析标志定义，参数列表不应包括命令名称。fn 的参数是 flag 和 value。必须在 FlagSet 中的所有标志定义之后、程序访问标志之前调用。如果设置了 -help 但未定义，返回值将是 ErrHelp。

##### (*FlagSet) Parsed 

``` go
func (f *FlagSet) Parsed() bool
```

​	Parsed 报告是否已调用 f.Parse。

##### (*FlagSet) PrintDefaults 

``` go
func (f *FlagSet) PrintDefaults()
```

​	PrintDefaults 打印所有已定义标志的默认值到标准错误（除非另有配置）。

##### (*FlagSet) Set 

``` go
func (f *FlagSet) Set(name, value string) error
```

​	Set 设置指定标志的值。

##### (*FlagSet) SetAnnotation 

``` go
func (f *FlagSet) SetAnnotation(name, key string, values []string) error
```

​	SetAnnotation 允许在 FlagSet 上设置任意注释。这有时由 spf13/cobra 程序使用，它们希望生成额外的 bash 完成信息。

##### (*FlagSet) SetInterspersed 

``` go
func (f *FlagSet) SetInterspersed(interspersed bool)
```

​	SetInterspersed 设置是否支持交错的选项/非选项参数。

##### (*FlagSet) SetNormalizeFunc 

``` go
func (f *FlagSet) SetNormalizeFunc(n func(f *FlagSet, name string) NormalizedName)
```

​	SetNormalizeFunc 允许您添加一个函数，用于转换标志名称。添加到 FlagSet 中的标志将被翻译，当任何东西尝试查找该标志时，也会被翻译。因此，可以创建名为 "getURL" 的标志，并将其翻译为 "geturl"。然后用户可以传递 "`--getUrl`"，这也可能被翻译为 "geturl"，一切都能正常工作。

##### (*FlagSet) SetOutput 

``` go
func (f *FlagSet) SetOutput(output io.Writer)
```

​	SetOutput 设置用于帮助和错误消息的目标。如果 output 为 nil，则使用 os.Stderr。

##### (*FlagSet) ShorthandLookup 

``` go
func (f *FlagSet) ShorthandLookup(name string) *Flag
```

​	ShorthandLookup 返回缩写标志的 Flag 结构，如果不存在则返回 nil。如果 name 的长度大于 1，则会引发 panic。

###### ShorthandLookup  Example

```go
package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	name := "verbose"
	short := name[:1]

	fs := pflag.NewFlagSet("Example", pflag.ContinueOnError)
	fs.BoolP(name, short, false, "verbose output")

	// len(short) must be == 1
	flag := fs.ShorthandLookup(short)

	fmt.Println(flag.Name)
}

```



##### (*FlagSet) String 

``` go
func (f *FlagSet) String(name string, value string, usage string) *string
```

​	String 定义了一个带有指定名称、默认值和用法说明的字符串标志。返回值是一个字符串变量的地址，用于存储标志的值。

##### (*FlagSet) StringArray 

``` go
func (f *FlagSet) StringArray(name string, value []string, usage string) *[]string
```

​	StringArray 定义了一个带有指定名称、默认值和用法说明的字符串标志。返回值是一个 []string 变量的地址，用于存储标志的值。每个参数的值不会尝试用逗号分隔。对于这种情况，请使用 StringSlice。

##### (*FlagSet) StringArrayP 

``` go
func (f *FlagSet) StringArrayP(name, shorthand string, value []string, usage string) *[]string
```

​	StringArrayP 类似于 StringArray，但允许在单个短划线后使用速记字母。

##### (*FlagSet) StringArrayVar 

``` go
func (f *FlagSet) StringArrayVar(p *[]string, name string, value []string, usage string)
```

​	StringArrayVar 定义了一个带有指定名称、默认值和用法说明的字符串标志。参数 p 指向一个 []string 变量，用于存储多个标志的值。每个参数的值不会尝试用逗号分隔。对于这种情况，请使用 StringSlice。

##### (*FlagSet) StringArrayVarP 

``` go
func (f *FlagSet) StringArrayVarP(p *[]string, name, shorthand string, value []string, usage string)
```

​	StringArrayVarP 类似于 StringArrayVar，但允许在单个短划线后使用速记字母。

##### (*FlagSet) StringP 

``` go
func (f *FlagSet) StringP(name, shorthand string, value string, usage string) *string
```

​	StringP 类似于 String，但允许在单个短划线后使用速记字母。

##### (*FlagSet) StringSlice 

``` go
func (f *FlagSet) StringSlice(name string, value []string, usage string) *[]string
```

​	StringSlice 定义了一个带有指定名称、默认值和用法说明的字符串标志。返回值是一个 []string 变量的地址，用于存储标志的值。与 StringArray 标志相比，StringSlice 标志将以逗号分隔的值作为参数，并相应地进行拆分。例如：

```
--ss="v1,v2" --ss="v3"
```

将导致：

```
[]string{"v1", "v2", "v3"}
```

##### (*FlagSet) StringSliceP 

``` go
func (f *FlagSet) StringSliceP(name, shorthand string, value []string, usage string) *[]string
```

​	StringSliceP 类似于 StringSlice，但允许在单个短划线后使用速记字母。

##### (*FlagSet) StringSliceVar 

``` go
func (f *FlagSet) StringSliceVar(p *[]string, name string, value []string, usage string)
```

​	StringSliceVar 定义了一个带有指定名称、默认值和用法说明的字符串标志。参数 p 指向一个 []string 变量，用于存储标志的值。与 StringArray 标志相比，StringSlice 标志将以逗号分隔的值作为参数，并相应地进行拆分。例如：

```
--ss="v1,v2" --ss="v3"
```

将导致：

```
[]string{"v1", "v2", "v3"}
```

##### (*FlagSet) StringSliceVarP 

``` go
func (f *FlagSet) StringSliceVarP(p *[]string, name, shorthand string, value []string, usage string)
```

​	StringSliceVarP 类似于 StringSliceVar，但允许在单个短划线后使用速记字母。

##### (*FlagSet) StringToInt  <- v1.0.3

``` go
func (f *FlagSet) StringToInt(name string, value map[string]int, usage string) *map[string]int
```

​	StringToInt 定义了一个带有指定名称、默认值和用法说明的字符串标志。返回值是一个 map[string]int 变量的地址，用于存储标志的值。每个参数的值不会尝试用逗号分隔。

##### (*FlagSet) StringToInt64  <- v1.0.5

``` go
func (f *FlagSet) StringToInt64(name string, value map[string]int64, usage string) *map[string]int64
```

​	StringToInt64 定义了一个带有指定名称、默认值和用法说明的字符串标志。返回值是一个 map[string]int64 变量的地址，用于存储标志的值。每个参数的值不会尝试用逗号分隔。

##### (*FlagSet) StringToInt64P  <- v1.0.5

``` go
func (f *FlagSet) StringToInt64P(name, shorthand string, value map[string]int64, usage string) *map[string]int64
```

​	StringToInt64P 类似于 StringToInt64，但允许在单个短划线后使用速记字母。

##### (*FlagSet) StringToInt64Var  <- v1.0.5

``` go
func (f *FlagSet) StringToInt64Var(p *map[string]int64, name string, value map[string]int64, usage string)
```

​	StringToInt64Var 定义了一个具有指定名称、默认值和用法字符串的字符串标志。参数 p 指向一个 map[string]int64 变量，用于存储多个标志的值。每个参数的值不会尝试用逗号分隔。

##### (*FlagSet) StringToInt64VarP  <- v1.0.5

``` go
func (f *FlagSet) StringToInt64VarP(p *map[string]int64, name, shorthand string, value map[string]int64, usage string)
```

​	StringToInt64VarP 类似于 StringToInt64Var，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) StringToIntP  <- v1.0.3

``` go
func (f *FlagSet) StringToIntP(name, shorthand string, value map[string]int, usage string) *map[string]int
```

​	StringToIntP 类似于 StringToInt，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) StringToIntVar  <- v1.0.3

``` go
func (f *FlagSet) StringToIntVar(p *map[string]int, name string, value map[string]int, usage string)
```

​	StringToIntVar 定义了一个具有指定名称、默认值和用法字符串的字符串标志。参数 p 指向一个 map[string]int 变量，用于存储多个标志的值。每个参数的值不会尝试用逗号分隔。

##### (*FlagSet) StringToIntVarP  <- v1.0.3

``` go
func (f *FlagSet) StringToIntVarP(p *map[string]int, name, shorthand string, value map[string]int, usage string)
```

​	StringToIntVarP 类似于 StringToIntVar，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) StringToString  <- v1.0.3

``` go
func (f *FlagSet) StringToString(name string, value map[string]string, usage string) *map[string]string
```

​	StringToString 定义了一个具有指定名称、默认值和用法字符串的字符串标志。返回值是一个 map[string]string 变量的地址，该变量用于存储标志的值。每个参数的值不会尝试用逗号分隔。

##### (*FlagSet) StringToStringP  <- v1.0.3

``` go
func (f *FlagSet) StringToStringP(name, shorthand string, value map[string]string, usage string) *map[string]string
```

​	StringToStringP 类似于 StringToString，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) StringToStringVar  <- v1.0.3

``` go
func (f *FlagSet) StringToStringVar(p *map[string]string, name string, value map[string]string, usage string)
```

​	StringToStringVar 定义了一个具有指定名称、默认值和用法字符串的字符串标志。参数 p 指向一个 map[string]string 变量，用于存储多个标志的值。每个参数的值不会尝试用逗号分隔。

##### (*FlagSet) StringToStringVarP  <- v1.0.3

``` go
func (f *FlagSet) StringToStringVarP(p *map[string]string, name, shorthand string, value map[string]string, usage string)
```

​	StringToStringVarP 类似于 StringToStringVar，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) StringVar 

``` go
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
```

​	StringVar 定义了一个具有指定名称、默认值和用法字符串的字符串标志。参数 p 指向一个字符串变量，用于存储标志的值。

##### (*FlagSet) StringVarP 

``` go
func (f *FlagSet) StringVarP(p *string, name, shorthand string, value string, usage string)
```

​	StringVarP 类似于 StringVar，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Uint 

``` go
func (f *FlagSet) Uint(name string, value uint, usage string) *uint
```

​	Uint 定义了一个具有指定名称、默认值和用法字符串的 uint 标志。返回值是一个指向 uint 变量的地址，用于存储标志的值。

##### (*FlagSet) Uint16 

``` go
func (f *FlagSet) Uint16(name string, value uint16, usage string) *uint16
```

​	Uint16 定义了一个具有指定名称、默认值和用法字符串的 uint16 标志。返回值是一个指向 uint16 变量的地址，用于存储标志的值。

##### (*FlagSet) Uint16P 

``` go
func (f *FlagSet) Uint16P(name, shorthand string, value uint16, usage string) *uint16
```

​	Uint16P 类似于 Uint16，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Uint16Var 

``` go
func (f *FlagSet) Uint16Var(p *uint16, name string, value uint16, usage string)
```

​	Uint16Var 定义了一个具有指定名称、默认值和用法字符串的 uint16 标志。参数 p 指向一个 uint16 变量，用于存储标志的值。

##### (*FlagSet) Uint16VarP 

``` go
func (f *FlagSet) Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string)
```

​	Uint16VarP 类似于 Uint16Var，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Uint32 

``` go
func (f *FlagSet) Uint32(name string, value uint32, usage string) *uint32
```

​	Uint32 定义了一个具有指定名称、默认值和用法字符串的 uint32 标志。返回值是一个指向 uint32 变量的地址，用于存储标志的值。

##### (*FlagSet) Uint32P 

``` go
func (f *FlagSet) Uint32P(name, shorthand string, value uint32, usage string) *uint32
```

​	Uint32P 类似于 Uint32，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Uint32Var 

``` go
func (f *FlagSet) Uint32Var(p *uint32, name string, value uint32, usage string)
```

​	Uint32Var 定义了一个具有指定名称、默认值和用法字符串的 uint32 标志。参数 p 指向一个 uint32 变量，用于存储标志的值。

##### (*FlagSet) Uint32VarP 

``` go
func (f *FlagSet) Uint32VarP(p *uint32, name, shorthand string, value uint32, usage string)
```

​	Uint32VarP 类似于 Uint32Var，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Uint64 

``` go
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
```

​	Uint64 方法定义了一个具有指定名称、默认值和用法字符串的 uint64 标志。返回值是一个指向 uint64 变量的地址，用于存储标志的值。

##### (*FlagSet) Uint64P 

``` go
func (f *FlagSet) Uint64P(name, shorthand string, value uint64, usage string) *uint64
```

​	Uint64P 方法类似于 Uint64方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Uint64Var 

``` go
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
```

​	Uint64Var 方法定义了一个具有指定名称、默认值和用法字符串的 uint64 标志。参数 p 指向一个 uint64 变量，用于存储标志的值。

##### (*FlagSet) Uint64VarP 

``` go
func (f *FlagSet) Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string)
```

​	Uint64VarP 方法类似于 Uint64Var方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Uint8 

``` go
func (f *FlagSet) Uint8(name string, value uint8, usage string) *uint8
```

​	Uint8 方法定义了一个具有指定名称、默认值和用法字符串的 uint8 标志。返回值是一个指向 uint8 变量的地址，用于存储标志的值。

##### (*FlagSet) Uint8P 

``` go
func (f *FlagSet) Uint8P(name, shorthand string, value uint8, usage string) *uint8
```

​	Uint8P 方法类似于 Uint8方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Uint8Var 

``` go
func (f *FlagSet) Uint8Var(p *uint8, name string, value uint8, usage string)
```

​	Uint8Var 方法定义了一个具有指定名称、默认值和用法字符串的 uint8 标志。参数 p 指向一个 uint8 变量，用于存储标志的值。

##### (*FlagSet) Uint8VarP 

``` go
func (f *FlagSet) Uint8VarP(p *uint8, name, shorthand string, value uint8, usage string)
```

​	Uint8VarP 方法类似于 Uint8Var方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) UintP 

``` go
func (f *FlagSet) UintP(name, shorthand string, value uint, usage string) *uint
```

​	UintP 方法类似于 Uint方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) UintSlice 

``` go
func (f *FlagSet) UintSlice(name string, value []uint, usage string) *[]uint
```

​	UintSlice 方法定义了一个具有指定名称、默认值和用法字符串的 []uint 标志。返回值是一个指向 []uint 变量的地址，用于存储标志的值。

##### (*FlagSet) UintSliceP 

``` go
func (f *FlagSet) UintSliceP(name, shorthand string, value []uint, usage string) *[]uint
```

​	UintSliceP 方法类似于 UintSlice方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) UintSliceVar 

``` go
func (f *FlagSet) UintSliceVar(p *[]uint, name string, value []uint, usage string)
```

​	UintSliceVar 方法定义了一个具有指定名称、默认值和用法字符串的 []uint 标志。参数 p 指向一个 []uint 变量，用于存储标志的值。

##### (*FlagSet) UintSliceVarP 

``` go
func (f *FlagSet) UintSliceVarP(p *[]uint, name, shorthand string, value []uint, usage string)
```

​	UintSliceVarP 方法类似于 UintSliceVar方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) UintVar 

``` go
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
```

​	UintVar 方法定义了一个具有指定名称、默认值和用法字符串的 uint 标志。参数 p 指向一个 uint 变量，用于存储标志的值。

##### (*FlagSet) UintVarP 

``` go
func (f *FlagSet) UintVarP(p *uint, name, shorthand string, value uint, usage string)
```

​	UintVarP 方法类似于 UintVar方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) Var 

``` go
func (f *FlagSet) Var(value Value, name string, usage string)
```

​	Var 方法定义了一个具有指定名称和用法字符串的标志。标志的类型和值由第一个参数表示，类型为 Value，通常包含用户定义的 Value 实现。例如，调用者可以创建一个标志，通过为切片分配方法，将逗号分隔的字符串转换为字符串切片；特别是，Set 方法会将逗号分隔的字符串分解为切片。

##### (*FlagSet) VarP 

``` go
func (f *FlagSet) VarP(value Value, name, shorthand, usage string)
```

​	VarP 方法类似于 Var方法，但接受一个可以在单个短划线后使用的速记字母。

##### (*FlagSet) VarPF 

``` go
func (f *FlagSet) VarPF(value Value, name, shorthand, usage string) *Flag
```

​	VarPF 方法类似于 VarP方法，但返回创建的标志。

##### (*FlagSet) Visit 

``` go
func (f *FlagSet) Visit(fn func(*Flag))
```

​	Visit 方法以词典顺序或原始顺序（如果 f.SortFlags 为 false）访问标志，对每个标志调用 fn。它仅访问已设置的标志。

##### (*FlagSet) VisitAll 

``` go
func (f *FlagSet) VisitAll(fn func(*Flag))
```

​	VisitAll 方法以词典顺序或原始顺序（如果 f.SortFlags 为 false）访问标志，对每个标志调用 fn。它访问所有标志，包括未设置的标志。

#### type NormalizedName 

``` go
type NormalizedName string
```

​	NormalizedName 是根据 FlagSet 规则进行标准化的标志名称（例如，使 '`-`' 和 '`_`' 等效）。

#### type ParseErrorsWhitelist  <- v1.0.1

``` go
type ParseErrorsWhitelist struct {
    // UnknownFlags 将忽略未知标志的错误并继续解析其余标志
	UnknownFlags bool
}
```

​	ParseErrorsWhitelist 结构体定义了可以忽略的解析错误。

#### type SliceValue  <- v1.0.5

``` go
type SliceValue interface {
    // Append 将指定的值添加到标志值列表的末尾。
	Append(string) error
    // Replace 将完全覆盖当前在标志值列表中的任何数据。
	Replace([]string) error
    // GetSlice 返回标志值列表作为字符串数组。
	GetSlice() []string
}
```

​	SliceValue 是所有持有值列表的标志的辅助接口。这允许对列表标志的值进行完全控制，并避免了复杂的 csv 编组和解组。

#### type Value 

``` go
type Value interface {
	String() string
	Set(string) error
	Type() string
}
```

​	Value 是存储在标志中的动态值的接口（默认值表示为字符串）。