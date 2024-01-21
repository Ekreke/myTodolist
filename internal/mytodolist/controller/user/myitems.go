package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) MyItems(c *gin.Context) {
	log.C(c).Infow("MyItems function called")
	var r v1.CommonRequestWizPagination
	if err := c.ShouldBind(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
	}
	username := c.GetString("X-Username")

	resp, err := ctrl.b.Users().LoadMyItems(c, &r, username)
	if err != nil {
		core.WriteResponse(c, errno.ErrLoadMydayItemFailed, nil)
	}
	core.WriteResponse(c, nil, resp)
}
