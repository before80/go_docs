+++
title = "迈向go 2的下一步"
weight = 14
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Next steps toward Go 2 - 迈向go 2的下一步

https://go.dev/blog/go2-next-steps

Robert Griesemer, for the Go team
26 June 2019

## Status 状态

We’re well on the way towards the release of Go 1.13, hopefully in early August of this year. This is the first release that will include concrete changes to the language (rather than just minor adjustments to the spec), after a longer moratorium on any such changes.

我们正朝着发布Go 1.13的方向前进，希望能在今年8月初发布。这是在较长时间暂停任何语言变化后的第一个版本，它将包括对语言的具体变化（而不仅仅是对规范的微小调整）。

To arrive at these language changes, we started out with a small set of viable proposals, selected from the much larger list of [Go 2 proposals](https://github.com/golang/go/issues?utf8=✓&q=is%3Aissue+is%3Aopen+label%3AGo2+label%3AProposal), per the new proposal evaluation process outlined in the “[Go 2, here we come!](https://blog.golang.org/go2-here-we-come)” blog post. We wanted our initial selection of proposals to be relatively minor and mostly uncontroversial, to have a reasonably high chance of having them make it through the process. The proposed changes had to be backward-compatible to be minimally disruptive since [modules](https://blog.golang.org/using-go-modules), which eventually will allow module-specific language version selection, are not the default build mode quite yet. In short, this initial round of changes was more about getting the ball rolling again and gaining experience with the new process, rather than tackling big issues.

为了实现这些语言变化，我们从一小部分可行的提案开始，根据 "Go 2，我们来了！"博文中概述的新提案评估过程，从更大的Go 2提案清单中选出。我们希望我们最初选择的提案相对较小，而且大多没有争议，以便有一个合理的高机会让它们通过这一过程。建议的修改必须是向后兼容的，以尽量减少破坏性，因为模块，最终将允许特定模块的语言版本选择，还不是默认的构建模式。简而言之，这一轮最初的修改更多的是为了让球再次滚动，并在新的过程中获得经验，而不是解决大问题。

Our [original list of proposals](https://blog.golang.org/go2-here-we-come) – [general Unicode identifiers](https://go.dev/issue/20706), [binary integer literals](https://go.dev/issue/19308), [separators for number literals](https://go.dev/issue/28493), [signed integer shift counts](https://go.dev/issue/19113) – got both trimmed and expanded. The general Unicode identifiers didn’t make the cut as we didn’t have a concrete design document in place in time. The proposal for binary integer literals was expanded significantly and led to a comprehensive overhaul and modernization of [Go’s number literal syntax](https://go.dev/design/19308-number-literals). And we added the Go 2 draft design proposal on [error inspection](https://go.dev/design/go2draft-error-inspection), which has been [partially accepted](https://go.dev/issue/29934#issuecomment-489682919).

我们最初的建议清单--一般的Unicode标识符、二进制整数字、数字字的分隔符、有符号的整数移位计数--既被修剪又被扩展。一般的Unicode标识符没有入选，因为我们没有及时准备好具体的设计文件。关于二进制整数字的建议被大幅扩展，并导致Go的数字字面语法的全面修改和现代化。我们还增加了关于错误检查的Go 2设计提案草案，该草案已被部分接受。

With these initial changes in place for Go 1.13, it’s now time to look forward to Go 1.14 and determine what we want to tackle next.

随着Go 1.13的这些初步变化，现在是时候展望Go 1.14并确定我们接下来要解决的问题了。

## Proposals for Go 1.14 关于Go 1.14的建议

The goals we have for Go today are the same as in 2007: to [make software development scale](https://blog.golang.org/toward-go2). The three biggest hurdles on this path to improved scalability for Go are package and version management, better error handling support, and generics.

我们今天对Go的目标与2007年一样：使软件开发规模化。在提高Go的可扩展性的道路上，最大的三个障碍是包和版本管理、更好的错误处理支持和泛型。

With Go module support getting increasingly stronger, support for package and version management is being addressed. This leaves better error handling support and generics. We have been working on both of these and presented [draft designs](https://go.dev/design/go2draft) at last year’s GopherCon in Denver. Since then we have been iterating those designs. For error handling, we have published a concrete, significantly revised and simplified proposal (see below). For generics, we are making progress, with a talk (“Generics in Go” by Ian Lance Taylor) [coming up](https://www.gophercon.com/agenda/session/49028) at this year’s GopherCon in San Diego, but we have not reached the concrete proposal stage yet.

随着Go模块的支持越来越强，对包和版本管理的支持正在得到解决。这就留下了更好的错误处理支持和泛型。我们一直在努力解决这两个问题，并在去年丹佛的GopherCon上展示了设计草案。从那时起，我们一直在迭代这些设计。对于错误处理，我们已经发布了一个具体的、经过大幅修改和简化的建议（见下文）。对于泛型，我们正在取得进展，在今年圣地亚哥的GopherCon上，我们将有一个讲座（Ian Lance Taylor的 "Go中的泛型"），但我们还没有达到具体建议的阶段。

We also want to continue with smaller improvements to the language. For Go 1.14, we have selected the following proposals:

我们还想继续对语言进行小的改进。对于Go 1.14，我们已经选择了以下提案：

[#32437](https://go.dev/issue/32437). A built-in Go error check function, “try” ([design doc](https://go.dev/design/32437-try-builtin)).

#32437. 一个内置的Go错误检查函数，"try"（设计文档）。

This is our concrete proposal for improved error handling. While the proposed, fully backwards-compatible language extension is minimal, we expect an outsize impact on error handling code. This proposal has already attracted an enormous amount of comments, and it’s not easy to follow up. We recommend starting with the [initial comment](https://go.dev/issue/32437#issue-452239211) for a quick outline and then to read the detailed design doc. The initial comment contains a couple of links leading to summaries of the feedback so far. Please follow the feedback recommendations (see the “Next steps” section below) before posting.

这是我们对改进错误处理的具体建议。虽然提议的、完全向后兼容的语言扩展是最小的，但我们期望对错误处理代码产生巨大的影响。这个建议已经吸引了大量的评论，而且不容易跟进。我们建议从最初的评论开始，快速浏览一下，然后再阅读详细的设计文档。最初的评论包含了几个链接，通往迄今为止的反馈摘要。在发布之前，请遵循反馈建议（见下面的 "下一步 "部分）。

[#6977](https://go.dev/issue/6977). Allow embedding overlapping interfaces ([design doc](https://go.dev/design/6977-overlapping-interfaces)).

#6977. 允许嵌入重叠的界面（设计文档）。

This is an old, backwards-compatible proposal for making interface embedding more tolerant.

这是一个旧的、向后兼容的建议，用于使接口嵌入更加宽容。

[#32479](https://go.dev/issue/32479) Diagnose `string(int)` conversion in `go vet`.

#32479 诊断go vet中的string(int)转换。

The `string(int)` conversion was introduced early in Go for convenience, but it is confusing to newcomers (`string(10)` is `"\n"` not `"10"`) and not justified anymore now that the conversion is available in the `unicode/utf8` package. Since removing this conversion is not a backwards-compatible change, we propose to start with a `vet` error instead.

string(int)转换是Go中早期引入的，目的是为了方便，但它对新人来说是混乱的（string(10)是"\n "而不是 "10"），而且现在unicode/utf8包中已经有了这种转换，因此不再有理由。由于删除这个转换并不是一个向后兼容的变化，我们建议从一个审核错误开始替代。

[#32466](https://go.dev/issue/32466) Adopt crypto principles ([design doc](https://go.dev/design/cryptography-principles)).

#32466 采用加密原则（设计文档）。

This is a request for feedback on a set of design principles for cryptographic libraries that we would like to adopt. See also the related [proposal to remove SSLv3 support](https://go.dev/issue/32716) from `crypto/tls`.

这是一个关于我们希望采用的一套加密库设计原则的反馈请求。也请参见相关提案，即从crypto/tls中移除SSLv3支持。

## Next steps 接下来的步骤

We are actively soliciting feedback on all these proposals. We are especially interested in fact-based evidence illustrating why a proposal might not work well in practice, or problematic aspects we might have missed in the design. Convincing examples in support of a proposal are also very helpful. On the other hand, comments containing only personal opinions are less actionable: we can acknowledge them but we can’t address them in any constructive way. Before posting, please take the time to read the detailed design docs and prior feedback or feedback summaries. Especially in long discussions, your concern may have already been raised and discussed in earlier comments.

我们正在积极征求对所有这些建议的反馈。我们对基于事实的证据特别感兴趣，这些证据可以说明为什么某项建议在实践中不能很好地工作，或者我们在设计中可能遗漏的问题。支持提案的令人信服的例子也很有帮助。另一方面，只包含个人意见的评论是不太可行的：我们可以承认它们，但我们不能以任何建设性的方式解决它们。在发帖之前，请花时间阅读详细的设计文档和先前的反馈或反馈摘要。特别是在长时间的讨论中，你所关心的问题可能已经在先前的评论中提出和讨论过了。

Unless there are strong reasons to not even proceed into the experimental phase with a given proposal, we are planning to have all these implemented at the start of the [Go 1.14 cycle](https://go.dev/wiki/Go-Release-Cycle) (beginning of August, 2019) so that they can be evaluated in practice. Per the [proposal evaluation process](https://blog.golang.org/go2-here-we-come), the final decision will be made at the end of the development cycle (beginning of November, 2019).

除非有充分的理由甚至不进入实验阶段的给定提案，否则我们计划在Go 1.14周期开始时（2019年8月初）实现所有这些提案，以便在实践中对它们进行评估。根据提案的评估过程，最终决定将在开发周期结束时（2019年11月初）作出。

Thank you for helping make Go a better language!

感谢你帮助Go成为更好的语言!
