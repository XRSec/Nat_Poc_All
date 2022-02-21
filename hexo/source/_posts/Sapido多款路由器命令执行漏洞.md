---
title: Sapido多款路由器命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - Sapido
---

# Sapido多款路由器命令执行漏洞

## 漏洞描述

Sapido多款路由器在未授权的情况下，导致任意访问者可以以Root权限执行命令

## 漏洞影响

> [!NOTE]
>
> BR270n-v2.1.03
>
> BRC76n-v2.1.03
>
> GR297-v2.1.3
>
> RB1732-v2.0.43

## FOFA

> [!NOTE]
>
> app="Sapido-路由器"

## 漏洞复现

固件中存在一个asp文件为 **syscmd.asp** 存在命令执行

![](/img/20210924015231234217.png)

访问目标:

```
http://xxx.xxx.xxx.xxx/syscmd.asp
http://xxx.xxx.xxx.xxx/syscmd.htm
```

![](/img/20210924015231590350.png)

直接输入就可以命令执行了

