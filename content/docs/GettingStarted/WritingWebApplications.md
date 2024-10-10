+++
title = "编写 Web 应用程序"
weight = 13
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false

+++
# Writing Web Applications - 编写 Web 应用程序

> 原文：[https://go.dev/doc/articles/wiki/](https://go.dev/doc/articles/wiki/)
>

## 简介

本教程中涉及的内容：

- 用加载和保存方法创建一个数据结构
- 使用`net/http`包来构建Web应用程序
- 使用`html/template`包来处理HTML模板
- 使用`regexp`包来验证用户的输入
- 使用闭包
  

前提知识：

- 有编程经验
- 了解基本的Web技术（HTTP、HTML）。
- 一些UNIX/DOS命令行知识

## 开始使用

​	目前，您需要有一台FreeBSD、Linux、macOS或Windows机器来运行Go。我们将用`$`来代表命令提示符。

​	安装Go（见[安装说明](../InstallingGo)）。

​	在您的`GOPATH`中为本教程建立一个新的目录，然后`cd`到它：

```shell
$ mkdir gowiki
$ cd gowiki
```

​	创建一个名为`wiki.go`的文件，用您喜欢的编辑器打开它，并添加以下几行：

```go 
package main

import (
    "fmt"
    "os"
)
```

​	我们从 Go 标准库中导入 `fmt` 和 `os` 包。以后，随着我们实现更多的功能，我们将在这个`import`声明中添加更多的包。

## 数据结构

​	让我们从定义数据结构开始。一个`wiki`由一系列相互关联的页面组成，每个页面都有一个标题和一个主体（页面内容）。在这里，我们将`Page`定义为一个结构，有两个字段代表标题和正文。

```go 
type Page struct {
    Title string
    Body  []byte
}
```

​	类型`[]byte`表示 "一个字节切片"。(有关切片的更多信息，请参见 [Slices: usage and internals]({{< ref "/goBlog/2011/GoSlicesUsageAndInternals" >}})。) `Body`元素是`[]byte`而不是`string`，因为这是我们将使用的`io`库所期望的类型，您会在下面看到。

​	`Page`结构描述如何将page数据据存储在内存中。但是持久性存储怎么办呢？我们可以通过在`Page`上创建一个`save`方法来解决这个问题。

```go 
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}
```

​	这个方法的签名是这样的 "这是一个名为`save`的方法，它的接收者是`p`，一个指向`Page`的指针。它不接受任何参数，并返回一个`error`类型的值"。

​	这个方法将把`Page`的`Body`保存到一个文本文件中。为了简单起见，我们将使用`Title`作为文件名。

​	`save`方法返回一个`error`值，因为那是`WriteFile`（一个向文件写入字节切片的标准库函数）的返回类型。`save`方法返回错误值，是为了让应用程序在写入文件时出现任何问题时进行处理。如果一切顺利，`Page.save()`将返回`nil`（指针、接口和其他一些类型的零值）。

​	八进制整数字面量`0600`作为`WriteFile`的第三个参数，表示文件应该以只针对当前用户的读写权限创建。(详见 Unix man page `open(2)`）。

​	除了保存页面之外，我们也想加载页面：

```go 
func loadPage(title string) *Page {
    filename := title + ".txt"
    body, _ := os.ReadFile(filename)
    return &Page{Title: title, Body: body}
}
```

​	函数（`是函数不是方法`）`loadPage`从`title` 参数中构造文件名，将文件内容读入一个新的变量`body`中，并返回一个指向用适当的标题和正文值构造的`Page`字面量指针。

​	函数可以返回多个值。标准库函数 `os.ReadFile` 返回 `[]byte` 和 `error`。在 `loadPage` 中，错误还没有被处理；由下划线 (`_`) 符号代表的 "空白标识符（blank identifier） "被用来丢弃错误的返回值（实质上，将该值赋给了 nothing（空值） ）。

​	但是如果`ReadFile`遇到了错误会怎样呢？例如，该文件可能不存在。我们不应该忽视这样的错误。让我们修改这个函数，以返回`*Page`和`error`。

```go 
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
```

​	这个函数的调用者现在可以检查第二个参数；如果它是`nil`，那么它已经成功加载了一个页面。如果不是，将是一个`error`，可以由调用者处理（详情见[语言规范（language specification）]({{< ref "/langSpec/Errors">}})）。

​	此时，我们有一个简单的数据结构，以及保存到文件和从文件加载的能力。让我们写一个`main`函数来测试我们所写的东西。

```go 
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
```

​	在编译和执行这段代码后，将创建一个名为`TestPage.txt`的文件，其中包含`p1`的内容。然后，该文件将被读入结构体`p2`，并将其`Body`元素打印到屏幕上。

​	您可以这样编译和运行该程序：

```shell
$ go build wiki.go
$ ./wiki
This is a sample Page.
```

(如果您使用的是Windows，您必须输入 "`wiki`"，不带有"`./`"来运行该程序。=> 应该说是在 Windows 的 cmd 的使用，若是在Windows 版的 VS Code 终端或 Powershell 上，则是需要使用 `./wiki`)

​	我们到目前为止所写的代码如下：

```go  title="wiki.go"
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
```

## 介绍`net/http`包（一个中间件）

下面是一个简单的Web服务器的完整工作实例：

```go 
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

​	`main`函数以调用`http.HandleFunc`开始，它告诉`http`包用`handler`来处理所有对Web根（"`/`"）的请求。

​	然后调用`http.ListenAndServe`，指定它应该在任何接口（`":8080"`）上监听8080端口。(现在不要担心它的第二个参数，`nil`）这个函数将会一直阻塞，直到程序被终止。

​	`ListenAndServe`总是返回一个错误，因为它只在发生意外错误时返回。为了记录这个错误，我们用`log.Fatal`来包装这个函数调用。

​	函数`handler`的类型是`http.HandlerFunc`。它接受一个`http.ResponseWriter`和一个`http.Request`作为它的参数。

一个`http.ResponseWriter`值组装了HTTP服务器的响应；通过写入它，我们向HTTP客户端发送数据。

​	一个`http.Request`是代表客户端HTTP请求的数据结构。`r.URL.Path`是请求URL的路径部分。后面的`[1:]`意味着 "从第1个字符到结尾，创建一个`Path`的子切片"。这就从路径名称中去掉了前面的"`/`"。

如果您运行这个程序并访问这个URL：

```
http://localhost:8080/monkeys
```

该程序将呈现一个包含以下内容的页面：

```
Hi there, I love monkeys!
```

## Using `net/http` to serve wiki pages

要使用 `net/http` 包，必须导入它：

```go 
import (
    "fmt"
    "os"
    "log"
    "net/http"
)
```

​	让我们创建一个处理程序，`viewHandler`，它将允许用户查看一个`wiki`页面。它将处理以 "`/view/`" 为前缀的 URL。

```go 
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
```

​	请注意，这里同样使用`_`来忽略`loadPage`的`error`返回值。这样做是为了简单，但通常被认为是不好的做法。我们将在后面讨论这个问题。

​	首先，这个函数从`r.URL.Path`中提取页面标题，即请求URL的路径组成。`Path`被重新分割为`[len("/view/"):]`，以去掉请求路径中的`"/view/"`成分。这是因为路径总是以`"/view/"`开始，这不是页面标题的一部分。

​	然后该函数加载页面数据，用一串简单的HTML格式化页面，并将其写入`w`，即`http.ResponseWriter`。

​	为了使用这个处理程序，我们重写我们的`main`函数，使用`viewHandler`初始化`http`，以处理路径`/view/`下的任何请求。

```go 
func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

​	我们到目前为止所写的代码如下：

```go  title="wiki.go"
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

​	让我们创建一些页面数据（如`test.txt`），编译我们的代码，并尝试为一个`wiki`页面服务。

​	在您的编辑器中打开`test.txt`文件，并在其中保存字符串 "Hello world"（不带引号）。

```shell
$ go build wiki.go
$ ./wiki
```

(如果您使用的是Windows，您必须输入 "`wiki`"，不带有"`./`"来运行该程序。=> 应该说是在 Windows 的 cmd 的使用，若是在Windows 版的 VS Code 终端或 Powershell 上，则是需要使用 `./wiki`)

​	在这个Web服务器运行的情况下，访问`http://localhost:8080/view/test`，应该会出现一个名为 "`test` "的页面，其中包含 "Hello world "的字样，如下图。

![image-20221119151739792](WritingWebApplications_img/image-20221119151739792.png)



## 编辑页面

​	没有编辑页面的能力，wiki 就不是wiki 。让我们创建两个新的处理程序：一个名为`editHandler`，用于显示 "编辑页面 "表单，另一个名为`saveHandler`，用于保存通过表单输入的数据。

​	首先，我们把它们添加到`main()`中。

```go 
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

​	函数 `editHandler` 加载页面（如果它不存在，则创建一个空的 `Page` 结构），并显示一个 HTML 表单。

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

​	这个函数可以正常工作，但所有这些硬编码的HTML都很难看。当然，有一个更好的方法。

## `html/template` 包

​	`html/template` 包是 Go 标准库的一部分。我们可以使用 `html/template` 将 HTML 保存在一个单独的文件中，这样我们就可以在不修改 Go 底层代码的情况下改变我们编辑页面的布局。

​	首先，我们必须将 `html/template` 添加到导入列表中。我们也不会再使用 `fmt`，所以我们必须删除它。

```go 
import (
    "html/template"
    "os"
    "net/http"
)
```

​	让我们创建一个包含HTML表单的模板文件。打开一个名为 `edit.html` 的新文件，并添加以下几行：

```html linenums="1"
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
```

修改 `editHandler` 以使用模板，而不是硬编码的 HTML：

```go
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

​	函数`template.ParseFiles`将读取`edit.html`的内容并返回一个`*template.Template`。

​	方法`t.Execute`执行模板，将生成的HTML写到`http.ResponseWriter`中。`.Title`和`.Body`的`.符号`指的是`p.Title`和`p.Body`。

​	`模板指令（template directives）`被括在`双大括号`中。`printf "%s" .Body` 指令是一个函数调用，它将 `.Body` 输出为一个字符串，而不是一个字节流，与调用 `fmt.Printf` 相同。`html/template`包有助于保证模板动作只生成安全和正确外观的HTML。例如，它自动转义任何大于号（`>`），用`&gt;`代替，以确保用户数据不会破坏表单的HTML。

​	由于我们现在正在使用模板，所以让我们为`viewHandler`创建一个名为`view.html`的模板：

```html linenums="1" title="view.html"
<h1>{{.Title}}</h1>

<p>[<a href="/edit/{{.Title}}">edit</a>]</p>

<div>{{printf "%s" .Body}}</div>
```

​	对`viewHandler`进行相应的修改：

```go 
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}
```

​	请注意，我们在两个处理程序中使用了几乎完全相同的模板代码。让我们通过把模板代码移到自己的函数中来消除这种重复。

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}
```

​	并修改处理程序以使用该函数：

```go 
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}
```

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
```

​	如果我们在`main`中注释掉我们未实现的保存处理程序的注册，我们就可以再次构建和测试我们的程序了。

​	我们到目前为止所写的代码如下：

{{< tabpane text=true >}}

{{< tab header="wiki.go" >}}

```go
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

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

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    //http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

{{< /tab  >}}

{{< tab header="view.html" >}}

```go
<h1>{{.Title}}</h1>

<p>[<a href="/edit/{{.Title}}">edit</a>]</p>

<div>{{printf "%s" .Body}}</div>
```

{{< /tab  >}}

{{< tab header="edit.html" >}}

```go
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
```

{{< /tab  >}}

{{< tab header="test.txt" >}}

```go
Hello world
```

{{< /tab  >}}

{{< /tabpane >}}

## 处理不存在的页面

What if you visit [`/view/APageThatDoesntExist`](http://localhost:8080/view/APageThatDoesntExist)? You'll see a page containing HTML. This is because it ignores the error return value from `loadPage` and continues to try and fill out the template with no data. Instead, if the requested Page doesn't exist, it should redirect the client to the edit Page so the content may be created:

​	如果您访问[`/view/APageThatDoesntExist`](http://localhost:8080/view/APageThatDoesntExist)怎么办？您会看到一个包含HTML的页面。这是因为它忽略了`loadPage`的错误返回值，并继续尝试在没有数据的情况下填充模板。相反，如果请求的页面不存在，它应该把客户端重定向到编辑页面，这样就可以创建内容了：

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

​	`http.Redirect`函数在HTTP响应中添加一个HTTP状态码`http.StatusFound`（302）和一个`Location`响应头。

## 保存页面

​	函数`saveHandler`将处理位于编辑页上的表单的提交。在取消对`main`中相关行的注释后，让我们来实现这个处理程序：

```go
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

​	页面标题（在URL中提供）和表单的唯一字段`Body`被存储在一个新的`Page`中。然后调用`save()`方法将数据写入一个文件，并将客户端重定向到`/view/`页面。

​	由`FormValue`返回的值是`string`类型的。我们必须将该值转换为`[]byte`，然后才能将其放入`Page`结构中。我们使用`[]byte(body)`来进行转换。

## 错误处理

​	在我们的程序中，有几个地方的错误被忽略了。这是不好的做法，尤其是当错误发生时，程序会产生意想不到的行为。一个更好的解决方案是处理错误并向用户返回一个错误信息。这样一来，如果真的出了问题，服务器将完全按照我们想要的方式运行，而用户可以得到通知。

首先，让我们在`renderTemplate`中处理这些错误：

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

​	`http.Error`函数会发送一个指定的HTTP响应码（本例中为 "`Internal Server Error`"）和错误信息。把`它`（指的是`renderTemplate`函数）放在一个单独的函数中的决定已经取得了成效。

现在让我们来修复`saveHandler`：

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

​	在`p.save()`过程中发生的任何错误都会报告给用户。

## 模板缓存

​	这段代码中有一个低效的地方：`renderTemplate`在每次渲染页面的时候都会调用`ParseFiles`。更好的方法是在程序初始化时调用一次`ParseFiles`，将所有模板解析成一个`*Template`。然后我们可以使用[ExecuteTemplate](https://go.dev/pkg/html/template/#Template.ExecuteTemplate)方法来渲染一个特定的模板。

首先我们创建一个名为`templates`的全局变量，并用`ParseFiles`初始化它。

```go 
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
```

​	函数`template.Must`是一个方便的包装器，当传递一个非nil的`error`值时，它就会发生`panic`，否则就会返回未更改的`*Template`。`panic`在这里是合适的；如果模板不能被加载，唯一明智的做法就是退出程序。

​	`ParseFiles`函数接受任意数量的字符串参数，这些参数标识我们的模板文件，并将这些文件解析为以基本文件名命名的模板。如果我们要在我们的程序中添加更多的模板，我们会把它们的名字添加到`ParseFiles`调用的参数中。

​	然后我们修改`renderTemplate`函数，用适当的模板名称调用`templates.ExecuteTemplate`方法：

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

注意，模板名称是模板文件的名称，所以我们必须在`tmpl`参数后加上"`.html`"。

## 验证

​	正如您可能已经观察到的，这个程序有一个严重的安全缺陷：用户可以提供一个任意的路径在服务器上被读/写。为了缓解这个问题，我们可以写一个函数，用`正则表达式`来验证标题。

​	首先，将 `"regexp"`添加到`import`列表中。然后我们可以创建一个全局变量来存储我们的验证表达式：

```go 
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
```

​	函数`regexp.MustCompile`将对正则表达式进行解析和编译，并返回一个`regexp.Regexp`。`MustCompile` 与 `Compile` 的不同之处在于，如果表达式编译失败，它将`panic` ，而 `Compile` 会返回一个`error`作为第二个参数。

​	现在，让我们写一个函数，使用`validPath`表达式来验证路径并提取页面标题：

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

​	如果标题是有效的，它将和一个`nil`错误值一起被返回。如果标题无效，该函数将向HTTP连接写一个 "`404 Not Found` "错误，并向处理程序返回一个错误。要创建一个新的错误，我们必须导入`errors`包。

让我们在每个处理程序中调用`getTitle`：

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

## 介绍函数字面量和闭包

​	在每个处理程序中捕捉错误条件，会引入大量的重复代码。如果我们能把每个处理程序都包在一个函数中，进行验证和错误检查呢？Go的[函数字面量]({{< ref "/langSpec/Expressions#function-literals">}})意义提供了一种强大的抽象功能的方法，可以帮助我们解决这个问题。

首先，我们重写每个处理程序的函数定义，接收一个标题字符串：

```go 
func viewHandler(w http.ResponseWriter, r *http.Request, title string)
func editHandler(w http.ResponseWriter, r *http.Request, title string)
func saveHandler(w http.ResponseWriter, r *http.Request, title string)
```

​	现在让我们定义一个包装函数，它接收一个上述类型的函数，并返回一个`http.HandlerFunc`类型的函数（适合传递给函数`http.HandleFunc`）：

```go 
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Here we will extract the page title from the Request,
        // and call the provided handler 'fn'
    }
}
```

​	返回的函数被称为`闭包（closure ）`，因为它包围了定义在它之外的值。在这种情况下，变量`fn`（`makeHandler`的单一参数）被闭包所包围。变量`fn`将是我们的 save， edit， view处理程序中的一个。

​	现在我们可以从`getTitle`中提取代码，并在这里使用它（做一些小修改）：

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

​	`makeHandler`返回的闭包是一个函数，它接收一个`http.ResponseWriter`和`http.Request`（换句话说，一个`http.HandlerFunc`）。该闭包从请求路径中提取`title`，并使用`validPath` regexp 对其进行验证。如果`title`无效，将使用`http.NotFound`函数向`ResponseWriter`写入一个错误。如果标题有效，封闭的处理函数`fn`将被调用，参数为`ResponseWriter`、`Request`和`title`。

​	现在我们可以用`main`中的`makeHandler`来包装处理函数，然后再将它们注册到`http`包中：

```go 
func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

​	最后，我们从处理函数中删除了对 `getTitle` 的调用，使它们变得更加简单：

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

## 试一试!

最终的代码如下:

{{< tabpane text=true >}}
{{< tab header="wiki.go" >}}

```go  title="wiki.go"
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
    "regexp"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

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

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

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

func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

```

{{< /tab >}}

{{< tab header="view.html" >}}

```html linenums="1"
<h1>{{.Title}}</h1>

<p>[<a href="/edit/{{.Title}}">edit</a>]</p>

<div>{{printf "%s" .Body}}</div>
```

{{< /tab >}}

{{< tab header="edit.html" >}}

```html linenums="1"
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
```

{{< /tab >}}

{{< tab header="test.txt" >}}

```text linenums="1"
Hello world
```

{{< /tab >}}

{{< /tabpane >}}

重新编译代码，并运行该应用程序：

```shell
$ go build wiki.go
$ ./wiki
```

​	访问http://localhost:8080/view/ANewPage，您应该看到页面编辑表单。然后您应该能够输入一些文本，点击 "Save"，并被重定向到新创建的页面。

## 其他任务

这里有一些您可能想自己解决的简单任务：

- 在`tmpl/`中存储模板，在`data/`中存储页面数据。
- 添加一个处理程序，使Web根目录重定向到`/view/FrontPage`。
- 通过使其成为有效的HTML并添加一些CSS规则来美化页面模板。
- 通过将`[PageName]`的实例转换为`<a href="/view/PageName">PageName</a>`实现页面间的链接。(提示：您可以使用`regexp.ReplaceAllFunc`来做这件事)

