---
title: 用友 NC bsh.servlet.BshServlet 远程命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - OA产品漏洞
 - 用友OA
---

# 用友 NC bsh.servlet.BshServlet 远程命令执行漏洞

## 漏洞描述

用友 NC bsh.servlet.BshServlet 存在远程命令执行漏洞，通过BeanShell 执行远程命令获取服务器权限

## 漏洞影响

> [!NOTE]
>
> 用友 NC 

## FOFA

> [!NOTE]
>
> icon_hash="1085941792"

## 漏洞复现

访问页面如下

![image-20210531220356962](/img/20210924020222153611.png)

漏洞Url为

```
/servlet/~ic/bsh.servlet.BshServlet
```

![image-20210531220503672](/img/20210924020223925833.png)