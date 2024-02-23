+++
title = "App Engine 1.7.1中的 Go 更新"
weight = 2
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go updates in App Engine 1.7.1  - App Engine 1.7.1中的 Go 更新

> 原文：[https://go.dev/blog/appengine-171](https://go.dev/blog/appengine-171)

Andrew Gerrand
22 August 2012

This week we released version 1.7.1 of the App Engine SDK. It includes some significant updates specific to the App Engine runtime for Go.

本周我们发布了App Engine SDK的1.7.1版本。它包括一些专门针对Go的App Engine运行时的重要更新。

The [memcache package](https://developers.google.com/appengine/docs/go/memcache/reference) has had some additions to its [Codec](https://developers.google.com/appengine/docs/go/memcache/reference#Codec) convenience type. The SetMulti, AddMulti, CompareAndSwap, and CompareAndSwapMulti methods make it easier to store and update encoded data in the [Memcache Service](https://developers.google.com/appengine/docs/go/memcache/overview).

memcache包对其Codec便利类型进行了一些补充。SetMulti、AddMulti、CompareAndSwap和CompareAndSwapMulti方法使得在Memcache服务中存储和更新编码数据更加容易。

The [bulkloader tool](https://developers.google.com/appengine/docs/go/tools/uploadingdata) can now be used with Go apps, allowing users to upload and download datastore records in bulk. This is useful for backups and offline processing, and a great help when migrating Python or Java apps to the Go runtime.

批量加载器工具现在可以与Go应用程序一起使用，允许用户批量上传和下载数据存储记录。这对于备份和离线处理非常有用，在将Python或Java应用程序迁移到Go运行时也有很大帮助。

The [Images Service](https://developers.google.com/appengine/docs/go/images/overview) is now available to Go users. The new [appengine/image package](https://developers.google.com/appengine/docs/go/images/reference) supports serving images directly from Blobstore and resizing or cropping those images on the fly. Note that this is not the full image service as provided by the Python and Java SDKs, as much of the equivalent functionality is available in the [standard Go image package](https://go.dev/pkg/image/) and external packages such as [graphics-go](http://code.google.com/p/graphics-go/).

图像服务现在可以提供给Go用户。新的appengine/image包支持直接从Blobstore提供图片，并对这些图片进行实时调整或裁剪。请注意，这并不是Python和Java SDK所提供的完整图像服务，因为标准Go图像包和外部包（如graphics-go）中提供了许多同等功能。

The new [runtime.RunInBackground](https://developers.google.com/appengine/docs/go/backends/runtime#RunInBackground) function allows backend requests to spawn a new request independent of the initial request. These can run in the background as long as the backend stays alive.

新的runtime.RunInBackground函数允许后端请求产生一个独立于初始请求的新请求。只要后端保持活力，这些都可以在后台运行。

Finally, we have filled in some missing functionality: the [xmpp package](https://developers.google.com/appengine/docs/go/xmpp/reference) now supports sending presence updates and chat invitations and retrieving the presence state of another user, and the [user package](https://developers.google.com/appengine/docs/go/users/reference) supports authenticating clients with OAuth.

最后，我们填补了一些缺失的功能：xmpp包现在支持发送存在感更新和聊天邀请，并检索另一个用户的存在感状态，而用户包支持用OAuth验证客户端。

You can grab the new SDK from the [App Engine downloads page](https://developers.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go) and browse the [updated documentation](https://developers.google.com/appengine/docs/go).

您可以从App Engine的下载页面获取新的SDK，并浏览更新的文档。
