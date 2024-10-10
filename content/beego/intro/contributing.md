+++
title = "贡献"
date = 2024-02-04T09:35:40+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/intro/contributing/]({{< ref "/beego/intro/contributing" >}})

# Contributing 贡献



## Contributing to Beego 对 Beego 的贡献

## Introduction 简介

Beego is free and open source software, which means that anyone can contribute to its development and progress under the Apache 2.0 License (http://www.apache.org/licenses/LICENSE-2.0.html). Beego’s source code is hosted on github (https://github.com/beego/beego).

​	Beego 是免费且开源的软件，这意味着任何人都可以在 Apache 2.0 许可证（http://www.apache.org/licenses/LICENSE-2.0.html）下为其开发和进步做出贡献。Beego 的源代码托管在 github（https://github.com/beego/beego）上。

### How can I become a contributor of Beego? 我如何成为 Beego 的贡献者？

You can fork, modify and then send a Pull Request to us. We will review your code and give you feedback on your changes as soon as possible.

​	您可以 fork、修改，然后向我们发送 Pull Request。我们将审查您的代码，并尽快就您的更改向您提供反馈。

## Pull Requests Pull Request

The process for pull requests for new features and bug fixes are not the same.

​	新功能和错误修复的 pull request 的流程不同。

### Bug fixes 错误修复

Pull requests for bug fixes do not need to create an issue first. If you have a solution to a bug, please describe your solution in detail in your pull request.

​	错误修复的 pull request 无需首先创建问题。如果您有错误的解决方案，请在您的 pull request 中详细描述您的解决方案。

### Documentation improvements 文档改进

You can help improve the documentation by submitting a pull request to the [beedoc](https://github.com/beego/beedoc) repository.

​	您可以通过向 beedoc 存储库提交 pull request 来帮助改进文档。

### New features proposals 新功能提案

Before you submit a pull request for a new feature, you should first create an issue with `[Proposal]` in the title, describing the new feature, as well as the implementation approach.

​	在您提交新功能的请求前，您应该首先创建一个标题中带有 `[Proposal]` 的问题，描述新功能以及实现方法。

Proposals will be reviewed and discussed by the core contributors, and can be adopted or potentially rejected.

​	核心贡献者将对其进行审查和商讨，并可能采纳或否决。

Once a proposal is accepted, create an implementation of the new features and submit it as a pull request. If the guidelines are not followed the pull request will be rejected immediately.

​	提案被接受后，创建新功能的实现并将其作为请求提交。如果不遵循指南，请求将立即被否决。

Since Beego follows the [Git Flow](http://nvie.com/posts/a-successful-git-branching-model/) branching model, ongoing development happens in the `develop` branch. Therefore, please base your pull requests on the HEAD of the `develop` branch.

​	由于 Beego 遵循 Git Flow 分支模型，正在进行的开发发生在 `develop` 分支中。因此，请将您的请求建立在 `develop` 分支的 HEAD 上。

### The git branches of Beego Beego 的 git 分支

The master branch is relatively stable and the dev branch is for developers. Here is a sample figure to show you how our branches work:

​	主分支是比较稳定的，而 dev 分支是给开发人员的。这里有一个简单的图例，向您展示我们的分支如何工作：

![img](./contributing_img/git-branch-1.png)

For more information about the branching model: http://nvie.com/posts/a-successful-git-branching-model/

​	有关分支模型的更多信息：http://nvie.com/posts/a-successful-git-branching-model/

## A simple guideline for Git command Git 命令的简单指南

You must have a github account, if not, please register one.

​	您必须有一个 GitHub 帐户，如果没有，请注册一个。

### Fork 代码

1. Click [https://github.com/beego/beego/v2](https://github.com/beego/beego)
   点击 https://github.com/beego/beego/v2
2. Click “Fork” button which is on top right corner
   点击右上角的“Fork”按钮

### Clone 代码

We recommend using official repo as `origin` repo, and then add a remote upstream to your repo.

​	我们建议使用官方仓库作为 `origin` 仓库，然后将远程上游添加到您的仓库。

If you already set SSH key, we recommend use SSH. The difference is that, we don’t need to input the username and password to push changes.

​	如果您已经设置了 SSH 密钥，我们建议使用 SSH。不同之处在于，我们无需输入用户名和密码即可推送更改。

Using SSH：

​	使用 SSH：

```bash
git clone git@github.com:astaxie/beego.git
cd beego
git remote add upstream 'git@github.com:<your github username>/beego.git'
```

Using HTTPS：

​	使用 HTTPS：

```bash
git clone https://github.com/beego/beego/v2.git
cd beego
git remote add  'https://github.com/<you github username>/beego.git'
```

The word `upstream` in command could be replaced with any word you like.

​	命令中的单词 `upstream` 可以替换为您喜欢的任何单词。

### fetch changes 获取更改

Every time you want to something, you’d better fetch remote changes:

​	每次您想执行操作时，最好获取远程更改：

```bash
git fetch
```

In this command, git only fetch `origin` repo。

​	在此命令中，git 仅获取 `origin` 仓库。

If we want to fetch our remote repo changes:

​	如果我们想获取远程仓库的更改：

```bash
git fetch upstream
```

You can replace `upstream` with your repo name

​	您可以用 `upstream` 替换您的仓库名称

### create feature branch 创建功能分支

我们在创建新的 feature 分支的时候，要先考虑清楚，从哪个分支切出来。 Before creating feature branch, we should think about choosing a branch as base branch.

Assume that we want to merge the new feature to develop branch. In such case:

​	假设我们要将新功能合并到开发分支。在这种情况下：

```bash
git checkout -b feature/my-feature origin/develop
```

Don’t forget to run `git fetch` before you create feature branch.

​	在创建功能分支之前，别忘了运行 `git fetch` 。

### push commit 推送提交

```bash
git add .
git commit
git push upstream my-feature
```

### make PR 创建 PR

Go to https://github.com/beego/beego, and make a Pull request

​	转到 https://github.com/beego/beego，并创建一个 Pull 请求
