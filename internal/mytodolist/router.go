package mytodolist

import (
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/gin-gonic/gin"
)

func installRouters(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		//
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

}
