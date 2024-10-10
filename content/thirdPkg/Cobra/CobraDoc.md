+++
title = "Cobra 在pkg.go.dev上的文档"
type = "docs"
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false
+++



# Cobra 在pkg.go.dev上的文档

> 原文：[https://pkg.go.dev/github.com/spf13/cobra](https://pkg.go.dev/github.com/spf13/cobra)
>
> 版本：v1.7.0
>
> 发布日期：2023.3.22
>
> github网址：[https://github.com/spf13/cobra](https://github.com/spf13/cobra)



## 概述

​	`cobra`包是一个命令行工具（commander），提供了一个简单的接口来创建强大的现代命令行界面（CLI界面）。除了提供接口外，Cobra还同时提供了一个控制器（controller ），用于组织应用程序代码。



## 常量

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/bash_completions.go#L29)

``` go
const (
	BashCompFilenameExt     = "cobra_annotation_bash_completion_filename_extensions"
	BashCompCustom          = "cobra_annotation_bash_completion_custom"
	BashCompOneRequiredFlag = "cobra_annotation_bash_completion_one_required_flag"
	BashCompSubdirsInDir    = "cobra_annotation_bash_completion_subdirs_in_dir"
)
```

​	用于 Bash 自动完成的注解。

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/completions.go#L26)

``` go
const (
	// ShellCompRequestCmd is the name of the hidden command that is used to request
	// completion results from the program.  It is used by the shell completion scripts.
    // ShellCompRequestCmd 是用于请求程序的完成结果的隐藏命令的名称。
	// 它被 shell 完成脚本使用。
	ShellCompRequestCmd = "__complete"
	// ShellCompNoDescRequestCmd is the name of the hidden command that is used to request
	// completion results without their description.  It is used by the shell completion scripts.
    // ShellCompNoDescRequestCmd 是用于请求完成结果但不包含描述的隐藏命令的名称。
	// 它被 shell 完成脚本使用。
	ShellCompNoDescRequestCmd = "__completeNoDesc"
)
```

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/command.go#L33)

``` go
const FlagSetByCobraAnnotation = "cobra_annotation_flag_set_by_cobra"
```

## 变量 

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L61)

```go
var EnableCaseInsensitive = defaultCaseInsensitive
```

​	`EnableCaseInsensitive` 变量允许命令名称不区分大小写。默认情况下是大小写敏感的。

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L58)

```go
var EnableCommandSorting = defaultCommandSorting
```

​	`EnableCommandSorting` 变量控制命令的排序， 默认情况下打开。要禁用排序，请将其设置为 false。

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L54)

```go
var EnablePrefixMatching = defaultPrefixMatching
```

​	`EnablePrefixMatching` 变量允许设置自动前缀匹配。自动前缀匹配在 CLI 工具中自动启用可能是一件危险的事情。将其设置为 true 以启用。

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L76)

```go
var MousetrapDisplayDuration = 5 * time.Second
```

​	`MousetrapDisplayDuration` 变量控制如果从 `explorer.exe` 启动 CLI，则 Windows 上的 `MousetrapHelpText` 消息显示的持续时间。将其设置为 0 以等待按下回车键。要禁用 `Mousetrap`，只需将 `MousetrapHelpText` 变量设置为空字符串（""）。仅在 Microsoft Windows 上工作。

[View Source](https://github.com/spf13/cobra/blob/v1.7.0/cobra.go#L67)

```go
var MousetrapHelpText = `This is a command line tool.

You need to open cmd.exe and run it from there.
`
```

​	`MousetrapHelpText`变量在 Windows 上启动 CLI 时，如果从 `explorer.exe` 启动 CLI，则会在屏幕上显示一个信息闪屏屏幕。要禁用 `Mousetrap`，只需将此变量设置为空字符串（""）。仅在 Microsoft Windows 上工作。

## 函数

#### func AddTemplateFunc 

``` go
func AddTemplateFunc(name string, tmplFunc interface{})
```

​	`AddTemplateFunc` 函数添加一个模板函数，可用于 Usage 和 Help 模板生成。

#### func AddTemplateFuncs 

``` go
func AddTemplateFuncs(tmplFuncs template.FuncMap)
```

​	`AddTemplateFuncs` 函数添加多个模板函数，可用于 Usage 和 Help 模板生成。

#### func AppendActiveHelp  <- v1.5.0

``` go
func AppendActiveHelp(compArray []string, activeHelpStr string) []string
```

AppendActiveHelp adds the specified string to the specified array to be used as ActiveHelp. Such strings will be processed by the completion script and will be shown as ActiveHelp to the user. The array parameter should be the array that will contain the completions. This function can be called multiple times before and/or after completions are added to the array. Each time this function is called with the same array, the new ActiveHelp line will be shown below the previous ones when completion is triggered.

​	`AppendActiveHelp` 函数将指定的字符串添加到指定的数组中，用作 ActiveHelp。这些字符串将由完成脚本处理，并将作为 ActiveHelp 显示给用户。数组参数应该是包含完成的数组。此函数可以在添加到数组之前和/或之后多次调用。每次在使用相同的数组调用此函数时，当触发完成时，新的 ActiveHelp 行将显示在前面的行下方。

#### func ArbitraryArgs 

``` go
func ArbitraryArgs(cmd *Command, args []string) error
```

​	`ArbitraryArgs`函数 从不返回错误。

#### func CheckErr  <- v1.1.2

``` go
func CheckErr(msg interface{})
```

​	`CheckErr` 函数打印带有前缀 'Error:' 的消息，并以错误码 1 退出。如果消息为 nil，则不执行任何操作。

#### func CompDebug  <- v1.0.0

``` go
func CompDebug(msg string, printToStdErr bool)
```

CompDebug prints the specified string to the same file as where the completion script prints its logs. Note that completion printouts should never be on stdout as they would be wrongly interpreted as actual completion choices by the completion script.

​	`CompDebug` 函数将指定的字符串打印到与完成脚本打印其日志的同一文件中。注意，完成的打印输出不应在 stdout 上，因为完成脚本会错误地将其解释为实际的完成选择。

#### func CompDebugln  <- v1.0.0

``` go
func CompDebugln(msg string, printToStdErr bool)
```

​	`CompDebugln` 函数将带有换行符的指定字符串打印到与完成脚本打印其日志的同一文件中。只有在用户设置了环境变量 `BASH_COMP_DEBUG_FILE` 为要使用的某个文件的路径时，这些日志才会被打印出来。

#### func CompError  <- v1.0.0

``` go
func CompError(msg string)
```

​	`CompError` 函数将指定的完成消息打印到 stderr。

#### func CompErrorln  <- v1.0.0

``` go
func CompErrorln(msg string)
```

​	`CompErrorln` 函数将带有换行符的指定完成消息打印到 stderr。

#### func Eq 

``` go
func Eq(a interface{}, b interface{}) bool
```

​	`Eq` 函数接受两种类型并检查它们是否相等。支持的类型是 int 和 string。不支持的类型将会引发 panic。

#### func FixedCompletions  <- v1.5.0

``` go
func FixedCompletions(choices []string, directive ShellCompDirective) func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)
```

​	`FixedCompletions` 函数可用于创建一个总是返回相同结果的完成函数。

#### func GetActiveHelpConfig  <- v1.5.0

``` go
func GetActiveHelpConfig(cmd *Command) string
```

​	`GetActiveHelpConfig` 函数返回 `ActiveHelp` 环境变量的值 `<PROGRAM>_ACTIVE_HELP`，其中 `<PROGRAM>` 是根命令的名称，全部大写，并将所有 `-` 替换为 `_`。如果全局环境变量 `COBRA_ACTIVE_HELP` 设置为 "0"，则它将始终返回 "0"。

#### func Gt 

``` go
func Gt(a interface{}, b interface{}) bool
```

​	`Gt` 函数接受两种类型并检查第一个类型是否大于第二个类型。对于类型 Arrays、Chans、Maps 和 Slices，将比较它们的长度。直接比较 Ints，而将字符串首先解析为 Ints，然后再进行比较。

#### func MarkFlagCustom 

``` go
func MarkFlagCustom(flags *pflag.FlagSet, name string, f string) error
```

MarkFlagCustom adds the BashCompCustom annotation to the named flag, if it exists. The bash completion script will call the bash function f for the flag.

​	`MarkFlagCustom` 函数为命名的标志添加常量`BashCompCustom` 注解（如果存在）。Bash 完成脚本将为该标志调用 bash 函数 `f`。

​	这仅适用于 Bash 完成。建议改为使用 `c.RegisterFlagCompletionFunc(...)`方法，它允许注册一个在所有 shell 中都有效的 Go 函数。

#### func MarkFlagDirname  <- v0.0.5

``` go
func MarkFlagDirname(flags *pflag.FlagSet, name string) error
```

MarkFlagDirname instructs the various shell completion implementations to limit completions for the named flag to directory names.

​	`MarkFlagDirname` 函数指示各种 shell 完成实现将命名标志的完成限制为目录名称。

#### func MarkFlagFilename 

``` go
func MarkFlagFilename(flags *pflag.FlagSet, name string, extensions ...string) error
```

MarkFlagFilename instructs the various shell completion implementations to limit completions for the named flag to the specified file extensions.

​	`MarkFlagFilename` 函数指示各种 shell 完成实现将命名标志的完成限制为指定的文件扩展名。

#### func MarkFlagRequired 

``` go
func MarkFlagRequired(flags *pflag.FlagSet, name string) error
```

​	`MarkFlagRequired` 函数指示各种 shell 完成实现在执行完成时优先考虑命名标志，并且如果在没有标志的情况下调用命令，则会报告错误。

#### func NoArgs 

``` go
func NoArgs(cmd *Command, args []string) error
```

NoArgs returns an error if any args are included.

​	`NoArgs` 函数如果包含任何参数，则返回错误。

#### func OnFinalize  <- v1.6.0

``` go
func OnFinalize(y ...func())
```

​	`OnFinalize` 函数在每个命令的 `Execute` 方法终止时运行传递的函数。

#### func OnInitialize 

``` go
func OnInitialize(y ...func())
```

​	`OnInitialize` 函数在每个命令的 `Execute` 方法调用时运行传递的函数。

#### func OnlyValidArgs 

``` go
func OnlyValidArgs(cmd *Command, args []string) error
```

​	`OnlyValidArgs` 函数如果有任何位置参数不在 `Command` 的 `ValidArgs` 字段中，则返回错误。

#### func WriteStringAndCheck  <- v1.1.2

``` go
func WriteStringAndCheck(b io.StringWriter, s string)
```

​	`WriteStringAndCheck` 函数将字符串写入缓冲区，并检查错误是否为 nil。

## 类型

### type Command 

``` go
type Command struct {
	// Use 是一行的用法消息。
	// 推荐的语法如下：
	//   [ ] 标识可选参数。不在括号内的参数是必需的。
	//   ... 表示您可以为前一个参数指定多个值。
	//   |   表示互斥的信息。您可以使用分隔符左侧或分隔符右侧的参数。在单个命令的使用中不能同时使用这两个参数。
	//   { } 用于分隔互斥的参数集，其中一个参数是必需的。如果参数是可选的，则用括号([ ])括起来。
	// 示例：add [-F file | -D dir]... [-f format] profile
	Use string

    // Aliases 是一个可以用来代替 Use 中第一个单词的别名数组。
	Aliases []string

    // SuggestFor 是一个命令名称的数组，用于建议此命令 - 类似于别名，但仅建议。
	SuggestFor []string

    // Short 是在 'help' 输出中显示的简短描述。
	Short string

    // 此子命令在父命令的 'help' 输出中分组的组 ID。
	GroupID string

    // Long 是在 'help <this-command>' 输出中显示的长消息。
	Long string

    // Example 是如何使用该命令的示例。
	Example string

    // ValidArgs 是接受 shell 完成中的所有有效的非标志参数的列表。
	ValidArgs []string	
    // ValidArgsFunction 是一个可选函数，用于为 shell 完成提供有效的非标志参数。
	// 它是 ValidArgs 的动态版本。
	// ValidArgs 和 ValidArgsFunction 只能在命令中使用一个。
	ValidArgsFunction func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)

    // 预期的参数
	Args PositionalArgs

    // ArgAliases 是 ValidArgs 的别名列表。
	// 这些在 shell 完成中不会向用户建议，但如果手动输入，则会被接受。
	ArgAliases []string

    // BashCompletionFunction 是由传统的 bash 自动完成生成器使用的自定义 bash 函数。
	// 为了与其他 shell 兼容，建议改用 ValidArgsFunction。
	BashCompletionFunction string

    // Deprecated 定义此命令是否已弃用，并且在使用时是否应打印此字符串。
	Deprecated string
	
    // 注解是可以由应用程序用来标识或分组命令的键/值对。
	Annotations map[string]string
	
    // Version 定义此命令的版本。如果此值非空且命令未定义 "version" 标志，将添加一个 "version" 布尔标志到命令中，
	// 并且如果指定，则会打印 "Version" 变量的内容。如果命令尚未定义 "v" 标志，则还会添加一个简写的 "v" 标志。
	Version string
	
    // *Run 函数按以下顺序执行：
	//   * PersistentPreRun()
	//   * PreRun()
	//   * Run()
	//   * PostRun()
	//   * PersistentPostRun()
	// 所有函数都获取相同的参数，即命令名称后的参数。
	//
	// PersistentPreRun：此命令的子命令将继承并执行。
	PersistentPreRun func(cmd *Command, args []string)
    // PersistentPreRunE：PersistentPreRun，但返回错误。
	PersistentPreRunE func(cmd *Command, args []string) error
    // PreRun：此命令的子命令不会继承。
	PreRun func(cmd *Command, args []string)
    // PreRunE：PreRun，但返回错误。
	PreRunE func(cmd *Command, args []string) error
    // Run：通常是实际的工作函数。大多数命令只会实现这个函数。
	Run func(cmd *Command, args []string)
    // RunE：Run，但返回错误。
	RunE func(cmd *Command, args []string) error
    // PostRun：在 Run 命令之后执行。
	PostRun func(cmd *Command, args []string)
    // PostRunE：PostRun，但返回错误。
	PostRunE func(cmd *Command, args []string) error
    // PersistentPostRun：此命令的子命令将继承并在 PostRun 之后执行。
	PersistentPostRun func(cmd *Command, args []string)
    // PersistentPostRunE：PersistentPostRun，但返回错误。
	PersistentPostRunE func(cmd *Command, args []string) error
    
    

    // FParseErrWhitelist 用于忽略标志解析错误
	FParseErrWhitelist FParseErrWhitelist

    // CompletionOptions 是控制 shell 完成处理方式的一组选项
	CompletionOptions CompletionOptions

    // TraverseChildren 在执行子命令之前解析所有父命令上的标志。
	TraverseChildren bool

    // Hidden 定义此命令是否隐藏，不应在可用命令列表中显示。
	Hidden bool

    // SilenceErrors 是一个选项，用于在下游消除错误。
	SilenceErrors bool

    // SilenceUsage 是一个选项，在发生错误时消除用法。
	SilenceUsage bool

    // DisableFlagParsing 禁用标志解析。
	// 如果为 true，则所有标志都将作为参数传递给命令。
	DisableFlagParsing bool

	// DisableAutoGenTag 定义是否打印生成文档时的生成标签 ("Auto generated by spf13/cobra...")。
	DisableAutoGenTag bool
	
    // DisableFlagsInUseLine 在打印帮助或生成文档时，禁用将 [flags] 添加到命令的用法行中。
	DisableFlagsInUseLine bool
	
    // DisableSuggestions 禁用基于 Levenshtein 距离的 'unknown command' 消息的建议。
	DisableSuggestions bool
	
    // SuggestionsMinimumDistance 定义显示建议的最小 Levenshtein 距离。
	// 必须 > 0。
	SuggestionsMinimumDistance int
    // 包含过滤或未公开的字段
}
```

​	Command 就是一个应用程序的命令。例如，'go run ...' 中的 'run' 就是命令。Cobra 要求您在命令定义中定义用法和描述，以确保可用性。

#### (*Command) AddCommand 

``` go
func (c *Command) AddCommand(cmds ...*Command)
```

​	AddCommand 方法将一个或多个命令添加到此父命令中。

#### (*Command) AddGroup  <- v1.6.0

``` go
func (c *Command) AddGroup(groups ...*Group)
```

​	AddGroup 方法将一个或多个命令组添加到此父命令中。	

#### (*Command) AllChildCommandsHaveGroup  <- v1.6.0

``` go
func (c *Command) AllChildCommandsHaveGroup() bool
```

​	AllChildCommandsHaveGroup 方法返回所有子命令是否都分配给了一个组。

#### (*Command) ArgsLenAtDash 

``` go
func (c *Command) ArgsLenAtDash() int
```

​	ArgsLenAtDash 方法将在参数解析过程中找到 '`--`' 时返回 c.Flags().Args 的长度。

#### (*Command) CalledAs  <- v0.0.2

``` go
func (c *Command) CalledAs() string
```

​	CalledAs 方法返回用于调用此命令的命令名称或别名，如果尚未调用命令，则返回空字符串。

#### (*Command) CommandPath 

``` go
func (c *Command) CommandPath() string
```

​	CommandPath 方法返回此命令的完整路径。

#### (*Command) CommandPathPadding 

``` go
func (c *Command) CommandPathPadding() int
```

​	CommandPathPadding 方法返回命令路径的填充。

#### (*Command) Commands 

``` go
func (c *Command) Commands() []*Command
```

​	Commands 方法返回已排序的子命令切片。

#### (*Command) ContainsGroup  <- v1.6.0

``` go
func (c *Command) ContainsGroup(groupID string) bool
```

​	ContainsGroup 方法返回 groupID 是否存在于命令组列表中。

#### (*Command) Context  <- v0.0.6

``` go
func (c *Command) Context() context.Context
```

​	Context 方法返回底层命令上下文。如果使用 ExecuteContext 执行了命令，或者使用 SetContext 设置了上下文，则将返回先前设置的上下文。否则，返回 nil。

​	注意，对 Execute 方法和 ExecuteC 方法的调用将使用命令的 nil 上下文替换为 context.Background，因此在调用这些函数之一后，Context 将返回一个 background 上下文。

#### (*Command) DebugFlags 

``` go
func (c *Command) DebugFlags()
```

​	DebugFlags 方法用于确定已将哪些标志分配给了哪些命令以及哪些标志保留。

#### (*Command) ErrOrStderr  <- v0.0.5

``` go
func (c *Command) ErrOrStderr() io.Writer
```

​	ErrOrStderr 方法返回 stderr 输出。

#### (*Command) Execute 

``` go
func (c *Command) Execute() error
```

​	Execute 方法使用 args（默认为 `os.Args[1:]`）并通过命令树运行，查找适合的命令和相应的标志。

#### (*Command) ExecuteC 

``` go
func (c *Command) ExecuteC() (cmd *Command, err error)
```

​	ExecuteC 方法执行命令。

#### (*Command) ExecuteContext  <- v0.0.6

``` go
func (c *Command) ExecuteContext(ctx context.Context) error
```

​	ExecuteContext 方法与 Execute() 方法相同，但在命令上设置 `ctx`。在 `*Run` 生命周期或 ValidArgs 函数内部调用 `cmd.Context()` 检索 `ctx`。

#### (*Command) ExecuteContextC  <- v1.2.0

``` go
func (c *Command) ExecuteContextC(ctx context.Context) (*Command, error)
```

​	ExecuteContextC 方法与 ExecuteC() 方法相同，但在命令上设置 `ctx`。在 `*Run` 生命周期或 ValidArgs 函数内部调用 `cmd.Context()` 检索 `ctx`。

#### (*Command) Find 

``` go
func (c *Command) Find(args []string) (*Command, []string, error)
```

​	Find 方法根据 args 和命令树查找目标命令。应在最高节点上运行。仅向下搜索。

#### (*Command) Flag 

``` go
func (c *Command) Flag(name string) (flag *flag.Flag)
```

​	Flag 方法通过向上遍历命令树查找匹配的标志。

#### (*Command) FlagErrorFunc 

``` go
func (c *Command) FlagErrorFunc() (f func(*Command, error) error)
```

​	FlagErrorFunc 方法返回为此命令或其父命令设置的函数，或者返回一个返回原始错误的函数。

#### (*Command) Flags 

``` go
func (c *Command) Flags() *flag.FlagSet
```

​	Flags 方法返回适用于此命令的完整 FlagSet（在此命令和所有父命令中本地和持久声明的标志）。

#### (*Command) GenBashCompletion 

``` go
func (c *Command) GenBashCompletion(w io.Writer) error
```

GenBashCompletion generates bash completion file and writes to the passed writer.

​	GenBashCompletion 方法生成 Bash 完成文件并写入传递的写入器。

#### (*Command) GenBashCompletionFile 

``` go
func (c *Command) GenBashCompletionFile(filename string) error
```

​	GenBashCompletionFile 方法生成 Bash 完成文件。

#### (*Command) GenBashCompletionFileV2  <- v1.2.0

``` go
func (c *Command) GenBashCompletionFileV2(filename string, includeDesc bool) error
```

GenBashCompletionFileV2 generates Bash completion version 2.

​	GenBashCompletionFileV2 方法生成 Bash 完成版本 2。

#### (*Command) GenBashCompletionV2  <- v1.2.0

``` go
func (c *Command) GenBashCompletionV2(w io.Writer, includeDesc bool) error
```

​	GenBashCompletionV2 方法生成 Bash 完成文件版本 2 并将其写入传递的写入器。

#### (*Command) GenFishCompletion  <- v1.0.0

``` go
func (c *Command) GenFishCompletion(w io.Writer, includeDesc bool) error
```

​	GenFishCompletion 方法生成 fish 完成文件并写入传递的写入器。

#### (*Command) GenFishCompletionFile  <- v1.0.0

``` go
func (c *Command) GenFishCompletionFile(filename string, includeDesc bool) error
```

​	GenFishCompletionFile 方法生成 fish 完成文件。

#### (*Command) GenPowerShellCompletion  <- v0.0.5

``` go
func (c *Command) GenPowerShellCompletion(w io.Writer) error
```

​	GenPowerShellCompletion 方法生成不带描述的 PowerShell 完成文件并将其写入传递的写入器。

#### (*Command) GenPowerShellCompletionFile  <- v0.0.5

``` go
func (c *Command) GenPowerShellCompletionFile(filename string) error
```

​	GenPowerShellCompletionFile 方法生成不带描述的 PowerShell 完成文件。

#### (*Command) GenPowerShellCompletionFileWithDesc  <- v1.1.2

``` go
func (c *Command) GenPowerShellCompletionFileWithDesc(filename string) error
```

​	GenPowerShellCompletionFileWithDesc 方法生成带有描述的 PowerShell 完成文件。

#### (*Command) GenPowerShellCompletionWithDesc  <- v1.1.2

``` go
func (c *Command) GenPowerShellCompletionWithDesc(w io.Writer) error
```

​	GenPowerShellCompletionWithDesc 方法生成带有描述的 PowerShell 完成文件并将其写入传递的写入器。

#### (*Command) GenZshCompletion 

``` go
func (c *Command) GenZshCompletion(w io.Writer) error
```

​	GenZshCompletion 方法生成包括描述的 zsh 完成文件并将其写入传递的写入器。

#### (*Command) GenZshCompletionFile 

``` go
func (c *Command) GenZshCompletionFile(filename string) error
```

​	GenZshCompletionFile 方法生成包括描述的 zsh 完成文件。

#### (*Command) GenZshCompletionFileNoDesc  <- v1.1.0

``` go
func (c *Command) GenZshCompletionFileNoDesc(filename string) error
```

​	GenZshCompletionFileNoDesc 方法生成不带描述的 zsh 完成文件。

#### (*Command) GenZshCompletionNoDesc  <- v1.1.0

``` go
func (c *Command) GenZshCompletionNoDesc(w io.Writer) error
```

​	GenZshCompletionNoDesc 方法生成不带描述的 zsh 完成文件并将其写入传递的写入器。

#### (*Command) GlobalNormalizationFunc 

``` go
func (c *Command) GlobalNormalizationFunc() func(f *flag.FlagSet, name string) flag.NormalizedName
```

​	GlobalNormalizationFunc 方法返回全局的归一化函数，如果不存在，则返回 nil。

#### (*Command) Groups  <- v1.6.0

``` go
func (c *Command) Groups() []*Group
```

​	Groups 方法返回已排序的子命令组切片。

#### (*Command) HasAlias 

``` go
func (c *Command) HasAlias(s string) bool
```

​	HasAlias 方法确定给定的字符串是否是该命令的别名。

#### (*Command) HasAvailableFlags 

``` go
func (c *Command) HasAvailableFlags() bool
```

​	HasAvailableFlags 方法检查命令是否包含任何标志（本地加上整个结构中的持久标志），这些标志不是隐藏的或已弃用的。

#### (*Command) HasAvailableInheritedFlags 

``` go
func (c *Command) HasAvailableInheritedFlags() bool
```

​	HasAvailableInheritedFlags 方法检查命令是否具有从其父命令继承的标志，这些标志不是隐藏的或已弃用的。

#### (*Command) HasAvailableLocalFlags 

``` go
func (c *Command) HasAvailableLocalFlags() bool
```

​	HasAvailableLocalFlags 方法检查命令是否具有本地明确声明的标志，这些标志不是隐藏的或已弃用的。

#### (*Command) HasAvailablePersistentFlags 

``` go
func (c *Command) HasAvailablePersistentFlags() bool
```

​	HasAvailablePersistentFlags 方法检查命令是否包含不是隐藏的或已弃用的持久性标志。

#### (*Command) HasAvailableSubCommands 

``` go
func (c *Command) HasAvailableSubCommands() bool
```

​	HasAvailableSubCommands 方法确定命令是否有可用的子命令，需要在默认用法/帮助模板下显示在“可用命令”下。

#### (*Command) HasExample 

``` go
func (c *Command) HasExample() bool
```

​	HasExample 方法确定命令是否有示例。

#### (*Command) HasFlags 

``` go
func (c *Command) HasFlags() bool
```

​	HasFlags 方法检查命令是否包含任何标志（本地加上整个结构中的持久标志）。

#### (*Command) HasHelpSubCommands 

``` go
func (c *Command) HasHelpSubCommands() bool
```

​	HasHelpSubCommands 方法确定命令是否具有任何可用的“帮助”子命令，需要在默认用法/帮助模板下显示在“附加帮助主题”下。

#### (*Command) HasInheritedFlags 

``` go
func (c *Command) HasInheritedFlags() bool
```

​	HasInheritedFlags 方法检查命令是否具有从其父命令继承的标志。

#### (*Command) HasLocalFlags 

``` go
func (c *Command) HasLocalFlags() bool
```

​	HasLocalFlags 方法检查命令是否具有本地明确声明的标志。

#### (*Command) HasParent 

``` go
func (c *Command) HasParent() bool
```

​	HasParent 方法确定命令是否为子命令。

#### (*Command) HasPersistentFlags 

``` go
func (c *Command) HasPersistentFlags() bool
```

​	HasPersistentFlags 方法检查命令是否包含持久性标志。

#### (*Command) HasSubCommands 

``` go
func (c *Command) HasSubCommands() bool
```

​	HasSubCommands 方法确定命令是否有子命令。

#### (*Command) Help 

``` go
func (c *Command) Help() error
```

​	Help 方法输出命令的帮助信息。在用户调用 help [command] 时使用。可以通过覆盖 HelpFunc 来定义用户的帮助信息。

#### (*Command) HelpFunc 

``` go
func (c *Command) HelpFunc() func(*Command, []string)
```

​	HelpFunc 方法返回为该命令或其父命令设置的函数，或者返回带有默认帮助行为的函数。

#### (*Command) HelpTemplate 

``` go
func (c *Command) HelpTemplate() string
```

​	HelpTemplate 方法返回命令的帮助模板。

#### (*Command) InOrStdin  <- v0.0.5

``` go
func (c *Command) InOrStdin() io.Reader
```

​	InOrStdin 方法返回用于stdin的输入。

#### (*Command) InheritedFlags 

``` go
func (c *Command) InheritedFlags() *flag.FlagSet
```

​	InheritedFlags 方法返回从父命令继承的所有标志。

#### (*Command) InitDefaultCompletionCmd  <- v1.6.0

``` go
func (c *Command) InitDefaultCompletionCmd()
```

​	InitDefaultCompletionCmd 方法向 `c` 添加默认的 'completion' 命令。如果满足以下任何条件，此函数将不执行任何操作：

- 1- 该功能已被程序明确禁用，

- 2- `c `没有子命令（以避免创建一个），

- 3- `c `已由程序提供 'completion' 命令。

  

#### (*Command) InitDefaultHelpCmd 

``` go
func (c *Command) InitDefaultHelpCmd()
```

​	InitDefaultHelpCmd 方法向 `c` 添加默认的帮助命令。在执行 `c` 或调用 help 和 usage 时，它会自动调用。如果 `c` 已经有帮助命令或 `c` 没有子命令，它将不执行任何操作。

#### (*Command) InitDefaultHelpFlag 

``` go
func (c *Command) InitDefaultHelpFlag()
```

​	InitDefaultHelpFlag 方法向 `c` 添加默认的帮助标志。在执行 `c `或调用 help 和 usage 时，它会自动调用。如果 `c` 已经有帮助标志，它将不执行任何操作。

#### (*Command) InitDefaultVersionFlag  <- v0.0.2

``` go
func (c *Command) InitDefaultVersionFlag()
```

​	InitDefaultVersionFlag 方法向 `c` 添加默认的版本标志。在执行 `c` 时，它会自动调用。如果 `c` 已经有版本标志，它将不执行任何操作。如果 `c.Version` 为空，它将不执行任何操作。

#### (*Command) IsAdditionalHelpTopicCommand 

``` go
func (c *Command) IsAdditionalHelpTopicCommand() bool
```

​	IsAdditionalHelpTopicCommand 方法确定命令是否为附加的帮助主题命令；附加的帮助主题命令是根据它不能运行/隐藏/废弃以及没有可以运行/隐藏/废弃的子命令来确定的。具体示例：https://github.com/spf13/cobra/issues/393#issuecomment-282741924。

#### (*Command) IsAvailableCommand 

``` go
func (c *Command) IsAvailableCommand() bool
```

​	IsAvailableCommand 方法确定命令是否作为非帮助命令可用（这包括所有未废弃/隐藏的命令）。

#### (*Command) LocalFlags 

``` go
func (c *Command) LocalFlags() *flag.FlagSet
```

​	LocalFlags 方法返回当前命令中明确设置的本地 FlagSet。

#### (*Command) LocalNonPersistentFlags 

``` go
func (c *Command) LocalNonPersistentFlags() *flag.FlagSet
```

​	LocalNonPersistentFlags 方法是特定于此命令的标志，不会传递给子命令。

#### (*Command) MarkFlagCustom 

``` go
func (c *Command) MarkFlagCustom(name string, f string) error
```

​	MarkFlagCustom 方法为指定的标志添加 BashCompCustom 注释，如果存在的话。Bash 完成脚本将为该标志调用 bash 函数 `f`。

​	这只适用于 Bash 完成。建议改用 c.RegisterFlagCompletionFunc(...)，该函数允许注册一个适用于所有 shell 的 Go 函数。

#### (*Command) MarkFlagDirname  <- v0.0.5

``` go
func (c *Command) MarkFlagDirname(name string) error
```

​	MarkFlagDirname 方法指示各种 shell 完成实现将完成限制为目录名的指定标志。

#### (*Command) MarkFlagFilename 

``` go
func (c *Command) MarkFlagFilename(name string, extensions ...string) error
```

​	MarkFlagFilename 方法指示各种 shell 完成实现将完成限制为指定的文件扩展名的指定标志。

#### (*Command) MarkFlagRequired 

``` go
func (c *Command) MarkFlagRequired(name string) error
```

​	MarkFlagRequired 方法指示各种 shell 完成实现在执行完成时优先考虑指定的标志，并在不带该标志调用命令时报告错误。

#### (*Command) MarkFlagsMutuallyExclusive  <- v1.5.0

``` go
func (c *Command) MarkFlagsMutuallyExclusive(flagNames ...string)
```

​	MarkFlagsMutuallyExclusive 方法使用注释标记给定的标志，以便在从给定的一组标志中使用多个标志调用命令时，Cobra 会引发错误。

#### (*Command) MarkFlagsRequiredTogether  <- v1.5.0

``` go
func (c *Command) MarkFlagsRequiredTogether(flagNames ...string)
```

​	MarkFlagsRequiredTogether 方法使用注释标记给定的标志，以便在使用给定标志的子集（但不是全部）调用命令时，Cobra 会引发错误。

#### (*Command) MarkPersistentFlagDirname  <- v0.0.5

``` go
func (c *Command) MarkPersistentFlagDirname(name string) error
```

​	MarkPersistentFlagDirname 方法指示各种 shell 完成实现将完成限制为目录名的指定持久性标志。

#### (*Command) MarkPersistentFlagFilename 

``` go
func (c *Command) MarkPersistentFlagFilename(name string, extensions ...string) error
```

​	MarkPersistentFlagFilename 方法指示各种 shell 完成实现将完成限制为指定的文件扩展名的指定持久性标志。

#### (*Command) MarkPersistentFlagRequired 

``` go
func (c *Command) MarkPersistentFlagRequired(name string) error
```

​	MarkPersistentFlagRequired 方法指示各种 shell 完成实现在执行完成时优先考虑指定的持久性标志，并在不带该标志调用命令时报告错误。

#### (*Command) MarkZshCompPositionalArgumentFile  <- v0.0.5

``` go
func (c *Command) MarkZshCompPositionalArgumentFile(argPosition int, patterns ...string) error
```

MarkZshCompPositionalArgumentFile only worked for zsh and its behavior was not consistent with Bash completion. It has therefore been disabled. Instead, when no other completion is specified, file completion is done by default for every argument. One can disable file completion on a per-argument basis by using ValidArgsFunction and ShellCompDirectiveNoFileComp. To achieve file extension filtering, one can use ValidArgsFunction and ShellCompDirectiveFilterFileExt.

​	MarkZshCompPositionalArgumentFile 方法仅适用于 zsh，并且其行为与 Bash 完成的行为不一致。因此，它已被禁用。相反，当未指定其他完成时，默认情况下为每个参数执行文件完成。可以使用 ValidArgsFunction字段 和 ShellCompDirectiveNoFileComp 在每个参数上禁用文件完成。要实现文件扩展名过滤，可以使用 ValidArgsFunction 和 ShellCompDirectiveFilterFileExt。

Deprecated

#### (*Command) MarkZshCompPositionalArgumentWords  <- v0.0.5

``` go
func (c *Command) MarkZshCompPositionalArgumentWords(argPosition int, words ...string) error
```

MarkZshCompPositionalArgumentWords only worked for zsh. It has therefore been disabled. To achieve the same behavior across all shells, one can use ValidArgs (for the first argument only) or ValidArgsFunction for any argument (can include the first one also).

​	MarkZshCompPositionalArgumentWords 方法仅适用于 zsh。因此，它已被禁用。要在所有 shell 中实现相同的行为，可以使用 ValidArgs字段（仅适用于第一个参数）或 ValidArgsFunction字段（对任何参数都可以，也可以包括第一个参数）。

Deprecated

#### (*Command) Name 

``` go
func (c *Command) Name() string
```

Name returns the command's name: the first word in the use line.

​	Name 方法返回命令的名称：use 行中的第一个单词。

#### (*Command) NameAndAliases 

``` go
func (c *Command) NameAndAliases() string
```

​	NameAndAliases 方法返回命令的名称和所有别名的列表。

#### (*Command) NamePadding 

``` go
func (c *Command) NamePadding() int
```

​	NamePadding 方法返回名称的填充。

#### (*Command) NonInheritedFlags 

``` go
func (c *Command) NonInheritedFlags() *flag.FlagSet
```

​	NonInheritedFlags 方法返回未从父命令继承的所有标志。

#### (*Command) OutOrStderr 

``` go
func (c *Command) OutOrStderr() io.Writer
```

​	OutOrStderr 方法返回用于标准错误输出的输出。

#### (*Command) OutOrStdout 

``` go
func (c *Command) OutOrStdout() io.Writer
```

​	OutOrStdout 方法返回用于标准输出的输出。

#### (*Command) Parent 

``` go
func (c *Command) Parent() *Command
```

​	Parent 方法返回命令的父命令。

#### (*Command) ParseFlags 

``` go
func (c *Command) ParseFlags(args []string) error
```

​	ParseFlags 方法解析持久性标志树和本地标志。

#### (*Command) PersistentFlags 

``` go
func (c *Command) PersistentFlags() *flag.FlagSet
```

​	PersistentFlags 方法返回当前命令中明确设置的持久性 FlagSet。

#### (*Command) Print 

``` go
func (c *Command) Print(i ...interface{})
```

​	Print 方法是一个方便的方法，用于将内容打印到已定义的输出，如果未设置，则回退到 Stderr。

#### (*Command) PrintErr  <- v0.0.5

``` go
func (c *Command) PrintErr(i ...interface{})
```

​	PrintErr 方法是一个方便的方法，用于将内容打印到已定义的 Err 输出，如果未设置，则回退到 Stderr。

#### (*Command) PrintErrf  <- v0.0.5

``` go
func (c *Command) PrintErrf(format string, i ...interface{})
```

​	PrintErrf 方法是一个方便的方法，用于将格式化的内容打印到已定义的 Err 输出，如果未设置，则回退到 Stderr。

#### (*Command) PrintErrln  <- v0.0.5

``` go
func (c *Command) PrintErrln(i ...interface{})
```

​	PrintErrln 方法是一个方便的方法，用于将内容打印为一行到已定义的 Err 输出，如果未设置，则回退到 Stderr。

#### (*Command) Printf 

``` go
func (c *Command) Printf(format string, i ...interface{})
```

​	Printf 方法是一个方便的方法，用于将格式化的内容打印到已定义的输出，如果未设置，则回退到 Stderr。

#### (*Command) Println 

``` go
func (c *Command) Println(i ...interface{})
```

​	Println 方法是一个方便的方法，用于将内容打印为一行到已定义的输出，如果未设置，则回退到 Stderr。

#### (*Command) RegisterFlagCompletionFunc  <- v1.0.0

``` go
func (c *Command) RegisterFlagCompletionFunc(flagName string, f func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)) error
```

​	RegisterFlagCompletionFunc 方法应该被调用以注册一个函数，为一个标志提供完成。

#### (*Command) RemoveCommand 

``` go
func (c *Command) RemoveCommand(cmds ...*Command)
```

​	RemoveCommand 方法从父命令中删除一个或多个命令。

#### (*Command) ResetCommands 

``` go
func (c *Command) ResetCommands()
```

​	ResetCommands 方法从 `c` 中删除父命令、子命令和帮助命令。

#### (*Command) ResetFlags 

``` go
func (c *Command) ResetFlags()
```

​	ResetFlags 方法删除命令中的所有标志。

#### (*Command) Root 

``` go
func (c *Command) Root() *Command
```

​	Root 方法查找根命令。

#### (*Command) Runnable 

``` go
func (c *Command) Runnable() bool
```

​	Runnable 方法确定命令是否本身可运行。

#### (*Command) SetArgs 

``` go
func (c *Command) SetArgs(a []string)
```

​	SetArgs 方法为命令设置参数。默认情况下，它设置为 os.Args[1:]，如果需要，可以覆盖，尤其是在测试时非常有用。

#### (*Command) SetCompletionCommandGroupID  <- v1.6.0

``` go
func (c *Command) SetCompletionCommandGroupID(groupID string)
```

​	SetCompletionCommandGroupID 方法设置完成命令的组 ID。

#### (*Command) SetContext  <- v1.5.0

``` go
func (c *Command) SetContext(ctx context.Context)
```

​	SetContext 为命令设置上下文。这个上下文将被 Command.ExecuteContext 或 Command.ExecuteContextC 覆盖。

#### (*Command) SetErr  <- v0.0.5

``` go
func (c *Command) SetErr(newErr io.Writer)
```

​	SetErr 方法设置错误消息的目标。如果 newErr 为 nil，则使用 os.Stderr。

#### (*Command) SetFlagErrorFunc 

``` go
func (c *Command) SetFlagErrorFunc(f func(*Command, error) error)
```

​	SetFlagErrorFunc 方法设置一个函数，以在标志解析失败时生成错误。

#### (*Command) SetGlobalNormalizationFunc 

``` go
func (c *Command) SetGlobalNormalizationFunc(n func(f *flag.FlagSet, name string) flag.NormalizedName)
```

​	SetGlobalNormalizationFunc 方法将标准化函数设置为所有标志集，以及子命令。用户不应在命令上具有循环依赖。

#### (*Command) SetHelpCommand 

``` go
func (c *Command) SetHelpCommand(cmd *Command)
```

​	SetHelpCommand 方法设置帮助命令。

#### (*Command) SetHelpCommandGroupID  <- v1.6.0

``` go
func (c *Command) SetHelpCommandGroupID(groupID string)
```

​	SetHelpCommandGroupID 方法设置帮助命令的组 ID。

#### (*Command) SetHelpFunc 

``` go
func (c *Command) SetHelpFunc(f func(*Command, []string))
```

​	SetHelpFunc 方法设置帮助函数。可以由应用程序定义。

#### (*Command) SetHelpTemplate 

``` go
func (c *Command) SetHelpTemplate(s string)
```

​	SetHelpTemplate 方法设置用于使用帮助的模板。应用程序可以使用它来设置自定义模板。

#### (*Command) SetIn  <- v0.0.5

``` go
func (c *Command) SetIn(newIn io.Reader)
```

​	SetIn 方法设置输入数据的源。如果 newIn 为 nil，则使用 os.Stdin。

#### (*Command) SetOut  <- v0.0.5

``` go
func (c *Command) SetOut(newOut io.Writer)
```

​	SetOut 方法设置用于使用消息的目标。如果 newOut 为 nil，则使用 os.Stdout。

#### (*Command) SetOutput 

``` go
func (c *Command) SetOutput(output io.Writer)
```

​	SetOutput 方法设置用于使用和错误消息的目标。如果 output 为 nil，则使用 os.Stderr。已弃用：请改用 SetOut 和/或 SetErr。

#### (*Command) SetUsageFunc 

``` go
func (c *Command) SetUsageFunc(f func(*Command) error)
```

​	SetUsageFunc 方法设置用法函数。用法可以由应用程序定义。

#### (*Command) SetUsageTemplate 

``` go
func (c *Command) SetUsageTemplate(s string)
```

​	SetUsageTemplate 方法设置用法模板。可以由应用程序定义。

#### (*Command) SetVersionTemplate  <- v0.0.2

``` go
func (c *Command) SetVersionTemplate(s string)
```

​	SetVersionTemplate 方法设置要使用的版本模板。应用程序可以使用它来设置自定义模板。

#### (*Command) SuggestionsFor 

``` go
func (c *Command) SuggestionsFor(typedName string) []string
```

​	SuggestionsFor 方法为 typedName 提供建议。

#### (*Command) Traverse 

``` go
func (c *Command) Traverse(args []string) (*Command, []string, error)
```

​	Traverse （遍历）命令树以查找命令，并为每个父级解析参数。

#### (*Command) Usage 

``` go
func (c *Command) Usage() error
```

​	Usage 方法输出命令的用法信息。当用户提供无效输入时使用。可以通过覆盖 UsageFunc 由用户定义。

#### (*Command) UsageFunc 

``` go
func (c *Command) UsageFunc() (f func(*Command) error)
```

​	UsageFunc 方法返回为该命令或父级设置的 SetUsageFunc 函数，如果没有则返回默认的用法函数。

#### (*Command) UsagePadding 

``` go
func (c *Command) UsagePadding() int
```

​	UsagePadding 方法返回用法的填充。

#### (*Command) UsageString 

``` go
func (c *Command) UsageString() string
```

​	UsageString 方法返回用法字符串。

#### (*Command) UsageTemplate 

``` go
func (c *Command) UsageTemplate() string
```

​	UsageTemplate 方法返回该命令的用法模板。

#### (*Command) UseLine 

``` go
func (c *Command) UseLine() string
```

​	UseLine 方法输出给定命令的完整用法（包括父级）。

#### (*Command) ValidateArgs 

``` go
func (c *Command) ValidateArgs(args []string) error
```

​	ValidateArgs 方法检查参数是否有效。

#### (*Command) ValidateFlagGroups  <- v1.6.0

``` go
func (c *Command) ValidateFlagGroups() error
```

ValidateFlagGroups validates the mutuallyExclusive/requiredAsGroup logic and returns the first error encountered.

​	ValidateFlagGroups 方法验证互斥/必需组逻辑并返回遇到的第一个错误。

#### (*Command) ValidateRequiredFlags  <- v1.6.0

``` go
func (c *Command) ValidateRequiredFlags() error
```

​	ValidateRequiredFlags 方法验证所有必需标志是否存在，否则返回错误。

#### (*Command) VersionTemplate  <- v0.0.2

``` go
func (c *Command) VersionTemplate() string
```

​	VersionTemplate 方法返回该命令的版本模板。

#### (*Command) VisitParents 

``` go
func (c *Command) VisitParents(fn func(*Command))
```

​	VisitParents 方法访问命令的所有父级并在每个父级上调用 `fn`。

#### type CompletionOptions  <- v1.2.0

``` go
type CompletionOptions struct {	
    // DisableDefaultCmd 防止 Cobra 创建默认的 'completion' 命令
	DisableDefaultCmd bool
	
    // DisableNoDescFlag 防止 Cobra 为支持完成描述的 shell 创建 '--no-descriptions' 标志
	DisableNoDescFlag bool

    // DisableDescriptions 关闭所有支持它们的 shell 的完成描述
	DisableDescriptions bool
    
    // HiddenDefaultCmd 使默认的 'completion' 命令隐藏
	HiddenDefaultCmd bool
}
```

​	CompletionOptions 结构体是控制 shell 完成的选项。

### type FParseErrWhitelist  <- v0.0.3

``` go
type FParseErrWhitelist flag.ParseErrorsWhitelist
```

​	FParseErrWhitelist 配置 将忽略的标志解析错误。

### type Group  <- v1.6.0

``` go
type Group struct {
	ID    string
	Title string
}
```

​	Group 结构体用于管理命令的分组。

### type PositionalArgs 

``` go
type PositionalArgs func(cmd *Command, args []string) error
```

#### func ExactArgs 

``` go
func ExactArgs(n int) PositionalArgs
```

​	ExactArgs 函数如果参数数量不是 n，则返回错误。

#### func ExactValidArgs DEPRECATED



#### func MatchAll  <- v1.3.0

``` go
func MatchAll(pargs ...PositionalArgs) PositionalArgs
```

​	MatchAll 函数允许将多个 PositionalArgs 结合在一起协同工作。

#### func MaximumNArgs 

``` go
func MaximumNArgs(n int) PositionalArgs
```

​	MaximumNArgs 函数如果参数数量超过 n，则返回错误。

#### func MinimumNArgs 

``` go
func MinimumNArgs(n int) PositionalArgs
```

​	MinimumNArgs 函数如果参数数量不足 n，则返回错误。

#### func RangeArgs 

``` go
func RangeArgs(min int, max int) PositionalArgs
```

​	RangeArgs 函数如果参数数量不在预期范围内，则返回错误。

#### type ShellCompDirective  <- v1.0.0

``` go
type ShellCompDirective int
```

ShellCompDirective is a bit map representing the different behaviors the shell can be instructed to have once completions have been provided.

​	`ShellCompDirective` 是一个位图（bit map），表示在提供完成后，可以向 shell 发出指令以确定不同的行为方式。

``` go
const (	
    // ShellCompDirectiveError 指示发生错误，应忽略完成。
	ShellCompDirectiveError ShellCompDirective = 1 << iota

	// ShellCompDirectiveNoSpace indicates that the shell should not add a space
	// after the completion even if there is a single completion provided.
    // ShellCompDirectiveNoSpace 指示即使提供了单个完成，shell 也不应在完成后添加空格。
	ShellCompDirectiveNoSpace
	
    // ShellCompDirectiveNoFileComp 指示即使没有提供完成，shell 也不应提供文件完成。
	ShellCompDirectiveNoFileComp
	
    // ShellCompDirectiveFilterFileExt 指示所提供的完成应用作文件扩展名过滤器。
	// 对于标志，使用 Command.MarkFlagFilename()方法 和 Command.MarkPersistentFlagFilename()方法
	// 是使用此指令的快捷方式。常量 BashCompFilenameExt 注解也可以用于为标志获得相同的行为。
	ShellCompDirectiveFilterFileExt

    // ShellCompDirectiveFilterDirs 指示只提供目录名称以进行文件完成。
	// 要在另一个目录中请求目录名称，返回的完成应指定要搜索的目录。
	// 常量BashCompSubdirsInDir 注解也可以用于为标志获得相同的行为，但仅限于目录。
	ShellCompDirectiveFilterDirs
	
    // ShellCompDirectiveKeepOrder 指示 shell 应保留提供完成的顺序
	ShellCompDirectiveKeepOrder
	
    // ShellCompDirectiveDefault 指示在提供完成后让 shell 执行其默认行为。
	// 必须将此项放在最后，以避免混淆 iota 计数。
	ShellCompDirectiveDefault ShellCompDirective = 0
)
```

#### func NoFileCompletions  <- v1.2.0

``` go
func NoFileCompletions(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)
```

NoFileCompletions can be used to disable file completion for commands that should not trigger file completions.

​	NoFileCompletions 函数可用于禁用不应触发文件完成的命令的文件完成。
