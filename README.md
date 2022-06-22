## VERSION：HEXO + Wikitten + CHEVERETO MAIN

## Poc 文章申请方式

- 对本仓库的poc进行同步更新，保证poc更新的稳定性
- issus 申请内测
- 由于近期特别忙，所以兄弟们帮忙一起维护这个项目了
- 祝大家早日建立属于自己的通用内网poc平台！
- 希望兄弟们能为 [dev](https://github.com/XRSec/Nat_Poc_All/tree/dev) 版本贡献一些代码

## 预览

![image-20220221234444368](https://dogefs.s3.ladydaily.com/ZYGG/storage/image-20220221234444368.png?fmt=webp&q=48)

![image-20220221234545002](https://dogefs.s3.ladydaily.com/ZYGG/storage/image-20220221234545002.png?fmt=webp&q=48)

![image-20220221234602514](https://dogefs.s3.ladydaily.com/ZYGG/storage/image-20220221234602514.png?fmt=webp&q=48)

## 安装

### 命令

```bash
# 启动
docker-compose -f /www/chevereto/docker-compose.yaml up

# 停止
docker-compose -f /www/chevereto/docker-compose.yaml down
```

### 配置

```
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

>  声明：维护不易，请更新文章时向我发送（图片 数据库 文章）用以同步漏洞库
