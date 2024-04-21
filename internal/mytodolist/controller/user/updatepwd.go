package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Updatepwd(c *gin.Context) {
	var r *v1.UpdatepwdRequest

	err := c.ShouldBind(&r)
	if err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	if r.Prepwd == r.Newpwd {
		log.C(c).Infow("previous password eauqls new password!")
		core.WriteResponse(c, errno.ErrPwdEquals, nil)
		return
	}

	username := c.GetString("X-Username")
	resp, err := ctrl.b.Users().UpdatePwd(c, username, r.Prepwd, r.Newpwd)
	if err != nil {
		log.C(c).Infow("controlle update pwd run failed")
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, resp)

}
