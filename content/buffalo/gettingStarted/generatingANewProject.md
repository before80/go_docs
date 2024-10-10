+++
title = "生成新项目"
date = 2024-02-04T21:05:50+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/getting_started/new-project/]({{< ref "/buffalo/gettingStarted/generatingANewProject" >}})

# Generating a New Project 生成新项目 

You now have a working Buffalo installation. In this section, you will learn how to create **a brand new web application**, using the `buffalo` command.

​	您现在拥有一个可用的 Buffalo 安装。在本节中，您将学习如何使用 `buffalo` 命令创建一个全新的 Web 应用程序。

## Create a New Project 创建新项目 

Buffalo aims to make building new web applications in Go as **quick and simple** as possible. What could be more simple than a *new application* generator?

​	Buffalo 旨在让在 Go 中构建新的 Web 应用程序尽可能快速和简单。还有什么比新的应用程序生成器更简单的呢？

Start by going to your preferred folder where you want to place your project, then:

​	首先转到您希望放置项目的首选文件夹，然后：

```bash
$ buffalo new coke
```

That will generate a whole new Buffalo application called **coke**, all ready to go:

​	这将生成一个全新的 Buffalo 应用程序，名为 coke，一切准备就绪：

- the **Buffalo framework layout** and default configuration ([pop/soda](https://github.com/gobuffalo/pop) with PostgreSQL support),
  Buffalo 框架布局和默认配置（pop/soda 支持 PostgreSQL），
- all necessary **Go dependencies** needed to run the current application,
  运行当前应用程序所需的所有必要的 Go 依赖项，
- **frontend dependencies** and working setup with [webpack](https://webpack.js.org/)
  前端依赖项和 webpack 的工作设置
- and an initial **Git repository**.
  以及一个初始 Git 存储库。

```bash
$ buffalo new coke
DEBU[2022-05-25T11:06:33-05:00] Step: 435aea40
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] Exec: go mod init coke
go: creating new go.mod: module coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/README.md
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/actions_test.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/app.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/home.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/home_test.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/actions/render.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/cmd/app/main.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/.codeclimate.yml
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/.env
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/fixtures/sample.toml
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/grifts/init.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/inflections.json
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/config/buffalo-app.toml
DEBU[2022-05-25T11:06:33-05:00] Step: 638bde0d
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/Dockerfile
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/.dockerignore
DEBU[2022-05-25T11:06:33-05:00] Step: 7065092d
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/grifts/db.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/models/models.go
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/models/models_test.go
DEBU[2022-05-25T11:06:33-05:00] Step: 916dfca0
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/database.yml
DEBU[2022-05-25T11:06:33-05:00] Step: ff1c6d38
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] File: /your/path/coke/.buffalo.dev.yml
DEBU[2022-05-25T11:06:33-05:00] Step: 103396c6
DEBU[2022-05-25T11:06:33-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:33-05:00] Exec: go install github.com/gobuffalo/buffalo-pop/v3@latest
DEBU[2022-05-25T11:06:34-05:00] Step: ae1260b1
DEBU[2022-05-25T11:06:34-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/config/buffalo-plugins.toml
DEBU[2022-05-25T11:06:34-05:00] Step: 4992ff46
DEBU[2022-05-25T11:06:34-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/actions/app.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/actions/home.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/actions/home_test.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/actions/render.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/locales/all.en-us.yaml
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/locales/embed.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/public/embed.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/public/robots.txt
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/templates/_flash.plush.html
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/templates/application.plush.html
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/templates/embed.go
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/templates/home/index.plush.html
DEBU[2022-05-25T11:06:34-05:00] Step: 69d71878
DEBU[2022-05-25T11:06:34-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:34-05:00] LookPath: yarn
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/css/_buffalo.scss
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/css/application.scss
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/images/favicon.ico
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/images/logo.svg
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/assets/js/application.js
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/.babelrc
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/package.json
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/postcss.config.js
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/public/assets/keep
DEBU[2022-05-25T11:06:34-05:00] File: /your/path/coke/webpack.config.js
DEBU[2022-05-25T11:06:34-05:00] LookPath: yarn
DEBU[2022-05-25T11:06:34-05:00] Exec: yarn --version
1.22.17
DEBU[2022-05-25T11:06:34-05:00] Exec: yarn set version berry
➤ YN0000: Retrieving https://repo.yarnpkg.com/3.2.1/packages/yarnpkg-cli/bin/yarn.js
➤ YN0000: Saving the new release in .yarn/releases/yarn-3.2.1.cjs
➤ YN0000: Done in 0s 637ms
DEBU[2022-05-25T11:06:37-05:00] Exec: yarn config set enableGlobalCache true
➤ YN0000: Successfully set enableGlobalCache to true
DEBU[2022-05-25T11:06:37-05:00] Exec: yarn config set logFilters --json [{"code":"YN0013","level":"discard"}]
➤ YN0000: Successfully set logFilters to [
  {
    code: 'YN0013',
    text: undefined,
    pattern: undefined,
    level: 'discard'
  }
]
DEBU[2022-05-25T11:06:37-05:00] Exec: yarn --version
3.2.1
DEBU[2022-05-25T11:06:38-05:00] Exec: yarn install
DEBU[2022-05-25T11:06:38-05:00] ➤ YN0000: ┌ Resolution step

DEBU[2022-05-25T11:06:39-05:00] ➤ YN0032: │ fsevents@npm:2.3.2: Implicit dependencies on node-gyp are discouraged

DEBU[2022-05-25T11:06:43-05:00] ➤ YN0000: └ Completed in 5s 401ms

DEBU[2022-05-25T11:06:43-05:00] ➤ YN0000: ┌ Fetch step

DEBU[2022-05-25T11:06:43-05:00] ➤ YN0000: └ Completed

DEBU[2022-05-25T11:06:43-05:00] ➤ YN0000: ┌ Link step

DEBU[2022-05-25T11:06:44-05:00] ➤ YN0000: │ ESM support for PnP uses the experimental loader API and is therefore experimental

DEBU[2022-05-25T11:06:44-05:00] ➤ YN0007: │ @fortawesome/fontawesome-free@npm:5.15.4 must be built because it never has been before or the last one failed

DEBU[2022-05-25T11:06:44-05:00] ➤ YN0000: └ Completed in 0s 906ms

DEBU[2022-05-25T11:06:44-05:00] ➤ YN0000: Done with warnings in 6s 462ms

DEBU[2022-05-25T11:06:44-05:00] Step: bb3c28ed
DEBU[2022-05-25T11:06:44-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:44-05:00] Exec: go mod tidy
go: finding module for package github.com/gobuffalo/buffalo
go: finding module for package github.com/gobuffalo/mw-paramlogger
go: finding module for package github.com/gobuffalo/buffalo-pop/v3/pop/popmw
go: finding module for package github.com/gobuffalo/mw-i18n/v2
go: finding module for package github.com/gobuffalo/buffalo/render
go: finding module for package github.com/gobuffalo/mw-forcessl
go: finding module for package github.com/gobuffalo/envy
go: finding module for package github.com/gobuffalo/mw-csrf
go: finding module for package github.com/unrolled/secure
go: finding module for package github.com/markbates/grift/grift
go: finding module for package github.com/gobuffalo/pop/v6
go: finding module for package github.com/gobuffalo/suite/v4
go: found github.com/gobuffalo/buffalo in github.com/gobuffalo/buffalo v0.18.7
go: found github.com/gobuffalo/buffalo-pop/v3/pop/popmw in github.com/gobuffalo/buffalo-pop/v3 v3.0.4
go: found github.com/gobuffalo/buffalo/render in github.com/gobuffalo/buffalo v0.18.7
go: found github.com/gobuffalo/envy in github.com/gobuffalo/envy v1.10.1
go: found github.com/gobuffalo/mw-csrf in github.com/gobuffalo/mw-csrf v1.0.0
go: found github.com/gobuffalo/mw-forcessl in github.com/gobuffalo/mw-forcessl v0.0.0-20220514125302-be60179938a4
go: found github.com/gobuffalo/mw-i18n/v2 in github.com/gobuffalo/mw-i18n/v2 v2.0.1
go: found github.com/gobuffalo/mw-paramlogger in github.com/gobuffalo/mw-paramlogger v1.0.0
go: found github.com/unrolled/secure in github.com/unrolled/secure v1.10.0
go: found github.com/markbates/grift/grift in github.com/markbates/grift v1.5.0
go: found github.com/gobuffalo/pop/v6 in github.com/gobuffalo/pop/v6 v6.0.4
go: found github.com/gobuffalo/suite/v4 in github.com/gobuffalo/suite/v4 v4.0.2
DEBU[2022-05-25T11:06:45-05:00] Exec: go mod download
DEBU[2022-05-25T11:06:46-05:00] Step: 218a906c
DEBU[2022-05-25T11:06:46-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:46-05:00] Step: a3cee09e
DEBU[2022-05-25T11:06:46-05:00] Chdir: /your/path/coke
DEBU[2022-05-25T11:06:46-05:00] File: /your/path/coke/.gitignore
DEBU[2022-05-25T11:06:46-05:00] Exec: git init
hint: Using 'master' as the name for the initial branch. This default branch name
hint: is subject to change. To configure the initial branch name to use in all
hint: of your new repositories, which will suppress this warning, call:
hint:
hint: 	git config --global init.defaultBranch <name>
hint:
hint: Names commonly chosen instead of 'master' are 'main', 'trunk' and
hint: 'development'. The just-created branch can be renamed via this command:
hint:
hint: 	git branch -m <name>
Initialized empty Git repository in /your/path/coke/.git/
DEBU[2022-05-25T11:06:46-05:00] Exec: git add .
DEBU[2022-05-25T11:06:46-05:00] Exec: git commit -q -m Initial Commit
INFO[2022-05-25T11:06:46-05:00] Congratulations! Your application, coke, has been successfully generated!
INFO[2022-05-25T11:06:46-05:00] You can find your new application at: /your/path/coke
INFO[2022-05-25T11:06:46-05:00] Please read the README.md file in your new application for next steps on running your application.
```

## Create a Customized App 创建自定义应用 

The default setup is great, but maybe it doesn’t fit you. Buffalo provides several options as flags for the `new` command.

​	默认设置很好，但可能不适合您。Buffalo 为 `new` 命令提供多个选项作为标志。

You can get the available flags list using the `help` command:

​	您可以使用 `help` 命令获取可用标志列表：

```bash
$ buffalo help new
Creates a new Buffalo application

Usage:
  buffalo new [name] [flags]

Flags:
      --api                  skip all front-end code and configure for an API server
      --ci-provider string   specify the type of ci file you would like buffalo to generate [none, travis, gitlab-ci, circleci] (default "none")
      --config string        config file (default is $HOME/.buffalo.yaml)
      --db-type string       specify the type of database you want to use [cockroach, mariadb, mysql, postgres, sqlite3] (default "postgres")
  -d, --dry-run              dry run
  -f, --force                delete and remake if the app already exists
  -h, --help                 help for new
      --module string        specify the root module (package) name. [defaults to 'automatic']
      --skip-config          skips using the config file
      --skip-docker          skips generating the Dockerfile
      --skip-pop             skips adding pop/soda to your app
      --skip-webpack         skips adding Webpack to your app
      --skip-yarn            use npm instead of yarn for frontend dependencies management
      --vcs string           specify the Version control system you would like to use [none, git, bzr] (default "git")
  -v, --verbose              verbosely print out the go get commands
```

You can choose to generate an API application, skipping the frontend stuff. Maybe you want to setup a CI to build your app on your favorite system? Or even use your own package to handle the database? Just use the flags!

​	您可以选择生成 API 应用程序，跳过前端内容。也许您想设置一个 CI 来在您喜欢的系统上构建您的应用？或者甚至使用您自己的软件包来处理数据库？只需使用标志即可！

## Override Default Config 覆盖默认配置 

By default `buffalo new` command will look for a configuration file at `$HOME/.buffalo.yml` and if it exists will try to load it. You can override the flags found in that file by passing the right ones in the command line or use the `--config` flag to specify a different YAML file. If the `--skip-config` flag is used `buffalo new` command will not load any config file and will use only the flags passed by the command line.

​	默认情况下， `buffalo new` 命令将在 `$HOME/.buffalo.yml` 处查找配置文件，如果存在，它将尝试加载它。您可以通过在命令行中传递正确的标志或使用 `--config` 标志指定不同的 YAML 文件来覆盖该文件中找到的标志。如果使用了 `--skip-config` 标志， `buffalo new` 命令将不会加载任何配置文件，并且只会使用命令行传递的标志。

An example of a `.buffalo.yml` config file can be:

​	 `.buffalo.yml` 配置文件的一个示例可以是：

```yaml
skip-yarn: true
db-type: postgres
bootstrap: 4
with-dep: true
```

## Running Your Application in Development 在开发中运行您的应用程序 

Before starting Buffalo for the first time, please head over to the [Database]({{< ref "/buffalo/database/gettingStartedWithPop" >}}) docs and read a little bit about setting up your databases.
在首次启动 Buffalo 之前，请前往数据库文档并阅读有关设置数据库的少量内容。

One of the downsides to Go development is the lack of code “reloading”. This means as you change your code **you need to manually stop** your application, rebuild it, then restart it. Buffalo finds this annoying, and wants to make life better for you.

​	Go 开发的一个缺点是缺少代码“重新加载”。这意味着在您更改代码时，您需要手动停止应用程序，重新构建它，然后重新启动它。Buffalo 发现这很烦人，并希望让您的生活更美好。

```bash
$ buffalo dev
```

The `dev` command will watch your `.go` and `.html` files and the [asset]({{< ref "/buffalo/frontend/assets" >}}) folder by default. It will **rebuild and restart your binary for you** automatically, so you don’t have to worry about such things.

​	 `dev` 命令将默认监视您的 `.go` 和 `.html` 文件以及资产文件夹。它将自动为您重新构建并重新启动您的二进制文件，因此您不必担心这些事情。

Just run the `buffalo dev` command and go to [localhost:3000/](http://localhost:3000/) to see all changes live!

​	只需运行 `buffalo dev` 命令并转到 localhost:3000/ 即可实时查看所有更改！

#### Run the dev server on a custom port 在自定义端口上运行开发服务器

Sometimes you will already have an app working on the 3000 port. You can configure the dev server port by providing the `PORT` environment variable:

​	有时您已经在 3000 端口上运行了一个应用程序。您可以通过提供 `PORT` 环境变量来配置开发服务器端口：

```bash
$ PORT=3001 buffalo dev
```

You can also take a look at the [Env Variables]({{< ref "/buffalo/gettingStarted/configuration" >}}) chapter for further information on Buffalo configuration.

​	您还可以查看环境变量章节以获取有关 Buffalo 配置的更多信息。

## Next Steps 后续步骤 

- [Directory Structure]({{< ref "/buffalo/gettingStarted/directoryStructure" >}}) - Learn more about Buffalo structure.
  目录结构 - 了解有关 Buffalo 结构的更多信息。
