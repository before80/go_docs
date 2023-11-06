+++
title = "json"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
https://pkg.go.dev/encoding/json@go1.20.1

Package json implements encoding and decoding of JSON as defined in [RFC 7159](https://rfc-editor.org/rfc/rfc7159.html). The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.

​	json包实现了[RFC 7159](https://rfc-editor.org/rfc/rfc7159.html)中定义的JSON编解码。JSON和Go值之间的映射在Marshal和Unmarshal函数的文档中有描述。

See "JSON and Go" for an introduction to this package: https://golang.org/doc/articles/json_and_go.html

​	有关此包的介绍，请参见[JSON和Go]({{< ref "/goBlog/2011/JSONAndGo">}})。

## Example (CustomMarshalJSON)
``` go 
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

## Example(TextMarshalJSON)
``` go 
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

### func Compact 

``` go 
func Compact(dst *bytes.Buffer, src []byte) error
```

Compact appends to dst the JSON-encoded src with insignificant space characters elided.

​	Compact函数将JSON编码的src附加到dst中，省略了不重要的空格字符。

### func HTMLEscape 

``` go 
func HTMLEscape(dst *bytes.Buffer, src []byte)
```

HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029 characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029 so that the JSON will be safe to embed inside HTML <script> tags. For historical reasons, web browsers don't honor standard HTML escaping within <script> tags, so an alternative JSON encoding must be used.

​	HTMLEscape函数将JSON编码的src附加到dst中，将字符串文字内的<、>、&、U+2028和U+2029字符更改为\u003c、\u003e、\u0026、\u2028、\u2029，以使JSON可以安全地嵌入HTML `<script>`标记中。由于历史原因，Web浏览器不支持在`<script>`标记中使用标准的HTML转义，因此必须使用替代的JSON编码。

#### HTMLEscape Example
``` go 
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

### func Indent 

``` go 
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
```

Indent appends to dst an indented form of the JSON-encoded src. Each element in a JSON object or array begins on a new, indented line beginning with prefix followed by one or more copies of indent according to the indentation nesting. The data appended to dst does not begin with the prefix nor any indentation, to make it easier to embed inside other formatted JSON data. Although leading space characters (space, tab, carriage return, newline) at the beginning of src are dropped, trailing space characters at the end of src are preserved and copied to dst. For example, if src has no trailing spaces, neither will dst; if src ends in a trailing newline, so will dst.

​	Indent函数将JSON编码的src的缩进形式附加到dst中。JSON对象或数组中的每个元素都在新的缩进行上开始，该行以prefix开头，后跟一个或多个indent的副本，具体取决于缩进嵌套。附加到dst的数据不以prefix或任何缩进开始，以使其更容易嵌入其他格式化的JSON数据中。虽然src开头的前导空格字符(空格、制表符、回车、换行符)会被删除，但src末尾的尾随空格字符会被保留并复制到dst中。例如，如果src没有尾随空格，则dst也没有；如果src以尾随换行符结束，则dst也是如此。

#### Indent Example
``` go 
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

### func Marshal 

``` go 
func Marshal(v any) ([]byte, error)
```

Marshal returns the JSON encoding of v.

​	Marshal函数返回v的JSON编码。

Marshal traverses the value v recursively. If an encountered value implements the Marshaler interface and is not a nil pointer, Marshal calls its MarshalJSON method to produce JSON. If no MarshalJSON method is present but the value implements encoding.TextMarshaler instead, Marshal calls its MarshalText method and encodes the result as a JSON string. The nil pointer exception is not strictly necessary but mimics a similar, necessary exception in the behavior of UnmarshalJSON.

​	Marshal函数递归遍历值v。如果遇到的值实现了Marshaler接口并且不是nil指针，则Marshal调用其MarshalJSON方法以生成JSON。如果没有MarshalJSON方法但该值代替实现了encoding.TextMarshaler，则Marshal调用其MarshalText方法并将结果编码为JSON字符串。nil指针异常并不是严格必要的，但模仿了在UnmarshalJSON的行为中必要的类似异常。

Otherwise, Marshal uses the following type-dependent default encodings:

​	否则，Marshal使用以下类型相关的默认编码：

Boolean values encode as JSON booleans.

​	布尔值被编码为JSON布尔值。

Floating point, integer, and Number values encode as JSON numbers.

​	浮点数、整数和Number类型的值被编码为JSON数字。

String values encode as JSON strings coerced to valid UTF-8, replacing invalid bytes with the Unicode replacement rune. So that the JSON will be safe to embed inside HTML <script> tags, the string is encoded using HTMLEscape, which replaces "<", ">", "&", U+2028, and U+2029 are escaped to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029". This replacement can be disabled when using an Encoder, by calling SetEscapeHTML(false).

​	字符串类型的值被编码为JSON字符串，被强制转换为有效的UTF-8编码，无效的字节会被替换为Unicode替换符。为了将JSON安全地嵌入到HTML的`<script>`标签中，使用HTMLEscape对字符串进行编码，将"`<`", "`>`", "`&`", U+2028, 和 U+2029转义为"\u003c", "\u003e", "\u0026", "\u2028"和"\u2029"。使用Encoder时，可以通过调用SetEscapeHTML(false)来禁用此替换。

Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON value.

​	数组和切片类型的值被编码为JSON数组，但[]byte类型的值会被编码为base64编码的字符串，而nil切片类型的值会被编码为null。

Struct values encode as JSON objects. Each exported struct field becomes a member of the object, using the field name as the object key, unless the field is omitted for one of the reasons given below.

​	结构体类型的值被编码为JSON对象。每个导出的结构体字段都成为对象的一个成员，使用字段名作为对象键，除非由于以下原因之一而省略该字段。

The encoding of each struct field can be customized by the format string stored under the "json" key in the struct field's tag. The format string gives the name of the field, possibly followed by a comma-separated list of options. The name may be empty in order to specify options without overriding the default field name.

​	每个结构体字段的编码可以通过存储在该字段标记的"json"键下的格式字符串进行自定义。格式字符串给出字段名，可能后跟一个用逗号分隔的选项列表。如果不想覆盖默认的字段名，则字段名可以为空。

The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value, defined as false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.

​	"omitempty"选项指定如果该字段具有空值(定义为false、0、nil指针、nil接口值以及任何空数组、切片、映射或字符串)，则应该从编码中省略该字段。

As a special case, if the field tag is "-", the field is always omitted. Note that a field with name "-" can still be generated using the tag "-,".

​	作为特殊情况，如果字段标记为"-"，则始终省略该字段。请注意，具有名称"-"的字段仍然可以使用标记"-,"生成。

Examples of struct field tags and their meanings:

​	以下是结构体字段标记及其含义的示例：

```
// Field appears in JSON as key "myName".
// Field 出现在JSON中作为键"myName"。
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
// Field 出现在JSON中作为键"myName"，
// 如果其值为空，将从对象中省略该字段。
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
// Field 出现在JSON中作为键"Field"(默认值)，
// 但如果为空则跳过该字段。
// 请注意前导逗号。
Field int `json:",omitempty"`

// Field is ignored by this package.
// 此包忽略 Field。
Field int `json:"-"`

// Field appears in JSON as key "-".
// Field 出现在JSON中作为键"-"。
Field int `json:"-,"`
```

The "string" option signals that a field is stored as JSON inside a JSON-encoded string. It applies only to fields of string, floating point, integer, or boolean types. This extra level of encoding is sometimes used when communicating with JavaScript programs:

​	"string"选项表示字段存储在JSON编码的字符串中。它仅适用于字符串、浮点数、整数或布尔类型的字段。在与JavaScript程序通信时，有时会使用这种额外的编码级别：

```
Int64String int64 `json:",string"`
```

The key name will be used if it's a non-empty string consisting of only Unicode letters, digits, and ASCII punctuation except quotation marks, backslash, and comma.

​	如果键名是由Unicode字母、数字和ASCII标点(引号、反斜杠和逗号除外)组成的非空字符串，则将使用该键名。

Anonymous struct fields are usually marshaled as if their inner exported fields were fields in the outer struct, subject to the usual Go visibility rules amended as described in the next paragraph. An anonymous struct field with a name given in its JSON tag is treated as having that name, rather than being anonymous. An anonymous struct field of interface type is treated the same as having that type as its name, rather than being anonymous.

​	匿名结构字段通常被解释为其内部导出字段在外部结构体中的字段，但受到下一段中描述的修饰性可见性规则的影响。具有在其JSON标签中给出的名称的匿名结构字段将被视为具有该名称，而不是匿名的。接口类型的匿名结构字段将与其类型相同，而不是匿名的。

The Go visibility rules for struct fields are amended for JSON when deciding which field to marshal or unmarshal. If there are multiple fields at the same level, and that level is the least nested (and would therefore be the nesting level selected by the usual Go rules), the following extra rules apply:

​	当决定哪个字段进行编组或取消编组时，针对JSON修饰性可见性规则修正结构字段的Go可见性规则。如果同一级别存在多个字段，并且该级别是最不嵌套的(因此将是常规Go规则选择的嵌套级别)，则以下额外规则适用：

1) Of those fields, if any are JSON-tagged, only tagged fields are considered, even if there are multiple untagged fields that would otherwise conflict.
2) 对于这些字段，如果有任何字段具有JSON标记，则只考虑带有标记的字段，即使有多个未标记的字段也会冲突。
3)  If there is exactly one field (tagged or not according to the first rule), that is selected.
4) 如果恰好有一个字段(根据第一条规则是否带有标记)，则选择该字段。
5) Otherwise there are multiple fields, and all are ignored; no error occurs.
6) 否则将忽略所有字段；不会发生错误。

Handling of anonymous struct fields is new in Go 1.1. Prior to Go 1.1, anonymous struct fields were ignored. To force ignoring of an anonymous struct field in both current and earlier versions, give the field a JSON tag of "-".

​	对匿名结构字段的处理是Go 1.1的新内容。在Go 1.1之前，匿名结构字段被忽略。要在当前版本和早期版本中强制忽略匿名结构字段，请给该字段一个"-"的JSON标签。

​	在Go 1.1中，对匿名结构字段的处理是新的。在Go 1.1之前，匿名结构字段被忽略。要在当前版本和早期版本中强制忽略匿名结构字段，请给该字段一个JSON标记"-"。

Map values encode as JSON objects. The map's key type must either be a string, an integer type, or implement encoding.TextMarshaler. The map keys are sorted and used as JSON object keys by applying the following rules, subject to the UTF-8 coercion described for string values above:

​	Map值编码为JSON对象。映射的键类型必须是字符串、整数类型或实现encoding.TextMarshaler。映射键按以下规则排序并用作JSON对象键，但要受上述字符串值的UTF-8强制规定：

- keys of any string type are used directly

- 任何字符串类型的键都直接使用。 

- encoding.TextMarshalers are marshaled

- encoding.TextMarshalers已被解组。 

- integer keys are converted to strings

- 整数键被转换为字符串。 

Pointer values encode as the value pointed to. A nil pointer encodes as the null JSON value.

​	指针值编码为指向的值。空指针编码为null JSON值。指针值编码为所指向的值。一个nil指针编码为空的JSON值。

Interface values encode as the value contained in the interface. A nil interface value encodes as the null JSON value.

​	接口值编码为接口中包含的值。空接口值编码为null JSON值。

Channel, complex, and function values cannot be encoded in JSON. Attempting to encode such a value causes Marshal to return an UnsupportedTypeError.

​	通道、复数和函数值不能编码为 JSON。尝试对此类值进行编码将导致 Marshal 返回 UnsupportedTypeError。

JSON cannot represent cyclic data structures and Marshal does not handle them. Passing cyclic structures to Marshal will result in an error.

​	JSON 不能表示循环数据结构，Marshal 也无法处理它们。将循环结构传递给 Marshal 将导致错误。

#### Marshal Example
``` go 
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

### func MarshalIndent 

``` go 
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
```

MarshalIndent is like Marshal but applies Indent to format the output. Each JSON element in the output will begin on a new line beginning with prefix followed by one or more copies of indent according to the indentation nesting.

​	MarshalIndent函数类似于 Marshal函数，但应用 Indent 以格式化输出。输出中的每个 JSON 元素都将在新行上开始，以前缀开头，后跟一个或多个 indent 副本，具体取决于缩进嵌套。

#### MarshalIndent Example
``` go 
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

### func Unmarshal 

``` go 
func Unmarshal(data []byte, v any) error
```

Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.

​	Unmarshal函数解析 JSON 编码的数据并将结果存储在 v 所指向的值中。如果 v 为 nil 或不是指针，则 Unmarshal 返回 InvalidUnmarshalError。

Unmarshal uses the inverse of the encodings that Marshal uses, allocating maps, slices, and pointers as necessary, with the following additional rules:

​	Unmarshal函数使用 Marshal函数所使用的编码的反向方式，根据以下附加规则分配映射、切片和指针：

To unmarshal JSON into a pointer, Unmarshal first handles the case of the JSON being the JSON literal null. In that case, Unmarshal sets the pointer to nil. Otherwise, Unmarshal unmarshals the JSON into the value pointed at by the pointer. If the pointer is nil, Unmarshal allocates a new value for it to point to.

​	要将 JSON 反序列化为指针，Unmarshal 首先处理 JSON 为 JSON 文字 null 的情况。在这种情况下，Unmarshal 将指针设置为 nil。否则，Unmarshal 将 JSON 反序列化为指针所指向的值。如果指针为 nil，则 Unmarshal 为其分配一个新值。

To unmarshal JSON into a value implementing the Unmarshaler interface, Unmarshal calls that value's UnmarshalJSON method, including when the input is a JSON null. Otherwise, if the value implements encoding.TextUnmarshaler and the input is a JSON quoted string, Unmarshal calls that value's UnmarshalText method with the unquoted form of the string.

​	要将 JSON 反序列化为实现 Unmarshaler 接口的值，Unmarshal函数调用该值的 UnmarshalJSON 方法，包括输入为 JSON null 的情况。否则，如果该值实现 encoding.TextUnmarshaler 并且输入是 JSON 引号字符串，则 Unmarshal 会使用字符串的未引用形式调用该值的 UnmarshalText 方法。

To unmarshal JSON into a struct, Unmarshal matches incoming object keys to the keys used by Marshal (either the struct field name or its tag), preferring an exact match but also accepting a case-insensitive match. By default, object keys which don't have a corresponding struct field are ignored (see Decoder.DisallowUnknownFields for an alternative).

​	要将 JSON 反序列化为结构体，Unmarshal 将传入的对象键与 Marshal 使用的键进行匹配(即结构字段名或其标记)，优先选择完全匹配，但也接受大小写不敏感的匹配。默认情况下，没有对应结构体字段的对象键将被忽略(有关另一种方式，请参见 Decoder.DisallowUnknownFields)。

To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:

​	要将 JSON 反序列化为接口值，Unmarshal 将以下其中之一存储在接口值中： 

```
bool, for JSON booleans
对于 JSON 布尔值，为 bool 类型；
float64, for JSON numbers
对于 JSON 数字，为 float64 类型；
string, for JSON strings
对于 JSON 字符串，为 string 类型；
[]interface{}, for JSON arrays
对于 JSON 数组，为 []interface{} 类型；
map[string]interface{}, for JSON objects
对于 JSON 对象，为 map[string]interface{} 类型；
nil for JSON null
对于 JSON null，为 nil 类型。
```

To unmarshal a JSON array into a slice, Unmarshal resets the slice length to zero and then appends each element to the slice. As a special case, to unmarshal an empty JSON array into a slice, Unmarshal replaces the slice with a new empty slice.

​	要将 JSON 数组反序列化为切片，Unmarshal 将切片长度重置为零，然后将每个元素附加到切片中。作为一个特例，要将空 JSON 数组反序列化为切片，Unmarshal 将该切片替换为一个新的空切片。

To unmarshal a JSON array into a Go array, Unmarshal decodes JSON array elements into corresponding Go array elements. If the Go array is smaller than the JSON array, the additional JSON array elements are discarded. If the JSON array is smaller than the Go array, the additional Go array elements are set to zero values.

​	将JSON数组解码到Go数组中，Unmarshal将JSON数组元素解码为对应的Go数组元素。如果Go数组小于JSON数组，则将多余的JSON数组元素丢弃。如果JSON数组小于Go数组，则将多余的Go数组元素设置为零值。

To unmarshal a JSON object into a map, Unmarshal first establishes a map to use. If the map is nil, Unmarshal allocates a new map. Otherwise Unmarshal reuses the existing map, keeping existing entries. Unmarshal then stores key-value pairs from the JSON object into the map. The map's key type must either be any string type, an integer, implement json.Unmarshaler, or implement encoding.TextUnmarshaler.

​	将JSON对象解码为映射时，Unmarshal首先创建一个映射以使用。如果映射为nil，则Unmarshal将分配一个新的映射。否则，Unmarshal将重用现有的映射，保留现有的条目。然后，Unmarshal将JSON对象中的键值对存储到映射中。映射的键类型必须是任何字符串类型、整数、实现json.Unmarshaler或实现encoding.TextUnmarshaler。

If the JSON-encoded data contain a syntax error, Unmarshal returns a SyntaxError.

​	如果JSON编码的数据包含语法错误，则Unmarshal返回一个SyntaxError。

If a JSON value is not appropriate for a given target type, or if a JSON number overflows the target type, Unmarshal skips that field and completes the unmarshaling as best it can. If no more serious errors are encountered, Unmarshal returns an UnmarshalTypeError describing the earliest such error. In any case, it's not guaranteed that all the remaining fields following the problematic one will be unmarshaled into the target object.

​	如果JSON值不适合给定的目标类型，或者JSON数字溢出了目标类型，则Unmarshal跳过该字段，并尽可能完成解组。如果没有遇到更严重的错误，Unmarshal将返回一个UnmarshalTypeError，描述最早出现的这种错误。在任何情况下，不能保证在问题字段后面的所有剩余字段都会解组到目标对象中。

The JSON null value unmarshals into an interface, map, pointer, or slice by setting that Go value to nil. Because null is often used in JSON to mean “not present,” unmarshaling a JSON null into any other Go type has no effect on the value and produces no error.

​	JSON null值通过将Go值设置为nil来解组到接口、映射、指针或切片中。因为null通常在JSON中用于表示"不存在"，将JSON null解组成任何其他Go类型对该值没有影响，并且不会产生错误。

When unmarshaling quoted strings, invalid UTF-8 or invalid UTF-16 surrogate pairs are not treated as an error. Instead, they are replaced by the Unicode replacement character U+FFFD.

​	解组引用字符串时，无效的UTF-8或无效的UTF-16代理对不会被视为错误。相反，它们将被替换为Unicode替换字符U+FFFD。

#### Unmarshal Example
``` go 
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

### func Valid  <- go1.9

``` go 
func Valid(data []byte) bool
```

​	Valid 函数判断数据 data 是否是有效的 JSON 编码。

#### Valid Example
``` go 
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

### type Decoder 

``` go 
type Decoder struct {
	// contains filtered or unexported fields
}
```

Valid reports whether data is a valid JSON encoding.

​	Decoder 从输入流中读取和解码 JSON 值。

#### Example

A Decoder reads and decodes JSON values from an input stream.

​	这个例子使用了 Decoder 来解码一系列不同的 JSON 值流。

``` go 
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

#### func NewDecoder 

``` go 
func NewDecoder(r io.Reader) *Decoder
```

NewDecoder returns a new decoder that reads from r.

​	NewDecoder函数返回一个从 r 读取数据的新的 Decoder。

The decoder introduces its own buffering and may read data from r beyond the JSON values requested.

​	Decoder 引入了自己的缓冲，可能会从 r 中读取超出请求的 JSON 值的数据。

#### (*Decoder) Buffered  <- go1.1

``` go 
func (dec *Decoder) Buffered() io.Reader
```

Buffered returns a reader of the data remaining in the Decoder's buffer. The reader is valid until the next call to Decode.

​	Buffered方法返回 Decoder 缓冲区中剩余的数据的 reader。该 reader 在下次调用 Decode 之前有效。

#### (*Decoder) Decode 

``` go 
func (dec *Decoder) Decode(v any) error
```

Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.

​	Decode方法从输入中读取下一个 JSON 编码的值并将其存储在 v 所指向的值中。

See the documentation for Unmarshal for details about the conversion of JSON into a Go value.

​	有关 JSON 转换为 Go 值的详细信息，请参见 Unmarshal 的文档。

##### Decode Example

This example uses a Decoder to decode a streaming array of JSON objects.

​	这个示例使用解码器(Decoder)来解码一个 JSON 对象流数组。

``` go 
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

#### (*Decoder) DisallowUnknownFields  <- go1.10

``` go 
func (dec *Decoder) DisallowUnknownFields()
```

DisallowUnknownFields causes the Decoder to return an error when the destination is a struct and the input contains object keys which do not match any non-ignored, exported fields in the destination.

​	DisallowUnknownFields方法会导致在目标对象为结构体且输入包含在目标中不存在的、未被忽略的公开字段的对象键时，解码器返回一个错误。

#### (*Decoder) InputOffset  <- go1.14

``` go 
func (dec *Decoder) InputOffset() int64
```

InputOffset returns the input stream byte offset of the current decoder position. The offset gives the location of the end of the most recently returned token and the beginning of the next token.

​	InputOffset方法返回当前解码器位置的输入流字节偏移量。该偏移量给出了最近返回的标记的结束位置和下一个标记的开始位置。

#### (*Decoder) More  <- go1.5

``` go 
func (dec *Decoder) More() bool
```

More reports whether there is another element in the current array or object being parsed.

​	More方法报告当前正在解析的数组或对象中是否还有另一个元素。

#### (*Decoder) Token  <- go1.5

``` go 
func (dec *Decoder) Token() (Token, error)
```

Token returns the next JSON token in the input stream. At the end of the input stream, Token returns nil, io.EOF.

​	Token方法返回输入流中的下一个JSON标记。在输入流的末尾，Token方法返回nil和io.EOF。

Token guarantees that the delimiters [ ] { } it returns are properly nested and matched: if Token encounters an unexpected delimiter in the input, it will return an error.

​	Token方法保证它返回的分隔符[ ] {}是正确嵌套和匹配的：如果Token方法在输入中遇到意外的分隔符，则它将返回一个错误。

The input stream consists of basic JSON values—bool, string, number, and null—along with delimiters [ ] { } of type Delim to mark the start and end of arrays and objects. Commas and colons are elided.

​	输入流由基本的JSON值——布尔值、字符串、数字和null——以及类型为Delim的[ ] {}分隔符组成，用于标记数组和对象的开始和结束。逗号和冒号被省略。

##### Token Example

This example uses a Decoder to decode a stream of distinct JSON values.

​	这个例子使用 Decoder 来解码一个由不同的 JSON 值组成的流。

``` go 
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

#### (*Decoder) UseNumber  <- go1.1

``` go 
func (dec *Decoder) UseNumber()
```

UseNumber causes the Decoder to unmarshal a number into an interface{} as a Number instead of as a float64.

​	UseNumber方法会导致解码器将数字解组为Number类型的interface{}，而不是解组为float64。

### type Delim  <- go1.5

``` go 
type Delim rune
```

A Delim is a JSON array or object delimiter, one of [ ] { or }.

​	Delim是JSON数组或对象分隔符，其中之一是[ ] { 或 }。

#### (Delim) String  <- go1.5

``` go 
func (d Delim) String() string
```

### type Encoder 

``` go 
type Encoder struct {
	// contains filtered or unexported fields
}
```

An Encoder writes JSON values to an output stream.

​	Encoder将JSON值写入输出流。

#### func NewEncoder 

``` go 
func NewEncoder(w io.Writer) *Encoder
```

NewEncoder returns a new encoder that writes to w.

​	NewEncoder 函数返回一个新的编码器，它将写入到 w。

#### (*Encoder) Encode 

``` go 
func (enc *Encoder) Encode(v any) error
```

Encode writes the JSON encoding of v to the stream, followed by a newline character.

​	Encode方法将 v 的 JSON 编码写入流中，随后加上一个换行符。

See the documentation for Marshal for details about the conversion of Go values to JSON.

​	有关将 Go 值转换为 JSON 的详细信息，请参见 Marshal 的文档。

#### (*Encoder) SetEscapeHTML  <- go1.7

``` go 
func (enc *Encoder) SetEscapeHTML(on bool)
```

SetEscapeHTML specifies whether problematic HTML characters should be escaped inside JSON quoted strings. The default behavior is to escape &, <, and > to \u0026, \u003c, and \u003e to avoid certain safety problems that can arise when embedding JSON in HTML.

​	SetEscapeHTML方法指定是否在 JSON 引用字符串内部转义有问题的 HTML 字符。默认行为是将 `&`、`<` 和 `>` 转义为 `\u0026`、`\u003c` 和 `\u003e`，以避免在 HTML 中嵌入 JSON 时出现某些安全问题。

In non-HTML settings where the escaping interferes with the readability of the output, SetEscapeHTML(false) disables this behavior.

​	在非 HTML 环境中，当转义干扰输出的可读性时，SetEscapeHTML(false) 将禁用此行为。

#### (*Encoder) SetIndent  <- go1.7

``` go 
func (enc *Encoder) SetIndent(prefix, indent string)
```

SetIndent instructs the encoder to format each subsequent encoded value as if indented by the package-level function Indent(dst, src, prefix, indent). Calling SetIndent("", "") disables indentation.

​	SetIndent方法指示编码器将每个后续编码值格式化为由 package-level 函数 Indent(dst、src、prefix、indent) 缩进的形式。调用 SetIndent("", "") 禁用缩进。

### type InvalidUTF8Error <-DEPRECATED

```go
type InvalidUTF8Error struct {
	S string // the whole string value that caused the error
}
```

Before Go 1.2, an InvalidUTF8Error was returned by Marshal when attempting to encode a string value with invalid UTF-8 sequences. As of Go 1.2, Marshal instead coerces the string to valid UTF-8 by replacing invalid bytes with the Unicode replacement rune U+FFFD.

Deprecated: No longer used; kept for compatibility.

#### (*InvalidUTF8Error) Error

```
func (e *InvalidUTF8Error) Error() string
```

### type InvalidUnmarshalError 

``` go 
type InvalidUnmarshalError struct {
	Type reflect.Type
}
```

An InvalidUnmarshalError describes an invalid argument passed to Unmarshal. (The argument to Unmarshal must be a non-nil pointer.)

​	InvalidUnmarshalError 描述传递给 Unmarshal 的无效参数。(传递给 Unmarshal 的参数必须是非 nil 指针。)

#### (*InvalidUnmarshalError) Error 

``` go 
func (e *InvalidUnmarshalError) Error() string
```

### type Marshaler 

``` go 
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}
```

Marshaler is the interface implemented by types that can marshal themselves into valid JSON.

​	Marshaler 是一种类型的接口，这种类型可以将自身编组为有效的 JSON。

### type MarshalerError 

``` go 
type MarshalerError struct {
	Type reflect.Type
	Err  error
	// contains filtered or unexported fields
}
```

A MarshalerError represents an error from calling a MarshalJSON or MarshalText method.

​	MarshalerError 表示调用 MarshalJSON 或 MarshalText 方法时的错误。

#### (*MarshalerError) Error 

``` go 
func (e *MarshalerError) Error() string
```

#### (*MarshalerError) Unwrap  <- go1.13

``` go 
func (e *MarshalerError) Unwrap() error
```

Unwrap returns the underlying error.

​	Unwrap方法返回基本的错误。

### type Number  <- go1.1

``` go 
type Number string
```

A Number represents a JSON number literal.

​	Number类型表示JSON数值字面。

#### (Number) Float64  <- go1.1

``` go 
func (n Number) Float64() (float64, error)
```

Float64 returns the number as a float64.

​	Float64方法将数值转换为float64。

#### (Number) Int64  <- go1.1

``` go 
func (n Number) Int64() (int64, error)
```

Int64 returns the number as an int64.

​	Int64方法将数值转换为int64。

#### (Number) String  <- go1.1

``` go 
func (n Number) String() string
```

String returns the literal text of the number.

​	String方法返回数值的文字表示形式。

### type RawMessage 

``` go 
type RawMessage []byte
```

RawMessage is a raw encoded JSON value. It implements Marshaler and Unmarshaler and can be used to delay JSON decoding or precompute a JSON encoding.

​	RawMessage是一个原始编码的JSON值。它实现了Marshaler和Unmarshaler接口，可以用于延迟JSON解码或预先计算JSON编码。

#### Example(Marshal)

This example uses RawMessage to use a precomputed JSON during marshal.

​	这个示例使用 RawMessage 来在编组时使用预计算的 JSON。

``` go 
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

#### Example (Unmarshal) 

This example uses RawMessage to delay parsing part of a JSON message.

​	这个示例使用 RawMessage 来延迟解析 JSON 消息的一部分。

``` go 
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

#### (RawMessage) MarshalJSON 

``` go 
func (m RawMessage) MarshalJSON() ([]byte, error)
```

MarshalJSON returns m as the JSON encoding of m.

​	MarshalJSON方法返回m的JSON编码。

#### (*RawMessage) UnmarshalJSON 

``` go 
func (m *RawMessage) UnmarshalJSON(data []byte) error
```

UnmarshalJSON sets *m to a copy of data.

​	UnmarshalJSON方法将data的一个副本设置为`*m`。

### type SyntaxError 

``` go 
type SyntaxError struct {
	Offset int64 // error occurred after reading Offset bytes // 读取Offset bytes后发生错误
	// contains filtered or unexported fields
}
```

A SyntaxError is a description of a JSON syntax error. Unmarshal will return a SyntaxError if the JSON can't be parsed.

​	SyntaxError描述了一个JSON语法错误。如果JSON不能被解析，Unmarshal函数将返回SyntaxError。

#### (*SyntaxError) Error 

``` go 
func (e *SyntaxError) Error() string
```

### type Token  <- go1.5

``` go 
type Token any
```

A Token holds a value of one of these types:

​	Token保存以下类型之一的值：

```
Delim, for the four JSON delimiters [ ] { }
bool, for JSON booleans
float64, for JSON numbers
Number, for JSON numbers
string, for JSON string literals
nil, for JSON null
```

### type UnmarshalFieldError <- DEPRECATED

```go
type UnmarshalFieldError struct {
	Key   string
	Type  reflect.Type
	Field reflect.StructField
}
```

An UnmarshalFieldError describes a JSON object key that led to an unexported (and therefore unwritable) struct field.

Deprecated: No longer used; kept for compatibility.

#### (*UnmarshalFieldError) Error

```go
func (e *UnmarshalFieldError) Error() string
```

### type UnmarshalTypeError 

``` go 
type UnmarshalTypeError struct {
	Value  string       // JSON值的描述 - "bool"、"array"、"number -5"
	Type   reflect.Type // 无法赋值给Go值的类型
	Offset int64        // 发生错误后读取的字节数
	Struct string       // 包含该字段的结构体类型的名称
	Field  string       // 从根节点到字段的完整路径
}
```

An UnmarshalTypeError describes a JSON value that was not appropriate for a value of a specific Go type.

​	UnmarshalTypeError描述了一个不适合特定Go类型值的JSON值。

#### (*UnmarshalTypeError) Error 

``` go 
func (e *UnmarshalTypeError) Error() string
```

### type Unmarshaler 

``` go 
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
```

Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. The input can be assumed to be a valid encoding of a JSON value. UnmarshalJSON must copy the JSON data if it wishes to retain the data after returning.

​	Unmarshaler是由可以解码JSON描述的类型实现的接口。可以假定输入是JSON值的有效编码。如果想在返回后保留数据，UnmarshalJSON必须复制JSON数据。

By convention, to approximate the behavior of Unmarshal itself, Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.

​	按照惯例，为了近似于Unmarshal本身的行为，Unmarshalers将UnmarshalJSON([]byte("null"))实现为无操作。

### type UnsupportedTypeError 

``` go 
type UnsupportedTypeError struct {
	Type reflect.Type
}
```

An UnsupportedTypeError is returned by Marshal when attempting to encode an unsupported value type.

​	在尝试对不受支持的值类型进行编码时，Marshal返回UnsupportedTypeError。

#### (*UnsupportedTypeError) Error 

``` go 
func (e *UnsupportedTypeError) Error() string
```

### type UnsupportedValueError 

``` go 
type UnsupportedValueError struct {
	Value reflect.Value
	Str   string
}
```

An UnsupportedValueError is returned by Marshal when attempting to encode an unsupported value.

​	在尝试对不受支持的值进行编码时，Marshal返回UnsupportedValueError。

#### (*UnsupportedValueError) Error 

``` go 
func (e *UnsupportedValueError) Error() string
```