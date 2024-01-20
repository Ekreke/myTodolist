package mytodolist

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/t"
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
	tc := t.New(store.S)
	// item controller
	// ic := item.New(store.S)
	tg := g.Group("/test")
	{
		// test graceful shutdown , send a request and get response delay 10 seconds
		tg.GET("/lazy", tc.Lazy)
	}

	ug := g.Group("/user")
	{
		// login
		ug.POST("/login", uc.Login)

		// register
		ug.POST("/register", uc.Register)
		// signed the next request need authn
		ug.Use(middleware.Authn())
		// get user info
		ug.GET("/info", uc.Info)
		// update info
		ug.POST("/updateinfo", uc.UpdateInfo)
		// update user pwd
		ug.POST("/updatepwd", uc.Updatepwd)
		// TODO:
		// get myday items
		ug.GET("/myday", uc.Myday)
		// get important items
		ug.GET("/important", uc.Important)
		// get collections
		ug.GET("/getcollction", uc.GetCollctions)
		// TODO:
		//get items
		// contain both items and nodes
		ug.GET("/items")
		// TODO:
		// get items user created
		ug.GET("/myitem")
		// TODO:
		// get items create by projects
		ug.GET("/nodes")
	}
	// item group
	ig := g.Group("/item")
	{
		// TODO:
		// create a item
		ig.POST("/create")
		// TODO:
		// delete a item
		ig.GET("/delete")
		// TODO:
		// update a item info
		ig.POST("/updateinfo")
		// TODO:
		// get a item info
		ig.GET("/getinfo")
		// TODO:
		// update the item status :done
		ig.GET("/setdone")
		// TODO:
		// update the item status :undone
		ig.GET("/setundone")
	}
	pg := g.Group("project")
	{
		// TODO:
		// user join a project -> request with join code
		pg.POST("/join")
		// TODO:
		// list projects belong to the projects
		pg.POST("/myprojects")
		// TODO:
		// quit a project
		pg.GET("/quit")
	}
	cg := g.Group("/collection")
	{
		// TODO:
		// user creat a collection
		cg.POST("/creatcollection")
		// user delete a collection
		// TODO:
		cg.GET("/deletecollection")
		// TODO:
		// user update a collection's info
		cg.POST("/updatecollectioninfo")
		// Todo:
		// cg.GET("/getcollectioninfo")
		// Todo:
		// cg.GET("/mycollections")
		// TODO:
		cg.GET("/getitems")
		// TODO:
		cg.GET("/additems")
		// TODO:
		cg.GET("/deleteitem")
	}

	return nil
}
