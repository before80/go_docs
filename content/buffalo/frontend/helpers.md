+++
title = "帮助程序"
date = 2024-02-04T21:10:38+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/helpers/]({{< ref "/buffalo/frontend/helpers" >}})

# Helpers 帮助程序 

This document only applies when using https://github.com/gobuffalo/buffalo/tree/main/render. Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.
此文档仅适用于使用 https://github.com/gobuffalo/buffalo/tree/main/render 时。有关底层模板包的更多详细信息，请参阅 github.com/gobuffalo/plush。

## Builtin Helpers 内置帮助器 

A full list of all helper functions for [`github.com/gobuffalo/plush`](https://godoc.org/github.com/gobuffalo/plush) can be found at [`github.com/gobuffalo/helpers`](https://godoc.org/github.com/gobuffalo/helpers).

​	可以在 `github.com/gobuffalo/helpers` 找到 `github.com/gobuffalo/plush` 的所有帮助器函数的完整列表。

## Path Helpers 路径帮助器 

Buffalo will generate path helpers for all of the routes you add to the App. The easiest way to see what all of the generated path helpers are and what they point to is to run `buffalo routes`. This will print out a list that looks something like this:

​	Buffalo 将为添加到应用程序的所有路由生成路径帮助器。查看所有生成的路径帮助器及其指向的内容的最简单方法是运行 `buffalo routes` 。这将打印出类似以下内容的列表：

```text
$ buffalo routes
METHOD | HOST                   | PATH                         | ALIASES | NAME              | HANDLER
------ | ----                   | ----                         | ------- | ----              | -------
GET    | http://127.0.0.1:3000  | /                            |         | rootPath          | github.com/gobuffalo/coke/actions.HomeHandler
GET    | http://127.0.0.1:3000  | /about                       |         | aboutPath         | github.com/gobuffalo/coke/actions.AboutHandler
GET    | http://127.0.0.1:3000  | /drinks                      |         | drinksPath        | github.com/gobuffalo/coke/actions.DrinksResource.List
POST   | http://127.0.0.1:3000  | /drinks                      |         | drinksPath        | github.com/gobuffalo/coke/actions.DrinksResource.Create
GET    | http://127.0.0.1:3000  | /drinks/new                  |         | newDrinksPath     | github.com/gobuffalo/coke/actions.DrinksResource.New
GET    | http://127.0.0.1:3000  | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Show
PUT    | http://127.0.0.1:3000  | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Update
DELETE | http://127.0.0.1:3000  | /drinks/{drink_id}           |         | drinkPath         | github.com/gobuffalo/coke/actions.DrinksResource.Destroy
GET    | http://127.0.0.1:3000  | /drinks/{drink_id}/edit      |         | editDrinkPath     | github.com/gobuffalo/coke/actions.DrinksResource.Edit
GET    | http://127.0.0.1:3000  | /api/v1/users                |         | apiV1UsersPath    | github.com/gobuffalo/coke/actions.UsersResource.List
POST   | http://127.0.0.1:3000  | /api/v1/users                |         | apiV1UsersPath    | github.com/gobuffalo/coke/actions.UsersResource.Create
GET    | http://127.0.0.1:3000  | /api/v1/users/new            |         | newApiV1UsersPath | github.com/gobuffalo/coke/actions.UsersResource.New
GET    | http://127.0.0.1:3000  | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Show
PUT    | http://127.0.0.1:3000  | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Update
DELETE | http://127.0.0.1:3000  | /api/v1/users/{user_id}      |         | apiV1UserPath     | github.com/gobuffalo/coke/actions.UsersResource.Destroy
GET    | http://127.0.0.1:3000  | /api/v1/users/{user_id}/edit |         | editApiV1UserPath | github.com/gobuffalo/coke/actions.UsersResource.Edit
```

Going down this list we start with the path *NAME*d `rootPath` which represents *PATH* `/` or the root route of the server and as a bonus, with all of these we can even see exactly which *HANDLER* code is being run for this METHOD+PATH combination.

​	向下滚动此列表，我们从名为 `rootPath` 的路径开始，它表示路径 `/` 或服务器的根路由，另外，使用所有这些，我们甚至可以看到确切地为这种 METHOD+PATH 组合运行哪个 HANDLER 代码。

Next we have a standard `app.GET("/about", AboutHandler)` which generates to `aboutPath`.

​	接下来，我们有一个标准 `app.GET("/about", AboutHandler)` ，它生成到 `aboutPath` 。

Then we use a resource `app.Resource("/drinks", DrinksResource{})`, which generates a path for each of our standard actions, and for each of those a helper to be used in templates. Those that take a parameter can be used like this `<%= drinkPath({drink_id: drink.ID}) %>`. All helpers take a `map[string]interface{}` that is used to fill-in parameters.

​	然后，我们使用资源 `app.Resource("/drinks", DrinksResource{})` ，它为我们的每个标准操作生成路径，并为每个操作生成一个用于模板的帮助器。那些带参数的可以像这样使用 `<%= drinkPath({drink_id: drink.ID}) %>` 。所有帮助器都采用 `map[string]interface{}` 来填充参数。

Finally, when we use a group we can see that this changes the generated helpers. Here is the routing for those last paths:

​	最后，当我们使用一个组时，我们可以看到这会改变生成的帮助器。以下是那些最后路径的路由：

```go
api := app.Group("/api/v1")
api.Resource("/users", UsersResource{})
```

**Note** that the helpers are generated to match the generated paths. It is possible to override the path names in the `App.Routes`, but it is highly advised that you find a different way to your goal than this. Slack is always open to these conversations.

​	请注意，帮助器是根据生成的路径生成的。可以在 `App.Routes` 中覆盖路径名称，但强烈建议您找到实现目标的其他方法。Slack 始终愿意进行这些对话。

### PathFor Helper PathFor 帮助器 

The [`github.com/gobuffalo/helpers/paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor) helper takes an `interface{}`, or a `slice` of them, and tries to convert it to a `/foos/{id}` style URL path.

​	 `github.com/gobuffalo/helpers/paths#PathFor` 助手接受一个 `interface{}` 或多个 `interface{}`，并尝试将其转换为 `/foos/{id}` 样式的 URL 路径。

Rules:

​	规则：

- if `string` it is returned as is
  如果 `string` ，则按原样返回
- if [`github.com/gobuffalo/helpers/paths#Pathable`](https://godoc.org/github.com/gobuffalo/helpers/paths#Pathable) the `ToPath` method is returned
  如果 `github.com/gobuffalo/helpers/paths#Pathable` ，则返回 `ToPath` 方法
- if `slice` or an `array` each element is run through the helper then joined
  如果 `slice` 或 `array` ，则每个元素都通过助手运行，然后连接
- if [`github.com/gobuffalo/helpers/paths#Paramable`](https://godoc.org/github.com/gobuffalo/helpers/paths#Paramable) the `ToParam` method is used to fill the `{id}` slot
  如果 `github.com/gobuffalo/helpers/paths#Paramable` 方法用于填充 `{id}` 插槽
- if `<T>.Slug` the slug is used to fill the `{id}` slot of the URL
  如果 `<T>.Slug` slug 用于填充 URL 的 `{id}` 插槽
- if `<T>.ID` the ID is used to fill the `{id}` slot of the URL
  如果 `<T>.ID` ID 用于填充 URL 的 `{id}` 插槽

### LinkTo Helpers LinkTo 帮助器 

### LinkTo

The [`github.com/gobuffalo/helpers/tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo) helpers creates HTML for a `<a>` tag using [`github.com/gobuffalo/tags`](https://godoc.org/github.com/gobuffalo/tags) to create tag with the given [`github.com/gobuffalo/tags#Options`](https://godoc.org/github.com/gobuffalo/tags#Options) and using [`github.com/gobuffalo/helpers/paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor) to set the `href` element.

​	 `github.com/gobuffalo/helpers/tags#LinkTo` 帮助器使用 `github.com/gobuffalo/tags` 创建具有给定 `github.com/gobuffalo/tags#Options` 的标签，并使用 `github.com/gobuffalo/helpers/paths#PathFor` 设置 `href` 元素，从而为 `<a>` 标签创建 HTML。

If given a block it will be interrupted and appended inside of the `<a>` tag.

​	如果给定一个块，它将被中断并追加到 `<a>` 标签内部。

#### Example 1: 示例 1：

```erb
<%= linkTo([user, widget], {class: "btn"}) %>

<a class="btn" href="/users/id/widget/slug"></a>
```

#### Example 2: 示例 2：

```erb
<%= linkTo("foo", {class: "btn"}) %>

<a class="btn" href="/foo"></a>
```

#### Example 3: 示例 3：

```erb
<%= linkTo(user, {class: "btn"}) { %>
Click Me!
<% } %>

<a class="btn" href="/users/id">Click Me!</a>
```

### RemoteLinkTo

The [`github.com/gobuffalo/helpers/tags#RemoteLinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#RemoteLinkTo) helper provides the same functionality as [`github.com/gobuffalo/helpers/tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo) but adds the `data-remote` element for use with https://www.npmjs.com/package/rails-ujs which is included in the default generated Webpack configuration.

​	 `github.com/gobuffalo/helpers/tags#RemoteLinkTo` 帮助程序提供与 `github.com/gobuffalo/helpers/tags#LinkTo` 相同的功能，但添加了 `data-remote` 元素，可与 https://www.npmjs.com/package/rails-ujs 配合使用，后者包含在默认生成的 Webpack 配置中。

#### Example 1: 示例 1：

```erb
<%= remoteLinkTo([user, widget], {class: "btn"}) %>

<a class="btn" data-remote="true" href="/users/id/widget/slug"></a>
```

#### Example 2: 示例 2：

```erb
<%= remoteLinkTo("foo", {class: "btn"}) %>

<a class="btn" data-remote="true" href="/foo"></a>
```

#### Example 3: 示例 3：

```erb
<%= remoteLinkTo(user, {class: "btn"}) { %>
Click Me!
<% } %>

<a class="btn" data-remote="true" href="/users/id">Click Me!</a>
```

## Content Helpers 内容帮助程序 

Plush ships with two complementary helpers that let you create dynamic HTML snippets and re-use them later in the template.

​	Plush 附带两个互补的帮助程序，可让您创建动态 HTML 片段并在模板中稍后重用它们。

### The `contentFor` and `contentOf` Helpers `contentFor` 和 `contentOf` 帮助程序 

The `contentFor` helper takes a block of HTML and holds on to it using the given name. This block can then be used elsewhere in a template file, even when the content defined in a `contentFor` block is in a yielded-to template and is expanded into a `contentOf` block in a `yield`-calling template. The default `templates/application.html` calls `yield` like this.

​	 `contentFor` 帮助程序获取一个 HTML 块并使用给定名称保存它。然后可以在模板文件的其他位置使用此块，即使在 `contentFor` 块中定义的内容位于已生成的模板中，并且在调用 `yield` 模板时扩展为 `contentOf` 块。默认 `templates/application.html` 会像这样调用 `yield` 。

Take the following example: suppose we have a `templates/application.html` that fully specifies everything in `<head>` and the outermost contents of `<body>`. This template yields to other subtemplates, like `templates/users/show.html`, to fill `<body>`. However, if we want to add or override something in the `<head>` from a subtemplate, we’ll need to use `contentFor`. In this example, we’ll add a way for subtemplates to add an extra chunk of CSS to the `<head>` of `application.html`:

​	以以下示例为例：假设我们有一个 `templates/application.html` ，它完全指定了 `<head>` 中的一切内容和 `<body>` 的最外层内容。此模板让位于其他子模板，如 `templates/users/show.html` ，以填充 `<body>` 。但是，如果我们想从子模板中添加或覆盖 `<head>` 中的内容，我们需要使用 `contentFor` 。在此示例中，我们将添加一种方法，以便子模板可以向 `application.html` 的 `<head>` 添加额外的 CSS 块：

```erb
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>My Site</title>
    <%= stylesheetTag("application.css") %>
    <%= contentOf("extraStyle") %>
  </head>
  <body>
    <div class="container">
      <%= partial("flash.html") %>
      <%= yield %>
    </div>
  </body>
</html>
```

As it turns out, our `users/index.html` template could use a little page-wide styling instead of adding a bunch of `style` attributes to different elements, so it defines a block of CSS that doesn’t show up anywhere inside the template:

​	事实证明，我们的 `users/index.html` 模板可以使用一些页面范围的样式，而不是向不同的元素添加一堆 `style` 属性，因此它定义了一个不会显示在模板内部任何位置的 CSS 块：

```erb
<div class="page-header">
  <h1>Users</h1>
</div>
<table class="table table-striped">
  <thead>
    <th>Username</th> <th>Password</th> <th>Email</th> <th>Admin?</th> <th>&nbsp;</th>
  </thead>
  <tbody>
    <%= for (user) in users { %>
      <!-- … -->
    <% } %>
  </tbody>
</table>

<% contentFor("extraStyle") { %>
  <style>
    .online {
      color: limegreen;
      background: black;
    }

    .offline {
      color: lightgray;
      background: darkgray;
    }
  </style>
<% } %>
```

The styling for the `online` and `offline` classes then appears at the end of `<head>` in `/users`. In other pages, nothing is added.

​	然后， `online` 和 `offline` 类的样式出现在 `/users` 中 `<head>` 的末尾。在其他页面中，不会添加任何内容。

Of course, if you’d rather do extensive processing on what goes into a chunk that goes on a webpage, you may want to do your processing in Go code instead of in templates. In that case, call, say, `c.Set("moonPhase", mp)` where `c` is a `buffalo.Context` in a function in an action like in `actions/users.go`, and `mp` is some string or object. Then, in your templates, refer to `<%= moonPhase %>` to display your expertly-calculated phase of the moon.

​	当然，如果您更愿意对网页上的块进行大量处理，您可能希望在 Go 代码中而不是在模板中进行处理。在这种情况下，调用，比如， `c.Set("moonPhase", mp)` ，其中 `c` 是函数中操作中的 `buffalo.Context` ，如 `actions/users.go` 中，而 `mp` 是某个字符串或对象。然后，在您的模板中，引用 `<%= moonPhase %>` 以显示您专业计算的月相。
