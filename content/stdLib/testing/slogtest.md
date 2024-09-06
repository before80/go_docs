+++
title = "slogtest"
date = 2023-11-05T14:33:11+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/testing/slogtest@go1.23.0](https://pkg.go.dev/testing/slogtest@go1.23.0)

> 注意
>
> ​	从go1.21.0开始才可以使用该包。

## 概述

Package slogtest implements support for testing implementations of log/slog.Handler.

## Example (Parsing)

This example demonstrates one technique for testing a handler with this package. The handler is given a [bytes.Buffer](https://pkg.go.dev/bytes#Buffer) to write to, and each line of the resulting output is parsed. For JSON output, [encoding/json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) produces a result in the desired format when given a pointer to a map[string]any.

``` go
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"log/slog"
	"testing/slogtest"
)

func main() {
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)

	results := func() []map[string]any {
		var ms []map[string]any
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				panic(err) // In a real test, use t.Fatal.
			}
			ms = append(ms, m)
		}
		return ms
	}
	err := slogtest.TestHandler(h, results)
	if err != nil {
		log.Fatal(err)
	}

}

```
## 常量

This section is empty.

## 变量

This section is empty.

## 函数 

### func Run <- go1.22.0

``` go
func Run(t *testing.T, newHandler func(*testing.T) slog.Handler, result func(*testing.T) map[string]any)
```

Run exercises a [slog.Handler](https://pkg.go.dev/log/slog#Handler) on the same test cases as [TestHandler](https://pkg.go.dev/testing/slogtest#TestHandler), but runs each case in a subtest. For each test case, it first calls newHandler to get an instance of the handler under test, then runs the test case, then calls result to get the result. If the test case fails, it calls t.Error.

​	`Run` 在与 [TestHandler](https://pkg.go.dev/testing/slogtest#TestHandler) 相同的测试用例上测试一个 [slog.Handler](https://pkg.go.dev/log/slog#Handler)，但每个测试用例都在子测试中运行。对于每个测试用例，它首先调用 `newHandler` 来获取待测试的处理器实例，然后运行测试用例，最后调用 `result` 来获取结果。如果测试用例失败，它会调用 `t.Error`。

### func TestHandler 

``` go
func TestHandler(h slog.Handler, results func() []map[string]any) error
```

TestHandler tests a [slog.Handler](https://pkg.go.dev/log/slog#Handler). If TestHandler finds any misbehaviors, it returns an error for each, combined into a single error with errors.Join.

​	`TestHandler` 测试一个 [slog.Handler](https://pkg.go.dev/log/slog#Handler)。如果 `TestHandler` 发现任何异常行为，它会将每个错误返回，并将它们合并成一个单一错误，通过 `errors.Join` 组合。

TestHandler installs the given Handler in a [slog.Logger](https://pkg.go.dev/log/slog#Logger) and makes several calls to the Logger's output methods.

​	`TestHandler` 将给定的 `Handler` 安装到一个 [slog.Logger](https://pkg.go.dev/log/slog#Logger) 中，并多次调用 Logger 的输出方法。

The results function is invoked after all such calls. It should return a slice of map[string]any, one for each call to a Logger output method. The keys and values of the map should correspond to the keys and values of the Handler's output. Each group in the output should be represented as its own nested map[string]any. The standard keys slog.TimeKey, slog.LevelKey and slog.MessageKey should be used.

​	`results` 函数在所有调用完成后执行。它应返回一个 `map[string]any` 的切片，每个元素对应 Logger 输出方法的每一次调用。`map` 的键和值应对应 `Handler` 的输出键和值。输出中的每个组都应表示为其自己的嵌套 `map[string]any`。标准键 `slog.TimeKey`、`slog.LevelKey` 和 `slog.MessageKey` 应该使用。

If the Handler outputs JSON, then calling [encoding/json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) with a `map[string]any` will create the right data structure.

​	如果 `Handler` 输出 JSON，那么调用 [encoding/json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) 和 `map[string]any` 将创建正确的数据结构。

If a Handler intentionally drops an attribute that is checked by a test, then the results function should check for its absence and add it to the map it returns.

​	如果一个 `Handler` 有意地丢弃了测试检查的某个属性，那么 `results` 函数应检查它是否缺失，并将其添加到返回的 `map` 中。

## 类型

This section is empty.