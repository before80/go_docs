+++
title = "更新 go 行为准则"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Updating the Go Code of Conduct - 更新 go 行为准则

https://go.dev/blog/conduct-2018

Steve Francia
23 May 2018

In November 2015, we introduced the Go Code of Conduct. It was developed in a collaboration between the Go team members at Google and the Go community. I was fortunate to be one of the community members invited to participate in both drafting and then enforcing the Go Code of Conduct. Since then, we have learned two lessons about limitations in our code of conduct that restricted us from being able to cultivate the safe culture essential to Go’s success.

2015年11月，我们推出了《Go行为准则》。它是由谷歌的Go团队成员和Go社区合作制定的。我很幸运地成为社区成员之一，被邀请参与起草和执行Go行为准则。从那时起，我们在行为准则的局限性方面学到了两个教训，这些局限性限制了我们培养对Go的成功至关重要的安全文化。

The first lesson we learned is that toxic behaviors by project participants in non-project spaces can have a negative impact on the project affecting the security and safety of community members. There were a few reported incidents where actions took place outside of project spaces but the impact was felt inside our community. The specific language in our code of conduct restricted our ability to respond only to actions happening “in the official forums operated by the Go project”. We needed a way to protect our community members wherever they are.

我们学到的第一个教训是，项目参与者在非项目场所的有毒行为会对项目产生负面影响，影响社区成员的安全和保障。有几起报告的事件是在项目空间之外发生的，但在我们的社区内却感受到了影响。我们的行为准则中的具体语言限制了我们的能力，使我们只能对 "在Go项目运作的官方论坛上 "发生的行为做出反应。我们需要一种方法来保护我们的社区成员，无论他们在哪里。

The second lesson we learned is that the demands required to enforce the code of conduct place too heavy of a burden on volunteers. The initial version of the code of conduct presented the working group as disciplinarians. It was soon clear that this was too much, so in early 2017 [we changed the group’s role](https://go.dev/cl/37014) to that of advisors and mediators. Still, working group community members reported feeling overwhelmed, untrained, and vulnerable. This well-intentioned shift left us without an enforcement mechanism without solving the issue with overburdened volunteers.

我们学到的第二个教训是，执行行为准则所需的要求给志愿者带来了太重的负担。行为准则的最初版本将工作组描述为惩戒者。很快，我们就发现这太过分了，所以在2017年初，我们把工作组的角色改为顾问和调解员。但是，工作组的社区成员仍然报告说，他们感到不堪重负，没有受过训练，而且很脆弱。这一用心良苦的转变使我们没有一个执行机制，而没有解决志愿者负担过重的问题。

In mid-2017, I represented the Go project in a meeting with Google’s Open Source Programs Office and Open Source Strategy Team to address the shortcomings in our respective codes of conduct, particularly in their enforcement. It quickly became clear that our problems had a lot in common, and that working together on a single code of conduct for all of Google’s open source projects made sense. We started with the text from the Contributor Covenant Code of Conduct v1.4 and then made changes, influenced by our experiences in Go community and our collective experiences in open source. This resulted in the Google [code of conduct template](https://opensource.google.com/docs/releasing/template/CODE_OF_CONDUCT/).

2017年年中，我代表Go项目参加了与谷歌开源项目办公室和开源战略团队的会议，以解决我们各自行为准则中的缺陷，特别是在执行方面。我们很快就发现，我们的问题有很多共同点，共同为谷歌所有的开源项目制定一个单一的行为准则是有意义的。我们从《贡献者盟约行为守则》V1.4版的文本开始，然后根据我们在Go社区的经验和我们在开源领域的集体经验进行了修改。这就形成了谷歌的行为准则模板。

Today the Go project is adopting this new code of conduct, and we’ve updated [golang.org/conduct](https://go.dev/conduct). This revised code of conduct retains much of the intent, structure and language of the original Go code of conduct while making two critical changes that address the shortcomings identified above.

今天，Go项目正在采用这个新的行为准则，我们已经更新了golang.org/conduct。这个修订后的行为准则保留了原始Go行为准则的大部分意图、结构和语言，同时做出了两个关键性的改变，以解决上述的缺陷。

First, [the new code of conduct makes clear](https://go.dev/conduct/#scope) that people who participate in any kind of harassment or inappropriate behavior, even outside our project spaces, are not welcome in our project spaces. This means that the Code of Conduct applies outside the project spaces when there is a reasonable belief that an individual’s behavior may have a negative impact on the project or its community.

首先，新的行为准则明确指出，参与任何形式的骚扰或不当行为的人，即使在我们的项目空间之外，也不欢迎他们进入我们的项目空间。这意味着，当有理由相信一个人的行为可能对项目或其社区产生负面影响时，该行为准则适用于项目空间之外。

Second, in the place of the working group, [the new code of conduct introduces a single Project Steward](https://go.dev/conduct/#reporting) who will have explicit training and support for this role. The Project Steward will receive reported violations and then work with a committee, consisting of representatives from the Open Source Programs Office and the Google Open Source Strategy team, to find a resolution.

第二，在工作组的位置上，新的行为准则引入了一个单一的项目管理员，他将有明确的培训和支持来扮演这个角色。项目管理人将接受报告的违规行为，然后与一个由开源项目办公室和谷歌开源战略团队的代表组成的委员会合作，寻找解决方案。

Our first Project Steward will be [Cassandra Salisbury](https://twitter.com/cassandraoid). She is well known to the Go community as a member of Go Bridge, an organizer of many Go meetups and conferences, and as a lead of the Go community outreach working group. Cassandra now works on the Go team at Google with a focus on advocating for and supporting the Go community.

我们的第一位项目管理员将是Cassandra Salisbury。她是Go Bridge的成员，是许多Go聚会和会议的组织者，也是Go社区推广工作小组的负责人，因此在Go社区中很有名气。Cassandra现在在Google的Go团队工作，主要负责宣传和支持Go社区。

We are grateful to everyone who served on the original Code of Conduct Working Group. Your efforts were essential in creating an inclusive and safe community.

我们感谢所有在最初的行为准则工作组中工作的人。你们的努力对于创建一个包容和安全的社区至关重要。

We believe the code of conduct has contributed to the Go project becoming more welcoming now than it was in 2015, and we should all be proud of that.

我们相信，行为准则为Go项目现在变得比2015年时更受欢迎做出了贡献，我们都应该为此感到自豪。

We hope that the new code of conduct will help protect our community members even more effectively.

我们希望新的行为准则能更有效地帮助保护我们的社区成员。
