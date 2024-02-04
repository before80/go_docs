+++
title = "目录结构"
date = 2024-02-04T21:06:17+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/getting_started/directory-structure/](https://gobuffalo.io/documentation/getting_started/directory-structure/)

# Directory Structure 目录结构 

Buffalo provides you **a minimal directory structure** to work on your project. This structure **keeps the project clean**, and allow the generators to work. Don’t try to reinvent the wheel, and **let Buffalo buy you more time** to do the interesting part of your app! :)

​	Buffalo 为您提供了最小的目录结构来处理您的项目。此结构保持项目干净，并允许生成器工作。不要试图重新发明轮子，让 Buffalo 为您争取更多时间来完成应用程序的有趣部分！ :)

Now that you have a minimal new project, let’s go through its contents.

​	现在您已经有一个最小的新项目，让我们浏览一下它的内容。

## The Root Directory 根目录 

Here is the structure of a Buffalo project:

​	以下是 Buffalo 项目的结构：

```erb
├── .yarn/
├── actions/
│	├── app.go
│	└── render.go
├── assets/
├── cmd/
│	└── app/
│		└── main.go
├── config/
├── fixtures/
├── grifts/
├── locales/
├── models/
├── public/
├── templates/
├── .babelrc
├── .buffalo.dev.yml
├── .codeclimate.yml
├── .docketignore
├── .env
├── .gitignore
├── .pnp.loader.mjs
├── .yarnrc.yml
├── database.yml
├── Dockerfile
├── go.mod
├── go.sum
├── inflections.json
├── package.json
├── postcss.config.js
├── README.md
├── webpack.config.js
└── yarn.lock
```

### actions

This directory handles the **Controller** part of the MVC pattern. It contains the handlers for your URLs, plus:

​	此目录处理 MVC 模式的控制器部分。它包含 URL 的处理程序，以及：

- The `app.go` file to setup your app & routes,
  `app.go` 文件，用于设置您的应用程序和路由，
- The `render.go` file to setup the template engine(s).
  `render.go` 文件，用于设置模板引擎。

### assets

This directory is optional. If you don’t need to use a frontend setup (API only, for instance), it can be removed.
此目录是可选的。如果您不需要使用前端设置（例如，仅限 API），则可以将其删除。

This directory contains **raw** assets which will be compiled/compressed & put in the [`public`](https://gobuffalo.io/documentation/getting_started/directory-structure/#public) directory.

​	此目录包含将被编译/压缩并放入 `public` 目录的原始资产。

### cmd

This folder contains the `main.go` file which bootstraps your app and starts it.

​	此文件夹包含引导您的应用程序并启动它的 `main.go` 文件。

### grifts

This directory is optional. If you don’t need to use [tasks](https://gobuffalo.io/documentation/guides/tasks), you can remove it.
此目录是可选的。如果您不需要使用任务，则可以将其删除。

This directory contains the [tasks](https://gobuffalo.io/documentation/guides/tasks) powered by [grift](https://github.com/gobuffalo/grift).

​	此目录包含由 grift 支持的任务。

### locales

This directory is optional. If you use only one language, you can remove it and the i18n module from the `app.go` file in the `actions` directory.
此目录是可选的。如果您只使用一种语言，则可以将其和 `actions` 目录中的 `app.go` 文件中的 i18n 模块删除。

This directory is used by the i18n system. It will fetch the translation strings from here.

​	此目录由 i18n 系统使用。它将从此处获取翻译字符串。

### models

If you use pop/soda with the integrated generator, it will generate the model files here.
如果您将 pop/soda 与集成生成器一起使用，它将在此处生成模型文件。

This directory is optional. If you don’t need to use a database, it can be removed.
此目录是可选的。如果您不需要使用数据库，则可以将其删除。

This directory handles the **Model** part of the MVC pattern. It contains the `models.go` file to initialize the datasource connection, and the model files to reflect objects from the datasource.

​	此目录处理 MVC 模式的模型部分。它包含用于初始化数据源连接的 `models.go` 文件，以及用于反映数据源对象的模型文件。

### public

The contents of this directory are auto-generated.
此目录的内容是自动生成的。

This directory contains the public (compiled/compressed) assets. If you use webpack, it will put its assets in this directory.

​	此目录包含公共（已编译/压缩）资源。如果您使用 webpack，它会将资源放入此目录。

### templates

This directory is optional. If you don’t need to use a frontend setup (API only, for instance), it can be removed.
此目录是可选的。如果您不需要使用前端设置（例如，仅限 API），则可以将其删除。

This directory handles the **View** part of the MVC pattern. It contains the project templates, used to render the views.

​	此目录处理 MVC 模式的视图部分。它包含用于呈现视图的项目模板。

### tmp

The contents of this directory are auto-generated.
此目录的内容是自动生成的。

This directory is used by the `buffalo dev` command to rebuild your project on every change. The temporary files that Buffalo works with are put here.

​	此目录由 `buffalo dev` 命令使用，以便在每次更改时重建您的项目。Buffalo 使用的临时文件放在这里。

### database.yml

This file is optional. If you don’t need a database, or if you want to handle the database without pop/soda, you can remove it.
此文件是可选的。如果您不需要数据库，或者您想在不使用 pop/soda 的情况下处理数据库，则可以将其删除。

This file contains the database configuration for [pop/soda](https://github.com/gobuffalo/pop).

​	此文件包含 pop/soda 的数据库配置。

## Next Steps 后续步骤 

- [Configuration](https://gobuffalo.io/documentation/getting_started/configuration) - Manage your app configuration.
  配置 - 管理您的应用配置。
- [Tooling Integration](https://gobuffalo.io/documentation/getting_started/integrations) - Work with Buffalo, using existing tools.
  工具集成 - 使用现有工具与 Buffalo 协同工作。