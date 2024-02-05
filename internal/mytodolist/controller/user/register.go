package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

const defaultMethods = "(GET)|(POST)|(PUT)|(DELETE)"

func (ctrl *UserController) Register(ctx *gin.Context) {
	log.C(ctx).Infow("Register function called")
	var r v1.RegisterRequest
	if err := ctx.ShouldBind(&r); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	resp, err := ctrl.b.Users().Register(ctx, &r)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	log.Debugw("nil:", "-->", ctrl.a)
	// add policy
	_, err = ctrl.a.AddNamedPolicy("p", r.Username, "/users/"+r.Username, defaultMethods)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, err, resp)

}
