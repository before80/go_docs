+++
title = "路径替换"
date = 2024-12-09T08:00:12+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/cli/substitutepath.md](https://github.com/go-delve/delve/blob/master/Documentation/cli/substitutepath.md)
>
> 收录该文档时间： `2024-12-09T08:00:12+08:00`

## 路径替换配置 Path substitution configuration



Normally Delve finds the path to the source code that was used to produce an executable by looking at the debug symbols of the executable. However, under [some circumstances](https://github.com/go-delve/delve/blob/master/Documentation/faq.md#substpath), the paths that end up inside the executable will be different from the paths to the source code on the machine that is running the debugger. If that is the case Delve will need extra configuration to convert the paths stored inside the executable to paths in your local filesystem.

​	通常，Delve 通过查看可执行文件的调试符号来查找用于生成可执行文件的源代码路径。然而，在[某些情况下](https://github.com/go-delve/delve/blob/master/Documentation/faq.md#substpath)，可执行文件内部的路径与运行调试器的机器上的源代码路径不同。如果是这种情况，Delve 需要额外的配置，将可执行文件内存储的路径转换为本地文件系统中的路径。

This configuration is done by specifying a list of path substitution rules.

​	此配置通过指定路径替换规则的列表来完成。

### 路径替换规则的指定位置 Where are path substitution rules specified



#### Delve 命令行客户端 Delve command line client



The command line client reads the path substitution rules from Delve's YAML configuration file located at `$XDG_CONFIG_HOME/dlv/config.yml` or `.dlv/config.yml` inside the home directory on Windows.

​	命令行客户端从 Delve 的 YAML 配置文件中读取路径替换规则，该文件位于 `$XDG_CONFIG_HOME/dlv/config.yml`，或者在 Windows 上位于主目录中的 `.dlv/config.yml`。

The `substitute-path` entry should look like this:

​	`substitute-path` 条目应该像这样：

```
substitute-path:
  - {from: "/compiler/machine/directory", to: "/debugger/machine/directory"}
  - {from: "", to: "/mapping/for/relative/paths"}
```



If you are starting a headless instance of Delve and connecting to it through `dlv connect` the configuration file that is used is the one that runs `dlv connect`.

​	如果您启动了一个无头实例的 Delve 并通过 `dlv connect` 连接它，则使用的配置文件是运行 `dlv connect` 的那个配置文件。

The rules can also be modified while Delve is running by using the [config substitute-path command](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md#config):

​	路径替换规则也可以在 Delve 运行时通过使用 [config substitute-path 命令](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md#config)进行修改：

```
(dlv) config substitute-path /from/path /to/path
```



Double quotes can be used to specify paths that contain spaces, or to specify empty paths:

​	可以使用双引号指定包含空格的路径，或指定空路径：

```
(dlv) config substitute-path "/path containing spaces/" /path-without-spaces/
(dlv) config substitute-path /make/this/path/relative ""
```



#### DAP server



If you connect to Delve using the DAP protocol then the substitute path rules are specified using the substitutePath option in [launch.json](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#launchjson-attributes).

​	如果通过 DAP 协议连接到 Delve，则使用 `launch.json` 中的 `substitutePath` 选项指定路径替换规则。[launch.json](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#launchjson-attributes) 示例：

```
	"substitutePath": [
		{ "from": "/from/path", "to": "/to/path" }
	]
```



The [debug console](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#dlv-command-from-debug-console) can also be used to modify the path substitution list:

​	[调试控制台](https://github.com/golang/vscode-go/blob/master/docs/debugging.md#dlv-command-from-debug-console) 也可以用来修改路径替换列表：

```
dlv config substitutePath /from/path /to/path
```



This command works similarly to the `config substitute-path` command described above.

​	此命令的工作方式类似于上面描述的 `config substitute-path` 命令。

### 路径替换规则的应用方式 How are path substitution rules applied



Regardless of how they are specified the path substitution rules are an ordered list of `(from-path, to-path)` pairs. When Delve needs to convert a path P found inside the executable file into a path in the local filesystem it will scan through the list of rules looking for the first one where P starts with from-path and replace from-path with to-path.

​	无论路径替换规则如何指定，它们都是 `(from-path, to-path)` 对的有序列表。当 Delve 需要将可执行文件中找到的路径 P 转换为本地文件系统中的路径时，它会扫描规则列表，查找第一个 P 以 `from-path` 开头的规则，并用 `to-path` 替换 `from-path`。

Empty paths in both from-path and to-path are special, they represent relative paths:

​	`from-path` 和 `to-path` 中的空路径是特殊的，它们表示相对路径：

- `(from="" to="/home/user/project/src")` converts all relative paths in the executable to absolute paths in `/home/user/project/src`
  - `(from="" to="/home/user/project/src")` 将可执行文件中的所有相对路径转换为 `/home/user/project/src` 中的绝对路径

- `(from="/build/dir" to="")` converts all paths in the executable that start with `/build/dir` into relative paths.
  - `(from="/build/dir" to="")` 将可执行文件中以 `/build/dir` 开头的所有路径转换为相对路径


The path substitution code is SubstitutePath in pkg/locspec/locations.go.

​	路径替换代码位于 `pkg/locspec/locations.go` 中的 `SubstitutePath`。

