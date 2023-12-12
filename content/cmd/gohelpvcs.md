+++
title = "go help vcs"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

The 'go get' command can run version control commands like git to download imported code. This functionality is critical to the decentralized Go package ecosystem, in which code can be imported from any server, but it is also a potential security problem, if a malicious server finds a way to cause the invoked version control command to run unintended code.

​	`go get` 命令可以运行版本控制命令（如 git）来下载导入的代码。这个功能对于去中心化的 Go 包生态系统至关重要，其中的代码可以从任何服务器导入，但这也是一个潜在的安全问题，如果恶意服务器找到一种方式来使调用的版本控制命令运行意外的代码。

To balance the functionality and security concerns, the 'go get' command by default will only use git and hg to download code from public servers. But it will use any known version control system (bzr, fossil, git, hg, svn) to download code from private servers, defined as those hosting packages matching the GOPRIVATE variable (see 'go help private'). The rationale behind allowing only Git and Mercurial is that these two systems have had the most attention to issues of being run as clients of untrusted servers. In contrast, Bazaar, Fossil, and Subversion have primarily been used in trusted, authenticated environments and are not as well scrutinized as attack surfaces.

​	为了平衡功能和安全性的考虑，`go get` 命令默认仅使用 git 和 hg 从公共服务器下载代码。但它将使用任何已知的版本控制系统（bzr、fossil、git、hg、svn）从私有服务器下载代码，私有服务器定义为托管与 GOPRIVATE 变量匹配的包的服务器（参见 'go help private'）。允许仅使用 Git 和 Mercurial 的原因是这两个系统在作为不受信任服务器的客户端运行时已经受到了最多的关注。相比之下，Bazaar、Fossil 和 Subversion 主要在受信任的身份验证环境中使用，并且作为攻击面没有得到同样深入的审查。

The version control command restrictions only apply when using direct version control access to download code. When downloading modules from a proxy, 'go get' uses the proxy protocol instead, which is always permitted.

​	版本控制命令的限制仅在使用直接版本控制访问下载代码时才适用。在从代理下载模块时，`go get` 使用代理协议，这是始终允许的。

By default, the 'go get' command uses the Go module mirror (proxy.golang.org) for public packages and only falls back to version control for private packages or when the mirror refuses to serve a public package (typically for legal reasons). Therefore, clients can still access public code served from Bazaar, Fossil, or Subversion repositories by default, because those downloads use the Go module mirror, which takes on the security risk of running the version control commands using a custom sandbox.

​	默认情况下，`go get` 命令使用 Go 模块镜像（proxy.golang.org）来获取公共包，并仅在对私有包进行版本控制或镜像拒绝提供公共包时才回退到版本控制（通常是由于法律原因）。因此，客户端仍然可以默认访问从 Bazaar、Fossil 或 Subversion 存储库提供的公共代码，因为这些下载使用 Go 模块镜像，该镜像承担了使用自定义沙箱运行版本控制命令的安全风险。

The GOVCS variable can be used to change the allowed version control systems for specific packages (identified by a module or import path).

​	GOVCS 变量可用于更改特定包（由模块或导入路径标识）允许的版本控制系统。

The GOVCS variable applies when building package in both module-aware mode and GOPATH mode. When using modules, the patterns match against the module path. When using GOPATH, the patterns match against the import path corresponding to the root of the version control repository.

​	GOVCS 变量在构建模块感知模式和 GOPATH 模式下均适用。使用模块时，模式与模块路径匹配。在使用 GOPATH 时，模式与对应于版本控制存储库根的导入路径匹配。

The general form of the GOVCS setting is a comma-separated list of pattern:vcslist rules. The pattern is a glob pattern that must match one or more leading elements of the module or import path. The vcslist is a pipe-separated list of allowed version control commands, or "all" to allow use of any known command, or "off" to disallow all commands. Note that if a module matches a pattern with vcslist "off", it may still be downloaded if the origin server uses the "mod" scheme, which instructs the go command to download the module using the GOPROXY protocol. The earliest matching pattern in the list applies, even if later patterns might also match.

​	GOVCS 设置的一般形式是一个逗号分隔的 pattern:vcslist 规则列表。模式是一个 glob 模式，必须与模块或导入路径的一个或多个前导元素匹配。vcslist 是一个用管道分隔的允许的版本控制命令列表，或者是 "all"，允许使用任何已知命令，或者是 "off"，禁止所有命令。请注意，如果模块与 vcslist "off" 的模式匹配，仍然可以下载，如果源服务器使用 "mod" 方案，该方案会指示 go 命令使用 GOPROXY 协议下载模块。列表中的最早匹配模式适用，即使后来的模式可能也匹配。

For example, consider:

​	例如，考虑：

        GOVCS=github.com:git,evil.com:off,*:git|hg

With this setting, code with a module or import path beginning with github.com/ can only use git; paths on evil.com cannot use any version control command, and all other paths (* matches everything) can use only git or hg.

​	使用此设置，以 github.com/ 开头的模块或导入路径只能使用 git；在 evil.com 上的路径不能使用任何版本控制命令，而所有其他路径（* 匹配所有）只能使用 git 或 hg。

The special patterns "public" and "private" match public and private module or import paths. A path is private if it matches the GOPRIVATE variable; otherwise it is public.

​	特殊模式 "public" 和 "private" 分别匹配公共和私有模块或导入路径。如果路径匹配 GOPRIVATE 变量，则路径为私有；否则为公共。

If no rules in the GOVCS variable match a particular module or import path, the 'go get' command applies its default rule, which can now be summarized in GOVCS notation as 'public:git|hg,private:all'.

​	如果 GOVCS 变量中没有规则与特定的模块或导入路径匹配，则 `go get` 命令将应用其默认规则，现在可以用 GOVCS 表示为 'public:git|hg,private:all'。

To allow unfettered use of any version control system for any package, use:

​	要允许对任何包使用任何版本控制系统，请使用：

        GOVCS=*:all

To disable all use of version control, use:

​	要禁用所有版本控制的使用，请使用：

        GOVCS=*:off

The 'go env -w' command (see 'go help env') can be used to set the GOVCS variable for future go command invocations.

​	`go env -w` 命令（参见 'go help env'）可用于设置 GOVCS 变量，以供将来的 go 命令调用使用。
