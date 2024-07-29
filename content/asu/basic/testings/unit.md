+++
title = "单元测试"
date = 2024-07-13T14:21:54+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

有什么好的被测试函数？



## 测试代码的组织

### 平铺模式



### xUnit家族模式



## 测试函数

### 

### testing.AllocsPerRun

### testing.CoverMode

### testing.Coverage

### testing.Init

### testing.Main

### testing.Register

### testing.RunBenchmarks

### testing.RunExamples

### testing.RunTests

### testing.Short

### testing.Verbose



### testing.T.Log

### testing.T.Logf

### testing.T.Error

### testing.T.Errorf

### testing.T.Fatal

### testing.T.Fatalf

### testing.T.Fail

### testing.T.FailNow

### testing.T.Failed

### testing.T.Parallel

### testing.T.Run

### testing.T.Skip

### testing.T.SkipNow

### testing.T.Skipf

### testing.T.Skipped

### testing.T.TempDir

### testing.T.Name

### testing.T.Cleanup

### testing.T.Deadline

### testing.T.Setenv

### testing.T.Helper



### testing.TB.Log

### testing.TB.Logf

### testing.TB.Error

### testing.TB.Errorf

### testing.TB.Fatal

### testing.TB.Fatalf

### testing.TB.Fail

### testing.TB.FailNow

### testing.TB.Failed

### testing.TB.Skip

### testing.TB.SkipNow

### testing.TB.Skipf

### testing.TB.Skipped

### testing.TB.TempDir

### testing.TB.Name

### testing.TB.Cleanup

### testing.TB.Setenv

### testing.TB.Helper

## 测试命令

go test 

go test .

go test ./...

go test ./onePkg

go test -v

go test -short

go test -fastfail

go test -p 1 ./...

go test -run regexpStr ./...

go test -timeout 50ms ./...

> 50ms是针对？

go test -count=1

> 这里的=符号是否可以省略？

