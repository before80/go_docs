+++
title = "Go 1.22 的路由增强功能"
date = 2024-02-22T20:32:05+08:00
weight = 980
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/routing-enhancements](https://go.dev/blog/routing-enhancements)

# Routing Enhancements for Go 1.22 - Go 1.22 的路由增强功能

Jonathan Amsterdam, on behalf of the Go team

​	Jonathan Amsterdam，代表 Go 团队

13 February 2024

​	2024 年 2 月 13 日

Go 1.22 brings two enhancements to the `net/http` package’s router: method matching and wildcards. These features let you express common routes as patterns instead of Go code. Although they are simple to explain and use, it was a challenge to come up with the right rules for selecting the winning pattern when several match a request.

​	Go 1.22 为 `net/http` 包的路由器带来了两项增强功能：方法匹配和通配符。这些功能让您可以将常见路由表示为模式，而不是 Go 代码。尽管它们很容易解释和使用，但当多个模式匹配请求时，想出用于选择获胜模式的正确规则却是一项挑战。

We made these changes as part of our continuing effort to make Go a great language for building production systems. We studied many third-party web frameworks, extracted what we felt were the most used features, and integrated them into `net/http`. Then we validated our choices and improved our design by collaborating with the community in a [GitHub discussion](https://github.com/golang/go/discussions/60227) and a [proposal issue](https://go.dev/issue/61410). Adding these features to the standard library means one fewer dependency for many projects. But third-party web frameworks remain a fine choice for current users or programs with advanced routing needs.

​	我们做出这些更改是为了继续努力，让 Go 成为构建生产系统的出色语言。我们研究了许多第三方 Web 框架，提取了我们认为使用最多的功能，并将它们集成到 `net/http` 中。然后，我们通过在 GitHub 讨论和提案问题中与社区合作来验证我们的选择并改进我们的设计。将这些功能添加到标准库意味着许多项目减少了一个依赖项。但对于当前用户或具有高级路由需求的程序来说，第三方 Web 框架仍然是一个不错的选择。

## Enhancements 增强功能

The new routing features almost exclusively affect the pattern string passed to the two `net/http.ServeMux` methods `Handle` and `HandleFunc`, and the corresponding top-level functions `http.Handle` and `http.HandleFunc`. The only API changes are two new methods on `net/http.Request` for working with wildcard matches.

​	新的路由功能几乎只影响传递给两个 `net/http.ServeMux` 方法 `Handle` 和 `HandleFunc` 的模式字符串，以及相应的顶级函数 `http.Handle` 和 `http.HandleFunc` 。唯一的 API 更改是 `net/http.Request` 上用于处理通配符匹配的两个新方法。

We’ll illustrate the changes with a hypothetical blog server in which every post has an integer identifier. A request like `GET /posts/234` retrieves the post with ID 234. Before Go 1.22, the code for handling those requests would start with a line like this:

​	我们将通过一个假设的博客服务器来说明这些更改，其中每篇文章都有一个整数标识符。像 `GET /posts/234` 这样的请求会检索 ID 为 234 的文章。在 Go 1.22 之前，处理这些请求的代码将以类似这样的行开头：

```go
http.Handle("/posts/", handlePost)
```

The trailing slash routes all requests beginning `/posts/` to the `handlePost` function, which would have to check that the HTTP method was GET, extract the identifier, and retrieve the post. Since the method check isn’t strictly necessary to satisfy the request, it would be a natural mistake to omit it. That would mean that a request like `DELETE /posts/234` would fetch the post, which is surprising at the least.

​	尾随斜杠将所有以 `/posts/` 开头的请求路由到 `handlePost` 函数，该函数必须检查 HTTP 方法是否为 GET，提取标识符并检索文章。由于方法检查对于满足请求并不是严格必要的，因此很容易忘记它。这意味着像 `DELETE /posts/234` 这样的请求将获取文章，这至少令人惊讶。

In Go 1.22, the existing code will continue to work, or you could instead write this:

​	在 Go 1.22 中，现有代码将继续工作，或者您可以改为编写以下内容：

```go
http.Handle("GET /posts/{id}", handlePost2)
```

This pattern matches a GET request whose path begins “/posts/” and has two segments. (As a special case, GET also matches HEAD; all the other methods match exactly.) The `handlePost2` function no longer needs to check the method, and extracting the identifier string can be written using the new `PathValue` method on `Request`:

​	此模式匹配路径以“/posts/”开头且具有两个片段的 GET 请求。（作为特例，GET 也匹配 HEAD；所有其他方法完全匹配。） `handlePost2` 函数不再需要检查方法，并且可以使用 `Request` 上的新 `PathValue` 方法来编写提取标识符字符串：

```go
idString := req.PathValue("id")
```

The rest of `handlePost2` would behave like `handlePost`, converting the string identifier to an integer and fetching the post.

​	 `handlePost2` 的其余部分将表现得像 `handlePost` ，将字符串标识符转换为整数并获取帖子。

Requests like `DELETE /posts/234` will fail if no other matching pattern is registered. In accordance with [HTTP semantics](https://httpwg.org/specs/rfc9110.html#status.405), a `net/http` server will reply to such a request with a `405 Method Not Allowed` error that lists the available methods in an `Allow` header.

​	如果未注册其他匹配模式，则类似 `DELETE /posts/234` 的请求将失败。根据 HTTP 语义， `net/http` 服务器将使用 `Allow` 标头中列出的可用方法来回复此类请求，并显示 `405 Method Not Allowed` 错误。

A wildcard can match an entire segment, like `{id}` in the example above, or if it ends in `...` it can match all the remaining segments of the path, as in the pattern `/files/{pathname...}`.

​	通配符可以匹配整个片段，如上例中的 `{id}` ，或者如果它以 `...` 结尾，它可以匹配路径的所有剩余片段，如模式 `/files/{pathname...}` 中所示。

There is one last bit of syntax. As we showed above, patterns ending in a slash, like `/posts/`, match all paths beginning with that string. To match only the path with the trailing slash, you can write `/posts/{$}`. That will match `/posts/` but not `/posts` or `/posts/234`.

​	还有一种最后的语法。如上所示，以斜杠结尾的模式（如 `/posts/` ）匹配以该字符串开头的所有路径。要仅匹配带有尾随斜杠的路径，您可以编写 `/posts/{$}` 。这将匹配 `/posts/` ，但不匹配 `/posts` 或 `/posts/234` 。

And there is one last bit of API: `net/http.Request` has a `SetPathValue` method so that routers outside the standard library can make the results of their own path parsing available via `Request.PathValue`.

​	最后还有一点 API： `net/http.Request` 有一个 `SetPathValue` 方法，因此标准库之外的路由器可以通过 `Request.PathValue` 提供它们自己的路径解析结果。

## Precedence 优先级

Every HTTP router must deal with overlapping patterns, like `/posts/{id}` and `/posts/latest`. Both of these patterns match the path “posts/latest”, but at most one can serve the request. Which pattern takes precedence?

​	每个 HTTP 路由器都必须处理重叠模式，例如 `/posts/{id}` 和 `/posts/latest` 。这两个模式都匹配路径“posts/latest”，但最多只能有一个模式为请求提供服务。哪个模式具有优先权？

Some routers disallow overlaps; others use the pattern that was registered last. Go has always allowed overlaps, and has chosen the longer pattern regardless of registration order. Preserving order-independence was important to us (and necessary for backwards compatibility), but we needed a better rule than “longest wins.” That rule would select `/posts/latest` over `/posts/{id}`, but would choose `/posts/{identifier}` over both. That seems wrong: the wildcard name shouldn’t matter. It feels like `/posts/latest` should always win this competition, because it matches a single path instead of many.

​	一些路由器不允许重叠；其他路由器使用最后注册的模式。Go 一直允许重叠，并且无论注册顺序如何，都会选择较长的模式。对我们来说，保持独立于顺序很重要（并且对于向后兼容性来说是必要的），但我们需要一个比“最长者获胜”更好的规则。该规则会选择 `/posts/latest` 而不是 `/posts/{id}` ，但会选择 `/posts/{identifier}` 而不是两者。这似乎是错误的：通配符名称不应重要。感觉上 `/posts/latest` 应该始终赢得此竞争，因为它匹配的是一个路径而不是多个路径。

Our quest for a good precedence rule led us to consider many properties of patterns. For example, we considered preferring the pattern with the longest literal (non-wildcard) prefix. That would choose `/posts/latest` over `/posts/ {id}`. But it wouldn’t distinguish between `/users/{u}/posts/latest` and `/users/{u}/posts/{id}`, and it seems like the former should take precedence.

​	我们为了寻找一个好的优先级规则，考虑了许多模式的属性。例如，我们考虑优先选择具有最长文字（非通配符）前缀的模式。这会选择 `/posts/latest` 而不是 `/posts/ {id}` 。但它不会区分 `/users/{u}/posts/latest` 和 `/users/{u}/posts/{id}` ，而前者似乎应该具有优先权。

We eventually chose a rule based on what the patterns mean instead of how they look. Every valid pattern matches a set of requests. For example, `/posts/latest` matches requests with the path `/posts/latest`, while `/posts/{id}` matches requests with any two-segment path whose first segment is “posts”. We say that one pattern is *more specific* than another if it matches a strict subset of requests. The pattern `/posts/latest` is more specific than `/posts/{id}` because the latter matches every request that the former does, and more.

​	我们最终选择了一个基于模式的含义而不是外观的规则。每个有效模式都匹配一组请求。例如， `/posts/latest` 匹配路径为 `/posts/latest` 的请求，而 `/posts/{id}` 匹配任何第一个片段为“posts”的两段路径的请求。如果一个模式匹配的请求严格是另一个模式的子集，则我们说一个模式比另一个模式更具体。模式 `/posts/latest` 比 `/posts/{id}` 更具体，因为后者匹配前者匹配的每个请求，以及更多请求。

The precedence rule is simple: the most specific pattern wins. This rule matches our intuition that `posts/latests` should be preferred to `posts/{id}`, and `/users/{u}/posts/latest` should be preferred to `/users/{u}/posts/{id}`. It also makes sense for methods. For example, `GET /posts/{id}` takes precedence over `/posts/{id}` because the first only matches GET and HEAD requests, while the second matches requests with any method.

​	优先级规则很简单：最具体的模式获胜。此规则符合我们的直觉，即 `posts/latests` 应优于 `posts/{id}` ，而 `/users/{u}/posts/latest` 应优于 `/users/{u}/posts/{id}` 。对于方法来说，这也是有意义的。例如， `GET /posts/{id}` 优先于 `/posts/{id}` ，因为第一个只匹配 GET 和 HEAD 请求，而第二个匹配任何方法的请求。

The “most specific wins” rule generalizes the original “longest wins” rule for the path parts of original patterns, those without wildcards or `{$}`. Such patterns only overlap when one is a prefix of the other, and the longer is the more specific.

​	“最具体获胜”规则概括了原始模式路径部分的原始“最长获胜”规则，这些模式没有通配符或 `{$}` 。此类模式仅在其中一个为另一个的前缀时才重叠，并且较长的模式更具体。

What if two patterns overlap but neither is more specific? For example, `/posts/{id}` and `/{resource}/latest` both match `/posts/latest`. There is no obvious answer to which takes precedence, so we consider these patterns to conflict with each other. Registering both of them (in either order!) will panic.

​	如果两个模式重叠，但都不是更具体的，该怎么办？例如， `/posts/{id}` 和 `/{resource}/latest` 都匹配 `/posts/latest` 。对于哪个优先级没有明显的答案，因此我们认为这些模式相互冲突。注册它们两者（按任何顺序！）都会引发恐慌。

The precedence rule works exactly as above for methods and paths, but we had to make one exception for hosts to preserve compatibility: if two patterns would otherwise conflict and one has a host while the other does not, then the pattern with the host takes precedence.

​	优先级规则对方法和路径的作用与上述完全相同，但我们不得不为主机设置一个例外以保持兼容性：如果两个模式本来会发生冲突，并且一个有主机而另一个没有，那么具有主机的模式优先。

Students of computer science may recall the beautiful theory of regular expressions and regular languages. Each regular expression picks out a regular language, the set of strings matched by the expression. Some questions are easier to pose and answer by talking about languages rather than expressions. Our precedence rule was inspired by this theory. Indeed, each routing pattern corresponds to a regular expression, and sets of matching requests play the role of regular languages.

​	计算机科学的学生可能记得正则表达式和正则语言的优美理论。每个正则表达式都会挑选一个正则语言，即与该表达式匹配的字符串集合。通过讨论语言而不是表达式，有些问题更容易提出和回答。我们的优先级规则受到这一理论的启发。事实上，每个路由模式都对应一个正则表达式，而匹配请求的集合则充当正则语言的角色。

Defining precedence by languages instead of expressions makes it easy to state and understand. But there is a downside to having a rule based on potentially infinite sets: it isn’t clear how to implement it efficiently. It turns out we can determine whether two patterns conflict by walking them segment by segment. Roughly speaking, if one pattern has a literal segment wherever the other has a wildcard, it is more specific; but if literals align with wildcards in both directions, the patterns conflict.

​	通过语言而不是表达式来定义优先级，使其易于表述和理解。但是，有一个基于潜在无限集的规则的缺点：不清楚如何有效地实现它。事实证明，我们可以通过逐段遍历模式来确定两个模式是否冲突。粗略地说，如果一个模式在另一个模式具有通配符的任何地方都有一个文字段，那么它就更具体；但是，如果文字与通配符在两个方向上都对齐，则模式冲突。

As new patterns are registered on a `ServeMux`, it checks for conflicts with previously registered patterns. But checking every pair of patterns would take quadratic time. We use an index to skip patterns that cannot conflict with a new pattern; in practice, it works quite well. In any case, this check happens when patterns are registered, usually at server startup. The time to match incoming requests in Go 1.22 hasn’t changed much from previous versions.

​	当新模式在 `ServeMux` 上注册时，它会检查与先前注册的模式是否存在冲突。但检查每一对模式将花费二次时间。我们使用索引来跳过与新模式不会冲突的模式；在实践中，它运行得很好。无论如何，此检查在模式注册时发生，通常在服务器启动时发生。Go 1.22 中匹配传入请求的时间与以前版本相比没有太大变化。

## Compatibility 兼容性

We made every effort to keep the new functionality compatible with older versions of Go. The new pattern syntax is a superset of the old, and the new precedence rule generalizes the old one. But there are a few edge cases. For example, previous versions of Go accepted patterns with braces and treated them literally, but Go 1.22 uses braces for wildcards. The GODEBUG setting `httpmuxgo121` restores the old behavior.

​	我们尽一切努力使新功能与旧版本的 Go 兼容。新模式语法是旧语法的超集，新优先级规则概括了旧规则。但有一些边缘情况。例如，以前版本的 Go 接受带有大括号的模式并按字面意思对待它们，但 Go 1.22 将大括号用于通配符。GODEBUG 设置 `httpmuxgo121` 恢复了旧行为。

For more details about these routing enhancements, see the [`net/http.ServeMux` documentation](https://go.dev/pkg/net/http#ServeMux).

​	有关这些路由增强的更多详细信息，请参阅 `net/http.ServeMux` 文档。