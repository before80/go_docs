+++
title = "控制器函数"
date = 2024-02-04T09:57:05+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/controller/]({{< ref "/beego/mvcIntroduction/controllers/controllerFuncs" >}})

# Controller funcs 控制器函数



## Introduction to controller 控制器简介

> Note: From version 1.6: `this.ServeJson()` has been changed to `this.ServeJSON()` and `this.TplNames` has been changed to `this.TplName`
>
> ​	注意：从 1.6 版本开始： `this.ServeJson()` 已更改为 `this.ServeJSON()` ， `this.TplNames` 已更改为 `this.TplName`

Beego’s controller needs to be embeded as `beego.Controller`:

​	Beego 的控制器需要嵌入为 `beego.Controller` ：

```
type xxxController struct {
    web.Controller
}
```

`web.Controller` implements interface `web.ControllerInterface`. `web.ControllerInterface` defines these functions:

​	 `web.Controller` 实现接口 `web.ControllerInterface` 。 `web.ControllerInterface` 定义了以下函数：

- Init(ct *context.Context, controllerName, actionName string, app interface{})

  This function will initialize Context, Controller name, template name, template variable container `Data`. `app` is the executing Controller’s reflecttype. This can be used to execute the subclass’s methods.

  ​	此函数将初始化 Context、Controller 名称、模板名称、模板变量容器 `Data` 。 `app` 是正在执行的 Controller 的反射类型。这可用于执行子类的函数。

- Prepare()

  This function is used for extension and will execute before the methods below. It can be overwritten to implement functions such as user validation.

  ​	此函数用于扩展，将在以下方法之前执行。可以覆盖它来实现诸如用户验证之类的函数。

- Get()

  This method will be executed if the HTTP request method is GET. It returns 403 by default. This can be used to handle GET requests by overwriting them in the struct of subclass.

  ​	如果 HTTP 请求方法是 GET，则将执行此方法。默认情况下，它返回 403。可以通过在子类的结构中覆盖它们来处理 GET 请求。

- Post()

  This method will be executed if the HTTP request method is POST. It returns 403 by default. This can be used to handle POST requests by overwriting them in the struct of subclass.

  ​	如果 HTTP 请求方法是 POST，则将执行此方法。默认情况下，它返回 403。可以通过在子类的结构中覆盖它们来处理 POST 请求。

- Delete()

  This method will be executed if the HTTP request method is DELETE. It returns 403 by default. This can be used to handle DELETE requests by overwriting them in the struct of subclass.

  ​	如果 HTTP 请求方法是 DELETE，则将执行此方法。默认情况下，它返回 403。可以通过在子类的结构中覆盖它们来处理 DELETE 请求。

- Put()

  This method will be executed if the HTTP request method is PUT. It returns 403 by default. This can be used to handle PUT requests by overwriting them in the struct of subclass.

  ​	如果 HTTP 请求方法是 PUT，则将执行此方法。默认情况下，它返回 403。这可用于通过在子类的结构中覆盖 PUT 请求来处理 PUT 请求。

- Head()

  This method will be executed if the HTTP request method is HEAD. It return 403 by default. This can be used to handle HEAD requests by overwriting them in the struct of subclass.

  ​	如果 HTTP 请求方法是 HEAD，则将执行此方法。默认情况下，它返回 403。这可用于通过在子类的结构中覆盖 HEAD 请求来处理 HEAD 请求。

- Patch()

  This method will be executed if the HTTP request method is PATCH. It returns 403 by default. This can be used to handle PATCH requests by overwriting them in the struct of subclass.

  ​	如果 HTTP 请求方法是 PATCH，则将执行此方法。默认情况下，它返回 403。这可用于通过在子类的结构中覆盖 PATCH 请求来处理 PATCH 请求。

- Options()

  This method will be executed if the HTTP request method is OPTIONS. It returns 403 by default. This can be used to handle OPTIONS requests by overwriting them in the struct of subclass.

  ​	如果 HTTP 请求方法是 OPTIONS，则将执行此方法。默认情况下，它返回 403。这可用于通过在子类的结构中覆盖 OPTIONS 请求来处理 OPTIONS 请求。

- Trace() error

  This method will be executed if the HTTP request method is TRACE. It returns 403 by default. This can be used to handle TRACE requests by overwriting them in the struct of subclass.

  ​	如果 HTTP 请求方法是 TRACE，则将执行此方法。默认情况下，它返回 403。这可用于通过在子类的结构中覆盖 TRACE 请求来处理 TRACE 请求。

- Finish()

  This method will be executed after finishing the related HTTP method. It is empty by default. This can be implemented by overwriting it in the struct of subclass. It is used for database closing, data cleaning and so on.

  ​	此方法将在完成相关的 HTTP 方法后执行。默认情况下为空。可以通过在子类的结构中覆盖它来实现。它用于数据库关闭、数据清理等。

- Render() error

  This method is used to render templates. It is only executed if `web.AutoRender` is set to true.

  ​	此方法用于渲染模板。仅当 `web.AutoRender` 设置为 true 时才执行。

- Mapping(method string, fn func())

  Register a method. Generally, the `method` is valid HTTP method name.

  ​	注册一个方法。通常， `method` 是有效的 HTTP 方法名称。

- HandlerFunc(fnname string) bool

  Execute the method that register by `Mapping` method. It return false when `fnname` not found.

  ​	执行通过 `Mapping` 方法注册的方法。当找不到 `fnname` 时，它返回 false。

- RenderBytes() ([]byte, error)

  Render the template and output the result as `[]byte`. This method doesn’t check `EnableRender` option, and it doesn’t output the result to response.

  ​	渲染模板并将结果输出为 `[]byte` 。此方法不检查 `EnableRender` 选项，并且不会将结果输出到响应。

- RenderString() (string, error)

  Similar with `RenderBytes()`

  ​	与 `RenderBytes()` 类似

- Redirect(url string, code int)

  Redirect the request to `url`

  ​	将请求重定向到 `url`

- SetData(data interface{})

  Store `data` to `controller`. Usually you won’t use this method.

  ​	将 `data` 存储到 `controller` 。通常您不会使用此方法。

- Abort(code string)

  Breaking current method with the code. [errors]({{< ref "/beego/mvcIntroduction/controllers/errorHanding" >}})

  ​	使用代码中断当前方法。错误

- CustomAbort(status int, body string)

  Breaking current method with the code. [errors]({{< ref "/beego/mvcIntroduction/controllers/errorHanding" >}})

  ​	使用代码中断当前方法。错误

- StopRun()

  Trigger panic.

  ​	触发恐慌。

- ServeXXX(encoding …bool) error

  Return response with the specific format. Supporting JSON, JSONP, XML, YAML. [Output]({{< ref "/beego/mvcIntroduction/controllers/responseFormats" >}})

  ​	返回具有特定格式的响应。支持 JSON、JSONP、XML、YAML。输出

- ServeFormatted(encoding …bool) error

  Return response with specific format passed from client’s `Accept` option. [Output]({{< ref "/beego/mvcIntroduction/controllers/responseFormats" >}})

  ​	返回具有从客户端的 `Accept` 选项传递的特定格式的响应。输出

- Input() (url.Values, error)

  Return all parameters.

  ​	返回所有参数。

- ParseForm(obj interface{}) error

  Deserialize form to obj.

  ​	将表单反序列化为 obj。

- GetXXX(key string, def…) XXX, err

  Read value from parameters. If the `def` not empty, return `def` when key not found or error. XXX coule be basic types, string or File.

  ​	从参数中读取值。如果 `def` 不为空，则在找不到键或出错时返回 `def` 。XXX 可以是基本类型、字符串或文件。

- SaveToFile(fromfile, tofile string) error

  Save the uploading file to file system. `fromfile` is uploading file name.

  ​	将上传的文件保存到文件系统。 `fromfile` 是上传的文件名。

- SetSession(name interface{}, value interface{}) error

  Put some value into session.

  ​	将某个值放入会话。

- GetSession(name interface{}) interface{}

  Read value from session.

  ​	从会话中读取值。

- DelSession(name interface{}) error

  Delete the value from session.

  ​	从会话中删除值。

- SessionRegenerateID() error

  Re-generate session id.

  ​	重新生成会话 ID。

- DestroySession() error

  ​	DestroySession() 错误

  Destroy session.

  ​	销毁会话。

- IsAjax() bool

  Check whether is ajax request.

  ​	检查是否是 ajax 请求。

- GetSecureCookie(Secret, key string) (string, bool)

  Read value from cookie.

  ​	从 cookie 中读取值。

- SetSecureCookie(Secret, name, value string, others …interface{})

  Put key-value pair into cookie.

  ​	将键值对放入 cookie 中。

- XSRFToken() string

  Create `CSRF` token.

  ​	创建 `CSRF` 令牌。

- CheckXSRFCookie() bool

  Check `CSRF` token

  ​	检查 `CSRF` 令牌

## Custom logic 自定义逻辑

Custom logic can be implemented by overwriting functions in struct. For example:

​	可以通过覆盖结构中的函数来实现自定义逻辑。例如：

```
type AddController struct {
    web.Controller
}

func (this *AddController) Prepare() {

}

func (this *AddController) Get() {
    this.Data["content"] = "value"
    this.Layout = "admin/layout.html"
    this.TplName = "admin/add.tpl"
}

func (this *AddController) Post() {
    pkgname := this.GetString("pkgname")
    content := this.GetString("content")
    pk := models.GetCruPkg(pkgname)
    if pk.Id == 0 {
        var pp models.PkgEntity
        pp.Pid = 0
        pp.Pathname = pkgname
        pp.Intro = pkgname
        models.InsertPkg(pp)
        pk = models.GetCruPkg(pkgname)
    }
    var at models.Article
    at.Pkgid = pk.Id
    at.Content = content
    models.InsertArticle(at)
    this.Ctx.Redirect(302, "/admin/index")
}
```

In the example above a RESTful structure has been implemented by overwriting functions.

​	在上面的示例中，通过覆盖函数来实现 RESTful 结构。

The following example implements a baseController and other initialization methods that will be inherited by other controllers:

​	以下示例实现了 baseController 和其他初始化方法，这些方法将由其他控制器继承：

```
type NestPreparer interface {
        NestPrepare()
}

// baseRouter implements global settings for all other routers.
type baseRouter struct {
        web.Controller
        i18n.Locale
        user    models.User
        isLogin bool
}
// Prepare implements Prepare method for baseRouter.
func (this *baseRouter) Prepare() {

        // page start time
        this.Data["PageStartTime"] = time.Now()

        // Setting properties.
        this.Data["AppDescription"] = utils.AppDescription
        this.Data["AppKeywords"] = utils.AppKeywords
        this.Data["AppName"] = utils.AppName
        this.Data["AppVer"] = utils.AppVer
        this.Data["AppUrl"] = utils.AppUrl
        this.Data["AppLogo"] = utils.AppLogo
        this.Data["AvatarURL"] = utils.AvatarURL
        this.Data["IsProMode"] = utils.IsProMode

        if app, ok := this.AppController.(NestPreparer); ok {
                app.NestPrepare()
        }
}
```

The above example defines a base class and initializes some variables. It will test if the executing Controller is an implementation of NestPreparer. If it is it calls the method of subclass.

​	上面的示例定义了一个基类并初始化了一些变量。它将测试正在执行的控制器是否是 NestPreparer 的实现。如果是，它将调用子类的函数。

The example below shows an implementation of `NestPreparer`:

​	下面的示例显示了 `NestPreparer` 的实现：

```
type BaseAdminRouter struct {
    baseRouter
}

func (this *BaseAdminRouter) NestPrepare() {
    if this.CheckActiveRedirect() {
            return
    }

    // if user isn't admin, then logout user
    if !this.user.IsAdmin {
            models.LogoutUser(&this.Controller)

            // write flash message
            this.FlashWrite("NotPermit", "true")

            this.Redirect("/login", 302)
            return
    }

    // current in admin page
    this.Data["IsAdmin"] = true

    if app, ok := this.AppController.(ModelPreparer); ok {
            app.ModelPrepare()
            return
    }
}

func (this *BaseAdminRouter) Get(){
	this.TplName = "Get.tpl"
}

func (this *BaseAdminRouter) Post(){
	this.TplName = "Post.tpl"
}
```

The above example first executes `Prepare`. Go will search for methods in the struct by looking in the parent classes. `BaseAdminRouter` will execute and checks whether there is a `Prepare` method. If not it keeps searching `baseRouter`. If there is it will execute the logic. `this.AppController` in `baseRouter` is the currently executing Controller `BaseAdminRouter`. Next, it will execute `BaseAdminRouter.NestPrepare` method. Finally, it will start executing the related `GET` or `POST` method.

​	上面的示例首先执行 `Prepare` 。Go 将通过在父类中查找来搜索结构中的方法。 `BaseAdminRouter` 将执行并检查是否存在 `Prepare` 方法。如果没有，它将继续搜索 `baseRouter` 。如果有，它将执行逻辑。 `baseRouter` 中的 `this.AppController` 是当前正在执行的控制器 `BaseAdminRouter` 。接下来，它将执行 `BaseAdminRouter.NestPrepare` 方法。最后，它将开始执行相关的 `GET` 或 `POST` 方法。

## Stop controller executing immediately 立即停止控制器执行

To stop the execution logic of a request and return the response immediately use `StopRun()`. For example, when a user authentication fails in `Prepare` method a response will be returned immediately.

​	要停止请求的执行逻辑并立即返回响应，请使用 `StopRun()` 。例如，当用户身份验证在 `Prepare` 方法中失败时，将立即返回响应。

```
type RController struct {
    web.Controller
}

func (this *RController) Prepare() {
    this.Data["json"] = map[string]interface{}{"name": "astaxie"}
    this.ServeJSON()
    this.StopRun()
}
```

> If you call `StopRun` the `Finish` method won’t be run. To free resources call `Finish` manually before calling `StopRun`.
>
> ​	如果您调用 `StopRun` ，则不会运行 `Finish` 方法。要在调用 `StopRun` 之前释放资源，请手动调用 `Finish` 。
