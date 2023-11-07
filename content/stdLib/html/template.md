+++
title = "template"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/html/template@go1.21.3



![Alert](template_img/alert_gm_grey_24dp.svg) [GO-2023-1703](https://pkg.go.dev/vuln/GO-2023-1703): Templates do not properly consider backticks (`) as Javascript string delimiters, and do not escape them as expected. Backticks are used, since ES6, for JS template literals. If a template contains a Go template action within a Javascript template literal, the contents of the action can be used to terminate the literal, injecting arbitrary Javascript code into the Go template. As ES6 template literals are rather complex, and themselves can do string interpolation, the decision was made to simply disallow Go template actions from being used inside of them (e.g. "var a = {{.}}"), since there is no obviously safe way to allow this behavior. This takes the same approach as github.com/google/safehtml. With fix, Template.Parse returns an Error when it encounters templates like this, with an ErrorCode of value 12. This ErrorCode is currently unexported, but will be exported in the release of Go 1.21. Users who rely on the previous behavior can re-enable it using the GODEBUG flag jstmpllitinterp=1, with the caveat that backticks will now be escaped. This should be used with caution.

[GO-2023-1703](https://pkg.go.dev/vuln/GO-2023-1703)：模板没有正确将反引号（`）视为JavaScript字符串分隔符，并且没有按预期进行转义。自ES6以来，反引号用于JS模板字面量。如果模板中包含一个Go模板操作符在JavaScript模板字面量内，操作符的内容可以用于终止字面量，从而将任意JavaScript代码注入到Go模板中。由于ES6模板字面量相当复杂，并且本身可以进行字符串插值，决定简单地禁止在其中使用Go模板操作符（例如 "var a = {{.}}")，因为没有明显安全的方式允许这种行为。这与github.com/google/safehtml采用了相同的方法。修复后，当Template.Parse遇到这样的模板时，会返回一个带有ErrorCode值为12的错误。目前，该ErrorCode是未导出的，但将在Go 1.21版本发布时进行导出。依赖先前行为的用户可以使用GODEBUG标志jstmpllitinterp=1重新启用它，但需注意反引号现在会被转义。请谨慎使用。

![Alert](template_img/alert_gm_grey_24dp.svg) [GO-2023-1751](https://pkg.go.dev/vuln/GO-2023-1751): Angle brackets (<>) are not considered dangerous characters when inserted into CSS contexts. Templates containing multiple actions separated by a '/' character can result in unexpectedly closing the CSS context and allowing for injection of unexpected HTML, if executed with untrusted input.

 [GO-2023-1751](https://pkg.go.dev/vuln/GO-2023-1751)：当插入CSS上下文时，尖括号（<>）不被视为危险字符。如果执行带有不受信任输入的多个操作由'/'字符分隔的模板，可能会意外关闭CSS上下文，并允许注入意外的HTML。

![Alert](template_img/alert_gm_grey_24dp.svg) [GO-2023-1752](https://pkg.go.dev/vuln/GO-2023-1752): Not all valid JavaScript whitespace characters are considered to be whitespace. Templates containing whitespace characters outside of the character set "\t\n\f\r\u0020\u2028\u2029" in JavaScript contexts that also contain actions may not be properly sanitized during execution.

 [GO-2023-1752](https://pkg.go.dev/vuln/GO-2023-1752)：并非所有有效的JavaScript空白字符都被视为空白。在JavaScript上下文中包含操作的模板中的空白字符，如果不在字符集"\t\n\f\r\u0020\u2028\u2029"之外，可能在执行过程中未能得到适当的过滤。

![Alert](template_img/alert_gm_grey_24dp.svg) [GO-2023-1753](https://pkg.go.dev/vuln/GO-2023-1753): Templates containing actions in unquoted HTML attributes (e.g. "attr={{.}}") executed with empty input can result in output with unexpected results when parsed due to HTML normalization rules. This may allow injection of arbitrary attributes into tags.

 [GO-2023-1753](https://pkg.go.dev/vuln/GO-2023-1753)：在未引用的HTML属性中包含操作的模板（例如 "attr={{.}}")，如果使用空输入执行，解析时可能导致输出出现意外结果，原因是HTML规范化规则。这可能允许将任意属性注入到标签中。



 [GO-2023-2043](https://pkg.go.dev/vuln/GO-2023-2043): The html/template package does not apply the proper rules for handling occurrences of "<script", "<!--", and "</script" within JS literals in `<script>` contexts. This may cause the template parser to improperly consider script contexts to be terminated early, causing actions to be improperly escaped. This could be leveraged to perform an XSS attack.

[GO-2023-2043](https://pkg.go.dev/vuln/GO-2023-2043): html/template包不适用于处理在`<script>`上下文中JS文字中出现的"<script"、"<!--"和"</script"的正确规则。这可能会导致模板解析器错误地认为脚本上下文被提前终止，从而导致操作被错误地转义。这可能被利用来执行跨站脚本攻击（XSS攻击）。



Package template (html/template) implements data-driven templates for generating HTML output safe against code injection. It provides the same interface as package text/template and should be used instead of text/template whenever the output is HTML.

​	`template`包（html/template）实现了用于生成安全的、防止代码注入的HTML输出的数据驱动模板。它提供了与text/template包相同的接口，应该在输出为HTML的情况下使用它来替代text/template。

The documentation here focuses on the security features of the package. For information about how to program the templates themselves, see the documentation for text/template.

​	这里的文档重点介绍了该包的安全特性。有关如何编写模板本身的信息，请参阅text/template的文档。

## 简介 Introduction  

This package wraps package text/template so you can share its template API to parse and execute HTML templates safely.

​	该包封装了text/template包，以便您可以共享其模板API来安全地解析和执行HTML模板。

```
tmpl, err := template.New("name").Parse(...)
// Error checking elided
err = tmpl.Execute(out, data)
```

If successful, tmpl will now be injection-safe. Otherwise, err is an error defined in the docs for ErrorCode.

​	如果成功，tmpl现在将是安全的，防止注入。否则，err是在ErrorCode的文档中定义的错误。

HTML templates treat data values as plain text which should be encoded so they can be safely embedded in an HTML document. The escaping is contextual, so actions can appear within JavaScript, CSS, and URI contexts.

​	HTML模板将数据值视为纯文本，应进行编码，以便它们可以安全地嵌入在HTML文档中。转义是有上下文的，因此操作可以出现在JavaScript、CSS和URI的上下文中。

The security model used by this package assumes that template authors are trusted, while Execute's data parameter is not. More details are provided below.

​	该包使用的安全模型假定模板作者是可信任的，而Execute的数据参数是不可信的。下面提供了更多详细信息。

示例

```
import "text/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
```

输出：

```
Hello, <script>alert('you have been pwned')</script>!
```

but the contextual autoescaping in html/template

但是，在html/template的上下文自动转义中：

```
import "html/template"
...
t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
```

produces safe, escaped HTML output

会产生安全的、经过转义的HTML输出：

```
Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!
```

## 上下文 Contexts  

This package understands HTML, CSS, JavaScript, and URIs. It adds sanitizing functions to each simple action pipeline, so given the excerpt

​	该包理解HTML、CSS、JavaScript和URI。它为每个简单操作管道添加了清理函数，因此在给定的摘录中

```
<a href="/search?q={{.}}">{{.}}</a>
```

At parse time each {{.}} is overwritten to add escaping functions as necessary. In this case it becomes

​	在解析时，每个{{.}}都会被覆盖以添加必要的转义函数。在这种情况下，它变成了

```
<a href="/search?q={{. | urlescaper | attrescaper}}">{{. | htmlescaper}}</a>
```

where urlescaper, attrescaper, and htmlescaper are aliases for internal escaping functions.

其中urlescaper、attrescaper和htmlescaper是内部转义函数的别名。

For these internal escaping functions, if an action pipeline evaluates to a nil interface value, it is treated as though it were an empty string.

​	对于这些内部转义函数，如果操作管道评估为nil接口值，则会被视为一个空字符串。

## 命名空间和数据属性 Namespaced and data- attributes  

Attributes with a namespace are treated as if they had no namespace. Given the excerpt

​	具有命名空间的属性被视为没有命名空间。给定摘录

```
<a my:href="{{.}}"></a>
```

At parse time the attribute will be treated as if it were just "href". So at parse time the template becomes:

​	在解析时，该属性将被视为只有 "href"。因此，在解析时，模板变为：

```
<a my:href="{{. | urlescaper | attrescaper}}"></a>
```

Similarly to attributes with namespaces, attributes with a "data-" prefix are treated as if they had no "data-" prefix. So given

​	类似于带有命名空间的属性，具有 "data-" 前缀的属性被视为没有 "data-" 前缀。因此，给定

```
<a data-href="{{.}}"></a>
```

At parse time this becomes

​	在解析时，它变为

```
<a data-href="{{. | urlescaper | attrescaper}}"></a>
```

If an attribute has both a namespace and a "data-" prefix, only the namespace will be removed when determining the context. For example

​	如果一个属性既有命名空间又有 "data-" 前缀，在确定上下文时，只有命名空间将被移除。例如

```
<a my:data-href="{{.}}"></a>
```

This is handled as if "my:data-href" was just "data-href" and not "href" as it would be if the "data-" prefix were to be ignored too. Thus at parse time this becomes just

​	这被处理为如果 "my:data-href" 只是 "data-href"，而不是 "href"，因为如果也忽略 "data-" 前缀的话，它将成为 "href"。因此，在解析时，它只变成了

```
<a my:data-href="{{. | attrescaper}}"></a>
```

As a special case, attributes with the namespace "xmlns" are always treated as containing URLs. Given the excerpts

​	作为特例，具有命名空间 "xmlns" 的属性始终被视为包含URL。给定摘录

```
<a xmlns:title="{{.}}"></a>
<a xmlns:href="{{.}}"></a>
<a xmlns:onclick="{{.}}"></a>
```

At parse time they become:

​	在解析时，它们变成：

```
<a xmlns:title="{{. | urlescaper | attrescaper}}"></a>
<a xmlns:href="{{. | urlescaper | attrescaper}}"></a>
<a xmlns:onclick="{{. | urlescaper | attrescaper}}"></a>
```

## 错误 Errors  

See the documentation of ErrorCode for details.

​	有关详细信息，请参阅 ErrorCode 的文档。

## 更全面的说明 A fuller picture  

The rest of this package comment may be skipped on first reading; it includes details necessary to understand escaping contexts and error messages. Most users will not need to understand these details.

​	在首次阅读时可以跳过包注释的其余部分；它包含了理解转义上下文和错误消息所必需的细节。大多数用户不需要理解这些细节。

## 上下文 Contexts  

Assuming {{.}} is `O'Reilly: How are <i>you</i>?`, the table below shows how {{.}} appears when used in the context to the left.

​	假设 {{.}} 是 `O'Reilly: How are <i>you</i>?`，下表显示了在左侧上下文中使用 {{.}} 时的结果。

```
Context                          {{.}} After
{{.}}                            O'Reilly: How are &lt;i&gt;you&lt;/i&gt;?
<a title='{{.}}'>                O&#39;Reilly: How are you?
<a href="/{{.}}">                O&#39;Reilly: How are %3ci%3eyou%3c/i%3e?
<a href="?q={{.}}">              O&#39;Reilly%3a%20How%20are%3ci%3e...%3f
<a onx='f("{{.}}")'>             O\x27Reilly: How are \x3ci\x3eyou...?
<a onx='f({{.}})'>               "O\x27Reilly: How are \x3ci\x3eyou...?"
<a onx='pattern = /{{.}}/;'>     O\x27Reilly: How are \x3ci\x3eyou...\x3f
```

If used in an unsafe context, then the value might be filtered out:

​	如果在不安全的上下文中使用，则该值可能会被过滤掉：

```
Context                          {{.}} After
<a href="{{.}}">                 #ZgotmplZ
```

since "O'Reilly:" is not an allowed protocol like "http:".

​	因为 "O'Reilly:" 不是像 "http:" 那样被允许的协议。

If {{.}} is the innocuous word, `left`, then it can appear more widely,

​	如果 {{.}} 是无害的单词 `left`，则它可以出现在更广泛的上下文中：

```
Context                              {{.}} After
{{.}}                                left
<a title='{{.}}'>                    left
<a href='{{.}}'>                     left
<a href='/{{.}}'>                    left
<a href='?dir={{.}}'>                left
<a style="border-{{.}}: 4px">        left
<a style="align: {{.}}">             left
<a style="background: '{{.}}'>       left
<a style="background: url('{{.}}')>  left
<style>p.{{.}} {color:red}</style>   left
```

Non-string values can be used in JavaScript contexts. If {{.}} is

非字符串值可以在 JavaScript 上下文中使用。如果 {{.}} 是

```
struct{A,B string}{ "foo", "bar" }
```

in the escaped template

在转义模板中

```
<script>var pair = {{.}};</script>
```

then the template output is

那么模板的输出将是

```
<script>var pair = {"A": "foo", "B": "bar"};</script>
```

See package json to understand how non-string content is marshaled for embedding in JavaScript contexts.

​	请参阅 json 包以了解如何将非字符串内容编组以嵌入到 JavaScript 上下文中。

## 类型化字符串 typed Strings 

By default, this package assumes that all pipelines produce a plain text string. It adds escaping pipeline stages necessary to correctly and safely embed that plain text string in the appropriate context.

​	默认情况下，该包假设所有管道生成的是纯文本字符串。它会添加必要的转义管道阶段，以正确而安全地

When a data value is not plain text, you can make sure it is not over-escaped by marking it with its type.

​	当数据值不是纯文本时，您可以通过标记其类型来确保它不会被过度转义。

Types HTML, JS, URL, and others from content.go can carry safe content that is exempted from escaping.

​	类型 HTML、JS、URL 和来自 content.go 的其他类型可以包含免于转义的安全内容。

The template 模板

```
Hello, {{.}}!
```

can be invoked with

可以使用以下方式调用：

```
tmpl.Execute(out, template.HTML(`<b>World</b>`))
```

to produce

生成结果为

```
Hello, <b>World</b>!
```

如果 {{.}} 是普通字符串则会生成

```
Hello, &lt;b&gt;World&lt;b&gt;!
```

that would have been produced if {{.}} was a regular string.



## 安全模型 Security Model 

https://rawgit.com/mikesamuel/sanitized-jquery-templates/trunk/safetemplate.html#problem_definition defines "safe" as used by this package.

https://rawgit.com/mikesamuel/sanitized-jquery-templates/trunk/safetemplate.html#problem_definition 定义了	此包使用的"安全"概念。

This package assumes that template authors are trusted, that Execute's data parameter is not, and seeks to preserve the properties below in the face of untrusted data:

​	该包假设模板作者是可信任的，而 Execute 的数据参数不可信，并努力在面对不可信数据时保持以下属性：

Structure Preservation Property: "... when a template author writes an HTML tag in a safe templating language, the browser will interpret the corresponding portion of the output as a tag regardless of the values of untrusted data, and similarly for other structures such as attribute boundaries and JS and CSS string boundaries."

​	结构保留属性："……当模板作者在安全的模板语言中编写一个 HTML 标签时，无论不可信数据的值如何，浏览器都会将输出的相应部分解释为标签，对于属性边界、JS 和 CSS 字符串边界等其他结构也是如此。"

Code Effect Property: "... only code specified by the template author should run as a result of injecting the template output into a page and all code specified by the template author should run as a result of the same."

​	代码效果属性："……只有模板作者指定的代码应该在将模板输出注入页面后运行，而且所有由模板作者指定的代码都应作为结果运行。"

Least Surprise Property: "A developer (or code reviewer) familiar with HTML, CSS, and JavaScript, who knows that contextual autoescaping happens should be able to look at a {{.}} and correctly infer what sanitization happens."

​	最小惊奇属性："熟悉 HTML、CSS 和 JavaScript 的开发人员（或代码审查人员）应该能够了解到上下文自动转义的存在，并能正确推断出进行了哪些净化。"

## Example
``` go 
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(os.Stdout, noItems)
	check(err)

}
Output:

<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>My page</title>
	</head>
	<body>
		<div>My photos</div><div>My blog</div>
	</body>
</html>
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>My another page</title>
	</head>
	<body>
		<div><strong>no rows</strong></div>
	</body>
</html>
```

## Example (Autoescaping)
``` go 
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	check(err)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	check(err)
}

Output:

Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!
```

## Example (Escape) 
``` go 
package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	v := []any{`"Fran & Freddie's Diner"`, ' ', `<tasty@example.com>`}

	fmt.Println(template.HTMLEscapeString(s))
	template.HTMLEscape(os.Stdout, []byte(s))
	fmt.Fprintln(os.Stdout, "")
	fmt.Println(template.HTMLEscaper(v...))

	fmt.Println(template.JSEscapeString(s))
	template.JSEscape(os.Stdout, []byte(s))
	fmt.Fprintln(os.Stdout, "")
	fmt.Println(template.JSEscaper(v...))

	fmt.Println(template.URLQueryEscaper(v...))

}
Output:

&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
&#34;Fran &amp; Freddie&#39;s Diner&#34;32&lt;tasty@example.com&gt;
\"Fran \u0026 Freddie\'s Diner\" \u003Ctasty@example.com\u003E
\"Fran \u0026 Freddie\'s Diner\" \u003Ctasty@example.com\u003E
\"Fran \u0026 Freddie\'s Diner\"32\u003Ctasty@example.com\u003E
%22Fran+%26+Freddie%27s+Diner%2232%3Ctasty%40example.com%3E
```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func HTMLEscape 

``` go 
func HTMLEscape(w io.Writer, b []byte)
```

HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.

​	HTMLEscape将纯文本数据b的转义HTML写入w。

### func HTMLEscapeString 

``` go 
func HTMLEscapeString(s string) string
```

HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.

​	HTMLEscapeString返回纯文本数据s的转义HTML。

### func HTMLEscaper 

``` go 
func HTMLEscaper(args ...any) string
```

HTMLEscaper returns the escaped HTML equivalent of the textual representation of its arguments.

​	HTMLEscaper返回其参数的文本表示的转义HTML。

### func IsTrue  <- go1.6

``` go 
func IsTrue(val any) (truth, ok bool)
```

IsTrue reports whether the value is 'true', in the sense of not the zero of its type, and whether the value has a meaningful truth value. This is the definition of truth used by if and other such actions.

​	IsTrue报告值是否为"true"，即不为其类型的零值，并且值具有有意义的真值。这是if和其他类似操作使用的真值定义。

### func JSEscape 

``` go 
func JSEscape(w io.Writer, b []byte)
```

JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.

​	JSEscape将纯文本数据b的转义JavaScript写入w。

### func JSEscapeString 

``` go 
func JSEscapeString(s string) string
```

JSEscapeString returns the escaped JavaScript equivalent of the plain text data s.

​	JSEscapeString返回纯文本数据s的转义JavaScript。

### func JSEscaper 

``` go 
func JSEscaper(args ...any) string
```

JSEscaper returns the escaped JavaScript equivalent of the textual representation of its arguments.

​	JSEscaper返回其参数的文本表示的转义JavaScript。

### func URLQueryEscaper 

``` go 
func URLQueryEscaper(args ...any) string
```

URLQueryEscaper returns the escaped value of the textual representation of its arguments in a form suitable for embedding in a URL query.

​	URLQueryEscaper以适合嵌入URL查询中的形式返回其参数的文本表示的转义值。

## 类型

### type CSS 

``` go 
type CSS string
```

CSS encapsulates known safe content that matches any of:

​	CSS封装了与以下任何内容匹配的已知安全内容： 

1. The CSS3 stylesheet production, such as `p { color: purple }`.
2. The CSS3 rule production, such as `a[href=~"https:"].foo#bar`.
3. CSS3 declaration productions, such as `color: red; margin: 2px`.
4. The CSS3 value production, such as `rgba(0, 0, 255, 127)`.
5. CSS3样式表产生，例如`p { color: purple }`。
6. CSS3规则产生，例如`a[href=~"https:"].foo#bar`。
7. CSS3声明产生，例如`color: red; margin: 2px`。
8. CSS3值产生，例如`rgba(0, 0, 255, 127)`。

See https://www.w3.org/TR/css3-syntax/#parsing and https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax#style

​	参见 https://www.w3.org/TR/css3-syntax/#parsing 和 https://web.archive.org/web/20090211114933/http://w3.org/TR/css3-syntax#style

Use of this type presents a security risk: the encapsulated content should come from a trusted source, as it will be included verbatim in the template output.

​	使用此类型存在安全风险：封装的内容应来自可信源，因为它将原样包含在模板输出中。

### type Error 

``` go 
type Error struct {
	// ErrorCode describes the kind of error.
    // ErrorCode描述错误的类型。
	ErrorCode ErrorCode
	// Node is the node that caused the problem, if known.
	// If not nil, it overrides Name and Line.
    // Node是导致问题的节点（如果已知）。
	// 如果非nil，则覆盖Name和Line。
	Node parse.Node
	// Name is the name of the template in which the error was encountered.
    // Name是遇到错误的模板的名称。
	Name string
	// Line is the line number of the error in the template source or 0.
    // Line是模板源中错误的行号，如果为0则表示未知。
	Line int
	// Description is a human-readable description of the problem.
    // Description是问题的人类可读描述。
	Description string
}
```

Error describes a problem encountered during template Escaping.

​	Error描述在模板转义过程中遇到的问题。

#### (*Error) Error 

``` go 
func (e *Error) Error() string
```

### type ErrorCode 

``` go 
type ErrorCode int
```

ErrorCode is a code for a kind of error.

​	ErrorCode是一种错误类型的代码。

``` go 
const (
	// OK indicates the lack of an error. OK表示没有错误。    
	OK ErrorCode = iota

	// ErrAmbigContext: "... appears in an ambiguous context within a URL"
	// Example:
	//   <a href="
	//      {{if .C}}
	//        /path/
	//      {{else}}
	//        /search?q=
	//      {{end}}
	//      {{.X}}
	//   ">
	// Discussion:
	//   {{.X}} is in an ambiguous URL context since, depending on {{.C}},
	//  it may be either a URL suffix or a query parameter.
	//   Moving {{.X}} into the condition removes the ambiguity:
	//   <a href="{{if .C}}/path/{{.X}}{{else}}/search?q={{.X}}">
    // ErrAmbigContext: "... 在 URL 的模糊上下文中出现"
	// 示例：
	//   <a href="
	//      {{if .C}}
	//        /path/
	//      {{else}}
	//        /search?q=
	//      {{end}}
	//      {{.X}}
	//   ">
	// 讨论：
	//   {{.X}} 处于模糊的 URL 上下文中，因为根据 {{.C}} 的值，它可能是 URL 后缀或查询参数。
	//   将 {{.X}} 移入条件语句中可以消除模糊性：
	//   <a href="{{if .C}}/path/{{.X}}{{else}}/search?q={{.X}}">
	ErrAmbigContext

	// ErrBadHTML: "expected space, attr name, or end of tag, but got ...",
	//   "... in unquoted attr", "... in attribute name"
	// Example:
	//   <a href = /search?q=foo>
	//   <href=foo>
	//   <form na<e=...>
	//   <option selected<
	// Discussion:
	//   This is often due to a typo in an HTML element, but some runes
	//   are banned in tag names, attribute names, and unquoted attribute
	//   values because they can tickle parser ambiguities.
	//   Quoting all attributes is the best policy.
    // ErrBadHTML: "预期空格、属性名称或标签结束符，但得到 ..."
	//   "... 在未引用的属性中", "... 在属性名称中"
	// 示例：
	//   <a href = /search?q=foo>
	//   <href=foo>
	//   <form na<e=...>
	//   <option selected<
	// 讨论：
	//   这通常是由于 HTML 元素中的拼写错误引起的，但某些符文在标签名称、属性名称和未引用的属性值中是被禁止的，因为它们可能引发解析器的歧义。
	//   引用所有属性是最佳策略。
	ErrBadHTML

	// ErrBranchEnd: "{{if}} branches end in different contexts"
	// Example:
	//   {{if .C}}<a href="{{end}}{{.X}}
	// Discussion:
	//   Package html/template statically examines each path through an
	//   {{if}}, {{range}}, or {{with}} to escape any following pipelines.
	//   The example is ambiguous since {{.X}} might be an HTML text node,
	//   or a URL prefix in an HTML attribute. The context of {{.X}} is
	//   used to figure out how to escape it, but that context depends on
	//   the run-time value of {{.C}} which is not statically known.
	//
	//   The problem is usually something like missing quotes or angle
	//   brackets, or can be avoided by refactoring to put the two contexts
	//   into different branches of an if, range or with. If the problem
	//   is in a {{range}} over a collection that should never be empty,
	//   adding a dummy {{else}} can help.
    // ErrBranchEnd: "{{if}} 的分支在不同的上下文中结束"
	// 示例：
	//   {{if .C}}<a href="{{end}}{{.X}}
	// 讨论：
	//   html/template 包静态地检查每个 {{if}}、{{range}} 或 {{with}} 中的路径，以转义任何后续的管道。
	//   该示例是模糊的，因为 {{.X}} 可能是 HTML 文本节点，也可能是 HTML 属性中的 URL 前缀。{{.X}} 的上下文用于确定如何进行转义，但该上下文取决于 {{.C}} 的运行时值，该值在静态上是未知的。
	//
	//   问题通常是缺少引号或尖括号，或者可以通过将两个上下文放入 if、range 或 with 的不同分支中来避免。如果问题出现在对不应为空的集合进行的 {{range}} 中，可以添加一个虚拟的 {{else}} 来帮助解决问题。
	ErrBranchEnd

	// ErrEndContext: "... ends in a non-text context: ..."
	// Examples:
	//   <div
	//   <div title="no close quote>
	//   <script>f()
	// Discussion:
	//   Executed templates should produce a DocumentFragment of HTML.
	//   Templates that end without closing tags will trigger this error.
	//   Templates that should not be used in an HTML context or that
	//   produce incomplete Fragments should not be executed directly.
	//
	//   {{define "main"}} <script>{{template "helper"}}</script> {{end}}
	//   {{define "helper"}} document.write(' <div title=" ') {{end}}
	//
	//   "helper" does not produce a valid document fragment, so should
	//   not be Executed directly.
    // ErrEndContext: "... 结束于非文本环境：..."
    // 示例：
    // <div
    // <div title="no close quote>
    // <script>f()
    // 讨论：
    // 执行的模板应该产生一个 HTML 的 DocumentFragment。
    // 如果模板没有关闭标签就结束，将触发此错误。
    // 不应直接执行不适用于 HTML 上下文或产生不完整的 Fragments 的模板。
    //
    // {{define "main"}} <script>{{template "helper"}}</script> {{end}}
    // {{define "helper"}} document.write(' <div title=" ') {{end}}
    //
    // "helper" 不会生成有效的 document fragment，因此不应直接执行。
	ErrEndContext

	// ErrNoSuchTemplate: "no such template ..."
	// Examples:
	//   {{define "main"}}<div {{template "attrs"}}>{{end}}
	//   {{define "attrs"}}href="{{.URL}}"{{end}}
	// Discussion:
	//   Package html/template looks through template calls to compute the
	//   context.
	//   Here the {{.URL}} in "attrs" must be treated as a URL when called
	//   from "main", but you will get this error if "attrs" is not defined
	//   when "main" is parsed.
    // ErrNoSuchTemplate: "没有找到模板..."
    // 示例：
    // {{define "main"}}<div {{template "attrs"}}>{{end}}
    // {{define "attrs"}}href="{{.URL}}"{{end}}
    // 讨论：
    // 包 html/template 会检查模板调用以计算上下文。
    // 这里在从 "main" 调用时，"attrs" 中的 {{.URL}} 必须被视为一个 URL，
    // 但如果 "main" 在解析时 "attrs" 未定义，就会出现该错误。
	ErrNoSuchTemplate

	// ErrOutputContext: "cannot compute output context for template ..."
	// Examples:
	//   {{define "t"}}{{if .T}}{{template "t" .T}}{{end}}{{.H}}",{{end}}
	// Discussion:
	//   A recursive template does not end in the same context in which it
	//   starts, and a reliable output context cannot be computed.
	//   Look for typos in the named template.
	//   If the template should not be called in the named start context,
	//   look for calls to that template in unexpected contexts.
	//   Maybe refactor recursive templates to not be recursive.
    // ErrOutputContext: "无法计算模板的输出上下文..."
    // 示例：
    // {{define "t"}}{{if .T}}{{template "t" .T}}{{end}}{{.H}}",{{end}}
    // 讨论：
    // 递归模板结束的上下文与开始的上下文不一致，无法可靠地计算输出上下文。
    // 检查命名模板中的拼写错误。
    // 如果命名的开始上下文中不应调用该模板，在意外的上下文中查找对该模板的调用。
    // 可能需要重构递归模板，使其不再是递归的。
	ErrOutputContext

	// ErrPartialCharset: "unfinished JS regexp charset in ..."
	// Example:
	//     <script>var pattern = /foo[{{.Chars}}]/</script>
	// Discussion:
	//   Package html/template does not support interpolation into regular
	//   expression literal character sets.
    // ErrPartialCharset: "未完成的 JS 正则表达式字符集：..."
    // 示例：
    // <script>var pattern = /foo[{{.Chars}}]/</script>
    // 讨论：
    // 包 html/template 不支持在正则表达式字面量字符集中进行插值。
	ErrPartialCharset

	// ErrPartialEscape: "unfinished escape sequence in ..."
	// Example:
	//   <script>alert("\{{.X}}")</script>
	// Discussion:
	//   Package html/template does not support actions following a
	//   backslash.
	//   This is usually an error and there are better solutions; for
	//   example
	//     <script>alert("{{.X}}")</script>
	//   should work, and if {{.X}} is a partial escape sequence such as
	//   "xA0", mark the whole sequence as safe content: JSStr(`\xA0`)
    // ErrPartialEscape: "未完成的转义序列：..."
    // 示例：
    // <script>alert("{{.X}}")</script>
    // 讨论：
    // 包 html/template 不支持在反斜杠后面跟随动作。
    // 这通常是一个错误，并且有更好的解决方案；例如
    // <script>alert("{{.X}}")</script>
    // 应该可以工作，如果 {{.X}} 是一个未完成的转义序列，比如 "xA0"，将整个序列标记为安全内容：JSStr(\xA0)
	ErrPartialEscape

	// ErrRangeLoopReentry: "on range loop re-entry: ..."
	// Example:
	//   <script>var x = [{{range .}}'{{.}},{{end}}]</script>
	// Discussion:
	//   If an iteration through a range would cause it to end in a
	//   different context than an earlier pass, there is no single context.
	//   In the example, there is missing a quote, so it is not clear
	//   whether {{.}} is meant to be inside a JS string or in a JS value
	//   context. The second iteration would produce something like
	//
	//     <script>var x = ['firstValue,'secondValue]</script>
    // ErrRangeLoopReentry: "在 range 循环重新进入时：..."
    // 示例：
    // <script>var x = [{{range .}}'{{.}},{{end}}]</script>
    // 讨论：
    // 如果通过 range 迭代会导致在之前的迭代中结束于不同的上下文，那么就没有单一的上下文。
    // 在示例中，缺少一个引号，因此不清楚 {{.}} 是应该在 JS 字符串内部还是在 JS 值上下文中。
    // 第二次迭代会产生类似下面的结果：
    //
    // <script>var x = ['firstValue,'secondValue]</script>
	ErrRangeLoopReentry

	// ErrSlashAmbig: '/' could start a division or regexp.
	// Example:
	//   <script>
	//     {{if .C}}var x = 1{{end}}
	//     /-{{.N}}/i.test(x) ? doThis : doThat();
	//   </script>
	// Discussion:
	//   The example above could produce `var x = 1/-2/i.test(s)...`
	//   in which the first '/' is a mathematical division operator or it
	//   could produce `/-2/i.test(s)` in which the first '/' starts a
	//   regexp literal.
	//   Look for missing semicolons inside branches, and maybe add
	//   parentheses to make it clear which interpretation you intend.
    // ErrSlashAmbig: '/' 可能开始一个除法或正则表达式。
    // 示例：
    // <script>
    // {{if .C}}var x = 1{{end}}
    // /-{{.N}}/i.test(x) ? doThis : doThat();
    // </script>
    // 讨论：
    // 上面的示例可能会产生 var x = 1/-2/i.test(s)...，
    // 其中第一个 '/' 是一个数学除法运算符，或者可能会产生 /-2/i.test(s)，
    // 其中第一个 '/' 开始一个正则表达式字面量。
    // 检查分支中是否缺少分号，并可能添加括号以明确您想要的解释方式。
	ErrSlashAmbig

	// ErrPredefinedEscaper: "predefined escaper ... disallowed in template"
	// Example:
	//   <div class={{. | html}}>Hello<div>
	// Discussion:
	//   Package html/template already contextually escapes all pipelines to
	//   produce HTML output safe against code injection. Manually escaping
	//   pipeline output using the predefined escapers "html" or "urlquery" is
	//   unnecessary, and may affect the correctness or safety of the escaped
	//   pipeline output in Go 1.8 and earlier.
	//
	//   In most cases, such as the given example, this error can be resolved by
	//   simply removing the predefined escaper from the pipeline and letting the
	//   contextual autoescaper handle the escaping of the pipeline. In other
	//   instances, where the predefined escaper occurs in the middle of a
	//   pipeline where subsequent commands expect escaped input, e.g.
	//     {{.X | html | makeALink}}
	//   where makeALink does
	//     return `<a href="`+input+`">link</a>`
	//   consider refactoring the surrounding template to make use of the
	//   contextual autoescaper, i.e.
	//     <a href="{{.X}}">link</a>
	//
	//   To ease migration to Go 1.9 and beyond, "html" and "urlquery" will
	//   continue to be allowed as the last command in a pipeline. However, if the
	//   pipeline occurs in an unquoted attribute value context, "html" is
	//   disallowed. Avoid using "html" and "urlquery" entirely in new templates.
    // ErrPredefinedEscaper: "预定义的转义器... 不允许在模板中使用"
    // 示例：
    // <div class={{. | html}}>Hello<div>
    // 讨论：
    // 包 html/template 已经根据上下文对所有 pipeline 进行了转义，以生成防止代码注入的安全 HTML 输出。
    // 手动使用预定义的转义器 "html" 或 "urlquery" 转义 pipeline 输出是不必要的，并且可能影响 Go 1.8 及更早版本中转义 pipeline 输出的正确性和安全性。
    //
    // 在大多数情况下，例如给定的示例，可以通过从 pipeline 中简单地移除预定义的转义器，
    // 让上下文自动转义器处理 pipeline 的转义来解决此错误。
    // 在其他情况下，如果预定义的转义器出现在 pipeline 的中间位置，后续的命令需要转义输入，例如：
    // {{.X | html | makeALink}}
    // 其中 makeALink 函数返回
    // return <a href="+input+">link</a>
    // 考虑重新设计周围的模板，以利用上下文自动转义器，即：
    // <a href="{{.X}}">link</a>
    //
    // 为了便于迁移到 Go 1.9 及更高版本，"html" 和 "urlquery" 仍将被允许作为 pipeline 中的最后一个命令。
    // 但是，如果 pipeline 出现在未引用的属性值上下文中，将禁止使用 "html"。
    // 在新的模板中完全避免使用 "html" 和 "urlquery"。
	ErrPredefinedEscaper
)
```

We define codes for each error that manifests while escaping templates, but escaped templates may also fail at runtime.

​	我们为每个在模板转义过程中出现的错误定义了错误代码，但转义后的模板也可能在运行时失败。

Output: "ZgotmplZ" Example:

​	输出："ZgotmplZ" 示例：

```
<img src="{{.X}}">
where {{.X}} evaluates to `javascript:...`
```

Discussion:

​	讨论：

```
"ZgotmplZ" is a special value that indicates that unsafe content reached a
CSS or URL context at runtime. The output of the example will be
  <img src="#ZgotmplZ">
If the data comes from a trusted source, use content types to exempt it
from filtering: URL(`javascript:...`).
```

### type FuncMap 

``` go 
type FuncMap = template.FuncMap
```

### type HTML 

``` go 
type HTML string
```

HTML encapsulates a known safe HTML document fragment. It should not be used for HTML from a third-party, or HTML with unclosed tags or comments. The outputs of a sound HTML sanitizer and a template escaped by this package are fine for use with HTML.

​	HTML 封装了一个已知安全的 HTML 文档片段。它不应该用于来自第三方的 HTML，或者包含未闭合的标签或注释的 HTML。经过安全的 HTML 清理器处理的输出和由该包转义的模板均可安全用于 HTML。

Use of this type presents a security risk: the encapsulated content should come from a trusted source, as it will be included verbatim in the template output.

​	使用此类型存在安全风险：封装的内容应来自可信任的源，因为它将直接包含在模板输出中。

### type HTMLAttr 

``` go 
type HTMLAttr string
```

HTMLAttr encapsulates an HTML attribute from a trusted source, for example, ` dir="ltr"`.

​	HTMLAttr 封装了来自可信任源的 HTML 属性，例如 `dir="ltr"`。

Use of this type presents a security risk: the encapsulated content should come from a trusted source, as it will be included verbatim in the template output.

​	使用此类型存在安全风险：封装的内容应来自可信任的源，因为它将直接包含在模板输出中。

### type JS 

``` go 
type JS string
```

JS encapsulates a known safe EcmaScript5 Expression, for example, `(x + y * z())`. Template authors are responsible for ensuring that typed expressions do not break the intended precedence and that there is no statement/expression ambiguity as when passing an expression like "{ foo: bar() }\n['foo']()", which is both a valid Expression and a valid Program with a very different meaning.

​	JS 封装了一个已知安全的 EcmaScript5 表达式，例如 `(x + y * z())`。模板作者负责确保类型化的表达式不会破坏预期的优先级，并且没有语句/表达式的歧义，例如传递一个表达式 "{ foo: bar() }\n['foo'](https://chat.openai.com/c/25703472-e130-4214-821f-951eb886f32c)"，它既是一个有效的表达式，也是一个具有完全不同含义的有效程序。

Use of this type presents a security risk: the encapsulated content should come from a trusted source, as it will be included verbatim in the template output.

​	使用此类型存在安全风险：封装的内容应来自可信任的源，因为它将直接包含在模板输出中。

Using JS to include valid but untrusted JSON is not safe. A safe alternative is to parse the JSON with json.Unmarshal and then pass the resultant object into the template, where it will be converted to sanitized JSON when presented in a JavaScript context.

​	使用 JS 来包含有效但不可信任的 JSON 是不安全的。一个安全的替代方法是使用 json.Unmarshal 解析 JSON，然后将结果对象传递到模板中，在 JavaScript 上下文中将其转换为经过清理的 JSON。

### type JSStr 

``` go 
type JSStr string
```

JSStr encapsulates a sequence of characters meant to be embedded between quotes in a JavaScript expression. The string must match a series of StringCharacters:

​	JSStr 封装了一系列字符，用于在 JavaScript 表达式中的引号之间嵌入。字符串必须匹配以下 StringCharacters 的规则：

```
StringCharacter :: SourceCharacter but not `\` or LineTerminator
                 | EscapeSequence
```

Note that LineContinuations are not allowed. JSStr("foo\\nbar") is fine, but JSStr("foo\\\nbar") is not.

​	请注意，不允许使用 LineContinuations。JSStr("foo\nbar") 是可以的，但 JSStr("foo\\nbar") 是不可以的。

Use of this type presents a security risk: the encapsulated content should come from a trusted source, as it will be included verbatim in the template output.

​	使用此类型存在安全风险：封装的内容应来自可信任的源，因为它将直接包含在模板输出中。

### type Srcset  <- go1.10

``` go 
type Srcset string
```

Srcset encapsulates a known safe srcset attribute (see https://w3c.github.io/html/semantics-embedded-content.html#element-attrdef-img-srcset).

​	Srcset 封装了一个已知安全的 srcset 属性（参见 https://w3c.github.io/html/semantics-embedded-content.html#element-attrdef-img-srcset）。

Use of this type presents a security risk: the encapsulated content should come from a trusted source, as it will be included verbatim in the template output.

​	使用此类型存在安全风险：封装的内容应来自可信任的源，因为它将直接包含在模板输出中。

### type Template 

``` go 
type Template struct {

	// The underlying template's parse tree, updated to be HTML-safe.
    // 底层模板的解析树，已更新为安全的 HTML。
	Tree *parse.Tree
	// contains filtered or unexported fields
}
```

Template is a specialized Template from "text/template" that produces a safe HTML document fragment.

Template 是来自 "text/template" 的专门用于生成安全的 HTML 文档片段的模板。

#### Example (Block) 
``` go 
package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

func main() {
	const (
		master  = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
		overlay = `{{define "list"}} {{join . ", "}}{{end}} `
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}
Output:

Names:
- Gamora
- Groot
- Nebula
- Rocket
- Star-Lord
Names: Gamora, Groot, Nebula, Rocket, Star-Lord
```

#### Example (Glob)

Here we demonstrate loading a set of templates from a directory.

``` go 
package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir := createTestDir([]templateFile{
		// T0.tmpl is a plain template file that just invokes T1.
		{"T0.tmpl", `T0 invokes T1: ({{template "T1"}})`},
		// T1.tmpl defines a template, T1 that invokes T2.
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmpl defines a template T2.
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	// Clean up after the test; another quirk of running as an example.
	defer os.RemoveAll(dir)

	// pattern is the glob pattern used to find all the template files.
	pattern := filepath.Join(dir, "*.tmpl")

	// Here starts the example proper.
	// T0.tmpl is the first name matched, so it becomes the starting template,
	// the value returned by ParseGlob.
	tmpl := template.Must(template.ParseGlob(pattern))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
Output:

T0 invokes T1: (T1 invokes T2: (This is T2))
```

#### Example (Helpers)

This example demonstrates one way to share some templates and use them in different contexts. In this variant we add multiple driver templates by hand to an existing bundle of templates.

``` go 
package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir := createTestDir([]templateFile{
		// T1.tmpl defines a template, T1 that invokes T2.
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmpl defines a template T2.
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})
	// Clean up after the test; another quirk of running as an example.
	defer os.RemoveAll(dir)

	// pattern is the glob pattern used to find all the template files.
	pattern := filepath.Join(dir, "*.tmpl")

	// Here starts the example proper.
	// Load the helpers.
	templates := template.Must(template.ParseGlob(pattern))
	// Add one driver template to the bunch; we do this with an explicit template definition.
	_, err := templates.Parse("{{define `driver1`}}Driver 1 calls T1: ({{template `T1`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver1: ", err)
	}
	// Add another driver template.
	_, err = templates.Parse("{{define `driver2`}}Driver 2 calls T2: ({{template `T2`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver2: ", err)
	}
	// We load all the templates before execution. This package does not require
	// that behavior but html/template's escaping does, so it's a good habit.
	err = templates.ExecuteTemplate(os.Stdout, "driver1", nil)
	if err != nil {
		log.Fatalf("driver1 execution: %s", err)
	}
	err = templates.ExecuteTemplate(os.Stdout, "driver2", nil)
	if err != nil {
		log.Fatalf("driver2 execution: %s", err)
	}
}
Output:

Driver 1 calls T1: (T1 invokes T2: (This is T2))
Driver 2 calls T2: (This is T2)
```

#### Example (Parsefiles)

Here we demonstrate loading a set of templates from files in different directories

``` go 
package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	// Here we create different temporary directories and populate them with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir1 := createTestDir([]templateFile{
		// T1.tmpl is a plain template file that just invokes T2.
		{"T1.tmpl", `T1 invokes T2: ({{template "T2"}})`},
	})

	dir2 := createTestDir([]templateFile{
		// T2.tmpl defines a template T2.
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})

	// Clean up after the test; another quirk of running as an example.
	defer func(dirs ...string) {
		for _, dir := range dirs {
			os.RemoveAll(dir)
		}
	}(dir1, dir2)

	// Here starts the example proper.
	// Let's just parse only dir1/T0 and dir2/T2
	paths := []string{
		filepath.Join(dir1, "T1.tmpl"),
		filepath.Join(dir2, "T2.tmpl"),
	}
	tmpl := template.Must(template.ParseFiles(paths...))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
Output:

T1 invokes T2: (This is T2)
```

#### Example (Share) 
``` go 
package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}

func createTestDir(files []templateFile) string {
	dir, err := os.MkdirTemp("", "template")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, file.contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func main() {
	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir := createTestDir([]templateFile{
		// T0.tmpl is a plain template file that just invokes T1.
		{"T0.tmpl", "T0 ({{.}} version) invokes T1: ({{template `T1`}})\n"},
		// T1.tmpl defines a template, T1 that invokes T2. Note T2 is not defined
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
	})
	// Clean up after the test; another quirk of running as an example.
	defer os.RemoveAll(dir)

	// pattern is the glob pattern used to find all the template files.
	pattern := filepath.Join(dir, "*.tmpl")

	// Here starts the example proper.
	// Load the drivers.
	drivers := template.Must(template.ParseGlob(pattern))

	// We must define an implementation of the T2 template. First we clone
	// the drivers, then add a definition of T2 to the template name space.

	// 1. Clone the helper set to create a new name space from which to run them.
	first, err := drivers.Clone()
	if err != nil {
		log.Fatal("cloning helpers: ", err)
	}
	// 2. Define T2, version A, and parse it.
	_, err = first.Parse("{{define `T2`}}T2, version A{{end}}")
	if err != nil {
		log.Fatal("parsing T2: ", err)
	}

	// Now repeat the whole thing, using a different version of T2.
	// 1. Clone the drivers.
	second, err := drivers.Clone()
	if err != nil {
		log.Fatal("cloning drivers: ", err)
	}
	// 2. Define T2, version B, and parse it.
	_, err = second.Parse("{{define `T2`}}T2, version B{{end}}")
	if err != nil {
		log.Fatal("parsing T2: ", err)
	}

	// Execute the templates in the reverse order to verify the
	// first is unaffected by the second.
	err = second.ExecuteTemplate(os.Stdout, "T0.tmpl", "second")
	if err != nil {
		log.Fatalf("second execution: %s", err)
	}
	err = first.ExecuteTemplate(os.Stdout, "T0.tmpl", "first")
	if err != nil {
		log.Fatalf("first: execution: %s", err)
	}

}
Output:

T0 (second version) invokes T1: (T1 invokes T2: (T2, version B))
T0 (first version) invokes T1: (T1 invokes T2: (T2, version A))
```

#### func Must 

``` go 
func Must(t *Template, err error) *Template
```

Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations such as

​	Must 是一个辅助函数，用于包装调用返回 (*Template, error) 的函数，并在错误非空时引发 panic。它适用于变量初始化，例如：

``` go 
var t = template.Must(template.New("name").Parse("html"))
```

#### func New 

``` go 
func New(name string) *Template
```

New allocates a new HTML template with the given name.

​	New 分配一个具有给定名称的新 HTML 模板。

#### func ParseFS  <- go1.16

``` go 
func ParseFS(fs fs.FS, patterns ...string) (*Template, error)
```

ParseFS is like ParseFiles or ParseGlob but reads from the file system fs instead of the host operating system's file system. It accepts a list of glob patterns. (Note that most file names serve as glob patterns matching only themselves.)

​	ParseFS 类似于 ParseFiles 或 ParseGlob，但从文件系统 fs 中读取，而不是主机操作系统的文件系统。它接受一系列的通配符模式。（请注意，大多数文件名本身都用作只匹配自身的通配符模式。）

#### func ParseFiles 

``` go 
func ParseFiles(filenames ...string) (*Template, error)
```

ParseFiles creates a new Template and parses the template definitions from the named files. The returned template's name will have the (base) name and (parsed) contents of the first file. There must be at least one file. If an error occurs, parsing stops and the returned *Template is nil.

​	ParseFiles 创建一个新的模板，并从指定的文件中解析模板定义。返回的模板的名称将具有第一个文件的（基本）名称和（已解析的）内容。必须至少指定一个文件。如果发生错误，解析将停止，并且返回的 *Template 为 nil。

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results. For instance, ParseFiles("a/foo", "b/foo") stores "b/foo" as the template named "foo", while "a/foo" is unavailable.

​	当解析具有相同名称但位于不同目录中的多个文件时，结果将是最后一个被提及的文件。例如，ParseFiles("a/foo", "b/foo") 将使用名为 "foo" 的模板存储为 "b/foo"，而 "a/foo" 将不可用。

#### func ParseGlob 

``` go 
func ParseGlob(pattern string) (*Template, error)
```

ParseGlob creates a new Template and parses the template definitions from the files identified by the pattern. The files are matched according to the semantics of filepath.Match, and the pattern must match at least one file. The returned template will have the (base) name and (parsed) contents of the first file matched by the pattern. ParseGlob is equivalent to calling ParseFiles with the list of files matched by the pattern.

​	ParseGlob 创建一个新的模板，并从与模式匹配的文件中解析模板定义。文件的匹配方式符合 filepath.Match 的语义，且模式必须匹配至少一个文件。返回的模板将具有由模式匹配的第一个文件的（基本）名称和（已解析的）内容。ParseGlob 等效于使用模式匹配的文件列表调用 ParseFiles。

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.

​	当解析具有相同名称但位于不同目录中的多个文件时，结果将是最后一个被提及的文件。

#### (*Template) AddParseTree 

``` go 
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)
```

AddParseTree creates a new template with the name and parse tree and associates it with t.

AddParseTree 使用给定的名称和解析树创建一个新模板，并将其与 t 关联。

It returns an error if t or any associated template has already been executed.

​	如果 t 或任何关联的模板已经执行过，则返回错误。

#### (*Template) Clone 

``` go 
func (t *Template) Clone() (*Template, error)
```

Clone returns a duplicate of the template, including all associated templates. The actual representation is not copied, but the name space of associated templates is, so further calls to Parse in the copy will add templates to the copy but not to the original. Clone can be used to prepare common templates and use them with variant definitions for other templates by adding the variants after the clone is made.

​	Clone 返回模板的副本，包括所有关联的模板。实际表示形式不会被复制，但关联模板的命名空间会被复制，因此在副本上进一步调用 Parse 将向副本添加模板，而不是原始模板。通过在创建副本之后添加变体定义，可以为其他模板准备常见的模板并使用它们。

It returns an error if t has already been executed.

如果 t 已经执行过，则返回错误。

#### (*Template) DefinedTemplates  <- go1.6

``` go 
func (t *Template) DefinedTemplates() string
```

DefinedTemplates returns a string listing the defined templates, prefixed by the string "; defined templates are: ". If there are none, it returns the empty string. Used to generate an error message.

​	DefinedTemplates 返回列出已定义模板的字符串，以字符串 "; defined templates are: " 为前缀。如果没有定义模板，则返回空字符串。用于生成错误消息。

#### (*Template) Delims 

``` go 
func (t *Template) Delims(left, right string) *Template
```

Delims sets the action delimiters to the specified strings, to be used in subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template definitions will inherit the settings. An empty delimiter stands for the corresponding default: {{ or }}. The return value is the template, so calls can be chained.

​	Delims 将动作定界符设置为指定的字符串，以便在后续调用 Parse、ParseFiles 或 ParseGlob 时使用。嵌套的模板定义将继承这些设置。空定界符表示相应的默认值：{{ 或 }}。返回值是模板本身，因此可以进行链式调用。

##### Delims Example
``` go 
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const text = "<<.Greeting>> {{.Name}}"

	data := struct {
		Greeting string
		Name     string
	}{
		Greeting: "Hello",
		Name:     "Joe",
	}

	t := template.Must(template.New("tpl").Delims("<<", ">>").Parse(text))

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}

}

Output:

Hello {{.Name}}
```

#### (*Template) Execute 

``` go 
func (t *Template) Execute(wr io.Writer, data any) error
```

Execute applies a parsed template to the specified data object, writing the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.

​	Execute 将解析的模板应用于指定的数据对象，并将输出写入 wr。如果执行模板或写入其输出时出现错误，执行将停止，但部分结果可能已经被写入输出写入器。模板可以安全地并行执行，尽管如果并行执行共享一个 Writer，则输出可能会交错。

#### (*Template) ExecuteTemplate 

``` go 
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error
```

ExecuteTemplate applies the template associated with t that has the given name to the specified data object and writes the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.

​	ExecuteTemplate 将与 t 关联且具有给定名称的模板应用于指定的数据对象，并将输出写入 wr。如果执行模板或写入其输出时出现错误，执行将停止，但部分结果可能已经被写入输出写入器。模板可以安全地并行执行，尽管如果并行执行共享一个 Writer，则输出可能会交错。

#### (*Template) Funcs 

``` go 
func (t *Template) Funcs(funcMap FuncMap) *Template
```

Funcs adds the elements of the argument map to the template's function map. It must be called before the template is parsed. It panics if a value in the map is not a function with appropriate return type. However, it is legal to overwrite elements of the map. The return value is the template, so calls can be chained.

​	Funcs 将参数映射的元素添加到模板的函数映射中。它必须在解析模板之前调用。如果映射中的值不是具有适当返回类型的函数，它将引发 panic。但是，覆盖映射中的元素是合法的。返回值是模板本身，因此可以进行链式调用。

#### (*Template) Lookup 

``` go 
func (t *Template) Lookup(name string) *Template
```

Lookup returns the template with the given name that is associated with t, or nil if there is no such template.

​	Lookup 返回与 t 关联且具有给定名称的模板，如果没有这样的模板，则返回 nil。

#### (*Template) Name 

``` go 
func (t *Template) Name() string
```

Name returns the name of the template.

​	Name 返回模板的名称。

#### (*Template) New 

``` go 
func (t *Template) New(name string) *Template
```

New allocates a new HTML template associated with the given one and with the same delimiters. The association, which is transitive, allows one template to invoke another with a {{template}} action.

​	New 为给定的模板分配一个新的 HTML 模板，并具有相同的分隔符。这种关联是传递的，允许一个模板使用 {{template}} 动作调用另一个模板。

If a template with the given name already exists, the new HTML template will replace it. The existing template will be reset and disassociated with t.

​	如果具有给定名称的模板已经存在，则新的 HTML 模板将替换它。现有的模板将被重置并与 t 解除关联。

#### (*Template) Option  <- go1.5

``` go 
func (t *Template) Option(opt ...string) *Template
```

Option sets options for the template. Options are described by strings, either a simple string or "key=value". There can be at most one equals sign in an option string. If the option string is unrecognized or otherwise invalid, Option panics.

​	Option 为模板设置选项。选项由字符串描述，可以是简单字符串或者是 "key=value" 形式。选项字符串中最多只能有一个等号。如果选项字符串无法识别或者无效，Option 会引发 panic。

Known options:

已知选项：

missingkey: Control the behavior during execution if a map is indexed with a key that is not present in the map.

​	missingkey: 控制在执行期间，如果对一个不存在于映射中的键进行索引的行为。

```
"missingkey=default" or "missingkey=invalid"
	The default behavior: Do nothing and continue execution.
	If printed, the result of the index operation is the string
	"<no value>".
"missingkey=zero"
	The operation returns the zero value for the map type's element.
"missingkey=error"
	Execution stops immediately with an error.
```

#### (*Template) Parse 

``` go 
func (t *Template) Parse(text string) (*Template, error)
```

Parse parses text as a template body for t. Named template definitions ({{define ...}} or {{block ...}} statements) in text define additional templates associated with t and are removed from the definition of t itself.

​	Parse 将文本解析为模板的主体。文本中的命名模板定义 ({{define ...}} 或者 {{block ...}} 语句) 将定义与 t 关联的其他模板，并从 t 本身的定义中移除。

Templates can be redefined in successive calls to Parse, before the first use of Execute on t or any associated template. A template definition with a body containing only white space and comments is considered empty and will not replace an existing template's body. This allows using Parse to add new named template definitions without overwriting the main template body.

​	可以在对 t 或者任何关联的模板进行第一次 Execute 调用之前的连续调用中重新定义模板。具有只包含空格和注释的主体的模板定义被视为空，并且不会替换现有模板的主体。这样可以使用 Parse 添加新的命名模板定义，而不覆盖主模板的主体。

#### (*Template) ParseFS  <- go1.16

``` go 
func (t *Template) ParseFS(fs fs.FS, patterns ...string) (*Template, error)
```

ParseFS is like ParseFiles or ParseGlob but reads from the file system fs instead of the host operating system's file system. It accepts a list of glob patterns. (Note that most file names serve as glob patterns matching only themselves.)

​	ParseFS 类似于 ParseFiles 或 ParseGlob，但它从文件系统 fs 中读取，而不是主机操作系统的文件系统。它接受一系列的 glob 模式。(注意，大多数文件名只匹配自身，不作为 glob 模式。)

#### (*Template) ParseFiles 

``` go 
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

ParseFiles parses the named files and associates the resulting templates with t. If an error occurs, parsing stops and the returned template is nil; otherwise it is t. There must be at least one file.

​	ParseFiles 解析指定的文件，并将生成的模板与 t 关联。如果发生错误，解析将停止，返回的模板将为 nil；否则为 t。必须至少指定一个文件。

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.

​	当在不同目录中解析具有相同名称的多个文件时，最后一个被提及的文件将是结果。

ParseFiles returns an error if t or any associated template has already been executed.

​	如果 t 或任何关联的模板已经执行过，ParseFiles 将返回一个错误。

#### (*Template) ParseGlob 

``` go 
func (t *Template) ParseGlob(pattern string) (*Template, error)
```

ParseGlob parses the template definitions in the files identified by the pattern and associates the resulting templates with t. The files are matched according to the semantics of filepath.Match, and the pattern must match at least one file. ParseGlob is equivalent to calling t.ParseFiles with the list of files matched by the pattern.

​	ParseGlob 解析由模式指定的文件中的模板定义，并将生成的模板与 t 关联。文件的匹配方式符合 filepath.Match 的语义，模式必须至少匹配一个文件。ParseGlob 等效于使用模式匹配的文件列表调用 t.ParseFiles。

When parsing multiple files with the same name in different directories, the last one mentioned will be the one that results.

​	当在不同目录中解析具有相同名称的多个文件时，最后一个被提及的文件将是结果。

ParseGlob returns an error if t or any associated template has already been executed.

​	如果 t 或任何关联的模板已经执行过，ParseGlob 将返回一个错误。

#### (*Template) Templates 

``` go 
func (t *Template) Templates() []*Template
```

Templates returns a slice of the templates associated with t, including t itself.

​	Templates 返回与 t 关联的模板的切片，包括 t 本身。

### type URL 

``` go 
type URL string
```

URL encapsulates a known safe URL or URL substring (see [RFC 3986](https://rfc-editor.org/rfc/rfc3986.html)). A URL like `javascript:checkThatFormNotEditedBeforeLeavingPage()` from a trusted source should go in the page, but by default dynamic `javascript:` URLs are filtered out since they are a frequently exploited injection vector.

​	URL 封装了已知安全的 URL 或 URL 子串（参见 [RFC 3986](https://rfc-editor.org/rfc/rfc3986.html)）。来自受信任来源的 URL，例如 `javascript:checkThatFormNotEditedBeforeLeavingPage()`，应该包含在页面中，但默认情况下会过滤掉动态的 `javascript:` URL，因为它们经常被利用为注入向量。

Use of this type presents a security risk: the encapsulated content should come from a trusted source, as it will be included verbatim in the template output.

​	使用此类型存在安全风险：封装的内容应来自受信任的来源，因为它将原样包含在模板输出中。