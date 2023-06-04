+++
title = "Cobra"
type = "docs"
date = 2023-05-21T16:24:15+08:00
weight = 1
description = ""
isCJKLanguage = true
draft = false
+++

# Cobra

> 原文：[https://cobra.dev/](https://cobra.dev/)

本文额外提及：

1. [pflag库](https://github.com/spf13/pflag)
2. [viper](https://github.com/spf13/viper) 
3. [hugo](https://gohugo.io/)



## 关于

{{< figure src="CobraDoc_img/logo.png" class="center" >}}

​	Cobra是Go语言的CLI框架。它包含一个用于创建强大现代CLI应用程序的库，以及一个用于快速生成基于Cobra的应用程序和命令文件的工具。

​	它由Go团队成员[spf13](https://twitter.com/spf13)为[hugo](https://gohugo.io/)创建，并已被最受欢迎的Go项目采用。

### Cobra 提供了

- 简单基于子命令的CLI：`app server`，`app fetch`等。
- 完全符合POSIX标准的标志（包括短标志和长标志）
- 嵌套子命令
- 全局、本地和级联标志
- 使用`cobra init appname`和`cobra add cmdname`轻松生成应用程序和命令
- 智能提示（`app srver`... 是否意味着 `app server`？）
- 为命令和标志自动生成帮助
- 自动识别帮助标志如`-h`，`--help`等
- 为您的应用程序自动生成的bash自动完成
- 为您的应用程序自动生成man页
- 命令别名，可以在不中断命令的情况下更改命令
- 灵活定义自己的帮助、用法等
- 可选与[viper](http://github.com/spf13/viper)紧密集成，用于[12因子应用程序](https://12factor.net/zh_cn/)

## 安装

​	使用Cobra非常简单。首先，使用`go get`安装最新版本的库。此命令将安装`cobra`生成器的可执行文件以及库及其依赖项：

```
go get -u github.com/spf13/cobra/cobra
```

然后，在您的应用程序中导入Cobra：

```go
import "github.com/spf13/cobra"
```



## 概念

​	Cobra是构建在命令、实参和标志的结构上。

​	**命令**代表操作，**参数**是事物，**标志**是对这些操作的修饰符。

​	最好的应用程序在使用时读起来像句子。用户将知道如何使用该应用程序，因为他们本能地了解如何使用它。

​	遵循的模式是 `APPNAME VERB NOUN --ADJECTIVE.` 或者 `APPNAME COMMAND ARG --FLAG`

​	一些好的真实世界的例子可能更好地说明这一点。

​	在下面的例子中，'server'是一个命令，'port'是一个标志：

```
hugo server --port=1313
```

​	在这个命令中，我们告诉Git（使用--bare，即裸仓库的形式，没有工作目录的方式，克隆之后只有一个以.git结尾的文件夹）克隆给定的url。

```
git clone URL --bare
```



## 命令

​	命令是应用程序的核心。应用程序支持的每个交互都将包含在一个命令中。一个命令可以有子命令，并可选择运行一个动作。

​	在上面的例子中，'server'是一个命令。

​	[更多关于cobra.Command的信息](https://pkg.go.dev/github.com/spf13/cobra?utm_source=godoc#Command)

## 标志

​	标志是修改命令行为的一种方式。Cobra支持完全符合POSIX标准的标志，以及Go的[flag包]({{< ref "/stdLib/flag" >}})。Cobra命令可以定义持续到子命令的标志，以及仅对该命令可用的标志。

​	在上面的例子中，'port'是一个标志。

​	标志功能由[pflag库](https://github.com/spf13/pflag)提供，它是flag标准库的一个分支，保持相同的接口同时添加POSIX兼容性。

## 入门指南

​	虽然您可以提供自己的组织结构，但通常基于Cobra的应用程序将遵循以下组织结构：

```
  ▾ appName/
    ▾ cmd/
        add.go
        your.go
        commands.go
        here.go
      main.go
```

​	在Cobra应用程序中，通常main.go文件非常简洁。它只有一个目的：初始化Cobra。

```go
package main

import (
  "{pathToYourApp}/cmd"
)

func main() {
  cmd.Execute()
}
```

## 使用Cobra生成器

​	Cobra提供了一个程序，可以创建您的应用程序并添加您想要的任何命令。这是将Cobra集成到您的应用程序中最简单的方法。

​	在[这里](https://github.com/spf13/cobra/blob/master/cobra/README.md)您可以找到更多关于Cobra的信息

## 使用Cobra库

​	要手动实现Cobra，您需要创建一个简单的main.go文件和一个rootCmd文件。您可以根据需要提供其他命令。

### 创建rootCmd

​	Cobra不需要任何特殊的构造函数。只需创建您的命令即可。

​	理想情况下，将其放置在 `app/cmd/root.go` 中：

```go
var rootCmd = &cobra.Command{
  Use:   "hugo",
  Short: "Hugo is a very fast static site generator",
  Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
```

​	您还可以在 `init()` 函数中定义标志和处理配置。

​	例如 cmd/root.go：

```go
import (
  "fmt"
  "os"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

func init() {
  cobra.OnInitialize(initConfig)
  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
  rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
  rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
  rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
  rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
  viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
  viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
  viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
  viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
  viper.SetDefault("license", "apache")
}

func initConfig() {
  // Don't forget to read config either from cfgFile or from home directory!
  // 不要忘记从cfgFile或家目录读取配置！
  if cfgFile != "" {
    // Use config file from the flag.
    // 使用配置文件中的标志。
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    // 查找家目录。
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".cobra" (without extension).
    // 在家目录中搜索名为".cobra"的配置文件（不包括扩展名）。
    viper.AddConfigPath(home)
    viper.SetConfigName(".cobra")
  }

  if err := viper.ReadInConfig(); err != nil {
    fmt.Println("Can't read config:", err)
    os.Exit(1)
  }
}
```

### 创建您的main.go

With the root command you need to have your main function execute it. Execute should be run on the root for clarity, though it can be called on any command.

​	有了 root 命令，您需要在 main 函数执行它。为了清晰起见，应该在 root 上调用 `Execute`，尽管可以在任何命令上调用它。

​	在Cobra应用程序中，通常`main.go`文件非常简洁。它只有一个目的，即初始化Cobra。

```go
package main

import (
  "{pathToYourApp}/cmd"
)

func main() {
  cmd.Execute()
}
```

### 创建其他命令

​	可以定义其他命令，通常每个命令都有自己的文件，放在 `cmd/` 目录下。

​	如果您想创建一个版本命令，您可以创建 `cmd/version.go` 文件，并使用以下内容填充：

```go
package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Hugo",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}
```

## 使用标志 - Working with Flags 

​	标志提供修饰符，用于控制操作命令的行为。

### 将标志分配给命令

​	由于标志在不同的位置定义和使用，我们需要在外部定义一个具有正确作用域的变量，以便将标志分配给其工作的命令。

```go
var Verbose bool
var Source string
```

​	有两种不同的方法来分配标志。

### 持久标志 - Persistent Flags

​	一个标志可以是"持久的（persistent）"，这意味着该标志将可用于分配给它的命令以及该命令下的每个命令。对于全局标志，将标志分配为根上的持久标志。

```go
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
```

### 局部标志 - Local Flags

​	一个标志也可以被局部分配，这仅适用于特定的命令。

```go
rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
```

### 父命令上的局部标志 - Local Flag on Parent Commands

​	默认情况下，Cobra仅解析目标命令上的局部标志，忽略父命令上的任何局部标志。通过启用 `Command.TraverseChildren`，Cobra将在执行目标命令之前解析每个命令上的局部标志。

```go
command := cobra.Command{
  Use: "print [OPTIONS] [COMMANDS]",
  TraverseChildren: true,
}
```

### 将标志与配置绑定 - Bind Flags with Config

​	您还可以将标志与 [viper](https://github.com/spf13/viper) 绑定：

```go
var author string

func init() {
  rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
  viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}
```

​	在此示例中，持久标志 `author` 与 `viper` 绑定。**注意**，当用户未提供 `--author` 标志时，变量 `author` 的值不会设置为配置中的值。

​	有关更多信息，请参阅 [viper 文档](https://github.com/spf13/viper#working-with-flags)。

### 必需的标志 - Required flags

​	默认情况下，标志是可选的。如果您希望命令在标志未设置时报告错误，请将其标记为必需：

```go
rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
rootCmd.MarkFlagRequired("region")
```

## 位置和自定义实参 - Positional and Custom Arguments

​	可以使用 `Command` 的 `Args` 字段来指定位置参数的验证规则。

​	以下是内置的验证器：

- `NoArgs` - 如果存在任何位置参数，该命令将报告错误。
- `ArbitraryArgs` - 该命令将接受任何参数。
- `OnlyValidArgs` - 如果存在任何不在 `Command` 的 `ValidArgs` 字段中的位置参数，该命令将报告错误。
- `MinimumNArgs(int)` - 如果位置参数少于 N 个，该命令将报告错误。
- `MaximumNArgs(int)` - 如果位置参数多于 N 个，该命令将报告错误。
- `ExactArgs(int)` - 如果位置参数数量不是恰好为 N 个，该命令将报告错误。
- `ExactValidArgs(int)` - 如果位置参数数量不是恰好为 N 个或者存在任何不在 `Command` 的 `ValidArgs` 字段中的位置参数，该命令将报告错误。
- `RangeArgs(min, max)` - 如果参数数量不在期望的最小值和最大值之间，该命令将报告错误。

​	下面是设置自定义验证器的示例：

```go
var cmd = &cobra.Command{
  Short: "hello",
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("requires at least one arg")
    }
    if myapp.IsValidColor(args[0]) {
      return nil
    }
    return fmt.Errorf("invalid color specified: %s", args[0])
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello, World!")
  },
}
```

## 示例

​	在下面的示例中，我们定义了三个命令。两个命令位于顶级，而一个命令（cmdTimes）是其中一个顶级命令的子命令。在这种情况下，根命令不可执行，这意味着需要一个子命令。这是通过不为 `rootCmd` 提供 'Run' 来实现的。

​	我们只为单个命令定义了一个标志。

​	有关标志的更多文档可在 [https://github.com/spf13/pflag](https://github.com/spf13/pflag) 查看。

```go
package main

import (
  "fmt"
  "strings"

  "github.com/spf13/cobra"
)

func main() {
  var echoTimes int

  var cmdPrint = &cobra.Command{
    Use:   "print [string to print]",
    Short: "Print anything to the screen",
    Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Print: " + strings.Join(args, " "))
    },
  }

  var cmdEcho = &cobra.Command{
    Use:   "echo [string to echo]",
    Short: "Echo anything to the screen",
    Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Print: " + strings.Join(args, " "))
    },
  }

  var cmdTimes = &cobra.Command{
    Use:   "times [# times] [string to echo]",
    Short: "Echo anything to the screen more times",
    Long: `echo things multiple times back to the user by providing
a count and a string.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      for i := 0; i < echoTimes; i++ {
        fmt.Println("Echo: " + strings.Join(args, " "))
      }
    },
  }

  cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

  var rootCmd = &cobra.Command{Use: "app"}
  rootCmd.AddCommand(cmdPrint, cmdEcho)
  cmdEcho.AddCommand(cmdTimes)
  rootCmd.Execute()
}
```

​	要查看更完整的较大应用程序示例，请查看 [Hugo](http://gohugo.io/)。

## help命令

​	当您拥有子命令时，Cobra会自动为应用程序添加`help`命令。当用户运行 `app help` 时，将调用此命令。此外，`help`还支持所有其他命令作为输入。例如，假设您有一个名为 `create` 的命令而没有其他配置；当调用 `app help create` 时，Cobra 将起作用。每个命令都会自动添加 `–help` 标志。

### 示例

The following output is automatically generated by Cobra. Nothing beyond the command and flag definitions are needed.

​	下面的输出是由Cobra自动生成的。除了命令和标志定义外，不需要其他内容。

```bash
$ cobra help

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra [command]

Available Commands:
  add         Add a command to a Cobra Application
  help        Help about any command
  init        Initialize a Cobra Application

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cobra
  -l, --license string   name of license for the project
      --viper            use Viper for configuration (default true)

Use "cobra [command] --help" for more information about a command.
```

​	`help`命令只是像其他命令一样的一个命令。它没有特殊的逻辑或行为。事实上，如果您需要的话，可以提供自己的帮助命令。

### 定义自己的帮助命令

​	您可以使用以下函数提供自己的Help命令或默认命令的模板：

```go
cmd.SetHelpCommand(cmd *Command)
cmd.SetHelpFunc(f func(*Command, []string))
cmd.SetHelpTemplate(s string)
```

​	后两个函数也适用于任何子命令。

## 用法信息 - Usage Message

​	当用户提供无效的标志或无效的命令时，Cobra会显示'usage'以响应用户。

### 示例

​	您可能会从上面的帮助中看到这个。这是因为默认帮助将用法嵌入到输出中。

```bash
$ cobra --invalid
Error: unknown flag: --invalid
Usage:
  cobra [command]

Available Commands:
  add         Add a command to a Cobra Application
  help        Help about any command
  init        Initialize a Cobra Application

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cobra
  -l, --license string   name of license for the project
      --viper            use Viper for configuration (default true)

Use "cobra [command] --help" for more information about a command.
```

### 定义您自己的用法

​	您可以为Cobra提供自己的usage函数或模板。与`help`命令一样，可以通过公共方法覆盖该函数和模板：

```go
cmd.SetUsageFunc(f func(*Command) error)
cmd.SetUsageTemplate(s string)
```

## 版本标志 - Version Flag

​	如果根命令的 Version 字段已设置，Cobra会添加顶层的"`--version`"标志。使用"`--version`"标志运行应用程序将使用版本模板将版本打印到标准输出(stdout)。可以使用 `cmd.SetVersionTemplate(s string)` 函数自定义模板。

## PreRun和PostRung钩子 - PreRun and PostRun Hooks

It is possible to run functions before or after the main `Run` function of your command. The `PersistentPreRun` and `PreRun` functions will be executed before `Run`. `PersistentPostRun` and `PostRun` will be executed after `Run`. The `Persistent*Run` functions will be inherited by children if they do not declare their own. These functions are run in the following order:

​	您可以在命令的主要 `Run` 函数之前或之后运行函数。`PersistentPreRun` 和 `PreRun` 函数将在 `Run` 函数之前执行。`PersistentPostRun` 和 `PostRun` 将在 `Run` 之后执行。如果子命令没有声明自己的钩子函数，它们将继承 `Persistent*Run` 函数。这些函数按照以下顺序运行：

- `PersistentPreRun`
- `PreRun`
- `Run`
- `PostRun`
- `PersistentPostRun`

​	以下是使用所有这些特性的两个命令的示例。当执行子命令时，它将运行根命令的 `PersistentPreRun`，但不会运行根命令的 `PersistentPostRun`：

```go
package main

import (
  "fmt"

  "github.com/spf13/cobra"
)

func main() {

  var rootCmd = &cobra.Command{
    Use:   "root [sub]",
    Short: "My root command",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
    },
    PreRun: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
    },
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside rootCmd Run with args: %v\n", args)
    },
    PostRun: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
    },
    PersistentPostRun: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
    },
  }

  var subCmd = &cobra.Command{
    Use:   "sub [no options!]",
    Short: "My subcommand",
    PreRun: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
    },
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside subCmd Run with args: %v\n", args)
    },
    PostRun: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
    },
    PersistentPostRun: func(cmd *cobra.Command, args []string) {
      fmt.Printf("Inside subCmd PersistentPostRun with args: %v\n", args)
    },
  }

  rootCmd.AddCommand(subCmd)

  rootCmd.SetArgs([]string{""})
  rootCmd.Execute()
  fmt.Println()
  rootCmd.SetArgs([]string{"sub", "arg1", "arg2"})
  rootCmd.Execute()
}
```

输出结果：

```
Inside rootCmd PersistentPreRun with args: []
Inside rootCmd PreRun with args: []
Inside rootCmd Run with args: []
Inside rootCmd PostRun with args: []
Inside rootCmd PersistentPostRun with args: []

Inside rootCmd PersistentPreRun with args: [arg1 arg2]
Inside subCmd PreRun with args: [arg1 arg2]
Inside subCmd Run with args: [arg1 arg2]
Inside subCmd PostRun with args: [arg1 arg2]
Inside subCmd PersistentPostRun with args: [arg1 arg2]
```

## 发生"unknown command"时的建议

​	当出现"unknown command" 错误时，Cobra将自动提供建议。这使得Cobra在发生拼写错误时的行为类似于`git`命令。例如：

```
$ hugo srever
Error: unknown command "srever" for "hugo"

Did you mean this?
        server

Run 'hugo --help' for usage.
```

​	建议是基于每个已注册的子命令，并使用[Levenshtein distance](http://en.wikipedia.org/wiki/Levenshtein_distance)的实现。所有已注册的命令中，与最小距离为2（忽略大小写）的命令将作为建议显示。

​	如果您需要禁用建议或调整命令中的字符串距离，可以使用：

```go
command.DisableSuggestions = true
```

或者

```go
command.SuggestionsMinimumDistance = 1
```

​	您还可以使用 `SuggestFor` 属性显式设置一个给定命令的建议的名称。这样可以对于在字符串距离上不接近但在命令集合中有意义的字符串进行建议，而且您不希望为它们设置别名。示例：

```
$ kubectl remove
Error: unknown command "remove" for "kubectl"

Did you mean this?
        delete

Run 'kubectl help' for usage.
```

## 为您的命令生成文档

​	Cobra可以基于子命令、标志等生成文档，支持以下格式：

- [Markdown](https://cobra.dev/doc/md_docs.md)
- [ReStructured Text](https://cobra.dev/doc/rest_docs.md)
- [Man Page](https://cobra.dev/doc/man_docs.md)

## 生成bash自动补全

​	Cobra可以生成bash自动补全文件。如果您为命令添加更多信息，这些自动补全功能将变得非常强大和灵活。详细了解请阅读[Bash Completions](https://cobra.dev/bash_completions.md)。

## 许可证

​	Cobra使用Apache 2.0许可证发布。请参阅[LICENSE.txt](https://github.com/spf13/cobra/blob/master/LICENSE.txt)。



