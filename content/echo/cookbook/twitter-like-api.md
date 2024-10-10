+++
title = "twitter-like-api"
date = 2023-07-09T22:05:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Twitter Like API

> 原文：[https://echo.labstack.com/docs/cookbook/twitter](https://echo.labstack.com/docs/cookbook/twitter)

This recipe demonstrates how to create a Twitter like REST API using MongoDB (Database), JWT (API security) and JSON (Data exchange).

## Models

cookbook/twitter/model/user.go

```go
package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Email     string        `json:"email" bson:"email"`
		Password  string        `json:"password,omitempty" bson:"password"`
		Token     string        `json:"token,omitempty" bson:"-"`
		Followers []string      `json:"followers,omitempty" bson:"followers,omitempty"`
	}
)
```



cookbook/twitter/model/post.go

```go
package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Post struct {
		ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
		To      string        `json:"to" bson:"to"`
		From    string        `json:"from" bson:"from"`
		Message string        `json:"message" bson:"message"`
	}
)
```



## Handlers

cookbook/twitter/handler/handler.go

```go
package handler

import (
	"gopkg.in/mgo.v2"
)

type (
	Handler struct {
		DB *mgo.Session
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
```



cookbook/twitter/handler/user.go

```go
package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echox/cookbook/twitter/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{ID: bson.NewObjectId()}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Follow(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	id := c.Param("id")

	// Add a follower to user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").
		UpdateId(bson.ObjectIdHex(id), bson.M{"$addToSet": bson.M{"followers": userID}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
	}

	return
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
```



cookbook/twitter/handler/post.go

```go
package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echox/cookbook/twitter/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) CreatePost(c echo.Context) (err error) {
	u := &model.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}
	p := &model.Post{
		ID:   bson.NewObjectId(),
		From: u.ID.Hex(),
	}
	if err = c.Bind(p); err != nil {
		return
	}

	// Validation
	if p.To == "" || p.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid to or message fields"}
	}

	// Find user from database
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	// Save post in database
	if err = db.DB("twitter").C("posts").Insert(p); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) FetchPost(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// Retrieve posts from database
	posts := []*model.Post{}
	db := h.DB.Clone()
	if err = db.DB("twitter").C("posts").
		Find(bson.M{"to": userID}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&posts); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, posts)
}
```



## Server

cookbook/twitter/server.go

```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echox/cookbook/twitter/handler"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	// Database connection
	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Create indices
	if err = db.Copy().DB("twitter").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/follow/:id", h.Follow)
	e.POST("/posts", h.CreatePost)
	e.GET("/feed", h.FetchPost)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
```



## API

### Signup

User signup

- Retrieve user credentials from the body and validate against database.
- For invalid email or password, send `400 - Bad Request` response.
- For valid email and password, save user in database and send `201 - Created` response.

#### Request

```sh
curl \
  -X POST \
  http://localhost:1323/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"jon@labstack.com","password":"shhh!"}'
```



#### Response

```
201 - Created
{
  "id": "58465b4ea6fe886d3215c6df",
  "email": "jon@labstack.com",
  "password": "shhh!"
}
```



### Login

User login

- Retrieve user credentials from the body and validate against database.

- For invalid credentials, send `401 - Unauthorized` response.

- For valid credentials, send

   

  ```
  200 - OK
  ```

   

  response:

  - Generate JWT for the user and send it as response.
  - Each subsequent request must include JWT in the `Authorization` header.

```
POST` `/login
```

#### Request

```sh
curl \
  -X POST \
  http://localhost:1323/login \
  -H "Content-Type: application/json" \
  -d '{"email":"jon@labstack.com","password":"shhh!"}'
```



#### Response

```
200 - OK
{
  "id": "58465b4ea6fe886d3215c6df",
  "email": "jon@labstack.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"
}
```



TIP

Client should store the token, for browsers, you may use local storage.

### Follow

Follow a user

- For invalid token, send `400 - Bad Request` response.
- For valid token:
  - If user is not found, send `404 - Not Found` response.
  - Add a follower to the specified user in the path parameter and send `200 - OK` response.

```
POST` `/follow/:id
```

#### Request

```sh
curl \
  -X POST \
  http://localhost:1323/follow/58465b4ea6fe886d3215c6df \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"
```



#### Response

```
200 - OK
```

### Post

Post a message to specified user

- For invalid request payload, send `400 - Bad Request` response.
- If user is not found, send `404 - Not Found` response.
- Otherwise save post in the database and return it via `201 - Created` response.

```
POST` `/posts
```

#### Request

```sh
curl \
  -X POST \
  http://localhost:1323/posts \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0" \
  -H "Content-Type: application/json" \
  -d '{"to":"58465b4ea6fe886d3215c6df","message":"hello"}'
```



#### Response

```
201 - Created
{
  "id": "584661b9a6fe8871a3804cba",
  "to": "58465b4ea6fe886d3215c6df",
  "from": "58465b4ea6fe886d3215c6df",
  "message": "hello"
}
```



### Feed

List most recent messages based on optional `page` and `limit` query parameters

```
GET` `/feed?page=1&limit=5
```

#### Request

```sh
curl \
  -X GET \
  http://localhost:1323/feed \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"
```



#### Response

```
200 - OK
[
  {
    "id": "584661b9a6fe8871a3804cba",
    "to": "58465b4ea6fe886d3215c6df",
    "from": "58465b4ea6fe886d3215c6df",
    "message": "hello"
  }
]
```