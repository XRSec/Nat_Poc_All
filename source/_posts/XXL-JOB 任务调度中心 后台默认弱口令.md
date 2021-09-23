---
title: XXL-JOB 任务调度中心 后台默认弱口令
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - XXL-JOB
---

# XXL-JOB 任务调度中心 后台默认弱口令

## 漏洞描述

XXL-JOB 任务调度中心 后台存在默认弱口令，攻击者可以在后台进一步攻击

## 漏洞影响

> [!NOTE]
>
> XXL-JOB

## FOFA

> [!NOTE]
>
> app="XXL-JOB" || title="任务调度中心"

## 漏洞复现

使用默认口令登录 admin 123456

![](/img/20210924015222611214.png)

![](/img/20210924015224157826.png)

## Goby & POC

> [!NOTE]
>
> 已上传 https://github.com/PeiQi0/PeiQi-WIKI-POC Goby & POC 目录中
>
> XXL_JOB_Default_password

![](/img/20210924015224446513.png)