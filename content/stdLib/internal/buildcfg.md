# buildcfg

https://pkg.go.dev/internal/buildcfg@go1.20.1



Package buildcfg provides access to the build configuration described by the current environment. It is for use by build tools such as cmd/go or cmd/compile and for setting up go/build's Default context.

Note that it does NOT provide access to the build configuration used to build the currently-running binary. For that, use runtime.GOOS etc as well as internal/goexperiment.











  
  

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/buildcfg/exp.go;l=42)

```
const DefaultGOEXPERIMENT = defaultGOEXPERIMENT
```

DefaultGOEXPERIMENT is the embedded default GOEXPERIMENT string. It is not guaranteed to be canonical.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/buildcfg/cfg.go;l=22)

```
var (
	GOROOT   = runtime.GOROOT() // cached for efficiency
	GOARCH   = envOr("GOARCH", defaultGOARCH)
	GOOS     = envOr("GOOS", defaultGOOS)
	GO386    = envOr("GO386", defaultGO386)
	GOAMD64  = goamd64()
	GOARM    = goarm()
	GOMIPS   = gomips()
	GOMIPS64 = gomips64()
	GOPPC64  = goppc64()
	GOWASM   = gowasm()
	ToolTags = toolTags()
	GO_LDSO  = defaultGO_LDSO
	Version  = version
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/buildcfg/cfg.go;l=39)

```
var Error error
```

Error is one of the errors found (if any) in the build configuration.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/buildcfg/exp.go;l=51)

```
var FramePointerEnabled = GOARCH == "amd64" || GOARCH == "arm64"
```

FramePointerEnabled enables the use of platform conventions for saving frame pointers.

This used to be an experiment, but now it's always enabled on platforms that support it.

Note: must agree with runtime.framepointer_enabled.

## 函数

#### func Check 

```
func Check()
```

Check exits the program with a fatal error if Error is non-nil.

#### func GOGOARCH  <- go1.20

```
func GOGOARCH() (name, value string)
```

GOGOARCH returns the name and value of the GO$GOARCH setting. For example, if GOARCH is "amd64" it might return "GOAMD64", "v2".

#### func Getgoextlinkenabled 

```
func Getgoextlinkenabled() string
```

## 类型

### type ExperimentFlags  <- go1.19

```
type ExperimentFlags struct {
	goexperiment.Flags
	// contains filtered or unexported fields
}
```

ExperimentFlags represents a set of GOEXPERIMENT flags relative to a baseline (platform-default) experiment configuration.

```
var Experiment ExperimentFlags = func() ExperimentFlags {
	flags, err := ParseGOEXPERIMENT(GOOS, GOARCH, envOr("GOEXPERIMENT", defaultGOEXPERIMENT))
	if err != nil {
		Error = err
		return ExperimentFlags{}
	}
	return *flags
}()
```

Experiment contains the toolchain experiments enabled for the current build.

(This is not necessarily the set of experiments the compiler itself was built with.)

experimentBaseline specifies the experiment flags that are enabled by default in the current toolchain. This is, in effect, the "control" configuration and any variation from this is an experiment.

#### func ParseGOEXPERIMENT 

```
func ParseGOEXPERIMENT(goos, goarch, goexp string) (*ExperimentFlags, error)
```

ParseGOEXPERIMENT parses a (GOOS, GOARCH, GOEXPERIMENT) configuration tuple and returns the enabled and baseline experiment flag sets.

TODO(mdempsky): Move to internal/goexperiment.

#### (*ExperimentFlags) All  <- go1.19

```
func (exp *ExperimentFlags) All() []string
```

All returns a list of all experiment settings. Disabled experiments appear in the list prefixed by "no".

#### (*ExperimentFlags) Enabled  <- go1.19

```
func (exp *ExperimentFlags) Enabled() []string
```

Enabled returns a list of enabled experiments, as lower-cased experiment names.

#### (*ExperimentFlags) String  <- go1.19

```
func (exp *ExperimentFlags) String() string
```

String returns the canonical GOEXPERIMENT string to enable this experiment configuration. (Experiments in the same state as in the baseline are elided.)