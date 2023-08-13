+++
title = "text templates"
date = 2023-08-07T13:50:27+08:00
weight = 42
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# text templates

> 原文：https://gobyexample.com/text-templates

```go
package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text") // Value: some text
	t1.Execute(os.Stdout, 5)           // Value: 5
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	}) // Value: [Go Rust C++ C#]

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"}) // Name: Jane Doe

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	}) // Name: Mickey Mouse

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty") // yes
	t3.Execute(os.Stdout, "")          // no

	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")

	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		}) // Range: Go Rust C++ C#
}

```

