package api

import (
	"gin-demo/cache"
	"gin-demo/serializer"
	"gin-demo/service"
	"gin-demo/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var s service.UserRegisterService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {
	var s service.LoginService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserTokenRefresh(c *gin.Context) {
	user := CurrentUser(c)
	var s service.TokenRefreshService
	res := s.Refresh(c, user)
	c.JSON(200, res)
}

func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

func UserLogout(c *gin.Context) {
	// 移动端登出
	token := c.GetHeader("X-Token")
	if token != "" {
		_ = cache.DeleteUserToken(token)
	} else {
		// web端登出
		s := sessions.Default(c)
		s.Clear()
		err := s.Save()
		if err != nil {
			utils.Log().Warning("Save session error ", err)
		}
	}
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
