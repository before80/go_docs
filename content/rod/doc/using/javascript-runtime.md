+++
title = "JavaScript 运行时"
date = 2024-11-21T08:09:20+08:00
weight = 80
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/javascript-runtime](https://go-rod.github.io/i18n/zh-CN/#/javascript-runtime)
>
> 收录该文档时间： `2024-11-21T08:09:20+08:00`

# JavaScript 运行时

​	我们可以使用 Rod 在页面上执行任意 JavaScript 代码。 例如用它来读取或修改页面的 HTML 内容。

## 在页面上 eval

​	例如，使用 `Page.Eval` 设置全局变量：

```go
page.MustEval(`() => window.a = {name: 'jack'}`)
```

​	我们可以使用 js 函数的参数来获取 golang 变量， 这是如何从 golang 中获取 `key` 和 `data` 数据并将其设置为页面 window 上的对象 `k`

```go
key := "a"
data := map[string]string{"name": "jack"}
page.MustEval(`(k, val) => {
    window[k] = val
}`, key, data)
```

​	从 Eval 获取返回值：

```go
val := page.MustEval(`() => a`).Get("name").Str()
fmt.Println(val) // output: jack
```

​	`Eval` 只接受 js 函数，如下代码是无效的:

```go
page.MustEval(`a`) // 将失败
```

## 在元素上 eval

​	`Element.Eval` 和 `Page.Eval` 类似，但是对于前者来说，`this` 对象代表当前元素。 例如在页面上有一个 `<button>Submit</button>` ，我们可以用 JS 来读取或修改元素：

```go
el := page.MustElement("button")
el.MustEval(`() => this.innerText = "Apply"`) // 更改内容
txt := el.MustEval(`() => this.innerText`).Str()
fmt.Println(txt) // 输出: Apply
```

## 将 Go 函数暴露给页面

​	我们使用 `Page.Expose` 来把回调函数暴露给页面。 例如，我们可以像这样暴露函数，来帮助页面计算 md5：

```go
page.MustExpose("md5", func(g gson.JSON) (interface{}, error) {
    return md5.Sum([]byte(g.Str())), nil
})
```

​	现在页面可以在 window 对象上调用这个方法：

```go
hash := page.MustEval(`() => window.md5("test")`).Str()
fmt.Println(hash)
```

​	下面是另一个例子来获取页面上的按钮事件：

```go
page.MustExpose("myClick", func(v gson.JSON) (interface{}, error) {
    fmt.Println("点击了")
    return nil, nil
})
```

​	点击按钮时调用 'myClick' 方法：

```go
page.MustElement("button").MustEval(`() => this.onclick = myClick`)
```
