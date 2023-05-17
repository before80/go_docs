+++
title = "模块发布和版本管理的工作流程"
weight = 2
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Module release and versioning workflow - 模块发布和版本管理的工作流程

> 原文：[https://go.dev/doc/modules/release-workflow](https://go.dev/doc/modules/release-workflow)

When you develop modules for use by other developers, you can follow a workflow that helps ensure a reliable, consistent experience for developers using the module. This topic describes the high-level steps in that workflow.

当你开发模块供其他开发者使用时，你可以遵循一个工作流程，以帮助确保使用该模块的开发者获得可靠、一致的体验。本主题描述了该工作流程中的高级步骤。

For an overview of module development, see [Developing and publishing modules](https://go.dev/doc/modules/developing).

关于模块开发的概述，请参阅开发和发布模块。

**See also**

另见

- If you’re merely wanting to use external packages in your code, be sure to see [Managing dependencies](https://go.dev/doc/modules/managing-dependencies).如果你只是想在你的代码中使用外部包，一定要看管理依赖性。
- With each new version, you signal the changes to your module with its version number. For more, see [Module version numbering](https://go.dev/doc/modules/version-numbers).在每一个新的版本中，你都要用版本号来表示你的模块的变化。更多信息，请参见模块版本号。

## Common workflow steps 常见的工作流程步骤

The following sequence illustrates release and versioning workflow steps for an example new module. For more about each step, see the sections in this topic.

下面的顺序说明了一个新模块例子的发布和版本管理工作流程步骤。关于每个步骤的更多信息，请参见本主题中的章节。

1. **Begin a module** and organize its sources to make it easier for developers to use and for you to maintain.开始一个模块，并组织它的源代码，使它更容易被开发人员使用，也更容易被你维护。

   If you’re brand new to developing modules, check out [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module).如果你是开发模块的新手，请查看教程。创建一个Go模块。

   In Go’s decentralized module publishing system, how you organize your code matters. For more, see [Managing module source](https://go.dev/doc/modules/managing-source).在Go的分散式模块发布系统中，你如何组织你的代码很重要。更多信息，请看管理模块源代码。

2. Set up to **write local client code** that calls functions in the unpublished module.设置编写本地客户端代码，调用未发布模块中的函数。

   Before you publish a module, it’s unavailable for the typical dependency management workflow using commands such as `go get`. A good way to test your module code at this stage is to try it while it is in a directory local to your calling code.在你发布一个模块之前，它在典型的依赖性管理工作流程中是不可用的，比如使用 go get 命令。在这个阶段，测试你的模块代码的一个好方法是，当它在你的调用代码的本地目录中时，就可以尝试它。

   See [Coding against an unpublished module](https://go.dev/doc/modules/release-workflow#unpublished) for more about local development.关于本地开发的更多信息，请参见针对未发布模块的编码。

3. When the module’s code is ready for other developers to try it out, **begin publishing v0 pre-releases** such as alphas and betas. See [Publishing pre-release versions](https://go.dev/doc/modules/release-workflow#pre-release) for more.当模块的代码准备好供其他开发者试用时，开始发布预发布版本，如alphas和betas。参见发布预发布版本以了解更多信息。

4. **Release a v0** that’s not guaranteed to be stable, but which users can try out. For more, see [Publishing the first (unstable) version](https://go.dev/doc/modules/release-workflow#first-unstable).发布一个不保证稳定的v0版本，但用户可以试用。更多内容请参见发布第一个（不稳定）版本。

5. After your v0 version is published, you can (and should!) continue to **release new versions** of it.在你的v0版本发布之后，你可以（而且应该！）继续发布新的版本。

   These new versions might include bug fixes (patch releases), additions to the module’s public API (minor releases), and even breaking changes. Because a v0 release makes no guarantees of stability or backward compatibility, you can make breaking changes in its versions.这些新版本可能包括错误的修正（补丁发布），对模块的公共API的补充（次要发布），甚至是突破性的改变。因为v0版本不保证稳定性或向后兼容性，所以你可以在其版本中进行破坏性的修改。

   For more, see [Publishing bug fixes](https://go.dev/doc/modules/release-workflow#bug-fixes) and [Publishing non-breaking API changes](https://go.dev/doc/modules/release-workflow#non-breaking).更多信息，请参阅发布错误修复和发布非破坏性的API变化。

6. When you’re getting a stable version ready for release, you **publish pre-releases as alphas and betas**. For more, see [Publishing pre-release versions](https://go.dev/doc/modules/release-workflow#pre-release).当你准备发布一个稳定的版本时，你可以将预发布版本作为alphas和betas发布。更多信息，请参见发布预发布版本。

7. Release a v1 as the **first stable release**.发布一个v1版本作为第一个稳定版本。

   This is the first release that makes commitments about the module’s stability. For more, see [Publishing the first stable version](https://go.dev/doc/modules/release-workflow#first-stable).这是对模块的稳定性做出承诺的第一个版本。更多信息，请参见发布第一个稳定版本。

8. In the v1 version, **continue to fix bugs** and, where necessary, make additions to the module’s public API.在v1版本中，继续修复bug，必要时对模块的公共API进行补充。

   For more, see [Publishing bug fixes](https://go.dev/doc/modules/release-workflow#bug-fixes) and [Publishing non-breaking API changes](https://go.dev/doc/modules/release-workflow#non-breaking).更多信息，请参见发布错误修复和发布非破坏性的API变化。

9. When it can’t be avoided, publish breaking changes in a **new major version**.当无法避免时，在一个新的主要版本中发布破坏性变化。

   A major version update – such as from v1.x.x to v2.x.x – can be a very disruptive upgrade for your module’s users. It should be a last resort. For more, see [Publishing breaking API changes](https://go.dev/doc/modules/release-workflow#breaking).一个主要的版本更新--比如从v1.x.x到v2.x.x--对你的模块用户来说可能是一个非常有破坏性的升级。这应该是最后的手段。更多信息，请看发布破坏性的API变化。

## Coding against an unpublished module 针对未发布的模块进行编码

When you begin developing a module or a new version of a module, you won’t yet have published it. Before you publish a module, you won’t be able to use Go commands to add the module as a dependency. Instead, at first, when writing client code in a different module that calls functions in the unpublished module, you’ll need to reference a copy of the module on the local file system.

当你开始开发一个模块或一个模块的新版本时，你还没有发布它。在你发布模块之前，你无法使用Go命令将该模块作为一个依赖项添加。相反，一开始，当在不同的模块中写客户端代码调用未发布的模块中的函数时，你需要在本地文件系统中引用该模块的一个副本。

You can reference a module locally from the client module’s go.mod file by using the `replace` directive in the client module’s go.mod file. For more information, see in [Requiring module code in a local directory](https://go.dev/doc/modules/managing-dependencies#local_directory).

你可以通过在客户端模块的go.mod文件中使用替换指令，在本地引用一个模块。更多信息，请参见《要求本地目录中的模块代码》。

## Publishing pre-release versions 发布预发布版本

You can publish pre-release versions to make a module available for others to try it out and give you feedback. A pre-release version includes no guarantee of stability.

你可以发布预发布版本，让其他人试用模块并给你反馈。预发布版本不包括稳定性的保证。

Pre-release version numbers are appended with a pre-release identifier. For more on version numbers, see [Module version numbering](https://go.dev/doc/modules/version-numbers).

预发布的版本号后面有一个预发布的标识符。更多关于版本号的信息，请参见模块版本号。

Here are two examples:

这里有两个例子：

```
v0.2.1-beta.1
v1.2.3-alpha
```

When making a pre-release available, keep in mind that developers using the pre-release will need to explicitly specify it by version with the `go get` command. That’s because, by default, the `go` command prefers release versions over pre-release versions when locating the module you’re asking for. So developers must get the pre-release by specifying it explicitly, as in the following example:

在提供预发布版本时，请记住，使用预发布版本的开发者需要用go get命令明确指定它的版本。这是因为，默认情况下，go命令在定位你所要求的模块时，更倾向于选择发布版本而不是预发布版本。因此，开发者必须通过明确指定预发布版本来获得预发布版本，如下面的例子：

```
go get example.com/theirmodule@v1.2.3-alpha
```

You publish a pre-release by tagging the module code in your repository, specifying the pre-release identifier in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

你可以通过在你的版本库中标记模块代码来发布预发布版本，在标记中指定预发布版本的标识符。更多信息请参见发布模块。

## Publishing the first (unstable) version 发布第一个（不稳定）版本

As when you publish a pre-release version, you can publish release versions that don’t guarantee stability or backward compatibility, but give your users an opportunity to try out the module and give you feedback.

就像你发布预发布版本一样，你可以发布不保证稳定性或向后兼容性的发布版本，但给你的用户一个机会来试用模块并给你反馈。

Unstable releases are those whose version numbers are in the v0.x.x range. A v0 version makes no stability or backward compatibility guarantees. But it gives you a way to get feedback and refine your API before making stability commitments with v1 and later. For more see, [Module version numbering](https://go.dev/doc/modules/version-numbers).

不稳定的版本是指那些版本号在v0.x.x范围内的版本。v0版本没有稳定性或向后兼容性的保证。但它为你提供了一种方法，让你在对v1及以后的版本做出稳定性承诺之前获得反馈并完善你的API。更多信息请见，模块的版本编号。

As with other published versions, you can increment the minor and patch parts of the v0 version number as you make changes toward releasing a stable v1 version. For example, after releasing a v.0.0.0, you might release a v0.0.1 with the first set of bug fixes.

与其他发布的版本一样，当你为发布稳定的v1版本而进行修改时，你可以增加v0版本号的次要部分和补丁部分。例如，在发布了v.0.0.0之后，你可以发布带有第一组错误修复的v0.0.1。

Here’s an example version number:

下面是一个版本号的例子：

```
v0.1.3
```

You publish an unstable release by tagging the module code in your repository, specifying a v0 version number in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

你可以通过在你的版本库中标记模块代码来发布一个不稳定的版本，并在标记中指定一个v0的版本号。更多信息，请看发布模块。

## Publishing the first stable version 发布第一个稳定版本

Your first stable release will have a v1.x.x version number. The first stable release follows pre-release and v0 releases through which you got feedback, fixed bugs, and stabilized the module for users.

你的第一个稳定版本会有一个v1.x.x的版本号。第一个稳定版本是在预发布和v0版本之后发布的，通过这些版本你可以得到反馈，修复错误，并为用户稳定模块。

With a v1 release, you’re making the following commitments to developers using your module:

有了v1版本，你就对使用你的模块的开发者做出了以下承诺：

- They can upgrade to the major version’s subsequent minor and patch releases without breaking their own code. 他们可以升级到主要版本的后续次要版本和补丁版本，而不会破坏自己的代码。
- You won’t be making further changes to the module’s public API – including its function and method signatures – that break backward compatibility.你不会对模块的公共API（包括其函数和方法签名）做进一步的修改，从而破坏向后的兼容性。
- You won’t be removing any exported types, which would break backward compatibility.你不会删除任何导出的类型，这将破坏后向兼容性。
- Future changes to your API (such as adding a new field to a struct) will be backward compatible and will be included in a new minor release.未来对你的API的改变（比如给结构添加一个新的字段）将是向后兼容的，并将包含在一个新的次要版本中。
- Bug fixes (such as a security fix) will be included in a patch release or as part of a minor release.错误修复（如安全修复）将包含在补丁发布中或作为小版本的一部分。

**Note:** While your first major version might be a v0 release, a v0 version does not signal stability or backward compatibility guarantees. As a result, when you increment from v0 to v1, you needn’t be mindful of breaking backward compatibility because the v0 release was not considered stable.

注意：虽然你的第一个主要版本可能是v0版本，但v0版本并不代表稳定性或向后兼容性的保证。因此，当你从v0增加到v1时，你不需要注意破坏后向兼容性，因为v0版本不被视为稳定。

For more about version numbers, see [Module version numbering](https://go.dev/doc/modules/version-numbers).

欲了解更多关于版本号的信息，请参见模块版本号。

Here’s an example of a stable version number:

下面是一个稳定版本号的例子：

```
v1.0.0
```

You publish a first stable release by tagging the module code in your repository, specifying a v1 version number in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

你通过在你的版本库中标记模块代码来发布第一个稳定版本，并在标记中指定v1的版本号。更多信息，请参见发布模块。

## Publishing bug fixes 发布错误修复

You can publish a release in which the changes are limited to bug fixes. This is known as a patch release.

你可以发布一个变化仅限于错误修复的版本。这就是所谓的补丁发布。

A *patch release* includes only minor changes. In particular, it includes no changes to the module’s public API. Developers of consuming code can upgrade to this version safely and without needing to change their code.

补丁发布只包括微小的变化。特别是，它不包括对模块的公共API的改变。消耗代码的开发者可以安全地升级到这个版本，而不需要改变他们的代码。

**Note:** Your patch release should try not to upgrade any of that module’s own transitive dependencies by more than a patch release. Otherwise, someone upgrading to the patch of your module could wind up accidentally pulling in a more invasive change to a transitive dependency that they use.

注意：你的补丁版本应该尽量不要将该模块自身的横向依赖关系升级到一个以上的补丁版本。否则，升级到你的模块的补丁的人可能会意外地对他们使用的过渡性依赖关系进行更多的修改。

A patch release increments the patch part of the module’s version number. For more see, [Module version numbering](https://go.dev/doc/modules/version-numbers).

补丁的发布会增加模块的版本号中的补丁部分。更多信息请看，模块版本号。

In the following example, v1.0.1 is a patch release.

在下面的例子中，v1.0.1是一个补丁版本。

Old version: `v1.0.0`

New version: `v1.0.1`

You publish a patch release by tagging the module code in your repository, incrementing the patch version number in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

你可以通过在版本库中标记模块代码来发布一个补丁版本，在标记中增加补丁版本号。更多信息请参见发布模块。

## Publishing non-breaking API changes 发布非破坏性的API变更

You can make non-breaking changes to your module’s public API and publish those changes in a *minor* version release.

你可以对你的模块的公共API进行非破坏性的修改，并将这些修改发布在一个小版本中。

This version changes the API, but not in a way that breaks calling code. This might include changes to a module’s own dependencies or the addition of new functions, methods, struct fields, or types. Even with the changes it includes, this kind of release guarantees backward compatibility and stability for existing code that calls the module’s functions.

这个版本会改变API，但不会破坏调用代码。这可能包括对模块本身的依赖性的改变，或者增加新的函数、方法、结构域或类型。即使它包括了一些变化，这种版本也能保证调用该模块功能的现有代码的向后兼容性和稳定性。

A minor release increments the minor part of the module’s version number. For more, see [Module version numbering](https://go.dev/doc/modules/version-numbers).

一个次要版本会增加模块的版本号的次要部分。更多信息请参见模块版本号。

In the following example, v1.1.0 is a minor release.

在下面的例子中，v1.1.0是一个次要版本。

Old version: `v1.0.1`

New version: `v1.1.0`

You publish a minor release by tagging the module code in your repository, incrementing the minor version number in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

你可以通过在版本库中标记模块代码来发布一个次要版本，在标记中增加次要版本号。更多信息请参见发布模块。

## Publishing breaking API changes 发布破坏性的API变化

You can publish a version that breaks backward compatibility by publishing a *major* version release.

你可以通过发布一个主要版本来发布一个破坏后向兼容性的版本。

A major version release doesn’t guarantee backward compatibility, typically because it includes changes to the module’s public API that would break code using the module’s previous versions.

一个主要版本的发布并不能保证向后兼容，通常是因为它包括对模块的公共API的修改，这些修改会破坏使用该模块以前版本的代码。

Given the disruptive effect a major version upgrade can have on code relying on the module, you should avoid a major version update if you can. For more about major version updates, see [Developing a major version update](https://go.dev/doc/modules/major-version). For strategies to avoid making breaking changes, see the blog post [Keeping your modules compatible](https://blog.golang.org/module-compatibility).

考虑到主要版本升级对依赖该模块的代码可能产生的破坏性影响，如果可以的话，你应该避免主要版本更新。更多关于主要版本更新的信息，请看开发主要版本更新。关于避免破坏性修改的策略，请看博客文章《保持你的模块的兼容性》。

Where publishing other kinds of versions requires essentially tagging the module code with the version number, publishing a major version update requires more steps.

发布其他类型的版本只需要在模块代码上标记版本号，而发布主要版本更新需要更多的步骤。

1. Before beginning development of the new major version, in your repository create a place for the new version’s source. 在开始开发新的主要版本之前，在你的版本库中为新版本的源代码创建一个地方。

   One way to do this is to create a new branch in your repository that is specifically for the new major version and its subsequent minor and patch versions. For more, see [Managing module source](https://go.dev/doc/modules/managing-source).一种方法是在你的版本库中创建一个新的分支，专门用于新的主版本及其后续的次版本和补丁版本。更多信息请参见管理模块源代码。

2. In the module’s go.mod file, revise the module path to append the new major version number, as in the following example:在模块的go.mod文件中，修改模块的路径，添加新的主要版本号，如下面的例子。

   ```
   example.com/mymodule/v2
   ```

   Given that the module path is the module’s identifier, this change effectively creates a new module. It also changes the package path, ensuring that developers won’t unintentionally import a version that breaks their code. Instead, those wanting to upgrade will explicitly replace occurrences of the old path with the new one.考虑到模块路径是模块的标识符，这一改变有效地创建了一个新模块。它也改变了包的路径，确保开发者不会无意中导入一个破坏他们代码的版本。相反，那些想要升级的人将明确地用新的路径来替换旧的路径的出现。

3. In your code, change any package paths where you’re importing packages in the module you’re updating, including packages in the module you’re updating. You need to do this because you changed your module path.在你的代码中，改变任何你要导入的模块中的包的路径，包括你要更新的模块中的包。你需要这样做是因为你改变了你的模块路径。

4. As with any new release, you should publish pre-release versions to get feedback and bug reports before publishing an official release.与任何新版本一样，在发布正式版本之前，你应该发布预发布版本以获得反馈和错误报告。

5. Publish the new major version by tagging the module code in your repository, incrementing the major version number in the tag – such as from v1.5.2 to v2.0.0.通过在你的版本库中标记模块代码来发布新的主要版本，在标记中增加主要版本号--比如从v1.5.2到v2.0.0。

   For more, see [Publishing a module](https://go.dev/doc/modules/publishing).
   
   更多信息请参见发布模块。