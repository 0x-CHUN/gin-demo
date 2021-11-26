package server

import (
	"gin-demo/api"
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", api.Ping)
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		auth := v1.Group("")
		auth.Use(middleware.Auth())
		{
			auth.GET("/user/me", api.UserMe)
			auth.DELETE("/user/logout", api.UserLogout)
			auth.POST("/user/refresh", api.UserTokenRefresh)
		}
	}
	return r
}
