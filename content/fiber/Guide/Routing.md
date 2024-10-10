+++
title = "Routing"
date = 2024-02-05T09:14:15+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/guide/routing]({{< ref "/fiber/Guide/Routing" >}})

# 🔌 Routing

## Handlers 处理程序

Registers a route bound to a specific [HTTP method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods).

​	注册绑定到特定 HTTP 方法的路由。

Signatures
签名

```go
// HTTP methods
func (app *App) Get(path string, handlers ...Handler) Router
func (app *App) Head(path string, handlers ...Handler) Router
func (app *App) Post(path string, handlers ...Handler) Router
func (app *App) Put(path string, handlers ...Handler) Router
func (app *App) Delete(path string, handlers ...Handler) Router
func (app *App) Connect(path string, handlers ...Handler) Router
func (app *App) Options(path string, handlers ...Handler) Router
func (app *App) Trace(path string, handlers ...Handler) Router
func (app *App) Patch(path string, handlers ...Handler) Router

// Add allows you to specifiy a method as value
func (app *App) Add(method, path string, handlers ...Handler) Router

// All will register the route on all HTTP methods
// Almost the same as app.Use but not bound to prefixes
func (app *App) All(path string, handlers ...Handler) Router
```



Examples
示例

```go
// Simple GET handler
app.Get("/api/list", func(c *fiber.Ctx) error {
  return c.SendString("I'm a GET request!")
})

// Simple POST handler
app.Post("/api/register", func(c *fiber.Ctx) error {
  return c.SendString("I'm a POST request!")
})
```



**Use** can be used for middleware packages and prefix catchers. These routes will only match the beginning of each path i.e. `/john` will match `/john/doe`, `/johnnnnn` etc

​	Use 可用于中间件包和前缀捕获器。这些路由将仅匹配每个路径的开头，即 `/john` 将匹配 `/john/doe` 、 `/johnnnnn` 等

Signature
签名

```go
func (app *App) Use(args ...interface{}) Router
```



Examples
示例

```go
// Match any request
app.Use(func(c *fiber.Ctx) error {
    return c.Next()
})

// Match request starting with /api
app.Use("/api", func(c *fiber.Ctx) error {
    return c.Next()
})

// Match requests starting with /api or /home (multiple-prefix support)
app.Use([]string{"/api", "/home"}, func(c *fiber.Ctx) error {
    return c.Next()
})

// Attach multiple handlers 
app.Use("/api", func(c *fiber.Ctx) error {
  c.Set("X-Custom-Header", random.String(32))
    return c.Next()
}, func(c *fiber.Ctx) error {
    return c.Next()
})
```



## Paths 路径

Route paths, combined with a request method, define the endpoints at which requests can be made. Route paths can be **strings** or **string patterns**.

​	路由路径与请求方法结合使用，定义可以进行请求的端点。路由路径可以是字符串或字符串模式。

**Examples of route paths based on strings
基于字符串的路由路径示例**

```go
// This route path will match requests to the root route, "/":
app.Get("/", func(c *fiber.Ctx) error {
      return c.SendString("root")
})

// This route path will match requests to "/about":
app.Get("/about", func(c *fiber.Ctx) error {
    return c.SendString("about")
})

// This route path will match requests to "/random.txt":
app.Get("/random.txt", func(c *fiber.Ctx) error {
    return c.SendString("random.txt")
})
```



As with the expressJs framework, the order of the route declaration plays a role. When a request is received, the routes are checked in the order in which they are declared.

​	与 expressJs 框架一样，路由声明的顺序起着作用。收到请求时，将按声明的顺序检查路由。

INFO
信息

So please be careful to write routes with variable parameters after the routes that contain fixed parts, so that these variable parts do not match instead and unexpected behavior occurs.

​	因此，请务必小心地在包含固定部分的路由之后编写具有可变参数的路由，以便这些可变部分不会匹配，从而导致意外行为。

## Parameters 参数

Route parameters are dynamic elements in the route, which are **named** or **not named segments**. This segments that are used to capture the values specified at their position in the URL. The obtained values can be retrieved using the [Params](https://fiber.wiki/context#params) function, with the name of the route parameter specified in the path as their respective keys or for unnamed parameters the character(*, +) and the counter of this.

​	路由参数是路由中的动态元素，它们是命名或未命名的段。这些段用于捕获在 URL 中其位置处指定的值。可以使用 Params 函数检索获得的值，其中路由参数的名称指定在路径中作为其各自的键，或者对于未命名的参数，使用字符(*, +)及其计数器。

The characters :, +, and * are characters that introduce a parameter.

​	字符 :, +, 和 * 是引入参数的字符。

Greedy parameters are indicated by wildcard(*) or plus(+) signs.

​	贪婪参数由通配符(*)或加号(+)符号表示。

The routing also offers the possibility to use optional parameters, for the named parameters these are marked with a final "?", unlike the plus sign which is not optional, you can use the wildcard character for a parameter range which is optional and greedy.

​	路由还提供使用可选参数的可能性，对于命名参数，这些参数用最终的“？”标记，与非可选的加号不同，您可以对参数范围使用通配符，该范围是可选的且贪婪的。

**Example of define routes with route parameters
使用路由参数定义路由的示例**

```go
// Parameters
app.Get("/user/:name/books/:title", func(c *fiber.Ctx) error {
    fmt.Fprintf(c, "%s\n", c.Params("name"))
    fmt.Fprintf(c, "%s\n", c.Params("title"))
    return nil
})
// Plus - greedy - not optional
app.Get("/user/+", func(c *fiber.Ctx) error {
    return c.SendString(c.Params("+"))
})

// Optional parameter
app.Get("/user/:name?", func(c *fiber.Ctx) error {
    return c.SendString(c.Params("name"))
})

// Wildcard - greedy - optional
app.Get("/user/*", func(c *fiber.Ctx) error {
    return c.SendString(c.Params("*"))
})

// This route path will match requests to "/v1/some/resource/name:customVerb", since the parameter character is escaped
app.Get(`/v1/some/resource/name\:customVerb`, func(c *fiber.Ctx) error {
    return c.SendString("Hello, Community")
})
```



INFO
信息

Since the hyphen (`-`) and the dot (`.`) are interpreted literally, they can be used along with route parameters for useful purposes.

​	由于连字符 ( `-` ) 和点 ( `.` ) 被逐字解释，因此可以与路由参数一起用于有用的目的。

INFO
信息

All special parameter characters can also be escaped with `"\\"` and lose their value, so you can use them in the route if you want, like in the custom methods of the [google api design guide](https://cloud.google.com/apis/design/custom_methods). It's recommended to use backticks ``` because in go's regex documentation, they always use backticks to make sure it is unambiguous and the escape character doesn't interfere with regex patterns in an unexpected way.

​	所有特殊参数字符也可以用 `"\\"` 转义并失去其值，因此您可以根据需要在路由中使用它们，例如在 google api 设计指南的自定义方法中。建议使用反引号 ``` ，因为在 go 的正则表达式文档中，它们始终使用反引号以确保它没有歧义，并且转义字符不会以意外的方式干扰正则表达式模式。

```go
// http://localhost:3000/plantae/prunus.persica
app.Get("/plantae/:genus.:species", func(c *fiber.Ctx) error {
    fmt.Fprintf(c, "%s.%s\n", c.Params("genus"), c.Params("species"))
    return nil // prunus.persica
})
```



```go
// http://localhost:3000/flights/LAX-SFO
app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
    fmt.Fprintf(c, "%s-%s\n", c.Params("from"), c.Params("to"))
    return nil // LAX-SFO
})
```



Our intelligent router recognizes that the introductory parameter characters should be part of the request route in this case and can process them as such.

​	我们的智能路由器识别出在这种情况下，介绍性参数字符应成为请求路由的一部分，并可以按原样处理它们。

```go
// http://localhost:3000/shop/product/color:blue/size:xs
app.Get("/shop/product/color::color/size::size", func(c *fiber.Ctx) error {
    fmt.Fprintf(c, "%s:%s\n", c.Params("color"), c.Params("size"))
    return nil // blue:xs
})
```



In addition, several parameters in a row and several unnamed parameter characters in the route, such as the wildcard or plus character, are possible, which greatly expands the possibilities of the router for the user.

​	此外，一行中的多个参数和路由中的多个未命名参数字符（例如通配符或加号字符）也是可能的，这极大地扩展了路由器对用户的功能。

```go
// GET /@v1
// Params: "sign" -> "@", "param" -> "v1"
app.Get("/:sign:param", handler)

// GET /api-v1
// Params: "name" -> "v1" 
app.Get("/api-:name", handler)

// GET /customer/v1/cart/proxy
// Params: "*1" -> "customer/", "*2" -> "/cart"
app.Get("/*v1*/proxy", handler)

// GET /v1/brand/4/shop/blue/xs
// Params: "*1" -> "brand/4", "*2" -> "blue/xs"
app.Get("/v1/*/shop/*", handler)
```



We have adapted the routing strongly to the express routing, but currently without the possibility of the regular expressions, because they are quite slow. The possibilities can be tested with version 0.1.7 (express 4) in the online [Express route tester](http://forbeslindesay.github.io/express-route-tester/).

​	我们已经将路由强烈地调整为 express 路由，但目前没有正则表达式的可能性，因为它们非常慢。可以在在线 Express 路由测试器中使用版本 0.1.7（express 4）测试这些可能性。

### Constraints 约束

Route constraints execute when a match has occurred to the incoming URL and the URL path is tokenized into route values by parameters. The feature was intorduced in `v2.37.0` and inspired by [.NET Core](https://docs.microsoft.com/en-us/aspnet/core/fundamentals/routing?view=aspnetcore-6.0#route-constraints).

​	当与传入 URL 匹配并且 URL 路径通过参数标记化为路由值时，路由约束会执行。此功能在 `v2.37.0` 中引入，并受到 .NET Core 的启发。

CAUTION
注意

Constraints aren't validation for parameters. If constraint aren't valid for parameter value, Fiber returns **404 handler**.

​	约束不是对参数的验证。如果约束对参数值无效，Fiber 会返回 404 处理程序。

| Constraint 约束   | Example 示例                    | Example matches 示例匹配                                     |
| ----------------- | ------------------------------- | ------------------------------------------------------------ |
| int               | :id<int>                        | 123456789, -123456789                                        |
| bool              | :active<bool>                   | true,false                                                   |
| guid              | :id<guid>                       | CD2C1638-1638-72D5-1638-DEADBEEF1638                         |
| float             | :weight<float>                  | 1.234, -1,001.01e8                                           |
| minLen(value)     | :username<minLen(4)>            | Test (must be at least 4 characters) 测试（必须至少 4 个字符） |
| maxLen(value)     | :filename<maxLen(8)>            | MyFile (must be no more than 8 characters MyFile（必须不超过 8 个字符 |
| len(length)       | :filename<len(12)>              | somefile.txt (exactly 12 characters) somefile.txt（恰好 12 个字符） |
| min(value)        | :age<min(18)>                   | 19 (Integer value must be at least 18) 19（整数必须至少为 18） |
| max(value)        | :age<max(120)>                  | 91 (Integer value must be no more than 120) 91（整数值必须不超过 120） |
| range(min,max)    | :age<range(18,120)>             | 91 (Integer value must be at least 18 but no more than 120) 91（整数值必须至少为 18 但不超过 120） |
| alpha             | :name<alpha>                    | Rick (String must consist of one or more alphabetical characters, a-z and case-insensitive) Rick（字符串必须包含一个或多个字母字符，a-z 且不区分大小写） |
| datetime          | :dob<datetime(2006\\-01\\-02)>  | 2005-11-01                                                   |
| regex(expression) | :date<regex(\d{4}-\d{2}-\d{2})> | 2022-08-27 (Must match regular expression) 2022-08-27（必须匹配正则表达式） |

**Examples
示例**

- Single Constraint
  单一约束
- Multiple Constraints
  多重约束
- Regex Constraint
  正则约束

```go
app.Get("/:test<min(5)>", func(c *fiber.Ctx) error {
  return c.SendString(c.Params("test"))
})

// curl -X GET http://localhost:3000/12
// 12

// curl -X GET http://localhost:3000/1
// Cannot GET /1
```



CAUTION
注意

You should use `\\` before routing-specific characters when to use datetime constraint (`*`, `+`, `?`, `:`, `/`, `<`, `>`, `;`, `(`, `)`), to avoid wrong parsing.

​	在使用日期时间约束（ `\\` 、 `*` 、 `+` 、 `?` 、 `:` 、 `/` 、 `<` 、 `>` 、 `;` 、 `(` 、 `)` ）时，您应该在路由特定字符之前使用 `\\` ，以避免错误解析。

**Optional Parameter Example
可选参数示例**

You can impose constraints on optional parameters as well.

​	您还可以对可选参数施加约束。

```go
app.Get("/:test<int>?", func(c *fiber.Ctx) error {
  return c.SendString(c.Params("test"))
})
// curl -X GET http://localhost:3000/42
// 42
// curl -X GET http://localhost:3000/
//
// curl -X GET http://localhost:3000/7.0
// Cannot GET /7.0
```



## Middleware 中间件

Functions that are designed to make changes to the request or response are called **middleware functions**. The [Next]({{< ref "/fiber/API/Ctx#next" >}}) is a **Fiber** router function, when called, executes the **next** function that **matches** the current route.

​	旨在对请求或响应进行更改的函数称为中间件函数。Next 是一个 Fiber 路由器函数，在调用时，执行与当前路由匹配的下一个函数。

**Example of a middleware function
中间件函数示例**

```go
app.Use(func(c *fiber.Ctx) error {
  // Set a custom header on all responses:
  c.Set("X-Custom-Header", "Hello, World")

  // Go to next middleware:
  return c.Next()
})

app.Get("/", func(c *fiber.Ctx) error {
  return c.SendString("Hello, World!")
})
```



`Use` method path is a **mount**, or **prefix** path, and limits middleware to only apply to any paths requested that begin with it.

​	 `Use` 方法路径是一个挂载或前缀路径，它限制中间件仅适用于以它开头的任何请求路径。

### Constraints on Adding Routes Dynamically 动态添加路由的约束

CAUTION
注意

Adding routes dynamically after the application has started is not supported due to design and performance considerations. Make sure to define all your routes before the application starts.

​	由于设计和性能方面的考虑，不支持在应用程序启动后动态添加路由。请务必在应用程序启动前定义所有路由。

## Grouping 分组

If you have many endpoints, you can organize your routes using `Group`.

​	如果您有许多端点，可以使用 `Group` 来组织您的路由。

```go
func main() {
  app := fiber.New()

  api := app.Group("/api", middleware) // /api

  v1 := api.Group("/v1", middleware)   // /api/v1
  v1.Get("/list", handler)             // /api/v1/list
  v1.Get("/user", handler)             // /api/v1/user

  v2 := api.Group("/v2", middleware)   // /api/v2
  v2.Get("/list", handler)             // /api/v2/list
  v2.Get("/user", handler)             // /api/v2/user

  log.Fatal(app.Listen(":3000"))
}
```



More information about this in our [Grouping Guide]({{< ref "/fiber/Guide/Grouping" >}})

​	有关此内容的更多信息，请参阅我们的分组指南
