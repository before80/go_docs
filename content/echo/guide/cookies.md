+++
title = "Cookies"
weight = 50
date = 2023-07-09T21:50:18+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Cookies

> 原文：[https://echo.labstack.com/docs/cookies](https://echo.labstack.com/docs/cookies)

​	Cookie（Cookie）是从网站服务器发送并存储在用户的网络浏览器中的一小段数据。每次用户加载网站时，浏览器会将 Cookie 发送回服务器，通知服务器用户的最新活动。Cookie 被设计为网站记住有状态信息（例如，在在线商店中添加到购物车的商品）或记录用户的浏览活动（例如，点击特定按钮、登录或用户先前访问的网站页面）的可靠机制。Cookie 还可以存储用户先前输入的表单内容，例如用户名、性别、年龄、地址等。

## Cookie 属性

| 属性       | 是否可选 |
| ---------- | -------- |
| `Name`     | No       |
| `Value`    | No       |
| `Path`     | Yes      |
| `Domain`   | Yes      |
| `Expires`  | Yes      |
| `Secure`   | Yes      |
| `HttpOnly` | Yes      |

​		Echo 使用 Go 标准的 `http.Cookie` 对象来在处理程序函数中添加/获取 Cookie。

## 创建 Cookie

```go
func writeCookie(c echo.Context) error {
    cookie := new(http.Cookie)
    cookie.Name = "username"
    cookie.Value = "jon"
    cookie.Expires = time.Now().Add(24 * time.Hour)
    c.SetCookie(cookie)
    return c.String(http.StatusOK, "write a cookie")
}
```



- 使用 `new(http.Cookie)` 创建 Cookie。
- 通过给 `http.Cookie` 实例的公共属性赋值来设置 Cookie 的属性。
- 最后，`c.SetCookie(cookie)` 在 HTTP 响应中添加一个 `Set-Cookie` 头。

## 读取 Cookie

```go
func readCookie(c echo.Context) error {
    cookie, err := c.Cookie("username")
    if err != nil {
        return err
    }
    fmt.Println(cookie.Name)
    fmt.Println(cookie.Value)
    return c.String(http.StatusOK, "read a cookie")
}
```



- 使用 `c.Cookie("username")` 按名称从 HTTP 请求中读取 Cookie。
- 使用 `Getter` 函数访问 Cookie 的属性。

## 读取所有 Cookie

```go
func readAllCookies(c echo.Context) error {
    for _, cookie := range c.Cookies() {
        fmt.Println(cookie.Name)
        fmt.Println(cookie.Value)
    }
    return c.String(http.StatusOK, "read all the cookies")
}
```