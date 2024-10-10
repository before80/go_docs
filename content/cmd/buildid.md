+++
title = "buildid"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# buildid



### Overview

Buildid displays or updates the build ID stored in a Go package or binary.

​	buildid 显示或更新存储在 Go 包或二进制文件中的构建 ID。

用法：

​	Usage:

```
go tool buildid [-w] file
```

By default, buildid prints the build ID found in the named file. If the -w option is given, buildid rewrites the build ID found in the file to accurately record a content hash of the file.

​	默认情况下，buildid 打印在指定文件中找到的构建 ID。如果给出了 -w 选项，buildid 会重写文件中找到的构建 ID，以准确记录文件的哈希值。

This tool is only intended for use by the go command or other build systems.

​	此工具仅供 go 命令或其他构建系统使用。

=== "buildid.go"

```
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"cmd/internal/buildid"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: go tool buildid [-w] file\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var wflag = flag.Bool("w", false, "write build ID")

func main() {
	log.SetPrefix("buildid: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}

	file := flag.Arg(0)
	id, err := buildid.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	if !*wflag {
		fmt.Printf("%s\n", id)
		return
	}

	// Keep in sync with src/cmd/go/internal/work/buildid.go:updateBuildID

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	matches, hash, err := buildid.FindAndHash(f, id, 0)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// <= go 1.7 doesn't embed the contentID or actionID, so no slash is present
	if !strings.Contains(id, "/") {
		log.Fatalf("%s: build ID is a legacy format...binary too old for this tool", file)
	}

	newID := id[:strings.LastIndex(id, "/")] + "/" + buildid.HashToString(hash)
	if len(newID) != len(id) {
		log.Fatalf("%s: build ID length mismatch %q vs %q", file, id, newID)
	}

	if len(matches) == 0 {
		return
	}

	f, err = os.OpenFile(file, os.O_RDWR, 0)
	if err != nil {
		log.Fatal(err)
	}
	if err := buildid.Rewrite(f, matches, newID); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
```

=== "doc.go"

```
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Buildid displays or updates the build ID stored in a Go package or binary.

Usage:

	go tool buildid [-w] file

By default, buildid prints the build ID found in the named file.
If the -w option is given, buildid rewrites the build ID found in
the file to accurately record a content hash of the file.

This tool is only intended for use by the go command or
other build systems.
*/
package main
```

