+++
title = "教程：用go和Gin开发一个RESTful API"
weight = 10
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Tutorial: Developing a RESTful API with Go and Gin - 教程：用Go和Gin开发一个RESTful API

> 原文：[https://go.dev/doc/tutorial/web-service-gin](https://go.dev/doc/tutorial/web-service-gin)

​	本教程介绍了用Go和[Gin Web 框架](https://gin-gonic.com/docs/)（Gin）编写RESTful Web服务API的基础知识。

​	如果您对Go及其工具有基本的了解，您会从本教程中得到最大的收获。如果这是您第一次接触Go，请参阅[Tutorial: Get started with Go](../TutorialGetStartedWithGo)的快速介绍。

​	`Gin`简化了许多与构建 Web 应用相关的编码任务，包括 web 服务。在本教程中，您将使用Gin来路由请求，检索请求细节，并将JSON用于响应。

​	在本教程中，您将建立一个有两个端点的RESTful API服务器。您的示例项目将是一个关于复古爵士乐唱片（vintage jazz records）的数据存储库。

本教程包括以下几个部分：

1. 设计API端点。
2. 为您的代码创建一个文件夹。
3. 创建数据。
4. 编写返回所有项的处理程序。
5. 编写添加新项的处理程序。
6. 编写返回特定项的处理程序。

注意：关于其他教程，请看[Tutorials](../Tutorials)。

要尝试在Google Cloud Shell中完成这个互动教程，请点击下面的按钮。

[![Open in Cloud Shell](TutorialDevelopingARESTfulAPIWithGoAndGin_img/open-btn.png)](https://ide.cloud.google.com/?cloudshell_workspace=~&walkthrough_tutorial_url=https://raw.githubusercontent.com/golang/tour/master/tutorial/web-service-gin.md)

## 前提条件 

- 安装 `Go 1.16` 或更高版本。关于安装说明，请参见[Installing Go](../InstallingGo)。
- 编辑代码的工具。您拥有的任何文本编辑器都可以工作。
- 命令终端。在 Linux 和 Mac 上使用任何终端，以及在 Windows 上使用 PowerShell 或 cmd，Go 都能很好地工作。
- `curl` 工具。在Linux和Mac上，这个工具应该已经安装。在Windows上，它包含在`Windows 10 Insider build 17063`及以后的版本中。对于早期的Windows版本，您可能需要安装它。更多信息，请参见 [Tar and Curl Come to Windows](https://docs.microsoft.com/en-us/virtualization/community/team-blog/2017/20171219-tar-and-curl-come-to-windows)。

## 设计API端点 

​	您将建立一个API，提供对一个销售黑胶唱片的商店的访问。所以您需要提供端点，客户端可以通过这些端点为用户获取和添加专辑。

​	在开发一个API时，您通常从设计端点开始。如果端点易于理解，您的API的用户将获得更多的成功。

以下是您在本教程中要创建的端点。

/albums

- `GET` – 获取所有专辑的列表，以JSON格式返回。
- `POST` – 从以JSON格式发送的请求数据中添加一个新专辑。

/albums/:id

- `GET` –通过其ID获得一个专辑，以JSON格式返回专辑数据。

接下来，您将为代码创建一个文件夹。

## 为您的代码创建一个文件夹

首先，为您要写的代码创建一个项目。

a. 打开一个命令提示符，切换到您的主目录。

在Linux或Mac上：

```shell
$ cd
```

在Windows上：

```shell
C:\> cd %HOMEPATH%
```

b. 使用命令提示符，为您的代码创建一个名为`web-service-gin`的目录。

```shell
$ mkdir web-service-gin
$ cd web-service-gin
```

c. 创建一个模块，您可以在其中管理依赖项。

​	运行`go mod init`命令，给它一个您的代码将在其中的模块的路径。

```shell
$ go mod init example/web-service-gin
go: creating new go.mod: module example/web-service-gin
```

​	该命令创建了一个`go.mod`文件，您添加的依赖项将被列在其中以便追踪。关于用模块路径命名模块的更多信息，请参见[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies#naming-a-module)。

接下来，您将设计处理数据的数据结构。

## 创建数据 

​	为了简化本教程，将数据存储在内存中。更典型的 API 将与数据库交互。

​	请注意，在内存中存储数据意味着每次您停止服务器时，这组专辑就会丢失，然后在您启动它时重新创建。

#### 编写代码

a. 使用您的文本编辑器，在`web-service`目录下创建一个名为`main.go`的文件。您将在这个文件中写下您的Go代码。

b. 在`main.go`文件的顶部，粘贴以下包声明。

```
package main
```

独立程序（相对于一个库）总是在`main`包中。

c. 在包声明的下面，粘贴以下`album`结构的声明。您将用它来存储内存中的专辑数据。

​	`结构标签`（如`json: "artist"`）指定了结构内容序列化为JSON时字段的名称。如果没有这些标签，JSON将使用结构中大写的字段名 —— 这种风格在JSON中并不常见。

```go linenums="1"
// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
```

d. 在您刚刚添加的结构声明下面，粘贴以下`album`结构切片，其中包含将用于启动的数据。

```go linenums="1"
// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```

接下来，您要写代码来实现您的第一个端点。

## 编写返回所有项的处理程序

​	当客户端使用`GET /albums`发出请求时，您想以JSON格式返回所有的专辑。

要做到这一点，您要写以下内容：

- 准备一个响应的逻辑
- 将请求路径映射到逻辑代码

请注意，这与它们在运行时的执行方式相反，但您首先要添加依赖关系，然后是依赖它们的代码。

#### 编写代码

a. 在上一节添加的结构代码下面，粘贴以下代码以获得专辑列表。

​	这个 `getAlbums` 函数从`album`结构切片中创建 JSON，将 JSON 写入响应。

```go linenums="1"
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
```

在这段代码中，您：

- 编写一个`getAlbums`函数，它需要一个[gin.Context](https://pkg.go.dev/github.com/gin-gonic/gin#Context) 参数。注意，您可以给这个函数起任何名字 ——Gin和Go都不要求特定的函数名称格式。

  ​	[gin.Context](https://pkg.go.dev/github.com/gin-gonic/gin#Context)是`Gin`中最重要的部分。它携带请求的细节，验证和序列化JSON，等等。(尽管名字相似，但这与Go的内置[context](https://go.dev/pkg/context/)包不同）。

-  调用 [Context.IndentedJSON](https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON)来将结构体序列化为JSON，并将其添加到响应中。

  ​	该函数的第一个参数是您想发送给客户端的HTTP状态代码。在这里，您要传递 `net/http` 包中的 [StatusOK](https://pkg.go.dev/net/http#StatusOK)常量，以指示 `200 OK`。

  Note that you can replace `Context.IndentedJSON` with a call to [Context.JSON](https://pkg.go.dev/github.com/gin-gonic/gin#Context.JSON) to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
  
  注意，您可以用调用 [Context.JSON](https://pkg.go.dev/github.com/gin-gonic/gin#Context.JSON) 来代替 [Context.IndentedJSON](https://pkg.go.dev/github.com/gin-gonic/gin#Context.IndentedJSON) 来发送更紧凑的 JSON。在实践中，缩进的形式在调试时更容易操作，而且大小差异通常很小。

b. 在`main.go`的顶部附近，就在`albums`切片声明下面，粘贴下面的代码，将处理函数分配给一个端点路径。

​	这就建立了一个关联，让`getAlbums`处理对`/albums`端点路径的请求。

```go linenums="1"
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
}
```

在这段代码中，您：

- 使用[Default](https://pkg.go.dev/github.com/gin-gonic/gin#Default)初始化一个Gin路由器。

- 使用[GET](https://pkg.go.dev/github.com/gin-gonic/gin#RouterGroup.GET)函数将`GET` HTTP方法和`/albums`路径与一个处理函数联系起来。

  注意，您传递的是`getAlbums`函数的名字。这与传递函数的**结果**不同，您可以通过传递`getAlbums()`来做到这一点（注意括号）。

- 使用[Run](https://pkg.go.dev/github.com/gin-gonic/gin#Engine.Run)函数将路由器连接到`http.Server`上并启动服务器。

c. 在`main.go`的顶部，就在包声明的下面，导入您需要的包来支持您刚刚写的代码。

第一行代码应该是这样的：

```go linenums="1"
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)
```

d. 保存`main.go`。

#### 运行代码

a. 开始跟踪Gin模块作为一个依赖项。

​	在命令行中，使用[go get](https://go.dev/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them)来添加`github.com/gin-gonic/gin`模块作为您的模块的依赖项。使用`点参数`表示 "获取当前目录下代码的依赖项"。

```shell
$ go get .
go get: added github.com/gin-gonic/gin v1.7.2
```

Go解决并下载这个依赖，以满足您在上一步添加的`import`声明。

b. 从包含`main.go`的目录中的命令行，运行代码。使用`点参数`表示 "在当前目录下运行代码"。

```shell
$ go run .
```

​	一旦代码运行，您就有了一个运行中的HTTP服务器，您可以向其发送请求。

c. 在一个新的命令行窗口中，使用`curl`向您正在运行的网络服务发出一个请求。

```shell
$ curl http://localhost:8080/albums
```

该命令应该显示您为该服务提供的数据。

```json linenums="1"
[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        }
]
```

​	您已经启动了一个API! 在下一节中，您将使用代码创建另一个端点来处理添加项的`POST`请求。

## 编写添加新项的处理程序 

​	当客户端向`/albums`发出`POST`请求时，您想把请求正文中描述的专辑添加到现有的专辑数据中。

要做到这一点，您要写以下内容：

- 将新专辑添加到现有列表中的逻辑。
- 一段将`POST`请求路由到逻辑的代码。

#### 编写代码

a. 添加代码，将专辑数据添加到专辑列表中。

​	在`import` 语句后的某个地方，粘贴以下代码。(文件的末尾是放置这段代码的好地方，但Go并不强制执行您声明函数的顺序。)

```go linenums="1"
// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
```

在这段代码中，您：

- 使用[Context.BindJSON](https://pkg.go.dev/github.com/gin-gonic/gin#Context.BindJSON)将请求主体绑定到`newAlbum`。

- 将从JSON中初始化的`album`结构追加到`album`切片中。

- 在响应中添加一个`201`状态代码，以及表示您添加的专辑的JSON。

  

b. 更改`main`函数，使其包括 `router.POST` 函数，如下所示。

```go linenums="1"
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
```

在这段代码中，您：

- 将`/albums`路径上的`POST`方法与`postAlbums`函数联系起来。

  ​	通过Gin，您可以将处理程序与HTTP 的 method-and-path 组合联系起来。通过这种方式，您可以根据客户端使用的方法，分别路由 sent 到单个路径的请求。

#### 运行代码

a. 如果服务器在上一节中仍在运行，请停止它。

b. 从包含`main.go`的目录中的命令行，运行代码。

```shell
$ go run .
```

c. 在另一个命令行窗口中，使用`curl`向正在运行的网络服务发出请求。

```shell
$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

该命令应该为添加的专辑显示 响应头和 JSON。

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Wed, 02 Jun 2021 00:34:12 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
```

d. 和上一节一样，使用`curl`检索完整的专辑列表，您可以用它来确认新专辑是否被添加。

```shell
$ curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

该命令应该显示专辑列表：

```json linenums="1"
[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        },
        {
                "id": "4",
                "title": "The Modern Sound of Betty Carter",
                "artist": "Betty Carter",
                "price": 49.99
        }
]
```

在下一节，您将添加代码来处理一个特定项的`GET`。

## 编写返回特定项的处理程序

​	当客户端发出`GET /albums/[id]`的请求时，您想返回ID与`id`路径参数相符的专辑。

​	要做到这一点，您需要：

- 添加逻辑来检索请求的专辑。
- 将路径映射到逻辑中。

#### 编写代码

a. 在上一节中添加的`postAlbums`函数下面，粘贴以下代码以检索特定的专辑。

​	这个`getAlbumByID`函数将提取请求路径中的ID，然后定位一个匹配的专辑。

```go linenums="1"
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```

在这段代码中，您：

- 使用[Context.Param](https://pkg.go.dev/github.com/gin-gonic/gin#Context.Param)从URL中检索出`id`路径参数。当您把这个处理程序映射到一个路径时，您会在路径中加入一个参数的占位符。

- 循环遍历切片中的`album`结构，寻找一个`ID`字段值与`id`参数值相匹配的结构。如果找到了，您就将该`album`结构序列化为JSON，并以`200 OK`的HTTP状态码作为响应返回。

  如上所述，现实世界的服务可能会使用数据库查询来执行这一查找。

- 如果没有找到专辑，就用[http.StatusNotFound](https://pkg.go.dev/net/http#StatusNotFound)返回一个HTTP `404` 错误。

b. 最后，更改您的`main`函数，使其包括对`router.GET`的新调用，现在的路径是`/albums/:id`，如以下例子所示。

```go linenums="1"
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
```

在这段代码中，您：

- 将`/albums/:id`路径与`getAlbumByID`函数联系起来。在Gin中，路径中项之前的冒号标志着该项是一个`路径参数 （path parameter.）`。

#### 运行代码

a. 如果服务器在上一节中仍在运行，请停止它。

b. 从包含`main.go`的目录中的命令行，运行代码来启动服务器。

```shell
$ go run .
```

c. 在另一个命令行窗口，使用`curl`向您正在运行的网络服务发出请求。

```shell
$ curl http://localhost:8080/albums/2
```

​	该命令应该为您使用的ID的专辑显示JSON。如果没有找到该专辑，您会得到带有错误信息的JSON。

```json linenums="1"
{
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
}
```

## 总结 

​	恭喜您！您刚刚用Go和Gin编写了一个简单的RESTful网络服务。您刚刚使用Go和Gin编写了一个简单的RESTful Web服务。

建议的下一个主题：

- 如果您是Go的新手，您会在[Effective Go](https://go.dev/doc/effective_go)和[How to write Go code](https://go.dev/doc/code)中找到有用的最佳实践。
- [Go Tour](https://go.dev/tour/)是对Go基础知识的一个很好的循序渐进的介绍。
- 关于Gin的更多信息，请参见[Gin Web Framework package documentation](https://pkg.go.dev/github.com/gin-gonic/gin)或[Gin Web Framework docs](https://gin-gonic.com/docs/)。

## 完整的代码 

本节包含您通过本教程构建的应用程序的代码。

```go title="main.go" linenums="1"
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```