package service

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/model"
)

// user login service

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (service *UserLoginService) Login(username, password string) {
	db := conf.DB
	u := &model.Users{}
	db.Find(u, "username =?", username)
	db.Debug()
}
