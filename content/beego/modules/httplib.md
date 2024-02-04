+++
title = "httplib 模块"
date = 2024-02-04T09:31:30+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/module/httplib/]({{< ref "/beego/modules/httplib" >}})

# Httplib Module - Httplib 模块



## Client Request 客户端请求

Similar to Curl, httplib is used to simulate http requests sent by clients. Similar to jQuery, it supports method chaining. It’s easy to use and it can be installed by:

​	与 Curl 类似，httplib 用于模拟客户端发送的 http 请求。与 jQuery 类似，它支持方法链接。它易于使用，可以通过以下方式安装：

```
go get github.com/beego/beego/v2/client/httplib
```

## Basic Usage 基本用法

Import package: 
​	导入包：

```
import (
	"github.com/beego/beego/v2/client/httplib"
)	
```

Initialize request method and url:

​	初始化请求方法和 url：

```
req := httplib.Get("http://beego.wiki/")
```

Send the request and retrieve the data in the response:

​	发送请求并在响应中检索数据：

```
str, err := req.String()
if err != nil {
	t.Fatal(err)
}
fmt.Println(str)
```

## Method Functions 方法函数

httplib supports these methods:

​	httplib 支持以下方法：

- `Get(url string)`
- `Post(url string)`
- `Put(url string)`
- `Delete(url string)`
- `Head(url string)`

## Debug Output 调试输出

Enable debug information output:

​	启用调试信息输出：

```
req.Debug(true)
```

Then it will output debug information:

​	然后它将输出调试信息：

```
httplib.Get("http://beego.wiki/").Debug(true).Response()

// Output
GET / HTTP/0.0
Host: beego.wiki
User-Agent: beegoServer
```

## HTTPS Request HTTPS 请求

If the requested scheme is https, we need to set the TLS of client:

​	如果请求的方案是 https，我们需要设置客户端的 TLS：

```
req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
```

[Learn more about TLS settings
了解有关 TLS 设置的更多信息](http://gowalker.org/crypto/tls#Config)

## Set Timeout 设置超时

Can set request timeout and data reading timeout by:

​	可以通过以下方式设置请求超时和数据读取超时：

```
req.SetTimeout(connectTimeout, readWriteTimeout)
```

It is a function of request object. So it can be done like this:

​	它是请求对象的一个函数。因此可以像这样完成：

```
httplib.Get("http://beego.wiki/").SetTimeout(100 * time.Second, 30 * time.Second).Response()
```

## Set Request Params 设置请求参数

For Put or Post requests, we may need to send parameters. Parameters can be set in the following manner:

​	对于 Put 或 Post 请求，我们可能需要发送参数。参数可以按以下方式设置：

```
req := httplib.Post("http://beego.wiki/")
req.Param("username","astaxie")
req.Param("password","123456")
```

## Send big data 发送大数据

To simulate file uploading or to send big data, one can use the `Body` function:

​	要模拟文件上传或发送大数据，可以使用 `Body` 函数：

```
req := httplib.Post("http://beego.wiki/")
bt,err:=ioutil.ReadFile("hello.txt")
if err!=nil{
	log.Fatal("read file err:",err)
}
req.Body(bt)
```

## Set header 设置标头

To simulate header values, e.g.:

​	要模拟标头值，例如：

```
Accept-Encoding:gzip,deflate,sdch
Host:beego.wiki
User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36
```

Can use `Header` function:

​	可以使用 `Header` 函数：

```
req := httplib.Post("http://beego.wiki/")
req.Header("Accept-Encoding","gzip,deflate,sdch")
req.Header("Host","beego.wiki")
req.Header("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
```

## Upload file 上传文件

PostFile function requires the first parameter to be the name of form and the second parameter is the filename or filepath you want to send.

​	PostFile 函数要求第一个参数是表单的名称，第二个参数是要发送的文件名或文件路径。

```
b:=httplib.Post("http://beego.wiki/")
b.Param("username","astaxie")
b.Param("password","123456")
b.PostFile("uploadfile1", "httplib.pdf")
b.PostFile("uploadfile2", "httplib.txt")
str, err := b.String()
if err != nil {
    t.Fatal(err)
}
```

## Get Response 获取响应

The settings above are before sending request, how can we get response after request? Here are the ways:

​	以上设置是在发送请求之前，我们如何在请求后获取响应？方法如下：

| Method 方法                      | Type 类型                 | Description 说明                                             |
| -------------------------------- | ------------------------- | ------------------------------------------------------------ |
| `req.Response()`                 | `(*http.Response, error)` | This is a `http.Response` object. You can get data from it. 这是一个 `http.Response` 对象。您可以从中获取数据。 |
| `req.Bytes()`                    | `([]byte, error)`         | Return raw response body. 返回原始响应主体。                 |
| `req.String()`                   | `(string, error)`         | Return raw response body. 返回原始响应主体。                 |
| `req.ToFile(filename string)`    | `error`                   | Save response body into a file. 将响应主体保存到文件中。     |
| `req.ToJSON(result interface{})` | `error`                   | Parse JSON response into the result object. 将 JSON 响应解析为结果对象。 |
| `req.ToXml(result interface{})`  | `error`                   | Parse XML response into the result object. 将 XML 响应解析为结果对象。 |

# Filter 过滤器

In order to support some AOP feature, e.g. logs, tracing, we designed `filter-chain` for httplib.

​	为了支持一些 AOP 功能，例如日志、跟踪，我们为 httplib 设计了 `filter-chain` 。

There are two key interfaces:

​	有两个关键接口：

```go
type FilterChain func(next Filter) Filter

type Filter func(ctx context.Context, req *BeegoHTTPRequest) (*http.Response, error)
```

This is a typical usage of `Filter-Chain` pattern. So you must invoke `next(...)` when you want to implement your own logic.

​	这是 `Filter-Chain` 模式的典型用法。因此，当您想要实现自己的逻辑时，必须调用 `next(...)` 。

Here is an example：

​	这里有一个示例：

```go
func myFilter(next httplib.Filter) httplib.Filter {
	return func(ctx context.Context, req *httplib.BeegoHTTPRequest) (*http.Response, error) {
		r := req.GetRequest()
		logs.Info("hello, here is the filter: ", r.URL)
		// Never forget invoke this. Or the request will not be sent
		return next(ctx, req)
	}
}
```

And we could register this filter as global filter:

​	我们可以将此过滤器注册为全局过滤器：

```go
	httplib.SetDefaultSetting(httplib.BeegoHTTPSettings{

		FilterChains: []httplib.FilterChain{
			myFilter,
		},

		UserAgent:        "beegoServer",
		ConnectTimeout:   60 * time.Second,
		ReadWriteTimeout: 60 * time.Second,
		Gzip:             true,
		DumpBody:         true,
	})
```

Sometimes you only want to use the filter for specific requests:

​	有时您只想将过滤器用于特定请求：

```go
req.AddFilters(myFilter)
```

We provide some filters.

​	我们提供了一些过滤器。

## Prometheus Filter Prometheus 过滤器

It’s used to support `Prometheus` framework to collect metric data.

​	它用于支持 `Prometheus` 框架收集指标数据。

```go
	builder := prometheus.FilterChainBuilder{
		AppName: "My-test",
		ServerName: "User-server-1",
		RunMode: "dev",
	}
	req := httplib.Get("http://beego.wiki/")
	// only work for this request, or using SetDefaultSetting to support all requests
	req.AddFilters(builder.FilterChain)

	resp, err := req.Response()
	if err != nil {
		logs.Error("could not get response: ", err)
	} else {
		logs.Info(resp)
	}
```

If you don’t use Beego’s admin service, you must expose `prometheus` port manually.

​	如果您不使用 Beego 的管理服务，则必须手动公开 `prometheus` 端口。

## Opentracing Filter Opentracing 过滤器

```go
	builder := opentracing.FilterChainBuilder{}
	req := httplib.Get("http://beego.wiki/")
	// only work for this request, or using SetDefaultSetting to support all requests
	req.AddFilters(builder.FilterChain)

	resp, err := req.Response()
	if err != nil {
		logs.Error("could not get response: ", err)
	} else {
		logs.Info(resp)
	}
```

Don’t forget to register `Opentracing` real implementation.

​	别忘了注册 `Opentracing` 真实实现。
