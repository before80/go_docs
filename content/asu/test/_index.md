+++
title = "代码测试"
date = 2023-06-12T17:20:58+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 目的

​	确保代码的功能正常，代码的性能满足需求。

## 测试方法

​	白盒测试（又称包内测试）、黑盒测试（又称包外测试）。	

## 测试类型

​	单元测试、性能测试、模糊测试、示例测试。

## 测试标志

​	另请参阅：[go命令中的go test部分]({{< ref "/cmd/go#go-test---测试包">}})  和 [go help testflag]({{< ref "/cmd/gohelptestflag">}})。

## 测试规范

### 测试文件名的命名规范

​	go的测试文件的名称必须以`_test.go`结尾。

### 测试包名的命名规范

​	对于白盒测试，测试包的包名和被测试包的包名一致；

​	对于黑盒测试，通常是测试包的包名加上`_test`；

### 测试函数名的命名规范

​	对于单元测试，测试函数名**必须**以`Test`开头，例如，`TestXxx`或`Test_xxx`；

​	对于性能测试，测试函数名**必须**以`Benchmark`开头，例如，`BenchmarkXxx`或`Benchmark_xxx`；

​	对于示例测试，测试函数名**必须**以`Example`开头，例如，`ExampleXxx`或`Example_xxx`；

​	若想用多个测试用例来对某一被函数或方法进行单元测试，可以将**测试场景**添加到函数名末尾，例如，`Strings.Compare`函数有这些测试函数：`TestCompare`、`TestCompareIdenticalString`、`TestCompareStrings`。

### 测试变量名的命名规范

​	go语言和go test没有对变量的命名做任何约束，但有一些规范值得遵守。

​	例如，为了清晰表达函数的实际输出和预期输出，可以将这两类输出命名为`expected/actual`或者`got/want`。

​	针对其他变量命名，可以遵循go语言推荐的变量命名方法。例如：

- go中的变量名应该短而不是长，这对于范围有限的局部变量来说尤其如此。
- 变量离声明越远，对名称的描述性要求越高。
- 像循环、索引之类的变量，变量名可以是单个字母（例如，i）。若是不常见的变量和全局变量，变量名就需要具有更多的描述性。