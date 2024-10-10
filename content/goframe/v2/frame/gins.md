+++
title = "gins"
date = 2024-03-21T17:51:35+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/frame/gins](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/frame/gins)

Package gins provides instances and core components management.

​	Package gins 提供实例和核心组件管理。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Config

```go
func Config(name ...string) *gcfg.Config
```

Config returns an instance of View with default settings. The parameter `name` is the name for the instance.

​	Config 返回具有默认设置的 View 实例。该参数 `name` 是实例的名称。

#### func Database

```go
func Database(name ...string) gdb.DB
```

Database returns an instance of database ORM object with specified configuration group name. Note that it panics if any error occurs duration instance creating.

​	Database 返回具有指定配置组名称的数据库 ORM 对象的实例。请注意，如果实例创建过程中发生任何错误，它会崩溃。

#### func HttpClient

```go
func HttpClient(name ...interface{}) *gclient.Client
```

HttpClient returns an instance of http client with specified name.

​	HttpClient 返回具有指定名称的 http 客户端实例。

#### func I18n

```go
func I18n(name ...string) *gi18n.Manager
```

I18n returns an instance of gi18n.Manager. The parameter `name` is the name for the instance.

​	I18n 返回 gi18n 的实例。经理。该参数 `name` 是实例的名称。

#### func Log

```go
func Log(name ...string) *glog.Logger
```

Log returns an instance of glog.Logger. The parameter `name` is the name for the instance. Note that it panics if any error occurs duration instance creating.

​	log 返回 glog 的实例。记录。该参数 `name` 是实例的名称。请注意，如果实例创建过程中发生任何错误，它会崩溃。

#### func Redis

```go
func Redis(name ...string) *gredis.Redis
```

Redis returns an instance of redis client with specified configuration group name. Note that it panics if any error occurs duration instance creating.

​	Redis 返回具有指定配置组名称的 redis 客户端实例。请注意，如果实例创建过程中发生任何错误，它会崩溃。

#### func Resource

```go
func Resource(name ...string) *gres.Resource
```

Resource returns an instance of Resource. The parameter `name` is the name for the instance.

​	Resource 返回 Resource 的实例。该参数 `name` 是实例的名称。

#### func Server

```go
func Server(name ...interface{}) *ghttp.Server
```

Server returns an instance of http server with specified name. Note that it panics if any error occurs duration instance creating.

​	Server 返回具有指定名称的 http 服务器实例。请注意，如果实例创建过程中发生任何错误，它会崩溃。

#### func View

```go
func View(name ...string) *gview.View
```

View returns an instance of View with default settings. The parameter `name` is the name for the instance. Note that it panics if any error occurs duration instance creating.

​	View 返回具有默认设置的 View 实例。该参数 `name` 是实例的名称。请注意，如果实例创建过程中发生任何错误，它会崩溃。

## 类型

This section is empty.