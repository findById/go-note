---
template:       article
title:          Ubuntu disk mount and umount
date:           2016-12-16 09:00:00 +0800
keywords:       ubuntu disk mount and umount
description:    ubuntu disk mount and umount
---

## 查看磁盘列表
```shell
sudo fdisk -l
```
## 创建挂载目录
```shell
sudo mkdir /media/storage0
```
## 挂载磁盘(/dev/sda7)到目录(/media/storage0)
```shell
sudo mount -o codepage=936,iocharset=cp936 /dev/sda7 /media/storage0
```
## 反挂载目录（/media/storage0）
```shell
sudo umount /media/storage0
```

## 挂载优盘命令如下 (并且能够正确显示中文)：
```shell
sudo mkdir /media/U
sudo mount /dev/sda1 /media/U/ -t vfat -o
```
## 开机挂载（自动将分区挂载到/mnt/d上）
### 创建/mnt/d目录
```shell
sudo mkdir /mnt/d
```
### 编辑`/etc/fstab`
```shell
sudo gedit /etc/fstab
```
### 加入以下一行
```
/dev/hda5 /mnt/d vfat defaults,codepage=936,iocharset=cp936 0 0
```
### 这样每次开机后就可以自动挂在分区
