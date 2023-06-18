+++
title = "HTTP/2服务器推送"
weight = 9
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# HTTP/2 Server Push - HTTP/2服务器推送

https://go.dev/blog/h2push

Jaana Burcu Dogan and Tom Bergan
24 March 2017

## Introduction 简介

HTTP/2 is designed to address many of the failings of HTTP/1.x. Modern web pages use many resources: HTML, stylesheets, scripts, images, and so on. In HTTP/1.x, each of these resources must be requested explicitly. This can be a slow process. The browser starts by fetching the HTML, then learns of more resources incrementally as it parses and evaluates the page. Since the server must wait for the browser to make each request, the network is often idle and underutilized.

HTTP/2旨在解决HTTP/1.x的许多缺陷。现代网页使用许多资源。HTML、样式表、脚本、图像等等。在HTTP/1.x中，这些资源中的每一个都必须被明确地请求。这可能是一个缓慢的过程。浏览器从获取HTML开始，然后在解析和评估页面的过程中逐步了解更多的资源。由于服务器必须等待浏览器发出每个请求，因此网络经常处于空闲状态，利用率低。

To improve latency, HTTP/2 introduced *server push*, which allows the server to push resources to the browser before they are explicitly requested. A server often knows many of the additional resources a page will need and can start pushing those resources as it responds to the initial request. This allows the server to fully utilize an otherwise idle network and improve page load times.

为了改善延迟，HTTP/2引入了服务器推送，允许服务器在浏览器明确请求之前向其推送资源。服务器通常知道一个页面将需要的许多额外资源，并可以在响应初始请求时开始推送这些资源。这使得服务器能够充分利用原本空闲的网络，并改善页面加载时间。

![img](HTTP2ServerPush_img/serverpush.svg)

At the protocol level, HTTP/2 server push is driven by `PUSH_PROMISE` frames. A `PUSH_PROMISE` describes a request that the server predicts the browser will make in the near future. As soon as the browser receives a `PUSH_PROMISE`, it knows that the server will deliver the resource. If the browser later discovers that it needs this resource, it will wait for the push to complete rather than sending a new request. This reduces the time the browser spends waiting on the network.

在协议层面，HTTP/2服务器推送是由PUSH_PROMISE帧驱动的。PUSH_PROMISE描述了一个服务器预测浏览器在不久的将来会提出的请求。一旦浏览器收到PUSH_PROMISE，它就知道服务器将交付资源。如果浏览器后来发现它需要这个资源，它将等待推送完成，而不是发送一个新的请求。这就减少了浏览器在网络上的等待时间。

## Server Push in net/http - net/http中的服务器推送

Go 1.8 introduced support for pushing responses from an [`http.Server`](https://go.dev/pkg/net/http/#Server). This feature is available if the running server is an HTTP/2 server and the incoming connection uses HTTP/2. In any HTTP handler, you can assert if the http.ResponseWriter supports server push by checking if it implements the new [`http.Pusher`](https://go.dev/pkg/net/http/#Pusher) interface.

Go 1.8 引入了对从 http.Server 推送响应的支持。如果运行的服务器是一个HTTP/2服务器，并且传入的连接使用HTTP/2，那么这个功能就可用。在任何HTTP处理程序中，您可以通过检查http.ResponseWriter是否实现了新的http.Pusher接口来断定它是否支持服务器推送。

For example, if the server knows that `app.js` will be required to render the page, the handler can initiate a push if `http.Pusher` is available:

例如，如果服务器知道需要app.js来渲染页面，处理程序可以在http.Pusher可用的情况下启动推送：

```go
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if pusher, ok := w.(http.Pusher); ok {
            // Push is supported.
            if err := pusher.Push("/app.js", nil); err != nil {
                log.Printf("Failed to push: %v", err)
            }
        }
        // ...
    })
```

The Push call creates a synthetic request for `/app.js`, synthesizes that request into a `PUSH_PROMISE` frame, then forwards the synthetic request to the server’s request handler, which will generate the pushed response. The second argument to Push specifies additional headers to include in the `PUSH_PROMISE`. For example, if the response to `/app.js` varies on Accept-Encoding, then the `PUSH_PROMISE` should include an Accept-Encoding value:

Push调用为/app.js创建一个合成请求，将该请求合成为一个PUSH_PROMISE框架，然后将该合成请求转发给服务器的请求处理程序，后者将生成推送的响应。推送的第二个参数指定了在PUSH_PROMISE中包含的额外头信息。例如，如果对/app.js的响应在Accept-Encoding上有变化，那么PUSH_PROMISE应该包括一个Accept-Encoding值：

```go
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if pusher, ok := w.(http.Pusher); ok {
            // Push is supported.
            options := &http.PushOptions{
                Header: http.Header{
                    "Accept-Encoding": r.Header["Accept-Encoding"],
                },
            }
            if err := pusher.Push("/app.js", options); err != nil {
                log.Printf("Failed to push: %v", err)
            }
        }
        // ...
    })
```

A fully working example is available at:

一个完整的工作实例可以在以下网站找到：

```
$ go get golang.org/x/blog/content/h2push/server
```

If you run the server and load [https://localhost:8080](https://localhost:8080/), your browser’s developer tools should show that `app.js` and `style.css` were pushed by the server.

如果您运行服务器并加载https://localhost:8080，您的浏览器的开发者工具应该显示app.js和style.css是由服务器推送的。

![img](HTTP2ServerPush_img/networktimeline.png)

## Start Your Pushes Before You Respond 在您回应之前开始您的推送

It’s a good idea to call the Push method before sending any bytes of the response. Otherwise it is possible to accidentally generate duplicate responses. For example, suppose you write part of an HTML response:

在发送响应的任何字节之前调用推送方法是个好主意。否则就有可能意外地产生重复的响应。例如，假设您写了一个HTML响应的一部分：

```html linenums="1"
<html>
<head>
    <link rel="stylesheet" href="a.css">...
```

Then you call Push("a.css", nil). The browser may parse this fragment of HTML before it receives your PUSH_PROMISE, in which case the browser will send a request for `a.css` in addition to receiving your `PUSH_PROMISE`. Now the server will generate two responses for `a.css`. Calling Push before writing the response avoids this possibility entirely.

然后您调用Push("a.css", nil)。浏览器可能会在收到您的PUSH_PROMISE之前解析这个HTML片段，在这种情况下，浏览器除了收到您的PUSH_PROMISE之外，还会发送一个关于a.css的请求。现在服务器将为a.css生成两个响应。在编写响应之前调用推送，可以完全避免这种可能性。

## When To Use Server Push 何时使用服务器推送

Consider using server push any time your network link is idle. Just finished sending the HTML for your web app? Don’t waste time waiting, start pushing the resources your client will need. Are you inlining resources into your HTML file to reduce latency? Instead of inlining, try pushing. Redirects are another good time to use push because there is almost always a wasted round trip while the client follows the redirect. There are many possible scenarios for using push – we are only getting started.

在您的网络链接处于空闲状态时，可以考虑使用服务器推送。刚刚为您的网络应用发送完HTML？不要浪费时间等待，开始推送您的客户需要的资源。您是否将资源内联到您的HTML文件以减少延迟？与其内联，不如尝试推送。重定向是另一个使用推送的好时机，因为当客户端跟随重定向时，几乎总是会有一个浪费的往返过程。使用推送有许多可能的情况--我们才刚刚开始。

We would be remiss if we did not mention a few caveats. First, you can only push resources your server is authoritative for – this means you cannot push resources that are hosted on third-party servers or CDNs. Second, don’t push resources unless you are confident they are actually needed by the client, otherwise your push wastes bandwidth. A corollary is to avoid pushing resources when it’s likely that the client already has those resources cached. Third, the naive approach of pushing all resources on your page often makes performance worse. When in doubt, measure.

如果我们不提及一些注意事项，那就是失职了。首先，您只能推送您的服务器是权威的资源--这意味着您不能推送托管在第三方服务器或CDN的资源。第二，除非您确信客户确实需要这些资源，否则不要推送，否则您的推送会浪费带宽。一个推论是，当客户端可能已经有这些资源的缓存时，要避免推送资源。第三，推送页面上所有资源的天真做法往往会使性能变差。当有疑问时，请测量。

The following links make for good supplemental reading:

下面的链接是很好的补充阅读：

- [HTTP/2 Push: The Details HTTP/2推送：细节](https://calendar.perfplanet.com/2016/http2-push-the-details/)
- [Innovating with HTTP/2 Server Push 创新的HTTP/2服务器推送](https://www.igvita.com/2013/06/12/innovating-with-http-2.0-server-push/)
- [Cache-Aware Server Push in H2O H2O中的缓存感知服务器推送](https://github.com/h2o/h2o/issues/421)
- [The PRPL Pattern PRPL模式](https://developers.google.com/web/fundamentals/performance/prpl-pattern/)
- [Rules of Thumb for HTTP/2 Push HTTP/2推送的经验法则](https://docs.google.com/document/d/1K0NykTXBbbbTlv60t5MyJvXjqKGsCVNYHyLEXIxYMv0)
- [Server Push in the HTTP/2 spec HTTP/2规范中的服务器推送](https://tools.ietf.org/html/rfc7540#section-8.2)

## Conclusion 总结

With Go 1.8, the standard library provides out-of-the-box support for HTTP/2 Server Push, giving you more flexibility to optimize your web applications.

在Go 1.8中，标准库为HTTP/2服务器推送提供了开箱即用的支持，使您可以更灵活地优化您的Web应用。

Go to our [HTTP/2 Server Push demo](https://http2.golang.org/serverpush) page to see it in action.

去我们的HTTP/2服务器推送演示页面看看它的运行情况吧。
