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
	// uercontroller
	uc := user.New(store.S)
	// item controller
	// ic := item.New(store.S)
	ug := g.Group("/user")
	{
		ug.POST("/login", uc.Login)
		ug.POST("/register", uc.Register)
		ug.Use(middleware.Authn())
		// get user info
		ug.GET("/info", uc.Info)
		// update info
		ug.POST("/updateinfo", uc.UpdateInfo)
		ug.POST("/updatepwd", uc.Updatepwd)
		// TODO:
		ug.GET("/myday", uc.Myday)
		ug.GET("/important", uc.Important)
		// TODO:
		ug.GET("/getcollction", uc.GetCollctions)
		// TODO:
		// contain both items and nodes
		ug.GET("/items")
		// TODO:
		ug.GET("/myitem")
		//TODO:
		ug.GET("/nodes")
	}
	ig := g.Group("/item")
	{
		// TODO:
		ig.POST("/creatitem")
		ig.GET("/deleteitem")
		ig.POST("/updateiteminfo")
		ig.GET("/getiteminfo")
		ig.GET("/setitemdone")
		ig.GET("/setitemundone")
	}
	pg := g.Group("project")
	{
		pg.POST("/joinproject")
		pg.POST("/myprojects")
		pg.GET("/exitproject")
	}
	cg := g.Group("/collection")
	{
		cg.POST("/creatcollection")
		cg.GET("/deletecollection")
		cg.POST("/updatecollectioninfo")
		cg.GET("/getcollectioninfo")
		cg.GET("/mycollections")
		cg.GET("/getitemsbycollections")
	}

	return nil
}
