package mytodolist

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/user"
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func installRouters(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})
	// 注册 pprof 路由
	pprof.Register(g)
	uc := user.New(store.S)
	g.POST("/login", uc.Login)
	g.POST("/register", uc.Register)
	return nil
}
