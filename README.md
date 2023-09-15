# chattyDiaryBot

## 概述
一个可以用来写日记的bot后端
支持多用户同时使用（大概）
接入go-cqhttp

## 部署

### 数据库
项目正常运行需要mysql和redis两个数据库  
如果只是下载下来玩玩的话，可以使用docker迅速部署mysql和redis  
（当然，如果真的需要用bot写日记的话，请使用可靠的方式运行mysql）  

### 配置
配置文件为config/config.yaml，按照文件里的提示进行配置即可  
bot的图像搜索功能需要bing的图像搜索API,获取方法请自行搜索

### 运行
在终端进入`chattyDiaryBot`文件夹内，执行
```bash
./chattyDiaryBot
```
然后bot就会跑起来啦

### 使用
bot的所有功能都在私聊里面实现
- `/r num` 查询近num天的日记
- `/d xxxxxxx` 记录日记
- `/s xxxxxxx` 搜索图片并随机返回一张

## Layout
```
.
├── README.md
├── cmd
│   └── main.go
├── config
│   ├── config-template.yaml
├── go.mod
├── go.sum
└── internal
    ├── config
    │   └── config.go
    ├── controller
    │   ├── params
    │   │   ├── param-http-post.go
    │   │   └── state-handling.go
    │   └── reverse-http-handling.go
    ├── middleware
    │   └── cronMessaging.go
    ├── model
    │   ├── diary.go
    │   ├── init-mysql.go
    │   ├── init-redis.go
    │   └── operation-redis.go
    ├── util
    │   ├── cq-code
    │   │   └── cq-marshal.go
    │   ├── encodeUtil.go
    │   ├── go-cqttp-apis.go
    │   ├── image-search.go
    │   ├── log.go
    │   └── params
    │       ├── request-param.go
    │       └── response-param.go
    └── view
        └── router.go
```
