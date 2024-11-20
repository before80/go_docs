+++
title = "cdp"
date = 2024-11-20T18:02:07+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/go-rod/rod/lib/cdp](https://pkg.go.dev/github.com/go-rod/rod/lib/cdp)
>
> 收录该文档时间：`2024-11-20T18:02:07+08:00`
>
> [Version: v0.116.2](https://pkg.go.dev/github.com/go-rod/rod/lib/cdp?tab=versions)

This client is directly based on this [doc](https://chromedevtools.github.io/devtools-protocol/).

​	该客户端直接基于此 [文档](https://chromedevtools.github.io/devtools-protocol/)。

You can treat it as a minimal example of how to use the DevTools Protocol, no complex abstraction.

​	您可以将其视为使用 DevTools 协议的最小示例，无复杂抽象。

It's thread-safe, and context first.

​	它是线程安全的，并以上下文为先。

For basic usage, check this [file](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/example_test.go).

​	有关基本用法，请查看此 [文件](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/example_test.go)。

> 个人添加
>
> [cdp_example_test.go](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/example_test.go) :
>
> ```go
> package cdp_test
> 
> import (
> 	"context"
> 	"fmt"
> 
> 	"github.com/go-rod/rod/lib/cdp"
> 	"github.com/go-rod/rod/lib/launcher"
> 	"github.com/go-rod/rod/lib/proto"
> 	"github.com/go-rod/rod/lib/utils"
> 	"github.com/ysmood/gson"
> )
> 
> func ExampleClient() {
> 	ctx := context.Background()
> 
> 	// launch a browser
> 	url := launcher.New().MustLaunch()
> 
> 	// create a controller
> 	client := cdp.New().Start(cdp.MustConnectWS(url))
> 
> 	go func() {
> 		for range client.Event() {
> 			// you must consume the events
> 			utils.Noop()
> 		}
> 	}()
> 
> 	// Such as call this endpoint on the api doc:
> 	// https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigate
> 	// This will create a new tab and navigate to the test.com
> 	res, err := client.Call(ctx, "", "Target.createTarget", map[string]string{
> 		"url": "http://test.com",
> 	})
> 	utils.E(err)
> 
> 	fmt.Println(len(gson.New(res).Get("targetId").Str()))
> 
> 	// close browser by using the proto lib to encode json
> 	_ = proto.BrowserClose{}.Call(client)
> 
> 	// Output: 32
> }
> 
> func Example_customize_cdp_log() {
> 	ws := cdp.MustConnectWS(launcher.New().MustLaunch())
> 
> 	cdp.New().
> 		Logger(utils.Log(func(args ...interface{}) {
> 			switch v := args[0].(type) {
> 			case *cdp.Request:
> 				fmt.Printf("id: %d", v.ID)
> 			}
> 		})).
> 		Start(ws)
> }
> ```
>
> 

For more info, check the unit tests.

​	有关更多信息，请查看单元测试。

Package cdp for application layer communication with browser.

​	Package cdp 用于与浏览器的应用层通信。

## Example (Customize_cdp_log)

``` go
package main

import (
	"fmt"

	"github.com/go-rod/rod/lib/cdp"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
)

func main() {
	ws := cdp.MustConnectWS(launcher.New().MustLaunch())

	cdp.New().
		Logger(utils.Log(func(args ...interface{}) {
			switch v := args[0].(type) {
			case *cdp.Request:
				fmt.Printf("id: %d", v.ID)
			}
		})).
		Start(ws)
}
Output:
```


## 常量

This section is empty.

## 变量 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L44)

``` go
var ErrCtxDestroyed = &Error{
	Code:    -32000,
	Message: "Execution context was destroyed.",
}
```

ErrCtxDestroyed type.

​	ErrCtxDestroyed 类型。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L26)

``` go
var ErrCtxNotFound = &Error{
	Code:    -32000,
	Message: "Cannot find context with specified id",
}
```

ErrCtxNotFound type.

​	ErrCtxNotFound 类型。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L56)

``` go
var ErrNodeNotFoundAtPos = &Error{
	Code:    -32000,
	Message: "No node found at given location",
}
```

ErrNodeNotFoundAtPos type.

​	ErrNodeNotFoundAtPos 类型。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L62)

``` go
var ErrNotAttachedToActivePage = &Error{
	Code:    -32000,
	Message: "Not attached to an active page",
}
```

ErrNotAttachedToActivePage type.

​	ErrNotAttachedToActivePage 类型。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L50)

``` go
var ErrObjNotFound = &Error{
	Code:    -32000,
	Message: "Could not find object with given id",
}
```

ErrObjNotFound type.

​	ErrObjNotFound 类型。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L38)

``` go
var ErrSearchSessionNotFound = &Error{
	Code:    -32000,
	Message: "No search session with given id found",
}
```

ErrSearchSessionNotFound type.

​	ErrSearchSessionNotFound 类型。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L32)

``` go
var ErrSessionNotFound = &Error{
	Code:    -32001,
	Message: "Session with given id not found.",
}
```

ErrSessionNotFound type.

​	ErrSessionNotFound 类型。

## 函数 

This section is empty.

## 类型

### type BadHandshakeError <-0.114.8

``` go
type BadHandshakeError struct {
	Status string
	Body   string
}
```

BadHandshakeError type.

​	BadHandshakeError 类型。

#### (*BadHandshakeError) Error <-0.114.8

``` go
func (e *BadHandshakeError) Error() string
```

### type Client 

``` go
type Client struct {
	// contains filtered or unexported fields
}
```

Client is a devtools protocol connection instance.

​	Client 是 DevTools 协议连接实例。

### Example

``` go
package main

import (
	"context"
	"fmt"

	"github.com/go-rod/rod/lib/cdp"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

func main() {
	ctx := context.Background()

	// launch a browser
    // 启动浏览器
	url := launcher.New().MustLaunch()

	// create a controller
	client := cdp.New().Start(cdp.MustConnectWS(url))

	go func() {
		for range client.Event() {
			// you must consume the events
			utils.Noop()
		}
	}()

	// Such as call this endpoint on the api doc:
    // 例如调用 API 文档中的此端点：
	// https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigate
	// This will create a new tab and navigate to the test.com
    // 这将创建一个新标签页并导航到 test.com
	res, err := client.Call(ctx, "", "Target.createTarget", map[string]string{
		"url": "http://test.com",
	})
	utils.E(err)

	fmt.Println(len(gson.New(res).Get("targetId").Str()))

	// close browser by using the proto lib to encode json
    // 使用 proto 库编码 JSON 关闭浏览器
	_ = proto.BrowserClose{}.Call(client)

}
Output:

32
```
### func MustStartWithURL <-0.106.0

``` go
func MustStartWithURL(ctx context.Context, u string, h http.Header) *Client
```

MustStartWithURL helper for ConnectURL.

​	MustStartWithURL 是 ConnectURL 的辅助函数。

### func New 

``` go
func New() *Client
```

New creates a cdp connection, all messages from Client.Event must be received or they will block the client.

​	New 创建一个 cdp 连接，必须接收 Client.Event 中的所有消息，否则将阻塞客户端。

### func StartWithURL <-0.106.0

``` go
func StartWithURL(ctx context.Context, u string, h http.Header) (*Client, error)
```

StartWithURL helper to connect to the u with the default websocket lib.

​	StartWithURL 是连接到 u 的辅助函数，使用默认的 WebSocket 库。

#### (*Client) Call 

``` go
func (cdp *Client) Call(ctx context.Context, sessionID, method string, params interface{}) ([]byte, error)
```

Call a method and wait for its response.

​	调用方法并等待响应。

#### (*Client) Event 

``` go
func (cdp *Client) Event() <-chan *Event
```

Event returns a channel that will emit browser devtools protocol events. Must be consumed or will block producer.

​	Event 返回一个通道，将发出浏览器 DevTools 协议事件。必须消费，否则会阻塞生产者。

#### (*Client) Logger <-0.70.0

``` go
func (cdp *Client) Logger(l utils.Logger) *Client
```

Logger sets the logger to log all the requests, responses, and events transferred between Rod and the browser. The default format for each type is in file format.go.

​	Logger 设置日志记录器，用于记录在 Rod 和浏览器之间传输的所有请求、响应和事件。每种类型的默认格式位于文件 format.go 中。

#### (*Client) Start <-0.106.0

``` go
func (cdp *Client) Start(ws WebSocketable) *Client
```

Start to browser.

​	启动到浏览器的连接。

### type Dialer <-0.75.0

``` go
type Dialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}
```

Dialer interface for WebSocket connection.

​	Dialer 是用于 WebSocket 连接的接口。

### type Error 

``` go
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
```

Error of the Response.

​	Response 的错误类型。

#### (*Error) Error 

``` go
func (e *Error) Error() string
```

Error stdlib interface.

​	Error 实现标准库接口。

#### (Error) Is <-0.74.0

``` go
func (e Error) Is(target error) bool
```

Is stdlib interface.

​	Is 实现标准库接口。

### type Event 

``` go
type Event struct {
	SessionID string          `json:"sessionId,omitempty"`
	Method    string          `json:"method"`
	Params    json.RawMessage `json:"params,omitempty"`
}
```

Event from browser.

​	浏览器事件。(Event) String <-0.70.0

``` go
func (e Event) String() string
```

### type Request 

``` go
type Request struct {
	ID        int         `json:"id"`
	SessionID string      `json:"sessionId,omitempty"`
	Method    string      `json:"method"`
	Params    interface{} `json:"params,omitempty"`
}
```

Request to send to browser.

​	发送到浏览器的请求。

#### (Request) String <-0.70.0

``` go
func (req Request) String() string
```

### type Response <-0.49.6

``` go
type Response struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}
```

Response from browser.

​	来自浏览器的响应。

#### (Response) String <-0.70.0

``` go
func (res Response) String() string
```

### type WebSocket <-0.75.0

``` go
type WebSocket struct {
	// Dialer is usually used for proxy
    // Dialer 通常用于代理
	Dialer Dialer
	// contains filtered or unexported fields
    // 包含过滤的或未导出的字段
}
```

WebSocket client for chromium. It only implements a subset of WebSocket protocol. Both the Read and Write are thread-safe. Limitation: https://bugs.chromium.org/p/chromium/issues/detail?id=1069431 Ref: https://tools.ietf.org/html/rfc6455

​	用于 Chromium 的 WebSocket 客户端。仅实现 WebSocket 协议的一个子集。Read 和 Write 均为线程安全。限制：https://bugs.chromium.org/p/chromium/issues/detail?id=1069431 参考：https://tools.ietf.org/html/rfc6455

#### (*WebSocket) Close <-0.106.0

``` go
func (ws *WebSocket) Close() error
```

Close the underlying connection.

​	关闭底层连接。

#### (*WebSocket) Connect <-0.75.0

``` go
func (ws *WebSocket) Connect(ctx context.Context, wsURL string, header http.Header) error
```

Connect to browser.

​	连接到浏览器。

#### (*WebSocket) Read <-0.75.0

``` go
func (ws *WebSocket) Read() ([]byte, error)
```

Read a message from browser.

​	从浏览器读取消息。

#### (*WebSocket) Send <-0.75.0

``` go
func (ws *WebSocket) Send(msg []byte) error
```

Send a message to browser. Because we use zero-copy design, it will modify the content of the msg. It won't allocate new memory.

​	发送消息到浏览器。由于我们使用零拷贝设计，它会修改 msg 的内容，不会分配新内存。

### type WebSocketable <-0.78.0

``` go
type WebSocketable interface {
	// Send text message only
    // 仅发送文本消息
	Send(data []byte) error
	// Read returns text message only
    // Read 返回仅文本消息
	Read() ([]byte, error)
}
```

WebSocketable enables you to choose the websocket lib you want to use. Such as you can easily wrap gorilla/websocket and use it as the transport layer.

​	WebSocketable 允许您选择所需的 WebSocket 库。例如，您可以轻松包装 gorilla/websocket 并将其用作传输层。

### func MustConnectWS <-0.106.0

``` go
func MustConnectWS(wsURL string) WebSocketable
```

MustConnectWS helper to make a websocket connection.

​	MustConnectWS 是建立 WebSocket 连接的辅助函数。
