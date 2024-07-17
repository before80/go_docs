package main

import (
	"fmt"
	"reflect"
)

func isNil(x any) {
	fmt.Println("-----------------------")
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	// 这里不能使用 reflect.TypeOf(x).Name(),
	//因为 Name()只能获取定义在本包内的类型的名称，其他包定义的类型，获取得到的都是空字符串
	str := fmt.Sprintf("using reflect.Value.IsNil -> %s's default value", t.String())

	// 这里之所以加上 recover， 是因为 reflect.ValueOf(x).IsNil() 可能在
	if err := recover(); err != nil {
		fmt.Println(str, err)
	}
	if v.IsNil() {
		fmt.Println(str, "is nil.")
	} else {
		fmt.Println(str, "is not nil.")
	}
}

func GetTypeName(x any) string {
	// 这里不能使用 reflect.TypeOf(x).Name(),
	//因为 Name()只能获取定义在本包内的类型的名称，其他包定义的类型，获取得到的都是空字符串
	return reflect.TypeOf(x).String()
}

func main() {
	var sli1 []int
	isNil(sli1)
	if sli1 == nil {
		fmt.Printf("using == nil -> %s's default value is nil.\n", GetTypeName(sli1))
	} else {
		fmt.Printf("using == nil -> %s's default value is not nil.\n", GetTypeName(sli1))
	}

	var ps *string
	isNil(ps)
	if ps == nil {
		fmt.Printf("using == nil -> %s's default value is nil.\n", GetTypeName(ps))
	} else {
		fmt.Printf("using == nil -> %s's default value is not nil.\n", GetTypeName(ps))
	}

	var m1 map[string]int
	isNil(m1)
	if ps == nil {
		fmt.Printf("using == nil -> %s's default value is nil.\n", GetTypeName(m1))
	} else {
		fmt.Printf("using == nil -> %s's default value is not nil.\n", GetTypeName(m1))
	}

	var m2 = make(map[string]int)
	isNil(m2)
	if ps == nil {
		fmt.Printf("using == nil -> %s's default value is nil.\n", GetTypeName(m2))
	} else {
		fmt.Printf("using == nil -> %s's default value is not nil.\n", GetTypeName(m2))
	}

}
