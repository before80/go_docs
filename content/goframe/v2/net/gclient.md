+++
title = "gclient"
date = 2024-03-21T17:52:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gclient

Package gclient provides convenient http client functionalities.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func LoadKeyCrt 

``` go
func LoadKeyCrt(crtFile, keyFile string) (*tls.Config, error)
```

LoadKeyCrt creates and returns a TLS configuration object with given certificate and key files.

### Types 

#### type Client 

``` go
type Client struct {
	http.Client // Underlying HTTP Client.
	// contains filtered or unexported fields
}
```

Client is the HTTP client for HTTP request management.

##### func New 

``` go
func New() *Client
```

New creates and returns a new HTTP client object.

##### Example

``` go
```
##### (*Client) BasicAuth 

``` go
func (c *Client) BasicAuth(user, pass string) *Client
```

BasicAuth is a chaining function, which sets HTTP basic authentication information for next request.

##### (*Client) Clone 

``` go
func (c *Client) Clone() *Client
```

Clone deeply clones current client and returns a new one.

##### Example

``` go
```
##### (*Client) Connect 

``` go
func (c *Client) Connect(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Connect send CONNECT request and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) ConnectBytes 

``` go
func (c *Client) ConnectBytes(ctx context.Context, url string, data ...interface{}) []byte
```

ConnectBytes sends a CONNECT request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) ConnectContent 

``` go
func (c *Client) ConnectContent(ctx context.Context, url string, data ...interface{}) string
```

ConnectContent is a convenience method for sending CONNECT request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) ConnectVar 

``` go
func (c *Client) ConnectVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

ConnectVar sends a CONNECT request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) ContentJson 

``` go
func (c *Client) ContentJson() *Client
```

ContentJson is a chaining function, which sets the HTTP content type as "application/json" for the next request.

Note that it also checks and encodes the parameter to JSON format automatically.

##### Example

``` go
```
##### (*Client) ContentType 

``` go
func (c *Client) ContentType(contentType string) *Client
```

ContentType is a chaining function, which sets HTTP content type for the next request.

##### (*Client) ContentXml 

``` go
func (c *Client) ContentXml() *Client
```

ContentXml is a chaining function, which sets the HTTP content type as "application/xml" for the next request.

Note that it also checks and encodes the parameter to XML format automatically.

##### (*Client) Cookie 

``` go
func (c *Client) Cookie(m map[string]string) *Client
```

Cookie is a chaining function, which sets cookie items with map for next request.

##### Example

``` go
```
##### (*Client) Delete 

``` go
func (c *Client) Delete(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Delete send DELETE request and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) DeleteBytes 

``` go
func (c *Client) DeleteBytes(ctx context.Context, url string, data ...interface{}) []byte
```

DeleteBytes sends a DELETE request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) DeleteContent 

``` go
func (c *Client) DeleteContent(ctx context.Context, url string, data ...interface{}) string
```

DeleteContent is a convenience method for sending DELETE request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) DeleteVar 

``` go
func (c *Client) DeleteVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

DeleteVar sends a DELETE request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) Discovery <-2.5.0

``` go
func (c *Client) Discovery(discovery gsvc.Discovery) *Client
```

Discovery is a chaining function, which sets the discovery for client. You can use `Discovery(nil)` to disable discovery feature for current client.

##### (*Client) DoRequest 

``` go
func (c *Client) DoRequest(ctx context.Context, method, url string, data ...interface{}) (resp *Response, err error)
```

DoRequest sends request with given HTTP method and data and returns the response object. Note that the response object MUST be closed if it'll never be used.

Note that it uses "multipart/form-data" as its Content-Type if it contains file uploading, else it uses "application/x-www-form-urlencoded". It also automatically detects the post content for JSON format, and for that it automatically sets the Content-Type as "application/json".

##### (*Client) DoRequestObj <-2.1.0

``` go
func (c *Client) DoRequestObj(ctx context.Context, req, res interface{}) error
```

DoRequestObj does HTTP request using standard request/response object. The request object `req` is defined like:

``` go
type UseCreateReq struct {
    g.Meta `path:"/user" method:"put"`
    // other fields....
}
```

The response object `res` should be a pointer type. It automatically converts result to given object `res` is success.

Example: var (

```
req = UseCreateReq{}
res *UseCreateRes
```

)

err := DoRequestObj(ctx, req, &res)

##### (*Client) Get 

``` go
func (c *Client) Get(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Get send GET request and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) GetBytes 

``` go
func (c *Client) GetBytes(ctx context.Context, url string, data ...interface{}) []byte
```

GetBytes sends a GET request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) GetContent 

``` go
func (c *Client) GetContent(ctx context.Context, url string, data ...interface{}) string
```

GetContent is a convenience method for sending GET request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) GetVar 

``` go
func (c *Client) GetVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

GetVar sends a GET request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) Head 

``` go
func (c *Client) Head(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Head send HEAD request and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) HeadBytes 

``` go
func (c *Client) HeadBytes(ctx context.Context, url string, data ...interface{}) []byte
```

HeadBytes sends a HEAD request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) HeadContent 

``` go
func (c *Client) HeadContent(ctx context.Context, url string, data ...interface{}) string
```

HeadContent is a convenience method for sending HEAD request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) HeadVar 

``` go
func (c *Client) HeadVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

HeadVar sends a HEAD request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) Header 

``` go
func (c *Client) Header(m map[string]string) *Client
```

Header is a chaining function, which sets custom HTTP headers with map for next request.

##### Example

``` go
```
##### (*Client) HeaderRaw 

``` go
func (c *Client) HeaderRaw(headers string) *Client
```

HeaderRaw is a chaining function, which sets custom HTTP header using raw string for next request.

##### Example

``` go
```
##### (*Client) Next 

``` go
func (c *Client) Next(req *http.Request) (*Response, error)
```

Next calls the next middleware. This should only be call in HandlerFunc.

##### (*Client) NoUrlEncode <-2.5.5

``` go
func (c *Client) NoUrlEncode() *Client
```

NoUrlEncode sets the mark that do not encode the parameters before sending request.

##### (*Client) Options 

``` go
func (c *Client) Options(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Options send OPTIONS request and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) OptionsBytes 

``` go
func (c *Client) OptionsBytes(ctx context.Context, url string, data ...interface{}) []byte
```

OptionsBytes sends an OPTIONS request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) OptionsContent 

``` go
func (c *Client) OptionsContent(ctx context.Context, url string, data ...interface{}) string
```

OptionsContent is a convenience method for sending OPTIONS request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) OptionsVar 

``` go
func (c *Client) OptionsVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

OptionsVar sends an OPTIONS request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) Patch 

``` go
func (c *Client) Patch(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Patch send PATCH request and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) PatchBytes 

``` go
func (c *Client) PatchBytes(ctx context.Context, url string, data ...interface{}) []byte
```

PatchBytes sends a PATCH request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) PatchContent 

``` go
func (c *Client) PatchContent(ctx context.Context, url string, data ...interface{}) string
```

PatchContent is a convenience method for sending PATCH request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) PatchVar 

``` go
func (c *Client) PatchVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

PatchVar sends a PATCH request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) Post 

``` go
func (c *Client) Post(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Post sends request using HTTP method POST and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) PostBytes 

``` go
func (c *Client) PostBytes(ctx context.Context, url string, data ...interface{}) []byte
```

PostBytes sends a POST request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) PostContent 

``` go
func (c *Client) PostContent(ctx context.Context, url string, data ...interface{}) string
```

PostContent is a convenience method for sending POST request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) PostForm 

``` go
func (c *Client) PostForm(ctx context.Context, url string, data map[string]string) (resp *Response, err error)
```

PostForm is different from net/http.PostForm. It's a wrapper of Post method, which sets the Content-Type as "multipart/form-data;". and It will automatically set boundary characters for the request body and Content-Type.

It's Seem like the following case:

Content-Type: multipart/form-data; boundary=----Boundarye4Ghaog6giyQ9ncN

And form data is like: ------Boundarye4Ghaog6giyQ9ncN Content-Disposition: form-data; name="checkType"

none

It's used for sending form data. Note that the response object MUST be closed if it'll never be used.

##### (*Client) PostVar 

``` go
func (c *Client) PostVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

PostVar sends a POST request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) Prefix 

``` go
func (c *Client) Prefix(prefix string) *Client
```

Prefix is a chaining function, which sets the URL prefix for next request of this client. Eg: Prefix("http://127.0.0.1:8199/api/v1") Prefix("http://127.0.0.1:8199/api/v2")

##### Example

``` go
```
##### (*Client) Proxy 

``` go
func (c *Client) Proxy(proxyURL string) *Client
```

Proxy is a chaining function, which sets proxy for next request. Make sure you pass the correct `proxyURL`. The correct pattern is like `[http://USER:PASSWORD@IP:PORT](http://USER:PASSWORD@ip:PORT/)` or `socks5://USER:PASSWORD@IP:PORT`. Only `http` and `socks5` proxies are supported currently.

##### Example

``` go
```
##### (*Client) Put 

``` go
func (c *Client) Put(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Put send PUT request and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) PutBytes 

``` go
func (c *Client) PutBytes(ctx context.Context, url string, data ...interface{}) []byte
```

PutBytes sends a PUT request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) PutContent 

``` go
func (c *Client) PutContent(ctx context.Context, url string, data ...interface{}) string
```

PutContent is a convenience method for sending PUT request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) PutVar 

``` go
func (c *Client) PutVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

PutVar sends a PUT request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) RedirectLimit 

``` go
func (c *Client) RedirectLimit(redirectLimit int) *Client
```

RedirectLimit is a chaining function, which sets the redirect limit the number of jumps for the request.

##### Example

``` go
```
##### (*Client) RequestBytes 

``` go
func (c *Client) RequestBytes(ctx context.Context, method string, url string, data ...interface{}) []byte
```

RequestBytes sends request using given HTTP method and data, retrieves returns the result as bytes. It reads and closes the response object internally automatically.

##### (*Client) RequestContent 

``` go
func (c *Client) RequestContent(ctx context.Context, method string, url string, data ...interface{}) string
```

RequestContent is a convenience method for sending custom http method request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) RequestVar 

``` go
func (c *Client) RequestVar(ctx context.Context, method string, url string, data ...interface{}) *gvar.Var
```

RequestVar sends request using given HTTP method and data, retrieves converts the result to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### (*Client) Retry 

``` go
func (c *Client) Retry(retryCount int, retryInterval time.Duration) *Client
```

Retry is a chaining function, which sets retry count and interval when failure for next request.

##### Example

``` go
```
##### (*Client) SetAgent 

``` go
func (c *Client) SetAgent(agent string) *Client
```

SetAgent sets the User-Agent header for client.

##### (*Client) SetBasicAuth 

``` go
func (c *Client) SetBasicAuth(user, pass string) *Client
```

SetBasicAuth sets HTTP basic authentication information for the client.

##### (*Client) SetBrowserMode 

``` go
func (c *Client) SetBrowserMode(enabled bool) *Client
```

SetBrowserMode enables browser mode of the client. When browser mode is enabled, it automatically saves and sends cookie content from and to server.

##### Example

``` go
```
##### (*Client) SetBuilder <-2.3.3

``` go
func (c *Client) SetBuilder(builder gsel.Builder)
```

SetBuilder sets the load balance builder for client.

##### (*Client) SetContentType 

``` go
func (c *Client) SetContentType(contentType string) *Client
```

SetContentType sets HTTP content type for the client.

##### (*Client) SetCookie 

``` go
func (c *Client) SetCookie(key, value string) *Client
```

SetCookie sets a cookie pair for the client.

##### (*Client) SetCookieMap 

``` go
func (c *Client) SetCookieMap(m map[string]string) *Client
```

SetCookieMap sets cookie items with map.

##### (*Client) SetDiscovery <-2.3.3

``` go
func (c *Client) SetDiscovery(discovery gsvc.Discovery)
```

SetDiscovery sets the load balance builder for client.

##### (*Client) SetHeader 

``` go
func (c *Client) SetHeader(key, value string) *Client
```

SetHeader sets a custom HTTP header pair for the client.

##### Example

``` go
```
##### (*Client) SetHeaderMap 

``` go
func (c *Client) SetHeaderMap(m map[string]string) *Client
```

SetHeaderMap sets custom HTTP headers with map.

##### (*Client) SetHeaderRaw 

``` go
func (c *Client) SetHeaderRaw(headers string) *Client
```

SetHeaderRaw sets custom HTTP header using raw string.

##### (*Client) SetNoUrlEncode <-2.5.5

``` go
func (c *Client) SetNoUrlEncode(noUrlEncode bool) *Client
```

SetNoUrlEncode sets the mark that do not encode the parameters before sending request.

##### (*Client) SetPrefix 

``` go
func (c *Client) SetPrefix(prefix string) *Client
```

SetPrefix sets the request server URL prefix.

##### (*Client) SetProxy 

``` go
func (c *Client) SetProxy(proxyURL string)
```

SetProxy set proxy for the client. This func will do nothing when the parameter `proxyURL` is empty or in wrong pattern. The correct pattern is like `[http://USER:PASSWORD@IP:PORT](http://USER:PASSWORD@ip:PORT/)` or `socks5://USER:PASSWORD@IP:PORT`. Only `http` and `socks5` proxies are supported currently.

##### Example

``` go
```
##### (*Client) SetRedirectLimit 

``` go
func (c *Client) SetRedirectLimit(redirectLimit int) *Client
```

SetRedirectLimit limits the number of jumps.

##### Example

``` go
```
##### (*Client) SetRetry 

``` go
func (c *Client) SetRetry(retryCount int, retryInterval time.Duration) *Client
```

SetRetry sets retry count and interval.

##### (*Client) SetTLSConfig 

``` go
func (c *Client) SetTLSConfig(tlsConfig *tls.Config) error
```

SetTLSConfig sets the TLS configuration of client.

##### Example

``` go
```
##### (*Client) SetTLSKeyCrt 

``` go
func (c *Client) SetTLSKeyCrt(crtFile, keyFile string) error
```

SetTLSKeyCrt sets the certificate and key file for TLS configuration of client.

##### Example

``` go
```
##### (*Client) SetTimeout 

``` go
func (c *Client) SetTimeout(t time.Duration) *Client
```

SetTimeout sets the request timeout for the client.

##### (*Client) Timeout 

``` go
func (c *Client) Timeout(t time.Duration) *Client
```

Timeout is a chaining function, which sets the timeout for next request.

##### (*Client) Trace 

``` go
func (c *Client) Trace(ctx context.Context, url string, data ...interface{}) (*Response, error)
```

Trace send TRACE request and returns the response object. Note that the response object MUST be closed if it'll never be used.

##### Example

``` go
```
##### (*Client) TraceBytes 

``` go
func (c *Client) TraceBytes(ctx context.Context, url string, data ...interface{}) []byte
```

TraceBytes sends a TRACE request, retrieves and returns the result content as bytes.

##### Example

``` go
```
##### (*Client) TraceContent 

``` go
func (c *Client) TraceContent(ctx context.Context, url string, data ...interface{}) string
```

TraceContent is a convenience method for sending TRACE request, which retrieves and returns the result content and automatically closes response object.

##### Example

``` go
```
##### (*Client) TraceVar 

``` go
func (c *Client) TraceVar(ctx context.Context, url string, data ...interface{}) *gvar.Var
```

TraceVar sends a TRACE request, retrieves and converts the result content to *gvar.Var. The client reads and closes the response object internally automatically. The result *gvar.Var can be conveniently converted to any type you want.

##### Example

``` go
```
##### (*Client) Use 

``` go
func (c *Client) Use(handlers ...HandlerFunc) *Client
```

Use adds one or more middleware handlers to client.

#### type HandlerFunc 

``` go
type HandlerFunc = func(c *Client, r *http.Request) (*Response, error)
```

HandlerFunc middleware handler func

#### type Response 

``` go
type Response struct {
	*http.Response // Response is the underlying http.Response object of certain request.
	// contains filtered or unexported fields
}
```

Response is the struct for client request response.

##### (*Response) Close 

``` go
func (r *Response) Close() error
```

Close closes the response when it will never be used.

##### (*Response) GetCookie 

``` go
func (r *Response) GetCookie(key string) string
```

GetCookie retrieves and returns the cookie value of specified `key`.

##### (*Response) GetCookieMap 

``` go
func (r *Response) GetCookieMap() map[string]string
```

GetCookieMap retrieves and returns a copy of current cookie values map.

##### (*Response) Raw 

``` go
func (r *Response) Raw() string
```

Raw returns the raw text of the request and the response.

##### (*Response) RawDump 

``` go
func (r *Response) RawDump()
```

RawDump outputs the raw text of the request and the response to stdout.

##### (*Response) RawRequest 

``` go
func (r *Response) RawRequest() string
```

RawRequest returns the raw content of the request.

##### (*Response) RawResponse 

``` go
func (r *Response) RawResponse() string
```

RawResponse returns the raw content of the response.

##### (*Response) ReadAll 

``` go
func (r *Response) ReadAll() []byte
```

ReadAll retrieves and returns the response content as []byte.

##### (*Response) ReadAllString 

``` go
func (r *Response) ReadAllString() string
```

ReadAllString retrieves and returns the response content as string.

##### (*Response) SetBodyContent <-2.1.0

``` go
func (r *Response) SetBodyContent(content []byte)
```

SetBodyContent overwrites response content with custom one.

#### type WebSocketClient 

``` go
type WebSocketClient struct {
	*websocket.Dialer
}
```

WebSocketClient wraps the underlying websocket client connection and provides convenient functions.

##### func NewWebSocket 

``` go
func NewWebSocket() *WebSocketClient
```

NewWebSocket creates and returns a new WebSocketClient object.