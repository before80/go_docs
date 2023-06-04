+++
title = "使用 go Cloud的便携式云编程"
weight = 10
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Portable Cloud Programming with Go Cloud - 使用 go Cloud的便携式云编程

https://go.dev/blog/go-cloud

Eno Compton and Cassandra Salisbury
24 July 2018

## Introduction 简介

Today, the Go team at Google is releasing a new open source project, [Go Cloud](https://github.com/google/go-cloud), a library and tools for developing on the [open cloud](https://cloud.google.com/open-cloud/). With this project, we aim to make Go the language of choice for developers building portable cloud applications.

今天，谷歌的Go团队发布了一个新的开源项目--Go Cloud，这是一个用于在开放云上开发的库和工具。通过这个项目，我们的目标是让Go成为开发人员构建可移植云应用程序的首选语言。

This post explains why we started this project, the details of how Go Cloud works, and how to get involved.

这篇文章解释了我们为什么要启动这个项目，Go Cloud如何工作的细节，以及如何参与。

## Why portable cloud programming? Why now? 为什么要进行可移植云编程？为什么是现在？

We estimate there are now [over one million](https://research.swtch.com/gophercount) Go developers worldwide. Go powers many of the most critical cloud infrastructure projects, including Kubernetes, Istio, and Docker. Companies like Lyft, Capital One, Netflix and [many more](https://github.com/golang/go/wiki/GoUsers) are depending on Go in production. Over the years, we’ve found that developers love Go for cloud development because of its efficiency, productivity, built-in concurrency, and low latency.

我们估计现在全球有超过一百万的Go开发者。Go为许多最关键的云基础设施项目提供动力，包括Kubernetes、Istio和Docker。像Lyft、Capital One、Netflix等公司都在生产中依赖Go。多年来，我们发现，开发人员喜欢Go的云开发，因为它的效率、生产力、内置并发性和低延迟。

As part of our work to support Go’s rapid growth, we have been interviewing teams who work with Go to understand how they use the language and how the Go ecosystem can improve further. One common theme with many organizations is the need for portability across cloud providers. These teams want to deploy robust applications in [multi-cloud](https://en.wikipedia.org/wiki/Cloud_computing#Multicloud) and [hybrid-cloud](https://en.wikipedia.org/wiki/Cloud_computing#Hybrid_cloud) environments, and migrate their workloads between cloud providers without significant changes to their code.

作为我们支持Go快速增长工作的一部分，我们一直在采访使用Go的团队，以了解他们如何使用该语言以及Go生态系统如何进一步改善。许多组织的一个共同主题是需要跨云供应商的可移植性。这些团队希望在多云和混合云环境中部署强大的应用程序，并在不对代码进行重大修改的情况下在云提供商之间迁移他们的工作负载。

To achieve this, some teams attempt to decouple their applications from provider-specific APIs in order to produce simpler and more portable code. However the short-term pressure to ship features means teams often sacrifice longer-term efforts toward portability. As a result, most Go applications running in the cloud are tightly coupled to their initial cloud provider.

为了实现这一目标，一些团队试图将他们的应用程序与供应商的特定API解耦，以产生更简单、更可移植的代码。然而，短期内的功能压力意味着团队往往会牺牲长期的努力来实现可移植性。因此，大多数在云中运行的 Go 应用程序都与最初的云提供商紧密相连。

As an alternative, teams can use Go Cloud, a set of open generic cloud APIs, to write simpler and more portable cloud applications. Go Cloud also sets the foundation for an ecosystem of portable cloud libraries to be built on top of these generic APIs. Go Cloud makes it possible for teams to meet their feature development goals while also preserving the long-term flexibility for multi-cloud and hybrid-cloud architectures. Go Cloud applications can also migrate to the cloud providers that best meet their needs.

作为替代方案，团队可以使用 Go Cloud（一套开放的通用云 API）来编写更简单、更便携的云应用程序。Go Cloud还为在这些通用API基础上建立一个可移植的云库生态系统奠定了基础。Go Cloud使团队有可能实现其功能开发目标，同时也为多云和混合云架构保留了长期的灵活性。Go Cloud应用程序还可以迁移到最能满足其需求的云供应商。

## What is Go Cloud? 什么是Go Cloud？

We have identified common services used by cloud applications and have created generic APIs to work across cloud providers. Today, Go Cloud is launching with blob storage, MySQL database access, runtime configuration, and an HTTP server configured with request logging, tracing, and health checking. Go Cloud offers support for Google Cloud Platform (GCP) and Amazon Web Services (AWS). We plan to work with cloud industry partners and the Go community to add support for additional cloud providers very soon.

我们已经确定了云应用程序使用的常见服务，并创建了通用的API，以跨云供应商工作。今天，Go Cloud推出了blob存储、MySQL数据库访问、运行时配置以及配置有请求记录、跟踪和健康检查的HTTP服务器。Go Cloud提供对谷歌云平台（GCP）和亚马逊网络服务（AWS）的支持。我们计划与云计算行业的合作伙伴和Go社区合作，很快增加对其他云提供商的支持。

Go Cloud aims to develop vendor-neutral generic APIs for the most-used services across cloud providers such that deploying a Go application on another cloud is simple and easy. Go Cloud also lays the foundation for other open source projects to write cloud libraries that work across providers. Community feedback, from all types of developers at all levels, will inform the priority of future APIs in Go Cloud.

Go Cloud的目标是为各云提供商最常用的服务开发供应商中立的通用API，以便在其他云上部署Go应用程序变得简单而容易。Go Cloud还为其他开源项目奠定了基础，使其能够编写跨供应商的云库。来自各个层次的各类开发人员的社区反馈，将为 Go Cloud 中未来 API 的优先级提供信息。

## How does it work? 它是如何工作的？

At the core of Go Cloud is a collection of generic APIs for portable cloud programming. Let’s look at an example of using blob storage. You can use the generic type [`*blob.Bucket`](https://godoc.org/github.com/google/go-cloud/blob#Bucket) to copy a file from a local disk to a cloud provider. Let’s start by opening an S3 bucket using the included [s3blob package](https://godoc.org/github.com/google/go-cloud/blob/s3blob):

Go Cloud的核心是用于可移植云编程的通用API集合。让我们看一下使用blob存储的例子。您可以使用通用类型*blob.Bucket将一个文件从本地磁盘复制到云提供商。让我们先用附带的s3blob包打开一个S3 bucket：

```go linenums="1"
// setupBucket opens an AWS bucket.
func setupBucket(ctx context.Context) (*blob.Bucket, error) {
    // Obtain AWS credentials.
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-2"),
    })
    if err != nil {
        return nil, err
    }
    // Open a handle to s3://go-cloud-bucket.
    return s3blob.OpenBucket(ctx, sess, "go-cloud-bucket")
}
```

Once a program has a `*blob.Bucket`, it can create a `*blob.Writer`, which implements `io.Writer`. From there, the program can use the `*blob.Writer` to write data to the bucket, checking that `Close` does not report an error.

一旦程序有了*blob.Bucket，它就可以创建一个*blob.Writer，它实现了io.Writer。在这里，程序可以使用*blob.Writer向Bucket写入数据，并检查Close是否报告错误。

```go linenums="1"
ctx := context.Background()
b, err := setupBucket(ctx)
if err != nil {
    log.Fatalf("Failed to open bucket: %v", err)
}
data, err := ioutil.ReadFile("gopher.png")
if err != nil {
    log.Fatalf("Failed to read file: %v", err)
}
w, err := b.NewWriter(ctx, "gopher.png", nil)
if err != nil {
    log.Fatalf("Failed to obtain writer: %v", err)
}
_, err = w.Write(data)
if err != nil {
    log.Fatalf("Failed to write to bucket: %v", err)
}
if err := w.Close(); err != nil {
    log.Fatalf("Failed to close: %v", err)
}
```

Notice how the logic of using the bucket does not refer to AWS S3. Go Cloud makes swapping out cloud storage a matter of changing the function used to open the `*blob.Bucket`. The application could instead use Google Cloud Storage by constructing a `*blob.Bucket` using [`gcsblob.OpenBucket`](https://godoc.org/github.com/google/go-cloud/blob/gcsblob#OpenBucket) without changing the code that copies the file:

请注意，使用桶的逻辑并没有提到AWS S3。Go Cloud使得更换云存储只是改变用于打开*blob.Bucket的函数的问题。应用程序可以通过使用gcsblob.OpenBucket构建一个*blob.Bucket来代替使用谷歌云存储，而无需改变复制文件的代码：

```go linenums="1"
// setupBucket opens a GCS bucket.
func setupBucket(ctx context.Context) (*blob.Bucket, error) {
    // Open GCS bucket.
    creds, err := gcp.DefaultCredentials(ctx)
    if err != nil {
        return nil, err
    }
    c, err := gcp.NewHTTPClient(gcp.DefaultTransport(), gcp.CredentialsTokenSource(creds))
    if err != nil {
        return nil, err
    }
    // Open a handle to gs://go-cloud-bucket.
    return gcsblob.OpenBucket(ctx, "go-cloud-bucket", c)
}
```

While different steps are needed to access buckets on different cloud providers, the resulting type used by your application is the same: `*blob.Bucket`. This isolates application code from cloud-specific code. To increase interoperability with existing Go libraries, Go Cloud leverages established interfaces like `io.Writer`, `io.Reader`, and `*sql.DB`.

虽然访问不同云提供商的桶需要不同的步骤，但您的应用程序使用的结果类型是相同的。*blob.Bucket。这就将应用程序代码与云端特定代码隔离开来。为了提高与现有 Go 库的互操作性，Go Cloud 利用了 io.Writer、io.Reader 和 *sql.DB 等既定接口。

The setup code needed to access cloud services tends to follow a pattern: higher abstractions are constructed from more basic abstractions. While you could write this code by hand, Go Cloud automates this with **Wire**, a tool that generates cloud-specific setup code for you. The [Wire documentation](https://github.com/google/go-cloud/tree/master/wire) explains how to install and use the tool and the [Guestbook sample](https://github.com/google/go-cloud/tree/master/samples/guestbook) shows Wire in action.

访问云服务所需的设置代码往往遵循一个模式：更高的抽象是由更基本的抽象构建的。虽然您可以手工编写这段代码，但Go Cloud通过Wire实现了自动化，这个工具可以为您生成特定的云设置代码。Wire 文档解释了如何安装和使用该工具，Guestbook 示例展示了 Wire 的运行情况。

## How can I get involved and learn more? 我怎样才能参与并了解更多？

To get started, we recommend following [the tutorial](https://github.com/google/go-cloud/tree/master/samples/tutorial) and then trying to build an application yourself. If you’re already using AWS or GCP, you can try migrating parts of your existing application to use Go Cloud. If you’re using a different cloud provider or an on-premise service, you can extend Go Cloud to support it by implementing the driver interfaces (like [`driver.Bucket`](https://godoc.org/github.com/google/go-cloud/blob/driver#Bucket)).

要想开始，我们建议按照教程进行，然后尝试自己建立一个应用程序。如果您已经在使用AWS或GCP，您可以尝试将您现有的应用程序的一部分迁移到使用Go Cloud。如果您正在使用不同的云提供商或内部服务，您可以通过实现驱动接口（如driver.Bucket）来扩展Go Cloud以支持它。

We appreciate any and all input you have about your experience. [Go Cloud’s](https://github.com/google/go-cloud) development is conducted on GitHub. We are looking forward to contributions, including pull requests. [File an issue](https://github.com/google/go-cloud/issues/new) to tell us what could be better or what future APIs the project should support. For updates and discussion about the project, join [the project’s mailing list](https://groups.google.com/forum/#!forum/go-cloud).

我们感谢您对您的经验的任何和所有的投入。Go Cloud的开发是在GitHub上进行的。我们期待着大家的贡献，包括拉动请求。请提交一个问题，告诉我们什么地方可以做得更好，或者项目未来应该支持哪些API。关于项目的更新和讨论，请加入项目的邮件列表。

The project requires contributors to sign the same Contributor License Agreement as that of the Go project. Read the [contribution guidelines](https://github.com/google/go-cloud/blob/master/CONTRIBUTING.md) for more details. Please note, Go Cloud is covered by the Go [Code of Conduct](https://github.com/google/go-cloud/blob/master/CODE_OF_CONDUCT.md).

本项目要求贡献者签署与Go项目相同的贡献者许可协议。请阅读贡献者指南以了解更多细节。请注意，Go Cloud受到Go行为准则的保护。

Thank you for taking the time to learn about Go Cloud. We are excited to work with you to make Go the language of choice for developers building portable cloud applications.

感谢您花时间了解 Go 云。我们很高兴能与您合作，使 Go 成为开发人员构建可移植云应用程序的首选语言。
