+++
title = "fix"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# fix

> 原文：[https://pkg.go.dev/cmd/fix@go1.19.3](https://pkg.go.dev/cmd/fix@go1.19.3)

### Overview 

Fix finds Go programs that use old APIs and rewrites them to use newer ones. After you update to a new Go release, fix helps make the necessary changes to your programs.

​	fix 查找使用旧 API 的 Go 程序并将其重写为使用较新的 API。在您更新到新的 Go 版本后，fix 有助于对您的程序进行必要的更改。

Usage:

​	用法：

```
go tool fix [-r name,...] [path ...]
```

Without an explicit path, fix reads standard input and writes the result to standard output.

​	在没有明确路径的情况下，fix 会读取标准输入并将结果写入标准输出。

If the named path is a file, fix rewrites the named files in place. If the named path is a directory, fix rewrites all .go files in that directory tree. When fix rewrites a file, it prints a line to standard error giving the name of the file and the rewrite applied.

​	如果命名的路径是一个文件，fix 会就地重写命名的文件。如果命名的路径是一个目录，fix 会重写该目录树中的所有 .go 文件。当 fix 重写一个文件时，它会向标准错误输出打印一行，给出文件名和应用的重写。

If the -diff flag is set, no files are rewritten. Instead fix prints the differences a rewrite would introduce.

​	如果设置了 -diff 标志，则不会重写任何文件。相反，fix 会打印重写将引入的差异。

The -r flag restricts the set of rewrites considered to those in the named list. By default fix considers all known rewrites. Fix's rewrites are idempotent, so that it is safe to apply fix to updated or partially updated code even without using the -r flag.

​	-r 标志将考虑的重写集限制为命名列表中的重写。默认情况下，fix 会考虑所有已知的重写。Fix 的重写是幂等的，因此即使不使用 -r 标志，也可以安全地将 fix 应用于已更新或部分更新的代码。

Fix prints the full list of fixes it can apply in its help output; to see them, run go tool fix -help.

​	Fix 在其帮助输出中打印它可以应用的完整修复列表；要查看它们，请运行 go tool fix -help。

Fix does not make backup copies of the files that it edits. Instead, use a version control system's “diff” functionality to inspect the changes that fix makes before committing them.

​	Fix 不会对其编辑的文件进行备份。相反，使用版本控制系统的“diff”功能来检查 fix 在提交之前所做的更改。