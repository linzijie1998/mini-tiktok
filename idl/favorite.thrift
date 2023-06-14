# favorite.thrift 点赞功能微服务接口描述文件
# 更多信息请参考: https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c
# kitex -service favoriteservice -module github.com/linzijie1998/mini-tiktok idl/favorite.thrift
namespace go douyin.favorite

include "feed.thrift"

struct ActionRequest {
    1: required string token,
    2: required i64 video_id,
    3: required i32 action_type,
}

struct ActionResponse {
    1: required i32 status_code,
    2: optional string status_msg,
}

struct ListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct ListResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required list<feed.Video> video_list,
}

service FavoriteService {
    ActionResponse favorite_action(1: ActionRequest req) (api.post="/douyin/favorite/action/")
    ListResponse favorite_list(1: ListRequest req) (api.get="/douyin/favorite/list/")
}