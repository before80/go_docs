+++
title = "mail"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# mail

https://pkg.go.dev/net/mail@go1.20.1



Package mail implements parsing of mail messages.

For the most part, this package follows the syntax as specified by [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) and extended by [RFC 6532](https://rfc-editor.org/rfc/rfc6532.html). Notable divergences:

- Obsolete address formats are not parsed, including addresses with embedded route information.
- The full range of spacing (the CFWS syntax element) is not supported, such as breaking addresses across lines.
- No unicode normalization is performed.
- The special characters ()[]:;@\, are allowed to appear unquoted in names.









## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=160)

``` go 
var ErrHeaderNotPresent = errors.New("mail: header not in message")
```

## 函数

#### func [ParseDate](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=99)  <- go1.8

``` go 
func ParseDate(date string) (time.Time, error)
```

ParseDate parses an [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) date string.

## 类型

### type [Address](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=183) 

``` go 
type Address struct {
	Name    string // Proper name; may be empty.
	Address string // user@domain
}
```

Address represents a single mail address. An address such as "Barry Gibbs <bg@example.com>" is represented as Address{Name: "Barry Gibbs", Address: "bg@example.com"}.

#### func [ParseAddress](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=189)  <- go1.1

``` go 
func ParseAddress(address string) (*Address, error)
```

ParseAddress parses a single [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) address, e.g. "Barry Gibbs <bg@example.com>"

##### Example
``` go 
```

#### func [ParseAddressList](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=194)  <- go1.1

``` go 
func ParseAddressList(list string) ([]*Address, error)
```

ParseAddressList parses the given string as a list of addresses.

##### Example
``` go 
```

#### (*Address) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=219) 

``` go 
func (a *Address) String() string
```

String formats the address as a valid [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) address. If the address's name contains non-ASCII characters the name will be rendered according to [RFC 2047](https://rfc-editor.org/rfc/rfc2047.html).

### type [AddressParser](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=199)  <- go1.5

``` go 
type AddressParser struct {
	// WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.
	WordDecoder *mime.WordDecoder
}
```

An AddressParser is an [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) address parser.

#### (*AddressParser) [Parse](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=206)  <- go1.5

``` go 
func (p *AddressParser) Parse(address string) (*Address, error)
```

Parse parses a single [RFC 5322](https://rfc-editor.org/rfc/rfc5322.html) address of the form "Gogh Fir <gf@example.com>" or "foo@example.com".

#### (*AddressParser) [ParseList](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=212)  <- go1.5

``` go 
func (p *AddressParser) ParseList(list string) ([]*Address, error)
```

ParseList parses the given string as a list of comma-separated addresses of the form "Gogh Fir <gf@example.com>" or "foo@example.com".

### type [Header](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=148) 

``` go 
type Header map[string][]string
```

A Header represents the key-value pairs in a mail message header.

#### (Header) [AddressList](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=172) 

``` go 
func (h Header) AddressList(key string) ([]*Address, error)
```

AddressList parses the named header field as a list of addresses.

#### (Header) [Date](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=163) 

``` go 
func (h Header) Date() (time.Time, error)
```

Date parses the Date header field.

#### (Header) [Get](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=156) 

``` go 
func (h Header) Get(key string) string
```

Get gets the first value associated with the given key. It is case insensitive; CanonicalMIMEHeaderKey is used to canonicalize the provided key. If there are no values associated with the key, Get returns "". To access multiple values of a key, or to use non-canonical keys, access the map directly.

### type [Message](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=45) 

``` go 
type Message struct {
	Header Header
	Body   io.Reader
}
```

A Message represents a parsed mail message.

#### func [ReadMessage](https://cs.opensource.google/go/go/+/go1.20.1:src/net/mail/message.go;l=53) 

``` go 
func ReadMessage(r io.Reader) (msg *Message, err error)
```

ReadMessage reads a message from r. The headers are parsed, and the body of the message will be available for reading from msg.Body.

##### Example
``` go 
```

