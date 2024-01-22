package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *ItemController) SetUnDone(ctx *gin.Context) {
	log.C(ctx).Infow("Item setundone function called")
	username := ctx.GetString("X-Username")
	var r v1.ItemRequesWizID
	if err := ctx.ShouldBind(&r); err != nil {
		log.C(ctx).Fatalw("request bind err")
		core.WriteResponse(ctx, err, nil)
	}
	resp, err := ctrl.b.Items().SetUnDone(int(r.ItemID), username)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		log.Errorw("Set item undone failed")
	}
	core.WriteResponse(ctx, nil, resp)
}
