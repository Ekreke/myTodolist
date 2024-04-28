package user

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/biz"
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
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

// ImportantAI implements IUserController.
func (uc *UserController) ImportantAI(c *gin.Context) {
	log.C(c).Infow("important ai function called")
	username := c.GetString("X-Username")
	resp, err := uc.b.Users().LoadImportantAI(c, username)
	if err != nil {
		core.WriteResponse(c, errno.ErrLoadMydayItemFailed, nil)
	}
	core.WriteResponse(c, nil, resp)
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
	MydayAi(c *gin.Context)
	ImportantAI(c *gin.Context)
}

var _ IUserController = (*UserController)(nil)

func New(ds store.Istore, a *auth.Authz) *UserController {
	return &UserController{b: biz.NewBiz(ds), a: a}
}
