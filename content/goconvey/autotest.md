+++
title = "Auto test"
date = 2024-12-15T11:17:59+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Auto-test](https://github.com/smartystreets/goconvey/wiki/Auto-test)
>
> 收录该文档时间： `2024-12-15T11:17:59+08:00`

# Auto test - 自动测试



gaku edited this page on Dec 20, 2014 · [7 revisions](https://github.com/smartystreets/goconvey/wiki/Auto-test/_history)

​	gaku 于 2014 年 12 月 20 日编辑了此页面 · [7 次修订](https://github.com/smartystreets/goconvey/wiki/Auto-test/_history)

There are two ways to run tests automatically with GoConvey: in the terminal or in the browser.

​	使用 GoConvey 有两种方式可以自动运行测试：在终端中或在浏览器中。

### 1. 在终端中：使用 auto-run.py - In the terminal: auto-run.py



For viewing test reports in the terminal, use something like [auto-run.py](https://gist.github.com/mdwhatcott/9107649), which was at one time bundled with GoConvey. First download the script and put it somewhere on your path.

​	如果希望在终端中查看测试报告，可以使用类似 [auto-run.py](https://gist.github.com/mdwhatcott/9107649) 的脚本，该脚本曾经是 GoConvey 的一部分。首先下载该脚本并将其放在您的路径中。

```sh
cd <folder_with_tests_or_packages>
auto-run.py -v
```



Run this from your project's top-level folder, and any changes to `.go` files under your current directory will trigger tests to run. The `-v` option enables "verbose" mode, and is entirely optional.

​	从项目的顶级文件夹运行此命令，当前目录下的 `.go` 文件发生任何更改时都会触发测试运行。`-v` 选项用于启用“详细”模式，完全是可选的。

### 2. 在浏览器中：通过 `goconvey` 服务器 In the browser: `goconvey` server



When you use the [web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) to watch test results, tests are already run automatically.

​	当您使用 [Web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) 查看测试结果时，测试会自动运行。

### 性能 Performance



One of the advantages with method #2 (in the browser) is that tests are run across multiple goroutines to bring results back as fast as possible.

​	使用方法 #2（在浏览器中）的优势之一是测试会在多个 goroutine 中运行，以尽可能快地返回结果。
