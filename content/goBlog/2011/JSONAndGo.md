+++
title = "JSON 和 go"
weight = 28
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# JSON and Go - JSON 和 go

> 原文：[https://go.dev/blog/json](https://go.dev/blog/json)

Andrew Gerrand
25 January 2011

2011年1月25日

## 介绍

​	JSON（JavaScript 对象表示法）是一种简单的数据交换格式。语法上类似于 JavaScript 中的对象和列表。它通常用于 Web 后端与运行在浏览器中的 JavaScript 程序之间的通信，但它在许多其他地方也被使用。它的主页 [json.org](http://json.org/) 提供了一个非常清晰和简明的标准定义。

​	使用 [json 包](https://go.dev/pkg/encoding/json/)，我们可以轻松地在 Go 程序中读取和写入 JSON 数据。

## 编码

​	要编码 JSON 数据，我们使用 Marshal 函数。

```go
func Marshal(v interface{}) ([]byte, error)
```

​	假设我们有 Go 数据结构体 Message，

```go
type Message struct {
    Name string
    Body string
    Time int64
}
```

和 Message 的一个实例

```
m := Message{"Alice", "Hello", 1294706395881547000}
```

​	我们可以使用 json.Marshal 将 m 编码为 JSON 格式：

```go
b, err := json.Marshal(m)
```

​	如果一切顺利，err 将为 nil，b 将包含以下 JSON 数据的 []byte：

```go
b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
```

​	只有那些可以表示为有效 JSON 的数据结构才能被编码：

- JSON 对象只支持字符串作为键；要编码 Go 的映射类型，它必须采用 map[string]T 的形式（其中 T 是 json 包支持的任何 Go 类型）。
- 通道、复数和函数类型无法编码。
- 不支持循环数据结构；它们将导致 Marshal 进入无限循环。
- 指针将被编码为它们指向的值（或如果指针为 nil，则为"null"）。

​	json 包仅访问结构体类型的公开字段（那些以大写字母开头的字段）。因此，结构体的仅公开字段将存在于 JSON 输出中。

## 解码

​	要解码JSON数据，我们使用Unmarshal函数。

```go
func Unmarshal(data []byte, v interface{}) error
```

​	我们必须先创建一个地方来存储解码后的数据

```go
var m Message
```

​	并调用json.Unmarshal，将一个[]byte的JSON数据和指向m的指针传递给它

```go
err := json.Unmarshal(b, &m)
```

​	如果b包含适合m的有效JSON，则在调用后err将为nil，并且数据将像通过赋值一样存储在结构体m中：

```go
m = Message{
    Name: "Alice",
    Body: "Hello",
    Time: 1294706395881547000,
}
```

​	Unmarshal如何确定在哪些字段中存储解码后的数据？对于给定的JSON键"Foo"，Unmarshal将查找目标结构的字段以查找（按优先顺序）：

- 具有标记"Foo"的导出字段（有关结构标记的更多信息，请参见[Go规范]({{< ref "/langSpec/Types#struct-types">}})）， 
- 一个名为"Foo"的导出字段，或
- 名为"FOO"或"FoO"或"Foo"的其他大小写不敏感匹配的导出字段。

​	当JSON数据的结构与Go类型不完全匹配时会发生什么？

```go
b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
var m Message
err := json.Unmarshal(b, &m)
```

​	Unmarshal将仅解码可以在目标类型中找到的字段。在这种情况下，只会填充m的Name字段，而忽略Food字段。当您想从一个大的JSON blob中只挑选一些特定字段时，这种行为特别有用。它还意味着目标结构中的任何未导出字段都不会受到Unmarshal的影响。

​	但是如果您事先不知道JSON数据的结构呢？

## 用接口处理通用JSON

​	`interface{}`（空接口）类型描述了一个带有零个方法的接口。每种Go类型都实现了至少零个方法，因此满足空接口。

​	空接口用作通用容器类型：

```go
var i interface{}
i = "a string"
i = 2011
i = 2.777
```

​	类型断言访问底层具体类型：

```go
r := i.(float64)
fmt.Println("the circle's area", math.Pi*r*r)
```

​	或者，如果底层类型未知，则类型切换确定类型：

```go
switch v := i.(type) {
case int:
    fmt.Println("twice i is", v*2)
case float64:
    fmt.Println("the reciprocal of i is", 1/v)
case string:
    h := len(v) / 2
    fmt.Println("i swapped by halves is", v[h:]+v[:h])
default:
    // i isn't one of the types above
}
```

​	json包使用`map[string]interface{}`和`[]interface{}`值存储任意的JSON对象和数组。它可以将任何有效的JSON数据转换为普通的`interface{}`值。默认的具体Go类型是：

- `bool`用于JSON布尔类型
- `float64`用于JSON数字类型
- `string`用于JSON字符串类型
- `nil`用于JSON空值类型

## 解码任意数据

​	考虑以下存储在变量b中的JSON数据：

```go
b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
```

​	在不知道数据结构的情况下，我们可以使用`Unmarshal`将其解码为一个`interface{}`值：

```go
var f interface{}
err := json.Unmarshal(b, &f)
```

​	此时f中的Go值将是一个映射，其键为字符串，其值本身存储为空接口值：

```go
f = map[string]interface{}{
    "Name": "Wednesday",
    "Age":  6,
    "Parents": []interface{}{
        "Gomez",
        "Morticia",
    },
}
```

​	我们可以使用类型断言来访问f的底层map[string]interface{}中的数据：

```go
m := f.(map[string]interface{})
```

​	然后我们可以使用一个range语句遍历该映射，并使用类型选择将其值访问为其具体类型：

```go
for k, v := range m {
    switch vv := v.(type) {
    case string:
        fmt.Println(k, "is string", vv)
    case float64:
        fmt.Println(k, "is float64", vv)
    case []interface{}:
        fmt.Println(k, "is an array:")
        for i, u := range vv {
            fmt.Println(i, u)
        }
    default:
        fmt.Println(k, "is of a type I don't know how to handle")
    }
}
```

​	通过这种方式，您可以使用未知的JSON数据，同时仍然享受类型安全的好处。

## 参考类型

​	让我们定义一个Go类型来包含前面示例的数据：

```go
type FamilyMember struct {
    Name    string
    Age     int
    Parents []string
}

var m FamilyMember
err := json.Unmarshal(b, &m)
```

​	将数据反序列化为FamilyMember值的结果与预期相符，但如果我们仔细观察，可以发现发生了一件非常特别的事情。使用var语句我们分配了一个FamilyMember结构体，然后将指向该值的指针提供给Unmarshal，但此时Parents字段是nil切片值。为了填充Parents字段，Unmarshal在幕后分配了一个新切片。这是Unmarshal与支持的引用类型（指针、切片和映射）一起工作的典型方式。

​	考虑将其反序列化为以下数据结构：

```go
type Foo struct {
    Bar *Bar
}
```

​	如果JSON对象中有Bar字段，则Unmarshal将分配一个新的Bar并填充它。如果没有，则Bar将保留为nil指针。

​	由此产生了一种有用的模式：如果您的应用程序接收几种不同类型的消息，则可以定义一个类似"接收器"的结构体，如下所示：

```go
type IncomingMessage struct {
    Cmd *Command
    Msg *Message
}
```

​	发送方可以填充顶层JSON对象的Cmd字段和/或Msg字段，具体取决于他们要通信的消息类型。当将JSON解码为IncomingMessage结构时，Unmarshal将仅分配存在于JSON数据中的数据结构。为了知道要处理哪些消息，程序员只需简单地测试Cmd或Msg是否为nil即可。

## 流式编码器和解码器

​	json包提供Decoder和Encoder类型来支持读取和写入JSON数据流的常见操作。NewDecoder和NewEncoder函数包装了[io.Reader](https://go.dev/pkg/io/#Reader)和[io.Writer](https://go.dev/pkg/io/#Writer)接口类型。

```go
func NewDecoder(r io.Reader) *Decoder
func NewEncoder(w io.Writer) *Encoder
```

​	这是一个示例程序，从标准输入读取一系列JSON对象，从每个对象中删除所有字段但Name字段，然后将对象写入标准输出：

```go
package main

import (
    "encoding/json"
    "log"
    "os"
)

func main() {
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            if k != "Name" {
                delete(v, k)
            }
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}
```

​	由于读取器和写入器的普遍性，这些编码器和解码器类型可以用于广泛的场景，例如读写HTTP连接、WebSockets或文件。

## 参考文献

​	更多信息请参见[json包文档](https://go.dev/pkg/encoding/json/)。有关json的示例用法，请参见[jsonrpc包](https://go.dev/pkg/net/rpc/jsonrpc/)的源文件。
