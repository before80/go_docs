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

#### func IsEmptyTree 

``` go 
func IsEmptyTree(n Node) bool
```

IsEmptyTree reports whether this tree (node) is empty of everything but space or comments.

#### func Parse 

``` go 
func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]any) (map[string]*Tree, error)
```

Parse returns a map from template name to parse.Tree, created by parsing the templates described in the argument string. The top-level template will be given the specified name. If an error is encountered, parsing stops and an empty map is returned with the error.

## 类型

### type ActionNode 

``` go 
type ActionNode struct {
	NodeType
	Pos

	Line int       // The line number in the input. Deprecated: Kept for compatibility.
	Pipe *PipeNode // The pipeline in the action.
	// contains filtered or unexported fields
}
```

ActionNode holds an action (something bounded by delimiters). Control actions have their own nodes; ActionNode represents simple ones such as field evaluations and parenthesized pipelines.

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

	True bool // The value of the boolean constant.
	// contains filtered or unexported fields
}
```

BoolNode holds a boolean constant.

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

	Line     int       // The line number in the input. Deprecated: Kept for compatibility.
	Pipe     *PipeNode // The pipeline to be evaluated.
	List     *ListNode // What to execute if the value is non-empty.
	ElseList *ListNode // What to execute if the value is empty (nil if absent).
	// contains filtered or unexported fields
}
```

BranchNode is the common representation of if, range, and with.

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
	Field []string // The identifiers in lexical order.
	// contains filtered or unexported fields
}
```

ChainNode holds a term followed by a chain of field accesses (identifier starting with '.'). The names may be chained ('.x.y'). The periods are dropped from each ident.

#### (*ChainNode) Add  <- go1.1

``` go 
func (c *ChainNode) Add(field string)
```

Add adds the named field (which should start with a period) to the end of the chain.

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

	Args []Node // Arguments in lexical order: Identifier, field, or constant.
	// contains filtered or unexported fields
}
```

CommandNode holds a command (a pipeline inside an evaluating action).

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

	Text string // Comment text.
	// contains filtered or unexported fields
}
```

CommentNode holds a comment.

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

ContinueNode represents a {{continue}} action.

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

	Ident []string // The identifiers in lexical order.
	// contains filtered or unexported fields
}
```

FieldNode holds a field (identifier starting with '.'). The names may be chained ('.x.y'). The period is dropped from each ident.

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

	Ident string // The identifier's name.
	// contains filtered or unexported fields
}
```

IdentifierNode holds an identifier.

#### func NewIdentifier 

``` go 
func NewIdentifier(ident string) *IdentifierNode
```

NewIdentifier returns a new IdentifierNode with the given identifier name.

#### (*IdentifierNode) Copy 

``` go 
func (i *IdentifierNode) Copy() Node
```

#### (*IdentifierNode) SetPos  <- go1.1

``` go 
func (i *IdentifierNode) SetPos(pos Pos) *IdentifierNode
```

SetPos sets the position. NewIdentifier is a public method so we can't modify its signature. Chained for convenience. TODO: fix one day?

#### (*IdentifierNode) SetTree  <- go1.4

``` go 
func (i *IdentifierNode) SetTree(t *Tree) *IdentifierNode
```

SetTree sets the parent tree for the node. NewIdentifier is a public method so we can't modify its signature. Chained for convenience. TODO: fix one day?

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

#### (*IfNode) Copy 

``` go 
func (i *IfNode) Copy() Node
```

### type ListNode 

``` go 
type ListNode struct {
	NodeType
	Pos

	Nodes []Node // The element nodes in lexical order.
	// contains filtered or unexported fields
}
```

ListNode holds a sequence of nodes.

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

``` go 
const (
	ParseComments Mode = 1 << iota // parse comments and add them to AST
	SkipFuncCheck                  // do not check that functions are defined
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
	Copy() Node
	Position() Pos // byte position of start of node in full original input string
	// contains filtered or unexported methods
}
```

A Node is an element in the parse tree. The interface is trivial. The interface contains an unexported method so that only types local to this package can satisfy it.

### type NodeType 

``` go 
type NodeType int
```

NodeType identifies the type of a parse tree node.

``` go 
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

#### (NodeType) Type 

``` go 
func (t NodeType) Type() NodeType
```

Type returns itself and provides an easy default implementation for embedding in a Node. Embedded in all non-trivial Nodes.

### type NumberNode 

``` go 
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

	Line     int             // The line number in the input. Deprecated: Kept for compatibility.
	IsAssign bool            // The variables are being assigned, not declared.
	Decl     []*VariableNode // Variables in lexical order.
	Cmds     []*CommandNode  // The commands in lexical order.
	// contains filtered or unexported fields
}
```

PipeNode holds a pipeline with optional declaration

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

#### (*RangeNode) Copy 

``` go 
func (r *RangeNode) Copy() Node
```

### type StringNode 

``` go 
type StringNode struct {
	NodeType
	Pos

	Quoted string // The original text of the string, with quotes.
	Text   string // The string, after quote processing.
	// contains filtered or unexported fields
}
```

StringNode holds a string constant. The value has been "unquoted".

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

	Line int       // The line number in the input. Deprecated: Kept for compatibility.
	Name string    // The name of the template (unquoted).
	Pipe *PipeNode // The command to evaluate as dot for the template.
	// contains filtered or unexported fields
}
```

TemplateNode represents a {{template}} action.

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

	Text []byte // The text; may span newlines.
	// contains filtered or unexported fields
}
```

TextNode holds plain text.

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
	Name      string    // name of the template represented by the tree.
	ParseName string    // name of the top-level template during parsing, for error messages.
	Root      *ListNode // top-level root of the tree.
	Mode      Mode      // parsing mode.
	// contains filtered or unexported fields
}
```

Tree is the representation of a single parsed template.

#### func New 

``` go 
func New(name string, funcs ...map[string]any) *Tree
```

New allocates a new parse tree with the given name.

#### (*Tree) Copy  <- go1.2

``` go 
func (t *Tree) Copy() *Tree
```

Copy returns a copy of the Tree. Any parsing state is discarded.

#### (*Tree) ErrorContext  <- go1.1

``` go 
func (t *Tree) ErrorContext(n Node) (location, context string)
```

ErrorContext returns a textual representation of the location of the node in the input text. The receiver is only used when the node does not have a pointer to the tree inside, which can occur in old code.

#### (*Tree) Parse 

``` go 
func (t *Tree) Parse(text, leftDelim, rightDelim string, treeSet map[string]*Tree, funcs ...map[string]any) (tree *Tree, err error)
```

Parse parses the template definition string to construct a representation of the template for execution. If either action delimiter string is empty, the default ("{{" or "}}") is used. Embedded template definitions are added to the treeSet map.

### type VariableNode 

``` go 
type VariableNode struct {
	NodeType
	Pos

	Ident []string // Variable name and fields in lexical order.
	// contains filtered or unexported fields
}
```

VariableNode holds a list of variable names, possibly with chained field accesses. The dollar sign is part of the (first) name.

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

WithNode represents a {{with}} action and its commands.

#### (*WithNode) Copy 

``` go 
func (w *WithNode) Copy() Node
```