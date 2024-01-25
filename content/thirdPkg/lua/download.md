+++
title = "下载"
date = 2024-01-25T21:53:44+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://www.lua.org/download.html](https://www.lua.org/download.html)

# Download 下载

## Source 源代码

Lua is [free software](https://www.lua.org/license.html) distributed in [source code](https://www.lua.org/ftp/). It may be used for any purpose, including commercial purposes, at absolutely no cost.

​	Lua 是以[源代码](https://www.lua.org/ftp/)形式分发的[免费软件](https://www.lua.org/license.html)。它可以用于任何目的，包括商业目的，绝对免费。

All [versions](https://www.lua.org/versions.html) are available for [download](https://www.lua.org/ftp/). The current version is [Lua 5.4](https://www.lua.org/versions.html#5.4) and its current release is [Lua 5.4.6](https://www.lua.org/ftp/lua-5.4.6.tar.gz).

​	所有[版本](https://www.lua.org/versions.html)均可[下载](https://www.lua.org/ftp/)。当前版本是 [Lua 5.4](https://www.lua.org/versions.html#5.4)，其当前版本是 Lua 5.4.6。

[lua-5.4.6.tar.gz](https://www.lua.org/ftp/lua-5.4.6.tar.gz)
2023-05-02, 355K, sha256:
7d5ea1b9cb6aa0b59ca3dde1c6adcb57ef83a1ba8e5432c0ecd06bf439b3ad88



## Tools 工具

The main repository of Lua modules is [LuaRocks](https://luarocks.org/). See also [Awesome Lua](https://github.com/LewisJEllis/awesome-lua). Pre-compiled Lua libraries and executables are available at [LuaBinaries](https://luabinaries.sourceforge.net/). The [lua-users wiki](http://lua-users.org/wiki/) lists many [user-contributed addons](http://lua-users.org/wiki/LuaAddons) for Lua.

​	Lua 模块的主要存储库是 [LuaRocks](https://luarocks.org/)。另请参阅 [Awesome Lua](https://github.com/LewisJEllis/awesome-lua)。预编译的 Lua 库和可执行文件可在 [LuaBinaries](https://luabinaries.sourceforge.net/)中获得。[lua-users wiki](http://lua-users.org/wiki/) 列出了许多[用户为 Lua 贡献的插件](http://lua-users.org/wiki/LuaAddons)。

## Building 构建

Lua is implemented in pure ANSI C and compiles unmodified in all platforms that have an ANSI C compiler. Lua also compiles cleanly as C++.

​	Lua 是用纯 ANSI C 实现的，并且在所有具有 ANSI C 编译器的平台上无需修改即可编译。Lua 也可以干净地编译为 C++。

Lua is very easy to build and install. There are [detailed instructions](https://www.lua.org/manual/5.4/readme.html) in the package but here is a simple terminal session that downloads the current release of Lua and builds it in Linux:

​	Lua 非常容易构建和安装。软件包中有[详细的说明](https://www.lua.org/manual/5.4/readme.html)，但以下是一个简单的终端会话，它下载了 Lua 的当前版本并在 Linux 中构建它：

```
curl -L -R -O https://www.lua.org/ftp/lua-5.4.6.tar.gz
tar zxf lua-5.4.6.tar.gz
cd lua-5.4.6
make all test
```

If you have trouble building Lua, [read the FAQ](https://www.lua.org/faq.html#1.1).

​	如果您在构建 Lua 时遇到问题，请[阅读常见问题解答](https://www.lua.org/faq.html#1.1)。

If you don't have the time or the inclination to compile Lua yourself, [get a binary](https://luabinaries.sourceforge.net/) or try the [live demo](https://www.lua.org/demo.html).

​	如果您没有时间或倾向自己编译 Lua，请[获取一个二进制文件](https://luabinaries.sourceforge.net/)或尝试[实时演示](https://www.lua.org/demo.html)。

## Giving credit 给予信用

If you use Lua, please give us credit, according to our [license](https://www.lua.org/license.html). A nice way to give us further credit is to include a [Lua logo](https://www.lua.org/images/) and a [link to our site](https://www.lua.org/) in a web page for your product.

​	如果您使用 Lua，请根据我们的许可证给我们信用。在您的产品网页中包含一个 Lua 徽标和指向我们网站的链接，这是给我们进一步认可的一种好方法。



## Supporting Lua 支持 Lua

You can help to [support the Lua project](https://www.lua.org/donations.html) by [buying a book](https://www.lua.org/donations.html#books) published by Lua.org and by [making a donation](https://www.lua.org/donations.html#donation).

​	您可以通过购买 Lua.org 出版的书籍和捐款来帮助支持 Lua 项目。

You can also help to spread the word about Lua by buying Lua products at [Zazzle](https://www.zazzle.com/store/lua_store).

​	您还可以在 Zazzle 购买 Lua 产品来帮助宣传 Lua。