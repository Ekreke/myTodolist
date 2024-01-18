package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

// update info only can update username / bio / link / avatar
// TODO: update avatar
func (ctrl *UserController) UpdateInfo(ctx *gin.Context) {
	log.C(ctx).Infow("UpdateInfo function called")
	var r v1.UpdateInfoRequest
	if err := ctx.ShouldBind(&r); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	username := ctx.GetString("X-Username")
	resp, err := ctrl.b.Users().UpdateInfo(ctx, &r, username)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
	}
	core.WriteResponse(ctx, err, resp)

}
