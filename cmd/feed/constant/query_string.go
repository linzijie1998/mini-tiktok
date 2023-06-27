package constant

const (
	UserBaseInfoQueryString     = "id, nickname, avatar, background_image, signature"
	UserCounterInfoQueryString  = "id, follow_count, follower_count, total_favorited, favorite_count, work_count"
	UserLoginInfoQueryString    = "id, password"
	UserRegisterInfoQueryString = "id"
	VideoBaseInfoQueryString    = "id, author_id, title, video_path, cover_path"
	VideoCounterInfoQueryString = "id, favorite_count, comment_count"
)
