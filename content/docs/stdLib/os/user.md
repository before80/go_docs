+++
title = "user"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# user

[https://pkg.go.dev/os/user@go1.20.1](https://pkg.go.dev/os/user@go1.20.1)

​	user包允许按名称或 ID 查找用户帐户。

​	对于大多数 Unix 系统，此包具有两种内部实现解析用户和组 ID 为名称，并列出附加组 ID。一种是纯 Go 编写并解析 `/etc/passwd` 和 `/etc/group`。另一种是基于 cgo 编写的，并依赖于标准 C 库(libc)例程，例如 getpwuid_r、getgrnam_r 和 getgrouplist。

​	当 cgo 可用，并且特定平台的 libc 中实现了所需例程时，将使用基于 cgo(libc 支持)的代码。这可以通过使用 osusergo 构建标签进行覆盖，该标签强制使用纯 Go 实现。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type [Group](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=61)  <- go1.7

``` go linenums="1"
type Group struct {
	Gid  string // group ID
	Name string // group name
}
```

​	Group 表示用户组。

​	在 POSIX 系统上，Gid 包含表示组 ID 的十进制数。

#### func [LookupGroup](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/lookup.go;l=57)  <- go1.7

``` go linenums="1"
func LookupGroup(name string) (*Group, error)
```

​	LookupGroup函数按名称查找组。如果找不到组，则返回的错误是 UnknownGroupError 类型。

#### func [LookupGroupId](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/lookup.go;l=63)  <- go1.7

``` go linenums="1"
func LookupGroupId(gid string) (*Group, error)
```

​	LookupGroupId函数按组 ID 查找组。如果找不到组，则返回的错误是 UnknownGroupIdError 类型。

### type [UnknownGroupError](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=91)  <- go1.7

``` go linenums="1"
type UnknownGroupError string
```

​	LookupGroup函数无法找到组时返回 UnknownGroupError。

#### (UnknownGroupError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=93)  <- go1.7

``` go linenums="1"
func (e UnknownGroupError) Error() string
```

### type [UnknownGroupIdError](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=83)  <- go1.7

``` go linenums="1"
type UnknownGroupIdError string
```

​	LookupGroupId函数无法找到组时返回 UnknownGroupIdError。

#### (UnknownGroupIdError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=85)  <- go1.7

``` go linenums="1"
func (e UnknownGroupIdError) Error() string
```

### type [UnknownUserError](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=75) 

``` go linenums="1"
type UnknownUserError string
```

​	Lookup函数无法找到用户时返回 UnknownUserError。

#### (UnknownUserError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=77) 

``` go linenums="1"
func (e UnknownUserError) Error() string
```

### type [UnknownUserIdError](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=67) 

``` go linenums="1"
type UnknownUserIdError int
```

​	LookupId函数无法找到用户时返回 UnknownUserIdError。

#### (UnknownUserIdError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=69) 

``` go linenums="1"
func (e UnknownUserIdError) Error() string
```

### type [User](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/user.go;l=34) 

``` go linenums="1"
type User struct {
	// Uid 是用户 ID。
	// 在 POSIX 系统上，这是表示 uid 的十进制数。
	// 在 Windows 上，这是以字符串格式表示的安全标识符 (SID)。
	// 在 Plan 9 上，这是 /dev/user 的内容。
	Uid string
	// Gid 是主要组 ID。
	// 在 POSIX 系统上，这是表示 gid 的十进制数。
	// 在 Windows 上，这是以字符串格式表示的 SID。
	// 在 Plan 9 上，这是 /dev/user 的内容。
	Gid string
	// Username 是登录名。
	Username string
	// Name 是用户的真实姓名或显示名称。
	// 可能为空。
	// 在 POSIX 系统上，这是 GECOS 字段列表中的第一个(或唯一)条目。
	// 在 Windows 上，这是用户的显示名称。
	// 在 Plan 9 上，这是 /dev/user 的内容。
	Name string
    
	// HomeDir 是用户主目录的路径(如果有)。
    HomeDir string
}
```

​	User 表示用户帐户。

#### func [Current](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/lookup.go;l=21) 

``` go linenums="1"
func Current() (*User, error)
```

​	Current函数返回当前用户。

​	第一次调用将缓存当前用户信息。后续调用将返回缓存值，不会反映当前用户的更改。

#### func [Lookup](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/lookup.go;l=39) 

``` go linenums="1"
func Lookup(username string) (*User, error)
```

​	Lookup函数按用户名查找用户。如果找不到用户，则返回的错误是 UnknownUserError 类型。

#### func [LookupId](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/lookup.go;l=48) 

``` go linenums="1"
func LookupId(uid string) (*User, error)
```

​	LookupId函数按用户 ID 查找用户。如果找不到用户，则返回的错误是 UnknownUserIdError 类型。

#### (*User) [GroupIds](https://cs.opensource.google/go/go/+/go1.20.1:src/os/user/lookup.go;l=68)  <- go1.7

``` go linenums="1"
func (u *User) GroupIds() ([]string, error)
```

​	GroupIds方法返回用户所属的组 ID 列表。