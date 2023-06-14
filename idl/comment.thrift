# kitex -service commentservice -module github.com/linzijie1998/mini-tiktok idl/comment.thrift
namespace go douyin.comment

include "user.thrift"

struct Comment {
    1: required i64 id,
    2: required user.User user,
    3: required string content,
    4: required string create_date,
}

struct ActionRequest {
    1: required string token,
    2: required i64 video_id,
    3: required i32 action_type,
    4: optional string comment_text,
    5: optional i64 comment_id,
}

struct ActionResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: optional Comment comment, // 评论成功返回评论内容, 无需重新拉取列表
}

struct ListRequest {
    1: required string token,
    2: required i64 video_id,
}

struct ListResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required list<Comment> comment_list,
}

service CommentService {
    ActionResponse comment_action(1: ActionRequest req) (api.post="/douyin/comment/action/")
    ListResponse comment_list(1: ListRequest req) (api.get="/douyin/comment/list/")
}
