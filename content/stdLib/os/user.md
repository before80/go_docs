+++
title = "user"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/os/user@go1.21.3

Package user allows user account lookups by name or id.

​	`user`包允许按名称或 ID 查找用户帐户。

For most Unix systems, this package has two internal implementations of resolving user and group ids to names, and listing supplementary group IDs. One is written in pure Go and parses /etc/passwd and /etc/group. The other is cgo-based and relies on the standard C library (libc) routines such as getpwuid_r, getgrnam_r, and getgrouplist.

​	对于大多数 Unix 系统，此包具有两种内部实现解析用户和组 ID 为名称，并列出附加组 ID。一种是纯 Go 编写并解析 `/etc/passwd` 和 `/etc/group`。另一种是基于 cgo 编写的，并依赖于标准 C 库(libc)例程，例如 getpwuid_r、getgrnam_r 和 getgrouplist。

When cgo is available, and the required routines are implemented in libc for a particular platform, cgo-based (libc-backed) code is used. This can be overridden by using osusergo build tag, which enforces the pure Go implementation.

​	当 cgo 可用，并且特定平台的 libc 中实现了所需例程时，将使用基于 cgo(libc 支持)的代码。这可以通过使用 osusergo 构建标签进行覆盖，该标签强制使用纯 Go 实现。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Group  <- go1.7

``` go 
type Group struct {
	Gid  string // group ID
	Name string // group name
}
```

Group represents a grouping of users.

​	Group 表示用户组。

On POSIX systems Gid contains a decimal number representing the group ID.

​	在 POSIX 系统上，Gid 包含表示组 ID 的十进制数。

#### func LookupGroup  <- go1.7

``` go 
func LookupGroup(name string) (*Group, error)
```

LookupGroup looks up a group by name. If the group cannot be found, the returned error is of type UnknownGroupError.

​	`LookupGroup`函数按名称查找组。如果找不到组，则返回的错误是 UnknownGroupError 类型。

#### func LookupGroupId  <- go1.7

``` go 
func LookupGroupId(gid string) (*Group, error)
```

LookupGroupId looks up a group by groupid. If the group cannot be found, the returned error is of type UnknownGroupIdError.

​	`LookupGroupId`函数按组 ID 查找组。如果找不到组，则返回的错误是 UnknownGroupIdError 类型。

### type UnknownGroupError  <- go1.7

``` go 
type UnknownGroupError string
```

UnknownGroupError is returned by LookupGroup when a group cannot be found.

​	`LookupGroup`函数无法找到组时返回 UnknownGroupError。

#### (UnknownGroupError) Error  <- go1.7

``` go 
func (e UnknownGroupError) Error() string
```

### type UnknownGroupIdError  <- go1.7

``` go 
type UnknownGroupIdError string
```

UnknownGroupIdError is returned by LookupGroupId when a group cannot be found.

​	`LookupGroupId`函数无法找到组时返回 UnknownGroupIdError。

#### (UnknownGroupIdError) Error  <- go1.7

``` go 
func (e UnknownGroupIdError) Error() string
```

### type UnknownUserError 

``` go 
type UnknownUserError string
```

UnknownUserError is returned by Lookup when a user cannot be found.

​	`Lookup`函数无法找到用户时返回 UnknownUserError。

#### (UnknownUserError) Error 

``` go 
func (e UnknownUserError) Error() string
```

### type UnknownUserIdError 

``` go 
type UnknownUserIdError int
```

UnknownUserIdError is returned by LookupId when a user cannot be found.

​	`LookupId`函数无法找到用户时返回 UnknownUserIdError。

#### (UnknownUserIdError) Error 

``` go 
func (e UnknownUserIdError) Error() string
```

### type User 

``` go 
type User struct {
    // Uid is the user ID.
	// On POSIX systems, this is a decimal number representing the uid.
	// On Windows, this is a security identifier (SID) in a string format.
	// On Plan 9, this is the contents of /dev/user.
	// Uid 是用户 ID。
	// 在 POSIX 系统上，这是表示 uid 的十进制数。
	// 在 Windows 上，这是以字符串格式表示的安全标识符 (SID)。
	// 在 Plan 9 上，这是 /dev/user 的内容。
	Uid string
    // Gid is the primary group ID.
	// On POSIX systems, this is a decimal number representing the gid.
	// On Windows, this is a SID in a string format.
	// On Plan 9, this is the contents of /dev/user.
	// Gid 是主要组 ID。
	// 在 POSIX 系统上，这是表示 gid 的十进制数。
	// 在 Windows 上，这是以字符串格式表示的 SID。
	// 在 Plan 9 上，这是 /dev/user 的内容。
	Gid string
    // Username is the login name.
	// Username 是登录名。
	Username string
    // Name is the user's real or display name.
	// It might be blank.
	// On POSIX systems, this is the first (or only) entry in the GECOS field
	// list.
	// On Windows, this is the user's display name.
	// On Plan 9, this is the contents of /dev/user.
	// Name 是用户的真实姓名或显示名称。
	// 可能为空。
	// 在 POSIX 系统上，这是 GECOS 字段列表中的第一个(或唯一)条目。
	// 在 Windows 上，这是用户的显示名称。
	// 在 Plan 9 上，这是 /dev/user 的内容。
	Name string
    
    // HomeDir is the path to the user's home directory (if they have one).
	// HomeDir 是用户主目录的路径(如果有)。
    HomeDir string
}
```

User represents a user account.

​	User 表示用户帐户。

#### func Current 

``` go 
func Current() (*User, error)
```

Current returns the current user.

​	`Current`函数返回当前用户。

The first call will cache the current user information. Subsequent calls will return the cached value and will not reflect changes to the current user.

​	第一次调用将缓存当前用户信息。后续调用将返回缓存值，不会反映当前用户的更改。

#### func Lookup 

``` go 
func Lookup(username string) (*User, error)
```

Lookup looks up a user by username. If the user cannot be found, the returned error is of type UnknownUserError.

​	`Lookup`函数按用户名查找用户。如果找不到用户，则返回的错误是 UnknownUserError 类型。

#### func LookupId 

``` go 
func LookupId(uid string) (*User, error)
```

LookupId looks up a user by userid. If the user cannot be found, the returned error is of type UnknownUserIdError.

​	`LookupId`函数按用户 ID 查找用户。如果找不到用户，则返回的错误是 UnknownUserIdError 类型。

#### (*User) GroupIds  <- go1.7

``` go 
func (u *User) GroupIds() ([]string, error)
```

GroupIds returns the list of group IDs that the user is a member of.

​	`GroupIds`方法返回用户所属的组 ID 列表。