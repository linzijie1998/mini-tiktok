# 【mini-tiktok】极简版抖音

<div align=left>
<img src="https://img.shields.io/badge/Golang-v1.19-blue"/>
<img src="https://img.shields.io/badge/Hertz-v0.6.2-lightBlue"/>
<img src="https://img.shields.io/badge/Kitex-v0.5.2-red"/>
<img src="https://img.shields.io/badge/LICENSE-MIT-green"/>
</div>

基于微服务RPC框架Kitex和微服务HTTP框架Hertz的极简版抖音后端（第五届字节跳动后端青训营结营项目）

## 一、项目结构
```
├── cmd: 
│    ├── api: API访问接口
│    │    ├── comment.go: 评论计数
│    │    ├── favorite.go: 点赞计数及状态获取
│    │    └── relation.go: 关注和粉丝计数及状态获取
│    ├── comment：评论发布和查询相关微服务
│    ├── favorite：视频点赞及列表相关微服务
│    ├── feed：视频流推送微服务
│    ├── message：用户聊天微服务
│    ├── publish：视频发布及列表查询相关微服务
│    ├── relation：关注、粉丝和朋友相关功能微服务
│    └── user：用户登录、注册、查询相关功能微服务
├── config: 配置信息结构, 对应yaml配置文件
├── idl: 接口描述文件
│    ├── comment.thrift：评论相关
│    ├── favorite.thrift：点赞相关
│    ├── feed.thrift：视频推送
│    ├── message.thrift：聊天相关
│    ├── publish.thrift：视频发布相关
│    ├── relation.thrift：关注相关
│    └── user.thrift：用户相关
├── kitex_gen: 由kitex自动生成工具生成
├── model: 数据实体
└── pkg: 工具
     ├── constant: 保存一些常量
     ├── errno: 错误信息
     ├── ffmpeg: 视频封面截取
     ├── path: 文件路径相关
     └── jwt: JWT生成和解析
```

## 二、项目启动

### 1. 环境准备
本项目在`Ubuntu 22.04 LTS`下进行开发和测试，Windows用户推荐使用WSL2编译运行，并且准备好以下环境：
- 请保证`Go编译器`版本为`1.9.x`，以确保代码能够正常编译运行；
- 需要`MySQL8.0`及以上版本，并且提前创建好一个数据库，字符集为`utf8mb4`；
- 本项目使用的`MongoDB`版本为`v6.0.6`，确保使用支持内存缓存和磁盘缓存的版本以达到最好的运行效果；
- 尽量使用最新版`Redis`作为缓存数据库；
- 搭建`Nginx`服务器映射视频文件上传目录，配置文件可参考以下代码：
    ```
    server {
        listen 8081; # 端口
        server_name localhost; # 服务名
        charset utf-8; # 避免中文乱码
        root /data; # 显示的根索引目录，即文件上传目录
    
        location / {
            autoindex on;             # 开启索引功能
            autoindex_exact_size off; # 关闭计算文件确切大小（单位bytes），只显示大概大小（单位kb、mb、gb）
            autoindex_localtime on;   # 显示本机时间而非 GMT 时间
        }
    }
    ```
- 安装`ffmpeg`，用于截取视频封面；
- 搭建`etcd`服务，用于服务的注册与发现。

### 2. 编译和运行
以`cmd/user`为例，1）进入目录 2）编译 3）运行（请根据`config_template.yaml`填写配置信息，并且将其重命名为`config.yaml`后运行）：
```shell
cd cmd/user
sh build.sh 
sh output/bootstrap.sh
```
