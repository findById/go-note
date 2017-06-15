---
template:       article
title:          Windows mysql install
date:           2016-12-16 09:00:00 +0800
keywords:       Mysql install
description:    Mysql install
---

# Windows 下MySql安装
### 配置环境变量
### 修改数据库配置文件 并 指定数据库绝对路径 `my-default.ini` => `my.ini`
```
basedir = C:\Program Files\mysql-5.7.15
datadir = C:\Program Files\mysql-5.7.15\data
```
## 初始化
```
mysqld --initialize-insecure
```
## 安装服务
```
mysqld -install
```
## 卸载服务
```
mysqld -remove
```
## 启动服务 or 停止服务
```
net start mysql
net stop mysql
```
## 设置初始密码 (无密码状态)
```
mysql -u root
SET PASSWORD FOR 'root'@'localhost'=PASSWORD('root');
```
