+++
title = "数据竞争检测器"
weight = 21
date = 2023-05-18T17:31:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Data Race Detector - 数据竞争检测器

> 原文：[https://go.dev/doc/articles/race_detector](https://go.dev/doc/articles/race_detector)

## 简介

​	数据竞争是并发系统中最常见和最难以调试的错误类型之一。当两个 Goroutine 并发访问同一个变量并且至少一个访问是写入操作时，就会发生数据竞争。有关详情，请参见[《Go 内存模型》](../References/TheGoMemoryModel)。

​	以下是导致崩溃和内存损坏的数据竞争示例：

```go linenums="1"
func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // 第一次冲突的访问。
		c <- true
	}()
	m["2"] = "b" // 第二次冲突的访问。
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
```

## 使用方法

​	为了帮助诊断此类错误，Go 包括一个内置的数据竞争检测器。要使用它，请将 `-race` 标志添加到 go 命令中：

```sh
$ go test -race mypkg    // 测试包
$ go run -race mysrc.go  // 运行源文件
$ go build -race mycmd   // 构建命令
$ go install -race mypkg // 安装包
```

## 报告格式

​	当数据竞争检测器在程序中发现数据竞争时，它会打印一个报告。该报告包含冲突访问的栈跟踪，以及涉及的 Goroutine 创建的栈。以下是一个示例：

```
WARNING: DATA RACE
Read by goroutine 185:
  net.(*pollServer).AddFD()
      src/net/fd_unix.go:89 +0x398
  net.(*pollServer).WaitWrite()
      src/net/fd_unix.go:247 +0x45
  net.(*netFD).Write()
      src/net/fd_unix.go:540 +0x4d4
  net.(*conn).Write()
      src/net/net.go:129 +0x101
  net.func·060()
      src/net/timeout_test.go:603 +0xaf

Previous write by goroutine 184:
  net.setWriteDeadline()
      src/net/sockopt_posix.go:135 +0xdf
  net.setDeadline()
      src/net/sockopt_posix.go:144 +0x9c
  net.(*conn).SetDeadline()
      src/net/net.go:161 +0xe3
  net.func·061()
      src/net/timeout_test.go:616 +0x3ed

Goroutine 185 (running) created at:
  net.func·061()
      src/net/timeout_test.go:609 +0x288

Goroutine 184 (running) created at:
  net.TestProlongTimeout()
      src/net/timeout_test.go:618 +0x298
  testing.tRunner()
      src/testing/testing.go:301 +0xe8
```

## 选项

​	`GORACE`环境变量设置了竞态检测器选项，格式如下：

```
GORACE="option1=val1 option2=val2"
```

选项如下：

- `log_path`(默认值`stderr`)：竞态检测器将其报告写入名为`log_path.pid`的文件中。特殊名称`stdout`和`stderr`会分别将报告写入标准输出和标准错误。
- `exitcode`(默认值`66`)：在检测到竞争后退出时使用的退出状态。
- `strip_path_prefix`(默认值`""`)：从所有报告的文件路径中删除该前缀，以使报告更加简洁。
- `history_size`(默认值`1`)：每个goroutine的内存访问历史记录为$32K * 2 ^{history\_size}$元素。增加此值可以避免在报告中出现"failed to restore the stack(无法恢复栈)" 错误，但会增加内存使用量。
- `halt_on_error`(默认值`0`)：控制程序在报告第一个数据竞争后是否退出。
- `atexit_sleep_ms`(默认值`1000`)：主goroutine在退出前休眠的毫秒数。

例如：

```sh
$ GORACE="log_path=/tmp/race/report strip_path_prefix=/my/go/sources/" go test -race
```

## 排除测试

​	当使用`-race`标志构建时，`go`命令会定义附加的[构建标签](https://go.dev/pkg/go/build/#hdr-Build_Constraints)race。您可以使用该标签来在运行竞态检测器时排除某些代码和测试。以下是一些示例：

```go linenums="1"
// +build !race

package foo

// The test contains a data race. See issue 123.
func TestFoo(t *testing.T) {
	// ...
}

// The test fails under the race detector due to timeouts.
func TestBar(t *testing.T) {
	// ...
}

// The test takes too long under the race detector.
func TestBaz(t *testing.T) {
	// ...
}
```

## 如何使用

​	首先使用竞态检测器运行测试(`go test -race`)。竞态检测器仅能发现发生在运行时的竞争，因此无法发现未执行的代码路径中的竞争。如果您的测试覆盖不完整，则可以通过在真实工作负载下运行使用`-race`构建的二进制文件来发现更多竞争。

## 典型的数据竞争

​	以下是一些典型的数据竞争。所有这些都可以被竞态检测器检测出来。

### 循环计数器竞争

```go linenums="1"
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) //  这里的 'i' 和你想要的不一样。
			wg.Done()
		}()
	}
	wg.Wait()
}
```

​	函数字面值中的变量 `i` 是循环中使用的相同变量，因此在 goroutine 中读取与循环增量的竞争。(此程序通常会打印 55555，而不是 01234)。可以通过创建变量副本来修复程序：

```go linenums="1"
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(j) // 好的。读取了循环计数器的本地副本。
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

### 意外共享变量

```go linenums="1"
// ParallelWrite 函数将数据写入 file1 和 file2 文件，并返回错误。
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// 这个 err 是与主 goroutine 共享的，
			// 所以写入与下面的写入相互竞争。
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // 第二个冲突的写入 err。
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}
```

​	修复方法是在 goroutine 中引入新变量(注意使用 `:=`)：

```go linenums="1"
			...
			_, err := f1.Write(data)
			...
			_, err := f2.Write(data)
			...
```

### 未受保护的全局变量

​	如果从多个 goroutine 调用以下代码，则会在 `service` map 上发生竞争。同一个 map 的并发读写是不安全的：

```go linenums="1"
var service map[string]net.Addr

func RegisterService(name string, addr net.Addr) {
	service[name] = addr
}

func LookupService(name string) net.Addr {
	return service[name]
}
```

​	为了使代码安全，使用 mutex 保护访问：

```go linenums="1"
var (
	service   map[string]net.Addr
	serviceMu sync.Mutex
)

func RegisterService(name string, addr net.Addr) {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	service[name] = addr
}

func LookupService(name string) net.Addr {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return service[name]
}
```

### 原始的未受保护的变量

​	原始数据类型的变量也可能发生数据竞争(例如`bool`、`int`、`int64`等)，如下所示：

```go linenums="1"
type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // 第一个冲突的访问。
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// 第二个冲突的访问。
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
```

​	即使是这种"innocent(无害)"的数据竞争也可能导致难以调试的问题，这些问题可能是由于内存访问的非原子性、与编译器优化的干扰或访问处理器内存的重新排序问题引起的。

​	这种竞态的典型解决方法是使用通道或互斥锁。为了保持无锁行为，也可以使用[sync/atomic](../StdLib/sync/atomic/)包。

```go linenums="1"
type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano())
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			}
		}
	}()
}
```

### 未同步的发送和关闭操作

​	正如这个例子所演示的那样，在同一个通道上未同步的发送和关闭操作也可能导致竞争条件：

```go linenums="1"
c := make(chan struct{}) // 或缓冲通道

// 竞争检测器无法推导出以下发送和关闭操作的发生顺序。
// 这两个操作是未同步的，会同时发生。
go func() { c <- struct{}{} }()
close(c)
```

​	根据Go内存模型，通道上的发送在对该通道的相应接收完成之前发生。要同步发送和关闭操作，请使用确保发送完成后再关闭的接收操作：

```go linenums="1" hl_lines="4 4"
c := make(chan struct{}) // 或缓冲通道

go func() { c <- struct{}{} }()
<-c
close(c)
```

## 要求

​	竞态检测器需要启用cgo，并且在非Darwin系统上需要安装C编译器。竞争检测器支持linux/amd64、linux/ppc64le、linux/arm64、freebsd/amd64、netbsd/amd64、darwin/amd64、darwin/arm64和windows/amd64。

## 运行时开销

​	竞争检测的成本因程序而异，但对于典型的程序，内存使用可能增加5-10倍，执行时间可能增加2-20倍。

​	竞争检测器当前为每个 `defer` 和 `recover` 语句分配额外的 8 个字节。这些额外的分配[在 goroutine  退出之前不会被恢复](https://go.dev/issue/26813)。这意味着，如果您有一个长时间运行的 goroutine，它会定期发出 `defer` 和 `recover`  调用，程序的内存使用量可能会无限增长。这些内存分配不会出现在 `runtime.ReadMemStats` 或 `runtime/pprof`  的输出中。