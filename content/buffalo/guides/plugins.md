+++
title = "插件"
date = 2024-02-04T21:17:12+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/plugins/](https://gobuffalo.io/documentation/guides/plugins/)

# Plugins 插件 

Since **0.9.1**
自 0.9.1 起



Plugins allow for 3rd party code to extend the `buffalo` command as well as its sub-commands.

​	插件允许第三方代码扩展 `buffalo` 命令及其子命令。

## Installing the buffalo-plugins Plugin 安装 buffalo-plugins 插件 

```bash
$ go get -u -v github.com/gobuffalo/buffalo-plugins
```

## Finding Available Plugins 查找可用插件 

A full list of Plugins can be found at https://toolkit.gobuffalo.io/tools?topic=plugin.

​	可以在 https://toolkit.gobuffalo.io/tools?topic=plugin 找到插件的完整列表。

To get your project listed on the Buffalo Toolkit you must tag your project on GitHub with `gobuffalo`.

​	要将您的项目列在 Buffalo Toolkit 上，您必须在 GitHub 上使用 `gobuffalo` 标记您的项目。

There are a few more tags that you can use that will help the Buffalo Toolkit better categorize your project. You can add as many of this tags to your project as is suitable. Please try to refrain from using more than just a few tags.

​	还有几个其他标签可供您使用，它们将帮助 Buffalo Toolkit 更好地对您的项目进行分类。您可以根据需要向您的项目添加任意数量的这些标签。请尽量不要使用超过几个标签。

- `plugin` - Plugins
  `plugin` - 插件
- `generator` - Generators
  `generator` - 生成器
- `middleware` - Middleware
  `middleware` - 中间件
- `pop` - Pop/Soda
  `pop` - 流行/苏打水
- `templating` - Templating
  `templating` - 模板化
- `grifts` - Grift Tasks
  `grifts` - 欺诈任务
- `deployment` - Deployment
  `deployment` - 部署
- `testing` - Testing
  `testing` - 测试
- `example` - Example Apps
  `example` - 示例应用程序
- `worker` - Workers/Adapters
  `worker` - 工作人员/适配器
- `webpack` - Webpack/Front-End
  `webpack` - Webpack/前端

Any other tags will still be indexed and searchable, but the tool may not show in the “known” categories section.

​	任何其他标签仍将被索引并可搜索，但该工具可能不会显示在“已知”类别部分。

## How Does Buffalo Find Plugins? Buffalo 如何查找插件？ 

Buffalo plugins have a set of rules that must be followed for them to be consider, by Buffalo, as a plugin.

​	Buffalo 插件有一组规则，必须遵循这些规则才能被 Buffalo 视为插件。

- Plugins must be named in the format of `buffalo-<plugin-name>`. For example, `buffalo-myplugin`.
  插件必须以 `buffalo-<plugin-name>` 的格式命名。例如， `buffalo-myplugin` 。

- Plugins must be executable and must be available in one of the following places:

  
  插件必须可执行，并且必须位于以下位置之一：

  - in the `$BUFFALO_PLUGIN_PATH`
    在 `$BUFFALO_PLUGIN_PATH`
  - if not set, `$GOPATH/bin`, is tried
    如果未设置，则尝试 `$GOPATH/bin`
  - in the `./plugins` folder of your Buffalo application
    在 Buffalo 应用程序的 `./plugins` 文件夹中

- Plugins must implement an `available` command that prints a JSON response listing the available commands.
  插件必须实现一个 `available` 命令，该命令打印一个 JSON 响应，列出可用的命令。

The `buffalo plugins list` command will print a table of plugins that Buffalo sees as “available” to you.

​	 `buffalo plugins list` 命令将打印一个 Buffalo 视为“可供您使用”的插件表。

## With Configuration 带配置 

Since **1.1.0**
自 1.1.0 起



When a `config/buffalo-plugins.toml` file is present, Buffalo will scope the list of plugins that are “available” to those listed in the configuration file.

​	当存在 `config/buffalo-plugins.toml` 文件时，Buffalo 将把“可用”插件列表限定为配置文件中列出的那些插件。

```bash
$ cat config/buffalo-plugins.toml

[[plugin]]
  binary = "buffalo-pop"
  go_get = "github.com/gobuffalo/buffalo-pop"
$ buffalo plugins list

Bin         |Command               |Description
---         |---                   |---
buffalo-pop |buffalo db            |[DEPRECATED] please use `buffalo pop` instead.
buffalo-pop |buffalo destroy model |Destroys model files.
buffalo-pop |buffalo pop           |A tasty treat for all your database needs
```

## Without Configuration 无配置 

Without a configuration file, Buffalo will try to aforementioned paths to find any, and all Buffalo plugins installed on the users system.

​	如果没有配置文件，Buffalo 将尝试上述路径以查找用户系统上安装的所有 Buffalo 插件。

```bash
$ buffalo plugins list

Bin              |Command                    |Description
---              |---                        |---
buffalo-auth     |buffalo generate auth      |Generates a full auth implementation
buffalo-pop      |buffalo db                 |[DEPRECATED] please use `buffalo pop` instead.
buffalo-goth     |buffalo generate goth-auth |Generates a full auth implementation use Goth
buffalo-goth     |buffalo generate goth      |generates a actions/auth.go file configured to the specified providers.
buffalo-heroku   |buffalo heroku             |helps with heroku setup and deployment for buffalo applications
buffalo-plugins  |buffalo events listen      |listens to github.com/gobuffalo/events
buffalo-pop      |buffalo destroy model      |Destroys model files.
buffalo-plugins  |buffalo generate plugin    |generates a new buffalo plugin
buffalo-plugins  |buffalo plugins            |tools for working with buffalo plugins
buffalo-pop      |buffalo pop                |A tasty treat for all your database needs
buffalo-trash    |buffalo trash              |destroys and recreates a buffalo app
buffalo-upgradex |buffalo upgradex           |updates Buffalo and/or Pop/Soda as well as your app
```

## Installing Plugins 安装插件 

Since **1.1.0**
自 1.1.0 起



To add support for the plugin manager, one can either manually edit `./config/buffalo-plugins.toml` or let `buffalo plugins install` create it for you.

​	要添加对插件管理器的支持，可以手动编辑 `./config/buffalo-plugins.toml` 或让 `buffalo plugins install` 为您创建它。

Install command
安装命令

Config file
配置文件

Resulting plugin list
生成的插件列表

```bash
// $ buffalo plugins install

go get github.com/gobuffalo/buffalo-pop
./config/buffalo-plugins.toml
```

The `buffalo-pop` plugin was automatically added because the application in this example is a Buffalo application that uses Pop.

​	 `buffalo-pop` 插件已自动添加，因为此示例中的应用程序是使用 Pop 的 Buffalo 应用程序。

New plugins can be install in bulk with the `install` command

​	可以使用 `install` 命令批量安装新插件

Bulk Install command
批量安装命令

Config file
配置文件

Resulting plugin list
生成的插件列表

```bash
// $ buffalo plugins install
$ buffalo plugins install github.com/markbates/buffalo-trash github.com/gobuffalo/buffalo-heroku

go get github.com/gobuffalo/buffalo-heroku
go get github.com/gobuffalo/buffalo-pop
go get github.com/markbates/buffalo-trash
./config/buffalo-plugins.toml
```

## Removing Plugins 删除插件 # 自 1.1.0 起

Since **1.1.0**
可以使用 命令删除插件。这只会将它们从配置文件中删除，而不会从用户系统中删除。 删除命令



Plugins can be removed with the `remove` command. This only removes them from the config file, not from the users system.

Remove command

Config file
配置文件

Resulting plugin list
生成的插件列表

```bash
// $ buffalo plugins remove
$ buffalo plugins remove github.com/gobuffalo/buffalo-heroku

./config/buffalo-plugins.toml
```

## Writing a Plugin 编写插件 

First, you must understand [how Buffalo finds plugins](https://gobuffalo.io/documentation/guides/plugins/#how-does-buffalo-find-plugins), before you can successfully write one.

​	首先，您必须了解 Buffalo 如何查找插件，然后才能成功编写一个插件。

The `buffalo-plugins` plugin adds a new generator to `buffalo generate` to help you build a new plugin quickly

​	 `buffalo-plugins` 插件为 `buffalo generate` 添加了一个新的生成器，以帮助您快速构建一个新插件

```bash
$ buffalo generate plugin -h

buffalo generate plugin github.com/foo/buffalo-bar

Usage:
  buffalo-plugins plugin [flags]

Flags:
  -a, --author string       author's name
  -d, --dry-run             run the generator without creating files or running commands
  -f, --force               will delete the target directory if it exists
  -h, --help                help for plugin
  -l, --license string      choose a license from: [agpl, isc, lgpl-v2.1, mozilla, no-license, artistic, bsd, eclipse, lgpl-v3, mit, apache, bsd-3-clause, unlicense, cc0, gpl-v2, gpl-v3] (default "mit")
  -s, --short-name string   a 'short' name for the package
      --with-gen            creates a generator plugin
```

LICENSE
许可证

Makefile

main.go

bar/version.go

cmd/available.go

cmd/bar.go

cmd/root.go

cmd/version.go

```text
The MIT License (MIT)

Copyright (c) 2018 Mark Bates

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Writing Non-Go Plugins 编写非 Go 插件 

Plugins do not need to be written in Go. They can be written in any language you would like, as long as they comply with the rules above.

​	插件无需用 Go 编写。只要符合上述规则，您就可以使用任何您喜欢的语言编写插件。

For example, we can write the following plugin using Ruby:

​	例如，我们可以使用 Ruby 编写以下插件：

```ruby
#!/usr/bin/env ruby
# ./plugins/buffalo-hello.rb

require 'json'

command = ARGV[0]

case command
when 'available'
  puts JSON.generate([{ name: 'hello', buffalo_command: 'root', description: 'says hello to you' }])
when 'hello'
  puts 'Hi there!'

end
```

To activate the plugin we need to add the file as `buffalo-hello.rb` to somewhere in the `$PATH` or in a directory called `plugins/` inside of a Buffalo application.

​	要激活插件，我们需要将文件作为 `buffalo-hello.rb` 添加到 `$PATH` 中的某个位置，或添加到 Buffalo 应用程序内的名为 `plugins/` 的目录中。

Finally the file needs to be made executable. On a Mac/Linux it can be done with `chmod +x buffalo-hello.rb`.

​	最后，该文件需要可执行。在 Mac/Linux 上，可以使用 `chmod +x buffalo-hello.rb` 来执行此操作。

```bash
$ buffalo plugins list

Bin              |Command                    |Description
---              |---                        |---
buffalo-auth     |buffalo generate auth      |Generates a full auth implementation
buffalo-pop      |buffalo db                 |[DEPRECATED] please use `buffalo pop` instead.
buffalo-goth     |buffalo generate goth-auth |Generates a full auth implementation use Goth
buffalo-goth     |buffalo generate goth      |generates a actions/auth.go file configured to the specified providers.
buffalo-hello.rb |buffalo hello              |says hello to you
buffalo-heroku   |buffalo heroku             |helps with heroku setup and deployment for buffalo applications
buffalo-plugins  |buffalo events listen      |listens to github.com/gobuffalo/events
buffalo-pop      |buffalo destroy model      |Destroys model files.
buffalo-plugins  |buffalo generate plugin    |generates a new buffalo plugin
buffalo-plugins  |buffalo plugins            |tools for working with buffalo plugins
buffalo-pop      |buffalo pop                |A tasty treat for all your database needs
buffalo-trash    |buffalo trash              |destroys and recreates a buffalo app
buffalo-upgradex |buffalo upgradex           |updates Buffalo and/or Pop/Soda as well as your app
```