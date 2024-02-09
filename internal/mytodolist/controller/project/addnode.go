package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

// AddNode implements IProjectController.
func (pc *ProjectController) AddNode(ctx *gin.Context) {
	log.C(ctx).Infow("project add node function called")
	projectid := ctx.Param("projectid")
	nodeid := ctx.Param("nodeid")
	userid := ctx.GetInt("X-UserID")
	resp, err := pc.b.Projects().AddNode(projectid, nodeid, int64(userid))
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectAddNode, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
