---
title: 泛微OA sysinterfacecodeEdit.jsp 页面任意文件上传
date: 2021-09-23 23:55:51
tags: PeiQi文库
categories:
 - OA产品漏洞
 - 泛微OA
---

# 泛微OA sysinterface/codeEdit.jsp 页面任意文件上传 WooYun-2015-0155705

## 漏洞描述

泛微OA sysinterface/codeEdit.jsp 页面任意文件上传导致可以上传恶意文件

## 漏洞描述

> [!NOTE]
>
> 较老版本，目前无准确版本

## 漏洞复现

`filename=******5308.java&filetype=javafilename为文件名称 为空时会自动创建一个`

```java
String fileid = "Ewv";<br>
    String readonly = "";<br>
    boolean isCreate = false;<br>
    if(StringHelper.isEmpty(fileName)) {<br>
     Date ndate = new Date();<br>
     SimpleDateFormat sf = new SimpleDateFormat("yyyyMMddHHmmss");<br>
     String datetime = sf.format(ndate);<br>
     fileid = fileid + datetime;<br>
     fileName= fileid + "." + filetype;<br>
     isCreate = true;<br>
    } else {<br>
        int pointIndex = fileName.indexOf(".");<br>
        if(pointIndex > -1) {<br>
            fileid = fileName.substring(0,pointIndex);<br>
        }}
```

![](/img/20210924013826310029.png)

![](/img/20210924013827143080.png)

![](/img/20210924013827844345.png)

![](/img/20210924013828039294.png)

![](/img/20210924013828340849.png)

## 参考文章

[泛微OA未授权可导致GetShell](https://www.uedbox.com/post/15730/)