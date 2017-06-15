---
template:       article
title:          Ubuntu git repo and gitweb
date:           2017-06-15 15:00:00 +0800
keywords:       git repo
description:    git repo
---

# 1. git server
## git install
```shell
sudo apt-get update
sudo apt-get install git -y
```
## 创建git用户 & 设置密码
```shell
sudo useradd git -d /home/git -m -s /bin/bash
sudo passwd git
```
## 切换用户
```shell
su git
cd ~
```
## 创建仓库 /home/git
```shell
git init --bare sample.git
```
## 客户端使用
```shell
git clone git@127.0.0.1:/home/git/sample.git
```
## 为安全考虑Git账号只允许使用git-shell。在passwd文件中找到git用户，把`/bin/bash`直接修改成`/usr/bin/git-shell` 登录root账号，并修改git的用户权限。
```shell
sudo vim /etc/passwd
```

# 2. nginx + gitweb
## 安装 spawn-fcgi 和 fcgiwrap
```shell
sudo apt install spawn-fcgi
sudo apt install fcgiwrap
```

## 启动 fcgiwrap
```shell
sudo spawn-fcgi -f /usr/sbin/fcgiwrap -p 9000
```

## 创建配置文件 `/etc/gitweb.conf` 并加入
```
# path to git projects (<project>.git)
$projectroot = "/home/git/";
```

## Nginx 配置
```
server {
    listen	8080;
    server_name	localhost;

    root /usr/share/gitweb/;

    # fix: 413 Request Entity Too Large
    client_max_body_size	100m;

    location / {
        access_log	off;
        expires		24h;

	index		index.cgi;
	include		fastcgi.conf;
	gzip		off;

	if ($uri ~ "/index.cgi") {
	    fastcgi_pass	127.0.0.1:9000;
	}
    }
}
```

### Note: 可使用cgit https://git.zx2c4.com/cgit/about/

