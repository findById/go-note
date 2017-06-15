---
template:       article
title:          Gitlab install
date:           2016-12-16 09:00:00 +0800
keywords:       Gitlab install
description:    Gitlab install
---

查看安装文档 https://about.gitlab.com/downloads/
(If you are located in China, try using https://mirror.tuna.tsinghua.edu.cn/help/gitlab-ce/)

## Example

## Install and configure the necessary dependencies
```
sudo apt-get install curl openssh-server ca-certificates postfix
```
## Add the GitLab package server and install the package
```
curl -sS https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | sudo bash
```
## 首先信任 GitLab 的 GPG 公钥:
```
curl https://packages.gitlab.com/gpg.key 2> /dev/null | sudo apt-key add - &>/dev/null
vi /etc/apt/sources.list.d/gitlab-ce.list 加入
deb https://mirrors.tuna.tsinghua.edu.cn/gitlab-ce/ubuntu xenial main
```
## 安装 gitlab-ce:
```
sudo apt-get update
sudo apt-get install gitlab-ce
```
## Configure and start GitLab
```
sudo gitlab-ctl reconfigure
```
## 打开 sshd 和 postfix 服务
```
service sshd start
service postfix start
```
## (可忽略)需要 GitLab 社区版的 Web 界面可以通过网络进行访问，我们需要允许 80 端口通过防火墙，这个端口是 GitLab 社区版的默认端口。为此需要运行下面的命令
```
sudo iptables -A INPUT -p tcp -m tcp --dport 80 -j ACCEPT
```
## 检查GitLab是否安装好并且已经正确运行
```
sudo gitlab-ctl status
```
## gitlab 默认web端口是80端口 So.
## 修改默认web端口
```
sudo vi /opt/gitlab/embedded/conf/nginx.conf
sudo vi /var/opt/gitlab/gitlab-rails/etc/gitlab.yml
sudo vi /var/opt/gitlab/nginx/conf/gitlab-http.conf
```
## 重启gitlab
```
./opt/gitlab/bin/gitlab-ctl stop
./opt/gitlab/bin/gitlab-ctl start
```
## 修改密码
Ps: http://docs.gitlab.com/ce/security/reset_root_password.html#how-to-reset-your-root-password

### 使用root权限 执行 gitlab-rails console production
```shell
root@work:~# gitlab-rails console production
Loading production environment (Rails 4.2.7.1)
irb(main):001:0> user = User.where(id: 1).first
=> #<User id: 1, email: "admin@example.com", ...
irb(main):002:0> user.password=12345678
=> 12345678
irb(main):003:0> user.password_confirmation=12345678
=> 12345678
irb(main):004:0> user.save!
=> true
irb(main):005:0> quit
```
