+++
title = "server_messaging"
date = 2024-12-15T21:22:14+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/messaging](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/messaging)
>
> 收录该文档时间： `2024-12-15T21:22:14+08:00`

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Folder 

``` go
type Folder struct {
	Path          string // key
	Root          string
	Ignored       bool
	Disabled      bool
	BuildTags     []string
	TestArguments []string
}
```

### type Folders 

``` go
type Folders map[string]*Folder
```

### type WatcherCommand 

``` go
type WatcherCommand struct {
	Instruction WatcherInstruction
	Details     string
}
```

### type WatcherInstruction 

``` go
type WatcherInstruction int
const (
	WatcherPause WatcherInstruction = iota
	WatcherResume
	WatcherIgnore
	WatcherReinstate
	WatcherAdjustRoot
	WatcherExecute
	WatcherStop
)
```

#### (WatcherInstruction) String 

``` go
func (this WatcherInstruction) String() string
```
