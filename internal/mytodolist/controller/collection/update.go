package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (cc *CollectionController) Update(ctx *gin.Context) {
	log.C(ctx).Infow("collection Update function called")
	username := ctx.GetString("X-Username")
	log.Debugw("the username is:", "username:", username)
	var r v1.CollectionUpdateRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := cc.b.Collections().Update(&r, username)
	if err != nil {
		log.Errorw("collection update error")
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
