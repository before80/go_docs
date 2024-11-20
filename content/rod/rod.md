+++
title = "rod"
date = 2024-11-20T18:01:04+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

#### [Documentation](https://go-rod.github.io/) | [API reference](https://pkg.go.dev/github.com/go-rod/rod?tab=doc) | [FAQ](https://go-rod.github.io/#/faq/README)

Rod is a high-level driver directly based on [DevTools Protocol](https://chromedevtools.github.io/devtools-protocol). It's designed for web automation and scraping for both high-level and low-level use, senior developers can use the low-level packages and functions to easily customize or build up their own version of Rod, the high-level functions are just examples to build a default version of Rod.

[中文 API 文档](https://pkg.go.dev/github.com/go-rod/go-rod-chinese)

#### Features

- Chained context design, intuitive to timeout or cancel the long-running task
- Auto-wait elements to be ready
- Debugging friendly, auto input tracing, remote monitoring headless browser
- Thread-safe for all operations
- Automatically find or download [browser](https://github.com/go-rod/rod/blob/v0.116.2/lib/launcher)
- High-level helpers like WaitStable, WaitRequestIdle, HijackRequests, WaitDownload, etc
- Two-step WaitEvent design, never miss an event ([how it works](https://github.com/ysmood/goob))
- Correctly handles nested iframes or shadow DOMs
- No zombie browser process after the crash ([how it works](https://github.com/ysmood/leakless))
- [CI](https://github.com/go-rod/rod/actions) enforced 100% test coverage

#### Examples

Please check the [examples_test.go](https://github.com/go-rod/rod/blob/v0.116.2/examples_test.go) file first, then check the [examples](https://github.com/go-rod/rod/blob/v0.116.2/lib/examples) folder.

For more detailed examples, please search the unit tests. Such as the usage of method `HandleAuth`, you can search all the `*_test.go` files that contain `HandleAuth`, for example, use Github online [search in repository](https://github.com/go-rod/rod/search?q=HandleAuth&unscoped_q=HandleAuth). You can also search the GitHub [issues](https://github.com/go-rod/rod/issues) or [discussions](https://github.com/go-rod/rod/discussions), a lot of usage examples are recorded there.

[Here](https://github.com/go-rod/rod/blob/v0.116.2/lib/examples/compare-chromedp) is a comparison of the examples between rod and Chromedp.

If you have questions, please raise an [issues](https://github.com/go-rod/rod/issues)/[discussions](https://github.com/go-rod/rod/discussions) or join the [chat room](https://discord.gg/CpevuvY).

#### Join us

Your help is more than welcome! Even just open an issue to ask a question may greatly help others.

Please read [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html) before you ask questions.

We use Github Projects to manage tasks, you can see the priority and progress of the issues [here](https://github.com/go-rod/rod/projects).

If you want to contribute please read the [Contributor Guide](https://github.com/go-rod/rod/blob/v0.116.2/.github/CONTRIBUTING.md).

Collapse ▴

## ![img](https://pkg.go.dev/static/shared/icon/code_gm_grey_24dp.svg) Documentation 

### Overview 

Package rod is a high-level driver directly based on DevTools Protocol.

##### Example (Basic)

``` go
```
##### Example (Context_and_EachEvent)

``` go
```
##### Example (Context_and_timeout)

``` go
```
##### Example (Customize_browser_launch)

``` go
```
##### Example (Customize_retry_strategy)

``` go
```
##### Example (Direct_cdp)

``` go
```
##### Example (Disable_headless_to_debug)

``` go
```
##### Example (Download_file)

``` go
```
##### Example (Error_handling)

``` go
```
##### Example (Eval_reuse_remote_object)

``` go
```
##### Example (Handle_events)

``` go
```
##### Example (Hijack_requests)

``` go
```
##### Example (Load_extension)

``` go
```
##### Example (Log_cdp_traffic)

``` go
```
##### Example (Page_pdf)

``` go
```
##### Example (Page_screenshot)

``` go
```
##### Example (Page_scroll_screenshot)

``` go
```
##### Example (Race_selectors)

``` go
```
##### Example (Search)

``` go
```
##### Example (States)

``` go
```
##### Example (Wait_for_animation)

``` go
```
##### Example (Wait_for_request)

``` go
```
### Index 

- [Variables](https://pkg.go.dev/github.com/go-rod/rod#pkg-variables)
- [func NotFoundSleeper() utils.Sleeper](https://pkg.go.dev/github.com/go-rod/rod#NotFoundSleeper)
- [func Try(fn func()) (err error)](https://pkg.go.dev/github.com/go-rod/rod#Try)
- [type Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser)
- - [func New() *Browser](https://pkg.go.dev/github.com/go-rod/rod#New)
- - [func (b *Browser) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res [\]byte, err error)](https://pkg.go.dev/github.com/go-rod/rod#Browser.Call)
  - [func (b *Browser) CancelTimeout() *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.CancelTimeout)
  - [func (b *Browser) Client(c CDPClient) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.Client)
  - [func (b *Browser) Close() error](https://pkg.go.dev/github.com/go-rod/rod#Browser.Close)
  - [func (b *Browser) Connect() error](https://pkg.go.dev/github.com/go-rod/rod#Browser.Connect)
  - [func (b *Browser) Context(ctx context.Context) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.Context)
  - [func (b *Browser) ControlURL(url string) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.ControlURL)
  - [func (b *Browser) DefaultDevice(d devices.Device) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.DefaultDevice)
  - [func (b *Browser) DisableDomain(sessionID proto.TargetSessionID, req proto.Request) (restore func())](https://pkg.go.dev/github.com/go-rod/rod#Browser.DisableDomain)
  - [func (b *Browser) EachEvent(callbacks ...interface{}) (wait func())](https://pkg.go.dev/github.com/go-rod/rod#Browser.EachEvent)
  - [func (b *Browser) EnableDomain(sessionID proto.TargetSessionID, req proto.Request) (restore func())](https://pkg.go.dev/github.com/go-rod/rod#Browser.EnableDomain)
  - [func (b *Browser) Event() <-chan *Message](https://pkg.go.dev/github.com/go-rod/rod#Browser.Event)
  - [func (b *Browser) GetContext() context.Context](https://pkg.go.dev/github.com/go-rod/rod#Browser.GetContext)
  - [func (b *Browser) GetCookies() ([\]*proto.NetworkCookie, error)](https://pkg.go.dev/github.com/go-rod/rod#Browser.GetCookies)
  - [func (b *Browser) HandleAuth(username, password string) func() error](https://pkg.go.dev/github.com/go-rod/rod#Browser.HandleAuth)
  - [func (b *Browser) HijackRequests() *HijackRouter](https://pkg.go.dev/github.com/go-rod/rod#Browser.HijackRequests)
  - [func (b *Browser) IgnoreCertErrors(enable bool) error](https://pkg.go.dev/github.com/go-rod/rod#Browser.IgnoreCertErrors)
  - [func (b *Browser) Incognito() (*Browser, error)](https://pkg.go.dev/github.com/go-rod/rod#Browser.Incognito)
  - [func (b *Browser) LoadState(sessionID proto.TargetSessionID, method proto.Request) (has bool)](https://pkg.go.dev/github.com/go-rod/rod#Browser.LoadState)
  - [func (b *Browser) Logger(l utils.Logger) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.Logger)
  - [func (b *Browser) Monitor(url string) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.Monitor)
  - [func (b *Browser) MustClose()](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustClose)
  - [func (b *Browser) MustConnect() *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustConnect)
  - [func (b *Browser) MustGetCookies() [\]*proto.NetworkCookie](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustGetCookies)
  - [func (b *Browser) MustHandleAuth(username, password string) (wait func())](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustHandleAuth)
  - [func (b *Browser) MustIgnoreCertErrors(enable bool) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustIgnoreCertErrors)
  - [func (b *Browser) MustIncognito() *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustIncognito)
  - [func (b *Browser) MustPage(url ...string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustPage)
  - [func (b *Browser) MustPageFromTargetID(targetID proto.TargetTargetID) *Page](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustPageFromTargetID)
  - [func (b *Browser) MustPages() Pages](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustPages)
  - [func (b *Browser) MustSetCookies(cookies ...*proto.NetworkCookie) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustSetCookies)
  - [func (b *Browser) MustVersion() *proto.BrowserGetVersionResult](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustVersion)
  - [func (b *Browser) MustWaitDownload() func() [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Browser.MustWaitDownload)
  - [func (b *Browser) NoDefaultDevice() *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.NoDefaultDevice)
  - [func (b *Browser) Page(opts proto.TargetCreateTarget) (p *Page, err error)](https://pkg.go.dev/github.com/go-rod/rod#Browser.Page)
  - [func (b *Browser) PageFromSession(sessionID proto.TargetSessionID) *Page](https://pkg.go.dev/github.com/go-rod/rod#Browser.PageFromSession)
  - [func (b *Browser) PageFromTarget(targetID proto.TargetTargetID) (*Page, error)](https://pkg.go.dev/github.com/go-rod/rod#Browser.PageFromTarget)
  - [func (b *Browser) Pages() (Pages, error)](https://pkg.go.dev/github.com/go-rod/rod#Browser.Pages)
  - [func (b *Browser) RemoveState(key interface{})](https://pkg.go.dev/github.com/go-rod/rod#Browser.RemoveState)
  - [func (b *Browser) ServeMonitor(host string) string](https://pkg.go.dev/github.com/go-rod/rod#Browser.ServeMonitor)
  - [func (b *Browser) SetCookies(cookies [\]*proto.NetworkCookieParam) error](https://pkg.go.dev/github.com/go-rod/rod#Browser.SetCookies)
  - [func (b *Browser) Sleeper(sleeper func() utils.Sleeper) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.Sleeper)
  - [func (b *Browser) SlowMotion(delay time.Duration) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.SlowMotion)
  - [func (b *Browser) Timeout(d time.Duration) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.Timeout)
  - [func (b *Browser) Trace(enable bool) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.Trace)
  - [func (b *Browser) Version() (*proto.BrowserGetVersionResult, error)](https://pkg.go.dev/github.com/go-rod/rod#Browser.Version)
  - [func (b *Browser) WaitDownload(dir string) func() (info *proto.PageDownloadWillBegin)](https://pkg.go.dev/github.com/go-rod/rod#Browser.WaitDownload)
  - [func (b *Browser) WaitEvent(e proto.Event) (wait func())](https://pkg.go.dev/github.com/go-rod/rod#Browser.WaitEvent)
  - [func (b *Browser) WithCancel() (*Browser, func())](https://pkg.go.dev/github.com/go-rod/rod#Browser.WithCancel)
  - [func (b *Browser) WithPanic(fail func(interface{})) *Browser](https://pkg.go.dev/github.com/go-rod/rod#Browser.WithPanic)
- [type CDPClient](https://pkg.go.dev/github.com/go-rod/rod#CDPClient)
- [type CoveredError](https://pkg.go.dev/github.com/go-rod/rod#CoveredError)
- - [func (e *CoveredError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#CoveredError.Error)
  - [func (e *CoveredError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#CoveredError.Is)
  - [func (e *CoveredError) Unwrap() error](https://pkg.go.dev/github.com/go-rod/rod#CoveredError.Unwrap)
- [type Element](https://pkg.go.dev/github.com/go-rod/rod#Element)
- - [func (el *Element) Attribute(name string) (*string, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Attribute)
  - [func (el *Element) BackgroundImage() ([\]byte, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.BackgroundImage)
  - [func (el *Element) Blur() error](https://pkg.go.dev/github.com/go-rod/rod#Element.Blur)
  - [func (el *Element) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res [\]byte, err error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Call)
  - [func (el *Element) CancelTimeout() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.CancelTimeout)
  - [func (el *Element) CanvasToImage(format string, quality float64) ([\]byte, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.CanvasToImage)
  - [func (el *Element) Click(button proto.InputMouseButton, clickCount int) error](https://pkg.go.dev/github.com/go-rod/rod#Element.Click)
  - [func (el *Element) ContainsElement(target *Element) (bool, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.ContainsElement)
  - [func (el *Element) Context(ctx context.Context) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.Context)
  - [func (el *Element) Describe(depth int, pierce bool) (*proto.DOMNode, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Describe)
  - [func (el *Element) Disabled() (bool, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Disabled)
  - [func (el *Element) Element(selector string) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Element)
  - [func (el *Element) ElementByJS(opts *EvalOptions) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementByJS)
  - [func (el *Element) ElementR(selector, jsRegex string) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementR)
  - [func (el *Element) ElementX(xPath string) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementX)
  - [func (el *Element) Elements(selector string) (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Elements)
  - [func (el *Element) ElementsByJS(opts *EvalOptions) (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementsByJS)
  - [func (el *Element) ElementsX(xpath string) (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementsX)
  - [func (el *Element) Equal(elm *Element) (bool, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Equal)
  - [func (el *Element) Eval(js string, params ...interface{}) (*proto.RuntimeRemoteObject, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Eval)
  - [func (el *Element) Evaluate(opts *EvalOptions) (*proto.RuntimeRemoteObject, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Evaluate)
  - [func (el *Element) Focus() error](https://pkg.go.dev/github.com/go-rod/rod#Element.Focus)
  - [func (el *Element) Frame() (*Page, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Frame)
  - [func (el *Element) GetContext() context.Context](https://pkg.go.dev/github.com/go-rod/rod#Element.GetContext)
  - [func (el *Element) GetSessionID() proto.TargetSessionID](https://pkg.go.dev/github.com/go-rod/rod#Element.GetSessionID)
  - [func (el *Element) GetXPath(optimized bool) (string, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.GetXPath)
  - [func (el *Element) HTML() (string, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.HTML)
  - [func (el *Element) Has(selector string) (bool, *Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Has)
  - [func (el *Element) HasR(selector, jsRegex string) (bool, *Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.HasR)
  - [func (el *Element) HasX(selector string) (bool, *Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.HasX)
  - [func (el *Element) Hover() error](https://pkg.go.dev/github.com/go-rod/rod#Element.Hover)
  - [func (el *Element) Input(text string) error](https://pkg.go.dev/github.com/go-rod/rod#Element.Input)
  - [func (el *Element) InputColor(color string) error](https://pkg.go.dev/github.com/go-rod/rod#Element.InputColor)
  - [func (el *Element) InputTime(t time.Time) error](https://pkg.go.dev/github.com/go-rod/rod#Element.InputTime)
  - [func (el *Element) Interactable() (pt *proto.Point, err error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Interactable)
  - [func (el *Element) KeyActions() (*KeyActions, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.KeyActions)
  - [func (el *Element) Matches(selector string) (bool, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Matches)
  - [func (el *Element) MoveMouseOut() error](https://pkg.go.dev/github.com/go-rod/rod#Element.MoveMouseOut)
  - [func (el *Element) MustAttribute(name string) *string](https://pkg.go.dev/github.com/go-rod/rod#Element.MustAttribute)
  - [func (el *Element) MustBackgroundImage() [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Element.MustBackgroundImage)
  - [func (el *Element) MustBlur() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustBlur)
  - [func (el *Element) MustCanvasToImage() [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Element.MustCanvasToImage)
  - [func (el *Element) MustClick() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustClick)
  - [func (el *Element) MustContainsElement(target *Element) bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustContainsElement)
  - [func (el *Element) MustDescribe() *proto.DOMNode](https://pkg.go.dev/github.com/go-rod/rod#Element.MustDescribe)
  - [func (el *Element) MustDisabled() bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustDisabled)
  - [func (el *Element) MustDoubleClick() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustDoubleClick)
  - [func (el *Element) MustElement(selector string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustElement)
  - [func (el *Element) MustElementByJS(js string, params ...interface{}) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustElementByJS)
  - [func (el *Element) MustElementR(selector, regex string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustElementR)
  - [func (el *Element) MustElementX(xpath string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustElementX)
  - [func (el *Element) MustElements(selector string) Elements](https://pkg.go.dev/github.com/go-rod/rod#Element.MustElements)
  - [func (el *Element) MustElementsByJS(js string, params ...interface{}) Elements](https://pkg.go.dev/github.com/go-rod/rod#Element.MustElementsByJS)
  - [func (el *Element) MustElementsX(xpath string) Elements](https://pkg.go.dev/github.com/go-rod/rod#Element.MustElementsX)
  - [func (el *Element) MustEqual(elm *Element) bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustEqual)
  - [func (el *Element) MustEval(js string, params ...interface{}) gson.JSON](https://pkg.go.dev/github.com/go-rod/rod#Element.MustEval)
  - [func (el *Element) MustFocus() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustFocus)
  - [func (el *Element) MustFrame() *Page](https://pkg.go.dev/github.com/go-rod/rod#Element.MustFrame)
  - [func (el *Element) MustGetXPath(optimized bool) string](https://pkg.go.dev/github.com/go-rod/rod#Element.MustGetXPath)
  - [func (el *Element) MustHTML() string](https://pkg.go.dev/github.com/go-rod/rod#Element.MustHTML)
  - [func (el *Element) MustHas(selector string) bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustHas)
  - [func (el *Element) MustHasR(selector, regex string) bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustHasR)
  - [func (el *Element) MustHasX(selector string) bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustHasX)
  - [func (el *Element) MustHover() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustHover)
  - [func (el *Element) MustInput(text string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustInput)
  - [func (el *Element) MustInputColor(color string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustInputColor)
  - [func (el *Element) MustInputTime(t time.Time) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustInputTime)
  - [func (el *Element) MustInteractable() bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustInteractable)
  - [func (el *Element) MustKeyActions() *KeyActions](https://pkg.go.dev/github.com/go-rod/rod#Element.MustKeyActions)
  - [func (el *Element) MustMatches(selector string) bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustMatches)
  - [func (el *Element) MustMoveMouseOut() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustMoveMouseOut)
  - [func (el *Element) MustNext() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustNext)
  - [func (el *Element) MustParent() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustParent)
  - [func (el *Element) MustParents(selector string) Elements](https://pkg.go.dev/github.com/go-rod/rod#Element.MustParents)
  - [func (el *Element) MustPrevious() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustPrevious)
  - [func (el *Element) MustProperty(name string) gson.JSON](https://pkg.go.dev/github.com/go-rod/rod#Element.MustProperty)
  - [func (el *Element) MustRelease()](https://pkg.go.dev/github.com/go-rod/rod#Element.MustRelease)
  - [func (el *Element) MustRemove()](https://pkg.go.dev/github.com/go-rod/rod#Element.MustRemove)
  - [func (el *Element) MustResource() [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Element.MustResource)
  - [func (el *Element) MustScreenshot(toFile ...string) [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Element.MustScreenshot)
  - [func (el *Element) MustScrollIntoView() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustScrollIntoView)
  - [func (el *Element) MustSelect(selectors ...string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustSelect)
  - [func (el *Element) MustSelectAllText() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustSelectAllText)
  - [func (el *Element) MustSelectText(regex string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustSelectText)
  - [func (el *Element) MustSetFiles(paths ...string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustSetFiles)
  - [func (el *Element) MustShadowRoot() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustShadowRoot)
  - [func (el *Element) MustShape() *proto.DOMGetContentQuadsResult](https://pkg.go.dev/github.com/go-rod/rod#Element.MustShape)
  - [func (el *Element) MustTap() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustTap)
  - [func (el *Element) MustText() string](https://pkg.go.dev/github.com/go-rod/rod#Element.MustText)
  - [func (el *Element) MustType(keys ...input.Key) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustType)
  - [func (el *Element) MustVisible() bool](https://pkg.go.dev/github.com/go-rod/rod#Element.MustVisible)
  - [func (el *Element) MustWait(js string, params ...interface{}) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustWait)
  - [func (el *Element) MustWaitEnabled() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustWaitEnabled)
  - [func (el *Element) MustWaitInteractable() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustWaitInteractable)
  - [func (el *Element) MustWaitInvisible() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustWaitInvisible)
  - [func (el *Element) MustWaitLoad() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustWaitLoad)
  - [func (el *Element) MustWaitStable() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustWaitStable)
  - [func (el *Element) MustWaitVisible() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustWaitVisible)
  - [func (el *Element) MustWaitWritable() *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.MustWaitWritable)
  - [func (el *Element) Next() (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Next)
  - [func (el *Element) Overlay(msg string) (removeOverlay func())](https://pkg.go.dev/github.com/go-rod/rod#Element.Overlay)
  - [func (el *Element) Page() *Page](https://pkg.go.dev/github.com/go-rod/rod#Element.Page)
  - [func (el *Element) Parent() (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Parent)
  - [func (el *Element) Parents(selector string) (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Parents)
  - [func (el *Element) Previous() (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Previous)
  - [func (el *Element) Property(name string) (gson.JSON, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Property)
  - [func (el *Element) Release() error](https://pkg.go.dev/github.com/go-rod/rod#Element.Release)
  - [func (el *Element) Remove() error](https://pkg.go.dev/github.com/go-rod/rod#Element.Remove)
  - [func (el *Element) Resource() ([\]byte, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Resource)
  - [func (el *Element) Screenshot(format proto.PageCaptureScreenshotFormat, quality int) ([\]byte, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Screenshot)
  - [func (el *Element) ScrollIntoView() error](https://pkg.go.dev/github.com/go-rod/rod#Element.ScrollIntoView)
  - [func (el *Element) Select(selectors [\]string, selected bool, t SelectorType) error](https://pkg.go.dev/github.com/go-rod/rod#Element.Select)
  - [func (el *Element) SelectAllText() error](https://pkg.go.dev/github.com/go-rod/rod#Element.SelectAllText)
  - [func (el *Element) SelectText(regex string) error](https://pkg.go.dev/github.com/go-rod/rod#Element.SelectText)
  - [func (el *Element) SetFiles(paths [\]string) error](https://pkg.go.dev/github.com/go-rod/rod#Element.SetFiles)
  - [func (el *Element) ShadowRoot() (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.ShadowRoot)
  - [func (el *Element) Shape() (*proto.DOMGetContentQuadsResult, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Shape)
  - [func (el *Element) Sleeper(sleeper func() utils.Sleeper) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.Sleeper)
  - [func (el *Element) String() string](https://pkg.go.dev/github.com/go-rod/rod#Element.String)
  - [func (el *Element) Tap() error](https://pkg.go.dev/github.com/go-rod/rod#Element.Tap)
  - [func (el *Element) Text() (string, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Text)
  - [func (el *Element) Timeout(d time.Duration) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.Timeout)
  - [func (el *Element) Type(keys ...input.Key) error](https://pkg.go.dev/github.com/go-rod/rod#Element.Type)
  - [func (el *Element) Visible() (bool, error)](https://pkg.go.dev/github.com/go-rod/rod#Element.Visible)
  - [func (el *Element) Wait(opts *EvalOptions) error](https://pkg.go.dev/github.com/go-rod/rod#Element.Wait)
  - [func (el *Element) WaitEnabled() error](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitEnabled)
  - [func (el *Element) WaitInteractable() (pt *proto.Point, err error)](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitInteractable)
  - [func (el *Element) WaitInvisible() error](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitInvisible)
  - [func (el *Element) WaitLoad() error](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitLoad)
  - [func (el *Element) WaitStable(d time.Duration) error](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitStable)
  - [func (el *Element) WaitStableRAF() error](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitStableRAF)
  - [func (el *Element) WaitVisible() error](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitVisible)
  - [func (el *Element) WaitWritable() error](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitWritable)
  - [func (el *Element) WithCancel() (*Element, func())](https://pkg.go.dev/github.com/go-rod/rod#Element.WithCancel)
  - [func (el *Element) WithPanic(fail func(interface{})) *Element](https://pkg.go.dev/github.com/go-rod/rod#Element.WithPanic)
- [type ElementNotFoundError](https://pkg.go.dev/github.com/go-rod/rod#ElementNotFoundError)
- - [func (e *ElementNotFoundError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#ElementNotFoundError.Error)
- [type Elements](https://pkg.go.dev/github.com/go-rod/rod#Elements)
- - [func (els Elements) Empty() bool](https://pkg.go.dev/github.com/go-rod/rod#Elements.Empty)
  - [func (els Elements) First() *Element](https://pkg.go.dev/github.com/go-rod/rod#Elements.First)
  - [func (els Elements) Last() *Element](https://pkg.go.dev/github.com/go-rod/rod#Elements.Last)
- [type EvalError](https://pkg.go.dev/github.com/go-rod/rod#EvalError)
- - [func (e *EvalError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#EvalError.Error)
  - [func (e *EvalError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#EvalError.Is)
- [type EvalOptions](https://pkg.go.dev/github.com/go-rod/rod#EvalOptions)
- - [func Eval(js string, args ...interface{}) *EvalOptions](https://pkg.go.dev/github.com/go-rod/rod#Eval)
- - [func (e *EvalOptions) ByObject() *EvalOptions](https://pkg.go.dev/github.com/go-rod/rod#EvalOptions.ByObject)
  - [func (e *EvalOptions) ByPromise() *EvalOptions](https://pkg.go.dev/github.com/go-rod/rod#EvalOptions.ByPromise)
  - [func (e *EvalOptions) ByUser() *EvalOptions](https://pkg.go.dev/github.com/go-rod/rod#EvalOptions.ByUser)
  - [func (e *EvalOptions) String() string](https://pkg.go.dev/github.com/go-rod/rod#EvalOptions.String)
  - [func (e *EvalOptions) This(obj *proto.RuntimeRemoteObject) *EvalOptions](https://pkg.go.dev/github.com/go-rod/rod#EvalOptions.This)
- [type ExpectElementError](https://pkg.go.dev/github.com/go-rod/rod#ExpectElementError)
- - [func (e *ExpectElementError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#ExpectElementError.Error)
  - [func (e *ExpectElementError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#ExpectElementError.Is)
- [type ExpectElementsError](https://pkg.go.dev/github.com/go-rod/rod#ExpectElementsError)
- - [func (e *ExpectElementsError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#ExpectElementsError.Error)
  - [func (e *ExpectElementsError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#ExpectElementsError.Is)
- [type Hijack](https://pkg.go.dev/github.com/go-rod/rod#Hijack)
- - [func (h *Hijack) ContinueRequest(cq *proto.FetchContinueRequest)](https://pkg.go.dev/github.com/go-rod/rod#Hijack.ContinueRequest)
  - [func (h *Hijack) LoadResponse(client *http.Client, loadBody bool) error](https://pkg.go.dev/github.com/go-rod/rod#Hijack.LoadResponse)
  - [func (h *Hijack) MustLoadResponse()](https://pkg.go.dev/github.com/go-rod/rod#Hijack.MustLoadResponse)
- [type HijackRequest](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest)
- - [func (ctx *HijackRequest) Body() string](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.Body)
  - [func (ctx *HijackRequest) Header(key string) string](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.Header)
  - [func (ctx *HijackRequest) Headers() proto.NetworkHeaders](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.Headers)
  - [func (ctx *HijackRequest) IsNavigation() bool](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.IsNavigation)
  - [func (ctx *HijackRequest) JSONBody() gson.JSON](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.JSONBody)
  - [func (ctx *HijackRequest) Method() string](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.Method)
  - [func (ctx *HijackRequest) Req() *http.Request](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.Req)
  - [func (ctx *HijackRequest) SetBody(obj interface{}) *HijackRequest](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.SetBody)
  - [func (ctx *HijackRequest) SetContext(c context.Context) *HijackRequest](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.SetContext)
  - [func (ctx *HijackRequest) Type() proto.NetworkResourceType](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.Type)
  - [func (ctx *HijackRequest) URL() *url.URL](https://pkg.go.dev/github.com/go-rod/rod#HijackRequest.URL)
- [type HijackResponse](https://pkg.go.dev/github.com/go-rod/rod#HijackResponse)
- - [func (ctx *HijackResponse) Body() string](https://pkg.go.dev/github.com/go-rod/rod#HijackResponse.Body)
  - [func (ctx *HijackResponse) Fail(reason proto.NetworkErrorReason) *HijackResponse](https://pkg.go.dev/github.com/go-rod/rod#HijackResponse.Fail)
  - [func (ctx *HijackResponse) Headers() http.Header](https://pkg.go.dev/github.com/go-rod/rod#HijackResponse.Headers)
  - [func (ctx *HijackResponse) Payload() *proto.FetchFulfillRequest](https://pkg.go.dev/github.com/go-rod/rod#HijackResponse.Payload)
  - [func (ctx *HijackResponse) SetBody(obj interface{}) *HijackResponse](https://pkg.go.dev/github.com/go-rod/rod#HijackResponse.SetBody)
  - [func (ctx *HijackResponse) SetHeader(pairs ...string) *HijackResponse](https://pkg.go.dev/github.com/go-rod/rod#HijackResponse.SetHeader)
- [type HijackRouter](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter)
- - [func (r *HijackRouter) Add(pattern string, resourceType proto.NetworkResourceType, handler func(*Hijack)) error](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Add)
  - [func (r *HijackRouter) MustAdd(pattern string, handler func(*Hijack)) *HijackRouter](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.MustAdd)
  - [func (r *HijackRouter) MustRemove(pattern string) *HijackRouter](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.MustRemove)
  - [func (r *HijackRouter) MustStop()](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.MustStop)
  - [func (r *HijackRouter) Remove(pattern string) error](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Remove)
  - [func (r *HijackRouter) Run()](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Run)
  - [func (r *HijackRouter) Stop() error](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Stop)
- [type InvisibleShapeError](https://pkg.go.dev/github.com/go-rod/rod#InvisibleShapeError)
- - [func (e *InvisibleShapeError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#InvisibleShapeError.Error)
  - [func (e *InvisibleShapeError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#InvisibleShapeError.Is)
  - [func (e *InvisibleShapeError) Unwrap() error](https://pkg.go.dev/github.com/go-rod/rod#InvisibleShapeError.Unwrap)
- [type KeyAction](https://pkg.go.dev/github.com/go-rod/rod#KeyAction)
- [type KeyActionType](https://pkg.go.dev/github.com/go-rod/rod#KeyActionType)
- [type KeyActions](https://pkg.go.dev/github.com/go-rod/rod#KeyActions)
- - [func (ka *KeyActions) Do() (err error)](https://pkg.go.dev/github.com/go-rod/rod#KeyActions.Do)
  - [func (ka *KeyActions) MustDo()](https://pkg.go.dev/github.com/go-rod/rod#KeyActions.MustDo)
  - [func (ka *KeyActions) Press(keys ...input.Key) *KeyActions](https://pkg.go.dev/github.com/go-rod/rod#KeyActions.Press)
  - [func (ka *KeyActions) Release(keys ...input.Key) *KeyActions](https://pkg.go.dev/github.com/go-rod/rod#KeyActions.Release)
  - [func (ka *KeyActions) Type(keys ...input.Key) *KeyActions](https://pkg.go.dev/github.com/go-rod/rod#KeyActions.Type)
- [type Keyboard](https://pkg.go.dev/github.com/go-rod/rod#Keyboard)
- - [func (k *Keyboard) MustType(key ...input.Key) *Keyboard](https://pkg.go.dev/github.com/go-rod/rod#Keyboard.MustType)
  - [func (k *Keyboard) Press(key input.Key) error](https://pkg.go.dev/github.com/go-rod/rod#Keyboard.Press)
  - [func (k *Keyboard) Release(key input.Key) error](https://pkg.go.dev/github.com/go-rod/rod#Keyboard.Release)
  - [func (k *Keyboard) Type(keys ...input.Key) (err error)](https://pkg.go.dev/github.com/go-rod/rod#Keyboard.Type)
- [type Message](https://pkg.go.dev/github.com/go-rod/rod#Message)
- - [func (msg *Message) Load(e proto.Event) bool](https://pkg.go.dev/github.com/go-rod/rod#Message.Load)
- [type Mouse](https://pkg.go.dev/github.com/go-rod/rod#Mouse)
- - [func (m *Mouse) Click(button proto.InputMouseButton, clickCount int) error](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Click)
  - [func (m *Mouse) Down(button proto.InputMouseButton, clickCount int) error](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Down)
  - [func (m *Mouse) MoveAlong(guide func() (proto.Point, bool)) error](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MoveAlong)
  - [func (m *Mouse) MoveLinear(to proto.Point, steps int) error](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MoveLinear)
  - [func (m *Mouse) MoveTo(p proto.Point) error](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MoveTo)
  - [func (m *Mouse) MustClick(button proto.InputMouseButton) *Mouse](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MustClick)
  - [func (m *Mouse) MustDown(button proto.InputMouseButton) *Mouse](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MustDown)
  - [func (m *Mouse) MustMoveTo(x, y float64) *Mouse](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MustMoveTo)
  - [func (m *Mouse) MustScroll(x, y float64) *Mouse](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MustScroll)
  - [func (m *Mouse) MustUp(button proto.InputMouseButton) *Mouse](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MustUp)
  - [func (m *Mouse) Position() proto.Point](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Position)
  - [func (m *Mouse) Scroll(offsetX, offsetY float64, steps int) error](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Scroll)
  - [func (m *Mouse) Up(button proto.InputMouseButton, clickCount int) error](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Up)
- [type NavigationError](https://pkg.go.dev/github.com/go-rod/rod#NavigationError)
- - [func (e *NavigationError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#NavigationError.Error)
  - [func (e *NavigationError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#NavigationError.Is)
- [type NoPointerEventsError](https://pkg.go.dev/github.com/go-rod/rod#NoPointerEventsError)
- - [func (e *NoPointerEventsError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#NoPointerEventsError.Error)
  - [func (e *NoPointerEventsError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#NoPointerEventsError.Is)
  - [func (e *NoPointerEventsError) Unwrap() error](https://pkg.go.dev/github.com/go-rod/rod#NoPointerEventsError.Unwrap)
- [type NoShadowRootError](https://pkg.go.dev/github.com/go-rod/rod#NoShadowRootError)
- - [func (e *NoShadowRootError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#NoShadowRootError.Error)
  - [func (e *NoShadowRootError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#NoShadowRootError.Is)
- [type NotInteractableError](https://pkg.go.dev/github.com/go-rod/rod#NotInteractableError)
- - [func (e *NotInteractableError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#NotInteractableError.Error)
- [type ObjectNotFoundError](https://pkg.go.dev/github.com/go-rod/rod#ObjectNotFoundError)
- - [func (e *ObjectNotFoundError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#ObjectNotFoundError.Error)
  - [func (e *ObjectNotFoundError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#ObjectNotFoundError.Is)
- [type Page](https://pkg.go.dev/github.com/go-rod/rod#Page)
- - [func (p *Page) Activate() (*Page, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Activate)
  - [func (p *Page) AddScriptTag(url, content string) error](https://pkg.go.dev/github.com/go-rod/rod#Page.AddScriptTag)
  - [func (p *Page) AddStyleTag(url, content string) error](https://pkg.go.dev/github.com/go-rod/rod#Page.AddStyleTag)
  - [func (p *Page) Browser() *Browser](https://pkg.go.dev/github.com/go-rod/rod#Page.Browser)
  - [func (p *Page) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res [\]byte, err error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Call)
  - [func (p *Page) CancelTimeout() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.CancelTimeout)
  - [func (p *Page) CaptureDOMSnapshot() (domSnapshot *proto.DOMSnapshotCaptureSnapshotResult, err error)](https://pkg.go.dev/github.com/go-rod/rod#Page.CaptureDOMSnapshot)
  - [func (p *Page) Close() error](https://pkg.go.dev/github.com/go-rod/rod#Page.Close)
  - [func (p *Page) Context(ctx context.Context) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.Context)
  - [func (p *Page) Cookies(urls [\]string) ([]*proto.NetworkCookie, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Cookies)
  - [func (p *Page) DisableDomain(method proto.Request) (restore func())](https://pkg.go.dev/github.com/go-rod/rod#Page.DisableDomain)
  - [func (p *Page) EachEvent(callbacks ...interface{}) (wait func())](https://pkg.go.dev/github.com/go-rod/rod#Page.EachEvent)
  - [func (p *Page) Element(selector string) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Element)
  - [func (p *Page) ElementByJS(opts *EvalOptions) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementByJS)
  - [func (p *Page) ElementFromNode(node *proto.DOMNode) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementFromNode)
  - [func (p *Page) ElementFromObject(obj *proto.RuntimeRemoteObject) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementFromObject)
  - [func (p *Page) ElementFromPoint(x, y int) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementFromPoint)
  - [func (p *Page) ElementR(selector, jsRegex string) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementR)
  - [func (p *Page) ElementX(xPath string) (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementX)
  - [func (p *Page) Elements(selector string) (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Elements)
  - [func (p *Page) ElementsByJS(opts *EvalOptions) (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementsByJS)
  - [func (p *Page) ElementsX(xpath string) (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementsX)
  - [func (p *Page) Emulate(device devices.Device) error](https://pkg.go.dev/github.com/go-rod/rod#Page.Emulate)
  - [func (p *Page) EnableDomain(method proto.Request) (restore func())](https://pkg.go.dev/github.com/go-rod/rod#Page.EnableDomain)
  - [func (p *Page) Eval(js string, args ...interface{}) (*proto.RuntimeRemoteObject, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Eval)
  - [func (p *Page) EvalOnNewDocument(js string) (remove func() error, err error)](https://pkg.go.dev/github.com/go-rod/rod#Page.EvalOnNewDocument)
  - [func (p *Page) Evaluate(opts *EvalOptions) (res *proto.RuntimeRemoteObject, err error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Evaluate)
  - [func (p *Page) Event() <-chan *Message](https://pkg.go.dev/github.com/go-rod/rod#Page.Event)
  - [func (p *Page) Expose(name string, fn func(gson.JSON) (interface{}, error)) (stop func() error, err error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Expose)
  - [func (p *Page) ExposeHelpers(list ...*js.Function)](https://pkg.go.dev/github.com/go-rod/rod#Page.ExposeHelpers)
  - [func (p *Page) GetContext() context.Context](https://pkg.go.dev/github.com/go-rod/rod#Page.GetContext)
  - [func (p *Page) GetNavigationHistory() (*proto.PageGetNavigationHistoryResult, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.GetNavigationHistory)
  - [func (p *Page) GetResource(url string) ([\]byte, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.GetResource)
  - [func (p *Page) GetSessionID() proto.TargetSessionID](https://pkg.go.dev/github.com/go-rod/rod#Page.GetSessionID)
  - [func (p *Page) GetWindow() (*proto.BrowserBounds, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.GetWindow)
  - [func (p *Page) HTML() (string, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.HTML)
  - [func (p *Page) HandleDialog() (wait func() *proto.PageJavascriptDialogOpening, ...)](https://pkg.go.dev/github.com/go-rod/rod#Page.HandleDialog)
  - [func (p *Page) HandleFileDialog() (func([\]string) error, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.HandleFileDialog)
  - [func (p *Page) Has(selector string) (bool, *Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Has)
  - [func (p *Page) HasR(selector, jsRegex string) (bool, *Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.HasR)
  - [func (p *Page) HasX(selector string) (bool, *Element, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.HasX)
  - [func (p *Page) HijackRequests() *HijackRouter](https://pkg.go.dev/github.com/go-rod/rod#Page.HijackRequests)
  - [func (p *Page) Info() (*proto.TargetTargetInfo, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Info)
  - [func (p *Page) InsertText(text string) error](https://pkg.go.dev/github.com/go-rod/rod#Page.InsertText)
  - [func (p *Page) IsIframe() bool](https://pkg.go.dev/github.com/go-rod/rod#Page.IsIframe)
  - [func (p *Page) KeyActions() *KeyActions](https://pkg.go.dev/github.com/go-rod/rod#Page.KeyActions)
  - [func (p *Page) LoadState(method proto.Request) (has bool)](https://pkg.go.dev/github.com/go-rod/rod#Page.LoadState)
  - [func (p *Page) MustActivate() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustActivate)
  - [func (p *Page) MustAddScriptTag(url string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustAddScriptTag)
  - [func (p *Page) MustAddStyleTag(url string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustAddStyleTag)
  - [func (p *Page) MustCaptureDOMSnapshot() (domSnapshot *proto.DOMSnapshotCaptureSnapshotResult)](https://pkg.go.dev/github.com/go-rod/rod#Page.MustCaptureDOMSnapshot)
  - [func (p *Page) MustClose()](https://pkg.go.dev/github.com/go-rod/rod#Page.MustClose)
  - [func (p *Page) MustCookies(urls ...string) [\]*proto.NetworkCookie](https://pkg.go.dev/github.com/go-rod/rod#Page.MustCookies)
  - [func (p *Page) MustElement(selector string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElement)
  - [func (p *Page) MustElementByJS(js string, params ...interface{}) *Element](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElementByJS)
  - [func (p *Page) MustElementFromNode(node *proto.DOMNode) *Element](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElementFromNode)
  - [func (p *Page) MustElementFromPoint(left, top int) *Element](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElementFromPoint)
  - [func (p *Page) MustElementR(selector, jsRegex string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElementR)
  - [func (p *Page) MustElementX(xPath string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElementX)
  - [func (p *Page) MustElements(selector string) Elements](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElements)
  - [func (p *Page) MustElementsByJS(js string, params ...interface{}) Elements](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElementsByJS)
  - [func (p *Page) MustElementsX(xpath string) Elements](https://pkg.go.dev/github.com/go-rod/rod#Page.MustElementsX)
  - [func (p *Page) MustEmulate(device devices.Device) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustEmulate)
  - [func (p *Page) MustEval(js string, params ...interface{}) gson.JSON](https://pkg.go.dev/github.com/go-rod/rod#Page.MustEval)
  - [func (p *Page) MustEvalOnNewDocument(js string)](https://pkg.go.dev/github.com/go-rod/rod#Page.MustEvalOnNewDocument)
  - [func (p *Page) MustEvaluate(opts *EvalOptions) *proto.RuntimeRemoteObject](https://pkg.go.dev/github.com/go-rod/rod#Page.MustEvaluate)
  - [func (p *Page) MustExpose(name string, fn func(gson.JSON) (interface{}, error)) (stop func())](https://pkg.go.dev/github.com/go-rod/rod#Page.MustExpose)
  - [func (p *Page) MustGetWindow() *proto.BrowserBounds](https://pkg.go.dev/github.com/go-rod/rod#Page.MustGetWindow)
  - [func (p *Page) MustHTML() string](https://pkg.go.dev/github.com/go-rod/rod#Page.MustHTML)
  - [func (p *Page) MustHandleDialog() (wait func() *proto.PageJavascriptDialogOpening, handle func(bool, string))](https://pkg.go.dev/github.com/go-rod/rod#Page.MustHandleDialog)
  - [func (p *Page) MustHandleFileDialog() func(...string)](https://pkg.go.dev/github.com/go-rod/rod#Page.MustHandleFileDialog)
  - [func (p *Page) MustHas(selector string) bool](https://pkg.go.dev/github.com/go-rod/rod#Page.MustHas)
  - [func (p *Page) MustHasR(selector, regex string) bool](https://pkg.go.dev/github.com/go-rod/rod#Page.MustHasR)
  - [func (p *Page) MustHasX(selector string) bool](https://pkg.go.dev/github.com/go-rod/rod#Page.MustHasX)
  - [func (p *Page) MustInfo() *proto.TargetTargetInfo](https://pkg.go.dev/github.com/go-rod/rod#Page.MustInfo)
  - [func (p *Page) MustInsertText(text string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustInsertText)
  - [func (p *Page) MustNavigate(url string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustNavigate)
  - [func (p *Page) MustNavigateBack() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustNavigateBack)
  - [func (p *Page) MustNavigateForward() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustNavigateForward)
  - [func (p *Page) MustObjectToJSON(obj *proto.RuntimeRemoteObject) gson.JSON](https://pkg.go.dev/github.com/go-rod/rod#Page.MustObjectToJSON)
  - [func (p *Page) MustObjectsToJSON(list [\]*proto.RuntimeRemoteObject) gson.JSON](https://pkg.go.dev/github.com/go-rod/rod#Page.MustObjectsToJSON)
  - [func (p *Page) MustPDF(toFile ...string) [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Page.MustPDF)
  - [func (p *Page) MustRelease(obj *proto.RuntimeRemoteObject) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustRelease)
  - [func (p *Page) MustReload() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustReload)
  - [func (p *Page) MustResetNavigationHistory() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustResetNavigationHistory)
  - [func (p *Page) MustScreenshot(toFile ...string) [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Page.MustScreenshot)
  - [func (p *Page) MustScreenshotFullPage(toFile ...string) [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Page.MustScreenshotFullPage)
  - [func (p *Page) MustScrollScreenshot(toFile ...string) [\]byte](https://pkg.go.dev/github.com/go-rod/rod#Page.MustScrollScreenshot)
  - [func (p *Page) MustSearch(query string) *Element](https://pkg.go.dev/github.com/go-rod/rod#Page.MustSearch)
  - [func (p *Page) MustSetBlockedURLs(urls ...string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustSetBlockedURLs)
  - [func (p *Page) MustSetCookies(cookies ...*proto.NetworkCookieParam) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustSetCookies)
  - [func (p *Page) MustSetDocumentContent(html string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustSetDocumentContent)
  - [func (p *Page) MustSetExtraHeaders(dict ...string) (cleanup func())](https://pkg.go.dev/github.com/go-rod/rod#Page.MustSetExtraHeaders)
  - [func (p *Page) MustSetUserAgent(req *proto.NetworkSetUserAgentOverride) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustSetUserAgent)
  - [func (p *Page) MustSetViewport(width, height int, deviceScaleFactor float64, mobile bool) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustSetViewport)
  - [func (p *Page) MustSetWindow(left, top, width, height int) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustSetWindow)
  - [func (p *Page) MustStopLoading() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustStopLoading)
  - [func (p *Page) MustTriggerFavicon() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustTriggerFavicon)
  - [func (p *Page) MustWait(js string, params ...interface{}) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWait)
  - [func (p *Page) MustWaitDOMStable() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWaitDOMStable)
  - [func (p *Page) MustWaitElementsMoreThan(selector string, num int) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWaitElementsMoreThan)
  - [func (p *Page) MustWaitIdle() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWaitIdle)
  - [func (p *Page) MustWaitLoad() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWaitLoad)
  - [func (p *Page) MustWaitNavigation() func()](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWaitNavigation)
  - [func (p *Page) MustWaitOpen() (wait func() (newPage *Page))](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWaitOpen)
  - [func (p *Page) MustWaitRequestIdle(excludes ...string) (wait func())](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWaitRequestIdle)
  - [func (p *Page) MustWaitStable() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWaitStable)
  - [func (p *Page) MustWindowFullscreen() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWindowFullscreen)
  - [func (p *Page) MustWindowMaximize() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWindowMaximize)
  - [func (p *Page) MustWindowMinimize() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWindowMinimize)
  - [func (p *Page) MustWindowNormal() *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.MustWindowNormal)
  - [func (p *Page) Navigate(url string) error](https://pkg.go.dev/github.com/go-rod/rod#Page.Navigate)
  - [func (p *Page) NavigateBack() error](https://pkg.go.dev/github.com/go-rod/rod#Page.NavigateBack)
  - [func (p *Page) NavigateForward() error](https://pkg.go.dev/github.com/go-rod/rod#Page.NavigateForward)
  - [func (p *Page) ObjectToJSON(obj *proto.RuntimeRemoteObject) (gson.JSON, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ObjectToJSON)
  - [func (p *Page) Overlay(left, top, width, height float64, msg string) (remove func())](https://pkg.go.dev/github.com/go-rod/rod#Page.Overlay)
  - [func (p *Page) PDF(req *proto.PagePrintToPDF) (*StreamReader, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.PDF)
  - [func (p *Page) Race() *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#Page.Race)
  - [func (p *Page) Release(obj *proto.RuntimeRemoteObject) error](https://pkg.go.dev/github.com/go-rod/rod#Page.Release)
  - [func (p *Page) Reload() error](https://pkg.go.dev/github.com/go-rod/rod#Page.Reload)
  - [func (p *Page) ResetNavigationHistory() error](https://pkg.go.dev/github.com/go-rod/rod#Page.ResetNavigationHistory)
  - [func (p *Page) Screenshot(fullPage bool, req *proto.PageCaptureScreenshot) ([\]byte, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Screenshot)
  - [func (p *Page) ScrollScreenshot(opt *ScrollScreenshotOptions) ([\]byte, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.ScrollScreenshot)
  - [func (p *Page) Search(query string) (*SearchResult, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.Search)
  - [func (p *Page) SetBlockedURLs(urls [\]string) error](https://pkg.go.dev/github.com/go-rod/rod#Page.SetBlockedURLs)
  - [func (p *Page) SetCookies(cookies [\]*proto.NetworkCookieParam) error](https://pkg.go.dev/github.com/go-rod/rod#Page.SetCookies)
  - [func (p *Page) SetDocumentContent(html string) error](https://pkg.go.dev/github.com/go-rod/rod#Page.SetDocumentContent)
  - [func (p *Page) SetExtraHeaders(dict [\]string) (func(), error)](https://pkg.go.dev/github.com/go-rod/rod#Page.SetExtraHeaders)
  - [func (p *Page) SetUserAgent(req *proto.NetworkSetUserAgentOverride) error](https://pkg.go.dev/github.com/go-rod/rod#Page.SetUserAgent)
  - [func (p *Page) SetViewport(params *proto.EmulationSetDeviceMetricsOverride) error](https://pkg.go.dev/github.com/go-rod/rod#Page.SetViewport)
  - [func (p *Page) SetWindow(bounds *proto.BrowserBounds) error](https://pkg.go.dev/github.com/go-rod/rod#Page.SetWindow)
  - [func (p *Page) Sleeper(sleeper func() utils.Sleeper) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.Sleeper)
  - [func (p *Page) StopLoading() error](https://pkg.go.dev/github.com/go-rod/rod#Page.StopLoading)
  - [func (p *Page) String() string](https://pkg.go.dev/github.com/go-rod/rod#Page.String)
  - [func (p *Page) Timeout(d time.Duration) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.Timeout)
  - [func (p *Page) TriggerFavicon() error](https://pkg.go.dev/github.com/go-rod/rod#Page.TriggerFavicon)
  - [func (p *Page) Wait(opts *EvalOptions) error](https://pkg.go.dev/github.com/go-rod/rod#Page.Wait)
  - [func (p *Page) WaitDOMStable(d time.Duration, diff float64) error](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitDOMStable)
  - [func (p *Page) WaitElementsMoreThan(selector string, num int) error](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitElementsMoreThan)
  - [func (p *Page) WaitEvent(e proto.Event) (wait func())](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitEvent)
  - [func (p *Page) WaitIdle(timeout time.Duration) (err error)](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitIdle)
  - [func (p *Page) WaitLoad() error](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitLoad)
  - [func (p *Page) WaitNavigation(name proto.PageLifecycleEventName) func()](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitNavigation)
  - [func (p *Page) WaitOpen() func() (*Page, error)](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitOpen)
  - [func (p *Page) WaitRepaint() error](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitRepaint)
  - [func (p *Page) WaitRequestIdle(d time.Duration, includes, excludes [\]string, ...) func()](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitRequestIdle)
  - [func (p *Page) WaitStable(d time.Duration) error](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitStable)
  - [func (p *Page) WithCancel() (*Page, func())](https://pkg.go.dev/github.com/go-rod/rod#Page.WithCancel)
  - [func (p *Page) WithPanic(fail func(interface{})) *Page](https://pkg.go.dev/github.com/go-rod/rod#Page.WithPanic)
- [type PageCloseCanceledError](https://pkg.go.dev/github.com/go-rod/rod#PageCloseCanceledError)
- - [func (e *PageCloseCanceledError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#PageCloseCanceledError.Error)
- [type PageNotFoundError](https://pkg.go.dev/github.com/go-rod/rod#PageNotFoundError)
- - [func (e *PageNotFoundError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#PageNotFoundError.Error)
- [type Pages](https://pkg.go.dev/github.com/go-rod/rod#Pages)
- - [func (ps Pages) Empty() bool](https://pkg.go.dev/github.com/go-rod/rod#Pages.Empty)
  - [func (ps Pages) Find(selector string) (*Page, error)](https://pkg.go.dev/github.com/go-rod/rod#Pages.Find)
  - [func (ps Pages) FindByURL(jsRegex string) (*Page, error)](https://pkg.go.dev/github.com/go-rod/rod#Pages.FindByURL)
  - [func (ps Pages) First() *Page](https://pkg.go.dev/github.com/go-rod/rod#Pages.First)
  - [func (ps Pages) Last() *Page](https://pkg.go.dev/github.com/go-rod/rod#Pages.Last)
  - [func (ps Pages) MustFind(selector string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Pages.MustFind)
  - [func (ps Pages) MustFindByURL(regex string) *Page](https://pkg.go.dev/github.com/go-rod/rod#Pages.MustFindByURL)
- [type Pool](https://pkg.go.dev/github.com/go-rod/rod#Pool)
- - [func NewBrowserPool(limit int) Pool[Browser\]](https://pkg.go.dev/github.com/go-rod/rod#NewBrowserPool)
  - [func NewPagePool(limit int) Pool[Page\]](https://pkg.go.dev/github.com/go-rod/rod#NewPagePool)
  - [func NewPool[T any](limit int) Pool[T\]](https://pkg.go.dev/github.com/go-rod/rod#NewPool)
- - [func (p Pool[T\]) Cleanup(iteratee func(*T))](https://pkg.go.dev/github.com/go-rod/rod#Pool.Cleanup)
  - [func (p Pool[T\]) Get(create func() (*T, error)) (elem *T, err error)](https://pkg.go.dev/github.com/go-rod/rod#Pool.Get)
  - [func (p Pool[T\]) MustGet(create func() *T) *T](https://pkg.go.dev/github.com/go-rod/rod#Pool.MustGet)
  - [func (p Pool[T\]) Put(elem *T)](https://pkg.go.dev/github.com/go-rod/rod#Pool.Put)
- [type RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext)
- - [func (rc *RaceContext) Do() (*Element, error)](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.Do)
  - [func (rc *RaceContext) Element(selector string) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.Element)
  - [func (rc *RaceContext) ElementByJS(opts *EvalOptions) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.ElementByJS)
  - [func (rc *RaceContext) ElementFunc(fn func(*Page) (*Element, error)) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.ElementFunc)
  - [func (rc *RaceContext) ElementR(selector, regex string) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.ElementR)
  - [func (rc *RaceContext) ElementX(selector string) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.ElementX)
  - [func (rc *RaceContext) Handle(callback func(*Element) error) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.Handle)
  - [func (rc *RaceContext) MustDo() *Element](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.MustDo)
  - [func (rc *RaceContext) MustElementByJS(js string, params [\]interface{}) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.MustElementByJS)
  - [func (rc *RaceContext) MustHandle(callback func(*Element)) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.MustHandle)
  - [func (rc *RaceContext) Search(query string) *RaceContext](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.Search)
- [type ScrollScreenshotOptions](https://pkg.go.dev/github.com/go-rod/rod#ScrollScreenshotOptions)
- [type SearchResult](https://pkg.go.dev/github.com/go-rod/rod#SearchResult)
- - [func (s *SearchResult) All() (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#SearchResult.All)
  - [func (s *SearchResult) Get(i, l int) (Elements, error)](https://pkg.go.dev/github.com/go-rod/rod#SearchResult.Get)
  - [func (s *SearchResult) Release()](https://pkg.go.dev/github.com/go-rod/rod#SearchResult.Release)
- [type SelectorType](https://pkg.go.dev/github.com/go-rod/rod#SelectorType)
- [type StreamReader](https://pkg.go.dev/github.com/go-rod/rod#StreamReader)
- - [func NewStreamReader(c proto.Client, h proto.IOStreamHandle) *StreamReader](https://pkg.go.dev/github.com/go-rod/rod#NewStreamReader)
- - [func (sr *StreamReader) Close() error](https://pkg.go.dev/github.com/go-rod/rod#StreamReader.Close)
  - [func (sr *StreamReader) Read(p [\]byte) (n int, err error)](https://pkg.go.dev/github.com/go-rod/rod#StreamReader.Read)
- [type Touch](https://pkg.go.dev/github.com/go-rod/rod#Touch)
- - [func (t *Touch) Cancel() error](https://pkg.go.dev/github.com/go-rod/rod#Touch.Cancel)
  - [func (t *Touch) End() error](https://pkg.go.dev/github.com/go-rod/rod#Touch.End)
  - [func (t *Touch) Move(points ...*proto.InputTouchPoint) error](https://pkg.go.dev/github.com/go-rod/rod#Touch.Move)
  - [func (t *Touch) MustCancel() *Touch](https://pkg.go.dev/github.com/go-rod/rod#Touch.MustCancel)
  - [func (t *Touch) MustEnd() *Touch](https://pkg.go.dev/github.com/go-rod/rod#Touch.MustEnd)
  - [func (t *Touch) MustMove(points ...*proto.InputTouchPoint) *Touch](https://pkg.go.dev/github.com/go-rod/rod#Touch.MustMove)
  - [func (t *Touch) MustStart(points ...*proto.InputTouchPoint) *Touch](https://pkg.go.dev/github.com/go-rod/rod#Touch.MustStart)
  - [func (t *Touch) MustTap(x, y float64) *Touch](https://pkg.go.dev/github.com/go-rod/rod#Touch.MustTap)
  - [func (t *Touch) Start(points ...*proto.InputTouchPoint) error](https://pkg.go.dev/github.com/go-rod/rod#Touch.Start)
  - [func (t *Touch) Tap(x, y float64) error](https://pkg.go.dev/github.com/go-rod/rod#Touch.Tap)
- [type TraceType](https://pkg.go.dev/github.com/go-rod/rod#TraceType)
- - [func (t TraceType) String() string](https://pkg.go.dev/github.com/go-rod/rod#TraceType.String)
- [type TryError](https://pkg.go.dev/github.com/go-rod/rod#TryError)
- - [func (e *TryError) Error() string](https://pkg.go.dev/github.com/go-rod/rod#TryError.Error)
  - [func (e *TryError) Is(err error) bool](https://pkg.go.dev/github.com/go-rod/rod#TryError.Is)
  - [func (e *TryError) Unwrap() error](https://pkg.go.dev/github.com/go-rod/rod#TryError.Unwrap)

#### Examples 

- [Package (Basic)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Basic)
- [Package (Context_and_EachEvent)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Context_and_EachEvent)
- [Package (Context_and_timeout)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Context_and_timeout)
- [Package (Customize_browser_launch)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Customize_browser_launch)
- [Package (Customize_retry_strategy)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Customize_retry_strategy)
- [Package (Direct_cdp)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Direct_cdp)
- [Package (Disable_headless_to_debug)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Disable_headless_to_debug)
- [Package (Download_file)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Download_file)
- [Package (Error_handling)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Error_handling)
- [Package (Eval_reuse_remote_object)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Eval_reuse_remote_object)
- [Package (Handle_events)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Handle_events)
- [Package (Hijack_requests)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Hijack_requests)
- [Package (Load_extension)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Load_extension)
- [Package (Log_cdp_traffic)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Log_cdp_traffic)
- [Package (Page_pdf)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Page_pdf)
- [Package (Page_screenshot)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Page_screenshot)
- [Package (Page_scroll_screenshot)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Page_scroll_screenshot)
- [Package (Race_selectors)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Race_selectors)
- [Package (Search)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Search)
- [Package (States)](https://pkg.go.dev/github.com/go-rod/rod#example-package-States)
- [Package (Wait_for_animation)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Wait_for_animation)
- [Package (Wait_for_request)](https://pkg.go.dev/github.com/go-rod/rod#example-package-Wait_for_request)
- [Browser (Pool)](https://pkg.go.dev/github.com/go-rod/rod#example-Browser-Pool)
- [Page (Pool)](https://pkg.go.dev/github.com/go-rod/rod#example-Page-Pool)

### Constants 

This section is empty.

### Variables 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/utils.go#L67)

``` go
var DefaultLogger = log.New(os.Stdout, "[rod] ", log.LstdFlags)
```

DefaultLogger for rod.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/utils.go#L76)

``` go
var DefaultSleeper = func() utils.Sleeper {
	return utils.BackoffSleeper(100*time.Millisecond, time.Second, nil)
}
```

DefaultSleeper generates the default sleeper for retry, it uses backoff to grow the interval. The growth looks like:

```
A(0) = 100ms, A(n) = A(n-1) * random[1.9, 2.1), A(n) < 1s
```

Why the default is not RequestAnimationFrame or DOM change events is because of if a retry never ends it can easily flood the program. But you can always easily config it into what you want.

### Functions 

#### func NotFoundSleeper <- 0.88.9

``` go
func NotFoundSleeper() utils.Sleeper
```

NotFoundSleeper returns ErrElementNotFound on the first call.

#### func Try <- 0.46.0

``` go
func Try(fn func()) (err error)
```

Try try fn with recover, return the panic as rod.ErrTry.

### Types 

#### type Browser 

``` go
type Browser struct {
	// BrowserContextID is the id for incognito window
	BrowserContextID proto.BrowserBrowserContextID
	// contains filtered or unexported fields
}
```

Browser represents the browser. It doesn't depends on file system, it should work with remote browser seamlessly. To check the env var you can use to quickly enable options from CLI, check here: https://pkg.go.dev/github.com/go-rod/rod/lib/defaults

##### Example (Pool)

``` go
```
#### func New 

``` go
func New() *Browser
```

New creates a controller. DefaultDevice to emulate is set to [devices.LaptopWithMDPIScreen](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#LaptopWithMDPIScreen).Landscape(), it will change the default user-agent and can make the actual view area smaller than the browser window on headful mode, you can use [Browser.NoDefaultDevice](https://pkg.go.dev/github.com/go-rod/rod#Browser.NoDefaultDevice) to disable it.

#### (*Browser) Call 

``` go
func (b *Browser) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res []byte, err error)
```

Call implements the [proto.Client](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#Client) to call raw cdp interface directly.

#### (*Browser) CancelTimeout 

``` go
func (b *Browser) CancelTimeout() *Browser
```

CancelTimeout cancels the current timeout context and returns a clone with the parent context.

#### (*Browser) Client 

``` go
func (b *Browser) Client(c CDPClient) *Browser
```

Client set the cdp client.

#### (*Browser) Close 

``` go
func (b *Browser) Close() error
```

Close the browser.

#### (*Browser) Connect 

``` go
func (b *Browser) Connect() error
```

Connect to the browser and start to control it. If fails to connect, try to launch a local browser, if local browser not found try to download one.

#### (*Browser) Context 

``` go
func (b *Browser) Context(ctx context.Context) *Browser
```

Context returns a clone with the specified ctx for chained sub-operations.

#### (*Browser) ControlURL 

``` go
func (b *Browser) ControlURL(url string) *Browser
```

ControlURL set the url to remote control browser.

#### (*Browser) DefaultDevice <- 0.71.0

``` go
func (b *Browser) DefaultDevice(d devices.Device) *Browser
```

DefaultDevice sets the default device for new page to emulate in the future. Default is [devices.LaptopWithMDPIScreen](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#LaptopWithMDPIScreen). Set it to [devices.Clear](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/devices#Clear) to disable it.

#### (*Browser) DisableDomain 

``` go
func (b *Browser) DisableDomain(sessionID proto.TargetSessionID, req proto.Request) (restore func())
```

DisableDomain and returns a restore function to restore previous state.

#### (*Browser) EachEvent 

``` go
func (b *Browser) EachEvent(callbacks ...interface{}) (wait func())
```

EachEvent is similar to [Page.EachEvent](https://pkg.go.dev/github.com/go-rod/rod#Page.EachEvent), but catches events of the entire browser.

#### (*Browser) EnableDomain 

``` go
func (b *Browser) EnableDomain(sessionID proto.TargetSessionID, req proto.Request) (restore func())
```

EnableDomain and returns a restore function to restore previous state.

#### (*Browser) Event 

``` go
func (b *Browser) Event() <-chan *Message
```

Event of the browser.

#### (*Browser) GetContext 

``` go
func (b *Browser) GetContext() context.Context
```

GetContext of current instance.

#### (*Browser) GetCookies <- 0.71.0

``` go
func (b *Browser) GetCookies() ([]*proto.NetworkCookie, error)
```

GetCookies from the browser.

#### (*Browser) HandleAuth 

``` go
func (b *Browser) HandleAuth(username, password string) func() error
```

HandleAuth for the next basic HTTP authentication. It will prevent the popup that requires user to input user name and password. Ref: https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication

#### (*Browser) HijackRequests 

``` go
func (b *Browser) HijackRequests() *HijackRouter
```

HijackRequests same as Page.HijackRequests, but can intercept requests of the entire browser.

#### (*Browser) IgnoreCertErrors <- 0.61.3

``` go
func (b *Browser) IgnoreCertErrors(enable bool) error
```

IgnoreCertErrors switch. If enabled, all certificate errors will be ignored.

#### (*Browser) Incognito 

``` go
func (b *Browser) Incognito() (*Browser, error)
```

Incognito creates a new incognito browser.

#### (*Browser) LoadState 

``` go
func (b *Browser) LoadState(sessionID proto.TargetSessionID, method proto.Request) (has bool)
```

LoadState into the method, sessionID can be empty.

#### (*Browser) Logger <- 0.70.0

``` go
func (b *Browser) Logger(l utils.Logger) *Browser
```

Logger overrides the default log functions for tracing.

#### (*Browser) Monitor <- 0.70.0

``` go
func (b *Browser) Monitor(url string) *Browser
```

Monitor address to listen if not empty. Shortcut for [Browser.ServeMonitor](https://pkg.go.dev/github.com/go-rod/rod#Browser.ServeMonitor).

#### (*Browser) MustClose <- 0.50.0

``` go
func (b *Browser) MustClose()
```

MustClose is similar to [Browser.Close](https://pkg.go.dev/github.com/go-rod/rod#Browser.Close).

#### (*Browser) MustConnect <- 0.50.0

``` go
func (b *Browser) MustConnect() *Browser
```

MustConnect is similar to [Browser.Connect](https://pkg.go.dev/github.com/go-rod/rod#Browser.Connect).

#### (*Browser) MustGetCookies <- 0.71.0

``` go
func (b *Browser) MustGetCookies() []*proto.NetworkCookie
```

MustGetCookies is similar to [Browser.GetCookies](https://pkg.go.dev/github.com/go-rod/rod#Browser.GetCookies).

#### (*Browser) MustHandleAuth <- 0.50.0

``` go
func (b *Browser) MustHandleAuth(username, password string) (wait func())
```

MustHandleAuth is similar to [Browser.HandleAuth](https://pkg.go.dev/github.com/go-rod/rod#Browser.HandleAuth).

#### (*Browser) MustIgnoreCertErrors <- 0.61.3

``` go
func (b *Browser) MustIgnoreCertErrors(enable bool) *Browser
```

MustIgnoreCertErrors is similar to [Browser.IgnoreCertErrors](https://pkg.go.dev/github.com/go-rod/rod#Browser.IgnoreCertErrors).

#### (*Browser) MustIncognito <- 0.50.0

``` go
func (b *Browser) MustIncognito() *Browser
```

MustIncognito is similar to [Browser.Incognito](https://pkg.go.dev/github.com/go-rod/rod#Browser.Incognito).

#### (*Browser) MustPage <- 0.50.0

``` go
func (b *Browser) MustPage(url ...string) *Page
```

MustPage is similar to [Browser.Page](https://pkg.go.dev/github.com/go-rod/rod#Browser.Page). The url list will be joined by "/".

#### (*Browser) MustPageFromTargetID <- 0.50.0

``` go
func (b *Browser) MustPageFromTargetID(targetID proto.TargetTargetID) *Page
```

MustPageFromTargetID is similar to [Browser.PageFromTargetID].

#### (*Browser) MustPages <- 0.50.0

``` go
func (b *Browser) MustPages() Pages
```

MustPages is similar to [Browser.Pages](https://pkg.go.dev/github.com/go-rod/rod#Browser.Pages).

#### (*Browser) MustSetCookies <- 0.71.0

``` go
func (b *Browser) MustSetCookies(cookies ...*proto.NetworkCookie) *Browser
```

MustSetCookies is similar to [Browser.SetCookies](https://pkg.go.dev/github.com/go-rod/rod#Browser.SetCookies). If the len(cookies) is 0 it will clear all the cookies.

#### (*Browser) MustVersion <- 0.107.0

``` go
func (b *Browser) MustVersion() *proto.BrowserGetVersionResult
```

MustVersion is similar to [Browser.Version](https://pkg.go.dev/github.com/go-rod/rod#Browser.Version).

#### (*Browser) MustWaitDownload <- 0.83.0

``` go
func (b *Browser) MustWaitDownload() func() []byte
```

MustWaitDownload is similar to [Browser.WaitDownload](https://pkg.go.dev/github.com/go-rod/rod#Browser.WaitDownload). It will read the file into bytes then remove the file.

#### (*Browser) NoDefaultDevice <- 0.81.1

``` go
func (b *Browser) NoDefaultDevice() *Browser
```

NoDefaultDevice is the same as [Browser.DefaultDevice](https://pkg.go.dev/github.com/go-rod/rod#Browser.DefaultDevice)(devices.Clear).

#### (*Browser) Page 

``` go
func (b *Browser) Page(opts proto.TargetCreateTarget) (p *Page, err error)
```

Page creates a new browser tab. If opts.URL is empty, the default target will be "about:blank".

#### (*Browser) PageFromSession <- 0.74.0

``` go
func (b *Browser) PageFromSession(sessionID proto.TargetSessionID) *Page
```

PageFromSession is used for low-level debugging.

#### (*Browser) PageFromTarget <- 0.50.0

``` go
func (b *Browser) PageFromTarget(targetID proto.TargetTargetID) (*Page, error)
```

PageFromTarget gets or creates a Page instance.

#### (*Browser) Pages 

``` go
func (b *Browser) Pages() (Pages, error)
```

Pages retrieves all visible pages.

#### (*Browser) RemoveState <- 0.74.0

``` go
func (b *Browser) RemoveState(key interface{})
```

RemoveState a state.

#### (*Browser) ServeMonitor 

``` go
func (b *Browser) ServeMonitor(host string) string
```

ServeMonitor starts the monitor server. The reason why not to use "chrome://inspect/#devices" is one target cannot be driven by multiple controllers.

#### (*Browser) SetCookies <- 0.71.0

``` go
func (b *Browser) SetCookies(cookies []*proto.NetworkCookieParam) error
```

SetCookies to the browser. If the cookies is nil it will clear all the cookies.

#### (*Browser) Sleeper <- 0.50.0

``` go
func (b *Browser) Sleeper(sleeper func() utils.Sleeper) *Browser
```

Sleeper returns a clone with the specified sleeper for chained sub-operations.

#### (*Browser) SlowMotion <- 0.77.0

``` go
func (b *Browser) SlowMotion(delay time.Duration) *Browser
```

SlowMotion set the delay for each control action, such as the simulation of the human inputs.

#### (*Browser) Timeout 

``` go
func (b *Browser) Timeout(d time.Duration) *Browser
```

Timeout returns a clone with the specified total timeout of all chained sub-operations.

#### (*Browser) Trace 

``` go
func (b *Browser) Trace(enable bool) *Browser
```

Trace enables/disables the visual tracing of the input actions on the page.

#### (*Browser) Version <- 0.107.0

``` go
func (b *Browser) Version() (*proto.BrowserGetVersionResult, error)
```

Version info of the browser.

#### (*Browser) WaitDownload <- 0.83.0

``` go
func (b *Browser) WaitDownload(dir string) func() (info *proto.PageDownloadWillBegin)
```

WaitDownload returns a helper to get the next download file. The file path will be:

```
filepath.Join(dir, info.GUID)
```

#### (*Browser) WaitEvent 

``` go
func (b *Browser) WaitEvent(e proto.Event) (wait func())
```

WaitEvent waits for the next event for one time. It will also load the data into the event object.

#### (*Browser) WithCancel <- 0.69.0

``` go
func (b *Browser) WithCancel() (*Browser, func())
```

WithCancel returns a clone with a context cancel function.

#### (*Browser) WithPanic <- 0.100.0

``` go
func (b *Browser) WithPanic(fail func(interface{})) *Browser
```

WithPanic returns a browser clone with the specified panic function. The fail must stop the current goroutine's execution immediately, such as use [runtime.Goexit](https://pkg.go.dev/runtime#Goexit) or panic inside it.

#### type CDPClient <- 0.70.0

``` go
type CDPClient interface {
	Event() <-chan *cdp.Event
	Call(ctx context.Context, sessionID, method string, params interface{}) ([]byte, error)
}
```

CDPClient is usually used to make rod side-effect free. Such as proxy all IO of rod.

#### type CoveredError <- 0.114.8

``` go
type CoveredError struct {
	*Element
}
```

CoveredError error.

#### (*CoveredError) Error <- 0.114.8

``` go
func (e *CoveredError) Error() string
```

Error ...

#### (*CoveredError) Is <- 0.114.8

``` go
func (e *CoveredError) Is(err error) bool
```

Is interface.

#### (*CoveredError) Unwrap <- 0.114.8

``` go
func (e *CoveredError) Unwrap() error
```

Unwrap ...

#### type Element 

``` go
type Element struct {
	Object *proto.RuntimeRemoteObject
	// contains filtered or unexported fields
}
```

Element represents the DOM element.

#### (*Element) Attribute 

``` go
func (el *Element) Attribute(name string) (*string, error)
```

Attribute of the DOM object. Attribute vs Property: https://stackoverflow.com/questions/6003819/what-is-the-difference-between-properties-and-attributes-in-html

#### (*Element) BackgroundImage <- 0.76.6

``` go
func (el *Element) BackgroundImage() ([]byte, error)
```

BackgroundImage returns the css background-image of the element.

#### (*Element) Blur 

``` go
func (el *Element) Blur() error
```

Blur removes focus from the element.

#### (*Element) Call <- 0.70.0

``` go
func (el *Element) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res []byte, err error)
```

Call implements the [proto.Client](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#Client).

#### (*Element) CancelTimeout 

``` go
func (el *Element) CancelTimeout() *Element
```

CancelTimeout cancels the current timeout context and returns a clone with the parent context.

#### (*Element) CanvasToImage <- 0.45.1

``` go
func (el *Element) CanvasToImage(format string, quality float64) ([]byte, error)
```

CanvasToImage get image data of a canvas. The default format is image/png. The default quality is 0.92. doc: https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/toDataURL

#### (*Element) Click 

``` go
func (el *Element) Click(button proto.InputMouseButton, clickCount int) error
```

Click will press then release the button just like a human. Before the action, it will try to scroll to the element, hover the mouse over it, wait until the it's interactable and enabled.

#### (*Element) ContainsElement <- 0.48.0

``` go
func (el *Element) ContainsElement(target *Element) (bool, error)
```

ContainsElement check if the target is equal or inside the element.

#### (*Element) Context 

``` go
func (el *Element) Context(ctx context.Context) *Element
```

Context returns a clone with the specified ctx for chained sub-operations.

#### (*Element) Describe 

``` go
func (el *Element) Describe(depth int, pierce bool) (*proto.DOMNode, error)
```

Describe the current element. The depth is the maximum depth at which children should be retrieved, defaults to 1, use -1 for the entire subtree or provide an integer larger than 0. The pierce decides whether or not iframes and shadow roots should be traversed when returning the subtree. The returned [proto.DOMNode.NodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMNode.NodeID) will always be empty, because NodeID is not stable (when [proto.DOMDocumentUpdated](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMDocumentUpdated) is fired all NodeID on the page will be reassigned to another value) we don't recommend using the NodeID, instead, use the [proto.DOMBackendNodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMBackendNodeID) to identify the element.

#### (*Element) Disabled <- 0.112.5

``` go
func (el *Element) Disabled() (bool, error)
```

Disabled checks if the element is disabled.

#### (*Element) Element 

``` go
func (el *Element) Element(selector string) (*Element, error)
```

Element returns the first child that matches the css selector.

#### (*Element) ElementByJS 

``` go
func (el *Element) ElementByJS(opts *EvalOptions) (*Element, error)
```

ElementByJS returns the element from the return value of the js.

#### (*Element) ElementR <- 0.57.0

``` go
func (el *Element) ElementR(selector, jsRegex string) (*Element, error)
```

ElementR returns the first child element that matches the css selector and its text matches the jsRegex.

#### (*Element) ElementX 

``` go
func (el *Element) ElementX(xPath string) (*Element, error)
```

ElementX returns the first child that matches the XPath selector.

#### (*Element) Elements 

``` go
func (el *Element) Elements(selector string) (Elements, error)
```

Elements returns all elements that match the css selector.

#### (*Element) ElementsByJS 

``` go
func (el *Element) ElementsByJS(opts *EvalOptions) (Elements, error)
```

ElementsByJS returns the elements from the return value of the js.

#### (*Element) ElementsX 

``` go
func (el *Element) ElementsX(xpath string) (Elements, error)
```

ElementsX returns all elements that match the XPath selector.

#### (*Element) Equal <- 0.85.7

``` go
func (el *Element) Equal(elm *Element) (bool, error)
```

Equal checks if the two elements are equal.

#### (*Element) Eval 

``` go
func (el *Element) Eval(js string, params ...interface{}) (*proto.RuntimeRemoteObject, error)
```

Eval is a shortcut for [Element.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Element.Evaluate) with AwaitPromise, ByValue and AutoExp set to true.

#### (*Element) Evaluate <- 0.67.0

``` go
func (el *Element) Evaluate(opts *EvalOptions) (*proto.RuntimeRemoteObject, error)
```

Evaluate is just a shortcut of [Page.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Page.Evaluate) with This set to current element.

#### (*Element) Focus 

``` go
func (el *Element) Focus() error
```

Focus sets focus on the specified element. Before the action, it will try to scroll to the element.

#### (*Element) Frame 

``` go
func (el *Element) Frame() (*Page, error)
```

Frame creates a page instance that represents the iframe.

#### (*Element) GetContext 

``` go
func (el *Element) GetContext() context.Context
```

GetContext of current instance.

#### (*Element) GetSessionID <- 0.72.0

``` go
func (el *Element) GetSessionID() proto.TargetSessionID
```

GetSessionID interface.

#### (*Element) GetXPath <- 0.109.3

``` go
func (el *Element) GetXPath(optimized bool) (string, error)
```

GetXPath returns the xpath of the element.

#### (*Element) HTML 

``` go
func (el *Element) HTML() (string, error)
```

HTML of the element.

#### (*Element) Has 

``` go
func (el *Element) Has(selector string) (bool, *Element, error)
```

Has an element that matches the css selector.

#### (*Element) HasR <- 0.61.0

``` go
func (el *Element) HasR(selector, jsRegex string) (bool, *Element, error)
```

HasR returns true if a child element that matches the css selector and its text matches the jsRegex.

#### (*Element) HasX 

``` go
func (el *Element) HasX(selector string) (bool, *Element, error)
```

HasX an element that matches the XPath selector.

#### (*Element) Hover <- 0.49.1

``` go
func (el *Element) Hover() error
```

Hover the mouse over the center of the element. Before the action, it will try to scroll to the element and wait until it's interactable.

#### (*Element) Input 

``` go
func (el *Element) Input(text string) error
```

Input focuses on the element and input text to it. Before the action, it will scroll to the element, wait until it's visible, enabled and writable. To empty the input you can use something like

```
el.SelectAllText().MustInput("")
```

#### (*Element) InputColor <- 0.114.3

``` go
func (el *Element) InputColor(color string) error
```

InputColor focuses on the element and inputs a color string to it. Before the action, it will scroll to the element, wait until it's visible, enabled and writable.

#### (*Element) InputTime <- 0.79.2

``` go
func (el *Element) InputTime(t time.Time) error
```

InputTime focuses on the element and input time to it. Before the action, it will scroll to the element, wait until it's visible, enabled and writable. It will wait until the element is visible, enabled and writable.

#### (*Element) Interactable <- 0.66.0

``` go
func (el *Element) Interactable() (pt *proto.Point, err error)
```

Interactable checks if the element is interactable with cursor. The cursor can be mouse, finger, stylus, etc. If not interactable err will be ErrNotInteractable, such as when covered by a modal,.

#### (*Element) KeyActions <- 0.107.0

``` go
func (el *Element) KeyActions() (*KeyActions, error)
```

KeyActions is similar with Page.KeyActions. Before the action, it will try to scroll to the element and focus on it.

#### (*Element) Matches <- 0.45.0

``` go
func (el *Element) Matches(selector string) (bool, error)
```

Matches checks if the element can be selected by the css selector.

#### (*Element) MoveMouseOut <- 0.97.13

``` go
func (el *Element) MoveMouseOut() error
```

MoveMouseOut of the current element.

#### (*Element) MustAttribute <- 0.50.0

``` go
func (el *Element) MustAttribute(name string) *string
```

MustAttribute is similar to [Element.Attribute](https://pkg.go.dev/github.com/go-rod/rod#Element.Attribute).

#### (*Element) MustBackgroundImage <- 0.76.6

``` go
func (el *Element) MustBackgroundImage() []byte
```

MustBackgroundImage is similar to [Element.BackgroundImage](https://pkg.go.dev/github.com/go-rod/rod#Element.BackgroundImage).

#### (*Element) MustBlur <- 0.50.0

``` go
func (el *Element) MustBlur() *Element
```

MustBlur is similar to [Element.Blur](https://pkg.go.dev/github.com/go-rod/rod#Element.Blur).

#### (*Element) MustCanvasToImage <- 0.50.0

``` go
func (el *Element) MustCanvasToImage() []byte
```

MustCanvasToImage is similar to [Element.CanvasToImage](https://pkg.go.dev/github.com/go-rod/rod#Element.CanvasToImage).

#### (*Element) MustClick <- 0.50.0

``` go
func (el *Element) MustClick() *Element
```

MustClick is similar to [Element.Click](https://pkg.go.dev/github.com/go-rod/rod#Element.Click).

#### (*Element) MustContainsElement <- 0.50.0

``` go
func (el *Element) MustContainsElement(target *Element) bool
```

MustContainsElement is similar to [Element.ContainsElement](https://pkg.go.dev/github.com/go-rod/rod#Element.ContainsElement).

#### (*Element) MustDescribe <- 0.50.0

``` go
func (el *Element) MustDescribe() *proto.DOMNode
```

MustDescribe is similar to [Element.Describe](https://pkg.go.dev/github.com/go-rod/rod#Element.Describe).

#### (*Element) MustDisabled <- 0.112.5

``` go
func (el *Element) MustDisabled() bool
```

MustDisabled is similar to [Element.Disabled](https://pkg.go.dev/github.com/go-rod/rod#Element.Disabled).

#### (*Element) MustDoubleClick <- 0.111.0

``` go
func (el *Element) MustDoubleClick() *Element
```

MustDoubleClick is similar to [Element.Click](https://pkg.go.dev/github.com/go-rod/rod#Element.Click).

#### (*Element) MustElement <- 0.50.0

``` go
func (el *Element) MustElement(selector string) *Element
```

MustElement is similar to [Element.Element](https://pkg.go.dev/github.com/go-rod/rod#Element.Element).

#### (*Element) MustElementByJS <- 0.50.0

``` go
func (el *Element) MustElementByJS(js string, params ...interface{}) *Element
```

MustElementByJS is similar to [Element.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementByJS).

#### (*Element) MustElementR <- 0.57.0

``` go
func (el *Element) MustElementR(selector, regex string) *Element
```

MustElementR is similar to [Element.ElementR](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementR).

#### (*Element) MustElementX <- 0.50.0

``` go
func (el *Element) MustElementX(xpath string) *Element
```

MustElementX is similar to [Element.ElementX](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementX).

#### (*Element) MustElements <- 0.50.0

``` go
func (el *Element) MustElements(selector string) Elements
```

MustElements is similar to [Element.Elements](https://pkg.go.dev/github.com/go-rod/rod#Element.Elements).

#### (*Element) MustElementsByJS <- 0.50.0

``` go
func (el *Element) MustElementsByJS(js string, params ...interface{}) Elements
```

MustElementsByJS is similar to [Element.ElementsByJS](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementsByJS).

#### (*Element) MustElementsX <- 0.50.0

``` go
func (el *Element) MustElementsX(xpath string) Elements
```

MustElementsX is similar to [Element.ElementsX](https://pkg.go.dev/github.com/go-rod/rod#Element.ElementsX).

#### (*Element) MustEqual <- 0.85.7

``` go
func (el *Element) MustEqual(elm *Element) bool
```

MustEqual is similar to [Element.Equal](https://pkg.go.dev/github.com/go-rod/rod#Element.Equal).

#### (*Element) MustEval <- 0.50.0

``` go
func (el *Element) MustEval(js string, params ...interface{}) gson.JSON
```

MustEval is similar to [Element.Eval](https://pkg.go.dev/github.com/go-rod/rod#Element.Eval).

#### (*Element) MustFocus <- 0.50.0

``` go
func (el *Element) MustFocus() *Element
```

MustFocus is similar to [Element.Focus](https://pkg.go.dev/github.com/go-rod/rod#Element.Focus).

#### (*Element) MustFrame <- 0.55.1

``` go
func (el *Element) MustFrame() *Page
```

MustFrame is similar to [Element.Frame](https://pkg.go.dev/github.com/go-rod/rod#Element.Frame).

#### (*Element) MustGetXPath <- 0.109.3

``` go
func (el *Element) MustGetXPath(optimized bool) string
```

MustGetXPath is similar to [Element.GetXPath](https://pkg.go.dev/github.com/go-rod/rod#Element.GetXPath).

#### (*Element) MustHTML <- 0.50.0

``` go
func (el *Element) MustHTML() string
```

MustHTML is similar to [Element.HTML](https://pkg.go.dev/github.com/go-rod/rod#Element.HTML).

#### (*Element) MustHas <- 0.50.0

``` go
func (el *Element) MustHas(selector string) bool
```

MustHas is similar to [Element.Has](https://pkg.go.dev/github.com/go-rod/rod#Element.Has).

#### (*Element) MustHasR <- 0.61.0

``` go
func (el *Element) MustHasR(selector, regex string) bool
```

MustHasR is similar to [Element.HasR](https://pkg.go.dev/github.com/go-rod/rod#Element.HasR).

#### (*Element) MustHasX <- 0.50.0

``` go
func (el *Element) MustHasX(selector string) bool
```

MustHasX is similar to [Element.HasX](https://pkg.go.dev/github.com/go-rod/rod#Element.HasX).

#### (*Element) MustHover <- 0.50.0

``` go
func (el *Element) MustHover() *Element
```

MustHover is similar to [Element.Hover](https://pkg.go.dev/github.com/go-rod/rod#Element.Hover).

#### (*Element) MustInput <- 0.50.0

``` go
func (el *Element) MustInput(text string) *Element
```

MustInput is similar to [Element.Input](https://pkg.go.dev/github.com/go-rod/rod#Element.Input).

#### (*Element) MustInputColor <- 0.114.3

``` go
func (el *Element) MustInputColor(color string) *Element
```

MustInputColor is similar to [Element.InputColor](https://pkg.go.dev/github.com/go-rod/rod#Element.InputColor).

#### (*Element) MustInputTime <- 0.79.2

``` go
func (el *Element) MustInputTime(t time.Time) *Element
```

MustInputTime is similar to [Element.Input](https://pkg.go.dev/github.com/go-rod/rod#Element.Input).

#### (*Element) MustInteractable <- 0.66.0

``` go
func (el *Element) MustInteractable() bool
```

MustInteractable is similar to [Element.Interactable](https://pkg.go.dev/github.com/go-rod/rod#Element.Interactable).

#### (*Element) MustKeyActions <- 0.107.0

``` go
func (el *Element) MustKeyActions() *KeyActions
```

MustKeyActions is similar to [Element.KeyActions](https://pkg.go.dev/github.com/go-rod/rod#Element.KeyActions).

#### (*Element) MustMatches <- 0.50.0

``` go
func (el *Element) MustMatches(selector string) bool
```

MustMatches is similar to [Element.Matches](https://pkg.go.dev/github.com/go-rod/rod#Element.Matches).

#### (*Element) MustMoveMouseOut <- 0.97.13

``` go
func (el *Element) MustMoveMouseOut() *Element
```

MustMoveMouseOut is similar to [Element.MoveMouseOut](https://pkg.go.dev/github.com/go-rod/rod#Element.MoveMouseOut).

#### (*Element) MustNext <- 0.50.0

``` go
func (el *Element) MustNext() *Element
```

MustNext is similar to [Element.Next](https://pkg.go.dev/github.com/go-rod/rod#Element.Next).

#### (*Element) MustParent <- 0.50.0

``` go
func (el *Element) MustParent() *Element
```

MustParent is similar to [Element.Parent](https://pkg.go.dev/github.com/go-rod/rod#Element.Parent).

#### (*Element) MustParents <- 0.50.0

``` go
func (el *Element) MustParents(selector string) Elements
```

MustParents is similar to [Element.Parents](https://pkg.go.dev/github.com/go-rod/rod#Element.Parents).

#### (*Element) MustPrevious <- 0.50.0

``` go
func (el *Element) MustPrevious() *Element
```

MustPrevious is similar to [Element.Previous](https://pkg.go.dev/github.com/go-rod/rod#Element.Previous).

#### (*Element) MustProperty <- 0.50.0

``` go
func (el *Element) MustProperty(name string) gson.JSON
```

MustProperty is similar to [Element.Property](https://pkg.go.dev/github.com/go-rod/rod#Element.Property).

#### (*Element) MustRelease <- 0.50.0

``` go
func (el *Element) MustRelease()
```

MustRelease is similar to [Element.Release](https://pkg.go.dev/github.com/go-rod/rod#Element.Release).

#### (*Element) MustRemove <- 0.66.0

``` go
func (el *Element) MustRemove()
```

MustRemove is similar to [Element.Remove](https://pkg.go.dev/github.com/go-rod/rod#Element.Remove).

#### (*Element) MustResource <- 0.50.0

``` go
func (el *Element) MustResource() []byte
```

MustResource is similar to [Element.Resource](https://pkg.go.dev/github.com/go-rod/rod#Element.Resource).

#### (*Element) MustScreenshot <- 0.50.0

``` go
func (el *Element) MustScreenshot(toFile ...string) []byte
```

MustScreenshot is similar to [Element.Screenshot](https://pkg.go.dev/github.com/go-rod/rod#Element.Screenshot).

#### (*Element) MustScrollIntoView <- 0.50.0

``` go
func (el *Element) MustScrollIntoView() *Element
```

MustScrollIntoView is similar to [Element.ScrollIntoView](https://pkg.go.dev/github.com/go-rod/rod#Element.ScrollIntoView).

#### (*Element) MustSelect <- 0.50.0

``` go
func (el *Element) MustSelect(selectors ...string) *Element
```

MustSelect is similar to [Element.Select](https://pkg.go.dev/github.com/go-rod/rod#Element.Select).

#### (*Element) MustSelectAllText <- 0.50.0

``` go
func (el *Element) MustSelectAllText() *Element
```

MustSelectAllText is similar to [Element.SelectAllText](https://pkg.go.dev/github.com/go-rod/rod#Element.SelectAllText).

#### (*Element) MustSelectText <- 0.50.0

``` go
func (el *Element) MustSelectText(regex string) *Element
```

MustSelectText is similar to [Element.SelectText](https://pkg.go.dev/github.com/go-rod/rod#Element.SelectText).

#### (*Element) MustSetFiles <- 0.50.0

``` go
func (el *Element) MustSetFiles(paths ...string) *Element
```

MustSetFiles is similar to [Element.SetFiles](https://pkg.go.dev/github.com/go-rod/rod#Element.SetFiles).

#### (*Element) MustShadowRoot <- 0.50.0

``` go
func (el *Element) MustShadowRoot() *Element
```

MustShadowRoot is similar to [Element.ShadowRoot](https://pkg.go.dev/github.com/go-rod/rod#Element.ShadowRoot).

#### (*Element) MustShape <- 0.66.0

``` go
func (el *Element) MustShape() *proto.DOMGetContentQuadsResult
```

MustShape is similar to [Element.Shape](https://pkg.go.dev/github.com/go-rod/rod#Element.Shape).

#### (*Element) MustTap <- 0.61.4

``` go
func (el *Element) MustTap() *Element
```

MustTap is similar to [Element.Tap](https://pkg.go.dev/github.com/go-rod/rod#Element.Tap).

#### (*Element) MustText <- 0.50.0

``` go
func (el *Element) MustText() string
```

MustText is similar to [Element.Text](https://pkg.go.dev/github.com/go-rod/rod#Element.Text).

#### (*Element) MustType <- 0.107.0

``` go
func (el *Element) MustType(keys ...input.Key) *Element
```

MustType is similar to [Element.Type](https://pkg.go.dev/github.com/go-rod/rod#Element.Type).

#### (*Element) MustVisible <- 0.50.0

``` go
func (el *Element) MustVisible() bool
```

MustVisible is similar to [Element.Visible](https://pkg.go.dev/github.com/go-rod/rod#Element.Visible).

#### (*Element) MustWait <- 0.50.0

``` go
func (el *Element) MustWait(js string, params ...interface{}) *Element
```

MustWait is similar to [Element.Wait](https://pkg.go.dev/github.com/go-rod/rod#Element.Wait).

#### (*Element) MustWaitEnabled <- 0.84.1

``` go
func (el *Element) MustWaitEnabled() *Element
```

MustWaitEnabled is similar to [Element.WaitEnabled](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitEnabled).

#### (*Element) MustWaitInteractable <- 0.88.0

``` go
func (el *Element) MustWaitInteractable() *Element
```

MustWaitInteractable is similar to [Element.WaitInteractable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitInteractable).

#### (*Element) MustWaitInvisible <- 0.50.0

``` go
func (el *Element) MustWaitInvisible() *Element
```

MustWaitInvisible is similar to [Element.WaitInvisible](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitInvisible)..

#### (*Element) MustWaitLoad <- 0.50.0

``` go
func (el *Element) MustWaitLoad() *Element
```

MustWaitLoad is similar to [Element.WaitLoad](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitLoad).

#### (*Element) MustWaitStable <- 0.50.0

``` go
func (el *Element) MustWaitStable() *Element
```

MustWaitStable is similar to [Element.WaitStable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitStable).

#### (*Element) MustWaitVisible <- 0.50.0

``` go
func (el *Element) MustWaitVisible() *Element
```

MustWaitVisible is similar to [Element.WaitVisible](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitVisible).

#### (*Element) MustWaitWritable <- 0.84.1

``` go
func (el *Element) MustWaitWritable() *Element
```

MustWaitWritable is similar to [Element.WaitWritable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitWritable).

#### (*Element) Next 

``` go
func (el *Element) Next() (*Element, error)
```

Next returns the next sibling element in the DOM tree.

#### (*Element) Overlay <- 0.88.0

``` go
func (el *Element) Overlay(msg string) (removeOverlay func())
```

Overlay msg on the element.

#### (*Element) Page <- 0.101.7

``` go
func (el *Element) Page() *Page
```

Page of the element.

#### (*Element) Parent 

``` go
func (el *Element) Parent() (*Element, error)
```

Parent returns the parent element in the DOM tree.

#### (*Element) Parents 

``` go
func (el *Element) Parents(selector string) (Elements, error)
```

Parents that match the selector.

#### (*Element) Previous 

``` go
func (el *Element) Previous() (*Element, error)
```

Previous returns the previous sibling element in the DOM tree.

#### (*Element) Property 

``` go
func (el *Element) Property(name string) (gson.JSON, error)
```

Property of the DOM object. Property vs Attribute: https://stackoverflow.com/questions/6003819/what-is-the-difference-between-properties-and-attributes-in-html

#### (*Element) Release 

``` go
func (el *Element) Release() error
```

Release is a shortcut for [Page.Release](https://pkg.go.dev/github.com/go-rod/rod#Page.Release) current element.

#### (*Element) Remove <- 0.66.0

``` go
func (el *Element) Remove() error
```

Remove the element from the page.

#### (*Element) Resource 

``` go
func (el *Element) Resource() ([]byte, error)
```

Resource returns the "src" content of current element. Such as the jpg of <img src="a.jpg">.

#### (*Element) Screenshot 

``` go
func (el *Element) Screenshot(format proto.PageCaptureScreenshotFormat, quality int) ([]byte, error)
```

Screenshot of the area of the element.

#### (*Element) ScrollIntoView 

``` go
func (el *Element) ScrollIntoView() error
```

ScrollIntoView scrolls the current element into the visible area of the browser window if it's not already within the visible area.

#### (*Element) Select 

``` go
func (el *Element) Select(selectors []string, selected bool, t SelectorType) error
```

Select the children option elements that match the selectors. Before the action, it will scroll to the element, wait until it's visible. If no option matches the selectors, it will return [ErrElementNotFound].

#### (*Element) SelectAllText 

``` go
func (el *Element) SelectAllText() error
```

SelectAllText selects all text Before the action, it will try to scroll to the element and focus on it.

#### (*Element) SelectText 

``` go
func (el *Element) SelectText(regex string) error
```

SelectText selects the text that matches the regular expression. Before the action, it will try to scroll to the element and focus on it.

#### (*Element) SetFiles 

``` go
func (el *Element) SetFiles(paths []string) error
```

SetFiles of the current file input element.

#### (*Element) ShadowRoot 

``` go
func (el *Element) ShadowRoot() (*Element, error)
```

ShadowRoot returns the shadow root of this element.

#### (*Element) Shape <- 0.66.0

``` go
func (el *Element) Shape() (*proto.DOMGetContentQuadsResult, error)
```

Shape of the DOM element content. The shape is a group of 4-sides polygons. A 4-sides polygon is not necessary a rectangle. 4-sides polygons can be apart from each other. For example, we use 2 4-sides polygons to describe the shape below:

```
  ____________          ____________
 /        ___/    =    /___________/    +     _________
/________/                                   /________/
```

#### (*Element) Sleeper <- 0.50.0

``` go
func (el *Element) Sleeper(sleeper func() utils.Sleeper) *Element
```

Sleeper returns a clone with the specified sleeper for chained sub-operations.

#### (*Element) String <- 0.88.0

``` go
func (el *Element) String() string
```

String interface.

#### (*Element) Tap <- 0.61.4

``` go
func (el *Element) Tap() error
```

Tap will scroll to the button and tap it just like a human. Before the action, it will try to scroll to the element and wait until it's interactable and enabled.

#### (*Element) Text 

``` go
func (el *Element) Text() (string, error)
```

Text that the element displays.

#### (*Element) Timeout 

``` go
func (el *Element) Timeout(d time.Duration) *Element
```

Timeout returns a clone with the specified total timeout of all chained sub-operations.

#### (*Element) Type <- 0.107.0

``` go
func (el *Element) Type(keys ...input.Key) error
```

Type is similar with Keyboard.Type. Before the action, it will try to scroll to the element and focus on it.

#### (*Element) Visible 

``` go
func (el *Element) Visible() (bool, error)
```

Visible returns true if the element is visible on the page.

#### (*Element) Wait 

``` go
func (el *Element) Wait(opts *EvalOptions) error
```

Wait until the js returns true.

#### (*Element) WaitEnabled <- 0.84.1

``` go
func (el *Element) WaitEnabled() error
```

WaitEnabled until the element is not disabled. Doc for readonly: https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/readonly

#### (*Element) WaitInteractable <- 0.88.0

``` go
func (el *Element) WaitInteractable() (pt *proto.Point, err error)
```

WaitInteractable waits for the element to be interactable. It will try to scroll to the element on each try.

#### (*Element) WaitInvisible 

``` go
func (el *Element) WaitInvisible() error
```

WaitInvisible until the element invisible.

#### (*Element) WaitLoad <- 0.49.0

``` go
func (el *Element) WaitLoad() error
```

WaitLoad for element like <img>.

#### (*Element) WaitStable 

``` go
func (el *Element) WaitStable(d time.Duration) error
```

WaitStable waits until no shape or position change for d duration. Be careful, d is not the max wait timeout, it's the least stable time. If you want to set a timeout you can use the [Element.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Element.Timeout) function.

#### (*Element) WaitStableRAF <- 0.84.1

``` go
func (el *Element) WaitStableRAF() error
```

WaitStableRAF waits until no shape or position change for 2 consecutive animation frames. If you want to wait animation that is triggered by JS not CSS, you'd better use [Element.WaitStable](https://pkg.go.dev/github.com/go-rod/rod#Element.WaitStable). About animation frame: https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame

#### (*Element) WaitVisible 

``` go
func (el *Element) WaitVisible() error
```

WaitVisible until the element is visible.

#### (*Element) WaitWritable <- 0.84.1

``` go
func (el *Element) WaitWritable() error
```

WaitWritable until the element is not readonly. Doc for disabled: https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/disabled

#### (*Element) WithCancel <- 0.69.0

``` go
func (el *Element) WithCancel() (*Element, func())
```

WithCancel returns a clone with a context cancel function.

#### (*Element) WithPanic <- 0.100.0

``` go
func (el *Element) WithPanic(fail func(interface{})) *Element
```

WithPanic returns an element clone with the specified panic function. The fail must stop the current goroutine's execution immediately, such as use [runtime.Goexit](https://pkg.go.dev/runtime#Goexit) or panic inside it.

#### type ElementNotFoundError <- 0.114.8

``` go
type ElementNotFoundError struct{}
```

ElementNotFoundError error.

#### (*ElementNotFoundError) Error <- 0.114.8

``` go
func (e *ElementNotFoundError) Error() string
```

#### type Elements 

``` go
type Elements []*Element
```

Elements provides some helpers to deal with element list.

#### (Elements) Empty 

``` go
func (els Elements) Empty() bool
```

Empty returns true if the list is empty.

#### (Elements) First 

``` go
func (els Elements) First() *Element
```

First returns the first element, if the list is empty returns nil.

#### (Elements) Last 

``` go
func (els Elements) Last() *Element
```

Last returns the last element, if the list is empty returns nil.

#### type EvalError <- 0.114.8

``` go
type EvalError struct {
	*proto.RuntimeExceptionDetails
}
```

EvalError error.

#### (*EvalError) Error <- 0.114.8

``` go
func (e *EvalError) Error() string
```

#### (*EvalError) Is <- 0.114.8

``` go
func (e *EvalError) Is(err error) bool
```

Is interface.

#### type EvalOptions <- 0.50.0

``` go
type EvalOptions struct {
	// If enabled the eval result will be a plain JSON value.
	// If disabled the eval result will be a reference of a remote js object.
	ByValue bool

	AwaitPromise bool

	// ThisObj represents the "this" object in the JS
	ThisObj *proto.RuntimeRemoteObject

	// JS function definition to execute.
	JS string

	// JSArgs represents the arguments that will be passed to JS.
	// If an argument is [*proto.RuntimeRemoteObject] type, the corresponding remote object will be used.
	// Or it will be passed as a plain JSON value.
	// When an arg in the args is a *js.Function, the arg will be cached on the page's js context.
	// When the arg.Name exists in the page's cache, it reuse the cache without sending
	// the definition to the browser again.
	// Useful when you need to eval a huge js expression many times.
	JSArgs []interface{}

	// Whether execution should be treated as initiated by user in the UI.
	UserGesture bool
}
```

EvalOptions for Page.Evaluate.

#### func Eval <- 0.67.0

``` go
func Eval(js string, args ...interface{}) *EvalOptions
```

Eval creates a [EvalOptions](https://pkg.go.dev/github.com/go-rod/rod#EvalOptions) with ByValue set to true.

#### (*EvalOptions) ByObject <- 0.50.0

``` go
func (e *EvalOptions) ByObject() *EvalOptions
```

ByObject disables ByValue.

#### (*EvalOptions) ByPromise <- 0.74.0

``` go
func (e *EvalOptions) ByPromise() *EvalOptions
```

ByPromise enables AwaitPromise.

#### (*EvalOptions) ByUser <- 0.64.0

``` go
func (e *EvalOptions) ByUser() *EvalOptions
```

ByUser enables UserGesture.

#### (*EvalOptions) String <- 0.88.0

``` go
func (e *EvalOptions) String() string
```

String interface.

#### (*EvalOptions) This <- 0.50.0

``` go
func (e *EvalOptions) This(obj *proto.RuntimeRemoteObject) *EvalOptions
```

This set the obj as ThisObj.

#### type ExpectElementError <- 0.114.8

``` go
type ExpectElementError struct {
	*proto.RuntimeRemoteObject
}
```

ExpectElementError error.

#### (*ExpectElementError) Error <- 0.114.8

``` go
func (e *ExpectElementError) Error() string
```

#### (*ExpectElementError) Is <- 0.114.8

``` go
func (e *ExpectElementError) Is(err error) bool
```

Is interface.

#### type ExpectElementsError <- 0.114.8

``` go
type ExpectElementsError struct {
	*proto.RuntimeRemoteObject
}
```

ExpectElementsError error.

#### (*ExpectElementsError) Error <- 0.114.8

``` go
func (e *ExpectElementsError) Error() string
```

#### (*ExpectElementsError) Is <- 0.114.8

``` go
func (e *ExpectElementsError) Is(err error) bool
```

Is interface.

#### type Hijack 

``` go
type Hijack struct {
	Request  *HijackRequest
	Response *HijackResponse
	OnError  func(error)

	// Skip to next handler
	Skip bool

	// CustomState is used to store things for this context
	CustomState interface{}
	// contains filtered or unexported fields
}
```

Hijack context.

#### (*Hijack) ContinueRequest <- 0.42.0

``` go
func (h *Hijack) ContinueRequest(cq *proto.FetchContinueRequest)
```

ContinueRequest without hijacking. The RequestID will be set by the router, you don't have to set it.

#### (*Hijack) LoadResponse 

``` go
func (h *Hijack) LoadResponse(client *http.Client, loadBody bool) error
```

LoadResponse will send request to the real destination and load the response as default response to override.

#### (*Hijack) MustLoadResponse <- 0.50.0

``` go
func (h *Hijack) MustLoadResponse()
```

MustLoadResponse is similar to [Hijack.LoadResponse](https://pkg.go.dev/github.com/go-rod/rod#Hijack.LoadResponse).

#### type HijackRequest 

``` go
type HijackRequest struct {
	// contains filtered or unexported fields
}
```

HijackRequest context.

#### (*HijackRequest) Body 

``` go
func (ctx *HijackRequest) Body() string
```

Body of the request, devtools API doesn't support binary data yet, only string can be captured.

#### (*HijackRequest) Header 

``` go
func (ctx *HijackRequest) Header(key string) string
```

Header via a key.

#### (*HijackRequest) Headers 

``` go
func (ctx *HijackRequest) Headers() proto.NetworkHeaders
```

Headers of request.

#### (*HijackRequest) IsNavigation <- 0.97.1

``` go
func (ctx *HijackRequest) IsNavigation() bool
```

IsNavigation determines whether the request is a navigation request.

#### (*HijackRequest) JSONBody 

``` go
func (ctx *HijackRequest) JSONBody() gson.JSON
```

JSONBody of the request.

#### (*HijackRequest) Method 

``` go
func (ctx *HijackRequest) Method() string
```

Method of the request.

#### (*HijackRequest) Req <- 0.52.0

``` go
func (ctx *HijackRequest) Req() *http.Request
```

Req returns the underlying http.Request instance that will be used to send the request.

#### (*HijackRequest) SetBody 

``` go
func (ctx *HijackRequest) SetBody(obj interface{}) *HijackRequest
```

SetBody of the request, if obj is []byte or string, raw body will be used, else it will be encoded as json.

#### (*HijackRequest) SetContext <- 0.57.1

``` go
func (ctx *HijackRequest) SetContext(c context.Context) *HijackRequest
```

SetContext of the underlying http.Request instance.

#### (*HijackRequest) Type <- 0.49.1

``` go
func (ctx *HijackRequest) Type() proto.NetworkResourceType
```

Type of the resource.

#### (*HijackRequest) URL 

``` go
func (ctx *HijackRequest) URL() *url.URL
```

URL of the request.

#### type HijackResponse 

``` go
type HijackResponse struct {
	RawResponse *http.Response
	// contains filtered or unexported fields
}
```

HijackResponse context.

#### (*HijackResponse) Body 

``` go
func (ctx *HijackResponse) Body() string
```

Body of the payload.

#### (*HijackResponse) Fail <- 0.48.1

``` go
func (ctx *HijackResponse) Fail(reason proto.NetworkErrorReason) *HijackResponse
```

Fail request.

#### (*HijackResponse) Headers 

``` go
func (ctx *HijackResponse) Headers() http.Header
```

Headers returns the clone of response headers. If you want to modify the response headers use HijackResponse.SetHeader .

#### (*HijackResponse) Payload <- 0.52.0

``` go
func (ctx *HijackResponse) Payload() *proto.FetchFulfillRequest
```

Payload to respond the request from the browser.

#### (*HijackResponse) SetBody 

``` go
func (ctx *HijackResponse) SetBody(obj interface{}) *HijackResponse
```

SetBody of the payload, if obj is []byte or string, raw body will be used, else it will be encoded as json.

#### (*HijackResponse) SetHeader 

``` go
func (ctx *HijackResponse) SetHeader(pairs ...string) *HijackResponse
```

SetHeader of the payload via key-value pairs.

#### type HijackRouter 

``` go
type HijackRouter struct {
	// contains filtered or unexported fields
}
```

HijackRouter context.

#### (*HijackRouter) Add 

``` go
func (r *HijackRouter) Add(pattern string, resourceType proto.NetworkResourceType, handler func(*Hijack)) error
```

Add a hijack handler to router, the doc of the pattern is the same as "proto.FetchRequestPattern.URLPattern".

#### (*HijackRouter) MustAdd <- 0.50.0

``` go
func (r *HijackRouter) MustAdd(pattern string, handler func(*Hijack)) *HijackRouter
```

MustAdd is similar to [HijackRouter.Add](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Add).

#### (*HijackRouter) MustRemove <- 0.50.0

``` go
func (r *HijackRouter) MustRemove(pattern string) *HijackRouter
```

MustRemove is similar to [HijackRouter.Remove](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Remove).

#### (*HijackRouter) MustStop <- 0.50.0

``` go
func (r *HijackRouter) MustStop()
```

MustStop is similar to [HijackRouter.Stop](https://pkg.go.dev/github.com/go-rod/rod#HijackRouter.Stop).

#### (*HijackRouter) Remove 

``` go
func (r *HijackRouter) Remove(pattern string) error
```

Remove handler via the pattern.

#### (*HijackRouter) Run 

``` go
func (r *HijackRouter) Run()
```

Run the router, after you call it, you shouldn't add new handler to it.

#### (*HijackRouter) Stop 

``` go
func (r *HijackRouter) Stop() error
```

Stop the router.

#### type InvisibleShapeError <- 0.114.8

``` go
type InvisibleShapeError struct {
	*Element
}
```

InvisibleShapeError error.

#### (*InvisibleShapeError) Error <- 0.114.8

``` go
func (e *InvisibleShapeError) Error() string
```

Error ...

#### (*InvisibleShapeError) Is <- 0.114.8

``` go
func (e *InvisibleShapeError) Is(err error) bool
```

Is interface.

#### (*InvisibleShapeError) Unwrap <- 0.114.8

``` go
func (e *InvisibleShapeError) Unwrap() error
```

Unwrap ...

#### type KeyAction <- 0.107.0

``` go
type KeyAction struct {
	Type KeyActionType
	Key  input.Key
}
```

KeyAction to perform.

#### type KeyActionType <- 0.107.0

``` go
type KeyActionType int
```

KeyActionType enum.

``` go
const (
	KeyActionPress KeyActionType = iota
	KeyActionRelease
	KeyActionTypeKey
)
```

KeyActionTypes.

#### type KeyActions <- 0.107.0

``` go
type KeyActions struct {
	Actions []KeyAction
	// contains filtered or unexported fields
}
```

KeyActions to simulate.

#### (*KeyActions) Do <- 0.107.0

``` go
func (ka *KeyActions) Do() (err error)
```

Do the actions.

#### (*KeyActions) MustDo <- 0.107.0

``` go
func (ka *KeyActions) MustDo()
```

MustDo is similar to [KeyActions.Do](https://pkg.go.dev/github.com/go-rod/rod#KeyActions.Do).

#### (*KeyActions) Press <- 0.107.0

``` go
func (ka *KeyActions) Press(keys ...input.Key) *KeyActions
```

Press keys is guaranteed to have a release at the end of actions.

#### (*KeyActions) Release <- 0.107.0

``` go
func (ka *KeyActions) Release(keys ...input.Key) *KeyActions
```

Release keys.

#### (*KeyActions) Type <- 0.107.0

``` go
func (ka *KeyActions) Type(keys ...input.Key) *KeyActions
```

Type will release the key immediately after the pressing.

#### type Keyboard 

``` go
type Keyboard struct {
	sync.Mutex
	// contains filtered or unexported fields
}
```

Keyboard represents the keyboard on a page, it's always related the main frame.

#### (*Keyboard) MustType <- 0.107.0

``` go
func (k *Keyboard) MustType(key ...input.Key) *Keyboard
```

MustType is similar to [Keyboard.Type](https://pkg.go.dev/github.com/go-rod/rod#Keyboard.Type).

#### (*Keyboard) Press 

``` go
func (k *Keyboard) Press(key input.Key) error
```

Press the key down. To input characters that are not on the keyboard, such as Chinese or Japanese, you should use method like [Page.InsertText](https://pkg.go.dev/github.com/go-rod/rod#Page.InsertText).

#### (*Keyboard) Release <- 0.107.0

``` go
func (k *Keyboard) Release(key input.Key) error
```

Release the key.

#### (*Keyboard) Type <- 0.107.0

``` go
func (k *Keyboard) Type(keys ...input.Key) (err error)
```

Type releases the key after the press.

#### type Message <- 0.74.0

``` go
type Message struct {
	SessionID proto.TargetSessionID
	Method    string
	// contains filtered or unexported fields
}
```

Message represents a cdp.Event.

#### (*Message) Load <- 0.74.0

``` go
func (msg *Message) Load(e proto.Event) bool
```

Load data into e, returns true if e matches the event type.

#### type Mouse 

``` go
type Mouse struct {
	sync.Mutex
	// contains filtered or unexported fields
}
```

Mouse represents the mouse on a page, it's always related the main frame.

#### (*Mouse) Click 

``` go
func (m *Mouse) Click(button proto.InputMouseButton, clickCount int) error
```

Click the button. It's the combination of [Mouse.Down](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Down) and [Mouse.Up](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Up).

#### (*Mouse) Down 

``` go
func (m *Mouse) Down(button proto.InputMouseButton, clickCount int) error
```

Down holds the button down.

#### (*Mouse) MoveAlong <- 0.112.0

``` go
func (m *Mouse) MoveAlong(guide func() (proto.Point, bool)) error
```

MoveAlong the guide function. Every time the guide function is called it should return the next mouse position, return true to stop. Read the source code of [Mouse.MoveLinear](https://pkg.go.dev/github.com/go-rod/rod#Mouse.MoveLinear) as an example to use this method.

#### (*Mouse) MoveLinear <- 0.112.0

``` go
func (m *Mouse) MoveLinear(to proto.Point, steps int) error
```

MoveLinear to the absolute position with the given steps. Such as move from (0,0) to (6,6) with 3 steps, the mouse will first move to (2,2) then (4,4) then (6,6).

#### (*Mouse) MoveTo <- 0.112.0

``` go
func (m *Mouse) MoveTo(p proto.Point) error
```

MoveTo the absolute position.

#### (*Mouse) MustClick <- 0.50.0

``` go
func (m *Mouse) MustClick(button proto.InputMouseButton) *Mouse
```

MustClick is similar to [Mouse.Click](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Click).

#### (*Mouse) MustDown <- 0.50.0

``` go
func (m *Mouse) MustDown(button proto.InputMouseButton) *Mouse
```

MustDown is similar to [Mouse.Down](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Down).

#### (*Mouse) MustMoveTo <- 0.112.0

``` go
func (m *Mouse) MustMoveTo(x, y float64) *Mouse
```

MustMoveTo is similar to [Mouse.Move].

#### (*Mouse) MustScroll <- 0.50.0

``` go
func (m *Mouse) MustScroll(x, y float64) *Mouse
```

MustScroll is similar to [Mouse.Scroll](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Scroll).

#### (*Mouse) MustUp <- 0.50.0

``` go
func (m *Mouse) MustUp(button proto.InputMouseButton) *Mouse
```

MustUp is similar to [Mouse.Up](https://pkg.go.dev/github.com/go-rod/rod#Mouse.Up).

#### (*Mouse) Position <- 0.112.0

``` go
func (m *Mouse) Position() proto.Point
```

Position of current cursor.

#### (*Mouse) Scroll 

``` go
func (m *Mouse) Scroll(offsetX, offsetY float64, steps int) error
```

Scroll the relative offset with specified steps.

#### (*Mouse) Up 

``` go
func (m *Mouse) Up(button proto.InputMouseButton, clickCount int) error
```

Up releases the button.

#### type NavigationError <- 0.114.8

``` go
type NavigationError struct {
	Reason string
}
```

NavigationError error.

#### (*NavigationError) Error <- 0.114.8

``` go
func (e *NavigationError) Error() string
```

#### (*NavigationError) Is <- 0.114.8

``` go
func (e *NavigationError) Is(err error) bool
```

Is interface.

#### type NoPointerEventsError <- 0.114.8

``` go
type NoPointerEventsError struct {
	*Element
}
```

NoPointerEventsError error.

#### (*NoPointerEventsError) Error <- 0.114.8

``` go
func (e *NoPointerEventsError) Error() string
```

Error ...

#### (*NoPointerEventsError) Is <- 0.114.8

``` go
func (e *NoPointerEventsError) Is(err error) bool
```

Is interface.

#### (*NoPointerEventsError) Unwrap <- 0.114.8

``` go
func (e *NoPointerEventsError) Unwrap() error
```

Unwrap ...

#### type NoShadowRootError <- 0.114.8

``` go
type NoShadowRootError struct {
	*Element
}
```

NoShadowRootError error.

#### (*NoShadowRootError) Error <- 0.114.8

``` go
func (e *NoShadowRootError) Error() string
```

Error ...

#### (*NoShadowRootError) Is <- 0.114.8

``` go
func (e *NoShadowRootError) Is(err error) bool
```

Is interface.

#### type NotInteractableError <- 0.114.8

``` go
type NotInteractableError struct{}
```

NotInteractableError error. Check the doc of Element.Interactable for details.

#### (*NotInteractableError) Error <- 0.114.8

``` go
func (e *NotInteractableError) Error() string
```

#### type ObjectNotFoundError <- 0.114.8

``` go
type ObjectNotFoundError struct {
	*proto.RuntimeRemoteObject
}
```

ObjectNotFoundError error.

#### (*ObjectNotFoundError) Error <- 0.114.8

``` go
func (e *ObjectNotFoundError) Error() string
```

#### (*ObjectNotFoundError) Is <- 0.114.8

``` go
func (e *ObjectNotFoundError) Is(err error) bool
```

Is interface.

#### type Page 

``` go
type Page struct {
	// TargetID is a unique ID for a remote page.
	// It's usually used in events sent from the browser to tell which page an event belongs to.
	TargetID proto.TargetTargetID

	// FrameID is a unique ID for a browsing context.
	// Usually, different FrameID means different javascript execution context.
	// Such as an iframe and the page it belongs to will have the same TargetID but different FrameIDs.
	FrameID proto.PageFrameID

	// SessionID is a unique ID for a page attachment to a controller.
	// It's usually used in transport layer to tell which page to send the control signal.
	// A page can attached to multiple controllers, the browser uses it distinguish controllers.
	SessionID proto.TargetSessionID

	// devices
	Mouse    *Mouse
	Keyboard *Keyboard
	Touch    *Touch
	// contains filtered or unexported fields
}
```

Page represents the webpage. We try to hold as less states as possible. When a page is closed by Rod or not all the ongoing operations an events on it will abort.

##### Example (Pool)

``` go
```
#### (*Page) Activate <- 0.86.3

``` go
func (p *Page) Activate() (*Page, error)
```

Activate (focuses) the page.

#### (*Page) AddScriptTag 

``` go
func (p *Page) AddScriptTag(url, content string) error
```

AddScriptTag to page. If url is empty, content will be used.

#### (*Page) AddStyleTag 

``` go
func (p *Page) AddStyleTag(url, content string) error
```

AddStyleTag to page. If url is empty, content will be used.

#### (*Page) Browser <- 0.101.7

``` go
func (p *Page) Browser() *Browser
```

Browser of the page.

#### (*Page) Call <- 0.70.0

``` go
func (p *Page) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res []byte, err error)
```

Call implements the [proto.Client](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#Client).

#### (*Page) CancelTimeout 

``` go
func (p *Page) CancelTimeout() *Page
```

CancelTimeout cancels the current timeout context and returns a clone with the parent context.

#### (*Page) CaptureDOMSnapshot <- 0.113.0

``` go
func (p *Page) CaptureDOMSnapshot() (domSnapshot *proto.DOMSnapshotCaptureSnapshotResult, err error)
```

CaptureDOMSnapshot Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened. `Documents` The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document. `Strings` Shared string table that all string properties refer to with indexes. Normally use `Strings` is enough.

#### (*Page) Close 

``` go
func (p *Page) Close() error
```

Close tries to close page, running its beforeunload hooks, if has any.

#### (*Page) Context 

``` go
func (p *Page) Context(ctx context.Context) *Page
```

Context returns a clone with the specified ctx for chained sub-operations.

#### (*Page) Cookies 

``` go
func (p *Page) Cookies(urls []string) ([]*proto.NetworkCookie, error)
```

Cookies returns the page cookies. By default it will return the cookies for current page. The urls is the list of URLs for which applicable cookies will be fetched.

#### (*Page) DisableDomain 

``` go
func (p *Page) DisableDomain(method proto.Request) (restore func())
```

DisableDomain and returns a restore function to restore previous state.

#### (*Page) EachEvent 

``` go
func (p *Page) EachEvent(callbacks ...interface{}) (wait func())
```

EachEvent of the specified event types, if any callback returns true the wait function will resolve, The type of each callback is (? means optional):

``` go
func(proto.Event, proto.TargetSessionID?) bool?
```

You can listen to multiple event types at the same time like:

```
browser.EachEvent(func(a *proto.A) {}, func(b *proto.B) {})
```

Such as subscribe the events to know when the navigation is complete or when the page is rendered. Here's an example to dismiss all dialogs/alerts on the page:

```
go page.EachEvent(func(e *proto.PageJavascriptDialogOpening) {
    _ = proto.PageHandleJavaScriptDialog{ Accept: false, PromptText: ""}.Call(page)
})()
```

#### (*Page) Element 

``` go
func (p *Page) Element(selector string) (*Element, error)
```

Element retries until an element in the page that matches the CSS selector, then returns the matched element.

#### (*Page) ElementByJS 

``` go
func (p *Page) ElementByJS(opts *EvalOptions) (*Element, error)
```

ElementByJS returns the element from the return value of the js function. If sleeper is nil, no retry will be performed. By default, it will retry until the js function doesn't return null. To customize the retry logic, check the examples of Page.Sleeper.

#### (*Page) ElementFromNode <- 0.47.0

``` go
func (p *Page) ElementFromNode(node *proto.DOMNode) (*Element, error)
```

ElementFromNode creates an Element from the node, [proto.DOMNodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMNodeID) or [proto.DOMBackendNodeID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#DOMBackendNodeID) must be specified.

#### (*Page) ElementFromObject <- 0.47.0

``` go
func (p *Page) ElementFromObject(obj *proto.RuntimeRemoteObject) (*Element, error)
```

ElementFromObject creates an Element from the remote object id.

#### (*Page) ElementFromPoint <- 0.48.0

``` go
func (p *Page) ElementFromPoint(x, y int) (*Element, error)
```

ElementFromPoint creates an Element from the absolute point on the page. The point should include the window scroll offset.

#### (*Page) ElementR <- 0.57.0

``` go
func (p *Page) ElementR(selector, jsRegex string) (*Element, error)
```

ElementR retries until an element in the page that matches the css selector and it's text matches the jsRegex, then returns the matched element.

#### (*Page) ElementX 

``` go
func (p *Page) ElementX(xPath string) (*Element, error)
```

ElementX retries until an element in the page that matches one of the XPath selectors, then returns the matched element.

#### (*Page) Elements 

``` go
func (p *Page) Elements(selector string) (Elements, error)
```

Elements returns all elements that match the css selector.

#### (*Page) ElementsByJS 

``` go
func (p *Page) ElementsByJS(opts *EvalOptions) (Elements, error)
```

ElementsByJS returns the elements from the return value of the js.

#### (*Page) ElementsX 

``` go
func (p *Page) ElementsX(xpath string) (Elements, error)
```

ElementsX returns all elements that match the XPath selector.

#### (*Page) Emulate <- 0.42.1

``` go
func (p *Page) Emulate(device devices.Device) error
```

Emulate the device, such as iPhone9. If device is devices.Clear, it will clear the override.

#### (*Page) EnableDomain 

``` go
func (p *Page) EnableDomain(method proto.Request) (restore func())
```

EnableDomain and returns a restore function to restore previous state.

#### (*Page) Eval 

``` go
func (p *Page) Eval(js string, args ...interface{}) (*proto.RuntimeRemoteObject, error)
```

Eval is a shortcut for [Page.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Page.Evaluate) with AwaitPromise, ByValue set to true.

#### (*Page) EvalOnNewDocument <- 0.44.0

``` go
func (p *Page) EvalOnNewDocument(js string) (remove func() error, err error)
```

EvalOnNewDocument Evaluates given script in every frame upon creation (before loading frame's scripts).

#### (*Page) Evaluate <- 0.67.0

``` go
func (p *Page) Evaluate(opts *EvalOptions) (res *proto.RuntimeRemoteObject, err error)
```

Evaluate js on the page.

#### (*Page) Event <- 0.70.2

``` go
func (p *Page) Event() <-chan *Message
```

Event of the page.

#### (*Page) Expose <- 0.49.1

``` go
func (p *Page) Expose(name string, fn func(gson.JSON) (interface{}, error)) (stop func() error, err error)
```

Expose fn to the page's window object with the name. The exposure survives reloads. Call stop to unbind the fn.

#### (*Page) ExposeHelpers <- 0.85.1

``` go
func (p *Page) ExposeHelpers(list ...*js.Function)
```

ExposeHelpers helper functions to page's js context so that we can use the Devtools' console to debug them.

#### (*Page) GetContext 

``` go
func (p *Page) GetContext() context.Context
```

GetContext of current instance.

#### (*Page) GetNavigationHistory <- 0.116.2

``` go
func (p *Page) GetNavigationHistory() (*proto.PageGetNavigationHistoryResult, error)
```

GetNavigationHistory get navigation history.

#### (*Page) GetResource <- 0.76.6

``` go
func (p *Page) GetResource(url string) ([]byte, error)
```

GetResource content by the url. Such as image, css, html, etc. Use the [proto.PageGetResourceTree](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#PageGetResourceTree) to list all the resources.

#### (*Page) GetSessionID <- 0.72.0

``` go
func (p *Page) GetSessionID() proto.TargetSessionID
```

GetSessionID interface.

#### (*Page) GetWindow 

``` go
func (p *Page) GetWindow() (*proto.BrowserBounds, error)
```

GetWindow position and size info.

#### (*Page) HTML <- 0.94.0

``` go
func (p *Page) HTML() (string, error)
```

HTML of the page.

#### (*Page) HandleDialog 

``` go
func (p *Page) HandleDialog() (
	wait func() *proto.PageJavascriptDialogOpening,
	handle func(*proto.PageHandleJavaScriptDialog) error,
)
```

HandleDialog accepts or dismisses next JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload). Because modal dialog will block js, usually you have to trigger the dialog in another goroutine. For example:

```
wait, handle := page.MustHandleDialog()
go page.MustElement("button").MustClick()
wait()
handle(true, "")
```

#### (*Page) HandleFileDialog <- 0.109.0

``` go
func (p *Page) HandleFileDialog() (func([]string) error, error)
```

HandleFileDialog return a functions that waits for the next file chooser dialog pops up and returns the element for the event.

#### (*Page) Has 

``` go
func (p *Page) Has(selector string) (bool, *Element, error)
```

Has an element that matches the css selector.

#### (*Page) HasR <- 0.61.0

``` go
func (p *Page) HasR(selector, jsRegex string) (bool, *Element, error)
```

HasR an element that matches the css selector and its display text matches the jsRegex.

#### (*Page) HasX 

``` go
func (p *Page) HasX(selector string) (bool, *Element, error)
```

HasX an element that matches the XPath selector.

#### (*Page) HijackRequests 

``` go
func (p *Page) HijackRequests() *HijackRouter
```

HijackRequests creates a new router instance for requests hijacking. When use Fetch domain outside the router should be stopped. Enabling hijacking disables page caching, but such as 304 Not Modified will still work as expected. The entire process of hijacking one request:

```
browser --req-> rod ---> server ---> rod --res-> browser
```

The --req-> and --res-> are the parts that can be modified.

#### (*Page) Info <- 0.42.1

``` go
func (p *Page) Info() (*proto.TargetTargetInfo, error)
```

Info of the page, such as the URL or title of the page.

#### (*Page) InsertText <- 0.107.0

``` go
func (p *Page) InsertText(text string) error
```

InsertText is like pasting text into the page.

#### (*Page) IsIframe 

``` go
func (p *Page) IsIframe() bool
```

IsIframe tells if it's iframe.

#### (*Page) KeyActions <- 0.107.0

``` go
func (p *Page) KeyActions() *KeyActions
```

KeyActions simulates the type actions on a physical keyboard. Useful when input shortcuts like ctrl+enter .

#### (*Page) LoadState 

``` go
func (p *Page) LoadState(method proto.Request) (has bool)
```

LoadState into the method.

#### (*Page) MustActivate <- 0.86.3

``` go
func (p *Page) MustActivate() *Page
```

MustActivate is similar to [Page.Activate](https://pkg.go.dev/github.com/go-rod/rod#Page.Activate).

#### (*Page) MustAddScriptTag <- 0.50.0

``` go
func (p *Page) MustAddScriptTag(url string) *Page
```

MustAddScriptTag is similar to [Page.AddScriptTag](https://pkg.go.dev/github.com/go-rod/rod#Page.AddScriptTag).

#### (*Page) MustAddStyleTag <- 0.50.0

``` go
func (p *Page) MustAddStyleTag(url string) *Page
```

MustAddStyleTag is similar to [Page.AddStyleTag](https://pkg.go.dev/github.com/go-rod/rod#Page.AddStyleTag).

#### (*Page) MustCaptureDOMSnapshot <- 0.113.0

``` go
func (p *Page) MustCaptureDOMSnapshot() (domSnapshot *proto.DOMSnapshotCaptureSnapshotResult)
```

MustCaptureDOMSnapshot is similar to [Page.CaptureDOMSnapshot](https://pkg.go.dev/github.com/go-rod/rod#Page.CaptureDOMSnapshot).

#### (*Page) MustClose <- 0.50.0

``` go
func (p *Page) MustClose()
```

MustClose is similar to [Page.Close](https://pkg.go.dev/github.com/go-rod/rod#Page.Close).

#### (*Page) MustCookies <- 0.50.0

``` go
func (p *Page) MustCookies(urls ...string) []*proto.NetworkCookie
```

MustCookies is similar to [Page.Cookies](https://pkg.go.dev/github.com/go-rod/rod#Page.Cookies).

#### (*Page) MustElement <- 0.50.0

``` go
func (p *Page) MustElement(selector string) *Element
```

MustElement is similar to [Page.Element](https://pkg.go.dev/github.com/go-rod/rod#Page.Element).

#### (*Page) MustElementByJS <- 0.50.0

``` go
func (p *Page) MustElementByJS(js string, params ...interface{}) *Element
```

MustElementByJS is similar to [Page.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementByJS).

#### (*Page) MustElementFromNode <- 0.50.0

``` go
func (p *Page) MustElementFromNode(node *proto.DOMNode) *Element
```

MustElementFromNode is similar to [Page.ElementFromNode](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementFromNode).

#### (*Page) MustElementFromPoint <- 0.50.0

``` go
func (p *Page) MustElementFromPoint(left, top int) *Element
```

MustElementFromPoint is similar to [Page.ElementFromPoint](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementFromPoint).

#### (*Page) MustElementR <- 0.57.0

``` go
func (p *Page) MustElementR(selector, jsRegex string) *Element
```

MustElementR is similar to [Page.ElementR](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementR).

#### (*Page) MustElementX <- 0.50.0

``` go
func (p *Page) MustElementX(xPath string) *Element
```

MustElementX is similar to [Page.ElementX](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementX).

#### (*Page) MustElements <- 0.50.0

``` go
func (p *Page) MustElements(selector string) Elements
```

MustElements is similar to [Page.Elements](https://pkg.go.dev/github.com/go-rod/rod#Page.Elements).

#### (*Page) MustElementsByJS <- 0.50.0

``` go
func (p *Page) MustElementsByJS(js string, params ...interface{}) Elements
```

MustElementsByJS is similar to [Page.ElementsByJS](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementsByJS).

#### (*Page) MustElementsX <- 0.50.0

``` go
func (p *Page) MustElementsX(xpath string) Elements
```

MustElementsX is similar to [Page.ElementsX](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementsX).

#### (*Page) MustEmulate <- 0.50.0

``` go
func (p *Page) MustEmulate(device devices.Device) *Page
```

MustEmulate is similar to [Page.Emulate](https://pkg.go.dev/github.com/go-rod/rod#Page.Emulate).

#### (*Page) MustEval <- 0.50.0

``` go
func (p *Page) MustEval(js string, params ...interface{}) gson.JSON
```

MustEval is similar to [Page.Eval](https://pkg.go.dev/github.com/go-rod/rod#Page.Eval).

#### (*Page) MustEvalOnNewDocument <- 0.50.0

``` go
func (p *Page) MustEvalOnNewDocument(js string)
```

MustEvalOnNewDocument is similar to [Page.EvalOnNewDocument](https://pkg.go.dev/github.com/go-rod/rod#Page.EvalOnNewDocument).

#### (*Page) MustEvaluate <- 0.67.0

``` go
func (p *Page) MustEvaluate(opts *EvalOptions) *proto.RuntimeRemoteObject
```

MustEvaluate is similar to [Page.Evaluate](https://pkg.go.dev/github.com/go-rod/rod#Page.Evaluate).

#### (*Page) MustExpose <- 0.50.0

``` go
func (p *Page) MustExpose(name string, fn func(gson.JSON) (interface{}, error)) (stop func())
```

MustExpose is similar to [Page.Expose](https://pkg.go.dev/github.com/go-rod/rod#Page.Expose).

#### (*Page) MustGetWindow <- 0.50.0

``` go
func (p *Page) MustGetWindow() *proto.BrowserBounds
```

MustGetWindow is similar to [Page.GetWindow](https://pkg.go.dev/github.com/go-rod/rod#Page.GetWindow).

#### (*Page) MustHTML <- 0.94.0

``` go
func (p *Page) MustHTML() string
```

MustHTML is similar to [Page.HTML](https://pkg.go.dev/github.com/go-rod/rod#Page.HTML).

#### (*Page) MustHandleDialog <- 0.50.0

``` go
func (p *Page) MustHandleDialog() (wait func() *proto.PageJavascriptDialogOpening, handle func(bool, string))
```

MustHandleDialog is similar to [Page.HandleDialog](https://pkg.go.dev/github.com/go-rod/rod#Page.HandleDialog).

#### (*Page) MustHandleFileDialog <- 0.109.0

``` go
func (p *Page) MustHandleFileDialog() func(...string)
```

MustHandleFileDialog is similar to [Page.HandleFileDialog](https://pkg.go.dev/github.com/go-rod/rod#Page.HandleFileDialog).

#### (*Page) MustHas <- 0.50.0

``` go
func (p *Page) MustHas(selector string) bool
```

MustHas is similar to [Page.Has](https://pkg.go.dev/github.com/go-rod/rod#Page.Has).

#### (*Page) MustHasR <- 0.61.0

``` go
func (p *Page) MustHasR(selector, regex string) bool
```

MustHasR is similar to [Page.HasR](https://pkg.go.dev/github.com/go-rod/rod#Page.HasR).

#### (*Page) MustHasX <- 0.50.0

``` go
func (p *Page) MustHasX(selector string) bool
```

MustHasX is similar to [Page.HasX](https://pkg.go.dev/github.com/go-rod/rod#Page.HasX).

#### (*Page) MustInfo <- 0.50.0

``` go
func (p *Page) MustInfo() *proto.TargetTargetInfo
```

MustInfo is similar to [Page.Info](https://pkg.go.dev/github.com/go-rod/rod#Page.Info).

#### (*Page) MustInsertText <- 0.107.0

``` go
func (p *Page) MustInsertText(text string) *Page
```

MustInsertText is similar to [Page.InsertText](https://pkg.go.dev/github.com/go-rod/rod#Page.InsertText).

#### (*Page) MustNavigate <- 0.50.0

``` go
func (p *Page) MustNavigate(url string) *Page
```

MustNavigate is similar to [Page.Navigate](https://pkg.go.dev/github.com/go-rod/rod#Page.Navigate).

#### (*Page) MustNavigateBack <- 0.61.4

``` go
func (p *Page) MustNavigateBack() *Page
```

MustNavigateBack is similar to [Page.NavigateBack](https://pkg.go.dev/github.com/go-rod/rod#Page.NavigateBack).

#### (*Page) MustNavigateForward <- 0.61.4

``` go
func (p *Page) MustNavigateForward() *Page
```

MustNavigateForward is similar to [Page.NavigateForward](https://pkg.go.dev/github.com/go-rod/rod#Page.NavigateForward).

#### (*Page) MustObjectToJSON <- 0.50.0

``` go
func (p *Page) MustObjectToJSON(obj *proto.RuntimeRemoteObject) gson.JSON
```

MustObjectToJSON is similar to [Page.ObjectToJSON](https://pkg.go.dev/github.com/go-rod/rod#Page.ObjectToJSON).

#### (*Page) MustObjectsToJSON <- 0.50.0

``` go
func (p *Page) MustObjectsToJSON(list []*proto.RuntimeRemoteObject) gson.JSON
```

MustObjectsToJSON is similar to [Page.ObjectsToJSON].

#### (*Page) MustPDF <- 0.50.0

``` go
func (p *Page) MustPDF(toFile ...string) []byte
```

MustPDF is similar to [Page.PDF](https://pkg.go.dev/github.com/go-rod/rod#Page.PDF). If the toFile is "", it Page.will save output to "tmp/pdf" folder, time as the file name.

#### (*Page) MustRelease <- 0.50.0

``` go
func (p *Page) MustRelease(obj *proto.RuntimeRemoteObject) *Page
```

MustRelease is similar to [Page.Release](https://pkg.go.dev/github.com/go-rod/rod#Page.Release).

#### (*Page) MustReload <- 0.61.4

``` go
func (p *Page) MustReload() *Page
```

MustReload is similar to [Page.Reload](https://pkg.go.dev/github.com/go-rod/rod#Page.Reload).

#### (*Page) MustResetNavigationHistory <- 0.116.2

``` go
func (p *Page) MustResetNavigationHistory() *Page
```

MustResetNavigationHistory is similar to [Page.ResetNavigationHistory](https://pkg.go.dev/github.com/go-rod/rod#Page.ResetNavigationHistory).

#### (*Page) MustScreenshot <- 0.50.0

``` go
func (p *Page) MustScreenshot(toFile ...string) []byte
```

MustScreenshot is similar to [Page.Screenshot](https://pkg.go.dev/github.com/go-rod/rod#Page.Screenshot). If the toFile is "", it Page.will save output to "tmp/screenshots" folder, time as the file name.

#### (*Page) MustScreenshotFullPage <- 0.50.0

``` go
func (p *Page) MustScreenshotFullPage(toFile ...string) []byte
```

MustScreenshotFullPage is similar to [Page.ScreenshotFullPage]. If the toFile is "", it Page.will save output to "tmp/screenshots" folder, time as the file name.

#### (*Page) MustScrollScreenshot <- 0.116.2

``` go
func (p *Page) MustScrollScreenshot(toFile ...string) []byte
```

MustScrollScreenshot is similar to [Page.ScrollScreenshot](https://pkg.go.dev/github.com/go-rod/rod#Page.ScrollScreenshot). If the toFile is "", it Page.will save output to "tmp/screenshots" folder, time as the file name.

#### (*Page) MustSearch <- 0.50.0

``` go
func (p *Page) MustSearch(query string) *Element
```

MustSearch is similar to [Page.Search](https://pkg.go.dev/github.com/go-rod/rod#Page.Search). It only returns the first element in the search result.

#### (*Page) MustSetBlockedURLs <- 0.112.3

``` go
func (p *Page) MustSetBlockedURLs(urls ...string) *Page
```

MustSetBlockedURLs is similar to [Page.SetBlockedURLs](https://pkg.go.dev/github.com/go-rod/rod#Page.SetBlockedURLs).

#### (*Page) MustSetCookies <- 0.50.0

``` go
func (p *Page) MustSetCookies(cookies ...*proto.NetworkCookieParam) *Page
```

MustSetCookies is similar to [Page.SetCookies](https://pkg.go.dev/github.com/go-rod/rod#Page.SetCookies). If the len(cookies) is 0 it will clear all the cookies.

#### (*Page) MustSetDocumentContent <- 0.104.0

``` go
func (p *Page) MustSetDocumentContent(html string) *Page
```

MustSetDocumentContent is similar to [Page.SetDocumentContent](https://pkg.go.dev/github.com/go-rod/rod#Page.SetDocumentContent).

#### (*Page) MustSetExtraHeaders <- 0.50.0

``` go
func (p *Page) MustSetExtraHeaders(dict ...string) (cleanup func())
```

MustSetExtraHeaders is similar to [Page.SetExtraHeaders](https://pkg.go.dev/github.com/go-rod/rod#Page.SetExtraHeaders).

#### (*Page) MustSetUserAgent <- 0.50.0

``` go
func (p *Page) MustSetUserAgent(req *proto.NetworkSetUserAgentOverride) *Page
```

MustSetUserAgent is similar to [Page.SetUserAgent](https://pkg.go.dev/github.com/go-rod/rod#Page.SetUserAgent).

#### (*Page) MustSetViewport <- 0.64.0

``` go
func (p *Page) MustSetViewport(width, height int, deviceScaleFactor float64, mobile bool) *Page
```

MustSetViewport is similar to [Page.SetViewport](https://pkg.go.dev/github.com/go-rod/rod#Page.SetViewport).

#### (*Page) MustSetWindow <- 0.64.0

``` go
func (p *Page) MustSetWindow(left, top, width, height int) *Page
```

MustSetWindow is similar to [Page.SetWindow](https://pkg.go.dev/github.com/go-rod/rod#Page.SetWindow).

#### (*Page) MustStopLoading <- 0.50.0

``` go
func (p *Page) MustStopLoading() *Page
```

MustStopLoading is similar to [Page.StopLoading](https://pkg.go.dev/github.com/go-rod/rod#Page.StopLoading).

#### (*Page) MustTriggerFavicon <- 0.113.2

``` go
func (p *Page) MustTriggerFavicon() *Page
```

MustTriggerFavicon is similar to [PageTriggerFavicon].

#### (*Page) MustWait <- 0.50.0

``` go
func (p *Page) MustWait(js string, params ...interface{}) *Page
```

MustWait is similar to [Page.Wait](https://pkg.go.dev/github.com/go-rod/rod#Page.Wait).

#### (*Page) MustWaitDOMStable <- 0.114.0

``` go
func (p *Page) MustWaitDOMStable() *Page
```

MustWaitDOMStable is similar to [Page.WaitDOMStable](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitDOMStable).

#### (*Page) MustWaitElementsMoreThan <- 0.97.3

``` go
func (p *Page) MustWaitElementsMoreThan(selector string, num int) *Page
```

MustWaitElementsMoreThan is similar to [Page.WaitElementsMoreThan](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitElementsMoreThan).

#### (*Page) MustWaitIdle <- 0.50.0

``` go
func (p *Page) MustWaitIdle() *Page
```

MustWaitIdle is similar to [Page.WaitIdle](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitIdle).

#### (*Page) MustWaitLoad <- 0.50.0

``` go
func (p *Page) MustWaitLoad() *Page
```

MustWaitLoad is similar to [Page.WaitLoad](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitLoad).

#### (*Page) MustWaitNavigation <- 0.63.2

``` go
func (p *Page) MustWaitNavigation() func()
```

MustWaitNavigation is similar to [Page.WaitNavigation](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitNavigation).

#### (*Page) MustWaitOpen <- 0.50.0

``` go
func (p *Page) MustWaitOpen() (wait func() (newPage *Page))
```

MustWaitOpen is similar to [Page.WaitOpen](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitOpen).

#### (*Page) MustWaitRequestIdle <- 0.50.0

``` go
func (p *Page) MustWaitRequestIdle(excludes ...string) (wait func())
```

MustWaitRequestIdle is similar to [Page.WaitRequestIdle](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitRequestIdle).

#### (*Page) MustWaitStable <- 0.113.0

``` go
func (p *Page) MustWaitStable() *Page
```

MustWaitStable is similar to [Page.WaitStable](https://pkg.go.dev/github.com/go-rod/rod#Page.WaitStable).

#### (*Page) MustWindowFullscreen <- 0.50.0

``` go
func (p *Page) MustWindowFullscreen() *Page
```

MustWindowFullscreen is similar to [Page.WindowFullscreen].

#### (*Page) MustWindowMaximize <- 0.50.0

``` go
func (p *Page) MustWindowMaximize() *Page
```

MustWindowMaximize is similar to [Page.WindowMaximize].

#### (*Page) MustWindowMinimize <- 0.50.0

``` go
func (p *Page) MustWindowMinimize() *Page
```

MustWindowMinimize is similar to [Page.WindowMinimize].

#### (*Page) MustWindowNormal <- 0.50.0

``` go
func (p *Page) MustWindowNormal() *Page
```

MustWindowNormal is similar to [Page.WindowNormal].

#### (*Page) Navigate 

``` go
func (p *Page) Navigate(url string) error
```

Navigate to the url. If the url is empty, "about:blank" will be used. It will return immediately after the server responds the http header.

#### (*Page) NavigateBack <- 0.61.4

``` go
func (p *Page) NavigateBack() error
```

NavigateBack history.

#### (*Page) NavigateForward <- 0.61.4

``` go
func (p *Page) NavigateForward() error
```

NavigateForward history.

#### (*Page) ObjectToJSON 

``` go
func (p *Page) ObjectToJSON(obj *proto.RuntimeRemoteObject) (gson.JSON, error)
```

ObjectToJSON by object id.

#### (*Page) Overlay 

``` go
func (p *Page) Overlay(left, top, width, height float64, msg string) (remove func())
```

Overlay a rectangle on the main frame with specified message.

#### (*Page) PDF 

``` go
func (p *Page) PDF(req *proto.PagePrintToPDF) (*StreamReader, error)
```

PDF prints page as PDF.

#### (*Page) Race <- 0.57.0

``` go
func (p *Page) Race() *RaceContext
```

Race creates a context to race selectors.

#### (*Page) Release 

``` go
func (p *Page) Release(obj *proto.RuntimeRemoteObject) error
```

Release the remote object. Usually, you don't need to call it. When a page is closed or reloaded, all remote objects will be released automatically. It's useful if the page never closes or reloads.

#### (*Page) Reload <- 0.61.4

``` go
func (p *Page) Reload() error
```

Reload page.

#### (*Page) ResetNavigationHistory <- 0.116.2

``` go
func (p *Page) ResetNavigationHistory() error
```

ResetNavigationHistory reset history.

#### (*Page) Screenshot 

``` go
func (p *Page) Screenshot(fullPage bool, req *proto.PageCaptureScreenshot) ([]byte, error)
```

Screenshot captures the screenshot of current page.

#### (*Page) ScrollScreenshot <- 0.114.7

``` go
func (p *Page) ScrollScreenshot(opt *ScrollScreenshotOptions) ([]byte, error)
```

ScrollScreenshot Scroll screenshot does not adjust the size of the viewport, but achieves it by scrolling and capturing screenshots in a loop, and then stitching them together. Note that this method also has a flaw: when there are elements with fixed positioning on the page (usually header navigation components), these elements will appear repeatedly, you can set the FixedTop parameter to optimize it.

Only support png and jpeg format yet, webP is not supported because no suitable processing library was found in golang.

#### (*Page) Search <- 0.47.0

``` go
func (p *Page) Search(query string) (*SearchResult, error)
```

Search for the given query in the DOM tree until the result count is not zero, before that it will keep retrying. The query can be plain text or css selector or xpath. It will search nested iframes and shadow doms too.

#### (*Page) SetBlockedURLs <- 0.112.3

``` go
func (p *Page) SetBlockedURLs(urls []string) error
```

SetBlockedURLs For some requests that do not want to be triggered, such as some dangerous operations, delete, quit logout, etc. Wildcards ('*') are allowed, such as ["*/api/logout/*","delete"]. NOTE: if you set empty pattern "", it will block all requests.

#### (*Page) SetCookies 

``` go
func (p *Page) SetCookies(cookies []*proto.NetworkCookieParam) error
```

SetCookies is similar to Browser.SetCookies .

#### (*Page) SetDocumentContent <- 0.104.0

``` go
func (p *Page) SetDocumentContent(html string) error
```

SetDocumentContent sets the page document html content.

#### (*Page) SetExtraHeaders 

``` go
func (p *Page) SetExtraHeaders(dict []string) (func(), error)
```

SetExtraHeaders whether to always send extra HTTP headers with the requests from this page.

#### (*Page) SetUserAgent 

``` go
func (p *Page) SetUserAgent(req *proto.NetworkSetUserAgentOverride) error
```

SetUserAgent (browser brand, accept-language, etc) of the page. If req is nil, a default user agent will be used, a typical mac chrome.

#### (*Page) SetViewport <- 0.62.0

``` go
func (p *Page) SetViewport(params *proto.EmulationSetDeviceMetricsOverride) error
```

SetViewport overrides the values of device screen dimensions.

#### (*Page) SetWindow <- 0.62.0

``` go
func (p *Page) SetWindow(bounds *proto.BrowserBounds) error
```

SetWindow location and size.

#### (*Page) Sleeper 

``` go
func (p *Page) Sleeper(sleeper func() utils.Sleeper) *Page
```

Sleeper returns a clone with the specified sleeper for chained sub-operations.

#### (*Page) StopLoading 

``` go
func (p *Page) StopLoading() error
```

StopLoading forces the page stop navigation and pending resource fetches.

#### (*Page) String <- 0.88.0

``` go
func (p *Page) String() string
```

String interface.

#### (*Page) Timeout 

``` go
func (p *Page) Timeout(d time.Duration) *Page
```

Timeout returns a clone with the specified total timeout of all chained sub-operations.

#### (*Page) TriggerFavicon <- 0.113.2

``` go
func (p *Page) TriggerFavicon() error
```

TriggerFavicon supports when browser in headless mode to trigger favicon's request. Pay attention to this function only supported when browser in headless mode, if you call it in no-headless mode, it will raise an error with the message "browser is no-headless".

#### (*Page) Wait 

``` go
func (p *Page) Wait(opts *EvalOptions) error
```

Wait until the js returns true.

#### (*Page) WaitDOMStable <- 0.114.0

``` go
func (p *Page) WaitDOMStable(d time.Duration, diff float64) error
```

WaitDOMStable waits until the change of the DOM tree is less or equal than diff percent for d duration. Be careful, d is not the max wait timeout, it's the least stable time. If you want to set a timeout you can use the [Page.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Page.Timeout) function.

#### (*Page) WaitElementsMoreThan <- 0.97.3

``` go
func (p *Page) WaitElementsMoreThan(selector string, num int) error
```

WaitElementsMoreThan waits until there are more than num elements that match the selector.

#### (*Page) WaitEvent 

``` go
func (p *Page) WaitEvent(e proto.Event) (wait func())
```

WaitEvent waits for the next event for one time. It will also load the data into the event object.

#### (*Page) WaitIdle 

``` go
func (p *Page) WaitIdle(timeout time.Duration) (err error)
```

WaitIdle waits until the next window.requestIdleCallback is called.

#### (*Page) WaitLoad 

``` go
func (p *Page) WaitLoad() error
```

WaitLoad waits for the `window.onload` event, it returns immediately if the event is already fired.

#### (*Page) WaitNavigation <- 0.63.2

``` go
func (p *Page) WaitNavigation(name proto.PageLifecycleEventName) func()
```

WaitNavigation wait for a page lifecycle event when navigating. Usually you will wait for [proto.PageLifecycleEventNameNetworkAlmostIdle](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#PageLifecycleEventNameNetworkAlmostIdle).

#### (*Page) WaitOpen 

``` go
func (p *Page) WaitOpen() func() (*Page, error)
```

WaitOpen waits for the next new page opened by the current one.

#### (*Page) WaitRepaint <- 0.84.1

``` go
func (p *Page) WaitRepaint() error
```

WaitRepaint waits until the next repaint. Doc: https://developer.mozilla.org/en-US/docs/Web/API/window/requestAnimationFrame

#### (*Page) WaitRequestIdle 

``` go
func (p *Page) WaitRequestIdle(
	d time.Duration,
	includes, excludes []string,
	excludeTypes []proto.NetworkResourceType,
) func()
```

WaitRequestIdle returns a wait function that waits until no request for d duration. Be careful, d is not the max wait timeout, it's the least idle time. If you want to set a timeout you can use the [Page.Timeout](https://pkg.go.dev/github.com/go-rod/rod#Page.Timeout) function. Use the includes and excludes regexp list to filter the requests by their url.

#### (*Page) WaitStable <- 0.113.0

``` go
func (p *Page) WaitStable(d time.Duration) error
```

WaitStable waits until the page is stable for d duration.

#### (*Page) WithCancel <- 0.69.0

``` go
func (p *Page) WithCancel() (*Page, func())
```

WithCancel returns a clone with a context cancel function.

#### (*Page) WithPanic <- 0.100.0

``` go
func (p *Page) WithPanic(fail func(interface{})) *Page
```

WithPanic returns a page clone with the specified panic function. The fail must stop the current goroutine's execution immediately, such as use [runtime.Goexit](https://pkg.go.dev/runtime#Goexit) or panic inside it.

#### type PageCloseCanceledError <- 0.114.8

``` go
type PageCloseCanceledError struct{}
```

PageCloseCanceledError error.

#### (*PageCloseCanceledError) Error <- 0.114.8

``` go
func (e *PageCloseCanceledError) Error() string
```

#### type PageNotFoundError <- 0.114.8

``` go
type PageNotFoundError struct{}
```

PageNotFoundError error.

#### (*PageNotFoundError) Error <- 0.114.8

``` go
func (e *PageNotFoundError) Error() string
```

#### type Pages 

``` go
type Pages []*Page
```

Pages provides some helpers to deal with page list.

#### (Pages) Empty <- 0.53.0

``` go
func (ps Pages) Empty() bool
```

Empty returns true if the list is empty.

#### (Pages) Find 

``` go
func (ps Pages) Find(selector string) (*Page, error)
```

Find the page that has the specified element with the css selector.

#### (Pages) FindByURL 

``` go
func (ps Pages) FindByURL(jsRegex string) (*Page, error)
```

FindByURL returns the page that has the url that matches the jsRegex.

#### (Pages) First <- 0.53.0

``` go
func (ps Pages) First() *Page
```

First returns the first page, if the list is empty returns nil.

#### (Pages) Last <- 0.53.0

``` go
func (ps Pages) Last() *Page
```

Last returns the last page, if the list is empty returns nil.

#### (Pages) MustFind <- 0.50.3

``` go
func (ps Pages) MustFind(selector string) *Page
```

MustFind is similar to [Browser.Find].

#### (Pages) MustFindByURL <- 0.50.0

``` go
func (ps Pages) MustFindByURL(regex string) *Page
```

MustFindByURL is similar to [Page.FindByURL].

#### type Pool <- 0.116.2

``` go
type Pool[T any] chan *T
```

Pool is used to thread-safely limit the number of elements at the same time. It's a common practice to use a channel to limit concurrency, it's not special for rod. This helper is more like an example to use Go Channel. Reference: https://golang.org/doc/effective_go#channels

#### func NewBrowserPool <- 0.101.7

``` go
func NewBrowserPool(limit int) Pool[Browser]
```

NewBrowserPool instance.

#### func NewPagePool <- 0.73.2

``` go
func NewPagePool(limit int) Pool[Page]
```

NewPagePool instance.

#### func NewPool <- 0.116.2

``` go
func NewPool[T any](limit int) Pool[T]
```

NewPool instance.

#### (Pool[T]) Cleanup <- 0.116.2

``` go
func (p Pool[T]) Cleanup(iteratee func(*T))
```

Cleanup helper.

#### (Pool[T]) Get <- 0.116.2

``` go
func (p Pool[T]) Get(create func() (*T, error)) (elem *T, err error)
```

Get a elem from the pool, allow error. Use the [Pool[T].Put] to make it reusable later.

#### (Pool[T]) MustGet <- 0.116.2

``` go
func (p Pool[T]) MustGet(create func() *T) *T
```

MustGet an elem from the pool. Use the [Pool[T].Put] to make it reusable later.

#### (Pool[T]) Put <- 0.116.2

``` go
func (p Pool[T]) Put(elem *T)
```

Put an elem back to the pool.

#### type RaceContext <- 0.57.0

``` go
type RaceContext struct {
	// contains filtered or unexported fields
}
```

RaceContext stores the branches to race.

#### (*RaceContext) Do <- 0.57.0

``` go
func (rc *RaceContext) Do() (*Element, error)
```

Do the race.

#### (*RaceContext) Element <- 0.57.0

``` go
func (rc *RaceContext) Element(selector string) *RaceContext
```

Element is similar to [Page.Element](https://pkg.go.dev/github.com/go-rod/rod#Page.Element).

#### (*RaceContext) ElementByJS <- 0.57.0

``` go
func (rc *RaceContext) ElementByJS(opts *EvalOptions) *RaceContext
```

ElementByJS is similar to [Page.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementByJS).

#### (*RaceContext) ElementFunc <- 0.107.1

``` go
func (rc *RaceContext) ElementFunc(fn func(*Page) (*Element, error)) *RaceContext
```

ElementFunc takes a custom function to determine race success.

#### (*RaceContext) ElementR <- 0.57.0

``` go
func (rc *RaceContext) ElementR(selector, regex string) *RaceContext
```

ElementR is similar to [Page.ElementR](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementR).

#### (*RaceContext) ElementX <- 0.57.0

``` go
func (rc *RaceContext) ElementX(selector string) *RaceContext
```

ElementX is similar to [Page.ElementX](https://pkg.go.dev/github.com/go-rod/rod#Page.ElementX).

#### (*RaceContext) Handle <- 0.81.0

``` go
func (rc *RaceContext) Handle(callback func(*Element) error) *RaceContext
```

Handle adds a callback function to the most recent chained selector. The callback function is run, if the corresponding selector is present first, in the Race condition.

#### (*RaceContext) MustDo <- 0.57.0

``` go
func (rc *RaceContext) MustDo() *Element
```

MustDo is similar to [RaceContext.Do](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.Do).

#### (*RaceContext) MustElementByJS <- 0.57.0

``` go
func (rc *RaceContext) MustElementByJS(js string, params []interface{}) *RaceContext
```

MustElementByJS is similar to [RaceContext.ElementByJS](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.ElementByJS).

#### (*RaceContext) MustHandle <- 0.81.0

``` go
func (rc *RaceContext) MustHandle(callback func(*Element)) *RaceContext
```

MustHandle is similar to [RaceContext.Handle](https://pkg.go.dev/github.com/go-rod/rod#RaceContext.Handle).

#### (*RaceContext) Search <- 0.112.0

``` go
func (rc *RaceContext) Search(query string) *RaceContext
```

Search is similar to [Page.Search](https://pkg.go.dev/github.com/go-rod/rod#Page.Search).

#### type ScrollScreenshotOptions <- 0.114.7

``` go
type ScrollScreenshotOptions struct {
	// Format (optional) Image compression format (defaults to png).
	Format proto.PageCaptureScreenshotFormat `json:"format,omitempty"`

	// Quality (optional) Compression quality from range [0..100] (jpeg only).
	Quality *int `json:"quality,omitempty"`

	// FixedTop (optional) The number of pixels to skip from the top.
	// It is suitable for optimizing the screenshot effect when there is a fixed
	// positioning element at the top of the page.
	FixedTop float64

	// FixedBottom (optional) The number of pixels to skip from the bottom.
	FixedBottom float64

	// WaitPerScroll until no animation (default is 300ms)
	WaitPerScroll time.Duration
}
```

ScrollScreenshotOptions is the options for the ScrollScreenshot.

#### type SearchResult <- 0.97.0

``` go
type SearchResult struct {
	*proto.DOMPerformSearchResult

	// First element in the search result
	First *Element
	// contains filtered or unexported fields
}
```

SearchResult handler.

#### (*SearchResult) All <- 0.97.0

``` go
func (s *SearchResult) All() (Elements, error)
```

All returns all elements.

#### (*SearchResult) Get <- 0.97.0

``` go
func (s *SearchResult) Get(i, l int) (Elements, error)
```

Get l elements at the index of i from the remote search result.

#### (*SearchResult) Release <- 0.97.0

``` go
func (s *SearchResult) Release()
```

Release the remote search result.

#### type SelectorType <- 0.68.0

``` go
type SelectorType string
```

SelectorType enum.

``` go
const (
	// SelectorTypeRegex type.
	SelectorTypeRegex SelectorType = "regex"
	// SelectorTypeCSSSector type.
	SelectorTypeCSSSector SelectorType = "css-selector"
	// SelectorTypeText type.
	SelectorTypeText SelectorType = "text"
)
```

#### type StreamReader <- 0.63.0

``` go
type StreamReader struct {
	Offset *int
	// contains filtered or unexported fields
}
```

StreamReader for browser data stream.

#### func NewStreamReader <- 0.63.0

``` go
func NewStreamReader(c proto.Client, h proto.IOStreamHandle) *StreamReader
```

NewStreamReader instance.

#### (*StreamReader) Close <- 0.102.0

``` go
func (sr *StreamReader) Close() error
```

Close the stream, discard any temporary backing storage.

#### (*StreamReader) Read <- 0.63.0

``` go
func (sr *StreamReader) Read(p []byte) (n int, err error)
```

#### type Touch <- 0.61.1

``` go
type Touch struct {
	// contains filtered or unexported fields
}
```

Touch presents a touch device, such as a hand with fingers, each finger is a [proto.InputTouchPoint](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#InputTouchPoint). Touch events is stateless, we use the struct here only as a namespace to make the API style unified.

#### (*Touch) Cancel <- 0.61.1

``` go
func (t *Touch) Cancel() error
```

Cancel touch action.

#### (*Touch) End <- 0.61.1

``` go
func (t *Touch) End() error
```

End touch action.

#### (*Touch) Move <- 0.61.1

``` go
func (t *Touch) Move(points ...*proto.InputTouchPoint) error
```

Move touch points. Use the [proto.InputTouchPoint.ID](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/proto#InputTouchPoint.ID) (Touch.identifier) to track points. Doc: https://developer.mozilla.org/en-US/docs/Web/API/Touch_events

#### (*Touch) MustCancel <- 0.61.1

``` go
func (t *Touch) MustCancel() *Touch
```

MustCancel is similar to [Touch.Cancel](https://pkg.go.dev/github.com/go-rod/rod#Touch.Cancel).

#### (*Touch) MustEnd <- 0.61.1

``` go
func (t *Touch) MustEnd() *Touch
```

MustEnd is similar to [Touch.End](https://pkg.go.dev/github.com/go-rod/rod#Touch.End).

#### (*Touch) MustMove <- 0.61.1

``` go
func (t *Touch) MustMove(points ...*proto.InputTouchPoint) *Touch
```

MustMove is similar to [Touch.Move](https://pkg.go.dev/github.com/go-rod/rod#Touch.Move).

#### (*Touch) MustStart <- 0.61.1

``` go
func (t *Touch) MustStart(points ...*proto.InputTouchPoint) *Touch
```

MustStart is similar to [Touch.Start](https://pkg.go.dev/github.com/go-rod/rod#Touch.Start).

#### (*Touch) MustTap <- 0.61.1

``` go
func (t *Touch) MustTap(x, y float64) *Touch
```

MustTap is similar to [Touch.Tap](https://pkg.go.dev/github.com/go-rod/rod#Touch.Tap).

#### (*Touch) Start <- 0.61.1

``` go
func (t *Touch) Start(points ...*proto.InputTouchPoint) error
```

Start a touch action.

#### (*Touch) Tap <- 0.61.1

``` go
func (t *Touch) Tap(x, y float64) error
```

Tap dispatches a touchstart and touchend event.

#### type TraceType <- 0.59.0

``` go
type TraceType string
```

TraceType for logger.

``` go
const (
	// TraceTypeWaitRequestsIdle type.
	TraceTypeWaitRequestsIdle TraceType = "wait requests idle"

	// TraceTypeWaitRequests type.
	TraceTypeWaitRequests TraceType = "wait requests"

	// TraceTypeQuery type.
	TraceTypeQuery TraceType = "query"

	// TraceTypeWait type.
	TraceTypeWait TraceType = "wait"

	// TraceTypeInput type.
	TraceTypeInput TraceType = "input"
)
```

#### (TraceType) String <- 0.88.0

``` go
func (t TraceType) String() string
```

String interface.

#### type TryError <- 0.114.8

``` go
type TryError struct {
	Value interface{}
	Stack string
}
```

TryError error.

#### (*TryError) Error <- 0.114.8

``` go
func (e *TryError) Error() string
```

#### (*TryError) Is <- 0.114.8

``` go
func (e *TryError) Is(err error) bool
```

Is interface.

#### (*TryError) Unwrap <- 0.114.8

``` go
func (e *TryError) Unwrap() error
```

Unwrap stdlib interface.
