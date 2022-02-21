---
title: Alibaba Canal config 云密钥信息泄露漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - Alibaba Canal
---

# Alibaba Canal config 云密钥信息泄露漏洞

## 漏洞描述

由于/api/v1/canal/config  未进行权限验证可直接访问，导致账户密码、accessKey、secretKey等一系列敏感信息泄露

## 漏洞影响

> [!NOTE]
>
> Alibaba Canal

## FOFA

> [!NOTE]
>
> title="Canal Admin"

## 漏洞复现

验证漏洞的Url为

```
/api/v1/canal/config/1/0
```

![](/img/20210924015211018043.png)

其中泄露了 aliyun.access 密钥，可以控制密钥下的所有服务器

云密钥泄露参考: red.peiqi.tech

其中还含有默认口令 **admin/123456**

![](/img/20210924015212406299.png)