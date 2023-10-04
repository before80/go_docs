+++
title = "Govulncheck v1.0.0 发布！"
date = 2023-08-21T14:54:43+08:00
weight = 93
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Govulncheck v1.0.0 is released! - Govulncheck v1.0.0 发布！

> 原文：[https://go.dev/blog/govulncheck](https://go.dev/blog/govulncheck)
>

Julie Qiu, for the Go security team
13 July 2023

Julie Qiu，Go 安全团队 2023 年 7 月 13 日

We are excited to announce that govulncheck v1.0.0 has been released, along with v1.0.0 of the API for integrating scanning into other tools!

​	我们很高兴地宣布，govulncheck v1.0.0 已经发布，同时还发布了将扫描集成到其他工具中的 API 的 v1.0.0 版本！

Go’s support for vulnerability management was [first announced](https://go.dev/blog/vuln) last September. We have made several changes since then, culminating in today’s release.

​	Go对漏洞管理的支持最初是在去年9月份[首次宣布的](https://go.dev/blog/vuln)。自那时以来，我们进行了几次更改，最终在今天发布了新版本。

This post describes Go’s updated vulnerability tooling, and how to get started using it. We also recently published a [security best practices guide](https://go.dev/security/best-practices) to help you prioritize security in your Go projects.

​	这篇文章介绍了Go更新后的漏洞工具，以及如何开始使用它。我们还最近发布了一个[安全最佳实践指南](https://go.dev/security/best-practices)，以帮助您在Go项目中优先考虑安全性。

## Govulncheck

[Govulncheck](https://golang.org/x/vuln/cmd/govulncheck) is a command-line tool that helps Go users find known vulnerabilities in their project dependencies. The tool can analyze both codebases and binaries, and it reduces noise by prioritizing vulnerabilities in functions that your code is actually calling.

​	[Govulncheck](https://golang.org/x/vuln/cmd/govulncheck) 是一个命令行工具，可以帮助Go用户在项目依赖中查找已知的漏洞。该工具可以分析代码库和二进制文件，通过优先考虑您的代码实际调用的函数中的漏洞，从而减少噪音。

You can install the latest version of govulncheck using [go install](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies):

​	您可以使用[go install](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)安装最新版本的 govulncheck：

```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
```

Then, run govulncheck inside your module:

​	然后，在您的模块内运行 govulncheck：

```bash
govulncheck ./...
```

See the [govulncheck tutorial](https://go.dev/doc/tutorial/govulncheck) for additional information on how to get started with using the tool.

​	有关如何开始使用该工具的更多信息，请参阅 [govulncheck 教程](https://go.dev/doc/tutorial/govulncheck)。

As of this release, there is now a stable API available, which is described at [golang.org/x/vuln/scan](https://golang.org/x/vuln/scan). This API provides the same functionality as the govulncheck command, enabling developers to integrate security scanners and other tools with govulncheck. As an example, see the [osv-scanner integration with govulncheck](https://github.com/google/osv-scanner/blob/d93d6b73e90ae392fe2b1b64a33bda6976b65b2d/internal/sourceanalysis/go.go#L20).

​	在这个版本中，现在有一个稳定的 API 可供使用，该API 的详细信息在[golang.org/x/vuln/scan](https://golang.org/x/vuln/scan)中有说明。这个API提供了与 govulncheck 命令相同的功能，使开发人员能够将安全扫描程序和其他工具与 govulncheck 集成。例如，查看 [osv-scanner 集成 govulncheck](https://github.com/google/osv-scanner/blob/d93d6b73e90ae392fe2b1b64a33bda6976b65b2d/internal/sourceanalysis/go.go#L20)。

## 数据库 Database

Govulncheck is powered by the Go vulnerability database, [https://vuln.go.dev](https://vuln.go.dev/), which provides a comprehensive source of information about known vulnerabilities in public Go modules. You can browse the entries in the database at [pkg.go.dev/vuln](https://pkg.go.dev/vuln).

​	Govulncheck 的数据来自 Go 漏洞数据库，[https://vuln.go.dev](https://vuln.go.dev/)，它为已知的公共Go模块中的漏洞提供了全面的信息来源。您可以在数据库中浏览条目，网址是 [pkg.go.dev/vuln](https://pkg.go.dev/vuln)。

Since the initial release, we have updated the [database API](https://go.dev/security/vuln/database#api) to improve performance and ensure long-term extensibility. An experimental tool to generate your own vulnerability database index is provided at [golang.org/x/vulndb/cmd/indexdb](https://golang.org/x/vulndb/cmd/indexdb).

​	自首次发布以来，我们已经更新了[数据库 API](https://go.dev/security/vuln/database#api)，以提高性能并确保长期可扩展性。一个用于生成自己的漏洞数据库索引的实验性工具位于[golang.org/x/vulndb/cmd/indexdb](https://golang.org/x/vulndb/cmd/indexdb)。

If you are a Go package maintainer, we encourage you to [contribute information](https://go.dev/s/vulndb-report-new) about public vulnerabilities in your projects.

​	如果您是Go包的维护者，我们鼓励您[提交信息](https://go.dev/s/vulndb-report-new)，关于您的项目中的公共漏洞。

For more information about the Go vulnerability database, see [go.dev/security/vuln/database](https://go.dev/security/vuln/database).

​	有关Go漏洞数据库的更多信息，请参见 [go.dev/security/vuln/database](https://go.dev/security/vuln/database)。

## 集成 Integrations

Vulnerability detection is now integrated into a suite of tools that are already part of many Go developers’ workflows.

​	漏洞检测现在已经集成到许多Go开发者工作流程中的一套工具中。

Data from the Go vulnerability database can be browsed at [pkg.go.dev/vuln](https://pkg.go.dev/vuln). Vulnerability information is also surfaced on the search and package pages of pkg.go.dev. For example, [the versions page of golang.org/x/text/language](https://pkg.go.dev/golang.org/x/text/language?tab=versions) shows vulnerabilities in older versions of the module.

​	Go漏洞数据库的数据可以在[pkg.go.dev/vuln](https://pkg.go.dev/vuln)中浏览。漏洞信息也显示在 pkg.go.dev 的搜索和包页面上。例如，[golang.org/x/text/language 的版本页面](https://pkg.go.dev/golang.org/x/text/language?tab=versions)显示了旧版本模块中的漏洞。

You can also run govulncheck directly in your editor using the Go extension for Visual Studio Code. See [the tutorial](https://go.dev/doc/tutorial/govulncheck-ide) to get started.

​	您还可以使用 Visual Studio Code 的 Go 扩展直接在编辑器中运行 govulncheck。请参阅[教程](https://go.dev/doc/tutorial/govulncheck-ide)以开始使用。

Lastly, we know that many developers will want to run govulncheck as part of their CI/CD systems. As a starting point, we have provided a [GitHub Action for govulncheck](https://github.com/marketplace/actions/golang-govulncheck-action) for integration with your projects.

​	最后，我们知道许多开发者希望将 govulncheck 作为他们的 CI/CD 系统的一部分运行。作为起点，我们为 govulncheck 提供了一个[GitHub Action for govulncheck](https://github.com/marketplace/actions/golang-govulncheck-action)，用于与您的项目集成。

## 视频演示 Video Walkthrough

If you are interested in a demo of the integrations described above, we presented a walkthrough of these tools at Google I/O this year, in our talk, [Build more secure apps with Go and Google](https://www.youtube.com/watch?v=HSt6FhsPT8c&ab_channel=TheGoProgrammingLanguage).

​	如果您对上述集成的演示感兴趣，我们今年在 Google I/O 上演示了这些工具，您可以在我们的演讲中观看，标题是 [Build more secure apps with Go and Google](https://www.youtube.com/watch?v=HSt6FhsPT8c&ab_channel=TheGoProgrammingLanguage)。

## 反馈 Feedback

As always, we welcome your feedback! See details on [how to contribute and help us make improvements](https://go.dev/security/vuln/#feedback).

​	和往常一样，我们欢迎您的反馈！有关[如何贡献并帮助我们进行改进的详细信息](https://go.dev/security/vuln/#feedback)。

We hope you’ll find the latest release of Go’s support for vulnerability management useful and work with us to build a more secure and reliable Go ecosystem.

​	我们希望您会发现Go对漏洞管理的最新版本对您有用，并与我们一起构建一个更安全可靠的Go生态系统。