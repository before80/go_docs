+++
title = "如何编写 Delve 客户端——非正式指南"
date = 2024-12-09T07:57:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/api/ClientHowto.md](https://github.com/go-delve/delve/blob/master/Documentation/api/ClientHowto.md)
>
> 收录该文档时间： `2024-12-09T07:57:43+08:00`

# How to write a Delve client, an informal guide - 如何编写 Delve 客户端——非正式指南



## 启动后端 Spawning the backend



The `dlv` binary built by our `Makefile` contains both the backend and a simple command line client. If you are writing your own client you will probably want to run only the backend, you can do this by specifying the `--headless` option, for example:

​	由 `Makefile` 构建的 `dlv` 二进制文件包含了后端和一个简单的命令行客户端。如果你正在编写自己的客户端，可能只想运行后端，可以通过指定 `--headless` 选项来实现，例如：

```
$ dlv --headless debug
```



The rest of the command line remains unchanged. You can use `debug`, `exec`, `test`, etc... along with `--headless` and they will work. If this project is part of a larger IDE integration then you probably have your own build system and do not wish to offload this task to Delve, in that case it's perfectly fine to always use the `dlv exec` command but do remember that:

​	命令行的其余部分保持不变。你可以与 `--headless` 一起使用 `debug`、`exec`、`test` 等命令，它们都能正常工作。如果这个项目是一个更大的 IDE 集成的一部分，那么你可能有自己的构建系统，不希望将这项任务交给 Delve。在这种情况下，始终使用 `dlv exec` 命令也是完全可以的，但请记住：

1. Delve may not have all the information necessary to properly debug optimized binaries, so it is recommended to disable them via: `-gcflags='all=-N -l`. Delve 可能没有足够的信息来正确调试优化过的二进制文件，因此建议通过 `-gcflags='all=-N -l'` 禁用优化。
2. your users *do want* to debug their tests so you should also provide some way to build the test executable (equivalent to `go test -c --gcflags='all=-N -l'`) and pass it to Delve. 用户确实希望调试他们的测试，因此你也应该提供某种方式来构建测试可执行文件（等同于 `go test -c --gcflags='all=-N -l'`）并将其传递给 Delve。

It would also be nice for your users if you provided a way to attach to a running process, like `dlv attach` does.

​	如果你提供一种方法让用户附加到正在运行的进程上（像 `dlv attach` 那样），这对用户来说也是很好的体验。

Command line arguments that should be handed to the inferior process should be specified on dlv's command line after a "--" argument:

​	应该传递给子进程的命令行参数应在 `dlv` 的命令行中指定，位置在 `--` 参数后面：

```
dlv exec --headless ./somebinary -- these arguments are for the inferior process
```



Specifying a static port number, like in the [README](https://github.com/go-delve/delve/tree/master/Documentation/README.md) example, can be done using `--listen=127.0.0.1:portnumber`.

​	指定静态端口号，像在 [README](https://github.com/go-delve/delve/tree/master/Documentation/README.md) 示例中那样，可以使用 `--listen=127.0.0.1:端口号`。

This will, however, cause problems if you actually spawn multiple instances of the debugger.

​	然而，如果你实际上启动多个调试器实例，这会引发问题。

It's probably better to let Delve pick a random unused port number on its own. To do this do not specify any `--listen` option and read one line of output from dlv's stdout. If the first line emitted by dlv starts with "API server listening at: " then dlv started correctly and the rest of the line specifies the address that Delve is listening at.

​	更好的做法是让 Delve 自动选择一个未使用的随机端口号。为此，不要指定任何 `--listen` 选项，直接读取 `dlv` 的标准输出中的一行。如果 `dlv` 输出的第一行以 “API server listening at: ” 开头，那么 Delve 就已经正确启动了，后续的内容就是 Delve 正在监听的地址。

The `--log-dest` option can be used to redirect the "API server listening at:" message to a file or to a file descriptor. If the flag is not specified, the message will be output to stdout while other log messages are output to stderr.

​	可以使用 `--log-dest` 选项将 “API server listening at:” 信息重定向到文件或文件描述符。如果未指定该标志，消息将输出到标准输出，而其他日志信息则输出到标准错误。

## 控制后端 Controlling the backend



Once you have a running headless instance you can connect to it and start sending commands. Delve's protocol is built on top of the [JSON-RPC 1.0 specification](https://www.jsonrpc.org/specification_v1).

​	一旦你启动了一个无头实例，你就可以连接到它并开始发送命令。Delve 的协议基于 [JSON-RPC 1.0 规范](https://www.jsonrpc.org/specification_v1)。

The methods of a `service/rpc2.RPCServer` are exposed through this connection, to find out which requests you can send see the documentation of RPCServer on [Go Reference](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer).

​	`service/rpc2.RPCServer` 的方法通过这个连接暴露出来，想要了解你可以发送哪些请求，请参见 [Go Reference](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer) 中的 RPCServer 文档。

### Example



Let's say you are trying to create a breakpoint. By looking at [Go Reference](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer) you'll find that there is a `CreateBreakpoint` method in `RPCServer`.

​	假设你想要创建一个断点。通过查看 [Go Reference](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer)，你会发现 `RPCServer` 中有一个 `CreateBreakpoint` 方法。

This method, like all other methods of RPCServer that you can call through the API, has two arguments: `args` and `out`: `args` contains all the input arguments of `CreateBreakpoint`, while `out` is what `CreateBreakpoint` will return to you.

​	这个方法就像所有其他可以通过 API 调用的 `RPCServer` 方法一样，有两个参数：`args` 和 `out`：`args` 包含 `CreateBreakpoint` 的所有输入参数，而 `out` 是 `CreateBreakpoint` 返回给你的内容。

The call that you could want to make, in pseudo-code, would be:

​	你可能会想要调用的伪代码是：

```
RPCServer.CreateBreakpoint(CreateBreakpointIn{ File: "/User/you/some/file.go", Line: 16 })
```



To actually send this request on the JSON-RPC connection you just have to convert the CreateBreakpointIn object to json and then wrap everything into a JSON-RPC envelope:

​	为了实际通过 JSON-RPC 连接发送这个请求，你只需将 `CreateBreakpointIn` 对象转换为 JSON，然后将所有内容包装进一个 JSON-RPC 信封：

```
{"method":"RPCServer.CreateBreakpoint","params":[{"Breakpoint":{"file":"/User/you/some/file.go","line":16}}],"id":27}
```



Delve will respond by sending a response packet that will look like this:

​	Delve 将会发送一个响应包，内容如下：

```
{"id":27, "result": {"Breakpoint": {"id":3, "name":"", "addr":4538829, "file":"/User/you/some/file.go", "line":16, "functionName":"main.main", "Cond":"", "continue":false, "goroutine":false, "stacktrace":0, "LoadArgs":null, "LoadLocals":null, "hitCount":{}, "totalHitCount":0}}, "error":null}
```



## 选择 API 版本 Selecting the API version



Delve currently supports two version of its API, APIv1 and APIv2. By default a headless instance of `dlv` will serve APIv1 for backward-compatibility with older clients, however new clients should use APIv2 as new features will only be made available through version 2. The preferred method of switching to APIv2 is to send the `RPCServer.SetApiVersion` command right after connecting to the backend. Alternatively the `--api-version=2` command line option can be used when spawning the backend.

​	Delve 当前支持两个版本的 API，APIv1 和 APIv2。默认情况下，无头实例的 `dlv` 将为兼容旧客户端而提供 APIv1，但新的客户端应该使用 APIv2，因为新功能仅会通过版本 2 提供。切换到 APIv2 的首选方法是在连接到后端后立即发送 `RPCServer.SetApiVersion` 命令。或者，也可以在启动后端时使用 `--api-version=2` 命令行选项。

## 诊断 Diagnostics



Just like any other program, both Delve and your client have bugs. To help with determining where the problem is you should log the exchange of messages between Delve and your client somehow.

​	像任何其他程序一样，Delve 和你的客户端都有可能存在 bug。为了帮助确定问题所在，你应该以某种方式记录 Delve 和客户端之间的信息交换。

If you don't want to do this yourself you can also pass the options `--log --log-output=rpc` to Delve. In fact the `--log-output` has many useful values and you should expose it to users, if possible, so that we can diagnose problems that are hard to reproduce.

​	如果你不想自己做这件事，你也可以将 `--log --log-output=rpc` 选项传递给 Delve。实际上，`--log-output` 具有许多有用的值，你应该尽可能将其暴露给用户，以便我们能够诊断那些难以重现的问题。

## Using RPCServer.Command



`Command` is probably the most important API entry point. It lets your client stop (`Name == "halt"`) and resume (`Name == "continue"`) execution of the inferior process.

​	`Command` 可能是最重要的 API 入口点。它允许你的客户端暂停（`Name == "halt"`）并恢复（`Name == "continue"`）子进程的执行。

The return value of `Command` is a `DebuggerState` object. If you lose the DebuggerState object returned by your last call to `Command` you can ask for a new copy with `RPCServer.State`.

​	`Command` 的返回值是一个 `DebuggerState` 对象。如果你丢失了通过 `Command` 调用返回的 `DebuggerState` 对象，你可以通过 `RPCServer.State` 请求一个新的副本。

### 处理同时触发的断点 Dealing with simultaneous breakpoints



Since Go is a programming language with a big emphasis on concurrency and parallelism it's possible that multiple goroutines will stop at a breakpoint simultaneously. This may at first seem incredibly unlikely but you must understand that between the time a breakpoint is triggered and the point where the debugger finishes stopping all threads of the inferior process thousands of CPU instructions have to be executed, which make simultaneous breakpoint triggering not that unlikely.

​	由于 Go 是一门高度强调并发和并行的编程语言，因此多个 goroutine 可能会在同一个断点处同时停止。这看起来可能不太可能，但你必须理解，在一个断点被触发和调试器停止所有子进程线程之间，成千上万的 CPU 指令可能已经执行完毕，这使得同时触发断点并非不太可能。

You should signal to your user *all* the breakpoints that occur after executing a command, not just the first one. To do this iterate through the `Threads` array in `DebuggerState` and note all the threads that have a non nil `Breakpoint` member.

​	你应该在执行命令后通知用户所有触发的断点，而不仅仅是第一个。为此，可以遍历 `DebuggerState` 中的 `Threads` 数组，并记录所有具有非空 `Breakpoint` 成员的线程。

### 特殊的继续命令 Special continue commands



In addition to "halt" and vanilla "continue" `Command` offers a few extra flavours of continue that automatically set interesting temporary breakpoints: "next" will continue until the next line of the program, "stepout" will continue until the function returns, "step" is just like "next" but it will step into function calls (but skip all calls to unexported runtime functions).

​	除了“halt”和普通的“continue”，`Command` 还提供了几种额外的继续命令，这些命令会自动设置一些有趣的临时断点：“next” 会继续直到程序的下一行，“stepout” 会继续直到函数返回，“step” 就像“next”一样，但它会进入函数调用（但会跳过所有调用未导出的运行时函数）。

All of "next", "step" and "stepout" operate on the selected goroutine. The selected goroutine is described by the `SelectedGoroutine` field of `DebuggerState`. Every time `Command` returns the selected goroutine will be reset to the goroutine that triggered the breakpoint.

​	“next”、“step”和“stepout”都会操作选定的 goroutine。选定的 goroutine 由 `DebuggerState` 中的 `SelectedGoroutine` 字段描述。每次 `Command` 返回时，选定的 goroutine 都会被重置为触发断点的 goroutine。

If multiple breakpoints are triggered simultaneously the selected goroutine will be chosen randomly between the goroutines that are stopped at a breakpoint. If a breakpoint is hit by a thread that is executing on the system stack *there will be no selected goroutine*. If the "halt" command is called *there may not be a selected goroutine*.

​	如果同时触发了多个断点，选定的 goroutine 会在停止在断点上的 goroutine 之间随机选择。如果一个断点是由一个在系统栈上执行的线程触发的，那么*将没有选定的 goroutine*。如果调用了“halt”命令*可能没有选定的 goroutine*。

The selected goroutine can be changed using the "switchGoroutine" command. If "switchGoroutine" is used to switch to a goroutine that's currently parked SelectedGoroutine and CurrentThread will be mismatched. Always prefer SelectedGoroutine over CurrentThread, you should ignore CurrentThread entirely unless SelectedGoroutine is nil.

​	可以使用“switchGoroutine”命令来更改选定的 goroutine。如果使用“switchGoroutine”切换到当前停驻的 goroutine，则 `SelectedGoroutine` 和 `CurrentThread` 会不匹配。始终优先使用 `SelectedGoroutine`，除非它为 nil，否则你应该完全忽略 `CurrentThread`。

### 特殊的继续命令和异步断点 Special continue commands and asynchronous breakpoints



Because of the way go internals work it is not possible for a debugger to resume a single goroutine. Therefore it's possible that after executing a next/step/stepout a goroutine other than the goroutine the next/step/stepout was executed on will hit a breakpoint.

​	由于 Go 的内部工作方式，调试器无法恢复单独的 goroutine。因此，执行 `next`、`step` 或 `stepout` 后，可能会有一个与执行该命令的 goroutine 不同的 goroutine 触发断点。

If this happens Delve will return a DebuggerState with NextInProgress set to true. When this happens your client has two options:

​	如果发生这种情况，Delve 会返回一个 `DebuggerState`，并且 `NextInProgress` 会被设置为 `true`。在这种情况下，你的客户端有两个选择：

- You can signal that a different breakpoint was hit and then automatically attempt to complete the next/step/stepout by calling `RPCServer.Command` with `Name == "continue"`
  - 你可以通知用户触发了不同的断点，然后自动尝试通过调用 `RPCServer.Command` 并设置 `Name == "continue"` 来完成 `next`、`step` 或 `stepout` 操作。

- You can abort the next/step/stepout operation using `RPCServer.CancelNext`.
  - 你可以使用 `RPCServer.CancelNext` 中止 `next`、`step` 或 `stepout` 操作。


It is important to note that while NextInProgress is true it is not possible to call next/step/stepout again without using CancelNext first. There can not be multiple next/step/stepout operations in progress at any time.

​	重要的是要注意，当 `NextInProgress` 为 `true` 时，必须先使用 `CancelNext` 才能再次调用 `next`、`step` 或 `stepout`。在任何时候，不能同时进行多个 `next`、`step` 或 `stepout` 操作。

### RPCServer.Command 和陈旧的可执行文件 RPCServer.Command and stale executable files



It's possible (albeit unfortunate) that your user will decide to change the source of the program being executed in the debugger, while the debugger is running. Because of this it would be advisable that your client check that the executable is not stale every time `Command` returns and notify the user that the executable being run is stale and line numbers may nor align properly anymore.

​	有可能（虽然不太理想）用户会决定在调试器运行时更改正在调试的程序的源代码。因为这个原因，建议您的客户端每次 `Command` 返回时检查可执行文件是否已经过时，并通知用户正在运行的可执行文件已经过时，行号可能已经不再对齐。

You can do this bookkeeping yourself, but Delve can also help you with the `LastModified` call that returns the LastModified time of the executable file when Delve started it.

​	您可以自己处理这些记录，但 Delve 也可以通过 `LastModified` 调用帮助您，它会返回 Delve 启动时可执行文件的最后修改时间。

## Using RPCServer.CreateBreakpoint



The only two fields you probably want to fill of the Breakpoint argument of CreateBreakpoint are File and Line. The file name should be the absolute path to the file as the compiler saw it.

​	您可能只需要填写 `CreateBreakpoint` 的 `Breakpoint` 参数中的两个字段：`File` 和 `Line`。文件名应为编译器看到的文件的绝对路径。

For example if the compiler saw this path:

​	例如，如果编译器看到如下路径：

```
/Users/you/go/src/something/something.go
```



But `/Users/you/go/src/something` is a symbolic link to `/Users/you/projects/golang/something` the path *must* be specified as `/Users/you/go/src/something/something.go` and `/Users/you/projects/golang/something/something.go` will not be recognized as valid.

​	但 `/Users/you/go/src/something` 是 `/Users/you/projects/golang/something` 的符号链接，路径 *必须* 指定为 `/Users/you/go/src/something/something.go`，而 `/Users/you/projects/golang/something/something.go` 将不被识别为有效路径。

If you want to let your users specify a breakpoint on a function selected from a list of all functions you should specify the name of the function in the FunctionName field of Breakpoint.

​	如果您希望让用户从所有函数的列表中选择一个函数来指定断点，您应该在 `Breakpoint` 的 `FunctionName` 字段中指定该函数的名称。

If you want to support the [same language as dlv's break and trace commands](https://github.com/go-delve/delve/tree/master/Documentation/cli/locspec.md) you should call RPCServer.FindLocation and then use the returned slice of Location objects to create Breakpoints to pass to CreateBreakpoint: just fill each Breakpoint.Addr with the contents of the corresponding Location.PC.

​	如果您想支持与 `dlv` 的 `break` 和 `trace` 命令相同的语言，您应该调用 `RPCServer.FindLocation`，然后使用返回的 `Location` 对象切片来创建断点并传递给 `CreateBreakpoint`：只需将每个 `Breakpoint.Addr` 填充为相应 `Location.PC` 的内容。

## 查看变量 Looking into variables



There are several API entry points to evaluate variables in Delve:

​	Delve 提供了多个 API 入口来评估变量：

- RPCServer.ListPackageVars returns all global variables in all packages
  - `RPCServer.ListPackageVars` 返回所有包中的全局变量

- PRCServer.ListLocalVars returns all local variables of a stack frame
  - `RPCServer.ListLocalVars` 返回堆栈帧中的所有局部变量

- RPCServer.ListFunctionArgs returns all function arguments of a stack frame
  - `RPCServer.ListFunctionArgs` 返回堆栈帧中的所有函数参数

- RPCServer.Eval evaluates an expression on a given stack frame
  - `RPCServer.Eval` 在给定堆栈帧上评估一个表达式


All those API calls take a LoadConfig argument. The LoadConfig specifies how much of the variable's value should actually be loaded. Because of LoadConfig a variable could be loaded incompletely, you should always notify the user of this:

​	所有这些 API 调用都接受一个 `LoadConfig` 参数。`LoadConfig` 指定应该加载变量值的多少。由于 `LoadConfig` 的存在，变量可能会被不完全加载，因此您应该始终通知用户这一点：

- For strings, arrays, slices *and structs* the load is incomplete if: `Variable.Len > len(Variable.Children)`. This can happen to structs even if LoadConfig.MaxStructFields is -1 when MaxVariableRecurse is reached.

  - 对于字符串、数组、切片 *和结构体*，如果加载不完整：`Variable.Len > len(Variable.Children)`。即使 `LoadConfig.MaxStructFields` 为 -1，当达到 `MaxVariableRecurse` 时，结构体也可能出现这种情况。

- For maps the load is incomplete if: `Variable.Len > len(Variable.Children) / 2`

  - 对于映射，如果加载不完整：`Variable.Len > len(Variable.Children) / 2`

- For interfaces the load is incomplete if the only children has the onlyAddr attribute set to true.

  - 对于接口类型，如果唯一的子项设置了 `onlyAddr` 属性为 `true`，则加载不完整。

  

### 加载更多的变量 Loading more of a Variable



You can also give the user an option to continue loading an incompletely loaded variable. To load a struct that wasn't loaded automatically evaluate the expression returned by:

​	您还可以给用户提供一个选项，继续加载一个未完全加载的变量。要加载一个未自动加载的结构体，评估返回的表达式：

```
fmt.Sprintf("*(*%q)(%#x)", v.Type, v.Addr)
```



where v is the variable that was truncated.

​	其中 `v` 是被截断的变量。

To load more elements from an array, slice or string:

​	要加载数组、切片或字符串中的更多元素：

```
fmt.Sprintf("(*(*%q)(%#x))[%d:]", v.Type, v.Addr, len(v.Children))
```



To load more elements from a map:

​	要加载映射中的更多元素：

```
fmt.Sprintf("(*(*%q)(%#x))[%d:]", v.Type, v.Addr, len(v.Children)/2)
```



All the evaluation API calls except ListPackageVars also take a EvalScope argument, this specifies which stack frame you are interested in. If you are interested in the topmost stack frame of the current goroutine (or thread) use: `EvalScope{ GoroutineID: -1, Frame: 0 }`.

​	所有评估 API 调用（除了 `ListPackageVars`）也接受一个 `EvalScope` 参数，它指定了您感兴趣的堆栈帧。如果您感兴趣的是当前 goroutine（或线程）的最上层堆栈帧，请使用：`EvalScope{ GoroutineID: -1, Frame: 0 }`。

More information on the expression language interpreted by RPCServer.Eval can be found [here](https://github.com/go-delve/delve/tree/master/Documentation/cli/expr.md).

​	有关 `RPCServer.Eval` 解释的表达式语言的更多信息，可以在 [此处](https://github.com/go-delve/delve/tree/master/Documentation/cli/expr.md) 查找。

### 变量遮蔽 Variable shadowing



Let's assume you are debugging a piece of code that looks like this:

​	假设您正在调试如下代码：

```
	for i := 0; i < N; i++ {
		for i := 0; i < M; i++ {
			f(i) // <-- debugger is stopped here
		}
	}
```



The response to a ListLocalVars request will list two variables named `i`, because at that point in the code two variables named `i` exist and are in scope. Only one (the innermost one), however, is visible to the user. The other one is *shadowed*.

​	对 `ListLocalVars` 请求的响应将列出两个名为 `i` 的变量，因为在代码的那个点，两个名为 `i` 的变量存在并且在作用域内。然而，只有一个（最内层的）对用户可见。另一个是 *遮蔽的*。

Delve will tell you which variable is shadowed through the `Flags` field of the `Variable` object. If `Flags` has the `VariableShadowed` bit set then the variable in question is shadowed.

​	Delve 会通过 `Variable` 对象的 `Flags` 字段告诉您哪个变量被遮蔽。如果 `Flags` 设置了 `VariableShadowed` 位，则该变量是被遮蔽的。

Users of your client should be able to distinguish between shadowed and non-shadowed variables.

​	您的客户端用户应该能够区分遮蔽和非遮蔽的变量。

## 优雅地结束调试会话 Gracefully ending the debug session



To ensure that Delve cleans up after itself by deleting the `debug` or `debug.test` binary it creates and killing any processes spawned by the program being debugged, the `Detach` command needs to be called. In case you are disconnecting a running program, ensure to halt the program before trying to detach.

​	为了确保 Delve 能够清理自己创建的 `debug` 或 `debug.test` 可执行文件，并杀死程序被调试时产生的任何进程，必须调用 `Detach` 命令。如果您要断开连接的程序正在运行，请确保在尝试分离之前停止该程序。

## 测试客户端 Testing the Client



A set of [example programs is available](https://github.com/aarzilli/delve_client_testing) to test corner cases in handling breakpoints and displaying data structures. Follow the instructions in the README.txt file.

​	一组 [示例程序](https://github.com/aarzilli/delve_client_testing) 可用来测试处理断点和显示数据结构时的边缘情况。请按照 README.txt 文件中的说明进行操作。
