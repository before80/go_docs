+++
title = "抗路径穿越的文件 API"
date = 2025-03-31T14:27:27+08:00
weight = 950
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/osroot](https://go.dev/blog/osroot)

# Traversal-resistant file APIs - 抗路径穿越的文件 API

Damien Neil  
12 March 2025  

A *path traversal vulnerability* arises when an attacker can trick a program into opening a file other than the one it intended. This post explains this class of vulnerability, some existing defenses against it, and describes how the new [`os.Root`](https://go.dev/pkg/os#Root) API added in Go 1.24 provides a simple and robust defense against unintentional path traversal.

​	*路径穿越漏洞* 发生在攻击者能够欺骗程序打开其未打算打开的文件时。本文将解释这类漏洞、一些现有的防御措施，并描述 Go 1.24 中新增的 [`os.Root`](https://go.dev/pkg/os#Root) API 如何为无意中的路径穿越提供简单而强大的防御。

## Path traversal attacks 路径穿越攻击

"Path traversal" covers a number of related attacks following a common pattern: A program attempts to open a file in some known location, but an attacker causes it to open a file in a different location.

​	"路径穿越"涵盖了一系列遵循共同模式的攻击：程序试图在某个已知位置打开文件，但攻击者使其打开另一个位置的文件。

If the attacker controls part of the filename, they may be able to use relative directory components ("..") to escape the intended location:

​	如果攻击者控制文件名的一部分，他们可能使用相对目录组件 ("..") 逃离预定位置：

```go
f, err := os.Open(filepath.Join(trustedLocation, "../../../../etc/passwd"))
```

On Windows systems, some names have special meaning:

​	在 Windows 系统上，一些名称具有特殊含义：

```go
// f will print to the console.
// f 将打印到控制台。
f, err := os.Create(filepath.Join(trustedLocation, "CONOUT$"))
```

If the attacker controls part of the local filesystem, they may be able to use symbolic links to cause a program to access the wrong file:

​	如果攻击者控制本地文件系统的一部分，他们可能通过符号链接使程序访问错误的文件：

```go
// Attacker links /home/user/.config to /home/otheruser/.config:
// 攻击者将 /home/user/.config 链接到 /home/otheruser/.config：
err := os.WriteFile("/home/user/.config/foo", config, 0o666)
```

If the program defends against symlink traversal by first verifying that the intended file does not contain any symlinks, it may still be vulnerable to [time-of-check/time-of-use (TOCTOU) races](https://en.wikipedia.org/wiki/Time-of-check_to_time-of-use), where the attacker creates a symlink after the program’s check:

​	如果程序通过首先验证目标文件不包含任何符号链接来防御符号链接穿越，它仍可能易受 [检查时/使用时 (TOCTOU) 竞争](https://en.wikipedia.org/wiki/Time-of-check_to_time-of-use) 的影响，此时攻击者在程序检查后创建符号链接：

```go
// Validate the path before use.
// 在使用前验证路径。
cleaned, err := filepath.EvalSymlinks(unsafePath)
if err != nil {
  return err
}
if !filepath.IsLocal(cleaned) {
  return errors.New("unsafe path")
}

// Attacker replaces part of the path with a symlink.
// The Open call follows the symlink:
// 攻击者将路径的一部分替换为符号链接。
// Open 调用会跟随符号链接：
f, err := os.Open(cleaned)
```

Another variety of TOCTOU race involves moving a directory that forms part of a path mid-traversal. For example, the attacker provides a path such as "a/b/c/../../etc/passwd", and renames "a/b/c" to "a/b" while the open operation is in progress.

​	另一种 TOCTOU 竞争涉及在遍历中移动构成路径一部分的目录。例如，攻击者提供路径如 "a/b/c/../../etc/passwd"，并在打开操作进行时将 "a/b/c" 重命名为 "a/b"。

## Path sanitization 路径清理

Before we tackle path traversal attacks in general, let’s start with path sanitization. When a program’s threat model does not include attackers with access to the local file system, it can be sufficient to validate untrusted input paths before use.

​	在全面处理路径穿越攻击之前，我们先从路径清理开始。当程序的威胁模型不包括能够访问本地文件系统的攻击者时，在使用前验证不受信任的输入路径可能就足够了。

Unfortunately, sanitizing paths can be surprisingly tricky, especially for portable programs that must handle both Unix and Windows paths. For example, on Windows `filepath.IsAbs(`\foo`)` reports `false`, because the path "\foo" is relative to the current drive.

​	不幸的是，清理路径可能出乎意料地棘手，特别是对于必须同时处理 Unix 和 Windows 路径的可移植程序。例如，在 Windows 上，`filepath.IsAbs(`\foo`)` 返回 `false`，因为路径 "\foo" 是相对于当前驱动器的。

In Go 1.20, we added the [`path/filepath.IsLocal`](https://go.dev/pkg/path/filepath#IsLocal) function, which reports whether a path is "local". A "local" path is one which:

​	在 Go 1.20 中，我们添加了 [`path/filepath.IsLocal`](https://go.dev/pkg/path/filepath#IsLocal) 函数，用于报告路径是否为"本地"。"本地"路径是指：

- does not escape the directory in which it is evaluated ("../etc/passwd" is not allowed); 不会逃离其被评估的目录 ("../etc/passwd" 不允许)；

- is not an absolute path ("/etc/passwd" is not allowed); 不是绝对路径 ("/etc/passwd" 不允许)；

- is not empty ("" is not allowed); 不是空的 ("" 不允许)；

- on Windows, is not a reserved name ("COM1" is not allowed). 在 Windows 上，不是保留名称 ("COM1" 不允许)。

  

In Go 1.23, we added the [`path/filepath.Localize`](https://go.dev/pkg/path/filepath#Localize) function, which converts a /-separated path into a local operating system path.

​	在 Go 1.23 中，我们添加了 [`path/filepath.Localize`](https://go.dev/pkg/path/filepath#Localize) 函数，将以 / 分隔的路径转换为本地操作系统路径。

Programs that accept and operate on potentially attacker-controlled paths should almost always use `filepath.IsLocal` or `filepath.Localize` to validate or sanitize those paths.

​	接受并操作可能由攻击者控制的路径的程序几乎总是应该使用 `filepath.IsLocal` 或 `filepath.Localize` 来验证或清理这些路径。

## Beyond sanitization 超越清理

Path sanitization is not sufficient when attackers may have access to part of the local filesystem.

​	当攻击者可能访问本地文件系统的一部分时，路径清理是不够的。

Multi-user systems are uncommon these days, but attacker access to the filesystem can still occur in a variety of ways. An unarchiving utility that extracts a tar or zip file may be induced to extract a symbolic link and then extract a file name that traverses that link. A container runtime may give untrusted code access to a portion of the local filesystem.

​	如今多用户系统并不常见，但攻击者访问文件系统仍可能以多种方式发生。一个解压 tar 或 zip 文件的解档工具可能被诱导提取符号链接，然后提取遍历该链接的文件名。容器运行时可能将不受信任的代码访问本地文件系统的一部分。

Programs may defend against unintended symlink traversal by using the [`path/filepath.EvalSymlinks`](https://go.dev/pkg/path/filepath#EvalSymlinks) function to resolve links in untrusted names before validation, but as described above this two-step process is vulnerable to TOCTOU races.

​	程序可以通过使用 [`path/filepath.EvalSymlinks`](https://go.dev/pkg/path/filepath#EvalSymlinks) 函数在验证前解析不受信任名称中的链接来防御意外的符号链接穿越，但如上所述，这种两步过程容易受到 TOCTOU 竞争的影响。

Before Go 1.24, the safer option was to use a package such as [github.com/google/safeopen](https://go.dev/pkg/github.com/google/safeopen), that provides path traversal-resistant functions for opening a potentially-untrusted filename within a specific directory.

​	在 Go 1.24 之前，更安全的选择是使用类似 [github.com/google/safeopen](https://go.dev/pkg/github.com/google/safeopen) 的包，它提供了抗路径穿越的功能，用于在特定目录内打开可能不受信任的文件名。

## Introducing `os.Root` - 介绍 `os.Root`

In Go 1.24, we are introducing new APIs in the `os` package to safely open a file in a location in a traversal-resistant fashion.

​	在 Go 1.24 中，我们在 `os` 包中引入了新的 API，以抗穿越的方式安全地在某个位置打开文件。

The new [`os.Root`](https://go.dev/pkg/os#Root) type represents a directory somewhere in the local filesystem. Open a root with the [`os.OpenRoot`](https://go.dev/pkg/os#OpenRoot) function:

​	新的 [`os.Root`](https://go.dev/pkg/os#Root) 类型表示本地文件系统中某处的目录。使用 [`os.OpenRoot`](https://go.dev/pkg/os#OpenRoot) 函数打开一个根目录：

```go
root, err := os.OpenRoot("/some/root/directory")
if err != nil {
  return err
}
defer root.Close()
```

`Root` provides methods to operate on files within the root. These methods all accept filenames relative to the root, and disallow any operations that would escape from the root either using relative path components ("..") or symlinks.

​	`Root` 提供了在根目录内操作文件的方法。这些方法都接受相对于根目录的文件名，并禁止任何使用相对路径组件 ("..") 或符号链接逃离根目录的操作。

```go
f, err := root.Open("path/to/file")
```

`Root` permits relative path components and symlinks that do not escape the root. For example, `root.Open("a/../b")` is permitted. Filenames are resolved using the semantics of the local platform: On Unix systems, this will follow any symlink in "a" (so long as that link does not escape the root); while on Windows systems this will open "b" (even if "a" does not exist).

​	`Root` 允许不逃离根目录的相对路径组件和符号链接。例如，`root.Open("a/../b")` 是允许的。文件名根据本地平台的语义解析：在 Unix 系统上，这将跟随 "a" 中的任何符号链接（只要该链接不逃离根目录）；而在 Windows 系统上，这将打开 "b"（即使 "a" 不存在）。

`Root` currently provides the following set of operations:

​	`Root` 当前提供以下操作集：

```go
func (*Root) Create(string) (*File, error)
func (*Root) Lstat(string) (fs.FileInfo, error)
func (*Root) Mkdir(string, fs.FileMode) error
func (*Root) Open(string) (*File, error)
func (*Root) OpenFile(string, int, fs.FileMode) (*File, error)
func (*Root) OpenRoot(string) (*Root, error)
func (*Root) Remove(string) error
func (*Root) Stat(string) (fs.FileInfo, error)
```

In addition to the `Root` type, the new [`os.OpenInRoot`](https://go.dev/pkg/os#OpenInRoot) function provides a simple way to open a potentially-untrusted filename within a specific directory:

​	除了 `Root` 类型外，新的 [`os.OpenInRoot`](https://go.dev/pkg/os#OpenInRoot) 函数提供了一种简单的方法，在特定目录内打开可能不受信任的文件名：

```go
f, err := os.OpenInRoot("/some/root/directory", untrustedFilename)
```

​	The `Root` type provides a simple, safe, portable API for operating with untrusted filenames.

`Root` 类型为操作不受信任的文件名提供了一个简单、安全、可移植的 API。

## Caveats and considerations 注意事项和考虑

### Unix

On Unix systems, `Root` is implemented using the `openat` family of system calls. A `Root` contains a file descriptor referencing its root directory and will track that directory across renames or deletion.

​	在 Unix 系统上，`Root` 使用 `openat` 系统调用家族实现。一个 `Root` 包含一个引用其根目录的文件描述符，并会在重命名或删除时跟踪该目录。

`Root` defends against symlink traversal but does not limit traversal of mount points. For example, `Root` does not prevent traversal of Linux bind mounts. Our threat model is that `Root` defends against filesystem constructs that may be created by ordinary users (such as symlinks), but does not handle ones that require root privileges to create (such as bind mounts).

​	`Root` 防御符号链接穿越，但不限制挂载点的穿越。例如，`Root` 不阻止 Linux 绑定挂载的穿越。我们的威胁模型是，`Root` 防御普通用户可能创建的文件系统结构（如符号链接），但不处理需要 root 权限创建的结构（如绑定挂载）。

### Windows 

On Windows, `Root` opens a handle referencing its root directory. The open handle prevents that directory from being renamed or deleted until the `Root` is closed.

在 Windows 上，`Root` 打开一个引用其根目录的句柄。打开的句柄阻止该目录在 `Root` 关闭之前被重命名或删除。

`Root` prevents access to reserved Windows device names such as `NUL` and `COM1`.

`Root` 阻止访问 Windows 保留的设备名称，如 `NUL` 和 `COM1`。

### WASI 

On WASI, the `os` package uses the WASI preview 1 filesystem API, which are intended to provide traversal-resistant filesystem access. Not all WASI implementations fully support filesystem sandboxing, however, and `Root`’s defense against traversal is limited to that provided by the WASI implementation.

​	在 WASI 上，`os` 包使用 WASI preview 1 文件系统 API，旨在提供抗穿越的文件系统访问。然而，并非所有 WASI 实现都完全支持文件系统沙箱，`Root` 对穿越的防御仅限于 WASI 实现提供的程度。

### GOOS=js

When GOOS=js, the `os` package uses the Node.js file system API. This API does not include the openat family of functions, and so `os.Root` is vulnerable to TOCTOU (time-of-check-time-of-use) races in symlink validation on this platform.

​	当 GOOS=js 时，`os` 包使用 Node.js 文件系统 API。此 API 不包括 openat 系列函数，因此在此平台上 `os.Root` 在符号链接验证中容易受到 TOCTOU（检查时-使用时）竞争的影响。

When GOOS=js, a `Root` references a directory name rather than a file descriptor, and does not track directories across renames.

​		当 GOOS=js 时，一个 `Root` 引用的是目录名称而非文件描述符，且不会在重命名时跟踪目录。

### Plan 9 

Plan 9 does not have symlinks. On Plan 9, a `Root` references a directory name and performs lexical sanitization of filenames.

Plan 9 没有符号链接。在 Plan 9 上，一个 `Root` 引用目录名称并对文件名进行词法清理。

### Performance 性能

`Root` operations on filenames containing many directory components can be much more expensive than the equivalent non-`Root` operation. Resolving ".." components can also be expensive. Programs that want to limit the cost of filesystem operations can use `filepath.Clean` to remove ".." components from input filenames, and may want to limit the number of directory components.

​	对包含许多目录组件的文件名执行 `Root` 操作可能比等效的非 `Root` 操作昂贵得多。解析 ".." 组件也可能很昂贵。希望限制文件系统操作成本的程序可以使用 `filepath.Clean` 从输入文件名中移除 ".." 组件，并可能希望限制目录组件的数量。

## Who should use os.Root? 谁应该使用 os.Root？

You should use `os.Root` or `os.OpenInRoot` if:

​	您应该在以下情况下使用 `os.Root` 或 `os.OpenInRoot`：

- you are opening a file in a directory; AND 您在某个目录中打开文件；并且
- the operation should not access a file outside that directory. 该操作不应访问该目录外的文件。

For example, an archive extractor writing files to an output directory should use `os.Root`, because the filenames are potentially untrusted and it would be incorrect to write a file outside the output directory.

​	例如，一个将文件写入输出目录的归档提取器应使用 `os.Root`，因为文件名可能是不可信的，将文件写入输出目录外是错误的。

However, a command-line program that writes output to a user-specified location should not use `os.Root`, because the filename is not untrusted and may refer to anywhere on the filesystem.

​	然而，一个将输出写入用户指定位置的命令行程序不应使用 `os.Root`，因为文件名并非不受信任，可能指向文件系统中的任何位置。

As a good rule of thumb, code which calls `filepath.Join` to combine a fixed directory and an externally-provided filename should probably use `os.Root` instead.

​	作为一个经验法则，调用 `filepath.Join` 将固定目录与外部提供的文件名组合的代码可能应该改用 `os.Root`。

```go
// This might open a file not located in baseDirectory.
// 这可能会打开不在 baseDirectory 中的文件。
f, err := os.Open(filepath.Join(baseDirectory, filename))

// This will only open files under baseDirectory.
// 这只会打开 baseDirectory 下的文件。
f, err := os.OpenInRoot(baseDirectory, filename)
```

## Future work 未来工作

The `os.Root` API is new in Go 1.24. We expect to make additions and refinements to it in future releases.

​	`os.Root` API 在 Go 1.24 中是新的。我们预计在未来的版本中会对其进行补充和改进。

The current implementation prioritizes correctness and safety over performance. Future versions will take advantage of platform-specific APIs, such as Linux’s `openat2`, to improve performance where possible.

​	当前实现优先考虑正确性和安全性而非性能。未来版本将利用特定于平台的 API，例如 Linux 的 `openat2`，在可能的情况下提高性能。

There are a number of filesystem operations which `Root` does not support yet, such as creating symbolic links and renaming files. Where possible, we will add support for these operations. A list of additional functions in progress is in [go.dev/issue/67002](https://go.dev/issue/67002).

​	`Root` 尚不支持许多文件系统操作，例如创建符号链接和重命名文件。在可能的情况下，我们将为这些操作添加支持。正在进行中的附加函数列表在 [go.dev/issue/67002](https://go.dev/issue/67002) 中。