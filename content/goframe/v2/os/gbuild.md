+++
title = "gbuild"
date = 2024-03-21T17:54:25+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gbuild

Package gbuild manages the build-in variables from "gf build".

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gbuild/gbuild.go#L31)

``` go
const (
	BuiltGit     = `builtGit`
	BuiltTime    = `builtTime`
	BuiltVersion = `builtVersion`
)
```

### Variables 

This section is empty.

### Functions 

##### func Data 

``` go
func Data() map[string]interface{}
```

Data returns the custom build-in variables as map.

##### func Get 

``` go
func Get(name string, def ...interface{}) *gvar.Var
```

Get retrieves and returns the build-in binary variable with given name.

### Types 

#### type BuildInfo 

``` go
type BuildInfo struct {
	GoFrame string                 // Built used GoFrame version.
	Golang  string                 // Built used Golang version.
	Git     string                 // Built used git repo. commit id and datetime.
	Time    string                 // Built datetime.
	Version string                 // Built version.
	Data    map[string]interface{} // All custom built data key-value pairs.
}
```

BuildInfo maintains the built info of current binary.

##### func Info 

``` go
func Info() BuildInfo
```

Info returns the basic built information of the binary as map. Note that it should be used with gf-cli tool "gf build", which automatically injects necessary information into the binary.