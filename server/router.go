package server

import (
	"fmt"

	"github.com/ekreke/myTodolist/api"
	_ "github.com/ekreke/myTodolist/docs"
	"github.com/ekreke/myTodolist/pkg/util/sdk"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Gin router

func NewRouter() *gin.Engine {
	r := gin.Default()
	// TODO: add middleware
	// r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	store := cookie.NewStore([]byte(sdk.VERSION))
	r.Use(sessions.Sessions("myssion", store))

	// router group
	v1 := r.Group("/api/v1")
	{
		fmt.Println("init")
		v1.GET("user/login", api.UserLogin)
	}
	return r
}
