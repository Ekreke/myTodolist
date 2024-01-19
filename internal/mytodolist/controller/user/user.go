package user

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/biz"
	"github.com/gin-gonic/gin"

	"github.com/ekreke/myTodolist/internal/mytodolist/store"
)

type UserController struct {
	// casbin
	// grpc
	b biz.IBiz
}

type IUserController interface {
	UpdateInfo(ctx *gin.Context)
	Login(c *gin.Context)
	Register(c *gin.Context)
	Info(c *gin.Context)
	Myday(c *gin.Context)
	Important(c *gin.Context)
	Updatepwd(c *gin.Context)
}

var _ IUserController = (*UserController)(nil)

func New(ds store.Istore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}
