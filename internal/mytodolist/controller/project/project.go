package project

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/biz"
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	b biz.IBiz
}

type IProjectController interface {
	Join(ctx *gin.Context)
	Quit(ctx *gin.Context)
	Myprojects(ctx *gin.Context)
	Info(ctx *gin.Context)
}

var _ IProjectController = (*ProjectController)(nil)

func New(ds store.Istore) *ProjectController {
	return &ProjectController{
		b: biz.NewBiz(ds),
	}
}
