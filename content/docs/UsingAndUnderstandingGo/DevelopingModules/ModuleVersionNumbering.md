+++
title = "模块的版本编号"
weight = 6
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Module version numbering - 模块的版本编号

> 原文：[https://go.dev/doc/modules/version-numbers](https://go.dev/doc/modules/version-numbers)

A module’s developer uses each part of a module’s version number to signal the version’s stability and backward compatibility. For each new release, a module’s release version number specifically reflects the nature of the module’s changes since the preceding release.

一个模块的开发者使用一个模块的版本号的每一部分来表示该版本的稳定性和向后兼容性。对于每一个新的版本，模块的发布版本号具体反映了模块自前一个版本以来的变化性质。

When you’re developing code that uses external modules, you can use the version numbers to understand an external module’s stability when you’re considering an upgrade. When you’re developing your own modules, your version numbers will signal your modules' stability and backward compatibility to other developers.

当你在开发使用外部模块的代码时，当你考虑升级时，你可以使用版本号来了解外部模块的稳定性。当你开发你自己的模块时，你的版本号将向其他开发者表明你的模块的稳定性和向后兼容性。

This topic describes what module version numbers mean.

本主题描述了模块版本号的含义。

**See also**

另见

- When you’re using external packages in your code, you can manage those dependencies with Go tools. For more, see [Managing dependencies](https://go.dev/doc/modules/managing-dependencies).当你在代码中使用外部包时，你可以用Go工具管理这些依赖关系。更多信息，请参见管理依赖关系。
- If you’re developing modules for others to use, you apply a version number when you publish the module, tagging the module in its repository. For more, see [Publishing a module](https://go.dev/doc/modules/publishing).如果你正在开发模块供他人使用，你可以在发布模块时应用一个版本号，在其仓库中标记该模块。更多信息，请参见发布模块。

A released module is published with a version number in the semantic versioning model, as in the following illustration:

发布的模块在语义版本模型中带有一个版本号，如下图所示：

![Diagram illustrating a semantic version number showing major version 1, minor version 4, patch version 0, and pre-release version beta 2](ModuleVersionNnumbering_img/version-number.png)

Diagram illustrating a semantic version number showing major version 1, minor version 4, patch version 0, and pre-release version beta 2

语义版本号示意图，显示主版本1、次版本4、补丁版本0和预发布版本β2

The following table describes how the parts of a version number signify a module’s stability and backward compatibility.

| Version stage 版本阶段                                       | Example  例子                             | Message to developers 给开发者的信息                         |
| ------------------------------------------------------------ | ----------------------------------------- | ------------------------------------------------------------ |
| [In development](https://go.dev/doc/modules/version-numbers#in-development) | Automatic pseudo-version numberv**0**.x.x | Signals that the module is still **in development and unstable**. This release carries no backward compatibility or stability guarantees.标志着该模块仍在开发中，不稳定。这个版本没有向后兼容性或稳定性保证。 |
| [Major version](https://go.dev/doc/modules/version-numbers#major) | v**1**.x.x                                | Signals **backward-incompatible public API changes**. This release carries no guarantee that it will be backward compatible with preceding major versions.主要版本v1.x.x 标志着向后不兼容的公共API变化。该版本不保证与之前的主要版本向后兼容。 |
| [Minor version](https://go.dev/doc/modules/version-numbers#minor) | vx.**4**.x                                | Signals **backward-compatible public API changes**. This release guarantees backward compatibility and stability.次要版本vx.4.x标志着向后兼容的公共API变化。此版本保证向后兼容和稳定。 |
| [Patch version](https://go.dev/doc/modules/version-numbers#patch) | vx.x.**1**                                | Signals **changes that don't affect the module's public API** or its dependencies. This release guarantees backward compatibility and stability.补丁版本vx.x.1 标志着不影响模块的公共API或其依赖关系的变化。这个版本保证了向后的兼容性和稳定性。 |
| [Pre-release version](https://go.dev/doc/modules/version-numbers#pre-release) | vx.x.x-**beta.2**                         | Signals that this is a **pre-release milestone, such as an alpha or beta**. This release carries no stability guarantees.预发布版本 vx.x.x-beta.2 标志着这是一个预发布的里程碑，例如 alpha 或 beta。这个版本没有稳定性保证。 |



## In development 开发中

Signals that the module is still in development and **unstable**. This release carries no backward compatibility or stability guarantees.

标志着该模块仍在开发中，不稳定。这个版本没有向后兼容性或稳定性保证。

The version number can take one of the following forms:

版本号可以采取以下形式之一：

**Pseudo-version number **伪版本号

> v0.0.0-20170915032832-14c0d48ead0c

**v0 number** v0编号

> v0.x.x 



### Pseudo-version number 伪版本号

When a module has not been tagged in its repository, Go tools will generate a pseudo-version number for use in the go.mod file of code that calls functions in the module.

当一个模块在其仓库中没有被标记时，Go工具将生成一个伪版本号，用于调用该模块函数的代码的go.mod文件中。

**Note:** As a best practice, always allow Go tools to generate the pseudo-version number rather than creating your own.

注意: 作为一个最佳实践，总是让Go工具生成伪版本号，而不是自己创建。

Pseudo-versions are useful when a developer of code consuming the module’s functions needs to develop against a commit that hasn’t been tagged with a semantic version tag yet.

当使用模块函数的代码开发人员需要针对尚未被标记为语义版本标签的提交进行开发时，伪版本号非常有用。

A pseudo-version number has three parts separated by dashes, as shown in the following form:

伪版本号有三个部分，用破折号隔开，如以下形式所示：

#### Syntax 语法

*baseVersionPrefix*-*timestamp*-*revisionIdentifier*

#### Parts 部件

- **baseVersionPrefix** (vX.0.0 or vX.Y.Z-0) is a value derived either from a semantic version tag that precedes the revision or from vX.0.0 if there is no such tag.timestamp（yymmddhhmmss）是该修订版的UTC时间。在Git中，这是提交时间，而不是作者时间。
- **timestamp** (yymmddhhmmss) is the UTC time the revision was created. In Git, this is the commit time, not the author time.revisionIdentifier (abcdefabcdef) 是提交哈希值的12个字符的前缀，或者在Subversion中，是一个零填充的版本号。
- **revisionIdentifier** (abcdefabcdef) is a 12-character prefix of the commit hash, or in Subversion, a zero-padded revision number.



### v0 number v0号

A module published with a v0 number will have a formal semantic version number with a major, minor, and patch part, as well as an optional pre-release identifier.

一个以v0编号发布的模块将有一个正式的语义上的版本号，包括主要、次要和补丁部分，以及一个可选的预发布标识符。

Though a v0 version can be used in production, it makes no stability or backward compatibility guarantees. In addition, versions v1 and later are allowed to break backward compatibility for code using the v0 versions. For this reason, a developer with code consuming functions in a v0 module is responsible for adapting to incompatible changes until v1 is released.

尽管v0版本可以在生产中使用，但它不做任何稳定性或向后兼容性的保证。此外，v1及以后的版本允许打破使用v0版本的代码的向后兼容性。由于这个原因，拥有消耗v0模块功能的代码的开发者有责任适应不兼容的变化，直到v1版本发布。



## Pre-release version 预发布版本

Signals that this is a pre-release milestone, such as an alpha or beta. This release carries no stability guarantees.

标志着这是一个预发布的里程碑，如alpha或beta。这个版本没有稳定性保证。

#### Example 例如

```
vx.x.x-beta.2
```

A module’s developer can use a pre-release identifier with any major.minor.patch combination by appending a hyphen and the pre-release identifier.

一个模块的开发者可以在任何major.minor.patch组合中使用预发布标识符，方法是在预发布标识符后面加上一个连字符。



## Minor version 次要版本

Signals backward-compatible changes to the module’s public API. This release guarantees backward compatibility and stability.

标志着模块的公共API向后兼容的变化。这个版本保证了向后的兼容性和稳定性。

#### Example 例如

```
vx.4.x
```

This version changes the module’s public API, but not in a way that breaks calling code. This might include changes to a module’s own dependencies or the addition of new functions, methods, struct fields, or types.

这个版本改变了模块的公共API，但不是以破坏调用代码的方式。这可能包括对模块本身的依赖性的改变，或增加新的函数、方法、结构域或类型。

In other words, this version might include enhancements through new functions that another developer might want to use. However, a developer using previous minor versions needn’t change their code otherwise.

换句话说，这个版本可能包括通过另一个开发者可能想要使用的新函数进行的增强。然而，使用以前次要版本的开发者不需要改变他们的代码。



## Patch version 补丁版本

Signals changes that don’t affect the module’s public API or its dependencies. This release guarantees backward compatibility and stability.

标志着不影响模块的公共API或其依赖关系的变化。这个版本保证了向后的兼容性和稳定性。

#### Example 例如

```
vx.x.1
```

An update that increments this number is only for minor changes such as bug fixes. Developers of consuming code can upgrade to this version safely without needing to change their code.

递增这个数字的更新只针对微小的变化，如错误修复。消耗代码的开发者可以安全地升级到这个版本而不需要改变他们的代码。



## Major version 主要版本

Signals backward-incompatible changes in a module’s public API. This release carries no guarantee that it will be backward compatible with preceding major versions.

标志着一个模块的公共API向后兼容的变化。这个版本不保证它能向后兼容之前的主要版本。

#### Example 例如

v1.x.x

A v1 or above version number signals that the module is stable for use (with exceptions for its pre-release versions).

一个v1或以上的版本号表示该模块可以稳定使用（其预发布版本除外）。

Note that because a version 0 makes no stability or backward compatibility guarantees, a developer upgrading a module from v0 to v1 is responsible for adapting to changes that break backward compatibility.

请注意，由于0版本没有稳定性或向后兼容性的保证，开发者要负责将模块从v0升级到v1，以适应破坏向后兼容性的变化。

A module developer should increment this number past v1 only when necessary because the version upgrade represents significant disruption for developers whose code uses function in the upgraded module. This disruption includes backward-incompatible changes to the public API, as well as the need for developers using the module to update the package path wherever they import packages from the module.

一个模块的开发者只有在必要的时候才应该把这个数字增加到v1以上，因为版本升级对那些代码使用升级后的模块功能的开发者来说意味着重大的破坏。这种干扰包括对公共API的向后兼容的改变，以及使用该模块的开发者需要在他们从该模块导入包的地方更新包路径。

A major version update to a number higher than v1 will also have a new module path. That’s because the module path will have the major version number appended, as in the following example:

一个主要的版本更新到高于v1的数字，也会有一个新的模块路径。这是因为模块路径将附加主要的版本号，如下面的例子：

```
module example.com/mymodule/v2 v2.0.0
```

A major version update makes this a new module with a separate history from the module’s previous version. If you’re developing modules to publish for others, see “Publishing breaking API changes” in [Module release and versioning workflow](https://go.dev/doc/modules/release-workflow).

一个主要版本的更新使得这是一个新的模块，与该模块之前的版本有一个独立的历史。如果你开发的模块要为他人发布，请参阅模块发布和版本工作流程中的 "发布破坏性的API变化"。

For more on the module directive, see [go.mod reference](https://go.dev/doc/modules/gomod-ref).

关于模块指令的更多信息，请参见go.mod参考。