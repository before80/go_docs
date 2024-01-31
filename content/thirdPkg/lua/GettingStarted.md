+++
title = "开始入门"
date = 2024-01-25T17:30:37+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：

## Welcome! 欢迎！

Lua is a powerful and fast programming language that is easy to learn and use and to embed into your application.

​	Lua 是一种功能强大且快速的编程语言，易于学习和使用，并且可以嵌入到您的应用程序中。

Lua is designed to be a lightweight embeddable scripting language. It is used for [all sorts of applications](https://www.lua.org/uses.html), from games to web applications and image processing.

​	Lua 被设计为一种轻量级的可嵌入脚本语言。它用于各种应用程序，从游戏到 Web 应用程序和图像处理。

See the [about](https://www.lua.org/about.html) page for details and some reasons why you should choose Lua.

​	请参阅关于页面以了解详细信息以及您应该选择 Lua 的一些原因。

See what Lua programs look and feel like in the [live demo](https://www.lua.org/demo.html).

​	在[live demo](https://www.lua.org/demo.html)中了解 Lua 程序的外观和感觉。



## Installing 安装

Use the [live demo](https://www.lua.org/demo.html) to play with Lua if you don't want to install anything on your computer.

​	如果您不想在计算机上安装任何东西，可以使用实时演示来玩 Lua。

To run Lua programs on your computer, you'll need a standalone Lua interpreter and perhaps some additional Lua libraries. Pre-compiled Lua libraries and executables are available at [LuaBinaries](http://luabinaries.sourceforge.net/). Use your favorite text editor to write your Lua programs. Make sure to save your programs as plain text. If you want an IDE, try [ZeroBrane Studio](https://studio.zerobrane.com/).

​	要在计算机上运行 Lua 程序，您需要一个独立的 Lua 解释器，可能还需要一些额外的 Lua 库。[LuaBinaries](http://luabinaries.sourceforge.net/)上提供了预编译的 Lua 库和可执行文件。使用您最喜欢的文本编辑器来编写 Lua 程序。确保将程序另存为纯文本。如果您想要一个 IDE，请尝试 [ZeroBrane Studio](https://studio.zerobrane.com/)。

If you use Linux or Mac OS X, Lua is either already installed on your system or there is a Lua package for it. Make sure you get the latest release of Lua (currently 5.4.6).

​	如果您使用 Linux 或 Mac OS X，则系统上可能已经安装了 Lua，或者有适用于它的 Lua 软件包。确保您获得最新版本的 Lua（目前为 5.4.6）。

Lua is also quite easy to build from source, as explained below.

​	Lua 也非常容易从源代码构建，如下所述。

## Learning 学习

[![buy from Amazon](./GettingStarted_img/pil4.jpg)](https://www.amazon.com/exec/obidos/ASIN/8590379868/lua-start-20)

A good place to start learning Lua is the book [Programming in Lua](https://www.lua.org/pil/), available in [paperback](https://www.amazon.com/exec/obidos/ASIN/8590379868/lua-start-20) and as an [e-book](https://store.feistyduck.com/products/programming-in-lua-fourth-edition-ebook). The first edition is freely available [online](https://www.lua.org/pil/contents.html). See also [course notes](http://www.dcc.ufrj.br/~fabiom/lua/) based on this book.

​	开始学习 Lua 的一个好地方是《用 Lua 编程》一书，该书以平装本和电子书的形式提供。第一版可在网上免费获得。另请参阅基于本书的课程笔记。

The official definition of the Lua language is given in the [reference manual](https://www.lua.org/manual/5.4/).

​	Lua 语言的官方定义在参考手册中给出。

See the [documentation](https://www.lua.org/docs.html) page and the [wiki](http://lua-users.org/wiki/LuaDirectory) for more.

​	有关更多信息，请参阅文档页面和 wiki。



### Building from source 从源代码构建

Lua is very easy to build and install. Just [download](https://www.lua.org/download.html) it and follow the [instructions](https://www.lua.org/manual/5.4/readme.html) in the package.

​	Lua 非常容易构建和安装。只需下载它并按照软件包中的说明进行操作即可。

Here is a simple terminal session that downloads the current release of Lua and builds it in a Linux system:

​	这是一个简单的终端会话，它下载了 Lua 的当前版本并在 Linux 系统中构建它：

```
curl -L -R -O https://www.lua.org/ftp/lua-5.4.6.tar.gz
tar zxf lua-5.4.6.tar.gz
cd lua-5.4.6
make all test
```

If you don't have curl, try wget.
如果您没有 curl，请尝试 wget。

If you use Windows and want to build Lua from source, there are [detailed instructions](http://lua-users.org/wiki/BuildingLuaInWindowsForNewbies) in the [wiki](http://lua-users.org/wiki/).

​	如果您使用 Windows 并想从源代码构建 Lua，wiki 中有详细的说明。

## Getting help 获取帮助

Our [community](https://www.lua.org/community.html) is friendly and will most probably help you if you need. Just visit the [mailing list](https://www.lua.org/lua-l.html), the [chat room](http://lua-users.org/wiki/IrcChannel), and [stackoverflow](http://stackoverflow.com/questions/tagged/lua).

​	我们的社区非常友好，如果您需要帮助，他们很可能会帮助您。只需访问邮件列表、聊天室和 Stackoverflow。

If you need help in Portuguese, join the [Lua BR](https://www.lua.org/lua-br.html) mailing list and visit [pt.stackoverflow](http://pt.stackoverflow.com/questions/tagged/lua).

​	如果您需要葡萄牙语帮助，请加入 Lua BR 邮件列表并访问 pt.stackoverflow。

See also the [FAQ](https://www.lua.org/faq.html), the community-maintained [wiki](http://lua-users.org/wiki/) and [LuaFaq](http://lua-users.org/wiki/LuaFaq), and the much longer [uFAQ](http://www.luafaq.org/).

​	另请参阅常见问题解答、社区维护的 wiki 和 LuaFaq，以及更长的 uFAQ。



### Embedding 嵌入

To embed Lua into your C or C++ program, you'll need the Lua headers to compile your program and a Lua library to link with it. If you're getting a ready-made Lua package for your platform, you'll probably need the development package as well. Otherwise, just [download](https://www.lua.org/download.html) Lua and add its source directory to your project.

​	要将 Lua 嵌入到您的 C 或 C++ 程序中，您需要 Lua 头文件来编译程序和一个 Lua 库来链接它。如果您正在为您的平台获取现成的 Lua 包，您可能还需要开发包。否则，只需下载 Lua 并将其源目录添加到您的项目中。

## Tools 工具

If you need to complement the standard Lua libraries to handle more complex tasks, visit [LuaRocks](http://luarocks.org/), the main repository of Lua modules. See also [Awesome Lua](https://github.com/LewisJEllis/awesome-lua), a curated list of quality Lua packages and resources. The [lua-users wiki](http://lua-users.org/wiki/) lists many [user-contributed addons](http://lua-users.org/wiki/LuaAddons) for Lua.

​	如果您需要补充标准 Lua 库以处理更复杂的任务，请访问 LuaRocks，这是 Lua 模块的主要存储库。另请参阅 Awesome Lua，这是一个精选的优质 Lua 包和资源列表。lua-users wiki 列出了许多用户为 Lua 贡献的插件。

## Supporting 支持

You can help to [support the Lua project](https://www.lua.org/donations.html) by [buying a book](https://www.lua.org/donations.html#books) published by Lua.org and by [making a donation](https://www.lua.org/donations.html#donation).

​	您可以通过购买 Lua.org 出版的书籍和捐款来帮助支持 Lua 项目。

You can also help to spread the word about Lua by buying Lua products at [Zazzle](http://www.zazzle.com/Lua_Store).

​	您还可以在 Zazzle 购买 Lua 产品来帮助宣传 Lua。







