+++
title = "http/cgi"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
https://pkg.go.dev/net/http/cgi@go1.20.1

Package cgi implements CGI (Common Gateway Interface) as specified in [RFC 3875](https://rfc-editor.org/rfc/rfc3875.html).

​	 cgi 包实现了 CGI（通用网关接口），如 [RFC 3875](https://rfc-editor.org/rfc/rfc3875.html) 中所规定。

Note that using CGI means starting a new process to handle each request, which is typically less efficient than using a long-running server. This package is intended primarily for compatibility with existing systems.

​	请注意，使用 CGI 意味着每个请求都会启动一个新的进程来处理，这通常比使用长时间运行的服务器效率要低。此包主要用于与现有系统的兼容性。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Request 

``` go 
func Request() (*http.Request, error)
```

Request returns the HTTP request as represented in the current environment. This assumes the current program is being run by a web server in a CGI environment. The returned Request's Body is populated, if applicable.

​	Request 函数返回当前环境中表示的 HTTP 请求。这假设当前程序在 CGI 环境中由 Web 服务器运行。如果适用，返回的 Request 的 Body 字段将被填充。

### func RequestFromMap 

``` go 
func RequestFromMap(params map[string]string) (*http.Request, error)
```

RequestFromMap creates an http.Request from CGI variables. The returned Request's Body field is not populated.

​	RequestFromMap 函数根据 CGI 变量创建一个 http.Request。返回的 Request 的 Body 字段不会被填充。

### func Serve 

``` go 
func Serve(handler http.Handler) error
```

Serve executes the provided Handler on the currently active CGI request, if any. If there's no current CGI environment an error is returned. The provided handler may be nil to use http.DefaultServeMux.

​	Serve 函数在当前活动的 CGI 请求上执行提供的 Handler。如果没有当前的 CGI 环境，则返回错误。提供的 handler 可以为 nil，以使用 http.DefaultServeMux。

## 类型

### type Handler 

``` go 
type Handler struct {
	Path string // 可执行文件的路径 path to the CGI executable
	Root string // 处理程序的根 URI 前缀，如果为空，则为 "/" root URI prefix of handler or empty for "/"

    // Dir specifies the CGI executable's working directory.
	// If Dir is empty, the base directory of Path is used.
	// If Path has no base directory, the current working
	// directory is used.
    // Dir 指定 CGI 可执行文件的工作目录。
	// 如果 Dir 为空，则使用 Path 的基目录。
	// 如果 Path 没有基目录，则使用当前工作目录。
	Dir string

	Env        []string    // 要设置的额外环境变量，如果有的话，格式为 "key=value" extra environment variables to set, if any, as "key=value"
	InheritEnv []string    // 从主机继承的环境变量，格式为 "key" environment variables to inherit from host, as "key"
	Logger     *log.Logger // 用于错误的可选日志，为 nil 则使用 log.Print optional log for errors or nil to use log.Print
	Args       []string    // 要传递给子进程的可选参数 optional arguments to pass to child process
	Stderr     io.Writer   // 子进程的可选 stderr；nil 表示 os.Stderr optional stderr for the child process; nil means os.Stderr

    // PathLocationHandler specifies the root http Handler that
	// should handle internal redirects when the CGI process
	// returns a Location header value starting with a "/", as
	// specified in RFC 3875 § 6.3.2. This will likely be
	// http.DefaultServeMux.
	//
	// If nil, a CGI response with a local URI path is instead sent
	// back to the client and not redirected internally.
   	// PathLocationHandler 指定在 CGI 进程返回以 "/" 开头的 Location 标头值时，
    // 应处理内部重定向的根 http 处理程序，
    // 如 RFC 3875 § 6.3.2 中所指定。
	// 这可能是 http.DefaultServeMux。
	//
	// 如果为 nil，则会将带有本地 URI 路径的 CGI 响应直接发送回客户端，而不进行内部重定向。
	PathLocationHandler http.Handler
}
```

Handler runs an executable in a subprocess with a CGI environment.

​	Handler 在子进程中使用 CGI 环境运行可执行文件。

#### (*Handler) ServeHTTP 

``` go 
func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request)
```

