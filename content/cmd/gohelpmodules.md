+++
title = "go help modules"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

A module is a collection of packages that are released, versioned, and distributed together. Modules may be downloaded directly from version control repositories or from module proxy servers.

For a series of tutorials on modules, see https://golang.org/doc/tutorial/create-module.

For a detailed reference on modules, see https://golang.org/ref/mod.

By default, the go command may download modules from https://proxy.golang.org. It may authenticate modules using the checksum database at https://sum.golang.org. Both services are operated by the Go team at Google.

The privacy policies for these services are available at https://proxy.golang.org/privacy and https://sum.golang.org/privacy, respectively.

The go command's download behavior may be configured using GOPROXY, GOSUMDB, GOPRIVATE, and other environment variables. See 'go help environment' and https://golang.org/ref/mod#private-module-privacy for more information.
