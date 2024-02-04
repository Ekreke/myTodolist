package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (pc *ProjectController) Join(ctx *gin.Context) {
	log.C(ctx).Infow("project join function called")
	userid := ctx.GetInt("X-UserID")

	var r *v1.JoinProjectRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := pc.b.Projects().Join(int64(userid), r.ProjectId, r.Password)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectJoin, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
