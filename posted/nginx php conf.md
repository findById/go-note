---
template:       article
title:          Nginx php conf
date:           2016-12-16 09:00:00 +0800
keywords:       Nginx php conf
description:    Nginx php conf
---

## Nginx + php
```
server {
	# linux start php
	# spawn-fcgi -a 127.0.0.1 -p 9000 -c 10 -u www-data -f /usr/bin/php-cgi
	# windows start php
	# d:\dev\php\php-nts\php-cgi.exe -b 127.0.0.1:9000 -c d:\dev\php\php-nts\php.ini

	listen 9090;
	server_name	localhost;
	#charset	UTF-8;
	#access_log	logs/host.access.log	main;

	root php;
	index index.html index.php;

	location / {
		try_files $uri $uri/ /index.php;
	}

	# pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
	location ~ \.php$ {
		try_files		$uri = 404;

		include			fastcgi.conf;
		fastcgi_pass	127.0.0.1:9000;
	}

	#error_page	404	/404.html;

	# redirect server error pages to the static page /50x.html
	error_page	500	502	503	504	/50x.html;
	location = /50x.html {
		root html;
	}
}
```
