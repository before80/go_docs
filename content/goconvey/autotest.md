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

# Auto test

[Edit](https://github.com/smartystreets/goconvey/wiki/Auto-test/_edit) [New page](https://github.com/smartystreets/goconvey/wiki/_new)

gaku edited this page on Dec 20, 2014 · [7 revisions](https://github.com/smartystreets/goconvey/wiki/Auto-test/_history)

There are two ways to run tests automatically with GoConvey: in the terminal or in the browser.

### 1. In the terminal: auto-run.py



For viewing test reports in the terminal, use something like [auto-run.py](https://gist.github.com/mdwhatcott/9107649), which was at one time bundled with GoConvey. First download the script and put it somewhere on your path.

```
cd <folder_with_tests_or_packages>
auto-run.py -v
```



Run this from your project's top-level folder, and any changes to `.go` files under your current directory will trigger tests to run. The `-v` option enables "verbose" mode, and is entirely optional.

### 2. In the browser: `goconvey` server



When you use the [web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) to watch test results, tests are already run automatically.

### Performance



One of the advantages with method #2 (in the browser) is that tests are run across multiple goroutines to bring results back as fast as possible.
