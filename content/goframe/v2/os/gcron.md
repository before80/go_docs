+++
title = "gcron"
date = 2024-03-21T17:55:06+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcron

Package gcron implements a cron pattern parser and job runner.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gcron/gcron.go#L18)

``` go
const (
	StatusReady   = gtimer.StatusReady
	StatusRunning = gtimer.StatusRunning
	StatusStopped = gtimer.StatusStopped
	StatusClosed  = gtimer.StatusClosed
)
```

### Variables 

This section is empty.

### Functions 

##### func DelayAdd 

``` go
func DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAdd adds a timed task to default cron object after `delay` time.

##### func DelayAddOnce 

``` go
func DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddOnce adds a timed task after `delay` time to default cron object. This timed task can be run only once.

##### func DelayAddSingleton 

``` go
func DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddSingleton adds a singleton timed task after `delay` time to default cron object.

##### func DelayAddTimes 

``` go
func DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string)
```

DelayAddTimes adds a timed task after `delay` time to default cron object. This timed task can be run specified times.

##### func GetLogger 

``` go
func GetLogger() glog.ILogger
```

GetLogger returns the global logger in the cron.

##### func Remove 

``` go
func Remove(name string)
```

Remove deletes scheduled task which named `name`.

##### func SetLogger 

``` go
func SetLogger(logger glog.ILogger)
```

SetLogger sets the global logger for cron.

##### func Size 

``` go
func Size() int
```

Size returns the size of the timed tasks of default cron.

##### func Start 

``` go
func Start(name ...string)
```

Start starts running the specified timed task named `name`. If no`name` specified, it starts the entire cron.

##### func Stop 

``` go
func Stop(name ...string)
```

Stop stops running the specified timed task named `name`. If no`name` specified, it stops the entire cron.

### Types 

#### type Cron 

``` go
type Cron struct {
	// contains filtered or unexported fields
}
```

Cron stores all the cron job entries.

##### func New 

``` go
func New() *Cron
```

New returns a new Cron object with default settings.

##### (*Cron) Add 

``` go
func (c *Cron) Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

Add adds a timed task. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

##### (*Cron) AddEntry 

``` go
func (c *Cron) AddEntry(
	ctx context.Context, pattern string, job JobFunc, times int, isSingleton bool, name ...string,
) (*Entry, error)
```

AddEntry creates and returns a new Entry object.

##### (*Cron) AddOnce 

``` go
func (c *Cron) AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddOnce adds a timed task which can be run only once. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

##### (*Cron) AddSingleton 

``` go
func (c *Cron) AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddSingleton adds a singleton timed task. A singleton timed task is that can only be running one single instance at the same time. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

##### (*Cron) AddTimes 

``` go
func (c *Cron) AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error)
```

AddTimes adds a timed task which can be run specified times. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

##### (*Cron) Close 

``` go
func (c *Cron) Close()
```

Close stops and closes current cron.

##### (*Cron) DelayAdd 

``` go
func (c *Cron) DelayAdd(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAdd adds a timed task after `delay` time.

##### (*Cron) DelayAddEntry 

``` go
func (c *Cron) DelayAddEntry(ctx context.Context, delay time.Duration, pattern string, job JobFunc, times int, isSingleton bool, name ...string)
```

DelayAddEntry adds a timed task after `delay` time.

##### (*Cron) DelayAddOnce 

``` go
func (c *Cron) DelayAddOnce(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddOnce adds a timed task after `delay` time. This timed task can be run only once.

##### (*Cron) DelayAddSingleton 

``` go
func (c *Cron) DelayAddSingleton(ctx context.Context, delay time.Duration, pattern string, job JobFunc, name ...string)
```

DelayAddSingleton adds a singleton timed task after `delay` time.

##### (*Cron) DelayAddTimes 

``` go
func (c *Cron) DelayAddTimes(ctx context.Context, delay time.Duration, pattern string, times int, job JobFunc, name ...string)
```

DelayAddTimes adds a timed task after `delay` time. This timed task can be run specified times.

##### (*Cron) Entries 

``` go
func (c *Cron) Entries() []*Entry
```

Entries return all timed tasks as slice(order by registered time asc).

##### (*Cron) GetLogger 

``` go
func (c *Cron) GetLogger() glog.ILogger
```

GetLogger returns the logger in the cron.

##### (*Cron) Remove 

``` go
func (c *Cron) Remove(name string)
```

Remove deletes scheduled task which named `name`.

##### (*Cron) Search 

``` go
func (c *Cron) Search(name string) *Entry
```

Search returns a scheduled task with the specified `name`. It returns nil if not found.

##### (*Cron) SetLogger 

``` go
func (c *Cron) SetLogger(logger glog.ILogger)
```

SetLogger sets the logger for cron.

##### (*Cron) Size 

``` go
func (c *Cron) Size() int
```

Size returns the size of the timed tasks.

##### (*Cron) Start 

``` go
func (c *Cron) Start(name ...string)
```

Start starts running the specified timed task named `name`. If no`name` specified, it starts the entire cron.

##### (*Cron) Stop 

``` go
func (c *Cron) Stop(name ...string)
```

Stop stops running the specified timed task named `name`. If no`name` specified, it stops the entire cron.

#### type Entry 

``` go
type Entry struct {
	Name string    // Entry name.
	Job  JobFunc   `json:"-"` // Callback function.
	Time time.Time // Registered time.
	// contains filtered or unexported fields
}
```

Entry is timing task entry.

##### func Add 

``` go
func Add(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

Add adds a timed task to default cron object. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

##### func AddOnce 

``` go
func AddOnce(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddOnce adds a timed task which can be run only once, to default cron object. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

##### func AddSingleton 

``` go
func AddSingleton(ctx context.Context, pattern string, job JobFunc, name ...string) (*Entry, error)
```

AddSingleton adds a singleton timed task, to default cron object. A singleton timed task is that can only be running one single instance at the same time. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

##### func AddTimes 

``` go
func AddTimes(ctx context.Context, pattern string, times int, job JobFunc, name ...string) (*Entry, error)
```

AddTimes adds a timed task which can be run specified times, to default cron object. A unique `name` can be bound with the timed task. It returns and error if the `name` is already used.

##### func Entries 

``` go
func Entries() []*Entry
```

Entries return all timed tasks as slice.

##### func Search 

``` go
func Search(name string) *Entry
```

Search returns a scheduled task with the specified `name`. It returns nil if no found.

##### (*Entry) Close 

``` go
func (e *Entry) Close()
```

Close stops and removes the entry from cron.

##### (*Entry) IsSingleton 

``` go
func (e *Entry) IsSingleton() bool
```

IsSingleton return whether this entry is a singleton timed task.

##### (*Entry) SetSingleton 

``` go
func (e *Entry) SetSingleton(enabled bool)
```

SetSingleton sets the entry running in singleton mode.

##### (*Entry) SetStatus 

``` go
func (e *Entry) SetStatus(status int) int
```

SetStatus sets the status of the entry.

##### (*Entry) SetTimes 

``` go
func (e *Entry) SetTimes(times int)
```

SetTimes sets the times which the entry can run.

##### (*Entry) Start 

``` go
func (e *Entry) Start()
```

Start starts running the entry.

##### (*Entry) Status 

``` go
func (e *Entry) Status() int
```

Status returns the status of entry.

##### (*Entry) Stop 

``` go
func (e *Entry) Stop()
```

Stop stops running the entry.

#### type JobFunc 

``` go
type JobFunc = gtimer.JobFunc
```

JobFunc is the timing called job function in cron.