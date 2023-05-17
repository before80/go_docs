+++
title = "template/parse"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# parse

https://pkg.go.dev/text/template/parse@go1.20.1



Package parse builds parse trees for templates as defined by text/template and html/template. Clients should use those packages to construct templates rather than this one, which provides shared internal data structures not intended for general use.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [IsEmptyTree](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/parse.go;l=270) 

``` go linenums="1"
func IsEmptyTree(n Node) bool
```

IsEmptyTree reports whether this tree (node) is empty of everything but space or comments.

#### func [Parse](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/parse.go;l=62) 

``` go linenums="1"
func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]any) (map[string]*Tree, error)
```

Parse returns a map from template name to parse.Tree, created by parsing the templates described in the argument string. The top-level template will be given the specified name. If an error is encountered, parsing stops and an empty map is returned with the error.

## 类型

### type [ActionNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=257) 

``` go linenums="1"
type ActionNode struct {
	NodeType
	Pos

	Line int       // The line number in the input. Deprecated: Kept for compatibility.
	Pipe *PipeNode // The pipeline in the action.
	// contains filtered or unexported fields
}
```

ActionNode holds an action (something bounded by delimiters). Control actions have their own nodes; ActionNode represents simple ones such as field evaluations and parenthesized pipelines.

#### (*ActionNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=285) 

``` go linenums="1"
func (a *ActionNode) Copy() Node
```

#### (*ActionNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=269) 

``` go linenums="1"
func (a *ActionNode) String() string
```

### type [BoolNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=582) 

``` go linenums="1"
type BoolNode struct {
	NodeType
	Pos

	True bool // The value of the boolean constant.
	// contains filtered or unexported fields
}
```

BoolNode holds a boolean constant.

#### (*BoolNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=608) 

``` go linenums="1"
func (b *BoolNode) Copy() Node
```

#### (*BoolNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=593) 

``` go linenums="1"
func (b *BoolNode) String() string
```

### type [BranchNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=841) 

``` go linenums="1"
type BranchNode struct {
	NodeType
	Pos

	Line     int       // The line number in the input. Deprecated: Kept for compatibility.
	Pipe     *PipeNode // The pipeline to be evaluated.
	List     *ListNode // What to execute if the value is non-empty.
	ElseList *ListNode // What to execute if the value is empty (nil if absent).
	// contains filtered or unexported fields
}
```

BranchNode is the common representation of if, range, and with.

#### (*BranchNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=886)  <- go1.4

``` go linenums="1"
func (b *BranchNode) Copy() Node
```

#### (*BranchNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=851) 

``` go linenums="1"
func (b *BranchNode) String() string
```

### type [BreakNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=913)  <- go1.18

``` go linenums="1"
type BreakNode struct {
	NodeType
	Pos
	Line int
	// contains filtered or unexported fields
}
```

BreakNode represents a {{break}} action.

#### (*BreakNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=924)  <- go1.18

``` go linenums="1"
func (b *BreakNode) Copy() Node
```

#### (*BreakNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=925)  <- go1.18

``` go linenums="1"
func (b *BreakNode) String() string
```

### type [ChainNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=529)  <- go1.1

``` go linenums="1"
type ChainNode struct {
	NodeType
	Pos

	Node  Node
	Field []string // The identifiers in lexical order.
	// contains filtered or unexported fields
}
```

ChainNode holds a term followed by a chain of field accesses (identifier starting with '.'). The names may be chained ('.x.y'). The periods are dropped from each ident.

#### (*ChainNode) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=542)  <- go1.1

``` go linenums="1"
func (c *ChainNode) Add(field string)
```

Add adds the named field (which should start with a period) to the end of the chain.

#### (*ChainNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=577)  <- go1.1

``` go linenums="1"
func (c *ChainNode) Copy() Node
```

#### (*ChainNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=553)  <- go1.1

``` go linenums="1"
func (c *ChainNode) String() string
```

### type [CommandNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=291) 

``` go linenums="1"
type CommandNode struct {
	NodeType
	Pos

	Args []Node // Arguments in lexical order: Identifier, field, or constant.
	// contains filtered or unexported fields
}
```

CommandNode holds a command (a pipeline inside an evaluating action).

#### (*CommandNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=331) 

``` go linenums="1"
func (c *CommandNode) Copy() Node
```

#### (*CommandNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=306) 

``` go linenums="1"
func (c *CommandNode) String() string
```

### type [CommentNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=156)  <- go1.16

``` go linenums="1"
type CommentNode struct {
	NodeType
	Pos

	Text string // Comment text.
	// contains filtered or unexported fields
}
```

CommentNode holds a comment.

#### (*CommentNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=183)  <- go1.16

``` go linenums="1"
func (c *CommentNode) Copy() Node
```

#### (*CommentNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=167)  <- go1.16

``` go linenums="1"
func (c *CommentNode) String() string
```

### type [ContinueNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=930)  <- go1.18

``` go linenums="1"
type ContinueNode struct {
	NodeType
	Pos
	Line int
	// contains filtered or unexported fields
}
```

ContinueNode represents a {{continue}} action.

#### (*ContinueNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=941)  <- go1.18

``` go linenums="1"
func (c *ContinueNode) Copy() Node
```

#### (*ContinueNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=942)  <- go1.18

``` go linenums="1"
func (c *ContinueNode) String() string
```

### type [DotNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=424) 

``` go linenums="1"
type DotNode struct {
	NodeType
	Pos
	// contains filtered or unexported fields
}
```

DotNode holds the special identifier '.'.

#### (*DotNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=453) 

``` go linenums="1"
func (d *DotNode) Copy() Node
```

#### (*DotNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=441) 

``` go linenums="1"
func (d *DotNode) String() string
```

#### (*DotNode) [Type](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=434) 

``` go linenums="1"
func (d *DotNode) Type() NodeType
```

### type [FieldNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=494) 

``` go linenums="1"
type FieldNode struct {
	NodeType
	Pos

	Ident []string // The identifiers in lexical order.
	// contains filtered or unexported fields
}
```

FieldNode holds a field (identifier starting with '.'). The names may be chained ('.x.y'). The period is dropped from each ident.

#### (*FieldNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=522) 

``` go linenums="1"
func (f *FieldNode) Copy() Node
```

#### (*FieldNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=505) 

``` go linenums="1"
func (f *FieldNode) String() string
```

### type [IdentifierNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=343) 

``` go linenums="1"
type IdentifierNode struct {
	NodeType
	Pos

	Ident string // The identifier's name.
	// contains filtered or unexported fields
}
```

IdentifierNode holds an identifier.

#### func [NewIdentifier](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=351) 

``` go linenums="1"
func NewIdentifier(ident string) *IdentifierNode
```

NewIdentifier returns a new IdentifierNode with the given identifier name.

#### (*IdentifierNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=383) 

``` go linenums="1"
func (i *IdentifierNode) Copy() Node
```

#### (*IdentifierNode) [SetPos](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=358)  <- go1.1

``` go linenums="1"
func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode
```

SetPos sets the position. NewIdentifier is a public method so we can't modify its signature. Chained for convenience. TODO: fix one day?

#### (*IdentifierNode) [SetTree](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=366)  <- go1.4

``` go linenums="1"
func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode
```

SetTree sets the parent tree for the node. NewIdentifier is a public method so we can't modify its signature. Chained for convenience. TODO: fix one day?

#### (*IdentifierNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=371) 

``` go linenums="1"
func (i *IdentifierNode) String() string
```

### type [IfNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=900) 

``` go linenums="1"
type IfNode struct {
	BranchNode
}
```

IfNode represents an {{if}} action and its commands.

#### (*IfNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=908) 

``` go linenums="1"
func (i *IfNode) Copy() Node
```

### type [ListNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=81) 

``` go linenums="1"
type ListNode struct {
	NodeType
	Pos

	Nodes []Node // The element nodes in lexical order.
	// contains filtered or unexported fields
}
```

ListNode holds a sequence of nodes.

#### (*ListNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=123) 

``` go linenums="1"
func (l *ListNode) Copy() Node
```

#### (*ListNode) [CopyList](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=112) 

``` go linenums="1"
func (l *ListNode) CopyList() *ListNode
```

#### (*ListNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=100) 

``` go linenums="1"
func (l *ListNode) String() string
```

### type [Mode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/parse.go;l=38)  <- go1.16

``` go linenums="1"
type Mode uint
```

A mode value is a set of flags (or 0). Modes control parser behavior.

``` go linenums="1"
const (
	ParseComments Mode = 1 << iota // parse comments and add them to AST
	SkipFuncCheck                  // do not check that functions are defined
)
```

### type [NilNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=458)  <- go1.1

``` go linenums="1"
type NilNode struct {
	NodeType
	Pos
	// contains filtered or unexported fields
}
```

NilNode holds the special identifier 'nil' representing an untyped nil constant.

#### (*NilNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=487)  <- go1.1

``` go linenums="1"
func (n *NilNode) Copy() Node
```

#### (*NilNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=475)  <- go1.1

``` go linenums="1"
func (n *NilNode) String() string
```

#### (*NilNode) [Type](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=468)  <- go1.1

``` go linenums="1"
func (n *NilNode) Type() NodeType
```

### type [Node](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=20) 

``` go linenums="1"
type Node interface {
	Type() NodeType
	String() string
	// Copy does a deep copy of the Node and all its components.
	// To avoid type assertions, some XxxNodes also have specialized
	// CopyXxx methods that return *XxxNode.
	Copy() Node
	Position() Pos // byte position of start of node in full original input string
	// contains filtered or unexported methods
}
```

A Node is an element in the parse tree. The interface is trivial. The interface contains an unexported method so that only types local to this package can satisfy it.

### type [NodeType](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=36) 

``` go linenums="1"
type NodeType int
```

NodeType identifies the type of a parse tree node.

``` go linenums="1"
const (
	NodeText    NodeType = iota // Plain text.
	NodeAction                  // A non-control action such as a field evaluation.
	NodeBool                    // A boolean constant.
	NodeChain                   // A sequence of field accesses.
	NodeCommand                 // An element of a pipeline.
	NodeDot                     // The cursor, dot.

	NodeField      // A field or method name.
	NodeIdentifier // An identifier; always a function name.
	NodeIf         // An if action.
	NodeList       // A list of Nodes.
	NodeNil        // An untyped nil constant.
	NodeNumber     // A numerical constant.
	NodePipe       // A pipeline of commands.
	NodeRange      // A range action.
	NodeString     // A string constant.
	NodeTemplate   // A template invocation action.
	NodeVariable   // A $ variable.
	NodeWith       // A with action.
	NodeComment    // A comment.
	NodeBreak      // A break action.
	NodeContinue   // A continue action.
)
```

#### (NodeType) [Type](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=48) 

``` go linenums="1"
func (t NodeType) Type() NodeType
```

Type returns itself and provides an easy default implementation for embedding in a Node. Embedded in all non-trivial Nodes.

### type [NumberNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=615) 

``` go linenums="1"
type NumberNode struct {
	NodeType
	Pos

	IsInt      bool       // Number has an integral value.
	IsUint     bool       // Number has an unsigned integral value.
	IsFloat    bool       // Number has a floating-point value.
	IsComplex  bool       // Number is complex.
	Int64      int64      // The signed integer value.
	Uint64     uint64     // The unsigned integer value.
	Float64    float64    // The floating-point value.
	Complex128 complex128 // The complex value.
	Text       string     // The original textual representation from the input.
	// contains filtered or unexported fields
}
```

NumberNode holds a number: signed or unsigned integer, float, or complex. The value is parsed and stored under all the types that can represent the value. This simulates in a small amount of code the behavior of Go's ideal constants.

#### (*NumberNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=745) 

``` go linenums="1"
func (n *NumberNode) Copy() Node
```

#### (*NumberNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=733) 

``` go linenums="1"
func (n *NumberNode) String() string
```

### type [PipeNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=188) 

``` go linenums="1"
type PipeNode struct {
	NodeType
	Pos

	Line     int             // The line number in the input. Deprecated: Kept for compatibility.
	IsAssign bool            // The variables are being assigned, not declared.
	Decl     []*VariableNode // Variables in lexical order.
	Cmds     []*CommandNode  // The commands in lexical order.
	// contains filtered or unexported fields
}
```

PipeNode holds a pipeline with optional declaration

#### (*PipeNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=250) 

``` go linenums="1"
func (p *PipeNode) Copy() Node
```

#### (*PipeNode) [CopyPipe](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=234) 

``` go linenums="1"
func (p *PipeNode) CopyPipe() *PipeNode
```

#### (*PipeNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=206) 

``` go linenums="1"
func (p *PipeNode) String() string
```

### type [Pos](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=40)  <- go1.1

``` go linenums="1"
type Pos int
```

Pos represents a byte position in the original input text from which this template was parsed.

#### (Pos) [Position](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=42)  <- go1.1

``` go linenums="1"
func (p Pos) Position() Pos
```

### type [RangeNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=947) 

``` go linenums="1"
type RangeNode struct {
	BranchNode
}
```

RangeNode represents a {{range}} action and its commands.

#### (*RangeNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=955) 

``` go linenums="1"
func (r *RangeNode) Copy() Node
```

### type [StringNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=752) 

``` go linenums="1"
type StringNode struct {
	NodeType
	Pos

	Quoted string // The original text of the string, with quotes.
	Text   string // The string, after quote processing.
	// contains filtered or unexported fields
}
```

StringNode holds a string constant. The value has been "unquoted".

#### (*StringNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=776) 

``` go linenums="1"
func (s *StringNode) Copy() Node
```

#### (*StringNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=764) 

``` go linenums="1"
func (s *StringNode) String() string
```

### type [TemplateNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=973) 

``` go linenums="1"
type TemplateNode struct {
	NodeType
	Pos

	Line int       // The line number in the input. Deprecated: Kept for compatibility.
	Name string    // The name of the template (unquoted).
	Pipe *PipeNode // The command to evaluate as dot for the template.
	// contains filtered or unexported fields
}
```

TemplateNode represents a {{template}} action.

#### (*TemplateNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=1006) 

``` go linenums="1"
func (t *TemplateNode) Copy() Node
```

#### (*TemplateNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=986) 

``` go linenums="1"
func (t *TemplateNode) String() string
```

### type [TextNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=128) 

``` go linenums="1"
type TextNode struct {
	NodeType
	Pos

	Text []byte // The text; may span newlines.
	// contains filtered or unexported fields
}
```

TextNode holds plain text.

#### (*TextNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=151) 

``` go linenums="1"
func (t *TextNode) Copy() Node
```

#### (*TextNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=139) 

``` go linenums="1"
func (t *TextNode) String() string
```

### type [Tree](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/parse.go;l=20) 

``` go linenums="1"
type Tree struct {
	Name      string    // name of the template represented by the tree.
	ParseName string    // name of the top-level template during parsing, for error messages.
	Root      *ListNode // top-level root of the tree.
	Mode      Mode      // parsing mode.
	// contains filtered or unexported fields
}
```

Tree is the representation of a single parsed template.

#### func [New](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/parse.go;l=131) 

``` go linenums="1"
func New(name string, funcs ...map[string]any) *Tree
```

New allocates a new parse tree with the given name.

#### (*Tree) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/parse.go;l=46)  <- go1.2

``` go linenums="1"
func (t *Tree) Copy() *Tree
```

Copy returns a copy of the Tree. Any parsing state is discarded.

#### (*Tree) [ErrorContext](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/parse.go;l=141)  <- go1.1

``` go linenums="1"
func (t *Tree) ErrorContext(n Node) (location, context string)
```

ErrorContext returns a textual representation of the location of the node in the input text. The receiver is only used when the node does not have a pointer to the tree inside, which can occur in old code.

#### (*Tree) [Parse](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/parse.go;l=245) 

``` go linenums="1"
func (t *Tree) Parse(text, leftDelim, rightDelim string, treeSet map[string]*Tree, funcs ...map[string]any) (tree *Tree, err error)
```

Parse parses the template definition string to construct a representation of the template for execution. If either action delimiter string is empty, the default ("{{" or "}}") is used. Embedded template definitions are added to the treeSet map.

### type [VariableNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=389) 

``` go linenums="1"
type VariableNode struct {
	NodeType
	Pos

	Ident []string // Variable name and fields in lexical order.
	// contains filtered or unexported fields
}
```

VariableNode holds a list of variable names, possibly with chained field accesses. The dollar sign is part of the (first) name.

#### (*VariableNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=419) 

``` go linenums="1"
func (v *VariableNode) Copy() Node
```

#### (*VariableNode) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=400) 

``` go linenums="1"
func (v *VariableNode) String() string
```

### type [WithNode](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=960) 

``` go linenums="1"
type WithNode struct {
	BranchNode
}
```

WithNode represents a {{with}} action and its commands.

#### (*WithNode) [Copy](https://cs.opensource.google/go/go/+/go1.20.1:src/text/template/parse/node.go;l=968) 

``` go linenums="1"
func (w *WithNode) Copy() Node
```