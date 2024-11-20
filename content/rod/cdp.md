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

You can treat it as a minimal example of how to use the DevTools Protocol, no complex abstraction.

It's thread-safe, and context first.

For basic usage, check this [file](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/example_test.go).

For more info, check the unit tests.

# Overview 

Package cdp for application layer communication with browser.

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


# 常量

This section is empty.

# 变量 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L44)

``` go
var ErrCtxDestroyed = &Error{
	Code:    -32000,
	Message: "Execution context was destroyed.",
}
```

ErrCtxDestroyed type.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L26)

``` go
var ErrCtxNotFound = &Error{
	Code:    -32000,
	Message: "Cannot find context with specified id",
}
```

ErrCtxNotFound type.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L56)

``` go
var ErrNodeNotFoundAtPos = &Error{
	Code:    -32000,
	Message: "No node found at given location",
}
```

ErrNodeNotFoundAtPos type.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L62)

``` go
var ErrNotAttachedToActivePage = &Error{
	Code:    -32000,
	Message: "Not attached to an active page",
}
```

ErrNotAttachedToActivePage type.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L50)

``` go
var ErrObjNotFound = &Error{
	Code:    -32000,
	Message: "Could not find object with given id",
}
```

ErrObjNotFound type.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L38)

``` go
var ErrSearchSessionNotFound = &Error{
	Code:    -32000,
	Message: "No search session with given id found",
}
```

ErrSearchSessionNotFound type.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/cdp/error.go#L32)

``` go
var ErrSessionNotFound = &Error{
	Code:    -32001,
	Message: "Session with given id not found.",
}
```

ErrSessionNotFound type.

# 函数 

This section is empty.

# 类型

## type BadHandshakeError <-0.114.8

``` go
type BadHandshakeError struct {
	Status string
	Body   string
}
```

BadHandshakeError type.

### (*BadHandshakeError) Error <-0.114.8

``` go
func (e *BadHandshakeError) Error() string
```

## type Client 

``` go
type Client struct {
	// contains filtered or unexported fields
}
```

Client is a devtools protocol connection instance.

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
	// https://chromedevtools.github.io/devtools-protocol/tot/Page#method-navigate
	// This will create a new tab and navigate to the test.com
	res, err := client.Call(ctx, "", "Target.createTarget", map[string]string{
		"url": "http://test.com",
	})
	utils.E(err)

	fmt.Println(len(gson.New(res).Get("targetId").Str()))

	// close browser by using the proto lib to encode json
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

### func New 

``` go
func New() *Client
```

New creates a cdp connection, all messages from Client.Event must be received or they will block the client.

### func StartWithURL <-0.106.0

``` go
func StartWithURL(ctx context.Context, u string, h http.Header) (*Client, error)
```

StartWithURL helper to connect to the u with the default websocket lib.

### (*Client) Call 

``` go
func (cdp *Client) Call(ctx context.Context, sessionID, method string, params interface{}) ([]byte, error)
```

Call a method and wait for its response.

### (*Client) Event 

``` go
func (cdp *Client) Event() <-chan *Event
```

Event returns a channel that will emit browser devtools protocol events. Must be consumed or will block producer.

### (*Client) Logger <-0.70.0

``` go
func (cdp *Client) Logger(l utils.Logger) *Client
```

Logger sets the logger to log all the requests, responses, and events transferred between Rod and the browser. The default format for each type is in file format.go.

### (*Client) Start <-0.106.0

``` go
func (cdp *Client) Start(ws WebSocketable) *Client
```

Start to browser.

## type Dialer <-0.75.0

``` go
type Dialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}
```

Dialer interface for WebSocket connection.

## type Error 

``` go
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
```

Error of the Response.

### (*Error) Error 

``` go
func (e *Error) Error() string
```

Error stdlib interface.

### (Error) Is <-0.74.0

``` go
func (e Error) Is(target error) bool
```

Is stdlib interface.

## type Event 

``` go
type Event struct {
	SessionID string          `json:"sessionId,omitempty"`
	Method    string          `json:"method"`
	Params    json.RawMessage `json:"params,omitempty"`
}
```

Event from browser.

### (Event) String <-0.70.0

``` go
func (e Event) String() string
```

## type Request 

``` go
type Request struct {
	ID        int         `json:"id"`
	SessionID string      `json:"sessionId,omitempty"`
	Method    string      `json:"method"`
	Params    interface{} `json:"params,omitempty"`
}
```

Request to send to browser.

### (Request) String <-0.70.0

``` go
func (req Request) String() string
```

## type Response <-0.49.6

``` go
type Response struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *Error          `json:"error,omitempty"`
}
```

Response from browser.

### (Response) String <-0.70.0

``` go
func (res Response) String() string
```

## type WebSocket <-0.75.0

``` go
type WebSocket struct {
	// Dialer is usually used for proxy
	Dialer Dialer
	// contains filtered or unexported fields
}
```

WebSocket client for chromium. It only implements a subset of WebSocket protocol. Both the Read and Write are thread-safe. Limitation: https://bugs.chromium.org/p/chromium/issues/detail?id=1069431 Ref: https://tools.ietf.org/html/rfc6455

### (*WebSocket) Close <-0.106.0

``` go
func (ws *WebSocket) Close() error
```

Close the underlying connection.

### (*WebSocket) Connect <-0.75.0

``` go
func (ws *WebSocket) Connect(ctx context.Context, wsURL string, header http.Header) error
```

Connect to browser.

### (*WebSocket) Read <-0.75.0

``` go
func (ws *WebSocket) Read() ([]byte, error)
```

Read a message from browser.

### (*WebSocket) Send <-0.75.0

``` go
func (ws *WebSocket) Send(msg []byte) error
```

Send a message to browser. Because we use zero-copy design, it will modify the content of the msg. It won't allocate new memory.

## type WebSocketable <-0.78.0

``` go
type WebSocketable interface {
	// Send text message only
	Send(data []byte) error
	// Read returns text message only
	Read() ([]byte, error)
}
```

WebSocketable enables you to choose the websocket lib you want to use. Such as you can easily wrap gorilla/websocket and use it as the transport layer.

### func MustConnectWS <-0.106.0

``` go
func MustConnectWS(wsURL string) WebSocketable
```

MustConnectWS helper to make a websocket connection.
