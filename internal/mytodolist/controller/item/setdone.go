package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *ItemController) SetDone(ctx *gin.Context) {
	log.C(ctx).Infow("Item function setdone called")
	username := ctx.GetString("X-Username")
	var r v1.ItemRequesWizID
	if err := ctx.ShouldBind(&r); err != nil {
		core.WriteResponse(ctx, err, nil)
	}
	resp, err := ctrl.b.Items().SetDone(int(r.ItemID), username)
	if err != nil {
		core.WriteResponse(ctx, errno.SetItemDoneFailed, nil)
	}
	core.WriteResponse(ctx, nil, resp)
}
