# calloc

> 原文：[https://pkg.go.dev/internal/coverage/calloc@go1.23.0](https://pkg.go.dev/internal/coverage/calloc@go1.23.0)






## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type BatchCounterAlloc 

``` go
type BatchCounterAlloc struct {
	// contains filtered or unexported fields
}
```

#### (*BatchCounterAlloc) AllocateCounters 

``` go
func (ca *BatchCounterAlloc) AllocateCounters(n int) []uint32
```