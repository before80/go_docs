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

​	当您开发模块供其他开发者使用时，您可以遵循一个工作流程，以帮助确保使用该模块的开发者获得可靠、一致的体验。本主题描述了该工作流程中的高级步骤。

​	关于模块开发的概述，请参阅[开发和发布模块](../DevelopingAndPublishingModules)。

另请参阅：

- 如果您只是想在您的代码中使用外部包，请一定要看[管理依赖项]({{< ref "/docs/UsingAndUnderstandingGo/ManagingDependencies">}})。
- 在每一个新的版本中，您都要用版本号来表示您的模块的变化。更多信息，请参见[模块版本号](../ModuleVersionNumbering)。

## 常见的工作流程步骤

​	下面的顺序说明了一个新模块例子的发布和版本管理工作流程步骤。关于每个步骤的更多信息，请参见本主题中的章节。

1. **启用一个模块**，并组织它的源代码，使它更容易被开发人员使用，也更容易被您维护。

   如果您是开发模块的新手，请查看[教程：创建一个Go模块]({{< ref "/docs/GettingStarted/TutorialCreateAGoModule">}})。

   在Go的分布式模块发布系统中，您如何组织您的代码很重要。更多信息，请看[管理模块来源](../ManagingModuleSource)。

2. 开始**编写本地客户端代码**，调用未发布模块中的函数。

   在您发布模块之前，它对于使用诸如 `go get` 命令的典型依赖管理工作流程是不可用的。在这个阶段测试模块代码的一个好方法是将其放在与调用代码相同的本地目录中，并在此目录中进行尝试。。

   关于本地开发的更多信息，请参见[针对未发布模块的编码](#针对未发布的模块进行编码)。

3. 当模块的代码已经准备好供其他开发人员尝试时，可以**开始发布 v0 的预发布版本**，如 alpha 版和 beta 版。更多信息请参阅[发布预发布版本](#发布预发布版本)。

4. **发布一个不保证稳定的v0版本**，但用户可以试用。更多内容请参见[发布第一个（不稳定）版本](#发布第一个不稳定版本)。

5. 在您的v0版本发布之后，您可以（而且应该！）继续**发布新的版本**。

   这些新版本可能包括 bug 修复（补丁版本 - patch releases）、模块公共 API 的添加（次要版本 - minor releases）甚至是破坏性的变更。由于 v0 版本不保证稳定性或向后兼容性，因此可以在其版本中进行破坏性的变更。

   更多信息，请参阅[发布bug修复](#发布bug修复)和[发布非破坏性的API变更](#发布非破坏性的API变更)。

6. 当您准备发布一个稳定版本时，您可以**将预发布版本发布为 alpha 和 beta 版本**。有关更多信息，请参阅[发布预发布版本](#发布预发布版本)。

7. 发布一个v1版本作为**第一个稳定版本**。

   这是对模块的稳定性做出承诺的第一个版本。更多信息，请参见[发布第一个稳定版本](#发布第一个稳定版本)。

8. 在v1版本中，**继续bug修复**，并在必要时对模块的公共 API 进行扩展。

   更多信息，请参见[发布bug修复](#发布bug修复)和[发布非破坏性的API变更](#发布非破坏性的API变更)。

9. 在无法避免的情况下，使用新的主要版本发布破坏性变更。

   主要版本更新（例如从 v1.x.x 到 v2.x.x）对于模块的用户来说可能是一个非常具有破坏性的升级。这应该是最后的选择。有关更多信息，请参阅[发布破坏性 API 变更](#发布破坏性的API变更)。

## 针对未发布的模块进行编码

​	当您开始开发一个模块或一个模块的新版本时，您还没有发布它。在发布模块之前，您将无法使用Go命令将该模块作为依赖项添加。相反，在最初阶段，在编写调用未发布模块中的函数的不同模块的客户端代码时，您需要在本地文件系统上引用该模块的副本。

You can reference a module locally from the client module’s go.mod file by using the `replace` directive in the client module’s go.mod file. For more information, see in [Requiring module code in a local directory](https://go.dev/doc/modules/managing-dependencies#local_directory).

​	您可以在客户端模块的go.mod文件中使用`replace`指令来从本地引用模块。有关更多信息，请参阅[在本地目录中要求模块代码](https://go.dev/doc/modules/managing-dependencies#local_directory)。

## 发布预发布版本

You can publish pre-release versions to make a module available for others to try it out and give you feedback. A pre-release version includes no guarantee of stability.

您可以发布预发布版本，让其他人试用模块并给您反馈。预发布版本不包括稳定性的保证。

Pre-release version numbers are appended with a pre-release identifier. For more on version numbers, see [Module version numbering](https://go.dev/doc/modules/version-numbers).

预发布的版本号后面有一个预发布的标识符。更多关于版本号的信息，请参见模块版本号。

Here are two examples:

这里有两个例子：

```
v0.2.1-beta.1
v1.2.3-alpha
```

When making a pre-release available, keep in mind that developers using the pre-release will need to explicitly specify it by version with the `go get` command. That’s because, by default, the `go` command prefers release versions over pre-release versions when locating the module you’re asking for. So developers must get the pre-release by specifying it explicitly, as in the following example:

在提供预发布版本时，请记住，使用预发布版本的开发者需要用go get命令明确指定它的版本。这是因为，默认情况下，go命令在定位您所要求的模块时，更倾向于选择发布版本而不是预发布版本。因此，开发者必须通过明确指定预发布版本来获得预发布版本，如下面的例子：

```
go get example.com/theirmodule@v1.2.3-alpha
```

You publish a pre-release by tagging the module code in your repository, specifying the pre-release identifier in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

您可以通过在您的版本库中标记模块代码来发布预发布版本，在标记中指定预发布版本的标识符。更多信息请参见发布模块。

## 发布第一个不稳定版本

As when you publish a pre-release version, you can publish release versions that don’t guarantee stability or backward compatibility, but give your users an opportunity to try out the module and give you feedback.

就像您发布预发布版本一样，您可以发布不保证稳定性或向后兼容性的发布版本，但给您的用户一个机会来试用模块并给您反馈。

Unstable releases are those whose version numbers are in the v0.x.x range. A v0 version makes no stability or backward compatibility guarantees. But it gives you a way to get feedback and refine your API before making stability commitments with v1 and later. For more see, [Module version numbering](https://go.dev/doc/modules/version-numbers).

不稳定的版本是指那些版本号在v0.x.x范围内的版本。v0版本没有稳定性或向后兼容性的保证。但它为您提供了一种方法，让您在对v1及以后的版本做出稳定性承诺之前获得反馈并完善您的API。更多信息请见，模块的版本编号。

As with other published versions, you can increment the minor and patch parts of the v0 version number as you make changes toward releasing a stable v1 version. For example, after releasing a v.0.0.0, you might release a v0.0.1 with the first set of bug fixes.

与其他发布的版本一样，当您为发布稳定的v1版本而进行修改时，您可以增加v0版本号的次要部分和补丁部分。例如，在发布了v.0.0.0之后，您可以发布带有第一组错误修复的v0.0.1。

Here’s an example version number:

下面是一个版本号的例子：

```
v0.1.3
```

You publish an unstable release by tagging the module code in your repository, specifying a v0 version number in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

您可以通过在您的版本库中标记模块代码来发布一个不稳定的版本，并在标记中指定一个v0的版本号。更多信息，请看发布模块。

## 发布第一个稳定版本

Your first stable release will have a v1.x.x version number. The first stable release follows pre-release and v0 releases through which you got feedback, fixed bugs, and stabilized the module for users.

您的第一个稳定版本会有一个v1.x.x的版本号。第一个稳定版本是在预发布和v0版本之后发布的，通过这些版本您可以得到反馈，修复错误，并为用户稳定模块。

With a v1 release, you’re making the following commitments to developers using your module:

有了v1版本，您就对使用您的模块的开发者做出了以下承诺：

- They can upgrade to the major version’s subsequent minor and patch releases without breaking their own code. 他们可以升级到主要版本的后续次要版本和补丁版本，而不会破坏自己的代码。
- You won’t be making further changes to the module’s public API – including its function and method signatures – that break backward compatibility.您不会对模块的公共API（包括其函数和方法签名）做进一步的修改，从而破坏向后的兼容性。
- You won’t be removing any exported types, which would break backward compatibility.您不会删除任何导出的类型，这将破坏后向兼容性。
- Future changes to your API (such as adding a new field to a struct) will be backward compatible and will be included in a new minor release.未来对您的API的改变（比如给结构添加一个新的字段）将是向后兼容的，并将包含在一个新的次要版本中。
- Bug fixes (such as a security fix) will be included in a patch release or as part of a minor release.错误修复（如安全修复）将包含在补丁发布中或作为小版本的一部分。

**Note:** While your first major version might be a v0 release, a v0 version does not signal stability or backward compatibility guarantees. As a result, when you increment from v0 to v1, you needn’t be mindful of breaking backward compatibility because the v0 release was not considered stable.

注意：虽然您的第一个主要版本可能是v0版本，但v0版本并不代表稳定性或向后兼容性的保证。因此，当您从v0增加到v1时，您不需要注意破坏后向兼容性，因为v0版本不被视为稳定。

For more about version numbers, see [Module version numbering](https://go.dev/doc/modules/version-numbers).

欲了解更多关于版本号的信息，请参见模块版本号。

Here’s an example of a stable version number:

下面是一个稳定版本号的例子：

```
v1.0.0
```

You publish a first stable release by tagging the module code in your repository, specifying a v1 version number in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

您通过在您的版本库中标记模块代码来发布第一个稳定版本，并在标记中指定v1的版本号。更多信息，请参见发布模块。

## 发布bug修复

You can publish a release in which the changes are limited to bug fixes. This is known as a patch release.

您可以发布一个变化仅限于错误修复的版本。这就是所谓的补丁发布。

A *patch release* includes only minor changes. In particular, it includes no changes to the module’s public API. Developers of consuming code can upgrade to this version safely and without needing to change their code.

补丁发布只包括微小的变化。特别是，它不包括对模块的公共API的改变。消耗代码的开发者可以安全地升级到这个版本，而不需要改变他们的代码。

**Note:** Your patch release should try not to upgrade any of that module’s own transitive dependencies by more than a patch release. Otherwise, someone upgrading to the patch of your module could wind up accidentally pulling in a more invasive change to a transitive dependency that they use.

注意：您的补丁版本应该尽量不要将该模块自身的横向依赖关系升级到一个以上的补丁版本。否则，升级到您的模块的补丁的人可能会意外地对他们使用的过渡性依赖关系进行更多的修改。

A patch release increments the patch part of the module’s version number. For more see, [Module version numbering](https://go.dev/doc/modules/version-numbers).

补丁的发布会增加模块的版本号中的补丁部分。更多信息请看，模块版本号。

In the following example, v1.0.1 is a patch release.

在下面的例子中，v1.0.1是一个补丁版本。

Old version: `v1.0.0`

New version: `v1.0.1`

You publish a patch release by tagging the module code in your repository, incrementing the patch version number in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

您可以通过在版本库中标记模块代码来发布一个补丁版本，在标记中增加补丁版本号。更多信息请参见发布模块。

## 发布非破坏性的API变更

You can make non-breaking changes to your module’s public API and publish those changes in a *minor* version release.

您可以对您的模块的公共API进行非破坏性的修改，并将这些修改发布在一个小版本中。

This version changes the API, but not in a way that breaks calling code. This might include changes to a module’s own dependencies or the addition of new functions, methods, struct fields, or types. Even with the changes it includes, this kind of release guarantees backward compatibility and stability for existing code that calls the module’s functions.

这个版本会改变API，但不会破坏调用代码。这可能包括对模块本身的依赖项的改变，或者增加新的函数、方法、结构域或类型。即使它包括了一些变化，这种版本也能保证调用该模块功能的现有代码的向后兼容性和稳定性。

A minor release increments the minor part of the module’s version number. For more, see [Module version numbering](https://go.dev/doc/modules/version-numbers).

一个次要版本会增加模块的版本号的次要部分。更多信息请参见模块版本号。

In the following example, v1.1.0 is a minor release.

在下面的例子中，v1.1.0是一个次要版本。

Old version: `v1.0.1`

New version: `v1.1.0`

You publish a minor release by tagging the module code in your repository, incrementing the minor version number in the tag. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).

您可以通过在版本库中标记模块代码来发布一个次要版本，在标记中增加次要版本号。更多信息请参见发布模块。

## 发布破坏性的API变更

You can publish a version that breaks backward compatibility by publishing a *major* version release.

您可以通过发布一个主要版本来发布一个破坏后向兼容性的版本。

A major version release doesn’t guarantee backward compatibility, typically because it includes changes to the module’s public API that would break code using the module’s previous versions.

一个主要版本的发布并不能保证向后兼容，通常是因为它包括对模块的公共API的修改，这些修改会破坏使用该模块以前版本的代码。

Given the disruptive effect a major version upgrade can have on code relying on the module, you should avoid a major version update if you can. For more about major version updates, see [Developing a major version update](https://go.dev/doc/modules/major-version). For strategies to avoid making breaking changes, see the blog post [Keeping your modules compatible](https://blog.golang.org/module-compatibility).

考虑到主要版本升级对依赖该模块的代码可能产生的破坏性影响，如果可以的话，您应该避免主要版本更新。更多关于主要版本更新的信息，请看开发主要版本更新。关于避免破坏性修改的策略，请看博客文章《保持您的模块的兼容性》。

Where publishing other kinds of versions requires essentially tagging the module code with the version number, publishing a major version update requires more steps.

发布其他类型的版本只需要在模块代码上标记版本号，而发布主要版本更新需要更多的步骤。

1. Before beginning development of the new major version, in your repository create a place for the new version’s source. 在开始开发新的主要版本之前，在您的版本库中为新版本的源代码创建一个地方。

   One way to do this is to create a new branch in your repository that is specifically for the new major version and its subsequent minor and patch versions. For more, see [Managing module source](https://go.dev/doc/modules/managing-source).一种方法是在您的版本库中创建一个新的分支，专门用于新的主版本及其后续的次版本和补丁版本。更多信息请参见管理模块源代码。

2. In the module’s go.mod file, revise the module path to append the new major version number, as in the following example:在模块的go.mod文件中，修改模块的路径，添加新的主要版本号，如下面的例子。

   ```
   example.com/mymodule/v2
   ```

   Given that the module path is the module’s identifier, this change effectively creates a new module. It also changes the package path, ensuring that developers won’t unintentionally import a version that breaks their code. Instead, those wanting to upgrade will explicitly replace occurrences of the old path with the new one.考虑到模块路径是模块的标识符，这一改变有效地创建了一个新模块。它也改变了包的路径，确保开发者不会无意中导入一个破坏他们代码的版本。相反，那些想要升级的人将明确地用新的路径来替换旧的路径的出现。

3. In your code, change any package paths where you’re importing packages in the module you’re updating, including packages in the module you’re updating. You need to do this because you changed your module path.在您的代码中，改变任何您要导入的模块中的包的路径，包括您要更新的模块中的包。您需要这样做是因为您改变了您的模块路径。

4. As with any new release, you should publish pre-release versions to get feedback and bug reports before publishing an official release.与任何新版本一样，在发布正式版本之前，您应该发布预发布版本以获得反馈和错误报告。

5. Publish the new major version by tagging the module code in your repository, incrementing the major version number in the tag – such as from v1.5.2 to v2.0.0.通过在您的版本库中标记模块代码来发布新的主要版本，在标记中增加主要版本号--比如从v1.5.2到v2.0.0。

   For more, see [Publishing a module](https://go.dev/doc/modules/publishing).
   
   更多信息请参见发布模块。