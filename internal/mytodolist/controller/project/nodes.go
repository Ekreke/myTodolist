package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// Nodes implements IProjectController.
func (pc *ProjectController) Nodes(ctx *gin.Context) {
	log.C(ctx).Infow("project nodes function called")
	projectid := ctx.Param("projectid")
	userid := ctx.GetInt("X-UserID")
	resp, err := pc.b.Projects().Nodes(projectid, userid)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectNodeInfo, nil)
		return
	}
	core.WriteResponse(ctx, nil, resp)

}
