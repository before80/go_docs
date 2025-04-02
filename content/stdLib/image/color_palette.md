+++
title = "color/palette"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/image/color/palette@go1.24.2](https://pkg.go.dev/image/color/palette@go1.24.2)

Package palette provides standard color palettes.

​	palette 包提供标准调色板。

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/color/palette/palette.go;l=23)

``` go 
var Plan9 = []color.Color{}/* 256 elements not displayed */
```

Plan9 is a 256-color palette that partitions the 24-bit RGB space into 4×4×4 subdivision, with 4 shades in each subcube. Compared to the WebSafe, the idea is to reduce the color resolution by dicing the color cube into fewer cells, and to use the extra space to increase the intensity resolution. This results in 16 gray shades (4 gray subcubes with 4 samples in each), 13 shades of each primary and secondary color (3 subcubes with 4 samples plus black) and a reasonable selection of colors covering the rest of the color cube. The advantage is better representation of continuous tones.

​	Plan9 是一个 256 色调色板，将 24 位 RGB 空间划分为 4×4×4 个子立方体，每个子立方体有 4 种色调。与 WebSafe 相比，其理念是通过将色立方体切成更少的单元来降低颜色分辨率，并利用额外的空间来提高强度分辨率。这产生了 16 种灰色（4 个灰色子立方体，每个子立方体有 4 个样本）、每种原色和次级色 13 种色调（3 个子立方体，每个子立方体有 4 个样本，外加黑色）以及覆盖色立方体其余部分的合理颜色选择。其优点是更好地表示连续色调。

This palette was used in the Plan 9 Operating System, described at https://9p.io/magic/man2html/6/color

​	此调色板用于 Plan 9 操作系统中，该系统在 https://9p.io/magic/man2html/6/color 中进行了描述

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/color/palette/palette.go;l=286)

``` go 
var WebSafe = []color.Color{}/* 216 elements not displayed */
```

WebSafe is a 216-color palette that was popularized by early versions of Netscape Navigator. It is also known as the Netscape Color Cube.

​	WebSafe 是一个 216 色调色板，由早期版本的 Netscape Navigator 普及。它也称为 Netscape 色彩立方体。

See https://en.wikipedia.org/wiki/Web_colors#Web-safe_colors for details.

​	有关详细信息，请参阅 https://en.wikipedia.org/wiki/Web_colors#Web-safe_colors。

## 函数

This section is empty.

## 类型

This section is empty.