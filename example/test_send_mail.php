<?php

try {
    $redis_handle = new Redis();
    $redis_handle->connect('127.0.0.1', 8699, 30);
    $result = $redis_handle->hmset('*', [
        "To" => "174171262@qq.com",
        "Cc" => "2795934612@qq.com",
        'Subject' => "运行测试通过",
        'Content' => '恭喜你完成设置，更多功能请及时关注<a href="https://github.com/jonnywang/goRedisMail">goRedisMail</a>',
        'Attach' => "jx.jpg,订阅号-开发权限.png"
    ]);
    var_dump($result);
} catch(Throwable $e) {
   echo $e->getMessage() . PHP_EOL;
}