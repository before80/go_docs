+++
title = "xml"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/encoding/xml@go1.20.1



Package xml implements a simple XML 1.0 parser that understands XML name spaces.

包xml实现了一个简单的XML 1.0解析器，它可以理解XML的名称空间。

##### Example(CustomMarshalXML)
``` go 
package main

import (
	"encoding/xml"
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

func (a *Animal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
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

func (a Animal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}
	return e.EncodeElement(s, start)
}

func main() {
	blob := `
	<animals>
		<animal>gopher</animal>
		<animal>armadillo</animal>
		<animal>zebra</animal>
		<animal>unknown</animal>
		<animal>gopher</animal>
		<animal>bee</animal>
		<animal>gopher</animal>
		<animal>zebra</animal>
	</animals>`
	var zoo struct {
		Animals []Animal `xml:"animal"`
	}
	if err := xml.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo.Animals {
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

##### Example (TextMarshalXML)
``` go 
package main

import (
	"encoding/xml"
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
	blob := `
	<sizes>
		<size>small</size>
		<size>regular</size>
		<size>large</size>
		<size>unrecognized</size>
		<size>small</size>
		<size>normal</size>
		<size>small</size>
		<size>large</size>
	</sizes>`
	var inventory struct {
		Sizes []Size `xml:"size"`
	}
	if err := xml.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}

	counts := make(map[Size]int)
	for _, size := range inventory.Sizes {
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

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/xml/marshal.go;l=19)

``` go 
const (
	// Header is a generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
    // Header是一个通用的XML头，适合与Marshal的输出一起使用。
	// 它不会被自动添加到这个包的任何输出中。
	// 它是作为一种方便提供的。
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/xml/xml.go;l=1866)

``` go 
var HTMLAutoClose []string = htmlAutoClose
```

HTMLAutoClose is the set of HTML elements that should be considered to close automatically.

HTMLAutoClose是应考虑自动关闭的HTML元素的集合。

See the Decoder.Strict and Decoder.Entity fields' documentation.

参见Decoder.Strict和Decoder.Entity字段的文档。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/xml/xml.go;l=1597)

``` go 
var HTMLEntity map[string]string = htmlEntity
```

HTMLEntity is an entity map containing translations for the standard HTML entity characters.

HTMLEntity是一个实体地图，包含标准HTML实体字符的翻译。

See the Decoder.Strict and Decoder.Entity fields' documentation.

参见Decoder.Strict和Decoder.Entity字段的文档。

## 函数

#### func Escape 

``` go 
func Escape(w io.Writer, s []byte)
```

Escape is like EscapeText but omits the error return value. It is provided for backwards compatibility with Go 1.0. Code targeting Go 1.1 or later should use EscapeText.

Escape和EscapeText一样，但省略了错误的返回值。它是为了向后兼容Go 1.0而提供的。针对Go 1.1或更高版本的代码应使用EscapeText。

#### func EscapeText  <- go1.1

``` go 
func EscapeText(w io.Writer, s []byte) error
```

EscapeText writes to w the properly escaped XML equivalent of the plain text data s.

EscapeText向w写出经过适当转义的纯文本数据s的XML等价物。

#### func Marshal 

``` go 
func Marshal(v any) ([]byte, error)
```

Marshal returns the XML encoding of v.

Marshal返回v的XML编码。

Marshal handles an array or slice by marshaling each of the elements. Marshal handles a pointer by marshaling the value it points at or, if the pointer is nil, by writing nothing. Marshal handles an interface value by marshaling the value it contains or, if the interface value is nil, by writing nothing. Marshal handles all other data by writing one or more XML elements containing the data.

Marshal通过处理每个元素来处理一个数组或片断。Marshal处理一个指针，对它所指向的值进行处理，如果指针为nil，则不写任何内容。Marshal处理一个接口值，对其包含的值进行处理，如果接口值为零，则不写任何内容。Marshal通过写一个或多个包含数据的XML元素来处理所有其他数据。

The name for the XML elements is taken from, in order of preference:

XML元素的名称按优先顺序取自：

- the tag on the XMLName field, if the data is a struct 如果数据是一个结构，XMLName字段上的标签
- the value of the XMLName field of type Name 类型为 "名称 "的 XMLName 字段的值
- the tag of the struct field used to obtain the data 用来获取数据的结构字段的标签
- the name of the struct field used to obtain the data 用于获取数据的结构字段的名称
- the name of the marshaled type 调解类型的名称

The XML element for a struct contains marshaled elements for each of the exported fields of the struct, with these exceptions:

结构的XML元素包含结构的每个输出字段的marshaled元素，但有这些例外：

- the XMLName field, described above, is omitted. 省略了上面描述的XMLName字段。
- a field with tag "-" is omitted. 标记为"-"的字段被省略。
- a field with tag "name,attr" becomes an attribute with the given name in the XML element. 标签为 "name,attr "的字段成为 XML 元素中具有给定名称的属性。
- a field with tag ",attr" becomes an attribute with the field name in the XML element. 一个带有标签",attr "的字段在XML元素中成为一个带有字段名的属性。
- a field with tag ",chardata" is written as character data, not as an XML element. 一个带有标签",chardata "的字段被写成字符数据，而不是一个 XML 元素。
- a field with tag ",cdata" is written as character data wrapped in one or more <![CDATA[ ... ]]> tags, not as an XML element. 带有标签",cdata "的字段被写成字符数据，被一个或多个<![CDATA[ ...]>标签包裹，而不是作为一个XML元素。
- a field with tag ",innerxml" is written verbatim, not subject to the usual marshaling procedure. 带有标签",innerxml "的字段被逐字写入，不受通常处理程序的约束。
- a field with tag ",comment" is written as an XML comment, not subject to the usual marshaling procedure. It must not contain the "--" string within it. 带有标签",comment "的字段被写成一个XML注释，不受通常处理程序的限制。它不能包含"--"字符串。
- a field with a tag including the "omitempty" option is omitted if the field value is empty. The empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero. 如果字段的值是空的，带有 "省略 "选项的字段将被省略。空值是 false、0、任何 nil 指针或接口值，以及任何长度为 0 的数组、片断、地图或字符串。
- an anonymous struct field is handled as if the fields of its value were part of the outer struct. 一个匿名的结构字段被处理为其值的字段是外部结构的一部分。
- a field implementing Marshaler is written by calling its MarshalXML method. 实现Marshaler的字段通过调用其MarshalXML方法来编写。
- a field implementing encoding.TextMarshaler is written by encoding the result of its MarshalText method as text. 实现encoding.TextMarshaler的字段通过将其MarshalText方法的结果编码为文本来写入。

If a field uses a tag "a>b>c", then the element c will be nested inside parent elements a and b. Fields that appear next to each other that name the same parent will be enclosed in one XML element.

如果一个字段使用标签 "a>b>c"，那么元素c将被嵌套在父元素a和b里面。出现在彼此旁边的命名相同父元素的字段将被包围在一个XML元素中。

If the XML name for a struct field is defined by both the field tag and the struct's XMLName field, the names must match.

如果结构字段的XML名称同时由字段标签和结构的XMLName字段定义，那么这些名称必须匹配。

See MarshalIndent for an example.

参见MarshalIndent的例子。

Marshal will return an error if asked to marshal a channel, function, or map.

如果要求Marshal对通道、函数或地图进行Marshal，将返回一个错误。

#### func MarshalIndent 

``` go 
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
```

MarshalIndent works like Marshal, but each XML element begins on a new indented line that starts with prefix and is followed by one or more copies of indent according to the nesting depth.

MarshalIndent的工作原理与Marshal类似，但每个XML元素都在一个新的缩进行上开始，该行以prefix开始，后面根据嵌套深度有一个或多个缩进副本。

##### Example
``` go 
```

#### func Unmarshal 

``` go 
func Unmarshal(data []byte, v any) error
```

Unmarshal parses the XML-encoded data and stores the result in the value pointed to by v, which must be an arbitrary struct, slice, or string. Well-formed data that does not fit into v is discarded.

Unmarshal解析XML编码的数据，并将结果存储在v所指向的值中，v必须是一个任意的结构、片断或字符串。不适合v的格式良好的数据被丢弃。

Because Unmarshal uses the reflect package, it can only assign to exported (upper case) fields. Unmarshal uses a case-sensitive comparison to match XML element names to tag values and struct field names.

因为Unmarshal使用reflect包，所以它只能赋值给导出的(大写)字段。Unmarshal使用区分大小写的比较方法，将XML元素名称与标签值和结构字段名称相匹配。

Unmarshal maps an XML element to a struct using the following rules. In the rules, the tag of a field refers to the value associated with the key 'xml' in the struct field's tag (see the example above).

Unmarshal使用以下规则将XML元素映射到结构中。在这些规则中，字段的标签指的是与结构字段标签中的键 "xml "相关的值(见上面的例子)。

- If the struct has a field of type []byte or string with tag ",innerxml", Unmarshal accumulates the raw XML nested inside the element in that field. The rest of the rules still apply. 如果该结构有一个类型为[]字节或字符串的字段，其标签为",innerxml"，那么Unmarshal将积累嵌套在该字段元素中的原始XML。其余的规则仍然适用。
- If the struct has a field named XMLName of type Name, Unmarshal records the element name in that field. 如果该结构有一个类型为Name的XMLName字段，Unmarshal会在该字段中记录元素名称。
- If the XMLName field has an associated tag of the form "name" or "namespace-URL name", the XML element must have the given name (and, optionally, name space) or else Unmarshal returns an error. 如果XMLName字段有一个形式为 "name "或 "namespace-URL name "的相关标签，那么XML元素必须具有给定的名称(以及可选的名称空间)，否则Unmarshal会返回一个错误。
- If the XML element has an attribute whose name matches a struct field name with an associated tag containing ",attr" or the explicit name in a struct field tag of the form "name,attr", Unmarshal records the attribute value in that field. 如果XML元素有一个属性，其名称与包含",attr "的关联标签的结构字段名称或形式为 "name,attr "的结构字段标签中的明确名称相匹配，Unmarshal在该字段中记录属性值。
- If the XML element has an attribute not handled by the previous rule and the struct has a field with an associated tag containing ",any,attr", Unmarshal records the attribute value in the first such field. 如果 XML 元素有一个未被前面规则处理的属性，并且该结构有一个包含",any,attr "关联标签的字段，Unmarshal 会在第一个这样的字段中记录属性值。
- If the XML element contains character data, that data is accumulated in the first struct field that has tag ",chardata". The struct field may have type []byte or string. If there is no such field, the character data is discarded. 如果 XML 元素包含字符数据，该数据将被累积到第一个具有标签",chardata "的结构字段中。该结构字段的类型可以是[]字节或字符串。如果没有这样的字段，字符数据将被丢弃。
- If the XML element contains comments, they are accumulated in the first struct field that has tag ",comment". The struct field may have type []byte or string. If there is no such field, the comments are discarded. 如果XML元素包含注释，它们将被累积到第一个具有",注释 "标签的结构字段中。该结构字段的类型可以是[]字节或字符串。如果没有这样的字段，注释将被丢弃。
- If the XML element contains a sub-element whose name matches the prefix of a tag formatted as "a" or "a>b>c", unmarshal will descend into the XML structure looking for elements with the given names, and will map the innermost elements to that struct field. A tag starting with ">" is equivalent to one starting with the field name followed by ">". 如果XML元素包含一个子元素，其名称与格式为 "a "或 "a>b>c "的标签前缀相匹配，unmarshal将在XML结构中寻找具有给定名称的元素，并将最内部的元素映射到该结构域。以">"开头的标签等同于以字段名后的">"开头的标签。
- If the XML element contains a sub-element whose name matches a struct field's XMLName tag and the struct field has no explicit name tag as per the previous rule, unmarshal maps the sub-element to that struct field. 如果XML元素包含一个子元素，其名称与结构字段的XMLName标签相匹配，并且按照之前的规则，结构字段没有明确的名称标签，那么unmarshal将该子元素映射到该结构字段。
- If the XML element contains a sub-element whose name matches a field without any mode flags (",attr", ",chardata", etc), Unmarshal maps the sub-element to that struct field. 如果 XML 元素包含一个子元素，其名称与没有任何模式标志(",attr",",chardata",等等)的字段相匹配，Unmarshal 将该子元素映射到该结构字段。
- If the XML element contains a sub-element that hasn't matched any of the above rules and the struct has a field with tag ",any", unmarshal maps the sub-element to that struct field. 如果XML元素包含的子元素不符合上述任何规则，并且该结构有一个标签为",any "的字段，那么Umarshal会将该子元素映射到该结构字段。
- An anonymous struct field is handled as if the fields of its value were part of the outer struct. 对匿名结构字段的处理，就像其值的字段是外部结构的一部分一样。
- A struct field with tag "-" is never unmarshaled into. 标签为"-"的结构字段永远不会被解密到。

If Unmarshal encounters a field type that implements the Unmarshaler interface, Unmarshal calls its UnmarshalXML method to produce the value from the XML element. Otherwise, if the value implements encoding.TextUnmarshaler, Unmarshal calls that value's UnmarshalText method.

如果Unmarshal遇到了一个实现了Unmarshaler接口的字段类型，Unmarshal会调用它的UnmarshalXML方法来产生来自XML元素的值。否则，如果该值实现了encoding.TextUnmarshaler，Unmarshal会调用该值的UnmarshalText方法。

Unmarshal maps an XML element to a string or []byte by saving the concatenation of that element's character data in the string or []byte. The saved []byte is never nil.

Unmarshal将一个XML元素映射到一个字符串或[]字节中，将该元素的字符数据的连接保存在字符串或[]字节中。保存的[]字节永远不会是零。

Unmarshal maps an attribute value to a string or []byte by saving the value in the string or slice.

Unmarshal将一个属性值映射到一个字符串或[]字节中，将该值保存在字符串或片断中。

Unmarshal maps an attribute value to an Attr by saving the attribute, including its name, in the Attr.

Unmarshal 将一个属性值映射到一个 Attr 中，通过保存该属性，包括它的名字，在 Attr 中。

Unmarshal maps an XML element or attribute value to a slice by extending the length of the slice and mapping the element or attribute to the newly created value.

Unmarshal 将一个 XML 元素或属性值映射到一个片断，通过扩展片断的长度并将元素或属性映射到新创建的值。

Unmarshal maps an XML element or attribute value to a bool by setting it to the boolean value represented by the string. Whitespace is trimmed and ignored.

Unmarshal 将一个 XML 元素或属性值映射为一个 bool，方法是将其设置为字符串所代表的布尔值。白色空间被修剪并被忽略。

Unmarshal maps an XML element or attribute value to an integer or floating-point field by setting the field to the result of interpreting the string value in decimal. There is no check for overflow. Whitespace is trimmed and ignored.

Unmarshal将一个XML元素或属性值映射到一个整数或浮点字段，方法是将该字段设置为以十进制解释字符串值的结果。没有对溢出的检查。白色的空间被修剪并被忽略。

Unmarshal maps an XML element to a Name by recording the element name.

Unmarshal通过记录元素名称将一个XML元素映射到一个Name。

Unmarshal maps an XML element to a pointer by setting the pointer to a freshly allocated value and then mapping the element to that value.

Unmarshal通过设置指针到一个新分配的值，然后将元素映射到该值，从而将XML元素映射到一个指针。

A missing element or empty attribute value will be unmarshaled as a zero value. If the field is a slice, a zero value will be appended to the field. Otherwise, the field will be set to its zero value.

一个缺失的元素或空属性值将被解封为一个零值。如果字段是一个片断，一个零值将被附加到字段上。否则，字段将被设置为其零值。

##### Unmarshal Example

This example demonstrates unmarshaling an XML excerpt into a value with some preset fields. Note that the Phone field isn't modified and that the XML <Company> element is ignored. Also, the Groups field is assigned considering the element path provided in its tag.

``` go 
package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
}

Output:

XMLName: xml.Name{Space:"", Local:"Person"}
Name: "Grace R. Emlin"
Phone: "none"
Email: [{home gre@example.com} {work gre@work.com}]
Groups: [Friends Squash]
Address: {Hanga Roa Easter Island}

```

## 类型

### type Attr 

``` go 
type Attr struct {
	Name  Name
	Value string
}
```

An Attr represents an attribute in an XML element (Name=Value).

一个Attr代表一个XML元素中的属性(Name=Value)。

### type CharData 

``` go 
type CharData []byte
```

A CharData represents XML character data (raw text), in which XML escape sequences have been replaced by the characters they represent.

一个CharData代表XML字符数据(原始文本)，其中XML转义序列已经被它们所代表的字符所取代。

#### (CharData) Copy 

``` go 
func (c CharData) Copy() CharData
```

Copy creates a new copy of CharData.

Copy 创建一个新的CharData的副本。

### type Comment 

``` go 
type Comment []byte
```

A Comment represents an XML comment of the form <!--comment-->. The bytes do not include the <!-- and --> comment markers.

一个Comment代表一个XML注释，其形式为<！--comment-->。字节不包括<！--和-->注释标记。

#### (Comment) Copy 

``` go 
func (c Comment) Copy() Comment
```

Copy creates a new copy of Comment.

Copy创建一个Comment的新副本。

### type Decoder 

``` go 
type Decoder struct {
	// Strict defaults to true, enforcing the requirements
	// of the XML specification.
	// If set to false, the parser allows input containing common
	// mistakes:
	//	* If an element is missing an end tag, the parser invents
	//	  end tags as necessary to keep the return values from Token
	//	  properly balanced.
	//	* In attribute values and character data, unknown or malformed
	//	  character entities (sequences beginning with &) are left alone.
	// Strict默认为true，强制执行XML规范的要求。
	// 如果设置为false，解析器允许包含常见错误的输入。
	// * 如果一个元素缺少一个结束标签，解析器会根据需要发明结束标签，以保持Token的返回值的正确平衡。
	// * 在属性值和字符数据中，未知的或畸形的字符实体(以&开头的序列)将不被考虑。
	//
	// Setting:
	//
	//	d.Strict = false
	//	d.AutoClose = xml.HTMLAutoClose
	//	d.Entity = xml.HTMLEntity
	//
	// creates a parser that can handle typical HTML.
	//
	// Strict mode does not enforce the requirements of the XML name spaces TR.
	// In particular it does not reject name space tags using undefined prefixes.
	// Such tags are recorded with the unknown prefix as the name space URL.
	// 创建一个可以处理典型HTML的解析器。
	//
	// 严格模式不强制执行XML名称空间TR的要求。
	// 尤其是它不拒绝使用未定义前缀的名称空间标签。
	// 这样的标签会以未知的前缀作为名称空间的URL来记录。
	Strict bool

	// When Strict == false, AutoClose indicates a set of elements to
	// consider closed immediately after they are opened, regardless
	// of whether an end element is present.
	// 当Strict == false时，AutoClose表示一组元素，在它们被打开后立即考虑关闭，不管是否有结束元素存在。
	AutoClose []string

	// Entity can be used to map non-standard entity names to string replacements.
	// The parser behaves as if these standard mappings are present in the map,
	// regardless of the actual map content:
	// 实体可以用来将非标准的实体名称映射为字符串替换。
	// 解析器的行为就像地图中存在这些标准的映射一样。 不管实际的地图内容是什么：
	//
	//	"lt": "<",
	//	"gt": ">",
	//	"amp": "&",
	//	"apos": "'",
	//	"quot": `"`,
	Entity map[string]string

	// CharsetReader, if non-nil, defines a function to generate
	// charset-conversion readers, converting from the provided
	// non-UTF-8 charset into UTF-8. If CharsetReader is nil or
	// returns an error, parsing stops with an error. One of the
	// CharsetReader's result values must be non-nil.
	// CharsetReader，如果不是nil，定义了一个函数来生成字符集转换阅读器，从提供的非UTF-8字符集转换为UTF-8。如果CharsetReader为零或返回错误，解析将以错误停止。CharsetReader的结果值之一必须是非零的。
	CharsetReader func(charset string, input io.Reader) (io.Reader, error)

	// DefaultSpace sets the default name space used for unadorned tags,
	// as if the entire XML stream were wrapped in an element containing
	// the attribute xmlns="DefaultSpace".
	// DefaultSpace设置用于未装饰标签的默认名称空间，就像整个XML流被包裹在一个包含xmlns="DefaultSpace "属性的元素中。
	DefaultSpace string
	// contains filtered or unexported fields
}
```

A Decoder represents an XML parser reading a particular input stream. The parser assumes that its input is encoded in UTF-8.

Decoder 代表一个读取特定输入流的XML解析器。该解析器假定其输入是以UTF-8编码的。

#### func NewDecoder 

``` go 
func NewDecoder(r io.Reader) *Decoder
```

NewDecoder creates a new XML parser reading from r. If r does not implement io.ByteReader, NewDecoder will do its own buffering.

如果r没有实现io.ByteReader，NewDecoder会自己做缓冲。

#### func NewTokenDecoder  <- go1.10

``` go 
func NewTokenDecoder(t TokenReader) *Decoder
```

NewTokenDecoder creates a new XML parser using an underlying token stream.

NewTokenDecoder使用底层令牌流创建一个新的XML解析器。

#### (*Decoder) Decode 

``` go 
func (d *Decoder) Decode(v any) error
```

Decode works like Unmarshal, except it reads the decoder stream to find the start element.

Decode的工作原理与Unmarshal类似，只是它读取解码器流来寻找起始元素。

#### (*Decoder) DecodeElement 

``` go 
func (d *Decoder) DecodeElement(v any, start *StartElement) error
```

DecodeElement works like Unmarshal except that it takes a pointer to the start XML element to decode into v. It is useful when a client reads some raw XML tokens itself but also wants to defer to Unmarshal for some elements.

DecodeElement的工作原理与Unmarshal类似，只是它需要一个指向起始XML元素的指针来解码成v。当客户端自己读取一些原始的XML标记，但也希望对一些元素推迟到Unmarshal时，它是非常有用的。

#### (*Decoder) InputOffset  <- go1.4

``` go 
func (d *Decoder) InputOffset() int64
```

InputOffset returns the input stream byte offset of the current decoder position. The offset gives the location of the end of the most recently returned token and the beginning of the next token.

InputOffset返回当前解码器位置的输入流字节偏移。这个偏移量给出了最近返回的令牌的结束位置和下一个令牌的开始位置。

#### (*Decoder) InputPos  <- go1.19

``` go 
func (d *Decoder) InputPos() (line, column int)
```

InputPos returns the line of the current decoder position and the 1 based input position of the line. The position gives the location of the end of the most recently returned token.

InputPos返回当前解码器位置的行，以及该行的基于1的输入位置。该位置给出了最近返回的token的结束位置。

#### (*Decoder) RawToken 

``` go 
func (d *Decoder) RawToken() (Token, error)
```

RawToken is like Token but does not verify that start and end elements match and does not translate name space prefixes to their corresponding URLs.

RawToken和Token一样，但是不验证开始和结束元素是否匹配，也不把名称空间前缀翻译成相应的URL。

#### (*Decoder) Skip 

``` go 
func (d *Decoder) Skip() error
```

Skip reads tokens until it has consumed the end element matching the most recent start element already consumed, skipping nested structures. It returns nil if it finds an end element matching the start element; otherwise it returns an error describing the problem.

跳过读取令牌，直到它消耗了与已经消耗的最近的开始元素相匹配的结束元素，跳过嵌套结构。如果它找到了与开始元素相匹配的结束元素，则返回nil；否则它将返回一个描述问题的错误。

#### (*Decoder) Token 

``` go 
func (d *Decoder) Token() (Token, error)
```

Token returns the next XML token in the input stream. At the end of the input stream, Token returns nil, io.EOF.

Token返回输入流中的下一个XML标记。在输入流结束时，Token返回nil, io.EOF。

Slices of bytes in the returned token data refer to the parser's internal buffer and remain valid only until the next call to Token. To acquire a copy of the bytes, call CopyToken or the token's Copy method.

返回的令牌数据中的字节片指的是解析器的内部缓冲区，并且只在下次调用Token之前保持有效。要获得字节的拷贝，可以调用CopyToken或令牌的Copy方法。

Token expands self-closing elements such as `<br>` into separate start and end elements returned by successive calls.

Token将自闭元素如`<br>`扩展成独立的开始和结束元素，由连续的调用返回。

Token guarantees that the StartElement and EndElement tokens it returns are properly nested and matched: if Token encounters an unexpected end element or EOF before all expected end elements, it will return an error.

Token 保证它返回的 StartElement 和 EndElement 令牌是正确嵌套和匹配的：如果 Token 遇到意外的结束元素或在所有预期结束元素之前遇到 EOF，它将返回一个错误。

Token implements XML name spaces as described by https://www.w3.org/TR/REC-xml-names/. Each of the Name structures contained in the Token has the Space set to the URL identifying its name space when known. If Token encounters an unrecognized name space prefix, it uses the prefix as the Space rather than report an error.

Token 实现了 XML 名称空间，如 https://www.w3.org/TR/REC-xml-names/ 所述。每个包含在 Token 中的 Name 结构的 Space 都被设置为识别其名称空间的 URL(当已知时)。如果 Token 遇到未被识别的名称空间前缀，它将使用该前缀作为 Space，而不是报告一个错误。

### type Directive 

``` go 
type Directive []byte
```

A Directive represents an XML directive of the form <!text>. The bytes do not include the <! and > markers.

Directive代表一个XML指令，其形式为<！text>。字节不包括<！和>标记。

#### (Directive) Copy 

``` go 
func (d Directive) Copy() Directive
```

Copy creates a new copy of Directive.

Copy创建Directive的一个新副本。

### type Encoder 

``` go 
type Encoder struct {
	// contains filtered or unexported fields
}
```

An Encoder writes XML data to an output stream.

Encoder 将XML数据写入一个输出流。

##### Example
``` go 
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

Output:

  <person id="13">
      <name>
          <first>John</first>
          <last>Doe</last>
      </name>
      <age>42</age>
      <Married>false</Married>
      <City>Hanga Roa</City>
      <State>Easter Island</State>
      <!-- Need more details. -->
  </person>
```

#### func NewEncoder 

``` go 
func NewEncoder(w io.Writer) *Encoder
```

NewEncoder returns a new encoder that writes to w.

NewEncoder返回一个向w写的新编码器。

#### (*Encoder) Close  <- go1.20

``` go 
func (enc *Encoder) Close() error
```

Close the Encoder, indicating that no more data will be written. It flushes any buffered XML to the underlying writer and returns an error if the written XML is invalid (e.g. by containing unclosed elements).

关闭Encoder，表示不再写入数据。它将任何缓冲的XML冲到底层写入器，如果写入的XML无效(例如包含未封闭的元素)，则返回一个错误。

#### (*Encoder) Encode 

``` go 
func (enc *Encoder) Encode(v any) error
```

Encode writes the XML encoding of v to the stream.

Encode将v的XML编码写到流中。

See the documentation for Marshal for details about the conversion of Go values to XML.

关于Go值转换为XML的细节，请参见Marshal的文档。

Encode calls Flush before returning.

Encode在返回之前调用Flush。

#### (*Encoder) EncodeElement  <- go1.2

``` go 
func (enc *Encoder) EncodeElement(v any, start StartElement) error
```

EncodeElement writes the XML encoding of v to the stream, using start as the outermost tag in the encoding.

EncodeElement将v的XML编码写到流中，使用start作为编码的最外层标签。

See the documentation for Marshal for details about the conversion of Go values to XML.

关于Go值转换为XML的细节，请参见Marshal的文档。

EncodeElement calls Flush before returning.

EncodeElement在返回前调用Flush。

#### (*Encoder) EncodeToken  <- go1.2

``` go 
func (enc *Encoder) EncodeToken(t Token) error
```

EncodeToken writes the given XML token to the stream. It returns an error if StartElement and EndElement tokens are not properly matched.

EncodeToken将给定的XML令牌写到流中。如果StartElement和EndElement令牌没有正确匹配，它会返回一个错误。

EncodeToken does not call Flush, because usually it is part of a larger operation such as Encode or EncodeElement (or a custom Marshaler's MarshalXML invoked during those), and those will call Flush when finished. Callers that create an Encoder and then invoke EncodeToken directly, without using Encode or EncodeElement, need to call Flush when finished to ensure that the XML is written to the underlying writer.

EncodeToken不调用Flush，因为通常它是一个更大的操作的一部分，如Encode或EncodeElement(或在这些操作中调用的自定义MarshalXML)，这些操作完成后将调用Flush。创建Encoder然后直接调用EncodeToken的调用者，不使用Encode或EncodeElement，需要在完成后调用Flush以确保XML被写入底层写入器。

EncodeToken allows writing a ProcInst with Target set to "xml" only as the first token in the stream.

EncodeToken允许写一个ProcInst，目标设置为 "xml"，只作为流中的第一个标记。

#### (*Encoder) Flush  <- go1.2

``` go 
func (enc *Encoder) Flush() error
```

Flush flushes any buffered XML to the underlying writer. See the EncodeToken documentation for details about when it is necessary.

Flush将任何缓冲的XML冲到底层写入器中。关于什么时候需要这样做，请参见EncodeToken文档。

#### (*Encoder) Indent  <- go1.1

``` go 
func (enc *Encoder) Indent(prefix, indent string)
```

Indent sets the encoder to generate XML in which each element begins on a new indented line that starts with prefix and is followed by one or more copies of indent according to the nesting depth.

缩进设置编码器生成XML，其中每个元素在一个新的缩进行中开始，该行以prefix开始，后面根据嵌套深度有一个或多个缩进副本。

### type EndElement 

``` go 
type EndElement struct {
	Name Name
}
```

An EndElement represents an XML end element.

一个EndElement代表一个XML的结束元素。

### type Marshaler  <- go1.2

``` go 
type Marshaler interface {
	MarshalXML(e *Encoder, start StartElement) error
}
```

Marshaler is the interface implemented by objects that can marshal themselves into valid XML elements.

Marshaler是由能够将自己 Marshal成有效的XML元素的对象实现的接口。

MarshalXML encodes the receiver as zero or more XML elements. By convention, arrays or slices are typically encoded as a sequence of elements, one per entry. Using start as the element tag is not required, but doing so will enable Unmarshal to match the XML elements to the correct struct field. One common implementation strategy is to construct a separate value with a layout corresponding to the desired XML and then to encode it using e.EncodeElement. Another common strategy is to use repeated calls to e.EncodeToken to generate the XML output one token at a time. The sequence of encoded tokens must make up zero or more valid XML elements.

MarshalXML将接收器编码为零个或多个XML元素。根据惯例，数组或片断通常被编码为一个元素序列，每个条目一个。使用 start 作为元素标签并不是必须的，但这样做将使 Unmarshal 能够将 XML 元素与正确的结构字段相匹配。一种常见的实现策略是构建一个单独的值，其布局对应于所需的 XML，然后使用 e.EncodeElement 对其进行编码。另一种常见的策略是使用对e.EncodeToken的重复调用，一次生成一个标记的XML输出。编码令牌的序列必须由零个或多个有效的XML元素组成。

### type MarshalerAttr  <- go1.2

``` go 
type MarshalerAttr interface {
	MarshalXMLAttr(name Name) (Attr, error)
}
```

MarshalerAttr is the interface implemented by objects that can marshal themselves into valid XML attributes.

MarshalerAttr 是由可以将自己 Marshal 成有效的 XML 属性的对象实现的接口。

MarshalXMLAttr returns an XML attribute with the encoded value of the receiver. Using name as the attribute name is not required, but doing so will enable Unmarshal to match the attribute to the correct struct field. If MarshalXMLAttr returns the zero attribute Attr{}, no attribute will be generated in the output. MarshalXMLAttr is used only for struct fields with the "attr" option in the field tag.

MarshalXMLAttr返回一个带有接收器编码值的XML属性。使用 name 作为属性名不是必须的，但这样做可以使 Unmarshal 将属性与正确的结构字段相匹配。如果 MarshalXMLAttr 返回零属性 Attr{}，输出中不会产生任何属性。MarshalXMLAttr 仅用于字段标签中带有 "attr "选项的结构字段。

### type Name 

``` go 
type Name struct {
	Space, Local string
}
```

A Name represents an XML name (Local) annotated with a name space identifier (Space). In tokens returned by Decoder.Token, the Space identifier is given as a canonical URL, not the short prefix used in the document being parsed.

一个Name代表一个XML名称(Local)，用一个名称空间标识符(Space)来注释。在Decoder.Token返回的令牌中，Space标识符是作为一个规范的URL给出的，而不是被解析的文档中使用的短前缀。

### type ProcInst 

``` go 
type ProcInst struct {
	Target string
	Inst   []byte
}
```

A ProcInst represents an XML processing instruction of the form <?target inst?>

一个ProcInst代表一个XML处理指令，其形式为`<?target inst?>`。

#### (ProcInst) Copy 

``` go 
func (p ProcInst) Copy() ProcInst
```

Copy creates a new copy of ProcInst.

Copy 创建ProcInst的新副本。

### type StartElement 

``` go 
type StartElement struct {
	Name Name
	Attr []Attr
}
```

A StartElement represents an XML start element.

StartElement代表一个XML起始元素。

#### (StartElement) Copy 

``` go 
func (e StartElement) Copy() StartElement
```

Copy creates a new copy of StartElement.

Copy 创建一个 StartElement 的新副本。

#### (StartElement) End  <- go1.2

``` go 
func (e StartElement) End() EndElement
```

End returns the corresponding XML end element.

End返回相应的XML结束元素。

### type SyntaxError 

``` go 
type SyntaxError struct {
	Msg  string
	Line int
}
```

A SyntaxError represents a syntax error in the XML input stream.

SyntaxError代表XML输入流中的一个语法错误。

#### (*SyntaxError) Error 

``` go 
func (e *SyntaxError) Error() string
```

### type TagPathError 

``` go 
type TagPathError struct {
	Struct       reflect.Type
	Field1, Tag1 string
	Field2, Tag2 string
}
```

A TagPathError represents an error in the unmarshaling process caused by the use of field tags with conflicting paths.

TagPathError表示在解封过程中，由于使用了路径冲突的字段标签而导致的错误。

#### (*TagPathError) Error 

``` go 
func (e *TagPathError) Error() string
```

### type Token 

``` go 
type Token any
```

A Token is an interface holding one of the token types: StartElement, EndElement, CharData, Comment, ProcInst, or Directive.

Token是一个接口，持有一个令牌类型。StartElement, EndElement, CharData, Comment, ProcInst, or Directive.

#### func CopyToken 

``` go 
func CopyToken(t Token) Token
```

CopyToken returns a copy of a Token.

CopyToken返回一个令牌的副本。

### type TokenReader  <- go1.10

``` go 
type TokenReader interface {
	Token() (Token, error)
}
```

A TokenReader is anything that can decode a stream of XML tokens, including a Decoder.

TokenReader是任何可以解码XML标记流的东西，包括解码器。

When Token encounters an error or end-of-file condition after successfully reading a token, it returns the token. It may return the (non-nil) error from the same call or return the error (and a nil token) from a subsequent call. An instance of this general case is that a TokenReader returning a non-nil token at the end of the token stream may return either io.EOF or a nil error. The next Read should return nil, io.EOF.

当Token在成功读取一个令牌后遇到错误或文件结束的情况时，它会返回该令牌。它可以从同一个调用中返回(非零)错误，或者从后续调用中返回错误(和一个零的令牌)。这个一般情况的一个例子是，TokenReader在令牌流的末端返回一个非空的令牌，可能会返回io.EOF或者一个nil错误。下一个Read应该返回nil, io.EOF。

Implementations of Token are discouraged from returning a nil token with a nil error. Callers should treat a return of nil, nil as indicating that nothing happened; in particular it does not indicate EOF.

不鼓励Token的实现在返回nil令牌时出现nil错误。调用者应该把返回的nil, nil看作是没有发生任何事情；特别是它并不表示EOF。

### type UnmarshalError 

``` go 
type UnmarshalError string
```

An UnmarshalError represents an error in the unmarshaling process.

一个UnmarshalError表示在解密过程中的一个错误。

#### (UnmarshalError) Error 

``` go 
func (e UnmarshalError) Error() string
```

### type Unmarshaler  <- go1.2

``` go 
type Unmarshaler interface {
	UnmarshalXML(d *Decoder, start StartElement) error
}
```

Unmarshaler is the interface implemented by objects that can unmarshal an XML element description of themselves.

Unmarshaler是由能够解读自己的XML元素描述的对象实现的接口。

UnmarshalXML decodes a single XML element beginning with the given start element. If it returns an error, the outer call to Unmarshal stops and returns that error. UnmarshalXML must consume exactly one XML element. One common implementation strategy is to unmarshal into a separate value with a layout matching the expected XML using d.DecodeElement, and then to copy the data from that value into the receiver. Another common strategy is to use d.Token to process the XML object one token at a time. UnmarshalXML may not use d.RawToken.

UnmarshalXML对从给定的start元素开始的单个XML元素进行解码。如果它返回一个错误，对Unmarshal的外部调用就会停止并返回该错误。UnmarshalXML必须正好消耗一个XML元素。一个常见的实现策略是使用d.DecodeElement将解压缩到一个单独的值，其布局与预期的XML相匹配，然后将该值中的数据复制到接收器中。另一个常见的策略是使用d.Token来一次处理XML对象的一个token。UnmarshalXML可能不会使用d.RawToken。

### type UnmarshalerAttr  <- go1.2

``` go 
type UnmarshalerAttr interface {
	UnmarshalXMLAttr(attr Attr) error
}
```

UnmarshalerAttr is the interface implemented by objects that can unmarshal an XML attribute description of themselves.

UnmarshalerAttr是由能够解读自身的XML属性描述的对象实现的接口。

UnmarshalXMLAttr decodes a single XML attribute. If it returns an error, the outer call to Unmarshal stops and returns that error. UnmarshalXMLAttr is used only for struct fields with the "attr" option in the field tag.

UnmarshalXMLAttr 解码一个单一的 XML 属性。如果它返回一个错误，外部对 Unmarshal 的调用就会停止并返回该错误。UnmarshalXMLAttr 仅用于字段标签中带有 "attr "选项的结构字段。

### type UnsupportedTypeError 

``` go 
type UnsupportedTypeError struct {
	Type reflect.Type
}
```

UnsupportedTypeError is returned when Marshal encounters a type that cannot be converted into XML.

当Marshal遇到不能转换为XML的类型时，UnsupportedTypeError被返回。

#### (*UnsupportedTypeError) Error 

``` go 
func (e *UnsupportedTypeError) Error() string
```

## Notes

## Bugs

- Mapping between XML elements and data structures is inherently flawed: an XML element is an order-dependent collection of anonymous values, while a data structure is an order-independent collection of named values. See package json for a textual representation more suitable to data structures.
- XML元素和数据结构之间的映射是有内在缺陷的：XML元素是一个匿名值的顺序依赖项集合，而数据结构是一个命名值的顺序无关的集合。请参阅包json，以获得更适合数据结构的文本表示法。