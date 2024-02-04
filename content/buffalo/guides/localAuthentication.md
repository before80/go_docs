+++
title = "本地身份验证"
date = 2024-02-04T21:17:27+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/auth/](https://gobuffalo.io/documentation/guides/auth/)

# Local Authentication 本地身份验证 

In many use-cases, you’ll need to implement user authentication in your apps.

​	在许多用例中，您需要在应用中实现用户身份验证。

Buffalo had a native support for Auth until version `v0.9.4`. Since then, it was moved into it’s own plugin, https://github.com/gobuffalo/buffalo-auth.

​	Buffalo 在版本 `v0.9.4` 之前本机支持身份验证。从那时起，它被移到了自己的插件中，https://github.com/gobuffalo/buffalo-auth。

## Installation 安装 

To install the `buffalo-auth` plugin, run the following command at your project route:

​	要安装 `buffalo-auth` 插件，请在项目路由中运行以下命令：

```bash
$ buffalo plugins install github.com/gobuffalo/buffalo-auth
```

## Generator 生成器 

```bash
$ buffalo g auth

create  models/user.go
create  models/user_test.go
    run  goimports -w actions/actions_test.go actions/app.go actions/home.go actions/home_test.go actions/render.go grifts/db.go grifts/init.go main.go models/models.go models/models_test.go models/user.go models/user_test.go
create  migrations/20180910062057_create_users.up.fizz
create  migrations/20180910062057_create_users.down.fizz
create  actions/auth.go
create  actions/auth_test.go
create  actions/users.go
create  actions/users_test.go
create  models/user_test.go
create  actions/home_test.go
create  templates/auth/new.html
create  templates/index.html
create  templates/users/new.html
```

## Example Usage 示例用法 

### Actions 操作 

actions/app.go

actions/auth.go

actions/auth_test.go

actions/home.go

actions/home_test.go

actions/users.go

actions/users_test.go

```go
// actions/app.go
package actions

import (
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/buffalo/middleware"
  "github.com/gobuffalo/buffalo/middleware/ssl"
  "github.com/gobuffalo/envy"
  "github.com/unrolled/secure"

  "coke/models"

  "github.com/gobuffalo/buffalo/middleware/csrf"
  "github.com/gobuffalo/buffalo/middleware/i18n"
  "github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
  if app == nil {
    app = buffalo.New(buffalo.Options{
      Env:         ENV,
      SessionName: "_coke_session",
    })
    // Automatically redirect to SSL
    app.Use(forceSSL())

    if ENV == "development" {
      app.Use(middleware.ParameterLogger)
    }

    // Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
    // Remove to disable this.
    app.Use(csrf.New)

    // Wraps each request in a transaction.
    //  c.Value("tx").(*pop.PopTransaction)
    // Remove to disable this.
    app.Use(middleware.PopTransaction(models.DB))

    // Setup and use translations:
    app.Use(translations())

    app.GET("/", HomeHandler)

    app.Use(SetCurrentUser)
    app.Use(Authorize)
    app.GET("/users/new", UsersNew)
    app.POST("/users", UsersCreate)
    app.GET("/signin", AuthNew)
    app.POST("/signin", AuthCreate)
    app.DELETE("/signout", AuthDestroy)
    app.Middleware.Skip(Authorize, HomeHandler, UsersNew, UsersCreate, AuthNew, AuthCreate)
    app.ServeFiles("/", assetsBox) // serve files from the public directory
  }

  return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
  var err error
  if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
    app.Stop(err)
  }
  return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
  return ssl.ForceSSL(secure.Options{
    SSLRedirect:     ENV == "production",
    SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
  })
}
```

### Models 模型 

models/user.go

models/user_test.go

```go
// models/user.go
package models

import (
  "encoding/json"
  "time"

  "github.com/gobuffalo/pop"
  "github.com/gobuffalo/uuid"
  "github.com/gobuffalo/validate"
  "github.com/gobuffalo/validate/validators"
"strings"
"github.com/pkg/errors"
"golang.org/x/crypto/bcrypt"
)

type User struct {
  ID           uuid.UUID `json:"id" db:"id"`
  CreatedAt    time.Time `json:"created_at" db:"created_at"`
  UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
  Email        string    `json:"email" db:"email"`
  PasswordHash string    `json:"password_hash" db:"password_hash"`
Password string `json:"-" db:"-"`
PasswordConfirmation string `json:"-" db:"-"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
  ju, _ := json.Marshal(u)
  return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
  ju, _ := json.Marshal(u)
  return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
    var err error
    return validate.Validate(
        &validators.StringIsPresent{Field: u.Email, Name: "Email"},
        &validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
        // check to see if the email address is already taken:
        &validators.FuncValidator{
            Field:   u.Email,
            Name:    "Email",
            Message: "%s is already taken",
            Fn: func() bool {
                var b bool
                q := tx.Where("email = ?", u.Email)
                if u.ID != uuid.Nil {
                    q = q.Where("id != ?", u.ID)
                }
                b, err = q.Exists(u)
                if err != nil {
                    return false
                }
                return !b
            },
        },
    ), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
    var err error
    return validate.Validate(
        &validators.StringIsPresent{Field: u.Password, Name: "Password"},
        &validators.StringsMatch{Name: "Password", Field: u.Password, Field2: u.PasswordConfirmation, Message: "Password does not match confirmation"},
    ), err
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
  return validate.NewErrors(), nil
}


// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
    u.Email = strings.ToLower(strings.TrimSpace(u.Email))
    ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return validate.NewErrors(), errors.WithStack(err)
    }
    u.PasswordHash = string(ph)
    return tx.ValidateAndCreate(u)
}
```

### Migrations 迁移 

Create Users Down
创建用户向下

Create Users Up
创建用户向上

```go
// migrations/20180910062057_create_users.down.fizz
drop_table("users")
```

### Templates 模板 

templates/auth/new.html

templates/new/new.html

templates/index.html

```html
// templates/auth/new.html
<style>
  .auth-wrapper{
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .auth-wrapper .sign-form{
    max-width: 350px;
    width: 100%;
    padding: 0 20px;
  }

  .auth-wrapper h1{
    margin-bottom: 20px;
  }
</style>

<div class="auth-wrapper">
  <div class="sign-form">
    <h1>Sign In</h1>

    <%= form_for(user, {action: signinPath()}) { %>
      <%= f.InputTag("Email") %>
      <%= f.InputTag("Password", {type: "password"}) %>
      <button class="btn btn-success">Sign In!</button>
    <% } %>
  </div>
</div>
```