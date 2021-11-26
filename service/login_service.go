package service

import (
	"fmt"
	"gin-demo/model"
	"gin-demo/serializer"
	"gin-demo/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Token    bool   `form:"token" json:"token"`
}

func (s *LoginService) setSession(c *gin.Context, user model.User) {
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	fmt.Println(user)
	err := session.Save()
	if err != nil {
		utils.Log().Warning("Session save error")
	}
}

func (s *LoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", s.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	if !user.CheckPassword(s.Password) {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	var token string
	var tokenExpire int64
	var err error
	if s.Token {
		token, tokenExpire, err = user.MakeToken()
		if err != nil {
			return serializer.DBErr("redis err", err)
		}
	} else {
		s.setSession(c, user)
	}
	data := serializer.BuildUser(user)
	data.Token = token
	data.TokenExpire = tokenExpire
	return serializer.Response{
		Data: data,
	}
}
