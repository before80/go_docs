+++
title = "go 中的指令PATH安全"
weight = 99
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Command PATH security in Go - go 中的指令PATH安全

[https://go.dev/blog/path-security](https://go.dev/blog/path-security)

Russ Cox
19 January 2021

Today’s [Go security release](https://go.dev/s/go-security-release-jan-2021) fixes an issue involving PATH lookups in untrusted directories that can lead to remote execution during the `go` `get` command. We expect people to have questions about what exactly this means and whether they might have issues in their own programs. This post details the bug, the fixes we have applied, how to decide whether your own programs are vulnerable to similar problems, and what you can do if they are.

​	今天的[Go安全版本](https://go.dev/s/go-security-release-jan-2021)修复了一个涉及不受信任目录中PATH查找的问题，该问题可能导致在`go get`命令中的远程执行。我们预计人们会有疑问，这到底意味着什么，以及他们自己的程序中是否会有问题。这篇文章详细介绍了这个bug、我们应用的修复程序、如何判断您自己的程序是否容易受到类似问题的影响，以及如果有的话您可以做什么。

## Go command & remote execution - Go命令和远程执行

One of the design goals for the `go` command is that most commands – including `go` `build`, `go` `doc`, `go` `get`, `go` `install`, and `go` `list` – do not run arbitrary code downloaded from the internet. There are a few obvious exceptions: clearly `go` `run`, `go` `test`, and `go` `generate` *do* run arbitrary code – that’s their job. But the others must not, for a variety of reasons including reproducible builds and security. So when `go` `get` can be tricked into executing arbitrary code, we consider that a security bug.

​	Go命令的设计目标之一是大多数命令 —— 包括`go build`、`go doc`、`go get`、`go install`和`go list` —— 不会运行从互联网下载的任意代码。有几个明显的例外：显然，`go run`、`go test`和`go generate`会运行任意代码 —— 那是它们的工作。但是其他的命令不能，因为有各种原因，包括可重复构建和安全。因此，当`go get`可以被欺骗执行任意代码时，我们认为这是一个安全漏洞。

If `go` `get` must not run arbitrary code, then unfortunately that means all the programs it invokes, such as compilers and version control systems, are also inside the security perimeter. For example, we’ve had issues in the past in which clever use of obscure compiler features or remote execution bugs in version control systems became remote execution bugs in Go. (On that note, Go 1.16 aims to improve the situation by introducing a GOVCS setting that allows configuration of exactly which version control systems are allowed and when.)

​	如果`go get`不能运行任意代码，那么不幸的是，这意味着它调用的所有程序，如编译器和版本控制系统，也在安全范围之内。例如，我们过去曾遇到过这样的问题：巧妙地使用晦涩的编译器功能或版本控制系统中的远程执行bug成为Go中的远程执行bug。(关于这一点，Go 1.16旨在通过引入`GOVCS`设置来改善这种情况，该设置允许配置到底允许哪些版本控制系统以及何时允许)。

Today’s bug, however, was entirely our fault, not a bug or obscure feature of `gcc` or `git`. The bug involves how Go and other programs find other executables, so we need to spend a little time looking at that before we can get to the details.

然而，今天的错误完全是我们的错，而不是gcc或git的错误或晦涩的功能。这个bug涉及到Go和其他程序如何寻找其他可执行文件的问题，所以我们需要花一点时间看看这个问题，然后再来讨论细节。

## Commands and PATHs and Go 命令和 PATHs 与 Go

All operating systems have a concept of an executable path (`$PATH` on Unix, `%PATH%` on Windows; for simplicity, we’ll just use the term PATH), which is a list of directories. When you type a command into a shell prompt, the shell looks in each of the listed directories, in turn, for an executable with the name you typed. It runs the first one it finds, or it prints a message like "command not found."

所有的操作系统都有一个可执行路径的概念（Unix的$PATH，Windows的%PATH%；为了简单起见，我们只使用PATH这个术语），它是一个目录列表。当您在shell提示符下输入一条命令时，shell会依次在列出的每个目录中寻找与您输入的名称相同的可执行文件。它运行它找到的第一个可执行文件，或者打印出一个类似 "未找到命令 "的信息。

On Unix, this idea first appeared in Seventh Edition Unix’s Bourne shell (1979). The manual explained:

在Unix中，这个想法首次出现在第七版Unix的Bourne shell中（1979年）。该手册解释说：

> The shell parameter `$PATH` defines the search path for the directory containing the command. Each alternative directory name is separated by a colon (`:`). The default path is `:/bin:/usr/bin`. If the command name contains a / then the search path is not used. Otherwise, each directory in the path is searched for an executable file.
>
> shell参数$PATH定义了包含命令的目录的搜索路径。每个备选的目录名都用冒号（:）隔开。默认路径是:/bin:/usr/bin。如果命令名包含一个/，那么搜索路径就不会被使用。否则，路径中的每个目录都会被搜索到可执行文件。

Note the default: the current directory (denoted here by an empty string, but let’s call it "dot") is listed ahead of `/bin` and `/usr/bin`. MS-DOS and then Windows chose to hard-code that behavior: on those systems, dot is always searched first, automatically, before considering any directories listed in `%PATH%`.

注意默认情况：当前目录（这里用一个空字符串表示，但让我们称它为 "dot"）被列在/bin和/usr/bin之前。MS-DOS和后来的Windows选择了硬编码的行为：在这些系统上，dot总是被自动首先搜索，然后再考虑%PATH%中列出的任何目录。

As Grampp and Morris pointed out in their classic paper "[UNIX Operating System Security](https://people.engr.ncsu.edu/gjin2/Classes/246/Spring2019/Security.pdf)" (1984), placing dot ahead of system directories in the PATH means that if you `cd` into a directory and run `ls`, you might get a malicious copy from that directory instead of the system utility. And if you can trick a system administrator to run `ls` in your home directory while logged in as `root`, then you can run any code you want. Because of this problem and others like it, essentially all modern Unix distributions set a new user’s default PATH to exclude dot. But Windows systems continue to search dot first, no matter what PATH says.

正如Grampp和Morris在他们的经典论文 "UNIX操作系统安全"（1984）中指出的那样，在PATH中把dot放在系统目录之前意味着如果您cd到一个目录并运行ls，您可能会从该目录中得到一个恶意的拷贝而不是系统工具。如果您能欺骗系统管理员在您的主目录中运行ls，同时以root身份登录，那么您就可以运行任何您想要的代码。由于这个问题和其他类似的问题，基本上所有现代Unix发行版都将新用户的默认PATH设置为不包括dot。但Windows系统仍然会先搜索dot，不管PATH怎么说。

For example, when you type the command

例如，当您输入命令

```shell linenums="1"
go version
```

on a typically-configured Unix, the shell runs a `go` executable from a system directory in your PATH. But when you type that command on Windows, `cmd.exe` checks dot first. If `.\go.exe` (or `.\go.bat` or many other choices) exists, `cmd.exe` runs that executable, not one from your PATH.

时，shell会从您的PATH中的一个系统目录中运行一个go可执行文件。但当您在Windows上键入该命令时，cmd.exe首先检查dot。如果.\go.exe（或.\go.bat或许多其他选择）存在，cmd.exe就会运行该可执行文件，而不是从您的PATH中选择。

For Go, PATH searches are handled by [`exec.LookPath`](https://pkg.go.dev/os/exec#LookPath), called automatically by [`exec.Command`](https://pkg.go.dev/os/exec#Command). And to fit well into the host system, Go’s `exec.LookPath` implements the Unix rules on Unix and the Windows rules on Windows. For example, this command

对于 Go，PATH 搜索由 exec.LookPath 处理，由 exec.Command 自动调用。为了很好地适应主机系统，Go 的 exec.LookPath 在 Unix 上执行 Unix 规则，在 Windows 上执行 Windows 规则。例如，这个命令

```
out, err := exec.Command("go", "version").CombinedOutput()
```

behaves the same as typing `go` `version` into the operating system shell. On Windows, it runs `.\go.exe` when that exists.

的行为与在操作系统shell中输入go版本的行为相同。在Windows上，当存在.\go.exe时，它会运行该程序。

(It is worth noting that Windows PowerShell changed this behavior, dropping the implicit search of dot, but `cmd.exe` and the Windows C library [`SearchPath function`](https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-searchpatha) continue to behave as they always have. Go continues to match `cmd.exe`.)

(值得注意的是，Windows PowerShell改变了这一行为，放弃了隐含的搜索点，但cmd.exe和Windows C库SearchPath函数的行为仍与以往相同。Go继续匹配cmd.exe）。

## The Bug

When `go` `get` downloads and builds a package that contains `import` `"C"`, it runs a program called `cgo` to prepare the Go equivalent of the relevant C code. The `go` command runs `cgo` in the directory containing the package sources. Once `cgo` has generated its Go output files, the `go` command itself invokes the Go compiler on the generated Go files and the host C compiler (`gcc` or `clang`) to build any C sources included with the package. All this works well. But where does the `go` command find the host C compiler? It looks in the PATH, of course. Luckily, while it runs the C compiler in the package source directory, it does the PATH lookup from the original directory where the `go` command was invoked:

当go get下载并构建一个包含import "C "的包时，它会运行一个名为cgo的程序来准备相关C代码的Go等价物。go命令在包含软件包源代码的目录中运行cgo。一旦cgo生成了它的Go输出文件，go命令本身就会在生成的Go文件上调用Go编译器和主机C编译器（gcc或clang）来构建软件包中包含的任何C源代码。所有这些都运作良好。但是go命令在哪里找到主机C编译器呢？当然是在PATH中寻找。幸运的是，当它在软件包源目录下运行C编译器时，它从调用go命令的原始目录中进行PATH查找：

```go linenums="1"
cmd := exec.Command("gcc", "file.c")
cmd.Dir = "badpkg"
cmd.Run()
```

So even if `badpkg\gcc.exe` exists on a Windows system, this code snippet will not find it. The lookup that happens in `exec.Command` does not know about the `badpkg` directory.

因此，即使badpkg\gcc.exe在Windows系统中存在，这个代码段也不会找到它。在exec.Command中发生的查找并不了解badpkg目录。

The `go` command uses similar code to invoke `cgo`, and in that case there’s not even a path lookup, because `cgo` always comes from GOROOT:

go命令使用类似的代码来调用cgo，在这种情况下，甚至没有路径查询，因为cgo总是来自GOROOT：

```go linenums="1"
cmd := exec.Command(GOROOT+"/pkg/tool/"+GOOS_GOARCH+"/cgo", "file.go")
cmd.Dir = "badpkg"
cmd.Run()
```

This is even safer than the previous snippet: there’s no chance of running any bad `cgo.exe` that may exist.

这比之前的片段更安全：没有机会运行可能存在的任何坏的cgo.exe。

But it turns out that cgo itself also invokes the host C compiler, on some temporary files it creates, meaning it executes this code itself:

但事实证明，cgo本身也调用了宿主的C语言编译器，在它创建的一些临时文件上，也就是说它自己执行了这些代码：

```go linenums="1"
// running in cgo in badpkg dir
cmd := exec.Command("gcc", "tmpfile.c")
cmd.Run()
```

Now, because cgo itself is running in `badpkg`, not in the directory where the `go` command was run, it will run `badpkg\gcc.exe` if that file exists, instead of finding the system `gcc`.

现在，由于cgo本身是在badpkg中运行的，而不是在运行go命令的目录中，如果该文件存在，它将运行badpkg\gcc.exe，而不是寻找系统中的gcc。



So an attacker can create a malicious package that uses cgo and includes a `gcc.exe`, and then any Windows user that runs `go` `get` to download and build the attacker’s package will run the attacker-supplied `gcc.exe` in preference to any `gcc` in the system path.

因此，攻击者可以创建一个使用cgo并包括gcc.exe的恶意软件包，然后任何运行go get来下载和构建攻击者软件包的Windows用户都会优先运行攻击者提供的gcc.exe，而不是系统路径中的任何gcc。

Unix systems avoid the problem first because dot is typically not in the PATH and second because module unpacking does not set execute bits on the files it writes. But Unix users who have dot ahead of system directories in their PATH and are using GOPATH mode would be as susceptible as Windows users. (If that describes you, today is a good day to remove dot from your path and to start using Go modules.)

Unix系统避免了这个问题，首先是因为dot通常不在PATH中，其次是因为模块解包不会在它写入的文件上设置执行位。但是，如果Unix用户在他们的PATH中把dot放在系统目录的前面，并且使用GOPATH模式，就会像Windows用户一样容易受到影响。(如果这描述了您，今天是一个从您的路径中删除dot并开始使用Go模块的好日子。)

(Thanks to [RyotaK](https://twitter.com/ryotkak) for [reporting this issue](https://go.dev/security) to us.)

(感谢RyotaK向我们报告这个问题）。

## The Fixes 修复方法

It’s obviously unacceptable for the `go` `get` command to download and run a malicious `gcc.exe`. But what’s the actual mistake that allows that? And then what’s the fix?

显然，go get命令下载并运行恶意的gcc.exe是不可接受的。但是，允许这种情况发生的实际错误是什么？然后又该如何修复呢？

One possible answer is that the mistake is that `cgo` does the search for the host C compiler in the untrusted source directory instead of in the directory where the `go` command was invoked. If that’s the mistake, then the fix is to change the `go` command to pass `cgo` the full path to the host C compiler, so that `cgo` need not do a PATH lookup in to the untrusted directory.

一个可能的答案是，错误在于cgo在不被信任的源代码目录下搜索主机C编译器，而不是在调用go命令的目录下搜索。如果这是一个错误，那么解决方法就是改变 go 命令，将主机 C 编译器的完整路径传递给 cgo，这样 cgo 就不需要在不受信任的目录中进行 PATH 查找了。

Another possible answer is that the mistake is to look in dot during PATH lookups, whether happens automatically on Windows or because of an explicit PATH entry on a Unix system. A user may want to look in dot to find a command they typed in a console or shell window, but it’s unlikely they also want to look there to find a subprocess of a subprocess of a typed command. If that’s the mistake, then the fix is to change the `cgo` command not to look in dot during a PATH lookup.

另一个可能的答案是，在PATH查找过程中，错误的是在dot中查找，不管是在Windows上自动发生，还是因为在Unix系统上有一个明确的PATH条目。用户可能想在dot中查找他们在控制台或shell窗口中输入的命令，但他们不可能也想在那里查找一个输入的命令的子进程。如果这是一个错误，那么解决的办法就是改变cgo命令，使其在PATH查找过程中不在dot中查找。

We decided both were mistakes, so we applied both fixes. The `go` command now passes the full host C compiler path to `cgo`. On top of that, `cgo`, `go`, and every other command in the Go distribution now use a variant of the `os/exec` package that reports an error if it would have previously used an executable from dot. The packages `go/build` and `go/import` use the same policy for their invocation of the `go` command and other tools. This should shut the door on any similar security problems that may be lurking.

我们认为这两个都是错误，所以我们应用了这两个修复方法。现在go命令将完整的主机C编译器路径传递给cgo。除此之外，cgo、go 以及 Go 发行版中的其他所有命令现在都使用了 os/exec 包的一个变体，如果它之前使用了 dot 中的可执行文件，则会报告错误。包go/build和go/import在调用go命令和其他工具时使用同样的策略。这应该可以关闭任何可能潜伏的类似安全问题的大门。

Out of an abundance of caution, we also made a similar fix in commands like `goimports` and `gopls`, as well as the libraries `golang.org/x/tools/go/analysis` and `golang.org/x/tools/go/packages`, which invoke the `go` command as a subprocess. If you run these programs in untrusted directories – for example, if you `git` `checkout` untrusted repositories and `cd` into them and then run programs like these, and you use Windows or use Unix with dot in your PATH – then you should update your copies of these commands too. If the only untrusted directories on your computer are the ones in the module cache managed by `go` `get`, then you only need the new Go release.

出于谨慎考虑，我们也对goimports和gopls等命令以及golang.org/x/tools/go/analysis和golang.org/x/tools/go/packages库进行了类似的修复，这些库将go命令作为一个子进程来调用。如果您在不被信任的目录中运行这些程序 —— 例如，如果您git签出不被信任的仓库并cd进入它们，然后运行像这样的程序，并且您使用Windows或使用Unix，在您的PATH中有dot —— 那么您也应该更新您的这些命令的副本。如果您的电脑上唯一不受信任的目录是由go get管理的模块缓存中的目录，那么您只需要新的Go版本。

After updating to the new Go release, you can update to the latest `gopls` by using:

更新到新的Go版本后，您可以通过以下方式更新到最新的gopls：

```
GO111MODULE=on \
go get golang.org/x/tools/gopls@v0.6.4
```

and you can update to the latest `goimports` or other tools by using:

并可以通过以下方式更新到最新的goimports或其他工具：

```
GO111MODULE=on \
go get golang.org/x/tools/cmd/goimports@v0.1.0
```

You can update programs that depend on `golang.org/x/tools/go/packages`, even before their authors do, by adding an explicit upgrade of the dependency during `go` `get`:

您可以更新依赖 golang.org/x/tools/go/packages 的程序，甚至比它们的作者更早，方法是在 go get 时加入明确的升级依赖关系：

```
GO111MODULE=on \
go get example.com/cmd/thecmd golang.org/x/tools@v0.1.0
```

For programs that use `go/build`, it is sufficient for you to recompile them using the updated Go release.

对于使用go/build的程序，您使用更新的Go版本重新编译它们就可以了。

Again, you only need to update these other programs if you are a Windows user or a Unix user with dot in the PATH *and* you run these programs in source directories you do not trust that may contain malicious programs.

同样，只有当您是Windows用户或Unix用户，PATH中有dot，并且您在您不信任的可能包含恶意程序的源文件目录中运行这些程序时，您才需要更新这些其他程序。

## Are your own programs affected? 您自己的程序是否受到影响？

If you use `exec.LookPath` or `exec.Command` in your own programs, you only need to be concerned if you (or your users) run your program in a directory with untrusted contents. If so, then a subprocess could be started using an executable from dot instead of from a system directory. (Again, using an executable from dot happens always on Windows and only with uncommon PATH settings on Unix.)

如果您在自己的程序中使用exec.LookPath或exec.Command，只有当您（或您的用户）在一个内容不受信任的目录中运行您的程序时，您才需要关注。如果是这样，那么子进程就可以使用来自dot的可执行文件而不是来自系统目录的可执行文件来启动。(同样，使用dot的可执行文件在Windows上总是发生，而在Unix上只有在不常见的PATH设置下才会发生）。

If you are concerned, then we’ve published the more restricted variant of `os/exec` as [`golang.org/x/sys/execabs`](https://pkg.go.dev/golang.org/x/sys/execabs). You can use it in your program by simply replacing

如果您担心，那么我们已经将os/exec的更多限制性变体发布为golang.org/x/sys/execabs。您可以在您的程序中使用它，只需替换掉

```go linenums="1"
import "os/exec"
```

with

替换为

```go linenums="1"
import exec "golang.org/x/sys/execabs"
```

and recompiling.

并重新编译。

## Securing os/exec by default 默认情况下保护os/exec的安全

We have been discussing on [golang.org/issue/38736](https://go.dev/issue/38736) whether the Windows behavior of always preferring the current directory in PATH lookups (during `exec.Command` and `exec.LookPath`) should be changed. The argument in favor of the change is that it closes the kinds of security problems discussed in this blog post. A supporting argument is that although the Windows `SearchPath` API and `cmd.exe` still always search the current directory, PowerShell, the successor to `cmd.exe`, does not, an apparent recognition that the original behavior was a mistake. The argument against the change is that it could break existing Windows programs that intend to find programs in the current directory. We don’t know how many such programs exist, but they would get unexplained failures if the PATH lookups started skipping the current directory entirely.

我们一直在 golang.org/issue/38736 上讨论是否应该改变 Windows 在 PATH 查找中（在 exec.Command 和 exec.LookPath 期间）总是优先选择当前目录的行为。赞成这一改变的观点是，它可以解决本博文中讨论的各种安全问题。一个支持的论点是，尽管Windows SearchPath API和cmd.exe仍然总是搜索当前目录，但PowerShell，即cmd.exe的继承者，并没有这样做，这显然是承认原来的行为是一个错误。反对这一变化的理由是，它可能会破坏现有的打算在当前目录中寻找程序的Windows程序。我们不知道有多少这样的程序存在，但如果PATH查找开始完全跳过当前目录，它们会得到无法解释的失败。

The approach we have taken in `golang.org/x/sys/execabs` may be a reasonable middle ground. It finds the result of the old PATH lookup and then returns a clear error rather than use a result from the current directory. The error returned from `exec.Command("prog")` when `prog.exe` exists looks like:

我们在golang.org/x/sys/execabs中采取的方法可能是一个合理的中间地带。它找到了旧的PATH查询的结果，然后返回一个明确的错误，而不是使用当前目录的结果。当prog.exe存在时，从exec.Command("prog")返回的错误看起来像：

```
prog resolves to executable in current directory (.\prog.exe)
```

For programs that do change behavior, this error should make very clear what has happened. Programs that intend to run a program from the current directory can use `exec.Command("./prog")` instead (that syntax works on all systems, even Windows).

对于那些确实改变了行为的程序，这个错误应该非常清楚地说明发生了什么。打算从当前目录运行程序的程序可以使用exec.Command("./prog")来代替（该语法在所有系统上都适用，甚至Windows）。

We have filed this idea as a new proposal, [golang.org/issue/43724](https://go.dev/issue/43724).

我们已经将这个想法作为一个新的提案提交，golang.org/issue/43724。
