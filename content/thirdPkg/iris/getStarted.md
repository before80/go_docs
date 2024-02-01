+++
title = "开始入门"
date = 2024-01-31T19:47:27+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://www.iris-go.com/docs/#/](https://www.iris-go.com/docs/#/)

## Installation 安装

Iris is a cross-platform software.

​	Iris 是一个跨平台软件。

The only requirement is the [Go Programming Language](https://go.dev/dl/), version 1.20 and above.

​	唯一的要求是 Go 编程语言 1.20 及更高版本。

```bash
$ mkdir myapp
$ cd myapp
$ go mod init myapp
$ go get github.com/kataras/iris/v12@latest
```

Import it in your code:

​	在您的代码中导入它：

```go
import "github.com/kataras/iris/v12"
```

### Troubleshooting 故障排除

If you get a network error during installation please make sure you set a valid [GOPROXY environment variable](https://github.com/golang/go/wiki/Modules#are-there-always-on-module-repositories-and-enterprise-proxies).

​	如果在安装过程中出现网络错误，请确保设置了有效的 GOPROXY 环境变量。

```sh
go env -w GOPROXY=https://goproxy.io,direct
```

Perform a clean of your go modules cache if none of the above worked:

​	如果上述方法均无效，请清除 go 模块缓存：

```sh
go clean --modcache
```

## Quick start 快速入门

```sh
# assume the following codes in main.go file
$ cat main.go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()

    booksAPI := app.Party("/books")
    {
        booksAPI.Use(iris.Compression)

        // GET: http://localhost:8080/books
        booksAPI.Get("/", list)
        // POST: http://localhost:8080/books
        booksAPI.Post("/", create)
    }

    app.Listen(":8080")
}

// Book example.
type Book struct {
    Title string `json:"title"`
}

func list(ctx iris.Context) {
    books := []Book{
        {"Mastering Concurrency in Go"},
        {"Go Design Patterns"},
        {"Black Hat Go"},
    }

    ctx.JSON(books)
    // TIP: negotiate the response between server's prioritizes
    // and client's requirements, instead of ctx.JSON:
    // 提示：在服务器的优先级和客户端的要求之间进行协商响应，而不是使用 ctx.JSON：
    // ctx.Negotiation().JSON().MsgPack().Protobuf()
    // ctx.Negotiate(books)
}

func create(ctx iris.Context) {
    var b Book
    err := ctx.ReadJSON(&b)
    // TIP: use ctx.ReadBody(&b) to bind
    // any type of incoming data instead.
    // 提示：使用 ctx.ReadBody(&b) 来绑定任何类型的传入数据
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Book creation failure").DetailErr(err))
        // TIP: use ctx.StopWithError(code, err) when only
        // plain text responses are expected on errors.
        // 提示：当仅在错误时期望纯文本响应时，请使用 ctx.StopWithError(code, err)。
        return
    }

    println("Received Book: " + b.Title)

    ctx.StatusCode(iris.StatusCreated)
}
```

**MVC** equivalent:

​	MVC 等价：

```go
import "github.com/kataras/iris/v12/mvc"
```

```go
m := mvc.New(booksAPI)
m.Handle(new(BookController))
```

```go
type BookController struct {
    /* dependencies */
}

// GET: http://localhost:8080/books
func (c *BookController) Get() []Book {
    return []Book{
        {"Mastering Concurrency in Go"},
        {"Go Design Patterns"},
        {"Black Hat Go"},
    }
}

// POST: http://localhost:8080/books
func (c *BookController) Post(b Book) int {
    println("Received Book: " + b.Title)

    return iris.StatusCreated
}
```

**Run** your Iris web server:

​	运行 Iris Web 服务器：

```sh
$ go run main.go
> Now listening on: http://localhost:8080
> Application started. Press CTRL+C to shut down.
```

**List** Books:

​	列出书籍：

```sh
$ curl --header 'Accept-Encoding:gzip' http://localhost:8080/books

[
  {
    "title": "Mastering Concurrency in Go"
  },
  {
    "title": "Go Design Patterns"
  },
  {
    "title": "Black Hat Go"
  }
]
```

**Create** a new Book:

​	创建新书：

```sh
$ curl -i -X POST \
--header 'Content-Encoding:gzip' \
--header 'Content-Type:application/json' \
--data "{\"title\":\"Writing An Interpreter In Go\"}" \
http://localhost:8080/books

> HTTP/1.1 201 Created
```

That's how an **error** response looks like:

​	错误响应如下所示：

```sh
$ curl -X POST --data "{\"title\" \"not valid one\"}" \
http://localhost:8080/books

> HTTP/1.1 400 Bad Request

{
  "status": 400,
  "title": "Book creation failure"
  "detail": "invalid character '\"' after object key",
}
```

[![run in the browser](./getStarted_img/Run-in%20the%20Browser-348798.svgstyle=for-the-badge&logo=repl.svg)](https://replit.com/@kataras/Iris-Hello-World-v1220?v=1)

## Benchmarks 基准

Iris uses a custom version of [muxie](https://github.com/kataras/muxie).

​	Iris 使用 muxie 的自定义版本。 查看所有基准

[See all benchmarks](https://github.com/kataras/server-benchmarks)

📖 Fires 200000 requests with a dynamic parameter of int, sends JSON as request body and receives JSON as response.

​	📖 使用 int 动态参数发送 200000 个请求，将 JSON 作为请求正文发送，并接收 JSON 作为响应。

| Name 名称                                        | Language 语言 | Reqs/sec 请求/秒 | Latency 延迟 | Throughput 吞吐量 | Time To Complete 完成时间 |
| ------------------------------------------------ | ------------- | ---------------- | ------------ | ----------------- | ------------------------- |
| [Iris](https://github.com/kataras/iris)          | Go            | 238954           | 521.69us     | 64.15MB           | 0.84s                     |
| [Gin](https://github.com/gin-gonic/gin)          | Go            | 229665           | 541.96us     | 62.86MB           | 0.87s                     |
| [Chi](https://github.com/go-chi/chi)             | Go            | 228072           | 545.78us     | 62.61MB           | 0.88s                     |
| [Echo ](https://github.com/labstack/echo)        | Go            | 224491           | 553.84us     | 61.70MB           | 0.89s                     |
| [Martini](https://github.com/go-martini/martini) | Go            | 198166           | 627.46us     | 54.47MB           | 1.01s                     |
| [Kestrel](https://github.com/dotnet/aspnetcore)  | C#            | 163486           | 766.90us     | 47.42MB           | 1.23s                     |
| [Buffalo](https://github.com/gobuffalo/buffalo)  | Go            | 102478           | 1.22ms       | 28.14MB           | 1.95s                     |
| [Koa](https://github.com/koajs/koa)              | Javascript    | 48425            | 2.56ms       | 15.39MB           | 4.14s                     |
| [Express](https://github.com/expressjs/express)  | Javascript    | 23622            | 5.25ms       | 9.04MB            | 8.41s                     |

## API Examples API 示例

You can find a number of ready-to-run examples at [Iris examples repository](https://github.com/iris-contrib/examples).

​	您可以在 Iris 示例存储库中找到许多可运行的示例。

### Using GET, POST, PUT, PATCH, DELETE and OPTIONS

```go
func main() {
    // Creates an iris application with default middleware:
    // 创建一个带有默认中间件的 Iris 应用程序：
    // Default with "debug" Logger Level.
    // 默认使用 "debug" 日志级别。
    // Localization enabled on "./locales" directory
    // and HTML templates on "./views" or "./templates" directory.
    // 在 "./locales" 目录启用本地化功能，
    // HTML 模板位于 "./views" 或 "./templates" 目录。
    // It runs with the AccessLog on "./access.log",
    // Recovery (crash-free) and Request ID middleware already attached.
    // 应用程序记录访问日志到 "./access.log"，
    // 已附加了Recovery（无崩溃）和 Request ID 中间件。
    app := iris.Default()

    app.Get("/someGet", getting)
    app.Post("/somePost", posting)
    app.Put("/somePut", putting)
    app.Delete("/someDelete", deleting)
    app.Patch("/somePatch", patching)
    app.Header("/someHead", head)
    app.Options("/someOptions", options)

    app.Listen(":8080")
}
```

### Parameters in path 路径中的参数

```go
func main() {
    app := iris.Default()

    // This handler will match /user/john but will not match /user/ or /user
    // 这个处理程序将匹配 /user/john，但不会匹配 /user/ 或 /user
    app.Get("/user/{name}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        ctx.Writef("Hello %s", name)
    })

    // However, this one will match /user/john/ and also /user/john/send
    // If no other routers match /user/john, it will redirect to /user/john/
    // 然而，这个处理程序将匹配 /user/john/ 以及 /user/john/send。
    // 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/。
    app.Get("/user/{name}/{action:path}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        action := ctx.Params().Get("action")
        message := name + " is " + action
        ctx.WriteString(message)
    })

    // For each matched request Context will hold the route definition
    app.Post("/user/{name:string}/{action:path}", func(ctx iris.Context) {
        ctx.GetCurrentRoute().Tmpl().Src == "/user/{name:string}/{action:path}" // true
    })

    app.Listen(":8080")
}
```

Builtin available parameter types:

​	内置可用参数类型：

| Param Type 参数类型 | Go Type Go 类型                          | Validation 验证                                              | Retrieve Helper 检索助手                                     |
| ------------------- | ---------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `:string`           | string 字符串                            | anything (single path segment) 任何内容（单路径段）          | `Params().Get`                                               |
| `:uuid`             | string 字符串                            | uuidv4 or v1 (single path segment) uuidv4 或 v1（单路径段）  | `Params().Get`                                               |
| `:int`              | int                                      | -9223372036854775808 to 9223372036854775807 (x64) or -2147483648 to 2147483647 (x32), depends on the host arch -9223372036854775808 到 9223372036854775807 (x64) 或 -2147483648 到 2147483647 (x32)，取决于主机体系结构 | `Params().GetInt`                                            |
| `:int8`             | int8                                     | -128 to 127 -128 到 127                                      | `Params().GetInt8`                                           |
| `:int16`            | int16                                    | -32768 to 32767 -32768 到 32767                              | `Params().GetInt16`                                          |
| `:int32`            | int32                                    | -2147483648 to 2147483647 -2147483648 到 2147483647          | `Params().GetInt32`                                          |
| `:int64`            | int64                                    | -9223372036854775808 to 9223372036854775807 -9223372036854775808 到 9223372036854775807 | `Params().GetInt64`                                          |
| `:uint`             | uint                                     | 0 to 18446744073709551615 (x64) or 0 to 4294967295 (x32), depends on the host arch 0 到 18446744073709551615 (x64) 或 0 到 4294967295 (x32)，取决于主机架构 | `Params().GetUint`                                           |
| `:uint8`            | uint8                                    | 0 to 255 0 到 255                                            | `Params().GetUint8`                                          |
| `:uint16`           | uint16                                   | 0 to 65535 0 到 65535                                        | `Params().GetUint16`                                         |
| `:uint32`           | uint32                                   | 0 to 4294967295 0 到 4294967295                              | `Params().GetUint32`                                         |
| `:uint64`           | uint64                                   | 0 to 18446744073709551615 0 到 18446744073709551615          | `Params().GetUint64`                                         |
| `:bool`             | bool                                     | "1" or "t" or "T" or "TRUE" or "true" or "True" or "0" or "f" or "F" or "FALSE" or "false" or "False" "1" 或 "t" 或 "T" 或 "TRUE" 或 "true" 或 "True" 或 "0" 或 "f" 或 "F" 或 "FALSE" 或 "false" 或 "False" | `Params().GetBool`                                           |
| `:alphabetical`     | string 字符串                            | lowercase or uppercase letters 小写或大写字母                | `Params().Get`                                               |
| `:file`             | string 字符串                            | lowercase or uppercase letters, numbers, underscore (_), dash (-), point (.) and no spaces or other special characters that are not valid for filenames 小写或大写字母、数字、下划线 (_)、破折号 (-)、句点 (.)，且没有空格或其他对文件名无效的特殊字符 | `Params().Get`                                               |
| `:path`             | string 字符串                            | anything, can be separated by slashes (path segments) but should be the last part of the route path 任意内容，可以用斜杠 (路径段) 分隔，但应为路由路径的最后一部分 | `Params().Get`                                               |
| `:mail`             | string 字符串                            | Email without domain validation 未经域验证的电子邮件         | `Params().Get`                                               |
| `:email`            | string 字符串                            | Email with domain validation 经过域验证的电子邮件            | `Params().Get`                                               |
| `:date`             | string                                   | yyyy/mm/dd format e.g. /blog/{param:date} matches /blog/2022/04/21 yyyy/mm/dd 格式，例如 /blog/{param:date} 匹配 /blog/2022/04/21 | `Params().GetTime` and `Params().SimpleDate` `Params().GetTime` 和 `Params().SimpleDate` |
| `:weekday`          | uint (0-6) or string uint (0-6) 或字符串 | string of time.Weekday longname format ("sunday" to "monday" or "Sunday" to "Monday") format e.g. /schedule/{param:weekday} matches /schedule/monday time.Weekday longname 格式的字符串（“sunday”到“monday”或“Sunday”到“Monday”）格式，例如 /schedule/{param:weekday} 匹配 /schedule/monday | `Params().GetWeekday`                                        |

More examples can be found at: [_examples/routing](https://github.com/kataras/iris/tree/main/_examples/routing).

​	更多示例可在以下位置找到：_examples/routing。

### Querystring parameters 查询字符串参数

```go
func main() {
    app := iris.Default()

    // Query string parameters are parsed using the existing underlying request object.
    // The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
    app.Get("/welcome", func(ctx iris.Context) {
        firstname := ctx.URLParamDefault("firstname", "Guest")
        lastname := ctx.URLParam("lastname") // shortcut for ctx.Request().URL.Query().Get("lastname")

        ctx.Writef("Hello %s %s", firstname, lastname)
    })
    app.Listen(":8080")
}
```

### Multipart/Urlencoded Form Multipart/Urlencoded 表单

```go
func main() {
    app := iris.Default()

    app.Post("/form_post", func(ctx iris.Context) {
        message := ctx.PostValue("message")
        nick := ctx.PostValueDefault("nick", "anonymous")

        ctx.JSON(iris.Map{
            "status":  "posted",
            "message": message,
            "nick":    nick,
        })
    })
    app.Listen(":8080")
}
```

### Another example: query + post form 另一个示例：查询 + post 表单

```
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=kataras&message=this_is_great
func main() {
    app := iris.Default()

    app.Post("/post", func(ctx iris.Context) {
        id, err := ctx.URLParamInt("id", 0)
        if err != nil {
            ctx.StopWithError(iris.StatusBadRequest, err)
            return
        }

        page := ctx.URLParamIntDefault("page", 0)
        name := ctx.PostValue("name")
        message := ctx.PostValue("message")

        ctx.Writef("id: %d; page: %d; name: %s; message: %s", id, page, name, message)
    })
    app.Listen(":8080")
}
id: 1234; page: 1; name: kataras; message: this_is_great
```

### Query and post form parameters 查询和 post 表单参数

```
POST /post?id=a&id=b&id=c&name=john&name=doe&name=kataras
Content-Type: application/x-www-form-urlencoded
func main() {
    app := iris.Default()

    app.Post("/post", func(ctx iris.Context) {

        ids := ctx.URLParamSlice("id")
        names, err := ctx.PostValues("name")
        if err != nil {
            ctx.StopWithError(iris.StatusBadRequest, err)
            return
        }

        ctx.Writef("ids: %v; names: %v", ids, names)
    })
    app.Listen(":8080")
}
ids: [a b c], names: [john doe kataras]
```

### Upload files 上传文件

#### Single file 单个文件

```go
const maxSize = 8 * iris.MB

func main() {
    app := iris.Default()

    app.Post("/upload", func(ctx iris.Context) {
        // Set a lower memory limit for multipart forms (default is 32 MiB)
        ctx.SetMaxRequestBodySize(maxSize)
        // OR
        // app.Use(iris.LimitRequestBodySize(maxSize))
        // OR
        // OR iris.WithPostMaxMemory(maxSize)

        // single file
        file, fileHeader, err:= ctx.FormFile("file")
        if err != nil {
            ctx.StopWithError(iris.StatusBadRequest, err)
            return
        }

        // Upload the file to specific destination.
        dest := filepath.Join("./uploads", fileHeader.Filename)
        ctx.SaveFormFile(fileHeader, dest)

        ctx.Writef("File: %s uploaded!", fileHeader.Filename)
    })

    app.Listen(":8080")
}
```

How to `curl`:

​	如何 `curl` ：

```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/kataras/test.zip" \
  -H "Content-Type: multipart/form-data"
```

#### Multiple files 多个文件

See the detail [example code](https://github.com/kataras/iris/tree/main/_examples/file-server/upload-files).

​	请参阅详细示例代码。

```go
func main() {
    app := iris.Default()
    app.Post("/upload", func(ctx iris.Context) {
        files, n, err := ctx.UploadFormFiles("./uploads")
        if err != nil {
            ctx.StopWithStatus(iris.StatusInternalServerError)
            return
        }

        ctx.Writef("%d files of %d total size uploaded!", len(files), n))
    })

    app.Listen(":8080", iris.WithPostMaxMemory(8 * iris.MB))
}
```

How to `curl`:

​	如何 `curl` ：

```bash
curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/kataras/test1.zip" \
  -F "upload[]=@/Users/kataras/test2.zip" \
  -H "Content-Type: multipart/form-data"
```

### Grouping routes 分组路由

```go
func main() {
    app := iris.Default()

    // Simple group: v1
    v1 := app.Party("/v1")
    {
        v1.Post("/login", loginEndpoint)
        v1.Post("/submit", submitEndpoint)
        v1.Post("/read", readEndpoint)
    }

    // Simple group: v2
    v2 := app.Party("/v2")
    {
        v2.Post("/login", loginEndpoint)
        v2.Post("/submit", submitEndpoint)
        v2.Post("/read", readEndpoint)
    }

    app.Listen(":8080")
}
```

### Blank Iris without middleware by default 默认情况下，空白 Iris 不带中间件

Use

```go
app := iris.New()
```

instead of

​	而不是

```go
// Default with "debug" Logger Level.
// Localization enabled on "./locales" directory
// and HTML templates on "./views" or "./templates" directory.
// It runs with the AccessLog on "./access.log",
// Recovery and Request ID middleware already attached.
app := iris.Default()
```

### Using middleware 使用中间件

```go
package main

import (
    "github.com/kataras/iris/v12"

    "github.com/kataras/iris/v12/middleware/recover"
)

func main() {
    // Creates an iris application without any middleware by default
    app := iris.New()

    // Global middleware using `UseRouter`.
    //
    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    app.UseRouter(recover.New())

    // Per route middleware, you can add as many as you desire.
    app.Get("/benchmark", MyBenchLogger(), benchEndpoint)

    // Authorization group
    // authorized := app.Party("/", AuthRequired())
    // exactly the same as:
    authorized := app.Party("/")
    // per group middleware! in this case we use the custom created
    // AuthRequired() middleware just in the "authorized" group.
    authorized.Use(AuthRequired())
    {
        authorized.Post("/login", loginEndpoint)
        authorized.Post("/submit", submitEndpoint)
        authorized.Post("/read", readEndpoint)

        // nested group
        testing := authorized.Party("testing")
        testing.Get("/analytics", analyticsEndpoint)
    }

    // Listen and serve on 0.0.0.0:8080
    app.Listen(":8080")
}
```

### Application File Logger 应用程序文件记录器

```go
func main() {
    app := iris.Default()
    // Logging to a file.
    // Colors are automatically disabled when writing to a file.
    f, _ := os.Create("iris.log")
    app.Logger().SetOutput(f)

    // Use the following code if you need to write the logs
    // to file and console at the same time.
    // app.Logger().AddOutput(os.Stdout)

    app.Get("/ping", func(ctx iris.Context) {
        ctx.WriteString("pong")
    })

   app.Listen(":8080")
}
```

### Controlling Log output coloring 控制日志输出着色

By default, logs output on console should be colorized depending on the detected TTY.

​	默认情况下，控制台上的日志输出应根据检测到的 TTY 着色。

Customize level title, text, color and styling at general.

​	在 general 中自定义级别标题、文本、颜色和样式。

Import `golog` and `pio`:

​	导入 `golog` 和 `pio` ：

```go
import (
    "github.com/kataras/golog"
    "github.com/kataras/pio"
    // [...]
)
```

Get a level to customize e.g. `DebugLevel`:

​	获取要自定义的级别，例如 `DebugLevel` ：

```go
level := golog.Levels[golog.DebugLevel]
```

You have full control over his text, title and style:

​	您可以完全控制其文本、标题和样式：

```go
// The Name of the Level
// that named (lowercased) will be used
// to convert a string level on `SetLevel`
// to the correct Level type.
Name string
// AlternativeNames are the names that can be referred to this specific log level.
// i.e Name = "warn"
// AlternativeNames = []string{"warning"}, it's an optional field,
// therefore we keep Name as a simple string and created this new field.
AlternativeNames []string
// Tha Title is the prefix of the log level.
// See `ColorCode` and `Style` too.
// Both `ColorCode` and `Style` should be respected across writers.
Title string
// ColorCode a color for the `Title`.
ColorCode int
// Style one or more rich options for the `Title`.
Style []pio.RichOption
```

Example Code:

​	示例代码：

```go
level := golog.Levels[golog.DebugLevel]
level.Name = "debug" // default
level.Title = "[DBUG]" // default
level.ColorCode = pio.Yellow // default
```

**To change the output format:
要更改输出格式：**

```go
app.Logger().SetFormat("json", "    ")
```

**To register a custom Formatter:
要注册自定义格式化程序：**

```go
app.Logger().RegisterFormatter(new(myFormatter))
```

The [golog.Formatter interface](https://github.com/kataras/golog/blob/master/formatter.go) looks like this:

​	golog.Formatter 接口如下所示：

```go
// Formatter is responsible to print a log to the logger's writer.
type Formatter interface {
    // The name of the formatter.
    String() string
    // Set any options and return a clone,
    // generic. See `Logger.SetFormat`.
    Options(opts ...interface{}) Formatter
    // Writes the "log" to "dest" logger.
    Format(dest io.Writer, log *Log) bool
}
```

**To change the output and the format per level:
要更改每个级别的输出和格式：**

```go
app.Logger().SetLevelOutput("error", os.Stderr)
app.Logger().SetLevelFormat("json")
```

### Request Logging 请求日志记录

The application logger we've seen above it's used to log application-releated information and errors. At the other hand, the Access Logger, we see below, is used to log the incoming HTTP requests and responses.

​	我们上面看到的应用程序记录器用于记录与应用程序相关的的信息和错误。另一方面，我们下面看到的访问记录器用于记录传入的 HTTP 请求和响应。

```go
package main

import (
    "os"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/accesslog"
)

// Read the example and its comments carefully.
func makeAccessLog() *accesslog.AccessLog {
    // Initialize a new access log middleware.
    ac := accesslog.File("./access.log")
    // Remove this line to disable logging to console:
    ac.AddOutput(os.Stdout)

    // The default configuration:
    ac.Delim = '|'
    ac.TimeFormat = "2006-01-02 15:04:05"
    ac.Async = false
    ac.IP = true
    ac.BytesReceivedBody = true
    ac.BytesSentBody = true
    ac.BytesReceived = false
    ac.BytesSent = false
    ac.BodyMinify = true
    ac.RequestBody = true
    ac.ResponseBody = false
    ac.KeepMultiLineError = true
    ac.PanicLog = accesslog.LogHandler

    // Default line format if formatter is missing:
    // Time|Latency|Code|Method|Path|IP|Path Params Query Fields|Bytes Received|Bytes Sent|Request|Response|
    //
    // Set Custom Formatter:
    ac.SetFormatter(&accesslog.JSON{
        Indent:    "  ",
        HumanTime: true,
    })
    // ac.SetFormatter(&accesslog.CSV{})
    // ac.SetFormatter(&accesslog.Template{Text: "{{.Code}}"})

    return ac
}

func main() {
    ac := makeAccessLog()
    defer ac.Close() // Close the underline file.

    app := iris.New()
    // Register the middleware (UseRouter to catch http errors too).
    app.UseRouter(ac.Handler)

    app.Get("/", indexHandler)

    app.Listen(":8080")
}

func indexHandler(ctx iris.Context) {
    ctx.WriteString("OK")
}
```

Read more examples at: [_examples/logging/request-logger](https://github.com/kataras/iris/tree/main/_examples/logging/request-logger).

​	在以下位置阅读更多示例：_examples/logging/request-logger。

### Model binding and validation 模型绑定和验证

To bind a request body into a type, use model binding. We currently support binding of `JSON`, `JSONProtobuf`, `Protobuf`, `MsgPack`, `XML`, `YAML` and standard form values (foo=bar&boo=baz).

​	要将请求正文绑定到类型，请使用模型绑定。我们目前支持绑定 `JSON` 、 `JSONProtobuf` 、 `Protobuf` 、 `MsgPack` 、 `XML` 、 `YAML` 和标准表单值 (foo=bar&boo=baz)。

```go
ReadJSON(outPtr interface{}) error
ReadJSONProtobuf(ptr proto.Message, opts ...ProtoUnmarshalOptions) error
ReadProtobuf(ptr proto.Message) error
ReadMsgPack(ptr interface{}) error
ReadXML(outPtr interface{}) error
ReadYAML(outPtr interface{}) error
ReadForm(formObject interface{}) error
ReadQuery(ptr interface{}) error
```

When using the `ReadBody`, Iris tries to infer the binder depending on the Content-Type header. If you are sure what you are binding, you can use the specific `ReadXXX` methods, e.g. `ReadJSON` or `ReadProtobuf` and e.t.c.

​	使用 `ReadBody` 时，Iris 会尝试根据 Content-Type 标头推断绑定器。如果您确定要绑定什么，可以使用特定的 `ReadXXX` 方法，例如 `ReadJSON` 或 `ReadProtobuf` 等。

```go
ReadBody(ptr interface{}) error
```

Iris, wisely, not features a builtin data validation. However, it does allow you to attach a validator which will automatically called on methods like `ReadJSON`, `ReadXML`.... In this example we will learn how to use the [go-playground/validator/v10](https://www.iris-go.com/docs/(https://github.com/go-playground/validator)) for request body validation.

​	Iris 明智地不提供内置数据验证。但是，它确实允许您附加一个验证器，该验证器将自动在 `ReadJSON` 、 `ReadXML` 等方法上调用.... 在此示例中，我们将学习如何使用 go-playground/validator/v10 进行请求正文验证。

Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set `json:"fieldname"`.

​	请注意，您需要在要绑定的所有字段上设置相应的绑定标记。例如，从 JSON 绑定时，设置 `json:"fieldname"` 。

You can also specify that specific fields are required. If a field is decorated with `binding:"required"` and has a empty value when binding, an error will be returned.

​	您还可以指定特定字段是必需的。如果某个字段用 `binding:"required"` 修饰，并且在绑定时为空值，则会返回错误。

```go
package main

import (
    "fmt"

    "github.com/kataras/iris/v12"
    "github.com/go-playground/validator/v10"
)

func main() {
    app := iris.New()
    app.Validator = validator.New()

    userRouter := app.Party("/user")
    {
        userRouter.Get("/validation-errors", resolveErrorsDocumentation)
        userRouter.Post("/", postUser)
    }
    app.Listen(":8080")
}

// User contains user information.
type User struct {
    FirstName      string     `json:"fname" validate:"required"`
    LastName       string     `json:"lname" validate:"required"`
    Age            uint8      `json:"age" validate:"gte=0,lte=130"`
    Email          string     `json:"email" validate:"required,email"`
    FavouriteColor string     `json:"favColor" validate:"hexcolor|rgb|rgba"`
    Addresses      []*Address `json:"addresses" validate:"required,dive,required"`
}

// Address houses a users address information.
type Address struct {
    Street string `json:"street" validate:"required"`
    City   string `json:"city" validate:"required"`
    Planet string `json:"planet" validate:"required"`
    Phone  string `json:"phone" validate:"required"`
}

type validationError struct {
    ActualTag string `json:"tag"`
    Namespace string `json:"namespace"`
    Kind      string `json:"kind"`
    Type      string `json:"type"`
    Value     string `json:"value"`
    Param     string `json:"param"`
}

func wrapValidationErrors(errs validator.ValidationErrors) []validationError {
    validationErrors := make([]validationError, 0, len(errs))
    for _, validationErr := range errs {
        validationErrors = append(validationErrors, validationError{
            ActualTag: validationErr.ActualTag(),
            Namespace: validationErr.Namespace(),
            Kind:      validationErr.Kind().String(),
            Type:      validationErr.Type().String(),
            Value:     fmt.Sprintf("%v", validationErr.Value()),
            Param:     validationErr.Param(),
        })
    }

    return validationErrors
}

func postUser(ctx iris.Context) {
    var user User
    err := ctx.ReadJSON(&user)
    if err != nil {
        // Handle the error, below you will find the right way to do that...

        if errs, ok := err.(validator.ValidationErrors); ok {
            // Wrap the errors with JSON format, the underline library returns the errors as interface.
            validationErrors := wrapValidationErrors(errs)

            // Fire an application/json+problem response and stop the handlers chain.
            ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
                Title("Validation error").
                Detail("One or more fields failed to be validated").
                Type("/user/validation-errors").
                Key("errors", validationErrors))

            return
        }

        // It's probably an internal JSON error, let's dont give more info here.
        ctx.StopWithStatus(iris.StatusInternalServerError)
        return
    }

    ctx.JSON(iris.Map{"message": "OK"})
}

func resolveErrorsDocumentation(ctx iris.Context) {
    ctx.WriteString("A page that should document to web developers or users of the API on how to resolve the validation errors")
}

```

**Sample request
示例请求**

```json
{
    "fname": "",
    "lname": "",
    "age": 45,
    "email": "mail@example.com",
    "favColor": "#000",
    "addresses": [{
        "street": "Eavesdown Docks",
        "planet": "Persphone",
        "phone": "none",
        "city": "Unknown"
    }]
}
```

**Sample response
示例响应**

```json
{
    "title": "Validation error",
    "detail": "One or more fields failed to be validated",
    "type": "http://localhost:8080/user/validation-errors",
    "status": 400,
    "fields": [
    {
        "tag": "required",
        "namespace": "User.FirstName",
        "kind": "string",
        "type": "string",
        "value": "",
        "param": ""
    },
    {
        "tag": "required",
        "namespace": "User.LastName",
        "kind": "string",
        "type": "string",
        "value": "",
        "param": ""
    }
    ]
}
```

Learn more about model validation at: https://github.com/go-playground/validator/blob/master/_examples

​	在以下网址了解有关模型验证的更多信息：https://github.com/go-playground/validator/blob/master/_examples

### Bind Query String 绑定查询字符串

The `ReadQuery` method only binds the query params and not the post data, use `ReadForm` instead to bind post data.

​	 `ReadQuery` 方法仅绑定查询参数，而不绑定 post 数据，请改用 `ReadForm` 来绑定 post 数据。

```go
package main

import "github.com/kataras/iris/v12"

type Person struct {
    Name    string `url:"name,required"`
    Address string `url:"address"`
}

func main() {
    app := iris.Default()
    app.Any("/", index)
    app.Listen(":8080")
}

func index(ctx iris.Context) {
    var person Person
    if err := ctx.ReadQuery(&person); err!=nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.Application().Logger().Infof("Person: %#+v", person)
    ctx.WriteString("Success")
}
```

### Bind Any 绑定任何内容

Bind request body to "ptr" depending on the content-type that client sends the data, e.g. JSON, XML, YAML, MessagePack, Protobuf, Form and URL Query.

​	根据客户端发送数据的 content-type 将请求正文绑定到“ptr”，例如 JSON、XML、YAML、MessagePack、Protobuf、表单和 URL 查询。

```go
package main

import (
    "time"

    "github.com/kataras/iris/v12"
)

type Person struct {
        Name       string    `form:"name" json:"name" url:"name" msgpack:"name"` 
        Address    string    `form:"address" json:"address" url:"address" msgpack:"address"`
        Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1" json:"birthday" url:"birthday" msgpack:"birthday"`
        CreateTime time.Time `form:"createTime" time_format:"unixNano" json:"create_time" url:"create_time" msgpack:"createTime"`
        UnixTime   time.Time `form:"unixTime" time_format:"unix" json:"unix_time" url:"unix_time" msgpack:"unixTime"`
}

func main() {
    app := iris.Default()
    app.Any("/", index)
    app.Listen(":8080")
}

func index(ctx iris.Context) {
    var person Person
    if err := ctx.ReadBody(&person); err!=nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.Application().Logger().Infof("Person: %#+v", person)
    ctx.WriteString("Success")
}
```

Test it with:

​	使用以下内容进行测试：

```sh
$ curl -X GET "localhost:8085/testing?name=kataras&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
```

### Bind URL Path Parameters 绑定 URL 路径参数

```go
package main

import "github.com/kataras/iris/v12"

type myParams struct {
    Name string   `param:"name"`
    Age  int      `param:"age"`
    Tail []string `param:"tail"`
}
// All parameters are required, as we already know,
// the router will fire 404 if name or int or tail are missing.

func main() {
    app := iris.Default()
    app.Get("/{name}/{age:int}/{tail:path}", func(ctx iris.Context) {
        var p myParams
        if err := ctx.ReadParams(&p); err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return
        }

        ctx.Writef("myParams: %#v", p)
    })
    app.Listen(":8088")
}
```

**Request
请求**

```sh
$ curl -v http://localhost:8080/kataras/27/iris/web/framework
```

### Bind Header 绑定标题

```go
package main

import "github.com/kataras/iris/v12"


type myHeaders struct {
    RequestID      string `header:"X-Request-Id,required"`
    Authentication string `header:"Authentication,required"`
}

func main() {
    app := iris.Default()
    r.GET("/", func(ctx iris.Context) {
        var hs myHeaders
        if err := ctx.ReadHeaders(&hs); err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return
        }

        ctx.JSON(hs)
    })

    app.Listen(":8080")
}
```

**Request
请求**

```sh
curl -H "x-request-id:373713f0-6b4b-42ea-ab9f-e2e04bc38e73" -H "authentication: Bearer my-token" \
http://localhost:8080
```

**Response
响应**

```json
{
  "RequestID": "373713f0-6b4b-42ea-ab9f-e2e04bc38e73",
  "Authentication": "Bearer my-token"
}
```

### Bind HTML checkboxes 绑定 HTML 复选框

```go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))

    app.Get("/", showForm)
    app.Post("/", handleForm)

    app.Listen(":8080")
}

func showForm(ctx iris.Context) {
    if err := ctx.View("form.html"); err!=nil {
        ctx.HTML("<h3>%s</h3>", err.Error())
        return
    }
}

type formExample struct {
    Colors []string `form:"colors[]"` // or just "colors".
}

func handleForm(ctx iris.Context) {
    var form formExample
    err := ctx.ReadForm(&form)
    if err != nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.JSON(iris.Map{"Colors": form.Colors})
}
```

**templates/form.html**

```html
<form action="/" method="POST">
    <p>Check one or more colors</p>

    <label for="red">Red</label>
    <!-- name can be "colors" too -->
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>
```

**Response
响应**

```json
{
  "Colors": [
    "red",
    "green",
    "blue"
  ]
}
```

### JSON, JSONP, XML, Markdown, YAML and MsgPack rendering JSON、JSONP、XML、Markdown、YAML 和 MsgPack 渲染

Detailed examples can be found [here](https://github.com/kataras/iris/tree/main/_examples/response-writer/write-rest).

​	此处可找到详细示例。

```go
func main() {
    app := iris.New()

    // iris.Map is an alias of map[string]interface{}
    app.Get("/json", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
    })

    // Use Secure field to prevent json hijacking.
    // It prepends `"while(1),"` to the body when the data is array.
    app.Get("/json_secure", func(ctx iris.Context) {
        response := []string{"val1", "val2", "val3"}
        options := iris.JSON{Indent: "", Secure: true}
        ctx.JSON(response, options)

        // Will output: while(1);["val1","val2","val3"]
    })

    // Use ASCII field to generate ASCII-only JSON
    // with escaped non-ASCII characters.
    app.Get("/json_ascii", func(ctx iris.Context) {
        response := iris.Map{"lang": "GO-虹膜", "tag": "<br>"}
        options := iris.JSON{Indent: "    ", ASCII: true}
        ctx.JSON(response, options)

        /* Will output:
           {
               "lang": "GO-\u8679\u819c",
               "tag": "\u003cbr\u003e"
           }
        */
    })

    // Normally, JSON replaces special HTML characters with their unicode entities.
    // If you want to encode such characters literally,
    // you SHOULD set the UnescapeHTML field to true.
    app.Get("/json_raw", func(ctx iris.Context) {
        options := iris.JSON{UnescapeHTML: true}
        ctx.JSON(iris.Map{
            "html": "<b>Hello, world!</b>",
        }, options)

        // Will output: {"html":"<b>Hello, world!</b>"}
    })

    app.Get("/json_struct", func(ctx iris.Context) {
        // You also can use a struct.
        var msg struct {
            Name    string `json:"user"`
            Message string
            Number  int
        }
        msg.Name = "Mariah"
        msg.Message = "hello"
        msg.Number = 42
        // Note that msg.Name becomes "user" in the JSON.
        // Will output: {"user": "Mariah", "Message": "hello", "Number": 42}
        ctx.JSON(msg)
    })

    app.Get("/jsonp", func(ctx iris.Context) {
        ctx.JSONP(iris.Map{"hello": "jsonp"}, iris.JSONP{Callback: "callbackName"})
    })

    app.Get("/xml", func(ctx iris.Context) {
        ctx.XML(iris.Map{"message": "hello", "status": iris.StatusOK})
    })

    app.Get("/markdown", func(ctx iris.Context) {
        ctx.Markdown([]byte("# Hello Dynamic Markdown -- iris"))
    })

    app.Get("/yaml", func(ctx iris.Context) {
        ctx.YAML(iris.Map{"message": "hello", "status": iris.StatusOK})
    })

    app.Get("/msgpack", func(ctx iris.Context) {
        u := User{
            Firstname: "John",
            Lastname:  "Doe",
            City:      "Neither FBI knows!!!",
            Age:       25,
        }

        ctx.MsgPack(u)
    })

    // Render using jsoniter instead of the encoding/json:
    app.Listen(":8080", iris.WithOptimizations)
}
```

#### Protobuf

Iris supports native protobuf with `Protobuf` and protobuf to JSON encode and decode.

​	Iris 支持原生 protobuf，并使用 `Protobuf` 和 protobuf 将 JSON 编码和解码。

```go
package main

import (
    "app/protos"

    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.New()

    app.Get("/", send)
    app.Get("/json", sendAsJSON)
    app.Post("/read", read)
    app.Post("/read_json", readFromJSON)

    app.Listen(":8080")
}

func send(ctx iris.Context) {
    response := &protos.HelloReply{Message: "Hello, World!"}
    ctx.Protobuf(response)
}

func sendAsJSON(ctx iris.Context) {
    response := &protos.HelloReply{Message: "Hello, World!"}
    options := iris.JSON{
        Proto: iris.ProtoMarshalOptions{
            AllowPartial: true,
            Multiline:    true,
            Indent:       "    ",
        },
    }

    ctx.JSON(response, options)
}

func read(ctx iris.Context) {
    var request protos.HelloRequest

    err := ctx.ReadProtobuf(&request)
    if err != nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.Writef("HelloRequest.Name = %s", request.Name)
}

func readFromJSON(ctx iris.Context) {
    var request protos.HelloRequest

    err := ctx.ReadJSONProtobuf(&request)
    if err != nil {
        ctx.StopWithError(iris.StatusBadRequest, err)
        return
    }

    ctx.Writef("HelloRequest.Name = %s", request.Name)
}
```

### Serving static files 提供静态文件

```go
func main() {
    app := iris.New()
    app.Favicon("./resources/favicon.ico")
    app.HandleDir("/assets", iris.Dir("./assets"))

    app.Listen(":8080")
}
```

The `HandleDir` method accepts a third, optional argument of `DirOptions`:

​	 `HandleDir` 方法接受第三个可选参数 `DirOptions` ：

```go
type DirOptions struct {
    // Defaults to "/index.html", if request path is ending with **/*/$IndexName
    // then it redirects to **/*(/) which another handler is handling it,
    // that another handler, called index handler, is auto-registered by the framework
    // if end developer does not managed to handle it by hand.
    IndexName string
    // PushTargets filenames (map's value) to
    // be served without additional client's requests (HTTP/2 Push)
    // when a specific request path (map's key WITHOUT prefix)
    // is requested and it's not a directory (it's an `IndexFile`).
    //
    // Example:
    //     "/": {
    //         "favicon.ico",
    //         "js/main.js",
    //         "css/main.css",
    //     }
    PushTargets map[string][]string
    // PushTargetsRegexp like `PushTargets` but accepts regexp which
    // is compared against all files under a directory (recursively).
    // The `IndexName` should be set.
    //
    // Example:
    // "/": regexp.MustCompile("((.*).js|(.*).css|(.*).ico)$")
    // See `iris.MatchCommonAssets` too.
    PushTargetsRegexp map[string]*regexp.Regexp

    // Cache to enable in-memory cache and pre-compress files.
    Cache DirCacheOptions
    // When files should served under compression.
    Compress bool

    // List the files inside the current requested directory if `IndexName` not found.
    ShowList bool
    // If `ShowList` is true then this function will be used instead
    // of the default one to show the list of files of a current requested directory(dir).
    // See `DirListRich` package-level function too.
    DirList DirListFunc

    // Files downloaded and saved locally.
    Attachments Attachments

    // Optional validator that loops through each requested resource.
    AssetValidator func(ctx *context.Context, name string) bool
}
```

Learn more about [file-server](https://github.com/kataras/iris/tree/main/_examples/file-server).

​	了解有关文件服务器的更多信息。

### Serving data from Context 从 Context 提供数据

```go
SendFile(filename string, destinationName string) error
SendFileWithRate(src, destName string, limit float64, burst int) error
```

**Usage
用法**

Force-Send a file to the client:

​	强制向客户端发送文件：

```go
func handler(ctx iris.Context) {
    src := "./files/first.zip"
    ctx.SendFile(src, "client.zip")
}
```

Limit download speed to ~50Kb/s with a burst of 100KB:

​	将下载速度限制为约 50Kb/s，突发速度为 100KB：

```go
func handler(ctx iris.Context) {
    src := "./files/big.zip"
    // optionally, keep it empty to resolve the filename based on the "src".
    dest := "" 

    limit := 50.0 * iris.KB
    burst := 100 * iris.KB
    ctx.SendFileWithRate(src, dest, limit, burst)
}
ServeContent(content io.ReadSeeker, filename string, modtime time.Time)
ServeContentWithRate(content io.ReadSeeker, filename string, modtime time.Time, limit float64, burst int)

ServeFile(filename string) error
ServeFileWithRate(filename string, limit float64, burst int) error
```

**Usage
用法**

```go
func handler(ctx iris.Context) {
    ctx.ServeFile("./public/main.js")
}
```

### Template rendering 模板渲染

Iris supports 8 template engines out-of-the-box, developers can still use any external golang template engine, as `Context.ResponseWriter()` is an `io.Writer`.

​	Iris 开箱即用支持 8 种模板引擎，开发人员仍然可以使用任何外部 golang 模板引擎，因为 `Context.ResponseWriter()` 是 `io.Writer` 。

All template engines share a common API i.e. Parse using embedded assets, Layouts and Party-specific layout, Template Funcs, Partial Render and more.

​	所有模板引擎共享一个通用 API，即使用嵌入式资产、布局和特定于 Party 的布局、模板函数、部分渲染等进行解析。

| #    | Name 名称   | Parser 解析器                                           |
| ---- | ----------- | ------------------------------------------------------- |
| 1    | HTML        | [html/template](https://pkg.go.dev/html/template)       |
| 2    | Blocks 区块 | [kataras/blocks](https://github.com/kataras/blocks)     |
| 3    | Django      | [flosch/pongo2](https://github.com/flosch/pongo2)       |
| 4    | Pug         | [Joker/jade](https://github.com/Joker/jade)             |
| 5    | Handlebars  | [aymerick/raymond](https://github.com/aymerick/raymond) |
| 6    | Amber       | [eknkc/amber](https://github.com/eknkc/amber)           |
| 7    | Jet         | [CloudyKit/jet](https://github.com/CloudyKit/jet)       |
| 8    | Ace         | [yosssi/ace](https://github.com/yosssi/ace)             |

[List of Examples](https://github.com/kataras/iris/tree/main/_examples/view).

​	示例列表。

[List of Benchmarks](https://dev.to/kataras/what-s-the-fastest-template-parser-in-go-4bal).

​	基准列表。

A view engine can be registered per-Party. To **register** a view engine use the `Application/Party.RegisterView(ViewEngine)` method as shown below.

​	可以为每个 Party 注册一个视图引擎。要注册视图引擎，请使用 `Application/Party.RegisterView(ViewEngine)` 方法，如下所示。

Load all templates from the "./views" folder where extension is ".html" and parse them using the standard `html/template` package.

​	从扩展名为“html”的“./views”文件夹中加载所有模板，并使用标准 `html/template` 包对其进行解析。

```go
// [app := iris.New...]
tmpl := iris.HTML("./views", ".html")
app.RegisterView(tmpl)
```

To **render or execute** a view use the `Context.View` method inside the main route's handler.

​	要在主路由的处理程序中呈现或执行视图，请使用 `Context.View` 方法。

```go
if err := ctx.View("hi.html"); err!=nil {
    ctx.HTML("<h3>%s</h3>", err.Error())
    return
}
```

To **bind** Go values with key-value pattern inside a view through middleware or main handler use the `Context.ViewData` method before the `Context.View` one.

​	要通过中间件或主处理程序在视图中使用键值模式绑定 Go 值，请在 `Context.View` 之前使用 `Context.ViewData` 方法。

Bind: `{{.message}}` with `"Hello world!"`.

​	绑定： `{{.message}}` 与 `"Hello world!"` 。

```go
ctx.ViewData("message", "Hello world!")
```

Root binding:

​	根绑定：

```go
if err := ctx.View("user-page.html", User{}); err!=nil {
    ctx.HTML("<h3>%s</h3>", err.Error())
    return
}

// root binding as {{.Name}}
```

To **add a template function** use the `AddFunc` method of the preferred view engine.

​	要添加模板函数，请使用首选视图引擎的 `AddFunc` 方法。

```go
//       func name, input arguments, render value
tmpl.AddFunc("greet", func(s string) string {
    return "Greetings " + s + "!"
})
```

To **reload on every request** call the view engine's `Reload` method.

​	要在每次请求时重新加载，请调用视图引擎的 `Reload` 方法。

```go
tmpl.Reload(true)
```

To use **embedded** templates and not depend on local file system use the [go-bindata](https://github.com/go-bindata/go-bindata) external tool and pass its `AssetFile()` generated function to the first input argument of the preferred view engine.

​	要使用嵌入式模板而不依赖于本地文件系统，请使用 go-bindata 外部工具，并将它生成的 `AssetFile()` 函数传递给首选视图引擎的第一个输入参数。

```go
 tmpl := iris.HTML(AssetFile(), ".html")
```

Example Code:

​	示例代码：

```go
// file: main.go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()

    // Parse all templates from the "./views" folder
    // where extension is ".html" and parse them
    // using the standard `html/template` package.
    tmpl := iris.HTML("./views", ".html")
    // Set custom delimeters.
    tmpl.Delims("{{", "}}")
    // Enable re-build on local template files changes.
    tmpl.Reload(true)

    // Default template funcs are:
    //
    // - {{ urlpath "myNamedRoute" "pathParameter_ifNeeded" }}
    // - {{ render "header.html" . }}
    // and partial relative path to current page:
    // - {{ render_r "header.html" . }} 
    // - {{ yield . }}
    // - {{ current }}
    // Register a custom template func:
    tmpl.AddFunc("greet", func(s string) string {
        return "Greetings " + s + "!"
    })

    // Register the view engine to the views,
    // this will load the templates.
    app.RegisterView(tmpl)

    // Method:    GET
    // Resource:  http://localhost:8080
    app.Get("/", func(ctx iris.Context) {
        // Bind: {{.message}} with "Hello world!"
        ctx.ViewData("message", "Hello world!")
        // Render template file: ./views/hi.html
        if err := ctx.View("hi.html"); err!=nil {
            ctx.HTML("<h3>%s</h3>", err.Error())
            return
        }
    })

    app.Listen(":8080")
}
<!-- file: ./views/hi.html -->
<html>
<head>
    <title>Hi Page</title>
</head>
<body>
    <h1>{{.message}}</h1>
    <strong>{{greet "to you"}}</strong>
</body>
</html>
```

Open a browser tab at [http://localhost:8080](http://localhost:8080/).

​	在 http://localhost:8080 打开一个浏览器选项卡。

The **rendered result** will look like this:

​	呈现的结果将如下所示：

```html
<html>
<head>
    <title>Hi Page</title>
</head>
<body>
    <h1>Hello world!</h1>
    <strong>Greetings to you!</strong>
</body>
</html>
```

### Multitemplate

Iris allows unlimited number of registered view engines per Application. Besides that, you can register a view engine **per Party or through middleware too!**.

​	Iris 允许每个应用程序注册无限数量的视图引擎。除此之外，您还可以为每个 Party 或通过中间件注册视图引擎！

```go
// Register a view engine per group of routes.
adminGroup := app.Party("/admin")
adminGroup.RegisterView(iris.Blocks("./views/admin", ".html"))
```

#### Through Middleware 通过中间件

```go
func middleware(views iris.ViewEngine) iris.Handler {
    return func(ctx iris.Context) {
        ctx.ViewEngine(views)
        ctx.Next()
    }
}
```

**Usage
用法**

```go
// Register a view engine on-fly for the current chain of handlers.
views := iris.Blocks("./views/on-fly", ".html")
views.Load()

app.Get("/", setViews(views), onFly)
```

### Redirects 重定向

Issuing a HTTP redirect is easy. Both internal and external locations are supported. By locations we mean, paths, subdomains, domains and e.t.c.

​	发出 HTTP 重定向很容易。内部和外部位置都受支持。位置是指路径、子域、域等。

#### From Handler 来自处理程序

```go
app.Get("/", func(ctx iris.Context) {
    ctx.Redirect("https://go.dev/dl", iris.StatusMovedPermanently)
})
```

Issuing a HTTP redirect from POST.

​	从 POST 发出 HTTP 重定向。

```go
app.Post("/", func(ctx iris.Context) {
    ctx.Redirect("/login", iris.StatusFound)
})
```

Issuing a local router redirect from a Handler, use `Application.ServeHTTPC` or `Exec()` like below.

​	从处理程序发出本地路由器重定向，请像下面一样使用 `Application.ServeHTTPC` 或 `Exec()` 。

```go
app.Get("/test", func(ctx iris.Context) {
    r := ctx.Request()
    r.URL.Path = "/test2"

    ctx.Application().ServeHTTPC(ctx)
    // OR
    // ctx.Exec("GET", "/test2")
})

app.Get("/test2", func(ctx iris.Context) {
    ctx.JSON(iris.Map{"hello": "world"})
})
```

#### Globally 全局

Use the syntax we all love.

​	使用我们都喜欢的语法。

```go
import "github.com/kataras/iris/v12/middleware/rewrite"
func main() {
    app := iris.New()
    // [...routes]
    redirects := rewrite.Load("redirects.yml")
    app.WrapRouter(redirects)
    app.Listen(":80")
}
```

The `"redirects.yml"` file looks like that:

​	 `"redirects.yml"` 文件如下所示：

```yaml
RedirectMatch:
  # Redirects /seo/* to /*
  - 301 /seo/(.*) /$1

  # Redirects /docs/v12* to /docs
  - 301 /docs/v12(.*) /docs

  # Redirects /old(.*) to /
  - 301 /old(.*) /

  # Redirects http or https://test.* to http or https://newtest.*
  - 301 ^(http|https)://test.(.*) $1://newtest.$2

  # Handles /*.json or .xml as *?format=json or xml,
  # without redirect. See /users route.
  # When Code is 0 then it does not redirect the request,
  # instead it changes the request URL
  # and leaves a route handle the request.
  - 0 /(.*).(json|xml) /$1?format=$2

# Redirects root domain to www.
# Creation of a www subdomain inside the Application is unnecessary,
# all requests are handled by the root Application itself.
PrimarySubdomain: www
```

The full code can be found at the [rewrite middleware example](https://github.com/kataras/iris/tree/main/_examples/routing/rewrite).

​	可以在重写中间件示例中找到完整代码。

### Custom Middleware 自定义中间件

```go
func Logger() iris.Handler {
    return func(ctx iris.Context) {
        t := time.Now()

        // Set a shared variable between handlers
        ctx.Values().Set("framework", "iris")

        // before request

        ctx.Next()

        // after request
        latency := time.Since(t)
        log.Print(latency)

        // access the status we are sending
        status := ctx.GetStatusCode()
        log.Println(status)
    }
}

func main() {
    app := iris.New()
    app.Use(Logger())

    app.Get("/test", func(ctx iris.Context) {
        // retrieve a value set by the middleware.
        framework := ctx.Values().GetString("framework")

        // it would print: "iris"
        log.Println(framework)
    })

    app.Listen(":8080")
}
```

### Using Basic Authentication 使用基本身份验证

HTTP Basic Authentication is the simplest technique for enforcing access controls to web resources because it does not require cookies, session identifiers, or login pages; rather, HTTP Basic authentication uses standard fields in the HTTP header.

​	HTTP 基本身份验证是实施对 Web 资源的访问控制的最简单技术，因为它不需要 cookie、会话标识符或登录页面；相反，HTTP 基本身份验证使用 HTTP 标头中的标准字段。

The Basic Authentication middleware [is included](https://github.com/kataras/iris/tree/main/middleware/basicauth) with the Iris framework, so you do not need to install it separately.

​	Iris 框架中包含基本身份验证中间件，因此您无需单独安装它。

**1.** Import the middleware

​	1. 导入中间件

```go
import "github.com/kataras/iris/v12/middleware/basicauth"
```

**2.** Configure the middleware with its `Options` struct:

​	2. 使用其 `Options` 结构配置中间件：

```go
opts := basicauth.Options{
    Allow: basicauth.AllowUsers(map[string]string{
        "username": "password",
    }),
    Realm:        "Authorization Required",
    ErrorHandler: basicauth.DefaultErrorHandler,
    // [...more options]
}
```

**3.** Initialize the middleware:

​	3. 初始化中间件：

```go
auth := basicauth.New(opts)
```

**3.1** The above steps are the same as the `Default` function:

​	3.1 上述步骤与 `Default` 函数相同：

```go
auth := basicauth.Default(map[string]string{
    "username": "password",
})
```

**3.2** Use a custom slice of Users:

​	3.2 使用自定义的用户切片：

```go
// The struct value MUST contain a Username and Passwords fields
// or GetUsername() string and GetPassword() string methods.
type User struct {
    Username string
    Password string
}

// [...]
auth := basicauth.Default([]User{...})
```

**3.3** Load users from a file optionally, passwords are encrypted using the [pkg.go.dev/golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) package:

​	3.3 从文件中加载用户（可选），使用 pkg.go.dev/golang.org/x/crypto/bcrypt 包对密码进行加密：

```go
auth := basicauth.Load("users.yml", basicauth.BCRYPT)
```

**3.3.1** The same can be achieved using the `Options` (recommended):

​	3.3.1 使用 `Options` 也可以实现相同的功能（推荐）：

```go
opts := basicauth.Options{
    Allow: basicauth.AllowUsersFile("users.yml", basicauth.BCRYPT),
    Realm: basicauth.DefaultRealm,
    // [...more options]
}

auth := basicauth.New(opts)
```

Where the `users.yml` may look like that:

​	其中 `users.yml` 可能如下所示：

```yaml
- username: kataras
  password: $2a$10$Irg8k8HWkDlvL0YDBKLCYee6j6zzIFTplJcvZYKA.B8/clHPZn2Ey
  # encrypted of kataras_pass
  role: admin
- username: makis
  password: $2a$10$3GXzp3J5GhHThGisbpvpZuftbmzPivDMo94XPnkTnDe7254x7sJ3O
  # encrypted of makis_pass
  role: member
```

**4.** Register the middleware:

​	4. 注册中间件：

```go
// Register to all matched routes
// under a Party and its children.
app.Use(auth)

// OR/and register to all http error routes.
app.UseError(auth)

// OR register under a path prefix of a specific Party,
// including all http errors of this path prefix.
app.UseRouter(auth)

// OR register to a specific Route before its main handler.
app.Post("/protected", auth, routeHandler)
```

**5.** Retrieve the username & password:

​	5. 检索用户名和密码：

```go
func routeHandler(ctx iris.Context) {
    username, password, _ := ctx.Request().BasicAuth()
    // [...]
}
```

**5.1** Retrieve the User value (useful when you register a slice of custom user struct at `Options.AllowUsers`):

​	5.1 检索用户值（当您在 `Options.AllowUsers` 处注册自定义用户结构的切片时很有用）：

```go
func routeHandler(ctx iris.Context) {
    user := ctx.User().(*iris.SimpleUser)
    // user.Username
    // user.Password
}
```

Read more authorization and authentication examples at [_examples/auth](https://github.com/kataras/iris/tree/main/_examples/auth).

​	在 _examples/auth 中阅读更多授权和身份验证示例。

### Goroutines inside a middleware 中间件内的 Goroutine

When starting new Goroutines inside a middleware or handler, you **SHOULD NOT** use the original context inside it, you have to use a read-only copy.

​	在中间件或处理程序内启动新的 Goroutine 时，您不应在其中使用原始上下文，您必须使用只读副本。

```go
func main() {
    app := iris.Default()

    app.Get("/long_async", func(ctx iris.Context) {
        // create a clone to be used inside the goroutine
        ctxCopy := ctx.Clone()
        go func() {
            // simulate a long task with time.Sleep(). 5 seconds
            time.Sleep(5 * time.Second)

            // note that you are using the copied context "ctxCopy", IMPORTANT
            log.Printf("Done! in path: %s", ctxCopy.Path())
        }()
    })

    app.Get("/long_sync", func(ctx iris.Context) {
        // simulate a long task with time.Sleep(). 5 seconds
        time.Sleep(5 * time.Second)

        // since we are NOT using a goroutine, we do not have to copy the context
        log.Printf("Done! in path: %s", ctx.Path())
    })

    app.Listen(":8080")
}
```

### Custom HTTP configuration 自定义 HTTP 配置

More than 12 examples about http server configuration can be found at the [_examples/http-server](https://github.com/kataras/iris/tree/main/_examples/http-server) folder.

​	可以在 _examples/http-server 文件夹中找到有关 http 服务器配置的 12 个以上的示例。

Use `http.ListenAndServe()` directly, like this:

​	直接使用 `http.ListenAndServe()` ，如下所示：

```go
func main() {
    app := iris.New()
    // [...routes]
    if err := app.Build(); err!=nil{
        panic(err)
    }
    http.ListenAndServe(":8080", app)
}
```

Note that you **SHOULD** call its `Build` method manually to build the application and the router before using it as an `http.Handler`.

​	请注意，您应该手动调用其 `Build` 方法来构建应用程序和路由器，然后才能将其用作 `http.Handler` 。

Another example:

​	另一个示例：

```go
func main() {
    app := iris.New()
    // [...routes]
    app.Build()

    srv := &http.Server{
        Addr:           ":8080",
        Handler:        app,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    srv.ListenAndServe()
}
```

However, you rarely need an external `http.Server` instance with Iris. You can listen using any tcp listener, http server or a custom function via `Application.Run` method.

​	但是，您很少需要使用 Iris 的外部 `http.Server` 实例。您可以通过 `Application.Run` 方法使用任何 tcp 侦听器、http 服务器或自定义函数进行侦听。

```go
app.Run(iris.Listener(l net.Listener)) // listen using a custom net.Listener
app.Run(iris.Server(srv *http.Server)) // listen using a custom http.Server
app.Run(iris.Addr(addr string)) // the app.Listen is a shortcut of this method.
app.Run(iris.TLS(addr string, certFileOrContents, keyFileOrContents string)) // listen TLS.
app.Run(iris.AutoTLS(addr, domain, email string)) // listen using letsencrypt (see below).

// and any custom function that returns an error:
app.Run(iris.Raw(f func() error))
```

### Socket Sharding 套接字分片

This option allows linear scaling server performance on **multi-CPU servers**. See https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/ for details. Enable with `iris.WithSocketSharding` configurator.

​	此选项允许在多 CPU 服务器上进行线性扩展服务器性能。有关详细信息，请参阅 https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/。使用 `iris.WithSocketSharding` 配置器启用。

*Example Code:
示例代码：*

```go
package main

import (
    "time"

    "github.com/kataras/iris/v12"
)

func main() {
    startup := time.Now()

    app := iris.New()
    app.Get("/", func(ctx iris.Context) {
        s := startup.Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
        ctx.Writef("This server started at: %s\n", s)
    })

    app.Listen(":8080", iris.WithSocketSharding)
    // or app.Run(..., iris.WithSocketSharding)
}
```

### Support Let's Encrypt 支持 Let's Encrypt

Example for 1-line LetsEncrypt HTTPS servers.

​	1 行 LetsEncrypt HTTPS 服务器的示例。

```go
package main

import (
    "log"

    "github.com/iris-gonic/autotls"
    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.Default()

    // Ping handler
    app.Get("/ping", func(ctx iris.Context) {
        ctx.WriteString("pong")
    })

    app.Run(iris.AutoTLS(":443", "example.com example2.com", "mail@example.com"))
}
```

Example for custom TLS (you can bind an autocert manager too):

​	自定义 TLS 的示例（您也可以绑定一个 autocert 管理器）：

```go
app.Run(
    iris.TLS(":443", "", "", func(su *iris.Supervisor) {
        su.Server.TLSConfig = &tls.Config{
            /* your custom fields */
        },
    }),
)
```

> All `iris.Runner` methods such as: Addr, TLS, AutoTLS, Server, Listener and e.t.c accept a variadic input argument of `func(*iris.Supervisor)` to configure the http server instance on build state.
>
> ​	所有 `iris.Runner` 方法（例如：Addr、TLS、AutoTLS、Server、Listener 等）都接受 `func(*iris.Supervisor)` 的变参输入参数，以便在构建状态下配置 http 服务器实例。

### Run multiple service using Iris 使用 Iris 运行多个服务

```go
package main

import (
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"

    "golang.org/x/sync/errgroup"
)

var g errgroup.Group

func startApp1() error {
    app := iris.New().SetName("app1")
    app.Use(recover.New())
    app.Get("/", func(ctx iris.Context) {
        app.Get("/", func(ctx iris.Context) {
            ctx.JSON(iris.Map{
                "code":  iris.StatusOK,
                "message": "Welcome server 1",
            })
        })
    })

    app.Build()
   return app.Listen(":8080")
}

func startApp2() error {
    app := iris.New().SetName("app2")
    app.Use(recover.New())
    app.Get("/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "code":  iris.StatusOK,
            "message": "Welcome server 2",
        })
    })

    return app.Listen(":8081")
}

func main() {
    g.Go(startApp1)
    g.Go(startApp2)

    if err := g.Wait(); err != nil {
        log.Fatal(err)
    }
}
```

Manage multiple Iris instances through the `apps` package. Read more [here](https://github.com/kataras/iris/blob/master/apps/README.md).

​	通过 `apps` 包管理多个 Iris 实例。在此处阅读更多信息。

### Graceful shutdown or restart 优雅地关闭或重新启动

There are a few approaches you can use to perform a graceful shutdown or restart. You can make use of third-party packages specifically built for that, or you can use the `app.Shutdown(context.Context)` method. Examples can be found [here](https://github.com/kataras/iris/tree/main/_examples/http-server/graceful-shutdown).

​	您可以使用几种方法来执行优雅地关闭或重新启动。您可以利用专门为此构建的第三方包，也可以使用 `app.Shutdown(context.Context)` 方法。可以在此处找到示例。

Register an event on CTRL/CMD+C using `iris.RegisterOnInterrupt`:

​	使用 `iris.RegisterOnInterrupt` 在 CTRL/CMD+C 上注册事件：

```go
idleConnsClosed := make(chan struct{})
iris.RegisterOnInterrupt(func() {
    timeout := 10 * time.Second
    ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
    defer cancel()
    // close all hosts.
    app.Shutdown(ctx)
    close(idleConnsClosed)
})

// [...]
app.Listen(":8080", iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
<-idleConnsClosed
```

### Build a single binary with templates 使用模板构建单个二进制文件

You can build a server into a single binary containing templates by using [go-bindata][https://github.com/go-bindata/go-bindata]'s `AssetFile` generated function.

​	您可以通过使用 [go-bindata][ https://github.com/go-bindata/go-bindata] 的 `AssetFile` 生成的函数，将服务器构建到包含模板的单个二进制文件中。

```sh
$ go get -u github.com/go-bindata/go-bindata/...
$ go-bindata -fs -prefix "templates" ./templates/...
$ go run .
```

Example Code:

​	示例代码：

```go
func main() {
    app := iris.New()

    tmpl := iris.HTML(AssetFile(), ".html")
    tmpl.Layout("layouts/layout.html")
    tmpl.AddFunc("greet", func(s string) string {
        return "Greetings " + s + "!"
    })
    app.RegisterView(tmpl)

    // [...]
}
```

See complete examples at the [_examples/view](https://github.com/kataras/iris/tree/main/_examples/view).

​	请参阅 _examples/view 中的完整示例。

### Try to bind body into different structs 尝试将正文绑定到不同的结构

The normal methods for binding request body consumes `ctx.Request().Body` and they cannot be called multiple times, **unless** the `iris.WithoutBodyConsumptionOnUnmarshal` configurator is passed to `app.Run/Listen`.

​	用于绑定请求正文的常规方法会消耗 `ctx.Request().Body` ，并且除非将 `iris.WithoutBodyConsumptionOnUnmarshal` 配置器传递给 `app.Run/Listen` ，否则无法多次调用这些方法。

```go
package main

import "github.com/kataras/iris/v12"

func main() {
    app := iris.New()

    app.Post("/", logAllBody, logJSON, logFormValues, func(ctx iris.Context) {
        // body, err := os.ReadAll(ctx.Request().Body) once or
        body, err := ctx.GetBody() // as many times as you need.
        if err != nil {
            ctx.StopWithError(iris.StatusInternalServerError, err)
            return
        }

        if len(body) == 0 {
            ctx.WriteString(`The body was empty.`)
        } else {
            ctx.WriteString("OK body is still:\n")
            ctx.Write(body)
        }
    })

    app.Listen(":8080", iris.WithoutBodyConsumptionOnUnmarshal)
}

func logAllBody(ctx iris.Context) {
    body, err := ctx.GetBody()
    if err == nil && len(body) > 0 {
        ctx.Application().Logger().Infof("logAllBody: %s", string(body))
    }

    ctx.Next()
}

func logJSON(ctx iris.Context) {
    var p interface{}
    if err := ctx.ReadJSON(&p); err == nil {
        ctx.Application().Logger().Infof("logJSON: %#+v", p)
    }

    ctx.Next()
}

func logFormValues(ctx iris.Context) {
    values := ctx.FormValues()
    if values != nil {
        ctx.Application().Logger().Infof("logFormValues: %v", values)
    }

    ctx.Next()
}
```

You can use the `ReadBody` to bind a struct to a request based on the client's content-type. You can also use [Content Negotiation](https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation). Here's a full example:

​	您可以使用 `ReadBody` 根据客户端的内容类型将结构绑定到请求。您还可以使用内容协商。这是一个完整的示例：

```go
package main

import (
    "github.com/kataras/iris/v12"
)

func main() {
    app := newApp()
    // See main_test.go for usage.
    app.Listen(":8080")
}

func newApp() *iris.Application {
    app := iris.New()
    // To automatically decompress using gzip:
    // app.Use(iris.GzipReader)

    app.Use(setAllowedResponses)

    app.Post("/", readBody)

    return app
}

type payload struct {
    Message string `json:"message" xml:"message" msgpack:"message" yaml:"Message" url:"message" form:"message"`
}

func readBody(ctx iris.Context) {
    var p payload

    // Bind request body to "p" depending on the content-type that client sends the data,
    // e.g. JSON, XML, YAML, MessagePack, Protobuf, Form and URL Query.
    err := ctx.ReadBody(&p)
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest,
            iris.NewProblem().Title("Parser issue").Detail(err.Error()))
        return
    }

    // For the sake of the example, log the received payload.
    ctx.Application().Logger().Infof("Received: %#+v", p)

    // Send back the payload depending on the accept content type and accept-encoding of the client,
    // e.g. JSON, XML and so on.
    ctx.Negotiate(p)
}

func setAllowedResponses(ctx iris.Context) {
    // Indicate that the Server can send JSON, XML, YAML and MessagePack for this request.
    ctx.Negotiation().JSON().XML().YAML().MsgPack()
    // Add more, allowed by the server format of responses, mime types here...

    // If client is missing an "Accept: " header then default it to JSON.
    ctx.Negotiation().Accept.JSON()

    ctx.Next()
}
```

### HTTP2 server push HTTP2 服务器推送

Full example code can be found at [_examples/response-writer/http2push](https://github.com/kataras/iris/tree/main/_examples/response-writer/http2push).

​	可以在 _examples/response-writer/http2push 中找到完整的示例代码。

Server push lets the server preemptively "push" website assets to the client without the user having explicitly asked for them. When used with care, we can send what we know the user is going to need for the page they're requesting.

​	服务器推送允许服务器在用户明确要求之前主动“推送”网站资产给客户端。如果谨慎使用，我们可以发送我们知道用户在请求的页面中需要的内容。

```go
package main

import (
    "net/http"

    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.New()
    app.Get("/", pushHandler)
    app.Get("/main.js", simpleAssetHandler)

    app.Run(iris.TLS("127.0.0.1:443", "mycert.crt", "mykey.key"))
    // $ openssl req -new -newkey rsa:4096 -x509 -sha256 \
    // -days 365 -nodes -out mycert.crt -keyout mykey.key
}

func pushHandler(ctx iris.Context) {
    // The target must either be an absolute path (like "/path") or an absolute
    // URL that contains a valid host and the same scheme as the parent request.
    // If the target is a path, it will inherit the scheme and host of the
    // parent request.
    target := "/main.js"

    if pusher, ok := ctx.ResponseWriter().Naive().(http.Pusher); ok {
        err := pusher.Push(target, nil)
        if err != nil {
            if err == iris.ErrPushNotSupported {
                ctx.StopWithText(iris.StatusHTTPVersionNotSupported, "HTTP/2 push not supported.")
            } else {
                ctx.StopWithError(iris.StatusInternalServerError, err)
            }
            return
        }
    }

    ctx.HTML(`<html><body><script src="%s"></script></body></html>`, target)
}

func simpleAssetHandler(ctx iris.Context) {
    ctx.ServeFile("./public/main.js")
}
```

### Set and get a cookie 设置和获取 Cookie

Secure cookies, encoding and decoding, sessions (and sessions scaling), flash messages and more can be found at the [_examples/cookies](https://github.com/kataras/iris/tree/main/_examples/cookies) and [_examples/sessions](https://github.com/kataras/iris/tree/main/_examples/sessions) directories respectfully.

​	可以在 _examples/cookies 和 _examples/sessions 目录中分别找到安全 Cookie、编码和解码、会话（和会话扩展）、闪存消息等。

```go
import "github.com/kataras/iris/v12"

func main() {
    app := iris.Default()

    app.Get("/cookie", func(ctx iris.Context) {
        value := ctx.GetCookie("my_cookie")

        if value == "" {
            value = "NotSet"
            ctx.SetCookieKV("my_cookie", value)
            // Alternatively: ctx.SetCookie(&http.Cookie{...})
            ctx.SetCookie("", "test", 3600, "/", "localhost", false, true)
        }

        ctx.Writef("Cookie value: %s \n", cookie)
    })

    app.Listen(":8080")
}
```

If you want to set custom the path:

​	如果您想设置自定义路径：

```go
ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))
```

If you want to be visible only to current request path:

​	如果您只想对当前请求路径可见：

```go
ctx.SetCookieKV(name, value, iris.CookieCleanPath /* or iris.CookiePath("") */)
```

More:

​	更多信息：

- `iris.CookieAllowReclaim`
- `iris.CookieAllowSubdomains`
- `iris.CookieSecure`
- `iris.CookieHTTPOnly`
- `iris.CookieSameSite`
- `iris.CookiePath`
- `iris.CookieCleanPath`
- `iris.CookieExpires`
- `iris.CookieEncoding`

You can add cookie options for the whole request in a middleware too:

​	您还可以在中间件中为整个请求添加 Cookie 选项：

```go
func setCookieOptions(ctx iris.Context) {
    ctx.AddCookieOptions(iris.CookieHTTPOnly(true), iris.CookieExpires(1*time.Hour))
    ctx.Next()
}
```

## JSON Web Tokens JSON Web 令牌

JSON Web Token (JWT) is an open standard ([RFC 7519](https://tools.ietf.org/html/rfc7519)) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed. JWTs can be signed using a secret (with the HMAC algorithm) or a public/private key pair using RSA or ECDSA.

​	JSON Web 令牌 (JWT) 是一种开放标准 (RFC 7519)，它定义了一种紧凑且独立的方式，可以作为 JSON 对象在各方之间安全地传输信息。此信息可以被验证和信任，因为它经过数字签名。JWT 可以使用密钥（采用 HMAC 算法）或使用 RSA 或 ECDSA 的公钥/私钥对进行签名。

### When should you use JSON Web Tokens? 您应该在何时使用 JSON Web 令牌？

Here are some scenarios where JSON Web Tokens are useful:

​	以下是一些 JSON Web 令牌有用的场景：

**Authorization**: This is the most common scenario for using JWT. Once the user is logged in, each subsequent request will include the JWT, allowing the user to access routes, services, and resources that are permitted with that token. Single Sign On is a feature that widely uses JWT nowadays, because of its small overhead and its ability to be easily used across different domains.

​	授权：这是使用 JWT 最常见的场景。用户登录后，每个后续请求都将包含 JWT，允许用户访问使用该令牌允许的路由、服务和资源。单点登录是如今广泛使用 JWT 的一项功能，因为它开销小，并且能够轻松地在不同域之间使用。

**Information Exchange**: JSON Web Tokens are a good way of securely transmitting information between parties. Because JWTs can be signed—for example, using public/private key pairs—you can be sure the senders are who they say they are. Additionally, as the signature is calculated using the header and the payload, you can also verify that the content hasn't been tampered with.

​	信息交换：JSON Web 令牌是安全地在各方之间传输信息的一种好方法。因为 JWT 可以签名——例如，使用公钥/私钥对——您可以确保发送者是他们所说的身份。此外，由于签名是使用标头和有效负载计算的，因此您还可以验证内容是否未被篡改。

> Read more about JWT at: https://jwt.io/introduction/
>
> ​	在以下网址阅读有关 JWT 的更多信息：https://jwt.io/introduction/

### Using JWT with Iris 将 JWT 与 Iris 配合使用

The Iris JWT [middleware](https://github.com/kataras/iris/tree/main/middleware/jwt) was designed with security, performance and simplicity in mind, it protects your tokens from [critical vulnerabilities that you may find in other libraries](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/). It is based on [kataras/jwt](https://github.com/kataras/jwt) package.

​	Iris JWT 中间件在设计时考虑了安全性、性能和简单性，它可以保护您的令牌免受您可能在其他库中发现的关键漏洞的影响。它基于 kataras/jwt 包。

```go
package main

import (
    "time"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/jwt"
)

var (
    sigKey = []byte("signature_hmac_secret_shared_key")
    encKey = []byte("GCM_AES_256_secret_shared_key_32")
)

type fooClaims struct {
    Foo string `json:"foo"`
}

func main() {
    app := iris.New()

    signer := jwt.NewSigner(jwt.HS256, sigKey, 10*time.Minute)
    // Enable payload encryption with:
    // signer.WithEncryption(encKey, nil)
    app.Get("/", generateToken(signer))

    verifier := jwt.NewVerifier(jwt.HS256, sigKey)
    // Enable server-side token block feature (even before its expiration time):
    verifier.WithDefaultBlocklist()
    // Enable payload decryption with:
    // verifier.WithDecryption(encKey, nil)
    verifyMiddleware := verifier.Verify(func() interface{} {
        return new(fooClaims)
    })

    protectedAPI := app.Party("/protected")
    // Register the verify middleware to allow access only to authorized clients.
    protectedAPI.Use(verifyMiddleware)
    // ^ or UseRouter(verifyMiddleware) to disallow unauthorized http error handlers too.

    protectedAPI.Get("/", protected)
    // Invalidate the token through server-side, even if it's not expired yet.
    protectedAPI.Get("/logout", logout)

    // http://localhost:8080
    // http://localhost:8080/protected?token=$token (or Authorization: Bearer $token)
    // http://localhost:8080/protected/logout?token=$token
    // http://localhost:8080/protected?token=$token (401)
    app.Listen(":8080")
}

func generateToken(signer *jwt.Signer) iris.Handler {
    return func(ctx iris.Context) {
        claims := fooClaims{Foo: "bar"}

        token, err := signer.Sign(claims)
        if err != nil {
            ctx.StopWithStatus(iris.StatusInternalServerError)
            return
        }

        ctx.Write(token)
    }
}

func protected(ctx iris.Context) {
    // Get the verified and decoded claims.
    claims := jwt.Get(ctx).(*fooClaims)

    // Optionally, get token information if you want to work with them.
    // Just an example on how you can retrieve all the standard claims (set by signer's max age, "exp").
    standardClaims := jwt.GetVerifiedToken(ctx).StandardClaims
    expiresAtString := standardClaims.ExpiresAt().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
    timeLeft := standardClaims.Timeleft()

    ctx.Writef("foo=%s\nexpires at: %s\ntime left: %s\n", claims.Foo, expiresAtString, timeLeft)
}

func logout(ctx iris.Context) {
    err := ctx.Logout()
    if err != nil {
        ctx.WriteString(err.Error())
    } else {
        ctx.Writef("token invalidated, a new token is required to access the protected API")
    }
}
```

> Learn about refresh tokens, blocklist and more at: [_examples/auth/jwt](https://github.com/kataras/iris/tree/main/_examples/auth/jwt).
>
> ​	在以下网址了解有关刷新令牌、阻止列表等的更多信息：_examples/auth/jwt。

## Testing 测试

Iris offers an incredible support for the [httpexpect](https://github.com/gavv/httpexpect), a Testing Framework for web applications. The `iris/httptest` subpackage provides helpers for Iris + httpexpect.

​	Iris 为 httpexpect 提供令人难以置信的支持，httpexpect 是一个用于 Web 应用程序的测试框架。 `iris/httptest` 子包为 Iris + httpexpect 提供帮助程序。

if you prefer the Go's standard [net/http/httptest](https://pkg.go.dev/net/http/httptest) package, you can still use it. Iris as much every http web framework is compatible with any external tool for testing, at the end it's HTTP.

​	如果您更喜欢 Go 的标准 net/http/httptest 包，您仍然可以使用它。Iris 与任何外部测试工具一样兼容任何 http Web 框架，最终它是 HTTP。

### Testing Basic Authentication 测试基本身份验证

In our first example we will use the `iris/httptest` subpackage to test Basic Authentication.

​	在我们的第一个示例中，我们将使用 `iris/httptest` 子包来测试基本身份验证。

**1.** The `main.go` source file looks like that:

​	1. `main.go` 源文件如下：

```go
package main

import (
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/basicauth"
)

func newApp() *iris.Application {
    app := iris.New()

    opts := basicauth.Options{
        Allow: basicauth.AllowUsers(map[string]string{"myusername": "mypassword"}),
    }

    authentication := basicauth.New(opts) // or just: basicauth.Default(map...)

    app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })

    // to party

    needAuth := app.Party("/admin", authentication)
    {
        //http://localhost:8080/admin
        needAuth.Get("/", h)
        // http://localhost:8080/admin/profile
        needAuth.Get("/profile", h)

        // http://localhost:8080/admin/settings
        needAuth.Get("/settings", h)
    }

    return app
}

func h(ctx iris.Context) {
    // username, password, _ := ctx.Request().BasicAuth()
    // third parameter it will be always true because the middleware
    // makes sure for that, otherwise this handler will not be executed.
    // OR:

    user := ctx.User().(*iris.SimpleUser)
    ctx.Writef("%s %s:%s", ctx.Path(), user.Username, user.Password)
    // ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}

func main() {
    app := newApp()
    app.Listen(":8080")
}
```

**2.** Now, create a `main_test.go` file and copy-paste the following.

​	2. 现在，创建一个 `main_test.go` 文件并复制粘贴以下内容。

```go
package main

import (
    "testing"

    "github.com/kataras/iris/v12/httptest"
)

func TestNewApp(t *testing.T) {
    app := newApp()
    e := httptest.New(t, app)

    // redirects to /admin without basic auth
    e.GET("/").Expect().Status(httptest.StatusUnauthorized)
    // without basic auth
    e.GET("/admin").Expect().Status(httptest.StatusUnauthorized)

    // with valid basic auth
    e.GET("/admin").WithBasicAuth("myusername", "mypassword").Expect().
        Status(httptest.StatusOK).Body().Equal("/admin myusername:mypassword")
    e.GET("/admin/profile").WithBasicAuth("myusername", "mypassword").Expect().
        Status(httptest.StatusOK).Body().Equal("/admin/profile myusername:mypassword")
    e.GET("/admin/settings").WithBasicAuth("myusername", "mypassword").Expect().
        Status(httptest.StatusOK).Body().Equal("/admin/settings myusername:mypassword")

    // with invalid basic auth
    e.GET("/admin/settings").WithBasicAuth("invalidusername", "invalidpassword").
        Expect().Status(httptest.StatusUnauthorized)

}
```

**3.** Open your command line and execute:

​	3. 打开命令行并执行：

```bash
$ go test -v
```

### Testing Cookies 测试 Cookie

```go
package main

import (
    "fmt"
    "testing"

    "github.com/kataras/iris/v12/httptest"
)

func TestCookiesBasic(t *testing.T) {
    app := newApp()
    e := httptest.New(t, app, httptest.URL("http://example.com"))

    cookieName, cookieValue := "my_cookie_name", "my_cookie_value"

    // Test Set A Cookie.
    t1 := e.GET(fmt.Sprintf("/cookies/%s/%s", cookieName, cookieValue)).
        Expect().Status(httptest.StatusOK)
    // Validate cookie's existence, it should be available now.
    t1.Cookie(cookieName).Value().Equal(cookieValue)
    t1.Body().Contains(cookieValue)

    path := fmt.Sprintf("/cookies/%s", cookieName)

    // Test Retrieve A Cookie.
    t2 := e.GET(path).Expect().Status(httptest.StatusOK)
    t2.Body().Equal(cookieValue)

    // Test Remove A Cookie.
    t3 := e.DELETE(path).Expect().Status(httptest.StatusOK)
    t3.Body().Contains(cookieName)

    t4 := e.GET(path).Expect().Status(httptest.StatusOK)
    t4.Cookies().Empty()
    t4.Body().Empty()
}
$ go test -v -run=TestCookiesBasic$
```

Iris web framework itself uses this package to test itself. In the [_examples repository directory](https://github.com/kataras/iris/tree/main/_examples) you will find some useful tests as well. For more information please take a look and read the [httpexpect's documentation](https://github.com/gavv/httpexpect).

​	Iris Web 框架本身使用此软件包来测试自身。在 _examples 存储库目录中，您还将找到一些有用的测试。有关更多信息，请查看并阅读 httpexpect 的文档。

## Localization 本地化

### Introduction 简介

Localization features provide a convenient way to retrieve strings in various languages, allowing you to easily support multiple languages within your application. Language strings are stored in files within the `./locales` directory. Within this directory there should be a subdirectory for each language supported by the application:

​	本地化功能提供了一种检索各种语言中字符串的便捷方式，使您能够在应用程序中轻松支持多种语言。语言字符串存储在 `./locales` 目录中的文件中。在此目录中，应该为应用程序支持的每种语言设置一个子目录：

```bash
│   main.go
└───locales
    ├───el-GR
    │       home.yml
    ├───en-US
    │       home.yml
    └───zh-CN
            home.yml
```

The default language for your application is the first registered language.

​	应用程序的默认语言是第一个注册的语言。

```go
app := iris.New()

// First parameter: Glob filpath patern,
// Second variadic parameter: Optional language tags,
// the first one is the default/fallback one.
app.I18n.Load("./locales/*/*", "en-US", "el-GR", "zh-CN")
```

Or if you load all languages by:

​	或者，如果您通过以下方式加载所有语言：

```go
app.I18n.Load("./locales/*/*")
// Then set the default language using:
app.I18n.SetDefault("en-US")
```

### Load embedded locales 加载嵌入式语言环境

You may want to embed locales with the new [embed directive](https://gobyexample.com/embed-directive) within your application executable.

​	您可能希望将区域设置与应用程序可执行文件中的新嵌入指令一起嵌入。

1. Import the embed package; if you don’t use any exported identifiers from this package, you can do a blank import with _ "embed".
   导入嵌入包；如果您不使用此包中的任何导出标识符，则可以使用 _ "embed" 进行空白导入。

```go
import (
    "embed"
)
```

1. Embed directives accept paths relative to the directory containing the Go source file. We can embed multiple files or even folders with wildcards. This uses a variable of the embed.FS type, which implements a simple virtual file system.
   嵌入指令接受相对于包含 Go 源文件的目录的路径。我们可以使用通配符嵌入多个文件甚至文件夹。这使用 embed.FS 类型的变量，它实现了一个简单的虚拟文件系统。

```go
//go:embed embedded/locales/*
var embeddedFS embed.FS
```

1. Instead of the `Load` method, we should use the `LoadFS` one.
   我们应该使用 `LoadFS` 方法，而不是 `Load` 方法。

```go
err := app.I18n.LoadFS(embeddedFS, "./embedded/locales/*/*.ini", "en-US", "el-GR")
// OR to load all languages by filename:
// app.I18n.LoadFS(embeddedFS, "./embedded/locales/*/*.ini")
// Then set the default language using:
// app.I18n.SetDefault("en-US")
```

### Defining Translations 定义翻译

Locale files can be written at YAML(recommended), JSON, TOML or INI form.

​	区域设置文件可以以 YAML（推荐）、JSON、TOML 或 INI 形式编写。

Each file should contain keys. Keys can have sub-keys(we call them "sections") too.

​	每个文件都应包含键。键也可以具有子键（我们称之为“部分”）。

Each key's value should be of form `string` or `map` containing by its translated text (or **template**) or/and its pluralized key-values.

​	每个键的值应为 `string` 或 `map` 形式，其中包含其翻译的文本（或模板）或/和其复数形式的键值。

Iris i18n module supports **pluralization** out-of-the-box, see below.

​	Iris i18n 模块开箱即用地支持复数形式，请参阅下文。 Fmt 样式

### Fmt Style

```yaml
hi: "Hi %s!"
ctx.Tr("Hi", "John")
// Outputs: Hi John!
```

### Template 模板

```yaml
hi: "Hi {{.Name}}!"
ctx.Tr("Hi", iris.Map{"Name": "John"})
// Outputs: Hi John!
```

### Pluralization 复数形式

Iris i18n supports plural variables. To define a per-locale variable you must define a new section of `Vars` key.

​	Iris i18n 支持复数变量。要定义每个区域设置的变量，您必须定义 `Vars` 键的新部分。

The acceptable keys for variables are:

​	变量的可接受键为：

- `one`
- `"=x"` where x is a number
  `"=x"` 其中 x 是一个数字
- `"<x"`
- `other`
- `format`

Example:

​	示例：

```yaml
Vars:
  - Minutes:
      one: "minute"
      other: "minutes"
  - Houses:
      one: "house"
      other: "houses"
```

Then, each message can use this variable, here's how:

​	然后，每条消息都可以使用此变量，方法如下：

```yaml
# Using variables in raw string
YouLate: "You are %[1]d ${Minutes} late."
# [x] is the argument position,
# variables always have priority other fmt-style arguments,
# that's why we see [1] for houses and [2] for the string argument.
HouseCount: "%[2]s has %[1]d ${Houses}."
ctx.Tr("YouLate", 1)
// Outputs: You are 1 minute late.
ctx.Tr("YouLate", 10)
// Outputs: You are 10 minutes late.

ctx.Tr("HouseCount", 2, "John")
// Outputs: John has 2 houses.
```

You can select what message will be shown based on a given plural count.

​	您可以根据给定的复数计数选择要显示的消息。

Except variables, each message can also have its plural form too!

​	除了变量之外，每条消息也可以有其复数形式！

Acceptable keys:

​	可接受的键：

- `zero`
- `one`
- `two`
- `"=x"`
- `"<x"`
- `">x"`
- `other`

Let's create a simple plural-featured message, it can use the Minutes variable we created above too.

​	让我们创建一个简单的复数特征消息，它也可以使用我们上面创建的 Minutes 变量。

```yaml
FreeDay:
  "=3": "You have three days and %[2]d ${Minutes} off." # "FreeDay" 3, 15
  one:  "You have a day off." # "FreeDay", 1
  other: "You have %[1]d free days." # "FreeDay", 5
ctx.Tr("FreeDay", 3, 15)
// Outputs: You have three days and 15 minutes off.
ctx.Tr("FreeDay", 1)
// Outputs: You have a day off.
ctx.Tr("FreeDay", 5)
// Outputs: You have 5 free days.
```

Let's continue with a bit more advanced example, using template text + functions + plural + variables.

​	让我们继续使用模板文本 + 函数 + 复数 + 变量来进行一个更高级的示例。

```yaml
Vars:
  - Houses:
      one: "house"
      other: "houses"
  - Gender:
      "=1": "She"
      "=2": "He"

VarTemplatePlural:
  one: "${Gender} is awesome!"
  other: "other (${Gender}) has %[3]d ${Houses}."
  "=5": "{{call .InlineJoin .Names}} are awesome."
const (
    female = iota + 1
    male
)

ctx.Tr("VarTemplatePlural", iris.Map{
    "PluralCount": 5,
    "Names":       []string{"John", "Peter"},
    "InlineJoin": func(arr []string) string {
        return strings.Join(arr, ", ")
    },
})
// Outputs: John, Peter are awesome

ctx.Tr("VarTemplatePlural", 1, female)
// Outputs: She is awesome!

ctx.Tr("VarTemplatePlural", 2, female, 5)
// Outputs: other (She) has 5 houses.
```

### Sections 部分

If the key is not a reserved one (e.g. one, two...) then it acts as a sub section. The sections are separated by dot characters (`.`).

​	如果键不是保留键（例如，一、二……），则它充当子部分。部分由点字符 ( `.` ) 分隔。

```yaml
Welcome:
  Message: "Welcome {{.Name}}"
ctx.Tr("Welcome.Message", iris.Map{"Name": "John"})
// Outputs: Welcome John
```

### Determining The Current Locale 确定当前语言环境

You may use the `context.GetLocale` method to determine the current locale or check if the locale is a given value:

​	您可以使用 `context.GetLocale` 方法来确定当前语言环境或检查语言环境是否为给定值：

```go
func(ctx iris.Context) {
    locale := ctx.GetLocale()
    // [...]
}
```

The **Locale** interface looks like this.

​	语言环境接口如下所示。

```go
// Locale is the interface which returns from a `Localizer.GetLocale` metod.
// It serves the transltions based on "key" or format. See `GetMessage`.
type Locale interface {
    // Index returns the current locale index from the languages list.
    Index() int
    // Tag returns the full language Tag attached tothis Locale,
    // it should be uniue across different Locales.
    Tag() *language.Tag
    // Language should return the exact languagecode of this `Locale`
    //that the user provided on `New` function.
    //
    // Same as `Tag().String()` but it's static.
    Language() string
    // GetMessage should return translated text based n the given "key".
    GetMessage(key string, args ...interface{}) string
}
```

### Retrieving Translation 检索翻译

Use of `context.Tr` method as a shortcut to get a translated text for this request.

​	使用 `context.Tr` 方法作为快捷方式，可获取此请求的翻译文本。

```go
func(ctx iris.Context) {
    text := ctx.Tr("hi", "name")
    // [...]
}
```

### Inside Views 在视图中

```go
func(ctx iris.Context) {
    err := ctx.View("index.html", iris.Map{
        "tr": ctx.Tr,
    })
    if err!=nil {
        ctx.HTML("<h3>%s</h3>", err.Error())
        return
    }
}
```

### Example

```go
package main

import (
    "github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
    app := iris.New()

    // Configure i18n.
    // First parameter: Glob filpath patern,
    // Second variadic parameter: Optional language tags, the first one is the default/fallback one.
    app.I18n.Load("./locales/*/*.ini", "en-US", "el-GR", "zh-CN")
    // app.I18n.LoadAssets for go-bindata.

    // Default values:
    // app.I18n.URLParameter = "lang"
    // app.I18n.Subdomain = true
    //
    // Set to false to disallow path (local) redirects,
    // see https://github.com/kataras/iris/issues/1369.
    // app.I18n.PathRedirect = true

    app.Get("/", func(ctx iris.Context) {
        hi := ctx.Tr("hi", "iris")

        locale := ctx.GetLocale()

        ctx.Writef("From the language %s translated output: %s", locale.Language(), hi)
    })

    app.Get("/some-path", func(ctx iris.Context) {
        ctx.Writef("%s", ctx.Tr("hi", "iris"))
    })

    app.Get("/other", func(ctx iris.Context) {
        language := ctx.GetLocale().Language()

        fromFirstFileValue := ctx.Tr("key1")
        fromSecondFileValue := ctx.Tr("key2")
        ctx.Writef("From the language: %s, translated output:\n%s=%s\n%s=%s",
            language, "key1", fromFirstFileValue,
            "key2", fromSecondFileValue)
    })

    // using in inside your views:
    view := iris.HTML("./views", ".html")
    app.RegisterView(view)

    app.Get("/templates", func(ctx iris.Context) {
        err := ctx.View("index.html", iris.Map{
            "tr": ctx.Tr, // word, arguments... {call .tr "hi" "iris"}}
        })
        if err != nil {
            ctx.HTML("<h3>%s</h3>", err.Error())
            return
        }

        // Note that,
        // Iris automatically adds a "tr" global template function as well,
        // the only difference is the way you call it inside your templates and
        // that it accepts a language code as its first argument.
    })
    //

    return app
}

func main() {
    app := newApp()

    // go to http://localhost:8080/el-gr/some-path
    // ^ (by path prefix)
    //
    // or http://el.mydomain.com8080/some-path
    // ^ (by subdomain - test locally with the hosts file)
    //
    // or http://localhost:8080/zh-CN/templates
    // ^ (by path prefix with uppercase)
    //
    // or http://localhost:8080/some-path?lang=el-GR
    // ^ (by url parameter)
    //
    // or http://localhost:8080 (default is en-US)
    // or http://localhost:8080/?lang=zh-CN
    //
    // go to http://localhost:8080/other?lang=el-GR
    // or http://localhost:8080/other (default is en-US)
    // or http://localhost:8080/other?lang=en-US
    //
    // or use cookies to set the language.
    app.Listen(":8080", iris.WithSitemap("http://localhost:8080"))
}
```

### Sitemap 站点地图

Sitemap translations are automatically set to each route by path prefix if `app.I18n.PathRedirect` is true or by subdomain if `app.I18n.Subdomain` is true or by URL query parameter if `app.I18n.URLParameter` is not empty.

​	如果 `app.I18n.PathRedirect` 为 true，则网站地图翻译会自动通过路径前缀设置为每个路由；如果 `app.I18n.Subdomain` 为 true，则通过子域设置；如果 `app.I18n.URLParameter` 不为空，则通过 URL 查询参数设置。

Read more at: https://support.google.com/webmasters/answer/189077?hl=en

​	阅读更多内容：https://support.google.com/webmasters/answer/189077?hl=en

```bash
GET http://localhost:8080/sitemap.xml
<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">
    <url>
        <loc>http://localhost:8080/</loc>
        <xhtml:link rel="alternate" hreflang="en-US" href="http://localhost:8080/"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="el-GR" href="http://localhost:8080/el-GR/"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="zh-CN" href="http://localhost:8080/zh-CN/"></xhtml:link>
    </url>
    <url>
        <loc>http://localhost:8080/some-path</loc>
        <xhtml:link rel="alternate" hreflang="en-US" href="http://localhost:8080/some-path"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="el-GR" href="http://localhost:8080/el-GR/some-path"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="zh-CN" href="http://localhost:8080/zh-CN/some-path"></xhtml:link>
    </url>
    <url>
        <loc>http://localhost:8080/other</loc>
        <xhtml:link rel="alternate" hreflang="en-US" href="http://localhost:8080/other"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="el-GR" href="http://localhost:8080/el-GR/other"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="zh-CN" href="http://localhost:8080/zh-CN/other"></xhtml:link>
    </url>
    <url>
        <loc>http://localhost:8080/templates</loc>
        <xhtml:link rel="alternate" hreflang="en-US" href="http://localhost:8080/templates"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="el-GR" href="http://localhost:8080/el-GR/templates"></xhtml:link>
        <xhtml:link rel="alternate" hreflang="zh-CN" href="http://localhost:8080/zh-CN/templates"></xhtml:link>
    </url>
</urlset>
```

That's all the basics about Iris. This document covers enough for beginners. Want to become an expert and a Certificated Iris Developer, learn about MVC, i18n, dependency-injection, gRPC, lambda functions, websockets, best practises and more? [Request the Iris E-Book](https://www.iris-go.com/#ebookDonateForm) today and be participated in the development of Iris!

​	这些就是关于 Iris 的所有基础知识。本文档涵盖了初学者需要了解的所有内容。想成为专家和认证的 Iris 开发人员，了解 MVC、i18n、依赖注入、gRPC、lambda 函数、websocket、最佳实践等内容吗？立即申请 Iris 电子书，参与 Iris 的开发！