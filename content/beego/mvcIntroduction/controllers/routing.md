+++
title = "路由"
date = 2024-02-04T09:56:45+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/router/]({{< ref "/beego/mvcIntroduction/controllers/routing" >}})

# Routing 路由



## Routing 路由

This chapter will cover the three types of routers incorporated into Beego.

​	本章将介绍 Beego 中包含的三种类型的路由器。

## Basic router 基本路由器

Beego supports a RESTful function router. This basic router includes the URI and closure functions.

​	Beego 支持 RESTful 函数路由器。此基本路由器包括 URI 和闭包函数。

### GET router GET 路由器

```
web.Get("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world"))
})
```

### POST router POST 路由器

```
web.Post("/alice",func(ctx *context.Context){
     ctx.Output.Body([]byte("bob"))
})
```

### support all HTTP routers 支持所有 HTTP 路由器

```
web.Any("/foo",func(ctx *context.Context){
     ctx.Output.Body([]byte("bar"))
})
```

### all the functions 所有功能

- web.Get(router, web.FilterFunc)
- web.Post(router, web.FilterFunc)
- web.Put(router, web.FilterFunc)
- web.Head(router, web.FilterFunc)
- web.Options(router, web.FilterFunc)
- web.Delete(router, web.FilterFunc)
- web.Any(router, web.FilterFunc)

### Handler register 处理程序注册

In cases where packages such as `net/http` are already implemented in a system they can be integrated into the web API or web system by following this procedure:

​	在系统中已经实现了诸如 `net/http` 之类的包的情况下，可以通过遵循此过程将其集成到 Web API 或 Web 系统中：

```
s := rpc.NewServer()
s.RegisterCodec(json.NewCodec(), "application/json")
s.RegisterService(new(HelloService), "")
web.Handler("/rpc", s)
```

`beego.Handler(router, http.Handler)` the first parameter represents the URI, and the second parameter represents `http.Handler`. When this is registered all requests to `/rpc` will call `http.Handler`.

​	 `beego.Handler(router, http.Handler)` 第一个参数表示 URI，第二个参数表示 `http.Handler` 。注册此项后，对 `/rpc` 的所有请求都将调用 `http.Handler` 。

There is also a third parameter, `isPrefix`. If this parameter is set to `true` all the matches will comply with prefix matching, meaning that the url `/rpc/user` will also call the register. By default this value is `false`.

​	还有一个第三个参数 `isPrefix` 。如果将此参数设置为 `true` ，则所有匹配项都将符合前缀匹配，这意味着 url `/rpc/user` 也将调用注册。默认情况下，此值为 `false` 。

## RESTful router RESTful 路由器

RESTful is a popular approach to API development that Beego supports implicitly, executing `Get` method for GET request and `Post` method for POST request. The default router is RESTful.

​	RESTful 是 Beego 隐式支持的 API 开发的流行方法，它为 GET 请求执行 `Get` 方法，为 POST 请求执行 `Post` 方法。默认路由器是 RESTful。

## Fixed router 固定路由器

A fixed router is a full matching router, such as:

​	固定路由器是完全匹配路由器，例如：

```
web.Router("/", &controllers.MainController{})
web.Router("/admin", &admin.UserController{})
web.Router("/admin/index", &admin.ArticleController{})
web.Router("/admin/addpkg", &admin.AddController{})
```

The fixed routers above are typical RESTful routers in their most common configuration, with one fixed router and one controller. This results in the execution of a different method based on each request method.

​	上面的固定路由器是其最常见配置中的典型 RESTful 路由器，具有一个固定路由器和一个控制器。这会导致根据每个请求方法执行不同的方法。

## Regex router 正则路由器

To simplify router configuration, Beego uses the router implementation approach found in Sinatra to support many router types.

​	为了简化路由器配置，Beego 使用 Sinatra 中的路由器实现方法来支持多种路由器类型。

- web.Router("/api/?:id", &controllers.RController{})

  *default matching* /api/123 :id = 123 *can match* /api/

  ​	默认匹配 /api/123 :id = 123 可以匹配 /api/

- web.Router("/api/:id", &controllers.RController{})

  *default matching* /api/123 :id = 123 *can’t match* /api/

  ​	默认匹配 /api/123 :id = 123 无法匹配 /api/

- web.Router("/api/:id([0-9]+)", &controllers.RController{})

  *Customized regex matching* /api/123 :id = 123

  ​	自定义正则匹配 /api/123 :id = 123

- web.Router("/user/:username([\w]+)", &controllers.RController{})

  *Regex string matching* /user/astaxie :username = astaxie

  ​	正则字符串匹配 /user/astaxie :username = astaxie

- web.Router("/download/*.*", &controllers.RController{})

  *matching* /download/file/api.xml :path= file/api :ext=xml

  ​	匹配 /download/file/api.xml :path= file/api :ext=xml

- web.Router("/download/ceshi/*", &controllers.RController{})

  *full matching* /download/ceshi/file/api.json :splat=file/api.json

  ​	完全匹配 /download/ceshi/file/api.json :splat=file/api.json

- web.Router("/:id:int", &controllers.RController{})

  *int type matching* :id is int type. web implements ([0-9]+) for you

  ​	int 类型匹配：:id 是 int 类型。web 为您实现 ([0-9]+)

- web.Router("/:hello:string", &controllers.RController{})

  *string type matching* :hello is string type. web implements ([\w]+) for you

  ​	string 类型匹配：:hello 是 string 类型。web 为您实现 ([\w]+)

- beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{})

  *has prefix regex* :id is the regex. *matching* cms_123.html :id = 123

  ​	具有前缀正则表达式：id 是正则表达式。匹配 cms_123.html：id = 123

The variables can be accessed in the controller like this:

​	可以在控制器中这样访问变量：

```
this.Ctx.Input.Param(":id")
this.Ctx.Input.Param(":username")
this.Ctx.Input.Param(":splat")
this.Ctx.Input.Param(":path")
this.Ctx.Input.Param(":ext")
```

## Custom methods and RESTful rules 自定义方法和 RESTful 规则

The examples above use default method names, where the request method name is same as the controller method name. For example as `GET` request executes `Get` method and `POST` request executes `Post` method. Different controller method names can be set like this:

​	上面的示例使用默认方法名称，其中请求方法名称与控制器方法名称相同。例如， `GET` 请求执行 `Get` 方法， `POST` 请求执行 `Post` 方法。可以这样设置不同的控制器方法名称：

```
web.Router("/",&IndexController{},"*:Index")
```

Use the third parameter which is the method you want to call in the controller. Here are some rules:

​	使用第三个参数，即要在控制器中调用的方法。以下是一些规则：

- - means any request method will execute this method.
    表示任何请求方法都将执行此方法。
- Use httpmethod:funcname format.
  使用 httpmethod:funcname 格式。
- Multiple formats can use `;` as the separator.
  多种格式可以使用 `;` 作为分隔符。
- Many HTTP methods mapping the same funcname, use `,` as the separator for HTTP methods.
  多个 HTTP 方法映射到同一个 funcname，使用 `,` 作为 HTTP 方法的分隔符。

Below are some examples of RESTful design:

​	以下是 RESTful 设计的一些示例：

```
web.Router("/api/list",&RestController{},"*:ListFood")
web.Router("/api/create",&RestController{},"post:CreateFood")
web.Router("/api/update",&RestController{},"put:UpdateFood")
web.Router("/api/delete",&RestController{},"delete:DeleteFood")
```

Below is an example of multiple HTTP methods mapping to the same controller method:

​	下面是一个将多个 HTTP 方法映射到同一个控制器方法的示例：

```
web.Router("/api",&RestController{},"get,post:ApiFunc")
```

Below is an example of different HTTP methods mapping to different controller methods. `;` as the separator:

​	以下是不同 HTTP 方法映射到不同控制器方法的示例。 `;` 作为分隔符：

```
web.Router("/simple",&SimpleController{},"get:GetFunc;post:PostFunc")
```

Below are the acceptable HTTP methods:

​	以下是可接受的 HTTP 方法：

- *：including all methods below
  *：包括以下所有方法
- get ：GET request
  get：GET 请求
- post ：POST request
  post：POST 请求
- put ：PUT request
  put：PUT 请求
- delete ：DELETE request
  delete：DELETE 请求
- patch ：PATCH request
  patch：PATCH 请求
- options ：OPTIONS request
  options：OPTIONS 请求
- head ：HEAD request
  head：HEAD 请求

If * and other HTTP methods are used together the HTTP method will be executed first. For example:

​	如果 * 和其他 HTTP 方法一起使用，则会首先执行 HTTP 方法。例如：

```
web.Router("/simple",&SimpleController{},"*:AllFunc;post:PostFunc")
```

The `PostFunc` rather than the `AllFunc` will be executed for POST requests.

​	对于 POST 请求，将执行 `PostFunc` 而不是 `AllFunc` 。

The router of custom methods does not support RESTful behaviour by default which means if you set the router like `web.Router("/api",&RestController{},"post:ApiFunc")` and the request method is `POST` then the `Post` method won’t be executed by default.

​	默认情况下，自定义方法的路由器不支持 RESTful 行为，这意味着如果您将路由器设置为 `web.Router("/api",&RestController{},"post:ApiFunc")` ，并且请求方法为 `POST` ，则默认情况下不会执行 `Post` 方法。

## Auto matching 自动匹配

To use auto matching the controller must be registered as an auto-router.

​	要使用自动匹配，控制器必须注册为自动路由器。

```
web.AutoRouter(&controllers.ObjectController{})
```

Beego will retrieve all the methods in that controller by reflection. The related methods can be called like this:

​	Beego 将通过反射检索该控制器中的所有方法。可以像这样调用相关方法：

```
/object/login   will call Login method of ObjectController
/object/logout  will call Logout method of ObjectController
```

Except `/:controller/:method` will match to controller and method. The remainder of the url path will be parsed as GET parameters and saved into `this.Ctx.Input.Param`:

​	除了 `/:controller/:method` 将匹配控制器和方法。其余的 URL 路径将被解析为 GET 参数并保存到 `this.Ctx.Input.Param` 中：

```
/object/blog/2013/09/12  will call Blog method of ObjectController with parameters `map[0:2013 1:09 2:12]`.
```

URL will match by lowercase conversion, so `object/LOGIN` will also map to `Login` method.

​	URL 将通过小写转换进行匹配，因此 `object/LOGIN` 也将映射到 `Login` 方法。

All the urls below will map to the `simple` method of `controller`.

​	以下所有 URL 都将映射到 `controller` 的 `simple` 方法。

```
/controller/simple
/controller/simple.html
/controller/simple.json
/controller/simple.xml
```

The extension name of the url can be reached by accessing `this.Ctx.Input.Param(":ext")`.

​	可以通过访问 `this.Ctx.Input.Param(":ext")` 来获取 URL 的扩展名。

## Annotations 注释

Not all routers need to be registered inside `router.go`. Only the controller needs to be registered using `Include`. For example:

​	并非所有路由器都需要在 `router.go` 中注册。仅控制器需要使用 `Include` 注册。例如：

```
// CMS API
type CMSController struct {
    web.Controller
}

func (c *CMSController) URLMapping() {
    c.Mapping("StaticBlock", c.StaticBlock)
    c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {

}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {
}
```

The routers can then be registered in `router.go`

​	然后可以将路由器注册在 `router.go` 中

```
web.Include(&CMSController{})
```

Beego will parse the source code automatically when under dev mode.

​	在开发模式下，Beego 会自动解析源代码。

The following routers will be supported:

​	将支持以下路由器：

- GET /staticblock/:key
- GET /all/:key

This is exactly same as registering by Router functions:

​	这与通过路由器函数注册完全相同：

```
web.Router("/staticblock/:key", &CMSController{}, "get:StaticBlock")
web.Router("/all/:key", &CMSController{}, "get:AllBlock")
```

If you do not use `URLMapping` Beego will find the function by reflection, otherwise Beego will find the function with the must faster `interface`.

​	如果您不使用 `URLMapping` ，Beego 将通过反射查找函数，否则 Beego 将使用更快的 `interface` 查找函数。

## Automatic Parameter Handling 自动参数处理

Beego supports automatic injection of http request parameters as method arguments, and method return values as http responses. For example, defining the following controller method:

​	Beego 支持将 http 请求参数自动注入为方法参数，并将方法返回值作为 http 响应。例如，定义以下控制器方法：

```
// @router /tasks/:id
func (c *TaskController) MyMethod(id int, field string) (map[string]interface{}, error) {
	if u, err := getObjectField(id, field); err == nil {
		return u, nil
	} else {
		return nil, context.NotFound
	}
}
```

will automatically route the http parameters id and field (i.e. `/tasks/5?field=name` ) to the correct method parameters, and will render the method return value as JSON. If the method returns an error it will be rendered as an http status code. If the parameter does not exist in the http request it will be passed to the method as the zero value for that parameter, unless that parameter is marked as ‘required’ using annotations. This will return an error without calling the method. For more information, see [Parameters]({{< ref "/beego/mvcIntroduction/controllers/requestParameters" >}})

​	将自动将 http 参数 id 和 field（即 `/tasks/5?field=name` ）路由到正确的方法参数，并将方法返回值呈现为 JSON。如果方法返回错误，它将被呈现为 http 状态代码。如果参数不存在于 http 请求中，它将作为该参数的零值传递给方法，除非该参数使用注释标记为“必需”。这将返回错误，而不调用该方法。有关更多信息，请参阅参数

## Method Expression Router 方法表达式路由器

The method expression router is to register routers by providing a controller method expresion. If the receiver of the controller method is a non-pointer type, then you can pass method expression as `pkg.controller.method`. If the receiver of method is a pointer, then you need to pass method expression as `(*pkg.controller).method`. However, if you register router in the same package as controller, then you don’t need to provide `pkg`.

​	方法表达式路由器是通过提供控制器方法表达式来注册路由器。如果控制器方法的接收者是非指针类型，则可以将方法表达式作为 `pkg.controller.method` 传递。如果方法的接收者是指针，则需要将方法表达式作为 `(*pkg.controller).method` 传递。但是，如果在与控制器相同的包中注册路由器，则无需提供 `pkg` 。

```golang
// Here are some examples:

type BaseController struct {
	web.Controller
}

func (b BaseController) Ping() {
	b.Data["json"] = "pong"
	b.ServeJSON()
}

func (b *BaseController) PingPointer() {
	b.Data["json"] = "pong_pointer"
	b.ServeJSON()
}

func main() {
	web.CtrlGet("/ping", BaseController.Ping)
	web.CtrlGet("/ping_pointer", (*BaseController).PingPointer)
	web.Run()
}
```

There are many other Method Expression Routers：

​	还有许多其他方法表达式路由器：

- web.CtrlGet(router, pkg.controller.method)
- web.CtrlPost(router, pkg.controller.method)
- web.CtrlPut(router, pkg.controller.method)
- web.CtrlPatch(router, pkg.controller.method)
- web.CtrlHead(router, pkg.controller.method)
- web.CtrlOptions(router, pkg.controller.method)
- web.CtrlDelete(router, pkg.controller.method)
- web.CtrlAny(router, pkg.controller.method)

It also provides namespace functions：

​	它还提供了命名空间函数：

- web.NSCtrlGet
- web.NSCtrlPost
- ……

## namespace 命名空间

```
//init namespace
ns :=
web.NewNamespace("/v1",
    web.NSCond(func(ctx *context.Context) bool {
        if ctx.Input.Domain() == "api.web.me" {
            return true
        }
        return false
    }),
    web.NSBefore(auth),
    web.NSGet("/notallowed", func(ctx *context.Context) {
        ctx.Output.Body([]byte("notAllowed"))
    }),
    web.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
    web.NSRouter("/changepassword", &UserController{}),
    web.NSNamespace("/shop",
        web.NSBefore(sentry),
        web.NSGet("/:id", func(ctx *context.Context) {
            ctx.Output.Body([]byte("notAllowed"))
        }),
    ),
    web.NSNamespace("/cms",
        web.NSInclude(
            &controllers.MainController{},
            &controllers.CMSController{},
            &controllers.BlockController{},
        ),
    ),
)

//register namespace
web.AddNamespace(ns)
```

the code set out above supports the URL:

​	上面列出的代码支持以下 URL：

- GET /v1/changepassword
- POST /v1/changepassword
- GET /v1/shop/123
- GET /v1/cms/ maps to annotation routers in MainController, CMSController, BlockController
  GET /v1/cms/ 映射到 MainController、CMSController、BlockController 中的注释路由器

namespace supports filter, condition and nested namespace

​	命名空间支持过滤器、条件和嵌套命名空间

namespace API:

​	命名空间 API：

- NewNamespace(prefix string,…interface{})

  Create a namespace object. The namespace object’s methods are listed below. For compatibility with the gofmt tool is is recommend that these method names begin with `NS`.

  ​	创建一个命名空间对象。命名空间对象的方法列在下面。为了与 gofmt 工具兼容，建议这些方法名以 `NS` 开头。

- NSCond(cond namespaceCond)

  if the namespaceCond returns true this namespace will be run.

  ​	如果 namespaceCond 返回 true，则将运行此命名空间。

- NSBefore(filterList …FilterFunc)

- NSAfter(filterList …FilterFunc)

  For `BeforeRouter` and `FinishRouter` filters. Multiple filters can be registered.

  ​	对于 `BeforeRouter` 和 `FinishRouter` 过滤器。可以注册多个过滤器。

- NSInclude(cList …ControllerInterface)

- NSRouter(rootpath string, c ControllerInterface, mappingMethods …string)

- NSGet(rootpath string, f FilterFunc)

- NSPost(rootpath string, f FilterFunc)

- NSDelete(rootpath string, f FilterFunc)

- NSPut(rootpath string, f FilterFunc)

- NSHead(rootpath string, f FilterFunc)

- NSOptions(rootpath string, f FilterFunc)

- NSPatch(rootpath string, f FilterFunc)

- NSAny(rootpath string, f FilterFunc)

- NSHandler(rootpath string, h http.Handler)

- NSAutoRouter(c ControllerInterface)

- NSAutoPrefix(prefix string, c ControllerInterface)

  These are methods to set up routers equivilant to the basic routers.

  ​	这些方法用于设置与基本方法等效的内容。

- NSNamespace(prefix string, params …innnerNamespace)

  ​	NSNamespace(前缀字符串、参数…内部命名空间)

  Nested namespaces

  ​	嵌套

  ```
  ns :=
  web.NewNamespace("/v1",
      web.NSNamespace("/shop",
          web.NSGet("/:id", func(ctx *context.Context) {
              ctx.Output.Body([]byte("shopinfo"))
          }),
      ),
      web.NSNamespace("/order",
          web.NSGet("/:id", func(ctx *context.Context) {
              ctx.Output.Body([]byte("orderinfo"))
          }),
      ),
      web.NSNamespace("/crm",
          web.NSGet("/:id", func(ctx *context.Context) {
              ctx.Output.Body([]byte("crminfo"))
          }),
      ),
  )
  ```

The methods below are for the`*Namespace` object and are not recommended. They have the same functionality as methods with `NS`, but are less elegant and harder to read.

​	以下方法适用于 `*Namespace` 对象，不建议使用。它们具有与 `NS` 方法相同的功能，但不太简洁且更难阅读。

- Cond(cond namespaceCond)

  ​	Cond(cond 命名空间条件)

  if the namespaceCond return true will run this namespace.

  ​	如果命名空间条件返回 true，将运行此命名空间。

- Filter(action string, filter FilterFunc)

  ​	Filter(操作字符串、过滤器 FilterFunc)

  action represents which position to run ,`before` and `after` is two validate value

  ​	操作表示要运行的位置， `before` 和 `after` 是两个验证值

- Router(rootpath string, c ControllerInterface, mappingMethods …string)

  ​	Router(根路径字符串、控制器 ControllerInterface、映射方法…字符串)

- AutoRouter(c ControllerInterface)

  ​	AutoRouter(控制器 ControllerInterface)

- AutoPrefix(prefix string, c ControllerInterface)

  ​	AutoPrefix(前缀字符串, c 控制器接口)

- Get(rootpath string, f FilterFunc)

  ​	Get(根路径字符串, f 过滤器函数)

- Post(rootpath string, f FilterFunc)

  ​	Post(根路径字符串, f 过滤器函数)

- Delete(rootpath string, f FilterFunc)

  ​	Delete(根路径字符串, f 过滤器函数)

- Put(rootpath string, f FilterFunc)

  ​	Put(根路径字符串, f 过滤器函数)

- Head(rootpath string, f FilterFunc)

  ​	Head(根路径字符串, f 过滤器函数)

- Options(rootpath string, f FilterFunc)

  ​	Options(根路径字符串, f 过滤器函数)

- Patch(rootpath string, f FilterFunc)

  ​	Patch(根路径字符串, f 过滤器函数)

- Any(rootpath string, f FilterFunc)

  ​	Any(根路径字符串, f 过滤器函数)

- Handler(rootpath string, h http.Handler)

  ​	Handler(根路径字符串, h http.Handler)

  these functions are the same as mentioned earlier

  ​	这些函数与前面提到的相同

- Namespace(ns …*Namespace)

More functions can be nested:

​	可以嵌套更多函数：

```go
//APIS
ns :=
	web.NewNamespace("/api",
		//It should verify the encrypted request in the production using
		web.NSCond(func(ctx *context.Context) bool {
			if ua := ctx.Input.Request.UserAgent(); ua != "" {
				return true
			}
			return false
		}),
		web.NSNamespace("/ios",
			//CRUD Create, Read, Update and Delete
			web.NSNamespace("/create",
				// /api/ios/create/node/
				web.NSRouter("/node", &apis.CreateNodeHandler{}),
				// /api/ios/create/topic/
				web.NSRouter("/topic", &apis.CreateTopicHandler{}),
			),
			web.NSNamespace("/read",
				web.NSRouter("/node", &apis.ReadNodeHandler{}),
				web.NSRouter("/topic", &apis.ReadTopicHandler{}),
			),
			web.NSNamespace("/update",
				web.NSRouter("/node", &apis.UpdateNodeHandler{}),
				web.NSRouter("/topic", &apis.UpdateTopicHandler{}),
			),
			web.NSNamespace("/delete",
				web.NSRouter("/node", &apis.DeleteNodeHandler{}),
				web.NSRouter("/topic", &apis.DeleteTopicHandler{}),
			)
		),
	)

web.AddNamespace(ns)
```
