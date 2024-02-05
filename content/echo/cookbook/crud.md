+++
title = "crud"
date = 2023-07-09T22:01:42+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# CRUD

> 原文：[https://echo.labstack.com/docs/cookbook/crud](https://echo.labstack.com/docs/cookbook/crud)

## Server

cookbook/crud/server.go

```go
package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
	lock  = sync.Mutex{}
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func getAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
```



## Client

### xxxxxxxxxx43 1package main2​3import (4    "net/http"5    "regexp"6​7    "github.com/labstack/echo/v4"8    "github.com/labstack/echo/v4/middleware"9)10​11var (12    users = []string{"Joe", "Veer", "Zion"}13)14​15func getUsers(c echo.Context) error {16    return c.JSON(http.StatusOK, users)17}18​19// allowOrigin takes the origin as an argument and returns true if the origin20// is allowed or false otherwise.21func allowOrigin(origin string) (bool, error) {22    // In this example we use a regular expression but we can imagine various23    // kind of custom logic. For example, an external datasource could be used24    // to maintain the list of allowed origins.25    return regexp.MatchString(`^https:\/\/labstack\.(net|com)$`, origin)26}27​28func main() {29    e := echo.New()30    e.Use(middleware.Logger())31    e.Use(middleware.Recover())32​33    // CORS restricted with a custom function to allow origins34    // and with the GET, PUT, POST or DELETE methods allowed.35    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{36        AllowOriginFunc: allowOrigin,37        AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},38    }))39​40    e.GET("/api/users", getUsers)41​42    e.Logger.Fatal(e.Start(":1323"))43}go

#### Request

```sh
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe Smith"}' \
  localhost:1323/users
```



#### Response

```js
{
  "id": 1,
  "name": "Joe Smith"
}
```



### Get user

#### Request

```sh
curl localhost:1323/users/1
```



#### Response

```js
{
  "id": 1,
  "name": "Joe Smith"
}
```



### Update user

#### Request

```sh
curl -X PUT \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe"}' \
  localhost:1323/users/1
```



#### Response

```js
{
  "id": 1,
  "name": "Joe"
}
```



### Delete user

#### Request

```sh
curl -X DELETE localhost:1323/users/1
```



#### Response

```
NoContent - 204
```