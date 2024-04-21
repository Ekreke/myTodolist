package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// MydayAi implements IUserController.
func (ctrl *UserController) MydayAi(c *gin.Context) {
	log.C(c).Infow("MydayAi function called")
	username := c.GetString("X-Username")
	resp, err := ctrl.b.Users().LoadMydayAI(c, username)
	if err != nil {
		core.WriteResponse(c, errno.ErrLoadMydayItemFailed, nil)
	}
	core.WriteResponse(c, nil, resp)
}
