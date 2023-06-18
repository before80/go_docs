+++
title = "JSON-RPC：接口的故事"
weight = 14
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# JSON-RPC: a tale of interfaces - JSON-RPC：接口的故事

https://go.dev/blog/json-rpc

Andrew Gerrand
27 April 2010

2010年4月27日	

​	这里我们提供一个例子，展示了Go语言中的[接口](https://go.dev/doc/effective_go.html#interfaces_and_types)如何使得重构现有代码变得更加灵活和可扩展。原来，标准库的[RPC包](https://go.dev/pkg/net/rpc/)使用了一个名为[gob](https://go.dev/pkg/encoding/gob/)的自定义数据格式。在某个应用程序中，我们想使用[JSON](https://go.dev/pkg/encoding/json/)作为备用的数据格式。

​	我们首先定义了一对接口来描述现有数据格式的功能，一个用于客户端，一个用于服务器端（如下所示）。

```go
type ServerCodec interface {
 ReadRequestHeader(*Request) error
 ReadRequestBody(interface{}) error
 WriteResponse(*Response, interface{}) error
 Close() error
}
```

​	在服务器端，我们随后更改了两个内部函数的签名，以接受`ServerCodec`接口而不是我们现有的`gob.Encoder`。下面是其中之一：

```go
func sendResponse(sending *sync.Mutex, req *Request,
 reply interface{}, enc *gob.Encoder, errmsg string)
```

变成

```go
func sendResponse(sending *sync.Mutex, req *Request,
  reply interface{}, enc ServerCodec, errmsg string)
```

​	然后，我们编写了一个简单的`gobServerCodec`包装器来复制原始功能。从那里开始，就很容易构建一个`jsonServerCodec`。

​	在对客户端进行类似的更改之后，这就是我们需要在RPC包上完成的全部工作了。整个过程大约需要20分钟！在整理和测试新代码之后，提交了[最终的更改集](https://github.com/golang/go/commit/dcff89057bc0e0d7cb14cf414f2df6f5fb1a41ec)。

​	在诸如Java或C++之类的继承式语言中，明显的路径是将RPC类概括化，并创建JsonRPC和GobRPC子类。然而，如果您想对该层次结构之外的另一个方向进行进一步的概括（例如，如果您要实现一种备用RPC标准），这种方法会变得棘手。在我们的Go包中，我们采取了一种概念上更简单且需要编写或更改的代码更少的方法。

​	对于任何代码库而言，重要的品质之一是可维护性。随着需求的变化，必须轻松而干净地调整代码，否则它将变得难以处理。我们相信Go语言的轻量级、组合导向的类型系统提供了一种可以扩展的代码结构方法。
