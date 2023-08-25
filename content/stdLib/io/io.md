+++
title = "io"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# io

https://pkg.go.dev/io@go1.20.1

​	io包提供了基本的I/O原语接口。它的主要任务是将这些原语的现有实现(例如os包中的实现)封装到共享的公共接口中，以抽象功能，以及一些其他相关的原语。

​	由于这些接口和原语封装了具有各种实现的底层操作，除非另有通知，否则客户端不应假设它们适合并行执行。

## 常量 

``` go 
const (
	SeekStart   = 0 // 相对于文件的起点进行寻找
	SeekCurrent = 1 // 相对于当前偏移量进行寻找
	SeekEnd     = 2 // 相对于结尾进行寻找
)
```

Seek whence values.

> 在Go语言中，Seek方法的whence参数用于确定相对于哪个位置进行偏移量的计算。

## 变量

``` go 
var EOF = errors.New("EOF")
```

​	EOF 是 Read方法在没有更多输入可用时返回的错误。(Read方法必须自己返回 EOF，而不是封装 EOF 的错误，因为调用者将使用 == 测试 EOF。)函数只应返回 EOF 来表示输入的优雅结束。如果 EOF 在结构化数据流中意外发生，则适当的错误是 ErrUnexpectedEOF 或提供更多详细信息的其他错误。

``` go 
var ErrClosedPipe = errors.New("io: read/write on closed pipe")
```

​	ErrClosedPipe是在关闭的管道上进行读取或写入操作时使用的错误。

``` go 
var ErrNoProgress = errors.New("multiple Read calls return no data or error")
```

​	ErrNoProgress 是一些 Reader 的客户端在多次调用 Read 后未返回任何数据或错误时返回的错误，通常表明 Reader 实现有问题。

``` go 
var ErrShortBuffer = errors.New("short buffer")
```

​	ErrShortBuffer表示读取所需的缓冲区比提供的缓冲区要长。

``` go 
var ErrShortWrite = errors.New("short write")
```

​	ErrShortWrite表示写入接受的字节数比请求的字节数少，但未返回明确的错误。

``` go 
var ErrUnexpectedEOF = errors.New("unexpected EOF")
```

​	ErrUnexpectedEOF表示在读取固定大小的块或数据结构的过程中遇到了EOF。

## 函数

#### func Copy 

``` go 
func Copy(dst Writer, src Reader) (written int64, err error)
```

​	Copy函数从src复制到dst，直到在src上到达EOF或发生错误。它返回复制的字节数和在复制过程中遇到的第一个错误(如果有)。

​	成功的Copy返回err==nil，而不是err==EOF。因为Copy定义为从src读取直到EOF，所以它不会将从Read返回的EOF视为要报告的错误。

​	如果src实现了`WriterTo`接口，则通过调用src.WriteTo(dst)来实现复制。否则，如果dst实现了`ReaderFrom`接口，则通过调用dst.ReadFrom(src)来实现复制。【如果src和dst都实现了所说的，以哪个为准？】

#####    Copy Example 

``` go 
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
Output:

some io.Reader stream to be read
```

##### Copy My Example 1

```go
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 创建一个字符串Reader
	reader := strings.NewReader("Hello World")

	// 创建一个文件Writer
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 将偏移量设置到结尾，同时也会把 whence 设置到结尾
	//_, _ = io.ReadAll(file)
	// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
	file.Seek(0, 2)
	//file.Seek(0, 0)
	var writer io.Writer
	writer = file

	// 调用 Copy 将reader的数据复制到writer
	written, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Copied %d bytes\n", written)

	// 确认文件已复制
	file.Seek(0, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)
}

// Output:
// Copied 11 bytes
//Hello WorldHello WorldHello WorldHello WorldHello World

```



##### Copy My Example 2



```go
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type MyReader1 struct {
	R *strings.Reader
}

func (mr *MyReader1) Read(p []byte) (n int, err error) {
	n, err = mr.R.Read(p)
	return
}

func (mr *MyReader1) WriteTo(w io.Writer) (n int64, err error) {
	fmt.Println("调用了MyReader1的WriteTo方法")

	data := make([]byte, mr.R.Size())
	//fmt.Println(mr.Read(data))
	//fmt.Println("data=", string(data))
	//fmt.Println(reflect.TypeOf(w).String())

	iw1, ok1 := w.(*MyWriter1)
	iw2, ok2 := w.(*MyWriter2)

	if !ok1 && !ok2 {
		return 0, errors.New(fmt.Sprintf("w's is neither  *MyWriter1 nor *MyWriter2, current w's type is %T", w))
	}

	if ok1 {
		iiw, ok := iw1.W.(*os.File)
		if !ok {
			return 0, errors.New(fmt.Sprintf("w.W's type is not *os.File, current type is %T", iw2.W))
		}

		// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
		iiw.Seek(0, 2)

		m, err := iiw.Write(data)
		//fmt.Println("m=", m)
		return int64(m), err
	}

	if ok2 {
		iiw, ok := iw2.W.(*os.File)
		if !ok {
			return 0, errors.New(fmt.Sprintf("w.W's type is not *os.File, current type is %T", iw2.W))
		}

		// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
		iiw.Seek(0, 2)

		m, err := iiw.Write(data)
		//fmt.Println("m=", m)
		return int64(m), err
	}

	return 0, nil
}

type MyReader2 struct {
	R *strings.Reader
}

func (mr *MyReader2) Read(p []byte) (n int, err error) {
	n, err = mr.R.Read(p)
	return
}

type MyWriter1 struct {
	W io.Writer
}

func (mw *MyWriter1) Write(p []byte) (n int, err error) {
	n, err = mw.W.Write(p)
	return
}

func (mw *MyWriter1) ReadFrom(r io.Reader) (n int64, err error) {
	fmt.Println("调用了MyWriter1的ReadFrom方法")

	ir1, ok1 := r.(*MyReader1)
	ir2, ok2 := r.(*MyReader2)

	if !ok1 && !ok2 {
		return 0, errors.New(fmt.Sprintf("r's is neither  *MyReader1 nor *MyReader2, current r's type is %T", r))
	}

	if ok1 {
		data := make([]byte, ir1.R.Size())
		r.Read(data)
		m, err := mw.W.Write(data)
		return int64(m), err
	}

	if ok2 {
		data := make([]byte, ir2.R.Size())
		r.Read(data)
		m, err := mw.W.Write(data)
		return int64(m), err
	}
	return 0, nil
}

type MyWriter2 struct {
	W io.Writer
}

func (mw *MyWriter2) Write(p []byte) (n int, err error) {
	n, err = mw.W.Write(p)
	return
}

func main() {
	fmt.Println("情况1：Copy 中的 src 实现了`WriterTo`接口 且 dst 实现了`ReaderFrom`接口 ")
	file1, err := os.OpenFile("data1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file1.Close()
	fmt.Println("起始文件字节数：", FileSize(file1))

	_, err = io.Copy(&MyWriter1{W: file1}, &MyReader1{R: strings.NewReader("Hello World!")})
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file1))

	fmt.Println("情况2：Copy 中的 src 实现了`WriterTo`接口 但 dst 没有实现`ReaderFrom`接口 ")
	file2, err := os.OpenFile("data2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file2.Close()
	fmt.Println("起始文件字节数：", FileSize(file2))

	_, err = io.Copy(&MyWriter2{W: file2}, &MyReader1{R: strings.NewReader("Hello World!")})
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file2))

	fmt.Println("情况3：Copy 中的 src 没有实现`WriterTo`接口 但 dst 实现了`ReaderFrom`接口 ")
	file3, err := os.OpenFile("data3.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file3.Close()
	fmt.Println("起始文件字节数：", FileSize(file3))

	_, err = io.Copy(&MyWriter1{W: file3}, &MyReader2{R: strings.NewReader("Hello World!")})
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file3))

	fmt.Println("情况4：Copy 中的 src 没有实现`WriterTo`接口 且 dst 没有实现`ReaderFrom`接口 ")
	file4, err := os.OpenFile("data4.txt", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file4.Close()
	fmt.Println("起始文件字节数：", FileSize(file4))

	_, err = io.Copy(&MyWriter2{W: file4}, &MyReader2{R: strings.NewReader("Hello World!")})
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file4))
}

func FileSize(file *os.File) int {
	file.Seek(0, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return len(data)
}
// Output:
//情况1：Copy 中的 src 实现了`WriterTo`接口 且 dst 实现了`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyReader1的WriteTo方法
//Copy操作后文件字节数： 12
//情况2：Copy 中的 src 实现了`WriterTo`接口 但 dst 没有实现`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyReader1的WriteTo方法
//Copy操作后文件字节数： 12
//情况3：Copy 中的 src 没有实现`WriterTo`接口 但 dst 实现了`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyWriter1的ReadFrom方法
//Copy操作后文件字节数： 12
//情况4：Copy 中的 src 没有实现`WriterTo`接口 且 dst 没有实现`ReaderFrom`接口
//起始文件字节数： 0
//Copy操作后文件字节数： 12 
```



#### func CopyBuffer  <- go1.5

``` go 
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
```

​	CopyBuffer函数与Copy函数相同，只是它通过提供的缓冲区(如果需要)进行阶段处理，而不是分配临时缓冲区。如果`buf`为nil，则会分配一个；否则，如果长度为零，则CopyBuffer函数会出现panic。

​	如果src实现了WriterTo或dst实现了ReaderFrom，则不会使用buf执行复制。

#####    CopyBuffer Example 

```go 
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}

}
Output:

first reader
second reader
```

##### CopyBuffer My Example 

```go
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type MyReader1 struct {
	R *strings.Reader
}

func (mr *MyReader1) Read(p []byte) (n int, err error) {
	n, err = mr.R.Read(p)
	return
}

func (mr *MyReader1) WriteTo(w io.Writer) (n int64, err error) {
	fmt.Println("调用了MyReader1的WriteTo方法")

	data := make([]byte, mr.R.Size())
	//fmt.Println(mr.Read(data))
	//fmt.Println("data=", string(data))
	//fmt.Println(reflect.TypeOf(w).String())

	iw1, ok1 := w.(*MyWriter1)
	iw2, ok2 := w.(*MyWriter2)

	if !ok1 && !ok2 {
		return 0, errors.New(fmt.Sprintf("w's is neither  *MyWriter1 nor *MyWriter2, current w's type is %T", w))
	}

	if ok1 {
		iiw, ok := iw1.W.(*os.File)
		if !ok {
			return 0, errors.New(fmt.Sprintf("w.W's type is not *os.File, current type is %T", iw2.W))
		}

		// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
		iiw.Seek(0, 2)

		m, err := iiw.Write(data)
		//fmt.Println("m=", m)
		return int64(m), err
	}

	if ok2 {
		iiw, ok := iw2.W.(*os.File)
		if !ok {
			return 0, errors.New(fmt.Sprintf("w.W's type is not *os.File, current type is %T", iw2.W))
		}

		// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
		iiw.Seek(0, 2)

		m, err := iiw.Write(data)
		//fmt.Println("m=", m)
		return int64(m), err
	}

	return 0, nil
}

type MyReader2 struct {
	R *strings.Reader
}

func (mr *MyReader2) Read(p []byte) (n int, err error) {
	n, err = mr.R.Read(p)
	return
}

type MyWriter1 struct {
	W io.Writer
}

func (mw *MyWriter1) Write(p []byte) (n int, err error) {
	n, err = mw.W.Write(p)
	return
}

func (mw *MyWriter1) ReadFrom(r io.Reader) (n int64, err error) {
	fmt.Println("调用了MyWriter1的ReadFrom方法")

	ir1, ok1 := r.(*MyReader1)
	ir2, ok2 := r.(*MyReader2)

	if !ok1 && !ok2 {
		return 0, errors.New(fmt.Sprintf("r's is neither  *MyReader1 nor *MyReader2, current r's type is %T", r))
	}

	if ok1 {
		data := make([]byte, ir1.R.Size())
		r.Read(data)
		m, err := mw.W.Write(data)
		return int64(m), err
	}

	if ok2 {
		data := make([]byte, ir2.R.Size())
		r.Read(data)
		m, err := mw.W.Write(data)
		return int64(m), err
	}
	return 0, nil
}

type MyWriter2 struct {
	W io.Writer
}

func (mw *MyWriter2) Write(p []byte) (n int, err error) {
	n, err = mw.W.Write(p)
	return
}

func main() {
	fmt.Println("情况1：Copy 中的 src 实现了`WriterTo`接口 且 dst 实现了`ReaderFrom`接口 ")
	file1, err := os.OpenFile("data1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file1.Close()
	fmt.Println("起始文件字节数：", FileSize(file1))

	_, err = io.CopyBuffer(&MyWriter1{W: file1}, &MyReader1{R: strings.NewReader("Hello World!")}, make([]byte, 8))
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file1))

	fmt.Println("情况2：Copy 中的 src 实现了`WriterTo`接口 但 dst 没有实现`ReaderFrom`接口 ")
	file2, err := os.OpenFile("data2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file2.Close()
	fmt.Println("起始文件字节数：", FileSize(file2))

	_, err = io.CopyBuffer(&MyWriter2{W: file2}, &MyReader1{R: strings.NewReader("Hello World!")}, make([]byte, 8))
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file2))

	fmt.Println("情况3：Copy 中的 src 没有实现`WriterTo`接口 但 dst 实现了`ReaderFrom`接口 ")
	file3, err := os.OpenFile("data3.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file3.Close()
	fmt.Println("起始文件字节数：", FileSize(file3))

	_, err = io.CopyBuffer(&MyWriter1{W: file3}, &MyReader2{R: strings.NewReader("Hello World!")}, make([]byte, 8))
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file3))

	fmt.Println("情况4：Copy 中的 src 没有实现`WriterTo`接口 且 dst 没有实现`ReaderFrom`接口 ")
	file4, err := os.OpenFile("data4.txt", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file4.Close()
	fmt.Println("起始文件字节数：", FileSize(file4))

	_, err = io.CopyBuffer(&MyWriter2{W: file4}, &MyReader2{R: strings.NewReader("Hello World!")}, make([]byte, 8))
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file4))
}

func FileSize(file *os.File) int {
	file.Seek(0, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return len(data)
}

// Output:
//情况1：Copy 中的 src 实现了`WriterTo`接口 且 dst 实现了`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyReader1的WriteTo方法
//Copy操作后文件字节数： 12
//情况2：Copy 中的 src 实现了`WriterTo`接口 但 dst 没有实现`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyReader1的WriteTo方法
//Copy操作后文件字节数： 12
//情况3：Copy 中的 src 没有实现`WriterTo`接口 但 dst 实现了`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyWriter1的ReadFrom方法
//Copy操作后文件字节数： 12
//情况4：Copy 中的 src 没有实现`WriterTo`接口 且 dst 没有实现`ReaderFrom`接口
//起始文件字节数： 0
//Copy操作后文件字节数： 12
```



#### func CopyN 

``` go 
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

​	CopyN函数从src复制n个字节(或直到出现错误)到dst。它返回已复制的字节数和在复制过程中遇到的最早的错误。如果err == nil，则written == n。

​	如果dst实现了ReaderFrom接口，则使用它来实现复制。

#####    CopyN Example 

```go 
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.CopyN(os.Stdout, r, 4); err != nil {
		log.Fatal(err)
	}

}
Output:

some
```

##### CopyN My Example

```go
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type MyReader1 struct {
	R *strings.Reader
}

func (mr *MyReader1) Read(p []byte) (n int, err error) {
	n, err = mr.R.Read(p)
	return
}

func (mr *MyReader1) WriteTo(w io.Writer) (n int64, err error) {
	fmt.Println("调用了MyReader1的WriteTo方法")

	data := make([]byte, mr.R.Size())
	//fmt.Println(mr.Read(data))
	//fmt.Println("data=", string(data))
	//fmt.Println(reflect.TypeOf(w).String())

	iw1, ok1 := w.(*MyWriter1)
	iw2, ok2 := w.(*MyWriter2)

	if !ok1 && !ok2 {
		return 0, errors.New(fmt.Sprintf("w's is neither  *MyWriter1 nor *MyWriter2, current w's type is %T", w))
	}

	if ok1 {
		iiw, ok := iw1.W.(*os.File)
		if !ok {
			return 0, errors.New(fmt.Sprintf("w.W's type is not *os.File, current type is %T", iw2.W))
		}

		// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
		iiw.Seek(0, 2)

		m, err := iiw.Write(data)
		//fmt.Println("m=", m)
		return int64(m), err
	}

	if ok2 {
		iiw, ok := iw2.W.(*os.File)
		if !ok {
			return 0, errors.New(fmt.Sprintf("w.W's type is not *os.File, current type is %T", iw2.W))
		}

		// whence: 0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾
		iiw.Seek(0, 2)

		m, err := iiw.Write(data)
		//fmt.Println("m=", m)
		return int64(m), err
	}

	return 0, nil
}

type MyReader2 struct {
	R *strings.Reader
}

func (mr *MyReader2) Read(p []byte) (n int, err error) {
	n, err = mr.R.Read(p)
	return
}

type MyWriter1 struct {
	W io.Writer
}

func (mw *MyWriter1) Write(p []byte) (n int, err error) {
	n, err = mw.W.Write(p)
	return
}

func (mw *MyWriter1) ReadFrom(r io.Reader) (n int64, err error) {
	fmt.Println("调用了MyWriter1的ReadFrom方法")

	ir1, ok1 := r.(*MyReader1)
	ir2, ok2 := r.(*MyReader2)

	if !ok1 && !ok2 {
		return 0, errors.New(fmt.Sprintf("r's is neither  *MyReader1 nor *MyReader2, current r's type is %T", r))
	}

	if ok1 {
		data := make([]byte, ir1.R.Size())
		r.Read(data)
		m, err := mw.W.Write(data)
		return int64(m), err
	}

	if ok2 {
		data := make([]byte, ir2.R.Size())
		r.Read(data)
		m, err := mw.W.Write(data)
		return int64(m), err
	}
	return 0, nil
}

type MyWriter2 struct {
	W io.Writer
}

func (mw *MyWriter2) Write(p []byte) (n int, err error) {
	n, err = mw.W.Write(p)
	return
}

func main() {
	fmt.Println("情况1：Copy 中的 src 实现了`WriterTo`接口 且 dst 实现了`ReaderFrom`接口 ")
	file1, err := os.OpenFile("data1.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file1.Close()
	fmt.Println("起始文件字节数：", FileSize(file1))

	_, err = io.CopyN(&MyWriter1{W: file1}, &MyReader1{R: strings.NewReader("Hello World!")}, 5)
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file1))

	fmt.Println("情况2：Copy 中的 src 实现了`WriterTo`接口 但 dst 没有实现`ReaderFrom`接口 ")
	file2, err := os.OpenFile("data2.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file2.Close()
	fmt.Println("起始文件字节数：", FileSize(file2))

	_, err = io.CopyN(&MyWriter2{W: file2}, &MyReader1{R: strings.NewReader("Hello World!")}, 5)
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file2))

	fmt.Println("情况3：Copy 中的 src 没有实现`WriterTo`接口 但 dst 实现了`ReaderFrom`接口 ")
	file3, err := os.OpenFile("data3.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file3.Close()
	fmt.Println("起始文件字节数：", FileSize(file3))

	_, err = io.CopyN(&MyWriter1{W: file3}, &MyReader2{R: strings.NewReader("Hello World!")}, 5)
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file3))

	fmt.Println("情况4：Copy 中的 src 没有实现`WriterTo`接口 且 dst 没有实现`ReaderFrom`接口 ")
	file4, err := os.OpenFile("data4.txt", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}
	defer file4.Close()
	fmt.Println("起始文件字节数：", FileSize(file4))

	_, err = io.CopyN(&MyWriter2{W: file4}, &MyReader2{R: strings.NewReader("Hello World!")}, 5)
	if err != nil {
		fmt.Println("发生错误：", err)
	}
	fmt.Println("Copy操作后文件字节数：", FileSize(file4))
}

func FileSize(file *os.File) int {
	file.Seek(0, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return len(data)
}

// Output:
//情况1：Copy 中的 src 实现了`WriterTo`接口 且 dst 实现了`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyWriter1的ReadFrom方法
//发生错误： r's is neither  *MyReader1 nor *MyReader2, current r's type is *io.LimitedReader
//Copy操作后文件字节数： 0
//情况2：Copy 中的 src 实现了`WriterTo`接口 但 dst 没有实现`ReaderFrom`接口
//起始文件字节数： 0
//Copy操作后文件字节数： 5
//情况3：Copy 中的 src 没有实现`WriterTo`接口 但 dst 实现了`ReaderFrom`接口
//起始文件字节数： 0
//调用了MyWriter1的ReadFrom方法
//发生错误： r's is neither  *MyReader1 nor *MyReader2, current r's type is *io.LimitedReader
//Copy操作后文件字节数： 0
//情况4：Copy 中的 src 没有实现`WriterTo`接口 且 dst 没有实现`ReaderFrom`接口
//起始文件字节数： 0
//Copy操作后文件字节数： 5
```

#### func Pipe 

``` go 
func Pipe() (*PipeReader, *PipeWriter)
```

​	Pipe 函数创建一个同步的内存管道。它可以用来连接期望io.Reader的代码和期望io.Writer的代码。   

​	除非多个Read被需要来消费一个Write，否则Pipe上的读和写是一一对应的。也就是说，每次向PipeWriter写入都会阻塞，直到一个或多个从PipeReader完全消费写入数据的Read满足写入。数据是直接从Write复制到相应的Read(或Reads)；没有内部缓冲。   

​	同时调用Read和Write或者与Close是安全的。对Read的并行调用和对Write的并行调用也是安全的：各个调用将依次进行。

#####    Pipe Example 

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
Output:

some io.Reader stream to be read
```

##### Pipe My Example

```go
package main

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 创建一个管道
	r, w := io.Pipe()
	defer w.Close()
	defer r.Close()

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			str := "write " + strconv.Itoa(j) + ": Hello world"
			w.Write([]byte(str))
		}(i)
	}

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			buf := make([]byte, 512)
			n, err := r.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("read " + strconv.Itoa(j) + " -> " + string(buf[:n]))
		}(i)
	}

	// 等待goroutine结束
	wg.Wait()
}
// Output:
//read 3 -> write 9: Hello world
//read 2 -> write 1: Hello world
//read 6 -> write 3: Hello world
//read 7 -> write 8: Hello world
//read 4 -> write 7: Hello world
//read 8 -> write 4: Hello world
//read 5 -> write 2: Hello world
//read 1 -> write 5: Hello world
//read 9 -> write 6: Hello world
```

#### func ReadAll  <- go1.16

``` go 
func ReadAll(r Reader) ([]byte, error)
```

​	ReadAll函数从r读取直到出现错误或EOF，并返回它读取的数据。成功调用返回err == nil，而不是err == EOF。因为ReadAll函数定义为从src读取直到EOF，所以它不会将从Read中读取的EOF视为错误报告。

#####    ReadAll Example 

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}
Output:

Go is a general-purpose language designed with systems programming in mind.
```

##### ReadAll My Example

![image-20230825200859284](io_img/image-20230825200859284.png)

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}
// Output:
//All content in one line and no newline!
```

#### func ReadAtLeast 

``` go 
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```

​	ReadAtLeast函数从`r`中读取到`buf`，直到它读取至少`min`个字节。它返回已复制的字节数和错误(如果读取的字节数少于`min`个)。如果没有读取任何字节，则错误为EOF。如果在读取少于min个字节后出现EOF，则ReadAtLeast返回ErrUnexpectedEOF。如果min大于buf的长度，则ReadAtLeast返回ErrShortBuffer。返回时，当且仅当err == nil时，n >= min。如果r在读取至少min个字节后返回错误，则将删除该错误。

#####    ReadAtLeast Example 

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Println("error:", err)
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err)
	}

}
Output:

some io.Reader
error: short buffer
error: unexpected EOF
```

##### ReadAtLeast My Example

```go
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func MyRead(r io.Reader, bufSize, min int) {
	buf := make([]byte, bufSize)
	ir, ok := r.(*strings.Reader)
	if !ok {
		log.Fatal("类型错误")
	}
	bytesSize := ir.Size()
	// 每次都是从头开始
	ir.Seek(0, 0)
	n, err := io.ReadAtLeast(r, buf, min)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("total bytes is %d,buf=%d,min=%d,n=%d,read content is:%s\n", bytesSize, bufSize, min, n, string(buf))
	fmt.Println("-----------------------------------------")
}

func main() {
	r := strings.NewReader("Do you code with go?\n") // 21个字符
	MyRead(r, 30, 30)
	MyRead(r, 30, 21)
	MyRead(r, 30, 20)
	MyRead(r, 30, 19)

	MyRead(r, 21, 30)
	MyRead(r, 21, 21)
	MyRead(r, 21, 19)

	MyRead(r, 20, 30)
	MyRead(r, 20, 21)
	MyRead(r, 20, 20)
	MyRead(r, 20, 19)

	MyRead(r, 19, 30)
	MyRead(r, 19, 21)
	MyRead(r, 19, 20)
	MyRead(r, 19, 19)
}

// Output:
//error: unexpected EOF
//total bytes is 21,buf=30,min=30,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=30,min=21,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=30,min=20,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=30,min=19,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=21,min=30,n=0,read content is:
//-----------------------------------------
//total bytes is 21,buf=21,min=21,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=21,min=19,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=20,min=30,n=0,read content is:
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=20,min=21,n=0,read content is:
//-----------------------------------------
//total bytes is 21,buf=20,min=20,n=20,read content is:Do you code with go?
//-----------------------------------------
//total bytes is 21,buf=20,min=19,n=20,read content is:Do you code with go?
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=19,min=30,n=0,read content is:
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=19,min=21,n=0,read content is:
//-----------------------------------------
//error: short buffer
//total bytes is 21,buf=19,min=20,n=0,read content is:
//-----------------------------------------
//total bytes is 21,buf=19,min=19,n=19,read content is:Do you code with go
//-----------------------------------------
```

#### func ReadFull 

``` go 
func ReadFull(r Reader, buf []byte) (n int, err error)
```

​	ReadFull函数从`r`中精确地读取len(buf)个字节到`buf`中。它返回已复制的字节数和错误(如果读取的字节数少于len(buf)个)。如果没有读取任何字节，则错误为EOF。如果在读取一些但不是所有字节后出现EOF，则ReadFull函数返回ErrUnexpectedEOF。返回时，当且仅当err == nil时，n == len(buf)。如果`r`在读取至少len(buf)个字节后返回错误，则将删除该错误。

#####    ReadFull Example 

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}

}
Output:

some
error: unexpected EOF
```

##### ReadFull My Example

```go
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func MyRead(r io.Reader, bufSize int) {
	buf := make([]byte, bufSize)
	ir, ok := r.(*strings.Reader)
	if !ok {
		log.Fatal("类型错误")
	}
	bytesSize := ir.Size()
	// 每次都是从头开始
	ir.Seek(0, 0)

	n, err := io.ReadFull(r, buf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("total bytes is %d,buf=%d,n=%d,read content is:%s\n", bytesSize, bufSize, n, string(buf))
	fmt.Println("-----------------------------------------")
}

func main() {
	r := strings.NewReader("Do you code with go?\n") // 21个字符

	MyRead(r, 30)
	MyRead(r, 21)
	MyRead(r, 20)
	MyRead(r, 19)
}

// Output:
//error: unexpected EOF
//total bytes is 21,buf=30,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=21,n=21,read content is:Do you code with go?
//
//-----------------------------------------
//total bytes is 21,buf=20,n=20,read content is:Do you code with go?
//-----------------------------------------
//total bytes is 21,buf=19,n=19,read content is:Do you code with go
//-----------------------------------------
```

#### func WriteString 

``` go 
func WriteString(w Writer, s string) (n int, err error)
```

​	WriteString函数将字符串`s`的内容写入接受字节切片的`w`中。如果`w`实现了StringWriter，则直接调用其WriteString方法。否则，将调用w.Write一次。

#####    WriteString Example 

``` go 
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		log.Fatal(err)
	}

}
Output:

Hello World
```

##### WriteString My Example

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type MyWriter struct {
	W io.Writer
}

func (mw *MyWriter) Write(p []byte) (n int, err error) {
	return mw.W.Write([]byte(p))
}

func (mw *MyWriter) WriteString(s string) (n int, err error) {
	fmt.Println("调用了MyWriter的WriteString方法")
	return mw.W.Write([]byte(s))
}

func main() {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("起始文件字节数：", FileSize(file))

	io.WriteString(file, "Hello World!")
	fmt.Println("第一次io.WriteString操作后文件字节数：", FileSize(file))

	io.WriteString(&MyWriter{file}, "Hello World!")
	fmt.Println("第二次io.WriteString操作后文件字节数：", FileSize(file))
}

func FileSize(file *os.File) int {
	file.Seek(0, 0)
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return len(data)
}
// Output:
//起始文件字节数： 0
//第一次io.WriteString操作后文件字节数： 12
//调用了MyWriter的WriteString方法
//第二次io.WriteString操作后文件字节数： 24
```

## 类型

### type ByteReader 

``` go 
type ByteReader interface {
	ReadByte() (byte, error)
}
```

​	ByteReader是封装了ReadByte方法的接口。

​	ReadByte读取并返回输入中的下一个字节或遇到的任何错误。如果ReadByte返回错误，则没有输入字节被消耗，返回的字节值未定义。

​	ReadByte提供了逐字节处理的有效接口。如果Reader未实现ByteReader，则可以使用bufio.NewReader进行封装以添加此方法。

### type ByteScanner 

``` go 
type ByteScanner interface {
	ByteReader
	UnreadByte() error
}
```

​	ByteScanner是在基本ReadByte方法上添加了UnreadByte方法的接口。

​	UnreadByte方法使得下一次调用ReadByte方法返回上次读取的最后一个字节。如果上次操作不是对ReadByte方法的成功调用，则UnreadByte方法可能会返回错误，未读取最后一个字节(或上一个未读取字节)，或(在支持Seeker接口的实现中)将偏移量设置为当前偏移量前一个字节。

### type ByteWriter  <- go1.1

``` go 
type ByteWriter interface {
	WriteByte(c byte) error
}
```

​	ByteWriter是封装了WriteByte方法的接口。

### type Closer 

``` go 
type Closer interface {
	Close() error
}
```

​	Closer是封装了基本Close方法的接口。

​	第一次调用Close方法后的行为是未定义的。具体的实现可能会记录自己的行为。

### type LimitedReader 

``` go 
type LimitedReader struct {
	R Reader // 底层的 Reader
	N int64  // 最多可以读取的字节数
}
```

​	LimitedReader结构体从 R 中读取数据，但是限制了返回数据的数量，只有最多 N 个字节。每次调用 Read 都会更新 N 的值以反映新的剩余字节数。当 N <= 0 或者底层的 R 返回 EOF 时，Read 会返回 EOF。

#### (*LimitedReader) Read 

``` go 
func (l *LimitedReader) Read(p []byte) (n int, err error)
```

### type OffsetWriter  <- go1.20

``` go 
type OffsetWriter struct {
	// contains filtered or unexported fields
    // 包含已过滤或未导出的字段
}
```

​	OffsetWriter结构体将写入基于 offset 的数据映射到底层 Writer 中基于 base+off 的偏移位置。

#### func NewOffsetWriter  <- go1.20

``` go 
func NewOffsetWriter(w WriterAt, off int64) *OffsetWriter
```

​	NewOffsetWriter函数返回一个 OffsetWriter，它从 offset 位置开始写入 w。

#### (*OffsetWriter) Seek  <- go1.20

``` go 
func (o *OffsetWriter) Seek(offset int64, whence int) (int64, error)
```

#### (*OffsetWriter) Write  <- go1.20

``` go 
func (o *OffsetWriter) Write(p []byte) (n int, err error)
```

#### (*OffsetWriter) WriteAt  <- go1.20

``` go 
func (o *OffsetWriter) WriteAt(p []byte, off int64) (n int, err error)
```

### type PipeReader 

``` go 
type PipeReader struct {
	// contains filtered or unexported fields
    // 包含已过滤或未导出的字段
}
```

​	PipeReader结构体是管道的读取端。

#### (*PipeReader) Close 

``` go 
func (r *PipeReader) Close() error
```

​	Close方法关闭读取器；后续对管道写入端的写入将返回 ErrClosedPipe 错误。

#### (*PipeReader) CloseWithError 

``` go 
func (r *PipeReader) CloseWithError(err error) error
```

​	CloseWithError方法关闭读取器；后续对管道写入端的写入将返回 err 错误。

​	如果存在先前的错误，CloseWithError不会覆盖它并始终返回nil。

#### (*PipeReader) Read 

``` go 
func (r *PipeReader) Read(data []byte) (n int, err error)
```

​	Read方法实现标准的Read接口：它从管道中读取数据，在写入方到达或写入端关闭之前阻塞。如果写入端以错误关闭，则将该错误作为err返回；否则err为EOF。

### type PipeWriter 

``` go 
type PipeWriter struct {
	// contains filtered or unexported fields
    // 包含已过滤或未导出的字段
}
```

​	PipeWriter结构体是管道的写入方。

#### (*PipeWriter) Close 

``` go 
func (w *PipeWriter) Close() error
```

​	Close关闭写入器；随后从读取器读取数据将不会返回字节和EOF。

#### (*PipeWriter) CloseWithError 

``` go 
func (w *PipeWriter) CloseWithError(err error) error
```

​	CloseWithError方法关闭写入器；随后从读取器读取数据将不会返回字节和错误err，如果err为nil，则返回EOF。

​	如果存在先前的错误，CloseWithError不会覆盖它并始终返回nil。

#### (*PipeWriter) Write 

``` go 
func (w *PipeWriter) Write(data []byte) (n int, err error)
```

​	Write方法实现标准的Write接口：它将数据写入管道，阻塞直到一个或多个读取器消耗了所有数据或读取端关闭。如果读取端以错误关闭，则将该err返回；否则err为ErrClosedPipe。

### type ReadCloser 

``` go 
type ReadCloser interface {
	Reader
	Closer
}
```

 	ReadCloser接口组合了基本的Read和Close方法。

#### func NopCloser  <- go1.16

``` go 
func NopCloser(r Reader) ReadCloser
```

​	NopCloser函数返回一个带有no-op Close方法的ReadCloser，封装提供的Reader r。如果r实现了WriterTo，则返回的ReadCloser将通过转发调用来实现WriterTo方法。

### type ReadSeekCloser  <- go1.16

``` go 
type ReadSeekCloser interface {
	Reader
	Seeker
	Closer
}
```

​	ReadSeekCloser接口组合了基本的Read、Seek和Close方法。

### type ReadSeeker 

``` go 
type ReadSeeker interface {
	Reader
	Seeker
}
```

​	ReadSeeker接口组合了基本的Read、Seek方法。

### type ReadWriteCloser 

``` go 
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

​	ReadWriteCloser接口组合了基本的Read、Write和Close方法。

### type ReadWriteSeeker 

``` go 
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}
```

​	ReadWriteSeeker接口组合了基本的Read、Write和Seek方法。

### type ReadWriter 

``` go 
type ReadWriter interface {
	Reader
	Writer
}
```

​	ReadWriter接口组合了基本的Read和Write方法。

### type Reader 

``` go 
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

​	Reader接口封装了基本的Read方法。

​	Read方法将最多len(p)个字节读取到p中。它返回读取的字节数(0 <= n <= len(p))以及任何遇到的错误。即使Read返回n < len(p)，它也可以在调用期间使用p作为临时空间。如果一些数据可用但不足len(p)字节，则Read通常返回可用的内容而不是等待更多。

​	当Read方法在成功读取n > 0个字节后遇到错误或文件结束条件时，它返回读取的字节数。它可能从同一调用返回(非零)错误，也可能从后续调用返回错误(n == 0)。一般情况下，一个返回非零字节数并在输入流结束时返回EOF或nil错误的Reader实例，下一次Read应该返回0，EOF。

​	在考虑错误err之前，调用者应始终处理返回的n > 0字节。这样做可以正确处理在读取一些字节之后发生的I/O错误，以及两种允许的EOF行为。

​	实现Read的方法不应该返回具有nil错误的零字节计数，除非len(p) == 0。调用者应该将返回0和nil视为表示没有发生任何事情；特别是它不表示EOF。

​	实现不能保留p。

#### func LimitReader 

``` go 
func LimitReader(r Reader, n int64) Reader
```

​	LimitReader函数返回一个从 r 读取但在读取 n 个字节后停止并返回 EOF 的 Reader。底层实现是一个 `*LimitedReader`。

#####    LimitReader Example 

``` go 
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

}
Output:

some
```

#### func MultiReader 

``` go 
func MultiReader(readers ...Reader) Reader
```

​	MultiReader函数返回一个 Reader，它是提供的输入 readers 的逻辑连接。它们按顺序读取。一旦所有输入都返回 EOF，Read 将返回 EOF。如果任何读取器返回非 nil、非 EOF 错误，则 Read 将返回该错误。

#####    MultiReaderExample 

``` go 
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
Output:

first reader second reader third reader
```

#### func TeeReader 

``` go 
func TeeReader(r Reader, w Writer) Reader
```

​	TeeReader函数返回一个 Reader，它从 r 中读取并将其写入 w。通过它执行的所有对 r 的读取都将与对 w 的相应写入匹配。没有内部缓冲区——写入必须在读取完成之前完成。任何在写入时遇到的错误都将作为读取错误报告。

#####    TeeReader Example 

``` go 
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var r io.Reader = strings.NewReader("some io.Reader stream to be read\n")

	r = io.TeeReader(r, os.Stdout)

	// Everything read from r will be copied to stdout.
	if _, err := io.ReadAll(r); err != nil {
		log.Fatal(err)
	}

}
Output:

some io.Reader stream to be read
```

### type ReaderAt 

``` go 
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}
```

​	ReaderAt 是封装基本 ReadAt 方法的接口。

​	ReadAt方法从基础输入源的偏移量 off 开始将 len(p) 个字节读入 p 中。它返回读取的字节数(0 <= n <= len(p))和任何遇到的错误。

​	当 ReadAt方法返回 n < len(p) 时，它返回一个非 nil 的错误，解释为什么没有返回更多的字节。在这方面，ReadAt方法比 Read方法更严格。

​	即使 ReadAt方法返回 n < len(p)，它也可以在调用期间使用 p 中的所有字节作为临时空间。如果有一些数据可用但不是 len(p) 字节，则 ReadAt 阻塞，直到所有数据都可用或发生错误。在这方面，ReadAt 不同于 Read。

​	如果 ReadAt方法返回的 n = len(p) 字节在输入源的末尾，则 ReadAt 可能返回 err == EOF 或 err == nil。

​	如果 ReadAt方法从具有 seek 偏移量的输入源中读取，则 ReadAt方法不应影响底层 seek 偏移量，也不应受其影响。

​	ReadAt方法的客户端可以在同一输入源上并行执行 ReadAt方法调用。

​	实现不得保留 p。

### type ReaderFrom 

``` go 
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}
```

​	ReaderFrom 是封装 ReadFrom 方法的接口。

​	ReadFrom方法从 r 中读取数据，直到 EOF 或错误。返回值 n 是读取的字节数。除了 EOF 之外，在读取过程中遇到的任何错误也将返回。

​	如果可用，Copy函数将使用 ReaderFrom方法。

### type RuneReader 

``` go 
type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}
```

​	RuneReader 是封装了 ReadRune 方法的接口。

​	ReadRune方法读取一个单一的编码的 Unicode 字符并返回该字符以及其所占用的字节数。如果没有字符可用，则 err 将被设置。

### type RuneScanner 

``` go 
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}
```

​	RuneScanner 是在基本 ReadRune 方法上添加了 UnreadRune 方法的接口。

​	UnreadRune方法会导致下一次调用 ReadRune方法返回上次读取的最后一个字符。如果上次操作不是成功的 ReadRune方法调用，则 UnreadRune方法可能会返回错误、未读取最后一个字符(或最后一个未读取字符之前的字符)，或者(在支持 Seeker 接口的实现中)定位到当前偏移量之前的字符的开头。

### type SectionReader 

``` go 
type SectionReader struct {
	// contains filtered or unexported fields
    // 包含已过滤或未导出的字段
}
```

​	SectionReader 在底层的 ReaderAt方法的一部分实现了 Read、Seek 和 ReadAt方法。

#####    SectionReader Example 

``` go 
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}

}
Output:

io.Reader stream
```

#### func NewSectionReader 

``` go 
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
```

​	NewSectionReader函数返回一个 SectionReader，该 Reader 从偏移量 off 处开始从 r 读取并在 n 字节后以 EOF 结束。

#### (*SectionReader) Read 

``` go 
func (s *SectionReader) Read(p []byte) (n int, err error)
```

##### Read Example 

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	buf := make([]byte, 9)
	if _, err := s.Read(buf); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)

}
Output:

io.Reader
```



#### (*SectionReader) ReadAt 

``` go 
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
```

#####    ReadAt Example 

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	buf := make([]byte, 6)
	if _, err := s.ReadAt(buf, 10); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)

}
Output:

stream
```



#### (*SectionReader) Seek 

``` go 
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
```

#####    Seek Example 

``` go 
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	if _, err := s.Seek(10, io.SeekStart); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}

}
Output:

stream
```



#### (*SectionReader) Size 

``` go 
func (s *SectionReader) Size() int64
```

​	Size方法返回该 SectionReader 中字节的大小。

#####    Size Example 

``` go 
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	fmt.Println(s.Size())

}
Output:

17
```



### type Seeker 

``` go 
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
```

​	Seeker 是封装基本 Seek 方法的接口。

​	Seek方法设置下一个 Read 或 Write 的偏移量，根据 whence 进行解释：SeekStart 表示相对于文件的开头，SeekCurrent 表示相对于当前偏移量，SeekEnd 表示相对于结尾(例如，offset = -2 表示文件的倒数第二个字节)。Seek 返回相对于文件开头的新偏移量，或者如果有的话，返回错误。

​	寻找到文件开头之前的偏移量是一个错误。寻找任何正偏移量可能是允许的，但是如果新的偏移量超过底层对象的大小，则随后的 I/O 操作的行为取决于实现。

### type StringWriter  <- go1.12

``` go 
type StringWriter interface {
	WriteString(s string) (n int, err error)
}
```

​	StringWriter 是封装 WriteString 方法的接口。

### type WriteCloser 

``` go 
type WriteCloser interface {
	Writer
	Closer
}
```

​	WriteCloser 是组合基本的 Write 和 Close 方法的接口。

### type WriteSeeker 

``` go 
type WriteSeeker interface {
	Writer
	Seeker
}
```

​	WriteSeeker 是组合基本的 Write 和 Seek 方法的接口。

### type Writer 

``` go 
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

​	Writer 是封装基本 Write 方法的接口。

​	Write方法从 p 中写入 len(p) 个字节到基础数据流中。它返回从 p 中写入的字节数 n(0 <= n <= len(p))和任何导致写入提前停止的错误。如果它返回 n < len(p)，则 Write 必须返回非 nil 错误。Write 必须不修改切片数据，即使是暂时的也不行。

​	实现不得保留p。

``` go 
var Discard Writer = discard{}
```

​	Discard 是一个 Writer，所有的写入调用都会成功地且不进行任何操作。

#### func MultiWriter 

``` go 
func MultiWriter(writers ...Writer) Writer
```

​	MultiWriter函数创建一个将其写入复制到所有提供的 writer 的 writer，类似于 Unix 的 tee(1) 命令。

​	每个写入都写入到每个列出的 writer 中，一个接一个地。如果列出的 writer 返回错误，整个写操作将停止并返回该错误；不会继续下一个 writer。

##### MultiWriter  Example 

``` go 
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2 strings.Builder
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())

}
Output:

some io.Reader stream to be read
some io.Reader stream to be read
```



### type WriterAt 

``` go 
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}
```

​	WriterAt 是封装 WriteAt 方法的接口。

​	WriteAt方法从偏移量off处将长度为len(p)的p写入底层数据流。它返回写入的字节数(0 <= n <= len(p))和任何导致写入提前停止的错误。如果返回n < len(p)，WriteAt必须返回非nil的错误。

​	如果WriteAt方法正在向带有寻址偏移量的目标写入，则WriteAt方法不应受到影响，也不应影响基础寻址偏移量。

​	如果WriteAt方法正在向目标写入数据，则客户端可以在同一目标上并行执行WriteAt方法调用，前提是范围不重叠。

​	实现不得保留p。

### type WriterTo 

``` go 
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
```

​	WriterTo是封装WriteTo方法的接口。

​	WriteTo方法写入数据直到没有更多数据可写或出现错误。返回值n是写入的字节数。任何在写入期间遇到的错误也将返回。

​	如果可用，Copy函数将使用WriterTo方法。