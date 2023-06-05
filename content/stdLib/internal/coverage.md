# coverage

https://pkg.go.dev/internal/coverage@go1.20.1























## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=325)

```
const CounterFilePref = "covcounters"
```

CounterFilePref is the file prefix used when emitting coverage data output files. CounterFileTemplate describes the format of the file name: prefix followed by meta-file hash followed by process ID followed by emit UnixNanoTime.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=327)

```
const CounterFileRegexp = `^%s\.(\S+)\.(\d+)\.(\d+)+$`
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=326)

```
const CounterFileTempl = "%s.%x.%d.%d"
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=277)

```
const CounterFileVersion = 1
```

CounterFileVersion stores the most recent counter data file version.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=96)

```
const CovMetaHeaderSize = 16 + 4 + 4 + 4 + 4 + 4 + 4 + 4 // keep in sync with above
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=374)

```
const FirstCtrOffset = 3
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=373)

```
const FuncIdOffset = 2
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=62)

```
const MetaFilePref = "covmeta"
```

MetaFilePref is a prefix used when emitting meta-data files; these files are of the form "covmeta.<hash>", where hash is a hash computed from the hashes of all the package meta-data symbols in the program.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=65)

```
const MetaFileVersion = 1
```

MetaFileVersion contains the current (most recent) meta-data file version.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/pkid.go;l=68)

```
const NotHardCoded = -1
```

NotHardCoded is a package pseudo-ID indicating that a given package is not part of the runtime and doesn't require a hard-coded ID.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=371)

```
const NumCtrsOffset = 0
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=372)

```
const PkgIdOffset = 1
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=274)

```
var CovCounterMagic = [4]byte{'\x00', '\x63', '\x77', '\x6d'}
```

CovCounterMagic holds the magic string for a coverage counter-data file.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/defs.go;l=56)

```
var CovMetaMagic = [4]byte{'\x00', '\x63', '\x76', '\x6d'}
```

CovMetaMagic holds the magic string for a meta-data file.

## 函数

#### func HardCodedPkgID 

```
func HardCodedPkgID(pkgpath string) int
```

HardCodedPkgID returns the hard-coded ID for the specified package path, or -1 if we don't use a hard-coded ID. Hard-coded IDs start at -2 and decrease as we go down the list.

#### func Round4 

```
func Round4(x int) int
```

## 类型

### type CounterFileFooter 

```
type CounterFileFooter struct {
	Magic [4]byte

	NumSegments uint32
	// contains filtered or unexported fields
}
```

CounterFileFooter appears at the tail end of a counter data file, and stores the number of segments it contains.

### type CounterFileHeader 

```
type CounterFileHeader struct {
	Magic     [4]byte
	Version   uint32
	MetaHash  [16]byte
	CFlavor   CounterFlavor
	BigEndian bool
	// contains filtered or unexported fields
}
```

CounterFileHeader stores files header information for a counter-data file.

### type CounterFlavor 

```
type CounterFlavor uint8
```

CounterFlavor describes how function and counters are stored/represented in the counter section of the file.

```
const (
	// "Raw" representation: all values (pkg ID, func ID, num counters,
	// and counters themselves) are stored as uint32's.
	CtrRaw CounterFlavor = iota + 1

	// "ULeb" representation: all values (pkg ID, func ID, num counters,
	// and counters themselves) are stored with ULEB128 encoding.
	CtrULeb128
)
```

### type CounterGranularity 

```
type CounterGranularity uint8
```

CounterGranularity tracks the granularity of the coverage counters being used in a given coverage-instrumented program.

```
const (
	CtrGranularityInvalid CounterGranularity = iota
	CtrGranularityPerBlock
	CtrGranularityPerFunc
)
```

#### (CounterGranularity) String 

```
func (cm CounterGranularity) String() string
```

### type CounterMode 

```
type CounterMode uint8
```

CounterMode tracks the "flavor" of the coverage counters being used in a given coverage-instrumented program.

```
const (
	CtrModeInvalid  CounterMode = iota
	CtrModeSet                  // "set" mode
	CtrModeCount                // "count" mode
	CtrModeAtomic               // "atomic" mode
	CtrModeRegOnly              // registration-only pseudo-mode
	CtrModeTestMain             // testmain pseudo-mode
)
```

#### func ParseCounterMode 

```
func ParseCounterMode(mode string) CounterMode
```

#### (CounterMode) String 

```
func (cm CounterMode) String() string
```

### type CounterSegmentHeader 

```
type CounterSegmentHeader struct {
	FcnEntries uint64
	StrTabLen  uint32
	ArgsLen    uint32
}
```

CounterSegmentHeader encapsulates information about a specific segment in a counter data file, which at the moment contains counters data from a single execution of a coverage-instrumented program. Following the segment header will be the string table and args table, and then (possibly) padding bytes to bring the byte size of the preamble up to a multiple of 4. Immediately following that will be the counter payloads.

The "args" section of a segment is used to store annotations describing where the counter data came from; this section is basically a series of key-value pairs (can be thought of as an encoded 'map[string]string'). At the moment we only write os.Args() data to this section, using pairs of the form "argc=<integer>", "argv0=<os.Args[0]>", "argv1=<os.Args[1]>", and so on. In the future the args table may also include things like GOOS/GOARCH values, and/or tags indicating which tests were run to generate the counter data.

### type CoverFixupConfig 

```
type CoverFixupConfig struct {
	// Name of the variable (created by cmd/cover) containing the
	// encoded meta-data for the package.
	MetaVar string

	// Length of the meta-data.
	MetaLen int

	// Hash computed by cmd/cover of the meta-data.
	MetaHash string

	// Instrumentation strategy. For now this is always set to
	// "normal", but in the future we may add new values (for example,
	// if panic paths are instrumented, or if the instrumenter
	// eliminates redundant counters).
	Strategy string

	// Prefix assigned to the names of counter variables generated
	// during instrumentation by cmd/cover.
	CounterPrefix string

	// Name chosen for the package ID variable generated during
	// instrumentation.
	PkgIdVar string

	// Counter mode (e.g. set/count/atomic)
	CounterMode string

	// Counter granularity (perblock or perfunc).
	CounterGranularity string
}
```

CoverFixupConfig contains annotations/notes generated by the cmd/cover tool (during instrumentation) to be passed on to the compiler when the instrumented code is compiled. The cmd/cover tool creates a struct of this type, JSON-encodes it, and emits the result to a file, which the Go command then passes to the compiler when the instrumented package is built.

### type CoverPkgConfig 

```
type CoverPkgConfig struct {
	// File into which cmd/cover should emit summary info
	// when instrumentation is complete.
	OutConfig string

	// Import path for the package being instrumented.
	PkgPath string

	// Package name.
	PkgName string

	// Instrumentation granularity: one of "perfunc" or "perblock" (default)
	Granularity string

	// Module path for this package (empty if no go.mod in use)
	ModulePath string

	// Local mode indicates we're doing a coverage build or test of a
	// package selected via local import path, e.g. "./..." or
	// "./foo/bar" as opposed to a non-relative import path. See the
	// corresponding field in cmd/go's PackageInternal struct for more
	// info.
	Local bool
}
```

CoverPkgConfig is a bundle of information passed from the Go command to the cover command during "go build -cover" runs. The Go command creates and fills in a struct as below, then passes file containing the encoded JSON for the struct to the "cover" tool when instrumenting the source files in a Go package.

### type CoverableUnit 

```
type CoverableUnit struct {
	StLine, StCol uint32
	EnLine, EnCol uint32
	NxStmts       uint32
	Parent        uint32
}
```

CoverableUnit describes the source characteristics of a single program unit for which we want to gather coverage info. Coverable units are either "simple" or "intraline"; a "simple" coverable unit corresponds to a basic block (region of straight-line code with no jumps or control transfers). An "intraline" unit corresponds to a logical clause nested within some other simple unit. A simple unit will have a zero Parent value; for an intraline unit NxStmts will be zero and and Parent will be set to 1 plus the index of the containing simple statement. Example:

```
L7:   q := 1
L8:   x := (y == 101 || launch() == false)
L9:   r := x * 2
```

For the code above we would have three simple units (one for each line), then an intraline unit describing the "launch() == false" clause in line 8, with Parent pointing to the index of the line 8 unit in the units array.

Note: in the initial version of the coverage revamp, only simple units will be in use.

### type FuncDesc 

```
type FuncDesc struct {
	Funcname string
	Srcfile  string
	Units    []CoverableUnit
	Lit      bool // true if this is a function literal
}
```

FuncDesc encapsulates the meta-data definitions for a single Go function. This version assumes that we're looking at a function before inlining; if we want to capture a post-inlining view of the world, the representations of source positions would need to be a good deal more complicated.

### type MetaFileHeader 

```
type MetaFileHeader struct {
	Magic        [4]byte
	Version      uint32
	TotalLength  uint64
	Entries      uint64
	MetaFileHash [16]byte
	StrTabOffset uint32
	StrTabLength uint32
	CMode        CounterMode
	CGranularity CounterGranularity
	// contains filtered or unexported fields
}
```

MetaFileHeader stores file header information for a meta-data file.

### type MetaSymbolHeader 

```
type MetaSymbolHeader struct {
	Length     uint32 // size of meta-symbol payload in bytes
	PkgName    uint32 // string table index
	PkgPath    uint32 // string table index
	ModulePath uint32 // string table index
	MetaHash   [16]byte

	NumFiles uint32
	NumFuncs uint32
	// contains filtered or unexported fields
}
```

MetaSymbolHeader stores header information for a single meta-data blob, e.g. the coverage meta-data payload computed for a given Go package.