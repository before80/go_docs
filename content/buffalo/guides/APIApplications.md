+++
title = "API 应用程序"
date = 2024-02-04T21:16:19+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/apis/]({{< ref "/buffalo/guides/APIApplications" >}})

# API Applications API 应用程序 

Applications that only serve API end-points, typically JSON, are very different from those that serve HTML, JavaScript, and CSS. In this guide, you’ll learn how to build an API-only app, using Buffalo.

​	仅提供 API 端点（通常为 JSON）的应用程序与提供 HTML、JavaScript 和 CSS 的应用程序非常不同。在本指南中，您将学习如何使用 Buffalo 构建仅限 API 的应用程序。

## Creating a New API Application 创建新的 API 应用程序 

When creating a new Buffalo application using the `buffalo new` command, the optional `--api` flag will generate an application that is better suited to serving APIs than a stock Buffalo application.

​	使用 `buffalo new` 命令创建新的 Buffalo 应用程序时，可选的 `--api` 标志将生成一个比普通 Buffalo 应用程序更适合提供 API 的应用程序。

```bash
$ buffalo new coke --api
```

### Slimmed Project Layout 精简的项目布局 

Applications generated with the `--api` flag don’t contain any front systems. This means there is no templating, stylesheets, etc…

​	使用 `--api` 标志生成的应用程序不包含任何前端系统。这意味着没有模板、样式表等……

API

Default 默认

```bash
$ buffalo new coke --api
├── actions/
│	├── app.go
│	└── render.go
├── cmd/
│	└── app/
│		└── main.go
├── config/
├── fixtures/
├── grifts/
├── locales/
├── models/
├── .buffalo.dev.yml
├── .codeclimate.yml
├── .docketignore
├── .env
├── .gitignore
├── database.yml
├── Dockerfile
├── go.mod
├── go.sum
├── inflections.json
├── README.md
```

### Tuned `actions/app.go actions/render.go` Files 调整的 `actions/app.go actions/render.go` 文件 

API applications have `actions/app.go` and `actions/render.go` files that are a good starting point for API applications.

​	API 应用程序具有 `actions/app.go` 和 `actions/render.go` 文件，它们是 API 应用程序的良好起点。

API

Default 默认

```bash
$ buffalo new coke --api
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_coke_session",
		})
		app.Use(forceSSL())
		app.Use(paramlogger.ParameterLogger)
		app.Use(contenttype.Set("application/json"))

		app.Use(popmw.Transaction(models.DB))
		app.GET("/", HomeHandler)
	}

	return app
}
func init() {
	r = render.New(render.Options{
		DefaultContentType: "application/json",
	})
}
```
