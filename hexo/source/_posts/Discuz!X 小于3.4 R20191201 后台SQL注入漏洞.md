---
title: Discuz!X 小于3.4 R20191201 后台SQL注入漏洞
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - CMS漏洞
 - Discuz!X
---

# Discuz!X <3.4 R20191201 后台SQL注入漏洞

## 漏洞描述

不久以前Discuz!X的后台披露了一个sql注入的漏洞，这里也要感谢漏洞的发现和研究者（无糖的kn1f3)。

## 影响版本

> [!NOTE]
>
> Discuz!X <3.4 R20191201 版本

## 环境搭建

[百度云盘下载链接](https://pan.baidu.com/s/1qcxgSp20tVGQ3oqts-kNTA)

**密码: 0515**

将 **upload**目录下的文件拷入**phpstudy**下的WWW目录打开网站按照步骤安装就行了

![](/img/20210924015650964381.png)

![](/img/20210924015651651032.png)

## 漏洞复现

来到后台页面, 在 **UCenter 应用 ID** 位置的参数添加单引号并抓包

![](/img/20210924015651859183.png)

发现出现SQL语句报错

![](/img/20210924015652094885.png)

使用报错注入去获取版本号

![](/img/20210924015652329747.png)

这里的参数为 `settingnew[uc][appid]`

查看文件 **\source\admincp\admincp_setting.php**， 在2677行找到了输入点

![](/img/20210924015652620957.png)

根据报错语句找到SQL语句执行点，在文件**uc_client\model\base.php** 中的 206行

![](/img/20210924015653101336.png)

通过这里的语句可以看到我们可以使用 **union注入** 的方法来写入恶意文件(**secure_file_priv不能为Null**)

![](/img/20210924015653504469.png)

```
1' union select "<?php phpinfo();?>"  into outfile 'D:/test.php';--+
```

也可以使用其他的方法

