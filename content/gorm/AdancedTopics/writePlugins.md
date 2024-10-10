+++
title = "编写插件"
date = 2023-10-28T14:36:59+08:00
weight = 12
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/write_plugins.html](https://gorm.io/docs/write_plugins.html)

## 回调 Callbacks

GORM itself is powered by `Callbacks`, it has callbacks for `Create`, `Query`, `Update`, `Delete`, `Row`, `Raw`, you could fully customize GORM with them as you want

​	GORM 本身由 `Callbacks` 提供支持，它具有用于 `Create`、`Query`、`Update`、`Delete`、`Row`、`Raw` 的回调，您可以根据需要完全自定义 GORM。

Callbacks are registered into the global `*gorm.DB`, not the session-level, if you require `*gorm.DB` with different callbacks, you need to initialize another `*gorm.DB`

​	回调被注册到全局的 `*gorm.DB` 中，而不是会话级别。如果您需要具有不同回调的 `*gorm.DB`，则需要初始化另一个 `*gorm.DB`。

### 注册回调 Register Callback

Register a callback into callbacks

​	将回调注册到回调中

``` go
func cropImage(db *gorm.DB) {
  if db.Statement.Schema != nil {
    // crop image fields and upload them to CDN, dummy code
    for _, field := range db.Statement.Schema.Fields {
      switch db.Statement.ReflectValue.Kind() {
      case reflect.Slice, reflect.Array:
        for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
          // Get value from field
          if fieldValue, isZero := field.ValueOf(db.Statement.Context, db.Statement.ReflectValue.Index(i)); !isZero {
            if crop, ok := fieldValue.(CropInterface); ok {
              crop.Crop()
            }
          }
        }
      case reflect.Struct:
        // Get value from field
        if fieldValue, isZero := field.ValueOf(db.Statement.Context, db.Statement.ReflectValue); !isZero {
          if crop, ok := fieldValue.(CropInterface); ok {
            crop.Crop()
          }
        }

        // Set value to field
        err := field.Set(db.Statement.Context, db.Statement.ReflectValue, "newValue")
      }
    }

    // All fields for current model
    db.Statement.Schema.Fields

    // All primary key fields for current model
    db.Statement.Schema.PrimaryFields

    // Prioritized primary key field: field with DB name `id` or the first defined primary key
    db.Statement.Schema.PrioritizedPrimaryField

    // All relationships for current model
    db.Statement.Schema.Relationships

    // Find field with field name or db name
    field := db.Statement.Schema.LookUpField("Name")

    // processing
  }
}

db.Callback().Create().Register("crop_image", cropImage)
// register a callback for Create process
```

### 删除回调 Delete Callback

Delete a callback from callbacks

​	从回调中删除回调

``` go
db.Callback().Create().Remove("gorm:create")
// delete callback `gorm:create` from Create callbacks
```

### 替换回调 Replace Callback

Replace a callback having the same name with the new one

​	用新函数替换具有相同名称的回调

``` go
db.Callback().Create().Replace("gorm:create", newCreateFunction)
// replace callback `gorm:create` with new function `newCreateFunction` for Create process
```

### 按顺序注册回调 Register Callback with orders

Register callbacks with orders

​	按顺序注册回调

``` go
// before gorm:create
db.Callback().Create().Before("gorm:create").Register("update_created_at", updateCreated)

// after gorm:create
db.Callback().Create().After("gorm:create").Register("update_created_at", updateCreated)

// after gorm:query
db.Callback().Query().After("gorm:query").Register("my_plugin:after_query", afterQuery)

// after gorm:delete
db.Callback().Delete().After("gorm:delete").Register("my_plugin:after_delete", afterDelete)

// before gorm:update
db.Callback().Update().Before("gorm:update").Register("my_plugin:before_update", beforeUpdate)

// before gorm:create and after gorm:before_create
db.Callback().Create().Before("gorm:create").After("gorm:before_create").Register("my_plugin:before_create", beforeCreate)

// before any other callbacks
db.Callback().Create().Before("*").Register("update_created_at", updateCreated)

// after any other callbacks
db.Callback().Create().After("*").Register("update_created_at", updateCreated)
```

### 定义的回调 Defined Callbacks

GORM has defined [some callbacks](https://github.com/go-gorm/gorm/blob/master/callbacks/callbacks.go) to power current GORM features, check them out before starting your plugins

​	GORM 有一些预定义的回调来增强当前 GORM 功能，在开始插件之前请先查看它们。

## 插件 Plugin

GORM provides a `Use` method to register plugins, the plugin needs to implement the `Plugin` interface

​	GORM 提供了一个 `Use` 方法来注册插件，插件需要实现 `Plugin` 接口。

``` go
type Plugin interface {
  Name() string
  Initialize(*gorm.DB) error
}
```

The `Initialize` method will be invoked when registering the plugin into GORM first time, and GORM will save the registered plugins, access them like:

​	`Initialize` 方法将在将插件注册到 GORM 时首次调用，GORM 将保存已注册的插件，如下所示：

``` go
db.Config.Plugins[pluginName]
```

Checkout [Prometheus](https://gorm.io/docs/prometheus.html) as example

​	以 [Prometheus](https://gorm.io/docs/prometheus.html)为例