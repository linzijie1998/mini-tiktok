# kitex -service relationservice -module github.com/linzijie1998/mini-tiktok idl/relation.thrift
namespace go douyin.relation

include "user.thrift"

struct FriendUser {
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
    12: optional string message,
    13: required i64 msgType,
}

struct ActionRequest {
    1: required string token,
    2: required i64 to_user_id,
    3: required i32 action_type,
}

struct ActionResponse {
    1: required i32 status_code,
    2: optional string status_msg,
}

struct FollowListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct FollowListResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required list<user.User> user_list,
}

struct FollowerListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct FollowerListResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required list<user.User> user_list,
}

struct FriendListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct FriendListResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required list<FriendUser> user_list,
}

service RelationService {
    ActionResponse relation_action(1: ActionRequest req) (api.post="/douyin/relation/action/")
    FollowListResponse relation_follow_list(1: FollowListRequest req) (api.get="/douyin/relation/follow/list/")
    FollowerListResponse relation_follower_list(1: FollowerListRequest req) (api.get="/douyin/relation/follower/list/")
    FriendListResponse relation_friend_list(1: FriendListRequest req) (api.get="/douyin/relation/friend/list/")
}
