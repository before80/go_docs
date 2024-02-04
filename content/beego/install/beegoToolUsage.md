+++
title = "bee 工具简介"
date = 2024-02-04T09:09:09+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/install/bee/]({{< ref "/beego/install/beegoToolUsage" >}})

# Introduction to bee tool - bee 工具简介



Bee tool is a project for rapid Beego development. With bee tool developers can create, auto compile and reload, develop, test, and deploy Beego applications quickly and easily.

​	Bee 工具是一个用于快速开发 Beego 的项目。借助 bee 工具，开发人员可以快速轻松地创建、自动编译和重新加载、开发、测试和部署 Beego 应用程序。

## Installing bee tool 安装 bee 工具

Install bee tool with the following command:

​	使用以下命令安装 bee 工具：

```
go get github.com/beego/bee/v2
```

Update the bee tool with the following command:

​	使用以下命令更新 bee 工具：

```
go get -u github.com/beego/bee/v2
```

`bee` is installed into `GOPATH/bin` by default. You need to add `GOPATH/bin` to your PATH, otherwise the `bee` command won’t work.

​	 `bee` 默认安装在 `GOPATH/bin` 中。您需要将 `GOPATH/bin` 添加到您的 PATH 中，否则 `bee` 命令将无法运行。

## bee tool commands bee 工具命令

Type `bee` in command line and the following messages with be displayed:

​	在命令行中键入 `bee` ，将显示以下消息：

```
bee is a tool for managing Beego framework.

Usage:

	bee command [arguments]

The commands are:

	new         Create a Beego application
	run         run the app and start a Web server for development
	pack        Compress a Beego project into a single file
	api         create an API Beego application
	bale        packs non-Go files to Go source files
	version     show the bee, Beego and Go version
	generate    source code generator
	migrate     run database migrations
```

### Command `new` 命令 `new`

The `new` command can create a new web project. You can create a new Beego project by typing `bee new <project name>` under `$GOPATH/src`. This will generate all the default project folders and files:

​	 `new` 命令可以创建一个新的 Web 项目。您可以在 `$GOPATH/src` 下键入 `bee new <project name>` 来创建一个新的 Beego 项目。这将生成所有默认项目文件夹和文件：

```
bee new myproject
[INFO] Creating application...
/gopath/src/myproject/
/gopath/src/myproject/conf/
/gopath/src/myproject/controllers/
/gopath/src/myproject/models/
/gopath/src/myproject/static/
/gopath/src/myproject/static/js/
/gopath/src/myproject/static/css/
/gopath/src/myproject/static/img/
/gopath/src/myproject/views/
/gopath/src/myproject/conf/app.conf
/gopath/src/myproject/controllers/default.go
/gopath/src/myproject/views/index.tpl
/gopath/src/myproject/main.go
13-11-25 09:50:39 [SUCC] New application successfully created!
myproject
├── conf
│   └── app.conf
├── controllers
│   └── default.go
├── main.go
├── models
├── routers
│   └── router.go
├── static
│   ├── css
│   ├── img
│   └── js
├── tests
│   └── default_test.go
└── views
    └── index.tpl

8 directories, 4 files
```

### Command `api` 命令 `api`

The `new` command is used for crafting new web applications. The `api` command is used to create new API applications. Here is the result of running `bee api project_name`:

​	 `new` 命令用于构建新的 Web 应用程序。 `api` 命令用于创建新的 API 应用程序。以下是运行 `bee api project_name` 的结果：

```
bee api apiproject
create app folder: /gopath/src/apiproject
create conf: /gopath/src/apiproject/conf
create controllers: /gopath/src/apiproject/controllers
create models: /gopath/src/apiproject/models
create tests: /gopath/src/apiproject/tests
create conf app.conf: /gopath/src/apiproject/conf/app.conf
create controllers default.go: /gopath/src/apiproject/controllers/default.go
create tests default.go: /gopath/src/apiproject/tests/default_test.go
create models object.go: /gopath/src/apiproject/models/object.go
create main.go: /gopath/src/apiproject/main.go
```

Below is the generated project structure of a new API application:

​	下面是新 API 应用程序的生成项目结构：

```
apiproject
├── conf
│   └── app.conf
├── controllers
│   └── object.go
│   └── user.go
├── docs
│   └── doc.go
├── main.go
├── models
│   └── object.go
│   └── user.go
├── routers
│   └── router.go
└── tests
    └── default_test.go
```

Compare this to the `bee new myproject` command seen earlier. Note that the new API application doesn’t have a `static` and `views` folder.

​	将此与前面看到的 `bee new myproject` 命令进行比较。请注意，新的 API 应用程序没有 `static` 和 `views` 文件夹。

You can also create a model and controller based on the database schema by providing database conn:

​	您还可以通过提供数据库连接来基于数据库架构创建模型和控制器：

```
bee api [appname] [-tables=""] [-driver=mysql] [-conn=root:@tcp(127.0.0.1:3306)/test]
```

### Command `run` 命令 `run`

The `bee run` command will supervise the file system of any Beego project using [inotify](http://en.wikipedia.org/wiki/Inotify). The results will autocompile and display immediately after any modification in the Beego project folders.

​	 `bee run` 命令将使用 inotify 监督任何 Beego 项目的文件系统。结果将在 Beego 项目文件夹中进行任何修改后立即自动编译并显示。

```
13-11-25 09:53:04 [INFO] Uses 'myproject' as 'appname'
13-11-25 09:53:04 [INFO] Initializing watcher...
13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject/controllers)
13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject/models)
13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject)
13-11-25 09:53:04 [INFO] Start building...
13-11-25 09:53:16 [SUCC] Build was successful
13-11-25 09:53:16 [INFO] Restarting myproject ...
13-11-25 09:53:16 [INFO] ./myproject is running...
```

Visting `http://localhost:8080/` with a web browser will display your app running:

​	使用 Web 浏览器访问 `http://localhost:8080/` 将显示正在运行的应用：

![img](./beegoToolUsage_img/beerun.png)

After modifying the `default.go` file in the `controllers` folder, the following output will be displayed in the command line:

​	修改 `controllers` 文件夹中的 `default.go` 文件后，命令行中将显示以下输出：

```
13-11-25 10:11:20 [EVEN] "/gopath/src/myproject/controllers/default.go": DELETE|MODIFY
13-11-25 10:11:20 [INFO] Start building...
13-11-25 10:11:20 [SKIP] "/gopath/src/myproject/controllers/default.go": CREATE
13-11-25 10:11:23 [SKIP] "/gopath/src/myproject/controllers/default.go": MODIFY
13-11-25 10:11:23 [SUCC] Build was successful
13-11-25 10:11:23 [INFO] Restarting myproject ...
13-11-25 10:11:23 [INFO] ./myproject is running...
```

Refresh the browser to show the results of the new modifications.

​	刷新浏览器以显示新修改的结果。

### Command `pack` 命令 `pack`

The `pack` command is used to compress the project into a single file. The compressed file can be deployed by uploading and extracting the zip file to the server.

​	 `pack` 命令用于将项目压缩成单个文件。可以通过将 zip 文件上传并解压到服务器来部署压缩文件。

```
bee pack
app path: /gopath/src/apiproject
GOOS darwin GOARCH amd64
build apiproject
build success
exclude prefix:
exclude suffix: .go:.DS_Store:.tmp
file write to `/gopath/src/apiproject/apiproject.tar.gz`
```

The compressed file will be in the project folder:

​	压缩文件将位于项目文件夹中：

```
rwxr-xr-x  1 astaxie  staff  8995376 11 25 22:46 apiproject
-rw-r--r--  1 astaxie  staff  2240288 11 25 22:58 apiproject.tar.gz
drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 conf
drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 controllers
-rw-r--r--  1 astaxie  staff      509 11 25 22:31 main.go
drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 models
drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 tests
```

### Command `bale` 命令 `bale`

This command is currently only available to the developer team. It is used to compress all static files in to a single binary file so that they do not need to carry static files including js, css, images and views when publishing the project. Those files will be self-extracting with non-overwrite when the program starts.

​	此命令目前仅对开发人员团队可用。它用于将所有静态文件压缩成单个二进制文件，以便在发布项目时无需携带包括 js、css、图像和视图在内的静态文件。这些文件将在程序启动时自动解压，且不会覆盖。

### Command `version` 命令 `version`

This command displays the version of `bee`, `beego`, and `go`.

​	此命令显示 `bee` 、 `beego` 和 `go` 的版本。 此命令尝试输出 beego 的版本。它适用于 GOPATH 模式。Bee 从 $GOPATH/src/astaxie/beego 目录中查找 beego 的版本。

```shell
$ bee version
bee   :1.2.2
Beego :1.4.2
Go    :go version go1.3.3 darwin/amd64
```

This command try to output beego’s version. It works well for GOPATH mode. Bee finds beego’s version from $GOPATH/src/astaxie/beego directory.

So when we use GOMOD mode, and we don’t download beego’s source code, Bee could not find the version’s information.

​	因此，当我们使用 GOMOD 模式，并且我们没有下载 beego 的源代码时，Bee 无法找到版本信息。

### Command `generate` 命令 `generate`

This command will generate the routers by analyzing the functions in controllers.

​	此命令将通过分析控制器中的函数来生成路由器。

```
bee generate scaffold [scaffoldname] [-fields=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    The generate scaffold command will do a number of things for you.
    -fields: a list of table fields. Format: field:type, ...
    -driver: [mysql | postgres | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
    example: bee generate scaffold post -fields="title:string,body:text"

bee generate model [modelname] [-fields=""]
    generate RESTful model based on fields
    -fields: a list of table fields. Format: field:type, ...

bee generate controller [controllerfile]
    generate RESTful controllers

bee generate view [viewpath]
    generate CRUD view in viewpath

bee generate migration [migrationfile] [-fields=""]
    generate migration file for making database schema update
    -fields: a list of table fields. Format: field:type, ...

bee generate docs
    generate swagger doc file
    
bee generate routers [-ctrlDir=/path/to/controller/directory] [-routersFile=/path/to/routers/file.go] [-routersPkg=myPackage]
    -ctrlDir: the directory contains controllers definition. Bee scans this directory and its subdirectory to generate routers info
    -routersFile: output file. All generated routers info will be output into this file. 
              If file not found, Bee create new one, or Bee truncates it.
              The default value is "routers/commentRouters.go"
    -routersPkg: package declaration.The default value is "routers". 
              When you pass routersFile parameter, you'd better pass this parameter

bee generate test [routerfile]
    generate testcase

bee generate appcode [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"] [-level=3]
    generate appcode based on an existing database
    -tables: a list of table names separated by ',', default is empty, indicating all tables
    -driver: [mysql | postgres | sqlite], the default is mysql
    -conn:   the connection string used by the driver.
             default for mysql:    root:@tcp(127.0.0.1:3306)/test
             default for postgres: postgres://postgres:postgres@127.0.0.1:5432/postgres
    -level:  [1 | 2 | 3], 1 = models; 2 = models,controllers; 3 = models,controllers,router
```

### Command `migrate` 命令 `migrate`

This command will run database migration scripts.

​	此命令将运行数据库迁移脚本。

```
bee migrate [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    run all outstanding migrations
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

bee migrate rollback [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    rollback the last migration operation
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

bee migrate reset [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    rollback all migrations
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

bee migrate refresh [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
    rollback all migrations and run them all again
    -driver: [mysql | postgresql | sqlite], the default is mysql
    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
```

## bee tool configuration bee 工具配置

The file `bee.json` in the bee tool source code folder is the Beego configuration file. This file is still under development, but some options are already available to use:

​	bee 工具源代码文件夹中的文件 `bee.json` 是 Beego 配置文件。此文件仍在开发中，但已有一些选项可供使用：

- `"version": 0`: version of file, for checking incompatible format version.
  `"version": 0` ：文件的版本，用于检查不兼容的格式版本。
- `"go_install": false`: if you use a full import path like `github.com/user/repo/subpkg` you can enable this option to run `go install` and speed up you build processes.
  `"go_install": false` ：如果您使用完整导入路径（如 `github.com/user/repo/subpkg` ），则可以启用此选项以运行 `go install` 并加快构建过程。
- `"watch_ext": []`: add other file extensions to watch (only watch `.go` files by default). For example, `.ini`, `.conf`, etc.
  `"watch_ext": []` ：添加其他要监视的文件扩展名（默认情况下仅监视 `.go` 文件）。例如， `.ini` 、 `.conf` 等。
- `"dir_structure":{}`: if your folder names are not the same as MVC classic names you can use this option to change them.
  `"dir_structure":{}` ：如果您的文件夹名称与 MVC 经典名称不同，可以使用此选项来更改它们。
- `"cmd_args": []`: add command arguments for every start.
  `"cmd_args": []` ：为每个启动添加命令参数。
- `"envs": []`: set environment variables for every start.
  `"envs": []` ：为每个启动设置环境变量。
