+++
title = "Beego 介绍"
date = 2024-02-04T09:40:48+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/intro/introduction/]({{< ref "/beego/intro/beegoIntro" >}})

# Beego introduction - Beego 介绍



## What is Beego? Beego 是什么？

Beego is a RESTful HTTP framework for the rapid development of Go applications including APIs, web apps and backend services with integrated Go specific features such as interfaces and struct embedding.

​	Beego 是一个 RESTful HTTP 框架，用于快速开发 Go 语言的应用，包括 API、Web 应用和后端服务，它集成了 Go 语言的特色功能，例如接口和结构体嵌入。

## The architecture of Beego Beego 的架构

Beego is built upon 8 loosely linked modules that can be used independently or as part of Beego’s HTTP logic. This high level of modularity gives Beego an unprecedented level of flexibility to meet developer needs.

​	Beego 构建在 8 个松散链接的模块之上，这些模块可以独立使用，也可以作为 Beego 的 HTTP 逻辑的一部分使用。这种高度的模块化使 Beego 具有前所未有的灵活性，可以满足开发人员的需求。

![img](./beegoIntro_img/architecture.png)

## The execution logic of Beego Beego 的执行逻辑

Beego uses a standard Model-View-Controller (MVC) architecture for logic execution.

​	Beego 使用标准的模型-视图-控制器 (MVC) 架构进行逻辑执行。

![img](./beegoIntro_img/flow.png)

## The project structure of Beego Beego 的项目结构

Here is the typical folder structure of a Beego project:

​	以下是 Beego 项目的典型文件夹结构：

```
├── conf
│   └── app.conf
├── controllers
│   ├── admin
│   └── default.go
├── main.go
├── models
│   └── models.go
├── static
│   ├── css
│   ├── ico
│   ├── img
│   └── js
└── views
    ├── admin
    └── index.tpl
```

M (models), V (views), C (controllers) each have top level folders. `main.go` is the entry point.

​	M（模型）、V（视图）、C（控制器）各有一个顶级文件夹。 `main.go` 是入口点。

## Creating a Beego project 创建 Beego 项目

Ready to try Beego? You can use the [bee tool to create a new project]({{< ref "/beego/install/beegoToolUsage" >}}).

​	准备尝试 Beego 了吗？您可以使用 bee 工具创建一个新项目。
