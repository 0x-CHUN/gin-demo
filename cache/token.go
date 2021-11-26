package cache

import "time"

func SaveToken(token string, userID string, exp time.Duration) error {
	err := RedisClient.Set(UserTokenKey(token), userID, exp).Err()
	return err
}

func GetUserIDByToken(token string) (string, error) {
	return RedisClient.Get(UserTokenKey(token)).Result()
}

func DeleteUserToken(token string) error {
	return RedisClient.Del(UserTokenKey(token)).Err()
}
