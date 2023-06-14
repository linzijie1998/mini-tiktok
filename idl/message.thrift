# kitex -service messageservice -module github.com/linzijie1998/mini-tiktok idl/message.thrift
namespace go douyin.message

struct Message {
    1: required i64 id,
    2: required i64 to_user_id,
    3: required i64 from_user_id,
    4: required string content,
    5: optional string create_date,
}

struct ChatRequest {
    1: required string token,
    2: required i64 to_user_id,
    3: required i64 pre_msg_time,
}

struct ChatResponse {
    1: required i32 status_code,
    2: optional string status_msg,
    3: required list<Message> message_list,
}

struct ActionRequest {
    1: required string token,
    2: required i64 to_user_id,
    3: required i32 action_type,
    4: required string content,
}

struct ActionResponse {
    1: required i32 status_code,
    2: optional string status_msg,
}

service MessageService {
    ChatResponse message_chat(1: ChatRequest req) (api.get="/douyin/message/chat/")
    ActionResponse message_action(1: ActionRequest req) (api.post="/douyin/message/action/")
}
