+++
title = "全局变量"
date = 2024-02-04T10:05:01+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/view/global_variables/]({{< ref "/beego/mvcIntroduction/views/globalVariables" >}})

# Global Variables 全局变量



## Global Variables 全局变量

Variables common to all requests can be added using a base controller using the `Prepare` method.

​	可以使用基本控制器通过 `Prepare` 方法添加对所有请求通用的变量。

Usage examples include the username of the current user for display in a toolbar or the request URL for highlighting active menu items.

​	使用示例包括工具栏中显示的当前用户的用户名或用于突出显示活动菜单项的请求 URL。

First, create a common/base controller in the controllers package:

​	首先，在 controllers 包中创建一个 common/base 控制器：

```
app_root
├── controllers
│   ├── base.go <-- base controller
│   └── default.go
└── main.go
```

Your base controller should embed the `web.Controller` from `github.com/beego/beego/v2/server/web`. From here, the `Prepare` method should be defined, containing any logic required for global variables:

​	基本控制器应从 `github.com/beego/beego/v2/server/web` 嵌入 `web.Controller` 。在此，应定义 `Prepare` 方法，其中包含全局变量所需的任何逻辑：

```go
// app_root/controllers/base.go
package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

// Runs after Init before request function execution
func (c *BaseController) Prepare() {
	c.Data["RequestUrl"] = c.Ctx.Input.URL()
}

// Runs after request function execution
func (c *BaseController) Finish() {
	// Any cleanup logic common to all requests goes here. Logging or metrics, for example.
}
```

All other controllers should embed `BaseController` instead of `web.Controller`:

​	所有其他控制器应嵌入 `BaseController` 而不是 `web.Controller` ：

```go
// app_root/controllers/default.go
package controllers

type DefaultController struct {
	BaseController
}

func (c *DefaultController) Index() {
	// your controller logic
}
```

From here your views can access these global variables, in both individual templates and all parent templates:

​	在此，视图可以在各个模板和所有父模板中访问这些全局变量：

```gotemplate
<div>{{ $.RequestUrl }}</div>
```
