package api

import (
	"github.com/ekreke/myTodolist/pkg/logging"
	"github.com/ekreke/myTodolist/service"
	"github.com/gin-gonic/gin"
)

// import "github.com/gin-gonic/gin"

//	func UserLogin(c *gin.Context) {
//		session := session.Default()
//	}
//
// @Summary userlogin
// @Schemes
// @Description user login
// @Tags User
// @Param Id query int true "Id"     参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	err := c.ShouldBind(&service)
	if err != nil {
		logging.Info(err)
	}
	username := c.PostForm("username")
	password := c.PostForm("password")

	if flag := service.Login(username, password); flag == true {
		c.JSON(200, res)
	}
}
