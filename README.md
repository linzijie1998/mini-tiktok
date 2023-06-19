# mini-tiktok
基于微服务RPC框架Kitex和微服务HTTP框架Hertz的极简版抖音后端（第五届字节跳动后端青训营结营项目）

## Redis缓存

### 用户信息

- 用户基本信息

| 字段             | 描述      |
|----------------|---------|
| id             | 用户唯一uid |
| nickname       | 用户昵称    |
| avatar         | 用户头像    |
| background_img | 用户主页背景  |
| signature      | 用户个性签名  |

- 用户计数信息

| 字段              | 描述      |
|-----------------|---------|
| id              | 用户唯一uid |
| follow_count    | 关注数     |
| follower_count  | 粉丝数     |
| favorite_count  | 点赞数     |
| work_count      | 发布作品数   |
| total_favorited | 作品获赞数   |
