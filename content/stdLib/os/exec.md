+++
title = "exec"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/os/exec@go1.20.1

​	exec包可以运行外部命令。它包装了 os.StartProcess，以便更轻松地重新映射 stdin 和 stdout，连接管道的 I/O，以及做其他调整。

​	与 C 语言和其他语言中的 "system" 库调用不同，os/exec 包故意不调用系统 shell，不扩展任何 glob 模式，也不处理 shell 通常执行的其他扩展、管道或重定向。该包的行为更像 C 的 "exec" 函数系列。为了扩展 glob 模式，请直接调用 shell，注意转义任何危险输入，或使用 path/filepath 包的 Glob 函数。为了扩展环境变量，请使用 os 包的 ExpandEnv。

​	请注意，此包中的示例假定 Unix 系统。它们可能无法在 Windows 上运行，并且它们不在 golang.org 和 godoc.org 使用的 Go Playground 中运行。

#### 当前目录中的可执行文件

​	函数 Command 和 LookPath 在列出的路径目录中查找程序，遵循主机操作系统的惯例。操作系统几十年来一直将当前目录包括在此搜索中，有时是隐式地，有时是默认配置为这样。现代实践是包括当前目录通常是意外的，而且经常导致安全问题。

​	为了避免这些安全问题，从 Go 1.19 开始，此包将不会使用相对于当前目录的隐式或显式路径条目解析程序。也就是说，如果您运行 exec.LookPath("go")，它将不会在 Unix 上成功返回 ./go，也不会在 Windows 上成功返回 .\go.exe，无论路径如何配置。相反，如果通常的路径算法会导致这样的答案，这些函数将返回一个满足 errors.Is(err, ErrDot) 的错误 err。

​	例如，考虑以下两个程序片段：

```
path, err := exec.LookPath("prog")
if err != nil {
	log.Fatal(err)
}
use(path)
```

和

```
cmd := exec.Command("prog")
if err := cmd.Run(); err != nil {
	log.Fatal(err)
}
```

​	无论当前路径如何配置，它们都无法找到并运行 `./prog` 或 `.\prog.exe`。

​	总是想从当前目录运行程序的代码可以重写为 "`./prog`"，而不是 "`prog`"。

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

​	设置环境变量GODEBUG=execerrdot=0将禁用生成ErrDot，为无法应用更有针对性的修复的程序暂时恢复Go 1.19之前的行为。将来的Go版本可能会删除对该变量的支持。

​	在添加这样的覆盖之前，请确保您了解这样做的安全性影响。有关更多信息，请参见https://go.dev/blog/path-security。


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/exec/exec.go;l=1303)

``` go 
var ErrDot = errors.New("cannot run executable found relative to current directory")
```

​	ErrDot 表示由于路径中存在 '.'(点号)，无论是隐式还是显式地，导致路径查找解析到了当前目录中的可执行文件。有关详细信息，请参阅包文档。

​	请注意，此包中的函数不会直接返回 ErrDot。代码应使用 errors.Is(err, ErrDot) 而不是 err == ErrDot 来测试返回的错误 err 是否是由于此条件而引起的。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/exec/lp_unix.go;l=20)

``` go 
var ErrNotFound = errors.New("executable file not found in $PATH")
```

​	ErrNotFound 是搜索路径失败未找到可执行文件时返回的错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/exec/exec.go;l=127)

``` go 
var ErrWaitDelay = errors.New("exec: WaitDelay expired before I/O complete")
```

​	如果进程以成功的状态码退出，但其输出管道在命令的 WaitDelay 过期之前未关闭，则 `(*Cmd).Wait` 返回 ErrWaitDelay。

## 函数

#### func LookPath 

``` go 
func LookPath(file string) (string, error)
```

​	LookPath函数在 PATH 环境变量指定的目录中查找名为 file 的可执行文件。如果 file 包含斜杠，则尝试直接使用它，不会查询 PATH。否则，如果成功，结果是绝对路径。

​	在较旧的 Go 版本中，LookPath 可能返回相对于当前目录的路径。从 Go 1.19 开始，如果常规的路径算法会导致该答案，则 LookPath 将返回该路径以及满足 errors.Is(err, ErrDot) 的错误。有关更多详细信息，请参阅包文档。

##### LookPath Example

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
	// Path 是要运行的命令的路径。
	//
	// 这是唯一必须设置为非零值的字段。
	// 如果 Path 是相对路径，则相对于 Dir 进行求值。
	Path string

	// Args 包含命令行参数，包括命令作为 Args[0]。
	// 如果 Args 字段为空或 nil，则 Run 使用 {Path}。
	//
	// 通常情况下，Path 和 Args 由调用 Command 来设置。
	Args []string

	// Env 指定进程的环境。
	// 每个条目的格式为 "key=value"。
	// 如果 Env 为 nil，则新进程使用当前进程的环境。
	// 如果 Env 包含重复的环境键，
    // 则只使用每个重复键的 slice 中的最后一个值。
	// 在 Windows 上，
    // 特殊情况下如果缺少 SYSTEMROOT 并且未显式设置为空字符串，
    // 则始终会添加它。
	Env []string

	// Dir 指定命令的工作目录。
	// 如果 Dir 是空字符串，则 Run 在调用进程的当前目录中运行命令。
	Dir string

	// Stdin 指定进程的标准输入。
	//
	// 如果 Stdin 是 nil，则进程从空设备(os.DevNull)读取。
	//
	// 如果 Stdin 是 *os.File，则进程的标准输入直接连接到该文件。
	//
	// 否则，在命令执行期间，
    // 单独的 goroutine 从 Stdin 读取数据并通过管道将其传递给命令。
	// 在这种情况下，Wait 不会完成，直到 goroutine 停止复制，
    // 要么是因为已经到达了 Stdin 的末尾(EOF 或读取错误)，
	// 要么是因为写入管道返回了错误，
    // 要么是因为设置了非零的 WaitDelay 并已过期。
	Stdin io.Reader

	// Stdout和Stderr指定进程的标准输出和错误。
	//
	// 如果它们中的任何一个为nil，
    // 则Run会将相应的文件描述符连接到null设备(os.DevNull)。
	//
	// 如果它们中的任何一个是*os.File，
    // 则进程的相应输出将直接连接到该文件。
	//
	// 否则，在执行命令期间，
    // 一个单独的goroutine会通过管道从进程读取数据，
    // 并将其传递到相应的Writer。
	// 在这种情况下，Wait不会完成，
    // 直到goroutine到达EOF或遇到错误或非零的WaitDelay到期。
	//
	// 如果Stdout和Stderr是同一个writer，
    // 并且具有可以与==比较的类型，
    // 则最多只有一个goroutine每次调用Write。
	Stdout io.Writer
	Stderr io.Writer

	// ExtraFiles指定要由新进程继承的其他打开文件。
    // 它不包括标准输入、标准输出或标准错误。
	// 如果非nil，则条目i变成文件描述符3+i。
	//
	// ExtraFiles在Windows上不受支持。
	ExtraFiles []*os.File

	// SysProcAttr保存可选的特定于操作系统的属性。
    // Run将其作为os.ProcAttr的Sys字段传递给os.StartProcess。
	SysProcAttr *syscall.SysProcAttr

	// Process是一旦启动就是底层进程。
	// 如果成功启动进程，Wait或Run将在命令完成时填充其ProcessState。
	Process *os.Process

	// ProcessState包含有关退出进程的信息。
	// 如果成功启动进程，Wait或Run将在命令完成时填充其ProcessState。
	ProcessState *os.ProcessState

	Err error // LookPath error, if any.// LookPath错误，如果有的话。

	// 如果Cancel非nil，则必须使用CommandContext创建命令，
    // 当命令的Context完成时，将调用Cancel。
	// 默认情况下，
    // CommandContext将Cancel设置为在命令的Process上调用Kill方法。
	//
	// 通常，自定义Cancel将向命令的Process发送一个信号，
    // 但它可能会采取其他措施来启动取消操作，
	// 例如关闭stdin或stdout管道或在网络套接字上发送关闭请求。
	//
	// 如果Cancel在调用后以成功状态退出，
    // 并且Cancel未返回与os.ErrProcessDone等效的错误，
	// 那么Wait和类似方法将返回非nil错误：
    // 包装由Cancel返回的错误或上下文的错误。
	// (如果命令以非成功状态退出，
    // 或者Cancel返回一个包装os.ErrProcessDone的错误，
	// Wait和类似方法将继续返回命令的常规退出状态。)
	//
	// 如果将Cancel设置为nil，
    // 则命令的Context完成时不会立即发生任何事情，
	// 但非零的WaitDelay仍然会生效。
    // 这可能是有用的，
    // 例如，为了解决不支持关闭信号但预计始终快速完成的命令中的死锁。
	//
	// 如果Start返回非nil错误，则不会调用Cancel。
	Cancel func() error

	// 如果WaitDelay不为零，
    // 则它限制了在Wait时等待两个意外延迟的来源所花费的时间：
	// 子进程在相关的上下文被取消后仍然无法退出，
    // 以及子进程退出但其I/O管道未关闭。
	//
	// 当关联的上下文完成或Wait的调用观察到子进程已退出时，
    // WaitDelay计时器启动，
	// 以确保在延迟时间后，命令将关闭子进程及/或其I/O管道。
	//
	// 如果子进程无法退出，
    // 可能是因为它忽略或未能接收来自取消函数的关闭信号，
	// 或因为未设置取消函数，则它将使用os.Process.Kill终止。
	//
	// 然后，如果与子进程通信的I/O管道仍然打开，
    // 则将关闭这些管道，
    // 以解除当前阻塞在Read或Write调用上的所有goroutine。
	//
	// 如果由于WaitDelay而关闭管道，
    // 没有发生取消调用，并且命令已以成功状态退出，
	// 则Wait等方法将返回ErrWaitDelay而不是nil。
	//
	// 如果WaitDelay为零(默认值)，
    // 则将读取I/O管道，直到EOF为止，
    // 这可能不会发生，直到命令的孤立子进程也关闭了它们的管道描述符。
	WaitDelay time.Duration
	// 包含已过滤或未导出的字段
}
```

​	Cmd 表示正在准备或运行的外部命令。

​	在调用其 Run、Output 或 CombinedOutput 方法之后，Cmd 不能被重用。

#### func Command 

``` go 
func Command(name string, arg ...string) *Cmd
```

​	Command函数返回一个 Cmd 结构，以给定参数执行指定的命令。

​	它只在返回的结构中设置 Path 和 Args 字段。

​	如果 name 不包含路径分隔符，则 Command 会尝试使用 LookPath 将 name 解析为完整路径。否则，它将直接使用 name 作为 Path。

​	返回的 Cmd 的 Args 字段由命令名称后跟 arg 的元素构成，因此 arg 不应包括命令本身的名称。例如，Command("echo", "hello")。Args[0] 总是 name，而不是可能解析后的 Path。

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

​	CommandContext函数类似于 Command，但包括一个 context。

​	提供的 context 在命令完成之前中止进程(通过调用 cmd.Cancel 或 os.Process.Kill)。

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

​	Run方法启动指定的命令并等待其完成。

​	如果命令运行良好，无法复制标准输入、标准输出和标准错误，并带有零退出状态，则返回的错误为nil。

​	如果命令启动但未成功完成，则返回的错误类型为*ExitError。其他错误类型可能会针对其他情况返回。

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

​	Start方法启动指定的命令，但不等待其完成。

​	如果Start成功返回，则c.Process字段将被设置。

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

​	StderrPipe方法返回一个管道，该管道将在命令启动时连接到命令的标准错误。

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

​	StdoutPipe方法返回一个管道，该管道将在命令启动时连接到命令的标准输出。

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

​	String方法返回c的人类可读描述。它仅用于调试。特别是，它不适合用作shell的输入。String的输出可能会因Go版本而异。

#### (*Cmd) Wait 

``` go 
func (c *Cmd) Wait() error
```

​	Wait方法等待命令退出，并等待从标准输入复制或从标准输出或标准错误复制完成。

​	该命令必须已由Start启动。

​	如果命令运行良好，无法复制标准输入、标准输出和标准错误，并带有零退出状态，则返回的错误为nil。

​	如果命令无法运行或未成功完成，则返回的错误类型为*ExitError。对于I/O问题，可能会返回其他错误类型。

​	如果c.Stdin、c.Stdout或c.Stderr中的任何一个不是*os.File，则Wait还将等待复制到或从进程中的相应I/O循环完成。

​	Wait释放与Cmd关联的任何资源。

### type Error 

``` go 
type Error struct {
	// Name 是出现错误的文件名。
	Name string
	// Err 是底层错误。
	Err error
}
```

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

	// Stderr 保存 Cmd.Output 方法的一部分标准错误输出，
    // 如果未收集标准错误输出，则会保存该部分。
	//
	// 如果错误输出很长，则 Stderr 可能只包含输出的前缀和后缀，
    // 其中间部分被替换为有关省略字节数的文本。
	//
	// Stderr 用于调试和包含在错误消息中。
    // 有其他需求的用户应根据需要重定向 Cmd.Stderr。
	Stderr []byte
}
```

​	ExitError 报告命令未能成功退出。

#### (*ExitError) Error 

``` go 
func (e *ExitError) Error() string
```