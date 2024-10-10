+++
title = "gtimer"
date = 2024-03-21T17:57:40+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gtimer](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gtimer)

Package gtimer implements timer for interval/delayed jobs running and management.

​	软件包 gtimer 为运行和管理的间隔/延迟作业实现计时器。

This package is designed for management for millions of timing jobs. The differences between gtimer and gcron are as follows:

​	该软件包专为管理数百万个计时作业而设计。gtimer 和 gcron 的区别如下：

1. package gcron is implemented based on package gtimer.
   软件包 GCRON 是基于软件包 Gtimer 实现的。
2. gtimer is designed for high performance and for millions of timing jobs.
   Gtimer 专为高性能和数百万个定时作业而设计。
3. gcron supports configuration pattern grammar like linux crontab, which is more manually readable.
   GCRON支持像Linux Crontab这样的配置模式语法，它更易于手动读取。
4. gtimer’s benchmark OP is measured in nanoseconds, and gcron’s benchmark OP is measured in microseconds.
   gtimer 的基准 OP 以纳秒为单位，gcron 的基准 OP 以微秒为单位。

ALSO VERY NOTE the common delay of the timer: https://github.com/golang/go/issues/14410

​	还要注意计时器的常见延迟：https://github.com/golang/go/issues/14410

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gtimer/gtimer.go#L51)

```go
const (
	StatusReady   = 0  // Job or Timer is ready for running.
	StatusRunning = 1  // Job or Timer is already running.
	StatusStopped = 2  // Job or Timer is stopped.
	StatusClosed  = -1 // Job or Timer is closed and waiting to be deleted.

)
```

## 变量

This section is empty.

## 函数

#### func DelayAdd

```go
func DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAdd adds a timing job after delay of `interval` duration. Also see Add.

​	DelayAdd 在 `interval` 持续时间延迟后添加计时作业。另请参阅添加。

#### func DelayAddEntry

```go
func DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int)
```

DelayAddEntry adds a timing job after delay of `interval` duration. Also see AddEntry.

​	DelayAddEntry 在 `interval` 持续时间延迟后添加计时作业。另请参阅 AddEntry。

#### func DelayAddOnce

```go
func DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddOnce adds a timing job after delay of `interval` duration. Also see AddOnce.

​	DelayAddOnce 在 `interval` 持续时间延迟后添加计时作业。另请参阅 AddOnce。

#### func DelayAddSingleton

```go
func DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddSingleton adds a timing job after delay of `interval` duration. Also see AddSingleton.

​	DelayAddSingleton 在 `interval` 持续时间延迟后添加计时作业。另请参阅 AddSingleton。

#### func DelayAddTimes

```go
func DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc)
```

DelayAddTimes adds a timing job after delay of `interval` duration. Also see AddTimes.

​	DelayAddTimes 在 `interval` 持续时间延迟后添加计时作业。另请参阅 AddTimes。

#### func Exit

```go
func Exit()
```

Exit is used in timing job internally, which exits and marks it closed from timer. The timing job will be automatically removed from timer later. It uses “panic-recover” mechanism internally implementing this feature, which is designed for simplification and convenience.

​	退出用于内部计时作业，该作业退出并从计时器中将其关闭。计时作业稍后将自动从计时器中删除。它在内部使用“紧急恢复”机制来实现此功能，旨在简化和方便。

#### func SetInterval

```go
func SetInterval(ctx context.Context, interval time.Duration, job JobFunc)
```

SetInterval runs the job every duration of `delay`. It is like the one in javascript.

​	SetInterval 每 的 `delay` 持续时间运行一次作业。它就像 javascript 中的一样。

#### func SetTimeout

```go
func SetTimeout(ctx context.Context, delay time.Duration, job JobFunc)
```

SetTimeout runs the job once after duration of `delay`. It is like the one in javascript.

​	SetTimeout 在持续时间为 后 `delay` 运行作业一次。它就像 javascript 中的一样。

## 类型

### type Entry

```go
type Entry struct {
	// contains filtered or unexported fields
}
```

Entry is the timing job.

​	条目是计时作业。

#### func Add

```go
func Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

Add adds a timing job to the default timer, which runs in interval of `interval`.

​	Add 将计时作业添加到默认计时器，该计时器以 的间隔运行 `interval` 。

##### Example

``` go
```

#### func AddEntry

```go
func AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry
```

AddEntry adds a timing job to the default timer with detailed parameters.

​	AddEntry 将计时作业添加到具有详细参数的默认计时器。

The parameter `interval` specifies the running interval of the job.

​	该参数 `interval` 指定作业的运行间隔。

The parameter `singleton` specifies whether the job running in singleton mode. There’s only one of the same job is allowed running when its a singleton mode job.

​	该参数 `singleton` 指定作业是否在单例模式下运行。当作业是单例模式作业时，只允许运行一个相同的作业。

The parameter `times` specifies limit for the job running times, which means the job exits if its run times exceeds the `times`.

​	该参数 `times` 指定作业运行时间的限制，这意味着如果作业的运行时间超过 . `times`

The parameter `status` specifies the job status when it’s firstly added to the timer.

​	该参数 `status` 指定首次添加到计时器时的作业状态。

#### func AddOnce

```go
func AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddOnce is a convenience function for adding a job which only runs once and then exits.

​	AddOnce 是一个方便的函数，用于添加仅运行一次然后退出的作业。

#### func AddSingleton

```go
func AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddSingleton is a convenience function for add singleton mode job.

​	AddSingleton 是添加单例模式作业的便捷函数。

#### func AddTimes

```go
func AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry
```

AddTimes is a convenience function for adding a job which is limited running times.

​	AddTimes 是一个方便的功能，用于添加运行时间有限的作业。

#### (*Entry) Close

```go
func (entry *Entry) Close()
```

Close closes the job, and then it will be removed from the timer.

​	关闭将关闭作业，然后将其从计时器中删除。

#### (*Entry) Ctx

```go
func (entry *Entry) Ctx() context.Context
```

Ctx returns the initialized context of this job.

​	Ctx 返回此作业的初始化上下文。

#### (*Entry) IsSingleton

```go
func (entry *Entry) IsSingleton() bool
```

IsSingleton checks and returns whether the job in singleton mode.

​	IsSingleton 检查并返回作业是否处于单例模式。

#### (*Entry) Job

```go
func (entry *Entry) Job() JobFunc
```

Job returns the job function of this job.

​	Job 返回此作业的作业函数。

#### (*Entry) Reset

```go
func (entry *Entry) Reset()
```

Reset resets the job, which resets its ticks for next running.

​	重置将重置作业，从而重置其下次运行的刻度。

#### (*Entry) Run

```go
func (entry *Entry) Run()
```

Run runs the timer job asynchronously.

​	Run 异步运行计时器作业。

#### (*Entry) SetSingleton

```go
func (entry *Entry) SetSingleton(enabled bool)
```

SetSingleton sets the job singleton mode.

​	SetSingleton 设置作业单例模式。

#### (*Entry) SetStatus

```go
func (entry *Entry) SetStatus(status int) int
```

SetStatus custom sets the status for the job.

​	SetStatus 自定义设置作业的状态。

#### (*Entry) SetTimes

```go
func (entry *Entry) SetTimes(times int)
```

SetTimes sets the limit running times for the job.

​	SetTimes 设置作业的限制运行时间。

#### (*Entry) Start

```go
func (entry *Entry) Start()
```

Start starts the job.

​	“开始”启动作业。

#### (*Entry) Status

```go
func (entry *Entry) Status() int
```

Status returns the status of the job.

​	Status 返回作业的状态。

#### (*Entry) Stop

```go
func (entry *Entry) Stop()
```

Stop stops the job.

​	停止停止作业。

### type JobFunc

```go
type JobFunc = func(ctx context.Context)
```

JobFunc is the timing called job function in timer.

​	JobFunc 是计时器中称为作业函数的计时。

### type Timer

```go
type Timer struct {
	// contains filtered or unexported fields
}
```

Timer is the timer manager, which uses ticks to calculate the timing interval.

​	计时器是计时器管理器，它使用刻度来计算计时间隔。

#### func New

```go
func New(options ...TimerOptions) *Timer
```

New creates and returns a Timer.

​	new 创建并返回一个计时器。

#### (*Timer) Add

```go
func (t *Timer) Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

Add adds a timing job to the timer, which runs in interval of `interval`.

​	Add 将计时作业添加到计时器，该作业以 的间隔运行。 `interval`

#### (*Timer) AddEntry

```go
func (t *Timer) AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry
```

AddEntry adds a timing job to the timer with detailed parameters.

​	AddEntry 将计时作业添加到具有详细参数的计时器中。

The parameter `interval` specifies the running interval of the job.

​	该参数 `interval` 指定作业的运行间隔。

The parameter `singleton` specifies whether the job running in singleton mode. There’s only one of the same job is allowed running when it’s a singleton mode job.

​	该参数 `singleton` 指定作业是否在单例模式下运行。当作业是单例模式作业时，只允许运行一个相同的作业。

The parameter `times` specifies limit for the job running times, which means the job exits if its run times exceeds the `times`.

​	该参数 `times` 指定作业运行时间的限制，这意味着如果作业的运行时间超过 . `times`

The parameter `status` specifies the job status when it’s firstly added to the timer.

​	该参数 `status` 指定首次添加到计时器时的作业状态。

#### (*Timer) AddOnce

```go
func (t *Timer) AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddOnce is a convenience function for adding a job which only runs once and then exits.

​	AddOnce 是一个方便的函数，用于添加仅运行一次然后退出的作业。

#### (*Timer) AddSingleton

```go
func (t *Timer) AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddSingleton is a convenience function for add singleton mode job.

​	AddSingleton 是添加单例模式作业的便捷函数。

#### (*Timer) AddTimes

```go
func (t *Timer) AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry
```

AddTimes is a convenience function for adding a job which is limited running times.

​	AddTimes 是一个方便的功能，用于添加运行时间有限的作业。

#### (*Timer) Close

```go
func (t *Timer) Close()
```

Close closes the timer.

​	关闭 关闭计时器。

#### (*Timer) DelayAdd

```go
func (t *Timer) DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAdd adds a timing job after delay of `delay` duration. Also see Add.

​	DelayAdd 在 `delay` 持续时间延迟后添加计时作业。另请参阅添加。

#### (*Timer) DelayAddEntry

```go
func (t *Timer) DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int)
```

DelayAddEntry adds a timing job after delay of `delay` duration. Also see AddEntry.

​	DelayAddEntry 在 `delay` 持续时间延迟后添加计时作业。另请参阅 AddEntry。

#### (*Timer) DelayAddOnce

```go
func (t *Timer) DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddOnce adds a timing job after delay of `delay` duration. Also see AddOnce.

​	DelayAddOnce 在 `delay` 持续时间延迟后添加计时作业。另请参阅 AddOnce。

#### (*Timer) DelayAddSingleton

```go
func (t *Timer) DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddSingleton adds a timing job after delay of `delay` duration. Also see AddSingleton.

​	DelayAddSingleton 在 `delay` 持续时间延迟后添加计时作业。另请参阅 AddSingleton。

#### (*Timer) DelayAddTimes

```go
func (t *Timer) DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc)
```

DelayAddTimes adds a timing job after delay of `delay` duration. Also see AddTimes.

​	DelayAddTimes 在 `delay` 持续时间延迟后添加计时作业。另请参阅 AddTimes。

#### (*Timer) Start

```go
func (t *Timer) Start()
```

Start starts the timer.

​	开始启动计时器。

#### (*Timer) Stop

```go
func (t *Timer) Stop()
```

Stop stops the timer.

​	停止停止计时器。

### type TimerOptions

```go
type TimerOptions struct {
	Interval time.Duration // (optional) Interval is the underlying rolling interval tick of the timer.
	Quick    bool          // Quick is used for quick timer, which means the timer will not wait for the first interval to be elapsed.
}
```

TimerOptions is the configuration object for Timer.

​	TimerOptions 是 Timer 的配置对象。

#### func DefaultOptions

```go
func DefaultOptions() TimerOptions
```

DefaultOptions creates and returns a default options object for Timer creation.

​	DefaultOptions 创建并返回用于创建计时器的默认选项对象。