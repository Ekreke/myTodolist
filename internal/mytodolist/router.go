package mytodolist

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/collection"
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/item"
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
	ic := item.New(store.S)
	cc := collection.New(store.S)
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
		// get myday items
		// record limit
		ug.GET("/myday", uc.Myday)
		// get important items
		// record limit
		ug.GET("/important", uc.Important)
		// get collections
		ug.GET("/collctions", uc.GetCollctions)
		//get items
		// contain both items and nodes
		// record limit
		ug.GET("/items", uc.Items)
		// get items user created
		// record limit
		ug.GET("/myitems", uc.MyItems)
		// get items create by projects
		// record limit
		ug.GET("/nodes", uc.Nodes)
	}
	// item group
	ig := g.Group("/item")
	{
		ig.Use(middleware.Authn())
		// create a item
		ig.POST("/create", ic.Create)
		// delete a item
		ig.GET("/delete", ic.Delete)
		// update a item info
		ig.POST("/update", ic.Update)
		// get a item info by item id
		ig.GET("/info", ic.Info)
		// update the item status :done
		ig.GET("/setdone", ic.SetDone)
		// update the item status :undone
		ig.GET("/setundone", ic.SetUnDone)
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
		// TODO:
		// create a project by root
		pg.POST("/create")
	}
	cg := g.Group("/collection")
	{
		cg.Use(middleware.Authn())
		// user creat a collection
		cg.POST("/create", cc.Create)
		// user delete a collection
		cg.GET("/delete", cc.Delete)
		// TODO:
		// user update a collection's info
		cg.POST("/update", cc.Update)
		// Todo:
		// cg.GET("/getcollectioninfo")
		// Todo:
		// cg.GET("/mycollections")
		cg.GET("/loaditems", cc.LoadItems)
		cg.GET("/additem", cc.AddItem)
		cg.GET("/deleteitem", cc.DeleteItem)
	}
	return nil
}
