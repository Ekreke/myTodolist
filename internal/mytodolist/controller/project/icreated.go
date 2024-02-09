package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// Projects implements IProjectController.
func (pc *ProjectController) Icreated(ctx *gin.Context) {
	log.C(ctx).Infow("project list function called")
	userid := ctx.GetInt("X-UserID")
	resp, err := pc.b.Projects().Icreated(int64(userid))
	if err != nil {
		core.WriteResponse(ctx, errno.ErrListProjectsICreated, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
