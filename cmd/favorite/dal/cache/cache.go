package cache

import "fmt"

const (
	userFavoriteKey = "user_favorite_uid%d"
)

func getUserFavoriteKey(id int64) string {
	return fmt.Sprintf(userFavoriteKey, id)
}
