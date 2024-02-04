+++
title = "grace 模块"
date = 2024-02-04T09:25:05+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/module/grace/](https://beego.wiki/docs/module/grace/)

# grace 模块



## What is Hot Upgrade? 什么是热升级？

What is hot upgrade? Students who know nginx know that nginx supports hot upgrade. It can use the old process to serve the previous link and the new process to serve the new link, that is, to complete the system upgrade and operating parameter modification without stopping the service. Then hot upgrade and hot compilation are different concepts. Hot compilation is to recompile by monitoring file changes, and then restart the process, For example `bee run` is such a tool

&zeroWidthSpace;什么是热升级？了解 nginx 的同学都知道 nginx 支持热升级，它可以利用旧进程服务旧链接，新进程服务新链接，即在不停止服务的情况下完成系统升级和运行参数修改。那么热升级和热编译是不同的概念，热编译是通过监控文件变化重新编译，然后重启进程，比如 `bee run` 就是这样的工具

## Is hot upgrade necessary？ 热升级有必要吗？

Many people think that it is necessary for HTTP applications to support hot upgrades? Then I can say very responsibly that it is very necessary. Uninterrupted service is always the goal we pursue. Although many people say that the server may be broken, etc., this belongs to the high-availability design category. Don’t confuse it. This is Predictable problems, so we need to avoid user unavailability caused by such an upgrade. Are you still troubled by the early morning upgrade for the previous upgrade? So hurry up and embrace the hot upgrade now.

&zeroWidthSpace;很多人认为 HTTP 应用有必要支持热升级吗？那么我可以很负责任的说，很有必要，不间断服务永远是我们追求的目标，虽然很多人说服务器可能会坏等，这个属于高可用设计范畴，不要混为一谈，这是可预见的问题，所以我们需要避免因为这样的升级导致用户不可用，你还在为以前升级凌晨升级而烦恼吗？那么现在赶紧拥抱热升级吧。

## grace module grace 模块

The grace module is a newly added module of beego that independently supports hot restart. The main idea comes from: http://grisha.org/blog/2014/06/03/graceful-restart-in-golang/

&zeroWidthSpace;grace 模块是 beego 新增加的一个模块，独立支持热重启。主要思想来源于：http://grisha.org/blog/2014/06/03/graceful-restart-in-golang/

## How to use hot upgrade 如何使用热升级

```
 import(
   "log"
	"net/http"
	"os"
    "strconv"

   "github.com/beego/beego/v2/server/web/grace"
 )

  func handler(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("WORLD!"))
      w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
  }

  func main() {
      mux := http.NewServeMux()
      mux.HandleFunc("/hello", handler)

      err := grace.ListenAndServe("localhost:8080", mux)
      if err != nil {
		   log.Println(err)
	    }
      log.Println("Server on 8080 stopped")
	     os.Exit(0)
    }
```

open two terminals

&zeroWidthSpace;打开两个终端

A terminal input: `ps -ef|grep appname`

&zeroWidthSpace;一个终端输入： `ps -ef|grep appname`

A terminal input request: `curl "http://127.0.0.1:8080/hello"`

&zeroWidthSpace;一个终端输入请求： `curl "http://127.0.0.1:8080/hello"`

hot upgrade 
&zeroWidthSpace;热升级

kill -HUP process ID

&zeroWidthSpace;kill -HUP 进程ID

Open a terminal and enter the request: `curl "http://127.0.0.1:8080/hello?sleep=0"`

&zeroWidthSpace;打开一个终端，输入请求： `curl "http://127.0.0.1:8080/hello?sleep=0"`