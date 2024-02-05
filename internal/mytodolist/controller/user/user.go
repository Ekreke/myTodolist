package user

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/biz"
	"github.com/ekreke/myTodolist/pkg/auth"
	"github.com/gin-gonic/gin"

	"github.com/ekreke/myTodolist/internal/mytodolist/store"
)

type UserController struct {
	// casbin
	a *auth.Authz
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
	GetCollctions(c *gin.Context)
}

var _ IUserController = (*UserController)(nil)

func New(ds store.Istore, a *auth.Authz) *UserController {
	return &UserController{b: biz.NewBiz(ds), a: a}
}
