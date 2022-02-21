---
title: 极致CMS 1.81以下版本 存储型XSS
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - CMS漏洞
 - 极致CMS
---

# 极致CMS 1.81以下版本 存储型XSS

### 漏洞复现

登录管理员添加模块

![](/img/20210924013831372462.png)

注册用户

![](/img/20210924013831621087.png)

点击发布文章

![](/img/20210924013831855621.png)

在文章标题处插入xss payload

```<details open ontoggle= confirm(document[`coo`+`kie`])>```

当管理员访问时XSS成功

![](/img/20210924013832084155.png)



### 参考

[极致CMS代码审计](https://xz.aliyun.com/t/7861)