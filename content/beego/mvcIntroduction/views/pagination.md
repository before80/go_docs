+++
title = "分页"
date = 2024-02-04T10:04:45+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/view/page/]({{< ref "/beego/mvcIntroduction/views/pagination" >}})



# Pagination 分页



## Pagination 分页

We use often pagination in templates. How do we do that? We have created this demo for pagination. We hope it’s useful to you.

​	我们在模板中经常使用分页。我们如何做到这一点？我们为此演示创建了分页。我们希望对您有用。

## Controllers 控制器

Before you can use the paginator in the view you have to set it in your controller:

​	在视图中使用分页器之前，您必须在控制器中设置它：

```
package controllers

type PostsController struct {
  web.Controller
}

func (this *PostsController) ListAllPosts() {
    // sets this.Data["paginator"] with the current offset (from the url query param)
    postsPerPage := 20
  	paginator := pagination.SetPaginator(this, postsPerPage, CountPosts())

    // fetch the next 20 posts
    this.Data["posts"] = ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
}
```

## Views 视图

Example templates (using Twitter Bootstrap):

​	示例模板（使用 Twitter Bootstrap）：

https://github.com/beego/wetalk/blob/master/views/base/paginator.html
