---
template:       article
title:          Ubuntu make ffmpeg
date:           2016-12-16 09:00:00 +0800
keywords:       Ubuntu make ffmpeg
description:    Ubuntu make ffmpeg
---

## 安装编译库
```shell
sudo apt-get install gcc
```
## 安装依赖
```shell
sudo apt-get install libfdk-aac-dev libfaac-dev libx264-dev libx265-dev libsdl1.2-dev
```
## 配置编译
```shell
./configure --prefix=/usr/local/ffmpeg --enable-version3 --enable-gpl --enable-nonfree --disable-asm --enable-shared --disable-yasm --enable-libfaac --enable-libfdk-aac --enable-libx264 --enable-libx265
```
### HLS切片
```shell
ffmpeg -i test.mp4 -codec copy -bsf:v h264_mp4toannexb -codec:a mp3 -hls_time 10 -hls_list_size 0 -hls_segment_filename 'file-ts-%05d.ts' playlist.m3u8
```
### HLS切片
```shell
ffmpeg -i a.rmvb -c:v libx264 -codec:a mp3 -hls_time 10 -hls_list_size 0 ./aaa/aaa.m3u8
```
