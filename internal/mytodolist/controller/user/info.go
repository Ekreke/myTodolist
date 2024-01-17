package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Info(c *gin.Context) {
	log.C(c).Infow("Info function called")
	username := c.GetString("X-Username")
	resp, err := ctrl.b.Users().Info(c, username)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, resp)
}
