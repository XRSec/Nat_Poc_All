# VERSION：HEXO + Wikitten + CHEVERETO CHEVERETO

## Poc 文章申请方式

- 对本仓库的poc进行同步更新，保证poc更新的稳定性
- issus 申请内测
- 由于近期特别忙，所以兄弟们帮忙一起维护这个项目了
- 祝大家早日建立属于自己的通用内网poc平台！

## 预览

![image-20220221234444368](https://xrsec.s3.bitiful.net/IMG/image-20220221234444368.png)

![image-20220221234545002](https://xrsec.s3.bitiful.net/IMG/image-20220221234545002.png)

![image-20220221234602514](https://xrsec.s3.bitiful.net/IMG/image-20220221234602514.png)

## PeiQi文库整理记录

> Peggy 0720; 336 articles; 113 zip | py; 1164 image results

### Markdown article first correction, classification optimization

```python
#!/usr/bin/env python3
# _*_ coding: utf-8 _*_

import os, re
from shutil import copyfile

## Poc 文章申请方式

- 对本仓库的poc进行同步更新，保证poc更新的稳定性
- issus 申请内测
- 由于近期特别忙，所以兄弟们帮忙一起维护这个项目了
- 祝大家早日建立属于自己的通用内网poc平台！
- 希望兄弟们能为 [dev](https://github.com/XRSec/Nat_Poc_All/tree/dev) 版本贡献一些代码

## 预览

if __name__ == '__main__':
    filepath = "/PeiQi_Wiki/test/"
    ds_store(filepath)
    g = os.walk(filepath)
    for path, dir_list, file_list in g:
        for file_name in file_list:
            if (".md" in os.path.join(path, file_name)) and ("README" not in os.path.join(path, file_name)) and ("peiqi.py" not in os.path.join(path, file_name)):
                l1 = (os.path.join(path).replace("/PeiQi_Wiki/test/", "")).split("/")
                categories1, categories2= l1[0],l1[1]
                date = "---\ntitle: " + file_name.replace(".md","") + "\ndate: 2021-09-23 23:55:51\ntags: PeiQi文库\ncategories:\n - "+ categories1 + "\n - "+ categories2 +"\n---\n\n"
                # print(date)
                open("/PeiQi_Wiki/test1/" + file_name, "w").close()
                with open("/PeiQi_Wiki/test1/" + file_name, "w") as hexo:
                    hexo.write(date + open(os.path.join(path, file_name),"r").read())
            if (".zip" in os.path.join(path, file_name)) or ("py" in os.path.join(path, file_name)):
                copyfile(os.path.join(path, file_name), "/PeiQi_Wiki/hexo/source/Poc/" + file_name)
```

### Picture update address

> First replace image link interference with `vscode`

`Pycharm` was used to match 1061 undescribed images

```python
#!/usr/bin/env python3
# _*_ coding: utf-8 _*_

import os, re, datetime, sys, requests
from shutil import copyfile


def ds_store(path):
    ds_store_file = path + ".DS_Store"
    while os.path.exists(ds_store_file):
        os.remove(ds_store_file)
        print("delete .DS_Store")
    return


def imgreplace(file, old_content, new_content):
    with open(file, encoding='UTF-8') as f:
        content = f.read()
        f.close()
    content = content.replace(old_content, new_content)
    with open(file, 'w', encoding='UTF-8') as f:
        f.write(content)
        f.close()


if __name__ == '__main__':
    filepath = "/PeiQi_Wiki/test1/"
    ds_store(filepath)
    file = os.listdir(filepath)
    for i in range(0, len(file)):
        file_content = open(filepath + file[i], "r").read()
        # imgs = re.findall('(/resource/(.*?).(png|jpg|jpeg|gif|webp))', file_content, re.S)
        imgs = re.findall("(?:!\[*\]\((.*?)\))", file_content, re.S)
        for j in range(0, len(imgs)):
            if len(imgs) > 0:
                new_name = datetime.datetime.now().strftime("%Y%m%d%H%M%S%f") + "." + re.match(r"^[\s\S]*\.(jpg|png|webp|jpeg|gif)", imgs[j]).group(1)
                newimg_path = "/PeiQi_Wiki/hexo/source/img/" + new_name
                with open(newimg_path, "wb") as temp:
                    temp.write((requests.get(imgs[j], timeout=5, verify=False)).content)
                newimg = "http://img.chion.tech/img/" + new_name
                mdpath = filepath + file[i]
                print(imgs[j])
                imgreplace(mdpath,imgs[j],newimg)
                print("1")
```

Of course, there was a little hiccup, which took me an hour to fix

So I will send out the `Hexo` version of `PeiQi Wiki` with images

## Go程序新需求

判断接收的是文件还是网址，如果是网址，则：

- 自动下载
- 判断文件类型
- 自动重命名

<hr>

## 往期需求

### 添加信息

> 当识别到文件后缀是.conf时自动在文件头部添加 [ title data tag ]等信息，文章尾部添加版权声明
>
> 由于网上poc很多，如不做细致的tag 区分，则可以把原作者附上，例如tag: peiqi
>
> 希望能够在保存新文件的时候将文件名中的特殊字符做处理，`space` 转化成 `_`， `<>` 转换成 `-` ，并且自动检测原文件内容是否包含 `title` `date` `tag` 关键字，以防冲突

- Input

```bash
# 启动
docker-compose -f /www/chevereto/docker-compose.yaml up

# 停止
docker-compose -f /www/chevereto/docker-compose.yaml down
```

### 配置

```ini
host img.chion.tech 虚拟机IP
host poc.chion.tech 虚拟机IP
```

### 图片上传

> Typora & Python3

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
    shutil.dogefs.s3ree(path_tmp)
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

> 声明：维护不易，请更新文章时向我发送（图片 数据库 文章）用以同步漏洞库
