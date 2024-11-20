+++
title = "launcher"
date = 2024-11-20T18:01:46+08:00
weight = 30
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/go-rod/rod/lib/launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher)
>
> 收录该文档时间：`2024-11-20T18:02:07+08:00`
>
> [Version: v0.116.2](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher?tab=versions)

A lib helps to find, launch or download the browser. You can also use it as a standalone lib without Rod.

​	一个库，用于查找、启动或下载浏览器。也可以作为独立库使用，而无需依赖 Rod。

Package launcher for launching browser utils.

​	用于启动浏览器工具的 `launcher` 包。

## Example (Custom_launch)

``` go
package main

import (
	"os/exec"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/leakless"
)

func main() {
	// get the browser executable path
    // 获取浏览器可执行文件路径
	path := launcher.NewBrowser().MustGet()

	// use the FormatArgs to construct args, this line is optional, you can construct the args manually
    // 使用 FormatArgs 构建参数（可选，可以手动构建参数）
	args := launcher.New().FormatArgs()

	var cmd *exec.Cmd
	if true { // decide whether to use leakless or not 决定是否使用 leakless
		cmd = leakless.New().Command(path, args...)
	} else {
		cmd = exec.Command(path, args...)
	}

	parser := launcher.NewURLParser()
	cmd.Stderr = parser
	utils.E(cmd.Start())
	u := launcher.MustResolveURL(<-parser.URL)

	rod.New().ControlURL(u).MustConnect()
}
Output:
```
## Example (Print_browser_CLI_output)

``` go
package main

import (
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Pipe the browser stderr and stdout to os.Stdout .
    // 将浏览器的 stderr 和 stdout 输出到 os.Stdout。
	u := launcher.New().Logger(os.Stdout).MustLaunch()
	rod.New().ControlURL(u).MustConnect()
}
Output:
```
## Example (Use_system_browser)

``` go
package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	if path, exists := launcher.LookPath(); exists {
		u := launcher.New().Bin(path).MustLaunch()
		rod.New().ControlURL(u).MustConnect()
	}
}
Output:
```
## 常量

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/manager.go#L17)

``` go
const (
	// HeaderName for remote launch.
    // HeaderName 用于远程启动。
	HeaderName = "Rod-Launcher"
)
```

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/revision.go#L6)

``` go
const RevisionDefault = 1321438
```

RevisionDefault for chromium.

​	RevisionDefault 为 Chromium。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/revision.go#L9)

``` go
const RevisionPlaywright = 1124
```

RevisionPlaywright for arm linux.

​	RevisionPlaywright 用于 ARM Linux。

## 变量 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/browser.go#L70)

``` go
var DefaultBrowserDir = filepath.Join(map[string]string{
	"windows": os.Getenv("APPDATA"),
	"darwin":  filepath.Join(os.Getenv("HOME"), ".cache"),
	"linux":   filepath.Join(os.Getenv("HOME"), ".cache"),
}[runtime.GOOS], "rod", "browser")
```

DefaultBrowserDir for downloaded browser. For unix is "$HOME/.cache/rod/browser", for Windows it's "%APPDATA%\rod\browser".

​	DefaultBrowserDir 为已下载的浏览器路径。Unix 系统路径为 `$HOME/.cache/rod/browser`，Windows 系统路径为 `%APPDATA%\rod\browser`。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/launcher.go#L24)

``` go
var DefaultUserDataDirPrefix = filepath.Join(os.TempDir(), "rod", "user-data")
```

DefaultUserDataDirPrefix ...

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/error.go#L6)

``` go
var ErrAlreadyLaunched = errors.New("already launched")
```

ErrAlreadyLaunched is an error that indicates the launcher has already been launched.

​	ErrAlreadyLaunched 表示启动器已被启动的错误。

## 函数 

## func HostGoogle 

``` go
func HostGoogle(revision int) string
```

HostGoogle to download browser.

​	HostGoogle 用于已下载浏览器。

## func HostNPM <- 0.102.0

``` go
func HostNPM(revision int) string
```

HostNPM to download browser.

​	HostNPM 用于已下载浏览器。

## func HostPlaywright <- 0.106.3

``` go
func HostPlaywright(revision int) string
```

HostPlaywright to download browser.

​	HostPlaywright 用于已下载浏览器。

## func LookPath <- 0.91.0

``` go
func LookPath() (found string, has bool)
```

LookPath searches for the browser executable from often used paths on current operating system.

​	LookPath 在当前操作系统的常用路径中搜索浏览器可执行文件。

## func MustResolveURL <- 0.65.0

``` go
func MustResolveURL(u string) string
```

MustResolveURL is similar to ResolveURL.

​	MustResolveURL 类似于 ResolveURL。

## func Open <- 0.91.0

``` go
func Open(url string)
```

Open tries to open the url via system's default browser.

​	Open 尝试通过系统默认浏览器打开指定的 URL。

## func ResolveURL <- 0.65.0

``` go
func ResolveURL(u string) (string, error)
```

ResolveURL by requesting the u, it will try best to normalize the u. The format of u can be "9222", ":9222", "host:9222", "ws://host:9222", "wss://host:9222", "[https://host:9222](https://host:9222/)" "[http://host:9222](http://host:9222/)". The return string will look like: "ws://host:9222/devtools/browser/4371405f-84df-4ad6-9e0f-eab81f7521cc"

​	ResolveURL 请求指定的 URL (`u`) 并尝试尽可能规范化 URL。`u` 的格式可以是：`"9222"`、`":9222"`、`"host:9222"`、`"ws://host:9222"`、`"wss://host:9222"`、`"https://host:9222"` 或 `"http://host:9222"`。返回的字符串格式为：`"ws://host:9222/devtools/browser/4371405f-84df-4ad6-9e0f-eab81f7521cc"`。

## 类型

### type Browser 

``` go
type Browser struct {
	Context context.Context

	// Hosts are the candidates to download the browser.
	// Such as [HostGoogle] or [HostNPM].
    // Hosts 是用于已下载浏览器的候选主机。
	// 例如 [HostGoogle] 或 [HostNPM]。
	Hosts []Host

	// Revision of the browser to use
    // Revision 表示使用的浏览器版本号。
	Revision int

	// RootDir to download different browser versions.
    // RootDir 表示用于已下载不同浏览器版本的根目录。
	RootDir string

	// Log to print output
    // Logger 用于打印输出日志。
	Logger utils.Logger

	// LockPort a tcp port to prevent race downloading. Default is 2968 .
    // LockPort 表示用于防止竞争下载的 TCP 端口，默认值为 2968。
	LockPort int

	// HTTPClient to download the browser
    // HTTPClient 用于已下载浏览器的 HTTP 客户端。
	HTTPClient *http.Client
}
```

Browser is a helper to download browser smartly.

​	Browser 是一个帮助智能已下载浏览器的辅助工具。

### func NewBrowser 

``` go
func NewBrowser() *Browser
```

NewBrowser with default values.

​	NewBrowser 创建一个带有默认值的浏览器实例。

#### (*Browser) BinPath <- 0.112.9

``` go
func (lc *Browser) BinPath() string
```

BinPath to download the browser executable.

​	BinPath 返回已下载浏览器可执行文件的路径。

#### (*Browser) Dir 

``` go
func (lc *Browser) Dir() string
```

Dir to download the browser.

​	Dir 返回用于已下载浏览器的目录。

#### (*Browser) Download 

``` go
func (lc *Browser) Download() error
```

Download browser from the fastest host. It will race downloading a TCP packet from each host and use the fastest host.

​	Download 从最快的主机下载浏览器。它会通过 TCP 包的下载速度进行竞争并使用最快的主机。

#### (*Browser) Get 

``` go
func (lc *Browser) Get() (string, error)
```

Get is a smart helper to get the browser executable path. If [Browser.BinPath](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.BinPath) is not valid it will auto download the browser to [Browser.BinPath](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.BinPath).

​	Get 是一个智能工具，用于获取浏览器可执行文件的路径。如果 [Browser.BinPath](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.BinPath) 无效，它将自动下载浏览器到 [Browser.BinPath](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.BinPath)。

#### (*Browser) MustGet <- 0.90.0

``` go
func (lc *Browser) MustGet() string
```

MustGet is similar with Get.

​	MustGet 类似于 Get。

#### (*Browser) Validate <- 0.110.0

``` go
func (lc *Browser) Validate() error
```

Validate returns nil if the browser executable is valid. If the executable is malformed it will return error.

​	Validate 如果浏览器可执行文件有效，则返回 nil。如果可执行文件格式错误，则返回错误。

### type Host <- 0.91.0

``` go
type Host func(revision int) string
```

Host formats a revision number to a downloadable URL for the browser.

​	Host 将版本号格式化为浏览器可下载的 URL。

### type Launcher 

``` go
type Launcher struct {
	Flags map[flags.Flag][]string `json:"flags"`
	// contains filtered or unexported fields
}
```

Launcher is a helper to launch browser binary smartly.

​	Launcher 是一个帮助智能启动浏览器可执行文件的工具。

### func MustNewManaged <- 0.98.0

``` go
func MustNewManaged(serviceURL string) *Launcher
```

MustNewManaged is similar to NewManaged.

​	MustNewManaged 类似于 NewManaged。

### func New 

``` go
func New() *Launcher
```

New returns the default arguments to start browser. Headless will be enabled by default. Leakless will be enabled by default. UserDataDir will use OS tmp dir by default, this folder will usually be cleaned up by the OS after reboot. It will auto download the browser binary according to the current platform, check [Launcher.Bin](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Bin) and [Launcher.Revision](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Revision) for more info.

​	New 返回启动浏览器的默认参数。默认启用无头模式（Headless）、防泄漏模式（Leakless），并将 UserDataDir 设为操作系统临时目录，该目录通常会在系统重启后清除。它会根据当前平台自动下载浏览器二进制文件，更多信息请参阅 [Launcher.Bin](#launcher-bin) 和 [Launcher.Revision](#launcher-revision---01030)。

### func NewAppMode <- 0.106.1

``` go
func NewAppMode(u string) *Launcher
```

NewAppMode is a preset to run the browser like a native application. The u should be a URL.

​	NewAppMode 是用于以本地应用程序模式运行浏览器的预设。`u` 应为 URL。

### func NewManaged <- 0.98.0

``` go
func NewManaged(serviceURL string) (*Launcher, error)
```

NewManaged creates a default Launcher instance from launcher.Manager. The serviceURL must point to a launcher.Manager. It will send a http request to the serviceURL to get the default settings of the Launcher instance. For example if the launcher.Manager running on a Linux machine will return different default settings from the one on Mac. If Launcher.Leakless is enabled, the remote browser will be killed after the websocket is closed.

​	NewManaged 从 `launcher.Manager` 创建一个默认的 Launcher 实例。`serviceURL` 必须指向 `launcher.Manager`。它将向 `serviceURL` 发送 HTTP 请求以获取 Launcher 实例的默认设置。例如，如果 `launcher.Manager` 运行在 Linux 机器上，将返回与 Mac 不同的默认设置。如果启用了 `Launcher.Leakless`，当 WebSocket 关闭时，远程浏览器将被杀死。

### func NewUserMode 

``` go
func NewUserMode() *Launcher
```

NewUserMode is a preset to enable reusing current user data. Useful for automation of personal browser. If you see any error, it may because you can't launch debug port for existing browser, the solution is to completely close the running browser. Unfortunately, there's no API for rod to tell it automatically yet.

​	NewUserMode 是一个预设，用于启用复用当前用户数据。适用于个人浏览器的自动化。如果遇到错误，可能是因为无法为现有浏览器启动调试端口，解决方法是完全关闭运行中的浏览器。不幸的是，目前 Rod 尚未自动提供相应的 API。

#### (*Launcher) AlwaysOpenPDFExternally <- 0.115.0

``` go
func (l *Launcher) AlwaysOpenPDFExternally() *Launcher
```

AlwaysOpenPDFExternally switch. It will set chromium user preferences to enable the always_open_pdf_externally option.

​	AlwaysOpenPDFExternally 开关。它将设置 Chromium 用户首选项以启用 `always_open_pdf_externally` 选项。

#### (*Launcher) Append <- 0.48.0

``` go
func (l *Launcher) Append(name flags.Flag, values ...string) *Launcher
```

Append values to the flag.

​	Append 向指定标志追加值。

#### (*Launcher) Bin 

``` go
func (l *Launcher) Bin(path string) *Launcher
```

Bin of the browser binary path to launch, if the path is not empty the auto download will be disabled.

​	Bin 设置浏览器二进制路径。如果`path` 不为空，则禁用自动下载。

#### (*Launcher) Cleanup <- 0.49.7

``` go
func (l *Launcher) Cleanup()
```

Cleanup wait until the Browser exits and remove [flags.UserDataDir](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/launcher/flags#UserDataDir).

​	Cleanup 等待浏览器退出并移除 [flags.UserDataDir]({{< ref "/rod/flags#type-flag">}})。

#### (*Launcher) Client 

``` go
func (l *Launcher) Client() (*cdp.Client, error)
```

Client for launching browser remotely via the launcher.Manager.

​	Client 用于通过 `launcher.Manager` 远程启动浏览器。

#### (*Launcher) ClientHeader <- 0.106.0

``` go
func (l *Launcher) ClientHeader() (string, http.Header)
```

ClientHeader for launching browser remotely via the launcher.Manager.

​	ClientHeader 用于通过 `launcher.Manager` 远程启动浏览器。

#### (*Launcher) Context 

``` go
func (l *Launcher) Context(ctx context.Context) *Launcher
```

Context sets the context.

​	Context 设置上下文。

#### (*Launcher) Delete 

``` go
func (l *Launcher) Delete(name flags.Flag) *Launcher
```

Delete a flag.

​	Delete 删除一个标志。

#### (*Launcher) Devtools 

``` go
func (l *Launcher) Devtools(autoOpenForTabs bool) *Launcher
```

Devtools switch to auto open devtools for each tab.

​	Devtools 开关，自动为每个标签页打开开发者工具。

#### (*Launcher) Env <- 0.56.0

``` go
func (l *Launcher) Env(env ...string) *Launcher
```

Env to launch the browser process. The default value is [os.Environ](https://pkg.go.dev/os#Environ)(). Usually you use it to set the timezone env. Such as:

​	Env 设置启动浏览器进程的环境变量。默认值为 [os.Environ]({{< ref "/stdLib/os/os#func-environ">}})()。通常用于设置时区环境变量，例如：

```
Env(append(os.Environ(), "TZ=Asia/Tokyo")...)
```

#### (*Launcher) FormatArgs 

``` go
func (l *Launcher) FormatArgs() []string
```

FormatArgs returns the formatted arg list for cli.

​	FormatArgs 返回格式化后的命令行参数列表。

#### (*Launcher) Get 

``` go
func (l *Launcher) Get(name flags.Flag) string
```

Get flag's first value.

​	Get 返回指定标志的第一个值。

#### (*Launcher) GetFlags 

``` go
func (l *Launcher) GetFlags(name flags.Flag) ([]string, bool)
```

GetFlags from settings.

​	GetFlags 获取指定标志的所有值及其是否存在。

#### (*Launcher) Has <- 0.98.0

``` go
func (l *Launcher) Has(name flags.Flag) bool
```

Has flag or not.

​	Has 检查指定标志是否存在。

#### (*Launcher) Headless 

``` go
func (l *Launcher) Headless(enable bool) *Launcher
```

Headless switch. Whether to run browser in headless mode. A mode without visible UI.

​	Headless 开关，决定是否以无头模式运行浏览器（无可见 UI）。

#### (*Launcher) HeadlessNew <- 0.116.1

``` go
func (l *Launcher) HeadlessNew(enable bool) *Launcher
```

HeadlessNew switch is the "--headless=new" switch: https://developer.chrome.com/docs/chromium/new-headless

​	HeadlessNew 开关，启用 `--headless=new` 模式：[参考文档](https://developer.chrome.com/docs/chromium/new-headless)。

#### (*Launcher) IgnoreCerts <- 0.112.1

``` go
func (l *Launcher) IgnoreCerts(pks []crypto.PublicKey) error
```

IgnoreCerts configure the Chrome's ignore-certificate-errors-spki-list argument with the public keys.

​	IgnoreCerts 配置 Chrome 的 `ignore-certificate-errors-spki-list` 参数，设置公钥以忽略证书错误。

#### (*Launcher) JSON 

``` go
func (l *Launcher) JSON() []byte
```

JSON serialization.

​	JSON 返回序列化后的 JSON 数据。

#### (*Launcher) KeepUserDataDir 

``` go
func (l *Launcher) KeepUserDataDir() *Launcher
```

KeepUserDataDir after remote browser is closed. By default launcher.FlagUserDataDir will be removed.

​	KeepUserDataDir 在关闭远程浏览器后保留用户数据目录。默认情况下 `Launcher.FlagUserDataDir`（个人注释：未找到该方法） 会被移除。

#### (*Launcher) Kill <- 0.59.0

``` go
func (l *Launcher) Kill()
```

Kill the browser process.

​	Kill 终止浏览器进程。

#### (*Launcher) Launch 

``` go
func (l *Launcher) Launch() (string, error)
```

Launch a standalone temp browser instance and returns the debug url. bin and profileDir are optional, set them to empty to use the default values. If you want to reuse sessions, such as cookies, set the [Launcher.UserDataDir](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.UserDataDir) to the same location.

​	Launch 启动一个独立的临时浏览器实例，并返回调试 URL。`bin` 和 `profileDir` 为可选参数，设置为空则使用默认值。如果希望重用会话（如 Cookie），可以将 [Launcher.UserDataDir](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.UserDataDir) 设置为相同位置。

Please note launcher can only be used once.

​	注意：Launcher 只能使用一次。

#### (*Launcher) Leakless <- 0.57.1

``` go
func (l *Launcher) Leakless(enable bool) *Launcher
```

Leakless switch. If enabled, the browser will be force killed after the Go process exits. The doc of leakless: https://github.com/ysmood/leakless.

​	Leakless 开关。如果启用，当 Go 进程退出时，浏览器将被强制终止。Leakless 文档：[leakless](https://github.com/ysmood/leakless)。

#### (*Launcher) Logger <- 0.56.0

``` go
func (l *Launcher) Logger(w io.Writer) *Launcher
```

Logger to handle stdout and stderr from browser. For example, pipe all browser output to stdout:

​	Logger 用于处理浏览器的标准输出和错误输出。例如，将所有浏览器输出重定向到标准输出：

```go
launcher.New().Logger(os.Stdout)
```

#### (*Launcher) MustClient <- 0.106.0

``` go
func (l *Launcher) MustClient() *cdp.Client
```

MustClient similar to Launcher.Client.

​	MustClient 类似于 `Launcher.Client`，但会在失败时直接引发错误。

#### (*Launcher) MustLaunch <- 0.50.0

``` go
func (l *Launcher) MustLaunch() string
```

MustLaunch is similar to Launch.

​	MustLaunch 类似于 `Launch`，但会在失败时直接引发错误。

#### (*Launcher) NoSandbox <- 0.94.3

``` go
func (l *Launcher) NoSandbox(enable bool) *Launcher
```

NoSandbox switch. Whether to run browser in no-sandbox mode. Linux users may face "running as root without `--no-sandbox` is not supported" in some Linux/Chrome combinations. This function helps switch mode easily. Be aware disabling sandbox is not trivial. Use at your own risk. Related doc: https://bugs.chromium.org/p/chromium/issues/detail?id=638180

​	NoSandbox 开关。是否以无沙盒模式运行浏览器。对于 Linux 用户，在某些 Linux/Chrome 组合中可能会遇到“以 root 身份运行时不支持无 `--no-sandbox`”问题，此函数可轻松切换模式。注意：禁用沙盒可能带来安全风险，需自行承担风险。相关文档：[No Sandbox](https://bugs.chromium.org/p/chromium/issues/detail?id=638180)。

#### (*Launcher) PID 

``` go
func (l *Launcher) PID() int
```

PID returns the browser process pid.

​	PID 返回浏览器进程的 PID。

#### (*Launcher) Preferences <- 0.114.4

``` go
func (l *Launcher) Preferences(pref string) *Launcher
```

Preferences set chromium user preferences, such as set the default search engine or disable the pdf viewer. The pref is a json string, the doc is here https://src.chromium.org/viewvc/chrome/trunk/src/chrome/common/pref_names.cc

​	Preferences 设置 Chromium 用户首选项，例如设置默认搜索引擎或禁用 PDF 查看器。`pref` 是一个 JSON 字符串，参考文档：[首选项文档](https://src.chromium.org/viewvc/chrome/trunk/src/chrome/common/pref_names.cc)。

#### (*Launcher) ProfileDir <- 0.78.3

``` go
func (l *Launcher) ProfileDir(dir string) *Launcher
```

ProfileDir is the browser profile the browser will use. When set to empty, the profile 'Default' is used. Related article: https://superuser.com/a/377195

​	ProfileDir 指定浏览器将使用的用户配置目录。如果设置为空，将使用默认的配置目录“Default”。相关文章：[浏览器配置目录](https://superuser.com/a/377195)。

#### (*Launcher) Proxy <- 0.57.2

``` go
func (l *Launcher) Proxy(host string) *Launcher
```

Proxy for the browser.

​	Proxy 为浏览器设置代理。

#### (*Launcher) RemoteDebuggingPort 

``` go
func (l *Launcher) RemoteDebuggingPort(port int) *Launcher
```

RemoteDebuggingPort to launch the browser. Zero for a random port. Zero is the default value. If it's not zero and the Launcher.Leakless is disabled, the launcher will try to reconnect to it first, if the reconnection fails it will launch a new browser.

​	RemoteDebuggingPort 设置浏览器的调试端口。默认值为 0（随机端口）。如果值为 0 并且 `Launcher.Leakless` 被禁用，Launcher 将尝试重新连接到它，如果重新连接失败，将启动新的浏览器。

#### (*Launcher) Revision <- 0.103.0

``` go
func (l *Launcher) Revision(rev int) *Launcher
```

Revision of the browser to auto download.

​	Revision 设置浏览器的版本号以自动下载。

#### (*Launcher) Set 

``` go
func (l *Launcher) Set(name flags.Flag, values ...string) *Launcher
```

Set a command line argument when launching the browser. Be careful the first argument is a flag name, it shouldn't contain values. The values the will be joined with comma. A flag can have multiple values. If no values are provided the flag will be a boolean flag. You can use the [Launcher.FormatArgs](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.FormatArgs) to debug the final CLI arguments. List of available flags: https://peter.sh/experiments/chromium-command-line-switches

​	Set 在启动浏览器时设置命令行参数。注意，第一个参数是标志名称，不应包含值。值将用逗号连接。一个标志可以有多个值。如果未提供值，标志将成为布尔标志。可以使用 [Launcher.FormatArgs](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.FormatArgs) 调试最终的 CLI 参数。可用标志列表：[Chromium 命令行标志](https://peter.sh/experiments/chromium-command-line-switches)。

#### (*Launcher) StartURL <- 0.81.3

``` go
func (l *Launcher) StartURL(u string) *Launcher
```

StartURL to launch.

​	StartURL 设置启动时的 URL。

#### (*Launcher) UserDataDir 

``` go
func (l *Launcher) UserDataDir(dir string) *Launcher
```

UserDataDir is where the browser will look for all of its state, such as cookie and cache. When set to empty, browser will use current OS home dir. Related doc: https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md

​	UserDataDir 设置浏览器用于存储状态（如 Cookie 和缓存）的目录。如果设置为空，浏览器将使用当前操作系统的主目录。相关文档：[用户数据目录](https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md)。

#### (*Launcher) WorkingDir <- 0.56.0

``` go
func (l *Launcher) WorkingDir(path string) *Launcher
```

WorkingDir to launch the browser process.

​	WorkingDir 设置启动浏览器进程的工作目录。

#### (*Launcher) XVFB <- 0.86.1

``` go
func (l *Launcher) XVFB(args ...string) *Launcher
```

XVFB enables to run browser in by XVFB. Useful when you want to run headful mode on linux.

​	XVFB 启用 XVFB 运行浏览器。适用于希望在 Linux 上以有头模式运行浏览器的情况。

### type Manager <- 0.98.0

``` go
type Manager struct {
	// Logger for key events
    // Logger 用于记录关键事件
	Logger utils.Logger

	// Defaults should return the default Launcher settings
    // Defaults 返回默认的 Launcher 设置
	Defaults func(http.ResponseWriter, *http.Request) *Launcher

	// BeforeLaunch hook is called right before the launching with the Launcher instance that will be used
	// to launch the browser.
	// Such as use it to filter malicious values of Launcher.UserDataDir, Launcher.Bin, or Launcher.WorkingDir.
    // BeforeLaunch 在启动浏览器之前调用，用于过滤 Launcher 实例的恶意值，
    // 如 Launcher.UserDataDir、Launcher.Bin 或 Launcher.WorkingDir。
	BeforeLaunch func(*Launcher, http.ResponseWriter, *http.Request)
}
```

Manager is used to launch browsers via http server on another machine. The reason why we have Manager is after we launcher a browser, we can't dynamically change its CLI arguments, such as "--headless". The Manager allows us to decide what CLI arguments to pass to the browser when launch it remotely. The work flow looks like:

​	Manager 用于通过 HTTP 服务器在另一台机器上启动浏览器。Manager 的设计目的是在启动浏览器后无法动态更改其 CLI 参数（如 `--headless`）。通过 Manager 可以在远程启动浏览器时决定传递的 CLI 参数。工作流程如下：

```
|      Machine X       |                             Machine Y                                    |
| NewManaged("a.com") -|-> http.ListenAndServe("a.com", launcher.NewManager()) --> launch browser |

1. X send a http request to Y, Y respond default Launcher settings based the OS of Y.
2. X start a websocket connect to Y with the Launcher settings
3. Y launches a browser with the Launcher settings X
4. Y transparently proxy the websocket connect between X and the launched browser

1. X 向 Y 发送 HTTP 请求，Y 根据 Y 的操作系统返回默认的 Launcher 设置。
2. X 使用 Launcher 设置开始与 Y 的 WebSocket 连接。
3. Y 使用 X 的 Launcher 设置启动浏览器。
4. Y 透明地代理 X 与启动的浏览器之间的 WebSocket 连接。
```

### func NewManager <- 0.98.0

``` go
func NewManager() *Manager
```

NewManager instance.

​	NewManager 创建一个新的 Manager 实例。

#### (*Manager) ServeHTTP <- 0.98.0

``` go
func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

### type URLParser <- 0.56.0

``` go
type URLParser struct {
	URL    chan string
	Buffer string // buffer for the browser stdout 缓存浏览器标准输出
	// contains filtered or unexported fields
}
```

URLParser to get control url from stderr.

​	NewURLParser 创建一个新的 URLParser 实例。

### func NewURLParser <- 0.56.0

``` go
func NewURLParser() *URLParser
```

NewURLParser instance.

​	NewURLParser 创建一个新的 URLParser 实例。

#### (*URLParser) Context <- 0.101.0

``` go
func (r *URLParser) Context(ctx context.Context) *URLParser
```

Context sets the context.

​	Context 设置上下文。

#### (*URLParser) Err <- 0.89.2

``` go
func (r *URLParser) Err() error
```

Err returns the common error parsed from stdout and stderr.

​	Err 返回从标准输出和标准错误中解析出的常见错误。

#### (*URLParser) Write <- 0.56.0

``` go
func (r *URLParser) Write(p []byte) (n int, err error)
```

Write interface.

​	Write 接口实现，用于写入数据。
