+++
title = "Casbin Auth"
weight = 40
date = 2023-07-09T21:54:16+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Casbin Auth

> 原文：[https://echo.labstack.com/docs/middleware/casbin-auth](https://echo.labstack.com/docs/middleware/casbin-auth)

> 注意
>
> ​	这是 Echo 社区的贡献。

​	[Casbin](https://github.com/casbin/casbin) 是一个功能强大且高效的 Go 开源访问控制库。它支持基于各种模型的授权策略。目前，Casbin 支持的访问控制模型有： 

- ACL（访问控制列表 Access Control List）
- 带超级用户的 ACL （ACL with superuser）
- 无用户的 ACL：适用于没有身份验证或用户登录的系统。
- 无资源的 ACL：某些场景可能针对的是一类资源，而不是个别资源，可以使用像写文章、读取日志之类的权限。它不控制对特定文章或日志的访问。
- RBAC（基于角色的访问控制 Role-Based Access Control）
- 带资源角色的 RBAC：用户和资源可以同时拥有角色（或组）。
- 带域/租户（domains/tenants）的 RBAC：用户可以为不同的域/租户设置不同的角色集合。
- ABAC（基于属性的访问控制 Attribute-Based Access Control）
- RESTful
- Deny-override：支持允许（allow）和拒绝（deny）两种授权方式，拒绝会覆盖允许。

> 信息
>
> ​	目前仅支持 HTTP 基本身份验证。

## 依赖

```go
import (
  "github.com/casbin/casbin"
  casbin_mw "github.com/labstack/echo-contrib/casbin"
)
```



## Usage

```go
e := echo.New()
enforcer, err := casbin.NewEnforcer("casbin_auth_model.conf", "casbin_auth_policy.csv")
e.Use(casbin_mw.Middleware(enforcer))
```



​	有关语法，请参阅：[模型语法](https://casbin.org/docs/en/syntax-for-models)。

## Custom Configuration

### Usage

```go
e := echo.New()
ce := casbin.NewEnforcer("casbin_auth_model.conf", "")
ce.AddRoleForUser("alice", "admin")
ce.AddPolicy(...)
e.Use(casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
  Enforcer: ce,
}))
```



## Configuration

```go
// Config 定义了 CasbinAuth 中间件的配置。
Config struct {
  // Skipper 定义一个用于跳过中间件的函数。
  Skipper middleware.Skipper

  // Enforcer 是 CasbinAuth 的主要规则。
  // 必填项。
  Enforcer *casbin.Enforcer
}
```



### Default Configuration

```go
// DefaultConfig 是 CasbinAuth 中间件的默认配置。
DefaultConfig = Config{
  Skipper: middleware.DefaultSkipper,
}
```



