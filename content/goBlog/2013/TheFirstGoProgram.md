+++
title = "第一个 Go 项目"
weight = 8
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# The first Go program - 第一个 Go 项目

> 原文：[https://go.dev/blog/first-go-program](https://go.dev/blog/first-go-program)

Andrew Gerrand
18 July 2013

Brad Fitzpatrick and I (Andrew Gerrand) recently started restructuring [godoc](https://go.dev/cmd/godoc/), and it occurred to me that it is one of the oldest Go programs. Robert Griesemer started writing it back in early 2009, and we’re still using it today.

Brad Fitzpatrick和我（Andrew Gerrand）最近开始重组godoc，我想到它是最古老的Go程序之一。Robert Griesemer早在2009年初就开始写它了，而我们至今还在使用它。

When I [tweeted](https://twitter.com/enneff/status/357403054632484865) about this, Dave Cheney replied with an [interesting question](https://twitter.com/davecheney/status/357406479415914497): what is the oldest Go program? Rob Pike dug into his mail and found it in an old message to Robert and Ken Thompson.

当我在推特上提到这个问题时，Dave Cheney回复了一个有趣的问题：最古老的Go程序是什么？罗伯-派克挖开他的邮件，在给罗伯特和肯-汤普森的一条旧信息中找到了它。

What follows is the first Go program. It was written by Rob in February 2008, when the team was just Rob, Robert, and Ken. They had a solid feature list (mentioned in [this blog post](https://commandcenter.blogspot.com.au/2012/06/less-is-exponentially-more.html)) and a rough language specfication. Ken had just finished the first working version of a Go compiler (it didn’t produce native code, but rather transliterated Go code to C for fast prototyping) and it was time to try writing a program with it.

下面的内容是第一个Go程序。它是由Rob在2008年2月写的，当时的团队只有Rob、Robert和Ken。他们有一个可靠的功能列表（在这篇博文中提到）和一个粗略的语言规范。Ken刚刚完成了Go编译器的第一个工作版本（它并不产生原生代码，而是将Go代码音译成C语言，以便快速制作原型），是时候尝试用它编写程序了。

Rob sent mail to the "Go team":

Rob给 "Go团队 "发了邮件:

```
From: Rob 'Commander' Pike
Date: Wed, Feb 6, 2008 at 3:42 PM
To: Ken Thompson, Robert Griesemer
Subject: slist

it works now.

roro=% a.out
(defn foo (add 12 34))
return: icounter = 4440
roro=%

here's the code.
some ugly hackery to get around the lack of strings.
```

(The `icounter` line in the program output is the number of executed statements, printed for debugging.)

(程序输出中的`icounter`行是已执行语句的数量，为调试而打印)。

```go
// +build OMIT

package main

// fake stuff
type char uint8;

// const char TESTSTRING[] = "(defn foo (add 'a 'b))\n";

type Atom struct {
        string  *[100]char;
        integer int;
        next    *Slist;  /* in hash bucket */
}

type List struct {
        car     *Slist;
        cdr     *Slist;
}

type Slist struct {
        isatom          bool;
        isstring        bool;
        //union {
        atom    Atom;
        list    List;
        //} u;

        Free method();
        Print method();
        PrintOne method(doparen bool);
        String method(*char <-);
        Integer method(int <-);
        Car method(*Slist <-);
        Cdr method(*Slist <-);
}

method (this *Slist) Car(*Slist <-) {
        return this.list.car;
}

method (this *Slist) Cdr(*Slist <-) {
        return this.list.cdr;
}

method (this *Slist) String(*[100]char <-) {
        return this.atom.string;
}

method (this *Slist) Integer(int <-) {
        return this.atom.integer;
}

function OpenFile();
function Parse(*Slist <-);

//Slist* atom(char *s, int i);

var token int;
var peekc int = -1;
var lineno int32 = 1;

var input [100*1000]char;
var inputindex int = 0;
var tokenbuf [100]char;

var EOF int = -1;  // BUG should be const

function main(int32 <-) {
        var list *Slist;

        OpenFile();
        for ;; {
                list = Parse();
                if list == nil {
                        break;
                }
                list.Print();
                list.Free();
                break;
        }

        return 0;
}

method (slist *Slist) Free(<-) {
        if slist == nil {
                return;
        }
        if slist.isatom {
//              free(slist.String());
        } else {
                slist.Car().Free();
                slist.Cdr().Free();
        }
//      free(slist);
}

method (slist *Slist) PrintOne(<- doparen bool) {
        if slist == nil {
                return;
        }
        if slist.isatom {
                if slist.isstring {
                        print(slist.String());
                } else {
                        print(slist.Integer());
                }
        } else {
                if doparen {
                        print("(");
                }
                slist.Car().PrintOne(true);
                if slist.Cdr() != nil {
                        print(" ");
                        slist.Cdr().PrintOne(false);
                }
                if doparen {
                        print(")");
                }
        }
}

method (slist *Slist) Print() {
        slist.PrintOne(true);
        print "\n";
}

function Get(int <-) {
        var c int;

        if peekc >= 0 {
                c = peekc;
                peekc = -1;
        } else {
                c = convert(int, input[inputindex]);
                inputindex = inputindex + 1; // BUG should be incr one expr
                if c == '\n' {
                        lineno = lineno + 1;
                }
                if c == '\0' {
                        inputindex = inputindex - 1;
                        c = EOF;
                }
        }
        return c;
}

function WhiteSpace(bool <- c int) {
        return c == ' ' || c == '\t' || c == '\r' || c == '\n';
}

function NextToken() {
        var i, c int;
        var backslash bool;

        tokenbuf[0] = '\0';     // clear previous token
        c = Get();
        while WhiteSpace(c)  {
                c = Get();
        }
        switch c {
                case EOF:
                        token = EOF;
                case '(':
                case ')':
                        token = c;
                        break;
                case:
                        for i = 0; i < 100 - 1; {  // sizeof tokenbuf - 1
                                tokenbuf[i] = convert(char, c);
                                i = i + 1;
                                c = Get();
                                if c == EOF {
                                        break;
                                }
                                if WhiteSpace(c) || c == ')' {
                                        peekc = c;
                                        break;
                                }
                        }
                        if i >= 100 - 1 {  // sizeof tokenbuf - 1
                                panic "atom too long\n";
                        }
                        tokenbuf[i] = '\0';
                        if '0' <= tokenbuf[0] && tokenbuf[0] <= '9' {
                                token = '0';
                        } else {
                                token = 'A';
                        }
        }
}

function Expect(<- c int) {
        if token != c {
                print "parse error: expected ", c, "\n";
                panic "parse";
        }
        NextToken();
}

// Parse a non-parenthesized list up to a closing paren or EOF
function ParseList(*Slist <-) {
        var slist, retval *Slist;

        slist = new(Slist);
        slist.list.car = nil;
        slist.list.cdr = nil;
        slist.isatom = false;
        slist.isstring = false;

        retval = slist;
        for ;; {
                slist.list.car = Parse();
                if token == ')' {       // empty cdr
                        break;
                }
                if token == EOF {       // empty cdr  BUG SHOULD USE ||
                        break;
                }
                slist.list.cdr = new(Slist);
                slist = slist.list.cdr;
        }
        return retval;
}

function atom(*Slist <- i int) {  // BUG: uses tokenbuf; should take argument
        var h, length int;
        var slist, tail *Slist;

        slist = new(Slist);
        if token == '0' {
                slist.atom.integer = i;
                slist.isstring = false;
        } else {
                slist.atom.string = new([100]char);
                var i int;
                for i = 0; ; i = i + 1 {
                        (*slist.atom.string)[i] = tokenbuf[i];
                        if tokenbuf[i] == '\0' {
                                break;
                        }
                }
                //slist.atom.string = "hello"; // BUG! s; //= strdup(s);
                slist.isstring = true;
        }
        slist.isatom = true;
        return slist;
}

function atoi(int <-) {  // BUG: uses tokenbuf; should take argument
        var v int = 0;
        for i := 0; '0' <= tokenbuf[i] && tokenbuf[i] <= '9'; i = i + 1 {
                v = 10 * v + convert(int, tokenbuf[i] - '0');
        }
        return v;
}

function Parse(*Slist <-) {
        var slist *Slist;

        if token == EOF || token == ')' {
                return nil;
        }
        if token == '(' {
                NextToken();
                slist = ParseList();
                Expect(')');
                return slist;
        } else {
                // Atom
                switch token {
                        case EOF:
                                return nil;
                        case '0':
                                slist = atom(atoi());
                        case '"':
                        case 'A':
                                slist = atom(0);
                        case:
                                slist = nil;
                                print "unknown token"; //, token, tokenbuf;
                }
                NextToken();
                return slist;
        }
        return nil;
}

function OpenFile() {
        //strcpy(input, TESTSTRING);
        //inputindex = 0;
        // (defn foo (add 12 34))\n
        inputindex = 0;
        peekc = -1;  // BUG
        EOF = -1;  // BUG
        i := 0;
        input[i] = '('; i = i + 1;
        input[i] = 'd'; i = i + 1;
        input[i] = 'e'; i = i + 1;
        input[i] = 'f'; i = i + 1;
        input[i] = 'n'; i = i + 1;
        input[i] = ' '; i = i + 1;
        input[i] = 'f'; i = i + 1;
        input[i] = 'o'; i = i + 1;
        input[i] = 'o'; i = i + 1;
        input[i] = ' '; i = i + 1;
        input[i] = '('; i = i + 1;
        input[i] = 'a'; i = i + 1;
        input[i] = 'd'; i = i + 1;
        input[i] = 'd'; i = i + 1;
        input[i] = ' '; i = i + 1;
        input[i] = '1'; i = i + 1;
        input[i] = '2'; i = i + 1;
        input[i] = ' '; i = i + 1;
        input[i] = '3'; i = i + 1;
        input[i] = '4'; i = i + 1;
        input[i] = ')'; i = i + 1;
        input[i] = ')'; i = i + 1;
        input[i] = '\n'; i = i + 1;
        NextToken();
}
```

The program parses and prints an [S-expression](https://en.wikipedia.org/wiki/S-expression). It takes no user input and has no imports, relying only on the built-in `print` facility for output. It was written literally the first day there was a [working but rudimentary compiler](https://go.dev/change/8b8615138da3). Much of the language wasn’t implemented and some of it wasn’t even specified.

这个程序解析并打印了一个S表达式。它不需要用户输入，也没有导入，只依靠内置的打印工具进行输出。它是在有一个可以工作但不成熟的编译器的第一天写的。该语言的大部分内容没有实现，有些甚至没有被指定。

Still, the basic flavor of the language today is recognizable in this program. Type and variable declarations, control flow, and package statements haven’t changed much.

不过，今天这种语言的基本味道在这个程序中还是可以识别的。类型和变量声明、控制流和包语句都没有什么变化。

But there are many differences and absences. Most significant are the lack of concurrency and interfaces—both considered essential since day 1 but not yet designed.

但也有许多不同和缺失。最重要的是缺乏并发性和接口--这两者从第一天起就被认为是必不可少的，但还没有设计出来。

A `func` was a `function`, and its signature specified return values *before* arguments, separating them with `<-`, which we now use as the channel send/receive operator. For example, the `WhiteSpace` function takes the integer `c` and returns a boolean.

func是一个函数，它的签名在参数之前指定了返回值，用<-分隔它们，我们现在用它作为通道发送/接收操作符。例如，WhiteSpace函数接收整数c并返回一个布尔值。

```go
function WhiteSpace(bool <- c int)
```

This arrow was a stop-gap measure until a better syntax arose for declaring multiple return values.

这个箭头是一个权宜之计，直到出现一个更好的语法来声明多个返回值。

Methods were distinct from functions and had their own keyword.

方法与函数不同，有自己的关键字。

```go
method (this *Slist) Car(*Slist <-) {
    return this.list.car;
}
```

And methods were pre-declared in the struct definition, although that changed soon.

方法是在结构定义中预先声明的，尽管这很快就改变了。

```go
type Slist struct {
    ...
    Car method(*Slist <-);
}
```

There were no strings, although they were in the spec. To work around this, Rob had to build the input string as an `uint8` array with a clumsy construction. (Arrays were rudimentary and slices hadn’t been designed yet, let alone implemented, although there was the unimplemented concept of an "open array".)

没有字符串，尽管在规范里有。为了解决这个问题，Rob不得不将输入的字符串构建为一个结构笨拙的uint8数组。(当时的数组还很初级，切片还没有被设计出来，更不用说实现了，虽然有一个未实现的 "开放数组 "的概念。)

```
input[i] = '('; i = i + 1;
input[i] = 'd'; i = i + 1;
input[i] = 'e'; i = i + 1;
input[i] = 'f'; i = i + 1;
input[i] = 'n'; i = i + 1;
input[i] = ' '; i = i + 1;
...
```

Both `panic` and `print` were built-in keywords, not pre-declared functions.

panic和print都是内置的关键字，不是预先声明的函数。

```
print "parse error: expected ", c, "\n";
panic "parse";
```

And there are many other little differences; see if you can identify some others.

还有很多其他的小区别；看看您是否能找出一些其他的区别。

Less than two years after this program was written, Go was released as an open source project. Looking back, it is striking how much the language has grown and matured. (The last thing to change between this proto-Go and the Go we know today was the elimination of semicolons.)

在写完这个程序不到两年的时间里，Go作为一个开源项目被发布。回顾过去，令人惊讶的是这门语言已经成长和成熟了。(在这个原版Go和我们今天所知道的Go之间，最后的变化是取消了分号)。

But even more striking is how much we have learned about *writing* Go code. For instance, Rob called his method receivers `this`, but now we use shorter context-specific names. There are hundreds of more significant examples and to this day we’re still discovering better ways to write Go code. (Check out the [glog package](https://github.com/golang/glog)’s clever trick for [handling verbosity levels](https://github.com/golang/glog/blob/c6f9652c7179652e2fd8ed7002330db089f4c9db/glog.go#L893).)

但更令人震惊的是我们在编写Go代码方面学到了很多东西。例如，Rob把他的方法称为接收者，但现在我们使用更短的特定上下文名称。还有数百个更重要的例子，直到今天我们还在发现更好的方法来编写Go代码。(看看glog包处理verbosity级别的巧妙技巧）。

I wonder what we’ll learn tomorrow.

我想知道我们明天会学到什么。
