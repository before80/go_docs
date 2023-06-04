+++
title = "go 如何减轻供应链攻击的影响"
weight = 95
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# How Go Mitigates Supply Chain Attacks  - Go如何减轻供应链攻击的影响

https://go.dev/blog/supply-chain

Filippo Valsorda
31 March 2022

Modern software engineering is collaborative, and based on reusing Open Source software. That exposes targets to supply chain attacks, where software projects are attacked by compromising their dependencies.

现代软件工程是协作性的，并以重复使用开源软件为基础。这使目标暴露在供应链攻击之下，即软件项目通过破坏其依赖关系而受到攻击。

Despite any process or technical measure, every dependency is unavoidably a trust relationship. However, the Go tooling and design help mitigate risk at various stages.

尽管有任何流程或技术措施，每一个依赖关系都不可避免地是一种信任关系。然而，Go的工具和设计有助于在各个阶段降低风险。

## All builds are "locked" 所有构建都是 "锁定 "的

There is no way for changes in the outside world—such as a new version of a dependency being published—to automatically affect a Go build.

外部世界的变化--例如依赖关系的新版本被发布--没有办法自动影响 Go 的构建。

Unlike most other package managers files, Go modules don’t have a separate list of constraints and a lock file pinning specific versions. The version of every dependency contributing to any Go build is fully determined by the [`go.mod` file](https://go.dev/ref/mod#go-mod-file) of the main module.

与大多数其他软件包管理器文件不同，Go模块没有单独的约束列表和锁定文件，而是将特定的版本固定下来。对任何 Go 构建作出贡献的每个依赖项的版本完全由主模块的 go.mod 文件决定。

Since Go 1.16, this determinism is enforced by default, and build commands (`go build`, `go test`, `go install`, `go run`, …) [will fail if the go.mod is incomplete](https://go.dev/ref/mod#go-mod-file-updates). The only commands that will change the `go.mod` (and therefore the build) are `go get` and `go mod tidy`. These commands are not expected to be run automatically or in CI, so changes to dependency trees must be made deliberately and have the opportunity to go through code review.

从Go 1.16开始，这种确定性被默认执行，如果go.mod不完整，构建命令（go build, go test, go install, go run, ...）将失败。唯一会改变go.mod（因此也会改变构建）的命令是go get和go mod tidy。这些命令不会自动或在CI中运行，所以对依赖关系树的改变必须有意进行，并有机会通过代码审查。

This is very important for security, because when a CI system or new machine runs `go build`, the checked-in source is the ultimate and complete source of truth for what will get built. There is no way for third parties to affect that.

这对安全非常重要，因为当CI系统或新机器运行go build时，签入的源码是最终的和完整的真相来源，说明什么会被构建。第三方没有办法影响这一点。

Moreover, when a dependency is added with `go get`, its transitive dependencies are added at the version specified in the dependency’s `go.mod` file, not at their latest versions, thanks to [Minimal version selection](https://go.dev/ref/mod#minimal-version-selection). The same happens for invocations of `go install example.com/cmd/devtoolx@latest`, [the equivalents of which in some ecosystems bypass pinning](https://research.swtch.com/npm-colors). In Go, the latest version of `example.com/cmd/devtoolx` will be fetched, but then all the dependencies will be set by its `go.mod` file.

此外，当用go get添加依赖时，由于最小版本的选择，它的交叉依赖会按照依赖的go.mod文件中指定的版本添加，而不是按照它们的最新版本。同样的情况也发生在调用 go install example.com/cmd/devtoolx@latest 的情况下，在某些生态系统中，其等价物会绕过 pinning。在Go中，example.com/cmd/devtoolx的最新版本将被获取，但所有的依赖关系将由其go.mod文件设定。

If a module gets compromised and a new malicious version is published, no one will be affected until they explicitly update that dependency, providing the opportunity to review the changes and time for the ecosystem to detect the event.

如果一个模块被破坏，新的恶意版本被发布，没有人会受到影响，直到他们明确地更新该依赖关系，提供了审查变化的机会和生态系统检测事件的时间。

## Version contents never change 版本内容永不改变

Another key property necessary to ensure third parties can’t affect builds is that the contents of a module version are immutable. If an attacker that compromises a dependency could re-upload an existing version, they could automatically compromise all projects that depend on it.

确保第三方不能影响构建的另一个关键属性是，模块版本的内容是不可改变的。如果一个攻击者破坏了一个依赖关系，可以重新上传一个现有的版本，他们可以自动破坏所有依赖它的项目。

That’s what the [`go.sum` file](https://go.dev/ref/mod#go-sum-files) is for. It contains a list of cryptographic hashes of each dependency that contributes to the build. Again, an incomplete `go.sum` causes an error, and only `go get` and `go mod tidy` will modify it, so any changes to it will accompany a deliberate dependency change. Other builds are guaranteed to have a full set of checksums.

这就是go.sum文件的作用。它包含了对构建有贡献的每个依赖项的加密哈希值的列表。同样，一个不完整的go.sum会导致一个错误，而且只有go get和go mod tidy会修改它，所以对它的任何修改都会伴随着一个故意的依赖项改变。其他的构建被保证有一套完整的校验和。

This is a common feature of most lock files. Go goes beyond it with the [Checksum Database](https://go.dev/ref/mod#checksum-database) (sumdb for short), a global append-only cryptographically-verifiable list of go.sum entries. When `go get` needs to add an entry to the `go.sum` file, it fetches it from the sumdb along with cryptographic proof of the sumdb integrity. This ensures that not only every build of a certain module uses the same dependency contents, but that every module out there uses the same dependency contents!

这是大多数锁文件的一个共同特征。Go通过校验和数据库（简称sumdb）超越了它，它是一个全局性的仅可附加的加密验证的go.sum条目列表。当go get需要在go.sum文件中添加一个条目时，它从sumdb中获取该条目，并对sumdb的完整性进行加密证明。这不仅确保了某一模块的每一次构建都使用相同的依赖内容，而且确保了每一个模块都使用相同的依赖内容。

The sumdb makes it impossible for compromised dependencies or even Google-operated Go infrastructure to target specific dependents with modified (e.g. backdoored) source. You’re guaranteed to be using the exact same code that everyone else who’s using e.g. v1.9.2 of `example.com/modulex` is using and has reviewed.

sumdb使得被破坏的依赖内容，甚至谷歌操作的Go基础设施不可能用修改过的（如反屏蔽）源代码来瞄准特定的依赖内容。保证您使用的代码与其他使用例如example.com/modulex的v1.9.2的人所使用的代码完全一样，并且已经过审查。

Finally, my favorite features of the sumdb: it doesn’t require any key management on the part of module authors, and it works seamlessly with the decentralized nature of Go modules.

最后，我最喜欢sumdb的特点：它不需要模块作者的任何密钥管理，而且它与Go模块的去中心化特性无缝连接。

## The VCS is the source of truth - VCS是真理之源

Most projects are developed through some version control system (VCS) and then, in other ecosystems, uploaded to the package repository. This means there are two accounts that could be compromised, the VCS host and the package repository, the latter of which is used more rarely and more likely to be overlooked. It also means it’s easier to hide malicious code in the version uploaded to the repository, especially if the source is routinely modified as part of the upload, for example to minimize it.

大多数项目是通过一些版本控制系统（VCS）开发的，然后在其他生态系统中，上传到软件包库。这意味着有两个账户可能被入侵，即VCS主机和软件包库，后者使用得更少，更容易被忽视。这也意味着在上传到仓库的版本中更容易隐藏恶意代码，尤其是当源码作为上传的一部分被例行修改时，例如将其最小化。

In Go, there is no such thing as a package repository account. The import path of a package embeds the information that `go mod download` [needs in order to fetch its module](https://pkg.go.dev/cmd/go#hdr-Remote_import_paths) directly from the VCS, where tags define versions.

在Go中，不存在所谓的包库账户。包的导入路径嵌入了go mod download所需要的信息，以便直接从VCS中获取其模块，其中标签定义了版本。

We do have the [Go Module Mirror](https://go.dev/blog/module-mirror-launch), but that’s only a proxy. Module authors don’t register an account and don’t upload versions to the proxy. The proxy uses the same logic that the `go` tool uses (in fact, the proxy runs `go mod download`) to fetch and cache a version. Since the Checksum Database guarantees that there can be only one source tree for a given module version, everyone using the proxy will see the same result as everyone bypassing it and fetching directly from the VCS. (If the version is not available anymore in the VCS or if its contents changed, fetching directly will lead to an error, while fetching from the proxy might still work, improving availability and protecting the ecosystem from ["left-pad" issues](https://blog.npmjs.org/post/141577284765/kik-left-pad-and-npm).)

我们确实有Go Module Mirror，但那只是一个代理。模块作者不需要注册账户，也不需要向代理上传版本。代理使用与go工具相同的逻辑（事实上，代理运行go模块下载）来获取和缓存一个版本。由于校验数据库保证一个给定的模块版本只能有一个源树，每个使用代理的人都会看到与绕过代理直接从VCS获取的结果相同。(如果该版本在VCS中不再可用，或者其内容发生了变化，直接获取将导致错误，而从代理获取可能仍然有效，提高了可用性并保护生态系统免受 "左键 "问题的影响）。

Running VCS tools on the client exposes a pretty large attack surface. That’s another place the Go Module Mirror helps: the `go` tool on the proxy runs inside a robust sandbox and is configured to support every VCS tool, while [the default is to only support the two major VCS systems](https://go.dev/ref/mod#vcs-govcs) (git and Mercurial). Anyone using the proxy can still fetch code published using off-by-default VCS systems, but attackers can’t reach that code in most installations.

在客户端运行VCS工具会暴露出一个相当大的攻击面。这是Go模块镜像的另一个作用：代理上的Go工具在一个强大的沙盒内运行，并被配置为支持所有的VCS工具，而默认情况下只支持两个主要的VCS系统（git和Mercurial）。任何使用代理的人仍然可以获取使用非默认的VCS系统发布的代码，但攻击者在大多数安装中无法接触到这些代码。

## Building code doesn’t execute it 构建代码并不执行它

It is an explicit security design goal of the Go toolchain that neither fetching nor building code will let that code execute, even if it is untrusted and malicious. This is different from most other ecosystems, many of which have first-class support for running code at package fetch time. These "post-install" hooks have been used in the past as the most convenient way to turn a compromised dependency into compromised developer machines, and to [worm](https://en.wikipedia.org/wiki/Computer_worm) through module authors.

Go工具链的一个明确的安全设计目标是，无论是获取还是构建代码，都不会让该代码执行，即使它是不被信任的和恶意的。这与其他大多数生态系统不同，许多生态系统在获取软件包时对运行代码有一流的支持。这些 "安装后 "的钩子在过去被用作最方便的方式，将受影响的依赖关系变成受影响的开发者机器，并通过模块作者进行蠕虫攻击。

To be fair, if you’re fetching some code it’s often to execute it shortly afterwards, either as part of tests on a developer machine or as part of a binary in production, so lacking post-install hooks is only going to slow down attackers. (There is no security boundary within a build: any package that contributes to a build can define an `init` function.) However, it can be a meaningful risk mitigation, since you might be executing a binary or testing a package that only uses a subset of the module’s dependencies. For example, if you build and execute `example.com/cmd/devtoolx` on macOS there is no way for a Windows-only dependency or a dependency of `example.com/cmd/othertool` to compromise your machine.

公平地说，如果您要获取一些代码，往往会在不久之后执行，要么作为开发者机器上的测试的一部分，要么作为生产中的二进制文件的一部分，所以缺乏安装后钩子只会减缓攻击者的速度。(在构建过程中没有安全边界：任何有助于构建的软件包都可以定义一个初始函数）。然而，它可以成为一个有意义的风险缓解措施，因为您可能正在执行一个二进制文件或测试一个包，而这个包只使用模块依赖的一个子集。例如，如果您在macOS上构建并执行example.com/cmd/devtoolx，那么只有Windows的依赖或example.com/cmd/othertool的依赖就不可能危害到您的机器。

In Go, modules that don’t contribute code to a specific build have no security impact on it.

在Go中，不为特定构建贡献代码的模块对其没有安全影响。

## "A little copying is better than a little dependency" "一点复制比一点依赖项好"

The final and maybe most important software supply chain risk mitigation in the Go ecosystem is the least technical one: Go has a culture of rejecting large dependency trees, and of preferring a bit of copying to adding a new dependency. It goes all the way back to one of the Go proverbs: ["a little copying is better than a little dependency"](https://youtube.com/clip/UgkxWCEmMJFW0-TvSMzcMEAHZcpt2FsVXP65). The label "zero dependencies" is proudly worn by high-quality reusable Go modules. If you find yourself in need of a library, you’re likely to find it will not cause you to take on a dependency on dozens of other modules by other authors and owners.

在Go生态系统中，最后一个也许也是最重要的软件供应链风险缓解措施是最没有技术含量的一个。Go有一种拒绝大型依赖树的文化，宁可复制一点也不愿意添加新的依赖关系。这可以追溯到Go的一句谚语。"一点复制比一点依赖项好"。"零依赖 "的标签被高质量的可重复使用的Go模块所自豪地佩戴。如果您发现自己需要一个库，您很可能会发现它不会导致您依赖其他作者和所有者的几十个模块。

That’s enabled also by the rich standard library and additional modules (the `golang.org/x/...` ones), which provide commonly used high-level building blocks such as an HTTP stack, a TLS library, JSON encoding, etc.

丰富的标准库和附加模块（golang.org/x/...的模块）也使之成为可能，这些模块提供了常用的高级构建模块，如HTTP栈、TLS库、JSON编码等。

All together this means it’s possible to build rich, complex applications with just a handful of dependencies. No matter how good the tooling is, it can’t eliminate the risk involved in reusing code, so the strongest mitigation will always be a small dependency tree.

所有这些意味着只需少量的依赖关系就可以建立丰富、复杂的应用程序。无论工具有多好，它都不能消除重复使用代码的风险，所以最有力的缓解措施永远是一个小的依赖树。
