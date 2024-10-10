+++
title = "表单"
date = 2024-02-04T21:11:17+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/forms/]({{< ref "/buffalo/frontend/forms" >}})

# Forms 表单 

Buffalo uses the [`github.com/gobuffalo/tags`](https://github.com/gobuffalo/tags) package to make form building easier.

​	使用 `github.com/gobuffalo/tags` 包来使表单构建更容易。

Plush includes two helpers from this package that produce Bootstrap v3 style forms. These helpers are `form` and `form_for`.

​	Plush 包含了来自此包的两个帮助程序，它们生成 v3 样式的表单。这些帮助程序是 `form` 和 `form_for` 。

Both types of form helpers have the following features in common:

​	类型的表单帮助程序具有以下共同特征：

- Automatically setting the CSRF authenticity token
  自动设置 CRSF 真实性令牌
- Support for all HTTP methods (PUT, POST, DELETE, etc…)
  支持所有方法（PUT、POST、GET 等）
- [Error Handling
  错误处理](https://gobuffalo.io/documentation/frontend-layer/forms/#error-handling)
- Multipart form support
  Multipart 表单支持
- Customizable input types
  可自定义输入类型
- Pass through HTML tag attributes
  通过标签

## Basic Forms 基本表单 

The `form` helper can be used to generate HTML forms. Since this type of form isn’t attached to any particular “model” all information must be passed as options to the form and it’s methods.

​	 `form` 帮助器可用于生成 HTML 表单。由于此类型的表单未附加到任何特定“模型”，因此所有信息都必须作为选项传递给表单及其方法。

templates/talks/edit.html

OUTPUT
输出

```erb
// templates/talks/edit.html

<%= form({action: talkPath({id: 3}), method: "PUT"}) { %>
  <div class="row">
    <div class="col-md-12">
      <%= f.InputTag({name:"Title", value: talk.Title }) %>
    </div>

    <div class="col-md-6">
      <%= f.TextArea({value: talk.Abstract, hide_label: true }) %>
    </div>

    <div class="col-md-6">
      <%= f.SelectTag({name: "TalkFormatID", value: talk.TalkFormatID, options: talk_formats}) %>
      <%= f.SelectTag({name: "AudienceLevel", value: talk.AudienceLevel, options: audience_levels }) %>
    </div>

    <div class="col-md-12">
      <%= f.TextArea({name: "Description", value: talk.Description, rows: 10}) %>
    </div>
    <div class="col-md-12">
      <%= f.TextArea({notes:"Notes", value: talk.Notes, rows: 10 }) %>
    </div>

  </div>
<% } %>
```

## Model Forms 模型表单 

The `form_for` helper can be used to generate HTML forms for a specified model. This makes the code easier to write, and maintains a level of “consistency” across your application.

​	 `form_for` 帮助器可用于为指定模型生成 HTML 表单。这使得代码更易于编写，并在应用程序中保持一定程度的“一致性”。

The `form_for` helper behaves in a similar matter to the `form` helper, with several key differences.

​	 `form_for` 帮助器的行为与 `form` 帮助器类似，但有几个关键区别。

The first difference is that the `form_for` takes a “model” as a first argument. This “model” only needs to be a `struct` it does not have to be database backed.

​	第一个区别是 `form_for` 将“模型”作为第一个参数。此“模型”只需要是 `struct` ，它不必由数据库支持。

The second difference is in the tag calls the models directly. These tags, such as `InputTag`, take the name of the attribute on the model you want to build a field for, then they take an optional set of options as the second argument.

​	第二个区别在于标签直接调用模型。这些标签（例如 `InputTag` ）采用您要为其构建字段的模型上的属性名称，然后将一组可选选项作为第二个参数。

models/talk.go

templates/talks/edit.html

OUTPUT
输出

```go
// models/talk.go
type Talk struct {
  ID            int          `json:"id" db:"id"`
  CreatedAt     time.Time    `json:"created_at" db:"created_at"`
  UpdatedAt     time.Time    `json:"updated_at" db:"updated_at"`
  UserID        int          `json:"user_id" db:"user_id"`
  Title         string       `json:"title" db:"title"`
  Description   nulls.String `json:"description" db:"description"`
  Notes         nulls.String `json:"notes" db:"notes"`
  ParentID      nulls.Int    `json:"parent_id" db:"parent_id"`
  Abstract      string       `json:"abstract" db:"abstract"`
  AudienceLevel string       `json:"audience_level" db:"audience_level"`
  IsPublic      nulls.Bool   `json:"is_public" db:"is_public"`
  TalkFormatID  int          `json:"talk_format_id" db:"talk_format_id"`
}
```

## Select Tags 选择标签 

To build your `<select>` tags inside forms Tags provide 3 convenient ways to add your `<select>` options: `form.SelectOptions`, `map[string]interface{}` or `[]string`, all of them by passing an `options` field into the `form.SelectTag` options like:

​	要在表单中构建 `<select>` 标签，标签提供了 3 种便捷方式来添加 `<select>` 选项： `form.SelectOptions` 、 `map[string]interface{}` 或 `[]string` ，所有这些方式都通过将 `options` 字段传递到 `form.SelectTag` 选项中，例如：

```erb
<%= f.SelectTag("TalkFormatID", {options: talkFormats}) %>
```

or

```erb
<%= f.SelectTag("TalkFormatID", {options: ["one", "two"]}) %>
```

Which will use the same value for the `value` attribute and the body of the option, or:

​	这将对 `value` 属性和选项的主体使用相同的值，或：

```erb
<%= f.SelectTag("TalkFormatID", {options: {"one": 1, "two": 2}}) %>
```

Which allows us to define the options map inside the view.

​	这允许我们在视图中定义选项映射。

### Selectable Interface 可选择接口 

Another alternative for the select options is to pass a list of structs that meet the `form.Selectable` interface.

​	选择选项的另一种替代方法是传递满足 `form.Selectable` 接口的结构列表。

Which consist of two functions:

​	它由两个函数组成：

```go
//Selectable allows any struct to become an option in the select tag.
type Selectable interface {
  SelectValue() interface{}
  SelectLabel() string
}
```

By implementing this interface tags will call `SelectValue` and `SelectLabel` to get the option Value and Label from implementer.

​	通过实现此接口，标签将调用 `SelectValue` 和 `SelectLabel` 从实现者那里获取选项值和标签。

### Selected 已选择 

Tags will add the `selected` attribute to the option that has the same value than the one it receives on the `value` option of the `form.SelectTag`, so you don’t have to look for the option that has equal value than the selected one manually, p.e:

​	标签会将 `selected` 属性添加到与在 `form.SelectTag` 的 `value` 选项中接收到的值相同的选项，因此您不必手动查找与所选值相等的值，例如：

```erb
<%= f.SelectTag("TalkFormatID", {options: {"one": 1, "two": 2}, value: 2}) %>
```

Produces:

​	生成：

```html
<div class="form-group">
  <label>TalkFormatID</label>
  <select class="form-control" id="talk-TalkFormatID" name="TalkFormatID">
    <option value="1">one</option>
    <option value="2" selected>two</option>
  </select>
</div>
```

And similarly with the `form.SelectOptions` slice:

​	类似地，对于 `form.SelectOptions` 切片：

```erb
<%= f.SelectTag("TalkFormatID", {options: talkFormats, value: 2}) %>
```

## Checkbox Tags 复选框标签 

Tags provide a convenient way to build an HTML `<input>` element with `type="checkbox"`:

​	标签提供了一种便捷的方法来构建一个带有 `type="checkbox"` 的 HTML `<input>` 元素：

```erb
<%= f.CheckboxTag("IsPublic") %>
```

That produces:

​	生成：

```html
<div class="form-group">
  <label>
    <input class="" id="talk-IsPublic" name="IsPublic" type="checkbox" value="true" checked="">
    IsPublic
  </label>
</div>
```

You can easily change the label content with

​	您可以使用轻松更改标签内容

```erb
<%= f.CheckboxTag("IsPublic", {label: "Is the talk public?"}) %>
```

That produces:

​	生成：

```html
<div class="form-group">
  <label>
    <input class="" id="post-IsPublic" name="IsPublic" type="checkbox" value="true" checked="">
     Is the Talk public?
  </label>
</div>
```

### Non-Checked Checkbox Values 未选中的复选框值 

By default when a checkbox is not “checked” no value will be sent to the server. Often, it is useful to send a value indicating a non-checked checkbox. This can be set by passing in a `unchecked` value.

​	默认情况下，当复选框未“选中”时，不会向服务器发送任何值。通常，发送一个指示未选中复选框的值很有用。这可以通过传入 `unchecked` 值来设置。

```erb
<%= f.CheckboxTag("IsPublic", {unchecked: false}) %>
<div class="form-group">
  <label>
    <input id="widget-IsPublic" name="IsPublic" type="checkbox" value="true">
    <input name="IsPublic" type="hidden" value="false"> IsPublic
  </label>
</div>
```

When the form is submitted the `hidden` tag will be posted and the server will see the `false` value.

​	提交表单时，将发布 `hidden` 标签，服务器将看到 `false` 值。

## Error Handling 错误处理 

Both `form` and `form_for` helpers have support for handling errors from the [`github.com/gobuffalo/validate`](https://github.com/gobuffalo/validate) package.

​	Both `form` 和 `form_for` 帮助器都支持处理来自 `github.com/gobuffalo/validate` 包的错误。

In an action simply set a value of type `*validate.Errors` on the context as `errors` and the form helpers will pick it up and add error messages to the appropriate form tags.

​	在操作中，只需将 `*validate.Errors` 类型的某个值设置为 `errors` 上的上下文，表单帮助器就会获取该值并将错误消息添加到相应的表单标记。

actions/widgets.go

templates/widgets/new.html

OUTPUT
输出

```go
// actions/widgets.go
func (v WidgetsResource) Create(c buffalo.Context) error {
  tx := c.Value("tx").(*pop.Connection)
  widget := &models.Widget{}
  if err := c.Bind(widget); err != nil {
    return err
  }
  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(widget)
  if err != nil {
    return errors.WithStack(err)
  }
  if verrs.HasAny() {
    c.Set("widget", widget)
    // Make the errors available inside the html template
    c.Set("errors", verrs)
    return c.Render(422, r.HTML("widgets/new.html"))
  }
  c.Flash().Add("success", "Widget was created successfully")
  return c.Redirect(302, "/widgets/%s", widget.ID)
}
```

## Using Non-Bootstrap Form Helpers 使用非 Bootstrap 表单帮助器 

The default form helpers, `form` and `form_for`, generate forms that are compatible with Bootstrap 3. If this is not for you, you can easily use the non-Bootstrap versions of these helpers.

​	默认表单帮助器 `form` 和 `form_for` 会生成与 Bootstrap 3 兼容的表单。如果您不想要这些表单，可以轻松使用这些帮助器的非 Bootstrap 版本。

Requires Plush version `v3.6.8` or greater

​	需要 Plush 版本 `v3.6.8` 或更高版本

actions/render.go

templates/widgets/new.html

OUTPUT
输出

```go
// actions/render.go
func init() {
  r = render.New(render.Options{
    // ...
    // Add template helpers here:
    Helpers: render.Helpers{
      "form":     plush.FormHelper,
      "form_for": plush.FormForHelper,
    },
    // ...
  })
}
```

## FAQs 常见问题解答 

### How Do I Map a Form to a Model/Struct? 如何将表单映射到模型/结构？ 

See the [Request Binding]({{< ref "/buffalo/requestHandling/requestBinding" >}}) page for more information on request binding.

​	请参阅请求绑定页面以获取有关请求绑定的更多信息。

### Can I Change the Name of the `f` Variable in My Template? 模板中的 `f` 变量的名称可以更改吗？ 

By default the form value inside the block is given the name `f`, however this can be changed when creating the form and passing the `var` option.

​	默认情况下，块内的表单值被赋予名称 `f` ，但是可以在创建表单并传递 `var` 选项时更改此名称。

```erb
<%= form({var: "xyz"}) { %>
  <%= xyz.InputTag({name: "Foo"}) %>
<% } %>
```

### How Do I Create a Multipart Form? 如何创建多部分表单？ 

```erb
<%= form({multipart: true}) { %>
<% } %>
<form enctype="multipart/form-data" method="POST">
</form>
```

### Can I Just Use My Own Form (Without the Use of the Form Helper)? 是否可以使用自己的表单（不使用表单助手）？ 

Yes! You most definitely can create and use your own form! The forms provided from having Buffalo generate your resources are simply a placeholder to get you up and running quickly! It is important to note, however, that asking Buffalo to generate your resources, using the supplied generators, will also generate the resource’s CRUD related routes. This is important to note since the route associated with the UPDATE action makes use of the PUT method and is not a valid value for an HTML form method according to the [HTML Standard](https://www.w3.org/TR/html5/forms.html#association-of-controls-and-forms). That being said, you need to ensure that you structure your form (for editing a resource) to use the POST method to tunnel the HTTP method, while using a hidden input to indicate your intention to make use of the PUT method server side. An example of this would look like the follow:

​	当然可以创建并使用自己的表单！从 Buffalo 生成的表单只是占位符，可帮助您快速启动并运行！但是，需要注意的是，使用提供的生成器让 Buffalo 生成资源时，还会生成资源的 CRUD 相关路由。这一点很重要，因为与 UPDATE 操作关联的路由使用 PUT 方法，而根据 HTML 标准，该方法不是 HTML 表单方法的有效值。也就是说，您需要确保将表单（用于编辑资源）构建为使用 POST 方法来隧道传输 HTTP 方法，同时使用隐藏的输入来指示您打算在服务器端使用 PUT 方法。以下是一个示例：

```html
<form method="POST" ...>
  <input type="hidden" name="_method" value="PUT" />
...
```

#### Can I use CSRF token if I disable SSL? 如果禁用 SSL，是否可以使用 CSRF 令牌？ 

The CSRF token is a secret value that is handled securely to remain valid during cookie-based sessions.

​	CSRF 令牌是一个秘密值，在基于 cookie 的会话期间安全地处理以保持有效。

In development environment you can run smoothly under `http`.

​	在开发环境中，您可以在 `http` 下顺利运行。

If you disable SSL (`https`) and post a form in production environment, you will get the message `CSRF token invalid`.

​	如果您禁用 SSL（ `https` ）并在生产环境中发布表单，您将收到消息 `CSRF token invalid` 。

More details: [PR #1851](https://github.com/gobuffalo/buffalo/pull/1851)

​	更多详细信息：PR #1851

#### How Do I Handle CSRF Tokens If I Use My Own Form? 如果我使用自己的表单，我该如何处理 CSRF 令牌？

If you do decide to use your own forms you are going to need a way to provide the form with the authenticity token. There are two ways to solve this issue.

​	如果您确实决定使用自己的表单，您将需要一种方法来为表单提供真实性令牌。有两种方法可以解决此问题。

The first way is to use the `authenticity_token` directly in form, since it is already in the context.

​	第一种方法是在表单中直接使用 `authenticity_token` ，因为它已经在上下文中。

```html
<form method="POST" ...>
  <input name="authenticity_token" type="hidden" value="<%= authenticity_token %>">
</form>
```

Another way is to write a helper to generate that line of code for you.

​	另一种方法是编写一个帮助器来为您生成该代码行。

```go
"csrf": func(ctx plush.HelperContext) (template.HTML, error) {
  tok, ok := ctx.Value("authenticity_token").(string)
  if !ok {
    return "", fmt.Errorf("expected CSRF token got %T", ctx.Value("authenticity_token"))
  }
  t := tags.New("input", tags.Options{
    "value": tok,
    "type":  "hidden",
    "name":  "authenticity_token",
  })
  return t.HTML(), nil
},
```

Now that you have defined a helper to use in your templates you can use your helper inside your form with `<%= csrf() %>`. So your custom form should end up looking like this:

​	现在您已经定义了一个帮助器以在模板中使用，您可以在表单中使用 `<%= csrf() %>` 来使用您的帮助器。因此，您的自定义表单最终应该如下所示：

```erb
<form method="POST" ...>
  <%= csrf() %>
</form>
```
