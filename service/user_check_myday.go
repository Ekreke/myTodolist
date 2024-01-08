package service

import (
	"strconv"
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

	// dao get itemids
	// select item_id from my_day where user_id = ? order by id ASC
	db.Debug().
		Select("item_id").
		Where("user_id = ?", u.Id).
		Limit(page.PageSize).
		Where("id > ?", page.NextID).
		Find(&my_days)

	// get item infos
	var items []model.Item
	for i := 0; i < len(my_days); i++ {
		tmp := my_days[i].ItemID

		var item model.Item
		err := db.Debug().Where("id =?", tmp).First(&item).Error
		items = append(items, item)
		if err != nil {
			logging.Fatal(err)
			code = e.ERROR_DB
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	// change page's info
	//TODO: time expire
	lastRecord := items[len(items)-1]
	page.NextID = strconv.Itoa(lastRecord.ID)
	//FIXME: page size
	logging.Info(page)
	proCurToken = utils.Encode(&page)
	resp := &serializer.My_Days{
		Items:       items,
		ProCurToken: proCurToken,
	}
	return serializer.Response{
		Data:   resp,
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
