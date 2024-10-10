+++
title = "内置函数"
date = 2024-07-13T11:00:48+08:00
weight = 600
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 目前有18个

## append()

### 定义

```go
func append(slice []Type, elems ...Type) []Type
```

```txt
PS F:\Hugos\go_docs> go doc builtin.append
package builtin // import "builtin"

func append(slice []Type, elems ...Type) []Type
    The append built-in function appends elements to the end of a slice.
    If it has sufficient capacity, the destination is resliced to accommodate
    the new elements. If it does not, a new underlying array will be allocated.
    Append returns the updated slice. It is therefore necessary to store the
    result of append, often in the variable holding the slice itself:

        slice = append(slice, elem1, elem2)
        slice = append(slice, anotherSlice...)

    As a special case, it is legal to append a string to a byte slice, like
    this:

        slice = append([]byte("hello "), "world"...)
```



### 注意

- 只能用于切片类型的变量！

- append函数是一个[可变参数]({{< ref "/langSpec/Types#function-types-函数型">}})函数！

### 特例

​	**作为一个特例**，如果第一个实参的[核心类型]({{< ref "/langSpec/PropertiesOfTypesAndValues#core-types-核心类型">}})是`[]byte`，`append`也接受核心类型是[bytestring]({{< ref "/langSpec/PropertiesOfTypesAndValues#core-types-核心类型">}})的第二个实参，后面跟随`...` 。这种形式追加了字节切片或字符串的字节。

```go
package main

import "fmt"

func main() {
	s := append([]byte("你好世界！"), "你好中国"...)
	fmt.Printf("%q,%T\n", s, s)
}

Output:
"你好世界！你好中国",[]uint8

```



## cap()

```txt
PS F:\Hugos\go_docs> go doc builtin.cap
package builtin // import "builtin"

func cap(v Type) int
    The cap built-in function returns the capacity of v, according to its type:

        Array: the number of elements in v (same as len(v)).
        Pointer to array: the number of elements in *v (same as len(v)).
        Slice: the maximum length the slice can reach when resliced;
        if v is nil, cap(v) is zero.
        Channel: the channel buffer capacity, in units of elements;
        if v is nil, cap(v) is zero.

    For some arguments, such as a simple array expression, the result can be a
    constant. See the Go language specification's "Length and capacity" section
    for details.
    
    cap 内置函数根据 v 的类型返回其容量：

        数组：v 中的元素数量（与 len(v) 相同）。
        指向数组的指针：*v 中的元素数量（与 len(v) 相同）。
        切片：切片在重新切片时能达到的最大长度；
              如果 v 是 nil，cap(v) 为零。
        通道：通道缓冲区的容量，以元素为单位；
              如果 v 是 nil，cap(v) 为零。

    对于某些参数，如简单的数组表达式，结果可以是一个常量。详情请参见 Go 语言规范中的“长度和容量”部分。
```



## clear() <- go 1.21

```txt
PS F:\Hugos\go_docs> go doc builtin.clear
package builtin // import "builtin"

func clear[T ~[]Type | ~map[Type]Type1](t T)
    The clear built-in function clears maps and slices. For maps, clear deletes
    all entries, resulting in an empty map. For slices, clear sets all elements
    up to the length of the slice to the zero value of the respective element
    type. If the argument type is a type parameter, the type parameter's type
    set must contain only map or slice types, and clear performs the operation
    implied by the type argument.
    
    clear 内置函数用于清空映射和切片。对于映射，clear 会删除所有条目，结果是一个空映射。对于切片，clear 会将切片长度范围内的所有元素设置为该元素类型的零值。如果参数类型是类型参数，则类型参数的类型集必须只包含映射或切片类型，clear 将根据类型参数执行相应的操作。
```



## close()

```txt
PS F:\Hugos\go_docs> go doc builtin.close
package builtin // import "builtin"

func close(c chan<- Type)
    The close built-in function closes a channel, which must be either
    bidirectional or send-only. It should be executed only by the sender,
    never the receiver, and has the effect of shutting down the channel after
    the last sent value is received. After the last value has been received
    from a closed channel c, any receive from c will succeed without blocking,
    returning the zero value for the channel element. The form

        x, ok := <-c

    will also set ok to false for a closed and empty channel.
    
    close 内置函数关闭一个通道，该通道必须是双向或只发送通道。它应由发送方执行，而不是接收方，关闭通道后，最后一个发送的值被接收到。关闭后的通道 c 上的接收操作将不阻塞，并返回该通道元素的零值。形式如下：

        x, ok := <-c

    对于关闭且为空的通道，ok 将被设置为 false。
```



## copy()

```txt
PS F:\Hugos\go_docs> go doc builtin.copy
package builtin // import "builtin"

func copy(dst, src []Type) int
    The copy built-in function copies elements from a source slice into a
    destination slice. (As a special case, it also will copy bytes from a string
    to a slice of bytes.) The source and destination may overlap. Copy returns
    the number of elements copied, which will be the minimum of len(src) and
    len(dst).
    
    copy 内置函数将元素从源切片复制到目标切片。（作为特例，它也可以将字节从字符串复制到字节切片。）源和目标可以重叠。copy 返回复制的元素数量，该数量为 len(src) 和 len(dst) 的最小值。
```



## complex()

```txt
PS F:\Hugos\go_docs> go doc builtin.complex
package builtin // import "builtin"

func complex(r, i FloatType) ComplexType
    The complex built-in function constructs a complex value from two
    floating-point values. The real and imaginary parts must be of the same
    size, either float32 or float64 (or assignable to them), and the return
    value will be the corresponding complex type (complex64 for float32,
    complex128 for float64).
    
    complex 内置函数从两个浮点值构造一个复数值。实部和虚部必须是相同大小的浮点数（float32 或 float64），返回值将是相应的复数类型（float32 对应 complex64，float64 对应 complex128）。
```



## delete()

```txt
PS F:\Hugos\go_docs> go doc builtin.delete
package builtin // import "builtin"

func delete(m map[Type]Type1, key Type)
    The delete built-in function deletes the element with the specified key
    (m[key]) from the map. If m is nil or there is no such element, delete is a
    no-op.
    
    delete 内置函数从映射中删除指定键的元素 (m[key])。如果 m 为 nil 或不存在该元素，delete 不执行任何操作。
```



## imag()

```txt
PS F:\Hugos\go_docs> go doc builtin.imag
package builtin // import "builtin"

func imag(c ComplexType) FloatType
    The imag built-in function returns the imaginary part of the complex number
    c. The return value will be floating point type corresponding to the type of
    c.
    
    imag 内置函数返回复数 c 的虚部。返回值将是与 c 类型对应的浮点类型
```



## len()

```txt
PS F:\Hugos\go_docs> go doc builtin.len
package builtin // import "builtin"

func len(v Type) int
    The len built-in function returns the length of v, according to its type:

        Array: the number of elements in v.
        Pointer to array: the number of elements in *v (even if v is nil).
        Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
        String: the number of bytes in v.
        Channel: the number of elements queued (unread) in the channel buffer;
                 if v is nil, len(v) is zero.

    For some arguments, such as a string literal or a simple array expression,
    the result can be a constant. See the Go language specification's "Length
    and capacity" section for details.
    
    len 内置函数根据 v 的类型返回其长度：

        数组：v 中的元素数量。
        指向数组的指针：*v 中的元素数量（即使 v 为 nil）。
        切片或map：v 中的元素数量；如果 v 为 nil，len(v) 为零。
        字符串：v 中的字节数。
        通道：通道缓冲区中排队（未读）元素的数量；如果 v 为 nil，len(v) 为零。

    对于某些参数，如字符串字面量或简单的数组表达式，结果可以是一个常量。详情请参见 Go 语言规范中的“长度和容量”部分。
```



## make()

```txt
PS F:\Hugos\go_docs> go doc builtin.make
package builtin // import "builtin"

func make(t Type, size ...IntegerType) Type
    The make built-in function allocates and initializes an object of type
    slice, map, or chan (only). Like new, the first argument is a type,
    not a value. Unlike new, make's return type is the same as the type of its
    argument, not a pointer to it. The specification of the result depends on
    the type:

        Slice: The size specifies the length. The capacity of the slice is
        equal to its length. A second integer argument may be provided to
        specify a different capacity; it must be no smaller than the
        length. For example, make([]int, 0, 10) allocates an underlying array
        of size 10 and returns a slice of length 0 and capacity 10 that is
        backed by this underlying array.
        Map: An empty map is allocated with enough space to hold the
        specified number of elements. The size may be omitted, in which case
        a small starting size is allocated.
        Channel: The channel's buffer is initialized with the specified
        buffer capacity. If zero, or the size is omitted, the channel is
        unbuffered.
        
     make 内置函数分配并初始化一个切片、映射或通道对象。与 new 类似，第一个参数是类型而不是值。与 new 不同的是，make 的返回类型与其参数类型相同，而不是该类型的指针。结果的规格取决于类型：

        切片：size 指定长度。切片的容量等于其长度。可以提供第二个整数参数来指定不同的容量；它必须不小于长度。例如，make([]int, 0, 10) 分配一个大小为 10 的底层数组，并返回一个长度为 0、容量为 10 的切片，该切片由这个底层数组支持。
        map：分配一个足以容纳指定数量元素的空映射。可以省略 size 参数，此情况下会分配一个小的初始大小。
        通道：通道的缓冲区容量以 size 指定。如果为零或省略 size，通道为无缓冲通道。
```



## max() <- go 1.21

```txt
PS F:\Hugos\go_docs> go doc builtin.max
package builtin // import "builtin"

func max[T cmp.Ordered](x T, y ...T) T
    The max built-in function returns the largest value of a fixed number
    of arguments of cmp.Ordered types. There must be at least one argument.
    If T is a floating-point type and any of the arguments are NaNs, max will
    return NaN.
    
    max 内置函数返回固定数量 cmp.Ordered 类型参数中的最大值。至少应有一个参数。如果 T 是浮点类型且任何参数为 NaN，max 将返回 NaN。
```



## min() <- go 1.21

```txt
PS F:\Hugos\go_docs> go doc builtin.min
package builtin // import "builtin"

func min[T cmp.Ordered](x T, y ...T) T
    The min built-in function returns the smallest value of a fixed number
    of arguments of cmp.Ordered types. There must be at least one argument.
    If T is a floating-point type and any of the arguments are NaNs, min will
    return NaN.
    
    min 内置函数返回固定数量 cmp.Ordered 类型参数中的最小值。至少应有一个参数。如果 T 是浮点类型且任何参数为 NaN，min 将返回 NaN。
```



## new()

```txt
PS F:\Hugos\go_docs> go doc builtin.new
package builtin // import "builtin"

func new(Type) *Type
    The new built-in function allocates memory. The first argument is a type,
    not a value, and the value returned is a pointer to a newly allocated zero
    value of that type.
    
    new 内置函数分配内存。第一个参数是类型而不是值，返回值是指向新分配的该类型零值的指针。
```



## panic()

```txt
PS F:\Hugos\go_docs> go doc builtin.panic
package builtin // import "builtin"

func panic(v any)
    The panic built-in function stops normal execution of the current goroutine.
    When a function F calls panic, normal execution of F stops immediately.
    Any functions whose execution was deferred by F are run in the usual way,
    and then F returns to its caller. To the caller G, the invocation of F then
    behaves like a call to panic, terminating G's execution and running any
    deferred functions. This continues until all functions in the executing
    goroutine have stopped, in reverse order. At that point, the program is
    terminated with a non-zero exit code. This termination sequence is called
    panicking and can be controlled by the built-in function recover.

    Starting in Go 1.21, calling panic with a nil interface value or an untyped
    nil causes a run-time error (a different panic). The GODEBUG setting
    panicnil=1 disables the run-time error.
    
    panic 内置函数停止当前 goroutine 的正常执行。当函数 F 调用 panic 时，F 的正常执行立即停止。任何由 F 延迟执行的函数都会按常规方式运行，然后 F 返回其调用者。对于调用者 G，F 的调用表现得像是调用 panic，终止 G 的执行并运行任何延迟函数。这种情况会一直持续到所有正在执行的 goroutine 都停止，顺序是反向的。此时，程序以非零退出码终止。这一终止序列称为 panic，可以通过内置函数 recover 控制。

    从 Go 1.21 开始，用 nil 接口值或无类型 nil 调用 panic 会引发运行时错误（另一种 panic）。GODEBUG 设置 panicnil=1 可禁用运行时错误。
```

## print()

```txt
PS F:\Hugos\go_docs> go doc builtin.print
package builtin // import "builtin"

func print(args ...Type)
    The print built-in function formats its arguments in an
    implementation-specific way and writes the result to standard error. Print
    is useful for bootstrapping and debugging; it is not guaranteed to stay in
    the language.
    
    print 内置函数以实现特定方式格式化其参数，并将结果写入标准错误。print 对引导和调试很有用；它不保证会保留在语言中。
```

## println()

```txt
PS F:\Hugos\go_docs> go doc builtin.println
package builtin // import "builtin"

func println(args ...Type)
    The println built-in function formats its arguments in an
    implementation-specific way and writes the result to standard error.
    Spaces are always added between arguments and a newline is appended. Println
    is useful for bo
    
    println 内置函数以实现特定方式格式化其参数，并将结果写入标准错误。参数之间始终加空格，末尾附加换行符。println 对引导和调试很有用；它不保证会保留在语言中。
```

## real()

```txt
PS F:\Hugos\go_docs> go doc builtin.real
package builtin // import "builtin"

func real(c ComplexType) FloatType
    The real built-in function returns the real part of the complex number c.
    The return value will be floating point type corresponding to the type of c.
    
    real 内置函数返回复数 c 的实部。返回值将是与 c 类型对应的浮点类型。
```

## recover()

```txt
PS F:\Hugos\go_docs> go doc builtin.recover
package builtin // import "builtin"

func recover() any
    The recover built-in function allows a program to manage behavior of
    a panicking goroutine. Executing a call to recover inside a deferred
    function (but not any function called by it) stops the panicking sequence
    by restoring normal execution and retrieves the error value passed to the
    call of panic. If recover is called outside the deferred function it will
    not stop a panicking sequence. In this case, or when the goroutine is not
    panicking, recover returns nil.

    Prior to Go 1.21, recover would also return nil if panic is called with a
    nil argument. See [panic] for details.
    
    recover 内置函数允许程序管理发生 panic 的 goroutine 的行为。在延迟函数内部（但不包括其调用的任何函数）执行 recover 调用会通过恢复正常执行来停止 panic 序列，并获取传递给 panic 调用的错误值。如果在延迟函数之外调用 recover，它不会停止 panic 序列。在这种情况下，或者当 goroutine 未发生 panic 时，recover 返回 nil。

	在 Go 1.21 之前，如果用 nil 参数调用 panic，recover 也会返回 nil。详情请参见 [panic]。
```
