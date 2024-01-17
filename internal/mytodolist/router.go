package mytodolist

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/user"
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func installRouters(g *gin.Engine) error {
	g.NoRoute(func(c *gin.Context) {
		//
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})
	// 注册 pprof 路由
	pprof.Register(g)
	uc := user.New(store.S)
	ug := g.Group("/user")
	{
		ug.POST("/login", uc.Login)
		ug.POST("/register", uc.Register)
		ug.Use(middleware.Authn())
		// get user info
		ug.GET("/info", uc.Info)
		// TODO:
		ug.POST("/updateinfo")
		// TODO:
		ug.POST("/updatepwd")
		// TODO:
		ug.POST("/updateavat/ar")
		// TODO:
		ug.GET("/myday")
		// TODO:
		ug.GET("/important")
		// TODO:
		ug.GET("/getcollction")
		// TODO:
		// contain both items and nodes
		ug.GET("/items")
		// TODO:
		ug.GET("/myitem")
		//TODO:
		ug.GET("/nodes")

	}
	return nil
}
