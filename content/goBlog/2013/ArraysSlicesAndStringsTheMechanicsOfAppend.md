+++
title = "数组、切片（和字符串）：append 的原理"
weight = 7
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Arrays, slices (and strings): The mechanics of 'append'  - 数组、切片（和字符串）：append 的原理

https://go.dev/blog/slices

Rob Pike
26 September 2013

## Introduction 简介

One of the most common features of procedural programming languages is the concept of an array. Arrays seem like simple things but there are many questions that must be answered when adding them to a language, such as:

程序性编程语言最常见的特征之一是数组的概念。数组看起来是很简单的东西，但是当把它们加入到语言中时，有很多问题必须要回答，比如说：

- fixed-size or variable-size? 固定大小还是可变大小？
- is the size part of the type? 大小是类型的一部分吗？
- what do multidimensional arrays look like? 多维数组是什么样子的？
- does the empty array have meaning? 空数组是否有意义？

The answers to these questions affect whether arrays are just a feature of the language or a core part of its design.

这些问题的答案会影响到数组是否只是语言的一个特征，还是其设计的核心部分。

In the early development of Go, it took about a year to decide the answers to these questions before the design felt right. The key step was the introduction of *slices*, which built on fixed-size *arrays* to give a flexible, extensible data structure. To this day, however, programmers new to Go often stumble over the way slices work, perhaps because experience from other languages has colored their thinking.

在Go的早期开发中，我们花了大约一年的时间来决定这些问题的答案，然后才觉得设计是正确的。关键的一步是引入了分片，在固定大小的数组基础上，提供了一个灵活的、可扩展的数据结构。然而，时至今日，刚接触Go的程序员往往对分片的工作方式感到困惑，也许是因为其他语言的经验影响了他们的思维。

In this post we’ll attempt to clear up the confusion. We’ll do so by building up the pieces to explain how the `append` built-in function works, and why it works the way it does.

在这篇文章中，我们将试图消除这种困惑。我们将通过建立碎片来解释append内置函数是如何工作的，以及为什么它能以这样的方式工作。

## Arrays 数组

Arrays are an important building block in Go, but like the foundation of a building they are often hidden below more visible components. We must talk about them briefly before we move on to the more interesting, powerful, and prominent idea of slices.

数组是Go中的一个重要构件，但就像建筑物的地基一样，它们往往被隐藏在更多可见的组件之下。在我们继续讨论更有趣、更强大、更突出的切片概念之前，我们必须简单地谈谈它们。

Arrays are not often seen in Go programs because the size of an array is part of its type, which limits its expressive power.

数组在Go程序中并不常见，因为数组的大小是其类型的一部分，这限制了其表达能力。

The declaration

声明

```go linenums="1"
var buffer [256]byte
```

declares the variable `buffer`, which holds 256 bytes. The type of `buffer` includes its size, `[256]byte`. An array with 512 bytes would be of the distinct type `[512]byte`.

声明了变量buffer，它可以容纳256字节。buffer的类型包括其大小，[256]字节。一个有512个字节的数组将是独特的类型[512]byte。

The data associated with an array is just that: an array of elements. Schematically, our buffer looks like this in memory,

与数组相关的数据只是：一个元素的数组。从原理上讲，我们的缓冲区在内存中看起来是这样的，

```
buffer: byte byte byte ... 256 times ... byte byte byte
```

That is, the variable holds 256 bytes of data and nothing else. We can access its elements with the familiar indexing syntax, `buffer[0]`, `buffer[1]`, and so on through `buffer[255]`. (The index range 0 through 255 covers 256 elements.) Attempting to index `buffer` with a value outside this range will crash the program.

也就是说，这个变量持有256个字节的数据，没有其他东西。我们可以用熟悉的索引语法访问它的元素，buffer[0], buffer[1], 以此类推直到buffer[255]。(索引范围0到255涵盖了256个元素。)试图用超出这个范围的值来索引缓冲区将使程序崩溃。

There is a built-in function called `len` that returns the number of elements of an array or slice and also of a few other data types. For arrays, it’s obvious what `len` returns. In our example, `len(buffer)` returns the fixed value 256.

有一个内置的函数叫len，它可以返回数组或片断以及其他一些数据类型的元素数。对于数组来说，len的返回值是很明显的。在我们的例子中，len(buffer)返回固定值256。

Arrays have their place—they are a good representation of a transformation matrix for instance—but their most common purpose in Go is to hold storage for a slice.

数组有它的用武之地--例如，它们是变换矩阵的良好代表--但是它们在 Go 中最常见的用途是为一个片断保存存储空间。

## Slices: The slice header 切片：切片头

Slices are where the action is, but to use them well one must understand exactly what they are and what they do.

分片是行动的地方，但要很好地使用它们，就必须准确地了解它们是什么以及它们做什么。

A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself. *A slice is not an array*. A slice *describes* a piece of an array.

分片是一个数据结构，描述了一个数组的连续部分，与分片变量本身分开存储。分片不是一个数组。分片描述的是一个数组的一个部分。

Given our `buffer` array variable from the previous section, we could create a slice that describes elements 100 through 150 (to be precise, 100 through 149, inclusive) by *slicing* the array:

考虑到上一节中我们的缓冲区数组变量，我们可以通过对数组的切片来创建一个描述元素100到150（准确的说是100到149，包括在内）的切片。

```go linenums="1"
var slice []byte = buffer[100:150]
```

In that snippet we used the full variable declaration to be explicit. The variable `slice` has type `[]byte`, pronounced "slice of bytes", and is initialized from the array, called `buffer`, by slicing elements 100 (inclusive) through 150 (exclusive). The more idiomatic syntax would drop the type, which is set by the initializing expression:

在该片段中，我们使用了完整的变量声明，以示明确。变量slice的类型是[]byte，读作 "字节的片断"，并从数组中初始化，称为buffer，通过对元素100（包括）到150（不包括）的片断。更恰当的语法是去掉类型，它是由初始化表达式设置的：

```go linenums="1"
var slice = buffer[100:150]
```

Inside a function we could use the short declaration form,

在一个函数中，我们可以使用简短的声明形式。

```go linenums="1"
slice := buffer[100:150]
```

What exactly is this slice variable? It’s not quite the full story, but for now think of a slice as a little data structure with two elements: a length and a pointer to an element of an array. You can think of it as being built like this behind the scenes:

这个slice变量到底是什么？这还不是很全面，但现在可以把slice想象成一个有两个元素的小数据结构：一个长度和一个指向数组元素的指针。您可以认为它在幕后是这样构建的：

```go linenums="1"
type sliceHeader struct {
    Length        int
    ZerothElement *byte
}

slice := sliceHeader{
    Length:        50,
    ZerothElement: &buffer[100],
}
```

Of course, this is just an illustration. Despite what this snippet says that `sliceHeader` struct is not visible to the programmer, and the type of the element pointer depends on the type of the elements, but this gives the general idea of the mechanics.

当然，这只是一个说明。尽管这个片段说sliceHeader结构对程序员来说是不可见的，而且元素指针的类型取决于元素的类型，但这给出了机械原理的大致概念。

So far we’ve used a slice operation on an array, but we can also slice a slice, like this:

到目前为止，我们已经在数组上使用了切片操作，但是我们也可以对一个切片进行切片操作，像这样：

```go linenums="1"
slice2 := slice[5:10]
```

Just as before, this operation creates a new slice, in this case with elements 5 through 9 (inclusive) of the original slice, which means elements 105 through 109 of the original array. The underlying `sliceHeader` struct for the `slice2` variable looks like this:

就像以前一样，这个操作创建了一个新的切片，在这种情况下，原始切片的元素5到9（包括），也就是原始数组的元素105到109。slice2变量的底层sliceHeader结构看起来像这样：

```go linenums="1"
slice2 := sliceHeader{
    Length:        5,
    ZerothElement: &buffer[105],
}
```

Notice that this header still points to the same underlying array, stored in the `buffer` variable.

注意，这个头仍然指向相同的底层数组，存储在缓冲区变量中。

We can also *reslice*, which is to say slice a slice and store the result back in the original slice structure. After

我们也可以重新切分，也就是切分一个片断，并将结果存储回原来的片断结构中。之后

```
slice = slice[5:10]
```

the `sliceHeader` structure for the `slice` variable looks just like it did for the `slice2` variable. You’ll see reslicing used often, for example to truncate a slice. This statement drops the first and last elements of our slice:

slice变量的sliceHeader结构看起来就像slice2变量的结构一样。您会看到经常使用重新切分，例如截断一个片断。这条语句删除了我们片断的第一个和最后一个元素：

```go linenums="1"
slice = slice[1:len(slice)-1]
```

[Exercise: Write out what the `sliceHeader` struct looks like after this assignment.]

[练习。写出这个赋值后sliceHeader结构的样子] 。

You’ll often hear experienced Go programmers talk about the "slice header" because that really is what’s stored in a slice variable. For instance, when you call a function that takes a slice as an argument, such as [bytes.IndexRune](https://go.dev/pkg/bytes/#IndexRune), that header is what gets passed to the function. In this call,

您会经常听到有经验的Go程序员谈论 "slice header"，因为这确实是存储在slice变量中的东西。例如，当您调用一个以分片为参数的函数时，比如byte.IndexRune，这个头就是传递给该函数的东西。在这个调用中，

```go linenums="1"
slashPos := bytes.IndexRune(slice, '/')
```

the `slice` argument that is passed to the `IndexRune` function is, in fact, a "slice header".

被传递给IndexRune函数的slice参数实际上是一个 "slice头"。

There’s one more data item in the slice header, which we talk about below, but first let’s see what the existence of the slice header means when you program with slices.

切片头中还有一个数据项，我们将在下面讨论，但首先让我们看看当您用切片编程时，切片头的存在意味着什么。

## Passing slices to functions 向函数传递片断

It’s important to understand that even though a slice contains a pointer, it is itself a value. Under the covers, it is a struct value holding a pointer and a length. It is *not* a pointer to a struct.

重要的是要理解，即使切片包含一个指针，它本身也是一个值。从表面上看，它是一个持有一个指针和一个长度的结构值。它不是一个指向结构的指针。

This matters.

这一点很重要。

When we called `IndexRune` in the previous example, it was passed a *copy* of the slice header. That behavior has important ramifications.

当我们在前面的例子中调用 IndexRune 时，它被传递了一份分片头的副本。这种行为有重要的影响。

Consider this simple function:

考虑一下这个简单的函数：

```go linenums="1"
func AddOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}
```

It does just what its name implies, iterating over the indices of a slice (using a `for` `range` loop), incrementing its elements.

它就像它的名字所暗示的那样，迭代一个片断的索引（使用for range循环），增加它的元素。

Try it:

试试吧：

```go linenums="1"
func main() {
    slice := buffer[10:20]
    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i)
    }
    fmt.Println("before", slice)
    AddOneToEachElement(slice)
    fmt.Println("after", slice)
}
```

Run

运行

(You can edit and re-execute these runnable snippets if you want to explore.)

(如果您想探索的话，您可以编辑并重新执行这些可运行的片段)。

Even though the slice *header* is passed by value, the header includes a pointer to elements of an array, so both the original slice header and the copy of the header passed to the function describe the same array. Therefore, when the function returns, the modified elements can be seen through the original slice variable.

尽管slice头是通过值传递的，但头包括一个指向数组元素的指针，所以原始slice头和传递给函数的头的副本都描述了同一个数组。因此，当函数返回时，可以通过原来的片头变量看到修改后的元素。

The argument to the function really is a copy, as this example shows:

如本例所示，该函数的参数实际上是一个副本：

```go linenums="1"
func SubtractOneFromLength(slice []byte) []byte {
    slice = slice[0 : len(slice)-1]
    return slice
}

func main() {
    fmt.Println("Before: len(slice) =", len(slice))
    newSlice := SubtractOneFromLength(slice)
    fmt.Println("After:  len(slice) =", len(slice))
    fmt.Println("After:  len(newSlice) =", len(newSlice))
}
```

Run

运行

Here we see that the *contents* of a slice argument can be modified by a function, but its *header* cannot. The length stored in the `slice` variable is not modified by the call to the function, since the function is passed a copy of the slice header, not the original. Thus if we want to write a function that modifies the header, we must return it as a result parameter, just as we have done here. The `slice` variable is unchanged but the returned value has the new length, which is then stored in `newSlice`,

在这里我们看到，slice参数的内容可以被一个函数修改，但它的头不能被修改。存储在slice变量中的长度不会因为调用函数而被修改，因为函数传递给它的是slice头的副本，而不是原件。因此，如果我们想写一个修改头的函数，我们必须把它作为一个结果参数返回，就像我们在这里做的那样。切片变量没有变化，但返回值有新的长度，然后存储在newSlice中。

## Pointers to slices: Method receivers 切片的指针：方法接收者

Another way to have a function modify the slice header is to pass a pointer to it. Here’s a variant of our previous example that does this:

另一种让函数修改切片头的方法是传递一个指向它的指针。下面是我们以前的例子的一个变体，它是这样做的：

```go linenums="1"
func PtrSubtractOneFromLength(slicePtr *[]byte) {
    slice := *slicePtr
    *slicePtr = slice[0 : len(slice)-1]
}

func main() {
    fmt.Println("Before: len(slice) =", len(slice))
    PtrSubtractOneFromLength(&slice)
    fmt.Println("After:  len(slice) =", len(slice))
}
```

Run

运行

It seems clumsy in that example, especially dealing with the extra level of indirection (a temporary variable helps), but there is one common case where you see pointers to slices. It is idiomatic to use a pointer receiver for a method that modifies a slice.

在这个例子中，它看起来很笨拙，特别是处理额外的指示层次（一个临时变量有助于），但有一个常见的情况，您看到指向切片的指针。在修改片断的方法中使用指针接收器是一种习惯。

Let’s say we wanted to have a method on a slice that truncates it at the final slash. We could write it like this:

假设我们想在一个切片上有一个方法，在最后的斜线处将其截断。我们可以这样写：

```go linenums="1"
type path []byte

func (p *path) TruncateAtFinalSlash() {
    i := bytes.LastIndex(*p, []byte("/"))
    if i >= 0 {
        *p = (*p)[0:i]
    }
}

func main() {
    pathName := path("/usr/bin/tso") // Conversion from string to path.
    pathName.TruncateAtFinalSlash()
    fmt.Printf("%s\n", pathName)
}
```

Run 运行

If you run this example you’ll see that it works properly, updating the slice in the caller.

如果您运行这个例子，您会看到它工作正常，更新了调用者中的片断。

[Exercise: Change the type of the receiver to be a value rather than a pointer and run it again. Explain what happens.]

[练习。把接收器的类型改成一个值而不是一个指针，然后再次运行。解释一下发生了什么]。

On the other hand, if we wanted to write a method for `path` that upper-cases the ASCII letters in the path (parochially ignoring non-English names), the method could be a value because the value receiver will still point to the same underlying array.

另一方面，如果我们想为path写一个方法，将路径中的ASCII字母大写（狭义上忽略了非英文名字），这个方法可以是一个值，因为值接收器仍然会指向同一个底层数组。

```go linenums="1"
type path []byte

func (p path) ToUpper() {
    for i, b := range p {
        if 'a' <= b && b <= 'z' {
            p[i] = b + 'A' - 'a'
        }
    }
}

func main() {
    pathName := path("/usr/bin/tso")
    pathName.ToUpper()
    fmt.Printf("%s\n", pathName)
}
```

Run

运行

Here the `ToUpper` method uses two variables in the `for` `range` construct to capture the index and slice element. This form of loop avoids writing `p[i]` multiple times in the body.

这里ToUpper方法在for range结构中使用了两个变量来捕获索引和片断元素。这种形式的循环避免了在正文中多次写入p[i]。

[Exercise: Convert the `ToUpper` method to use a pointer receiver and see if its behavior changes.]

[练习。将ToUpper方法转换为使用一个指针接收器，看看它的行为是否有变化。］

[Advanced exercise: Convert the `ToUpper` method to handle Unicode letters, not just ASCII.]

[高级练习。转换ToUpper方法来处理Unicode字母，而不仅仅是ASCII。］

## Capacity 容量

Look at the following function that extends its argument slice of `ints` by one element:

看看下面这个函数，它将其参数slice的ints扩展了一个元素：

```go linenums="1"
func Extend(slice []int, element int) []int {
    n := len(slice)
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}
```

(Why does it need to return the modified slice?) Now run it:

(为什么需要返回修改后的slice？) 现在运行它。

```go linenums="1"
func main() {
    var iBuffer [10]int
    slice := iBuffer[0:0]
    for i := 0; i < 20; i++ {
        slice = Extend(slice, i)
        fmt.Println(slice)
    }
}
```

Run

运行

See how the slice grows until… it doesn’t.

看看slice是如何增长的，直到......它不增长了。

It’s time to talk about the third component of the slice header: its *capacity*. Besides the array pointer and length, the slice header also stores its capacity:

现在是时候谈谈分片头的第三个组成部分了：它的容量。除了数组指针和长度，切片头还存储了它的容量：

```go linenums="1"
type sliceHeader struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}
```

The `Capacity` field records how much space the underlying array actually has; it is the maximum value the `Length` can reach. Trying to grow the slice beyond its capacity will step beyond the limits of the array and will trigger a panic.

Capacity字段记录了底层数组实际有多少空间；它是Length可以达到的最大值。试图使分片的增长超过它的容量，就会超出数组的限制，并会引发恐慌。

After our example slice is created by

在我们的例子中，分片是通过以下方式创建的

```go linenums="1"
slice := iBuffer[0:0]
```

its header looks like this:

它的头看起来像这样：

```go linenums="1"
slice := sliceHeader{
    Length:        0,
    Capacity:      10,
    ZerothElement: &iBuffer[0],
}
```

The `Capacity` field is equal to the length of the underlying array, minus the index in the array of the first element of the slice (zero in this case). If you want to inquire what the capacity is for a slice, use the built-in function `cap`:

容量字段等于底层数组的长度，减去分片第一个元素在数组中的索引（本例中为0）。如果您想查询一个片断的容量是多少，可以使用内置函数cap：

```go linenums="1"
if cap(slice) == len(slice) {
    fmt.Println("slice is full!")
}
```

## Make

What if we want to grow the slice beyond its capacity? You can’t! By definition, the capacity is the limit to growth. But you can achieve an equivalent result by allocating a new array, copying the data over, and modifying the slice to describe the new array.

如果我们想让片断增长到超过其容量怎么办？您不能这样做! 根据定义，容量是增长的极限。但是您可以通过分配一个新的数组，把数据复制过来，然后修改slice来描述新的数组来达到同等的效果。

Let’s start with allocation. We could use the `new` built-in function to allocate a bigger array and then slice the result, but it is simpler to use the `make` built-in function instead. It allocates a new array and creates a slice header to describe it, all at once. The `make` function takes three arguments: the type of the slice, its initial length, and its capacity, which is the length of the array that `make` allocates to hold the slice data. This call creates a slice of length 10 with room for 5 more (15-10), as you can see by running it:

让我们从分配开始。我们可以使用新的内置函数来分配一个更大的数组，然后对结果进行切片，但是使用内置函数make更简单。它分配了一个新的数组，并创建了一个切片头来描述它，一次完成。make函数需要三个参数：分片的类型，它的初始长度，以及它的容量，也就是make分配的用来存放分片数据的数组的长度。这个调用创建了一个长度为10的分片，并留有容纳5个分片的空间（15-10），您可以通过运行它看到这一点：

```go linenums="1"
    slice := make([]int, 10, 15)
    fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
```

Run 运行

This snippet doubles the capacity of our `int` slice but keeps its length the same:

运行这段代码后，在需要再次重新分配之前，片子有了更大的增长空间：

```go linenums="1"
    slice := make([]int, 10, 15)
    fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
    newSlice := make([]int, len(slice), 2*cap(slice))
    for i := range slice {
        newSlice[i] = slice[i]
    }
    slice = newSlice
    fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
```

Run 运行

After running this code the slice has much more room to grow before needing another reallocation.

运行这段代码后，在需要再次重新分配之前，片子有了更大的增长空间。

When creating slices, it’s often true that the length and capacity will be same. The `make` built-in has a shorthand for this common case. The length argument defaults to the capacity, so you can leave it out to set them both to the same value. After

在创建分片时，通常情况下，长度和容量是一致的。内置的make软件对这种常见的情况有一个速记法。长度参数的默认值是容量，所以您可以省略它，把它们都设置为相同的值。之后

```go linenums="1"
gophers := make([]Gopher, 10)
```

the `gophers` slice has both its length and capacity set to 10.

## Copy 复制

When we doubled the capacity of our slice in the previous section, we wrote a loop to copy the old data to the new slice. Go has a built-in function, `copy`, to make this easier. Its arguments are two slices, and it copies the data from the right-hand argument to the left-hand argument. Here’s our example rewritten to use `copy`:

当我们在上一节中把分片的容量增加一倍时，我们写了一个循环，把旧的数据复制到新的分片中。Go有一个内置的函数，copy，使之更容易。它的参数是两个片断，它把数据从右边的参数复制到左边的参数。下面是我们使用copy重写的例子：

```go linenums="1"
    newSlice := make([]int, len(slice), 2*cap(slice))
    copy(newSlice, slice)
```

Run 运行

The `copy` function is smart. It only copies what it can, paying attention to the lengths of both arguments. In other words, the number of elements it copies is the minimum of the lengths of the two slices. This can save a little bookkeeping. Also, `copy` returns an integer value, the number of elements it copied, although it’s not always worth checking.

复制函数是很聪明的。它只复制它能复制的东西，注意两个参数的长度。换句话说，它复制的元素数量是两个切片长度的最小值。这可以节省一点簿记工作。另外，copy返回一个整数值，即它所复制的元素的数量，尽管这并不总是值得检查。

The `copy` function also gets things right when source and destination overlap, which means it can be used to shift items around in a single slice. Here’s how to use `copy` to insert a value into the middle of a slice.

当源片和目的片重叠时，copy函数也能做出正确的判断，这意味着它可以被用来在一个片中移动项目。下面是如何使用copy将一个值插入到一个片断的中间。

```go linenums="1"
// Insert inserts the value into the slice at the specified index,
// which must be in range.
// The slice must have room for the new element.
func Insert(slice []int, index, value int) []int {
    // Grow the slice by one element.
    slice = slice[0 : len(slice)+1]
    // Use copy to move the upper part of the slice out of the way and open a hole.
    copy(slice[index+1:], slice[index:])
    // Store the new value.
    slice[index] = value
    // Return the result.
    return slice
}
```

There are a couple of things to notice in this function. First, of course, it must return the updated slice because its length has changed. Second, it uses a convenient shorthand. The expression

在这个函数中，有几件事情需要注意。首先，当然，它必须返回更新的slice，因为它的长度已经改变。第二，它使用了一种方便的速记方法。表达式是

```go linenums="1"
slice[i:]
```

means exactly the same as

的意思与

```go linenums="1"
slice[i:len(slice)]
```

Also, although we haven’t used the trick yet, we can leave out the first element of a slice expression too; it defaults to zero. Thus

另外，虽然我们还没有使用这个技巧，但我们也可以省去slice表达式的第一个元素；它默认为零。因此

```go linenums="1"
slice[:]
```

just means the slice itself, which is useful when slicing an array. This expression is the shortest way to say "a slice describing all the elements of the array":

就是指分片本身，这在对数组进行分片时很有用。这个表达式是 "描述数组中所有元素的分片 "的最简短说法：

```go linenums="1"
array[:]
```

Now that’s out of the way, let’s run our `Insert` function.

现在，我们来运行我们的插入函数。

```go linenums="1"
    slice := make([]int, 10, 20) // Note capacity > length: room to add element.
    for i := range slice {
        slice[i] = i
    }
    fmt.Println(slice)
    slice = Insert(slice, 5, 99)
    fmt.Println(slice)
```

Run 运行

## Append: An example 追加：一个例子

A few sections back, we wrote an `Extend` function that extends a slice by one element. It was buggy, though, because if the slice’s capacity was too small, the function would crash. (Our `Insert` example has the same problem.) Now we have the pieces in place to fix that, so let’s write a robust implementation of `Extend` for integer slices.

在前几节，我们写了一个Extend函数，可以将一个slice扩展一个元素。但这是个错误，因为如果片断的容量太小，函数就会崩溃。(我们的Insert例子也有同样的问题。)现在我们已经有了解决这个问题的方法，所以让我们为整数片写一个强大的Extend实现。

```go linenums="1"
func Extend(slice []int, element int) []int {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]int, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}
```

In this case it’s especially important to return the slice, since when it reallocates the resulting slice describes a completely different array. Here’s a little snippet to demonstrate what happens as the slice fills up:

在这种情况下，返回slice是特别重要的，因为当它重新分配时，产生的slice描述了一个完全不同的数组。这里有一个小片段来演示当slice填满时发生了什么：

```go linenums="1"
    slice := make([]int, 0, 5)
    for i := 0; i < 10; i++ {
        slice = Extend(slice, i)
        fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
        fmt.Println("address of 0th element:", &slice[0])
    }
```

Run 运行

Notice the reallocation when the initial array of size 5 is filled up. Both the capacity and the address of the zeroth element change when the new array is allocated.

注意当初始的5号数组被填满时的重新分配。当新的数组被分配时，容量和第2个元素的地址都会改变。

With the robust `Extend` function as a guide we can write an even nicer function that lets us extend the slice by multiple elements. To do this, we use Go’s ability to turn a list of function arguments into a slice when the function is called. That is, we use Go’s variadic function facility.

有了强大的Extend函数作为指导，我们可以写一个更漂亮的函数，让我们通过多个元素来扩展分片。要做到这一点，我们利用Go的能力，在函数被调用时将函数参数列表变成一个片断。也就是说，我们使用Go的变量函数工具。

Let’s call the function `Append`. For the first version, we can just call `Extend` repeatedly so the mechanism of the variadic function is clear. The signature of `Append` is this:

让我们把这个函数称为Append。对于第一个版本，我们可以直接重复调用Extend，这样变量函数的机制就清楚了。Append的签名是这样的：

```go linenums="1"
func Append(slice []int, items ...int) []int
```

What that says is that `Append` takes one argument, a slice, followed by zero or more `int` arguments. Those arguments are exactly a slice of `int` as far as the implementation of `Append` is concerned, as you can see:

这就是说，Append需要一个参数，一个片断，然后是0个或更多的int参数。就Append的实现而言，这些参数正好是一个int的切片，您可以看到：

```go linenums="1"
// Append appends the items to the slice.
// First version: just loop calling Extend.
func Append(slice []int, items ...int) []int {
    for _, item := range items {
        slice = Extend(slice, item)
    }
    return slice
}
```

Notice the `for` `range` loop iterating over the elements of the `items` argument, which has implied type `[]int`. Also notice the use of the blank identifier `_` to discard the index in the loop, which we don’t need in this case.

注意for range循环遍历了items参数的元素，它的隐含类型是[]int。还注意到在循环中使用了空白标识符_来丢弃索引，在这种情况下我们不需要。

Try it:

试试吧：

```go linenums="1"
    slice := []int{0, 1, 2, 3, 4}
    fmt.Println(slice)
    slice = Append(slice, 5, 6, 7, 8)
    fmt.Println(slice)
```

Run

运行

Another new technique in this example is that we initialize the slice by writing a composite literal, which consists of the type of the slice followed by its elements in braces:

这个例子中的另一个新技术是，我们通过写一个复合字面来初始化slice，它由slice的类型和它在大括号中的元素组成：

```go linenums="1"
    slice := []int{0, 1, 2, 3, 4}
```

The `Append` function is interesting for another reason. Not only can we append elements, we can append a whole second slice by "exploding" the slice into arguments using the `...` notation at the call site:

Append函数之所以有趣，还有一个原因。我们不仅可以追加元素，还可以通过在调用处使用......符号将切片 "爆炸 "成参数来追加整个第二个切片：

```go linenums="1"
    slice1 := []int{0, 1, 2, 3, 4}
    slice2 := []int{55, 66, 77}
    fmt.Println(slice1)
    slice1 = Append(slice1, slice2...) // The '...' is essential!
    fmt.Println(slice1)
```

Run 运行

Of course, we can make `Append` more efficient by allocating no more than once, building on the innards of `Extend`:

当然，我们可以在Extend的内部基础上，通过不超过一次的分配，使Append更有效率：

```go linenums="1"
// Append appends the elements to the slice.
// Efficient version.
func Append(slice []int, elements ...int) []int {
    n := len(slice)
    total := len(slice) + len(elements)
    if total > cap(slice) {
        // Reallocate. Grow to 1.5 times the new size, so we can still grow.
        newSize := total*3/2 + 1
        newSlice := make([]int, total, newSize)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[:total]
    copy(slice[n:], elements)
    return slice
}
```

Here, notice how we use `copy` twice, once to move the slice data to the newly allocated memory, and then to copy the appending items to the end of the old data.

在这里，注意到我们是如何使用copy两次的，一次是将slice数据移动到新分配的内存中，另一次是将追加的项目复制到旧数据的末尾。

Try it; the behavior is the same as before:

试试吧，行为和之前一样：

```go linenums="1"
    slice1 := []int{0, 1, 2, 3, 4}
    slice2 := []int{55, 66, 77}
    fmt.Println(slice1)
    slice1 = Append(slice1, slice2...) // The '...' is essential!
    fmt.Println(slice1)
```

Run 运行

## Append: The built-in function - Append：内置函数

And so we arrive at the motivation for the design of the `append` built-in function. It does exactly what our `Append` example does, with equivalent efficiency, but it works for any slice type.

于是我们就有了设计append内置函数的动机。它的作用与我们的Append例子完全一样，效率相当，但它对任何片断类型都有效。

A weakness of Go is that any generic-type operations must be provided by the run-time. Some day that may change, but for now, to make working with slices easier, Go provides a built-in generic `append` function. It works the same as our `int` slice version, but for *any* slice type.

Go的一个弱点是，任何通用类型的操作都必须由运行时提供。有一天这可能会改变，但是现在，为了使处理分片更容易，Go提供了一个内置的通用追加函数。它的工作原理与我们的int slice版本相同，但适用于任何slice类型。

Remember, since the slice header is always updated by a call to `append`, you need to save the returned slice after the call. In fact, the compiler won’t let you call append without saving the result.

请记住，由于片头总是通过调用append来更新，您需要在调用后保存返回的片头。事实上，编译器不会让您调用append而不保存结果。

Here are some one-liners intermingled with print statements. Try them, edit them and explore:

这里有一些夹杂着打印语句的单行代码。试试它们，编辑它们，并进行探索：

```go linenums="1"
    // Create a couple of starter slices.
    slice := []int{1, 2, 3}
    slice2 := []int{55, 66, 77}
    fmt.Println("Start slice: ", slice)
    fmt.Println("Start slice2:", slice2)

    // Add an item to a slice.
    slice = append(slice, 4)
    fmt.Println("Add one item:", slice)

    // Add one slice to another.
    slice = append(slice, slice2...)
    fmt.Println("Add one slice:", slice)

    // Make a copy of a slice (of int).
    slice3 := append([]int(nil), slice...)
    fmt.Println("Copy a slice:", slice3)

    // Copy a slice to the end of itself.
    fmt.Println("Before append to self:", slice)
    slice = append(slice, slice...)
    fmt.Println("After append to self:", slice)
```

Run 运行

It’s worth taking a moment to think about the final one-liner of that example in detail to understand how the design of slices makes it possible for this simple call to work correctly.

值得花点时间详细思考一下这个例子的最后一句话，以理解切片的设计是如何使这个简单的调用能够正确工作的。

There are lots more examples of `append`, `copy`, and other ways to use slices on the community-built ["Slice Tricks" Wiki page](https://go.dev/wiki/SliceTricks).

在社区建立的 "Slice Tricks "Wiki页面上有更多关于append、copy和其他使用slices的例子。

## Nil

As an aside, with our newfound knowledge we can see what the representation of a `nil` slice is. Naturally, it is the zero value of the slice header:

作为一个旁观者，利用我们新发现的知识，我们可以看到nil片的表示方法是什么。自然，它是切片头的零值：

```go linenums="1"
sliceHeader{
    Length:        0,
    Capacity:      0,
    ZerothElement: nil,
}
```

or just 或只是

```go linenums="1"
sliceHeader{}
```

The key detail is that the element pointer is `nil` too. The slice created by

关键的细节是，元素指针也是nil。通过以下方式创建的片断

```go linenums="1"
array[0:0]
```

has length zero (and maybe even capacity zero) but its pointer is not `nil`, so it is not a nil slice.

数组[0:0]所创建的片断，其长度为零（甚至容量为零），但其指针并非为零。

As should be clear, an empty slice can grow (assuming it has non-zero capacity), but a `nil` slice has no array to put values in and can never grow to hold even one element.

应该很清楚，一个空的片断可以增长（假设它的容量不是零），但是一个nil片断没有数组可以放值，而且永远不可能增长到容纳一个元素。



That said, a `nil` slice is functionally equivalent to a zero-length slice, even though it points to nothing. It has length zero and can be appended to, with allocation. As an example, look at the one-liner above that copies a slice by appending to a `nil` slice.

也就是说，一个nil slice在功能上等同于一个零长度的slice，尽管它没有指向任何东西。它的长度为零，并且可以通过分配被追加到上面。作为一个例子，请看上面的单行代码，它通过附加到一个nil slice来复制一个slice。

## Strings 字符串

Now a brief section about strings in Go in the context of slices. 

现在简单介绍一下Go中与分片相关的字符串。

Strings are actually very simple: they are just read-only slices of bytes with a bit of extra syntactic support from the language.

字符串其实很简单：它们只是只读的字节片，有语言的额外语法支持。

Because they are read-only, there is no need for a capacity (you can’t grow them), but otherwise for most purposes you can treat them just like read-only slices of bytes.

因为它们是只读的，所以不需要容量（您不能增长它们），但是对于大多数目的来说，您可以把它们当作只读的字节片。

For starters, we can index them to access individual bytes:

对于初学者来说，我们可以通过索引来访问单个字节：

```go linenums="1"
slash := "/usr/ken"[0] // yields the byte value '/'.
```

We can slice a string to grab a substring:

我们可以对一个字符串进行切分以获取一个子串：

```go linenums="1"
usr := "/usr/ken"[0:4] // yields the string "/usr"
```

It should be obvious now what’s going on behind the scenes when we slice a string.

现在应该很明显，当我们切分一个字符串时，背后发生了什么。

We can also take a normal slice of bytes and create a string from it with the simple conversion:

我们也可以用一个普通的字节切片，通过简单的转换从它那里创建一个字符串。

```go linenums="1"
str := string(slice)
```

and go in the reverse direction as well:

并以相反的方向进行：

```go linenums="1"
slice := []byte(usr)
```

The array underlying a string is hidden from view; there is no way to access its contents except through the string. That means that when we do either of these conversions, a copy of the array must be made. Go takes care of this, of course, so you don’t have to. After either of these conversions, modifications to the array underlying the byte slice don’t affect the corresponding string.

字符串底层的数组是隐藏的；除了通过字符串，没有办法访问其内容。这意味着当我们进行上述任何一种转换时，必须对数组进行复制。当然，Go会处理这个问题，所以您不必这样做。在这两种转换之后，对字节片下的数组的修改不会影响到相应的字符串。

An important consequence of this slice-like design for strings is that creating a substring is very efficient. All that needs to happen is the creation of a two-word string header. Since the string is read-only, the original string and the string resulting from the slice operation can share the same array safely.

字符串的这种片状设计的一个重要结果是，创建子串是非常有效的。所需要做的就是创建一个两个字的字符串头。由于字符串是只读的，原始字符串和切片操作产生的字符串可以安全地共享同一个数组。

A historical note: The earliest implementation of strings always allocated, but when slices were added to the language, they provided a model for efficient string handling. Some of the benchmarks saw huge speedups as a result.

一个历史性的说明：最早的字符串的实现总是分配的，但是当切片被添加到语言中时，它们提供了一个高效的字符串处理模型。一些基准测试因此而出现了巨大的速度提升。

There’s much more to strings, of course, and a [separate blog post](https://blog.golang.org/strings) covers them in greater depth.

当然，字符串还有很多东西，另外一篇博文会更深入地介绍它们。

## Conclusion 总结

To understand how slices work, it helps to understand how they are implemented. There is a little data structure, the slice header, that is the item associated with the slice variable, and that header describes a section of a separately allocated array. When we pass slice values around, the header gets copied but the array it points to is always shared.

要理解分片的工作原理，有助于理解它们是如何实现的。有一个小的数据结构，即分片头，是与分片变量相关的项目，该头描述了一个单独分配的数组的一部分。当我们传递分片值时，头被复制，但它所指向的数组始终是共享的。

Once you appreciate how they work, slices become not only easy to use, but powerful and expressive, especially with the help of the `copy` and `append` built-in functions.

一旦您理解了它们的工作原理，分片不仅容易使用，而且功能强大，表现力强，特别是在复制和追加内置函数的帮助下。

## More reading 更多阅读

There’s lots to find around the intertubes about slices in Go. As mentioned earlier, the ["Slice Tricks" Wiki page](https://go.dev/wiki/SliceTricks) has many examples. The [Go Slices](https://blog.golang.org/go-slices-usage-and-internals) blog post describes the memory layout details with clear diagrams. Russ Cox’s [Go Data Structures](https://research.swtch.com/godata) article includes a discussion of slices along with some of Go’s other internal data structures.

在互联网上有很多关于Go中的分片的内容可以找到。如前所述，"切片技巧 "维基页面有许多例子。Go Slices 博文以清晰的图表描述了内存布局的细节。Russ Cox的Go数据结构文章包括对分片的讨论，以及Go的一些其他内部数据结构。

There is much more material available, but the best way to learn about slices is to use them.

还有很多可用的材料，但学习分片的最好方法是使用它们。
