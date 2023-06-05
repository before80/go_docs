+++
title = "验证模块"
date = 2023-05-17T09:59:21+08:00
weight = 14
description = ""
isCJKLanguage = true
draft = false
+++
## Authenticating modules 验证模块

> 原文：[https://go.dev/ref/mod#authenticating](https://go.dev/ref/mod#authenticating)

​	当`go`命令下载模块[zip 文件](../ModuleZipFiles)或`go.mod`文件到[module cache（模块缓存）](../ModuleCache)时，它会计算加密散列并将其与已知值进行比较，以验证该文件自首次下载以来没有更改。如果下载的文件没有正确的散列值，`go`命令会报告安全错误。

​	对于`go.mod`文件，`go`命令从文件内容中计算出散列值。对于模块zip文件，`go`命令以确定的顺序从归档文件中的文件名称和内容计算散列值。散列值不受文件顺序、压缩、对齐和其他元数据的影响。有关散列值的实现细节，请参见 [golang.org/x/mod/sumdb/dirhash](https://pkg.go.dev/golang.org/x/mod/sumdb/dirhash?tab=doc)。

​	`go`命令将每个散列值与主模块的`go.sum`文件中的对应行进行比较。如果散列值与`go.sum`中的散列值不同，`go`命令将报告安全错误，并删除下载的文件而不将其加入模块缓存。

​	如果`go.sum`文件不存在，或者它不包含下载文件的散列值，则`go`命令可以使用校验和数据库（公共可用模块的全局散列源）验证散列值。散列值被验证之后，`go`命令会将其加入`go.sum`，并将下载的文件加入模块缓存中。如果模块是私有的（被`GOPRIVATE`或`GONOSUMDB`环境变量匹配），或者如果校验和数据库被禁用（通过设置`GOSUMDB=off`），则`go`命令接受散列值并将文件添加到模块缓存中，而不进行验证。

​	模块缓存通常由系统中的所有Go项目共享，每个模块可能有自己的`go.sum`文件，其散列值可能不同。为了避免信任其他模块，`go`命令在访问模块缓存中的文件时，使用主模块的`go.sum`来验证散列值。zip文件的散列值计算成本很高，所以`go`命令检查与zip文件一起存储的预先计算的散列值，而不是重新散列文件。[go mod verify](../gomodFiles#go-mod-verify)命令可以用来检查zip文件和提取的目录自从被添加到模块缓存后是否被修改过。

### go.sum files

​	模块的根目录中可能有一个名为`go.sum`的文本文件，以及跟随它的`go.mod`文件。`go.sum`文件包含模块的直接和间接依赖的加密散列值。当`go`命令将模块`.mod`或`.zip`文件下载到[module cache（模块缓存）](../ModuleCache)中时，它会计算散列并检查该散列是否与主模块的`go.sum`文件中的相应散列匹配。如果模块没有依赖项，或者如果使用[replace 指令](../gomodFiles#replace-directive)将所有依赖项替换为本地目录，那么`go.sum`可能为空或不存在。

​	`go.sum`中的每一行都有三个用空格隔开的字段：模块路径，版本（可能以`/go.mod`结尾），以及散列值。

- 模块路径是散列值所属模块的名称。

- 版本是散列值所属的模块的版本。如果版本以`/go.mod`结尾，则散列值仅适用于该模块的`go.mod`文件；否则，散列值适用于该模块的`.zip`文件中的文件。

- 散列值由算法名称（如`h1`）和一个base64编码的加密散列组成，用冒号（`:`）分开。目前，SHA-256（`h1`）是唯一支持的散列算法。如果将来发现了SHA-256的漏洞，将增加对另一种算法（命名为`h2`，以此类推）的支持。

  

​	`go.sum`文件可能包含一个模块的多个版本的散列值。`go`命令可能需要从依赖项的多个版本加载`go.mod`文件，以进行[最小的版本选择](../MVS)。`go.sum`也可能包含不再需要的模块版本的散列值（例如，在升级之后）。[go mod tidy](../gomodFiles#go-mod-tidy)将添加缺失的散列值，并从`go.sum`中删除不必要的散列值。

### Checksum database 校验和数据库

​	校验和数据库是`go.sum`行的全局源。`go`命令可以在许多情况下使用它来检测代理或源服务器的错误行为。

​	校验和数据库允许所有公开可用的模块版本的全局一致性和可靠性。它使不被信任的代理成为可能，因为它们无法在不被发现的情况下提供错误的代码。它还确保与特定版本相关的bit位不会从一天到另一天发生变化，即使模块的作者后来更改了其存储库中的标签。

​	校验和数据库由Google运营的[sum.golang.org](https://sum.golang.org/)提供服务。它是由[Trillian](https://github.com/google/trillian)支持的`go.sum`行散列的透明日志(或"`Merkle 树`")。  `Merkle 树`的主要优点是，独立的审计人员可以验证它没有被篡改，因此它比简单的数据库更值得信赖。

​	`go`命令使用最初在[Proposal: Secure the Public Go Module Ecosystem](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#checksum-database)中概述的协议与校验和数据库进行交互。

​	下表列出了校验和数据库必须响应的查询。对于每个路径，`$base`是校验和数据库URL的路径部分，`$module`是模块路径，而`$version`是版本。例如，如果校验和数据库的URL是`https://sum.golang.org`，而客户端请求的是`golang.org/x/text`模块的记录，版本为`v0.3.2`，那么客户端将发送`https://sum.golang.org/lookup/golang.org/x/text@v0.3.2`的`GET`请求。

​	为了避免在从不区分大小写的文件系统中提供服务时出现歧义，`$module`和`$version`元素被进行了[case-encoded（大小写编码）](https://pkg.go.dev/golang.org/x/mod/module#EscapePath)，将每个大写字母替换为感叹号，后跟相应的小写字母。这允许模块`example.com/M`和`example.com/m`同时被存储在磁盘上，因为前者被编码为`example.com/!m`。

​	路径中由方括号包围的部分，如`[.p/$W]`表示可选的值。

| Path | Description |
| ---- | ----------- |
|      |             |

#### `$base/latest`

​	返回最新日志的有符号编码的树描述。这个签名的描述是以[note（注释）](https://pkg.go.dev/golang.org/x/mod/sumdb/note)的形式出现的，它是由一个或多个服务器密钥签名的文本，可以用服务器的公钥进行验证。树描述提供了树的大小和该大小的树头的散列值。这中编码在`golang.org/x/mod/sumdb/tlog#FormatTree`中有描述。



#### `$base/lookup/$module@$version`

​	返回`$version`中关于`$module`的条目的日志记录编号，后跟记录的数据（即`$version`中的`$module`的`go.sum`行）和包含该记录的有符号编码的树描述。



#### `$base/tile/$H/$L/$K[.p/$W]`

​	返回一个[log tile](https://research.swtch.com/tlog#serving_tiles)，它是组成日志的一部分的一组散列。每个tile定义在tile级别`$L`，左起`$K`的二维坐标中，tile高度为`$H`。可选的`.p/$W`后缀表示只有`$W`散列的部分日志块。如果没有找到部分tile，客户必须回退到获取完整的tile。



#### `$base/tile/$H/data/$K[.p/$W]`

​	返回`/tile/$H/0/$K[.p/$W]`中的叶子散列的记录数据（带有字面`data`路径元素）。

--------------------------------------------------------------------





​	如果`go`命令查询校验和数据库，那么第一步就是通过`/lookup`端点检索记录数据。如果日志中尚未记录模块版本，校验和数据库将尝试在回复前从原服务器上获取它。这个`/lookup`数据提供了这个模块版本的sum（和）以及它在日志中的位置，这就通知了客户端应该取哪些tile来执行证明。在向主模块的`go.sum`文件添加新的`go.sum`行之前，`go`命令执行 "inclusion（包含性）"证明（即日志中存在特定记录）和 "consistency（一致性）"证明（即树没有被篡改）。重要的是，如果没有首先根据签名的树散列值对其进行身份认证，以及根据客户的签名树散列值时间线对签名树散列值进行认证，就不应该使用来自`/lookup`的数据。

​	已签名的树状散列值和由校验数据库提供的新瓦片都存储在模块缓存中，所以`go`命令只需要获取缺少的瓦片。

`go`命令不需要直接连接到校验和数据库。它可以通过镜像[mirrors the checksum database（校验和数据库）](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#proxying-a-checksum-database)并支持上述协议的模块代理来请求模块和。这对于阻止组织外部请求的私有企业代理特别有帮助。

​	`GOSUMDB`环境变量标识要使用的校验和数据库的名称，还可以选择其公钥和URL，如：

```
GOSUMDB="sum.golang.org"
GOSUMDB="sum.golang.org+<publickey>"
GOSUMDB="sum.golang.org+<publickey> https://sum.golang.org"
```

​	`go`命令知道`sum.golang.org`的公钥，也知道`sum.golang.google.cn`这个名字（在中国大陆可用）连接到`sum.golang.org`的校验和数据库；使用其他数据库需要显式给出公钥。URL默认为`https://`，后跟数据库名称。

​	`GOSUMDB`默认为`sum.golang.org`，这是由Google运营的Go校验和数据库。有关该服务的隐私政策，请参见 [https://sum.golang.org/privacy](https://sum.golang.org/privacy)。

​	如果 `GOSUMDB` 被设置为`off`，或者在调用 `go get` 时使用了 `-insecure` 标志，则不会查询校验和数据库，并接受所有未识别的模块，代价是放弃了对所有模块的可重复下载的安全保证。一个更好的方法是使用`GOPRIVATE`或`GONOSUMDB`环境变量来绕过特定模块的校验和数据库。详见[Private Modules（私有模块）](../PrivateModules)。

​	`go env -w`命令可用于为将来的`go`命令调用[设置这些变量]({{< ref "/cmd/go#print-go-environment-information">}})。