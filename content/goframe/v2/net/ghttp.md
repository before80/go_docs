+++
title = "ghttp"
date = 2024-03-21T17:52:53+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/ghttp](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/ghttp)

Package ghttp provides powerful http server and simple client implements.

​	软件包 ghttp 提供了强大的 http 服务器和简单的客户端实现。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/ghttp/ghttp.go#L133)

```go
const (
	HeaderXUrlPath                     = "x-url-path"         // Used for custom route handler, which does not change URL.Path.
	HookBeforeServe       HookName     = "HOOK_BEFORE_SERVE"  // Hook handler before route handler/file serving.
	HookAfterServe        HookName     = "HOOK_AFTER_SERVE"   // Hook handler after route handler/file serving.
	HookBeforeOutput      HookName     = "HOOK_BEFORE_OUTPUT" // Hook handler before response output.
	HookAfterOutput       HookName     = "HOOK_AFTER_OUTPUT"  // Hook handler after response output.
	ServerStatusStopped   ServerStatus = 0
	ServerStatusRunning   ServerStatus = 1
	DefaultServerName                  = "default"
	DefaultDomainName                  = "default"
	HandlerTypeHandler    HandlerType  = "handler"
	HandlerTypeObject     HandlerType  = "object"
	HandlerTypeMiddleware HandlerType  = "middleware"
	HandlerTypeHook       HandlerType  = "hook"
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/ghttp/ghttp_server_config.go#L39)

```go
const (
	UriTypeDefault  = iota // Method names to the URI converting type, which converts name to its lower case and joins the words using char '-'.
	UriTypeFullName        // Method names to the URI converting type, which does not convert to the method name.
	UriTypeAllLower        // Method names to the URI converting type, which converts name to its lower case.
	UriTypeCamel           // Method names to the URI converting type, which converts name to its camel case.
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/ghttp/ghttp_server_websocket.go#L17)

```go
const (
	// WsMsgText TextMessage denotes a text data message.
	// The text message payload is interpreted as UTF-8 encoded text data.
	WsMsgText = websocket.TextMessage

	// WsMsgBinary BinaryMessage denotes a binary data message.
	WsMsgBinary = websocket.BinaryMessage

	// WsMsgClose CloseMessage denotes a close control message.
	// The optional message payload contains a numeric code and text.
	// Use the FormatCloseMessage function to format a close message payload.
	WsMsgClose = websocket.CloseMessage

	// WsMsgPing PingMessage denotes a ping control message.
	// The optional message payload is UTF-8 encoded text.
	WsMsgPing = websocket.PingMessage

	// WsMsgPong PongMessage denotes a pong control message.
	// The optional message payload is UTF-8 encoded text.
	WsMsgPong = websocket.PongMessage
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/ghttp/ghttp.go#L128)

```go
const (
	// FreePortAddress marks the server listens using random free port.
	FreePortAddress = ":0"
)
```

## 变量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/ghttp/ghttp.go#L208)

```go
var (
	ErrNeedJsonBody = gerror.NewWithOption(gerror.Option{
		Text: "the request body content should be JSON format",
		Code: gcode.CodeInvalidRequest,
	})
)
```

## 函数

#### func BuildParams

```go
func BuildParams(params interface{}, noUrlEncode ...bool) (encodedParamStr string)
```

BuildParams builds the request string for the http client. The `params` can be type of: string/[]byte/map/struct/*struct.

​	BuildParams 为 http 客户端生成请求字符串。类型 `params` 可以是：string/[]byte/map/struct/*struct。

The optional parameter `noUrlEncode` specifies whether to ignore the url encoding for the data.

​	可选参数 `noUrlEncode` 指定是否忽略数据的 url 编码。

#### func MiddlewareCORS

```go
func MiddlewareCORS(r *Request)
```

MiddlewareCORS is a middleware handler for CORS with default options.

​	中间件CORS是具有默认选项的CORS中间件处理程序。

#### func MiddlewareHandlerResponse

```go
func MiddlewareHandlerResponse(r *Request)
```

MiddlewareHandlerResponse is the default middleware handling handler response object and its error.

#### func MiddlewareJsonBody <-2.1.3

```go
func MiddlewareJsonBody(r *Request)
```

MiddlewareJsonBody validates and returns request body whether JSON format.

​	MiddlewareJsonBody 验证并返回请求正文是否为 JSON 格式。

#### func MiddlewareNeverDoneCtx <-2.6.2

```go
func MiddlewareNeverDoneCtx(r *Request)
```

MiddlewareNeverDoneCtx sets the context never done for current process.

​	MiddlewareNeverDoneCtx 为当前进程设置从未完成的上下文。

#### func RestartAllServer

```go
func RestartAllServer(ctx context.Context, newExeFilePath string) error
```

RestartAllServer restarts all the servers of the process gracefully. The optional parameter `newExeFilePath` specifies the new binary file for creating process.

​	RestartAllServer 优雅地重新启动进程的所有服务器。可选参数 `newExeFilePath` 指定用于创建进程的新二进制文件。

#### func ShutdownAllServer

```go
func ShutdownAllServer(ctx context.Context) error
```

ShutdownAllServer shuts down all servers of current process gracefully.

​	ShutdownAllServer 正常关闭当前进程的所有服务器。

#### func StartPProfServer

```go
func StartPProfServer(port int, pattern ...string)
```

StartPProfServer starts and runs a new server for pprof.

​	StartPProfServer 启动并运行 pprof 的新服务器。

#### func SupportedMethods <-2.4.2

```go
func SupportedMethods() []string
```

SupportedMethods returns all supported HTTP methods.

​	SupportedMethods 返回所有支持的 HTTP 方法。

#### func Wait

```go
func Wait()
```

Wait blocks to wait for all servers done. It’s commonly used in multiple server situation.

​	等待块等待所有服务器完成。它通常用于多服务器情况。

## 类型

### type CORSOptions

```go
type CORSOptions struct {
	AllowDomain      []string // Used for allowing requests from custom domains
	AllowOrigin      string   // Access-Control-Allow-Origin
	AllowCredentials string   // Access-Control-Allow-Credentials
	ExposeHeaders    string   // Access-Control-Expose-Headers
	MaxAge           int      // Access-Control-Max-Age
	AllowMethods     string   // Access-Control-Allow-Methods
	AllowHeaders     string   // Access-Control-Allow-Headers
}
```

CORSOptions is the options for CORS feature. See https://www.w3.org/TR/cors/ .

​	CORSOptions 是 CORS 功能的选项。请参见 https://www.w3.org/TR/cors/ 。

### type Cookie

```go
type Cookie struct {
	// contains filtered or unexported fields
}
```

Cookie for HTTP COOKIE management.

​	用于 HTTP COOKIE 管理的 Cookie。

#### func GetCookie

```go
func GetCookie(r *Request) *Cookie
```

GetCookie creates or retrieves a cookie object with given request. It retrieves and returns an existing cookie object if it already exists with given request. It creates and returns a new cookie object if it does not exist with given request.

​	GetCookie 使用给定的请求创建或检索 cookie 对象。如果给定请求中已存在现有 cookie 对象，则它会检索并返回该对象。如果给定请求中不存在新的 cookie 对象，则它会创建并返回该对象。

#### (*Cookie) Contains

```go
func (c *Cookie) Contains(key string) bool
```

Contains checks if given key exists and not expire in cookie.

​	包含检查给定密钥是否存在并且不会在 cookie 中过期。

#### (*Cookie) Flush

```go
func (c *Cookie) Flush()
```

Flush outputs the cookie items to the client.

​	Flush 将 cookie 项输出到客户端。

#### (*Cookie) Get

```go
func (c *Cookie) Get(key string, def ...string) *gvar.Var
```

Get retrieves and returns the value with specified key. It returns `def` if specified key does not exist and `def` is given.

​	Get 检索并返回具有指定键的值。如果指定的键不存在并且 `def` 已给出，则返回 `def` 。

#### (*Cookie) GetSessionId

```go
func (c *Cookie) GetSessionId() string
```

GetSessionId retrieves and returns the session id from cookie.

​	GetSessionId 从 cookie 中检索并返回会话 ID。

#### (*Cookie) Map

```go
func (c *Cookie) Map() map[string]string
```

Map returns the cookie items as map[string]string.

​	Map 以 map[string]string 的形式返回 cookie 项。

#### (*Cookie) Remove

```go
func (c *Cookie) Remove(key string)
```

Remove deletes specified key and its value from cookie using default domain and path. It actually tells the http client that the cookie is expired, do not send it to server next time.

​	使用默认域和路径从 cookie 中删除删除指定键及其值。它实际上告诉 http 客户端 cookie 已过期，下次不要将其发送到服务器。

#### (*Cookie) RemoveCookie

```go
func (c *Cookie) RemoveCookie(key, domain, path string)
```

RemoveCookie deletes specified key and its value from cookie using given domain and path. It actually tells the http client that the cookie is expired, do not send it to server next time.

​	RemoveCookie 使用给定的域和路径从 cookie 中删除指定的键及其值。它实际上告诉 http 客户端 cookie 已过期，下次不要将其发送到服务器。

#### (*Cookie) Set

```go
func (c *Cookie) Set(key, value string)
```

Set sets cookie item with default domain, path and expiration age.

​	设置具有默认域、路径和到期期限的 cookie 项。

#### (*Cookie) SetCookie

```go
func (c *Cookie) SetCookie(key, value, domain, path string, maxAge time.Duration, options ...CookieOptions)
```

SetCookie sets cookie item with given domain, path and expiration age. The optional parameter `options` specifies extra security configurations, which is usually empty.

​	SetCookie 设置具有给定域、路径和到期期限的 cookie 项。optional 参数 `options` 指定额外的安全配置，该配置通常为空。

#### (*Cookie) SetHttpCookie

```go
func (c *Cookie) SetHttpCookie(httpCookie *http.Cookie)
```

SetHttpCookie sets cookie with *http.Cookie.

​	SetHttpCookie 使用 *http 设置 cookie。饼干。

#### (*Cookie) SetSessionId

```go
func (c *Cookie) SetSessionId(id string)
```

SetSessionId sets session id in the cookie.

​	SetSessionId 在 Cookie 中设置会话 ID。

### type CookieOptions

```go
type CookieOptions struct {
	SameSite http.SameSite // cookie SameSite property
	Secure   bool          // cookie Secure property
	HttpOnly bool          // cookie HttpOnly property
}
```

CookieOptions provides security config for cookies

​	CookieOptions 为 Cookie 提供安全配置

### type DefaultHandlerResponse

```go
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}
```

DefaultHandlerResponse is the default implementation of HandlerResponse.

### type Domain

```go
type Domain struct {
	// contains filtered or unexported fields
}
```

Domain is used for route register for domains.

​	域用于域的路由寄存器。

#### (*Domain) BindHandler

```go
func (d *Domain) BindHandler(pattern string, handler interface{})
```

BindHandler binds the handler for the specified pattern.

​	BindHandler 绑定指定模式的处理程序。

#### (*Domain) BindHookHandler

```go
func (d *Domain) BindHookHandler(pattern string, hook HookName, handler HandlerFunc)
```

BindHookHandler binds the hook handler for the specified pattern.

​	BindHookHandler 绑定指定模式的挂钩处理程序。

#### (*Domain) BindHookHandlerByMap

```go
func (d *Domain) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc)
```

BindHookHandlerByMap binds the hook handler for the specified pattern.

​	BindHookHandlerByMap 绑定指定模式的挂钩处理程序。

#### (*Domain) BindMiddleware

```go
func (d *Domain) BindMiddleware(pattern string, handlers ...HandlerFunc)
```

BindMiddleware binds the middleware for the specified pattern.

​	BindMiddleware 绑定指定模式的中间件。

#### (*Domain) BindMiddlewareDefault

```go
func (d *Domain) BindMiddlewareDefault(handlers ...HandlerFunc)
```

BindMiddlewareDefault binds the default middleware for the specified pattern.

​	BindMiddlewareDefault 绑定指定模式的默认中间件。

#### (*Domain) BindObject

```go
func (d *Domain) BindObject(pattern string, obj interface{}, methods ...string)
```

BindObject binds the object for the specified pattern.

​	BindObject 绑定指定模式的对象。

#### (*Domain) BindObjectMethod

```go
func (d *Domain) BindObjectMethod(pattern string, obj interface{}, method string)
```

BindObjectMethod binds the method for the specified pattern.

​	BindObjectMethod 绑定指定模式的方法。

#### (*Domain) BindObjectRest

```go
func (d *Domain) BindObjectRest(pattern string, obj interface{})
```

BindObjectRest binds the RESTful API for the specified pattern.

​	BindObjectRest 绑定指定模式的 RESTful API。

#### (*Domain) BindStatusHandler

```go
func (d *Domain) BindStatusHandler(status int, handler HandlerFunc)
```

BindStatusHandler binds the status handler for the specified pattern.

​	BindStatusHandler 绑定指定模式的状态处理程序。

#### (*Domain) BindStatusHandlerByMap

```go
func (d *Domain) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc)
```

BindStatusHandlerByMap binds the status handler for the specified pattern.

​	BindStatusHandlerByMap 绑定指定模式的状态处理程序。

#### (*Domain) EnablePProf

```go
func (d *Domain) EnablePProf(pattern ...string)
```

EnablePProf enables PProf feature for server of specified domain.

​	EnablePProf 为指定域的服务器启用 PProf 功能。

#### (*Domain) Group

```go
func (d *Domain) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup
```

Group creates and returns a RouterGroup object, which is bound to a specified domain.

​	Group 创建并返回绑定到指定域的 RouterGroup 对象。

#### (*Domain) Use

```go
func (d *Domain) Use(handlers ...HandlerFunc)
```

Use adds middleware to the domain.

​	使用将中间件添加到域中。

### type HandlerFunc

```go
type HandlerFunc = func(r *Request)
```

HandlerFunc is request handler function.

#### func WrapF

```go
func WrapF(f http.HandlerFunc) HandlerFunc
```

WrapF is a helper function for wrapping http.HandlerFunc and returns a ghttp.HandlerFunc.

​	WrapF 是用于包装 http 的帮助程序函数。HandlerFunc 并返回一个 ghttp。HandlerFunc。

#### func WrapH

```go
func WrapH(h http.Handler) HandlerFunc
```

WrapH is a helper function for wrapping http.Handler and returns a ghttp.HandlerFunc.

​	WrapH 是用于包装 http 的帮助函数。Handler 并返回一个 ghttp。HandlerFunc。

### type HandlerItem <-2.1.0

```go
type HandlerItem struct {
	// Unique handler item id mark.
	// Note that the handler function may be registered multiple times as different handler items,
	// which have different handler item id.
	Id         int
	Name       string          // Handler name, which is automatically retrieved from runtime stack when registered.
	Type       HandlerType     // Handler type: object/handler/middleware/hook.
	Info       handlerFuncInfo // Handler function information.
	InitFunc   HandlerFunc     // Initialization function when request enters the object (only available for object register type).
	ShutFunc   HandlerFunc     // Shutdown function when request leaves out the object (only available for object register type).
	Middleware []HandlerFunc   // Bound middleware array.
	HookName   HookName        // Hook type name, only available for the hook type.
	Router     *Router         // Router object.
	Source     string          // Registering source file `path:line`.
}
```

HandlerItem is the registered handler for route handling, including middleware and hook functions.

​	HandlerItem 是路由处理的注册处理程序，包括中间件和挂钩函数。

#### (HandlerItem) MarshalJSON

```go
func (item HandlerItem) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

### type HandlerItemParsed <-2.2.2

```go
type HandlerItemParsed struct {
	Handler *HandlerItem      // Handler information.
	Values  map[string]string // Router values parsed from URL.Path.
}
```

HandlerItemParsed is the item parsed from URL.Path.

#### (HandlerItemParsed) MarshalJSON

```go
func (item HandlerItemParsed) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

### type HandlerType <-2.5.0

```go
type HandlerType string
```

HandlerType is the route handler enum type.

### type HookName <-2.5.0

```go
type HookName string
```

HookName is the route hook name enum type.

### type Plugin

```go
type Plugin interface {
	Name() string            // Name returns the name of the plugin.
	Author() string          // Author returns the author of the plugin.
	Version() string         // Version returns the version of the plugin, like "v1.0.0".
	Description() string     // Description returns the description of the plugin.
	Install(s *Server) error // Install installs the plugin BEFORE the server starts.
	Remove() error           // Remove removes the plugin when server shuts down.
}
```

Plugin is the interface for server plugin.

​	插件是服务器插件的接口。

### type Request

```go
type Request struct {
	*http.Request
	Server     *Server           // Server.
	Cookie     *Cookie           // Cookie.
	Session    *gsession.Session // Session.
	Response   *Response         // Corresponding Response of this request.
	Router     *Router           // Matched Router for this request. Note that it's not available in HOOK handler.
	EnterTime  int64             // Request starting time in milliseconds.
	LeaveTime  int64             // Request to end time in milliseconds.
	Middleware *middleware       // Middleware manager.
	StaticFile *staticFile       // Static file object for static file serving.
	// contains filtered or unexported fields
}
```

Request is the context object for a request.

​	Request 是请求的上下文对象。

#### func RequestFromCtx

```go
func RequestFromCtx(ctx context.Context) *Request
```

RequestFromCtx retrieves and returns the Request object from context.

​	RequestFromCtx 从上下文中检索并返回 Request 对象。

#### (*Request) Assign

```go
func (r *Request) Assign(key string, value interface{})
```

Assign binds a template variable to current request.

​	Assign 将模板变量绑定到当前请求。

#### (*Request) Assigns

```go
func (r *Request) Assigns(data gview.Params)
```

Assigns binds multiple template variables to current request.

​	Assign 将多个模板变量绑定到当前请求。

#### (*Request) BasicAuth

```go
func (r *Request) BasicAuth(user, pass string, tips ...string) bool
```

BasicAuth enables the http basic authentication feature with a given passport and password and asks client for authentication. It returns true if authentication success, else returns false if failure.

​	BasicAuth 使用给定的 passport 和密码启用 http 基本身份验证功能，并要求客户端进行身份验证。如果身份验证成功，则返回 true，如果失败，则返回 false。

#### (*Request) Context

```go
func (r *Request) Context() context.Context
```

Context is alias for function GetCtx. This function overwrites the http.Request.Context function. See GetCtx.

​	Context 是函数 GetCtx 的别名。此函数将覆盖 http.Request.Context 函数。请参阅 GetCtx。

#### (*Request) Exit

```go
func (r *Request) Exit()
```

Exit exits executing of current HTTP handler.

​	退出执行当前 HTTP 处理程序的出口。

#### (*Request) ExitAll

```go
func (r *Request) ExitAll()
```

ExitAll exits executing of current and following HTTP handlers.

​	ExitAll 退出执行当前和后续 HTTP 处理程序。

#### (*Request) ExitHook

```go
func (r *Request) ExitHook()
```

ExitHook exits executing of current and following HTTP HOOK handlers.

​	ExitHook 退出执行当前和后续 HTTP HOOK 处理程序。

#### (*Request) Get

```go
func (r *Request) Get(key string, def ...interface{}) *gvar.Var
```

Get is alias of GetRequest, which is one of the most commonly used functions for retrieving parameter. See r.GetRequest.

​	Get 是 GetRequest 的别名，它是检索参数最常用的函数之一。请参阅 r.GetRequest。

#### (*Request) GetBody

```go
func (r *Request) GetBody() []byte
```

GetBody retrieves and returns request body content as bytes. It can be called multiple times retrieving the same body content.

​	GetBody 检索请求正文内容并将其返回为字节。它可以多次调用，检索相同的正文内容。

#### (*Request) GetBodyString

```go
func (r *Request) GetBodyString() string
```

GetBodyString retrieves and returns request body content as string. It can be called multiple times retrieving the same body content.

​	GetBodyString 检索请求正文内容并将其作为字符串返回。它可以多次调用，检索相同的正文内容。

#### (*Request) GetClientIp

```go
func (r *Request) GetClientIp() string
```

GetClientIp returns the client ip of this request without port. Note that this ip address might be modified by client header.

​	GetClientIp 返回此请求的客户端 ip，不带端口。请注意，此 IP 地址可能会被客户端标头修改。

#### (*Request) GetCtx

```go
func (r *Request) GetCtx() context.Context
```

GetCtx retrieves and returns the request’s context. Its alias of function Context,to be relevant with function SetCtx.

​	GetCtx 检索并返回请求的上下文。它的函数 Context 别名，与函数 SetCtx 相关。

#### (*Request) GetCtxVar

```go
func (r *Request) GetCtxVar(key interface{}, def ...interface{}) *gvar.Var
```

GetCtxVar retrieves and returns a Var with a given key name. The optional parameter `def` specifies the default value of the Var if given `key` does not exist in the context.

​	GetCtxVar 检索并返回具有给定键名称的 Var。可选参数 `def` 指定 Var 的默认值（如果给定 `key` 的上下文中不存在）。

#### (*Request) GetError

```go
func (r *Request) GetError() error
```

GetError returns the error occurs in the procedure of the request. It returns nil if there’s no error.

​	GetError 返回在请求过程中发生的错误。如果没有错误，它将返回 nil。

#### (*Request) GetForm

```go
func (r *Request) GetForm(key string, def ...interface{}) *gvar.Var
```

GetForm retrieves and returns parameter `key` from form. It returns `def` if `key` does not exist in the form and `def` is given, or else it returns nil.

​	GetForm `key` 从窗体中检索并返回参数。如果 `key` 形式中不存在并且 `def` 被给定，则返回 `def` ，否则返回 nil。

#### (*Request) GetFormMap

```go
func (r *Request) GetFormMap(kvMap ...map[string]interface{}) map[string]interface{}
```

GetFormMap retrieves and returns all form parameters passed from client as map. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

​	GetFormMap 检索并返回从客户端传递的所有表单参数作为映射。该参数 `kvMap` 指定从客户端参数中检索的键，如果客户端未通过，则关联的值为默认值。

#### (*Request) GetFormMapStrStr

```go
func (r *Request) GetFormMapStrStr(kvMap ...map[string]interface{}) map[string]string
```

GetFormMapStrStr retrieves and returns all form parameters passed from client as map[string]string. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

​	GetFormMapStrStr 检索并返回从客户端传递的所有表单参数作为 map[string]string。该参数 `kvMap` 指定从客户端参数中检索的键，如果客户端未通过，则关联的值为默认值。

#### (*Request) GetFormMapStrVar

```go
func (r *Request) GetFormMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var
```

GetFormMapStrVar retrieves and returns all form parameters passed from client as map[string]*gvar.Var. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

​	GetFormMapStrVar 检索并返回从客户端传递的所有表单参数，作为 map[string]*gvar.Var。该参数 `kvMap` 指定从客户端参数中检索的键，如果客户端未通过，则关联的值为默认值。

#### (*Request) GetFormStruct

```go
func (r *Request) GetFormStruct(pointer interface{}, mapping ...map[string]string) error
```

GetFormStruct retrieves all form parameters passed from client and converts them to given struct object. Note that the parameter `pointer` is a pointer to the struct object. The optional parameter `mapping` is used to specify the key to attribute mapping.

​	GetFormStruct 检索从客户端传递的所有表单参数，并将它们转换为给定的结构对象。请注意，该参数 `pointer` 是指向 struct 对象的指针。可选参数 `mapping` 用于指定属性映射的键。

#### (*Request) GetHandlerResponse

```go
func (r *Request) GetHandlerResponse() interface{}
```

GetHandlerResponse retrieves and returns the handler response object and its error.

​	GetHandlerResponse 检索并返回处理程序响应对象及其错误。

#### (*Request) GetHeader

```go
func (r *Request) GetHeader(key string) string
```

GetHeader retrieves and returns the header value with given `key`.

​	GetHeader 检索并返回给定 `key` .

#### (*Request) GetHost

```go
func (r *Request) GetHost() string
```

GetHost returns current request host name, which might be a domain or an IP without port.

​	GetHost 返回当前请求主机名，该主机名可能是域或没有端口的 IP。

#### (*Request) GetJson

```go
func (r *Request) GetJson() (*gjson.Json, error)
```

GetJson parses current request content as JSON format, and returns the JSON object. Note that the request content is read from request BODY, not from any field of FORM.

​	GetJson 将当前请求内容解析为 JSON 格式，并返回 JSON 对象。请注意，请求内容是从请求 BODY 读取的，而不是从 FORM 的任何字段读取的。

#### (*Request) GetMap

```go
func (r *Request) GetMap(def ...map[string]interface{}) map[string]interface{}
```

GetMap is an alias and convenient function for GetRequestMap. See GetRequestMap.

​	GetMap 是 GetRequestMap 的别名和方便的函数。请参阅 GetRequestMap。

#### (*Request) GetMapStrStr

```go
func (r *Request) GetMapStrStr(def ...map[string]interface{}) map[string]string
```

GetMapStrStr is an alias and convenient function for GetRequestMapStrStr. See GetRequestMapStrStr.

​	GetMapStrStr 是 GetRequestMapStrStr 的别名和方便的函数。请参阅 GetRequestMapStrStr。

#### (*Request) GetMultipartFiles

```go
func (r *Request) GetMultipartFiles(name string) []*multipart.FileHeader
```

GetMultipartFiles parses and returns the post files array. Note that the request form should be type of multipart.

​	GetMultipartFiles 分析并返回 post files 数组。请注意，申请表应为多部分类型。

#### (*Request) GetMultipartForm

```go
func (r *Request) GetMultipartForm() *multipart.Form
```

GetMultipartForm parses and returns the form as multipart forms.

​	GetMultipartForm 分析表单并将其作为多部分表单返回。

#### (*Request) GetNeverDoneCtx

```go
func (r *Request) GetNeverDoneCtx() context.Context
```

GetNeverDoneCtx creates and returns a never done context object, which forbids the context manually done, to make the context can be propagated to asynchronous goroutines, which will not be affected by the HTTP request ends.

​	GetNeverDoneCtx 创建并返回一个从未完成的上下文对象，该对象禁止手动完成的上下文，以使上下文可以传播到异步 goroutine，这不会受到 HTTP 请求结束的影响。

This change is considered for common usage habits of developers for context propagation in multiple goroutines creation in one HTTP request.

​	此更改是针对开发人员在一个 HTTP 请求中创建多个 goroutines 时进行上下文传播的常见使用习惯考虑的。

#### (*Request) GetPage

```go
func (r *Request) GetPage(totalSize, pageSize int) *gpage.Page
```

GetPage creates and returns the pagination object for given `totalSize` and `pageSize`. NOTE THAT the page parameter name from clients is constantly defined as gpage.DefaultPageName for simplification and convenience.

​	GetPage 创建并返回给定 `totalSize` 和 `pageSize` 的分页对象。请注意，客户端中的页面参数名称始终定义为 gpage。DefaultPageName 为简化和方便起见。

#### (*Request) GetParam

```go
func (r *Request) GetParam(key string, def ...interface{}) *gvar.Var
```

GetParam returns custom parameter with a given name `key`. It returns `def` if `key` does not exist. It returns nil if `def` is not passed.

​	GetParam 返回具有给定名称 `key` 的自定义参数。如果 `key` 不存在，则返回 `def` 。如果 `def` 未通过，则返回 nil。

#### (*Request) GetQuery

```go
func (r *Request) GetQuery(key string, def ...interface{}) *gvar.Var
```

GetQuery retrieves and return parameter with the given name `key` from query string and request body. It returns `def` if `key` does not exist in the query and `def` is given, or else it returns nil.

​	GetQuery 从查询字符串和请求正文中检索并返回具有给定名称 `key` 的参数。如果 `key` 查询中不存在并且 `def` 给定，则返回 `def` ，否则返回 nil。

Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote in order of priority: query > body.

​	请注意，如果有多个同名的参数，则按优先级顺序检索和覆盖这些参数：查询>正文。

#### (*Request) GetQueryMap

```go
func (r *Request) GetQueryMap(kvMap ...map[string]interface{}) map[string]interface{}
```

GetQueryMap retrieves and returns all parameters passed from the client using HTTP GET method as the map. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

​	GetQueryMap 使用 HTTP GET 方法作为映射检索并返回从客户端传递的所有参数。该参数 `kvMap` 指定从客户端参数中检索的键，如果客户端未通过，则关联的值为默认值。

Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote in order of priority: query > body.

​	请注意，如果有多个同名的参数，则按优先级顺序检索和覆盖这些参数：查询>正文。

#### (*Request) GetQueryMapStrStr

```go
func (r *Request) GetQueryMapStrStr(kvMap ...map[string]interface{}) map[string]string
```

GetQueryMapStrStr retrieves and returns all parameters passed from the client using the HTTP GET method as a

​	GetQueryMapStrStr 使用 HTTP GET 方法检索并返回从客户端传递的所有参数，作为

```
map[string]string. The parameter `kvMap` specifies the keys
```

retrieving from client parameters, the associated values are the default values if the client does not pass.

​	从客户端参数中检索，如果客户端未传递，则关联的值是默认值。

#### (*Request) GetQueryMapStrVar

```go
func (r *Request) GetQueryMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var
```

GetQueryMapStrVar retrieves and returns all parameters passed from the client using the HTTP GET method as map[string]*gvar.Var. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

​	GetQueryMapStrVar 检索并返回使用 HTTP GET 方法从客户端传递的所有参数，如 map[string]*gvar.Var。该参数 `kvMap` 指定从客户端参数中检索的键，如果客户端未通过，则关联的值为默认值。

#### (*Request) GetQueryStruct

```go
func (r *Request) GetQueryStruct(pointer interface{}, mapping ...map[string]string) error
```

GetQueryStruct retrieves all parameters passed from the client using the HTTP GET method and converts them to a given struct object. Note that the parameter `pointer` is a pointer to the struct object. The optional parameter `mapping` is used to specify the key to attribute mapping.

​	GetQueryStruct 使用 HTTP GET 方法检索从客户端传递的所有参数，并将它们转换为给定的结构对象。请注意，该参数 `pointer` 是指向 struct 对象的指针。可选参数 `mapping` 用于指定属性映射的键。

#### (*Request) GetReferer

```go
func (r *Request) GetReferer() string
```

GetReferer returns referer of this request.

#### (*Request) GetRemoteIp

```go
func (r *Request) GetRemoteIp() string
```

GetRemoteIp returns the ip from RemoteAddr.

​	GetRemoteIp 从 RemoteAddr 返回 ip。

#### (*Request) GetRequest

```go
func (r *Request) GetRequest(key string, def ...interface{}) *gvar.Var
```

GetRequest retrieves and returns the parameter named `key` passed from the client and custom params as interface{}, no matter what HTTP method the client is using. The parameter `def` specifies the default value if the `key` does not exist.

​	无论客户端使用哪种 HTTP 方法，GetRequest 都会检索并返回从客户端传递的名为 `key` Map 的参数，并将自定义参数作为 interface{}。如果 不存在 `key` ，则该参数 `def` 指定默认值。

GetRequest is one of the most commonly used functions for retrieving parameters.

​	GetRequest 是用于检索参数的最常用函数之一。

Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote in order of priority: router < query < body < form < custom.

​	请注意，如果有多个同名参数，则按优先级顺序检索和覆盖参数：路由器<查询正文<<<自定义表单。

#### (*Request) GetRequestMap

```go
func (r *Request) GetRequestMap(kvMap ...map[string]interface{}) map[string]interface{}
```

GetRequestMap retrieves and returns all parameters passed from the client and custom params as the map, no matter what HTTP method the client is using. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass the according keys.

​	无论客户端使用哪种 HTTP 方法，GetRequestMap 都会检索并返回从客户端传递的所有参数和自定义参数作为映射。该参数 `kvMap` 指定从客户端参数中检索的密钥，如果客户端未传递相应的密钥，则关联的值为默认值。

GetRequestMap is one of the most commonly used functions for retrieving parameters.

​	GetRequestMap 是用于检索参数的最常用函数之一。

Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote in order of priority: router < query < body < form < custom.

​	请注意，如果有多个同名参数，则按优先级顺序检索和覆盖参数：路由器<查询正文<<<自定义表单。

#### (*Request) GetRequestMapStrStr

```go
func (r *Request) GetRequestMapStrStr(kvMap ...map[string]interface{}) map[string]string
```

GetRequestMapStrStr retrieve and returns all parameters passed from the client and custom params as map[string]string, no matter what HTTP method the client is using. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

​	GetRequestMapStrStr 检索并返回从客户端和自定义参数传递的所有参数作为 map[string]string，无论客户端使用什么 HTTP 方法。该参数 `kvMap` 指定从客户端参数中检索的键，如果客户端未通过，则关联的值为默认值。

#### (*Request) GetRequestMapStrVar

```go
func (r *Request) GetRequestMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var
```

GetRequestMapStrVar retrieve and returns all parameters passed from the client and custom params as map[string]*gvar.Var, no matter what HTTP method the client is using. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

​	GetRequestMapStrVar 检索并返回从客户端和自定义参数传递的所有参数，作为 map[string]*gvar。Var，无论客户端使用什么 HTTP 方法。该参数 `kvMap` 指定从客户端参数中检索的键，如果客户端未通过，则关联的值为默认值。

#### (*Request) GetRequestStruct

```go
func (r *Request) GetRequestStruct(pointer interface{}, mapping ...map[string]string) error
```

GetRequestStruct retrieves all parameters passed from the client and custom params no matter what HTTP method the client is using, and converts them to give the struct object. Note that the parameter `pointer` is a pointer to the struct object. The optional parameter `mapping` is used to specify the key to attribute mapping.

​	无论客户端使用什么 HTTP 方法，GetRequestStruct 都会检索从客户端和自定义参数传递的所有参数，并将它们转换为 struct 对象。请注意，该参数 `pointer` 是指向 struct 对象的指针。可选参数 `mapping` 用于指定属性映射的键。

#### (*Request) GetRouter

```go
func (r *Request) GetRouter(key string, def ...interface{}) *gvar.Var
```

GetRouter retrieves and returns the router value with given key name `key`. It returns `def` if `key` does not exist.

​	GetRouter 检索并返回具有给定键名称 `key` 的路由器值。如果 `key` 不存在，则返回 `def` 。

#### (*Request) GetRouterMap

```go
func (r *Request) GetRouterMap() map[string]string
```

GetRouterMap retrieves and returns a copy of the router map.

​	GetRouterMap 检索并返回路由器映射的副本。

#### (*Request) GetServeHandler

```go
func (r *Request) GetServeHandler() *HandlerItemParsed
```

GetServeHandler retrieves and returns the user defined handler used to serve this request.

​	GetServeHandler 检索并返回用于处理此请求的用户定义处理程序。

#### (*Request) GetSessionId

```go
func (r *Request) GetSessionId() string
```

GetSessionId retrieves and returns session id from cookie or header.

​	GetSessionId 从 Cookie 或标头中检索并返回会话 ID。

#### (*Request) GetStruct

```go
func (r *Request) GetStruct(pointer interface{}, mapping ...map[string]string) error
```

GetStruct is an alias and convenient function for GetRequestStruct. See GetRequestStruct.

​	GetStruct 是 GetRequestStruct 的别名和方便的函数。请参阅 GetRequestStruct。

#### (*Request) GetUploadFile

```go
func (r *Request) GetUploadFile(name string) *UploadFile
```

GetUploadFile retrieves and returns the uploading file with specified form name. This function is used for retrieving single uploading file object, which is uploaded using multipart form content type.

​	GetUploadFile 检索并返回具有指定窗体名称的上载文件。此函数用于检索使用多部分表单内容类型上传的单个上传文件对象。

It returns nil if retrieving failed or no form file with given name posted.

​	如果检索失败或没有发布给定名称的表单文件，则返回 nil。

Note that the `name` is the file field name of the multipart form from client.

​	请注意，这是 `name` 来自客户端的多部分表单的文件字段名称。

#### (*Request) GetUploadFiles

```go
func (r *Request) GetUploadFiles(name string) UploadFiles
```

GetUploadFiles retrieves and returns multiple uploading files with specified form name. This function is used for retrieving multiple uploading file objects, which are uploaded using multipart form content type.

​	GetUploadFiles 检索并返回具有指定窗体名称的多个上传文件。此函数用于检索多个上传文件对象，这些对象是使用多部分表单内容类型上传的。

It returns nil if retrieving failed or no form file with given name posted.

​	如果检索失败或未发布指定名称的表单文件，则返回 nil。

Note that the `name` is the file field name of the multipart form from client.

​	请注意，这是 `name` 来自客户端的多部分表单的文件字段名称。

#### (*Request) GetUrl

```go
func (r *Request) GetUrl() string
```

GetUrl returns current URL of this request.

​	GetUrl 返回此请求的当前 URL。

#### (*Request) GetView

```go
func (r *Request) GetView() *gview.View
```

GetView returns the template view engine object for this request.

​	GetView 返回此请求的模板视图引擎对象。

#### (*Request) IsAjaxRequest

```go
func (r *Request) IsAjaxRequest() bool
```

IsAjaxRequest checks and returns whether current request is an AJAX request.

​	IsAjaxRequest 检查并返回当前请求是否为 AJAX 请求。

#### (*Request) IsExited

```go
func (r *Request) IsExited() bool
```

IsExited checks and returns whether current request is exited.

​	IsExited 检查并返回当前请求是否已退出。

#### (*Request) IsFileRequest

```go
func (r *Request) IsFileRequest() bool
```

IsFileRequest checks and returns whether current request is serving file.

​	IsFileRequest 检查并返回当前请求是否正在提供文件。

#### (*Request) MakeBodyRepeatableRead

```go
func (r *Request) MakeBodyRepeatableRead(repeatableRead bool) []byte
```

MakeBodyRepeatableRead marks the request body could be repeatedly readable or not. It also returns the current content of the request body.

​	MakeBodyRepeatableRead 标记请求正文是否可以重复读取。它还返回请求正文的当前内容。

#### (*Request) Parse

```go
func (r *Request) Parse(pointer interface{}) error
```

Parse is the most commonly used function, which converts request parameters to struct or struct slice. It also automatically validates the struct or every element of the struct slice according to the validation tag of the struct.

​	Parse 是最常用的函数，它将请求参数转换为结构或结构切片。它还会根据结构的验证标记自动验证结构或结构切片的每个元素。

The parameter `pointer` can be type of: *struct/**struct/*[]struct/*[]*struct.

​	参数的类型 `pointer` 可以是：struct/**struct/[]struct/*[]*struct。

It supports single and multiple struct converting: 1. Single struct, post content like: {“id”:1, “name”:“john”} or ?id=1&name=john 2. Multiple struct, post content like: [{“id”:1, “name”:“john”}, {“id”:, “name”:“smith”}]

​	它支持单结构和多结构转换： 1.单结构，发布内容如：{“id”：1， “name”：“john”} 或 ？id=1&name=john 2.多个结构，发布内容如下： [{“id”：1， “name”：“john”}， {“id”：， “name”：“smith”}]

TODO: Improve the performance by reducing duplicated reflect usage on the same variable across packages.

​	TODO：通过减少跨包对同一变量的重复反映使用来提高性能。

#### (*Request) ParseForm

```go
func (r *Request) ParseForm(pointer interface{}) error
```

ParseForm performs like function Parse, but only parses the form parameters or the body content.

​	ParseForm 的执行方式类似于函数 Parse，但仅分析表单参数或正文内容。

#### (*Request) ParseQuery

```go
func (r *Request) ParseQuery(pointer interface{}) error
```

ParseQuery performs like function Parse, but only parses the query parameters.

​	ParseQuery 的执行方式与函数 Parse 类似，但仅分析查询参数。

#### (*Request) ReloadParam

```go
func (r *Request) ReloadParam()
```

ReloadParam is used for modifying request parameter. Sometimes, we want to modify request parameters through middleware, but directly modifying Request.Body is invalid, so it clears the parsed* marks of Request to make the parameters reparsed.

​	ReloadParam 用于修改请求参数。有时，我们想通过中间件修改请求参数，但直接修改 Request.Body 是无效的，所以它会清除 Request 的解析*标记，使参数重新解析。

#### (*Request) SetCtx

```go
func (r *Request) SetCtx(ctx context.Context)
```

SetCtx custom context for current request.

​	当前请求的 SetCtx 自定义上下文。

#### (*Request) SetCtxVar

```go
func (r *Request) SetCtxVar(key interface{}, value interface{})
```

SetCtxVar sets custom parameter to context with key-value pairs.

​	SetCtxVar 使用键值对将自定义参数设置为上下文。

#### (*Request) SetError

```go
func (r *Request) SetError(err error)
```

SetError sets custom error for current request.

#### (*Request) SetForm

```go
func (r *Request) SetForm(key string, value interface{})
```

SetForm sets custom form value with key-value pairs.

#### (*Request) SetParam

```go
func (r *Request) SetParam(key string, value interface{})
```

SetParam sets custom parameter with key-value pairs.

#### (*Request) SetParamMap

```go
func (r *Request) SetParamMap(data map[string]interface{})
```

SetParamMap sets custom parameter with key-value pair maps.

#### (*Request) SetQuery

```go
func (r *Request) SetQuery(key string, value interface{})
```

SetQuery sets custom query value with key-value pairs.

​	SetQuery 使用键值对设置自定义查询值。

#### (*Request) SetView

```go
func (r *Request) SetView(view *gview.View)
```

SetView sets template view engine object for this request.

​	SetView 为此请求设置模板视图引擎对象。

#### (*Request) WebSocket

```go
func (r *Request) WebSocket() (*WebSocket, error)
```

WebSocket upgrades current request as a websocket request. It returns a new WebSocket object if success, or the error if failure. Note that the request should be a websocket request, or it will surely fail upgrading.

​	WebSocket 将当前请求升级为 websocket 请求。如果成功，它将返回一个新的 WebSocket 对象，如果失败，则返回错误。请注意，该请求应该是 websocket 请求，否则升级肯定会失败。

### type Response

```go
type Response struct {
	*ResponseWriter                 // Underlying ResponseWriter.
	Server          *Server         // Parent server.
	Writer          *ResponseWriter // Alias of ResponseWriter.
	Request         *Request        // According request.
}
```

Response is the http response manager. Note that it implements the http.ResponseWriter interface with buffering feature.

​	Response 是 http 响应管理器。请注意，它实现了 http.具有缓冲功能的 ResponseWriter 接口。

#### (*Response) Buffer

```go
func (r *Response) Buffer() []byte
```

Buffer returns the buffered content as []byte.

#### (*Response) BufferLength

```go
func (r *Response) BufferLength() int
```

BufferLength returns the length of the buffered content.

​	BufferLength 返回缓冲内容的长度。

#### (*Response) BufferString

```go
func (r *Response) BufferString() string
```

BufferString returns the buffered content as string.

​	BufferString 以字符串形式返回缓冲内容。

#### (*Response) CORS

```go
func (r *Response) CORS(options CORSOptions)
```

CORS sets custom CORS options. See https://www.w3.org/TR/cors/ .

​	CORS 设置自定义 CORS 选项。请参见 https://www.w3.org/TR/cors/ 。

#### (*Response) CORSAllowedOrigin

```go
func (r *Response) CORSAllowedOrigin(options CORSOptions) bool
```

CORSAllowedOrigin CORSAllowed checks whether the current request origin is allowed cross-domain.

​	CORSAllowedOrigin CORSAllowed检查当前请求源是否允许跨域。

#### (*Response) CORSDefault

```go
func (r *Response) CORSDefault()
```

CORSDefault sets CORS with default CORS options, which allows any cross-domain request.

​	CORSDefault 使用默认的 CORS 选项设置 CORS，该选项允许任何跨域请求。

#### (*Response) ClearBuffer

```go
func (r *Response) ClearBuffer()
```

ClearBuffer clears the response buffer.

​	ClearBuffer 清除响应缓冲区。

#### (*Response) DefaultCORSOptions

```go
func (r *Response) DefaultCORSOptions() CORSOptions
```

DefaultCORSOptions returns the default CORS options, which allows any cross-domain request.

​	DefaultCORSOptions 返回默认的 CORS 选项，该选项允许任何跨域请求。

#### (*Response) Flush

```go
func (r *Response) Flush()
```

Flush outputs the buffer content to the client and clears the buffer.

​	Flush 将缓冲区内容输出到客户端并清除缓冲区。

#### (*Response) ParseTpl

```go
func (r *Response) ParseTpl(tpl string, params ...gview.Params) (string, error)
```

ParseTpl parses given template file `tpl` with given template variables `params` and returns the parsed template content.

​	ParseTpl 使用给定的 `params` 模板变量解析给定的模板文件 `tpl` ，并返回解析的模板内容。

#### (*Response) ParseTplContent

```go
func (r *Response) ParseTplContent(content string, params ...gview.Params) (string, error)
```

ParseTplContent parses given template file `file` with given template parameters `params` and returns the parsed template content.

​	ParseTplContent 使用给定的模板参数 `params` 分析给定的模板文件 `file` ，并返回分析后的模板内容。

#### (*Response) ParseTplDefault

```go
func (r *Response) ParseTplDefault(params ...gview.Params) (string, error)
```

ParseTplDefault parses the default template file with params.

​	ParseTplDefault 使用参数分析默认模板文件。

#### (*Response) RedirectBack

```go
func (r *Response) RedirectBack(code ...int)
```

RedirectBack redirects the client back to referer. The optional parameter `code` specifies the http status code for redirecting, which commonly can be 301 or 302. It’s 302 in default.

​	RedirectBack 将客户端重定向回 referer。可选参数 `code` 指定用于重定向的 http 状态代码，通常可以是 301 或 302。默认为 302。

#### (*Response) RedirectTo

```go
func (r *Response) RedirectTo(location string, code ...int)
```

RedirectTo redirects the client to another location. The optional parameter `code` specifies the http status code for redirecting, which commonly can be 301 or 302. It’s 302 in default.

​	RedirectTo 将客户端重定向到另一个位置。可选参数 `code` 指定用于重定向的 http 状态代码，通常可以是 301 或 302。默认为 302。

#### (*Response) ServeContent

```go
func (r *Response) ServeContent(name string, modTime time.Time, content io.ReadSeeker)
```

ServeContent replies to the request using the content in the provided ReadSeeker. The main benefit of ServeContent over io.Copy is that it handles Range requests properly, sets the MIME type, and handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since, and If-Range requests.

​	ServeContent 使用提供的 ReadSeeker 中的内容回复请求。ServeContent 相对于 io 的主要优势。复制是指它正确处理 Range 请求，设置 MIME 类型，并处理 If-Match、If-Unmodified-Since、If-None-Match、If-Modified-Since 和 If-Range 请求。

See http.ServeContent

​	请参见 http。服务内容

#### (*Response) ServeFile

```go
func (r *Response) ServeFile(path string, allowIndex ...bool)
```

ServeFile serves the file to the response.

​	ServeFile 为响应提供文件。

#### (*Response) ServeFileDownload

```go
func (r *Response) ServeFileDownload(path string, name ...string)
```

ServeFileDownload serves file downloading to the response.

​	ServeFileDownload 将文件下载到响应。

#### (*Response) SetBuffer

```go
func (r *Response) SetBuffer(data []byte)
```

SetBuffer overwrites the buffer with `data`.

​	SetBuffer 使用 `data` 覆盖缓冲区。

#### (*Response) Write

```go
func (r *Response) Write(content ...interface{})
```

Write writes `content` to the response buffer.

#### (*Response) WriteExit

```go
func (r *Response) WriteExit(content ...interface{})
```

WriteExit writes `content` to the response buffer and exits executing of current handler. The “Exit” feature is commonly used to replace usage of return statements in the handler, for convenience.

​	WriteExit 写入 `content` 响应缓冲区并退出当前处理程序的执行。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

#### (*Response) WriteJson

```go
func (r *Response) WriteJson(content interface{})
```

WriteJson writes `content` to the response with JSON format.

​	WriteJson 使用 JSON 格式写 `content` 入响应。

#### (*Response) WriteJsonExit

```go
func (r *Response) WriteJsonExit(content interface{})
```

WriteJsonExit writes `content` to the response with JSON format and exits executing of current handler if success. The “Exit” feature is commonly used to replace usage of return statements in the handler, for convenience.

​	WriteJsonExit 使用 JSON 格式写 `content` 入响应，如果成功，则退出执行当前处理程序。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

#### (*Response) WriteJsonP

```go
func (r *Response) WriteJsonP(content interface{})
```

WriteJsonP writes `content` to the response with JSONP format.

​	WriteJsonP 使用 JSONP 格式写 `content` 入响应。

Note that there should be a “callback” parameter in the request for JSONP format.

​	请注意，JSONP 格式的请求中应该有一个“callback”参数。

#### (*Response) WriteJsonPExit

```go
func (r *Response) WriteJsonPExit(content interface{})
```

WriteJsonPExit writes `content` to the response with JSONP format and exits executing of current handler if success. The “Exit” feature is commonly used to replace usage of return statements in the handler, for convenience.

​	WriteJsonPExit 使用 JSONP 格式写 `content` 入响应，如果成功，则退出当前处理程序的执行。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

Note that there should be a “callback” parameter in the request for JSONP format.

​	请注意，JSONP 格式的请求中应该有一个“callback”参数。

#### (*Response) WriteOver

```go
func (r *Response) WriteOver(content ...interface{})
```

WriteOver overwrites the response buffer with `content`.

​	WriteOver 使用 `content` 覆盖响应缓冲区。

#### (*Response) WriteOverExit

```go
func (r *Response) WriteOverExit(content ...interface{})
```

WriteOverExit overwrites the response buffer with `content` and exits executing of current handler. The “Exit” feature is commonly used to replace usage of return statements in the handler, for convenience.

​	WriteOverExit 使用 `content` 当前处理程序覆盖响应缓冲区并退出执行。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

#### (*Response) WriteStatus

```go
func (r *Response) WriteStatus(status int, content ...interface{})
```

WriteStatus writes HTTP `status` and `content` to the response. Note that it does not set a Content-Type header here.

​	WriteStatus 将 HTTP `status` 和 `content` 写入响应。请注意，它未在此处设置 Content-Type 标头。

#### (*Response) WriteStatusExit

```go
func (r *Response) WriteStatusExit(status int, content ...interface{})
```

WriteStatusExit writes HTTP `status` and `content` to the response and exits executing of current handler if success. The “Exit” feature is commonly used to replace usage of return statements in the handler, for convenience.

​	WriteStatusExit 将 HTTP `status` 和 `content` 写入响应，如果成功，则退出执行当前处理程序。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

#### (*Response) WriteTpl

```go
func (r *Response) WriteTpl(tpl string, params ...gview.Params) error
```

WriteTpl parses and responses given template file. The parameter `params` specifies the template variables for parsing.

​	WriteTpl 解析并响应给定的模板文件。该参数 `params` 指定用于分析的模板变量。

#### (*Response) WriteTplContent

```go
func (r *Response) WriteTplContent(content string, params ...gview.Params) error
```

WriteTplContent parses and responses the template content. The parameter `params` specifies the template variables for parsing.

​	WriteTplContent 分析并响应模板内容。该参数 `params` 指定用于分析的模板变量。

#### (*Response) WriteTplDefault

```go
func (r *Response) WriteTplDefault(params ...gview.Params) error
```

WriteTplDefault parses and responses the default template file. The parameter `params` specifies the template variables for parsing.

​	WriteTplDefault 分析并响应默认模板文件。该参数 `params` 指定用于分析的模板变量。

#### (*Response) WriteXml

```go
func (r *Response) WriteXml(content interface{}, rootTag ...string)
```

WriteXml writes `content` to the response with XML format.

​	WriteXml 使用 XML 格式写 `content` 入响应。

#### (*Response) WriteXmlExit

```go
func (r *Response) WriteXmlExit(content interface{}, rootTag ...string)
```

WriteXmlExit writes `content` to the response with XML format and exits executing of current handler if success. The “Exit” feature is commonly used to replace usage of return statements in the handler, for convenience.

​	WriteXmlExit 使用 XML 格式写 `content` 入响应，如果成功，则退出当前处理程序的执行。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

#### (*Response) Writef

```go
func (r *Response) Writef(format string, params ...interface{})
```

Writef writes the response with fmt.Sprintf.

​	Writef 使用 fmt 写入响应。斯普林特夫。

#### (*Response) WritefExit

```go
func (r *Response) WritefExit(format string, params ...interface{})
```

WritefExit writes the response with fmt.Sprintf and exits executing of current handler. The “Exit” feature is commonly used to replace usage of return statements in the handler, for convenience.

​	WritefExit 使用 fmt 写入响应。Sprintf 并退出执行当前处理程序。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

#### (*Response) Writefln

```go
func (r *Response) Writefln(format string, params ...interface{})
```

Writefln writes the response with fmt.Sprintf and new line.

​	Writefln 使用 fmt 写入响应。Sprintf 和新线。

#### (*Response) WriteflnExit

```go
func (r *Response) WriteflnExit(format string, params ...interface{})
```

WriteflnExit writes the response with fmt.Sprintf and new line and exits executing of current handler. The “Exit” feature is commonly used to replace usage of return statement in the handler, for convenience.

​	WriteflnExit 使用 fmt 写入响应。Sprintf 和 new 行和出口执行当前处理程序。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

#### (*Response) Writeln

```go
func (r *Response) Writeln(content ...interface{})
```

Writeln writes the response with `content` and new line.

​	Writeln 用 `content` 和 换行符写入响应。

#### (*Response) WritelnExit

```go
func (r *Response) WritelnExit(content ...interface{})
```

WritelnExit writes the response with `content` and new line and exits executing of current handler. The “Exit” feature is commonly used to replace usage of return statements in the handler, for convenience.

​	WritelnExit 使用 `content` 和 换行写入响应，并退出执行当前处理程序。为方便起见，“Exit”功能通常用于替换处理程序中 return 语句的用法。

### type ResponseWriter

```go
type ResponseWriter struct {
	Status int // HTTP status.
	// contains filtered or unexported fields
}
```

ResponseWriter is the custom writer for http response.

​	ResponseWriter 是 http 响应的自定义编写器。

#### (*ResponseWriter) Flush

```go
func (w *ResponseWriter) Flush()
```

Flush outputs the buffer to clients and clears the buffer.

​	Flush 将缓冲区输出到客户端并清除缓冲区。

#### (*ResponseWriter) Header

```go
func (w *ResponseWriter) Header() http.Header
```

Header implements the interface function of http.ResponseWriter.Header.

​	Header 实现 http 的接口功能。ResponseWriter.Header。

#### (*ResponseWriter) Hijack

```go
func (w *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error)
```

Hijack implements the interface function of http.Hijacker.Hijack.

​	Hijack 实现了 http 的接口功能。劫机者劫持。

#### (*ResponseWriter) RawWriter

```go
func (w *ResponseWriter) RawWriter() http.ResponseWriter
```

RawWriter returns the underlying ResponseWriter.

#### (*ResponseWriter) Write

```go
func (w *ResponseWriter) Write(data []byte) (int, error)
```

Write implements the interface function of http.ResponseWriter.Write.

​	write 实现了 http 的接口函数。ResponseWriter.Write。

#### (*ResponseWriter) WriteHeader

```go
func (w *ResponseWriter) WriteHeader(status int)
```

WriteHeader implements the interface of http.ResponseWriter.WriteHeader.

​	WriteHeader 实现了 http 的接口。ResponseWriter.WriteHeader。

### type Router

```go
type Router struct {
	Uri      string   // URI.
	Method   string   // HTTP method
	Domain   string   // Bound domain.
	RegRule  string   // Parsed regular expression for route matching.
	RegNames []string // Parsed router parameter names.
	Priority int      // Just for reference.
}
```

Router object.

​	路由器对象。

### type RouterGroup

```go
type RouterGroup struct {
	// contains filtered or unexported fields
}
```

RouterGroup is a group wrapping multiple routes and middleware.

​	RouterGroup 是一个包装多个路由和中间件的组。

#### (*RouterGroup) ALL

```go
func (g *RouterGroup) ALL(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

ALL register an http handler to give the route pattern and all http methods.

​	ALL 注册一个 http 处理程序以提供路由模式和所有 http 方法。

#### (*RouterGroup) ALLMap

```go
func (g *RouterGroup) ALLMap(m map[string]interface{})
```

ALLMap registers http handlers for http methods using map.

​	ALLMap 使用 map 为 http 方法注册 http 处理程序。

#### (*RouterGroup) Bind

```go
func (g *RouterGroup) Bind(handlerOrObject ...interface{}) *RouterGroup
```

Bind does batch route registering feature for a router group.

​	Bind 为路由器组提供批量路由注册功能。

#### (*RouterGroup) CONNECT

```go
func (g *RouterGroup) CONNECT(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

CONNECT registers an http handler to give the route pattern and the http method: CONNECT.

​	CONNECT 注册一个 http 处理程序来提供路由模式和 http 方法：CONNECT。

#### (*RouterGroup) Clone

```go
func (g *RouterGroup) Clone() *RouterGroup
```

Clone returns a new router group which is a clone of the current group.

​	Clone 返回一个新的路由器组，该组是当前组的克隆。

#### (*RouterGroup) DELETE

```go
func (g *RouterGroup) DELETE(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

DELETE registers an http handler to give the route pattern and the http method: DELETE.

​	DELETE 注册一个 http 处理程序以提供路由模式和 http 方法：DELETE。

#### (*RouterGroup) GET

```go
func (g *RouterGroup) GET(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

GET registers an http handler to give the route pattern and the http method: GET.

​	GET 注册一个 http 处理程序来提供路由模式和 http 方法：GET。

#### (*RouterGroup) Group

```go
func (g *RouterGroup) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup
```

Group creates and returns a subgroup of the current router group.

​	组创建并返回当前路由器组的子组。

#### (*RouterGroup) HEAD

```go
func (g *RouterGroup) HEAD(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

HEAD registers an http handler to give the route pattern and the http method: HEAD.

​	HEAD 注册一个 http 处理程序来提供路由模式和 http 方法：HEAD。

#### (*RouterGroup) Hook

```go
func (g *RouterGroup) Hook(pattern string, hook HookName, handler HandlerFunc) *RouterGroup
```

Hook registers a hook to given route pattern.

​	Hook 将钩子注册到给定的路由模式。

#### (*RouterGroup) Map

```go
func (g *RouterGroup) Map(m map[string]interface{})
```

Map registers http handlers for http methods using map.

​	Map 使用 map 为 http 方法注册 http 处理程序。

#### (*RouterGroup) Middleware

```go
func (g *RouterGroup) Middleware(handlers ...HandlerFunc) *RouterGroup
```

Middleware binds one or more middleware to the router group.

​	中间件将一个或多个中间件绑定到路由器组。

#### (*RouterGroup) OPTIONS

```go
func (g *RouterGroup) OPTIONS(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

OPTIONS register an http handler to give the route pattern and the http method: OPTIONS.

​	OPTIONS 注册一个 http 处理程序来提供路由模式和 http 方法：OPTIONS。

#### (*RouterGroup) PATCH

```go
func (g *RouterGroup) PATCH(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

PATCH registers an http handler to give the route pattern and the http method: PATCH.

​	PATCH 注册一个 http 处理程序来提供路由模式和 http 方法：PATCH。

#### (*RouterGroup) POST

```go
func (g *RouterGroup) POST(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

POST registers an http handler to give the route pattern and the http method: POST.

​	POST 注册一个 http 处理程序来提供路由模式和 http 方法：POST。

#### (*RouterGroup) PUT

```go
func (g *RouterGroup) PUT(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

PUT registers an http handler to give the route pattern and the http method: PUT.

​	PUT 注册一个 http 处理程序来提供路由模式和 http 方法：PUT。

#### (*RouterGroup) REST

```go
func (g *RouterGroup) REST(pattern string, object interface{}) *RouterGroup
```

REST registers an http handler to give the route pattern according to REST rule.

​	REST注册一个http处理程序，以根据REST规则提供路由模式。

#### (*RouterGroup) TRACE

```go
func (g *RouterGroup) TRACE(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

TRACE registers an http handler to give the route pattern and the http method: TRACE.

​	TRACE 注册一个 http 处理程序来提供路由模式和 http 方法：TRACE。

### type RouterItem

```go
type RouterItem struct {
	Handler          *HandlerItem // The handler.
	Server           string       // Server name.
	Address          string       // Listening address.
	Domain           string       // Bound domain.
	Type             HandlerType  // Route handler type.
	Middleware       string       // Bound middleware.
	Method           string       // Handler method name.
	Route            string       // Route URI.
	Priority         int          // Just for reference.
	IsServiceHandler bool         // Is service handler.
}
```

RouterItem is just for route dumps.

​	RouterItem 仅用于路由转储。

### type Server

```go
type Server struct {
	// contains filtered or unexported fields
}
```

Server wraps the http.Server and provides more rich features.

​	服务器包装 http.服务器，并提供更丰富的功能。

#### func GetServer

```go
func GetServer(name ...interface{}) *Server
```

GetServer creates and returns a server instance using given name and default configurations. Note that the parameter `name` should be unique for different servers. It returns an existing server instance if given `name` is already existing in the server mapping.

​	GetServer 使用给定名称和默认配置创建并返回服务器实例。请注意，对于不同的服务器，该参数 `name` 应该是唯一的。如果给定 `name` 的服务器映射中已存在现有服务器实例，则返回现有服务器实例。

#### (*Server) AddSearchPath

```go
func (s *Server) AddSearchPath(path string)
```

AddSearchPath add searching directory path for static file service.

​	AddSearchPath 为静态文件服务添加搜索目录路径。

#### (*Server) AddStaticPath

```go
func (s *Server) AddStaticPath(prefix string, path string)
```

AddStaticPath sets the uri to static directory path mapping for static file service.

​	AddStaticPath 将 uri 设置为静态文件服务的静态目录路径映射。

#### (*Server) BindHandler

```go
func (s *Server) BindHandler(pattern string, handler interface{})
```

BindHandler registers a handler function to server with a given pattern.

​	BindHandler 使用给定模式将处理程序函数注册到服务器。

Note that the parameter `handler` can be type of: 1. func(*ghttp.Request) 2. func(context.Context, BizRequest)(BizResponse, error)

​	请注意，该参数 `handler` 的类型可以是：1. func（*ghttp.请求） 2.func（上下文。Context， BizRequest （BizResponse， error）

#### (*Server) BindHookHandler

```go
func (s *Server) BindHookHandler(pattern string, hook HookName, handler HandlerFunc)
```

BindHookHandler registers handler for specified hook.

​	BindHookHandler 注册指定挂钩的处理程序。

#### (*Server) BindHookHandlerByMap

```go
func (s *Server) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc)
```

BindHookHandlerByMap registers handler for specified hook.

​	BindHookHandlerByMap 注册指定挂钩的处理程序。

#### (*Server) BindMiddleware

```go
func (s *Server) BindMiddleware(pattern string, handlers ...HandlerFunc)
```

BindMiddleware registers one or more global middleware to the server. Global middleware can be used standalone without service handler, which intercepts all dynamic requests before or after service handler. The parameter `pattern` specifies what route pattern the middleware intercepts, which is usually a “fuzzy” pattern like “/:name”, “/*any” or “/{field}”.

​	BindMiddleware 将一个或多个全局中间件注册到服务器。全局中间件可以独立使用，无需服务处理程序，服务处理程序会拦截服务处理程序之前或之后的所有动态请求。该参数 `pattern` 指定中间件截获的路由模式，通常是“模糊”模式，如“/：name”、“/*any”或“/{field}”。

#### (*Server) BindMiddlewareDefault

```go
func (s *Server) BindMiddlewareDefault(handlers ...HandlerFunc)
```

BindMiddlewareDefault registers one or more global middleware to the server using default pattern “/*”. Global middleware can be used standalone without service handler, which intercepts all dynamic requests before or after service handler.

​	BindMiddlewareDefault 使用默认模式“/*”将一个或多个全局中间件注册到服务器。全局中间件可以独立使用，无需服务处理程序，服务处理程序会拦截服务处理程序之前或之后的所有动态请求。

#### (*Server) BindObject

```go
func (s *Server) BindObject(pattern string, object interface{}, method ...string)
```

BindObject registers object to server routes with a given pattern.

​	BindObject 使用给定模式将对象注册到服务器路由。

The optional parameter `method` is used to specify the method to be registered, which supports multiple method names; multiple methods are separated by char ‘,’, case-sensitive.

​	可选参数 `method` 用于指定要注册的方法，支持多个方法名称;多个方法用字符“，”分隔，区分大小写。

#### (*Server) BindObjectMethod

```go
func (s *Server) BindObjectMethod(pattern string, object interface{}, method string)
```

BindObjectMethod registers specified method of the object to server routes with a given pattern.

​	BindObjectMethod 使用给定模式将对象的指定方法注册到服务器路由。

The optional parameter `method` is used to specify the method to be registered, which does not support multiple method names but only one, case-sensitive.

​	可选参数 `method` 用于指定要注册的方法，该方法不支持多个方法名称，而仅支持一个方法名称，区分大小写。

#### (*Server) BindObjectRest

```go
func (s *Server) BindObjectRest(pattern string, object interface{})
```

BindObjectRest registers object in REST API styles to server with a specified pattern.

​	BindObjectRest 使用 REST API 样式将对象注册到具有指定模式的服务器。

#### (*Server) BindStatusHandler

```go
func (s *Server) BindStatusHandler(status int, handler HandlerFunc)
```

BindStatusHandler registers handler for given status code.

#### (*Server) BindStatusHandlerByMap

```go
func (s *Server) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc)
```

BindStatusHandlerByMap registers handler for given status code using map.

#### (*Server) Domain

```go
func (s *Server) Domain(domains string) *Domain
```

Domain creates and returns a domain object for management for one or more domains.

​	Domain 创建并返回用于管理一个或多个域的域对象。

#### (*Server) EnableAdmin

```go
func (s *Server) EnableAdmin(pattern ...string)
```

EnableAdmin enables the administration feature for the process. The optional parameter `pattern` specifies the URI for the administration page.

​	EnableAdmin 启用进程的管理功能。optional 参数 `pattern` 指定管理页的 URI。

#### (*Server) EnableHTTPS

```go
func (s *Server) EnableHTTPS(certFile, keyFile string, tlsConfig ...*tls.Config)
```

EnableHTTPS enables HTTPS with given certification and key files for the server. The optional parameter `tlsConfig` specifies custom TLS configuration.

​	EnableHTTPS 使用给定的证书和服务器密钥文件启用 HTTPS。可选参数 `tlsConfig` 指定自定义 TLS 配置。

#### (*Server) EnablePProf

```go
func (s *Server) EnablePProf(pattern ...string)
```

EnablePProf enables PProf feature for server.

​	EnablePProf 为服务器启用 PProf 功能。

#### (*Server) GetCookieDomain

```go
func (s *Server) GetCookieDomain() string
```

GetCookieDomain returns CookieDomain of server.

#### (*Server) GetCookieHttpOnly

```go
func (s *Server) GetCookieHttpOnly() bool
```

#### (*Server) GetCookieMaxAge

```go
func (s *Server) GetCookieMaxAge() time.Duration
```

GetCookieMaxAge returns the CookieMaxAge of the server.

#### (*Server) GetCookiePath

```go
func (s *Server) GetCookiePath() string
```

GetCookiePath returns the CookiePath of server.

#### (*Server) GetCookieSameSite

```go
func (s *Server) GetCookieSameSite() http.SameSite
```

GetCookieSameSite return CookieSameSite of server.

#### (*Server) GetCookieSecure

```go
func (s *Server) GetCookieSecure() bool
```

#### (*Server) GetHandler

```go
func (s *Server) GetHandler() func(w http.ResponseWriter, r *http.Request)
```

GetHandler returns the request handler of the server.

#### (*Server) GetIndexFiles

```go
func (s *Server) GetIndexFiles() []string
```

GetIndexFiles retrieves and returns the index files from the server.

​	GetIndexFiles 从服务器检索并返回索引文件。

#### (*Server) GetListenedAddress

```go
func (s *Server) GetListenedAddress() string
```

GetListenedAddress retrieves and returns the address string which are listened by current server.

​	GetListenedAddress 检索并返回当前服务器侦听的地址字符串。

#### (*Server) GetListenedPort

```go
func (s *Server) GetListenedPort() int
```

GetListenedPort retrieves and returns one port which is listened by current server.

​	GetListenedPort 检索并返回当前服务器侦听的一个端口。

#### (*Server) GetListenedPorts

```go
func (s *Server) GetListenedPorts() []int
```

GetListenedPorts retrieves and returns the ports which are listened by current server.

​	GetListenedPorts 检索并返回当前服务器侦听的端口。

#### (*Server) GetLogPath

```go
func (s *Server) GetLogPath() string
```

GetLogPath returns the log path.

#### (*Server) GetName

```go
func (s *Server) GetName() string
```

GetName returns the name of the server.

​	GetName 返回服务器的名称。

#### (*Server) GetOpenApi

```go
func (s *Server) GetOpenApi() *goai.OpenApiV3
```

GetOpenApi returns the OpenApi specification management object of current server.

​	GetOpenApi 返回当前服务器的 OpenApi 规范管理对象。

#### (*Server) GetRegistrar

```go
func (s *Server) GetRegistrar() gsvc.Registrar
```

GetRegistrar returns the Registrar of server.

#### (*Server) GetRoutes

```go
func (s *Server) GetRoutes() []RouterItem
```

GetRoutes retrieves and returns the router array.

​	GetRoutes 检索并返回路由器阵列。

#### (*Server) GetSessionCookieMaxAge

```go
func (s *Server) GetSessionCookieMaxAge() time.Duration
```

GetSessionCookieMaxAge returns the SessionCookieMaxAge of server.

​	GetSessionCookieMaxAge 返回服务器的 SessionCookieMaxAge。

#### (*Server) GetSessionIdName

```go
func (s *Server) GetSessionIdName() string
```

GetSessionIdName returns the SessionIdName of server.

​	GetSessionIdName 返回服务器的 SessionIdName。

#### (*Server) GetSessionMaxAge

```go
func (s *Server) GetSessionMaxAge() time.Duration
```

GetSessionMaxAge returns the SessionMaxAge of server.

​	GetSessionMaxAge 返回服务器的 SessionMaxAge。

#### (*Server) Group

```go
func (s *Server) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup
```

Group creates and returns a RouterGroup object.

​	Group 创建并返回 RouterGroup 对象。

#### (*Server) IsAccessLogEnabled

```go
func (s *Server) IsAccessLogEnabled() bool
```

IsAccessLogEnabled checks whether the access log enabled.

​	IsAccessLogEnabled检查访问日志是否开启。

#### (*Server) IsErrorLogEnabled

```go
func (s *Server) IsErrorLogEnabled() bool
```

IsErrorLogEnabled checks whether the error log enabled.

#### (*Server) Logger

```go
func (s *Server) Logger() *glog.Logger
```

Logger is alias of GetLogger.

#### (*Server) Plugin

```go
func (s *Server) Plugin(plugin ...Plugin)
```

Plugin adds plugin to the server.

#### (*Server) Run

```go
func (s *Server) Run()
```

Run starts server listening in blocking way. It’s commonly used for single server situation.

​	Run 以阻塞方式启动服务器侦听。它通常用于单服务器情况。

##### Example

#### (*Server) ServeHTTP

```go
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

ServeHTTP is the default handler for http request. It should not create new goroutine handling the request as it’s called by am already created new goroutine from http.Server.

​	ServeHTTP 是 http 请求的默认处理程序。它不应该创建处理请求的新 goroutine，因为它是由 am 调用的，已经从 http 创建了新的 goroutine。服务器。

This function also makes serve implementing the interface of http.Handler.

​	这个函数也使 serve 实现了 http 的接口。处理器。

#### (*Server) SetAccessLogEnabled

```go
func (s *Server) SetAccessLogEnabled(enabled bool)
```

SetAccessLogEnabled enables/disables the access log.

​	SetAccessLogEnabled 启用/禁用访问日志。

#### (*Server) SetAddr

```go
func (s *Server) SetAddr(address string)
```

SetAddr sets the listening address for the server. The address is like ‘:80’, ‘0.0.0.0:80’, ‘127.0.0.1:80’, ‘180.18.99.10:80’, etc.

#### (*Server) SetClientMaxBodySize

```go
func (s *Server) SetClientMaxBodySize(maxSize int64)
```

SetClientMaxBodySize sets the ClientMaxBodySize for server.

#### (*Server) SetConfig

```go
func (s *Server) SetConfig(c ServerConfig) error
```

SetConfig sets the configuration for the server.

#### (*Server) SetConfigWithMap

```go
func (s *Server) SetConfigWithMap(m map[string]interface{}) error
```

SetConfigWithMap sets the configuration for the server using map.

​	SetConfigWithMap 使用 map 设置服务器的配置。

#### (*Server) SetCookieDomain

```go
func (s *Server) SetCookieDomain(domain string)
```

SetCookieDomain sets the CookieDomain for server.

​	SetCookieDomain 设置服务器的 CookieDomain。

#### (*Server) SetCookieMaxAge

```go
func (s *Server) SetCookieMaxAge(ttl time.Duration)
```

SetCookieMaxAge sets the CookieMaxAge for server.

​	SetCookieMaxAge 设置服务器的 CookieMaxAge。

#### (*Server) SetCookiePath

```go
func (s *Server) SetCookiePath(path string)
```

SetCookiePath sets the CookiePath for server.

​	SetCookiePath 设置服务器的 CookiePath。

#### (*Server) SetDumpRouterMap

```go
func (s *Server) SetDumpRouterMap(enabled bool)
```

SetDumpRouterMap sets the DumpRouterMap for server. If DumpRouterMap is enabled, it automatically dumps the route map when server starts.

​	SetDumpRouterMap 设置服务器的 DumpRouterMap。如果启用了 DumpRouterMap，则会在服务器启动时自动转储路由映射。

#### (*Server) SetEndpoints

```go
func (s *Server) SetEndpoints(endpoints []string)
```

SetEndpoints sets the Endpoints for the server.

​	SetEndpoints 设置服务器的终结点。

#### (*Server) SetErrorLogEnabled

```go
func (s *Server) SetErrorLogEnabled(enabled bool)
```

SetErrorLogEnabled enables/disables the error log.

#### (*Server) SetErrorStack

```go
func (s *Server) SetErrorStack(enabled bool)
```

SetErrorStack enables/disables the error stack feature.

​	SetErrorStack 启用/禁用错误堆栈功能。

#### (*Server) SetFileServerEnabled

```go
func (s *Server) SetFileServerEnabled(enabled bool)
```

SetFileServerEnabled enables/disables the static file service. It’s the main switch for the static file service. When static file service configuration functions like SetServerRoot, AddSearchPath and AddStaticPath are called, this configuration is automatically enabled.

​	SetFileServerEnabled 启用/禁用静态文件服务。它是静态文件服务的主开关。调用 SetServerRoot、AddSearchPath 和 AddStaticPath 等静态文件服务配置函数时，将自动启用此配置。

#### (*Server) SetFormParsingMemory

```go
func (s *Server) SetFormParsingMemory(maxMemory int64)
```

SetFormParsingMemory sets the FormParsingMemory for server.

#### (*Server) SetHTTPSAddr

```go
func (s *Server) SetHTTPSAddr(address string)
```

SetHTTPSAddr sets the HTTPS listening ports for the server.

#### (*Server) SetHTTPSPort

```go
func (s *Server) SetHTTPSPort(port ...int)
```

SetHTTPSPort sets the HTTPS listening ports for the server. The listening ports can be multiple like: SetHTTPSPort(443, 500).

#### (*Server) SetHandler

```go
func (s *Server) SetHandler(h func(w http.ResponseWriter, r *http.Request))
```

SetHandler sets the request handler for server.

#### (*Server) SetIdleTimeout

```go
func (s *Server) SetIdleTimeout(t time.Duration)
```

SetIdleTimeout sets the IdleTimeout for the server.

​	SetIdleTimeout 设置服务器的 IdleTimeout。

#### (*Server) SetIndexFiles

```go
func (s *Server) SetIndexFiles(indexFiles []string)
```

SetIndexFiles sets the index files for server.

​	SetIndexFiles 设置服务器的索引文件。

#### (*Server) SetIndexFolder

```go
func (s *Server) SetIndexFolder(enabled bool)
```

SetIndexFolder enables/disables listing the sub-files if requesting a directory.

​	如果请求目录，SetIndexFolder 启用/禁用列出子文件。

#### (*Server) SetKeepAlive

```go
func (s *Server) SetKeepAlive(enabled bool)
```

SetKeepAlive sets the KeepAlive for the server.

​	SetKeepAlive 设置服务器的 KeepAlive。

#### (*Server) SetListener

```go
func (s *Server) SetListener(listeners ...net.Listener) error
```

SetListener set the custom listener for the server.

​	SetListener 设置服务器的自定义侦听器。

#### (*Server) SetLogLevel

```go
func (s *Server) SetLogLevel(level string)
```

SetLogLevel sets logging level by level string.

#### (*Server) SetLogPath

```go
func (s *Server) SetLogPath(path string) error
```

SetLogPath sets the log path for server. It logs content to file only if the log path is set.

​	SetLogPath 设置服务器的日志路径。仅当设置了日志路径时，它才会将内容记录到文件中。

#### (*Server) SetLogStdout

```go
func (s *Server) SetLogStdout(enabled bool)
```

SetLogStdout sets whether output the logging content to stdout.

​	SetLogStdout 设置是否将日志记录内容输出到 stdout。

#### (*Server) SetLogger

```go
func (s *Server) SetLogger(logger *glog.Logger)
```

SetLogger sets the logger for logging responsibility. Note that it cannot be set in runtime as there may be concurrent safety issue.

​	SetLogger 设置记录器的日志记录责任。请注意，它不能在运行时设置，因为可能存在并发安全问题。

#### (*Server) SetMaxHeaderBytes

```go
func (s *Server) SetMaxHeaderBytes(b int)
```

SetMaxHeaderBytes sets the MaxHeaderBytes for the server.

#### (*Server) SetName

```go
func (s *Server) SetName(name string)
```

SetName sets the name for the server.

​	SetName 设置服务器的名称。

#### (*Server) SetNameToUriType

```go
func (s *Server) SetNameToUriType(t int)
```

SetNameToUriType sets the NameToUriType for server.

​	SetNameToUriType 设置服务器的 NameToUriType。

#### (*Server) SetOpenApiPath

```go
func (s *Server) SetOpenApiPath(path string)
```

SetOpenApiPath sets the OpenApiPath for server.

​	SetOpenApiPath 设置服务器的 OpenApiPath。

#### (*Server) SetPort

```go
func (s *Server) SetPort(port ...int)
```

SetPort sets the listening ports for the server. The listening ports can be multiple like: SetPort(80, 8080).

​	SetPort 设置服务器的侦听端口。侦听端口可以是多个，例如：SetPort（80， 8080）。

#### (*Server) SetReadTimeout

```go
func (s *Server) SetReadTimeout(t time.Duration)
```

SetReadTimeout sets the ReadTimeout for the server.

​	SetReadTimeout 设置服务器的 ReadTimeout。

#### (*Server) SetRegistrar

```go
func (s *Server) SetRegistrar(registrar gsvc.Registrar)
```

SetRegistrar sets the Registrar for server.

#### (*Server) SetRewrite

```go
func (s *Server) SetRewrite(uri string, rewrite string)
```

SetRewrite sets rewrites for static URI for server.

#### (*Server) SetRewriteMap

```go
func (s *Server) SetRewriteMap(rewrites map[string]string)
```

SetRewriteMap sets the rewritten map for server.

​	SetRewriteMap 为服务器设置重写的映射。

#### (*Server) SetRouteOverWrite

```go
func (s *Server) SetRouteOverWrite(enabled bool)
```

SetRouteOverWrite sets the RouteOverWrite for server.

​	SetRouteOverWrite 设置服务器的 RouteOverWrite。

#### (*Server) SetServerAgent

```go
func (s *Server) SetServerAgent(agent string)
```

SetServerAgent sets the ServerAgent for the server.

​	SetServerAgent 设置服务器的 ServerAgent。

#### (*Server) SetServerRoot

```go
func (s *Server) SetServerRoot(root string)
```

SetServerRoot sets the document root for static service.

​	SetServerRoot 设置静态服务的文档根目录。

#### (*Server) SetSessionCookieMaxAge

```go
func (s *Server) SetSessionCookieMaxAge(maxAge time.Duration)
```

SetSessionCookieMaxAge sets the SessionCookieMaxAge for server.

#### (*Server) SetSessionCookieOutput

```go
func (s *Server) SetSessionCookieOutput(enabled bool)
```

SetSessionCookieOutput sets the SetSessionCookieOutput for server.

#### (*Server) SetSessionIdName

```go
func (s *Server) SetSessionIdName(name string)
```

SetSessionIdName sets the SessionIdName for server.

​	SetSessionIdName 设置服务器的 SessionIdName。

#### (*Server) SetSessionMaxAge

```go
func (s *Server) SetSessionMaxAge(ttl time.Duration)
```

SetSessionMaxAge sets the SessionMaxAge for server.

​	SetSessionMaxAge 设置服务器的 SessionMaxAge。

#### (*Server) SetSessionStorage

```go
func (s *Server) SetSessionStorage(storage gsession.Storage)
```

SetSessionStorage sets the SessionStorage for server.

​	SetSessionStorage 设置服务器的 SessionStorage。

#### (*Server) SetSwaggerPath

```go
func (s *Server) SetSwaggerPath(path string)
```

SetSwaggerPath sets the SwaggerPath for server.

​	SetSwaggerPath 设置服务器的 SwaggerPath。

#### (*Server) SetSwaggerUITemplate

```go
func (s *Server) SetSwaggerUITemplate(swaggerUITemplate string)
```

SetSwaggerUITemplate sets the Swagger template for server.

​	SetSwaggerUITemplate 设置服务器的 Swagger 模板。

#### (*Server) SetTLSConfig

```go
func (s *Server) SetTLSConfig(tlsConfig *tls.Config)
```

SetTLSConfig sets custom TLS configuration and enables HTTPS feature for the server.

​	SetTLSConfig 设置自定义 TLS 配置并为服务器启用 HTTPS 功能。

#### (*Server) SetView

```go
func (s *Server) SetView(view *gview.View)
```

SetView sets the View for the server.

​	SetView 设置服务器的视图。

#### (*Server) SetWriteTimeout

```go
func (s *Server) SetWriteTimeout(t time.Duration)
```

SetWriteTimeout sets the WriteTimeout for the server.

​	SetWriteTimeout 设置服务器的 WriteTimeout。

#### (*Server) Shutdown

```go
func (s *Server) Shutdown() error
```

Shutdown shuts down current server.

#### (*Server) Start

```go
func (s *Server) Start() error
```

Start starts listening on configured port. This function does not block the process, you can use function Wait blocking the process.

​	Start 开始侦听配置的端口。此函数不会阻塞进程，可以使用函数 Wait 阻塞进程。

#### (*Server) Status

```go
func (s *Server) Status() ServerStatus
```

Status retrieves and returns the server status.

#### (*Server) Use

```go
func (s *Server) Use(handlers ...HandlerFunc)
```

Use is the alias of BindMiddlewareDefault. See BindMiddlewareDefault.

​	Use 是 BindMiddlewareDefault 的别名。请参阅 BindMiddlewareDefault。

### type ServerConfig

```go
type ServerConfig struct {

	// Service name, which is for service registry and discovery.
	Name string `json:"name"`

	// Address specifies the server listening address like "port" or ":port",
	// multiple addresses joined using ','.
	Address string `json:"address"`

	// HTTPSAddr specifies the HTTPS addresses, multiple addresses joined using char ','.
	HTTPSAddr string `json:"httpsAddr"`

	// Listeners specifies the custom listeners.
	Listeners []net.Listener `json:"listeners"`

	// Endpoints are custom endpoints for service register, it uses Address if empty.
	Endpoints []string `json:"endpoints"`

	// HTTPSCertPath specifies certification file path for HTTPS service.
	HTTPSCertPath string `json:"httpsCertPath"`

	// HTTPSKeyPath specifies the key file path for HTTPS service.
	HTTPSKeyPath string `json:"httpsKeyPath"`

	// TLSConfig optionally provides a TLS configuration for use
	// by ServeTLS and ListenAndServeTLS. Note that this value is
	// cloned by ServeTLS and ListenAndServeTLS, so it's not
	// possible to modify the configuration with methods like
	// tls.Config.SetSessionTicketKeys. To use
	// SetSessionTicketKeys, use Server.Serve with a TLS Listener
	// instead.
	TLSConfig *tls.Config `json:"tlsConfig"`

	// Handler the handler for HTTP request.
	Handler func(w http.ResponseWriter, r *http.Request) `json:"-"`

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout time.Duration `json:"readTimeout"`

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	WriteTimeout time.Duration `json:"writeTimeout"`

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alive are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, there is no timeout.
	IdleTimeout time.Duration `json:"idleTimeout"`

	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	//
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's 10240 bytes in default.
	MaxHeaderBytes int `json:"maxHeaderBytes"`

	// KeepAlive enables HTTP keep-alive.
	KeepAlive bool `json:"keepAlive"`

	// ServerAgent specifies the server agent information, which is wrote to
	// HTTP response header as "Server".
	ServerAgent string `json:"serverAgent"`

	// View specifies the default template view object for the server.
	View *gview.View `json:"view"`

	// Rewrites specifies the URI rewrite rules map.
	Rewrites map[string]string `json:"rewrites"`

	// IndexFiles specifies the index files for static folder.
	IndexFiles []string `json:"indexFiles"`

	// IndexFolder specifies if listing sub-files when requesting folder.
	// The server responses HTTP status code 403 if it is false.
	IndexFolder bool `json:"indexFolder"`

	// ServerRoot specifies the root directory for static service.
	ServerRoot string `json:"serverRoot"`

	// SearchPaths specifies additional searching directories for static service.
	SearchPaths []string `json:"searchPaths"`

	// StaticPaths specifies URI to directory mapping array.
	StaticPaths []staticPathItem `json:"staticPaths"`

	// FileServerEnabled is the global switch for static service.
	// It is automatically set enabled if any static path is set.
	FileServerEnabled bool `json:"fileServerEnabled"`

	// CookieMaxAge specifies the max TTL for cookie items.
	CookieMaxAge time.Duration `json:"cookieMaxAge"`

	// CookiePath specifies cookie path.
	// It also affects the default storage for session id.
	CookiePath string `json:"cookiePath"`

	// CookieDomain specifies cookie domain.
	// It also affects the default storage for session id.
	CookieDomain string `json:"cookieDomain"`

	// CookieSameSite specifies cookie SameSite property.
	// It also affects the default storage for session id.
	CookieSameSite string `json:"cookieSameSite"`

	// CookieSameSite specifies cookie Secure property.
	// It also affects the default storage for session id.
	CookieSecure bool `json:"cookieSecure"`

	// CookieSameSite specifies cookie HttpOnly property.
	// It also affects the default storage for session id.
	CookieHttpOnly bool `json:"cookieHttpOnly"`

	// SessionIdName specifies the session id name.
	SessionIdName string `json:"sessionIdName"`

	// SessionMaxAge specifies max TTL for session items.
	SessionMaxAge time.Duration `json:"sessionMaxAge"`

	// SessionPath specifies the session storage directory path for storing session files.
	// It only makes sense if the session storage is type of file storage.
	SessionPath string `json:"sessionPath"`

	// SessionStorage specifies the session storage.
	SessionStorage gsession.Storage `json:"sessionStorage"`

	// SessionCookieMaxAge specifies the cookie ttl for session id.
	// If it is set 0, it means it expires along with browser session.
	SessionCookieMaxAge time.Duration `json:"sessionCookieMaxAge"`

	// SessionCookieOutput specifies whether automatic outputting session id to cookie.
	SessionCookieOutput bool `json:"sessionCookieOutput"`

	Logger           *glog.Logger `json:"logger"`           // Logger specifies the logger for server.
	LogPath          string       `json:"logPath"`          // LogPath specifies the directory for storing logging files.
	LogLevel         string       `json:"logLevel"`         // LogLevel specifies the logging level for logger.
	LogStdout        bool         `json:"logStdout"`        // LogStdout specifies whether printing logging content to stdout.
	ErrorStack       bool         `json:"errorStack"`       // ErrorStack specifies whether logging stack information when error.
	ErrorLogEnabled  bool         `json:"errorLogEnabled"`  // ErrorLogEnabled enables error logging content to files.
	ErrorLogPattern  string       `json:"errorLogPattern"`  // ErrorLogPattern specifies the error log file pattern like: error-{Ymd}.log
	AccessLogEnabled bool         `json:"accessLogEnabled"` // AccessLogEnabled enables access logging content to files.
	AccessLogPattern string       `json:"accessLogPattern"` // AccessLogPattern specifies the error log file pattern like: access-{Ymd}.log

	PProfEnabled bool   `json:"pprofEnabled"` // PProfEnabled enables PProf feature.
	PProfPattern string `json:"pprofPattern"` // PProfPattern specifies the PProf service pattern for router.

	OpenApiPath       string `json:"openapiPath"`       // OpenApiPath specifies the OpenApi specification file path.
	SwaggerPath       string `json:"swaggerPath"`       // SwaggerPath specifies the swagger UI path for route registering.
	SwaggerUITemplate string `json:"swaggerUITemplate"` // SwaggerUITemplate specifies the swagger UI custom template

	// ClientMaxBodySize specifies the max body size limit in bytes for client request.
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's `8MB` in default.
	ClientMaxBodySize int64 `json:"clientMaxBodySize"`

	// FormParsingMemory specifies max memory buffer size in bytes which can be used for
	// parsing multimedia form.
	// It can be configured in configuration file using string like: 1m, 10m, 500kb etc.
	// It's 1MB in default.
	FormParsingMemory int64 `json:"formParsingMemory"`

	// NameToUriType specifies the type for converting struct method name to URI when
	// registering routes.
	NameToUriType int `json:"nameToUriType"`

	// RouteOverWrite allows to overwrite the route if duplicated.
	RouteOverWrite bool `json:"routeOverWrite"`

	// DumpRouterMap specifies whether automatically dumps router map when server starts.
	DumpRouterMap bool `json:"dumpRouterMap"`

	// Graceful enables graceful reload feature for all servers of the process.
	Graceful bool `json:"graceful"`

	// GracefulTimeout set the maximum survival time (seconds) of the parent process.
	GracefulTimeout uint8 `json:"gracefulTimeout"`

	// GracefulShutdownTimeout set the maximum survival time (seconds) before stopping the server.
	GracefulShutdownTimeout uint8 `json:"gracefulShutdownTimeout"`
}
```

ServerConfig is the HTTP Server configuration manager.

​	ServerConfig 是 HTTP 服务器配置管理器。

#### func ConfigFromMap

```go
func ConfigFromMap(m map[string]interface{}) (ServerConfig, error)
```

ConfigFromMap creates and returns a ServerConfig object with given map and default configuration object.

​	ConfigFromMap 创建并返回具有给定映射和默认配置对象的 ServerConfig 对象。

#### func NewConfig

```go
func NewConfig() ServerConfig
```

NewConfig creates and returns a ServerConfig object with default configurations. Note that, do not define this default configuration to local package variable, as there are some pointer attributes that may be shared in different servers.

​	NewConfig 创建并返回具有默认配置的 ServerConfig 对象。请注意，不要将此默认配置定义为本地包变量，因为某些指针属性可能在不同的服务器中共享。

### type ServerStatus <-2.5.0

```go
type ServerStatus = int
```

ServerStatus is the server status enum type.

### type Session

```go
type Session = gsession.Session
```

Session is actually an alias of gsession.Session, which is bound to a single request.

​	Session 实际上是 gsession 的别名。会话，绑定到单个请求。

### type UploadFile

```go
type UploadFile struct {
	*multipart.FileHeader `json:"-"`
	// contains filtered or unexported fields
}
```

UploadFile wraps the multipart uploading file with more and convenient features.

​	UploadFile 为分段上传文件包装了更多便捷的功能。

#### (UploadFile) MarshalJSON

```go
func (f UploadFile) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*UploadFile) Save

```go
func (f *UploadFile) Save(dirPath string, randomlyRename ...bool) (filename string, err error)
```

Save saves the single uploading file to directory path and returns the saved file name.

​	Save 将单个上传文件保存到目录路径，并返回保存的文件名。

The parameter `dirPath` should be a directory path, or it returns error.

​	该参数 `dirPath` 应为目录路径，否则将返回错误。

Note that it will OVERWRITE the target file if there’s already a same name file exist.

​	请注意，如果已经存在同名文件，它将覆盖目标文件。

##### Example

### type UploadFiles

```go
type UploadFiles []*UploadFile
```

UploadFiles is an array type of *UploadFile.

#### (UploadFiles) Save

```go
func (fs UploadFiles) Save(dirPath string, randomlyRename ...bool) (filenames []string, err error)
```

Save saves all uploading files to specified directory path and returns the saved file names.

​	Save 将所有上传文件保存到指定的目录路径，并返回保存的文件名。

The parameter `dirPath` should be a directory path or it returns error.

​	该参数 `dirPath` 应为目录路径，否则将返回错误。

The parameter `randomlyRename` specifies whether randomly renames all the file names.

​	该参数 `randomlyRename` 指定是否随机重命名所有文件名。

### type WebSocket

```go
type WebSocket struct {
	*websocket.Conn
}
```

WebSocket wraps the underlying websocket connection and provides convenient functions.

​	WebSocket 包装底层 websocket 连接，提供便捷的功能。