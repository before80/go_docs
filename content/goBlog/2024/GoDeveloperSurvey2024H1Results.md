+++
title = "Go 开发者调查 2024 年上半年结果"
date = 2024-05-30T10:13:24+08:00
weight = 950
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/survey2024-h1-results](https://go.dev/blog/survey2024-h1-results)

# Go Developer Survey 2024 H1 Results - Go 开发者调查 2024 年上半年结果

Alice Merrick and Todd Kulesza

​	爱丽丝·梅里克和托德·库莱扎
9 April 2024

​	2024 年 4 月 9 日

## Background 背景

This post shares the results of our most recent Go Developer Survey, conducted in January and February 2024. Along with capturing sentiments and challenges around using Go and Go tooling, our primary focus areas for this survey were about how developers are starting to use Go (or other languages) for AI-related use cases, and particular challenges for those who are learning Go or looking to expand their Go skill set.

​	这篇博文分享了我们最近在 2024 年 1 月和 2 月进行的 Go 开发者调查的结果。除了收集有关使用 Go 和 Go 工具的感受和挑战外，我们这次调查的主要关注点是开发者如何开始将 Go（或其他语言）用于与 AI 相关的用例，以及对于正在学习 Go 或希望扩展其 Go 技能集的人来说的特定挑战。

We recruited participants from the Go blog and through randomized prompts in the VS Code Go plug-in. This year, with the help of [JetBrains](https://jetbrains.com/), we also included a randomized survey prompt in the [GoLand IDE](https://jetbrains.com/go/), allowing us to recruit a more representative sample of Go developers. We received a total of 6,224 responses! A huge thank you to all those who contributed to making this possible.

​	我们从 Go 博客和 VS Code Go 插件中的随机提示中招募了参与者。今年，在 JetBrains 的帮助下，我们还在 GoLand IDE 中加入了一个随机调查提示，这让我们能够招募到更有代表性的 Go 开发者样本。我们总共收到了 6,224 份回复！非常感谢所有为实现这一目标做出贡献的人。

## Highlights 亮点

- Developer sentiment remains high, with 93% of respondents expressing satisfaction with Go over the past year.
  93% 的受访者对过去一年的 Go 表示满意，开发人员情绪依然高涨。
- A majority of respondents (80%) said they trust the Go team to “do what’s best” for developers like themselves when maintaining and evolving the language.
  大多数受访者 (80%) 表示他们相信 Go 团队在维护和发展语言时会“为像他们这样的开发者做最好的事情”。
- Among survey respondents who build AI-powered applications and services, there is a shared sense that Go is a strong platform for running these types of applications in production. For example, a majority of respondents working with AI-powered applications already use Go or would like to migrate to Go for their AI-powered workloads, and the most serious challenges developers encounter are related to the library and documentation ecosystems rather than the core language and runtime. That said, the most commonly documented paths for getting started are currently Python-centric, resulting in many organizations starting AI-powered work in Python before moving to a more production-ready language.
  在构建 AI 驱动的应用程序和服务的调查受访者中，有一种共同的看法，即 Go 是用于在生产中运行这些类型的应用程序的强大平台。例如，大多数使用 AI 驱动的应用程序的受访者已经使用 Go 或希望将其 AI 驱动的负载迁移到 Go，而开发人员遇到的最严重挑战与库和文档生态系统有关，而不是核心语言和运行时。也就是说，目前最常见的入门记录路径以 Python 为中心，导致许多组织在转向更适合生产的语言之前，先使用 Python 开始 AI 驱动的任务。
- The most common kinds of AI-powered services respondents are building include summarization tools, text generation tools, and chatbots. Responses suggest that many of these use cases are internal-facing, such as chatbots trained upon an organization’s internal documentation and intended to answer employee questions. We hypothesize that organizations are intentionally starting with internal use cases to develop in-house expertise with LLMs while avoiding potential public embarrassment when AI-powered agents behave unexpectedly.
  受访者构建的最常见的 AI 驱动的服务类型包括摘要工具、文本生成工具和聊天机器人。回复表明，其中许多用例都是面向内部的，例如根据组织内部文档训练的聊天机器人，旨在回答员工问题。我们假设组织有意从内部用例开始，以在避免 AI 驱动的代理意外行为导致潜在的公开尴尬的同时，培养内部LLMs专业知识。
- Lack of time or opportunities was the most commonly cited challenge for respondents to reaching their Go-related learning goals, suggesting that language learning is difficult to prioritize without a specific goal or business case in mind. The next most common challenge was in learning new best practices, concepts, and idioms that are particular to Go when coming from other language ecosystems.
  受访者最常提到的实现 Go 相关学习目标的挑战是时间或机会不足，这表明在没有具体目标或商业案例的情况下，很难将语言学习作为优先事项。其次最常见的挑战是在来自其他语言生态系统时学习 Go 特有的最佳实践、概念和惯用语。

## Contents 内容

- [Developer sentiment
  开发者情绪](https://go.dev/blog/survey2024-h1-results#sentiment)
- [Developer environments
  开发者环境](https://go.dev/blog/survey2024-h1-results#devenv)
- [Resource and performance priorities
  资源和性能优先级](https://go.dev/blog/survey2024-h1-results#priorities)
- [Understanding AI use cases for Go
  了解 Go 的 AI 用例](https://go.dev/blog/survey2024-h1-results#mlai)
- [Learning challenges
  学习挑战](https://go.dev/blog/survey2024-h1-results#learn)
- [Demographics
  人口统计学](https://go.dev/blog/survey2024-h1-results#demographics)
- [Firmographics
  人口统计](https://go.dev/blog/survey2024-h1-results#firmographics)
- [Methodology
  方法论](https://go.dev/blog/survey2024-h1-results#methodology)
- [Closing
  闭幕](https://go.dev/blog/survey2024-h1-results#closing)

## Developer sentiment 开发者情绪

Overall satisfaction remains high in the survey with 93% of respondents saying they were somewhat or very satisfied with Go during the last year. This isn’t surprising, considering our audience is those who have voluntarily taken our survey. But even among those who were randomly sampled from both VS Code and GoLand, we still see comparable rates of satisfaction (92%). Although the exact percentages fluctuate slightly from survey to survey, we do not see any statistically significant differences from [2023 H2](https://go.dev/blog/survey2023-h2-results), when the satisfaction rate was 90%.

​	在调查中，总体满意度仍然很高，93% 的受访者表示他们在去年对 Go 比较满意或非常满意。这并不奇怪，因为我们的受众是自愿参加我们调查的人。但即使在从 VS Code 和 GoLand 中随机抽取的人群中，我们仍然看到相当高的满意度 (92%)。尽管确切的百分比在不同的调查中略有波动，但我们没有看到与 2023 年下半年（当时满意度为 90%）有任何统计学上的显着差异。

![Chart of developer satisfaction with Go](./GoDeveloperSurvey2024H1Results_img/csat.svg+xml)

### Trust 信任

This year we introduced a new metric for measuring developer trust. This was an experimental question and its wording may change over time as we learn more about how respondents interpreted it. Because this is the first time we asked this question, we don’t have previous years to give us context for our results. We found that 80% of respondents somewhat or strongly agree that they trust the Go team to do what’s best for users like them. Respondents with 5 or more years of experience with Go tended to agree more (83%) than those with less than 2 years of experience (77%). This could reflect [survivorship bias](https://en.wikipedia.org/wiki/Survivorship_bias) in that those who trust the Go team more are more likely to continue using Go, or may reflect how trust is calibrated over time.

​	今年我们引入了一个新的指标来衡量开发人员的信任度。这是一个实验性问题，随着我们更多地了解受访者如何解释它，其措辞可能会随着时间的推移而改变。由于这是我们第一次提出这个问题，我们没有往年的数据为我们的结果提供背景。我们发现 80% 的受访者或多或少同意他们相信 Go 团队会为像他们这样的用户做最好的事情。拥有 5 年或以上 Go 经验的受访者往往比拥有不到 2 年经验的受访者（77%）更同意（83%）。这可能反映了幸存者偏差，即那些更信任 Go 团队的人更有可能继续使用 Go，或者可能反映了信任是如何随着时间推移而校准的。

![Chart of developer trust with the Go team](./GoDeveloperSurvey2024H1Results_img/trust_go.svg+xml)

### Community satisfaction 社区满意度

In the last year, almost a third of respondents (32%) said they participated in the Go developer community either online or at in-person events. More experienced Go developers were more likely to have participated in a community event and were more satisfied with community events overall. Although we can’t draw causal conclusions from this data, we did see a positive correlation between community satisfaction and overall satisfaction with Go. It could be that participating in the Go community increases satisfaction through increased social interaction or technical support. In general, we also found that respondents with less experience were less likely to have participated in events in the last year. This may mean they haven’t discovered events or found opportunities yet to be involved.

​	在去年，近三分之一的受访者（32%）表示他们在线或在面对面活动中参与了 Go 开发者社区。更有经验的 Go 开发者更有可能参与社区活动，并且对社区活动总体上更满意。尽管我们无法从这些数据中得出因果关系，但我们确实看到了社区满意度与对 Go 的总体满意度之间的正相关关系。参与 Go 社区可能会通过增加社交互动或技术支持来提高满意度。总体而言，我们还发现，经验较少的受访者在去年不太可能参与活动。这可能意味着他们尚未发现活动或找到参与的机会。

![Chart of participation in community events](./GoDeveloperSurvey2024H1Results_img/community_events.svg+xml) ![Chart of community satisfaction](./GoDeveloperSurvey2024H1Results_img/community_sat.svg+xml)

### Biggest challenges 最大的挑战

For several years, this survey has asked participants about their biggest challenge when using Go. This has always been in the form of an open text box and has elicited a wide variety of responses. In this cycle we introduced a closed form of the question, where we provided the most common write-in responses from prior years. Respondents were randomly shown either the open or closed forms of the question. The closed form helps us validate how we’ve historically interpreted these responses, while also increasing the number of Go developers we hear from: this year participants who saw the closed form were 2.5x more likely to answer than those who saw the open form. This higher number of responses narrows our margin of error and increases our confidence when interpreting survey results.

​	多年来，本调查一直询问参与者在使用 Go 时遇到的最大挑战。这始终以开放文本框的形式出现，并引发了各种各样的回应。在此周期中，我们引入了该问题的封闭形式，其中我们提供了前几年最常见的书面回复。受访者随机看到该问题的开放或封闭形式。封闭形式有助于我们验证我们过去如何解释这些回应，同时还可以增加我们听到的 Go 开发人员的数量：今年看到封闭形式的参与者比看到开放形式的参与者回答的可能性高 2.5 倍。如此多的回应缩小了我们的误差范围，并在解释调查结果时增加了我们的信心。

In the closed-form, only 8% of respondents selected “Other”, which suggests we captured the majority of common challenges with our response choices. Interestingly, 13% of respondents said they don’t face any challenges using Go. In the open text version of this question, only 2% of respondents gave this response. The top responses in the closed-form were learning how to write Go effectively (15%) and the verbosity of error handling (13%). This matches what we saw in the open-text form, where 11% of responses mentioned learning Go, learning best practices, or issues with documentation as their biggest challenge, and another 11% mentioned error handling.

​	在封闭式问题中，只有 8% 的受访者选择了“其他”，这表明我们通过我们的回答选项捕捉到了大多数常见挑战。有趣的是，13% 的受访者表示他们使用 Go 时没有遇到任何挑战。在这个问题的开放式文本版本中，只有 2% 的受访者给出了这个回答。封闭式问题中的主要回答是学习如何有效地编写 Go（15%）和错误处理的冗长（13%）。这与我们在开放式文本中看到的情况相符，其中 11% 的回答提到学习 Go、学习最佳实践或文档问题是他们遇到的最大挑战，而另有 11% 提到了错误处理。

![Chart of closed form biggest challenges using Go](./GoDeveloperSurvey2024H1Results_img/biggest_challenge_closed.svg+xml) ![Chart of open text biggest challenges using Go](./GoDeveloperSurvey2024H1Results_img/text_biggest_challenge.svg+xml)

Respondents who saw the closed form of the question also received a follow-up open-text question to give them an opportunity to tell us more about their biggest challenge in case they had wanted to provide more nuanced answers, additional challenges, or anything else they felt was important.The most common response mentioned Go’s type system, and often asked specifically for enums, option types, or sum types in Go. Often we did not get much context for these requests, but we suspect this is due to some recent proposals and community discussions related to enums, an increase in folks coming from other language ecosystems where these features are common, or the expectation that these features will reduce writing boilerplate code. One of the more comprehensive comments related to the type system explained as follows:

​	看到该问题的封闭式形式的受访者还收到一个后续开放式文本问题，以便有机会告诉我们更多有关他们遇到的最大挑战的信息，以防他们想要提供更多细致入微的答案、其他挑战或他们认为重要的任何其他信息。最常见的回答提到了 Go 的类型系统，并且经常具体要求 Go 中的枚举、选项类型或总和类型。通常我们没有为这些请求获得太多上下文，但我们怀疑这是由于与枚举相关的某些最新提议和社区讨论、来自这些功能很常见的其他语言生态系统的人数增加，或期望这些功能将减少编写样板代码。与类型系统相关的更全面的评论之一解释如下：

> “These aren’t big challenges, but more conveniences I miss in the language. There’s ways around all of them, but it would be nice not to have to think about it.
>
> ​	“这些都不是什么大问题，但都是我在语言中错过的便利之处。所有这些问题都有解决办法，但不必考虑这些问题会很好。

> Sum types/closed enums can be emulated but its a lot of faff. It’s a very handy feature to have when interacting with APIs that only have a limited set of values for a particular element/field in a response and a value outside of it is an error. It helps with validation and catching issues at the point of entry and can often directly be generated from API specifications like JSON Schema, OpenAPI or heaven forbid XML Schema Definitions.
>
> ​	和类型/封闭枚举可以模拟，但需要很多麻烦。当与仅对响应中的特定元素/字段有一组有限的值的 API 交互时，这是一个非常方便的功能，并且该值之外的值是一个错误。它有助于在入口点验证和捕获问题，并且通常可以直接从 API 规范（如 JSON Schema、OpenAPI 或禁止 XML Schema 定义）生成。

> I don’t mind the error checking verbosity at all, but the nil-checking with pointers gets tedious especially when [I] need to drill into a deeply nested struct of pointer fields. Some form of Optional/Result type or an ability to chase through a chain of pointers and simply get a nil back instead of triggering a runtime panic would be appreciated.”
>
> ​	我一点也不介意错误检查的冗长，但是使用指针进行 nil 检查会很乏味，尤其是在我需要深入嵌套的指针字段结构时。Optional/Result 类型或能够追溯指针链并简单地返回 nil 而不是触发运行时恐慌的某种形式将不胜感激。

![Chart of anything else related to biggest challenges using Go](./GoDeveloperSurvey2024H1Results_img/text_biggest_challenge_anything.svg+xml)

## Developer environments 开发人员环境

As in previous years, most survey respondents develop with Go on Linux (61%) and macOS (58%) systems. Although the numbers haven’t changed much from year to year, we did see some interesting differences in our self-selected sample. The randomly sampled groups from JetBrains and VS Code were more likely (31% and 33%, respectively) to develop on Windows than the self-selected group (19%). We don’t know exactly why the self-selected group is so different, but we hypothesize that, because they likely encountered the survey from reading the Go Blog, these respondents are some of the most engaged and experienced developers in the community. Their operating system preferences might be reflective of historical priorities of the core development team who typically developed on Linux and macOS. Thankfully we have the random samples from JetBrains and VS Code to provide a more representative view of developer preferences.

​	与往年一样，大多数调查受访者使用 Linux (61%) 和 macOS (58%) 系统开发 Go。尽管这些数字年复一年变化不大，但我们在自选样本中确实看到了一些有趣的差异。与自选组 (19%) 相比，JetBrains 和 VS Code 的随机抽样组更有可能 (分别为 31% 和 33%) 在 Windows 上开发。我们不知道自选组如此不同的确切原因，但我们假设，由于他们可能从阅读 Go 博客中遇到调查，因此这些受访者是社区中最活跃和最有经验的开发人员。他们的操作系统偏好可能反映了通常在 Linux 和 macOS 上开发的核心开发团队的历史优先级。值得庆幸的是，我们有来自 JetBrains 和 VS Code 的随机样本，可以提供更具代表性的开发人员偏好视图。

![Chart of operating systems respondents use when developing Go software](./GoDeveloperSurvey2024H1Results_img/os_dev.svg+xml) ![Chart of operating systems respondents use when developing Go software, split by difference sample sources](./GoDeveloperSurvey2024H1Results_img/os_dev_src.svg+xml) ![Chart of operating systems respondents use when developing Go software, split by duration of experience](./GoDeveloperSurvey2024H1Results_img/os_dev_exp.svg+xml)

As a followup for the 17% of respondents who develop on WSL, we asked which version they’re using. 93% of respondents who develop on WSL are using version 2, so going forward, [the Go team at Microsoft has decided to focus their efforts on WSL2.](https://go.dev/issue/63503)

​	对于在 WSL 上进行开发的 17% 的受访者，我们询问了他们正在使用哪个版本。在 WSL 上进行开发的受访者中有 93% 使用的是版本 2，因此，Microsoft 的 Go 团队决定将他们的精力集中在 WSL2 上。

![Chart of WSL versions usage](./GoDeveloperSurvey2024H1Results_img/wsl_version.svg+xml)

Given that two of our sample populations were recruited from within VS Code or GoLand, they are strongly biased towards preferring those editors. To avoid skewing the results, we show the data here from the self-selected group only. Similar to previous years, the most common code editors among Go Developer Survey respondents continue to be [VS Code](https://code.visualstudio.com/) (43%) and [GoLand](https://www.jetbrains.com/go/) (33%). We don’t see any statistically significant differences from mid-2023, (44% and 31%, respectively).

​	鉴于我们的两个样本人群是从 VS Code 或 GoLand 中招募的，因此他们强烈倾向于偏好这些编辑器。为了避免扭曲结果，我们仅在此处显示来自自选组的数据。与往年类似，Go 开发人员调查受访者中最常见的代码编辑器仍然是 VS Code（43%）和 GoLand（33%）。我们没有看到与 2023 年年中（分别为 44% 和 31%）有任何统计学上的显着差异。

![Chart of code editors respondents prefer to use with Go](./GoDeveloperSurvey2024H1Results_img/editor.svg+xml)

With the prevalence of Go for cloud development and containerized workloads, it’s no surprise that Go developers primarily deploy to Linux environments (93%). We didn’t see any significant changes from last year.

​	随着 Go 在云开发和容器化工作负载中的普及，Go 开发人员主要部署到 Linux 环境（93%）也就不足为奇了。我们没有看到去年有任何重大变化。

![Chart of platforms respondents deploy Go software to](./GoDeveloperSurvey2024H1Results_img/os_deploy.svg+xml)

Go is a popular language for modern cloud-based development, so we typically include survey questions to help us understand which cloud platforms Go developers are using and how satisfied they are with the three most popular platforms: Amazon Web Services (AWS), Microsoft Azure, and Google Cloud. This section was only shown to respondents who said they use Go for their primary job, about 76% of total respondents. 98% of those who saw this question work on Go software that integrates with cloud services. Over half of respondents used AWS (52%), while 27% used GCP for their Go development and deployments. For both AWS and Google Cloud, we don’t see any differences between small or large companies in their likelihood to use either provider. Microsoft Azure is the only cloud provider that is significantly more likely to be used in large organizations (companies with > 1,000 employees) than smaller shops. We didn’t see any significant differences in usage based on the size of the organization for any other cloud providers.

​	Go 是一种流行的现代云端开发语言，因此我们通常会包含调查问题，以帮助我们了解 Go 开发人员正在使用哪些云平台，以及他们对三大最流行平台的满意度：亚马逊网络服务 (AWS)、Microsoft Azure 和 Google Cloud。此部分仅向那些表示他们将 Go 用于其主要工作的受访者展示，约占受访者总数的 76%。98% 看到此问题的人员使用与云服务集成的 Go 软件。超过一半的受访者使用 AWS（52%），而 27% 的受访者使用 GCP 进行 Go 开发和部署。对于 AWS 和 Google Cloud，我们没有看到小型或大型公司在使用任一提供商的可能性方面有任何差异。Microsoft Azure 是唯一一家在大型组织（员工人数 > 1,000 人的公司）中使用可能性明显高于小型商店的云提供商。我们没有看到任何其他云提供商在使用方面基于组织规模存在任何显著差异。

The rates of satisfaction for using Go with AWS and Google Cloud were both 77%. Historically these rates have been about the same. As in previous years, the satisfaction rate for Microsoft Azure was lower (57%).

​	使用 Go 与 AWS 和 Google Cloud 的满意度均为 77%。从历史上看，这些比率一直大致相同。与往年一样，Microsoft Azure 的满意度较低（57%）。

![Chart of cloud platforms respondents](./GoDeveloperSurvey2024H1Results_img/cloud_platform.svg+xml) ![Chart of satisfaction with Go on AWS in the last year](./GoDeveloperSurvey2024H1Results_img/cloud_sat_aws.svg+xml) ![Chart of satisfaction with using Go on Google Cloud in the last year](./GoDeveloperSurvey2024H1Results_img/cloud_sat_gcp.svg+xml) ![Chart of satisfaction with using Go on Microsoft Azure in the last year](./GoDeveloperSurvey2024H1Results_img/cloud_sat_azure.svg+xml)

## Resource and Security Priorities 资源和安全优先级

To help prioritize the Go team’s work, we wanted to understand the top resource cost and security concerns for teams using Go. About half of respondents using Go at work reported having at least one resource cost concern in the last year (52%). The engineering costs of writing and maintaining Go services was more common (28%) than concern for the costs of running Go services (10%) or both about equally (12%). We didn’t see any significant differences in resource concerns between small and large organizations. To address concerns about resource costs, the Go team is continuing to optimize Go and enhance profile-guided optimization (PGO).

​	为了帮助优先考虑 Go 团队的工作，我们希望了解使用 Go 的团队最关心的资源成本和安全问题。大约一半在工作中使用 Go 的受访者报告称，在过去一年中至少有一个资源成本问题 (52%)。编写和维护 Go 服务的工程成本比运行 Go 服务的成本 (10%) 或两者都差不多 (12%) 更常见 (28%)。我们没有看到小型和大型组织在资源问题上有什么显着差异。为了解决对资源成本的担忧，Go 团队正在继续优化 Go 并增强基于配置文件的优化 (PGO)。

![Chart of cost concerns respondents have had related to their Go usage in the last year](./GoDeveloperSurvey2024H1Results_img/cost_concern.svg+xml)

As for security priorities, we asked respondents to tell us up to three of their top concerns. Of those who did have security concerns, overall, the top concern was insecure coding practices (42%), followed by system misconfiguration (29%). Our main takeaway is that respondents are especially interested in tooling to help find and fix potential security issues while they’re writing code. This aligns with what we’ve learned from prior research into how developers find and address security vulnerabilities.

​	至于安全优先级，我们要求受访者告诉我们他们最关心的三个问题。对于那些确实有安全顾虑的人来说，总体而言，最关心的问题是不安全的编码实践（42%），其次是系统错误配置（29%）。我们的主要收获是，受访者特别感兴趣的是在编写代码时帮助查找和修复潜在安全问题的工具。这与我们从先前关于开发人员如何发现和解决安全漏洞的研究中了解到的内容是一致的。

![Chart of cost concerns respondents have had related Go usage in the last year](./GoDeveloperSurvey2024H1Results_img/security_concern.svg+xml)

### Performance Tooling 性能工具

Our goals for this section were to measure how respondents perceive the ease or difficulty of diagnosing performance issues and determine whether this task is more or less difficult depending on their editor or IDE usage. Specifically, we wanted to know if it’s more difficult to diagnose performance issues from the command line, and if we should invest in improving the integration of performance diagnostic tooling within VS Code to make this task easier. In our analyses, we show comparisons between respondents who prefer VS Code or GoLand to highlight what we learned about the experience of using VS Code compared to another common editor.

​	本节的目标是衡量受访者如何感知诊断性能问题容易还是困难，并确定此任务是否根据其编辑器或 IDE 使用情况或多或少困难。具体来说，我们想知道从命令行诊断性能问题是否更困难，以及我们是否应该投资于改进 VS Code 中性能诊断工具的集成以使此任务更轻松。在我们的分析中，我们展示了更喜欢 VS Code 或 GoLand 的受访者之间的比较，以突出我们了解到的使用 VS Code 与另一个常见编辑器相比的体验。

We first asked a general question about different kinds of tools and techniques respondents use with Go to have some points of comparison. We found that only 40% of respondents use tools to improve code performance or efficiency. We didn’t see any significant differences based on editor or IDE preference, that is, VS Code users and GoLand users were about equally likely to use tools to improve code performance or efficiency.

​	我们首先询问了一个关于不同类型的工具和技术的问题，受访者使用 Go 来进行一些比较。我们发现只有 40% 的受访者使用工具来提高代码性能或效率。我们没有看到基于编辑器或 IDE 偏好的任何显著差异，也就是说，VS Code 用户和 GoLand 用户使用工具来提高代码性能或效率的可能性大致相同。

![Chart of different techniques used for security, quality and performance](./GoDeveloperSurvey2024H1Results_img/dev_techniques.svg+xml)

Most respondents (73%) told us that identifying and addressing performance issues is at least moderately important. Again, we didn’t see any significant differences here between GoLand and VS Code users in how important they found diagnosing performance issues.

​	大多数受访者 (73%) 告诉我们，识别和解决性能问题至少是中等重要性的。同样，我们没有看到 GoLand 和 VS Code 用户之间在诊断性能问题的重要性方面有任何显着差异。

![Chart of the importance of identifying and addressing performance issues ](./GoDeveloperSurvey2024H1Results_img/perf_importance.svg+xml)

Overall, respondents did not find diagnosing performance issues easy, with 30% reporting it was somewhat or very difficult and 46% saying it was neither easy nor difficult. Contrary to our hypothesis, VS Code users were not more likely to report challenges when diagnosing performance issues vs. other respondents. Those using their command line for diagnosing performance issues, regardless of their preferred editor, also did not report this task as more challenging than those using their IDE. Years of experience was the only significant factor we observed, where less experienced Go developers found it overall more difficult to diagnose performance issues than more experienced Go developers.

​	总体而言，受访者并未发现诊断性能问题很容易，30% 的人报告说这有点或非常困难，46% 的人说既不简单也不困难。与我们的假设相反，与其他受访者相比，VS Code 用户在诊断性能问题时不太可能报告挑战。无论首选编辑器如何，那些使用命令行来诊断性能问题的人也没有报告说这项任务比那些使用其 IDE 的人更具挑战性。我们观察到的唯一重要因素是经验年限，经验较少的 Go 开发人员总体上发现诊断性能问题比经验较丰富的 Go 开发人员更困难。

![Chart of how easy or difficult respondents found diagnosing performance issues ](./GoDeveloperSurvey2024H1Results_img/perf_easiness.svg+xml) ![Chart of how easy or difficult respondents found diagnosing performance issues split by duration of experience](./GoDeveloperSurvey2024H1Results_img/perf_easiness_exp.svg+xml) ![Chart of how easy or difficult respondents found diagnosing performance issues split by where they use performance diagnostic tools ](./GoDeveloperSurvey2024H1Results_img/perf_easiness_where.svg+xml)

To answer our original question, most developers found it difficult to diagnose performance issues in Go, regardless of their preferred editor or tooling. This was especially true for developers with less than two years of experience in Go.

​	为了回答我们的最初问题，大多数开发人员发现诊断 Go 中的性能问题很困难，无论他们偏好的编辑器或工具是什么。对于在 Go 中拥有不到两年经验的开发人员来说，这一点尤其正确。

We also included a follow-up for respondents who rated diagnosing performance issues as at least slightly important to understand which issues were most important to them. Latency, total memory, and total CPU were the top concerns. There could be several explanations to the significance of these areas. First, they are measurable and easily convertible into business costs. Secondly, total memory and CPU usage represent physical constraints that necessitate hardware upgrades or software optimizations for improvement. Moreover, latency, total memory, and total CPU are more manageable by developers and can impact even straightforward services. In contrast, GC performance and memory allocation may only be relevant in rare cases or for exceptionally heavy workloads. Additionally, latency stands out as the most user-visible metric, as high latency results in slow services and dissatisfied users.

​	我们还为受访者纳入了后续问题，他们将诊断性能问题评级为至少略微重要，以了解哪些问题对他们来说最重要。延迟、总内存和总 CPU 是最受关注的问题。这些领域的重要性可能有几种解释。首先，它们是可衡量的，并且很容易转换为业务成本。其次，总内存和 CPU 使用率代表了需要硬件升级或软件优化才能改进的物理限制。此外，延迟、总内存和总 CPU 更易于开发者管理，并且甚至可以影响直接的服务。相比之下，GC 性能和内存分配可能只在极少数情况下或对于特别繁重的负载才相关。此外，延迟作为最直观的用户可见指标脱颖而出，因为高延迟会导致服务缓慢和用户不满意。

![Chart of which performance issues are the highest concern to respondents](./GoDeveloperSurvey2024H1Results_img/perf_concerns.svg+xml)

## Understanding AI use cases for Go 了解 Go 的 AI 用例

Our [previous survey](https://go.dev/blog/survey2023-h2-results#mlai) asked Go developers about their early experiences with generative AI systems. To go a bit deeper this cycle, we asked several AI-related questions to understand how respondents are building AI-powered (more specifically, LLM-powered) services. We found that half of survey respondents (50%) work at organizations that are building or exploring AI-powered services. Of these, just over half (56%) said they were involved with adding AI capabilities to their organization’s services. Our remaining AI-related questions were only shown to this slice of respondents.

​	我们之前的调查询问了 Go 开发人员他们对生成式 AI 系统的早期体验。为了在这个周期中深入了解，我们询问了几个与 AI 相关的问卷，以了解受访者如何构建 AI 驱动的（更具体地说，是 LLM 驱动的）服务。我们发现，一半的受访者（50%）在构建或探索 AI 驱动的服务的组织中工作。在这些人中，略多于一半（56%）的人表示他们参与了向其组织的服务中添加 AI 功能。我们剩下的与 AI 相关的问卷仅向这一部分受访者展示。

Please be cautious about generalizing these participant responses to the overall population of Go developers. Because only about ¼ of survey respondents are working with AI-powered services, we suggest using this data to understand the early adopters in this space, with the caveat that early adopters tend to be a bit different than the majority of people who will eventually adopt a technology. As an example, we expect that this audience is experimenting with more models and SDKs than may be the case a year or two from now, and encountering more challenges related to integrating those services into their existing code base.

​	请谨慎将这些参与者的回复推广到 Go 开发人员的总体群体。由于只有大约 1/4 的调查受访者正在使用 AI 驱动的服务，我们建议使用此数据来了解该领域的早期采用者，但需要注意的是，早期采用者往往与最终采用该技术的多数人有些不同。例如，我们预计，与一两年后相比，该受众正在尝试更多模型和 SDK，并且在将这些服务集成到其现有代码库中时会遇到更多挑战。

![Chart of respondents whose org is currently building or exploring ML/AI based services](./GoDeveloperSurvey2024H1Results_img/ai_org.svg+xml) ![Chart of respondents who are currently involved in their orgs AI based development](./GoDeveloperSurvey2024H1Results_img/ai_involved.svg+xml)

Among the audience of Go developers working professionally with generative AI (GenAI) systems, a solid majority (81%) reported using OpenAI’s ChatGPT or DALL-E models. A collection of open-source models also saw high adoption, with a majority of respondents (53%) using at least one of Llama, Mistral, or another OSS model. We see some early evidence that larger organizations (1,000+ employees) are a bit less likely to be using OpenAI models (74% vs. 83%) and a bit more likely to be using other proprietary models (22% vs. 11%). We do not, however, see any evidence of differences in adoption of OSS models based on organization size–both smaller companies and larger enterprises show small majorities adopting OSS models (51% and 53%, respectively). Overall we found that a plurality of respondents prefer to use open-source models (47%) with only 19% preferring proprietary models; 37% said they had no preference.

​	在使用生成式 AI (GenAI) 系统进行专业工作的 Go 开发人员受众中，绝大多数 (81%) 报告使用 OpenAI 的 ChatGPT 或 DALL-E 模型。一系列开源模型也获得了高采用率，大多数受访者 (53%) 使用至少一个 Llama、Mistral 或其他 OSS 模型。我们看到一些早期证据表明，较大的组织（1,000 名以上员工）使用 OpenAI 模型的可能性稍低（74% 对比 83%），而使用其他专有模型的可能性稍高（22% 对比 11%）。然而，我们没有看到任何证据表明基于组织规模采用 OSS 模型存在差异——小型公司和大型企业都显示出采用 OSS 模型的小多数（分别为 51% 和 53%）。总体而言，我们发现大多数受访者更喜欢使用开源模型（47%），只有 19% 的人更喜欢专有模型；37% 的人表示他们没有偏好。

![Chart of which generative AI models respondents' orgs are using](./GoDeveloperSurvey2024H1Results_img/generative_models.svg+xml) ![Chart of which AI related services and libraries respondents' orgs are using](./GoDeveloperSurvey2024H1Results_img/ai_libs.svg+xml)

The most common kinds of services respondents are building include summarization tools (56%), text generation tools (55%), and chatbots (46%). Open-text responses suggested that many of these use cases are internal-facing, such as chat bots trained upon an organization’s internal documentation and intended to answer employee questions. Respondents raised several concerns about external-facing AI features, most notably due to reliability (e.g., do slight changes in my question lead to very different results?) and accuracy (e.g., are the results trustworthy?) issues. An interesting theme running through these responses was a sense of tension between the risk of not adopting AI tooling at all (and thereby losing a potential competitive advantage should generative AI become necessary in the future), balanced against the risk of negative publicity or violating regulations/laws by using untested AI in high-criticality customer-facing domains.

​	受访者构建的最常见服务类型包括摘要工具（56%）、文本生成工具（55%）和聊天机器人（46%）。开放式文本回复表明，其中许多用例都是面向内部的，例如针对组织内部文档进行训练并旨在回答员工问题的聊天机器人。受访者对面向外部的 AI 功能提出了若干担忧，最值得注意的是由于可靠性（例如，我的问题中轻微的更改会导致截然不同的结果吗？）和准确性（例如，结果是否可信？）问题。贯穿这些回复的一个有趣主题是对根本不采用 AI 工具的风险（从而在未来生成式 AI 变得必要时失去潜在的竞争优势）与通过在高度关键的客户面临领域使用未经测试的 AI 而产生负面宣传或违反法规/法律的风险之间的紧张感。

We found evidence that Go is already being used in the GenAI space, and there appears to be an appetite for more. Roughly ⅓ of respondents who were building AI-powered features told us they were already using Go for a variety of GenAI tasks, including prototyping new features and integrating services with LLMs. These proportions tick up slightly for two areas where we believe Go is a particularly well-suited tool: data pipelines for ML/AI systems (37%) and hosting API endpoints for ML/AI models (41%). In addition to these (likely early) adopters, we found that about ¼ of respondents *want* to use Go for these types of uses, but are currently blocked by something. We’ll return to these blockers shortly, after exploring why respondents wanted to use Go for these tasks in the first place.

​	我们发现证据表明 Go 已在 GenAI 领域中使用，并且似乎有更多需求。大约 ⅓ 正在构建 AI 驱动的功能的受访者告诉我们，他们已经在使用 Go 来执行各种 GenAI 任务，包括对新功能进行原型设计以及将服务与 LLMs 集成。对于我们认为 Go 是特别合适的工具的两个领域，这些比例略有上升：ML/AI 系统的数据管道 (37%) 和托管 ML/AI 模型的 API 端点 (41%)。除了这些（可能是早期）采用者之外，我们发现大约 ¼ 的受访者希望将 Go 用于这些类型的用途，但目前受到某些因素的阻碍。在探讨受访者最初希望将 Go 用于这些任务的原因之后，我们将很快回到这些阻碍因素。

![Chart of the kinds of Generative AI apps respondents work on](./GoDeveloperSurvey2024H1Results_img/ai_apps.svg+xml) ![Chart of the kinds of AI apps respondents' orgs are currently working on or considering](./GoDeveloperSurvey2024H1Results_img/ai_uses_interest.svg+xml)

### Reasons for using Go with generative AI systems 使用 Go 与生成式 AI 系统的原因

To help us understand what benefits developers hope to derive from using Go in their AI/ML services, we asked developers why they feel Go is a good choice for this domain. A clear majority (61%) of respondents mentioned one or more of Go’s core principles or features, such as simplicity, runtime safety, concurrency, or single-binary deployments. One third of respondents cited existing familiarity with Go, including a desire to avoid introducing new languages if they can avoid it. Rounding out the most common responses were various challenges with Python (particularly for running production services) at 14%.

​	为了帮助我们了解开发者希望从在 AI/ML 服务中使用 Go 中获得哪些好处，我们询问了开发者他们为何觉得 Go 是该领域的不错选择。绝大多数 (61%) 受访者提到了 Go 的一项或多项核心原则或特性，例如简单性、运行时安全性、并发性或单二进制部署。三分之一的受访者提到了对 Go 的现有熟悉度，包括如果可以避免的话，希望避免引入新语言。最常见的回答中，有 14% 提到了 Python 的各种挑战（尤其是对于运行生产服务）。

> “I think that the robustness, simplicity, performance and native binaries that the language offers make it a far stronger choice for AI workloads.” *— Open-source Go developer at a large organization with up to 1 year of experience*
>
> ​	“我认为该语言提供的健壮性、简单性、性能和原生二进制文件使其成为 AI 工作负载的更强大选择。”——在大型组织中拥有长达 1 年经验的开源 Go 开发人员

> “We want to keep our tech stack as homogenous as possible across the organization to make it easier for everybody to develop on all areas. Since we are already writing all our backends in Go, it is of interest to us to be able to write ML model deployments in Go and avoid having to rewrite parts of the stack for logging, monitoring, etc… in a separate language [like] Python.” *— Professional Go developer at a mid-sized organization with 5 – 7 years of experience*
>
> ​	“我们希望在整个组织中尽可能保持我们的技术栈同质化，以便每个人都能更轻松地在所有领域进行开发。由于我们已经用 Go 编写了所有后端，因此我们有兴趣能够用 Go 编写 ML 模型部署，并避免不得不为日志记录、监控等重写堆栈的某些部分，用另一种语言 [如] Python。”——在一家拥有 5-7 年经验的中型组织中担任专业 Go 开发人员

> “Go is better for us at running API servers and background tasks on worker pools. Go’s lower resource usage has allowed us to grow without using more resources. And we have found that Go projects are easier to maintain over time both in code changes and when updating dependencies. We run the models as a separate service written in Python and interact with them in Go.” *— Professional Go developer at a large organization with 5 – 7 years of experience*
>
> ​	“Go 更适合我们运行 API 服务器和工作池上的后台任务。Go 较低的资源使用率让我们得以在不使用更多资源的情况下实现增长。我们发现，Go 项目在代码更改和更新依赖项时都更容易维护。我们以 Python 编写的模型作为一项独立服务运行，并通过 Go 与它们进行交互。”——在一家拥有 5-7 年经验的大型组织中担任专业 Go 开发人员

It appears that among Go developers who are interested in ML/AI, there is a shared sense that 1) Go is inherently a good language for this domain (for the reasons articulated above), and 2) there is reluctance to introduce a new language once organizations have already invested in Go (this point reasonably generalizes to any language). Some respondents also expressed frustration with Python for reasons such as type safety, code quality, and challenging deployments.

​	在对 ML/AI 感兴趣的 Go 开发者中，似乎存在一种共识，即 1) Go 本质上是适合此领域的良好语言（出于上述原因），并且 2) 一旦组织已经投资于 Go，就不愿意引入新语言（这一点合理地概括为任何语言）。一些受访者还对 Python 表示失望，原因包括类型安全、代码质量和部署挑战。

![Chart of respondents' reasons for why Go is a good choice for their AI related use case](./GoDeveloperSurvey2024H1Results_img/text_ml_interest.svg+xml)

### Challenges when using Go with GenAI systems 使用 Go 与 GenAI 系统时遇到的挑战

Respondents were largely unified on what currently prevents them from using Go with AI-powered services: the ecosystem is centered around Python, their favorite libraries/frameworks are all in Python, getting started documentation assumes Python familiarity, and the data scientists or researchers exploring these models are already familiar with Python.

​	受访者在目前阻碍他们将 Go 与 AI 驱动的服务结合使用的问题上基本达成一致：生态系统以 Python 为中心，他们最喜欢的库/框架都在 Python 中，入门文档假定熟悉 Python，而探索这些模型的数据科学家或研究人员已经熟悉 Python。

> “Python just seems to have all the libraries. PyTorch for example is widely used to run models. If there were frameworks in Go to run these models, we’d much rather be doing that.” *— Professional Go developer at a large organization with 2 – 4 years of experience*
>
> ​	“Python 似乎拥有所有库。例如，PyTorch 广泛用于运行模型。如果 Go 中有用于运行这些模型的框架，我们更愿意这样做。”——在一家拥有 2-4 年经验的大型组织中担任专业 Go 开发人员

> “Python tools are substantially more mature and usable out of the box, making them a significantly lower cost to implement.” *— Professional Go developer at a small organization with 2 – 4 years of experience*
>
> ​	“Python 工具开箱即用，成熟度和可用性都更高，因此实施成本也显著降低。”——在一家小型组织工作 2-4 年的专业 Go 开发人员

> “[The] Go world is missing many AI libraries. If I have a LLM PyTorch model, I can’t even serve it (or I’m unaware how to do it). With Python it’s basically a few lines of code.” *— Professional Go developer at a small organization with up to 1 year of experience*
>
> ​	[Go]世界缺少许多 AI 库。如果我有一个 LLM PyTorch 模型，我甚至无法提供它（或者我不知道如何提供它）。使用 Python，它基本上是几行代码。——在一家拥有最多 1 年经验的小型组织中担任Go专业开发人员

These findings triangulate well with our observation above that Go developers believe Go *should* be a great language for building production-ready AI services: only 3% of respondents said that something specific to Go was blocking their path forward, and only 2% cited specific interoperability challenges with Python. In other words, most blockers developers face could be resolved in the module and documentation ecosystem, rather than necessitating core language or runtime changes.

​	这些调查结果与我们上述的观察结果很好地吻合，即 Go 开发者认为 Go 应该是构建可用于生产环境的 AI 服务的出色语言：只有 3% 的受访者表示 Go 中的某些特定内容阻碍了他们的前进道路，只有 2% 的受访者提到了与 Python 的特定互操作性挑战。换句话说，大多数阻碍开发者的问题都可以通过模块和文档生态系统解决，而不是需要核心语言或运行时更改。

![Chart of what is blocking respondents from using Go with their AI powered apps](./GoDeveloperSurvey2024H1Results_img/text_ml_blockers.svg+xml)

We also asked survey participants whether they were already working with Python for GenAI, and if so, whether they’d prefer to use Go. Respondents who said they’d prefer to use Go rather than Python also received a follow-up about what would enable them to use Go with GenAI systems.

​	我们还询问了调查参与者他们是否已经使用 Python 进行 GenAI，如果是，他们是否更愿意使用 Go。那些表示他们更愿意使用 Go 而不是 Python 的受访者还收到了有关如何使用 Go 与 GenAI 系统的后续问题。

A solid majority (62%) of respondents reported already using Python to integrate with generative AI models; of this group, 57% would rather use Go instead. Given that our survey audience are all Go developers, we should expect this to be an approximate upper bound on the proportion of overall developers who are interested in moving from Python to Go for GenAI tasks, given the state of each ecosystem today.

​	62% 的受访者报告称已经使用 Python 与生成式 AI 模型集成；在这一群体中，57% 的人更愿意使用 Go。鉴于我们的调查受众都是 Go 开发人员，我们应该期望这成为对希望从 Python 转向 Go 以执行 GenAI 任务的整体开发人员比例的近似上限，因为这是当今每个生态系统的发展状态。

Of the respondents who are already using Python but would prefer to use Go, the vast majority (92%) said that the availability of Go equivalents for Python libraries would enable them to integrate Go with GenAI systems. However, we should be cautious when interpreting this result; the open-text responses and a separate set of contextual interviews with developers working on GenAI services describe a Python-centric ecosystem around GenAI; it’s not only that Go lacks many libraries when compared with the Python ecosystem, but also that the perceived level of investment into Go libraries is lower, documentation and examples are predominantly in Python, and the network of experts working in this area are already comfortable with Python. Experimenting and building proofs-of-concept in Python is almost certain to continue, and the lack of Go variants of Python libraries (for example, [pandas](https://pandas.pydata.org/)) is only the first barrier developers would encounter when trying to port from Python to Go. Libraries and SDKs are necessary, but unlikely by themselves to be sufficient, to build a robust Go ecosystem for production ML/AI applications.

​	在已经使用 Python 但更愿意使用 Go 的受访者中，绝大多数（92%）表示，如果 Go 提供与 Python 库等效的库，他们就能将 Go 与 GenAI 系统集成。但是，在解读这一结果时我们应保持谨慎；开放文本回复和对从事 GenAI 服务的开发人员进行的另一组情境访谈描述了围绕 GenAI 的以 Python 为中心的生态系统；与 Python 生态系统相比，Go 不仅缺少许多库，而且 Go 库的投资水平被认为较低，文档和示例主要使用 Python，而且在这个领域工作的专家网络已经习惯了 Python。几乎可以肯定的是，在 Python 中进行实验和构建概念验证将继续进行，而缺少 Python 库（例如 pandas）的 Go 变体只是开发人员在尝试从 Python 移植到 Go 时遇到的第一个障碍。库和 SDK 是必需的，但仅靠它们自己不太可能足以构建一个用于生产 ML/AI 应用程序的强大的 Go 生态系统。

Further, contextual interviews with Go developers building AI-powered services suggest that *calling* APIs from Go is not a major issue, particularly with hosted models such as [GPT-4](https://openai.com/gpt-4) or [Gemini](https://gemini.google.com/). Building, evaluating, and hosting custom models is seen as challenging in Go (primarily due to the lack of frameworks and libraries that support this in Python), but interview participants distinguished between hobbyist use cases (e.g., playing around with custom models at home) and business use cases. The hobbyist cases are dominated by Python for all of the reasons enumerated above, but the business use cases are more focused around reliability, accuracy, and performance while calling hosted models. This is an area where Go can shine *without* building a large ecosystem of ML/AI/data science libraries, though we expect developers will still benefit from documentation, best practice guidance, and examples.

​	此外，与构建 AI 驱动的服务的 Go 开发人员进行的背景访谈表明，从 Go 调用 API 并不是一个主要问题，特别是对于托管模型（例如 GPT-4 或 Gemini）。在 Go 中构建、评估和托管自定义模型被视为一项挑战（主要是因为缺少 Python 中支持此功能的框架和库），但访谈参与者区分了业余爱好用例（例如，在家中使用自定义模型）和业务用例。业余爱好用例因上述所有原因而以 Python 为主，但业务用例更注重在调用托管模型时的可靠性、准确性和性能。这是一个 Go 可以大放异彩的领域，而无需构建一个大型的 ML/AI/数据科学库生态系统，尽管我们预计开发人员仍将受益于文档、最佳实践指南和示例。

Because the field of GenAI is so novel, best practices are still being identified and tested. Initial contextual interviews with developers have suggested that one of their goals is to be prepared for a future in which GenAI becomes a competitive advantage; by making some investment in this area now, they hope to moderate future risk. They’re also still trying to understand what GenAI systems might be helpful for and what the return on investment (if any) may look like. Due to these unknowns, our early data suggests that organizations (especially outside the tech industry) may be hesitant to make long-term commitments here, and will instead pursue a lean or scrappy approach until either a reliable use case with clear benefits emerges, or their industry peers begin to make large, public investments in this space.

​	由于生成式人工智能领域非常新颖，最佳实践仍在识别和测试中。与开发人员进行的初步背景访谈表明，他们的目标之一是为生成式人工智能成为竞争优势的未来做好准备；通过现在对该领域进行一些投资，他们希望缓和未来的风险。他们还仍在尝试了解生成式人工智能系统可能对哪些方面有帮助，以及投资回报（如果有的话）可能是什么样子。由于这些未知因素，我们的早期数据表明，组织（尤其是在科技行业之外）可能不愿在此做出长期承诺，而是在出现具有明显优势的可靠用例或其行业同行开始在此领域进行大规模公开投资之前，采用精益或拼凑式方法。

![Chart showing high usage of Python to integrate with gen AI models](./GoDeveloperSurvey2024H1Results_img/python_usage.svg+xml) ![Chart showing preference to use Go rather than Python to integrate with gen AI models](./GoDeveloperSurvey2024H1Results_img/go_python_pref.svg+xml) ![Chart of what would enable respondents to use Go where they are currently using Python](./GoDeveloperSurvey2024H1Results_img/enable_go.svg+xml) ![Chart of biggest challenges for respondents integrating backend services with gen AI models](./GoDeveloperSurvey2024H1Results_img/text_ml_challenge.svg+xml)

## Learning challenges 学习挑战

In order to improve the experience of learning Go, we wanted to hear from inexperienced Go developers, as well as those who might have already mastered the basics on what they see as their biggest challenge to meeting their learning goals. We also wanted to hear from developers who might primarily be focused on helping others get started with Go rather than their own learning goals, since they might have some insights on common challenges they see when onboarding developers.

​	为了改善学习 Go 的体验，我们希望听取没有经验的 Go 开发人员以及那些可能已经掌握了基础知识的人对他们认为实现学习目标的最大挑战的看法。我们还希望听取主要专注于帮助他人开始使用 Go 而不是他们自己的学习目标的开发人员的意见，因为他们可能对在开发人员入职时遇到的常见挑战有一些见解。

Only 3% of respondents said that they were currently learning the basics of Go. This isn’t too surprising, considering most of our survey respondents have at least a year of experience with Go. Meanwhile, 40% of respondents said that they have already learned the basics but want to learn more advanced topics and another 40% said that they help other developers learn Go. Only 15% said they didn’t have any learning goals related to Go.

​	只有 3% 的受访者表示他们目前正在学习 Go 的基础知识。这并不令人惊讶，因为我们调查的大多数受访者都有至少一年的 Go 经验。与此同时，40% 的受访者表示他们已经学习了基础知识，但希望学习更多高级主题，另有 40% 的受访者表示他们帮助其他开发人员学习 Go。只有 15% 的人表示他们没有任何与 Go 相关的学习目标。

![Chart of respondents' learning goals for Go](./GoDeveloperSurvey2024H1Results_img/learning_goal.svg+xml)

When we looked at more finely grained time segments of Go experience, we found that 30% of those who’ve been using Go for less than three months say they’re learning the basics of Go, while about two-thirds of them say that they’ve already learned the basics. That’s good evidence that someone can at least feel like they’ve learned the basics of Go in a short amount of time, but it also means we don’t have as much feedback from this group who are at the beginning of their learning journey.

​	当我们查看 Go 体验的更细粒度的时间段时，我们发现 30% 使用 Go 不到三个月的人表示他们正在学习 Go 的基础知识，而大约三分之二的人表示他们已经学完了基础知识。这很好地证明了人们至少可以在短时间内感觉自己已经学完了 Go 的基础知识，但这也意味着我们没有从这群刚开始学习的人那里获得太多反馈。

![Chart of respondents' learning goals for Go split by finer units of time](./GoDeveloperSurvey2024H1Results_img/learning_goal_go_exp.svg+xml)

To determine what kinds of learning materials might be most needed in the community, we asked what kind of learning content respondents preferred for topics related to software development. They were able to select multiple options so the numbers here exceed 100%. 87% of respondents said they preferred written content, which was by far the most preferred format. 52% said they preferred video content, and in particular this format was more often preferred by developers with less experience. This could indicate a growing desire for learning content in video format. The less experienced demographic did not prefer written content any less than other groups, however. [Providing both written and video formats together has been shown to improve learning outcomes](https://www.sciencedirect.com/science/article/abs/pii/S0360131514001353) and [helps developers with different learning preferences and abilities](https://udlguidelines.cast.org/representation/perception), which could increase the accessibility of learning content in the Go community.

​	为了确定社区中可能最需要哪种学习资料，我们询问了受访者在与软件开发相关的主题中更喜欢哪种学习内容。他们可以选择多个选项，因此这里的数字超过了 100%。87% 的受访者表示他们更喜欢书面内容，这是迄今为止最受欢迎的格式。52% 的人表示他们更喜欢视频内容，尤其是经验较少的开发人员更喜欢这种格式。这可能表明对视频格式的学习内容的需求不断增长。然而，经验较少的群体并不比其他群体更不喜欢书面内容。事实证明，同时提供书面和视频格式可以改善学习成果，并帮助具有不同学习偏好和能力的开发人员，这可以增加 Go 社区中学习内容的可访问性。

![Chart of respondents' preferred formats for learning content, split by years of Go experience](./GoDeveloperSurvey2024H1Results_img/learning_content_exp.svg+xml)

We asked respondents who said they had a learning goal related to Go what their biggest challenge was to reaching their goal. This was intentionally left broad enough that someone who was just getting started or who had already mastered the basics could respond to this question. We also wanted to give respondents the opportunity to tell us about a wide range of challenges, not just topics they find difficult.

​	我们询问了那些表示自己有与Go相关的学习目标的受访者，他们实现目标的最大挑战是什么。这个问题故意留得足够宽泛，以便刚入门或已经掌握基础知识的人都可以回答这个问题。我们还希望给受访者一个机会，让他们告诉我们各种挑战，而不仅仅是他们觉得困难的话题。

Overwhelmingly, the most common challenge mentioned was a lack of time or other personal limitations such as focus or motivation to learn or (44%). Although we can’t give respondents more time, we should be mindful when we’re producing learning materials or introducing changes in the ecosystem that users may be operating under significant time constraints. There may also be opportunities for educators to produce resources that are [digestible in smaller portions](https://web.cortland.edu/frieda/id/IDtheories/26.html) or [at a regular cadence](https://psychology.ucsd.edu/undergraduate-program/undergraduate-resources/academic-writing-resources/effective-studying/spaced-practice.html#:~:text=This is known as spaced,information and retain it longer.) to keep learners motivated.

​	绝大多数情况下，最常见的问题是缺乏时间或其他个人限制，例如专注力或学习动力（44%）。尽管我们无法给受访者更多时间，但我们在制作学习材料或引入生态系统变化时应注意，用户可能面临着严重的时间限制。教育工作者还可以制作一些资源，这些资源可以分小部分消化，或以规律的节奏来保持学习者的动力。

Other than time, the top challenge was learning new concepts, idioms or best practices that are unique to Go (11%). In particular, adapting to a statically typed compiled language from Python or JavaScript and learning how to organize Go code can be particularly challenging. Respondents also asked for more examples (6%), both in documentation and real world applications to learn from. Developers coming from a larger developer community expected to be able to find more existing solutions and examples.

​	除了时间，最大的挑战是学习 Go 特有的新概念、惯用语或最佳实践（11%）。特别是，从 Python 或 JavaScript 适应静态类型编译语言，以及学习如何组织 Go 代码可能特别具有挑战性。受访者还要求提供更多示例（6%），包括文档和现实世界中的应用程序以供学习。来自较大开发者社区的开发者希望能够找到更多现有的解决方案和示例。

> “Moving from a language like Python to a statically typed, compiled language has been challenging, but Go itself hasn’t been. I like to learn through quick feedback, so Python’s REPL was great for that. So now I need to focus on really reading documentation and examples to be able to learn. Some of the documentation for Go is quite sparse and could do with more examples.” *— Respondent with less than 3 years of experience with Go.*
>
> ​	“从 Python 这样的语言转向静态类型、编译语言一直很有挑战性，但 Go 本身并没有。我喜欢通过快速反馈来学习，所以 Python 的 REPL 对此非常棒。因此，我现在需要专注于真正阅读文档和示例才能学习。Go 的一些文档相当稀疏，可以增加更多示例。”——Go 经验不足 3 年的受访者。

> “My main challenge is the lack of example projects for enterprise-level applications. How to organize a big Go project is something I would like to have more examples as reference. I would like to refactor the current project I am working [on] to a more modular/clean architecture style, and I find it difficult in Go due to lack of examples / a more opinionated ‘folder/package’ reference.” *— Respondent with 1–2 years of experience with Go.*
>
> ​	“我的主要挑战是缺乏针对企业级应用程序的示例项目。如何组织一个大型 Go 项目是我希望有更多示例作为参考的内容。我想将我正在处理的当前项目重构为更模块化/干净的架构风格，由于缺乏示例/更主观的“文件夹/包”参考，我发现用 Go 很难做到这一点。”——具有 1-2 年 Go 经验的受访者。

> “It’s a smaller ecosystem than I am used to so online searches don’t yield as many results to specific issues. The resources that are out there are incredibly helpful and I usually am able to solve issues eventually, it just takes a little longer."*— Respondent with less than 3 months of experience with Go.*
>
> ​	“这是一个比我习惯的更小的生态系统，因此在线搜索不会产生针对特定问题的那么多结果。现有的资源非常有帮助，我通常最终能够解决问题，只是需要花费更长的时间。”——使用 Go 不到 3 个月的受访者。

![Chart of biggest challenges to reaching respondents' learning goals](./GoDeveloperSurvey2024H1Results_img/text_learning_challenge.svg+xml)

For respondents whose primary learning goal was to help others get started with Go, we asked what might make it easier for developers to get started with Go. We got a wide range of responses including documentation suggestions, comments on difficult topics (e.g., using pointers or concurrency), as well as requests for adding more familiar features from other languages. For categories that made up less than 2% of responses, we lumped them into “Other” responses. Interestingly, nobody mentioned “more time.” We think this is because lack of time or motivation is most often a challenge when there isn’t an immediate necessity to learn something new related to Go. For those helping others get started with Go, there may be a business reason for doing so, making it easier to prioritize, and hence “lack of time” is not as much of a challenge.

​	对于主要学习目标是帮助他人开始使用 Go 的受访者，我们询问了哪些因素可以使开发人员更容易开始使用 Go。我们得到了广泛的答复，包括文档建议、对困难主题的评论（例如，使用指针或并发），以及要求添加其他语言中更熟悉的特性。对于少于 2% 答复的类别，我们将其归为“其他”答复。有趣的是，没有人提到“更多时间”。我们认为这是因为当没有立即必要学习与 Go 相关的新知识时，时间或动机不足通常是一个挑战。对于那些帮助他人开始使用 Go 的人来说，可能存在一个商业原因，使其更容易确定优先级，因此“时间不足”并不是一个很大的挑战。

Consistent with the previous results, 16% of those who help others get started with Go told us that new Go developers would benefit from having more realistic examples or project-based exercises to learn from. They also saw the need to help developers coming from other language ecosystems through comparisons between them. [Previous research tells us that experience with one programming language can interfere with learning a new one](https://dl.acm.org/doi/abs/10.1145/3377811.3380352), especially when new concepts and tooling are different from what developers are used to. There are existing resources that aim to address this issue (just try searching for “Golang for [language] developers” for examples), but it could be difficult for new Go developers to search for concepts they don’t have the vocabulary for yet or these kinds of resources might not adequately address specific tasks. In the future we would like to learn more about how and when to present language comparisons to facilitate learning new concepts.

​	与之前的结果一致，16% 的帮助他人开始使用 Go 的人告诉我们，新的 Go 开发人员将受益于有更多现实示例或基于项目的练习来学习。他们还看到了通过比较来帮助来自其他语言生态系统的开发人员的必要性。先前的研究告诉我们，一种编程语言的经验会干扰学习另一种语言，尤其是当新概念和工具与开发人员习惯的不同时。有现有的资源旨在解决这个问题（只需尝试搜索“Golang for [language] developers”即可获取示例），但对于新的 Go 开发人员来说，搜索他们还没有词汇的概念可能会很困难，或者这些类型的资源可能无法充分解决特定任务。在未来，我们希望更多地了解如何以及何时展示语言比较以促进学习新概念。

A related need that this group reported was more explanations behind Go’s philosophy and best practices. It could be the case that learning not only *what* makes Go different but also *why* would help new Go developers understand new concepts or ways of doing tasks that might be different from their previous experience.

​	这个小组报告的另一个相关需求是更多地解释 Go 的理念和最佳实践。这可能是这种情况，不仅学习 Go 的不同之处，而且学习原因将有助于新的 Go 开发人员理解新概念或执行任务的方式，这可能与他们之前的经验不同。

![Chart of ideas from respondents  who help others get started with Go](./GoDeveloperSurvey2024H1Results_img/text_onboard_others.svg+xml)

## Demographics 人口统计学

We ask similar demographic questions during each cycle of this survey so we can understand how comparable the year-over-year results may be. For example, if a majority of respondents reported having less than one year of experience with Go in one survey cycle, it’d be very likely that any other differences in results from prior cycles stem from this major demographic shift. We also use these questions to provide comparisons between groups, such as satisfaction according to how long respondents have been using Go.

​	在每次调查周期中，我们都会询问类似的人口统计问题，以便了解年复一年的结果的比较性。例如，如果大多数受访者报告说在一次调查周期中使用 Go 的经验不足一年，那么很可能先前周期中结果的任何其他差异都源于这一主要的人口统计变化。我们还使用这些问题来提供组之间的比较，例如根据受访者使用 Go 的时间长短来比较满意度。

This year we introduced some minor changes to how we ask about experience with Go to match the JetBrains developer survey. This allowed us to make comparisons between our survey populations and facilitated data analysis.

​	今年，我们对询问 Go 经验的方式进行了一些微小的更改，以匹配 JetBrains 开发人员调查。这使我们能够在我们的调查人群之间进行比较，并促进了数据分析。

![Chart of how long respondents have been working with Go](./GoDeveloperSurvey2024H1Results_img/go_exp.svg+xml)

We saw some differences in experience level depending on how developers discovered our survey. The population who responded to survey notifications in VS Code skewed toward less experience with Go; we suspect this a reflection of VS Code’s popularity with new Go developers, who may not be ready to invest in an IDE license while they’re still learning. With respect to years of Go experience, the respondents randomly selected from GoLand are more similar to our self-selected population who found the survey through the Go Blog. Seeing consistencies between samples such as these allows us to more confidently generalize findings to the rest of the community.

​	我们发现，开发人员发现我们调查的方式不同，经验水平也有所不同。在 VS Code 中响应调查通知的人群倾向于对 Go 的经验较少；我们怀疑这是 VS Code 在 Go 新手开发者中受欢迎程度的反映，他们可能还没有准备好投资 IDE 许可证，因为他们仍在学习。就 Go 经验年限而言，从 GoLand 中随机选出的受访者与通过 Go 博客找到调查的我们自选人群更相似。看到诸如此类的样本之间的一致性，使我们能够更自信地将调查结果推广到社区的其他人。

![Chart of how long respondents have been working with Go, split by different sample sources](./GoDeveloperSurvey2024H1Results_img/go_exp_src.svg+xml)

In addition to years of experience with Go, this year we also measured years of professional coding experience. We were surprised to find that 26% of respondents have 16 or more years of professional coding experience. For comparison, the [JetBrains Developer Survey audience](https://www.jetbrains.com/lp/devecosystem-2023/demographics/#code_yrs) from 2023 had a majority of respondents with 3–5 years of professional experience. Having a more experienced demographic could affect differences in responses. For example, we saw significant differences in what kinds of learning content respondents with different levels of experience preferred.

​	除了多年的 Go 经验外，今年我们还衡量了多年的专业编码经验。我们惊讶地发现，26% 的受访者拥有 16 年或以上的专业编码经验。相比之下，2023 年的 JetBrains 开发者调查受众中，大多数受访者拥有 3-5 年的专业经验。拥有更资深的受访者可能会影响回答的差异。例如，我们发现具有不同经验水平的受访者在偏好的学习内容类型方面存在显着差异。

![Chart of respondents' years of professional developer experience](./GoDeveloperSurvey2024H1Results_img/dev_exp.svg+xml)

When we looked at our different samples, the self-selected group was even more experienced than the randomly selected groups, with 29% having 16 or more years of professional experience. This suggests that our self-selected group is generally more experienced than our randomly selected groups and can help explain some of the differences we see in this group.

​	当我们审视不同的样本时，自选组比随机选取的组更有经验，29% 的人有 16 年或以上的专业经验。这表明我们的自选组通常比我们的随机选取的组更有经验，并且可以帮助解释我们在这个组中看到的某些差异。

![Chart of respondents' years of professional developer experience](./GoDeveloperSurvey2024H1Results_img/dev_exp_src.svg+xml)

We introduced another demographic question during this cycle on employment status to help us make comparisons with [JetBrains’ Developer Survey](https://www.jetbrains.com/lp/devecosystem-2023/demographics/#employment_status). We found that 81% of respondents were fully employed, significantly more than 63% on the JetBrains survey. We also found significantly fewer students in our population (4%) compared to 15% on the JetBrains survey. When we look at our individual samples, we see a small but significant difference within our respondents from VS Code, who are slightly less likely to be fully employed and slightly more likely to be students. This makes sense given that VS Code is free.

​	在本次周期中，我们引入了另一个关于就业状况的人口统计问题，以帮助我们与 JetBrains 的开发者调查进行比较。我们发现 81% 的受访者为全职员工，显著高于 JetBrains 调查中的 63%。我们还发现，与 JetBrains 调查中的 15% 相比，我们的人群中学生明显较少（4%）。当我们查看我们的个人样本时，我们看到来自 VS Code 的受访者中存在一个细微但显著的差异，他们不太可能成为全职员工，而更有可能成为学生。鉴于 VS Code 是免费的，这一点是合理的。

![Chart of respondents' employment status](./GoDeveloperSurvey2024H1Results_img/employment.svg+xml)

Similar to previous years, the most common use cases for Go were API/RPC services (74%) and command line tools (63%). We’ve heard that Go’s built-in HTTP server and concurrency primitives, ease of cross-compilation, and single-binary deployments make Go a good choice for these kinds of applications.

​	与往年类似，Go 最常见的用例是 API/RPC 服务 (74%) 和命令行工具 (63%)。我们听说 Go 的内置 HTTP 服务器和并发原语、易于交叉编译以及单一二进制部署使 Go 成为此类应用程序的理想选择。

We also looked for differences based on respondents’ level of experience with Go and organization size. More experienced Go developers reported building a wider variety of applications in Go. This trend was consistent across every category of app or service. We did not find any notable differences in what respondents are building based on their organization size.

​	我们还根据受访者对 Go 的经验水平和组织规模寻找差异。更有经验的 Go 开发者报告说使用 Go 构建了更多种类的应用程序。这种趋势在每个类别的应用程序或服务中都是一致的。我们没有发现受访者根据其组织规模构建的内容有任何显着差异。

![Chart of the types of things respondents are building with Go](./GoDeveloperSurvey2024H1Results_img/what.svg+xml)

## Firmographics 人口统计

We heard from respondents at a variety of different organizations. About 27% worked at large organizations with 1,000 or more employees, 25% were from midsize organizations of 100–1,000 employees, and 43% worked at smaller organizations with less than 100 employees. As in previous years, the most common industry people work in was technology (48%) while the second most common was financial services (13%) .

​	我们从不同组织的受访者那里得到了反馈。大约 27% 的受访者在拥有 1,000 名或更多员工的大型组织工作，25% 的受访者来自拥有 100-1,000 名员工的中型组织，43% 的受访者在拥有不到 100 名员工的小型组织工作。与往年一样，人们工作最多的行业是技术行业（48%），其次是金融服务行业（13%）。

This is statistically unchanged from the past few Go Developer Surveys—we continue to hear from people in different countries and in organizations of different sizes and industries at consistent rates year after year.

​	这与过去几项 Go 开发者调查在统计上没有变化——我们持续听到来自不同国家/地区、不同规模和不同行业的人员的反馈，年复一年保持着稳定的增长率。

![Chart of the different organization sizes where respondents use Go](./GoDeveloperSurvey2024H1Results_img/org_size.svg+xml)

![Chart of the different industries where respondents use Go](./GoDeveloperSurvey2024H1Results_img/industry.svg+xml)

![Chart of countries or regions where respondents are located](./GoDeveloperSurvey2024H1Results_img/location.svg+xml)

## Methodology 方法论

Prior to 2021, we announced the survey primarily through the Go Blog, where it was picked up on various social channels like Twitter, Reddit, or Hacker News. In 2021 we introduced a new way to recruit respondents by using the VS Code Go plugin to randomly select users to be shown a prompt asking if they’d like to participate in the survey. This created a random sample that we used to compare the self-selected respondents from our traditional channels and helped identify potential effects of [self-selection bias](https://en.wikipedia.org/wiki/Self-selection_bias). For this cycle, our friends at JetBrains generously provided us with an additional random sample by prompting a random subset of GoLand users to take the survey!

​	在 2021 年之前，我们主要通过 Go 博客宣布调查，该博客在 Twitter、Reddit 或 Hacker News 等各种社交渠道上被采用。2021 年，我们引入了一种新的招募受访者方式，即使用 VS Code Go 插件随机选择用户，向他们显示一个提示，询问他们是否愿意参与调查。这创建了一个随机样本，我们用它来比较来自我们传统渠道的自我选择受访者，并帮助识别自我选择偏差的潜在影响。对于此周期，JetBrains 的朋友们慷慨地为我们提供了一个额外的随机样本，他们提示 GoLand 用户的随机子集参加调查！

64% of survey respondents “self-selected” to take the survey, meaning they found it on the Go blog or other social Go channels. People who don’t follow these channels are less likely to learn about the survey from them, and in some cases, they respond differently than people who do closely follow them. For example, they might be new to the Go community and not yet aware of the Go blog. About 36% of respondents were randomly sampled, meaning they responded to the survey after seeing a prompt in VS Code (25%) or GoLand (11%). Over the period of January 23 – February 13, there was roughly a 10% chance that users would have seen this prompt. By examining how the randomly sampled groups differ from the self-selected responses, as well as from each other, we’re able to more confidently generalize findings to the larger community of Go developers.

​	64% 的调查受访者“自选”参加调查，这意味着他们在 Go 博客或其他社交 Go 频道上找到了它。不关注这些频道的人不太可能从这些频道了解到调查，在某些情况下，他们的回答与密切关注这些频道的人不同。例如，他们可能是 Go 社区的新手，还不了解 Go 博客。大约 36% 的受访者是随机抽样的，这意味着他们在看到 VS Code（25%）或 GoLand（11%）中的提示后回答了调查。在 1 月 23 日至 2 月 13 日期间，用户看到此提示的几率约为 10%。通过检查随机抽样组与自选回答以及彼此之间的差异，我们能够更自信地将调查结果推广到更大的 Go 开发者社区。

![Chart of different sources of survey respondents](./GoDeveloperSurvey2024H1Results_img/source.svg+xml)

### How to read these results 如何阅读这些结果

Throughout this report we use charts of survey responses to provide supporting evidence for our findings. All of these charts use a similar format. The title is the exact question that survey respondents saw. Unless otherwise noted, questions were multiple choice and participants could only select a single response choice; each chart’s subtitle will tell the reader if the question allowed multiple response choices or was an open-ended text box instead of a multiple choice question. For charts of open-ended text responses, a Go team member read and manually categorized all of the responses. Many open-ended questions elicited a wide variety of responses; to keep the chart sizes reasonable, we condensed them to a maximum of the top 10-12 themes, with additional themes all grouped under “Other”. The percentage labels shown in charts are rounded to the nearest integer (e.g., 1.4% and 0.8% will both be displayed as 1%), but the length of each bar and row ordering are based on the unrounded values.

​	在整个报告中，我们使用调查回复图表来为我们的发现提供支持证据。所有这些图表都使用类似的格式。标题是调查受访者看到的确切问题。除非另有说明，否则问题是多项选择题，参与者只能选择一个回复选项；每个图表的小标题将告诉读者该问题是否允许多项回复选项，或者是一个开放式文本框，而不是多项选择题。对于开放式文本回复的图表，Go 团队成员阅读并手动对所有回复进行分类。许多开放式问题引发了各种各样的回复；为了保持图表大小合理，我们将它们浓缩到最多 10-12 个主题，其他所有主题都归入“其他”之下。图表中显示的百分比标签四舍五入到最接近的整数（例如，1.4% 和 0.8% 都将显示为 1%），但每条栏和行顺序的长度基于未舍入的值。

To help readers understand the weight of evidence underlying each finding, we included error bars showing the 95% [confidence interval](https://en.wikipedia.org/wiki/Confidence_interval) for responses; narrower bars indicate increased confidence. Sometimes two or more responses have overlapping error bars, which means the relative order of those responses is not statistically meaningful (i.e., the responses are effectively tied). The lower right of each chart shows the number of people whose responses are included in the chart, in the form “n = [number of respondents]”. In cases where we found interesting differences in responses between groups, (e.g., years of experience, organization size, or sample source) we showed a color-coded breakdown of the differences.

​	为了帮助读者理解每项发现背后的证据权重，我们包含了显示响应 95% 置信区间的误差线；较窄的线表示置信度更高。有时，两个或更多响应具有重叠的误差线，这意味着这些响应的相对顺序在统计上没有意义（即，响应实际上是平局）。每个图表右下角显示了图表中包含的响应人数，格式为“n = [受访者人数]”。在我们在组之间发现响应存在有趣差异的情况下（例如，经验年数、组织规模或样本来源），我们显示了差异的彩色编码细分。

## Closing 闭幕

And that’s it for our semi-annual Go Developer Survey. Many thanks to everyone who shared their thoughts on Go and everyone who contributed to making this survey happen! It means the world to us and truly helps us improve Go.

​	这就是我们的半年一度的 Go 开发者调查。非常感谢所有分享他们对 Go 的想法的人，以及所有为这次调查的发生做出贡献的人！这对我们来说意义重大，并且真正帮助我们改进了 Go。

This year we’re also excited to announce the forthcoming release of this survey’s dataset. We expect to share this anonymized data by the end of April, allowing anyone to slice and dice survey responses as needed to answer their own questions about the Go ecosystem.

​	今年，我们还很高兴地宣布即将发布此调查的数据集。我们预计在 4 月底前分享这些匿名数据，以便任何人都可以根据需要对调查回复进行切分和分析，以回答他们自己关于 Go 生态系统的问题。

Updated 2024-05-03: We unfortunately need to delay the release of this dataset. We’re still working to make this happen, but we don’t expect to be able to share it until the second half of 2024.

​	更新 2024-05-03：我们很遗憾地需要延迟发布此数据集。我们仍在努力实现这一目标，但我们预计要到 2024 年下半年才能分享它。

— Alice and Todd (on behalf of the Go team at Google)

​	——爱丽丝和托德（代表 Google 的 Go 团队）