package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (pc *ProjectController) Info(ctx *gin.Context) {
	log.C(ctx).Infow("project info function called")
	var r *v1.ProjectInfoRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := pc.b.Projects().Info(r.ProjectId)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectInfo, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
