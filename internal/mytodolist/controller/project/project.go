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
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	// icreated controller
	Icreated(ctx *gin.Context)
	// addnode controller
	AddNode(ctx *gin.Context)
	// delete node controller
	DeleteNode(ctx *gin.Context)
	// update node controller
	UpdateNode(ctx *gin.Context)
	// node info controller
	NodeInfo(ctx *gin.Context)
	Nodes(ctx *gin.Context)
	Queryallproject(ctx *gin.Context)
}

var _ IProjectController = (*ProjectController)(nil)

func New(ds store.Istore) *ProjectController {
	return &ProjectController{
		b: biz.NewBiz(ds),
	}
}
