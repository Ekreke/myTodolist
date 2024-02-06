package mytodolist

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/collection"
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/item"
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/project"
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/t"
	"github.com/ekreke/myTodolist/internal/mytodolist/controller/user"
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	"github.com/ekreke/myTodolist/internal/pkg/core"
	"github.com/ekreke/myTodolist/internal/pkg/errno"
	"github.com/ekreke/myTodolist/internal/pkg/middleware"
	"github.com/ekreke/myTodolist/pkg/auth"
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
	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	// controllers
	uc := user.New(store.S, authz)
	tc := t.New(store.S)
	ic := item.New(store.S)
	cc := collection.New(store.S)
	pc := project.New(store.S)
	// ic := item.New(store.S)
	tg := g.Group("/test")
	{
		tg.Use(middleware.Authn())
		tg.Use(middleware.Authz(authz))
		// test graceful shutdown , send a request and get response delay 10 seconds
		tg.GET("/lazy", tc.Lazy)
		// get logs from remote use grpc server
		tg.GET("/logs", tc.Logs)
		// get user list remote use grpc server
		tg.GET("/users", tc.GetUsers)
		// get casbin: user/rules remote use grpc server
		tg.GET("/rules", tc.Rules)
	}

	ug := g.Group("/user")
	{
		// login
		ug.POST("/login", uc.Login)

		// register
		ug.POST("/register", uc.Register)
		// signed the next request need authn
		ug.Use(middleware.Authn())
		ug.Use(middleware.Authz(authz))
		// get user info
		ug.GET("/info", uc.Info)
		// update info
		ug.POST("/info", uc.UpdateInfo)
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
		ig.POST("", ic.Create)
		// delete a item
		ig.DELETE("", ic.Delete)
		// update a item info
		ig.PUT("", ic.Update)
		// get a item info by item id
		ig.GET("", ic.Info)
		// update the item status :done
		ig.GET("/setdone", ic.SetDone)
		// update the item status :undone
		ig.GET("/setundone", ic.SetUnDone)
	}
	pg := g.Group("project")
	{
		pg.Use(middleware.Authn())
		// user join a project -> request with join code
		pg.POST("/join", pc.Join)
		// TODO:
		// list projects belong to the projects
		pg.GET("/myprojects", pc.Myprojects)
		// quit a project
		pg.GET("/quit", pc.Quit)
		// TODO:
		// create a project by root
		pg.POST("/create")
		// get project info by id
		pg.GET("/info", pc.Info)

		// TODO:create a project
		pg.POST("/create", pc.Create)
		// TODO:delete a project
		pg.DELETE("", pc.Delete)
		// TODO:update project info
		pg.POST("/update", pc.Update)
		// TODO:get all projects i created
		pg.GET("/projects", pc.projects)
		// TODO:add node to one project
		pg.POST("/addnode", pc.AddNode)
		// TODO:delete a node from one project
		pg.GET("/deletenode", pc.DeleteNode)
		// TODO:update node info
		pg.POST("/updatenode", pc.UpdateNode)
		// TODO:get nodes from project
		pg.GET("/nodes", pc.Nodes)
		// TODO:get node info
		pg.GET("/nodeinfo", pc.NodeInfo)

	}
	cg := g.Group("/collection")
	{
		cg.Use(middleware.Authn())
		// user creat a collection
		cg.POST("/create", cc.Create)
		// user delete a collection
		cg.GET("/delete", cc.Delete)
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
