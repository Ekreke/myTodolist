package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

// Update implements IProjectController.
func (pc *ProjectController) Update(ctx *gin.Context) {
	log.C(ctx).Infow("project update function called")
	var r v1.ProjectUpdateRequest
	userid := ctx.GetInt("X-UserID")
	if err := ctx.ShouldBind(&r); err != nil {
		log.C(ctx).Errorw("project update function called", "error", err)
	}
	resp, err := pc.b.Projects().Update(r.Description, r.Endtime, r.Name, int64(userid))
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectUpdate, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
