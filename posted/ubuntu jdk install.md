---
template:       article
title:          Ubuntu JDK install
date:           2016-12-05 09:00:00 +0800
keywords:       jdk install
description:    ubuntu jdk install
---

# jdk install
```shell
sudo mkdir /usr/lib/java
sudo tar -zxvf jdk-8u112-linux-x64.tar.gz -C /usr/lib/java/

sudo update-alternatives --install /usr/bin/java java /usr/lib/java/jdk1.8.0_60/bin/java 300
sudo update-alternatives --install /usr/bin/javac javac /usr/lib/java/jdk1.8.0_60/bin/javac 300
sudo update-alternatives --install /usr/bin/javah javah /usr/lib/java/jdk1.8.0_60/bin/javah 300
sudo update-alternatives --install /usr/bin/javap javap /usr/lib/java/jdk1.8.0_60/bin/javap 300 
```

# OR
```
export JAVA_HOME=/usr/lib/java/jdk1.8.0_102

export JRE_HOME=$JAVA_HOME/jre

export CLASSPATH=.:$CLASSPATH:$JAVA_HOME/lib:$JRE_HOME/lib

export PATH=$PATH:$JAVA_HOME/bin:$JRE_HOME/bin
```
