+++
title = "IP 地址"
weight = 80
date = 2023-07-09T21:51:10+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# IP Address - IP 地址

> 原文：[https://echo.labstack.com/docs/ip-address](https://echo.labstack.com/docs/ip-address)

​	IP 地址在 HTTP 中起着基础性的作用；它用于访问控制（access control）、审计（auditing）、基于地理位置的访问分析（geo-based access analysis）等等。Echo 提供了一个方便的方法 [`Context#RealIP()`](https://godoc.org/github.com/labstack/echo#Context) 来获取 IP 地址。

​	然而，从请求中可靠地获取*真实的* IP 地址并不是一件简单的事情，尤其是当您在应用程序之前放置 L7 代理时。在这种情况下，*真实的* IP 需要在 HTTP 层次上从代理传递到您的应用程序，但是您不能盲目地相信 HTTP 标头。否则，您可能会给别人一个欺骗您的机会。**这是一个安全风险！**

​	为了可靠/安全地检索 IP 地址，您必须让您的应用程序了解您的基础架构的整个体系结构。在 Echo 中，可以通过适当地配置 `Echo#IPExtractor` 来实现这一点。本指南将向您展示为什么以及如何进行配置。

> 注意
>
> ​	注意：如果您没有显式设置 `Echo#IPExtractor`，Echo 将回退到旧的行为，这不是一个好的选择。

​	让我们从两个问题开始，了解正确的方向： 

1. 您是否在应用程序之前放置了任何 HTTP（L7）代理？
   - 这包括云解决方案（例如 AWS ALB 或 GCP HTTP LB）和 OSS 解决方案（例如 Nginx、Envoy 或 Istio 入口网关）。
2. 如果是，您的代理使用哪个 HTTP 标头将客户端 IP 传递给应用程序？

## 情况 1. 没有代理

​	如果您没有使用代理（例如：直接面向互联网），您需要（也必须）查看的是网络层的 IP 地址。任何 HTTP 标头都是不可信的，因为客户端可以完全控制要设置的标头。

​	在这种情况下，请使用 `echo.ExtractIPDirect()`：

```go
e.IPExtractor = echo.ExtractIPDirect()
```



## 情况 2. 使用 X-Forwarded-For 标头的代理

​	[`X-Forwared-For` (XFF)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For) 是传递客户端 IP 地址的常用标头。在代理的每个跳点上，它们将请求的 IP 地址附加到标头的末尾。

​	下面的示例图示了此行为：

```text
            ┌──────────┐            ┌──────────┐            ┌──────────┐
───────────>│ Proxy 1  │───────────>│ Proxy 2  │───────────>│ Your app │
            │ (IP: b)  │            │ (IP: c)  │            │          │
            └──────────┘            └──────────┘            └──────────┘

Case 1.
XFF:  ""                    "a"                     "a, b"
                                                    ~~~~~~
Case 2.
XFF:  "x"                   "x, a"                  "x, a, b"
                                                    ~~~~~~~~~
                                                    ↑ What your app will see
```



​	在这种情况下，使用**从右侧读取的第一个不可信（untrustable） IP**。永远不要使用从左侧读取的第一个 IP，因为它可以由客户端进行配置。这里的 "可信（trustable）" 意味着 "您确信 IP 地址属于您的基础架构"。在上面的示例中，如果 `b` 和 `c` 是可信的，那么客户端的 IP 地址对于两种情况都是 `a`，而不会是 `x`。

​	在 Echo 中，请使用 `ExtractIPFromXFFHeader(...TrustOption)`：

```go
e.IPExtractor = echo.ExtractIPFromXFFHeader()
```



​	默认情况下，它信任内部 IP 地址（环回（loopback）、链路本地单播（link-local unicast）、私有使用地址和来自 [RFC6890](https://tools.ietf.org/html/rfc6890)、[RFC4291](https://tools.ietf.org/html/rfc4291) 和 [RFC4193](https://tools.ietf.org/html/rfc4193) 的唯一本地地址）。要控制此行为，可以使用 [`TrustOption`](https://godoc.org/github.com/labstack/echo#TrustOption)s。

E.g.:

```go
e.IPExtractor = echo.ExtractIPFromXFFHeader(
    TrustLinkLocal(false),
    TrustIPRanges(lbIPRange),
)
```



- 参考：[https://godoc.org/github.com/labstack/echo#TrustOption](https://godoc.org/github.com/labstack/echo#TrustOption)

## 情况 3. 使用 X-Real-IP 标头的代理

​	`X-Real-IP` 是另一个传递客户端 IP 地址的 HTTP 标头，但它只携带一个地址，而不像 XFF 那样携带多个地址。

​	如果您的代理设置了此标头，请使用 `ExtractIPFromRealIPHeader(...TrustOption)`：

```go
e.IPExtractor = echo.ExtractIPFromRealIPHeader()
```



​	同样，默认情况下，它信任内部 IP 地址（环回、链路本地单播、私有使用和来自 [RFC6890](https://tools.ietf.org/html/rfc6890)、[RFC4291](https://tools.ietf.org/html/rfc4291) 和 [RFC4193](https://tools.ietf.org/html/rfc4193) 的唯一本地地址）。要控制此行为，可以使用 [`TrustOption`](https://godoc.org/github.com/labstack/echo#TrustOption)。 

- 参考： [https://godoc.org/github.com/labstack/echo#TrustOption](https://godoc.org/github.com/labstack/echo#TrustOption)

> ​	**永远不要忘记**配置最外层的代理（即在基础架构的边缘）**不要通过传入的标头**。否则，存在欺诈的可能性，因为客户端可以控制这一点。

## 关于默认行为

​	在默认行为中，Echo 查看第一个 XFF 标头、X-Real-IP 标头和来自网络层的 IP。

​	正如您可能已经注意到的，在阅读本文之后，这是不好的。之所以这是默认值，只是为了向后兼容性。