package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")
	var r v1.LoginRequest
	if err := c.ShouldBind(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, resp)
}
