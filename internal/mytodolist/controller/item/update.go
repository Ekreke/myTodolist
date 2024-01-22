package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *ItemController) Update(ctx *gin.Context) {
	log.C(ctx).Infow("Item Update function called")
	username := ctx.GetString("X-Username")
	var r v1.ItemUpdateRequest
	err := ctx.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
	}
	resp, err := ctrl.b.Items().Update(&r, username)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrUpdateItemFailed, nil)
		log.C(ctx).Errorw("Item update failed", "the err is :", err)
	}
	core.WriteResponse(ctx, nil, resp)
}
