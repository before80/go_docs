+++
title = "gproc"
date = 2024-03-21T17:56:37+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gproc

Package gproc implements management and communication for processes.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func AddSigHandler 

``` go
func AddSigHandler(handler SigHandler, signals ...os.Signal)
```

AddSigHandler adds custom signal handler for custom one or more signals.

##### func AddSigHandlerShutdown 

``` go
func AddSigHandlerShutdown(handler ...SigHandler)
```

AddSigHandlerShutdown adds custom signal handler for shutdown signals: syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGABRT.

##### func IsChild 

``` go
func IsChild() bool
```

IsChild checks and returns whether current process is a child process. A child process is forked by another gproc process.

##### func Listen 

``` go
func Listen()
```

Listen blocks and does signal listening and handling.

##### func MustShell 

``` go
func MustShell(ctx context.Context, cmd string, out io.Writer, in io.Reader)
```

MustShell performs as Shell, but it panics if any error occurs.

##### func MustShellExec 

``` go
func MustShellExec(ctx context.Context, cmd string, environment ...[]string) string
```

MustShellExec performs as ShellExec, but it panics if any error occurs.

##### func MustShellRun 

``` go
func MustShellRun(ctx context.Context, cmd string)
```

MustShellRun performs as ShellRun, but it panics if any error occurs.

##### func PPid 

``` go
func PPid() int
```

PPid returns the custom parent pid if exists, or else it returns the system parent pid.

##### func PPidOS 

``` go
func PPidOS() int
```

PPidOS returns the system parent pid of current process. Note that the difference between PPidOS and PPid function is that the PPidOS returns the system ppid, but the PPid functions may return the custom pid by gproc if the custom ppid exists.

##### func Pid 

``` go
func Pid() int
```

Pid returns the pid of current process.

##### func SearchBinary 

``` go
func SearchBinary(file string) string
```

SearchBinary searches the binary `file` in current working folder and PATH environment.

##### func SearchBinaryPath 

``` go
func SearchBinaryPath(file string) string
```

SearchBinaryPath searches the binary `file` in PATH environment.

##### func Send 

``` go
func Send(pid int, data []byte, group ...string) error
```

Send sends data to specified process of given pid.

##### func SetPPid 

``` go
func SetPPid(ppid int) error
```

SetPPid sets custom parent pid for current process.

##### func Shell 

``` go
func Shell(ctx context.Context, cmd string, out io.Writer, in io.Reader) error
```

Shell executes command `cmd` synchronously with given input pipe `in` and output pipe `out`. The command `cmd` reads the input parameters from input pipe `in`, and writes its output automatically to output pipe `out`.

##### func ShellExec 

``` go
func ShellExec(ctx context.Context, cmd string, environment ...[]string) (result string, err error)
```

ShellExec executes given command `cmd` synchronously and returns the command result.

##### func ShellRun 

``` go
func ShellRun(ctx context.Context, cmd string) error
```

ShellRun executes given command `cmd` synchronously and outputs the command result to the stdout.

##### func StartTime 

``` go
func StartTime() time.Time
```

StartTime returns the start time of current process.

##### func Uptime 

``` go
func Uptime() time.Duration
```

Uptime returns the duration which current process has been running

### Types 

#### type Manager 

``` go
type Manager struct {
	// contains filtered or unexported fields
}
```

Manager is a process manager maintaining multiple processes.

##### func NewManager 

``` go
func NewManager() *Manager
```

NewManager creates and returns a new process manager.

##### (*Manager) AddProcess 

``` go
func (m *Manager) AddProcess(pid int)
```

AddProcess adds a process to current manager. It does nothing if the process with given `pid` does not exist.

##### (*Manager) Clear 

``` go
func (m *Manager) Clear()
```

Clear removes all processes in current manager.

##### (*Manager) GetProcess 

``` go
func (m *Manager) GetProcess(pid int) *Process
```

GetProcess retrieves and returns a Process object. It returns nil if it does not find the process with given `pid`.

##### (*Manager) KillAll 

``` go
func (m *Manager) KillAll() error
```

KillAll kills all processes in current manager.

##### (*Manager) NewProcess 

``` go
func (m *Manager) NewProcess(path string, args []string, environment []string) *Process
```

NewProcess creates and returns a Process object.

##### (*Manager) Pids 

``` go
func (m *Manager) Pids() []int
```

Pids retrieves and returns all process id array in current manager.

##### (*Manager) Processes 

``` go
func (m *Manager) Processes() []*Process
```

Processes retrieves and returns all processes in current manager.

##### (*Manager) RemoveProcess 

``` go
func (m *Manager) RemoveProcess(pid int)
```

RemoveProcess removes a process from current manager.

##### (*Manager) Send 

``` go
func (m *Manager) Send(data []byte)
```

Send sends data bytes to all processes in current manager.

##### (*Manager) SendTo 

``` go
func (m *Manager) SendTo(pid int, data []byte) error
```

SendTo sneds data bytes to specified processe in current manager.

##### (*Manager) SignalAll 

``` go
func (m *Manager) SignalAll(sig os.Signal) error
```

SignalAll sends a signal `sig` to all processes in current manager.

##### (*Manager) Size 

``` go
func (m *Manager) Size() int
```

Size returns the size of processes in current manager.

##### (*Manager) WaitAll 

``` go
func (m *Manager) WaitAll()
```

WaitAll waits until all process exit.

#### type MsgRequest 

``` go
type MsgRequest struct {
	SenderPid   int    // Sender PID.
	ReceiverPid int    // Receiver PID.
	Group       string // Message group name.
	Data        []byte // Request data.
}
```

MsgRequest is the request structure for process communication.

##### func Receive 

``` go
func Receive(group ...string) *MsgRequest
```

Receive blocks and receives message from other process using local TCP listening. Note that, it only enables the TCP listening service when this function called.

#### type MsgResponse 

``` go
type MsgResponse struct {
	Code    int    // 1: OK; Other: Error.
	Message string // Response message.
	Data    []byte // Response data.
}
```

MsgResponse is the response structure for process communication.

#### type Process 

``` go
type Process struct {
	exec.Cmd
	Manager *Manager
	PPid    int
}
```

Process is the struct for a single process.

##### func NewProcess 

``` go
func NewProcess(path string, args []string, environment ...[]string) *Process
```

NewProcess creates and returns a new Process.

##### func NewProcessCmd 

``` go
func NewProcessCmd(cmd string, environment ...[]string) *Process
```

NewProcessCmd creates and returns a process with given command and optional environment variable array.

##### (*Process) Kill 

``` go
func (p *Process) Kill() (err error)
```

Kill causes the Process to exit immediately.

##### (*Process) Pid 

``` go
func (p *Process) Pid() int
```

Pid retrieves and returns the PID for the process.

##### (*Process) Release 

``` go
func (p *Process) Release() error
```

Release releases any resources associated with the Process p, rendering it unusable in the future. Release only needs to be called if Wait is not.

##### (*Process) Run 

``` go
func (p *Process) Run(ctx context.Context) error
```

Run executes the process in blocking way.

##### (*Process) Send 

``` go
func (p *Process) Send(data []byte) error
```

Send sends custom data to the process.

##### (*Process) Signal 

``` go
func (p *Process) Signal(sig os.Signal) error
```

Signal sends a signal to the Process. Sending Interrupt on Windows is not implemented.

##### (*Process) Start 

``` go
func (p *Process) Start(ctx context.Context) (int, error)
```

Start starts executing the process in non-blocking way. It returns the pid if success, or else it returns an error.

#### type SigHandler 

``` go
type SigHandler func(sig os.Signal)
```

SigHandler defines a function type for signal handling.