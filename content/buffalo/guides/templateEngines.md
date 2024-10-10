+++
title = "模板引擎"
date = 2024-02-04T21:18:46+08:00
weight = 13
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/logging/]({{< ref "/buffalo/guides/templateEngines" >}})

# Template Engines 模板引擎 

## Map Template Engines by Extension 按扩展名映射模板引擎 

Since **0.10.0**
自 0.10.0 起



Previously you were able to define a single implementation of [`render.TemplateEngine`](https://godoc.org/github.com/gobuffalo/buffalo/render#TemplateEngine) per [`render.Engine`](https://godoc.org/github.com/gobuffalo/buffalo/render#Engine). This has been removed in favor of a map of `render.TemplateEngine`. Now you can map a file extension to an implementation of `render.TemplateEngine`. This means, not only, can you now use multiple template engines in one application, but you can also chain them together.

​	以前，您能够为每个 `render.Engine` 定义 `render.TemplateEngine` 的单个实现。这已被移除，取而代之的是 `render.TemplateEngine` 的映射。现在，您可以将文件扩展名映射到 `render.TemplateEngine` 的实现。这意味着，您不仅可以在一个应用程序中使用多个模板引擎，还可以将它们链接在一起。

For example, if the file was `foo.tmpl.html` it would, by default, first be processed as a Go template, then that result would be sent to the Plush engine.

​	例如，如果文件是 `foo.tmpl.html` ，则默认情况下，它将首先作为 Go 模板进行处理，然后将该结果发送到 Plush 引擎。

Here is a list of default implementations:

​	以下是默认实现的列表：

- `.html` - processed as a Plush template, unchanged from previous releases.
  `.html` - 作为 Plush 模板处理，与以前的版本保持不变。
- `.md` - processed first as Markdown, then as a Plush template, unchanged from previous releases.
  `.md` - 首先作为 Markdown 处理，然后作为 Plush 模板处理，与以前的版本保持不变。
- `.tmpl` - processed as a Go template.
  `.tmpl` - 作为 Go 模板处理。
- `.js` - processed as a Plush template.
  `.js` - 作为 Plush 模板处理。

```go
func init() {
  r = render.New(render.Options{
    // ...
    TemplateEngines: map[string]render.TemplateEngine{
      ".tmpl": GoTemplateEngine,
    },
    // ...
  })
}

func GoTemplateEngine(input string, data map[string]interface{}, helpers map[string]interface{}) (string, error) {
  // since go templates don't have the concept of an optional map argument like Plush does
  // add this "null" map so it can be used in templates like this:
  // {{ partial "flash.html" .nilOpts }}
  data["nilOpts"] = map[string]interface{}{}

  t := template.New(input)
  if helpers != nil {
    t = t.Funcs(helpers)
  }

  t, err := t.Parse(input)
  if err != nil {
    return "", err
  }

  bb := &bytes.Buffer{}
  err = t.Execute(bb, data)
  return bb.String(), err
}
```
