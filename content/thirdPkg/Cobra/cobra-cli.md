+++
title = "cobra-cli 之 Cobra Generator"
date = 2023-08-18T21:30:53+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Cobra Generator - Cobra 生成器

> 原文：[https://github.com/spf13/cobra-cli](https://github.com/spf13/cobra-cli)



​	Cobra 提供了自己的程序，可以为你创建应用程序并添加任何你想要的命令。这是将 Cobra 集成到你的应用程序中的最简单方法。

​	使用以下命令安装 cobra 生成器：

```bash
go install github.com/spf13/cobra-cli@latest
```

​	Go 将自动将其安装在你的 `$GOPATH/bin` 目录中，该目录应该在你的 `$PATH` 环境变量中。

​	安装完成后，你应该可以使用 `cobra-cli` 命令。在命令行中输入 `cobra-cli` 来确认。

​	目前 Cobra 生成器只支持两个操作：

## 操作1：cobra-cli init

​	`cobra-cli init [app]` 命令将为你创建初始的应用程序代码。这是一个非常强大的应用程序，它会为你的程序填充正确的结构，以便你可以立即享受 Cobra 的所有优势。它还可以将你指定的许可证应用于你的应用程序。

​	随着 Go 模块的引入，Cobra 生成器已经简化，以利用模块的优势。Cobra 生成器在 Go 模块内部工作。

### 初始化模块

​	**如果你已经有一个模块，请跳过此步骤。**

​	如果你想要初始化一个新的 Go 模块： 

1. 创建一个新的目录
2. 使用 `cd` 进入该目录
3. 运行 `go mod init <MODNAME>`

例如：

```
cd $HOME/code 
mkdir myapp
cd myapp
go mod init github.com/spf13/myapp
```



#### 初始化一个 Cobra CLI 应用程序

​	在 Go 模块内部运行 `cobra-cli init`。这将为你创建一个新的基本项目，供你编辑。

> 个人注释
>
> ​	截止到2023.8.18来看，目前`cobra-cli init ` 在执行的时候，需要在一个有go.mod文件的目录下，才能执行！但生成的目录结构可能对你来说多了一层目录，目前你可以手动将生成的内层文件手动剪切到外层目录！

​	你应该可以立即运行你的新应用程序。尝试使用 `go run main.go`。

​	你需要打开并编辑 'cmd/root.go'，并提供你自己的描述和逻辑。

例如：

```
cd $HOME/code/myapp
cobra-cli init
go run main.go
```

​	`cobra-cli init` 也可以从子目录中运行，就像 [cobra 生成器本身的组织方式](https://github.com/spf13/cobra-cli) 一样。如果你想将应用程序代码与库代码分开，这是非常有用的。

### 可选标志：

​	你可以使用 `--author` 标志提供作者名称。例如：`cobra-cli init --author "Steve Francia spf@spf13.com"`

​	你可以使用 `--license` 标志提供许可证。例如：`cobra-cli init --license apache`

​	使用 `--viper` 标志自动设置 [viper](https://github.com/spf13/viper)

​	Viper 是 Cobra 的伴侣，旨在轻松处理环境变量和配置文件，并将它们与应用程序标志无缝连接。

## 操作2：向项目添加命令

​	一旦初始化了 Cobra 应用程序，你可以继续使用 Cobra 生成器向你的应用程序添加其他命令。执行此操作的命令是 `cobra-cli add`。

​	假设你创建了一个应用程序，并希望为其添加以下命令： 

- app serve
- app config
- app config create

​	在项目目录（包含你的 main.go 文件的目录）中，你将运行以下命令：

```
cobra-cli add serve
cobra-cli add config
cobra-cli add create -p 'configCmd'
```

​	`cobra-cli add` 支持与 `cobra-cli init` 相同的可选标志（如上所述）。

​	你会注意到最后一个命令有一个 `-p` 标志。这用于将新添加的命令分配给一个父命令。在这种情况下，我们希望将 "create" 命令分配给 "config" 命令。如果未指定，默认情况下所有命令的父命令是 `rootCmd`。

​	默认情况下，`cobra-cli` 将名称后附加 `Cmd` 并将此名称用作内部变量名。在指定父命令时，请确保与代码中使用的变量名匹配。

​	*注意：对于命令名称，请使用 camelCase（而不是 snake_case/kebab-case）。否则，你会遇到错误。例如，`cobra-cli add add-user` 是不正确的，但 `cobra-cli add addUser` 是有效的。*

​	运行这三个命令后，你的应用程序结构将类似于以下内容：

```
  ▾ app/
    ▾ cmd/
        config.go
        create.go
        serve.go
        root.go
      main.go
```

​	在此阶段，你可以运行 `go run main.go`，它将运行你的应用程序。`go run main.go serve`、`go run main.go config`、`go run main.go config create` 以及 `go run main.go help serve` 等都将起作用。

​	现在你已经拥有一个基于 Cobra 的基本应用程序了。下一步是编辑 cmd 中的文件，并根据你的应用程序进行自定义。

​	有关使用 Cobra 库的完整详细信息，请阅读 [Cobra 用户指南](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#using-the-cobra-library)。

​	玩得开心！

## 配置 Cobra 生成器

​	如果你提供一个简单的配置文件，Cobra 生成器将更容易使用，这将帮助你在 flags 中消除重复提供大量信息的问题。

​	一个示例的 `~/.cobra.yaml` 文件：

```yml
author: Steve Francia <spf@spf13.com>
license: MIT
useViper: true
```

​	你还可以使用内置的许可证。例如，**GPLv2**、**GPLv3**、**LGPL**、**AGPL**、**MIT**、**2-Clause BSD** 或 **3-Clause BSD**。

​	你可以通过将 `license` 设置为 `none` 来指定没有许可证，或者你可以指定一个自定义许可证：

```yml
author: Steve Francia <spf@spf13.com>
year: 2020
license:
  header: This file is part of CLI application foo.
  text: |
    {{ .copyright }}

    This is my license. There are many like it, but this one is mine.
    My license is my best friend. It is my life. I must master it as I must
    master my life.
```

​	在上述自定义许可证配置中，许可证文本中的 `copyright` 行是从 `author` 和 `year` 属性生成的。`LICENSE` 文件的内容如下：

```
Copyright © 2020 Steve Francia <spf@spf13.com>

This is my license. There are many like it, but this one is mine.
My license is my best friend. It is my life. I must master it as I must
master my life.
```

​	`header` 属性用作许可证头文件。不执行插值。这是 go 文件头文件的示例。

```
/*
Copyright © 2020 Steve Francia <spf@spf13.com>
This file is part of CLI application foo.
*/
```