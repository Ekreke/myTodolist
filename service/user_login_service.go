package service

import (
	"github.com/ekreke/myTodolist/conf"
	"github.com/ekreke/myTodolist/model"
	"github.com/ekreke/myTodolist/pkg/e"
	"github.com/ekreke/myTodolist/pkg/logging"
	"github.com/ekreke/myTodolist/pkg/util"
	"github.com/ekreke/myTodolist/serializer"
	"github.com/jinzhu/gorm"
)

// user login service

// form -> front end json -> backend
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

	// check password
	if checkPassword(*u, password) == false {
		return serializer.Response{
			Data:   "password error",
			Status: e.ERROR_PASSWORD,
			Msg:    e.GetMsg(code),
		}
	}

	// try to get token
	token, err := util.GenerateUserToken(u.Username, u.Password)
	if err != nil {
		logging.Info(err)
		code = e.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	usr := *u
	//get token success
	return serializer.Response{
		Data:   serializer.TokenData{User: serializer.BuildUser(usr), Token: token},
		Status: code,
		Msg:    e.GetMsg(code),
	}

}

func checkPassword(u model.Users, password string) bool {
	return u.Password == password
}
