package service

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/model"
	"github.com/ekreke/myTodolist/serializer"
)

// user login service

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (service *UserLoginService) Login(username, password string) serializer.Response {
	db := conf.DB
	u := &model.Users{}
	db.Debug().Where("username = ?", "users").First(&u)
	return serializer.Response{
		Data:
		Status:

	}
}
