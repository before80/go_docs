+++
title = "模块版本编号"
weight = 6
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false

+++
# Module version numbering - 模块版本编号

> 原文：[https://go.dev/doc/modules/version-numbers](https://go.dev/doc/modules/version-numbers)

​	模块的开发者使用模块版本号的每个部分来表示版本的稳定性和向后兼容性。对于每个新版本，模块的发布版本号特别反映了自上一个版本以来模块的变化性质。

​	当您开发使用外部模块的代码时，您可以使用版本号来了解在考虑升级时外部模块的稳定性。当您开发自己的模块时，您的版本号将向其他开发者表明您的模块的稳定性和向后兼容性。

​	本主题描述了模块版本号的含义。

另请参阅：

- 当您在代码中使用外部包时，您可以使用Go工具来管理这些依赖项。更多信息，请参阅[管理依赖项]({{< ref "/docs/UsingAndUnderstandingGo/ManagingDependencies" >}})。
- 如果您正在开发供他人使用的模块，您可以在发布模块时应用版本号，并在其存储库中为该模块打标签。更多信息，请参阅[发布模块](../PublishingAModule)。

​	一个发布的模块以语义化版本模型的版本号发布，如下图所示：

![Diagram illustrating a semantic version number showing major version 1, minor version 4, patch version 0, and pre-release version beta 2](ModuleVersionNnumbering_img/version-number.png)

Diagram illustrating a semantic version number showing major version 1, minor version 4, patch version 0, and pre-release version beta 2

图示语义化版本号，显示主版本号1、次版本号4、补丁版本号0和预发布版本号beta 2

​	以下表格描述了版本号的各个部分如何表示一个模块的稳定性和向后兼容性。

| Version stage 版本阶段    | 示例                          | 给开发者的信息                                               |
| ------------------------- | ----------------------------- | ------------------------------------------------------------ |
| [开发中](#开发中)         | 自动生成的伪版本号为**0**.x.x | 表明该模块仍处于**开发中且不稳定**阶段。此发布版本不提供向后兼容性或稳定性的保证。 |
| [主要版本](#主要版本)     | v**1**.x.x                    | 表明**向后不兼容的公共API变更**。此发布版本不保证与之前的主版本兼容。 |
| [次要版本](#次要版本)     | vx.**4**.x                    | 表明**向后兼容的公共API变更**。此发布版本保证向后兼容性和稳定性。 |
| [补丁版本](#补丁版本)     | vx.x.**1**                    | 表明**对该模块的公共API或其依赖项没有影响的变更**。此发布版本保证向后兼容性和稳定性。 |
| [预发布版本](#预发布版本) | vx.x.x-**beta.2**             | 表明这是一个**预发布里程碑，例如alpha或beta版本**。此发布版本不提供稳定性的保证。 |



## 开发中

​	表明该模块仍处于开发阶段，**不稳定**。此版本不提供向后兼容性或稳定性保证。

​	版本号可以采用以下形式之一：

**伪版本号**

```
v0.0.0-20170915032832-14c0d48ead0c
```

**v0编号** 

```
v0.x.x 
```



### 伪版本号

​	当一个模块在其存储库中没有被标记时，Go工具将为调用该模块函数的代码生成一个伪版本号，用于go.mod文件。

**注意：**作为最佳实践，始终允许Go工具生成伪版本号，而不是自己创建。

​	当使用模块函数的代码开发人员需要针对尚未用语义化版本标签标记的提交进行开发时，伪版本号非常有用。

​	伪版本号由三部分组成，以破折号分隔，如下所示：

#### 语法

```
baseVersionPrefix-timestamp-revisionIdentifier
```



#### 组成部分

- **baseVersionPrefix**（vX.0.0或vX.Y.Z-0）是从修订版本之前的语义化版本标签或vX.0.0（如果没有这样的标签）中派生的值。
- **timestamp**（yymmddhhmmss）是修订版本的创建时间，采用UTC时间。在Git中，这是提交时间，而不是作者时间。
- **revisionIdentifier**（abcdefabcdef）是该提交哈希的12个字符前缀，或者在Subversion中，是填充为零的修订号。



### v0编号

​	以v0数字发布的模块将具有正式的语义化版本号，包括主版本、次版本和修订版本部分，以及可选的预发布标识符。

​	尽管v0版本可以在生产环境中使用，但它不提供稳定性或向后兼容性的保证。此外，v1及以后的版本允许打破使用v0版本的代码的向后兼容性。因此，使用v0模块函数的开发者需要适应不兼容的变更，直到v1版本发布为止。

## 预发布版本

​	表明这是一个预发布的里程碑，比如alpha或beta版本。此版本不提供稳定性保证。

#### 示例

```
vx.x.x-beta.2
```

​	模块的开发者可以通过在`主版本.次版本.修订版本（major.minor.patch）`组合后添加连字符（hyphen ）和预发布标识符来使用预发布标识符。



## 次要版本

​	表明对模块的公共API进行了向后兼容的变更。此版本保证向后兼容性和稳定性。

#### 示例

```
vx.4.x
```

​	此版本变更了模块的公共API，但不会破坏调用代码。这可能包括对模块自身的依赖项的变更或添加新的函数、方法、结构字段或类型。

​	换句话说，此版本可能通过新的函数包含其他开发者可能想要使用的增强功能。然而，使用先前的次版本的开发者无需修改他们的代码。

## 补丁版本

​	表明对该模块的公共API或其依赖项没有影响的变更。此版本保证向后兼容性和稳定性。

#### 示例

```
vx.x.1
```

​	此数字的递增仅用于次要变更（minor changes），例如bug修复。使用此代码的开发者可以安全地升级到此版本，而无需修改他们的代码。

## 主要版本

​	表明对模块的公共API进行了不向后兼容的变更。此版本不保证它能向后兼容之前的主要版本。

#### 示例

```
v1.x.x
```

​	v1或更高版本号表示该模块已经稳定可用（预发布版本除外）。

​	请注意，因为版本0不提供稳定性或向后兼容性的保证，将模块从v0升级到v1的开发者需要适应破坏向后兼容性的变更。

​	模块开发者应仅在必要时将此数字增加到v1之后，因为版本升级对于使用升级后模块中的函数的开发者来说代表着重大的变动。此变动包括对公共API的不兼容变更，以及使用该模块的开发者需要更新导入该模块中包的包路径。

​	大于v1的主要版本更新还将具有新的模块路径。这是因为模块路径将附加主要版本号，如以下示例所示：

```
module example.com/mymodule/v2 v2.0.0
```

A major version update makes this a new module with a separate history from the module’s previous version. If you’re developing modules to publish for others, see "Publishing breaking API changes" in [Module release and versioning workflow](https://go.dev/doc/modules/release-workflow).

​	主要版本更新将使其成为一个新的模块，与模块的先前版本有着不同的历史记录。如果您正在开发用于他人使用（原文中使用 publish ，我想应该是要表达工他人使用的意思才对吧！）的模块，请参阅[模块发布和版本控制工作流程](../ModuleReleaseAndVersioningWorkflow)中的"发布破坏性的API变更"。

​	有关模块指令的更多信息，请参阅[go.mod参考]({{< ref "/docs/References/gomodFileReference">}})。