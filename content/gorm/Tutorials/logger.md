+++
title = "日志记录器"
date = 2023-10-28T14:31:06+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/logger.html](https://gorm.io/docs/logger.html)

Gorm has a [default logger implementation](https://github.com/go-gorm/gorm/blob/master/logger/logger.go), it will print Slow SQL and happening errors by default

​	Gorm有一个[默认的日志记录器实现](https://github.com/go-gorm/gorm/blob/master/logger/logger.go)，默认情况下会打印慢SQL和发生的错误。

The logger accepts few options, you can customize it during initialization, for example:

``` go
newLogger := logger.New(
  log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
  logger.Config{
    SlowThreshold:              time.Second,   // Slow SQL阈值 Slow SQL threshold
    LogLevel:                   logger.Silent, // Log级别 Log level
    IgnoreRecordNotFoundError: true,           // 为日志记录器忽略ErrRecordNotFound错误 Ignore ErrRecordNotFound error for logger
    ParameterizedQueries:      true,           // 不在SQL日志中包含参数 Don't include params in the SQL log
    Colorful:                  false,          // 禁用颜色 Disable color
  },
)

// 全局模式 Globally mode
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
  Logger: newLogger,
})

// 连续会话模式 Continuous session mode
tx := db.Session(&Session{Logger: newLogger})
tx.First(&user)
tx.Model(&user).Update("Age", 18)
```

### 日志级别 Log Levels

GORM defined log levels: `Silent`, `Error`, `Warn`, `Info`

​	GORM定义了以下日志级别：`Silent`，`Error`，`Warn`，`Info`

``` go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
  Logger: logger.Default.LogMode(logger.Silent),
})
```

### 调试 Debug

Debug a single operation, change current operation’s log level to logger.Info

​	调试单个操作，将当前操作的日志级别更改为logger.Info

``` go
db.Debug().Where("name = ?", "jinzhu").First(&User{})
```

## 自定义日志记录器 Customize Logger

Refer to GORM’s [default logger](https://github.com/go-gorm/gorm/blob/master/logger/logger.go) for how to define your own one

​	参考GORM的[默认日志记录器](https://github.com/go-gorm/gorm/blob/master/logger/logger.go)以了解如何定义自己的一个

The logger needs to implement the following interface, it accepts `context`, so you can use it for log tracing

​	日志记录器需要实现以下接口，它接受`context`，因此您可以将其用于日志追踪。

``` go
type Interface interface {
  LogMode(LogLevel) Interface
  Info(context.Context, string, ...interface{})
  Warn(context.Context, string, ...interface{})
  Error(context.Context, string, ...interface{})
  Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
}
```