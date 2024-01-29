+++
title = "编写 Web 应用程序"
date = 2024-01-29T13:06:30+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/doc/articles/wiki/](https://go.dev/doc/articles/wiki/)

# Writing Web Applications 编写 Web 应用程序

## Introduction 简介

Covered in this tutorial:

​	本教程涵盖的内容：

- Creating a data structure with load and save methods
  创建具有加载和保存方法的数据结构
- Using the `net/http` package to build web applications
  使用 `net/http` 包构建 Web 应用程序
- Using the `html/template` package to process HTML templates
  使用 `html/template` 包处理 HTML 模板
- Using the `regexp` package to validate user input
  使用 `regexp` 包验证用户输入
- Using closures
  使用闭包

Assumed knowledge:

​	假定知识：

- Programming experience
  编程经验
- Understanding of basic web technologies (HTTP, HTML)
  了解基本 Web 技术（HTTP、HTML）
- Some UNIX/DOS command-line knowledge
  一些 UNIX/DOS 命令行知识

## Getting Started 开始

At present, you need to have a FreeBSD, Linux, macOS, or Windows machine to run Go. We will use `$` to represent the command prompt.

​	目前，您需要一台 FreeBSD、Linux、macOS 或 Windows 机器来运行 Go。我们将使用 `$` 来表示命令提示符。

Install Go (see the [Installation Instructions](https://go.dev/doc/install)).

​	安装 Go（请参阅安装说明）。

Make a new directory for this tutorial inside your `GOPATH` and cd to it:

​	在您的 `GOPATH` 内为本教程创建一个新目录并 cd 到它：

```
$ mkdir gowiki
$ cd gowiki
```

Create a file named `wiki.go`, open it in your favorite editor, and add the following lines:

​	创建一个名为 `wiki.go` 的文件，在您喜欢的编辑器中打开它，并添加以下行：

```
package main

import (
    "fmt"
    "os"
)
```

We import the `fmt` and `os` packages from the Go standard library. Later, as we implement additional functionality, we will add more packages to this `import` declaration.

​	我们从 Go 标准库中导入 `fmt` 和 `os` 包。稍后，当我们实现其他功能时，我们将向此 `import` 声明中添加更多包。

## Data Structures 数据结构

Let's start by defining the data structures. A wiki consists of a series of interconnected pages, each of which has a title and a body (the page content). Here, we define `Page` as a struct with two fields representing the title and body.

​	让我们从定义数据结构开始。Wiki 由一系列相互连接的页面组成，每个页面都有一个标题和一个正文（页面内容）。在此，我们将 `Page` 定义为一个具有两个字段的结构，分别表示标题和正文。

```
type Page struct {
    Title string
    Body  []byte
}
```

The type `[]byte` means "a `byte` slice". (See [Slices: usage and internals](https://go.dev/doc/articles/slices_usage_and_internals.html) for more on slices.) The `Body` element is a `[]byte` rather than `string` because that is the type expected by the `io` libraries we will use, as you'll see below.

​	类型 `[]byte` 表示“ `byte` 切片”。（有关切片的更多信息，请参阅切片：用法和内部结构。） `Body` 元素是 `[]byte` 而不是 `string` ，因为这是我们将在下面看到， `io` 库期望的类型。

The `Page` struct describes how page data will be stored in memory. But what about persistent storage? We can address that by creating a `save` method on `Page`:

​	 `Page` 结构描述了页面数据将在内存中存储的方式。但持久性存储呢？我们可以通过在 `Page` 上创建一个 `save` 方法来解决这个问题：

```
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}
```

This method's signature reads: "This is a method named `save` that takes as its receiver `p`, a pointer to `Page` . It takes no parameters, and returns a value of type `error`."

​	此方法的签名为：“这是一个名为 `save` 的方法，它以 `p` 为接收者， `Page` 的指针。它不接受任何参数，并返回类型为 `error` 的值。”

This method will save the `Page`'s `Body` to a text file. For simplicity, we will use the `Title` as the file name.

​	此方法将 `Page` 的 `Body` 保存到文本文件中。为简单起见，我们将使用 `Title` 作为文件名。

The `save` method returns an `error` value because that is the return type of `WriteFile` (a standard library function that writes a byte slice to a file). The `save` method returns the error value, to let the application handle it should anything go wrong while writing the file. If all goes well, `Page.save()` will return `nil` (the zero-value for pointers, interfaces, and some other types).

​	 `save` 方法返回 `error` 值，因为这是 `WriteFile` 的返回类型（一个将字节切片写入文件的标准库函数）。 `save` 方法返回错误值，以便应用程序在写入文件时出现任何问题时进行处理。如果一切顺利， `Page.save()` 将返回 `nil` （指针、接口和其他某些类型的零值）。

The octal integer literal `0600`, passed as the third parameter to `WriteFile`, indicates that the file should be created with read-write permissions for the current user only. (See the Unix man page `open(2)` for details.)

​	八进制整数文字 `0600` 作为第三个参数传递给 `WriteFile` ，表示应仅为当前用户创建具有读写权限的文件。（有关详细信息，请参阅 Unix 手册页 `open(2)` 。）

In addition to saving pages, we will want to load pages, too:

​	除了保存页面外，我们还希望加载页面：

```
func loadPage(title string) *Page {
    filename := title + ".txt"
    body, _ := os.ReadFile(filename)
    return &Page{Title: title, Body: body}
}
```

The function `loadPage` constructs the file name from the title parameter, reads the file's contents into a new variable `body`, and returns a pointer to a `Page` literal constructed with the proper title and body values.

​	函数 `loadPage` 从标题参数构造文件名，将文件内容读入新变量 `body` ，并返回使用适当的标题和正文值构造的 `Page` 文字的指针。

Functions can return multiple values. The standard library function `os.ReadFile` returns `[]byte` and `error`. In `loadPage`, error isn't being handled yet; the "blank identifier" represented by the underscore (`_`) symbol is used to throw away the error return value (in essence, assigning the value to nothing).

​	函数可以返回多个值。标准库函数 `os.ReadFile` 返回 `[]byte` 和 `error` 。在 `loadPage` 中，错误尚未得到处理；由下划线 ( `_` ) 符号表示的“空白标识符”用于丢弃错误返回值（本质上，将值分配给无）。

But what happens if `ReadFile` encounters an error? For example, the file might not exist. We should not ignore such errors. Let's modify the function to return `*Page` and `error`.

​	但是，如果 `ReadFile` 遇到错误会发生什么？例如，文件可能不存在。我们不应该忽略此类错误。让我们修改函数以返回 `*Page` 和 `error` 。

```
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
```

Callers of this function can now check the second parameter; if it is `nil` then it has successfully loaded a Page. If not, it will be an `error` that can be handled by the caller (see the [language specification](https://go.dev/ref/spec#Errors) for details).

​	此函数的调用方现在可以检查第二个参数；如果它是 `nil` ，则它已成功加载页面。如果不是，它将是一个 `error` ，可以由调用方处理（有关详细信息，请参阅语言规范）。

At this point we have a simple data structure and the ability to save to and load from a file. Let's write a `main` function to test what we've written:

​	在这一点上，我们有一个简单的数据结构以及保存到文件和从文件中加载文件的能力。让我们编写一个 `main` 函数来测试我们编写的内容：

```
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
```

After compiling and executing this code, a file named `TestPage.txt` would be created, containing the contents of `p1`. The file would then be read into the struct `p2`, and its `Body` element printed to the screen.

​	编译并执行此代码后，将创建一个名为 `TestPage.txt` 的文件，其中包含 `p1` 的内容。然后将该文件读入结构 `p2` ，并将它的 `Body` 元素打印到屏幕上。

You can compile and run the program like this:

​	您可以像这样编译并运行程序：

```
$ go build wiki.go
$ ./wiki
This is a sample Page.
```

(If you're using Windows you must type "`wiki`" without the "`./`" to run the program.)

​	（如果您使用的是 Windows，您必须键入“ `wiki` ”，不带“ `./` ”才能运行程序。）

[Click here to view the code we've written so far.
单击此处查看我们迄今为止编写的代码。](https://go.dev/doc/articles/wiki/part1.go)

## Introducing the `net/http` package (an interlude) 介绍 `net/http` 包（插曲）

Here's a full working example of a simple web server:

​	这是一个简单的 Web 服务器的完整工作示例：

```
//go:build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

The `main` function begins with a call to `http.HandleFunc`, which tells the `http` package to handle all requests to the web root (`"/"`) with `handler`.

​	 `main` 函数以对 `http.HandleFunc` 的调用开头，该调用告诉 `http` 包使用 `handler` 处理对 Web 根目录（ `"/"` ）的所有请求。

It then calls `http.ListenAndServe`, specifying that it should listen on port 8080 on any interface (`":8080"`). (Don't worry about its second parameter, `nil`, for now.) This function will block until the program is terminated.

​	然后它调用 `http.ListenAndServe` ，指定它应在任何接口（ `":8080"` ）上的端口 8080 上侦听。（现在，不用担心它的第二个参数 `nil` 。）此函数将阻塞，直到程序终止。

`ListenAndServe` always returns an error, since it only returns when an unexpected error occurs. In order to log that error we wrap the function call with `log.Fatal`.

​	 `ListenAndServe` 始终返回一个错误，因为它仅在发生意外错误时才返回。为了记录该错误，我们使用 `log.Fatal` 包装函数调用。

The function `handler` is of the type `http.HandlerFunc`. It takes an `http.ResponseWriter` and an `http.Request` as its arguments.

​	函数 `handler` 的类型为 `http.HandlerFunc` 。它以 `http.ResponseWriter` 和 `http.Request` 作为其参数。

An `http.ResponseWriter` value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.

​	 `http.ResponseWriter` 值组装 HTTP 服务器的响应；通过写入它，我们将数据发送到 HTTP 客户端。

An `http.Request` is a data structure that represents the client HTTP request. `r.URL.Path` is the path component of the request URL. The trailing `[1:]` means "create a sub-slice of `Path` from the 1st character to the end." This drops the leading "/" from the path name.

​	 `http.Request` 是一个表示客户端 HTTP 请求的数据结构。 `r.URL.Path` 是请求 URL 的路径组件。尾随的 `[1:]` 表示“从第 1 个字符到末尾创建 `Path` 的子切片”。这会从路径名称中删除前导“/”。

If you run this program and access the URL:

​	如果您运行此程序并访问 URL：

```
http://localhost:8080/monkeys
```

the program would present a page containing:

​	该程序将显示包含以下内容的页面：

```
Hi there, I love monkeys!
```

## Using `net/http` to serve wiki pages 使用 `net/http` 来提供 wiki 页面

To use the `net/http` package, it must be imported:

​	要使用 `net/http` 包，必须导入它：

```go
import (
    "fmt"
    "os"
    "log"
    "net/http"
)
```

Let's create a handler, `viewHandler` that will allow users to view a wiki page. It will handle URLs prefixed with "/view/".

​	让我们创建一个处理程序 `viewHandler` ，它将允许用户查看 wiki 页面。它将处理以“/view/”为前缀的 URL。

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
```

Again, note the use of `_` to ignore the `error` return value from `loadPage`. This is done here for simplicity and generally considered bad practice. We will attend to this later.

​	同样，请注意使用 `_` 来忽略 `loadPage` 的 `error` 返回值。这样做是为了简单起见，通常被认为是错误的做法。我们稍后会处理这个问题。

First, this function extracts the page title from `r.URL.Path`, the path component of the request URL. The `Path` is re-sliced with `[len("/view/"):]` to drop the leading `"/view/"` component of the request path. This is because the path will invariably begin with `"/view/"`, which is not part of the page's title.

​	首先，此函数从 `r.URL.Path` （请求 URL 的路径组件）中提取页面标题。使用 `[len("/view/"):]` 重新对 `Path` 进行切片，以删除请求路径的前导 `"/view/"` 组件。这是因为路径总是以 `"/view/"` 开头，它不是页面标题的一部分。

The function then loads the page data, formats the page with a string of simple HTML, and writes it to `w`, the `http.ResponseWriter`.

​	然后，该函数加载页面数据，使用简单的 HTML 字符串格式化页面，并将其写入 `w` （ `http.ResponseWriter` ）。

To use this handler, we rewrite our `main` function to initialize `http` using the `viewHandler` to handle any requests under the path `/view/`.

​	要使用此处理程序，我们重写 `main` 函数以使用 `viewHandler` 初始化 `http` ，以便处理路径 `/view/` 下的任何请求。

```go
func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

[Click here to view the code we've written so far.
单击此处查看我们到目前为止编写的代码。](https://go.dev/doc/articles/wiki/part2.go)

Let's create some page data (as `test.txt`), compile our code, and try serving a wiki page.

​	让我们创建一些页面数据（作为 `test.txt` ），编译我们的代码，并尝试提供一个 wiki 页面。

Open `test.txt` file in your editor, and save the string "Hello world" (without quotes) in it.

​	在编辑器中打开 `test.txt` 文件，并在其中保存字符串“Hello world”（不带引号）。

```cmd
$ go build wiki.go
$ ./wiki
```

(If you're using Windows you must type "`wiki`" without the "`./`" to run the program.)

​	（如果您使用的是 Windows，则必须键入“ `wiki` ”而不带“ `./` ”才能运行该程序。）

With this web server running, a visit to `http://localhost:8080/view/test` should show a page titled "test" containing the words "Hello world".

​	在此 Web 服务器运行时，访问 `http://localhost:8080/view/test` 应显示一个名为“test”的页面，其中包含单词“Hello world”。

## Editing Pages 编辑页面

A wiki is not a wiki without the ability to edit pages. Let's create two new handlers: one named `editHandler` to display an 'edit page' form, and the other named `saveHandler` to save the data entered via the form.

​	如果没有编辑页面的能力，wiki 就不是 wiki。让我们创建两个新的处理程序：一个名为 `editHandler` 的处理程序，用于显示“编辑页面”表单，另一个名为 `saveHandler` 的处理程序，用于保存通过表单输入的数据。

First, we add them to `main()`:

​	首先，我们将它们添加到 `main()` ：

```go
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

The function `editHandler` loads the page (or, if it doesn't exist, create an empty `Page` struct), and displays an HTML form.

​	函数 `editHandler` 加载页面（或者，如果页面不存在，则创建一个空的 `Page` 结构），并显示一个 HTML 表单。

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}
```

This function will work fine, but all that hard-coded HTML is ugly. Of course, there is a better way.

​	此函数将正常工作，但所有这些硬编码的 HTML 都不美观。当然，有更好的方法。

## The `html/template` package `html/template` 包

The `html/template` package is part of the Go standard library. We can use `html/template` to keep the HTML in a separate file, allowing us to change the layout of our edit page without modifying the underlying Go code.

​	 `html/template` 包是 Go 标准库的一部分。我们可以使用 `html/template` 将 HTML 保存在单独的文件中，这样我们就可以更改编辑页面的布局，而无需修改底层的 Go 代码。

First, we must add `html/template` to the list of imports. We also won't be using `fmt` anymore, so we have to remove that.

​	首先，我们必须将 `html/template` 添加到导入列表中。我们也不会再使用 `fmt` ，因此我们必须将其删除。

```go
import (
    "html/template"
    "os"
    "net/http"
)
```

Let's create a template file containing the HTML form. Open a new file named `edit.html`, and add the following lines:

​	让我们创建一个包含 HTML 表单的模板文件。打开一个名为 `edit.html` 的新文件，并添加以下行：

```go
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
```

Modify `editHandler` to use the template, instead of the hard-coded HTML:

​	修改 `editHandler` 以使用模板，而不是硬编码的 HTML：

```
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}
```

The function `template.ParseFiles` will read the contents of `edit.html` and return a `*template.Template`.

​	函数 `template.ParseFiles` 将读取 `edit.html` 的内容并返回一个 `*template.Template` 。

The method `t.Execute` executes the template, writing the generated HTML to the `http.ResponseWriter`. The `.Title` and `.Body` dotted identifiers refer to `p.Title` and `p.Body`.

​	方法 `t.Execute` 执行模板，将生成的 HTML 写入 `http.ResponseWriter` 。点标识符 `.Title` 和 `.Body` 引用 `p.Title` 和 `p.Body` 。

Template directives are enclosed in double curly braces. The `printf "%s" .Body` instruction is a function call that outputs `.Body` as a string instead of a stream of bytes, the same as a call to `fmt.Printf`. The `html/template` package helps guarantee that only safe and correct-looking HTML is generated by template actions. For instance, it automatically escapes any greater than sign (`>`), replacing it with `>`, to make sure user data does not corrupt the form HTML.

​	模板指令用双花括号括起来。 `printf "%s" .Body` 指令是一个函数调用，它将 `.Body` 作为字符串输出，而不是字节流，与调用 `fmt.Printf` 相同。 `html/template` 包有助于确保模板操作仅生成安全且外观正确的 HTML。例如，它会自动转义任何大于号 ( `>` )，用 `>` 替换它，以确保用户数据不会破坏表单 HTML。

Since we're working with templates now, let's create a template for our `viewHandler` called `view.html`:

​	既然我们现在正在使用模板，那么让我们为我们的 `viewHandler` 创建一个模板，称为 `view.html` ：

```go
<h1>{{.Title}}</h1>

<p>[<a href="/edit/{{.Title}}">edit</a>]</p>

<div>{{printf "%s" .Body}}</div>
```

Modify `viewHandler` accordingly:

​	相应地修改 `viewHandler` ：

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}
```

Notice that we've used almost exactly the same templating code in both handlers. Let's remove this duplication by moving the templating code to its own function:

​	请注意，我们在两个处理程序中使用了几乎完全相同的模板代码。让我们通过将模板代码移至其自身函数来消除此重复：

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}
```

And modify the handlers to use that function:

​	并修改处理程序以使用该函数：

```
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
```

If we comment out the registration of our unimplemented save handler in `main`, we can once again build and test our program. [Click here to view the code we've written so far.](https://go.dev/doc/articles/wiki/part3.go)

​	如果我们在 `main` 中注释掉未实现的保存处理程序的注册，我们就可以再次构建和测试我们的程序。单击此处查看我们到目前为止编写的代码。

## Handling non-existent pages 处理不存在的页面

What if you visit [`/view/APageThatDoesntExist`](http://localhost:8080/view/APageThatDoesntExist)? You'll see a page containing HTML. This is because it ignores the error return value from `loadPage` and continues to try and fill out the template with no data. Instead, if the requested Page doesn't exist, it should redirect the client to the edit Page so the content may be created:

​	如果您访问 `/view/APageThatDoesntExist` 怎么办？您会看到包含 HTML 的页面。这是因为它忽略了 `loadPage` 的错误返回值，并继续尝试用没有数据填充模板。相反，如果请求的页面不存在，它应该将客户端重定向到编辑页面，以便可以创建内容：

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
```

The `http.Redirect` function adds an HTTP status code of `http.StatusFound` (302) and a `Location` header to the HTTP response.

​	 `http.Redirect` 函数向 HTTP 响应添加 HTTP 状态代码 `http.StatusFound` （302）和 `Location` 标头。

## Saving Pages 保存页面

The function `saveHandler` will handle the submission of forms located on the edit pages. After uncommenting the related line in `main`, let's implement the handler:

​	函数 `saveHandler` 将处理位于编辑页面上的表单的提交。在取消注释 `main` 中的相关行后，我们来实现处理程序：

```
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

The page title (provided in the URL) and the form's only field, `Body`, are stored in a new `Page`. The `save()` method is then called to write the data to a file, and the client is redirected to the `/view/` page.

​	页面标题（在 URL 中提供）和表单的唯一字段 `Body` 存储在一个新的 `Page` 中。然后调用 `save()` 方法将数据写入文件，并将客户端重定向到 `/view/` 页面。

The value returned by `FormValue` is of type `string`. We must convert that value to `[]byte` before it will fit into the `Page` struct. We use `[]byte(body)` to perform the conversion.

​	 `FormValue` 返回的值为 `string` 类型。我们必须将该值转换为 `[]byte` ，然后才能放入 `Page` 结构。我们使用 `[]byte(body)` 执行转换。

## Error handling 错误处理

There are several places in our program where errors are being ignored. This is bad practice, not least because when an error does occur the program will have unintended behavior. A better solution is to handle the errors and return an error message to the user. That way if something does go wrong, the server will function exactly how we want and the user can be notified.

​	我们的程序中有多处忽略了错误。这是一个不好的做法，尤其是当错误发生时，程序的行为将是不可预期的。更好的解决方案是处理错误并向用户返回错误消息。这样，如果出现问题，服务器将按我们想要的方式运行，并且可以通知用户。

First, let's handle the errors in `renderTemplate`:

​	首先，让我们处理 `renderTemplate` 中的错误：

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

The `http.Error` function sends a specified HTTP response code (in this case "Internal Server Error") and error message. Already the decision to put this in a separate function is paying off.

​	 `http.Error` 函数发送指定的 HTTP 响应代码（在本例中为“内部服务器错误”）和错误消息。将此放入单独的函数中的决定已经得到了回报。

Now let's fix up `saveHandler`:

​	现在让我们修复 `saveHandler` ：

```go
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

Any errors that occur during `p.save()` will be reported to the user.

​	 `p.save()` 期间发生的任何错误都将报告给用户。

## Template caching 模板缓存

There is an inefficiency in this code: `renderTemplate` calls `ParseFiles` every time a page is rendered. A better approach would be to call `ParseFiles` once at program initialization, parsing all templates into a single `*Template`. Then we can use the [`ExecuteTemplate`](https://go.dev/pkg/html/template/#Template.ExecuteTemplate) method to render a specific template.

​	此代码中存在低效之处： `renderTemplate` 在每次呈现页面时都会调用 `ParseFiles` 。更好的方法是在程序初始化时调用 `ParseFiles` 一次，将所有模板解析为单个 `*Template` 。然后，我们可以使用 `ExecuteTemplate` 方法来呈现特定模板。

First we create a global variable named `templates`, and initialize it with `ParseFiles`.

​	首先，我们创建一个名为 `templates` 的全局变量，并使用 `ParseFiles` 对其进行初始化。

```
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
```

The function `template.Must` is a convenience wrapper that panics when passed a non-nil `error` value, and otherwise returns the `*Template` unaltered. A panic is appropriate here; if the templates can't be loaded the only sensible thing to do is exit the program.

​	 `template.Must` 函数是一个便利包装器，当传递非零 `error` 值时会引发恐慌，否则返回未更改的 `*Template` 。这里恐慌是合适的；如果无法加载模板，唯一明智的做法是退出程序。

The `ParseFiles` function takes any number of string arguments that identify our template files, and parses those files into templates that are named after the base file name. If we were to add more templates to our program, we would add their names to the `ParseFiles` call's arguments.

​	 `ParseFiles` 函数采用任意数量的字符串参数来标识我们的模板文件，并将这些文件解析为以基本文件名命名的模板。如果要向程序中添加更多模板，我们会将它们的名称添加到 `ParseFiles` 调用参数中。

We then modify the `renderTemplate` function to call the `templates.ExecuteTemplate` method with the name of the appropriate template:

​	然后，我们修改 `renderTemplate` 函数以使用适当模板的名称调用 `templates.ExecuteTemplate` 方法：

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

Note that the template name is the template file name, so we must append `".html"` to the `tmpl` argument.

​	请注意，模板名称是模板文件名，因此我们必须将 `".html"` 附加到 `tmpl` 参数。

## Validation 验证

As you may have observed, this program has a serious security flaw: a user can supply an arbitrary path to be read/written on the server. To mitigate this, we can write a function to validate the title with a regular expression.

​	您可能已经注意到，此程序存在严重的安全漏洞：用户可以提供要在服务器上读取/写入的任意路径。为了缓解此问题，我们可以编写一个函数来使用正则表达式验证标题。

First, add `"regexp"` to the `import` list. Then we can create a global variable to store our validation expression:

​	首先，将 `"regexp"` 添加到 `import` 列表中。然后，我们可以创建一个全局变量来存储我们的验证表达式：

```
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
```

The function `regexp.MustCompile` will parse and compile the regular expression, and return a `regexp.Regexp`. `MustCompile` is distinct from `Compile` in that it will panic if the expression compilation fails, while `Compile` returns an `error` as a second parameter.

​	函数 `regexp.MustCompile` 将解析并编译正则表达式，并返回一个 `regexp.Regexp` 。 `MustCompile` 与 `Compile` 的区别在于，如果表达式编译失败，它会引发恐慌，而 `Compile` 会将 `error` 作为第二个参数返回。

Now, let's write a function that uses the `validPath` expression to validate path and extract the page title:

​	现在，让我们编写一个使用 `validPath` 表达式来验证路径并提取页面标题的函数：

```go
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}
```

If the title is valid, it will be returned along with a `nil` error value. If the title is invalid, the function will write a "404 Not Found" error to the HTTP connection, and return an error to the handler. To create a new error, we have to import the `errors` package.

​	如果标题有效，它将与 `nil` 错误值一起返回。如果标题无效，该函数将向 HTTP 连接写入“404 未找到”错误，并向处理程序返回错误。要创建新错误，我们必须导入 `errors` 包。

Let's put a call to `getTitle` in each of the handlers:

​	让我们在每个处理程序中调用 `getTitle` ：

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err = p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

## Introducing Function Literals and Closures 介绍函数字面量和闭包

Catching the error condition in each handler introduces a lot of repeated code. What if we could wrap each of the handlers in a function that does this validation and error checking? Go's [function literals](https://go.dev/ref/spec#Function_literals) provide a powerful means of abstracting functionality that can help us here.

​	在每个处理程序中捕获错误条件会引入大量重复的代码。如果我们能用一个执行此验证和错误检查的函数包装每个处理程序，该怎么办？Go 的函数字面量提供了一种强大的抽象功能的方法，可以帮助我们解决此问题。

First, we re-write the function definition of each of the handlers to accept a title string:

​	首先，我们重写每个处理程序的函数定义以接受标题字符串：

```go
func viewHandler(w http.ResponseWriter, r *http.Request, title string)
func editHandler(w http.ResponseWriter, r *http.Request, title string)
func saveHandler(w http.ResponseWriter, r *http.Request, title string)
```

Now let's define a wrapper function that *takes a function of the above type*, and returns a function of type `http.HandlerFunc` (suitable to be passed to the function `http.HandleFunc`):

​	现在让我们定义一个包装函数，它接受上述类型的函数，并返回类型为 `http.HandlerFunc` 的函数（适合传递给函数 `http.HandleFunc` ）：

```go
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Here we will extract the page title from the Request,
        // and call the provided handler 'fn'
    }
}
```

The returned function is called a closure because it encloses values defined outside of it. In this case, the variable `fn` (the single argument to `makeHandler`) is enclosed by the closure. The variable `fn` will be one of our save, edit, or view handlers.

​	返回的函数称为闭包，因为它包含在它外部定义的值。在这种情况下，变量 `fn` （ `makeHandler` 的单个参数）被闭包包含。变量 `fn` 将是我们保存、编辑或查看处理程序之一。

Now we can take the code from `getTitle` and use it here (with some minor modifications):

​	现在我们可以从 `getTitle` 中获取代码并在此处使用它（略作修改）：

```go
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}
```

The closure returned by `makeHandler` is a function that takes an `http.ResponseWriter` and `http.Request` (in other words, an `http.HandlerFunc`). The closure extracts the `title` from the request path, and validates it with the `validPath` regexp. If the `title` is invalid, an error will be written to the `ResponseWriter` using the `http.NotFound` function. If the `title` is valid, the enclosed handler function `fn` will be called with the `ResponseWriter`, `Request`, and `title` as arguments.

​	 `makeHandler` 返回的闭包是一个函数，它接受 `http.ResponseWriter` 和 `http.Request` （换句话说，一个 `http.HandlerFunc` ）。闭包从请求路径中提取 `title` ，并使用 `validPath` 正则表达式对其进行验证。如果 `title` 无效，则会使用 `http.NotFound` 函数将错误写入 `ResponseWriter` 。如果 `title` 有效，则会使用 `ResponseWriter` 、 `Request` 和 `title` 作为参数调用封闭的处理程序函数 `fn` 。

Now we can wrap the handler functions with `makeHandler` in `main`, before they are registered with the `http` package:

​	现在，我们可以在将处理程序函数注册到 `http` 包之前，使用 `makeHandler` 在 `main` 中包装它们：

```go
func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Finally we remove the calls to `getTitle` from the handler functions, making them much simpler:

​	最后，我们从处理程序函数中删除对 `getTitle` 的调用，使它们变得更简单：

```go
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

## Try it out! 试试看！

[Click here to view the final code listing.
单击此处查看最终代码清单。](https://go.dev/doc/articles/wiki/final.go)

> 完整代码：
>
> ```go
> // Copyright 2010 The Go Authors. All rights reserved.
> // Use of this source code is governed by a BSD-style
> // license that can be found in the LICENSE file.
> 
> //go:build ignore
> 
> package main
> 
> import (
> 	"html/template"
> 	"log"
> 	"net/http"
> 	"os"
> 	"regexp"
> )
> 
> type Page struct {
> 	Title string
> 	Body  []byte
> }
> 
> func (p *Page) save() error {
> 	filename := p.Title + ".txt"
> 	return os.WriteFile(filename, p.Body, 0600)
> }
> 
> func loadPage(title string) (*Page, error) {
> 	filename := title + ".txt"
> 	body, err := os.ReadFile(filename)
> 	if err != nil {
> 		return nil, err
> 	}
> 	return &Page{Title: title, Body: body}, nil
> }
> 
> func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
> 	p, err := loadPage(title)
> 	if err != nil {
> 		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
> 		return
> 	}
> 	renderTemplate(w, "view", p)
> }
> 
> func editHandler(w http.ResponseWriter, r *http.Request, title string) {
> 	p, err := loadPage(title)
> 	if err != nil {
> 		p = &Page{Title: title}
> 	}
> 	renderTemplate(w, "edit", p)
> }
> 
> func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
> 	body := r.FormValue("body")
> 	p := &Page{Title: title, Body: []byte(body)}
> 	err := p.save()
> 	if err != nil {
> 		http.Error(w, err.Error(), http.StatusInternalServerError)
> 		return
> 	}
> 	http.Redirect(w, r, "/view/"+title, http.StatusFound)
> }
> 
> var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
> 
> func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
> 	err := templates.ExecuteTemplate(w, tmpl+".html", p)
> 	if err != nil {
> 		http.Error(w, err.Error(), http.StatusInternalServerError)
> 	}
> }
> 
> var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
> 
> func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
> 	return func(w http.ResponseWriter, r *http.Request) {
> 		m := validPath.FindStringSubmatch(r.URL.Path)
> 		if m == nil {
> 			http.NotFound(w, r)
> 			return
> 		}
> 		fn(w, r, m[2])
> 	}
> }
> 
> func main() {
> 	http.HandleFunc("/view/", makeHandler(viewHandler))
> 	http.HandleFunc("/edit/", makeHandler(editHandler))
> 	http.HandleFunc("/save/", makeHandler(saveHandler))
> 
> 	log.Fatal(http.ListenAndServe(":8080", nil))
> }
> ```
>
> 

Recompile the code, and run the app:

​	重新编译代码并运行应用：

```
$ go build wiki.go
$ ./wiki
```

Visiting http://localhost:8080/view/ANewPage should present you with the page edit form. You should then be able to enter some text, click 'Save', and be redirected to the newly created page.

​	访问 http://localhost:8080/view/ANewPage 应会显示页面编辑表单。然后，您应该能够输入一些文本，单击“保存”，并重定向到新创建的页面。

## Other tasks 其他任务

Here are some simple tasks you might want to tackle on your own:

​	以下是一些您可能想要自己解决的简单任务：

- Store templates in `tmpl/` and page data in `data/`.
  将模板存储在 `tmpl/` 中，将页面数据存储在 `data/` 中。
- Add a handler to make the web root redirect to `/view/FrontPage`.
  添加一个处理程序，使 Web 根目录重定向到 `/view/FrontPage` 。
- Spruce up the page templates by making them valid HTML and adding some CSS rules.
  通过使页面模板成为有效的 HTML 并添加一些 CSS 规则来整理页面模板。
- Implement inter-page linking by converting instances of `[PageName]` to
  通过将 `[PageName]` 的实例转换为 `[PageName]` 来实现页面间链接。（提示：您可以使用 来执行此操作）
  `<a href="/view/PageName">PageName</a>`. (hint: you could use `regexp.ReplaceAllFunc` to do this)
  go.dev 使用 Google 的 Cookie 来提供和增强其服务质量并分析流量。了解更多。