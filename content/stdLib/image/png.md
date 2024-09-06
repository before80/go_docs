+++
title = "png"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/image/png@go1.23.0](https://pkg.go.dev/image/png@go1.23.0)

Package png implements a PNG image decoder and encoder.

​	png 包实现了一个 PNG 图像解码器和编码器。

The PNG specification is at https://www.w3.org/TR/PNG/.

​	PNG 规范位于 [https://www.w3.org/TR/PNG/](https://www.w3.org/TR/PNG/)。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Decode

```go
func Decode(r io.Reader) (image.Image, error)
```

Decode reads a PNG image from r and returns it as an image.Image. The type of Image returned depends on the PNG contents.

​	Decode 从 r 读取 PNG 图像并将其作为 image.Image 返回。返回的 Image 类型取决于 PNG 内容。

#### Decode Example

```go
package main

import (
	"encoding/base64"
	"fmt"
	"image/color"
	"image/png"
	"io"
	"log"
	"strings"
)

const gopher = `iVBORw0KGgoAAAANSUhEUgAAAEsAAAA8CAAAAAALAhhPAAAFfUlEQVRYw62XeWwUVRzHf2+OPbo9d7tsWyiyaZti6eWGAhISoIGKECEKCAiJJkYTiUgTMYSIosYYBBIUIxoSPIINEBDi2VhwkQrVsj1ESgu9doHWdrul7ba73WNm3vOPtsseM9MdwvvrzTs+8/t95ze/33sI5BqiabU6m9En8oNjduLnAEDLUsQXFF8tQ5oxK3vmnNmDSMtrncks9Hhtt/qeWZapHb1ha3UqYSWVl2ZmpWgaXMXGohQAvmeop3bjTRtv6SgaK/Pb9/bFzUrYslbFAmHPp+3WhAYdr+7GN/YnpN46Opv55VDsJkoEpMrY/vO2BIYQ6LLvm0ThY3MzDzzeSJeeWNyTkgnIE5ePKsvKlcg/0T9QMzXalwXMlj54z4c0rh/mzEfr+FgWEz2w6uk8dkzFAgcARAgNp1ZYef8bH2AgvuStbc2/i6CiWGj98y2tw2l4FAXKkQBIf+exyRnteY83LfEwDQAYCoK+P6bxkZm/0966LxcAAILHB56kgD95PPxltuYcMtFTWw/FKkY/6Opf3GGd9ZF+Qp6mzJxzuRSractOmJrH1u8XTvWFHINNkLQLMR+XHXvfPPHw967raE1xxwtA36IMRfkAAG29/7mLuQcb2WOnsJReZGfpiHsSBX81cvMKywYZHhX5hFPtOqPGWZCXnhWGAu6lX91ElKXSalcLXu3UaOXVay57ZSe5f6Gpx7J2MXAsi7EqSp09b/MirKSyJfnfEEgeDjl8FgDAfvewP03zZ+AJ0m9aFRM8eEHBDRKjfcreDXnZdQuAxXpT2NRJ7xl3UkLBhuVGU16gZiGOgZmrSbRdqkILuL/yYoSXHHkl9KXgqNu3PB8oRg0geC5vFmLjad6mUyTKLmF3OtraWDIfACyXqmephaDABawfpi6tqqBZytfQMqOz6S09iWXhktrRaB8Xz4Yi/8gyABDm5NVe6qq/3VzPrcjELWrebVuyY2T7ar4zQyybUCtsQ5Es1FGaZVrRVQwAgHGW2ZCRZshI5bGQi7HesyE972pOSeMM0dSktlzxRdrlqb3Osa6CCS8IJoQQQgBAbTAa5l5epO34rJszibJI8rxLfGzcp1dRosutGeb2VDNgqYrwTiPNsLxXiPi3dz7LiS1WBRBDBOnqEjyy3aQb+/bLiJzz9dIkscVBBLxMfSEac7kO4Fpkngi0ruNBeSOal+u8jgOuqPz12nryMLCniEjtOOOmpt+KEIqsEdocJjYXwrh9OZqWJQyPCTo67LNS/TdxLAv6R5ZNK9npEjbYdT33gRo4o5oTqR34R+OmaSzDBWsAIPhuRcgyoteNi9gF0KzNYWVItPf2TLoXEg+7isNC7uJkgo1iQWOfRSP9NR11RtbZZ3OMG/VhL6jvx+J1m87+RCfJChAtEBQkSBX2PnSiihc/Twh3j0h7qdYQAoRVsRGmq7HU2QRbaxVGa1D6nIOqaIWRjyRZpHMQKWKpZM5feA+lzC4ZFultV8S6T0mzQGhQohi5I8iw+CsqBSxhFMuwyLgSwbghGb0AiIKkSDmGZVmJSiKihsiyOAUs70UkywooYP0bii9GdH4sfr1UNysd3fUyLLMQN+rsmo3grHl9VNJHbbwxoa47Vw5gupIqrZcjPh9R4Nye3nRDk199V+aetmvVtDRE8/+cbgAAgMIWGb3UA0MGLE9SCbWX670TDy1y98c3D27eppUjsZ6fql3jcd5rUe7+ZIlLNQny3Rd+E5Tct3WVhTM5RBCEdiEK0b6B+/ca2gYU393nFj/n1AygRQxPIUA043M42u85+z2SnssKrPl8Mx76NL3E6eXc3be7OD+H4WHbJkKI8AU8irbITQjZ+0hQcPEgId/Fn/pl9crKH02+5o2b9T/eMx7pKoskYgAAAABJRU5ErkJggg==`

// gopherPNG creates an io.Reader by decoding the base64 encoded image data string in the gopher constant.
func gopherPNG() io.Reader { return base64.NewDecoder(base64.StdEncoding, strings.NewReader(gopher)) }

func main() {
	// This example uses png.Decode which can only decode PNG images.
	// Consider using the general image.Decode as it can sniff and decode any registered image format.
	img, err := png.Decode(gopherPNG())
	if err != nil {
		log.Fatal(err)
	}

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}
Output:
```

### func DecodeConfig

```go
func DecodeConfig(r io.Reader) (image.Config, error)
```

DecodeConfig returns the color model and dimensions of a PNG image without decoding the entire image.

​	DecodeConfig 返回 PNG 图像的颜色模型和尺寸，而无需解码整个图像。

### func Encode

```go
func Encode(w io.Writer, m image.Image) error
```

Encode writes the Image m to w in PNG format. Any Image may be encoded, but images that are not image.NRGBA might be encoded lossily.

​	Encode 将图像 m 以 PNG 格式写入 w。可以对任何图像进行编码，但不是 image.NRGBA 的图像可能会以有损方式进行编码。

#### Encode Example

```go
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	const width, height = 256, 256

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
Output:
```

## 类型

### type CompressionLevel <- go1.4

```go
type CompressionLevel int
```

CompressionLevel indicates the compression level.

​	CompressionLevel 指示压缩级别。

```go
const (
	DefaultCompression CompressionLevel = 0
	NoCompression      CompressionLevel = -1
	BestSpeed          CompressionLevel = -2
	BestCompression    CompressionLevel = -3
)
```

### type Encoder <- go1.4

```go
type Encoder struct {
	CompressionLevel CompressionLevel

	// BufferPool optionally specifies a buffer pool to get temporary
	// EncoderBuffers when encoding an image.
	BufferPool EncoderBufferPool
}
```

Encoder configures encoding PNG images.

​	Encoder 配置 PNG 图像的编码。

#### (*Encoder) Encode <- go1.4

```go
func (enc *Encoder) Encode(w io.Writer, m image.Image) error
```

Encode writes the Image m to w in PNG format.

​	Encode 将图像 m 以 PNG 格式写入 w。

### type EncoderBuffer <- go1.9

```go
type EncoderBuffer encoder
```

EncoderBuffer holds the buffers used for encoding PNG images.

​	EncoderBuffer 保存用于编码 PNG 图像的缓冲区。

### type EncoderBufferPool <- go1.9

```go
type EncoderBufferPool interface {
	Get() *EncoderBuffer
	Put(*EncoderBuffer)
}
```

EncoderBufferPool is an interface for getting and returning temporary instances of the EncoderBuffer struct. This can be used to reuse buffers when encoding multiple images.

​	EncoderBufferPool 是一个用于获取和返回 EncoderBuffer 结构的临时实例的接口。这可用于在编码多张图像时重用缓冲区。

### type FormatError

```go
type FormatError string
```

A FormatError reports that the input is not a valid PNG.

​	FormatError 报告输入不是有效的 PNG。

#### (FormatError) Error

```go
func (e FormatError) Error() string
```

### type UnsupportedError

```go
type UnsupportedError string
```

An UnsupportedError reports that the input uses a valid but unimplemented PNG feature.

​	UnsupportedError 报告输入使用了有效但未实现的 PNG 功能。

#### (UnsupportedError) Error 

``` go 
func (e UnsupportedError) Error() string
```