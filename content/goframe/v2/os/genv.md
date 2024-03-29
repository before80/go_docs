+++
title = "genv"
date = 2024-03-21T17:55:23+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/genv](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/genv)

Package genv provides operations for environment variables of system.

​	软件包 genv 提供系统环境变量的操作。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func All

```go
func All() []string
```

All returns a copy of strings representing the environment, in the form “key=value”.

​	All 返回表示环境的字符串副本，格式为“key=value”。

#### func Build

```go
func Build(m map[string]string) []string
```

Build builds a map to an environment variable slice.

​	Build 生成到环境变量切片的映射。

#### func Contains

```go
func Contains(key string) bool
```

Contains checks whether the environment variable named `key` exists.

​	包含检查名为 `key` 的环境变量是否存在。

#### func Filter <-2.1.0

```go
func Filter(envs []string) []string
```

Filter filters repeated items from given environment variables.

​	筛选器从给定的环境变量中筛选重复的项目。

#### func Get

```go
func Get(key string, def ...interface{}) *gvar.Var
```

Get creates and returns a Var with the value of the environment variable named by the `key`. It uses the given `def` if the variable does not exist in the environment.

​	Get 创建并返回一个 Var，其值为 `key` .如果环境中不存在变量，则使用给定 `def` 变量。

#### func GetWithCmd

```go
func GetWithCmd(key string, def ...interface{}) *gvar.Var
```

GetWithCmd returns the environment value specified `key`. If the environment value does not exist, then it retrieves and returns the value from command line options. It returns the default value `def` if none of them exists.

​	GetWithCmd 返回指定的 `key` 环境值。如果环境值不存在，则它会从命令行选项中检索并返回该值。如果它们都不存在，则返回默认值 `def` 。

Fetching Rules: 1. Environment arguments are in uppercase format, eg: GF__； 2. Command line arguments are in lowercase format, eg: gf..;

​	获取规则： 1.环境参数为大写格式，例如：GF__;2.命令行参数为小写格式，例如：gf..;

#### func Map

```go
func Map() map[string]string
```

Map returns a copy of strings representing the environment as a map.

​	Map 返回将环境表示为映射的字符串副本。

#### func MapFromEnv <-2.1.0

```go
func MapFromEnv(envs []string) map[string]string
```

MapFromEnv converts environment variables from slice to map.

​	MapFromEnv 将环境变量从切片转换为映射。

#### func MapToEnv <-2.1.0

```go
func MapToEnv(m map[string]string) []string
```

MapToEnv converts environment variables from map to slice.

​	MapToEnv 将环境变量从地图转换为切片。

#### func MustRemove

```go
func MustRemove(key ...string)
```

MustRemove performs as Remove, but it panics if any error occurs.

​	MustRemove 以 Remove 的形式执行，但如果发生任何错误，它会崩溃。

#### func MustSet

```go
func MustSet(key, value string)
```

MustSet performs as Set, but it panics if any error occurs.

​	MustSet 以 Set 的形式执行，但如果发生任何错误，它会崩溃。

#### func Remove

```go
func Remove(key ...string) (err error)
```

Remove deletes one or more environment variables.

​	“删除”（Remove） 将删除一个或多个环境变量。

#### func Set

```go
func Set(key, value string) (err error)
```

Set sets the value of the environment variable named by the `key`. It returns an error, if any.

​	Set 设置由 `key` .它返回错误（如果有）。

#### func SetMap

```go
func SetMap(m map[string]string) (err error)
```

SetMap sets the environment variables using map.

​	SetMap 使用 map 设置环境变量。

## 类型

This section is empty.