+++
title = "模型"
date = 2024-02-04T09:10:33+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/quickstart/model/]({{< ref "/beego/quickStart/models" >}})

# Models 模型



## Creating models 创建模型

Models are normally the best way to handle the numerous databases used in web applications. The `bee new` project does not contain an example of models. Demos on implementing and using models can instead be found in `bee api` projects.

​	模型通常是处理 Web 应用程序中使用的众多数据库的最佳方式。 `bee new` 项目不包含模型示例。相反，可以在 `bee api` 项目中找到有关实现和使用模型的演示。

The Controller can automatically handle models for simple applications.

​	对于简单的应用程序，控制器可以自动处理模型。

Larger applications with more reusable code requiring logic separation must use models. Reusable logic can be factored out into a Model and used to handle database interactions. The following is an example:

​	需要逻辑分离且具有更多可重用代码的大型应用程序必须使用模型。可将可重用逻辑分解到模型中，并用于处理数据库交互。以下是一个示例：

```
package models

import (
	"loggo/utils"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	NotPV []string = []string{"css", "js", "class", "gif", "jpg", "jpeg", "png", "bmp", "ico", "rss", "xml", "swf"}
)

const big = 0xFFFFFF

func LogPV(urls string) bool {
	ext := filepath.Ext(urls)
	if ext == "" {
		return true
	}
	for _, v := range NotPV {
		if v == strings.ToLower(ext) {
			return false
		}
	}
	return true
}
```

Please see [MVC Models]({{< ref "/beego/mvcIntroduction/models/overview" >}}) for the specific examples of database models and Beego’s ORM framework. [The next section]({{< ref "/beego/quickStart/view" >}}) will cover writing views.

​	有关数据库模型和 Beego 的 ORM 框架的具体示例，请参阅 MVC 模型。下一节将介绍编写视图。
