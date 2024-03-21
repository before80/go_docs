+++
title = "gcmd"
date = 2024-03-21T17:54:54+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcmd

Package gcmd provides console operations, like options/arguments reading and command running.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gcmd/gcmd.go#L20)

``` go
const (
	CtxKeyParser    gctx.StrKey = `CtxKeyParser`
	CtxKeyCommand   gctx.StrKey = `CtxKeyCommand`
	CtxKeyArguments gctx.StrKey = `CtxKeyArguments`
)
```

### Variables 

This section is empty.

### Functions 

##### func BuildOptions 

``` go
func BuildOptions(m map[string]string, prefix ...string) string
```

BuildOptions builds the options as string.

##### func GetArg 

``` go
func GetArg(index int, def ...string) *gvar.Var
```

GetArg returns the argument at `index` as gvar.Var.

##### Example

``` go
```
##### func GetArgAll 

``` go
func GetArgAll() []string
```

GetArgAll returns all parsed arguments.

##### Example

``` go
```
##### func GetOpt 

``` go
func GetOpt(name string, def ...string) *gvar.Var
```

GetOpt returns the option value named `name` as gvar.Var.

##### Example

``` go
```
##### func GetOptAll 

``` go
func GetOptAll() map[string]string
```

GetOptAll returns all parsed options.

##### Example

``` go
```
##### func GetOptWithEnv 

``` go
func GetOptWithEnv(key string, def ...interface{}) *gvar.Var
```

GetOptWithEnv returns the command line argument of the specified `key`. If the argument does not exist, then it returns the environment variable with specified `key`. It returns the default value `def` if none of them exists.

Fetching Rules: 1. Command line arguments are in lowercase format, eg: gf.`package name`.<variable name>; 2. Environment arguments are in uppercase format, eg: GF_`package name`_<variable name>；

##### Example

``` go
```
##### func Init 

``` go
func Init(args ...string)
```

Init does custom initialization.

##### Example

``` go
```
##### func Scan 

``` go
func Scan(info ...interface{}) string
```

Scan prints `info` to stdout, reads and returns user input, which stops by '\n'.

##### Example

``` go
```
##### func Scanf 

``` go
func Scanf(format string, info ...interface{}) string
```

Scanf prints `info` to stdout with `format`, reads and returns user input, which stops by '\n'.

##### Example

``` go
```
### Types 

#### type Argument 

``` go
type Argument struct {
	Name   string // Option name.
	Short  string // Option short.
	Brief  string // Brief info about this Option, which is used in help info.
	IsArg  bool   // IsArg marks this argument taking value from command line argument instead of option.
	Orphan bool   // Whether this Option having or having no value bound to it.
}
```

Argument is the command value that are used by certain command.

#### type Command 

``` go
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

##### func CommandFromCtx 

``` go
func CommandFromCtx(ctx context.Context) *Command
```

CommandFromCtx retrieves and returns Command from context.

##### Example

``` go
```
##### func NewFromObject 

``` go
func NewFromObject(object interface{}) (rootCmd *Command, err error)
```

NewFromObject creates and returns a root command object using given object.

##### (*Command) AddCommand 

``` go
func (c *Command) AddCommand(commands ...*Command) error
```

AddCommand adds one or more sub-commands to current command.

##### Example

``` go
```
##### (*Command) AddObject 

``` go
func (c *Command) AddObject(objects ...interface{}) error
```

AddObject adds one or more sub-commands to current command using struct object.

##### Example

``` go
```
##### (*Command) Print 

``` go
func (c *Command) Print()
```

Print prints help info to stdout for current command.

##### Example

``` go
```
##### (*Command) PrintTo <-2.1.0

``` go
func (c *Command) PrintTo(writer io.Writer)
```

PrintTo prints help info to custom io.Writer.

##### (*Command) Run 

``` go
func (c *Command) Run(ctx context.Context)
```

Run calls custom function in os.Args that bound to this command. It exits this process with exit code 1 if any error occurs.

##### (*Command) RunWithError 

``` go
func (c *Command) RunWithError(ctx context.Context) (err error)
```

RunWithError calls custom function in os.Args that bound to this command with error output.

##### (*Command) RunWithSpecificArgs <-2.6.3

``` go
func (c *Command) RunWithSpecificArgs(ctx context.Context, args []string) (value interface{}, err error)
```

RunWithSpecificArgs calls custom function in specific args that bound to this command with value and error output.

##### (*Command) RunWithValue 

``` go
func (c *Command) RunWithValue(ctx context.Context) (value interface{})
```

RunWithValue calls custom function in os.Args that bound to this command with value output. It exits this process with exit code 1 if any error occurs.

##### (*Command) RunWithValueError 

``` go
func (c *Command) RunWithValueError(ctx context.Context) (value interface{}, err error)
```

RunWithValueError calls custom function in os.Args that bound to this command with value and error output.

#### type FuncWithValue 

``` go
type FuncWithValue func(ctx context.Context, parser *Parser) (out interface{}, err error)
```

FuncWithValue is similar like Func but with output parameters that can interact with command caller.

#### type Function 

``` go
type Function func(ctx context.Context, parser *Parser) (err error)
```

Function is a custom command callback function that is bound to a certain argument.

#### type Parser 

``` go
type Parser struct {
	// contains filtered or unexported fields
}
```

Parser for arguments.

##### func Parse 

``` go
func Parse(supportedOptions map[string]bool, option ...ParserOption) (*Parser, error)
```

Parse creates and returns a new Parser with os.Args and supported options.

Note that the parameter `supportedOptions` is as [option name: need argument], which means the value item of `supportedOptions` indicates whether corresponding option name needs argument or not.

The optional parameter `strict` specifies whether stops parsing and returns error if invalid option passed.

##### Example

``` go
```
##### func ParseArgs 

``` go
func ParseArgs(args []string, supportedOptions map[string]bool, option ...ParserOption) (*Parser, error)
```

ParseArgs creates and returns a new Parser with given arguments and supported options.

Note that the parameter `supportedOptions` is as [option name: need argument], which means the value item of `supportedOptions` indicates whether corresponding option name needs argument or not.

The optional parameter `strict` specifies whether stops parsing and returns error if invalid option passed.

##### Example

``` go
```
##### func ParserFromCtx 

``` go
func ParserFromCtx(ctx context.Context) *Parser
```

ParserFromCtx retrieves and returns Parser from context.

##### Example

``` go
```
##### (*Parser) GetArg 

``` go
func (p *Parser) GetArg(index int, def ...string) *gvar.Var
```

GetArg returns the argument at `index` as gvar.Var.

##### Example

``` go
```
##### (*Parser) GetArgAll 

``` go
func (p *Parser) GetArgAll() []string
```

GetArgAll returns all parsed arguments.

##### (*Parser) GetOpt 

``` go
func (p *Parser) GetOpt(name string, def ...interface{}) *gvar.Var
```

GetOpt returns the option value named `name` as gvar.Var.

##### (*Parser) GetOptAll 

``` go
func (p *Parser) GetOptAll() map[string]string
```

GetOptAll returns all parsed options.

##### (Parser) MarshalJSON 

``` go
func (p Parser) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

#### type ParserOption <-2.1.0

``` go
type ParserOption struct {
	CaseSensitive bool // Marks options parsing in case-sensitive way.
	Strict        bool // Whether stops parsing and returns error if invalid option passed.
}
```

ParserOption manages the parsing options.