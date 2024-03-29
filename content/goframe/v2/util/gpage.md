+++
title = "gpage"
date = 2024-03-21T17:59:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gpage](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gpage)

Package gpage provides useful paging functionality for web pages.

​	Package gpage 为网页提供了有用的分页功能。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/util/gpage/gpage.go#L38)

```go
const (
	DefaultPageName        = "page"    // DefaultPageName defines the default page name.
	DefaultPagePlaceHolder = "{.page}" // DefaultPagePlaceHolder defines the place holder for the url template.
)
```

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Page

```go
type Page struct {
	TotalSize      int    // Total size.
	TotalPage      int    // Total page, which is automatically calculated.
	CurrentPage    int    // Current page number >= 1.
	UrlTemplate    string // Custom url template for page url producing.
	LinkStyle      string // CSS style name for HTML link tag `a`.
	SpanStyle      string // CSS style name for HTML span tag `span`, which is used for first, current and last page tag.
	SelectStyle    string // CSS style name for HTML select tag `select`.
	NextPageTag    string // Tag name for next p.
	PrevPageTag    string // Tag name for prev p.
	FirstPageTag   string // Tag name for first p.
	LastPageTag    string // Tag name for last p.
	PrevBarTag     string // Tag string for prev bar.
	NextBarTag     string // Tag string for next bar.
	PageBarNum     int    // Page bar number for displaying.
	AjaxActionName string // Ajax function name. Ajax is enabled if this attribute is not empty.
}
```

Page is the pagination implementer. All the attributes are public, you can change them when necessary.

​	Page 是分页实现者。所有属性都是公开的，您可以在必要时更改它们。

#### func New

```go
func New(totalSize, pageSize, currentPage int, urlTemplate string) *Page
```

New creates and returns a pagination manager. Note that the parameter `urlTemplate` specifies the URL producing template, like: /user/list/{.page}, /user/list/{.page}.html, /user/list?page={.page}&type=1, etc. The build-in variable in `urlTemplate` “{.page}” specifies the page number, which will be replaced by certain page number when producing.

​	New 创建并返回分页管理器。请注意，该参数 `urlTemplate` 指定 URL 生成模板，例如：/user/list/{.page}、/user/list/{.page}.html、/user/list？page={.page}&type=1 等。“{.page}”中的 `urlTemplate` 内置变量指定页码，在生成时将替换为特定的页码。

#### (*Page) FirstPage

```go
func (p *Page) FirstPage() string
```

FirstPage returns the HTML content for the first page.

​	FirstPage 返回第一页的 HTML 内容。

#### (*Page) GetContent

```go
func (p *Page) GetContent(mode int) string
```

GetContent returns the page content for predefined mode. These predefined contents are mainly for chinese localization purpose. You can defines your own page function retrieving the page content according to the implementation of this function.

​	GetContent 返回预定义模式的页面内容。这些预定义内容主要用于中文本地化目的。您可以定义自己的页面函数，根据该函数的实现来检索页面内容。

#### (*Page) GetLink

```go
func (p *Page) GetLink(page int, text, title string) string
```

GetLink returns the HTML link tag `a` content for given page number.

​	GetLink 返回给定页码的 HTML 链接标记 `a` 内容。

#### (*Page) GetUrl

```go
func (p *Page) GetUrl(page int) string
```

GetUrl parses the UrlTemplate with given page number and returns the URL string. Note that the UrlTemplate attribute can be either an URL or an URI string with “{.page}” place holder specifying the page number position.

​	GetUrl 分析具有给定页码的 UrlTemplate 并返回 URL 字符串。请注意，UrlTemplate 属性可以是 URL 或 URI 字符串，其中“{.page}”占位符指定页码位置。

#### (*Page) LastPage

```go
func (p *Page) LastPage() string
```

LastPage returns the HTML content for the last page.

​	LastPage 返回最后一页的 HTML 内容。

#### (*Page) NextPage

```go
func (p *Page) NextPage() string
```

NextPage returns the HTML content for the next page.

​	NextPage 返回下一页的 HTML 内容。

#### (*Page) PageBar

```go
func (p *Page) PageBar() string
```

PageBar returns the HTML page bar content with link and span tags.

​	PageBar 返回带有 link 和 span 标记的 HTML 页面栏内容。

#### (*Page) PrevPage

```go
func (p *Page) PrevPage() string
```

PrevPage returns the HTML content for the previous page.

​	PrevPage 返回上一页的 HTML 内容。

#### (*Page) SelectBar

```go
func (p *Page) SelectBar() string
```

SelectBar returns the select HTML content for pagination.

​	SelectBar 返回要分页的选定 HTML 内容。