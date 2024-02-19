package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// Delete implements IProjectController.
func (pc *ProjectController) Delete(ctx *gin.Context) {
	log.C(ctx).Infow("project delete function called")
	projectid := ctx.Param("projectid")
	userid := ctx.GetInt("X-UserID")
	resp, err := pc.b.Projects().Delete(projectid, int64(userid))
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectDelete, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
