+++
title = "You're holding it wrong"
date = 2024-12-15T11:36:32+08:00
weight = 17
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/You're-holding-it-wrong](https://github.com/smartystreets/goconvey/wiki/You're-holding-it-wrong)
>
> 收录该文档时间： `2024-12-15T11:36:32+08:00`

# You're holding it wrong

mdwhatcott edited this page on Nov 13, 2014 · [9 revisions](https://github.com/smartystreets/goconvey/wiki/You're-holding-it-wrong/_history)

​	mdwhatcott 编辑了此页面，日期为 2014年11月13日 · [9 个修订](https://github.com/smartystreets/goconvey/wiki/You're-holding-it-wrong/_history)

When a customer emailed Steve Jobs about bad reception on the iPhone 4, [Jobs replied and basically told him](http://www.engadget.com/2010/06/24/apple-responds-over-iphone-4-reception-issues-youre-holding-th/), you're holding it wrong.

​	当有顾客给 Steve Jobs 发邮件，反映 iPhone 4 信号差时，[Jobs 回复并基本上告诉他](http://www.engadget.com/2010/06/24/apple-responds-over-iphone-4-reception-issues-youre-holding-th/)，是你拿错了。

Similarly, here are some of the wrong ways to "hold" GoConvey, some of which are not very obvious (and understandably so).

​	类似地，以下是一些错误的方式来“使用” GoConvey，其中一些可能并不那么明显（这也是可以理解的）。

### • Go files not in `src` folder in a Go workspace

Go 文件不在 Go 工作区的 `src` 文件夹中

When using the web UI, Go files being tested must be in the GOPATH under the `src` directory. The server expects the `src` directory to be present after the GOPATH so it can determine the full package name.

​	在使用 Web UI 时，被测试的 Go 文件必须位于 GOPATH 下的 `src` 目录中。服务器需要 `src` 目录存在于 GOPATH 后，以便它能够确定完整的包名。

### • Calling `panic(nil)`



Never, never, never call `panic(nil)` because GoConvey can't recover the error (it thinks the `Convey` passed) see [issue 98](https://github.com/smartystreets/goconvey/issues/98) and especially [this Stack Overflow question](http://stackoverflow.com/questions/19662527/how-to-detect-panicnil-and-normal-execution-in-deferred-function-go).

​	绝对不要、绝对不要调用 `panic(nil)`，因为 GoConvey 无法恢复该错误（它认为 `Convey` 已通过），请参见 [issue 98](https://github.com/smartystreets/goconvey/issues/98)，特别是 [这个 Stack Overflow 问题](http://stackoverflow.com/questions/19662527/how-to-detect-panicnil-and-normal-execution-in-deferred-function-go)。
