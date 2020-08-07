[TOC]

# moc

a http moc server

# 启动服务

1. 直接启动

```
moc
```

这种启动方式，按照

当前目录 -> ~/.moc/ -> /etc/moc/

的顺序查找配置文件conf.json。若找到了，则使用该配置文件启动服务。若未找到，则服务启动失败。即配置文件的优先级为

conf.json > ~/.moc/conf.json > /etc/moc/conf.json

2. 命令行指定配置文件

```
moc 配置文件完整路径
```

这种方式使用指定的配置文件启动服务

# 配置文件

配置文件实例

```json
{
    "port": 8017,
    "response_root_path": "/home/my_user/moc",
    "moc": [
        {
            "path": "/user/login",
            "method": "post",
            "sleep": 5000,
            "rules": [
                {
                    "body": ".*",
                    "request": ".*",
                    "response": "ccc",
                    "response_file": "user_login.json",
                    "response_shell": "user_login.py"
                }
            ]
        }
    ]
}
```

| 字段 | 必填 | 说明 |
| ---  | --- |  --- |
| port | 是 | 服务端口号 |
| response_root_path | 否 | 指定response_root_path后，response_file、response_shell 可使用相对路径 |
| moc | 是 | moc规则集 |

字段说明

| 字段 | 必填 | 说明 |
| --- | --- | --- |
| path   | 是   | url的path部分 |
| method | 是   | 请求的方法 |
| slee   | 否   | 休眠一段时间再向客户端返回，单位：毫秒 |
| rules  | 是   | 规则集 |

rules字段说明

|字段|必填|说明|
|---|---|---|
|body    |否|post请求body正则表达式，若不填或为空视为匹配成功|
|request |否|查询字符串正则表达式，url"?"后面的内容，若不填或为空视为匹配成功|
|response|否|返回的响应数据|
|response_file|否|响应返回指定文件中的内容|
|response_shell|否|获取指定脚本的标准输出并返回|

请求的查询字符串、body与给定的正则表达式完全匹配，则返回response、response_file、response_shell指定的内容。

响应字段的优先级为：response > response_file > response_shell

即当三个字段同时存在时

* 若response不为空，则返回response指定的内容

* 若response为空，response_file指定的文件存在，则返回该文件的内容

* 若response为空，response_file为空或指定的文件不存在，response_shell指定的文件存在，则以执行response_shell的标准输出作为响应内容。目前response_shell只支持python脚本

* 否则返回默认内容

对于一个请求

```json
curl -X POST -d 'this is a body' http://127.0.0.1:9008/user/login?name=n&password=123
```

依次对查询字符串"name=n&password=123"和body"this is a body"进行正则匹配，若全部匹配成功，则返回指定的内容。

特别的，body和request字段均可为空，则默认匹配成功

## 实例1 简单的登录接口

配置文件

```json
{
    "port": 8017,
    "moc": [
        {
            "path": "/user/login",
            "method": "get",
            "rules": [
                {
                    "response": "{\"status\":0,\"msg\":\"ok\"}"
                }
            ]
        }
    ]
}
```



## 实例2 复杂一点的登录接口

moc一个登录接口，用户"zhangsan"无论输入什么都视为登录成功，用户"lisi"无论输入什么都视为登录失败，其他用户输入密码"123456"视为登录成功。

请求url

```json
http://127.0.0.1:9008/user/login?name=zhangsan&password=123
```

配置文件

```json
{
    "port": 8017,
    "moc": [
        {
            "path": "/user/login",
            "method": "get",
            "rules": [
                {
                    "request": "name=zhangsan&password=.*",
                    "response": "{\"status\":0,\"msg\":\"ok\"}"
                },
                {
                    "request": "name=lisi&password=.*",
                    "response": "{\"status\":1,\"msg\":\"login failed\"}"
                },
                {
                    "request": "name=.*&password=123456$",
                    "response": "{\"status\":0,\"msg\":\"ok\"}"
                },
                {
                    "response": "{\"status\":1,\"msg\":\"login failed\"}"
                }
            ]
        }
    ]
}
```

## 实例3 以文件作为响应内容

配置文件

```json
{
    "port": 8017,
    "response_root_path": "/home/my_user/moc",
    "moc": [
        {
            "path": "/user/login",
            "method": "post",
            "rules": [
                {
                    "response_file": "user_login.json"
                }
            ]
        }
    ]
}
```

## 实例4 以脚本标准输出作为响应内容

配置文件

```json
{
    "port": 8017,
    "response_root_path": "/home/my_user/moc",
    "moc": [
        {
            "path": "/user/login",
            "method": "post",
            "rules": [
                {
                    "response_file": "user_login.py"
                }
            ]
        }
    ]
}
```

user_login.py脚本内容为

```python
import sys


print("*****************")
print(len(sys.argv))
print(sys.argv[0])
print(sys.argv[1])
print(sys.argv[2])
print("*****************")
```

请求

```
curl -X POST -d 'this is a body' http://127.0.0.1:9008/user/login?name=n&password=123
```

响应

```
*****************
3
/home/my_user/moc/response/user_login.py
name=n&password=123
this is a body
*****************
```
