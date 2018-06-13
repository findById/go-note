---
template:       article
title:          docker work host
date:           2018-06-13 18:00:00 +0800
keywords:       docker work host
description:    docker host
---

# jenkins
```shell
mkdir /home/work/jenkins

sudo docker pull jenkins
# remove
sudo docker stop jenkins
sudo docker container rm jenkins
# add and run
sudo docker run -itd -p 8082:8080 --name jenkins --privileged=true  -v /home/work/jenkins:/var/jenkins_home -v /home/work/Android/Sdk:/var/jenkins_home/android-sdk-linux -v /home/work/.android:/var/jenkins_home/.android -v /home/work/.gradle:/var/jenkins_home/.gradle jenkins

sudo docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```
## android sdk
```shell
ln -s /home/work/Android/Sdk /home/work/jenkins/android-sdk-linux
```
## gradle
```shell
ln -s /home/work/jenkins/.gradle /home/work/.gradle
```
## git ssh access
```shell
docker exec -it <containerId or name> /bin/sh

ssh-keygen -t rsa -C "your.email@example.com" -b 4096

cat ~/.ssh/id_rsa.pub | clip
```

# sonatype nexus
```shell
mkdir /home/work/nexus-data
chmod -R 777 /home/work/nexus-data

sudo docker pull sonatype/nexus3
# docker run -itd --name nexus --restart=always -p 8081:8081 -v /home/work/nexus-data:/nexus-data
sudo docker run -d -p 8081:8081 --name nexus -v /home/work/nexus-data:/nexus-data --restart=always sonatype/nexus3
```

# redis
```shell
mkdir /home/work/redis
mkdir /home/work/redis/data

sudo docker run -d -p 6379:6379 --name redis-server -v /home/work/redis/redis.conf:/etc/redis/redis.conf -v /home/work/redis/data:/data docker.io/redis redis-server /etc/redis/redis.conf --appendonly yes
```
## redis-tools
```shell
sudo apt install redis-tools

redic-cli
```
