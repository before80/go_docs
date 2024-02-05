+++
title = "Go for App Engine is now generally available"
weight = 14
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go for App Engine is now generally available

> 原文：[https://go.dev/blog/appengine-ga](https://go.dev/blog/appengine-ga)

Andrew Gerrand
21 July 2011

The Go and App Engine teams are excited to announce that the Go runtime for App Engine is now generally available. This means you can take that Go app you’ve been working on (or meaning to work on) and deploy it to App Engine right now with the new [1.5.2 SDK](http://code.google.com/appengine/downloads.html).

Go和App Engine团队很高兴地宣布，适用于App Engine的Go运行时现在已经普遍可用。这意味着您可以利用新的1.5.2 SDK将您一直在做的（或打算做的）Go应用部署到App Engine上。

Since we announced the Go runtime at Google I/O we have continued to [improve and extend](http://code.google.com/p/googleappengine/wiki/SdkForGoReleaseNotes) Go support for the App Engine APIs and have added the Channels API. The Go Datastore API now supports transactions and ancestor queries, too. See the [Go App Engine documentation](https://code.google.com/appengine/docs/go/) for all the details.

自从我们在Google I/O上宣布Go运行时以来，我们一直在改进和扩展Go对App Engine API的支持，并增加了Channels API。Go Datastore API 现在也支持交易和祖先查询。所有细节见Go App Engine文档。

For those who have been using the Go SDK already, please note that the 1.5.2 release introduces `api_version` 2. This is because the new SDK is based on Go `release.r58.1` (the current stable version of Go) and is not backwards compatible with the previous release. Existing apps may require changes as per the [r58 release notes](https://go.dev/doc/devel/release.html#r58). Once you’ve updated your code, you should redeploy your app with the line `api_version: 2` in its `app.yaml` file. Apps written against `api_version` 1 will stop working after the 18th of August.

对于那些已经在使用Go SDK的用户，请注意1.5.2版本引入了api_version 2。这是因为新的SDK是基于Go release.r58.1（Go的当前稳定版本），并不向后兼容以前的版本。现有的应用程序可能需要按照r58的发布说明进行修改。一旦您更新了您的代码，您应该重新部署您的应用程序，在其app.yaml文件中加入api_version: 2一行。针对api_version 1编写的应用程序将在8月18日之后停止工作。

Finally, we owe a huge thanks to our trusted testers and their many bug reports. Their help was invaluable in reaching this important milestone.

最后，我们要感谢我们信任的测试人员和他们的许多错误报告。他们的帮助对于达到这个重要的里程碑是非常宝贵的。

*The fastest way to get started with Go on App Engine is with the* [*Getting Started guide*](http://code.google.com/appengine/docs/go/gettingstarted/).

在App Engine上开始使用Go的最快方法是使用入门指南。

*Note that the Go runtime is still considered experimental; it is not as well-supported as the Python and Java runtimes.*

请注意，Go运行时仍被认为是实验性的；它不像Python和Java运行时那样得到良好的支持。
