+++
title = "dlv_trace"
date = 2024-12-09T08:06:45+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_trace.md](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv_trace.md)
>
> 收录该文档时间： `2024-12-09T08:06:45+08:00`

## dlv trace



Compile and begin tracing program.

​	编译并开始跟踪程序。

### Synopsis



Trace program execution.

​	跟踪程序的执行。

The trace sub command will set a tracepoint on every function matching the provided regular expression and output information when tracepoint is hit. This is useful if you do not want to begin an entire debug session, but merely want to know what functions your process is executing.

​	`trace` 子命令将在每个匹配给定正则表达式的函数上设置跟踪点，并在跟踪点被触发时输出信息。如果你不想开始一个完整的调试会话，而只是想了解程序执行了哪些函数，这非常有用。

The output of the trace sub command is printed to stderr, so if you would like to only see the output of the trace operations you can redirect stdout.

​	`trace` 子命令的输出将打印到 `stderr`，因此，如果你只想查看跟踪操作的输出，可以将 `stdout` 重定向到其他地方。

```
dlv trace [package] regexp [flags]
```



### Options



```
      --ebpf               Trace using eBPF (experimental).
      						使用 eBPF 进行跟踪（实验性功能）。
  -e, --exec string        Binary file to exec and trace.
  								要执行并跟踪的二进制文件。
      --follow-calls int   Trace all children of the function to the required depth
      						跟踪函数的所有子函数，直到所需的深度。
      						
  -h, --help               help for trace
  								显示 trace 命令的帮助信息。
  								
      --output string      Output path for the binary.
      						指定二进制文件的输出路径。
      						
  -p, --pid int            Pid to attach to.
  								要附加的进程 ID。
  								
  -s, --stack int          Show stack trace with given depth. (Ignored with --ebpf)
  								显示给定深度的堆栈跟踪。（使用 `--ebpf` 时忽略此选项）
  								
  -t, --test               Trace a test binary.
  								跟踪测试二进制文件。
  								
      --timestamp          Show timestamp in the output
      						在输出中显示时间戳。
```



### 从父命令继承的选项 Options inherited from parent commands



```
      --backend string         Backend selection (see 'dlv help backend'). (default "default")
      							后端选择（请参阅 'dlv help backend'）。默认值为 "default"
      							
      --build-flags string     Build flags, to be passed to the compiler. For example: --build-flags="-tags=integration -mod=vendor -cover -v"
      							要传递给编译器的构建标志。例如：--build-flags="-tags=integration -mod=vendor -cover -v"
      							
      --check-go-version       Exits if the version of Go in use is not compatible (too old or too new) with the version of Delve. (default true)
      							如果正在使用的 Go 版本与 Delve 的版本不兼容（过旧或过新），则退出。（默认值为 true）
      							
      --disable-aslr           Disables address space randomization
      							禁用地址空间随机化
      							
      --log                    Enable debugging server logging.
      							启用调试服务器日志。
      							
      --log-dest string        Writes logs to the specified file or file descriptor (see 'dlv help log').
      							将日志写入指定的文件或文件描述符（参见 'dlv help log'）。
      							
      --log-output string      Comma separated list of components that should produce debug output (see 'dlv help log')
      							逗号分隔的组件列表，这些组件应生成调试输出（参见 'dlv help log'）。
      							
  -r, --redirect stringArray   Specifies redirect rules for target process (see 'dlv help redirect')
  									指定目标进程的重定向规则（参见 'dlv help redirect'）。
  									
      --wd string              Working directory for running the program.
      							运行程序的工作目录。
```



### SEE ALSO



- [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve is a debugger for the Go programming language.
  - [dlv](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md) - Delve 是 Go 编程语言的调试器。
