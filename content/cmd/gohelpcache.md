+++
title = "go help cache"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 510
description = ""
isCJKLanguage = true
draft = false

+++

​	

The go command caches build outputs for reuse in future builds. The default location for cache data is a subdirectory named go-build in the standard user cache directory for the current operating system. Setting the GOCACHE environment variable overrides this default,and running 'go env GOCACHE' prints the current cache directory.

​	go 命令会对构建输出进行缓存，以便在将来的构建中重用。缓存数据的默认位置是当前操作系统的标准用户缓存目录中的名为 go-build 的子目录。设置 GOCACHE 环境变量将覆盖此默认值，并且运行 'go env GOCACHE' 将打印当前的缓存目录。

The go command periodically deletes cached data that has not been used recently. Running 'go clean -cache' deletes all cached data.

​	go 命令会定期删除长时间未使用的缓存数据。运行 'go clean -cache' 将删除所有缓存数据。

The build cache correctly accounts for changes to Go source files,compilers, compiler options, and so on: cleaning the cache explicitly should not be necessary in typical use. However, the build cache does not detect changes to C libraries imported with cgo. If you have made changes to the C libraries on your system, you will need to clean the cache explicitly or else use the -a build flag (see 'go help build') to force rebuilding of packages that depend on the updated C libraries.

​	构建缓存会正确处理 Go 源文件、编译器、编译器选项等的更改：在典型用法中，不应该明确清理缓存。但是，构建缓存不会检测导入 cgo 的 C 库的更改。如果您对系统上的 C 库进行了更改，则需要明确清理缓存，否则使用 -a 构建标志（参见 'go help build'）来强制重新构建依赖于已更新的 C 库的包。

The go command also caches successful package test results. See 'go help test' for details. Running 'go clean -testcache' removes all cached test results (but not cached build results).

​	go 命令还会缓存成功的包测试结果。有关详细信息，请参阅 'go help test'。运行 'go clean -testcache' 将删除所有缓存的测试结果（但不会删除缓存的构建结果）。

The go command also caches values used in fuzzing with 'go test -fuzz', specifically, values that expanded code coverage when passed to a fuzz function. These values are not used for regular building and testing, but they're stored in a subdirectory of the build cache.Running 'go clean -fuzzcache' removes all cached fuzzing values.This may make fuzzing less effective, temporarily.

​	go 命令还会缓存在 'go test -fuzz' 中使用的模糊测试值，具体而言，这些值在传递给模糊函数时扩展了代码覆盖范围。这些值不用于常规构建和测试，但它们存储在构建缓存的子目录中。运行 'go clean -fuzzcache' 将删除所有缓存的模糊测试值。这可能会使模糊测试在短时间内变得不那么有效。

The GODEBUG environment variable can enable printing of debugging information about the state of the cache:

​	GODEBUG 环境变量可以启用有关缓存状态的调试信息：

- GODEBUG=gocacheverify=1 causes the go command to bypass the use of any cache entries and instead rebuild everything and check that the results match existing cache entries.
- GODEBUG=gocacheverify=1 导致 go 命令绕过使用任何缓存条目，而是重新构建一切并检查结果是否与现有缓存条目匹配。


- GODEBUG=gocachehash=1 causes the go command to print the inputs for all of the content hashes it uses to construct cache lookup keys.
- GODEBUG=gocachehash=1 导致 go 命令打印用于构建缓存查找键的所有内容哈希的输入。

The output is voluminous but can be useful for debugging the cache.

​	输出庞大，但在调试缓存时可能很有用。

GODEBUG=gocachetest=1 causes the go command to print details of its decisions about whether to reuse a cached test result.

​	GODEBUG=gocachetest=1 导致 go 命令打印关于是否重用缓存的测试结果的详细信息。
