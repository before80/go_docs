+++
title = "go 云开发工具包的新内容"
weight = 19
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# What's new in the Go Cloud Development Kit - go 云开发工具包的新内容

https://go.dev/blog/go-cloud2019

The Go Cloud Development Kit team at Google  谷歌的Go云开发工具包团队

4 March 2019 

2019年3月4日

## Introduction 简介

Last July, we [introduced](https://blog.golang.org/go-cloud) the [Go Cloud Development Kit](https://gocloud.dev/) (previously referred to as simply "Go Cloud"), an open source project building libraries and tools to improve the experience of developing for the cloud with Go. We’ve made a lot of progress since then – thank you to early contributors! We look forward to growing the Go CDK community of users and contributors, and are excited to work closely with early adopters.

去年 7 月，我们推出了 Go 云开发工具包（以前简称为 "Go 云"），这是一个构建库和工具的开源项目，旨在改善使用 Go 进行云开发的体验。从那时起，我们已经取得了很大的进展--感谢早期的贡献者们！我们期待着Go CD的发展。我们期待着Go CDK用户和贡献者社区的发展，并很高兴能与早期采用者密切合作。

## Portable APIs 可移植的API

Our first initiative is a set of portable APIs for common cloud services. You write your application using these APIs, and then deploy it on any combination of providers, including AWS, GCP, Azure, on-premise, or on a single developer machine for testing. Additional providers can be added by implementing an interface.

我们的第一项举措是为常见的云服务提供了一套可移植的API。您可以使用这些API编写您的应用程序，然后将其部署在任何供应商的组合上，包括AWS、GCP、Azure、内部部署，或在单个开发人员机器上进行测试。可以通过实现一个接口来添加额外的供应商。

These portable APIs are a great fit if any of the following are true:

如果以下情况属实，这些可移植的API是非常合适的：

- You develop cloud applications locally. 您在本地开发云应用程序。
- You have on-premise applications that you want to run in the cloud (permanently, or as part of a migration). 您有希望在云中运行的内部应用程序（永久的，或作为迁移的一部分）。
- You want portability across multiple clouds. 您希望在多个云端有可移植性。
- You are creating a new Go application that will use cloud services. 您正在创建一个将使用云服务的新Go应用程序。

Unlike traditional approaches where you would need to write new application code for each cloud provider, with the Go CDK you write your application code once using our portable APIs to access the set of services listed below. Then, you can run your application on any supported cloud with minimal config changes.

与传统方法不同的是，您需要为每个云提供商编写新的应用程序代码，而使用 Go CDK，您只需使用我们的可移植 API 编写一次应用程序代码，以访问下面列出的一组服务。然后，您可以在任何支持的云上运行您的应用程序，只需进行最小的配置更改。

Our current set of APIs includes:

我们目前的API集包括：

- [blob](https://godoc.org/gocloud.dev/blob), for persistence of blob data. Supported providers include: AWS S3, Google Cloud Storage (GCS), Azure Storage, the filesystem, and in-memory. blob，用于blob数据的持久化。支持的供应商包括。AWS S3、谷歌云存储（GCS）、Azure存储、文件系统和内存。
- [pubsub](https://godoc.org/gocloud.dev/pubsub) for publishing/subscribing of messages to a topic. Supported providers include: Amazon SNS/SQS, Google Pub/Sub, Azure Service Bus, RabbitMQ, and in-memory. pubsub，用于向主题发布/订阅消息。支持的提供者包括。亚马逊SNS/SQS、谷歌Pub/Sub、Azure服务总线、RabbitMQ和内存中。
- [runtimevar](https://godoc.org/gocloud.dev/runtimevar), for watching external configuration variables. Supported providers include AWS Parameter Store, Google Runtime Configurator, etcd, and the filesystem. runtimevar，用于观察外部配置变量。支持的供应商包括：AWS参数存储、谷歌运行时配置器、etcd和文件系统。
- [secrets](https://godoc.org/gocloud.dev/secrets), for encryption/decryption. Supported providers include AWS KMS, GCP KMS, Hashicorp Vault, and local symmetric keys. secrets，用于加密/解密。支持的供应商包括AWS KMS、GCP KMS、Hashicorp Vault和本地对称密钥。
- Helpers for connecting to cloud SQL providers. Supported providers include AWS RDS and Google Cloud SQL. 用于连接云SQL提供商的助手。支持的供应商包括AWS RDS和谷歌云SQL。
- We are also working on a document storage API (e.g. MongoDB, DynamoDB, Firestore). 我们还在开发一个文档存储API（如MongoDB、DynamoDB、Firestore）。

## Feedback 反馈意见

We hope you’re as excited about the Go CDK as we are – check out our [godoc](https://godoc.org/gocloud.dev), walk through our [tutorial](https://github.com/google/go-cloud/tree/master/samples/tutorial), and use the Go CDK in your application(s). We’d love to hear your ideas for other APIs and API providers you’d like to see.

我们希望您和我们一样对 Go CDK 感到兴奋 - 查看我们的 godoc，学习我们的教程，并在您的应用程序中使用 Go CDK。我们希望听到您对其他 API 和 API 供应商的想法。

If you’re digging into Go CDK please share your experiences with us:

如果您正在研究Go CDK，请与我们分享您的经验：

- What went well? 哪些方面进展顺利？
- Were there any pain points using the APIs? 使用API时有什么痛点吗？
- Are there any features missing in the API you used? 在您使用的API中是否有任何功能缺失？
- Suggestions for documentation improvements. 对文档的改进建议。

To send feedback, you can:

要发送反馈，您可以：

- Submit issues to our public [GitHub repository](https://github.com/google/go-cloud/issues/new/choose). 向我们的公共GitHub仓库提交问题。
- Email [go-cdk-feedback@google.com](mailto:go-cdk-feedback@google.com). 发送电子邮件到 go-cdk-feedback@google.com。
- Post to our [public Google group](https://groups.google.com/forum/#!forum/go-cloud). 发布到我们的公共谷歌小组。

Thanks!

谢谢!
