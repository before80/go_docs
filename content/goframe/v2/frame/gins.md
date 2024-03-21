+++
title = "gins"
date = 2024-03-21T17:51:35+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/frame/gins

Package gins provides instances and core components management.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Config 

``` go
func Config(name ...string) *gcfg.Config
```

Config returns an instance of View with default settings. The parameter `name` is the name for the instance.

##### func Database 

``` go
func Database(name ...string) gdb.DB
```

Database returns an instance of database ORM object with specified configuration group name. Note that it panics if any error occurs duration instance creating.

##### func HttpClient 

``` go
func HttpClient(name ...interface{}) *gclient.Client
```

HttpClient returns an instance of http client with specified name.

##### func I18n 

``` go
func I18n(name ...string) *gi18n.Manager
```

I18n returns an instance of gi18n.Manager. The parameter `name` is the name for the instance.

##### func Log 

``` go
func Log(name ...string) *glog.Logger
```

Log returns an instance of glog.Logger. The parameter `name` is the name for the instance. Note that it panics if any error occurs duration instance creating.

##### func Redis 

``` go
func Redis(name ...string) *gredis.Redis
```

Redis returns an instance of redis client with specified configuration group name. Note that it panics if any error occurs duration instance creating.

##### func Resource 

``` go
func Resource(name ...string) *gres.Resource
```

Resource returns an instance of Resource. The parameter `name` is the name for the instance.

##### func Server 

``` go
func Server(name ...interface{}) *ghttp.Server
```

Server returns an instance of http server with specified name. Note that it panics if any error occurs duration instance creating.

##### func View 

``` go
func View(name ...string) *gview.View
```

View returns an instance of View with default settings. The parameter `name` is the name for the instance. Note that it panics if any error occurs duration instance creating.

### Types 

This section is empty.