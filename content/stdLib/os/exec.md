+++
title = "exec"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/os/exec@go1.23.0](https://pkg.go.dev/os/exec@go1.23.0)

Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.

​	`exec`包可以运行外部命令。它包装了 os.StartProcess，以便更轻松地重新映射 stdin 和 stdout，连接管道的 I/O，以及做其他调整。

Unlike the "system" library call from C and other languages, the os/exec package intentionally does not invoke the system shell and does not expand any glob patterns or handle other expansions, pipelines, or redirections typically done by shells. The package behaves more like C's "exec" family of functions. To expand glob patterns, either call the shell directly, taking care to escape any dangerous input, or use the path/filepath package's Glob function. To expand environment variables, use package os's ExpandEnv.

​	与 C 语言和其他语言中的 "system" 库调用不同，os/exec 包故意不调用系统 shell，不扩展任何 glob 模式，也不处理 shell 通常执行的其他扩展、管道或重定向。该包的行为更像 C 的 "exec" 函数系列。为了扩展 glob 模式，请直接调用 shell，注意转义任何危险输入，或使用 path/filepath 包的 Glob 函数。为了扩展环境变量，请使用 os 包的 ExpandEnv。

Note that the examples in this package assume a Unix system. They may not run on Windows, and they do not run in the Go Playground used by golang.org and godoc.org.

​	请注意，此包中的示例假定 Unix 系统。它们可能无法在 Windows 上运行，并且它们不在 golang.org 和 godoc.org 使用的 Go Playground 中运行。

## 当前目录中的可执行文件 Executables in the current directory

The functions Command and LookPath look for a program in the directories listed in the current path, following the conventions of the host operating system. Operating systems have for decades included the current directory in this search, sometimes implicitly and sometimes configured explicitly that way by default. Modern practice is that including the current directory is usually unexpected and often leads to security problems.

​	函数 Command 和 LookPath 在列出的路径目录中查找程序，遵循主机操作系统的惯例。操作系统几十年来一直将当前目录包括在此搜索中，有时是隐式地，有时是默认配置为这样。现代实践是包括当前目录通常是意外的，而且经常导致安全问题。

To avoid those security problems, as of Go 1.19, this package will not resolve a program using an implicit or explicit path entry relative to the current directory. That is, if you run exec.LookPath("go"), it will not successfully return ./go on Unix nor .\go.exe on Windows, no matter how the path is configured. Instead, if the usual path algorithms would result in that answer, these functions return an error err satisfying errors.Is(err, ErrDot).

​	为了避免这些安全问题，从 Go 1.19 开始，此包将不会使用相对于当前目录的隐式或显式路径条目解析程序。也就是说，如果您运行 exec.LookPath("go")，它将不会在 Unix 上成功返回 ./go，也不会在 Windows 上成功返回 .\go.exe，无论路径如何配置。相反，如果通常的路径算法会导致这样的答案，这些函数将返回一个满足 errors.Is(err, ErrDot) 的错误 err。

For example, consider these two program snippets:

​	例如，考虑以下两个程序片段：

```
path, err := exec.LookPath("prog")
if err != nil {
	log.Fatal(err)
}
use(path)
```

and

和

```
cmd := exec.Command("prog")
if err := cmd.Run(); err != nil {
	log.Fatal(err)
}
```

These will not find and run ./prog or .\prog.exe, no matter how the current path is configured.

​	无论当前路径如何配置，它们都无法找到并运行 `./prog` 或 `.\prog.exe`。

Code that always wants to run a program from the current directory can be rewritten to say "./prog" instead of "prog".

​	总是想从当前目录运行程序的代码可以重写为 "`./prog`"，而不是 "`prog`"。

Code that insists on including results from relative path entries can instead override the error using an errors.Is check:

​	坚持包括来自相对路径条目的结果的代码可以使用 errors.Is 检查覆盖错误：

```
path, err := exec.LookPath("prog")
if errors.Is(err, exec.ErrDot) {
	err = nil
}
if err != nil {
	log.Fatal(err)
}
use(path)
```

and

和

```
cmd := exec.Command("prog")
if errors.Is(cmd.Err, exec.ErrDot) {
	cmd.Err = nil
}
if err := cmd.Run(); err != nil {
	log.Fatal(err)
}
```

Setting the environment variable GODEBUG=execerrdot=0 disables generation of ErrDot entirely, temporarily restoring the pre-Go 1.19 behavior for programs that are unable to apply more targeted fixes. A future version of Go may remove support for this variable.

​	设置环境变量GODEBUG=execerrdot=0将禁用生成ErrDot，为无法应用更有针对性的修复的程序暂时恢复Go 1.19之前的行为。将来的Go版本可能会删除对该变量的支持。

Before adding such overrides, make sure you understand the security implications of doing so. See https://go.dev/blog/path-security for more information.

​	在添加这样的覆盖之前，请确保您了解这样做的安全性影响。有关更多信息，请参见https://go.dev/blog/path-security。


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/exec/exec.go;l=1303)

``` go 
var ErrDot = errors.New("cannot run executable found relative to current directory")
```

ErrDot indicates that a path lookup resolved to an executable in the current directory due to ‘.’ being in the path, either implicitly or explicitly. See the package documentation for details.

​	ErrDot 表示由于路径中存在 '.'(点号)，无论是隐式还是显式地，导致路径查找解析到了当前目录中的可执行文件。有关详细信息，请参阅包文档。

Note that functions in this package do not return ErrDot directly. Code should use errors.Is(err, ErrDot), not err == ErrDot, to test whether a returned error err is due to this condition.

​	请注意，此包中的函数不会直接返回 ErrDot。代码应使用 errors.Is(err, ErrDot) 而不是 err == ErrDot 来测试返回的错误 err 是否是由于此条件而引起的。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/exec/lp_unix.go;l=20)

``` go 
var ErrNotFound = errors.New("executable file not found in $PATH")
```

ErrNotFound is the error resulting if a path search failed to find an executable file.

​	ErrNotFound 是搜索路径失败未找到可执行文件时返回的错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/exec/exec.go;l=127)

``` go 
var ErrWaitDelay = errors.New("exec: WaitDelay expired before I/O complete")
```

ErrWaitDelay is returned by (*Cmd).Wait if the process exits with a successful status code but its output pipes are not closed before the command's WaitDelay expires.

​	如果进程以成功的状态码退出，但其输出管道在命令的 WaitDelay 过期之前未关闭，则 `(*Cmd).Wait` 返回 ErrWaitDelay。

## 函数

### func LookPath 

``` go 
func LookPath(file string) (string, error)
```

LookPath searches for an executable named file in the directories named by the PATH environment variable. If file contains a slash, it is tried directly and the PATH is not consulted. Otherwise, on success, the result is an absolute path.

​	LookPath函数在 PATH 环境变量指定的目录中查找名为 file 的可执行文件。如果 file 包含斜杠，则尝试直接使用它，不会查询 PATH。否则，如果成功，结果是绝对路径。

In older versions of Go, LookPath could return a path relative to the current directory. As of Go 1.19, LookPath will instead return that path along with an error satisfying errors.Is(err, ErrDot). See the package documentation for more details.

​	在较旧的 Go 版本中，LookPath 可能返回相对于当前目录的路径。从 Go 1.19 开始，如果常规的路径算法会导致该答案，则 LookPath 将返回该路径以及满足 errors.Is(err, ErrDot) 的错误。有关更多详细信息，请参阅包文档。

#### LookPath Example

``` go 
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("fortune")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", path)
}

```



## 类型

### type Cmd 

``` go 
type Cmd struct {
    // Path is the path of the command to run.
	// Path 是要运行的命令的路径。
	//
    // This is the only field that must be set to a non-zero
	// value. If Path is relative, it is evaluated relative
	// to Dir.
	// 这是唯一必须设置为非零值的字段。
	// 如果 Path 是相对路径，则相对于 Dir 进行求值。
	Path string

    // Args holds command line arguments, including the command as Args[0].
	// If the Args field is empty or nil, Run uses {Path}.
	// Args 包含命令行参数，包括命令作为 Args[0]。
	// 如果 Args 字段为空或 nil，则 Run 使用 {Path}。
	//
    // In typical use, both Path and Args are set by calling Command.
	// 通常情况下，Path 和 Args 由调用 Command 来设置。
	Args []string

    // Env specifies the environment of the process.
	// Each entry is of the form "key=value".
	// If Env is nil, the new process uses the current process's
	// environment.
	// If Env contains duplicate environment keys, only the last
	// value in the slice for each duplicate key is used.
	// As a special case on Windows, SYSTEMROOT is always added if
	// missing and not explicitly set to the empty string.
	// Env 指定进程的环境。
	// 每个条目的格式为 "key=value"。
	// 如果 Env 为 nil，则新进程使用当前进程的环境。
	// 如果 Env 包含重复的环境键，
    // 则只使用每个重复键的 slice 中的最后一个值。
	// 在 Windows 上，
    // 特殊情况下如果缺少 SYSTEMROOT 并且未显式设置为空字符串，
    // 则始终会添加它。
	Env []string

    // Dir specifies the working directory of the command.
	// If Dir is the empty string, Run runs the command in the
	// calling process's current directory.
	// Dir 指定命令的工作目录。
	// 如果 Dir 是空字符串，则 Run 在调用进程的当前目录中运行命令。
	Dir string

    // Stdin specifies the process's standard input.
	// Stdin 指定进程的标准输入。
	//
    // If Stdin is nil, the process reads from the null device (os.DevNull).
	// 如果 Stdin 是 nil，则进程从空设备(os.DevNull)读取。
	//
    // If Stdin is an *os.File, the process's standard input is connected
	// directly to that file.
	// 如果 Stdin 是 *os.File，则进程的标准输入直接连接到该文件。
	//
    // Otherwise, during the execution of the command a separate
	// goroutine reads from Stdin and delivers that data to the command
	// over a pipe. In this case, Wait does not complete until the goroutine
	// stops copying, either because it has reached the end of Stdin
	// (EOF or a read error), or because writing to the pipe returned an error,
	// or because a nonzero WaitDelay was set and expired.
	// 否则，在命令执行期间，
    // 单独的 goroutine 从 Stdin 读取数据并通过管道将其传递给命令。
	// 在这种情况下，Wait 不会完成，直到 goroutine 停止复制，
    // 要么是因为已经到达了 Stdin 的末尾(EOF 或读取错误)，
	// 要么是因为写入管道返回了错误，
    // 要么是因为设置了非零的 WaitDelay 并已过期。
	Stdin io.Reader

    // Stdout and Stderr specify the process's standard output and error.
	// Stdout和Stderr指定进程的标准输出和错误。
	//
    // If either is nil, Run connects the corresponding file descriptor
	// to the null device (os.DevNull).
	// 如果它们中的任何一个为nil，
    // 则Run会将相应的文件描述符连接到null设备(os.DevNull)。
	//
    // If either is an *os.File, the corresponding output from the process
	// is connected directly to that file.
	// 如果它们中的任何一个是*os.File，
    // 则进程的相应输出将直接连接到该文件。
	//
    // Otherwise, during the execution of the command a separate goroutine
	// reads from the process over a pipe and delivers that data to the
	// corresponding Writer. In this case, Wait does not complete until the
	// goroutine reaches EOF or encounters an error or a nonzero WaitDelay
	// expires.
	// 否则，在执行命令期间，
    // 一个单独的goroutine会通过管道从进程读取数据，
    // 并将其传递到相应的Writer。
	// 在这种情况下，Wait不会完成，
    // 直到goroutine到达EOF或遇到错误或非零的WaitDelay到期。
	//
    // If Stdout and Stderr are the same writer, and have a type that can
	// be compared with ==, at most one goroutine at a time will call Write.
	// 如果Stdout和Stderr是同一个writer，
    // 并且具有可以与==比较的类型，
    // 则最多只有一个goroutine每次调用Write。
	Stdout io.Writer
	Stderr io.Writer

    // ExtraFiles specifies additional open files to be inherited by the
	// new process. It does not include standard input, standard output, or
	// standard error. If non-nil, entry i becomes file descriptor 3+i.
	// ExtraFiles指定要由新进程继承的其他打开文件。
    // 它不包括标准输入、标准输出或标准错误。
	// 如果非nil，则条目i变成文件描述符3+i。
	//
    // ExtraFiles is not supported on Windows.
	// ExtraFiles在Windows上不受支持。
	ExtraFiles []*os.File

    // SysProcAttr holds optional, operating system-specific attributes.
	// Run passes it to os.StartProcess as the os.ProcAttr's Sys field.
	// SysProcAttr保存可选的特定于操作系统的属性。
    // Run将其作为os.ProcAttr的Sys字段传递给os.StartProcess。
	SysProcAttr *syscall.SysProcAttr

    // Process is the underlying process, once started.
	// Process是一旦启动就是底层进程。
	// 如果成功启动进程，Wait或Run将在命令完成时填充其ProcessState。
	Process *os.Process

    // ProcessState contains information about an exited process.
	// If the process was started successfully, Wait or Run will
	// populate its ProcessState when the command completes.
	// ProcessState包含有关退出进程的信息。
	// 如果成功启动进程，Wait或Run将在命令完成时填充其ProcessState。
	ProcessState *os.ProcessState

	Err error // LookPath error, if any.// LookPath错误，如果有的话。

    // If Cancel is non-nil, the command must have been created with
	// CommandContext and Cancel will be called when the command's
	// Context is done. By default, CommandContext sets Cancel to
	// call the Kill method on the command's Process.
	// 如果Cancel非nil，则必须使用CommandContext创建命令，
    // 当命令的Context完成时，将调用Cancel。
	// 默认情况下，
    // CommandContext将Cancel设置为在命令的Process上调用Kill方法。
	//
    // Typically a custom Cancel will send a signal to the command's
	// Process, but it may instead take other actions to initiate cancellation,
	// such as closing a stdin or stdout pipe or sending a shutdown request on a
	// network socket.
	// 通常，自定义Cancel将向命令的Process发送一个信号，
    // 但它可能会采取其他措施来启动取消操作，
	// 例如关闭stdin或stdout管道或在网络套接字上发送关闭请求。
	//
    // If the command exits with a success status after Cancel is
	// called, and Cancel does not return an error equivalent to
	// os.ErrProcessDone, then Wait and similar methods will return a non-nil
	// error: either an error wrapping the one returned by Cancel,
	// or the error from the Context.
	// (If the command exits with a non-success status, or Cancel
	// returns an error that wraps os.ErrProcessDone, Wait and similar methods
	// continue to return the command's usual exit status.)
	// 如果Cancel在调用后以成功状态退出，
    // 并且Cancel未返回与os.ErrProcessDone等效的错误，
	// 那么Wait和类似方法将返回非nil错误：
    // 包装由Cancel返回的错误或上下文的错误。
	// (如果命令以非成功状态退出，
    // 或者Cancel返回一个包装os.ErrProcessDone的错误，
	// Wait和类似方法将继续返回命令的常规退出状态。)
	//
    // If Cancel is set to nil, nothing will happen immediately when the command's
	// Context is done, but a nonzero WaitDelay will still take effect. That may
	// be useful, for example, to work around deadlocks in commands that do not
	// support shutdown signals but are expected to always finish quickly.
	// 如果将Cancel设置为nil，
    // 则命令的Context完成时不会立即发生任何事情，
	// 但非零的WaitDelay仍然会生效。
    // 这可能是有用的，
    // 例如，为了解决不支持关闭信号但预计始终快速完成的命令中的死锁。
	//
    // Cancel will not be called if Start returns a non-nil error.
	// 如果Start返回非nil错误，则不会调用Cancel。
	Cancel func() error

    // If WaitDelay is non-zero, it bounds the time spent waiting on two sources
	// of unexpected delay in Wait: a child process that fails to exit after the
	// associated Context is canceled, and a child process that exits but leaves
	// its I/O pipes unclosed.
	// 如果WaitDelay不为零，
    // 则它限制了在Wait时等待两个意外延迟的来源所花费的时间：
	// 子进程在相关的上下文被取消后仍然无法退出，
    // 以及子进程退出但其I/O管道未关闭。
	//
    // The WaitDelay timer starts when either the associated Context is done or a
	// call to Wait observes that the child process has exited, whichever occurs
	// first. When the delay has elapsed, the command shuts down the child process
	// and/or its I/O pipes.
	// 当关联的上下文完成或Wait的调用观察到子进程已退出时，
    // WaitDelay计时器启动，
	// 以确保在延迟时间后，命令将关闭子进程及/或其I/O管道。
	//
    // If the child process has failed to exit — perhaps because it ignored or
	// failed to receive a shutdown signal from a Cancel function, or because no
	// Cancel function was set — then it will be terminated using os.Process.Kill.
	// 如果子进程无法退出，
    // 可能是因为它忽略或未能接收来自取消函数的关闭信号，
	// 或因为未设置取消函数，则它将使用os.Process.Kill终止。
	//
    // Then, if the I/O pipes communicating with the child process are still open,
	// those pipes are closed in order to unblock any goroutines currently blocked
	// on Read or Write calls.
	// 然后，如果与子进程通信的I/O管道仍然打开，
    // 则将关闭这些管道，
    // 以解除当前阻塞在Read或Write调用上的所有goroutine。
	//
    // If pipes are closed due to WaitDelay, no Cancel call has occurred,
	// and the command has otherwise exited with a successful status, Wait and
	// similar methods will return ErrWaitDelay instead of nil.
	// 如果由于WaitDelay而关闭管道，
    // 没有发生取消调用，并且命令已以成功状态退出，
	// 则Wait等方法将返回ErrWaitDelay而不是nil。
	//
    // If WaitDelay is zero (the default), I/O pipes will be read until EOF,
	// which might not occur until orphaned subprocesses of the command have
	// also closed their descriptors for the pipes.
	// 如果WaitDelay为零(默认值)，
    // 则将读取I/O管道，直到EOF为止，
    // 这可能不会发生，直到命令的孤立子进程也关闭了它们的管道描述符。
	WaitDelay time.Duration
	// 包含已过滤或未导出的字段
}
```

Cmd represents an external command being prepared or run.

​	Cmd 表示正在准备或运行的外部命令。

A Cmd cannot be reused after calling its Run, Output or CombinedOutput methods.

​	在调用其 Run、Output 或 CombinedOutput 方法之后，Cmd 不能被重用。

#### func Command 

``` go 
func Command(name string, arg ...string) *Cmd
```

Command returns the Cmd struct to execute the named program with the given arguments.

​	Command函数返回一个 Cmd 结构，以给定参数执行指定的命令。

It sets only the Path and Args in the returned structure.

​	它只在返回的结构中设置 Path 和 Args 字段。

If name contains no path separators, Command uses LookPath to resolve name to a complete path if possible. Otherwise it uses name directly as Path.

​	如果 name 不包含路径分隔符，则 Command 会尝试使用 LookPath 将 name 解析为完整路径。否则，它将直接使用 name 作为 Path。

The returned Cmd's Args field is constructed from the command name followed by the elements of arg, so arg should not include the command name itself. For example, Command("echo", "hello"). Args[0] is always name, not the possibly resolved Path.

​	返回的 Cmd 的 Args 字段由命令名称后跟 arg 的元素构成，因此 arg 不应包括命令本身的名称。例如，Command("echo", "hello")。Args[0] 总是 name，而不是可能解析后的 Path。

On Windows, processes receive the whole command line as a single string and do their own parsing. Command combines and quotes Args into a command line string with an algorithm compatible with applications using CommandLineToArgvW (which is the most common way). Notable exceptions are msiexec.exe and cmd.exe (and thus, all batch files), which have a different unquoting algorithm. In these or other similar cases, you can do the quoting yourself and provide the full command line in SysProcAttr.CmdLine, leaving Args empty.

​	在 Windows 上，进程接收整个命令行作为单个字符串，并进行自己的解析。Command 将 Args 组合并用算法兼容于使用 CommandLineToArgvW 的应用程序的引号为命令行字符串。值得注意的例外是 msiexec.exe 和 cmd.exe(因此所有的批处理文件)，它们有不同的去引号算法。在这些或其他类似的情况下，您可以自己引用并在 SysProcAttr.CmdLine 中提供完整的命令行，将 Args 留空。

##### Command Example

``` go 
package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}

```



##### Command Example(Environment) 

``` go 
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("prog")
	cmd.Env = append(os.Environ(),
		"FOO=duplicate_value", // ignored
		"FOO=actual_value",    // this value is used
	)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

```



#### func CommandContext  <- go1.7

``` go 
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd
```

CommandContext is like Command but includes a context.

​	CommandContext函数类似于 Command，但包括一个 context。

The provided context is used to interrupt the process (by calling cmd.Cancel or os.Process.Kill) if the context becomes done before the command completes on its own.

​	提供的 context 在命令完成之前中止进程(通过调用 cmd.Cancel 或 os.Process.Kill)。

CommandContext sets the command's Cancel function to invoke the Kill method on its Process, and leaves its WaitDelay unset. The caller may change the cancellation behavior by modifying those fields before starting the command.

​	CommandContext函数将命令的 Cancel 函数设置为在其 Process 上调用 Kill 方法，并将其 WaitDelay 留空。调用者可以在启动命令之前修改这些字段以改变取消行为。

##### CommandContext Example

``` go 
package main

import (
	"context"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
	}
}

```



#### (*Cmd) CombinedOutput 

``` go 
func (c *Cmd) CombinedOutput() ([]byte, error)
```

CombinedOutput runs the command and returns its combined standard output and standard error.

​	CombinedOutput方法运行命令并返回其合并的标准输出和标准错误。

##### CombinedOutput Example

``` go 
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

```



#### (*Cmd) Environ  <- go1.19

``` go 
func (c *Cmd) Environ() []string
```

Environ returns a copy of the environment in which the command would be run as it is currently configured.

​	Environ方法返回命令将按当前配置运行的环境的副本。

##### Environ Example

``` go 
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("pwd")

	// Set Dir before calling cmd.Environ so that it will include an
	// updated PWD variable (on platforms where that is used).
	cmd.Dir = ".."
	cmd.Env = append(cmd.Environ(), "POSIXLY_CORRECT=1")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}

```



#### (*Cmd) Output 

``` go 
func (c *Cmd) Output() ([]byte, error)
```

Output runs the command and returns its standard output. Any returned error will usually be of type *ExitError. If c.Stderr was nil, Output populates ExitError.Stderr.

​	Output方法运行命令并返回其标准输出。任何返回的错误通常都是`*ExitError`类型。如果c.Stderr为nil，则Output将填充ExitError.Stderr。

``` go 
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

```



#### (*Cmd) Run 

``` go 
func (c *Cmd) Run() error
```

Run starts the specified command and waits for it to complete.

​	Run方法启动指定的命令并等待其完成。

The returned error is nil if the command runs, has no problems copying stdin, stdout, and stderr, and exits with a zero exit status.

​	如果命令运行良好，无法复制标准输入、标准输出和标准错误，并带有零退出状态，则返回的错误为nil。

If the command starts but does not complete successfully, the error is of type *ExitError. Other error types may be returned for other situations.

​	如果命令启动但未成功完成，则返回的错误类型为*ExitError。其他错误类型可能会针对其他情况返回。

If the calling goroutine has locked the operating system thread with runtime.LockOSThread and modified any inheritable OS-level thread state (for example, Linux or Plan 9 name spaces), the new process will inherit the caller's thread state.

​	如果调用的goroutine使用runtime.LockOSThread锁定操作系统线程并修改任何可继承的操作系统级线程状态(例如，Linux或Plan 9名称空间)，则新进程将继承调用者的线程状态。

##### Run Example

``` go 
package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

```





#### (*Cmd) Start 

``` go 
func (c *Cmd) Start() error
```

Start starts the specified command but does not wait for it to complete.

​	Start方法启动指定的命令，但不等待其完成。

If Start returns successfully, the c.Process field will be set.

​	如果Start成功返回，则c.Process字段将被设置。

After a successful call to Start the Wait method must be called in order to release associated system resources.

​	在成功调用Start方法之后，必须调用Wait方法以释放相关的系统资源。

##### Start Example

``` go 
package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

```



#### (*Cmd) StderrPipe 

``` go 
func (c *Cmd) StderrPipe() (io.ReadCloser, error)
```

StderrPipe returns a pipe that will be connected to the command's standard error when the command starts.

​	StderrPipe方法返回一个管道，该管道将在命令启动时连接到命令的标准错误。

Wait will close the pipe after seeing the command exit, so most callers need not close the pipe themselves. It is thus incorrect to call Wait before all reads from the pipe have completed. For the same reason, it is incorrect to use Run when using StderrPipe. See the StdoutPipe example for idiomatic usage.

​	Wait将在看到命令退出后关闭该管道，因此大多数调用方不需要自己关闭该管道。因此，在所有从管道中读取完成之前调用Wait是不正确的。出于同样的原因，当使用StderrPipe时使用Run是不正确的。有关惯用用法，请参阅StdoutPipe示例。

##### StderrPipe Example

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

```



#### (*Cmd) StdinPipe 

``` go 
func (c *Cmd) StdinPipe() (io.WriteCloser, error)
```

StdinPipe returns a pipe that will be connected to the command's standard input when the command starts. The pipe will be closed automatically after Wait sees the command exit. A caller need only call Close to force the pipe to close sooner. For example, if the command being run will not exit until standard input is closed, the caller must close the pipe.

​	StdinPipe方法返回一个管道，该管道将在命令启动时连接到命令的标准输入。当Wait看到命令退出后，管道将自动关闭。调用方只需调用Close来更快地关闭管道。例如，如果运行的命令不会退出，直到标准输入被关闭，那么调用方必须关闭管道。

##### StdinPipe Example

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

```



#### (*Cmd) StdoutPipe 

``` go 
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
```

StdoutPipe returns a pipe that will be connected to the command's standard output when the command starts.

​	StdoutPipe方法返回一个管道，该管道将在命令启动时连接到命令的标准输出。

Wait will close the pipe after seeing the command exit, so most callers need not close the pipe themselves. It is thus incorrect to call Wait before all reads from the pipe have completed. For the same reason, it is incorrect to call Run when using StdoutPipe. See the example for idiomatic usage.

​	Wait将在看到命令退出后关闭该管道，因此大多数调用方不需要自己关闭该管道。因此，在所有从管道中读取完成之前调用Wait是不正确的。出于同样的原因，在使用StdoutPipe时调用Run也是不正确的。有关惯用用法，请参阅示例。

##### StdoutPipe Example

``` go 
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

```



#### (*Cmd) String  <- go1.13

``` go 
func (c *Cmd) String() string
```

String returns a human-readable description of c. It is intended only for debugging. In particular, it is not suitable for use as input to a shell. The output of String may vary across Go releases.

​	String方法返回c的人类可读描述。它仅用于调试。特别是，它不适合用作shell的输入。String的输出可能会因Go版本而异。

#### (*Cmd) Wait 

``` go 
func (c *Cmd) Wait() error
```

Wait waits for the command to exit and waits for any copying to stdin or copying from stdout or stderr to complete.

​	Wait方法等待命令退出，并等待从标准输入复制或从标准输出或标准错误复制完成。

The command must have been started by Start.

​	该命令必须已由Start启动。

The returned error is nil if the command runs, has no problems copying stdin, stdout, and stderr, and exits with a zero exit status.

​	如果命令运行良好，无法复制标准输入、标准输出和标准错误，并带有零退出状态，则返回的错误为nil。

If the command fails to run or doesn't complete successfully, the error is of type *ExitError. Other error types may be returned for I/O problems.

​	如果命令无法运行或未成功完成，则返回的错误类型为*ExitError。对于I/O问题，可能会返回其他错误类型。

If any of c.Stdin, c.Stdout or c.Stderr are not an *os.File, Wait also waits for the respective I/O loop copying to or from the process to complete.

​	如果c.Stdin、c.Stdout或c.Stderr中的任何一个不是*os.File，则Wait还将等待复制到或从进程中的相应I/O循环完成。

Wait releases any resources associated with the Cmd.

​	Wait释放与Cmd关联的任何资源。

### type Error 

``` go 
type Error struct {
    // Name is the file name for which the error occurred.
	// Name 是出现错误的文件名。
	Name string
    // Err is the underlying error.
	// Err 是底层错误。
	Err error
}
```

Error is returned by LookPath when it fails to classify a file as an executable.

​	当 LookPath函数无法将文件分类为可执行文件时，将返回 Error。

#### (*Error) Error 

``` go 
func (e *Error) Error() string
```

#### (*Error) Unwrap  <- go1.13

``` go 
func (e *Error) Unwrap() error
```

### type ExitError 

``` go 
type ExitError struct {
	*os.ProcessState

    // Stderr holds a subset of the standard error output from the
	// Cmd.Output method if standard error was not otherwise being
	// collected.
	// Stderr 保存 Cmd.Output 方法的一部分标准错误输出，
    // 如果未收集标准错误输出，则会保存该部分。
	//
    // If the error output is long, Stderr may contain only a prefix
	// and suffix of the output, with the middle replaced with
	// text about the number of omitted bytes.
	// 如果错误输出很长，则 Stderr 可能只包含输出的前缀和后缀，
    // 其中间部分被替换为有关省略字节数的文本。
	//
    // Stderr is provided for debugging, for inclusion in error messages.
	// Users with other needs should redirect Cmd.Stderr as needed.
	// Stderr 用于调试和包含在错误消息中。
    // 有其他需求的用户应根据需要重定向 Cmd.Stderr。
	Stderr []byte
}
```

An ExitError reports an unsuccessful exit by a command.

​	ExitError 报告命令未能成功退出。

#### (*ExitError) Error 

``` go 
func (e *ExitError) Error() string
```