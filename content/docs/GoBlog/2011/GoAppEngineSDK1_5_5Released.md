+++
title = "Go App Engine SDK 1.5.5 released"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go App Engine SDK 1.5.5 released

https://go.dev/blog/appengine-155

Andrew Gerrand
11 October 2011

Today we released version 1.5.5 the Go App Engine SDK. You can download it from the [App Engine downloads page](http://code.google.com/appengine/downloads.html).

今天我们发布了1.5.5版本的Go App Engine SDK。您可以从App Engine的下载页面下载它。

This release includes changes and improvements to the App Engine APIs and brings the supporting Go tool chain to [release.r60.2](https://go.dev/doc/devel/release.html#r60) (the current stable release). Also included in this release are the [godoc](https://go.dev/cmd/godoc/), [gofmt](https://go.dev/cmd/gofmt/), and [gofix](https://go.dev/cmd/gofix/) tools from the Go tool chain. They can be found in the root directory of the SDK.

这个版本包括对App Engine APIs的修改和改进，并将支持Go的工具链提升到release.r60.2（当前的稳定版本）。这个版本还包括Go工具链中的godoc、gofmt和gofix工具。它们可以在SDK的根目录中找到。

Some changes made in this release are backwards-incompatible, so we have incremented the SDK `api_version` to 3. Existing apps will require code changes when migrating to `api_version` 3.

这个版本中的一些变化是向后兼容的，所以我们将SDK的api_version增加到了3。现有的应用程序在迁移到api_version 3时需要修改代码。

The gofix tool that ships with the SDK has been customized with App Engine-specific modules. It can be used to automatically update Go apps to work with the latest appengine packages and the updated Go standard library. To update your apps, run:

与SDK一起提供的gofix工具已经用App Eng的特定模块进行了定制。它可以用来自动更新Go应用程序，使其与最新的appengine软件包和更新的Go标准库一起使用。要更新您的应用程序，请运行：

```go linenums="1"
/path/to/sdk/gofix /path/to/your/app
```

The SDK now includes the appengine package source code, so you can use the local godoc to read App Engine API documentation:

SDK现在包括appengine包的源代码，所以您可以使用本地godoc来阅读App Engine API文档：

```go linenums="1"
/path/to/sdk/godoc appengine/datastore Get
```

**Important note:** We have deprecated `api_version` 2. Go apps that use `api_version` 2 will stop working after 16 December 2011. Please update your apps to use `api_version` 3 before then.

重要提示：我们已经废弃了api_version 2。使用api_version 2的Go应用程序将在2011年12月16日后停止工作。请在那之前更新您的应用程序以使用api_version 3。

See the [release notes](http://code.google.com/p/googleappengine/wiki/SdkForGoReleaseNotes) for a full list of changes. Please direct any questions about the new SDK to the [Go App Engine discussion group](http://groups.google.com/group/google-appengine-go).

请参阅发行说明以了解完整的变化列表。关于新SDK的任何问题，请直接向Go App Engine讨论组提出。
