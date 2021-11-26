package cache

import "fmt"

func UserTokenKey(token string) string {
	return fmt.Sprintf("user:token:%s", token)
}
