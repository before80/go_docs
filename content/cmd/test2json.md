+++
title = "test2json"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# test2json

> 原文：[https://pkg.go.dev/cmd/test2json@go1.19.3](https://pkg.go.dev/cmd/test2json@go1.19.3)

### Overview

Test2json converts go test output to a machine-readable JSON stream.

​	Test2json 将 go 测试输出转换为可供机器读取的 JSON 流。

Usage:

​	用法：

```
go tool test2json [-p pkg] [-t] [./pkg.test -test.v [-test.paniconexit0]]
```

Test2json runs the given test command and converts its output to JSON; with no command specified, test2json expects test output on standard input. It writes a corresponding stream of JSON events to standard output. There is no unnecessary input or output buffering, so that the JSON stream can be read for “live updates” of test status.

​	Test2json 运行给定的测试命令并将输出转换为 JSON；未指定命令时，test2json 期望标准输入上的测试输出。它将相应的 JSON 事件流写入标准输出。没有不必要的输入或输出缓冲，因此可以读取 JSON 流以获取测试状态的“实时更新”。

The -p flag sets the package reported in each test event.

​	-p 标志设置每个测试事件中报告的包。

The -t flag requests that time stamps be added to each test event.

​	-t 标志请求将时间戳添加到每个测试事件。

The test must be invoked with -test.v. Additionally passing -test.paniconexit0 will cause test2json to exit with a non-zero status if one of the tests being run calls os.Exit(0).

​	必须使用 -test.v 调用测试。另外传递 -test.paniconexit0 将导致 test2json 在运行的测试之一调用 os.Exit(0) 时退出并显示非零状态。

Note that test2json is only intended for converting a single test binary's output. To convert the output of a "go test" command, use "go test -json" instead of invoking test2json directly.

​	请注意，test2json 仅用于转换单个测试二进制文件的输出。要转换“go test”命令的输出，请使用“go test -json”，而不是直接调用 test2json。

#### Output Format [¶](https://pkg.go.dev/cmd/test2json@go1.19.3#hdr-Output_Format)

The JSON stream is a newline-separated sequence of TestEvent objects corresponding to the Go struct:

​	JSON 流是与 Go 结构对应的 TestEvent 对象的新行分隔序列：

```
type TestEvent struct {
	Time    time.Time // encodes as an RFC3339-format string
	Action  string
	Package string
	Test    string
	Elapsed float64 // seconds
	Output  string
}
```

The Time field holds the time the event happened. It is conventionally omitted for cached test results.

​	Time 字段保存事件发生的时间。对于缓存的测试结果，通常会省略它。

The Action field is one of a fixed set of action descriptions:

​	Action 字段是固定的一组操作描述之一：

```
run    - the test has started running
pause  - the test has been paused
cont   - the test has continued running
pass   - the test passed
bench  - the benchmark printed log output but did not fail
fail   - the test or benchmark failed
output - the test printed output
skip   - the test was skipped or the package contained no tests
```

The Package field, if present, specifies the package being tested. When the go command runs parallel tests in -json mode, events from different tests are interlaced; the Package field allows readers to separate them.

​	如果存在，Package 字段指定正在测试的包。当 go 命令以 -json 模式运行并行测试时，来自不同测试的事件会交织在一起；Package 字段允许读者将它们分开。

The Test field, if present, specifies the test, example, or benchmark function that caused the event. Events for the overall package test do not set Test.

​	如果存在，Test 字段指定导致事件的测试、示例或基准函数。整个包测试的事件不会设置 Test。

The Elapsed field is set for "pass" and "fail" events. It gives the time elapsed for the specific test or the overall package test that passed or failed.

​	Elapsed 字段针对“pass”和“fail”事件设置。它给出通过或失败的特定测试或整个包测试所经过的时间。

The Output field is set for Action == "output" and is a portion of the test's output (standard output and standard error merged together). The output is unmodified except that invalid UTF-8 output from a test is coerced into valid UTF-8 by use of replacement characters. With that one exception, the concatenation of the Output fields of all output events is the exact output of the test execution.

​	Output 字段针对 Action == "output" 设置，并且是测试输出（标准输出和标准错误合并在一起）的一部分。输出未经修改，但来自测试的无效 UTF-8 输出通过使用替换字符强制转换为有效的 UTF-8。除了这一例外，所有输出事件的 Output 字段的连接是测试执行的确切输出。

When a benchmark runs, it typically produces a single line of output giving timing results. That line is reported in an event with Action == "output" and no Test field. If a benchmark logs output or reports a failure (for example, by using b.Log or b.Error), that extra output is reported as a sequence of events with Test set to the benchmark name, terminated by a final event with Action == "bench" or "fail". Benchmarks have no events with Action == "run", "pause", or "cont".

​	当基准运行时，它通常会生成一行输出，给出计时结果。该行在事件中报告，其中 Action == "output" 且没有 Test 字段。如果基准记录输出或报告失败（例如，通过使用 b.Log 或 b.Error），则该额外输出将作为一系列事件报告，其中 Test 设置为基准名称，并以 Action == "bench" 或 "fail" 的最终事件终止。基准没有 Action == "run"、"pause" 或 "cont" 的事件。



=== "main.go"

    ``` go 
    // Copyright 2017 The Go Authors. All rights reserved.
    // Use of this source code is governed by a BSD-style
    // license that can be found in the LICENSE file.
    
    // Test2json converts go test output to a machine-readable JSON stream.
    //
    // Usage:
    //
    //	go tool test2json [-p pkg] [-t] [./pkg.test -test.v [-test.paniconexit0]]
    //
    // Test2json runs the given test command and converts its output to JSON;
    // with no command specified, test2json expects test output on standard input.
    // It writes a corresponding stream of JSON events to standard output.
    // There is no unnecessary input or output buffering, so that
    // the JSON stream can be read for "live updates" of test status.
    //
    // The -p flag sets the package reported in each test event.
    //
    // The -t flag requests that time stamps be added to each test event.
    //
    // The test must be invoked with -test.v. Additionally passing
    // -test.paniconexit0 will cause test2json to exit with a non-zero
    // status if one of the tests being run calls os.Exit(0).
    //
    // Note that test2json is only intended for converting a single test
    // binary's output. To convert the output of a "go test" command,
    // use "go test -json" instead of invoking test2json directly.
    //
    // # Output Format
    //
    // The JSON stream is a newline-separated sequence of TestEvent objects
    // corresponding to the Go struct:
    //
    //	type TestEvent struct {
    //		Time    time.Time // encodes as an RFC3339-format string
    //		Action  string
    //		Package string
    //		Test    string
    //		Elapsed float64 // seconds
    //		Output  string
    //	}
    //
    // The Time field holds the time the event happened.
    // It is conventionally omitted for cached test results.
    //
    // The Action field is one of a fixed set of action descriptions:
    //
    //	run    - the test has started running
    //	pause  - the test has been paused
    //	cont   - the test has continued running
    //	pass   - the test passed
    //	bench  - the benchmark printed log output but did not fail
    //	fail   - the test or benchmark failed
    //	output - the test printed output
    //	skip   - the test was skipped or the package contained no tests
    //
    // The Package field, if present, specifies the package being tested.
    // When the go command runs parallel tests in -json mode, events from
    // different tests are interlaced; the Package field allows readers to
    // separate them.
    //
    // The Test field, if present, specifies the test, example, or benchmark
    // function that caused the event. Events for the overall package test
    // do not set Test.
    //
    // The Elapsed field is set for "pass" and "fail" events. It gives the time
    // elapsed for the specific test or the overall package test that passed or failed.
    //
    // The Output field is set for Action == "output" and is a portion of the test's output
    // (standard output and standard error merged together). The output is
    // unmodified except that invalid UTF-8 output from a test is coerced
    // into valid UTF-8 by use of replacement characters. With that one exception,
    // the concatenation of the Output fields of all output events is the exact
    // output of the test execution.
    //
    // When a benchmark runs, it typically produces a single line of output
    // giving timing results. That line is reported in an event with Action == "output"
    // and no Test field. If a benchmark logs output or reports a failure
    // (for example, by using b.Log or b.Error), that extra output is reported
    // as a sequence of events with Test set to the benchmark name, terminated
    // by a final event with Action == "bench" or "fail".
    // Benchmarks have no events with Action == "run", "pause", or "cont".
    package main
    
    import (
    	"flag"
    	"fmt"
    	"io"
    	"os"
    	"os/exec"
    
    	"cmd/internal/test2json"
    )
    
    var (
    	flagP = flag.String("p", "", "report `pkg` as the package being tested in each event")
    	flagT = flag.Bool("t", false, "include timestamps in events")
    )
    
    func usage() {
    	fmt.Fprintf(os.Stderr, "usage: go tool test2json [-p pkg] [-t] [./pkg.test -test.v]\n")
    	os.Exit(2)
    }
    
    func main() {
    	flag.Usage = usage
    	flag.Parse()
    
    	var mode test2json.Mode
    	if *flagT {
    		mode |= test2json.Timestamp
    	}
    	c := test2json.NewConverter(os.Stdout, *flagP, mode)
    	defer c.Close()
    
    	if flag.NArg() == 0 {
    		io.Copy(c, os.Stdin)
    	} else {
    		args := flag.Args()
    		cmd := exec.Command(args[0], args[1:]...)
    		w := &countWriter{0, c}
    		cmd.Stdout = w
    		cmd.Stderr = w
    		err := cmd.Run()
    		if err != nil {
    			if w.n > 0 {
    				// Assume command printed why it failed.
    			} else {
    				fmt.Fprintf(c, "test2json: %v\n", err)
    			}
    		}
    		c.Exited(err)
    		if err != nil {
    			c.Close()
    			os.Exit(1)
    		}
    	}
    }
    
    type countWriter struct {
    	n int64
    	w io.Writer
    }
    
    func (w *countWriter) Write(b []byte) (int, error) {
    	w.n += int64(len(b))
    	return w.w.Write(b)
    }
    ```

