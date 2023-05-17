+++
title = "draw"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# draw

https://pkg.go.dev/image/draw@go1.20.1



Package draw provides image composition functions.

See "The Go image/draw package" for an introduction to this package: https://golang.org/doc/articles/image_draw.html







## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [Draw](https://cs.opensource.google/go/go/+/go1.20.1:src/image/draw/draw.go;l=110) 

``` go linenums="1"
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
```

Draw calls DrawMask with a nil mask.

#### func [DrawMask](https://cs.opensource.google/go/go/+/go1.20.1:src/image/draw/draw.go;l=116) 

``` go linenums="1"
func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op)
```

DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r in dst with the result of a Porter-Duff composition. A nil mask is treated as opaque.

## 类型

### type [Drawer](https://cs.opensource.google/go/go/+/go1.20.1:src/image/draw/draw.go;l=60)  <- go1.2

``` go linenums="1"
type Drawer interface {
	// Draw aligns r.Min in dst with sp in src and then replaces the
	// rectangle r in dst with the result of drawing src on dst.
	Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}
```

Drawer contains the Draw method.

##### Example
``` go linenums="1"
```

``` go linenums="1"
var FloydSteinberg Drawer = floydSteinberg{}
```

FloydSteinberg is a Drawer that is the Src Op with Floyd-Steinberg error diffusion.

### type [Image](https://cs.opensource.google/go/go/+/go1.20.1:src/image/draw/draw.go;l=21) 

``` go linenums="1"
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}
```

Image is an image.Image with a Set method to change a single pixel.

### type [Op](https://cs.opensource.google/go/go/+/go1.20.1:src/image/draw/draw.go;l=44) 

``` go linenums="1"
type Op int
```

Op is a Porter-Duff compositing operator.

``` go linenums="1"
const (
	// Over specifies "(src in mask) over dst".
	Over Op = iota
	// Src specifies "src in mask".
	Src
)
```

#### (Op) [Draw](https://cs.opensource.google/go/go/+/go1.20.1:src/image/draw/draw.go;l=55)  <- go1.2

``` go linenums="1"
func (op Op) Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
```

Draw implements the Drawer interface by calling the Draw function with this Op.

### type [Quantizer](https://cs.opensource.google/go/go/+/go1.20.1:src/image/draw/draw.go;l=37)  <- go1.2

``` go linenums="1"
type Quantizer interface {
	// Quantize appends up to cap(p) - len(p) colors to p and returns the
	// updated palette suitable for converting m to a paletted image.
	Quantize(p color.Palette, m image.Image) color.Palette
}
```

Quantizer produces a palette for an image.

### type [RGBA64Image](https://cs.opensource.google/go/go/+/go1.20.1:src/image/draw/draw.go;l=30)  <- go1.17

``` go linenums="1"
type RGBA64Image interface {
	image.RGBA64Image
	Set(x, y int, c color.Color)
	SetRGBA64(x, y int, c color.RGBA64)
}
```

RGBA64Image extends both the Image and image.RGBA64Image interfaces with a SetRGBA64 method to change a single pixel. SetRGBA64 is equivalent to calling Set, but it can avoid allocations from converting concrete color types to the color.Color interface type.