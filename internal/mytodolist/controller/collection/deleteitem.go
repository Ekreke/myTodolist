package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (cc *CollectionController) DeleteItem(ctx *gin.Context) {
	log.C(ctx).Infow("delete item function called")
	username := ctx.GetString("X-Username")
	log.Debugw("the username is:", "username:", username)
	var r v1.CollectionDeleteItemRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := cc.b.Collections().DeleteItem(&r, username)
	if err != nil {
		log.Errorw("collection delete item error")
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
