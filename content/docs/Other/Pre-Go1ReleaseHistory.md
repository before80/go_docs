+++
title = "Pre-Go 1 Release History"
weight = 2
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Pre-Go 1 Release History

> 原文：[https://go.dev/doc/devel/pre_go1](https://go.dev/doc/devel/pre_go1)

This page summarizes the changes between stable releases of Go prior to Go 1. See the [Release History](https://go.dev/doc/devel/release.html) page for notes on recent releases.

## r60 (released 2011/09/07)

The r60 release corresponds to `weekly.2011-08-17`. This section highlights the most significant changes in this release. For a more detailed summary, see the [weekly release notes](https://go.dev/doc/devel/weekly.html#2011-08-17). For complete information, see the [Mercurial change list](https://code.google.com/p/go/source/list?r=release-branch.r60).

### Language

An "else" block is now required to have braces except if the body of the "else" is another "if". Since gofmt always puts those braces in anyway, gofmt-formatted programs will not be affected. To fix other programs, run gofmt.

### Packages

[Package http](https://go.dev/pkg/http/)'s URL parsing and query escaping code (such as `ParseURL` and `URLEscape`) has been moved to the new [url package](https://go.dev/pkg/url/), with several simplifications to the names. Client code can be updated automatically with gofix.

[Package image](https://go.dev/pkg/image/) has had significant changes made to the `Pix` field of struct types such as [image.RGBA](https://go.dev/pkg/image/#RGBA) and [image.NRGBA](https://go.dev/pkg/image/#NRGBA). The [image.Image](https://go.dev/pkg/image/#Image) interface type has not changed, though, and you should not need to change your code if you don't explicitly refer to `Pix` fields. For example, if you decode a number of images using the [image/jpeg](https://go.dev/pkg/image/jpeg/) package, compose them using [image/draw](https://go.dev/pkg/image/draw/), and then encode the result using [image/png](https://go.dev/pkg/img/png), then your code should still work as before. If your code *does* refer to `Pix` fields see the [weekly.2011-07-19](https://go.dev/doc/devel/weekly.html#2011-07-19) snapshot notes for how to update your code.

[Package template](https://go.dev/pkg/template/) has been replaced with a new templating package (formerly `exp/template`). The original template package is still available as [old/template](https://go.dev/pkg/old/template/). The `old/template` package is deprecated and will be removed. The Go tree has been updated to use the new template package. We encourage users of the old template package to switch to the new one. Code that uses `template` or `exp/template` will need to change its import lines to `"old/template"` or `"template"`, respectively.

### Tools

[Goinstall](https://go.dev/cmd/goinstall/) now uses a new tag selection scheme. When downloading or updating, goinstall looks for a tag or branch with the `"go."` prefix that corresponds to the local Go version. For Go `release.r58` it looks for `go.r58`. For `weekly.2011-06-03` it looks for `go.weekly.2011-06-03`. If the specific `go.X` tag or branch is not found, it chooses the closest earlier version. If an appropriate tag or branch is found, goinstall uses that version of the code. Otherwise it uses the default version selected by the version control system. Library authors are encouraged to use the appropriate tag or branch names in their repositories to make their libraries more accessible.

### Minor revisions

r60.1 includes a [linker fix](https://go.dev/change/1824581bf62d), a pair of [goplay](https://go.dev/change/9ef4429c2c64) [fixes](https://go.dev/change/d42ed8c3098e), and a `json` package [fix](https://go.dev/change/d5e97874fe84) and a new [struct tag option](https://go.dev/change/4f0e6269213f).

r60.2 [fixes](https://go.dev/change/ff19536042ac) a memory leak involving maps.

r60.3 fixes a [reflect bug](https://go.dev/change/01fa62f5e4e5).

## r59 (released 2011/08/01)

The r59 release corresponds to `weekly.2011-07-07`. This section highlights the most significant changes in this release. For a more detailed summary, see the [weekly release notes](https://go.dev/doc/devel/weekly.html#2011-07-07). For complete information, see the [Mercurial change list](https://code.google.com/p/go/source/list?r=release-branch.r59).

### Language

This release includes a language change that restricts the use of `goto`. In essence, a `goto` statement outside a block cannot jump to a label inside that block. Your code may require changes if it uses `goto`. See [this changeset](https://go.dev/change/dc6d3cf9279d) for how the new rule affected the Go tree.

### Packages

As usual, [gofix](https://go.dev/cmd/gofix/) will handle the bulk of the rewrites necessary for these changes to package APIs.

[Package http](https://go.dev/pkg/http) has a new [FileSystem](https://go.dev/pkg/http/#FileSystem) interface that provides access to files. The [FileServer](https://go.dev/pkg/http/#FileServer) helper now takes a `FileSystem` argument instead of an explicit file system root. By implementing your own `FileSystem` you can use the `FileServer` to serve arbitrary data.

[Package os](https://go.dev/pkg/os/)'s `ErrorString` type has been hidden. Most uses of `os.ErrorString` can be replaced with [os.NewError](https://go.dev/pkg/os/#NewError).

[Package reflect](https://go.dev/pkg/reflect/) supports a new struct tag scheme that enables sharing of struct tags between multiple packages. In this scheme, the tags must be of the form:

```
	`key:"value" key2:"value2"`
```

The [StructField](https://go.dev/pkg/reflect/#StructField) type's Tag field now has type [StructTag](https://go.dev/pkg/reflect/#StructTag), which has a `Get` method. Clients of [json](https://go.dev/pkg/json) and [xml](https://go.dev/pkg/xml) will need to be updated. Code that says

```
	type T struct {
		X int "name"
	}
```

should become

```
	type T struct {
		X int `json:"name"`  // or `xml:"name"`
	}
```

Use [govet](https://go.dev/cmd/govet/) to identify struct tags that need to be changed to use the new syntax.

[Package sort](https://go.dev/pkg/sort/)'s `IntArray` type has been renamed to [IntSlice](https://go.dev/pkg/sort/#IntSlice), and similarly for [Float64Slice](https://go.dev/pkg/sort/#Float64Slice) and [StringSlice](https://go.dev/pkg/sort/#StringSlice).

[Package strings](https://go.dev/pkg/strings/)'s `Split` function has itself been split into [Split](https://go.dev/pkg/strings/#Split) and [SplitN](https://go.dev/pkg/strings/#SplitN). `SplitN` is the same as the old `Split`. The new `Split` is equivalent to `SplitN` with a final argument of -1.

[Package image/draw](https://go.dev/pkg/image/draw/)'s [Draw](https://go.dev/pkg/image/draw/#Draw) function now takes an additional argument, a compositing operator. If in doubt, use [draw.Over](https://go.dev/pkg/image/draw/#Op).



### Tools

[Goinstall](https://go.dev/cmd/goinstall/) now installs packages and commands from arbitrary remote repositories (not just Google Code, Github, and so on). See the [goinstall documentation](https://go.dev/cmd/goinstall/) for details.

## r58 (released 2011/06/29)

The r58 release corresponds to `weekly.2011-06-09` with additional bug fixes. This section highlights the most significant changes in this release. For a more detailed summary, see the [weekly release notes](https://go.dev/doc/devel/weekly.html#2011-06-09). For complete information, see the [Mercurial change list](https://code.google.com/p/go/source/list?r=release-branch.r58).

### Language

This release fixes a [use of uninitialized memory in programs that misuse `goto`](https://go.dev/change/b720749486e1).

### Packages

As usual, [gofix](https://go.dev/cmd/gofix/) will handle the bulk of the rewrites necessary for these changes to package APIs.

[Package http](https://go.dev/pkg/http/) drops the `finalURL` return value from the [Client.Get](https://go.dev/pkg/http/#Client.Get) method. The value is now available via the new `Request` field on [http.Response](https://go.dev/pkg/http/#Response). Most instances of the type map[string][]string in have been replaced with the new [Values](https://go.dev/pkg/http/#Values) type.

[Package exec](https://go.dev/pkg/exec/) has been redesigned with a more convenient and succinct API.

[Package strconv](https://go.dev/pkg/strconv/)'s [Quote](https://go.dev/pkg/strconv/#Quote) function now escapes only those Unicode code points not classified as printable by [unicode.IsPrint](https://go.dev/pkg/unicode/#IsPrint). Previously Quote would escape all non-ASCII characters. This also affects the [fmt](https://go.dev/pkg/fmt/) package's `"%q"` formatting directive. The previous quoting behavior is still available via strconv's new [QuoteToASCII](https://go.dev/pkg/strconv/#QuoteToASCII) function.

[Package os/signal](https://go.dev/pkg/os/signal/)'s [Signal](https://go.dev/pkg/os/#Signal) and [UnixSignal](https://go.dev/pkg/os/#UnixSignal) types have been moved to the [os](https://go.dev/pkg/os/) package.

[Package image/draw](https://go.dev/pkg/image/draw/) is the new name for `exp/draw`. The GUI-related code from `exp/draw` is now located in the [exp/gui](https://go.dev/pkg/exp/gui/) package.

### Tools

[Goinstall](https://go.dev/cmd/goinstall/) now observes the GOPATH environment variable to build and install your own code and external libraries outside of the Go tree (and avoid writing Makefiles).

### Minor revisions

r58.1 adds [build](https://go.dev/change/293c25943586) and [runtime](https://go.dev/change/bf17e96b6582) changes to make Go run on OS X 10.7 Lion.

## r57 (released 2011/05/03)

The r57 release corresponds to `weekly.2011-04-27` with additional bug fixes. This section highlights the most significant changes in this release. For a more detailed summary, see the [weekly release notes](https://go.dev/doc/devel/weekly.html#2011-04-27). For complete information, see the [Mercurial change list](https://code.google.com/p/go/source/list?r=release-branch.r57).

The new [gofix](https://go.dev/cmd/gofix) tool finds Go programs that use old APIs and rewrites them to use newer ones. After you update to a new Go release, gofix helps make the necessary changes to your programs. Gofix will handle the http, os, and syscall package changes described below, and we will update the program to keep up with future changes to the libraries. Gofix can’t handle all situations perfectly, so read and test the changes it makes before committing them. See [the gofix blog post](https://blog.golang.org/2011/04/introducing-gofix.html) for more information.

### Language

[Multiple assignment syntax](https://go.dev/doc/go_spec.html#Receive_operator) replaces the `closed` function. The syntax for channel receives allows an optional second assigned value, a boolean value indicating whether the channel is closed. This code:

```
	v := <-ch
	if closed(ch) {
		// channel is closed
	}
```

should now be written as:

```
	v, ok := <-ch
	if !ok {
		// channel is closed
	}
```

[Unused labels are now illegal](https://go.dev/doc/go_spec.html#Label_scopes), just as unused local variables are.

### Packages

[Package gob](https://go.dev/pkg/gob/) will now encode and decode values of types that implement the [GobEncoder](https://go.dev/pkg/gob/#GobEncoder) and [GobDecoder](https://go.dev/pkg/gob/#GobDecoder) interfaces. This allows types with unexported fields to transmit self-consistent descriptions; examples include [big.Int](https://go.dev/pkg/big/#Int.GobDecode) and [big.Rat](https://go.dev/pkg/big/#Rat.GobDecode).

[Package http](https://go.dev/pkg/http/) has been redesigned. For clients, there are new [Client](https://go.dev/pkg/http/#Client) and [Transport](https://go.dev/pkg/http/#Transport) abstractions that give more control over HTTP details such as headers sent and redirections followed. These abstractions make it easy to implement custom clients that add functionality such as [OAuth2](https://code.google.com/p/goauth2/source/browse/oauth/oauth.go). For servers, [ResponseWriter](https://go.dev/pkg/http/#ResponseWriter) has dropped its non-essential methods. The Hijack and Flush methods are no longer required; code can test for them by checking whether a specific value implements [Hijacker](https://go.dev/pkg/http/#Hijacker) or [Flusher](https://go.dev/pkg/http/#Flusher). The RemoteAddr and UsingTLS methods are replaced by [Request](https://go.dev/pkg/http/#Request)'s RemoteAddr and TLS fields. The SetHeader method is replaced by a Header method; its result, of type [Header](https://go.dev/pkg/http/#Header), implements Set and other methods.

[Package net](https://go.dev/pkg/net/) drops the `laddr` argument from [Dial](https://go.dev/pkg/net/#Conn.Dial) and drops the `cname` return value from [LookupHost](https://go.dev/pkg/net/#LookupHost). The implementation now uses [cgo](https://go.dev/cmd/cgo/) to implement network name lookups using the C library getaddrinfo(3) function when possible. This ensures that Go and C programs resolve names the same way and also avoids the OS X application-level firewall.

[Package os](https://go.dev/pkg/os/) introduces simplified [Open](https://go.dev/pkg/os/#File.Open) and [Create](https://go.dev/pkg/os/#File.Create) functions. The original Open is now available as [OpenFile](https://go.dev/pkg/os/#File.OpenFile). The final three arguments to [StartProcess](https://go.dev/pkg/os/#Process.StartProcess) have been replaced by a pointer to a [ProcAttr](https://go.dev/pkg/os/#ProcAttr).

[Package reflect](https://go.dev/pkg/reflect/) has been redesigned. [Type](https://go.dev/pkg/reflect/#Type) is now an interface that implements all the possible type methods. Instead of a type switch on a Type `t`, switch on `t.Kind()`. [Value](https://go.dev/pkg/reflect/#Value) is now a struct value that implements all the possible value methods. Instead of a type switch on a Value `v`, switch on `v.Kind()`. Typeof and NewValue are now called [TypeOf](https://go.dev/pkg/reflect/#Type.TypeOf) and [ValueOf](https://go.dev/pkg/reflect/#Value.ValueOf) To create a writable Value, use `New(t).Elem()` instead of `Zero(t)`. See [the change description](https://go.dev/change/843855f3c026) for the full details. The new API allows a more efficient implementation of Value that avoids many of the allocations required by the previous API.

Remember that gofix will handle the bulk of the rewrites necessary for these changes to package APIs.

### Tools

[Gofix](https://go.dev/cmd/gofix/), a new command, is described above.

[Gotest](https://go.dev/cmd/gotest/) is now a Go program instead of a shell script. The new `-test.short` flag in combination with package testing's Short function allows you to write tests that can be run in normal or "short" mode; all.bash runs tests in short mode to reduce installation time. The Makefiles know about the flag: use `make testshort`.

The run-time support now implements CPU and memory profiling. Gotest's new [`-test.cpuprofile` and `-test.memprofile` flags](https://go.dev/cmd/gotest/) make it easy to profile tests. To add profiling to your web server, see the [http/pprof](https://go.dev/pkg/http/pprof/) documentation. For other uses, see the [runtime/pprof](https://go.dev/pkg/runtime/pprof/) documentation.

### Minor revisions

r57.1 fixes a [nil pointer dereference in http.FormFile](https://go.dev/change/ff2bc62726e7145eb2ecc1e0f076998e4a8f86f0).

r57.2 fixes a [use of uninitialized memory in programs that misuse `goto`](https://go.dev/change/063b0ff67d8277df03c956208abc068076818dae).

## r56 (released 2011/03/16)

The r56 release was the first stable release and corresponds to `weekly.2011-03-07.1`. The numbering starts at 56 because before this release, what we now consider weekly snapshots were called releases.