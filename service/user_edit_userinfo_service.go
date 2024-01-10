package service

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/model"
	"github.com/ekreke/myTodolist/pkg/e"
	"github.com/ekreke/myTodolist/serializer"
)

// username when created is uneditable 4ever
type UserEditUserInfoService struct {
	Password string `form:"password" json:"password"`
	Link     string `form:"link" json:"link"`
	Bio      string `form:"bio" json:"bio"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// edit user info
func (service *UserEditUserInfoService) EditUserInfo(password, link, bio, avatar, username string) serializer.Response {
	code := e.SUCCESS
	db := conf.DB
	u := &model.Users{}
	err := db.Debug().First(&u, "username =?", username).Error
	if err != nil {
		code = e.ERROR_NOT_EXIST_USER
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	u.Password = password
	u.Avatar = avatar
	u.Link = link
	u.Bio = bio
	db.Debug().Save(u)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(*u),
	}
}
