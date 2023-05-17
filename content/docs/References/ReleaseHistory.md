+++
title = "Release History"
date = 2023-05-17T09:59:21+08:00
weight = 3
description = ""
isCJKLanguage = true
draft = false
+++
# Release History

> 原文：[https://go.dev/doc/devel/release](https://go.dev/doc/devel/release)

本页总结了Go官方稳定版本之间的变化。[变更日志](https://go.dev/change)中有完整的细节。

要更新到一个特定的版本，请使用：

```
git fetch --tags
git checkout goX.Y.Z
```

## Release Policy 发布政策

​	每个主要的 Go 版本都被支持，直到有两个更新的主要版本。例如，Go 1.5 被支持到 Go 1.7 发布，而 Go 1.6 被支持到 Go 1.8 发布。我们会根据需要通过发布小修订版（例如，Go 1.6.1、Go 1.6.2，等等）来修复支持的版本中的关键问题，包括[关键的安全问题（critical security problems）](https://go.dev/security)。

## go1.20 (released 2023-02-01)

Go 1.20 is a major release of Go. Read the [Go 1.20 Release Notes](https://go.dev/doc/go1.20) for more information.

go 1.20是Go的一个重要版本。请阅读 [Go 1.20 发行说明](https://go.dev/doc/go1.20)了解更多信息。

### Minor revisions

go1.20.1 (released 2023-02-14) includes security fixes to the `crypto/tls`, `mime/multipart`, `net/http`, and `path/filepath` packages, as well as bug fixes to the compiler, the `go` command, the linker, the runtime, and the `time` package. See the [Go 1.20.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.20.1+label%3ACherryPickApproved) on our issue tracker for details.

go1.20.1（2023-02-14发布）包括对`crypto/tls`、`mime/multipart`、`net/http`和`path/filepath`包的安全修正，以及对编译器、go命令、链接器、`runtime`和`time`包的错误修正。详情请见我们问题跟踪器上的[Go 1.20.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.20.1+label%3ACherryPickApproved)。

## go1.19 (released 2022-08-02)

Go 1.19是Go的一个主要版本。请阅读 [Go 1.19 发行说明](https://go.dev/doc/go1.19)了解更多信息。

### Minor revisions

go1.19.1（2022-09-06发布）包括对`net/http`和`net/url`包的安全修复，以及对编译器、`go`命令、`pprof`命令、链接器、`runtime`和`crypto/tls`和`crypto/x509`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.19.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.19.1+label%3ACherryPickApproved)。

go1.19.2（2022-10-04发布）包括对`archive/tar`、`net/http/httputil`和`regexp`包的安全修复，以及对编译器、链接器、`runtime`和`go/types`包的错误修复。详情请见我们问题跟踪器上的[Go 1.19.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.19.2+label%3ACherryPickApproved)。

go1.19.3（2022-11-01发布）包括对`os/exec`和`syscall`包的安全修复，以及对编译器和`runtime`的错误修复。详情请见我们问题跟踪器上的 [Go 1.19.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.19.3+label%3ACherryPickApproved)。

go1.19.4（2022-12-06发布）包括对`net/http`和`os`包的安全修复，以及对编译器、`runtime`、`crypto/x509`、`os/exec`和`sync/atomic`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.19.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.19.4+label%3ACherryPickApproved)。

## go1.18 (released 2022-03-15)

Go 1.18是Go的一个重要版本。阅读 [Go 1.18 发行说明](https://go.dev/doc/go1.18)以了解更多信息。

### Minor revisions

go1.18.1（2022-04-12发布）包括对`crypto/elliptic`、`crypto/x509`和`encoding/pem`包的安全修复，以及对编译器、链接器、`runtime`、`go`命令、`vet`以及`bytes`、`crypto/x509`和`go/types`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.18.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.1+label%3ACherryPickApproved)。

go1.18.2（2022-05-10发布）包括对`syscall`包的安全修复，以及对编译器、`runtime`、`go`命令和`crypto/x509`、`go/types`、`net/http/httptest`、`reflect`和`sync/atomic`包的错误修复。详情请见我们问题跟踪器上的[Go 1.18.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.2+label%3ACherryPickApproved)。

go1.18.3（2022-06-01发布）包括对`crypto/rand`、`crypto/tls`、`os/exec`和`path/filepath`包的安全修复，以及对编译器、`crypto/tls`和`text/template/parse`包的错误修复。详情请见我们问题跟踪器上的[Go 1.18.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.3+label%3ACherryPickApproved)。

go1.18.4（2022-07-12发布）包括对`compress/gzip`、`encoding/gob`、`encoding/xml`、`go/parser`、`io/fs`、`net/http`和`path/filepath`包的安全修复，以及对编译器、`go`命令、链接器、`runtime`和`runtime`/计量包的错误修复。详情请见我们问题跟踪器上的[Go 1.18.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.4+label%3ACherryPickApproved)。

go1.18.5（2022-08-01发布）包括`encoding/gob`和`math/big`包的安全修复，以及编译器、`go`命令、`runtime`和`testing`包的错误修复。详情请见我们问题跟踪器上的[Go 1.18.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.5+label%3ACherryPickApproved)。

go1.18.6（2022-09-06发布）包括对`net/http`包的安全修复，以及对编译器、`go`命令、`pprof`命令、`runtime`以及`crypto/tls`、`encoding/xml`和`net`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.18.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.6+label%3ACherryPickApproved)。

go1.18.7（2022-10-04发布）包括对`archive/tar`、`net/http/httputil`和`regexp`包的安全修复，以及对编译器、链接器和`go/types`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.18.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.7+label%3ACherryPickApproved)。

go1.18.8（2022-11-01发布）包括对`os/exec`和`syscall`包的安全修复，以及对`runtime`的错误修复。详情请参见问题跟踪器上的 [Go 1.18.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.18.8+label%3ACherryPickApproved)。

## go1.17 (released 2021-08-16)

Go 1.17是Go的一个重要版本。阅读 [Go 1.17 发行说明](https://go.dev/doc/go1.17)以了解更多信息。

### Minor revisions

go1.17.1（2021-09-09发布）包括对`archive/zip`包的安全修复，以及对编译器、链接器、`go`命令和`crypto/rand`、`embed`、`go/types`、`html/template`和`net/http`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.17.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.1+label%3ACherryPickApproved)。

go1.17.2（2021-10-07发布）包括对linker和`misc/wasm`目录的安全修复，以及对编译器、`runtime`、`go`命令、`text/template`和`time`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.17.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.2+label%3ACherryPickApproved)。

go1.17.3（2021-11-04发布）包括对`archive/zip`和`debug/macho`包的安全修复，以及对编译器、链接器、`runtime`、`go`命令、`misc/wasm`目录以及`net/http`和`syscall`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.17.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.3+label%3ACherryPickApproved)。

go1.17.4（2021-12-02发布）包括对编译器、链接器、`runtime`以及`go/types`、`net/http`和`time`包的修复。详情请参见问题追踪器上的[Go 1.17.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.4+label%3ACherryPickApproved)。

go1.17.5（2021-12-09发布）包括`net/http`和`syscall`包的安全修复。详情请参见问题追踪器上的[Go 1.17.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.5+label%3ACherryPickApproved)。

go1.17.6 (2022-01-06发布) 包括对编译器、链接器、`runtime`以及 `crypto/x509`、`net/http` 和 `reflect` 包的修复。详情请参见问题跟踪器上的 [Go 1.17.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.6+label%3ACherryPickApproved)。

go1.17.7（2022-02-10发布）包括对`go`命令、`crypto/elliptic`和`math/big`包的安全修复，以及对编译器、链接器、`runtime`、`go`命令和`debug/macho`、`debug/pe`、`net/http/httptest`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.17.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.7+label%3ACherryPickApproved)。

go1.17.8（2022-03-03发布）包括对`regexp/syntax`包的安全修复，以及对编译器、`runtime`、`go`命令以及`crypto/x509`和`net`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.17.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.8+label%3ACherryPickApproved)。

go1.17.9（2022-04-12发布）包括对`crypto/elliptic`和`encoding/pem`包的安全修复，以及对链接器和`runtime`的错误修复。详情请参见问题跟踪器上的 [Go 1.17.9 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.9+label%3ACherryPickApproved)。

go1.17.10 (2022-05-10发布) 包括对`syscall`包的安全修复，以及对编译器、`runtime`、`crypto/x509`和`net/http/httptest`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.17.10 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.10+label%3ACherryPickApproved)。

go1.17.11（2022-06-01发布）包括对`crypto/rand`、`crypto/tls`、`os/exec`和`path/filepath`包的安全修复，以及对`crypto/tls`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.17.11 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.11+label%3ACherryPickApproved)。

go1.17.12（2022-07-12发布）包括对`compress/gzip`、`encoding/gob`、`encoding/xml`、`go/parser`、`io/fs`、`net/http`和`path/filepath`包的安全修复，以及对编译器、go命令、`runtime`和`runtime/metrics`包的错误修复。详情请见我们的问题跟踪器上的 [Go 1.17.12 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.12+label%3ACherryPickApproved)。

go1.17.13（2022-08-01发布）包括`encoding/gob`和`math/big`包的安全修复，以及编译器和`runtime`的错误修复。详情请参见问题跟踪器上的[Go 1.17.13 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.13+label%3ACherryPickApproved)。

## go1.16 (released 2021-02-16)

Go 1.16是Go的一个重要版本。请阅读 [Go 1.16 发行说明](https://go.dev/doc/go1.16)以了解更多信息。

### Minor revisions

go1.16.1（2021-03-10发布）包括对`archive/zip`和`encoding/xml`包的安全修复。详情请见我们问题追踪器上的 [Go 1.16.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.1+label%3ACherryPickApproved)。

go1.16.2（2021-03-11发布）包括对`cgo`、编译器、链接器、`go`命令、以及`syscall`和`time`包的修复。详情请参见问题追踪器上的 [Go 1.16.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.2+label%3ACherryPickApproved)。

go1.16.3（2021-04-01发布）包括对编译器、链接器、`runtime`、`go`命令以及`testing`和`time`包的修复。详情请参见问题跟踪器上的 [Go 1.16.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.3+label%3ACherryPickApproved)。

go1.16.4（2021-05-06发布）包括对`net/http`包的安全修复，以及对编译器、`runtime`、`archive/zip`、`syscall`和`time`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.16.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.4+label%3ACherryPickApproved)。

go1.16.5（2021-06-03发布）包括对`archive/zip`、`math/big`、`net`和`net/http/httputil`包的安全修复，以及对链接器、`go`命令和`net/http`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.16.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.5+label%3ACherryPickApproved)。

go1.16.6（2021-07-12发布）包括对`crypto/tls`包的安全修复，以及对编译器、`net`和`net/http`包的错误修复。详情请参见问题跟踪器上的[Go 1.16.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.6+label%3ACherryPickApproved)。

go1.16.7（2021-08-05发布）包括对`net/http/httputil`包的安全修复，以及对编译器、链接器、`runtime`、`go`命令和`net/http`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.16.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.7+label%3ACherryPickApproved)。

go1.16.8（2021-09-09发布）包括对`archive/zip`包的安全修复，以及对`archive/zip`、`go/internal/gccgoimporter`、`html/template`、`net/http`和`runtime/pprof`包的错误修复。详情请参见问题跟踪器上的 [Go 1.16.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.8+label%3ACherryPickApproved)。

go1.16.9（2021-10-07发布）包括对linker和`misc/wasm`目录的安全修复，以及对`runtime`和`text/template`包的错误修复。详情请见我们问题跟踪器上的[Go 1.16.9 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.9+label%3ACherryPickApproved)。

go1.16.10（2021-11-04发布）包括对`archive/zip`和`debug/macho`包的安全修复，以及对编译器、链接器、`runtime`、`misc/wasm`目录和`net/http`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.16.10 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.10+label%3ACherryPickApproved)。

go1.16.11（2021-12-02发布）包括对编译器、`runtime`以及`net/http`、`net/http/httptest`和`time`包的修复。详情请参见问题跟踪器上的[Go 1.16.11 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.11+label%3ACherryPickApproved)。

go1.16.12（2021-12-09发布）包括`net/http`和`syscall`包的安全修复。详情请参见问题追踪器上的[Go 1.16.12 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.12+label%3ACherryPickApproved)。

go1.16.13（2022-01-06发布）包括对编译器、链接器、`runtime`和`net/http`包的修复。详情请参见问题追踪器上的 [Go 1.16.13 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.13+label%3ACherryPickApproved)。

go1.16.14（2022-02-10发布）包括对`go`命令、`crypto/elliptic`和`math/big`包的安全修复，以及对编译器、链接器、`runtime`、`go`命令和`debug/macho`、`debug/pe`、`net/http/httptest`以及`testing`包的错误修复。详情请见我们问题跟踪器上的[Go 1.16.14 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.14+label%3ACherryPickApproved)。

go1.16.15（2022-03-03发布）包括对`regexp/syntax`包的安全修复，以及对编译器、`runtime`、`go`命令和`net`包的错误修复。详情请见我们问题跟踪器上的[Go 1.16.15 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.16.15+label%3ACherryPickApproved)。



## go1.15 (released 2020-08-11)

Go 1.15是Go的一个重要版本。请阅读 [Go 1.15 发行说明]()以了解更多信息。

### Minor revisions

go1.15.1（2020-09-01发布）包括`net/http/cgi`和`net/http/fcgi`包的安全修复。详情请见我们问题跟踪器上的 [Go 1.15.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.1+label%3ACherryPickApproved)。

go1.15.2（2020-09-09发布）包括对编译器、`runtime`、文档、`go`命令以及`net/mail`、`os`、`sync`和`testing`包的修复。详情请见我们问题跟踪器上的[Go 1.15.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.2+label%3ACherryPickApproved)。

go1.15.3（2020-10-14发布）包括对`cgo`、编译器、`runtime`、`go`命令，以及`bytes`、`plugin`和`testing`包的修复。详情请见我们问题跟踪器上的[Go 1.15.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.3+label%3ACherryPickApproved)。

go1.15.4（2020-11-05发布）包括对`cgo`、编译器、链接器、`runtime`，以及`compress/flate`、`net/http`、`reflect`和`time`包的修复。详情请见我们问题跟踪器上的[Go 1.15.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.4+label%3ACherryPickApproved)。

go1.15.5（2020-11-12发布）包括对`go`命令和`math/big`包的安全修复。详情请参见问题追踪器上的[Go 1.15.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.5+label%3ACherryPickApproved)。

go1.15.6（2020-12-03发布）包括对编译器、链接器、`runtime`、`go`命令和`io`包的修复。详情请参见问题追踪器上的[Go 1.15.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.6+label%3ACherryPickApproved)。

go1.15.7（2021-01-19发布）包括对`go`命令和`crypto/elliptic`包的安全修复。详情请参见问题追踪器上的[Go 1.15.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.7+label%3ACherryPickApproved)。

go1.15.8（2021-02-04发布）包括对编译器、链接器、`runtime`、`go`命令和`net/http`包的修复。详情请参见问题跟踪器上的 [Go 1.15.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.8+label%3ACherryPickApproved)。

go1.15.9（2021-03-10发布）包括对`encoding/xml`包的安全修复。详见问题追踪器上的[Go 1.15.9 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.9+label%3ACherryPickApproved)。

go1.15.10（2021-03-11发布）包括对编译器、`go`命令以及`net/http`、`os`、`syscall`和`time`包的修复。详情请参见问题跟踪器上的[Go 1.15.10 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.10+label%3ACherryPickApproved)。

go1.15.11（2021-04-01发布）包括对cgo、编译器、链接器、`runtime`、`go`命令以及`database/sql`和`net/http`包的修复。详情请见我们问题跟踪器上的[Go 1.15.11 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.11+label%3ACherryPickApproved)。

go1.15.12（2021-05-06发布）包括net/http包的安全修复，以及编译器、`runtime`、`archive/zip`、`syscall`和`time`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.15.12 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.12+label%3ACherryPickApproved)。

go1.15.13（2021-06-03发布）包括对`archive/zip`、`math/big`、`net`和`net/http/httputil`包的安全修复，以及对链接器、`go`命令、`math/big`和`net/http`包的错误修复。详情请见我们问题跟踪器上的 [Go 1.15.13 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.13+label%3ACherryPickApproved)。

go1.15.14（2021-07-12发布）包括对`crypto/tls`包的安全修复，以及对链接器和`net`包的错误修复。详情请参见问题跟踪器上的[Go 1.15.14 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.14+label%3ACherryPickApproved)。

go1.15.15（2021-08-05发布）包括对`net/http/httputil`包的安全修复，以及对编译器、`runtime`、`go`命令和`net/http`包的错误修复。详情请见我们问题跟踪器上的[Go 1.15.15 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.15.15+label%3ACherryPickApproved)。

## go1.14 (released 2020-02-25)

Go 1.14是Go的一个重要版本。请阅读 [Go 1.14 发行说明](https://go.dev/doc/go1.14)以了解更多信息。

### Minor revisions

go1.14.1（2020-03-19发布）包括对`go`命令、工具和`runtime`的修正。请参阅我们问题跟踪器上的 [Go 1.14.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.1+label%3ACherryPickApproved)以了解详情。

go1.14.2（2020-04-08发布）包括对`cgo`、`go`命令、`runtime`以及`os/exec`和`testing`包的修正。详情请参见问题追踪器上的 [Go 1.14.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.2+label%3ACherryPickApproved)。

go1.14.3（2020-05-14发布）包括对`cgo`、编译器、`runtime`以及`go/doc`和`math/big`包的修正。详情请参见问题追踪器上的 [Go 1.14.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.3+label%3ACherryPickApproved)。

go1.14.4（2020-06-01发布）包括对`go doc`命令、`runtime`、以及`encoding/json`和`os`包的修复。详情请参见问题追踪器上的 [Go 1.14.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.4+label%3ACherryPickApproved)。

go1.14.5（2020-07-14发布）包括对`crypto/x509`和`net/http`包的安全修复。详情请参见问题追踪器上的 [Go 1.14.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.5+label%3ACherryPickApproved)。

go1.14.6（2020-07-16发布）包括对`go`命令、编译器、链接器、`vet`以及`database/sql`、`encoding/json`、`net/http`、`reflect`和`testing`包的修复。详情请见我们的问题跟踪器上的 [Go 1.14.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.6+label%3ACherryPickApproved)。

go1.14.7（2020-08-06发布）包括对`encoding/binary`包的安全修复。详见问题追踪器上的[Go 1.14.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.7+label%3ACherryPickApproved)。

go1.14.8（2020-09-01发布）包括 `net/http/cgi` 和 `net/http/fcgi` 包的安全修复。详情请参见问题追踪器上的 [Go 1.14.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.8+label%3ACherryPickApproved)。

go1.14.9（2020-09-09发布）包括对编译器、链接器、`runtime`、文档，以及`net/http`和`testing`包的修复。详情请参见问题跟踪器上的 [Go 1.14.9 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.9+label%3ACherryPickApproved)。

go1.14.10（2020-10-14发布）包括对编译器、`runtime`以及`plugin`和`testing`包的修正。详情请参见问题追踪器上的[Go 1.14.10 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.10+label%3ACherryPickApproved)。

go1.14.11（2020-11-05发布）包括对`runtime`、`net/http`和`time`包的修复。详情请参见问题追踪器上的 [Go 1.14.11 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.11+label%3ACherryPickApproved)。

go1.14.12（2020-11-12发布）包括对`go`命令和`math/big`包的安全修复。详情请参见问题追踪器上的[Go 1.14.12 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.12+label%3ACherryPickApproved)。

go1.14.13（2020-12-03发布）包括对编译器、`runtime`和`go`命令的修复。详情请参见问题追踪器上的[Go 1.14.13 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.13+label%3ACherryPickApproved)。

go1.14.14（2021-01-19发布）包括对`go`命令和`crypto/elliptic`包的安全修复。详情请参见问题追踪器上的[Go 1.14.14 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.14+label%3ACherryPickApproved)。

go1.14.15（2021-02-04发布）包括对编译器、`runtime`、`go`命令和`net/http`包的修复。详情请见我们问题追踪器上的[Go 1.14.15 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.14.15+label%3ACherryPickApproved)。

## go1.13 (released 2019-09-03)

Go 1.13是Go的一个重要版本。阅读 [Go 1.13 发行说明](https://go.dev/doc/go1.13)以了解更多信息。

### Minor revisions

go1.13.1（2019-09-25发布）包括`net/http`和`net/textproto`包的安全修复。详情请见我们问题跟踪器上的[Go 1.13.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.1+label%3ACherryPickApproved)。

go1.13.2（2019-10-17发布）包括对编译器和`crypto/dsa`包的安全修复。详情请见我们问题追踪器上的[Go 1.13.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.2+label%3ACherryPickApproved)。

go1.13.3（2019-10-17发布）包括对`go`命令、工具链、`runtime`以及`crypto/ecdsa`、`net`、`net/http`和`syscall`包的修复。详情请见我们问题跟踪器上的[Go 1.13.3 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.3+label%3ACherryPickApproved)。

go1.13.4（2019-10-31发布）包括对`net/http`和`syscall`包的修复。它还修复了macOS 10.15 Catalina上的一个问题，即未经公证的安装程序和二进制文件[被Gatekeeper拒绝](https://go.dev/issue/34986)。详情请见我们问题追踪器上的 [Go 1.13.4 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.4+label%3ACherryPickApproved)。

go1.13.5（2019-12-04发布）包括对`go`命令、`runtime`、链接器和`net/http`包的修复。详情请见我们问题追踪器上的[Go 1.13.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.5+label%3ACherryPickApproved)。

go1.13.6（2020-01-09发布）包括对`runtime`和`net/http`包的修正。详情请参见问题追踪器上的 [Go 1.13.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.6+label%3ACherryPickApproved)。

go1.13.7（2020-01-28发布）包括对 `crypto/x509` 包的两个安全修复。详情请参见问题追踪器上的[Go 1.13.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.7+label%3ACherryPickApproved)。

go1.13.8（2020-02-12发布）包括对`runtime`、`crypto/x509`和`net/http`包的修复。详情请参见问题追踪器上的 [Go 1.13.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.8+label%3ACherryPickApproved)。

go1.13.9（2020-03-19发布）包括对`go`命令、工具、`runtime`、工具链和`crypto/cypher`包的修复。详情请见我们问题跟踪器上的[Go 1.13.9 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.9+label%3ACherryPickApproved)。

go1.13.10（2020-04-08发布）包括对`go`命令、`runtime`以及`os/exec`和`time`包的修复。详情请参见问题追踪器上的 [Go 1.13.10 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.10+label%3ACherryPickApproved)。

go1.13.11（2020-05-14发布）包括对编译器的修复。详见问题追踪器上的 [Go 1.13.11 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.11+label%3ACherryPickApproved)。

go1.13.12（2020-06-01发布）包括对`runtime`以及`go/types`和`math/big`包的修正。详情请参见问题追踪器上的 [Go 1.13.12 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.12+label%3ACherryPickApproved)。

go1.13.13（2020-07-14发布）包括对`crypto/x509`和`net/http`包的安全修复。详情请参见问题追踪器上的 [Go 1.13.13 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.13+label%3ACherryPickApproved)。

go1.13.14（2020-07-16发布）包括对编译器、`vet`以及`database/sql`、`net/http`和`reflect`包的修复。详情请参见问题跟踪器上的 [Go 1.13.14 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.14+label%3ACherryPickApproved)。

go1.13.15（2020-08-06发布）包括对`encoding/binary`包的安全修复。详情请见我们问题追踪器上的[Go 1.13.15 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.15+label%3ACherryPickApproved)。

## go1.12 (released 2019-02-25)

Go 1.12是Go的一个重要版本。阅读 [Go 1.12 发行说明](https://go.dev/doc/go1.12)以了解更多信息。

### Minor revisions

go1.12.1（2019-03-14发布）包括对`cgo`、编译器、`go`命令以及`fmt`、`net/smtp`、`os`、`path/filepath`、`sync`和`text/template`包的修正。详情请见我们问题跟踪器上的 [Go 1.12.1 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.1+label%3ACherryPickApproved)。

go1.12.2（2019-04-05发布）包括对`runtime`的安全修复，以及对编译器、`go`命令、`doc`、`net`、`net/http/httputil`和`os`包的错误修复。详情请参见问题跟踪器上的 [Go 1.12.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.2+label%3ACherryPickApproved)。

go1.12.3（2019-04-08发布）意外地在没有预定修复的情况下发布。它与go1.12.2相同，只是版本号不同。预定的修复在go1.12.4中。

go1.12.4（2019-04-11发布）修复了一个问题，即在旧版本的GNU/Linux上使用预编译的二进制版本会导致链接使用`cgo`的程序时[出现故障](https://go.dev/issues/31293)。只有遇到这个问题的Linux用户才需要更新。

go1.12.5（2019-05-06发布）包括对编译器、链接器、`go`命令、`runtime`和`os`包的修复。详情请见我们问题跟踪器上的[Go 1.12.5 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.5+label%3ACherryPickApproved)。

go1.12.6（2019-06-11发布）包括对编译器、链接器、`go`命令以及`crypto/x509`、`net/http`和`os`包的修复。详情请见我们问题跟踪器上的[Go 1.12.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.6+label%3ACherryPickApproved)。

go1.12.7（2019-07-08发布）包括对`cgo`、编译器和链接器的修复。详情请见我们的问题跟踪器上的[Go 1.12.7 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.7+label%3ACherryPickApproved)。

go1.12.8（2019-08-13发布）包括对`net/http`和`net/url`包的安全修复。详情请见我们问题追踪器上的[Go 1.12.8 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.12.8+label%3ACherryPickApproved)。

go1.12.9（2019-08-15发布）包括对链接器以及`math/big`和`os`包的修复。详情请见我们[问题追踪器上的Go 1.12.9里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.12.9+label%3ACherryPickApproved)。

go1.12.10（2019-09-25发布）包括对`net/http`和`net/textproto`包的安全修复。详情请见我们[问题追踪器上的Go 1.12.10里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.12.10+label%3ACherryPickApproved)。

go1.12.11（2019-10-17发布）包括对`crypto/dsa`包的安全修复。详情请参见[问题追踪器上的Go 1.12.11里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.12.11+label%3ACherryPickApproved)。

go1.12.12（2019-10-17发布）包括对`go`命令、`runtime`以及`net`和`syscall`包的修复。详情请见我们[问题追踪器上的Go 1.12.12里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.12.12+label%3ACherryPickApproved)。

go1.12.13（2019-10-31发布）修复了macOS 10.15 Catalina上的一个问题，即未经公证的安装程序和二进制文件[被Gatekeeper拒绝](https://go.dev/issue/34986)。只有遇到这个问题的macOS用户才需要更新。

go1.12.14（2019-12-04发布）包括一个对`runtime`的修复。详情请见我们[问题跟踪器上的Go 1.12.14里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.12.14+label%3ACherryPickApproved)。

go1.12.15（2020-01-09发布）包括对`runtime`和`net/http`包的修复。详情请参见[问题追踪器上的 Go 1.12.15 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.12.15+label%3ACherryPickApproved)。

go1.12.16（2020-01-28发布）包括对 `crypto/x509` 包的两个安全修复。详情请参见[问题追踪器上的 Go 1.12.16 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.12.16+label%3ACherryPickApproved)。

go1.12.17（2020-02-12发布）包括一个对`runtime`的修复。详见[问题追踪器上的Go 1.12.17里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.12.17+label%3ACherryPickApproved)。

## go1.11 (released 2018-08-24)

Go 1.11是Go的一个重要版本。阅读 [Go 1.11 发行说明](https://go.dev/doc/go1.11)以了解更多信息。

### Minor revisions

go1.11.1（2018-10-01发布）包括对编译器、文档、`go`命令、`runtime`以及`crypto/x509`、`encoding/json`、`go/types`、`net`、`net/http`和`reflect`包的修复。详情请见我们[问题跟踪器上的Go 1.11.1里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.1+label%3ACherryPickApproved)。

go1.11.2（2018-11-02发布）包括对编译器、链接器、文档、`go`命令，以及`database/sql`和`go/types`包的修复。详情请见我们[问题跟踪器上的Go 1.11.2里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.2+label%3ACherryPickApproved)。

go1.11.3（2018-12-12发布）包括对 "`go get`"和`crypto/x509`包的三个安全修复。详情请见我们[问题追踪器上的Go 1.11.3里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.3+label%3ACherryPickApproved)。

go1.11.4（2018-12-14发布）包括对`cgo`、编译器、链接器、`runtime`、文档、`go`命令以及`go/types`和`net/http`包的修复。它包括对Go 1.11.3中引入的一个bug的修复，该bug会破坏`go get`对于包含"`...`"的导入路径模式。详情请参见[问题跟踪器上的 Go 1.11.4 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.4+label%3ACherryPickApproved)。

go1.11.5（2019-01-23发布）包括一个对`crypto/elliptic`包的安全修复。详情请参见[问题追踪器上的Go 1.11.5里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.5+label%3ACherryPickApproved)。

go1.11.6（2019-03-14发布）包括对`cgo`、编译器、链接器、`runtime`、`go`命令以及`crypto/x509`、`encoding/json`、`net`和`net/url`包的修复。详情请见我们[问题跟踪器上的 Go 1.11.6 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.6+label%3ACherryPickApproved)。

go1.11.7（2019-04-05发布）包括对`runtime`和`net`包的修复。详见[问题追踪器上的Go 1.11.7里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.7+label%3ACherryPickApproved)。

go1.11.8（2019-04-08发布）是意外发布的，没有预定的修复。它与go1.11.7相同，只是版本号不同。预定的修复在go1.11.9中。

go1.11.9（2019-04-11发布）修复了一个问题，即在旧版本的GNU/Linux上使用预置二进制版本会导致链接使用`cgo`的程序时[出现故障](https://go.dev/issues/31293)。只有遇到这个问题的Linux用户才需要更新。

go1.11.10（2019-05-06发布）包括对`runtime`的安全修复，以及对链接器的错误修复。详情请见我们[问题跟踪器上的Go 1.11.10里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.10+label%3ACherryPickApproved)。

go1.11.11（2019-06-11发布）包括对`crypto/x509`包的修复。详情请见我们[问题追踪器上的Go 1.11.11里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.11+label%3ACherryPickApproved)。

go1.11.12（2019-07-08发布）包括对编译器和链接器的修复。详见[问题追踪器上的Go 1.11.12里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.12+label%3ACherryPickApproved)。

go1.11.13（2019-08-13发布）包括对`net/http`和`net/url`包的安全修复。详情请见我们的[问题跟踪器上的Go 1.11.13里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.11.13+label%3ACherryPickApproved)。

## go1.10 (released 2018-02-16)

Go 1.10是Go的一个重要版本。阅读 [Go 1.10 发行说明](https://go.dev/doc/go1.10)以了解更多信息。

### Minor revisions

go1.10.1（2018-03-28发布）包括对go命令的安全修复，以及对编译器、`runtime`和`archive/zip`、`crypto/tls`、`crypto/x509`、`encoding/json`、`net`、`net/http`和`net/http/pprof`包的错误修复。详情请见我们[问题跟踪器上的 Go 1.10.1 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.10.1+label%3ACherryPickApproved)。

go1.10.2（2018-05-01发布）包括对编译器、链接器和`go`命令的修复。详情请见我们[问题追踪器上的Go 1.10.2里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.10.2+label%3ACherryPickApproved)。

go1.10.3（2018-06-05发布）包括对`go`命令以及`crypto/tls`、`crypto/x509`和`strings`包的修复。特别是，它为[`go`命令增加了对`vgo`过渡的最小支持](https://go.googlesource.com/go/+/d4e21288e444d3ffd30d1a0737f15ea3fc3b8ad9)。详情请见我们[问题跟踪器上的 Go 1.10.3 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.10.3+label%3ACherryPickApproved)。

go1.10.4（2018-08-24发布）包括对go命令、链接器以及`bytes`、`mime/multipart`、`net/http`和`strings`包的修复。详情请见我们[问题跟踪器上的Go 1.10.4里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.10.4+label%3ACherryPickApproved)。

go1.10.5（2018-11-02发布）包括对go命令、链接器、`runtime`和`database/sql`包的修复。详情请见我们[问题追踪器上的Go 1.10.5里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.10.5+label%3ACherryPickApproved)。

go1.10.6（2018-12-12发布）包括对 "`go get`"和`crypto/x509`包的三个安全修正。它包含与Go 1.11.3相同的修复，并在同一时间发布。详情请见我们[问题跟踪器上的Go 1.10.6里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.10.6+label%3ACherryPickApproved)。

go1.10.7（2018-12-14发布）包含对Go 1.10.6中引入的一个错误的修复，该错误会破坏对包含"`...`"的导入路径模式的`go get`。详情请参见[问题追踪器上的 Go 1.10.7 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.10.7+label%3ACherryPickApproved)。

go1.10.8（2019-01-23发布）包括一个对`crypto/elliptic`包的安全修复。详情请见我们[问题追踪器上的Go 1.10.8里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.10.8+label%3ACherryPickApproved)。

## go1.9 (released 2017-08-24)

Go 1.9是Go的一个重要版本。阅读 [Go 1.9 发行说明](https://go.dev/doc/go1.9)以了解更多信息。

### Minor revisions

go1.9.1（2017-10-04发布）包括两个安全修复。请参阅我们的[问题跟踪器上的Go 1.9.1里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.9.1+label%3ACherryPickApproved)以了解详情。

go1.9.2（2017-10-25发布）包括对编译器、链接器、`runtime`、文档、`go`命令以及`crypto/x509`、`database/sql`、`log`和`net/smtp`包的修正。它包括对 Go 1.9.1 中引入的一个错误的修复，该错误在某些情况下会破坏非 Git 仓库的 `go get`。详情请参见[问题追踪器上的 Go 1.9.2 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.9.2+label%3ACherryPickApproved)。

go1.9.3（2018-01-22发布）包括`net/url`包的安全修复，以及编译器、`runtime`、`database/sql`、`math/big`和`net/http`包的错误修复。详情请见我们[问题跟踪器上的 Go 1.9.3 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.9.3+label%3ACherryPickApproved)。

go1.9.4（2018-02-07发布）包括一个对 "`go get`"的安全修复。详见我们[问题追踪器上的Go 1.9.4里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.9.4+label%3ACherryPickApproved)。

go1.9.5（2018-03-28发布）包括对`go`命令的安全修复，以及对编译器、`go`命令和`net/http/pprof`包的错误修复。详情请见我们[问题跟踪器上的Go 1.9.5里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.9.5+label%3ACherryPickApproved)。

go1.9.6（2018-05-01发布）包括对编译器和`go`命令的修复。详情请见我们[问题追踪器上的Go 1.9.6里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.9.6+label%3ACherryPickApproved)。

go1.9.7（2018-06-05发布）包括对`go`命令以及`crypto/x509`和`strings`包的修复。特别是，它为[`go`命令增加了对`vgo`过渡的最小支持](https://go.googlesource.com/go/+/d4e21288e444d3ffd30d1a0737f15ea3fc3b8ad9)。详情请见我们[问题跟踪器上的Go 1.9.7里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.9.7+label%3ACherryPickApproved)。

## go1.8 (released 2017-02-16)

Go 1.8是Go的一个重要版本。阅读 [Go 1.8 发行说明](https://go.dev/doc/go1.8)以了解更多信息。

### Minor revisions

go1.8.1（2017-04-07发布）包括对编译器、链接器、`runtime`、文档、`go`命令以及`crypto/tls`、`encoding/xml`、`image/png`、`net`、`net/http`、`reflect`、`text/template`和`time`包的修复。详情请见我们[问题跟踪器上的Go 1.8.1里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.8.1)。

go1.8.2（2017-05-23发布）包括一个对`crypto/elliptic`包的安全修复。详情请见我们[问题追踪器上的Go 1.8.2里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.8.2)。

go1.8.3（2017-05-24发布）包括对编译器、`runtime`、文档和`database/sql`包的修复。详情请见我们[问题追踪器上的Go 1.8.3里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.8.3)。

go1.8.4（2017-10-04发布）包括两个安全修复。它包含了与Go 1.9.1相同的修复，并在同一时间发布。详情请见我们[问题跟踪器上的Go 1.8.4里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.8.4)。

go1.8.5（2017-10-25发布）包括对编译器、链接器、`runtime`、文档、`go`命令以及`crypto/x509`和`net/smtp`包的修正。它包括对Go 1.8.4中引入的一个错误的修复，该错误在某些条件下破坏了非Git仓库的`go get`。详情请参见[问题跟踪器上的 Go 1.8.5 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.8.5)。

go1.8.6（2018-01-22发布）包括与Go 1.9.3相同的`math/big`中的修复，并在同一时间发布。详情请见我们[问题跟踪器上的Go 1.8.6里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.8.6)。

go1.8.7（2018-02-07发布）包含一个对 "`go get`"的安全修复。它包含与Go 1.9.4相同的修复，并在同一时间发布。详情请见我们[问题跟踪器上的Go 1.8.7里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.8.7)。

## go1.7 (released 2016-08-15)

Go 1.7是Go的一个重要版本。阅读 [Go 1.7 发行说明](https://go.dev/doc/go1.7)以了解更多信息。

### Minor revisions

go1.7.1（2016-09-07发布）包括对编译器、`runtime`、文档以及`compress/flate`、`hash/crc32`、`io`、`net`、`net/http`、`path/filepath`、`reflect`和`syscall`包的修正。详情请见我们[问题跟踪器上的 Go 1.7.1 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.7.1)。

go1.7.2不应被使用。它已被标记但未完全发布。由于最后一分钟的错误报告，该版本被推迟了。请使用go1.7.3来代替，并参考下面的变化摘要。

go1.7.3（2016-10-19发布）包括对编译器、`runtime`以及`crypto/cipher`、`crypto/tls`、`net/http`和`strings`包的修复。详情请见我们[问题跟踪器上的Go 1.7.3里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.7.3)。

go1.7.4（2016-12-01发布）包括两个安全修复。详见[问题追踪器上的 Go 1.7.4 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.7.4)。

go1.7.5（2017-01-26发布）包括对编译器、`runtime`以及`crypto/x509`和`time`包的修复。详情请见我们的[问题跟踪器上的Go 1.7.5里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.7.5)。

go1.7.6（2017-05-23发布）包括与Go 1.8.2相同的安全修复，并在同一时间发布。详情请见我们[问题追踪器上的Go 1.8.2里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.8.2)。

## go1.6 (released 2016-02-17)

Go 1.6是Go的一个重要版本。阅读 [Go 1.6 发行说明](https://go.dev/doc/go1.6)以了解更多信息。

### Minor revisions

go1.6.1（2016-04-12发布）包括两个安全修复。请参阅我们的[问题跟踪器上的 Go 1.6.1 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.6.1)了解详情。

go1.6.2（2016-04-20发布）包括对编译器、`runtime`、工具、文档，以及`mime/multipart`、`net/http`和`sort`包的修复。详情请见我们[问题跟踪器上的Go 1.6.2里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.6.2)。

go1.6.3（2016-07-17发布）包括在`CGI`环境下使用`net/http/cgi`包和`net/http`包时的安全修复。详情请见我们[问题跟踪器上的Go 1.6.3里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.6.3)。

go1.6.4（2016-12-01发布）包括两个安全修复。它包含与Go 1.7.4相同的修复，并在同一时间发布。详情请见我们[问题跟踪器上的Go 1.7.4里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.7.4)。

## go1.5 (released 2015-08-19)

Go 1.5是Go的一个重要版本。阅读 [Go 1.5 发行说明](https://go.dev/doc/go1.5)以了解更多信息。

### Minor revisions

go1.5.1（2015-09-08发布）包括对编译器、汇编器以及`fmt`、`net/textproto`、`net/http`和`runtime`包的错误修复。详情请见我们[问题跟踪器上的Go 1.5.1里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.5.1)。

go1.5.2（2015-12-02发布）包括对编译器、链接器以及`mime/multipart`、`net`和`runtime`包的错误修复。详情请见我们[问题跟踪器上的Go 1.5.2里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.5.2)。

go1.5.3（2016-01-13发布）包括对影响`crypto/tls`包的`math/big`包的安全修复。详见[发布公告](https://go.dev/s/go153announce)。

go1.5.4（2016-04-12发布）包括两个安全修复。它包含与Go 1.6.1相同的修复，并在同一时间发布。详情请见我们[问题跟踪器上的Go 1.6.1里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.6.1)。

## go1.4 (released 2014-12-10)

Go 1.4是Go的一个重要版本。阅读 [Go 1.4 发行说明](https://go.dev/doc/go1.4)以了解更多信息。

### Minor revisions

go1.4.1（2015-01-15发布）包括对链接器和`log`、`syscall`和`runtime`包的错误修复。详情请参见我们[问题跟踪器上的 Go 1.4.1 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.4.1)。

go1.4.2（2015-02-17发布）包括对编译器的安全修复，以及对`go`命令、编译器和链接器以及`runtime`、`syscall`、`reflect`和`math/big`包的错误修复。详情请见我们[问题跟踪器上的 Go 1.4.2 里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.4.2)。

go1.4.3（2015-09-22发布）包括`net/http`包的安全修复和`runtime`包的错误修复。详情请见我们[问题追踪器上的Go 1.4.3里程碑](https://github.com/golang/go/issues?q=milestone%3AGo1.4.3)。

## go1.3 (released 2014-06-18)

Go 1.3是Go的一个重要版本。阅读 [Go 1.3 发行说明](https://go.dev/doc/go1.3)以了解更多信息。

### Minor revisions

go1.3.1（2014-08-13发布）包括对编译器以及`runtime`、`net`和`crypto/rsa`包的错误修复。详细内容请参见[变更历史](https://github.com/golang/go/commits/go1.3.1)。

go1.3.2（2014-09-25发布）包括对`crypto/tls`包的安全修复和对`cgo`的错误修复。详见[变更历史](https://github.com/golang/go/commits/go1.3.2)。

go1.3.3（2014-09-30发布）包括对`cgo`、`runtime` 包和`nacl`端口的进一步错误修复。详见[变更历史](https://github.com/golang/go/commits/go1.3.3)。

## go1.2 (released 2013-12-01)

Go 1.2是Go的一个重要版本。请阅读 [Go 1.2 发行说明](https://go.dev/doc/go1.2) 以了解更多信息。

### Minor revisions

go1.2.1（2014-03-02发布）包括对`runtime`、`net`和`database/sql`包的错误修复。详情请参见[变更历史](https://github.com/golang/go/commits/go1.2.1)。

go1.2.2（2014-05-05发布）包括一个影响到二进制发行版中包含的tour binary的[安全修复](https://github.com/golang/go/commits/go1.2.2)（感谢Guillaume T）。

## go1.1 (released 2013-05-13)

Go 1.1是Go的一个重要版本。请阅读 [Go 1.1 发行说明](https://go.dev/doc/go1.1) 以了解更多信息。

### Minor revisions

go1.1.1（2013-06-13发布）包括对编译器的安全修复以及对编译器和`runtime`的若干错误修复。详情请参见[变更历史](https://github.com/golang/go/commits/go1.1.1)。

go1.1.2（2013-08-13发布）包括对`gc`编译器和`cgo`，以及`bufio`、`runtime`、`syscall`和`time`包的修复。详情请参见[变更历史](https://github.com/golang/go/commits/go1.1.2)。如果你在ARM或386架构的Linux下使用`syscall`包的`Getrlimit`和`Setrlimit`函数，请注意[11803043](https://go.dev/cl/11803043)的变动，它修复了[issue 5949](https://go.dev/issue/5949)。

## go1 (released 2012-03-28)

Go 1是Go的一个重要版本，将长期保持稳定。请阅读 [Go 1 发行说明](https://go.dev/doc/go1.html)了解更多信息。

我们希望为 Go 1 编写的程序能够在未来的 Go 1 版本下继续正确地编译和运行，不做任何改变。 阅读 [Go 1 兼容性文件](https://go.dev/doc/go1compat.html)，了解更多关于 Go 1 的未来。

go1版本对应的是`weekly.2012-03-27`。

### Minor revisions

go1.0.1（2012-04-25发布）是为了[修复](https://go.dev/cl/6061043)一个可能导致内存损坏的[转义分析错误](https://go.dev/issue/3545)。它还包括几个小的代码和文档修复。

go1.0.2（2012-06-13发布）修正了使用结构体或数组键的映射实现中的两个错误：[issue 3695](https://go.dev/issue/3695)和[issue 3573](https://go.dev/issue/3573)。它还包括许多小的代码和文档修复。

go1.0.3（2012-09-21发布）包括一些小的代码和文档修正。

完整的变化列表请参见[go1发布分支历史](https://github.com/golang/go/commits/release-branch.go1)。

## Older releases 较早的版本

请参阅[Go1发布前的历史](../../Other/Pre-Go1ReleaseHistory)页面，了解早期发布的注意事项。