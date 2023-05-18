+++
title = "html"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# html

https://pkg.go.dev/html@go1.20.1



Package html provides functions for escaping and unescaping HTML text.








## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [EscapeString](https://cs.opensource.google/go/go/+/go1.20.1:src/html/escape.go;l=178) 

``` go 
func EscapeString(s string) string
```

EscapeString escapes special characters like "<" to become "&lt;". It escapes only five such characters: <, >, &, ' and ". UnescapeString(EscapeString(s)) == s always holds, but the converse isn't always true.

##### Example
``` go 
```

#### func [UnescapeString](https://cs.opensource.google/go/go/+/go1.20.1:src/html/escape.go;l=187) 

``` go 
func UnescapeString(s string) string
```

UnescapeString unescapes entities like "&lt;" to become "<". It unescapes a larger range of entities than EscapeString escapes. For example, "&aacute;" unescapes to "á", as does "&#225;" and "&#xE1;". UnescapeString(EscapeString(s)) == s always holds, but the converse isn't always true.

##### Example
``` go 
```

## 类型

This section is empty.