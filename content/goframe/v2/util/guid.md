+++
title = "guid"
date = 2024-03-21T17:59:59+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/guid](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/guid)

Package guid provides simple and high performance unique id generation functionality.

​	包 guid 提供简单且高性能的唯一 ID 生成功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func S

```go
func S(data ...[]byte) string
```

S creates and returns a global unique string in 32 bytes that meets most common usages without strict UUID algorithm. It returns a unique string using default unique algorithm if no `data` is given.

​	S 创建并返回一个 32 字节的全局唯一字符串，该字符串满足大多数常见用法，无需严格的 UUID 算法。如果给出 no `data` ，则使用默认的唯一算法返回一个唯一的字符串。

The specified `data` can be no more than 2 parts. No matter how long each of the `data` size is, each of them will be hashed into 7 bytes as part of the result. If given `data` parts is less than 2, the leftover size of the result bytes will be token by random string.

​	指定的 `data` 不能超过 2 个部分。无论每个 `data` 大小有多长，它们中的每一个都将作为结果的一部分被哈希为 7 个字节。如果给定 `data` 的部分小于 2，则结果字节的剩余大小将由随机字符串标记。

The returned string is composed with: 1. Default: MACHash(7) + PID(4) + TimestampNano(12) + Sequence(3) + RandomString(6) 2. CustomData: DataHash(7/14) + TimestampNano(12) + Sequence(3) + RandomString(3/10)

​	返回的字符串由以下部分组成： 1. 默认值：MACHash（7） + PID（4） + TimestampNano（12） + Sequence（3） + RandomString（6） 2.自定义数据：DataHash（7/14） + TimestampNano（12） + Sequence（3） + RandomString（3/10）

Note that：

​	请注意：

1. The returned length is fixed to 32 bytes for performance purpose.
   出于性能目的，返回的长度固定为 32 字节。
2. The custom parameter `data` composed should have unique attribute in your business scenario.
   `data` 编写的自定义参数在业务场景中应具有唯一属性。

## 类型

This section is empty.