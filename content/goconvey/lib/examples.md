+++
title = "examples"
date = 2024-12-15T21:20:15+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/examples](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/examples)
>
> 收录该文档时间： `2024-12-15T21:20:15+08:00`

## Overview 

Package examples contains, well, examples of how to use goconvey to specify behavior of a system under test. It contains a well-known example by Robert C. Martin called "Bowling Game Kata" as well as another very trivial example that demonstrates Reset() and some of the assertions.

### Index 

- [type Game](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/examples#Game)
- - [func NewGame() *Game](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/examples#NewGame)
- - [func (self *Game) Roll(pins int)](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/examples#Game.Roll)
  - [func (self *Game) Score() (sum int)](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/examples#Game.Score)

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Game 

``` go
type Game struct {
	// contains filtered or unexported fields
}
```

Game contains the state of a bowling game.

### func NewGame 

``` go
func NewGame() *Game
```

NewGame allocates and starts a new game of bowling.

#### (*Game) Roll 

``` go
func (self *Game) Roll(pins int)
```

Roll rolls the ball and knocks down the number of pins specified by pins.

#### (*Game) Score 

``` go
func (self *Game) Score() (sum int)
```

Score calculates and returns the player's current score.
