+++
title = "launcher"
date = 2024-11-20T18:01:46+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

### Overview

A lib helps to find, launch or download the browser. You can also use it as a standalone lib without Rod.

## ![img](https://pkg.go.dev/static/shared/icon/code_gm_grey_24dp.svg) Documentation 

[Rendered for](https://go.dev/about#build-context)                   linux/amd64                   windows/amd64                   darwin/amd64                   js/wasm                

### Overview 

Package launcher for launching browser utils.

##### Example (Custom_launch)

``` go
```
##### Example (Print_browser_CLI_output)

``` go
```
##### Example (Use_system_browser)

``` go
```
### Index 

- [Constants](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#pkg-constants)
- [Variables](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#pkg-variables)
- [func HostGoogle(revision int) string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#HostGoogle)
- [func HostNPM(revision int) string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#HostNPM)
- [func HostPlaywright(revision int) string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#HostPlaywright)
- [func LookPath() (found string, has bool)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#LookPath)
- [func MustResolveURL(u string) string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#MustResolveURL)
- [func Open(url string)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Open)
- [func ResolveURL(u string) (string, error)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#ResolveURL)
- [type Browser](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser)
- - [func NewBrowser() *Browser](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#NewBrowser)
- - [func (lc *Browser) BinPath() string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.BinPath)
  - [func (lc *Browser) Dir() string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.Dir)
  - [func (lc *Browser) Download() error](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.Download)
  - [func (lc *Browser) Get() (string, error)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.Get)
  - [func (lc *Browser) MustGet() string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.MustGet)
  - [func (lc *Browser) Validate() error](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.Validate)
- [type Host](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Host)
- [type Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher)
- - [func MustNewManaged(serviceURL string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#MustNewManaged)
  - [func New() *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#New)
  - [func NewAppMode(u string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#NewAppMode)
  - [func NewManaged(serviceURL string) (*Launcher, error)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#NewManaged)
  - [func NewUserMode() *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#NewUserMode)
- - [func (l *Launcher) AlwaysOpenPDFExternally() *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.AlwaysOpenPDFExternally)
  - [func (l *Launcher) Append(name flags.Flag, values ...string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Append)
  - [func (l *Launcher) Bin(path string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Bin)
  - [func (l *Launcher) Cleanup()](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Cleanup)
  - [func (l *Launcher) Client() (*cdp.Client, error)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Client)
  - [func (l *Launcher) ClientHeader() (string, http.Header)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.ClientHeader)
  - [func (l *Launcher) Context(ctx context.Context) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Context)
  - [func (l *Launcher) Delete(name flags.Flag) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Delete)
  - [func (l *Launcher) Devtools(autoOpenForTabs bool) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Devtools)
  - [func (l *Launcher) Env(env ...string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Env)
  - [func (l *Launcher) FormatArgs() [\]string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.FormatArgs)
  - [func (l *Launcher) Get(name flags.Flag) string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Get)
  - [func (l *Launcher) GetFlags(name flags.Flag) ([\]string, bool)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.GetFlags)
  - [func (l *Launcher) Has(name flags.Flag) bool](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Has)
  - [func (l *Launcher) Headless(enable bool) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Headless)
  - [func (l *Launcher) HeadlessNew(enable bool) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.HeadlessNew)
  - [func (l *Launcher) IgnoreCerts(pks [\]crypto.PublicKey) error](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.IgnoreCerts)
  - [func (l *Launcher) JSON() [\]byte](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.JSON)
  - [func (l *Launcher) KeepUserDataDir() *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.KeepUserDataDir)
  - [func (l *Launcher) Kill()](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Kill)
  - [func (l *Launcher) Launch() (string, error)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Launch)
  - [func (l *Launcher) Leakless(enable bool) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Leakless)
  - [func (l *Launcher) Logger(w io.Writer) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Logger)
  - [func (l *Launcher) MustClient() *cdp.Client](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.MustClient)
  - [func (l *Launcher) MustLaunch() string](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.MustLaunch)
  - [func (l *Launcher) NoSandbox(enable bool) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.NoSandbox)
  - [func (l *Launcher) PID() int](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.PID)
  - [func (l *Launcher) Preferences(pref string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Preferences)
  - [func (l *Launcher) ProfileDir(dir string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.ProfileDir)
  - [func (l *Launcher) Proxy(host string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Proxy)
  - [func (l *Launcher) RemoteDebuggingPort(port int) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.RemoteDebuggingPort)
  - [func (l *Launcher) Revision(rev int) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Revision)
  - [func (l *Launcher) Set(name flags.Flag, values ...string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Set)
  - [func (l *Launcher) StartURL(u string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.StartURL)
  - [func (l *Launcher) UserDataDir(dir string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.UserDataDir)
  - [func (l *Launcher) WorkingDir(path string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.WorkingDir)
  - [func (l *Launcher) XVFB(args ...string) *Launcher](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.XVFB)
- [type Manager](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Manager)
- - [func NewManager() *Manager](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#NewManager)
- - [func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Manager.ServeHTTP)
- [type URLParser](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#URLParser)
- - [func NewURLParser() *URLParser](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#NewURLParser)
- - [func (r *URLParser) Context(ctx context.Context) *URLParser](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#URLParser.Context)
  - [func (r *URLParser) Err() error](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#URLParser.Err)
  - [func (r *URLParser) Write(p [\]byte) (n int, err error)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#URLParser.Write)

#### Examples 

- [Package (Custom_launch)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#example-package-Custom_launch)
- [Package (Print_browser_CLI_output)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#example-package-Print_browser_CLI_output)
- [Package (Use_system_browser)](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#example-package-Use_system_browser)

### Constants 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/manager.go#L17)

``` go
const (
	// HeaderName for remote launch.
	HeaderName = "Rod-Launcher"
)
```

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/revision.go#L6)

``` go
const RevisionDefault = 1321438
```

RevisionDefault for chromium.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/revision.go#L9)

``` go
const RevisionPlaywright = 1124
```

RevisionPlaywright for arm linux.

### Variables 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher/browser.go#L70)

``` go
var DefaultBrowserDir = filepath.Join(map[string]string{
	"windows": os.Getenv("APPDATA"),
	"darwin":  filepath.Join(os.Getenv("HOME"), ".cache"),
	"linux":   filepath.Join(os.Getenv("HOME"), ".cache"),
}[runtime.GOOS], "rod", "browser")
```

DefaultBrowserDir for downloaded browser. For unix is "$HOME/.cache/rod/browser", for Windows it's "%APPDATA%\rod\browser".

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

### Functions 

#### func HostGoogle 

``` go
func HostGoogle(revision int) string
```

HostGoogle to download browser.

#### func HostNPM <- 0.102.0

``` go
func HostNPM(revision int) string
```

HostNPM to download browser.

#### func HostPlaywright <- 0.106.3

``` go
func HostPlaywright(revision int) string
```

HostPlaywright to download browser.

#### func LookPath <- 0.91.0

``` go
func LookPath() (found string, has bool)
```

LookPath searches for the browser executable from often used paths on current operating system.

#### func MustResolveURL <- 0.65.0

``` go
func MustResolveURL(u string) string
```

MustResolveURL is similar to ResolveURL.

#### func Open <- 0.91.0

``` go
func Open(url string)
```

Open tries to open the url via system's default browser.

#### func ResolveURL <- 0.65.0

``` go
func ResolveURL(u string) (string, error)
```

ResolveURL by requesting the u, it will try best to normalize the u. The format of u can be "9222", ":9222", "host:9222", "ws://host:9222", "wss://host:9222", "[https://host:9222](https://host:9222/)" "[http://host:9222](http://host:9222/)". The return string will look like: "ws://host:9222/devtools/browser/4371405f-84df-4ad6-9e0f-eab81f7521cc"

### Types 

#### type Browser 

``` go
type Browser struct {
	Context context.Context

	// Hosts are the candidates to download the browser.
	// Such as [HostGoogle] or [HostNPM].
	Hosts []Host

	// Revision of the browser to use
	Revision int

	// RootDir to download different browser versions.
	RootDir string

	// Log to print output
	Logger utils.Logger

	// LockPort a tcp port to prevent race downloading. Default is 2968 .
	LockPort int

	// HTTPClient to download the browser
	HTTPClient *http.Client
}
```

Browser is a helper to download browser smartly.

#### func NewBrowser 

``` go
func NewBrowser() *Browser
```

NewBrowser with default values.

#### (*Browser) BinPath <- 0.112.9

``` go
func (lc *Browser) BinPath() string
```

BinPath to download the browser executable.

#### (*Browser) Dir 

``` go
func (lc *Browser) Dir() string
```

Dir to download the browser.

#### (*Browser) Download 

``` go
func (lc *Browser) Download() error
```

Download browser from the fastest host. It will race downloading a TCP packet from each host and use the fastest host.

#### (*Browser) Get 

``` go
func (lc *Browser) Get() (string, error)
```

Get is a smart helper to get the browser executable path. If [Browser.BinPath](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.BinPath) is not valid it will auto download the browser to [Browser.BinPath](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Browser.BinPath).

#### (*Browser) MustGet <- 0.90.0

``` go
func (lc *Browser) MustGet() string
```

MustGet is similar with Get.

#### (*Browser) Validate <- 0.110.0

``` go
func (lc *Browser) Validate() error
```

Validate returns nil if the browser executable is valid. If the executable is malformed it will return error.

#### type Host <- 0.91.0

``` go
type Host func(revision int) string
```

Host formats a revision number to a downloadable URL for the browser.

#### type Launcher 

``` go
type Launcher struct {
	Flags map[flags.Flag][]string `json:"flags"`
	// contains filtered or unexported fields
}
```

Launcher is a helper to launch browser binary smartly.

#### func MustNewManaged <- 0.98.0

``` go
func MustNewManaged(serviceURL string) *Launcher
```

MustNewManaged is similar to NewManaged.

#### func New 

``` go
func New() *Launcher
```

New returns the default arguments to start browser. Headless will be enabled by default. Leakless will be enabled by default. UserDataDir will use OS tmp dir by default, this folder will usually be cleaned up by the OS after reboot. It will auto download the browser binary according to the current platform, check [Launcher.Bin](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Bin) and [Launcher.Revision](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.Revision) for more info.

#### func NewAppMode <- 0.106.1

``` go
func NewAppMode(u string) *Launcher
```

NewAppMode is a preset to run the browser like a native application. The u should be a URL.

#### func NewManaged <- 0.98.0

``` go
func NewManaged(serviceURL string) (*Launcher, error)
```

NewManaged creates a default Launcher instance from launcher.Manager. The serviceURL must point to a launcher.Manager. It will send a http request to the serviceURL to get the default settings of the Launcher instance. For example if the launcher.Manager running on a Linux machine will return different default settings from the one on Mac. If Launcher.Leakless is enabled, the remote browser will be killed after the websocket is closed.

#### func NewUserMode 

``` go
func NewUserMode() *Launcher
```

NewUserMode is a preset to enable reusing current user data. Useful for automation of personal browser. If you see any error, it may because you can't launch debug port for existing browser, the solution is to completely close the running browser. Unfortunately, there's no API for rod to tell it automatically yet.

#### (*Launcher) AlwaysOpenPDFExternally <- 0.115.0

``` go
func (l *Launcher) AlwaysOpenPDFExternally() *Launcher
```

AlwaysOpenPDFExternally switch. It will set chromium user preferences to enable the always_open_pdf_externally option.

#### (*Launcher) Append <- 0.48.0

``` go
func (l *Launcher) Append(name flags.Flag, values ...string) *Launcher
```

Append values to the flag.

#### (*Launcher) Bin 

``` go
func (l *Launcher) Bin(path string) *Launcher
```

Bin of the browser binary path to launch, if the path is not empty the auto download will be disabled.

#### (*Launcher) Cleanup <- 0.49.7

``` go
func (l *Launcher) Cleanup()
```

Cleanup wait until the Browser exits and remove [flags.UserDataDir](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/launcher/flags#UserDataDir).

#### (*Launcher) Client 

``` go
func (l *Launcher) Client() (*cdp.Client, error)
```

Client for launching browser remotely via the launcher.Manager.

#### (*Launcher) ClientHeader <- 0.106.0

``` go
func (l *Launcher) ClientHeader() (string, http.Header)
```

ClientHeader for launching browser remotely via the launcher.Manager.

#### (*Launcher) Context 

``` go
func (l *Launcher) Context(ctx context.Context) *Launcher
```

Context sets the context.

#### (*Launcher) Delete 

``` go
func (l *Launcher) Delete(name flags.Flag) *Launcher
```

Delete a flag.

#### (*Launcher) Devtools 

``` go
func (l *Launcher) Devtools(autoOpenForTabs bool) *Launcher
```

Devtools switch to auto open devtools for each tab.

#### (*Launcher) Env <- 0.56.0

``` go
func (l *Launcher) Env(env ...string) *Launcher
```

Env to launch the browser process. The default value is [os.Environ](https://pkg.go.dev/os#Environ)(). Usually you use it to set the timezone env. Such as:

```
Env(append(os.Environ(), "TZ=Asia/Tokyo")...)
```

#### (*Launcher) FormatArgs 

``` go
func (l *Launcher) FormatArgs() []string
```

FormatArgs returns the formatted arg list for cli.

#### (*Launcher) Get 

``` go
func (l *Launcher) Get(name flags.Flag) string
```

Get flag's first value.

#### (*Launcher) GetFlags 

``` go
func (l *Launcher) GetFlags(name flags.Flag) ([]string, bool)
```

GetFlags from settings.

#### (*Launcher) Has <- 0.98.0

``` go
func (l *Launcher) Has(name flags.Flag) bool
```

Has flag or not.

#### (*Launcher) Headless 

``` go
func (l *Launcher) Headless(enable bool) *Launcher
```

Headless switch. Whether to run browser in headless mode. A mode without visible UI.

#### (*Launcher) HeadlessNew <- 0.116.1

``` go
func (l *Launcher) HeadlessNew(enable bool) *Launcher
```

HeadlessNew switch is the "--headless=new" switch: https://developer.chrome.com/docs/chromium/new-headless

#### (*Launcher) IgnoreCerts <- 0.112.1

``` go
func (l *Launcher) IgnoreCerts(pks []crypto.PublicKey) error
```

IgnoreCerts configure the Chrome's ignore-certificate-errors-spki-list argument with the public keys.

#### (*Launcher) JSON 

``` go
func (l *Launcher) JSON() []byte
```

JSON serialization.

#### (*Launcher) KeepUserDataDir 

``` go
func (l *Launcher) KeepUserDataDir() *Launcher
```

KeepUserDataDir after remote browser is closed. By default launcher.FlagUserDataDir will be removed.

#### (*Launcher) Kill <- 0.59.0

``` go
func (l *Launcher) Kill()
```

Kill the browser process.

#### (*Launcher) Launch 

``` go
func (l *Launcher) Launch() (string, error)
```

Launch a standalone temp browser instance and returns the debug url. bin and profileDir are optional, set them to empty to use the default values. If you want to reuse sessions, such as cookies, set the [Launcher.UserDataDir](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.UserDataDir) to the same location.

Please note launcher can only be used once.

#### (*Launcher) Leakless <- 0.57.1

``` go
func (l *Launcher) Leakless(enable bool) *Launcher
```

Leakless switch. If enabled, the browser will be force killed after the Go process exits. The doc of leakless: https://github.com/ysmood/leakless.

#### (*Launcher) Logger <- 0.56.0

``` go
func (l *Launcher) Logger(w io.Writer) *Launcher
```

Logger to handle stdout and stderr from browser. For example, pipe all browser output to stdout:

```
launcher.New().Logger(os.Stdout)
```

#### (*Launcher) MustClient <- 0.106.0

``` go
func (l *Launcher) MustClient() *cdp.Client
```

MustClient similar to Launcher.Client.

#### (*Launcher) MustLaunch <- 0.50.0

``` go
func (l *Launcher) MustLaunch() string
```

MustLaunch is similar to Launch.

#### (*Launcher) NoSandbox <- 0.94.3

``` go
func (l *Launcher) NoSandbox(enable bool) *Launcher
```

NoSandbox switch. Whether to run browser in no-sandbox mode. Linux users may face "running as root without --no-sandbox is not supported" in some Linux/Chrome combinations. This function helps switch mode easily. Be aware disabling sandbox is not trivial. Use at your own risk. Related doc: https://bugs.chromium.org/p/chromium/issues/detail?id=638180

#### (*Launcher) PID 

``` go
func (l *Launcher) PID() int
```

PID returns the browser process pid.

#### (*Launcher) Preferences <- 0.114.4

``` go
func (l *Launcher) Preferences(pref string) *Launcher
```

Preferences set chromium user preferences, such as set the default search engine or disable the pdf viewer. The pref is a json string, the doc is here https://src.chromium.org/viewvc/chrome/trunk/src/chrome/common/pref_names.cc

#### (*Launcher) ProfileDir <- 0.78.3

``` go
func (l *Launcher) ProfileDir(dir string) *Launcher
```

ProfileDir is the browser profile the browser will use. When set to empty, the profile 'Default' is used. Related article: https://superuser.com/a/377195

#### (*Launcher) Proxy <- 0.57.2

``` go
func (l *Launcher) Proxy(host string) *Launcher
```

Proxy for the browser.

#### (*Launcher) RemoteDebuggingPort 

``` go
func (l *Launcher) RemoteDebuggingPort(port int) *Launcher
```

RemoteDebuggingPort to launch the browser. Zero for a random port. Zero is the default value. If it's not zero and the Launcher.Leakless is disabled, the launcher will try to reconnect to it first, if the reconnection fails it will launch a new browser.

#### (*Launcher) Revision <- 0.103.0

``` go
func (l *Launcher) Revision(rev int) *Launcher
```

Revision of the browser to auto download.

#### (*Launcher) Set 

``` go
func (l *Launcher) Set(name flags.Flag, values ...string) *Launcher
```

Set a command line argument when launching the browser. Be careful the first argument is a flag name, it shouldn't contain values. The values the will be joined with comma. A flag can have multiple values. If no values are provided the flag will be a boolean flag. You can use the [Launcher.FormatArgs](https://pkg.go.dev/github.com/go-rod/rod/lib/launcher#Launcher.FormatArgs) to debug the final CLI arguments. List of available flags: https://peter.sh/experiments/chromium-command-line-switches

#### (*Launcher) StartURL <- 0.81.3

``` go
func (l *Launcher) StartURL(u string) *Launcher
```

StartURL to launch.

#### (*Launcher) UserDataDir 

``` go
func (l *Launcher) UserDataDir(dir string) *Launcher
```

UserDataDir is where the browser will look for all of its state, such as cookie and cache. When set to empty, browser will use current OS home dir. Related doc: https://chromium.googlesource.com/chromium/src/+/master/docs/user_data_dir.md

#### (*Launcher) WorkingDir <- 0.56.0

``` go
func (l *Launcher) WorkingDir(path string) *Launcher
```

WorkingDir to launch the browser process.

#### (*Launcher) XVFB <- 0.86.1

``` go
func (l *Launcher) XVFB(args ...string) *Launcher
```

XVFB enables to run browser in by XVFB. Useful when you want to run headful mode on linux.

#### type Manager <- 0.98.0

``` go
type Manager struct {
	// Logger for key events
	Logger utils.Logger

	// Defaults should return the default Launcher settings
	Defaults func(http.ResponseWriter, *http.Request) *Launcher

	// BeforeLaunch hook is called right before the launching with the Launcher instance that will be used
	// to launch the browser.
	// Such as use it to filter malicious values of Launcher.UserDataDir, Launcher.Bin, or Launcher.WorkingDir.
	BeforeLaunch func(*Launcher, http.ResponseWriter, *http.Request)
}
```

Manager is used to launch browsers via http server on another machine. The reason why we have Manager is after we launcher a browser, we can't dynamically change its CLI arguments, such as "--headless". The Manager allows us to decide what CLI arguments to pass to the browser when launch it remotely. The work flow looks like:

```
|      Machine X       |                             Machine Y                                    |
| NewManaged("a.com") -|-> http.ListenAndServe("a.com", launcher.NewManager()) --> launch browser |

1. X send a http request to Y, Y respond default Launcher settings based the OS of Y.
2. X start a websocket connect to Y with the Launcher settings
3. Y launches a browser with the Launcher settings X
4. Y transparently proxy the websocket connect between X and the launched browser
```

#### func NewManager <- 0.98.0

``` go
func NewManager() *Manager
```

NewManager instance.

#### (*Manager) ServeHTTP <- 0.98.0

``` go
func (m *Manager) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

#### type URLParser <- 0.56.0

``` go
type URLParser struct {
	URL    chan string
	Buffer string // buffer for the browser stdout
	// contains filtered or unexported fields
}
```

URLParser to get control url from stderr.

#### func NewURLParser <- 0.56.0

``` go
func NewURLParser() *URLParser
```

NewURLParser instance.

#### (*URLParser) Context <- 0.101.0

``` go
func (r *URLParser) Context(ctx context.Context) *URLParser
```

Context sets the context.

#### (*URLParser) Err <- 0.89.2

``` go
func (r *URLParser) Err() error
```

Err returns the common error parsed from stdout and stderr.

#### (*URLParser) Write <- 0.56.0

``` go
func (r *URLParser) Write(p []byte) (n int, err error)
```

Write interface.
