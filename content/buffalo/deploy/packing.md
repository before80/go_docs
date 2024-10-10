+++
title = "打包"
date = 2024-02-04T21:19:27+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/deploy/packing/]({{< ref "/buffalo/deploy/packing" >}})

# Packing 打包 

Now, your project is ready to be deployed. In this section, you will learn how to package a version of your app to deploy it on a server.

​	现在，您的项目已准备好部署。在本节中，您将学习如何打包应用程序的版本以将其部署到服务器上。

## The build Command 构建命令 

Buffalo features a command, `build`, that will build a **full binary** of your application including, but not limited to; assets, migrations, templates, etc. If you buy into the “Buffalo Way”, things just work. It’s a wonderful experience. :)

​	Buffalo 提供了一个命令 `build` ，它将构建应用程序的完整二进制文件，包括但不限于资产、迁移、模板等。如果您采用“Buffalo Way”，一切都会正常工作。这是一次美妙的体验。 :)

```bash
$ buffalo build
Buffalo version v0.18.14

--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/coke -ldflags -X main.version=b5dffda -X main.buildTime="2017-03-20T11:05:23-04:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

When the build finishes, you have a fresh baked binary in the `bin` folder. It will also have the **compilation time** and the **git commit SHA** burnt in, thus making the binaries “versioned”.

​	构建完成后，您将在 `bin` 文件夹中获得一个新鲜烘焙的二进制文件。它还将具有编译时间和 git 提交 SHA，从而使二进制文件“版本化”。

## Customize the Build 自定义构建 

To get the list of available options, use the help command:

​	要获取可用选项的列表，请使用帮助命令：

```bash
$ buffalo help build
Build the application binary, including bundling of webpack assets

Usage:
  buffalo build [flags]

Aliases:
  build, b, bill, install

Flags:
    --build-flags strings        Additional comma-separated build flags to feed to go build
    --clean-assets               will delete public/assets before calling webpack
    --dry-run                    runs the build 'dry'
    --environment string         set the environment for the binary (default "development")
-e, --extract-assets             extract the assets and put them in a distinct archive
-h, --help                       help for build
    --ldflags string             set any ldflags to be passed to the go build
    --mod string                 -mod flag for go build
-o, --output string              set the name of the binary
-k, --skip-assets                skip running webpack and building assets
    --skip-build-deps            skip building dependencies
    --skip-template-validation   skip validating templates
-s, --static                     build a static binary using  --ldflags '-linkmode external -extldflags "-static"'
-t, --tags string                compile with specific build tags
-v, --verbose                    print debugging information
```

### Binary name / location 二进制名称/位置 

By default, your application will be built in the `bin` directory of your project, and the name of the executable will be the name you used to create the project with the `new` command.

​	默认情况下，您的应用程序将在项目的 `bin` 目录中构建，可执行文件的名称将是您使用 `new` 命令创建项目的名称。

You can change this default name by using the `-o` or `-output` flag:

​	您可以使用 `-o` 或 `-output` 标志更改此默认名称：

```bash
$ buffalo build -o bin/cookies
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o bin/cookies -ldflags -X main.version="2017-04-02T08:32:28+02:00" -X main.buildTime="2017-04-02T08:32:28+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

In fact, you can change the target directory too:

​	事实上，您也可以更改目标目录：

```bash
$ # Put the app in my home directory, as "coke"
$ buffalo build -o ~/coke
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> packing .../coke/actions/actions-packr.go
--> running go build -v -o ~/coke -ldflags -X main.version="2017-04-02T08:32:28+02:00" -X main.buildTime="2017-04-02T08:32:28+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

### Extract Assets in a Zip File 将资产提取到一个 Zip 文件中 

By default, your whole app is packed into a single executable, assets included. In production setups, you may want to serve these assets with a proxy server (like Apache or NGINX), to lower the app load. You may even use a *CDN* to handle your assets.

​	默认情况下，您的整个应用程序都打包到一个可执行文件中，包括资产。在生产环境中，您可能希望使用代理服务器（如 Apache 或 NGINX）来提供这些资产，以降低应用程序负载。您甚至可以使用 CDN 来处理您的资产。

Buffalo provides a way to extract compiled app assets into a single archive, using the `-e` or `-extract-assets` flag:

​	Buffalo 提供了一种使用 `-e` 或 `-extract-assets` 标志将已编译的应用程序资产提取到单个存档中的方法：

```bash
$ buffalo build -e
--> cleaning up target dir
--> running node_modules/.bin/webpack
--> build assets archive
--> disable self assets handling
--> running go build -v -o bin/coke -ldflags -X main.version="2017-04-02T08:45:58+02:00" -X main.buildTime="2017-04-02T08:45:58+02:00"
--> cleaning up build
----> cleaning up buffalo_build_main.go
----> cleaning up a
----> cleaning up a/a.go
----> cleaning up a/database.go
----> cleaning up buffalo_build_main.go
----> cleaning up ...coke/actions/actions-packr.go
```

By default, the assets archive is put in the *bin* directory, but if you change the executable output directory with the `-o` flag, the assets will be put in the same directory.

​	默认情况下，资产存档放在 bin 目录中，但如果您使用 `-o` 标志更改可执行输出目录，则资产将放在同一目录中。

```bash
$ ls -la bin
total 36280
drwxr-xr--@  4 markbates  staff   136B Apr  3 10:10 ./
drwxr-xr-x@ 20 markbates  staff   680B Apr  3 10:10 ../
-rwxr-xr-x@  1 markbates  staff    17M Apr  3 10:10 coke*
-rw-r--r--@  1 markbates  staff   691K Apr  3 10:10 coke-assets.zip
```

## Advanced Options 高级选项 

### Building “Static”/CGO Binaries 构建“静态”/CGO 二进制文件 

Building statically linked binaries that contain CGO, think SQLite3, can be tricky. By using the `--static` flag with `buffalo build`, the flags `--ldflags '-linkmode external -extldflags "-static"'` will be added to the `go build` command.

​	构建包含 CGO（例如 SQLite3）的静态链接二进制文件可能很棘手。通过将 `--static` 标志与 `buffalo build` 一起使用，标志 `--ldflags '-linkmode external -extldflags "-static"'` 将被添加到 `go build` 命令中。

### Build Tags 构建标记 

When building a Buffalo binary using the `buffalo build` command, you can pass `--tags` and `--ldflags` to the built binary; just as you normally would when using the `go build` tools.

​	使用 `buffalo build` 命令构建 Buffalo 二进制文件时，可以将 `--tags` 和 `--ldflags` 传递给构建的二进制文件；就像在使用 `go build` 工具时一样。

```bash
$ buffalo build --tags="mytag" --ldflags="-X foo.Bar=baz"
```

## Binary Commands 二进制命令 

### Modes 模式 

Binaries, by default, run in `development` mode, which means all of the sub-commands will run in that mode as well. To change the mode, you must use the `GO_ENV` environment variable.

​	二进制文件默认在 `development` 模式下运行，这意味着所有子命令也将以该模式运行。要更改模式，必须使用 `GO_ENV` 环境变量。

```bash
$ GO_ENV=production ./coke
```

### Available commands 可用命令 

Once a binary has been built, there are several sub-commands that can be run on that binary:

​	构建二进制文件后，可以对该二进制文件运行多个子命令：

#### Default 默认 

The default command, if you just run the binary, will start the application.

​	默认命令（如果您只运行二进制文件）将启动应用程序。

#### migrate

The `migrate` sub-command will run the migrations for the application.

​	 `migrate` 子命令将运行应用程序的迁移。

#### version 版本 

The `version` sub-command will output the version information for the binary, including the name, the git commit SHA used to build the binary, and the time the binary was built.

​	 `version` 子命令将输出二进制文件的版本信息，包括名称、用于构建二进制文件的 git 提交 SHA 以及构建二进制文件的时间。

```bash
$ ./coke version
coke version 69b6a8b ("2017-04-03T10:19:46-04:00")
```

#### task 任务 

The `task` sub-command runs tasks.

​	 `task` 子命令运行任务。

```bash
$ ./coke task greet

Hello World!
```

## Next Steps 后续步骤 

- [Using a Proxy]({{< ref "/buffalo/deploy/usingAProxy" >}}) - Integrate your app with a server like NGINX.
  使用代理 - 将您的应用与 NGINX 等服务器集成。
- [Systemd Service]({{< ref "/buffalo/deploy/systemdService" >}}) - Run your app as a systemd service.
  Systemd 服务 - 将您的应用作为 systemd 服务运行。
- [Cloud Providers]({{< ref "/buffalo/deploy/cloudProviders" >}}) - Deploy your app on a cloud provider.
  云提供商 - 在云提供商上部署您的应用。
