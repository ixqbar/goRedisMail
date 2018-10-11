
[TOC]

### purpose
* 发送邮件

### version
```
v0.0.2
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

try {
    $redis_handle = new Redis();
    $redis_handle->connect('127.0.0.1', 8699, 30);
    $result = $redis_handle->hmset('*', [
        "To" => "174171262@qq.com,xiaochengxu@caoxianjie.com",
        'Subject' => "运行测试通过",
        'Content' => '恭喜你完成设置，更多功能请及时关注<a href="https://github.com/jonnywang/goRedisMail">goRedisMail</a>']
    );
    var_dump($result);
} catch(Throwable $e) {
   echo $e->getMessage() . PHP_EOL;
}
```

### deps
* https://github.com/jonnywang/go-kits/redis
* https://github.com/go-gomail/gomail

### faq
* qq群 233415606