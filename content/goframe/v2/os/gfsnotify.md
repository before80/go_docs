+++
title = "gfsnotify"
date = 2024-03-21T17:55:52+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfsnotify](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfsnotify)

Package gfsnotify provides a platform-independent interface for file system notifications.

​	软件包 gfsnotify 为文件系统通知提供了一个独立于平台的接口。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Exit

```go
func Exit()
```

Exit is only used in the callback function, which can be used to remove current callback of itself from the watcher.

​	Exit 仅在回调函数中使用，该函数可用于从观察程序中删除自身的当前回调。

#### func Remove

```go
func Remove(path string) error
```

Remove removes all monitoring callbacks of given `path` from watcher recursively.

​	Remove 以递归方式从观察程序中删除给定 `path` 的所有监视回调。

#### func RemoveCallback

```go
func RemoveCallback(callbackId int) error
```

RemoveCallback removes specified callback with given id from watcher.

​	RemoveCallback 从观察程序中删除具有给定 ID 的指定回调。

## 类型

### type Callback

```go
type Callback struct {
	Id   int                // Unique id for callback object.
	Func func(event *Event) // Callback function.
	Path string             // Bound file path (absolute).
	// contains filtered or unexported fields
}
```

Callback is the callback function for Watcher.

​	Callback 是 Watcher 的回调函数。

#### func Add

```go
func Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error)
```

Add monitors `path` using default watcher with callback function `callbackFunc`. The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

​	 `path` 使用带有回调函数 `callbackFunc` 的默认观察程序添加监视器。可选参数 `recursive` 指定是否以递归方式监视 ， `path` 默认为 true。

#### func AddOnce

```go
func AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error)
```

AddOnce monitors `path` using default watcher with callback function `callbackFunc` only once using unique name `name`. If AddOnce is called multiple times with the same `name` parameter, `path` is only added to monitor once. It returns error if it’s called twice with the same `name`.

​	AddOnce `path` 使用带有回调函数 `callbackFunc` 的默认观察程序仅使用唯一名称 `name` 进行一次监视。如果使用同一 `name` 参数多次调用 AddOnce， `path` 则仅添加一次监视。如果使用相同的 `name` .

The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

​	可选参数 `recursive` 指定是否以递归方式监视 ， `path` 默认为 true。

### type Event

```go
type Event struct {
	Path    string   // Absolute file path.
	Op      Op       // File operation.
	Watcher *Watcher // Parent watcher.
	// contains filtered or unexported fields
}
```

Event is the event produced by underlying fsnotify.

​	Event 是由底层 fsnotify 产生的事件。

#### (*Event) IsChmod

```go
func (e *Event) IsChmod() bool
```

IsChmod checks whether current event contains file/folder chmod event.

​	IsChmod 检查当前事件是否包含文件/文件夹 chmod 事件。

#### (*Event) IsCreate

```go
func (e *Event) IsCreate() bool
```

IsCreate checks whether current event contains file/folder create event.

​	IsCreate 检查当前事件是否包含文件/文件夹创建事件。

#### (*Event) IsRemove

```go
func (e *Event) IsRemove() bool
```

IsRemove checks whether current event contains file/folder remove event.

​	IsRemove 检查当前事件是否包含文件/文件夹删除事件。

#### (*Event) IsRename

```go
func (e *Event) IsRename() bool
```

IsRename checks whether current event contains file/folder rename event.

​	IsRename 检查当前事件是否包含文件/文件夹重命名事件。

#### (*Event) IsWrite

```go
func (e *Event) IsWrite() bool
```

IsWrite checks whether current event contains file/folder write event.

​	IsWrite 检查当前事件是否包含文件/文件夹写入事件。

#### (*Event) String

```go
func (e *Event) String() string
```

String returns current event as string.

​	String 以字符串形式返回当前事件。

### type Op

```go
type Op uint32
```

Op is the bits union for file operations.

​	Op 是文件操作的位联合。

```go
const (
	CREATE Op = 1 << iota
	WRITE
	REMOVE
	RENAME
	CHMOD
)
```

### type Watcher

```go
type Watcher struct {
	// contains filtered or unexported fields
}
```

Watcher is the monitor for file changes.

​	观察程序是文件更改的监视器。

#### func New

```go
func New() (*Watcher, error)
```

New creates and returns a new watcher. Note that the watcher number is limited by the file handle setting of the system. Eg: fs.inotify.max_user_instances system variable in linux systems.

​	new 创建并返回新的观察程序。请注意，观察程序编号受系统文件句柄设置的限制。例如：linux系统中的fs.inotify.max_user_instances系统变量。

#### (*Watcher) Add

```go
func (w *Watcher) Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error)
```

Add monitors `path` with callback function `callbackFunc` to the watcher. The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

​	将具有回调函数 `callbackFunc` 的 `path` 监视器添加到观察程序。可选参数 `recursive` 指定是否以递归方式监视 ， `path` 默认为 true。

#### (*Watcher) AddOnce

```go
func (w *Watcher) AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error)
```

AddOnce monitors `path` with callback function `callbackFunc` only once using unique name `name` to the watcher. If AddOnce is called multiple times with the same `name` parameter, `path` is only added to monitor once.

​	 `path` AddOnce 仅使用观察程序的唯一名称 `name` 监视一次具有回调函数 `callbackFunc` 的监视器。如果使用同一 `name` 参数多次调用 AddOnce， `path` 则仅添加一次监视。

It returns error if it’s called twice with the same `name`.

​	如果使用相同的 `name` .

The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

​	可选参数 `recursive` 指定是否以递归方式监视 ， `path` 默认为 true。

#### (*Watcher) Close

```go
func (w *Watcher) Close()
```

Close closes the watcher.

​	关闭 关闭观察程序。

#### (*Watcher) Remove

```go
func (w *Watcher) Remove(path string) error
```

Remove removes monitor and all callbacks associated with the `path` recursively.

​	Remove 将删除与递归关联的 `path` 监视器和所有回调。

#### (*Watcher) RemoveCallback

```go
func (w *Watcher) RemoveCallback(callbackId int)
```

RemoveCallback removes callback with given callback id from watcher.

​	RemoveCallback 从观察程序中删除具有给定回调 ID 的回调。