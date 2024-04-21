package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *ItemController) Create(ctx *gin.Context) {
	log.C(ctx).Infow("Item create function called")
	username := ctx.GetString("X-Username")
	// log.Debugw("the username is:", "username:", username)
	var r v1.ItemCreateRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := ctrl.b.Items().Create(&r, username)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
