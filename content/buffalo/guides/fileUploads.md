+++
title = "文件上传"
date = 2024-02-04T21:16:31+08:00
weight =2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/file-uploads/]({{< ref "/buffalo/guides/fileUploads" >}})

# File Uploads 文件上传 

Since **0.10.3**
自 0.10.3 起



Buffalo allows for the easily handling of files uploaded from a form. Storing those files, such as to disk or S3, is up to you the end developer: Buffalo just gives you easy access to the file from the request.

​	Buffalo 允许轻松处理从表单上传的文件。存储这些文件（例如，存储到磁盘或 S3）由您作为最终开发者决定：Buffalo 只让您可以轻松地从请求中访问文件。

## Configuring the Form 配置表单 

The `f.FileTag` form helper can be used to quickly add a file element to the form. When using this the `enctype` of the form is *automatically* switched to be `multipart/form-data`.

​	 `f.FileTag` 表单帮助器可用于快速向表单添加文件元素。使用此表单时，表单的 `enctype` 会自动切换为 `multipart/form-data` 。

```html
<%= form_for(widget, {action: widgetsPath(), method: "POST"}) { %>
  <%= f.InputTag("Name") %>
  <%= f.FileTag("MyFile") %>
  <button class="btn btn-success" role="submit">Save</button>
  <a href="<%= widgetsPath() %>" class="btn btn-warning" data-confirm="Are you sure?">Cancel</a>
<% } %>
```

## Accessing a Form File 访问表单文件 

In the [`buffalo.Context`](https://godoc.org/github.com/gobuffalo/buffalo#Context) the `c.File` takes a string, the name of the form file parameter and will return a [`binding.File`](https://godoc.org/github.com/gobuffalo/buffalo/binding#File) that can be used to easily retrieve a file from the from.

​	在 `buffalo.Context` 中， `c.File` 采用一个字符串，即表单文件参数的名称，并将返回一个 `binding.File` ，可用于轻松地从表单中检索文件。

```go
func SomeHandler(c buffalo.Context) error {
  // ...
  f, err := c.File("someFile")
  if err != nil {
    return errors.WithStack(err)
  }
  // ...
}
```

## Binding to a Struct 绑定到结构 

The [`c.Bind`](https://godoc.org/github.com/gobuffalo/buffalo#Context) allows form elements to be bound to a struct, but it can also attach uploaded files to the struct. To do this, the type of the struct attribute **must** be a `binding.File` type.

​	 `c.Bind` 允许将表单元素绑定到结构，但它还可以将上传的文件附加到结构。为此，结构属性的类型必须是 `binding.File` 类型。

In the example below you can see a model, which is configured to have a `MyFile` attribute that is of type `binding.File`. There is an `AfterCreate` callback on this example model that saves the file to disk after the model has been successfully saved to the database.

​	在下面的示例中，您可以看到一个模型，该模型被配置为具有类型为 `binding.File` 的 `MyFile` 属性。此示例模型上有一个 `AfterCreate` 回调，该回调在模型成功保存到数据库后将文件保存到磁盘。

```go
// models/widget.go
type Widget struct {
  ID        uuid.UUID    `json:"id" db:"id"`
  CreatedAt time.Time    `json:"created_at" db:"created_at"`
  UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
  Name      string       `json:"name" db:"name"`
  MyFile    binding.File `db:"-" form:"someFile"`
}

func (w *Widget) AfterCreate(tx *pop.Connection) error {
  if !w.MyFile.Valid() {
    return nil
  }
  dir := filepath.Join(".", "uploads")
  if err := os.MkdirAll(dir, 0755); err != nil {
    return errors.WithStack(err)
  }
  f, err := os.Create(filepath.Join(dir, w.MyFile.Filename))
  if err != nil {
    return errors.WithStack(err)
  }
  defer f.Close()
  _, err = io.Copy(f, w.MyFile)
  return err
}
```

**Note:** The `MyFile` attribute is not being saved to the database because of the `db:"-"` struct tag.

​	注意：由于 `db:"-"` 结构标记， `MyFile` 属性未保存到数据库。

## Testing File Uploads 测试文件上传 # HTTP 测试库 （包含在 Buffalo 用于测试的 包中）已更新，其中包括两个新函数： 和 。

The HTTP testing library, [`github.com/gobuffalo/httptest`](https://github.com/gobuffalo/httptest) (which is included in the [`github.com/gobuffalo/suite`](https://github.com/gobuffalo/suite) package that Buffalo uses for testing) has been updated to include two new functions: [`MultiPartPost`](https://godoc.org/github.com/gobuffalo/httptest#Request.MultiPartPost) and [`MultiPartPut`](https://godoc.org/github.com/gobuffalo/httptest#Request.MultiPartPut).

​	这些方法的工作方式与 `github.com/gobuffalo/httptest` 和 `github.com/gobuffalo/suite` 方法相同，但它们提交多部分表单，并且可以接受要上传的文件。

These methods work just like the `Post` and `Put` methods, but instead they submit a multipart form, and can accept files for upload.

​	与 `Post` 和 `Put` 一样， 和 将结构或映射作为第一个参数：这相当于您要发布的 HTML 表单。这些方法采用一个可变的第二个参数 。

Like `Post` and `Put`, `MultiPartPost` and `MultiPartPut`, take a struct, or map, as the first argument: this is the equivalent of the HTML form you would post. The methods take a variadic second argument, [`httptest.File`](https://godoc.org/github.com/gobuffalo/httptest#File).

​	 `Post` 需要表单参数的名称 `Put` ；文件的名称 `MultiPartPost` ；以及 `MultiPartPut` ，大概是您要上传的文件。

A `httptest.File` requires the name of the form parameter, `ParamName`; the name of the file, `FileName`; and an `io.Reader`, presumably the file you want to upload.

​	models/widgets.go

actions/widgets_test.go

actions/widgets.go

models/widgets.go

```go
// actions/widgets_test.go

func (as *ActionSuite) Test_WidgetsResource_Create() {
  // clear out the uploads directory
  os.RemoveAll("./uploads")

  // setup a new Widget
  w := &models.Widget{Name: "Foo"}

  // find the file we want to upload
  r, err := os.Open("./logo.svg")
  as.NoError(err)
  // setup a new httptest.File to hold the file information
  f := httptest.File{
    // ParamName is the name of the form parameter
    ParamName: "someFile",
    // FileName is the name of the file being uploaded
    FileName: r.Name(),
    // Reader is the file that is to be uploaded, any io.Reader works
    Reader: r,
  }

  // Post the Widget and the File(s) to /widgets
  res, err := as.HTML("/widgets").MultiPartPost(w, f)
  as.NoError(err)
  as.Equal(302, res.Code)

  // assert the file exists on disk
  _, err = os.Stat("./uploads/logo.svg")
  as.NoError(err)

  // assert the Widget was saved to the DB correctly
  as.NoError(as.DB.First(w))
  as.Equal("Foo", w.Name)
  as.NotZero(w.ID)
}
```
