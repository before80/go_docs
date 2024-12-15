+++
title = "Execution"
date = 2024-12-15T11:20:49+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Execution](https://github.com/smartystreets/goconvey/wiki/Execution)
>
> 收录该文档时间： `2024-12-15T11:20:49+08:00`

# Execution



Alex edited this page on Oct 10, 2018 · [4 revisions](https://github.com/smartystreets/goconvey/wiki/Execution/_history)

To run your tests, do what you do best in your terminal:

```
go test
```



**Example output:** (real output is colorized)

```
.....

5 assertions and counting

....

9 assertions and counting

PASS
ok  	github.com/smartystreets/goconvey/examples	0.022s
```



You can also use -v for verbose:

```
go test -v
```



**Example output:** (real output is colorized)

```
=== RUN TestScoring

  Subject: Bowling Game Scoring 
    Given a fresh score card 
      When all gutter balls are thrown 
        The score should be zero ✔
      When all throws knock down only one pin 
        The score should be 20 ✔
      When a spare is thrown 
        The score should include a spare bonus. ✔
      When a strike is thrown 
        The score should include a strike bonus. ✔
      When all strikes are thrown 
        The score should be 300. ✔

5 assertions and counting

--- PASS: TestScoring (0.00 seconds)
=== RUN TestSpec

  Subject: Integer incrementation and decrementation 
    Given a starting integer value 
      When incremented 
        The value should be greater by one ✔
        The value should NOT be what it used to be ✔
      When decremented 
        The value should be lesser by one ✔
        The value should NOT be what it used to be ✔

9 assertions and counting

--- PASS: TestSpec (0.00 seconds)
PASS
ok  	github.com/smartystreets/goconvey/examples	0.023s
```



### Auto-test and web UI



If you're tired of hitting ↑, Enter all the time, try [running tests automatically](https://github.com/smartystreets/goconvey/wiki/Auto-test).

If you're tired of the terminal altogether, check out the [web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) which displays test output elegantly in the browser.

### [Next](https://github.com/smartystreets/goconvey/wiki/Skip)



Finally, [learn about Skip](https://github.com/smartystreets/goconvey/wiki/Skip) to skip/ignore scopes and assertions.
