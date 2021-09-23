---
title: 中国移动 禹路由 simple-index.asp 越权访问漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - 中国移动
---

# 中国移动 禹路由 simple-index.asp 越权访问漏洞

## 漏洞描述

中国移动 禹路由 simple-index.asp 存在越权访问漏洞，攻击者通过漏洞可以获取账号密码等敏感信息

## 漏洞影响

> [!NOTE]
>
> 中国移动 禹路由

## FOFA

> [!NOTE]
>
> title="互联世界 物联未来-登录"

## 漏洞复现

登录页面如下

![image-20210618173233203](/img/20210924020314932354.png)

访问Url

```
/simple-index.asp
```

​	![image-20210618173314789](/img/20210924020315091288.png)

越过管理员验证获取Wifl账号密码等信息