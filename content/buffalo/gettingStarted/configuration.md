+++
title = "configuration"
date = 2024-02-04T21:06:28+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/getting_started/configuration/](https://gobuffalo.io/documentation/getting_started/configuration/)

# Configuration 配置 

In this chapter, you’ll learn how to manage configuration with Buffalo.

​	在本章中，您将学习如何使用 Buffalo 管理配置。

Environment variables are a good way to separate environment specific values, or secrets, from your application code base ([as described in The Twelve Factor app](https://12factor.net/config)). It can help define behavior that is based on the context of the app (as requiring SSL on production) and isolate secrets from your code base. This way, all developers don’t have to know the productions keys to sensitive services, such as a bank API, and they can use sandbox API keys.

​	环境变量是一种将特定于环境的值或机密与应用程序代码库分隔开来的好方法（如十二要素应用程序中所述）。它可以帮助定义基于应用程序上下文的行为（如在生产中要求 SSL）并将机密与代码库隔离。这样，所有开发人员都不必知道敏感服务的生产密钥，例如银行 API，并且他们可以使用沙盒 API 密钥。

## Available Environment Variables 可用的环境变量 

The following variables are used by Buffalo:

​	Buffalo 使用以下变量：

| Variable 变量            | Default 默认值                                    | Usage 用法                                                   |
| ------------------------ | ------------------------------------------------- | ------------------------------------------------------------ |
| `GO_ENV`                 | `development`                                     | The environment (dev, qa, production etc.) that Buffalo is run in. Buffalo 运行的环境（dev、qa、production 等）。 |
| `GO_BIN`                 | `go`                                              | The Go compiler to use. 要使用的 Go 编译器。                 |
| `BUFFALO_PLUGIN_PATH`    | `$PATH`                                           | Where Buffalo looks for plugins. Buffalo 在哪里查找插件。    |
| `BUFFALO_PLUGIN_TIMEOUT` | `1s`                                              | How long Buffalo waits for a plugin to respond. Buffalo 等待插件响应的时间。 |
| `ADDR`                   | `127.0.0.1` or `0.0.0.0` `127.0.0.1` 或 `0.0.0.0` | Which address to bind the server to. 将服务器绑定到哪个地址。 |
| `PORT`                   | `3000`                                            | Which port to bind the server to. 将服务器绑定到哪个端口。   |
| `HOST`                   | `http://127.0.0.1:$PORT`                          | The “URL” of the application (i.e. what end users type in). 应用程序的“URL”（即最终用户输入的内容）。 |
| `SESSION_SECRET`         | `""`                                              | A salt used for securing sessions. 用于保护会话的盐。        |

## Custom Configuration 自定义配置 

You still can provide your own variables, and retrieve them from within your application. The [envy](https://github.com/gobuffalo/envy) package makes it easy!

​	您仍然可以提供自己的变量，并在应用程序中检索它们。envy 包让这一切变得很容易！

```go
import "github.com/gobuffalo/envy"

// [...]

// Get MYSECRET env variable, default to empty string if it's not set
var MYSECRET = envy.Get("MYSECRET", "")

// Get REQUIREDSECRET env variable, return an error if it's not set
REQUIREDSECRET, err := envy.MustGet("REQUIREDSECRET")
```

## Support for .env Files 对 .env 文件的支持 

Buffalo ships with `.env` support (**since buffalo >= 0.10.3**), meaning buffalo will load `.env` files into environment variables once the application starts. To do it, Buffalo uses [`envy.Load`](https://github.com/gobuffalo/envy/blob/e613c80275b86293880eddeb27417c9a7c670ff3/envy.go#L53) which will look for `.env` file at the root of your app.

​	Buffalo 附带 `.env` 支持（自 buffalo >= 0.10.3 起），这意味着 buffalo 会在应用程序启动后将 `.env` 文件加载到环境变量中。为此，Buffalo 使用 `envy.Load` ，它将在您的应用程序根目录中查找 `.env` 文件。

If you’re not familiar with how a `.env` file looks, here is an example:

​	如果您不熟悉 `.env` 文件的外观，这里有一个示例：

```text
SENDGRID_API_KEY=ABCCOQ7GFRVCW0ODHPFQ3FTP5SLL1Q
SENDGRID_EMAIL=email@myapp.com

APP_DEBUG=true
APP_LOG_LEVEL=debug
APP_URL=https://myapp.com
```

Generated apps (**with buffalo >= 0.10.3**) will also create a default `.env` file in your application root. This file will be watched by Buffalo for changes, but will be ignored by git (added in the `.gitignore`). This is a good way to prevent developers to push credentials by mistake.

​	生成的应用程序（buffalo >= 0.10.3）还将在您的应用程序根目录中创建一个默认 `.env` 文件。Buffalo 会监视此文件是否有更改，但会忽略 git（添加到 `.gitignore` 中）。这是一个防止开发人员错误地推送凭据的好方法。

## Next Steps 后续步骤 

- [Tooling Integration](https://gobuffalo.io/documentation/getting_started/integrations) - Work with Buffalo, using existing tools.
  工具集成 - 使用现有工具与 Buffalo 协同工作。