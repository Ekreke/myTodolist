package service

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/model"
	"github.com/ekreke/myTodolist/pkg/e"
	"github.com/ekreke/myTodolist/serializer"
	"github.com/jinzhu/gorm"
)

// user register service

type UserRegisterService struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (service UserRegisterService) Register() serializer.Response {
	db := conf.DB
	u := &model.Users{}
	code := e.SUCCESS
	username := service.UserName
	password := service.Password
	msg := ""

	// username existed
	err := db.Debug().Where("username =?", username).First(&u).Error
	// record not found ; username not exist
	if gorm.IsRecordNotFoundError(err) {
		// continue register
		nU := &model.Users{
			Username: username,
			Password: password,
		}
		err := db.Debug().Create(&nU).Error
		if err != nil {
			code = e.ERROR_DB
			msg = e.GetMsg(code)
			return serializer.Response{
				Status: code,
				Msg:    msg,
			}
		}
	} else {
		code = e.ERROR_USER_EXIST
		msg = "username existed!"
		return serializer.Response{
			Status: code,
			Msg:    msg,
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "success",
	}
}
