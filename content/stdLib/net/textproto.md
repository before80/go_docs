+++
title = "textproto"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/textproto@go1.21.3](https://pkg.go.dev/net/textproto@go1.21.3)

Package textproto implements generic support for text-based request/response protocols in the style of HTTP, NNTP, and SMTP.

​	textproto 包为 HTTP、NNTP 和 SMTP 风格的基于文本的请求/响应协议实现了通用支持。

The package provides:

​	该包提供：

Error, which represents a numeric error response from a server.

​	Error，表示来自服务器的数字错误响应。

Pipeline, to manage pipelined requests and responses in a client.

​	Pipeline，在客户端管理流水线请求和响应。

Reader, to read numeric response code lines, key: value headers, lines wrapped with leading spaces on continuation lines, and whole text blocks ending with a dot on a line by itself.

​	Reader，读取数字响应代码行、键：值标头、在延续行上用前导空格包装的行以及以单独一行上的点结尾的整个文本块。

Writer, to write dot-encoded text blocks.

​	Writer，写入点编码的文本块。

Conn, a convenient packaging of Reader, Writer, and Pipeline for use with a single network connection.

​	Conn，一个方便的 Reader、Writer 和 Pipeline 的打包，可与单个网络连接一起使用。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func CanonicalMIMEHeaderKey

```go
func CanonicalMIMEHeaderKey(s string) string
```

CanonicalMIMEHeaderKey returns the canonical format of the MIME header key s. The canonicalization converts the first letter and any letter following a hyphen to upper case; the rest are converted to lowercase. For example, the canonical key for “accept-encoding” is “Accept-Encoding”. MIME header keys are assumed to be ASCII only. If s contains a space or invalid header field bytes, it is returned without modifications.

​	CanonicalMIMEHeaderKey 返回 MIME 头部键 s 的规范格式。规范化将第一个字母和连字符后面的任何字母转换为大写；其余的转换为小写。例如，“accept-encoding”的规范键是“Accept-Encoding”。MIME 头部键假定仅为 ASCII。如果 s 包含空格或无效的头部字段字节，则会原样返回。

### func TrimBytes <- go1.1

```go
func TrimBytes(b []byte) []byte
```

TrimBytes returns b without leading and trailing ASCII space.

​	TrimBytes 返回不带前导和尾随 ASCII 空格的 b。

### func TrimString <- go1.1

```go
func TrimString(s string) string
```

TrimString returns s without leading and trailing ASCII space.

​	TrimString 返回不带前导和尾随 ASCII 空格的 s。

## 类型

### type Conn

```go
type Conn struct {
	Reader
	Writer
	Pipeline
	// contains filtered or unexported fields
}
```

A Conn represents a textual network protocol connection. It consists of a Reader and Writer to manage I/O and a Pipeline to sequence concurrent requests on the connection. These embedded types carry methods with them; see the documentation of those types for details.

​	Conn 表示文本网络协议连接。它由一个 Reader 和 Writer 组成，用于管理 I/O，以及一个 Pipeline，用于对连接上的并发请求进行排序。这些嵌入式类型带有方法；有关详细信息，请参阅这些类型的文档。

#### func Dial

```go
func Dial(network, addr string) (*Conn, error)
```

Dial connects to the given address on the given network using net.Dial and then returns a new Conn for the connection.

​	Dial 使用 net.Dial 连接到给定网络上的给定地址，然后为该连接返回一个新的 Conn。

#### func NewConn

```go
func NewConn(conn io.ReadWriteCloser) *Conn
```

NewConn returns a new Conn using conn for I/O.

​	NewConn 返回一个使用 conn 进行 I/O 的新 Conn。

#### (*Conn) Close

```go
func (c *Conn) Close() error
```

Close closes the connection.

​	Close 关闭连接。

#### (*Conn) Cmd

```go
func (c *Conn) Cmd(format string, args ...any) (id uint, err error)
```

Cmd is a convenience method that sends a command after waiting its turn in the pipeline. The command text is the result of formatting format with args and appending \r\n. Cmd returns the id of the command, for use with StartResponse and EndResponse.

​	Cmd 是一种便捷方法，它在管道中等待轮到它之后发送命令。命令文本是使用 args 格式化 format 并追加 \r\n 的结果。Cmd 返回命令的 id，可与 StartResponse 和 EndResponse 配合使用。

For example, a client might run a HELP command that returns a dot-body by using:

​	例如，客户端可以使用以下命令运行返回点正文的 HELP 命令：

```
id, err := c.Cmd("HELP")
if err != nil {
	return nil, err
}

c.StartResponse(id)
defer c.EndResponse(id)

if _, _, err = c.ReadCodeLine(110); err != nil {
	return nil, err
}
text, err := c.ReadDotBytes()
if err != nil {
	return nil, err
}
return c.ReadCodeLine(250)
```

### type Error

```go
type Error struct {
	Code int
	Msg  string
}
```

An Error represents a numeric error response from a server.

​	Error 表示来自服务器的数字错误响应。

#### (*Error) Error

```go
func (e *Error) Error() string
```

### type MIMEHeader

```go
type MIMEHeader map[string][]string
```

A MIMEHeader represents a MIME-style header mapping keys to sets of values.

​	MIMEHeader 表示将键映射到值集的 MIME 样式标头。

#### (MIMEHeader) Add

```go
func (h MIMEHeader) Add(key, value string)
```

Add adds the key, value pair to the header. It appends to any existing values associated with key.

​	Add 将键值对添加到标头。它追加到与键关联的任何现有值。

#### (MIMEHeader) Del

```go
func (h MIMEHeader) Del(key string)
```

Del deletes the values associated with key.

​	Del 删除与键关联的值。

#### (MIMEHeader) Get

```go
func (h MIMEHeader) Get(key string) string
```

Get gets the first value associated with the given key. It is case insensitive; CanonicalMIMEHeaderKey is used to canonicalize the provided key. If there are no values associated with the key, Get returns “”. To use non-canonical keys, access the map directly.

​	Get 获取与给定键关联的第一个值。它不区分大小写；CanonicalMIMEHeaderKey 用于规范化提供的键。如果与键没有关联的值，则 Get 返回“”。要使用非规范键，请直接访问映射。

#### (MIMEHeader) Set

```go
func (h MIMEHeader) Set(key, value string)
```

Set sets the header entries associated with key to the single element value. It replaces any existing values associated with key.

​	Set 将与键关联的标头条目设置为单个元素值。它替换与键关联的任何现有值。

#### (MIMEHeader) Values <- go1.14

```go
func (h MIMEHeader) Values(key string) []string
```

Values returns all values associated with the given key. It is case insensitive; CanonicalMIMEHeaderKey is used to canonicalize the provided key. To use non-canonical keys, access the map directly. The returned slice is not a copy.

​	Values 返回与给定键关联的所有值。它不区分大小写；CanonicalMIMEHeaderKey 用于规范化提供的键。要使用非规范键，请直接访问映射。返回的切片不是副本。

### type Pipeline

```go
type Pipeline struct {
	// contains filtered or unexported fields
}
```

A Pipeline manages a pipelined in-order request/response sequence.

​	Pipeline 管理流水线中的按顺序请求/响应序列。

To use a Pipeline p to manage multiple clients on a connection, each client should run:

​	要使用管道 p 来管理连接上的多个客户端，每个客户端都应运行：

```
id := p.Next()	// take a number

p.StartRequest(id)	// wait for turn to send request
«send request»
p.EndRequest(id)	// notify Pipeline that request is sent

p.StartResponse(id)	// wait for turn to read response
«read response»
p.EndResponse(id)	// notify Pipeline that response is read
```

A pipelined server can use the same calls to ensure that responses computed in parallel are written in the correct order.

​	管道式服务器可以使用相同的调用来确保并行计算的响应按正确顺序写入。

#### (*Pipeline) EndRequest

```go
func (p *Pipeline) EndRequest(id uint)
```

EndRequest notifies p that the request with the given id has been sent (or, if this is a server, received).

​	EndRequest 通知 p 已发送（或如果这是服务器，已收到）具有给定 ID 的请求。

#### (*Pipeline) EndResponse

```go
func (p *Pipeline) EndResponse(id uint)
```

EndResponse notifies p that the response with the given id has been received (or, if this is a server, sent).

​	EndResponse 通知 p 已收到（或如果这是服务器，已发送）具有给定 ID 的响应。

#### (*Pipeline) Next

```go
func (p *Pipeline) Next() uint
```

Next returns the next id for a request/response pair.

​	Next 返回请求/响应对的下一个 ID。

#### (*Pipeline) StartRequest

```go
func (p *Pipeline) StartRequest(id uint)
```

StartRequest blocks until it is time to send (or, if this is a server, receive) the request with the given id.

​	StartRequest 阻塞，直到可以发送（或如果这是服务器，可以接收）具有给定 ID 的请求。

#### (*Pipeline) StartResponse 

```go
func (p *Pipeline) StartResponse(id uint)
```

StartResponse blocks until it is time to receive (or, if this is a server, send) the request with the given id.

​	StartResponse 会阻塞，直到该请求具有给定 ID 时接收（或如果这是服务器，则发送）请求。

### type ProtocolError

```go
type ProtocolError string
```

A ProtocolError describes a protocol violation such as an invalid response or a hung-up connection.

​	ProtocolError 描述协议违规，例如无效响应或挂起的连接。

#### (ProtocolError) Error 

```go
func (p ProtocolError) Error() string
```

### type Reader

```go
type Reader struct {
	R *bufio.Reader
	// contains filtered or unexported fields
}
```

A Reader implements convenience methods for reading requests or responses from a text protocol network connection.

​	Reader 实现从文本协议网络连接读取请求或响应的便捷方法。

#### func NewReader

```go
func NewReader(r *bufio.Reader) *Reader
```

NewReader returns a new Reader reading from r.

​	NewReader 返回一个从 r 读取的新 Reader。

To avoid denial of service attacks, the provided bufio.Reader should be reading from an io.LimitReader or similar Reader to bound the size of responses.

​	为了避免拒绝服务攻击，提供的 bufio.Reader 应从 io.LimitReader 或类似的 Reader 读取，以限制响应的大小。

#### (*Reader) DotReader

```go
func (r *Reader) DotReader() io.Reader
```

DotReader returns a new Reader that satisfies Reads using the decoded text of a dot-encoded block read from r. The returned Reader is only valid until the next call to a method on r.

​	DotReader 返回一个新的 Reader，该 Reader 使用从 r 读取的点编码块的解码文本满足读取。返回的 Reader 仅在对 r 上的方法进行下一次调用之前有效。

Dot encoding is a common framing used for data blocks in text protocols such as SMTP. The data consists of a sequence of lines, each of which ends in “\r\n”. The sequence itself ends at a line containing just a dot: “.\r\n”. Lines beginning with a dot are escaped with an additional dot to avoid looking like the end of the sequence.

​	点编码是文本协议（如 SMTP）中数据块常用的框架。数据由一系列行组成，每行以“\r\n”结尾。序列本身以仅包含一个点的行结束：“.\r\n”。以点开头的行用一个额外的点转义，以避免看起来像序列的结尾。

The decoded form returned by the Reader’s Read method rewrites the “\r\n” line endings into the simpler “\n”, removes leading dot escapes if present, and stops with error io.EOF after consuming (and discarding) the end-of-sequence line.

​	Reader 的 Read 方法返回的解码形式将“\r\n”行尾重写为更简单的“\n”，如果存在，则删除前导点转义，并在使用（并丢弃）序列结束行后停止，并出现错误 io.EOF。

#### (*Reader) ReadCodeLine

```go
func (r *Reader) ReadCodeLine(expectCode int) (code int, message string, err error)
```

ReadCodeLine reads a response code line of the form

​	ReadCodeLine 读取形式为以下的响应代码行

```
code message
```

where code is a three-digit status code and the message extends to the rest of the line. An example of such a line is:

​	其中 code 是一个三位状态代码，message 扩展到行的其余部分。此类行的示例如下：

```
220 plan9.bell-labs.com ESMTP
```

If the prefix of the status does not match the digits in expectCode, ReadCodeLine returns with err set to &Error{code, message}. For example, if expectCode is 31, an error will be returned if the status is not in the range [310,319].

​	如果状态的前缀与 expectCode 中的数字不匹配，ReadCodeLine 返回，并将 err 设置为 &Error{code, message}。例如，如果 expectCode 为 31，如果状态不在范围 [310,319] 中，则会返回错误。

If the response is multi-line, ReadCodeLine returns an error.

​	如果响应是多行的，ReadCodeLine 返回错误。

An expectCode <= 0 disables the check of the status code.

​	expectCode <= 0 禁用状态代码检查。

#### (*Reader) ReadContinuedLine

```go
func (r *Reader) ReadContinuedLine() (string, error)
```

ReadContinuedLine reads a possibly continued line from r, eliding the final trailing ASCII white space. Lines after the first are considered continuations if they begin with a space or tab character. In the returned data, continuation lines are separated from the previous line only by a single space: the newline and leading white space are removed.

​	ReadContinuedLine 从 r 读取可能已继续的行，省略最终的尾随 ASCII 空白。如果以空格或制表符字符开头，则第一行之后的行被视为续行。在返回的数据中，续行仅由单个空格与前一行分隔：删除换行符和前导空白。

For example, consider this input:

​	例如，考虑以下输入：

```
Line 1
  continued...
Line 2
```

The first call to ReadContinuedLine will return “Line 1 continued…” and the second will return “Line 2”.

​	对 ReadContinuedLine 的第一次调用将返回“Line 1 continued…”，第二次将返回“Line 2”。

Empty lines are never continued.

​	空行永远不会继续。

#### (*Reader) ReadContinuedLineBytes

```go
func (r *Reader) ReadContinuedLineBytes() ([]byte, error)
```

ReadContinuedLineBytes is like ReadContinuedLine but returns a []byte instead of a string.

​	ReadContinuedLineBytes 与 ReadContinuedLine 类似，但返回 []byte 而不是字符串。

#### (*Reader) ReadDotBytes

```go
func (r *Reader) ReadDotBytes() ([]byte, error)
```

ReadDotBytes reads a dot-encoding and returns the decoded data.

​	ReadDotBytes 读取点编码并返回解码后的数据。

See the documentation for the DotReader method for details about dot-encoding.

​	有关点编码的详细信息，请参阅 DotReader 方法的文档。

#### (*Reader) ReadDotLines

```go
func (r *Reader) ReadDotLines() ([]string, error)
```

ReadDotLines reads a dot-encoding and returns a slice containing the decoded lines, with the final \r\n or \n elided from each.

​	ReadDotLines 读取点编码并返回包含已解码行的切片，每个切片中最后的 \r\n 或 \n 已省略。

See the documentation for the DotReader method for details about dot-encoding.

​	有关点编码的详细信息，请参阅 DotReader 方法的文档。

#### (*Reader) ReadLine

```go
func (r *Reader) ReadLine() (string, error)
```

ReadLine reads a single line from r, eliding the final \n or \r\n from the returned string.

​	ReadLine 从 r 读取单行，从返回的字符串中省略最后的 \n 或 \r\n。

#### (*Reader) ReadLineBytes

```go
func (r *Reader) ReadLineBytes() ([]byte, error)
```

ReadLineBytes is like ReadLine but returns a []byte instead of a string.

​	ReadLineBytes 与 ReadLine 类似，但返回 []byte 而不是字符串。

#### (*Reader) ReadMIMEHeader

```go
func (r *Reader) ReadMIMEHeader() (MIMEHeader, error)
```

ReadMIMEHeader reads a MIME-style header from r. The header is a sequence of possibly continued Key: Value lines ending in a blank line. The returned map m maps CanonicalMIMEHeaderKey(key) to a sequence of values in the same order encountered in the input.

​	ReadMIMEHeader 从 r 读取 MIME 样式的标头。标头是一系列可能继续的 Key: Value 行，以空行结尾。返回的映射 m 将 CanonicalMIMEHeaderKey(key) 映射到输入中遇到的相同顺序的值序列。

For example, consider this input:

​	例如，考虑以下输入：

```
My-Key: Value 1
Long-Key: Even
       Longer Value
My-Key: Value 2
```

Given that input, ReadMIMEHeader returns the map:

​	给定该输入，ReadMIMEHeader 返回映射：

```
map[string][]string{
	"My-Key": {"Value 1", "Value 2"},
	"Long-Key": {"Even Longer Value"},
}
```

#### (*Reader) ReadResponse

```go
func (r *Reader) ReadResponse(expectCode int) (code int, message string, err error)
```

ReadResponse reads a multi-line response of the form:

​	ReadResponse 读取如下形式的多行响应：

```
code-message line 1
code-message line 2
...
code message line n
```

where code is a three-digit status code. The first line starts with the code and a hyphen. The response is terminated by a line that starts with the same code followed by a space. Each line in message is separated by a newline (\n).

​	其中 code 是一个三位数状态代码。第一行以 code 和连字符开头。响应以一行结尾，该行以相同的 code 后跟一个空格开头。message 中的每一行都以换行符 (\n) 分隔。

See page 36 of [RFC 959](https://rfc-editor.org/rfc/rfc959.html) (https://www.ietf.org/rfc/rfc959.txt) for details of another form of response accepted:

​	有关接受的另一种响应形式的详细信息，请参阅 RFC 959（https://www.ietf.org/rfc/rfc959.txt）的第 36 页：

```
code-message line 1
message line 2
...
code message line n
```

If the prefix of the status does not match the digits in expectCode, ReadResponse returns with err set to &Error{code, message}. For example, if expectCode is 31, an error will be returned if the status is not in the range [310,319].

​	如果状态的前缀与 expectCode 中的数字不匹配，ReadResponse 返回，并将 err 设置为 &Error{code, message}。例如，如果 expectCode 为 31，则如果状态不在 [310,319] 范围内，将返回错误。

An expectCode <= 0 disables the check of the status code.

​	expectCode <= 0 禁用状态代码检查。

### type Writer

```go
type Writer struct {
	W *bufio.Writer
	// contains filtered or unexported fields
}
```

A Writer implements convenience methods for writing requests or responses to a text protocol network connection.

​	Writer 实现将请求或响应写入文本协议网络连接的便捷方法。

#### func NewWriter

```go
func NewWriter(w *bufio.Writer) *Writer
```

NewWriter returns a new Writer writing to w.

​	NewWriter 返回一个新的 Writer，写入 w。

#### (*Writer) DotWriter

```go
func (w *Writer) DotWriter() io.WriteCloser
```

DotWriter returns a writer that can be used to write a dot-encoding to w. It takes care of inserting leading dots when necessary, translating line-ending \n into \r\n, and adding the final .\r\n line when the DotWriter is closed. The caller should close the DotWriter before the next call to a method on w.

​	DotWriter 返回一个可用于向 w 写入点编码的编写器。它负责在必要时插入前导点，将行尾 \n 转换为 \r\n，并在 DotWriter 关闭时添加最终的 .\r\n 行。调用者应在对 w 上的方法进行下一次调用之前关闭 DotWriter。

See the documentation for Reader’s DotReader method for details about dot-encoding.

​	有关点编码的详细信息，请参阅 Reader 的 DotReader 方法的文档。

#### (*Writer) PrintfLine

```go
func (w *Writer) PrintfLine(format string, args ...any) error
```

PrintfLine writes the formatted output followed by \r\n.

​	PrintfLine 编写格式化的输出，后跟 \r\n。