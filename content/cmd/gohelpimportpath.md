+++
title = "go help importpath"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

An import path (see 'go help packages') denotes a package stored in the local file system. In general, an import path denotes either a standard package (such as "unicode/utf8") or a package found in one of the work spaces (For more details see: 'go help gopath').

Relative import paths

An import path beginning with ./ or ../ is called a relative path. The toolchain supports relative import paths as a shortcut in two ways.

First, a relative path can be used as a shorthand on the command line. If you are working in the directory containing the code imported as "unicode" and want to run the tests for "unicode/utf8", you can type "go test ./utf8" instead of needing to specify the full path.

Similarly, in the reverse situation, "go test .." will test "unicode" from the "unicode/utf8" directory. Relative patterns are also allowed, like "go test ./..." to test all subdirectories. See 'go help packages' for details on the pattern syntax.

Second, if you are compiling a Go program not in a work space, you can use a relative path in an import statement in that program to refer to nearby code also not in a work space.

This makes it easy to experiment with small multipackage programs outside of the usual work spaces, but such programs cannot be installed with "go install" (there is no work space in which to install them),so they are rebuilt from scratch each time they are built.

To avoid ambiguity, Go programs cannot use relative import paths within a work space.

Remote import paths

Certain import paths also describe how to obtain the source code for the package using a revision control system.

A few common code hosting sites have special syntax:

        Bitbucket (Git, Mercurial)
    
                import "bitbucket.org/user/project"
                import "bitbucket.org/user/project/sub/directory"
    
        GitHub (Git)
    
                import "github.com/user/project"
                import "github.com/user/project/sub/directory"
    
        Launchpad (Bazaar)
    
                import "launchpad.net/project"
                import "launchpad.net/project/series"
                import "launchpad.net/project/series/sub/directory"
    
                import "launchpad.net/~user/project/branch"
                import "launchpad.net/~user/project/branch/sub/directory"
    
        IBM DevOps Services (Git)
    
                import "hub.jazz.net/git/user/project"
                import "hub.jazz.net/git/user/project/sub/directory"

For code hosted on other servers, import paths may either be qualified with the version control type, or the go tool can dynamically fetch the import path over https/http and discover where the code resides from a `<meta>` tag in the HTML.

To declare the code location, an import path of the form

        repository.vcs/path

specifies the given repository, with or without the .vcs suffix, using the named version control system, and then the path inside that repository. The supported version control systems are:

        Bazaar      .bzr
        Fossil      .fossil
        Git         .git
        Mercurial   .hg
        Subversion  .svn

For example,

```go
import "example.org/user/foo.hg"
```

denotes the root directory of the Mercurial repository at example.org/user/foo or foo.hg, and

```go
import "example.org/repo.git/foo/bar"
```

denotes the foo/bar directory of the Git repository at example.org/repo or repo.git.

When a version control system supports multiple protocols, each is tried in turn when downloading. For example, a Git download tries https://, then git+ssh://.

By default, downloads are restricted to known secure protocols (e.g. https, ssh). To override this setting for Git downloads, the GIT_ALLOW_PROTOCOL environment variable can be set (For more details see: 'go help environment').

If the import path is not a known code hosting site and also lacks a version control qualifier, the go tool attempts to fetch the import over https/http and looks for a `<meta>` tag in the document's HTML `<head>`.


The meta tag has the form:

```html
<meta name="go-import" content="import-prefix vcs repo-root">
```

The import-prefix is the import path corresponding to the repository root. It must be a prefix or an exact match of the package being fetched with "go get". If it's not an exact match, another http request is made at the prefix to verify the <meta> tags match.

The meta tag should appear as early in the file as possible.
In particular, it should appear before any raw JavaScript or CSS, to avoid confusing the go command's restricted parser.

The vcs is one of "bzr", "fossil", "git", "hg", "svn".

The repo-root is the root of the version control system containing a scheme and not containing a .vcs qualifier.

For example,

```go
import "example.org/pkg/foo"
```

will result in the following requests:

        https://example.org/pkg/foo?go-get=1 (preferred)
        http://example.org/pkg/foo?go-get=1  (fallback, only with use of correctly set GOINSECURE)

If that page contains the meta tag

```html
<meta name="go-import" content="example.org git https://code.org/r/p/exproj">
```

the go tool will verify that https://example.org/?go-get=1 contains the same meta tag and then git clone https://code.org/r/p/exproj into GOPATH/src/example.org.

When using GOPATH, downloaded packages are written to the first directory listed in the GOPATH environment variable. (See 'go help gopath-get' and 'go help gopath'.)

When using modules, downloaded packages are stored in the module cache. See https://golang.org/ref/mod#module-cache.

When using modules, an additional variant of the go-import meta tag is recognized and is preferred over those listing version control systems. That variant uses "mod" as the vcs in the content value, as in:

```html
<meta name="go-import" content="example.org mod https://code.org/moduleproxy">
```

This tag means to fetch modules with paths beginning with example.org from the module proxy available at the URL https://code.org/moduleproxy. See https://golang.org/ref/mod#goproxy-protocol for details about the proxy protocol.

Import path checking

When the custom import path feature described above redirects to a known code hosting site, each of the resulting packages has two possible import paths, using the custom domain or the known hosting site.

A package statement is said to have an "import comment" if it is immediately followed (before the next newline) by a comment of one of these two forms:

```go
    package math // import "path"
    package math /* import "path" */
```

The go command will refuse to install a package with an import comment unless it is being referred to by that import path. In this way, import comments let package authors make sure the custom import path is used and not a direct path to the underlying code hosting site.

Import path checking is disabled for code found within vendor trees. This makes it possible to copy code into alternate locations in vendor trees without needing to update import comments.

Import path checking is also disabled when using modules. Import path comments are obsoleted by the go.mod file's module statement.

See https://golang.org/s/go14customimport for details.
