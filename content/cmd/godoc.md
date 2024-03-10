+++
title = "godoc"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# godoc

> 原文：[https://pkg.go.dev/golang.org/x/tools/cmd/godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

### Overview 概述

Godoc extracts and generates documentation for Go programs.

​	godoc 提取并生成Go程序的文档。

It runs as a web server and presents the documentation as a web page.

​	它以Web服务器的形式运行，并以网页的形式展示文档。

```
godoc -http=:6060
```

Usage:

​	使用方法：

```
godoc [flag]
```

The flags are:

​	标志有：

```
-v
	冗长的模式
-timestamps=true
	显示目录列表中的时间戳
-index
	启用标识符和全文搜索索引
	(如果没有设置-index，则不显示搜索框)
-index_files=""
	指定索引文件的glob模式；如果不是空的话。
	索引是按照排序从这些文件中读取的。
-index_throttle=0.75
	索引节流值；值为0意味着没有时间分配给索引器（索引器将永远不会完成），值为1.0意味着索引创建以全速运行（当索引建立时，其他goroutine可能没有时间）。
-index_interval=0
	建立索引的时间间隔；值为0时，设置为5分钟，负值时，只在启动时建立一次索引。
-play=false
	启用playground
-links=true
	将标识符链接到它们的声明中
-write_index=false
	将索引写到一个文件中；文件名必须用-index_files
-maxresults=10000
	显示全文搜索结果的最大数量
	(如果maxresults <= 0，则不建立全文索引)
-notes="BUG"
	正则表达式匹配要显示的注释标记
	(例如，"BUG|TODO", ".*")
-goroot=$GOROOT
	Go的根目录
-http=addr
	HTTP服务地址（例如，'127.0.0.1:6060'或只是':6060')
-templates=""
	包含备用模板文件的目录；如果设置，该目录可以为$GOROOT/lib/godoc中的文件提供备用模板文件
-url=path
	将HTTP请求提供的数据打印到标准输出中，该数据将由path提供。
-zip=""
	提供文件系统服务的zip文件；如果为空则禁用 
```

By default, godoc looks at the packages it finds via $GOROOT and $GOPATH (if set). This behavior can be altered by providing an alternative $GOROOT with the -goroot flag.

​	默认情况下，godoc 通过`$GOROOT`和`$GOPATH`（如果设置）查看它发现的包。这个行为可以通过使用`-goroot`标志提供一个替代的`$GOROOT`来改变。

When the -index flag is set, a search index is maintained. The index is created at startup.

​	当`-index`标志被设置时，一个搜索索引被维护。该索引在启动时被创建。

The index contains both identifier and full text search information (searchable via regular expressions). The maximum number of full text search results shown can be set with the -maxresults flag; if set to 0, no full text results are shown, and only an identifier index but no full text search index is created.

​	该索引包含标识符和全文搜索信息（可通过正则表达式搜索）。显示全文搜索结果的最大数量可以用`-maxresults`标志设置；如果设置为`0`，不显示全文搜索结果，只创建标识符索引，不创建全文搜索索引。

By default, godoc uses the system's GOOS/GOARCH. You can provide the URL parameters "GOOS" and "GOARCH" to set the output on the web page for the target system.

​	默认情况下，godoc使用系统的`GOOS/GOARCH`。您可以提供URL参数 "GOOS "和 "GOARCH"，为目标系统设置网页上的输出。

The presentation mode of web pages served by godoc can be controlled with the "m" URL parameter; it accepts a comma-separated list of flag names as value:

​	由godoc提供的网页的表现模式可以用 "`m`" URL参数控制；它接受一个用逗号分隔的标志名称列表作为值：

```
all	    show documentation for all declarations, not just the exported ones
		显示所有声明的文档，而不仅仅是导出的声明
		
methods	 show all embedded methods, not just those of unexported anonymous fields
		 显示所有嵌入的方法，而不仅仅是那些未导出的匿名字段的方法
		
src 	 show the original source code rather than the extracted documentation
		 显示原始源代码，而不是提取的文档
		
flat	 present flat (not indented) directory listings using full paths
		 显示使用完整路径的扁平（非缩进）目录列表
```

For instance, https://golang.org/pkg/math/big/?m=all shows the documentation for all (not just the exported) declarations of package big.

​	例如，https://golang.org/pkg/math/big/?m=all 显示`big`包的所有（而不仅仅是导出的）声明的文档。

By default, godoc serves files from the file system of the underlying OS. Instead, a .zip file may be provided via the -zip flag, which contains the file system to serve. The file paths stored in the .zip file must use slash ('/') as path separator; and they must be unrooted. $GOROOT (or -goroot) must be set to the .zip file directory path containing the Go root directory. For instance, for a .zip file created by the command:

​	默认情况下，godoc从底层操作系统的文件系统中提供文件。相反，可以通过`-zip`标志提供一个.zip文件，其中包含要服务的文件系统。存储在.zip文件中的文件路径必须使用斜线（'/'）作为路径分隔符；并且它们必须是无根的。`$GOROOT`（或`-goroot`）必须设置为包含Go根目录的.zip文件目录路径。例如，对于一个由命令创建的.zip文件：

```
zip -r go.zip $HOME/go
```

one may run godoc as follows:

可以按以下方式运行godoc：

```
godoc -http=:6060 -zip=go.zip -goroot=$HOME/go
```

Godoc documentation is converted to HTML or to text using the go/doc package; see https://golang.org/pkg/go/doc/#ToHTML for the exact rules. Godoc also shows example code that is runnable by the testing package; see https://golang.org/pkg/testing/#hdr-Examples for the conventions. See "Godoc: documenting Go code" for how to write good comments for godoc: https://golang.org/doc/articles/godoc_documenting_go_code.html

​	godoc文档被转换为HTML或使用go/doc包转换为文本，具体规则见[https://golang.org/pkg/go/doc/#ToHTML](https://golang.org/pkg/go/doc/#ToHTML)。godoc还显示可由测试包运行的示例代码；具体规则见[https://golang.org/pkg/testing/#hdr-Examples](https://golang.org/pkg/testing/#hdr-Examples)。参见 "Godoc：记录Go代码"，了解如何为godoc写好注释：[https://golang.org/doc/articles/godoc_documenting_go_code.html](https://golang.org/doc/articles/godoc_documenting_go_code.html)

Deprecated: godoc cannot select what version of a package is displayed. Instead, use golang.org/x/pkgsite/cmd/pkgsite.

​	已弃用：godoc不能选择要显示包的版本。相反，请使用golang.org/x/pkgsite/cmd/pkgsite。