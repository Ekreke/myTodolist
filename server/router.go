package server

import (
	"github.com/ekreke/myTodolist/api"
	"github.com/ekreke/myTodolist/docs"
	_ "github.com/ekreke/myTodolist/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Gin router

func NewRouter() *gin.Engine {
	r := gin.Default()
	// TODO: add middleware
	// r.Use(middleware.Cors())

	// swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		// TODO : session
		// store := cookie.NewStore([]byte(sdk.VERSION))
		// r.Use(sessions.Sessions("myssion", store))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		// user login
		v1.POST("user/login", api.UserLogin)

		// user register
		v1.POST("user/register", api.UserRegister)
	}

	return r
}
