+++
title = "hooks"
date = 2023-06-25T09:41:29+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Hooks

https://github.com/sirupsen/logrus/wiki/Hooks

Owefsad edited this page on Feb 1 · [35 revisions](https://github.com/sirupsen/logrus/wiki/Hooks/_history)



This page describes the list of known hooks services. The list is provided as is. They are not scrutinized nor reviewed by logrus developers.

| Hook                                                         | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [Airbrake "legacy"](https://github.com/gemnasium/logrus-airbrake-legacy-hook) | Send errors to an exception tracking service compatible with the Airbrake API V2. Uses [`airbrake-go`](https://github.com/tobi/airbrake-go) behind the scenes. |
| [Airbrake](https://github.com/gemnasium/logrus-airbrake-hook) | Send errors to the Airbrake API V3. Uses the official [`gobrake`](https://github.com/airbrake/gobrake) behind the scenes. |
| [Amazon Kinesis](https://github.com/evalphobia/logrus_kinesis) | Hook for logging to [Amazon Kinesis](https://aws.amazon.com/kinesis/) |
| [Amazon SNS](https://github.com/stvvan/logrus-sns)           | Hook for logging to [Amazon Simple Notification Service (SNS)](https://aws.amazon.com/sns/) |
| [Amazon EventBridge](https://github.com/teddy-schmitz/eventbridge_logrus) | Hook for logging to [Amazon EventBridge](https://aws.amazon.com/eventbridge/) |
| [Amqp-Hook](https://github.com/vladoatanasov/logrus_amqp)    | Hook for logging to Amqp broker (Like RabbitMQ)              |
| [Anexia CloudLog](https://github.com/anexia-it/go-logrus-cloudlog) | Hook for logging to Anexia CloudLog                          |
| [Application Insights](https://github.com/jjcollinge/logrus-appinsights) | Hook for logging to [Application Insights](https://azure.microsoft.com/en-us/services/application-insights/) |
| [AzureTableHook](https://github.com/kpfaulkner/azuretablehook/) | Hook for logging to Azure Table Storage                      |
| [Bugsnag](https://github.com/Shopify/logrus-bugsnag/blob/master/bugsnag.go) | Send errors to the Bugsnag exception tracking service.       |
| [ClickHouse](https://github.com/oxgrouby/logrus-clickhouse-hook) | Send logs to [ClickHouse](https://clickhouse.yandex/)        |
| [Datadog Log](https://github.com/bin3377/logrus-datadog-hook) | Hook for logging to [Datadog](https://www.datadoghq.com/) over HTTP endpoint |
| [Discord Bot Hook](https://github.com/outdead/discordbotrus) | Hook for logging to [Discord](https://discordapp.com/) using Discord app |
| [Discordrus](https://github.com/kz/discordrus)               | Hook for logging to [Discord](https://discordapp.com/)       |
| [Elastic APM](https://godoc.org/go.elastic.co/apm/module/apmlogrus#Hook) | Hook for logging errors to [Elastic APM](https://www.elastic.co/solutions/apm) |
| [ElasticSearch](https://github.com/sohlich/elogrus)          | Hook for logging to ElasticSearch                            |
| [ElasticSearch (with the official client)](https://github.com/go-extras/elogrus) | Hook for logging to ElasticSearch                            |
| [Firehose](https://github.com/beaubrewer/logrus_firehose)    | Hook for logging to [Amazon Firehose](https://aws.amazon.com/kinesis/firehose/) |
| [Fluentd](https://github.com/evalphobia/logrus_fluent)       | Hook for logging to fluentd                                  |
| [Go-Slack](https://github.com/multiplay/go-slack)            | Hook for logging to [Slack](https://slack.com/)              |
| [Graylog](https://github.com/gemnasium/logrus-graylog-hook)  | Hook for logging to [Graylog](http://graylog2.org/)          |
| [Hiprus](https://github.com/nubo/hiprus)                     | Send errors to a channel in hipchat.                         |
| [Honeybadger](https://github.com/agonzalezro/logrus_honeybadger) | Hook for sending exceptions to Honeybadger                   |
| [InfluxDB](https://github.com/Abramovic/logrus_influxdb)     | Hook for logging to influxdb                                 |
| [Influxus](http://github.com/vlad-doru/influxus)             | Hook for concurrently logging to [InfluxDB](http://influxdata.com/) |
| [Journalhook](https://github.com/wercker/journalhook)        | Hook for logging to `systemd-journald`                       |
| [KafkaLogrus](https://github.com/tracer0tong/kafkalogrus)    | Hook for logging to Kafka                                    |
| [Kafka REST Proxy](https://github.com/Nordstrom/logrus-kafka-rest-proxy) | Hook for logging to [Kafka REST Proxy](https://docs.confluent.io/current/kafka-rest/docs) |
| [LFShook](https://github.com/rifflock/lfshook)               | Hook for logging to the local filesystem                     |
| [Logbeat](https://github.com/macandmia/logbeat)              | Hook for logging to [Opbeat](https://opbeat.com/)            |
| [Logentries](https://github.com/jcftang/logentriesrus)       | Hook for logging to [Logentries](https://logentries.com/)    |
| [Logentrus](https://github.com/puddingfactory/logentrus)     | Hook for logging to [Logentries](https://logentries.com/)    |
| [Logmatic.io](https://github.com/logmatic/logmatic-go)       | Hook for logging to [Logmatic.io](http://logmatic.io/)       |
| [Logrus Boltdb Hook](https://github.com/trK54Ylmz/logrus-boltdb-hook) | Hook for logging to boltdb                                   |
| [Logrus Bolt Hook](https://github.com/kennykarnama/logrus-bolt-hook) | Hook for logging to boltdb                                   |
| [Logrusly](https://github.com/sebest/logrusly)               | Send logs to [Loggly](https://www.loggly.com/)               |
| [Logstash](https://github.com/bshuster-repo/logrus-logstash-hook) | Hook for logging to [Logstash](https://www.elastic.co/products/logstash) |
| [Loki](https://github.com/YuKitsune/lokirus)                 | Hook for logging to [Loki](https://grafana.com/oss/loki/)    |
| [Lumberjackrus](https://github.com/orandin/lumberjackrus)    | Hook for logging to the local filesystem (with logrotate and a file per log level) |
| [Mail](https://github.com/zbindenren/logrus_mail)            | Hook for sending exceptions via mail                         |
| [Mattermost](https://github.com/shuLhan/mattermost-integration/tree/master/hooks/logrus) | Hook for logging to [Mattermost](https://mattermost.com/)    |
| [Mongodb](https://github.com/weekface/mgorus)                | Hook for logging to mongodb                                  |
| [MongoDB](https://github.com/LyricTian/logrus-mongo-hook)    | An asynchronous MongoDB Hook                                 |
| [MongoDB](https://github.com/geronimo794/go-mongolog)        | Hook for logging to MongoDB with [MongoDB Go driver](https://www.mongodb.com/docs/drivers/go/current/) |
| [MySQL](https://github.com/LyricTian/logrus-mysql-hook)      | An asynchronous MySQL Hook                                   |
| [NATS-Hook](https://github.com/rybit/nats_logrus_hook)       | Hook for logging to [NATS](https://nats.io/)                 |
| [New Relic](https://github.com/abrunner94/rusrelic)          | Hook for logging to [New Relic](https://newrelic.com/)       |
| [NXLog](https://github.com/hybridtheory/logrus-nxlog-hook)   | Hook for logging to [NXLog](https://nxlog.co/)               |
| [Octokit](https://github.com/dorajistyle/logrus-octokit-hook) | Hook for logging to github via octokit                       |
| [OpsGenie](https://github.com/JackFazackerley/logrus-opsgenie-hook) | Hook for OpsGenie.                                           |
| [Papertrail](https://github.com/polds/logrus-papertrail-hook) | Send errors to the [Papertrail](https://papertrailapp.com/) hosted logging service via UDP. |
| [PostgreSQL](https://github.com/gemnasium/logrus-postgresql-hook) | Send logs to [PostgreSQL](http://postgresql.org/)            |
| [Promrus](https://github.com/weaveworks/promrus)             | Expose number of log messages as [Prometheus](https://prometheus.io/) metrics |
| [Pushover](https://github.com/toorop/logrus_pushover)        | Send error via [Pushover](https://pushover.net/)             |
| [Raygun](https://github.com/squirkle/logrus-raygun-hook)     | Hook for logging to [Raygun.io](http://raygun.io/)           |
| [Redactrus](https://github.com/whuang8/redactrus)            | Redact sensitive information from your logs                  |
| [Redis-Hook](https://github.com/rogierlommers/logrus-redis-hook) | Hook for logging to a ELK stack (through Redis)              |
| [Rollrus](https://github.com/heroku/rollrus)                 | Hook for sending errors to rollbar                           |
| [Rocketrus](https://github.com/miraclesu/rocketrus)          | Hook for RocketChat.                                         |
| [Scribe](https://github.com/sagar8192/logrus-scribe-hook)    | Hook for logging to [Scribe](https://github.com/facebookarchive/scribe) |
| [Sentrus](https://github.com/orandin/sentrus)                | Send errors to [Sentry](https://sentry.io/) (using the latest Sentry SDK: `sentry-go`) |
| [Sentry](https://github.com/evalphobia/logrus_sentry)        | Send errors to the Sentry error logging and aggregation service. |
| [Seq](https://github.com/nullseed/logruseq)                  | Hook for logging to [Seq](https://getseq.net/)               |
| [Slackrus](https://github.com/johntdyer/slackrus)            | Hook for Slack chat.                                         |
| [Splunk](https://github.com/Franco-Poveda/logrus-splunk-hook) | Hook for sending events to [Splunk](https://www.splunk.com/) |
| [Stackdriver](https://github.com/knq/sdhook)                 | Hook for logging to [Google Stackdriver](https://cloud.google.com/logging/) |
| [Sumologrus](https://github.com/mmarinm/sumologrus)          | Hook for logging to [SumoLogic](https://www.sumologic.com/)  |
| [Sumorus](https://github.com/doublefree/sumorus)             | Hook for logging to [SumoLogic](https://www.sumologic.com/)  |
| [Syslog](https://github.com/sirupsen/logrus/blob/master/hooks/syslog/syslog.go) | Send errors to remote syslog server. Uses standard library `log/syslog` behind the scenes. |
| [Syslog TLS](https://github.com/shinji62/logrus-syslog-ng)   | Send errors to remote syslog server with TLS support.        |
| [SQS-Hook](https://github.com/tsarpaul/logrus_sqs)           | Hook for logging to [Amazon Simple Queue Service (SQS)](https://aws.amazon.com/sqs/) |
| [Telegram](https://github.com/rossmcdonald/telegram_hook)    | Hook for logging errors to [Telegram](https://telegram.org/) |
| [Telegram](https://github.com/krasun/logrus2telegram)        | Hook for sending logs to [Telegram](https://telegram.org/)   |
| [Tencent Cloud CLS](https://github.com/chuangbo/logruscls)   | Hook for [Tencent Cloud CLS](https://intl.cloud.tencent.com/document/product/614) ([zh-CN](https://cloud.tencent.com/document/product/614)) |
| [TraceView](https://github.com/evalphobia/logrus_appneta)    | Hook for logging to [AppNeta TraceView](https://www.appneta.com/products/traceview/) |
| [Typetalk](https://github.com/dragon3/logrus-typetalk-hook)  | Hook for logging to [Typetalk](https://www.typetalk.in/)     |
| [Vkrus](https://github.com/SevereCloud/vkrus)                | Hook for logging to [VK](https://vk.com/)                    |
| [Windows Event Log](https://github.com/freman/eventloghook)  | Hook for Windows Event Log                                   |
| [Yandex Cloud Logging](https://github.com/DavyJohnes/logrus-yc-hoook) | Hook for logging to [Yandex Cloud Logging](https://cloud.yandex.ru/docs/logging/) |
| [DingTalk](https://github.com/exexute/logrus-webhook#send-log-to-dingtalk-robot) | Hook for logging to [DingTalk Rebot](https://open.dingtalk.com/document/group/call-robot-api-operations) |
| [FeiShu](https://github.com/exexute/logrus-webhook#send-log-to-feishu-webhook) | Hook for logging to [FeiShu Rebot](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bot-v3/bot-overview) |
| [Aliyun SLS](https://github.com/exexute/logrus-webhook#send-log-to-aliyun-sls) | Hook for logging to [Aliyun SLS](https://help.aliyun.com/document_detail/48869.html) |