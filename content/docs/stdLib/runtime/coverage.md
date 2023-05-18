+++
title = "coverage"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# coverage

https://pkg.go.dev/runtime/coverage@go1.20.1



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [ClearCounters](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/coverage/apis.go;l=89) 

``` go 
func ClearCounters() error
```

ClearCounters clears/resets all coverage counter variables in the currently running program. It returns an error if the program in question was not built with the "-cover" flag. Clearing of coverage counters is also not supported for programs not using atomic counter mode (see more detailed comments below for the rationale here).

#### func [WriteCounters](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/coverage/apis.go;l=62) 

``` go 
func WriteCounters(w io.Writer) error
```

WriteCounters writes coverage counter-data content for the currently running program to the writer 'w'. An error will be returned if the operation can't be completed successfully (for example, if the currently running program was not built with "-cover", or if a write fails). The counter data written will be a snapshot taken at the point of the invocation.

#### func [WriteCountersDir](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/coverage/apis.go;l=52) 

``` go 
func WriteCountersDir(dir string) error
```

WriteCountersDir writes a coverage counter-data file for the currently running program to the directory specified in 'dir'. An error will be returned if the operation can't be completed successfully (for example, if the currently running program was not built with "-cover", or if the directory does not exist). The counter data written will be a snapshot taken at the point of the call.

#### func [WriteMeta](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/coverage/apis.go;l=34) 

``` go 
func WriteMeta(w io.Writer) error
```

WriteMeta writes the meta-data content (the payload that would normally be emitted to a meta-data file) for the currently running program to the the writer 'w'. An error will be returned if the operation can't be completed successfully (for example, if the currently running program was not built with "-cover", or if a write fails).

#### func [WriteMetaDir](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/coverage/apis.go;l=21) 

``` go 
func WriteMetaDir(dir string) error
```

WriteMetaDir writes a coverage meta-data file for the currently running program to the directory specified in 'dir'. An error will be returned if the operation can't be completed successfully (for example, if the currently running program was not built with "-cover", or if the directory does not exist).

## 类型

This section is empty.