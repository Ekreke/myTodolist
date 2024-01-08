package api

import (
	"github.com/ekreke/myTodolist/pkg/logging"
	"github.com/ekreke/myTodolist/serializer"
	"github.com/ekreke/myTodolist/service"
	"github.com/ekreke/myTodolist/utils"
	"github.com/gin-gonic/gin"
)

// @Summary
// @Description
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body model.Users true
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	err := c.ShouldBind(&service)
	if err != nil {
		logging.Info(err)
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	resp := service.Login(username, password)
	c.JSON(200, resp)
}

// user regiseter
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	err := c.ShouldBind(&service)
	if err != nil {
		logging.Info(err)
	}
	resp := service.Register()
	c.JSON(200, resp)
}

// SetUserInfo  test : input token in post body
func SetUserInfo(c *gin.Context) {
	var service service.UserEditUserInfoService
	err := c.ShouldBind(&service)
	if err != nil {
		logging.Info(err)
	}
	password := c.PostForm("password")
	link := c.PostForm("link")
	bio := c.PostForm("bio")
	avatar := c.PostForm("avatar")
	token := c.Request.Header.Get("Authorization")
	resp := service.EditUserInfo(password, link, bio, avatar, token)
	c.JSON(200, resp)
}

// TODO:checkMyDay
func UserCheckMyDay(c *gin.Context) {
	var service service.UserCheckMyDayService
	err := c.ShouldBind(&service)
	if err != nil {
		logging.Info(err)
	}
	// get token
	token := c.Request.Header.Get("Authorization")
	// get proejct cur token
	proCurToken := c.PostForm("proCurToken")
	// return response
	resp := service.UserCheckMyDay(token, proCurToken)
	c.JSON(200, resp)
}

// TODO: getProjects
// func UserGetProjectsIds(c *gin.Context) {
// 	var service service.UserGetProjectsIds
// 	err := c.ShouldBind(&service)
// 	if err != nil {
// 		logging.Info(err)
// 	}
// 	token := c.Request.Header.Get("Authorization")
// 	resp := service.GetProjectsIds(token)
// 	c.JSON(200, resp)
// }

// TODO: setAccountAvatar
func SetAccountAvatar(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
		"data":   "ok",
	})
}

func CheckToken(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "ok",
	})
}

func GetPageToken(c *gin.Context) {
	var page utils.Page
	page.NextID = "0"
	page.PageSize = 10
	proCurtoken := utils.Encode(&page)
	c.JSON(200, serializer.Response{
		Status: 200,
		Data:   proCurtoken,
	})
}
