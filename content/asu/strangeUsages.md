+++
title = "那些看起来有些奇怪的用法"
date = 2023-08-18T14:13:48+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# 那些看起来有些奇怪的用法



## 代码中的指令



### //go:build



### //go:cgo_dynamic_linker



### //go:cgo_export_dynamic



### //go:cgo_export_static



### //go:cgo_import_dynamic



### //go:cgo_import_static



### //go:cgo_ldflag



### //go:cgo_unsafe_args





### //go:embed



### //go:generate



### //go:linkname



### //go:name



### //go:noescape





### //go:noinline



### //go:nosplit



### //go:registerparams



### //go:systemstack



### //go:uintptrescapes



### //go:xxx



### //go:yeswritebarrierrec



## 2006-01-02T15:04:05 有什么用？





## fstest 中的 var _ fs.FS = MapFS(nil) 有什么用？

## fstest 中的 var _ fs.File = (*openMapFile)(nil) 有什么用？



## n := 2; fmt.Println(string(n)) 输出却是空的！或者不能编译通过！为什么？

​	因为：`string()`转化：conversion from int to string yields a string of one rune, not a string of digits （将 int类型的值通过string()进行转化，会产生某一个rune字符对应的字符串，而不是产生对应数值的字符串），而这里的 n 是 int 类型！

```go
package main

func main() {
    n := 2
    str := string(n)
    fmt.Println(str)
    //fmt.Printf("str=%q,len=%d\n",str,len(str))  // str="\x02",len=1
}
```

​	又因`\x02` ，十进制编码值为 2，对应ASCII 码中的 "Start of Text" 控制字符，显示为 `^B` 

​	`\x02`在不同的环境下显示效果可能不同：

- 在终端中，它通常不显示任何可见字符 
- 在文本编辑器中，可能显示为 `^B` 
- 在浏览器中，可能显示为一个小方块 `□` 

