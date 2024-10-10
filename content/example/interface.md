+++
title = "interface"
date = 2023-08-07T13:35:34+08:00
weight = 20
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# interface

```go
package main

import "fmt"

type Biology interface {
	Move()
	GetEnergy()
}

type Cow struct {
	Food string
}

func (c Cow) SetFood1(food string) {
	c.Food = food
}

func (c *Cow) SetFood2(food string) {
	c.Food = food
}

func (c Cow) Move() {
	fmt.Println("Moving use four leg.")
}

func (c Cow) GetEnergy() {
	fmt.Println("Getting energy from eating ", c.Food, " .")
}

type Human struct {
	Name string
	Food string
}

func (h Human) SetName1(name string) {
	h.Name = name
}

func (h *Human) SetName2(name string) {
	h.Name = name
}

func (h Human) SetFood1(food string) {
	h.Food = food
}

func (h *Human) SetFood2(food string) {
	h.Food = food
}

func (h Human) Move() {
	fmt.Println(h.Name, ",Moving use two leg.")
}

func (h Human) GetEnergy() {
	fmt.Println(h.Name, ",Getting energy from eating ", h.Food, " .")
}

func main() {
	var bio Biology

	fmt.Println("------------- bio  Cow -------------")
	bio = Cow{}
	bio.Move() // Moving use four leg.
	//bio.SetFood1("grass") //编译报错：bio.SetFood1 undefined (type Biology has no field or method SetFood1)
	//bio.SetFood2("grass") //编译报错：bio.SetFood2 undefined (type Biology has no field or method SetFood2)
	bio.GetEnergy() // Getting energy from eating    .

	fmt.Println("------------- bio  Human -------------")
	bio = Human{}
	bio.Move()
	//bio.SetName1("A") //编译报错：bio.SetName1 undefined (type Biology has no field or method SetName1)
	//bio.SetName2("A") //编译报错：bio.SetName2 undefined (type Biology has no field or method SetName2)
	//bio.SetFood1("rice") //编译报错：bio.SetFood1 undefined (type Biology has no field or method SetFood1)
	//bio.SetFood2("rice") //编译报错：bio.SetFood2 undefined (type Biology has no field or method SetFood2)
	bio.GetEnergy() //  ,Getting energy from eating    .

	fmt.Println("------------- c Cow -------------")
	c := Cow{}
	c.Move() // Moving use four leg.
	c.SetFood1("grass")
	c.GetEnergy() // Getting energy from eating    .
	c.SetFood2("leaf")
	c.GetEnergy() // Getting energy from eating  leaf  .

	fmt.Println("------------- cp Cow -------------")
	cp := &Cow{}
	cp.Move() // Moving use four leg.
	cp.SetFood1("grass")
	cp.GetEnergy() // Getting energy from eating    .
	cp.SetFood2("leaf")
	cp.GetEnergy() // Getting energy from eating  leaf  .

	fmt.Println("------------- h Human -------------")
	h := Human{}
	h.SetName1("A")
	h.Move() //  ,Moving use two leg.
	h.SetName2("B")
	h.Move() // B ,Moving use two leg.

	h.SetFood1("rice")
	h.GetEnergy() // B ,Getting energy from eating    .
	h.SetFood2("fruit")
	h.GetEnergy() // B ,Getting energy from eating  fruit  .

	fmt.Println("------------- hp Human -------------")
	hp := &Human{}
	hp.SetName1("A")
	hp.Move() //  ,Moving use two leg.
	hp.SetName2("B")
	hp.Move() // B ,Moving use two leg.

	hp.SetFood1("rice")
	hp.GetEnergy() // B ,Getting energy from eating    .
	hp.SetFood2("fruit")
	hp.GetEnergy() // B ,Getting energy from eating  fruit  .

}

```

