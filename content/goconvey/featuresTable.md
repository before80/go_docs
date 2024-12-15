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

# Features Table- 功能对比表



Marcos Nils edited this page on May 24, 2015 · [1 revision](https://github.com/smartystreets/goconvey/wiki/Features-Table/_history)

​	Marcos Nils 于 2015 年 5 月 24 日编辑了此页面 · [1 次修订](https://github.com/smartystreets/goconvey/wiki/Features-Table/_history)

> **✓** = Has it 支持
>
> **~** = Kind of has it 部分支持
>
> **?** = Not sure yet 暂不确定

| Feature                                                      | [GoConvey](https://github.com/smartystreets/goconvey) | Native [`go test`](http://golang.org/pkg/testing/) | [Testify](https://github.com/stretchr/testify) | [goblin](https://github.com/franela/goblin) | [PrettyTest](https://github.com/remogatto/prettytest) | [Ginkgo](https://github.com/onsi/ginkgo) |
| ------------------------------------------------------------ | ----------------------------------------------------- | -------------------------------------------------- | ---------------------------------------------- | ------------------------------------------- | ----------------------------------------------------- | ---------------------------------------- |
| Uses `go test` 使用 `go test`                                | ✓                                                     | ✓                                                  | ✓                                              | ✓                                           | ✓                                                     | ~                                        |
| Web UI                                                       | ✓                                                     |                                                    |                                                |                                             |                                                       |                                          |
| Web UI reports traditional Go tests - Web UI 报告传统 Go 测试 | ✓                                                     |                                                    |                                                |                                             |                                                       |                                          |
| Open files in Sublime Text 在 Sublime Text 中打开文件        | ✓                                                     |                                                    |                                                |                                             |                                                       |                                          |
| Auto-test 自动测试                                           | ✓                                                     |                                                    | ?                                              | ?                                           | ✓                                                     |                                          |
| Test code generator 测试代码生成器                           | ✓                                                     |                                                    |                                                |                                             |                                                       | ~                                        |
| Custom assertions/matchers 自定义断言/匹配器                 | ✓                                                     | ✓                                                  | ?                                              | ~                                           | ✓                                                     | ✓                                        |
| Optional verbose output 可选详细输出                         | ✓                                                     | ✓                                                  | ?                                              | ?                                           | ✓                                                     | ✓                                        |
| Colorized console output 彩色控制台输出                      | ✓                                                     |                                                    | ?                                              | ✓                                           | ✓                                                     | ✓                                        |
| Non-IT-readable output 非 IT 人员可读的输出                  | ✓                                                     |                                                    |                                                | ✓                                           | ✓                                                     | ✓                                        |
| Randomized test execution 随机化测试执行                     |                                                       |                                                    |                                                |                                             |                                                       | ✓                                        |
| Coverage report 覆盖率报告                                   | ✓                                                     |                                                    |                                                |                                             |                                                       | ✓                                        |
| Skip test blocks/assertions 跳过测试块/断言                  | ✓                                                     | ✓                                                  | ?                                              | ✓                                           | ?                                                     | ✓                                        |
| Flexible DSL 灵活的 DSL                                      | ✓                                                     | ✓                                                  |                                                |                                             |                                                       |                                          |
| Supports BDD/TDD/Acceptance 支持 BDD/TDD/验收测试            | ✓                                                     |                                                    |                                                | ✓                                           |                                                       | ✓                                        |
| **Assertions** 断言支持                                      |                                                       |                                                    |                                                |                                             |                                                       |                                          |
| - Equal 相等                                                 | ✓                                                     |                                                    | ✓                                              | ✓                                           | ✓                                                     | ✓                                        |
| - DeepEqual 深度相等                                         | ✓                                                     |                                                    |                                                | ✓                                           | ✓                                                     | ✓                                        |
| - True 为真                                                  | ✓                                                     |                                                    | ✓                                              |                                             | ✓                                                     | ✓                                        |
| - False 为假                                                 | ✓                                                     |                                                    | ✓                                              |                                             | ✓                                                     | ✓                                        |
| - Nil                                                        | ✓                                                     |                                                    | ✓                                              |                                             | ✓                                                     | ✓                                        |
| - Empty 为空                                                 | ✓                                                     |                                                    | ✓                                              |                                             |                                                       | ✓                                        |
| - (a whole bunch more) 更多断言                              | ✓                                                     |                                                    | ✓                                              |                                             | ✓                                                     | ✓                                        |

*Some table data compiled from [Go Test It](https://github.com/shageman/gotestit)*

​	部分表格数据来自 [Go Test It](https://github.com/shageman/gotestit)
