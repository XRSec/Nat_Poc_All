# VERSION：HEXO + Wikitten + Go_UPDATE GOLANG

前言：由于[上一版方案](https://github.com/XRSec/Nat_Poc_All/tree/v0.01)过于臃肿，所以决定重新整理，刚好最近golang挺火的，就当练个手

## Poc 文章申请方式

- 对本仓库的poc进行同步更新，保证poc更新的稳定性
- issus 申请内测
- 由于近期特别忙，所以兄弟们帮忙一起维护这个项目了
- 祝大家早日建立属于自己的通用内网poc平台！
- 希望兄弟们能为 dev 版本贡献一些代码

## 预览

![image-20220221234444368](https://rmt.ladydaily.com/fetch/ZYGG/storage/image-20220221234444368.png)

![image-20220221234545002](https://rmt.ladydaily.com/fetch/ZYGG/storage/image-20220221234545002.png)

![image-20220221234602514](https://rmt.ladydaily.com/fetch/ZYGG/storage/image-20220221234602514.png)

## 架构

web：hexo ( theme [Wikitten](https://github.com/zthxxx/hexo-theme-Wikitten))

图床：图床与笔记同步整合

```ini
➜  Nat_Poc_All git:(main) ✗ tree -N
.
├── images ( 图像上传文件夹 )
│   └── 20210922113701.png
├── main.go ( 主程序 )
├── markdown ( Markdown上传文件夹，同 hexo /hexo/source/_posts/ )
│   └── 任务.md
├── upd ( 临时文件夹 )
│   ├── poc.png
│   └── 任务.md
├── upload.gtpl ( web页面渲染 )
└── upload.py ( 上传测试脚本，内容需要修改 )

3 directories, 7 files
```

## 需求

### 添加信息

> 当识别到文件后缀是.conf时自动在文件头部添加 [ title data tag ]等信息，文章尾部添加版权声明
>
> 由于网上poc很多，如不做细致的tag 区分，则可以把原作者附上，例如tag: peiqi
>
> 希望能够在保存新文件的时候将文件名中的特殊字符做处理，`space` 转化成 `_`， `<>` 转换成 `-` ，并且自动检测原文件内容是否包含 `title` `date` `tag` 关键字，以防冲突

- 快速生成 tag 等信息 

```bash
#!/bin/zsh

posts="/markdown/"
title=$(basename $1 .md)
newdoc=$posts$title.md

echo "---\ntitle: $title\ndate: $2 $3\ntag: \n---\n\n" >> $newdoc
cat $1 >> $newdoc
echo "\n> Poc++ has the right to modify and interpret this article.
```
- 将本地写好的markdown文件复制到指定位置

```ini
if file == /xxxx/xxxx/test.md
$1 == /xxxx/xxxx/test.md
basename $1 .md == test
newdoc=$posts$title.md == /markdown/test.md
```

### 页面美化

> 虽然说所有的步骤可以直接通过请求解决，但是还是需要美化一下web上传页面，理论上同html css

### 图文上传

Typora 是一个很好的写作平台，同时提供图片上传选项，所以图片上传是使用的 `python 脚本` 或者 `go 脚本`

文件上传主要上传Markdown文件，总不可能每次都 `docker cp` 或者 `scp`  吧，再部属个 `FTP` 显示太臃肿，刚好主程序已经能够接收 `Markdown` 文件,所以只需要写一个上传文件的脚本

so - 图文脚本合并，判断后缀是个不错的选择,这个脚本是 `v0.0.1` 的 `chevereto` 上传脚本，希望兄弟能改成`golang` 语言

```python
#!/usr/bin/env python3
# _*_ coding: utf-8 _*_
import shutil

import requests, sys, base64, os, re, urllib3, imghdr, tempfile, datetime, urllib3

# 创建临时文件夹 (防止空格导致报错)

path_tmp = tempfile.gettempdir() + '/typora/'
# 判断结果
if not os.path.exists(path_tmp):
    os.makedirs(path_tmp)
else:
    shutil.rmtree(path_tmp)
    os.mkdir(path_tmp)


def upload_img():
    for i in range(1, len(sys.argv)):
        if os.path.isfile(sys.argv[i]):
            new_name = path_tmp + datetime.datetime.now().strftime("%Y%m%d%H%M%S%f") + "." + imghdr.what(
                sys.argv[i])  # 时间戳重命名
            shutil.move(sys.argv[i], new_name)
            img_file = base64.b64encode(open(new_name, 'rb').read())
        elif "http://" or "https://" in sys.argv[i]:
            img_file = sys.argv[i]
            print("error")
        data = {
            "source": img_file,
            "url": 'http://img.chion.tech/api/1/upload',  # 请求URL
            "action": "upload",
            "key": "APIKEY",  # API v1 的密钥
            "format": 'json'
        }
        res = requests.post(data['url'], data=data).json()
        if str(res["status_code"]) == "200":
            if res["image"]["url"]:
                print(res["image"]["url"])
            else:
                print("日常抽风！")
                webbrowser.open("http://img.chion.tech/explore/recent", new=0, autoraise=True)
            shutil.remove(sys.argv[i])
            shutil.remove(new_name)
        elif "Duplicated upload" in str(res):
            print("Duplicated upload")
            webbrowser.open("http://img.chion.tech/explore/recent", new=0, autoraise=True)
            shutil.move(new_name, sys.argv[i])
        else:
            print("img_file=", sys.argv[i], res)
            shutil.move(new_name, sys.argv[i])


if __name__ == '__main__':
    upload_img()
```

## 使用

### 图片

- `Typora`  **OR** 

```bash
$ curl localhost:8081/upload -X POST -F "uploadFile=@/Nat_Poc_All/upd/poc.png" -x "http://localhost:8080"

# unix
$ echo 'curl localhost:8081/upload -X POST -F "uploadFile=@/$@" -x "http://localhost:8080"' > natpoc && chmod +x natpoc && ./natpoc '/Nat_Poc_All/markdown/任务.md'
# windows
$ echo 'curl localhost:8081/upload -X POST -F "uploadFile=@/%*" -x "http://localhost:8080"' > natpoc && chmod +x natpoc && ./natpoc '/Nat_Poc_All/markdown/任务.md'
```

### 域名

```ini
# 待定
img.chion.tech 虚拟机IP
poc.chion.tech 虚拟机IP
```

## 同步

只需要将 `Markdown` 和图片一起打包就好了，图片使用的是时间戳，可以用 `vscode`  将 `Markdown` 中的图片链接提取并下载，其他的就不用改了

## 测试

```bash
curl localhost:8081/upload -X POST -F "uploadFile=@/$@" -x "http://localhost:8080"
```

