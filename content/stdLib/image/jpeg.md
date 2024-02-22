+++
title = "jpeg"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/image/jpeg@go1.21.3](https://pkg.go.dev/image/jpeg@go1.21.3)

Package jpeg implements a JPEG image decoder and encoder.

​	 jpeg 包实现了一个 JPEG 图像解码器和编码器。

JPEG is defined in ITU-T T.81: https://www.w3.org/Graphics/JPEG/itu-t81.pdf.

​	JPEG 在 ITU-T T.81 中定义：[https://www.w3.org/Graphics/JPEG/itu-t81.pdf](https://www.w3.org/Graphics/JPEG/itu-t81.pdf)。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/jpeg/writer.go;l=565)

``` go 
const DefaultQuality = 75
```

DefaultQuality is the default quality encoding parameter.

​	DefaultQuality 是默认质量编码参数。

## 变量

This section is empty.

## 函数

### func Decode 

``` go 
func Decode(r io.Reader) (image.Image, error)
```

Decode reads a JPEG image from r and returns it as an image.Image.

​	Decode 从 r 读取 JPEG 图像并将其作为 image.Image 返回。

### func DecodeConfig

```go
func DecodeConfig(r io.Reader) (image.Config, error)
```

DecodeConfig returns the color model and dimensions of a JPEG image without decoding the entire image.

​	DecodeConfig 返回 JPEG 图像的颜色模型和尺寸，而无需解码整个图像。

### func Encode

```go
func Encode(w io.Writer, m image.Image, o *Options) error
```

Encode writes the Image m to w in JPEG 4:2:0 baseline format with the given options. Default parameters are used if a nil *Options is passed.

​	Encode 以给定的选项将图像 m 写入 JPEG 4:2:0 基准格式。如果传递 nil *Options，则使用默认参数。

## 类型

### type FormatError 

``` go 
type FormatError string
```

A FormatError reports that the input is not a valid JPEG.

​	FormatError 报告输入不是有效的 JPEG。

#### (FormatError) Error 

``` go 
func (e FormatError) Error() string
```

### type Options 

``` go 
type Options struct {
	Quality int
}
```

Options are the encoding parameters. Quality ranges from 1 to 100 inclusive, higher is better.

​	Options 是编码参数。质量范围从 1 到 100（含），越高越好。

### type Reader <- DEPRECATED

```go
type Reader interface {
	io.ByteReader
	io.Reader
}
```

Deprecated: Reader is not used by the image/jpeg package and should not be used by others. It is kept for compatibility.

​	已弃用：image/jpeg 包不使用 Reader，其他人也不应使用它。保留它是为了兼容性。

### type UnsupportedError 

``` go 
type UnsupportedError string
```

An UnsupportedError reports that the input uses a valid but unimplemented JPEG feature.

​	UnsupportedError 报告输入使用了有效但未实现的 JPEG 功能。

#### (UnsupportedError) Error 

``` go 
func (e UnsupportedError) Error() string
```