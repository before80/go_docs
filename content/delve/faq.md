+++
title = "常见问题解答"
date = 2024-12-09T08:09:20+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/faq.md](https://github.com/go-delve/delve/blob/master/Documentation/faq.md)
>
> 收录该文档时间： `2024-12-09T08:09:20+08:00`

## 常见问题解答 Frequently Asked Questions

### 编译 Delve 时遇到错误 / 不支持的架构和操作系统 I'm getting an error while compiling Delve / unsupported architectures and OSs



The most likely cause of this is that you are running an unsupported Operating System or architecture. Currently Delve supports (GOOS / GOARCH):

​	最可能的原因是您正在使用一个不支持的操作系统或架构。目前，Delve 支持以下操作系统（GOOS）/架构（GOARCH）：

- linux / amd64 (86x64)
- linux / arm64 (AARCH64)
- linux / 386
- windows / amd64
- darwin (macOS) / amd64

There is no planned ETA for support of other architectures or operating systems. Bugs tracking requested support are:

​	目前没有计划支持其他架构或操作系统。有关请求支持的错误追踪：

- [32bit ARM support](https://github.com/go-delve/delve/issues/328)
- [PowerPC support](https://github.com/go-delve/delve/issues/1564)
- [OpenBSD](https://github.com/go-delve/delve/issues/1477)

See also: [backend test health](https://github.com/go-delve/delve/blob/master/Documentation/backend_test_health.md).

​	另请参阅：[后端测试健康状况](https://github.com/go-delve/delve/blob/master/Documentation/backend_test_health.md)。

### 如何在 Docker 中使用 Delve？ How do I use Delve with Docker?



When running the container you should pass the `--security-opt=seccomp:unconfined` option to Docker. You can start a headless instance of Delve inside the container like this:

​	运行容器时，您应该向 Docker 传递 `--security-opt=seccomp:unconfined` 选项。您可以通过以下方式在容器中启动一个无头实例的 Delve：

```
dlv exec --headless --listen :4040 /path/to/executable
```



And then connect to it from outside the container:

​	然后从容器外部连接到它：

```
dlv connect :4040
```



The program will not start executing until you connect to Delve and send the `continue` command. If you want the program to start immediately you can do that by passing the `--continue` and `--accept-multiclient` options to Delve:

​	程序不会在连接到 Delve 并发送 `continue` 命令之前开始执行。如果您希望程序立即开始，可以通过传递 `--continue` 和 `--accept-multiclient` 选项来启动 Delve：

```
dlv exec --headless --continue --listen :4040 --accept-multiclient /path/to/executable
```



Note that the connection to Delve is unauthenticated and will allow arbitrary remote code execution: *do not do this in production*.

​	请注意，连接到 Delve 是不经过身份验证的，这将允许任意的远程代码执行：*不要在生产环境中这样做*。

### 如何使用 Delve 调试 CLI 应用程序？ How can I use Delve to debug a CLI application?



There are three good ways to go about this

​	有三种有效的方式可以做到这一点：

1. Run your CLI application in a separate terminal and then attach to it via `dlv attach`. 在单独的终端中运行您的 CLI 应用程序，然后通过 `dlv attach` 附加到它。
2. Run Delve in headless mode via `dlv debug --headless` and then connect to it from another terminal. This will place the process in the foreground and allow it to access the terminal TTY. 通过 `dlv debug --headless` 在无头模式下运行 Delve，然后从另一个终端连接到它。这将把进程放到前台并允许它访问终端 TTY。
3. Assign the process its own TTY. This can be done on UNIX systems via the `--tty` flag for the `dlv debug` and `dlv exec` commands. For the best experience, you should create your own PTY and assign it as the TTY. This can be done via [ptyme](https://github.com/derekparker/ptyme). 为进程分配一个独立的 TTY。可以通过 UNIX 系统上的 `--tty` 标志为 `dlv debug` 和 `dlv exec` 命令完成。为了获得最佳体验，您应该创建自己的 PTY 并将其指定为 TTY。可以通过 [ptyme](https://github.com/derekparker/ptyme) 来实现。

### 如何使用 Delve 进行远程调试？ How can I use Delve for remote debugging?



It is best not to use remote debugging on a public network. If you have to do this, we recommend using ssh tunnels or a vpn connection.

​	最好不要在公共网络上使用远程调试。如果必须这样做，建议使用 ssh 隧道或 VPN 连接。

##### `Example `



Remote server:

​	远程服务器：

```
dlv exec --headless --listen localhost:4040 /path/to/executable
```



Local client:

​	本地客户端：

1. connect to the server and start a local port forward 连接到服务器并启动本地端口转发

```
ssh -NL 4040:localhost:4040 user@remote.ip
```



2. connect local port  连接本地端口

```
dlv connect :4040
```



### 无法在复杂的调试环境中设置断点或查看源代码列表 Can not set breakpoints or see source listing in a complicated debugging environment



This problem manifests when one or more of these things happen:

​	当发生以下情况时，会出现此问题：

- Can not see source code when the program stops at a breakpoint
  - 程序在断点处停止时无法看到源代码

- Setting a breakpoint using full path, or through an IDE, does not work
  - 使用完整路径或通过 IDE 设置断点无效


While doing one of the following things:

​	发生上述情况时：

- **The program is built and run inside a container** and Delve (or an IDE) is remotely connecting to it
  - **程序在容器内构建和运行**，并且 Delve（或 IDE）正在远程连接到它

- Generally, every time the build environment (VM, container, computer...) differs from the environment where Delve's front-end (dlv or a IDE) runs
  - 通常，每当构建环境（虚拟机、容器、计算机等）与 Delve 前端（dlv 或 IDE）运行的环境不同

- Using `-trimpath` or `-gcflags=-trimpath`
  - 使用 `-trimpath` 或 `-gcflags=-trimpath`

- Using a build system other than `go build` (eg. bazel)
  - 使用 `go build` 以外的构建系统（例如 bazel）

- Using symlinks in your source tree
  - 在源代码树中使用符号链接


If you are affected by this problem then the `list main.main` command (in the command line interface) will have this result:

​	如果您遇到此问题，则在命令行界面执行 `list main.main` 命令时，将出现以下结果：

```
(dlv) list main.main
Showing /path/to/the/mainfile.go:42 (PC: 0x47dfca)
Command failed: open /path/to/the/mainfile.go: no such file or directory
(dlv)
```



This is not a bug. The Go compiler embeds the paths of source files into the executable so that debuggers, including Delve, can use them. Doing any of the things listed above will prevent this feature from working seamlessly.

​	这不是一个错误。Go 编译器将源文件的路径嵌入到可执行文件中，以便包括 Delve 在内的调试器可以使用这些路径。执行上述任何操作都将阻止此功能顺利工作。

The substitute-path feature can be used to solve this problem, see `help config` or the `substitutePath` option in launch.json.

​	可以使用 `substitute-path` 功能来解决此问题，查看 `help config` 或 `launch.json` 中的 `substitutePath` 选项。

The `sources` command could also be useful in troubleshooting this problem, it shows the list of file paths that has been embedded by the compiler into the executable.

​	`sources` 命令也有助于排查此问题，它显示了编译器将文件路径嵌入到可执行文件中的列表。	

For more information on path substitution see [path substitution](https://github.com/go-delve/delve/blob/master/Documentation/cli/substitutepath.md).

​	有关路径替换的更多信息，请参见 [路径替换](https://github.com/go-delve/delve/blob/master/Documentation/cli/substitutepath.md)。

If you still think this is a bug in Delve and not a configuration problem, open an [issue](https://github.com/go-delve/delve/issues), filling the issue template and including the logs produced by delve with the options `--log --log-output=rpc,dap`.

​	如果您仍然认为这是 Delve 的 bug 而不是配置问题，请打开一个 [问题](https://github.com/go-delve/delve/issues)，填写问题模板，并包含 Delve 使用 `--log --log-output=rpc,dap` 选项产生的日志。

### 使用 Delve 调试 Go 运行时 Using Delve to debug the Go runtime



It's possible to use Delve to debug the Go runtime, however there are some caveats to keep in mind

​	可以使用 Delve 调试 Go 运行时，但需要注意以下几点：

- The `runtime` package is always compiled with optimizations and inlining, all of the caveats that apply to debugging optimized binaries apply to the runtime package. In particular some variables could be unavailable or have stale values and it could expose some bugs with the compiler assigning line numbers to instructions.
  - `runtime` 包始终会进行优化和内联，因此调试优化后的二进制文件时的所有限制都适用于 `runtime` 包。特别是，某些变量可能不可用或有过时的值，并且它可能会暴露编译器将行号分配给指令的错误。

- Next, step and stepout try to follow the current goroutine, if you debug one of the functions in the runtime that modify the curg pointer they will get confused. The 'step-instruction' command should be used instead.
  - 接下来的 `next`、`step` 和 `stepout` 尝试跟踪当前 goroutine，如果您调试运行时的某个函数，该函数修改 `curg` 指针，它们会变得混乱。应该使用 `step-instruction` 命令。
- When executing a stacktrace from g0 Delve will return the top frame and then immediately switch to the goroutine stack. If you want to see the g0 stacktrace use `stack -mode simple`.
  - 当执行 `g0` 的栈追踪时，Delve 会返回顶层栈帧，然后立即切换到 goroutine 栈。如果您想查看 `g0` 的栈追踪，可以使用 `stack -mode simple`。
- The step command only steps into private runtime functions if it is already inside a runtime function. To step inside a private runtime function inserted into user code by the compiler set a breakpoint and then use `runtime.curg.goid == <current goroutine id>` as condition.
  - `step` 命令只有在已经进入运行时函数时才会进入私有运行时函数。要进入编译器将私有运行时函数插入到用户代码中的运行时函数，请设置断点，然后使用 `runtime.curg.goid == <current goroutine id>` 作为条件。
