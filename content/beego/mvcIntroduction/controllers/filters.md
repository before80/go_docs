+++
title = "过滤器"
date = 2024-02-04T09:57:22+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/filter/]({{< ref "/beego/mvcIntroduction/controllers/filters" >}})

# Filters 过滤器



## Filters 过滤器

Beego supports custom filter middlewares. E.g.: user authentication and force redirection.

​	Beego 支持自定义过滤器中间件。例如：用户认证和强制重定向。

## Activating Filters 激活过滤器

Before filters can be used, filters must be activated.

​	在使用过滤器之前，必须先激活过滤器。

Filters can be activated at the code level:

​	可以在代码级别激活过滤器：

```
web.BConfig.WebConfig.Session.SessionOn = true
```

Filters can also be activated in the configuration file:

​	也可以在配置文件中激活过滤器：

```
SessionOn = true
```

Attempting to use a filter without activation will cause a `Handler crashed with error runtime error: invalid memory address or nil pointer dereference` error

​	尝试在未激活的情况下使用过滤器将导致 `Handler crashed with error runtime error: invalid memory address or nil pointer dereference` 错误

## Inserting Filters 插入过滤器

A filter function can be inserted as follows:

​	可以按如下方式插入过滤器函数：

```go
web.InsertFilter(pattern string, pos int, filter FilterFunc, opts ...FilterOpt)
```

This is the FilterFunc signature:

​	这是 FilterFunc 签名：

```go
type FilterFunc func(*context.Context)
```

The *context* must be imported if this has not already been done:

​	如果尚未导入上下文，则必须导入上下文：

```go
import "github.com/beego/beego/v2/server/web/context"
```

InsertFilter’s four parameters:

​	InsertFilter 的四个参数：

- `pattern`: string or regex to match against router rules. Use `/*` to match all.
  `pattern` ：用于与路由规则匹配的字符串或正则表达式。使用 `/*` 匹配所有内容。

- ```
  pos
  ```

  : the place to execute the Filter. There are five fixed parameters representing different execution processes.

  
  `pos` ：用于运行过滤器的场所。有五个固定变量表示不同的运行进程。

  - web.BeforeStatic: Before finding the static file.
    web.BeforeRouter：在查找路由器前。
  - web.BeforeRouter: Before finding router.
    web.BeforeRouter：在查找路由器前。
  - web.BeforeExec: After finding router and before executing the matched Controller.
    web.BeforeRouter：在查找路由器前。
  - web.AfterExec: After executing Controller.
    web.BeforeRouter：在查找路由器前。
  - web.FinishRouter: After finishing router.
    web.BeforeRouter：在查找路由器前。

- `filter`: filter function type FilterFunc func(*context.Context)
  `filter` ：过滤器 func 类型Filterfunc func(*context.Context)

- ```
  opts
  ```

  :

  1. web.WithReturnOnOutput: whether to continue running if has output. default is false.
     web.WithReturnOnOutput：如果已输出，则确定是停止还是继运行。默认值为 false。
  2. web.WithResetParams: whether to reset parameters to their previous values after the filter has completed.
     web.WithRestParams：在过滤器运行完成后，是将变量重置为其先前的值。
  3. web.WithCaseSensitive: whether case sensitive
     web.WithCaseSensitive：是否区分大小写

> from beego version 1.3 AddFilter has been removed
>
> ​	从 beego 版本 1.3 开始，已移除 AddFilter

Here is an example to authenticate if the user is logged in for all requests:

​	以下是一个示例，用于验证用户是否已登录所有请求：

```go
var FilterUser = func(ctx *context.Context) {
    if strings.HasPrefix(ctx.Input.URL(), "/login") {
    	return
    }
    
    _, ok := ctx.Input.Session("uid").(int)
    if !ok {
        ctx.Redirect(302, "/login")
    }
}

web.InsertFilter("/*", web.BeforeRouter, FilterUser)
```

> Filters which use session must be executed after `BeforeRouter` because session is not initialized before that. web session module must be enabled first. (see [Session control]({{< ref "/beego/mvcIntroduction/controllers/sessionControl" >}}))
>
> ​	使用会话的过滤器必须在 `BeforeRouter` 之后执行，因为在此之前不会初始化会话。必须首先启用 web 会话模块。（请参阅会话控制）

Filters can be run against requests which use a regex router rule for matching:

​	可以使用正则表达式路由规则进行匹配来对请求运行过滤器：

```go
var FilterUser = func(ctx *context.Context) {
    _, ok := ctx.Input.Session("uid").(int)
    if !ok {
        ctx.Redirect(302, "/login")
    }
}
web.InsertFilter("/user/:id([0-9]+)", web.BeforeRouter, FilterUser)
```

## Filter Implementation UrlManager 过滤器实现 UrlManager

Context.Input has new features `RunController` and `RunMethod` from beego version 1.1.2. These can control the router in the filter and skip the Beego router rule.

​	Context.Input 从 beego 版本 1.1.2 开始具有新功能 `RunController` 和 `RunMethod` 。这些功能可以在过滤器中控制路由器并跳过 Beego 路由规则。

For example:

​	例如：

```go
var UrlManager = func(ctx *context.Context) {
    // read urlMapping data from database
	urlMapping := model.GetUrlMapping()
	for baseurl,rule := range urlMapping {
		if baseurl == ctx.Request.RequestURI {
			ctx.Input.RunController = rule.controller
			ctx.Input.RunMethod = rule.method
			break
		}
	}
}

web.InsertFilter("/*", web.BeforeRouter, UrlManager)
```

## Filter和FilterChain 过滤器和 FilterChain

In v1.x, we can’t invoke next `Filter` inside a `Filter`. So we got a problem: we could not do something “surrounding” request execution.

​	在 v1.x 中，我们无法在 `Filter` 内调用下一个 `Filter` 。因此我们遇到了一个问题：我们无法对请求执行进行“环绕”。

For example, if we want to do:

​	例如，如果我们想执行以下操作：

```
func filter() {
    // do something before serving request
    handleRequest()
    // do something after serving request
}
```

The typical cases are tracing and metrics.

​	典型的案例是跟踪和指标。

So we enhance `Filter` by designing a new interface:

​	因此，我们通过设计一个新接口来增强 `Filter` ：

```go
type FilterChain func(next FilterFunc) FilterFunc
```

Here is a simple example:

​	这是一个简单的示例：

```go
package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func main() {
	web.InsertFilterChain("/*", func(next web.FilterFunc) web.FilterFunc {
		return func(ctx *context.Context) {
			// do something
			logs.Info("hello")
			// don't forget this
			next(ctx)

			// do something
		}
	})
}
```

In this example, we only output “hello” and then we invoke next filter.

​	在此示例中，我们仅输出“hello”，然后调用下一个过滤器。

### Prometheus例子 Prometheus 示例

```go
package main

import (
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/prometheus"
)

func main() {
	// we start admin service
	// Prometheus will fetch metrics data from admin service's port
	web.BConfig.Listen.EnableAdmin = true

	web.BConfig.AppName = "my app"

	ctrl := &MainController{}
	web.Router("/hello", ctrl, "get:Hello")
	fb := &prometheus.FilterChainBuilder{}
	web.InsertFilterChain("/*", fb.FilterChain)
	web.Run(":8080")
	// after you start the server
	// and GET http://localhost:8080/hello
	// access http://localhost:8088/metrics
	// you can see something looks like:
	// http_request_web_sum{appname="my app",duration="1002",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1002
	// http_request_web_count{appname="my app",duration="1002",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1
	// http_request_web_sum{appname="my app",duration="1004",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1004
	// http_request_web_count{appname="my app",duration="1004",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) Hello() {
	time.Sleep(time.Second)
	ctrl.Ctx.ResponseWriter.Write([]byte("Hello, world"))
}
```

If you don’t use Beego’s admin service, don’t forget to expose `prometheus`’s port.

​	如果您不使用 Beego 的管理服务，请不要忘记公开 `prometheus` 的端口。

### Opentracing例子 Opentracing 示例

```go
package main

import (
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/opentracing"
)

func main() {
	// don't forget this to inject the opentracing API's implementation
	// opentracing2.SetGlobalTracer()

	web.BConfig.AppName = "my app"

	ctrl := &MainController{}
	web.Router("/hello", ctrl, "get:Hello")
	fb := &opentracing.FilterChainBuilder{}
	web.InsertFilterChain("/*", fb.FilterChain)
	web.Run(":8080")
	// after you start the server
}

type MainController struct {
	web.Controller
}

func (ctrl *MainController) Hello() {
	time.Sleep(time.Second)
	ctrl.Ctx.ResponseWriter.Write([]byte("Hello, world"))
}
```

Don’t forget to using `SetGlobalTracer` to initialize opentracing.

​	不要忘记使用 `SetGlobalTracer` 初始化 opentracing。
