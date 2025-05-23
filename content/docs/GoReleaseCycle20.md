+++
title = "Go 发布周期(v20)"
weight = 26
date = 2023-05-18T17:26:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go Release Cycle

> 原文：[https://github.com/golang/go/wiki/Go-Release-Cycle](https://github.com/golang/go/wiki/Go-Release-Cycle)

Heschi Kreinick edited this page last week · [20 revisions](https://github.com/golang/go/wiki/Go-Release-Cycle/_history)

​	本维基页面由Go团队维护。请向 golang-dev [发送评论](https://groups.google.com/group/golang-dev)或[提交问题](https://go.dev/issue)，而不是直接进行修改。

​	短链接：https://go.dev/s/release。

## 概述

​	Go每六个月发布一个版本。每个发布周期都分为大约4个月的开发阶段，然后是为期3个月的测试和打磨阶段，称为发布冻结期。如果一切顺利，下一个版本的工作将在上一个版本发布之前开始，导致大约一个月的重叠。

​	版本初始发布后，将通过修复严重的错误和安全问题来支持次要版本发布。

## 时间轴

​	当前发布周期的开始时间为每年的1月中旬和7月中旬。发布周期的目标里程碑如下所述。我们尽可能准确地达到目标，同时仍要发布质量高的版本。

​	为了给团队准备时间，并解决意外问题，我们更喜欢在工作日的早期或中期进行发布工作。这意味着确切的日期会因年份而异，因此里程碑被指定为特定月份的周数。第1周是从月份的第一个星期一开始的一周。所有日期都可能根据当年的假期时间而改变。

![release](https://user-images.githubusercontent.com/24611692/223832580-b613c098-cd8b-4d48-b5c4-cf349e7cc269.svg)

#### 1月/7月第1周：开始进行发布计划。

​	在[golang-dev](https://groups.google.com/group/golang-dev)上宣布即将到来的发布周期的重要工作的规划。

​	例如：[Go 1.20](https://groups.google.com/g/golang-dev/c/V8ez4YunkeE)

#### 1月/7月第3周：开始发布工作。

​	在先前的版本进入最后的稳定期后，代码库对一般开发人员开放。在此期间，欢迎进行各种类型的开发。最好在开发窗口结束前使大型或特别风险的更改生效，以便有时间修复可能出现的任何问题。

#### 5月/11月第4周：开始发布冻结。

​	该里程碑标志着发布周期的第二部分，发布冻结。发布冻结适用于整个主存储库，以及构建包含在发布中的二进制文件所需的子存储库中的代码，特别是vet和其工具子存储库中的所有依赖项。

​	在冻结期间，只接受修复错误和更新文档。偶尔在冻结期间可能会进行新的工作，但仅在特殊情况下，并且通常仅在截止日期之前提出和批准该工作。这些更改必须是低风险的。请参见下面的[冻结例外](#freeze-exceptions)。

​	发布周期的这一部分专注于通过测试和修复发现的错误来改善发布的质量。但是，必须评估每个修复，以平衡可能修复带来的好处与现在在发布中有不太经过测试的代码（修复）之间的成本。在发布周期的早期，平衡倾向于接受修复。在发布周期的后期，平衡倾向于拒绝修复，除非可以证明修复既是低风险又是高回报的情况。

​	在周期晚期适合进行的低风险更改的示例包括更改文档和修复当前版本中引入的新功能的错误（因为与较早版本相比不存在引入回归的可能性）。

​	在冻结开始后不久，几乎所有已知的错误都应该被修复或明确推迟（推迟到下一个版本或无限期推迟）。其余问题通常应被跟踪为发布阻塞，并紧急处理。

#### 6月/12月第2周：发布候选版本1。

​	发布候选版本的目的是尽可能接近实际的发布版本。发布候选版本意味着Go团队对代码树没有重要的bug有高度的信心。特别是因为Google持续跟踪Go的开发版本，所以在发布候选版本时，其近似版本至少已经在Google生产环境中运行了一两周。

​	一旦发布候选版本，只能进行文档更改和解决关键性bug的更改。在这一点上，对于bug修复的标准通常比小版本的bug修复标准甚至稍微更高。我们可能更倾向于发布已知但非常罕见的崩溃版本，而不是发布经过新但未经生产测试的修复版本。

​	如果有关键性bug被报告并解决，可能会发布其他的候选版本，但通常不会超过每两周发布一次。

​	再次强调，发布候选版本应尽可能不含bug。鼓励组织在适当的组织特定测试后部署它在生产环境中。

​	在发布候选版本和最终发布之间的平静时期是进行额外测试或讨论下一个发布的好时机（请参见上面的计划里程碑）。

#### 8月/2月第2周：发布版本。

​	最后，发布本身！

​	发布不应包含自上一个候选版本以来的重大更改：重要的是，发布版本的所有代码都经过了充分的测试。发布版本表明发布测试已经确认了候选版本的高度信心，即代码树没有关键性bug。

​	即使发布顺利，还有多余的时间，我们也更倾向于按计划进行。额外的测试只能改善版本的稳定性，并且它还给正在进行Go发布工作的开发人员更多的时间来考虑和计划下一个发布，以免代码改变再次涌现。

​	到最终发布的时候，Google将已经使用这个版本的Go近两个月了。虽然Google的成功使用并不能保证没有问题的存在，但我们的经验表明，它确实有助于提高发布的质量。我们强烈鼓励其他组织像他们所能的那样积极测试候选版本，并报告他们发现的问题。

​	一旦版本稳定下来，下一个版本的工作，包括代码审查和提交新代码，就可以开始了，周期就这样循环下去。需要注意的是，如果发布被延迟，下一个发布的工作也可能被延迟。

## 发布维护

​	次要版本发布旨在解决一个或多个关键问题，而没有解决方案（通常与稳定性或安全性相关）。版本发布中包括的唯一代码更改是针对特定关键问题的修复。重要的仅文档更改和安全的测试更新（例如禁用测试）也可以包括在内，但不会有更多内容。次要版本尽可能保持向后兼容性，并且不会引入新的API。

​	针对Go 1.x的次要版本用于解决问题（包括安全问题），一旦发布了Go 1.x+2，就会停止。有关安全更新的更多信息，请参见[安全策略](../GoSecurity/GoSecurity)。

​	另请参阅[MinorReleases](https://go.dev/wiki/MinorReleases)维基页面。

## 冻结例外 Freeze Exceptions

​	在冻结之前，任何冻结例外都必须与Go Release团队沟通并明确批准。如果您想请求例外，请在问题跟踪器中提交一个带有"[freeze exception]"后缀的问题，并包括"CC @golang/release"（[example](https://go.dev/issue/42747)）。我们将根据情况逐个处理任何请求，并强烈建议在冻结后不允许更改。

## 历史notes

​	这个时间表的一个版本，具有较短的开发窗口，最初是在2016年Go 1.7版本中采用的。经过多年的困难版本发布，在2022年和2023年的测试和流程改进导致了及时的1.19发布。对于1.20，开发窗口进行了扩展，进行了晚期冻结和早期解冻。这些变化已为1.21版本正式化。我们预计将继续按时发布。