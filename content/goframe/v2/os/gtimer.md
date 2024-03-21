+++
title = "gtimer"
date = 2024-03-21T17:57:40+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gtimer

Package gtimer implements timer for interval/delayed jobs running and management.

This package is designed for management for millions of timing jobs. The differences between gtimer and gcron are as follows:

1. package gcron is implemented based on package gtimer.
2. gtimer is designed for high performance and for millions of timing jobs.
3. gcron supports configuration pattern grammar like linux crontab, which is more manually readable.
4. gtimer's benchmark OP is measured in nanoseconds, and gcron's benchmark OP is measured in microseconds.

ALSO VERY NOTE the common delay of the timer: https://github.com/golang/go/issues/14410

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gtimer/gtimer.go#L51)

``` go
const (
	StatusReady   = 0  // Job or Timer is ready for running.
	StatusRunning = 1  // Job or Timer is already running.
	StatusStopped = 2  // Job or Timer is stopped.
	StatusClosed  = -1 // Job or Timer is closed and waiting to be deleted.

)
```

### Variables 

This section is empty.

### Functions 

##### func DelayAdd 

``` go
func DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAdd adds a timing job after delay of `interval` duration. Also see Add.

##### func DelayAddEntry 

``` go
func DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int)
```

DelayAddEntry adds a timing job after delay of `interval` duration. Also see AddEntry.

##### func DelayAddOnce 

``` go
func DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddOnce adds a timing job after delay of `interval` duration. Also see AddOnce.

##### func DelayAddSingleton 

``` go
func DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddSingleton adds a timing job after delay of `interval` duration. Also see AddSingleton.

##### func DelayAddTimes 

``` go
func DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc)
```

DelayAddTimes adds a timing job after delay of `interval` duration. Also see AddTimes.

##### func Exit 

``` go
func Exit()
```

Exit is used in timing job internally, which exits and marks it closed from timer. The timing job will be automatically removed from timer later. It uses "panic-recover" mechanism internally implementing this feature, which is designed for simplification and convenience.

##### func SetInterval 

``` go
func SetInterval(ctx context.Context, interval time.Duration, job JobFunc)
```

SetInterval runs the job every duration of `delay`. It is like the one in javascript.

##### func SetTimeout 

``` go
func SetTimeout(ctx context.Context, delay time.Duration, job JobFunc)
```

SetTimeout runs the job once after duration of `delay`. It is like the one in javascript.

### Types 

#### type Entry 

``` go
type Entry struct {
	// contains filtered or unexported fields
}
```

Entry is the timing job.

##### func Add 

``` go
func Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

Add adds a timing job to the default timer, which runs in interval of `interval`.

##### Example

``` go
```
##### func AddEntry 

``` go
func AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry
```

AddEntry adds a timing job to the default timer with detailed parameters.

The parameter `interval` specifies the running interval of the job.

The parameter `singleton` specifies whether the job running in singleton mode. There's only one of the same job is allowed running when its a singleton mode job.

The parameter `times` specifies limit for the job running times, which means the job exits if its run times exceeds the `times`.

The parameter `status` specifies the job status when it's firstly added to the timer.

##### func AddOnce 

``` go
func AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddOnce is a convenience function for adding a job which only runs once and then exits.

##### func AddSingleton 

``` go
func AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddSingleton is a convenience function for add singleton mode job.

##### func AddTimes 

``` go
func AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry
```

AddTimes is a convenience function for adding a job which is limited running times.

##### (*Entry) Close 

``` go
func (entry *Entry) Close()
```

Close closes the job, and then it will be removed from the timer.

##### (*Entry) Ctx 

``` go
func (entry *Entry) Ctx() context.Context
```

Ctx returns the initialized context of this job.

##### (*Entry) IsSingleton 

``` go
func (entry *Entry) IsSingleton() bool
```

IsSingleton checks and returns whether the job in singleton mode.

##### (*Entry) Job 

``` go
func (entry *Entry) Job() JobFunc
```

Job returns the job function of this job.

##### (*Entry) Reset 

``` go
func (entry *Entry) Reset()
```

Reset resets the job, which resets its ticks for next running.

##### (*Entry) Run 

``` go
func (entry *Entry) Run()
```

Run runs the timer job asynchronously.

##### (*Entry) SetSingleton 

``` go
func (entry *Entry) SetSingleton(enabled bool)
```

SetSingleton sets the job singleton mode.

##### (*Entry) SetStatus 

``` go
func (entry *Entry) SetStatus(status int) int
```

SetStatus custom sets the status for the job.

##### (*Entry) SetTimes 

``` go
func (entry *Entry) SetTimes(times int)
```

SetTimes sets the limit running times for the job.

##### (*Entry) Start 

``` go
func (entry *Entry) Start()
```

Start starts the job.

##### (*Entry) Status 

``` go
func (entry *Entry) Status() int
```

Status returns the status of the job.

##### (*Entry) Stop 

``` go
func (entry *Entry) Stop()
```

Stop stops the job.

#### type JobFunc 

``` go
type JobFunc = func(ctx context.Context)
```

JobFunc is the timing called job function in timer.

#### type Timer 

``` go
type Timer struct {
	// contains filtered or unexported fields
}
```

Timer is the timer manager, which uses ticks to calculate the timing interval.

##### func New 

``` go
func New(options ...TimerOptions) *Timer
```

New creates and returns a Timer.

##### (*Timer) Add 

``` go
func (t *Timer) Add(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

Add adds a timing job to the timer, which runs in interval of `interval`.

##### (*Timer) AddEntry 

``` go
func (t *Timer) AddEntry(ctx context.Context, interval time.Duration, job JobFunc, isSingleton bool, times int, status int) *Entry
```

AddEntry adds a timing job to the timer with detailed parameters.

The parameter `interval` specifies the running interval of the job.

The parameter `singleton` specifies whether the job running in singleton mode. There's only one of the same job is allowed running when it's a singleton mode job.

The parameter `times` specifies limit for the job running times, which means the job exits if its run times exceeds the `times`.

The parameter `status` specifies the job status when it's firstly added to the timer.

##### (*Timer) AddOnce 

``` go
func (t *Timer) AddOnce(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddOnce is a convenience function for adding a job which only runs once and then exits.

##### (*Timer) AddSingleton 

``` go
func (t *Timer) AddSingleton(ctx context.Context, interval time.Duration, job JobFunc) *Entry
```

AddSingleton is a convenience function for add singleton mode job.

##### (*Timer) AddTimes 

``` go
func (t *Timer) AddTimes(ctx context.Context, interval time.Duration, times int, job JobFunc) *Entry
```

AddTimes is a convenience function for adding a job which is limited running times.

##### (*Timer) Close 

``` go
func (t *Timer) Close()
```

Close closes the timer.

##### (*Timer) DelayAdd 

``` go
func (t *Timer) DelayAdd(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAdd adds a timing job after delay of `delay` duration. Also see Add.

##### (*Timer) DelayAddEntry 

``` go
func (t *Timer) DelayAddEntry(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc, isSingleton bool, times int, status int)
```

DelayAddEntry adds a timing job after delay of `delay` duration. Also see AddEntry.

##### (*Timer) DelayAddOnce 

``` go
func (t *Timer) DelayAddOnce(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddOnce adds a timing job after delay of `delay` duration. Also see AddOnce.

##### (*Timer) DelayAddSingleton 

``` go
func (t *Timer) DelayAddSingleton(ctx context.Context, delay time.Duration, interval time.Duration, job JobFunc)
```

DelayAddSingleton adds a timing job after delay of `delay` duration. Also see AddSingleton.

##### (*Timer) DelayAddTimes 

``` go
func (t *Timer) DelayAddTimes(ctx context.Context, delay time.Duration, interval time.Duration, times int, job JobFunc)
```

DelayAddTimes adds a timing job after delay of `delay` duration. Also see AddTimes.

##### (*Timer) Start 

``` go
func (t *Timer) Start()
```

Start starts the timer.

##### (*Timer) Stop 

``` go
func (t *Timer) Stop()
```

Stop stops the timer.

#### type TimerOptions 

``` go
type TimerOptions struct {
	Interval time.Duration // (optional) Interval is the underlying rolling interval tick of the timer.
	Quick    bool          // Quick is used for quick timer, which means the timer will not wait for the first interval to be elapsed.
}
```

TimerOptions is the configuration object for Timer.

##### func DefaultOptions 

``` go
func DefaultOptions() TimerOptions
```

DefaultOptions creates and returns a default options object for Timer creation.