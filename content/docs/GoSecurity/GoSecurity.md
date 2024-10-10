+++
title = "Go 的安全问题"
weight = 1
date = 2023-05-18T16:50:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go Security - Go的安全问题

> 原文：https://go.dev/security/

## Overview 概述

本页提供了关于使用Go编写安全可靠的软件的信息。

## Go Security Go的安全问题

### Go Security Policy  Go 安全政策

​	[Go Security Policy](../GoSecurityPolicy)解释了如何向 Go 团队报告 Go 标准库和子库的安全问题。

### Go Security Releases Go 安全发布

​	[Go Release History](../../../References/ReleaseHistory)包括过去安全问题的发布说明。根据[发布政策 （release policy）](../../../References/ReleaseHistory#release-policy)，我们对Go的`两个最新的主要版本`发布了安全补丁。

## Go Vulnerability Management Go 漏洞管理

​	[Go’s vulnerability management](./GoVulnerabilityManagement)支持帮助开发人员找到可能影响其 Go 项目的已知公共漏洞。

## Go Fuzzing

​	[Go native fuzzing](../../UsingAndUnderstandingGo/Fuzzing) 提供了一种自动化测试，它不断地操纵程序的输入以发现漏洞。

​	从 Go 1.18 开始，Go 在其标准工具链中支持模糊测试。[OSS-Fuzz](https://google.github.io/oss-fuzz/getting-started/new-project-guide/go-lang/#native-go-fuzzing-support)支持原生Go模糊测试。请尝试使用[Go的模糊测试教程](../../GettingStarted/TutorialGettingStartedWithFuzzing_)。

## Go Cryptography Go 密码学

​	Go 密码学库是 Go 标准库和子库中的 [crypto/…](https://pkg.go.dev/crypto) 和 [golang.org/x/crypto/…](https://pkg.go.dev/golang.org/x/crypto)包，并按照[这些原则](https://go.googlesource.com/proposal/+/master/design/cryptography-principles.md)开发。