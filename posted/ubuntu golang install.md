---
template:       article
title:          Ubuntu golang install
date:           2016-12-01 15:00:00 +0800
keywords:       Ubuntu golang install
description:    Ubuntu golang install
---

# golang install
```shell
sudo tar -C /usr/lib/ -xzf gox.x.x.linux-amd64.tar.gz
```

## environment

```
export GOROOT=/usr/lib/go
export GOPATH=$HOME/dev/go/
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin

```
