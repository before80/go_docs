+++
title = "模块缓存"
date = 2023-05-17T09:59:21+08:00
weight = 13
description = ""
isCJKLanguage = true
draft = false
+++
## Module cache 模块缓存

> 原文：[https://go.dev/ref/mod#module-cache](https://go.dev/ref/mod#module-cache)

​	模块缓存是`go`命令存储下载模块文件的目录。模块缓存与构建缓存不同，构建缓存包含已编译的包和其他构建构件。

​	模块缓存的默认位置是 `$GOPATH/pkg/mod`。要使用其他位置，请设置`GOMODCACHE`[environment variable（环境变量）](../EnvironmentVariables)。

​	模块缓存没有最大大小，`go`命令不会自动删除其内容。

​	缓存可以由同一台机器上开发的多个 Go 项目共享。无论主模块的位置如何，`go`命令都会使用相同的缓存。`go`命令的多个实例可以同时安全地访问同一个模块缓存。

​	`go`命令在缓存中创建具有只读权限的模块源文件和目录，以防止模块下载后被意外更改。这有一个不幸的副作用，就是难以用`rm -rf`等命令删除缓存。取而代之的是，缓存可以用[go clean -modcache](../gomodFiles#go-clean-modcache)来删除。另外，当使用`-modcacherw`标志时，`go`命令将创建具有读写权限的新目录。这增加了编辑、测试和其他程序修改模块缓存中文件的风险。[go mod verify]() 命令可以用来检测对主模块的依赖项的修改。它扫描每个模块依赖项的提取内容，并确认它们与`go.sum`中预期的散列匹配。

​	下表解释了模块缓存中大多数文件的用途。省略了一些临时文件(锁定文件、临时目录)。对于每个路径，`$module`是一个模块路径，`$version`是一个版本。以斜线（`/`）结尾的路径是目录。模块路径和版本中的大写字母用感叹号转义（`Azure`被转义为`!azure`）以避免在不区分大小写的文件系统上发生冲突。



| Path | Description |
| ---- | ----------- |
|      |             |

### `$module@$version/`

​	包含模块`.zip`文件的提取内容的目录。它用作已下载模块的模块根目录。如果原始模块没有`go.mod`文件，那么它将不包含`go.mod`文件。



### `cache/download/`

​	包含从模块代理下载的文件和从[版本控制系统](../VersionControlSystems)获得的文件的目录。这个目录的布局遵循[GOPROXY 协议](../ModuleProxies#goproxy-protocol)，因此当由HTTP文件服务器提供服务或用`file://`URL引用此目录时，此目录可以作为一个代理使用。



### `cache/download/$module/@v/list`

​	已知版本的列表（见[GOPROXY 协议](../ModuleProxies#goproxy-protocol)）。这可能会随着时间的推移而改变，因此`go`命令通常会获取新的副本，而不是重新使用这个文件。



### `cache/download/$module/@v/$version.info`

​	有关版本的JSON元数据。(见[GOPROXY 协议](../ModuleProxies#goproxy-protocol))。这可能会随着时间的推移而改变，因此`go`命令通常会获取新的副本，而不是重复使用这个文件。



### `cache/download/$module/@v/$version.mod`

​	这个版本的`go.mod`文件（见[GOPROXY 协议](../ModuleProxies#goproxy-protocol)）。如果原来的模块没有`go.mod`文件，则这就是一个没有需求的合成文件。



### `cache/download/$module/@v/$version.zip`

​	模块的压缩内容（见[GOPROXY 协议](../ModuleProxies#goproxy-protocol)和[模块压缩文件](../ModuleZipFiles)）。



### `cache/download/$module/@v/$version.ziphash`

​	`.zip`文件中的文件的加密哈希值。请注意`.zip`文件本身没有被哈希化，因此文件的顺序、压缩、对齐和元数据不会影响哈希值。当使用一个模块时，`go`命令会验证这个哈希值是否与`go.sum`中的对应行匹配。[go mod verify](../gomodFiles#go-mod-verify)命令检查模块`.zip`文件和提取的目录的哈希值是否与这些文件相符。



### `cache/download/sumdb/`

​	包含从[checksum database（校验和数据库）](../AuthenticatingModules#checksum-database)（通常是`sum.golang.org`）下载的文件的目录。



### `cache/vcs/`

​	包含直接从源获取的模块的克隆的版本控制存储库。目录名称是由存储库类型和 URL 衍生出来的十六进制编码的哈希值。存储库在磁盘上的大小是经过优化的。例如，克隆的Git存储库在可能的情况下是裸露和浅层的。
