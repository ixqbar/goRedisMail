<?php

try {
    $redis_handle = new Redis();
    $redis_handle->connect('127.0.0.1', 8699, 30);
    $result = $redis_handle->hmset('*', ["To" => "174171262@qq.com,xiaochengxu@caoxianjie.com", 'Subject' => "test", 'Content' => '我知道了在弹出的窗口中输入待加入白名单的邮箱地址，点击“确定”完成添加']);
    var_dump($result);
} catch(Throwable $e) {
   echo $e->getMessage() . PHP_EOL;
}