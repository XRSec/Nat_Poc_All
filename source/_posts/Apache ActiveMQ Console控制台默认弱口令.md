---
title: Apache ActiveMQ Console控制台默认弱口令
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web服务器漏洞
 - Apache
---

# Apache ActiveMQ Console 存在默认弱口令

## 漏洞描述

Apache ActiveMQ Console 存在默认弱口令 admin:admin，进入控制台后可被进一步恶意利用

## 漏洞影响

> [!NOTE]
>
> Apache ActiveMQ

## 漏洞复现

Apache ActiveMQ 默认开启了 8186 控制台

访问目标: http://xxx.xxx.xxx.xxx:8161/admin

![](/img/20210924015530749634.png)

使用默认口令 admin:admin

![](/img/20210924015530907761.png)

## Goby & POC

> [!NOTE]
>
> Apache ActiveMQ_Console Weak Password

![](/img/20210924015531156427.png)

![](/img/20210924015531439855.png)