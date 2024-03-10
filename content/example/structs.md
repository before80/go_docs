+++
title = "struct"
date = 2023-08-07T13:35:18+08:00
weight = 18
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# struct

```go
package main

import "fmt"

type Op interface {
	Walk()
	Say()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Walk() {
	fmt.Println(p.Name, ",I can walk.")
}

func (p Person) Say() {
	fmt.Println("My Name is ", p.Name, ",and I am ", p.Age, " years old.")
}

type Student struct {
	Person
	Class string
}

func (s Student) DoHomework() {
	fmt.Println(s.Name, ",I need to do same homework!")
}

type Teacher struct {
	Person
	Course string
}

func (t Teacher) Walk() {
	fmt.Println(t.Name, ",I walk fast!")
}

func (t Teacher) Say() {
	fmt.Println("My Name is ", t.Name, ",and I am ", t.Age, " years old. I am a ", t.Course, " teacher.")
}

func (t Teacher) Teach() {
	fmt.Println(t.Name, ": Listen to me! Follow me! A B C ...")
}

func main() {
	fmt.Println("------------- p1 ------------------")
	p1 := Person{"PA", 18}
	p1.Walk()
	p1.Say()

	fmt.Println("------------- p2 ------------------")
	p2 := Person{Name: "PB", Age: 12}
	p2.Walk()
	p2.Say()

	s1 := Student{Person{"SA", 10}, "五年级1班"}
	s1.Walk()
	s1.Say()
	s1.DoHomework()

	fmt.Println("------------- s1 ------------------")
	s2 := Student{Person: Person{"SB", 11}, Class: "六年级2班"}
	s2.Walk()
	s2.Say()
	s2.DoHomework()

	fmt.Println("------------- s2 ------------------")
	s3 := Student{Person: Person{Name: "SC", Age: 12}, Class: "初一年1班"}
	s3.Walk()
	s3.Say()
	s3.DoHomework()

	fmt.Println("------------- t1 ------------------")
	t1 := Teacher{Person{"TA", 25}, "English"}
	t1.Walk()
	t1.Say()
	t1.Teach()

	fmt.Println("------------- t2 ------------------")
	t2 := Teacher{Person: Person{"TB", 26}, Course: "Math"}
	t2.Walk()
	t2.Say()
	t2.Teach()

	fmt.Println("------------- t3 ------------------")
	t3 := Teacher{Person: Person{Name: "TC", Age: 27}, Course: "Chinese"}
	t3.Walk()
	t3.Say()
	t3.Teach()

	t3.Age = 28
	t3.Say()

	t3.Person.Age = 29
	t3.Say()

	fmt.Println("------------- tp3 ------------------")
	tp3 := &t3
	tp3.Walk()
	tp3.Say()
	tp3.Teach()

	tp3.Age = 30
	tp3.Say()

	tp3.Person.Age = 31
	tp3.Say()

}

// Output:
//------------- p1 ------------------
//PA ,I can walk.
//My Name is  PA ,and I am  18  years old.
//------------- p2 ------------------
//PB ,I can walk.
//My Name is  PB ,and I am  12  years old.
//SA ,I can walk.
//My Name is  SA ,and I am  10  years old.
//SA ,I need to do same homework!
//------------- s1 ------------------
//SB ,I can walk.
//My Name is  SB ,and I am  11  years old.
//SB ,I need to do same homework!
//------------- s2 ------------------
//SC ,I can walk.
//My Name is  SC ,and I am  12  years old.
//SC ,I need to do same homework!
//------------- t1 ------------------
//TA ,I walk fast!
//My Name is  TA ,and I am  25  years old. I am a  English  teacher.
//TA : Listen to me! Follow me! A B C ...
//------------- t2 ------------------
//TB ,I walk fast!
//My Name is  TB ,and I am  26  years old. I am a  Math  teacher.
//TB : Listen to me! Follow me! A B C ...
//------------- t3 ------------------
//TC ,I walk fast!
//My Name is  TC ,and I am  27  years old. I am a  Chinese  teacher.
//TC : Listen to me! Follow me! A B C ...
//My Name is  TC ,and I am  28  years old. I am a  Chinese  teacher.
//My Name is  TC ,and I am  29  years old. I am a  Chinese  teacher.
//------------- tp3 ------------------
//TC ,I walk fast!
//My Name is  TC ,and I am  29  years old. I am a  Chinese  teacher.
//TC : Listen to me! Follow me! A B C ...
//My Name is  TC ,and I am  30  years old. I am a  Chinese  teacher.
//My Name is  TC ,and I am  31  years old. I am a  Chinese  teacher.

```

