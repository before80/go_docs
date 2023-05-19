+++
title = "signal"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# signal

https://pkg.go.dev/os/signal@go1.20.1

​	signal包实现对进入的信号的访问。

​	信号主要在类 Unix 的系统上使用。有关在 Windows 和 Plan 9 上使用此包的信息，请参见下文。

### 信号的类型

​	SIGKILL 和 SIGSTOP 信号可能不会被程序捕获，因此无法受到此包的影响。

​	同步信号是由程序执行中的错误触发的信号：SIGBUS、SIGFPE 和 SIGSEGV。只有在程序执行时引起的情况下，才被认为是同步的，而不是通过 os.Process.Kill、kill 程序或某些类似机制发送的信号。通常情况下，除非在下文中讨论的情况下，Go 程序将把同步信号转换为运行时 panic。

​	剩余的信号是异步信号。它们不是由程序错误触发的，而是从内核或其他程序发送的。

​	在异步信号中，当程序失去控制终端时，会发送 SIGHUP 信号。当控制终端的用户按下中断字符时，会发送 SIGINT 信号，默认情况下是 ^C(Control-C)。当控制终端的用户按下退出字符时，会发送 SIGQUIT 信号，默认情况下是 ^\(Control-Backslash)。通常情况下，你可以通过按下 ^C 使程序简单退出，并通过按下 ^\ 使程序退出并出现堆栈转储。

#### Go 程序中信号的默认行为

​	默认情况下，同步信号会转换为运行时 panic。SIGHUP、SIGINT 或 SIGTERM 信号会导致程序退出。SIGQUIT、SIGILL、SIGTRAP、SIGABRT、SIGSTKFLT、SIGEMT 或 SIGSYS 信号会导致程序退出并出现堆栈转储。SIGTSTP、SIGTTIN 或 SIGTTOU 信号获取系统默认行为(这些信号由 shell 用于作业控制)。SIGPROF 信号由 Go 运行时直接处理，以实现 runtime.CPUProfile。其他信号将被捕获，但不会采取任何措施。

​	如果使用忽略信号处理程序(signal handler)SIG_IGN 启动 Go 程序，SIGHUP 或 SIGINT 信号将保持被忽略状态。

​	如果使用非空信号掩码启动 Go 程序，通常将受到控制。然而，某些信号是明确解除阻止的：同步信号、SIGILL、SIGTRAP、SIGSTKFLT、SIGCHLD、SIGPROF，以及在 Linux 上，信号 32(SIGCANCEL)和 33(SIGSETXID)(SIGCANCEL 和 SIGSETXID 在 glibc 内部使用)。由 os.Exec 或 os/exec 包启动的子进程将继承修改后的信号掩码。

#### 更改Go程序中信号的行为

​	这个包中的函数允许程序更改Go程序处理信号的方式。

​	Notify禁用给定一组异步信号的默认行为，并通过一个或多个注册的通道将其传递。具体来说，它适用于信号SIGHUP、SIGINT、SIGQUIT、SIGABRT和SIGTERM。它还适用于作业控制信号SIGTSTP、SIGTTIN和SIGTTOU，在这种情况下，系统默认行为不会发生。它还适用于某些否则不会引起任何操作的信号：SIGUSR1、SIGUSR2、SIGPIPE、SIGALRM、SIGCHLD、SIGCONT、SIGURG、SIGXCPU、SIGXFSZ、SIGVTALRM、SIGWINCH、SIGIO、SIGPWR、SIGSYS、SIGINFO、SIGTHR、SIGWAITING、SIGLWP、SIGFREEZE、SIGTHAW、SIGLOST、SIGXRES、SIGJVM1、SIGJVM2和系统上使用的任何实时信号。注意，并非所有这些信号都在所有系统上可用。

​	如果程序启动时忽略了SIGHUP或SIGINT，并且为这两个信号调用了Notify，那么一个信号处理程序将被安装到该信号上，它将不再被忽略。如果稍后为该信号调用Reset或Ignore，或为该信号传递到Notify的所有通道调用Stop，则该信号将再次被忽略。Reset将恢复该信号的系统默认行为，而Ignore将导致系统完全忽略该信号。

​	如果程序启动时具有非空信号掩码，则某些信号将被明确解除阻塞，如上所述。如果为阻塞信号调用Notify，它将被解除阻塞。如果稍后为该信号调用Reset，或为该信号传递到Notify的所有通道调用Stop，则该信号将再次被阻塞。

#### SIGPIPE 

​	当 Go 程序写入到一个已断开的管道时，内核会触发一个 SIGPIPE 信号。

​	如果程序没有调用 Notify 接收 SIGPIPE 信号，那么行为取决于文件描述符号。向标准输出或标准错误的已断开管道上写入将导致程序退出并收到 SIGPIPE 信号。在其他文件描述符上向已断开管道写入将不会对 SIGPIPE 信号采取任何行动，并且写入将失败并返回 EPIPE 错误。

​	如果程序已经调用了 Notify 来接收 SIGPIPE 信号，则文件描述符号无关紧要。SIGPIPE 信号将被传递到 Notify 通道，而写入将失败并返回 EPIPE 错误。

​	这意味着，默认情况下，命令行程序的行为将像典型的 Unix 命令行程序一样，而其他程序在向已关闭的网络连接写入时将不会崩溃。

#### 使用 cgo 或 SWIG 的 Go 程序

​	在包含非 Go 代码的 Go 程序中，通常是使用 cgo 或 SWIG 访问 C/C++ 代码，Go 的启动代码通常会先运行。它会在非 Go 启动代码运行之前按照 Go 运行时的期望配置信号处理程序。如果非 Go 启动代码希望安装自己的信号处理程序，则必须采取某些步骤来确保 Go 能正常运行。本节记录了这些步骤以及非 Go 代码对信号处理程序设置所产生的整体影响。在罕见情况下，非 Go 代码可能会在 Go 代码之前运行，在这种情况下下一节也适用。

​	如果 Go 程序调用的非 Go 代码未更改任何信号处理程序或掩码，则其行为与纯 Go 程序相同。

​	如果非 Go 代码安装了任何信号处理程序，则必须在 sigaction 中使用 SA_ONSTACK 标志。如果未这样做，则在接收到信号时程序可能会崩溃。Go 程序通常使用有限堆栈运行，因此会设置备用信号堆栈。

​	如果非 Go 代码为任何同步信号(SIGBUS、SIGFPE、SIGSEGV)安装了信号处理程序，则应记录现有的 Go 信号处理程序。如果这些信号在执行 Go 代码时发生，则应调用 Go 信号处理程序(可以通过查看传递给信号处理程序的 PC 来确定信号是否在执行 Go 代码时发生)。否则，某些 Go 运行时 panic 将不会按预期发生。

​	如果非 Go 代码为任何异步信号安装了信号处理程序，则可以按其选择调用 Go 信号处理程序或不调用。自然，如果它不调用 Go 信号处理程序，则上面描述的 Go 行为将不会发生。这可能是特别关注 SIGPROF 信号的问题。

​	非 Go 代码不应更改由 Go 运行时创建的任何线程的信号掩码。如果非 Go 代码启动自己的新线程，则可以自由地设置信号掩码。

​	如果非 Go 代码启动新线程，更改信号掩码，然后在该线程中调用 Go 函数，则 Go 运行时将自动解除阻止某些信号：同步信号 SIGILL、SIGTRAP、SIGSTKFLT、SIGCHLD、SIGPROF、SIGCANCEL 和 SIGSETXID。当 Go 函数返回时，将恢复非 Go 信号掩码。

​	如果 Go 信号处理程序在非 Go 线程上调用且未运行 Go 代码，则处理程序通常将信号转发给非 Go 代码，如下所示。如果信号是 SIGPROF，则 Go 处理程序什么也不做。否则，Go 处理程序将其自身删除、解除信号阻止并再次引发该信号，以调用任何非 Go 处理程序或默认系统处理程序。如果程序未退出，则 Go 处理程序将重新安装自身并继续执行程序。

​	如果接收到 SIGPIPE 信号，则 Go 程序将在 Go 线程上接收到 SIGPIPE 时调用上述的特殊处理程序。如果在非 Go 线程上接收到 SIGPIPE，则该信号将转发到非 Go 处理程序(如果有的话)，如果没有，则默认的系统处理程序将导致程序终止。

#### 调用 Go 代码的非 Go 程序

​	当使用像-buildmode=c-shared这样的选项构建 Go 代码时，它将作为现有非 Go 程序的一部分运行。在 Go 代码启动时，非 Go 代码可能已经安装了信号处理程序(在使用 cgo 或 SWIG 时也可能会发生这种情况；在这种情况下，此处的讨论也适用)。对于-buildmode=c-archive，Go 运行时将在全局构造函数时间初始化信号。对于-buildmode=c-shared，Go 运行时将在加载共享库时初始化信号。

​	如果 Go 运行时看到 SIGCANCEL 或 SIGSETXID 信号的现有信号处理程序(仅在 Linux 上使用)，它将打开 SA_ONSTACK 标志并保留信号处理程序。

​	对于同步信号和 SIGPIPE，Go 运行时将安装信号处理程序。它将保存任何现有的信号处理程序。如果在执行非 Go 代码时出现同步信号，Go 运行时将调用现有的信号处理程序而不是 Go 信号处理程序。

​	使用-buildmode=c-archive或-buildmode=c-shared构建的 Go 代码默认不会安装任何其他信号处理程序。如果存在现有的信号处理程序，则 Go 运行时将打开 SA_ONSTACK 标志并保留信号处理程序。如果对于异步信号调用了 Notify，将为该信号安装 Go 信号处理程序。如果稍后调用 Reset 以该信号，则将重新安装该信号的原始处理方式，如果有，则恢复非 Go 信号处理程序。

​	未使用-buildmode=c-archive或-buildmode=c-shared构建的 Go 代码将为上述异步信号安装信号处理程序，并保存任何现有信号处理程序。如果将信号传递到非 Go 线程，则将像上面描述的那样操作，除非存在现有的非 Go 信号处理程序，在引发信号之前将安装该处理程序。

#### Windows 

​	Windows 上，^C(Control-C)或^BREAK(Control-Break)通常会导致程序退出。如果对 os.Interrupt 调用了 Notify，则^C 或^BREAK 将导致在该通道上发送 os.Interrupt，并且程序不会退出。如果调用 Reset 或对于 Notify 传递的所有通道调用 Stop，则将恢复默认行为。

​	此外，如果调用了 Notify，并且 Windows 将 CTRL_CLOSE_EVENT、CTRL_LOGOFF_EVENT 或 CTRL_SHUTDOWN_EVENT 发送到进程，则 Notify 将返回 syscall.SIGTERM。与 Control-C 和 Control-Break 不同，当收到 CTRL_CLOSE_EVENT、CTRL_LOGOFF_EVENT 或 CTRL_SHUTDOWN_EVENT 时，Notify 不会更改进程行为——除非退出，否则进程仍将被终止。但是，收到 syscall.SIGTERM 将使进程有机会在终止之前清理。

#### Plan 9 

​	Plan 9 上，信号具有 syscall.Note 类型，它是一个字符串。使用 syscall.Note 调用 Notify 将导致在发布该字符串作为注释时将该值发送到通道。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Ignore  <- go1.5

``` go 
func Ignore(sig ...os.Signal)
```

​	Ignore函数将提供的信号设置为被忽略。如果程序收到这些信号，将不会发生任何事情。 Ignore函数会撤销之前对所提供信号的Notify调用的效果。如果没有提供信号，则会忽略所有传入的信号。

#### func Ignored  <- go1.11

``` go 
func Ignored(sig os.Signal) bool
```

​	Ignored函数报告sig信号当前是否被忽略。

#### func Notify 

``` go 
func Notify(c chan<- os.Signal, sig ...os.Signal)
```

​	Notify函数使得signal包将传入的信号转发到通道c。如果没有提供信号，则所有传入的信号都将被转发到c。否则，只转发提供的信号。

​	signal包不会阻止向c发送信号：调用者必须确保c具有足够的缓冲空间来跟上预期的信号率。对于用于通知一个信号值的通道，大小为1的缓冲区就足够了。

​	可以多次调用Notify函数并传入相同的通道，每次调用都会扩展该通道接收的信号集。从集合中删除信号的唯一方法是调用Stop函数。

​	可以多次调用Notify函数并传入不同的通道和相同的信号：每个通道都会独立地接收传入信号的副本。

##### Notify Example

``` go 
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// 设置信号通知的通道。
	// 必须使用带缓冲的通道，
    // 否则如果我们没有准备好接收信号，就可能会错过信号。
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// 阻塞直到接收到信号。
	s := <-c
	fmt.Println("Got signal:", s)
}

```

##### Notify Example (AllSignals) 

``` go 
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// 设置信号通知的通道。
	// 必须使用带缓冲的通道，
    // 否则如果我们没有准备好接收信号，就可能会错过信号。
	c := make(chan os.Signal, 1)

	// 将没有信号传递给 Notify 表示将所有信号发送到通道。
	signal.Notify(c)

	// 阻塞直到接收到任何信号。
	s := <-c
	fmt.Println("Got signal:", s)
}

```



#### func NotifyContext  <- go1.16

``` go 
func NotifyContext(parent context.Context, signals ...os.Signal) (ctx context.Context, stop context.CancelFunc)
```

​	NotifyContext函数返回一个父上下文的副本，当其中一个列出的信号到达、返回的stop函数被调用或父上下文的Done通道关闭时，该上下文将被标记为完成(其Done通道被关闭)。

​	stop函数取消注册信号行为，这与signal.Reset函数一样，可能会为给定信号恢复默认行为。例如，接收os.Interrupt的Go程序的默认行为是退出。调用NotifyContext(parent, os.Interrupt)将更改行为以取消返回的上下文。未来收到的中断不会触发默认(退出)行为，直到调用返回的stop函数。

​	stop函数释放与之关联的资源，因此代码应尽快在完成此上下文中运行的操作并且信号不再需要转发到上下文时调用stop函数。

##### NotifyContext Example

这个例子使用了带有信号的上下文，以便在接收到信号后告诉一个阻塞的函数它应该放弃它的工作。

``` go 
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Fatal(err)
	}

	// 在类Unix系统中，
    // 按下键盘上的 Ctrl+C 键会向正在执行的程序的进程发送 SIGINT 信号。
    // 这个示例通过向自身发送 SIGINT 信号来模拟这种情况。
	if err := p.Signal(os.Interrupt); err != nil {
		log.Fatal(err)
	}

	select {
	case <-time.After(time.Second):
		fmt.Println("missed signal")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // 输出 "context canceled"
		stop()                 // 尽快停止接收信号通知。
	}

}

```



#### func Reset  <- go1.5

``` go 
func Reset(sig ...os.Signal)
```

​	Reset函数将提供的信号的Notify调用效果撤销。如果没有提供信号，则将重置所有信号处理程序。

#### func Stop  <- go1.1

``` go 
func Stop(c chan<- os.Signal)
```

​	Stop函数使包 signal 停止将传入信号转发到 c。它会撤销所有先前使用 c 调用 Notify 的效果。当 Stop 返回时，保证 c 不会再收到任何信号。

## 类型

This section is empty.