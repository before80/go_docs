+++
title = "gvalid"
date = 2024-03-21T18:00:16+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/util/gvalid

Package gvalid implements powerful and useful data/form validation functionality.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func DeleteRule 

``` go
func DeleteRule(rules ...string)
```

DeleteRule deletes custom defined validation one or more rules and associated functions from global package.

##### func GetRegisteredRuleMap 

``` go
func GetRegisteredRuleMap() map[string]RuleFunc
```

GetRegisteredRuleMap returns all the custom registered rules and associated functions.

##### func GetTags 

``` go
func GetTags() []string
```

GetTags returns the validation tags.

##### func ParseTagValue 

``` go
func ParseTagValue(tag string) (field, rule, msg string)
```

ParseTagValue parses one sequence tag to field, rule and error message. The sequence tag is like: [alias@]rule[...#msg...]

##### func RegisterRule 

``` go
func RegisterRule(rule string, f RuleFunc)
```

RegisterRule registers custom validation rule and function for package.

##### func RegisterRuleByMap 

``` go
func RegisterRuleByMap(m map[string]RuleFunc)
```

RegisterRuleByMap registers custom validation rules using map for package.

### Types 

#### type CustomMsg 

``` go
type CustomMsg = map[string]interface{}
```

CustomMsg is the custom error message type, like: map[field] => string|map[rule]string

#### type Error 

``` go
type Error interface {
	Code() gcode.Code
	Current() error
	Error() string
	FirstItem() (key string, messages map[string]error)
	FirstRule() (rule string, err error)
	FirstError() (err error)
	Items() (items []map[string]map[string]error)
	Map() map[string]error
	Maps() map[string]map[string]error
	String() string
	Strings() (errs []string)
}
```

Error is the validation error for validation result.

#### type RuleFunc 

``` go
type RuleFunc func(ctx context.Context, in RuleFuncInput) error
```

RuleFunc is the custom function for data validation.

#### type RuleFuncInput 

``` go
type RuleFuncInput struct {
	// Rule specifies the validation rule string, like "required", "between:1,100", etc.
	Rule string

	// Message specifies the custom error message or configured i18n message for this rule.
	Message string

	// Field specifies the field for this rule to validate.
	Field string

	// ValueType specifies the type of the value, which might be nil.
	ValueType reflect.Type

	// Value specifies the value for this rule to validate.
	Value *gvar.Var

	// Data specifies the `data` which is passed to the Validator. It might be a type of map/struct or a nil value.
	// You can ignore the parameter `Data` if you do not really need it in your custom validation rule.
	Data *gvar.Var
}
```

RuleFuncInput holds the input parameters that passed to custom rule function RuleFunc.

#### type Validator 

``` go
type Validator struct {
	// contains filtered or unexported fields
}
```

Validator is the validation manager for chaining operations.

##### func New 

``` go
func New() *Validator
```

New creates and returns a new Validator.

##### Example

``` go
```
##### (*Validator) Assoc 

``` go
func (v *Validator) Assoc(assoc interface{}) *Validator
```

Assoc is a chaining operation function, which sets associated validation data for current operation. The optional parameter `assoc` is usually type of map, which specifies the parameter map used in union validation. Calling this function with `assoc` also sets `useAssocInsteadOfObjectAttributes` true

##### Example

``` go
```
##### (*Validator) Bail 

``` go
func (v *Validator) Bail() *Validator
```

Bail sets the mark for stopping validation after the first validation error.

##### Example

``` go
```
##### (*Validator) Ci 

``` go
func (v *Validator) Ci() *Validator
```

Ci sets the mark for Case-Insensitive for those rules that need value comparison.

##### Example

``` go
```
##### (*Validator) Clone 

``` go
func (v *Validator) Clone() *Validator
```

Clone creates and returns a new Validator which is a shallow copy of current one.

##### Example

``` go
```
##### (*Validator) Data 

``` go
func (v *Validator) Data(data interface{}) *Validator
```

Data is a chaining operation function, which sets validation data for current operation.

##### Example

``` go
```
##### (*Validator) Foreach <-2.2.0

``` go
func (v *Validator) Foreach() *Validator
```

Foreach tells the next validation using current value as an array and validates each of its element. Note that this decorating rule takes effect just once for next validation rule, specially for single value validation.

##### (*Validator) I18n 

``` go
func (v *Validator) I18n(i18nManager *gi18n.Manager) *Validator
```

I18n sets the i18n manager for the validator.

##### Example

``` go
```
##### (*Validator) Messages 

``` go
func (v *Validator) Messages(messages interface{}) *Validator
```

Messages is a chaining operation function, which sets custom error messages for current operation. The parameter `messages` can be type of string/[]string/map[string]string. It supports sequence in error result if `rules` is type of []string.

##### Example

``` go
```
##### (*Validator) RuleFunc 

``` go
func (v *Validator) RuleFunc(rule string, f RuleFunc) *Validator
```

RuleFunc registers one custom rule function to current Validator.

##### Example

``` go
```
##### (*Validator) RuleFuncMap 

``` go
func (v *Validator) RuleFuncMap(m map[string]RuleFunc) *Validator
```

RuleFuncMap registers multiple custom rule functions to current Validator.

##### Example

``` go
```
##### (*Validator) Rules 

``` go
func (v *Validator) Rules(rules interface{}) *Validator
```

Rules is a chaining operation function, which sets custom validation rules for current operation.

##### Example

``` go
```
##### (*Validator) Run 

``` go
func (v *Validator) Run(ctx context.Context) Error
```

Run starts validating the given data with rules and messages.

Example Run

``` go
package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	// check value mode
	if err := g.Validator().Data(16).Rules("min:18").Run(context.Background()); err != nil {
		fmt.Println("check value err:", err)
	}
	// check map mode
	data := map[string]interface{}{
		"passport":  "",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := map[string]string{
		"passport":  "required|length:6,16",
		"password":  "required|length:6,16|same:password2",
		"password2": "required|length:6,16",
	}
	if err := g.Validator().Data(data).Rules(rules).Run(context.Background()); err != nil {
		fmt.Println("check map err:", err)
	}
	// check struct mode
	type Params struct {
		Page      int    `v:"required|min:1"`
		Size      int    `v:"required|between:1,100"`
		ProjectId string `v:"between:1,10000"`
	}
	rules = map[string]string{
		"Page":      "required|min:1",
		"Size":      "required|between:1,100",
		"ProjectId": "between:1,10000",
	}
	obj := &Params{
		Page: 0,
		Size: 101,
	}
	if err := g.Validator().Data(obj).Run(context.Background()); err != nil {
		fmt.Println("check struct err:", err)
	}

	// May Output:
	// check value err: The value `16` must be equal or greater than 18
	// check map err: The passport field is required; The passport value `` length must be between 6 and 16; The password value `123456` must be the same as field password2
	// check struct err: The Page value `0` must be equal or greater than 1; The Size value `101` must be between 1 and 100
}

Output:
```

