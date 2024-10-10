+++
title = "基准测试"
date = 2024-02-05T09:14:15+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文： [https://docs.gofiber.io/extra/benchmarks]({{< ref "/fiber/Extra/Benchmarks" >}})

# 📊 Benchmarks  基准测试

## TechEmpower TechEmpower TechEmpower 提供了许多 Web 应用程序框架的性能比较，这些框架执行基本任务，例如 JSON 序列化、数据库访问和服务器端模板组合。

[TechEmpower](https://www.techempower.com/benchmarks/#section=data-r19&hw=ph&test=composite) provides a performance comparison of many web application frameworks executing fundamental tasks such as JSON serialization, database access, and server-side template composition.

​	每个框架都在现实的生产配置中运行。结果在云实例和物理硬件上捕获。测试实现主要由社区贡献，所有源代码都可以在 GitHub 存储库中找到。

Each framework is operating in a realistic production configuration. Results are captured on cloud instances and on physical hardware. The test implementations are largely community-contributed and all source is available at the [GitHub repository](https://github.com/TechEmpower/FrameworkBenchmarks).

​	Fiber

- Fiber `v1.10.0`
- 28 HT Cores Intel(R) Xeon(R) Gold 5120 CPU @ 2.20GHz
  28 个 HT 内核英特尔(R) 至强(R) 金牌 5120 CPU @ 2.20GHz
- 32GB RAM
  32GB 内存
- Ubuntu 18.04.3 4.15.0-88-generic
- Dedicated Cisco 10-Gbit Ethernet switch.
  专用思科 10 千兆以太网交换机。

### Plaintext 纯文本 

The Plaintext test is an exercise of the request-routing fundamentals only, designed to demonstrate the capacity of high-performance platforms in particular. Requests will be sent using HTTP pipelining. The response payload is still small, meaning good performance is still necessary in order to saturate the gigabit Ethernet of the test environment.

​	纯文本测试仅是对请求路由基础知识的练习，旨在展示高性能平台的容量。将使用 HTTP 管道发送请求。响应负载仍然很小，这意味着为了使测试环境的千兆以太网饱和，仍然需要良好的性能。

See [Plaintext requirements](https://github.com/TechEmpower/FrameworkBenchmarks/wiki/Project-Information-Framework-Tests-Overview#single-database-query)

​	请参阅纯文本要求

**Fiber** - **6,162,556** responses per second with an average latency of **2.0** ms.

​	光纤 - 每秒 6,162,556 次响应，平均延迟为 2.0 毫秒。
**Express** - **367,069** responses per second with an average latency of **354.1** ms.

​	Express - 每秒 367,069 次响应，平均延迟为 354.1 毫秒。

![img](./Benchmarks_img/plaintext-e25d187f782d18fdd35b84e3d7c625eb.png)

![Fiber vs Express](./Benchmarks_img/plaintext_express-ef6522843412bb5b14b3c6b6a4f032de.png)

### Data Updates 数据更新 

**Fiber** handled **11,846** responses per second with an average latency of **42.8** ms.

​	光纤每秒处理 11,846 次响应，平均延迟为 42.8 毫秒。
**Express** handled **2,066** responses per second with an average latency of **390.44** ms.

​	Express 每秒处理 2,066 个响应，平均延迟为 390.44 毫秒。

![img](./Benchmarks_img/data_updates-3be85c418d6971091854c5086af9ed10.png)

![Fiber vs Express](./Benchmarks_img/data_updates_express-2f55d1b0975ec391d29d823b48faf617.png)

### Multiple Queries 多重查询 

**Fiber** handled **19,664** responses per second with an average latency of **25.7** ms.

​	Fiber 每秒处理 19,664 个响应，平均延迟为 25.7 毫秒。
**Express** handled **4,302** responses per second with an average latency of **117.2** ms.

​	Express 每秒处理 4,302 个响应，平均延迟为 117.2 毫秒。

![img](./Benchmarks_img/multiple_queries-2c2e81674208b90b9aeb1cb791a3f0dc.png)

![Fiber vs Express](./Benchmarks_img/multiple_queries_express-ec4dc8013e85dc2a2fa4f5eeb55ce8dd.png)

### Single Query 单一查询 

**Fiber** handled **368,647** responses per second with an average latency of **0.7** ms.

​	Fiber 每秒处理 368,647 个响应，平均延迟为 0.7 毫秒。
**Express** handled **57,880** responses per second with an average latency of **4.4** ms.

​	Express 每秒处理 57,880 个响应，平均延迟为 4.4 毫秒。

![img](./Benchmarks_img/single_query-4f7782d3c3ff91e92ac27e382b09f6ac.png)

![Fiber vs Express](./Benchmarks_img/single_query_express-d8e41422b4f5c0a9496272e4a66a97c4.png)

### JSON Serialization JSON 序列化 

**Fiber** handled **1,146,667** responses per second with an average latency of **0.4** ms.

​	Fiber 每秒处理 1,146,667 个响应，平均延迟为 0.4 毫秒。
**Express** handled **244,847** responses per second with an average latency of **1.1** ms.

​	Express 每秒处理 244,847 个响应，平均延迟为 1.1 毫秒。

![img](./Benchmarks_img/json-62868f61b34e3790f3a8b3b52b1a3a3b.png)

![Fiber vs Express](./Benchmarks_img/json_express-aa631b2de86808970aa4bb7c9c9d3edf.png)

## Go web framework benchmark Go web 框架基准 

🔗 https://github.com/smallnest/go-web-framework-benchmark

- **CPU** Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz
- **MEM** 4GB
  内存 4GB
- **GO** go1.13.6 linux/amd64
- **OS** Linux
  操作系统 Linux

The first test case is to mock **0 ms**, **10 ms**, **100 ms**, **500 ms** processing time in handlers.

​	第一个测试用例是在处理程序中模拟 0 毫秒、10 毫秒、100 毫秒、500 毫秒的处理时间。

![img](./Benchmarks_img/benchmark-18e23fcf42afc7f5e12ea23aceb27885.png)

The concurrency clients are **5000**.

​	并发客户端为 5000。

![img](./Benchmarks_img/benchmark_latency-b67a470cf1b261c3092b80cbf42ef16b.png)

Latency is the time of real processing time by web servers. *The smaller is the better.*

​	延迟是 Web 服务器实际处理时间。越小越好。

![img](./Benchmarks_img/benchmark_alloc-dec96faa96e07bcec84f40a4dfc8d187.png)

Allocs is the heap allocations by web servers when test is running. The unit is MB. *The smaller is the better.*

​	在测试运行时，Allocs 是 Web 服务器的堆分配。单位是 MB。越小越好。

If we enable **http pipelining**, test result as below:

​	如果我们启用 http 流水线，测试结果如下：

![img](./Benchmarks_img/benchmark-pipeline-b49cbb1db36293acdfb0e6c96d844e1a.png)

Concurrency test in **30 ms** processing time, the test result for **100**, **1000**, **5000** clients is:

​	在 30 毫秒处理时间内的并发测试，100、1000、5000 个客户端的测试结果为：

![img](./Benchmarks_img/concurrency-1307e1d23c01a561a4b2a0f5bdd7e1bc.png)

![img](./Benchmarks_img/concurrency_latency-5a223848a8bee8df21cc02451f0db2b6.png)

![img](./Benchmarks_img/concurrency_alloc-6f2d485576803f7de2fe0a1deca21a09.png)

If we enable **http pipelining**, test result as below:

​	如果我们启用 http 管道，测试结果如下：

![img](./Benchmarks_img/concurrency-pipeline-b0d3c211d9c7cb5474fd191223a41241.png)

Dependency graph for `v1.9.0`

​	 `v1.9.0` 的依赖关系图

![img](./Benchmarks_img/graph-afbd400b1c3e1c6f137dae3cfc1890ce.svg)
