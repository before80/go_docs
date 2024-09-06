+++
title = "mail"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/mail@go1.23.0](https://pkg.go.dev/net/mail@go1.23.0)

Package mail implements parsing of mail messages.

​	mail 包实现了邮件消息的解析。

For the most part, this package follows the syntax as specified by [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) and extended by [RFC 6532](https://rfc-editor.org/rfc/rfc6532.html). Notable divergences:

​	在大多数情况下，此软件包遵循 RFC 5322 指定并由 RFC 6532 扩展的语法。值得注意的差异：

- Obsolete address formats are not parsed, including addresses with embedded route information.
  不解析过时的地址格式，包括包含嵌入式路由信息的地址。
- The full range of spacing (the CFWS syntax element) is not supported, such as breaking addresses across lines.
  不支持全范围的空格（CFWS 语法元素），例如跨行中断地址。
- No unicode normalization is performed.
  不执行 unicode 规范化。
- The special characters ()[]:;@, are allowed to appear unquoted in names.
  允许特殊字符 ()[]:;@, 在名称中不加引号出现。

## 常量

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=160)

```go
var ErrHeaderNotPresent = errors.New("mail: header not in message")
```

## 函数

### func ParseDate <- go1.8

```go
func ParseDate(date string) (time.Time, error)
```

ParseDate parses an [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) date string.

​	ParseDate 解析 RFC 5322 日期字符串。

## 类型

### type Address

```go
type Address struct {
	Name    string // Proper name; may be empty.
	Address string // user@domain
}
```

Address represents a single mail address. An address such as “Barry Gibbs [bg@example.com](mailto:bg@example.com)” is represented as Address{Name: “Barry Gibbs”, Address: “[bg@example.com](mailto:bg@example.com)”}.

​	Address 表示单个邮件地址。例如，“Barry Gibbs bg@example.com”这样的地址表示为 Address{Name: “Barry Gibbs”, Address: “ bg@example.com”}。

#### func ParseAddress <- go1.1

```go
func ParseAddress(address string) (*Address, error)
```

ParseAddress parses a single [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) address, e.g. “Barry Gibbs [bg@example.com](mailto:bg@example.com)”

​	ParseAddress 解析单个 RFC 5322 地址，例如“Barry Gibbs bg@example.com”

##### ParseAddress Example

```go
package main

import (
	"fmt"
	"log"
	"net/mail"
)

func main() {
	e, err := mail.ParseAddress("Alice <alice@example.com>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e.Name, e.Address)

}
Output:

Alice alice@example.com
```

#### func ParseAddressList <- go1.1

```go
func ParseAddressList(list string) ([]*Address, error)
```

ParseAddressList parses the given string as a list of addresses.

​	ParseAddressList 将给定字符串解析为地址列表。

##### ParseAddressList Example

```go
package main

import (
	"fmt"
	"log"
	"net/mail"
)

func main() {
	const list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
	emails, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range emails {
		fmt.Println(v.Name, v.Address)
	}

}
Output:

Alice alice@example.com
Bob bob@example.com
Eve eve@example.com
```

#### (*Address) String

```go
func (a *Address) String() string
```

String formats the address as a valid [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) address. If the address’s name contains non-ASCII characters the name will be rendered according to [RFC 2047](https://rfc-editor.org/rfc/rfc2047.html).

​	String 将地址格式化为有效的 RFC 5322 地址。如果地址的名称包含非 ASCII 字符，则该名称将按照 RFC 2047 呈现。

### type AddressParser <- go1.5

```go
type AddressParser struct {
	// WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.
	WordDecoder *mime.WordDecoder
}
```

An AddressParser is an [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) address parser.

​	AddressParser 是 RFC 5322 地址解析器。

#### (*AddressParser) Parse <- go1.5

```go
func (p *AddressParser) Parse(address string) (*Address, error)
```

Parse parses a single [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) address of the form “Gogh Fir [gf@example.com](mailto:gf@example.com)” or “[foo@example.com](mailto:foo@example.com)”.

​	Parse 解析单个 RFC 5322 地址，格式为“Gogh Fir gf@example.com”或“foo@example.com”。

#### (*AddressParser) ParseList <- go1.5

```go
func (p *AddressParser) ParseList(list string) ([]*Address, error)
```

ParseList parses the given string as a list of comma-separated addresses of the form “Gogh Fir [gf@example.com](mailto:gf@example.com)” or “[foo@example.com](mailto:foo@example.com)”.

​	ParseList 将给定字符串解析为以逗号分隔的地址列表，格式为“Gogh Fir gf@example.com”或“foo@example.com”。

### type Header

```go
type Header map[string][]string
```

A Header represents the key-value pairs in a mail message header.

​	Header 表示邮件报头中的键值对。

#### (Header) AddressList

```go
func (h Header) AddressList(key string) ([]*Address, error)
```

AddressList parses the named header field as a list of addresses.

​	AddressList 将命名的标头字段解析为地址列表。

#### (Header) Date

```go
func (h Header) Date() (time.Time, error)
```

Date parses the Date header field.

​	Date 解析 Date 标头字段。

#### (Header) Get

```go
func (h Header) Get(key string) string
```

Get gets the first value associated with the given key. It is case insensitive; CanonicalMIMEHeaderKey is used to canonicalize the provided key. If there are no values associated with the key, Get returns “”. To access multiple values of a key, or to use non-canonical keys, access the map directly.

​	Get 获取与给定键关联的第一个值。它不区分大小写；CanonicalMIMEHeaderKey 用于规范化提供的键。如果与键没有关联的值，则 Get 返回“”。要访问键的多个值或使用非规范键，请直接访问映射。

### type Message

```go
type Message struct {
	Header Header
	Body   io.Reader
}
```

A Message represents a parsed mail message.

​	Message 表示已解析的邮件消息。

#### func ReadMessage

```go
func ReadMessage(r io.Reader) (msg *Message, err error)
```

ReadMessage reads a message from r. The headers are parsed, and the body of the message will be available for reading from msg.Body.

​	ReadMessage 从 r 读取消息。将解析标头，并且消息正文可供从 msg.Body 读取。

##### ReadMessage Example
``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"strings"
)

func main() {
	msg := `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body
`

	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}
Output:

Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon
Message body
```

