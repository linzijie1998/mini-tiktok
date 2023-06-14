# 用户微服务RPC服务端
本服务负责用户的注册（user_register）、登录（user_login）以及用户信息查找（user_info）功能。

## 一、数据设计

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

## 二、服务逻辑

### 2.1 用户注册