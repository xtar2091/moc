[TOC]

# moc

a http moc server

# 启动服务

moc --port=9008 --conf=conf.json

# 配置文件

配置文件实例

```json
{
    "moc": [
        {
            "path": "/user/login",
            "method": "post",
            "sleep": 5000,
            "rules": [
                {
                    "body": ".*",
                    "request": ".*",
                    "response": "ccc"
                }
            ]
        }
    ]
}
```

字段说明

| 字段   | 必填 | 说明                                   |
| ------ | ---- | -------------------------------------- |
| path   | 是   | url的path部分                          |
| method | 是   | 请求的方法                             |
| slee   | 否   | 休眠一段时间再向客户端返回，单位：毫秒 |
| rules  | 是   | 规则集                                 |

rules字段说明

| 字段     | 必填 | 说明                                                         |
| -------- | ---- | ------------------------------------------------------------ |
| body     | 否   | post请求body正则表达式，若不填或为空视为匹配成功             |
| request  | 否   | 查询字符串正则表达式，url"?"后面的内容，若不填或为空视为匹配成功 |
| response | 是   | 返回的响应数据                                               |

请求的查询字符串、body与给定的正则表达式完全匹配，则返回response指定的内容。

对于一个请求

```json
curl -X POST -d '{}' http://127.0.0.1:9008/user/login?name=n&password=123
```

依次对查询字符串"name=n&password=123"和body"{}"进行正则匹配，若全部匹配成功，则返回resposne指定的内容。

特别的，body和request字段均可为空，则默认匹配成功

## 实例1 简单的登录接口

配置文件

```json
{
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

