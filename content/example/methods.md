+++
title = "method"
date = 2023-08-07T13:35:26+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# method

```go
package main

import (
	"fmt"
	"strings"
)

type MyInt int

func (mi MyInt) ToHexStr() string {
	return fmt.Sprintf("%X", mi)
}

type Student struct {
	Courses []string
	Name    string
}

func (s Student) SetCourses1(courses ...string) {
	s.Courses = courses
}

func (s Student) SetName1(name string) {
	s.Name = name
}

func (s *Student) SetCourses2(courses ...string) {
	s.Courses = courses
}

func (s *Student) SetName2(name string) {
	s.Name = name
}

func (s Student) Say1() {
	fmt.Println(s.Name, ",I had learnt courses: ", strings.Join(s.Courses, "、"), ".")
}

func (s Student) Say2() {
	fmt.Println(s.Name, ",I had learnt courses: ", strings.Join(s.Courses, "、"), ".")
}

func main() {
	var i MyInt = 15
	fmt.Println(i.ToHexStr()) // F

	i = 16
	fmt.Println(i.ToHexStr()) // 10

	ip := &i
	fmt.Println(ip.ToHexStr()) // 10

	fmt.Println("---------------- s1 --------------------")
	s1 := Student{}
	s1.SetName1("A")
	s1.SetCourses1([]string{"Math", "Chinese"}...)
	s1.Say1() //  ,I had learnt courses:   .

	s1.SetName1("A")
	s1.SetCourses1([]string{"Math", "Chinese"}...)
	s1.Say2() //  ,I had learnt courses:   .

	s1.SetName2("A")
	s1.SetCourses1([]string{"Math", "Chinese"}...)
	s1.Say1() // A ,I had learnt courses:   .

	s1.SetName2("A")
	s1.SetCourses2([]string{"Math", "Chinese"}...)
	s1.Say1() // A ,I had learnt courses:  Math、Chinese .

	s1.SetName2("A")
	s1.SetCourses2([]string{"Math", "Chinese"}...)
	s1.Say2() // A ,I had learnt courses:  Math、Chinese .

	fmt.Println("---------------- s2 --------------------")
	s2 := &Student{}
	s2.SetName1("B")
	s2.SetCourses1([]string{"Math", "Chinese"}...)
	s2.Say1() //  ,I had learnt courses:   .

	s2.SetName1("B")
	s2.SetCourses1([]string{"Math", "Chinese"}...)
	s2.Say2() //  ,I had learnt courses:   .

	s2.SetName2("B")
	s2.SetCourses1([]string{"Math", "Chinese"}...)
	s2.Say1() // B ,I had learnt courses:   .

	s2.SetName2("B")
	s2.SetCourses2([]string{"Math", "Chinese"}...)
	s2.Say1() // B ,I had learnt courses:  Math、Chinese .

	s2.SetName2("B")
	s2.SetCourses2([]string{"Math", "Chinese"}...)
	s2.Say2() // B ,I had learnt courses:  Math、Chinese .

}

```

