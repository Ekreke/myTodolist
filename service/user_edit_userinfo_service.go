package service

import (
	"time"

	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/model"
	"github.com/ekreke/myTodolist/pkg/e"
	"github.com/ekreke/myTodolist/pkg/logging"
	"github.com/ekreke/myTodolist/serializer"
	"github.com/ekreke/myTodolist/utils"
)

// username when created is uneditable 4ever
type UserEditUserInfoService struct {
	Password string `form:"password" json:"password"`
	Link     string `form:"link" json:"link"`
	Bio      string `form:"bio" json:"bio"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// edit user info
func (service *UserEditUserInfoService) EditUserInfo(password, link, bio, avatar, token string) serializer.Response {
	code := e.SUCCESS
	username := ""
	// can't get token
	if token == "" {
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	} else {
		// get claims
		claims, err := utils.ParseUserToken(token)
		if err != nil {
			logging.Info(err)
			code = e.ERROR_AUTH_TOKEN
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		} else {
			username = claims.Username
		}
	}
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
