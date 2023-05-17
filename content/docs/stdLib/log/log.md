+++
title = "log"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# log

[https://pkg.go.dev/log@go1.20.1](https://pkg.go.dev/log@go1.20.1)

​	log包实现了一个简单的日志包。它定义了一个类型Logger，该类型具有格式化输出的方法。它还有一个预定义的"标准"Logger，可以通过Print[f|ln]、Fatal[f|ln]和Panic[f|ln]帮助函数访问，这比手动创建Logger更容易。该记录器(logger )将写入标准错误，并打印每个记录消息的日期和时间。每条日志消息都在单独的一行上输出：如果要打印的消息没有以换行符结尾，该记录器(logger )将添加一个换行符。Fatal函数在写入日志消息后调用os.Exit(1)。Panic函数在写入日志消息后调用panic。


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=41)

```go linenums="1"
const (
	Ldate         = 1 << iota     // 当地时区的日期：2009/01/23
	Ltime                         // 当地时区的时间：01:23:23
	Lmicroseconds // 微秒级分辨率：01:23:23.123123.假设是Ltime。
	Llongfile    // 完整的文件名和行数：/a/b/c/d.go:23
	Lshortfile  // 最后的文件名元素和行数: d.go:23. 覆盖Llongfile。
	LUTC  // 如果设置了Ldate或Ltime，则使用UTC而不是本地时区。
	Lmsgprefix  // 将 "前缀"从行首移至信息之前
	LstdFlags     = Ldate | Ltime // 标准记录器的初始值
)
```

​	这些标志定义了记录器生成每个日志条目时要添加的文本。位或运算将它们连接在一起以控制打印内容。除了 Lmsgprefix 标志外，没有控制它们出现的顺序(按此处列出的顺序)或它们呈现的格式(如注释中所述)。只有当指定了 Llongfile 或 Lshortfile 时，前缀后面才跟有一个冒号。例如，标志 Ldate | Ltime(或 LstdFlags)会产生以下输出：

```
2009/01/23 01:23:23 message
```

而标志 Ldate | Ltime | Lmicroseconds | Llongfile 会产生以下输出：

```
2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
```

## 变量

This section is empty.

## 函数

#### func [Fatal](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=363) 

```go linenums="1"
func Fatal(v ...any)
```

​	Fatal函数等效于 Print()，然后调用 os.Exit(1)。

#### func [Fatalf](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=369) 

```go linenums="1"
func Fatalf(format string, v ...any)
```

​	Fatalf函数等效于 Printf()，然后调用 os.Exit(1)。

#### func [Fatalln](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=375) 

```go linenums="1"
func Fatalln(v ...any)
```

​	Fatalln函数等效于 Println()，然后调用 os.Exit(1)。

#### func [Flags](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=308) 

```go linenums="1"
func Flags() int
```

​	Flags函数返回标准日志记录器的输出标志。标志位是 Ldate、Ltime 等等。

#### func [Output](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=408)  <- go1.5

```go linenums="1"
func Output(calldepth int, s string) error
```

​	Output 方法用于写入日志事件的输出。字符串 s 包含在 Logger 的标志指定前缀之后要打印的文本。如果 s 的最后一个字符不是换行符，则会附加一个换行符。当设置 Llongfile 或 Lshortfile 标志时，calldepth 是要跳过的帧数，以计算文件名和行号；当 calldepth 的值为 1 时，将会打印 Output 的调用者的细节信息。

#### func [Panic](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=381) 

```go linenums="1"
func Panic(v ...any)
```

​	Panic函数等效于 Print()，然后调用 panic()。

#### func [Panicf](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=388) 

```go linenums="1"
func Panicf(format string, v ...any)
```

​	Panicf函数相当于调用Printf()后再调用panic()。

#### func [Panicln](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=395) 

```go linenums="1"
func Panicln(v ...any)
```

​	Panicln函数相当于调用Println()后再调用panic()。

#### func [Prefix](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=319) 

```go linenums="1"
func Prefix() string
```

​	Prefix函数返回标准记录器的输出前缀。

#### func [Print](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=337) 

```go linenums="1"
func Print(v ...any)
```

​	Print函数通过调用Output函数向标准记录器打印内容。参数的处理方式类似于fmt.Print。

#### func [Printf](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=346) 

```go linenums="1"
func Printf(format string, v ...any)
```

​	Printf函数通过调用Output函数向标准记录器打印内容。参数的处理方式类似于fmt.Printf。

#### func [Println](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=355) 

```go linenums="1"
func Println(v ...any)
```

​	Println函数通过调用Output函数向标准记录器打印内容。参数的处理方式类似于fmt.Println。

#### func [SetFlags](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=314) 

```go linenums="1"
func SetFlags(flag int)
```

​	SetFlags函数设置标准记录器的输出标志。标志位包括Ldate、Ltime等等。

#### func [SetOutput](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=302) 

```go linenums="1"
func SetOutput(w io.Writer)
```

​	SetOutput函数设置标准记录器的输出目的地。

#### func [SetPrefix](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=324) 

```go linenums="1"
func SetPrefix(prefix string)
```

​	SetPrefix函数设置标准记录器的输出前缀。

#### func [Writer](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=329)  <- go1.13

```go linenums="1"
func Writer() io.Writer
```

​	Writer函数返回标准记录器的输出目的地。

## 类型

### type [Logger](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=56) 

```go linenums="1"
type Logger struct {
    mu        sync.Mutex // 保证原子写入；保护以下字段
	prefix    string // 每行日志前缀以标识记录器(参见 Lmsgprefix)
	flag      int        // 属性
	out       io.Writer  // 输出目的地
	buf       []byte     // 用于累积要写入的文本
	isDiscard int32 // 原子布尔值：是否 out == io.Discard
}
```

​	Logger表示一个活动的日志记录对象，它生成输出行到一个io.Writer。每个日志操作都会调用Writer的Write方法。一个Logger可以同时从多个goroutine使用；它保证对Writer的访问是序列化的。

##### Logger Example

```go linenums="1"
package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
}
Output:

logger: example_test.go:19: Hello, log file!
```



#### func [Default](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=89)  <- go1.16

```go linenums="1"
func Default() *Logger
```

​	Default函数返回由包级别输出函数使用的标准记录器。

#### func [New](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=70) 

```go linenums="1"
func New(out io.Writer, prefix string, flag int) *Logger
```

​	New创建一个新的Logger。out变量设置将写入日志数据的目的地。prefix出现在每个生成的日志行的开头，或者如果提供了Lmsgprefix标志，则出现在日志标题之后。flag参数定义日志属性。

#### (*Logger) [Fatal](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=226) 

```go linenums="1"
func (l *Logger) Fatal(v ...any)
```

​	Fatal方法相当于 l.Print() 后面跟一个 os.Exit(1) 的调用。

#### (*Logger) [Fatalf](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=232) 

```go linenums="1"
func (l *Logger) Fatalf(format string, v ...any)
```

​	Fatalf方法相当于 l.Printf() 后面跟一个 os.Exit(1) 的调用。

#### (*Logger) [Fatalln](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=238) 

```go linenums="1"
func (l *Logger) Fatalln(v ...any)
```

​	Fatalln方法相当于 l.Println() 后面跟一个 os.Exit(1) 的调用。

#### (*Logger) [Flags](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=266) 

```go linenums="1"
func (l *Logger) Flags() int
```

​	Flags方法返回 Logger 的输出标志。标志位包括 Ldate、Ltime 等。

#### (*Logger) [Output](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=171) 

```go linenums="1"
func (l *Logger) Output(calldepth int, s string) error
```

​	Output方法写入日志事件的输出。字符串 s 包含在 Logger 的标志指定的前缀之后要打印的文本。如果 s 的最后一个字符不是换行符，则附加一个换行符。Calldepth 用于恢复 PC，并提供一般性，尽管目前在所有预定义的路径上它将是 2。

##### Output  Example

```go linenums="1"
package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)

		infof = func(info string) {
			logger.Output(2, info)
		}
	)

	infof("Hello world")

	fmt.Print(&buf)
}
Output:

INFO: example_test.go:36: Hello world
```



#### (*Logger) [Panic](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=244) 

```go linenums="1"
func (l *Logger) Panic(v ...any)
```

​	Panic方法相当于 l.Print() 后面跟一个 panic() 的调用。

#### (*Logger) [Panicf](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=251) 

```go linenums="1"
func (l *Logger) Panicf(format string, v ...any)
```

​	Panicf方法相当于 l.Printf() 后面跟一个 panic() 的调用。

#### (*Logger) [Panicln](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=258) 

```go linenums="1"
func (l *Logger) Panicln(v ...any)
```

​	Panicln方法相当于 l.Println() 后面跟一个 panic() 的调用。

#### (*Logger) [Prefix](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=281) 

```go linenums="1"
func (l *Logger) Prefix() string
```

​	Prefix方法返回 Logger 的输出前缀。

#### (*Logger) [Print](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=209) 

```go linenums="1"
func (l *Logger) Print(v ...any)
```

​	Print方法调用 l.Output 来向 Logger 打印信息。参数的处理方式类似于 fmt.Print。

#### (*Logger) [Printf](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=200) 

```go linenums="1"
func (l *Logger) Printf(format string, v ...any)
```

​	Printf方法调用 l.Output 来向 Logger 打印信息。参数的处理方式类似于 fmt.Printf。

#### (*Logger) [Println](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=218) 

```go linenums="1"
func (l *Logger) Println(v ...any)
```

​	Println方法调用 l.Output 来向 Logger 打印信息。参数的处理方式类似于 fmt.Println。

#### (*Logger) [SetFlags](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=274) 

```go linenums="1"
func (l *Logger) SetFlags(flag int)
```

​	SetFlags方法设置 Logger 的输出标志。标志位包括 Ldate、Ltime 等。

#### (*Logger) [SetOutput](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=79)  <- go1.5

```go linenums="1"
func (l *Logger) SetOutput(w io.Writer)
```

​	SetOutput方法设置 Logger 的输出目的地。

#### (*Logger) [SetPrefix](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=288) 

```go linenums="1"
func (l *Logger) SetPrefix(prefix string)
```

​	SetPrefix方法设置Logger 的输出前缀。

#### (*Logger) [Writer](https://cs.opensource.google/go/go/+/go1.20.1:src/log/log.go;l=295)  <- go1.12

```go linenums="1"
func (l *Logger) Writer() io.Writer
```

​	Writer方法返回Logger 的输出目的地。