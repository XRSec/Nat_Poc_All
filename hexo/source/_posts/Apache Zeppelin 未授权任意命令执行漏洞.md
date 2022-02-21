---
title: Apache Zeppelin 未授权任意命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web服务器漏洞
 - Apache
---

# Apache Zeppelin 未授权任意命令执行漏洞

## 漏洞描述

Apache Zeppelin 存在未授权的用户访问命令执行接口，导致了任意用户都可以执行恶意命令获取服务器权限

## 漏洞影响

> [!NOTE]
>
> Apache Zeppelin

## FOFA

> [!NOTE]
>
> icon_hash="960250052"

## 漏洞复现

含有漏洞的页面如下

![](/img/20210924015238099051.png)

点击 创建一个匿名用户在用户页面执行命令即可

![](/img/20210924015238354004.png)

