# profile

> 原文：[https://pkg.go.dev/internal/profile@go1.20.1](https://pkg.go.dev/internal/profile@go1.20.1)



Package profile provides a representation of github.com/google/pprof/proto/profile.proto and methods to encode/decode/merge profiles in this format.












  
  

  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  




## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/profile/legacy_profile.go;l=21)

``` go
var (

	// LegacyHeapAllocated instructs the heapz parsers to use the
	// allocated memory stats instead of the default in-use memory. Note
	// that tcmalloc doesn't provide all allocated memory, only in-use
	// stats.
	LegacyHeapAllocated bool
)
```

## 函数

This section is empty.

## 类型

### type Demangler 

``` go
type Demangler func(name []string) (map[string]string, error)
```

Demangler maps symbol names to a human-readable form. This may include C++ demangling and additional simplification. Names that are not demangled may be missing from the resulting map.

### type Function 

``` go
type Function struct {
	ID         uint64
	Name       string
	SystemName string
	Filename   string
	StartLine  int64
	// contains filtered or unexported fields
}
```

Function corresponds to Profile.Function

### type Label 

``` go
type Label struct {
	// contains filtered or unexported fields
}
```

Label corresponds to Profile.Label

### type Line 

``` go
type Line struct {
	Function *Function
	Line     int64
	// contains filtered or unexported fields
}
```

Line corresponds to Profile.Line

### type Location 

``` go
type Location struct {
	ID       uint64
	Mapping  *Mapping
	Address  uint64
	Line     []Line
	IsFolded bool
	// contains filtered or unexported fields
}
```

Location corresponds to Profile.Location

### type Mapping 

``` go
type Mapping struct {
	ID              uint64
	Start           uint64
	Limit           uint64
	Offset          uint64
	File            string
	BuildID         string
	HasFunctions    bool
	HasFilenames    bool
	HasLineNumbers  bool
	HasInlineFrames bool
	// contains filtered or unexported fields
}
```

Mapping corresponds to Profile.Mapping

### type Profile 

``` go
type Profile struct {
	SampleType        []*ValueType
	DefaultSampleType string
	Sample            []*Sample
	Mapping           []*Mapping
	Location          []*Location
	Function          []*Function
	Comments          []string

	DropFrames string
	KeepFrames string

	TimeNanos     int64
	DurationNanos int64
	PeriodType    *ValueType
	Period        int64
	// contains filtered or unexported fields
}
```

Profile is an in-memory representation of profile.proto.

#### func Merge 

``` go
func Merge(srcs []*Profile) (*Profile, error)
```

Merge merges all the profiles in profs into a single Profile. Returns a new profile independent of the input profiles. The merged profile is compacted to eliminate unused samples, locations, functions and mappings. Profiles must have identical profile sample and period types or the merge will fail. profile.Period of the resulting profile will be the maximum of all profiles, and profile.TimeNanos will be the earliest nonzero one.

#### func Parse 

``` go
func Parse(r io.Reader) (*Profile, error)
```

Parse parses a profile and checks for its validity. The input may be a gzip-compressed encoded protobuf or one of many legacy profile formats which may be unsupported in the future.

#### func ParseTracebacks 

``` go
func ParseTracebacks(b []byte) (*Profile, error)
```

ParseTracebacks parses a set of tracebacks and returns a newly populated profile. It will accept any text file and generate a Profile out of it with any hex addresses it can identify, including a process map if it can recognize one. Each sample will include a tag "source" with the addresses recognized in string format.

#### (*Profile) Aggregate 

``` go
func (p *Profile) Aggregate(inlineFrame, function, filename, linenumber, address bool) error
```

Aggregate merges the locations in the profile into equivalence classes preserving the request attributes. It also updates the samples to point to the merged locations.

#### (*Profile) CheckValid 

``` go
func (p *Profile) CheckValid() error
```

CheckValid tests whether the profile is valid. Checks include, but are not limited to:

- len(Profile.Sample[n].value) == len(Profile.value_unit)
- Sample.id has a corresponding Profile.Location

#### (*Profile) Compatible 

``` go
func (p *Profile) Compatible(pb *Profile) error
```

Compatible determines if two profiles can be compared/merged. returns nil if the profiles are compatible; otherwise an error with details on the incompatibility.

#### (*Profile) Copy 

``` go
func (p *Profile) Copy() *Profile
```

Copy makes a fully independent copy of a profile.

#### (*Profile) Demangle 

``` go
func (p *Profile) Demangle(d Demangler) error
```

Demangle attempts to demangle and optionally simplify any function names referenced in the profile. It works on a best-effort basis: it will silently preserve the original names in case of any errors.

#### (*Profile) Empty 

``` go
func (p *Profile) Empty() bool
```

Empty reports whether the profile contains no samples.

#### (*Profile) FilterSamplesByName 

``` go
func (p *Profile) FilterSamplesByName(focus, ignore, hide *regexp.Regexp) (fm, im, hm bool)
```

FilterSamplesByName filters the samples in a profile and only keeps samples where at least one frame matches focus but none match ignore. Returns true is the corresponding regexp matched at least one sample.

#### (*Profile) FilterSamplesByTag 

``` go
func (p *Profile) FilterSamplesByTag(focus, ignore TagMatch) (fm, im bool)
```

FilterSamplesByTag removes all samples from the profile, except those that match focus and do not match the ignore regular expression.

#### (*Profile) HasFileLines 

``` go
func (p *Profile) HasFileLines() bool
```

HasFileLines determines if all locations in this profile have symbolized file and line number information.

#### (*Profile) HasFunctions 

``` go
func (p *Profile) HasFunctions() bool
```

HasFunctions determines if all locations in this profile have symbolized function information.

#### (*Profile) Merge 

``` go
func (p *Profile) Merge(pb *Profile, r float64) error
```

Merge adds profile p adjusted by ratio r into profile p. Profiles must be compatible (same Type and SampleType). TODO(rsilvera): consider normalizing the profiles based on the total samples collected.

#### (*Profile) Normalize 

``` go
func (p *Profile) Normalize(pb *Profile) error
```

Normalize normalizes the source profile by multiplying each value in profile by the ratio of the sum of the base profile's values of that sample type to the sum of the source profile's value of that sample type.

#### (*Profile) ParseMemoryMap 

``` go
func (p *Profile) ParseMemoryMap(rd io.Reader) error
```

ParseMemoryMap parses a memory map in the format of /proc/self/maps, and overrides the mappings in the current profile. It renumbers the samples and locations in the profile correspondingly.

#### (*Profile) Prune 

``` go
func (p *Profile) Prune(dropRx, keepRx *regexp.Regexp)
```

Prune removes all nodes beneath a node matching dropRx, and not matching keepRx. If the root node of a Sample matches, the sample will have an empty stack.

#### (*Profile) RemoveUninteresting 

``` go
func (p *Profile) RemoveUninteresting() error
```

RemoveUninteresting prunes and elides profiles using built-in tables of uninteresting function names.

#### (*Profile) Scale 

``` go
func (p *Profile) Scale(ratio float64)
```

Scale multiplies all sample values in a profile by a constant.

#### (*Profile) ScaleN 

``` go
func (p *Profile) ScaleN(ratios []float64) error
```

ScaleN multiplies each sample values in a sample by a different amount.

#### (*Profile) String 

``` go
func (p *Profile) String() string
```

Print dumps a text representation of a profile. Intended mainly for debugging purposes.

#### (*Profile) Write 

``` go
func (p *Profile) Write(w io.Writer) error
```

Write writes the profile as a gzip-compressed marshaled protobuf.

### type Sample 

``` go
type Sample struct {
	Location []*Location
	Value    []int64
	Label    map[string][]string
	NumLabel map[string][]int64
	NumUnit  map[string][]string
	// contains filtered or unexported fields
}
```

Sample corresponds to Profile.Sample

### type TagMatch 

``` go
type TagMatch func(key, val string, nval int64) bool
```

TagMatch selects tags for filtering

### type ValueType 

``` go
type ValueType struct {
	Type string // cpu, wall, inuse_space, etc
	Unit string // seconds, nanoseconds, bytes, etc
	// contains filtered or unexported fields
}
```

ValueType corresponds to Profile.ValueType