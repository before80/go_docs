+++
title = "第三方包"
date = 2024-07-29T08:05:29+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 框架

### web框架

#### gin

[https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

Gin is a web framework written in [Go](https://go.dev/). It features a martini-like API with performance that is up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter). If you need performance and good productivity, you will love Gin.

## 概况型

### 数据类型

#### constraints

> import方式："golang.org/x/exp/constraints"
>
> 所在代码仓库：[https://cs.opensource.google/go/x/exp](https://cs.opensource.google/go/x/exp)

Package constraints defines a set of useful constraints to be used with type parameters.

## 指导作用

### 代码布局

#### project-layout

[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

This is a basic layout for Go application projects. Note that it's basic in terms of content because it's focusing only on the general layout and not what you have inside. It's also basic because it's very high level and it doesn't go into great details in terms of how you can structure your project even further. For example, it doesn't try to cover the project structure you'd have with something like Clean Architecture.

This is **`NOT an official standard defined by the core Go dev team`**. This is a set of common historical and emerging project layout patterns in the Go ecosystem. Some of these patterns are more popular than others. It also has a number of small enhancements along with several supporting directories common to any large enough real world application. Note that the **core Go team provides a great set of general guidelines about structuring Go projects** and what it means for your project when it's imported and when it's installed. See the [`Organizing a Go module`](https://go.dev/doc/modules/layout) page in the official Go docs for more details. It includes the `internal` and `cmd` directory patterns (described below) and other useful information.

**`If you are trying to learn Go or if you are building a PoC or a simple project for yourself this project layout is an overkill. Start with something really simple instead (a single `main.go`file and`go.mod` is more than enough).`** As your project grows keep in mind that it'll be important to make sure your code is well structured otherwise you'll end up with a messy code with lots of hidden dependencies and global state. When you have more people working on the project you'll need even more structure. That's when it's important to introduce a common way to manage packages/libraries. When you have an open source project or when you know other projects import the code from your project repository that's when it's important to have private (aka `internal`) packages and code. Clone the repository, keep what you need and delete everything else! Just because it's there it doesn't mean you have to use it all. None of these patterns are used in every single project. Even the `vendor` pattern is not universal.

With Go 1.14 [`Go Modules`](https://go.dev/wiki/Modules) are finally ready for production. Use [`Go Modules`](https://blog.golang.org/using-go-modules) unless you have a specific reason not to use them and if you do then you don’t need to worry about $GOPATH and where you put your project. The basic `go.mod` file in the repo assumes your project is hosted on GitHub, but it's not a requirement. The module path can be anything though the first module path component should have a dot in its name (the current version of Go doesn't enforce it anymore, but if you are using slightly older versions don't be surprised if your builds fail without it). See Issues [`37554`](https://github.com/golang/go/issues/37554) and [`32819`](https://github.com/golang/go/issues/32819) if you want to know more about it.

This project layout is intentionally generic and it doesn't try to impose a specific Go package structure.

This is a community effort. Open an issue if you see a new pattern or if you think one of the existing patterns needs to be updated.

If you need help with naming, formatting and style start by running [`gofmt`](https://golang.org/cmd/gofmt/) and [`staticcheck`](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck). The previous standard linter, golint, is now deprecated and not maintained; use of a maintained linter such as staticcheck is recommended. Also make sure to read these Go code style guidelines and recommendations:

- https://talks.golang.org/2014/names.slide
- https://golang.org/doc/effective_go.html#names
- https://blog.golang.org/package-names
- https://go.dev/wiki/CodeReviewComments
- [Style guideline for Go packages](https://rakyll.org/style-packages) (rakyll/JBD)

See [`Go Project Layout`](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) for additional background information.

More about naming and organizing packages as well as other code structure recommendations:

- [GopherCon EU 2018: Peter Bourgon - Best Practices for Industrial Programming](https://www.youtube.com/watch?v=PTE4VJIdHPg)
- [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices.](https://www.youtube.com/watch?v=MzTcsI6tn-0)
- [GopherCon 2017: Edward Muller - Go Anti-Patterns](https://www.youtube.com/watch?v=ltqV6pDKZD8)
- [GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)

A Chinese post about Package-Oriented-Design guidelines and Architecture layer

- [面向包的设计和架构分层](https://github.com/danceyoung/paper-code/blob/master/package-oriented-design/packageorienteddesign.md)

## 某类操作

### 管理命令

#### cobra

[https://github.com/spf13/cobra](https://github.com/spf13/cobra)

Cobra is a library for creating powerful modern CLI applications.

Cobra is used in many Go projects such as [Kubernetes](https://kubernetes.io/), [Hugo](https://gohugo.io/), and [GitHub CLI](https://github.com/cli/cli) to name a few. [This list](https://github.com/spf13/cobra/blob/main/site/content/projects_using_cobra.md) contains a more extensive list of projects using Cobra.

### 管理配置

#### viper

[https://github.com/spf13/viper](https://github.com/spf13/viper)

Viper is a complete configuration solution for Go applications including [12-Factor apps](https://12factor.net/#the_twelve_factors). It is designed to work within an application, and can handle all types of configuration needs and formats. It supports:

- setting defaults
- reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
- live watching and re-reading of config files (optional)
- reading from environment variables
- reading from remote config systems (etcd or Consul), and watching changes
- reading from command line flags
- reading from buffer
- setting explicit values

Viper can be thought of as a registry for all of your applications configuration needs.

### 操作数据库

#### gorm

[https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)

The fantastic ORM library for Golang, aims to be developer friendly.

### 记录日志

#### zap

[https://github.com/uber-go/zap](https://github.com/uber-go/zap)

Blazing fast, structured, leveled logging in Go.

#### logrus

[https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.

### 验证

#### validator

[https://github.com/go-playground/validator](https://github.com/go-playground/validator)

Package validator implements value validations for structs and individual fields based on tags.

It has the following **unique** features:

- Cross Field and Cross Struct validations by using validation tags or custom validators.
- Slice, Array and Map diving, which allows any or all levels of a multidimensional field to be validated.
- Ability to dive into both map keys and values for validation
- Handles type interface by determining it's underlying type prior to validation.
- Handles custom field types such as sql driver Valuer see [Valuer](https://golang.org/src/database/sql/driver/types.go?s=1210:1293#L29)
- Alias validation tags, which allows for mapping of several validations to a single tag for easier defining of validations on structs
- Extraction of custom defined Field Name e.g. can specify to extract the JSON name while validating and have it available in the resulting FieldError
- Customizable i18n aware error messages.
- Default validator for the [gin](https://github.com/gin-gonic/gin) web framework; upgrading from v8 to v9 in gin see [here](https://github.com/go-playground/validator/tree/master/_examples/gin-upgrading-overriding)



### 测试

#### testify

[https://github.com/stretchr/testify](https://github.com/stretchr/testify)

A toolkit with common assertions and mocks that plays nicely with the standard library。



### 产生、获取特定数据

#### UUID

[https://github.com/google/uuid](https://github.com/google/uuid)

​	The uuid package generates and inspects UUIDs based on [RFC 9562](https://datatracker.ietf.org/doc/html/rfc9562) and DCE 1.1: Authentication and Security Services.

​	This package is based on the github.com/pborman/uuid package (previously named code.google.com/p/go-uuid). It differs from these earlier packages in that a UUID is a 16 byte array rather than a byte slice. One loss due to this change is the ability to represent an invalid UUID (vs a NIL UUID).



#### goid

[https://github.com/petermattis/goid](https://github.com/petermattis/goid)

Programatically retrieve the current goroutine's ID. See [the CI configuration](https://github.com/petermattis/goid/blob/master/.github/workflows/go.yml) for supported Go versions.

