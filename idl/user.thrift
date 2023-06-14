# user.thrift 用户相关功能微服务接口描述文件
# 更多信息请参考: https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c
# kitex -service userservice -module github.com/linzijie1998/mini-tiktok idl/user.thrift
namespace go douyin.user

# 用户信息
struct User {
    1: required i64 id,                   // 用户ID
    2: required string name,              // 用户名称
    3: optional i64 follow_count,         // 关注总数
    4: optional i64 follower_count,       // 粉丝总数
    5: required bool is_follow,           // 是否关注
    6: optional string avatar,            // 用户头像
    7: optional string background_image,  // 用户个人页顶部大图
    8: optional string signature,         // 个人简介
    9: optional i64 total_favorited,      // 获赞数
    10: optional i64 work_count,          // 作品数
    11: optional i64 favorite_count,      // 点赞数
}

# 注册功能HTTP请求体
struct RegisterRequest {
    1: required string username,
    2: required string password,
}

# 注册功能HTTP响应体
struct RegisterResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required i64 user_id,
    4: required string token,
}

# 登录功能HTTP请求体
struct LoginRequest {
    1: required string username,
    2: required string password,
}

# 登录功能HTTP响应体
struct LoginResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required i64 user_id,
    4: required string token,
}

# 用户信息查询HTTP请求体
struct InfoRequest {
    1: required i64 user_id,
    2: required string token,
}

# 用户信息查询HTTP响应体
struct InfoResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required User user,
}

service UserService {
    RegisterResponse user_register(1: RegisterRequest req) (api.post="/douyin/user/register/"); // 注册功能
    LoginResponse user_login(1: LoginRequest req) (api.post="/douyin/user/login/");             // 登录功能
    InfoResponse user_info(1: InfoRequest req) (api.get="/douyin/user/");                       // 用户信息查询功能
}
