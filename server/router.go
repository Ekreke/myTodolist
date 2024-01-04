package server

import (
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
		v1.POST("user/login")
	}
	return r
}
