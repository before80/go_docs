+++
title = "syntax"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/regexp/syntax@go1.21.3](https://pkg.go.dev/regexp/syntax@go1.21.3)

Package syntax parses regular expressions into parse trees and compiles parse trees into programs. Most clients of regular expressions will use the facilities of package regexp (such as Compile and Match) instead of this package.

​	syntax 包将正则表达式解析为解析树，并将解析树编译为程序。大多数正则表达式的客户端将使用 regexp 包的工具(例如 Compile 和 Match)而不是此包。

## 语法 Syntax 

The regular expression syntax understood by this package when parsing with the Perl flag is as follows. Parts of the syntax can be disabled by passing alternate flags to Parse.

​	当使用 Perl 标志进行解析时，此包理解的正则表达式语法如下。通过将替代标志传递给 Parse，可以禁用语法的某些部分。

Single characters:

单个字符：

```
.              任何字符，可能包括换行符(标志 s=true)
[xyz]          字符类
[^xyz]         反向字符类
\d             Perl字符类
\D             反向 Perl 字符类
[[:alpha:]]    ASCII 字符类
[[:^alpha:]]   反向 ASCII 字符类
\pN             Unicode 字符类(单字母名称)
\p{Greek}      Unicode 字符类
\PN            反向 Unicode 字符类(单字母名称)
\P{Greek}      反向 Unicode 字符类
```

Composites:

复合：

```
xy             x 后跟 y
x|y            x 或 y(更偏向于 x)
```

Repetitions:

重复：

```
x*             零个或多个 x，更偏向于多个
x+             一个或多个 x，更偏向于多个
x?             零个或一个 x，更偏向于一个
x{n,m}         n 或 n+1 或 ... 或 m 个 x，更偏向于多个
x{n,}          n 或更多个 x，更偏向于多个
x{n}           恰好 n 个 x
x*?            零个或多个 x，更偏向于较少
x+?            一个或多个 x，更偏向于较少
x??            零个或一个 x，更偏向于零
x{n,m}?        n 或 n+1 或 ... 或 m 个 x，更偏向于较少
x{n,}?         n 或更多个 x，更偏向于较少
x{n}?          恰好 n 个 x
```

Implementation restriction: The counting forms x{n,m}, x{n,}, and x{n} reject forms that create a minimum or maximum repetition count above 1000. Unlimited repetitions are not subject to this restriction.

​	实现限制：计数形式 `x{n,m}`、`x{n,}` 和 `x{n}` 拒绝创建最小或最大重复计数大于 1000 的形式。无限重复不受此限制。

Grouping:

分组：

```
(re)           编号捕获组(子匹配)
(?P<name>re)   命名和编号捕获组(子匹配)
(?:re)         非捕获组
(?flags)       在当前组内设置标志；非捕获
(?flags:re)    在 re 中设置标志；非捕获

标志语法为 xyz(设置)或 -xyz(清除)或 xy-z(设置 xy，清除 z)。
标志包括：

i              不区分大小写(默认为 false)
m              多行模式：^ 和 $ 匹配行的开头和结尾以及文本的开头和结尾(默认为 false)
s              让 . 匹配 \n(默认为 false)
U              非贪婪模式：交换 x* 和 x*？x+ 和 x+？等的含义(默认为 false)
```

Empty strings:

空字符串：

```
^              文本或行的开头(标志 m=true)
$              文本的结尾(类似于 \z 而不是 \Z)或行的结尾(标志 m=true)
\A             文本的开头
\b             ASCII 单词边界(\w 在一侧，\W、\A 或 \z 在另一侧)
\B             非 ASCII 单词边界
\z             文本的结尾
```

Escape sequences:

转义序列：

```
\a             响铃(== \007)
\f             换页符(== \014)
\t             水平制表符(== \011)
\n             换行符(== \012)
\r             回车符(== \015)
\v             垂直制表符(== \013)
\*             字符 *(对于任何标点符号字符 * 都是字面量)
\123           八进制字符代码(最多三位数字)
\x7F           十六进制字符代码(恰好两位数字)
\x{10FFFF}     十六进制字符代码
\Q...\E        字面文本...即使...中有标点符号
```

Character class elements:

字符类元素：

```
x              单个字符
A-Z            字符范围(包括)
\d             Perl 字符类
[:foo:]        ASCII 字符类 foo
\p{Foo}        Unicode 字符类 Foo
\pF            字符类 F(一个字母名称)
```

Named character classes as character class elements:

命名字符类作为字符类元素：

```
[\d]           数字(== \d)
[^\d]          非数字(== \D)
[\D]           非数字(== \D)
[^\D]          非非数字(== \d)
[[:name:]]     字符类内部的命名 ASCII 类(== [:name:])
[^[:name:]]    在否定字符类中命名 ASCII 类(== [:^name:])
[\p{Name}]     字符类内的命名 Unicode 属性(== \p{Name})
[^\p{Name}]    在否定字符类中的命名 Unicode 属性(== \P{Name})
```

Perl character classes (all ASCII-only):

Perl字符类(仅ASCII)：

```
\d             数字(== [0-9])
\D             非数字(== [^0-9])
\s             空格字符(== [\t\n\f\r ])
\S             非空格字符
\w             单词字符 (== [0-9A-Za-z_]) 
\W             非单词字符 (== [^0-9A-Za-z_])
```

ASCII character classes:

ASCII字符类：

```
[[:alnum:]]    字母数字 (== [0-9A-Za-z])
[[:alpha:]]    字母 (== [A-Za-z])
[[:ascii:]]    ASCII字符 (== [\x00-\x7F])
[[:blank:]]    空格或制表符 (== [\t])
[[:cntrl:]]    控制字符 (== [\x00-\x1F\x7F])
[[:digit:]]    数字 (== [0-9])
[[:graph:]]    图形字符 (== [!-~] == [A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\]^_`{|}~] )
[[:lower:]]    小写字母 (== [a-z])
[[:print:]]    可打印字符 (== [ -~] == [ [:graph:]] )
[[:punct:]]    punctuation (== [!-/:-@[-`{-~]) 标点符号 (== [!-/:-@[-`{-~] )
[[:space:]]    空格字符 (== [\t\n\v\f\r])
[[:upper:]]    大写字母 (== [A-Z]) 大写 (== [A-Z])
[[:word:]]     单词字符 (== [0-9A-Za-z_])
[[:xdigit:]]   十六进制数字 (== [0-9A-Fa-f])
```

Unicode character classes are those in unicode.Categories and unicode.Scripts.

​	Unicode字符类包括unicode.Categories和unicode.Scripts中的类别。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func IsWordChar 

``` go 
func IsWordChar(r rune) bool
```

IsWordChar reports whether r is considered a “word character” during the evaluation of the \b and \B zero-width assertions. These assertions are ASCII-only: the word characters are [A-Za-z0-9_].

​	`IsWordChar`函数报告 r 是否在评估 \b 和 \B 零宽断言期间被认为是"单词字符"。这些断言仅适用于 ASCII：单词字符是 `[A-Za-z0-9_]`。

## 类型

### type EmptyOp 

``` go 
type EmptyOp uint8
```

An EmptyOp specifies a kind or mixture of zero-width assertions.

​	`EmptyOp` 指定零宽断言的一种或多种种类。

``` go 
const (
	EmptyBeginLine EmptyOp = 1 << iota
	EmptyEndLine
	EmptyBeginText
	EmptyEndText
	EmptyWordBoundary
	EmptyNoWordBoundary
)
```

#### func EmptyOpContext 

``` go 
func EmptyOpContext(r1, r2 rune) EmptyOp
```

EmptyOpContext returns the zero-width assertions satisfied at the position between the runes r1 and r2. Passing r1 == -1 indicates that the position is at the beginning of the text. Passing r2 == -1 indicates that the position is at the end of the text.

​	`EmptyOpContext`函数返回在符文 r1 和 r2 之间位置的满足零宽断言的内容。传递 r1 == -1 表示位置位于文本开头。传递 r2 == -1 表示位置位于文本末尾。

### type Error 

``` go 
type Error struct {
	Code ErrorCode
	Expr string
}
```

An Error describes a failure to parse a regular expression and gives the offending expression.

​	`Error`结构体描述了无法解析正则表达式的失败，并提供了出错的表达式。

#### (*Error) Error 

``` go 
func (e *Error) Error() string
```

### type ErrorCode 

``` go 
type ErrorCode string
```

An ErrorCode describes a failure to parse a regular expression.

​	`ErrorCode` 描述了无法解析正则表达式的失败。

``` go 
const (
    // Unexpected error
	// 意外的错误
	ErrInternalError ErrorCode = "regexp/syntax: internal error"

    // Parse errors
	// 解析错误
	ErrInvalidCharClass      ErrorCode = "invalid character class"
	ErrInvalidCharRange      ErrorCode = "invalid character class range"
	ErrInvalidEscape         ErrorCode = "invalid escape sequence"
	ErrInvalidNamedCapture   ErrorCode = "invalid named capture"
	ErrInvalidPerlOp         ErrorCode = "invalid or unsupported Perl syntax"
	ErrInvalidRepeatOp       ErrorCode = "invalid nested repetition operator"
	ErrInvalidRepeatSize     ErrorCode = "invalid repeat count"
	ErrInvalidUTF8           ErrorCode = "invalid UTF-8"
	ErrMissingBracket        ErrorCode = "missing closing ]"
	ErrMissingParen          ErrorCode = "missing closing )"
	ErrMissingRepeatArgument ErrorCode = "missing argument to repetition operator"
	ErrTrailingBackslash     ErrorCode = "trailing backslash at end of expression"
	ErrUnexpectedParen       ErrorCode = "unexpected )"
	ErrNestingDepth          ErrorCode = "expression nests too deeply"
	ErrLarge                 ErrorCode = "expression too large"
)
```

#### (ErrorCode) String 

``` go 
func (e ErrorCode) String() string
```

### type Flags 

``` go 
type Flags uint16
```

Flags control the behavior of the parser and record information about regexp context.

​	`Flags` 控制解析器的行为并记录有关正则表达式上下文的信息。

``` go 
const (
	FoldCase      Flags = 1 << iota // 不区分大小写匹配 case-insensitive match
	Literal                         // 将模式视为字面字符串 treat pattern as literal string
	ClassNL  // 允许字符类(如[^a-z]和[[:space:]])匹配换行符 allow character classes like [^a-z] and [[:space:]] to match newline
	DotNL                           // 允许"."匹配换行符 allow . to match newline
	OneLine // 将 "^" 和 "$" 视为只匹配文本开头和结尾 treat ^ and $ as only matching at beginning and end of text
	NonGreedy // 重复操作符默认为非贪婪 make repetition operators default to non-greedy
	PerlX  // 允许Perl扩展 allow Perl extensions
	UnicodeGroups // 允许\p{Han}、\P{Han}等Unicode组和否定 allow \p{Han}, \P{Han} for Unicode group and negation
	WasDollar // 允许\p{Han}、\P{Han}等Unicode组和否定 regexp OpEndText was $, not \z
	Simple  // regexp不包含计数的重复 regexp contains no counted repetition

	MatchNL = ClassNL | DotNL

	Perl = ClassNL | OneLine | PerlX | UnicodeGroups //尽可能接近Perl as close to Perl as possible
	POSIX Flags = 0    //POSIX语法 POSIX syntax
)
```

### type Inst 

``` go 
type Inst struct {
	Op   InstOp
	Out  uint32 // 所有操作除了InstMatch和InstFail all but InstMatch, InstFail
	Arg  uint32 // InstAlt, InstAltMatch, InstCapture, InstEmptyWidth InstAlt, InstAltMatch, InstCapture, InstEmptyWidth
	Rune []rune
}
```

An Inst is a single instruction in a regular expression program.

​	`Inst`是正则表达式程序中的单个指令。

#### (*Inst) MatchEmptyWidth 

``` go 
func (i *Inst) MatchEmptyWidth(before rune, after rune) bool
```

MatchEmptyWidth reports whether the instruction matches an empty string between the runes before and after. It should only be called when i.Op == InstEmptyWidth.

​	`MatchEmptyWidth`方法报告指令是否在before和after之间匹配空字符串。仅在i.Op == InstEmptyWidth时才应调用它。

#### (*Inst) MatchRune 

``` go 
func (i *Inst) MatchRune(r rune) bool
```

MatchRune reports whether the instruction matches (and consumes) r. It should only be called when i.Op == InstRune.

​	`MatchRune`方法报告指令是否与r匹配(并消耗r)。仅在i.Op == InstRune时才应调用它。

#### (*Inst) MatchRunePos  <- go1.3

``` go 
func (i *Inst) MatchRunePos(r rune) int
```

MatchRunePos checks whether the instruction matches (and consumes) r. If so, MatchRunePos returns the index of the matching rune pair (or, when len(i.Rune) == 1, rune singleton). If not, MatchRunePos returns -1. MatchRunePos should only be called when i.Op == InstRune.

​	`MatchRunePos`方法检查指令是否与r匹配(并消耗r)。如果是，则MatchRunePos返回匹配符号对的索引(或者当len(i.Rune) == 1时，返回符号单例)。如果不是，则MatchRunePos返回-1。MatchRunePos仅在i.Op == InstRune时应调用。

#### (*Inst) String 

``` go 
func (i *Inst) String() string
```

### type InstOp 

``` go 
type InstOp uint8
```

An InstOp is an instruction opcode.

​	InstOp是一个指令操作码。

``` go 
const (
	InstAlt InstOp = iota
	InstAltMatch
	InstCapture
	InstEmptyWidth
	InstMatch
	InstFail
	InstNop
	InstRune
	InstRune1
	InstRuneAny
	InstRuneAnyNotNL
)
```

#### (InstOp) String  <- go1.3

``` go 
func (i InstOp) String() string
```

### type Op 

``` go 
type Op uint8
```

An Op is a single regular expression operator.

​	一个Op是一个单一的正则表达式运算符。

``` go 
const (
	OpNoMatch        Op = 1 + iota // 不匹配任何字符串 matches no strings
	OpEmptyMatch                   // 匹配空字符串 matches empty string
	OpLiteral                      // 匹配字符序列Runes matches Runes sequence
	OpCharClass                    // 匹配将Runes解释为范围对列表 matches Runes interpreted as range pair list
	OpAnyCharNotNL                 // 匹配除换行符外的任何字符 matches any character except newline
	OpAnyChar                      // 匹配任何字符 matches any character
	OpBeginLine                    // 在行首匹配空字符串 matches empty string at beginning of line
	OpEndLine                      // 在行尾匹配空字符串 matches empty string at end of line
	OpBeginText                    // 在文本开头匹配空字符串 matches empty string at beginning of text
	OpEndText                      // 在文本末尾匹配空字符串 matches empty string at end of text
	OpWordBoundary                 // 匹配单词边界 \b matches word boundary `\b`
	OpNoWordBoundary               // 匹配非单词边界 \B matches word non-boundary `\B`
	OpCapture                      // 捕获具有索引Cap和可选名称Name的子表达式 capturing subexpression with index Cap, optional name Name
	OpStar                         // 零次或多次匹配Sub [0] matches Sub[0] zero or more times
	OpPlus                         // 一次或多次匹配Sub [0] matches Sub[0] one or more times
	OpQuest                        // 零次或一次匹配Sub [0] matches Sub[0] zero or one times
	OpRepeat                       // 至少匹配Sub [0] Min次，最多匹配Max次(Max == -1没有限制) matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
	OpConcat                       // 匹配Subs的连接 matches concatenation of Subs
	OpAlternate                    // 匹配Subs的交替 matches alternation of Subs
)
```

#### (Op) String  <- go1.11

``` go 
func (i Op) String() string
```

### type Prog 

``` go 
type Prog struct {
	Inst   []Inst
	Start  int // 起始指令的索引 index of start instruction
	NumCap int // re中InstCapture指令的数量 number of InstCapture insts in re
}
```

A Prog is a compiled regular expression program.

​	一个Prog是一个已编译的正则表达式程序。

#### func Compile 

``` go 
func Compile(re *Regexp) (*Prog, error)
```

Compile compiles the regexp into a program to be executed. The regexp should have been simplified already (returned from re.Simplify).

​	`Compile`方法将正则表达式编译成可执行程序。 正则表达式应该已经被简化(从re.Simplify返回)。

#### (*Prog) Prefix 

``` go 
func (p *Prog) Prefix() (prefix string, complete bool)
```

Prefix returns a literal string that all matches for the regexp must start with. Complete is true if the prefix is the entire match.

​	`Prefix`方法返回一个字面字符串，该字符串是所有匹配该正则表达式的字符串必须以此开头。如果该字符串是整个匹配，则 complete 参数为 true。

#### (*Prog) StartCond 

``` go 
func (p *Prog) StartCond() EmptyOp
```

StartCond returns the leading empty-width conditions that must be true in any match. It returns ^EmptyOp(0) if no matches are possible.

​	`StartCond`方法返回必须在任何匹配中为真的前导空宽度条件。如果没有匹配，则返回 ^EmptyOp(0)。

#### (*Prog) String 

``` go 
func (p *Prog) String() string
```

​	`String`方法返回表示程序的字符串。

### type Regexp 

``` go 
type Regexp struct {
	Op       Op // 运算符 operator
	Flags    Flags
	Sub      []*Regexp  // 子表达式(如果有) subexpressions, if any
	Sub0     [1]*Regexp // 短子表达式的存储 storage for short Sub
	Rune     []rune     // 匹配的符文(对于 OpLiteral、OpCharClass) matched runes, for OpLiteral, OpCharClass
	Rune0    [2]rune    // 短符文的存储  storage for short Rune
	Min, Max int        // OpRepeat 的最小值和最大值 
	Cap      int        // OpCapture 的捕获索引 capturing index, for OpCapture
	Name     string     // OpCapture 的捕获名称 capturing name, for OpCapture
}
```

A Regexp is a node in a regular expression syntax tree.

​	Regexp结构体是正则表达式语法树中的一个节点。

#### func Parse 

``` go 
func Parse(s string, flags Flags) (*Regexp, error)
```

Parse parses a regular expression string s, controlled by the specified Flags, and returns a regular expression parse tree. The syntax is described in the top-level comment.

​	`Parse`方法解析一个受 flags 控制的正则表达式字符串 s，并返回一个正则表达式语法树。语法在顶层注释中描述。

#### (*Regexp) CapNames 

``` go 
func (re *Regexp) CapNames() []string
```

CapNames walks the regexp to find the names of capturing groups.

​	`CapNames`方法遍历正则表达式以查找捕获组的名称。

#### (*Regexp) Equal 

``` go 
func (x *Regexp) Equal(y *Regexp) bool
```

Equal reports whether x and y have identical structure.

​	`Equal`方法报告 x 和 y 是否具有相同的结构。

#### (*Regexp) MaxCap 

``` go 
func (re *Regexp) MaxCap() int
```

MaxCap walks the regexp to find the maximum capture index.

​	`MaxCap`方法遍历正则表达式以查找最大捕获索引。

#### (*Regexp) Simplify 

``` go 
func (re *Regexp) Simplify() *Regexp
```

Simplify returns a regexp equivalent to re but without counted repetitions and with various other simplifications, such as rewriting /(?:a+)+/ to /a+/. The resulting regexp will execute correctly but its string representation will not produce the same parse tree, because capturing parentheses may have been duplicated or removed. For example, the simplified form for /(x){1,2}/ is /(x)(x)?/ but both parentheses capture as $1. The returned regexp may share structure with or be the original.

​	`Simplify`方法返回等效于 re 的正则表达式，但不包括计数重复，并进行各种其他简化，例如将 `/(?:a+)+/` 重写为 `/a+/`。生成的正则表达式将正确执行，但其字符串表示形式不会生成相同的解析树，因为捕获括号可能已复制或删除。例如，`/(x){1,2}/` 的简化形式是`/(x)(x)?/`，但两个括号都捕获为 `$1`。返回的正则表达式可能与原始表达式共享结构，也可能是原始表达式的副本。

#### (*Regexp) String 

``` go 
func (re *Regexp) String() string
```

