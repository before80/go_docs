+++
title = "gin快速入门"
date = 2023-06-05T08:55:39+08:00
type = "docs"
weight = 2
description = ""
isCJKLanguage = true
draft = false

+++

# Gin Quick Start - gin快速入门

> 原文：[https://github.com/gin-gonic/gin/blob/master/docs/doc.md](https://github.com/gin-gonic/gin/blob/master/docs/doc.md)

## 构建标签

### 使用替代的json包构建

​	gin默认使用`encoding/json`作为json包，但你可以通过使用其他标签进行构建来更改它。

[jsoniter](https://github.com/json-iterator/go)

```sh
go build -tags=jsoniter .
```

[go-json](https://github.com/goccy/go-json)

```sh
go build -tags=go_json .
```

[sonic](https://github.com/bytedance/sonic) (你必须确保你的CPU支持AVX指令。)

```sh
$ go build -tags="sonic avx" .
```

### 不使用`MsgPack`渲染特性构建

​	gin默认启用`MsgPack`渲染特性。但你可以通过指定`nomsgpack`构建标签来禁用此特性：

```sh
go build -tags=nomsgpack .
```

​	这有助于减小可执行文件的大小。详细信息请参见[详细信息](https://github.com/gin-gonic/gin/pull/1852)。

## API 示例

​	你可以在[Gin示例仓库](https://github.com/gin-gonic/examples)中找到许多可直接运行的示例。

### 使用GET, POST, PUT, PATCH, DELETE and OPTIONS

```go
func main() {
  // 创建一个带有默认中间件的gin router：
  // logger和recovery（无崩溃 crash-free）中间件
  router := gin.Default()

  router.GET("/someGet", getting)
  router.POST("/somePost", posting)
  router.PUT("/somePut", putting)
  router.DELETE("/someDelete", deleting)
  router.PATCH("/somePatch", patching)
  router.HEAD("/someHead", head)
  router.OPTIONS("/someOptions", options)

  // 默认情况下它监听 :8080，除非定义了一个 PORT 环境变量。
  router.Run()
 // 如果要使用硬编码端口，则使用 router.Run(":3000")
}
```

### 路径参数

```go
func main() {
  router := gin.Default()

  // 这个handler将匹配 /user/john，但不会匹配 /user/ 或 /user
  router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  // 然而，这个handler将匹配 /user/john/ 和 /user/john/send
  // 如果没有其他路由器匹配 /user/john，它将重定向到 /user/john/  
  router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })

  // 对于每个匹配的请求，Context将保存路由定义
  router.POST("/user/:name/*action", func(c *gin.Context) {
    b := c.FullPath() == "/user/:name/*action" // true
    c.String(http.StatusOK, "%t", b)
  })

  // 这个handler将为 /user/groups 添加一个新的路由器。
  // 精确匹配的路由器在param routes之前解析，无论它们的定义顺序如何。
  // 以 /user/groups 开头的路由永远不会被解释为 /user/:name/... 路由
  router.GET("/user/groups", func(c *gin.Context) {
    c.String(http.StatusOK, "The available groups are [...]")
  })

  router.Run(":8080")
}
```

### 查询字符串参数

```go
func main() {
  router := gin.Default()

  // 查询字符串参数使用现有的底层请求对象进行解析。
  // 请求响应的 URL 匹配：/welcome?firstname=Jane&lastname=Doe  
  router.GET("/welcome", func(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") //  c.Request.URL.Query().Get("lastname") 的简写

    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })
  router.Run(":8080")
}
```

### Multipart/Urlencoded 表单

```go
func main() {
  router := gin.Default()

  router.POST("/form_post", func(c *gin.Context) {
    message := c.PostForm("message")
    nick := c.DefaultPostForm("nick", "anonymous")

    c.JSON(http.StatusOK, gin.H{
      "status":  "posted",
      "message": message,
      "nick":    nick,
    })
  })
  router.Run(":8080")
}
```

### 另一个示例: query + post 表单

```sh
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
```

```go
func main() {
  router := gin.Default()

  router.POST("/post", func(c *gin.Context) {

    id := c.Query("id")
    page := c.DefaultQuery("page", "0")
    name := c.PostForm("name")
    message := c.PostForm("message")

    fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
  })
  router.Run(":8080")
}
```

```sh
id: 1234; page: 1; name: manu; message: this_is_great
```

### 将 Map 作为查询字符串或 post 表单参数

```sh
POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
Content-Type: application/x-www-form-urlencoded

names[first]=thinkerou&names[second]=tianou
```

```go
func main() {
  router := gin.Default()

  router.POST("/post", func(c *gin.Context) {

    ids := c.QueryMap("ids")
    names := c.PostFormMap("names")

    fmt.Printf("ids: %v; names: %v", ids, names)
  })
  router.Run(":8080")
}
```

```sh
ids: map[b:hello a:1234]; names: map[second:tianou first:thinkerou]
```

### 上传文件

#### 单个文件

​	参考问题[#774](https://github.com/gin-gonic/gin/issues/774)和详细的[示例代码](https://github.com/gin-gonic/examples/tree/master/upload-file/single)。

​	`file.Filename` **SHOULD NOT 不应该**被信任。参见[`Content-Disposition`在MDN上的说明](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition#Directives)和[#1693](https://github.com/gin-gonic/gin/issues/1693)。

> ​	文件名始终是可选的，应用程序不应该盲目使用它：应该去除路径信息，并进行服务器文件系统规则的转换。

```go
func main() {
  router := gin.Default()
  // 设置 multipart 表单的内存限制（默认是 32 MiB）
  router.MaxMultipartMemory = 8 << 20  // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
    // 单个文件
    file, _ := c.FormFile("file")
    log.Println(file.Filename)

    // 将文件上传到指定的目标位置
    c.SaveUploadedFile(file, dst)

    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
  })
  router.Run(":8080")
}
```

​	使用 `curl` ：

```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/appleboy/test.zip" \
  -H "Content-Type: multipart/form-data"
```

#### 多个文件

​	参见详细的[示例代码](https://github.com/gin-gonic/examples/tree/master/upload-file/multiple)。

```go
func main() {
  router := gin.Default()
  // 设置 multipart 表单的内存限制（默认是 32 MiB）
  router.MaxMultipartMemory = 8 << 20  // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
    // Multipart form
    form, _ := c.MultipartForm()
    files := form.File["upload[]"]

    for _, file := range files {
      log.Println(file.Filename)
      
      // 将文件上传到指定的目标位置
      c.SaveUploadedFile(file, dst)
    }
    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
  })
  router.Run(":8080")
}
```

​	使用 `curl` ：

```bash
curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/appleboy/test1.zip" \
  -F "upload[]=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"
```

### 路由分组

```go
func main() {
  router := gin.Default()

  // 简单分组：v1 
  v1 := router.Group("/v1")
  {
    v1.POST("/login", loginEndpoint)
    v1.POST("/submit", submitEndpoint)
    v1.POST("/read", readEndpoint)
  }

  // 简单分组：v2
  v2 := router.Group("/v2")
  {
    v2.POST("/login", loginEndpoint)
    v2.POST("/submit", submitEndpoint)
    v2.POST("/read", readEndpoint)
  }

  router.Run(":8080")
}
```

### 默认情况下没有中间件的空白gin

使用

```go
r := gin.New()
```

而不是

```go
// 默认情况下，附带Logger和Recovery中间件
r := gin.Default()
```

### 使用中间件

```go
func main() { 
  // 默认情况下创建一个没有任何中间件的路由器
  r := gin.New()

  // 全局中间件
  // Logger中间件会将日志写入gin.DefaultWriter，即使你设置了 GIN_MODE=release。
  // 默认情况下，gin.DefaultWriter = os.Stdout 
  r.Use(gin.Logger())

  // Recovery中间件会在出现恐慌时恢复，并在有恐慌时写入一个500响应。
  r.Use(gin.Recovery())

  // 为每个路由添加中间件，你可以添加任意多个。
  r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

  // 授权组 - Authorization group
  // authorized := r.Group("/", AuthRequired())
  // 与上面代码完全等价:
  authorized := r.Group("/")
  // 在授权组中使用自定义创建的 AuthRequired() 中间件。
  // 每个分组的中间件！在这种情况下，我们在"authorized"分组中使用自定义创建的AuthRequired()中间件。
  authorized.Use(AuthRequired())
  {
    authorized.POST("/login", loginEndpoint)
    authorized.POST("/submit", submitEndpoint)
    authorized.POST("/read", readEndpoint)

    // 嵌套组
    testing := authorized.Group("testing")

    // 访问 0.0.0.0:8080/testing/analytics
    testing.GET("/analytics", analyticsEndpoint)
  }

  // 监听并在 0.0.0.0:8080 上提供服务
  r.Run(":8080")
}
```

### 自定义Recovery 行为

```go
func main() {
  // 默认情况下创建一个没有任何中间件的路由器
  r := gin.New()

  // 全局中间件
  // Logger中间件会将日志写入gin.DefaultWriter，即使你设置了 GIN_MODE=release。
  // 默认情况下，gin.DefaultWriter = os.Stdout
  r.Use(gin.Logger())

  // Recovery中间件会在出现恐慌时恢复，并在有恐慌时写入一个500响应。 
  r.Use(gin.CustomRecovery(func(c *gin.Context, recovered any) {
    if err, ok := recovered.(string); ok {
      c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
    }
    c.AbortWithStatus(http.StatusInternalServerError)
  }))

  r.GET("/panic", func(c *gin.Context) {
   // panic with a string —— 自定义的中间件可以将其保存到数据库或向用户报告。 
    panic("foo")
  })

  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "ohai")
  })

  // 监听并在 0.0.0.0:8080 上提供服务
  r.Run(":8080")
}
```

### 如何写日志文件

```go
func main() {
  // 禁用控制台颜色，当将日志写入文件时，不需要控制台颜色。
  gin.DisableConsoleColor()

  // 将日志写入文件。
  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f)

  // 如果需要将日志同时写入文件和控制台，请使用以下代码。
  // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
    
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

   router.Run(":8080")
}
```

### 自定义日志格式

```go
func main() {
  router := gin.New()

  // LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter。
  // 默认情况下，gin.DefaultWriter = os.Stdout
  router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

    // 你自定义的格式
    return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
        param.ClientIP,
        param.TimeStamp.Format(time.RFC1123),
        param.Method,
        param.Path,
        param.Request.Proto,
        param.StatusCode,
        param.Latency,
        param.Request.UserAgent(),
        param.ErrorMessage,
    )
  }))
  router.Use(gin.Recovery())

  router.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

示例输出

```sh
::1 - [Fri, 07 Dec 2018 17:04:38 JST] "GET /ping HTTP/1.1 200 122.767µs "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36" "
```

### 控制日志输出颜色

​	默认情况下，根据检测到的TTY对控制台上的日志进行颜色化。

​	永远不要给日志着色：

```go
func main() {
  // 禁用日志颜色
  gin.DisableConsoleColor()

  // 创建一个具有默认中间件的 gin 路由器：日志记录器（logger）和recovery（crash-free）中间件。
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

​	始终给日志着色：

```go
func main() {
  // 强制给日志着色 
  gin.ForceConsoleColor()

  // // 创建一个具有默认中间件的 gin 路由器：日志记录器（logger）和recovery（crash-free）中间件。
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
      c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

### Model绑定和验证

​	使用模型绑定将请求体绑定到类型上。我们目前支持JSON、XML、YAML、TOML和标准表单值（`foo=bar&boo=baz`）的绑定。

​	gin使用[go-playground/validator/v10](https://github.com/go-playground/validator)进行验证。可以在此处查看有关标签使用的完整文档 [here](https://pkg.go.dev/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags)

> 请注意，您需要在要绑定的所有字段上设置相应的绑定标签。例如，当从JSON绑定时，设置`json:"fieldname"`。
>

​	此外，Gin提供了两组用于绑定的方法：

- **Type** —— 必须绑定
  - **Methods** —— `Bind`, `BindJSON`, `BindXML`, `BindQuery`, `BindYAML`, `BindHeader`, `BindTOML`
  - **Behavior** —— 这些方法在内部使用`MustBindWith`。使用`c.AbortWithError(400, err).SetType(ErrorTypeBind)`，如果有绑定错误，请求会被中止。这将把响应状态码设置为400，并将`Content-Type`头设置为`text/plain; charset=utf-8`。请注意，如果在此之后尝试设置响应代码，将会出现警告`[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422`。如果您希望更好地控制行为，请考虑使用相应的`ShouldBind`方法。
- **Type** —— 应该绑定
  - **Methods** —— `ShouldBind`, `ShouldBindJSON`, `ShouldBindXML`, `ShouldBindQuery`, `ShouldBindYAML`, `ShouldBindHeader`, `ShouldBindTOML`,
  - **Behavior** —— 这些方法在内部使用`ShouldBindWith`。如果有绑定错误，将返回错误，开发人员有责任适当处理请求和错误。

​	当使用Bind-method方法时，gin会根据Content-Type头来推断binder。如果您确定要绑定的内容，可以使用`MustBindWith`或`ShouldBindWith`。

​	您还可以指定特定字段是必需的。如果字段被装饰为`binding:"required"`并且在绑定时具有空值，将返回错误。

```go
// 从JSON绑定
type Login struct {
  User     string `form:"user" json:"user" xml:"user"  binding:"required"`
  Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
  router := gin.Default()
 
  // 以JSON为例进行绑定({"user": "manu", "password": "123"}) 
  router.POST("/loginJSON", func(c *gin.Context) {
    var json Login
    if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if json.User != "manu" || json.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // 以XML为例进行绑定 (
  //  <?xml version="1.0" encoding="UTF-8"?>
  //  <root>
  //    <user>manu</user>
  //    <password>123</password>
  //  </root>)
  router.POST("/loginXML", func(c *gin.Context) {
    var xml Login
    if err := c.ShouldBindXML(&xml); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if xml.User != "manu" || xml.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // 以HTML表单为例进行绑定 (user=manu&password=123)
  router.POST("/loginForm", func(c *gin.Context) {
    var form Login
    // 这将根据Content-Type头来推断要使用的binder。
    if err := c.ShouldBind(&form); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if form.User != "manu" || form.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // 监听并在0.0.0.0:8080上提供服务
  router.Run(":8080")
}
```

示例请求：

```sh
$ curl -v -X POST \
  http://localhost:8080/loginJSON \
  -H 'content-type: application/json' \
  -d '{ "user": "manu" }'
> POST /loginJSON HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.51.0
> Accept: */*
> content-type: application/json
> Content-Length: 18
>
* upload completely sent off: 18 out of 18 bytes
< HTTP/1.1 400 Bad Request
< Content-Type: application/json; charset=utf-8
< Date: Fri, 04 Aug 2017 03:51:31 GMT
< Content-Length: 100
<
{"error":"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
```

跳过验证：使用上面的`curl`命令运行上面的示例时，它会返回错误。因为示例中为`Password`使用了`binding:"required"`。如果为`Password`使用`binding:"-"`，再次运行上面的代码示例时，将不会返回错误。

### 自定义验证器

​	您还可以注册自定义验证器。请参阅[示例代码](https://github.com/gin-gonic/examples/tree/master/custom-validation/server.go)。

```go
package main

import (
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
  "github.com/go-playground/validator/v10"
)

// Booking 包含绑定和验证的数据。
type Booking struct {
  CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
  CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
  date, ok := fl.Field().Interface().(time.Time)
  if ok {
    today := time.Now()
    if today.After(date) {
      return false
    }
  }
  return true
}

func main() {
  route := gin.Default()

  if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("bookabledate", bookableDate)
  }

  route.GET("/bookable", getBookable)
  route.Run(":8085")
}

func getBookable(c *gin.Context) {
  var b Booking
  if err := c.ShouldBindWith(&b, binding.Query); err == nil {
    c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
  } else {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
}
```

```console
$ curl "localhost:8085/bookable?check_in=2030-04-16&check_out=2030-04-17"
{"message":"Booking dates are valid!"}

$ curl "localhost:8085/bookable?check_in=2030-03-10&check_out=2030-03-09"
{"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}

$ curl "localhost:8085/bookable?check_in=2000-03-09&check_out=2000-03-10"
{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag"}%
```

​	[结构级别验证](https://github.com/go-playground/validator/releases/tag/v8.7)也可以通过此方式注册。

​	请参阅[struct-lvl-validation example](https://github.com/gin-gonic/examples/tree/master/struct-lvl-validations)了解更多信息。

### 仅绑定查询字符串

​	`ShouldBindQuery`函数只绑定查询参数而不绑定post数据。请参阅[详细信息](https://github.com/gin-gonic/gin/issues/742#issuecomment-315953017)。

```go
package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

type Person struct {
  Name    string `form:"name"`
  Address string `form:"address"`
}

func main() {
  route := gin.Default()
  route.Any("/testing", startPage)
  route.Run(":8085")
}

func startPage(c *gin.Context) {
  var person Person
  if c.ShouldBindQuery(&person) == nil {
    log.Println("====== Only Bind By Query String ======")
    log.Println(person.Name)
    log.Println(person.Address)
  }
  c.String(http.StatusOK, "Success")
}

```

### 绑定查询字符串或Post数据

​	请参阅[详细信息](https://github.com/gin-gonic/gin/issues/742#issuecomment-264681292)。

```go
package main

import (
  "log"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
)

type Person struct {
  Name       string    `form:"name"`
  Address    string    `form:"address"`
  Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
  CreateTime time.Time `form:"createTime" time_format:"unixNano"`
  UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
  route := gin.Default()
  route.GET("/testing", startPage)
  route.Run(":8085")
}

func startPage(c *gin.Context) {
  var person Person
  // 如果是`GET`请求，仅使用`Form`绑定引擎（`query`）。
  // 如果是`POST`请求，首先检查`content-type`是否为`JSON`或`XML`，然后使用`Form`（`form-data`）。
  // 请参阅 https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88 了解更多信息。
  if c.ShouldBind(&person) == nil {
    log.Println(person.Name)
    log.Println(person.Address)
    log.Println(person.Birthday)
    log.Println(person.CreateTime)
    log.Println(person.UnixTime)
  }

  c.String(http.StatusOK, "Success")
}
```

使用以下命令进行测试：

```sh
curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
```

### 绑定 Uri

​	请参阅[详细信息](https://github.com/gin-gonic/gin/issues/846)。

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type Person struct {
  ID string `uri:"id" binding:"required,uuid"`
  Name string `uri:"name" binding:"required"`
}

func main() {
  route := gin.Default()
  route.GET("/:name/:id", func(c *gin.Context) {
    var person Person
    if err := c.ShouldBindUri(&person); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
      return
    }
    c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
  })
  route.Run(":8088")
}
```

​	使用以下命令进行测试：

```sh
curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
curl -v localhost:8088/thinkerou/not-uuid
```

### 绑定 Header

```go
package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

type testHeader struct {
  Rate   int    `header:"Rate"`
  Domain string `header:"Domain"`
}

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    h := testHeader{}

    if err := c.ShouldBindHeader(&h); err != nil {
      c.JSON(http.StatusOK, err)
    }

    fmt.Printf("%#v\n", h)
    c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
  })

  r.Run()

// client
// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/
// output
// {"Domain":"music","Rate":300}
}
```

### 绑定HTML复选框

​	请参阅[详细信息](https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092)

main.go

```go
...

type myForm struct {
    Colors []string `form:"colors[]"`
}

...

func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}

...

```

form.html

```html
<form action="/" method="POST">
    <p>Check some colors</p>
    <label for="red">Red</label>
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>
```

result:

```json
{"color":["red","green","blue"]}
```

### Multipart/Urlencoded binding

```go
type ProfileForm struct {
  Name   string                `form:"name" binding:"required"`
  Avatar *multipart.FileHeader `form:"avatar" binding:"required"`

  // 或者多个文件
  // Avatars []*multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
  router := gin.Default()
  router.POST("/profile", func(c *gin.Context) {
    // 可以使用明确的绑定声明绑定multipart form：
    // c.ShouldBindWith(&form, binding.Form)
    // 或者可以简单地使用 ShouldBind 方法进行自动绑定：
    var form ProfileForm
 
    // 在这种情况下，将自动选择适当的绑定方式
    if err := c.ShouldBind(&form); err != nil {
      c.String(http.StatusBadRequest, "bad request")
      return
    }

    err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
    if err != nil {
      c.String(http.StatusInternalServerError, "unknown error")
      return
    }

    // db.Save(&form)

    c.String(http.StatusOK, "ok")
  })
  router.Run(":8080")
}
```

​	使用以下命令进行测试：

```sh
curl -X POST -v --form name=user --form "avatar=@./avatar.png" http://localhost:8080/profile
```

### XML, JSON, YAML, TOML and ProtoBuf rendering

```go
func main() {
  r := gin.Default()

  // gin.H 是 map[string]interface{} 的快捷方式
  r.GET("/someJSON", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/moreJSON", func(c *gin.Context) {  
    // 也可以使用结构体
    var msg struct {
      Name    string `json:"user"`
      Message string
      Number  int
    }
    msg.Name = "Lena"
    msg.Message = "hey"
    msg.Number = 123  
    // 注意 msg.Name 在 JSON 中变为 "user"
	// 输出：{"user": "Lena", "Message": "hey", "Number": 123}		
    c.JSON(http.StatusOK, msg)
  })

  r.GET("/someXML", func(c *gin.Context) {
    c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/someYAML", func(c *gin.Context) {
    c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/someTOML", func(c *gin.Context) {
    c.TOML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/someProtoBuf", func(c *gin.Context) {
    reps := []int64{int64(1), int64(2)}
    label := "test"
    // protobuf 的具体定义在 testdata/protoexample 文件中
    data := &protoexample.Test{
      Label: &label,
      Reps:  reps,
    }

    // 注意 data 在响应中成为二进制数据
	// 输出 protoexample.Test 的序列化 protobuf 数据
    c.ProtoBuf(http.StatusOK, data)
  })

  // 在 0.0.0.0:8080 上监听和服务
  r.Run(":8080")
}
```

#### SecureJSON

​	使用 SecureJSON 防止 JSON 劫持。默认情况下，如果给定的结构体是数组值，则在响应体前添加 `"while(1),"`。

```go
func main() {
  r := gin.Default()

  // 也可以使用自定义的安全 JSON 前缀
  // r.SecureJsonPrefix(")]}',\n")
  r.GET("/someJSON", func(c *gin.Context) {
    names := []string{"lena", "austin", "foo"}

    // 输出：while(1);["lena","austin","foo"] 
    c.SecureJSON(http.StatusOK, names)
  })

  // 在 0.0.0.0:8080 上监听和服务
  r.Run(":8080")
}
```

#### JSONP

​	使用 JSONP 从不同域请求数据。如果查询参数中存在回调函数，则将回调函数添加到响应体中。

```go
func main() {
  r := gin.Default()

  r.GET("/JSONP", func(c *gin.Context) {
    data := gin.H{
      "foo": "bar",
    }

    // 回调函数为 x
	// 输出：x({\"foo\":\"bar\"})
    c.JSONP(http.StatusOK, data)
  })

  // 在 0.0.0.0:8080 上监听和服务 
  r.Run(":8080")

    	// 客户端
		// curl http://127.0.0.1:8080/JSONP?callback=x
}
```

#### AsciiJSON

​	使用 AsciiJSON 生成只包含 ASCII 字符的 JSON，将非 ASCII 字符转义。

```go
func main() {
  r := gin.Default()

  r.GET("/someJSON", func(c *gin.Context) {
    data := gin.H{
      "lang": "GO语言",
      "tag":  "<br>",
    }

    // 输出： {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
    c.AsciiJSON(http.StatusOK, data)
  })

  // 在 0.0.0.0:8080 上监听和服务
  r.Run(":8080")
}
```

#### PureJSON

​	通常，JSON 会将特殊的 HTML 字符替换为 Unicode 实体，例如 `<` 会变为 `\u003c`。如果要直接编码这些字符，可以使用 PureJSON。 

​	该功能在 Go 1.6 及更低版本不可用。

```go
func main() {
  r := gin.Default()

  // 服务端以 Unicode 实体形式返回 
  r.GET("/json", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

  // 服务端以字面字符形式返回
  r.GET("/purejson", func(c *gin.Context) {
    c.PureJSON(http.StatusOK, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

  // 监听并在 0.0.0.0:8080 上服务 
  r.Run(":8080")
}
```

### 提供静态文件 Serving static files

```go
func main() {
  router := gin.Default()
  router.Static("/assets", "./assets")
  router.StaticFS("/more_static", http.Dir("my_file_system"))
  router.StaticFile("/favicon.ico", "./resources/favicon.ico")
  router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))  
 
  // 监听并在 0.0.0.0:8080 上服务 
  router.Run(":8080")
}
```

### 从文件中提供数据 Serving data from file

```go
func main() {
  router := gin.Default()

  router.GET("/local/file", func(c *gin.Context) {
    c.File("local/file.go")
  })

  var fs http.FileSystem = // ...
  router.GET("/fs/file", func(c *gin.Context) {
    c.FileFromFS("fs/file.go", fs)
  })
}

```

### 从reader中提供数据 Serving data from reader

```go
func main() {
  router := gin.Default()
  router.GET("/someDataFromReader", func(c *gin.Context) {
    response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
    if err != nil || response.StatusCode != http.StatusOK {
      c.Status(http.StatusServiceUnavailable)
      return
    }

    reader := response.Body
     defer reader.Close()
    contentLength := response.ContentLength
    contentType := response.Header.Get("Content-Type")

    extraHeaders := map[string]string{
      "Content-Disposition": `attachment; filename="gopher.png"`,
    }

    c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
  })
  router.Run(":8080")
}
```

### HTML 渲染

​	使用 LoadHTMLGlob() 或 LoadHTMLFiles()

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")
  //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
  router.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Main website",
    })
  })
  router.Run(":8080")
}
```

templates/index.tmpl

```html
<html>
  <h1>
    {{ .title }}
  </h1>
</html>
```

​	在不同目录中使用相同名称的模板

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/**/*")
  router.GET("/posts/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
      "title": "Posts",
    })
  })
  router.GET("/users/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
      "title": "Users",
    })
  })
  router.Run(":8080")
}
```

{{< tabpane text=true >}}

{{< tab header="templates/posts/index.tmpl" >}}


```html
{{ define "posts/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>Using posts/index.tmpl</p>
</html>
{{ end }}
```

{{< /tab >}}

{{< tab header="templates/users/index.tmpl" >}}

```html
{{ define "users/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>Using users/index.tmpl</p>
</html>
{{ end }}
```

{{< /tab >}}

{{< /tabpane >}}



#### 自定义模板渲染器

​	还可以使用自定义的 HTML 模板渲染器

```go
import "html/template"

func main() {
  router := gin.Default()
  html := template.Must(template.ParseFiles("file1", "file2"))
  router.SetHTMLTemplate(html)
  router.Run(":8080")
}
```

#### 自定义分隔符

​	可以使用自定义分隔符

```go
  r := gin.Default()
  r.Delims("{[{", "}]}")
  r.LoadHTMLGlob("/path/to/templates")
```

#### 自定义模板函数

​	请参考详细的 [示例代码](https://github.com/gin-gonic/examples/tree/master/template)。

main.go

```go
import (
  "fmt"
  "html/template"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
  year, month, day := t.Date()
  return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
  router := gin.Default()
  router.Delims("{[{", "}]}")
  router.SetFuncMap(template.FuncMap{
      "formatAsDate": formatAsDate,
  })
  router.LoadHTMLFiles("./testdata/template/raw.tmpl")

  router.GET("/raw", func(c *gin.Context) {
      c.HTML(http.StatusOK, "raw.tmpl", gin.H{
          "now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
      })
  })

  router.Run(":8080")
}

```

raw.tmpl

```html
Date: {[{.now | formatAsDate}]}
```

Result:

```sh
Date: 2017/07/01
```

### Multitemplate

​	Gin 默认只允许使用一个 html.Template。可以使用 [multitemplate render](https://github.com/gin-contrib/multitemplate) 来使用类似于 Go 1.6 中的 `block template` 功能。

### 重定向

​	发出（Issuing ） HTTP 重定向很容易。支持内部和外部locations 。

```go
r.GET("/test", func(c *gin.Context) {
  c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
})
```

​	从 POST 请求发出（Issuing ） HTTP 重定向。参考问题：[#444](https://github.com/gin-gonic/gin/issues/444)

```go
r.POST("/test", func(c *gin.Context) {
  c.Redirect(http.StatusFound, "/foo")
})
```

​	发出（Issuing ）路由器重定向，使用 `HandleContext` 如下：

``` go
r.GET("/test", func(c *gin.Context) {
    c.Request.URL.Path = "/test2"
    r.HandleContext(c)
})
r.GET("/test2", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": "world"})
})
```

### 自定义中间件

```go
func Logger() gin.HandlerFunc {
  return func(c *gin.Context) {
    t := time.Now()

    // 设置示例变量
    c.Set("example", "12345")

	// 请求之前
    c.Next()

    // 请求之后
    latency := time.Since(t)
    log.Print(latency)

    // 访问我们发送的状态
    status := c.Writer.Status()
    log.Println(status)
  }
}

func main() {
  r := gin.New()
  r.Use(Logger())

  r.GET("/test", func(c *gin.Context) {
    example := c.MustGet("example").(string)

    // 它会打印："12345"
    log.Println(example)
  })

  // 监听并在 0.0.0.0:8080 上服务
  r.Run(":8080")
}
```

### 使用 BasicAuth() 中间件

```go
// simulate some private data
// 模拟一些私密数据
var secrets = gin.H{
  "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
  "austin": gin.H{"email": "austin@example.com", "phone": "666"},
  "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
  r := gin.Default()

  // 使用 gin.BasicAuth() 中间件创建分组
  // gin.Accounts 是 map[string]string 的快捷方式
  authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
    "foo":    "bar",
    "austin": "1234",
    "lena":   "hello2",
    "manu":   "4321",
  }))

  // /admin/secrets 端点
  // hit "localhost:8080/admin/secrets
  authorized.GET("/secrets", func(c *gin.Context) {
    // 获取 user，它由 BasicAuth 中间件设置
    user := c.MustGet(gin.AuthUserKey).(string)
    if secret, ok := secrets[user]; ok {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
    } else {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
    }
  })

  // 监听并在 0.0.0.0:8080 上服务
  r.Run(":8080")
}
```

### 中间件内部的 goroutines 

​	在中间件或处理程序内部启动新的 goroutines 时，**SHOULD NOT （不应）**在其中使用原始context，而应使用只读副本。

```go
func main() {
  r := gin.Default()

  r.GET("/long_async", func(c *gin.Context) {
    // 创建一个副本，在 goroutine 内使用副本
    cCp := c.Copy()
    go func() {
      // 模拟一个耗时的任务，使用 time.Sleep() 休眠 5 秒
      time.Sleep(5 * time.Second)

      // 注意你在这里使用的是副本 "cCp" 的上下文，非常重要
      log.Println("Done! in path " + cCp.Request.URL.Path)
    }()
  })

  r.GET("/long_sync", func(c *gin.Context) {
    // 模拟一个耗时的任务，使用 time.Sleep() 休眠 5 秒
    time.Sleep(5 * time.Second)

    // 由于我们没有使用 goroutine，所以不需要复制context
    log.Println("Done! in path " + c.Request.URL.Path)
  })

  // 监听并在 0.0.0.0:8080 上服务
  r.Run(":8080")
}
```

### 自定义 HTTP 配置

​	可以直接使用 `http.ListenAndServe()` 进行自定义的 HTTP 配置，如下所示：

```go
func main() {
  router := gin.Default()
  http.ListenAndServe(":8080", router)
}
```

或者

```go
func main() {
  router := gin.Default()

  s := &http.Server{
    Addr:           ":8080",
    Handler:        router,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  s.ListenAndServe()
}
```

### 支持 Let's Encrypt

example for 1-line LetsEncrypt HTTPS servers.

​	使用自动化的一行代码实现 LetsEncrypt HTTPS 服务器。

```go
package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/autotls"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // Ping 处理程序
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
```

​	使用自定义的 autocert 管理器的示例：

```go
package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/autotls"
  "github.com/gin-gonic/gin"
  "golang.org/x/crypto/acme/autocert"
)

func main() {
  r := gin.Default()

  // Ping 处理程序
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  m := autocert.Manager{
    Prompt:     autocert.AcceptTOS,
    HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
    Cache:      autocert.DirCache("/var/www/.cache"),
  }

  log.Fatal(autotls.RunWithManager(r, &m))
}
```

### 使用 gin 运行多个服务

​	参考问题[#346](https://github.com/gin-gonic/gin/issues/346) 并尝试以下示例：

```go
package main

import (
  "log"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
  "golang.org/x/sync/errgroup"
)

var (
  g errgroup.Group
)

func router01() http.Handler {
  e := gin.New()
  e.Use(gin.Recovery())
  e.GET("/", func(c *gin.Context) {
    c.JSON(
      http.StatusOK,
      gin.H{
        "code":  http.StatusOK,
        "error": "Welcome server 01",
      },
    )
  })

  return e
}

func router02() http.Handler {
  e := gin.New()
  e.Use(gin.Recovery())
  e.GET("/", func(c *gin.Context) {
    c.JSON(
      http.StatusOK,
      gin.H{
        "code":  http.StatusOK,
        "error": "Welcome server 02",
      },
    )
  })

  return e
}

func main() {
  server01 := &http.Server{
    Addr:         ":8080",
    Handler:      router01(),
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
  }

  server02 := &http.Server{
    Addr:         ":8081",
    Handler:      router02(),
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
  }

  g.Go(func() error {
    err := server01.ListenAndServe()
    if err != nil && err != http.ErrServerClosed {
      log.Fatal(err)
    }
    return err
  })

  g.Go(func() error {
    err := server02.ListenAndServe()
    if err != nil && err != http.ErrServerClosed {
      log.Fatal(err)
    }
    return err
  })

  if err := g.Wait(); err != nil {
    log.Fatal(err)
  }
}
```

### 优雅地关闭或重启

​	有几种方法可以实现优雅地关闭或重启。您可以使用专门用于此目的的第三方软件包，也可以使用内置软件包的函数和方法手动执行相同的操作。

#### 第三方包

​	我们可以使用 [fvbock/endless](https://github.com/fvbock/endless) 来替代默认的 `ListenAndServe`。有关更多详细信息，请参考问题 [#296](https://github.com/gin-gonic/gin/issues/296)。

```go
router := gin.Default()
router.GET("/", handler)
// [...]
endless.ListenAndServe(":4242", router)
```

其他选择：

* [grace](https://github.com/facebookgo/grace)：Go 服务器的优雅重启和零停机部署。
* [graceful](https://github.com/tylerb/graceful)：graceful 是一个 Go 包，用于优雅地关闭 http.Handler 服务器。
* [manners](https://github.com/braintree/manners)：Manners 是一个友好的 Go HTTP 服务器，能够优雅地关闭。

#### 手动操作

​	如果您使用的是 Go 1.8 或更高版本，可能不需要使用那些库。考虑使用 `http.Server` 内置的 [Shutdown()](https://pkg.go.dev/net/http#Server.Shutdown) 方法进行优雅的关闭。下面的示例描述了它的用法，并且我们在[这里](https://github.com/gin-gonic/examples/tree/master/graceful-shutdown)有更多使用 gin 的示例 。

```go
// +build go1.8

package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
    time.Sleep(5 * time.Second)
    c.String(http.StatusOK, "Welcome Gin Server")
  })

  srv := &http.Server{
    Addr:    ":8080",
    Handler: router,
  }

  // 在 goroutine 中初始化服务器，以便不会阻塞下面的优雅关闭处理
  go func() {
    if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
      log.Printf("listen: %s\n", err)
    }
  }()

  // 等待中断信号以便优雅地关闭服务器，超时时间为 5 秒。
  quit := make(chan os.Signal)
  // kill (无参数) 默认发送 syscall.SIGTERM
  // kill -2 则发送 syscall.SIGINT
  // kill -9 则发送 syscall.SIGKILL，但无法被捕获，因此不需要添加
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  log.Println("Shutting down server...")

  // context用于通知服务器有 5 秒时间来完成当前正在处理的请求
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  if err := srv.Shutdown(ctx); err != nil {
    log.Fatal("Server forced to shutdown:", err)
  }

  log.Println("Server exiting")
}
```

### 使用模板构建单个二进制文件

​	您可以使用 [embed](https://pkg.go.dev/embed) 包将服务器构建为包含模板的单个二进制文件。

```go
package main

import (
  "embed"
  "html/template"
  "net/http"

  "github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS

func main() {
  router := gin.Default()
  templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
  router.SetHTMLTemplate(templ)

  // 示例： /public/assets/images/example.png
  router.StaticFS("/public", http.FS(f))

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Main website",
    })
  })

  router.GET("/foo", func(c *gin.Context) {
    c.HTML(http.StatusOK, "bar.tmpl", gin.H{
      "title": "Foo website",
    })
  })

  router.GET("favicon.ico", func(c *gin.Context) {
    file, _ := f.ReadFile("assets/favicon.ico")
    c.Data(
      http.StatusOK,
      "image/x-icon",
      file,
    )
  })

  router.Run(":8080")
}
```

​	在[https://github.com/gin-gonic/examples/tree/master/assets-in-binary/example02](https://github.com/gin-gonic/examples/tree/master/assets-in-binary/example02)目录中可以找到完整示例。

### 将 form-data 请求绑定到自定义结构体

​	使用自定义结构体的示例：

```go
type StructA struct {
    FieldA string `form:"field_a"`
}

type StructB struct {
    NestedStruct StructA
    FieldB string `form:"field_b"`
}

type StructC struct {
    NestedStructPointer *StructA
    FieldC string `form:"field_c"`
}

type StructD struct {
    NestedAnonyStruct struct {
        FieldX string `form:"field_x"`
    }
    FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
    var b StructB
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "a": b.NestedStruct,
        "b": b.FieldB,
    })
}

func GetDataC(c *gin.Context) {
    var b StructC
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "a": b.NestedStructPointer,
        "c": b.FieldC,
    })
}

func GetDataD(c *gin.Context) {
    var b StructD
    c.Bind(&b)
    c.JSON(http.StatusOK, gin.H{
        "x": b.NestedAnonyStruct,
        "d": b.FieldD,
    })
}

func main() {
    r := gin.Default()
    r.GET("/getb", GetDataB)
    r.GET("/getc", GetDataC)
    r.GET("/getd", GetDataD)

    r.Run()
}
```

​	使用 `curl` 命令进行测试：

```sh
$ curl "http://localhost:8080/getb?field_a=hello&field_b=world"
{"a":{"FieldA":"hello"},"b":"world"}
$ curl "http://localhost:8080/getc?field_a=hello&field_c=world"
{"a":{"FieldA":"hello"},"c":"world"}
$ curl "http://localhost:8080/getd?field_x=hello&field_d=world"
{"d":"world","x":{"FieldX":"hello"}}
```

### 尝试将请求体绑定到不同的结构体中

​	常规的请求体绑定方法会消耗 `c.Request.Body`，而且不能多次调用。

```go
type formA struct {
  Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
  Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
  objA := formA{}
  objB := formB{}
  // 这个 c.ShouldBind 会消耗 c.Request.Body，并且不能被重复使用。
  if errA := c.ShouldBind(&objA); errA == nil {
    c.String(http.StatusOK, `the body should be formA`)
  // Always an error is occurred by this because c.Request.Body is EOF now.
  // 因为 c.Request.Body 现在已经是 EOF，所以总是会发生错误。
  } else if errB := c.ShouldBind(&objB); errB == nil {
    c.String(http.StatusOK, `the body should be formB`)
  } else {
    ...
  }
}
```

​	为此，您可以使用 `c.ShouldBindBodyWith` 方法：

```go
func SomeHandler(c *gin.Context) {
  objA := formA{}
  objB := formB{}
  // This reads c.Request.Body and stores the result into the context.
  // 这个方法会读取 c.Request.Body 并将结果存储到context中。
  if errA := c.ShouldBindBodyWith(&objA, binding.Form); errA == nil {
    c.String(http.StatusOK, `the body should be formA`)
  // At this time, it reuses body stored in the context.
  // 这时，它会重用存储在context中的请求体。
  } else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
    c.String(http.StatusOK, `the body should be formB JSON`)
  // 它还可以接受其他格式
  } else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
    c.String(http.StatusOK, `the body should be formB XML`)
  } else {
    ...
  }
}
```

1. `c.ShouldBindBodyWith` 在绑定之前将请求体存储到context中。这对性能有一些影响，所以如果您只需要调用一次绑定，就不应该使用这个方法。
2. 这个功能只适用于某些格式，如 `JSON`、`XML`、`MsgPack`、`ProtoBuf`。对于其他格式，如 `Query`、`Form`、`FormPost`、`FormMultipart`，可以多次调用 `c.ShouldBind()` 而不会对性能造成任何损害（参见 [#1341](https://github.com/gin-gonic/gin/pull/1341)）。

### 将 form-data 请求绑定到自定义结构体和自定义标签

```go
const (
  customerTag = "url"
  defaultMemory = 32 << 20
)

type customerBinding struct {}

func (customerBinding) Name() string {
  return "form"
}

func (customerBinding) Bind(req *http.Request, obj any) error {
  if err := req.ParseForm(); err != nil {
    return err
  }
  if err := req.ParseMultipartForm(defaultMemory); err != nil {
    if err != http.ErrNotMultipart {
      return err
    }
  }
  if err := binding.MapFormWithTag(obj, req.Form, customerTag); err != nil {
    return err
  }
  return validate(obj)
}

func validate(obj any) error {
  if binding.Validator == nil {
    return nil
  }
  return binding.Validator.ValidateStruct(obj)
}

// 现在我们可以这样做！！！
// FormA 是一个外部类型，我们无法修改它的标签
type FormA struct {
  FieldA string `url:"field_a"`
}

func ListHandler(s *Service) func(ctx *gin.Context) {
  return func(ctx *gin.Context) {
    var urlBinding = customerBinding{}
    var opt FormA
    err := ctx.MustBindWith(&opt, urlBinding)
    if err != nil {
      ...
    }
    ...
  }
}
```

### http2 server push

​	http.Pusher 只支持 **go1.8+**。请参阅 [golang 博客]({{< ref "/goBlog/2017/HTTP2ServerPush" >}}) 获取详细信息。

```go
package main

import (
  "html/template"
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
  r := gin.Default()
  r.Static("/assets", "./assets")
  r.SetHTMLTemplate(html)

  r.GET("/", func(c *gin.Context) {
    if pusher := c.Writer.Pusher(); pusher != nil {
      // 使用 pusher.Push() 进行服务器推送
      if err := pusher.Push("/assets/app.js", nil); err != nil {
        log.Printf("Failed to push: %v", err)
      }
    }
    c.HTML(http.StatusOK, "https", gin.H{
      "status": "success",
    })
  })

  // 在 https://127.0.0.1:8080 上监听和提供服务 
  r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
```

### 定义路由日志的格式

​	默认的路由日志格式如下：

```sh
[GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
[GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
[GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
```

​	如果你希望以指定格式（如 JSON、键值对或其他格式）记录这些信息，你可以使用 `gin.DebugPrintRouteFunc` 来定义日志格式。 

​	在下面的示例中，我们使用标准日志包记录所有的路由，但你也可以使用适合你需求的其他日志工具。

```go
import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
    log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
  }

  r.POST("/foo", func(c *gin.Context) {
    c.JSON(http.StatusOK, "foo")
  })

  r.GET("/bar", func(c *gin.Context) {
    c.JSON(http.StatusOK, "bar")
  })

  r.GET("/status", func(c *gin.Context) {
    c.JSON(http.StatusOK, "ok")
  })

  // 在 http://0.0.0.0:8080 上监听和提供服务
  r.Run()
}
```

### 设置和获取 cookie

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/cookie", func(c *gin.Context) {

      cookie, err := c.Cookie("gin_cookie")

      if err != nil {
          cookie = "NotSet"
          c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
      }

      fmt.Printf("Cookie value: %s \n", cookie)
  })

  router.Run()
}
```

## 不信任所有代理

​	gin 允许您指定要保存真实客户端IP的头部（如果有的话），并指定您信任的代理服务器（或直接客户端）来指定这些头部之一。

​	使用`gin.Engine`上的`SetTrustedProxies()`函数，可以指定客户端IP相关请求头部可信任的网络地址或网络CIDR。它们可以是IPv4地址、IPv4 CIDR、IPv6地址或IPv6 CIDR。

**注意：** 如果您没有使用上述函数指定受信任的代理服务器，默认情况下gin信任所有代理服务器，**这是不安全的**。同时，如果您不使用任何代理服务器，可以通过使用`Engine.SetTrustedProxies(nil)`来禁用此功能，然后`Context.ClientIP()`将直接返回远程地址，避免一些不必要的计算。

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.SetTrustedProxies([]string{"192.168.1.2"})

  router.GET("/", func(c *gin.Context) {
    // 如果客户端是192.168.1.2，
    // 则使用X-Forwarded-For头部来推断原始客户端IP，从该头部的可信任部分获取。
    // 否则，直接返回客户端的真实IP。
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })
  router.Run()
}
```

**注意：** 如果你正在使用 CDN 服务，你可以将 `Engine.TrustedPlatform` 设置为跳过 TrustedProxies 检查，它的优先级比 TrustedProxies 更高。 

​	看下面的示例：

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  // Use predefined header gin.PlatformXXX
  // 使用预定义的头部 gin.PlatformXXX
  router.TrustedPlatform = gin.PlatformGoogleAppEngine
  // 或者设置自己的受信任代理服务的请求头部
  // 不要将其设置为任何可疑的请求头部，这是不安全的
  router.TrustedPlatform = "X-CDN-IP"

  router.GET("/", func(c *gin.Context) {
    // 如果你设置了 TrustedPlatform，ClientIP() 将解析相应的头部并直接返回 IP
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })
  router.Run()
}
```

## 测试

​	`net/http/httptest` 包是进行 HTTP 测试的首选方法。

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })
  return r
}

func main() {
  r := setupRouter()
  r.Run(":8080")
}
```

​	上述代码示例的测试：

```go
package main

import (
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
  router := setupRouter()

  w := httptest.NewRecorder()
  req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
  router.ServeHTTP(w, req)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Equal(t, "pong", w.Body.String())
}
```