+++
title = "JSON"
date = 2023-08-07T13:50:55+08:00
weight = 44
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# JSON

> 原文：https://gobyexample.com/json

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Printf("%v,%T,%q\n", bolB, bolB, string(bolB)) // [116 114 117 101],[]uint8,"true"

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 1

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) // 2.34

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB)) // "gopher"

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Printf("%v,%T,%q\n", slcB, slcB, string(slcB)) // [91 34 97 112 112 108 101 34 44 34 112 101 97 99 104 34 44 34 112 101 97 114 34 93],[]uint8,"[\"apple\",\"peach\",\"pear\"]"
	fmt.Println(string(slcB))                          // ["apple","peach","pear"]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Printf("%v,%T,%q\n", mapB, mapB, string(mapB)) // [123 34 97 112 112 108 101 34 58 53 44 34 108 101 116 116 117 99 101 34 58 55 125],[]uint8,"{\"apple\":5,\"lettuce\":7}"
	fmt.Println(string(mapB))                          // {"apple":5,"lettuce":7}

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // {"page":1,"fruits":["apple","peach","pear"]}

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat) // map[num:6.13 strs:[a b]]

	num := dat["num"].(float64)
	fmt.Println(num) // 6.13

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)

	fmt.Println(str1) // a

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}

	json.Unmarshal([]byte(str), &res)
	fmt.Printf("%#v,%T\n", res, res) // main.response2{Page:1, Fruits:[]string{"apple", "peach"}},main.response2
	fmt.Println(res.Fruits[0])       // apple

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // {"apple":5,"lettuce":7}
	
}

```

