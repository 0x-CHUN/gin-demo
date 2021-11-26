package service

import (
	"gin-demo/model"
	"gin-demo/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

func (s *UserRegisterService) valid() *serializer.Response {
	if s.PasswordConfirm != s.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}
	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", s.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}
	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", s.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}
	return nil
}

func (s *UserRegisterService) Register() serializer.Response {
	user := model.User{
		Nickname: s.Nickname,
		UserName: s.UserName,
		Status:   model.Active,
	}
	if err := s.valid(); err != nil {
		return *err
	}
	if err := user.SetPassword(s.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}
	return serializer.BuildUserResponse(user)
}
