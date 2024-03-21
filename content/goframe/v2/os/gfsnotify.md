+++
title = "gfsnotify"
date = 2024-03-21T17:55:52+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfsnotify

Package gfsnotify provides a platform-independent interface for file system notifications.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Exit 

``` go
func Exit()
```

Exit is only used in the callback function, which can be used to remove current callback of itself from the watcher.

##### func Remove 

``` go
func Remove(path string) error
```

Remove removes all monitoring callbacks of given `path` from watcher recursively.

##### func RemoveCallback 

``` go
func RemoveCallback(callbackId int) error
```

RemoveCallback removes specified callback with given id from watcher.

### Types 

#### type Callback 

``` go
type Callback struct {
	Id   int                // Unique id for callback object.
	Func func(event *Event) // Callback function.
	Path string             // Bound file path (absolute).
	// contains filtered or unexported fields
}
```

Callback is the callback function for Watcher.

##### func Add 

``` go
func Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error)
```

Add monitors `path` using default watcher with callback function `callbackFunc`. The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

##### func AddOnce 

``` go
func AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error)
```

AddOnce monitors `path` using default watcher with callback function `callbackFunc` only once using unique name `name`. If AddOnce is called multiple times with the same `name` parameter, `path` is only added to monitor once. It returns error if it's called twice with the same `name`.

The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

#### type Event 

``` go
type Event struct {
	Path    string   // Absolute file path.
	Op      Op       // File operation.
	Watcher *Watcher // Parent watcher.
	// contains filtered or unexported fields
}
```

Event is the event produced by underlying fsnotify.

##### (*Event) IsChmod 

``` go
func (e *Event) IsChmod() bool
```

IsChmod checks whether current event contains file/folder chmod event.

##### (*Event) IsCreate 

``` go
func (e *Event) IsCreate() bool
```

IsCreate checks whether current event contains file/folder create event.

##### (*Event) IsRemove 

``` go
func (e *Event) IsRemove() bool
```

IsRemove checks whether current event contains file/folder remove event.

##### (*Event) IsRename 

``` go
func (e *Event) IsRename() bool
```

IsRename checks whether current event contains file/folder rename event.

##### (*Event) IsWrite 

``` go
func (e *Event) IsWrite() bool
```

IsWrite checks whether current event contains file/folder write event.

##### (*Event) String 

``` go
func (e *Event) String() string
```

String returns current event as string.

#### type Op 

``` go
type Op uint32
```

Op is the bits union for file operations.

``` go
const (
	CREATE Op = 1 << iota
	WRITE
	REMOVE
	RENAME
	CHMOD
)
```

#### type Watcher 

``` go
type Watcher struct {
	// contains filtered or unexported fields
}
```

Watcher is the monitor for file changes.

##### func New 

``` go
func New() (*Watcher, error)
```

New creates and returns a new watcher. Note that the watcher number is limited by the file handle setting of the system. Eg: fs.inotify.max_user_instances system variable in linux systems.

##### (*Watcher) Add 

``` go
func (w *Watcher) Add(path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error)
```

Add monitors `path` with callback function `callbackFunc` to the watcher. The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

##### (*Watcher) AddOnce 

``` go
func (w *Watcher) AddOnce(name, path string, callbackFunc func(event *Event), recursive ...bool) (callback *Callback, err error)
```

AddOnce monitors `path` with callback function `callbackFunc` only once using unique name `name` to the watcher. If AddOnce is called multiple times with the same `name` parameter, `path` is only added to monitor once.

It returns error if it's called twice with the same `name`.

The optional parameter `recursive` specifies whether monitoring the `path` recursively, which is true in default.

##### (*Watcher) Close 

``` go
func (w *Watcher) Close()
```

Close closes the watcher.

##### (*Watcher) Remove 

``` go
func (w *Watcher) Remove(path string) error
```

Remove removes monitor and all callbacks associated with the `path` recursively.

##### (*Watcher) RemoveCallback 

``` go
func (w *Watcher) RemoveCallback(callbackId int)
```

RemoveCallback removes callback with given callback id from watcher.