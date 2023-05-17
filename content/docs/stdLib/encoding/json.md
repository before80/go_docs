+++
title = "json"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# json

[https://pkg.go.dev/encoding/json@go1.20.1](https://pkg.go.dev/encoding/json@go1.20.1)

​	json包实现了[RFC 7159](https://rfc-editor.org/rfc/rfc7159.html)中定义的JSON编解码。JSON和Go值之间的映射在Marshal和Unmarshal函数的文档中有描述。

​	有关此包的介绍，请参见[JSON和Go](https://golang.org/doc/articles/json_and_go.html)。

##### Example (CustomMarshalJSON)
``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

func main() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}
Output:

Zoo Census:
* Gophers: 3
* Zebras:  2
* Unknown: 3
```

##### Example(TextMarshalJSON)
``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Size int

const (
	Unrecognized Size = iota
	Small
	Large
)

func (s *Size) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	default:
		*s = Unrecognized
	case "small":
		*s = Small
	case "large":
		*s = Large
	}
	return nil
}

func (s Size) MarshalText() ([]byte, error) {
	var name string
	switch s {
	default:
		name = "unrecognized"
	case Small:
		name = "small"
	case Large:
		name = "large"
	}
	return []byte(name), nil
}

func main() {
	blob := `["small","regular","large","unrecognized","small","normal","small","large"]`
	var inventory []Size
	if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}

	counts := make(map[Size]int)
	for _, size := range inventory {
		counts[size] += 1
	}

	fmt.Printf("Inventory Counts:\n* Small:        %d\n* Large:        %d\n* Unrecognized: %d\n",
		counts[Small], counts[Large], counts[Unrecognized])

}
Output:

Inventory Counts:
* Small:        3
* Large:        2
* Unrecognized: 3
```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [Compact](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/indent.go;l=13) 

``` go linenums="1"
func Compact(dst *bytes.Buffer, src []byte) error
```

​	Compact函数将JSON编码的src附加到dst中，省略了不重要的空格字符。

#### func [HTMLEscape](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=192) 

``` go linenums="1"
func HTMLEscape(dst *bytes.Buffer, src []byte)
```

​	HTMLEscape函数将JSON编码的src附加到dst中，将字符串文字内的<、>、&、U+2028和U+2029字符更改为\u003c、\u003e、\u0026、\u2028、\u2029，以使JSON可以安全地嵌入HTML `<script>`标记中。由于历史原因，Web浏览器不支持在`<script>`标记中使用标准的HTML转义，因此必须使用替代的JSON编码。

##### HTMLEscape Example
``` go linenums="1"
package main

import (
	"bytes"
	"encoding/json"
	"os"
)

func main() {
	var out bytes.Buffer
	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout)
}
Output:

{"Name":"\u003cb\u003eHTML content\u003c/b\u003e"}
```

#### func [Indent](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/indent.go;l=81) 

``` go linenums="1"
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
```

​	Indent函数将JSON编码的src的缩进形式附加到dst中。JSON对象或数组中的每个元素都在新的缩进行上开始，该行以prefix开头，后跟一个或多个indent的副本，具体取决于缩进嵌套。附加到dst的数据不以prefix或任何缩进开始，以使其更容易嵌入其他格式化的JSON数据中。虽然src开头的前导空格字符(空格、制表符、回车、换行符)会被删除，但src末尾的尾随空格字符会被保留并复制到dst中。例如，如果src没有尾随空格，则dst也没有；如果src以尾随换行符结束，则dst也是如此。

##### Example
``` go linenums="1"
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func main() {
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")
	out.WriteTo(os.Stdout)
}
Output:

[
=	{
=		"Name": "Diamond Fork",
=		"Number": 29
=	},
=	{
=		"Name": "Sheep Creek",
=		"Number": 51
=	}
=]
```

#### func [Marshal](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=157) 

``` go linenums="1"
func Marshal(v any) ([]byte, error)
```

​	Marshal函数返回v的JSON编码。

​	Marshal函数递归遍历值v。如果遇到的值实现了Marshaler接口并且不是nil指针，则Marshal调用其MarshalJSON方法以生成JSON。如果没有MarshalJSON方法但该值代替实现了encoding.TextMarshaler，则Marshal调用其MarshalText方法并将结果编码为JSON字符串。nil指针异常并不是严格必要的，但模仿了在UnmarshalJSON的行为中必要的类似异常。

​	否则，Marshal使用以下类型相关的默认编码：

​	布尔值被编码为JSON布尔值。

​	浮点数、整数和Number类型的值被编码为JSON数字。

​	字符串类型的值被编码为JSON字符串，被强制转换为有效的UTF-8编码，无效的字节会被替换为Unicode替换符。为了将JSON安全地嵌入到HTML的<script>标签中，使用HTMLEscape对字符串进行编码，将"<", ">", "&", U+2028, 和 U+2029转义为"\u003c", "\u003e", "\u0026", "\u2028"和"\u2029"。使用Encoder时，可以通过调用SetEscapeHTML(false)来禁用此替换。

​	数组和切片类型的值被编码为JSON数组，但[]byte类型的值会被编码为base64编码的字符串，而nil切片类型的值会被编码为null。

​	结构体类型的值被编码为JSON对象。每个导出的结构体字段都成为对象的一个成员，使用字段名作为对象键，除非由于以下原因之一而省略该字段。

​	每个结构体字段的编码可以通过存储在该字段标记的"json"键下的格式字符串进行自定义。格式字符串给出字段名，可能后跟一个用逗号分隔的选项列表。如果不想覆盖默认的字段名，则字段名可以为空。

​	"omitempty"选项指定如果该字段具有空值(定义为false、0、nil指针、nil接口值以及任何空数组、切片、映射或字符串)，则应该从编码中省略该字段。

​	作为特殊情况，如果字段标记为"-"，则始终省略该字段。请注意，具有名称"-"的字段仍然可以使用标记"-,"生成。

​	以下是结构体字段标记及其含义的示例：

```
// Field 出现在JSON中作为键"myName"。
Field int `json:"myName"`

// Field 出现在JSON中作为键"myName"，
// 如果其值为空，将从对象中省略该字段。
Field int `json:"myName,omitempty"`

// Field 出现在JSON中作为键"Field"(默认值)，
// 但如果为空则跳过该字段。
// 请注意前导逗号。
Field int `json:",omitempty"`

// 此包忽略 Field。
Field int `json:"-"`

// Field 出现在JSON中作为键"-"。
Field int `json:"-,"`
```

​	"string"选项表示字段存储在JSON编码的字符串中。它仅适用于字符串、浮点数、整数或布尔类型的字段。在与JavaScript程序通信时，有时会使用这种额外的编码级别：

```
Int64String int64 `json:",string"`
```

​	如果键名是由Unicode字母、数字和ASCII标点(引号、反斜杠和逗号除外)组成的非空字符串，则将使用该键名。

​	匿名结构字段通常被解释为其内部导出字段在外部结构体中的字段，但受到下一段中描述的修饰性可见性规则的影响。具有在其JSON标签中给出的名称的匿名结构字段将被视为具有该名称，而不是匿名的。接口类型的匿名结构字段将与其类型相同，而不是匿名的。

​	当决定哪个字段进行编组或取消编组时，针对JSON修饰性可见性规则修正结构字段的Go可见性规则。如果同一级别存在多个字段，并且该级别是最不嵌套的(因此将是常规Go规则选择的嵌套级别)，则以下额外规则适用：

\1) Of those fields, if any are JSON-tagged, only tagged fields are considered, even if there are multiple untagged fields that would otherwise conflict.

1)对于这些字段，如果有任何字段具有JSON标记，则只考虑带有标记的字段，即使有多个未标记的字段也会冲突。

2)如果恰好有一个字段(根据第一条规则是否带有标记)，则选择该字段。

3)否则将忽略所有字段；不会发生错误。

对匿名结构字段的处理是Go 1.1的新内容。在Go 1.1之前，匿名结构字段被忽略。要在当前版本和早期版本中强制忽略匿名结构字段，请给该字段一个"-"的JSON标签。

​	在Go 1.1中，对匿名结构字段的处理是新的。在Go 1.1之前，匿名结构字段被忽略。要在当前版本和早期版本中强制忽略匿名结构字段，请给该字段一个JSON标记"-"。

​	Map值编码为JSON对象。映射的键类型必须是字符串、整数类型或实现encoding.TextMarshaler。映射键按以下规则排序并用作JSON对象键，但要受上述字符串值的UTF-8强制规定：

- 任何字符串类型的键都直接使用。 

- encoding.TextMarshalers已被解组。 

- 整数键被转换为字符串。 


​	指针值编码为指向的值。空指针编码为null JSON值。指针值编码为所指向的值。一个nil指针编码为空的JSON值。

​	接口值编码为接口中包含的值。空接口值编码为null JSON值。

​	通道、复数和函数值不能编码为 JSON。尝试对此类值进行编码将导致 Marshal 返回 UnsupportedTypeError。

​	JSON 不能表示循环数据结构，Marshal 也无法处理它们。将循环结构传递给 Marshal 将导致错误。

##### Marshal Example
``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
Output:

{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
```

#### func [MarshalIndent](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=173) 

``` go linenums="1"
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
```

​	MarshalIndent函数类似于 Marshal函数，但应用 Indent 以格式化输出。输出中的每个 JSON 元素都将在新行上开始，以前缀开头，后跟一个或多个 indent 副本，具体取决于缩进嵌套。

##### MarshalIndent Example
``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	b, err := json.MarshalIndent(data, "<prefix>", "<indent>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
Output:

{
<prefix><indent>"a": 1,
<prefix><indent>"b": 2
<prefix>}
```

#### func [Unmarshal](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=97) 

``` go linenums="1"
func Unmarshal(data []byte, v any) error
```

​	Unmarshal函数解析 JSON 编码的数据并将结果存储在 v 所指向的值中。如果 v 为 nil 或不是指针，则 Unmarshal 返回 InvalidUnmarshalError。

​	Unmarshal函数使用 Marshal函数所使用的编码的反向方式，根据以下附加规则分配映射、切片和指针：

​	要将 JSON 反序列化为指针，Unmarshal 首先处理 JSON 为 JSON 文字 null 的情况。在这种情况下，Unmarshal 将指针设置为 nil。否则，Unmarshal 将 JSON 反序列化为指针所指向的值。如果指针为 nil，则 Unmarshal 为其分配一个新值。

​	要将 JSON 反序列化为实现 Unmarshaler 接口的值，Unmarshal函数调用该值的 UnmarshalJSON 方法，包括输入为 JSON null 的情况。否则，如果该值实现 encoding.TextUnmarshaler 并且输入是 JSON 引号字符串，则 Unmarshal 会使用字符串的未引用形式调用该值的 UnmarshalText 方法。

​	要将 JSON 反序列化为结构体，Unmarshal 将传入的对象键与 Marshal 使用的键进行匹配(即结构字段名或其标记)，优先选择完全匹配，但也接受大小写不敏感的匹配。默认情况下，没有对应结构体字段的对象键将被忽略(有关另一种方式，请参见 Decoder.DisallowUnknownFields)。

​	要将 JSON 反序列化为接口值，Unmarshal 将以下其中之一存储在接口值中： 

对于 JSON 布尔值，为 bool 类型；
对于 JSON 数字，为 float64 类型；
对于 JSON 字符串，为 string 类型；
对于 JSON 数组，为 []interface{} 类型；
对于 JSON 对象，为 map[string]interface{} 类型；
对于 JSON null，为 nil 类型。

​	要将 JSON 数组反序列化为切片，Unmarshal 将切片长度重置为零，然后将每个元素附加到切片中。作为一个特例，要将空 JSON 数组反序列化为切片，Unmarshal 将该切片替换为一个新的空切片。

​	将JSON数组解码到Go数组中，Unmarshal将JSON数组元素解码为对应的Go数组元素。如果Go数组小于JSON数组，则将多余的JSON数组元素丢弃。如果JSON数组小于Go数组，则将多余的Go数组元素设置为零值。

​	将JSON对象解码为映射时，Unmarshal首先创建一个映射以使用。如果映射为nil，则Unmarshal将分配一个新的映射。否则，Unmarshal将重用现有的映射，保留现有的条目。然后，Unmarshal将JSON对象中的键值对存储到映射中。映射的键类型必须是任何字符串类型、整数、实现json.Unmarshaler或实现encoding.TextUnmarshaler。

​	如果JSON编码的数据包含语法错误，则Unmarshal返回一个SyntaxError。

​	如果JSON值不适合给定的目标类型，或者JSON数字溢出了目标类型，则Unmarshal跳过该字段，并尽可能完成解组。如果没有遇到更严重的错误，Unmarshal将返回一个UnmarshalTypeError，描述最早出现的这种错误。在任何情况下，不能保证在问题字段后面的所有剩余字段都会解组到目标对象中。

​	JSON null值通过将Go值设置为nil来解组到接口、映射、指针或切片中。因为null通常在JSON中用于表示"不存在"，将JSON null解组成任何其他Go类型对该值没有影响，并且不会产生错误。

​	解组引用字符串时，无效的UTF-8或无效的UTF-16代理对不会被视为错误。相反，它们将被替换为Unicode替换字符U+FFFD。

##### Unmarshal Example
``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
Output:

[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
```

#### func [Valid](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/scanner.go;l=22)  <- go1.9

``` go linenums="1"
func Valid(data []byte) bool
```

​	Valid 函数判断数据 data 是否是有效的 JSON 编码。

##### Valid Example
``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))
}
Output:

true false
```

## 类型

### type [Decoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=14) 

``` go linenums="1"
type Decoder struct {
	// contains filtered or unexported fields
}
```

​	Decoder 从输入流中读取和解码 JSON 值。

##### Example

​	这个例子使用了 Decoder 来解码一系列不同的 JSON 值流。

``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

```

#### func [NewDecoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=31) 

``` go linenums="1"
func NewDecoder(r io.Reader) *Decoder
```

​	NewDecoder函数返回一个从 r 读取数据的新的 Decoder。

​	Decoder 引入了自己的缓冲，可能会从 r 中读取超出请求的 JSON 值的数据。

#### (*Decoder) [Buffered](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=83)  <- go1.1

``` go linenums="1"
func (dec *Decoder) Buffered() io.Reader
```

​	Buffered方法返回 Decoder 缓冲区中剩余的数据的 reader。该 reader 在下次调用 Decode 之前有效。

#### (*Decoder) [Decode](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=49) 

``` go linenums="1"
func (dec *Decoder) Decode(v any) error
```

​	Decode方法从输入中读取下一个 JSON 编码的值并将其存储在 v 所指向的值中。

​	有关 JSON 转换为 Go 值的详细信息，请参见 Unmarshal 的文档。

##### Decode Example

​	这个示例使用解码器(Decoder)来解码一个 JSON 对象流数组。

``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

}
Output:

json.Delim: [
Ed: Knock knock.
Sam: Who's there?
Ed: Go fmt.
Sam: Go fmt who?
Ed: Go fmt yourself!
json.Delim: ]
```

#### (*Decoder) [DisallowUnknownFields](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=42)  <- go1.10

``` go linenums="1"
func (dec *Decoder) DisallowUnknownFields()
```

​	DisallowUnknownFields方法会导致在目标对象为结构体且输入包含在目标中不存在的、未被忽略的公开字段的对象键时，解码器返回一个错误。

#### (*Decoder) [InputOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=513)  <- go1.14

``` go linenums="1"
func (dec *Decoder) InputOffset() int64
```

​	InputOffset方法返回当前解码器位置的输入流字节偏移量。该偏移量给出了最近返回的标记的结束位置和下一个标记的开始位置。

#### (*Decoder) [More](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=486)  <- go1.5

``` go linenums="1"
func (dec *Decoder) More() bool
```

​	More方法报告当前正在解析的数组或对象中是否还有另一个元素。

#### (*Decoder) [Token](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=371)  <- go1.5

``` go linenums="1"
func (dec *Decoder) Token() (Token, error)
```

​	Token方法返回输入流中的下一个JSON标记。在输入流的末尾，Token方法返回nil和io.EOF。

​	Token方法保证它返回的分隔符[ ] {}是正确嵌套和匹配的：如果Token方法在输入中遇到意外的分隔符，则它将返回一个错误。

​	输入流由基本的JSON值——布尔值、字符串、数字和null——以及类型为Delim的[ ] {}分隔符组成，用于标记数组和对象的开始和结束。逗号和冒号被省略。

##### Token Example

​	这个例子使用 Decoder 来解码一个由不同的 JSON 值组成的流。

``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
}
Output:

json.Delim: { (more)
string: Message (more)
string: Hello (more)
string: Array (more)
json.Delim: [ (more)
float64: 1 (more)
float64: 2 (more)
float64: 3
json.Delim: ] (more)
string: Null (more)
<nil>: <nil> (more)
string: Number (more)
float64: 1.234
json.Delim: }
```

#### (*Decoder) [UseNumber](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=37)  <- go1.1

``` go linenums="1"
func (dec *Decoder) UseNumber()
```

​	UseNumber方法会导致解码器将数字解组为Number类型的接口{}，而不是解组为float64。

### type [Delim](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=354)  <- go1.5

``` go linenums="1"
type Delim rune
```

​	Delim是JSON数组或对象分隔符，其中之一是[ ] { 或 }。

#### (Delim) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=356)  <- go1.5

``` go linenums="1"
func (d Delim) String() string
```

### type [Encoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=181) 

``` go linenums="1"
type Encoder struct {
	// contains filtered or unexported fields
}
```

​	Encoder将JSON值写入输出流。

#### func [NewEncoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=192) 

``` go linenums="1"
func NewEncoder(w io.Writer) *Encoder
```

​	NewEncoder 函数返回一个新的编码器，它将写入到 w。

#### (*Encoder) [Encode](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=201) 

``` go linenums="1"
func (enc *Encoder) Encode(v any) error
```

​	Encode方法将 v 的 JSON 编码写入流中，随后加上一个换行符。

​	有关将 Go 值转换为 JSON 的详细信息，请参见 Marshal 的文档。

#### (*Encoder) [SetEscapeHTML](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=255)  <- go1.7

``` go linenums="1"
func (enc *Encoder) SetEscapeHTML(on bool)
```

​	SetEscapeHTML方法指定是否在 JSON 引用字符串内部转义有问题的 HTML 字符。默认行为是将 `&`、`<` 和 `>` 转义为 `\u0026`、`\u003c` 和 `\u003e`，以避免在 HTML 中嵌入 JSON 时出现某些安全问题。

​	在非 HTML 环境中，当转义干扰输出的可读性时，SetEscapeHTML(false) 将禁用此行为。

#### (*Encoder) [SetIndent](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=243)  <- go1.7

``` go linenums="1"
func (enc *Encoder) SetIndent(prefix, indent string)
```

​	SetIndent方法指示编码器将每个后续编码值格式化为由 package-level 函数 Indent(dst、src、prefix、indent) 缩进的形式。调用 SetIndent("", "") 禁用缩进。

### type [InvalidUnmarshalError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=156) 

``` go linenums="1"
type InvalidUnmarshalError struct {
	Type reflect.Type
}
```

​	InvalidUnmarshalError 描述传递给 Unmarshal 的无效参数。(传递给 Unmarshal 的参数必须是非 nil 指针。)

#### (*InvalidUnmarshalError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=160) 

``` go linenums="1"
func (e *InvalidUnmarshalError) Error() string
```

### type [Marshaler](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=223) 

``` go linenums="1"
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}
```

​	Marshaler 是一种类型的接口，这种类型可以将自身编组为有效的 JSON。

### type [MarshalerError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=263) 

``` go linenums="1"
type MarshalerError struct {
	Type reflect.Type
	Err  error
	// contains filtered or unexported fields
}
```

​	MarshalerError 表示调用 MarshalJSON 或 MarshalText 方法时的错误。

#### (*MarshalerError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=269) 

``` go linenums="1"
func (e *MarshalerError) Error() string
```

#### (*MarshalerError) [Unwrap](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=280)  <- go1.13

``` go linenums="1"
func (e *MarshalerError) Unwrap() error
```

​	Unwrap方法返回基本的错误。

### type [Number](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=189)  <- go1.1

``` go linenums="1"
type Number string
```

​	Number类型表示JSON数值字面。

#### (Number) [Float64](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=195)  <- go1.1

``` go linenums="1"
func (n Number) Float64() (float64, error)
```

​	Float64方法将数值转换为float64。

#### (Number) [Int64](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=200)  <- go1.1

``` go linenums="1"
func (n Number) Int64() (int64, error)
```

​	Int64方法将数值转换为int64。

#### (Number) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=192)  <- go1.1

``` go linenums="1"
func (n Number) String() string
```

​	String方法返回数值的文字表示形式。

### type [RawMessage](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=262) 

``` go linenums="1"
type RawMessage []byte
```

​	RawMessage是一个原始编码的JSON值。它实现了Marshaler和Unmarshaler接口，可以用于延迟JSON解码或预先计算JSON编码。

##### Example(Marshal)

​	这个示例使用 RawMessage 来在编组时使用预计算的 JSON。

``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}
Output:

{
	"header": {
		"precomputed": true
	},
	"body": "Hello Gophers!"
}
```

##### Example (Unmarshal) 

​	这个示例使用 RawMessage 来延迟解析 JSON 消息的一部分。

``` go linenums="1"
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
	{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
	{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for _, c := range colors {
		var dst any
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}
Output:

YCbCr &{255 0 -10}
RGB &{98 218 255}
```

#### (RawMessage) [MarshalJSON](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=265) 

``` go linenums="1"
func (m RawMessage) MarshalJSON() ([]byte, error)
```

​	MarshalJSON方法返回m的JSON编码。

#### (*RawMessage) [UnmarshalJSON](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=273) 

``` go linenums="1"
func (m *RawMessage) UnmarshalJSON(data []byte) error
```

​	UnmarshalJSON方法将data的一个副本设置为`*m`。

### type [SyntaxError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/scanner.go;l=47) 

``` go linenums="1"
type SyntaxError struct {
	Offset int64 // error occurred after reading Offset bytes // 读取Offset bytes后发生错误
	// contains filtered or unexported fields
}
```

​	SyntaxError描述了一个JSON语法错误。如果JSON不能被解析，Unmarshal函数将返回SyntaxError。

#### (*SyntaxError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/scanner.go;l=52) 

``` go linenums="1"
func (e *SyntaxError) Error() string
```

### type [Token](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/stream.go;l=292)  <- go1.5

``` go linenums="1"
type Token any
```

Token保存以下类型之一的值：

```
Delim, for the four JSON delimiters [ ] { }
bool, for JSON booleans
float64, for JSON numbers
Number, for JSON numbers
string, for JSON string literals
nil, for JSON null
```

### type [UnmarshalTypeError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=125) 

``` go linenums="1"
type UnmarshalTypeError struct {
	Value  string       // JSON值的描述 - "bool"、"array"、"number -5"
	Type   reflect.Type // 无法赋值给Go值的类型
	Offset int64        // 发生错误后读取的字节数
	Struct string       // 包含该字段的结构体类型的名称
	Field  string       // 从根节点到字段的完整路径
}
```

​	UnmarshalTypeError描述了一个不适合特定Go类型值的JSON值。

#### (*UnmarshalTypeError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=133) 

``` go linenums="1"
func (e *UnmarshalTypeError) Error() string
```

### type [Unmarshaler](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/decode.go;l=119) 

``` go linenums="1"
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
```

​	Unmarshaler是由可以解码JSON描述的类型实现的接口。可以假定输入是JSON值的有效编码。如果想在返回后保留数据，UnmarshalJSON必须复制JSON数据。

​	按照惯例，为了近似于Unmarshal本身的行为，Unmarshalers将UnmarshalJSON([]byte("null"))实现为无操作。

### type [UnsupportedTypeError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=229) 

``` go linenums="1"
type UnsupportedTypeError struct {
	Type reflect.Type
}
```

​	在尝试对不受支持的值类型进行编码时，Marshal返回UnsupportedTypeError。

#### (*UnsupportedTypeError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=233) 

``` go linenums="1"
func (e *UnsupportedTypeError) Error() string
```

### type [UnsupportedValueError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=239) 

``` go linenums="1"
type UnsupportedValueError struct {
	Value reflect.Value
	Str   string
}
```

​	在尝试对不受支持的值进行编码时，Marshal返回UnsupportedValueError。

#### (*UnsupportedValueError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/json/encode.go;l=244) 

``` go linenums="1"
func (e *UnsupportedValueError) Error() string
```