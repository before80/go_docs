+++
title = "smtp"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/smtp@go1.21.3](https://pkg.go.dev/net/smtp@go1.21.3)

Package smtp implements the Simple Mail Transfer Protocol as defined in [RFC 5321](https://rfc-editor.org/rfc/rfc5321.html). It also implements the following extensions:

​	smtp 包实现了 RFC 5321 中定义的简单邮件传输协议。它还实现了以下扩展：

```
8BITMIME  RFC 1652
AUTH      RFC 2554
STARTTLS  RFC 3207
```

Additional extensions may be handled by clients.

​	客户端可以处理其他扩展。

The smtp package is frozen and is not accepting new features. Some external packages provide more functionality. See:

​	smtp 包已冻结，不接受新功能。一些外部包提供了更多功能。请参阅：

```
https://godoc.org/?q=smtp
```

## Example 示例

```go
package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("sender@example.org"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("recipient@example.net"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
```

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func SendMail

```go
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
```

SendMail connects to the server at addr, switches to TLS if possible, authenticates with the optional mechanism a if possible, and then sends an email from address from, to addresses to, with message msg. The addr must include a port, as in “mail.example.com:smtp”.

​	SendMail 连接到 addr 处的服务器，如果可能，切换到 TLS，如果可能，使用可选机制 a 进行身份验证，然后从地址 from 向地址 to 发送电子邮件，其中包含消息 msg。addr 必须包含端口，例如“mail.example.com:smtp”。

The addresses in the to parameter are the SMTP RCPT addresses.

​	to 参数中的地址是 SMTP RCPT 地址。

The msg parameter should be an [RFC 822](https://rfc-editor.org/rfc/rfc822.html)-style email with headers first, a blank line, and then the message body. The lines of msg should be CRLF terminated. The msg headers should usually include fields such as “From”, “To”, “Subject”, and “Cc”. Sending “Bcc” messages is accomplished by including an email address in the to parameter but not including it in the msg headers.

​	msg 参数应为 RFC 822 样式的电子邮件，首先是标题、一个空行，然后是邮件正文。msg 的行应以 CRLF 结尾。msg 标题通常应包括“发件人”、“收件人”、“主题”和“抄送”等字段。发送“密件抄送”邮件可通过在 to 参数中包含电子邮件地址，但不将其包含在 msg 标题中来实现。

The SendMail function and the net/smtp package are low-level mechanisms and provide no support for DKIM signing, MIME attachments (see the mime/multipart package), or other mail functionality. Higher-level packages exist outside of the standard library.

​	SendMail 函数和 net/smtp 包是低级机制，不提供对 DKIM 签名、MIME 附件（请参阅 mime/multipart 包）或其他邮件功能的支持。高级包存在于标准库之外。

#### SendMail Example

```go
package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
```

## 类型

### type Auth

```go
type Auth interface {
	// Start begins an authentication with a server.
	// It returns the name of the authentication protocol
	// and optionally data to include in the initial AUTH message
	// sent to the server.
	// If it returns a non-nil error, the SMTP client aborts
	// the authentication attempt and closes the connection.
	Start(server *ServerInfo) (proto string, toServer []byte, err error)

	// Next continues the authentication. The server has just sent
	// the fromServer data. If more is true, the server expects a
	// response, which Next should return as toServer; otherwise
	// Next should return toServer == nil.
	// If Next returns a non-nil error, the SMTP client aborts
	// the authentication attempt and closes the connection.
	Next(fromServer []byte, more bool) (toServer []byte, err error)
}
```

Auth is implemented by an SMTP authentication mechanism.

​	Auth 由 SMTP 身份验证机制实现。

#### func CRAMMD5Auth

```go
func CRAMMD5Auth(username, secret string) Auth
```

CRAMMD5Auth returns an Auth that implements the CRAM-MD5 authentication mechanism as defined in [RFC 2195](https://rfc-editor.org/rfc/rfc2195.html). The returned Auth uses the given username and secret to authenticate to the server using the challenge-response mechanism.

​	CRAMMD5Auth 返回一个 Auth，该 Auth 实现 RFC 2195 中定义的 CRAM-MD5 身份验证机制。返回的 Auth 使用给定的用户名和密码使用质询-响应机制向服务器进行身份验证。

#### func PlainAuth

```go
func PlainAuth(identity, username, password, host string) Auth
```

PlainAuth returns an Auth that implements the PLAIN authentication mechanism as defined in [RFC 4616](https://rfc-editor.org/rfc/rfc4616.html). The returned Auth uses the given username and password to authenticate to host and act as identity. Usually identity should be the empty string, to act as username.

​	PlainAuth 返回一个 Auth，该 Auth 实现 RFC 4616 中定义的 PLAIN 身份验证机制。返回的 Auth 使用给定的用户名和密码向主机进行身份验证并充当身份。通常身份应为空字符串，以充当用户名。

PlainAuth will only send the credentials if the connection is using TLS or is connected to localhost. Otherwise authentication will fail with an error, without sending the credentials.

​	PlainAuth 仅在连接使用 TLS 或连接到 localhost 时才发送凭据。否则，身份验证将失败并出现错误，而不会发送凭据。

##### PlainAuth Example

```go
package main

import (
	"log"
	"net/smtp"
)

// variables to make ExamplePlainAuth compile, without adding
// unnecessary noise there.
var (
	from       = "gopher@example.net"
	msg        = []byte("dummy message")
	recipients = []string{"foo@example.com"}
)

func main() {
	// hostname is used by PlainAuth to validate the TLS certificate.
	hostname := "mail.example.com"
	auth := smtp.PlainAuth("", "user@example.com", "password", hostname)

	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}
```

### type Client 

```go
type Client struct {
	// Text is the textproto.Conn used by the Client. It is exported to allow for
	// clients to add extensions.
	Text *textproto.Conn
	// contains filtered or unexported fields
}
```

A Client represents a client connection to an SMTP server.

​	Client 表示与 SMTP 服务器的客户端连接。

#### func Dial

```go
func Dial(addr string) (*Client, error)
```

Dial returns a new Client connected to an SMTP server at addr. The addr must include a port, as in “mail.example.com:smtp”.

​	Dial 返回一个新的 Client，该 Client 连接到 addr 处的 SMTP 服务器。addr 必须包含端口，例如“mail.example.com:smtp”。

#### func NewClient

```go
func NewClient(conn net.Conn, host string) (*Client, error)
```

NewClient returns a new Client using an existing connection and host as a server name to be used when authenticating.

​	NewClient 使用现有连接和主机作为服务器名称返回一个新的 Client，该服务器名称将在进行身份验证时使用。

#### (*Client) Auth

```go
func (c *Client) Auth(a Auth) error
```

Auth authenticates a client using the provided authentication mechanism. A failed authentication closes the connection. Only servers that advertise the AUTH extension support this function.

​	Auth 使用提供的身份验证机制对客户端进行身份验证。身份验证失败将关闭连接。只有宣传 AUTH 扩展的服务器才支持此功能。

#### (*Client) Close <- go1.2

```go
func (c *Client) Close() error
```

Close closes the connection.

​	Close 关闭连接。

#### (*Client) Data 

```go
func (c *Client) Data() (io.WriteCloser, error)
```

Data issues a DATA command to the server and returns a writer that can be used to write the mail headers and body. The caller should close the writer before calling any more methods on c. A call to Data must be preceded by one or more calls to Rcpt.

​	数据向服务器发出 DATA 命令，并返回一个可用于编写邮件头和正文的编写器。调用者应在对 c 调用更多方法之前关闭编写器。必须在调用 Rcpt 一次或多次之后才能调用 Data。 （*客户端）扩展 扩展报告服务器是否支持扩展。扩展名不区分大小写。如果支持扩展，扩展还会返回一个字符串，其中包含服务器为扩展指定的所有参数。 （*客户端）Hello <- go1.1 Hello 将 HELO 或 EHLO 作为给定的主机名发送到服务器。仅当客户端需要控制所用主机名时，才需要调用此方法。否则，客户端会自动将自身介绍为“localhost”。如果调用了 Hello，则必须在调用其他任何方法之前调用它。 （*客户端）邮件 邮件使用提供的电子邮件地址向服务器发出 MAIL 命令。如果服务器支持 8BITMIME 扩展，则 Mail 会添加 BODY=8BITMIME 参数。如果服务器支持 SMTPUTF8 扩展，则 Mail 会添加 SMTPUTF8 参数。这会启动邮件事务，然后是调用 Rcpt 一次或多次。

#### (*Client) Extension

```go
func (c *Client) Extension(ext string) (bool, string)
```

Extension reports whether an extension is support by the server. The extension name is case-insensitive. If the extension is supported, Extension also returns a string that contains any parameters the server specifies for the extension.

#### (*Client) Hello <- go1.1

```go
func (c *Client) Hello(localName string) error
```

Hello sends a HELO or EHLO to the server as the given host name. Calling this method is only necessary if the client needs control over the host name used. The client will introduce itself as “localhost” automatically otherwise. If Hello is called, it must be called before any of the other methods.

#### (*Client) Mail

```go
func (c *Client) Mail(from string) error
```

Mail issues a MAIL command to the server using the provided email address. If the server supports the 8BITMIME extension, Mail adds the BODY=8BITMIME parameter. If the server supports the SMTPUTF8 extension, Mail adds the SMTPUTF8 parameter. This initiates a mail transaction and is followed by one or more Rcpt calls.

#### (*Client) Noop <- go1.10

```go
func (c *Client) Noop() error
```

Noop sends the NOOP command to the server. It does nothing but check that the connection to the server is okay.

​	Noop 向服务器发送 NOOP 命令。它除了检查与服务器的连接是否正常外，什么也不做。

#### (*Client) Quit 

```go
func (c *Client) Quit() error
```

Quit sends the QUIT command and closes the connection to the server.

​	Quit 发送 QUIT 命令并关闭与服务器的连接。

#### (*Client) Rcpt 

```go
func (c *Client) Rcpt(to string) error
```

Rcpt issues a RCPT command to the server using the provided email address. A call to Rcpt must be preceded by a call to Mail and may be followed by a Data call or another Rcpt call.

​	Rcpt 使用提供的电子邮件地址向服务器发出 RCPT 命令。必须在调用 Rcpt 之前调用 Mail，并且可以在调用 Rcpt 之后调用 Data 调用或另一个 Rcpt 调用。

#### (*Client) Reset 

```go
func (c *Client) Reset() error
```

Reset sends the RSET command to the server, aborting the current mail transaction.

​	Reset 向服务器发送 RSET 命令，中止当前邮件事务。

#### (*Client) StartTLS 

```go
func (c *Client) StartTLS(config *tls.Config) error
```

StartTLS sends the STARTTLS command and encrypts all further communication. Only servers that advertise the STARTTLS extension support this function.

​	StartTLS 发送 STARTTLS 命令并加密所有进一步的通信。只有宣传 STARTTLS 扩展的服务器才支持此功能。

#### (*Client) TLSConnectionState <- go1.5 

```go
func (c *Client) TLSConnectionState() (state tls.ConnectionState, ok bool)
```

TLSConnectionState returns the client’s TLS connection state. The return values are their zero values if StartTLS did not succeed.

​	TLSConnectionState 返回客户端的 TLS 连接状态。如果 StartTLS 未成功，则返回值为其零值。

#### (*Client) Verify 

```go
func (c *Client) Verify(addr string) error
```

Verify checks the validity of an email address on the server. If Verify returns nil, the address is valid. A non-nil return does not necessarily indicate an invalid address. Many servers will not verify addresses for security reasons.

​	Verify 检查服务器上电子邮件地址的有效性。如果 Verify 返回 nil，则该地址有效。非 nil 返回并不一定表示地址无效。出于安全原因，许多服务器不会验证地址。

### type ServerInfo

```go
type ServerInfo struct {
	Name string   // SMTP server name
	TLS  bool     // using TLS, with valid certificate for Name
	Auth []string // advertised authentication mechanisms
}
```

ServerInfo records information about an SMTP server.

​	ServerInfo 记录有关 SMTP 服务器的信息。