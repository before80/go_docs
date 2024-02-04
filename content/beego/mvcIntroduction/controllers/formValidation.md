+++
title = "formValidation"
date = 2024-02-04T09:58:57+08:00
weight = 11
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/validation/]({{< ref "/beego/mvcIntroduction/controllers/formValidation" >}})

# Form validation 表单验证



## Form validation 表单验证

The Form validation module is used for data validation and error collection.

​	表单验证模块用于数据验证和错误收集。

## Installing and testing 安装和测试

Installing:

​	安装：

```
go get github.com/beego/beego/v2/core/validation
```

Testing:

​	测试：

```
go test github.com/beego/beego/v2/core/validation
```

## Localization 本地化

In order to localize validation error messages, one might use `SetDefaultMessage` function of the `validation` package.

​	为了本地化验证错误消息，可以使用 `SetDefaultMessage` 包的 `validation` 函数。

Note that format markers (`%d`, `%s`) must be preserved in translated text to provide resulting messages with validation context values.

​	请注意，格式标记（ `%d` 、 `%s` ）必须保留在翻译后的文本中，以便为结果消息提供验证上下文值。

Default template messages are present in `validation.MessageTmpls` variable.

​	默认模板消息存在于 `validation.MessageTmpls` 变量中。

Simple message localization for Russian language:

​	俄语的简单消息本地化：

```go
import "github.com/beego/beego/v2/core/validation"

func init() {
	validation.SetDefaultMessage(map[string]string{
		"Required":     "Должно быть заполнено",
		"Min":          "Минимально допустимое значение %d",
		"Max":          "Максимально допустимое значение %d",
		"Range":        "Должно быть в диапазоне от %d до %d",
		"MinSize":      "Минимально допустимая длина %d",
		"MaxSize":      "Максимально допустимая длина %d",
		"Length":       "Длина должна быть равна %d",
		"Alpha":        "Должно состоять из букв",
		"Numeric":      "Должно состоять из цифр",
		"AlphaNumeric": "Должно состоять из букв или цифр",
		"Match":        "Должно совпадать с %s",
		"NoMatch":      "Не должно совпадать с %s",
		"AlphaDash":    "Должно состоять из букв, цифр или символов (-_)",
		"Email":        "Должно быть в правильном формате email",
		"IP":           "Должен быть правильный IP адрес",
		"Base64":       "Должно быть представлено в правильном формате base64",
		"Mobile":       "Должно быть правильным номером мобильного телефона",
		"Tel":          "Должно быть правильным номером телефона",
		"Phone":        "Должно быть правильным номером телефона или мобильного телефона",
		"ZipCode":      "Должно быть правильным почтовым индексом",
	})
}
```

## Examples: 示例：

Direct use:

​	直接使用：

```go
import (
    "github.com/beego/beego/v2/core/validation"
    "log"
)

type User struct {
    Name string
    Age int
}

func main() {
    u := User{"man", 40}
    valid := validation.Validation{}
    valid.Required(u.Name, "name")
    valid.MaxSize(u.Name, 15, "nameMax")
    valid.Range(u.Age, 0, 18, "age")

    if valid.HasErrors() {
        // If there are error messages it means the validation didn't pass
        // Print error message
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    // or use like this
    if v := valid.Max(u.Age, 140, "age"); !v.Ok {
        log.Println(v.Error.Key, v.Error.Message)
    }
    // Customize error messages
    minAge := 18
    valid.Min(u.Age, minAge, "age").Message("18+ only!!")
    // Format error messages
    valid.Min(u.Age, minAge, "age").Message("%d+", minAge)
}
```

Use through StructTag

​	通过 StructTag 使用

```go
import (
    "log"
    "strings"

    "github.com/beego/beego/v2/core/validation"
)

// Set validation function in "valid" tag
// Use ";" as the separator of multiple functions. Spaces accept after ";"
// Wrap parameters with "()" and separate parameter with ",". Spaces accept after ","
// Wrap regex match with "//"
// 
// 各个函数的结果的key值为字段名.验证函数名
type user struct {
    Id     int
    Name   string `valid:"Required;Match(/^Bee.*/)"` // Name can't be empty or start with Bee
    Age    int    `valid:"Range(1, 140)"` // 1 <= Age <= 140, only valid in this range
    Email  string `valid:"Email; MaxSize(100)"` // Need to be a valid Email address and no more than 100 characters.
    Mobile string `valid:"Mobile"` // Must be a valid mobile number
    IP     string `valid:"IP"` // Must be a valid IPv4 address
}

// If your struct implemented interface `validation.ValidFormer`
// When all tests in StructTag succeed, it will execute Valid function for custom validation
func (u *user) Valid(v *validation.Validation) {
    if strings.Index(u.Name, "admin") != -1 {
        // Set error messages of Name by SetError and HasErrors will return true
        v.SetError("Name", "Can't contain 'admin' in Name")
    }
}

func main() {
    valid := validation.Validation{}
    u := user{Name: "Beego", Age: 2, Email: "dev@beego.wiki"}
    b, err := valid.Valid(&u)
    if err != nil {
        // handle error
    }
    if !b {
        // validation does not pass
        // blabla...
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
}
```

Available validation functions in StrucTag:

​	StrucTag 中可用的验证函数：

- `Required` not empty. :TODO 不为空，即各个类型要求不为其零值
- `Min(min int)` minimum value. Valid type is `int`, all other types are invalid.
  `Min(min int)` 最小值。有效类型为 `int` ，所有其他类型无效。
- `Max(max int)` maximum value. Valid type is `int`, all other types are invalid.
  `Max(max int)` 最大值。有效类型为 `int` ，所有其他类型无效。
- `Range(min, max int)` Value range. Valid type is `int`, all other types are invalid.
  `Range(min, max int)` 值范围。有效类型为 `int` ，所有其他类型无效。
- `MinSize(min int)` minimum length. Valid type is `string slice`, all other types are invalid.
  `MinSize(min int)` 最小长度。有效类型为 `string slice` ，所有其他类型无效。
- `MaxSize(max int)` maximum length. Valid type is `string slice`, all other types are invalid.
  `MaxSize(max int)` 最大长度。有效类型为 `string slice` ，所有其他类型无效。
- `Length(length int)` fixed length. Valid type is `string slice`, all other types are invalid.
  `Length(length int)` 固定长度。有效类型为 `string slice` ，所有其他类型无效。
- `Alpha` alpha characters. Valid type is `string`, all other types are invalid.
  `Alpha` 字母字符。有效类型为 `string` ，所有其他类型无效。
- `Numeric` numerics. Valid type is `string`, all other types are invalid.
  `Numeric` 数字。有效类型为 `string` ，所有其他类型无效。
- `AlphaNumeric` alpha characters or numerics. Valid type is `string`, all other types are invalid.
  `AlphaNumeric` 字母字符或数字。有效类型为 `string` ，所有其他类型无效。
- `Match(pattern string)` regex matching. Valid type is `string`, all other types will be cast to string then match. (fmt.Sprintf("%v", obj).Match)
  `Match(pattern string)` 正则表达式匹配。有效类型为 `string` ，所有其他类型将转换为字符串然后匹配。（fmt.Sprintf("%v", obj).Match）
- `AlphaDash` alpha characters or numerics or `-_`. Valid type is `string`, all other types are invalid.
  `AlphaDash` 字母字符或数字或 `-_` 。有效类型为 `string` ，所有其他类型无效。
- `Email` Email address. Valid type is `string`, all other types are invalid.
  `Email` 电子邮件地址。有效类型为 `string` ，所有其他类型无效。
- `IP` IP address，Only support IPv4 address. Valid type is `string`, all other types are invalid.
  `IP` IP 地址，仅支持 IPv4 地址。有效类型为 `string` ，所有其他类型无效。
- `Base64` base64 encoding. Valid type is `string`, all other types are invalid.
  `Base64` base64 编码。有效类型为 `string` ，所有其他类型无效。
- `Mobile` mobile number. Valid type is `string`, all other types are invalid.
  `Mobile` 手机号码。有效类型为 `string` ，所有其他类型无效。
- `Tel` telephone number. Valid type is `string`, all other types are invalid.
  `Tel` 电话号码。有效类型为 `string` ，所有其他类型无效。
- `Phone` mobile number or telephone number. Valid type is `string`, all other types are invalid.
  `Phone` 手机号码或电话号码。有效类型为 `string` ，所有其他类型无效。
- `ZipCode` zip code. Valid type is `string`, all other types are invalid.
  `ZipCode` 邮政编码。有效类型为 `string` ，所有其他类型无效。

### API doc API 文档

Please see [Go Walker](http://gowalker.org/github.com/beego/beego/v2/core/validation)

​	请参阅 Go Walker
