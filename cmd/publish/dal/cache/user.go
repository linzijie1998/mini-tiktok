package cache

import "context"

func GetUserInfoNullKey(ctx context.Context, uid int64) error {
	key := getUserInfoNullKey(uid)
	return getNullKey(ctx, key)
}

func incrByUserField(ctx context.Context, uid int64, field string) error {
	key := getUserCounterKey(uid)
	return change(ctx, key, field, 1)
}

func IncrWorkCount(ctx context.Context, uid int64) error {
	return incrByUserField(ctx, uid, "work_count")
}
