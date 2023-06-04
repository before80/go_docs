+++
title = "发布模块"
weight = 5
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Publishing a module - 发布模块

> 原文：[https://go.dev/doc/modules/publishing](https://go.dev/doc/modules/publishing)

When you want to make a module available for other developers, you publish it so that it’s visible to Go tools. Once you’ve published the module, developers importing its packages will be able to resolve a dependency on the module by running commands such as `go get`.

当您想把一个模块提供给其他开发者时，您可以发布它，这样它就能被Go工具看到。一旦您发布了模块，导入其软件包的开发者就可以通过运行go get等命令来解决对该模块的依赖。

> **Note:** Don’t change a tagged version of a module after publishing it. For developers using the module, Go tools authenticate a downloaded module against the first downloaded copy. If the two differ, Go tools will return a security error. Instead of changing the code for a previously published version, publish a new version.注意：不要在发布模块后改变模块的标记版本。对于使用该模块的开发者来说，Go工具会根据第一个下载的副本来验证下载的模块。如果两者不同，Go工具将返回一个安全错误。不要改变之前发布的版本的代码，而是发布一个新的版本。

**See also**

另见

- For an overview of module development, see [Developing and publishing modules](https://go.dev/doc/modules/developing)关于模块开发的概述，见开发和发布模块
- For a high-level module development workflow – which includes publishing – see [Module release and versioning workflow](https://go.dev/doc/modules/release-workflow).关于高层次的模块开发工作流程--包括发布--见模块发布和版本工作流程。

## Publishing steps 发布步骤

Use the following steps to publish a module.

使用下面的步骤来发布模块。

1. Open a command prompt and change to your module’s root directory in the local repository.打开命令提示符，切换到本地版本库中模块的根目录。

2. Run `go mod tidy`, which removes any dependencies the module might have accumulated that are no longer necessary.运行go mod tidy，删除模块可能积累的不再需要的任何依赖关系。

   ```
   $ go mod tidy
   ```

3. Run `go test ./...` a final time to make sure everything is working.最后一次运行go test ./...以确保一切正常。

   This runs the unit tests you’ve written to use the Go testing framework.这将运行您编写的单元测试，以使用Go测试框架。

   ```
   $ go test ./...
   ok      example.com/mymodule       0.015s
   ```

4. Tag the project with a new version number using the `git tag` command.使用git tag命令给项目贴上新的版本号。

   For the version number, use a number that signals to users the nature of changes in this release. For more, see [Module version numbering](https://go.dev/doc/modules/version-numbers).对于版本号，请使用一个能向用户表明此版本中变化性质的数字。更多信息请参见模块的版本号。

   ```
   $ git commit -m "mymodule: changes for v0.1.0"
   $ git tag v0.1.0
   ```

5. Push the new tag to the origin repository.推送新的标签到原点仓库。

   ```
   $ git push origin v0.1.0
   ```

6. Make the module available by running the [`go list` command](https://go.dev/cmd/go/#hdr-List_packages_or_modules) to prompt Go to update its index of modules with information about the module you’re publishing.通过运行 go list 命令使模块可用，以提示 Go 用您要发布的模块的信息更新其模块索引。

   Precede the command with a statement to set the `GOPROXY` environment variable to a Go proxy. This will ensure that your request reaches the proxy.在命令的前面加上一条语句，将GOPROXY环境变量设置为Go代理。这将确保您的请求到达代理。

   ```
   $ GOPROXY=proxy.golang.org go list -m example.com/mymodule@v0.1.0
   ```

Developers interested in your module import a package from it and run the [`go get` command](https://go.dev/doc/modules/publishing) just as they would with any other module. They can run the [`go get` command](https://go.dev/doc/modules/publishing) for latest versions or they can specify a particular version, as in the following example:

对您的模块感兴趣的开发者从您的模块中导入一个包，然后像对待其他模块一样运行 go get 命令。他们可以运行go get命令获取最新的版本，也可以指定一个特定的版本，就像下面的例子：

```
$ go get example.com/mymodule@v0.1.0
```