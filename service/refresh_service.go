package service

import (
	"gin-demo/model"
	"gin-demo/serializer"
	"github.com/gin-gonic/gin"
)

type TokenRefreshService struct {
}

func (s *TokenRefreshService) Refresh(_ *gin.Context, user *model.User) serializer.Response {
	token, tokenExpire, err := user.MakeToken()
	if err != nil {
		return serializer.DBErr("redis err", err)
	}
	data := serializer.BuildUser(*user)
	data.Token = token
	data.TokenExpire = tokenExpire
	return serializer.Response{
		Data: data,
	}
}
