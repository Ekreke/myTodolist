package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

/*
// UpdateNode implements IProjectController.
func (pc *ProjectController) UpdateNode(ctx *gin.Context) {
	log.C(ctx).Infow("project update node function called")
	projectid := ctx.Param("projectid")
	nodeid := ctx.Param("nodeid")
	userid := ctx.GetInt("X-UserID")
	var r v1.ProjectUpdateNodeRequest
	if err := ctx.ShouldBind(&r); err != nil {
		log.C(ctx).Errorw("project update node function called", "error", err)
	}
	nTime := time.Unix(r.Deadline, 0)
	item := &model.Items{
		ItemName:     r.ItemName,
		Description:  r.Description,
		ProjectId:    r.ProjectId,
		Deadline:     nTime,
		Important:    r.Important,
		Checkpoint:   r.Checkpoint,
		Node:         r.Node,
		Myday:        r.Myday,
		CreatedTime:  r.CreatedTime,
		CollectionId: r.CollectionId,
	}
	resp, err := pc.b.Projects().UpdateNode(projectid, nodeid, userid, item)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectUpdateNode, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)

}

*/
// NodeInfo implements IProjectController.
func (pc *ProjectController) NodeInfo(ctx *gin.Context) {
	log.C(ctx).Infow("project node info function called")
	projectid := ctx.Param("projectid")
	nodeid := ctx.Param("nodeid")
	userid := ctx.GetInt("X-UserID")
	resp, err := pc.b.Projects().NodeInfo(projectid, nodeid, userid)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectInfo, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
