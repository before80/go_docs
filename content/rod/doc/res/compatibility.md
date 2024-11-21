+++
title = "兼容性"
date = 2024-11-21T08:12:24+08:00
weight = 50
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/compatibility](https://go-rod.github.io/i18n/zh-CN/#/compatibility)
>
> 收录该文档时间： `2024-11-21T08:12:24+08:00`

# 兼容性

## 操作系统

​	一般来说你可以在 Golang 支持的所有主要平台上无感知地编译和运行 Rod。 推荐使用 [docker 方式](https://go-rod.github.io/i18n/zh-CN/#/custom-launch?id=remotely-manage-the-launcher) 在服务器上运行 Rod 。 在某些平台上，您可能需要手动安装浏览器。 Rod 无法保证自动下载的浏览器能够正常工作。 如果您想要Rod 支持一个平台，请为此提出问题。

​	在网上可以很轻松地搜索到如何在你的系统中安装浏览器。 比如说，对于 Ubuntu 或 Debian，可以搜索到这种安装浏览器的方法：

```bash
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
apt install ./google-chrome-stable_current_amd64.deb
```

​	对于 CentOS：

```bash
wget https://dl.google.com/linux/direct/google-chrome-stable_current_x86_64.rpm
yum localinstall -y google-chrome-stable_current_x86_64.rpm
```

​	对于 Alpine：

```bash
apk add chromium
```

## 支持的浏览器

​	Rod 支持任何使用 [DevTools 协议](https://chromedevtools.github.io/devtools-protocol/)的浏览器。

- 支持 Microsoft Edge。
- Firefox 目前正在[支持](https://wiki.mozilla.org/Remote)这一协议。
- Safari 目前还没有支持它的计划。
- IE 不会支持它的。

## 浏览器和 cdp 协议版本号

​	cdp 协议总是与 [launcher.DefaultRevision](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#DefaultRevision) 相同。 如果 Rod 找不到本地浏览器，它会下载 `launcher.DefaultRevision` 版本的浏览器。

每个版本的 Rod 只保证支持它的 `launcher.DefaultRevision` 版本的浏览器。

## API 版本号

​	采用 [Semver](https://semver.org/)。

​	在 `v1.0.0` 之前，如果版本号的第二个部分改变了，比如说由 `v0.1.0` 变为了 `v0.2.0`，那么肯定有公有 API 发生了改变，比如说函数名或参数类型发生了变更。 如果仅仅是版本号的最后一部分改变了，则公有 API 不会变更。

​	你可以使用 Github 的版本比较来查看自动生成的更新日志，例如，[比较 v0.75.2 与 v0.76.0](https://github.com/go-rod/rod/compare/v0.75.2...v0.76.0)。

## API 文档版本

​	参考这里 。

## Doc 网站版本

​	我们使用 github 来管理文档，很容易查看任何版本的文档：

1. 克隆 doc [repo](https://github.com/go-rod/go-rod.github.io.git)
2. Git checkout 到你想要的 Rod 版本发布日期附近的 commit
3. 安装 [docsify-cli](https://docsify.js.org/#/quickstart)。
4. 在仓库的根目录下运行 `docsify serve -o`
