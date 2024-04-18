package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Myday(c *gin.Context) {
	log.C(c).Infow("Myday function called")
	var r v1.MydayRequest
	if err := c.ShouldBind(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
	}
	username := c.GetString("X-Username")
	// log.Infow("Myday Request: %+v", r)
	resp, err := ctrl.b.Users().LoadMydayItems(c, &r, username)
	if err != nil {
		core.WriteResponse(c, errno.ErrLoadMydayItemFailed, nil)
	}
	core.WriteResponse(c, nil, resp)
}
