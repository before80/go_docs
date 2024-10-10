+++
title = "tabwriter"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/text/tabwriter@go1.23.0](https://pkg.go.dev/text/tabwriter@go1.23.0)

Package tabwriter implements a write filter (tabwriter.Writer) that translates tabbed columns in input into properly aligned text.

​	`tabwriter`包实现了一个写入过滤器（tabwriter.Writer），它将输入中的制表符列转换为正确对齐的文本。

The package is using the Elastic Tabstops algorithm described at http://nickgravgaard.com/elastictabstops/index.html.

​	该包使用了在[http://nickgravgaard.com/elastictabstops/index.html](http://nickgravgaard.com/elastictabstops/index.html)上描述的弹性制表符算法。

The text/tabwriter package is frozen and is not accepting new features.

​	`text/tabwriter`包已经冻结，不再接受新功能。

## Example (Elastic 灵活的)
``` go 
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	// Observe how the b's and the d's, despite appearing in the
	// second cell of each line, belong to different columns.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // trailing tab
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()

}

//Output:

//....a|..b|c
//...aa|.bb|cc
//..aaa|
//.aaaa|.dddd|eeee
```

## Example (TrailingTab)
``` go 
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	// Observe that the third line has no trailing tab,
	// so its final cell is not part of an aligned column.
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned") // no trailing tab
	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
	w.Flush()

}

//Output:

//------a|------b|---aligned|
//-----aa|-----bb|---aligned|
//----aaa|----bbb|unaligned
//---aaaa|---bbbb|---aligned|
```



## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/text/tabwriter/tabwriter.go;l=169)

``` go 
const (
	// Ignore html tags and treat entities (starting with '&'
	// and ending in ';') as single characters (width = 1).
    // 忽略HTML标签，并将实体（以'&'开头，以';'结尾）视为单个字符（宽度=1）。
	FilterHTML uint = 1 << iota

	// Strip Escape characters bracketing escaped text segments
	// instead of passing them through unchanged with the text.
    // 剥离转义字符，而不是将其与文本一起保持不变。
	StripEscape

	// Force right-alignment of cell content.
	// Default is left-alignment.
    // 强制将单元格内容右对齐。
	// 默认为左对齐。
	AlignRight

	// Handle empty columns as if they were not present in
	// the input in the first place.
    // 将空列处理为在输入中不存在的列。
	DiscardEmptyColumns

	// Always use tabs for indentation columns (i.e., padding of
	// leading empty cells on the left) independent of padchar.
    // 对缩进列始终使用制表符（即左侧的前导空单元格填充），而不受padchar的影响。
	TabIndent

	// Print a vertical bar ('|') between columns (after formatting).
	// Discarded columns appear as zero-width columns ("||").
    // 在列之间打印垂直条（' | '）（在格式化之后）。
	// 被丢弃的列出现为零宽度列（'||'）。
	Debug
)
```

Formatting can be controlled with these flags.

​	可以使用这些标志来控制格式化。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/text/tabwriter/tabwriter.go;l=426)

``` go 
const Escape = '\xff'
```

To escape a text segment, bracket it with Escape characters. For instance, the tab in this string "Ignore this tab: \xff\t\xff" does not terminate a cell and constitutes a single character of width one for formatting purposes.

​	要转义文本段，请使用转义字符括起来。例如，字符串 "Ignore this tab: \xff\t\xff" 中的制表符不会终止单元格，并且在格式化时宽度为一个字符。

The value 0xff was chosen because it cannot appear in a valid UTF-8 sequence.

​	选择值0xff是因为它不能出现在有效的UTF-8序列中。

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Writer 

``` go 
type Writer struct {
	// contains filtered or unexported fields
}
```

A Writer is a filter that inserts padding around tab-delimited columns in its input to align them in the output.

​	Writer是一个过滤器，它在输入中的制表符分隔的列周围插入填充，以便在输出中对齐它们。

The Writer treats incoming bytes as UTF-8-encoded text consisting of cells terminated by horizontal ('\t') or vertical ('\v') tabs, and newline ('\n') or formfeed ('\f') characters; both newline and formfeed act as line breaks.

​	Writer将输入字节视为UTF-8编码的文本，由水平（'\t'）或垂直（'\v'）制表符以及换行（'\n'）或换页（'\f'）字符组成；换行和换页都作为换行符。

Tab-terminated cells in contiguous lines constitute a column. The Writer inserts padding as needed to make all cells in a column have the same width, effectively aligning the columns. It assumes that all characters have the same width, except for tabs for which a tabwidth must be specified. Column cells must be tab-terminated, not tab-separated: non-tab terminated trailing text at the end of a line forms a cell but that cell is not part of an aligned column. For instance, in this example (where | stands for a horizontal tab):

​	连续行中的以制表符结尾的单元格构成一列。Writer根据需要插入填充，使得一列中的所有单元格具有相同的宽度，从而对齐列。它假设所有字符的宽度相同，除了制表符，必须指定一个制表符宽度。列单元格必须以制表符结尾，而不是以制表符分隔：在行末尾的非制表符结尾的尾随文本形成一个单元格，但该单元格不是对齐列的一部分。例如，在以下示例中（其中|代表水平制表符）：

```
aaaa|bbb|d
aa  |b  |dd
a   |
aa  |cccc|eee
```

the b and c are in distinct columns (the b column is not contiguous all the way). The d and e are not in a column at all (there's no terminating tab, nor would the column be contiguous).

​	b和c位于不同的列（b列不是一直连续的）。d和e根本不在列中（没有终止制表符，也不会连续的列）。

The Writer assumes that all Unicode code points have the same width; this may not be true in some fonts or if the string contains combining characters.

​	Writer假设所有Unicode码点具有相同的宽度；在某些字体中或字符串包含组合字符时，可能不成立。

If DiscardEmptyColumns is set, empty columns that are terminated entirely by vertical (or "soft") tabs are discarded. Columns terminated by horizontal (or "hard") tabs are not affected by this flag.

​	如果设置了DiscardEmptyColumns，以垂直（或"软"）制表符完全终止的空列将被丢弃。以水平（或"硬"）制表符终止的列不受此标志的影响。

If a Writer is configured to filter HTML, HTML tags and entities are passed through. The widths of tags and entities are assumed to be zero (tags) and one (entities) for formatting purposes.

​	如果Writer配置为过滤HTML，则会通过HTML标签和实体。对于格式化目的，标签和实体的宽度被假定为零（标签）和一（实体）。

A segment of text may be escaped by bracketing it with Escape characters. The tabwriter passes escaped text segments through unchanged. In particular, it does not interpret any tabs or line breaks within the segment. If the StripEscape flag is set, the Escape characters are stripped from the output; otherwise they are passed through as well. For the purpose of formatting, the width of the escaped text is always computed excluding the Escape characters.

​	文本段可以通过用转义字符括起来来进行转义。tabwriter将转义的文本段原样传递。特别是，它不会解释段内的任何制表符或换行符。如果设置了StripEscape标志，将从输出中剥离转义字符；否则也将它们原样传递。对于格式化目的，转义文本的宽度始终是不包括转义字符在内的。

The formfeed character acts like a newline but it also terminates all columns in the current line (effectively calling Flush). Tab- terminated cells in the next line start new columns. Unless found inside an HTML tag or inside an escaped text segment, formfeed characters appear as newlines in the output.

换页字符的行为类似于换行符，但它还会终止当前行中的所有列（实际上调用Flush）。在下一行中，以制表符结尾的单元格将开始新的列。除非在HTML标签或转义的文本段内，换页字符在输出中显示为换行符。

The Writer must buffer input internally, because proper spacing of one line may depend on the cells in future lines. Clients must call Flush when done calling Write.

​	Writer必须在内部缓冲输入，因为一行的正确间距可能取决于未来行中的单元格。在调用Write完成后，客户端必须调用Flush。

#### func NewWriter 

``` go 
func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
```

NewWriter allocates and initializes a new tabwriter.Writer. The parameters are the same as for the Init function.

​	NewWriter函数分配并初始化一个新的tabwriter.Writer。参数与Init函数相同。

#### (*Writer) Flush 

``` go 
func (b *Writer) Flush() error
```

Flush should be called after the last call to Write to ensure that any data buffered in the Writer is written to output. Any incomplete escape sequence at the end is considered complete for formatting purposes.

​	Flush方法应在最后一次调用Write后调用，以确保将Writer中缓冲的任何数据写入输出。对于格式化目的，尾部的任何不完整的转义序列都被视为完整。

#### (*Writer) Init 

``` go 
func (b *Writer) Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
```

A Writer must be initialized with a call to Init. The first parameter (output) specifies the filter output. The remaining parameters control the formatting:

​	必须使用Init调用来初始化Writer。第一个参数（output）指定过滤器的输出。其余参数控制格式化：

```
minwidth	minimal cell width including any padding 包括任何填充的最小单元格宽度
tabwidth	width of tab characters (equivalent number of spaces) 制表符字符的宽度（相当于空格的数量）
padding		padding added to a cell before computing its width 在计算单元格宽度之前添加的填充
padchar		ASCII char used for padding 用于填充的ASCII字符
		if padchar == '\t', the Writer will assume that the
		width of a '\t' in the formatted output is tabwidth,
		and cells are left-aligned independent of align_left
		(for correct-looking results, tabwidth must correspond
		to the tab width in the viewer displaying the result)
		如果padchar == '\t'，则Writer将假定格式化输出中'\t'的宽度为tabwidth，
		并且单元格是左对齐的，不受align_left的影响
		（为了获得正确的结果，tabwidth必须与显示结果的查看器中的制表符宽度相对应）
flags		formatting control 格式化控制
```

##### Init Example
``` go 
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()

	// Format right-aligned in space-separated columns of minimal width 5
	// and at least one blank of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()

}

//Output:

//a	b	c	d		.
//123	12345	1234567	123456789	.

//    a     b       c         d.
//  123 12345 1234567 123456789.
```

#### (*Writer) Write 

``` go 
func (b *Writer) Write(buf []byte) (n int, err error)
```

Write writes buf to the writer b. The only errors returned are ones encountered while writing to the underlying output stream.

​	Write方法将buf写入写入器b。仅返回在写入底层输出流时遇到的错误。