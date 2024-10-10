+++
title = "Prometheus"
date = 2023-10-28T14:35:08+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/prometheus.html](https://gorm.io/docs/prometheus.html)

GORM provides Prometheus plugin to collect [DBStats](https://pkg.go.dev/database/sql?tab=doc#DBStats) or user-defined metrics

​	GORM 提供了 Prometheus 插件来收集 [DBStats](https://pkg.go.dev/database/sql?tab=doc#DBStats) 或用户自定义指标

https://github.com/go-gorm/prometheus

## 用法 Usage

``` go
import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "gorm.io/plugin/prometheus"
)

db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

db.Use(prometheus.New(prometheus.Config{
  DBName:          "db1", // 使用 `DBName` 作为指标标签 use `DBName` as metrics label
  RefreshInterval: 15,    // 使用 `DBName` 作为指标标签 Refresh metrics interval (default 15 seconds)
  PushAddr:        "prometheus pusher address", // 使用 `DBName` 作为指标标签 push metrics if `PushAddr` configured
  StartServer:     true,  // 启动 HTTP 服务器以暴露指标 start http server to expose metrics
  HTTPServerPort:  8080,  // 配置 HTTP 服务器端口，默认端口为 8080（如果已配置多个实例，仅第一个 `HTTPServerPort` 将用于启动服务器） configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
  MetricsCollector: []prometheus.MetricsCollector {
    &prometheus.MySQL{
      VariableNames: []string{"Threads_running"},
    },
  },  // 用户自定义指标 user defined metrics
}))
```

## 用户自定义指标 User-Defined Metrics

You can define your metrics and collect them with GORM Prometheus plugin, which needs to implements `MetricsCollector` interface

​	你可以定义自己的指标并使用 GORM Prometheus 插件收集它们，需要实现 `MetricsCollector` 接口

``` go
type MetricsCollector interface {
  Metrics(*Prometheus) []prometheus.Collector
}
```

### MySQL

GORM provides an example for how to collect MySQL Status as metrics, check it out [prometheus.MySQL](https://github.com/go-gorm/prometheus/blob/master/mysql.go)

​	GORM 提供了一个示例，展示了如何收集 MySQL 状态作为指标，请查看 [prometheus.MySQL](https://github.com/go-gorm/prometheus/blob/master/mysql.go)

``` go
&prometheus.MySQL{
  Prefix: "gorm_status_",
  // 指标名称前缀，默认为 `gorm_status_` Metrics name prefix, default is `gorm_status_`
  // 例如，Threads_running 的指标名称为 `gorm_status_Threads_running` For example, Threads_running's metric name is `gorm_status_Threads_running`
  Interval: 100,
  // 获取间隔，默认使用 Prometheus 的 RefreshInterval Fetch interval, default use Prometheus's RefreshInterval
  VariableNames: []string{"Threads_running"},
  // 从 SHOW STATUS 中选择变量，如果没有设置，则使用所有状态变量 Select variables from SHOW STATUS, if not set, uses all status variables
}
```