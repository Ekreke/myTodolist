package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

/*
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
*/
// DeleteNode implements IProjectController.
func (pc *ProjectController) DeleteNode(ctx *gin.Context) {
	log.C(ctx).Infow("delete node function called")
	userid := ctx.GetInt("X-UserID")
	projectid := ctx.Param("projectid")
	nodeid := ctx.Param("nodeid")
	resp, err := pc.b.Projects().DeleteNode(projectid, nodeid, int64(userid))
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectDeleteNode, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
