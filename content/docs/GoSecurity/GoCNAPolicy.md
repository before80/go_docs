+++
title = "Go CNA策略"
weight = 3
date = 2023-05-18T16:50:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go CNA Policy - Go CNA策略

[https://go.dev/security/vuln/cna](https://go.dev/security/vuln/cna)

## 概述

​	Go CNA是[CVE编号机构](https://www.cve.org/ProgramOrganization/CNAs)，为Go生态系统中的公共漏洞发行[CVE IDs](https://www.cve.org/ResourcesSupport/Glossary?activeTerm=glossaryCVEID)并发布[CVE记录](https://www.cve.org/ResourcesSupport/Glossary?activeTerm=glossaryRecord)。它是Google CNA的子CNA。

## 范围

​	Go CNA涵盖Go项目（Go[标准库](https://go.dev/pkg)和[子存储库](https://pkg.go.dev/golang.org/x)）中的漏洞以及不受其他CNA覆盖的可导入Go模块中的公共漏洞。

​	该范围旨在明确排除在Go中编写但不可导入的应用程序或包的漏洞（例如，任何位于main包中的内容）。有关排除报告的更多信息，请参见[go.dev/security/vuln/database#excluded-reports](https://go.dev/security/vuln/database#excluded-reports)。

​	要报告Go项目中的潜在新漏洞，请参阅[go.dev/security/policy](https://go.dev/security/policy)。

## 为公共漏洞请求CVE ID

​	**重要提示**：以下链接的表单在问题跟踪器上创建公共问题，因此不得用于报告未披露的Go漏洞（请参见我们的[安全策略](../GoSecurityPolicy)，了解报告未披露问题的说明）。

​	如果要为现有公共漏洞在Go生态系统中请求CVE ID，[请通过此表单提交请求](https://go.dev/s/vulndb-report-new)。

​	如果已经公开披露了漏洞，或者该漏洞存在于您维护的软件包中，并且您已准备好公开披露，请将其视为公共漏洞。