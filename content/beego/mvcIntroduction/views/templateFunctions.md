+++
title = "templateFunctions"
date = 2024-02-04T10:04:21+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/view/template/]({{< ref "/beego/mvcIntroduction/views/templateFunctions" >}})

# Template Functions 模板函数



## Template Functions 模板函数

Beego supports custom template functions but it must be set as below before `web.Run()`:

​	Beego 支持自定义模板函数，但必须在 `web.Run()` 之前进行如下设置：

```
func hello(in string)(out string){
	out = in + "world"
	return
}

web.AddFuncMap("hi",hello)
```

Then you can use these functions in template:

​	然后你可以在模板中使用这些函数：

```
{{.Content | hi}}
```

Here are Beego’s built-in template functions:

​	以下是 Beego 内置的模板函数：

- dateformat

  Format time, return string. {{dateformat .Time “2006-01-02T15:04:05Z07:00”}}

  ​	格式化时间，返回字符串。{{dateformat .Time “2006-01-02T15:04:05Z07:00”}}

- date

  ​	日期

  This is similar to date function in PHP. It can easily return time by string {{date .T “Y-m-d H:i:s”}}

  ​	这类似于 PHP 中的 date 函数。它可以通过字符串轻松返回时间 {{date .T “Y-m-d H:i:s”}}

- compare

  ​	比较

  Compare two objects. If they are the same return true otherwise return false. {{compare .A .B}}

  ​	比较两个对象。如果它们相同则返回 true，否则返回 false。{{compare .A .B}}

- substr

  Return sub string. supports UTF-8 string. {{substr .Str 0 30}}

  ​	返回子字符串。支持 UTF-8 字符串。{{substr .Str 0 30}}

- html2str

  Parse html to string by removing tags like script and css and return text. {{html2str .Htmlinfo}}

  ​	通过删除脚本和 CSS 等标签将 html 解析为字符串并返回文本。{{html2str .Htmlinfo}}

- str2html Parse string to HTML, no escaping. {{str2html .Strhtml}}

  ​	str2html 将字符串解析为 HTML，不转义。{{str2html .Strhtml}}

- htmlquote

  Escaping html. {{htmlquote .quote}}

  ​	转义 html。{{htmlquote .quote}}

- htmlunquote

  Unescaping to html. {{htmlunquote .unquote}}

  ​	将 html 转义。{{htmlunquote .unquote}}

- renderform

  Generate form from StructTag. {{&struct | renderform}}

  ​	从 StructTag 生成表单。{{&struct | renderform}}

- assets_js

  Create a `<script>` tag from js src. {{assets_js src}}

  ​	从 js src 创建 `<script>` 标记。 {{assets_js src}}

- assets_css

  Create a `<link>` tag from css src. {{assets_css src}}

  ​	从 css src 创建 `<link>` 标签。 {{assets_css src}}

- config

  Get the value of AppConfig. {{config configType configKey defaultValue}}. configType must be String, Bool, Int, Int64, Float, or DIY

  ​	获取 AppConfig 的值。 {{config configType configKey defaultValue}}。configType 必须是 String、Bool、Int、Int64、Float 或 DIY

- map_get

  Get value of `map` by key

  ​	按键获取 `map` 的值

  ```
    // In controller
    Data["m"] = map[string]interface{} {
        "a": 1,
        "1": map[string]float64{
            "c": 4,
        },
    }
  
    // In view
    {{ map_get m "a" }} // return 1
    {{ map_get m 1 "c" }} // return 4
  ```

- urlfor

  Get the URL of a controller method

  ​	获取控制器方法的 URL

  ```
    {{urlfor "TestController.List"}}
  ```

  [more details
  更多详情]({{< ref "/beego/mvcIntroduction/controllers/urlBuilding" >}})
