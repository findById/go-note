---
template:       article
title:          ubuntu server handle lid switch ignore to do nothing
date:           2016-12-16 09:00:00 +0800
keywords:       ubuntu server handle lid switch ignore to do nothing
description:    ubuntu server handle lid switch ignore to do nothing
---

To get started:

1. Press Ctrl+Alt+T on keyboard to open the terminal. When it opens, run the command below to open the configuration file via gedit editor:
```
sudo gedit /etc/systemd/logind.conf
```
Replace gedit with vi or other text editor if you’re on Server edition.

2. Find out the line #HandleLidSwitch=suspend, remove the # and change it to:
```
HandleLidSwitch=poweroff to shutdown computer when lid is closed
HandleLidSwitch=hibernate to hibernate computer when lid is closed
HandleLidSwitch=ignore to do nothing
```
3. Save the file and restart the service or just restart your laptop to apply changes.
```
sudo restart systemd-logind
```
Update: for Ubuntu 16.04, the command to restart the systemd service should be:
```
systemctl restart systemd-logind.service
```
That’s it. Enjoy!