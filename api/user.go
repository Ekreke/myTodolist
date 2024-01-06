package api

import (
	"github.com/ekreke/myTodolist/pkg/logging"
	"github.com/ekreke/myTodolist/service"
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
