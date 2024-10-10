+++
title = "测试"
weight = 140
date = 2023-07-09T21:52:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Testing - 测试

> 原文：[https://echo.labstack.com/docs/testing](https://echo.labstack.com/docs/testing)

## 测试处理程序

`GET` `/users/:id`

Handler below retrieves user by id from the database. If user is not found it returns `404` error with a message.

​	下面的处理程序从数据库中根据用户ID检索用户。如果找不到用户，则返回`404`错误和一条消息。

### CreateUser

`POST` `/users`

- 接受JSON负载
- 成功时返回`201 - Created`
- 错误时返回`500 - Internal Server Error`

### GetUser

`GET` `/users/:email`

- 成功时返回`200 - OK`
- 错误时，如果找不到用户，则返回`404 - Not Found`，否则返回`500 - Internal Server Error`

handler.go

```go
package handler

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

type (
    User struct {
        Name  string `json:"name" form:"name"`
        Email string `json:"email" form:"email"`
    }
    handler struct {
        db map[string]*User
    }
)

func (h *handler) createUser(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, u)
}

func (h *handler) getUser(c echo.Context) error {
    email := c.Param("email")
    user := h.db[email]
    if user == nil {
        return echo.NewHTTPError(http.StatusNotFound, "user not found")
    }
    return c.JSON(http.StatusOK, user)
}
```

handler_test.go

```go
package handler

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

var (
    mockDB = map[string]*User{
        "jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
    }
    userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestCreateUser(t *testing.T) {
    // Setup
    e := echo.New()
    req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    h := &handler{mockDB}

    // Assertions
    if assert.NoError(t, h.createUser(c)) {
        assert.Equal(t, http.StatusCreated, rec.Code)
        assert.Equal(t, userJSON, rec.Body.String())
    }
}

func TestGetUser(t *testing.T) {
    // Setup
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetPath("/users/:email")
    c.SetParamNames("email")
    c.SetParamValues("jon@labstack.com")
    h := &handler{mockDB}

    // Assertions
    if assert.NoError(t, h.getUser(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Equal(t, userJSON, rec.Body.String())
    }
}
```



### 使用 Form Payload

```go
// import "net/url"
f := make(url.Values)
f.Set("name", "Jon Snow")
f.Set("email", "jon@labstack.com")
req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
```



### 设置 Path Params

```go
c.SetParamNames("id", "email")
c.SetParamValues("1", "jon@labstack.com")
```



### 设置 Query Params

```go
// import "net/url"
q := make(url.Values)
q.Set("email", "jon@labstack.com")
req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
```



## 测试 Middleware

*TBD（*待定*）*

​	目前，您可以查看内置（built-in）中间件的[测试用例](https://github.com/labstack/echo/tree/master/middleware)。