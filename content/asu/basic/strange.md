+++
title = "那些奇怪的用法"
date = 2024-07-13T11:14:41+08:00
weight = 1500
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++



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



## func New() gdb.Driver {   return &Driver{} } 为什么需要使用 &?

```go
func New() gdb.Driver {
    return &Driver{}
}
```

gdb.Driver的定义

```go
type Driver interface {
	// New creates and returns a database object for specified database server.
	New(core *Core, node *ConfigNode) (DB, error)
}
```

Driver的定义

```go
type Driver struct {
	*gdb.Core
}
```

gdb.Core的定义

```go
type Core struct {
	db            DB              // DB interface object.
	ctx           context.Context // Context for chaining operation only. Do not set a default value in Core initialization.
	group         string          // Configuration group name.
	schema        string          // Custom schema for this object.
	debug         *gtype.Bool     // Enable debug mode for the database, which can be changed in runtime.
	cache         *gcache.Cache   // Cache manager, SQL result cache only.
	links         *gmap.StrAnyMap // links caches all created links by node.
	logger        glog.ILogger    // Logger for logging functionality.
	config        *ConfigNode     // Current config node.
	dynamicConfig dynamicConfig   // Dynamic configurations, which can be changed in runtime.
}
```



​	我们先将以上相关代码做下简化:

```go
package main

import "fmt"

// gdb.Driver -> A
type A interface {
	F()
}

// Driver -> a
type a struct {
	b
}

func (ar a) F() {}

// gdb.Core -> b
type b struct {
	V string
}

func New1() A {
	return &a{}
}

func New2() A {
	return a{}
}

func main() {
	v1 := New1()
	v2 := New2()

	fmt.Printf("%T, %#v\n", v1, v1)
	fmt.Printf("%T, %#v\n", v2, v2)
}

*main.a, &main.a{b:main.b{V:""}}
main.a, main.a{b:main.b{V:""}}
```

​	我们发现实际上无论是否使用`&`这个符号，返回的值都是满足 `A`这个接口的。使用`&`只是减少数据的拷贝， 如果 `a` 是一个很大的结构体(包含很多字段或嵌入字段)，不使用`&`，势必会发生拷贝时间长等问题。
