+++
title = "task 模块"
date = 2024-02-04T09:31:48+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/module/task/]({{< ref "/beego/modules/task" >}})

# Task Module 任务模块



## Core task Module 核心任务模块

## Installation 安装

```
go get github.com/beego/beego/v2/task
```

## Tasks 任务

Tasks work very similarly to cron jobs. Tasks are used to run a job outside the normal request/response cycle. These can be adhoc or scheduled to run regularly. Examples include: Reporting memory and goroutine status, periodically triggering GC or cleaning up log files at fixed intervals.

​	任务的工作方式与 cron 作业非常相似。任务用于在常规请求/响应周期之外运行作业。这些可以是临时性的，也可以被安排定期运行。示例包括：报告内存和 goroutine 状态、定期触发 GC 或按固定间隔清理日志文件。

### Creating a new Task 创建新任务

To initialize a task implement :

​	要初始化任务，请实现：

```
tk1 := task.NewTask("tk1", "0 12 * * * *", func(ctx context.Context) error {
	fmt.Println("tk1")
	return nil
})
```

The NewTask signature:

​	NewTask 签名：

```
NewTask(tname string, spec string, f TaskFunc) *Task	
```

- `tname`: Task name
  `tname` ：任务名称
- `spec`: Task format. See below for details.
  `spec` ：任务格式。有关详细信息，请参见下文。
- `f`: The function which will be run as the task.
  `f` ：将作为任务运行的函数。

To implement this task, add it to the global task list and start it.

​	要实现此任务，请将其添加到全局任务列表并启动它。

```
task.AddTask("tk1", tk1)
task.StartTask()
defer task.StopTask()
```

### Testing the TaskFunc 测试 TaskFunc

Use the code below to test if the TaskFunc is working correctly.

​	使用以下代码测试 TaskFunc 是否工作正常。

```
err := tk.Run()
if err != nil {
	t.Fatal(err)
}
```

### spec in detail 详细说明

`spec` specifies when the new Task will be run. Its format is the same as that of traditional crontab:

​	 `spec` 指定新任务何时运行。其格式与传统的 crontab 相同：

```
// The first 6 parts are:
//       second: 0-59
//       minute: 0-59
//       hour: 1-23
//       day: 1-31
//       month: 1-12
//       weekdays: 0-6（0 is Sunday）

// Some special sign:
//       *: any time
//       ,: separator. E.g.: 2,4 in the third part means run at 2 and 4 o'clock
//　　    －: range. E.g.: 1-5 in the third part means run between 1 and 5 o'clock
//       /n : run once every n time. E.g.: */1 in the third part means run once every an hour. Same as 1-23/1
/////////////////////////////////////////////////////////
//	0/30 * * * * *                        run every 30 seconds
//	0 43 21 * * *                         run at 21:43
//	0 15 05 * * *                         run at 05:15
//	0 0 17 * * *                          run at 17:00
//	0 0 17 * * 1                          run at 17:00 of every Monday
//	0 0,10 17 * * 0,2,3                   run at 17:00 and 17:10 of every Sunday, Tuesday and Wednesday
//	0 0-10 17 1 * *                       run once every minute from 17:00 to 7:10 on 1st day of every month
//	0 0 0 1,15 * 1                        run at 0:00 on 1st and 15th of each month and every Monday
//	0 42 4 1 * *                          run at 4:42 on 1st of every month
//	0 0 21 * * 1-6                        run at 21:00 from Monday to Saturday
//	0 0,10,20,30,40,50 * * * *            run every 10 minutes
//	0 */10 * * * *                        run every 10 minutes
//	0 * 1 * * *                           run every one minute from 1:00 to 1:59
//	0 0 1 * * *                           run at 1:00
//	0 0 */1 * * *                         run at :00 of every hour
//	0 0 * * * *                           run at :00 of every hour
//	0 2 8-20/3 * * *                      run at 8:02, 11:02, 14:02, 17:02 and 20:02
//	0 30 5 1,15 * *                       run at 5:30 of 1st and 15th of every month
```

## Debug module (Already moved to utils module) 调试模块（已移至 utils 模块）

We always use print for debugging. But the default output is not good enough for debugging. Beego provides this debug module

​	我们始终使用 print 进行调试。但默认输出不足以进行调试。Beego 提供此调试模块

- Display() print result to console
  Display() 将打印结果打印到控制台
- GetDisplayString() return the string
  GetDisplayString() 返回字符串

It print key/value pairs. The following code:

​	它打印键/值对。以下代码：

```
Display("v1", 1, "v2", 2, "v3", 3)
```

will output:

​	将输出：

```
2013/12/16 23:48:41 [Debug] at TestPrint() [/Users/astaxie/github/beego/task/debug_test.go:13]

[Variables]
v1 = 1
v2 = 2
v3 = 3	
```

For pointer type:

​	对于指针类型：

```
type mytype struct {
	next *mytype
	prev *mytype
}	

var v1 = new(mytype)
var v2 = new(mytype)

v1.prev = nil
v1.next = v2

v2.prev = v1
v2.next = nil

Display("v1", v1, "v2", v2)
```

The output result

​	输出结果

```
2013/12/16 23:48:41 [Debug] at TestPrintPoint() [/Users/astaxie/github/beego/task/debug_test.go:26]

[Variables]
v1 = &task.mytype{
    next: &task.mytype{
        next: nil,
        prev: 0x210335420,
    },
    prev: nil,
}
v2 = &task.mytype{
    next: nil,
    prev: &task.mytype{
        next: 0x210335430,
        prev: nil,
    },
}	
```
