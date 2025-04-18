+++
title = "robotgo"
date = 2025-04-18T09:26:20+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/go-vgo/robotgo
>
> 收录该文档时间： `2025-04-18T09:26:20+08:00`
>
> 版本：v0.110.7
>
> Published: Apr 11, 2025 

> Golang Desktop Automation. Control the mouse, keyboard, read the screen, process, Window Handle, image and bitmap and global event listener.
>
> ​	Golang 桌面自动化。控制鼠标、键盘、读取屏幕、进程、窗口句柄、图像和位图以及全局事件监听器。

RobotGo supports Mac, Windows, and Linux(X11); and robotgo supports arm64 and x86-amd64.

​	RobotGo 支持 macOS、Windows 和 Linux（X11）；并且兼容 arm64 和 x86‑amd64 架构。

## Docs

- [GoDoc](https://godoc.org/github.com/go-vgo/robotgo)
- [API Docs](https://github.com/go-vgo/robotgo/blob/master/docs/doc.md) (Deprecated, no updated) （已弃用，暂无更新）

## Binding

[ADB](https://github.com/vcaesar/adb), packaging android adb API.

[Robotn](https://github.com/vcaesar/robotn), binding JavaScript and other, support more language.

​	[Robotn](https://github.com/vcaesar/robotn)，绑定 JavaScript 等语言，支持更多语言调用。

## Requirements

Now, Please make sure `Golang, GCC` is installed correctly before installing RobotGo.

​	请在安装 RobotGo 之前，确保已正确安装 `Golang` 和 `GCC`。

### ALL:

```txt
Golang

GCC
```

#### For MacOS:

```sh
brew install go
```

Xcode Command Line Tools (And Privacy setting: [#277](https://github.com/go-vgo/robotgo/issues/277))

​	安装 Xcode 命令行工具（及隐私设置见 [#277](https://github.com/go-vgo/robotgo/issues/277)）：

```sh
xcode-select --install
```

#### For Windows:

```sh
winget install Golang.go
winget install MartinStorsjo.LLVM-MinGW.UCRT
```

Or [MinGW-w64](https://sourceforge.net/projects/mingw-w64/files) (Use recommended) or others Mingw [llvm-mingw](https://github.com/mstorsjo/llvm-mingw);

Download the Mingw, then set system environment variables `C:\mingw64\bin` to the Path. [Set environment variables to run GCC from command line](https://www.youtube.com/results?search_query=Set+environment+variables+to+run+GCC+from+command+line).

​	或使用 [MinGW-w64](https://sourceforge.net/projects/mingw-w64/files)（推荐）或其他如 [llvm-mingw](https://github.com/mstorsjo/llvm-mingw)；

 	下载后将 `C:\mingw64\bin` 添加到系统环境变量 Path。详细设置请参考 “[Set environment variables to run GCC from command line]”。

`Or the other GCC` (But you should compile the "libpng" with yourself when use the [bitmap](https://github.com/vcaesar/bitmap).)

​	或者使用其他 GCC，但在使用 [bitmap](https://github.com/vcaesar/bitmap) 时需自行编译 “libpng”。

#### For everything else:

```txt
GCC

X11 with the XTest extension (the Xtst library)

"Clipboard": xsel xclip


"Bitmap": libpng (Just used by the "bitmap".)

"Event-Gohook": xcb, xkb, libxkbcommon (Just used by the "hook".)
```

#### Ubuntu:

```sh
# sudo apt install golang
sudo snap install go  --classic

# gcc
sudo apt install gcc libc6-dev

# x11
sudo apt install libx11-dev xorg-dev libxtst-dev

# Clipboard
sudo apt install xsel xclip

#
# Bitmap
sudo apt install libpng++-dev

# GoHook
sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev libxkbcommon-dev
```

#### Fedora:

```sh
# x11
sudo dnf install libXtst-devel

# Clipboard
sudo dnf install xsel xclip

#
# Bitmap
sudo dnf install libpng-devel

# GoHook
sudo dnf install libxkbcommon-devel libxkbcommon-x11-devel xorg-x11-xkb-utils-devel
```

## Installation:

With Go module support (Go 1.11+), just import:

​	在支持 Go module（Go 1.11 及以上）的项目中，只需导入：

```go
import "github.com/go-vgo/robotgo"
```

Otherwise, to install the robotgo package, run the command:

```sh
go get github.com/go-vgo/robotgo
```

png.h: No such file or directory? Please see [issues/47](https://github.com/go-vgo/robotgo/issues/47).

​	如出现 “png.h: No such file or directory” 错误，请参考 [issues/47](https://github.com/go-vgo/robotgo/issues/47)。

## Update:

```sh
go get -u github.com/go-vgo/robotgo
```

Note go1.10.x C file compilation cache problem, [golang #24355](https://github.com/golang/go/issues/24355). `go mod vendor` problem, [golang #26366](https://github.com/golang/go/issues/26366).

​	注意 go1.10.x 的 C 文件编译缓存问题，见 [golang/go#24355]；`go mod vendor` 问题，见 [golang/go#26366]。

## Examples

### Mouse

``` go
package main

import (
  "fmt"
  "github.com/go-vgo/robotgo"
)

func main() {
  robotgo.MouseSleep = 300

  robotgo.Move(100, 100)
  fmt.Println(robotgo.Location())
  robotgo.Move(100, -200) // multi screen supported
  robotgo.MoveSmooth(120, -150)
  fmt.Println(robotgo.Location())

  robotgo.ScrollDir(10, "up")
  robotgo.ScrollDir(20, "right")

  robotgo.Scroll(0, -10)
  robotgo.Scroll(100, 0)

  robotgo.MilliSleep(100)
  robotgo.ScrollSmooth(-10, 6)
  // robotgo.ScrollRelative(10, -100)

  robotgo.Move(10, 20)
  robotgo.MoveRelative(0, -10)
  robotgo.DragSmooth(10, 10)

  robotgo.Click("wheelRight")
  robotgo.Click("left", true)
  robotgo.MoveSmooth(100, 200, 1.0, 10.0)

  robotgo.Toggle("left")
  robotgo.Toggle("left", "up")
}
```

### Keyboard

``` go
package main

import (
  "fmt"

  "github.com/go-vgo/robotgo"
)

func main() {
  robotgo.TypeStr("Hello World")
  robotgo.TypeStr("だんしゃり", 0, 1)
  // robotgo.TypeStr("テストする")

  robotgo.TypeStr("Hi, Seattle space needle, Golden gate bridge, One world trade center.")
  robotgo.TypeStr("Hi galaxy, hi stars, hi MT.Rainier, hi sea. こんにちは世界.")
  robotgo.Sleep(1)

  // ustr := uint32(robotgo.CharCodeAt("Test", 0))
  // robotgo.UnicodeType(ustr)

  robotgo.KeySleep = 100
  robotgo.KeyTap("enter")
  // robotgo.TypeStr("en")
  robotgo.KeyTap("i", "alt", "cmd")

  arr := []string{"alt", "cmd"}
  robotgo.KeyTap("i", arr)

  robotgo.MilliSleep(100)
  robotgo.KeyToggle("a")
  robotgo.KeyToggle("a", "up")

  robotgo.WriteAll("Test")
  text, err := robotgo.ReadAll()
  if err == nil {
    fmt.Println(text)
  }
}
```

### Screen

``` go
package main

import (
  "fmt"
  "strconv"

  "github.com/go-vgo/robotgo"
  "github.com/vcaesar/imgo"
)

func main() {
  x, y := robotgo.Location()
  fmt.Println("pos: ", x, y)

  color := robotgo.GetPixelColor(100, 200)
  fmt.Println("color---- ", color)

  sx, sy := robotgo.GetScreenSize()
  fmt.Println("get screen size: ", sx, sy)

  bit := robotgo.CaptureScreen(10, 10, 30, 30)
  defer robotgo.FreeBitmap(bit)

  img := robotgo.ToImage(bit)
  imgo.Save("test.png", img)

  num := robotgo.DisplaysNum()
  for i := 0; i < num; i++ {
    robotgo.DisplayID = i
    img1, _ := robotgo.CaptureImg()
    path1 := "save_" + strconv.Itoa(i)
    robotgo.Save(img1, path1+".png")
    robotgo.SaveJpeg(img1, path1+".jpeg", 50)

    img2, _ := robotgo.CaptureImg(10, 10, 20, 20)
    robotgo.Save(img2, "test_"+strconv.Itoa(i)+".png")

    x, y, w, h := robotgo.GetDisplayBounds(i)
    img3, err := robotgo.CaptureImg(x, y, w, h)
    fmt.Println("Capture error: ", err)
    robotgo.Save(img3, path1+"_1.png")
  }
}
```

### Bitmap

``` go
package main

import (
  "fmt"

  "github.com/go-vgo/robotgo"
  "github.com/vcaesar/bitmap"
)

func main() {
  bit := robotgo.CaptureScreen(10, 20, 30, 40)
  // use `defer robotgo.FreeBitmap(bit)` to free the bitmap
  defer robotgo.FreeBitmap(bit)

  fmt.Println("bitmap...", bit)
  img := robotgo.ToImage(bit)
  // robotgo.SavePng(img, "test_1.png")
  robotgo.Save(img, "test_1.png")

  bit2 := robotgo.ToCBitmap(robotgo.ImgToBitmap(img))
  fx, fy := bitmap.Find(bit2)
  fmt.Println("FindBitmap------ ", fx, fy)
  robotgo.Move(fx, fy)

  arr := bitmap.FindAll(bit2)
  fmt.Println("Find all bitmap: ", arr)

  fx, fy = bitmap.Find(bit)
  fmt.Println("FindBitmap------ ", fx, fy)

  bitmap.Save(bit, "test.png")
}
```

### OpenCV

``` go
package main

import (
  "fmt"
  "math/rand"

  "github.com/go-vgo/robotgo"
  "github.com/vcaesar/gcv"
  "github.com/vcaesar/bitmap"
)

func main() {
  opencv()
}

func opencv() {
  name := "test.png"
  name1 := "test_001.png"
  robotgo.SaveCapture(name1, 10, 10, 30, 30)
  robotgo.SaveCapture(name)

  fmt.Print("gcv find image: ")
  fmt.Println(gcv.FindImgFile(name1, name))
  fmt.Println(gcv.FindAllImgFile(name1, name))

  bit := bitmap.Open(name1)
  defer robotgo.FreeBitmap(bit)
  fmt.Print("find bitmap: ")
  fmt.Println(bitmap.Find(bit))

  // bit0 := robotgo.CaptureScreen()
  // img := robotgo.ToImage(bit0)
  // bit1 := robotgo.CaptureScreen(10, 10, 30, 30)
  // img1 := robotgo.ToImage(bit1)
  // defer robotgo.FreeBitmapArr(bit0, bit1)
  img, _ := robotgo.CaptureImg()
  img1, _ := robotgo.CaptureImg(10, 10, 30, 30)

  fmt.Print("gcv find image: ")
  fmt.Println(gcv.FindImg(img1, img))
  fmt.Println()

  res := gcv.FindAllImg(img1, img)
  fmt.Println(res[0].TopLeft.Y, res[0].Rects.TopLeft.X, res)
  x, y := res[0].TopLeft.X, res[0].TopLeft.Y
  robotgo.Move(x, y-rand.Intn(5))
  robotgo.MilliSleep(100)
  robotgo.Click()

  res = gcv.FindAll(img1, img) // use find template and sift
  fmt.Println("find all: ", res)
  res1 := gcv.Find(img1, img)
  fmt.Println("find: ", res1)

  img2, _, _ := robotgo.DecodeImg("test_001.png")
  x, y = gcv.FindX(img2, img)
  fmt.Println(x, y)
}
```

### Event

``` go
package main

import (
  "fmt"

  // "github.com/go-vgo/robotgo"
  hook "github.com/robotn/gohook"
)

func main() {
  add()
  low()
  event()
}

func add() {
  fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
  hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
    fmt.Println("ctrl-shift-q")
    hook.End()
  })

  fmt.Println("--- Please press w---")
  hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
    fmt.Println("w")
  })

  s := hook.Start()
  <-hook.Process(s)
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}

func event() {
  ok := hook.AddEvents("q", "ctrl", "shift")
  if ok {
    fmt.Println("add events...")
  }

  keve := hook.AddEvent("k")
  if keve {
    fmt.Println("you press... ", "k")
  }

  mleft := hook.AddEvent("mleft")
  if mleft {
    fmt.Println("you press... ", "mouse left button")
  }
}
```

### Window

``` go
package main

import (
  "fmt"

  "github.com/go-vgo/robotgo"
)

func main() {
  fpid, err := robotgo.FindIds("Google")
  if err == nil {
    fmt.Println("pids... ", fpid)

    if len(fpid) > 0 {
      robotgo.TypeStr("Hi galaxy!", fpid[0])
      robotgo.KeyTap("a", fpid[0], "cmd")

      robotgo.KeyToggle("a", fpid[0])
      robotgo.KeyToggle("a", fpid[0], "up")

      robotgo.ActivePid(fpid[0])

      robotgo.Kill(fpid[0])
    }
  }

  robotgo.ActiveName("chrome")

  isExist, err := robotgo.PidExists(100)
  if err == nil && isExist {
    fmt.Println("pid exists is", isExist)

    robotgo.Kill(100)
  }

  abool := robotgo.Alert("test", "robotgo")
  if abool {
 	  fmt.Println("ok@@@ ", "ok")
  }

  title := robotgo.GetTitle()
  fmt.Println("title@@@ ", title)
}
```

## Authors

- [The author is vz](https://github.com/vcaesar)
- [Maintainers](https://github.com/orgs/go-vgo/people)
- [Contributors](https://github.com/go-vgo/robotgo/graphs/contributors)

## Plans

- Refactor some C code to Go (such as x11, windows) 将部分 C 代码重构为 Go（例如 x11、Windows）
- Better multiscreen support 更完善的多屏支持
- Wayland support
- Update Window Handle
- Try to support Android and IOS 尝试支持 Android 和 iOS

## Contributors

- See [contributors page](https://github.com/go-vgo/robotgo/graphs/contributors) for full list of contributors.
- See [Contribution Guidelines](https://github.com/go-vgo/robotgo/blob/master/CONTRIBUTING.md).

## License

Robotgo is primarily distributed under the terms of "both the MIT license and the Apache License (Version 2.0)", with portions covered by various BSD-like licenses.

See [LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0), [LICENSE-MIT](https://github.com/go-vgo/robotgo/raw/master/LICENSE).

Collapse ▴

## Overview 

Package robotgo Go native cross-platform system automation.

​	robotgo 包是 Go 原生的跨平台系统自动化库。

Please make sure Golang, GCC is installed correctly before installing RobotGo;

​	请在安装 RobotGo 之前，确保已正确安装 Golang 和 GCC；

See Requirements:

​	请查看安装要求：

```txt
https://github.com/go-vgo/robotgo#requirements
```

Installation:

With Go module support (Go 1.11+), just import:

​	在支持 Go module（Go 1.11 及以上）的项目中，只需导入：

```go
import "github.com/go-vgo/robotgo"
```

Otherwise, to install the robotgo package, run the command:

​	否则，可运行以下命令安装：

```sh
go get -u github.com/go-vgo/robotgo
```

+bulid linux,next

+bulid windows,next



## 常量

[View Source](https://github.com/go-vgo/robotgo/blob/v0.110.7/key.go#L33)

``` go
const (
	// KeyA define key "a"
	KeyA = "a"
	KeyB = "b"
	KeyC = "c"
	KeyD = "d"
	KeyE = "e"
	KeyF = "f"
	KeyG = "g"
	KeyH = "h"
	KeyI = "i"
	KeyJ = "j"
	KeyK = "k"
	KeyL = "l"
	KeyM = "m"
	KeyN = "n"
	KeyO = "o"
	KeyP = "p"
	KeyQ = "q"
	KeyR = "r"
	KeyS = "s"
	KeyT = "t"
	KeyU = "u"
	KeyV = "v"
	KeyW = "w"
	KeyX = "x"
	KeyY = "y"
	KeyZ = "z"
	//
	CapA = "A"
	CapB = "B"
	CapC = "C"
	CapD = "D"
	CapE = "E"
	CapF = "F"
	CapG = "G"
	CapH = "H"
	CapI = "I"
	CapJ = "J"
	CapK = "K"
	CapL = "L"
	CapM = "M"
	CapN = "N"
	CapO = "O"
	CapP = "P"
	CapQ = "Q"
	CapR = "R"
	CapS = "S"
	CapT = "T"
	CapU = "U"
	CapV = "V"
	CapW = "W"
	CapX = "X"
	CapY = "Y"
	CapZ = "Z"
	//
	Key0 = "0"
	Key1 = "1"
	Key2 = "2"
	Key3 = "3"
	Key4 = "4"
	Key5 = "5"
	Key6 = "6"
	Key7 = "7"
	Key8 = "8"
	Key9 = "9"

	// Backspace backspace key string
	Backspace = "backspace"
	Delete    = "delete"
	Enter     = "enter"
	Tab       = "tab"
	Esc       = "esc"
	Escape    = "escape"
	Up        = "up"    // Up arrow key
	Down      = "down"  // Down arrow key
	Right     = "right" // Right arrow key
	Left      = "left"  // Left arrow key
	Home      = "home"
	End       = "end"
	Pageup    = "pageup"
	Pagedown  = "pagedown"

	F1  = "f1"
	F2  = "f2"
	F3  = "f3"
	F4  = "f4"
	F5  = "f5"
	F6  = "f6"
	F7  = "f7"
	F8  = "f8"
	F9  = "f9"
	F10 = "f10"
	F11 = "f11"
	F12 = "f12"
	F13 = "f13"
	F14 = "f14"
	F15 = "f15"
	F16 = "f16"
	F17 = "f17"
	F18 = "f18"
	F19 = "f19"
	F20 = "f20"
	F21 = "f21"
	F22 = "f22"
	F23 = "f23"
	F24 = "f24"

	Cmd  = "cmd"  // is the "win" key for windows
	Lcmd = "lcmd" // left command
	Rcmd = "rcmd" // right command
	// "command"
	Alt     = "alt"
	Lalt    = "lalt" // left alt
	Ralt    = "ralt" // right alt
	Ctrl    = "ctrl"
	Lctrl   = "lctrl" // left ctrl
	Rctrl   = "rctrl" // right ctrl
	Control = "control"
	Shift   = "shift"
	Lshift  = "lshift" // left shift
	Rshift  = "rshift" // right shift
	// "right_shift"
	Capslock    = "capslock"
	Space       = "space"
	Print       = "print"
	Printscreen = "printscreen" // No Mac support
	Insert      = "insert"
	Menu        = "menu" // Windows only

	AudioMute    = "audio_mute"     // Mute the volume
	AudioVolDown = "audio_vol_down" // Lower the volume
	AudioVolUp   = "audio_vol_up"   // Increase the volume
	AudioPlay    = "audio_play"
	AudioStop    = "audio_stop"
	AudioPause   = "audio_pause"
	AudioPrev    = "audio_prev"    // Previous Track
	AudioNext    = "audio_next"    // Next Track
	AudioRewind  = "audio_rewind"  // Linux only
	AudioForward = "audio_forward" // Linux only
	AudioRepeat  = "audio_repeat"  //  Linux only
	AudioRandom  = "audio_random"  //  Linux only

	Num0    = "num0" // numpad 0
	Num1    = "num1"
	Num2    = "num2"
	Num3    = "num3"
	Num4    = "num4"
	Num5    = "num5"
	Num6    = "num6"
	Num7    = "num7"
	Num8    = "num8"
	Num9    = "num9"
	NumLock = "num_lock"

	NumDecimal = "num."
	NumPlus    = "num+"
	NumMinus   = "num-"
	NumMul     = "num*"
	NumDiv     = "num/"
	NumClear   = "num_clear"
	NumEnter   = "num_enter"
	NumEqual   = "num_equal"

	LightsMonUp     = "lights_mon_up"     // Turn up monitor brightness			No Windows support
	LightsMonDown   = "lights_mon_down"   // Turn down monitor brightness		No Windows support
	LightsKbdToggle = "lights_kbd_toggle" // Toggle keyboard backlight on/off		No Windows support
	LightsKbdUp     = "lights_kbd_up"     // Turn up keyboard backlight brightness	No Windows support
	LightsKbdDown   = "lights_kbd_down"
)
```

Defining a bunch of constants.

[View Source](https://github.com/go-vgo/robotgo/blob/v0.110.7/keycode.go#L22)

``` go
const (
	// Mleft mouse left button
	Mleft      = "left"
	Mright     = "right"
	Center     = "center"
	WheelDown  = "wheelDown"
	WheelUp    = "wheelUp"
	WheelLeft  = "wheelLeft"
	WheelRight = "wheelRight"
)
```

[View Source](https://github.com/go-vgo/robotgo/blob/v0.110.7/robotgo.go#L62)

``` go
const (
	// Version get the robotgo version
	Version = "v1.00.0.1189, MT. Baker!"
)
```

## 变量

[View Source](https://github.com/go-vgo/robotgo/blob/v0.110.7/robotgo.go#L72)

``` go
var (
	// MouseSleep set the mouse default millisecond sleep time
	MouseSleep = 0
	// KeySleep set the key default millisecond sleep time
	KeySleep = 0

	// DisplayID set the screen display id
	DisplayID = -1

	// NotPid used the hwnd not pid in windows
	NotPid bool
	// Scale option the os screen scale
	Scale bool
)
```

[View Source](https://github.com/go-vgo/robotgo/blob/v0.110.7/keycode.go#L34)

``` go
var Keycode = keycode.Keycode
```

Keycode robotgo hook key's code map

[View Source](https://github.com/go-vgo/robotgo/blob/v0.110.7/keycode.go#L20)

``` go
var MouseMap = keycode.MouseMap
```

MouseMap robotgo hook mouse's code map

[View Source](https://github.com/go-vgo/robotgo/blob/v0.110.7/keycode.go#L37)

``` go
var Special = keycode.Special
```

Special is the special key map

## 函数

### func ActiveName

``` go
func ActiveName(name string) error
```

ActiveName active the window by name

​	ActiveName 根据窗口名称激活对应窗口

#### Examples:

``` go
robotgo.ActiveName("chrome")
```

### func ActivePid ← 0.110.0

``` go
func ActivePid(pid int, args ...int) error
```

ActivePid active the window by Pid, If args[0] > 0 on the Windows platform via a window handle to active, If args[0] > 0 on the unix platform via a xid to active 
 	ActivePid 根据进程 PID 激活窗口， 若 args[0] > 0，则在 Windows 上通过窗口句柄激活，在 Unix 上通过 XID 激活。

### func ActivePidC ← 0.110.0

``` go
func ActivePidC(pid int, args ...int) error
```

ActivePidC active the window by Pid, If args[0] > 0 on the unix platform via a xid to active

​	 ActivePidC 根据 PID 激活窗口，Unix 平台若 args[0] > 0 则通过 XID 激活。

### func Alert ← 0.110.0

``` go
func Alert(title, msg string, args ...string) bool
```

Alert show a alert window Displays alert with the attributes. If cancel button is not given, only the default button is displayed

​	 Alert 显示一个原生弹窗，参数可指定标题、内容及按钮；若未传入取消按钮，则仅显示默认按钮。

#### Examples

``` go
robotgo.Alert("hi", "window", "ok", "cancel")
```

### func ByteToImg ← 0.93.2

``` go
func ByteToImg(b []byte) (image.Image, error)
```

ByteToImg convert []byte to image.Image

​	ByteToImg 将字节切片转换为 image.Image。

### func Capture ← 0.110.0

``` go
func Capture(args ...int) (*image.RGBA, error)
```

Capture capture the screenshot, use the CaptureImg default

​	Capture 抓取屏幕到 RGBA，等同于默认使用 CaptureImg。

### func CaptureImg ← 0.100.1

``` go
func CaptureImg(args ...int) (image.Image, error)
```

CaptureImg capture the screen and return image.Image, error

 	CaptureImg 抓取屏幕并返回 image.Image。

### func CharCodeAt

``` go
func CharCodeAt(s string, n int) rune
```

CharCodeAt char code at utf-8

​	CharCodeAt 获取 UTF-8 字符串中第 n 个字符的码点。

### func CheckMouse

``` go
func CheckMouse(btn string) C.MMMouseButton
```

CheckMouse check the mouse button

​	CheckMouse 检测并返回指定的鼠标按键。

### func Click

``` go
func Click(args ...interface{})
```

Click click the mouse button

​	robotgo.Click(button string, double bool)

#### Examples

``` go
robotgo.Click()               // 默认左键单击  
robotgo.Click("right")        // 右键单击  
robotgo.Click("wheelLeft")    // 鼠标中键左侧点击  
```

### func CloseMainDisplay ← 0.110.3

``` go
func CloseMainDisplay()
```

CloseMainDisplay close the main X11 display

​	CloseMainDisplay 关闭主 X11 显示器连接。

### func CloseWindow

``` go
func CloseWindow(args ...int)
```

CloseWindow close the window

​	CloseWindow 关闭指定窗口。

### func CmdCtrl ← 0.110.0

``` go
func CmdCtrl() string
```

CmdCtrl If the operating system is macOS, return the key string "cmd", otherwise return the key string "ctrl"

​	CmdCtrl 返回当前系统的主控键，在 macOS 上返回 "cmd"，其它平台返回 "ctrl"。

### func DecodeImg

``` go
func DecodeImg(path string) (image.Image, string, error)
```

DecodeImg decode the image to image.Image and return

​	DecodeImg 解码图像文件为 image.Image，并返回格式信息。

### func DisplaysNum ← 0.110.0

``` go
func DisplaysNum() int
```

DisplaysNum get the count of displays

​	DisplaysNum 获取当前所有显示器的数量。

### func Drag ← DEPRECATED

``` go
func Drag(x, y int, args ...string)
```

Deprecated: use the DragSmooth(),

Drag drag the mouse to (x, y), It's not valid now, use the DragSmooth()

​	已弃用：请使用 DragSmooth()；

​	Drag 在屏幕上拖动鼠标到指定坐标（现已失效）。

### func DragMouse ← DEPRECATED

``` go
func DragMouse(x, y int, args ...interface{})
```

Deprecated: use the DragSmooth(),

DragMouse drag the mouse to (x, y), It's same with the DragSmooth() now

​	已弃用：请使用 DragSmooth()；

​	DragMouse 功能同 DragSmooth()。

### func DragSmooth

``` go
func DragSmooth(x, y int, args ...interface{})
```

DragSmooth drag the mouse like smooth to (x, y)

​	DragSmooth 以平滑轨迹拖动鼠标到指定坐标。

#### Examples

``` go
robotgo.DragSmooth(10, 10)
```

### func FindIds

``` go
func FindIds(name string) ([]int, error)
```

FindIds finds the all processes named with a subset of "name" (case insensitive), return matched IDs.

​	FindIds 查找所有进程名称中包含指定子串（不区分大小写）的进程，返回匹配的 PID 列表。

### func FindName

``` go
func FindName(pid int) (string, error)
```

FindName find the process name by the process id

​	FindName 根据 PID 查找进程名称

------

### func FindNames

``` go
func FindNames() ([]string, error)
```

FindNames find the all process name

​	FindNames 查找所有进程名称

------

### func FindPath

``` go
func FindPath(pid int) (string, error)
```

FindPath find the process path by the process pid

​	FindPath 根据 PID 查找进程可执行文件路径

------

### func FreeBitmap

``` go
func FreeBitmap(bitmap CBitmap)
```

FreeBitmap free and dealloc the C bitmap

FreeBitmap 释放并回收 CBitmap

------

### func FreeBitmapArr ← 0.100.1

``` go
func FreeBitmapArr(bit ...CBitmap)
```

FreeBitmapArr free and dealloc the C bitmap array

​	FreeBitmapArr 释放并回收多个 CBitmap

------

### func GetActiveC ← 0.110.3

``` go
func GetActiveC() C.MData
```

GetActiveC get the active window

​	GetActiveC 获取当前激活窗口信息

------

### func GetBHandle ← DEPRECATED

``` go
func GetBHandle() int
```

Deprecated: use the GetHandle(),

 GetBHandle get the window handle, Wno-deprecated

This function will be removed in version v1.0.0

​	已弃用：请使用 GetHandle()

​	GetBHandle 返回窗口句柄（将于 v1.0.0 移除）

------

### func GetBounds

``` go
func GetBounds(pid int, args ...int) (int, int, int, int)
```

GetBounds get the window bounds

​	GetBounds 获取指定 PID 窗口的边界 (x, y, w, h)

------

### func GetClient ← 0.110.0

``` go
func GetClient(pid int, args ...int) (int, int, int, int)
```

GetClient get the window client bounds

GetClient 获取指定 PID 窗口的客户区边界 (x, y, w, h)

------

### func GetDisplayBounds ← 0.110.0

``` go
func GetDisplayBounds(i int) (x, y, w, h int)
```

GetDisplayBounds gets the display screen bounds

​	GetDisplayBounds 获取第 i 个显示器的屏幕边界 (x, y, w, h)

------

### func GetHWNDByPid ← 0.110.0

``` go
func GetHWNDByPid(pid int) int
```

GetHWNDByPid get the hwnd by pid

GetHWNDByPid 根据 PID 返回 Windows 窗口句柄

------

### func GetHandByPidC ← 0.110.3

``` go
func GetHandByPidC(pid int, args ...int) C.MData
```

GetHandByPidC get handle mdata by pid

​	GetHandByPidC 根据 PID 返回跨平台窗口句柄数据

------

### func GetHandle

``` go
func GetHandle() int
```

GetHandle get the window handle

​	GetHandle 获取当前活动窗口的句柄

------

### func GetLocationColor ← 0.110.0

``` go
func GetLocationColor(displayId ...int) string
```

GetLocationColor get the location pos's color

​	GetLocationColor 获取当前鼠标位置或指定显示器上某点的像素颜色

------

### func GetMainId ← 0.110.0

``` go
func GetMainId() int
```

GetMainId get the main display id

GetMainId 获取主显示器 ID

------

### func GetMousePos ← DEPRECATED

``` go
func GetMousePos() (int, int)
```

Deprecated: use the function Location()

GetMousePos get the mouse's position return x, y

​	已弃用：请使用 Location()
​	 GetMousePos 获取鼠标当前位置 (x, y)

------

### func GetPid ← 0.110.0

``` go
func GetPid() int
```

GetPid get the process id return int32

​	GetPid 获取当前进程的 PID

------

### func GetPixelColor

``` go
func GetPixelColor(x, y int, displayId ...int) string
```

GetPixelColor get the pixel color return string

​	GetPixelColor 获取指定坐标的像素颜色（十六进制字符串）

------

### func GetPxColor

``` go
func GetPxColor(x, y int, displayId ...int) C.MMRGBHex
```

GetPxColor get the pixel color return C.MMRGBHex

​	GetPxColor 获取指定坐标的像素颜色（C.MMRGBHex 结构）

------

### func GetScaleSize

``` go
func GetScaleSize(displayId ...int) (int, int)
```

GetScaleSize get the screen scale size

​	GetScaleSize 获取指定显示器的缩放后尺寸 (宽, 高)

------

### func GetScreenSize

``` go
func GetScreenSize() (int, int)
```

GetScreenSize get the screen size

​	GetScreenSize 获取屏幕分辨率 (宽, 高)

------

### func GetText

``` go
func GetText(imgPath string, args ...string) (string, error)
```

GetText get the image text by tesseract ocr

 robotgo.GetText(imgPath, lang string)

​	GetText 使用 Tesseract OCR 识别图像中的文字

 robotgo.GetText(imgPath, lang string)

------

### func GetTitle

``` go
func GetTitle(args ...int) string
```

GetTitle get the window title return string

​	GetTitle 获取指定窗口或当前活动窗口的标题

#### Examples:

``` go
fmt.Println(robotgo.GetTitle())

ids, _ := robotgo.FindIds()
robotgo.GetTitle(ids[0])
```



------

### func GetVersion

``` go
func GetVersion() string
```

GetVersion get the robotgo version

​	GetVersion 获取 RobotGo 的版本号

------

### func GetXDisplayName

``` go
func GetXDisplayName() string
```

GetXDisplayName get XDisplay name (Linux)

​	GetXDisplayName 获取 X11 系统的 DISPLAY 名称

------

### func GetXid ← 0.110.0

``` go
func GetXid(xu *xgbutil.XUtil, pid int) (xproto.Window, error)
```

GetXid get the xid return window and error

​	GetXid 通过 PID 获取 X11 窗口的 XID

------

### func GetXidFromPid

``` go
func GetXidFromPid(xu *xgbutil.XUtil, pid int) (xproto.Window, error)
```

GetXidFromPid get the xid from pid

​	GetXidFromPid 同上，通过 PID 获取 XID

------

### func GoString

``` go
func GoString(char *C.char) string
```

GoString trans C.char to string

​	GoString 将 C.char 转换为 Go 字符串

------

### func Height ← 0.100.2

``` go
func Height(img image.Image) int
```

Height return the image.Image height

​	Height 返回 image.Image 的高度

------

### func HexToRgb

``` go
func HexToRgb(hex uint32) *C.uint8_t
```

HexToRgb trans hex to rgb

​	HexToRgb 将十六进制颜色值转换为 RGB

------

### func ImgSize ← 0.100.2

``` go
func ImgSize(path string) (int, int, error)
```

ImgSize get the file image size

​	ImgSize 获取图像文件的尺寸 (宽, 高)

------

### func Is64Bit

``` go
func Is64Bit() bool
```

Is64Bit determine whether the sys is 64bit

Is64Bit 判断当前系统是否为 64 位

------

### func IsMain ← 0.110.0

``` go
func IsMain(displayId int) bool
```

IsMain is main display

​	IsMain 判断指定显示器是否为主显示器

------

### func IsValid

``` go
func IsValid() bool
```

IsValid valid the window

​	IsValid 判断当前活动窗口是否有效

------

### func KeyDown ← 0.100.0

``` go
func KeyDown(key string, args ...interface{}) error
```

KeyDown press down a key

​	KeyDown 模拟按下指定按键

------

### func KeyPress ← 0.100.0

``` go
func KeyPress(key string, args ...interface{}) error
```

KeyPress press key string

​	KeyPress 模拟按键按下并释放

------

### func KeyTap

``` go
func KeyTap(key string, args ...interface{}) error
```

KeyTap taps the keyboard code;

​	KeyTap 模拟单击键盘按键；

​	 See keys supported: 支持按键列表见：

[https://github.com/go-vgo/robotgo/blob/master/docs/keys.md#keys](https://github.com/go-vgo/robotgo/blob/master/docs/keys.md#keys)

#### Examples:

``` go
robotgo.KeySleep = 100 // 100 millisecond
robotgo.KeyTap("a")
robotgo.KeyTap("i", "alt", "command")

arr := []string{"alt", "command"}
robotgo.KeyTap("i", arr)

robotgo.KeyTap("k", pid int)
```

------

### func KeyToggle

``` go
func KeyToggle(key string, args ...interface{}) error
```

KeyToggle toggles the keyboard, if there not have args default is "down"

​	KeyToggle 切换按键状态，若未指定状态则默认为 "down"	

See keys:
 [https://github.com/go-vgo/robotgo/blob/master/docs/keys.md#keys]( https://github.com/go-vgo/robotgo/blob/master/docs/keys.md#keys)

Examples:

``` go
robotgo.KeyToggle("a")
robotgo.KeyToggle("a", "up")

robotgo.KeyToggle("a", "up", "alt", "cmd")
robotgo.KeyToggle("k", pid int)
```



------

### func KeyUp ← 0.100.0

``` go
func KeyUp(key string, args ...interface{}) error
```

KeyUp press up a key

​	KeyUp 模拟释放指定按键

### func Kill

``` go
func Kill(pid int) error
```

Kill kill the process by PID

​	Kill 通过 PID 杀死进程

------

### func Location ← 0.110.0

``` go
func Location() (int, int)
```

Location get the mouse location position return x, y

​	Location 获取鼠标当前位置，返回 x, y

------

### func MaxWindow

``` go
func MaxWindow(pid int, args ...interface{})
```

MaxWindow set the window max

​	MaxWindow 将指定窗口最大化

------

### func MicroSleep DEPRECATED

``` go
func MicroSleep(tm float64)
```

Deprecated: use the MilliSleep(),

MicroSleep time C.microsleep(tm)

​	已弃用：请使用 MilliSleep()，

------

### func MilliSleep

``` go
func MilliSleep(tm int)
```

MilliSleep sleep tm milli second

​	MilliSleep 以毫秒为单位暂停 tm 毫秒

------

### func MinWindow

``` go
func MinWindow(pid int, args ...interface{})
```

MinWindow set the window min

​	MinWindow 将指定窗口最小化

------

### func MouseClick ← DEPRECATED

``` go
func MouseClick(args ...interface{})
```

Deprecated: use the Click(),

MouseClick click the mouse

​	已弃用：请使用 Click()，

​	MouseClick 单击鼠标

------

### func MouseDown ← 0.110.0

``` go
func MouseDown(key ...interface{}) error
```

MouseDown send mouse down event

​	MouseDown 发送鼠标按下事件

------

### func MouseUp ← 0.110.0

``` go
func MouseUp(key ...interface{}) error
```

MouseUp send mouse up event

​	MouseUp 发送鼠标抬起事件

------

### func Move

``` go
func Move(x, y int, displayId ...int)
```

Move move the mouse to (x, y)

​	Move 将鼠标移动到 (x, y)

#### Examples:

``` go
robotgo.MouseSleep = 100  // 100 millisecond  
robotgo.Move(10, 10)
```

------

### func MoveArgs

``` go
func MoveArgs(x, y int) (int, int)
```

MoveArgs get the mouse relative args

​	MoveArgs 获取鼠标移动的相对坐标

------

### func MoveClick

``` go
func MoveClick(x, y int, args ...interface{})
```

MoveClick move and click the mouse

​	MoveClick 移动鼠标并单击

robotgo.MoveClick(x, y int, button string, double bool)

#### Examples:

``` go
robotgo.MouseSleep = 100  
robotgo.MoveClick(10, 10)
```

------

### func MoveMouse ← DEPRECATED

``` go
func MoveMouse(x, y int)
```

Deprecated: use the Move(),

MoveMouse move the mouse

​	已弃用：请使用 Move()，

 	MoveMouse 将鼠标移动到指定位置

------

### func MoveMouseSmooth ← DEPRECATED

``` go
func MoveMouseSmooth(x, y int, args ...interface{}) bool
```

Deprecated: use the MoveSmooth(),

MoveMouseSmooth move the mouse smooth, moves mouse to x, y human like, with the mouse button up.

​	已弃用：请使用 MoveSmooth()，

​	MoveMouseSmooth 平滑移动鼠标到 (x, y)

------

### func MoveRelative

``` go
func MoveRelative(x, y int)
```

MoveRelative move mouse with relative

​	MoveRelative 相对移动鼠标

------

### func MoveScale ← 0.110.0

``` go
func MoveScale(x, y int, displayId ...int) (int, int)
```

MoveScale calculate the os scale factor x, y

​	MoveScale 根据系统缩放计算坐标

------

### func MoveSmooth

``` go
func MoveSmooth(x, y int, args ...interface{}) bool
```

MoveSmooth move the mouse smooth, moves mouse to x, y human like, with the mouse button up.

​	 MoveSmooth 平滑移动鼠标到 (x, y)

 robotgo.MoveSmooth(x, y int, low, high float64, mouseDelay int)

#### Examples:

``` go
robotgo.MoveSmooth(10, 10)  
robotgo.MoveSmooth(10, 10, 1.0, 2.0)
```



------

### func MoveSmoothRelative

``` go
func MoveSmoothRelative(x, y int, args ...interface{})
```

MoveSmoothRelative move mouse smooth with relative

​	MoveSmoothRelative 相对平滑移动鼠标

------

### func MovesClick

``` go
func MovesClick(x, y int, args ...interface{})
```

MovesClick move smooth and click the mouse

use the `robotgo.MouseSleep = 100`

​	MovesClick 平滑移动并单击鼠标

​	使用 `robotgo.MouseSleep = 100`

------

### func Mul ← DEPRECATED

``` go
func Mul(x int) int
```

Deprecated: use the ScaledF(),

Mul mul the scale, drop

​	已弃用：请使用 ScaledF()，

​	Mul 按系统缩放计算，已废弃

------

### func OpenImg

``` go
func OpenImg(path string) ([]byte, error)
```

OpenImg open the image return []byte

​	OpenImg 打开图像并返回`[]byte`

------

### func PadHex

``` go
func PadHex(hex C.MMRGBHex) string
```

PadHex trans C.MMRGBHex to string

​	PadHex 将 C.MMRGBHex 转换为字符串

------

### func PadHexs ← 0.110.0

``` go
func PadHexs(hex CHex) string
```

PadHexs trans CHex to string

​	PadHexs 将 CHex 转换为字符串

------

### func PasteStr

``` go
func PasteStr(str string) error
```

PasteStr paste a string (support UTF-8), write the string to clipboard and tap `cmd + v`

​	PasteStr 粘贴字符串（支持 UTF-8），写入剪贴板并模拟 `cmd+v`

------

### func PidExists

``` go
func PidExists(pid int) (bool, error)
```

PidExists determine whether the process exists

​	PidExists 判断指定 PID 的进程是否存在

------

### func Pids

``` go
func Pids() ([]int, error)
```

Pids get the all process id

​	Pids 获取所有进程的 PID 列表

------

### func Read ← 0.100.0

``` go
func Read(path string) (image.Image, error)
```

Read read the file return image.Image

​	Read 读取图像文件并返回 image.Image

------

### func ReadAll

``` go
func ReadAll() (string, error)
```

ReadAll read string from clipboard

​	ReadAll 从剪贴板读取文本

------

### func RgbToHex

``` go
func RgbToHex(r, g, b uint8) C.uint32_t
```

RgbToHex trans rgb to hex

​	RgbToHex 将 RGB 值转换为十六进制

------

### func Run ← 0.100.9

``` go
func Run(path string) ([]byte, error)
```

Run run a cmd shell

​	Run 在系统 shell 中执行命令并返回输出

------

### func Save ← 0.100.0

``` go
func Save(img image.Image, path string, quality ...int) error
```

Save create a image file with the image.Image

​	Save 使用 image.Image 创建图像文件，可指定质量

------

### func SaveCapture

``` go
func SaveCapture(path string, args ...int) error
```

SaveCapture capture screen and save the screenshot to image

​	SaveCapture 捕获屏幕并保存为图像文件

------

### func SaveImg

``` go
func SaveImg(b []byte, path string) error
```

SaveImg save the image by []byte
 SaveImg 将字节切片保存为图像文件

------

### func SaveJpeg ← 0.100.0

``` go
func SaveJpeg(img image.Image, path string, quality ...int) error
```

SaveJpeg save the image by image.Image

​	SaveJpeg 以 JPEG 格式保存 image.Image

------

### func SavePng ← 0.92.0

``` go
func SavePng(img image.Image, path string) error
```

SavePng save the image by image.Image

​	SavePng 以 PNG 格式保存 image.Image

------

### func Scale0 DEPRECATED

``` go
func Scale0() int
```

Deprecated: use the ScaledF(), Scale0 return ScaleX() / 0.96, drop

​	已弃用：请使用 ScaledF()，

​	 Scale0 返回 ScaleX()/0.96，已废弃

------

### func Scale1 DEPRECATED

``` go
func Scale1() int
```

Deprecated: use the ScaledF(),
 Scale1 get the screen scale (only windows old), drop
 已弃用：请使用 ScaledF()，
 Scale1 获取旧版 Windows 缩放值，已废弃

### func ScaleF ← 0.100.8

``` go
func ScaleF(displayId ...int) float64
```

ScaleF get the system scale val

​	ScaleF 获取系统缩放值

------

### func ScaleX ← DEPRECATED

``` go
func ScaleX() int
```

Deprecated: use the ScaledF(),

ScaleX get the primary display horizontal DPI scale factor, drop

​	已弃用：请使用 ScaledF()， ScaleX 获取主显示器水平 DPI 缩放因子，已废弃

------

### func Scaled

``` go
func Scaled(x int, displayId ...int) int
```

Scaled get the screen scaled return scale size

​	Scaled 根据系统缩放计算并返回缩放后的尺寸

------

### func Scaled0 ← 0.100.8

``` go
func Scaled0(x int, f float64) int
```

Scaled0 return int(x * f)

​	Scaled0 返回 int(x * f)

------

### func Scaled1 ← 0.110.0

``` go
func Scaled1(x int, f float64) int
```

Scaled1 return int(x / f)

​	Scaled1 返回 int(x / f)

------

### func Scroll

``` go
func Scroll(x, y int, args ...int)
```

Scroll scroll the mouse to (x, y)

​	Scroll 将鼠标滚动到 (x, y)

robotgo.Scroll(x, y, msDelay int)

#### Examples:

``` go
robotgo.Scroll(10, 10)
```

------

### func ScrollDir ← 0.110.0

``` go
func ScrollDir(x int, direction ...interface{})
```

ScrollDir scroll the mouse with direction to (x, "up") supported: "up", "down", "left", "right"

​	ScrollDir 按方向滚动鼠标，支持："up"、"down"、"left"、"right"

#### Examples:

``` go
robotgo.ScrollDir(10, "down")
robotgo.ScrollDir(10, "up")
```

------

### func ScrollRelative ← 0.100.1

``` go
func ScrollRelative(x, y int, args ...int)
```

ScrollRelative scroll mouse with relative

​	ScrollRelative 相对滚动鼠标

#### Examples:

``` go
robotgo.ScrollRelative(10, 10)
```

------

### func ScrollSmooth ← 0.100.7

``` go
func ScrollSmooth(to int, args ...int)
```

ScrollSmooth scroll the mouse smooth, default scroll 5 times and sleep 100 millisecond

​	ScrollSmooth 平滑滚动鼠标，默认滚动 5 次并暂停 100 毫秒

 robotgo.ScrollSmooth(toy, num, sleep, tox)

#### Examples:

``` go
robotgo.ScrollSmooth(-10)
robotgo.ScrollSmooth(-10, 6, 200, -10)
```

------

### func SetActive

``` go
func SetActive(win Handle)
```

SetActive set the window active

​	SetActive 激活指定窗口

------

### func SetActiveC ← 0.110.3

``` go
func SetActiveC(win C.MData)
```

SetActiveC set the window active

​	SetActiveC 激活指定窗口（跨平台句柄）

------

### func SetDelay

``` go
func SetDelay(d ...int)
```

SetDelay sets the key and mouse delay

robotgo.SetDelay(100) option the robotgo.KeySleep and robotgo.MouseSleep = d

​	SetDelay 设置按键与鼠标延迟

​	robotgo.SetDelay(延迟毫秒) 同时设置 robotgo.KeySleep 与 robotgo.MouseSleep

------

### func SetHandle

``` go
func SetHandle(hwnd int)
```

SetHandle set the window handle

​	SetHandle 设置当前操作的窗口句柄

------

### func SetHandlePid

``` go
func SetHandlePid(pid int, args ...int)
```

SetHandlePid set the window handle by pid

​	SetHandlePid 根据 PID 设置窗口句柄

------

### func SetXDisplayName

``` go
func SetXDisplayName(name string) error
```

SetXDisplayName set XDisplay name (Linux)

​	SetXDisplayName 设置 X11 DISPLAY 名称（Linux）

------

### func Sleep

``` go
func Sleep(tm int)
```

Sleep time.Sleep tm second

​	Sleep 调用 time.Sleep 暂停 tm 秒

------

### func StrToImg ← 0.94.0

``` go
func StrToImg(data string) (image.Image, error)
```

StrToImg convert base64 string to image.Image

​	StrToImg 将 Base64 编码字符串转换为 image.Image

------

### func SysScale

``` go
func SysScale(displayId ...int) float64
```

SysScale get the sys scale

​	SysScale 获取系统缩放因子

------

### func ToByteImg ← 0.94.0

``` go
func ToByteImg(img image.Image, fm ...string) []byte
```

ToByteImg convert image.Image to []byte

​	ToByteImg 将 image.Image 转换为字节切片

------

### func ToImage ← 0.92.0

``` go
func ToImage(bit CBitmap) image.Image
```

ToImage convert C.MMBitmapRef to standard image.Image

​	ToImage 将 C.MMBitmapRef 转换为标准 image.Image

------

### func ToInterfaces ← 0.110.0

``` go
func ToInterfaces(fields []string) []interface{}
```

ToInterfaces convert []string to []interface{}
 ToInterfaces 将 []string 转换为 []interface{}

------

### func ToMMBitmapRef

``` go
func ToMMBitmapRef(bit CBitmap) C.MMBitmapRef
```

ToMMBitmapRef trans CBitmap to C.MMBitmapRef

​	ToMMBitmapRef 将 CBitmap 转换为 C.MMBitmapRef

------

### func ToMMRGBHex

``` go
func ToMMRGBHex(hex CHex) C.MMRGBHex
```

ToMMRGBHex trans CHex to C.MMRGBHex
 ToMMRGBHex 将 CHex 转换为 C.MMRGBHex

------

### func ToRGBA ← 0.99.9

``` go
func ToRGBA(bit CBitmap) *image.RGBA
```

ToRGBA convert C.MMBitmapRef to standard image.RGBA

​	ToRGBA 将 C.MMBitmapRef 转换为标准 *image.RGBA

------

### func ToRGBAGo ← 0.99.9

``` go
func ToRGBAGo(bmp1 Bitmap) *image.RGBA
```

ToRGBAGo convert Bitmap to standard image.RGBA
 ToRGBAGo 将 Bitmap 转换为标准 *image.RGBA

------

### func ToStringImg ← 0.94.0

``` go
func ToStringImg(img image.Image, fm ...string) string
```

ToStringImg convert image.Image to string

​	ToStringImg 将 image.Image 转换为字符串（Base64 等格式）

------

### func ToStrings ← 0.110.0

``` go
func ToStrings(fields []interface{}) []string
```

ToStrings convert []interface{} to []string

​	ToStrings 将 []interface{} 转换为 []string

------

### func ToUC ← 0.99.1

``` go
func ToUC(text string) []string
```

ToUC trans string to unicode []string

​	ToUC 将字符串转换为 Unicode 码点切片

------

### func ToUint8p ← 0.99.9

``` go
func ToUint8p(dst []uint8) *uint8
```

ToUint8p convert the []uint8 to uint8 pointer

​	ToUint8p 将 []uint8 转换为 `*uint8`

------

### func Toggle ← 0.100.7

``` go
func Toggle(key ...interface{}) error
```

Toggle toggle the mouse, support button:

​	Toggle 切换鼠标按键状态，支持按钮：

```txt
 "left", "center", "right", "wheelDown", "wheelUp", "wheelLeft", "wheelRight"
```

####  Examples:

```go
 robotgo.Toggle("left") // default is down
 robotgo.Toggle("left", "up")
```



------

### func Try

``` go
func Try(fun func(), handler func(interface{}))
```

Try handler(err)

​	Try 运行 fun 并在发生 panic 时调用 handler(err)

------

### func TypeStr

``` go
func TypeStr(str string, args ...int)
```

TypeStr send a string (supported UTF-8)

robotgo.TypeStr(string: "The string to send", int: pid, "milli_sleep time", "x11 option")

#### Examples:

```go
 robotgo.TypeStr("abc@123, Hi galaxy, こんにちは")
 robotgo.TypeStr("To be or not to be, this is questions.", pid int)
```

 

------

### func TypeStrDelay

``` go
func TypeStrDelay(str string, delay int)
```

TypeStrDelay type string with delayed

 And you can use robotgo.KeySleep = 100 to delayed not this function

​	TypeStrDelay 带延迟地输入字符串

​	可使用 robotgo.KeySleep = 100 代替此函数进行延迟

------

### func TypeStringDelayed ← DEPRECATED

``` go
func TypeStringDelayed(str string, delay int)
```

Deprecated: use the TypeStr(),

 TypeStringDelayed type string delayed, Wno-deprecated

 This function will be removed in version v1.0.0

​	已弃用：请使用 TypeStr()，

​	TypeStringDelayed 延迟输入字符串，已弃用

​	此函数将在 v1.0.0 移除

------

### func U32ToHex

``` go
func U32ToHex(hex C.uint32_t) C.MMRGBHex
```

U32ToHex trans C.uint32_t to C.MMRGBHex

​	U32ToHex 将 C.uint32_t 转换为 C.MMRGBHex

------

### func U8ToHex

``` go
func U8ToHex(hex *C.uint8_t) C.MMRGBHex
```

U8ToHex trans `*C.uint8_t` to C.MMRGBHex

​	U8ToHex 将 `*C.uint8_t` 转换为 C.MMRGBHex

------

### func UnicodeType

``` go
func UnicodeType(str uint32, args ...int)
```

UnicodeType tap the uint32 unicode

​	UnicodeType 输入单个 Unicode 码点

------

### func Width ← 0.100.2

``` go
func Width(img image.Image) int
```

Width return the image.Image width

​	Width 返回 image.Image 的宽度

------

### func WriteAll

``` go
func WriteAll(text string) error
```

WriteAll write string to clipboard

​	WriteAll 将字符串写入剪贴板

------

## Types

### type Bitmap

``` go
type Bitmap struct {
    ImgBuf        *uint8
    Width, Height int

    Bytewidth     int
    BitsPixel     uint8
    BytesPerPixel uint8
}
```

Bitmap define the go Bitmap struct

​	Bitmap 定义 Go 端的 Bitmap 结构体

------

#### func CaptureGo ← 0.110.0

``` go
func CaptureGo(args ...int) Bitmap
```

CaptureGo capture the screen and return bitmap(go struct)

​	CaptureGo 抓取屏幕并返回 Go 结构体 Bitmap

------

#### func ImgToBitmap ← 0.99.9

``` go
func ImgToBitmap(m image.Image) (bit Bitmap)
```

ImgToBitmap convert the standard image.Image to Bitmap

​	ImgToBitmap 将标准 image.Image 转换为 Bitmap 结构体

------

#### func RGBAToBitmap ← 0.99.9

``` go
func RGBAToBitmap(r1 *image.RGBA) (bit Bitmap)
```

RGBAToBitmap convert the standard image.RGBA to Bitmap

​	RGBAToBitmap 将标准 *image.RGBA 转换为 Bitmap 结构体

------

#### func ToBitmap

``` go
func ToBitmap(bit CBitmap) Bitmap
```

ToBitmap trans C.MMBitmapRef to Bitmap

​	ToBitmap 将 C.MMBitmapRef 转换为 Bitmap

------

### type CBitmap

``` go
type CBitmap C.MMBitmapRef
```

CBitmap define CBitmap as C.MMBitmapRef type

​	CBitmap 定义为 C.MMBitmapRef 类型

------

#### func ByteToCBitmap ← 0.110.0

``` go
func ByteToCBitmap(by []byte) CBitmap
```

ByteToCBitmap trans []byte to CBitmap

​	ByteToCBitmap 将 []byte 转换为 CBitmap

------

#### func CaptureScreen

``` go
func CaptureScreen(args ...int) CBitmap
```

CaptureScreen capture the screen return bitmap(c struct), use `defer robotgo.FreeBitmap(bitmap)` to free the bitmap

robotgo.CaptureScreen(x, y, w, h int)

​	CaptureScreen 抓取屏幕并返回 CBitmap，需要使用 `defer robotgo.FreeBitmap(bitmap)` 释放

------

#### func ImgToCBitmap ← 0.110.0

``` go
func ImgToCBitmap(img image.Image) CBitmap
```

ImgToCBitmap trans image.Image to CBitmap

​	ImgToCBitmap 将 image.Image 转换为 CBitmap

------

#### func ToCBitmap

``` go
func ToCBitmap(bit Bitmap) CBitmap
```

ToCBitmap trans Bitmap to C.MMBitmapRef

​	ToCBitmap 将 Bitmap 转换为 C.MMBitmapRef

------

### type CHex

``` go
type CHex C.MMRGBHex
```

CHex define CHex as c rgb Hex type (C.MMRGBHex)

​	CHex 定义为 C.MMRGBHex 类型

------

#### func UintToHex

``` go
func UintToHex(u uint32) CHex
```

UintToHex trans uint32 to robotgo.CHex

​	UintToHex 将 uint32 转换为 CHex

------

### type Handle ← 0.110.3

``` go
type Handle C.MData
```

Handle define window Handle as C.MData type

​	Handle 定义为 C.MData 类型，表示窗口句柄

------

#### func GetActive

``` go
func GetActive() Handle
```

GetActive get the active window

​	GetActive 获取当前激活窗口句柄

------

#### func GetHandById ← 0.110.3

``` go
func GetHandById(id int, args ...int) Handle
```

GetHandById get handle mdata by id

​	GetHandById 根据窗口 ID 获取 Handle

------

#### func GetHandByPid ← 0.110.3

``` go
func GetHandByPid(pid int, args ...int) Handle
```

GetHandByPid get handle mdata by pid

​	GetHandByPid 根据 PID 获取 Handle

------

#### func GetHandPid ← DEPRECATED

``` go
func GetHandPid(pid int, args ...int) Handle
```

Deprecated: use the GetHandByPid(),

GetHandPid get handle mdata by pid

​	已弃用：请使用 GetHandByPid()，

​	GetHandPid 根据 PID 获取 Handle

------

### type Map

``` go
type Map map[string]interface{}
```

Map a `map[string]interface{}`

​	Map 定义为 `map[string]interface{}`

------

### type Nps

``` go
type Nps struct {
    Pid  int
    Name string
}
```

Nps process struct

​	Nps 表示进程信息的结构体

------

#### func Process

``` go
func Process() ([]Nps, error)
```

Process get the all process struct

​	Process 获取所有进程信息列表

------

### type Point ← 0.100.3

``` go
type Point struct {
    X int
    Y int
}
```

Point is point struct

​	Point 表示坐标点的结构体

------

### type Rect ← 0.100.3

``` go
type Rect struct {
    Point
    Size
}
```

Rect is rect structure

​	Rect 表示矩形区域的结构体，包含 Point 和 Size

------

#### func GetDisplayRect ← 0.110.0

``` go
func GetDisplayRect(i int) Rect
```

GetDisplayRect gets the display rect

​	GetDisplayRect 获取第 i 个显示器的矩形区域

------

#### func GetScreenRect ← 0.100.3

``` go
func GetScreenRect(displayId ...int) Rect
```

GetScreenRect get the screen rect (x, y, w, h)

​	GetScreenRect 获取屏幕矩形区域 (x, y, w, h)

------

### type Size ← 0.100.3

```
type Size struct {
    W, H int
}
```

Size is size structure

​	Size 表示尺寸的结构体，包含宽 W 和高 H
