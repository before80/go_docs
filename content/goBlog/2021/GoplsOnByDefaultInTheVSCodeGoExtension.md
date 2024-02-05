+++
title = "在VS Code Go扩展中默认打开gopls"
weight = 98
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Gopls on by default in the VS Code Go extension - 在VS Code Go扩展中默认打开gopls

> 原文：[https://go.dev/blog/gopls-vscode-go](https://go.dev/blog/gopls-vscode-go)

Go tools team
1 February 2021

We’re happy to announce that the VS Code Go extension now enables the [gopls language server](https://github.com/golang/tools/blob/master/gopls/README.md) by default, to deliver more robust IDE features and better support for Go modules.

​	我们很高兴地宣布，VS Code Go扩展现在默认启用 [gopls 语言服务器](https://github.com/golang/tools/blob/master/gopls/README.md)，以提供更强大的IDE功能和对Go模块的更好支持。

![img](GoplsOnByDefaultInTheVSCodeGoExtension_img/features.gif)

*(`gopls` provides IDE features, such as intelligent autocompletion, signature help, refactoring, and workspace symbol search.)*

(`gopls`提供IDE功能，如智能自动完成、签名帮助、重构和工作区符号搜索)。

When [Go modules](https://go.dev/blog/using-go-modules) were released two years ago, they completely changed the landscape of Go developer tooling. Tools like `goimports` and `godef` previously depended on the fact that code was stored in your `$GOPATH`. When the Go team began rewriting these tools to work with modules, we immediately realized that we needed a more systematic approach to bridge the gap.

​	当[Go modules](../../2019/UsingGoModules)在两年前发布时，它们完全改变了Go开发工具的面貌。像`goimports`和`godef`这样的工具以前依赖于代码存储在`$GOPATH`中的事实。当 Go 团队开始重写这些工具以便与模块一起工作时，我们立即意识到我们需要一种更系统的方法来弥补这一差距。

As a result, we began working on a single Go [language server](https://microsoft.github.io/language-server-protocol/), `gopls`, which provides IDE features, such as autocompletion, formatting, and diagnostics to any compatible editor frontend. This persistent and unified server is a [fundamental shift](https://www.youtube.com/watch?v=EFJfdWzBHwE&t=1s) from the earlier collections of command-line tools.

​	因此，我们开始着手开发一个单一的[Go语言服务器](https://microsoft.github.io/language-server-protocol/)，`gopls`，它为任何兼容的编辑器前端提供IDE功能，如自动完成、格式化和诊断功能。这种持久而统一的服务器是对早期的命令行工具集合的[根本性转变](https://www.youtube.com/watch?v=EFJfdWzBHwE&t=1s)。

In addition to working on `gopls`, we sought other ways of creating a stable ecosystem of editor tooling. Last year, the Go team took responsibility for the [Go extension for VS Code](https://blog.golang.org/vscode-go). As part of this work, we smoothed the extension’s integration with the language server—automating `gopls` updates, rearranging and clarifying `gopls` settings, improving the troubleshooting workflow, and soliciting feedback through a survey. We’ve also continued to foster a community of active users and contributors who have helped us improve the stability, performance, and user experience of the Go extension.

​	除了在`gopls`上的工作，我们还寻求其他方法来创建一个稳定的编辑器工具的生态系统。去年，Go团队负责了[VS Code的Go扩展](../../2021/TheVSCodeGoExtensionJoinsTheGoProject)。作为这项工作的一部分，我们平滑了扩展与语言服务器的整合 —— 自动更新`gopls`，重新安排和明确`gopls`的设置，改进故障排除工作流程，并通过调查征求反馈意见。我们还继续培养了一个活跃的用户和贡献者社区，他们帮助我们提高了Go扩展的稳定性、性能和用户体验。

## Announcement 公告

January 28 marked a major milestone in both the `gopls` and VS Code Go journeys, as `gopls` is now enabled by default in the Go extension for VS Code.

​	1月28日是`gopls`和VS Code Go旅程中的一个重要里程碑，因为现在VS Code的Go扩展中默认启用了`gopls`。

In advance of this switch we spent a long time iterating on the design, feature set, and user experience of `gopls`, focusing on improving performance and stability. For more than a year, `gopls` has been the default in most plugins for Vim, Emacs, and other editors. We’ve had 24 `gopls` releases, and we’re incredibly grateful to our users for consistently providing feedback and reporting issues on each and every one.

​	在这次转换之前，我们花了很长时间对`gopls`的设计、功能设置和用户体验进行了迭代，重点是提高性能和稳定性。一年多来，`gopls`一直是Vim、Emacs和其他编辑器的大多数插件中的默认配置。我们已经发布了24个`gopls`版本，我们非常感谢我们的用户不断地提供反馈和报告每一个版本的问题。

We’ve also dedicated time to smoothing the new user experience. We hope that VS Code Go with `gopls` will be intuitive with clear error messages, but if you have a question or need to adjust some configuration, you’ll be able to find answers in our [updated documentation](https://github.com/golang/vscode-go/blob/master/README.md). We have also recorded [a screencast](https://www.youtube.com/watch?v=1MXIGYrMk80) to help you get started, as well as [animations](https://github.com/golang/vscode-go/blob/master/docs/features.md) to show off some hard-to-find features.

​	我们也花了很多时间来使新的用户体验更加顺畅。我们希望使用`gopls`的VS Code Go能有直观的错误信息，但如果您有问题或需要调整一些配置，您可以在我们[更新的文档](https://github.com/golang/vscode-go/blob/master/README.md)中找到答案。我们还录制了[一个截屏](https://www.youtube.com/watch?v=1MXIGYrMk80)来帮助您入门，还有[动画](https://github.com/golang/vscode-go/blob/master/docs/features.md)来展示一些难以发现的功能。

Gopls is the best way of working with Go code, especially with Go modules. With the upcoming arrival of Go 1.16, in which modules are enabled by default, VS Code Go users will have the best possible experience out-of-the-box.

​	Gopls是处理Go代码的最佳方式，尤其是处理Go模块。随着Go 1.16的即将到来，其中的模块是默认启用的，VS Code Go用户将拥有开箱即用的最佳体验。

Still, this switch does not mean that `gopls` is complete. We will continue working on bug fixes, new features, and general stability. Our next area of focus will be improving the user experience when [working with multiple modules](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md). Feedback from our larger user base will help inform our next steps.

​	不过，这种转换并不意味着`gopls`已经完成。我们将继续致力于错误修复、新功能和总体稳定性。我们的下一个重点领域将是改善[使用多个模块](https://github.com/golang/tools/blob/master/gopls/doc/workspace.md)时的用户体验。来自我们广大用户的反馈将有助于我们采取下一步行动。

## So, what should you do? 那么，您应该怎么做？

If you use VS Code, you don’t need to do anything. When you get the next VS Code Go update, `gopls` will be enabled automatically.

​	如果您使用VS Code，您不需要做任何事情。当您得到下一个VS Code Go更新时，`gopls`将自动启用。

If you use another editor, you are likely using `gopls` already. If not, see [the `gopls` user guide](https://github.com/golang/tools/blob/master/gopls/README.md) to learn how to enable `gopls` in your preferred editor. The Language Server Protocol ensures that `gopls` will continue to offer the same features to every editor.

​	如果您使用其他编辑器，您可能已经在使用`gopls`了。如果没有，请参阅[gopls 用户指南](https://github.com/golang/tools/blob/master/gopls/README.md)，了解如何在您喜欢的编辑器中启用`gopls`。语言服务器协议确保`gopls`将继续为每个编辑器提供相同的功能。

If `gopls` is not working for you, please see our [detailed troubleshooting guide](https://github.com/golang/vscode-go/blob/master/docs/troubleshooting.md) and file an issue. If you need to, you can always [disable `gopls` in VS Code](https://github.com/golang/vscode-go/blob/master/docs/settings.md#gouselanguageserver).

​	如果`gopls`不能为您工作，请看我们[详细的故障排除指南](https://github.com/golang/vscode-go/blob/master/docs/troubleshooting.md)并提出问题。如果您需要，您可以随时在[VS Code中禁用gopls](https://github.com/golang/vscode-go/blob/master/docs/settings.md#gouselanguageserver)。

## Thank you 谢谢您

To our existing users, thank you for bearing with us as we rewrote our caching layer for the third time. To our new users, we look forward to hearing your experience reports and feedback.

​	对于我们的现有用户，感谢您们在我们第三次重写缓存层时的支持。对于我们的新用户，我们期待着听到您们的体验报告和反馈。

Finally, no discussion of Go tooling is complete without mentioning the valuable contributions of the Go tools community. Thank you for the lengthy discussions, detailed bug reports, integration tests, and most importantly, thank you for the fantastic contributions. The most exciting `gopls` features come from our passionate open-source contributors, and we are appreciative of your hard work and dedication.

​	最后，如果不提及Go工具社区的宝贵贡献，关于Go工具的讨论就不完整。感谢您们长时间的讨论、详细的错误报告、集成测试，最重要的是，感谢您们的精彩贡献。最激动人心的`gopls`功能来自于我们热情的开源贡献者，我们对您们的辛勤工作和奉献表示感谢。

## Learn more 了解更多

Watch [the screencast](https://www.youtube.com/watch?v=1MXIGYrMk80) for a walk-through of how to get started with `gopls` and VS Code Go, and see the [VS Code Go README](https://github.com/golang/vscode-go/blob/master/README.md) for additional information.

​	观看[截屏视频](https://www.youtube.com/watch?v=1MXIGYrMk80)，了解如何开始使用`gopls`和VS Code Go，并查看[VS Code Go的README](https://github.com/golang/vscode-go/blob/master/README.md)，了解更多信息。

If you’d like to read about `gopls` in more detail, see the [`gopls` README](https://github.com/golang/tools/blob/master/gopls/README.md).

​	如果您想更详细地了解`gopls`，请看[gopls README](https://github.com/golang/tools/blob/master/gopls/README.md)。
