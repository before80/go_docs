+++
title = "关于"
date = 2024-01-25T17:43:19+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://www.lua.org/about.html](https://www.lua.org/about.html)

## What is Lua? Lua 是什么？

Lua is a powerful, efficient, lightweight, embeddable scripting language. It supports procedural programming, object-oriented programming, functional programming, data-driven programming, and data description.

​	Lua 是一种功能强大、高效、轻量级、可嵌入的脚本语言。它支持过程式编程、面向对象编程、函数式编程、数据驱动编程和数据描述。

Lua combines simple procedural syntax with powerful data description constructs based on associative arrays and extensible semantics. Lua is dynamically typed, runs by interpreting bytecode with a register-based virtual machine, and has automatic memory management with incremental garbage collection, making it ideal for configuration, scripting, and rapid prototyping.

​	Lua 将简单的过程语法与基于关联数组和可扩展语义的强大数据描述结构结合在一起。Lua 是动态类型的，通过使用基于寄存器的虚拟机解释字节码来运行，并具有增量垃圾回收的自动内存管理，使其非常适合配置、脚本编写和快速原型设计。

## Where does Lua come from? Lua 来自哪里？

Lua is designed, implemented, and maintained by a [team](https://www.lua.org/authors.html) at [PUC-Rio](http://www.puc-rio.br/), the Pontifical Catholic University of Rio de Janeiro in Brazil. Lua was born and raised in [Tecgraf](https://www.tecgraf.puc-rio.br/), formerly the Computer Graphics Technology Group of PUC-Rio. Lua is now housed at [LabLua](http://www.lua.inf.puc-rio.br/), a laboratory of the [Department of Computer Science](http://www.inf.puc-rio.br/) of PUC-Rio.

​	Lua 由巴西里约热内卢教皇天主教大学 (PUC-Rio) 的一个团队设计、实施和维护。Lua 诞生并成长于 Tecgraf，以前是 PUC-Rio 的计算机图形技术组。Lua 目前位于 PUC-Rio 计算机科学系的实验室 LabLua。

## What's in a name? 名字的由来？

"Lua" (pronounced **LOO-ah**) means "Moon" in Portuguese. As such, it is neither an acronym nor an abbreviation, but a noun. More specifically, "Lua" is a name, the name of the Earth's moon and the name of the language. Like most names, it should be written in lower case with an initial capital, that is, "Lua". Please do not write it as "LUA", which is both ugly and confusing, because then it becomes an acronym with [different meanings](http://acronyms.thefreedictionary.com/lua) for different people. So, please, write "Lua" right!

​	“Lua”（发音为 LOO-ah）在葡萄牙语中意为“月亮”。因此，它既不是首字母缩写词，也不是缩写，而是一个名词。更具体地说，“Lua”是一个名字，是地球的月亮的名字，也是这门语言的名字。像大多数名字一样，它应该用小写字母开头，即“Lua”。请不要将其写成“LUA”，这既难看又容易混淆，因为它会变成一个首字母缩写词，对不同的人有不同的含义。所以，请正确地写“Lua”！

## Why choose Lua? 为什么选择 Lua？

### Lua is a proven, robust language -  Lua 是一种经过验证的、强大的语言

Lua has been used in [many industrial applications](https://en.wikipedia.org/wiki/Category:Lua_(programming_language)-scriptable_software) (e.g., [Adobe's Photoshop Lightroom](http://since1968.com/article/190/mark-hamburg-interview-adobe-photoshop-lightroom-part-2-of-2)), with an emphasis on embedded systems (e.g., the [Ginga](http://www.ginga.org.br/) middleware for digital TV in Brazil) and [games](https://en.wikipedia.org/wiki/Category:Lua_(programming_language)-scripted_video_games) (e.g., [World of Warcraft](http://www.wowwiki.com/Lua) and Angry Birds). Lua is currently [the leading scripting language in games](http://www.satori.org/2009/03/the-engine-survey-general-results/). Lua has a solid [reference manual](https://www.lua.org/manual/5.4/) and there are [several books about it](https://www.lua.org/docs.html#books). Several [versions](https://www.lua.org/versions.html) of Lua have been released and used in [real applications](https://sites.google.com/site/marbux/home/where-lua-is-used) since its creation in 1993. Lua featured in [HOPL III, the Third ACM SIGPLAN History of Programming Languages Conference](https://en.wikipedia.org/wiki/History_of_Programming_Languages#HOPL_III), in 2007. Lua won the [Front Line Award 2011](http://www.gamasutra.com/view/news/129084/Announcing_Game_Developer_magazines_2011_Front_Line_Award_winners.php) from the [Game Developers Magazine](https://www.gdmag.com/).

​	Lua 已用于许多工业应用（例如，Adobe 的 Photoshop Lightroom），重点是嵌入式系统（例如，巴西数字电视的 Ginga 中间件）和游戏（例如，魔兽世界和愤怒的小鸟）。Lua 目前是游戏中领先的脚本语言。Lua 有一个可靠的参考手册，还有几本关于它的书籍。自 1993 年创建以来，已经发布了多个版本的 Lua，并将其用于实际应用。Lua 在 2007 年的第三届 ACM SIGPLAN 编程语言历史会议 HOPL III 中亮相。Lua 赢得了游戏开发者杂志颁发的 2011 年前线奖。

### Lua is fast - Lua 很快

Lua has a deserved reputation for performance. To claim to be "as fast as Lua" is an aspiration of other scripting languages. Several benchmarks show Lua as the fastest language in the realm of interpreted scripting languages. Lua is fast not only in fine-tuned benchmark programs, but in real life too. Substantial fractions of large applications have been written in Lua.

​	Lua 以性能著称。声称“与 Lua 一样快”是其他脚本语言的愿望。多项基准测试表明 Lua 是解释型脚本语言领域中最快的语言。Lua 不仅在经过微调的基准程序中速度很快，而且在现实生活中也很速度。大量应用程序的实质性部分都是用 Lua 编写的。

If you need even more speed, try [LuaJIT](https://luajit.org/), an independent implementation of Lua using a just-in-time compiler.

​	如果您需要更快的速度，请尝试 LuaJIT，这是一个使用即时编译器的独立 Lua 实现。

### Lua is portable - Lua 是可移植的

Lua is [distributed](https://www.lua.org/download.html) in a small package and builds out-of-the-box in all platforms that have a standard C compiler. Lua runs on all flavors of Unix and Windows, on mobile devices (running Android, iOS, BREW, Symbian, Windows Phone), on embedded microprocessors (such as ARM and Rabbit, for applications like Lego MindStorms), on IBM mainframes, etc.

​	Lua 以一个小软件包的形式进行分发，并且可以在所有具有标准 C 编译器的平台上开箱即用。Lua 可以运行在所有版本的 Unix 和 Windows 上，在移动设备上（运行 Android、iOS、BREW、Symbian、Windows Phone），在嵌入式微处理器上（例如 ARM 和 Rabbit，用于乐高 MindStorms 等应用程序），在 IBM 大型机上，等等。

For specific reasons why Lua is a good choice also for constrained devices, read [this summary](http://lua-users.org/lists/lua-l/2007-11/msg00248.html) by Mike Pall. See also a [poster](http://www.schulze-mueller.de/download/lua-poster-090207.pdf) created by Timm Müller.

​	有关 Lua 为什么也是受限设备的良好选择这一问题的具体原因，请阅读 Mike Pall 的这份摘要。另请参阅 Timm Müller 创建的海报。

### Lua is embeddable - Lua 可嵌入

Lua is a fast language engine with small footprint that you can embed easily into your application. Lua has a simple and well documented API that allows strong integration with code written in other languages. It is easy to extend Lua with libraries written in other languages. It is also easy to extend programs written in other languages with Lua. Lua has been used to extend programs written not only in C and C++, but also in Java, C#, Smalltalk, Fortran, Ada, Erlang, and even in other scripting languages, such as Perl and Ruby.

​	Lua 是一个快速的小型语言引擎，您可以轻松地将其嵌入到您的应用程序中。Lua 具有简单且有据可查的 API，允许与用其他语言编写的代码进行强集成。使用其他语言编写的库可以轻松扩展 Lua。使用 Lua 扩展用其他语言编写的程序也很容易。Lua 不仅被用来扩展用 C 和 C++ 编写的程序，还被用来扩展用 Java、C#、Smalltalk、Fortran、Ada、Erlang 甚至其他脚本语言（例如 Perl 和 Ruby）编写的程序。

### Lua is powerful (but simple) - Lua 功能强大（但简单）

A fundamental concept in the design of Lua is to provide *meta-mechanisms* for implementing features, instead of providing a host of features directly in the language. For example, although Lua is not a pure object-oriented language, it does provide meta-mechanisms for implementing classes and inheritance. Lua's meta-mechanisms bring an economy of concepts and keep the language small, while allowing the semantics to be extended in unconventional ways.

​	Lua 设计中的一个基本概念是提供元机制来实现特性，而不是直接在语言中提供大量特性。例如，尽管 Lua 不是纯面向对象语言，但它确实提供了元机制来实现类和继承。Lua 的元机制带来了概念的经济性，并保持了语言的小巧，同时允许以非常规方式扩展语义。

### Lua is small - Lua 很小

Adding Lua to an application does not bloat it. The [tarball for Lua 5.4.6](https://www.lua.org/ftp/lua-5.4.6.tar.gz), which contains source code and documentation, takes 355K compressed and 1.4M uncompressed. The source contains around 30000 lines of C. Under 64-bit Linux, the Lua interpreter built with all standard Lua libraries takes 282K and the Lua library takes 470K.

​	将 Lua 添加到应用程序不会使应用程序膨胀。包含源代码和文档的 Lua 5.4.6 的 tarball 压缩后为 355K，未压缩为 1.4M。源代码包含大约 30000 行 C 代码。在 64 位 Linux 下，使用所有标准 Lua 库构建的 Lua 解释器占用 282K，Lua 库占用 470K。

### Lua is free - Lua 是免费的

Lua is free open-source software, distributed under a [very liberal license](https://www.lua.org/license.html) (the well-known MIT license). It may be used for any purpose, including commercial purposes, at absolutely no cost. Just [download](https://www.lua.org/download.html) it and use it.

​	Lua 是免费的开源软件，在非常宽松的许可证（著名的 MIT 许可证）下发行。它可以用于任何目的，包括商业目的，完全免费。只需下载并使用即可。



## Joining the community 加入社区

There are several [meeting places](https://www.lua.org/community.html#meeting) for the Lua [community](https://www.lua.org/community.html) where you can go to learn and help others and [contribute](https://www.lua.org/community.html#contributing) in other ways. One of the focal points is the [mailing list](https://www.lua.org/lua-l.html), which is very [active](https://www.lua.org/lua-l-stats.html) and friendly.

​	Lua 社区有几个聚会场所，您可以在那里学习、帮助他人并以其他方式做出贡献。其中一个重点是邮件列表，它非常活跃且友好。

You can meet part of the Lua community in person by attending a [Lua Workshop](https://www.lua.org/community.html#workshop).

​	您可以通过参加 Lua 研讨会亲自结识部分 Lua 社区成员。

## Supporting Lua 支持 Lua

You can help to [support the Lua project](https://www.lua.org/donations.html) by [buying a book](https://www.lua.org/donations.html#books) published by Lua.org and by [making a donation](https://www.lua.org/donations.html#donation).

​	您可以通过购买 Lua.org 出版的书籍和捐款来帮助支持 Lua 项目。

You can also help to spread the word about Lua by buying Lua products at [Zazzle](https://www.zazzle.com/Lua_Store).

​	您还可以在 Zazzle 购买 Lua 产品来帮助宣传 Lua。

Lua.org is an Amazon Associate and we get commissions for qualifying purchases made through links in this site.

​	Lua.org 是亚马逊的合作伙伴，我们通过此网站中的链接获得符合条件的购买佣金。