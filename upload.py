#!/usr/bin/env python3
# _*_ coding: utf-8 _*_
import datetime, os, re, requests, sys, json, tempfile, shutil, imghdr, base64

requests.packages.urllib3.disable_warnings()

# burpsuite
proxies = {
    "http": "http://127.0.0.1:8080",
    "https": "http://127.0.0.1:8080"
}
print(requests.post(url="http://localhost:8081/upload", data="",
                    files={'uploadFile': ("任务.md", open(os.getcwd() + "/upd/任务.md", 'rb'))},
                    timeout=5, verify=False, proxies=proxies).text)
print(requests.post(url="http://localhost:8081/upload", data="",
                    files={'uploadFile': ("poc.png", open(os.getcwd() + "/upd/poc.png", 'rb'))},
                    timeout=5, verify=False, proxies=proxies).text)