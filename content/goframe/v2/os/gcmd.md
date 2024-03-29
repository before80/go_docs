+++
title = "gcmd"
date = 2024-03-21T17:54:54+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcmd](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcmd)

Package gcmd provides console operations, like options/arguments reading and command running.

​	软件包 gcmd 提供控制台操作，例如选项/参数读取和命令运行。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gcmd/gcmd.go#L20)

```go
const (
	CtxKeyParser    gctx.StrKey = `CtxKeyParser`
	CtxKeyCommand   gctx.StrKey = `CtxKeyCommand`
	CtxKeyArguments gctx.StrKey = `CtxKeyArguments`
)
```

## 变量

This section is empty.

## 函数

#### func BuildOptions

```go
func BuildOptions(m map[string]string, prefix ...string) string
```

BuildOptions builds the options as string.

​	BuildOptions 将选项构建为字符串。

#### func GetArg

```go
func GetArg(index int, def ...string) *gvar.Var
```

GetArg returns the argument at `index` as gvar.Var.

​	GetArg skilar röksemdafærslunni eins `index` og var.意识到的。

##### Example
``` go
```

#### func GetArgAll

```go
func GetArgAll() []string
```

GetArgAll returns all parsed arguments.

​	GetArgAll 返回所有已分析的参数。

##### Example

``` go
```

#### func GetOpt

```go
func GetOpt(name string, def ...string) *gvar.Var
```

GetOpt returns the option value named `name` as gvar.Var.

​	GetOpt 返回名为 `name` gvar 的选项值。是。

##### Example

``` go
```

#### func GetOptAll

```go
func GetOptAll() map[string]string
```

GetOptAll returns all parsed options.

​	GetOptAll 返回所有已分析的选项。

##### Example

``` go
```

#### func GetOptWithEnv

```go
func GetOptWithEnv(key string, def ...interface{}) *gvar.Var
```

GetOptWithEnv returns the command line argument of the specified `key`. If the argument does not exist, then it returns the environment variable with specified `key`. It returns the default value `def` if none of them exists.

​	GetOptWithEnv 返回指定 `key` .如果参数不存在，则返回具有指定 `key` .如果它们都不存在，则返回默认值 `def` 。

Fetching Rules: 1. Command line arguments are in lowercase format, eg: gf.`package name`.; 2. Environment arguments are in uppercase format, eg: GF_`package name`_；

​	获取规则： 1.命令行参数为小写格式，例如：gf。 `package name` .;2. 环境参数为大写格式，例如：GF_ `package name` _;

##### Example

``` go
```

#### func Init

```go
func Init(args ...string)
```

Init does custom initialization.

​	Init 执行自定义初始化。

##### Example

``` go
```

#### func Scan

```go
func Scan(info ...interface{}) string
```

Scan prints `info` to stdout, reads and returns user input, which stops by ‘\n’.

​	扫描打印 `info` 到 stdout，读取并返回用户输入，该输入以“\n”停止。

##### Example

``` go
```

#### func Scanf

```go
func Scanf(format string, info ...interface{}) string
```

Scanf prints `info` to stdout with `format`, reads and returns user input, which stops by ‘\n’.

​	Scanf 打印 `info` 到 stdout with `format` ，读取并返回用户输入，该输入以 '\n' 停止。

##### Example

``` go
```

## 类型

### type Argument

```go
type Argument struct {
	Name   string // Option name.
	Short  string // Option short.
	Brief  string // Brief info about this Option, which is used in help info.
	IsArg  bool   // IsArg marks this argument taking value from command line argument instead of option.
	Orphan bool   // Whether this Option having or having no value bound to it.
}
```

Argument is the command value that are used by certain command.

​	参数是某个命令使用的命令值。

### type Command

```go
type Command struct {
	Name          string        // Command name(case-sensitive).
	Usage         string        // A brief line description about its usage, eg: gf build main.go [OPTION]
	Brief         string        // A brief info that describes what this command will do.
	Description   string        // A detailed description.
	Arguments     []Argument    // Argument array, configuring how this command act.
	Func          Function      // Custom function.
	FuncWithValue FuncWithValue // Custom function with output parameters that can interact with command caller.
	HelpFunc      Function      // Custom help function
	Examples      string        // Usage examples.
	Additional    string        // Additional info about this command, which will be appended to the end of help info.
	Strict        bool          // Strict parsing options, which means it returns error if invalid option given.
	CaseSensitive bool          // CaseSensitive parsing options, which means it parses input options in case-sensitive way.
	Config        string        // Config node name, which also retrieves the values from config component along with command line.
	// contains filtered or unexported fields
}
```

Command holds the info about an argument that can handle custom logic.

​	Command 保存有关可处理自定义逻辑的参数的信息。

#### func CommandFromCtx

```go
func CommandFromCtx(ctx context.Context) *Command
```

CommandFromCtx retrieves and returns Command from context.

​	CommandFromCtx 从上下文中检索并返回 Command。

##### Example

``` go
```

#### func NewFromObject

```go
func NewFromObject(object interface{}) (rootCmd *Command, err error)
```

NewFromObject creates and returns a root command object using given object.

​	NewFromObject 使用给定对象创建并返回根命令对象。

#### (*Command) AddCommand

```go
func (c *Command) AddCommand(commands ...*Command) error
```

AddCommand adds one or more sub-commands to current command.

​	AddCommand 将一个或多个子命令添加到当前命令。

##### Example

``` go
```

#### (*Command) AddObject

```go
func (c *Command) AddObject(objects ...interface{}) error
```

AddObject adds one or more sub-commands to current command using struct object.

​	AddObject 使用 struct object 将一个或多个子命令添加到当前命令。

##### Example

``` go
```

#### (*Command) Print

```go
func (c *Command) Print()
```

Print prints help info to stdout for current command.

​	打印 打印当前命令的帮助信息到 stdout。

##### Example

``` go
```

#### (*Command) PrintTo

```go
func (c *Command) PrintTo(writer io.Writer)
```

PrintTo prints help info to custom io.Writer.

​	PrintTo 将帮助信息打印到自定义 io。作家。

#### (*Command) Run

```go
func (c *Command) Run(ctx context.Context)
```

Run calls custom function in os.Args that bound to this command. It exits this process with exit code 1 if any error occurs.

​	在操作系统中运行调用自定义函数。绑定到此命令的参数。如果发生任何错误，它将退出此进程，并显示退出代码 1。

#### (*Command) RunWithError

```go
func (c *Command) RunWithError(ctx context.Context) (err error)
```

RunWithError calls custom function in os.Args that bound to this command with error output.

​	RunWithError 在 os 中调用自定义函数。绑定到此命令的参数，输出错误。

#### (*Command) RunWithSpecificArgs

```go
func (c *Command) RunWithSpecificArgs(ctx context.Context, args []string) (value interface{}, err error)
```

RunWithSpecificArgs calls custom function in specific args that bound to this command with value and error output.

​	RunWithSpecificArgs 在绑定到此命令的特定参数中调用自定义函数，并提供值和错误输出。

#### (*Command) RunWithValue

```go
func (c *Command) RunWithValue(ctx context.Context) (value interface{})
```

RunWithValue calls custom function in os.Args that bound to this command with value output. It exits this process with exit code 1 if any error occurs.

​	RunWithValue 在 os 中调用自定义函数。绑定到此命令的参数，并带有值输出。如果发生任何错误，它将退出此进程，并显示退出代码 1。

#### (*Command) RunWithValueError

```go
func (c *Command) RunWithValueError(ctx context.Context) (value interface{}, err error)
```

RunWithValueError calls custom function in os.Args that bound to this command with value and error output.

​	RunWithValueError 在 os 中调用自定义函数。绑定到此命令的参数，并输出值和错误。

### type FuncWithValue

```go
type FuncWithValue func(ctx context.Context, parser *Parser) (out interface{}, err error)
```

FuncWithValue is similar like Func but with output parameters that can interact with command caller.

​	FuncWithValue 类似于 Func，但具有可以与命令调用者交互的输出参数。

### type Function

```go
type Function func(ctx context.Context, parser *Parser) (err error)
```

Function is a custom command callback function that is bound to a certain argument.

​	函数是绑定到某个参数的自定义命令回调函数。

### type Parser

```go
type Parser struct {
	// contains filtered or unexported fields
}
```

Parser for arguments.

​	参数的解析器。

#### func Parse

```go
func Parse(supportedOptions map[string]bool, option ...ParserOption) (*Parser, error)
```

Parse creates and returns a new Parser with os.Args and supported options.

​	Parse 创建并返回一个带有 os 的新解析器。参数和支持的选项。

Note that the parameter `supportedOptions` is as [option name: need argument], which means the value item of `supportedOptions` indicates whether corresponding option name needs argument or not.

​	需要注意的是，参数 `supportedOptions` 为[option name： need argument]，表示值 `supportedOptions` 项表示对应的选项名称是否需要参数。

The optional parameter `strict` specifies whether stops parsing and returns error if invalid option passed.

​	可选参数 `strict` 指定是否停止解析，如果传递了无效选项，则返回错误。

##### Example

``` go
```

#### func ParseArgs

```go
func ParseArgs(args []string, supportedOptions map[string]bool, option ...ParserOption) (*Parser, error)
```

ParseArgs creates and returns a new Parser with given arguments and supported options.

​	ParseArgs 创建并返回一个具有给定参数和支持选项的新解析器。

Note that the parameter `supportedOptions` is as [option name: need argument], which means the value item of `supportedOptions` indicates whether corresponding option name needs argument or not.

​	需要注意的是，参数 `supportedOptions` 为[option name： need argument]，表示值 `supportedOptions` 项表示对应的选项名称是否需要参数。

The optional parameter `strict` specifies whether stops parsing and returns error if invalid option passed.

​	可选参数 `strict` 指定是否停止解析，如果传递了无效选项，则返回错误。

##### Example

``` go
```

#### func ParserFromCtx

```go
func ParserFromCtx(ctx context.Context) *Parser
```

ParserFromCtx retrieves and returns Parser from context.

​	ParserFromCtx 从上下文中检索并返回 Parser。

##### Example

``` go
```

#### (*Parser) GetArg

```go
func (p *Parser) GetArg(index int, def ...string) *gvar.Var
```

GetArg returns the argument at `index` as gvar.Var.

​	GetArg skilar röksemdafærslunni eins `index` og var.意识到的。

##### Example

``` go
```

#### (*Parser) GetArgAll

```go
func (p *Parser) GetArgAll() []string
```

GetArgAll returns all parsed arguments.

​	GetArgAll 返回所有已分析的参数。

#### (*Parser) GetOpt

```go
func (p *Parser) GetOpt(name string, def ...interface{}) *gvar.Var
```

GetOpt returns the option value named `name` as gvar.Var.

​	GetOpt 返回名为 `name` gvar 的选项值。是。

#### (*Parser) GetOptAll

```go
func (p *Parser) GetOptAll() map[string]string
```

GetOptAll returns all parsed options.

​	GetOptAll 返回所有已分析的选项。

#### (Parser) MarshalJSON

```go
func (p Parser) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

### type ParserOption <-2.1.0

```go
type ParserOption struct {
	CaseSensitive bool // Marks options parsing in case-sensitive way.
	Strict        bool // Whether stops parsing and returns error if invalid option passed.
}
```

ParserOption manages the parsing options.

​	ParserOption 管理解析选项。