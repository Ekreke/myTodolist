package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Important(c *gin.Context) {
	log.C(c).Infow("Important function called")
	var r v1.ImportantRequest
	err := c.ShouldBindQuery(&r)
	if err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
	}
	username := c.GetString("X-Username")
	resp, token, err := ctrl.b.Users().LoadImportantItems(c, &r, username)
	resp.CursorToken = string(token)
	if err != nil {
		core.WriteResponse(c, errno.ErrLoadImportantItemFailed, nil)
	}
	core.WriteResponse(c, nil, resp)
}
