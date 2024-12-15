+++
title = "Features Table"
date = 2024-12-15T11:21:13+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Features-Table](https://github.com/smartystreets/goconvey/wiki/Features-Table)
>
> 收录该文档时间： `2024-12-15T11:21:13+08:00`

# Features Table



Marcos Nils edited this page on May 24, 2015 · [1 revision](https://github.com/smartystreets/goconvey/wiki/Features-Table/_history)

> **✓** = Has it
>
> **~** = Kind of has it
>
> **?** = Not sure yet

| Feature                             | [GoConvey](https://github.com/smartystreets/goconvey) | Native [`go test`](http://golang.org/pkg/testing/) | [Testify](https://github.com/stretchr/testify) | [goblin](https://github.com/franela/goblin) | [PrettyTest](https://github.com/remogatto/prettytest) | [Ginkgo](https://github.com/onsi/ginkgo) |
| ----------------------------------- | ----------------------------------------------------- | -------------------------------------------------- | ---------------------------------------------- | ------------------------------------------- | ----------------------------------------------------- | ---------------------------------------- |
| Uses `go test`                      | ✓                                                     | ✓                                                  | ✓                                              | ✓                                           | ✓                                                     | ~                                        |
| Web UI                              | ✓                                                     |                                                    |                                                |                                             |                                                       |                                          |
| Web UI reports traditional Go tests | ✓                                                     |                                                    |                                                |                                             |                                                       |                                          |
| Open files in Sublime Text          | ✓                                                     |                                                    |                                                |                                             |                                                       |                                          |
| Auto-test                           | ✓                                                     |                                                    | ?                                              | ?                                           | ✓                                                     |                                          |
| Test code generator                 | ✓                                                     |                                                    |                                                |                                             |                                                       | ~                                        |
| Custom assertions/matchers          | ✓                                                     | ✓                                                  | ?                                              | ~                                           | ✓                                                     | ✓                                        |
| Optional verbose output             | ✓                                                     | ✓                                                  | ?                                              | ?                                           | ✓                                                     | ✓                                        |
| Colorized console output            | ✓                                                     |                                                    | ?                                              | ✓                                           | ✓                                                     | ✓                                        |
| Non-IT-readable output              | ✓                                                     |                                                    |                                                | ✓                                           | ✓                                                     | ✓                                        |
| Randomized test execution           |                                                       |                                                    |                                                |                                             |                                                       | ✓                                        |
| Coverage report                     | ✓                                                     |                                                    |                                                |                                             |                                                       | ✓                                        |
| Skip test blocks/assertions         | ✓                                                     | ✓                                                  | ?                                              | ✓                                           | ?                                                     | ✓                                        |
| Flexible DSL                        | ✓                                                     | ✓                                                  |                                                |                                             |                                                       |                                          |
| Supports BDD/TDD/Acceptance         | ✓                                                     |                                                    |                                                | ✓                                           |                                                       | ✓                                        |
| **Assertions**                      |                                                       |                                                    |                                                |                                             |                                                       |                                          |
| - Equal                             | ✓                                                     |                                                    | ✓                                              | ✓                                           | ✓                                                     | ✓                                        |
| - DeepEqual                         | ✓                                                     |                                                    |                                                | ✓                                           | ✓                                                     | ✓                                        |
| - True                              | ✓                                                     |                                                    | ✓                                              |                                             | ✓                                                     | ✓                                        |
| - False                             | ✓                                                     |                                                    | ✓                                              |                                             | ✓                                                     | ✓                                        |
| - Nil                               | ✓                                                     |                                                    | ✓                                              |                                             | ✓                                                     | ✓                                        |
| - Empty                             | ✓                                                     |                                                    | ✓                                              |                                             |                                                       | ✓                                        |
| - (a whole bunch more)              | ✓                                                     |                                                    | ✓                                              |                                             | ✓                                                     | ✓                                        |

*Some table data compiled from [Go Test It](https://github.com/shageman/gotestit)*
