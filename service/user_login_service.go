package service

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/model"
	"github.com/ekreke/myTodolist/pkg/e"
	"github.com/ekreke/myTodolist/pkg/logging"
	"github.com/ekreke/myTodolist/serializer"
	"github.com/jinzhu/gorm"
)

// user login service

type UserLoginService struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (service *UserLoginService) Login(username, password string) serializer.Response {
	db := conf.DB
	u := &model.Users{}
	code := e.SUCCESS
	err := db.Debug().Where("username = ?", username).First(&u).Error
	// check if the record of username exists
	if gorm.IsRecordNotFoundError(err) {
		logging.Info(err)
		code = e.ERROR_NOT_EXIST_USER
		return serializer.Response{
			Data:   "user not exist",
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// check if the password is correct
	if f := checkPassword(*u, password); f {
		return serializer.Response{
			Data:   "login success !",
			Status: code,
			Msg:    e.GetMsg(code),
		}
	} else {
		code = e.ERROR_PASSWORD
		return serializer.Response{
			Data:   "password not correct",
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
}

func checkPassword(u model.Users, password string) bool {
	return u.Password == password
}
