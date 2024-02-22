+++
title = "html"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/html@go1.21.3](https://pkg.go.dev/html@go1.21.3)

Package html provides functions for escaping and unescaping HTML text.

​	`html`包提供了用于转义和反转义HTML文本的函数。


## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func EscapeString 

``` go 
func EscapeString(s string) string
```

EscapeString escapes special characters like "`<`" to become "`&lt;`". It escapes only five such characters: `<`, `>`, `&`, `'` and `"`. `UnescapeString(EscapeString(s)) == s` always holds, but the converse isn't always true.

​	EscapeString用于转义HTML文本中的特殊字符，例如将"`<`"转义为"`&lt;`"。它只转义五个特殊字符： `<`, `>`, `&`, `'` 和 `"`。`UnescapeString(EscapeString(s)) == s` 总是成立，但反过来不一定成立。

#### EscapeString Example
``` go 
package main

import (
	"fmt"
	"html"
)

func main() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
}

Output:

&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
```

### func UnescapeString 

``` go 
func UnescapeString(s string) string
```

UnescapeString unescapes entities like "`&lt;`" to become "`<`". It unescapes a larger range of entities than EscapeString escapes. For example, "`&aacute;`" unescapes to "`á`", as does "`&#225;`" and "`&#xE1;`". UnescapeString(EscapeString(s)) == s always holds, but the converse isn't always true.

​	UnescapeString用于反转义实体，例如将"`&lt;`"反转义为"`<`"。它比EscapeString转义更多的实体。例如，"`&aacute;`"反转义为"`á`"，"`&#225;`"和"`&#xE1;`"也是如此。`UnescapeString(EscapeString(s)) == s` 总是成立，但反过来不一定成立。

#### UnescapeString Example
``` go 
package main

import (
	"fmt"
	"html"
)

func main() {
	const s = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(s))
}

Output:

"Fran & Freddie's Diner" <tasty@example.com>
```

## 类型

This section is empty.