+++
title = "2015年GopherCon综述"
weight = 5
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# GopherCon 2015 Roundup - 2015年GopherCon综述

https://go.dev/blog/gophercon2015

Andrew Gerrand
28 July 2015

A few weeks ago, Go programmers from around the world descended on Denver, Colorado for GopherCon 2015. The two-day, single-track conference attracted more than 1,250 attendees—nearly double last year’s number—and featured 22 talks presented by Go community members.

几周前，来自世界各地的Go程序员来到科罗拉多州丹佛市参加GopherCon 2015。这场为期两天的单轨会议吸引了超过1250名与会者--几乎是去年的两倍，并且有22场由Go社区成员主持的讲座。

![img](GopherCon2015Roundup_img/cowboy.jpg)

The Cowboy Gopher (a toy given to each attendee) watches over the ranch.

牛仔地鼠（送给每个参会者的玩具）在牧场上观看。

*Photograph by [Nathan Youngman](https://twitter.com/nathany/status/619861336399351808). Gopher by Renee French.*

摄影：Nathan Youngman。地鼠由Renee French拍摄。

Today the organizers have posted the videos online so you can now enjoy the conference from afar:

今天，组织者已经将视频发布到网上，所以你现在可以从远处欣赏到会议：

[Day 1](http://gophercon.com/schedule/8july/): 第1天。

- Go, Open Source, Community — Russ Cox ([video](https://www.youtube.com/watch?v=XvZOdpd_9tc)) ([text](https://blog.golang.org/open-source)) Go、开源、社区--Russ Cox (视频) (文字)
- Go kit: A Standard Library for Distributed Programming — Peter Bourgon ([video](https://www.youtube.com/watch?v=1AjaZi4QuGo)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Go kit/go-kit.pdf)) Go套件。分布式编程的标准库 - Peter Bourgon (视频) (幻灯片)
- Delve Into Go — Derek Parker ([video](https://www.youtube.com/watch?v=InG72scKPd4)) ([slides](http://go-talks.appspot.com/github.com/derekparker/talks/gophercon-2015/delve-into-go.slide)) 深入了解Go - Derek Parker (视频) (幻灯片)
- How a complete beginner learned Go as her first backend language in 5 weeks — Audrey Lim ([video](https://www.youtube.com/watch?v=fZh8uCInEfw)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Audrey Lim - How a Complete Beginner Picked Up Go as Her First Backend Language in 5 weeks/audreylim_slides.pdf)) 一个完全的初学者如何在5周内学会Go作为她的第一个后端语言 - Audrey Lim (视频) (幻灯片)
- A Practical Guide to Preventing Deadlocks and Leaks in Go — Richard Fliam ([video](https://www.youtube.com/watch?v=3EW1hZ8DVyw)) Go中防止死锁和泄漏的实用指南 - Richard Fliam (视频)
- Go GC: Solving the Latency Problem — Rick Hudson ([video](https://www.youtube.com/watch?v=aiv1JOfMjm0)) ([slides](https://go.dev/talks/2015/go-gc.pdf)) Go GC：解决延迟问题 - Rick Hudson (视频) (幻灯片)
- Simplicity and Go — Katherine Cox-Buday ([video](https://www.youtube.com/watch?v=S6mEo_FHZ5Y)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Katherine Cox-Buday: Simplicity %26 Go/Simplicity %26 Go.pdf)) 简洁性与Go - Katherine Cox-Buday (视频) (幻灯片)
- Rebuilding Parse.com in Go - an opinionated rewrite — Abhishek Kona ([video](https://www.youtube.com/watch?v=_f9LS-OWfeA)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Abhishek Kona Rewriting Parse in GO/myslides.pdf)) 在Go中重建Parse.com--一个有观点的重写--Abhishek Kona (视频) (幻灯片)
- Prometheus: Designing and Implementing a Modern Monitoring Solution in Go — Björn Rabenstein ([video](https://www.youtube.com/watch?v=1V7eJ0jN8-E)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Björn Rabenstein - Prometheus/slides.pdf)) 普罗米修斯。在Go中设计和实施现代监控解决方案 - Björn Rabenstein (视频) (幻灯片)
- What Could Go Wrong? — Kevin Cantwell ([video](https://www.youtube.com/watch?v=VC3QXZ-x5yI)) 什么会出错？- Kevin Cantwell (视频)
- The Roots of Go — Baishampayan Ghose ([video](https://www.youtube.com/watch?v=0hPOopcJ8-E)) ([slides](https://speakerdeck.com/bg/the-roots-of-go)) Go的根源 - Baishampayan Ghose (视频) (幻灯片)

[Day 2](http://gophercon.com/schedule/9july/): 第二天。

- The Evolution of Go — Robert Griesemer ([video](https://www.youtube.com/watch?v=0ReKdcpNyQg)) ([slides](https://go.dev/talks/2015/gophercon-goevolution.slide)) Go的演变 - Robert Griesemer (视频) (幻灯片)
- Static Code Analysis Using SSA — Ben Johnson ([video](https://www.youtube.com/watch?v=D2-gaMvWfQY)) ([slides](https://speakerdeck.com/benbjohnson/static-code-analysis-using-ssa)) 使用SSA进行静态代码分析 - Ben Johnson (视频) (幻灯片)
- Go on Mobile — Hana Kim ([video](https://www.youtube.com/watch?v=sQ6-HyPxHKg)) ([slides](https://go.dev/talks/2015/gophercon-go-on-mobile.slide)) 移动版Go - Hana Kim (视频) (幻灯片)
- Go Dynamic Tools — Dmitry Vyukov ([video](https://www.youtube.com/watch?v=a9xrxRsIbSU)) ([slides](https://go.dev/talks/2015/dynamic-tools.slide)) Go动态工具 - Dmitry Vyukov (视频) (幻灯片)
- Embrace the Interface — Tomás Senart ([video](https://www.youtube.com/watch?v=xyDkyFjzFVc)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Tomás Senart - Embrace the Interface/ETI.pdf)) 拥抱界面 - Tomás Senart (视频) (幻灯片)
- Uptime: Building Resilient Services with Go — Blake Caldwell ([video](https://www.youtube.com/watch?v=PyBJQA4clfc)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Blake Caldwell - Uptime: Building Resilient Services with Go/2015-GopherCon-Talk-Uptime.pdf)) 正常运行时间。用Go构建有弹性的服务 - Blake Caldwell (视频) (幻灯片) 
- Cayley: Building a Graph Database — Barak Michener ([video](https://www.youtube.com/watch?v=-9kWbPmSyCI)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Barak Michener - Cayley: Building a Graph Database/Cayley - Building a Graph Database.pdf)) Cayley。构建图形数据库 - Barak Michener (视频) (幻灯片)
- Code Generation For The Sake Of Consistency — Sarah Adams ([video](https://www.youtube.com/watch?v=kGAgHwfjg1s)) 为了一致性而进行的代码生成 - Sarah Adams (视频)
- The Many Faces of Struct Tags — Sam Helman and Kyle Erf ([video](https://www.youtube.com/watch?v=_SCRvMunkdA)) ([slides](https://github.com/gophercon/2015-talks/blob/master/Sam Helman %26 Kyle Erf - The Many Faces of Struct Tags/StructTags.pdf)) Struct Tags的多面性 - Sam Helman 和 Kyle Erf (视频) (幻灯片)
- Betting the Company on Go and Winning — Kelsey Hightower ([video](https://www.youtube.com/watch?v=wqVbLlHqAeY)) 将公司押在Go上并赢得胜利 - Kelsey Hightower (视频)
- How Go Was Made — Andrew Gerrand ([video](https://www.youtube.com/watch?v=0ht89TxZZnk)) ([slides](https://go.dev/talks/2015/how-go-was-made.slide)) Go是如何诞生的 - Andrew Gerrand (视频) (幻灯片)

The [hack day](http://gophercon.com/schedule/10july/) was also a ton of fun, with hours of [lightning talks](https://www.youtube.com/playlist?list=PL2ntRZ1ySWBeHqlHM8DmvS8axgbrpvF9b) and a range of activities from programming robots to a Magic: the Gathering tournament.

黑客日也非常有趣，有几个小时的闪电讲座和一系列活动，从机器人编程到《魔法：集会》比赛。

Huge thanks to the event organizers Brian Ketelsen and Eric St. Martin and their production team, the sponsors, the speakers, and the attendees for making this such a fun and action-packed conference. Hope to see you there next year!

非常感谢活动组织者Brian Ketelsen和Eric St. Martin以及他们的制作团队、赞助商、演讲者和参会者，他们使这次会议变得如此有趣和充满行动。希望明年能在那里见到你!
