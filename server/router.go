package server

import (
	"fmt"

	"github.com/ekreke/myTodolist/api"
	"github.com/ekreke/myTodolist/pkg/util/sdk"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Gin router

func NewRouter() *gin.Engine {
	r := gin.Default()
	// TODO: add middleware
	// r.Use(middleware.Cors())
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
