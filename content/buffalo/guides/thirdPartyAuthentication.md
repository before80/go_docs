+++
title = "第三方认证"
date = 2024-02-04T21:17:48+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/goth/](https://gobuffalo.io/documentation/guides/goth/)

# Third Party Authentication 第三方认证 

In many use-cases, you’ll need to implement user authentication in your apps. [Goth](https://github.com/markbates/goth) provides a simple, clean, and idiomatic way to write authentication packages for Go web applications.

​	在许多用例中，您需要在应用程序中实现用户认证。Goth 为编写适用于 Go Web 应用程序的认证包提供了一种简单、干净且惯用的方式。

If you’re looking for authentication through Facebook, Google and others, that’s probably the solution you’re looking for.

​	如果您正在寻找通过 Facebook、Google 等进行认证，那么这可能是您正在寻找的解决方案。

Buffalo had a native support for Goth until version `v0.9.4`. Since then, it was moved into it’s own plugin, https://github.com/gobuffalo/buffalo-goth.
Buffalo 在版本 `v0.9.4` 之前对 Goth 提供了原生支持。从那时起，它被移到了自己的插件中，https://github.com/gobuffalo/buffalo-goth。

## Installation 安装 

To install the `buffalo-goth` plugin, run the following command:

​	要安装 `buffalo-goth` 插件，请运行以下命令：

```bash
$ buffalo plugins install github.com/gobuffalo/buffalo-goth
```

## Generator 生成器 

```bash
$ buffalo g goth twitter facebook linkedin github

--> actions/auth.go
--> go get github.com/markbates/goth/...
--> goimports -w .
```

## Example Usage 示例用法 

```go
// actions/app.go
package actions

import (
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/buffalo/middleware"
  "github.com/gobuffalo/buffalo/middleware/csrf"
  "github.com/gobuffalo/buffalo/middleware/i18n"

  "github.com/markbates/coke/models"

  "github.com/gobuffalo/envy"
  "github.com/gobuffalo/packr"

  "github.com/markbates/goth/gothic"
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

    if ENV == "development" {
      app.Use(middleware.ParameterLogger)
    }
    if ENV != "test" {
      // Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
      // Remove to disable this.
      app.Use(csrf.Middleware)
    }

    // Wraps each request in a transaction.
    //  c.Value("tx").(*pop.PopTransaction)
    // Remove to disable this.
    app.Use(middleware.PopTransaction(models.DB))

    // Setup and use translations:
    var err error
    if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
      app.Stop(err)
    }
    app.Use(T.Middleware())

    app.GET("/", HomeHandler)

    app.ServeFiles("/assets", packr.NewBox("../public/assets"))
    auth := app.Group("/auth")
    auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
    auth.GET("/{provider}/callback", AuthCallback)
  }

  return app
}
// actions/auth.go
package actions

import (
  "fmt"
  "os"

  "github.com/gobuffalo/buffalo"
  "github.com/markbates/goth"
  "github.com/markbates/goth/gothic"
  "github.com/markbates/goth/providers/facebook"
  "github.com/markbates/goth/providers/github"
  "github.com/markbates/goth/providers/linkedin"
  "github.com/markbates/goth/providers/twitter"
)

func init() {
  gothic.Store = App().SessionStore

  goth.UseProviders(
    twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/twitter/callback")),
    facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/facebook/callback")),
    linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/linkedin/callback")),
    github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), fmt.Sprintf("%s%s", App().Host, "/auth/github/callback")),
  )
}

func AuthCallback(c buffalo.Context) error {
  user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
  if err != nil {
    return c.Error(401, err)
  }
  // Do something with the user, maybe register them/sign them in
  return c.Render(200, r.JSON(user))
}
```

## Video Tutorial 视频教程 

{{< vimeo "223666374">}}

## See Also 另请参阅 # 本地认证 - 管理内部用户认证。

- [Local Authentication](https://gobuffalo.io/documentation/guides/auth) - Manage internal users auth.