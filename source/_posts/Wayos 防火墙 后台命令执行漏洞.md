---
title: Wayos 防火墙 后台命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - WayOS
---

# Wayos 防火墙 后台命令执行漏洞

## 漏洞描述

Wayos 防火墙 后台存在命令执行漏洞，通过命令注入可以执行远程命令

## 漏洞影响

> [!NOTE]
>
> Wayos 防火墙 

## FOFA

> [!NOTE]
>
> body="Get_Verify_Info(hex_md5(user_string)."

## 漏洞复现

登录页面如下

![image-20210531224015364](/img/20210924020226801539.png)

登录后台后 ping 模块命令执行

![image-20210531224745218](/img/20210924020227162540.png)



