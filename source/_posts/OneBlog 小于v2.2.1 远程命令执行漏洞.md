---
title: OneBlog 小于v2.2.1 远程命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - OneBlog
---

# OneBlog 小于v2.2.1 远程命令执行漏洞

## 漏洞描述

OneBlog 小于v2.2.1 由于使用含有漏洞版本的Apache Shiro和默认的密钥导致存在远程命令执行漏洞

## 漏洞影响

> [!NOTE]
>
> OneBlog <= v2.2.1

## FOFA

> [!NOTE]
>
> app="OneBlog开源博客后台管理系统"

## 漏洞复现

登陆页面如下

![](/img/20210924015321272679.png)

使用工具直接利用Apache Shiro漏洞即可

![](/img/20210924015321518546.png)