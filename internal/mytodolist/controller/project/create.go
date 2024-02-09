package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

// Create implements IProjectController.
func (pc *ProjectController) Create(ctx *gin.Context) {
	log.C(ctx).Infow("project create function called")
	userid := ctx.GetInt("X-UserID")
	var r v1.ProjectCreateRequest
	if err := ctx.ShouldBind(&r); err != nil {
		log.C(ctx).Errorw("project create function called", "error", err)
	}
	resp, err := pc.b.Projects().Create(r.Description, r.Endtime, r.Name, int64(userid))
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectCreate, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
