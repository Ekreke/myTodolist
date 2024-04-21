package project

import (
	"strconv"

	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (pc *ProjectController) Join(ctx *gin.Context) {
	log.C(ctx).Infow("project join function called")
	userid := ctx.GetInt("X-UserID")
	// var r *v1.JoinProjectRequest

	projectId := ctx.Param("projectid")
	pwd := ctx.Param("pwd")

	log.Infow("项目id和密码", "value:", projectId)
	intProjectId, err := strconv.Atoi(projectId)

	if err != nil {
		log.Infow("绑定projectid和pwd的时候出错")
		core.WriteResponse(ctx, errno.ErrProjectJoin, nil)
		return

	}

	// log.Infow("request:", "value:", r)
	// err = ctx.ShouldBindJSON(r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := pc.b.Projects().Join(int64(userid), int64(intProjectId), pwd)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrProjectJoin, resp)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
