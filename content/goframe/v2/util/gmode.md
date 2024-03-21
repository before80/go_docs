+++
title = "gmode"
date = 2024-03-21T17:59:37+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gmode

Package gmode provides release mode management for project.

It uses string to mark the mode instead of integer, which is convenient for configuration.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gmode/gmode.go#L18)

``` go
const (
	NOT_SET = "not-set"
	DEVELOP = "develop"
	TESTING = "testing"
	STAGING = "staging"
	PRODUCT = "product"
)
```

### Variables 

This section is empty.

### Functions 

##### func IsDevelop 

``` go
func IsDevelop() bool
```

IsDevelop checks and returns whether current application is running in DEVELOP mode.

##### func IsProduct 

``` go
func IsProduct() bool
```

IsProduct checks and returns whether current application is running in PRODUCT mode.

##### func IsStaging 

``` go
func IsStaging() bool
```

IsStaging checks and returns whether current application is running in STAGING mode.

##### func IsTesting 

``` go
func IsTesting() bool
```

IsTesting checks and returns whether current application is running in TESTING mode.

##### func Mode 

``` go
func Mode() string
```

Mode returns current application mode set.

##### func Set 

``` go
func Set(mode string)
```

Set sets the mode for current application.

##### func SetDevelop 

``` go
func SetDevelop()
```

SetDevelop sets current mode DEVELOP for current application.

##### func SetProduct 

``` go
func SetProduct()
```

SetProduct sets current mode PRODUCT for current application.

##### func SetStaging 

``` go
func SetStaging()
```

SetStaging sets current mode STAGING for current application.

##### func SetTesting 

``` go
func SetTesting()
```

SetTesting sets current mode TESTING for current application.

### Types 

This section is empty.