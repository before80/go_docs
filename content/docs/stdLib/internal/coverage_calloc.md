# calloc

https://pkg.go.dev/internal/coverage/calloc@go1.20.1






## 常量 [¶](https://pkg.go.dev/internal/coverage/calloc@go1.20.1#pkg-constants)

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type [BatchCounterAlloc](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/calloc/batchcounteralloc.go;l=13) [¶](https://pkg.go.dev/internal/coverage/calloc@go1.20.1#BatchCounterAlloc)

```
type BatchCounterAlloc struct {
	// contains filtered or unexported fields
}
```

#### (*BatchCounterAlloc) [AllocateCounters](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/calloc/batchcounteralloc.go;l=17) [¶](https://pkg.go.dev/internal/coverage/calloc@go1.20.1#BatchCounterAlloc.AllocateCounters)

```
func (ca *BatchCounterAlloc) AllocateCounters(n int) []uint32
```