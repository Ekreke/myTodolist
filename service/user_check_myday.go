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

type UserCheckMyDayService struct {
}

func (serivce UserCheckMyDayService) UserCheckMyDay(token, proCurToken string) serializer.Response {
	// init items list
	var my_days []model.My_Days
	code := e.SUCCESS
	var page utils.Page
	db := conf.DB
	var username string

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

	//  check proCurToken
	if proCurToken == "" {
		code = e.ERROR_PRODUCTS_CURSOR_TOKEN_INVALID
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// get the page info
	page = utils.Decode(proCurToken)
	logging.Debug(page)

	// get userid by username
	// id := GetUserIdByUsername(username)
	var u model.Users
	err := db.Debug().Where("username = ?", username).Select("id").First(&u).Error
	if err != nil {
		logging.Info("getuseridbyusername err:", err)
	}

	// dao get items
	// select item_id from my_day where user_id = ? order by id ASC
	db.Debug().
		Select("item_id").
		Where("user_id = ?", u.Id).
		Limit(page.PageSize).
		Where("id > ?", page.NextID).
		Find(&my_days)

	return serializer.Response{
		Data:   my_days,
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
