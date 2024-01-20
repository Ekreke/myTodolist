package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) GetCollctions(c *gin.Context) {
	log.C(c).Infow("Get Collections function called")
	username := c.GetString("X-Username")
	resp, err := ctrl.b.Users().GetCollctions(c, username)
	if err != nil {
		log.Errorw("get collectinos failed")
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, resp)
}
