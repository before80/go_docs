+++
title = "go 1.14版发布了"
weight = 14
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.14 is released - go 1.14版发布了

https://go.dev/blog/go1.14

Alex Rakoczy
25 February 2020

Today the Go team is very happy to announce the release of Go 1.14. You can get it from the [download page](https://go.dev/dl).

今天Go团队非常高兴地宣布Go 1.14的发布。您可以从下载页面获得它。

Some of the highlights include:

其中的一些亮点包括：

- Module support in the `go` command is now ready for production use. We encourage all users to [migrate to `go` modules for dependency management](https://go.dev/doc/go1.14#introduction). go命令中的模块支持现在已经可以用于生产了。我们鼓励所有用户迁移到go模块来进行依赖项管理。
- [Embedding interfaces with overlapping method sets ](https://go.dev/doc/go1.14#language)嵌入具有重叠方法集的接口
- [Improved defer performance](https://go.dev/doc/go1.14#runtime) 改进了defer的性能
- [Goroutines are asynchronously preemptible](https://go.dev/doc/go1.14#runtime) Goroutines是可异步抢占的
- [The page allocator is more efficient](https://go.dev/doc/go1.14#runtime) 页面分配器更加高效
- [Internal timers are more efficient](https://go.dev/doc/go1.14#runtime) 内部定时器更有效

For the complete list of changes and more information about the improvements above, see the [**Go 1.14 release notes**](https://go.dev/doc/go1.14).

关于完整的变化列表和有关上述改进的更多信息，请参见Go 1.14发布说明。

We want to thank everyone who contributed to this release by writing code, filing bugs, providing feedback, and/or testing the beta and release candidate. Your contributions and diligence helped to ensure that Go 1.14 is as stable as possible. That said, if you notice any problems, please [file an issue](https://go.dev/issue/new).

我们要感谢所有通过编写代码、提交错误、提供反馈和/或测试测试版和候选版而为这个版本做出贡献的人。您的贡献和勤奋有助于确保Go 1.14尽可能的稳定。也就是说，如果您发现任何问题，请提出问题。

We hope you enjoy the new release!

我们希望您喜欢这个新版本
