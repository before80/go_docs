+++
title = "flag"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/flag@go1.23.0](https://pkg.go.dev/flag@go1.23.0)

Package flag implements command-line flag parsing.

​	`flag`包实现了命令行标志解析。

## Usage 

Define flags using flag.String(), Bool(), Int(), etc.

​	使用 flag.String()、Bool()、Int() 等函数定义标志。

This declares an integer flag, -n, stored in the pointer nFlag, with type *int:

​	下面的代码声明了一个整数标志 -n，存储在指针 nFlag 中，类型为 `*int`：

```go 
import "flag"
var nFlag = flag.Int("n", 1234, "help message for flag n")
```

If you like, you can bind the flag to a variable using the Var() functions.

​	如果您喜欢，可以使用 Var() 函数将标志绑定到变量。

```go 
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

Or you can create custom flags that satisfy the Value interface (with pointer receivers) and couple them to flag parsing by

或者，您可以创建满足 Value 接口(使用指针接收器)的自定义标志，并将它们与标志解析耦合：

```
flag.Var(&flagVal, "name", "help message for flagname")
```

For such flags, the default value is just the initial value of the variable.

​	对于这种标志，默认值就是变量的初始值。

After all flags are defined, call

​	在定义好所有标志之后，调用

```
flag.Parse()
```

to parse the command line into the defined flags.

将命令行参数解析为定义好的标志。

Flags may then be used directly. If you're using the flags themselves, they are all pointers; if you bind to variables, they're values.

​	然后可以直接使用标志。**如果使用标志本身，它们都是指针**；**如果与变量绑定，它们就是值**。

```
fmt.Println("ip has value ", *ip)
fmt.Println("flagvar has value ", flagvar)
```

After parsing, the arguments following the flags are available as the slice flag.Args() or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1.

​	解析后，跟在标志后面的参数可以作为 `flag.Args()` 切片使用，也可以单独作为 flag.Arg(i) 使用。参数的索引从 0 到 `flag.NArg()-1`。

## Command line flag syntax 

The following forms are permitted:

​	以下形式是允许的：

```
-flag
--flag   // 也允许双短线
-flag=x
-flag x  // 只适用于非布尔标志
```

One or two dashes may be used; they are equivalent. The last form is not permitted for boolean flags because the meaning of the command

​	可以使用一个或两个短线，它们是等价的。最后一种形式不适用于布尔标志，因为命令

```
cmd -x *
```

where * is a Unix shell wildcard, will change if there is a file called 0, false, etc. You must use the -flag=false form to turn off a boolean flag.

其中 `*` 是 Unix shell 通配符，如果存在名为 `0`、`false` 等的文件，其含义将发生变化。您必须使用 -flag=false 形式来关闭布尔标志。

Flag parsing stops just before the first non-flag argument ("-" is a non-flag argument) or after the terminator "--".

​	标志解析在第一个非标志参数("-"是非标志参数)之前或终止符"--"之后停止。

Integer flags accept 1234, 0664, 0x1234 and may be negative. Boolean flags may be:

​	整数标志接受 1234、0664、0x1234 等值，也可以是负数。布尔标志可以是：

```
1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
```

Duration flags accept any input valid for time.ParseDuration.

​	时间间隔标志接受任何符合 time.ParseDuration 的输入。

The default set of command-line flags is controlled by top-level functions. The FlagSet type allows one to define independent sets of flags, such as to implement subcommands in a command-line interface. The methods of FlagSet are analogous to the top-level functions for the command-line flag set.

​	默认的命令行标志集由顶级函数控制。FlagSet 类型允许定义独立的标志集，例如实现命令行接口中的子命令。FlagSet 的方法类似于命令行标志集的顶级函数。

### Example

``` go 
// These examples demonstrate more intricate uses of the flag package.
//这些例子展示了flag包的更复杂的用途。
package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// Example 1: A single string flag called "species" with default value "gopher".
// 示例1：一个名为 "species "的单一字符串标志，默认值为 "gopher"。
var species = flag.String("species", "gopher", "the species we are studying")

// Example 2: Two flags sharing a variable, so we can have a shorthand.
// The order of initialization is undefined, so make sure both use the
// same default value. They must be set up with an init function.
// 例二：两个flag共享一个变量，所以我们可以有一个速记的方法。
// 初始化的顺序是未定义的，所以要确保两者使用相同的默认值。
// 它们必须用一个init函数来设置。
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

// Example 3: A user-defined flag type, a slice of durations.
// 例3：一个用户定义的标志类型，一个持续时间的切片。
type interval []time.Duration

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
// String是用来格式化标志值的方法，是flag.Value接口的一部分。
// String方法的输出将在诊断中使用。
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
// Set是设置标志值的方法，是flag.Value接口的一部分。
// Set的参数是一个要解析的字符串，用来设置标志。
// 它是一个逗号分隔的列表，所以我们把它拆开。
func (i *interval) Set(value string) error {
    // If we wanted to allow the flag to be set multiple times,
	// accumulating values, we would delete this if statement.
	// That would permit usages such as
	//	-deltaT 10s -deltaT 15s
	// and other combinations.
    // 如果我们想让这个标志被设置多次，累积数值，我们可以删除这个if语句。
	// 这将允许诸如-deltaT 10s -deltaT 15s和其他组合的使用。
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// Define a flag to accumulate durations. Because it has a special type,
// we need to use the Var function and therefore create the flag during
// init.
// 定义一个标志来累积持续时间。因为它有一个特殊的类型。
// 我们需要使用Var函数，因此在init期间创建该标志。

var intervalFlag interval

func init() {
    // Tie the command-line flag to the intervalFlag variable and
	// set a usage message.
    // 将命令行标志与intervalFlag变量绑定，并设置一个使用信息。
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
    // All the interesting pieces are with the variables declared above, but
	// to enable the flag package to see the flags defined there, one must
	// execute, typically at the start of main (not init!):
    // 所有有趣的部分都与上面声明的变量在一起，
    // 但为了使flag包能够看到那里定义的标志，必须执行，通常在main的开始(不是init！):
	// flag.Parse()
    // We don't call it here because this code is a function called "Example"
	// that is part of the testing suite for the package, which has already
	// parsed the flags. When viewed at pkg.go.dev, however, the function is
	// renamed to "main" and it could be run as a standalone example.
	// 我们不在这里调用它，因为这段代码是一个叫做 "Example"的函数，是包测试套件的一部分，
    // 它已经解析了标志。但在pkg.go.dev查看时，该函数被重命名为 "main"，
    // 它可以作为一个独立的例子运行。
}

```


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/flag/flag.go;l=1150)

``` go 
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

CommandLine is the default set of command-line flags, parsed from os.Args. The top-level functions such as BoolVar, Arg, and so on are wrappers for the methods of CommandLine.

​	CommandLine 是默认的命令行标志集，从 os.Args 解析而来。诸如 BoolVar、Arg 等顶级函数是 CommandLine 方法的包装器。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/flag/flag.go;l=100)

``` go 
var ErrHelp = errors.New("flag: help requested")
```

ErrHelp is the error returned if the -help or -h flag is invoked but no such flag is defined.

​	ErrHelp 是在调用 -help 或 -h 标志，但没有定义这样的标志时返回的错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/flag/flag.go;l=674)

``` go 
var Usage = func() {
	fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	PrintDefaults()
}
```

Usage prints a usage message documenting all defined command-line flags to CommandLine's output, which by default is os.Stderr. It is called when an error occurs while parsing flags. The function is a variable that may be changed to point to a custom function. By default it prints a simple header and calls PrintDefaults; for details about the format of the output and how to control it, see the documentation for PrintDefaults. Custom usage functions may choose to exit the program; by default exiting happens anyway as the command line's error handling strategy is set to ExitOnError.

​	Usage 打印一个使用说明，文档化所有已定义的命令行标志到 CommandLine 的输出，这个输出默认是 os.Stderr。当解析标志时发生错误时会调用它。该函数是一个变量，可以更改指向自定义函数。默认情况下，它会打印一个简单的标题并调用 [PrintDefaults](#func-printdefaults)；关于输出格式及如何控制它的详细信息，请参阅 PrintDefaults 的文档。自定义的使用函数可能选择退出程序；默认情况下，退出仍会发生，因为命令行的错误处理策略设置为 `ExitOnError`。

## 函数

### func Arg 

``` go 
func Arg(i int) string
```

Arg returns the i'th command-line argument. Arg(0) is the first remaining argument after flags have been processed. Arg returns an empty string if the requested element does not exist.

​	Arg函数返回第 i 个命令行参数。Arg(0) 是在处理标志后剩下的第一个参数。如果请求的元素不存在，则 Arg 返回空字符串。

### func Args 

``` go 
func Args() []string
```

Args returns the non-flag command-line arguments.

​	Args函数返回非标志命令行参数。

### func Bool 

``` go 
func Bool(name string, value bool, usage string) *bool
```

Bool defines a bool flag with specified name, default value, and usage string. The return value is the address of a bool variable that stores the value of the flag.

​	Bool函数定义一个具有指定名称、默认值和用法字符串的布尔标志。返回值是存储标志值的布尔变量的地址。

### func BoolFunc <- go1.21.0

```go
func BoolFunc(name, usage string, fn func(string) error)
```

BoolFunc defines a flag with the specified name and usage string without requiring values. Each time the flag is seen, fn is called with the value of the flag. If fn returns a non-nil error, it will be treated as a flag value parsing error.

​	`BoolFunc` 定义了一个具有指定名称和使用说明的标志，无需提供值。每次看到该标志时，都会调用 `fn` 并传递该标志的值。如果 `fn` 返回非空错误，则会将其视为标志值解析错误。

#### BoolFunc Example

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fs := flag.NewFlagSet("ExampleBoolFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	fs.BoolFunc("log", "logs a dummy message", func(s string) error {
		fmt.Println("dummy message:", s)
		return nil
	})
	fs.Parse([]string{"-log"})
	fs.Parse([]string{"-log=0"})

}
Output:

dummy message: true
dummy message: 0
```



### func BoolVar 

``` go 
func BoolVar(p *bool, name string, value bool, usage string)
```

BoolVar defines a bool flag with specified name, default value, and usage string. The argument p points to a bool variable in which to store the value of the flag.

​	BoolVar函数定义一个具有指定名称、默认值和用法字符串的布尔标志。参数 p 指向一个布尔变量，用于存储标志的值。

### func Duration 

``` go 
func Duration(name string, value time.Duration, usage string) *time.Duration
```

Duration defines a time.Duration flag with specified name, default value, and usage string. The return value is the address of a time.Duration variable that stores the value of the flag. The flag accepts a value acceptable to time.ParseDuration.

​	Duration函数定义一个具有指定名称、默认值和用法字符串的 time.Duration 标志。返回值是存储标志值的 time.Duration 变量的地址。标志接受符合 time.ParseDuration 的值。

### func DurationVar 

``` go 
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

DurationVar defines a time.Duration flag with specified name, default value, and usage string. The argument p points to a time.Duration variable in which to store the value of the flag. The flag accepts a value acceptable to time.ParseDuration.

​	DurationVar函数定义一个具有指定名称、默认值和用法字符串的 time.Duration 标志。参数 p 指向一个 time.Duration 变量，用于存储标志的值。标志接受符合 time.ParseDuration 的值。

### func Float64 

``` go 
func Float64(name string, value float64, usage string) *float64
```

Float64 defines a float64 flag with specified name, default value, and usage string. The return value is the address of a float64 variable that stores the value of the flag.

​	Float64函数定义一个具有指定名称、默认值和用法字符串的 float64 标志。返回值是存储标志值的 float64 变量的地址。

### func Float64Var 

``` go 
func Float64Var(p *float64, name string, value float64, usage string)
```

Float64Var defines a float64 flag with specified name, default value, and usage string. The argument p points to a float64 variable in which to store the value of the flag.

​	Float64Var函数定义一个具有指定名称、默认值和用法字符串的 float64 标志。参数 p 指向一个 float64 变量，用于存储标志的值。

### func Func  <- go1.16

``` go 
func Func(name, usage string, fn func(string) error)
```

Func defines a flag with the specified name and usage string. Each time the flag is seen, fn is called with the value of the flag. If fn returns a non-nil error, it will be treated as a flag value parsing error.

​	Func函数定义一个具有指定名称和用法字符串的标志。每次看到标志时，fn 都会使用标志的值调用。如果 fn 返回非 nil 错误，则将其视为标志值解析错误。

##### Func Example
``` go 
package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	fs := flag.NewFlagSet("ExampleFunc", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	fs.Func("ip", "`IP address` to parse", func(s string) error {
		ip = net.ParseIP(s)
		if ip == nil {
			return errors.New("could not parse IP")
		}
		return nil
	})
	fs.Parse([]string{"-ip", "127.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

    // 256 is not a valid IPv4 component
	// 256不是一个有效的IPv4组成部分
	fs.Parse([]string{"-ip", "256.0.0.1"})
	fmt.Printf("{ip: %v, loopback: %t}\n\n", ip, ip.IsLoopback())

}
Output:

{ip: 127.0.0.1, loopback: true}

invalid value "256.0.0.1" for flag -ip: could not parse IP
Usage of ExampleFunc:
  -ip IP address
    	IP address to parse
{ip: <nil>, loopback: false}
```

### func Int 

``` go 
func Int(name string, value int, usage string) *int
```

Int defines an int flag with specified name, default value, and usage string. The return value is the address of an int variable that stores the value of the flag.

​	Int函数定义一个具有指定名称、默认值和用法字符串的 int 标志。返回值是存储标志值的 int 变量的地址。

### func Int64 

``` go 
func Int64(name string, value int64, usage string) *int64
```

Int64 defines an int64 flag with specified name, default value, and usage string. The return value is the address of an int64 variable that stores the value of the flag.

​	Int64函数定义一个具有指定名称、默认值和用法字符串的 int64 标志。返回值是存储标志值的 int64 变量的地址。

### func Int64Var 

``` go 
func Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var defines an int64 flag with specified name, default value, and usage string. The argument p points to an int64 variable in which to store the value of the flag.

​	Int64Var函数定义一个具有指定名称、默认值和用法字符串的 int64 标志。参数 p 指向一个 int64 变量，用于存储标志的值。

### func IntVar 

``` go 
func IntVar(p *int, name string, value int, usage string)
```

IntVar defines an int flag with specified name, default value, and usage string. The argument p points to an int variable in which to store the value of the flag.

​	IntVar函数定义一个具有指定名称、默认值和用法字符串的 int 标志。参数 p 指向一个 int 变量，用于存储标志的值。

### func NArg 

``` go 
func NArg() int
```

NArg is the number of arguments remaining after flags have been processed.

​	NArg函数返回在处理标志后剩余的参数数目。

### func NFlag 

``` go 
func NFlag() int
```

NFlag returns the number of command-line flags that have been set.

​	NFlag函数返回已设置的命令行标志数目。

### func Parse 

``` go 
func Parse()
```

Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.

​	Parse函数从 `os.Args[1:]` 解析命令行标志。必须在定义所有标志并在程序访问标志之前调用它。

### func Parsed 

``` go 
func Parsed() bool
```

Parsed reports whether the command-line flags have been parsed.

​	Parsed函数报告命令行标志是否已被解析。

### func PrintDefaults 

``` go 
func PrintDefaults()
```

PrintDefaults prints, to standard error unless configured otherwise, a usage message showing the default settings of all defined command-line flags. For an integer valued flag x, the default output has the form

​	PrintDefaults函数打印一个使用说明，显示所有已定义命令行标志的默认设置。对于值为整数的标志 x，默认输出的格式为：

```
-x int
	usage-message-for-x (default 7)
```

The usage message will appear on a separate line for anything but a bool flag with a one-byte name. For bool flags, the type is omitted and if the flag name is one byte the usage message appears on the same line. The parenthetical default is omitted if the default is the zero value for the type. The listed type, here int, can be changed by placing a back-quoted name in the flag's usage string; the first such item in the message is taken to be a parameter name to show in the message and the back quotes are stripped from the message when displayed. For instance, given

​	使用方法会在除了具有一个字节名称的 bool 标志外的其他地方显示在单独的一行上。对于 bool 标志，类型被省略，如果标志名称为一个字节，则使用说明消息出现在同一行上。如果默认值是类型的零值，则括号中的默认值被省略。列出的类型(这里是 int)可以通过在标志的使用说明字符串中放置一个带反引号的名称来更改；消息中的第一个此类项被认为是要在消息中显示的参数名称，并在显示消息时从消息中删除反引号。例如，给定

```
flag.String("I", "", "search `directory` for include files")
```

the output will be

输出将是

```
-I directory
	search directory for include files.
```

To change the destination for flag messages, call CommandLine.SetOutput.

​	要更改标志消息的目标，请调用 CommandLine.SetOutput。

### func Set 

``` go 
func Set(name, value string) error
```

Set sets the value of the named command-line flag.

​	Set函数设置命名的命令行标志的值。

### func String 

``` go 
func String(name string, value string, usage string) *string
```

String defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag.

​	String函数定义一个具有指定名称、默认值和用法字符串的字符串标志。返回值是存储标志值的字符串变量的地址。

### func StringVar 

``` go 
func StringVar(p *string, name string, value string, usage string)
```

StringVar defines a string flag with specified name, default value, and usage string. The argument p points to a string variable in which to store the value of the flag.

​	StringVar函数定义一个具有指定名称、默认值和用法字符串的字符串标志。参数 p 指向一个字符串变量，用于存储标志的值。

### func TextVar  <- go1.19

``` go 
func TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string)
```

TextVar defines a flag with a specified name, default value, and usage string. The argument p must be a pointer to a variable that will hold the value of the flag, and p must implement encoding.TextUnmarshaler. If the flag is used, the flag value will be passed to p's UnmarshalText method. The type of the default value must be the same as the type of p.

​	TextVar函数定义一个具有指定名称、默认值和用法字符串的标志。参数 p 必须是一个指向将保存标志值的变量的指针，并且 p 必须实现 encoding.TextUnmarshaler 接口。如果使用标志，则将标志值传递给 p 的 UnmarshalText 方法。默认值的类型必须与 p 的类型相同。

#### TextVar Example
``` go 
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	fs := flag.NewFlagSet("ExampleTextVar", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	var ip net.IP
	fs.TextVar(&ip, "ip", net.IPv4(192, 168, 0, 100), "`IP address` to parse")
	fs.Parse([]string{"-ip", "127.0.0.1"})
	fmt.Printf("{ip: %v}\n\n", ip)

    // 256 is not a valid IPv4 component
	// 256 is not a valid IPv4 component
	ip = nil
	fs.Parse([]string{"-ip", "256.0.0.1"})
	fmt.Printf("{ip: %v}\n\n", ip)

}
Output:

{ip: 127.0.0.1}

invalid value "256.0.0.1" for flag -ip: invalid IP address: 256.0.0.1
Usage of ExampleTextVar:
  -ip IP address
    	IP address to parse (default 192.168.0.100)
{ip: <nil>}
```

### func Uint 

``` go 
func Uint(name string, value uint, usage string) *uint
```

Uint defines a uint flag with specified name, default value, and usage string. The return value is the address of a uint variable that stores the value of the flag.

​	Uint函数定义一个具有指定名称、默认值和用法字符串的 uint 标志。返回值是存储标志值的 uint 变量的地址。

### func Uint64 

``` go 
func Uint64(name string, value uint64, usage string) *uint64
```

Uint64 defines a uint64 flag with specified name, default value, and usage string. The return value is the address of a uint64 variable that stores the value of the flag.

​	Uint64函数定义一个指定名称、默认值和使用说明的 uint64 标志。返回值是一个 uint64 变量的地址，该变量存储标志的值。

### func Uint64Var 

``` go 
func Uint64Var(p *uint64, name string, value uint64, usage string)
```

Uint64Var defines a uint64 flag with specified name, default value, and usage string. The argument p points to a uint64 variable in which to store the value of the flag.

​	Uint64Var函数定义一个指定名称、默认值和使用说明的 uint64 标志。参数 p 指向一个 uint64 变量，用于存储标志的值。

### func UintVar 

``` go 
func UintVar(p *uint, name string, value uint, usage string)
```

UintVar defines a uint flag with specified name, default value, and usage string. The argument p points to a uint variable in which to store the value of the flag.

​	UintVar函数定义一个指定名称、默认值和使用说明的 uint 标志。参数 p 指向一个 uint 变量，用于存储标志的值。

### func UnquoteUsage  <- go1.5

``` go 
func UnquoteUsage(flag *Flag) (name string, usage string)
```

UnquoteUsage extracts a back-quoted name from the usage string for a flag and returns it and the un-quoted usage. Given "a `name` to show" it returns ("name", "a name to show"). If there are no back quotes, the name is an educated guess of the type of the flag's value, or the empty string if the flag is boolean.

​	UnquoteUsage函数从标志的使用说明中提取带反引号的名称，并返回未引用的名称和使用说明。给定 "a `name` to show"，返回 ("name", "a name to show")。如果没有反引号，则名称是标志值类型的猜测，如果标志是布尔值，则为空字符串。

### func Var 

``` go 
func Var(value Value, name string, usage string)
```

Var defines a flag with the specified name and usage string. The type and value of the flag are represented by the first argument, of type Value, which typically holds a user-defined implementation of Value. For instance, the caller could create a flag that turns a comma-separated string into a slice of strings by giving the slice the methods of Value; in particular, Set would decompose the comma-separated string into the slice.

​	Var函数定义一个具有指定名称和使用说明的标志。标志的类型和值由第一个参数 Value 表示，该参数通常持有 Value 的用户定义实现。例如，调用者可以创建一个标志，将逗号分隔的字符串转换为字符串切片，给定的切片具有 Value 的方法；特别是，Set 会将逗号分隔的字符串分解为切片。

### func Visit 

``` go 
func Visit(fn func(*Flag))
```

Visit visits the command-line flags in lexicographical order, calling fn for each. It visits only those flags that have been set.

​	Visit函数按字典顺序访问命令行标志，并为每个标志调用 fn。它仅访问已设置的标志。

### func VisitAll 

``` go 
func VisitAll(fn func(*Flag))
```

VisitAll visits the command-line flags in lexicographical order, calling fn for each. It visits all flags, even those not set.

​	VisitAll函数按字典顺序访问命令行标志，并为每个标志调用 fn。它访问所有标志，即使未设置。

## 类型

### type ErrorHandling 

``` go 
type ErrorHandling int
```

ErrorHandling defines how FlagSet.Parse behaves if the parse fails.

​	ErrorHandling定义了FlagSet.Parse在解析失败时的行为。

``` go 
const (
	ContinueOnError ErrorHandling = iota // 返回一个描述性错误。Return a descriptive error.
	ExitOnError  // 调用os.Exit(2) 或对于-h/-help 则调用Exit(0)。 Call os.Exit(2) or for -h/-help Exit(0).
	PanicOnError  // 调用一个带描述性错误的panic。 Call panic with a descriptive error.
)
```

These constants cause FlagSet.Parse to behave as described if the parse fails.

​	如果解析失败，这些常量会使FlagSet.Parse按照描述的行为进行操作。

### type Flag 

``` go 
type Flag struct {
	Name     string // 在命令行中出现的名称 name as it appears on command line
	Usage    string // 帮助信息 help message
	Value    Value  // 设定的值 value as set
	DefValue string // 默认值(以文本形式)；用于帮助信息 default value (as text); for usage message
}
```

A Flag represents the state of a flag.

​	Flag表示标志的状态。

#### func Lookup 

``` go 
func Lookup(name string) *Flag
```

Lookup returns the Flag structure of the named command-line flag, returning nil if none exists.

​	Lookup函数返回命名的命令行标志的Flag结构，如果不存在则返回nil。

### type FlagSet 

``` go 
type FlagSet struct {
    // Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler. What happens after Usage is called depends
	// on the ErrorHandling setting; for the command line, this defaults
	// to ExitOnError, which exits the program after calling Usage.
	// Usage是解析标志时发生错误时调用的函数。
	// 该字段是一个函数(不是方法)，可以更改以指向自定义错误处理程序。
	// 在调用Usage后发生的情况取决于ErrorHandling设置；
    // 对于命令行，默认值为ExitOnError，
	// 它在调用Usage后退出程序。
	Usage func()
	// 包含已过滤或未导出的字段
}
```

A FlagSet represents a set of defined flags. The zero value of a FlagSet has no name and has ContinueOnError error handling.

​	FlagSet表示一组已定义的标志。FlagSet的零值没有名称并具有ContinueOnError错误处理。

Flag names must be unique within a FlagSet. An attempt to define a flag whose name is already in use will cause a panic.

​	标志名称必须在FlagSet内唯一。尝试定义其名称已在使用中的标志将导致恐慌。

#### func NewFlagSet 

``` go 
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
```

NewFlagSet returns a new, empty flag set with the specified name and error handling property. If the name is not empty, it will be printed in the default usage message and in error messages.

​	NewFlagSet函数返回一个新的、空的FlagSet，其中包含指定的名称和错误处理属性。如果名称不为空，它将在默认用法消息和错误消息中打印。

#### (*FlagSet) Arg 

``` go 
func (f *FlagSet) Arg(i int) string
```

Arg returns the i'th argument. Arg(0) is the first remaining argument after flags have been processed. Arg returns an empty string if the requested element does not exist.

​	Arg方法返回第i个参数。Arg(0)是在处理标志后剩余的第一个参数。如果请求的元素不存在，则Arg返回空字符串。

#### (*FlagSet) Args 

``` go 
func (f *FlagSet) Args() []string
```

Args returns the non-flag arguments.

​	Args方法返回非标志参数。

#### (*FlagSet) Bool 

``` go 
func (f *FlagSet) Bool(name string, value bool, usage string) *bool
```

Bool defines a bool flag with specified name, default value, and usage string. The return value is the address of a bool variable that stores the value of the flag.

​	Bool方法定义具有指定名称、默认值和用法字符串的布尔标志。返回值是一个bool变量的地址，该变量存储标志的值。

#### (*FlagSet) BoolFunc <-go1.21.0

```go
func (f *FlagSet) BoolFunc(name, usage string, fn func(string) error)
```

BoolFunc defines a flag with the specified name and usage string without requiring values. Each time the flag is seen, fn is called with the value of the flag. If fn returns a non-nil error, it will be treated as a flag value parsing error.

#### (*FlagSet) BoolVar 

``` go 
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
```

BoolVar defines a bool flag with specified name, default value, and usage string. The argument p points to a bool variable in which to store the value of the flag.

​	BoolVar方法定义具有指定名称、默认值和用法字符串的布尔标志。参数p指向一个bool变量，用于存储标志的值。

#### (*FlagSet) Duration 

``` go 
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
```

Duration defines a time.Duration flag with specified name, default value, and usage string. The return value is the address of a time.Duration variable that stores the value of the flag. The flag accepts a value acceptable to time.ParseDuration.

​	Duration方法定义具有指定名称、默认值和用法字符串的time.Duration标志。返回值是一个time.Duration变量的地址，该变量存储标志的值。该标志接受一个可接受的time.ParseDuration值。

#### (*FlagSet) DurationVar 

``` go 
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

DurationVar defines a time.Duration flag with specified name, default value, and usage string. The argument p points to a time.Duration variable in which to store the value of the flag. The flag accepts a value acceptable to time.ParseDuration.

​	DurationVar方法定义具有指定名称、默认值和用法字符串的time.Duration标志。参数p指向一个time.Duration变量，用于存储标志的值。该标志接受一个可接受的time.ParseDuration值。

#### (*FlagSet) ErrorHandling  <- go1.10

```go 
func (f *FlagSet) ErrorHandling() ErrorHandling
```

ErrorHandling returns the error handling behavior of the flag set.

​	ErrorHandling方法返回标志集的错误处理行为。

#### (*FlagSet) Float64 

```go 
func (f *FlagSet) Float64(name string, value float64, usage string) *float64
```

Float64 defines a float64 flag with specified name, default value, and usage string. The return value is the address of a float64 variable that stores the value of the flag.

​	Float64方法定义了一个指定名称、默认值和用法说明的 float64 标志(flag)。返回值是一个 float64 变量的地址，该变量存储标志的值。

#### (*FlagSet) Float64Var 

```go 
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
```

Float64Var defines a float64 flag with specified name, default value, and usage string. The argument p points to a float64 variable in which to store the value of the flag.

​	Float64Var方法定义了一个指定名称、默认值和用法说明的 float64 标志。参数 p 是一个指向 float64 变量的指针，用于存储标志的值。

#### (*FlagSet) Func  <- go1.16

```go 
func (f *FlagSet) Func(name, usage string, fn func(string) error)
```

Func defines a flag with the specified name and usage string. Each time the flag is seen, fn is called with the value of the flag. If fn returns a non-nil error, it will be treated as a flag value parsing error.

​	Func方法定义了一个指定名称和用法说明的标志。每次遇到该标志时，都会将标志的值传递给 fn 函数。如果 fn 函数返回一个非 nil 错误，则将其视为标志值解析错误。

#### (*FlagSet) Init 

```go 
func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
```

Init sets the name and error handling property for a flag set. By default, the zero FlagSet uses an empty name and the ContinueOnError error handling policy.

​	Init方法为一个标志集设置名称和错误处理属性。默认情况下，零值 FlagSet 使用一个空名称和 ContinueOnError 错误处理策略。

#### (*FlagSet) Int 

```go 
func (f *FlagSet) Int(name string, value int, usage string) *int
```

Int defines an int flag with specified name, default value, and usage string. The return value is the address of an int variable that stores the value of the flag.

​	Int方法定义了一个指定名称、默认值和用法说明的 int 标志。返回值是一个 int 变量的地址，该变量存储标志的值。

#### (*FlagSet) Int64 

``` go 
func (f *FlagSet) Int64(name string, value int64, usage string) *int64
```

Int64 defines an int64 flag with specified name, default value, and usage string. The return value is the address of an int64 variable that stores the value of the flag.

​	Int64 函数定义了一个指定名称、默认值和用法说明的 int64 标志。返回值是一个 int64 变量的地址，该变量存储标志的值。

#### (*FlagSet) Int64Var 

```go 
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var defines an int64 flag with specified name, default value, and usage string. The argument p points to an int64 variable in which to store the value of the flag.

​	Int64Var方法定义了一个指定名称、默认值和用法说明的 int64 标志。参数 p 是一个指向 int64 变量的指针，用于存储标志的值。

#### (*FlagSet) IntVar 

```go 
func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
```

IntVar defines an int flag with specified name, default value, and usage string. The argument p points to an int variable in which to store the value of the flag.

​	IntVar方法定义了一个指定名称、默认值和用法说明的 int 标志。参数 p 是一个指向 int 变量的指针，用于存储标志的值。

#### (*FlagSet) Lookup 

```go 
func (f *FlagSet) Lookup(name string) *Flag
```

Lookup returns the Flag structure of the named flag, returning nil if none exists.

​	Lookup方法返回指定名称的标志(flag)结构体，如果不存在则返回nil。

#### (*FlagSet) NArg 

```go 
func (f *FlagSet) NArg() int
```

NArg is the number of arguments remaining after flags have been processed.

​	NArg方法返回处理完flag后剩余的参数个数。

#### (*FlagSet) NFlag 

```go 
func (f *FlagSet) NFlag() int
```

NFlag returns the number of flags that have been set.

​	NFlag方法返回已设置的flag的数量。

#### (*FlagSet) Name  <- go1.10

```go 
func (f *FlagSet) Name() string
```

Name returns the name of the flag set.

​	Name返回设置的标志的名称。

​	Name方法返回flag set的名称。

#### (*FlagSet) Output  <- go1.10

```go 
func (f *FlagSet) Output() io.Writer
```

Output returns the destination for usage and error messages. os.Stderr is returned if output was not set or was set to nil.

​	Output方法返回用于使用说明和错误信息的目标io.Writer。如果未设置或设置为nil，则返回os.Stderr。

#### (*FlagSet) Parse 

```go 
func (f *FlagSet) Parse(arguments []string) error
```

Parse parses flag definitions from the argument list, which should not include the command name. Must be called after all flags in the FlagSet are defined and before flags are accessed by the program. The return value will be ErrHelp if -help or -h were set but not defined.

​	Parse方法从参数列表中解析flag定义，该列表不应包括命令名称。必须在FlagSet中定义所有flag并在程序访问flag之前调用。如果设置了但未定义-help或-h，则返回ErrHelp。

#### (*FlagSet) Parsed 

```go 
func (f *FlagSet) Parsed() bool
```

Parsed reports whether f.Parse has been called.

​	Parsed方法返回是否已调用f.Parse。

#### (*FlagSet) PrintDefaults 

```go 
func (f *FlagSet) PrintDefaults()
```

PrintDefaults prints, to standard error unless configured otherwise, the default values of all defined command-line flags in the set. See the documentation for the global function PrintDefaults for more information.

​	PrintDefaults方法将标志集合中所有定义的命令行标志的默认值打印到标准错误(除非另有配置)。有关更多信息，请参见全局函数PrintDefaults的文档。

#### (*FlagSet) Set 

```go 
func (f *FlagSet) Set(name, value string) error
```

Set sets the value of the named flag.

​	Set方法设置指定名称的flag的值。

#### (*FlagSet) SetOutput 

```go 
func (f *FlagSet) SetOutput(output io.Writer)
```

SetOutput sets the destination for usage and error messages. If output is nil, os.Stderr is used.

​	SetOutput方法设置用于帮助信息和错误信息输出的目标。如果output为nil，则使用os.Stderr。

#### (*FlagSet) String 

```go 
func (f *FlagSet) String(name string, value string, usage string) *string
```

String defines a string flag with specified name, default value, and usage string. The return value is the address of a string variable that stores the value of the flag.

​	String方法定义了一个指定名称、默认值和用法说明的字符串标志。返回值是一个string类型变量的地址，该变量存储标志的值。

#### (*FlagSet) StringVar 

```go 
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
```

StringVar defines a string flag with specified name, default value, and usage string. The argument p points to a string variable in which to store the value of the flag.

​	StringVar方法定义了一个指定名称、默认值和用法说明的字符串标志。参数p指向一个string类型的变量，用于存储标志的值。

#### (*FlagSet) TextVar  <- go1.19

```go 
func (f *FlagSet) TextVar(p encoding.TextUnmarshaler, name string, value encoding.TextMarshaler, usage string)
```

TextVar defines a flag with a specified name, default value, and usage string. The argument p must be a pointer to a variable that will hold the value of the flag, and p must implement encoding.TextUnmarshaler. If the flag is used, the flag value will be passed to p's UnmarshalText method. The type of the default value must be the same as the type of p.

​	TextVar方法定义了一个指定名称、默认值和用法说明的标志。参数p必须是一个指向将保存标志值的变量的指针，p必须实现encoding.TextUnmarshaler接口。如果使用了该标志，则标志值将传递给p的UnmarshalText方法。默认值的类型必须与p的类型相同。

#### (*FlagSet) Uint 

```go 
func (f *FlagSet) Uint(name string, value uint, usage string) *uint
```

Uint defines a uint flag with specified name, default value, and usage string. The return value is the address of a uint variable that stores the value of the flag.

​	Uint方法定义了一个指定名称、默认值和用法说明的uint标志。返回值是一个uint变量的地址，该变量存储标志的值。

#### (*FlagSet) Uint64 

```go 
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
```

Uint64 defines a uint64 flag with specified name, default value, and usage string. The return value is the address of a uint64 variable that stores the value of the flag.

​	Uint64方法定义了一个指定名称、默认值和用法说明的uint64标志。返回值是一个uint64变量的地址，该变量存储标志的值。

#### (*FlagSet) Uint64Var 

```go 
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
```

Uint64Var defines a uint64 flag with specified name, default value, and usage string. The argument p points to a uint64 variable in which to store the value of the flag.

​	Uint64Var方法定义了一个指定名称、默认值和用法说明的uint64标志。参数p指向一个uint64类型的变量，用于存储标志的值。

#### (*FlagSet) UintVar 

```go 
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
```

UintVar defines a uint flag with specified name, default value, and usage string. The argument p points to a uint variable in which to store the value of the flag.

​	UintVar方法定义了一个指定名称、默认值和用法说明的uint标志。参数p指向一个uint类型的变量，用于存储标志的值。

#### (*FlagSet) Var 

```go 
func (f *FlagSet) Var(value Value, name string, usage string)
```

Var defines a flag with the specified name and usage string. The type and value of the flag are represented by the first argument, of type Value, which typically holds a user-defined implementation of Value. For instance, the caller could create a flag that turns a comma-separated string into a slice of strings by giving the slice the methods of Value; in particular, Set would decompose the comma-separated string into the slice.

​	Var方法定义了一个指定名称和用法说明的标志。标志的类型和值由第一个参数Value表示，通常Value持有一个用户定义的实现。例如，调用方可以创建一个标志，通过给该slice的方法提供Value，将逗号分隔的字符串转换为字符串的slice；特别地，Set会将逗号分隔的字符串分解为slice。

#### (*FlagSet) Visit 

```go 
func (f *FlagSet) Visit(fn func(*Flag))
```

Visit visits the flags in lexicographical order, calling fn for each. It visits only those flags that have been set.

​	Visit方法按词典顺序访问设置了值的标志，对于每个标志调用fn函数。它只访问已设置的标志。

#### (*FlagSet) VisitAll 

```go 
func (f *FlagSet) VisitAll(fn func(*Flag))
```

VisitAll visits the flags in lexicographical order, calling fn for each. It visits all flags, even those not set.

​	VisitAll方法按词典顺序访问所有标志，对于每个标志调用fn函数，即使标志未设置也会访问。

### type Getter  <- go1.2

```go 
type Getter interface {
	Value
	Get() any
}
```

Getter is an interface that allows the contents of a Value to be retrieved. It wraps the Value interface, rather than being part of it, because it appeared after Go 1 and its compatibility rules. All Value types provided by this package satisfy the Getter interface, except the type used by Func.

​	Getter是一个接口，允许检索存储在标志中的值。它包装了Value接口，而不是作为其一部分，因为它出现在Go 1之后，并具有其兼容性规则。由此包提供的所有Value类型都满足Getter接口，除了Func使用的类型。

### type Value 

```go 
type Value interface {
	String() string
	Set(string) error
}
```

Value is the interface to the dynamic value stored in a flag. (The default value is represented as a string.)

​	Value是标志中存储的动态值的接口(默认值表示为字符串)。

If a Value has an IsBoolFlag() bool method returning true, the command-line parser makes -name equivalent to -name=true rather than using the next command-line argument.

​	如果Value具有IsBoolFlag() bool方法返回true，则命令行解析器将-name等同于-name=true，而不是使用下一个命令行参数。

Set is called once, in command line order, for each flag present. The flag package may call the String method with a zero-valued receiver, such as a nil pointer.

​	对于每个存在的标志，Set在命令行顺序中调用一次。标志包可以使用零值接收器(如nil指针)调用String方法。

#### Value Example

```go 
package main

import (
	"flag"
	"fmt"
	"net/url"
)

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func main() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&URLValue{u}, "url", "URL to parse")

	fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

}
Output:

{scheme: "https", host: "golang.org", path: "/pkg/flag/"}

```





