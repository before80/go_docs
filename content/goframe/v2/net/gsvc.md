+++
title = "gsvc"
date = 2024-03-21T17:53:32+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gsvc](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gsvc)

Package gsvc provides service registry and discovery definition.

​	软件包 gsvc 提供服务注册表和发现定义。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/net/gsvc/gsvc.go#L116)

```go
const (
	Schema                    = `service`            // Schema is the schema of service.
	DefaultHead               = `service`            // DefaultHead is the default head of service.
	DefaultDeployment         = `default`            // DefaultDeployment is the default deployment of service.
	DefaultNamespace          = `default`            // DefaultNamespace is the default namespace of service.
	DefaultVersion            = `latest`             // DefaultVersion is the default version of service.
	EnvPrefix                 = `GF_GSVC_PREFIX`     // EnvPrefix is the environment variable prefix.
	EnvDeployment             = `GF_GSVC_DEPLOYMENT` // EnvDeployment is the environment variable deployment.
	EnvNamespace              = `GF_GSVC_NAMESPACE`  // EnvNamespace is the environment variable namespace.
	EnvName                   = `GF_GSVC_Name`       // EnvName is the environment variable name.
	EnvVersion                = `GF_GSVC_VERSION`    // EnvVersion is the environment variable version.
	MDProtocol                = `protocol`           // MDProtocol is the metadata key for protocol.
	MDInsecure                = `insecure`           // MDInsecure is the metadata key for insecure.
	MDWeight                  = `weight`             // MDWeight is the metadata key for weight.
	DefaultProtocol           = `http`               // DefaultProtocol is the default protocol of service.
	DefaultSeparator          = "/"                  // DefaultSeparator is the default separator of service.
	EndpointHostPortDelimiter = ":"                  // EndpointHostPortDelimiter is the delimiter of host and port.

	EndpointsDelimiter = "," // EndpointsDelimiter is the delimiter of endpoints.
)
```

## 变量

This section is empty.

## 函数

#### func Deregister

```go
func Deregister(ctx context.Context, service Service) error
```

Deregister removes `service` from default registry.

​	取消注册将从默认注册表中删除 `service` 。

#### func SetRegistry

```go
func SetRegistry(registry Registry)
```

SetRegistry sets the default Registry implements as your own implemented interface.

​	SetRegistry 将默认的 Registry 实现设置为您自己的实现接口。

## 类型

### type Discovery

```go
type Discovery interface {
	// Search searches and returns services with specified condition.
	Search(ctx context.Context, in SearchInput) (result []Service, err error)

	// Watch watches specified condition changes.
	// The `key` is the prefix of service key.
	Watch(ctx context.Context, key string) (watcher Watcher, err error)
}
```

Discovery interface for service discovery.

​	用于服务发现的发现接口。

### type Endpoint <-2.1.0

```go
type Endpoint interface {
	// Host returns the IPv4/IPv6 address of a service.
	Host() string

	// Port returns the port of a service.
	Port() int

	// String formats and returns the Endpoint as a string.
	String() string
}
```

Endpoint interface for service.

​	服务的端点接口。

#### func NewEndpoint <-2.1.0

```go
func NewEndpoint(address string) Endpoint
```

NewEndpoint creates and returns an Endpoint from address string of pattern “host:port”, eg: “192.168.1.100:80”.

​	NewEndpoint 从模式为“host：port”的地址字符串创建并返回一个 Endpoint，例如：“192.168.1.100：80”。

### type Endpoints <-2.1.0

```go
type Endpoints []Endpoint
```

Endpoints are composed by multiple Endpoint.

​	终结点由多个终结点组成。

#### func NewEndpoints <-2.1.0

```go
func NewEndpoints(addresses string) Endpoints
```

NewEndpoints creates and returns Endpoints from multiple addresses like: “192.168.1.100:80,192.168.1.101:80”.

​	NewEndpoints 创建并返回来自多个地址的终结点，例如：“192.168.1.100：80,192.168.1.101：80”。

#### (Endpoints) String

```go
func (es Endpoints) String() string
```

String formats and returns the Endpoints as a string like: “192.168.1.100:80,192.168.1.101:80”

​	字符串格式化并将终结点返回为字符串，如下所示：“192.168.1.100：80,192.168.1.101：80”

### type LocalEndpoint <-2.1.0

```go
type LocalEndpoint struct {
	// contains filtered or unexported fields
}
```

LocalEndpoint implements interface Endpoint.

​	LocalEndpoint 实现接口 Endpoint。

#### (*LocalEndpoint) Host

```go
func (e *LocalEndpoint) Host() string
```

Host returns the IPv4/IPv6 address of a service.

​	Host 返回服务的 IPv4/IPv6 地址。

#### (*LocalEndpoint) Port

```go
func (e *LocalEndpoint) Port() int
```

Port returns the port of a service.

​	Port 返回服务的端口。

#### (*LocalEndpoint) String

```go
func (e *LocalEndpoint) String() string
```

String formats and returns the Endpoint as a string, like: 192.168.1.100:80.

​	String 格式化并将 Endpoint 作为字符串返回，如：192.168.1.100：80。

### type LocalService <-2.1.0

```go
type LocalService struct {
	Head       string    // Service custom head string in service key.
	Deployment string    // Service deployment name, eg: dev, qa, staging, prod, etc.
	Namespace  string    // Service Namespace, to indicate different services in the same environment with the same Name.
	Name       string    // Name for the service.
	Version    string    // Service version, eg: v1.0.0, v2.1.1, etc.
	Endpoints  Endpoints // Service Endpoints, pattern: IP:port, eg: 192.168.1.2:8000.
	Metadata   Metadata  // Custom data for this service, which can be set using JSON by environment or command-line.
}
```

LocalService provides a default implements for interface Service.

​	LocalService 为接口服务提供默认实现。

#### (*LocalService) GetEndpoints

```go
func (s *LocalService) GetEndpoints() Endpoints
```

GetEndpoints returns the Endpoints of service. The Endpoints contain multiple host/port information of service.

​	GetEndpoints 返回服务的终结点。端点包含多个主机/端口服务信息。

#### (*LocalService) GetKey

```go
func (s *LocalService) GetKey() string
```

GetKey formats and returns a unique key string for service. The result key is commonly used for key-value registrar server.

​	GetKey 格式化并返回服务的唯一键字符串。结果键通常用于键值注册器服务器。

#### (*LocalService) GetMetadata

```go
func (s *LocalService) GetMetadata() Metadata
```

GetMetadata returns the Metadata map of service. The Metadata is key-value pair map specifying extra attributes of a service.

​	GetMetadata 返回服务的元数据映射。元数据是键值对映射，用于指定服务的额外属性。

#### (*LocalService) GetName

```go
func (s *LocalService) GetName() string
```

GetName returns the name of the service. The name is necessary for a service, and should be unique among services.

​	GetName 返回服务的名称。该名称对于服务是必需的，并且在服务中应该是唯一的。

#### (*LocalService) GetPrefix

```go
func (s *LocalService) GetPrefix() string
```

GetPrefix formats and returns the key prefix string. The result prefix string is commonly used in key-value registrar server for service searching.

​	GetPrefix 格式化并返回键前缀字符串。结果前缀字符串通常用于键值注册器服务器中的服务搜索。

Take etcd server for example, the prefix string is used like: `etcdctl get /services/prod/hello.svc --prefix`

​	以 etcd server 为例，前缀字符串的使用方式如下： `etcdctl get /services/prod/hello.svc --prefix`

#### (*LocalService) GetValue

```go
func (s *LocalService) GetValue() string
```

GetValue formats and returns the value of the service. The result value is commonly used for key-value registrar server.

​	GetValue 格式化并返回服务的值。结果值通常用于键值注册器服务器。

#### (*LocalService) GetVersion

```go
func (s *LocalService) GetVersion() string
```

GetVersion returns the version of the service. It is suggested using GNU version naming like: v1.0.0, v2.0.1, v2.1.0-rc. A service can have multiple versions deployed at once. If no version set in service, the default version of service is “latest”.

​	GetVersion 返回服务的版本。建议使用 GNU 版本命名，例如：v1.0.0、v2.0.1、v2.1.0-rc。一个服务可以同时部署多个版本。如果未在服务中设置版本，则默认服务版本为“最新”。

### type Metadata

```go
type Metadata map[string]interface{}
```

Metadata stores custom key-value pairs.

​	元数据存储自定义键值对。

#### (Metadata) Get

```go
func (m Metadata) Get(key string) *gvar.Var
```

Get retrieves and returns value of specified key as gvar.

​	获取检索并返回指定键的值作为 gvar。

#### (Metadata) IsEmpty

```go
func (m Metadata) IsEmpty() bool
```

IsEmpty checks and returns whether current Metadata is empty.

​	IsEmpty 检查并返回当前元数据是否为空。

#### (Metadata) Set

```go
func (m Metadata) Set(key string, value interface{})
```

Set sets key-value pair into metadata.

​	将键值对设置为元数据。

#### (Metadata) Sets

```go
func (m Metadata) Sets(kvs map[string]interface{})
```

Sets sets key-value pairs into metadata.

​	Sets 将键值对设置为元数据。

### type Registrar

```go
type Registrar interface {
	// Register registers `service` to Registry.
	// Note that it returns a new Service if it changes the input Service with custom one.
	Register(ctx context.Context, service Service) (registered Service, err error)

	// Deregister off-lines and removes `service` from the Registry.
	Deregister(ctx context.Context, service Service) error
}
```

Registrar interface for service registrar.

​	服务注册器的注册器接口。

### type Registry

```go
type Registry interface {
	Registrar
	Discovery
}
```

Registry interface for service.

​	服务的注册表接口。

#### func GetRegistry

```go
func GetRegistry() Registry
```

GetRegistry returns the default Registry that is previously set. It returns nil if no Registry is set.

​	GetRegistry 返回以前设置的默认注册表。如果未设置注册表，则返回 nil。

### type SearchInput

```go
type SearchInput struct {
	Prefix   string   // Search by key prefix.
	Name     string   // Search by service name.
	Version  string   // Search by service version.
	Metadata Metadata // Filter by metadata if there are multiple result.
}
```

SearchInput is the input for service searching.

​	SearchInput 是服务搜索的输入。

### type Service

```go
type Service interface {
	// GetName returns the name of the service.
	// The name is necessary for a service, and should be unique among services.
	GetName() string

	// GetVersion returns the version of the service.
	// It is suggested using GNU version naming like: v1.0.0, v2.0.1, v2.1.0-rc.
	// A service can have multiple versions deployed at once.
	// If no version set in service, the default version of service is "latest".
	GetVersion() string

	// GetKey formats and returns a unique key string for service.
	// The result key is commonly used for key-value registrar server.
	GetKey() string

	// GetValue formats and returns the value of the service.
	// The result value is commonly used for key-value registrar server.
	GetValue() string

	// GetPrefix formats and returns the key prefix string.
	// The result prefix string is commonly used in key-value registrar server
	// for service searching.
	//
	// Take etcd server for example, the prefix string is used like:
	// `etcdctl get /services/prod/hello.svc --prefix`
	GetPrefix() string

	// GetMetadata returns the Metadata map of service.
	// The Metadata is key-value pair map specifying extra attributes of a service.
	GetMetadata() Metadata

	// GetEndpoints returns the Endpoints of service.
	// The Endpoints contain multiple host/port information of service.
	GetEndpoints() Endpoints
}
```

Service interface for service definition.

​	用于服务定义的服务接口。

#### func Get

```go
func Get(ctx context.Context, name string) (service Service, err error)
```

Get retrieves and returns the service by service name.

​	Get 按服务名称检索并返回服务。

#### func GetAndWatch <-2.1.0

```go
func GetAndWatch(ctx context.Context, name string, watch ServiceWatch) (service Service, err error)
```

GetAndWatch is used to getting the service with custom watch callback function.

​	GetAndWatch 习惯于使用自定义监视回调函数获取服务。

#### func GetAndWatchWithDiscovery <-2.3.3

```go
func GetAndWatchWithDiscovery(ctx context.Context, discovery Discovery, name string, watch ServiceWatch) (service Service, err error)
```

GetAndWatchWithDiscovery is used to getting the service with custom watch callback function in `discovery`.

​	GetAndWatchWithDiscovery 用于获取具有自定义监视回调函数的服务 `discovery` 。

#### func GetWithDiscovery <-2.3.3

```go
func GetWithDiscovery(ctx context.Context, discovery Discovery, name string) (service Service, err error)
```

GetWithDiscovery retrieves and returns the service by service name in `discovery`.

​	GetWithDiscovery 按 中的 `discovery` 服务名称检索并返回服务。

#### func NewServiceWithKV

```go
func NewServiceWithKV(key, value string) (Service, error)
```

NewServiceWithKV creates and returns a default implements for interface Service by key-value pair string.

​	NewServiceWithKV 通过键值对字符串创建并返回接口服务的默认实现。

#### func NewServiceWithName

```go
func NewServiceWithName(name string) Service
```

NewServiceWithName creates and returns a default implements for interface Service by service name.

​	NewServiceWithName 按服务名称创建并返回接口 Service 的默认实现。

#### func Register

```go
func Register(ctx context.Context, service Service) (Service, error)
```

Register registers `service` to default registry..

​	将寄 `service` 存器注册到默认注册表..

#### func Search

```go
func Search(ctx context.Context, in SearchInput) ([]Service, error)
```

Search searches and returns services with specified condition.

​	搜索 搜索并返回具有指定条件的服务。

### type ServiceWatch

```go
type ServiceWatch func(service Service)
```

ServiceWatch is used to watch the service status.

​	ServiceWatch 用于监视服务状态。

### type Watcher

```go
type Watcher interface {
	// Proceed proceeds watch in blocking way.
	// It returns all complete services that watched by `key` if any change.
	Proceed() (services []Service, err error)

	// Close closes the watcher.
	Close() error
}
```

Watcher interface for service.

​	服务的观察程序接口。

#### func Watch

```go
func Watch(ctx context.Context, key string) (Watcher, error)
```

Watch watches specified condition changes.

​	手表监视指定的条件变化。