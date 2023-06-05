+++
title = "dist"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# dist

### Overview 

Dist helps bootstrap, build, and test the Go distribution.

Usage:

```
go tool dist [command]
```

The commands are:

```
banner         print installation banner
bootstrap      rebuild everything
clean          deletes all built files
env [-p]       print environment (-p: include $PATH)
install [dir]  install individual directory
list [-json]   list all supported platforms
test [-h]      run Go test(s)
version        print Go version
```