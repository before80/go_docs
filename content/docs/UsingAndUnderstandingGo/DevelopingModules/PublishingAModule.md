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

​	当您想把一个模块提供给其他开发者时，您可以发布它，这样它就能被Go工具看到。一旦您发布了模块，导入其包的开发者就可以通过运行`go get`等命令来解决对该模块的依赖。

> 注意
>
> ​	不要在发布模块后改变模块的标记版本。对于使用该模块的开发者来说，Go工具会根据第一个下载的副本来验证下载的模块。如果两者不同，Go工具将返回一个安全错误。不要改变之前发布的版本的代码，而是发布一个新的版本。

另请参阅

- 关于模块开发的概述，请参阅[开发和发布模块](../DevelopingAndPublishingModules)
- 关于高层次的模块开发工作流程 —— 包括发布 —— 请参阅[模块发布和版本工作流程](../ModuleReleaseAndVersioningWorkflow)。

## 发布步骤

​	使用下面的步骤来发布模块。

1. 打开命令提示符，切换到本地存储库中模块的根目录。

2. 运行`go mod tidy`，删除模块可能积累的不再需要的任何依赖项。

   ```bash
   $ go mod tidy
   ```

3. 最后一次运行`go test ./...`以确保一切正常。

   这将运行为使用Go测试框架进行编写的单元测试。

   ```bash
   $ go test ./...
   ok      example.com/mymodule       0.015s
   ```

4. 使用`git tag`命令给该项目贴上新的版本号。

   对于版本号，请使用一个能向用户表明此版本中变化性质的数字。更多信息请参见[模块版本编号](../ModuleVersionNumbering)。

   ```bash
   $ git commit -m "mymodule: changes for v0.1.0"
   $ git tag v0.1.0
   ```

5. 推送新的标签到原始存储库。

   ```bash
   $ git push origin v0.1.0
   ```

6. 通过运行 [go list 命令]({{< ref "/cmd/go#go-list---列出包或模块">}})使该模块可用，以提示 Go 用您正在发布的模块的信息更新其模块索引。

   在该命令的前面加上一条语句，将`GOPROXY`环境变量设置为Go代理。这将确保您的请求到达该代理。

   ```bash
   $ GOPROXY=proxy.golang.org go list -m example.com/mymodule@v0.1.0
   ```

​	对您的模块感兴趣的开发者从您的模块中导入一个包，然后像对待其他模块一样运行 [go get 命令]({{< ref "/cmd/go#go-get---添加依赖项到当前模块并安装它们">}})。他们可以运行go get命令获取最新的版本，也可以指定一个特定的版本，就像下面的例子：

```bash
$ go get example.com/mymodule@v0.1.0
```