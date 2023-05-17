# cmerge

https://pkg.go.dev/internal/coverage/cmerge@go1.20.1






  
  
  
  
  

## 常量 [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#pkg-constants)

This section is empty.

## 变量

This section is empty.

## 函数

#### func [SaturatingAdd](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/cmerge/merge.go;l=61) [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#SaturatingAdd)

```
func SaturatingAdd(dst, src uint32) (uint32, bool)
```

Saturating add does a saturing addition of 'dst' and 'src', returning added value or math.MaxUint32 plus an overflow flag.

## 类型

### type [Merger](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/cmerge/merge.go;l=20) [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#Merger)

```
type Merger struct {
	// contains filtered or unexported fields
}
```

Merger provides state and methods to help manage the process of merging together coverage counter data for a given function, for tools that need to implicitly merge counter as they read multiple coverage counter data files.

#### (*Merger) [Granularity](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/cmerge/merge.go;l=102) [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#Merger.Granularity)

```
func (cm *Merger) Granularity() coverage.CounterGranularity
```

#### (*Merger) [MergeCounters](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/cmerge/merge.go;l=28) [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#Merger.MergeCounters)

```
func (m *Merger) MergeCounters(dst, src []uint32) (error, bool)
```

MergeCounters takes the counter values in 'src' and merges them into 'dst' according to the correct counter mode.

#### (*Merger) [Mode](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/cmerge/merge.go;l=98) [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#Merger.Mode)

```
func (cm *Merger) Mode() coverage.CounterMode
```

#### (*Merger) [ResetModeAndGranularity](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/cmerge/merge.go;l=92) [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#Merger.ResetModeAndGranularity)

```
func (cm *Merger) ResetModeAndGranularity()
```

#### (*Merger) [SaturatingAdd](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/cmerge/merge.go;l=51) [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#Merger.SaturatingAdd)

```
func (m *Merger) SaturatingAdd(dst, src uint32) uint32
```

Saturating add does a saturating addition of 'dst' and 'src', returning added value or math.MaxUint32 if there is an overflow. Overflows are recorded in case the client needs to track them.

#### (*Merger) [SetModeAndGranularity](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/cmerge/merge.go;l=77) [¶](https://pkg.go.dev/internal/coverage/cmerge@go1.20.1#Merger.SetModeAndGranularity)

```
func (cm *Merger) SetModeAndGranularity(mdf string, cmode coverage.CounterMode, cgran coverage.CounterGranularity) error
```

SetModeAndGranularity records the counter mode and granularity for the current merge. In the specific case of merging across coverage data files from different binaries, where we're combining data from more than one meta-data file, we need to check for mode/granularity clashes.