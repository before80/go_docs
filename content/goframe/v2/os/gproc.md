+++
title = "gproc"
date = 2024-03-21T17:56:37+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gproc](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gproc)

Package gproc implements management and communication for processes.

​	软件包 gproc 实现流程的管理和通信。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func AddSigHandler

```go
func AddSigHandler(handler SigHandler, signals ...os.Signal)
```

AddSigHandler adds custom signal handler for custom one or more signals.

​	AddSigHandler 为自定义一个或多个信号添加自定义信号处理程序。

#### func AddSigHandlerShutdown

```go
func AddSigHandlerShutdown(handler ...SigHandler)
```

AddSigHandlerShutdown adds custom signal handler for shutdown signals: syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGABRT.

​	AddSigHandlerShutdown 为关机信号添加自定义信号处理程序：syscall。SIGINT，系统调用。SIGQUIT，系统调用。SIGKILL，系统调用。SIGTERM，系统调用。西加布特。

#### func IsChild

```go
func IsChild() bool
```

IsChild checks and returns whether current process is a child process. A child process is forked by another gproc process.

​	IsChild 检查并返回当前进程是否为子进程。子进程由另一个 gproc 进程分叉。

#### func Listen

```go
func Listen()
```

Listen blocks and does signal listening and handling.

​	监听块并发出监听和处理的信号。

#### func MustShell

```go
func MustShell(ctx context.Context, cmd string, out io.Writer, in io.Reader)
```

MustShell performs as Shell, but it panics if any error occurs.

​	MustShell 以 Shell 的形式执行，但如果发生任何错误，它会崩溃。

#### func MustShellExec

```go
func MustShellExec(ctx context.Context, cmd string, environment ...[]string) string
```

MustShellExec performs as ShellExec, but it panics if any error occurs.

​	MustShellExec 以 ShellExec 的形式执行，但如果发生任何错误，它会崩溃。

#### func MustShellRun

```go
func MustShellRun(ctx context.Context, cmd string)
```

MustShellRun performs as ShellRun, but it panics if any error occurs.

​	MustShellRun 以 ShellRun 的形式执行，但如果发生任何错误，它会崩溃。

#### func PPid

```go
func PPid() int
```

PPid returns the custom parent pid if exists, or else it returns the system parent pid.

​	PPid 返回自定义父 pid（如果存在），否则返回系统父 pid。

#### func PPidOS

```go
func PPidOS() int
```

PPidOS returns the system parent pid of current process. Note that the difference between PPidOS and PPid function is that the PPidOS returns the system ppid, but the PPid functions may return the custom pid by gproc if the custom ppid exists.

​	PPidOS 返回当前进程的系统父 pid。请注意，PPidOS 和 PPid 函数之间的区别在于 PPidOS 返回系统 ppid，但如果自定义 ppid 存在，PPid 函数可能会通过 gproc 返回自定义 pid。

#### func Pid

```go
func Pid() int
```

Pid returns the pid of current process.

​	Pid 返回当前进程的 pid。

#### func SearchBinary

```go
func SearchBinary(file string) string
```

SearchBinary searches the binary `file` in current working folder and PATH environment.

​	SearchBinary 在当前工作文件夹和 PATH 环境中搜索二进制文件 `file` 。

#### func SearchBinaryPath

```go
func SearchBinaryPath(file string) string
```

SearchBinaryPath searches the binary `file` in PATH environment.

​	SearchBinaryPath 在 PATH 环境中搜索二进制文件 `file` 。

#### func Send

```go
func Send(pid int, data []byte, group ...string) error
```

Send sends data to specified process of given pid.

​	发送 将数据发送到给定 pid 的指定进程。

#### func SetPPid

```go
func SetPPid(ppid int) error
```

SetPPid sets custom parent pid for current process.

​	SetPPid 为当前进程设置自定义父 pid。

#### func Shell

```go
func Shell(ctx context.Context, cmd string, out io.Writer, in io.Reader) error
```

Shell executes command `cmd` synchronously with given input pipe `in` and output pipe `out`. The command `cmd` reads the input parameters from input pipe `in`, and writes its output automatically to output pipe `out`.

​	Shell 与给定的输入管道 `in` 和输出管道同步执行命令 `out` `cmd` 。该命令 `cmd` 从输入管道读取输入参数 `in` ，并将其输出自动写入输出管道 `out` 。

#### func ShellExec

```go
func ShellExec(ctx context.Context, cmd string, environment ...[]string) (result string, err error)
```

ShellExec executes given command `cmd` synchronously and returns the command result.

​	ShellExec 同步执行给定的命令 `cmd` 并返回命令结果。

#### func ShellRun

```go
func ShellRun(ctx context.Context, cmd string) error
```

ShellRun executes given command `cmd` synchronously and outputs the command result to the stdout.

​	ShellRun 同步执行给定的命令 `cmd` ，并将命令结果输出到 stdout。

#### func StartTime

```go
func StartTime() time.Time
```

StartTime returns the start time of current process.

​	StartTime 返回当前进程的开始时间。

#### func Uptime

```go
func Uptime() time.Duration
```

Uptime returns the duration which current process has been running

​	Uptime 返回当前进程运行的持续时间

## 类型

### type Manager

```go
type Manager struct {
	// contains filtered or unexported fields
}
```

Manager is a process manager maintaining multiple processes.

​	管理器是维护多个进程的进程管理器。

#### func NewManager

```go
func NewManager() *Manager
```

NewManager creates and returns a new process manager.

​	NewManager 创建并返回一个新的进程管理器。

#### (*Manager) AddProcess

```go
func (m *Manager) AddProcess(pid int)
```

AddProcess adds a process to current manager. It does nothing if the process with given `pid` does not exist.

​	AddProcess 将进程添加到当前管理器。如果给定 `pid` 的进程不存在，则它不执行任何操作。

#### (*Manager) Clear

```go
func (m *Manager) Clear()
```

Clear removes all processes in current manager.

​	清除将删除当前管理器中的所有进程。

#### (*Manager) GetProcess

```go
func (m *Manager) GetProcess(pid int) *Process
```

GetProcess retrieves and returns a Process object. It returns nil if it does not find the process with given `pid`.

​	GetProcess 检索并返回一个 Process 对象。如果它没有找到给定 `pid` 的进程，则返回 nil。

#### (*Manager) KillAll

```go
func (m *Manager) KillAll() error
```

KillAll kills all processes in current manager.

​	KillAll 会终止当前管理器中的所有进程。

#### (*Manager) NewProcess

```go
func (m *Manager) NewProcess(path string, args []string, environment []string) *Process
```

NewProcess creates and returns a Process object.

​	NewProcess 创建并返回一个 Process 对象。

#### (*Manager) Pids

```go
func (m *Manager) Pids() []int
```

Pids retrieves and returns all process id array in current manager.

​	Pids 检索并返回当前管理器中的所有进程 ID 数组。

#### (*Manager) Processes

```go
func (m *Manager) Processes() []*Process
```

Processes retrieves and returns all processes in current manager.

​	进程检索并返回当前管理器中的所有进程。

#### (*Manager) RemoveProcess

```go
func (m *Manager) RemoveProcess(pid int)
```

RemoveProcess removes a process from current manager.

​	RemoveProcess 从当前管理器中删除进程。

#### (*Manager) Send

```go
func (m *Manager) Send(data []byte)
```

Send sends data bytes to all processes in current manager.

​	发送将数据字节发送到当前管理器中的所有进程。

#### (*Manager) SendTo

```go
func (m *Manager) SendTo(pid int, data []byte) error
```

SendTo sneds data bytes to specified processe in current manager.

​	SendTo 将数据字节发送到当前管理器中的指定进程。

#### (*Manager) SignalAll

```go
func (m *Manager) SignalAll(sig os.Signal) error
```

SignalAll sends a signal `sig` to all processes in current manager.

​	SignalAll 向当前管理器中的所有进程发送信号 `sig` 。

#### (*Manager) Size

```go
func (m *Manager) Size() int
```

Size returns the size of processes in current manager.

​	Size 返回当前管理器中进程的大小。

#### (*Manager) WaitAll

```go
func (m *Manager) WaitAll()
```

WaitAll waits until all process exit.

​	WaitAll 等待所有进程退出。

### type MsgRequest

```go
type MsgRequest struct {
	SenderPid   int    // Sender PID.
	ReceiverPid int    // Receiver PID.
	Group       string // Message group name.
	Data        []byte // Request data.
}
```

MsgRequest is the request structure for process communication.

​	MsgRequest 是进程通信的请求结构。

#### func Receive

```go
func Receive(group ...string) *MsgRequest
```

Receive blocks and receives message from other process using local TCP listening. Note that, it only enables the TCP listening service when this function called.

​	使用本地 TCP 侦听接收块并接收来自其他进程的消息。请注意，它仅在调用此函数时启用 TCP 侦听服务。

### type MsgResponse

```go
type MsgResponse struct {
	Code    int    // 1: OK; Other: Error.
	Message string // Response message.
	Data    []byte // Response data.
}
```

MsgResponse is the response structure for process communication.

​	MsgResponse 是进程通信的响应结构。

### type Process

```go
type Process struct {
	exec.Cmd
	Manager *Manager
	PPid    int
}
```

Process is the struct for a single process.

​	Process 是单个进程的结构。

#### func NewProcess

```go
func NewProcess(path string, args []string, environment ...[]string) *Process
```

NewProcess creates and returns a new Process.

​	NewProcess 创建并返回一个新 Process。

#### func NewProcessCmd

```go
func NewProcessCmd(cmd string, environment ...[]string) *Process
```

NewProcessCmd creates and returns a process with given command and optional environment variable array.

​	NewProcessCmd 使用给定的命令和可选的环境变量数组创建并返回进程。

#### (*Process) Kill

```go
func (p *Process) Kill() (err error)
```

Kill causes the Process to exit immediately.

​	Kill 会导致进程立即退出。

#### (*Process) Pid

```go
func (p *Process) Pid() int
```

Pid retrieves and returns the PID for the process.

​	Pid 检索并返回进程的 PID。

#### (*Process) Release

```go
func (p *Process) Release() error
```

Release releases any resources associated with the Process p, rendering it unusable in the future. Release only needs to be called if Wait is not.

​	Release 会释放与 Process p 关联的任何资源，使其将来无法使用。仅当 Wait 不是时，才需要调用 Release。

#### (*Process) Run

```go
func (p *Process) Run(ctx context.Context) error
```

Run executes the process in blocking way.

​	Run 以阻塞方式执行进程。

#### (*Process) Send

```go
func (p *Process) Send(data []byte) error
```

Send sends custom data to the process.

​	发送将自定义数据发送到进程。

#### (*Process) Signal

```go
func (p *Process) Signal(sig os.Signal) error
```

Signal sends a signal to the Process. Sending Interrupt on Windows is not implemented.

​	信号向进程发送信号。未实现在 Windows 上发送中断。

#### (*Process) Start

```go
func (p *Process) Start(ctx context.Context) (int, error)
```

Start starts executing the process in non-blocking way. It returns the pid if success, or else it returns an error.

​	Start 开始以非阻塞方式执行进程。如果成功，它将返回 pid，否则返回错误。

### type SigHandler

```go
type SigHandler func(sig os.Signal)
```

SigHandler defines a function type for signal handling.

​	SigHandler 定义了用于信号处理的函数类型。