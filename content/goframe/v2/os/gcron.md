+++
title = "gcron"
date = 2024-03-21T17:55:06+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcron](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcron)

Package gcron implements a cron pattern parser and job runner.

​	软件包 gcron 实现了 cron 模式解析器和作业运行程序。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gcron/gcron.go#L18)

```go
const (
	StatusReady   = gtimer.StatusReady
	StatusRunning = gtimer.StatusRunning
	StatusStopped = gtimer.StatusStopped
	StatusClosed  = gtimer.StatusClosed
)
```

## 变量

This section is empty.

## 函数

#### func DelayAdd

```go
func DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAdd adds a timed task to default cron object after `delay` time.

​	DelayAdd 在一段时间后 `delay` 将定时任务添加到默认的 cron 对象。

#### func DelayAddOnce

```go
func DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddOnce adds a timed task after `delay` time to default cron object. This timed task can be run only once.

​	DelayAddOnce 在 time 之后 `delay` 将定时任务添加到默认 cron 对象。此定时任务只能运行一次。

#### func DelayAddSingleton

```go
func DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddSingleton adds a singleton timed task after `delay` time to default cron object.

​	DelayAddSingleton 在 time 之后 `delay` 将单例定时任务添加到默认 cron 对象。

#### func DelayAddTimes

```go
func DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string)
```

DelayAddTimes adds a timed task after `delay` time to default cron object. This timed task can be run specified times.

​	DelayAddTimes 在 time 之后 `delay` 向默认 cron 对象添加一个定时任务。此定时任务可以运行指定时间。

#### func GetLogger

```go
func GetLogger() glog.ILogger
```

GetLogger returns the global logger in the cron.

​	GetLogger 返回 cron 中的全局记录器。

#### func Remove

```go
func Remove(name string)
```

Remove deletes scheduled task which named `name`.

​	删除删除名为 `name` .

#### func SetLogger

```go
func SetLogger(logger glog.ILogger)
```

SetLogger sets the global logger for cron.

​	SetLogger 为 cron 设置全局记录器。

#### func Size

```go
func Size() int
```

Size returns the size of the timed tasks of default cron.

​	Size 返回默认 cron 的定时任务的大小。

#### func Start

```go
func Start(name ...string)
```

Start starts running the specified timed task named `name`. If no`name` specified, it starts the entire cron.

​	Start 开始运行名为 `name` 的指定定时任务。如果未 `name` 指定，则启动整个 cron。

#### func Stop

```go
func Stop(name ...string)
```

Stop stops running the specified timed task named `name`. If no`name` specified, it stops the entire cron.

​	停止运行名为 `name` 的指定定时任务。如果未 `name` 指定，则停止整个 cron。

## 类型

### type Cron

```go
type Cron struct {
	// contains filtered or unexported fields
}
```

Cron stores all the cron job entries.

​	Cron 存储所有 cron 作业条目。

#### func New

```go
func New() *Cron
```

New returns a new Cron object with default settings.

​	New 返回具有默认设置的新 Cron 对象。

#### (*Cron) Add

```go
func (c *Cron) Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

Add adds a timed task. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

​	添加添加定时任务。唯一 `name` 可以与定时任务绑定。如果已使用， `name` 则返回 and 错误。

#### (*Cron) AddEntry

```go
func (c *Cron) AddEntry(
	ctx context.Context, pattern string, job JobFunc, times int, isSingleton bool, name ...string,
) (*Entry, error)
```

AddEntry creates and returns a new Entry object.

​	AddEntry 创建并返回一个新的 Entry 对象。

#### (*Cron) AddOnce

```go
func (c *Cron) AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddOnce adds a timed task which can be run only once. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

​	AddOnce 添加了一个只能运行一次的定时任务。唯一 `name` 可以与定时任务绑定。如果已使用， `name` 则返回 and 错误。

#### (*Cron) AddSingleton

```go
func (c *Cron) AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddSingleton adds a singleton timed task. A singleton timed task is that can only be running one single instance at the same time. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

​	AddSingleton 添加一个单例定时任务。单例定时任务是指只能同时运行一个实例。唯一 `name` 可以与定时任务绑定。如果已使用， `name` 则返回 and 错误。

#### (*Cron) AddTimes

```go
func (c *Cron) AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error)
```

AddTimes adds a timed task which can be run specified times. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

​	AddTimes 添加了一个定时任务，该任务可以运行指定的时间。唯一 `name` 可以与定时任务绑定。如果已使用， `name` 则返回 and 错误。

#### (*Cron) Close

```go
func (c *Cron) Close()
```

Close stops and closes current cron.

​	关闭停止并关闭当前 cron。

#### (*Cron) DelayAdd

```go
func (c *Cron) DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAdd adds a timed task after `delay` time.

​	DelayAdd 一次又一次地 `delay` 添加一个定时任务。

#### (*Cron) DelayAddEntry

```go
func (c *Cron) DelayAddEntry(ctx context.Context, delay time.Duration, pattern string, job JobFunc, times int, isSingleton bool, name ...string)
```

DelayAddEntry adds a timed task after `delay` time.

​	DelayAddEntry 在一段时间后 `delay` 添加一个定时任务。

#### (*Cron) DelayAddOnce

```go
func (c *Cron) DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddOnce adds a timed task after `delay` time. This timed task can be run only once.

​	DelayAddOnce 在一段时间后 `delay` 添加一个定时任务。此定时任务只能运行一次。

#### (*Cron) DelayAddSingleton

```go
func (c *Cron) DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddSingleton adds a singleton timed task after `delay` time.

​	DelayAddSingleton 在一段时间后 `delay` 添加一个单一的计时任务。

#### (*Cron) DelayAddTimes

```go
func (c *Cron) DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string)
```

DelayAddTimes adds a timed task after `delay` time. This timed task can be run specified times.

​	DelayAddTimes 会逐 `delay` 个添加定时任务。此定时任务可以运行指定时间。

#### (*Cron) Entries

```go
func (c *Cron) Entries() []*Entry
```

Entries return all timed tasks as slice(order by registered time asc).

​	条目以切片形式返回所有计时任务（按注册时间排序）。

#### (*Cron) GetLogger

```go
func (c *Cron) GetLogger() glog.ILogger
```

GetLogger returns the logger in the cron.

​	GetLogger 返回 cron 中的记录器。

#### (*Cron) Remove

```go
func (c *Cron) Remove(name string)
```

Remove deletes scheduled task which named `name`.

​	删除删除名为 `name` .

#### (*Cron) Search

```go
func (c *Cron) Search(name string) *Entry
```

Search returns a scheduled task with the specified `name`. It returns nil if not found.

​	搜索返回具有指定 `name` .如果未找到，则返回 nil。

#### (*Cron) SetLogger

```go
func (c *Cron) SetLogger(logger glog.ILogger)
```

SetLogger sets the logger for cron.

​	SetLogger 为 cron 设置记录器。

#### (*Cron) Size

```go
func (c *Cron) Size() int
```

Size returns the size of the timed tasks.

​	Size 返回定时任务的大小。

#### (*Cron) Start

```go
func (c *Cron) Start(name ...string)
```

Start starts running the specified timed task named `name`. If no`name` specified, it starts the entire cron.

​	Start 开始运行名为 `name` 的指定定时任务。如果未 `name` 指定，则启动整个 cron。

#### (*Cron) Stop

```go
func (c *Cron) Stop(name ...string)
```

Stop stops running the specified timed task named `name`. If no`name` specified, it stops the entire cron.

​	停止运行名为 `name` 的指定定时任务。如果未 `name` 指定，则停止整个 cron。

### type Entry

```go
type Entry struct {
	Name string    // Entry name.
	Job  JobFunc   `json:"-"` // Callback function.
	Time time.Time // Registered time.
	// contains filtered or unexported fields
}
```

Entry is timing task entry.

​	条目是计时任务条目。

#### func Add

```go
func Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

Add adds a timed task to default cron object. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

​	Add 将定时任务添加到默认 cron 对象。唯一 `name` 可以与定时任务绑定。如果已使用， `name` 则返回 and 错误。

#### func AddOnce

```go
func AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddOnce adds a timed task which can be run only once, to default cron object. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

​	AddOnce 将一个只能运行一次的定时任务添加到默认的 cron 对象中。唯一 `name` 可以与定时任务绑定。如果已使用， `name` 则返回 and 错误。

#### func AddSingleton

```go
func AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddSingleton adds a singleton timed task, to default cron object. A singleton timed task is that can only be running one single instance at the same time. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

​	AddSingleton 将单例定时任务添加到默认 cron 对象。单例定时任务是指只能同时运行一个实例。唯一 `name` 可以与定时任务绑定。如果已使用， `name` 则返回 and 错误。

#### func AddTimes

```go
func AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error)
```

AddTimes adds a timed task which can be run specified times, to default cron object. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

​	AddTimes 将一个可以运行指定时间的定时任务添加到默认的 cron 对象中。唯一 `name` 可以与定时任务绑定。如果已使用， `name` 则返回 and 错误。

#### func Entries

```go
func Entries() []*Entry
```

Entries return all timed tasks as slice.

​	条目以切片形式返回所有计时任务。

#### func Search

```go
func Search(name string) *Entry
```

Search returns a scheduled task with the specified `name`. It returns nil if no found.

​	搜索返回具有指定 `name` .如果未找到，则返回 nil。

#### (*Entry) Close

```go
func (e *Entry) Close()
```

Close stops and removes the entry from cron.

​	关闭停止并从 cron 中删除条目。

#### (*Entry) IsSingleton

```go
func (e *Entry) IsSingleton() bool
```

IsSingleton return whether this entry is a singleton timed task.

​	IsSingleton 返回此条目是否为单例定时任务。

#### (*Entry) SetSingleton

```go
func (e *Entry) SetSingleton(enabled bool)
```

SetSingleton sets the entry running in singleton mode.

​	SetSingleton 设置在单例模式下运行的条目。

#### (*Entry) SetStatus

```go
func (e *Entry) SetStatus(status int) int
```

SetStatus sets the status of the entry.

​	SetStatus 设置条目的状态。

#### (*Entry) SetTimes

```go
func (e *Entry) SetTimes(times int)
```

SetTimes sets the times which the entry can run.

​	SetTimes 设置条目可以运行的时间。

#### (*Entry) Start

```go
func (e *Entry) Start()
```

Start starts running the entry.

​	Start 开始运行该条目。

#### (*Entry) Status

```go
func (e *Entry) Status() int
```

Status returns the status of entry.

​	Status 返回进入状态。

#### (*Entry) Stop

```go
func (e *Entry) Stop()
```

Stop stops running the entry.

​	Stop 停止运行条目。

### type JobFunc

```go
type JobFunc = gtimer.JobFunc
```

JobFunc is the timing called job function in cron.

​	JobFunc 是在 cron 中称为作业函数的计时。