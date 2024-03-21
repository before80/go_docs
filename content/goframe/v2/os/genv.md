+++
title = "genv"
date = 2024-03-21T17:55:23+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/genv

Package genv provides operations for environment variables of system.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func All 

``` go
func All() []string
```

All returns a copy of strings representing the environment, in the form "key=value".

##### func Build 

``` go
func Build(m map[string]string) []string
```

Build builds a map to an environment variable slice.

##### func Contains 

``` go
func Contains(key string) bool
```

Contains checks whether the environment variable named `key` exists.

##### func Filter <-2.1.0

``` go
func Filter(envs []string) []string
```

Filter filters repeated items from given environment variables.

##### func Get 

``` go
func Get(key string, def ...interface{}) *gvar.Var
```

Get creates and returns a Var with the value of the environment variable named by the `key`. It uses the given `def` if the variable does not exist in the environment.

##### func GetWithCmd 

``` go
func GetWithCmd(key string, def ...interface{}) *gvar.Var
```

GetWithCmd returns the environment value specified `key`. If the environment value does not exist, then it retrieves and returns the value from command line options. It returns the default value `def` if none of them exists.

Fetching Rules: 1. Environment arguments are in uppercase format, eg: GF_<package name>_<variable name>； 2. Command line arguments are in lowercase format, eg: gf.<package name>.<variable name>;

##### func Map 

``` go
func Map() map[string]string
```

Map returns a copy of strings representing the environment as a map.

##### func MapFromEnv <-2.1.0

``` go
func MapFromEnv(envs []string) map[string]string
```

MapFromEnv converts environment variables from slice to map.

##### func MapToEnv <-2.1.0

``` go
func MapToEnv(m map[string]string) []string
```

MapToEnv converts environment variables from map to slice.

##### func MustRemove 

``` go
func MustRemove(key ...string)
```

MustRemove performs as Remove, but it panics if any error occurs.

##### func MustSet 

``` go
func MustSet(key, value string)
```

MustSet performs as Set, but it panics if any error occurs.

##### func Remove 

``` go
func Remove(key ...string) (err error)
```

Remove deletes one or more environment variables.

##### func Set 

``` go
func Set(key, value string) (err error)
```

Set sets the value of the environment variable named by the `key`. It returns an error, if any.

##### func SetMap 

``` go
func SetMap(m map[string]string) (err error)
```

SetMap sets the environment variables using map.

### Types 

This section is empty.