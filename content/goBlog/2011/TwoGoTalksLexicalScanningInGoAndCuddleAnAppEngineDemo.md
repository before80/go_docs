+++
title = "两场 Go 讲座： Go 中的词汇扫描 和 Cuddle：一个App Engine Demo"
weight = 13
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Two Go Talks: "Lexical Scanning in Go" and "Cuddle: an App Engine Demo" - 两场 Go 讲座： Go 中的词汇扫描 和 Cuddle：一个App Engine Demo

> 原文：[https://go.dev/blog/sydney-gtug](https://go.dev/blog/sydney-gtug)

Andrew Gerrand
1 September 2011

On Tuesday night Rob Pike and Andrew Gerrand each presented at the [Sydney Google Technology User Group](http://www.sydney-gtug.org/).

周二晚上，Rob Pike和Andrew Gerrand分别在悉尼谷歌技术用户组发表演讲。

Rob’s talk, "[Lexical Scanning in Go](http://www.youtube.com/watch?v=HxaD_trXwRE)", discusses the design of a particularly interesting and idiomatic piece of Go code, the lexer component of the new [template package.](https://go.dev/pkg/exp/template/)

Rob的演讲是 "Go中的词汇扫描"，讨论了Go代码中一个特别有趣和习惯的部分的设计，即新模板包的词汇器组件。

{{< youtube "HxaD_trXwRE">}}

The slides are [available here](http://cuddle.googlecode.com/hg/talk/lex.html). The new template package is available as [exp/template](https://go.dev/pkg/exp/template/) in Go release r59. In a future release it will replace the old template package.

幻灯片可以在这里找到。新的模板包在Go版本r59中作为exp/template提供。在未来的版本中，它将取代旧的模板包。

Andrew’s talk, "[Cuddle: an App Engine Demo](http://www.youtube.com/watch?v=HQtLRqqB-Kk)", describes the construction of a simple real-time chat application that uses App Engine’s [Datastore](http://code.google.com/appengine/docs/go/datastore/overview.html), [Channel](http://code.google.com/appengine/docs/go/channel/overview.html), and [Memcache](http://code.google.com/appengine/docs/go/datastore/memcache.html) APIs. It also includes a question and answer session that covers [Go for App Engine](http://code.google.com/appengine/docs/go/gettingstarted/) and Go more generally.

Andrew的演讲 "Cuddle: an App Engine Demo "描述了一个简单的实时聊天应用的构建，该应用使用了App Engine的Datastore、Channel和Memcache APIs。它还包括一个问答环节，涉及App Engine的Go和更广泛的Go。

{{< youtube "HQtLRqqB-Kk">}}

The slides are [available here](http://cuddle.googlecode.com/hg/talk/index.html). The code is available at the [cuddle Google Code project](http://code.google.com/p/cuddle/).

幻灯片可以在这里找到。代码可以在cuddle Google Code项目中找到。
