+++
title = "测试"
date = 2024-01-31T19:08:51+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-chi.io/#/pages/testing](https://go-chi.io/#/pages/testing)

# Testing 🧪 测试

Writing tests for APIs is easy. We can use the inbuilt `net/http/httptest` lib to test our apis.

​	为 API 编写测试很容易。我们可以使用内置的 `net/http/httptest` 库来测试我们的 API。

### Usage 用法

First we will create a simple Hello World Api

​	首先，我们将创建一个简单的 Hello World API

```go
package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    s := CreateNewServer()
    s.MountHandlers()
    http.ListenAndServe(":3000", s.Router)
}

// HelloWorld api Handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}

type Server struct {
    Router *chi.Mux
    // Db, config can be added here
}

func CreateNewServer() *Server {
    s := &Server{}
    s.Router = chi.NewRouter()
    return s
}

func (s *Server) MountHandlers() {
    // Mount all Middleware here
    s.Router.Use(middleware.Logger)

    // Mount all handlers here
    s.Router.Get("/", HelloWorld)

}
```

This is how a standard api would look, with a `Server` struct where we can add our router, and database connection...etc.

​	标准 API 的外观如下，其中包含一个 `Server` 结构，我们可以在其中添加路由器、数据库连接等。

We then write a `CreateNewServer` function to return a New Server with a `chi.Mux` Router

​	然后，我们编写一个 `CreateNewServer` 函数来返回一个带有 `chi.Mux` 路由器的新服务器

We can then Mount all Handlers and middlewares in a single server method `MountHandlers`

​	然后，我们可以在单个服务器方法 `MountHandlers` 中挂载所有处理程序和中间件

We can now start writing tests for this.

​	我们现在可以开始为此编写测试。

When writing tests, we will assert what values our api will return

​	在编写测试时，我们将断言我们的 API 将返回哪些值

So for the route `/` our api should return `Hello World!` and a status code of `200`

​	因此，对于路由 `/` ，我们的 API 应返回 `Hello World!` 和状态代码 `200`

Now in another file `main_test.go`

​	现在在另一个文件中 `main_test.go`

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "os"
    "testing"

    "github.com/stretchr/testify/require"
)

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    s.Router.ServeHTTP(rr, req)

    return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func TestHelloWorld(t *testing.T) {
    // Create a New Server Struct
    s := CreateNewServer()
    // Mount Handlers
    s.MountHandlers()

    // Create a New Request
    req, _ := http.NewRequest("GET", "/", nil)

    // Execute Request
    response := executeRequest(req, s)

    // Check the response code
    checkResponseCode(t, http.StatusOK, response.Code)

    // We can use testify/require to assert values, as it is more convenient
    require.Equal(t, "Hello World!", response.Body.String())
}
```

Now run `go test ./... -v -cover`

​	现在运行 `go test ./... -v -cover`

Voila, your tests work now.

​	瞧，你的测试现在可以运行了。