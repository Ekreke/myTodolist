package project

import (
	"time"

	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

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
	cTime := time.Unix(time.Now().Unix(), 0)
	item := &model.Items{
		ItemName:     r.ItemName,
		Description:  r.Description,
		ProjectId:    r.ProjectId,
		Deadline:     nTime,
		Important:    r.Important,
		Checkpoint:   r.Checkpoint,
		Node:         r.Node,
		Myday:        r.Myday,
		CreatedTime:  cTime,
		CollectionId: r.CollectionId,
	}
	resp, err := pc.b.Projects().UpdateNode(projectid, nodeid, userid, item)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectUpdateNode, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)

}
