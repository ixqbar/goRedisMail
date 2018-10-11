
[TOC]

### purpose
* 发送邮件

### version
```
v0.0.1
```

### usage
```
./bin/goRedisMail --config=config.xml
```

### config
```
<?xml version="1.0" encoding="UTF-8" ?>
<config>
    <listen>0.0.0.0:8699</listen>
    <mail>
        <host>smtp.qq.com</host>
        <port>465</port>
        <user>你的邮件账号</user>
        <password>你的邮件密码</password>
    </mail>
</config>
```
* 邮件密码建议开启授权码
* qq邮件请参考 https://service.mail.qq.com/cgi-bin/help?subtype=1&&no=1001256&&id=28
* 163邮件请参考 http://help.mail.163.com/faq.do?m=list&categoryID=197

### example
```
<?php

$redis_handle = new Redis();
$redis_handle->connect('127.0.0.1', 8699, 30);
$result = $redis_handle->hmset('*', ["To" => "174171262@qq.com,xiaochengxu@caoxianjie.com", 'Subject' => "test", 'Content' => 'nice']);
var_dump($result);
```

### deps
* https://github.com/jonnywang/go-kits/redis
* https://github.com/go-gomail/gomail

### faq
* qq群 233415606