# 用户微服务RPC服务端
本服务负责用户的注册（user_register）、登录（user_login）以及用户信息查找（user_info）功能

| 序号 | 数据字段名            | 字段描述 | 使用场景                     |
|----|------------------|------|--------------------------|
| 1  | username         | 登录名  | 用于用户登录的唯一用户名，在注册和登录时会使用到 |
| 2  | password         | 登录密码 | 用于用户登录的密码，在注册和登录时会使用到    |
| 3  | nickname         | 昵称   | 用户在平台上展示的用户名称            |
| 4  | avatar           | 头像   | 用户在平台上展示的个性头像            |
| 5  | background_image | 背景图  | 用户在平台上展示的背景图，在用户主页上显示    |
| 6  | signature        | 签名   | 用户在平台上展示的个性签名，在用户主页上显示   |
| 7  | follow_count     | 关注数  | 用户的关注用户计数，在用户主页上显示       |
| 8  | follower_count   | 粉丝数  | 用户的粉丝用户计数，在用户主页上显示       |
| 9  | total_favorited  | 获赞数  | 用户获得的点赞计数，在用户主页上显示       |
| 10 | favorite_count   | 点赞数  | 用户的点赞视频计数，在用户主页上显示       |
| 11 | work_count       | 作品数  | 用户在平台上发布的视频计数，在用户主页上显示   |

参考配置：`config_template.yaml`，启动前修改为`config.yaml`

```yaml
# MySQL数据库连接配置
mysql:
  host: 127.0.0.1
  port: 3306
  db-name: tiktok
  username: root
  password: 123456
  config: charset=utf8mb4&parseTime=True&loc=Local

# JWT配置
jwt:
  signing-key: bytedance-camp
  expires-time: 168h
  issuer: linzijie
  subject: mini-tiktok

# Etcd服务注册
etcd:
  host: 127.0.0.1
  port: 2379

# RPC服务
rpc_server:
  service-name: user_service
  host: 127.0.0.1
  port: 8880

# 静态资源
static_resource:
  default-avatar: https://imgse.com/i/p912u5V
  default-background-image: https://imgse.com/i/p912VDs

# Redis数据库连接配置
redis:
  host: 127.0.0.1
  port: 6379
  db: 0
  password: ""

# Redis缓存过期时间
cache_expire:
  null-key: 30m # 空值缓存
  user-base-info: 24h
  video-info: 24h

# MongoDB数据库连接配置
mongodb:
  host: 127.0.0.1
  port: 27017
  username: admin
  password: 123456
  database: tiktok
```
