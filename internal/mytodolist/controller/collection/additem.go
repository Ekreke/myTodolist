package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (cc *CollectionController) AddItem(ctx *gin.Context) {
	log.C(ctx).Infow("additem function called")
	username := ctx.GetString("X-Username")
	log.Debugw("the username is:", "username:", username)
	var r v1.CollectionAddItemRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := cc.b.Collections().AddItem(&r, username)
	if err != nil {
		log.Errorw("collection add item error")
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
