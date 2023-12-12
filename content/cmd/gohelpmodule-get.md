+++
title = "go help module-get"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

The 'go get' command changes behavior depending on whether the go command is running in module-aware mode or legacy GOPATH mode. This help text, accessible as 'go help module-get' even in legacy GOPATH mode, describes 'go get' as it operates in module-aware mode.

Usage:

```cmd
go get [-t] [-u] [-v] [build flags] [packages]
```

Get resolves its command-line arguments to packages at specific module versions, updates go.mod to require those versions, and downloads source code into the module cache.

To add a dependency for a package or upgrade it to its latest version:

        go get example.com/pkg

To upgrade or downgrade a package to a specific version:

        go get example.com/pkg@v1.2.3

To remove a dependency on a module and downgrade modules that require it:

        go get example.com/mod@none

To upgrade the minimum required Go version to the latest released Go version:

        go get go@latest

To upgrade the Go toolchain to the latest patch release of the current Go toolchain:

        go get toolchain@patch

See https://golang.org/ref/mod#go-get for details.

In earlier versions of Go, 'go get' was used to build and install packages. Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install' may be used to build and install commands instead. When a version is specified, 'go install' runs in module-aware mode and ignores the go.mod file in the current directory. For example:

        go install example.com/pkg@v1.2.3
        go install example.com/pkg@latest

See 'go help install' or https://golang.org/ref/mod#go-install for details.

'go get' accepts the following flags.

The -t flag instructs get to consider modules needed to build tests of packages specified on the command line.

The -u flag instructs get to update modules providing dependencies of packages named on the command line to use newer minor or patch releases when available.

The -u=patch flag (not -u patch) also instructs get to update dependencies, but changes the default to select patch releases.

When the -t and -u flags are used together, get will update test dependencies as well.

The -x flag prints commands as they are executed. This is useful for debugging version control commands when a module is downloaded directly from a repository.

For more about modules, see https://golang.org/ref/mod.

For more about using 'go get' to update the minimum Go version and suggested Go toolchain, see https://go.dev/doc/toolchain.

For more about specifying packages, see 'go help packages'.

This text describes the behavior of get using modules to manage source code and dependencies. If instead the go command is running in GOPATH mode, the details of get's flags and effects change, as does 'go help get'. See 'go help gopath-get'.

See also: go build, go install, go clean, go mod.
