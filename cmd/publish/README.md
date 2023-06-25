# 视频发布微服务RPC服务端
本服务负责视频的上传和发布功能

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
  signing-key: bytedance-project
  expires-time: 168h
  issuer: linzijie
  subject: mini-tiktok

# Etcd服务注册
etcd:
  host: 127.0.0.1
  port: 2379

# RPC服务
rpc_server:
  service-name: publish_service
  host: 127.0.0.1
  port: 8881

# 静态资源
static_resource:
  default-avatar: https://s1.ax1x.com/2023/04/29/p912u5V.jpg
  default-background-image: https://s1.ax1x.com/2023/04/29/p912VDs.jpg

# Redis数据库连接配置
redis:
  host: 127.0.0.1
  port: 6379
  db: 0
  password: ""

# 文件上传和访问
file_access:
  upload_path: /opt/tiktok # 文件上传路径
  nginx_url: http://192.168.31.160:8081/ # nginx服务地址

# Redis缓存过期时间
cache_expire:
  null-key: 30m # 空值缓存
  user-base-info: 24h
  video-info: 24h

```
