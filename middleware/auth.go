package middleware

import (
	"gin-demo/cache"
	"gin-demo/model"
	"gin-demo/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid interface{}
		token := c.GetHeader("X-Token")
		if token != "" {
			uid, _ = cache.GetUserIDByToken(token)
		} else {
			session := sessions.Default(c)
			uid = session.Get("user_id")
		}
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
