# feed.thrift 视频流推送微服务接口描述文件
# 更多信息请参考: https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c

# kitex -service feedservice -module github.com/linzijie1998/mini-tiktok idl/feed.thrift

namespace go douyin.feed

include "user.thrift"

# 视频信息
struct Video {
    1: required i64 id,               // 视频ID
    2: required user.User author,     // 视频作者信息
    3: required string play_url,      // 视频播放链接
    4: required string cover_url,     // 视频封面链接
    5: required i64 favorite_count,   // 视频点赞数
    6: required i64 comment_count,    // 视频评论数
    7: required bool is_favorite,     // 是否点赞
    8: required string title,         // 视频标题
}

# 视频流HTTP请求体
struct FeedRequest {
    1: optional i64 latest_time,
    2: optional string token,
}

# 视频流HTTP响应体
struct FeedResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required list<Video> video_list,
    4: optional i64 next_time,
}

service FeedService {
    FeedResponse feed(1: FeedRequest req) (api.get="/douyin/feed/")
}
