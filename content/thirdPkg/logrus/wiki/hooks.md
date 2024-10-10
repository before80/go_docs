+++
title = "hooks"
date = 2023-06-25T09:41:29+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Hooks

> 原文：[https://github.com/sirupsen/logrus/wiki/Hooks](https://github.com/sirupsen/logrus/wiki/Hooks)

Owefsad 在 2 月 1 日编辑了此页面·[35 次修订](https://github.com/sirupsen/logrus/wiki/Hooks/_history)



​	本页面列出了已知的钩子服务列表。这些列表按原样提供，未经 logrus 开发人员的审核或审查。

| Hook                                                         | 描述                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [Airbrake "legacy"](https://github.com/gemnasium/logrus-airbrake-legacy-hook) | 将错误发送到与Airbrake API V2兼容的异常追踪服务。在幕后使用[`airbrake-go`](https://github.com/tobi/airbrake-go)。 |
| [Airbrake](https://github.com/gemnasium/logrus-airbrake-hook) | 将错误发送到Airbrake API V3。在幕后使用官方的[`gobrake`](https://github.com/airbrake/gobrake)。 |
| [Amazon Kinesis](https://github.com/evalphobia/logrus_kinesis) | 用于将日志记录到[Amazon Kinesis](https://aws.amazon.com/kinesis/)的钩子 |
| [Amazon SNS](https://github.com/stvvan/logrus-sns)           | 用于将日志记录到[Amazon Simple Notification Service (SNS)](https://aws.amazon.com/sns/)的钩子 |
| [Amazon EventBridge](https://github.com/teddy-schmitz/eventbridge_logrus) | 用于将日志记录到[Amazon EventBridge](https://aws.amazon.com/eventbridge/)的Hook |
| [Amqp-Hook](https://github.com/vladoatanasov/logrus_amqp)    | 用于将日志记录到Amqp代理（如RabbitMQ）的Hook                 |
| [Anexia CloudLog](https://github.com/anexia-it/go-logrus-cloudlog) | 用于将日志记录到Anexia CloudLog的钩子                        |
| [Application Insights](https://github.com/jjcollinge/logrus-appinsights) | 用于将日志记录到[Application Insights](https://azure.microsoft.com/en-us/services/application-insights/)的钩子 |
| [AzureTableHook](https://github.com/kpfaulkner/azuretablehook/) | 用于将日志记录到Azure Table Storage的钩子                    |
| [Bugsnag](https://github.com/Shopify/logrus-bugsnag/blob/master/bugsnag.go) | 将错误发送到Bugsnag异常追踪服务的钩子                        |
| [ClickHouse](https://github.com/oxgrouby/logrus-clickhouse-hook) | 将日志发送到[ClickHouse](https://clickhouse.yandex/)的钩子   |
| [Datadog Log](https://github.com/bin3377/logrus-datadog-hook) | 用于通过HTTP端点将日志记录到[Datadog](https://www.datadoghq.com/)的钩子 |
| [Discord Bot Hook](https://github.com/outdead/discordbotrus) | 用于通过Discord应用将日志记录到[Discord](https://discordapp.com/)的钩子 |
| [Discordrus](https://github.com/kz/discordrus)               | 用于将日志记录到[Discord](https://discordapp.com/)的钩子     |
| [Elastic APM](https://godoc.org/go.elastic.co/apm/module/apmlogrus#Hook) | 用于将错误日志记录到[Elastic APM](https://www.elastic.co/solutions/apm)的钩子 |
| [ElasticSearch（使用官方客户端）](https://github.com/go-extras/elogrus) | 用于将日志记录到ElasticSearch的钩子                          |
| [Firehose](https://github.com/beaubrewer/logrus_firehose)    | 用于将日志记录到[Amazon Firehose](https://aws.amazon.com/kinesis/firehose/)的钩子 |
| [Fluentd](https://github.com/evalphobia/logrus_fluent)       | 用于将日志记录到fluentd的钩子                                |
| [Go-Slack](https://github.com/multiplay/go-slack)            | 用于将日志记录到[Slack](https://slack.com/)的钩子            |
| [Graylog](https://github.com/gemnasium/logrus-graylog-hook)  | 用于将日志记录到[Graylog](http://graylog2.org/)的钩子        |
| [Hiprus](https://github.com/nubo/hiprus)                     | 将错误发送到Hipchat频道的钩子                                |
| [Honeybadger](https://github.com/agonzalezro/logrus_honeybadger) | 用于将异常发送到Honeybadger的钩子                            |
| [InfluxDB](https://github.com/Abramovic/logrus_influxdb)     | 用于将日志记录到influxdb的钩子                               |
| [Influxus](http://github.com/vlad-doru/influxus)             | 用于同时将日志记录到[InfluxDB](http://influxdata.com/)的钩子 |
| [Journalhook](https://github.com/wercker/journalhook)        | 用于将日志记录到`systemd-journald`的钩子                     |
| [KafkaLogrus](https://github.com/tracer0tong/kafkalogrus)    | 用于将日志记录到Kafka的钩子                                  |
| [Kafka REST Proxy](https://github.com/Nordstrom/logrus-kafka-rest-proxy) | 用于将日志记录到[Kafka REST Proxy](https://docs.confluent.io/current/kafka-rest/docs)的钩子 |
| [LFShook](https://github.com/rifflock/lfshook)               | 用于将日志记录到本地文件系统的钩子                           |
| [Logbeat](https://github.com/macandmia/logbeat)              | 用于将日志记录到[Opbeat](https://opbeat.com/)的钩子          |
| [Logentries](https://github.com/jcftang/logentriesrus)       | 用于将日志记录到[Logentries](https://logentries.com/)的钩子  |
| [Logentrus](https://github.com/puddingfactory/logentrus)     | 用于将日志记录到[Logentries](https://logentries.com/)的钩子  |
| [Logmatic.io](https://github.com/logmatic/logmatic-go)       | 用于将日志记录到 [Logmatic.io](http://logmatic.io/) 的钩子   |
| [Logrus Boltdb Hook](https://github.com/trK54Ylmz/logrus-boltdb-hook) | 用于将日志记录到 boltdb 的钩子                               |
| [Logrus Bolt Hook](https://github.com/kennykarnama/logrus-bolt-hook) | 用于将日志记录到 boltdb 的钩子                               |
| [Logrusly](https://github.com/sebest/logrusly)               | 将日志发送到 [Loggly](https://www.loggly.com/) 的钩子        |
| [Logstash](https://github.com/bshuster-repo/logrus-logstash-hook) | 用于将日志记录到 [Logstash](https://www.elastic.co/products/logstash) 的钩子 |
| [Loki](https://github.com/YuKitsune/lokirus)                 | 用于将日志记录到 [Loki](https://grafana.com/oss/loki/) 的钩子 |
| [Lumberjackrus](https://github.com/orandin/lumberjackrus)    | 用于将日志记录到本地文件系统（具有日志轮转和每个日志级别一个文件）的钩子 |
| [Mail](https://github.com/zbindenren/logrus_mail)            | 用于通过电子邮件发送异常的钩子                               |
| [Mattermost](https://github.com/shuLhan/mattermost-integration/tree/master/hooks/logrus) | 用于将日志记录到 [Mattermost](https://mattermost.com/) 的钩子 |
| [Mongodb](https://github.com/weekface/mgorus)                | 用于将日志记录到 mongodb 的钩子                              |
| [MongoDB](https://github.com/LyricTian/logrus-mongo-hook)    | 异步 MongoDB 钩子                                            |
| [MongoDB](https://github.com/geronimo794/go-mongolog)        | 用于使用 [MongoDB Go driver](https://www.mongodb.com/docs/drivers/go/current/) 将日志记录到 MongoDB 的钩子 |
| [MySQL](https://github.com/LyricTian/logrus-mysql-hook)      | 异步 MySQL 钩子                                              |
| [NATS-Hook](https://github.com/rybit/nats_logrus_hook)       | 用于将日志记录到 [NATS](https://nats.io/) 的钩子             |
| [New Relic](https://github.com/abrunner94/rusrelic)          | 用于将日志记录到 [New Relic](https://newrelic.com/) 的钩子   |
| [NXLog](https://github.com/hybridtheory/logrus-nxlog-hook)   | 用于将日志记录到 [NXLog](https://nxlog.co/) 的钩子           |
| [Octokit](https://github.com/dorajistyle/logrus-octokit-hook) | 用于通过 octokit 将日志记录到 GitHub 的钩子                  |
| [OpsGenie](https://github.com/JackFazackerley/logrus-opsgenie-hook) | OpsGenie 的钩子                                              |
| [Papertrail](https://github.com/polds/logrus-papertrail-hook) | 通过 UDP 将错误日志发送到托管日志服务 [Papertrail](https://papertrailapp.com/) |
| [PostgreSQL](https://github.com/gemnasium/logrus-postgresql-hook) | 将日志发送到 [PostgreSQL](http://postgresql.org/)            |
| [Promrus](https://github.com/weaveworks/promrus)             | 将日志消息数作为 [Prometheus](https://prometheus.io/) 指标暴露出来的钩子 |
| [Pushover](https://github.com/toorop/logrus_pushover)        | 通过 [Pushover](https://pushover.net/) 发送错误的钩子        |
| [Raygun](https://github.com/squirkle/logrus-raygun-hook)     | 用于将日志记录到 [Raygun.io](http://raygun.io/) 的钩子       |
| [Redactrus](https://github.com/whuang8/redactrus)            | 从日志中删除敏感信息的钩子                                   |
| [Redis-Hook](https://github.com/rogierlommers/logrus-redis-hook) | 用于将日志记录到 ELK 堆栈（通过 Redis）的钩子                |
| [Rollrus](https://github.com/heroku/rollrus)                 | 将错误发送到 Rollbar 的钩子                                  |
| [Rocketrus](https://github.com/miraclesu/rocketrus)          | RocketChat 的钩子                                            |
| [Scribe](https://github.com/sagar8192/logrus-scribe-hook)    | 用于将日志记录到 [Scribe](https://github.com/facebookarchive/scribe) 的钩子 |
| [Sentrus](https://github.com/orandin/sentrus)                | 将错误发送到 [Sentry](https://sentry.io/)（使用最新的 Sentry SDK：`sentry-go`）的钩子 |
| [Sentry](https://github.com/evalphobia/logrus_sentry)        | 将错误发送到 Sentry 错误日志和聚合服务的钩子                 |
| [Seq](https://github.com/nullseed/logruseq)                  | 用于将日志记录到 [Seq](https://getseq.net/) 的钩子           |
| [Slackrus](https://github.com/johntdyer/slackrus)            | Slack 聊天的钩子                                             |
| [Splunk](https://github.com/Franco-Poveda/logrus-splunk-hook) | 将事件发送到 [Splunk](https://www.splunk.com/) 的钩子        |
| [Stackdriver](https://github.com/knq/sdhook)                 | 用于将日志记录到 [Google Stackdriver](https://cloud.google.com/logging/) 的钩子 |
| [Sumologrus](https://github.com/mmarinm/sumologrus)          | 用于将日志记录到 [SumoLogic](https://www.sumologic.com/) 的钩子 |
| [Sumorus](https://github.com/doublefree/sumorus)             | 用于将日志记录到 [SumoLogic](https://www.sumologic.com/) 的钩子 |
| [Syslog](https://github.com/sirupsen/logrus/blob/master/hooks/syslog/syslog.go) | 将错误发送到远程 syslog 服务器的钩子。在幕后使用标准库 `log/syslog`。 |
| [Syslog TLS](https://github.com/shinji62/logrus-syslog-ng)   | 带有 TLS 支持将错误发送到远程 syslog 服务器的钩子            |
| [SQS-Hook](https://github.com/tsarpaul/logrus_sqs)           | 用于将日志记录到 [Amazon Simple Queue Service (SQS)](https://aws.amazon.com/sqs/) 的钩子 |
| [Telegram](https://github.com/rossmcdonald/telegram_hook)    | 将错误日志记录到 [Telegram](https://telegram.org/) 的钩子    |
| [Telegram](https://github.com/krasun/logrus2telegram)        | 将日志发送到 [Telegram](https://telegram.org/) 的钩子        |
| [Tencent Cloud CLS](https://github.com/chuangbo/logruscls)   | 用于将日志记录到 [腾讯云 CLS](https://intl.cloud.tencent.com/document/product/614) 的钩子 |
| [TraceView](https://github.com/evalphobia/logrus_appneta)    | 用于将日志记录到 [AppNeta TraceView](https://www.appneta.com/products/traceview/) 的钩子 |
| [Typetalk](https://github.com/dragon3/logrus-typetalk-hook)  | 用于将日志记录到 [Typetalk](https://www.typetalk.in/) 的钩子 |
| [Vkrus](https://github.com/SevereCloud/vkrus)                | 用于将日志记录到 [VK](https://vk.com/) 的钩子                |
| [Windows Event Log](https://github.com/freman/eventloghook)  | Windows 事件日志的钩子                                       |
| [Yandex Cloud Logging](https://github.com/DavyJohnes/logrus-yc-hoook) | 用于将日志记录到 [Yandex Cloud Logging](https://cloud.yandex.ru/docs/logging/) 的钩子 |
| [DingTalk](https://github.com/exexute/logrus-webhook#send-log-to-dingtalk-robot) | 用于将日志记录到 [钉钉机器人](https://open.dingtalk.com/document/group/call-robot-api-operations) 的钩子 |
| [FeiShu](https://github.com/exexute/logrus-webhook#send-log-to-feishu-webhook) | 用于将日志记录到 [飞书机器人](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bot-v3/bot-overview) 的钩子 |
| [Aliyun SLS](https://github.com/exexute/logrus-webhook#send-log-to-aliyun-sls) | 用于将日志记录到 [阿里云 SLS](https://help.aliyun.com/document_detail/48869.html) 的钩子 |


