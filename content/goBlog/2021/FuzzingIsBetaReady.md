+++
title = "模糊测试已经准备好进入Beta测试"
weight = 92
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Fuzzing is Beta Ready - 模糊测试已经准备好进入Beta测试

[https://go.dev/blog/fuzz-beta](https://go.dev/blog/fuzz-beta)

Katie Hockman and Jay Conrod
3 June 2021

2021年6月3日

​	我们很高兴地宣布，原生模糊测试已经准备好在tip上进行beta测试了!

​	Fuzzing是一种自动化测试，它持续操纵程序的输入，以发现panic或bug等问题。这些半随机的数据突变可以发现现有单元测试可能错过的新的代码覆盖率，并发现边缘用例的bug，否则这些bug就不会被注意到。由于 fuzzing 可以接触到这些边缘用例，所以模糊测试对于发现安全漏洞和漏洞特别有价值。

​	关于该功能的更多细节，请参见 [golang.org/s/draft-fuzzing-design](https://go.dev/s/draft-fuzzing-design)。

## 开始使用

​	要开始使用，您可以运行以下命令

```shell linenums="1"
$ go install golang.org/dl/gotip@latest
$ gotip download
```

​	这将从主分支构建Go工具链。运行这个命令后，`gotip`可以作为`go`命令的直接替代者。现在您可以运行如下命令

```shell linenums="1"
$ gotip test -fuzz=Fuzz
```

## 编写模糊测试

​	一个模糊测试必须在`*_test.go`文件中以`FuzzXxx`的形式成为一个函数。这个函数必须被传递一个`*testing.F`参数，就像`TestXxx`函数被传递`*testing.T`参数一样。

​	以下是一个测试[net/url](https://pkg.go.dev/net/url#ParseQuery)包行为的模糊测试的示例。

```go linenums="1"
//go:build go1.18
// +build go1.18

package fuzz

import (
    "net/url"
    "reflect"
    "testing"
)

func FuzzParseQuery(f *testing.F) {
    f.Add("x=1&y=2")
    f.Fuzz(func(t *testing.T, queryStr string) {
        query, err := url.ParseQuery(queryStr)
        if err != nil {
            t.Skip()
        }
        queryStr2 := query.Encode()
        query2, err := url.ParseQuery(queryStr2)
        if err != nil {
            t.Fatalf("ParseQuery failed to decode a valid encoded query %s: %v", queryStr2, err)
        }
        if !reflect.DeepEqual(query, query2) {
            t.Errorf("ParseQuery gave different query after being encoded\nbefore: %v\nafter: %v", query, query2)
        }
    })
}
```

​	您可以在pkg.go.dev阅读更多关于fuzzing的内容，包括[Go的fuzzing概述](https://pkg.go.dev/testing@master#hdr-Fuzzing)和[新的testing.F类型的godoc](https://pkg.go.dev/testing@master#F)。

## 期望

​	这是一项仍处于测试阶段的新功能，因此您应该会遇到一些错误和不完整的功能集。请查看[问题追踪器中标有 "fuzz"](https://github.com/golang/go/issues?q=is%3Aopen+is%3Aissue+label%3Afuzz)的问题，以了解现有错误和缺失功能的最新情况。

​	请注意，fuzzing会消耗大量内存，在运行时可能会影响您的机器性能。 `go test -fuzz`默认为在`$GOMAXPROCS`进程中并行运行fuzzing。您可以通过在`go test`中明确设置`-parallel`标志来降低模糊处理时使用的进程数。如果您想了解更多信息，可以运行 `gotip help testflag` 阅读 `go test` 命令的文档。

​	还要注意的是，模糊测试引擎在运行时，会将扩大测试范围的数值写入`$GOCACHE/fuzz`的模糊缓存目录中。目前对写入模糊缓存的文件数量或总字节数没有限制，所以它可能会占用大量的存储空间（即几个GB）。您可以通过运行 `gotip clean -fuzzcache` 来清除 fuzz 缓存。

## 下一步是什么？

​	这项功能将在Go 1.18中开始使用。

​	如果您遇到任何问题或有任何关于该功能的想法，请[提交问题](https://github.com/golang/go/issues/new/?&labels=fuzz)。

​	关于该功能的讨论和一般反馈，您也可以参与Gophers Slack的[#fuzzing频道](https://gophers.slack.com/archives/CH5KV1AKE)。

​	Happy fuzzing!
