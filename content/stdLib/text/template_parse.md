+++
title = "template/parse"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/text/template/parse@go1.23.0](https://pkg.go.dev/text/template/parse@go1.23.0)

Package parse builds parse trees for templates as defined by text/template and html/template. Clients should use those packages to construct templates rather than this one, which provides shared internal data structures not intended for general use.

​	`parse`包根据`text/template`和`html/template`定义的模板构建解析树。客户端应使用这些包来构建模板，而不是使用本包，本包提供了一些共享的内部数据结构，不适用于一般用途。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func IsEmptyTree 

``` go 
func IsEmptyTree(n Node) bool
```

IsEmptyTree reports whether this tree (node) is empty of everything but space or comments.

​	IsEmptyTree报告此树（节点）是否除了空格或注释之外没有任何内容。

### func Parse 

``` go 
func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]any) (map[string]*Tree, error)
```

Parse returns a map from template name to parse.Tree, created by parsing the templates described in the argument string. The top-level template will be given the specified name. If an error is encountered, parsing stops and an empty map is returned with the error.

​	Parse根据参数字符串中描述的模板解析信息，返回一个从模板名称到parse.Tree的映射。顶级模板将被赋予指定的名称。如果遇到错误，解析停止并返回一个空映射和错误。

## 类型

### type ActionNode 

``` go 
type ActionNode struct {
	NodeType
	Pos

	Line int       // The line number in the input. Deprecated: Kept for compatibility. 输入中的行号。已弃用：保留兼容性。
	Pipe *PipeNode // The pipeline in the action. 动作中的管道。
	// contains filtered or unexported fields
}
```

ActionNode holds an action (something bounded by delimiters). Control actions have their own nodes; ActionNode represents simple ones such as field evaluations and parenthesized pipelines.

​	ActionNode保存一个动作（由分隔符界定的内容）。控制动作有自己的节点；ActionNode表示简单的动作，例如字段求值和带括号的管道。

#### (*ActionNode) Copy 

``` go 
func (a *ActionNode) Copy() Node
```

#### (*ActionNode) String 

``` go 
func (a *ActionNode) String() string
```

### type BoolNode 

``` go 
type BoolNode struct {
	NodeType
	Pos

	True bool // The value of the boolean constant. 布尔常量的值。
	// contains filtered or unexported fields
}
```

BoolNode holds a boolean constant.

​	BoolNode保存一个布尔常量。

#### (*BoolNode) Copy 

``` go 
func (b *BoolNode) Copy() Node
```

#### (*BoolNode) String 

``` go 
func (b *BoolNode) String() string
```

### type BranchNode 

``` go 
type BranchNode struct {
	NodeType
	Pos

	Line     int       // The line number in the input. Deprecated: Kept for compatibility. 输入中的行号。已弃用：保留兼容性。
	Pipe     *PipeNode // The pipeline to be evaluated. 要评估的管道。
	List     *ListNode // What to execute if the value is non-empty. 如果值非空，则执行的内容。
	ElseList *ListNode // What to execute if the value is empty (nil if absent). 如果值为空时执行的内容（如果不存在，则为nil）。
	// contains filtered or unexported fields
}
```

BranchNode is the common representation of if, range, and with.

​	BranchNode是if、range和with的通用表示。

#### (*BranchNode) Copy  <- go1.4

``` go 
func (b *BranchNode) Copy() Node
```

#### (*BranchNode) String 

``` go 
func (b *BranchNode) String() string
```

### type BreakNode  <- go1.18

``` go 
type BreakNode struct {
	NodeType
	Pos
	Line int
	// contains filtered or unexported fields
}
```

BreakNode represents a {{break}} action.

​	BreakNode表示{{break}}动作。

#### (*BreakNode) Copy  <- go1.18

``` go 
func (b *BreakNode) Copy() Node
```

#### (*BreakNode) String  <- go1.18

``` go 
func (b *BreakNode) String() string
```

### type ChainNode  <- go1.1

``` go 
type ChainNode struct {
	NodeType
	Pos

	Node  Node
	Field []string // The identifiers in lexical order. 以词法顺序的标识符。
	// contains filtered or unexported fields
}
```

ChainNode holds a term followed by a chain of field accesses (identifier starting with '.'). The names may be chained ('.x.y'). The periods are dropped from each ident.

​	ChainNode保存一个术语，后面是一系列的字段访问（以'.'开头的标识符）。名称可以链接（'.x.y'）。每个标识符中的句点都被去除。

#### (*ChainNode) Add  <- go1.1

``` go 
func (c *ChainNode) Add(field string)
```

Add adds the named field (which should start with a period) to the end of the chain.

​	Add方法将指定的字段（应以句点开头）添加到链的末尾。

#### (*ChainNode) Copy  <- go1.1

``` go 
func (c *ChainNode) Copy() Node
```

#### (*ChainNode) String  <- go1.1

``` go 
func (c *ChainNode) String() string
```

### type CommandNode 

``` go 
type CommandNode struct {
	NodeType
	Pos

	Args []Node // Arguments in lexical order: Identifier, field, or constant. 参数的词法顺序：标识符、字段或常量。
	// contains filtered or unexported fields
}
```

CommandNode holds a command (a pipeline inside an evaluating action).

​	CommandNode保存一个命令（在评估动作中的管道内部）。

#### (*CommandNode) Copy 

``` go 
func (c *CommandNode) Copy() Node
```

#### (*CommandNode) String 

``` go 
func (c *CommandNode) String() string
```

### type CommentNode  <- go1.16

``` go 
type CommentNode struct {
	NodeType
	Pos

	Text string // Comment text. 注释文本。
	// contains filtered or unexported fields
}
```

CommentNode holds a comment.

​	CommentNode保存一个注释。

#### (*CommentNode) Copy  <- go1.16

``` go 
func (c *CommentNode) Copy() Node
```

#### (*CommentNode) String  <- go1.16

``` go 
func (c *CommentNode) String() string
```

### type ContinueNode  <- go1.18

``` go 
type ContinueNode struct {
	NodeType
	Pos
	Line int
	// contains filtered or unexported fields
}
```

ContinueNode represents a `{{continue}}` action.

​	ContinueNode表示`{{continue}}`动作。

#### (*ContinueNode) Copy  <- go1.18

``` go 
func (c *ContinueNode) Copy() Node
```

#### (*ContinueNode) String  <- go1.18

``` go 
func (c *ContinueNode) String() string
```

### type DotNode 

``` go 
type DotNode struct {
	NodeType
	Pos
	// contains filtered or unexported fields
}
```

DotNode holds the special identifier '.'.

​	DotNode保存特殊标识符'.'。

#### (*DotNode) Copy 

``` go 
func (d *DotNode) Copy() Node
```

#### (*DotNode) String 

``` go 
func (d *DotNode) String() string
```

#### (*DotNode) Type 

``` go 
func (d *DotNode) Type() NodeType
```

### type FieldNode 

``` go 
type FieldNode struct {
	NodeType
	Pos

	Ident []string // The identifiers in lexical order. 以词法顺序的标识符。
	// contains filtered or unexported fields
}
```

FieldNode holds a field (identifier starting with '.'). The names may be chained ('.x.y'). The period is dropped from each ident.

​	FieldNode保存一个字段（以'.'开头的标识符）。名称可以链接（'.x.y'）。每个标识符中的句点都被去除。

#### (*FieldNode) Copy 

``` go 
func (f *FieldNode) Copy() Node
```

#### (*FieldNode) String 

``` go 
func (f *FieldNode) String() string
```

### type IdentifierNode 

``` go 
type IdentifierNode struct {
	NodeType
	Pos

	Ident string // The identifier's name. 标识符的名称。
	// contains filtered or unexported fields
}
```

IdentifierNode holds an identifier.

​	IdentifierNode保存一个标识符。

#### func NewIdentifier 

``` go 
func NewIdentifier(ident string) *IdentifierNode
```

NewIdentifier returns a new IdentifierNode with the given identifier name.

​	NewIdentifier使用给定的标识符名称返回一个新的IdentifierNode。

#### (*IdentifierNode) Copy 

``` go 
func (i *IdentifierNode) Copy() Node
```

#### (*IdentifierNode) SetPos  <- go1.1

``` go 
func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode
```

SetPos sets the position. NewIdentifier is a public method so we can't modify its signature. Chained for convenience. TODO: fix one day?

​	SetPos设置位置。NewIdentifier是一个公共方法，所以我们不能修改它的签名。为了方便链式调用。待修复：TODO。

#### (*IdentifierNode) SetTree  <- go1.4

``` go 
func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode
```

SetTree sets the parent tree for the node. NewIdentifier is a public method so we can't modify its signature. Chained for convenience. TODO: fix one day?

​	SetTree设置节点的父树。NewIdentifier是一个公共方法，所以我们不能修改它的签名。为了方便链式调用。待修复：TODO。

#### (*IdentifierNode) String 

``` go 
func (i *IdentifierNode) String() string
```

### type IfNode 

``` go 
type IfNode struct {
	BranchNode
}
```

IfNode represents an {{if}} action and its commands.

​	IfNode表示{{if}}动作及其命令。

#### (*IfNode) Copy 

``` go 
func (i *IfNode) Copy() Node
```

### type ListNode 

``` go 
type ListNode struct {
	NodeType
	Pos

	Nodes []Node // The element nodes in lexical order. 以词法顺序的元素节点。
	// contains filtered or unexported fields
}
```

ListNode holds a sequence of nodes.

​	ListNode保存一系列节点。

#### (*ListNode) Copy 

``` go 
func (l *ListNode) Copy() Node
```

#### (*ListNode) CopyList 

``` go 
func (l *ListNode) CopyList() *ListNode
```

#### (*ListNode) String 

``` go 
func (l *ListNode) String() string
```

### type Mode  <- go1.16

``` go 
type Mode uint
```

A mode value is a set of flags (or 0). Modes control parser behavior.

​	模式值是一组标志（或0）。模式控制解析器的行为。

``` go 
const (
	ParseComments Mode = 1 << iota // parse comments and add them to AST 解析注释并将其添加到AST
	SkipFuncCheck                  // do not check that functions are defined 不检查函数是否被定义
)
```

### type NilNode  <- go1.1

``` go 
type NilNode struct {
	NodeType
	Pos
	// contains filtered or unexported fields
}
```

NilNode holds the special identifier 'nil' representing an untyped nil constant.

​	NilNode保存特殊标识符'nil'，表示未类型化的nil常量。

#### (*NilNode) Copy  <- go1.1

``` go 
func (n *NilNode) Copy() Node
```

#### (*NilNode) String  <- go1.1

``` go 
func (n *NilNode) String() string
```

#### (*NilNode) Type  <- go1.1

``` go 
func (n *NilNode) Type() NodeType
```

### type Node 

``` go 
type Node interface {
	Type() NodeType
	String() string
	// Copy does a deep copy of the Node and all its components.
	// To avoid type assertions, some XxxNodes also have specialized
	// CopyXxx methods that return *XxxNode.
    // Copy方法对Node及其所有组件进行深拷贝。 为了避免类型断言，某些XxxNodes还具有专门的CopyXxx方法，返回*XxxNode。
	Copy() Node
	Position() Pos // byte position of start of node in full original input string 在完整原始输入字符串中节点起始位置的字节位置
	// contains filtered or unexported methods
}
```

A Node is an element in the parse tree. The interface is trivial. The interface contains an unexported method so that only types local to this package can satisfy it.

​	Node是解析树中的一个元素。该接口是平凡的。该接口包含一个未导出的方法，因此只有本包中的类型才能满足它。

### type NodeType 

``` go 
type NodeType int
```

NodeType identifies the type of a parse tree node.

​	NodeType标识解析树节点的类型。

``` go 
const (
	NodeText    NodeType = iota // Plain text. 纯文本。
	NodeAction                  // A non-control action such as a field evaluation. 非控制操作，如字段评估。
	NodeBool                    // A boolean constant. 布尔常量。
	NodeChain                   // A sequence of field accesses. 字段访问的序列。
	NodeCommand                 // An element of a pipeline. 管道中的元素。
	NodeDot                     // The cursor, dot. 光标，点。

	NodeField      // A field or method name. 字段或方法名。
	NodeIdentifier // An identifier; always a function name. 标识符；始终为函数名。
	NodeIf         // An if action. if操作。
	NodeList       // A list of Nodes. 节点列表。
	NodeNil        // An untyped nil constant. 未类型化的nil常量。
	NodeNumber     // A numerical constant. 数值常量。
	NodePipe       // A pipeline of commands. 命令管道。
	NodeRange      // A range action.  range操作。
	NodeString     // A string constant. 字符串常量。
	NodeTemplate   // A template invocation action. 模板调用操作。
	NodeVariable   // A $ variable. $变量。
	NodeWith       // A with action. with操作。
	NodeComment    // A comment. 注释。
	NodeBreak      // A break action. break操作。
	NodeContinue   // A continue action. continue操作。
)
```

#### (NodeType) Type 

``` go 
func (t NodeType) Type() NodeType
```

Type returns itself and provides an easy default implementation for embedding in a Node. Embedded in all non-trivial Nodes.

​	Type方法返回自身，并为嵌入在所有非平凡节点中提供了简单的默认实现。

### type NumberNode 

``` go 
type NumberNode struct {
	NodeType
	Pos
 
	IsInt      bool       // Number has an integral value. 数字具有整数值。
	IsUint     bool       // Number has an unsigned integral value. 数字具有无符号整数值。
	IsFloat    bool       // Number has a floating-point value. 数字具有浮点数值。
	IsComplex  bool       // Number is complex.  数字是复数。
	Int64      int64      // The signed integer value. 有符号整数值。
	Uint64     uint64     // The unsigned integer value. 无符号整数值。
	Float64    float64    // The floating-point value. 浮点数值。
	Complex128 complex128 // The complex value. 复数值。
	Text       string     // The original textual representation from the input. 输入中的原始文本表示。
	// contains filtered or unexported fields
}
```

NumberNode holds a number: signed or unsigned integer, float, or complex. The value is parsed and stored under all the types that can represent the value. This simulates in a small amount of code the behavior of Go's ideal constants.

​	NumberNode表示一个数字：有符号或无符号整数、浮点数或复数。该值被解析并存储在可以表示该值的所有类型下。这在很少的代码中模拟了Go理想常量的行为。

#### (*NumberNode) Copy 

``` go 
func (n *NumberNode) Copy() Node
```

#### (*NumberNode) String 

``` go 
func (n *NumberNode) String() string
```

### type PipeNode 

``` go 
type PipeNode struct {
	NodeType
	Pos

	Line     int             // The line number in the input. Deprecated: Kept for compatibility. 输入中的行号。已弃用：为了兼容性而保留。
	IsAssign bool            // The variables are being assigned, not declared. 变量正在被赋值，而不是声明。
	Decl     []*VariableNode // Variables in lexical order. 按词法顺序的变量。
	Cmds     []*CommandNode  // The commands in lexical order. 按词法顺序的命令。
	// contains filtered or unexported fields
}
```

PipeNode holds a pipeline with optional declaration

​	PipeNode表示具有可选声明的管道。

#### (*PipeNode) Copy 

``` go 
func (p *PipeNode) Copy() Node
```

#### (*PipeNode) CopyPipe 

``` go 
func (p *PipeNode) CopyPipe() *PipeNode
```

#### (*PipeNode) String 

``` go 
func (p *PipeNode) String() string
```

### type Pos  <- go1.1

``` go 
type Pos int
```

Pos represents a byte position in the original input text from which this template was parsed.

​	Pos表示从中解析此模板的原始输入文本中的字节位置。

#### (Pos) Position  <- go1.1

``` go 
func (p Pos) Position() Pos
```

### type RangeNode 

``` go 
type RangeNode struct {
	BranchNode
}
```

RangeNode represents a {{range}} action and its commands.

​	RangeNode表示一个{{range}}操作及其命令。

#### (*RangeNode) Copy 

``` go 
func (r *RangeNode) Copy() Node
```

### type StringNode 

``` go 
type StringNode struct {
	NodeType
	Pos

	Quoted string // The original text of the string, with quotes. 带引号的字符串的原始文本。
	Text   string // The string, after quote processing. 经过引号处理后的字符串。
	// contains filtered or unexported fields
}
```

StringNode holds a string constant. The value has been "unquoted".

​	StringNode表示一个字符串常量。该值已经被"去引号"。

#### (*StringNode) Copy 

``` go 
func (s *StringNode) Copy() Node
```

#### (*StringNode) String 

``` go 
func (s *StringNode) String() string
```

### type TemplateNode 

``` go 
type TemplateNode struct {
	NodeType
	Pos

	Line int       // The line number in the input. Deprecated: Kept for compatibility. 输入中的行号。已弃用：为了兼容性而保留。
	Name string    // The name of the template (unquoted). 模板的名称（未引号化）。
	Pipe *PipeNode // The command to evaluate as dot for the template. 模板的名称（未引号化）。
	// contains filtered or unexported fields
}
```

TemplateNode represents a {{template}} action.

​	TemplateNode表示一个{{template}}操作。

#### (*TemplateNode) Copy 

``` go 
func (t *TemplateNode) Copy() Node
```

#### (*TemplateNode) String 

``` go 
func (t *TemplateNode) String() string
```

### type TextNode 

``` go 
type TextNode struct {
	NodeType
	Pos

	Text []byte // The text; may span newlines. 文本；可能跨越多行。
	// contains filtered or unexported fields
}
```

TextNode holds plain text.

​	TextNode保存纯文本。

#### (*TextNode) Copy 

``` go 
func (t *TextNode) Copy() Node
```

#### (*TextNode) String 

``` go 
func (t *TextNode) String() string
```

### type Tree 

``` go 
type Tree struct {
	Name      string    // name of the template represented by the tree. 树表示的模板名称。
	ParseName string    // name of the top-level template during parsing, for error messages. 解析过程中顶级模板的名称，用于错误消息。
	Root      *ListNode // top-level root of the tree. 树的顶级根节点。
	Mode      Mode      // parsing mode. 解析模式。
	// contains filtered or unexported fields
}
```

Tree is the representation of a single parsed template.

​	Tree是单个解析模板的表示。

#### func New 

``` go 
func New(name string, funcs ...map[string]any) *Tree
```

New allocates a new parse tree with the given name.

​	New函数分配一个带有给定名称的新解析树。

#### (*Tree) Copy  <- go1.2

``` go 
func (t *Tree) Copy() *Tree
```

Copy returns a copy of the Tree. Any parsing state is discarded.

​	Copy方法返回树的副本。丢弃任何解析状态。

#### (*Tree) ErrorContext  <- go1.1

``` go 
func (t *Tree) ErrorContext(n Node) (location, context string)
```

ErrorContext returns a textual representation of the location of the node in the input text. The receiver is only used when the node does not have a pointer to the tree inside, which can occur in old code.

​	ErrorContext方法返回节点在输入文本中的位置的文本表示。仅当节点没有内部指向树的指针时，接收者才会被使用，这可能发生在旧代码中。

#### (*Tree) Parse 

``` go 
func (t *Tree) Parse(text, leftDelim, rightDelim string, treeSet map[string]*Tree, funcs ...map[string]any) (tree *Tree, err error)
```

Parse parses the template definition string to construct a representation of the template for execution. If either action delimiter string is empty, the default ("{{" or "}}") is used. Embedded template definitions are added to the treeSet map.

​	Parse方法解析模板定义字符串以构建模板的表示形式以供执行。如果任一操作分隔符字符串为空，则使用默认值（"{{"或"}}"）。嵌入的模板定义将添加到treeSet映射中。

### type VariableNode 

``` go 
type VariableNode struct {
	NodeType
	Pos

	Ident []string // Variable name and fields in lexical order. 按词法顺序的变量名和字段。
	// contains filtered or unexported fields
}
```

VariableNode holds a list of variable names, possibly with chained field accesses. The dollar sign is part of the (first) name.

​	VariableNode保存变量名的列表，可能包含链式字段访问。美元符号是（第一个）名称的一部分。

#### (*VariableNode) Copy 

``` go 
func (v *VariableNode) Copy() Node
```

#### (*VariableNode) String 

``` go 
func (v *VariableNode) String() string
```

### type WithNode 

``` go 
type WithNode struct {
	BranchNode
}
```

WithNode represents a `{{with}}` action and its commands.

​	WithNode表示`{{with}}`操作及其命令。

#### (*WithNode) Copy 

``` go 
func (w *WithNode) Copy() Node
```