+++
title = "go 模块镜像、索引和校验数据库"
weight = 27
date = 2023-05-18T17:31:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go Module Mirror, Index, and Checksum Database - go 模块镜像、索引和校验数据库

> 原文：[https://proxy.golang.org/](https://proxy.golang.org/)

​	Go团队正在提供以下由谷歌运行的服务：用于加速Go模块下载的模块镜像，用于发现新模块的索引，以及用于验证模块内容的全局go.sum数据库。

​	从Go 1.13开始，go命令默认使用Go模块镜像和Go校验数据库来下载和验证模块。关于这些服务的隐私信息，请参见[proxy.golang.org/privacy](https://proxy.golang.org/privacy)，关于配置细节，包括如何禁用这些服务器或使用不同的服务器，请参见[go命令文档]({{< ref "/docs/References/GoModulesReference/PrivateModules">}})。如果您依赖非公开的模块，请参阅[配置环境的文档]({{< ref "/cmd/go#用于下载非公共代码的配置">}})。

## 服务 Services

[**proxy.golang.org**](https://proxy.golang.org/) - a module mirror which implements the [module proxy protocol](https://golang.org/ref/mod#goproxy-protocol). For users downloading large numbers of modules (e.g. for bulk static analysis), the mirror supports a non-standard header, `Disable-Module-Fetch: true` that instructs it to return only cached content. This will avoid slow downloads, at the cost of possibly missing some rarely-used modules.

**proxy.golang.org** —— 一个模块镜像，实现了[模块代理协议]({{< ref "/docs/References/GoModulesReference/ModuleProxies#goproxy-protocol">}})。对于下载大量模块的用户（例如用于批量静态分析），该镜像支持一个非标准的头：`Disable-Module-Fetch: true`，指示它只返回缓存的内容。这将避免缓慢的下载，但代价是可能会错过一些很少使用的模块。

[**sum.golang.org**](https://sum.golang.org/) - an auditable checksum database which will be used by the go command to authenticate modules. Learn more in the [go command documentation](https://golang.org/ref/mod#checksum-database).

**sum.golang.org** —— 一个可审计的校验数据库，将被go命令用来验证模块。在 [go 命令的文档]({{< ref "/docs/References/GoModulesReference/AuthenticatingModules#checksum-database-校验和数据库">}})中了解更多。

[**index.golang.org**](https://index.golang.org/) - an index which serves a feed of new module versions that become available by proxy.golang.org. The feed can be viewed at https://index.golang.org/index. The feed is served as new line delimited JSON, providing the module path (as Path), the module version (as Version), and the time it was first cached by proxy.golang.org (as Timestamp). The list is sorted in chronological order. There are two optional parameters:

**index.golang.org** —— 一个索引，为proxy.golang.org提供新模块版本的信息。可以在 [https://index.golang.org/index](https://index.golang.org/index ) 上查看该信息。反馈是以新行分隔的JSON格式提供的，提供了模块的路径（Path），模块的版本（Version），以及它第一次被proxy.golang.org缓存的时间（Timestamp）。该列表按时间顺序排序。有两个可选参数：

- 'since': the oldest allowable timestamp (RFC3339 format) for module versions in the returned list. Default is the beginning of time, e.g. https://index.golang.org/index?since=2019-04-10T19:08:52.997264Z
- 'since'：在返回的列表中，模块版本允许的最老的时间戳（RFC3339格式）。默认是时间的开始，例如：[https://index.golang.org/index?since=2019-04-10T19:08:52.997264Z](https://index.golang.org/index?since=2019-04-10T19:08:52.997264Z)
- 'limit': the maximum length of the returned list. Default = 2000, Max = 2000, e.g. https://index.golang.org/index?limit=10
- 'limit'：返回列表的最大长度。Default =2000，Max =2000，例如：[https://index.golang.org/index?limit=10](https://index.golang.org/index?limit=10)

If you use the index to download many modules from the module mirror, you will want to set the `Disable-Module-Fetch` header, described above.

​	如果您使用该索引从模块镜像中下载许多模块，您要按如上所述设置`Disable-Module-Fetch`头。

## 状态：已启动 Status: Launched 

These services are ready for production use. Please [file issues](https://golang.org/issue/new?title=proxy.golang.org: ) if you spot them, with the title prefix "proxy.golang.org:" (or "index.golang.org:", or "sum.golang.org:").

​	这些服务已经可以在生产中使用了。如果您发现了这些问题，请[提交问题](https://golang.org/issue/new?title=proxy.golang.org:)，标题前缀为 "proxy.golang.org:"（或 "index.golang.org:"，或 "sum.golang.org:"）。

## 环境设置 Environment setup 

These services can only access publicly available source code. If you depend on private modules, set `GOPRIVATE` to a glob pattern that covers them. See [Module configuration for non-public modules](https://pkg.go.dev/cmd/go#hdr-Configuration_for_downloading_non_public_code) in the go command documentation for more details.

​	这些服务只能访问公开的源代码。如果您依赖私有模块，请将`GOPRIVATE`设置为涵盖它们的glob模式。更多细节请参见 go 命令文档中的[非公开模块的模块配置]({{< ref "/cmd/go#用于下载非公共代码的配置">}})。

To opt-out of this module mirror, you can turn it off by setting `GOPROXY=direct`

​	要退出这个模块镜像，您可以通过设置`GOPROXY=direct`来关闭它。

See the [go command documentation](https://golang.org/ref/mod#private-module-privacy) for other configuration details.

​	其他配置细节见[go命令文档]({{< ref "/docs/References/GoModulesReference/PrivateModules">}})。

## FAQ

### I committed a new change (or released a new version) to a repository, why isn't it showing up when I run `go get -u` or `go list -m --versions`?

我向存储库提交了一个新的变更（或发布了一个新的版本），为什么当我运行 `go get -u` 或 `go list -m --versions` 时，它没有显示出来？

In order to improve our services' caching and serving latencies, new versions may not show up right away. If you want new code to be immediately available in the mirror, then first make sure there is a semantically versioned tag for this revision in the underlying source repository. Then explicitly request that version via `go get module@version`. The new version should be available within one minute. Note that if someone requested the version before the tag was pushed, it may take up to 30 minutes for the mirror's cache to expire and fresh data about the version to become available. If the version is still not available after 30 minutes, please [file an issue](https://golang.org/issue/new?title=proxy.golang.org%3A+).

​	为了改善我们服务的缓存和服务延迟，新版本可能不会马上显示出来。如果您想让新的代码立即出现在镜像中，那么首先要确保在底层源码存储库中有一个语义上的版本标签。然后通过 `go get module@version` 明确请求该版本。新版本应该在一分钟内可用。请注意，如果有人在标签被推送之前请求该版本，可能需要长达30分钟的时间让镜像的缓存过期，关于该版本的新数据才会变得可用。如果30分钟后该版本仍然不可用，请[提交一个问题](https://golang.org/issue/new?title=proxy.golang.org%3A+)。

### I removed a bad release from my repository but it still appears in the mirror, what should I do?

我从我的存储库中删除了一个坏的版本，但它仍然出现在镜像中，我应该怎么做？

Whenever possible, the mirror aims to cache content in order to avoid breaking builds for people that depend on your package, so this bad release may still be available in the mirror even if it is not available at the origin. The same situation applies if you delete your entire repository. We suggest creating a new version and encouraging people to use that one instead.

​	只要有可能，镜像的目的是缓存内容，以避免破坏依赖您的包的人的构建，所以这个坏版本可能仍然在镜像中可用，即使它在原点（origin）不可用。如果您删除了您的整个存储库，同样的情况也适用。我们建议创建一个新的版本并鼓励人们使用该版本。

If you would like to hide versions of a module from the `go` command, as well as [pkg.go.dev](https://pkg.go.dev/), you should retract them. Retracting a module version involves adding a [retract directive](https://golang.org/ref/mod#go-mod-file-retract) to your go.mod file and publishing a new version. See the Go blog post [New module changes in Go 1.16](https://go.dev/blog/go116-module-changes#module-retraction) and the [modules reference](https://go.dev/ref/mod#go-mod-file-retract) for details.

​	如果您想从 `go` 命令以及 [pkg.go.dev](https://pkg.go.dev/) 中隐藏模块的版本，您应该撤回它们。撤回模块版本需要在go.mod文件中加入[retract指令]({{< ref "/docs/References/GoModulesReference/gomodFiles#retract-directive">}})并发布新的版本。详情请参见Go博客文章[Go 1.16中的新模块变化]({{< ref "/goBlog/2021/NewModuleChangesInGo1_16#module-retraction-模块撤回">}})和[模块参考]({{< ref "/docs/References/gomodFileReference#retract-撤回">}})。

### I'm running the go command in an environment that can't use the mirror.

我在一个不能使用镜像的环境中运行go命令。

The [go command documentation](https://golang.org/ref/mod#private-module-privacy) describes the configuration details including how to disable the use of these servers or use different ones.

​	[go命令文档]({{< ref "/docs/References/GoModulesReference/PrivateModules">}})描述了配置细节，包括如何禁止使用这些服务器或使用不同的服务器。

### If I don't set `GOPRIVATE` and request a private module from these services, what leaks?

如果我不设置`GOPRIVATE`，从这些服务中请求一个私有模块，会有什么泄露？

The proxy and checksum database protocols only send module paths and versions to the remote server. If you request a private module, the mirror will try to download it just as any Go user would and fail in the same way. Information about failed requests isn't published anywhere. The only trace of the request will be in internal logs, which are governed by the [privacy policy](https://proxy.golang.org/privacy).

​	代理和校验数据库协议只向远程服务器发送模块的路径和版本。如果您请求一个私有模块，镜像将尝试下载它，就像任何Go用户一样，并以同样的方式失败。关于失败的请求的信息不会在任何地方公布。请求的唯一痕迹将出现在内部日志中，这受[隐私政策](https://proxy.golang.org/privacy)的约束。

### Why did a previously available module become unavailable in the mirror?

为什么以前可用的模块在镜像中变得不可用？

[proxy.golang.org](https://proxy.golang.org/) does not save all modules forever. There are a number of reasons for this, but one reason is if [proxy.golang.org](https://proxy.golang.org/) is not able to detect a suitable license. In this case, only a temporarily cached copy of the module will be made available, and may become unavailable if it is removed from the original source and becomes outdated. The checksums will still remain in the checksum database regardless of whether or not they have become unavailable in the mirror.

​	[proxy.golang.org](https://proxy.golang.org/)不会永远保存所有模块。这有很多原因，但其中一个原因是如果[proxy.golang.org](https://proxy.golang.org/)无法检测到合适的许可证。在这种情况下，只有一个临时缓存的模块副本将被提供，如果它被从原始源头移除并变得过时，则可能变得不可用。无论镜像中是否变得不可用，校验和仍将保留在校验和数据库中。

### I have discovered a malicious module version in the mirror. Where do I report it?

我在镜像中发现了一个恶意的模块版本。我在哪里报告它？

Following [the security policy](https://golang.org/security#reporting), send an email to [security@golang.org](mailto:security@golang.org) with the word "vulnerability" in the message somewhere.

​	遵循[安全政策](https://golang.org/security#reporting)，发送电子邮件到 [security@golang.org](mailto:security@golang.org)，并在邮件的某处写上 "vulnerability "一词。