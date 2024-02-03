package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

// Quit implements IProjectController.
func (pc *ProjectController) Quit(ctx *gin.Context) {
	log.C(ctx).Infow("project quit function called")
	userid := ctx.GetInt("X-UserID")
	var r *v1.ProjectQuitRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := pc.b.Projects().Quit(int64(userid), r.ProjectId)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, resp)

}
