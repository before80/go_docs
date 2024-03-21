+++
title = "ghttp"
date = 2024-03-21T17:52:53+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/ghttp

Package ghttp provides powerful http server and simple client implements.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/ghttp/ghttp.go#L133)

``` go
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

``` go
const (
	UriTypeDefault  = iota // Method names to the URI converting type, which converts name to its lower case and joins the words using char '-'.
	UriTypeFullName        // Method names to the URI converting type, which does not convert to the method name.
	UriTypeAllLower        // Method names to the URI converting type, which converts name to its lower case.
	UriTypeCamel           // Method names to the URI converting type, which converts name to its camel case.
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/ghttp/ghttp_server_websocket.go#L17)

``` go
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

``` go
const (
	// FreePortAddress marks the server listens using random free port.
	FreePortAddress = ":0"
)
```

### Variables 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/ghttp/ghttp.go#L208)

``` go
var (
	ErrNeedJsonBody = gerror.NewWithOption(gerror.Option{
		Text: "the request body content should be JSON format",
		Code: gcode.CodeInvalidRequest,
	})
)
```

### Functions 

##### func BuildParams 

``` go
func BuildParams(params interface{}, noUrlEncode ...bool) (encodedParamStr string)
```

BuildParams builds the request string for the http client. The `params` can be type of: string/[]byte/map/struct/*struct.

The optional parameter `noUrlEncode` specifies whether to ignore the url encoding for the data.

##### func MiddlewareCORS 

``` go
func MiddlewareCORS(r *Request)
```

MiddlewareCORS is a middleware handler for CORS with default options.

##### func MiddlewareHandlerResponse 

``` go
func MiddlewareHandlerResponse(r *Request)
```

MiddlewareHandlerResponse is the default middleware handling handler response object and its error.

##### func MiddlewareJsonBody <-2.1.3

``` go
func MiddlewareJsonBody(r *Request)
```

MiddlewareJsonBody validates and returns request body whether JSON format.

##### func MiddlewareNeverDoneCtx <-2.6.2

``` go
func MiddlewareNeverDoneCtx(r *Request)
```

MiddlewareNeverDoneCtx sets the context never done for current process.

##### func RestartAllServer 

``` go
func RestartAllServer(ctx context.Context, newExeFilePath string) error
```

RestartAllServer restarts all the servers of the process gracefully. The optional parameter `newExeFilePath` specifies the new binary file for creating process.

##### func ShutdownAllServer 

``` go
func ShutdownAllServer(ctx context.Context) error
```

ShutdownAllServer shuts down all servers of current process gracefully.

##### func StartPProfServer 

``` go
func StartPProfServer(port int, pattern ...string)
```

StartPProfServer starts and runs a new server for pprof.

##### func SupportedMethods <-2.4.2

``` go
func SupportedMethods() []string
```

SupportedMethods returns all supported HTTP methods.

##### func Wait 

``` go
func Wait()
```

Wait blocks to wait for all servers done. It's commonly used in multiple server situation.

### Types 

#### type CORSOptions 

``` go
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

#### type Cookie 

``` go
type Cookie struct {
	// contains filtered or unexported fields
}
```

Cookie for HTTP COOKIE management.

##### func GetCookie 

``` go
func GetCookie(r *Request) *Cookie
```

GetCookie creates or retrieves a cookie object with given request. It retrieves and returns an existing cookie object if it already exists with given request. It creates and returns a new cookie object if it does not exist with given request.

##### (*Cookie) Contains 

``` go
func (c *Cookie) Contains(key string) bool
```

Contains checks if given key exists and not expire in cookie.

##### (*Cookie) Flush 

``` go
func (c *Cookie) Flush()
```

Flush outputs the cookie items to the client.

##### (*Cookie) Get 

``` go
func (c *Cookie) Get(key string, def ...string) *gvar.Var
```

Get retrieves and returns the value with specified key. It returns `def` if specified key does not exist and `def` is given.

##### (*Cookie) GetSessionId 

``` go
func (c *Cookie) GetSessionId() string
```

GetSessionId retrieves and returns the session id from cookie.

##### (*Cookie) Map 

``` go
func (c *Cookie) Map() map[string]string
```

Map returns the cookie items as map[string]string.

##### (*Cookie) Remove 

``` go
func (c *Cookie) Remove(key string)
```

Remove deletes specified key and its value from cookie using default domain and path. It actually tells the http client that the cookie is expired, do not send it to server next time.

##### (*Cookie) RemoveCookie 

``` go
func (c *Cookie) RemoveCookie(key, domain, path string)
```

RemoveCookie deletes specified key and its value from cookie using given domain and path. It actually tells the http client that the cookie is expired, do not send it to server next time.

##### (*Cookie) Set 

``` go
func (c *Cookie) Set(key, value string)
```

Set sets cookie item with default domain, path and expiration age.

##### (*Cookie) SetCookie 

``` go
func (c *Cookie) SetCookie(key, value, domain, path string, maxAge time.Duration, options ...CookieOptions)
```

SetCookie sets cookie item with given domain, path and expiration age. The optional parameter `options` specifies extra security configurations, which is usually empty.

##### (*Cookie) SetHttpCookie 

``` go
func (c *Cookie) SetHttpCookie(httpCookie *http.Cookie)
```

SetHttpCookie sets cookie with *http.Cookie.

##### (*Cookie) SetSessionId 

``` go
func (c *Cookie) SetSessionId(id string)
```

SetSessionId sets session id in the cookie.

#### type CookieOptions 

``` go
type CookieOptions struct {
	SameSite http.SameSite // cookie SameSite property
	Secure   bool          // cookie Secure property
	HttpOnly bool          // cookie HttpOnly property
}
```

CookieOptions provides security config for cookies

#### type DefaultHandlerResponse 

``` go
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}
```

DefaultHandlerResponse is the default implementation of HandlerResponse.

#### type Domain 

``` go
type Domain struct {
	// contains filtered or unexported fields
}
```

Domain is used for route register for domains.

##### (*Domain) BindHandler 

``` go
func (d *Domain) BindHandler(pattern string, handler interface{})
```

BindHandler binds the handler for the specified pattern.

##### (*Domain) BindHookHandler 

``` go
func (d *Domain) BindHookHandler(pattern string, hook HookName, handler HandlerFunc)
```

BindHookHandler binds the hook handler for the specified pattern.

##### (*Domain) BindHookHandlerByMap 

``` go
func (d *Domain) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc)
```

BindHookHandlerByMap binds the hook handler for the specified pattern.

##### (*Domain) BindMiddleware 

``` go
func (d *Domain) BindMiddleware(pattern string, handlers ...HandlerFunc)
```

BindMiddleware binds the middleware for the specified pattern.

##### (*Domain) BindMiddlewareDefault 

``` go
func (d *Domain) BindMiddlewareDefault(handlers ...HandlerFunc)
```

BindMiddlewareDefault binds the default middleware for the specified pattern.

##### (*Domain) BindObject 

``` go
func (d *Domain) BindObject(pattern string, obj interface{}, methods ...string)
```

BindObject binds the object for the specified pattern.

##### (*Domain) BindObjectMethod 

``` go
func (d *Domain) BindObjectMethod(pattern string, obj interface{}, method string)
```

BindObjectMethod binds the method for the specified pattern.

##### (*Domain) BindObjectRest 

``` go
func (d *Domain) BindObjectRest(pattern string, obj interface{})
```

BindObjectRest binds the RESTful API for the specified pattern.

##### (*Domain) BindStatusHandler 

``` go
func (d *Domain) BindStatusHandler(status int, handler HandlerFunc)
```

BindStatusHandler binds the status handler for the specified pattern.

##### (*Domain) BindStatusHandlerByMap 

``` go
func (d *Domain) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc)
```

BindStatusHandlerByMap binds the status handler for the specified pattern.

##### (*Domain) EnablePProf 

``` go
func (d *Domain) EnablePProf(pattern ...string)
```

EnablePProf enables PProf feature for server of specified domain.

##### (*Domain) Group 

``` go
func (d *Domain) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup
```

Group creates and returns a RouterGroup object, which is bound to a specified domain.

##### (*Domain) Use 

``` go
func (d *Domain) Use(handlers ...HandlerFunc)
```

Use adds middleware to the domain.

#### type HandlerFunc 

``` go
type HandlerFunc = func(r *Request)
```

HandlerFunc is request handler function.

##### func WrapF 

``` go
func WrapF(f http.HandlerFunc) HandlerFunc
```

WrapF is a helper function for wrapping http.HandlerFunc and returns a ghttp.HandlerFunc.

##### func WrapH 

``` go
func WrapH(h http.Handler) HandlerFunc
```

WrapH is a helper function for wrapping http.Handler and returns a ghttp.HandlerFunc.

#### type HandlerItem <-2.1.0

``` go
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

##### (HandlerItem) MarshalJSON <-2.1.0

``` go
func (item HandlerItem) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

#### type HandlerItemParsed <-2.2.2

``` go
type HandlerItemParsed struct {
	Handler *HandlerItem      // Handler information.
	Values  map[string]string // Router values parsed from URL.Path.
}
```

HandlerItemParsed is the item parsed from URL.Path.

##### (HandlerItemParsed) MarshalJSON <-2.2.2

``` go
func (item HandlerItemParsed) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

#### type HandlerType <-2.5.0

``` go
type HandlerType string
```

HandlerType is the route handler enum type.

#### type HookName <-2.5.0

``` go
type HookName string
```

HookName is the route hook name enum type.

#### type Plugin 

``` go
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

#### type Request 

``` go
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

##### func RequestFromCtx 

``` go
func RequestFromCtx(ctx context.Context) *Request
```

RequestFromCtx retrieves and returns the Request object from context.

##### (*Request) Assign 

``` go
func (r *Request) Assign(key string, value interface{})
```

Assign binds a template variable to current request.

##### (*Request) Assigns 

``` go
func (r *Request) Assigns(data gview.Params)
```

Assigns binds multiple template variables to current request.

##### (*Request) BasicAuth 

``` go
func (r *Request) BasicAuth(user, pass string, tips ...string) bool
```

BasicAuth enables the http basic authentication feature with a given passport and password and asks client for authentication. It returns true if authentication success, else returns false if failure.

##### (*Request) Context 

``` go
func (r *Request) Context() context.Context
```

Context is alias for function GetCtx. This function overwrites the http.Request.Context function. See GetCtx.

##### (*Request) Exit 

``` go
func (r *Request) Exit()
```

Exit exits executing of current HTTP handler.

##### (*Request) ExitAll 

``` go
func (r *Request) ExitAll()
```

ExitAll exits executing of current and following HTTP handlers.

##### (*Request) ExitHook 

``` go
func (r *Request) ExitHook()
```

ExitHook exits executing of current and following HTTP HOOK handlers.

##### (*Request) Get 

``` go
func (r *Request) Get(key string, def ...interface{}) *gvar.Var
```

Get is alias of GetRequest, which is one of the most commonly used functions for retrieving parameter. See r.GetRequest.

##### (*Request) GetBody 

``` go
func (r *Request) GetBody() []byte
```

GetBody retrieves and returns request body content as bytes. It can be called multiple times retrieving the same body content.

##### (*Request) GetBodyString 

``` go
func (r *Request) GetBodyString() string
```

GetBodyString retrieves and returns request body content as string. It can be called multiple times retrieving the same body content.

##### (*Request) GetClientIp 

``` go
func (r *Request) GetClientIp() string
```

GetClientIp returns the client ip of this request without port. Note that this ip address might be modified by client header.

##### (*Request) GetCtx 

``` go
func (r *Request) GetCtx() context.Context
```

GetCtx retrieves and returns the request's context. Its alias of function Context,to be relevant with function SetCtx.

##### (*Request) GetCtxVar 

``` go
func (r *Request) GetCtxVar(key interface{}, def ...interface{}) *gvar.Var
```

GetCtxVar retrieves and returns a Var with a given key name. The optional parameter `def` specifies the default value of the Var if given `key` does not exist in the context.

##### (*Request) GetError 

``` go
func (r *Request) GetError() error
```

GetError returns the error occurs in the procedure of the request. It returns nil if there's no error.

##### (*Request) GetForm 

``` go
func (r *Request) GetForm(key string, def ...interface{}) *gvar.Var
```

GetForm retrieves and returns parameter `key` from form. It returns `def` if `key` does not exist in the form and `def` is given, or else it returns nil.

##### (*Request) GetFormMap 

``` go
func (r *Request) GetFormMap(kvMap ...map[string]interface{}) map[string]interface{}
```

GetFormMap retrieves and returns all form parameters passed from client as map. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

##### (*Request) GetFormMapStrStr 

``` go
func (r *Request) GetFormMapStrStr(kvMap ...map[string]interface{}) map[string]string
```

GetFormMapStrStr retrieves and returns all form parameters passed from client as map[string]string. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

##### (*Request) GetFormMapStrVar 

``` go
func (r *Request) GetFormMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var
```

GetFormMapStrVar retrieves and returns all form parameters passed from client as map[string]*gvar.Var. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

##### (*Request) GetFormStruct 

``` go
func (r *Request) GetFormStruct(pointer interface{}, mapping ...map[string]string) error
```

GetFormStruct retrieves all form parameters passed from client and converts them to given struct object. Note that the parameter `pointer` is a pointer to the struct object. The optional parameter `mapping` is used to specify the key to attribute mapping.

##### (*Request) GetHandlerResponse 

``` go
func (r *Request) GetHandlerResponse() interface{}
```

GetHandlerResponse retrieves and returns the handler response object and its error.

##### (*Request) GetHeader 

``` go
func (r *Request) GetHeader(key string) string
```

GetHeader retrieves and returns the header value with given `key`.

##### (*Request) GetHost 

``` go
func (r *Request) GetHost() string
```

GetHost returns current request host name, which might be a domain or an IP without port.

##### (*Request) GetJson 

``` go
func (r *Request) GetJson() (*gjson.Json, error)
```

GetJson parses current request content as JSON format, and returns the JSON object. Note that the request content is read from request BODY, not from any field of FORM.

##### (*Request) GetMap 

``` go
func (r *Request) GetMap(def ...map[string]interface{}) map[string]interface{}
```

GetMap is an alias and convenient function for GetRequestMap. See GetRequestMap.

##### (*Request) GetMapStrStr 

``` go
func (r *Request) GetMapStrStr(def ...map[string]interface{}) map[string]string
```

GetMapStrStr is an alias and convenient function for GetRequestMapStrStr. See GetRequestMapStrStr.

##### (*Request) GetMultipartFiles 

``` go
func (r *Request) GetMultipartFiles(name string) []*multipart.FileHeader
```

GetMultipartFiles parses and returns the post files array. Note that the request form should be type of multipart.

##### (*Request) GetMultipartForm 

``` go
func (r *Request) GetMultipartForm() *multipart.Form
```

GetMultipartForm parses and returns the form as multipart forms.

##### (*Request) GetNeverDoneCtx <-2.5.0

``` go
func (r *Request) GetNeverDoneCtx() context.Context
```

GetNeverDoneCtx creates and returns a never done context object, which forbids the context manually done, to make the context can be propagated to asynchronous goroutines, which will not be affected by the HTTP request ends.

This change is considered for common usage habits of developers for context propagation in multiple goroutines creation in one HTTP request.

##### (*Request) GetPage 

``` go
func (r *Request) GetPage(totalSize, pageSize int) *gpage.Page
```

GetPage creates and returns the pagination object for given `totalSize` and `pageSize`. NOTE THAT the page parameter name from clients is constantly defined as gpage.DefaultPageName for simplification and convenience.

##### (*Request) GetParam 

``` go
func (r *Request) GetParam(key string, def ...interface{}) *gvar.Var
```

GetParam returns custom parameter with a given name `key`. It returns `def` if `key` does not exist. It returns nil if `def` is not passed.

##### (*Request) GetQuery 

``` go
func (r *Request) GetQuery(key string, def ...interface{}) *gvar.Var
```

GetQuery retrieves and return parameter with the given name `key` from query string and request body. It returns `def` if `key` does not exist in the query and `def` is given, or else it returns nil.

Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote in order of priority: query > body.

##### (*Request) GetQueryMap 

``` go
func (r *Request) GetQueryMap(kvMap ...map[string]interface{}) map[string]interface{}
```

GetQueryMap retrieves and returns all parameters passed from the client using HTTP GET method as the map. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote in order of priority: query > body.

##### (*Request) GetQueryMapStrStr 

``` go
func (r *Request) GetQueryMapStrStr(kvMap ...map[string]interface{}) map[string]string
```

GetQueryMapStrStr retrieves and returns all parameters passed from the client using the HTTP GET method as a

```
map[string]string. The parameter `kvMap` specifies the keys
```

retrieving from client parameters, the associated values are the default values if the client does not pass.

##### (*Request) GetQueryMapStrVar 

``` go
func (r *Request) GetQueryMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var
```

GetQueryMapStrVar retrieves and returns all parameters passed from the client using the HTTP GET method as map[string]*gvar.Var. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

##### (*Request) GetQueryStruct 

``` go
func (r *Request) GetQueryStruct(pointer interface{}, mapping ...map[string]string) error
```

GetQueryStruct retrieves all parameters passed from the client using the HTTP GET method and converts them to a given struct object. Note that the parameter `pointer` is a pointer to the struct object. The optional parameter `mapping` is used to specify the key to attribute mapping.

##### (*Request) GetReferer 

``` go
func (r *Request) GetReferer() string
```

GetReferer returns referer of this request.

##### (*Request) GetRemoteIp 

``` go
func (r *Request) GetRemoteIp() string
```

GetRemoteIp returns the ip from RemoteAddr.

##### (*Request) GetRequest 

``` go
func (r *Request) GetRequest(key string, def ...interface{}) *gvar.Var
```

GetRequest retrieves and returns the parameter named `key` passed from the client and custom params as interface{}, no matter what HTTP method the client is using. The parameter `def` specifies the default value if the `key` does not exist.

GetRequest is one of the most commonly used functions for retrieving parameters.

Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote in order of priority: router < query < body < form < custom.

##### (*Request) GetRequestMap 

``` go
func (r *Request) GetRequestMap(kvMap ...map[string]interface{}) map[string]interface{}
```

GetRequestMap retrieves and returns all parameters passed from the client and custom params as the map, no matter what HTTP method the client is using. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass the according keys.

GetRequestMap is one of the most commonly used functions for retrieving parameters.

Note that if there are multiple parameters with the same name, the parameters are retrieved and overwrote in order of priority: router < query < body < form < custom.

##### (*Request) GetRequestMapStrStr 

``` go
func (r *Request) GetRequestMapStrStr(kvMap ...map[string]interface{}) map[string]string
```

GetRequestMapStrStr retrieve and returns all parameters passed from the client and custom params as map[string]string, no matter what HTTP method the client is using. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

##### (*Request) GetRequestMapStrVar 

``` go
func (r *Request) GetRequestMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var
```

GetRequestMapStrVar retrieve and returns all parameters passed from the client and custom params as map[string]*gvar.Var, no matter what HTTP method the client is using. The parameter `kvMap` specifies the keys retrieving from client parameters, the associated values are the default values if the client does not pass.

##### (*Request) GetRequestStruct 

``` go
func (r *Request) GetRequestStruct(pointer interface{}, mapping ...map[string]string) error
```

GetRequestStruct retrieves all parameters passed from the client and custom params no matter what HTTP method the client is using, and converts them to give the struct object. Note that the parameter `pointer` is a pointer to the struct object. The optional parameter `mapping` is used to specify the key to attribute mapping.

##### (*Request) GetRouter 

``` go
func (r *Request) GetRouter(key string, def ...interface{}) *gvar.Var
```

GetRouter retrieves and returns the router value with given key name `key`. It returns `def` if `key` does not exist.

##### (*Request) GetRouterMap 

``` go
func (r *Request) GetRouterMap() map[string]string
```

GetRouterMap retrieves and returns a copy of the router map.

##### (*Request) GetServeHandler <-2.2.2

``` go
func (r *Request) GetServeHandler() *HandlerItemParsed
```

GetServeHandler retrieves and returns the user defined handler used to serve this request.

##### (*Request) GetSessionId 

``` go
func (r *Request) GetSessionId() string
```

GetSessionId retrieves and returns session id from cookie or header.

##### (*Request) GetStruct 

``` go
func (r *Request) GetStruct(pointer interface{}, mapping ...map[string]string) error
```

GetStruct is an alias and convenient function for GetRequestStruct. See GetRequestStruct.

##### (*Request) GetUploadFile 

``` go
func (r *Request) GetUploadFile(name string) *UploadFile
```

GetUploadFile retrieves and returns the uploading file with specified form name. This function is used for retrieving single uploading file object, which is uploaded using multipart form content type.

It returns nil if retrieving failed or no form file with given name posted.

Note that the `name` is the file field name of the multipart form from client.

##### (*Request) GetUploadFiles 

``` go
func (r *Request) GetUploadFiles(name string) UploadFiles
```

GetUploadFiles retrieves and returns multiple uploading files with specified form name. This function is used for retrieving multiple uploading file objects, which are uploaded using multipart form content type.

It returns nil if retrieving failed or no form file with given name posted.

Note that the `name` is the file field name of the multipart form from client.

##### (*Request) GetUrl 

``` go
func (r *Request) GetUrl() string
```

GetUrl returns current URL of this request.

##### (*Request) GetView 

``` go
func (r *Request) GetView() *gview.View
```

GetView returns the template view engine object for this request.

##### (*Request) IsAjaxRequest 

``` go
func (r *Request) IsAjaxRequest() bool
```

IsAjaxRequest checks and returns whether current request is an AJAX request.

##### (*Request) IsExited 

``` go
func (r *Request) IsExited() bool
```

IsExited checks and returns whether current request is exited.

##### (*Request) IsFileRequest 

``` go
func (r *Request) IsFileRequest() bool
```

IsFileRequest checks and returns whether current request is serving file.

##### (*Request) MakeBodyRepeatableRead <-2.4.2

``` go
func (r *Request) MakeBodyRepeatableRead(repeatableRead bool) []byte
```

MakeBodyRepeatableRead marks the request body could be repeatedly readable or not. It also returns the current content of the request body.

##### (*Request) Parse 

``` go
func (r *Request) Parse(pointer interface{}) error
```

Parse is the most commonly used function, which converts request parameters to struct or struct slice. It also automatically validates the struct or every element of the struct slice according to the validation tag of the struct.

The parameter `pointer` can be type of: *struct/**struct/*[]struct/*[]*struct.

It supports single and multiple struct converting: 1. Single struct, post content like: {"id":1, "name":"john"} or ?id=1&name=john 2. Multiple struct, post content like: [{"id":1, "name":"john"}, {"id":, "name":"smith"}]

TODO: Improve the performance by reducing duplicated reflect usage on the same variable across packages.

##### (*Request) ParseForm 

``` go
func (r *Request) ParseForm(pointer interface{}) error
```

ParseForm performs like function Parse, but only parses the form parameters or the body content.

##### (*Request) ParseQuery 

``` go
func (r *Request) ParseQuery(pointer interface{}) error
```

ParseQuery performs like function Parse, but only parses the query parameters.

##### (*Request) ReloadParam 

``` go
func (r *Request) ReloadParam()
```

ReloadParam is used for modifying request parameter. Sometimes, we want to modify request parameters through middleware, but directly modifying Request.Body is invalid, so it clears the parsed* marks of Request to make the parameters reparsed.

##### (*Request) SetCtx 

``` go
func (r *Request) SetCtx(ctx context.Context)
```

SetCtx custom context for current request.

##### (*Request) SetCtxVar 

``` go
func (r *Request) SetCtxVar(key interface{}, value interface{})
```

SetCtxVar sets custom parameter to context with key-value pairs.

##### (*Request) SetError 

``` go
func (r *Request) SetError(err error)
```

SetError sets custom error for current request.

##### (*Request) SetForm 

``` go
func (r *Request) SetForm(key string, value interface{})
```

SetForm sets custom form value with key-value pairs.

##### (*Request) SetParam 

``` go
func (r *Request) SetParam(key string, value interface{})
```

SetParam sets custom parameter with key-value pairs.

##### (*Request) SetParamMap 

``` go
func (r *Request) SetParamMap(data map[string]interface{})
```

SetParamMap sets custom parameter with key-value pair maps.

##### (*Request) SetQuery 

``` go
func (r *Request) SetQuery(key string, value interface{})
```

SetQuery sets custom query value with key-value pairs.

##### (*Request) SetView 

``` go
func (r *Request) SetView(view *gview.View)
```

SetView sets template view engine object for this request.

##### (*Request) WebSocket 

``` go
func (r *Request) WebSocket() (*WebSocket, error)
```

WebSocket upgrades current request as a websocket request. It returns a new WebSocket object if success, or the error if failure. Note that the request should be a websocket request, or it will surely fail upgrading.

#### type Response 

``` go
type Response struct {
	*ResponseWriter                 // Underlying ResponseWriter.
	Server          *Server         // Parent server.
	Writer          *ResponseWriter // Alias of ResponseWriter.
	Request         *Request        // According request.
}
```

Response is the http response manager. Note that it implements the http.ResponseWriter interface with buffering feature.

##### (*Response) Buffer 

``` go
func (r *Response) Buffer() []byte
```

Buffer returns the buffered content as []byte.

##### (*Response) BufferLength 

``` go
func (r *Response) BufferLength() int
```

BufferLength returns the length of the buffered content.

##### (*Response) BufferString 

``` go
func (r *Response) BufferString() string
```

BufferString returns the buffered content as string.

##### (*Response) CORS 

``` go
func (r *Response) CORS(options CORSOptions)
```

CORS sets custom CORS options. See https://www.w3.org/TR/cors/ .

##### (*Response) CORSAllowedOrigin 

``` go
func (r *Response) CORSAllowedOrigin(options CORSOptions) bool
```

CORSAllowedOrigin CORSAllowed checks whether the current request origin is allowed cross-domain.

##### (*Response) CORSDefault 

``` go
func (r *Response) CORSDefault()
```

CORSDefault sets CORS with default CORS options, which allows any cross-domain request.

##### (*Response) ClearBuffer 

``` go
func (r *Response) ClearBuffer()
```

ClearBuffer clears the response buffer.

##### (*Response) DefaultCORSOptions 

``` go
func (r *Response) DefaultCORSOptions() CORSOptions
```

DefaultCORSOptions returns the default CORS options, which allows any cross-domain request.

##### (*Response) Flush 

``` go
func (r *Response) Flush()
```

Flush outputs the buffer content to the client and clears the buffer.

##### (*Response) ParseTpl 

``` go
func (r *Response) ParseTpl(tpl string, params ...gview.Params) (string, error)
```

ParseTpl parses given template file `tpl` with given template variables `params` and returns the parsed template content.

##### (*Response) ParseTplContent 

``` go
func (r *Response) ParseTplContent(content string, params ...gview.Params) (string, error)
```

ParseTplContent parses given template file `file` with given template parameters `params` and returns the parsed template content.

##### (*Response) ParseTplDefault 

``` go
func (r *Response) ParseTplDefault(params ...gview.Params) (string, error)
```

ParseTplDefault parses the default template file with params.

##### (*Response) RedirectBack 

``` go
func (r *Response) RedirectBack(code ...int)
```

RedirectBack redirects the client back to referer. The optional parameter `code` specifies the http status code for redirecting, which commonly can be 301 or 302. It's 302 in default.

##### (*Response) RedirectTo 

``` go
func (r *Response) RedirectTo(location string, code ...int)
```

RedirectTo redirects the client to another location. The optional parameter `code` specifies the http status code for redirecting, which commonly can be 301 or 302. It's 302 in default.

##### (*Response) ServeContent <-2.2.6

``` go
func (r *Response) ServeContent(name string, modTime time.Time, content io.ReadSeeker)
```

ServeContent replies to the request using the content in the provided ReadSeeker. The main benefit of ServeContent over io.Copy is that it handles Range requests properly, sets the MIME type, and handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since, and If-Range requests.

See http.ServeContent

##### (*Response) ServeFile 

``` go
func (r *Response) ServeFile(path string, allowIndex ...bool)
```

ServeFile serves the file to the response.

##### (*Response) ServeFileDownload 

``` go
func (r *Response) ServeFileDownload(path string, name ...string)
```

ServeFileDownload serves file downloading to the response.

##### (*Response) SetBuffer 

``` go
func (r *Response) SetBuffer(data []byte)
```

SetBuffer overwrites the buffer with `data`.

##### (*Response) Write 

``` go
func (r *Response) Write(content ...interface{})
```

Write writes `content` to the response buffer.

##### (*Response) WriteExit 

``` go
func (r *Response) WriteExit(content ...interface{})
```

WriteExit writes `content` to the response buffer and exits executing of current handler. The "Exit" feature is commonly used to replace usage of return statements in the handler, for convenience.

##### (*Response) WriteJson 

``` go
func (r *Response) WriteJson(content interface{})
```

WriteJson writes `content` to the response with JSON format.

##### (*Response) WriteJsonExit 

``` go
func (r *Response) WriteJsonExit(content interface{})
```

WriteJsonExit writes `content` to the response with JSON format and exits executing of current handler if success. The "Exit" feature is commonly used to replace usage of return statements in the handler, for convenience.

##### (*Response) WriteJsonP 

``` go
func (r *Response) WriteJsonP(content interface{})
```

WriteJsonP writes `content` to the response with JSONP format.

Note that there should be a "callback" parameter in the request for JSONP format.

##### (*Response) WriteJsonPExit 

``` go
func (r *Response) WriteJsonPExit(content interface{})
```

WriteJsonPExit writes `content` to the response with JSONP format and exits executing of current handler if success. The "Exit" feature is commonly used to replace usage of return statements in the handler, for convenience.

Note that there should be a "callback" parameter in the request for JSONP format.

##### (*Response) WriteOver 

``` go
func (r *Response) WriteOver(content ...interface{})
```

WriteOver overwrites the response buffer with `content`.

##### (*Response) WriteOverExit 

``` go
func (r *Response) WriteOverExit(content ...interface{})
```

WriteOverExit overwrites the response buffer with `content` and exits executing of current handler. The "Exit" feature is commonly used to replace usage of return statements in the handler, for convenience.

##### (*Response) WriteStatus 

``` go
func (r *Response) WriteStatus(status int, content ...interface{})
```

WriteStatus writes HTTP `status` and `content` to the response. Note that it does not set a Content-Type header here.

##### (*Response) WriteStatusExit 

``` go
func (r *Response) WriteStatusExit(status int, content ...interface{})
```

WriteStatusExit writes HTTP `status` and `content` to the response and exits executing of current handler if success. The "Exit" feature is commonly used to replace usage of return statements in the handler, for convenience.

##### (*Response) WriteTpl 

``` go
func (r *Response) WriteTpl(tpl string, params ...gview.Params) error
```

WriteTpl parses and responses given template file. The parameter `params` specifies the template variables for parsing.

##### (*Response) WriteTplContent 

``` go
func (r *Response) WriteTplContent(content string, params ...gview.Params) error
```

WriteTplContent parses and responses the template content. The parameter `params` specifies the template variables for parsing.

##### (*Response) WriteTplDefault 

``` go
func (r *Response) WriteTplDefault(params ...gview.Params) error
```

WriteTplDefault parses and responses the default template file. The parameter `params` specifies the template variables for parsing.

##### (*Response) WriteXml 

``` go
func (r *Response) WriteXml(content interface{}, rootTag ...string)
```

WriteXml writes `content` to the response with XML format.

##### (*Response) WriteXmlExit 

``` go
func (r *Response) WriteXmlExit(content interface{}, rootTag ...string)
```

WriteXmlExit writes `content` to the response with XML format and exits executing of current handler if success. The "Exit" feature is commonly used to replace usage of return statements in the handler, for convenience.

##### (*Response) Writef 

``` go
func (r *Response) Writef(format string, params ...interface{})
```

Writef writes the response with fmt.Sprintf.

##### (*Response) WritefExit 

``` go
func (r *Response) WritefExit(format string, params ...interface{})
```

WritefExit writes the response with fmt.Sprintf and exits executing of current handler. The "Exit" feature is commonly used to replace usage of return statements in the handler, for convenience.

##### (*Response) Writefln 

``` go
func (r *Response) Writefln(format string, params ...interface{})
```

Writefln writes the response with fmt.Sprintf and new line.

##### (*Response) WriteflnExit 

``` go
func (r *Response) WriteflnExit(format string, params ...interface{})
```

WriteflnExit writes the response with fmt.Sprintf and new line and exits executing of current handler. The "Exit" feature is commonly used to replace usage of return statement in the handler, for convenience.

##### (*Response) Writeln 

``` go
func (r *Response) Writeln(content ...interface{})
```

Writeln writes the response with `content` and new line.

##### (*Response) WritelnExit 

``` go
func (r *Response) WritelnExit(content ...interface{})
```

WritelnExit writes the response with `content` and new line and exits executing of current handler. The "Exit" feature is commonly used to replace usage of return statements in the handler, for convenience.

#### type ResponseWriter 

``` go
type ResponseWriter struct {
	Status int // HTTP status.
	// contains filtered or unexported fields
}
```

ResponseWriter is the custom writer for http response.

##### (*ResponseWriter) Flush 

``` go
func (w *ResponseWriter) Flush()
```

Flush outputs the buffer to clients and clears the buffer.

##### (*ResponseWriter) Header 

``` go
func (w *ResponseWriter) Header() http.Header
```

Header implements the interface function of http.ResponseWriter.Header.

##### (*ResponseWriter) Hijack 

``` go
func (w *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error)
```

Hijack implements the interface function of http.Hijacker.Hijack.

##### (*ResponseWriter) RawWriter 

``` go
func (w *ResponseWriter) RawWriter() http.ResponseWriter
```

RawWriter returns the underlying ResponseWriter.

##### (*ResponseWriter) Write 

``` go
func (w *ResponseWriter) Write(data []byte) (int, error)
```

Write implements the interface function of http.ResponseWriter.Write.

##### (*ResponseWriter) WriteHeader 

``` go
func (w *ResponseWriter) WriteHeader(status int)
```

WriteHeader implements the interface of http.ResponseWriter.WriteHeader.

#### type Router 

``` go
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

#### type RouterGroup 

``` go
type RouterGroup struct {
	// contains filtered or unexported fields
}
```

RouterGroup is a group wrapping multiple routes and middleware.

##### (*RouterGroup) ALL 

``` go
func (g *RouterGroup) ALL(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

ALL register an http handler to give the route pattern and all http methods.

##### (*RouterGroup) ALLMap 

``` go
func (g *RouterGroup) ALLMap(m map[string]interface{})
```

ALLMap registers http handlers for http methods using map.

##### (*RouterGroup) Bind 

``` go
func (g *RouterGroup) Bind(handlerOrObject ...interface{}) *RouterGroup
```

Bind does batch route registering feature for a router group.

##### (*RouterGroup) CONNECT 

``` go
func (g *RouterGroup) CONNECT(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

CONNECT registers an http handler to give the route pattern and the http method: CONNECT.

##### (*RouterGroup) Clone 

``` go
func (g *RouterGroup) Clone() *RouterGroup
```

Clone returns a new router group which is a clone of the current group.

##### (*RouterGroup) DELETE 

``` go
func (g *RouterGroup) DELETE(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

DELETE registers an http handler to give the route pattern and the http method: DELETE.

##### (*RouterGroup) GET 

``` go
func (g *RouterGroup) GET(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

GET registers an http handler to give the route pattern and the http method: GET.

##### (*RouterGroup) Group 

``` go
func (g *RouterGroup) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup
```

Group creates and returns a subgroup of the current router group.

##### (*RouterGroup) HEAD 

``` go
func (g *RouterGroup) HEAD(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

HEAD registers an http handler to give the route pattern and the http method: HEAD.

##### (*RouterGroup) Hook 

``` go
func (g *RouterGroup) Hook(pattern string, hook HookName, handler HandlerFunc) *RouterGroup
```

Hook registers a hook to given route pattern.

##### (*RouterGroup) Map 

``` go
func (g *RouterGroup) Map(m map[string]interface{})
```

Map registers http handlers for http methods using map.

##### (*RouterGroup) Middleware 

``` go
func (g *RouterGroup) Middleware(handlers ...HandlerFunc) *RouterGroup
```

Middleware binds one or more middleware to the router group.

##### (*RouterGroup) OPTIONS 

``` go
func (g *RouterGroup) OPTIONS(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

OPTIONS register an http handler to give the route pattern and the http method: OPTIONS.

##### (*RouterGroup) PATCH 

``` go
func (g *RouterGroup) PATCH(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

PATCH registers an http handler to give the route pattern and the http method: PATCH.

##### (*RouterGroup) POST 

``` go
func (g *RouterGroup) POST(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

POST registers an http handler to give the route pattern and the http method: POST.

##### (*RouterGroup) PUT 

``` go
func (g *RouterGroup) PUT(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

PUT registers an http handler to give the route pattern and the http method: PUT.

##### (*RouterGroup) REST 

``` go
func (g *RouterGroup) REST(pattern string, object interface{}) *RouterGroup
```

REST registers an http handler to give the route pattern according to REST rule.

##### (*RouterGroup) TRACE 

``` go
func (g *RouterGroup) TRACE(pattern string, object interface{}, params ...interface{}) *RouterGroup
```

TRACE registers an http handler to give the route pattern and the http method: TRACE.

#### type RouterItem 

``` go
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

#### type Server 

``` go
type Server struct {
	// contains filtered or unexported fields
}
```

Server wraps the http.Server and provides more rich features.

##### func GetServer 

``` go
func GetServer(name ...interface{}) *Server
```

GetServer creates and returns a server instance using given name and default configurations. Note that the parameter `name` should be unique for different servers. It returns an existing server instance if given `name` is already existing in the server mapping.

##### (*Server) AddSearchPath 

``` go
func (s *Server) AddSearchPath(path string)
```

AddSearchPath add searching directory path for static file service.

##### (*Server) AddStaticPath 

``` go
func (s *Server) AddStaticPath(prefix string, path string)
```

AddStaticPath sets the uri to static directory path mapping for static file service.

##### (*Server) BindHandler 

``` go
func (s *Server) BindHandler(pattern string, handler interface{})
```

BindHandler registers a handler function to server with a given pattern.

Note that the parameter `handler` can be type of: 1. func(*ghttp.Request) 2. func(context.Context, BizRequest)(BizResponse, error)

##### (*Server) BindHookHandler 

``` go
func (s *Server) BindHookHandler(pattern string, hook HookName, handler HandlerFunc)
```

BindHookHandler registers handler for specified hook.

##### (*Server) BindHookHandlerByMap 

``` go
func (s *Server) BindHookHandlerByMap(pattern string, hookMap map[HookName]HandlerFunc)
```

BindHookHandlerByMap registers handler for specified hook.

##### (*Server) BindMiddleware 

``` go
func (s *Server) BindMiddleware(pattern string, handlers ...HandlerFunc)
```

BindMiddleware registers one or more global middleware to the server. Global middleware can be used standalone without service handler, which intercepts all dynamic requests before or after service handler. The parameter `pattern` specifies what route pattern the middleware intercepts, which is usually a "fuzzy" pattern like "/:name", "/*any" or "/{field}".

##### (*Server) BindMiddlewareDefault 

``` go
func (s *Server) BindMiddlewareDefault(handlers ...HandlerFunc)
```

BindMiddlewareDefault registers one or more global middleware to the server using default pattern "/*". Global middleware can be used standalone without service handler, which intercepts all dynamic requests before or after service handler.

##### (*Server) BindObject 

``` go
func (s *Server) BindObject(pattern string, object interface{}, method ...string)
```

BindObject registers object to server routes with a given pattern.

The optional parameter `method` is used to specify the method to be registered, which supports multiple method names; multiple methods are separated by char ',', case-sensitive.

##### (*Server) BindObjectMethod 

``` go
func (s *Server) BindObjectMethod(pattern string, object interface{}, method string)
```

BindObjectMethod registers specified method of the object to server routes with a given pattern.

The optional parameter `method` is used to specify the method to be registered, which does not support multiple method names but only one, case-sensitive.

##### (*Server) BindObjectRest 

``` go
func (s *Server) BindObjectRest(pattern string, object interface{})
```

BindObjectRest registers object in REST API styles to server with a specified pattern.

##### (*Server) BindStatusHandler 

``` go
func (s *Server) BindStatusHandler(status int, handler HandlerFunc)
```

BindStatusHandler registers handler for given status code.

##### (*Server) BindStatusHandlerByMap 

``` go
func (s *Server) BindStatusHandlerByMap(handlerMap map[int]HandlerFunc)
```

BindStatusHandlerByMap registers handler for given status code using map.

##### (*Server) Domain 

``` go
func (s *Server) Domain(domains string) *Domain
```

Domain creates and returns a domain object for management for one or more domains.

##### (*Server) EnableAdmin 

``` go
func (s *Server) EnableAdmin(pattern ...string)
```

EnableAdmin enables the administration feature for the process. The optional parameter `pattern` specifies the URI for the administration page.

##### (*Server) EnableHTTPS 

``` go
func (s *Server) EnableHTTPS(certFile, keyFile string, tlsConfig ...*tls.Config)
```

EnableHTTPS enables HTTPS with given certification and key files for the server. The optional parameter `tlsConfig` specifies custom TLS configuration.

##### (*Server) EnablePProf 

``` go
func (s *Server) EnablePProf(pattern ...string)
```

EnablePProf enables PProf feature for server.

##### (*Server) GetCookieDomain 

``` go
func (s *Server) GetCookieDomain() string
```

GetCookieDomain returns CookieDomain of server.

##### (*Server) GetCookieHttpOnly 

``` go
func (s *Server) GetCookieHttpOnly() bool
```

##### (*Server) GetCookieMaxAge 

``` go
func (s *Server) GetCookieMaxAge() time.Duration
```

GetCookieMaxAge returns the CookieMaxAge of the server.

##### (*Server) GetCookiePath 

``` go
func (s *Server) GetCookiePath() string
```

GetCookiePath returns the CookiePath of server.

##### (*Server) GetCookieSameSite 

``` go
func (s *Server) GetCookieSameSite() http.SameSite
```

GetCookieSameSite return CookieSameSite of server.

##### (*Server) GetCookieSecure 

``` go
func (s *Server) GetCookieSecure() bool
```

##### (*Server) GetHandler 

``` go
func (s *Server) GetHandler() func(w http.ResponseWriter, r *http.Request)
```

GetHandler returns the request handler of the server.

##### (*Server) GetIndexFiles 

``` go
func (s *Server) GetIndexFiles() []string
```

GetIndexFiles retrieves and returns the index files from the server.

##### (*Server) GetListenedAddress <-2.2.0

``` go
func (s *Server) GetListenedAddress() string
```

GetListenedAddress retrieves and returns the address string which are listened by current server.

##### (*Server) GetListenedPort 

``` go
func (s *Server) GetListenedPort() int
```

GetListenedPort retrieves and returns one port which is listened by current server.

##### (*Server) GetListenedPorts 

``` go
func (s *Server) GetListenedPorts() []int
```

GetListenedPorts retrieves and returns the ports which are listened by current server.

##### (*Server) GetLogPath 

``` go
func (s *Server) GetLogPath() string
```

GetLogPath returns the log path.

##### (*Server) GetName 

``` go
func (s *Server) GetName() string
```

GetName returns the name of the server.

##### (*Server) GetOpenApi 

``` go
func (s *Server) GetOpenApi() *goai.OpenApiV3
```

GetOpenApi returns the OpenApi specification management object of current server.

##### (*Server) GetRegistrar <-2.3.3

``` go
func (s *Server) GetRegistrar() gsvc.Registrar
```

GetRegistrar returns the Registrar of server.

##### (*Server) GetRoutes 

``` go
func (s *Server) GetRoutes() []RouterItem
```

GetRoutes retrieves and returns the router array.

##### (*Server) GetSessionCookieMaxAge 

``` go
func (s *Server) GetSessionCookieMaxAge() time.Duration
```

GetSessionCookieMaxAge returns the SessionCookieMaxAge of server.

##### (*Server) GetSessionIdName 

``` go
func (s *Server) GetSessionIdName() string
```

GetSessionIdName returns the SessionIdName of server.

##### (*Server) GetSessionMaxAge 

``` go
func (s *Server) GetSessionMaxAge() time.Duration
```

GetSessionMaxAge returns the SessionMaxAge of server.

##### (*Server) Group 

``` go
func (s *Server) Group(prefix string, groups ...func(group *RouterGroup)) *RouterGroup
```

Group creates and returns a RouterGroup object.

##### (*Server) IsAccessLogEnabled 

``` go
func (s *Server) IsAccessLogEnabled() bool
```

IsAccessLogEnabled checks whether the access log enabled.

##### (*Server) IsErrorLogEnabled 

``` go
func (s *Server) IsErrorLogEnabled() bool
```

IsErrorLogEnabled checks whether the error log enabled.

##### (*Server) Logger 

``` go
func (s *Server) Logger() *glog.Logger
```

Logger is alias of GetLogger.

##### (*Server) Plugin 

``` go
func (s *Server) Plugin(plugin ...Plugin)
```

Plugin adds plugin to the server.

##### (*Server) Run 

``` go
func (s *Server) Run()
```

Run starts server listening in blocking way. It's commonly used for single server situation.

##### Example

``` go
```
##### (*Server) ServeHTTP 

``` go
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

ServeHTTP is the default handler for http request. It should not create new goroutine handling the request as it's called by am already created new goroutine from http.Server.

This function also makes serve implementing the interface of http.Handler.

##### (*Server) SetAccessLogEnabled 

``` go
func (s *Server) SetAccessLogEnabled(enabled bool)
```

SetAccessLogEnabled enables/disables the access log.

##### (*Server) SetAddr 

``` go
func (s *Server) SetAddr(address string)
```

SetAddr sets the listening address for the server. The address is like ':80', '0.0.0.0:80', '127.0.0.1:80', '180.18.99.10:80', etc.

##### (*Server) SetClientMaxBodySize 

``` go
func (s *Server) SetClientMaxBodySize(maxSize int64)
```

SetClientMaxBodySize sets the ClientMaxBodySize for server.

##### (*Server) SetConfig 

``` go
func (s *Server) SetConfig(c ServerConfig) error
```

SetConfig sets the configuration for the server.

##### (*Server) SetConfigWithMap 

``` go
func (s *Server) SetConfigWithMap(m map[string]interface{}) error
```

SetConfigWithMap sets the configuration for the server using map.

##### (*Server) SetCookieDomain 

``` go
func (s *Server) SetCookieDomain(domain string)
```

SetCookieDomain sets the CookieDomain for server.

##### (*Server) SetCookieMaxAge 

``` go
func (s *Server) SetCookieMaxAge(ttl time.Duration)
```

SetCookieMaxAge sets the CookieMaxAge for server.

##### (*Server) SetCookiePath 

``` go
func (s *Server) SetCookiePath(path string)
```

SetCookiePath sets the CookiePath for server.

##### (*Server) SetDumpRouterMap 

``` go
func (s *Server) SetDumpRouterMap(enabled bool)
```

SetDumpRouterMap sets the DumpRouterMap for server. If DumpRouterMap is enabled, it automatically dumps the route map when server starts.

##### (*Server) SetEndpoints <-2.5.0

``` go
func (s *Server) SetEndpoints(endpoints []string)
```

SetEndpoints sets the Endpoints for the server.

##### (*Server) SetErrorLogEnabled 

``` go
func (s *Server) SetErrorLogEnabled(enabled bool)
```

SetErrorLogEnabled enables/disables the error log.

##### (*Server) SetErrorStack 

``` go
func (s *Server) SetErrorStack(enabled bool)
```

SetErrorStack enables/disables the error stack feature.

##### (*Server) SetFileServerEnabled 

``` go
func (s *Server) SetFileServerEnabled(enabled bool)
```

SetFileServerEnabled enables/disables the static file service. It's the main switch for the static file service. When static file service configuration functions like SetServerRoot, AddSearchPath and AddStaticPath are called, this configuration is automatically enabled.

##### (*Server) SetFormParsingMemory 

``` go
func (s *Server) SetFormParsingMemory(maxMemory int64)
```

SetFormParsingMemory sets the FormParsingMemory for server.

##### (*Server) SetHTTPSAddr 

``` go
func (s *Server) SetHTTPSAddr(address string)
```

SetHTTPSAddr sets the HTTPS listening ports for the server.

##### (*Server) SetHTTPSPort 

``` go
func (s *Server) SetHTTPSPort(port ...int)
```

SetHTTPSPort sets the HTTPS listening ports for the server. The listening ports can be multiple like: SetHTTPSPort(443, 500).

##### (*Server) SetHandler 

``` go
func (s *Server) SetHandler(h func(w http.ResponseWriter, r *http.Request))
```

SetHandler sets the request handler for server.

##### (*Server) SetIdleTimeout 

``` go
func (s *Server) SetIdleTimeout(t time.Duration)
```

SetIdleTimeout sets the IdleTimeout for the server.

##### (*Server) SetIndexFiles 

``` go
func (s *Server) SetIndexFiles(indexFiles []string)
```

SetIndexFiles sets the index files for server.

##### (*Server) SetIndexFolder 

``` go
func (s *Server) SetIndexFolder(enabled bool)
```

SetIndexFolder enables/disables listing the sub-files if requesting a directory.

##### (*Server) SetKeepAlive 

``` go
func (s *Server) SetKeepAlive(enabled bool)
```

SetKeepAlive sets the KeepAlive for the server.

##### (*Server) SetListener <-2.1.0

``` go
func (s *Server) SetListener(listeners ...net.Listener) error
```

SetListener set the custom listener for the server.

##### (*Server) SetLogLevel 

``` go
func (s *Server) SetLogLevel(level string)
```

SetLogLevel sets logging level by level string.

##### (*Server) SetLogPath 

``` go
func (s *Server) SetLogPath(path string) error
```

SetLogPath sets the log path for server. It logs content to file only if the log path is set.

##### (*Server) SetLogStdout 

``` go
func (s *Server) SetLogStdout(enabled bool)
```

SetLogStdout sets whether output the logging content to stdout.

##### (*Server) SetLogger 

``` go
func (s *Server) SetLogger(logger *glog.Logger)
```

SetLogger sets the logger for logging responsibility. Note that it cannot be set in runtime as there may be concurrent safety issue.

##### (*Server) SetMaxHeaderBytes 

``` go
func (s *Server) SetMaxHeaderBytes(b int)
```

SetMaxHeaderBytes sets the MaxHeaderBytes for the server.

##### (*Server) SetName 

``` go
func (s *Server) SetName(name string)
```

SetName sets the name for the server.

##### (*Server) SetNameToUriType 

``` go
func (s *Server) SetNameToUriType(t int)
```

SetNameToUriType sets the NameToUriType for server.

##### (*Server) SetOpenApiPath 

``` go
func (s *Server) SetOpenApiPath(path string)
```

SetOpenApiPath sets the OpenApiPath for server.

##### (*Server) SetPort 

``` go
func (s *Server) SetPort(port ...int)
```

SetPort sets the listening ports for the server. The listening ports can be multiple like: SetPort(80, 8080).

##### (*Server) SetReadTimeout 

``` go
func (s *Server) SetReadTimeout(t time.Duration)
```

SetReadTimeout sets the ReadTimeout for the server.

##### (*Server) SetRegistrar <-2.3.3

``` go
func (s *Server) SetRegistrar(registrar gsvc.Registrar)
```

SetRegistrar sets the Registrar for server.

##### (*Server) SetRewrite 

``` go
func (s *Server) SetRewrite(uri string, rewrite string)
```

SetRewrite sets rewrites for static URI for server.

##### (*Server) SetRewriteMap 

``` go
func (s *Server) SetRewriteMap(rewrites map[string]string)
```

SetRewriteMap sets the rewritten map for server.

##### (*Server) SetRouteOverWrite 

``` go
func (s *Server) SetRouteOverWrite(enabled bool)
```

SetRouteOverWrite sets the RouteOverWrite for server.

##### (*Server) SetServerAgent 

``` go
func (s *Server) SetServerAgent(agent string)
```

SetServerAgent sets the ServerAgent for the server.

##### (*Server) SetServerRoot 

``` go
func (s *Server) SetServerRoot(root string)
```

SetServerRoot sets the document root for static service.

##### (*Server) SetSessionCookieMaxAge 

``` go
func (s *Server) SetSessionCookieMaxAge(maxAge time.Duration)
```

SetSessionCookieMaxAge sets the SessionCookieMaxAge for server.

##### (*Server) SetSessionCookieOutput 

``` go
func (s *Server) SetSessionCookieOutput(enabled bool)
```

SetSessionCookieOutput sets the SetSessionCookieOutput for server.

##### (*Server) SetSessionIdName 

``` go
func (s *Server) SetSessionIdName(name string)
```

SetSessionIdName sets the SessionIdName for server.

##### (*Server) SetSessionMaxAge 

``` go
func (s *Server) SetSessionMaxAge(ttl time.Duration)
```

SetSessionMaxAge sets the SessionMaxAge for server.

##### (*Server) SetSessionStorage 

``` go
func (s *Server) SetSessionStorage(storage gsession.Storage)
```

SetSessionStorage sets the SessionStorage for server.

##### (*Server) SetSwaggerPath 

``` go
func (s *Server) SetSwaggerPath(path string)
```

SetSwaggerPath sets the SwaggerPath for server.

##### (*Server) SetSwaggerUITemplate <-2.6.2

``` go
func (s *Server) SetSwaggerUITemplate(swaggerUITemplate string)
```

SetSwaggerUITemplate sets the Swagger template for server.

##### (*Server) SetTLSConfig 

``` go
func (s *Server) SetTLSConfig(tlsConfig *tls.Config)
```

SetTLSConfig sets custom TLS configuration and enables HTTPS feature for the server.

##### (*Server) SetView 

``` go
func (s *Server) SetView(view *gview.View)
```

SetView sets the View for the server.

##### (*Server) SetWriteTimeout 

``` go
func (s *Server) SetWriteTimeout(t time.Duration)
```

SetWriteTimeout sets the WriteTimeout for the server.

##### (*Server) Shutdown 

``` go
func (s *Server) Shutdown() error
```

Shutdown shuts down current server.

##### (*Server) Start 

``` go
func (s *Server) Start() error
```

Start starts listening on configured port. This function does not block the process, you can use function Wait blocking the process.

##### (*Server) Status 

``` go
func (s *Server) Status() ServerStatus
```

Status retrieves and returns the server status.

##### (*Server) Use 

``` go
func (s *Server) Use(handlers ...HandlerFunc)
```

Use is the alias of BindMiddlewareDefault. See BindMiddlewareDefault.

#### type ServerConfig 

``` go
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

##### func ConfigFromMap 

``` go
func ConfigFromMap(m map[string]interface{}) (ServerConfig, error)
```

ConfigFromMap creates and returns a ServerConfig object with given map and default configuration object.

##### func NewConfig 

``` go
func NewConfig() ServerConfig
```

NewConfig creates and returns a ServerConfig object with default configurations. Note that, do not define this default configuration to local package variable, as there are some pointer attributes that may be shared in different servers.

#### type ServerStatus <-2.5.0

``` go
type ServerStatus = int
```

ServerStatus is the server status enum type.

#### type Session 

``` go
type Session = gsession.Session
```

Session is actually an alias of gsession.Session, which is bound to a single request.

#### type UploadFile 

``` go
type UploadFile struct {
	*multipart.FileHeader `json:"-"`
	// contains filtered or unexported fields
}
```

UploadFile wraps the multipart uploading file with more and convenient features.

##### (UploadFile) MarshalJSON <-2.1.0

``` go
func (f UploadFile) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*UploadFile) Save 

``` go
func (f *UploadFile) Save(dirPath string, randomlyRename ...bool) (filename string, err error)
```

Save saves the single uploading file to directory path and returns the saved file name.

The parameter `dirPath` should be a directory path, or it returns error.

Note that it will OVERWRITE the target file if there's already a same name file exist.

##### Example

``` go
```
#### type UploadFiles 

``` go
type UploadFiles []*UploadFile
```

UploadFiles is an array type of *UploadFile.

##### (UploadFiles) Save 

``` go
func (fs UploadFiles) Save(dirPath string, randomlyRename ...bool) (filenames []string, err error)
```

Save saves all uploading files to specified directory path and returns the saved file names.

The parameter `dirPath` should be a directory path or it returns error.

The parameter `randomlyRename` specifies whether randomly renames all the file names.

#### type WebSocket 

``` go
type WebSocket struct {
	*websocket.Conn
}
```

WebSocket wraps the underlying websocket connection and provides convenient functions.