+++
title = "gin文档"
date = 2023-06-05T08:55:39+08:00
type = "docs"
weight = 1
description = ""
isCJKLanguage = true
draft = false
+++

# Gin Web Framework

https://pkg.go.dev/github.com/gin-gonic/gin

当前版本：1.9.0 发布时间：2023.2.21

​	Gin是一个用Go编写的网络框架。它的特点是有一个类似于martini-like的API，由于[httprouter](https://github.com/julienschmidt/httprouter)的存在，其性能可提高40倍。如果你需要性能和良好的生产力，你会喜欢Gin。

Gin的主要特点：

- 零分配路由器（Zero allocation router）：指的是在处理 HTTP 请求时，不需要动态地分配内存来创建路由，从而提高了性能和减少了资源占用。
- 快速
- 支持中间件
- 不会出现崩溃（Crash-free）
- JSON 验证
- 路由组
- 错误管理
- 内置渲染
- 可扩展性

## 开始入门

### 前提条件

- **[Go](https://go.dev/)**: ~~any one of the **three latest major** [releases](https://go.dev/doc/devel/release)~~ (现在需要1.16以上版本).

### 获得Gin

有了[Go module](https://github.com/golang/go/wiki/Modules)的支持，只需添加以下导入即可

```go
import "github.com/gin-gonic/gin"
```

到你的代码中，然后`go [build|run|test]`会自动获取必要的依赖项。

否则，运行下面的Go命令来安装gin包：

```sh
$ go get -u github.com/gin-gonic/gin
```

!!! warning "注意"
	从go1.17版本开始，已经弃用这种安装方式，详见：[Go1.17发布说明](../../../ReleaseNotes/Go1_17ReleaseNotes#go-get)。这种方式只能在存在 `go.mod`文件的情况下，更改`go.mod`中的依赖项。


​	

### 运行Gin

​	首先你需要导入Gin包来使用Gin，一个最简单的例子就是下面这个`example.go`。

```go linenums="1"
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

然后使用go命令来运行该示例：

```sh
# run example.go and visit 0.0.0.0:8080/ping on browser
$ go run example.go
```

### 了解更多的例子

#### 快速入门

​	学习和练习更多的例子，请阅读[Gin Quick Start](https://github.com/gin-gonic/gin/blob/v1.9.0/docs/doc.md)，其中包括API使用示例和构建标签。

#### 示例

​	在[Gin示例库](https://github.com/gin-gonic/examples)中有许多可运行的实例，演示了Gin的各种使用情况。

### 文档

​	参见[API文档和包](https://pkg.go.dev/github.com/gin-gonic/gin?utm_source=godoc)的描述。

​	所有的文档都可以在Gin网站上找到。

- [English](https://gin-gonic.com/docs/)
- [简体中文](https://gin-gonic.com/zh-cn/docs/)
- [繁體中文](https://gin-gonic.com/zh-tw/docs/)
- [日本語](https://gin-gonic.com/ja/docs/)
- [Español](https://gin-gonic.com/es/docs/)
- [한국어](https://gin-gonic.com/ko-kr/docs/)
- [Turkish](https://gin-gonic.com/tr/docs/)
- [Persian](https://gin-gonic.com/fa/docs/)

#### 关于Gin的文章

一个精心策划的关于Gin框架的列表。

- [教程：用Go和Gin开发一个RESTful API](../../../GettingStarted/TutorialDevelopingARESTfulAPIWithGoAndGin)

### 基准测试

​	Gin使用自定义版本的[HttpRouter](https://github.com/julienschmidt/httprouter)，请看[所有的基准测试细节](https://github.com/gin-gonic/gin/blob/v1.9.0/BENCHMARKS.md)。

| Benchmark name                 | (1)       | (2)             | (3)          | (4)             |
| ------------------------------ | --------- | --------------- | ------------ | --------------- |
| BenchmarkGin_GithubAll         | **43550** | **27364 ns/op** | **0 B/op**   | **0 allocs/op** |
| BenchmarkAce_GithubAll         | 40543     | 29670 ns/op     | 0 B/op       | 0 allocs/op     |
| BenchmarkAero_GithubAll        | 57632     | 20648 ns/op     | 0 B/op       | 0 allocs/op     |
| BenchmarkBear_GithubAll        | 9234      | 216179 ns/op    | 86448 B/op   | 943 allocs/op   |
| BenchmarkBeego_GithubAll       | 7407      | 243496 ns/op    | 71456 B/op   | 609 allocs/op   |
| BenchmarkBone_GithubAll        | 420       | 2922835 ns/op   | 720160 B/op  | 8620 allocs/op  |
| BenchmarkChi_GithubAll         | 7620      | 238331 ns/op    | 87696 B/op   | 609 allocs/op   |
| BenchmarkDenco_GithubAll       | 18355     | 64494 ns/op     | 20224 B/op   | 167 allocs/op   |
| BenchmarkEcho_GithubAll        | 31251     | 38479 ns/op     | 0 B/op       | 0 allocs/op     |
| BenchmarkGocraftWeb_GithubAll  | 4117      | 300062 ns/op    | 131656 B/op  | 1686 allocs/op  |
| BenchmarkGoji_GithubAll        | 3274      | 416158 ns/op    | 56112 B/op   | 334 allocs/op   |
| BenchmarkGojiv2_GithubAll      | 1402      | 870518 ns/op    | 352720 B/op  | 4321 allocs/op  |
| BenchmarkGoJsonRest_GithubAll  | 2976      | 401507 ns/op    | 134371 B/op  | 2737 allocs/op  |
| BenchmarkGoRestful_GithubAll   | 410       | 2913158 ns/op   | 910144 B/op  | 2938 allocs/op  |
| BenchmarkGorillaMux_GithubAll  | 346       | 3384987 ns/op   | 251650 B/op  | 1994 allocs/op  |
| BenchmarkGowwwRouter_GithubAll | 10000     | 143025 ns/op    | 72144 B/op   | 501 allocs/op   |
| BenchmarkHttpRouter_GithubAll  | 55938     | 21360 ns/op     | 0 B/op       | 0 allocs/op     |
| BenchmarkHttpTreeMux_GithubAll | 10000     | 153944 ns/op    | 65856 B/op   | 671 allocs/op   |
| BenchmarkKocha_GithubAll       | 10000     | 106315 ns/op    | 23304 B/op   | 843 allocs/op   |
| BenchmarkLARS_GithubAll        | 47779     | 25084 ns/op     | 0 B/op       | 0 allocs/op     |
| BenchmarkMacaron_GithubAll     | 3266      | 371907 ns/op    | 149409 B/op  | 1624 allocs/op  |
| BenchmarkMartini_GithubAll     | 331       | 3444706 ns/op   | 226551 B/op  | 2325 allocs/op  |
| BenchmarkPat_GithubAll         | 273       | 4381818 ns/op   | 1483152 B/op | 26963 allocs/op |
| BenchmarkPossum_GithubAll      | 10000     | 164367 ns/op    | 84448 B/op   | 609 allocs/op   |
| BenchmarkR2router_GithubAll    | 10000     | 160220 ns/op    | 77328 B/op   | 979 allocs/op   |
| BenchmarkRivet_GithubAll       | 14625     | 82453 ns/op     | 16272 B/op   | 167 allocs/op   |
| BenchmarkTango_GithubAll       | 6255      | 279611 ns/op    | 63826 B/op   | 1618 allocs/op  |
| BenchmarkTigerTonic_GithubAll  | 2008      | 687874 ns/op    | 193856 B/op  | 4474 allocs/op  |
| BenchmarkTraffic_GithubAll     | 355       | 3478508 ns/op   | 820744 B/op  | 14114 allocs/op |
| BenchmarkVulcan_GithubAll      | 6885      | 193333 ns/op    | 19894 B/op   | 609 allocs/op   |

- (1): 在固定时间内实现的总重复次数，越高表示结果越可信
- (2): 单次重复耗时（ns/op），数值越小表示越好
- (3): 堆内存使用量（B/op），数值越小表示越好
- (4): 平均每次重复的分配数（allocs/op），数值越小表示越好。

### 中间件

​	你可以在[gin-contrib](https://github.com/gin-contrib)找到许多有用的Gin中间件。

### Users

使用[Gin](https://github.com/gin-gonic/gin)框架的令人敬畏的项目列表：

- [gorush](https://github.com/appleboy/gorush): 用Go编写的推送通知服务器。

- [fnproject](https://github.com/fnproject/fn): 容器原生的、与云平台无关的无服务器平台。

- [photoprism](https://github.com/photoprism/photoprism): 由 Go 和 Google TensorFlow 提供支持的个人照片管理工具。

- [lura](https://github.com/luraproject/lura): 带有中间件的超高性能 API 网关。

- [picfit](https://github.com/thoas/picfit): 用 Go 编写的图像调整服务器。

- [dkron](https://github.com/distribworks/dkron): 分布式、容错的工作调度系统。

  

### 贡献

​	Gin 是数百个贡献者的成果。我们感谢您的帮助！

​	关于提交补丁和贡献工作流程的细节，请参见[CONTRIBUTING](https://github.com/gin-gonic/gin/blob/v1.9.0/CONTRIBUTING.md)。



## 概述

​	gin包实现了一个名为gin的HTTP网络框架。

​	关于gin的更多信息，请参见https://gin-gonic.com/。



## 常量

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/context.go#L28)

``` go linenums="1"
const (
	MIMEJSON              = binding.MIMEJSON
	MIMEHTML              = binding.MIMEHTML
	MIMEXML               = binding.MIMEXML
	MIMEXML2              = binding.MIMEXML2
	MIMEPlain             = binding.MIMEPlain
	MIMEPOSTForm          = binding.MIMEPOSTForm
	MIMEMultipartPOSTForm = binding.MIMEMultipartPOSTForm
	MIMEYAML              = binding.MIMEYAML
	MIMETOML              = binding.MIMETOML
)
```

这是最常见的数据格式的Content-Type MIME类型。

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/gin.go#L73)

``` go linenums="1"
const (
    // PlatformGoogleAppEngine 当运行在 Google App Engine 时，使用 X-Appengine-Remote-Addr 来确定客户端的 IP。
	PlatformGoogleAppEngine = "X-Appengine-Remote-Addr"
    // PlatformCloudflare 当使用 Cloudflare 的 CDN 时，使用 CF-Connecting-IP 来确定客户端的 IP。
	PlatformCloudflare = "CF-Connecting-IP"
)
```

可信平台（Trusted platforms 通常指一些被认为是安全可靠的互联网服务或软件平台）

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/mode.go#L18)

``` go linenums="1"
const (
	// DebugMode 表示 gin 运行在 debug 模式下。    
	DebugMode = "debug"
	// ReleaseMode 表示 gin 运行在 release 模式下。
	ReleaseMode = "release"
	// TestMode 表示 gin 运行在 test 模式下。
	TestMode = "test"
)
```

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/auth.go#L17)

``` go linenums="1"
const AuthUserKey = "user"
```

AuthUserKey 是基本身份验证（basic auth）中用于存储用户凭证的 cookie 名称。

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/utils.go#L19)

``` go linenums="1"
const BindKey = "_gin-gonic/gin/bindkey"
```

BindKey 表示默认的绑定键。

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/context.go#L41)

``` go linenums="1"
const BodyBytesKey = "_gin-gonic/gin/bodybyteskey"
```

BodyBytesKey 表示默认的请求体字节键。

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/context.go#L44)

``` go linenums="1"
const ContextKey = "_gin-gonic/gin/contextkey"
```

ContextKey 是一个 Context 返回给自身的键。

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/mode.go#L16)

``` go linenums="1"
const EnvGinMode = "GIN_MODE"
```

EnvGinMode 表示 gin 模式的环境名称。

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/version.go#L8)

``` go linenums="1"
const Version = "v1.9.0"
```

Version 是当前 gin 框架的版本号。

## 变量

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/debug.go#L24)

```
var DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)
```

​	DebugPrintRouteFunc 表示调试日志输出格式。

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/mode.go#L44)

```
var DefaultErrorWriter io.Writer = os.Stderr
```

​	DefaultErrorWriter 是 Gin 用于调试错误的默认 io.Writer。

[View Source](https://github.com/gin-gonic/gin/blob/v1.9.0/mode.go#L41)

```
var DefaultWriter io.Writer = os.Stdout
```

​	DefaultWriter 是 Gin 用于调试输出和中间件输出（如 Logger() 或 Recovery()）的默认 io.Writer。请注意，Logger 和 Recovery 都提供自定义方式来配置其输出 io.Writer。要支持 Windows 中的着色，请使用：

```
import "github.com/mattn/go-colorable"
gin.DefaultWriter = colorable.NewColorableStdout()
```

## 函数

#### func CreateTestContext 

``` go linenums="1"
func CreateTestContext(w http.ResponseWriter) (c *Context, r *Engine)
```

​	CreateTestContext 用于测试目的返回一个新的 engine 和 context。

#### func Dir 

``` go linenums="1"
func Dir(root string, listDirectory bool) http.FileSystem
```

​	Dir返回一个http.FileSystem，可供http.FileServer()使用。如果listDirectory == true，则它与http.Dir()的作用相同，否则会返回一个文件系统，防止http.FileServer()列出目录文件。

#### func DisableBindValidation 

``` go linenums="1"
func DisableBindValidation()
```

​	DisableBindValidation关闭默认验证器。

#### func DisableConsoleColor 

``` go linenums="1"
func DisableConsoleColor()
```

​	DisableConsoleColor可以禁止控制台的颜色输出。

#### func EnableJsonDecoderDisallowUnknownFields 

``` go linenums="1"
func EnableJsonDecoderDisallowUnknownFields()
```

​	EnableJsonDecoderDisallowUnknownFields为binding.EnableDecoderDisallowUnknownFields设置为true，以便在JSON解码器实例上调用DisallowUnknownFields方法。

#### func EnableJsonDecoderUseNumber 

``` go linenums="1"
func EnableJsonDecoderUseNumber()
```

​	EnableJsonDecoderUseNumber为binding.EnableDecoderUseNumber设置为true，以便在JSON解码器实例上调用UseNumber方法。

#### func ForceConsoleColor 

``` go linenums="1"
func ForceConsoleColor()
```

​	ForceConsoleColor强制在控制台中输出颜色。

#### func IsDebugging 

``` go linenums="1"
func IsDebugging() bool
```

​	如果框架在调试模式下运行，IsDebugging返回true。使用SetMode(gin.ReleaseMode)来禁用调试模式。

#### func Mode 

``` go linenums="1"
func Mode() string
```

​	Mode返回当前的gin模式。

#### func SetMode 

``` go linenums="1"
func SetMode(value string)
```

​	SetMode根据传入的字符串设置gin的模式。

## 类型

### type Accounts 

``` go linenums="1"
type Accounts map[string]string
```

​	Accounts定义了授权登录的用户/密码列表的键/值。

### type Context 

``` go linenums="1"
type Context struct {    
    writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine       *Engine
	params       *Params
	skippedNodes *[]skippedNode

	// This mutex protects Keys map.
	mu sync.RWMutex

	// Keys 是一个键/值对，专门用于存储每个请求上下文的信息。
	Keys map[string]any

	// Errors 是一个列表，包含所有使用此上下文的处理程序/中间件附加的错误信息。
	Errors errorMsgs

	// Accepted 定义了手动接受的内容协商格式的列表。
	Accepted []string

	// queryCache 缓存了 c.Request.URL.Query() 的查询结果。
	queryCache url.Values

   // formCache 缓存了 c.Request.PostForm 中的解析后的表单数据，
   //该数据来自于 POST、PATCH 或 PUT 请求的 body 参数。
	formCache url.Values

	// SameSite 允许服务器定义一个 cookie 属性，
   // 使得浏览器在跨站点请求中无法发送该 cookie。
	sameSite http.SameSite
}
```

​	Context 是 Gin 中最重要的部分，它允许我们在中间件之间传递变量、管理请求的处理流程、验证请求中的 JSON 数据以及生成 JSON 响应等等。

#### func CreateTestContextOnly 

``` go linenums="1"
func CreateTestContextOnly(w http.ResponseWriter, r *Engine) (c *Context)
```

​	CreateTestContextOnly函数在engine的基础上返回一个新的用于测试目的上下文。

#### (*Context) Abort 

``` go linenums="1"
func (c *Context) Abort()
```

​	Abort方法可以阻止pending handlers（即：尚未执行的处理程序）被调用。注意，这不会停止当前的处理程序。假设你有一个（验证当前请求是否被授权的）授权中间件。如果授权失败（例如：密码不匹配），可以调用 Abort 来确保不会调用此请求的剩余处理程序。

!!! waring "注释"

	在这里，"pending handlers" 指的是还没有被调用的后续处理函数。例如，如果你在 Gin 框架中使用了多个中间件，每个中间件都有一个处理函数。那么，这些中间件中尚未执行的处理函数就是 "pending handlers"。在调用 Abort 后，这些 "pending handlers" 将不再被调用。

#### (*Context) AbortWithError 

``` go linenums="1"
func (c *Context) AbortWithError(code int, err error) *Error
```

​	AbortWithError方法在内部调用`AbortWithStatus()`和`Error()`方法。这个方法会中止处理链，写入状态码，并将指定的错误推送到`c.Errors`中。更多细节见Context.Error()。

#### (*Context) AbortWithStatus 

``` go linenums="1"
func (c *Context) AbortWithStatus(code int)
```

​	AbortWithStatus方法调用`Abort()`并写入写入指定状态码的响应头。例如，认证请求失败时可以使用：context.AbortWithStatus(401)。

#### (*Context) AbortWithStatusJSON 

``` go linenums="1"
func (c *Context) AbortWithStatusJSON(code int, jsonObj any)
```

​	AbortWithStatusJSON方法在内部先调用`Abort()`方法接着调用`JSON()`方法。这个方法中止了请求处理链，写入了指定的 HTTP 状态码并返回一个 JSON 格式的响应体。此外，它还将 Content-Type 响应头设置为 "application/json"。

#### (*Context) AddParam 

``` go linenums="1"
func (c *Context) AddParam(key, value string)
```

​	AddParam方法是用于在 e2e 测试中向上下文添加参数，以便替换路径参数键的值。例如，对于路由 "/user/:id"，调用 AddParam("id", 1) 的结果是 "/user/1"。

!!! waring "注释"

	"e2e" 是 "end-to-end" 的缩写，指的是端到端测试。这种测试方式是从用户的角度出发，测试整个软件系统是否能够正确地工作。它涉及到系统的各个组成部分，包括用户界面、服务器端、数据库、网络等等。在软件开发的过程中，端到端测试可以帮助保证整个系统的功能和性能符合预期。

#### (*Context) AsciiJSON 

``` go linenums="1"
func (c *Context) AsciiJSON(code int, obj any)
```

​	AsciiJSON方法将给定的结构体序列化为JSON，并将unicode转为ASCII字符串，放入响应主体。同时，它还将 Content-Type 响应头设置为 "application/json"。

#### (*Context) Bind 

``` go linenums="1"
func (c *Context) Bind(obj any) error
```

​	Bind方法检查请求的HTTP Method 和Content-Type，以自动选择一个绑定引擎，根据 "Content-Type"头部信息，使用不同的绑定方式，例如：

```
"application/json" --> JSON binding
"application/xml"  --> XML binding
```

​	如果Content-Type == "application/json"，它将请求的主体解析为JSON格式，使用JSON或XML作为JSON方法的输入。它将json的有效载荷解码为作为指针指定的结构体。如果输入无效，它会写入一个400错误并在响应中设置Content-Type为"text/plain"。

#### (*Context) BindHeader 

``` go linenums="1"
func (c *Context) BindHeader(obj any) error
```

​	BindHeader方法是c.MustBindWith(obj, binding.Header)的快捷方式。

#### (*Context) BindJSON 

``` go linenums="1"
func (c *Context) BindJSON(obj any) error
```

​	BindJSON方法是c.MustBindWith(obj, binding.JSON)的快捷方式。

#### (*Context) BindQuery 

``` go linenums="1"
func (c *Context) BindQuery(obj any) error
```

​	BindQuery方法是c.MustBindWith(obj, binding.Query)的快捷方式。

#### (*Context) BindTOML 

``` go linenums="1"
func (c *Context) BindTOML(obj interface{}) error
```

​	BindTOML方法是c.MustBindWith(obj, binding.TOML)的快捷方式。

#### (*Context) BindUri 

``` go linenums="1"
func (c *Context) BindUri(obj any) error
```

​	BindUri方法使用binding.Uri来绑定传递的结构体指针。如果发生任何错误，它将以HTTP 400中止请求。

#### (*Context) BindWith 

``` go linenums="1"
func (c *Context) BindWith(obj any, b binding.Binding) error
```

​	BindWith方法使用指定的binding engine绑定传递的结构体指针。参见binding 包。

#### (*Context) BindXML 

``` go linenums="1"
func (c *Context) BindXML(obj any) error
```

​	BindXML方法是c.MustBindWith(obj, binding.BindXML)的快捷方式。

#### (*Context) BindYAML 

``` go linenums="1"
func (c *Context) BindYAML(obj any) error
```

​	BindYAML方法是c.MustBindWith(obj, binding.YAML)的快捷方式。

#### (*Context) ClientIP 

``` go linenums="1"
func (c *Context) ClientIP() string
```

​	ClientIP 方法实现了一个尽力而为的算法来返回真实的客户端 IP。它在内部调用 c.RemoteIP() 来检查远程 IP 是否是可信的代理。如果是，它就会尝试解析 Engine.RemoteIPHeaders 中定义的标头（默认为 [X-Forwarded-For，X-Real-Ip]）。如果标头在语法上无效或远程 IP 不对应于可信代理，则返回来自 Request.RemoteAddr 的远程 IP。

#### (*Context) ContentType 

``` go linenums="1"
func (c *Context) ContentType() string
```

​	ContentType方法返回请求的Content-Type标头。

#### (*Context) Cookie 

``` go linenums="1"
func (c *Context) Cookie(name string) (string, error)
```

​	Cookie方法返回请求中指定名称的cookie值，如果未找到则返回ErrNoCookie。返回的cookie未经过转义。如果多个cookie与指定名称匹配，则只返回一个cookie。

#### (*Context) Copy 

``` go linenums="1"
func (c *Context) Copy() *Context
```

​	Copy方法返回当前上下文的一个副本，可以安全地在请求范围之外使用。当上下文需要被传递给一个goroutine时，**必须使用**这个方法。

#### (*Context) Data 

``` go linenums="1"
func (c *Context) Data(code int, contentType string, data []byte)
```

​	Data 方法将一些数据写入响应体，并更新 HTTP 状态码。

#### (*Context) DataFromReader 

``` go linenums="1"
func (c *Context) DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string)
```

​	DataFromReader方法将指定的读取器写入响应体流，并更新 HTTP 状态码。

#### (*Context) Deadline 

``` go linenums="1"
func (c *Context) Deadline() (deadline time.Time, ok bool)
```

​	当c.Request没有Context时，Deadline方法返回没有最后期限（ok==false）。

#### (*Context) DefaultPostForm 

``` go linenums="1"
func (c *Context) DefaultPostForm(key, defaultValue string) string
```

​	DefaultPostForm方法在POST urlencoded form或multipart form存在的情况下返回指定的key，否则它返回指定的defaultValue字符串。有关更多信息，请参见：PostForm()和GetPostForm()方法。

!!! warning "注释"

    "POST urlencoded form" 是指在 HTTP 请求体中使用 `application/x-www-form-urlencoded` 格式提交的表单数据，通常是通过表单提交、AJAX 等方式将数据提交给服务器。该格式会将数据用 `key1=value1&key2=value2` 的方式进行编码。
    "multipart form" 是指在 HTTP 请求体中使用 `multipart/form-data` 格式提交的表单数据，通常用于上传文件等场景。该格式会将数据分成多个部分，每部分包含一个头部和一个实体。每个实体可以是文本、二进制数据或者文件，多个实体之间用一个特殊的边界分隔符进行分隔。

#### (*Context) DefaultQuery 

``` go linenums="1"
func (c *Context) DefaultQuery(key, defaultValue string) string
```

​	DefaultQuery方法用于返回指定的url查询参数值，如果指定的参数不存在，则它返回defaultValue参数指定的字符串。有关更多信息，请参见：Query()和GetQuery()方法。

```
GET /?name=Manu&lastname=
c.DefaultQuery("name", "unknown") == "Manu"
c.DefaultQuery("id", "none") == "none"
c.DefaultQuery("lastname", "none") == ""
```

#### (*Context) Done 

``` go linenums="1"
func (c *Context) Done() <-chan struct{}
```

​	Done方法在 c.Request 没有 Context 的时候会返回 nil值的通道（它是一个无限等待的通道）。

#### (*Context) Err 

``` go linenums="1"
func (c *Context) Err() error
```

​	Err方法在 c.Request 没有 Context 的时候会返回 nil。

#### (*Context) Error 

``` go linenums="1"
func (c *Context) Error(err error) *Error
```

​	Error方法将一个错误附加到当前上下文。该错误被推送到错误列表中。在处理请求过程中，为每个发生的错误调用Error方法是个好主意。可以用一个中间件来收集所有的错误，并把它们一起推送到数据库中，打印日志，或者将其附加在 HTTP 响应中。如果err为nil，Error方法会引发panic。

#### (*Context) File 

``` go linenums="1"
func (c *Context) File(filepath string)
```

​	File方法以高效的方式将指定文件写入响应体流中。

#### (*Context) FileAttachment 

``` go linenums="1"
func (c *Context) FileAttachment(filepath, filename string)
```

​	FileAttachment方法将指定的文件以高效的方式写入响应主体中。在客户端，通常会使用给定的文件名（filename参数指定的）下载该文件。

#### (*Context) FileFromFS 

``` go linenums="1"
func (c *Context) FileFromFS(filepath string, fs http.FileSystem)
```

​	FileFromFS方法从 http.FileSystem 中读取指定的文件并以高效的方式写入响应体流中。

#### (*Context) FormFile 

``` go linenums="1"
func (c *Context) FormFile(name string) (*multipart.FileHeader, error)
```

​	FormFile方法返回给定表单键的第一个文件。

#### (*Context) FullPath 

``` go linenums="1"
func (c *Context) FullPath() string
```

​	FullPath方法返回匹配的路由的完整路径。如果没有找到匹配的路由则返回空字符串。

```
router.GET("/user/:id", func(c *gin.Context) {
    c.FullPath() == "/user/:id" // true
})
```

#### (*Context) Get 

``` go linenums="1"
func (c *Context) Get(key string) (value any, exists bool)
```

​	Get方法返回指定键的值，该值存在则返回（value，true）。如果该值不存在，则返回（nil, false）。

#### (*Context) GetBool 

``` go linenums="1"
func (c *Context) GetBool(key string) (b bool)
```

​	GetBool方法从请求参数中获取指定的 key 对应的值，并将该值解析为bool类型后返回。如果该key不存在，则会返回 `false`。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是bool类型，或者该 key 不存在，该方法会返回bool类型的零值。

#### (*Context) GetDuration 

``` go linenums="1"
func (c *Context) GetDuration(key string) (d time.Duration)
```

​	GetDuration方法从请求参数中获取指定的key对应的值，并将该值解析为 `time.Duration` 类型后返回。
!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是time.Duration类型，或者该 key 不存在，该方法会返回time.Duration类型的零值。


#### (*Context) GetFloat64 

``` go linenums="1"
func (c *Context) GetFloat64(key string) (f64 float64)
```

​	GetFloat64方法从请求参数中获取指定key对应的float64类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是float64类型，或者该 key 不存在，该方法会返回float64类型的零值。

#### (*Context) GetHeader 

``` go linenums="1"
func (c *Context) GetHeader(key string) string
```

​	GetHeader方法从请求头信息中获取指定key对应的值。

#### (*Context) GetInt 

``` go linenums="1"
func (c *Context) GetInt(key string) (i int)
```

​	GetInt方法从请求参数中获取指定key对应的 int 类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是int类型，或者该 key 不存在，该方法会返回int类型的零值。

#### (*Context) GetInt64 

``` go linenums="1"
func (c *Context) GetInt64(key string) (i64 int64)
```

​	GetInt64方法从请求参数中获取指定key对应的 int64 类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是int64类型，或者该 key 不存在，该方法会返回int64类型的零值。

#### (*Context) GetPostForm 

``` go linenums="1"
func (c *Context) GetPostForm(key string) (string, bool)
```

​	GetPostForm方法与PostForm(key)方法类似。当POST urlencoded form或multipart form存在时，它返回指定的key的值`(value, true)`（即使value是空字符串），否则它返回("", false)。例如，在更新用户的电子邮件的PATCH请求期间：

```
email=mail@example.com  -->  ("mail@example.com", true) := GetPostForm("email") // set email to "mail@example.com"
email=  -->  ("", true) := GetPostForm("email") // set email to ""
      -->  ("", false) := GetPostForm("email") // do nothing with email
```

!!! warning "注释"

    "POST urlencoded form" 是指在 HTTP 请求体中使用 `application/x-www-form-urlencoded` 格式提交的表单数据，通常是通过表单提交、AJAX 等方式将数据提交给服务器。该格式会将数据用 `key1=value1&key2=value2` 的方式进行编码。
    "multipart form" 是指在 HTTP 请求体中使用 `multipart/form-data` 格式提交的表单数据，通常用于上传文件等场景。该格式会将数据分成多个部分，每部分包含一个头部和一个实体。每个实体可以是文本、二进制数据或者文件，多个实体之间用一个特殊的边界分隔符进行分隔。



#### (*Context) GetPostFormArray 

``` go linenums="1"
func (c *Context) GetPostFormArray(key string) (values []string, ok bool)
```

​	GetPostFormArray方法根据给定的表单键返回一个字符串切片以及一个布尔值，该布尔值表示给定键是否至少存在一个值。

#### (*Context) GetPostFormMap 

``` go linenums="1"
func (c *Context) GetPostFormMap(key string) (map[string]string, bool)
```

​	GetPostFormMap方法根据给定的表单键返回一个映射以及一个布尔值，该布尔值表示给定键是否存在至少一个值。

#### (*Context) GetQuery 

``` go linenums="1"
func (c *Context) GetQuery(key string) (string, bool)
```

​	GetQuery方法类似于Query()方法，如果在url查询值中存在指定的key，则返回`(value, true)`（即使value是空字符串），否则它返回`(""", false)`。该方法是`c.Request.URL.Query().Get(key)`的快捷方式。	

```
GET /?name=Manu&lastname=
("Manu", true) == c.GetQuery("name")
("", false) == c.GetQuery("id")
("", true) == c.GetQuery("lastname")
```

#### (*Context) GetQueryArray 

``` go linenums="1"
func (c *Context) GetQueryArray(key string) (values []string, ok bool)
```

​	GetQueryArray方法根据给定的查询键返回一个字符串切片以及一个布尔值，该布尔值表示给定键是否至少存在一个值。

#### (*Context) GetQueryMap 

``` go linenums="1"
func (c *Context) GetQueryMap(key string) (map[string]string, bool)
```

​	GetQueryMap方法根据给定的查询键返回一个映射以及一个布尔值，该布尔值表示给定键是否存在至少一个值。

#### (*Context) GetRawData 

``` go linenums="1"
func (c *Context) GetRawData() ([]byte, error)
```

​	GetRawData方法获取请求体中的原始数据，返回的是字节切片。

#### (*Context) GetString 

``` go linenums="1"
func (c *Context) GetString(key string) (s string)
```

​	GetString方法从请求参数中获取指定key对应的string类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是string类型，或者该 key 不存在，该方法会返回string类型的零值。

#### (*Context) GetStringMap 

``` go linenums="1"
func (c *Context) GetStringMap(key string) (sm map[string]any)
```

​	GetStringMap方法从请求参数中获取指定key对应的`map[string]any`类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是map[string]any类型，或者该 key 不存在，该方法会返回map[string]any类型的零值。

#### (*Context) GetStringMapString 

``` go linenums="1"
func (c *Context) GetStringMapString(key string) (sms map[string]string)
```

​	GetStringMapString方法从请求参数中获取指定key对应的`map[string]string`类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是map[string]string类型，或者该 key 不存在，该方法会返回map[string]string类型的零值。

#### (*Context) GetStringMapStringSlice 

``` go linenums="1"
func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string)
```

​	GetStringMapStringSlice方法从请求参数中获取指定key对应的`map[string][]string`类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是map[string][]string类型，或者该 key 不存在，该方法会返回map[string][]string类型的零值。

#### (*Context) GetStringSlice 

``` go linenums="1"
func (c *Context) GetStringSlice(key string) (ss []string)
```

​	GetStringSlice方法从请求参数中获取指定key对应的`[]string`类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是[]string类型，或者该 key 不存在，该方法会返回[]string类型的零值。

#### (*Context) GetTime 

``` go linenums="1"
func (c *Context) GetTime(key string) (t time.Time)
```

​	GetTime 方法从请求参数中获取指定key对应的`time.Time`类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是time.Time类型，或者该 key 不存在，该方法会返回time.Time类型的零值。

#### (*Context) GetUint 

``` go linenums="1"
func (c *Context) GetUint(key string) (ui uint)
```

​	GetUint方法从请求参数中获取指定key对应的`uint`类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是uint类型，或者该 key 不存在，该方法会返回uint类型的零值。

#### (*Context) GetUint64 

``` go linenums="1"
func (c *Context) GetUint64(key string) (ui64 uint64)
```

​	GetUint64方法从请求参数中获取指定key对应的`uint64`类型的值。

!!! warning "注意"

	需要注意的是，如果请求中指定 key 的值不是uint64类型，或者该 key 不存在，该方法会返回uint64类型的零值。

#### (*Context) HTML 

``` go linenums="1"
func (c *Context) HTML(code int, name string, obj any)
```

​	HTML方法根据指定的模板文件名渲染 HTTP 模板。它还更新HTTP状态码并将Content-Type设置为"text/html"。参见[http://golang.org/doc/articles/wiki/](http://golang.org/doc/articles/wiki/)。

#### (*Context) Handler 

``` go linenums="1"
func (c *Context) Handler() HandlerFunc
```

​	Handler方法返回主处理程序。

#### (*Context) HandlerName 

``` go linenums="1"
func (c *Context) HandlerName() string
```

​	HandlerName方法返回主处理程序的名称。例如，如果处理程序是 "handleGetUsers()"，这个方法将返回 "main.handleGetUsers"。

#### (*Context) HandlerNames 

``` go linenums="1"
func (c *Context) HandlerNames() []string
```

​	HandlerNames方法按照HandlerName()方法的语义，以降序返回该上下文的所有注册处理程序的列表。

#### (*Context) Header 

``` go linenums="1"
func (c *Context) Header(key, value string)
```

​	Header方法是c.Writer.Header().Set(key, value)的一个智能快捷方式。它在响应中写了一个头信息。如果 value == ""，该方法将删除用key指定的头信息，相当于调用了`c.Writer.Header().Del(key)`。

#### (*Context) IndentedJSON 

``` go linenums="1"
func (c *Context) IndentedJSON(code int, obj any)
```

​	IndentedJSON方法将给定的结构体序列化为漂亮的JSON（带缩进和换行符）并写入响应体中。它还将Content-Type设置为"application/json”。

> 警告：我们建议仅在开发过程中使用此方法，因为打印漂亮的JSON更耗费CPU和带宽。请改用Context.JSON()。

#### (*Context) IsAborted 

``` go linenums="1"
func (c *Context) IsAborted() bool
```

​	如果当前的上下文被中止了，IsAborted方法返回true。

#### (*Context) IsWebsocket 

``` go linenums="1"
func (c *Context) IsWebsocket() bool
```

​	如果请求头指示客户端正在启动websocket握手，则IsWebsocket方法返回true。

#### (*Context) JSON 

``` go linenums="1"
func (c *Context) JSON(code int, obj any)
```

​	JSON方法将给定的结构体序列化为JSON，写入响应体。它还将Content-Type设置为"application/json"。

#### (*Context) JSONP 

``` go linenums="1"
func (c *Context) JSONP(code int, obj any)
```

​	JSONP将给定的结构体序列化为JSON格式的响应体。它为响应体添加了填充，以请求来自与客户端不同域的服务器的数据。它还将Content-Type设置为 "application/javascript"。

> ​	当一个网页从一个域名请求另一个域名的资源时，浏览器会根据同源策略限制请求。为了解决这个问题，JSONP（JSON with padding）应运而生。
>
> ​	JSONP（JSON with padding）是一种用于解决跨域数据请求的方法，它通过在响应中添加一个包含回调函数的JavaScript函数调用，来让页面能够访问从不同域的服务器获取的JSON数据。JSONP是一种hack方法，由于它的不安全性已被广泛批评，现在已经被CORS（Cross-Origin Resource Sharing：跨源资源共享）等更现代的技术所取代。
>
> otherSite.com中使用gin的JSONP实现：
>
> ```go linenums="1"
> func main() {
> 	r := gin.Default()
> 	r.GET("/jsonp", func(c *gin.Context) {
> 		c.JSONP(200, gin.H{"name": "zlongx"})
> 	})
> 	r.Run(":8080")
> }
> ```
>
> 本站点前端HTML页面：
>
> ```html
> <script type="text/javascript">
> function sayHello(data){
> 	var obj = JSON.stringify(data);
>   	alert(obj.name);
> }
> </script>
> <script type="text/javascript" src="http://otherSite.com/jsonp?callback=sayHello"></script>
> ```
>
> 当浏览器执行到 `<script type="text/javascript" src="http://otherSite.com/jsonp?callback=sayHello"></script>`,就会执行 `http://otherSite.com/jsonp?callback=sayHello`请求，从而从otherSite.com网站中获取到数据，最终又调用到本站点HTML中的sayHello函数。

#### (*Context) MultipartForm 

``` go linenums="1"
func (c *Context) MultipartForm() (*multipart.Form, error)
```

​	MultipartForm方法获取经过解析的multipart form，包括文件上传。

#### (*Context) MustBindWith 

``` go linenums="1"
func (c *Context) MustBindWith(obj any, b binding.Binding) error
```

​	MustBindWith方法用指定的绑定引擎来绑定传递的结构指针。如果发生任何错误，它将以HTTP 400中止请求。参见binding包。

#### (*Context) MustGet 

``` go linenums="1"
func (c *Context) MustGet(key string) any
```

​	MustGet方法返回给定key的值，如果不存在则会抛出panic。

> 一般来说，建议在使用 MustGet 方法之前，先使用该 Context 的 Value 方法获取指定 key 对应的值，并对返回值进行类型断言，以确保程序不会抛出异常。

#### (*Context) Negotiate 

``` go linenums="1"
func (c *Context) Negotiate(code int, config Negotiate)
```

​	Negotiate方法根据可接受的Accept格式调用不同的Render。

#### (*Context) NegotiateFormat 

``` go linenums="1"
func (c *Context) NegotiateFormat(offered ...string) string
```

​	NegotiateFormat方法返回一个可接受的Accept格式。

> `Context.NegotiateFormat()` 方法是用来协商响应数据格式的。它会根据客户端支持的数据格式（`Accept` 请求头）来确定服务器返回的数据格式。在 GIN 框架中，`NegotiateFormat()` 方法会根据客户端请求头的 `Accept` 字段内容来决定采用何种响应数据格式，比如 JSON 或者 XML 等。在调用此方法前，需要先注册相应的渲染函数，例如 JSON 渲染函数和 XML 渲染函数等。如果客户端支持多种数据格式，那么 `NegotiateFormat()` 方法将选择客户端优先级最高的数据格式返回。

#### (*Context) Next 

``` go linenums="1"
func (c *Context) Next()
```

​	Next方法应仅在中间件内部使用。它在调用处理程序内部执行待处理程序链中的挂起处理程序。见GitHub中的例子。

> ​	在 Gin 框架中，`Next` 方法用于执行链中的下一个中间件或路由处理程序。它只应在中间件中使用，因为在处理请求的控制器或路由处理程序中调用它会导致请求的处理过早结束。
>
> ​	当 `Next` 方法被调用时，它将在调用当前处理程序之前暂停执行，并在调用下一个处理程序后继续执行。如果没有更多的处理程序，则控制流程将回到上一个处理程序，并从那里继续执行。
>
> 这种"委托式”的流程控制使得 Gin 的中间件非常灵活，可以很容易地添加、修改或删除处理步骤。下面是一个示例：
>
> ```go linenums="1"
> func AuthMiddleware() gin.HandlerFunc {
>   return func(c *gin.Context) {
>     // Perform authentication
>     if authenticated {
>       // Call the next middleware/handler in the chain
>       c.Next()
>     } else {
>       // Abort the request and send an error response
>       c.AbortWithStatus(http.StatusUnauthorized)
>     }
>   }
> }
> 
> func MyHandler(c *gin.Context) {
>   // Do some processing
>   c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
> }
> 
> func main() {
>   r := gin.Default()
> 
>   r.Use(AuthMiddleware())
> 
>   r.GET("/", MyHandler)
> 
>   r.Run(":8080")
> }
> 
> ```
>
> ​	在上面的示例中，`AuthMiddleware` 中间件执行身份验证，如果身份验证成功，则调用 `Next` 方法继续执行下一个处理程序。如果身份验证失败，则调用 `AbortWithStatus` 方法终止请求并发送错误响应。在 `MyHandler` 处理程序中，我们将返回一个 JSON 响应。
>
> ​	请注意，我们在路由中使用 `r.Use(AuthMiddleware())` 将 `AuthMiddleware` 中间件注册为路由级别的中间件。这意味着该中间件将在路由处理程序之前执行。当 `AuthMiddleware` 中间件调用 `Next` 方法时，控制流程将继续到 `MyHandler` 处理程序。如果 `AuthMiddleware` 中间件调用 `AbortWithStatus` 方法，则控制流程将不会到达 `MyHandler` 处理程序。

#### (*Context) Param 

``` go linenums="1"
func (c *Context) Param(key string) string
```

​	Param方法返回URL参数的值。它是c.Params.ByName(key)的快捷方式。

```
router.GET("/user/:id", func(c *gin.Context) {
    // a GET request to /user/john
    id := c.Param("id") // id == "/john"
    // a GET request to /user/john/
    id := c.Param("id") // id == "/john/"
})
```

!!! warning "注意"

    `Context.Param(key string)` 方法用于获取路由参数中的值，例如在路由模式为 `/users/:id` 时，使用该方法可以获取请求中的 `id` 值。
    `Context.Get(key string)` 方法用于获取请求参数中的值，无论是 `GET` 请求中的查询参数还是 `POST` 请求中的表单参数。
    因此，这两个方法的作用对象不同，一个用于获取路由参数，一个用于获取请求参数。

#### (*Context) PostForm 

``` go linenums="1"
func (c *Context) PostForm(key string) (value string)
```

​	PostForm方法返回POST urlencoded form或multipart form中指定的key的值，如果不存在该key则返回空字符串。

> urlencoded form 和 multipart form 的说明，请参照：[DefaultPostForm方法](#context-defaultpostform)

#### (*Context) PostFormArray 

``` go linenums="1"
func (c *Context) PostFormArray(key string) (values []string)
```

​	PostFormArray方法返回给定表单键的字符串切片。切片的长度取决于具有给定键的参数数量。

> 假设客户端发送了一个 POST 请求，请求体是：
>
> ```
> fruit=apple&fruit=banana&fruit=orange
> ```
>
> 那么在处理这个请求的处理器函数中，我们可以这样使用 `PostFormArray` 方法：
>
> ```go linenums="1"
> func handleFruit(c *gin.Context) {
>     fruits := c.PostFormArray("fruit")
>     fmt.Printf("Length of the fruits slice: %d\n", len(fruits))
>     fmt.Printf("Fruits: %v\n", fruits)
>     // ...
> }
> ```
>
> 运行这个处理器函数，输出将是：
>
> ```
> Length of the fruits slice: 3
> Fruits: [apple banana orange]
> ```
>
> 这里，`PostFormArray("fruit")` 方法返回一个长度为 3 的字符串切片，包含请求中的所有水果。

#### (*Context) PostFormMap 

``` go linenums="1"
func (c *Context) PostFormMap(key string) (dicts map[string]string)
```

​	PostFormMap方法返回给定表单键的映射。

> `Context.PostFormMap()` 方法用于获取表单中指定 key 对应的值并以 map 的形式返回，其中 map 中的 key 为表单中指定的 key，value 为相应的 value 值。

#### (*Context) ProtoBuf 

``` go linenums="1"
func (c *Context) ProtoBuf(code int, obj any)
```

​	ProtoBuf方法将给定的结构体序列化为 ProtoBuf，并将其写入响应体中。

> ProtoBuf 是 Protocol Buffers 的缩写，是 Google 开发的一种轻便高效的结构化数据序列化方法，主要用于数据存储、通信协议等领域。ProtoBuf 通过将结构化的数据序列化为二进制数据，从而实现了高效地数据传输和存储。在 Gin 框架中，ProtoBuf 用于将结构体序列化为 ProtoBuf 格式，然后发送到客户端。这样可以实现更高效的数据传输，同时也提供了一种跨语言的数据交换格式。

#### (*Context) PureJSON 

``` go linenums="1"
func (c *Context) PureJSON(code int, obj any)
```

​	PureJSON方法将给定的结构体序列化为 JSON，并将其写入响应体中。与 JSON方法不同，PureJSON方法不会用其Unicode 实体替换特殊的 HTML 字符。

> `JSON()` 方法是将结构体转换为 JSON 格式，并使用 `gin` 的默认字符编码设置响应头的 `Content-Type` 为 `application/json`。此外，`JSON()` 方法还会将特殊的 HTML 字符（如 `<`, `>`, `&`）替换为它们的 Unicode 实体，以避免浏览器将其解释为 HTML 标签或其他特殊字符。
>
> `PureJSON()` 方法则不会将特殊的 HTML 字符替换为 Unicode 实体。这意味着如果返回的 JSON 数据中包含特殊字符，则客户端将看到这些字符，而不是它们的实体。`PureJSON()` 方法仅仅将结构体转换为 JSON 格式并返回给客户端，并不会设置响应头的 `Content-Type`，需要在自己的代码中设置。
>
> 因此，如果需要在 JSON 数据中包含特殊字符，可以使用 `PureJSON()` 方法。如果不需要，则可以使用 `JSON()` 方法，它会在大多数情况下更加安全。

#### (*Context) Query 

``` go linenums="1"
func (c *Context) Query(key string) (value string)
```

​	如果key存在url查询字符串中的话，则Query方法返回key的值，否则返回一个空字符串`("")`。它是`c.Request.URL.Query().Get(key)`的快捷方式。

```
GET /path?id=1234&name=Manu&value=
	   c.Query("id") == "1234"
	   c.Query("name") == "Manu"
	   c.Query("value") == ""
	   c.Query("wtf") == ""
```

#### (*Context) QueryArray 

``` go linenums="1"
func (c *Context) QueryArray(key string) (values []string)
```

​	QueryArray 方法返回给定查询key（即该key在url查询字符串）的字符串切片。切片的长度取决于具有给定键的参数数量。

> 例如：请求url为：`http://example.com/colors?color=red&color=green&color=blue`
>
> 通过以下代码：
>
> ```go linenums="1"
> colors := c.QueryArray("color")
> fmt.Printf("%#v",colors)
> ```
>
> 获取的colors的值为：["red","green","blue"]

#### (*Context) QueryMap 

``` go linenums="1"
func (c *Context) QueryMap(key string) (dicts map[string]string)
```

​	QueryMap方法返回给定查询key（即该key在url查询字符串）的映射。

> 例如：请求url为：`http://example.com/nameAge?name[zw]=30&name[lx]=66`
>
> 通过以下代码：
>
> ```go linenums="1"
> s := c.QueryArray("name")
> fmt.Printf("%#v",s)
> ```
>
> 获取的name的值为：map[string]string{"zw":"30","lx":"66"}

#### (*Context) Redirect 

``` go linenums="1"
func (c *Context) Redirect(code int, location string)
```

​	Redirect方法将 HTTP 重定向到指定位置。

> 在 `Context.Redirect()` 方法中，第一个参数是重定向的状态码，例如可以使用 `http.StatusMovedPermanently`，表示永久性重定向；第二个参数是重定向的 URL。

#### (*Context) RemoteIP 

``` go linenums="1"
func (c *Context) RemoteIP() string
```

​	RemoteIP方法从 Request.RemoteAddr 解析 IP，将其标准化并返回 IP（不包括端口）。

#### (*Context) Render 

``` go linenums="1"
func (c *Context) Render(code int, r render.Render)
```

​	Render方法写入响应头并调用 render.Render方法渲染数据。

#### (*Context) SSEvent 

``` go linenums="1"
func (c *Context) SSEvent(name string, message any)
```

​	SSEvent方法将服务器发送事件写入响应体流中。

#### (*Context) SaveUploadedFile 

``` go linenums="1"
func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error
```

​	SaveUploadedFile方法将表单文件上传到指定的dst。

#### (*Context) SecureJSON 

``` go linenums="1"
func (c *Context) SecureJSON(code int, obj any)
```

​	SecureJSON方法将给定的结构体序列化为安全的JSON，放入响应体。如果给定的结构体是数组值，默认会在响应体前面添加 "while(1),"。它还将Content-Type设置为 "application/json"。

> 可以通过 Engine.SecureJsonPrefix() 方法，设置防劫持前缀以替换默认的"while(1),"。

#### (*Context) Set 

``` go linenums="1"
func (c *Context) Set(key string, value any)
```

​	Set方法是用来存储一个新的键/值对，专门用于这个上下文。如果之前未使用 c.Keys，则还会惰性初始化。

#### (*Context) SetAccepted 

``` go linenums="1"
func (c *Context) SetAccepted(formats ...string)
```

​	SetAccepted方法设置Accept标头数据。

#### (*Context) SetCookie 

``` go linenums="1"
func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
```

​	SetCookie方法向 ResponseWriter【一种接口，在下文定义[ResponseWriter接口](#type-responsewriter)】 的标头添加 Set-Cookie 标头。所提供的 cookie 必须具有有效的 Name。无效的 cookie 可能会被静默丢弃。

#### (*Context) SetSameSite 

``` go linenums="1"
func (c *Context) SetSameSite(samesite http.SameSite) {
	c.sameSite = samesite
}
```

SetSameSite with cookie

带cookie的SetSameSite

使用 SetSameSite 与 cookie。

> 参见：[http.SameSite类型](../../../StdLib/net/http#type-samesite)

> Context.SetSameSite() 方法用于设置 Cookie 的 SameSite 属性，SameSite 属性用于防止 CSRF（跨站请求伪造）攻击和其他类似攻击。该方法有一个参数，类型为字符串，取值可以是 "Strict"、"Lax" 或 "None"。
>
> - "Strict" 表示严格模式，Cookie 在跨站请求时不会发送，只能在同站点请求中使用。
> - "Lax" 表示宽松模式，Cookie 在安全的跨站请求（例如 GET 请求）时可以发送，但在其他跨站请求（例如 POST 请求）时不会发送。
> - "None" 表示不使用 SameSite 属性，Cookie 可以在任何跨站请求中发送，不过需要设置 secure 属性为 true，表示只有在 HTTPS 连接中才能发送。
>
> 以下是一个示例：
>
> ```go linenums="1"
> func handler(c *gin.Context) {
>     cookie := http.Cookie{
>         Name:     "test",
>         Value:    "123",
>         SameSite:  "Lax",
>         Secure:    true,
>     }
>     http.SetCookie(c.Writer, &cookie)
>     c.String(http.StatusOK, "Cookie set!")
> }
> 
> ```
>
> 上述代码中，我们使用 http.Cookie 结构体创建了一个名为 "test" 的 Cookie，设置了其值为 "123"，SameSite 属性为 "Lax"，secure 属性为 true，表示只能在 HTTPS 连接中发送。然后，我们通过 http.SetCookie() 方法将该 Cookie 添加到响应头中，并在响应体中返回 "Cookie set!" 字符串。

#### (*Context) ShouldBind 

``` go linenums="1"
func (c *Context) ShouldBind(obj any) error
```

​	ShouldBind方法检查 Method 和 Content-Type 请求头以自动选择binding引擎。根据 "Content-Type" 请求头使用不同的binding，例如：

```
"application/json" --> JSON binding
"application/xml"  --> XML binding
```

​	如果 Content-Type == "application/json"，它将请求的主体解析为 JSON，使用 JSON 或 XML 作为 JSON方法输入。将 JSON 负载解码为指定为指针的结构体。类似于 c.Bind()，但是如果输入无效，此方法不会将响应状态码设置为 400 或中止。

#### (*Context) ShouldBindBodyWith 

``` go linenums="1"
func (c *Context) ShouldBindBodyWith(obj any, bb binding.BindingBody) (err error)
```

​	ShouldBindBodyWith方法类似于 ShouldBindWith方法，但它将请求主体存储到该上下文中，并在再次调用时重用。

> 注意：此方法在绑定之前读取请求主体。因此，如果只需要调用一次，应使用 ShouldBindWith方法获得更好的性能。

#### (*Context) ShouldBindHeader 

``` go linenums="1"
func (c *Context) ShouldBindHeader(obj any) error
```

​	ShouldBindHeader方法是 c.ShouldBindWith(obj, binding.Header) 的快捷方式。

#### (*Context) ShouldBindJSON 

``` go linenums="1"
func (c *Context) ShouldBindJSON(obj any) error
```

​	ShouldBindJSON方法是 c.ShouldBindWith(obj, binding.JSON) 的快捷方式。

#### (*Context) ShouldBindQuery 

``` go linenums="1"
func (c *Context) ShouldBindQuery(obj any) error
```

​	ShouldBindQuery方法是 c.ShouldBindWith(obj, binding.Query) 的快捷方式。

#### (*Context) ShouldBindTOML 

``` go linenums="1"
func (c *Context) ShouldBindTOML(obj interface{}) error
```

​	ShouldBindTOML方法是 c.ShouldBindWith(obj, binding.TOML) 的快捷方式。

#### (*Context) ShouldBindUri 

``` go linenums="1"
func (c *Context) ShouldBindUri(obj any) error
```

​	ShouldBindUri方法使用指定的binding引擎绑定传递的结构体指针。

#### (*Context) ShouldBindWith 

``` go linenums="1"
func (c *Context) ShouldBindWith(obj any, b binding.Binding) error
```

​	ShouldBindWith方法使用指定的binding引擎绑定传递的结构体指针。请参阅binding包。

#### (*Context) ShouldBindXML 

``` go linenums="1"
func (c *Context) ShouldBindXML(obj any) error
```

​	ShouldBindXML方法是 c.ShouldBindWith(obj, binding.XML) 的快捷方式。

#### (*Context) ShouldBindYAML 

``` go linenums="1"
func (c *Context) ShouldBindYAML(obj any) error
```

​	ShouldBindYAML方法是 c.ShouldBindWith(obj, binding.YAML) 的快捷方式。

#### (*Context) Status 

``` go linenums="1"
func (c *Context) Status(code int)
```

​	Status方法设置HTTP响应状态码。

#### (*Context) Stream 

``` go linenums="1"
func (c *Context) Stream(step func(w io.Writer) bool) bool
```

​	Stream方法发送流式响应，并返回一个布尔值，指示"客户端在流程中是否断开连接”。

#### (*Context) String 

``` go linenums="1"
func (c *Context) String(code int, format string, values ...any)
```

​	String方法将给定的字符串写入响应体中。

#### (*Context) TOML 

``` go linenums="1"
func (c *Context) TOML(code int, obj interface{})
```

​	TOML方法将给定的结构体序列化为 TOML，并将其写入响应正文中。

#### (*Context) Value 

``` go linenums="1"
func (c *Context) Value(key any) any
```

​	Value方法返回与该上下文关联的键 key 的值，如果没有与该 key 关联的值，则返回nil。用相同的key连续调用Value方法将返回相同的结果。

#### (*Context) XML 

``` go linenums="1"
func (c *Context) XML(code int, obj any)
```

​	XML方法将给定的结构体序列化为 XML，并将 Content-Type 设置为 "application/xml"。

#### (*Context) YAML 

``` go linenums="1"
func (c *Context) YAML(code int, obj any)
```

​	YAML方法将给定的结构体序列化为 YAML，并将其写入响应正文中。

### type Engine 

``` go linenums="1"
type Engine struct {
	RouterGroup
    
    // 如果启用 RedirectTrailingSlash，并且当前路由无法匹配，
    // 但是存在一个带（或不带）尾部斜杠的路径处理程序，则会自动重定向。
    // 例如，如果请求了 /foo/，但是只有 /foo 的路由存在，
    // 则将客户端重定向到 /foo，
    // 对于 GET 请求，状态码为 301，
    // 对于所有其他请求方法，状态码为 307。
	RedirectTrailingSlash bool

    // 如果启用 RedirectFixedPath且当前请求路径未注册处理程序，
    // 则路由器将尝试修复它。
    // 首先删除多余的路径元素，如 ../ 或 //。
    // 然后，路由器对已清理的路径进行不区分大小写的查找。
    // 如果可以为此路由找到处理程序，
    // 则路由器将以状态码301(对于GET请求)
    // 和状态码307（对于所有其他请求方法）重定向到已更正的路径。
    // 例如，/FOO 和 /..//Foo 可能会重定向到 /foo。
    // RedirectTrailingSlash 与此选项无关。
	RedirectFixedPath bool

    // 如果启用 HandleMethodNotAllowed，且当无法路由当前请求时，
    // 则路由器会检查当前路由是否允许其他方法处理。
    // 如果是这种情况，则响应请求以"Method Not Allowed”和405HTTP状态码。
    // 如果没有其他方法允许，则请求将委托给 NotFound 处理程序。
	HandleMethodNotAllowed bool

    // 如果启用 ForwardedByClientIP，
    // 则客户端 IP 将从请求的头部解析，
    // 这些头部与(*gin.Engine).RemoteIPHeaders存储的头部相匹配。
    // 如果未获取到 IP，则返回到从(*gin.Context).Request.RemoteAddr获取的IP。
	ForwardedByClientIP bool

    // AppEngine 已被弃用。
    // 请使用带有值 gin.PlatformGoogleAppEngine 的 TrustedPlatform 代替。
    // #726 #755 如果启用，它将信任某些以"X-AppEngine…”开头的标头，
    // 以更好地与该 PaaS 集成。
    // #726 和 #755 是该库的 GitHub 仓库中相关 issue 的编号。
    // #726 是关于在 Google App Engine 上使用 gin 的问题，
    // #755 是关于 App Engine Standard 环境中的用户 IP 的问题。
    // 在这段注释中，提到这两个 issue 是因为它们与 AppEngine 字段的弃用相关。
    // PaaS 是 Platform as a Service 的缩写，即"平台即服务”。
    // PaaS 提供一种平台来运行应用程序，
    // 可以让开发人员更专注于开发应用程序的核心功能，
    // 而不必担心运维的问题，如服务器管理、网络配置、数据库维护等。
    // 常见的 PaaS 服务提供商包括 Heroku、Google App Engine、Microsoft Azure 等。
	AppEngine bool

    // 如果启用UseRawPath，则将使用 url.RawPath 查找参数。
	UseRawPath bool

    // 如果UnescapePathValues为true，则路径值将是未经转义的。
    // 如果UseRawPath为false（默认值），则UnescapePathValues实际上为true，
    // 因为将使用(未经转义的）url.Path。
	UnescapePathValues bool

	// RemoveExtraSlash选项可以让URL中包含的额外斜杠不影响参数的解析。
    // 请参考 PR #1817 和 issue #1644。
	RemoveExtraSlash bool

    // RemoteIPHeaders是一个用于获取客户端 IP 的头列表。
    // 当(*gin.Engine).ForwardedByClientIP 为 true 
    // 且 (*gin.Context).Request.RemoteAddr被(*gin.Engine).SetTrustedProxies()
    // 定义的网络来源列表中的至少一个所匹配时，
    // 将使用此列表中的头来获取客户端 IP。
	RemoteIPHeaders []string

	// 如果将 TrustedPlatform 设置为 gin.Platform* 常量的值，
    // 则信任由该平台设置的头部信息，例如用于确定客户端 IP。
	TrustedPlatform string
	
    // MaxMultipartMemory是传递给http.Request的ParseMultipartForm方法调用的maxMemory参数的值。
	MaxMultipartMemory int64

    // UseH2C开启h2c的支持。
    // h2c 是基于 HTTP/2 协议的扩展，是一种轻量级的协议，
    // 通常用于支持 WebSockets 和服务器推送等功能。
    // 它允许无需预先建立 SSL 连接即可直接使用 HTTP/2 协议进行通信。
	UseH2C bool

    // ContextWithFallback启用回退Context.Deadline()、Context.Done()、
    // Context.Err()和Context.Value()，当Context.Request.Context()不为nil时。
	ContextWithFallback bool

	HTMLRender render.HTMLRender
	FuncMap    template.FuncMap
	// contains filtered or unexported fields
}
```

​	Engine 是 Gin 框架的实例，它包含路由器 (muxer)、中间件和配置设置。可以通过使用New()方法或Default()方法创建一个Engine实例。

#### func Default 

``` go linenums="1"
func Default() *Engine
```

​	Default函数返回一个已经附加了Logger和Recovery中间件的Engine实例。

#### func New 

``` go linenums="1"
func New() *Engine
```

​	New函数返回一个新的没有附加任何中间件的Engine 实例。

​	默认配置为：

```
- RedirectTrailingSlash: true 
- RedirectFixedPath: false 
- HandleMethodNotAllowed: false 
- ForwardedByClientIP: true 
- UseRawPath: false - UnescapePathValues: true
```



#### (*Engine) Delims 

``` go linenums="1"
func (engine *Engine) Delims(left, right string) *Engine
```

​	Delims方法设置模板的左右定界符并返回一个 Engine 实例。

#### (*Engine) HandleContext 

``` go linenums="1"
func (engine *Engine) HandleContext(c *Context)
```

​	HandleContext方法重新进入已被重写的上下文。可以通过将 c.Request.URL.Path 设置为新目标来实现。免责声明：您可以通过循环处理此问题，请明智使用。

#### (*Engine) Handler 

``` go linenums="1"
func (engine *Engine) Handler() http.Handler
```

#### (*Engine) LoadHTMLFiles 

``` go linenums="1"
func (engine *Engine) LoadHTMLFiles(files ...string)
```

​	LoadHTMLFiles方法加载一组 HTML 文件并将结果与 HTML 渲染器关联。

#### (*Engine) LoadHTMLGlob 

``` go linenums="1"
func (engine *Engine) LoadHTMLGlob(pattern string)
```

​	LoadHTMLGlob方法加载由 glob 模式识别的 HTML 文件并将结果与 HTML 渲染器关联。

#### (*Engine) NoMethod 

``` go linenums="1"
func (engine *Engine) NoMethod(handlers ...HandlerFunc)
```

NoMethod 设置当 Engine.HandleMethodNotAllowed 为 true 时的处理程序。

#### (*Engine) NoRoute 

``` go linenums="1"
func (engine *Engine) NoRoute(handlers ...HandlerFunc)
```

​	NoRoute方法添加 NoRoute 的处理程序。默认情况下返回 404状态码。

> 在 Gin 框架中，`NoRoute()` 方法用于设置在没有匹配到路由时的处理函数，即 404 页面处理。它可以设置一个或多个处理函数，处理函数的参数是 `Context` 对象。
>
> 示例代码如下：
>
> ```go linenums="1"
> package main
> 
> import (
> 	"github.com/gin-gonic/gin"
> 	"net/http"
> )
> 
> func main() {
> 	r := gin.Default()
> 
> 	// 定义404处理函数
> 	r.NoRoute(func(c *gin.Context) {
> 		c.JSON(http.StatusNotFound, gin.H{"message": "404 Not Found"})
> 	})
> 
> 	r.GET("/hello", func(c *gin.Context) {
> 		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin!"})
> 	})
> 
> 	r.Run(":8080")
> }
> 
> ```
>
> 在上面的示例代码中，`NoRoute()` 方法设置了一个处理函数，该函数返回 404 Not Found 的 JSON 格式数据。当用户访问未定义的路由时，Gin 框架会调用该处理函数，返回 404 状态码和相应的 JSON 数据。而当用户访问 `/hello` 路由时，则会调用该路由的处理函数，返回 Hello Gin! 的 JSON 数据。

#### (*Engine) Routes 

``` go linenums="1"
func (engine *Engine) Routes() (routes RoutesInfo)
```

​	Routes方法返回已注册路由的切片，包括一些有用的信息，如 HTTP 方法、路径和处理程序名称。

#### (*Engine) Run 

``` go linenums="1"
func (engine *Engine) Run(addr ...string) (err error)
```

​	Run方法将路由器附加到 http.Server 上并开始侦听和服务 HTTP 请求。这是 http.ListenAndServe(addr, router) 的快捷方式。注意：除非出现错误，否则此方法将无限期地阻塞调用 goroutine。

#### (*Engine) RunFd 

``` go linenums="1"
func (engine *Engine) RunFd(fd int) (err error)
```

​	RunFd方法将路由器附加到 http.Server 上并开始通过指定的文件描述符侦听和服务 HTTP 请求。注意：除非出现错误，否则此方法将无限期地阻塞调用 goroutine。

#### (*Engine) RunListener 

``` go linenums="1"
func (engine *Engine) RunListener(listener net.Listener) (err error)
```

​	RunListener方法将路由器附加到 http.Server 上并开始通过指定的 net.Listener 侦听和服务 HTTP 请求。

#### (*Engine) RunTLS 

``` go linenums="1"
func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error)
```

​	RunTLS方法将路由器附加到 http.Server 上并开始侦听和服务 HTTPS（安全）请求。这是 http.ListenAndServeTLS(addr, certFile, keyFile, router) 的快捷方式。注意：除非出现错误，否则此方法将无限期地阻塞调用 goroutine。

#### (*Engine) RunUnix 

``` go linenums="1"
func (engine *Engine) RunUnix(file string) (err error)
```

​	RunUnix方法将路由器附加到 http.Server 上并开始通过指定的 Unix 套接字（即文件）侦听和服务 HTTP 请求。注意：除非出现错误，否则此方法将无限期地阻塞调用 goroutine。

#### (*Engine) SecureJsonPrefix 

``` go linenums="1"
func (engine *Engine) SecureJsonPrefix(prefix string) *Engine
```

​	SecureJsonPrefix方法设置在 Context.SecureJSON()方法中使用的 secureJSONPrefix。

#### (*Engine) ServeHTTP 

``` go linenums="1"
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)
```

ServeHTTP方法符合 http.Handler 接口。

> `Engine.ServeHTTP()` 方法是 `gin` 框架的核心方法，它实现了 `http.Handler` 接口，可以处理来自客户端的 HTTP 请求。
>
> 当请求到达服务器时，HTTP Server 会调用 `Engine.ServeHTTP()` 方法，`Engine` 会根据请求的 URL 和 HTTP 方法选择合适的路由进行处理，并将请求交给对应的中间件和处理函数来处理。
>
> 下面是一个简单的示例：
>
> ```go linenums="1"
> package main
> 
> import (
>     "github.com/gin-gonic/gin"
>     "net/http"
> )
> 
> func main() {
>     // 创建一个新的 gin 引擎实例
>     r := gin.Default()
> 
>     // 添加路由和对应的处理函数
>     r.GET("/", func(c *gin.Context) {
>         c.String(http.StatusOK, "Hello, World!")
>     })
> 
>     // 将 gin 引擎实例作为 http.Handler 注册到 HTTP Server
>     http.ListenAndServe(":8080", r)
> }
> 
> ```
>
> ​	在上面的示例中，我们首先创建了一个 `gin` 引擎实例 `r`，然后通过 `r.GET()` 方法为根路由 `/` 添加了一个 GET 请求处理函数。最后，我们将 `r` 注册到 `http.ListenAndServe()` 方法中，该方法会创建一个 HTTP Server 并监听端口 `8080`，并将每个请求交给 `r` 进行处理。
>
> ​	当请求到达服务器时，HTTP Server 会调用 `r.ServeHTTP()` 方法，该方法会将请求交给 `gin` 引擎实例 `r` 进行处理，`r` 会根据请求的 URL 和 HTTP 方法选择合适的路由进行处理，并将请求交给对应的中间件和处理函数来处理。在本例中，当用户访问根路由时，`r` 会将请求交给我们定义的处理函数，处理函数会返回一个包含 `Hello, World!` 字符串的响应。最后，`http.ListenAndServe()` 方法会将响应发送给客户端。

#### (*Engine) SetFuncMap 

``` go linenums="1"
func (engine *Engine) SetFuncMap(funcMap template.FuncMap)
```

​	SetFuncMap方法设置用于 template.FuncMap 的 FuncMap。

#### (*Engine) SetHTMLTemplate 

``` go linenums="1"
func (engine *Engine) SetHTMLTemplate(templ *template.Template)
```

​	SetHTMLTemplate方法将模板与 HTML 渲染器关联。

#### (*Engine) SetTrustedProxies 

``` go linenums="1"
func (engine *Engine) SetTrustedProxies(trustedProxies []string) error
```

​	SetTrustedProxies方法设置一个网络来源列表（IPv4 地址、IPv4 CIDR、IPv6 地址或 IPv6 CIDR），当 `(*gin.Engine).ForwardedByClientIP` 为 `true` 时，用于信任包含备用客户端 IP 的请求标头。`TrustedProxies` 功能默认已启用，并且默认情况下还信任所有代理。如果要禁用此功能，请使用 Engine.SetTrustedProxies(nil)，然后 Context.ClientIP() 将直接返回远程地址。

#### (*Engine) Use 

``` go linenums="1"
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes
```

​	Use 方法将全局中间件附加到路由器上，即通过 Use() 附加的中间件将包含在每个单个请求的处理程序链中。甚至包括404、405、静态文件等请求。例如，这是一个记录器或错误管理中间件的正确位置。

### type Error 

``` go linenums="1"
type Error struct {
	Err  error
	Type ErrorType
	Meta any
}
```

​	Error 表示一个错误的规范。 

#### (Error) Error 

``` go linenums="1"
func (msg Error) Error() string
```

​	Error方法实现了error 接口。

#### (*Error) IsType 

``` go linenums="1"
func (msg *Error) IsType(flags ErrorType) bool
```

​	IsType方法判断一个错误。

#### (*Error) JSON 

``` go linenums="1"
func (msg *Error) JSON() any
```

​	JSON方法创建一个正确格式的JSON。

#### (*Error) MarshalJSON 

``` go linenums="1"
func (msg *Error) MarshalJSON() ([]byte, error)
```

​	MarshalJSON方法实现了json.Marshaller接口。

#### (*Error) SetMeta 

``` go linenums="1"
func (msg *Error) SetMeta(data any) *Error
```

​	SetMeta方法设置错误的元数据。

#### (*Error) SetType 

``` go linenums="1"
func (msg *Error) SetType(flags ErrorType) *Error
```

​	SetType方法设置错误的类型。

#### (*Error) Unwrap 

``` go linenums="1"
func (msg *Error) Unwrap() error
```

​	Unwrap方法返回被包装的错误，以便与errors.Is()、errors.As()和errors.Unwrap()互操作。

### type ErrorType 

``` go linenums="1"
type ErrorType uint64
```

​	ErrorType类型是一个无符号的64位错误代码，在gin规范中定义。

``` go linenums="1"
const (
	// ErrorTypeBind 表示当 Context.Bind() 失败时使用的错误类型。
	ErrorTypeBind ErrorType = 1 << 63
	// ErrorTypeRender 表示当 Context.Render() 失败时使用的错误类型。
	ErrorTypeRender ErrorType = 1 << 62
	// ErrorTypePrivate 表示一个私有错误。
	ErrorTypePrivate ErrorType = 1 << 0
	// ErrorTypePublic 表示一个公共错误。
	ErrorTypePublic ErrorType = 1 << 1
	// ErrorTypeAny 表示任何其他类型的错误。
	ErrorTypeAny ErrorType = 1<<64 - 1
	// ErrorTypeNu 表示任何其他类型的错误。
	ErrorTypeNu = 2
)
```

### type H 

``` go linenums="1"
type H map[string]any
```

​	H 是 map[string]interface{}类型的简称（别名）。 

#### (H) MarshalXML 

``` go linenums="1"
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```

​	MarshalXML方法允许H类型与xml.Marshal方法一起使用。

### type HandlerFunc 

``` go linenums="1"
type HandlerFunc func(*Context)
```

​	HandlerFunc将 gin 中间件使用的处理程序定义为返回值。 

#### func BasicAuth 

``` go linenums="1"
func BasicAuth(accounts Accounts) HandlerFunc
```

​	BasicAuth函数返回一个基本的 HTTP 授权中间件，它的参数是一个 map[string]string 类型，其中 key 是用户名，value 是密码。

#### func BasicAuthForRealm 

``` go linenums="1"
func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc
```

​	BasicAuthForRealm函数返回一个基本的 HTTP 授权中间件，它的参数是一个 map[string]string 类型，其中 key 是用户名，value 是密码，另外还需要指定 Realm 的名称。如果 Realm 是空的，则默认使用 "Authorization Required"。（参见 [http://tools.ietf.org/html/rfc2617#section-1.2](http://tools.ietf.org/html/rfc2617#section-1.2)）

#### func Bind 

``` go linenums="1"
func Bind(val any) HandlerFunc
```

​	Bind函数是一个给定接口对象的辅助函数，并返回一个Gin中间件。

#### func CustomRecovery 

``` go linenums="1"
func CustomRecovery(handle RecoveryFunc) HandlerFunc
```

​	CustomRecovery函数返回一个中间件，可以从任何panic中恢复并调用提供的处理函数来处理它。

#### func CustomRecoveryWithWriter 

``` go linenums="1"
func CustomRecoveryWithWriter(out io.Writer, handle RecoveryFunc) HandlerFunc
```

​	CustomRecoveryWithWriter函数返回一个使用指定 writer 的中间件，用于从任何 panic 中恢复，并调用提供的 handle 函数来处理 panic。

#### func ErrorLogger 

``` go linenums="1"
func ErrorLogger() HandlerFunc
```

​	ErrorLogger函数返回一个适用于任何错误类型的 HandlerFunc。

#### func ErrorLoggerT 

``` go linenums="1"
func ErrorLoggerT(typ ErrorType) HandlerFunc
```

​	ErrorLoggerT函数返回适用于给定错误类型的 HandlerFunc。

#### func Logger 

``` go linenums="1"
func Logger() HandlerFunc
```

​	Logger函数实例化一个 Logger 中间件，它将日志写入 gin.DefaultWriter。默认情况下，gin.DefaultWriter = os.Stdout。

#### func LoggerWithConfig 

``` go linenums="1"
func LoggerWithConfig(conf LoggerConfig) HandlerFunc
```

​	LoggerWithConfig函数实例化一个具有指定配置的 Logger 中间件。

#### func LoggerWithFormatter 

``` go linenums="1"
func LoggerWithFormatter(f LogFormatter) HandlerFunc
```

​	LoggerWithFormatte函数实例化一个具有指定日志格式函数的 Logger 中间件。

#### func LoggerWithWriter 

``` go linenums="1"
func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc
```

​	LoggerWithWriter函数实例化一个具有指定 writer 缓冲区的 Logger 中间件。例如：os.Stdout、以写模式打开的文件、套接字等等。

#### func Recovery 

``` go linenums="1"
func Recovery() HandlerFunc
```

​	Recovery函数返回一个从任何 panic 中恢复的中间件，并在有 panic 时写入 500。

#### func RecoveryWithWriter 

``` go linenums="1"
func RecoveryWithWriter(out io.Writer, recovery ...RecoveryFunc) HandlerFunc
```

​	RecoveryWithWriter函数返回一个使用指定 writer 的中间件，用于从任何 panic 中恢复，并在有 panic 时写入 500。

#### func WrapF 

``` go linenums="1"
func WrapF(f http.HandlerFunc) HandlerFunc
```

​	WrapF函数是一个辅助函数，用于将 http.HandlerFunc 包装为 Gin 中间件。

#### func WrapH 

``` go linenums="1"
func WrapH(h http.Handler) HandlerFunc
```

​	WrapH函数是一个辅助函数，用于将 http.Handler 包装为 Gin 中间件。

### type HandlersChain 

``` go linenums="1"
type HandlersChain []HandlerFunc
```

​	HandlersChain类型定义了一个HandlerFunc切片。

#### (HandlersChain) Last 

``` go linenums="1"
func (c HandlersChain) Last() HandlerFunc
```

​	Last方法返回链中的最后一个处理程序，即最后一个处理程序是主要处理程序。

### type IRouter 

``` go linenums="1"
type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}
```

​	IRouter 定义了包括单个和组路由在内的所有路由器处理接口。 

### type IRoutes 

``` go linenums="1"
type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
	DELETE(string, ...HandlerFunc) IRoutes
	PATCH(string, ...HandlerFunc) IRoutes
	PUT(string, ...HandlerFunc) IRoutes
	OPTIONS(string, ...HandlerFunc) IRoutes
	HEAD(string, ...HandlerFunc) IRoutes
	Match([]string, string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}
```

​	IRoutes定义了所有路由处理接口。

### type LogFormatter 

``` go linenums="1"
type LogFormatter func(params LogFormatterParams) string
```

​	LogFormatter类型给出了传递给 LoggerWithFormatter 的格式化函数的签名。 

### type LogFormatterParams 

``` go linenums="1"
type LogFormatterParams struct {
	Request *http.Request

	// TimeStamp 展示服务器返回响应的时间。
	TimeStamp time.Time
   // StatusCode 是 HTTP 响应码。
	StatusCode int
   // Latency 是服务器处理某个请求所需的时间。
	Latency time.Duration
   // ClientIP 等于 Context 的 ClientIP 方法。
	ClientIP string
   // Method 是请求所使用的 HTTP 方法。
	Method string
   // Path 是客户端请求的路径。
	Path string
   // ErrorMessage 如果在处理请求时发生错误，则设置此字段。
	ErrorMessage string

   // BodySize 是响应体的大小。
	BodySize int
   // Keys 是请求上下文中设置的键。
	Keys map[string]any
	// contains filtered or unexported fields
}
```

​	LogFormatterParams是任何格式化程序在记录日志时将要使用的结构体。 

#### (*LogFormatterParams) IsOutputColor 

``` go linenums="1"
func (p *LogFormatterParams) IsOutputColor() bool
```

​	IsOutputColor方法表示是否可以在日志中输出颜色。

#### (*LogFormatterParams) MethodColor 

``` go linenums="1"
func (p *LogFormatterParams) MethodColor() string
```

​	MethodColor方法是用于在终端中适当记录HTTP方法的ANSI颜色。。

#### (*LogFormatterParams) ResetColor 

``` go linenums="1"
func (p *LogFormatterParams) ResetColor() string
```

​	ResetColor方法重置所有转义属性。

#### (*LogFormatterParams) StatusCodeColor 

``` go linenums="1"
func (p *LogFormatterParams) StatusCodeColor() string
```

​	StatusCodeColor方法是用于在终端上适当记录HTTP状态码的ANSI颜色。

### type LoggerConfig 

``` go linenums="1"
type LoggerConfig struct {
   // 可选项。默认值为gin.defaultLogFormatter。
	Formatter LogFormatter

	// Output是日志写入的输出流。
	// 可选项。默认值为gin.DefaultWriter。
	Output io.Writer

	// SkipPaths是一个URL路径的数组，日志不会被写入。
	// 可选项。
	SkipPaths []string
}
```

​	LoggerConfig结构体定义了 Logger 中间件的配置。

### type Negotiate 

``` go linenums="1"
type Negotiate struct {
	Offered  []string
	HTMLName string
	HTMLData any
	JSONData any
	XMLData  any
	YAMLData any
	Data     any
	TOMLData any
}
```

​	Negotiate结构体包含所有协商数据。

### type Param 

``` go linenums="1"
type Param struct {
	Key   string
	Value string
}
```

​	Param结构体是一个 URL 参数，由键和值组成。 

### type Params 

``` go linenums="1"
type Params []Param
```

​	Params 是一个 Param 切片，由路由器返回。该切片是有序的，第一个 URL 参数也是第一个切片值。因此，可以通过索引安全地读取值。

#### (Params) ByName 

``` go linenums="1"
func (ps Params) ByName(name string) (va string)
```

​	ByName方法返回与给定名称匹配的第一个Param的值。如果找不到匹配的Param，则返回空字符串。

#### (Params) Get 

``` go linenums="1"
func (ps Params) Get(name string) (string, bool)
```

​	Get方法返回与给定名称匹配的第一个Param的值和一个布尔值true。如果找不到匹配的Param，则返回空字符串和布尔值false。

### type RecoveryFunc 

``` go linenums="1"
type RecoveryFunc func(c *Context, err any)
```

​	RecoveryFunc类型定义了可传递给 CustomRecovery函数的函数。 

### type ResponseWriter 

``` go linenums="1"
type ResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier

	// Status 方法返回当前请求的 HTTP 响应状态码。
	Status() int

	// Size 方法返回已经写入 HTTP 响应主体的字节数。参见 Written() 方法。
	Size() int

	// WriteString 方法向响应主体写入字符串。
	WriteString(string) (int, error)

	// 如果 HTTP 响应主体已经写入，Written 方法返回 true。
	Written() bool

	// WriteHeaderNow forces to write the http header (status code + headers).
   // WriteHeaderNow 方法强制写入 HTTP 头部（状态码 + 头部）
	WriteHeaderNow()

	// Pusher 方法获取 http.Pusher 接口，以便进行服务端推送。
	Pusher() http.Pusher
}
```

ResponseWriter ...

### type RouteInfo 

``` go linenums="1"
type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}
```

​	RouteInfo 表示请求路由的规范，包含请求方法、路径和对应的处理函数。

### type RouterGroup 

``` go linenums="1"
type RouterGroup struct {
	Handlers HandlersChain
	// contains filtered or unexported fields
}
```

​		RouterGroup结构体用于内部配置路由器，每个 RouterGroup 关联一个前缀和一组处理函数（中间件）。

#### (*RouterGroup) Any 

``` go linenums="1"
func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes
```

​	Any方法注册一个匹配所有HTTP方法的路由。GET、POST、PUT、PATCH、HEAD、OPTIONS、DELETE、CONNECT、TRACE。

#### (*RouterGroup) BasePath 

``` go linenums="1"
func (group *RouterGroup) BasePath() string
```

​	BasePath方法返回路由组的基础路径。例如，如果v:=router.Group("/rest/n/v1/api")，v.BasePath()将返回"/rest/n/v1/api"。

#### (*RouterGroup) DELETE 

``` go linenums="1"
func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes
```

​	DELETE方法是router.Handle("DELETE", path, handlers)的快捷方式。

#### (*RouterGroup) GET 

``` go linenums="1"
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes
```

​	GET方法是router.Handle("GET", path, handlers)的快捷方式。

#### (*RouterGroup) Group 

``` go linenums="1"
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup
```

​	Group方法创建一个新的路由组。您应该添加所有具有常见中间件或相同路径前缀的路由。例如，所有使用公共授权中间件的路由都可以被分组。

#### (*RouterGroup) HEAD 

``` go linenums="1"
func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes
```

​	HEAD方法是router.Handle("HEAD", path, handlers)的快捷方式。

#### (*RouterGroup) Handle 

``` go linenums="1"
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes
```

​	Handle方法注册一个具有给定路径和method的新请求处理和中间件。最后一个处理程序应该是真正的处理程序，其他的应该是可以和不同路由共享的中间件。请参见GitHub中的示例代码。

​	对于GET、POST、PUT、PATCH和DELETE请求，可以使用相应的快捷方法。

​	此方法用于批量加载和允许使用不太频繁使用的、非标准化的或自定义方法（例如，用于与代理的内部通信）。

#### (*RouterGroup) Match 

``` go linenums="1"
func (group *RouterGroup) Match(methods []string, relativePath string, handlers ...HandlerFunc) IRoutes
```

​	Match方法方法注册一个匹配指定method的路由。

#### (*RouterGroup) OPTIONS 

``` go linenums="1"
func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes
```

​	OPTIONS方法是router.Handle("OPTIONS", path, handlers)的快捷方式。

#### (*RouterGroup) PATCH 

``` go linenums="1"
func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes
```

​	PATCH方法是router.Handle("PATCH", path, handlers)的快捷方式。

#### (*RouterGroup) POST 

``` go linenums="1"
func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes
```

​	POST方法是router.Handle("POST", path, handlers)的快捷方式。

#### (*RouterGroup) PUT 

``` go linenums="1"
func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes
```

​	PUT方法是router.Handle("PUT", path, handlers)的快捷方式。

#### (*RouterGroup) Static 

``` go linenums="1"
func (group *RouterGroup) Static(relativePath, root string) IRoutes
```

​	Static方法从给定的文件系统根目录提供文件。内部使用http.FileServer，因此使用http.NotFound而不是路由器的NotFound处理程序。要使用操作系统的文件系统实现，请使用：

```
router.Static("/static", "/var/www")
```

#### (*RouterGroup) StaticFS 

``` go linenums="1"
func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes
```

​	StaticFS方法的用法与Static()相同，但可以使用自定义的http.FileSystem。Gin默认使用gin.Dir()。

#### (*RouterGroup) StaticFile 

``` go linenums="1"
func (group *RouterGroup) StaticFile(relativePath, filepath string) IRoutes
```

​	StaticFile方法会注册一个路由，以便于用于服务本地文件系统中的单个文件。例如：`router.StaticFile("favicon.ico", "./resources/favicon.ico")`。

#### (*RouterGroup) StaticFileFS 

``` go linenums="1"
func (group *RouterGroup) StaticFileFS(relativePath, filepath string, fs http.FileSystem) IRoutes
```

​	StaticFileFS方法的用法与StaticFile方法相同，但可以使用自定义的http.FileSystem。例如：`router.StaticFileFS("favicon.ico", "./resources/favicon.ico", Dir{".", false})`。 Gin默认使用`gin.Dir()`。

#### (*RouterGroup) Use 

``` go linenums="1"
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes
```

​	Use方法向组添加中间件，参见GitHub中的示例代码。

### type RoutesInfo 

``` go linenums="1"
type RoutesInfo []RouteInfo
```

​	RoutesInfo类型定义了一个 RouteInfo 切片。