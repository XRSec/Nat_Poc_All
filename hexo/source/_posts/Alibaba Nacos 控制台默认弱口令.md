---
title: Alibaba Nacos 控制台默认弱口令
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - Alibaba Nacos
---

# Alibaba Nacos 控制台默认弱口令

## 漏洞描述

Alibaba Nacos 控制台存在默认弱口令 **nacos/nacos**，可登录后台查看敏感信息

## 漏洞影响

> [!NOTE]
>
> Alibaba Nacos

## 漏洞复现

发送如下请求

![](/img/20210924015238596205.png)

返回200说明成功登录

## Goby & POC

> [!NOTE]
>
> 已上传 https://github.com/PeiQi0/PeiQi-WIKI-POC Goby & POC 目录中
>
> Alibaba_Nacos_Default_password.json

![](/img/20210924015238827681.png)