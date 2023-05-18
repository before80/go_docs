+++
title = "真正的 go 项目：SmartTwitter 和web.go"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Real Go Projects: SmartTwitter and web.go - 真正的 go 项目：SmartTwitter 和web.go

https://go.dev/blog/smarttwitter

Michael Hoisie
19 October 2010

2010年10月19日

​	本周的文章由[*Michael Hoisie*](http://www.hoisie.com/)撰写。他是一位位于旧金山的程序员，是Go的早期采用者和几个流行的Go库的作者。他描述了使用Go的经历：

​	我是通过[Hacker News](http://news.ycombinator.com/)上的一篇文章认识到Go的。大约一个小时后，我就爱上了它。当时我正在一家网络创业公司工作，并且一直在使用Python开发内部测试应用程序。Go提供了速度、更好的并发支持和合理的Unicode处理，所以我很想将我的程序移植到这种语言中。当时还没有一种容易的方式来使用Go编写Web应用程序，因此我决定构建一个简单的Web框架[web.go](http://github.com/hoisie/web.go)。它是基于我之前使用过的流行Python框架[web.py](http://webpy.org/)建模的。在开发web.go期间，我参与了Go社区，提交了一堆错误报告，并对一些标准库包（主要是[http](https://go.dev/pkg/http/)和[json](https://go.dev/pkg/json/)）进行了修改。

​	几周后，我注意到web.go在Github上受到了关注。这令人惊讶，因为我从来没有真正推广这个项目。我认为简单、快速的Web应用程序有市场，而我认为Go可以填补这个市场空缺。

​	有一个周末我决定编写一个简单的Facebook应用程序：它会将您的Twitter状态更新重新发布到您的Facebook个人资料中。有一个官方的Twitter应用程序可以做到这一点，但它会重新发布所有内容，在您的Facebook feed中制造噪音。我的应用程序允许您过滤转发、提及、哈希标签、回复等等。这变成了[Smart Twitter](http://www.facebook.com/apps/application.php?id=135488932982)，目前已经拥有近90,000名用户。

​	整个程序都是用Go编写的，并使用[Redis](https://redis.io/)作为其存储后端。它非常快速和稳健。它目前每秒处理约两打推文，并且大量使用Go的通道。它在具有2GB RAM的单个虚拟专用服务器实例上运行，处理负载没有问题。Smart Twitter使用的CPU时间很少，几乎完全是内存绑定的，因为整个数据库都保存在内存中。在任何给定时间，大约有10个goroutine同时运行：一个接受HTTP连接，另一个从Twitter Streaming API中读取，一对用于错误处理，其余的则处理Web请求或重新发布传入的推文。

​	Smart Twitter还衍生出其他开源Go项目：[mustache.go](http://github.com/hoisie/mustache.go)、[redis.go](http://github.com/hoisie/redis.go)和[twitterstream](http://github.com/hoisie/twitterstream)。

​	我认为web.go还有很多工作要做。例如，我想为流式连接、WebSockets、路由过滤器、共享主机的更好支持和文档的改进增加更好的支持。最近我离开了初创公司从事软件自由职业工作，计划在可能的情况下使用Go。这意味着我可能会将其用作个人应用程序的后端，以及为喜欢使用尖端技术的客户的后端。

​	最后，我要感谢Go团队的所有努力。Go是一个很棒的平台，我认为它有着光明的未来。我希望看到语言围绕社区的需求不断成长。社区中正在发生很多有趣的事情，我期待着看到人们用这种语言编写的创新作品。
