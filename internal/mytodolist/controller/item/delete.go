package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *ItemController) Delete(ctx *gin.Context) {
	log.C(ctx).Infow("Item delete function called")
	username := ctx.GetString("X-Username")
	var r v1.ItemRequesWizID
	if err := ctx.ShouldBind(&r); err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}
	resp, err := ctrl.b.Items().Delete(int(r.ItemID), username)
	if err != nil {
		core.WriteResponse(ctx, errno.ErrDeleteItemFailed, nil)
		return
	}
	core.WriteResponse(ctx, nil, resp)
}
