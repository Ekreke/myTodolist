package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// Myprojects implements IProjectController.
func (pc *ProjectController) Queryallproject(ctx *gin.Context) {
	log.C(ctx).Infow("query all project function called")
	userid := ctx.GetInt("X-UserID")
	resp, err := pc.b.Projects().QueryAllProject(int64(userid))
	if err != nil {
		core.WriteResponse(ctx, errno.ErrListProjects, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
