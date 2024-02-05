+++
title = "让 Fiber 更快"
date = 2024-02-05T09:14:15+08:00
weight = 60
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/guide/faster-fiber]({{< ref "/fiber/Guide/MakeFiberFaster" >}})

# ⚡ Make Fiber Faster  让 Fiber 更快

## Custom JSON Encoder/Decoder 自定义 JSON 编码器/解码器 

Since Fiber v2.32.0, we use **encoding/json** as default json library due to stability and producibility. However, the standard library is a bit slow compared to 3rd party libraries. If you're not happy with the performance of **encoding/json**, we recommend you to use these libraries:

​	自 Fiber v2.32.0 起，我们使用 encoding/json 作为默认 json 库，因为它稳定且可生产。但是，与第三方库相比，标准库有点慢。如果您对 encoding/json 的性能不满意，我们建议您使用以下库：

- [goccy/go-json](https://github.com/goccy/go-json)
- [bytedance/sonic](https://github.com/bytedance/sonic)
- [segmentio/encoding](https://github.com/segmentio/encoding)
- [mailru/easyjson](https://github.com/mailru/easyjson)
- [minio/simdjson-go](https://github.com/minio/simdjson-go)
- [wI2L/jettison](https://github.com/wI2L/jettison)

Example
示例

```go
package main

import "github.com/gofiber/fiber/v2"
import "github.com/goccy/go-json"

func main() {
    app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
    })

    # ...
}
```



### References 参考 

- [Set custom JSON encoder for client
  为客户端设置自定义 JSON 编码器]({{< ref "/fiber/API/Client#jsonencoder" >}})
- [Set custom JSON decoder for client
  为客户端设置自定义 JSON 解码器]({{< ref "/fiber/API/Client#jsondecoder" >}})
- [Set custom JSON encoder for application
  为应用程序设置自定义 JSON 编码器]({{< ref "/fiber/API/Fiber#config" >}})
- [Set custom JSON decoder for application
  为应用程序设置自定义 JSON 解码器]({{< ref "/fiber/API/Fiber#config" >}})
