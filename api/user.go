package api

import (
	"fmt"

	"github.com/ekreke/myTodolist/pkg/logging"
	"github.com/ekreke/myTodolist/service"
	"github.com/gin-gonic/gin"
)

// import "github.com/gin-gonic/gin"

// func UserLogin(c *gin.Context) {
// 	session := session.Default()
// }

func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	err := c.ShouldBind(&service)
	if err != nil {
		logging.Info(err)
	}
	// res := service.Login(user)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
	fmt.Println("gin init")
}
