+++
title = "go 发布周期(v18)"
weight = 25
date = 2023-05-18T17:26:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go Release Cycle - Go 发布周期

> 原文：[https://github.com/golang/go/wiki/Go-Release-Cycle](https://github.com/golang/go/wiki/Go-Release-Cycle)

Ian Lance Taylor edited this page on Jun 29 · [18 revisions](https://github.com/golang/go/wiki/Go-Release-Cycle/_history)

​	本维基页面由Go团队维护。请向 golang-dev [发送评论](https://groups.google.com/group/golang-dev)或[提交问题](https://go.dev/issue)，而不是直接进行修改。

​	短链接：https://go.dev/s/release。

## 概述

​	在Go 1.0和Go 1.1之间经过了14个月的时间，Go团队采用了一个发布时间表来简化工作、完成和发布的过程。总体目标是每六个月发布一个主要版本，分为三个月的总体开发，然后是三个月的测试和完善，称为发布冻结期。一个版本的维护是通过发布小版本来纠正关键问题，如崩溃或安全问题。

​	请注意，这个页面记录了我们在即将发布的版本中打算做什么。如果您熟悉我们过去所做的事情，请特别注意[历史部分](#history)所描述的差异。

## 时间轴

The current release cycle is aligned to start on February 1 and August 1 of each year. The target milestones for a release cycle are as described below. We try to hit the targets as closely as possible, while still delivering a quality release.

​	目前的发布周期被安排在每年的2月1日和8月1日开始。一个发布周期的目标里程碑如下所述。我们尽可能地达到目标，同时仍然提供一个高质量的版本。

Betas, release candidates, and releases have traditionally been cut and issued mid-week, often on Wednesdays. We avoid Mondays to give the Go team a work day before to prepare for the cut, and we avoid Fridays to give the Go team a work day after in case the cut runs into unexpected problems.

传统上，测试版、候选版和发行版都是在周中切割和发布的，通常在星期三。我们避开周一，以便给Go团队一个工作日的时间来准备剪接，我们避开周五，以便给Go团队一个工作日的时间，以防剪接时遇到意外问题。

![img](GoReleaseCycle18_img/release-cycle.png)

#### January 15 / July 15: Planning for release begins. 

#### 1月15日/7月15日：开始规划发布工作。

Planning of major work for upcoming release cycle is announced on [golang-dev](https://groups.google.com/group/golang-dev).

在golang-dev上公布即将到来的发布周期的主要工作规划。

Examples: [Go 1.4](https://groups.google.com/d/msg/golang-dev/eqBihsj7x-k/3h59pc8FDAoJ), [Go 1.5](https://groups.google.com/d/msg/golang-dev/2ZUi792oztM/GNsw1i76LhsJ), [Go 1.6](https://groups.google.com/d/msg/golang-dev/vNboccLL95c/XOD3vHKOCQAJ).

例如。Go 1.4, Go 1.5, Go 1.6。

#### February 1 / August 1: Release work begins.

2月1日/8月1日：发布工作开始。

This milestone is delayed if the previous release has not yet been issued, but note that later milestones are not delayed. That is, failure to get one release out on time takes time away from the development cycle of the subsequent release.

如果前一个版本还没有发布，这个里程碑就会被推迟，但请注意，后面的里程碑不会被推迟。也就是说，如果不能按时发布一个版本，就会占用后续版本的开发周期的时间。

Note that incoming bug reports should be handled and bugs fixed during ordinary release work. It is not appropriate to leave all bug fixes for the release freeze. See the [history](https://github.com/golang/go/wiki/Go-Release-Cycle#history) section for more discussion.

注意传入的bug报告应该在普通的发布工作中处理并修复bug。把所有的bug修复工作留到版本冻结期是不合适的。更多讨论见历史部分。

#### May 1 / November 1: Release freeze begins.

5月1日/11月1日：发布冻结期开始。

This milestone begins the second half of the release cycle, the release freeze. The release freeze applies to the entire main repository as well as to the code in subrepositories that is needed to build the binaries included in the release, particularly godoc and all its dependencies in the tools subrepository.

这个里程碑开始了发布周期的后半段，即发布冻结期。发布冻结适用于整个主仓库以及子仓库中构建二进制文件所需的代码，特别是 godoc 及其在工具子仓库中的所有依赖项。

Changes that were mailed before the freeze can be submitted if they are reviewed promptly after the freeze begins. During the freeze, only bug fixes and doc updates are accepted. On occasion new work may be done during the freeze, but only in exceptional circumstances and typically only if the work was proposed and approved before the cutoff. Such changes must be low risk.

如果在冻结开始后及时审查，在冻结前寄出的修改可以提交。在冻结期间，只接受错误修复和文档更新。偶尔，新的工作可以在冻结期间完成，但只有在特殊情况下，而且通常只有在截止日期前提出和批准的情况下。这种变化必须是低风险的。

This part of the release cycle is focused on improving the quality of the release, by testing it and fixing bugs that are found. However, every fix must be evaluated to balance the benefit of a possible fix against the cost of now having not as well tested code (the fix) in the release. Early in the release cycle, the balance tends toward accepting a fix. Late in the release cycle, the balance tends toward rejecting a fix, unless a case can be made that the fix is both low risk and high reward.

发布周期的这一部分主要是通过测试和修复发现的错误来提高发布的质量。然而，每一个修复都必须被评估，以平衡可能的修复的好处和现在在版本中没有经过良好测试的代码（修复）的代价。在发布周期的早期，这种平衡倾向于接受一个修复。在发布周期的晚期，这种平衡倾向于拒绝修复，除非有理由认为该修复是低风险和高回报的。

Examples of low risk changes appropriate late in the cycle include changes to documentation and fixes to new features being introduced in the current release (since there is no chance of introducing a regression compared to an earlier release).

在周期后期适合的低风险修改的例子包括对文档的修改和对当前版本中引入的新功能的修复（因为与早期版本相比，没有引入回归的可能性）。

By the end of the first month of the freeze, nearly all known bugs should have been fixed or explicitly postponed (either to the next release or indefinitely). There should be few known bugs left, perhaps only the ones that have proven too elusive to chase down.

在冻结期的第一个月末，几乎所有已知的bug都应该被修复或明确地推迟（要么到下一个版本或无限期地推迟）。应该很少有已知的bug了，也许只有那些被证明太难追寻的bug。

#### June 1 / December 1: Beta 1 issued.

6月1日/12月1日：发布测试版1。

A beta release is meant to encourage testing to discover new bugs. Issuing a beta is an indication that the Go team has fixed nearly all the known bugs that are planned to be fixed for the release, and now it is time to look for as yet unknown bugs.

测试版的目的是鼓励测试，以发现新的错误。发布测试版表明Go团队已经修复了几乎所有计划在发布时修复的已知错误，现在是时候寻找尚不清楚的错误了。

The first beta includes a complete draft of the eventual release notes, but marked clearly as a draft to avoid confusion when people link to them on the internet.

第一个测试版包括最终发布说明的完整草案，但明确标记为草案，以避免人们在互联网上链接时产生混淆。

If a release is ahead of schedule, it is acceptable and even encouraged to issue a beta a few weeks ahead of this target.

如果一个版本是提前发布的，那么比这个目标提前几周发布一个测试版是可以接受的，甚至是鼓励的。

As bugs are reported and fixed, additional betas may be released, provided there are significant code changes to test anew. Typically betas should not be issued more frequently than two weeks. It is important not to issue too many betas, nor too many release candidates: we are asking our users for their time to help us test the release and must not waste their good will by making too many requests.

随着错误的报告和修复，如果有重大的代码变化需要重新测试，可以发布更多的测试版。一般来说，测试版的发布频率不应超过两周。重要的是不要发布太多的测试版，也不要发布太多的候选版：我们要求我们的用户花时间来帮助我们测试版本，不能因为提出太多的要求而浪费他们的善意。

A beta is not expected to be bug-free and should not be used in production settings where failures or misbehavior cannot be tolerated. Organizations can run integration or other tests against a beta or even use it in a canary setting, but they should be discouraged from deploying the beta for unrestricted production usage.

测试版不应该是没有错误的，也不应该在不能容忍失败或错误行为的生产环境中使用。组织可以对测试版进行集成或其他测试，甚至可以在金丝雀环境中使用它，但他们不应该将测试版部署到不受限制的生产中去。

#### July 1 / January 1: Release candidate 1 issued.

7月1日/1月1日：发布候选版本1。

A release candidate is meant to be as close as possible to the actual release bits. Issuing a release candidate is an indication that the Go team has high confidence that the tree is free of critical bugs.

候选发布版是为了尽可能地接近实际的发布位。发布候选版本表明Go团队对这棵树没有关键的错误有很高的信心。

Once a release candidate is issued, only documentation changes and changes to address critical bugs should be made. In general the bar for bug fixes at this point is even slightly higher than the bar for bug fixes in a minor release. We may prefer to issue a release with a known but very rare crash than to issue a release with a new but not production-tested fix.

一旦发布了候选版本，就应该只对文档进行修改和对关键错误进行修改。一般来说，这时的错误修复标准甚至比次要版本中的错误修复标准还要高一些。我们可能更愿意发布一个已知但非常罕见的崩溃的版本，而不是发布一个新的但没有经过生产测试的修复的版本。

If a release is ahead of schedule, it is acceptable and even encouraged to issue a release candidate a few weeks ahead of this target. Extended release testing is a great way to deliver a robust release.

如果一个版本比计划提前，可以接受甚至鼓励比这个目标提前几周发布候选版本。延长发布测试是提供一个强大的版本的好方法。

If critical bugs are reported and fixed, additional release candidates may be issued, but typically not more than one every two weeks.

如果关键的错误被报告和修复，可以发布额外的候选版本，但通常不会超过每两周一次。

Again, a release candidate is meant to be bug-free, as much as possible. Organizations are encouraged to deploy it in production settings after appropriate organization-specific testing.

同样，候选发布版的目的是要尽可能地没有错误。我们鼓励各组织在进行适当的特定测试后，将其部署在生产环境中。

One of the criteria for issuing a release candidate is that Google be using that version of the code for new production builds by default: if we at Google are not willing to run it for production use, we shouldn't be asking others to. We may issue the release candidate a few days in advance of Google changing over, depending on how the calendar falls. For example, the change inside Google makes more sense to do on Mondays, so we may issue the release candidate the Wednesday before or the Wednesday after Google converts to the new release by default.

发布候选版本的标准之一是，谷歌在新的生产构建中默认使用该版本的代码：如果我们谷歌不愿意在生产中使用它，我们就不应该要求其他人使用。我们可能会在谷歌改变之前提前几天发布候选版本，这取决于日历的变化。例如，谷歌内部的变化在周一进行更有意义，所以我们可能在谷歌默认转换为新版本的前一个或后一个星期三发布候选版本。



The calm period between a release candidate and the final release is a good time for additional testing or for discussing the next release (see the January 15 milestone above).

在候选版本和最终版本之间的平静期是进行额外测试或讨论下一个版本的好时机（见上文1月15日的里程碑）。

#### August 1 / February 1: Release issued.

8月1日/2月1日：发布。

Finally, the release itself!

最后，发行版本身

A release should not contain significant changes since the last release candidate: it is important that all code in the release has been well tested. Issuing a release is an indication that release testing has confirmed the release candidate's high confidence that the tree is free of critical bugs.

一个版本不应该包含自上一个候选版本以来的重大变化：重要的是，版本中的所有代码都经过了良好的测试。发布版本是一个迹象，表明发布测试已经证实了候选发布者的高度自信，即该树没有关键的错误。

One of the criteria for issuing a release is that the release candidate has been available for four weeks and any problems that need to be addressed have been.

发布版本的标准之一是候选版本已经有四个星期的时间，任何需要解决的问题都已经解决了。

If a release process has run ahead of schedule, with an early beta and early release candidate, release candidate testing should absorb any extra time, leaving the actual release on time, not early. This improves the stability of the release, and it also gives developers working on the Go release more time to think about and plan the next release before code changes start pouring in again.

如果一个发布过程已经提前进行，有了早期的测试版和早期的候选版，候选版的测试应该吸收任何额外的时间，让实际的发布准时进行，而不是提前进行。这可以提高版本的稳定性，同时也给从事Go版本工作的开发人员更多的时间来思考和计划下一个版本，然后再开始大量的代码修改。

If a release is behind schedule, it is acceptable (but certainly not ideal) to issue a release sooner than four weeks after the release candidate, but no sooner than two weeks after. Abbreviated release testing is a great way to deliver a buggy release.

如果一个版本落后于计划，早于候选版本四周发布是可以接受的（但肯定不理想），但不能早于两周。简略的发布测试是提供一个有缺陷的版本的好方法。

Because Google runs the release candidate as the default version of Go, four weeks of release testing means that at the least Google has been using this version of Go for four weeks before it becomes an official release. While Google's successful use does not guarantee the absence of problems, our experience has been that it certainly helps improve the quality of the release. We strongly encourage other organizations to test release candidates as aggressively as they are able and to report problems that they find.

因为谷歌将候选版本作为Go的默认版本运行，四周的发布测试意味着至少谷歌在成为正式发布版本之前已经使用这个版本的Go四周了。虽然谷歌的成功使用并不能保证没有问题，但我们的经验是，这肯定有助于提高版本的质量。我们强烈鼓励其他组织尽可能积极地测试候选版本，并报告他们发现的问题。

Once a release is issued, work on the next release, including code reviews and submission of new code, can begin, and the cycle repeats. Note that if a release is delayed, so is work on the next release.

一旦发布了一个版本，下一个版本的工作，包括代码审查和提交新的代码，就可以开始了，这个循环不断重复。请注意，如果一个版本被推迟，下一个版本的工作也会被推迟。

## Release Maintenance 发布维护

A minor release is issued to address one or more critical problems for which there is no workaround (typically related to stability or security). The only code changes included in the release are the fixes for the specific critical problems. Important documentation-only changes and safe test updates (such as disabling tests), may also be included as well, but nothing more.

一个次要的版本是为了解决一个或多个没有解决方法的关键问题而发布的（通常与稳定性或安全性有关）。该版本中唯一的代码修改是对特定关键问题的修复。重要的纯文档修改和安全测试更新（如禁用测试），也可能包括在内，但仅此而已。

Minor releases to address non-security problems for Go 1.x stop once Go 1.x+2 is released.

解决Go 1.x的非安全问题的小版本在Go 1.x+2发布后就停止了。

Minor releases to address security problems for Go 1.x stop once Go 1.x+2 is released. For more about security updates, see the [security policy](https://go.dev/security).

解决Go 1.x的安全问题的小版本在Go 1.x+2发布后就停止了。关于安全更新的更多信息，请参见安全政策。

See also the [MinorReleases](https://go.dev/wiki/MinorReleases) wiki page.

也请参见 MinorReleases 维基页面。

## Freeze Exceptions 冻结的例外情况

Any exceptions to the freeze must be communicated to and explicitly approved by the Go Release Team before the freeze. If you’d like to request an exception, please file an issue in the issue tracker with "[freeze exception]" as a suffix and include "CC @golang/release" ([example](https://go.dev/issue/42747)). We will address any requests on a case-by-case basis with a strong preference for not permitting changes after the freeze.

任何冻结的例外情况都必须在冻结前与Go发布团队沟通并得到明确的批准。如果你想申请一个例外，请在问题跟踪器中以"[冻结例外]"为后缀提交一个问题，并包括 "CC @golang/release"（示例）。我们将根据具体情况处理任何请求，强烈建议不允许在冻结后进行修改。

## History 历史

The Go release cycle was discussed and adopted after the fourteen month effort to release Go 1.1. Go 1.2, Go 1.3, and Go 1.4 followed a six-month cycle beginning and ending (alternately) on December 1 and June 1. After experience with calendar problems in that cycle, we extended Go 1.5's development phase by two months to shift the cycle to begin and end on February 1 and August 1, as described above.

Go的发布周期是在经过14个月的努力发布Go 1.1后讨论并通过的。Go 1.2、Go 1.3和Go 1.4遵循6个月的周期，从12月1日和6月1日交替开始和结束。在经历了该周期的日历问题后，我们将Go 1.5的开发阶段延长了两个月，将该周期转移到2月1日和8月1日开始和结束，如上所述。

The [original proposal](https://go.dev/s/release-old) did not contain enough detail about the milestones during the freeze, and over the course of a few releases development work took over much of the freeze. Compared to the goal set above of issuing a beta one month into the release freeze, the first betas for Go 1.3, Go 1.4, Go 1.5, and Go 1.6 were three, four, five, and six weeks late, respectively. (Go 1.6 beta 1 was only two weeks late, but it was full of known bugs that we still intended to fix, primarily to get something out for testing before the winter holidays. Go 1.6's first real beta by the above definition was beta 2.)

最初的建议没有包含关于冻结期间的里程碑的足够细节，在几个版本的过程中，开发工作占据了冻结期的大部分时间。与上面设定的在发布冻结期一个月内发布测试版的目标相比，Go 1.3、Go 1.4、Go 1.5和Go 1.6的第一个测试版分别晚了三、四、五和六周。(Go 1.6 beta1只晚了两周，但它充满了已知的错误，我们仍然打算修复，主要是为了在寒假前推出一些测试。按照上述定义，Go 1.6的第一个真正的测试版是测试版2）。

When the beta is late, everything that follows the beta—shaking out the final bugs, thorough testing of the release candidates, and the shipping of the release—gets rushed, leading to more bugs in the final release and usually a delay in starting the next cycle.

当测试版姗姗来迟时，测试版之后的一切工作--找出最后的错误，对候选发布版进行彻底的测试，以及发布版的发货--都会变得匆忙，导致最终发布版中出现更多的错误，并且通常会推迟开始下一个周期。

The beta was ready later and later in those four cycles primarily because we both postponed too many bugs to the freeze and then allowed too many non-essential bug fixes during the freeze.

在这四个周期中，测试版的准备时间越来越晚，主要是因为我们把太多的错误推迟到冻结期，然后在冻结期允许太多的非必要的错误修复。

For Go 1.7 and later we will need to make sure that bugs are fixed before the freeze. That is, we need to follow the schedule above, not what we've done in the past.

对于Go 1.7及以后的版本，我们将需要确保在冻结前修复bug。也就是说，我们需要遵循上面的时间表，而不是我们过去的做法。

