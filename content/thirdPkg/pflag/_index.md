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
> github：[https://github.com/spf13/pflag](https://github.com/spf13/pflag)

## 描述

​	pflag是Go语言的`flag`包的一个替代品，实现了POSIX/GNU风格的`--flags`。

​	pflag兼容[GNU对命令行选项的POSIX建议的扩展](http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html)。更详细的描述请参见下面的"命令行标志语法"部分。

​	pflag采用与Go语言相同的BSD许可证授权，许可证的内容可以在LICENSE文件中找到。

## 安装

​	使用标准的`go get`命令可以获取pflag。

​	运行以下命令进行安装：

```
go get github.com/spf13/pflag
```

​	运行以下命令进行测试：

```
go test github.com/spf13/pflag
```

## 用法

​	pflag是Go原生`flag`包的一个替代品。如果您将pflag导入为"flag"，则所有的代码都应该继续正常工作，无需进行任何更改。

```
import flag "github.com/spf13/pflag"
```

​	有一个例外情况：如果您直接实例化Flag结构体，还需要设置一个额外的字段"Shorthand"。大多数代码不会直接实例化这个结构体，而是使用诸如String()、BoolVar()和Var()等函数，因此不受影响。

​	使用flag.String()、Bool()、Int()等函数来定义标志。

​	以下是一个声明整数标志的示例，标志名为`-flagname`，存储在指针`ip`中，类型为`*int`。

```
var ip *int = flag.Int("flagname", 1234, "help message for flagname")
```

​	如果需要，您可以使用Var()函数将标志绑定到一个变量上。

```
var flagvar int
func init() {
    flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

​	或者您可以创建满足Value接口的自定义标志（使用指针接收器），并通过以下方式将它们与标志解析相关联：

```
flag.Var(&flagVal, "name", "help message for flagname")
```

​	对于这样的标志，其默认值就是变量的初始值。

​	在定义所有标志后，调用

```
flag.Parse()
```

来将命令行解析进已定义的标志。

​	然后可以直接使用标志。**如果使用的是标志本身，则它们都是指针；如果绑定到变量，则它们是值**。

```
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

​	如果您使用 `FlagSet`，并且发现在代码中跟踪所有指针变得困难，那么可以使用一些辅助函数来获取 `Flag`（结构体） 中存储的值如果您有一个名为'flagname'、类型为`int`的`pflag.FlagSet`，您可以使用`GetInt()`来获取int值。但请注意，'flagname'必须存在且为int类型，否则`GetString("flagname")`将失败。

```
i, err := flagset.GetInt("flagname")
```

​	在解析后，该标志之后的实参可以作为`flag.Args()`的切片或作为`flag.Arg(i)`单独使用。实参的索引范围是从0到`flag.NArg()-1`。

​	pflag包还定义了一些在 `flag` 包中不存在的新函数，它们为标志提供了一字母缩写。您可以通过在定义标志的任何函数名称后附加 '`P`' 来使用这些函数。

```
var ip = flag.IntP("flagname", "f", 1234, "help message")
var flagvar bool
func init() {
	flag.BoolVarP(&flagvar, "boolname", "b", true, "help message")
}
flag.VarP(&flagVal, "varname", "v", "help message")
```

​	缩写字母可以在命令行上使用单破折号。布尔型缩写标志可以与其他缩写标志结合使用。

​	默认的命令行标志集由顶层函数控制。FlagSet类型允许定义独立的标志集，例如在命令行接口中实现子命令。FlagSet的方法类似于顶层函数用于命令行标志集的方法。

## Setting no option default values for flags 为标志设置无选项默认值

After you create a flag it is possible to set the pflag.NoOptDefVal for the given flag. Doing this changes the meaning of the flag slightly. If a flag has a NoOptDefVal and the flag is set on the command line without an option the flag will be set to the NoOptDefVal. For example given:

在创建标志后，可以为给定的标志设置pflag.NoOptDefVal。这样做会略微改变标志的含义。如果一个标志具有NoOptDefVal，并且在命令行上设置该标志而没有选项，那么该标志将被设置为NoOptDefVal。例如：

```
var ip = flag.IntP("flagname", "f", 1234, "help message")
flag.Lookup("flagname").NoOptDefVal = "4321"
```

Would result in something like

会产生类似以下的结果

| Parsed Arguments | Resulting Value |
| ---------------- | --------------- |
| --flagname=1357  | ip=1357         |
| --flagname       | ip=4321         |
| [nothing]        | ip=1234         |

## Command line flag syntax 命令行标志语法

```
--flag    // boolean flags, or flags with no option default values
--flag x  // only on flags without a default value
--flag=x
```

Unlike the flag package, a single dash before an option means something different than a double dash. Single dashes signify a series of shorthand letters for flags. All but the last shorthand letter must be boolean flags or a flag with a default value

与flag包不同，选项之前的单破折号和双破折号有不同的含义。单破折号表示一系列标志的缩写字母。除了最后一个缩写字母可以是布尔型标志或具有默认值的标志外，其他都必须是布尔型标志。

```
// boolean or flags where the 'no option default value' is set
-f
-f=true
-abc
but
-b true is INVALID

// non-boolean and flags without a 'no option default value'
-n 1234
-n=1234
-n1234

// mixed
-abcs "hello"
-absd="hello"
-abcs1234
```

Flag parsing stops after the terminator "--". Unlike the flag package, flags can be interspersed with arguments anywhere on the command line before this terminator.

在" -- "之后，标志解析将停止。与flag包不同，在这个终止符之前，标志可以与参数混合在命令行的任何位置。

Integer flags accept 1234, 0664, 0x1234 and may be negative. Boolean flags (in their long form) accept 1, 0, t, f, true, false, TRUE, FALSE, True, False. Duration flags accept any input valid for time.ParseDuration.

整数型标志接受1234、0664、0x1234等，并且可以为负数。布尔型标志（长格式）接受1、0、t、f、true、false、TRUE、FALSE、True、False。持续时间标志接受任何对于time.ParseDuration有效的输入。

## Mutating or "Normalizing" Flag names 修改或"规范化"标志名称

It is possible to set a custom flag name 'normalization function.' It allows flag names to be mutated both when created in the code and when used on the command line to some 'normalized' form. The 'normalized' form is used for comparison. Two examples of using the custom normalization func follow.

可以设置自定义的标志名称"规范化函数"。它允许标志名称在代码中创建时和在命令行上使用时以某种"规范化"的形式进行变换。比较时使用"规范化"的形式。下面是两个使用自定义规范化函数的示例。

**Example #1**: You want -, _, and . in flags to compare the same. aka --my-flag == --my_flag == --my.flag

**示例＃1**：您希望在标志中比较 -、_ 和 . 时得到相同的结果。也就是说 --my-flag == --my_flag == --my.flag

```
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

**Example #2**: You want to alias two flags. aka --old-flag-name == --new-flag-name

**示例＃2**：您希望给两个标志设置别名。也就是说 --old-flag-name == --new-flag-name

```
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

## Deprecating a flag or its shorthand 弃用标志或其缩写

It is possible to deprecate a flag, or just its shorthand. Deprecating a flag/shorthand hides it from help text and prints a usage message when the deprecated flag/shorthand is used.

可以弃用一个标志，或仅弃用它的缩写。弃用标志/缩写将在帮助文本中隐藏，并在使用被弃用的标志/缩写时显示使用消息。

**Example #1**: You want to deprecate a flag named "badflag" as well as inform the users what flag they should use instead.

**示例＃1**：您希望弃用一个名为"badflag"的标志，并告知用户应该使用哪个标志代替它。

```
// deprecate a flag by specifying its name and a usage message
flags.MarkDeprecated("badflag", "please use --good-flag instead")
```

This hides "badflag" from help text, and prints `Flag --badflag has been deprecated, please use --good-flag instead` when "badflag" is used.

这将在帮助文本中隐藏"badflag"，并在使用"badflag"时打印`Flag --badflag已被弃用，请使用--good-flag代替`。

**Example #2**: You want to keep a flag name "noshorthandflag" but deprecate its shortname "n".

**示例＃2**：您希望保留一个名为"noshorthandflag"的标志，但弃用其缩写"n"。

```
// deprecate a flag shorthand by specifying its flag name and a usage message
flags.MarkShorthandDeprecated("noshorthandflag", "please use --noshorthandflag only")
```

This hides the shortname "n" from help text, and prints `Flag shorthand -n has been deprecated, please use --noshorthandflag only` when the shorthand "n" is used.

这将在帮助文本中隐藏缩写"n"，并在使用缩写"n"时打印`Flag shorthand -n已被弃用，请只使用--noshorthandflag`。

Note that usage message is essential here, and it should not be empty.

请注意，这里的用法消息是必要的，不应为空。

## Hidden flags 隐藏标志

It is possible to mark a flag as hidden, meaning it will still function as normal, however will not show up in usage/help text.

可以将一个标志标记为隐藏，这意味着它仍然会正常工作，但不会显示在使用/帮助文本中。

**Example**: You have a flag named "secretFlag" that you need for internal use only and don't want it showing up in help text, or for its usage text to be available.

**示例**：您有一个名为"secretFlag"的标志，仅供内部使用，不希望它显示在帮助文本中或可用于使用文本。

```
// hide a flag by specifying its name
flags.MarkHidden("secretFlag")
```

## Disable sorting of flags 禁用标志的排序

`pflag` allows you to disable sorting of flags for help and usage message.

`pflag`允许您禁用帮助和使用消息的标志排序。

****示例****:

```
flags.BoolP("verbose", "v", false, "verbose output")
flags.String("coolflag", "yeaah", "it's really cool flag")
flags.Int("usefulflag", 777, "sometimes it's very useful")
flags.SortFlags = false
flags.PrintDefaults()
```

**输出**:

```
  -v, --verbose           verbose output
      --coolflag string   it's really cool flag (default "yeaah")
      --usefulflag int    sometimes it's very useful (default 777)
```

## Supporting Go flags when using pflag 在使用pflag时支持Go标志

In order to support flags defined using Go's `flag` package, they must be added to the `pflag` flagset. This is usually necessary to support flags defined by third-party dependencies (e.g. `golang/glog`).

为了支持使用Go的`flag`包定义的标志，它们必须添加到`pflag`的标志集中。这通常是为了支持由第三方依赖（例如`golang/glog`）定义的标志。

**Example**: You want to add the Go flags to the `CommandLine` flagset

**示例**：您希望将Go标志添加到`CommandLine`的标志集中

```
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

You can see the full reference documentation of the pflag package [at godoc.org](http://godoc.org/github.com/spf13/pflag), or through go's standard documentation system by running `godoc -http=:6060` and browsing to http://localhost:6060/pkg/github.com/spf13/pflag after installation.

您可以在[godoc.org](http://godoc.org/github.com/spf13/pflag)上查看pflag包的完整参考文档，或者在安装后通过运行`godoc -http=:6060`并浏览http://localhost:6060/pkg/github.com/spf13/pflag来使用Go的标准文档系统。



## 文档概述

Package pflag is a drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.

pflag包是Go标准库flag包的一个替代品，实现了POSIX/GNU风格的--flags。



pflag is compatible with the GNU extensions to the POSIX recommendations for command-line options. See http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html

pflag与GNU对POSIX命令行选项的推荐扩展兼容。请参阅http://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html

用法：

pflag is a drop-in replacement of Go's native flag package. If you import pflag under the name "flag" then all code should continue to function with no changes.

pflag是Go原生flag包的一个替代品。如果您使用名称"flag"导入pflag，则所有代码应该继续正常工作，无需更改。

```
import flag "github.com/spf13/pflag"
```

There is one exception to this: if you directly instantiate the Flag struct there is one more field "Shorthand" that you will need to set. Most code never instantiates this struct directly, and instead uses functions such as String(), BoolVar(), and Var(), and is therefore unaffected.

但有一个例外：如果直接实例化Flag结构体，则需要设置一个额外的字段"Shorthand"。大多数代码不会直接实例化此结构体，而是使用诸如String()、BoolVar()和Var()等函数，因此不受影响。

Define flags using flag.String(), Bool(), Int(), etc.

使用flag.String()、Bool()、Int()等来定义标志。

This declares an integer flag, -flagname, stored in the pointer ip, with type *int.

这将声明一个整数标志-flagname，存储在指针ip中，类型为*int。

```
var ip = flag.Int("flagname", 1234, "help message for flagname")
```

If you like, you can bind the flag to a variable using the Var() functions.

如果愿意，可以使用Var()函数将标志绑定到变量。

```
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

Or you can create custom flags that satisfy the Value interface (with pointer receivers) and couple them to flag parsing by

或者，您可以创建满足Value接口（具有指针接收者）的自定义标志，并将它们与标志解析耦合

```
flag.Var(&flagVal, "name", "help message for flagname")
```

For such flags, the default value is just the initial value of the variable.

对于这些标志，默认值只是变量的初始值。

After all flags are defined, call

在定义所有标志后，调用

```
flag.Parse()
```

to parse the command line into the defined flags.

将命令行解析为已定义的标志。

Flags may then be used directly. If you're using the flags themselves, they are all pointers; if you bind to variables, they're values.

然后可以直接使用标志。如果使用标志本身，它们都是指针；如果绑定到变量，则是值。

```
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

After parsing, the arguments after the flag are available as the slice flag.Args() or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1.

解析后，标志后面的参数可以作为切片flag.Args()或单独作为flag.Arg(i)使用。参数从0到flag.NArg()-1进行索引。

The pflag package also defines some new functions that are not in flag, that give one-letter shorthands for flags. You can use these by appending 'P' to the name of any function that defines a flag.

pflag包还定义了一些在flag中不存在的新函数，为标志提供了一字母的缩写形式。您可以通过将定义标志的任何函数的名称附加'P'来使用这些函数。

```
var ip = flag.IntP("flagname", "f", 1234, "help message")
var flagvar bool
func init() {
	flag.BoolVarP(&flagvar, "boolname", "b", true, "help message")
}
flag.VarP(&flagval, "varname", "v", "help message")
```

Shorthand letters can be used with single dashes on the command line. Boolean shorthand flags can be combined with other shorthand flags.

单个短横线加上字母可以在命令行中使用。布尔型短标志可以与其他短标志组合使用。

Command line flag syntax:

命令行标志语法：

```
--flag    // boolean flags only
--flag=x
```

Unlike the flag package, a single dash before an option means something different than a double dash. Single dashes signify a series of shorthand letters for flags. All but the last shorthand letter must be boolean flags.

与flag包不同，选项之前的单个短横线与双短横线有不同的含义。单个短横线表示一系列用于标志的简写字母。除了最后一个简写字母，其他字母必须是布尔型标志。

```
// boolean flags
-f
-abc
// non-boolean flags
-n 1234
-Ifile
// mixed
-abcs "hello"
-abcn1234
```

Flag parsing stops after the terminator "--". Unlike the flag package, flags can be interspersed with arguments anywhere on the command line before this terminator.

在终止符"--"之后，标志解析停止。与flag包不同，标志可以与参数交错出现在终止符之前的命令行中的任何位置。

Integer flags accept 1234, 0664, 0x1234 and may be negative. Boolean flags (in their long form) accept 1, 0, t, f, true, false, TRUE, FALSE, True, False. Duration flags accept any input valid for time.ParseDuration.

整数标志接受1234、0664、0x1234等值，可以为负数。布尔型标志（以其长形式表示）接受1、0、t、f、true、false、TRUE、FALSE、True、False。持续时间标志接受任何time.ParseDuration可接受的输入。

The default set of command-line flags is controlled by top-level functions. The FlagSet type allows one to define independent sets of flags, such as to implement subcommands in a command-line interface. The methods of FlagSet are analogous to the top-level functions for the command-line flag set.

默认的命令行标志集由顶级函数控制。FlagSet类型允许定义独立的标志集，例如在命令行界面中实现子命令。FlagSet的方法与用于命令行标志集的顶级函数类似。



#### Examples 

- [FlagSet.ShorthandLookup](https://pkg.go.dev/github.com/spf13/pflag#example-FlagSet.ShorthandLookup)
- [ShorthandLookup](https://pkg.go.dev/github.com/spf13/pflag#example-ShorthandLookup)

### Constants 

This section is empty.

### Variables 

[View Source](https://github.com/spf13/pflag/blob/v1.0.5/flag.go#L1212)

```
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

CommandLine is the default set of command-line flags, parsed from os.Args.

[View Source](https://github.com/spf13/pflag/blob/v1.0.5/flag.go#L113)

```
var ErrHelp = errors.New("pflag: help requested")
```

ErrHelp is the error returned if the flag -help is invoked but no such flag is defined.

[View Source](https://github.com/spf13/pflag/blob/v1.0.5/flag.go#L773)

```
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	PrintDefaults()
}
```

Usage prints to standard error a usage message documenting all defined command-line flags. The function is a variable that may be changed to point to a custom function. By default it prints a simple header and calls PrintDefaults; for details about the format of the output and how to control it, see the documentation for PrintDefaults.

### Functions 

#### func Arg 

```
func Arg(i int) string
```

Arg returns the i'th command-line argument. Arg(0) is the first remaining argument after flags have been processed.

#### func Args 

```
func Args() []string
```

Args returns the non-flag command-line arguments.

#### func Bool 

```
func Bool(name string, value bool, usage string) *bool
```

Bool defines a bool flag with specified name, default value, and usage string. The return value is the address of a bool variable that stores the value of the flag.

#### func BoolP 

```
func BoolP(name, shorthand string, value bool, usage string) *bool
```

BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.

#### func BoolSlice 

```
func BoolSlice(name string, value []bool, usage string) *[]bool
```

BoolSlice defines a []bool flag with specified name, default value, and usage string. The return value is the address of a []bool variable that stores the value of the flag.

#### func BoolSliceP 

```
func BoolSliceP(name, shorthand string, value []bool, usage string) *[]bool
```

BoolSliceP is like BoolSlice, but accepts a shorthand letter that can be used after a single dash.

#### func BoolSliceVar 

```
func BoolSliceVar(p *[]bool, name string, value []bool, usage string)
```

BoolSliceVar defines a []bool flag with specified name, default value, and usage string. The argument p points to a []bool variable in which to store the value of the flag.

#### func BoolSliceVarP 

```
func BoolSliceVarP(p *[]bool, name, shorthand string, value []bool, usage string)
```

BoolSliceVarP is like BoolSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func BoolVar 

```
func BoolVar(p *bool, name string, value bool, usage string)
```

BoolVar defines a bool flag with specified name, default value, and usage string. The argument p points to a bool variable in which to store the value of the flag.

#### func BoolVarP 

```
func BoolVarP(p *bool, name, shorthand string, value bool, usage string)
```

BoolVarP is like BoolVar, but accepts a shorthand letter that can be used after a single dash.

#### func BytesBase64  <- v1.0.2

```
func BytesBase64(name string, value []byte, usage string) *[]byte
```

BytesBase64 defines an []byte flag with specified name, default value, and usage string. The return value is the address of an []byte variable that stores the value of the flag.

#### func BytesBase64P  <- v1.0.2

```
func BytesBase64P(name, shorthand string, value []byte, usage string) *[]byte
```

BytesBase64P is like BytesBase64, but accepts a shorthand letter that can be used after a single dash.

#### func BytesBase64Var  <- v1.0.2

```
func BytesBase64Var(p *[]byte, name string, value []byte, usage string)
```

BytesBase64Var defines an []byte flag with specified name, default value, and usage string. The argument p points to an []byte variable in which to store the value of the flag.

#### func BytesBase64VarP  <- v1.0.2

```
func BytesBase64VarP(p *[]byte, name, shorthand string, value []byte, usage string)
```

BytesBase64VarP is like BytesBase64Var, but accepts a shorthand letter that can be used after a single dash.

#### func BytesHex  <- v1.0.1

```
func BytesHex(name string, value []byte, usage string) *[]byte
```

BytesHex defines an []byte flag with specified name, default value, and usage string. The return value is the address of an []byte variable that stores the value of the flag.

#### func BytesHexP  <- v1.0.1

```
func BytesHexP(name, shorthand string, value []byte, usage string) *[]byte
```

BytesHexP is like BytesHex, but accepts a shorthand letter that can be used after a single dash.

#### func BytesHexVar  <- v1.0.1

```
func BytesHexVar(p *[]byte, name string, value []byte, usage string)
```

BytesHexVar defines an []byte flag with specified name, default value, and usage string. The argument p points to an []byte variable in which to store the value of the flag.

#### func BytesHexVarP  <- v1.0.1

```
func BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string)
```

BytesHexVarP is like BytesHexVar, but accepts a shorthand letter that can be used after a single dash.

#### func Count 

```
func Count(name string, usage string) *int
```

Count defines a count flag with specified name, default value, and usage string. The return value is the address of an int variable that stores the value of the flag. A count flag will add 1 to its value evey time it is found on the command line

#### func CountP 

```
func CountP(name, shorthand string, usage string) *int
```

CountP is like Count only takes a shorthand for the flag name.

#### func CountVar 

```
func CountVar(p *int, name string, usage string)
```

CountVar like CountVar only the flag is placed on the CommandLine instead of a given flag set

#### func CountVarP 

```
func CountVarP(p *int, name, shorthand string, usage string)
```

CountVarP is like CountVar only take a shorthand for the flag name.

#### func Duration 

```
func Duration(name string, value time.Duration, usage string) *time.Duration
```

Duration defines a time.Duration flag with specified name, default value, and usage string. The return value is the address of a time.Duration variable that stores the value of the flag.

#### func DurationP 

```
func DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration
```

DurationP is like Duration, but accepts a shorthand letter that can be used after a single dash.

#### func DurationSlice  <- v1.0.1

```
func DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration
```

DurationSlice defines a []time.Duration flag with specified name, default value, and usage string. The return value is the address of a []time.Duration variable that stores the value of the flag.

#### func DurationSliceP  <- v1.0.1

```
func DurationSliceP(name, shorthand string, value []time.Duration, usage string) *[]time.Duration
```

DurationSliceP is like DurationSlice, but accepts a shorthand letter that can be used after a single dash.

#### func DurationSliceVar  <- v1.0.1

```
func DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string)
```

DurationSliceVar defines a duration[] flag with specified name, default value, and usage string. The argument p points to a duration[] variable in which to store the value of the flag.

#### func DurationSliceVarP  <- v1.0.1

```
func DurationSliceVarP(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string)
```

DurationSliceVarP is like DurationSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func DurationVar 

```
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

DurationVar defines a time.Duration flag with specified name, default value, and usage string. The argument p points to a time.Duration variable in which to store the value of the flag.

#### func DurationVarP 

```
func DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string)
```

DurationVarP is like DurationVar, but accepts a shorthand letter that can be used after a single dash.

#### func Float32 

```
func Float32(name string, value float32, usage string) *float32
```

Float32 defines a float32 flag with specified name, default value, and usage string. The return value is the address of a float32 variable that stores the value of the flag.

#### func Float32P 

```
func Float32P(name, shorthand string, value float32, usage string) *float32
```

Float32P is like Float32, but accepts a shorthand letter that can be used after a single dash.

#### func Float32Slice  <- v1.0.5

```
func Float32Slice(name string, value []float32, usage string) *[]float32
```

Float32Slice defines a []float32 flag with specified name, default value, and usage string. The return value is the address of a []float32 variable that stores the value of the flag.

#### func Float32SliceP  <- v1.0.5

```
func Float32SliceP(name, shorthand string, value []float32, usage string) *[]float32
```

Float32SliceP is like Float32Slice, but accepts a shorthand letter that can be used after a single dash.

#### func Float32SliceVar  <- v1.0.5

```
func Float32SliceVar(p *[]float32, name string, value []float32, usage string)
```

Float32SliceVar defines a float32[] flag with specified name, default value, and usage string. The argument p points to a float32[] variable in which to store the value of the flag.

#### func Float32SliceVarP  <- v1.0.5

```
func Float32SliceVarP(p *[]float32, name, shorthand string, value []float32, usage string)
```

Float32SliceVarP is like Float32SliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func Float32Var 

```
func Float32Var(p *float32, name string, value float32, usage string)
```

Float32Var defines a float32 flag with specified name, default value, and usage string. The argument p points to a float32 variable in which to store the value of the flag.

#### func Float32VarP 

```
func Float32VarP(p *float32, name, shorthand string, value float32, usage string)
```

Float32VarP is like Float32Var, but accepts a shorthand letter that can be used after a single dash.

#### func Float64 

```
func Float64(name string, value float64, usage string) *float64
```

Float64 defines a float64 flag with specified name, default value, and usage string. The return value is the address of a float64 variable that stores the value of the flag.

#### func Float64P 

```
func Float64P(name, shorthand string, value float64, usage string) *float64
```

Float64P is like Float64, but accepts a shorthand letter that can be used after a single dash.

#### func Float64Slice  <- v1.0.5

```
func Float64Slice(name string, value []float64, usage string) *[]float64
```

Float64Slice defines a []float64 flag with specified name, default value, and usage string. The return value is the address of a []float64 variable that stores the value of the flag.

#### func Float64SliceP  <- v1.0.5

```
func Float64SliceP(name, shorthand string, value []float64, usage string) *[]float64
```

Float64SliceP is like Float64Slice, but accepts a shorthand letter that can be used after a single dash.

#### func Float64SliceVar  <- v1.0.5

```
func Float64SliceVar(p *[]float64, name string, value []float64, usage string)
```

Float64SliceVar defines a float64[] flag with specified name, default value, and usage string. The argument p points to a float64[] variable in which to store the value of the flag.

#### func Float64SliceVarP  <- v1.0.5

```
func Float64SliceVarP(p *[]float64, name, shorthand string, value []float64, usage string)
```

Float64SliceVarP is like Float64SliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func Float64Var 

```
func Float64Var(p *float64, name string, value float64, usage string)
```

Float64Var defines a float64 flag with specified name, default value, and usage string. The argument p points to a float64 variable in which to store the value of the flag.

#### func Float64VarP 

```
func Float64VarP(p *float64, name, shorthand string, value float64, usage string)
```

Float64VarP is like Float64Var, but accepts a shorthand letter that can be used after a single dash.

#### func IP 

```
func IP(name string, value net.IP, usage string) *net.IP
```

IP defines an net.IP flag with specified name, default value, and usage string. The return value is the address of an net.IP variable that stores the value of the flag.

#### func IPMask 

```
func IPMask(name string, value net.IPMask, usage string) *net.IPMask
```

IPMask defines an net.IPMask flag with specified name, default value, and usage string. The return value is the address of an net.IPMask variable that stores the value of the flag.

#### func IPMaskP 

```
func IPMaskP(name, shorthand string, value net.IPMask, usage string) *net.IPMask
```

IPMaskP is like IP, but accepts a shorthand letter that can be used after a single dash.

#### func IPMaskVar 

```
func IPMaskVar(p *net.IPMask, name string, value net.IPMask, usage string)
```

IPMaskVar defines an net.IPMask flag with specified name, default value, and usage string. The argument p points to an net.IPMask variable in which to store the value of the flag.

#### func IPMaskVarP 

```
func IPMaskVarP(p *net.IPMask, name, shorthand string, value net.IPMask, usage string)
```

IPMaskVarP is like IPMaskVar, but accepts a shorthand letter that can be used after a single dash.

#### func IPNet 

```
func IPNet(name string, value net.IPNet, usage string) *net.IPNet
```

IPNet defines an net.IPNet flag with specified name, default value, and usage string. The return value is the address of an net.IPNet variable that stores the value of the flag.

#### func IPNetP 

```
func IPNetP(name, shorthand string, value net.IPNet, usage string) *net.IPNet
```

IPNetP is like IPNet, but accepts a shorthand letter that can be used after a single dash.

#### func IPNetVar 

```
func IPNetVar(p *net.IPNet, name string, value net.IPNet, usage string)
```

IPNetVar defines an net.IPNet flag with specified name, default value, and usage string. The argument p points to an net.IPNet variable in which to store the value of the flag.

#### func IPNetVarP 

```
func IPNetVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string)
```

IPNetVarP is like IPNetVar, but accepts a shorthand letter that can be used after a single dash.

#### func IPP 

```
func IPP(name, shorthand string, value net.IP, usage string) *net.IP
```

IPP is like IP, but accepts a shorthand letter that can be used after a single dash.

#### func IPSlice 

```
func IPSlice(name string, value []net.IP, usage string) *[]net.IP
```

IPSlice defines a []net.IP flag with specified name, default value, and usage string. The return value is the address of a []net.IP variable that stores the value of the flag.

#### func IPSliceP 

```
func IPSliceP(name, shorthand string, value []net.IP, usage string) *[]net.IP
```

IPSliceP is like IPSlice, but accepts a shorthand letter that can be used after a single dash.

#### func IPSliceVar 

```
func IPSliceVar(p *[]net.IP, name string, value []net.IP, usage string)
```

IPSliceVar defines a []net.IP flag with specified name, default value, and usage string. The argument p points to a []net.IP variable in which to store the value of the flag.

#### func IPSliceVarP 

```
func IPSliceVarP(p *[]net.IP, name, shorthand string, value []net.IP, usage string)
```

IPSliceVarP is like IPSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func IPVar 

```
func IPVar(p *net.IP, name string, value net.IP, usage string)
```

IPVar defines an net.IP flag with specified name, default value, and usage string. The argument p points to an net.IP variable in which to store the value of the flag.

#### func IPVarP 

```
func IPVarP(p *net.IP, name, shorthand string, value net.IP, usage string)
```

IPVarP is like IPVar, but accepts a shorthand letter that can be used after a single dash.

#### func Int 

```
func Int(name string, value int, usage string) *int
```

Int defines an int flag with specified name, default value, and usage string. The return value is the address of an int variable that stores the value of the flag.

#### func Int16  <- v1.0.1

```
func Int16(name string, value int16, usage string) *int16
```

Int16 defines an int16 flag with specified name, default value, and usage string. The return value is the address of an int16 variable that stores the value of the flag.

#### func Int16P  <- v1.0.1

```
func Int16P(name, shorthand string, value int16, usage string) *int16
```

Int16P is like Int16, but accepts a shorthand letter that can be used after a single dash.

#### func Int16Var  <- v1.0.1

```
func Int16Var(p *int16, name string, value int16, usage string)
```

Int16Var defines an int16 flag with specified name, default value, and usage string. The argument p points to an int16 variable in which to store the value of the flag.

#### func Int16VarP  <- v1.0.1

```
func Int16VarP(p *int16, name, shorthand string, value int16, usage string)
```

Int16VarP is like Int16Var, but accepts a shorthand letter that can be used after a single dash.

#### func Int32 

```
func Int32(name string, value int32, usage string) *int32
```

Int32 defines an int32 flag with specified name, default value, and usage string. The return value is the address of an int32 variable that stores the value of the flag.

#### func Int32P 

```
func Int32P(name, shorthand string, value int32, usage string) *int32
```

Int32P is like Int32, but accepts a shorthand letter that can be used after a single dash.

#### func Int32Slice  <- v1.0.5

```
func Int32Slice(name string, value []int32, usage string) *[]int32
```

Int32Slice defines a []int32 flag with specified name, default value, and usage string. The return value is the address of a []int32 variable that stores the value of the flag.

#### func Int32SliceP  <- v1.0.5

```
func Int32SliceP(name, shorthand string, value []int32, usage string) *[]int32
```

Int32SliceP is like Int32Slice, but accepts a shorthand letter that can be used after a single dash.

#### func Int32SliceVar  <- v1.0.5

```
func Int32SliceVar(p *[]int32, name string, value []int32, usage string)
```

Int32SliceVar defines a int32[] flag with specified name, default value, and usage string. The argument p points to a int32[] variable in which to store the value of the flag.

#### func Int32SliceVarP  <- v1.0.5

```
func Int32SliceVarP(p *[]int32, name, shorthand string, value []int32, usage string)
```

Int32SliceVarP is like Int32SliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func Int32Var 

```
func Int32Var(p *int32, name string, value int32, usage string)
```

Int32Var defines an int32 flag with specified name, default value, and usage string. The argument p points to an int32 variable in which to store the value of the flag.

#### func Int32VarP 

```
func Int32VarP(p *int32, name, shorthand string, value int32, usage string)
```

Int32VarP is like Int32Var, but accepts a shorthand letter that can be used after a single dash.

#### func Int64 

```
func Int64(name string, value int64, usage string) *int64
```

Int64 defines an int64 flag with specified name, default value, and usage string. The return value is the address of an int64 variable that stores the value of the flag.

#### func Int64P 

```
func Int64P(name, shorthand string, value int64, usage string) *int64
```

Int64P is like Int64, but accepts a shorthand letter that can be used after a single dash.

#### func Int64Slice  <- v1.0.5

```
func Int64Slice(name string, value []int64, usage string) *[]int64
```

Int64Slice defines a []int64 flag with specified name, default value, and usage string. The return value is the address of a []int64 variable that stores the value of the flag.

#### func Int64SliceP  <- v1.0.5

```
func Int64SliceP(name, shorthand string, value []int64, usage string) *[]int64
```

Int64SliceP is like Int64Slice, but accepts a shorthand letter that can be used after a single dash.

#### func Int64SliceVar  <- v1.0.5

```
func Int64SliceVar(p *[]int64, name string, value []int64, usage string)
```

Int64SliceVar defines a int64[] flag with specified name, default value, and usage string. The argument p points to a int64[] variable in which to store the value of the flag.

#### func Int64SliceVarP  <- v1.0.5

```
func Int64SliceVarP(p *[]int64, name, shorthand string, value []int64, usage string)
```

Int64SliceVarP is like Int64SliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func Int64Var 

```
func Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var defines an int64 flag with specified name, default value, and usage string. The argument p points to an int64 variable in which to store the value of the flag.

#### func Int64VarP 

```
func Int64VarP(p *int64, name, shorthand string, value int64, usage string)
```

Int64VarP is like Int64Var, but accepts a shorthand letter that can be used after a single dash.

#### func Int8 

```
func Int8(name string, value int8, usage string) *int8
```

Int8 defines an int8 flag with specified name, default value, and usage string. The return value is the address of an int8 variable that stores the value of the flag.

#### func Int8P 

```
func Int8P(name, shorthand string, value int8, usage string) *int8
```

Int8P is like Int8, but accepts a shorthand letter that can be used after a single dash.

#### func Int8Var 

```
func Int8Var(p *int8, name string, value int8, usage string)
```

Int8Var defines an int8 flag with specified name, default value, and usage string. The argument p points to an int8 variable in which to store the value of the flag.

#### func Int8VarP 

```
func Int8VarP(p *int8, name, shorthand string, value int8, usage string)
```

Int8VarP is like Int8Var, but accepts a shorthand letter that can be used after a single dash.

#### func IntP 

```
func IntP(name, shorthand string, value int, usage string) *int
```

IntP is like Int, but accepts a shorthand letter that can be used after a single dash.

#### func IntSlice 

```
func IntSlice(name string, value []int, usage string) *[]int
```

IntSlice defines a []int flag with specified name, default value, and usage string. The return value is the address of a []int variable that stores the value of the flag.

#### func IntSliceP 

```
func IntSliceP(name, shorthand string, value []int, usage string) *[]int
```

IntSliceP is like IntSlice, but accepts a shorthand letter that can be used after a single dash.

#### func IntSliceVar 

```
func IntSliceVar(p *[]int, name string, value []int, usage string)
```

IntSliceVar defines a int[] flag with specified name, default value, and usage string. The argument p points to a int[] variable in which to store the value of the flag.

#### func IntSliceVarP 

```
func IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string)
```

IntSliceVarP is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func IntVar 

```
func IntVar(p *int, name string, value int, usage string)
```

IntVar defines an int flag with specified name, default value, and usage string. The argument p points to an int variable in which to store the value of the flag.

#### func IntVarP 

```
func IntVarP(p *int, name, shorthand string, value int, usage string)
```

IntVarP is like IntVar, but accepts a shorthand letter that can be used after a single dash.

#### func NArg 

```
func NArg() int
```

NArg is the number of arguments remaining after flags have been processed.

#### func NFlag 

```
func NFlag() int
```

NFlag returns the number of command-line flags that have been set.

#### func Parse 

```
func Parse()
```

Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.

#### func ParseAll 

```
func ParseAll(fn func(flag *Flag, value string) error)
```

ParseAll parses the command-line flags from os.Args[1:] and called fn for each. The arguments for fn are flag and value. Must be called after all flags are defined and before flags are accessed by the program.

#### func ParseIPv4Mask 

```
func ParseIPv4Mask(s string) net.IPMask
```

ParseIPv4Mask written in IP form (e.g. 255.255.255.0). This function should really belong to the net package.

#### func Parsed 

```
func Parsed() bool
```

Parsed returns true if the command-line flags have been parsed.

#### func PrintDefaults 

```
func PrintDefaults()
```

PrintDefaults prints to standard error the default values of all defined command-line flags.

#### func Set 

```
func Set(name, value string) error
```

Set sets the value of the named command-line flag.

#### func SetInterspersed 

```
func SetInterspersed(interspersed bool)
```

SetInterspersed sets whether to support interspersed option/non-option arguments.

#### func String 

```
func String(name string, value string, usage string) *string
```

String defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag.

#### func StringArray 

```
func StringArray(name string, value []string, usage string) *[]string
```

StringArray defines a string flag with specified name, default value, and usage string. The return value is the address of a []string variable that stores the value of the flag. The value of each argument will not try to be separated by comma. Use a StringSlice for that.

#### func StringArrayP 

```
func StringArrayP(name, shorthand string, value []string, usage string) *[]string
```

StringArrayP is like StringArray, but accepts a shorthand letter that can be used after a single dash.

#### func StringArrayVar 

```
func StringArrayVar(p *[]string, name string, value []string, usage string)
```

StringArrayVar defines a string flag with specified name, default value, and usage string. The argument p points to a []string variable in which to store the value of the flag. The value of each argument will not try to be separated by comma. Use a StringSlice for that.

#### func StringArrayVarP 

```
func StringArrayVarP(p *[]string, name, shorthand string, value []string, usage string)
```

StringArrayVarP is like StringArrayVar, but accepts a shorthand letter that can be used after a single dash.

#### func StringP 

```
func StringP(name, shorthand string, value string, usage string) *string
```

StringP is like String, but accepts a shorthand letter that can be used after a single dash.

#### func StringSlice 

```
func StringSlice(name string, value []string, usage string) *[]string
```

StringSlice defines a string flag with specified name, default value, and usage string. The return value is the address of a []string variable that stores the value of the flag. Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly. For example:

```
--ss="v1,v2" --ss="v3"
```

will result in

```
[]string{"v1", "v2", "v3"}
```

#### func StringSliceP 

```
func StringSliceP(name, shorthand string, value []string, usage string) *[]string
```

StringSliceP is like StringSlice, but accepts a shorthand letter that can be used after a single dash.

#### func StringSliceVar 

```
func StringSliceVar(p *[]string, name string, value []string, usage string)
```

StringSliceVar defines a string flag with specified name, default value, and usage string. The argument p points to a []string variable in which to store the value of the flag. Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly. For example:

```
--ss="v1,v2" --ss="v3"
```

will result in

```
[]string{"v1", "v2", "v3"}
```

#### func StringSliceVarP 

```
func StringSliceVarP(p *[]string, name, shorthand string, value []string, usage string)
```

StringSliceVarP is like StringSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func StringToInt  <- v1.0.3

```
func StringToInt(name string, value map[string]int, usage string) *map[string]int
```

StringToInt defines a string flag with specified name, default value, and usage string. The return value is the address of a map[string]int variable that stores the value of the flag. The value of each argument will not try to be separated by comma

#### func StringToInt64  <- v1.0.5

```
func StringToInt64(name string, value map[string]int64, usage string) *map[string]int64
```

StringToInt64 defines a string flag with specified name, default value, and usage string. The return value is the address of a map[string]int64 variable that stores the value of the flag. The value of each argument will not try to be separated by comma

#### func StringToInt64P  <- v1.0.5

```
func StringToInt64P(name, shorthand string, value map[string]int64, usage string) *map[string]int64
```

StringToInt64P is like StringToInt64, but accepts a shorthand letter that can be used after a single dash.

#### func StringToInt64Var  <- v1.0.5

```
func StringToInt64Var(p *map[string]int64, name string, value map[string]int64, usage string)
```

StringToInt64Var defines a string flag with specified name, default value, and usage string. The argument p point64s to a map[string]int64 variable in which to store the value of the flag. The value of each argument will not try to be separated by comma

#### func StringToInt64VarP  <- v1.0.5

```
func StringToInt64VarP(p *map[string]int64, name, shorthand string, value map[string]int64, usage string)
```

StringToInt64VarP is like StringToInt64Var, but accepts a shorthand letter that can be used after a single dash.

#### func StringToIntP  <- v1.0.3

```
func StringToIntP(name, shorthand string, value map[string]int, usage string) *map[string]int
```

StringToIntP is like StringToInt, but accepts a shorthand letter that can be used after a single dash.

#### func StringToIntVar  <- v1.0.3

```
func StringToIntVar(p *map[string]int, name string, value map[string]int, usage string)
```

StringToIntVar defines a string flag with specified name, default value, and usage string. The argument p points to a map[string]int variable in which to store the value of the flag. The value of each argument will not try to be separated by comma

#### func StringToIntVarP  <- v1.0.3

```
func StringToIntVarP(p *map[string]int, name, shorthand string, value map[string]int, usage string)
```

StringToIntVarP is like StringToIntVar, but accepts a shorthand letter that can be used after a single dash.

#### func StringToString  <- v1.0.3

```
func StringToString(name string, value map[string]string, usage string) *map[string]string
```

StringToString defines a string flag with specified name, default value, and usage string. The return value is the address of a map[string]string variable that stores the value of the flag. The value of each argument will not try to be separated by comma

#### func StringToStringP  <- v1.0.3

```
func StringToStringP(name, shorthand string, value map[string]string, usage string) *map[string]string
```

StringToStringP is like StringToString, but accepts a shorthand letter that can be used after a single dash.

#### func StringToStringVar  <- v1.0.3

```
func StringToStringVar(p *map[string]string, name string, value map[string]string, usage string)
```

StringToStringVar defines a string flag with specified name, default value, and usage string. The argument p points to a map[string]string variable in which to store the value of the flag. The value of each argument will not try to be separated by comma

#### func StringToStringVarP  <- v1.0.3

```
func StringToStringVarP(p *map[string]string, name, shorthand string, value map[string]string, usage string)
```

StringToStringVarP is like StringToStringVar, but accepts a shorthand letter that can be used after a single dash.

#### func StringVar 

```
func StringVar(p *string, name string, value string, usage string)
```

StringVar defines a string flag with specified name, default value, and usage string. The argument p points to a string variable in which to store the value of the flag.

#### func StringVarP 

```
func StringVarP(p *string, name, shorthand string, value string, usage string)
```

StringVarP is like StringVar, but accepts a shorthand letter that can be used after a single dash.

#### func Uint 

```
func Uint(name string, value uint, usage string) *uint
```

Uint defines a uint flag with specified name, default value, and usage string. The return value is the address of a uint variable that stores the value of the flag.

#### func Uint16 

```
func Uint16(name string, value uint16, usage string) *uint16
```

Uint16 defines a uint flag with specified name, default value, and usage string. The return value is the address of a uint variable that stores the value of the flag.

#### func Uint16P 

```
func Uint16P(name, shorthand string, value uint16, usage string) *uint16
```

Uint16P is like Uint16, but accepts a shorthand letter that can be used after a single dash.

#### func Uint16Var 

```
func Uint16Var(p *uint16, name string, value uint16, usage string)
```

Uint16Var defines a uint flag with specified name, default value, and usage string. The argument p points to a uint variable in which to store the value of the flag.

#### func Uint16VarP 

```
func Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string)
```

Uint16VarP is like Uint16Var, but accepts a shorthand letter that can be used after a single dash.

#### func Uint32 

```
func Uint32(name string, value uint32, usage string) *uint32
```

Uint32 defines a uint32 flag with specified name, default value, and usage string. The return value is the address of a uint32 variable that stores the value of the flag.

#### func Uint32P 

```
func Uint32P(name, shorthand string, value uint32, usage string) *uint32
```

Uint32P is like Uint32, but accepts a shorthand letter that can be used after a single dash.

#### func Uint32Var 

```
func Uint32Var(p *uint32, name string, value uint32, usage string)
```

Uint32Var defines a uint32 flag with specified name, default value, and usage string. The argument p points to a uint32 variable in which to store the value of the flag.

#### func Uint32VarP 

```
func Uint32VarP(p *uint32, name, shorthand string, value uint32, usage string)
```

Uint32VarP is like Uint32Var, but accepts a shorthand letter that can be used after a single dash.

#### func Uint64 

```
func Uint64(name string, value uint64, usage string) *uint64
```

Uint64 defines a uint64 flag with specified name, default value, and usage string. The return value is the address of a uint64 variable that stores the value of the flag.

#### func Uint64P 

```
func Uint64P(name, shorthand string, value uint64, usage string) *uint64
```

Uint64P is like Uint64, but accepts a shorthand letter that can be used after a single dash.

#### func Uint64Var 

```
func Uint64Var(p *uint64, name string, value uint64, usage string)
```

Uint64Var defines a uint64 flag with specified name, default value, and usage string. The argument p points to a uint64 variable in which to store the value of the flag.

#### func Uint64VarP 

```
func Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string)
```

Uint64VarP is like Uint64Var, but accepts a shorthand letter that can be used after a single dash.

#### func Uint8 

```
func Uint8(name string, value uint8, usage string) *uint8
```

Uint8 defines a uint8 flag with specified name, default value, and usage string. The return value is the address of a uint8 variable that stores the value of the flag.

#### func Uint8P 

```
func Uint8P(name, shorthand string, value uint8, usage string) *uint8
```

Uint8P is like Uint8, but accepts a shorthand letter that can be used after a single dash.

#### func Uint8Var 

```
func Uint8Var(p *uint8, name string, value uint8, usage string)
```

Uint8Var defines a uint8 flag with specified name, default value, and usage string. The argument p points to a uint8 variable in which to store the value of the flag.

#### func Uint8VarP 

```
func Uint8VarP(p *uint8, name, shorthand string, value uint8, usage string)
```

Uint8VarP is like Uint8Var, but accepts a shorthand letter that can be used after a single dash.

#### func UintP 

```
func UintP(name, shorthand string, value uint, usage string) *uint
```

UintP is like Uint, but accepts a shorthand letter that can be used after a single dash.

#### func UintSlice 

```
func UintSlice(name string, value []uint, usage string) *[]uint
```

UintSlice defines a []uint flag with specified name, default value, and usage string. The return value is the address of a []uint variable that stores the value of the flag.

#### func UintSliceP 

```
func UintSliceP(name, shorthand string, value []uint, usage string) *[]uint
```

UintSliceP is like UintSlice, but accepts a shorthand letter that can be used after a single dash.

#### func UintSliceVar 

```
func UintSliceVar(p *[]uint, name string, value []uint, usage string)
```

UintSliceVar defines a uint[] flag with specified name, default value, and usage string. The argument p points to a uint[] variable in which to store the value of the flag.

#### func UintSliceVarP 

```
func UintSliceVarP(p *[]uint, name, shorthand string, value []uint, usage string)
```

UintSliceVarP is like the UintSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### func UintVar 

```
func UintVar(p *uint, name string, value uint, usage string)
```

UintVar defines a uint flag with specified name, default value, and usage string. The argument p points to a uint variable in which to store the value of the flag.

#### func UintVarP 

```
func UintVarP(p *uint, name, shorthand string, value uint, usage string)
```

UintVarP is like UintVar, but accepts a shorthand letter that can be used after a single dash.

#### func UnquoteUsage 

```
func UnquoteUsage(flag *Flag) (name string, usage string)
```

UnquoteUsage extracts a back-quoted name from the usage string for a flag and returns it and the un-quoted usage. Given "a `name` to show" it returns ("name", "a name to show"). If there are no back quotes, the name is an educated guess of the type of the flag's value, or the empty string if the flag is boolean.

#### func Var 

```
func Var(value Value, name string, usage string)
```

Var defines a flag with the specified name and usage string. The type and value of the flag are represented by the first argument, of type Value, which typically holds a user-defined implementation of Value. For instance, the caller could create a flag that turns a comma-separated string into a slice of strings by giving the slice the methods of Value; in particular, Set would decompose the comma-separated string into the slice.

#### func VarP 

```
func VarP(value Value, name, shorthand, usage string)
```

VarP is like Var, but accepts a shorthand letter that can be used after a single dash.

#### func Visit 

```
func Visit(fn func(*Flag))
```

Visit visits the command-line flags in lexicographical order or in primordial order if f.SortFlags is false, calling fn for each. It visits only those flags that have been set.

#### func VisitAll 

```
func VisitAll(fn func(*Flag))
```

VisitAll visits the command-line flags in lexicographical order or in primordial order if f.SortFlags is false, calling fn for each. It visits all flags, even those not set.

### Types 

#### type ErrorHandling 

```
type ErrorHandling int
```

ErrorHandling defines how to handle flag parsing errors.

```
const (
	// ContinueOnError will return an err from Parse() if an error is found
	ContinueOnError ErrorHandling = iota
	// ExitOnError will call os.Exit(2) if an error is found when parsing
	ExitOnError
	// PanicOnError will panic() if an error is found when parsing flags
	PanicOnError
)
```

#### type Flag 

```
type Flag struct {
	Name                string              // name as it appears on command line
	Shorthand           string              // one-letter abbreviated flag
	Usage               string              // help message
	Value               Value               // value as set
	DefValue            string              // default value (as text); for usage message
	Changed             bool                // If the user set the value (or if left to default)
	NoOptDefVal         string              // default value (as text); if the flag is on the command line without any options
	Deprecated          string              // If this flag is deprecated, this string is the new or now thing to use
	Hidden              bool                // used by cobra.Command to allow flags to be hidden from help/usage text
	ShorthandDeprecated string              // If the shorthand of this flag is deprecated, this string is the new or now thing to use
	Annotations         map[string][]string // used by cobra.Command bash autocomple code
}
```

A Flag represents the state of a flag.

#### func Lookup 

```
func Lookup(name string) *Flag
```

Lookup returns the Flag structure of the named command-line flag, returning nil if none exists.

#### func PFlagFromGoFlag 

```
func PFlagFromGoFlag(goflag *goflag.Flag) *Flag
```

PFlagFromGoFlag will return a *pflag.Flag given a *flag.Flag If the *flag.Flag.Name was a single character (ex: `v`) it will be accessiblei with both `-v` and `--v` in flags. If the golang flag was more than a single character (ex: `verbose`) it will only be accessible via `--verbose`

#### func ShorthandLookup 

```
func ShorthandLookup(name string) *Flag
```

ShorthandLookup returns the Flag structure of the short handed flag, returning nil if none exists.

##### Example

#### type FlagSet 

```
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler.
	Usage func()

	// SortFlags is used to indicate, if user wants to have sorted flags in
	// help/usage messages.
	SortFlags bool

	// ParseErrorsWhitelist is used to configure a whitelist of errors
	ParseErrorsWhitelist ParseErrorsWhitelist
	// contains filtered or unexported fields
}
```

A FlagSet represents a set of defined flags.

#### func NewFlagSet 

```
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
```

NewFlagSet returns a new, empty flag set with the specified name, error handling property and SortFlags set to true.

#### (*FlagSet) AddFlag 

```
func (f *FlagSet) AddFlag(flag *Flag)
```

AddFlag will add the flag to the FlagSet

#### (*FlagSet) AddFlagSet 

```
func (f *FlagSet) AddFlagSet(newSet *FlagSet)
```

AddFlagSet adds one FlagSet to another. If a flag is already present in f the flag from newSet will be ignored.

#### (*FlagSet) AddGoFlag 

```
func (f *FlagSet) AddGoFlag(goflag *goflag.Flag)
```

AddGoFlag will add the given *flag.Flag to the pflag.FlagSet

#### (*FlagSet) AddGoFlagSet 

```
func (f *FlagSet) AddGoFlagSet(newSet *goflag.FlagSet)
```

AddGoFlagSet will add the given *flag.FlagSet to the pflag.FlagSet

#### (*FlagSet) Arg 

```
func (f *FlagSet) Arg(i int) string
```

Arg returns the i'th argument. Arg(0) is the first remaining argument after flags have been processed.

#### (*FlagSet) Args 

```
func (f *FlagSet) Args() []string
```

Args returns the non-flag arguments.

#### (*FlagSet) ArgsLenAtDash 

```
func (f *FlagSet) ArgsLenAtDash() int
```

ArgsLenAtDash will return the length of f.Args at the moment when a -- was found during arg parsing. This allows your program to know which args were before the -- and which came after.

#### (*FlagSet) Bool 

```
func (f *FlagSet) Bool(name string, value bool, usage string) *bool
```

Bool defines a bool flag with specified name, default value, and usage string. The return value is the address of a bool variable that stores the value of the flag.

#### (*FlagSet) BoolP 

```
func (f *FlagSet) BoolP(name, shorthand string, value bool, usage string) *bool
```

BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) BoolSlice 

```
func (f *FlagSet) BoolSlice(name string, value []bool, usage string) *[]bool
```

BoolSlice defines a []bool flag with specified name, default value, and usage string. The return value is the address of a []bool variable that stores the value of the flag.

#### (*FlagSet) BoolSliceP 

```
func (f *FlagSet) BoolSliceP(name, shorthand string, value []bool, usage string) *[]bool
```

BoolSliceP is like BoolSlice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) BoolSliceVar 

```
func (f *FlagSet) BoolSliceVar(p *[]bool, name string, value []bool, usage string)
```

BoolSliceVar defines a boolSlice flag with specified name, default value, and usage string. The argument p points to a []bool variable in which to store the value of the flag.

#### (*FlagSet) BoolSliceVarP 

```
func (f *FlagSet) BoolSliceVarP(p *[]bool, name, shorthand string, value []bool, usage string)
```

BoolSliceVarP is like BoolSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) BoolVar 

```
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
```

BoolVar defines a bool flag with specified name, default value, and usage string. The argument p points to a bool variable in which to store the value of the flag.

#### (*FlagSet) BoolVarP 

```
func (f *FlagSet) BoolVarP(p *bool, name, shorthand string, value bool, usage string)
```

BoolVarP is like BoolVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) BytesBase64  <- v1.0.2

```
func (f *FlagSet) BytesBase64(name string, value []byte, usage string) *[]byte
```

BytesBase64 defines an []byte flag with specified name, default value, and usage string. The return value is the address of an []byte variable that stores the value of the flag.

#### (*FlagSet) BytesBase64P  <- v1.0.2

```
func (f *FlagSet) BytesBase64P(name, shorthand string, value []byte, usage string) *[]byte
```

BytesBase64P is like BytesBase64, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) BytesBase64Var  <- v1.0.2

```
func (f *FlagSet) BytesBase64Var(p *[]byte, name string, value []byte, usage string)
```

BytesBase64Var defines an []byte flag with specified name, default value, and usage string. The argument p points to an []byte variable in which to store the value of the flag.

#### (*FlagSet) BytesBase64VarP  <- v1.0.2

```
func (f *FlagSet) BytesBase64VarP(p *[]byte, name, shorthand string, value []byte, usage string)
```

BytesBase64VarP is like BytesBase64Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) BytesHex  <- v1.0.1

```
func (f *FlagSet) BytesHex(name string, value []byte, usage string) *[]byte
```

BytesHex defines an []byte flag with specified name, default value, and usage string. The return value is the address of an []byte variable that stores the value of the flag.

#### (*FlagSet) BytesHexP  <- v1.0.1

```
func (f *FlagSet) BytesHexP(name, shorthand string, value []byte, usage string) *[]byte
```

BytesHexP is like BytesHex, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) BytesHexVar  <- v1.0.1

```
func (f *FlagSet) BytesHexVar(p *[]byte, name string, value []byte, usage string)
```

BytesHexVar defines an []byte flag with specified name, default value, and usage string. The argument p points to an []byte variable in which to store the value of the flag.

#### (*FlagSet) BytesHexVarP  <- v1.0.1

```
func (f *FlagSet) BytesHexVarP(p *[]byte, name, shorthand string, value []byte, usage string)
```

BytesHexVarP is like BytesHexVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Changed 

```
func (f *FlagSet) Changed(name string) bool
```

Changed returns true if the flag was explicitly set during Parse() and false otherwise

#### (*FlagSet) Count 

```
func (f *FlagSet) Count(name string, usage string) *int
```

Count defines a count flag with specified name, default value, and usage string. The return value is the address of an int variable that stores the value of the flag. A count flag will add 1 to its value every time it is found on the command line

#### (*FlagSet) CountP 

```
func (f *FlagSet) CountP(name, shorthand string, usage string) *int
```

CountP is like Count only takes a shorthand for the flag name.

#### (*FlagSet) CountVar 

```
func (f *FlagSet) CountVar(p *int, name string, usage string)
```

CountVar defines a count flag with specified name, default value, and usage string. The argument p points to an int variable in which to store the value of the flag. A count flag will add 1 to its value every time it is found on the command line

#### (*FlagSet) CountVarP 

```
func (f *FlagSet) CountVarP(p *int, name, shorthand string, usage string)
```

CountVarP is like CountVar only take a shorthand for the flag name.

#### (*FlagSet) Duration 

```
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
```

Duration defines a time.Duration flag with specified name, default value, and usage string. The return value is the address of a time.Duration variable that stores the value of the flag.

#### (*FlagSet) DurationP 

```
func (f *FlagSet) DurationP(name, shorthand string, value time.Duration, usage string) *time.Duration
```

DurationP is like Duration, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) DurationSlice  <- v1.0.1

```
func (f *FlagSet) DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration
```

DurationSlice defines a []time.Duration flag with specified name, default value, and usage string. The return value is the address of a []time.Duration variable that stores the value of the flag.

#### (*FlagSet) DurationSliceP  <- v1.0.1

```
func (f *FlagSet) DurationSliceP(name, shorthand string, value []time.Duration, usage string) *[]time.Duration
```

DurationSliceP is like DurationSlice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) DurationSliceVar  <- v1.0.1

```
func (f *FlagSet) DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string)
```

DurationSliceVar defines a durationSlice flag with specified name, default value, and usage string. The argument p points to a []time.Duration variable in which to store the value of the flag.

#### (*FlagSet) DurationSliceVarP  <- v1.0.1

```
func (f *FlagSet) DurationSliceVarP(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string)
```

DurationSliceVarP is like DurationSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) DurationVar 

```
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

DurationVar defines a time.Duration flag with specified name, default value, and usage string. The argument p points to a time.Duration variable in which to store the value of the flag.

#### (*FlagSet) DurationVarP 

```
func (f *FlagSet) DurationVarP(p *time.Duration, name, shorthand string, value time.Duration, usage string)
```

DurationVarP is like DurationVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) FlagUsages 

```
func (f *FlagSet) FlagUsages() string
```

FlagUsages returns a string containing the usage information for all flags in the FlagSet

#### (*FlagSet) FlagUsagesWrapped 

```
func (f *FlagSet) FlagUsagesWrapped(cols int) string
```

FlagUsagesWrapped returns a string containing the usage information for all flags in the FlagSet. Wrapped to `cols` columns (0 for no wrapping)

#### (*FlagSet) Float32 

```
func (f *FlagSet) Float32(name string, value float32, usage string) *float32
```

Float32 defines a float32 flag with specified name, default value, and usage string. The return value is the address of a float32 variable that stores the value of the flag.

#### (*FlagSet) Float32P 

```
func (f *FlagSet) Float32P(name, shorthand string, value float32, usage string) *float32
```

Float32P is like Float32, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Float32Slice  <- v1.0.5

```
func (f *FlagSet) Float32Slice(name string, value []float32, usage string) *[]float32
```

Float32Slice defines a []float32 flag with specified name, default value, and usage string. The return value is the address of a []float32 variable that stores the value of the flag.

#### (*FlagSet) Float32SliceP  <- v1.0.5

```
func (f *FlagSet) Float32SliceP(name, shorthand string, value []float32, usage string) *[]float32
```

Float32SliceP is like Float32Slice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Float32SliceVar  <- v1.0.5

```
func (f *FlagSet) Float32SliceVar(p *[]float32, name string, value []float32, usage string)
```

Float32SliceVar defines a float32Slice flag with specified name, default value, and usage string. The argument p points to a []float32 variable in which to store the value of the flag.

#### (*FlagSet) Float32SliceVarP  <- v1.0.5

```
func (f *FlagSet) Float32SliceVarP(p *[]float32, name, shorthand string, value []float32, usage string)
```

Float32SliceVarP is like Float32SliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Float32Var 

```
func (f *FlagSet) Float32Var(p *float32, name string, value float32, usage string)
```

Float32Var defines a float32 flag with specified name, default value, and usage string. The argument p points to a float32 variable in which to store the value of the flag.

#### (*FlagSet) Float32VarP 

```
func (f *FlagSet) Float32VarP(p *float32, name, shorthand string, value float32, usage string)
```

Float32VarP is like Float32Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Float64 

```
func (f *FlagSet) Float64(name string, value float64, usage string) *float64
```

Float64 defines a float64 flag with specified name, default value, and usage string. The return value is the address of a float64 variable that stores the value of the flag.

#### (*FlagSet) Float64P 

```
func (f *FlagSet) Float64P(name, shorthand string, value float64, usage string) *float64
```

Float64P is like Float64, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Float64Slice  <- v1.0.5

```
func (f *FlagSet) Float64Slice(name string, value []float64, usage string) *[]float64
```

Float64Slice defines a []float64 flag with specified name, default value, and usage string. The return value is the address of a []float64 variable that stores the value of the flag.

#### (*FlagSet) Float64SliceP  <- v1.0.5

```
func (f *FlagSet) Float64SliceP(name, shorthand string, value []float64, usage string) *[]float64
```

Float64SliceP is like Float64Slice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Float64SliceVar  <- v1.0.5

```
func (f *FlagSet) Float64SliceVar(p *[]float64, name string, value []float64, usage string)
```

Float64SliceVar defines a float64Slice flag with specified name, default value, and usage string. The argument p points to a []float64 variable in which to store the value of the flag.

#### (*FlagSet) Float64SliceVarP  <- v1.0.5

```
func (f *FlagSet) Float64SliceVarP(p *[]float64, name, shorthand string, value []float64, usage string)
```

Float64SliceVarP is like Float64SliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Float64Var 

```
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
```

Float64Var defines a float64 flag with specified name, default value, and usage string. The argument p points to a float64 variable in which to store the value of the flag.

#### (*FlagSet) Float64VarP 

```
func (f *FlagSet) Float64VarP(p *float64, name, shorthand string, value float64, usage string)
```

Float64VarP is like Float64Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) GetBool 

```
func (f *FlagSet) GetBool(name string) (bool, error)
```

GetBool return the bool value of a flag with the given name

#### (*FlagSet) GetBoolSlice 

```
func (f *FlagSet) GetBoolSlice(name string) ([]bool, error)
```

GetBoolSlice returns the []bool value of a flag with the given name.

#### (*FlagSet) GetBytesBase64  <- v1.0.2

```
func (f *FlagSet) GetBytesBase64(name string) ([]byte, error)
```

GetBytesBase64 return the []byte value of a flag with the given name

#### (*FlagSet) GetBytesHex  <- v1.0.1

```
func (f *FlagSet) GetBytesHex(name string) ([]byte, error)
```

GetBytesHex return the []byte value of a flag with the given name

#### (*FlagSet) GetCount 

```
func (f *FlagSet) GetCount(name string) (int, error)
```

GetCount return the int value of a flag with the given name

#### (*FlagSet) GetDuration 

```
func (f *FlagSet) GetDuration(name string) (time.Duration, error)
```

GetDuration return the duration value of a flag with the given name

#### (*FlagSet) GetDurationSlice  <- v1.0.1

```
func (f *FlagSet) GetDurationSlice(name string) ([]time.Duration, error)
```

GetDurationSlice returns the []time.Duration value of a flag with the given name

#### (*FlagSet) GetFloat32 

```
func (f *FlagSet) GetFloat32(name string) (float32, error)
```

GetFloat32 return the float32 value of a flag with the given name

#### (*FlagSet) GetFloat32Slice  <- v1.0.5

```
func (f *FlagSet) GetFloat32Slice(name string) ([]float32, error)
```

GetFloat32Slice return the []float32 value of a flag with the given name

#### (*FlagSet) GetFloat64 

```
func (f *FlagSet) GetFloat64(name string) (float64, error)
```

GetFloat64 return the float64 value of a flag with the given name

#### (*FlagSet) GetFloat64Slice  <- v1.0.5

```
func (f *FlagSet) GetFloat64Slice(name string) ([]float64, error)
```

GetFloat64Slice return the []float64 value of a flag with the given name

#### (*FlagSet) GetIP 

```
func (f *FlagSet) GetIP(name string) (net.IP, error)
```

GetIP return the net.IP value of a flag with the given name

#### (*FlagSet) GetIPNet 

```
func (f *FlagSet) GetIPNet(name string) (net.IPNet, error)
```

GetIPNet return the net.IPNet value of a flag with the given name

#### (*FlagSet) GetIPSlice 

```
func (f *FlagSet) GetIPSlice(name string) ([]net.IP, error)
```

GetIPSlice returns the []net.IP value of a flag with the given name

#### (*FlagSet) GetIPv4Mask 

```
func (f *FlagSet) GetIPv4Mask(name string) (net.IPMask, error)
```

GetIPv4Mask return the net.IPv4Mask value of a flag with the given name

#### (*FlagSet) GetInt 

```
func (f *FlagSet) GetInt(name string) (int, error)
```

GetInt return the int value of a flag with the given name

#### (*FlagSet) GetInt16  <- v1.0.1

```
func (f *FlagSet) GetInt16(name string) (int16, error)
```

GetInt16 returns the int16 value of a flag with the given name

#### (*FlagSet) GetInt32 

```
func (f *FlagSet) GetInt32(name string) (int32, error)
```

GetInt32 return the int32 value of a flag with the given name

#### (*FlagSet) GetInt32Slice  <- v1.0.5

```
func (f *FlagSet) GetInt32Slice(name string) ([]int32, error)
```

GetInt32Slice return the []int32 value of a flag with the given name

#### (*FlagSet) GetInt64 

```
func (f *FlagSet) GetInt64(name string) (int64, error)
```

GetInt64 return the int64 value of a flag with the given name

#### (*FlagSet) GetInt64Slice  <- v1.0.5

```
func (f *FlagSet) GetInt64Slice(name string) ([]int64, error)
```

GetInt64Slice return the []int64 value of a flag with the given name

#### (*FlagSet) GetInt8 

```
func (f *FlagSet) GetInt8(name string) (int8, error)
```

GetInt8 return the int8 value of a flag with the given name

#### (*FlagSet) GetIntSlice 

```
func (f *FlagSet) GetIntSlice(name string) ([]int, error)
```

GetIntSlice return the []int value of a flag with the given name

#### (*FlagSet) GetNormalizeFunc 

```
func (f *FlagSet) GetNormalizeFunc() func(f *FlagSet, name string) NormalizedName
```

GetNormalizeFunc returns the previously set NormalizeFunc of a function which does no translation, if not set previously.

#### (*FlagSet) GetString 

```
func (f *FlagSet) GetString(name string) (string, error)
```

GetString return the string value of a flag with the given name

#### (*FlagSet) GetStringArray 

```
func (f *FlagSet) GetStringArray(name string) ([]string, error)
```

GetStringArray return the []string value of a flag with the given name

#### (*FlagSet) GetStringSlice 

```
func (f *FlagSet) GetStringSlice(name string) ([]string, error)
```

GetStringSlice return the []string value of a flag with the given name

#### (*FlagSet) GetStringToInt  <- v1.0.3

```
func (f *FlagSet) GetStringToInt(name string) (map[string]int, error)
```

GetStringToInt return the map[string]int value of a flag with the given name

#### (*FlagSet) GetStringToInt64  <- v1.0.5

```
func (f *FlagSet) GetStringToInt64(name string) (map[string]int64, error)
```

GetStringToInt64 return the map[string]int64 value of a flag with the given name

#### (*FlagSet) GetStringToString  <- v1.0.3

```
func (f *FlagSet) GetStringToString(name string) (map[string]string, error)
```

GetStringToString return the map[string]string value of a flag with the given name

#### (*FlagSet) GetUint 

```
func (f *FlagSet) GetUint(name string) (uint, error)
```

GetUint return the uint value of a flag with the given name

#### (*FlagSet) GetUint16 

```
func (f *FlagSet) GetUint16(name string) (uint16, error)
```

GetUint16 return the uint16 value of a flag with the given name

#### (*FlagSet) GetUint32 

```
func (f *FlagSet) GetUint32(name string) (uint32, error)
```

GetUint32 return the uint32 value of a flag with the given name

#### (*FlagSet) GetUint64 

```
func (f *FlagSet) GetUint64(name string) (uint64, error)
```

GetUint64 return the uint64 value of a flag with the given name

#### (*FlagSet) GetUint8 

```
func (f *FlagSet) GetUint8(name string) (uint8, error)
```

GetUint8 return the uint8 value of a flag with the given name

#### (*FlagSet) GetUintSlice 

```
func (f *FlagSet) GetUintSlice(name string) ([]uint, error)
```

GetUintSlice returns the []uint value of a flag with the given name.

#### (*FlagSet) HasAvailableFlags 

```
func (f *FlagSet) HasAvailableFlags() bool
```

HasAvailableFlags returns a bool to indicate if the FlagSet has any flags that are not hidden.

#### (*FlagSet) HasFlags 

```
func (f *FlagSet) HasFlags() bool
```

HasFlags returns a bool to indicate if the FlagSet has any flags defined.

#### (*FlagSet) IP 

```
func (f *FlagSet) IP(name string, value net.IP, usage string) *net.IP
```

IP defines an net.IP flag with specified name, default value, and usage string. The return value is the address of an net.IP variable that stores the value of the flag.

#### (*FlagSet) IPMask 

```
func (f *FlagSet) IPMask(name string, value net.IPMask, usage string) *net.IPMask
```

IPMask defines an net.IPMask flag with specified name, default value, and usage string. The return value is the address of an net.IPMask variable that stores the value of the flag.

#### (*FlagSet) IPMaskP 

```
func (f *FlagSet) IPMaskP(name, shorthand string, value net.IPMask, usage string) *net.IPMask
```

IPMaskP is like IPMask, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IPMaskVar 

```
func (f *FlagSet) IPMaskVar(p *net.IPMask, name string, value net.IPMask, usage string)
```

IPMaskVar defines an net.IPMask flag with specified name, default value, and usage string. The argument p points to an net.IPMask variable in which to store the value of the flag.

#### (*FlagSet) IPMaskVarP 

```
func (f *FlagSet) IPMaskVarP(p *net.IPMask, name, shorthand string, value net.IPMask, usage string)
```

IPMaskVarP is like IPMaskVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IPNet 

```
func (f *FlagSet) IPNet(name string, value net.IPNet, usage string) *net.IPNet
```

IPNet defines an net.IPNet flag with specified name, default value, and usage string. The return value is the address of an net.IPNet variable that stores the value of the flag.

#### (*FlagSet) IPNetP 

```
func (f *FlagSet) IPNetP(name, shorthand string, value net.IPNet, usage string) *net.IPNet
```

IPNetP is like IPNet, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IPNetVar 

```
func (f *FlagSet) IPNetVar(p *net.IPNet, name string, value net.IPNet, usage string)
```

IPNetVar defines an net.IPNet flag with specified name, default value, and usage string. The argument p points to an net.IPNet variable in which to store the value of the flag.

#### (*FlagSet) IPNetVarP 

```
func (f *FlagSet) IPNetVarP(p *net.IPNet, name, shorthand string, value net.IPNet, usage string)
```

IPNetVarP is like IPNetVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IPP 

```
func (f *FlagSet) IPP(name, shorthand string, value net.IP, usage string) *net.IP
```

IPP is like IP, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IPSlice 

```
func (f *FlagSet) IPSlice(name string, value []net.IP, usage string) *[]net.IP
```

IPSlice defines a []net.IP flag with specified name, default value, and usage string. The return value is the address of a []net.IP variable that stores the value of that flag.

#### (*FlagSet) IPSliceP 

```
func (f *FlagSet) IPSliceP(name, shorthand string, value []net.IP, usage string) *[]net.IP
```

IPSliceP is like IPSlice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IPSliceVar 

```
func (f *FlagSet) IPSliceVar(p *[]net.IP, name string, value []net.IP, usage string)
```

IPSliceVar defines a ipSlice flag with specified name, default value, and usage string. The argument p points to a []net.IP variable in which to store the value of the flag.

#### (*FlagSet) IPSliceVarP 

```
func (f *FlagSet) IPSliceVarP(p *[]net.IP, name, shorthand string, value []net.IP, usage string)
```

IPSliceVarP is like IPSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IPVar 

```
func (f *FlagSet) IPVar(p *net.IP, name string, value net.IP, usage string)
```

IPVar defines an net.IP flag with specified name, default value, and usage string. The argument p points to an net.IP variable in which to store the value of the flag.

#### (*FlagSet) IPVarP 

```
func (f *FlagSet) IPVarP(p *net.IP, name, shorthand string, value net.IP, usage string)
```

IPVarP is like IPVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Init 

```
func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
```

Init sets the name and error handling property for a flag set. By default, the zero FlagSet uses an empty name and the ContinueOnError error handling policy.

#### (*FlagSet) Int 

```
func (f *FlagSet) Int(name string, value int, usage string) *int
```

Int defines an int flag with specified name, default value, and usage string. The return value is the address of an int variable that stores the value of the flag.

#### (*FlagSet) Int16  <- v1.0.1

```
func (f *FlagSet) Int16(name string, value int16, usage string) *int16
```

Int16 defines an int16 flag with specified name, default value, and usage string. The return value is the address of an int16 variable that stores the value of the flag.

#### (*FlagSet) Int16P  <- v1.0.1

```
func (f *FlagSet) Int16P(name, shorthand string, value int16, usage string) *int16
```

Int16P is like Int16, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int16Var  <- v1.0.1

```
func (f *FlagSet) Int16Var(p *int16, name string, value int16, usage string)
```

Int16Var defines an int16 flag with specified name, default value, and usage string. The argument p points to an int16 variable in which to store the value of the flag.

#### (*FlagSet) Int16VarP  <- v1.0.1

```
func (f *FlagSet) Int16VarP(p *int16, name, shorthand string, value int16, usage string)
```

Int16VarP is like Int16Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int32 

```
func (f *FlagSet) Int32(name string, value int32, usage string) *int32
```

Int32 defines an int32 flag with specified name, default value, and usage string. The return value is the address of an int32 variable that stores the value of the flag.

#### (*FlagSet) Int32P 

```
func (f *FlagSet) Int32P(name, shorthand string, value int32, usage string) *int32
```

Int32P is like Int32, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int32Slice  <- v1.0.5

```
func (f *FlagSet) Int32Slice(name string, value []int32, usage string) *[]int32
```

Int32Slice defines a []int32 flag with specified name, default value, and usage string. The return value is the address of a []int32 variable that stores the value of the flag.

#### (*FlagSet) Int32SliceP  <- v1.0.5

```
func (f *FlagSet) Int32SliceP(name, shorthand string, value []int32, usage string) *[]int32
```

Int32SliceP is like Int32Slice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int32SliceVar  <- v1.0.5

```
func (f *FlagSet) Int32SliceVar(p *[]int32, name string, value []int32, usage string)
```

Int32SliceVar defines a int32Slice flag with specified name, default value, and usage string. The argument p points to a []int32 variable in which to store the value of the flag.

#### (*FlagSet) Int32SliceVarP  <- v1.0.5

```
func (f *FlagSet) Int32SliceVarP(p *[]int32, name, shorthand string, value []int32, usage string)
```

Int32SliceVarP is like Int32SliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int32Var 

```
func (f *FlagSet) Int32Var(p *int32, name string, value int32, usage string)
```

Int32Var defines an int32 flag with specified name, default value, and usage string. The argument p points to an int32 variable in which to store the value of the flag.

#### (*FlagSet) Int32VarP 

```
func (f *FlagSet) Int32VarP(p *int32, name, shorthand string, value int32, usage string)
```

Int32VarP is like Int32Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int64 

```
func (f *FlagSet) Int64(name string, value int64, usage string) *int64
```

Int64 defines an int64 flag with specified name, default value, and usage string. The return value is the address of an int64 variable that stores the value of the flag.

#### (*FlagSet) Int64P 

```
func (f *FlagSet) Int64P(name, shorthand string, value int64, usage string) *int64
```

Int64P is like Int64, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int64Slice  <- v1.0.5

```
func (f *FlagSet) Int64Slice(name string, value []int64, usage string) *[]int64
```

Int64Slice defines a []int64 flag with specified name, default value, and usage string. The return value is the address of a []int64 variable that stores the value of the flag.

#### (*FlagSet) Int64SliceP  <- v1.0.5

```
func (f *FlagSet) Int64SliceP(name, shorthand string, value []int64, usage string) *[]int64
```

Int64SliceP is like Int64Slice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int64SliceVar  <- v1.0.5

```
func (f *FlagSet) Int64SliceVar(p *[]int64, name string, value []int64, usage string)
```

Int64SliceVar defines a int64Slice flag with specified name, default value, and usage string. The argument p points to a []int64 variable in which to store the value of the flag.

#### (*FlagSet) Int64SliceVarP  <- v1.0.5

```
func (f *FlagSet) Int64SliceVarP(p *[]int64, name, shorthand string, value []int64, usage string)
```

Int64SliceVarP is like Int64SliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int64Var 

```
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var defines an int64 flag with specified name, default value, and usage string. The argument p points to an int64 variable in which to store the value of the flag.

#### (*FlagSet) Int64VarP 

```
func (f *FlagSet) Int64VarP(p *int64, name, shorthand string, value int64, usage string)
```

Int64VarP is like Int64Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int8 

```
func (f *FlagSet) Int8(name string, value int8, usage string) *int8
```

Int8 defines an int8 flag with specified name, default value, and usage string. The return value is the address of an int8 variable that stores the value of the flag.

#### (*FlagSet) Int8P 

```
func (f *FlagSet) Int8P(name, shorthand string, value int8, usage string) *int8
```

Int8P is like Int8, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Int8Var 

```
func (f *FlagSet) Int8Var(p *int8, name string, value int8, usage string)
```

Int8Var defines an int8 flag with specified name, default value, and usage string. The argument p points to an int8 variable in which to store the value of the flag.

#### (*FlagSet) Int8VarP 

```
func (f *FlagSet) Int8VarP(p *int8, name, shorthand string, value int8, usage string)
```

Int8VarP is like Int8Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IntP 

```
func (f *FlagSet) IntP(name, shorthand string, value int, usage string) *int
```

IntP is like Int, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IntSlice 

```
func (f *FlagSet) IntSlice(name string, value []int, usage string) *[]int
```

IntSlice defines a []int flag with specified name, default value, and usage string. The return value is the address of a []int variable that stores the value of the flag.

#### (*FlagSet) IntSliceP 

```
func (f *FlagSet) IntSliceP(name, shorthand string, value []int, usage string) *[]int
```

IntSliceP is like IntSlice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IntSliceVar 

```
func (f *FlagSet) IntSliceVar(p *[]int, name string, value []int, usage string)
```

IntSliceVar defines a intSlice flag with specified name, default value, and usage string. The argument p points to a []int variable in which to store the value of the flag.

#### (*FlagSet) IntSliceVarP 

```
func (f *FlagSet) IntSliceVarP(p *[]int, name, shorthand string, value []int, usage string)
```

IntSliceVarP is like IntSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) IntVar 

```
func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
```

IntVar defines an int flag with specified name, default value, and usage string. The argument p points to an int variable in which to store the value of the flag.

#### (*FlagSet) IntVarP 

```
func (f *FlagSet) IntVarP(p *int, name, shorthand string, value int, usage string)
```

IntVarP is like IntVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Lookup 

```
func (f *FlagSet) Lookup(name string) *Flag
```

Lookup returns the Flag structure of the named flag, returning nil if none exists.

#### (*FlagSet) MarkDeprecated 

```
func (f *FlagSet) MarkDeprecated(name string, usageMessage string) error
```

MarkDeprecated indicated that a flag is deprecated in your program. It will continue to function but will not show up in help or usage messages. Using this flag will also print the given usageMessage.

#### (*FlagSet) MarkHidden 

```
func (f *FlagSet) MarkHidden(name string) error
```

MarkHidden sets a flag to 'hidden' in your program. It will continue to function but will not show up in help or usage messages.

#### (*FlagSet) MarkShorthandDeprecated 

```
func (f *FlagSet) MarkShorthandDeprecated(name string, usageMessage string) error
```

MarkShorthandDeprecated will mark the shorthand of a flag deprecated in your program. It will continue to function but will not show up in help or usage messages. Using this flag will also print the given usageMessage.

#### (*FlagSet) NArg 

```
func (f *FlagSet) NArg() int
```

NArg is the number of arguments remaining after flags have been processed.

#### (*FlagSet) NFlag 

```
func (f *FlagSet) NFlag() int
```

NFlag returns the number of flags that have been set.

#### (*FlagSet) Parse 

```
func (f *FlagSet) Parse(arguments []string) error
```

Parse parses flag definitions from the argument list, which should not include the command name. Must be called after all flags in the FlagSet are defined and before flags are accessed by the program. The return value will be ErrHelp if -help was set but not defined.

#### (*FlagSet) ParseAll 

```
func (f *FlagSet) ParseAll(arguments []string, fn func(flag *Flag, value string) error) error
```

ParseAll parses flag definitions from the argument list, which should not include the command name. The arguments for fn are flag and value. Must be called after all flags in the FlagSet are defined and before flags are accessed by the program. The return value will be ErrHelp if -help was set but not defined.

#### (*FlagSet) Parsed 

```
func (f *FlagSet) Parsed() bool
```

Parsed reports whether f.Parse has been called.

#### (*FlagSet) PrintDefaults 

```
func (f *FlagSet) PrintDefaults()
```

PrintDefaults prints, to standard error unless configured otherwise, the default values of all defined flags in the set.

#### (*FlagSet) Set 

```
func (f *FlagSet) Set(name, value string) error
```

Set sets the value of the named flag.

#### (*FlagSet) SetAnnotation 

```
func (f *FlagSet) SetAnnotation(name, key string, values []string) error
```

SetAnnotation allows one to set arbitrary annotations on a flag in the FlagSet. This is sometimes used by spf13/cobra programs which want to generate additional bash completion information.

#### (*FlagSet) SetInterspersed 

```
func (f *FlagSet) SetInterspersed(interspersed bool)
```

SetInterspersed sets whether to support interspersed option/non-option arguments.

#### (*FlagSet) SetNormalizeFunc 

```
func (f *FlagSet) SetNormalizeFunc(n func(f *FlagSet, name string) NormalizedName)
```

SetNormalizeFunc allows you to add a function which can translate flag names. Flags added to the FlagSet will be translated and then when anything tries to look up the flag that will also be translated. So it would be possible to create a flag named "getURL" and have it translated to "geturl". A user could then pass "--getUrl" which may also be translated to "geturl" and everything will work.

#### (*FlagSet) SetOutput 

```
func (f *FlagSet) SetOutput(output io.Writer)
```

SetOutput sets the destination for usage and error messages. If output is nil, os.Stderr is used.

#### (*FlagSet) ShorthandLookup 

```
func (f *FlagSet) ShorthandLookup(name string) *Flag
```

ShorthandLookup returns the Flag structure of the short handed flag, returning nil if none exists. It panics, if len(name) > 1.

##### Example

#### (*FlagSet) String 

```
func (f *FlagSet) String(name string, value string, usage string) *string
```

String defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag.

#### (*FlagSet) StringArray 

```
func (f *FlagSet) StringArray(name string, value []string, usage string) *[]string
```

StringArray defines a string flag with specified name, default value, and usage string. The return value is the address of a []string variable that stores the value of the flag. The value of each argument will not try to be separated by comma. Use a StringSlice for that.

#### (*FlagSet) StringArrayP 

```
func (f *FlagSet) StringArrayP(name, shorthand string, value []string, usage string) *[]string
```

StringArrayP is like StringArray, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringArrayVar 

```
func (f *FlagSet) StringArrayVar(p *[]string, name string, value []string, usage string)
```

StringArrayVar defines a string flag with specified name, default value, and usage string. The argument p points to a []string variable in which to store the values of the multiple flags. The value of each argument will not try to be separated by comma. Use a StringSlice for that.

#### (*FlagSet) StringArrayVarP 

```
func (f *FlagSet) StringArrayVarP(p *[]string, name, shorthand string, value []string, usage string)
```

StringArrayVarP is like StringArrayVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringP 

```
func (f *FlagSet) StringP(name, shorthand string, value string, usage string) *string
```

StringP is like String, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringSlice 

```
func (f *FlagSet) StringSlice(name string, value []string, usage string) *[]string
```

StringSlice defines a string flag with specified name, default value, and usage string. The return value is the address of a []string variable that stores the value of the flag. Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly. For example:

```
--ss="v1,v2" --ss="v3"
```

will result in

```
[]string{"v1", "v2", "v3"}
```

#### (*FlagSet) StringSliceP 

```
func (f *FlagSet) StringSliceP(name, shorthand string, value []string, usage string) *[]string
```

StringSliceP is like StringSlice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringSliceVar 

```
func (f *FlagSet) StringSliceVar(p *[]string, name string, value []string, usage string)
```

StringSliceVar defines a string flag with specified name, default value, and usage string. The argument p points to a []string variable in which to store the value of the flag. Compared to StringArray flags, StringSlice flags take comma-separated value as arguments and split them accordingly. For example:

```
--ss="v1,v2" --ss="v3"
```

will result in

```
[]string{"v1", "v2", "v3"}
```

#### (*FlagSet) StringSliceVarP 

```
func (f *FlagSet) StringSliceVarP(p *[]string, name, shorthand string, value []string, usage string)
```

StringSliceVarP is like StringSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringToInt  <- v1.0.3

```
func (f *FlagSet) StringToInt(name string, value map[string]int, usage string) *map[string]int
```

StringToInt defines a string flag with specified name, default value, and usage string. The return value is the address of a map[string]int variable that stores the value of the flag. The value of each argument will not try to be separated by comma

#### (*FlagSet) StringToInt64  <- v1.0.5

```
func (f *FlagSet) StringToInt64(name string, value map[string]int64, usage string) *map[string]int64
```

StringToInt64 defines a string flag with specified name, default value, and usage string. The return value is the address of a map[string]int64 variable that stores the value of the flag. The value of each argument will not try to be separated by comma

#### (*FlagSet) StringToInt64P  <- v1.0.5

```
func (f *FlagSet) StringToInt64P(name, shorthand string, value map[string]int64, usage string) *map[string]int64
```

StringToInt64P is like StringToInt64, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringToInt64Var  <- v1.0.5

```
func (f *FlagSet) StringToInt64Var(p *map[string]int64, name string, value map[string]int64, usage string)
```

StringToInt64Var defines a string flag with specified name, default value, and usage string. The argument p point64s to a map[string]int64 variable in which to store the values of the multiple flags. The value of each argument will not try to be separated by comma

#### (*FlagSet) StringToInt64VarP  <- v1.0.5

```
func (f *FlagSet) StringToInt64VarP(p *map[string]int64, name, shorthand string, value map[string]int64, usage string)
```

StringToInt64VarP is like StringToInt64Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringToIntP  <- v1.0.3

```
func (f *FlagSet) StringToIntP(name, shorthand string, value map[string]int, usage string) *map[string]int
```

StringToIntP is like StringToInt, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringToIntVar  <- v1.0.3

```
func (f *FlagSet) StringToIntVar(p *map[string]int, name string, value map[string]int, usage string)
```

StringToIntVar defines a string flag with specified name, default value, and usage string. The argument p points to a map[string]int variable in which to store the values of the multiple flags. The value of each argument will not try to be separated by comma

#### (*FlagSet) StringToIntVarP  <- v1.0.3

```
func (f *FlagSet) StringToIntVarP(p *map[string]int, name, shorthand string, value map[string]int, usage string)
```

StringToIntVarP is like StringToIntVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringToString  <- v1.0.3

```
func (f *FlagSet) StringToString(name string, value map[string]string, usage string) *map[string]string
```

StringToString defines a string flag with specified name, default value, and usage string. The return value is the address of a map[string]string variable that stores the value of the flag. The value of each argument will not try to be separated by comma

#### (*FlagSet) StringToStringP  <- v1.0.3

```
func (f *FlagSet) StringToStringP(name, shorthand string, value map[string]string, usage string) *map[string]string
```

StringToStringP is like StringToString, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringToStringVar  <- v1.0.3

```
func (f *FlagSet) StringToStringVar(p *map[string]string, name string, value map[string]string, usage string)
```

StringToStringVar defines a string flag with specified name, default value, and usage string. The argument p points to a map[string]string variable in which to store the values of the multiple flags. The value of each argument will not try to be separated by comma

#### (*FlagSet) StringToStringVarP  <- v1.0.3

```
func (f *FlagSet) StringToStringVarP(p *map[string]string, name, shorthand string, value map[string]string, usage string)
```

StringToStringVarP is like StringToStringVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) StringVar 

```
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
```

StringVar defines a string flag with specified name, default value, and usage string. The argument p points to a string variable in which to store the value of the flag.

#### (*FlagSet) StringVarP 

```
func (f *FlagSet) StringVarP(p *string, name, shorthand string, value string, usage string)
```

StringVarP is like StringVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Uint 

```
func (f *FlagSet) Uint(name string, value uint, usage string) *uint
```

Uint defines a uint flag with specified name, default value, and usage string. The return value is the address of a uint variable that stores the value of the flag.

#### (*FlagSet) Uint16 

```
func (f *FlagSet) Uint16(name string, value uint16, usage string) *uint16
```

Uint16 defines a uint flag with specified name, default value, and usage string. The return value is the address of a uint variable that stores the value of the flag.

#### (*FlagSet) Uint16P 

```
func (f *FlagSet) Uint16P(name, shorthand string, value uint16, usage string) *uint16
```

Uint16P is like Uint16, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Uint16Var 

```
func (f *FlagSet) Uint16Var(p *uint16, name string, value uint16, usage string)
```

Uint16Var defines a uint flag with specified name, default value, and usage string. The argument p points to a uint variable in which to store the value of the flag.

#### (*FlagSet) Uint16VarP 

```
func (f *FlagSet) Uint16VarP(p *uint16, name, shorthand string, value uint16, usage string)
```

Uint16VarP is like Uint16Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Uint32 

```
func (f *FlagSet) Uint32(name string, value uint32, usage string) *uint32
```

Uint32 defines a uint32 flag with specified name, default value, and usage string. The return value is the address of a uint32 variable that stores the value of the flag.

#### (*FlagSet) Uint32P 

```
func (f *FlagSet) Uint32P(name, shorthand string, value uint32, usage string) *uint32
```

Uint32P is like Uint32, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Uint32Var 

```
func (f *FlagSet) Uint32Var(p *uint32, name string, value uint32, usage string)
```

Uint32Var defines a uint32 flag with specified name, default value, and usage string. The argument p points to a uint32 variable in which to store the value of the flag.

#### (*FlagSet) Uint32VarP 

```
func (f *FlagSet) Uint32VarP(p *uint32, name, shorthand string, value uint32, usage string)
```

Uint32VarP is like Uint32Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Uint64 

```
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
```

Uint64 defines a uint64 flag with specified name, default value, and usage string. The return value is the address of a uint64 variable that stores the value of the flag.

#### (*FlagSet) Uint64P 

```
func (f *FlagSet) Uint64P(name, shorthand string, value uint64, usage string) *uint64
```

Uint64P is like Uint64, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Uint64Var 

```
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
```

Uint64Var defines a uint64 flag with specified name, default value, and usage string. The argument p points to a uint64 variable in which to store the value of the flag.

#### (*FlagSet) Uint64VarP 

```
func (f *FlagSet) Uint64VarP(p *uint64, name, shorthand string, value uint64, usage string)
```

Uint64VarP is like Uint64Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Uint8 

```
func (f *FlagSet) Uint8(name string, value uint8, usage string) *uint8
```

Uint8 defines a uint8 flag with specified name, default value, and usage string. The return value is the address of a uint8 variable that stores the value of the flag.

#### (*FlagSet) Uint8P 

```
func (f *FlagSet) Uint8P(name, shorthand string, value uint8, usage string) *uint8
```

Uint8P is like Uint8, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Uint8Var 

```
func (f *FlagSet) Uint8Var(p *uint8, name string, value uint8, usage string)
```

Uint8Var defines a uint8 flag with specified name, default value, and usage string. The argument p points to a uint8 variable in which to store the value of the flag.

#### (*FlagSet) Uint8VarP 

```
func (f *FlagSet) Uint8VarP(p *uint8, name, shorthand string, value uint8, usage string)
```

Uint8VarP is like Uint8Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) UintP 

```
func (f *FlagSet) UintP(name, shorthand string, value uint, usage string) *uint
```

UintP is like Uint, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) UintSlice 

```
func (f *FlagSet) UintSlice(name string, value []uint, usage string) *[]uint
```

UintSlice defines a []uint flag with specified name, default value, and usage string. The return value is the address of a []uint variable that stores the value of the flag.

#### (*FlagSet) UintSliceP 

```
func (f *FlagSet) UintSliceP(name, shorthand string, value []uint, usage string) *[]uint
```

UintSliceP is like UintSlice, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) UintSliceVar 

```
func (f *FlagSet) UintSliceVar(p *[]uint, name string, value []uint, usage string)
```

UintSliceVar defines a uintSlice flag with specified name, default value, and usage string. The argument p points to a []uint variable in which to store the value of the flag.

#### (*FlagSet) UintSliceVarP 

```
func (f *FlagSet) UintSliceVarP(p *[]uint, name, shorthand string, value []uint, usage string)
```

UintSliceVarP is like UintSliceVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) UintVar 

```
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
```

UintVar defines a uint flag with specified name, default value, and usage string. The argument p points to a uint variable in which to store the value of the flag.

#### (*FlagSet) UintVarP 

```
func (f *FlagSet) UintVarP(p *uint, name, shorthand string, value uint, usage string)
```

UintVarP is like UintVar, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) Var 

```
func (f *FlagSet) Var(value Value, name string, usage string)
```

Var defines a flag with the specified name and usage string. The type and value of the flag are represented by the first argument, of type Value, which typically holds a user-defined implementation of Value. For instance, the caller could create a flag that turns a comma-separated string into a slice of strings by giving the slice the methods of Value; in particular, Set would decompose the comma-separated string into the slice.

#### (*FlagSet) VarP 

```
func (f *FlagSet) VarP(value Value, name, shorthand, usage string)
```

VarP is like Var, but accepts a shorthand letter that can be used after a single dash.

#### (*FlagSet) VarPF 

```
func (f *FlagSet) VarPF(value Value, name, shorthand, usage string) *Flag
```

VarPF is like VarP, but returns the flag created

#### (*FlagSet) Visit 

```
func (f *FlagSet) Visit(fn func(*Flag))
```

Visit visits the flags in lexicographical order or in primordial order if f.SortFlags is false, calling fn for each. It visits only those flags that have been set.

#### (*FlagSet) VisitAll 

```
func (f *FlagSet) VisitAll(fn func(*Flag))
```

VisitAll visits the flags in lexicographical order or in primordial order if f.SortFlags is false, calling fn for each. It visits all flags, even those not set.

#### type NormalizedName 

```
type NormalizedName string
```

NormalizedName is a flag name that has been normalized according to rules for the FlagSet (e.g. making '-' and '_' equivalent).

#### type ParseErrorsWhitelist  <- v1.0.1

```
type ParseErrorsWhitelist struct {
	// UnknownFlags will ignore unknown flags errors and continue parsing rest of the flags
	UnknownFlags bool
}
```

ParseErrorsWhitelist defines the parsing errors that can be ignored

#### type SliceValue  <- v1.0.5

```
type SliceValue interface {
	// Append adds the specified value to the end of the flag value list.
	Append(string) error
	// Replace will fully overwrite any data currently in the flag value list.
	Replace([]string) error
	// GetSlice returns the flag value list as an array of strings.
	GetSlice() []string
}
```

SliceValue is a secondary interface to all flags which hold a list of values. This allows full control over the value of list flags, and avoids complicated marshalling and unmarshalling to csv.

#### type Value 

```
type Value interface {
	String() string
	Set(string) error
	Type() string
}
```

Value is the interface to the dynamic value stored in a flag. (The default value is represented as a string.)