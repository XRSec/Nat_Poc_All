#!/usr/bin/env python3
# _*_ coding: utf-8 _*_

import os, re
from shutil import copyfile


def ds_store(path):
    ds_store_file = path + ".DS_Store"
    while os.path.exists(ds_store_file):
        os.remove(ds_store_file)
        print("delete .DS_Store")
    return


def get_filelist(dir, Filelist):
    if os.path.isfile(dir):
        Filelist.append(dir)

    elif os.path.isdir(dir):
        for s in os.listdir(dir):
            if s == "resource":
                continue
            newDir = os.path.join(dir, s)
            get_filelist(newDir, Filelist)
    return Filelist


if __name__ == '__main__':
    filepath = "/Users/xr/Downloads/PeiQi_Wiki/test/"
    ds_store(filepath)
    g = os.walk(filepath)
    for path, dir_list, file_list in g:
        for file_name in file_list:
            if (".md" in os.path.join(path, file_name)) and ("README" not in os.path.join(path, file_name)) and ("peiqi.py" not in os.path.join(path, file_name)):
                l1 = (os.path.join(path).replace("/Users/xr/Downloads/PeiQi_Wiki/test/", "")).split("/")
                categories1, categories2= l1[0],l1[1]
                date = "---\ntitle: " + file_name.replace(".md","") + "\ndate: 2021-09-23 23:55:51\ntags: PeiQi文库\ncategories:\n - "+ categories1 + "\n - "+ categories2 +"\n---\n\n"
                # print(date)
                open("/Users/xr/Downloads/PeiQi_Wiki/test1/" + file_name, "w").close()
                with open("/Users/xr/Downloads/PeiQi_Wiki/test1/" + file_name, "w") as hexo:
                    hexo.write(date + open(os.path.join(path, file_name),"r").read())
            if (".zip" in os.path.join(path, file_name)) or ("py" in os.path.join(path, file_name)):
                copyfile(os.path.join(path, file_name), "/Users/xr/Downloads/PeiQi_Wiki/hexo/source/Poc/" + file_name)


