+++
title = "coverage"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/runtime/coverage@go1.24.2](https://pkg.go.dev/runtime/coverage@go1.24.2)



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func ClearCounters 

``` go 
func ClearCounters() error
```

ClearCounters clears/resets all coverage counter variables in the currently running program. It returns an error if the program in question was not built with the “-cover” flag. Clearing of coverage counters is also not supported for programs not using atomic counter mode (see more detailed comments below for the rationale here).

​	ClearCounters 清除/重置当前正在运行程序中的所有覆盖率计数器变量。如果所讨论的程序不是使用“-cover”标志构建的，它将返回一个错误。对于不使用原子计数器模式的程序，也不支持清除覆盖率计数器（有关此处原理的更详细注释，请参见下文）。

### func WriteCounters

```go
func WriteCounters(w io.Writer) error
```

WriteCounters writes coverage counter-data content for the currently running program to the writer ‘w’. An error will be returned if the operation can’t be completed successfully (for example, if the currently running program was not built with “-cover”, or if a write fails). The counter data written will be a snapshot taken at the point of the invocation.

​	WriteCounters 将当前正在运行程序的覆盖率计数器数据内容写入编写器“w”。如果无法成功完成操作（例如，如果当前正在运行的程序不是使用“-cover”构建的，或者写入失败），将返回一个错误。写入的计数器数据将是调用时拍摄的快照。

### func WriteCountersDir

```go
func WriteCountersDir(dir string) error
```

WriteCountersDir writes a coverage counter-data file for the currently running program to the directory specified in ‘dir’. An error will be returned if the operation can’t be completed successfully (for example, if the currently running program was not built with “-cover”, or if the directory does not exist). The counter data written will be a snapshot taken at the point of the call.

​	WriteCountersDir 将当前正在运行程序的覆盖率计数器数据文件写入“dir”中指定的目录。如果无法成功完成操作（例如，如果当前正在运行的程序不是使用“-cover”构建的，或者目录不存在），将返回一个错误。写入的计数器数据将是调用时拍摄的快照。

### func WriteMeta

```go
func WriteMeta(w io.Writer) error
```

WriteMeta writes the meta-data content (the payload that would normally be emitted to a meta-data file) for the currently running program to the the writer ‘w’. An error will be returned if the operation can’t be completed successfully (for example, if the currently running program was not built with “-cover”, or if a write fails).

​	WriteMeta 将当前运行程序的元数据内容（通常会输出到元数据文件的有效负载）写入编写器“w”。如果操作无法成功完成（例如，如果当前运行的程序不是使用“-cover”构建的，或者写入失败），则会返回错误。

### func WriteMetaDir

```go
func WriteMetaDir(dir string) error
```

WriteMetaDir writes a coverage meta-data file for the currently running program to the directory specified in ‘dir’. An error will be returned if the operation can’t be completed successfully (for example, if the currently running program was not built with “-cover”, or if the directory does not exist).

​	WriteMetaDir 将当前运行程序的覆盖率元数据文件写入“dir”中指定的目录。如果操作无法成功完成（例如，如果当前运行的程序不是使用“-cover”构建的，或者目录不存在），则会返回错误。

## 类型

This section is empty.