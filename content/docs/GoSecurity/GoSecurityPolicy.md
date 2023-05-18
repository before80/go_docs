+++
title = "Go安全策略"
weight = 3
date = 2023-05-18T16:50:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go Security Policy - Go安全策略

> 原文：[https://go.dev/security/policy](https://go.dev/security/policy)

## 概述

​	本文档介绍Go安全团队处理报告问题的流程以及期望得到的反馈。

## 报告安全漏洞

​	Go发布的所有安全漏洞都应通过电子邮件发送到[security@golang.org](mailto:security@golang.org)进行报告。此邮件将被送到Go安全团队。

​	为确保您的报告不被标记为垃圾邮件，请在电子邮件中任意位置包含单词"vulnerability（漏洞）"。请为报告电子邮件使用一个描述性的主题行。

​	您的电子邮件将在7天内得到确认，并且在解决问题之前您将被及时通知进展情况。您的问题将在90天内得到修复或公开。

​	如果您在7天内没有收到对您的电子邮件的回复，请再次通过[security@golang.org](mailto:security@golang.org)与Go安全团队联系。请确保在您的电子邮件中包含"vulnerability（漏洞）"这个词。

​	如果再过3天您仍未收到报告确认，请注意您的电子邮件可能已被标记为垃圾邮件。在这种情况下，请[在此处](https://g.co/vulnz)提交问题。选择"I want to report a technical security or an abuse risk related bug in a Google product (SQLi, XSS, etc.)"，并将"Go"列为受影响的产品。

## 跟踪

​	根据您的问题性质，Go安全团队将其分类为公开（PUBLIC）、私有（PRIVATE）或紧急（URGENT）跟踪。所有安全问题都将获得CVE编号。

### 公开（PUBLIC）

​	公开跟踪的问题影响小众配置，影响非常有限，或者已经广泛知晓。

​	公开跟踪的问题在**公开修复**，并被回溯到下一个计划的[小版本发布](https://go.dev/wiki/MinorReleases)中（约每月一次）。发布公告包括这些问题的详细信息，但没有预先公告。

​	过去公开问题的例子包括：

- [#44916](https://go.dev/issue/44916): archive/zip: 在调用Reader.Open时可能会出现panic

- [#44913](https://go.dev/issue/44913): encoding/xml: 使用自定义的TokenReader时，xml.NewTokenDecoder可能会出现无限循环

- [#43786](https://go.dev/issue/43786): crypto/elliptic: 在P-224曲线上执行不正确的操作

- [#40928](https://go.dev/issue/40928): net/http/cgi,net/http/fcgi: 在未指定Content-Type时存在跨站脚本（XSS）漏洞

- [#40618](https://go.dev/issue/40618): encoding/binary: ReadUvarint和ReadVarint可以从无效输入中读取无限字节数

- [#36834](https://go.dev/issue/36834): crypto/x509: 在Windows 10上存在证书验证绕过漏洞

  

### 私有（PRIVATE）

​	PRIVATE 跟踪的问题违反了已承诺的安全属性。

​	私有问题在下一个计划的[小版本发布](https://go.dev/wiki/MinorReleases)中修复，并保持私有状态。

​	在发布之前的三到七天，会向 golang-announce 发送预先公告，宣布即将发布的版本中存在一个或多个安全修复，并说明这些问题是否影响标准库、工具链或两者，以及每个修复所预留的 CVE ID。

​	过去 PRIVATE 问题的一些示例包括：

- [#42552](https://go.dev/issue/42552): math/big: 在递归除法处理极大数时出现 panic
- [#34902](https://go.dev/issue/34902): net/http: 在 httputil.ReverseProxy 中使用 Expect 100-continue 会导致 panic
- [#39360](https://go.dev/issue/39360): crypto/x509: Certificate.Verify 方法在 Windows 上似乎忽略了 EKU 要求
- [#34960](https://go.dev/issue/34960): crypto/dsa: 无效的公钥会导致在 dsa.Verify 中出现 panic
- [#34540](https://go.dev/issue/34540): net/http: 规范化无效标头，从而允许请求走私
- [#29098](https://go.dev/issue/29098): net/url: URL.Parse 存在多个解析问题
- [#53416](https://go.dev/issue/53416): path/filepath: 在 Glob 中存在栈耗尽
- [#53616](https://go.dev/issue/53616): go/parser: 在所有 Parse* 函数中存在栈耗尽
- [#54658](https://go.dev/issue/54658): net/http: 在发送 GOAWAY 后处理服务器错误
- [#56284](https://go.dev/issue/56284): syscall、os/exec: 环境变量中未经过消毒的 NUL

### 紧急（URGENT）

​	URGENT 跟踪的问题对 Go 生态系统的完整性构成威胁，或正在被积极利用，在野外导致严重损害。尽管没有最近的例子，但它们可能包括 net/http 中的远程代码执行或 crypto/tls 中的实际密钥恢复。

​	URGENT 跟踪问题在私有环境中修复，并触发立即的专用安全发布，可能没有预先公告。。

## 将现有问题标记为安全相关

​	如果您认为现有问题与安全有关，我们要求您发送电子邮件至 [security@golang.org](mailto:security@golang.org)。该电子邮件应包括问题 ID 和一个简短的描述，说明为什么应根据此安全策略进行处理。

## 披露流程

​	Go项目采用以下披露流程：

1. 一旦收到安全报告，会指定一位主要处理者。这个人会协调修复和发布流程。

2. 确认问题并确定受影响的软件列表。

3. 审查代码以查找任何潜在的类似问题。

4. 如果经过与提交者协商确定需要CVE号码，主要处理者将获取它。

5. 为最近的两个主要版本和当前版本准备修复程序。修复程序将准备好并合并到最近的两个主要版本和当前版本。

6. 在修正程序被应用的那一天，会向[golang-announce](https://groups.google.com/group/golang-announce)、[golang-dev](https://groups.google.com/group/golang-dev)和[golang-nuts](https://groups.google.com/group/golang-nuts)发送公告。

   

​	这个过程可能需要一些时间，尤其是需要与其他项目的维护者协调的情况下。我们将尽一切努力尽快处理错误，但是遵循上述流程非常重要，以确保披露得到一致的处理。

​	对于包括分配CVE号码的安全问题，该问题会在[CVEDetails网站上公开列出"Golang"产品](https://www.cvedetails.com/vulnerability-list/vendor_id-14185/Golang.html)以及[国家漏洞披露网站](https://web.nvd.nist.gov/view/vuln/search)。

## 接收安全更新

​	接收安全公告的最佳方法是订阅[golang-announce](https://groups.google.com/forum/#!forum/golang-announce)邮件列表。与安全问题相关的任何消息都将以"[security]"为前缀。

## 对这个政策的评论

​	如果您有任何建议改进此政策，请[提交一个议题](https://go.dev/issue/new)以供讨论。