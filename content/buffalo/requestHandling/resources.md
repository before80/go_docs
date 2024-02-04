+++
title = "资源"
date = 2024-02-04T21:07:51+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/request_handling/resources/]({{< ref "/buffalo/requestHandling/resources" >}})

# Resources 资源 

Often web applications need to build very similar [“CRUD”](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) end-points. To help reduce the amount of thought and complexity involved in this, Buffalo supports the concept of a “Resource”.

​	通常，Web 应用程序需要构建非常相似的“CRUD”端点。为了帮助减少与此相关的思考和复杂性，Buffalo 支持“资源”的概念。

The [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface allows Buffalo to map common routes and respond to common requests.

​	 `github.com/gobuffalo/buffalo#Resource` 接口允许 Buffalo 映射常见路由并响应常见请求。

Since **0.14.1**
自 0.14.1 起



```go
type Resource interface {
	List(Context) error
	Show(Context) error
	Create(Context) error
	Update(Context) error
	Destroy(Context) error
}
```

The [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface was made smaller in release `v0.14.1`. The `New` and `Edit` methods, which serve the HTML forms to edit the resource, are now optional.

​	 `github.com/gobuffalo/buffalo#Resource` 接口在版本 `v0.14.1` 中变得更小。用于提供用于编辑资源的 HTML 表单的 `New` 和 `Edit` 方法现在是可选的。

Here’s what the interface looked like before:

​	以下是该接口以前的样子：

```go
type Resource interface {
	List(Context) error
	Show(Context) error
	New(Context) error
	Create(Context) error
	Edit(Context) error
	Update(Context) error
	Destroy(Context) error
}
```

## Using Resources 使用资源 

After implementing the necessary methods on the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface, the resource can then be mapped to the application using the [`github.com/gobuffalo/buffalo#App.Resource`](https://godoc.org/github.com/gobuffalo/buffalo#App.Resource) method.

​	在 `github.com/gobuffalo/buffalo#Resource` 接口上实现必要的方法后，可以使用 `github.com/gobuffalo/buffalo#App.Resource` 方法将资源映射到应用程序。

```go
// action/users.go
type UsersResource struct{ }

func (u UsersResource) List(c buffalo.Context) error {
  // do work
}

func (u UsersResource) Show(c buffalo.Context) error {
  // do work
}

func (u UsersResource) Create(c buffalo.Context) error {
  // do work
}

func (u UsersResource) Update(c buffalo.Context) error {
  // do work
}

func (u UsersResource) Destroy(c buffalo.Context) error {
  // do work
}
```

Mapping the Resource in app.go:

​	在 app.go 中映射资源：

```go
// actions/app.go
app.Resource("/users", UsersResource{})
```

The above code example would be the equivalent of the following:

​	上面的代码示例相当于以下内容：

```go
// actions/app.go
ur := UsersResource{}

app.GET("/users", ur.List)
app.GET("/users/{user_id}", ur.Show)
app.POST("/users", ur.Create)
app.PUT("/users/{user_id}", ur.Update)
app.DELETE("/users/{user_id}", ur.Destroy)
```

It will produce a routing table that looks similar to:

​	它将生成一个类似于以下内容的路由表：

```bash
$ buffalo routes

METHOD | HOST                  | PATH                    | ALIASES | NAME                 | HANDLER
------ | ----                  | ----                    | ------- | ----                 | -------
GET    | http://127.0.0.1:3000 | /users/                 |         | usersPath            | coke/actions.UsersResource.List
POST   | http://127.0.0.1:3000 | /users/                 |         | usersPath            | coke/actions.UsersResource.Create
GET    | http://127.0.0.1:3000 | /users/{user_id}/       |         | userPath             | coke/actions.UsersResource.Show
PUT    | http://127.0.0.1:3000 | /users/{user_id}/       |         | userPath             | coke/actions.UsersResource.Update
DELETE | http://127.0.0.1:3000 | /users/{user_id}/       |         | userPath             | coke/actions.UsersResource.Destroy
```

## Optional Resource Methods 可选资源方法 # 自 0.14.1 起

Since **0.14.1**
在 中， 变小了，现在以下方法是可选的：



In `v0.14.1` the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) was made smaller with the following methods now being optional:

​	实现 `v0.14.1` 和 `github.com/gobuffalo/buffalo#Resource` 方法时，将向路由表添加以下内容：

```go
New(Context) error
Edit(Context) error
```

When implemented the `New` and `Edit` methods will add the following to the routing table:

​	生成资源 #

```bash
METHOD | HOST                   | PATH                   | ALIASES | NAME         | HANDLER
------ | ----                   | ----                   | ------- | ----         | -------
GET    | http://127.0.0.1:3000  | /users/new             |         | newUsersPath | coke/actions.UsersResource.New
GET    | http://127.0.0.1:3000  | /users/{user_id}/edit/ |         | editUserPath | coke/actions.UsersResource.Edit
```

## Generating Resources 命令将生成必要的模型、迁移、Go 代码和 HTML 来对资源进行 CRUD 操作。

The `buffalo generate resource` command will generate the necessary models, migrations, Go code, and HTML to CRUD the resource.

​	在 API 应用程序中运行生成器时，Buffalo 将生成代码以满足 `buffalo generate resource` 接口。

When running the generator in an API application Buffalo will generate code to meet the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface.

​	在 Web 应用程序中运行生成器时，Buffalo 将生成代码以满足 `github.com/gobuffalo/buffalo#Resource` 接口，以及可选的 和 方法。

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  Create(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

When running the generator in a Web application Buffalo will generate code to the meet the [`github.com/gobuffalo/buffalo#Resource`](https://godoc.org/github.com/gobuffalo/buffalo#Resource) interface, as well as the optional `New` and `Edit` methods.

```go
type Resource interface {
  List(Context) error
  Show(Context) error
  New(Context) error
  Create(Context) error
  Edit(Context) error
  Update(Context) error
  Destroy(Context) error
}
```

## Example Resource Generation 示例资源生成 

In this example Buffalo will generate the code needed to CRUD a resource named `widget` (Go: `Widget`) that has the following attributes:

​	在此示例中，Buffalo 将生成创建、读取、更新和删除名为 `widget` (Go： `Widget` ) 的资源所需的代码，该资源具有以下属性：

|               | Model Attribute 模型属性 | Go Type Go 类型                                              | DB type 数据库类型   | Form Type 表单类型 |
| ------------- | ------------------------ | ------------------------------------------------------------ | -------------------- | ------------------ |
| `title`       | `Title`                  | `string`                                                     | `varchar`            | `text`             |
| `description` | `Description`            | [`nulls.String`](https://godoc.org/github.com/gobuffalo/pop/nulls#String) | `varchar (nullable)` | `textarea`         |

```bash
$ buffalo generate resource widget title description:nulls.Text
```

actions/app.go

actions/widgets.go

actions/widgets_test.go

```go
package actions

import (
	"net/http"

	"coke/locales"
	"coke/models"
	"coke/public"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	i18n "github.com/gobuffalo/mw-i18n/v2"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

var (
	app *buffalo.App
	T   *i18n.Translator
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_coke_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//   c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))
		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)

		app.Resource("/widgets", WidgetsResource{})
		app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
```

locales/widgets.en-us.yaml

```yaml
- id: "widget.created.success"
  translation: "Widget was successfully created."
- id: "widget.updated.success"
  translation: "Widget was successfully updated."
- id: "widget.destroyed.success"
  translation: "Widget was successfully destroyed."
```

migrations/20181005153028_create_widgets.up.fizz

migrations/20181005153028_create_widgets.down.fizz

```erb
create_table("widgets") {
	t.Column("id", "uuid", {primary: true})
	t.Column("title", "string", {})
	t.Column("description", "text", {null: true})
	t.Timestamps()
}
```

models/widget.go

models/widget_test.go

```go
package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Widget is used by pop to map your widgets database table to your go code.
type Widget struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	Title       string       `json:"title" db:"title"`
	Description nulls.String `json:"description" db:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (w Widget) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Widgets is not required by pop and may be deleted
type Widgets []Widget

// String is not required by pop and may be deleted
func (w Widgets) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *Widget) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: w.Title, Name: "Title"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *Widget) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *Widget) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
```

templates/widgets/_form.plush.html

templates/widgets/edit.plush.html

templates/widgets/index.plush.html

templates/widgets/new.plush.html

templates/widgets/show.plush.html

```html
<%= f.InputTag("Title") %>
<%= f.TextAreaTag("Description", {rows: 10}) %>
<button class="btn btn-success" role="submit">Save</button>
```

## Destroying Resources 销毁资源 

You can remove files generated by this generator by running:

​	您可以通过运行以下命令来删除此生成器生成的文件：

```bash
$ buffalo destroy resource users
```

This command will ask you which files you want to remove, you can either answer each of the questions with `y/n` or you can pass the `-y` flag to the command like:

​	此命令会询问您要删除哪些文件，您可以用 `y/n` 回答每个问题，或者可以像下面这样将 `-y` 标志传递给命令：

```bash
$ buffalo destroy resource users -y
```

Or in short form:

​	或简短形式：

```bash
$ buffalo d r users -y
```

## Nesting Resources 嵌套资源 

To simplify creating resource hierarchies, Buffalo supports nesting resources.

​	为了简化创建资源层次结构，Buffalo 支持嵌套资源。

```go
type WidgetsResource struct {
	buffalo.Resource
}

type ImagesResource struct {
  buffalo.Resource
}

w := app.Resource("/widgets", WidgetsResource{})
w.Resource("/images", ImagesResource{})
```

This results in the following routes:

​	这将导致以下路由：

```bash
$ buffalo routes

METHOD | HOST                  | PATH                                         | ALIASES | NAME                | HANDLER
------ | ----                  | ----                                         | ------- | ----                | -------
GET    | http://127.0.0.1:3000 | /                                            |         | rootPath            | coke/actions.HomeHandler
GET    | http://127.0.0.1:3000 | /widgets/                                    |         | widgetsPath         | coke/actions.WidgetsResource.List
POST   | http://127.0.0.1:3000 | /widgets/                                    |         | widgetsPath         | coke/actions.WidgetsResource.Create
GET    | http://127.0.0.1:3000 | /widgets/new/                                |         | newWidgetsPath      | coke/actions.WidgetsResource.New
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/                        |         | widgetPath          | coke/actions.WidgetsResource.Show
PUT    | http://127.0.0.1:3000 | /widgets/{widget_id}/                        |         | widgetPath          | coke/actions.WidgetsResource.Update
DELETE | http://127.0.0.1:3000 | /widgets/{widget_id}/                        |         | widgetPath          | coke/actions.WidgetsResource.Destroy
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/edit/                   |         | editWidgetPath      | coke/actions.WidgetsResource.Edit
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/                 |         | widgetImagesPath    | coke/actions.ImagesResource.List
POST   | http://127.0.0.1:3000 | /widgets/{widget_id}/images/                 |         | widgetImagesPath    | coke/actions.ImagesResource.Create
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/new/             |         | newWidgetImagesPath | coke/actions.ImagesResource.New
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/{image_id}/      |         | widgetImagePath     | coke/actions.ImagesResource.Show
PUT    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/{image_id}/      |         | widgetImagePath     | coke/actions.ImagesResource.Update
DELETE | http://127.0.0.1:3000 | /widgets/{widget_id}/images/{image_id}/      |         | widgetImagePath     | coke/actions.ImagesResource.Destroy
GET    | http://127.0.0.1:3000 | /widgets/{widget_id}/images/{image_id}/edit/ |         | editWidgetImagePath | coke/actions.ImagesResource.Edit
```

## buffalo.BaseResource

When a resource is generated it has [`buffalo.BaseResource`](https://godoc.org/github.com/gobuffalo/buffalo#BaseResource) embedded into it.

​	生成资源时，其中嵌入了 `buffalo.BaseResource` 。

```go
type Widget struct {
  buffalo.BaseResource
}
```

The `buffalo.BaseResource` has basic implementations for all of the methods required by `buffalo.Resource`. These methods all `404`.

​	 `buffalo.BaseResource` 具有 `buffalo.Resource` 所需的所有方法的基本实现。这些方法全部 `404` 。

```go
// Edit default implementation. Returns a 404
func (v BaseResource) Edit(c Context) error {
  return c.Error(http.StatusNotFound, errors.New("resource not implemented"))
}
```

## Video Presentation 视频演示 

{{< vimeo "212302823">}}

## Related Content 相关内容 

- [Actions]({{< ref "/buffalo/requestHandling/actionController" >}}) - Learn more about Buffalo actions.
  操作 - 了解有关 Buffalo 操作的更多信息。

## Next Steps 后续步骤 

- [Context]({{< ref "/buffalo/requestHandling/context" >}}) - Learn more about Buffalo context.
  上下文 - 了解有关 Buffalo 上下文的更多信息。
